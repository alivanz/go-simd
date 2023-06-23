package x86

/*
#cgo CFLAGS: -march=native
#include <immintrin.h>
*/
import "C"

// _mm_cvtpd_pi32
// __always_inline__, __nodebug__, __target__("mmx,sse2"), __min_vector_width__(64)
func MmCvtpdPi32(v0 M128D) M64 {
	return C._mm_cvtpd_pi32(v0)
}

// _mm_cvttpd_pi32
// __always_inline__, __nodebug__, __target__("mmx,sse2"), __min_vector_width__(64)
func MmCvttpdPi32(v0 M128D) M64 {
	return C._mm_cvttpd_pi32(v0)
}

// _mm_cvtpi32_pd
// __always_inline__, __nodebug__, __target__("mmx,sse2"), __min_vector_width__(64)
func MmCvtpi32Pd(v0 M64) M128D {
	return C._mm_cvtpi32_pd(v0)
}

// _mm_add_si64
// __always_inline__, __nodebug__, __target__("mmx,sse2"), __min_vector_width__(64)
func MmAddSi64(v0 M64, v1 M64) M64 {
	return C._mm_add_si64(v0, v1)
}

// _mm_mul_su32
// __always_inline__, __nodebug__, __target__("mmx,sse2"), __min_vector_width__(64)
func MmMulSu32(v0 M64, v1 M64) M64 {
	return C._mm_mul_su32(v0, v1)
}

// _mm_sub_si64
// __always_inline__, __nodebug__, __target__("mmx,sse2"), __min_vector_width__(64)
func MmSubSi64(v0 M64, v1 M64) M64 {
	return C._mm_sub_si64(v0, v1)
}
