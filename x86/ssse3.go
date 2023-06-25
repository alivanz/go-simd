package x86

/*
#cgo CFLAGS: -march=native
#include <immintrin.h>
*/
import "C"

// Compute the absolute value of packed signed 8-bit integers in "a", and store the unsigned results in "dst".
func MmAbsEpi8(v0 M128I) M128I {
	return C._mm_abs_epi8(v0)
}

// Compute the absolute value of packed signed 16-bit integers in "a", and store the unsigned results in "dst".
func MmAbsEpi16(v0 M128I) M128I {
	return C._mm_abs_epi16(v0)
}

// Compute the absolute value of packed signed 32-bit integers in "a", and store the unsigned results in "dst".
func MmAbsEpi32(v0 M128I) M128I {
	return C._mm_abs_epi32(v0)
}

// Horizontally add adjacent pairs of 16-bit integers in "a" and "b", and pack the signed 16-bit results in "dst".
func MmHaddEpi16(v0 M128I, v1 M128I) M128I {
	return C._mm_hadd_epi16(v0, v1)
}

// Horizontally add adjacent pairs of 32-bit integers in "a" and "b", and pack the signed 32-bit results in "dst".
func MmHaddEpi32(v0 M128I, v1 M128I) M128I {
	return C._mm_hadd_epi32(v0, v1)
}

// Horizontally add adjacent pairs of signed 16-bit integers in "a" and "b" using saturation, and pack the signed 16-bit results in "dst".
func MmHaddsEpi16(v0 M128I, v1 M128I) M128I {
	return C._mm_hadds_epi16(v0, v1)
}

// Horizontally subtract adjacent pairs of 16-bit integers in "a" and "b", and pack the signed 16-bit results in "dst".
func MmHsubEpi16(v0 M128I, v1 M128I) M128I {
	return C._mm_hsub_epi16(v0, v1)
}

// Horizontally subtract adjacent pairs of 32-bit integers in "a" and "b", and pack the signed 32-bit results in "dst".
func MmHsubEpi32(v0 M128I, v1 M128I) M128I {
	return C._mm_hsub_epi32(v0, v1)
}

// Horizontally subtract adjacent pairs of signed 16-bit integers in "a" and "b" using saturation, and pack the signed 16-bit results in "dst".
func MmHsubsEpi16(v0 M128I, v1 M128I) M128I {
	return C._mm_hsubs_epi16(v0, v1)
}

// Vertically multiply each unsigned 8-bit integer from "a" with the corresponding signed 8-bit integer from "b", producing intermediate signed 16-bit integers. Horizontally add adjacent pairs of intermediate signed 16-bit integers, and pack the saturated results in "dst".
func MmMaddubsEpi16(v0 M128I, v1 M128I) M128I {
	return C._mm_maddubs_epi16(v0, v1)
}

// Multiply packed signed 16-bit integers in "a" and "b", producing intermediate signed 32-bit integers. Truncate each intermediate integer to the 18 most significant bits, round by adding 1, and store bits [16:1] to "dst".
func MmMulhrsEpi16(v0 M128I, v1 M128I) M128I {
	return C._mm_mulhrs_epi16(v0, v1)
}

// Shuffle packed 8-bit integers in "a" according to shuffle control mask in the corresponding 8-bit element of "b", and store the results in "dst".
func MmShuffleEpi8(v0 M128I, v1 M128I) M128I {
	return C._mm_shuffle_epi8(v0, v1)
}

// Negate packed 8-bit integers in "a" when the corresponding signed 8-bit integer in "b" is negative, and store the results in "dst". Element in "dst" are zeroed out when the corresponding element in "b" is zero.
func MmSignEpi8(v0 M128I, v1 M128I) M128I {
	return C._mm_sign_epi8(v0, v1)
}

// Negate packed 16-bit integers in "a" when the corresponding signed 16-bit integer in "b" is negative, and store the results in "dst". Element in "dst" are zeroed out when the corresponding element in "b" is zero.
func MmSignEpi16(v0 M128I, v1 M128I) M128I {
	return C._mm_sign_epi16(v0, v1)
}

// Negate packed 32-bit integers in "a" when the corresponding signed 32-bit integer in "b" is negative, and store the results in "dst". Element in "dst" are zeroed out when the corresponding element in "b" is zero.
func MmSignEpi32(v0 M128I, v1 M128I) M128I {
	return C._mm_sign_epi32(v0, v1)
}
