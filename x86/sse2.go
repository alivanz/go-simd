package x86

/*
#cgo CFLAGS: 
#include <immintrin.h>
*/
import "C"

// _mm_add_sd
// __always_inline__, __nodebug__, __target__("sse2"), __min_vector_width__(128)
func MmAddSd(v0 M128D, v1 M128D) M128D {
	return C._mm_add_sd(v0, v1)
}

// _mm_add_pd
// __always_inline__, __nodebug__, __target__("sse2"), __min_vector_width__(128)
func MmAddPd(v0 M128D, v1 M128D) M128D {
	return C._mm_add_pd(v0, v1)
}

// _mm_sub_sd
// __always_inline__, __nodebug__, __target__("sse2"), __min_vector_width__(128)
func MmSubSd(v0 M128D, v1 M128D) M128D {
	return C._mm_sub_sd(v0, v1)
}

// _mm_sub_pd
// __always_inline__, __nodebug__, __target__("sse2"), __min_vector_width__(128)
func MmSubPd(v0 M128D, v1 M128D) M128D {
	return C._mm_sub_pd(v0, v1)
}

// _mm_mul_sd
// __always_inline__, __nodebug__, __target__("sse2"), __min_vector_width__(128)
func MmMulSd(v0 M128D, v1 M128D) M128D {
	return C._mm_mul_sd(v0, v1)
}

// _mm_mul_pd
// __always_inline__, __nodebug__, __target__("sse2"), __min_vector_width__(128)
func MmMulPd(v0 M128D, v1 M128D) M128D {
	return C._mm_mul_pd(v0, v1)
}

// _mm_div_sd
// __always_inline__, __nodebug__, __target__("sse2"), __min_vector_width__(128)
func MmDivSd(v0 M128D, v1 M128D) M128D {
	return C._mm_div_sd(v0, v1)
}

// _mm_div_pd
// __always_inline__, __nodebug__, __target__("sse2"), __min_vector_width__(128)
func MmDivPd(v0 M128D, v1 M128D) M128D {
	return C._mm_div_pd(v0, v1)
}

// _mm_sqrt_sd
// __always_inline__, __nodebug__, __target__("sse2"), __min_vector_width__(128)
func MmSqrtSd(v0 M128D, v1 M128D) M128D {
	return C._mm_sqrt_sd(v0, v1)
}

// _mm_sqrt_pd
// __always_inline__, __nodebug__, __target__("sse2"), __min_vector_width__(128)
func MmSqrtPd(v0 M128D) M128D {
	return C._mm_sqrt_pd(v0)
}

// _mm_min_sd
// __always_inline__, __nodebug__, __target__("sse2"), __min_vector_width__(128)
func MmMinSd(v0 M128D, v1 M128D) M128D {
	return C._mm_min_sd(v0, v1)
}

// _mm_min_pd
// __always_inline__, __nodebug__, __target__("sse2"), __min_vector_width__(128)
func MmMinPd(v0 M128D, v1 M128D) M128D {
	return C._mm_min_pd(v0, v1)
}

// _mm_max_sd
// __always_inline__, __nodebug__, __target__("sse2"), __min_vector_width__(128)
func MmMaxSd(v0 M128D, v1 M128D) M128D {
	return C._mm_max_sd(v0, v1)
}

// _mm_max_pd
// __always_inline__, __nodebug__, __target__("sse2"), __min_vector_width__(128)
func MmMaxPd(v0 M128D, v1 M128D) M128D {
	return C._mm_max_pd(v0, v1)
}

// _mm_and_pd
// __always_inline__, __nodebug__, __target__("sse2"), __min_vector_width__(128)
func MmAndPd(v0 M128D, v1 M128D) M128D {
	return C._mm_and_pd(v0, v1)
}

// _mm_andnot_pd
// __always_inline__, __nodebug__, __target__("sse2"), __min_vector_width__(128)
func MmAndnotPd(v0 M128D, v1 M128D) M128D {
	return C._mm_andnot_pd(v0, v1)
}

// _mm_or_pd
// __always_inline__, __nodebug__, __target__("sse2"), __min_vector_width__(128)
func MmOrPd(v0 M128D, v1 M128D) M128D {
	return C._mm_or_pd(v0, v1)
}

// _mm_xor_pd
// __always_inline__, __nodebug__, __target__("sse2"), __min_vector_width__(128)
func MmXorPd(v0 M128D, v1 M128D) M128D {
	return C._mm_xor_pd(v0, v1)
}

// _mm_cmpeq_pd
// __always_inline__, __nodebug__, __target__("sse2"), __min_vector_width__(128)
func MmCmpeqPd(v0 M128D, v1 M128D) M128D {
	return C._mm_cmpeq_pd(v0, v1)
}

// _mm_cmplt_pd
// __always_inline__, __nodebug__, __target__("sse2"), __min_vector_width__(128)
func MmCmpltPd(v0 M128D, v1 M128D) M128D {
	return C._mm_cmplt_pd(v0, v1)
}

// _mm_cmple_pd
// __always_inline__, __nodebug__, __target__("sse2"), __min_vector_width__(128)
func MmCmplePd(v0 M128D, v1 M128D) M128D {
	return C._mm_cmple_pd(v0, v1)
}

// _mm_cmpgt_pd
// __always_inline__, __nodebug__, __target__("sse2"), __min_vector_width__(128)
func MmCmpgtPd(v0 M128D, v1 M128D) M128D {
	return C._mm_cmpgt_pd(v0, v1)
}

// _mm_cmpge_pd
// __always_inline__, __nodebug__, __target__("sse2"), __min_vector_width__(128)
func MmCmpgePd(v0 M128D, v1 M128D) M128D {
	return C._mm_cmpge_pd(v0, v1)
}

// _mm_cmpord_pd
// __always_inline__, __nodebug__, __target__("sse2"), __min_vector_width__(128)
func MmCmpordPd(v0 M128D, v1 M128D) M128D {
	return C._mm_cmpord_pd(v0, v1)
}

// _mm_cmpunord_pd
// __always_inline__, __nodebug__, __target__("sse2"), __min_vector_width__(128)
func MmCmpunordPd(v0 M128D, v1 M128D) M128D {
	return C._mm_cmpunord_pd(v0, v1)
}

// _mm_cmpneq_pd
// __always_inline__, __nodebug__, __target__("sse2"), __min_vector_width__(128)
func MmCmpneqPd(v0 M128D, v1 M128D) M128D {
	return C._mm_cmpneq_pd(v0, v1)
}

// _mm_cmpnlt_pd
// __always_inline__, __nodebug__, __target__("sse2"), __min_vector_width__(128)
func MmCmpnltPd(v0 M128D, v1 M128D) M128D {
	return C._mm_cmpnlt_pd(v0, v1)
}

// _mm_cmpnle_pd
// __always_inline__, __nodebug__, __target__("sse2"), __min_vector_width__(128)
func MmCmpnlePd(v0 M128D, v1 M128D) M128D {
	return C._mm_cmpnle_pd(v0, v1)
}

// _mm_cmpngt_pd
// __always_inline__, __nodebug__, __target__("sse2"), __min_vector_width__(128)
func MmCmpngtPd(v0 M128D, v1 M128D) M128D {
	return C._mm_cmpngt_pd(v0, v1)
}

// _mm_cmpnge_pd
// __always_inline__, __nodebug__, __target__("sse2"), __min_vector_width__(128)
func MmCmpngePd(v0 M128D, v1 M128D) M128D {
	return C._mm_cmpnge_pd(v0, v1)
}

// _mm_cmpeq_sd
// __always_inline__, __nodebug__, __target__("sse2"), __min_vector_width__(128)
func MmCmpeqSd(v0 M128D, v1 M128D) M128D {
	return C._mm_cmpeq_sd(v0, v1)
}

// _mm_cmplt_sd
// __always_inline__, __nodebug__, __target__("sse2"), __min_vector_width__(128)
func MmCmpltSd(v0 M128D, v1 M128D) M128D {
	return C._mm_cmplt_sd(v0, v1)
}

// _mm_cmple_sd
// __always_inline__, __nodebug__, __target__("sse2"), __min_vector_width__(128)
func MmCmpleSd(v0 M128D, v1 M128D) M128D {
	return C._mm_cmple_sd(v0, v1)
}

// _mm_cmpgt_sd
// __always_inline__, __nodebug__, __target__("sse2"), __min_vector_width__(128)
func MmCmpgtSd(v0 M128D, v1 M128D) M128D {
	return C._mm_cmpgt_sd(v0, v1)
}

// _mm_cmpge_sd
// __always_inline__, __nodebug__, __target__("sse2"), __min_vector_width__(128)
func MmCmpgeSd(v0 M128D, v1 M128D) M128D {
	return C._mm_cmpge_sd(v0, v1)
}

// _mm_cmpord_sd
// __always_inline__, __nodebug__, __target__("sse2"), __min_vector_width__(128)
func MmCmpordSd(v0 M128D, v1 M128D) M128D {
	return C._mm_cmpord_sd(v0, v1)
}

// _mm_cmpunord_sd
// __always_inline__, __nodebug__, __target__("sse2"), __min_vector_width__(128)
func MmCmpunordSd(v0 M128D, v1 M128D) M128D {
	return C._mm_cmpunord_sd(v0, v1)
}

// _mm_cmpneq_sd
// __always_inline__, __nodebug__, __target__("sse2"), __min_vector_width__(128)
func MmCmpneqSd(v0 M128D, v1 M128D) M128D {
	return C._mm_cmpneq_sd(v0, v1)
}

// _mm_cmpnlt_sd
// __always_inline__, __nodebug__, __target__("sse2"), __min_vector_width__(128)
func MmCmpnltSd(v0 M128D, v1 M128D) M128D {
	return C._mm_cmpnlt_sd(v0, v1)
}

// _mm_cmpnle_sd
// __always_inline__, __nodebug__, __target__("sse2"), __min_vector_width__(128)
func MmCmpnleSd(v0 M128D, v1 M128D) M128D {
	return C._mm_cmpnle_sd(v0, v1)
}

// _mm_cmpngt_sd
// __always_inline__, __nodebug__, __target__("sse2"), __min_vector_width__(128)
func MmCmpngtSd(v0 M128D, v1 M128D) M128D {
	return C._mm_cmpngt_sd(v0, v1)
}

// _mm_cmpnge_sd
// __always_inline__, __nodebug__, __target__("sse2"), __min_vector_width__(128)
func MmCmpngeSd(v0 M128D, v1 M128D) M128D {
	return C._mm_cmpnge_sd(v0, v1)
}

// _mm_comieq_sd
// __always_inline__, __nodebug__, __target__("sse2"), __min_vector_width__(128)
func MmComieqSd(v0 M128D, v1 M128D) Int {
	return C._mm_comieq_sd(v0, v1)
}

// _mm_comilt_sd
// __always_inline__, __nodebug__, __target__("sse2"), __min_vector_width__(128)
func MmComiltSd(v0 M128D, v1 M128D) Int {
	return C._mm_comilt_sd(v0, v1)
}

// _mm_comile_sd
// __always_inline__, __nodebug__, __target__("sse2"), __min_vector_width__(128)
func MmComileSd(v0 M128D, v1 M128D) Int {
	return C._mm_comile_sd(v0, v1)
}

// _mm_comigt_sd
// __always_inline__, __nodebug__, __target__("sse2"), __min_vector_width__(128)
func MmComigtSd(v0 M128D, v1 M128D) Int {
	return C._mm_comigt_sd(v0, v1)
}

// _mm_comige_sd
// __always_inline__, __nodebug__, __target__("sse2"), __min_vector_width__(128)
func MmComigeSd(v0 M128D, v1 M128D) Int {
	return C._mm_comige_sd(v0, v1)
}

// _mm_comineq_sd
// __always_inline__, __nodebug__, __target__("sse2"), __min_vector_width__(128)
func MmComineqSd(v0 M128D, v1 M128D) Int {
	return C._mm_comineq_sd(v0, v1)
}

// _mm_ucomieq_sd
// __always_inline__, __nodebug__, __target__("sse2"), __min_vector_width__(128)
func MmUcomieqSd(v0 M128D, v1 M128D) Int {
	return C._mm_ucomieq_sd(v0, v1)
}

// _mm_ucomilt_sd
// __always_inline__, __nodebug__, __target__("sse2"), __min_vector_width__(128)
func MmUcomiltSd(v0 M128D, v1 M128D) Int {
	return C._mm_ucomilt_sd(v0, v1)
}

// _mm_ucomile_sd
// __always_inline__, __nodebug__, __target__("sse2"), __min_vector_width__(128)
func MmUcomileSd(v0 M128D, v1 M128D) Int {
	return C._mm_ucomile_sd(v0, v1)
}

// _mm_ucomigt_sd
// __always_inline__, __nodebug__, __target__("sse2"), __min_vector_width__(128)
func MmUcomigtSd(v0 M128D, v1 M128D) Int {
	return C._mm_ucomigt_sd(v0, v1)
}

// _mm_ucomige_sd
// __always_inline__, __nodebug__, __target__("sse2"), __min_vector_width__(128)
func MmUcomigeSd(v0 M128D, v1 M128D) Int {
	return C._mm_ucomige_sd(v0, v1)
}

// _mm_ucomineq_sd
// __always_inline__, __nodebug__, __target__("sse2"), __min_vector_width__(128)
func MmUcomineqSd(v0 M128D, v1 M128D) Int {
	return C._mm_ucomineq_sd(v0, v1)
}

// _mm_cvtpd_ps
// __always_inline__, __nodebug__, __target__("sse2"), __min_vector_width__(128)
func MmCvtpdPs(v0 M128D) M128 {
	return C._mm_cvtpd_ps(v0)
}

// _mm_cvtps_pd
// __always_inline__, __nodebug__, __target__("sse2"), __min_vector_width__(128)
func MmCvtpsPd(v0 M128) M128D {
	return C._mm_cvtps_pd(v0)
}

// _mm_cvtepi32_pd
// __always_inline__, __nodebug__, __target__("sse2"), __min_vector_width__(128)
func MmCvtepi32Pd(v0 M128I) M128D {
	return C._mm_cvtepi32_pd(v0)
}

// _mm_cvtpd_epi32
// __always_inline__, __nodebug__, __target__("sse2"), __min_vector_width__(128)
func MmCvtpdEpi32(v0 M128D) M128I {
	return C._mm_cvtpd_epi32(v0)
}

// _mm_cvtsd_si32
// __always_inline__, __nodebug__, __target__("sse2"), __min_vector_width__(128)
func MmCvtsdSi32(v0 M128D) Int {
	return C._mm_cvtsd_si32(v0)
}

// _mm_cvtsd_ss
// __always_inline__, __nodebug__, __target__("sse2"), __min_vector_width__(128)
func MmCvtsdSs(v0 M128, v1 M128D) M128 {
	return C._mm_cvtsd_ss(v0, v1)
}

// _mm_cvtsi32_sd
// __always_inline__, __nodebug__, __target__("sse2"), __min_vector_width__(128)
func MmCvtsi32Sd(v0 M128D, v1 Int) M128D {
	return C._mm_cvtsi32_sd(v0, v1)
}

// _mm_cvtss_sd
// __always_inline__, __nodebug__, __target__("sse2"), __min_vector_width__(128)
func MmCvtssSd(v0 M128D, v1 M128) M128D {
	return C._mm_cvtss_sd(v0, v1)
}

// _mm_cvttpd_epi32
// __always_inline__, __nodebug__, __target__("sse2"), __min_vector_width__(128)
func MmCvttpdEpi32(v0 M128D) M128I {
	return C._mm_cvttpd_epi32(v0)
}

// _mm_cvttsd_si32
// __always_inline__, __nodebug__, __target__("sse2"), __min_vector_width__(128)
func MmCvttsdSi32(v0 M128D) Int {
	return C._mm_cvttsd_si32(v0)
}

// _mm_cvtsd_f64
// __always_inline__, __nodebug__, __target__("sse2"), __min_vector_width__(128)
func MmCvtsdF64(v0 M128D) Double {
	return C._mm_cvtsd_f64(v0)
}

// _mm_undefined_pd
// __always_inline__, __nodebug__, __target__("sse2"), __min_vector_width__(128)
func MmUndefinedPd() M128D {
	return C._mm_undefined_pd()
}

// _mm_set_sd
// __always_inline__, __nodebug__, __target__("sse2"), __min_vector_width__(128)
func MmSetSd(v0 Double) M128D {
	return C._mm_set_sd(v0)
}

// _mm_set1_pd
// __always_inline__, __nodebug__, __target__("sse2"), __min_vector_width__(128)
func MmSet1Pd(v0 Double) M128D {
	return C._mm_set1_pd(v0)
}

// _mm_set_pd1
// __always_inline__, __nodebug__, __target__("sse2"), __min_vector_width__(128)
func MmSetPd1(v0 Double) M128D {
	return C._mm_set_pd1(v0)
}

// _mm_set_pd
// __always_inline__, __nodebug__, __target__("sse2"), __min_vector_width__(128)
func MmSetPd(v0 Double, v1 Double) M128D {
	return C._mm_set_pd(v0, v1)
}

// _mm_setr_pd
// __always_inline__, __nodebug__, __target__("sse2"), __min_vector_width__(128)
func MmSetrPd(v0 Double, v1 Double) M128D {
	return C._mm_setr_pd(v0, v1)
}

// _mm_setzero_pd
// __always_inline__, __nodebug__, __target__("sse2"), __min_vector_width__(128)
func MmSetzeroPd() M128D {
	return C._mm_setzero_pd()
}

// _mm_move_sd
// __always_inline__, __nodebug__, __target__("sse2"), __min_vector_width__(128)
func MmMoveSd(v0 M128D, v1 M128D) M128D {
	return C._mm_move_sd(v0, v1)
}

// _mm_add_epi8
// __always_inline__, __nodebug__, __target__("sse2"), __min_vector_width__(128)
func MmAddEpi8(v0 M128I, v1 M128I) M128I {
	return C._mm_add_epi8(v0, v1)
}

// _mm_add_epi16
// __always_inline__, __nodebug__, __target__("sse2"), __min_vector_width__(128)
func MmAddEpi16(v0 M128I, v1 M128I) M128I {
	return C._mm_add_epi16(v0, v1)
}

// _mm_add_epi32
// __always_inline__, __nodebug__, __target__("sse2"), __min_vector_width__(128)
func MmAddEpi32(v0 M128I, v1 M128I) M128I {
	return C._mm_add_epi32(v0, v1)
}

// _mm_add_epi64
// __always_inline__, __nodebug__, __target__("sse2"), __min_vector_width__(128)
func MmAddEpi64(v0 M128I, v1 M128I) M128I {
	return C._mm_add_epi64(v0, v1)
}

// _mm_adds_epi8
// __always_inline__, __nodebug__, __target__("sse2"), __min_vector_width__(128)
func MmAddsEpi8(v0 M128I, v1 M128I) M128I {
	return C._mm_adds_epi8(v0, v1)
}

// _mm_adds_epi16
// __always_inline__, __nodebug__, __target__("sse2"), __min_vector_width__(128)
func MmAddsEpi16(v0 M128I, v1 M128I) M128I {
	return C._mm_adds_epi16(v0, v1)
}

// _mm_adds_epu8
// __always_inline__, __nodebug__, __target__("sse2"), __min_vector_width__(128)
func MmAddsEpu8(v0 M128I, v1 M128I) M128I {
	return C._mm_adds_epu8(v0, v1)
}

// _mm_adds_epu16
// __always_inline__, __nodebug__, __target__("sse2"), __min_vector_width__(128)
func MmAddsEpu16(v0 M128I, v1 M128I) M128I {
	return C._mm_adds_epu16(v0, v1)
}

// _mm_avg_epu8
// __always_inline__, __nodebug__, __target__("sse2"), __min_vector_width__(128)
func MmAvgEpu8(v0 M128I, v1 M128I) M128I {
	return C._mm_avg_epu8(v0, v1)
}

// _mm_avg_epu16
// __always_inline__, __nodebug__, __target__("sse2"), __min_vector_width__(128)
func MmAvgEpu16(v0 M128I, v1 M128I) M128I {
	return C._mm_avg_epu16(v0, v1)
}

// _mm_madd_epi16
// __always_inline__, __nodebug__, __target__("sse2"), __min_vector_width__(128)
func MmMaddEpi16(v0 M128I, v1 M128I) M128I {
	return C._mm_madd_epi16(v0, v1)
}

// _mm_max_epi16
// __always_inline__, __nodebug__, __target__("sse2"), __min_vector_width__(128)
func MmMaxEpi16(v0 M128I, v1 M128I) M128I {
	return C._mm_max_epi16(v0, v1)
}

// _mm_max_epu8
// __always_inline__, __nodebug__, __target__("sse2"), __min_vector_width__(128)
func MmMaxEpu8(v0 M128I, v1 M128I) M128I {
	return C._mm_max_epu8(v0, v1)
}

// _mm_min_epi16
// __always_inline__, __nodebug__, __target__("sse2"), __min_vector_width__(128)
func MmMinEpi16(v0 M128I, v1 M128I) M128I {
	return C._mm_min_epi16(v0, v1)
}

// _mm_min_epu8
// __always_inline__, __nodebug__, __target__("sse2"), __min_vector_width__(128)
func MmMinEpu8(v0 M128I, v1 M128I) M128I {
	return C._mm_min_epu8(v0, v1)
}

// _mm_mulhi_epi16
// __always_inline__, __nodebug__, __target__("sse2"), __min_vector_width__(128)
func MmMulhiEpi16(v0 M128I, v1 M128I) M128I {
	return C._mm_mulhi_epi16(v0, v1)
}

// _mm_mulhi_epu16
// __always_inline__, __nodebug__, __target__("sse2"), __min_vector_width__(128)
func MmMulhiEpu16(v0 M128I, v1 M128I) M128I {
	return C._mm_mulhi_epu16(v0, v1)
}

// _mm_mullo_epi16
// __always_inline__, __nodebug__, __target__("sse2"), __min_vector_width__(128)
func MmMulloEpi16(v0 M128I, v1 M128I) M128I {
	return C._mm_mullo_epi16(v0, v1)
}

// _mm_mul_epu32
// __always_inline__, __nodebug__, __target__("sse2"), __min_vector_width__(128)
func MmMulEpu32(v0 M128I, v1 M128I) M128I {
	return C._mm_mul_epu32(v0, v1)
}

// _mm_sad_epu8
// __always_inline__, __nodebug__, __target__("sse2"), __min_vector_width__(128)
func MmSadEpu8(v0 M128I, v1 M128I) M128I {
	return C._mm_sad_epu8(v0, v1)
}

// _mm_sub_epi8
// __always_inline__, __nodebug__, __target__("sse2"), __min_vector_width__(128)
func MmSubEpi8(v0 M128I, v1 M128I) M128I {
	return C._mm_sub_epi8(v0, v1)
}

// _mm_sub_epi16
// __always_inline__, __nodebug__, __target__("sse2"), __min_vector_width__(128)
func MmSubEpi16(v0 M128I, v1 M128I) M128I {
	return C._mm_sub_epi16(v0, v1)
}

// _mm_sub_epi32
// __always_inline__, __nodebug__, __target__("sse2"), __min_vector_width__(128)
func MmSubEpi32(v0 M128I, v1 M128I) M128I {
	return C._mm_sub_epi32(v0, v1)
}

// _mm_sub_epi64
// __always_inline__, __nodebug__, __target__("sse2"), __min_vector_width__(128)
func MmSubEpi64(v0 M128I, v1 M128I) M128I {
	return C._mm_sub_epi64(v0, v1)
}

// _mm_subs_epi8
// __always_inline__, __nodebug__, __target__("sse2"), __min_vector_width__(128)
func MmSubsEpi8(v0 M128I, v1 M128I) M128I {
	return C._mm_subs_epi8(v0, v1)
}

// _mm_subs_epi16
// __always_inline__, __nodebug__, __target__("sse2"), __min_vector_width__(128)
func MmSubsEpi16(v0 M128I, v1 M128I) M128I {
	return C._mm_subs_epi16(v0, v1)
}

// _mm_subs_epu8
// __always_inline__, __nodebug__, __target__("sse2"), __min_vector_width__(128)
func MmSubsEpu8(v0 M128I, v1 M128I) M128I {
	return C._mm_subs_epu8(v0, v1)
}

// _mm_subs_epu16
// __always_inline__, __nodebug__, __target__("sse2"), __min_vector_width__(128)
func MmSubsEpu16(v0 M128I, v1 M128I) M128I {
	return C._mm_subs_epu16(v0, v1)
}

// _mm_and_si128
// __always_inline__, __nodebug__, __target__("sse2"), __min_vector_width__(128)
func MmAndSi128(v0 M128I, v1 M128I) M128I {
	return C._mm_and_si128(v0, v1)
}

// _mm_andnot_si128
// __always_inline__, __nodebug__, __target__("sse2"), __min_vector_width__(128)
func MmAndnotSi128(v0 M128I, v1 M128I) M128I {
	return C._mm_andnot_si128(v0, v1)
}

// _mm_or_si128
// __always_inline__, __nodebug__, __target__("sse2"), __min_vector_width__(128)
func MmOrSi128(v0 M128I, v1 M128I) M128I {
	return C._mm_or_si128(v0, v1)
}

// _mm_xor_si128
// __always_inline__, __nodebug__, __target__("sse2"), __min_vector_width__(128)
func MmXorSi128(v0 M128I, v1 M128I) M128I {
	return C._mm_xor_si128(v0, v1)
}

// _mm_slli_epi16
// __always_inline__, __nodebug__, __target__("sse2"), __min_vector_width__(128)
func MmSlliEpi16(v0 M128I, v1 Int) M128I {
	return C._mm_slli_epi16(v0, v1)
}

// _mm_sll_epi16
// __always_inline__, __nodebug__, __target__("sse2"), __min_vector_width__(128)
func MmSllEpi16(v0 M128I, v1 M128I) M128I {
	return C._mm_sll_epi16(v0, v1)
}

// _mm_slli_epi32
// __always_inline__, __nodebug__, __target__("sse2"), __min_vector_width__(128)
func MmSlliEpi32(v0 M128I, v1 Int) M128I {
	return C._mm_slli_epi32(v0, v1)
}

// _mm_sll_epi32
// __always_inline__, __nodebug__, __target__("sse2"), __min_vector_width__(128)
func MmSllEpi32(v0 M128I, v1 M128I) M128I {
	return C._mm_sll_epi32(v0, v1)
}

// _mm_slli_epi64
// __always_inline__, __nodebug__, __target__("sse2"), __min_vector_width__(128)
func MmSlliEpi64(v0 M128I, v1 Int) M128I {
	return C._mm_slli_epi64(v0, v1)
}

// _mm_sll_epi64
// __always_inline__, __nodebug__, __target__("sse2"), __min_vector_width__(128)
func MmSllEpi64(v0 M128I, v1 M128I) M128I {
	return C._mm_sll_epi64(v0, v1)
}

// _mm_srai_epi16
// __always_inline__, __nodebug__, __target__("sse2"), __min_vector_width__(128)
func MmSraiEpi16(v0 M128I, v1 Int) M128I {
	return C._mm_srai_epi16(v0, v1)
}

// _mm_sra_epi16
// __always_inline__, __nodebug__, __target__("sse2"), __min_vector_width__(128)
func MmSraEpi16(v0 M128I, v1 M128I) M128I {
	return C._mm_sra_epi16(v0, v1)
}

// _mm_srai_epi32
// __always_inline__, __nodebug__, __target__("sse2"), __min_vector_width__(128)
func MmSraiEpi32(v0 M128I, v1 Int) M128I {
	return C._mm_srai_epi32(v0, v1)
}

// _mm_sra_epi32
// __always_inline__, __nodebug__, __target__("sse2"), __min_vector_width__(128)
func MmSraEpi32(v0 M128I, v1 M128I) M128I {
	return C._mm_sra_epi32(v0, v1)
}

// _mm_srli_epi16
// __always_inline__, __nodebug__, __target__("sse2"), __min_vector_width__(128)
func MmSrliEpi16(v0 M128I, v1 Int) M128I {
	return C._mm_srli_epi16(v0, v1)
}

// _mm_srl_epi16
// __always_inline__, __nodebug__, __target__("sse2"), __min_vector_width__(128)
func MmSrlEpi16(v0 M128I, v1 M128I) M128I {
	return C._mm_srl_epi16(v0, v1)
}

// _mm_srli_epi32
// __always_inline__, __nodebug__, __target__("sse2"), __min_vector_width__(128)
func MmSrliEpi32(v0 M128I, v1 Int) M128I {
	return C._mm_srli_epi32(v0, v1)
}

// _mm_srl_epi32
// __always_inline__, __nodebug__, __target__("sse2"), __min_vector_width__(128)
func MmSrlEpi32(v0 M128I, v1 M128I) M128I {
	return C._mm_srl_epi32(v0, v1)
}

// _mm_srli_epi64
// __always_inline__, __nodebug__, __target__("sse2"), __min_vector_width__(128)
func MmSrliEpi64(v0 M128I, v1 Int) M128I {
	return C._mm_srli_epi64(v0, v1)
}

// _mm_srl_epi64
// __always_inline__, __nodebug__, __target__("sse2"), __min_vector_width__(128)
func MmSrlEpi64(v0 M128I, v1 M128I) M128I {
	return C._mm_srl_epi64(v0, v1)
}

// _mm_cmpeq_epi8
// __always_inline__, __nodebug__, __target__("sse2"), __min_vector_width__(128)
func MmCmpeqEpi8(v0 M128I, v1 M128I) M128I {
	return C._mm_cmpeq_epi8(v0, v1)
}

// _mm_cmpeq_epi16
// __always_inline__, __nodebug__, __target__("sse2"), __min_vector_width__(128)
func MmCmpeqEpi16(v0 M128I, v1 M128I) M128I {
	return C._mm_cmpeq_epi16(v0, v1)
}

// _mm_cmpeq_epi32
// __always_inline__, __nodebug__, __target__("sse2"), __min_vector_width__(128)
func MmCmpeqEpi32(v0 M128I, v1 M128I) M128I {
	return C._mm_cmpeq_epi32(v0, v1)
}

// _mm_cmpgt_epi8
// __always_inline__, __nodebug__, __target__("sse2"), __min_vector_width__(128)
func MmCmpgtEpi8(v0 M128I, v1 M128I) M128I {
	return C._mm_cmpgt_epi8(v0, v1)
}

// _mm_cmpgt_epi16
// __always_inline__, __nodebug__, __target__("sse2"), __min_vector_width__(128)
func MmCmpgtEpi16(v0 M128I, v1 M128I) M128I {
	return C._mm_cmpgt_epi16(v0, v1)
}

// _mm_cmpgt_epi32
// __always_inline__, __nodebug__, __target__("sse2"), __min_vector_width__(128)
func MmCmpgtEpi32(v0 M128I, v1 M128I) M128I {
	return C._mm_cmpgt_epi32(v0, v1)
}

// _mm_cmplt_epi8
// __always_inline__, __nodebug__, __target__("sse2"), __min_vector_width__(128)
func MmCmpltEpi8(v0 M128I, v1 M128I) M128I {
	return C._mm_cmplt_epi8(v0, v1)
}

// _mm_cmplt_epi16
// __always_inline__, __nodebug__, __target__("sse2"), __min_vector_width__(128)
func MmCmpltEpi16(v0 M128I, v1 M128I) M128I {
	return C._mm_cmplt_epi16(v0, v1)
}

// _mm_cmplt_epi32
// __always_inline__, __nodebug__, __target__("sse2"), __min_vector_width__(128)
func MmCmpltEpi32(v0 M128I, v1 M128I) M128I {
	return C._mm_cmplt_epi32(v0, v1)
}

// _mm_cvtsi64_sd
// __always_inline__, __nodebug__, __target__("sse2"), __min_vector_width__(128)
func MmCvtsi64Sd(v0 M128D, v1 Longlong) M128D {
	return C._mm_cvtsi64_sd(v0, v1)
}

// _mm_cvtsd_si64
// __always_inline__, __nodebug__, __target__("sse2"), __min_vector_width__(128)
func MmCvtsdSi64(v0 M128D) Longlong {
	return C._mm_cvtsd_si64(v0)
}

// _mm_cvttsd_si64
// __always_inline__, __nodebug__, __target__("sse2"), __min_vector_width__(128)
func MmCvttsdSi64(v0 M128D) Longlong {
	return C._mm_cvttsd_si64(v0)
}

// _mm_cvtepi32_ps
// __always_inline__, __nodebug__, __target__("sse2"), __min_vector_width__(128)
func MmCvtepi32Ps(v0 M128I) M128 {
	return C._mm_cvtepi32_ps(v0)
}

// _mm_cvtps_epi32
// __always_inline__, __nodebug__, __target__("sse2"), __min_vector_width__(128)
func MmCvtpsEpi32(v0 M128) M128I {
	return C._mm_cvtps_epi32(v0)
}

// _mm_cvttps_epi32
// __always_inline__, __nodebug__, __target__("sse2"), __min_vector_width__(128)
func MmCvttpsEpi32(v0 M128) M128I {
	return C._mm_cvttps_epi32(v0)
}

// _mm_cvtsi32_si128
// __always_inline__, __nodebug__, __target__("sse2"), __min_vector_width__(128)
func MmCvtsi32Si128(v0 Int) M128I {
	return C._mm_cvtsi32_si128(v0)
}

// _mm_cvtsi64_si128
// __always_inline__, __nodebug__, __target__("sse2"), __min_vector_width__(128)
func MmCvtsi64Si128(v0 Longlong) M128I {
	return C._mm_cvtsi64_si128(v0)
}

// _mm_cvtsi128_si32
// __always_inline__, __nodebug__, __target__("sse2"), __min_vector_width__(128)
func MmCvtsi128Si32(v0 M128I) Int {
	return C._mm_cvtsi128_si32(v0)
}

// _mm_cvtsi128_si64
// __always_inline__, __nodebug__, __target__("sse2"), __min_vector_width__(128)
func MmCvtsi128Si64(v0 M128I) Longlong {
	return C._mm_cvtsi128_si64(v0)
}

// _mm_undefined_si128
// __always_inline__, __nodebug__, __target__("sse2"), __min_vector_width__(128)
func MmUndefinedSi128() M128I {
	return C._mm_undefined_si128()
}

// _mm_set_epi64x
// __always_inline__, __nodebug__, __target__("sse2"), __min_vector_width__(128)
func MmSetEpi64X(v0 Longlong, v1 Longlong) M128I {
	return C._mm_set_epi64x(v0, v1)
}

// _mm_set_epi64
// __always_inline__, __nodebug__, __target__("sse2"), __min_vector_width__(128)
func MmSetEpi64(v0 M64, v1 M64) M128I {
	return C._mm_set_epi64(v0, v1)
}

// _mm_set_epi32
// __always_inline__, __nodebug__, __target__("sse2"), __min_vector_width__(128)
func MmSetEpi32(v0 Int, v1 Int, v2 Int, v3 Int) M128I {
	return C._mm_set_epi32(v0, v1, v2, v3)
}

// _mm_set_epi16
// __always_inline__, __nodebug__, __target__("sse2"), __min_vector_width__(128)
func MmSetEpi16(v0 Short, v1 Short, v2 Short, v3 Short, v4 Short, v5 Short, v6 Short, v7 Short) M128I {
	return C._mm_set_epi16(v0, v1, v2, v3, v4, v5, v6, v7)
}

// _mm_set_epi8
// __always_inline__, __nodebug__, __target__("sse2"), __min_vector_width__(128)
func MmSetEpi8(v0 Char, v1 Char, v2 Char, v3 Char, v4 Char, v5 Char, v6 Char, v7 Char, v8 Char, v9 Char, v10 Char, v11 Char, v12 Char, v13 Char, v14 Char, v15 Char) M128I {
	return C._mm_set_epi8(v0, v1, v2, v3, v4, v5, v6, v7, v8, v9, v10, v11, v12, v13, v14, v15)
}

// _mm_set1_epi64x
// __always_inline__, __nodebug__, __target__("sse2"), __min_vector_width__(128)
func MmSet1Epi64X(v0 Longlong) M128I {
	return C._mm_set1_epi64x(v0)
}

// _mm_set1_epi64
// __always_inline__, __nodebug__, __target__("sse2"), __min_vector_width__(128)
func MmSet1Epi64(v0 M64) M128I {
	return C._mm_set1_epi64(v0)
}

// _mm_set1_epi32
// __always_inline__, __nodebug__, __target__("sse2"), __min_vector_width__(128)
func MmSet1Epi32(v0 Int) M128I {
	return C._mm_set1_epi32(v0)
}

// _mm_set1_epi16
// __always_inline__, __nodebug__, __target__("sse2"), __min_vector_width__(128)
func MmSet1Epi16(v0 Short) M128I {
	return C._mm_set1_epi16(v0)
}

// _mm_set1_epi8
// __always_inline__, __nodebug__, __target__("sse2"), __min_vector_width__(128)
func MmSet1Epi8(v0 Char) M128I {
	return C._mm_set1_epi8(v0)
}

// _mm_setr_epi64
// __always_inline__, __nodebug__, __target__("sse2"), __min_vector_width__(128)
func MmSetrEpi64(v0 M64, v1 M64) M128I {
	return C._mm_setr_epi64(v0, v1)
}

// _mm_setr_epi32
// __always_inline__, __nodebug__, __target__("sse2"), __min_vector_width__(128)
func MmSetrEpi32(v0 Int, v1 Int, v2 Int, v3 Int) M128I {
	return C._mm_setr_epi32(v0, v1, v2, v3)
}

// _mm_setr_epi16
// __always_inline__, __nodebug__, __target__("sse2"), __min_vector_width__(128)
func MmSetrEpi16(v0 Short, v1 Short, v2 Short, v3 Short, v4 Short, v5 Short, v6 Short, v7 Short) M128I {
	return C._mm_setr_epi16(v0, v1, v2, v3, v4, v5, v6, v7)
}

// _mm_setr_epi8
// __always_inline__, __nodebug__, __target__("sse2"), __min_vector_width__(128)
func MmSetrEpi8(v0 Char, v1 Char, v2 Char, v3 Char, v4 Char, v5 Char, v6 Char, v7 Char, v8 Char, v9 Char, v10 Char, v11 Char, v12 Char, v13 Char, v14 Char, v15 Char) M128I {
	return C._mm_setr_epi8(v0, v1, v2, v3, v4, v5, v6, v7, v8, v9, v10, v11, v12, v13, v14, v15)
}

// _mm_setzero_si128
// __always_inline__, __nodebug__, __target__("sse2"), __min_vector_width__(128)
func MmSetzeroSi128() M128I {
	return C._mm_setzero_si128()
}

// _mm_packs_epi16
// __always_inline__, __nodebug__, __target__("sse2"), __min_vector_width__(128)
func MmPacksEpi16(v0 M128I, v1 M128I) M128I {
	return C._mm_packs_epi16(v0, v1)
}

// _mm_packs_epi32
// __always_inline__, __nodebug__, __target__("sse2"), __min_vector_width__(128)
func MmPacksEpi32(v0 M128I, v1 M128I) M128I {
	return C._mm_packs_epi32(v0, v1)
}

// _mm_packus_epi16
// __always_inline__, __nodebug__, __target__("sse2"), __min_vector_width__(128)
func MmPackusEpi16(v0 M128I, v1 M128I) M128I {
	return C._mm_packus_epi16(v0, v1)
}

// _mm_movemask_epi8
// __always_inline__, __nodebug__, __target__("sse2"), __min_vector_width__(128)
func MmMovemaskEpi8(v0 M128I) Int {
	return C._mm_movemask_epi8(v0)
}

// _mm_unpackhi_epi8
// __always_inline__, __nodebug__, __target__("sse2"), __min_vector_width__(128)
func MmUnpackhiEpi8(v0 M128I, v1 M128I) M128I {
	return C._mm_unpackhi_epi8(v0, v1)
}

// _mm_unpackhi_epi16
// __always_inline__, __nodebug__, __target__("sse2"), __min_vector_width__(128)
func MmUnpackhiEpi16(v0 M128I, v1 M128I) M128I {
	return C._mm_unpackhi_epi16(v0, v1)
}

// _mm_unpackhi_epi32
// __always_inline__, __nodebug__, __target__("sse2"), __min_vector_width__(128)
func MmUnpackhiEpi32(v0 M128I, v1 M128I) M128I {
	return C._mm_unpackhi_epi32(v0, v1)
}

// _mm_unpackhi_epi64
// __always_inline__, __nodebug__, __target__("sse2"), __min_vector_width__(128)
func MmUnpackhiEpi64(v0 M128I, v1 M128I) M128I {
	return C._mm_unpackhi_epi64(v0, v1)
}

// _mm_unpacklo_epi8
// __always_inline__, __nodebug__, __target__("sse2"), __min_vector_width__(128)
func MmUnpackloEpi8(v0 M128I, v1 M128I) M128I {
	return C._mm_unpacklo_epi8(v0, v1)
}

// _mm_unpacklo_epi16
// __always_inline__, __nodebug__, __target__("sse2"), __min_vector_width__(128)
func MmUnpackloEpi16(v0 M128I, v1 M128I) M128I {
	return C._mm_unpacklo_epi16(v0, v1)
}

// _mm_unpacklo_epi32
// __always_inline__, __nodebug__, __target__("sse2"), __min_vector_width__(128)
func MmUnpackloEpi32(v0 M128I, v1 M128I) M128I {
	return C._mm_unpacklo_epi32(v0, v1)
}

// _mm_unpacklo_epi64
// __always_inline__, __nodebug__, __target__("sse2"), __min_vector_width__(128)
func MmUnpackloEpi64(v0 M128I, v1 M128I) M128I {
	return C._mm_unpacklo_epi64(v0, v1)
}

// _mm_movepi64_pi64
// __always_inline__, __nodebug__, __target__("sse2"), __min_vector_width__(128)
func MmMovepi64Pi64(v0 M128I) M64 {
	return C._mm_movepi64_pi64(v0)
}

// _mm_movpi64_epi64
// __always_inline__, __nodebug__, __target__("sse2"), __min_vector_width__(128)
func MmMovpi64Epi64(v0 M64) M128I {
	return C._mm_movpi64_epi64(v0)
}

// _mm_move_epi64
// __always_inline__, __nodebug__, __target__("sse2"), __min_vector_width__(128)
func MmMoveEpi64(v0 M128I) M128I {
	return C._mm_move_epi64(v0)
}

// _mm_unpackhi_pd
// __always_inline__, __nodebug__, __target__("sse2"), __min_vector_width__(128)
func MmUnpackhiPd(v0 M128D, v1 M128D) M128D {
	return C._mm_unpackhi_pd(v0, v1)
}

// _mm_unpacklo_pd
// __always_inline__, __nodebug__, __target__("sse2"), __min_vector_width__(128)
func MmUnpackloPd(v0 M128D, v1 M128D) M128D {
	return C._mm_unpacklo_pd(v0, v1)
}

// _mm_movemask_pd
// __always_inline__, __nodebug__, __target__("sse2"), __min_vector_width__(128)
func MmMovemaskPd(v0 M128D) Int {
	return C._mm_movemask_pd(v0)
}

// _mm_castpd_ps
// __always_inline__, __nodebug__, __target__("sse2"), __min_vector_width__(128)
func MmCastpdPs(v0 M128D) M128 {
	return C._mm_castpd_ps(v0)
}

// _mm_castpd_si128
// __always_inline__, __nodebug__, __target__("sse2"), __min_vector_width__(128)
func MmCastpdSi128(v0 M128D) M128I {
	return C._mm_castpd_si128(v0)
}

// _mm_castps_pd
// __always_inline__, __nodebug__, __target__("sse2"), __min_vector_width__(128)
func MmCastpsPd(v0 M128) M128D {
	return C._mm_castps_pd(v0)
}

// _mm_castps_si128
// __always_inline__, __nodebug__, __target__("sse2"), __min_vector_width__(128)
func MmCastpsSi128(v0 M128) M128I {
	return C._mm_castps_si128(v0)
}

// _mm_castsi128_ps
// __always_inline__, __nodebug__, __target__("sse2"), __min_vector_width__(128)
func MmCastsi128Ps(v0 M128I) M128 {
	return C._mm_castsi128_ps(v0)
}

// _mm_castsi128_pd
// __always_inline__, __nodebug__, __target__("sse2"), __min_vector_width__(128)
func MmCastsi128Pd(v0 M128I) M128D {
	return C._mm_castsi128_pd(v0)
}
