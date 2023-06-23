package writer

import (
	"io"
	"os"
	"path/filepath"
)

func WriteToFile(dst string, fn func(w io.Writer) error) error {
	if len(dst) == 0 {
		return nil
	}
	dst, err := filepath.Abs(dst)
	if err != nil {
		return err
	}
	if err := os.MkdirAll(filepath.Dir(dst), os.ModePerm); err != nil {
		return err
	}
	f, err := os.OpenFile(dst, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, os.ModePerm)
	if err != nil {
		return err
	}
	defer f.Close()
	return fn(f)
}
