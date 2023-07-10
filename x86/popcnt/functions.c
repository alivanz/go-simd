#include <immintrin.h>

void MmPopcntU32(int* r, unsigned int* v0) { *r = _mm_popcnt_u32(*v0); }
void MmPopcntU64(long long* r, unsigned long long* v0) { *r = _mm_popcnt_u64(*v0); }
