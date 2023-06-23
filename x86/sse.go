package x86

/*
#cgo CFLAGS: 
#include <immintrin.h>
*/
import "C"

// _mm_add_ss
// __always_inline__, __nodebug__, __target__("sse"), __min_vector_width__(128)
func MmAddSs(v0 M128, v1 M128) M128 {
	return C._mm_add_ss(v0, v1)
}

// _mm_add_ps
// __always_inline__, __nodebug__, __target__("sse"), __min_vector_width__(128)
func MmAddPs(v0 M128, v1 M128) M128 {
	return C._mm_add_ps(v0, v1)
}

// _mm_sub_ss
// __always_inline__, __nodebug__, __target__("sse"), __min_vector_width__(128)
func MmSubSs(v0 M128, v1 M128) M128 {
	return C._mm_sub_ss(v0, v1)
}

// _mm_sub_ps
// __always_inline__, __nodebug__, __target__("sse"), __min_vector_width__(128)
func MmSubPs(v0 M128, v1 M128) M128 {
	return C._mm_sub_ps(v0, v1)
}

// _mm_mul_ss
// __always_inline__, __nodebug__, __target__("sse"), __min_vector_width__(128)
func MmMulSs(v0 M128, v1 M128) M128 {
	return C._mm_mul_ss(v0, v1)
}

// _mm_mul_ps
// __always_inline__, __nodebug__, __target__("sse"), __min_vector_width__(128)
func MmMulPs(v0 M128, v1 M128) M128 {
	return C._mm_mul_ps(v0, v1)
}

// _mm_div_ss
// __always_inline__, __nodebug__, __target__("sse"), __min_vector_width__(128)
func MmDivSs(v0 M128, v1 M128) M128 {
	return C._mm_div_ss(v0, v1)
}

// _mm_div_ps
// __always_inline__, __nodebug__, __target__("sse"), __min_vector_width__(128)
func MmDivPs(v0 M128, v1 M128) M128 {
	return C._mm_div_ps(v0, v1)
}

// _mm_sqrt_ss
// __always_inline__, __nodebug__, __target__("sse"), __min_vector_width__(128)
func MmSqrtSs(v0 M128) M128 {
	return C._mm_sqrt_ss(v0)
}

// _mm_sqrt_ps
// __always_inline__, __nodebug__, __target__("sse"), __min_vector_width__(128)
func MmSqrtPs(v0 M128) M128 {
	return C._mm_sqrt_ps(v0)
}

// _mm_rcp_ss
// __always_inline__, __nodebug__, __target__("sse"), __min_vector_width__(128)
func MmRcpSs(v0 M128) M128 {
	return C._mm_rcp_ss(v0)
}

// _mm_rcp_ps
// __always_inline__, __nodebug__, __target__("sse"), __min_vector_width__(128)
func MmRcpPs(v0 M128) M128 {
	return C._mm_rcp_ps(v0)
}

// _mm_rsqrt_ss
// __always_inline__, __nodebug__, __target__("sse"), __min_vector_width__(128)
func MmRsqrtSs(v0 M128) M128 {
	return C._mm_rsqrt_ss(v0)
}

// _mm_rsqrt_ps
// __always_inline__, __nodebug__, __target__("sse"), __min_vector_width__(128)
func MmRsqrtPs(v0 M128) M128 {
	return C._mm_rsqrt_ps(v0)
}

// _mm_min_ss
// __always_inline__, __nodebug__, __target__("sse"), __min_vector_width__(128)
func MmMinSs(v0 M128, v1 M128) M128 {
	return C._mm_min_ss(v0, v1)
}

// _mm_min_ps
// __always_inline__, __nodebug__, __target__("sse"), __min_vector_width__(128)
func MmMinPs(v0 M128, v1 M128) M128 {
	return C._mm_min_ps(v0, v1)
}

// _mm_max_ss
// __always_inline__, __nodebug__, __target__("sse"), __min_vector_width__(128)
func MmMaxSs(v0 M128, v1 M128) M128 {
	return C._mm_max_ss(v0, v1)
}

// _mm_max_ps
// __always_inline__, __nodebug__, __target__("sse"), __min_vector_width__(128)
func MmMaxPs(v0 M128, v1 M128) M128 {
	return C._mm_max_ps(v0, v1)
}

// _mm_and_ps
// __always_inline__, __nodebug__, __target__("sse"), __min_vector_width__(128)
func MmAndPs(v0 M128, v1 M128) M128 {
	return C._mm_and_ps(v0, v1)
}

// _mm_andnot_ps
// __always_inline__, __nodebug__, __target__("sse"), __min_vector_width__(128)
func MmAndnotPs(v0 M128, v1 M128) M128 {
	return C._mm_andnot_ps(v0, v1)
}

// _mm_or_ps
// __always_inline__, __nodebug__, __target__("sse"), __min_vector_width__(128)
func MmOrPs(v0 M128, v1 M128) M128 {
	return C._mm_or_ps(v0, v1)
}

// _mm_xor_ps
// __always_inline__, __nodebug__, __target__("sse"), __min_vector_width__(128)
func MmXorPs(v0 M128, v1 M128) M128 {
	return C._mm_xor_ps(v0, v1)
}

// _mm_cmpeq_ss
// __always_inline__, __nodebug__, __target__("sse"), __min_vector_width__(128)
func MmCmpeqSs(v0 M128, v1 M128) M128 {
	return C._mm_cmpeq_ss(v0, v1)
}

// _mm_cmpeq_ps
// __always_inline__, __nodebug__, __target__("sse"), __min_vector_width__(128)
func MmCmpeqPs(v0 M128, v1 M128) M128 {
	return C._mm_cmpeq_ps(v0, v1)
}

// _mm_cmplt_ss
// __always_inline__, __nodebug__, __target__("sse"), __min_vector_width__(128)
func MmCmpltSs(v0 M128, v1 M128) M128 {
	return C._mm_cmplt_ss(v0, v1)
}

// _mm_cmplt_ps
// __always_inline__, __nodebug__, __target__("sse"), __min_vector_width__(128)
func MmCmpltPs(v0 M128, v1 M128) M128 {
	return C._mm_cmplt_ps(v0, v1)
}

// _mm_cmple_ss
// __always_inline__, __nodebug__, __target__("sse"), __min_vector_width__(128)
func MmCmpleSs(v0 M128, v1 M128) M128 {
	return C._mm_cmple_ss(v0, v1)
}

// _mm_cmple_ps
// __always_inline__, __nodebug__, __target__("sse"), __min_vector_width__(128)
func MmCmplePs(v0 M128, v1 M128) M128 {
	return C._mm_cmple_ps(v0, v1)
}

// _mm_cmpgt_ss
// __always_inline__, __nodebug__, __target__("sse"), __min_vector_width__(128)
func MmCmpgtSs(v0 M128, v1 M128) M128 {
	return C._mm_cmpgt_ss(v0, v1)
}

// _mm_cmpgt_ps
// __always_inline__, __nodebug__, __target__("sse"), __min_vector_width__(128)
func MmCmpgtPs(v0 M128, v1 M128) M128 {
	return C._mm_cmpgt_ps(v0, v1)
}

// _mm_cmpge_ss
// __always_inline__, __nodebug__, __target__("sse"), __min_vector_width__(128)
func MmCmpgeSs(v0 M128, v1 M128) M128 {
	return C._mm_cmpge_ss(v0, v1)
}

// _mm_cmpge_ps
// __always_inline__, __nodebug__, __target__("sse"), __min_vector_width__(128)
func MmCmpgePs(v0 M128, v1 M128) M128 {
	return C._mm_cmpge_ps(v0, v1)
}

// _mm_cmpneq_ss
// __always_inline__, __nodebug__, __target__("sse"), __min_vector_width__(128)
func MmCmpneqSs(v0 M128, v1 M128) M128 {
	return C._mm_cmpneq_ss(v0, v1)
}

// _mm_cmpneq_ps
// __always_inline__, __nodebug__, __target__("sse"), __min_vector_width__(128)
func MmCmpneqPs(v0 M128, v1 M128) M128 {
	return C._mm_cmpneq_ps(v0, v1)
}

// _mm_cmpnlt_ss
// __always_inline__, __nodebug__, __target__("sse"), __min_vector_width__(128)
func MmCmpnltSs(v0 M128, v1 M128) M128 {
	return C._mm_cmpnlt_ss(v0, v1)
}

// _mm_cmpnlt_ps
// __always_inline__, __nodebug__, __target__("sse"), __min_vector_width__(128)
func MmCmpnltPs(v0 M128, v1 M128) M128 {
	return C._mm_cmpnlt_ps(v0, v1)
}

// _mm_cmpnle_ss
// __always_inline__, __nodebug__, __target__("sse"), __min_vector_width__(128)
func MmCmpnleSs(v0 M128, v1 M128) M128 {
	return C._mm_cmpnle_ss(v0, v1)
}

// _mm_cmpnle_ps
// __always_inline__, __nodebug__, __target__("sse"), __min_vector_width__(128)
func MmCmpnlePs(v0 M128, v1 M128) M128 {
	return C._mm_cmpnle_ps(v0, v1)
}

// _mm_cmpngt_ss
// __always_inline__, __nodebug__, __target__("sse"), __min_vector_width__(128)
func MmCmpngtSs(v0 M128, v1 M128) M128 {
	return C._mm_cmpngt_ss(v0, v1)
}

// _mm_cmpngt_ps
// __always_inline__, __nodebug__, __target__("sse"), __min_vector_width__(128)
func MmCmpngtPs(v0 M128, v1 M128) M128 {
	return C._mm_cmpngt_ps(v0, v1)
}

// _mm_cmpnge_ss
// __always_inline__, __nodebug__, __target__("sse"), __min_vector_width__(128)
func MmCmpngeSs(v0 M128, v1 M128) M128 {
	return C._mm_cmpnge_ss(v0, v1)
}

// _mm_cmpnge_ps
// __always_inline__, __nodebug__, __target__("sse"), __min_vector_width__(128)
func MmCmpngePs(v0 M128, v1 M128) M128 {
	return C._mm_cmpnge_ps(v0, v1)
}

// _mm_cmpord_ss
// __always_inline__, __nodebug__, __target__("sse"), __min_vector_width__(128)
func MmCmpordSs(v0 M128, v1 M128) M128 {
	return C._mm_cmpord_ss(v0, v1)
}

// _mm_cmpord_ps
// __always_inline__, __nodebug__, __target__("sse"), __min_vector_width__(128)
func MmCmpordPs(v0 M128, v1 M128) M128 {
	return C._mm_cmpord_ps(v0, v1)
}

// _mm_cmpunord_ss
// __always_inline__, __nodebug__, __target__("sse"), __min_vector_width__(128)
func MmCmpunordSs(v0 M128, v1 M128) M128 {
	return C._mm_cmpunord_ss(v0, v1)
}

// _mm_cmpunord_ps
// __always_inline__, __nodebug__, __target__("sse"), __min_vector_width__(128)
func MmCmpunordPs(v0 M128, v1 M128) M128 {
	return C._mm_cmpunord_ps(v0, v1)
}

// _mm_comieq_ss
// __always_inline__, __nodebug__, __target__("sse"), __min_vector_width__(128)
func MmComieqSs(v0 M128, v1 M128) Int {
	return C._mm_comieq_ss(v0, v1)
}

// _mm_comilt_ss
// __always_inline__, __nodebug__, __target__("sse"), __min_vector_width__(128)
func MmComiltSs(v0 M128, v1 M128) Int {
	return C._mm_comilt_ss(v0, v1)
}

// _mm_comile_ss
// __always_inline__, __nodebug__, __target__("sse"), __min_vector_width__(128)
func MmComileSs(v0 M128, v1 M128) Int {
	return C._mm_comile_ss(v0, v1)
}

// _mm_comigt_ss
// __always_inline__, __nodebug__, __target__("sse"), __min_vector_width__(128)
func MmComigtSs(v0 M128, v1 M128) Int {
	return C._mm_comigt_ss(v0, v1)
}

// _mm_comige_ss
// __always_inline__, __nodebug__, __target__("sse"), __min_vector_width__(128)
func MmComigeSs(v0 M128, v1 M128) Int {
	return C._mm_comige_ss(v0, v1)
}

// _mm_comineq_ss
// __always_inline__, __nodebug__, __target__("sse"), __min_vector_width__(128)
func MmComineqSs(v0 M128, v1 M128) Int {
	return C._mm_comineq_ss(v0, v1)
}

// _mm_ucomieq_ss
// __always_inline__, __nodebug__, __target__("sse"), __min_vector_width__(128)
func MmUcomieqSs(v0 M128, v1 M128) Int {
	return C._mm_ucomieq_ss(v0, v1)
}

// _mm_ucomilt_ss
// __always_inline__, __nodebug__, __target__("sse"), __min_vector_width__(128)
func MmUcomiltSs(v0 M128, v1 M128) Int {
	return C._mm_ucomilt_ss(v0, v1)
}

// _mm_ucomile_ss
// __always_inline__, __nodebug__, __target__("sse"), __min_vector_width__(128)
func MmUcomileSs(v0 M128, v1 M128) Int {
	return C._mm_ucomile_ss(v0, v1)
}

// _mm_ucomigt_ss
// __always_inline__, __nodebug__, __target__("sse"), __min_vector_width__(128)
func MmUcomigtSs(v0 M128, v1 M128) Int {
	return C._mm_ucomigt_ss(v0, v1)
}

// _mm_ucomige_ss
// __always_inline__, __nodebug__, __target__("sse"), __min_vector_width__(128)
func MmUcomigeSs(v0 M128, v1 M128) Int {
	return C._mm_ucomige_ss(v0, v1)
}

// _mm_ucomineq_ss
// __always_inline__, __nodebug__, __target__("sse"), __min_vector_width__(128)
func MmUcomineqSs(v0 M128, v1 M128) Int {
	return C._mm_ucomineq_ss(v0, v1)
}

// _mm_cvtss_si32
// __always_inline__, __nodebug__, __target__("sse"), __min_vector_width__(128)
func MmCvtssSi32(v0 M128) Int {
	return C._mm_cvtss_si32(v0)
}

// _mm_cvt_ss2si
// __always_inline__, __nodebug__, __target__("sse"), __min_vector_width__(128)
func MmCvtSs2Si(v0 M128) Int {
	return C._mm_cvt_ss2si(v0)
}

// _mm_cvtss_si64
// __always_inline__, __nodebug__, __target__("sse"), __min_vector_width__(128)
func MmCvtssSi64(v0 M128) Longlong {
	return C._mm_cvtss_si64(v0)
}

// _mm_cvttss_si32
// __always_inline__, __nodebug__, __target__("sse"), __min_vector_width__(128)
func MmCvttssSi32(v0 M128) Int {
	return C._mm_cvttss_si32(v0)
}

// _mm_cvtt_ss2si
// __always_inline__, __nodebug__, __target__("sse"), __min_vector_width__(128)
func MmCvttSs2Si(v0 M128) Int {
	return C._mm_cvtt_ss2si(v0)
}

// _mm_cvttss_si64
// __always_inline__, __nodebug__, __target__("sse"), __min_vector_width__(128)
func MmCvttssSi64(v0 M128) Longlong {
	return C._mm_cvttss_si64(v0)
}

// _mm_cvtsi32_ss
// __always_inline__, __nodebug__, __target__("sse"), __min_vector_width__(128)
func MmCvtsi32Ss(v0 M128, v1 Int) M128 {
	return C._mm_cvtsi32_ss(v0, v1)
}

// _mm_cvt_si2ss
// __always_inline__, __nodebug__, __target__("sse"), __min_vector_width__(128)
func MmCvtSi2Ss(v0 M128, v1 Int) M128 {
	return C._mm_cvt_si2ss(v0, v1)
}

// _mm_cvtsi64_ss
// __always_inline__, __nodebug__, __target__("sse"), __min_vector_width__(128)
func MmCvtsi64Ss(v0 M128, v1 Longlong) M128 {
	return C._mm_cvtsi64_ss(v0, v1)
}

// _mm_cvtss_f32
// __always_inline__, __nodebug__, __target__("sse"), __min_vector_width__(128)
func MmCvtssF32(v0 M128) Float {
	return C._mm_cvtss_f32(v0)
}

// _mm_undefined_ps
// __always_inline__, __nodebug__, __target__("sse"), __min_vector_width__(128)
func MmUndefinedPs() M128 {
	return C._mm_undefined_ps()
}

// _mm_set_ss
// __always_inline__, __nodebug__, __target__("sse"), __min_vector_width__(128)
func MmSetSs(v0 Float) M128 {
	return C._mm_set_ss(v0)
}

// _mm_set1_ps
// __always_inline__, __nodebug__, __target__("sse"), __min_vector_width__(128)
func MmSet1Ps(v0 Float) M128 {
	return C._mm_set1_ps(v0)
}

// _mm_set_ps1
// __always_inline__, __nodebug__, __target__("sse"), __min_vector_width__(128)
func MmSetPs1(v0 Float) M128 {
	return C._mm_set_ps1(v0)
}

// _mm_set_ps
// __always_inline__, __nodebug__, __target__("sse"), __min_vector_width__(128)
func MmSetPs(v0 Float, v1 Float, v2 Float, v3 Float) M128 {
	return C._mm_set_ps(v0, v1, v2, v3)
}

// _mm_setr_ps
// __always_inline__, __nodebug__, __target__("sse"), __min_vector_width__(128)
func MmSetrPs(v0 Float, v1 Float, v2 Float, v3 Float) M128 {
	return C._mm_setr_ps(v0, v1, v2, v3)
}

// _mm_setzero_ps
// __always_inline__, __nodebug__, __target__("sse"), __min_vector_width__(128)
func MmSetzeroPs() M128 {
	return C._mm_setzero_ps()
}

// _mm_unpackhi_ps
// __always_inline__, __nodebug__, __target__("sse"), __min_vector_width__(128)
func MmUnpackhiPs(v0 M128, v1 M128) M128 {
	return C._mm_unpackhi_ps(v0, v1)
}

// _mm_unpacklo_ps
// __always_inline__, __nodebug__, __target__("sse"), __min_vector_width__(128)
func MmUnpackloPs(v0 M128, v1 M128) M128 {
	return C._mm_unpacklo_ps(v0, v1)
}

// _mm_move_ss
// __always_inline__, __nodebug__, __target__("sse"), __min_vector_width__(128)
func MmMoveSs(v0 M128, v1 M128) M128 {
	return C._mm_move_ss(v0, v1)
}

// _mm_movehl_ps
// __always_inline__, __nodebug__, __target__("sse"), __min_vector_width__(128)
func MmMovehlPs(v0 M128, v1 M128) M128 {
	return C._mm_movehl_ps(v0, v1)
}

// _mm_movelh_ps
// __always_inline__, __nodebug__, __target__("sse"), __min_vector_width__(128)
func MmMovelhPs(v0 M128, v1 M128) M128 {
	return C._mm_movelh_ps(v0, v1)
}

// _mm_movemask_ps
// __always_inline__, __nodebug__, __target__("sse"), __min_vector_width__(128)
func MmMovemaskPs(v0 M128) Int {
	return C._mm_movemask_ps(v0)
}
