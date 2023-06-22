package scanner

import (
	"fmt"
	"io"
	"strings"

	"github.com/iancoleman/strcase"
)

type Type struct {
	Name       string
	Full       string
	Attributes []string
}

func (t *Type) C() string {
	if !strings.Contains(t.Name, " ") {
		return t.Name
	}
	s := strings.Replace(t.Name, "unsigned", "u", -1)
	s = strings.Replace(s, " ", "", -1)
	return s
}

func (t *Type) Go() string {
	s := strings.TrimSuffix(string(t.Name), "_t")
	return strcase.ToCamel(s)
}

func (t *Type) Declare(w io.Writer) error {
	var err error
	if len(t.Full) > 0 {
		_, err = fmt.Fprintf(w, "\n// %s\ntype %s = C.%s\n", t.Full, t.Go(), t.C())
	} else {
		_, err = fmt.Fprintf(w, "\ntype %s = C.%s\n", t.Go(), t.C())
	}
	return err
}
