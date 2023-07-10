package popcnt

import (
	"github.com/alivanz/go-simd/x86"
)

/*
#cgo CFLAGS: -mpopcnt
#include <immintrin.h>
*/
import "C"

// Count the number of bits set to 1 in unsigned 32-bit integer "a", and return that count in "dst".
//
//go:linkname MmPopcntU32 MmPopcntU32
//go:noescape
func MmPopcntU32(r *x86.Int, v0 *x86.Uint)

// Count the number of bits set to 1 in unsigned 64-bit integer "a", and return that count in "dst".
//
//go:linkname MmPopcntU64 MmPopcntU64
//go:noescape
func MmPopcntU64(r *x86.Longlong, v0 *x86.Ulonglong)
