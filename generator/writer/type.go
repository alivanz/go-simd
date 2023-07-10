package writer

import (
	"fmt"
	"io"

	"github.com/alivanz/go-simd/generator/types"
)

func DeclareType(w io.Writer, t types.Type) error {
	var err error
	if len(t.Full) > 0 {
		_, err = fmt.Fprintf(w, "\n// %s\ntype %s = C.%s\n", t.Full, t.Go(""), t.CGO())
	} else {
		_, err = fmt.Fprintf(w, "\ntype %s = C.%s\n", t.Go(""), t.CGO())
	}
	return err
}
