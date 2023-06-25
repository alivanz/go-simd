package x86

/*
#cgo CFLAGS: -march=native
#include <immintrin.h>
*/
import "C"

// Convert packed double-precision (64-bit) floating-point elements in "a" to packed 32-bit integers, and store the results in "dst".
func MmCvtpdPi32(v0 M128D) M64 {
	return C._mm_cvtpd_pi32(v0)
}

// Convert packed double-precision (64-bit) floating-point elements in "a" to packed 32-bit integers with truncation, and store the results in "dst".
func MmCvttpdPi32(v0 M128D) M64 {
	return C._mm_cvttpd_pi32(v0)
}

// Convert packed signed 32-bit integers in "a" to packed double-precision (64-bit) floating-point elements, and store the results in "dst".
func MmCvtpi32Pd(v0 M64) M128D {
	return C._mm_cvtpi32_pd(v0)
}

// Add 64-bit integers "a" and "b", and store the result in "dst".
func MmAddSi64(v0 M64, v1 M64) M64 {
	return C._mm_add_si64(v0, v1)
}

// Multiply the low unsigned 32-bit integers from "a" and "b", and store the unsigned 64-bit result in "dst".
func MmMulSu32(v0 M64, v1 M64) M64 {
	return C._mm_mul_su32(v0, v1)
}

// Subtract 64-bit integer "b" from 64-bit integer "a", and store the result in "dst".
func MmSubSi64(v0 M64, v1 M64) M64 {
	return C._mm_sub_si64(v0, v1)
}
