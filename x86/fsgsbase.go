package x86

/*
#cgo CFLAGS: -march=native
#include <immintrin.h>
*/
import "C"

// _readfsbase_u32
// __always_inline__, __nodebug__, __target__("fsgsbase")
func ReadfsbaseU32() Uint {
	return C._readfsbase_u32()
}

// _readfsbase_u64
// __always_inline__, __nodebug__, __target__("fsgsbase")
func ReadfsbaseU64() Ulonglong {
	return C._readfsbase_u64()
}

// _readgsbase_u32
// __always_inline__, __nodebug__, __target__("fsgsbase")
func ReadgsbaseU32() Uint {
	return C._readgsbase_u32()
}

// _readgsbase_u64
// __always_inline__, __nodebug__, __target__("fsgsbase")
func ReadgsbaseU64() Ulonglong {
	return C._readgsbase_u64()
}

// _writefsbase_u32
// __always_inline__, __nodebug__, __target__("fsgsbase")
func WritefsbaseU32(v0 Uint) {
	C._writefsbase_u32(v0)
}

// _writefsbase_u64
// __always_inline__, __nodebug__, __target__("fsgsbase")
func WritefsbaseU64(v0 Ulonglong) {
	C._writefsbase_u64(v0)
}

// _writegsbase_u32
// __always_inline__, __nodebug__, __target__("fsgsbase")
func WritegsbaseU32(v0 Uint) {
	C._writegsbase_u32(v0)
}

// _writegsbase_u64
// __always_inline__, __nodebug__, __target__("fsgsbase")
func WritegsbaseU64(v0 Ulonglong) {
	C._writegsbase_u64(v0)
}
