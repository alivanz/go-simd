package x86

/*
#cgo CFLAGS: -march=native
#include <immintrin.h>
*/
import "C"

// _bzhi_u32
// __always_inline__, __nodebug__, __target__("bmi2")
func BzhiU32(v0 Uint, v1 Uint) Uint {
	return C._bzhi_u32(v0, v1)
}

// _pdep_u32
// __always_inline__, __nodebug__, __target__("bmi2")
func PdepU32(v0 Uint, v1 Uint) Uint {
	return C._pdep_u32(v0, v1)
}

// _pext_u32
// __always_inline__, __nodebug__, __target__("bmi2")
func PextU32(v0 Uint, v1 Uint) Uint {
	return C._pext_u32(v0, v1)
}

// _bzhi_u64
// __always_inline__, __nodebug__, __target__("bmi2")
func BzhiU64(v0 Ulonglong, v1 Ulonglong) Ulonglong {
	return C._bzhi_u64(v0, v1)
}

// _pdep_u64
// __always_inline__, __nodebug__, __target__("bmi2")
func PdepU64(v0 Ulonglong, v1 Ulonglong) Ulonglong {
	return C._pdep_u64(v0, v1)
}

// _pext_u64
// __always_inline__, __nodebug__, __target__("bmi2")
func PextU64(v0 Ulonglong, v1 Ulonglong) Ulonglong {
	return C._pext_u64(v0, v1)
}
