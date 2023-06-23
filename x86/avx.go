package x86

/*
#cgo CFLAGS: -march=native
#include <immintrin.h>
*/
import "C"

// _mm256_add_pd
// __always_inline__, __nodebug__, __target__("avx"), __min_vector_width__(256)
func Mm256AddPd(v0 M256D, v1 M256D) M256D {
	return C._mm256_add_pd(v0, v1)
}

// _mm256_add_ps
// __always_inline__, __nodebug__, __target__("avx"), __min_vector_width__(256)
func Mm256AddPs(v0 M256, v1 M256) M256 {
	return C._mm256_add_ps(v0, v1)
}

// _mm256_sub_pd
// __always_inline__, __nodebug__, __target__("avx"), __min_vector_width__(256)
func Mm256SubPd(v0 M256D, v1 M256D) M256D {
	return C._mm256_sub_pd(v0, v1)
}

// _mm256_sub_ps
// __always_inline__, __nodebug__, __target__("avx"), __min_vector_width__(256)
func Mm256SubPs(v0 M256, v1 M256) M256 {
	return C._mm256_sub_ps(v0, v1)
}

// _mm256_addsub_pd
// __always_inline__, __nodebug__, __target__("avx"), __min_vector_width__(256)
func Mm256AddsubPd(v0 M256D, v1 M256D) M256D {
	return C._mm256_addsub_pd(v0, v1)
}

// _mm256_addsub_ps
// __always_inline__, __nodebug__, __target__("avx"), __min_vector_width__(256)
func Mm256AddsubPs(v0 M256, v1 M256) M256 {
	return C._mm256_addsub_ps(v0, v1)
}

// _mm256_div_pd
// __always_inline__, __nodebug__, __target__("avx"), __min_vector_width__(256)
func Mm256DivPd(v0 M256D, v1 M256D) M256D {
	return C._mm256_div_pd(v0, v1)
}

// _mm256_div_ps
// __always_inline__, __nodebug__, __target__("avx"), __min_vector_width__(256)
func Mm256DivPs(v0 M256, v1 M256) M256 {
	return C._mm256_div_ps(v0, v1)
}

// _mm256_max_pd
// __always_inline__, __nodebug__, __target__("avx"), __min_vector_width__(256)
func Mm256MaxPd(v0 M256D, v1 M256D) M256D {
	return C._mm256_max_pd(v0, v1)
}

// _mm256_max_ps
// __always_inline__, __nodebug__, __target__("avx"), __min_vector_width__(256)
func Mm256MaxPs(v0 M256, v1 M256) M256 {
	return C._mm256_max_ps(v0, v1)
}

// _mm256_min_pd
// __always_inline__, __nodebug__, __target__("avx"), __min_vector_width__(256)
func Mm256MinPd(v0 M256D, v1 M256D) M256D {
	return C._mm256_min_pd(v0, v1)
}

// _mm256_min_ps
// __always_inline__, __nodebug__, __target__("avx"), __min_vector_width__(256)
func Mm256MinPs(v0 M256, v1 M256) M256 {
	return C._mm256_min_ps(v0, v1)
}

// _mm256_mul_pd
// __always_inline__, __nodebug__, __target__("avx"), __min_vector_width__(256)
func Mm256MulPd(v0 M256D, v1 M256D) M256D {
	return C._mm256_mul_pd(v0, v1)
}

// _mm256_mul_ps
// __always_inline__, __nodebug__, __target__("avx"), __min_vector_width__(256)
func Mm256MulPs(v0 M256, v1 M256) M256 {
	return C._mm256_mul_ps(v0, v1)
}

// _mm256_sqrt_pd
// __always_inline__, __nodebug__, __target__("avx"), __min_vector_width__(256)
func Mm256SqrtPd(v0 M256D) M256D {
	return C._mm256_sqrt_pd(v0)
}

// _mm256_sqrt_ps
// __always_inline__, __nodebug__, __target__("avx"), __min_vector_width__(256)
func Mm256SqrtPs(v0 M256) M256 {
	return C._mm256_sqrt_ps(v0)
}

// _mm256_rsqrt_ps
// __always_inline__, __nodebug__, __target__("avx"), __min_vector_width__(256)
func Mm256RsqrtPs(v0 M256) M256 {
	return C._mm256_rsqrt_ps(v0)
}

// _mm256_rcp_ps
// __always_inline__, __nodebug__, __target__("avx"), __min_vector_width__(256)
func Mm256RcpPs(v0 M256) M256 {
	return C._mm256_rcp_ps(v0)
}

// _mm256_and_pd
// __always_inline__, __nodebug__, __target__("avx"), __min_vector_width__(256)
func Mm256AndPd(v0 M256D, v1 M256D) M256D {
	return C._mm256_and_pd(v0, v1)
}

// _mm256_and_ps
// __always_inline__, __nodebug__, __target__("avx"), __min_vector_width__(256)
func Mm256AndPs(v0 M256, v1 M256) M256 {
	return C._mm256_and_ps(v0, v1)
}

// _mm256_andnot_pd
// __always_inline__, __nodebug__, __target__("avx"), __min_vector_width__(256)
func Mm256AndnotPd(v0 M256D, v1 M256D) M256D {
	return C._mm256_andnot_pd(v0, v1)
}

// _mm256_andnot_ps
// __always_inline__, __nodebug__, __target__("avx"), __min_vector_width__(256)
func Mm256AndnotPs(v0 M256, v1 M256) M256 {
	return C._mm256_andnot_ps(v0, v1)
}

// _mm256_or_pd
// __always_inline__, __nodebug__, __target__("avx"), __min_vector_width__(256)
func Mm256OrPd(v0 M256D, v1 M256D) M256D {
	return C._mm256_or_pd(v0, v1)
}

// _mm256_or_ps
// __always_inline__, __nodebug__, __target__("avx"), __min_vector_width__(256)
func Mm256OrPs(v0 M256, v1 M256) M256 {
	return C._mm256_or_ps(v0, v1)
}

// _mm256_xor_pd
// __always_inline__, __nodebug__, __target__("avx"), __min_vector_width__(256)
func Mm256XorPd(v0 M256D, v1 M256D) M256D {
	return C._mm256_xor_pd(v0, v1)
}

// _mm256_xor_ps
// __always_inline__, __nodebug__, __target__("avx"), __min_vector_width__(256)
func Mm256XorPs(v0 M256, v1 M256) M256 {
	return C._mm256_xor_ps(v0, v1)
}

// _mm256_hadd_pd
// __always_inline__, __nodebug__, __target__("avx"), __min_vector_width__(256)
func Mm256HaddPd(v0 M256D, v1 M256D) M256D {
	return C._mm256_hadd_pd(v0, v1)
}

// _mm256_hadd_ps
// __always_inline__, __nodebug__, __target__("avx"), __min_vector_width__(256)
func Mm256HaddPs(v0 M256, v1 M256) M256 {
	return C._mm256_hadd_ps(v0, v1)
}

// _mm256_hsub_pd
// __always_inline__, __nodebug__, __target__("avx"), __min_vector_width__(256)
func Mm256HsubPd(v0 M256D, v1 M256D) M256D {
	return C._mm256_hsub_pd(v0, v1)
}

// _mm256_hsub_ps
// __always_inline__, __nodebug__, __target__("avx"), __min_vector_width__(256)
func Mm256HsubPs(v0 M256, v1 M256) M256 {
	return C._mm256_hsub_ps(v0, v1)
}

// _mm_permutevar_pd
// __always_inline__, __nodebug__, __target__("avx"), __min_vector_width__(128)
func MmPermutevarPd(v0 M128D, v1 M128I) M128D {
	return C._mm_permutevar_pd(v0, v1)
}

// _mm256_permutevar_pd
// __always_inline__, __nodebug__, __target__("avx"), __min_vector_width__(256)
func Mm256PermutevarPd(v0 M256D, v1 M256I) M256D {
	return C._mm256_permutevar_pd(v0, v1)
}

// _mm_permutevar_ps
// __always_inline__, __nodebug__, __target__("avx"), __min_vector_width__(128)
func MmPermutevarPs(v0 M128, v1 M128I) M128 {
	return C._mm_permutevar_ps(v0, v1)
}

// _mm256_permutevar_ps
// __always_inline__, __nodebug__, __target__("avx"), __min_vector_width__(256)
func Mm256PermutevarPs(v0 M256, v1 M256I) M256 {
	return C._mm256_permutevar_ps(v0, v1)
}

// _mm256_blendv_pd
// __always_inline__, __nodebug__, __target__("avx"), __min_vector_width__(256)
func Mm256BlendvPd(v0 M256D, v1 M256D, v2 M256D) M256D {
	return C._mm256_blendv_pd(v0, v1, v2)
}

// _mm256_blendv_ps
// __always_inline__, __nodebug__, __target__("avx"), __min_vector_width__(256)
func Mm256BlendvPs(v0 M256, v1 M256, v2 M256) M256 {
	return C._mm256_blendv_ps(v0, v1, v2)
}

// _mm256_cvtepi32_pd
// __always_inline__, __nodebug__, __target__("avx"), __min_vector_width__(256)
func Mm256Cvtepi32Pd(v0 M128I) M256D {
	return C._mm256_cvtepi32_pd(v0)
}

// _mm256_cvtepi32_ps
// __always_inline__, __nodebug__, __target__("avx"), __min_vector_width__(256)
func Mm256Cvtepi32Ps(v0 M256I) M256 {
	return C._mm256_cvtepi32_ps(v0)
}

// _mm256_cvtpd_ps
// __always_inline__, __nodebug__, __target__("avx"), __min_vector_width__(256)
func Mm256CvtpdPs(v0 M256D) M128 {
	return C._mm256_cvtpd_ps(v0)
}

// _mm256_cvtps_epi32
// __always_inline__, __nodebug__, __target__("avx"), __min_vector_width__(256)
func Mm256CvtpsEpi32(v0 M256) M256I {
	return C._mm256_cvtps_epi32(v0)
}

// _mm256_cvtps_pd
// __always_inline__, __nodebug__, __target__("avx"), __min_vector_width__(256)
func Mm256CvtpsPd(v0 M128) M256D {
	return C._mm256_cvtps_pd(v0)
}

// _mm256_cvttpd_epi32
// __always_inline__, __nodebug__, __target__("avx"), __min_vector_width__(256)
func Mm256CvttpdEpi32(v0 M256D) M128I {
	return C._mm256_cvttpd_epi32(v0)
}

// _mm256_cvtpd_epi32
// __always_inline__, __nodebug__, __target__("avx"), __min_vector_width__(256)
func Mm256CvtpdEpi32(v0 M256D) M128I {
	return C._mm256_cvtpd_epi32(v0)
}

// _mm256_cvttps_epi32
// __always_inline__, __nodebug__, __target__("avx"), __min_vector_width__(256)
func Mm256CvttpsEpi32(v0 M256) M256I {
	return C._mm256_cvttps_epi32(v0)
}

// _mm256_cvtsd_f64
// __always_inline__, __nodebug__, __target__("avx"), __min_vector_width__(256)
func Mm256CvtsdF64(v0 M256D) Double {
	return C._mm256_cvtsd_f64(v0)
}

// _mm256_cvtsi256_si32
// __always_inline__, __nodebug__, __target__("avx"), __min_vector_width__(256)
func Mm256Cvtsi256Si32(v0 M256I) Int {
	return C._mm256_cvtsi256_si32(v0)
}

// _mm256_cvtss_f32
// __always_inline__, __nodebug__, __target__("avx"), __min_vector_width__(256)
func Mm256CvtssF32(v0 M256) Float {
	return C._mm256_cvtss_f32(v0)
}

// _mm256_movehdup_ps
// __always_inline__, __nodebug__, __target__("avx"), __min_vector_width__(256)
func Mm256MovehdupPs(v0 M256) M256 {
	return C._mm256_movehdup_ps(v0)
}

// _mm256_moveldup_ps
// __always_inline__, __nodebug__, __target__("avx"), __min_vector_width__(256)
func Mm256MoveldupPs(v0 M256) M256 {
	return C._mm256_moveldup_ps(v0)
}

// _mm256_movedup_pd
// __always_inline__, __nodebug__, __target__("avx"), __min_vector_width__(256)
func Mm256MovedupPd(v0 M256D) M256D {
	return C._mm256_movedup_pd(v0)
}

// _mm256_unpackhi_pd
// __always_inline__, __nodebug__, __target__("avx"), __min_vector_width__(256)
func Mm256UnpackhiPd(v0 M256D, v1 M256D) M256D {
	return C._mm256_unpackhi_pd(v0, v1)
}

// _mm256_unpacklo_pd
// __always_inline__, __nodebug__, __target__("avx"), __min_vector_width__(256)
func Mm256UnpackloPd(v0 M256D, v1 M256D) M256D {
	return C._mm256_unpacklo_pd(v0, v1)
}

// _mm256_unpackhi_ps
// __always_inline__, __nodebug__, __target__("avx"), __min_vector_width__(256)
func Mm256UnpackhiPs(v0 M256, v1 M256) M256 {
	return C._mm256_unpackhi_ps(v0, v1)
}

// _mm256_unpacklo_ps
// __always_inline__, __nodebug__, __target__("avx"), __min_vector_width__(256)
func Mm256UnpackloPs(v0 M256, v1 M256) M256 {
	return C._mm256_unpacklo_ps(v0, v1)
}

// _mm_testz_pd
// __always_inline__, __nodebug__, __target__("avx"), __min_vector_width__(128)
func MmTestzPd(v0 M128D, v1 M128D) Int {
	return C._mm_testz_pd(v0, v1)
}

// _mm_testc_pd
// __always_inline__, __nodebug__, __target__("avx"), __min_vector_width__(128)
func MmTestcPd(v0 M128D, v1 M128D) Int {
	return C._mm_testc_pd(v0, v1)
}

// _mm_testnzc_pd
// __always_inline__, __nodebug__, __target__("avx"), __min_vector_width__(128)
func MmTestnzcPd(v0 M128D, v1 M128D) Int {
	return C._mm_testnzc_pd(v0, v1)
}

// _mm_testz_ps
// __always_inline__, __nodebug__, __target__("avx"), __min_vector_width__(128)
func MmTestzPs(v0 M128, v1 M128) Int {
	return C._mm_testz_ps(v0, v1)
}

// _mm_testc_ps
// __always_inline__, __nodebug__, __target__("avx"), __min_vector_width__(128)
func MmTestcPs(v0 M128, v1 M128) Int {
	return C._mm_testc_ps(v0, v1)
}

// _mm_testnzc_ps
// __always_inline__, __nodebug__, __target__("avx"), __min_vector_width__(128)
func MmTestnzcPs(v0 M128, v1 M128) Int {
	return C._mm_testnzc_ps(v0, v1)
}

// _mm256_testz_pd
// __always_inline__, __nodebug__, __target__("avx"), __min_vector_width__(256)
func Mm256TestzPd(v0 M256D, v1 M256D) Int {
	return C._mm256_testz_pd(v0, v1)
}

// _mm256_testc_pd
// __always_inline__, __nodebug__, __target__("avx"), __min_vector_width__(256)
func Mm256TestcPd(v0 M256D, v1 M256D) Int {
	return C._mm256_testc_pd(v0, v1)
}

// _mm256_testnzc_pd
// __always_inline__, __nodebug__, __target__("avx"), __min_vector_width__(256)
func Mm256TestnzcPd(v0 M256D, v1 M256D) Int {
	return C._mm256_testnzc_pd(v0, v1)
}

// _mm256_testz_ps
// __always_inline__, __nodebug__, __target__("avx"), __min_vector_width__(256)
func Mm256TestzPs(v0 M256, v1 M256) Int {
	return C._mm256_testz_ps(v0, v1)
}

// _mm256_testc_ps
// __always_inline__, __nodebug__, __target__("avx"), __min_vector_width__(256)
func Mm256TestcPs(v0 M256, v1 M256) Int {
	return C._mm256_testc_ps(v0, v1)
}

// _mm256_testnzc_ps
// __always_inline__, __nodebug__, __target__("avx"), __min_vector_width__(256)
func Mm256TestnzcPs(v0 M256, v1 M256) Int {
	return C._mm256_testnzc_ps(v0, v1)
}

// _mm256_testz_si256
// __always_inline__, __nodebug__, __target__("avx"), __min_vector_width__(256)
func Mm256TestzSi256(v0 M256I, v1 M256I) Int {
	return C._mm256_testz_si256(v0, v1)
}

// _mm256_testc_si256
// __always_inline__, __nodebug__, __target__("avx"), __min_vector_width__(256)
func Mm256TestcSi256(v0 M256I, v1 M256I) Int {
	return C._mm256_testc_si256(v0, v1)
}

// _mm256_testnzc_si256
// __always_inline__, __nodebug__, __target__("avx"), __min_vector_width__(256)
func Mm256TestnzcSi256(v0 M256I, v1 M256I) Int {
	return C._mm256_testnzc_si256(v0, v1)
}

// _mm256_movemask_pd
// __always_inline__, __nodebug__, __target__("avx"), __min_vector_width__(256)
func Mm256MovemaskPd(v0 M256D) Int {
	return C._mm256_movemask_pd(v0)
}

// _mm256_movemask_ps
// __always_inline__, __nodebug__, __target__("avx"), __min_vector_width__(256)
func Mm256MovemaskPs(v0 M256) Int {
	return C._mm256_movemask_ps(v0)
}

// _mm256_zeroall
// __always_inline__, __nodebug__, __target__("avx")
func Mm256Zeroall() {
	C._mm256_zeroall()
}

// _mm256_zeroupper
// __always_inline__, __nodebug__, __target__("avx")
func Mm256Zeroupper() {
	C._mm256_zeroupper()
}

// _mm256_undefined_pd
// __always_inline__, __nodebug__, __target__("avx"), __min_vector_width__(256)
func Mm256UndefinedPd() M256D {
	return C._mm256_undefined_pd()
}

// _mm256_undefined_ps
// __always_inline__, __nodebug__, __target__("avx"), __min_vector_width__(256)
func Mm256UndefinedPs() M256 {
	return C._mm256_undefined_ps()
}

// _mm256_undefined_si256
// __always_inline__, __nodebug__, __target__("avx"), __min_vector_width__(256)
func Mm256UndefinedSi256() M256I {
	return C._mm256_undefined_si256()
}

// _mm256_set_pd
// __always_inline__, __nodebug__, __target__("avx"), __min_vector_width__(256)
func Mm256SetPd(v0 Double, v1 Double, v2 Double, v3 Double) M256D {
	return C._mm256_set_pd(v0, v1, v2, v3)
}

// _mm256_set_ps
// __always_inline__, __nodebug__, __target__("avx"), __min_vector_width__(256)
func Mm256SetPs(v0 Float, v1 Float, v2 Float, v3 Float, v4 Float, v5 Float, v6 Float, v7 Float) M256 {
	return C._mm256_set_ps(v0, v1, v2, v3, v4, v5, v6, v7)
}

// _mm256_set_epi32
// __always_inline__, __nodebug__, __target__("avx"), __min_vector_width__(256)
func Mm256SetEpi32(v0 Int, v1 Int, v2 Int, v3 Int, v4 Int, v5 Int, v6 Int, v7 Int) M256I {
	return C._mm256_set_epi32(v0, v1, v2, v3, v4, v5, v6, v7)
}

// _mm256_set_epi16
// __always_inline__, __nodebug__, __target__("avx"), __min_vector_width__(256)
func Mm256SetEpi16(v0 Short, v1 Short, v2 Short, v3 Short, v4 Short, v5 Short, v6 Short, v7 Short, v8 Short, v9 Short, v10 Short, v11 Short, v12 Short, v13 Short, v14 Short, v15 Short) M256I {
	return C._mm256_set_epi16(v0, v1, v2, v3, v4, v5, v6, v7, v8, v9, v10, v11, v12, v13, v14, v15)
}

// _mm256_set_epi8
// __always_inline__, __nodebug__, __target__("avx"), __min_vector_width__(256)
func Mm256SetEpi8(v0 Char, v1 Char, v2 Char, v3 Char, v4 Char, v5 Char, v6 Char, v7 Char, v8 Char, v9 Char, v10 Char, v11 Char, v12 Char, v13 Char, v14 Char, v15 Char, v16 Char, v17 Char, v18 Char, v19 Char, v20 Char, v21 Char, v22 Char, v23 Char, v24 Char, v25 Char, v26 Char, v27 Char, v28 Char, v29 Char, v30 Char, v31 Char) M256I {
	return C._mm256_set_epi8(v0, v1, v2, v3, v4, v5, v6, v7, v8, v9, v10, v11, v12, v13, v14, v15, v16, v17, v18, v19, v20, v21, v22, v23, v24, v25, v26, v27, v28, v29, v30, v31)
}

// _mm256_set_epi64x
// __always_inline__, __nodebug__, __target__("avx"), __min_vector_width__(256)
func Mm256SetEpi64X(v0 Longlong, v1 Longlong, v2 Longlong, v3 Longlong) M256I {
	return C._mm256_set_epi64x(v0, v1, v2, v3)
}

// _mm256_setr_pd
// __always_inline__, __nodebug__, __target__("avx"), __min_vector_width__(256)
func Mm256SetrPd(v0 Double, v1 Double, v2 Double, v3 Double) M256D {
	return C._mm256_setr_pd(v0, v1, v2, v3)
}

// _mm256_setr_ps
// __always_inline__, __nodebug__, __target__("avx"), __min_vector_width__(256)
func Mm256SetrPs(v0 Float, v1 Float, v2 Float, v3 Float, v4 Float, v5 Float, v6 Float, v7 Float) M256 {
	return C._mm256_setr_ps(v0, v1, v2, v3, v4, v5, v6, v7)
}

// _mm256_setr_epi32
// __always_inline__, __nodebug__, __target__("avx"), __min_vector_width__(256)
func Mm256SetrEpi32(v0 Int, v1 Int, v2 Int, v3 Int, v4 Int, v5 Int, v6 Int, v7 Int) M256I {
	return C._mm256_setr_epi32(v0, v1, v2, v3, v4, v5, v6, v7)
}

// _mm256_setr_epi16
// __always_inline__, __nodebug__, __target__("avx"), __min_vector_width__(256)
func Mm256SetrEpi16(v0 Short, v1 Short, v2 Short, v3 Short, v4 Short, v5 Short, v6 Short, v7 Short, v8 Short, v9 Short, v10 Short, v11 Short, v12 Short, v13 Short, v14 Short, v15 Short) M256I {
	return C._mm256_setr_epi16(v0, v1, v2, v3, v4, v5, v6, v7, v8, v9, v10, v11, v12, v13, v14, v15)
}

// _mm256_setr_epi8
// __always_inline__, __nodebug__, __target__("avx"), __min_vector_width__(256)
func Mm256SetrEpi8(v0 Char, v1 Char, v2 Char, v3 Char, v4 Char, v5 Char, v6 Char, v7 Char, v8 Char, v9 Char, v10 Char, v11 Char, v12 Char, v13 Char, v14 Char, v15 Char, v16 Char, v17 Char, v18 Char, v19 Char, v20 Char, v21 Char, v22 Char, v23 Char, v24 Char, v25 Char, v26 Char, v27 Char, v28 Char, v29 Char, v30 Char, v31 Char) M256I {
	return C._mm256_setr_epi8(v0, v1, v2, v3, v4, v5, v6, v7, v8, v9, v10, v11, v12, v13, v14, v15, v16, v17, v18, v19, v20, v21, v22, v23, v24, v25, v26, v27, v28, v29, v30, v31)
}

// _mm256_setr_epi64x
// __always_inline__, __nodebug__, __target__("avx"), __min_vector_width__(256)
func Mm256SetrEpi64X(v0 Longlong, v1 Longlong, v2 Longlong, v3 Longlong) M256I {
	return C._mm256_setr_epi64x(v0, v1, v2, v3)
}

// _mm256_set1_pd
// __always_inline__, __nodebug__, __target__("avx"), __min_vector_width__(256)
func Mm256Set1Pd(v0 Double) M256D {
	return C._mm256_set1_pd(v0)
}

// _mm256_set1_ps
// __always_inline__, __nodebug__, __target__("avx"), __min_vector_width__(256)
func Mm256Set1Ps(v0 Float) M256 {
	return C._mm256_set1_ps(v0)
}

// _mm256_set1_epi32
// __always_inline__, __nodebug__, __target__("avx"), __min_vector_width__(256)
func Mm256Set1Epi32(v0 Int) M256I {
	return C._mm256_set1_epi32(v0)
}

// _mm256_set1_epi16
// __always_inline__, __nodebug__, __target__("avx"), __min_vector_width__(256)
func Mm256Set1Epi16(v0 Short) M256I {
	return C._mm256_set1_epi16(v0)
}

// _mm256_set1_epi8
// __always_inline__, __nodebug__, __target__("avx"), __min_vector_width__(256)
func Mm256Set1Epi8(v0 Char) M256I {
	return C._mm256_set1_epi8(v0)
}

// _mm256_set1_epi64x
// __always_inline__, __nodebug__, __target__("avx"), __min_vector_width__(256)
func Mm256Set1Epi64X(v0 Longlong) M256I {
	return C._mm256_set1_epi64x(v0)
}

// _mm256_setzero_pd
// __always_inline__, __nodebug__, __target__("avx"), __min_vector_width__(256)
func Mm256SetzeroPd() M256D {
	return C._mm256_setzero_pd()
}

// _mm256_setzero_ps
// __always_inline__, __nodebug__, __target__("avx"), __min_vector_width__(256)
func Mm256SetzeroPs() M256 {
	return C._mm256_setzero_ps()
}

// _mm256_setzero_si256
// __always_inline__, __nodebug__, __target__("avx"), __min_vector_width__(256)
func Mm256SetzeroSi256() M256I {
	return C._mm256_setzero_si256()
}

// _mm256_castpd_ps
// __always_inline__, __nodebug__, __target__("avx"), __min_vector_width__(256)
func Mm256CastpdPs(v0 M256D) M256 {
	return C._mm256_castpd_ps(v0)
}

// _mm256_castpd_si256
// __always_inline__, __nodebug__, __target__("avx"), __min_vector_width__(256)
func Mm256CastpdSi256(v0 M256D) M256I {
	return C._mm256_castpd_si256(v0)
}

// _mm256_castps_pd
// __always_inline__, __nodebug__, __target__("avx"), __min_vector_width__(256)
func Mm256CastpsPd(v0 M256) M256D {
	return C._mm256_castps_pd(v0)
}

// _mm256_castps_si256
// __always_inline__, __nodebug__, __target__("avx"), __min_vector_width__(256)
func Mm256CastpsSi256(v0 M256) M256I {
	return C._mm256_castps_si256(v0)
}

// _mm256_castsi256_ps
// __always_inline__, __nodebug__, __target__("avx"), __min_vector_width__(256)
func Mm256Castsi256Ps(v0 M256I) M256 {
	return C._mm256_castsi256_ps(v0)
}

// _mm256_castsi256_pd
// __always_inline__, __nodebug__, __target__("avx"), __min_vector_width__(256)
func Mm256Castsi256Pd(v0 M256I) M256D {
	return C._mm256_castsi256_pd(v0)
}

// _mm256_castpd256_pd128
// __always_inline__, __nodebug__, __target__("avx"), __min_vector_width__(256)
func Mm256Castpd256Pd128(v0 M256D) M128D {
	return C._mm256_castpd256_pd128(v0)
}

// _mm256_castps256_ps128
// __always_inline__, __nodebug__, __target__("avx"), __min_vector_width__(256)
func Mm256Castps256Ps128(v0 M256) M128 {
	return C._mm256_castps256_ps128(v0)
}

// _mm256_castsi256_si128
// __always_inline__, __nodebug__, __target__("avx"), __min_vector_width__(256)
func Mm256Castsi256Si128(v0 M256I) M128I {
	return C._mm256_castsi256_si128(v0)
}

// _mm256_castpd128_pd256
// __always_inline__, __nodebug__, __target__("avx"), __min_vector_width__(256)
func Mm256Castpd128Pd256(v0 M128D) M256D {
	return C._mm256_castpd128_pd256(v0)
}

// _mm256_castps128_ps256
// __always_inline__, __nodebug__, __target__("avx"), __min_vector_width__(256)
func Mm256Castps128Ps256(v0 M128) M256 {
	return C._mm256_castps128_ps256(v0)
}

// _mm256_castsi128_si256
// __always_inline__, __nodebug__, __target__("avx"), __min_vector_width__(256)
func Mm256Castsi128Si256(v0 M128I) M256I {
	return C._mm256_castsi128_si256(v0)
}

// _mm256_zextpd128_pd256
// __always_inline__, __nodebug__, __target__("avx"), __min_vector_width__(256)
func Mm256Zextpd128Pd256(v0 M128D) M256D {
	return C._mm256_zextpd128_pd256(v0)
}

// _mm256_zextps128_ps256
// __always_inline__, __nodebug__, __target__("avx"), __min_vector_width__(256)
func Mm256Zextps128Ps256(v0 M128) M256 {
	return C._mm256_zextps128_ps256(v0)
}

// _mm256_zextsi128_si256
// __always_inline__, __nodebug__, __target__("avx"), __min_vector_width__(256)
func Mm256Zextsi128Si256(v0 M128I) M256I {
	return C._mm256_zextsi128_si256(v0)
}

// _mm256_set_m128
// __always_inline__, __nodebug__, __target__("avx"), __min_vector_width__(256)
func Mm256SetM128(v0 M128, v1 M128) M256 {
	return C._mm256_set_m128(v0, v1)
}

// _mm256_set_m128d
// __always_inline__, __nodebug__, __target__("avx"), __min_vector_width__(256)
func Mm256SetM128D(v0 M128D, v1 M128D) M256D {
	return C._mm256_set_m128d(v0, v1)
}

// _mm256_set_m128i
// __always_inline__, __nodebug__, __target__("avx"), __min_vector_width__(256)
func Mm256SetM128I(v0 M128I, v1 M128I) M256I {
	return C._mm256_set_m128i(v0, v1)
}

// _mm256_setr_m128
// __always_inline__, __nodebug__, __target__("avx"), __min_vector_width__(256)
func Mm256SetrM128(v0 M128, v1 M128) M256 {
	return C._mm256_setr_m128(v0, v1)
}

// _mm256_setr_m128d
// __always_inline__, __nodebug__, __target__("avx"), __min_vector_width__(256)
func Mm256SetrM128D(v0 M128D, v1 M128D) M256D {
	return C._mm256_setr_m128d(v0, v1)
}

// _mm256_setr_m128i
// __always_inline__, __nodebug__, __target__("avx"), __min_vector_width__(256)
func Mm256SetrM128I(v0 M128I, v1 M128I) M256I {
	return C._mm256_setr_m128i(v0, v1)
}
