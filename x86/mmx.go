package x86

/*
#cgo CFLAGS: -march=native
#include <immintrin.h>
*/
import "C"

// Empty the MMX state, which marks the x87 FPU registers as available for use by x87 instructions. This instruction must be used at the end of all MMX technology procedures.
func MmEmpty() {
	C._mm_empty()
}

// Copy 32-bit integer "a" to the lower elements of "dst", and zero the upper element of "dst".
func MmCvtsi32Si64(v0 Int) M64 {
	return C._mm_cvtsi32_si64(v0)
}

// Copy the lower 32-bit integer in "a" to "dst".
func MmCvtsi64Si32(v0 M64) Int {
	return C._mm_cvtsi64_si32(v0)
}

// Copy 64-bit integer "a" to "dst".
func MmCvtsi64M64(v0 Longlong) M64 {
	return C._mm_cvtsi64_m64(v0)
}

// Copy 64-bit integer "a" to "dst".
func MmCvtm64Si64(v0 M64) Longlong {
	return C._mm_cvtm64_si64(v0)
}

// Convert packed signed 16-bit integers from "a" and "b" to packed 8-bit integers using signed saturation, and store the results in "dst".
func MmPacksPi16(v0 M64, v1 M64) M64 {
	return C._mm_packs_pi16(v0, v1)
}

// Convert packed signed 32-bit integers from "a" and "b" to packed 16-bit integers using signed saturation, and store the results in "dst".
func MmPacksPi32(v0 M64, v1 M64) M64 {
	return C._mm_packs_pi32(v0, v1)
}

// Convert packed signed 16-bit integers from "a" and "b" to packed 8-bit integers using unsigned saturation, and store the results in "dst".
func MmPacksPu16(v0 M64, v1 M64) M64 {
	return C._mm_packs_pu16(v0, v1)
}

// Unpack and interleave 8-bit integers from the high half of "a" and "b", and store the results in "dst".
func MmUnpackhiPi8(v0 M64, v1 M64) M64 {
	return C._mm_unpackhi_pi8(v0, v1)
}

// Unpack and interleave 16-bit integers from the high half of "a" and "b", and store the results in "dst".
func MmUnpackhiPi16(v0 M64, v1 M64) M64 {
	return C._mm_unpackhi_pi16(v0, v1)
}

// Unpack and interleave 32-bit integers from the high half of "a" and "b", and store the results in "dst".
func MmUnpackhiPi32(v0 M64, v1 M64) M64 {
	return C._mm_unpackhi_pi32(v0, v1)
}

// Unpack and interleave 8-bit integers from the low half of "a" and "b", and store the results in "dst".
func MmUnpackloPi8(v0 M64, v1 M64) M64 {
	return C._mm_unpacklo_pi8(v0, v1)
}

// Unpack and interleave 16-bit integers from the low half of "a" and "b", and store the results in "dst".
func MmUnpackloPi16(v0 M64, v1 M64) M64 {
	return C._mm_unpacklo_pi16(v0, v1)
}

// Unpack and interleave 32-bit integers from the low half of "a" and "b", and store the results in "dst".
func MmUnpackloPi32(v0 M64, v1 M64) M64 {
	return C._mm_unpacklo_pi32(v0, v1)
}

// Add packed 8-bit integers in "a" and "b", and store the results in "dst".
func MmAddPi8(v0 M64, v1 M64) M64 {
	return C._mm_add_pi8(v0, v1)
}

// Add packed 16-bit integers in "a" and "b", and store the results in "dst".
func MmAddPi16(v0 M64, v1 M64) M64 {
	return C._mm_add_pi16(v0, v1)
}

// Add packed 32-bit integers in "a" and "b", and store the results in "dst".
func MmAddPi32(v0 M64, v1 M64) M64 {
	return C._mm_add_pi32(v0, v1)
}

// Add packed signed 8-bit integers in "a" and "b" using saturation, and store the results in "dst".
func MmAddsPi8(v0 M64, v1 M64) M64 {
	return C._mm_adds_pi8(v0, v1)
}

// Add packed signed 16-bit integers in "a" and "b" using saturation, and store the results in "dst".
func MmAddsPi16(v0 M64, v1 M64) M64 {
	return C._mm_adds_pi16(v0, v1)
}

// Add packed unsigned 8-bit integers in "a" and "b" using saturation, and store the results in "dst".
func MmAddsPu8(v0 M64, v1 M64) M64 {
	return C._mm_adds_pu8(v0, v1)
}

// Add packed unsigned 16-bit integers in "a" and "b" using saturation, and store the results in "dst".
func MmAddsPu16(v0 M64, v1 M64) M64 {
	return C._mm_adds_pu16(v0, v1)
}

// Subtract packed 8-bit integers in "b" from packed 8-bit integers in "a", and store the results in "dst".
func MmSubPi8(v0 M64, v1 M64) M64 {
	return C._mm_sub_pi8(v0, v1)
}

// Subtract packed 16-bit integers in "b" from packed 16-bit integers in "a", and store the results in "dst".
func MmSubPi16(v0 M64, v1 M64) M64 {
	return C._mm_sub_pi16(v0, v1)
}

// Subtract packed 32-bit integers in "b" from packed 32-bit integers in "a", and store the results in "dst".
func MmSubPi32(v0 M64, v1 M64) M64 {
	return C._mm_sub_pi32(v0, v1)
}

// Subtract packed signed 8-bit integers in "b" from packed 8-bit integers in "a" using saturation, and store the results in "dst".
func MmSubsPi8(v0 M64, v1 M64) M64 {
	return C._mm_subs_pi8(v0, v1)
}

// Subtract packed signed 16-bit integers in "b" from packed 16-bit integers in "a" using saturation, and store the results in "dst".
func MmSubsPi16(v0 M64, v1 M64) M64 {
	return C._mm_subs_pi16(v0, v1)
}

// Subtract packed unsigned 8-bit integers in "b" from packed unsigned 8-bit integers in "a" using saturation, and store the results in "dst".
func MmSubsPu8(v0 M64, v1 M64) M64 {
	return C._mm_subs_pu8(v0, v1)
}

// Subtract packed unsigned 16-bit integers in "b" from packed unsigned 16-bit integers in "a" using saturation, and store the results in "dst".
func MmSubsPu16(v0 M64, v1 M64) M64 {
	return C._mm_subs_pu16(v0, v1)
}

// Multiply packed signed 16-bit integers in "a" and "b", producing intermediate signed 32-bit integers. Horizontally add adjacent pairs of intermediate 32-bit integers, and pack the results in "dst".
func MmMaddPi16(v0 M64, v1 M64) M64 {
	return C._mm_madd_pi16(v0, v1)
}

// Multiply the packed signed 16-bit integers in "a" and "b", producing intermediate 32-bit integers, and store the high 16 bits of the intermediate integers in "dst".
func MmMulhiPi16(v0 M64, v1 M64) M64 {
	return C._mm_mulhi_pi16(v0, v1)
}

// Multiply the packed 16-bit integers in "a" and "b", producing intermediate 32-bit integers, and store the low 16 bits of the intermediate integers in "dst".
func MmMulloPi16(v0 M64, v1 M64) M64 {
	return C._mm_mullo_pi16(v0, v1)
}

// Shift packed 16-bit integers in "a" left by "count" while shifting in zeros, and store the results in "dst".
func MmSllPi16(v0 M64, v1 M64) M64 {
	return C._mm_sll_pi16(v0, v1)
}

// Shift packed 16-bit integers in "a" left by "imm8" while shifting in zeros, and store the results in "dst".
func MmSlliPi16(v0 M64, v1 Int) M64 {
	return C._mm_slli_pi16(v0, v1)
}

// Shift packed 32-bit integers in "a" left by "count" while shifting in zeros, and store the results in "dst".
func MmSllPi32(v0 M64, v1 M64) M64 {
	return C._mm_sll_pi32(v0, v1)
}

// Shift packed 32-bit integers in "a" left by "imm8" while shifting in zeros, and store the results in "dst".
func MmSlliPi32(v0 M64, v1 Int) M64 {
	return C._mm_slli_pi32(v0, v1)
}

// Shift 64-bit integer "a" left by "count" while shifting in zeros, and store the result in "dst".
func MmSllSi64(v0 M64, v1 M64) M64 {
	return C._mm_sll_si64(v0, v1)
}

// Shift 64-bit integer "a" left by "imm8" while shifting in zeros, and store the result in "dst".
func MmSlliSi64(v0 M64, v1 Int) M64 {
	return C._mm_slli_si64(v0, v1)
}

// Shift packed 16-bit integers in "a" right by "count" while shifting in sign bits, and store the results in "dst".
func MmSraPi16(v0 M64, v1 M64) M64 {
	return C._mm_sra_pi16(v0, v1)
}

// Shift packed 16-bit integers in "a" right by "imm8" while shifting in sign bits, and store the results in "dst".
func MmSraiPi16(v0 M64, v1 Int) M64 {
	return C._mm_srai_pi16(v0, v1)
}

// Shift packed 32-bit integers in "a" right by "count" while shifting in sign bits, and store the results in "dst".
func MmSraPi32(v0 M64, v1 M64) M64 {
	return C._mm_sra_pi32(v0, v1)
}

// Shift packed 32-bit integers in "a" right by "imm8" while shifting in sign bits, and store the results in "dst".
func MmSraiPi32(v0 M64, v1 Int) M64 {
	return C._mm_srai_pi32(v0, v1)
}

// Shift packed 16-bit integers in "a" right by "count" while shifting in zeros, and store the results in "dst".
func MmSrlPi16(v0 M64, v1 M64) M64 {
	return C._mm_srl_pi16(v0, v1)
}

// Shift packed 16-bit integers in "a" right by "imm8" while shifting in zeros, and store the results in "dst".
func MmSrliPi16(v0 M64, v1 Int) M64 {
	return C._mm_srli_pi16(v0, v1)
}

// Shift packed 32-bit integers in "a" right by "count" while shifting in zeros, and store the results in "dst".
func MmSrlPi32(v0 M64, v1 M64) M64 {
	return C._mm_srl_pi32(v0, v1)
}

// Shift packed 32-bit integers in "a" right by "imm8" while shifting in zeros, and store the results in "dst".
func MmSrliPi32(v0 M64, v1 Int) M64 {
	return C._mm_srli_pi32(v0, v1)
}

// Shift 64-bit integer "a" right by "count" while shifting in zeros, and store the result in "dst".
func MmSrlSi64(v0 M64, v1 M64) M64 {
	return C._mm_srl_si64(v0, v1)
}

// Shift 64-bit integer "a" right by "imm8" while shifting in zeros, and store the result in "dst".
func MmSrliSi64(v0 M64, v1 Int) M64 {
	return C._mm_srli_si64(v0, v1)
}

// Compute the bitwise AND of 64 bits (representing integer data) in "a" and "b", and store the result in "dst".
func MmAndSi64(v0 M64, v1 M64) M64 {
	return C._mm_and_si64(v0, v1)
}

// Compute the bitwise NOT of 64 bits (representing integer data) in "a" and then AND with "b", and store the result in "dst".
func MmAndnotSi64(v0 M64, v1 M64) M64 {
	return C._mm_andnot_si64(v0, v1)
}

// Compute the bitwise OR of 64 bits (representing integer data) in "a" and "b", and store the result in "dst".
func MmOrSi64(v0 M64, v1 M64) M64 {
	return C._mm_or_si64(v0, v1)
}

// Compute the bitwise XOR of 64 bits (representing integer data) in "a" and "b", and store the result in "dst".
func MmXorSi64(v0 M64, v1 M64) M64 {
	return C._mm_xor_si64(v0, v1)
}

// Compare packed 8-bit integers in "a" and "b" for equality, and store the results in "dst".
func MmCmpeqPi8(v0 M64, v1 M64) M64 {
	return C._mm_cmpeq_pi8(v0, v1)
}

// Compare packed 16-bit integers in "a" and "b" for equality, and store the results in "dst".
func MmCmpeqPi16(v0 M64, v1 M64) M64 {
	return C._mm_cmpeq_pi16(v0, v1)
}

// Compare packed 32-bit integers in "a" and "b" for equality, and store the results in "dst".
func MmCmpeqPi32(v0 M64, v1 M64) M64 {
	return C._mm_cmpeq_pi32(v0, v1)
}

// Compare packed signed 8-bit integers in "a" and "b" for greater-than, and store the results in "dst".
func MmCmpgtPi8(v0 M64, v1 M64) M64 {
	return C._mm_cmpgt_pi8(v0, v1)
}

// Compare packed signed 16-bit integers in "a" and "b" for greater-than, and store the results in "dst".
func MmCmpgtPi16(v0 M64, v1 M64) M64 {
	return C._mm_cmpgt_pi16(v0, v1)
}

// Compare packed signed 32-bit integers in "a" and "b" for greater-than, and store the results in "dst".
func MmCmpgtPi32(v0 M64, v1 M64) M64 {
	return C._mm_cmpgt_pi32(v0, v1)
}

// Return vector of type __m64 with all elements set to zero.
func MmSetzeroSi64() M64 {
	return C._mm_setzero_si64()
}

// Set packed 32-bit integers in "dst" with the supplied values.
func MmSetPi32(v0 Int, v1 Int) M64 {
	return C._mm_set_pi32(v0, v1)
}

// Set packed 16-bit integers in "dst" with the supplied values.
func MmSetPi16(v0 Short, v1 Short, v2 Short, v3 Short) M64 {
	return C._mm_set_pi16(v0, v1, v2, v3)
}

// Set packed 8-bit integers in "dst" with the supplied values.
func MmSetPi8(v0 Char, v1 Char, v2 Char, v3 Char, v4 Char, v5 Char, v6 Char, v7 Char) M64 {
	return C._mm_set_pi8(v0, v1, v2, v3, v4, v5, v6, v7)
}

// Broadcast 32-bit integer "a" to all elements of "dst".
func MmSet1Pi32(v0 Int) M64 {
	return C._mm_set1_pi32(v0)
}

// Broadcast 16-bit integer "a" to all all elements of "dst".
func MmSet1Pi16(v0 Short) M64 {
	return C._mm_set1_pi16(v0)
}

// Broadcast 8-bit integer "a" to all elements of "dst".
func MmSet1Pi8(v0 Char) M64 {
	return C._mm_set1_pi8(v0)
}

// Set packed 32-bit integers in "dst" with the supplied values in reverse order.
func MmSetrPi32(v0 Int, v1 Int) M64 {
	return C._mm_setr_pi32(v0, v1)
}

// Set packed 16-bit integers in "dst" with the supplied values in reverse order.
func MmSetrPi16(v0 Short, v1 Short, v2 Short, v3 Short) M64 {
	return C._mm_setr_pi16(v0, v1, v2, v3)
}

// Set packed 8-bit integers in "dst" with the supplied values in reverse order.
func MmSetrPi8(v0 Char, v1 Char, v2 Char, v3 Char, v4 Char, v5 Char, v6 Char, v7 Char) M64 {
	return C._mm_setr_pi8(v0, v1, v2, v3, v4, v5, v6, v7)
}
