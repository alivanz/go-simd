package scanner

import (
	"fmt"
	"io"
	"strings"

	"github.com/iancoleman/strcase"
)

type Function struct {
	Name       string
	Args       []Type
	Return     Type
	Attributes []string
	Comment    string
}

type Arg struct {
	Name string
	Type string
}

const funcTemplate = `
// %s
func %s(%s) %s {
	return C.%s(%s)
}
`

func (f *Function) Declare(w io.Writer) error {
	var comment string
	if len(f.Comment) > 0 {
		comment = f.Comment
	} else {
		comment = f.Name
	}
	if len(f.Attributes) > 0 {
		comment += fmt.Sprintf("\n// %s", strings.Join(f.Attributes, ", "))
	}
	_, err := fmt.Fprintf(
		w,
		funcTemplate,
		comment,
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
	return err
}
