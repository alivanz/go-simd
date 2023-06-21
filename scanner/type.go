package scanner

import (
	"fmt"
	"strings"

	"github.com/iancoleman/strcase"
)

type Type struct {
	Name string
	Full string
}

func (t *Type) C() string {
	return t.Name
}

func (t *Type) Go() string {
	s := strings.TrimSuffix(string(t.Name), "_t")
	return strcase.ToCamel(s)
}

func (t *Type) Declare() string {
	if len(t.Full) > 0 {
		return fmt.Sprintf("\n// %s\ntype %s = C.%s\n", t.Full, t.Go(), t.C())
	}
	return fmt.Sprintf("\ntype %s = C.%s\n", t.Go(), t.C())
}
