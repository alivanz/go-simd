package scanner

import (
	"fmt"
	"io"
	"regexp"
	"strings"

	"github.com/iancoleman/strcase"
)

type Function struct {
	Name      string
	Args      []Type
	Return    Type
	Attribute string
	Comment   string
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

var (
	regTarget = regexp.MustCompile(`__target__\("([a-z0-9\s,]+)"\)`)
)

func (f *Function) Target() string {
	match := regTarget.FindStringSubmatch(f.Attribute)
	if match == nil {
		return ""
	}
	return match[1]
}

func (f *Function) Declare(w io.Writer) error {
	var comment string
	if len(f.Comment) > 0 {
		comment = f.Comment
	} else {
		comment = f.Name
	}
	if len(f.Attribute) > 0 {
		comment += fmt.Sprintf("\n// %s", f.Attribute)
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
