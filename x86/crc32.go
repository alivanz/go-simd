package x86

/*
#cgo CFLAGS: -march=native
#include <immintrin.h>
*/
import "C"

// Starting with the initial value in "crc", accumulates a CRC32 value for unsigned 8-bit integer "v", and stores the result in "dst".
func MmCrc32U8(v0 Uint, v1 Uchar) Uint {
	return C._mm_crc32_u8(v0, v1)
}

// Starting with the initial value in "crc", accumulates a CRC32 value for unsigned 16-bit integer "v", and stores the result in "dst".
func MmCrc32U16(v0 Uint, v1 Ushort) Uint {
	return C._mm_crc32_u16(v0, v1)
}

// Starting with the initial value in "crc", accumulates a CRC32 value for unsigned 32-bit integer "v", and stores the result in "dst".
func MmCrc32U32(v0 Uint, v1 Uint) Uint {
	return C._mm_crc32_u32(v0, v1)
}

// Starting with the initial value in "crc", accumulates a CRC32 value for unsigned 64-bit integer "v", and stores the result in "dst".
func MmCrc32U64(v0 Ulonglong, v1 Ulonglong) Ulonglong {
	return C._mm_crc32_u64(v0, v1)
}
