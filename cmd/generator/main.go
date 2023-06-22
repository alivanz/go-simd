package main

import (
	"bytes"
	"fmt"
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
			&cli.StringFlag{
				Name:  "types",
				Usage: "types output file",
			},
			&cli.StringFlag{
				Name:  "funcs",
				Usage: "functions output file",
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
	if fname := cli.String("raw"); len(fname) > 0 {
		f, err := os.OpenFile(fname, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, os.ModePerm)
		if err != nil {
			return err
		}
		defer f.Close()
		f.Write(src)
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
	}
	if fname := cli.String("types"); len(fname) > 0 {
		f, err := os.OpenFile(fname, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, os.ModePerm)
		if err != nil {
			return err
		}
		defer f.Close()
		pkg.Types = result.Types
		pkg.WriteTo(f)
		pkg.Types = nil
	}
	if fname := cli.String("funcs"); len(fname) > 0 {
		f, err := os.OpenFile(fname, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, os.ModePerm)
		if err != nil {
			return err
		}
		defer f.Close()
		pkg.Functions = result.Functions
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
		pkg.WriteTo(f)
		pkg.Functions = nil
	}
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
