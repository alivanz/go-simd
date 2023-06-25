package x86

/*
#cgo CFLAGS: -march=native
#include <immintrin.h>
*/
import "C"

// Read the FS segment base register and store the 32-bit result in "dst".
func ReadfsbaseU32() Uint {
	return C._readfsbase_u32()
}

// Read the FS segment base register and store the 64-bit result in "dst".
func ReadfsbaseU64() Ulonglong {
	return C._readfsbase_u64()
}

// Read the GS segment base register and store the 32-bit result in "dst".
func ReadgsbaseU32() Uint {
	return C._readgsbase_u32()
}

// Read the GS segment base register and store the 64-bit result in "dst".
func ReadgsbaseU64() Ulonglong {
	return C._readgsbase_u64()
}

// Write the unsigned 32-bit integer "a" to the FS segment base register.
func WritefsbaseU32(v0 Uint) {
	C._writefsbase_u32(v0)
}

// Write the unsigned 64-bit integer "a" to the FS segment base register.
func WritefsbaseU64(v0 Ulonglong) {
	C._writefsbase_u64(v0)
}

// Write the unsigned 32-bit integer "a" to the GS segment base register.
func WritegsbaseU32(v0 Uint) {
	C._writegsbase_u32(v0)
}

// Write the unsigned 64-bit integer "a" to the GS segment base register.
func WritegsbaseU64(v0 Ulonglong) {
	C._writegsbase_u64(v0)
}
