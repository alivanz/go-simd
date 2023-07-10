package mmx_sse

import (
	"github.com/alivanz/go-simd/x86"
)

/*
#cgo CFLAGS: -mmmx -msse
#include <immintrin.h>
*/
import "C"

// Convert packed single-precision (32-bit) floating-point elements in "a" to packed 32-bit integers, and store the results in "dst".
//
//go:linkname MmCvtpsPi32 MmCvtpsPi32
//go:noescape
func MmCvtpsPi32(r *x86.M64, v0 *x86.M128)

// Convert packed single-precision (32-bit) floating-point elements in "a" to packed 32-bit integers, and store the results in "dst".
//
//go:linkname MmCvtPs2Pi MmCvtPs2Pi
//go:noescape
func MmCvtPs2Pi(r *x86.M64, v0 *x86.M128)

// Convert packed single-precision (32-bit) floating-point elements in "a" to packed 32-bit integers with truncation, and store the results in "dst".
//
//go:linkname MmCvttpsPi32 MmCvttpsPi32
//go:noescape
func MmCvttpsPi32(r *x86.M64, v0 *x86.M128)

// Convert packed single-precision (32-bit) floating-point elements in "a" to packed 32-bit integers with truncation, and store the results in "dst".
//
//go:linkname MmCvttPs2Pi MmCvttPs2Pi
//go:noescape
func MmCvttPs2Pi(r *x86.M64, v0 *x86.M128)

// Convert packed 32-bit integers in "b" to packed single-precision (32-bit) floating-point elements, store the results in the lower 2 elements of "dst", and copy the upper 2 packed elements from "a" to the upper elements of "dst".
//
//go:linkname MmCvtpi32Ps MmCvtpi32Ps
//go:noescape
func MmCvtpi32Ps(r *x86.M128, v0 *x86.M128, v1 *x86.M64)

// Convert packed signed 32-bit integers in "b" to packed single-precision (32-bit) floating-point elements, store the results in the lower 2 elements of "dst", and copy the upper 2 packed elements from "a" to the upper elements of "dst".
//
//go:linkname MmCvtPi2Ps MmCvtPi2Ps
//go:noescape
func MmCvtPi2Ps(r *x86.M128, v0 *x86.M128, v1 *x86.M64)

// Compare packed signed 16-bit integers in "a" and "b", and store packed maximum values in "dst".
//
//go:linkname MmMaxPi16 MmMaxPi16
//go:noescape
func MmMaxPi16(r *x86.M64, v0 *x86.M64, v1 *x86.M64)

// Compare packed unsigned 8-bit integers in "a" and "b", and store packed maximum values in "dst".
//
//go:linkname MmMaxPu8 MmMaxPu8
//go:noescape
func MmMaxPu8(r *x86.M64, v0 *x86.M64, v1 *x86.M64)

// Compare packed signed 16-bit integers in "a" and "b", and store packed minimum values in "dst".
//
//go:linkname MmMinPi16 MmMinPi16
//go:noescape
func MmMinPi16(r *x86.M64, v0 *x86.M64, v1 *x86.M64)

// Compare packed unsigned 8-bit integers in "a" and "b", and store packed minimum values in "dst".
//
//go:linkname MmMinPu8 MmMinPu8
//go:noescape
func MmMinPu8(r *x86.M64, v0 *x86.M64, v1 *x86.M64)

// Create mask from the most significant bit of each 8-bit element in "a", and store the result in "dst".
//
//go:linkname MmMovemaskPi8 MmMovemaskPi8
//go:noescape
func MmMovemaskPi8(r *x86.Int, v0 *x86.M64)

// Multiply the packed unsigned 16-bit integers in "a" and "b", producing intermediate 32-bit integers, and store the high 16 bits of the intermediate integers in "dst".
//
//go:linkname MmMulhiPu16 MmMulhiPu16
//go:noescape
func MmMulhiPu16(r *x86.M64, v0 *x86.M64, v1 *x86.M64)

// Average packed unsigned 8-bit integers in "a" and "b", and store the results in "dst".
//
//go:linkname MmAvgPu8 MmAvgPu8
//go:noescape
func MmAvgPu8(r *x86.M64, v0 *x86.M64, v1 *x86.M64)

// Average packed unsigned 16-bit integers in "a" and "b", and store the results in "dst".
//
//go:linkname MmAvgPu16 MmAvgPu16
//go:noescape
func MmAvgPu16(r *x86.M64, v0 *x86.M64, v1 *x86.M64)

// Compute the absolute differences of packed unsigned 8-bit integers in "a" and "b", then horizontally sum each consecutive 8 differences to produce four unsigned 16-bit integers, and pack these unsigned 16-bit integers in the low 16 bits of "dst".
//
//go:linkname MmSadPu8 MmSadPu8
//go:noescape
func MmSadPu8(r *x86.M64, v0 *x86.M64, v1 *x86.M64)

// Convert packed 16-bit integers in "a" to packed single-precision (32-bit) floating-point elements, and store the results in "dst".
//
//go:linkname MmCvtpi16Ps MmCvtpi16Ps
//go:noescape
func MmCvtpi16Ps(r *x86.M128, v0 *x86.M64)

// Convert packed unsigned 16-bit integers in "a" to packed single-precision (32-bit) floating-point elements, and store the results in "dst".
//
//go:linkname MmCvtpu16Ps MmCvtpu16Ps
//go:noescape
func MmCvtpu16Ps(r *x86.M128, v0 *x86.M64)

// Convert the lower packed 8-bit integers in "a" to packed single-precision (32-bit) floating-point elements, and store the results in "dst".
//
//go:linkname MmCvtpi8Ps MmCvtpi8Ps
//go:noescape
func MmCvtpi8Ps(r *x86.M128, v0 *x86.M64)

// Convert the lower packed unsigned 8-bit integers in "a" to packed single-precision (32-bit) floating-point elements, and store the results in "dst".
//
//go:linkname MmCvtpu8Ps MmCvtpu8Ps
//go:noescape
func MmCvtpu8Ps(r *x86.M128, v0 *x86.M64)

// Convert packed signed 32-bit integers in "a" to packed single-precision (32-bit) floating-point elements, store the results in the lower 2 elements of "dst", then covert the packed signed 32-bit integers in "b" to single-precision (32-bit) floating-point element, and store the results in the upper 2 elements of "dst".
//
//go:linkname MmCvtpi32X2Ps MmCvtpi32X2Ps
//go:noescape
func MmCvtpi32X2Ps(r *x86.M128, v0 *x86.M64, v1 *x86.M64)

// Convert packed single-precision (32-bit) floating-point elements in "a" to packed 16-bit integers, and store the results in "dst". Note: this intrinsic will generate 0x7FFF, rather than 0x8000, for input values between 0x7FFF and 0x7FFFFFFF.
//
//go:linkname MmCvtpsPi16 MmCvtpsPi16
//go:noescape
func MmCvtpsPi16(r *x86.M64, v0 *x86.M128)

// Convert packed single-precision (32-bit) floating-point elements in "a" to packed 8-bit integers, and store the results in lower 4 elements of "dst". Note: this intrinsic will generate 0x7F, rather than 0x80, for input values between 0x7F and 0x7FFFFFFF.
//
//go:linkname MmCvtpsPi8 MmCvtpsPi8
//go:noescape
func MmCvtpsPi8(r *x86.M64, v0 *x86.M128)
