package types

import (
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

func (t *Type) Go(pkg string) string {
	s := strings.TrimSuffix(string(t.Name), "_t")
	s = strcase.ToCamel(s)
	if len(pkg) > 0 {
		return pkg + "." + s
	}
	return s
}