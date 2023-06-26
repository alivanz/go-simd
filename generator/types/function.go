package types

import (
	"regexp"
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
