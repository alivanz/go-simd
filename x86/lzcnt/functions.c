#include <immintrin.h>

void Lzcnt32(unsigned int* r, unsigned int* v0) { *r = __lzcnt32(*v0); }
void LzcntU32(unsigned int* r, unsigned int* v0) { *r = _lzcnt_u32(*v0); }
void LzcntU64(unsigned long long* r, unsigned long long* v0) { *r = _lzcnt_u64(*v0); }
