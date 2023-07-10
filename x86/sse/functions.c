#include <immintrin.h>

void MmAddSs(__m128* r, __m128* v0, __m128* v1) { *r = _mm_add_ss(*v0, *v1); }
void MmAddPs(__m128* r, __m128* v0, __m128* v1) { *r = _mm_add_ps(*v0, *v1); }
void MmSubSs(__m128* r, __m128* v0, __m128* v1) { *r = _mm_sub_ss(*v0, *v1); }
void MmSubPs(__m128* r, __m128* v0, __m128* v1) { *r = _mm_sub_ps(*v0, *v1); }
void MmMulSs(__m128* r, __m128* v0, __m128* v1) { *r = _mm_mul_ss(*v0, *v1); }
void MmMulPs(__m128* r, __m128* v0, __m128* v1) { *r = _mm_mul_ps(*v0, *v1); }
void MmDivSs(__m128* r, __m128* v0, __m128* v1) { *r = _mm_div_ss(*v0, *v1); }
void MmDivPs(__m128* r, __m128* v0, __m128* v1) { *r = _mm_div_ps(*v0, *v1); }
void MmSqrtSs(__m128* r, __m128* v0) { *r = _mm_sqrt_ss(*v0); }
void MmSqrtPs(__m128* r, __m128* v0) { *r = _mm_sqrt_ps(*v0); }
void MmRcpSs(__m128* r, __m128* v0) { *r = _mm_rcp_ss(*v0); }
void MmRcpPs(__m128* r, __m128* v0) { *r = _mm_rcp_ps(*v0); }
void MmRsqrtSs(__m128* r, __m128* v0) { *r = _mm_rsqrt_ss(*v0); }
void MmRsqrtPs(__m128* r, __m128* v0) { *r = _mm_rsqrt_ps(*v0); }
void MmMinSs(__m128* r, __m128* v0, __m128* v1) { *r = _mm_min_ss(*v0, *v1); }
void MmMinPs(__m128* r, __m128* v0, __m128* v1) { *r = _mm_min_ps(*v0, *v1); }
void MmMaxSs(__m128* r, __m128* v0, __m128* v1) { *r = _mm_max_ss(*v0, *v1); }
void MmMaxPs(__m128* r, __m128* v0, __m128* v1) { *r = _mm_max_ps(*v0, *v1); }
void MmAndPs(__m128* r, __m128* v0, __m128* v1) { *r = _mm_and_ps(*v0, *v1); }
void MmAndnotPs(__m128* r, __m128* v0, __m128* v1) { *r = _mm_andnot_ps(*v0, *v1); }
void MmOrPs(__m128* r, __m128* v0, __m128* v1) { *r = _mm_or_ps(*v0, *v1); }
void MmXorPs(__m128* r, __m128* v0, __m128* v1) { *r = _mm_xor_ps(*v0, *v1); }
void MmCmpeqSs(__m128* r, __m128* v0, __m128* v1) { *r = _mm_cmpeq_ss(*v0, *v1); }
void MmCmpeqPs(__m128* r, __m128* v0, __m128* v1) { *r = _mm_cmpeq_ps(*v0, *v1); }
void MmCmpltSs(__m128* r, __m128* v0, __m128* v1) { *r = _mm_cmplt_ss(*v0, *v1); }
void MmCmpltPs(__m128* r, __m128* v0, __m128* v1) { *r = _mm_cmplt_ps(*v0, *v1); }
void MmCmpleSs(__m128* r, __m128* v0, __m128* v1) { *r = _mm_cmple_ss(*v0, *v1); }
void MmCmplePs(__m128* r, __m128* v0, __m128* v1) { *r = _mm_cmple_ps(*v0, *v1); }
void MmCmpgtSs(__m128* r, __m128* v0, __m128* v1) { *r = _mm_cmpgt_ss(*v0, *v1); }
void MmCmpgtPs(__m128* r, __m128* v0, __m128* v1) { *r = _mm_cmpgt_ps(*v0, *v1); }
void MmCmpgeSs(__m128* r, __m128* v0, __m128* v1) { *r = _mm_cmpge_ss(*v0, *v1); }
void MmCmpgePs(__m128* r, __m128* v0, __m128* v1) { *r = _mm_cmpge_ps(*v0, *v1); }
void MmCmpneqSs(__m128* r, __m128* v0, __m128* v1) { *r = _mm_cmpneq_ss(*v0, *v1); }
void MmCmpneqPs(__m128* r, __m128* v0, __m128* v1) { *r = _mm_cmpneq_ps(*v0, *v1); }
void MmCmpnltSs(__m128* r, __m128* v0, __m128* v1) { *r = _mm_cmpnlt_ss(*v0, *v1); }
void MmCmpnltPs(__m128* r, __m128* v0, __m128* v1) { *r = _mm_cmpnlt_ps(*v0, *v1); }
void MmCmpnleSs(__m128* r, __m128* v0, __m128* v1) { *r = _mm_cmpnle_ss(*v0, *v1); }
void MmCmpnlePs(__m128* r, __m128* v0, __m128* v1) { *r = _mm_cmpnle_ps(*v0, *v1); }
void MmCmpngtSs(__m128* r, __m128* v0, __m128* v1) { *r = _mm_cmpngt_ss(*v0, *v1); }
void MmCmpngtPs(__m128* r, __m128* v0, __m128* v1) { *r = _mm_cmpngt_ps(*v0, *v1); }
void MmCmpngeSs(__m128* r, __m128* v0, __m128* v1) { *r = _mm_cmpnge_ss(*v0, *v1); }
void MmCmpngePs(__m128* r, __m128* v0, __m128* v1) { *r = _mm_cmpnge_ps(*v0, *v1); }
void MmCmpordSs(__m128* r, __m128* v0, __m128* v1) { *r = _mm_cmpord_ss(*v0, *v1); }
void MmCmpordPs(__m128* r, __m128* v0, __m128* v1) { *r = _mm_cmpord_ps(*v0, *v1); }
void MmCmpunordSs(__m128* r, __m128* v0, __m128* v1) { *r = _mm_cmpunord_ss(*v0, *v1); }
void MmCmpunordPs(__m128* r, __m128* v0, __m128* v1) { *r = _mm_cmpunord_ps(*v0, *v1); }
void MmComieqSs(int* r, __m128* v0, __m128* v1) { *r = _mm_comieq_ss(*v0, *v1); }
void MmComiltSs(int* r, __m128* v0, __m128* v1) { *r = _mm_comilt_ss(*v0, *v1); }
void MmComileSs(int* r, __m128* v0, __m128* v1) { *r = _mm_comile_ss(*v0, *v1); }
void MmComigtSs(int* r, __m128* v0, __m128* v1) { *r = _mm_comigt_ss(*v0, *v1); }
void MmComigeSs(int* r, __m128* v0, __m128* v1) { *r = _mm_comige_ss(*v0, *v1); }
void MmComineqSs(int* r, __m128* v0, __m128* v1) { *r = _mm_comineq_ss(*v0, *v1); }
void MmUcomieqSs(int* r, __m128* v0, __m128* v1) { *r = _mm_ucomieq_ss(*v0, *v1); }
void MmUcomiltSs(int* r, __m128* v0, __m128* v1) { *r = _mm_ucomilt_ss(*v0, *v1); }
void MmUcomileSs(int* r, __m128* v0, __m128* v1) { *r = _mm_ucomile_ss(*v0, *v1); }
void MmUcomigtSs(int* r, __m128* v0, __m128* v1) { *r = _mm_ucomigt_ss(*v0, *v1); }
void MmUcomigeSs(int* r, __m128* v0, __m128* v1) { *r = _mm_ucomige_ss(*v0, *v1); }
void MmUcomineqSs(int* r, __m128* v0, __m128* v1) { *r = _mm_ucomineq_ss(*v0, *v1); }
void MmCvtssSi32(int* r, __m128* v0) { *r = _mm_cvtss_si32(*v0); }
void MmCvtSs2Si(int* r, __m128* v0) { *r = _mm_cvt_ss2si(*v0); }
void MmCvtssSi64(long long* r, __m128* v0) { *r = _mm_cvtss_si64(*v0); }
void MmCvttssSi32(int* r, __m128* v0) { *r = _mm_cvttss_si32(*v0); }
void MmCvttSs2Si(int* r, __m128* v0) { *r = _mm_cvtt_ss2si(*v0); }
void MmCvttssSi64(long long* r, __m128* v0) { *r = _mm_cvttss_si64(*v0); }
void MmCvtsi32Ss(__m128* r, __m128* v0, int* v1) { *r = _mm_cvtsi32_ss(*v0, *v1); }
void MmCvtSi2Ss(__m128* r, __m128* v0, int* v1) { *r = _mm_cvt_si2ss(*v0, *v1); }
void MmCvtsi64Ss(__m128* r, __m128* v0, long long* v1) { *r = _mm_cvtsi64_ss(*v0, *v1); }
void MmCvtssF32(float* r, __m128* v0) { *r = _mm_cvtss_f32(*v0); }
void MmUndefinedPs(__m128* r) { *r = _mm_undefined_ps(); }
void MmSetSs(__m128* r, float* v0) { *r = _mm_set_ss(*v0); }
void MmSet1Ps(__m128* r, float* v0) { *r = _mm_set1_ps(*v0); }
void MmSetPs1(__m128* r, float* v0) { *r = _mm_set_ps1(*v0); }
void MmSetPs(__m128* r, float* v0, float* v1, float* v2, float* v3) { *r = _mm_set_ps(*v0, *v1, *v2, *v3); }
void MmSetrPs(__m128* r, float* v0, float* v1, float* v2, float* v3) { *r = _mm_setr_ps(*v0, *v1, *v2, *v3); }
void MmSetzeroPs(__m128* r) { *r = _mm_setzero_ps(); }
void MmUnpackhiPs(__m128* r, __m128* v0, __m128* v1) { *r = _mm_unpackhi_ps(*v0, *v1); }
void MmUnpackloPs(__m128* r, __m128* v0, __m128* v1) { *r = _mm_unpacklo_ps(*v0, *v1); }
void MmMoveSs(__m128* r, __m128* v0, __m128* v1) { *r = _mm_move_ss(*v0, *v1); }
void MmMovehlPs(__m128* r, __m128* v0, __m128* v1) { *r = _mm_movehl_ps(*v0, *v1); }
void MmMovelhPs(__m128* r, __m128* v0, __m128* v1) { *r = _mm_movelh_ps(*v0, *v1); }
void MmMovemaskPs(int* r, __m128* v0) { *r = _mm_movemask_ps(*v0); }
