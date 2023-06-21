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
	return C.%s(%s)
}
`

func (f *Function) String() string {
	return fmt.Sprintf(
		funcTemplate,
		strcase.ToCamel(f.Name),
		strings.Join(transform(f.Args, func(i int, t Type) string {
			return fmt.Sprintf("v%d %s", i, t.Go())
		}), ", "),
		f.Return.Go(),
		f.Name,
		strings.Join(transform(f.Args, func(i int, t Type) string {
			return fmt.Sprintf("v%d", i)
		}), ", "),
	)
}
