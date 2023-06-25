package x86

/*
#cgo CFLAGS: -march=native
#include <immintrin.h>
*/
import "C"

// Add the lower double-precision (64-bit) floating-point element in "a" and "b", store the result in the lower element of "dst", and copy the upper element from "a" to the upper element of "dst".
func MmAddSd(v0 M128D, v1 M128D) M128D {
	return C._mm_add_sd(v0, v1)
}

// Add packed double-precision (64-bit) floating-point elements in "a" and "b", and store the results in "dst".
func MmAddPd(v0 M128D, v1 M128D) M128D {
	return C._mm_add_pd(v0, v1)
}

// Subtract the lower double-precision (64-bit) floating-point element in "b" from the lower double-precision (64-bit) floating-point element in "a", store the result in the lower element of "dst", and copy the upper element from "a" to the upper element of "dst".
func MmSubSd(v0 M128D, v1 M128D) M128D {
	return C._mm_sub_sd(v0, v1)
}

// Subtract packed double-precision (64-bit) floating-point elements in "b" from packed double-precision (64-bit) floating-point elements in "a", and store the results in "dst".
func MmSubPd(v0 M128D, v1 M128D) M128D {
	return C._mm_sub_pd(v0, v1)
}

// Multiply the lower double-precision (64-bit) floating-point element in "a" and "b", store the result in the lower element of "dst", and copy the upper element from "a" to the upper element of "dst".
func MmMulSd(v0 M128D, v1 M128D) M128D {
	return C._mm_mul_sd(v0, v1)
}

// Multiply packed double-precision (64-bit) floating-point elements in "a" and "b", and store the results in "dst".
func MmMulPd(v0 M128D, v1 M128D) M128D {
	return C._mm_mul_pd(v0, v1)
}

// Divide the lower double-precision (64-bit) floating-point element in "a" by the lower double-precision (64-bit) floating-point element in "b", store the result in the lower element of "dst", and copy the upper element from "a" to the upper element of "dst".
func MmDivSd(v0 M128D, v1 M128D) M128D {
	return C._mm_div_sd(v0, v1)
}

// Divide packed double-precision (64-bit) floating-point elements in "a" by packed elements in "b", and store the results in "dst".
func MmDivPd(v0 M128D, v1 M128D) M128D {
	return C._mm_div_pd(v0, v1)
}

// Compute the square root of the lower double-precision (64-bit) floating-point element in "b", store the result in the lower element of "dst", and copy the upper element from "a" to the upper element of "dst".
func MmSqrtSd(v0 M128D, v1 M128D) M128D {
	return C._mm_sqrt_sd(v0, v1)
}

// Compute the square root of packed double-precision (64-bit) floating-point elements in "a", and store the results in "dst".
func MmSqrtPd(v0 M128D) M128D {
	return C._mm_sqrt_pd(v0)
}

// Compare the lower double-precision (64-bit) floating-point elements in "a" and "b", store the minimum value in the lower element of "dst", and copy the upper element from "a" to the upper element of "dst". [min_float_note]
func MmMinSd(v0 M128D, v1 M128D) M128D {
	return C._mm_min_sd(v0, v1)
}

// Compare packed double-precision (64-bit) floating-point elements in "a" and "b", and store packed minimum values in "dst". [min_float_note]
func MmMinPd(v0 M128D, v1 M128D) M128D {
	return C._mm_min_pd(v0, v1)
}

// Compare the lower double-precision (64-bit) floating-point elements in "a" and "b", store the maximum value in the lower element of "dst", and copy the upper element from "a" to the upper element of "dst". [max_float_note]
func MmMaxSd(v0 M128D, v1 M128D) M128D {
	return C._mm_max_sd(v0, v1)
}

// Compare packed double-precision (64-bit) floating-point elements in "a" and "b", and store packed maximum values in "dst". [max_float_note]
func MmMaxPd(v0 M128D, v1 M128D) M128D {
	return C._mm_max_pd(v0, v1)
}

// Compute the bitwise AND of packed double-precision (64-bit) floating-point elements in "a" and "b", and store the results in "dst".
func MmAndPd(v0 M128D, v1 M128D) M128D {
	return C._mm_and_pd(v0, v1)
}

// Compute the bitwise NOT of packed double-precision (64-bit) floating-point elements in "a" and then AND with "b", and store the results in "dst".
func MmAndnotPd(v0 M128D, v1 M128D) M128D {
	return C._mm_andnot_pd(v0, v1)
}

// Compute the bitwise OR of packed double-precision (64-bit) floating-point elements in "a" and "b", and store the results in "dst".
func MmOrPd(v0 M128D, v1 M128D) M128D {
	return C._mm_or_pd(v0, v1)
}

// Compute the bitwise XOR of packed double-precision (64-bit) floating-point elements in "a" and "b", and store the results in "dst".
func MmXorPd(v0 M128D, v1 M128D) M128D {
	return C._mm_xor_pd(v0, v1)
}

// Compare packed double-precision (64-bit) floating-point elements in "a" and "b" for equality, and store the results in "dst".
func MmCmpeqPd(v0 M128D, v1 M128D) M128D {
	return C._mm_cmpeq_pd(v0, v1)
}

// Compare packed double-precision (64-bit) floating-point elements in "a" and "b" for less-than, and store the results in "dst".
func MmCmpltPd(v0 M128D, v1 M128D) M128D {
	return C._mm_cmplt_pd(v0, v1)
}

// Compare packed double-precision (64-bit) floating-point elements in "a" and "b" for less-than-or-equal, and store the results in "dst".
func MmCmplePd(v0 M128D, v1 M128D) M128D {
	return C._mm_cmple_pd(v0, v1)
}

// Compare packed double-precision (64-bit) floating-point elements in "a" and "b" for greater-than, and store the results in "dst".
func MmCmpgtPd(v0 M128D, v1 M128D) M128D {
	return C._mm_cmpgt_pd(v0, v1)
}

// Compare packed double-precision (64-bit) floating-point elements in "a" and "b" for greater-than-or-equal, and store the results in "dst".
func MmCmpgePd(v0 M128D, v1 M128D) M128D {
	return C._mm_cmpge_pd(v0, v1)
}

// Compare packed double-precision (64-bit) floating-point elements in "a" and "b" to see if neither is NaN, and store the results in "dst".
func MmCmpordPd(v0 M128D, v1 M128D) M128D {
	return C._mm_cmpord_pd(v0, v1)
}

// Compare packed double-precision (64-bit) floating-point elements in "a" and "b" to see if either is NaN, and store the results in "dst".
func MmCmpunordPd(v0 M128D, v1 M128D) M128D {
	return C._mm_cmpunord_pd(v0, v1)
}

// Compare packed double-precision (64-bit) floating-point elements in "a" and "b" for not-equal, and store the results in "dst".
func MmCmpneqPd(v0 M128D, v1 M128D) M128D {
	return C._mm_cmpneq_pd(v0, v1)
}

// Compare packed double-precision (64-bit) floating-point elements in "a" and "b" for not-less-than, and store the results in "dst".
func MmCmpnltPd(v0 M128D, v1 M128D) M128D {
	return C._mm_cmpnlt_pd(v0, v1)
}

// Compare packed double-precision (64-bit) floating-point elements in "a" and "b" for not-less-than-or-equal, and store the results in "dst".
func MmCmpnlePd(v0 M128D, v1 M128D) M128D {
	return C._mm_cmpnle_pd(v0, v1)
}

// Compare packed double-precision (64-bit) floating-point elements in "a" and "b" for not-greater-than, and store the results in "dst".
func MmCmpngtPd(v0 M128D, v1 M128D) M128D {
	return C._mm_cmpngt_pd(v0, v1)
}

// Compare packed double-precision (64-bit) floating-point elements in "a" and "b" for not-greater-than-or-equal, and store the results in "dst".
func MmCmpngePd(v0 M128D, v1 M128D) M128D {
	return C._mm_cmpnge_pd(v0, v1)
}

// Compare the lower double-precision (64-bit) floating-point elements in "a" and "b" for equality, store the result in the lower element of "dst", and copy the upper element from "a" to the upper element of "dst".
func MmCmpeqSd(v0 M128D, v1 M128D) M128D {
	return C._mm_cmpeq_sd(v0, v1)
}

// Compare the lower double-precision (64-bit) floating-point elements in "a" and "b" for less-than, store the result in the lower element of "dst", and copy the upper element from "a" to the upper element of "dst".
func MmCmpltSd(v0 M128D, v1 M128D) M128D {
	return C._mm_cmplt_sd(v0, v1)
}

// Compare the lower double-precision (64-bit) floating-point elements in "a" and "b" for less-than-or-equal, store the result in the lower element of "dst", and copy the upper element from "a" to the upper element of "dst".
func MmCmpleSd(v0 M128D, v1 M128D) M128D {
	return C._mm_cmple_sd(v0, v1)
}

// Compare the lower double-precision (64-bit) floating-point elements in "a" and "b" for greater-than, store the result in the lower element of "dst", and copy the upper element from "a" to the upper element of "dst".
func MmCmpgtSd(v0 M128D, v1 M128D) M128D {
	return C._mm_cmpgt_sd(v0, v1)
}

// Compare the lower double-precision (64-bit) floating-point elements in "a" and "b" for greater-than-or-equal, store the result in the lower element of "dst", and copy the upper element from "a" to the upper element of "dst".
func MmCmpgeSd(v0 M128D, v1 M128D) M128D {
	return C._mm_cmpge_sd(v0, v1)
}

// Compare the lower double-precision (64-bit) floating-point elements in "a" and "b" to see if neither is NaN, store the result in the lower element of "dst", and copy the upper element from "a" to the upper element of "dst".
func MmCmpordSd(v0 M128D, v1 M128D) M128D {
	return C._mm_cmpord_sd(v0, v1)
}

// Compare the lower double-precision (64-bit) floating-point elements in "a" and "b" to see if either is NaN, store the result in the lower element of "dst", and copy the upper element from "a" to the upper element of "dst".
func MmCmpunordSd(v0 M128D, v1 M128D) M128D {
	return C._mm_cmpunord_sd(v0, v1)
}

// Compare the lower double-precision (64-bit) floating-point elements in "a" and "b" for not-equal, store the result in the lower element of "dst", and copy the upper element from "a" to the upper element of "dst".
func MmCmpneqSd(v0 M128D, v1 M128D) M128D {
	return C._mm_cmpneq_sd(v0, v1)
}

// Compare the lower double-precision (64-bit) floating-point elements in "a" and "b" for not-less-than, store the result in the lower element of "dst", and copy the upper element from "a" to the upper element of "dst".
func MmCmpnltSd(v0 M128D, v1 M128D) M128D {
	return C._mm_cmpnlt_sd(v0, v1)
}

// Compare the lower double-precision (64-bit) floating-point elements in "a" and "b" for not-less-than-or-equal, store the result in the lower element of "dst", and copy the upper element from "a" to the upper element of "dst".
func MmCmpnleSd(v0 M128D, v1 M128D) M128D {
	return C._mm_cmpnle_sd(v0, v1)
}

// Compare the lower double-precision (64-bit) floating-point elements in "a" and "b" for not-greater-than, store the result in the lower element of "dst", and copy the upper element from "a" to the upper element of "dst".
func MmCmpngtSd(v0 M128D, v1 M128D) M128D {
	return C._mm_cmpngt_sd(v0, v1)
}

// Compare the lower double-precision (64-bit) floating-point elements in "a" and "b" for not-greater-than-or-equal, store the result in the lower element of "dst", and copy the upper element from "a" to the upper element of "dst".
func MmCmpngeSd(v0 M128D, v1 M128D) M128D {
	return C._mm_cmpnge_sd(v0, v1)
}

// Compare the lower double-precision (64-bit) floating-point element in "a" and "b" for equality, and return the boolean result (0 or 1).
func MmComieqSd(v0 M128D, v1 M128D) Int {
	return C._mm_comieq_sd(v0, v1)
}

// Compare the lower double-precision (64-bit) floating-point element in "a" and "b" for less-than, and return the boolean result (0 or 1).
func MmComiltSd(v0 M128D, v1 M128D) Int {
	return C._mm_comilt_sd(v0, v1)
}

// Compare the lower double-precision (64-bit) floating-point element in "a" and "b" for less-than-or-equal, and return the boolean result (0 or 1).
func MmComileSd(v0 M128D, v1 M128D) Int {
	return C._mm_comile_sd(v0, v1)
}

// Compare the lower double-precision (64-bit) floating-point element in "a" and "b" for greater-than, and return the boolean result (0 or 1).
func MmComigtSd(v0 M128D, v1 M128D) Int {
	return C._mm_comigt_sd(v0, v1)
}

// Compare the lower double-precision (64-bit) floating-point element in "a" and "b" for greater-than-or-equal, and return the boolean result (0 or 1).
func MmComigeSd(v0 M128D, v1 M128D) Int {
	return C._mm_comige_sd(v0, v1)
}

// Compare the lower double-precision (64-bit) floating-point element in "a" and "b" for not-equal, and return the boolean result (0 or 1).
func MmComineqSd(v0 M128D, v1 M128D) Int {
	return C._mm_comineq_sd(v0, v1)
}

// Compare the lower double-precision (64-bit) floating-point element in "a" and "b" for equality, and return the boolean result (0 or 1). This instruction will not signal an exception for QNaNs.
func MmUcomieqSd(v0 M128D, v1 M128D) Int {
	return C._mm_ucomieq_sd(v0, v1)
}

// Compare the lower double-precision (64-bit) floating-point element in "a" and "b" for less-than, and return the boolean result (0 or 1). This instruction will not signal an exception for QNaNs.
func MmUcomiltSd(v0 M128D, v1 M128D) Int {
	return C._mm_ucomilt_sd(v0, v1)
}

// Compare the lower double-precision (64-bit) floating-point element in "a" and "b" for less-than-or-equal, and return the boolean result (0 or 1). This instruction will not signal an exception for QNaNs.
func MmUcomileSd(v0 M128D, v1 M128D) Int {
	return C._mm_ucomile_sd(v0, v1)
}

// Compare the lower double-precision (64-bit) floating-point element in "a" and "b" for greater-than, and return the boolean result (0 or 1). This instruction will not signal an exception for QNaNs.
func MmUcomigtSd(v0 M128D, v1 M128D) Int {
	return C._mm_ucomigt_sd(v0, v1)
}

// Compare the lower double-precision (64-bit) floating-point element in "a" and "b" for greater-than-or-equal, and return the boolean result (0 or 1). This instruction will not signal an exception for QNaNs.
func MmUcomigeSd(v0 M128D, v1 M128D) Int {
	return C._mm_ucomige_sd(v0, v1)
}

// Compare the lower double-precision (64-bit) floating-point element in "a" and "b" for not-equal, and return the boolean result (0 or 1). This instruction will not signal an exception for QNaNs.
func MmUcomineqSd(v0 M128D, v1 M128D) Int {
	return C._mm_ucomineq_sd(v0, v1)
}

// Convert packed double-precision (64-bit) floating-point elements in "a" to packed single-precision (32-bit) floating-point elements, and store the results in "dst".
func MmCvtpdPs(v0 M128D) M128 {
	return C._mm_cvtpd_ps(v0)
}

// Convert packed single-precision (32-bit) floating-point elements in "a" to packed double-precision (64-bit) floating-point elements, and store the results in "dst".
func MmCvtpsPd(v0 M128) M128D {
	return C._mm_cvtps_pd(v0)
}

// Convert packed signed 32-bit integers in "a" to packed double-precision (64-bit) floating-point elements, and store the results in "dst".
func MmCvtepi32Pd(v0 M128I) M128D {
	return C._mm_cvtepi32_pd(v0)
}

// Convert packed double-precision (64-bit) floating-point elements in "a" to packed 32-bit integers, and store the results in "dst".
func MmCvtpdEpi32(v0 M128D) M128I {
	return C._mm_cvtpd_epi32(v0)
}

// Convert the lower double-precision (64-bit) floating-point element in "a" to a 32-bit integer, and store the result in "dst".
func MmCvtsdSi32(v0 M128D) Int {
	return C._mm_cvtsd_si32(v0)
}

// Convert the lower double-precision (64-bit) floating-point element in "b" to a single-precision (32-bit) floating-point element, store the result in the lower element of "dst", and copy the upper 3 packed elements from "a" to the upper elements of "dst".
func MmCvtsdSs(v0 M128, v1 M128D) M128 {
	return C._mm_cvtsd_ss(v0, v1)
}

// Convert the signed 32-bit integer "b" to a double-precision (64-bit) floating-point element, store the result in the lower element of "dst", and copy the upper element from "a" to the upper element of "dst".
func MmCvtsi32Sd(v0 M128D, v1 Int) M128D {
	return C._mm_cvtsi32_sd(v0, v1)
}

// Convert the lower single-precision (32-bit) floating-point element in "b" to a double-precision (64-bit) floating-point element, store the result in the lower element of "dst", and copy the upper element from "a" to the upper element of "dst".
func MmCvtssSd(v0 M128D, v1 M128) M128D {
	return C._mm_cvtss_sd(v0, v1)
}

// Convert packed double-precision (64-bit) floating-point elements in "a" to packed 32-bit integers with truncation, and store the results in "dst".
func MmCvttpdEpi32(v0 M128D) M128I {
	return C._mm_cvttpd_epi32(v0)
}

// Convert the lower double-precision (64-bit) floating-point element in "a" to a 32-bit integer with truncation, and store the result in "dst".
func MmCvttsdSi32(v0 M128D) Int {
	return C._mm_cvttsd_si32(v0)
}

// Copy the lower double-precision (64-bit) floating-point element of "a" to "dst".
func MmCvtsdF64(v0 M128D) Double {
	return C._mm_cvtsd_f64(v0)
}

// Return vector of type __m128d with undefined elements.
func MmUndefinedPd() M128D {
	return C._mm_undefined_pd()
}

// Copy double-precision (64-bit) floating-point element "a" to the lower element of "dst", and zero the upper element.
func MmSetSd(v0 Double) M128D {
	return C._mm_set_sd(v0)
}

// Broadcast double-precision (64-bit) floating-point value "a" to all elements of "dst".
func MmSet1Pd(v0 Double) M128D {
	return C._mm_set1_pd(v0)
}

// Broadcast double-precision (64-bit) floating-point value "a" to all elements of "dst".
func MmSetPd1(v0 Double) M128D {
	return C._mm_set_pd1(v0)
}

// Set packed double-precision (64-bit) floating-point elements in "dst" with the supplied values.
func MmSetPd(v0 Double, v1 Double) M128D {
	return C._mm_set_pd(v0, v1)
}

// Set packed double-precision (64-bit) floating-point elements in "dst" with the supplied values in reverse order.
func MmSetrPd(v0 Double, v1 Double) M128D {
	return C._mm_setr_pd(v0, v1)
}

// Return vector of type __m128d with all elements set to zero.
func MmSetzeroPd() M128D {
	return C._mm_setzero_pd()
}

// Move the lower double-precision (64-bit) floating-point element from "b" to the lower element of "dst", and copy the upper element from "a" to the upper element of "dst".
func MmMoveSd(v0 M128D, v1 M128D) M128D {
	return C._mm_move_sd(v0, v1)
}

// Add packed 8-bit integers in "a" and "b", and store the results in "dst".
func MmAddEpi8(v0 M128I, v1 M128I) M128I {
	return C._mm_add_epi8(v0, v1)
}

// Add packed 16-bit integers in "a" and "b", and store the results in "dst".
func MmAddEpi16(v0 M128I, v1 M128I) M128I {
	return C._mm_add_epi16(v0, v1)
}

// Add packed 32-bit integers in "a" and "b", and store the results in "dst".
func MmAddEpi32(v0 M128I, v1 M128I) M128I {
	return C._mm_add_epi32(v0, v1)
}

// Add packed 64-bit integers in "a" and "b", and store the results in "dst".
func MmAddEpi64(v0 M128I, v1 M128I) M128I {
	return C._mm_add_epi64(v0, v1)
}

// Add packed signed 8-bit integers in "a" and "b" using saturation, and store the results in "dst".
func MmAddsEpi8(v0 M128I, v1 M128I) M128I {
	return C._mm_adds_epi8(v0, v1)
}

// Add packed signed 16-bit integers in "a" and "b" using saturation, and store the results in "dst".
func MmAddsEpi16(v0 M128I, v1 M128I) M128I {
	return C._mm_adds_epi16(v0, v1)
}

// Add packed unsigned 8-bit integers in "a" and "b" using saturation, and store the results in "dst".
func MmAddsEpu8(v0 M128I, v1 M128I) M128I {
	return C._mm_adds_epu8(v0, v1)
}

// Add packed unsigned 16-bit integers in "a" and "b" using saturation, and store the results in "dst".
func MmAddsEpu16(v0 M128I, v1 M128I) M128I {
	return C._mm_adds_epu16(v0, v1)
}

// Average packed unsigned 8-bit integers in "a" and "b", and store the results in "dst".
func MmAvgEpu8(v0 M128I, v1 M128I) M128I {
	return C._mm_avg_epu8(v0, v1)
}

// Average packed unsigned 16-bit integers in "a" and "b", and store the results in "dst".
func MmAvgEpu16(v0 M128I, v1 M128I) M128I {
	return C._mm_avg_epu16(v0, v1)
}

// Multiply packed signed 16-bit integers in "a" and "b", producing intermediate signed 32-bit integers. Horizontally add adjacent pairs of intermediate 32-bit integers, and pack the results in "dst".
func MmMaddEpi16(v0 M128I, v1 M128I) M128I {
	return C._mm_madd_epi16(v0, v1)
}

// Compare packed signed 16-bit integers in "a" and "b", and store packed maximum values in "dst".
func MmMaxEpi16(v0 M128I, v1 M128I) M128I {
	return C._mm_max_epi16(v0, v1)
}

// Compare packed unsigned 8-bit integers in "a" and "b", and store packed maximum values in "dst".
func MmMaxEpu8(v0 M128I, v1 M128I) M128I {
	return C._mm_max_epu8(v0, v1)
}

// Compare packed signed 16-bit integers in "a" and "b", and store packed minimum values in "dst".
func MmMinEpi16(v0 M128I, v1 M128I) M128I {
	return C._mm_min_epi16(v0, v1)
}

// Compare packed unsigned 8-bit integers in "a" and "b", and store packed minimum values in "dst".
func MmMinEpu8(v0 M128I, v1 M128I) M128I {
	return C._mm_min_epu8(v0, v1)
}

// Multiply the packed signed 16-bit integers in "a" and "b", producing intermediate 32-bit integers, and store the high 16 bits of the intermediate integers in "dst".
func MmMulhiEpi16(v0 M128I, v1 M128I) M128I {
	return C._mm_mulhi_epi16(v0, v1)
}

// Multiply the packed unsigned 16-bit integers in "a" and "b", producing intermediate 32-bit integers, and store the high 16 bits of the intermediate integers in "dst".
func MmMulhiEpu16(v0 M128I, v1 M128I) M128I {
	return C._mm_mulhi_epu16(v0, v1)
}

// Multiply the packed 16-bit integers in "a" and "b", producing intermediate 32-bit integers, and store the low 16 bits of the intermediate integers in "dst".
func MmMulloEpi16(v0 M128I, v1 M128I) M128I {
	return C._mm_mullo_epi16(v0, v1)
}

// Multiply the low unsigned 32-bit integers from each packed 64-bit element in "a" and "b", and store the unsigned 64-bit results in "dst".
func MmMulEpu32(v0 M128I, v1 M128I) M128I {
	return C._mm_mul_epu32(v0, v1)
}

// Compute the absolute differences of packed unsigned 8-bit integers in "a" and "b", then horizontally sum each consecutive 8 differences to produce two unsigned 16-bit integers, and pack these unsigned 16-bit integers in the low 16 bits of 64-bit elements in "dst".
func MmSadEpu8(v0 M128I, v1 M128I) M128I {
	return C._mm_sad_epu8(v0, v1)
}

// Subtract packed 8-bit integers in "b" from packed 8-bit integers in "a", and store the results in "dst".
func MmSubEpi8(v0 M128I, v1 M128I) M128I {
	return C._mm_sub_epi8(v0, v1)
}

// Subtract packed 16-bit integers in "b" from packed 16-bit integers in "a", and store the results in "dst".
func MmSubEpi16(v0 M128I, v1 M128I) M128I {
	return C._mm_sub_epi16(v0, v1)
}

// Subtract packed 32-bit integers in "b" from packed 32-bit integers in "a", and store the results in "dst".
func MmSubEpi32(v0 M128I, v1 M128I) M128I {
	return C._mm_sub_epi32(v0, v1)
}

// Subtract packed 64-bit integers in "b" from packed 64-bit integers in "a", and store the results in "dst".
func MmSubEpi64(v0 M128I, v1 M128I) M128I {
	return C._mm_sub_epi64(v0, v1)
}

// Subtract packed signed 8-bit integers in "b" from packed 8-bit integers in "a" using saturation, and store the results in "dst".
func MmSubsEpi8(v0 M128I, v1 M128I) M128I {
	return C._mm_subs_epi8(v0, v1)
}

// Subtract packed signed 16-bit integers in "b" from packed 16-bit integers in "a" using saturation, and store the results in "dst".
func MmSubsEpi16(v0 M128I, v1 M128I) M128I {
	return C._mm_subs_epi16(v0, v1)
}

// Subtract packed unsigned 8-bit integers in "b" from packed unsigned 8-bit integers in "a" using saturation, and store the results in "dst".
func MmSubsEpu8(v0 M128I, v1 M128I) M128I {
	return C._mm_subs_epu8(v0, v1)
}

// Subtract packed unsigned 16-bit integers in "b" from packed unsigned 16-bit integers in "a" using saturation, and store the results in "dst".
func MmSubsEpu16(v0 M128I, v1 M128I) M128I {
	return C._mm_subs_epu16(v0, v1)
}

// Compute the bitwise AND of 128 bits (representing integer data) in "a" and "b", and store the result in "dst".
func MmAndSi128(v0 M128I, v1 M128I) M128I {
	return C._mm_and_si128(v0, v1)
}

// Compute the bitwise NOT of 128 bits (representing integer data) in "a" and then AND with "b", and store the result in "dst".
func MmAndnotSi128(v0 M128I, v1 M128I) M128I {
	return C._mm_andnot_si128(v0, v1)
}

// Compute the bitwise OR of 128 bits (representing integer data) in "a" and "b", and store the result in "dst".
func MmOrSi128(v0 M128I, v1 M128I) M128I {
	return C._mm_or_si128(v0, v1)
}

// Compute the bitwise XOR of 128 bits (representing integer data) in "a" and "b", and store the result in "dst".
func MmXorSi128(v0 M128I, v1 M128I) M128I {
	return C._mm_xor_si128(v0, v1)
}

// Shift packed 16-bit integers in "a" left by "imm8" while shifting in zeros, and store the results in "dst".
func MmSlliEpi16(v0 M128I, v1 Int) M128I {
	return C._mm_slli_epi16(v0, v1)
}

// Shift packed 16-bit integers in "a" left by "count" while shifting in zeros, and store the results in "dst".
func MmSllEpi16(v0 M128I, v1 M128I) M128I {
	return C._mm_sll_epi16(v0, v1)
}

// Shift packed 32-bit integers in "a" left by "imm8" while shifting in zeros, and store the results in "dst".
func MmSlliEpi32(v0 M128I, v1 Int) M128I {
	return C._mm_slli_epi32(v0, v1)
}

// Shift packed 32-bit integers in "a" left by "count" while shifting in zeros, and store the results in "dst".
func MmSllEpi32(v0 M128I, v1 M128I) M128I {
	return C._mm_sll_epi32(v0, v1)
}

// Shift packed 64-bit integers in "a" left by "imm8" while shifting in zeros, and store the results in "dst".
func MmSlliEpi64(v0 M128I, v1 Int) M128I {
	return C._mm_slli_epi64(v0, v1)
}

// Shift packed 64-bit integers in "a" left by "count" while shifting in zeros, and store the results in "dst".
func MmSllEpi64(v0 M128I, v1 M128I) M128I {
	return C._mm_sll_epi64(v0, v1)
}

// Shift packed 16-bit integers in "a" right by "imm8" while shifting in sign bits, and store the results in "dst".
func MmSraiEpi16(v0 M128I, v1 Int) M128I {
	return C._mm_srai_epi16(v0, v1)
}

// Shift packed 16-bit integers in "a" right by "count" while shifting in sign bits, and store the results in "dst".
func MmSraEpi16(v0 M128I, v1 M128I) M128I {
	return C._mm_sra_epi16(v0, v1)
}

// Shift packed 32-bit integers in "a" right by "imm8" while shifting in sign bits, and store the results in "dst".
func MmSraiEpi32(v0 M128I, v1 Int) M128I {
	return C._mm_srai_epi32(v0, v1)
}

// Shift packed 32-bit integers in "a" right by "count" while shifting in sign bits, and store the results in "dst".
func MmSraEpi32(v0 M128I, v1 M128I) M128I {
	return C._mm_sra_epi32(v0, v1)
}

// Shift packed 16-bit integers in "a" right by "imm8" while shifting in zeros, and store the results in "dst".
func MmSrliEpi16(v0 M128I, v1 Int) M128I {
	return C._mm_srli_epi16(v0, v1)
}

// Shift packed 16-bit integers in "a" right by "count" while shifting in zeros, and store the results in "dst".
func MmSrlEpi16(v0 M128I, v1 M128I) M128I {
	return C._mm_srl_epi16(v0, v1)
}

// Shift packed 32-bit integers in "a" right by "imm8" while shifting in zeros, and store the results in "dst".
func MmSrliEpi32(v0 M128I, v1 Int) M128I {
	return C._mm_srli_epi32(v0, v1)
}

// Shift packed 32-bit integers in "a" right by "count" while shifting in zeros, and store the results in "dst".
func MmSrlEpi32(v0 M128I, v1 M128I) M128I {
	return C._mm_srl_epi32(v0, v1)
}

// Shift packed 64-bit integers in "a" right by "imm8" while shifting in zeros, and store the results in "dst".
func MmSrliEpi64(v0 M128I, v1 Int) M128I {
	return C._mm_srli_epi64(v0, v1)
}

// Shift packed 64-bit integers in "a" right by "count" while shifting in zeros, and store the results in "dst".
func MmSrlEpi64(v0 M128I, v1 M128I) M128I {
	return C._mm_srl_epi64(v0, v1)
}

// Compare packed 8-bit integers in "a" and "b" for equality, and store the results in "dst".
func MmCmpeqEpi8(v0 M128I, v1 M128I) M128I {
	return C._mm_cmpeq_epi8(v0, v1)
}

// Compare packed 16-bit integers in "a" and "b" for equality, and store the results in "dst".
func MmCmpeqEpi16(v0 M128I, v1 M128I) M128I {
	return C._mm_cmpeq_epi16(v0, v1)
}

// Compare packed 32-bit integers in "a" and "b" for equality, and store the results in "dst".
func MmCmpeqEpi32(v0 M128I, v1 M128I) M128I {
	return C._mm_cmpeq_epi32(v0, v1)
}

// Compare packed signed 8-bit integers in "a" and "b" for greater-than, and store the results in "dst".
func MmCmpgtEpi8(v0 M128I, v1 M128I) M128I {
	return C._mm_cmpgt_epi8(v0, v1)
}

// Compare packed signed 16-bit integers in "a" and "b" for greater-than, and store the results in "dst".
func MmCmpgtEpi16(v0 M128I, v1 M128I) M128I {
	return C._mm_cmpgt_epi16(v0, v1)
}

// Compare packed signed 32-bit integers in "a" and "b" for greater-than, and store the results in "dst".
func MmCmpgtEpi32(v0 M128I, v1 M128I) M128I {
	return C._mm_cmpgt_epi32(v0, v1)
}

// Compare packed signed 8-bit integers in "a" and "b" for less-than, and store the results in "dst". Note: This intrinsic emits the pcmpgtb instruction with the order of the operands switched.
func MmCmpltEpi8(v0 M128I, v1 M128I) M128I {
	return C._mm_cmplt_epi8(v0, v1)
}

// Compare packed signed 16-bit integers in "a" and "b" for less-than, and store the results in "dst". Note: This intrinsic emits the pcmpgtw instruction with the order of the operands switched.
func MmCmpltEpi16(v0 M128I, v1 M128I) M128I {
	return C._mm_cmplt_epi16(v0, v1)
}

// Compare packed signed 32-bit integers in "a" and "b" for less-than, and store the results in "dst". Note: This intrinsic emits the pcmpgtd instruction with the order of the operands switched.
func MmCmpltEpi32(v0 M128I, v1 M128I) M128I {
	return C._mm_cmplt_epi32(v0, v1)
}

// Convert the signed 64-bit integer "b" to a double-precision (64-bit) floating-point element, store the result in the lower element of "dst", and copy the upper element from "a" to the upper element of "dst".
func MmCvtsi64Sd(v0 M128D, v1 Longlong) M128D {
	return C._mm_cvtsi64_sd(v0, v1)
}

// Convert the lower double-precision (64-bit) floating-point element in "a" to a 64-bit integer, and store the result in "dst".
func MmCvtsdSi64(v0 M128D) Longlong {
	return C._mm_cvtsd_si64(v0)
}

// Convert the lower double-precision (64-bit) floating-point element in "a" to a 64-bit integer with truncation, and store the result in "dst".
func MmCvttsdSi64(v0 M128D) Longlong {
	return C._mm_cvttsd_si64(v0)
}

// Convert packed signed 32-bit integers in "a" to packed single-precision (32-bit) floating-point elements, and store the results in "dst".
func MmCvtepi32Ps(v0 M128I) M128 {
	return C._mm_cvtepi32_ps(v0)
}

// Convert packed single-precision (32-bit) floating-point elements in "a" to packed 32-bit integers, and store the results in "dst".
func MmCvtpsEpi32(v0 M128) M128I {
	return C._mm_cvtps_epi32(v0)
}

// Convert packed single-precision (32-bit) floating-point elements in "a" to packed 32-bit integers with truncation, and store the results in "dst".
func MmCvttpsEpi32(v0 M128) M128I {
	return C._mm_cvttps_epi32(v0)
}

// Copy 32-bit integer "a" to the lower elements of "dst", and zero the upper elements of "dst".
func MmCvtsi32Si128(v0 Int) M128I {
	return C._mm_cvtsi32_si128(v0)
}

// Copy 64-bit integer "a" to the lower element of "dst", and zero the upper element.
func MmCvtsi64Si128(v0 Longlong) M128I {
	return C._mm_cvtsi64_si128(v0)
}

// Copy the lower 32-bit integer in "a" to "dst".
func MmCvtsi128Si32(v0 M128I) Int {
	return C._mm_cvtsi128_si32(v0)
}

// Copy the lower 64-bit integer in "a" to "dst".
func MmCvtsi128Si64(v0 M128I) Longlong {
	return C._mm_cvtsi128_si64(v0)
}

// Return vector of type __m128i with undefined elements.
func MmUndefinedSi128() M128I {
	return C._mm_undefined_si128()
}

// Set packed 64-bit integers in "dst" with the supplied values.
func MmSetEpi64X(v0 Longlong, v1 Longlong) M128I {
	return C._mm_set_epi64x(v0, v1)
}

// Set packed 64-bit integers in "dst" with the supplied values.
func MmSetEpi64(v0 M64, v1 M64) M128I {
	return C._mm_set_epi64(v0, v1)
}

// Set packed 32-bit integers in "dst" with the supplied values.
func MmSetEpi32(v0 Int, v1 Int, v2 Int, v3 Int) M128I {
	return C._mm_set_epi32(v0, v1, v2, v3)
}

// Set packed 16-bit integers in "dst" with the supplied values.
func MmSetEpi16(v0 Short, v1 Short, v2 Short, v3 Short, v4 Short, v5 Short, v6 Short, v7 Short) M128I {
	return C._mm_set_epi16(v0, v1, v2, v3, v4, v5, v6, v7)
}

// Set packed 8-bit integers in "dst" with the supplied values.
func MmSetEpi8(v0 Char, v1 Char, v2 Char, v3 Char, v4 Char, v5 Char, v6 Char, v7 Char, v8 Char, v9 Char, v10 Char, v11 Char, v12 Char, v13 Char, v14 Char, v15 Char) M128I {
	return C._mm_set_epi8(v0, v1, v2, v3, v4, v5, v6, v7, v8, v9, v10, v11, v12, v13, v14, v15)
}

// Broadcast 64-bit integer "a" to all elements of "dst". This intrinsic may generate the "vpbroadcastq".
func MmSet1Epi64X(v0 Longlong) M128I {
	return C._mm_set1_epi64x(v0)
}

// Broadcast 64-bit integer "a" to all elements of "dst".
func MmSet1Epi64(v0 M64) M128I {
	return C._mm_set1_epi64(v0)
}

// Broadcast 32-bit integer "a" to all elements of "dst". This intrinsic may generate "vpbroadcastd".
func MmSet1Epi32(v0 Int) M128I {
	return C._mm_set1_epi32(v0)
}

// Broadcast 16-bit integer "a" to all all elements of "dst". This intrinsic may generate "vpbroadcastw".
func MmSet1Epi16(v0 Short) M128I {
	return C._mm_set1_epi16(v0)
}

// Broadcast 8-bit integer "a" to all elements of "dst". This intrinsic may generate "vpbroadcastb".
func MmSet1Epi8(v0 Char) M128I {
	return C._mm_set1_epi8(v0)
}

// Set packed 64-bit integers in "dst" with the supplied values in reverse order.
func MmSetrEpi64(v0 M64, v1 M64) M128I {
	return C._mm_setr_epi64(v0, v1)
}

// Set packed 32-bit integers in "dst" with the supplied values in reverse order.
func MmSetrEpi32(v0 Int, v1 Int, v2 Int, v3 Int) M128I {
	return C._mm_setr_epi32(v0, v1, v2, v3)
}

// Set packed 16-bit integers in "dst" with the supplied values in reverse order.
func MmSetrEpi16(v0 Short, v1 Short, v2 Short, v3 Short, v4 Short, v5 Short, v6 Short, v7 Short) M128I {
	return C._mm_setr_epi16(v0, v1, v2, v3, v4, v5, v6, v7)
}

// Set packed 8-bit integers in "dst" with the supplied values in reverse order.
func MmSetrEpi8(v0 Char, v1 Char, v2 Char, v3 Char, v4 Char, v5 Char, v6 Char, v7 Char, v8 Char, v9 Char, v10 Char, v11 Char, v12 Char, v13 Char, v14 Char, v15 Char) M128I {
	return C._mm_setr_epi8(v0, v1, v2, v3, v4, v5, v6, v7, v8, v9, v10, v11, v12, v13, v14, v15)
}

// Return vector of type __m128i with all elements set to zero.
func MmSetzeroSi128() M128I {
	return C._mm_setzero_si128()
}

// Convert packed signed 16-bit integers from "a" and "b" to packed 8-bit integers using signed saturation, and store the results in "dst".
func MmPacksEpi16(v0 M128I, v1 M128I) M128I {
	return C._mm_packs_epi16(v0, v1)
}

// Convert packed signed 32-bit integers from "a" and "b" to packed 16-bit integers using signed saturation, and store the results in "dst".
func MmPacksEpi32(v0 M128I, v1 M128I) M128I {
	return C._mm_packs_epi32(v0, v1)
}

// Convert packed signed 16-bit integers from "a" and "b" to packed 8-bit integers using unsigned saturation, and store the results in "dst".
func MmPackusEpi16(v0 M128I, v1 M128I) M128I {
	return C._mm_packus_epi16(v0, v1)
}

// Create mask from the most significant bit of each 8-bit element in "a", and store the result in "dst".
func MmMovemaskEpi8(v0 M128I) Int {
	return C._mm_movemask_epi8(v0)
}

// Unpack and interleave 8-bit integers from the high half of "a" and "b", and store the results in "dst".
func MmUnpackhiEpi8(v0 M128I, v1 M128I) M128I {
	return C._mm_unpackhi_epi8(v0, v1)
}

// Unpack and interleave 16-bit integers from the high half of "a" and "b", and store the results in "dst".
func MmUnpackhiEpi16(v0 M128I, v1 M128I) M128I {
	return C._mm_unpackhi_epi16(v0, v1)
}

// Unpack and interleave 32-bit integers from the high half of "a" and "b", and store the results in "dst".
func MmUnpackhiEpi32(v0 M128I, v1 M128I) M128I {
	return C._mm_unpackhi_epi32(v0, v1)
}

// Unpack and interleave 64-bit integers from the high half of "a" and "b", and store the results in "dst".
func MmUnpackhiEpi64(v0 M128I, v1 M128I) M128I {
	return C._mm_unpackhi_epi64(v0, v1)
}

// Unpack and interleave 8-bit integers from the low half of "a" and "b", and store the results in "dst".
func MmUnpackloEpi8(v0 M128I, v1 M128I) M128I {
	return C._mm_unpacklo_epi8(v0, v1)
}

// Unpack and interleave 16-bit integers from the low half of "a" and "b", and store the results in "dst".
func MmUnpackloEpi16(v0 M128I, v1 M128I) M128I {
	return C._mm_unpacklo_epi16(v0, v1)
}

// Unpack and interleave 32-bit integers from the low half of "a" and "b", and store the results in "dst".
func MmUnpackloEpi32(v0 M128I, v1 M128I) M128I {
	return C._mm_unpacklo_epi32(v0, v1)
}

// Unpack and interleave 64-bit integers from the low half of "a" and "b", and store the results in "dst".
func MmUnpackloEpi64(v0 M128I, v1 M128I) M128I {
	return C._mm_unpacklo_epi64(v0, v1)
}

// Copy the lower 64-bit integer in "a" to "dst".
func MmMovepi64Pi64(v0 M128I) M64 {
	return C._mm_movepi64_pi64(v0)
}

// Copy the 64-bit integer "a" to the lower element of "dst", and zero the upper element.
func MmMovpi64Epi64(v0 M64) M128I {
	return C._mm_movpi64_epi64(v0)
}

// Copy the lower 64-bit integer in "a" to the lower element of "dst", and zero the upper element.
func MmMoveEpi64(v0 M128I) M128I {
	return C._mm_move_epi64(v0)
}

// Unpack and interleave double-precision (64-bit) floating-point elements from the high half of "a" and "b", and store the results in "dst".
func MmUnpackhiPd(v0 M128D, v1 M128D) M128D {
	return C._mm_unpackhi_pd(v0, v1)
}

// Unpack and interleave double-precision (64-bit) floating-point elements from the low half of "a" and "b", and store the results in "dst".
func MmUnpackloPd(v0 M128D, v1 M128D) M128D {
	return C._mm_unpacklo_pd(v0, v1)
}

// Set each bit of mask "dst" based on the most significant bit of the corresponding packed double-precision (64-bit) floating-point element in "a".
func MmMovemaskPd(v0 M128D) Int {
	return C._mm_movemask_pd(v0)
}

// Cast vector of type __m128d to type __m128. This intrinsic is only used for compilation and does not generate any instructions, thus it has zero latency.
func MmCastpdPs(v0 M128D) M128 {
	return C._mm_castpd_ps(v0)
}

// Cast vector of type __m128d to type __m128i. This intrinsic is only used for compilation and does not generate any instructions, thus it has zero latency.
func MmCastpdSi128(v0 M128D) M128I {
	return C._mm_castpd_si128(v0)
}

// Cast vector of type __m128 to type __m128d. This intrinsic is only used for compilation and does not generate any instructions, thus it has zero latency.
func MmCastpsPd(v0 M128) M128D {
	return C._mm_castps_pd(v0)
}

// Cast vector of type __m128 to type __m128i. This intrinsic is only used for compilation and does not generate any instructions, thus it has zero latency.
func MmCastpsSi128(v0 M128) M128I {
	return C._mm_castps_si128(v0)
}

// Cast vector of type __m128i to type __m128. This intrinsic is only used for compilation and does not generate any instructions, thus it has zero latency.
func MmCastsi128Ps(v0 M128I) M128 {
	return C._mm_castsi128_ps(v0)
}

// Cast vector of type __m128i to type __m128d. This intrinsic is only used for compilation and does not generate any instructions, thus it has zero latency.
func MmCastsi128Pd(v0 M128I) M128D {
	return C._mm_castsi128_pd(v0)
}
