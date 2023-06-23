package x86

/*
#cgo CFLAGS: -march=native
#include <immintrin.h>
*/
import "C"

// __lzcnt32
// __always_inline__, __nodebug__, __target__("lzcnt")
func Lzcnt32(v0 Uint) Uint {
	return C.__lzcnt32(v0)
}

// _lzcnt_u32
// __always_inline__, __nodebug__, __target__("lzcnt")
func LzcntU32(v0 Uint) Uint {
	return C._lzcnt_u32(v0)
}

// _lzcnt_u64
// __always_inline__, __nodebug__, __target__("lzcnt")
func LzcntU64(v0 Ulonglong) Ulonglong {
	return C._lzcnt_u64(v0)
}
