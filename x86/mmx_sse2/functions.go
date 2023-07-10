package mmx_sse2

import (
	"github.com/alivanz/go-simd/x86"
)

/*
#cgo CFLAGS: -mmmx -msse2
#include <immintrin.h>
*/
import "C"

// Convert packed double-precision (64-bit) floating-point elements in "a" to packed 32-bit integers, and store the results in "dst".
//
//go:linkname MmCvtpdPi32 MmCvtpdPi32
//go:noescape
func MmCvtpdPi32(r *x86.M64, v0 *x86.M128D)

// Convert packed double-precision (64-bit) floating-point elements in "a" to packed 32-bit integers with truncation, and store the results in "dst".
//
//go:linkname MmCvttpdPi32 MmCvttpdPi32
//go:noescape
func MmCvttpdPi32(r *x86.M64, v0 *x86.M128D)

// Convert packed signed 32-bit integers in "a" to packed double-precision (64-bit) floating-point elements, and store the results in "dst".
//
//go:linkname MmCvtpi32Pd MmCvtpi32Pd
//go:noescape
func MmCvtpi32Pd(r *x86.M128D, v0 *x86.M64)

// Add 64-bit integers "a" and "b", and store the result in "dst".
//
//go:linkname MmAddSi64 MmAddSi64
//go:noescape
func MmAddSi64(r *x86.M64, v0 *x86.M64, v1 *x86.M64)

// Multiply the low unsigned 32-bit integers from "a" and "b", and store the unsigned 64-bit result in "dst".
//
//go:linkname MmMulSu32 MmMulSu32
//go:noescape
func MmMulSu32(r *x86.M64, v0 *x86.M64, v1 *x86.M64)

// Subtract 64-bit integer "b" from 64-bit integer "a", and store the result in "dst".
//
//go:linkname MmSubSi64 MmSubSi64
//go:noescape
func MmSubSi64(r *x86.M64, v0 *x86.M64, v1 *x86.M64)
