package x86

/*
#cgo CFLAGS: -march=native
#include <immintrin.h>
*/
import "C"

// _mm_popcnt_u32
// __always_inline__, __nodebug__, __target__("popcnt")
func MmPopcntU32(v0 Uint) Int {
	return C._mm_popcnt_u32(v0)
}

// _mm_popcnt_u64
// __always_inline__, __nodebug__, __target__("popcnt")
func MmPopcntU64(v0 Ulonglong) Longlong {
	return C._mm_popcnt_u64(v0)
}
