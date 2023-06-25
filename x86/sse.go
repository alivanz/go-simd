package x86

/*
#cgo CFLAGS: -march=native
#include <immintrin.h>
*/
import "C"

// Add the lower single-precision (32-bit) floating-point element in "a" and "b", store the result in the lower element of "dst", and copy the upper 3 packed elements from "a" to the upper elements of "dst".
func MmAddSs(v0 M128, v1 M128) M128 {
	return C._mm_add_ss(v0, v1)
}

// Add packed single-precision (32-bit) floating-point elements in "a" and "b", and store the results in "dst".
func MmAddPs(v0 M128, v1 M128) M128 {
	return C._mm_add_ps(v0, v1)
}

// Subtract the lower single-precision (32-bit) floating-point element in "b" from the lower single-precision (32-bit) floating-point element in "a", store the result in the lower element of "dst", and copy the upper 3 packed elements from "a" to the upper elements of "dst".
func MmSubSs(v0 M128, v1 M128) M128 {
	return C._mm_sub_ss(v0, v1)
}

// Subtract packed single-precision (32-bit) floating-point elements in "b" from packed single-precision (32-bit) floating-point elements in "a", and store the results in "dst".
func MmSubPs(v0 M128, v1 M128) M128 {
	return C._mm_sub_ps(v0, v1)
}

// Multiply the lower single-precision (32-bit) floating-point element in "a" and "b", store the result in the lower element of "dst", and copy the upper 3 packed elements from "a" to the upper elements of "dst".
func MmMulSs(v0 M128, v1 M128) M128 {
	return C._mm_mul_ss(v0, v1)
}

// Multiply packed single-precision (32-bit) floating-point elements in "a" and "b", and store the results in "dst".
func MmMulPs(v0 M128, v1 M128) M128 {
	return C._mm_mul_ps(v0, v1)
}

// Divide the lower single-precision (32-bit) floating-point element in "a" by the lower single-precision (32-bit) floating-point element in "b", store the result in the lower element of "dst", and copy the upper 3 packed elements from "a" to the upper elements of "dst".
func MmDivSs(v0 M128, v1 M128) M128 {
	return C._mm_div_ss(v0, v1)
}

// Divide packed single-precision (32-bit) floating-point elements in "a" by packed elements in "b", and store the results in "dst".
func MmDivPs(v0 M128, v1 M128) M128 {
	return C._mm_div_ps(v0, v1)
}

// Compute the square root of the lower single-precision (32-bit) floating-point element in "a", store the result in the lower element of "dst", and copy the upper 3 packed elements from "a" to the upper elements of "dst".
func MmSqrtSs(v0 M128) M128 {
	return C._mm_sqrt_ss(v0)
}

// Compute the square root of packed single-precision (32-bit) floating-point elements in "a", and store the results in "dst".
func MmSqrtPs(v0 M128) M128 {
	return C._mm_sqrt_ps(v0)
}

// Compute the approximate reciprocal of the lower single-precision (32-bit) floating-point element in "a", store the result in the lower element of "dst", and copy the upper 3 packed elements from "a" to the upper elements of "dst". The maximum relative error for this approximation is less than 1.5*2^-12.
func MmRcpSs(v0 M128) M128 {
	return C._mm_rcp_ss(v0)
}

// Compute the approximate reciprocal of packed single-precision (32-bit) floating-point elements in "a", and store the results in "dst". The maximum relative error for this approximation is less than 1.5*2^-12.
func MmRcpPs(v0 M128) M128 {
	return C._mm_rcp_ps(v0)
}

// Compute the approximate reciprocal square root of the lower single-precision (32-bit) floating-point element in "a", store the result in the lower element of "dst", and copy the upper 3 packed elements from "a" to the upper elements of "dst". The maximum relative error for this approximation is less than 1.5*2^-12.
func MmRsqrtSs(v0 M128) M128 {
	return C._mm_rsqrt_ss(v0)
}

// Compute the approximate reciprocal square root of packed single-precision (32-bit) floating-point elements in "a", and store the results in "dst". The maximum relative error for this approximation is less than 1.5*2^-12.
func MmRsqrtPs(v0 M128) M128 {
	return C._mm_rsqrt_ps(v0)
}

// Compare the lower single-precision (32-bit) floating-point elements in "a" and "b", store the minimum value in the lower element of "dst", and copy the upper 3 packed elements from "a" to the upper element of "dst". [min_float_note]
func MmMinSs(v0 M128, v1 M128) M128 {
	return C._mm_min_ss(v0, v1)
}

// Compare packed single-precision (32-bit) floating-point elements in "a" and "b", and store packed minimum values in "dst". [min_float_note]
func MmMinPs(v0 M128, v1 M128) M128 {
	return C._mm_min_ps(v0, v1)
}

// Compare the lower single-precision (32-bit) floating-point elements in "a" and "b", store the maximum value in the lower element of "dst", and copy the upper 3 packed elements from "a" to the upper element of "dst". [max_float_note]
func MmMaxSs(v0 M128, v1 M128) M128 {
	return C._mm_max_ss(v0, v1)
}

// Compare packed single-precision (32-bit) floating-point elements in "a" and "b", and store packed maximum values in "dst". [max_float_note]
func MmMaxPs(v0 M128, v1 M128) M128 {
	return C._mm_max_ps(v0, v1)
}

// Compute the bitwise AND of packed single-precision (32-bit) floating-point elements in "a" and "b", and store the results in "dst".
func MmAndPs(v0 M128, v1 M128) M128 {
	return C._mm_and_ps(v0, v1)
}

// Compute the bitwise NOT of packed single-precision (32-bit) floating-point elements in "a" and then AND with "b", and store the results in "dst".
func MmAndnotPs(v0 M128, v1 M128) M128 {
	return C._mm_andnot_ps(v0, v1)
}

// Compute the bitwise OR of packed single-precision (32-bit) floating-point elements in "a" and "b", and store the results in "dst".
func MmOrPs(v0 M128, v1 M128) M128 {
	return C._mm_or_ps(v0, v1)
}

// Compute the bitwise XOR of packed single-precision (32-bit) floating-point elements in "a" and "b", and store the results in "dst".
func MmXorPs(v0 M128, v1 M128) M128 {
	return C._mm_xor_ps(v0, v1)
}

// Compare the lower single-precision (32-bit) floating-point elements in "a" and "b" for equality, store the result in the lower element of "dst", and copy the upper 3 packed elements from "a" to the upper elements of "dst".
func MmCmpeqSs(v0 M128, v1 M128) M128 {
	return C._mm_cmpeq_ss(v0, v1)
}

// Compare packed single-precision (32-bit) floating-point elements in "a" and "b" for equality, and store the results in "dst".
func MmCmpeqPs(v0 M128, v1 M128) M128 {
	return C._mm_cmpeq_ps(v0, v1)
}

// Compare the lower single-precision (32-bit) floating-point elements in "a" and "b" for less-than, store the result in the lower element of "dst", and copy the upper 3 packed elements from "a" to the upper elements of "dst".
func MmCmpltSs(v0 M128, v1 M128) M128 {
	return C._mm_cmplt_ss(v0, v1)
}

// Compare packed single-precision (32-bit) floating-point elements in "a" and "b" for less-than, and store the results in "dst".
func MmCmpltPs(v0 M128, v1 M128) M128 {
	return C._mm_cmplt_ps(v0, v1)
}

// Compare the lower single-precision (32-bit) floating-point elements in "a" and "b" for less-than-or-equal, store the result in the lower element of "dst", and copy the upper 3 packed elements from "a" to the upper elements of "dst".
func MmCmpleSs(v0 M128, v1 M128) M128 {
	return C._mm_cmple_ss(v0, v1)
}

// Compare packed single-precision (32-bit) floating-point elements in "a" and "b" for less-than-or-equal, and store the results in "dst".
func MmCmplePs(v0 M128, v1 M128) M128 {
	return C._mm_cmple_ps(v0, v1)
}

// Compare the lower single-precision (32-bit) floating-point elements in "a" and "b" for greater-than, store the result in the lower element of "dst", and copy the upper 3 packed elements from "a" to the upper elements of "dst".
func MmCmpgtSs(v0 M128, v1 M128) M128 {
	return C._mm_cmpgt_ss(v0, v1)
}

// Compare packed single-precision (32-bit) floating-point elements in "a" and "b" for greater-than, and store the results in "dst".
func MmCmpgtPs(v0 M128, v1 M128) M128 {
	return C._mm_cmpgt_ps(v0, v1)
}

// Compare the lower single-precision (32-bit) floating-point elements in "a" and "b" for greater-than-or-equal, store the result in the lower element of "dst", and copy the upper 3 packed elements from "a" to the upper elements of "dst".
func MmCmpgeSs(v0 M128, v1 M128) M128 {
	return C._mm_cmpge_ss(v0, v1)
}

// Compare packed single-precision (32-bit) floating-point elements in "a" and "b" for greater-than-or-equal, and store the results in "dst".
func MmCmpgePs(v0 M128, v1 M128) M128 {
	return C._mm_cmpge_ps(v0, v1)
}

// Compare the lower single-precision (32-bit) floating-point elements in "a" and "b" for not-equal, store the result in the lower element of "dst", and copy the upper 3 packed elements from "a" to the upper elements of "dst".
func MmCmpneqSs(v0 M128, v1 M128) M128 {
	return C._mm_cmpneq_ss(v0, v1)
}

// Compare packed single-precision (32-bit) floating-point elements in "a" and "b" for not-equal, and store the results in "dst".
func MmCmpneqPs(v0 M128, v1 M128) M128 {
	return C._mm_cmpneq_ps(v0, v1)
}

// Compare the lower single-precision (32-bit) floating-point elements in "a" and "b" for not-less-than, store the result in the lower element of "dst", and copy the upper 3 packed elements from "a" to the upper elements of "dst".
func MmCmpnltSs(v0 M128, v1 M128) M128 {
	return C._mm_cmpnlt_ss(v0, v1)
}

// Compare packed single-precision (32-bit) floating-point elements in "a" and "b" for not-less-than, and store the results in "dst".
func MmCmpnltPs(v0 M128, v1 M128) M128 {
	return C._mm_cmpnlt_ps(v0, v1)
}

// Compare the lower single-precision (32-bit) floating-point elements in "a" and "b" for not-less-than-or-equal, store the result in the lower element of "dst", and copy the upper 3 packed elements from "a" to the upper elements of "dst".
func MmCmpnleSs(v0 M128, v1 M128) M128 {
	return C._mm_cmpnle_ss(v0, v1)
}

// Compare packed single-precision (32-bit) floating-point elements in "a" and "b" for not-less-than-or-equal, and store the results in "dst".
func MmCmpnlePs(v0 M128, v1 M128) M128 {
	return C._mm_cmpnle_ps(v0, v1)
}

// Compare the lower single-precision (32-bit) floating-point elements in "a" and "b" for not-greater-than, store the result in the lower element of "dst", and copy the upper 3 packed elements from "a" to the upper elements of "dst".
func MmCmpngtSs(v0 M128, v1 M128) M128 {
	return C._mm_cmpngt_ss(v0, v1)
}

// Compare packed single-precision (32-bit) floating-point elements in "a" and "b" for not-greater-than, and store the results in "dst".
func MmCmpngtPs(v0 M128, v1 M128) M128 {
	return C._mm_cmpngt_ps(v0, v1)
}

// Compare the lower single-precision (32-bit) floating-point elements in "a" and "b" for not-greater-than-or-equal, store the result in the lower element of "dst", and copy the upper 3 packed elements from "a" to the upper elements of "dst".
func MmCmpngeSs(v0 M128, v1 M128) M128 {
	return C._mm_cmpnge_ss(v0, v1)
}

// Compare packed single-precision (32-bit) floating-point elements in "a" and "b" for not-greater-than-or-equal, and store the results in "dst".
func MmCmpngePs(v0 M128, v1 M128) M128 {
	return C._mm_cmpnge_ps(v0, v1)
}

// Compare the lower single-precision (32-bit) floating-point elements in "a" and "b" to see if neither is NaN, store the result in the lower element of "dst", and copy the upper 3 packed elements from "a" to the upper elements of "dst".
func MmCmpordSs(v0 M128, v1 M128) M128 {
	return C._mm_cmpord_ss(v0, v1)
}

// Compare packed single-precision (32-bit) floating-point elements in "a" and "b" to see if neither is NaN, and store the results in "dst".
func MmCmpordPs(v0 M128, v1 M128) M128 {
	return C._mm_cmpord_ps(v0, v1)
}

// Compare the lower single-precision (32-bit) floating-point elements in "a" and "b" to see if either is NaN, store the result in the lower element of "dst", and copy the upper 3 packed elements from "a" to the upper elements of "dst".
func MmCmpunordSs(v0 M128, v1 M128) M128 {
	return C._mm_cmpunord_ss(v0, v1)
}

// Compare packed single-precision (32-bit) floating-point elements in "a" and "b" to see if either is NaN, and store the results in "dst".
func MmCmpunordPs(v0 M128, v1 M128) M128 {
	return C._mm_cmpunord_ps(v0, v1)
}

// Compare the lower single-precision (32-bit) floating-point element in "a" and "b" for equality, and return the boolean result (0 or 1).
func MmComieqSs(v0 M128, v1 M128) Int {
	return C._mm_comieq_ss(v0, v1)
}

// Compare the lower single-precision (32-bit) floating-point element in "a" and "b" for less-than, and return the boolean result (0 or 1).
func MmComiltSs(v0 M128, v1 M128) Int {
	return C._mm_comilt_ss(v0, v1)
}

// Compare the lower single-precision (32-bit) floating-point element in "a" and "b" for less-than-or-equal, and return the boolean result (0 or 1).
func MmComileSs(v0 M128, v1 M128) Int {
	return C._mm_comile_ss(v0, v1)
}

// Compare the lower single-precision (32-bit) floating-point element in "a" and "b" for greater-than, and return the boolean result (0 or 1).
func MmComigtSs(v0 M128, v1 M128) Int {
	return C._mm_comigt_ss(v0, v1)
}

// Compare the lower single-precision (32-bit) floating-point element in "a" and "b" for greater-than-or-equal, and return the boolean result (0 or 1).
func MmComigeSs(v0 M128, v1 M128) Int {
	return C._mm_comige_ss(v0, v1)
}

// Compare the lower single-precision (32-bit) floating-point element in "a" and "b" for not-equal, and return the boolean result (0 or 1).
func MmComineqSs(v0 M128, v1 M128) Int {
	return C._mm_comineq_ss(v0, v1)
}

// Compare the lower single-precision (32-bit) floating-point element in "a" and "b" for equality, and return the boolean result (0 or 1). This instruction will not signal an exception for QNaNs.
func MmUcomieqSs(v0 M128, v1 M128) Int {
	return C._mm_ucomieq_ss(v0, v1)
}

// Compare the lower single-precision (32-bit) floating-point element in "a" and "b" for less-than, and return the boolean result (0 or 1). This instruction will not signal an exception for QNaNs.
func MmUcomiltSs(v0 M128, v1 M128) Int {
	return C._mm_ucomilt_ss(v0, v1)
}

// Compare the lower single-precision (32-bit) floating-point element in "a" and "b" for less-than-or-equal, and return the boolean result (0 or 1). This instruction will not signal an exception for QNaNs.
func MmUcomileSs(v0 M128, v1 M128) Int {
	return C._mm_ucomile_ss(v0, v1)
}

// Compare the lower single-precision (32-bit) floating-point element in "a" and "b" for greater-than, and return the boolean result (0 or 1). This instruction will not signal an exception for QNaNs.
func MmUcomigtSs(v0 M128, v1 M128) Int {
	return C._mm_ucomigt_ss(v0, v1)
}

// Compare the lower single-precision (32-bit) floating-point element in "a" and "b" for greater-than-or-equal, and return the boolean result (0 or 1). This instruction will not signal an exception for QNaNs.
func MmUcomigeSs(v0 M128, v1 M128) Int {
	return C._mm_ucomige_ss(v0, v1)
}

// Compare the lower single-precision (32-bit) floating-point element in "a" and "b" for not-equal, and return the boolean result (0 or 1). This instruction will not signal an exception for QNaNs.
func MmUcomineqSs(v0 M128, v1 M128) Int {
	return C._mm_ucomineq_ss(v0, v1)
}

// Convert the lower single-precision (32-bit) floating-point element in "a" to a 32-bit integer, and store the result in "dst".
func MmCvtssSi32(v0 M128) Int {
	return C._mm_cvtss_si32(v0)
}

// Convert the lower single-precision (32-bit) floating-point element in "a" to a 32-bit integer, and store the result in "dst".
func MmCvtSs2Si(v0 M128) Int {
	return C._mm_cvt_ss2si(v0)
}

// Convert the lower single-precision (32-bit) floating-point element in "a" to a 64-bit integer, and store the result in "dst".
func MmCvtssSi64(v0 M128) Longlong {
	return C._mm_cvtss_si64(v0)
}

// Convert the lower single-precision (32-bit) floating-point element in "a" to a 32-bit integer with truncation, and store the result in "dst".
func MmCvttssSi32(v0 M128) Int {
	return C._mm_cvttss_si32(v0)
}

// Convert the lower single-precision (32-bit) floating-point element in "a" to a 32-bit integer with truncation, and store the result in "dst".
func MmCvttSs2Si(v0 M128) Int {
	return C._mm_cvtt_ss2si(v0)
}

// Convert the lower single-precision (32-bit) floating-point element in "a" to a 64-bit integer with truncation, and store the result in "dst".
func MmCvttssSi64(v0 M128) Longlong {
	return C._mm_cvttss_si64(v0)
}

// Convert the signed 32-bit integer "b" to a single-precision (32-bit) floating-point element, store the result in the lower element of "dst", and copy the upper 3 packed elements from "a" to the upper elements of "dst".
func MmCvtsi32Ss(v0 M128, v1 Int) M128 {
	return C._mm_cvtsi32_ss(v0, v1)
}

// Convert the signed 32-bit integer "b" to a single-precision (32-bit) floating-point element, store the result in the lower element of "dst", and copy the upper 3 packed elements from "a" to the upper elements of "dst".
func MmCvtSi2Ss(v0 M128, v1 Int) M128 {
	return C._mm_cvt_si2ss(v0, v1)
}

// Convert the signed 64-bit integer "b" to a single-precision (32-bit) floating-point element, store the result in the lower element of "dst", and copy the upper 3 packed elements from "a" to the upper elements of "dst".
func MmCvtsi64Ss(v0 M128, v1 Longlong) M128 {
	return C._mm_cvtsi64_ss(v0, v1)
}

// Copy the lower single-precision (32-bit) floating-point element of "a" to "dst".
func MmCvtssF32(v0 M128) Float {
	return C._mm_cvtss_f32(v0)
}

// Return vector of type __m128 with undefined elements.
func MmUndefinedPs() M128 {
	return C._mm_undefined_ps()
}

// Copy single-precision (32-bit) floating-point element "a" to the lower element of "dst", and zero the upper 3 elements.
func MmSetSs(v0 Float) M128 {
	return C._mm_set_ss(v0)
}

// Broadcast single-precision (32-bit) floating-point value "a" to all elements of "dst".
func MmSet1Ps(v0 Float) M128 {
	return C._mm_set1_ps(v0)
}

// Broadcast single-precision (32-bit) floating-point value "a" to all elements of "dst".
func MmSetPs1(v0 Float) M128 {
	return C._mm_set_ps1(v0)
}

// Set packed single-precision (32-bit) floating-point elements in "dst" with the supplied values.
func MmSetPs(v0 Float, v1 Float, v2 Float, v3 Float) M128 {
	return C._mm_set_ps(v0, v1, v2, v3)
}

// Set packed single-precision (32-bit) floating-point elements in "dst" with the supplied values in reverse order.
func MmSetrPs(v0 Float, v1 Float, v2 Float, v3 Float) M128 {
	return C._mm_setr_ps(v0, v1, v2, v3)
}

// Return vector of type __m128 with all elements set to zero.
func MmSetzeroPs() M128 {
	return C._mm_setzero_ps()
}

// Unpack and interleave single-precision (32-bit) floating-point elements from the high half "a" and "b", and store the results in "dst".
func MmUnpackhiPs(v0 M128, v1 M128) M128 {
	return C._mm_unpackhi_ps(v0, v1)
}

// Unpack and interleave single-precision (32-bit) floating-point elements from the low half of "a" and "b", and store the results in "dst".
func MmUnpackloPs(v0 M128, v1 M128) M128 {
	return C._mm_unpacklo_ps(v0, v1)
}

// Move the lower single-precision (32-bit) floating-point element from "b" to the lower element of "dst", and copy the upper 3 packed elements from "a" to the upper elements of "dst".
func MmMoveSs(v0 M128, v1 M128) M128 {
	return C._mm_move_ss(v0, v1)
}

// Move the upper 2 single-precision (32-bit) floating-point elements from "b" to the lower 2 elements of "dst", and copy the upper 2 elements from "a" to the upper 2 elements of "dst".
func MmMovehlPs(v0 M128, v1 M128) M128 {
	return C._mm_movehl_ps(v0, v1)
}

// Move the lower 2 single-precision (32-bit) floating-point elements from "b" to the upper 2 elements of "dst", and copy the lower 2 elements from "a" to the lower 2 elements of "dst".
func MmMovelhPs(v0 M128, v1 M128) M128 {
	return C._mm_movelh_ps(v0, v1)
}

// Set each bit of mask "dst" based on the most significant bit of the corresponding packed single-precision (32-bit) floating-point element in "a".
func MmMovemaskPs(v0 M128) Int {
	return C._mm_movemask_ps(v0)
}
