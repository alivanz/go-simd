package bmi2

import (
	"github.com/alivanz/go-simd/x86"
)

/*
#cgo CFLAGS: -mbmi2
#include <immintrin.h>
*/
import "C"

// Copy all bits from unsigned 32-bit integer "a" to "dst", and reset (set to 0) the high bits in "dst" starting at "index".
//
//go:linkname BzhiU32 BzhiU32
//go:noescape
func BzhiU32(r *x86.Uint, v0 *x86.Uint, v1 *x86.Uint)

// Deposit contiguous low bits from unsigned 32-bit integer "a" to "dst" at the corresponding bit locations specified by "mask"; all other bits in "dst" are set to zero.
//
//go:linkname PdepU32 PdepU32
//go:noescape
func PdepU32(r *x86.Uint, v0 *x86.Uint, v1 *x86.Uint)

// Extract bits from unsigned 32-bit integer "a" at the corresponding bit locations specified by "mask" to contiguous low bits in "dst"; the remaining upper bits in "dst" are set to zero.
//
//go:linkname PextU32 PextU32
//go:noescape
func PextU32(r *x86.Uint, v0 *x86.Uint, v1 *x86.Uint)

// Copy all bits from unsigned 64-bit integer "a" to "dst", and reset (set to 0) the high bits in "dst" starting at "index".
//
//go:linkname BzhiU64 BzhiU64
//go:noescape
func BzhiU64(r *x86.Ulonglong, v0 *x86.Ulonglong, v1 *x86.Ulonglong)

// Deposit contiguous low bits from unsigned 64-bit integer "a" to "dst" at the corresponding bit locations specified by "mask"; all other bits in "dst" are set to zero.
//
//go:linkname PdepU64 PdepU64
//go:noescape
func PdepU64(r *x86.Ulonglong, v0 *x86.Ulonglong, v1 *x86.Ulonglong)

// Extract bits from unsigned 64-bit integer "a" at the corresponding bit locations specified by "mask" to contiguous low bits in "dst"; the remaining upper bits in "dst" are set to zero.
//
//go:linkname PextU64 PextU64
//go:noescape
func PextU64(r *x86.Ulonglong, v0 *x86.Ulonglong, v1 *x86.Ulonglong)
