package sse3

import (
	"github.com/alivanz/go-simd/x86"
)

/*
#cgo CFLAGS: -msse3
#include <immintrin.h>
*/
import "C"

// Alternatively add and subtract packed single-precision (32-bit) floating-point elements in "a" to/from packed elements in "b", and store the results in "dst".
//
//go:linkname MmAddsubPs MmAddsubPs
//go:noescape
func MmAddsubPs(r *x86.M128, v0 *x86.M128, v1 *x86.M128)

// Horizontally add adjacent pairs of single-precision (32-bit) floating-point elements in "a" and "b", and pack the results in "dst".
//
//go:linkname MmHaddPs MmHaddPs
//go:noescape
func MmHaddPs(r *x86.M128, v0 *x86.M128, v1 *x86.M128)

// Horizontally add adjacent pairs of single-precision (32-bit) floating-point elements in "a" and "b", and pack the results in "dst".
//
//go:linkname MmHsubPs MmHsubPs
//go:noescape
func MmHsubPs(r *x86.M128, v0 *x86.M128, v1 *x86.M128)

// Duplicate odd-indexed single-precision (32-bit) floating-point elements from "a", and store the results in "dst".
//
//go:linkname MmMovehdupPs MmMovehdupPs
//go:noescape
func MmMovehdupPs(r *x86.M128, v0 *x86.M128)

// Duplicate even-indexed single-precision (32-bit) floating-point elements from "a", and store the results in "dst".
//
//go:linkname MmMoveldupPs MmMoveldupPs
//go:noescape
func MmMoveldupPs(r *x86.M128, v0 *x86.M128)

// Alternatively add and subtract packed double-precision (64-bit) floating-point elements in "a" to/from packed elements in "b", and store the results in "dst".
//
//go:linkname MmAddsubPd MmAddsubPd
//go:noescape
func MmAddsubPd(r *x86.M128D, v0 *x86.M128D, v1 *x86.M128D)

// Horizontally add adjacent pairs of double-precision (64-bit) floating-point elements in "a" and "b", and pack the results in "dst".
//
//go:linkname MmHaddPd MmHaddPd
//go:noescape
func MmHaddPd(r *x86.M128D, v0 *x86.M128D, v1 *x86.M128D)

// Horizontally subtract adjacent pairs of double-precision (64-bit) floating-point elements in "a" and "b", and pack the results in "dst".
//
//go:linkname MmHsubPd MmHsubPd
//go:noescape
func MmHsubPd(r *x86.M128D, v0 *x86.M128D, v1 *x86.M128D)

// Duplicate the low double-precision (64-bit) floating-point element from "a", and store the results in "dst".
//
//go:linkname MmMovedupPd MmMovedupPd
//go:noescape
func MmMovedupPd(r *x86.M128D, v0 *x86.M128D)

// Hint to the processor that it can enter an implementation-dependent-optimized state while waiting for an event or store operation to the address range specified by MONITOR.
//
//go:linkname MmMwait MmMwait
//go:noescape
func MmMwait(v0 *x86.Unsigned, v1 *x86.Unsigned)
