package crc32

import (
	"github.com/alivanz/go-simd/x86"
)

/*
#cgo CFLAGS: -mcrc32
#include <immintrin.h>
*/
import "C"

// Starting with the initial value in "crc", accumulates a CRC32 value for unsigned 8-bit integer "v", and stores the result in "dst".
//
//go:linkname MmCrc32U8 MmCrc32U8
//go:noescape
func MmCrc32U8(r *x86.Uint, v0 *x86.Uint, v1 *x86.Uchar)

// Starting with the initial value in "crc", accumulates a CRC32 value for unsigned 16-bit integer "v", and stores the result in "dst".
//
//go:linkname MmCrc32U16 MmCrc32U16
//go:noescape
func MmCrc32U16(r *x86.Uint, v0 *x86.Uint, v1 *x86.Ushort)

// Starting with the initial value in "crc", accumulates a CRC32 value for unsigned 32-bit integer "v", and stores the result in "dst".
//
//go:linkname MmCrc32U32 MmCrc32U32
//go:noescape
func MmCrc32U32(r *x86.Uint, v0 *x86.Uint, v1 *x86.Uint)

// Starting with the initial value in "crc", accumulates a CRC32 value for unsigned 64-bit integer "v", and stores the result in "dst".
//
//go:linkname MmCrc32U64 MmCrc32U64
//go:noescape
func MmCrc32U64(r *x86.Ulonglong, v0 *x86.Ulonglong, v1 *x86.Ulonglong)
