package scanner

import (
	"strings"

	"github.com/iancoleman/strcase"
)

type Type string

func (t Type) GoString() string {
	switch t {
	default:
		s := strings.TrimSuffix(string(t), "_t")
		return strcase.ToCamel(s)
	case "float32_t":
		return "float32"
	case "float64_t":
		return "float64"
	case "uint8_t":
		return "uint8"
	case "uint16_t":
		return "uint16"
	case "uint32_t":
		return "uint32"
	case "uint64_t":
		return "uint64"
	case "int8_t":
		return "int8"
	case "int16_t":
		return "int16"
	case "int32_t":
		return "int32"
	case "int64_t":
		return "int64"
	}
}
