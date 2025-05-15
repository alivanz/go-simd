package neon

/*
#cgo CFLAGS: -march=armv8.5-a+crypto+i8mm
#include <arm_neon.h>
*/
import "C"

type int8x8 = C.int8x8_t

func vmulS8_cgo(r, v0, v1 *int8x8) {
	*r = C.vmul_s8(*v0, *v1)
}
