package writer

import (
	"fmt"
	"strings"
)

func Cflags(flags []string) string {
	return fmt.Sprintf("#cgo CFLAGS: %s", strings.Join(flags, " "))
}

func Includes(headers []string) []string {
	out := make([]string, len(headers))
	for i, h := range headers {
		out[i] = fmt.Sprintf("#include <%s>", h)
	}
	return out
}
