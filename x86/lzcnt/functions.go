package lzcnt

import (
	"github.com/alivanz/go-simd/x86"
)

/*
#cgo CFLAGS: -mlzcnt
#include <immintrin.h>
*/
import "C"

// __lzcnt32
//
//go:linkname Lzcnt32 Lzcnt32
//go:noescape
func Lzcnt32(r *x86.Uint, v0 *x86.Uint)

// Count the number of leading zero bits in unsigned 32-bit integer "a", and return that count in "dst".
//
//go:linkname LzcntU32 LzcntU32
//go:noescape
func LzcntU32(r *x86.Uint, v0 *x86.Uint)

// Count the number of leading zero bits in unsigned 64-bit integer "a", and return that count in "dst".
//
//go:linkname LzcntU64 LzcntU64
//go:noescape
func LzcntU64(r *x86.Ulonglong, v0 *x86.Ulonglong)
