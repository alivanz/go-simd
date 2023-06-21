package main

import (
	"bytes"
	"log"
	"os"
	"os/exec"

	"github.com/alivanz/go-simd/scanner"
)

func main() {
	src, err := Source()
	if err != nil {
		log.Fatal(err)
	}
	result, err := scanner.Scan(src)
	if err != nil {
		log.Fatal(err)
	}
	pkg := &scanner.Package{
		Name: "neon",
		C: []string{
			"#include <arm_neon.h>",
		},
		Types:     result.Types,
		Functions: result.Functions,
	}
	pkg.WriteTo(os.Stdout)
}

func Source() ([]byte, error) {
	cmd := exec.Command("clang", "-E", "-")
	cmd.Stdin = bytes.NewBufferString("#include <arm_neon.h>")
	return cmd.Output()
}
