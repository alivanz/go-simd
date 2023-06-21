package main

import (
	"bytes"
	"io"
	"log"
	"os"
	"os/exec"

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
		},
		Action: action,
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func action(cli *cli.Context) error {
	src, err := Source()
	if err != nil {
		return err
	}
	result, err := scanner.Scan(src)
	if err != nil {
		return err
	}
	pkg := &scanner.Package{
		Name: cli.String("package"),
		C: []string{
			"#include <arm_neon.h>",
		},
		Types:     result.Types,
		Functions: result.Functions,
	}
	if !cli.Bool("types") {
		pkg.Types = nil
	}
	if !cli.Bool("funcs") {
		pkg.Functions = nil
	} else {
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
	pkg.WriteTo(w)
	return nil
}

func Source() ([]byte, error) {
	cmd := exec.Command("clang", "-E", "-")
	cmd.Stdin = bytes.NewBufferString("#include <arm_neon.h>")
	return cmd.Output()
}
