package scanner

import (
	"fmt"
	"strings"

	"github.com/iancoleman/strcase"
)

type Function struct {
	Name   string
	Args   []Type
	Return Type
}

type Arg struct {
	Name string
	Type string
}

const funcTemplate = `func %s(%s) %s {
	return %s(C.%s(%s))
}`

func (f *Function) String() string {
	return fmt.Sprintf(
		funcTemplate,
		strcase.ToCamel(f.Name),
		strings.Join(transform(f.Args, func(i int, t Type) string {
			return fmt.Sprintf("v%d %s", i, t.GoString())
		}), ", "),
		f.Return.GoString(),
		f.Return.GoString(),
		f.Name,
		strings.Join(transform(f.Args, func(i int, t Type) string {
			return fmt.Sprintf("C.%s(v%d)", string(t), i)
		}), ", "),
	)
}
