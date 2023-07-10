package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strings"

	"github.com/alivanz/go-simd/generator/scanner"
	"github.com/alivanz/go-simd/generator/types"
	"github.com/alivanz/go-simd/generator/utils"
	"github.com/alivanz/go-simd/generator/writer"
)

var (
	regComma = regexp.MustCompile(`\s*,\s*`)
)

func main() {
	// generate
	cmd := exec.Command("clang", "-march=native", "-E", "-")
	cmd.Stdin = bytes.NewBufferString(strings.Join([]string{
		"#include <immintrin.h>",
	}, "\n"))
	cmd.Stderr = os.Stderr
	src, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	// raw
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
	// filter functions
	mfunc := make(map[string]bool)
	result.Functions = utils.Filter(result.Functions, func(fn types.Function) bool {
		if mfunc[fn.Name] {
			return false
		}
		if len(fn.Target()) == 0 {
			return false
		}
		mfunc[fn.Name] = true
		return true
	})
	// filter types
	mtype := make(map[string]bool)
	for _, fn := range result.Functions {
		if fn.Return != nil {
			mtype[fn.Return.Name] = true
			// append type
			result.Types = append(result.Types, *fn.Return)
		}
		for _, arg := range fn.Args {
			mtype[arg.Name] = true
			result.Types = append(result.Types, arg)
		}
	}
	result.Types = utils.Filter(result.Types, func(t types.Type) bool {
		if !mtype[t.Name] {
			return false
		}
		// remove dup
		delete(mtype, t.Name)
		return true
	})
	// types
	if err := writer.WriteToFile("types.go", func(w io.Writer) error {
		if err := writer.Package(w, "x86"); err != nil {
			return err
		}
		if err := writer.ImportC(w, func(w io.Writer) error {
			fmt.Fprintf(w, "#include <immintrin.h>")
			return err
		}); err != nil {
			return err
		}
		if err := writer.Types(w, result.Types); err != nil {
			return err
		}
		return nil
	}); err != nil {
		log.Fatal(err)
	}
	// patch funcs
	intrins, err := GetIntrinsic()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%+v", intrins[0])
	mintrin := make(map[string]*Intrinsic)
	for _, intrin := range intrins {
		mintrin[intrin.Name] = intrin
	}
	log.Printf("%+v", mintrin["_mm_fmsubadd_pd"])
	for i, fn := range result.Functions {
		if intrin, found := mintrin[fn.Name]; found {
			result.Functions[i].Comment = intrin.Description
		}
	}
	// group funcs by target
	mf := make(map[string][]types.Function)
	for _, fn := range result.Functions {
		target := fn.Target()
		mf[target] = append(mf[target], fn)
	}
	// funcs
	for target, funcs := range mf {
		target = regComma.ReplaceAllString(target, "_")
		cname := fmt.Sprintf("%s/functions.c", target)
		fname := fmt.Sprintf("%s/functions.go", target)
		// write C
		if err := writer.WriteToFile(cname, func(w io.Writer) error {
			if _, err := io.WriteString(w, "#include <immintrin.h>\n\n"); err != nil {
				return err
			}
			for _, fn := range funcs {
				if fn.Blacklisted() {
					continue
				}
				if err := writer.RewriteC(w, fn); err != nil {
					return err
				}
			}
			return nil
		}); err != nil {
			log.Fatal(err)
		}
		// write Go
		if err := writer.WriteToFile(fname, func(w io.Writer) error {
			if err := writer.Package(w, target); err != nil {
				return err
			}
			if err := writer.Import(w, []string{
				"github.com/alivanz/go-simd/x86",
			}); err != nil {
				return err
			}
			if err := writer.ImportC(w, func(w io.Writer) error {
				feats := strings.Split(target, "_")
				if len(feats) > 0 {
					fmt.Fprintf(w, "#cgo CFLAGS: %s\n", strings.Join(utils.Transform(feats, func(i int, feat string) string {
						return "-m" + feat
					}), " "))
				}
				fmt.Fprintf(w, "#include <immintrin.h>")
				return err
			}); err != nil {
				return err
			}
			for _, fn := range funcs {
				if fn.Blacklisted() {
					continue
				}
				if err := writer.DeclareFuncBypass(w, fn, "x86"); err != nil {
					return err
				}
			}
			return nil
		}); err != nil {
			log.Fatal(err)
		}
	}
}
