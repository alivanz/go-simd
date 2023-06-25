package x86

/*
#cgo CFLAGS: -march=native
#include <immintrin.h>
*/
import "C"

// Compute the absolute value of packed signed 8-bit integers in "a", and store the unsigned results in "dst".
func Mm256AbsEpi8(v0 M256I) M256I {
	return C._mm256_abs_epi8(v0)
}

// Compute the absolute value of packed signed 16-bit integers in "a", and store the unsigned results in "dst".
func Mm256AbsEpi16(v0 M256I) M256I {
	return C._mm256_abs_epi16(v0)
}

// Compute the absolute value of packed signed 32-bit integers in "a", and store the unsigned results in "dst".
func Mm256AbsEpi32(v0 M256I) M256I {
	return C._mm256_abs_epi32(v0)
}

// Convert packed signed 16-bit integers from "a" and "b" to packed 8-bit integers using signed saturation, and store the results in "dst".
func Mm256PacksEpi16(v0 M256I, v1 M256I) M256I {
	return C._mm256_packs_epi16(v0, v1)
}

// Convert packed signed 32-bit integers from "a" and "b" to packed 16-bit integers using signed saturation, and store the results in "dst".
func Mm256PacksEpi32(v0 M256I, v1 M256I) M256I {
	return C._mm256_packs_epi32(v0, v1)
}

// Convert packed signed 16-bit integers from "a" and "b" to packed 8-bit integers using unsigned saturation, and store the results in "dst".
func Mm256PackusEpi16(v0 M256I, v1 M256I) M256I {
	return C._mm256_packus_epi16(v0, v1)
}

// Convert packed signed 32-bit integers from "a" and "b" to packed 16-bit integers using unsigned saturation, and store the results in "dst".
func Mm256PackusEpi32(v0 M256I, v1 M256I) M256I {
	return C._mm256_packus_epi32(v0, v1)
}

// Add packed 8-bit integers in "a" and "b", and store the results in "dst".
func Mm256AddEpi8(v0 M256I, v1 M256I) M256I {
	return C._mm256_add_epi8(v0, v1)
}

// Add packed 16-bit integers in "a" and "b", and store the results in "dst".
func Mm256AddEpi16(v0 M256I, v1 M256I) M256I {
	return C._mm256_add_epi16(v0, v1)
}

// Add packed 32-bit integers in "a" and "b", and store the results in "dst".
func Mm256AddEpi32(v0 M256I, v1 M256I) M256I {
	return C._mm256_add_epi32(v0, v1)
}

// Add packed 64-bit integers in "a" and "b", and store the results in "dst".
func Mm256AddEpi64(v0 M256I, v1 M256I) M256I {
	return C._mm256_add_epi64(v0, v1)
}

// Add packed 8-bit integers in "a" and "b" using saturation, and store the results in "dst".
func Mm256AddsEpi8(v0 M256I, v1 M256I) M256I {
	return C._mm256_adds_epi8(v0, v1)
}

// Add packed 16-bit integers in "a" and "b" using saturation, and store the results in "dst".
func Mm256AddsEpi16(v0 M256I, v1 M256I) M256I {
	return C._mm256_adds_epi16(v0, v1)
}

// Add packed unsigned 8-bit integers in "a" and "b" using saturation, and store the results in "dst".
func Mm256AddsEpu8(v0 M256I, v1 M256I) M256I {
	return C._mm256_adds_epu8(v0, v1)
}

// Add packed unsigned 16-bit integers in "a" and "b" using saturation, and store the results in "dst".
func Mm256AddsEpu16(v0 M256I, v1 M256I) M256I {
	return C._mm256_adds_epu16(v0, v1)
}

// Compute the bitwise AND of 256 bits (representing integer data) in "a" and "b", and store the result in "dst".
func Mm256AndSi256(v0 M256I, v1 M256I) M256I {
	return C._mm256_and_si256(v0, v1)
}

// Compute the bitwise NOT of 256 bits (representing integer data) in "a" and then AND with "b", and store the result in "dst".
func Mm256AndnotSi256(v0 M256I, v1 M256I) M256I {
	return C._mm256_andnot_si256(v0, v1)
}

// Average packed unsigned 8-bit integers in "a" and "b", and store the results in "dst".
func Mm256AvgEpu8(v0 M256I, v1 M256I) M256I {
	return C._mm256_avg_epu8(v0, v1)
}

// Average packed unsigned 16-bit integers in "a" and "b", and store the results in "dst".
func Mm256AvgEpu16(v0 M256I, v1 M256I) M256I {
	return C._mm256_avg_epu16(v0, v1)
}

// Blend packed 8-bit integers from "a" and "b" using "mask", and store the results in "dst".
func Mm256BlendvEpi8(v0 M256I, v1 M256I, v2 M256I) M256I {
	return C._mm256_blendv_epi8(v0, v1, v2)
}

// Compare packed 8-bit integers in "a" and "b" for equality, and store the results in "dst".
func Mm256CmpeqEpi8(v0 M256I, v1 M256I) M256I {
	return C._mm256_cmpeq_epi8(v0, v1)
}

// Compare packed 16-bit integers in "a" and "b" for equality, and store the results in "dst".
func Mm256CmpeqEpi16(v0 M256I, v1 M256I) M256I {
	return C._mm256_cmpeq_epi16(v0, v1)
}

// Compare packed 32-bit integers in "a" and "b" for equality, and store the results in "dst".
func Mm256CmpeqEpi32(v0 M256I, v1 M256I) M256I {
	return C._mm256_cmpeq_epi32(v0, v1)
}

// Compare packed 64-bit integers in "a" and "b" for equality, and store the results in "dst".
func Mm256CmpeqEpi64(v0 M256I, v1 M256I) M256I {
	return C._mm256_cmpeq_epi64(v0, v1)
}

// Compare packed signed 8-bit integers in "a" and "b" for greater-than, and store the results in "dst".
func Mm256CmpgtEpi8(v0 M256I, v1 M256I) M256I {
	return C._mm256_cmpgt_epi8(v0, v1)
}

// Compare packed signed 16-bit integers in "a" and "b" for greater-than, and store the results in "dst".
func Mm256CmpgtEpi16(v0 M256I, v1 M256I) M256I {
	return C._mm256_cmpgt_epi16(v0, v1)
}

// Compare packed signed 32-bit integers in "a" and "b" for greater-than, and store the results in "dst".
func Mm256CmpgtEpi32(v0 M256I, v1 M256I) M256I {
	return C._mm256_cmpgt_epi32(v0, v1)
}

// Compare packed signed 64-bit integers in "a" and "b" for greater-than, and store the results in "dst".
func Mm256CmpgtEpi64(v0 M256I, v1 M256I) M256I {
	return C._mm256_cmpgt_epi64(v0, v1)
}

// Horizontally add adjacent pairs of 16-bit integers in "a" and "b", and pack the signed 16-bit results in "dst".
func Mm256HaddEpi16(v0 M256I, v1 M256I) M256I {
	return C._mm256_hadd_epi16(v0, v1)
}

// Horizontally add adjacent pairs of 32-bit integers in "a" and "b", and pack the signed 32-bit results in "dst".
func Mm256HaddEpi32(v0 M256I, v1 M256I) M256I {
	return C._mm256_hadd_epi32(v0, v1)
}

// Horizontally add adjacent pairs of signed 16-bit integers in "a" and "b" using saturation, and pack the signed 16-bit results in "dst".
func Mm256HaddsEpi16(v0 M256I, v1 M256I) M256I {
	return C._mm256_hadds_epi16(v0, v1)
}

// Horizontally subtract adjacent pairs of 16-bit integers in "a" and "b", and pack the signed 16-bit results in "dst".
func Mm256HsubEpi16(v0 M256I, v1 M256I) M256I {
	return C._mm256_hsub_epi16(v0, v1)
}

// Horizontally subtract adjacent pairs of 32-bit integers in "a" and "b", and pack the signed 32-bit results in "dst".
func Mm256HsubEpi32(v0 M256I, v1 M256I) M256I {
	return C._mm256_hsub_epi32(v0, v1)
}

// Horizontally subtract adjacent pairs of signed 16-bit integers in "a" and "b" using saturation, and pack the signed 16-bit results in "dst".
func Mm256HsubsEpi16(v0 M256I, v1 M256I) M256I {
	return C._mm256_hsubs_epi16(v0, v1)
}

// Vertically multiply each unsigned 8-bit integer from "a" with the corresponding signed 8-bit integer from "b", producing intermediate signed 16-bit integers. Horizontally add adjacent pairs of intermediate signed 16-bit integers, and pack the saturated results in "dst".
func Mm256MaddubsEpi16(v0 M256I, v1 M256I) M256I {
	return C._mm256_maddubs_epi16(v0, v1)
}

// Multiply packed signed 16-bit integers in "a" and "b", producing intermediate signed 32-bit integers. Horizontally add adjacent pairs of intermediate 32-bit integers, and pack the results in "dst".
func Mm256MaddEpi16(v0 M256I, v1 M256I) M256I {
	return C._mm256_madd_epi16(v0, v1)
}

// Compare packed signed 8-bit integers in "a" and "b", and store packed maximum values in "dst".
func Mm256MaxEpi8(v0 M256I, v1 M256I) M256I {
	return C._mm256_max_epi8(v0, v1)
}

// Compare packed signed 16-bit integers in "a" and "b", and store packed maximum values in "dst".
func Mm256MaxEpi16(v0 M256I, v1 M256I) M256I {
	return C._mm256_max_epi16(v0, v1)
}

// Compare packed signed 32-bit integers in "a" and "b", and store packed maximum values in "dst".
func Mm256MaxEpi32(v0 M256I, v1 M256I) M256I {
	return C._mm256_max_epi32(v0, v1)
}

// Compare packed unsigned 8-bit integers in "a" and "b", and store packed maximum values in "dst".
func Mm256MaxEpu8(v0 M256I, v1 M256I) M256I {
	return C._mm256_max_epu8(v0, v1)
}

// Compare packed unsigned 16-bit integers in "a" and "b", and store packed maximum values in "dst".
func Mm256MaxEpu16(v0 M256I, v1 M256I) M256I {
	return C._mm256_max_epu16(v0, v1)
}

// Compare packed unsigned 32-bit integers in "a" and "b", and store packed maximum values in "dst".
func Mm256MaxEpu32(v0 M256I, v1 M256I) M256I {
	return C._mm256_max_epu32(v0, v1)
}

// Compare packed signed 8-bit integers in "a" and "b", and store packed minimum values in "dst".
func Mm256MinEpi8(v0 M256I, v1 M256I) M256I {
	return C._mm256_min_epi8(v0, v1)
}

// Compare packed signed 16-bit integers in "a" and "b", and store packed minimum values in "dst".
func Mm256MinEpi16(v0 M256I, v1 M256I) M256I {
	return C._mm256_min_epi16(v0, v1)
}

// Compare packed signed 32-bit integers in "a" and "b", and store packed minimum values in "dst".
func Mm256MinEpi32(v0 M256I, v1 M256I) M256I {
	return C._mm256_min_epi32(v0, v1)
}

// Compare packed unsigned 8-bit integers in "a" and "b", and store packed minimum values in "dst".
func Mm256MinEpu8(v0 M256I, v1 M256I) M256I {
	return C._mm256_min_epu8(v0, v1)
}

// Compare packed unsigned 16-bit integers in "a" and "b", and store packed minimum values in "dst".
func Mm256MinEpu16(v0 M256I, v1 M256I) M256I {
	return C._mm256_min_epu16(v0, v1)
}

// Compare packed unsigned 32-bit integers in "a" and "b", and store packed minimum values in "dst".
func Mm256MinEpu32(v0 M256I, v1 M256I) M256I {
	return C._mm256_min_epu32(v0, v1)
}

// Create mask from the most significant bit of each 8-bit element in "a", and store the result in "dst".
func Mm256MovemaskEpi8(v0 M256I) Int {
	return C._mm256_movemask_epi8(v0)
}

// Sign extend packed 8-bit integers in "a" to packed 16-bit integers, and store the results in "dst".
func Mm256Cvtepi8Epi16(v0 M128I) M256I {
	return C._mm256_cvtepi8_epi16(v0)
}

// Sign extend packed 8-bit integers in "a" to packed 32-bit integers, and store the results in "dst".
func Mm256Cvtepi8Epi32(v0 M128I) M256I {
	return C._mm256_cvtepi8_epi32(v0)
}

// Sign extend packed 8-bit integers in the low 8 bytes of "a" to packed 64-bit integers, and store the results in "dst".
func Mm256Cvtepi8Epi64(v0 M128I) M256I {
	return C._mm256_cvtepi8_epi64(v0)
}

// Sign extend packed 16-bit integers in "a" to packed 32-bit integers, and store the results in "dst".
func Mm256Cvtepi16Epi32(v0 M128I) M256I {
	return C._mm256_cvtepi16_epi32(v0)
}

// Sign extend packed 16-bit integers in "a" to packed 64-bit integers, and store the results in "dst".
func Mm256Cvtepi16Epi64(v0 M128I) M256I {
	return C._mm256_cvtepi16_epi64(v0)
}

// Sign extend packed 32-bit integers in "a" to packed 64-bit integers, and store the results in "dst".
func Mm256Cvtepi32Epi64(v0 M128I) M256I {
	return C._mm256_cvtepi32_epi64(v0)
}

// Zero extend packed unsigned 8-bit integers in "a" to packed 16-bit integers, and store the results in "dst".
func Mm256Cvtepu8Epi16(v0 M128I) M256I {
	return C._mm256_cvtepu8_epi16(v0)
}

// Zero extend packed unsigned 8-bit integers in "a" to packed 32-bit integers, and store the results in "dst".
func Mm256Cvtepu8Epi32(v0 M128I) M256I {
	return C._mm256_cvtepu8_epi32(v0)
}

// Zero extend packed unsigned 8-bit integers in the low 8 byte sof "a" to packed 64-bit integers, and store the results in "dst".
func Mm256Cvtepu8Epi64(v0 M128I) M256I {
	return C._mm256_cvtepu8_epi64(v0)
}

// Zero extend packed unsigned 16-bit integers in "a" to packed 32-bit integers, and store the results in "dst".
func Mm256Cvtepu16Epi32(v0 M128I) M256I {
	return C._mm256_cvtepu16_epi32(v0)
}

// Zero extend packed unsigned 16-bit integers in "a" to packed 64-bit integers, and store the results in "dst".
func Mm256Cvtepu16Epi64(v0 M128I) M256I {
	return C._mm256_cvtepu16_epi64(v0)
}

// Zero extend packed unsigned 32-bit integers in "a" to packed 64-bit integers, and store the results in "dst".
func Mm256Cvtepu32Epi64(v0 M128I) M256I {
	return C._mm256_cvtepu32_epi64(v0)
}

// Multiply the low signed 32-bit integers from each packed 64-bit element in "a" and "b", and store the signed 64-bit results in "dst".
func Mm256MulEpi32(v0 M256I, v1 M256I) M256I {
	return C._mm256_mul_epi32(v0, v1)
}

// Multiply packed signed 16-bit integers in "a" and "b", producing intermediate signed 32-bit integers. Truncate each intermediate integer to the 18 most significant bits, round by adding 1, and store bits [16:1] to "dst".
func Mm256MulhrsEpi16(v0 M256I, v1 M256I) M256I {
	return C._mm256_mulhrs_epi16(v0, v1)
}

// Multiply the packed unsigned 16-bit integers in "a" and "b", producing intermediate 32-bit integers, and store the high 16 bits of the intermediate integers in "dst".
func Mm256MulhiEpu16(v0 M256I, v1 M256I) M256I {
	return C._mm256_mulhi_epu16(v0, v1)
}

// Multiply the packed signed 16-bit integers in "a" and "b", producing intermediate 32-bit integers, and store the high 16 bits of the intermediate integers in "dst".
func Mm256MulhiEpi16(v0 M256I, v1 M256I) M256I {
	return C._mm256_mulhi_epi16(v0, v1)
}

// Multiply the packed signed 16-bit integers in "a" and "b", producing intermediate 32-bit integers, and store the low 16 bits of the intermediate integers in "dst".
func Mm256MulloEpi16(v0 M256I, v1 M256I) M256I {
	return C._mm256_mullo_epi16(v0, v1)
}

// Multiply the packed signed 32-bit integers in "a" and "b", producing intermediate 64-bit integers, and store the low 32 bits of the intermediate integers in "dst".
func Mm256MulloEpi32(v0 M256I, v1 M256I) M256I {
	return C._mm256_mullo_epi32(v0, v1)
}

// Multiply the low unsigned 32-bit integers from each packed 64-bit element in "a" and "b", and store the unsigned 64-bit results in "dst".
func Mm256MulEpu32(v0 M256I, v1 M256I) M256I {
	return C._mm256_mul_epu32(v0, v1)
}

// Compute the bitwise OR of 256 bits (representing integer data) in "a" and "b", and store the result in "dst".
func Mm256OrSi256(v0 M256I, v1 M256I) M256I {
	return C._mm256_or_si256(v0, v1)
}

// Compute the absolute differences of packed unsigned 8-bit integers in "a" and "b", then horizontally sum each consecutive 8 differences to produce four unsigned 16-bit integers, and pack these unsigned 16-bit integers in the low 16 bits of 64-bit elements in "dst".
func Mm256SadEpu8(v0 M256I, v1 M256I) M256I {
	return C._mm256_sad_epu8(v0, v1)
}

// Shuffle 8-bit integers in "a" within 128-bit lanes according to shuffle control mask in the corresponding 8-bit element of "b", and store the results in "dst".
func Mm256ShuffleEpi8(v0 M256I, v1 M256I) M256I {
	return C._mm256_shuffle_epi8(v0, v1)
}

// Negate packed signed 8-bit integers in "a" when the corresponding signed 8-bit integer in "b" is negative, and store the results in "dst". Element in "dst" are zeroed out when the corresponding element in "b" is zero.
func Mm256SignEpi8(v0 M256I, v1 M256I) M256I {
	return C._mm256_sign_epi8(v0, v1)
}

// Negate packed signed 16-bit integers in "a" when the corresponding signed 16-bit integer in "b" is negative, and store the results in "dst". Element in "dst" are zeroed out when the corresponding element in "b" is zero.
func Mm256SignEpi16(v0 M256I, v1 M256I) M256I {
	return C._mm256_sign_epi16(v0, v1)
}

// Negate packed signed 32-bit integers in "a" when the corresponding signed 32-bit integer in "b" is negative, and store the results in "dst". Element in "dst" are zeroed out when the corresponding element in "b" is zero.
func Mm256SignEpi32(v0 M256I, v1 M256I) M256I {
	return C._mm256_sign_epi32(v0, v1)
}

// Shift packed 16-bit integers in "a" left by "imm8" while shifting in zeros, and store the results in "dst".
func Mm256SlliEpi16(v0 M256I, v1 Int) M256I {
	return C._mm256_slli_epi16(v0, v1)
}

// Shift packed 16-bit integers in "a" left by "count" while shifting in zeros, and store the results in "dst".
func Mm256SllEpi16(v0 M256I, v1 M128I) M256I {
	return C._mm256_sll_epi16(v0, v1)
}

// Shift packed 32-bit integers in "a" left by "imm8" while shifting in zeros, and store the results in "dst".
func Mm256SlliEpi32(v0 M256I, v1 Int) M256I {
	return C._mm256_slli_epi32(v0, v1)
}

// Shift packed 32-bit integers in "a" left by "count" while shifting in zeros, and store the results in "dst".
func Mm256SllEpi32(v0 M256I, v1 M128I) M256I {
	return C._mm256_sll_epi32(v0, v1)
}

// Shift packed 64-bit integers in "a" left by "imm8" while shifting in zeros, and store the results in "dst".
func Mm256SlliEpi64(v0 M256I, v1 Int) M256I {
	return C._mm256_slli_epi64(v0, v1)
}

// Shift packed 64-bit integers in "a" left by "count" while shifting in zeros, and store the results in "dst".
func Mm256SllEpi64(v0 M256I, v1 M128I) M256I {
	return C._mm256_sll_epi64(v0, v1)
}

// Shift packed 16-bit integers in "a" right by "imm8" while shifting in sign bits, and store the results in "dst".
func Mm256SraiEpi16(v0 M256I, v1 Int) M256I {
	return C._mm256_srai_epi16(v0, v1)
}

// Shift packed 16-bit integers in "a" right by "count" while shifting in sign bits, and store the results in "dst".
func Mm256SraEpi16(v0 M256I, v1 M128I) M256I {
	return C._mm256_sra_epi16(v0, v1)
}

// Shift packed 32-bit integers in "a" right by "imm8" while shifting in sign bits, and store the results in "dst".
func Mm256SraiEpi32(v0 M256I, v1 Int) M256I {
	return C._mm256_srai_epi32(v0, v1)
}

// Shift packed 32-bit integers in "a" right by "count" while shifting in sign bits, and store the results in "dst".
func Mm256SraEpi32(v0 M256I, v1 M128I) M256I {
	return C._mm256_sra_epi32(v0, v1)
}

// Shift packed 16-bit integers in "a" right by "imm8" while shifting in zeros, and store the results in "dst".
func Mm256SrliEpi16(v0 M256I, v1 Int) M256I {
	return C._mm256_srli_epi16(v0, v1)
}

// Shift packed 16-bit integers in "a" right by "count" while shifting in zeros, and store the results in "dst".
func Mm256SrlEpi16(v0 M256I, v1 M128I) M256I {
	return C._mm256_srl_epi16(v0, v1)
}

// Shift packed 32-bit integers in "a" right by "imm8" while shifting in zeros, and store the results in "dst".
func Mm256SrliEpi32(v0 M256I, v1 Int) M256I {
	return C._mm256_srli_epi32(v0, v1)
}

// Shift packed 32-bit integers in "a" right by "count" while shifting in zeros, and store the results in "dst".
func Mm256SrlEpi32(v0 M256I, v1 M128I) M256I {
	return C._mm256_srl_epi32(v0, v1)
}

// Shift packed 64-bit integers in "a" right by "imm8" while shifting in zeros, and store the results in "dst".
func Mm256SrliEpi64(v0 M256I, v1 Int) M256I {
	return C._mm256_srli_epi64(v0, v1)
}

// Shift packed 64-bit integers in "a" right by "count" while shifting in zeros, and store the results in "dst".
func Mm256SrlEpi64(v0 M256I, v1 M128I) M256I {
	return C._mm256_srl_epi64(v0, v1)
}

// Subtract packed 8-bit integers in "b" from packed 8-bit integers in "a", and store the results in "dst".
func Mm256SubEpi8(v0 M256I, v1 M256I) M256I {
	return C._mm256_sub_epi8(v0, v1)
}

// Subtract packed 16-bit integers in "b" from packed 16-bit integers in "a", and store the results in "dst".
func Mm256SubEpi16(v0 M256I, v1 M256I) M256I {
	return C._mm256_sub_epi16(v0, v1)
}

// Subtract packed 32-bit integers in "b" from packed 32-bit integers in "a", and store the results in "dst".
func Mm256SubEpi32(v0 M256I, v1 M256I) M256I {
	return C._mm256_sub_epi32(v0, v1)
}

// Subtract packed 64-bit integers in "b" from packed 64-bit integers in "a", and store the results in "dst".
func Mm256SubEpi64(v0 M256I, v1 M256I) M256I {
	return C._mm256_sub_epi64(v0, v1)
}

// Subtract packed signed 8-bit integers in "b" from packed 8-bit integers in "a" using saturation, and store the results in "dst".
func Mm256SubsEpi8(v0 M256I, v1 M256I) M256I {
	return C._mm256_subs_epi8(v0, v1)
}

// Subtract packed signed 16-bit integers in "b" from packed 16-bit integers in "a" using saturation, and store the results in "dst".
func Mm256SubsEpi16(v0 M256I, v1 M256I) M256I {
	return C._mm256_subs_epi16(v0, v1)
}

// Subtract packed unsigned 8-bit integers in "b" from packed unsigned 8-bit integers in "a" using saturation, and store the results in "dst".
func Mm256SubsEpu8(v0 M256I, v1 M256I) M256I {
	return C._mm256_subs_epu8(v0, v1)
}

// Subtract packed unsigned 16-bit integers in "b" from packed unsigned 16-bit integers in "a" using saturation, and store the results in "dst".
func Mm256SubsEpu16(v0 M256I, v1 M256I) M256I {
	return C._mm256_subs_epu16(v0, v1)
}

// Unpack and interleave 8-bit integers from the high half of each 128-bit lane in "a" and "b", and store the results in "dst".
func Mm256UnpackhiEpi8(v0 M256I, v1 M256I) M256I {
	return C._mm256_unpackhi_epi8(v0, v1)
}

// Unpack and interleave 16-bit integers from the high half of each 128-bit lane in "a" and "b", and store the results in "dst".
func Mm256UnpackhiEpi16(v0 M256I, v1 M256I) M256I {
	return C._mm256_unpackhi_epi16(v0, v1)
}

// Unpack and interleave 32-bit integers from the high half of each 128-bit lane in "a" and "b", and store the results in "dst".
func Mm256UnpackhiEpi32(v0 M256I, v1 M256I) M256I {
	return C._mm256_unpackhi_epi32(v0, v1)
}

// Unpack and interleave 64-bit integers from the high half of each 128-bit lane in "a" and "b", and store the results in "dst".
func Mm256UnpackhiEpi64(v0 M256I, v1 M256I) M256I {
	return C._mm256_unpackhi_epi64(v0, v1)
}

// Unpack and interleave 8-bit integers from the low half of each 128-bit lane in "a" and "b", and store the results in "dst".
func Mm256UnpackloEpi8(v0 M256I, v1 M256I) M256I {
	return C._mm256_unpacklo_epi8(v0, v1)
}

// Unpack and interleave 16-bit integers from the low half of each 128-bit lane in "a" and "b", and store the results in "dst".
func Mm256UnpackloEpi16(v0 M256I, v1 M256I) M256I {
	return C._mm256_unpacklo_epi16(v0, v1)
}

// Unpack and interleave 32-bit integers from the low half of each 128-bit lane in "a" and "b", and store the results in "dst".
func Mm256UnpackloEpi32(v0 M256I, v1 M256I) M256I {
	return C._mm256_unpacklo_epi32(v0, v1)
}

// Unpack and interleave 64-bit integers from the low half of each 128-bit lane in "a" and "b", and store the results in "dst".
func Mm256UnpackloEpi64(v0 M256I, v1 M256I) M256I {
	return C._mm256_unpacklo_epi64(v0, v1)
}

// Compute the bitwise XOR of 256 bits (representing integer data) in "a" and "b", and store the result in "dst".
func Mm256XorSi256(v0 M256I, v1 M256I) M256I {
	return C._mm256_xor_si256(v0, v1)
}

// Broadcast the low single-precision (32-bit) floating-point element from "a" to all elements of "dst".
func MmBroadcastssPs(v0 M128) M128 {
	return C._mm_broadcastss_ps(v0)
}

// Broadcast the low double-precision (64-bit) floating-point element from "a" to all elements of "dst".
// func MmBroadcastsdPd(v0 M128D) M128D {
// 	return C._mm_broadcastsd_pd(v0)
// }

// Broadcast the low single-precision (32-bit) floating-point element from "a" to all elements of "dst".
func Mm256BroadcastssPs(v0 M128) M256 {
	return C._mm256_broadcastss_ps(v0)
}

// Broadcast the low double-precision (64-bit) floating-point element from "a" to all elements of "dst".
func Mm256BroadcastsdPd(v0 M128D) M256D {
	return C._mm256_broadcastsd_pd(v0)
}

// Broadcast 128 bits of integer data from "a" to all 128-bit lanes in "dst".
func Mm256Broadcastsi128Si256(v0 M128I) M256I {
	return C._mm256_broadcastsi128_si256(v0)
}

// Broadcast the low packed 8-bit integer from "a" to all elements of "dst".
func Mm256BroadcastbEpi8(v0 M128I) M256I {
	return C._mm256_broadcastb_epi8(v0)
}

// Broadcast the low packed 16-bit integer from "a" to all elements of "dst".
func Mm256BroadcastwEpi16(v0 M128I) M256I {
	return C._mm256_broadcastw_epi16(v0)
}

// Broadcast the low packed 32-bit integer from "a" to all elements of "dst".
func Mm256BroadcastdEpi32(v0 M128I) M256I {
	return C._mm256_broadcastd_epi32(v0)
}

// Broadcast the low packed 64-bit integer from "a" to all elements of "dst".
func Mm256BroadcastqEpi64(v0 M128I) M256I {
	return C._mm256_broadcastq_epi64(v0)
}

// Broadcast the low packed 8-bit integer from "a" to all elements of "dst".
func MmBroadcastbEpi8(v0 M128I) M128I {
	return C._mm_broadcastb_epi8(v0)
}

// Broadcast the low packed 16-bit integer from "a" to all elements of "dst".
func MmBroadcastwEpi16(v0 M128I) M128I {
	return C._mm_broadcastw_epi16(v0)
}

// Broadcast the low packed 32-bit integer from "a" to all elements of "dst".
func MmBroadcastdEpi32(v0 M128I) M128I {
	return C._mm_broadcastd_epi32(v0)
}

// Broadcast the low packed 64-bit integer from "a" to all elements of "dst".
func MmBroadcastqEpi64(v0 M128I) M128I {
	return C._mm_broadcastq_epi64(v0)
}

// Shuffle 32-bit integers in "a" across lanes using the corresponding index in "idx", and store the results in "dst".
func Mm256Permutevar8X32Epi32(v0 M256I, v1 M256I) M256I {
	return C._mm256_permutevar8x32_epi32(v0, v1)
}

// Shuffle single-precision (32-bit) floating-point elements in "a" across lanes using the corresponding index in "idx".
func Mm256Permutevar8X32Ps(v0 M256, v1 M256I) M256 {
	return C._mm256_permutevar8x32_ps(v0, v1)
}

// Shift packed 32-bit integers in "a" left by the amount specified by the corresponding element in "count" while shifting in zeros, and store the results in "dst".
func Mm256SllvEpi32(v0 M256I, v1 M256I) M256I {
	return C._mm256_sllv_epi32(v0, v1)
}

// Shift packed 32-bit integers in "a" left by the amount specified by the corresponding element in "count" while shifting in zeros, and store the results in "dst".
func MmSllvEpi32(v0 M128I, v1 M128I) M128I {
	return C._mm_sllv_epi32(v0, v1)
}

// Shift packed 64-bit integers in "a" left by the amount specified by the corresponding element in "count" while shifting in zeros, and store the results in "dst".
func Mm256SllvEpi64(v0 M256I, v1 M256I) M256I {
	return C._mm256_sllv_epi64(v0, v1)
}

// Shift packed 64-bit integers in "a" left by the amount specified by the corresponding element in "count" while shifting in zeros, and store the results in "dst".
func MmSllvEpi64(v0 M128I, v1 M128I) M128I {
	return C._mm_sllv_epi64(v0, v1)
}

// Shift packed 32-bit integers in "a" right by the amount specified by the corresponding element in "count" while shifting in sign bits, and store the results in "dst".
func Mm256SravEpi32(v0 M256I, v1 M256I) M256I {
	return C._mm256_srav_epi32(v0, v1)
}

// Shift packed 32-bit integers in "a" right by the amount specified by the corresponding element in "count" while shifting in sign bits, and store the results in "dst".
func MmSravEpi32(v0 M128I, v1 M128I) M128I {
	return C._mm_srav_epi32(v0, v1)
}

// Shift packed 32-bit integers in "a" right by the amount specified by the corresponding element in "count" while shifting in zeros, and store the results in "dst".
func Mm256SrlvEpi32(v0 M256I, v1 M256I) M256I {
	return C._mm256_srlv_epi32(v0, v1)
}

// Shift packed 32-bit integers in "a" right by the amount specified by the corresponding element in "count" while shifting in zeros, and store the results in "dst".
func MmSrlvEpi32(v0 M128I, v1 M128I) M128I {
	return C._mm_srlv_epi32(v0, v1)
}

// Shift packed 64-bit integers in "a" right by the amount specified by the corresponding element in "count" while shifting in zeros, and store the results in "dst".
func Mm256SrlvEpi64(v0 M256I, v1 M256I) M256I {
	return C._mm256_srlv_epi64(v0, v1)
}

// Shift packed 64-bit integers in "a" right by the amount specified by the corresponding element in "count" while shifting in zeros, and store the results in "dst".
func MmSrlvEpi64(v0 M128I, v1 M128I) M128I {
	return C._mm_srlv_epi64(v0, v1)
}
