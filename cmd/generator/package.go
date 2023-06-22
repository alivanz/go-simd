package main

import (
	"fmt"
	"io"
	"strings"

	"github.com/alivanz/go-simd/scanner"
)

func Package(w io.Writer, pkg string) error {
	_, err := fmt.Fprintf(w, "package %s\n", pkg)
	return err
}

func Import(w io.Writer, pkgs []string) error {
	if len(pkgs) == 0 {
		return nil
	}
	_, err := fmt.Fprintf(w, "\nimport (\n\t%s\n)\n", strings.Join(pkgs, "\n\t"))
	return err
}

func ImportC(w io.Writer, s string) error {
	_, err := fmt.Fprintf(w, "\n/*\n%s\n*/\nimport \"C\"\n", s)
	return err
}

func Types(w io.Writer, types []scanner.Type) error {
types:
	for _, t := range types {
		for _, blacklist := range []string{
			"__darwin",
			"__int",
			"__uint",
			"__mm_storeh",
			"_tile",
			"_aligned",
			"float16",
		} {
			if strings.Contains(t.Name, blacklist) {
				continue types
			}
		}
		if err := t.Declare(w); err != nil {
			return err
		}
	}
	return nil
}

func Funcs(w io.Writer, funcs []scanner.Function) error {
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
		if err := fn.Declare(w); err != nil {
			return err
		}
	}
	return nil
}
