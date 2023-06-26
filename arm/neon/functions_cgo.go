package neon

/*
#include <arm_neon.h>
*/
import "C"

func vmulS8_cgo(r *Int8X8, v0 *Int8X8, v1 *Int8X8) {
	*r = C.vmul_s8(*v0, *v1)
}
