package main

import (
	"log"

	"github.com/alivanz/go-simd/arm"
	"github.com/alivanz/go-simd/arm/neon"
)

func main() {
	var a, b arm.Int8X8
	var add, mul arm.Int16X8
	for i := 0; i < 8; i++ {
		a[i] = arm.Int8(i)
		b[i] = arm.Int8(i * i)
	}
	log.Printf("a = %+v", b)
	log.Printf("b = %+v", a)
	neon.VaddlS8(&add, &a, &b)
	neon.VmullS8(&mul, &a, &b)
	log.Printf("add = %+v", add)
	log.Printf("mul = %+v", mul)
}
