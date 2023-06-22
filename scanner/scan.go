package scanner

import (
	"bytes"
	"regexp"
)

var (
	name             = `(\w+?)`
	args             = `([\w\s,_]*?)`
	attr             = `(?:__attribute__\(\(` + `([\w\s,\(\)"]+?)` + `\)\))`
	regTypedefSimple = regexp.MustCompile(`typedef\s+` + attr + `?[\w\s]+? ` + name + `\s*` + attr + `?;`)
	regTypedefStruct = regexp.MustCompile(`typedef struct \w+? {.+?}\s*?` + name + `;`)
	regFunction      = regexp.MustCompile(name + `\s+` + attr + `?\s*` + name + `\s*\(` + args + `\)` + `\s*` + `{.*?}`)
	regArg           = regexp.MustCompile(`(\w+)(\s+\w+)?`)
	regWhitespace    = regexp.MustCompile(`\s+`)
	regComma         = regexp.MustCompile(`\s*,\s*`)
)

type ScanResult struct {
	Types     []Type
	Functions []Function
}

func Scan(raw []byte) (*ScanResult, error) {
	var buf bytes.Buffer
	// filter #
	for _, line := range bytes.Split(raw, []byte("\n")) {
		if !bytes.HasPrefix(line, []byte("#")) {
			buf.Write(line)
		}
	}
	// remove duplicates whitespace
	raw = regWhitespace.ReplaceAll(buf.Bytes(), []byte(" "))
	s := string(raw)
	var result ScanResult
	// types
	result.Types = joinAll(
		transform(
			regTypedefSimple.FindAllStringSubmatch(s, -1),
			func(i int, e []string) Type {
				return Type{
					Name:       e[2],
					Full:       e[0],
					Attributes: commaSplit(e[1], e[3]),
				}
			},
		),
		transform(
			regTypedefStruct.FindAllStringSubmatch(s, -1),
			func(i int, e []string) Type {
				return Type{
					Name: e[1],
					Full: e[0],
				}
			},
		),
	)
	// functions
	result.Functions = transform(
		regFunction.FindAllStringSubmatch(s, -1),
		func(i int, match []string) Function {
			sargs := regArg.FindAllStringSubmatch(match[4], -1)
			return Function{
				Name:      match[3],
				Attribute: match[2],
				Return: Type{
					Name: match[1],
					Full: match[1],
				},
				Args: transform(sargs, func(i int, e []string) Type {
					return Type{
						Name: e[1],
						Full: e[0],
					}
				}),
			}
		},
	)
	return &result, nil
}
