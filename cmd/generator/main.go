package main

import (
	"bytes"
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
			&cli.BoolFlag{
				Name:  "types",
				Usage: "skip types",
			},
			&cli.BoolFlag{
				Name:  "funcs",
				Usage: "skip functions",
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
		Name: "neon",
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
	}
	pkg.WriteTo(os.Stdout)
	return nil
}

func Source() ([]byte, error) {
	cmd := exec.Command("clang", "-E", "-")
	cmd.Stdin = bytes.NewBufferString("#include <arm_neon.h>")
	return cmd.Output()
}
