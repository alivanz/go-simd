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
types:
	for _, t := range pkg.Types {
		for _, blacklist := range []string{
			"__darwin",
			"__int",
			"__uint",
			"__mm_storeh",
			"_tile",
			"_aligned",
		} {
			if strings.Contains(t.Name, blacklist) {
				continue types
			}
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
			"__extension__",
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
