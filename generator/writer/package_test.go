package writer

import (
	"bytes"
	"io"
	"strings"
	"testing"
)

func TestPackage(t *testing.T) {
	var buf bytes.Buffer
	Package(&buf, "abc")
	if buf.String() != "package abc\n" {
		t.Fatal(buf.String())
	}
}

func TestImport(t *testing.T) {
	var buf bytes.Buffer
	Import(&buf, []string{
		"pkg1",
		"pkg2",
		"pkg3",
	})
	if buf.String() != `
import (
	"pkg1"
	"pkg2"
	"pkg3"
)
` {
		t.Fatal(buf.String())
	}
}

func TestImportC(t *testing.T) {
	var buf bytes.Buffer
	ImportC(&buf, func(w io.Writer) error {
		io.WriteString(w, strings.Join([]string{
			`#include <abc.h>`,
			`#include <def.h>`,
		}, "\n"))
		return nil
	})
	ref := `
/*
#include <abc.h>
#include <def.h>
*/
import "C"
`
	if buf.String() != ref {
		t.Fatal(buf.String())
	}
}
