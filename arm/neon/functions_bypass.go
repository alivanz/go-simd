package neon

/*
#include <arm_neon.h>
void vmulS8_bypass(int8x8_t* r, int8x8_t* v0, int8x8_t* v1) { *r = vmul_s8(*v0, *v1); }
void vmulS8_full(int8_t* r, int8_t* v0, int8_t* v1, int n) {
	int8x8_t* pr = (int8x8_t*)r;
	int8x8_t* pa = (int8x8_t*)v0;
	int8x8_t* pb = (int8x8_t*)v1;
	for (int i=0; i<n; i+=8) {
		*pr = vmul_s8(*pa, *pb);
		pr += 1;
		pa += 1;
		pb += 1;
	}
}
*/
import "C"

//go:linkname vmulS8_bypass vmulS8_bypass
//go:noescape
func vmulS8_bypass(r *Int8X8, v0 *Int8X8, v1 *Int8X8)

//go:linkname vmulS8_full vmulS8_full
//go:noescape
func vmulS8_full(r *int8, v0 *int8, v1 *int8, n int)
