package main

import (
	"github.com/alivanz/go-simd/scanner"
)

func splitTarget(funcs []scanner.Function) map[string][]scanner.Function {
	m := make(map[string][]scanner.Function)
	for _, fn := range funcs {
		target := fn.Target()
		m[target] = append(m[target], fn)
	}
	return m
}
