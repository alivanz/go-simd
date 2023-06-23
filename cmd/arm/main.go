package main

import (
	"bytes"
	"io"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/alivanz/go-simd/cmd/utils"
	"github.com/alivanz/go-simd/cmd/writer"
	"github.com/alivanz/go-simd/scanner"
)

func Source() ([]byte, error) {
	cmd := exec.Command("clang", "-E", "-")
	cmd.Stdin = bytes.NewBufferString(strings.Join(writer.Includes([]string{
		"arm_neon.h",
	}), "\n"))
	cmd.Stderr = os.Stderr
	return cmd.Output()
}

func main() {
	src, err := Source()
	if err != nil {
		log.Fatal(err)
	}
	// write raw
	if err := writer.WriteToFile("raw.h", func(w io.Writer) error {
		_, err := w.Write(src)
		return err
	}); err != nil {
		log.Fatal(err)
	}
	// scan
	result, err := scanner.Scan(src)
	if err != nil {
		log.Fatal(err)
	}
	// filter types
	mtype := make(map[string]bool)
	for _, fn := range result.Functions {
		if fn.Return != nil {
			mtype[fn.Return.Name] = true
		}
		for _, arg := range fn.Args {
			mtype[arg.Name] = true
		}
	}
	result.Types = utils.Filter(result.Types, func(t scanner.Type) bool {
		return mtype[t.Name]
	})
	// write types
	if err := writer.WriteToFile("types.go", func(w io.Writer) error {
		if err := writer.Package(w, "neon"); err != nil {
			return err
		}
		if err := writer.ImportC(w, strings.Join(writer.Includes([]string{
			"arm_neon.h",
		}), "\n")); err != nil {
			return err
		}
		if err := writer.Types(w, result.Types); err != nil {
			return err
		}
		return nil
	}); err != nil {
		log.Fatal(err)
	}
	// patch intrinsics info
	intrins, err := GetIntrinsics()
	if err != nil {
		log.Fatal(err)
	}
	for i, fn := range result.Functions {
		if info := intrins.Find(fn.Name); info != nil {
			result.Functions[i].Comment = info.Description
		}
	}
	// write functions
	if err := writer.WriteToFile("functions.go", func(w io.Writer) error {
		if err := writer.Package(w, "neon"); err != nil {
			return err
		}
		if err := writer.ImportC(w, strings.Join(writer.Includes([]string{
			"arm_neon.h",
		}), "\n")); err != nil {
			return err
		}
		return writer.Funcs(w, result.Functions, "")
	}); err != nil {
		log.Fatal(err)
	}
}
