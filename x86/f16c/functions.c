#include <immintrin.h>

void CvtshSs(float* r, unsigned short* v0) { *r = _cvtsh_ss(*v0); }
void MmCvtphPs(__m128* r, __m128i* v0) { *r = _mm_cvtph_ps(*v0); }
void Mm256CvtphPs(__m256* r, __m128i* v0) { *r = _mm256_cvtph_ps(*v0); }
