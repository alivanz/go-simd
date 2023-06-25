package x86

/*
#cgo CFLAGS: -march=native
#include <immintrin.h>
*/
import "C"

// Count the number of bits set to 1 in unsigned 32-bit integer "a", and return that count in "dst".
func MmPopcntU32(v0 Uint) Int {
	return C._mm_popcnt_u32(v0)
}

// Count the number of bits set to 1 in unsigned 64-bit integer "a", and return that count in "dst".
func MmPopcntU64(v0 Ulonglong) Longlong {
	return C._mm_popcnt_u64(v0)
}
