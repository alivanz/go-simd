package scanner

import (
	"bytes"
	"regexp"
	"strings"
)

var (
	regTypedef       = regexp.MustCompile(`typedef [\w\s\(\){}]+? ([0-9a-z]\w+?_t);`)
	regTypedefStruct = regexp.MustCompile(`typedef struct \w+ {.+?(\w+);`)
	regTypedefAttr   = regexp.MustCompile(`typedef __attribute__.+? (\w+);`)
	regFunction      = regexp.MustCompile(`(\w+) (\w+)\(([\w\s,_]*?)\) {.*?}`)
	regArg           = regexp.MustCompile(`(\w+)(\s+\w+)?`)
)

type ScanResult struct {
	Types     []Type
	Functions []Function
}

func Scan(raw []byte) (*ScanResult, error) {
	var buf bytes.Buffer
	for _, line := range strings.Split(string(raw), "\n") {
		if !strings.HasPrefix(line, "#") {
			buf.WriteString(line)
		}
	}
	var result ScanResult
	// types
	result.Types = joinAll(
		transform(
			regTypedef.FindAllStringSubmatch(buf.String(), -1),
			func(i int, e []string) Type {
				return Type{
					Name: e[1],
					Full: e[0],
				}
			},
		),
		transform(
			regTypedefStruct.FindAllStringSubmatch(buf.String(), -1),
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
		regFunction.FindAllStringSubmatch(buf.String(), -1),
		func(i int, match []string) Function {
			sargs := regArg.FindAllStringSubmatch(match[3], -1)
			return Function{
				Name: match[2],
				Return: Type{
					Name: match[1],
					Full: match[0],
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
