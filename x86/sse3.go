package x86

/*
#cgo CFLAGS: -march=native
#include <immintrin.h>
*/
import "C"

// Alternatively add and subtract packed single-precision (32-bit) floating-point elements in "a" to/from packed elements in "b", and store the results in "dst".
func MmAddsubPs(v0 M128, v1 M128) M128 {
	return C._mm_addsub_ps(v0, v1)
}

// Horizontally add adjacent pairs of single-precision (32-bit) floating-point elements in "a" and "b", and pack the results in "dst".
func MmHaddPs(v0 M128, v1 M128) M128 {
	return C._mm_hadd_ps(v0, v1)
}

// Horizontally add adjacent pairs of single-precision (32-bit) floating-point elements in "a" and "b", and pack the results in "dst".
func MmHsubPs(v0 M128, v1 M128) M128 {
	return C._mm_hsub_ps(v0, v1)
}

// Duplicate odd-indexed single-precision (32-bit) floating-point elements from "a", and store the results in "dst".
func MmMovehdupPs(v0 M128) M128 {
	return C._mm_movehdup_ps(v0)
}

// Duplicate even-indexed single-precision (32-bit) floating-point elements from "a", and store the results in "dst".
func MmMoveldupPs(v0 M128) M128 {
	return C._mm_moveldup_ps(v0)
}

// Alternatively add and subtract packed double-precision (64-bit) floating-point elements in "a" to/from packed elements in "b", and store the results in "dst".
func MmAddsubPd(v0 M128D, v1 M128D) M128D {
	return C._mm_addsub_pd(v0, v1)
}

// Horizontally add adjacent pairs of double-precision (64-bit) floating-point elements in "a" and "b", and pack the results in "dst".
func MmHaddPd(v0 M128D, v1 M128D) M128D {
	return C._mm_hadd_pd(v0, v1)
}

// Horizontally subtract adjacent pairs of double-precision (64-bit) floating-point elements in "a" and "b", and pack the results in "dst".
func MmHsubPd(v0 M128D, v1 M128D) M128D {
	return C._mm_hsub_pd(v0, v1)
}

// Duplicate the low double-precision (64-bit) floating-point element from "a", and store the results in "dst".
func MmMovedupPd(v0 M128D) M128D {
	return C._mm_movedup_pd(v0)
}

// Hint to the processor that it can enter an implementation-dependent-optimized state while waiting for an event or store operation to the address range specified by MONITOR.
func MmMwait(v0 Unsigned, v1 Unsigned) {
	C._mm_mwait(v0, v1)
}
