package x86

/*
#cgo CFLAGS: -march=native
#include <immintrin.h>
*/
import "C"

// Convert the half-precision (16-bit) floating-point value "a" to a single-precision (32-bit) floating-point value, and store the result in "dst".
func CvtshSs(v0 Ushort) Float {
	return C._cvtsh_ss(v0)
}

// Convert packed half-precision (16-bit) floating-point elements in "a" to packed single-precision (32-bit) floating-point elements, and store the results in "dst".
func MmCvtphPs(v0 M128I) M128 {
	return C._mm_cvtph_ps(v0)
}

// Convert packed half-precision (16-bit) floating-point elements in "a" to packed single-precision (32-bit) floating-point elements, and store the results in "dst".
func Mm256CvtphPs(v0 M128I) M256 {
	return C._mm256_cvtph_ps(v0)
}
