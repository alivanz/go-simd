package x86

/*
#cgo CFLAGS: -march=native
#include <immintrin.h>
*/
import "C"

// Copy all bits from unsigned 32-bit integer "a" to "dst", and reset (set to 0) the high bits in "dst" starting at "index".
func BzhiU32(v0 Uint, v1 Uint) Uint {
	return C._bzhi_u32(v0, v1)
}

// Deposit contiguous low bits from unsigned 32-bit integer "a" to "dst" at the corresponding bit locations specified by "mask"; all other bits in "dst" are set to zero.
func PdepU32(v0 Uint, v1 Uint) Uint {
	return C._pdep_u32(v0, v1)
}

// Extract bits from unsigned 32-bit integer "a" at the corresponding bit locations specified by "mask" to contiguous low bits in "dst"; the remaining upper bits in "dst" are set to zero.
func PextU32(v0 Uint, v1 Uint) Uint {
	return C._pext_u32(v0, v1)
}

// Copy all bits from unsigned 64-bit integer "a" to "dst", and reset (set to 0) the high bits in "dst" starting at "index".
func BzhiU64(v0 Ulonglong, v1 Ulonglong) Ulonglong {
	return C._bzhi_u64(v0, v1)
}

// Deposit contiguous low bits from unsigned 64-bit integer "a" to "dst" at the corresponding bit locations specified by "mask"; all other bits in "dst" are set to zero.
func PdepU64(v0 Ulonglong, v1 Ulonglong) Ulonglong {
	return C._pdep_u64(v0, v1)
}

// Extract bits from unsigned 64-bit integer "a" at the corresponding bit locations specified by "mask" to contiguous low bits in "dst"; the remaining upper bits in "dst" are set to zero.
func PextU64(v0 Ulonglong, v1 Ulonglong) Ulonglong {
	return C._pext_u64(v0, v1)
}
