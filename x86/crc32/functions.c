#include <immintrin.h>

void MmCrc32U8(unsigned int* r, unsigned int* v0, unsigned char* v1) { *r = _mm_crc32_u8(*v0, *v1); }
void MmCrc32U16(unsigned int* r, unsigned int* v0, unsigned short* v1) { *r = _mm_crc32_u16(*v0, *v1); }
void MmCrc32U32(unsigned int* r, unsigned int* v0, unsigned int* v1) { *r = _mm_crc32_u32(*v0, *v1); }
void MmCrc32U64(unsigned long long* r, unsigned long long* v0, unsigned long long* v1) { *r = _mm_crc32_u64(*v0, *v1); }
