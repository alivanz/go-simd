package scanner

import (
	"bytes"
	"regexp"

	"github.com/alivanz/go-simd/generator/types"
	"github.com/alivanz/go-simd/generator/utils"
)

var (
	name             = `(\w+?)`
	args             = `([\w\s,_]*?)`
	attr             = `(?:__attribute__\(\(` + `([\w\s,\(\)"]+?)` + `\)\))`
	regTypedefSimple = regexp.MustCompile(`typedef\s+` + attr + `?[\w\s]+? ` + name + `\s*` + attr + `?;`)
	regTypedefStruct = regexp.MustCompile(`typedef struct \w+? {.+?}\s*?` + name + `;`)
	regFunction      = regexp.MustCompile(name + `\s+` + attr + `?\s*` + name + `\s*\(` + args + `\)` + `\s*` + `{.*?}`)
	regArg           = regexp.MustCompile(`\s*(([\w\s]+)\s(?:\w+))`)
	regWhitespace    = regexp.MustCompile(`\s+`)
	regComma         = regexp.MustCompile(`\s*,\s*`)
	regLongLong      = regexp.MustCompile(`long\s+long`)
	regULongLong     = regexp.MustCompile(`unsigned\s+long\s+long`)
	regUlong         = regexp.MustCompile(`unsigned\s+long`)
	regUint          = regexp.MustCompile(`unsigned\s+int`)
	regUshort        = regexp.MustCompile(`unsigned\s+short`)
	regUchar         = regexp.MustCompile(`unsigned\s+char`)
)

type ScanResult struct {
	Types     []types.Type
	Functions []types.Function
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
	// replace known types
	raw = regLongLong.ReplaceAll(raw, []byte("longlong"))
	raw = regULongLong.ReplaceAll(raw, []byte("ulonglong"))
	raw = regUlong.ReplaceAll(raw, []byte("ulong"))
	raw = regUint.ReplaceAll(raw, []byte("uint"))
	raw = regUshort.ReplaceAll(raw, []byte("ushort"))
	raw = regUchar.ReplaceAll(raw, []byte("uchar"))
	s := string(raw)
	var result ScanResult
	// types
	result.Types = utils.Merge(
		utils.Transform(
			regTypedefSimple.FindAllStringSubmatch(s, -1),
			func(i int, e []string) types.Type {
				return types.Type{
					Name:       e[2],
					Full:       e[0],
					Attributes: commaSplit(e[1], e[3]),
				}
			},
		),
		utils.Transform(
			regTypedefStruct.FindAllStringSubmatch(s, -1),
			func(i int, e []string) types.Type {
				return types.Type{
					Name: e[1],
					Full: e[0],
				}
			},
		),
	)
	// functions
	result.Functions = utils.Transform(
		regFunction.FindAllStringSubmatch(s, -1),
		func(i int, match []string) types.Function {
			var args []types.Type
			for _, arg := range regArg.FindAllStringSubmatch(match[4], -1) {
				if arg[2] == "void" {
					continue
				}
				args = append(args, types.Type{
					Name: arg[2],
					Full: arg[1],
				})
			}
			var ret *types.Type
			if match[1] != "void" {
				ret = &types.Type{
					Name: match[1],
					Full: match[1],
				}
			}
			return types.Function{
				Name:      match[3],
				Attribute: match[2],
				Return:    ret,
				Args:      args,
			}
		},
	)
	return &result, nil
}
