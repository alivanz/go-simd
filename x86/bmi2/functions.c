#include <immintrin.h>

void BzhiU32(unsigned int* r, unsigned int* v0, unsigned int* v1) { *r = _bzhi_u32(*v0, *v1); }
void PdepU32(unsigned int* r, unsigned int* v0, unsigned int* v1) { *r = _pdep_u32(*v0, *v1); }
void PextU32(unsigned int* r, unsigned int* v0, unsigned int* v1) { *r = _pext_u32(*v0, *v1); }
void BzhiU64(unsigned long long* r, unsigned long long* v0, unsigned long long* v1) { *r = _bzhi_u64(*v0, *v1); }
void PdepU64(unsigned long long* r, unsigned long long* v0, unsigned long long* v1) { *r = _pdep_u64(*v0, *v1); }
void PextU64(unsigned long long* r, unsigned long long* v0, unsigned long long* v1) { *r = _pext_u64(*v0, *v1); }
