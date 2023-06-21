package scanner

import (
	"fmt"
	"io"
	"strings"
)

type Package struct {
	Name      string
	C         []string
	Types     []Type
	Functions []Function
}

func (pkg *Package) WriteTo(w io.Writer) error {
	if _, err := fmt.Fprintf(w, "package %s\n", pkg.Name); err != nil {
		return err
	}
	if len(pkg.C) > 0 {
		if _, err := fmt.Fprintf(w, "\n/*\n%s\n*/\nimport \"C\"\n", strings.Join(pkg.C, "\n")); err != nil {
			return err
		}
	}
	for _, t := range pkg.Types {
		if strings.HasPrefix(t.C(), "__") {
			continue
		}
		if strings.Contains(t.C(), "float16") {
			continue
		}
		if _, err := io.WriteString(w, t.Declare()); err != nil {
			return err
		}
	}
funcs:
	for _, fn := range pkg.Functions {
		for _, blacklist := range []string{
			"f16",
			"vcmla",
		} {
			if strings.Contains(fn.Name, blacklist) {
				continue funcs
			}
		}
		if _, err := io.WriteString(w, fn.Declare()); err != nil {
			return err
		}
	}
	return nil
}
