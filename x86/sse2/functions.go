package sse2

import (
	"github.com/alivanz/go-simd/x86"
)

/*
#cgo CFLAGS: -msse2
#include <immintrin.h>
*/
import "C"

// Add the lower double-precision (64-bit) floating-point element in "a" and "b", store the result in the lower element of "dst", and copy the upper element from "a" to the upper element of "dst".
//
//go:linkname MmAddSd MmAddSd
//go:noescape
func MmAddSd(r *x86.M128D, v0 *x86.M128D, v1 *x86.M128D)

// Add packed double-precision (64-bit) floating-point elements in "a" and "b", and store the results in "dst".
//
//go:linkname MmAddPd MmAddPd
//go:noescape
func MmAddPd(r *x86.M128D, v0 *x86.M128D, v1 *x86.M128D)

// Subtract the lower double-precision (64-bit) floating-point element in "b" from the lower double-precision (64-bit) floating-point element in "a", store the result in the lower element of "dst", and copy the upper element from "a" to the upper element of "dst".
//
//go:linkname MmSubSd MmSubSd
//go:noescape
func MmSubSd(r *x86.M128D, v0 *x86.M128D, v1 *x86.M128D)

// Subtract packed double-precision (64-bit) floating-point elements in "b" from packed double-precision (64-bit) floating-point elements in "a", and store the results in "dst".
//
//go:linkname MmSubPd MmSubPd
//go:noescape
func MmSubPd(r *x86.M128D, v0 *x86.M128D, v1 *x86.M128D)

// Multiply the lower double-precision (64-bit) floating-point element in "a" and "b", store the result in the lower element of "dst", and copy the upper element from "a" to the upper element of "dst".
//
//go:linkname MmMulSd MmMulSd
//go:noescape
func MmMulSd(r *x86.M128D, v0 *x86.M128D, v1 *x86.M128D)

// Multiply packed double-precision (64-bit) floating-point elements in "a" and "b", and store the results in "dst".
//
//go:linkname MmMulPd MmMulPd
//go:noescape
func MmMulPd(r *x86.M128D, v0 *x86.M128D, v1 *x86.M128D)

// Divide the lower double-precision (64-bit) floating-point element in "a" by the lower double-precision (64-bit) floating-point element in "b", store the result in the lower element of "dst", and copy the upper element from "a" to the upper element of "dst".
//
//go:linkname MmDivSd MmDivSd
//go:noescape
func MmDivSd(r *x86.M128D, v0 *x86.M128D, v1 *x86.M128D)

// Divide packed double-precision (64-bit) floating-point elements in "a" by packed elements in "b", and store the results in "dst".
//
//go:linkname MmDivPd MmDivPd
//go:noescape
func MmDivPd(r *x86.M128D, v0 *x86.M128D, v1 *x86.M128D)

// Compute the square root of the lower double-precision (64-bit) floating-point element in "b", store the result in the lower element of "dst", and copy the upper element from "a" to the upper element of "dst".
//
//go:linkname MmSqrtSd MmSqrtSd
//go:noescape
func MmSqrtSd(r *x86.M128D, v0 *x86.M128D, v1 *x86.M128D)

// Compute the square root of packed double-precision (64-bit) floating-point elements in "a", and store the results in "dst".
//
//go:linkname MmSqrtPd MmSqrtPd
//go:noescape
func MmSqrtPd(r *x86.M128D, v0 *x86.M128D)

// Compare the lower double-precision (64-bit) floating-point elements in "a" and "b", store the minimum value in the lower element of "dst", and copy the upper element from "a" to the upper element of "dst". [min_float_note]
//
//go:linkname MmMinSd MmMinSd
//go:noescape
func MmMinSd(r *x86.M128D, v0 *x86.M128D, v1 *x86.M128D)

// Compare packed double-precision (64-bit) floating-point elements in "a" and "b", and store packed minimum values in "dst". [min_float_note]
//
//go:linkname MmMinPd MmMinPd
//go:noescape
func MmMinPd(r *x86.M128D, v0 *x86.M128D, v1 *x86.M128D)

// Compare the lower double-precision (64-bit) floating-point elements in "a" and "b", store the maximum value in the lower element of "dst", and copy the upper element from "a" to the upper element of "dst". [max_float_note]
//
//go:linkname MmMaxSd MmMaxSd
//go:noescape
func MmMaxSd(r *x86.M128D, v0 *x86.M128D, v1 *x86.M128D)

// Compare packed double-precision (64-bit) floating-point elements in "a" and "b", and store packed maximum values in "dst". [max_float_note]
//
//go:linkname MmMaxPd MmMaxPd
//go:noescape
func MmMaxPd(r *x86.M128D, v0 *x86.M128D, v1 *x86.M128D)

// Compute the bitwise AND of packed double-precision (64-bit) floating-point elements in "a" and "b", and store the results in "dst".
//
//go:linkname MmAndPd MmAndPd
//go:noescape
func MmAndPd(r *x86.M128D, v0 *x86.M128D, v1 *x86.M128D)

// Compute the bitwise NOT of packed double-precision (64-bit) floating-point elements in "a" and then AND with "b", and store the results in "dst".
//
//go:linkname MmAndnotPd MmAndnotPd
//go:noescape
func MmAndnotPd(r *x86.M128D, v0 *x86.M128D, v1 *x86.M128D)

// Compute the bitwise OR of packed double-precision (64-bit) floating-point elements in "a" and "b", and store the results in "dst".
//
//go:linkname MmOrPd MmOrPd
//go:noescape
func MmOrPd(r *x86.M128D, v0 *x86.M128D, v1 *x86.M128D)

// Compute the bitwise XOR of packed double-precision (64-bit) floating-point elements in "a" and "b", and store the results in "dst".
//
//go:linkname MmXorPd MmXorPd
//go:noescape
func MmXorPd(r *x86.M128D, v0 *x86.M128D, v1 *x86.M128D)

// Compare packed double-precision (64-bit) floating-point elements in "a" and "b" for equality, and store the results in "dst".
//
//go:linkname MmCmpeqPd MmCmpeqPd
//go:noescape
func MmCmpeqPd(r *x86.M128D, v0 *x86.M128D, v1 *x86.M128D)

// Compare packed double-precision (64-bit) floating-point elements in "a" and "b" for less-than, and store the results in "dst".
//
//go:linkname MmCmpltPd MmCmpltPd
//go:noescape
func MmCmpltPd(r *x86.M128D, v0 *x86.M128D, v1 *x86.M128D)

// Compare packed double-precision (64-bit) floating-point elements in "a" and "b" for less-than-or-equal, and store the results in "dst".
//
//go:linkname MmCmplePd MmCmplePd
//go:noescape
func MmCmplePd(r *x86.M128D, v0 *x86.M128D, v1 *x86.M128D)

// Compare packed double-precision (64-bit) floating-point elements in "a" and "b" for greater-than, and store the results in "dst".
//
//go:linkname MmCmpgtPd MmCmpgtPd
//go:noescape
func MmCmpgtPd(r *x86.M128D, v0 *x86.M128D, v1 *x86.M128D)

// Compare packed double-precision (64-bit) floating-point elements in "a" and "b" for greater-than-or-equal, and store the results in "dst".
//
//go:linkname MmCmpgePd MmCmpgePd
//go:noescape
func MmCmpgePd(r *x86.M128D, v0 *x86.M128D, v1 *x86.M128D)

// Compare packed double-precision (64-bit) floating-point elements in "a" and "b" to see if neither is NaN, and store the results in "dst".
//
//go:linkname MmCmpordPd MmCmpordPd
//go:noescape
func MmCmpordPd(r *x86.M128D, v0 *x86.M128D, v1 *x86.M128D)

// Compare packed double-precision (64-bit) floating-point elements in "a" and "b" to see if either is NaN, and store the results in "dst".
//
//go:linkname MmCmpunordPd MmCmpunordPd
//go:noescape
func MmCmpunordPd(r *x86.M128D, v0 *x86.M128D, v1 *x86.M128D)

// Compare packed double-precision (64-bit) floating-point elements in "a" and "b" for not-equal, and store the results in "dst".
//
//go:linkname MmCmpneqPd MmCmpneqPd
//go:noescape
func MmCmpneqPd(r *x86.M128D, v0 *x86.M128D, v1 *x86.M128D)

// Compare packed double-precision (64-bit) floating-point elements in "a" and "b" for not-less-than, and store the results in "dst".
//
//go:linkname MmCmpnltPd MmCmpnltPd
//go:noescape
func MmCmpnltPd(r *x86.M128D, v0 *x86.M128D, v1 *x86.M128D)

// Compare packed double-precision (64-bit) floating-point elements in "a" and "b" for not-less-than-or-equal, and store the results in "dst".
//
//go:linkname MmCmpnlePd MmCmpnlePd
//go:noescape
func MmCmpnlePd(r *x86.M128D, v0 *x86.M128D, v1 *x86.M128D)

// Compare packed double-precision (64-bit) floating-point elements in "a" and "b" for not-greater-than, and store the results in "dst".
//
//go:linkname MmCmpngtPd MmCmpngtPd
//go:noescape
func MmCmpngtPd(r *x86.M128D, v0 *x86.M128D, v1 *x86.M128D)

// Compare packed double-precision (64-bit) floating-point elements in "a" and "b" for not-greater-than-or-equal, and store the results in "dst".
//
//go:linkname MmCmpngePd MmCmpngePd
//go:noescape
func MmCmpngePd(r *x86.M128D, v0 *x86.M128D, v1 *x86.M128D)

// Compare the lower double-precision (64-bit) floating-point elements in "a" and "b" for equality, store the result in the lower element of "dst", and copy the upper element from "a" to the upper element of "dst".
//
//go:linkname MmCmpeqSd MmCmpeqSd
//go:noescape
func MmCmpeqSd(r *x86.M128D, v0 *x86.M128D, v1 *x86.M128D)

// Compare the lower double-precision (64-bit) floating-point elements in "a" and "b" for less-than, store the result in the lower element of "dst", and copy the upper element from "a" to the upper element of "dst".
//
//go:linkname MmCmpltSd MmCmpltSd
//go:noescape
func MmCmpltSd(r *x86.M128D, v0 *x86.M128D, v1 *x86.M128D)

// Compare the lower double-precision (64-bit) floating-point elements in "a" and "b" for less-than-or-equal, store the result in the lower element of "dst", and copy the upper element from "a" to the upper element of "dst".
//
//go:linkname MmCmpleSd MmCmpleSd
//go:noescape
func MmCmpleSd(r *x86.M128D, v0 *x86.M128D, v1 *x86.M128D)

// Compare the lower double-precision (64-bit) floating-point elements in "a" and "b" for greater-than, store the result in the lower element of "dst", and copy the upper element from "a" to the upper element of "dst".
//
//go:linkname MmCmpgtSd MmCmpgtSd
//go:noescape
func MmCmpgtSd(r *x86.M128D, v0 *x86.M128D, v1 *x86.M128D)

// Compare the lower double-precision (64-bit) floating-point elements in "a" and "b" for greater-than-or-equal, store the result in the lower element of "dst", and copy the upper element from "a" to the upper element of "dst".
//
//go:linkname MmCmpgeSd MmCmpgeSd
//go:noescape
func MmCmpgeSd(r *x86.M128D, v0 *x86.M128D, v1 *x86.M128D)

// Compare the lower double-precision (64-bit) floating-point elements in "a" and "b" to see if neither is NaN, store the result in the lower element of "dst", and copy the upper element from "a" to the upper element of "dst".
//
//go:linkname MmCmpordSd MmCmpordSd
//go:noescape
func MmCmpordSd(r *x86.M128D, v0 *x86.M128D, v1 *x86.M128D)

// Compare the lower double-precision (64-bit) floating-point elements in "a" and "b" to see if either is NaN, store the result in the lower element of "dst", and copy the upper element from "a" to the upper element of "dst".
//
//go:linkname MmCmpunordSd MmCmpunordSd
//go:noescape
func MmCmpunordSd(r *x86.M128D, v0 *x86.M128D, v1 *x86.M128D)

// Compare the lower double-precision (64-bit) floating-point elements in "a" and "b" for not-equal, store the result in the lower element of "dst", and copy the upper element from "a" to the upper element of "dst".
//
//go:linkname MmCmpneqSd MmCmpneqSd
//go:noescape
func MmCmpneqSd(r *x86.M128D, v0 *x86.M128D, v1 *x86.M128D)

// Compare the lower double-precision (64-bit) floating-point elements in "a" and "b" for not-less-than, store the result in the lower element of "dst", and copy the upper element from "a" to the upper element of "dst".
//
//go:linkname MmCmpnltSd MmCmpnltSd
//go:noescape
func MmCmpnltSd(r *x86.M128D, v0 *x86.M128D, v1 *x86.M128D)

// Compare the lower double-precision (64-bit) floating-point elements in "a" and "b" for not-less-than-or-equal, store the result in the lower element of "dst", and copy the upper element from "a" to the upper element of "dst".
//
//go:linkname MmCmpnleSd MmCmpnleSd
//go:noescape
func MmCmpnleSd(r *x86.M128D, v0 *x86.M128D, v1 *x86.M128D)

// Compare the lower double-precision (64-bit) floating-point elements in "a" and "b" for not-greater-than, store the result in the lower element of "dst", and copy the upper element from "a" to the upper element of "dst".
//
//go:linkname MmCmpngtSd MmCmpngtSd
//go:noescape
func MmCmpngtSd(r *x86.M128D, v0 *x86.M128D, v1 *x86.M128D)

// Compare the lower double-precision (64-bit) floating-point elements in "a" and "b" for not-greater-than-or-equal, store the result in the lower element of "dst", and copy the upper element from "a" to the upper element of "dst".
//
//go:linkname MmCmpngeSd MmCmpngeSd
//go:noescape
func MmCmpngeSd(r *x86.M128D, v0 *x86.M128D, v1 *x86.M128D)

// Compare the lower double-precision (64-bit) floating-point element in "a" and "b" for equality, and return the boolean result (0 or 1).
//
//go:linkname MmComieqSd MmComieqSd
//go:noescape
func MmComieqSd(r *x86.Int, v0 *x86.M128D, v1 *x86.M128D)

// Compare the lower double-precision (64-bit) floating-point element in "a" and "b" for less-than, and return the boolean result (0 or 1).
//
//go:linkname MmComiltSd MmComiltSd
//go:noescape
func MmComiltSd(r *x86.Int, v0 *x86.M128D, v1 *x86.M128D)

// Compare the lower double-precision (64-bit) floating-point element in "a" and "b" for less-than-or-equal, and return the boolean result (0 or 1).
//
//go:linkname MmComileSd MmComileSd
//go:noescape
func MmComileSd(r *x86.Int, v0 *x86.M128D, v1 *x86.M128D)

// Compare the lower double-precision (64-bit) floating-point element in "a" and "b" for greater-than, and return the boolean result (0 or 1).
//
//go:linkname MmComigtSd MmComigtSd
//go:noescape
func MmComigtSd(r *x86.Int, v0 *x86.M128D, v1 *x86.M128D)

// Compare the lower double-precision (64-bit) floating-point element in "a" and "b" for greater-than-or-equal, and return the boolean result (0 or 1).
//
//go:linkname MmComigeSd MmComigeSd
//go:noescape
func MmComigeSd(r *x86.Int, v0 *x86.M128D, v1 *x86.M128D)

// Compare the lower double-precision (64-bit) floating-point element in "a" and "b" for not-equal, and return the boolean result (0 or 1).
//
//go:linkname MmComineqSd MmComineqSd
//go:noescape
func MmComineqSd(r *x86.Int, v0 *x86.M128D, v1 *x86.M128D)

// Compare the lower double-precision (64-bit) floating-point element in "a" and "b" for equality, and return the boolean result (0 or 1). This instruction will not signal an exception for QNaNs.
//
//go:linkname MmUcomieqSd MmUcomieqSd
//go:noescape
func MmUcomieqSd(r *x86.Int, v0 *x86.M128D, v1 *x86.M128D)

// Compare the lower double-precision (64-bit) floating-point element in "a" and "b" for less-than, and return the boolean result (0 or 1). This instruction will not signal an exception for QNaNs.
//
//go:linkname MmUcomiltSd MmUcomiltSd
//go:noescape
func MmUcomiltSd(r *x86.Int, v0 *x86.M128D, v1 *x86.M128D)

// Compare the lower double-precision (64-bit) floating-point element in "a" and "b" for less-than-or-equal, and return the boolean result (0 or 1). This instruction will not signal an exception for QNaNs.
//
//go:linkname MmUcomileSd MmUcomileSd
//go:noescape
func MmUcomileSd(r *x86.Int, v0 *x86.M128D, v1 *x86.M128D)

// Compare the lower double-precision (64-bit) floating-point element in "a" and "b" for greater-than, and return the boolean result (0 or 1). This instruction will not signal an exception for QNaNs.
//
//go:linkname MmUcomigtSd MmUcomigtSd
//go:noescape
func MmUcomigtSd(r *x86.Int, v0 *x86.M128D, v1 *x86.M128D)

// Compare the lower double-precision (64-bit) floating-point element in "a" and "b" for greater-than-or-equal, and return the boolean result (0 or 1). This instruction will not signal an exception for QNaNs.
//
//go:linkname MmUcomigeSd MmUcomigeSd
//go:noescape
func MmUcomigeSd(r *x86.Int, v0 *x86.M128D, v1 *x86.M128D)

// Compare the lower double-precision (64-bit) floating-point element in "a" and "b" for not-equal, and return the boolean result (0 or 1). This instruction will not signal an exception for QNaNs.
//
//go:linkname MmUcomineqSd MmUcomineqSd
//go:noescape
func MmUcomineqSd(r *x86.Int, v0 *x86.M128D, v1 *x86.M128D)

// Convert packed double-precision (64-bit) floating-point elements in "a" to packed single-precision (32-bit) floating-point elements, and store the results in "dst".
//
//go:linkname MmCvtpdPs MmCvtpdPs
//go:noescape
func MmCvtpdPs(r *x86.M128, v0 *x86.M128D)

// Convert packed single-precision (32-bit) floating-point elements in "a" to packed double-precision (64-bit) floating-point elements, and store the results in "dst".
//
//go:linkname MmCvtpsPd MmCvtpsPd
//go:noescape
func MmCvtpsPd(r *x86.M128D, v0 *x86.M128)

// Convert packed signed 32-bit integers in "a" to packed double-precision (64-bit) floating-point elements, and store the results in "dst".
//
//go:linkname MmCvtepi32Pd MmCvtepi32Pd
//go:noescape
func MmCvtepi32Pd(r *x86.M128D, v0 *x86.M128I)

// Convert packed double-precision (64-bit) floating-point elements in "a" to packed 32-bit integers, and store the results in "dst".
//
//go:linkname MmCvtpdEpi32 MmCvtpdEpi32
//go:noescape
func MmCvtpdEpi32(r *x86.M128I, v0 *x86.M128D)

// Convert the lower double-precision (64-bit) floating-point element in "a" to a 32-bit integer, and store the result in "dst".
//
//go:linkname MmCvtsdSi32 MmCvtsdSi32
//go:noescape
func MmCvtsdSi32(r *x86.Int, v0 *x86.M128D)

// Convert the lower double-precision (64-bit) floating-point element in "b" to a single-precision (32-bit) floating-point element, store the result in the lower element of "dst", and copy the upper 3 packed elements from "a" to the upper elements of "dst".
//
//go:linkname MmCvtsdSs MmCvtsdSs
//go:noescape
func MmCvtsdSs(r *x86.M128, v0 *x86.M128, v1 *x86.M128D)

// Convert the signed 32-bit integer "b" to a double-precision (64-bit) floating-point element, store the result in the lower element of "dst", and copy the upper element from "a" to the upper element of "dst".
//
//go:linkname MmCvtsi32Sd MmCvtsi32Sd
//go:noescape
func MmCvtsi32Sd(r *x86.M128D, v0 *x86.M128D, v1 *x86.Int)

// Convert the lower single-precision (32-bit) floating-point element in "b" to a double-precision (64-bit) floating-point element, store the result in the lower element of "dst", and copy the upper element from "a" to the upper element of "dst".
//
//go:linkname MmCvtssSd MmCvtssSd
//go:noescape
func MmCvtssSd(r *x86.M128D, v0 *x86.M128D, v1 *x86.M128)

// Convert packed double-precision (64-bit) floating-point elements in "a" to packed 32-bit integers with truncation, and store the results in "dst".
//
//go:linkname MmCvttpdEpi32 MmCvttpdEpi32
//go:noescape
func MmCvttpdEpi32(r *x86.M128I, v0 *x86.M128D)

// Convert the lower double-precision (64-bit) floating-point element in "a" to a 32-bit integer with truncation, and store the result in "dst".
//
//go:linkname MmCvttsdSi32 MmCvttsdSi32
//go:noescape
func MmCvttsdSi32(r *x86.Int, v0 *x86.M128D)

// Copy the lower double-precision (64-bit) floating-point element of "a" to "dst".
//
//go:linkname MmCvtsdF64 MmCvtsdF64
//go:noescape
func MmCvtsdF64(r *x86.Double, v0 *x86.M128D)

// Return vector of type __m128d with undefined elements.
//
//go:linkname MmUndefinedPd MmUndefinedPd
//go:noescape
func MmUndefinedPd(r *x86.M128D, )

// Copy double-precision (64-bit) floating-point element "a" to the lower element of "dst", and zero the upper element.
//
//go:linkname MmSetSd MmSetSd
//go:noescape
func MmSetSd(r *x86.M128D, v0 *x86.Double)

// Broadcast double-precision (64-bit) floating-point value "a" to all elements of "dst".
//
//go:linkname MmSet1Pd MmSet1Pd
//go:noescape
func MmSet1Pd(r *x86.M128D, v0 *x86.Double)

// Broadcast double-precision (64-bit) floating-point value "a" to all elements of "dst".
//
//go:linkname MmSetPd1 MmSetPd1
//go:noescape
func MmSetPd1(r *x86.M128D, v0 *x86.Double)

// Set packed double-precision (64-bit) floating-point elements in "dst" with the supplied values.
//
//go:linkname MmSetPd MmSetPd
//go:noescape
func MmSetPd(r *x86.M128D, v0 *x86.Double, v1 *x86.Double)

// Set packed double-precision (64-bit) floating-point elements in "dst" with the supplied values in reverse order.
//
//go:linkname MmSetrPd MmSetrPd
//go:noescape
func MmSetrPd(r *x86.M128D, v0 *x86.Double, v1 *x86.Double)

// Return vector of type __m128d with all elements set to zero.
//
//go:linkname MmSetzeroPd MmSetzeroPd
//go:noescape
func MmSetzeroPd(r *x86.M128D, )

// Move the lower double-precision (64-bit) floating-point element from "b" to the lower element of "dst", and copy the upper element from "a" to the upper element of "dst".
//
//go:linkname MmMoveSd MmMoveSd
//go:noescape
func MmMoveSd(r *x86.M128D, v0 *x86.M128D, v1 *x86.M128D)

// Add packed 8-bit integers in "a" and "b", and store the results in "dst".
//
//go:linkname MmAddEpi8 MmAddEpi8
//go:noescape
func MmAddEpi8(r *x86.M128I, v0 *x86.M128I, v1 *x86.M128I)

// Add packed 16-bit integers in "a" and "b", and store the results in "dst".
//
//go:linkname MmAddEpi16 MmAddEpi16
//go:noescape
func MmAddEpi16(r *x86.M128I, v0 *x86.M128I, v1 *x86.M128I)

// Add packed 32-bit integers in "a" and "b", and store the results in "dst".
//
//go:linkname MmAddEpi32 MmAddEpi32
//go:noescape
func MmAddEpi32(r *x86.M128I, v0 *x86.M128I, v1 *x86.M128I)

// Add packed 64-bit integers in "a" and "b", and store the results in "dst".
//
//go:linkname MmAddEpi64 MmAddEpi64
//go:noescape
func MmAddEpi64(r *x86.M128I, v0 *x86.M128I, v1 *x86.M128I)

// Add packed signed 8-bit integers in "a" and "b" using saturation, and store the results in "dst".
//
//go:linkname MmAddsEpi8 MmAddsEpi8
//go:noescape
func MmAddsEpi8(r *x86.M128I, v0 *x86.M128I, v1 *x86.M128I)

// Add packed signed 16-bit integers in "a" and "b" using saturation, and store the results in "dst".
//
//go:linkname MmAddsEpi16 MmAddsEpi16
//go:noescape
func MmAddsEpi16(r *x86.M128I, v0 *x86.M128I, v1 *x86.M128I)

// Add packed unsigned 8-bit integers in "a" and "b" using saturation, and store the results in "dst".
//
//go:linkname MmAddsEpu8 MmAddsEpu8
//go:noescape
func MmAddsEpu8(r *x86.M128I, v0 *x86.M128I, v1 *x86.M128I)

// Add packed unsigned 16-bit integers in "a" and "b" using saturation, and store the results in "dst".
//
//go:linkname MmAddsEpu16 MmAddsEpu16
//go:noescape
func MmAddsEpu16(r *x86.M128I, v0 *x86.M128I, v1 *x86.M128I)

// Average packed unsigned 8-bit integers in "a" and "b", and store the results in "dst".
//
//go:linkname MmAvgEpu8 MmAvgEpu8
//go:noescape
func MmAvgEpu8(r *x86.M128I, v0 *x86.M128I, v1 *x86.M128I)

// Average packed unsigned 16-bit integers in "a" and "b", and store the results in "dst".
//
//go:linkname MmAvgEpu16 MmAvgEpu16
//go:noescape
func MmAvgEpu16(r *x86.M128I, v0 *x86.M128I, v1 *x86.M128I)

// Multiply packed signed 16-bit integers in "a" and "b", producing intermediate signed 32-bit integers. Horizontally add adjacent pairs of intermediate 32-bit integers, and pack the results in "dst".
//
//go:linkname MmMaddEpi16 MmMaddEpi16
//go:noescape
func MmMaddEpi16(r *x86.M128I, v0 *x86.M128I, v1 *x86.M128I)

// Compare packed signed 16-bit integers in "a" and "b", and store packed maximum values in "dst".
//
//go:linkname MmMaxEpi16 MmMaxEpi16
//go:noescape
func MmMaxEpi16(r *x86.M128I, v0 *x86.M128I, v1 *x86.M128I)

// Compare packed unsigned 8-bit integers in "a" and "b", and store packed maximum values in "dst".
//
//go:linkname MmMaxEpu8 MmMaxEpu8
//go:noescape
func MmMaxEpu8(r *x86.M128I, v0 *x86.M128I, v1 *x86.M128I)

// Compare packed signed 16-bit integers in "a" and "b", and store packed minimum values in "dst".
//
//go:linkname MmMinEpi16 MmMinEpi16
//go:noescape
func MmMinEpi16(r *x86.M128I, v0 *x86.M128I, v1 *x86.M128I)

// Compare packed unsigned 8-bit integers in "a" and "b", and store packed minimum values in "dst".
//
//go:linkname MmMinEpu8 MmMinEpu8
//go:noescape
func MmMinEpu8(r *x86.M128I, v0 *x86.M128I, v1 *x86.M128I)

// Multiply the packed signed 16-bit integers in "a" and "b", producing intermediate 32-bit integers, and store the high 16 bits of the intermediate integers in "dst".
//
//go:linkname MmMulhiEpi16 MmMulhiEpi16
//go:noescape
func MmMulhiEpi16(r *x86.M128I, v0 *x86.M128I, v1 *x86.M128I)

// Multiply the packed unsigned 16-bit integers in "a" and "b", producing intermediate 32-bit integers, and store the high 16 bits of the intermediate integers in "dst".
//
//go:linkname MmMulhiEpu16 MmMulhiEpu16
//go:noescape
func MmMulhiEpu16(r *x86.M128I, v0 *x86.M128I, v1 *x86.M128I)

// Multiply the packed 16-bit integers in "a" and "b", producing intermediate 32-bit integers, and store the low 16 bits of the intermediate integers in "dst".
//
//go:linkname MmMulloEpi16 MmMulloEpi16
//go:noescape
func MmMulloEpi16(r *x86.M128I, v0 *x86.M128I, v1 *x86.M128I)

// Multiply the low unsigned 32-bit integers from each packed 64-bit element in "a" and "b", and store the unsigned 64-bit results in "dst".
//
//go:linkname MmMulEpu32 MmMulEpu32
//go:noescape
func MmMulEpu32(r *x86.M128I, v0 *x86.M128I, v1 *x86.M128I)

// Compute the absolute differences of packed unsigned 8-bit integers in "a" and "b", then horizontally sum each consecutive 8 differences to produce two unsigned 16-bit integers, and pack these unsigned 16-bit integers in the low 16 bits of 64-bit elements in "dst".
//
//go:linkname MmSadEpu8 MmSadEpu8
//go:noescape
func MmSadEpu8(r *x86.M128I, v0 *x86.M128I, v1 *x86.M128I)

// Subtract packed 8-bit integers in "b" from packed 8-bit integers in "a", and store the results in "dst".
//
//go:linkname MmSubEpi8 MmSubEpi8
//go:noescape
func MmSubEpi8(r *x86.M128I, v0 *x86.M128I, v1 *x86.M128I)

// Subtract packed 16-bit integers in "b" from packed 16-bit integers in "a", and store the results in "dst".
//
//go:linkname MmSubEpi16 MmSubEpi16
//go:noescape
func MmSubEpi16(r *x86.M128I, v0 *x86.M128I, v1 *x86.M128I)

// Subtract packed 32-bit integers in "b" from packed 32-bit integers in "a", and store the results in "dst".
//
//go:linkname MmSubEpi32 MmSubEpi32
//go:noescape
func MmSubEpi32(r *x86.M128I, v0 *x86.M128I, v1 *x86.M128I)

// Subtract packed 64-bit integers in "b" from packed 64-bit integers in "a", and store the results in "dst".
//
//go:linkname MmSubEpi64 MmSubEpi64
//go:noescape
func MmSubEpi64(r *x86.M128I, v0 *x86.M128I, v1 *x86.M128I)

// Subtract packed signed 8-bit integers in "b" from packed 8-bit integers in "a" using saturation, and store the results in "dst".
//
//go:linkname MmSubsEpi8 MmSubsEpi8
//go:noescape
func MmSubsEpi8(r *x86.M128I, v0 *x86.M128I, v1 *x86.M128I)

// Subtract packed signed 16-bit integers in "b" from packed 16-bit integers in "a" using saturation, and store the results in "dst".
//
//go:linkname MmSubsEpi16 MmSubsEpi16
//go:noescape
func MmSubsEpi16(r *x86.M128I, v0 *x86.M128I, v1 *x86.M128I)

// Subtract packed unsigned 8-bit integers in "b" from packed unsigned 8-bit integers in "a" using saturation, and store the results in "dst".
//
//go:linkname MmSubsEpu8 MmSubsEpu8
//go:noescape
func MmSubsEpu8(r *x86.M128I, v0 *x86.M128I, v1 *x86.M128I)

// Subtract packed unsigned 16-bit integers in "b" from packed unsigned 16-bit integers in "a" using saturation, and store the results in "dst".
//
//go:linkname MmSubsEpu16 MmSubsEpu16
//go:noescape
func MmSubsEpu16(r *x86.M128I, v0 *x86.M128I, v1 *x86.M128I)

// Compute the bitwise AND of 128 bits (representing integer data) in "a" and "b", and store the result in "dst".
//
//go:linkname MmAndSi128 MmAndSi128
//go:noescape
func MmAndSi128(r *x86.M128I, v0 *x86.M128I, v1 *x86.M128I)

// Compute the bitwise NOT of 128 bits (representing integer data) in "a" and then AND with "b", and store the result in "dst".
//
//go:linkname MmAndnotSi128 MmAndnotSi128
//go:noescape
func MmAndnotSi128(r *x86.M128I, v0 *x86.M128I, v1 *x86.M128I)

// Compute the bitwise OR of 128 bits (representing integer data) in "a" and "b", and store the result in "dst".
//
//go:linkname MmOrSi128 MmOrSi128
//go:noescape
func MmOrSi128(r *x86.M128I, v0 *x86.M128I, v1 *x86.M128I)

// Compute the bitwise XOR of 128 bits (representing integer data) in "a" and "b", and store the result in "dst".
//
//go:linkname MmXorSi128 MmXorSi128
//go:noescape
func MmXorSi128(r *x86.M128I, v0 *x86.M128I, v1 *x86.M128I)

// Shift packed 16-bit integers in "a" left by "imm8" while shifting in zeros, and store the results in "dst".
//
//go:linkname MmSlliEpi16 MmSlliEpi16
//go:noescape
func MmSlliEpi16(r *x86.M128I, v0 *x86.M128I, v1 *x86.Int)

// Shift packed 16-bit integers in "a" left by "count" while shifting in zeros, and store the results in "dst".
//
//go:linkname MmSllEpi16 MmSllEpi16
//go:noescape
func MmSllEpi16(r *x86.M128I, v0 *x86.M128I, v1 *x86.M128I)

// Shift packed 32-bit integers in "a" left by "imm8" while shifting in zeros, and store the results in "dst".
//
//go:linkname MmSlliEpi32 MmSlliEpi32
//go:noescape
func MmSlliEpi32(r *x86.M128I, v0 *x86.M128I, v1 *x86.Int)

// Shift packed 32-bit integers in "a" left by "count" while shifting in zeros, and store the results in "dst".
//
//go:linkname MmSllEpi32 MmSllEpi32
//go:noescape
func MmSllEpi32(r *x86.M128I, v0 *x86.M128I, v1 *x86.M128I)

// Shift packed 64-bit integers in "a" left by "imm8" while shifting in zeros, and store the results in "dst".
//
//go:linkname MmSlliEpi64 MmSlliEpi64
//go:noescape
func MmSlliEpi64(r *x86.M128I, v0 *x86.M128I, v1 *x86.Int)

// Shift packed 64-bit integers in "a" left by "count" while shifting in zeros, and store the results in "dst".
//
//go:linkname MmSllEpi64 MmSllEpi64
//go:noescape
func MmSllEpi64(r *x86.M128I, v0 *x86.M128I, v1 *x86.M128I)

// Shift packed 16-bit integers in "a" right by "imm8" while shifting in sign bits, and store the results in "dst".
//
//go:linkname MmSraiEpi16 MmSraiEpi16
//go:noescape
func MmSraiEpi16(r *x86.M128I, v0 *x86.M128I, v1 *x86.Int)

// Shift packed 16-bit integers in "a" right by "count" while shifting in sign bits, and store the results in "dst".
//
//go:linkname MmSraEpi16 MmSraEpi16
//go:noescape
func MmSraEpi16(r *x86.M128I, v0 *x86.M128I, v1 *x86.M128I)

// Shift packed 32-bit integers in "a" right by "imm8" while shifting in sign bits, and store the results in "dst".
//
//go:linkname MmSraiEpi32 MmSraiEpi32
//go:noescape
func MmSraiEpi32(r *x86.M128I, v0 *x86.M128I, v1 *x86.Int)

// Shift packed 32-bit integers in "a" right by "count" while shifting in sign bits, and store the results in "dst".
//
//go:linkname MmSraEpi32 MmSraEpi32
//go:noescape
func MmSraEpi32(r *x86.M128I, v0 *x86.M128I, v1 *x86.M128I)

// Shift packed 16-bit integers in "a" right by "imm8" while shifting in zeros, and store the results in "dst".
//
//go:linkname MmSrliEpi16 MmSrliEpi16
//go:noescape
func MmSrliEpi16(r *x86.M128I, v0 *x86.M128I, v1 *x86.Int)

// Shift packed 16-bit integers in "a" right by "count" while shifting in zeros, and store the results in "dst".
//
//go:linkname MmSrlEpi16 MmSrlEpi16
//go:noescape
func MmSrlEpi16(r *x86.M128I, v0 *x86.M128I, v1 *x86.M128I)

// Shift packed 32-bit integers in "a" right by "imm8" while shifting in zeros, and store the results in "dst".
//
//go:linkname MmSrliEpi32 MmSrliEpi32
//go:noescape
func MmSrliEpi32(r *x86.M128I, v0 *x86.M128I, v1 *x86.Int)

// Shift packed 32-bit integers in "a" right by "count" while shifting in zeros, and store the results in "dst".
//
//go:linkname MmSrlEpi32 MmSrlEpi32
//go:noescape
func MmSrlEpi32(r *x86.M128I, v0 *x86.M128I, v1 *x86.M128I)

// Shift packed 64-bit integers in "a" right by "imm8" while shifting in zeros, and store the results in "dst".
//
//go:linkname MmSrliEpi64 MmSrliEpi64
//go:noescape
func MmSrliEpi64(r *x86.M128I, v0 *x86.M128I, v1 *x86.Int)

// Shift packed 64-bit integers in "a" right by "count" while shifting in zeros, and store the results in "dst".
//
//go:linkname MmSrlEpi64 MmSrlEpi64
//go:noescape
func MmSrlEpi64(r *x86.M128I, v0 *x86.M128I, v1 *x86.M128I)

// Compare packed 8-bit integers in "a" and "b" for equality, and store the results in "dst".
//
//go:linkname MmCmpeqEpi8 MmCmpeqEpi8
//go:noescape
func MmCmpeqEpi8(r *x86.M128I, v0 *x86.M128I, v1 *x86.M128I)

// Compare packed 16-bit integers in "a" and "b" for equality, and store the results in "dst".
//
//go:linkname MmCmpeqEpi16 MmCmpeqEpi16
//go:noescape
func MmCmpeqEpi16(r *x86.M128I, v0 *x86.M128I, v1 *x86.M128I)

// Compare packed 32-bit integers in "a" and "b" for equality, and store the results in "dst".
//
//go:linkname MmCmpeqEpi32 MmCmpeqEpi32
//go:noescape
func MmCmpeqEpi32(r *x86.M128I, v0 *x86.M128I, v1 *x86.M128I)

// Compare packed signed 8-bit integers in "a" and "b" for greater-than, and store the results in "dst".
//
//go:linkname MmCmpgtEpi8 MmCmpgtEpi8
//go:noescape
func MmCmpgtEpi8(r *x86.M128I, v0 *x86.M128I, v1 *x86.M128I)

// Compare packed signed 16-bit integers in "a" and "b" for greater-than, and store the results in "dst".
//
//go:linkname MmCmpgtEpi16 MmCmpgtEpi16
//go:noescape
func MmCmpgtEpi16(r *x86.M128I, v0 *x86.M128I, v1 *x86.M128I)

// Compare packed signed 32-bit integers in "a" and "b" for greater-than, and store the results in "dst".
//
//go:linkname MmCmpgtEpi32 MmCmpgtEpi32
//go:noescape
func MmCmpgtEpi32(r *x86.M128I, v0 *x86.M128I, v1 *x86.M128I)

// Compare packed signed 8-bit integers in "a" and "b" for less-than, and store the results in "dst". Note: This intrinsic emits the pcmpgtb instruction with the order of the operands switched.
//
//go:linkname MmCmpltEpi8 MmCmpltEpi8
//go:noescape
func MmCmpltEpi8(r *x86.M128I, v0 *x86.M128I, v1 *x86.M128I)

// Compare packed signed 16-bit integers in "a" and "b" for less-than, and store the results in "dst". Note: This intrinsic emits the pcmpgtw instruction with the order of the operands switched.
//
//go:linkname MmCmpltEpi16 MmCmpltEpi16
//go:noescape
func MmCmpltEpi16(r *x86.M128I, v0 *x86.M128I, v1 *x86.M128I)

// Compare packed signed 32-bit integers in "a" and "b" for less-than, and store the results in "dst". Note: This intrinsic emits the pcmpgtd instruction with the order of the operands switched.
//
//go:linkname MmCmpltEpi32 MmCmpltEpi32
//go:noescape
func MmCmpltEpi32(r *x86.M128I, v0 *x86.M128I, v1 *x86.M128I)

// Convert the signed 64-bit integer "b" to a double-precision (64-bit) floating-point element, store the result in the lower element of "dst", and copy the upper element from "a" to the upper element of "dst".
//
//go:linkname MmCvtsi64Sd MmCvtsi64Sd
//go:noescape
func MmCvtsi64Sd(r *x86.M128D, v0 *x86.M128D, v1 *x86.Longlong)

// Convert the lower double-precision (64-bit) floating-point element in "a" to a 64-bit integer, and store the result in "dst".
//
//go:linkname MmCvtsdSi64 MmCvtsdSi64
//go:noescape
func MmCvtsdSi64(r *x86.Longlong, v0 *x86.M128D)

// Convert the lower double-precision (64-bit) floating-point element in "a" to a 64-bit integer with truncation, and store the result in "dst".
//
//go:linkname MmCvttsdSi64 MmCvttsdSi64
//go:noescape
func MmCvttsdSi64(r *x86.Longlong, v0 *x86.M128D)

// Convert packed signed 32-bit integers in "a" to packed single-precision (32-bit) floating-point elements, and store the results in "dst".
//
//go:linkname MmCvtepi32Ps MmCvtepi32Ps
//go:noescape
func MmCvtepi32Ps(r *x86.M128, v0 *x86.M128I)

// Convert packed single-precision (32-bit) floating-point elements in "a" to packed 32-bit integers, and store the results in "dst".
//
//go:linkname MmCvtpsEpi32 MmCvtpsEpi32
//go:noescape
func MmCvtpsEpi32(r *x86.M128I, v0 *x86.M128)

// Convert packed single-precision (32-bit) floating-point elements in "a" to packed 32-bit integers with truncation, and store the results in "dst".
//
//go:linkname MmCvttpsEpi32 MmCvttpsEpi32
//go:noescape
func MmCvttpsEpi32(r *x86.M128I, v0 *x86.M128)

// Copy 32-bit integer "a" to the lower elements of "dst", and zero the upper elements of "dst".
//
//go:linkname MmCvtsi32Si128 MmCvtsi32Si128
//go:noescape
func MmCvtsi32Si128(r *x86.M128I, v0 *x86.Int)

// Copy 64-bit integer "a" to the lower element of "dst", and zero the upper element.
//
//go:linkname MmCvtsi64Si128 MmCvtsi64Si128
//go:noescape
func MmCvtsi64Si128(r *x86.M128I, v0 *x86.Longlong)

// Copy the lower 32-bit integer in "a" to "dst".
//
//go:linkname MmCvtsi128Si32 MmCvtsi128Si32
//go:noescape
func MmCvtsi128Si32(r *x86.Int, v0 *x86.M128I)

// Copy the lower 64-bit integer in "a" to "dst".
//
//go:linkname MmCvtsi128Si64 MmCvtsi128Si64
//go:noescape
func MmCvtsi128Si64(r *x86.Longlong, v0 *x86.M128I)

// Return vector of type __m128i with undefined elements.
//
//go:linkname MmUndefinedSi128 MmUndefinedSi128
//go:noescape
func MmUndefinedSi128(r *x86.M128I, )

// Set packed 64-bit integers in "dst" with the supplied values.
//
//go:linkname MmSetEpi64X MmSetEpi64X
//go:noescape
func MmSetEpi64X(r *x86.M128I, v0 *x86.Longlong, v1 *x86.Longlong)

// Set packed 64-bit integers in "dst" with the supplied values.
//
//go:linkname MmSetEpi64 MmSetEpi64
//go:noescape
func MmSetEpi64(r *x86.M128I, v0 *x86.M64, v1 *x86.M64)

// Set packed 32-bit integers in "dst" with the supplied values.
//
//go:linkname MmSetEpi32 MmSetEpi32
//go:noescape
func MmSetEpi32(r *x86.M128I, v0 *x86.Int, v1 *x86.Int, v2 *x86.Int, v3 *x86.Int)

// Set packed 16-bit integers in "dst" with the supplied values.
//
//go:linkname MmSetEpi16 MmSetEpi16
//go:noescape
func MmSetEpi16(r *x86.M128I, v0 *x86.Short, v1 *x86.Short, v2 *x86.Short, v3 *x86.Short, v4 *x86.Short, v5 *x86.Short, v6 *x86.Short, v7 *x86.Short)

// Set packed 8-bit integers in "dst" with the supplied values.
//
//go:linkname MmSetEpi8 MmSetEpi8
//go:noescape
func MmSetEpi8(r *x86.M128I, v0 *x86.Char, v1 *x86.Char, v2 *x86.Char, v3 *x86.Char, v4 *x86.Char, v5 *x86.Char, v6 *x86.Char, v7 *x86.Char, v8 *x86.Char, v9 *x86.Char, v10 *x86.Char, v11 *x86.Char, v12 *x86.Char, v13 *x86.Char, v14 *x86.Char, v15 *x86.Char)

// Broadcast 64-bit integer "a" to all elements of "dst". This intrinsic may generate the "vpbroadcastq".
//
//go:linkname MmSet1Epi64X MmSet1Epi64X
//go:noescape
func MmSet1Epi64X(r *x86.M128I, v0 *x86.Longlong)

// Broadcast 64-bit integer "a" to all elements of "dst".
//
//go:linkname MmSet1Epi64 MmSet1Epi64
//go:noescape
func MmSet1Epi64(r *x86.M128I, v0 *x86.M64)

// Broadcast 32-bit integer "a" to all elements of "dst". This intrinsic may generate "vpbroadcastd".
//
//go:linkname MmSet1Epi32 MmSet1Epi32
//go:noescape
func MmSet1Epi32(r *x86.M128I, v0 *x86.Int)

// Broadcast 16-bit integer "a" to all all elements of "dst". This intrinsic may generate "vpbroadcastw".
//
//go:linkname MmSet1Epi16 MmSet1Epi16
//go:noescape
func MmSet1Epi16(r *x86.M128I, v0 *x86.Short)

// Broadcast 8-bit integer "a" to all elements of "dst". This intrinsic may generate "vpbroadcastb".
//
//go:linkname MmSet1Epi8 MmSet1Epi8
//go:noescape
func MmSet1Epi8(r *x86.M128I, v0 *x86.Char)

// Set packed 64-bit integers in "dst" with the supplied values in reverse order.
//
//go:linkname MmSetrEpi64 MmSetrEpi64
//go:noescape
func MmSetrEpi64(r *x86.M128I, v0 *x86.M64, v1 *x86.M64)

// Set packed 32-bit integers in "dst" with the supplied values in reverse order.
//
//go:linkname MmSetrEpi32 MmSetrEpi32
//go:noescape
func MmSetrEpi32(r *x86.M128I, v0 *x86.Int, v1 *x86.Int, v2 *x86.Int, v3 *x86.Int)

// Set packed 16-bit integers in "dst" with the supplied values in reverse order.
//
//go:linkname MmSetrEpi16 MmSetrEpi16
//go:noescape
func MmSetrEpi16(r *x86.M128I, v0 *x86.Short, v1 *x86.Short, v2 *x86.Short, v3 *x86.Short, v4 *x86.Short, v5 *x86.Short, v6 *x86.Short, v7 *x86.Short)

// Set packed 8-bit integers in "dst" with the supplied values in reverse order.
//
//go:linkname MmSetrEpi8 MmSetrEpi8
//go:noescape
func MmSetrEpi8(r *x86.M128I, v0 *x86.Char, v1 *x86.Char, v2 *x86.Char, v3 *x86.Char, v4 *x86.Char, v5 *x86.Char, v6 *x86.Char, v7 *x86.Char, v8 *x86.Char, v9 *x86.Char, v10 *x86.Char, v11 *x86.Char, v12 *x86.Char, v13 *x86.Char, v14 *x86.Char, v15 *x86.Char)

// Return vector of type __m128i with all elements set to zero.
//
//go:linkname MmSetzeroSi128 MmSetzeroSi128
//go:noescape
func MmSetzeroSi128(r *x86.M128I, )

// Convert packed signed 16-bit integers from "a" and "b" to packed 8-bit integers using signed saturation, and store the results in "dst".
//
//go:linkname MmPacksEpi16 MmPacksEpi16
//go:noescape
func MmPacksEpi16(r *x86.M128I, v0 *x86.M128I, v1 *x86.M128I)

// Convert packed signed 32-bit integers from "a" and "b" to packed 16-bit integers using signed saturation, and store the results in "dst".
//
//go:linkname MmPacksEpi32 MmPacksEpi32
//go:noescape
func MmPacksEpi32(r *x86.M128I, v0 *x86.M128I, v1 *x86.M128I)

// Convert packed signed 16-bit integers from "a" and "b" to packed 8-bit integers using unsigned saturation, and store the results in "dst".
//
//go:linkname MmPackusEpi16 MmPackusEpi16
//go:noescape
func MmPackusEpi16(r *x86.M128I, v0 *x86.M128I, v1 *x86.M128I)

// Create mask from the most significant bit of each 8-bit element in "a", and store the result in "dst".
//
//go:linkname MmMovemaskEpi8 MmMovemaskEpi8
//go:noescape
func MmMovemaskEpi8(r *x86.Int, v0 *x86.M128I)

// Unpack and interleave 8-bit integers from the high half of "a" and "b", and store the results in "dst".
//
//go:linkname MmUnpackhiEpi8 MmUnpackhiEpi8
//go:noescape
func MmUnpackhiEpi8(r *x86.M128I, v0 *x86.M128I, v1 *x86.M128I)

// Unpack and interleave 16-bit integers from the high half of "a" and "b", and store the results in "dst".
//
//go:linkname MmUnpackhiEpi16 MmUnpackhiEpi16
//go:noescape
func MmUnpackhiEpi16(r *x86.M128I, v0 *x86.M128I, v1 *x86.M128I)

// Unpack and interleave 32-bit integers from the high half of "a" and "b", and store the results in "dst".
//
//go:linkname MmUnpackhiEpi32 MmUnpackhiEpi32
//go:noescape
func MmUnpackhiEpi32(r *x86.M128I, v0 *x86.M128I, v1 *x86.M128I)

// Unpack and interleave 64-bit integers from the high half of "a" and "b", and store the results in "dst".
//
//go:linkname MmUnpackhiEpi64 MmUnpackhiEpi64
//go:noescape
func MmUnpackhiEpi64(r *x86.M128I, v0 *x86.M128I, v1 *x86.M128I)

// Unpack and interleave 8-bit integers from the low half of "a" and "b", and store the results in "dst".
//
//go:linkname MmUnpackloEpi8 MmUnpackloEpi8
//go:noescape
func MmUnpackloEpi8(r *x86.M128I, v0 *x86.M128I, v1 *x86.M128I)

// Unpack and interleave 16-bit integers from the low half of "a" and "b", and store the results in "dst".
//
//go:linkname MmUnpackloEpi16 MmUnpackloEpi16
//go:noescape
func MmUnpackloEpi16(r *x86.M128I, v0 *x86.M128I, v1 *x86.M128I)

// Unpack and interleave 32-bit integers from the low half of "a" and "b", and store the results in "dst".
//
//go:linkname MmUnpackloEpi32 MmUnpackloEpi32
//go:noescape
func MmUnpackloEpi32(r *x86.M128I, v0 *x86.M128I, v1 *x86.M128I)

// Unpack and interleave 64-bit integers from the low half of "a" and "b", and store the results in "dst".
//
//go:linkname MmUnpackloEpi64 MmUnpackloEpi64
//go:noescape
func MmUnpackloEpi64(r *x86.M128I, v0 *x86.M128I, v1 *x86.M128I)

// Copy the lower 64-bit integer in "a" to "dst".
//
//go:linkname MmMovepi64Pi64 MmMovepi64Pi64
//go:noescape
func MmMovepi64Pi64(r *x86.M64, v0 *x86.M128I)

// Copy the 64-bit integer "a" to the lower element of "dst", and zero the upper element.
//
//go:linkname MmMovpi64Epi64 MmMovpi64Epi64
//go:noescape
func MmMovpi64Epi64(r *x86.M128I, v0 *x86.M64)

// Copy the lower 64-bit integer in "a" to the lower element of "dst", and zero the upper element.
//
//go:linkname MmMoveEpi64 MmMoveEpi64
//go:noescape
func MmMoveEpi64(r *x86.M128I, v0 *x86.M128I)

// Unpack and interleave double-precision (64-bit) floating-point elements from the high half of "a" and "b", and store the results in "dst".
//
//go:linkname MmUnpackhiPd MmUnpackhiPd
//go:noescape
func MmUnpackhiPd(r *x86.M128D, v0 *x86.M128D, v1 *x86.M128D)

// Unpack and interleave double-precision (64-bit) floating-point elements from the low half of "a" and "b", and store the results in "dst".
//
//go:linkname MmUnpackloPd MmUnpackloPd
//go:noescape
func MmUnpackloPd(r *x86.M128D, v0 *x86.M128D, v1 *x86.M128D)

// Set each bit of mask "dst" based on the most significant bit of the corresponding packed double-precision (64-bit) floating-point element in "a".
//
//go:linkname MmMovemaskPd MmMovemaskPd
//go:noescape
func MmMovemaskPd(r *x86.Int, v0 *x86.M128D)

// Cast vector of type __m128d to type __m128. This intrinsic is only used for compilation and does not generate any instructions, thus it has zero latency.
//
//go:linkname MmCastpdPs MmCastpdPs
//go:noescape
func MmCastpdPs(r *x86.M128, v0 *x86.M128D)

// Cast vector of type __m128d to type __m128i. This intrinsic is only used for compilation and does not generate any instructions, thus it has zero latency.
//
//go:linkname MmCastpdSi128 MmCastpdSi128
//go:noescape
func MmCastpdSi128(r *x86.M128I, v0 *x86.M128D)

// Cast vector of type __m128 to type __m128d. This intrinsic is only used for compilation and does not generate any instructions, thus it has zero latency.
//
//go:linkname MmCastpsPd MmCastpsPd
//go:noescape
func MmCastpsPd(r *x86.M128D, v0 *x86.M128)

// Cast vector of type __m128 to type __m128i. This intrinsic is only used for compilation and does not generate any instructions, thus it has zero latency.
//
//go:linkname MmCastpsSi128 MmCastpsSi128
//go:noescape
func MmCastpsSi128(r *x86.M128I, v0 *x86.M128)

// Cast vector of type __m128i to type __m128. This intrinsic is only used for compilation and does not generate any instructions, thus it has zero latency.
//
//go:linkname MmCastsi128Ps MmCastsi128Ps
//go:noescape
func MmCastsi128Ps(r *x86.M128, v0 *x86.M128I)

// Cast vector of type __m128i to type __m128d. This intrinsic is only used for compilation and does not generate any instructions, thus it has zero latency.
//
//go:linkname MmCastsi128Pd MmCastsi128Pd
//go:noescape
func MmCastsi128Pd(r *x86.M128D, v0 *x86.M128I)
