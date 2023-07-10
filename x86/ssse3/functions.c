#include <immintrin.h>

void MmAbsEpi8(__m128i* r, __m128i* v0) { *r = _mm_abs_epi8(*v0); }
void MmAbsEpi16(__m128i* r, __m128i* v0) { *r = _mm_abs_epi16(*v0); }
void MmAbsEpi32(__m128i* r, __m128i* v0) { *r = _mm_abs_epi32(*v0); }
void MmHaddEpi16(__m128i* r, __m128i* v0, __m128i* v1) { *r = _mm_hadd_epi16(*v0, *v1); }
void MmHaddEpi32(__m128i* r, __m128i* v0, __m128i* v1) { *r = _mm_hadd_epi32(*v0, *v1); }
void MmHaddsEpi16(__m128i* r, __m128i* v0, __m128i* v1) { *r = _mm_hadds_epi16(*v0, *v1); }
void MmHsubEpi16(__m128i* r, __m128i* v0, __m128i* v1) { *r = _mm_hsub_epi16(*v0, *v1); }
void MmHsubEpi32(__m128i* r, __m128i* v0, __m128i* v1) { *r = _mm_hsub_epi32(*v0, *v1); }
void MmHsubsEpi16(__m128i* r, __m128i* v0, __m128i* v1) { *r = _mm_hsubs_epi16(*v0, *v1); }
void MmMaddubsEpi16(__m128i* r, __m128i* v0, __m128i* v1) { *r = _mm_maddubs_epi16(*v0, *v1); }
void MmMulhrsEpi16(__m128i* r, __m128i* v0, __m128i* v1) { *r = _mm_mulhrs_epi16(*v0, *v1); }
void MmShuffleEpi8(__m128i* r, __m128i* v0, __m128i* v1) { *r = _mm_shuffle_epi8(*v0, *v1); }
void MmSignEpi8(__m128i* r, __m128i* v0, __m128i* v1) { *r = _mm_sign_epi8(*v0, *v1); }
void MmSignEpi16(__m128i* r, __m128i* v0, __m128i* v1) { *r = _mm_sign_epi16(*v0, *v1); }
void MmSignEpi32(__m128i* r, __m128i* v0, __m128i* v1) { *r = _mm_sign_epi32(*v0, *v1); }
