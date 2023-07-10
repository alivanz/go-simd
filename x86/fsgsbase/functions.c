#include <immintrin.h>

void ReadfsbaseU32(unsigned int* r) { *r = _readfsbase_u32(); }
void ReadfsbaseU64(unsigned long long* r) { *r = _readfsbase_u64(); }
void ReadgsbaseU32(unsigned int* r) { *r = _readgsbase_u32(); }
void ReadgsbaseU64(unsigned long long* r) { *r = _readgsbase_u64(); }
void WritefsbaseU32(unsigned int* v0) { _writefsbase_u32(*v0); }
void WritefsbaseU64(unsigned long long* v0) { _writefsbase_u64(*v0); }
void WritegsbaseU32(unsigned int* v0) { _writegsbase_u32(*v0); }
void WritegsbaseU64(unsigned long long* v0) { _writegsbase_u64(*v0); }
