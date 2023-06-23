package x86

/*
#cgo CFLAGS: -march=native
#include <immintrin.h>
*/
import "C"

// _mm256_abs_epi8
// __always_inline__, __nodebug__, __target__("avx2"), __min_vector_width__(256)
func Mm256AbsEpi8(v0 M256I) M256I {
	return C._mm256_abs_epi8(v0)
}

// _mm256_abs_epi16
// __always_inline__, __nodebug__, __target__("avx2"), __min_vector_width__(256)
func Mm256AbsEpi16(v0 M256I) M256I {
	return C._mm256_abs_epi16(v0)
}

// _mm256_abs_epi32
// __always_inline__, __nodebug__, __target__("avx2"), __min_vector_width__(256)
func Mm256AbsEpi32(v0 M256I) M256I {
	return C._mm256_abs_epi32(v0)
}

// _mm256_packs_epi16
// __always_inline__, __nodebug__, __target__("avx2"), __min_vector_width__(256)
func Mm256PacksEpi16(v0 M256I, v1 M256I) M256I {
	return C._mm256_packs_epi16(v0, v1)
}

// _mm256_packs_epi32
// __always_inline__, __nodebug__, __target__("avx2"), __min_vector_width__(256)
func Mm256PacksEpi32(v0 M256I, v1 M256I) M256I {
	return C._mm256_packs_epi32(v0, v1)
}

// _mm256_packus_epi16
// __always_inline__, __nodebug__, __target__("avx2"), __min_vector_width__(256)
func Mm256PackusEpi16(v0 M256I, v1 M256I) M256I {
	return C._mm256_packus_epi16(v0, v1)
}

// _mm256_packus_epi32
// __always_inline__, __nodebug__, __target__("avx2"), __min_vector_width__(256)
func Mm256PackusEpi32(v0 M256I, v1 M256I) M256I {
	return C._mm256_packus_epi32(v0, v1)
}

// _mm256_add_epi8
// __always_inline__, __nodebug__, __target__("avx2"), __min_vector_width__(256)
func Mm256AddEpi8(v0 M256I, v1 M256I) M256I {
	return C._mm256_add_epi8(v0, v1)
}

// _mm256_add_epi16
// __always_inline__, __nodebug__, __target__("avx2"), __min_vector_width__(256)
func Mm256AddEpi16(v0 M256I, v1 M256I) M256I {
	return C._mm256_add_epi16(v0, v1)
}

// _mm256_add_epi32
// __always_inline__, __nodebug__, __target__("avx2"), __min_vector_width__(256)
func Mm256AddEpi32(v0 M256I, v1 M256I) M256I {
	return C._mm256_add_epi32(v0, v1)
}

// _mm256_add_epi64
// __always_inline__, __nodebug__, __target__("avx2"), __min_vector_width__(256)
func Mm256AddEpi64(v0 M256I, v1 M256I) M256I {
	return C._mm256_add_epi64(v0, v1)
}

// _mm256_adds_epi8
// __always_inline__, __nodebug__, __target__("avx2"), __min_vector_width__(256)
func Mm256AddsEpi8(v0 M256I, v1 M256I) M256I {
	return C._mm256_adds_epi8(v0, v1)
}

// _mm256_adds_epi16
// __always_inline__, __nodebug__, __target__("avx2"), __min_vector_width__(256)
func Mm256AddsEpi16(v0 M256I, v1 M256I) M256I {
	return C._mm256_adds_epi16(v0, v1)
}

// _mm256_adds_epu8
// __always_inline__, __nodebug__, __target__("avx2"), __min_vector_width__(256)
func Mm256AddsEpu8(v0 M256I, v1 M256I) M256I {
	return C._mm256_adds_epu8(v0, v1)
}

// _mm256_adds_epu16
// __always_inline__, __nodebug__, __target__("avx2"), __min_vector_width__(256)
func Mm256AddsEpu16(v0 M256I, v1 M256I) M256I {
	return C._mm256_adds_epu16(v0, v1)
}

// _mm256_and_si256
// __always_inline__, __nodebug__, __target__("avx2"), __min_vector_width__(256)
func Mm256AndSi256(v0 M256I, v1 M256I) M256I {
	return C._mm256_and_si256(v0, v1)
}

// _mm256_andnot_si256
// __always_inline__, __nodebug__, __target__("avx2"), __min_vector_width__(256)
func Mm256AndnotSi256(v0 M256I, v1 M256I) M256I {
	return C._mm256_andnot_si256(v0, v1)
}

// _mm256_avg_epu8
// __always_inline__, __nodebug__, __target__("avx2"), __min_vector_width__(256)
func Mm256AvgEpu8(v0 M256I, v1 M256I) M256I {
	return C._mm256_avg_epu8(v0, v1)
}

// _mm256_avg_epu16
// __always_inline__, __nodebug__, __target__("avx2"), __min_vector_width__(256)
func Mm256AvgEpu16(v0 M256I, v1 M256I) M256I {
	return C._mm256_avg_epu16(v0, v1)
}

// _mm256_blendv_epi8
// __always_inline__, __nodebug__, __target__("avx2"), __min_vector_width__(256)
func Mm256BlendvEpi8(v0 M256I, v1 M256I, v2 M256I) M256I {
	return C._mm256_blendv_epi8(v0, v1, v2)
}

// _mm256_cmpeq_epi8
// __always_inline__, __nodebug__, __target__("avx2"), __min_vector_width__(256)
func Mm256CmpeqEpi8(v0 M256I, v1 M256I) M256I {
	return C._mm256_cmpeq_epi8(v0, v1)
}

// _mm256_cmpeq_epi16
// __always_inline__, __nodebug__, __target__("avx2"), __min_vector_width__(256)
func Mm256CmpeqEpi16(v0 M256I, v1 M256I) M256I {
	return C._mm256_cmpeq_epi16(v0, v1)
}

// _mm256_cmpeq_epi32
// __always_inline__, __nodebug__, __target__("avx2"), __min_vector_width__(256)
func Mm256CmpeqEpi32(v0 M256I, v1 M256I) M256I {
	return C._mm256_cmpeq_epi32(v0, v1)
}

// _mm256_cmpeq_epi64
// __always_inline__, __nodebug__, __target__("avx2"), __min_vector_width__(256)
func Mm256CmpeqEpi64(v0 M256I, v1 M256I) M256I {
	return C._mm256_cmpeq_epi64(v0, v1)
}

// _mm256_cmpgt_epi8
// __always_inline__, __nodebug__, __target__("avx2"), __min_vector_width__(256)
func Mm256CmpgtEpi8(v0 M256I, v1 M256I) M256I {
	return C._mm256_cmpgt_epi8(v0, v1)
}

// _mm256_cmpgt_epi16
// __always_inline__, __nodebug__, __target__("avx2"), __min_vector_width__(256)
func Mm256CmpgtEpi16(v0 M256I, v1 M256I) M256I {
	return C._mm256_cmpgt_epi16(v0, v1)
}

// _mm256_cmpgt_epi32
// __always_inline__, __nodebug__, __target__("avx2"), __min_vector_width__(256)
func Mm256CmpgtEpi32(v0 M256I, v1 M256I) M256I {
	return C._mm256_cmpgt_epi32(v0, v1)
}

// _mm256_cmpgt_epi64
// __always_inline__, __nodebug__, __target__("avx2"), __min_vector_width__(256)
func Mm256CmpgtEpi64(v0 M256I, v1 M256I) M256I {
	return C._mm256_cmpgt_epi64(v0, v1)
}

// _mm256_hadd_epi16
// __always_inline__, __nodebug__, __target__("avx2"), __min_vector_width__(256)
func Mm256HaddEpi16(v0 M256I, v1 M256I) M256I {
	return C._mm256_hadd_epi16(v0, v1)
}

// _mm256_hadd_epi32
// __always_inline__, __nodebug__, __target__("avx2"), __min_vector_width__(256)
func Mm256HaddEpi32(v0 M256I, v1 M256I) M256I {
	return C._mm256_hadd_epi32(v0, v1)
}

// _mm256_hadds_epi16
// __always_inline__, __nodebug__, __target__("avx2"), __min_vector_width__(256)
func Mm256HaddsEpi16(v0 M256I, v1 M256I) M256I {
	return C._mm256_hadds_epi16(v0, v1)
}

// _mm256_hsub_epi16
// __always_inline__, __nodebug__, __target__("avx2"), __min_vector_width__(256)
func Mm256HsubEpi16(v0 M256I, v1 M256I) M256I {
	return C._mm256_hsub_epi16(v0, v1)
}

// _mm256_hsub_epi32
// __always_inline__, __nodebug__, __target__("avx2"), __min_vector_width__(256)
func Mm256HsubEpi32(v0 M256I, v1 M256I) M256I {
	return C._mm256_hsub_epi32(v0, v1)
}

// _mm256_hsubs_epi16
// __always_inline__, __nodebug__, __target__("avx2"), __min_vector_width__(256)
func Mm256HsubsEpi16(v0 M256I, v1 M256I) M256I {
	return C._mm256_hsubs_epi16(v0, v1)
}

// _mm256_maddubs_epi16
// __always_inline__, __nodebug__, __target__("avx2"), __min_vector_width__(256)
func Mm256MaddubsEpi16(v0 M256I, v1 M256I) M256I {
	return C._mm256_maddubs_epi16(v0, v1)
}

// _mm256_madd_epi16
// __always_inline__, __nodebug__, __target__("avx2"), __min_vector_width__(256)
func Mm256MaddEpi16(v0 M256I, v1 M256I) M256I {
	return C._mm256_madd_epi16(v0, v1)
}

// _mm256_max_epi8
// __always_inline__, __nodebug__, __target__("avx2"), __min_vector_width__(256)
func Mm256MaxEpi8(v0 M256I, v1 M256I) M256I {
	return C._mm256_max_epi8(v0, v1)
}

// _mm256_max_epi16
// __always_inline__, __nodebug__, __target__("avx2"), __min_vector_width__(256)
func Mm256MaxEpi16(v0 M256I, v1 M256I) M256I {
	return C._mm256_max_epi16(v0, v1)
}

// _mm256_max_epi32
// __always_inline__, __nodebug__, __target__("avx2"), __min_vector_width__(256)
func Mm256MaxEpi32(v0 M256I, v1 M256I) M256I {
	return C._mm256_max_epi32(v0, v1)
}

// _mm256_max_epu8
// __always_inline__, __nodebug__, __target__("avx2"), __min_vector_width__(256)
func Mm256MaxEpu8(v0 M256I, v1 M256I) M256I {
	return C._mm256_max_epu8(v0, v1)
}

// _mm256_max_epu16
// __always_inline__, __nodebug__, __target__("avx2"), __min_vector_width__(256)
func Mm256MaxEpu16(v0 M256I, v1 M256I) M256I {
	return C._mm256_max_epu16(v0, v1)
}

// _mm256_max_epu32
// __always_inline__, __nodebug__, __target__("avx2"), __min_vector_width__(256)
func Mm256MaxEpu32(v0 M256I, v1 M256I) M256I {
	return C._mm256_max_epu32(v0, v1)
}

// _mm256_min_epi8
// __always_inline__, __nodebug__, __target__("avx2"), __min_vector_width__(256)
func Mm256MinEpi8(v0 M256I, v1 M256I) M256I {
	return C._mm256_min_epi8(v0, v1)
}

// _mm256_min_epi16
// __always_inline__, __nodebug__, __target__("avx2"), __min_vector_width__(256)
func Mm256MinEpi16(v0 M256I, v1 M256I) M256I {
	return C._mm256_min_epi16(v0, v1)
}

// _mm256_min_epi32
// __always_inline__, __nodebug__, __target__("avx2"), __min_vector_width__(256)
func Mm256MinEpi32(v0 M256I, v1 M256I) M256I {
	return C._mm256_min_epi32(v0, v1)
}

// _mm256_min_epu8
// __always_inline__, __nodebug__, __target__("avx2"), __min_vector_width__(256)
func Mm256MinEpu8(v0 M256I, v1 M256I) M256I {
	return C._mm256_min_epu8(v0, v1)
}

// _mm256_min_epu16
// __always_inline__, __nodebug__, __target__("avx2"), __min_vector_width__(256)
func Mm256MinEpu16(v0 M256I, v1 M256I) M256I {
	return C._mm256_min_epu16(v0, v1)
}

// _mm256_min_epu32
// __always_inline__, __nodebug__, __target__("avx2"), __min_vector_width__(256)
func Mm256MinEpu32(v0 M256I, v1 M256I) M256I {
	return C._mm256_min_epu32(v0, v1)
}

// _mm256_movemask_epi8
// __always_inline__, __nodebug__, __target__("avx2"), __min_vector_width__(256)
func Mm256MovemaskEpi8(v0 M256I) Int {
	return C._mm256_movemask_epi8(v0)
}

// _mm256_cvtepi8_epi16
// __always_inline__, __nodebug__, __target__("avx2"), __min_vector_width__(256)
func Mm256Cvtepi8Epi16(v0 M128I) M256I {
	return C._mm256_cvtepi8_epi16(v0)
}

// _mm256_cvtepi8_epi32
// __always_inline__, __nodebug__, __target__("avx2"), __min_vector_width__(256)
func Mm256Cvtepi8Epi32(v0 M128I) M256I {
	return C._mm256_cvtepi8_epi32(v0)
}

// _mm256_cvtepi8_epi64
// __always_inline__, __nodebug__, __target__("avx2"), __min_vector_width__(256)
func Mm256Cvtepi8Epi64(v0 M128I) M256I {
	return C._mm256_cvtepi8_epi64(v0)
}

// _mm256_cvtepi16_epi32
// __always_inline__, __nodebug__, __target__("avx2"), __min_vector_width__(256)
func Mm256Cvtepi16Epi32(v0 M128I) M256I {
	return C._mm256_cvtepi16_epi32(v0)
}

// _mm256_cvtepi16_epi64
// __always_inline__, __nodebug__, __target__("avx2"), __min_vector_width__(256)
func Mm256Cvtepi16Epi64(v0 M128I) M256I {
	return C._mm256_cvtepi16_epi64(v0)
}

// _mm256_cvtepi32_epi64
// __always_inline__, __nodebug__, __target__("avx2"), __min_vector_width__(256)
func Mm256Cvtepi32Epi64(v0 M128I) M256I {
	return C._mm256_cvtepi32_epi64(v0)
}

// _mm256_cvtepu8_epi16
// __always_inline__, __nodebug__, __target__("avx2"), __min_vector_width__(256)
func Mm256Cvtepu8Epi16(v0 M128I) M256I {
	return C._mm256_cvtepu8_epi16(v0)
}

// _mm256_cvtepu8_epi32
// __always_inline__, __nodebug__, __target__("avx2"), __min_vector_width__(256)
func Mm256Cvtepu8Epi32(v0 M128I) M256I {
	return C._mm256_cvtepu8_epi32(v0)
}

// _mm256_cvtepu8_epi64
// __always_inline__, __nodebug__, __target__("avx2"), __min_vector_width__(256)
func Mm256Cvtepu8Epi64(v0 M128I) M256I {
	return C._mm256_cvtepu8_epi64(v0)
}

// _mm256_cvtepu16_epi32
// __always_inline__, __nodebug__, __target__("avx2"), __min_vector_width__(256)
func Mm256Cvtepu16Epi32(v0 M128I) M256I {
	return C._mm256_cvtepu16_epi32(v0)
}

// _mm256_cvtepu16_epi64
// __always_inline__, __nodebug__, __target__("avx2"), __min_vector_width__(256)
func Mm256Cvtepu16Epi64(v0 M128I) M256I {
	return C._mm256_cvtepu16_epi64(v0)
}

// _mm256_cvtepu32_epi64
// __always_inline__, __nodebug__, __target__("avx2"), __min_vector_width__(256)
func Mm256Cvtepu32Epi64(v0 M128I) M256I {
	return C._mm256_cvtepu32_epi64(v0)
}

// _mm256_mul_epi32
// __always_inline__, __nodebug__, __target__("avx2"), __min_vector_width__(256)
func Mm256MulEpi32(v0 M256I, v1 M256I) M256I {
	return C._mm256_mul_epi32(v0, v1)
}

// _mm256_mulhrs_epi16
// __always_inline__, __nodebug__, __target__("avx2"), __min_vector_width__(256)
func Mm256MulhrsEpi16(v0 M256I, v1 M256I) M256I {
	return C._mm256_mulhrs_epi16(v0, v1)
}

// _mm256_mulhi_epu16
// __always_inline__, __nodebug__, __target__("avx2"), __min_vector_width__(256)
func Mm256MulhiEpu16(v0 M256I, v1 M256I) M256I {
	return C._mm256_mulhi_epu16(v0, v1)
}

// _mm256_mulhi_epi16
// __always_inline__, __nodebug__, __target__("avx2"), __min_vector_width__(256)
func Mm256MulhiEpi16(v0 M256I, v1 M256I) M256I {
	return C._mm256_mulhi_epi16(v0, v1)
}

// _mm256_mullo_epi16
// __always_inline__, __nodebug__, __target__("avx2"), __min_vector_width__(256)
func Mm256MulloEpi16(v0 M256I, v1 M256I) M256I {
	return C._mm256_mullo_epi16(v0, v1)
}

// _mm256_mullo_epi32
// __always_inline__, __nodebug__, __target__("avx2"), __min_vector_width__(256)
func Mm256MulloEpi32(v0 M256I, v1 M256I) M256I {
	return C._mm256_mullo_epi32(v0, v1)
}

// _mm256_mul_epu32
// __always_inline__, __nodebug__, __target__("avx2"), __min_vector_width__(256)
func Mm256MulEpu32(v0 M256I, v1 M256I) M256I {
	return C._mm256_mul_epu32(v0, v1)
}

// _mm256_or_si256
// __always_inline__, __nodebug__, __target__("avx2"), __min_vector_width__(256)
func Mm256OrSi256(v0 M256I, v1 M256I) M256I {
	return C._mm256_or_si256(v0, v1)
}

// _mm256_sad_epu8
// __always_inline__, __nodebug__, __target__("avx2"), __min_vector_width__(256)
func Mm256SadEpu8(v0 M256I, v1 M256I) M256I {
	return C._mm256_sad_epu8(v0, v1)
}

// _mm256_shuffle_epi8
// __always_inline__, __nodebug__, __target__("avx2"), __min_vector_width__(256)
func Mm256ShuffleEpi8(v0 M256I, v1 M256I) M256I {
	return C._mm256_shuffle_epi8(v0, v1)
}

// _mm256_sign_epi8
// __always_inline__, __nodebug__, __target__("avx2"), __min_vector_width__(256)
func Mm256SignEpi8(v0 M256I, v1 M256I) M256I {
	return C._mm256_sign_epi8(v0, v1)
}

// _mm256_sign_epi16
// __always_inline__, __nodebug__, __target__("avx2"), __min_vector_width__(256)
func Mm256SignEpi16(v0 M256I, v1 M256I) M256I {
	return C._mm256_sign_epi16(v0, v1)
}

// _mm256_sign_epi32
// __always_inline__, __nodebug__, __target__("avx2"), __min_vector_width__(256)
func Mm256SignEpi32(v0 M256I, v1 M256I) M256I {
	return C._mm256_sign_epi32(v0, v1)
}

// _mm256_slli_epi16
// __always_inline__, __nodebug__, __target__("avx2"), __min_vector_width__(256)
func Mm256SlliEpi16(v0 M256I, v1 Int) M256I {
	return C._mm256_slli_epi16(v0, v1)
}

// _mm256_sll_epi16
// __always_inline__, __nodebug__, __target__("avx2"), __min_vector_width__(256)
func Mm256SllEpi16(v0 M256I, v1 M128I) M256I {
	return C._mm256_sll_epi16(v0, v1)
}

// _mm256_slli_epi32
// __always_inline__, __nodebug__, __target__("avx2"), __min_vector_width__(256)
func Mm256SlliEpi32(v0 M256I, v1 Int) M256I {
	return C._mm256_slli_epi32(v0, v1)
}

// _mm256_sll_epi32
// __always_inline__, __nodebug__, __target__("avx2"), __min_vector_width__(256)
func Mm256SllEpi32(v0 M256I, v1 M128I) M256I {
	return C._mm256_sll_epi32(v0, v1)
}

// _mm256_slli_epi64
// __always_inline__, __nodebug__, __target__("avx2"), __min_vector_width__(256)
func Mm256SlliEpi64(v0 M256I, v1 Int) M256I {
	return C._mm256_slli_epi64(v0, v1)
}

// _mm256_sll_epi64
// __always_inline__, __nodebug__, __target__("avx2"), __min_vector_width__(256)
func Mm256SllEpi64(v0 M256I, v1 M128I) M256I {
	return C._mm256_sll_epi64(v0, v1)
}

// _mm256_srai_epi16
// __always_inline__, __nodebug__, __target__("avx2"), __min_vector_width__(256)
func Mm256SraiEpi16(v0 M256I, v1 Int) M256I {
	return C._mm256_srai_epi16(v0, v1)
}

// _mm256_sra_epi16
// __always_inline__, __nodebug__, __target__("avx2"), __min_vector_width__(256)
func Mm256SraEpi16(v0 M256I, v1 M128I) M256I {
	return C._mm256_sra_epi16(v0, v1)
}

// _mm256_srai_epi32
// __always_inline__, __nodebug__, __target__("avx2"), __min_vector_width__(256)
func Mm256SraiEpi32(v0 M256I, v1 Int) M256I {
	return C._mm256_srai_epi32(v0, v1)
}

// _mm256_sra_epi32
// __always_inline__, __nodebug__, __target__("avx2"), __min_vector_width__(256)
func Mm256SraEpi32(v0 M256I, v1 M128I) M256I {
	return C._mm256_sra_epi32(v0, v1)
}

// _mm256_srli_epi16
// __always_inline__, __nodebug__, __target__("avx2"), __min_vector_width__(256)
func Mm256SrliEpi16(v0 M256I, v1 Int) M256I {
	return C._mm256_srli_epi16(v0, v1)
}

// _mm256_srl_epi16
// __always_inline__, __nodebug__, __target__("avx2"), __min_vector_width__(256)
func Mm256SrlEpi16(v0 M256I, v1 M128I) M256I {
	return C._mm256_srl_epi16(v0, v1)
}

// _mm256_srli_epi32
// __always_inline__, __nodebug__, __target__("avx2"), __min_vector_width__(256)
func Mm256SrliEpi32(v0 M256I, v1 Int) M256I {
	return C._mm256_srli_epi32(v0, v1)
}

// _mm256_srl_epi32
// __always_inline__, __nodebug__, __target__("avx2"), __min_vector_width__(256)
func Mm256SrlEpi32(v0 M256I, v1 M128I) M256I {
	return C._mm256_srl_epi32(v0, v1)
}

// _mm256_srli_epi64
// __always_inline__, __nodebug__, __target__("avx2"), __min_vector_width__(256)
func Mm256SrliEpi64(v0 M256I, v1 Int) M256I {
	return C._mm256_srli_epi64(v0, v1)
}

// _mm256_srl_epi64
// __always_inline__, __nodebug__, __target__("avx2"), __min_vector_width__(256)
func Mm256SrlEpi64(v0 M256I, v1 M128I) M256I {
	return C._mm256_srl_epi64(v0, v1)
}

// _mm256_sub_epi8
// __always_inline__, __nodebug__, __target__("avx2"), __min_vector_width__(256)
func Mm256SubEpi8(v0 M256I, v1 M256I) M256I {
	return C._mm256_sub_epi8(v0, v1)
}

// _mm256_sub_epi16
// __always_inline__, __nodebug__, __target__("avx2"), __min_vector_width__(256)
func Mm256SubEpi16(v0 M256I, v1 M256I) M256I {
	return C._mm256_sub_epi16(v0, v1)
}

// _mm256_sub_epi32
// __always_inline__, __nodebug__, __target__("avx2"), __min_vector_width__(256)
func Mm256SubEpi32(v0 M256I, v1 M256I) M256I {
	return C._mm256_sub_epi32(v0, v1)
}

// _mm256_sub_epi64
// __always_inline__, __nodebug__, __target__("avx2"), __min_vector_width__(256)
func Mm256SubEpi64(v0 M256I, v1 M256I) M256I {
	return C._mm256_sub_epi64(v0, v1)
}

// _mm256_subs_epi8
// __always_inline__, __nodebug__, __target__("avx2"), __min_vector_width__(256)
func Mm256SubsEpi8(v0 M256I, v1 M256I) M256I {
	return C._mm256_subs_epi8(v0, v1)
}

// _mm256_subs_epi16
// __always_inline__, __nodebug__, __target__("avx2"), __min_vector_width__(256)
func Mm256SubsEpi16(v0 M256I, v1 M256I) M256I {
	return C._mm256_subs_epi16(v0, v1)
}

// _mm256_subs_epu8
// __always_inline__, __nodebug__, __target__("avx2"), __min_vector_width__(256)
func Mm256SubsEpu8(v0 M256I, v1 M256I) M256I {
	return C._mm256_subs_epu8(v0, v1)
}

// _mm256_subs_epu16
// __always_inline__, __nodebug__, __target__("avx2"), __min_vector_width__(256)
func Mm256SubsEpu16(v0 M256I, v1 M256I) M256I {
	return C._mm256_subs_epu16(v0, v1)
}

// _mm256_unpackhi_epi8
// __always_inline__, __nodebug__, __target__("avx2"), __min_vector_width__(256)
func Mm256UnpackhiEpi8(v0 M256I, v1 M256I) M256I {
	return C._mm256_unpackhi_epi8(v0, v1)
}

// _mm256_unpackhi_epi16
// __always_inline__, __nodebug__, __target__("avx2"), __min_vector_width__(256)
func Mm256UnpackhiEpi16(v0 M256I, v1 M256I) M256I {
	return C._mm256_unpackhi_epi16(v0, v1)
}

// _mm256_unpackhi_epi32
// __always_inline__, __nodebug__, __target__("avx2"), __min_vector_width__(256)
func Mm256UnpackhiEpi32(v0 M256I, v1 M256I) M256I {
	return C._mm256_unpackhi_epi32(v0, v1)
}

// _mm256_unpackhi_epi64
// __always_inline__, __nodebug__, __target__("avx2"), __min_vector_width__(256)
func Mm256UnpackhiEpi64(v0 M256I, v1 M256I) M256I {
	return C._mm256_unpackhi_epi64(v0, v1)
}

// _mm256_unpacklo_epi8
// __always_inline__, __nodebug__, __target__("avx2"), __min_vector_width__(256)
func Mm256UnpackloEpi8(v0 M256I, v1 M256I) M256I {
	return C._mm256_unpacklo_epi8(v0, v1)
}

// _mm256_unpacklo_epi16
// __always_inline__, __nodebug__, __target__("avx2"), __min_vector_width__(256)
func Mm256UnpackloEpi16(v0 M256I, v1 M256I) M256I {
	return C._mm256_unpacklo_epi16(v0, v1)
}

// _mm256_unpacklo_epi32
// __always_inline__, __nodebug__, __target__("avx2"), __min_vector_width__(256)
func Mm256UnpackloEpi32(v0 M256I, v1 M256I) M256I {
	return C._mm256_unpacklo_epi32(v0, v1)
}

// _mm256_unpacklo_epi64
// __always_inline__, __nodebug__, __target__("avx2"), __min_vector_width__(256)
func Mm256UnpackloEpi64(v0 M256I, v1 M256I) M256I {
	return C._mm256_unpacklo_epi64(v0, v1)
}

// _mm256_xor_si256
// __always_inline__, __nodebug__, __target__("avx2"), __min_vector_width__(256)
func Mm256XorSi256(v0 M256I, v1 M256I) M256I {
	return C._mm256_xor_si256(v0, v1)
}

// _mm_broadcastss_ps
// __always_inline__, __nodebug__, __target__("avx2"), __min_vector_width__(128)
func MmBroadcastssPs(v0 M128) M128 {
	return C._mm_broadcastss_ps(v0)
}

// _mm_broadcastsd_pd
// __always_inline__, __nodebug__, __target__("avx2"), __min_vector_width__(128)
// func MmBroadcastsdPd(v0 M128D) M128D {
// 	return C._mm_broadcastsd_pd(v0)
// }

// _mm256_broadcastss_ps
// __always_inline__, __nodebug__, __target__("avx2"), __min_vector_width__(256)
func Mm256BroadcastssPs(v0 M128) M256 {
	return C._mm256_broadcastss_ps(v0)
}

// _mm256_broadcastsd_pd
// __always_inline__, __nodebug__, __target__("avx2"), __min_vector_width__(256)
func Mm256BroadcastsdPd(v0 M128D) M256D {
	return C._mm256_broadcastsd_pd(v0)
}

// _mm256_broadcastsi128_si256
// __always_inline__, __nodebug__, __target__("avx2"), __min_vector_width__(256)
func Mm256Broadcastsi128Si256(v0 M128I) M256I {
	return C._mm256_broadcastsi128_si256(v0)
}

// _mm256_broadcastb_epi8
// __always_inline__, __nodebug__, __target__("avx2"), __min_vector_width__(256)
func Mm256BroadcastbEpi8(v0 M128I) M256I {
	return C._mm256_broadcastb_epi8(v0)
}

// _mm256_broadcastw_epi16
// __always_inline__, __nodebug__, __target__("avx2"), __min_vector_width__(256)
func Mm256BroadcastwEpi16(v0 M128I) M256I {
	return C._mm256_broadcastw_epi16(v0)
}

// _mm256_broadcastd_epi32
// __always_inline__, __nodebug__, __target__("avx2"), __min_vector_width__(256)
func Mm256BroadcastdEpi32(v0 M128I) M256I {
	return C._mm256_broadcastd_epi32(v0)
}

// _mm256_broadcastq_epi64
// __always_inline__, __nodebug__, __target__("avx2"), __min_vector_width__(256)
func Mm256BroadcastqEpi64(v0 M128I) M256I {
	return C._mm256_broadcastq_epi64(v0)
}

// _mm_broadcastb_epi8
// __always_inline__, __nodebug__, __target__("avx2"), __min_vector_width__(128)
func MmBroadcastbEpi8(v0 M128I) M128I {
	return C._mm_broadcastb_epi8(v0)
}

// _mm_broadcastw_epi16
// __always_inline__, __nodebug__, __target__("avx2"), __min_vector_width__(128)
func MmBroadcastwEpi16(v0 M128I) M128I {
	return C._mm_broadcastw_epi16(v0)
}

// _mm_broadcastd_epi32
// __always_inline__, __nodebug__, __target__("avx2"), __min_vector_width__(128)
func MmBroadcastdEpi32(v0 M128I) M128I {
	return C._mm_broadcastd_epi32(v0)
}

// _mm_broadcastq_epi64
// __always_inline__, __nodebug__, __target__("avx2"), __min_vector_width__(128)
func MmBroadcastqEpi64(v0 M128I) M128I {
	return C._mm_broadcastq_epi64(v0)
}

// _mm256_permutevar8x32_epi32
// __always_inline__, __nodebug__, __target__("avx2"), __min_vector_width__(256)
func Mm256Permutevar8X32Epi32(v0 M256I, v1 M256I) M256I {
	return C._mm256_permutevar8x32_epi32(v0, v1)
}

// _mm256_permutevar8x32_ps
// __always_inline__, __nodebug__, __target__("avx2"), __min_vector_width__(256)
func Mm256Permutevar8X32Ps(v0 M256, v1 M256I) M256 {
	return C._mm256_permutevar8x32_ps(v0, v1)
}

// _mm256_sllv_epi32
// __always_inline__, __nodebug__, __target__("avx2"), __min_vector_width__(256)
func Mm256SllvEpi32(v0 M256I, v1 M256I) M256I {
	return C._mm256_sllv_epi32(v0, v1)
}

// _mm_sllv_epi32
// __always_inline__, __nodebug__, __target__("avx2"), __min_vector_width__(128)
func MmSllvEpi32(v0 M128I, v1 M128I) M128I {
	return C._mm_sllv_epi32(v0, v1)
}

// _mm256_sllv_epi64
// __always_inline__, __nodebug__, __target__("avx2"), __min_vector_width__(256)
func Mm256SllvEpi64(v0 M256I, v1 M256I) M256I {
	return C._mm256_sllv_epi64(v0, v1)
}

// _mm_sllv_epi64
// __always_inline__, __nodebug__, __target__("avx2"), __min_vector_width__(128)
func MmSllvEpi64(v0 M128I, v1 M128I) M128I {
	return C._mm_sllv_epi64(v0, v1)
}

// _mm256_srav_epi32
// __always_inline__, __nodebug__, __target__("avx2"), __min_vector_width__(256)
func Mm256SravEpi32(v0 M256I, v1 M256I) M256I {
	return C._mm256_srav_epi32(v0, v1)
}

// _mm_srav_epi32
// __always_inline__, __nodebug__, __target__("avx2"), __min_vector_width__(128)
func MmSravEpi32(v0 M128I, v1 M128I) M128I {
	return C._mm_srav_epi32(v0, v1)
}

// _mm256_srlv_epi32
// __always_inline__, __nodebug__, __target__("avx2"), __min_vector_width__(256)
func Mm256SrlvEpi32(v0 M256I, v1 M256I) M256I {
	return C._mm256_srlv_epi32(v0, v1)
}

// _mm_srlv_epi32
// __always_inline__, __nodebug__, __target__("avx2"), __min_vector_width__(128)
func MmSrlvEpi32(v0 M128I, v1 M128I) M128I {
	return C._mm_srlv_epi32(v0, v1)
}

// _mm256_srlv_epi64
// __always_inline__, __nodebug__, __target__("avx2"), __min_vector_width__(256)
func Mm256SrlvEpi64(v0 M256I, v1 M256I) M256I {
	return C._mm256_srlv_epi64(v0, v1)
}

// _mm_srlv_epi64
// __always_inline__, __nodebug__, __target__("avx2"), __min_vector_width__(128)
func MmSrlvEpi64(v0 M128I, v1 M128I) M128I {
	return C._mm_srlv_epi64(v0, v1)
}
