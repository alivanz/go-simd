#include <immintrin.h>

void MmCvtpdPi32(__m64* r, __m128d* v0) { *r = _mm_cvtpd_pi32(*v0); }
void MmCvttpdPi32(__m64* r, __m128d* v0) { *r = _mm_cvttpd_pi32(*v0); }
void MmCvtpi32Pd(__m128d* r, __m64* v0) { *r = _mm_cvtpi32_pd(*v0); }
void MmAddSi64(__m64* r, __m64* v0, __m64* v1) { *r = _mm_add_si64(*v0, *v1); }
void MmMulSu32(__m64* r, __m64* v0, __m64* v1) { *r = _mm_mul_su32(*v0, *v1); }
void MmSubSi64(__m64* r, __m64* v0, __m64* v1) { *r = _mm_sub_si64(*v0, *v1); }
