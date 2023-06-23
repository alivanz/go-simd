package x86

/*
#cgo CFLAGS: -march=native
#include <immintrin.h>
*/
import "C"

// _mm_aesenc_si128
// __always_inline__, __nodebug__, __target__("aes"), __min_vector_width__(128)
func MmAesencSi128(v0 M128I, v1 M128I) M128I {
	return C._mm_aesenc_si128(v0, v1)
}

// _mm_aesenclast_si128
// __always_inline__, __nodebug__, __target__("aes"), __min_vector_width__(128)
func MmAesenclastSi128(v0 M128I, v1 M128I) M128I {
	return C._mm_aesenclast_si128(v0, v1)
}

// _mm_aesdec_si128
// __always_inline__, __nodebug__, __target__("aes"), __min_vector_width__(128)
func MmAesdecSi128(v0 M128I, v1 M128I) M128I {
	return C._mm_aesdec_si128(v0, v1)
}

// _mm_aesdeclast_si128
// __always_inline__, __nodebug__, __target__("aes"), __min_vector_width__(128)
func MmAesdeclastSi128(v0 M128I, v1 M128I) M128I {
	return C._mm_aesdeclast_si128(v0, v1)
}

// _mm_aesimc_si128
// __always_inline__, __nodebug__, __target__("aes"), __min_vector_width__(128)
func MmAesimcSi128(v0 M128I) M128I {
	return C._mm_aesimc_si128(v0)
}
