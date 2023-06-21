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
	if _, err := fmt.Fprintf(w, "package %s\n\n", pkg.Name); err != nil {
		return err
	}
	if len(pkg.C) > 0 {
		if _, err := fmt.Fprintf(w, "/*\n%s\n*/\nimport \"C\"\n\n", strings.Join(pkg.C, "\n")); err != nil {
			return err
		}
	}
	for _, t := range pkg.Types {
		if strings.Contains(string(t), "float16") {
			continue
		}
		if _, err := fmt.Fprintf(w, "type %s C.%s\n\n", t.GoString(), t); err != nil {
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
		if _, err := fmt.Fprintf(w, "%s\n\n", fn.String()); err != nil {
			return err
		}
	}
	return nil
}
