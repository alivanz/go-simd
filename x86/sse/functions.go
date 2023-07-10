package sse

import (
	"github.com/alivanz/go-simd/x86"
)

/*
#cgo CFLAGS: -msse
#include <immintrin.h>
*/
import "C"

// Add the lower single-precision (32-bit) floating-point element in "a" and "b", store the result in the lower element of "dst", and copy the upper 3 packed elements from "a" to the upper elements of "dst".
//
//go:linkname MmAddSs MmAddSs
//go:noescape
func MmAddSs(r *x86.M128, v0 *x86.M128, v1 *x86.M128)

// Add packed single-precision (32-bit) floating-point elements in "a" and "b", and store the results in "dst".
//
//go:linkname MmAddPs MmAddPs
//go:noescape
func MmAddPs(r *x86.M128, v0 *x86.M128, v1 *x86.M128)

// Subtract the lower single-precision (32-bit) floating-point element in "b" from the lower single-precision (32-bit) floating-point element in "a", store the result in the lower element of "dst", and copy the upper 3 packed elements from "a" to the upper elements of "dst".
//
//go:linkname MmSubSs MmSubSs
//go:noescape
func MmSubSs(r *x86.M128, v0 *x86.M128, v1 *x86.M128)

// Subtract packed single-precision (32-bit) floating-point elements in "b" from packed single-precision (32-bit) floating-point elements in "a", and store the results in "dst".
//
//go:linkname MmSubPs MmSubPs
//go:noescape
func MmSubPs(r *x86.M128, v0 *x86.M128, v1 *x86.M128)

// Multiply the lower single-precision (32-bit) floating-point element in "a" and "b", store the result in the lower element of "dst", and copy the upper 3 packed elements from "a" to the upper elements of "dst".
//
//go:linkname MmMulSs MmMulSs
//go:noescape
func MmMulSs(r *x86.M128, v0 *x86.M128, v1 *x86.M128)

// Multiply packed single-precision (32-bit) floating-point elements in "a" and "b", and store the results in "dst".
//
//go:linkname MmMulPs MmMulPs
//go:noescape
func MmMulPs(r *x86.M128, v0 *x86.M128, v1 *x86.M128)

// Divide the lower single-precision (32-bit) floating-point element in "a" by the lower single-precision (32-bit) floating-point element in "b", store the result in the lower element of "dst", and copy the upper 3 packed elements from "a" to the upper elements of "dst".
//
//go:linkname MmDivSs MmDivSs
//go:noescape
func MmDivSs(r *x86.M128, v0 *x86.M128, v1 *x86.M128)

// Divide packed single-precision (32-bit) floating-point elements in "a" by packed elements in "b", and store the results in "dst".
//
//go:linkname MmDivPs MmDivPs
//go:noescape
func MmDivPs(r *x86.M128, v0 *x86.M128, v1 *x86.M128)

// Compute the square root of the lower single-precision (32-bit) floating-point element in "a", store the result in the lower element of "dst", and copy the upper 3 packed elements from "a" to the upper elements of "dst".
//
//go:linkname MmSqrtSs MmSqrtSs
//go:noescape
func MmSqrtSs(r *x86.M128, v0 *x86.M128)

// Compute the square root of packed single-precision (32-bit) floating-point elements in "a", and store the results in "dst".
//
//go:linkname MmSqrtPs MmSqrtPs
//go:noescape
func MmSqrtPs(r *x86.M128, v0 *x86.M128)

// Compute the approximate reciprocal of the lower single-precision (32-bit) floating-point element in "a", store the result in the lower element of "dst", and copy the upper 3 packed elements from "a" to the upper elements of "dst". The maximum relative error for this approximation is less than 1.5*2^-12.
//
//go:linkname MmRcpSs MmRcpSs
//go:noescape
func MmRcpSs(r *x86.M128, v0 *x86.M128)

// Compute the approximate reciprocal of packed single-precision (32-bit) floating-point elements in "a", and store the results in "dst". The maximum relative error for this approximation is less than 1.5*2^-12.
//
//go:linkname MmRcpPs MmRcpPs
//go:noescape
func MmRcpPs(r *x86.M128, v0 *x86.M128)

// Compute the approximate reciprocal square root of the lower single-precision (32-bit) floating-point element in "a", store the result in the lower element of "dst", and copy the upper 3 packed elements from "a" to the upper elements of "dst". The maximum relative error for this approximation is less than 1.5*2^-12.
//
//go:linkname MmRsqrtSs MmRsqrtSs
//go:noescape
func MmRsqrtSs(r *x86.M128, v0 *x86.M128)

// Compute the approximate reciprocal square root of packed single-precision (32-bit) floating-point elements in "a", and store the results in "dst". The maximum relative error for this approximation is less than 1.5*2^-12.
//
//go:linkname MmRsqrtPs MmRsqrtPs
//go:noescape
func MmRsqrtPs(r *x86.M128, v0 *x86.M128)

// Compare the lower single-precision (32-bit) floating-point elements in "a" and "b", store the minimum value in the lower element of "dst", and copy the upper 3 packed elements from "a" to the upper element of "dst". [min_float_note]
//
//go:linkname MmMinSs MmMinSs
//go:noescape
func MmMinSs(r *x86.M128, v0 *x86.M128, v1 *x86.M128)

// Compare packed single-precision (32-bit) floating-point elements in "a" and "b", and store packed minimum values in "dst". [min_float_note]
//
//go:linkname MmMinPs MmMinPs
//go:noescape
func MmMinPs(r *x86.M128, v0 *x86.M128, v1 *x86.M128)

// Compare the lower single-precision (32-bit) floating-point elements in "a" and "b", store the maximum value in the lower element of "dst", and copy the upper 3 packed elements from "a" to the upper element of "dst". [max_float_note]
//
//go:linkname MmMaxSs MmMaxSs
//go:noescape
func MmMaxSs(r *x86.M128, v0 *x86.M128, v1 *x86.M128)

// Compare packed single-precision (32-bit) floating-point elements in "a" and "b", and store packed maximum values in "dst". [max_float_note]
//
//go:linkname MmMaxPs MmMaxPs
//go:noescape
func MmMaxPs(r *x86.M128, v0 *x86.M128, v1 *x86.M128)

// Compute the bitwise AND of packed single-precision (32-bit) floating-point elements in "a" and "b", and store the results in "dst".
//
//go:linkname MmAndPs MmAndPs
//go:noescape
func MmAndPs(r *x86.M128, v0 *x86.M128, v1 *x86.M128)

// Compute the bitwise NOT of packed single-precision (32-bit) floating-point elements in "a" and then AND with "b", and store the results in "dst".
//
//go:linkname MmAndnotPs MmAndnotPs
//go:noescape
func MmAndnotPs(r *x86.M128, v0 *x86.M128, v1 *x86.M128)

// Compute the bitwise OR of packed single-precision (32-bit) floating-point elements in "a" and "b", and store the results in "dst".
//
//go:linkname MmOrPs MmOrPs
//go:noescape
func MmOrPs(r *x86.M128, v0 *x86.M128, v1 *x86.M128)

// Compute the bitwise XOR of packed single-precision (32-bit) floating-point elements in "a" and "b", and store the results in "dst".
//
//go:linkname MmXorPs MmXorPs
//go:noescape
func MmXorPs(r *x86.M128, v0 *x86.M128, v1 *x86.M128)

// Compare the lower single-precision (32-bit) floating-point elements in "a" and "b" for equality, store the result in the lower element of "dst", and copy the upper 3 packed elements from "a" to the upper elements of "dst".
//
//go:linkname MmCmpeqSs MmCmpeqSs
//go:noescape
func MmCmpeqSs(r *x86.M128, v0 *x86.M128, v1 *x86.M128)

// Compare packed single-precision (32-bit) floating-point elements in "a" and "b" for equality, and store the results in "dst".
//
//go:linkname MmCmpeqPs MmCmpeqPs
//go:noescape
func MmCmpeqPs(r *x86.M128, v0 *x86.M128, v1 *x86.M128)

// Compare the lower single-precision (32-bit) floating-point elements in "a" and "b" for less-than, store the result in the lower element of "dst", and copy the upper 3 packed elements from "a" to the upper elements of "dst".
//
//go:linkname MmCmpltSs MmCmpltSs
//go:noescape
func MmCmpltSs(r *x86.M128, v0 *x86.M128, v1 *x86.M128)

// Compare packed single-precision (32-bit) floating-point elements in "a" and "b" for less-than, and store the results in "dst".
//
//go:linkname MmCmpltPs MmCmpltPs
//go:noescape
func MmCmpltPs(r *x86.M128, v0 *x86.M128, v1 *x86.M128)

// Compare the lower single-precision (32-bit) floating-point elements in "a" and "b" for less-than-or-equal, store the result in the lower element of "dst", and copy the upper 3 packed elements from "a" to the upper elements of "dst".
//
//go:linkname MmCmpleSs MmCmpleSs
//go:noescape
func MmCmpleSs(r *x86.M128, v0 *x86.M128, v1 *x86.M128)

// Compare packed single-precision (32-bit) floating-point elements in "a" and "b" for less-than-or-equal, and store the results in "dst".
//
//go:linkname MmCmplePs MmCmplePs
//go:noescape
func MmCmplePs(r *x86.M128, v0 *x86.M128, v1 *x86.M128)

// Compare the lower single-precision (32-bit) floating-point elements in "a" and "b" for greater-than, store the result in the lower element of "dst", and copy the upper 3 packed elements from "a" to the upper elements of "dst".
//
//go:linkname MmCmpgtSs MmCmpgtSs
//go:noescape
func MmCmpgtSs(r *x86.M128, v0 *x86.M128, v1 *x86.M128)

// Compare packed single-precision (32-bit) floating-point elements in "a" and "b" for greater-than, and store the results in "dst".
//
//go:linkname MmCmpgtPs MmCmpgtPs
//go:noescape
func MmCmpgtPs(r *x86.M128, v0 *x86.M128, v1 *x86.M128)

// Compare the lower single-precision (32-bit) floating-point elements in "a" and "b" for greater-than-or-equal, store the result in the lower element of "dst", and copy the upper 3 packed elements from "a" to the upper elements of "dst".
//
//go:linkname MmCmpgeSs MmCmpgeSs
//go:noescape
func MmCmpgeSs(r *x86.M128, v0 *x86.M128, v1 *x86.M128)

// Compare packed single-precision (32-bit) floating-point elements in "a" and "b" for greater-than-or-equal, and store the results in "dst".
//
//go:linkname MmCmpgePs MmCmpgePs
//go:noescape
func MmCmpgePs(r *x86.M128, v0 *x86.M128, v1 *x86.M128)

// Compare the lower single-precision (32-bit) floating-point elements in "a" and "b" for not-equal, store the result in the lower element of "dst", and copy the upper 3 packed elements from "a" to the upper elements of "dst".
//
//go:linkname MmCmpneqSs MmCmpneqSs
//go:noescape
func MmCmpneqSs(r *x86.M128, v0 *x86.M128, v1 *x86.M128)

// Compare packed single-precision (32-bit) floating-point elements in "a" and "b" for not-equal, and store the results in "dst".
//
//go:linkname MmCmpneqPs MmCmpneqPs
//go:noescape
func MmCmpneqPs(r *x86.M128, v0 *x86.M128, v1 *x86.M128)

// Compare the lower single-precision (32-bit) floating-point elements in "a" and "b" for not-less-than, store the result in the lower element of "dst", and copy the upper 3 packed elements from "a" to the upper elements of "dst".
//
//go:linkname MmCmpnltSs MmCmpnltSs
//go:noescape
func MmCmpnltSs(r *x86.M128, v0 *x86.M128, v1 *x86.M128)

// Compare packed single-precision (32-bit) floating-point elements in "a" and "b" for not-less-than, and store the results in "dst".
//
//go:linkname MmCmpnltPs MmCmpnltPs
//go:noescape
func MmCmpnltPs(r *x86.M128, v0 *x86.M128, v1 *x86.M128)

// Compare the lower single-precision (32-bit) floating-point elements in "a" and "b" for not-less-than-or-equal, store the result in the lower element of "dst", and copy the upper 3 packed elements from "a" to the upper elements of "dst".
//
//go:linkname MmCmpnleSs MmCmpnleSs
//go:noescape
func MmCmpnleSs(r *x86.M128, v0 *x86.M128, v1 *x86.M128)

// Compare packed single-precision (32-bit) floating-point elements in "a" and "b" for not-less-than-or-equal, and store the results in "dst".
//
//go:linkname MmCmpnlePs MmCmpnlePs
//go:noescape
func MmCmpnlePs(r *x86.M128, v0 *x86.M128, v1 *x86.M128)

// Compare the lower single-precision (32-bit) floating-point elements in "a" and "b" for not-greater-than, store the result in the lower element of "dst", and copy the upper 3 packed elements from "a" to the upper elements of "dst".
//
//go:linkname MmCmpngtSs MmCmpngtSs
//go:noescape
func MmCmpngtSs(r *x86.M128, v0 *x86.M128, v1 *x86.M128)

// Compare packed single-precision (32-bit) floating-point elements in "a" and "b" for not-greater-than, and store the results in "dst".
//
//go:linkname MmCmpngtPs MmCmpngtPs
//go:noescape
func MmCmpngtPs(r *x86.M128, v0 *x86.M128, v1 *x86.M128)

// Compare the lower single-precision (32-bit) floating-point elements in "a" and "b" for not-greater-than-or-equal, store the result in the lower element of "dst", and copy the upper 3 packed elements from "a" to the upper elements of "dst".
//
//go:linkname MmCmpngeSs MmCmpngeSs
//go:noescape
func MmCmpngeSs(r *x86.M128, v0 *x86.M128, v1 *x86.M128)

// Compare packed single-precision (32-bit) floating-point elements in "a" and "b" for not-greater-than-or-equal, and store the results in "dst".
//
//go:linkname MmCmpngePs MmCmpngePs
//go:noescape
func MmCmpngePs(r *x86.M128, v0 *x86.M128, v1 *x86.M128)

// Compare the lower single-precision (32-bit) floating-point elements in "a" and "b" to see if neither is NaN, store the result in the lower element of "dst", and copy the upper 3 packed elements from "a" to the upper elements of "dst".
//
//go:linkname MmCmpordSs MmCmpordSs
//go:noescape
func MmCmpordSs(r *x86.M128, v0 *x86.M128, v1 *x86.M128)

// Compare packed single-precision (32-bit) floating-point elements in "a" and "b" to see if neither is NaN, and store the results in "dst".
//
//go:linkname MmCmpordPs MmCmpordPs
//go:noescape
func MmCmpordPs(r *x86.M128, v0 *x86.M128, v1 *x86.M128)

// Compare the lower single-precision (32-bit) floating-point elements in "a" and "b" to see if either is NaN, store the result in the lower element of "dst", and copy the upper 3 packed elements from "a" to the upper elements of "dst".
//
//go:linkname MmCmpunordSs MmCmpunordSs
//go:noescape
func MmCmpunordSs(r *x86.M128, v0 *x86.M128, v1 *x86.M128)

// Compare packed single-precision (32-bit) floating-point elements in "a" and "b" to see if either is NaN, and store the results in "dst".
//
//go:linkname MmCmpunordPs MmCmpunordPs
//go:noescape
func MmCmpunordPs(r *x86.M128, v0 *x86.M128, v1 *x86.M128)

// Compare the lower single-precision (32-bit) floating-point element in "a" and "b" for equality, and return the boolean result (0 or 1).
//
//go:linkname MmComieqSs MmComieqSs
//go:noescape
func MmComieqSs(r *x86.Int, v0 *x86.M128, v1 *x86.M128)

// Compare the lower single-precision (32-bit) floating-point element in "a" and "b" for less-than, and return the boolean result (0 or 1).
//
//go:linkname MmComiltSs MmComiltSs
//go:noescape
func MmComiltSs(r *x86.Int, v0 *x86.M128, v1 *x86.M128)

// Compare the lower single-precision (32-bit) floating-point element in "a" and "b" for less-than-or-equal, and return the boolean result (0 or 1).
//
//go:linkname MmComileSs MmComileSs
//go:noescape
func MmComileSs(r *x86.Int, v0 *x86.M128, v1 *x86.M128)

// Compare the lower single-precision (32-bit) floating-point element in "a" and "b" for greater-than, and return the boolean result (0 or 1).
//
//go:linkname MmComigtSs MmComigtSs
//go:noescape
func MmComigtSs(r *x86.Int, v0 *x86.M128, v1 *x86.M128)

// Compare the lower single-precision (32-bit) floating-point element in "a" and "b" for greater-than-or-equal, and return the boolean result (0 or 1).
//
//go:linkname MmComigeSs MmComigeSs
//go:noescape
func MmComigeSs(r *x86.Int, v0 *x86.M128, v1 *x86.M128)

// Compare the lower single-precision (32-bit) floating-point element in "a" and "b" for not-equal, and return the boolean result (0 or 1).
//
//go:linkname MmComineqSs MmComineqSs
//go:noescape
func MmComineqSs(r *x86.Int, v0 *x86.M128, v1 *x86.M128)

// Compare the lower single-precision (32-bit) floating-point element in "a" and "b" for equality, and return the boolean result (0 or 1). This instruction will not signal an exception for QNaNs.
//
//go:linkname MmUcomieqSs MmUcomieqSs
//go:noescape
func MmUcomieqSs(r *x86.Int, v0 *x86.M128, v1 *x86.M128)

// Compare the lower single-precision (32-bit) floating-point element in "a" and "b" for less-than, and return the boolean result (0 or 1). This instruction will not signal an exception for QNaNs.
//
//go:linkname MmUcomiltSs MmUcomiltSs
//go:noescape
func MmUcomiltSs(r *x86.Int, v0 *x86.M128, v1 *x86.M128)

// Compare the lower single-precision (32-bit) floating-point element in "a" and "b" for less-than-or-equal, and return the boolean result (0 or 1). This instruction will not signal an exception for QNaNs.
//
//go:linkname MmUcomileSs MmUcomileSs
//go:noescape
func MmUcomileSs(r *x86.Int, v0 *x86.M128, v1 *x86.M128)

// Compare the lower single-precision (32-bit) floating-point element in "a" and "b" for greater-than, and return the boolean result (0 or 1). This instruction will not signal an exception for QNaNs.
//
//go:linkname MmUcomigtSs MmUcomigtSs
//go:noescape
func MmUcomigtSs(r *x86.Int, v0 *x86.M128, v1 *x86.M128)

// Compare the lower single-precision (32-bit) floating-point element in "a" and "b" for greater-than-or-equal, and return the boolean result (0 or 1). This instruction will not signal an exception for QNaNs.
//
//go:linkname MmUcomigeSs MmUcomigeSs
//go:noescape
func MmUcomigeSs(r *x86.Int, v0 *x86.M128, v1 *x86.M128)

// Compare the lower single-precision (32-bit) floating-point element in "a" and "b" for not-equal, and return the boolean result (0 or 1). This instruction will not signal an exception for QNaNs.
//
//go:linkname MmUcomineqSs MmUcomineqSs
//go:noescape
func MmUcomineqSs(r *x86.Int, v0 *x86.M128, v1 *x86.M128)

// Convert the lower single-precision (32-bit) floating-point element in "a" to a 32-bit integer, and store the result in "dst".
//
//go:linkname MmCvtssSi32 MmCvtssSi32
//go:noescape
func MmCvtssSi32(r *x86.Int, v0 *x86.M128)

// Convert the lower single-precision (32-bit) floating-point element in "a" to a 32-bit integer, and store the result in "dst".
//
//go:linkname MmCvtSs2Si MmCvtSs2Si
//go:noescape
func MmCvtSs2Si(r *x86.Int, v0 *x86.M128)

// Convert the lower single-precision (32-bit) floating-point element in "a" to a 64-bit integer, and store the result in "dst".
//
//go:linkname MmCvtssSi64 MmCvtssSi64
//go:noescape
func MmCvtssSi64(r *x86.Longlong, v0 *x86.M128)

// Convert the lower single-precision (32-bit) floating-point element in "a" to a 32-bit integer with truncation, and store the result in "dst".
//
//go:linkname MmCvttssSi32 MmCvttssSi32
//go:noescape
func MmCvttssSi32(r *x86.Int, v0 *x86.M128)

// Convert the lower single-precision (32-bit) floating-point element in "a" to a 32-bit integer with truncation, and store the result in "dst".
//
//go:linkname MmCvttSs2Si MmCvttSs2Si
//go:noescape
func MmCvttSs2Si(r *x86.Int, v0 *x86.M128)

// Convert the lower single-precision (32-bit) floating-point element in "a" to a 64-bit integer with truncation, and store the result in "dst".
//
//go:linkname MmCvttssSi64 MmCvttssSi64
//go:noescape
func MmCvttssSi64(r *x86.Longlong, v0 *x86.M128)

// Convert the signed 32-bit integer "b" to a single-precision (32-bit) floating-point element, store the result in the lower element of "dst", and copy the upper 3 packed elements from "a" to the upper elements of "dst".
//
//go:linkname MmCvtsi32Ss MmCvtsi32Ss
//go:noescape
func MmCvtsi32Ss(r *x86.M128, v0 *x86.M128, v1 *x86.Int)

// Convert the signed 32-bit integer "b" to a single-precision (32-bit) floating-point element, store the result in the lower element of "dst", and copy the upper 3 packed elements from "a" to the upper elements of "dst".
//
//go:linkname MmCvtSi2Ss MmCvtSi2Ss
//go:noescape
func MmCvtSi2Ss(r *x86.M128, v0 *x86.M128, v1 *x86.Int)

// Convert the signed 64-bit integer "b" to a single-precision (32-bit) floating-point element, store the result in the lower element of "dst", and copy the upper 3 packed elements from "a" to the upper elements of "dst".
//
//go:linkname MmCvtsi64Ss MmCvtsi64Ss
//go:noescape
func MmCvtsi64Ss(r *x86.M128, v0 *x86.M128, v1 *x86.Longlong)

// Copy the lower single-precision (32-bit) floating-point element of "a" to "dst".
//
//go:linkname MmCvtssF32 MmCvtssF32
//go:noescape
func MmCvtssF32(r *x86.Float, v0 *x86.M128)

// Return vector of type __m128 with undefined elements.
//
//go:linkname MmUndefinedPs MmUndefinedPs
//go:noescape
func MmUndefinedPs(r *x86.M128, )

// Copy single-precision (32-bit) floating-point element "a" to the lower element of "dst", and zero the upper 3 elements.
//
//go:linkname MmSetSs MmSetSs
//go:noescape
func MmSetSs(r *x86.M128, v0 *x86.Float)

// Broadcast single-precision (32-bit) floating-point value "a" to all elements of "dst".
//
//go:linkname MmSet1Ps MmSet1Ps
//go:noescape
func MmSet1Ps(r *x86.M128, v0 *x86.Float)

// Broadcast single-precision (32-bit) floating-point value "a" to all elements of "dst".
//
//go:linkname MmSetPs1 MmSetPs1
//go:noescape
func MmSetPs1(r *x86.M128, v0 *x86.Float)

// Set packed single-precision (32-bit) floating-point elements in "dst" with the supplied values.
//
//go:linkname MmSetPs MmSetPs
//go:noescape
func MmSetPs(r *x86.M128, v0 *x86.Float, v1 *x86.Float, v2 *x86.Float, v3 *x86.Float)

// Set packed single-precision (32-bit) floating-point elements in "dst" with the supplied values in reverse order.
//
//go:linkname MmSetrPs MmSetrPs
//go:noescape
func MmSetrPs(r *x86.M128, v0 *x86.Float, v1 *x86.Float, v2 *x86.Float, v3 *x86.Float)

// Return vector of type __m128 with all elements set to zero.
//
//go:linkname MmSetzeroPs MmSetzeroPs
//go:noescape
func MmSetzeroPs(r *x86.M128, )

// Unpack and interleave single-precision (32-bit) floating-point elements from the high half "a" and "b", and store the results in "dst".
//
//go:linkname MmUnpackhiPs MmUnpackhiPs
//go:noescape
func MmUnpackhiPs(r *x86.M128, v0 *x86.M128, v1 *x86.M128)

// Unpack and interleave single-precision (32-bit) floating-point elements from the low half of "a" and "b", and store the results in "dst".
//
//go:linkname MmUnpackloPs MmUnpackloPs
//go:noescape
func MmUnpackloPs(r *x86.M128, v0 *x86.M128, v1 *x86.M128)

// Move the lower single-precision (32-bit) floating-point element from "b" to the lower element of "dst", and copy the upper 3 packed elements from "a" to the upper elements of "dst".
//
//go:linkname MmMoveSs MmMoveSs
//go:noescape
func MmMoveSs(r *x86.M128, v0 *x86.M128, v1 *x86.M128)

// Move the upper 2 single-precision (32-bit) floating-point elements from "b" to the lower 2 elements of "dst", and copy the upper 2 elements from "a" to the upper 2 elements of "dst".
//
//go:linkname MmMovehlPs MmMovehlPs
//go:noescape
func MmMovehlPs(r *x86.M128, v0 *x86.M128, v1 *x86.M128)

// Move the lower 2 single-precision (32-bit) floating-point elements from "b" to the upper 2 elements of "dst", and copy the lower 2 elements from "a" to the lower 2 elements of "dst".
//
//go:linkname MmMovelhPs MmMovelhPs
//go:noescape
func MmMovelhPs(r *x86.M128, v0 *x86.M128, v1 *x86.M128)

// Set each bit of mask "dst" based on the most significant bit of the corresponding packed single-precision (32-bit) floating-point element in "a".
//
//go:linkname MmMovemaskPs MmMovemaskPs
//go:noescape
func MmMovemaskPs(r *x86.Int, v0 *x86.M128)
