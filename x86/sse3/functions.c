#include <immintrin.h>

void MmAddsubPs(__m128* r, __m128* v0, __m128* v1) { *r = _mm_addsub_ps(*v0, *v1); }
void MmHaddPs(__m128* r, __m128* v0, __m128* v1) { *r = _mm_hadd_ps(*v0, *v1); }
void MmHsubPs(__m128* r, __m128* v0, __m128* v1) { *r = _mm_hsub_ps(*v0, *v1); }
void MmMovehdupPs(__m128* r, __m128* v0) { *r = _mm_movehdup_ps(*v0); }
void MmMoveldupPs(__m128* r, __m128* v0) { *r = _mm_moveldup_ps(*v0); }
void MmAddsubPd(__m128d* r, __m128d* v0, __m128d* v1) { *r = _mm_addsub_pd(*v0, *v1); }
void MmHaddPd(__m128d* r, __m128d* v0, __m128d* v1) { *r = _mm_hadd_pd(*v0, *v1); }
void MmHsubPd(__m128d* r, __m128d* v0, __m128d* v1) { *r = _mm_hsub_pd(*v0, *v1); }
void MmMovedupPd(__m128d* r, __m128d* v0) { *r = _mm_movedup_pd(*v0); }
void MmMwait(unsigned* v0, unsigned* v1) { _mm_mwait(*v0, *v1); }
