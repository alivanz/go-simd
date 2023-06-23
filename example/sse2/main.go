package main

import (
	"log"

	"github.com/alivanz/go-simd/x86"
)

func main() {
	a := x86.MmSetrEpi8(0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15)
	b := x86.MmSetrEpi8(15, 14, 13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0)
	add := x86.MmAddEpi8(a, b)
	log.Print(a)
	log.Print(b)
	log.Print(add)
}
