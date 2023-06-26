package types

import (
	"regexp"
	"strings"
)

type Function struct {
	Name      string
	Args      []Type
	Return    *Type
	Attribute string
	Comment   string
}

type Arg struct {
	Name string
	Type string
}

var (
	regTarget = regexp.MustCompile(`__target__\("([a-z0-9\s,]+)"\)`)
)

func (f *Function) Target() string {
	match := regTarget.FindStringSubmatch(f.Attribute)
	if match == nil {
		return ""
	}
	return match[1]
}

func (fn *Function) Blacklisted() bool {
	for _, blacklist := range []string{
		"f16",
		"vcmla",
		"__extension__",
	} {
		if strings.Contains(fn.Name, blacklist) {
			return true
		}
	}
	return false
}
