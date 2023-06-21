package scanner

import (
	"bytes"
	"regexp"
	"strings"
)

var (
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
	var (
		result   ScanResult
		submatch [][]string
	)
	// types
	submatch = regTypedefStruct.FindAllStringSubmatch(buf.String(), -1)
	result.Types = transform(submatch, func(i int, e []string) Type {
		return Type(e[1])
	})
	submatch = regTypedefAttr.FindAllStringSubmatch(buf.String(), -1)
	result.Types = append(result.Types, transform(submatch, func(i int, e []string) Type {
		return Type(e[1])
	})...)
	// functions
	submatch = regFunction.FindAllStringSubmatch(buf.String(), -1)
	result.Functions = make([]Function, len(submatch))
	for i, match := range submatch {
		sargs := regArg.FindAllStringSubmatch(match[3], -1)
		result.Functions[i] = Function{
			Name:   match[2],
			Return: Type(match[1]),
			Args: transform(sargs, func(i int, e []string) Type {
				return Type(e[1])
			}),
		}
	}
	return &result, nil
}
