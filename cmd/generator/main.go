package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/alivanz/go-simd/scanner"
	"github.com/urfave/cli/v2"
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
			&cli.BoolFlag{
				Name:  "types",
				Usage: "skip types",
			},
			&cli.BoolFlag{
				Name:  "funcs",
				Usage: "skip functions",
			},
			&cli.StringFlag{
				Name:    "output",
				Aliases: []string{"o"},
				Usage:   "output file",
			},
			&cli.BoolFlag{
				Name:  "raw",
				Usage: "raw unprocessed",
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
	var w io.Writer
	if output := cli.String("output"); len(output) > 0 {
		f, err := os.OpenFile(output, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, os.ModePerm)
		if err != nil {
			return err
		}
		defer f.Close()
		w = f
	} else {
		w = os.Stdout
	}
	if cli.Bool("raw") {
		w.Write(src)
		return nil
	}
	result, err := scanner.Scan(src)
	if err != nil {
		return err
	}
	pkg := &scanner.Package{
		Name: cli.String("package"),
		C: append(
			[]string{cflags(flags)},
			includes(headers)...,
		),
		Types:     result.Types,
		Functions: result.Functions,
	}
	if !cli.Bool("types") {
		pkg.Types = nil
	}
	if !cli.Bool("funcs") {
		pkg.Functions = nil
	} else {
		switch cli.String("package") {
		case "neon":
			intrins, err := GetIntrinsics()
			if err != nil {
				return err
			}
			for i, fn := range pkg.Functions {
				if info := intrins.Find(fn.Name); info != nil {
					pkg.Functions[i].Comment = info.Description
				}
			}
		}
	}
	pkg.WriteTo(w)
	return nil
}

func Source(flags, headers []string) ([]byte, error) {
	args := append(flags, "-E", "-")
	log.Printf("args: %s", strings.Join(args, " "))
	cmd := exec.Command("clang", args...)
	cmd.Stdin = bytes.NewBufferString(strings.Join(includes(headers), "\n"))
	return cmd.Output()
}

func cflags(flags []string) string {
	return fmt.Sprintf("#cgo CFLAGS: %s", strings.Join(flags, " "))
}

func includes(headers []string) []string {
	out := make([]string, len(headers))
	for i, h := range headers {
		out[i] = fmt.Sprintf("#include <%s>", h)
	}
	return out
}
