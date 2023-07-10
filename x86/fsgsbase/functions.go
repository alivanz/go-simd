package fsgsbase

import (
	"github.com/alivanz/go-simd/x86"
)

/*
#cgo CFLAGS: -mfsgsbase
#include <immintrin.h>
*/
import "C"

// Read the FS segment base register and store the 32-bit result in "dst".
//
//go:linkname ReadfsbaseU32 ReadfsbaseU32
//go:noescape
func ReadfsbaseU32(r *x86.Uint, )

// Read the FS segment base register and store the 64-bit result in "dst".
//
//go:linkname ReadfsbaseU64 ReadfsbaseU64
//go:noescape
func ReadfsbaseU64(r *x86.Ulonglong, )

// Read the GS segment base register and store the 32-bit result in "dst".
//
//go:linkname ReadgsbaseU32 ReadgsbaseU32
//go:noescape
func ReadgsbaseU32(r *x86.Uint, )

// Read the GS segment base register and store the 64-bit result in "dst".
//
//go:linkname ReadgsbaseU64 ReadgsbaseU64
//go:noescape
func ReadgsbaseU64(r *x86.Ulonglong, )

// Write the unsigned 32-bit integer "a" to the FS segment base register.
//
//go:linkname WritefsbaseU32 WritefsbaseU32
//go:noescape
func WritefsbaseU32(v0 *x86.Uint)

// Write the unsigned 64-bit integer "a" to the FS segment base register.
//
//go:linkname WritefsbaseU64 WritefsbaseU64
//go:noescape
func WritefsbaseU64(v0 *x86.Ulonglong)

// Write the unsigned 32-bit integer "a" to the GS segment base register.
//
//go:linkname WritegsbaseU32 WritegsbaseU32
//go:noescape
func WritegsbaseU32(v0 *x86.Uint)

// Write the unsigned 64-bit integer "a" to the GS segment base register.
//
//go:linkname WritegsbaseU64 WritegsbaseU64
//go:noescape
func WritegsbaseU64(v0 *x86.Ulonglong)
