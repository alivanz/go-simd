package writer

import (
	"fmt"
	"io"
	"strings"

	"github.com/alivanz/go-simd/generator/types"
	"github.com/alivanz/go-simd/generator/utils"
	"github.com/iancoleman/strcase"
)

func DeclareFunc(w io.Writer, f types.Function, typePkg string) error {
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
	fmt.Fprintf(w, "%s", strings.Join(utils.Transform(f.Args, func(i int, t types.Type) string {
		return fmt.Sprintf("v%d %s", i, t.Go(typePkg))
	}), ", "))
	if f.Return == nil {
		fmt.Fprintf(w, ") {\n")
	} else {
		fmt.Fprintf(w, ") %s {\n", f.Return.Go(typePkg))
	}
	if f.Return == nil {
		fmt.Fprintf(w, "\tC.%s(%s)\n", f.Name, strings.Join(utils.Transform(f.Args, func(i int, t types.Type) string {
			return fmt.Sprintf("v%d", i)
		}), ", "))
	} else {
		fmt.Fprintf(w, "\treturn C.%s(%s)\n", f.Name, strings.Join(utils.Transform(f.Args, func(i int, t types.Type) string {
			return fmt.Sprintf("v%d", i)
		}), ", "))
	}
	fmt.Fprintf(w, "}\n")
	return nil
}

func DeclareFuncBypass(w io.Writer, f types.Function, typePkg string) error {
	fmt.Fprintf(w, "\n")
	if len(f.Comment) > 0 {
		fmt.Fprintf(w, "// %s\n", f.Comment)
	} else {
		fmt.Fprintf(w, "// %s\n", f.Name)
	}
	fmt.Fprintf(w, "//\n")
	fmt.Fprintf(w, "//go:linkname %s %s\n", strcase.ToCamel(f.Name), strcase.ToCamel(f.Name))
	fmt.Fprintf(w, "//go:noescape\n")
	fmt.Fprintf(w, "func %s(", strcase.ToCamel(f.Name))
	if f.Return != nil {
		fmt.Fprintf(w, "r *%s, ", f.Return.Go(typePkg))
	}
	fmt.Fprintf(w, "%s)\n", strings.Join(utils.Transform(f.Args, func(i int, t types.Type) string {
		return fmt.Sprintf("v%d *%s", i, t.Go(typePkg))
	}), ", "))
	return nil
}

func RewriteC(w io.Writer, f types.Function) error {
	fmt.Fprintf(w, "void %s(", strcase.ToCamel(f.Name))
	if f.Return != nil {
		fmt.Fprintf(w, "%s* r, ", f.Return.Name)
	}
	fmt.Fprintf(w, "%s) { ", strings.Join(utils.Transform(f.Args, func(i int, t types.Type) string {
		return fmt.Sprintf("%s* v%d", t.Name, i)
	}), ", "))
	if f.Return != nil {
		fmt.Fprintf(w, "*r = ")
	}
	fmt.Fprintf(w, "%s(%s); }\n", f.Name, strings.Join(utils.Transform(f.Args, func(i int, t types.Type) string {
		return fmt.Sprintf("*v%d", i)
	}), ", "))
	return nil
}
