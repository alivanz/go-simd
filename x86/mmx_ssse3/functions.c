#include <immintrin.h>

void MmAbsPi8(__m64* r, __m64* v0) { *r = _mm_abs_pi8(*v0); }
void MmAbsPi16(__m64* r, __m64* v0) { *r = _mm_abs_pi16(*v0); }
void MmAbsPi32(__m64* r, __m64* v0) { *r = _mm_abs_pi32(*v0); }
void MmHaddPi16(__m64* r, __m64* v0, __m64* v1) { *r = _mm_hadd_pi16(*v0, *v1); }
void MmHaddPi32(__m64* r, __m64* v0, __m64* v1) { *r = _mm_hadd_pi32(*v0, *v1); }
void MmHaddsPi16(__m64* r, __m64* v0, __m64* v1) { *r = _mm_hadds_pi16(*v0, *v1); }
void MmHsubPi16(__m64* r, __m64* v0, __m64* v1) { *r = _mm_hsub_pi16(*v0, *v1); }
void MmHsubPi32(__m64* r, __m64* v0, __m64* v1) { *r = _mm_hsub_pi32(*v0, *v1); }
void MmHsubsPi16(__m64* r, __m64* v0, __m64* v1) { *r = _mm_hsubs_pi16(*v0, *v1); }
void MmMaddubsPi16(__m64* r, __m64* v0, __m64* v1) { *r = _mm_maddubs_pi16(*v0, *v1); }
void MmMulhrsPi16(__m64* r, __m64* v0, __m64* v1) { *r = _mm_mulhrs_pi16(*v0, *v1); }
void MmShufflePi8(__m64* r, __m64* v0, __m64* v1) { *r = _mm_shuffle_pi8(*v0, *v1); }
void MmSignPi8(__m64* r, __m64* v0, __m64* v1) { *r = _mm_sign_pi8(*v0, *v1); }
void MmSignPi16(__m64* r, __m64* v0, __m64* v1) { *r = _mm_sign_pi16(*v0, *v1); }
void MmSignPi32(__m64* r, __m64* v0, __m64* v1) { *r = _mm_sign_pi32(*v0, *v1); }
