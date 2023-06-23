package main

import (
	"bytes"
	"io"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strings"

	"github.com/alivanz/go-simd/cmd/utils"
	"github.com/alivanz/go-simd/cmd/writer"
	"github.com/alivanz/go-simd/scanner"
	"github.com/urfave/cli/v2"
)

var (
	regComma = regexp.MustCompile(`\s*,\s*`)
)

func main() {
	app := cli.App{
		Name: "go-simd-generator",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "package",
				Usage:    "Go package name",
				Required: true,
			},
			&cli.StringFlag{
				Name:  "types",
				Usage: "types output file",
			},
			&cli.StringFlag{
				Name:  "funcs",
				Usage: "functions output file",
			},
			&cli.BoolFlag{
				Name:  "split-target",
				Usage: "split by target",
			},
			&cli.StringFlag{
				Name:  "raw",
				Usage: "raw unprocessed output file",
			},
			&cli.StringSliceFlag{
				Name:    "header",
				Aliases: []string{"head"},
				Usage:   "C header file",
			},
		},
		Action: action,
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func action(cli *cli.Context) error {
	var (
		flags   []string
		headers []string
	)
	flags = cli.Args().Slice()
	headers = cli.StringSlice("header")
	src, err := Source(flags, headers)
	if err != nil {
		return err
	}
	// raw
	if err := writer.WriteToFile(cli.String("raw"), func(w io.Writer) error {
		_, err := w.Write(src)
		return err
	}); err != nil {
		return err
	}
	// scan
	result, err := scanner.Scan(src)
	if err != nil {
		return err
	}
	pkg := cli.String("package")
	// filter functions
	mfunc := make(map[string]bool)
	result.Functions = utils.Filter(result.Functions, func(fn scanner.Function) bool {
		if mfunc[fn.Name] {
			return false
		}
		if cli.Bool("split-target") && len(fn.Target()) == 0 {
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
	result.Types = utils.Filter(result.Types, func(t scanner.Type) bool {
		if !mtype[t.Name] {
			return false
		}
		// remove dup
		delete(mtype, t.Name)
		return true
	})
	// types
	if err := writer.WriteToFile(cli.String("types"), func(w io.Writer) error {
		if err := writer.Package(w, pkg); err != nil {
			return err
		}
		if err := writer.ImportC(w, strings.Join(append(
			[]string{writer.Cflags(flags)},
			writer.Includes(headers)...,
		), "\n")); err != nil {
			return err
		}
		if err := writer.Types(w, result.Types); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return err
	}
	// funcs
	if !cli.Bool("split-target") {
		if err := writer.WriteToFile(cli.String("funcs"), wrapFuncs(pkg, "", flags, headers, result.Functions)); err != nil {
			return err
		}
	} else {
		mf := splitTarget(result.Functions)
		fname := cli.String("funcs")
		for target, funcs := range mf {
			target = regComma.ReplaceAllString(target, "_")
			fname := strings.ReplaceAll(fname, "*", target)
			if err := writer.WriteToFile(strings.ReplaceAll(fname, "*", target), wrapFuncs(pkg, "", flags, headers, funcs)); err != nil {
				return err
			}
		}
	}
	return nil
}

func Source(flags, headers []string) ([]byte, error) {
	args := append(flags, "-E", "-")
	log.Printf("args: %s", strings.Join(args, " "))
	cmd := exec.Command("clang", args...)
	cmd.Stdin = bytes.NewBufferString(strings.Join(writer.Includes(headers), "\n"))
	return cmd.Output()
}

func wrapFuncs(pkg, typePkg string, flags, headers []string, funcs []scanner.Function) func(w io.Writer) error {
	return func(w io.Writer) error {
		if err := writer.Package(w, pkg); err != nil {
			return err
		}
		if err := writer.ImportC(w, strings.Join(append(
			[]string{writer.Cflags(flags)},
			writer.Includes(headers)...,
		), "\n")); err != nil {
			return err
		}
		// if err := Import(w, []string{
		// 	"github.com/alivanz/go-simd/x86",
		// }); err != nil {
		// 	return err
		// }
		switch pkg {
		case "neon":
			intrins, err := GetIntrinsics()
			if err != nil {
				return err
			}
			for i, fn := range funcs {
				if info := intrins.Find(fn.Name); info != nil {
					funcs[i].Comment = info.Description
				}
			}
		}
		return writer.Funcs(w, funcs, typePkg)
	}
}
