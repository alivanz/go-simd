package neon

/*
#include <arm_neon.h>
void vmulS8_bypass(int8x8_t* r, int8x8_t* v0, int8x8_t* v1) { *r = vmul_s8(*v0, *v1); }
*/
import "C"

//go:linkname vmulS8_bypass vmulS8_bypass
//go:noescape
func vmulS8_bypass(r *Int8X8, v0 *Int8X8, v1 *Int8X8)
