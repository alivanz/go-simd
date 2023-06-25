package x86

/*
#cgo CFLAGS: -march=native
#include <immintrin.h>
*/
import "C"

// Compute the absolute value of packed signed 8-bit integers in "a", and store the unsigned results in "dst".
func MmAbsPi8(v0 M64) M64 {
	return C._mm_abs_pi8(v0)
}

// Compute the absolute value of packed signed 16-bit integers in "a", and store the unsigned results in "dst".
func MmAbsPi16(v0 M64) M64 {
	return C._mm_abs_pi16(v0)
}

// Compute the absolute value of packed signed 32-bit integers in "a", and store the unsigned results in "dst".
func MmAbsPi32(v0 M64) M64 {
	return C._mm_abs_pi32(v0)
}

// Horizontally add adjacent pairs of 16-bit integers in "a" and "b", and pack the signed 16-bit results in "dst".
func MmHaddPi16(v0 M64, v1 M64) M64 {
	return C._mm_hadd_pi16(v0, v1)
}

// Horizontally add adjacent pairs of 32-bit integers in "a" and "b", and pack the signed 32-bit results in "dst".
func MmHaddPi32(v0 M64, v1 M64) M64 {
	return C._mm_hadd_pi32(v0, v1)
}

// Horizontally add adjacent pairs of signed 16-bit integers in "a" and "b" using saturation, and pack the signed 16-bit results in "dst".
func MmHaddsPi16(v0 M64, v1 M64) M64 {
	return C._mm_hadds_pi16(v0, v1)
}

// Horizontally subtract adjacent pairs of 16-bit integers in "a" and "b", and pack the signed 16-bit results in "dst".
func MmHsubPi16(v0 M64, v1 M64) M64 {
	return C._mm_hsub_pi16(v0, v1)
}

// Horizontally subtract adjacent pairs of 32-bit integers in "a" and "b", and pack the signed 32-bit results in "dst".
func MmHsubPi32(v0 M64, v1 M64) M64 {
	return C._mm_hsub_pi32(v0, v1)
}

// Horizontally subtract adjacent pairs of signed 16-bit integers in "a" and "b" using saturation, and pack the signed 16-bit results in "dst".
func MmHsubsPi16(v0 M64, v1 M64) M64 {
	return C._mm_hsubs_pi16(v0, v1)
}

// Vertically multiply each unsigned 8-bit integer from "a" with the corresponding signed 8-bit integer from "b", producing intermediate signed 16-bit integers. Horizontally add adjacent pairs of intermediate signed 16-bit integers, and pack the saturated results in "dst".
func MmMaddubsPi16(v0 M64, v1 M64) M64 {
	return C._mm_maddubs_pi16(v0, v1)
}

// Multiply packed signed 16-bit integers in "a" and "b", producing intermediate signed 32-bit integers. Truncate each intermediate integer to the 18 most significant bits, round by adding 1, and store bits [16:1] to "dst".
func MmMulhrsPi16(v0 M64, v1 M64) M64 {
	return C._mm_mulhrs_pi16(v0, v1)
}

// Shuffle packed 8-bit integers in "a" according to shuffle control mask in the corresponding 8-bit element of "b", and store the results in "dst".
func MmShufflePi8(v0 M64, v1 M64) M64 {
	return C._mm_shuffle_pi8(v0, v1)
}

// Negate packed 8-bit integers in "a" when the corresponding signed 8-bit integer in "b" is negative, and store the results in "dst". Element in "dst" are zeroed out when the corresponding element in "b" is zero.
func MmSignPi8(v0 M64, v1 M64) M64 {
	return C._mm_sign_pi8(v0, v1)
}

// Negate packed 16-bit integers in "a" when the corresponding signed 16-bit integer in "b" is negative, and store the results in "dst". Element in "dst" are zeroed out when the corresponding element in "b" is zero.
func MmSignPi16(v0 M64, v1 M64) M64 {
	return C._mm_sign_pi16(v0, v1)
}

// Negate packed 32-bit integers in "a" when the corresponding signed 32-bit integer in "b" is negative, and store the results in "dst". Element in "dst" are zeroed out when the corresponding element in "b" is zero.
func MmSignPi32(v0 M64, v1 M64) M64 {
	return C._mm_sign_pi32(v0, v1)
}
