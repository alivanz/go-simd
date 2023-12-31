#include <immintrin.h>

void MmFmaddPs(__m128* r, __m128* v0, __m128* v1, __m128* v2) { *r = _mm_fmadd_ps(*v0, *v1, *v2); }
void MmFmaddPd(__m128d* r, __m128d* v0, __m128d* v1, __m128d* v2) { *r = _mm_fmadd_pd(*v0, *v1, *v2); }
void MmFmaddSs(__m128* r, __m128* v0, __m128* v1, __m128* v2) { *r = _mm_fmadd_ss(*v0, *v1, *v2); }
void MmFmaddSd(__m128d* r, __m128d* v0, __m128d* v1, __m128d* v2) { *r = _mm_fmadd_sd(*v0, *v1, *v2); }
void MmFmsubPs(__m128* r, __m128* v0, __m128* v1, __m128* v2) { *r = _mm_fmsub_ps(*v0, *v1, *v2); }
void MmFmsubPd(__m128d* r, __m128d* v0, __m128d* v1, __m128d* v2) { *r = _mm_fmsub_pd(*v0, *v1, *v2); }
void MmFmsubSs(__m128* r, __m128* v0, __m128* v1, __m128* v2) { *r = _mm_fmsub_ss(*v0, *v1, *v2); }
void MmFmsubSd(__m128d* r, __m128d* v0, __m128d* v1, __m128d* v2) { *r = _mm_fmsub_sd(*v0, *v1, *v2); }
void MmFnmaddPs(__m128* r, __m128* v0, __m128* v1, __m128* v2) { *r = _mm_fnmadd_ps(*v0, *v1, *v2); }
void MmFnmaddPd(__m128d* r, __m128d* v0, __m128d* v1, __m128d* v2) { *r = _mm_fnmadd_pd(*v0, *v1, *v2); }
void MmFnmaddSs(__m128* r, __m128* v0, __m128* v1, __m128* v2) { *r = _mm_fnmadd_ss(*v0, *v1, *v2); }
void MmFnmaddSd(__m128d* r, __m128d* v0, __m128d* v1, __m128d* v2) { *r = _mm_fnmadd_sd(*v0, *v1, *v2); }
void MmFnmsubPs(__m128* r, __m128* v0, __m128* v1, __m128* v2) { *r = _mm_fnmsub_ps(*v0, *v1, *v2); }
void MmFnmsubPd(__m128d* r, __m128d* v0, __m128d* v1, __m128d* v2) { *r = _mm_fnmsub_pd(*v0, *v1, *v2); }
void MmFnmsubSs(__m128* r, __m128* v0, __m128* v1, __m128* v2) { *r = _mm_fnmsub_ss(*v0, *v1, *v2); }
void MmFnmsubSd(__m128d* r, __m128d* v0, __m128d* v1, __m128d* v2) { *r = _mm_fnmsub_sd(*v0, *v1, *v2); }
void MmFmaddsubPs(__m128* r, __m128* v0, __m128* v1, __m128* v2) { *r = _mm_fmaddsub_ps(*v0, *v1, *v2); }
void MmFmaddsubPd(__m128d* r, __m128d* v0, __m128d* v1, __m128d* v2) { *r = _mm_fmaddsub_pd(*v0, *v1, *v2); }
void MmFmsubaddPs(__m128* r, __m128* v0, __m128* v1, __m128* v2) { *r = _mm_fmsubadd_ps(*v0, *v1, *v2); }
void MmFmsubaddPd(__m128d* r, __m128d* v0, __m128d* v1, __m128d* v2) { *r = _mm_fmsubadd_pd(*v0, *v1, *v2); }
void Mm256FmaddPs(__m256* r, __m256* v0, __m256* v1, __m256* v2) { *r = _mm256_fmadd_ps(*v0, *v1, *v2); }
void Mm256FmaddPd(__m256d* r, __m256d* v0, __m256d* v1, __m256d* v2) { *r = _mm256_fmadd_pd(*v0, *v1, *v2); }
void Mm256FmsubPs(__m256* r, __m256* v0, __m256* v1, __m256* v2) { *r = _mm256_fmsub_ps(*v0, *v1, *v2); }
void Mm256FmsubPd(__m256d* r, __m256d* v0, __m256d* v1, __m256d* v2) { *r = _mm256_fmsub_pd(*v0, *v1, *v2); }
void Mm256FnmaddPs(__m256* r, __m256* v0, __m256* v1, __m256* v2) { *r = _mm256_fnmadd_ps(*v0, *v1, *v2); }
void Mm256FnmaddPd(__m256d* r, __m256d* v0, __m256d* v1, __m256d* v2) { *r = _mm256_fnmadd_pd(*v0, *v1, *v2); }
void Mm256FnmsubPs(__m256* r, __m256* v0, __m256* v1, __m256* v2) { *r = _mm256_fnmsub_ps(*v0, *v1, *v2); }
void Mm256FnmsubPd(__m256d* r, __m256d* v0, __m256d* v1, __m256d* v2) { *r = _mm256_fnmsub_pd(*v0, *v1, *v2); }
void Mm256FmaddsubPs(__m256* r, __m256* v0, __m256* v1, __m256* v2) { *r = _mm256_fmaddsub_ps(*v0, *v1, *v2); }
void Mm256FmaddsubPd(__m256d* r, __m256d* v0, __m256d* v1, __m256d* v2) { *r = _mm256_fmaddsub_pd(*v0, *v1, *v2); }
void Mm256FmsubaddPs(__m256* r, __m256* v0, __m256* v1, __m256* v2) { *r = _mm256_fmsubadd_ps(*v0, *v1, *v2); }
void Mm256FmsubaddPd(__m256d* r, __m256d* v0, __m256d* v1, __m256d* v2) { *r = _mm256_fmsubadd_pd(*v0, *v1, *v2); }
