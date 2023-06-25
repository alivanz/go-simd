package x86

/*
#cgo CFLAGS: -march=native
#include <immintrin.h>
*/
import "C"

// Convert packed single-precision (32-bit) floating-point elements in "a" to packed 32-bit integers, and store the results in "dst".
func MmCvtpsPi32(v0 M128) M64 {
	return C._mm_cvtps_pi32(v0)
}

// Convert packed single-precision (32-bit) floating-point elements in "a" to packed 32-bit integers, and store the results in "dst".
func MmCvtPs2Pi(v0 M128) M64 {
	return C._mm_cvt_ps2pi(v0)
}

// Convert packed single-precision (32-bit) floating-point elements in "a" to packed 32-bit integers with truncation, and store the results in "dst".
func MmCvttpsPi32(v0 M128) M64 {
	return C._mm_cvttps_pi32(v0)
}

// Convert packed single-precision (32-bit) floating-point elements in "a" to packed 32-bit integers with truncation, and store the results in "dst".
func MmCvttPs2Pi(v0 M128) M64 {
	return C._mm_cvtt_ps2pi(v0)
}

// Convert packed 32-bit integers in "b" to packed single-precision (32-bit) floating-point elements, store the results in the lower 2 elements of "dst", and copy the upper 2 packed elements from "a" to the upper elements of "dst".
func MmCvtpi32Ps(v0 M128, v1 M64) M128 {
	return C._mm_cvtpi32_ps(v0, v1)
}

// Convert packed signed 32-bit integers in "b" to packed single-precision (32-bit) floating-point elements, store the results in the lower 2 elements of "dst", and copy the upper 2 packed elements from "a" to the upper elements of "dst".
func MmCvtPi2Ps(v0 M128, v1 M64) M128 {
	return C._mm_cvt_pi2ps(v0, v1)
}

// Compare packed signed 16-bit integers in "a" and "b", and store packed maximum values in "dst".
func MmMaxPi16(v0 M64, v1 M64) M64 {
	return C._mm_max_pi16(v0, v1)
}

// Compare packed unsigned 8-bit integers in "a" and "b", and store packed maximum values in "dst".
func MmMaxPu8(v0 M64, v1 M64) M64 {
	return C._mm_max_pu8(v0, v1)
}

// Compare packed signed 16-bit integers in "a" and "b", and store packed minimum values in "dst".
func MmMinPi16(v0 M64, v1 M64) M64 {
	return C._mm_min_pi16(v0, v1)
}

// Compare packed unsigned 8-bit integers in "a" and "b", and store packed minimum values in "dst".
func MmMinPu8(v0 M64, v1 M64) M64 {
	return C._mm_min_pu8(v0, v1)
}

// Create mask from the most significant bit of each 8-bit element in "a", and store the result in "dst".
func MmMovemaskPi8(v0 M64) Int {
	return C._mm_movemask_pi8(v0)
}

// Multiply the packed unsigned 16-bit integers in "a" and "b", producing intermediate 32-bit integers, and store the high 16 bits of the intermediate integers in "dst".
func MmMulhiPu16(v0 M64, v1 M64) M64 {
	return C._mm_mulhi_pu16(v0, v1)
}

// Average packed unsigned 8-bit integers in "a" and "b", and store the results in "dst".
func MmAvgPu8(v0 M64, v1 M64) M64 {
	return C._mm_avg_pu8(v0, v1)
}

// Average packed unsigned 16-bit integers in "a" and "b", and store the results in "dst".
func MmAvgPu16(v0 M64, v1 M64) M64 {
	return C._mm_avg_pu16(v0, v1)
}

// Compute the absolute differences of packed unsigned 8-bit integers in "a" and "b", then horizontally sum each consecutive 8 differences to produce four unsigned 16-bit integers, and pack these unsigned 16-bit integers in the low 16 bits of "dst".
func MmSadPu8(v0 M64, v1 M64) M64 {
	return C._mm_sad_pu8(v0, v1)
}

// Convert packed 16-bit integers in "a" to packed single-precision (32-bit) floating-point elements, and store the results in "dst".
func MmCvtpi16Ps(v0 M64) M128 {
	return C._mm_cvtpi16_ps(v0)
}

// Convert packed unsigned 16-bit integers in "a" to packed single-precision (32-bit) floating-point elements, and store the results in "dst".
func MmCvtpu16Ps(v0 M64) M128 {
	return C._mm_cvtpu16_ps(v0)
}

// Convert the lower packed 8-bit integers in "a" to packed single-precision (32-bit) floating-point elements, and store the results in "dst".
func MmCvtpi8Ps(v0 M64) M128 {
	return C._mm_cvtpi8_ps(v0)
}

// Convert the lower packed unsigned 8-bit integers in "a" to packed single-precision (32-bit) floating-point elements, and store the results in "dst".
func MmCvtpu8Ps(v0 M64) M128 {
	return C._mm_cvtpu8_ps(v0)
}

// Convert packed signed 32-bit integers in "a" to packed single-precision (32-bit) floating-point elements, store the results in the lower 2 elements of "dst", then covert the packed signed 32-bit integers in "b" to single-precision (32-bit) floating-point element, and store the results in the upper 2 elements of "dst".
func MmCvtpi32X2Ps(v0 M64, v1 M64) M128 {
	return C._mm_cvtpi32x2_ps(v0, v1)
}

// Convert packed single-precision (32-bit) floating-point elements in "a" to packed 16-bit integers, and store the results in "dst". Note: this intrinsic will generate 0x7FFF, rather than 0x8000, for input values between 0x7FFF and 0x7FFFFFFF.
func MmCvtpsPi16(v0 M128) M64 {
	return C._mm_cvtps_pi16(v0)
}

// Convert packed single-precision (32-bit) floating-point elements in "a" to packed 8-bit integers, and store the results in lower 4 elements of "dst". Note: this intrinsic will generate 0x7F, rather than 0x80, for input values between 0x7F and 0x7FFFFFFF.
func MmCvtpsPi8(v0 M128) M64 {
	return C._mm_cvtps_pi8(v0)
}
