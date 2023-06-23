package x86

/*
#cgo CFLAGS: -march=native
#include <immintrin.h>
*/
import "C"

// _mm_crc32_u8
// __always_inline__, __nodebug__, __target__("crc32")
func MmCrc32U8(v0 Uint, v1 Uchar) Uint {
	return C._mm_crc32_u8(v0, v1)
}

// _mm_crc32_u16
// __always_inline__, __nodebug__, __target__("crc32")
func MmCrc32U16(v0 Uint, v1 Ushort) Uint {
	return C._mm_crc32_u16(v0, v1)
}

// _mm_crc32_u32
// __always_inline__, __nodebug__, __target__("crc32")
func MmCrc32U32(v0 Uint, v1 Uint) Uint {
	return C._mm_crc32_u32(v0, v1)
}

// _mm_crc32_u64
// __always_inline__, __nodebug__, __target__("crc32")
func MmCrc32U64(v0 Ulonglong, v1 Ulonglong) Ulonglong {
	return C._mm_crc32_u64(v0, v1)
}
