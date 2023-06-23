package x86

/*
#cgo CFLAGS: -march=native
#include <immintrin.h>
*/
import "C"

// _cvtsh_ss
// __always_inline__, __nodebug__, __target__("f16c"), __min_vector_width__(128)
func CvtshSs(v0 Ushort) Float {
	return C._cvtsh_ss(v0)
}

// _mm_cvtph_ps
// __always_inline__, __nodebug__, __target__("f16c"), __min_vector_width__(128)
func MmCvtphPs(v0 M128I) M128 {
	return C._mm_cvtph_ps(v0)
}

// _mm256_cvtph_ps
// __always_inline__, __nodebug__, __target__("f16c"), __min_vector_width__(256)
func Mm256CvtphPs(v0 M128I) M256 {
	return C._mm256_cvtph_ps(v0)
}
