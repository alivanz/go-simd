package scanner

import (
	"bytes"
	"regexp"
)

var (
	name             = `(\w+?)`
	attr             = `(?:__attribute__\(\([\w\s,\(\)]+?\)\))`
	regTypedefSimple = regexp.MustCompile(`typedef\s+` + attr + `?[\w\s]+? ` + name + `\s*` + attr + `?;`)
	regTypedefStruct = regexp.MustCompile(`typedef struct \w+? {.+?}\s*?` + name + `;`)
	regFunction      = regexp.MustCompile(`(\w+) (\w+)\(([\w\s,_]*?)\) {.*?}`)
	regArg           = regexp.MustCompile(`(\w+)(\s+\w+)?`)
	regWhitespace    = regexp.MustCompile(`\s+`)
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
					Name: e[1],
					Full: e[0],
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
			sargs := regArg.FindAllStringSubmatch(match[3], -1)
			return Function{
				Name: match[2],
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
