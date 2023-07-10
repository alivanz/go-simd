package fma

import (
	"github.com/alivanz/go-simd/x86"
)

/*
#cgo CFLAGS: -mfma
#include <immintrin.h>
*/
import "C"

// Multiply packed single-precision (32-bit) floating-point elements in "a" and "b", add the intermediate result to packed elements in "c", and store the results in "dst".
//
//go:linkname MmFmaddPs MmFmaddPs
//go:noescape
func MmFmaddPs(r *x86.M128, v0 *x86.M128, v1 *x86.M128, v2 *x86.M128)

// Multiply packed double-precision (64-bit) floating-point elements in "a" and "b", add the intermediate result to packed elements in "c", and store the results in "dst".
//
//go:linkname MmFmaddPd MmFmaddPd
//go:noescape
func MmFmaddPd(r *x86.M128D, v0 *x86.M128D, v1 *x86.M128D, v2 *x86.M128D)

// Multiply the lower single-precision (32-bit) floating-point elements in "a" and "b", and add the intermediate result to the lower element in "c". Store the result in the lower element of "dst", and copy the upper 3 packed elements from "a" to the upper elements of "dst".
//
//go:linkname MmFmaddSs MmFmaddSs
//go:noescape
func MmFmaddSs(r *x86.M128, v0 *x86.M128, v1 *x86.M128, v2 *x86.M128)

// Multiply the lower double-precision (64-bit) floating-point elements in "a" and "b", and add the intermediate result to the lower element in "c". Store the result in the lower element of "dst", and copy the upper element from "a" to the upper element of "dst".
//
//go:linkname MmFmaddSd MmFmaddSd
//go:noescape
func MmFmaddSd(r *x86.M128D, v0 *x86.M128D, v1 *x86.M128D, v2 *x86.M128D)

// Multiply packed single-precision (32-bit) floating-point elements in "a" and "b", subtract packed elements in "c" from the intermediate result, and store the results in "dst".
//
//go:linkname MmFmsubPs MmFmsubPs
//go:noescape
func MmFmsubPs(r *x86.M128, v0 *x86.M128, v1 *x86.M128, v2 *x86.M128)

// Multiply packed double-precision (64-bit) floating-point elements in "a" and "b", subtract packed elements in "c" from the intermediate result, and store the results in "dst".
//
//go:linkname MmFmsubPd MmFmsubPd
//go:noescape
func MmFmsubPd(r *x86.M128D, v0 *x86.M128D, v1 *x86.M128D, v2 *x86.M128D)

// Multiply the lower single-precision (32-bit) floating-point elements in "a" and "b", and subtract the lower element in "c" from the intermediate result. Store the result in the lower element of "dst", and copy the upper 3 packed elements from "a" to the upper elements of "dst".
//
//go:linkname MmFmsubSs MmFmsubSs
//go:noescape
func MmFmsubSs(r *x86.M128, v0 *x86.M128, v1 *x86.M128, v2 *x86.M128)

// Multiply the lower double-precision (64-bit) floating-point elements in "a" and "b", and subtract the lower element in "c" from the intermediate result. Store the result in the lower element of "dst", and copy the upper element from "a" to the upper element of "dst".
//
//go:linkname MmFmsubSd MmFmsubSd
//go:noescape
func MmFmsubSd(r *x86.M128D, v0 *x86.M128D, v1 *x86.M128D, v2 *x86.M128D)

// Multiply packed single-precision (32-bit) floating-point elements in "a" and "b", add the negated intermediate result to packed elements in "c", and store the results in "dst".
//
//go:linkname MmFnmaddPs MmFnmaddPs
//go:noescape
func MmFnmaddPs(r *x86.M128, v0 *x86.M128, v1 *x86.M128, v2 *x86.M128)

// Multiply packed double-precision (64-bit) floating-point elements in "a" and "b", add the negated intermediate result to packed elements in "c", and store the results in "dst".
//
//go:linkname MmFnmaddPd MmFnmaddPd
//go:noescape
func MmFnmaddPd(r *x86.M128D, v0 *x86.M128D, v1 *x86.M128D, v2 *x86.M128D)

// Multiply the lower single-precision (32-bit) floating-point elements in "a" and "b", and add the negated intermediate result to the lower element in "c". Store the result in the lower element of "dst", and copy the upper 3 packed elements from "a" to the upper elements of "dst".
//
//go:linkname MmFnmaddSs MmFnmaddSs
//go:noescape
func MmFnmaddSs(r *x86.M128, v0 *x86.M128, v1 *x86.M128, v2 *x86.M128)

// Multiply the lower double-precision (64-bit) floating-point elements in "a" and "b", and add the negated intermediate result to the lower element in "c". Store the result in the lower element of "dst", and copy the upper element from "a" to the upper element of "dst".
//
//go:linkname MmFnmaddSd MmFnmaddSd
//go:noescape
func MmFnmaddSd(r *x86.M128D, v0 *x86.M128D, v1 *x86.M128D, v2 *x86.M128D)

// Multiply packed single-precision (32-bit) floating-point elements in "a" and "b", subtract packed elements in "c" from the negated intermediate result, and store the results in "dst".
//
//go:linkname MmFnmsubPs MmFnmsubPs
//go:noescape
func MmFnmsubPs(r *x86.M128, v0 *x86.M128, v1 *x86.M128, v2 *x86.M128)

// Multiply packed double-precision (64-bit) floating-point elements in "a" and "b", subtract packed elements in "c" from the negated intermediate result, and store the results in "dst".
//
//go:linkname MmFnmsubPd MmFnmsubPd
//go:noescape
func MmFnmsubPd(r *x86.M128D, v0 *x86.M128D, v1 *x86.M128D, v2 *x86.M128D)

// Multiply the lower single-precision (32-bit) floating-point elements in "a" and "b", and subtract the lower element in "c" from the negated intermediate result. Store the result in the lower element of "dst", and copy the upper 3 packed elements from "a" to the upper elements of "dst".
//
//go:linkname MmFnmsubSs MmFnmsubSs
//go:noescape
func MmFnmsubSs(r *x86.M128, v0 *x86.M128, v1 *x86.M128, v2 *x86.M128)

// Multiply the lower double-precision (64-bit) floating-point elements in "a" and "b", and subtract the lower element in "c" from the negated intermediate result. Store the result in the lower element of "dst", and copy the upper element from "a" to the upper element of "dst".
//
//go:linkname MmFnmsubSd MmFnmsubSd
//go:noescape
func MmFnmsubSd(r *x86.M128D, v0 *x86.M128D, v1 *x86.M128D, v2 *x86.M128D)

// Multiply packed single-precision (32-bit) floating-point elements in "a" and "b", alternatively add and subtract packed elements in "c" to/from the intermediate result, and store the results in "dst".
//
//go:linkname MmFmaddsubPs MmFmaddsubPs
//go:noescape
func MmFmaddsubPs(r *x86.M128, v0 *x86.M128, v1 *x86.M128, v2 *x86.M128)

// Multiply packed double-precision (64-bit) floating-point elements in "a" and "b", alternatively add and subtract packed elements in "c" to/from the intermediate result, and store the results in "dst".
//
//go:linkname MmFmaddsubPd MmFmaddsubPd
//go:noescape
func MmFmaddsubPd(r *x86.M128D, v0 *x86.M128D, v1 *x86.M128D, v2 *x86.M128D)

// Multiply packed single-precision (32-bit) floating-point elements in "a" and "b", alternatively subtract and add packed elements in "c" from/to the intermediate result, and store the results in "dst".
//
//go:linkname MmFmsubaddPs MmFmsubaddPs
//go:noescape
func MmFmsubaddPs(r *x86.M128, v0 *x86.M128, v1 *x86.M128, v2 *x86.M128)

// Multiply packed double-precision (64-bit) floating-point elements in "a" and "b", alternatively subtract and add packed elements in "c" from/to the intermediate result, and store the results in "dst".
//
//go:linkname MmFmsubaddPd MmFmsubaddPd
//go:noescape
func MmFmsubaddPd(r *x86.M128D, v0 *x86.M128D, v1 *x86.M128D, v2 *x86.M128D)

// Multiply packed single-precision (32-bit) floating-point elements in "a" and "b", add the intermediate result to packed elements in "c", and store the results in "dst".
//
//go:linkname Mm256FmaddPs Mm256FmaddPs
//go:noescape
func Mm256FmaddPs(r *x86.M256, v0 *x86.M256, v1 *x86.M256, v2 *x86.M256)

// Multiply packed double-precision (64-bit) floating-point elements in "a" and "b", add the intermediate result to packed elements in "c", and store the results in "dst".
//
//go:linkname Mm256FmaddPd Mm256FmaddPd
//go:noescape
func Mm256FmaddPd(r *x86.M256D, v0 *x86.M256D, v1 *x86.M256D, v2 *x86.M256D)

// Multiply packed single-precision (32-bit) floating-point elements in "a" and "b", subtract packed elements in "c" from the intermediate result, and store the results in "dst".
//
//go:linkname Mm256FmsubPs Mm256FmsubPs
//go:noescape
func Mm256FmsubPs(r *x86.M256, v0 *x86.M256, v1 *x86.M256, v2 *x86.M256)

// Multiply packed double-precision (64-bit) floating-point elements in "a" and "b", subtract packed elements in "c" from the intermediate result, and store the results in "dst".
//
//go:linkname Mm256FmsubPd Mm256FmsubPd
//go:noescape
func Mm256FmsubPd(r *x86.M256D, v0 *x86.M256D, v1 *x86.M256D, v2 *x86.M256D)

// Multiply packed single-precision (32-bit) floating-point elements in "a" and "b", add the negated intermediate result to packed elements in "c", and store the results in "dst".
//
//go:linkname Mm256FnmaddPs Mm256FnmaddPs
//go:noescape
func Mm256FnmaddPs(r *x86.M256, v0 *x86.M256, v1 *x86.M256, v2 *x86.M256)

// Multiply packed double-precision (64-bit) floating-point elements in "a" and "b", add the negated intermediate result to packed elements in "c", and store the results in "dst".
//
//go:linkname Mm256FnmaddPd Mm256FnmaddPd
//go:noescape
func Mm256FnmaddPd(r *x86.M256D, v0 *x86.M256D, v1 *x86.M256D, v2 *x86.M256D)

// Multiply packed single-precision (32-bit) floating-point elements in "a" and "b", subtract packed elements in "c" from the negated intermediate result, and store the results in "dst".
//
//go:linkname Mm256FnmsubPs Mm256FnmsubPs
//go:noescape
func Mm256FnmsubPs(r *x86.M256, v0 *x86.M256, v1 *x86.M256, v2 *x86.M256)

// Multiply packed double-precision (64-bit) floating-point elements in "a" and "b", subtract packed elements in "c" from the negated intermediate result, and store the results in "dst".
//
//go:linkname Mm256FnmsubPd Mm256FnmsubPd
//go:noescape
func Mm256FnmsubPd(r *x86.M256D, v0 *x86.M256D, v1 *x86.M256D, v2 *x86.M256D)

// Multiply packed single-precision (32-bit) floating-point elements in "a" and "b", alternatively add and subtract packed elements in "c" to/from the intermediate result, and store the results in "dst".
//
//go:linkname Mm256FmaddsubPs Mm256FmaddsubPs
//go:noescape
func Mm256FmaddsubPs(r *x86.M256, v0 *x86.M256, v1 *x86.M256, v2 *x86.M256)

// Multiply packed double-precision (64-bit) floating-point elements in "a" and "b", alternatively add and subtract packed elements in "c" to/from the intermediate result, and store the results in "dst".
//
//go:linkname Mm256FmaddsubPd Mm256FmaddsubPd
//go:noescape
func Mm256FmaddsubPd(r *x86.M256D, v0 *x86.M256D, v1 *x86.M256D, v2 *x86.M256D)

// Multiply packed single-precision (32-bit) floating-point elements in "a" and "b", alternatively subtract and add packed elements in "c" from/to the intermediate result, and store the results in "dst".
//
//go:linkname Mm256FmsubaddPs Mm256FmsubaddPs
//go:noescape
func Mm256FmsubaddPs(r *x86.M256, v0 *x86.M256, v1 *x86.M256, v2 *x86.M256)

// Multiply packed double-precision (64-bit) floating-point elements in "a" and "b", alternatively subtract and add packed elements in "c" from/to the intermediate result, and store the results in "dst".
//
//go:linkname Mm256FmsubaddPd Mm256FmsubaddPd
//go:noescape
func Mm256FmsubaddPd(r *x86.M256D, v0 *x86.M256D, v1 *x86.M256D, v2 *x86.M256D)
