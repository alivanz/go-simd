package scanner

import (
	"fmt"
	"io"
	"regexp"
	"strings"

	"github.com/alivanz/go-simd/generator/utils"
	"github.com/iancoleman/strcase"
)

type Function struct {
	Name      string
	Args      []Type
	Return    *Type
	Attribute string
	Comment   string
}

type Arg struct {
	Name string
	Type string
}

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

func (f *Function) Declare(w io.Writer, typePkg string) error {
	fmt.Fprintf(w, "\n")
	if len(f.Comment) > 0 {
		fmt.Fprintf(w, "// %s\n", f.Comment)
	} else {
		fmt.Fprintf(w, "// %s\n", f.Name)
	}
	// if len(f.Attribute) > 0 {
	// 	fmt.Fprintf(w, "// %s\n", f.Attribute)
	// }
	fmt.Fprintf(w, "func %s(", strcase.ToCamel(f.Name))
	fmt.Fprintf(w, "%s", strings.Join(utils.Transform(f.Args, func(i int, t Type) string {
		return fmt.Sprintf("v%d %s", i, t.Go(typePkg))
	}), ", "))
	if f.Return == nil {
		fmt.Fprintf(w, ") {\n")
	} else {
		fmt.Fprintf(w, ") %s {\n", f.Return.Go(typePkg))
	}
	if f.Return == nil {
		fmt.Fprintf(w, "\tC.%s(%s)\n", f.Name, strings.Join(utils.Transform(f.Args, func(i int, t Type) string {
			return fmt.Sprintf("v%d", i)
		}), ", "))
	} else {
		fmt.Fprintf(w, "\treturn C.%s(%s)\n", f.Name, strings.Join(utils.Transform(f.Args, func(i int, t Type) string {
			return fmt.Sprintf("v%d", i)
		}), ", "))
	}
	fmt.Fprintf(w, "}\n")
	return nil
}
