package f16c

import (
	"github.com/alivanz/go-simd/x86"
)

/*
#cgo CFLAGS: -mf16c
#include <immintrin.h>
*/
import "C"

// Convert the half-precision (16-bit) floating-point value "a" to a single-precision (32-bit) floating-point value, and store the result in "dst".
//
//go:linkname CvtshSs CvtshSs
//go:noescape
func CvtshSs(r *x86.Float, v0 *x86.Ushort)

// Convert packed half-precision (16-bit) floating-point elements in "a" to packed single-precision (32-bit) floating-point elements, and store the results in "dst".
//
//go:linkname MmCvtphPs MmCvtphPs
//go:noescape
func MmCvtphPs(r *x86.M128, v0 *x86.M128I)

// Convert packed half-precision (16-bit) floating-point elements in "a" to packed single-precision (32-bit) floating-point elements, and store the results in "dst".
//
//go:linkname Mm256CvtphPs Mm256CvtphPs
//go:noescape
func Mm256CvtphPs(r *x86.M256, v0 *x86.M128I)
