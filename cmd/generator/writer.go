package main

import (
	"io"
	"os"
)

func writeToFile(dst string, fn func(w io.Writer) error) error {
	if len(dst) == 0 {
		return nil
	}
	f, err := os.OpenFile(dst, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, os.ModePerm)
	if err != nil {
		return err
	}
	defer f.Close()
	return fn(f)
}
