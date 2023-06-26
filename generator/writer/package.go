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
	for _, t := range types {
		if t.Blacklisted() {
			continue
		}
		if err := DeclareType(w, t); err != nil {
			return err
		}
	}
	return nil
}

func Funcs(w io.Writer, funcs []types.Function, typePkg string) error {
	for _, fn := range funcs {
		if fn.Blacklisted() {
			continue
		}
		if err := DeclareFunc(w, fn, typePkg); err != nil {
			return err
		}
	}
	return nil
}
