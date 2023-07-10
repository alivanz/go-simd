package avx

import (
	"github.com/alivanz/go-simd/x86"
)

/*
#cgo CFLAGS: -mavx
#include <immintrin.h>
*/
import "C"

// Add packed double-precision (64-bit) floating-point elements in "a" and "b", and store the results in "dst".
//
//go:linkname Mm256AddPd Mm256AddPd
//go:noescape
func Mm256AddPd(r *x86.M256D, v0 *x86.M256D, v1 *x86.M256D)

// Add packed single-precision (32-bit) floating-point elements in "a" and "b", and store the results in "dst".
//
//go:linkname Mm256AddPs Mm256AddPs
//go:noescape
func Mm256AddPs(r *x86.M256, v0 *x86.M256, v1 *x86.M256)

// Subtract packed double-precision (64-bit) floating-point elements in "b" from packed double-precision (64-bit) floating-point elements in "a", and store the results in "dst".
//
//go:linkname Mm256SubPd Mm256SubPd
//go:noescape
func Mm256SubPd(r *x86.M256D, v0 *x86.M256D, v1 *x86.M256D)

// Subtract packed single-precision (32-bit) floating-point elements in "b" from packed single-precision (32-bit) floating-point elements in "a", and store the results in "dst".
//
//go:linkname Mm256SubPs Mm256SubPs
//go:noescape
func Mm256SubPs(r *x86.M256, v0 *x86.M256, v1 *x86.M256)

// Alternatively add and subtract packed double-precision (64-bit) floating-point elements in "a" to/from packed elements in "b", and store the results in "dst".
//
//go:linkname Mm256AddsubPd Mm256AddsubPd
//go:noescape
func Mm256AddsubPd(r *x86.M256D, v0 *x86.M256D, v1 *x86.M256D)

// Alternatively add and subtract packed single-precision (32-bit) floating-point elements in "a" to/from packed elements in "b", and store the results in "dst".
//
//go:linkname Mm256AddsubPs Mm256AddsubPs
//go:noescape
func Mm256AddsubPs(r *x86.M256, v0 *x86.M256, v1 *x86.M256)

// Divide packed double-precision (64-bit) floating-point elements in "a" by packed elements in "b", and store the results in "dst".
//
//go:linkname Mm256DivPd Mm256DivPd
//go:noescape
func Mm256DivPd(r *x86.M256D, v0 *x86.M256D, v1 *x86.M256D)

// Divide packed single-precision (32-bit) floating-point elements in "a" by packed elements in "b", and store the results in "dst".
//
//go:linkname Mm256DivPs Mm256DivPs
//go:noescape
func Mm256DivPs(r *x86.M256, v0 *x86.M256, v1 *x86.M256)

// Compare packed double-precision (64-bit) floating-point elements in "a" and "b", and store packed maximum values in "dst". [max_float_note]
//
//go:linkname Mm256MaxPd Mm256MaxPd
//go:noescape
func Mm256MaxPd(r *x86.M256D, v0 *x86.M256D, v1 *x86.M256D)

// Compare packed single-precision (32-bit) floating-point elements in "a" and "b", and store packed maximum values in "dst". [max_float_note]
//
//go:linkname Mm256MaxPs Mm256MaxPs
//go:noescape
func Mm256MaxPs(r *x86.M256, v0 *x86.M256, v1 *x86.M256)

// Compare packed double-precision (64-bit) floating-point elements in "a" and "b", and store packed minimum values in "dst". [min_float_note]
//
//go:linkname Mm256MinPd Mm256MinPd
//go:noescape
func Mm256MinPd(r *x86.M256D, v0 *x86.M256D, v1 *x86.M256D)

// Compare packed single-precision (32-bit) floating-point elements in "a" and "b", and store packed minimum values in "dst". [min_float_note]
//
//go:linkname Mm256MinPs Mm256MinPs
//go:noescape
func Mm256MinPs(r *x86.M256, v0 *x86.M256, v1 *x86.M256)

// Multiply packed double-precision (64-bit) floating-point elements in "a" and "b", and store the results in "dst".
//
//go:linkname Mm256MulPd Mm256MulPd
//go:noescape
func Mm256MulPd(r *x86.M256D, v0 *x86.M256D, v1 *x86.M256D)

// Multiply packed single-precision (32-bit) floating-point elements in "a" and "b", and store the results in "dst".
//
//go:linkname Mm256MulPs Mm256MulPs
//go:noescape
func Mm256MulPs(r *x86.M256, v0 *x86.M256, v1 *x86.M256)

// Compute the square root of packed double-precision (64-bit) floating-point elements in "a", and store the results in "dst".
//
//go:linkname Mm256SqrtPd Mm256SqrtPd
//go:noescape
func Mm256SqrtPd(r *x86.M256D, v0 *x86.M256D)

// Compute the square root of packed single-precision (32-bit) floating-point elements in "a", and store the results in "dst".
//
//go:linkname Mm256SqrtPs Mm256SqrtPs
//go:noescape
func Mm256SqrtPs(r *x86.M256, v0 *x86.M256)

// Compute the approximate reciprocal square root of packed single-precision (32-bit) floating-point elements in "a", and store the results in "dst". The maximum relative error for this approximation is less than 1.5*2^-12.
//
//go:linkname Mm256RsqrtPs Mm256RsqrtPs
//go:noescape
func Mm256RsqrtPs(r *x86.M256, v0 *x86.M256)

// Compute the approximate reciprocal of packed single-precision (32-bit) floating-point elements in "a", and store the results in "dst". The maximum relative error for this approximation is less than 1.5*2^-12.
//
//go:linkname Mm256RcpPs Mm256RcpPs
//go:noescape
func Mm256RcpPs(r *x86.M256, v0 *x86.M256)

// Compute the bitwise AND of packed double-precision (64-bit) floating-point elements in "a" and "b", and store the results in "dst".
//
//go:linkname Mm256AndPd Mm256AndPd
//go:noescape
func Mm256AndPd(r *x86.M256D, v0 *x86.M256D, v1 *x86.M256D)

// Compute the bitwise AND of packed single-precision (32-bit) floating-point elements in "a" and "b", and store the results in "dst".
//
//go:linkname Mm256AndPs Mm256AndPs
//go:noescape
func Mm256AndPs(r *x86.M256, v0 *x86.M256, v1 *x86.M256)

// Compute the bitwise NOT of packed double-precision (64-bit) floating-point elements in "a" and then AND with "b", and store the results in "dst".
//
//go:linkname Mm256AndnotPd Mm256AndnotPd
//go:noescape
func Mm256AndnotPd(r *x86.M256D, v0 *x86.M256D, v1 *x86.M256D)

// Compute the bitwise NOT of packed single-precision (32-bit) floating-point elements in "a" and then AND with "b", and store the results in "dst".
//
//go:linkname Mm256AndnotPs Mm256AndnotPs
//go:noescape
func Mm256AndnotPs(r *x86.M256, v0 *x86.M256, v1 *x86.M256)

// Compute the bitwise OR of packed double-precision (64-bit) floating-point elements in "a" and "b", and store the results in "dst".
//
//go:linkname Mm256OrPd Mm256OrPd
//go:noescape
func Mm256OrPd(r *x86.M256D, v0 *x86.M256D, v1 *x86.M256D)

// Compute the bitwise OR of packed single-precision (32-bit) floating-point elements in "a" and "b", and store the results in "dst".
//
//go:linkname Mm256OrPs Mm256OrPs
//go:noescape
func Mm256OrPs(r *x86.M256, v0 *x86.M256, v1 *x86.M256)

// Compute the bitwise XOR of packed double-precision (64-bit) floating-point elements in "a" and "b", and store the results in "dst".
//
//go:linkname Mm256XorPd Mm256XorPd
//go:noescape
func Mm256XorPd(r *x86.M256D, v0 *x86.M256D, v1 *x86.M256D)

// Compute the bitwise XOR of packed single-precision (32-bit) floating-point elements in "a" and "b", and store the results in "dst".
//
//go:linkname Mm256XorPs Mm256XorPs
//go:noescape
func Mm256XorPs(r *x86.M256, v0 *x86.M256, v1 *x86.M256)

// Horizontally add adjacent pairs of double-precision (64-bit) floating-point elements in "a" and "b", and pack the results in "dst".
//
//go:linkname Mm256HaddPd Mm256HaddPd
//go:noescape
func Mm256HaddPd(r *x86.M256D, v0 *x86.M256D, v1 *x86.M256D)

// Horizontally add adjacent pairs of single-precision (32-bit) floating-point elements in "a" and "b", and pack the results in "dst".
//
//go:linkname Mm256HaddPs Mm256HaddPs
//go:noescape
func Mm256HaddPs(r *x86.M256, v0 *x86.M256, v1 *x86.M256)

// Horizontally subtract adjacent pairs of double-precision (64-bit) floating-point elements in "a" and "b", and pack the results in "dst".
//
//go:linkname Mm256HsubPd Mm256HsubPd
//go:noescape
func Mm256HsubPd(r *x86.M256D, v0 *x86.M256D, v1 *x86.M256D)

// Horizontally subtract adjacent pairs of single-precision (32-bit) floating-point elements in "a" and "b", and pack the results in "dst".
//
//go:linkname Mm256HsubPs Mm256HsubPs
//go:noescape
func Mm256HsubPs(r *x86.M256, v0 *x86.M256, v1 *x86.M256)

// Shuffle double-precision (64-bit) floating-point elements in "a" using the control in "b", and store the results in "dst".
//
//go:linkname MmPermutevarPd MmPermutevarPd
//go:noescape
func MmPermutevarPd(r *x86.M128D, v0 *x86.M128D, v1 *x86.M128I)

// Shuffle double-precision (64-bit) floating-point elements in "a" within 128-bit lanes using the control in "b", and store the results in "dst".
//
//go:linkname Mm256PermutevarPd Mm256PermutevarPd
//go:noescape
func Mm256PermutevarPd(r *x86.M256D, v0 *x86.M256D, v1 *x86.M256I)

// Shuffle single-precision (32-bit) floating-point elements in "a" using the control in "b", and store the results in "dst".
//
//go:linkname MmPermutevarPs MmPermutevarPs
//go:noescape
func MmPermutevarPs(r *x86.M128, v0 *x86.M128, v1 *x86.M128I)

// Shuffle single-precision (32-bit) floating-point elements in "a" within 128-bit lanes using the control in "b", and store the results in "dst".
//
//go:linkname Mm256PermutevarPs Mm256PermutevarPs
//go:noescape
func Mm256PermutevarPs(r *x86.M256, v0 *x86.M256, v1 *x86.M256I)

// Blend packed double-precision (64-bit) floating-point elements from "a" and "b" using "mask", and store the results in "dst".
//
//go:linkname Mm256BlendvPd Mm256BlendvPd
//go:noescape
func Mm256BlendvPd(r *x86.M256D, v0 *x86.M256D, v1 *x86.M256D, v2 *x86.M256D)

// Blend packed single-precision (32-bit) floating-point elements from "a" and "b" using "mask", and store the results in "dst".
//
//go:linkname Mm256BlendvPs Mm256BlendvPs
//go:noescape
func Mm256BlendvPs(r *x86.M256, v0 *x86.M256, v1 *x86.M256, v2 *x86.M256)

// Convert packed signed 32-bit integers in "a" to packed double-precision (64-bit) floating-point elements, and store the results in "dst".
//
//go:linkname Mm256Cvtepi32Pd Mm256Cvtepi32Pd
//go:noescape
func Mm256Cvtepi32Pd(r *x86.M256D, v0 *x86.M128I)

// Convert packed signed 32-bit integers in "a" to packed single-precision (32-bit) floating-point elements, and store the results in "dst".
//
//go:linkname Mm256Cvtepi32Ps Mm256Cvtepi32Ps
//go:noescape
func Mm256Cvtepi32Ps(r *x86.M256, v0 *x86.M256I)

// Convert packed double-precision (64-bit) floating-point elements in "a" to packed single-precision (32-bit) floating-point elements, and store the results in "dst".
//
//go:linkname Mm256CvtpdPs Mm256CvtpdPs
//go:noescape
func Mm256CvtpdPs(r *x86.M128, v0 *x86.M256D)

// Convert packed single-precision (32-bit) floating-point elements in "a" to packed 32-bit integers, and store the results in "dst".
//
//go:linkname Mm256CvtpsEpi32 Mm256CvtpsEpi32
//go:noescape
func Mm256CvtpsEpi32(r *x86.M256I, v0 *x86.M256)

// Convert packed single-precision (32-bit) floating-point elements in "a" to packed double-precision (64-bit) floating-point elements, and store the results in "dst".
//
//go:linkname Mm256CvtpsPd Mm256CvtpsPd
//go:noescape
func Mm256CvtpsPd(r *x86.M256D, v0 *x86.M128)

// Convert packed double-precision (64-bit) floating-point elements in "a" to packed 32-bit integers with truncation, and store the results in "dst".
//
//go:linkname Mm256CvttpdEpi32 Mm256CvttpdEpi32
//go:noescape
func Mm256CvttpdEpi32(r *x86.M128I, v0 *x86.M256D)

// Convert packed double-precision (64-bit) floating-point elements in "a" to packed 32-bit integers, and store the results in "dst".
//
//go:linkname Mm256CvtpdEpi32 Mm256CvtpdEpi32
//go:noescape
func Mm256CvtpdEpi32(r *x86.M128I, v0 *x86.M256D)

// Convert packed single-precision (32-bit) floating-point elements in "a" to packed 32-bit integers with truncation, and store the results in "dst".
//
//go:linkname Mm256CvttpsEpi32 Mm256CvttpsEpi32
//go:noescape
func Mm256CvttpsEpi32(r *x86.M256I, v0 *x86.M256)

// Copy the lower double-precision (64-bit) floating-point element of "a" to "dst".
//
//go:linkname Mm256CvtsdF64 Mm256CvtsdF64
//go:noescape
func Mm256CvtsdF64(r *x86.Double, v0 *x86.M256D)

// Copy the lower 32-bit integer in "a" to "dst".
//
//go:linkname Mm256Cvtsi256Si32 Mm256Cvtsi256Si32
//go:noescape
func Mm256Cvtsi256Si32(r *x86.Int, v0 *x86.M256I)

// Copy the lower single-precision (32-bit) floating-point element of "a" to "dst".
//
//go:linkname Mm256CvtssF32 Mm256CvtssF32
//go:noescape
func Mm256CvtssF32(r *x86.Float, v0 *x86.M256)

// Duplicate odd-indexed single-precision (32-bit) floating-point elements from "a", and store the results in "dst".
//
//go:linkname Mm256MovehdupPs Mm256MovehdupPs
//go:noescape
func Mm256MovehdupPs(r *x86.M256, v0 *x86.M256)

// Duplicate even-indexed single-precision (32-bit) floating-point elements from "a", and store the results in "dst".
//
//go:linkname Mm256MoveldupPs Mm256MoveldupPs
//go:noescape
func Mm256MoveldupPs(r *x86.M256, v0 *x86.M256)

// Duplicate even-indexed double-precision (64-bit) floating-point elements from "a", and store the results in "dst".
//
//go:linkname Mm256MovedupPd Mm256MovedupPd
//go:noescape
func Mm256MovedupPd(r *x86.M256D, v0 *x86.M256D)

// Unpack and interleave double-precision (64-bit) floating-point elements from the high half of each 128-bit lane in "a" and "b", and store the results in "dst".
//
//go:linkname Mm256UnpackhiPd Mm256UnpackhiPd
//go:noescape
func Mm256UnpackhiPd(r *x86.M256D, v0 *x86.M256D, v1 *x86.M256D)

// Unpack and interleave double-precision (64-bit) floating-point elements from the low half of each 128-bit lane in "a" and "b", and store the results in "dst".
//
//go:linkname Mm256UnpackloPd Mm256UnpackloPd
//go:noescape
func Mm256UnpackloPd(r *x86.M256D, v0 *x86.M256D, v1 *x86.M256D)

// Unpack and interleave single-precision (32-bit) floating-point elements from the high half of each 128-bit lane in "a" and "b", and store the results in "dst".
//
//go:linkname Mm256UnpackhiPs Mm256UnpackhiPs
//go:noescape
func Mm256UnpackhiPs(r *x86.M256, v0 *x86.M256, v1 *x86.M256)

// Unpack and interleave single-precision (32-bit) floating-point elements from the low half of each 128-bit lane in "a" and "b", and store the results in "dst".
//
//go:linkname Mm256UnpackloPs Mm256UnpackloPs
//go:noescape
func Mm256UnpackloPs(r *x86.M256, v0 *x86.M256, v1 *x86.M256)

// Compute the bitwise AND of 128 bits (representing double-precision (64-bit) floating-point elements) in "a" and "b", producing an intermediate 128-bit value, and set "ZF" to 1 if the sign bit of each 64-bit element in the intermediate value is zero, otherwise set "ZF" to 0. Compute the bitwise NOT of "a" and then AND with "b", producing an intermediate value, and set "CF" to 1 if the sign bit of each 64-bit element in the intermediate value is zero, otherwise set "CF" to 0. Return the "ZF" value.
//
//go:linkname MmTestzPd MmTestzPd
//go:noescape
func MmTestzPd(r *x86.Int, v0 *x86.M128D, v1 *x86.M128D)

// Compute the bitwise AND of 128 bits (representing double-precision (64-bit) floating-point elements) in "a" and "b", producing an intermediate 128-bit value, and set "ZF" to 1 if the sign bit of each 64-bit element in the intermediate value is zero, otherwise set "ZF" to 0. Compute the bitwise NOT of "a" and then AND with "b", producing an intermediate value, and set "CF" to 1 if the sign bit of each 64-bit element in the intermediate value is zero, otherwise set "CF" to 0. Return the "CF" value.
//
//go:linkname MmTestcPd MmTestcPd
//go:noescape
func MmTestcPd(r *x86.Int, v0 *x86.M128D, v1 *x86.M128D)

// Compute the bitwise AND of 128 bits (representing double-precision (64-bit) floating-point elements) in "a" and "b", producing an intermediate 128-bit value, and set "ZF" to 1 if the sign bit of each 64-bit element in the intermediate value is zero, otherwise set "ZF" to 0. Compute the bitwise NOT of "a" and then AND with "b", producing an intermediate value, and set "CF" to 1 if the sign bit of each 64-bit element in the intermediate value is zero, otherwise set "CF" to 0. Return 1 if both the "ZF" and "CF" values are zero, otherwise return 0.
//
//go:linkname MmTestnzcPd MmTestnzcPd
//go:noescape
func MmTestnzcPd(r *x86.Int, v0 *x86.M128D, v1 *x86.M128D)

// Compute the bitwise AND of 128 bits (representing single-precision (32-bit) floating-point elements) in "a" and "b", producing an intermediate 128-bit value, and set "ZF" to 1 if the sign bit of each 32-bit element in the intermediate value is zero, otherwise set "ZF" to 0. Compute the bitwise NOT of "a" and then AND with "b", producing an intermediate value, and set "CF" to 1 if the sign bit of each 32-bit element in the intermediate value is zero, otherwise set "CF" to 0. Return the "ZF" value.
//
//go:linkname MmTestzPs MmTestzPs
//go:noescape
func MmTestzPs(r *x86.Int, v0 *x86.M128, v1 *x86.M128)

// Compute the bitwise AND of 128 bits (representing single-precision (32-bit) floating-point elements) in "a" and "b", producing an intermediate 128-bit value, and set "ZF" to 1 if the sign bit of each 32-bit element in the intermediate value is zero, otherwise set "ZF" to 0. Compute the bitwise NOT of "a" and then AND with "b", producing an intermediate value, and set "CF" to 1 if the sign bit of each 32-bit element in the intermediate value is zero, otherwise set "CF" to 0. Return the "CF" value.
//
//go:linkname MmTestcPs MmTestcPs
//go:noescape
func MmTestcPs(r *x86.Int, v0 *x86.M128, v1 *x86.M128)

// Compute the bitwise AND of 128 bits (representing single-precision (32-bit) floating-point elements) in "a" and "b", producing an intermediate 128-bit value, and set "ZF" to 1 if the sign bit of each 32-bit element in the intermediate value is zero, otherwise set "ZF" to 0. Compute the bitwise NOT of "a" and then AND with "b", producing an intermediate value, and set "CF" to 1 if the sign bit of each 32-bit element in the intermediate value is zero, otherwise set "CF" to 0. Return 1 if both the "ZF" and "CF" values are zero, otherwise return 0.
//
//go:linkname MmTestnzcPs MmTestnzcPs
//go:noescape
func MmTestnzcPs(r *x86.Int, v0 *x86.M128, v1 *x86.M128)

// Compute the bitwise AND of 256 bits (representing double-precision (64-bit) floating-point elements) in "a" and "b", producing an intermediate 256-bit value, and set "ZF" to 1 if the sign bit of each 64-bit element in the intermediate value is zero, otherwise set "ZF" to 0. Compute the bitwise NOT of "a" and then AND with "b", producing an intermediate value, and set "CF" to 1 if the sign bit of each 64-bit element in the intermediate value is zero, otherwise set "CF" to 0. Return the "ZF" value.
//
//go:linkname Mm256TestzPd Mm256TestzPd
//go:noescape
func Mm256TestzPd(r *x86.Int, v0 *x86.M256D, v1 *x86.M256D)

// Compute the bitwise AND of 256 bits (representing double-precision (64-bit) floating-point elements) in "a" and "b", producing an intermediate 256-bit value, and set "ZF" to 1 if the sign bit of each 64-bit element in the intermediate value is zero, otherwise set "ZF" to 0. Compute the bitwise NOT of "a" and then AND with "b", producing an intermediate value, and set "CF" to 1 if the sign bit of each 64-bit element in the intermediate value is zero, otherwise set "CF" to 0. Return the "CF" value.
//
//go:linkname Mm256TestcPd Mm256TestcPd
//go:noescape
func Mm256TestcPd(r *x86.Int, v0 *x86.M256D, v1 *x86.M256D)

// Compute the bitwise AND of 256 bits (representing double-precision (64-bit) floating-point elements) in "a" and "b", producing an intermediate 256-bit value, and set "ZF" to 1 if the sign bit of each 64-bit element in the intermediate value is zero, otherwise set "ZF" to 0. Compute the bitwise NOT of "a" and then AND with "b", producing an intermediate value, and set "CF" to 1 if the sign bit of each 64-bit element in the intermediate value is zero, otherwise set "CF" to 0. Return 1 if both the "ZF" and "CF" values are zero, otherwise return 0.
//
//go:linkname Mm256TestnzcPd Mm256TestnzcPd
//go:noescape
func Mm256TestnzcPd(r *x86.Int, v0 *x86.M256D, v1 *x86.M256D)

// Compute the bitwise AND of 256 bits (representing single-precision (32-bit) floating-point elements) in "a" and "b", producing an intermediate 256-bit value, and set "ZF" to 1 if the sign bit of each 32-bit element in the intermediate value is zero, otherwise set "ZF" to 0. Compute the bitwise NOT of "a" and then AND with "b", producing an intermediate value, and set "CF" to 1 if the sign bit of each 32-bit element in the intermediate value is zero, otherwise set "CF" to 0. Return the "ZF" value.
//
//go:linkname Mm256TestzPs Mm256TestzPs
//go:noescape
func Mm256TestzPs(r *x86.Int, v0 *x86.M256, v1 *x86.M256)

// Compute the bitwise AND of 256 bits (representing single-precision (32-bit) floating-point elements) in "a" and "b", producing an intermediate 256-bit value, and set "ZF" to 1 if the sign bit of each 32-bit element in the intermediate value is zero, otherwise set "ZF" to 0. Compute the bitwise NOT of "a" and then AND with "b", producing an intermediate value, and set "CF" to 1 if the sign bit of each 32-bit element in the intermediate value is zero, otherwise set "CF" to 0. Return the "CF" value.
//
//go:linkname Mm256TestcPs Mm256TestcPs
//go:noescape
func Mm256TestcPs(r *x86.Int, v0 *x86.M256, v1 *x86.M256)

// Compute the bitwise AND of 256 bits (representing single-precision (32-bit) floating-point elements) in "a" and "b", producing an intermediate 256-bit value, and set "ZF" to 1 if the sign bit of each 32-bit element in the intermediate value is zero, otherwise set "ZF" to 0. Compute the bitwise NOT of "a" and then AND with "b", producing an intermediate value, and set "CF" to 1 if the sign bit of each 32-bit element in the intermediate value is zero, otherwise set "CF" to 0. Return 1 if both the "ZF" and "CF" values are zero, otherwise return 0.
//
//go:linkname Mm256TestnzcPs Mm256TestnzcPs
//go:noescape
func Mm256TestnzcPs(r *x86.Int, v0 *x86.M256, v1 *x86.M256)

// Compute the bitwise AND of 256 bits (representing integer data) in "a" and "b", and set "ZF" to 1 if the result is zero, otherwise set "ZF" to 0. Compute the bitwise NOT of "a" and then AND with "b", and set "CF" to 1 if the result is zero, otherwise set "CF" to 0. Return the "ZF" value.
//
//go:linkname Mm256TestzSi256 Mm256TestzSi256
//go:noescape
func Mm256TestzSi256(r *x86.Int, v0 *x86.M256I, v1 *x86.M256I)

// Compute the bitwise AND of 256 bits (representing integer data) in "a" and "b", and set "ZF" to 1 if the result is zero, otherwise set "ZF" to 0. Compute the bitwise NOT of "a" and then AND with "b", and set "CF" to 1 if the result is zero, otherwise set "CF" to 0. Return the "CF" value.
//
//go:linkname Mm256TestcSi256 Mm256TestcSi256
//go:noescape
func Mm256TestcSi256(r *x86.Int, v0 *x86.M256I, v1 *x86.M256I)

// Compute the bitwise AND of 256 bits (representing integer data) in "a" and "b", and set "ZF" to 1 if the result is zero, otherwise set "ZF" to 0. Compute the bitwise NOT of "a" and then AND with "b", and set "CF" to 1 if the result is zero, otherwise set "CF" to 0. Return 1 if both the "ZF" and "CF" values are zero, otherwise return 0.
//
//go:linkname Mm256TestnzcSi256 Mm256TestnzcSi256
//go:noescape
func Mm256TestnzcSi256(r *x86.Int, v0 *x86.M256I, v1 *x86.M256I)

// Set each bit of mask "dst" based on the most significant bit of the corresponding packed double-precision (64-bit) floating-point element in "a".
//
//go:linkname Mm256MovemaskPd Mm256MovemaskPd
//go:noescape
func Mm256MovemaskPd(r *x86.Int, v0 *x86.M256D)

// Set each bit of mask "dst" based on the most significant bit of the corresponding packed single-precision (32-bit) floating-point element in "a".
//
//go:linkname Mm256MovemaskPs Mm256MovemaskPs
//go:noescape
func Mm256MovemaskPs(r *x86.Int, v0 *x86.M256)

// Zero the contents of all XMM or YMM registers.
//
//go:linkname Mm256Zeroall Mm256Zeroall
//go:noescape
func Mm256Zeroall()

// Zero the upper 128 bits of all YMM registers; the lower 128-bits of the registers are unmodified.
//
//go:linkname Mm256Zeroupper Mm256Zeroupper
//go:noescape
func Mm256Zeroupper()

// Return vector of type __m256d with undefined elements.
//
//go:linkname Mm256UndefinedPd Mm256UndefinedPd
//go:noescape
func Mm256UndefinedPd(r *x86.M256D, )

// Return vector of type __m256 with undefined elements.
//
//go:linkname Mm256UndefinedPs Mm256UndefinedPs
//go:noescape
func Mm256UndefinedPs(r *x86.M256, )

// Return vector of type __m256i with undefined elements.
//
//go:linkname Mm256UndefinedSi256 Mm256UndefinedSi256
//go:noescape
func Mm256UndefinedSi256(r *x86.M256I, )

// Set packed double-precision (64-bit) floating-point elements in "dst" with the supplied values.
//
//go:linkname Mm256SetPd Mm256SetPd
//go:noescape
func Mm256SetPd(r *x86.M256D, v0 *x86.Double, v1 *x86.Double, v2 *x86.Double, v3 *x86.Double)

// Set packed single-precision (32-bit) floating-point elements in "dst" with the supplied values.
//
//go:linkname Mm256SetPs Mm256SetPs
//go:noescape
func Mm256SetPs(r *x86.M256, v0 *x86.Float, v1 *x86.Float, v2 *x86.Float, v3 *x86.Float, v4 *x86.Float, v5 *x86.Float, v6 *x86.Float, v7 *x86.Float)

// Set packed 32-bit integers in "dst" with the supplied values.
//
//go:linkname Mm256SetEpi32 Mm256SetEpi32
//go:noescape
func Mm256SetEpi32(r *x86.M256I, v0 *x86.Int, v1 *x86.Int, v2 *x86.Int, v3 *x86.Int, v4 *x86.Int, v5 *x86.Int, v6 *x86.Int, v7 *x86.Int)

// Set packed 16-bit integers in "dst" with the supplied values.
//
//go:linkname Mm256SetEpi16 Mm256SetEpi16
//go:noescape
func Mm256SetEpi16(r *x86.M256I, v0 *x86.Short, v1 *x86.Short, v2 *x86.Short, v3 *x86.Short, v4 *x86.Short, v5 *x86.Short, v6 *x86.Short, v7 *x86.Short, v8 *x86.Short, v9 *x86.Short, v10 *x86.Short, v11 *x86.Short, v12 *x86.Short, v13 *x86.Short, v14 *x86.Short, v15 *x86.Short)

// Set packed 8-bit integers in "dst" with the supplied values.
//
//go:linkname Mm256SetEpi8 Mm256SetEpi8
//go:noescape
func Mm256SetEpi8(r *x86.M256I, v0 *x86.Char, v1 *x86.Char, v2 *x86.Char, v3 *x86.Char, v4 *x86.Char, v5 *x86.Char, v6 *x86.Char, v7 *x86.Char, v8 *x86.Char, v9 *x86.Char, v10 *x86.Char, v11 *x86.Char, v12 *x86.Char, v13 *x86.Char, v14 *x86.Char, v15 *x86.Char, v16 *x86.Char, v17 *x86.Char, v18 *x86.Char, v19 *x86.Char, v20 *x86.Char, v21 *x86.Char, v22 *x86.Char, v23 *x86.Char, v24 *x86.Char, v25 *x86.Char, v26 *x86.Char, v27 *x86.Char, v28 *x86.Char, v29 *x86.Char, v30 *x86.Char, v31 *x86.Char)

// Set packed 64-bit integers in "dst" with the supplied values.
//
//go:linkname Mm256SetEpi64X Mm256SetEpi64X
//go:noescape
func Mm256SetEpi64X(r *x86.M256I, v0 *x86.Longlong, v1 *x86.Longlong, v2 *x86.Longlong, v3 *x86.Longlong)

// Set packed double-precision (64-bit) floating-point elements in "dst" with the supplied values in reverse order.
//
//go:linkname Mm256SetrPd Mm256SetrPd
//go:noescape
func Mm256SetrPd(r *x86.M256D, v0 *x86.Double, v1 *x86.Double, v2 *x86.Double, v3 *x86.Double)

// Set packed single-precision (32-bit) floating-point elements in "dst" with the supplied values in reverse order.
//
//go:linkname Mm256SetrPs Mm256SetrPs
//go:noescape
func Mm256SetrPs(r *x86.M256, v0 *x86.Float, v1 *x86.Float, v2 *x86.Float, v3 *x86.Float, v4 *x86.Float, v5 *x86.Float, v6 *x86.Float, v7 *x86.Float)

// Set packed 32-bit integers in "dst" with the supplied values in reverse order.
//
//go:linkname Mm256SetrEpi32 Mm256SetrEpi32
//go:noescape
func Mm256SetrEpi32(r *x86.M256I, v0 *x86.Int, v1 *x86.Int, v2 *x86.Int, v3 *x86.Int, v4 *x86.Int, v5 *x86.Int, v6 *x86.Int, v7 *x86.Int)

// Set packed 16-bit integers in "dst" with the supplied values in reverse order.
//
//go:linkname Mm256SetrEpi16 Mm256SetrEpi16
//go:noescape
func Mm256SetrEpi16(r *x86.M256I, v0 *x86.Short, v1 *x86.Short, v2 *x86.Short, v3 *x86.Short, v4 *x86.Short, v5 *x86.Short, v6 *x86.Short, v7 *x86.Short, v8 *x86.Short, v9 *x86.Short, v10 *x86.Short, v11 *x86.Short, v12 *x86.Short, v13 *x86.Short, v14 *x86.Short, v15 *x86.Short)

// Set packed 8-bit integers in "dst" with the supplied values in reverse order.
//
//go:linkname Mm256SetrEpi8 Mm256SetrEpi8
//go:noescape
func Mm256SetrEpi8(r *x86.M256I, v0 *x86.Char, v1 *x86.Char, v2 *x86.Char, v3 *x86.Char, v4 *x86.Char, v5 *x86.Char, v6 *x86.Char, v7 *x86.Char, v8 *x86.Char, v9 *x86.Char, v10 *x86.Char, v11 *x86.Char, v12 *x86.Char, v13 *x86.Char, v14 *x86.Char, v15 *x86.Char, v16 *x86.Char, v17 *x86.Char, v18 *x86.Char, v19 *x86.Char, v20 *x86.Char, v21 *x86.Char, v22 *x86.Char, v23 *x86.Char, v24 *x86.Char, v25 *x86.Char, v26 *x86.Char, v27 *x86.Char, v28 *x86.Char, v29 *x86.Char, v30 *x86.Char, v31 *x86.Char)

// Set packed 64-bit integers in "dst" with the supplied values in reverse order.
//
//go:linkname Mm256SetrEpi64X Mm256SetrEpi64X
//go:noescape
func Mm256SetrEpi64X(r *x86.M256I, v0 *x86.Longlong, v1 *x86.Longlong, v2 *x86.Longlong, v3 *x86.Longlong)

// Broadcast double-precision (64-bit) floating-point value "a" to all elements of "dst".
//
//go:linkname Mm256Set1Pd Mm256Set1Pd
//go:noescape
func Mm256Set1Pd(r *x86.M256D, v0 *x86.Double)

// Broadcast single-precision (32-bit) floating-point value "a" to all elements of "dst".
//
//go:linkname Mm256Set1Ps Mm256Set1Ps
//go:noescape
func Mm256Set1Ps(r *x86.M256, v0 *x86.Float)

// Broadcast 32-bit integer "a" to all elements of "dst". This intrinsic may generate the "vpbroadcastd".
//
//go:linkname Mm256Set1Epi32 Mm256Set1Epi32
//go:noescape
func Mm256Set1Epi32(r *x86.M256I, v0 *x86.Int)

// Broadcast 16-bit integer "a" to all all elements of "dst". This intrinsic may generate the "vpbroadcastw".
//
//go:linkname Mm256Set1Epi16 Mm256Set1Epi16
//go:noescape
func Mm256Set1Epi16(r *x86.M256I, v0 *x86.Short)

// Broadcast 8-bit integer "a" to all elements of "dst". This intrinsic may generate the "vpbroadcastb".
//
//go:linkname Mm256Set1Epi8 Mm256Set1Epi8
//go:noescape
func Mm256Set1Epi8(r *x86.M256I, v0 *x86.Char)

// Broadcast 64-bit integer "a" to all elements of "dst". This intrinsic may generate the "vpbroadcastq".
//
//go:linkname Mm256Set1Epi64X Mm256Set1Epi64X
//go:noescape
func Mm256Set1Epi64X(r *x86.M256I, v0 *x86.Longlong)

// Return vector of type __m256d with all elements set to zero.
//
//go:linkname Mm256SetzeroPd Mm256SetzeroPd
//go:noescape
func Mm256SetzeroPd(r *x86.M256D, )

// Return vector of type __m256 with all elements set to zero.
//
//go:linkname Mm256SetzeroPs Mm256SetzeroPs
//go:noescape
func Mm256SetzeroPs(r *x86.M256, )

// Return vector of type __m256i with all elements set to zero.
//
//go:linkname Mm256SetzeroSi256 Mm256SetzeroSi256
//go:noescape
func Mm256SetzeroSi256(r *x86.M256I, )

// Cast vector of type __m256d to type __m256.	This intrinsic is only used for compilation and does not generate any instructions, thus it has zero latency.
//
//go:linkname Mm256CastpdPs Mm256CastpdPs
//go:noescape
func Mm256CastpdPs(r *x86.M256, v0 *x86.M256D)

// Cast vector of type __m256d to type __m256i. This intrinsic is only used for compilation and does not generate any instructions, thus it has zero latency.
//
//go:linkname Mm256CastpdSi256 Mm256CastpdSi256
//go:noescape
func Mm256CastpdSi256(r *x86.M256I, v0 *x86.M256D)

// Cast vector of type __m256 to type __m256d.	This intrinsic is only used for compilation and does not generate any instructions, thus it has zero latency.
//
//go:linkname Mm256CastpsPd Mm256CastpsPd
//go:noescape
func Mm256CastpsPd(r *x86.M256D, v0 *x86.M256)

// Cast vector of type __m256 to type __m256i. This intrinsic is only used for compilation and does not generate any instructions, thus it has zero latency.
//
//go:linkname Mm256CastpsSi256 Mm256CastpsSi256
//go:noescape
func Mm256CastpsSi256(r *x86.M256I, v0 *x86.M256)

// Cast vector of type __m256i to type __m256. This intrinsic is only used for compilation and does not generate any instructions, thus it has zero latency.
//
//go:linkname Mm256Castsi256Ps Mm256Castsi256Ps
//go:noescape
func Mm256Castsi256Ps(r *x86.M256, v0 *x86.M256I)

// Cast vector of type __m256i to type __m256d. This intrinsic is only used for compilation and does not generate any instructions, thus it has zero latency.
//
//go:linkname Mm256Castsi256Pd Mm256Castsi256Pd
//go:noescape
func Mm256Castsi256Pd(r *x86.M256D, v0 *x86.M256I)

// Cast vector of type __m256d to type __m128d. This intrinsic is only used for compilation and does not generate any instructions, thus it has zero latency.
//
//go:linkname Mm256Castpd256Pd128 Mm256Castpd256Pd128
//go:noescape
func Mm256Castpd256Pd128(r *x86.M128D, v0 *x86.M256D)

// Cast vector of type __m256 to type __m128. This intrinsic is only used for compilation and does not generate any instructions, thus it has zero latency.
//
//go:linkname Mm256Castps256Ps128 Mm256Castps256Ps128
//go:noescape
func Mm256Castps256Ps128(r *x86.M128, v0 *x86.M256)

// Cast vector of type __m256i to type __m128i. This intrinsic is only used for compilation and does not generate any instructions, thus it has zero latency.
//
//go:linkname Mm256Castsi256Si128 Mm256Castsi256Si128
//go:noescape
func Mm256Castsi256Si128(r *x86.M128I, v0 *x86.M256I)

// Cast vector of type __m128d to type __m256d; the upper 128 bits of the result are undefined. This intrinsic is only used for compilation and does not generate any instructions, thus it has zero latency.
//
//go:linkname Mm256Castpd128Pd256 Mm256Castpd128Pd256
//go:noescape
func Mm256Castpd128Pd256(r *x86.M256D, v0 *x86.M128D)

// Cast vector of type __m128 to type __m256; the upper 128 bits of the result are undefined. This intrinsic is only used for compilation and does not generate any instructions, thus it has zero latency.
//
//go:linkname Mm256Castps128Ps256 Mm256Castps128Ps256
//go:noescape
func Mm256Castps128Ps256(r *x86.M256, v0 *x86.M128)

// Cast vector of type __m128i to type __m256i; the upper 128 bits of the result are undefined. This intrinsic is only used for compilation and does not generate any instructions, thus it has zero latency.
//
//go:linkname Mm256Castsi128Si256 Mm256Castsi128Si256
//go:noescape
func Mm256Castsi128Si256(r *x86.M256I, v0 *x86.M128I)

// Cast vector of type __m128d to type __m256d; the upper 128 bits of the result are zeroed. This intrinsic is only used for compilation and does not generate any instructions, thus it has zero latency.
//
//go:linkname Mm256Zextpd128Pd256 Mm256Zextpd128Pd256
//go:noescape
func Mm256Zextpd128Pd256(r *x86.M256D, v0 *x86.M128D)

// Cast vector of type __m128 to type __m256; the upper 128 bits of the result are zeroed. This intrinsic is only used for compilation and does not generate any instructions, thus it has zero latency.
//
//go:linkname Mm256Zextps128Ps256 Mm256Zextps128Ps256
//go:noescape
func Mm256Zextps128Ps256(r *x86.M256, v0 *x86.M128)

// Cast vector of type __m128i to type __m256i; the upper 128 bits of the result are zeroed. This intrinsic is only used for compilation and does not generate any instructions, thus it has zero latency.
//
//go:linkname Mm256Zextsi128Si256 Mm256Zextsi128Si256
//go:noescape
func Mm256Zextsi128Si256(r *x86.M256I, v0 *x86.M128I)

// Set packed __m256 vector "dst" with the supplied values.
//
//go:linkname Mm256SetM128 Mm256SetM128
//go:noescape
func Mm256SetM128(r *x86.M256, v0 *x86.M128, v1 *x86.M128)

// Set packed __m256d vector "dst" with the supplied values.
//
//go:linkname Mm256SetM128D Mm256SetM128D
//go:noescape
func Mm256SetM128D(r *x86.M256D, v0 *x86.M128D, v1 *x86.M128D)

// Set packed __m256i vector "dst" with the supplied values.
//
//go:linkname Mm256SetM128I Mm256SetM128I
//go:noescape
func Mm256SetM128I(r *x86.M256I, v0 *x86.M128I, v1 *x86.M128I)

// Set packed __m256 vector "dst" with the supplied values.
//
//go:linkname Mm256SetrM128 Mm256SetrM128
//go:noescape
func Mm256SetrM128(r *x86.M256, v0 *x86.M128, v1 *x86.M128)

// Set packed __m256d vector "dst" with the supplied values.
//
//go:linkname Mm256SetrM128D Mm256SetrM128D
//go:noescape
func Mm256SetrM128D(r *x86.M256D, v0 *x86.M128D, v1 *x86.M128D)

// Set packed __m256i vector "dst" with the supplied values.
//
//go:linkname Mm256SetrM128I Mm256SetrM128I
//go:noescape
func Mm256SetrM128I(r *x86.M256I, v0 *x86.M128I, v1 *x86.M128I)
