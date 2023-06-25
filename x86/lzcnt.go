package x86

/*
#cgo CFLAGS: -march=native
#include <immintrin.h>
*/
import "C"

// __lzcnt32
func Lzcnt32(v0 Uint) Uint {
	return C.__lzcnt32(v0)
}

// Count the number of leading zero bits in unsigned 32-bit integer "a", and return that count in "dst".
func LzcntU32(v0 Uint) Uint {
	return C._lzcnt_u32(v0)
}

// Count the number of leading zero bits in unsigned 64-bit integer "a", and return that count in "dst".
func LzcntU64(v0 Ulonglong) Ulonglong {
	return C._lzcnt_u64(v0)
}
