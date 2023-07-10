#include <immintrin.h>

void AndnU32(unsigned int* r, unsigned int* v0, unsigned int* v1) { *r = __andn_u32(*v0, *v1); }
void BextrU32(unsigned int* r, unsigned int* v0, unsigned int* v1) { *r = __bextr_u32(*v0, *v1); }
void BextrU32(unsigned int* r, unsigned int* v0, unsigned int* v1, unsigned int* v2) { *r = _bextr_u32(*v0, *v1, *v2); }
void Bextr2U32(unsigned int* r, unsigned int* v0, unsigned int* v1) { *r = _bextr2_u32(*v0, *v1); }
void BlsiU32(unsigned int* r, unsigned int* v0) { *r = __blsi_u32(*v0); }
void BlsmskU32(unsigned int* r, unsigned int* v0) { *r = __blsmsk_u32(*v0); }
void BlsrU32(unsigned int* r, unsigned int* v0) { *r = __blsr_u32(*v0); }
void AndnU64(unsigned long long* r, unsigned long long* v0, unsigned long long* v1) { *r = __andn_u64(*v0, *v1); }
void BextrU64(unsigned long long* r, unsigned long long* v0, unsigned long long* v1) { *r = __bextr_u64(*v0, *v1); }
void BextrU64(unsigned long long* r, unsigned long long* v0, unsigned int* v1, unsigned int* v2) { *r = _bextr_u64(*v0, *v1, *v2); }
void Bextr2U64(unsigned long long* r, unsigned long long* v0, unsigned long long* v1) { *r = _bextr2_u64(*v0, *v1); }
void BlsiU64(unsigned long long* r, unsigned long long* v0) { *r = __blsi_u64(*v0); }
void BlsmskU64(unsigned long long* r, unsigned long long* v0) { *r = __blsmsk_u64(*v0); }
void BlsrU64(unsigned long long* r, unsigned long long* v0) { *r = __blsr_u64(*v0); }
