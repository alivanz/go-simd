#include <immintrin.h>

void MmCvtpsPi32(__m64* r, __m128* v0) { *r = _mm_cvtps_pi32(*v0); }
void MmCvtPs2Pi(__m64* r, __m128* v0) { *r = _mm_cvt_ps2pi(*v0); }
void MmCvttpsPi32(__m64* r, __m128* v0) { *r = _mm_cvttps_pi32(*v0); }
void MmCvttPs2Pi(__m64* r, __m128* v0) { *r = _mm_cvtt_ps2pi(*v0); }
void MmCvtpi32Ps(__m128* r, __m128* v0, __m64* v1) { *r = _mm_cvtpi32_ps(*v0, *v1); }
void MmCvtPi2Ps(__m128* r, __m128* v0, __m64* v1) { *r = _mm_cvt_pi2ps(*v0, *v1); }
void MmMaxPi16(__m64* r, __m64* v0, __m64* v1) { *r = _mm_max_pi16(*v0, *v1); }
void MmMaxPu8(__m64* r, __m64* v0, __m64* v1) { *r = _mm_max_pu8(*v0, *v1); }
void MmMinPi16(__m64* r, __m64* v0, __m64* v1) { *r = _mm_min_pi16(*v0, *v1); }
void MmMinPu8(__m64* r, __m64* v0, __m64* v1) { *r = _mm_min_pu8(*v0, *v1); }
void MmMovemaskPi8(int* r, __m64* v0) { *r = _mm_movemask_pi8(*v0); }
void MmMulhiPu16(__m64* r, __m64* v0, __m64* v1) { *r = _mm_mulhi_pu16(*v0, *v1); }
void MmAvgPu8(__m64* r, __m64* v0, __m64* v1) { *r = _mm_avg_pu8(*v0, *v1); }
void MmAvgPu16(__m64* r, __m64* v0, __m64* v1) { *r = _mm_avg_pu16(*v0, *v1); }
void MmSadPu8(__m64* r, __m64* v0, __m64* v1) { *r = _mm_sad_pu8(*v0, *v1); }
void MmCvtpi16Ps(__m128* r, __m64* v0) { *r = _mm_cvtpi16_ps(*v0); }
void MmCvtpu16Ps(__m128* r, __m64* v0) { *r = _mm_cvtpu16_ps(*v0); }
void MmCvtpi8Ps(__m128* r, __m64* v0) { *r = _mm_cvtpi8_ps(*v0); }
void MmCvtpu8Ps(__m128* r, __m64* v0) { *r = _mm_cvtpu8_ps(*v0); }
void MmCvtpi32X2Ps(__m128* r, __m64* v0, __m64* v1) { *r = _mm_cvtpi32x2_ps(*v0, *v1); }
void MmCvtpsPi16(__m64* r, __m128* v0) { *r = _mm_cvtps_pi16(*v0); }
void MmCvtpsPi8(__m64* r, __m128* v0) { *r = _mm_cvtps_pi8(*v0); }
