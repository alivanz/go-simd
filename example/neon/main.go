package main

import (
	"log"

	"github.com/alivanz/go-simd/arm/neon"
)

func main() {
	var a, b neon.Int8X8
	var add, mul neon.Int16X8
	for i := 0; i < 8; i++ {
		a[i] = neon.Int8(i)
		b[i] = neon.Int8(i * i)
	}
	log.Printf("a = %+v", b)
	log.Printf("b = %+v", a)
	neon.VaddlS8(&add, &a, &b)
	neon.VmullS8(&mul, &a, &b)
	log.Printf("add = %+v", add)
	log.Printf("mul = %+v", mul)
}
