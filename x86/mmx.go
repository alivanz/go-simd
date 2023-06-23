package x86

/*
#cgo CFLAGS: 
#include <immintrin.h>
*/
import "C"

// _mm_empty
// __always_inline__, __nodebug__, __target__("mmx")
func MmEmpty() {
	C._mm_empty()
}

// _mm_cvtsi32_si64
// __always_inline__, __nodebug__, __target__("mmx"), __min_vector_width__(64)
func MmCvtsi32Si64(v0 Int) M64 {
	return C._mm_cvtsi32_si64(v0)
}

// _mm_cvtsi64_si32
// __always_inline__, __nodebug__, __target__("mmx"), __min_vector_width__(64)
func MmCvtsi64Si32(v0 M64) Int {
	return C._mm_cvtsi64_si32(v0)
}

// _mm_cvtsi64_m64
// __always_inline__, __nodebug__, __target__("mmx"), __min_vector_width__(64)
func MmCvtsi64M64(v0 Longlong) M64 {
	return C._mm_cvtsi64_m64(v0)
}

// _mm_cvtm64_si64
// __always_inline__, __nodebug__, __target__("mmx"), __min_vector_width__(64)
func MmCvtm64Si64(v0 M64) Longlong {
	return C._mm_cvtm64_si64(v0)
}

// _mm_packs_pi16
// __always_inline__, __nodebug__, __target__("mmx"), __min_vector_width__(64)
func MmPacksPi16(v0 M64, v1 M64) M64 {
	return C._mm_packs_pi16(v0, v1)
}

// _mm_packs_pi32
// __always_inline__, __nodebug__, __target__("mmx"), __min_vector_width__(64)
func MmPacksPi32(v0 M64, v1 M64) M64 {
	return C._mm_packs_pi32(v0, v1)
}

// _mm_packs_pu16
// __always_inline__, __nodebug__, __target__("mmx"), __min_vector_width__(64)
func MmPacksPu16(v0 M64, v1 M64) M64 {
	return C._mm_packs_pu16(v0, v1)
}

// _mm_unpackhi_pi8
// __always_inline__, __nodebug__, __target__("mmx"), __min_vector_width__(64)
func MmUnpackhiPi8(v0 M64, v1 M64) M64 {
	return C._mm_unpackhi_pi8(v0, v1)
}

// _mm_unpackhi_pi16
// __always_inline__, __nodebug__, __target__("mmx"), __min_vector_width__(64)
func MmUnpackhiPi16(v0 M64, v1 M64) M64 {
	return C._mm_unpackhi_pi16(v0, v1)
}

// _mm_unpackhi_pi32
// __always_inline__, __nodebug__, __target__("mmx"), __min_vector_width__(64)
func MmUnpackhiPi32(v0 M64, v1 M64) M64 {
	return C._mm_unpackhi_pi32(v0, v1)
}

// _mm_unpacklo_pi8
// __always_inline__, __nodebug__, __target__("mmx"), __min_vector_width__(64)
func MmUnpackloPi8(v0 M64, v1 M64) M64 {
	return C._mm_unpacklo_pi8(v0, v1)
}

// _mm_unpacklo_pi16
// __always_inline__, __nodebug__, __target__("mmx"), __min_vector_width__(64)
func MmUnpackloPi16(v0 M64, v1 M64) M64 {
	return C._mm_unpacklo_pi16(v0, v1)
}

// _mm_unpacklo_pi32
// __always_inline__, __nodebug__, __target__("mmx"), __min_vector_width__(64)
func MmUnpackloPi32(v0 M64, v1 M64) M64 {
	return C._mm_unpacklo_pi32(v0, v1)
}

// _mm_add_pi8
// __always_inline__, __nodebug__, __target__("mmx"), __min_vector_width__(64)
func MmAddPi8(v0 M64, v1 M64) M64 {
	return C._mm_add_pi8(v0, v1)
}

// _mm_add_pi16
// __always_inline__, __nodebug__, __target__("mmx"), __min_vector_width__(64)
func MmAddPi16(v0 M64, v1 M64) M64 {
	return C._mm_add_pi16(v0, v1)
}

// _mm_add_pi32
// __always_inline__, __nodebug__, __target__("mmx"), __min_vector_width__(64)
func MmAddPi32(v0 M64, v1 M64) M64 {
	return C._mm_add_pi32(v0, v1)
}

// _mm_adds_pi8
// __always_inline__, __nodebug__, __target__("mmx"), __min_vector_width__(64)
func MmAddsPi8(v0 M64, v1 M64) M64 {
	return C._mm_adds_pi8(v0, v1)
}

// _mm_adds_pi16
// __always_inline__, __nodebug__, __target__("mmx"), __min_vector_width__(64)
func MmAddsPi16(v0 M64, v1 M64) M64 {
	return C._mm_adds_pi16(v0, v1)
}

// _mm_adds_pu8
// __always_inline__, __nodebug__, __target__("mmx"), __min_vector_width__(64)
func MmAddsPu8(v0 M64, v1 M64) M64 {
	return C._mm_adds_pu8(v0, v1)
}

// _mm_adds_pu16
// __always_inline__, __nodebug__, __target__("mmx"), __min_vector_width__(64)
func MmAddsPu16(v0 M64, v1 M64) M64 {
	return C._mm_adds_pu16(v0, v1)
}

// _mm_sub_pi8
// __always_inline__, __nodebug__, __target__("mmx"), __min_vector_width__(64)
func MmSubPi8(v0 M64, v1 M64) M64 {
	return C._mm_sub_pi8(v0, v1)
}

// _mm_sub_pi16
// __always_inline__, __nodebug__, __target__("mmx"), __min_vector_width__(64)
func MmSubPi16(v0 M64, v1 M64) M64 {
	return C._mm_sub_pi16(v0, v1)
}

// _mm_sub_pi32
// __always_inline__, __nodebug__, __target__("mmx"), __min_vector_width__(64)
func MmSubPi32(v0 M64, v1 M64) M64 {
	return C._mm_sub_pi32(v0, v1)
}

// _mm_subs_pi8
// __always_inline__, __nodebug__, __target__("mmx"), __min_vector_width__(64)
func MmSubsPi8(v0 M64, v1 M64) M64 {
	return C._mm_subs_pi8(v0, v1)
}

// _mm_subs_pi16
// __always_inline__, __nodebug__, __target__("mmx"), __min_vector_width__(64)
func MmSubsPi16(v0 M64, v1 M64) M64 {
	return C._mm_subs_pi16(v0, v1)
}

// _mm_subs_pu8
// __always_inline__, __nodebug__, __target__("mmx"), __min_vector_width__(64)
func MmSubsPu8(v0 M64, v1 M64) M64 {
	return C._mm_subs_pu8(v0, v1)
}

// _mm_subs_pu16
// __always_inline__, __nodebug__, __target__("mmx"), __min_vector_width__(64)
func MmSubsPu16(v0 M64, v1 M64) M64 {
	return C._mm_subs_pu16(v0, v1)
}

// _mm_madd_pi16
// __always_inline__, __nodebug__, __target__("mmx"), __min_vector_width__(64)
func MmMaddPi16(v0 M64, v1 M64) M64 {
	return C._mm_madd_pi16(v0, v1)
}

// _mm_mulhi_pi16
// __always_inline__, __nodebug__, __target__("mmx"), __min_vector_width__(64)
func MmMulhiPi16(v0 M64, v1 M64) M64 {
	return C._mm_mulhi_pi16(v0, v1)
}

// _mm_mullo_pi16
// __always_inline__, __nodebug__, __target__("mmx"), __min_vector_width__(64)
func MmMulloPi16(v0 M64, v1 M64) M64 {
	return C._mm_mullo_pi16(v0, v1)
}

// _mm_sll_pi16
// __always_inline__, __nodebug__, __target__("mmx"), __min_vector_width__(64)
func MmSllPi16(v0 M64, v1 M64) M64 {
	return C._mm_sll_pi16(v0, v1)
}

// _mm_slli_pi16
// __always_inline__, __nodebug__, __target__("mmx"), __min_vector_width__(64)
func MmSlliPi16(v0 M64, v1 Int) M64 {
	return C._mm_slli_pi16(v0, v1)
}

// _mm_sll_pi32
// __always_inline__, __nodebug__, __target__("mmx"), __min_vector_width__(64)
func MmSllPi32(v0 M64, v1 M64) M64 {
	return C._mm_sll_pi32(v0, v1)
}

// _mm_slli_pi32
// __always_inline__, __nodebug__, __target__("mmx"), __min_vector_width__(64)
func MmSlliPi32(v0 M64, v1 Int) M64 {
	return C._mm_slli_pi32(v0, v1)
}

// _mm_sll_si64
// __always_inline__, __nodebug__, __target__("mmx"), __min_vector_width__(64)
func MmSllSi64(v0 M64, v1 M64) M64 {
	return C._mm_sll_si64(v0, v1)
}

// _mm_slli_si64
// __always_inline__, __nodebug__, __target__("mmx"), __min_vector_width__(64)
func MmSlliSi64(v0 M64, v1 Int) M64 {
	return C._mm_slli_si64(v0, v1)
}

// _mm_sra_pi16
// __always_inline__, __nodebug__, __target__("mmx"), __min_vector_width__(64)
func MmSraPi16(v0 M64, v1 M64) M64 {
	return C._mm_sra_pi16(v0, v1)
}

// _mm_srai_pi16
// __always_inline__, __nodebug__, __target__("mmx"), __min_vector_width__(64)
func MmSraiPi16(v0 M64, v1 Int) M64 {
	return C._mm_srai_pi16(v0, v1)
}

// _mm_sra_pi32
// __always_inline__, __nodebug__, __target__("mmx"), __min_vector_width__(64)
func MmSraPi32(v0 M64, v1 M64) M64 {
	return C._mm_sra_pi32(v0, v1)
}

// _mm_srai_pi32
// __always_inline__, __nodebug__, __target__("mmx"), __min_vector_width__(64)
func MmSraiPi32(v0 M64, v1 Int) M64 {
	return C._mm_srai_pi32(v0, v1)
}

// _mm_srl_pi16
// __always_inline__, __nodebug__, __target__("mmx"), __min_vector_width__(64)
func MmSrlPi16(v0 M64, v1 M64) M64 {
	return C._mm_srl_pi16(v0, v1)
}

// _mm_srli_pi16
// __always_inline__, __nodebug__, __target__("mmx"), __min_vector_width__(64)
func MmSrliPi16(v0 M64, v1 Int) M64 {
	return C._mm_srli_pi16(v0, v1)
}

// _mm_srl_pi32
// __always_inline__, __nodebug__, __target__("mmx"), __min_vector_width__(64)
func MmSrlPi32(v0 M64, v1 M64) M64 {
	return C._mm_srl_pi32(v0, v1)
}

// _mm_srli_pi32
// __always_inline__, __nodebug__, __target__("mmx"), __min_vector_width__(64)
func MmSrliPi32(v0 M64, v1 Int) M64 {
	return C._mm_srli_pi32(v0, v1)
}

// _mm_srl_si64
// __always_inline__, __nodebug__, __target__("mmx"), __min_vector_width__(64)
func MmSrlSi64(v0 M64, v1 M64) M64 {
	return C._mm_srl_si64(v0, v1)
}

// _mm_srli_si64
// __always_inline__, __nodebug__, __target__("mmx"), __min_vector_width__(64)
func MmSrliSi64(v0 M64, v1 Int) M64 {
	return C._mm_srli_si64(v0, v1)
}

// _mm_and_si64
// __always_inline__, __nodebug__, __target__("mmx"), __min_vector_width__(64)
func MmAndSi64(v0 M64, v1 M64) M64 {
	return C._mm_and_si64(v0, v1)
}

// _mm_andnot_si64
// __always_inline__, __nodebug__, __target__("mmx"), __min_vector_width__(64)
func MmAndnotSi64(v0 M64, v1 M64) M64 {
	return C._mm_andnot_si64(v0, v1)
}

// _mm_or_si64
// __always_inline__, __nodebug__, __target__("mmx"), __min_vector_width__(64)
func MmOrSi64(v0 M64, v1 M64) M64 {
	return C._mm_or_si64(v0, v1)
}

// _mm_xor_si64
// __always_inline__, __nodebug__, __target__("mmx"), __min_vector_width__(64)
func MmXorSi64(v0 M64, v1 M64) M64 {
	return C._mm_xor_si64(v0, v1)
}

// _mm_cmpeq_pi8
// __always_inline__, __nodebug__, __target__("mmx"), __min_vector_width__(64)
func MmCmpeqPi8(v0 M64, v1 M64) M64 {
	return C._mm_cmpeq_pi8(v0, v1)
}

// _mm_cmpeq_pi16
// __always_inline__, __nodebug__, __target__("mmx"), __min_vector_width__(64)
func MmCmpeqPi16(v0 M64, v1 M64) M64 {
	return C._mm_cmpeq_pi16(v0, v1)
}

// _mm_cmpeq_pi32
// __always_inline__, __nodebug__, __target__("mmx"), __min_vector_width__(64)
func MmCmpeqPi32(v0 M64, v1 M64) M64 {
	return C._mm_cmpeq_pi32(v0, v1)
}

// _mm_cmpgt_pi8
// __always_inline__, __nodebug__, __target__("mmx"), __min_vector_width__(64)
func MmCmpgtPi8(v0 M64, v1 M64) M64 {
	return C._mm_cmpgt_pi8(v0, v1)
}

// _mm_cmpgt_pi16
// __always_inline__, __nodebug__, __target__("mmx"), __min_vector_width__(64)
func MmCmpgtPi16(v0 M64, v1 M64) M64 {
	return C._mm_cmpgt_pi16(v0, v1)
}

// _mm_cmpgt_pi32
// __always_inline__, __nodebug__, __target__("mmx"), __min_vector_width__(64)
func MmCmpgtPi32(v0 M64, v1 M64) M64 {
	return C._mm_cmpgt_pi32(v0, v1)
}

// _mm_setzero_si64
// __always_inline__, __nodebug__, __target__("mmx"), __min_vector_width__(64)
func MmSetzeroSi64() M64 {
	return C._mm_setzero_si64()
}

// _mm_set_pi32
// __always_inline__, __nodebug__, __target__("mmx"), __min_vector_width__(64)
func MmSetPi32(v0 Int, v1 Int) M64 {
	return C._mm_set_pi32(v0, v1)
}

// _mm_set_pi16
// __always_inline__, __nodebug__, __target__("mmx"), __min_vector_width__(64)
func MmSetPi16(v0 Short, v1 Short, v2 Short, v3 Short) M64 {
	return C._mm_set_pi16(v0, v1, v2, v3)
}

// _mm_set_pi8
// __always_inline__, __nodebug__, __target__("mmx"), __min_vector_width__(64)
func MmSetPi8(v0 Char, v1 Char, v2 Char, v3 Char, v4 Char, v5 Char, v6 Char, v7 Char) M64 {
	return C._mm_set_pi8(v0, v1, v2, v3, v4, v5, v6, v7)
}

// _mm_set1_pi32
// __always_inline__, __nodebug__, __target__("mmx"), __min_vector_width__(64)
func MmSet1Pi32(v0 Int) M64 {
	return C._mm_set1_pi32(v0)
}

// _mm_set1_pi16
// __always_inline__, __nodebug__, __target__("mmx"), __min_vector_width__(64)
func MmSet1Pi16(v0 Short) M64 {
	return C._mm_set1_pi16(v0)
}

// _mm_set1_pi8
// __always_inline__, __nodebug__, __target__("mmx"), __min_vector_width__(64)
func MmSet1Pi8(v0 Char) M64 {
	return C._mm_set1_pi8(v0)
}

// _mm_setr_pi32
// __always_inline__, __nodebug__, __target__("mmx"), __min_vector_width__(64)
func MmSetrPi32(v0 Int, v1 Int) M64 {
	return C._mm_setr_pi32(v0, v1)
}

// _mm_setr_pi16
// __always_inline__, __nodebug__, __target__("mmx"), __min_vector_width__(64)
func MmSetrPi16(v0 Short, v1 Short, v2 Short, v3 Short) M64 {
	return C._mm_setr_pi16(v0, v1, v2, v3)
}

// _mm_setr_pi8
// __always_inline__, __nodebug__, __target__("mmx"), __min_vector_width__(64)
func MmSetrPi8(v0 Char, v1 Char, v2 Char, v3 Char, v4 Char, v5 Char, v6 Char, v7 Char) M64 {
	return C._mm_setr_pi8(v0, v1, v2, v3, v4, v5, v6, v7)
}
