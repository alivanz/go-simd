package main

import (
	"fmt"
	"io"
	"strings"
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
