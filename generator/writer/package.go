package writer

import (
	"fmt"
	"io"
	"strings"

	"github.com/alivanz/go-simd/generator/types"
)

func Package(w io.Writer, pkg string) error {
	_, err := fmt.Fprintf(w, "package %s\n", pkg)
	return err
}

func Import(w io.Writer, pkgs []string) error {
	if len(pkgs) == 0 {
		return nil
	}
	_, err := fmt.Fprintf(w, "\nimport (\n\t\"%s\"\n)\n", strings.Join(pkgs, "\"\n\t\""))
	return err
}

func ImportC(w io.Writer, fn func(w io.Writer) error) error {
	if _, err := fmt.Fprintf(w, "\n/*\n"); err != nil {
		return err
	}
	if err := fn(w); err != nil {
		return err
	}
	if _, err := fmt.Fprintf(w, "\n*/\nimport \"C\"\n"); err != nil {
		return err
	}
	return nil
}

func Types(w io.Writer, types []types.Type) error {
types:
	for _, t := range types {
		for _, blacklist := range []string{
			"__darwin",
			"__int",
			"__uint",
			"__mm_storeh",
			"_tile",
			"_aligned",
			// float16
			"float16",
			"f16",
			"v8bf",
			"v8hf",
			"m128h",
			"m128bh",
			// windows?
			"crt",
			"_pi_",
			"mbstate_t",
		} {
			if strings.Contains(t.Name, blacklist) {
				continue types
			}
		}
		if err := DeclareType(w, t); err != nil {
			return err
		}
	}
	return nil
}

func Funcs(w io.Writer, funcs []types.Function, typePkg string) error {
funcs:
	for _, fn := range funcs {
		for _, blacklist := range []string{
			"f16",
			"vcmla",
			"__extension__",
		} {
			if strings.Contains(fn.Name, blacklist) {
				continue funcs
			}
		}
		if err := DeclareFunc(w, fn, typePkg); err != nil {
			return err
		}
	}
	return nil
}
