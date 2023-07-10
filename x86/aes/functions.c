#include <immintrin.h>

void MmAesencSi128(__m128i* r, __m128i* v0, __m128i* v1) { *r = _mm_aesenc_si128(*v0, *v1); }
void MmAesenclastSi128(__m128i* r, __m128i* v0, __m128i* v1) { *r = _mm_aesenclast_si128(*v0, *v1); }
void MmAesdecSi128(__m128i* r, __m128i* v0, __m128i* v1) { *r = _mm_aesdec_si128(*v0, *v1); }
void MmAesdeclastSi128(__m128i* r, __m128i* v0, __m128i* v1) { *r = _mm_aesdeclast_si128(*v0, *v1); }
void MmAesimcSi128(__m128i* r, __m128i* v0) { *r = _mm_aesimc_si128(*v0); }
