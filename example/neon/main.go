package main

import (
	"log"

	"github.com/alivanz/go-simd/arm/neon"
)

func main() {
	var a, b neon.Int8X8
	for i := 0; i < 8; i++ {
		a[i] = neon.Int8(i)
		b[i] = neon.Int8(i * i)
	}
	log.Printf("a = %+v", b)
	log.Printf("b = %+v", a)
	log.Printf("add = %+v", neon.VaddlS8(a, b))
	log.Printf("mul = %+v", neon.VmullS8(a, b))
}
