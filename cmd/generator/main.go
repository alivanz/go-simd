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
	// raw
	if err := writeToFile(cli.String("raw"), func(w io.Writer) error {
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
	// types
	if err := writeToFile(cli.String("types"), func(w io.Writer) error {
		if err := Package(w, pkg); err != nil {
			return err
		}
		if err := ImportC(w, strings.Join(append(
			[]string{cflags(flags)},
			includes(headers)...,
		), "\n")); err != nil {
			return err
		}
		if err := Types(w, result.Types); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return err
	}
	// funcs
	if err := writeToFile(cli.String("funcs"), func(w io.Writer) error {
		if err := Package(w, pkg); err != nil {
			return err
		}
		if err := ImportC(w, strings.Join(append(
			[]string{cflags(flags)},
			includes(headers)...,
		), "\n")); err != nil {
			return err
		}
		switch cli.String("package") {
		case "neon":
			intrins, err := GetIntrinsics()
			if err != nil {
				return err
			}
			for i, fn := range result.Functions {
				if info := intrins.Find(fn.Name); info != nil {
					result.Functions[i].Comment = info.Description
				}
			}
		}
		return Funcs(w, result.Functions)
	}); err != nil {
		return err
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
