package x86

/*
#cgo CFLAGS: -march=native
#include <immintrin.h>
*/
import "C"

// Multiply packed single-precision (32-bit) floating-point elements in "a" and "b", add the intermediate result to packed elements in "c", and store the results in "dst".
func MmFmaddPs(v0 M128, v1 M128, v2 M128) M128 {
	return C._mm_fmadd_ps(v0, v1, v2)
}

// Multiply packed double-precision (64-bit) floating-point elements in "a" and "b", add the intermediate result to packed elements in "c", and store the results in "dst".
func MmFmaddPd(v0 M128D, v1 M128D, v2 M128D) M128D {
	return C._mm_fmadd_pd(v0, v1, v2)
}

// Multiply the lower single-precision (32-bit) floating-point elements in "a" and "b", and add the intermediate result to the lower element in "c". Store the result in the lower element of "dst", and copy the upper 3 packed elements from "a" to the upper elements of "dst".
func MmFmaddSs(v0 M128, v1 M128, v2 M128) M128 {
	return C._mm_fmadd_ss(v0, v1, v2)
}

// Multiply the lower double-precision (64-bit) floating-point elements in "a" and "b", and add the intermediate result to the lower element in "c". Store the result in the lower element of "dst", and copy the upper element from "a" to the upper element of "dst".
func MmFmaddSd(v0 M128D, v1 M128D, v2 M128D) M128D {
	return C._mm_fmadd_sd(v0, v1, v2)
}

// Multiply packed single-precision (32-bit) floating-point elements in "a" and "b", subtract packed elements in "c" from the intermediate result, and store the results in "dst".
func MmFmsubPs(v0 M128, v1 M128, v2 M128) M128 {
	return C._mm_fmsub_ps(v0, v1, v2)
}

// Multiply packed double-precision (64-bit) floating-point elements in "a" and "b", subtract packed elements in "c" from the intermediate result, and store the results in "dst".
func MmFmsubPd(v0 M128D, v1 M128D, v2 M128D) M128D {
	return C._mm_fmsub_pd(v0, v1, v2)
}

// Multiply the lower single-precision (32-bit) floating-point elements in "a" and "b", and subtract the lower element in "c" from the intermediate result. Store the result in the lower element of "dst", and copy the upper 3 packed elements from "a" to the upper elements of "dst".
func MmFmsubSs(v0 M128, v1 M128, v2 M128) M128 {
	return C._mm_fmsub_ss(v0, v1, v2)
}

// Multiply the lower double-precision (64-bit) floating-point elements in "a" and "b", and subtract the lower element in "c" from the intermediate result. Store the result in the lower element of "dst", and copy the upper element from "a" to the upper element of "dst".
func MmFmsubSd(v0 M128D, v1 M128D, v2 M128D) M128D {
	return C._mm_fmsub_sd(v0, v1, v2)
}

// Multiply packed single-precision (32-bit) floating-point elements in "a" and "b", add the negated intermediate result to packed elements in "c", and store the results in "dst".
func MmFnmaddPs(v0 M128, v1 M128, v2 M128) M128 {
	return C._mm_fnmadd_ps(v0, v1, v2)
}

// Multiply packed double-precision (64-bit) floating-point elements in "a" and "b", add the negated intermediate result to packed elements in "c", and store the results in "dst".
func MmFnmaddPd(v0 M128D, v1 M128D, v2 M128D) M128D {
	return C._mm_fnmadd_pd(v0, v1, v2)
}

// Multiply the lower single-precision (32-bit) floating-point elements in "a" and "b", and add the negated intermediate result to the lower element in "c". Store the result in the lower element of "dst", and copy the upper 3 packed elements from "a" to the upper elements of "dst".
func MmFnmaddSs(v0 M128, v1 M128, v2 M128) M128 {
	return C._mm_fnmadd_ss(v0, v1, v2)
}

// Multiply the lower double-precision (64-bit) floating-point elements in "a" and "b", and add the negated intermediate result to the lower element in "c". Store the result in the lower element of "dst", and copy the upper element from "a" to the upper element of "dst".
func MmFnmaddSd(v0 M128D, v1 M128D, v2 M128D) M128D {
	return C._mm_fnmadd_sd(v0, v1, v2)
}

// Multiply packed single-precision (32-bit) floating-point elements in "a" and "b", subtract packed elements in "c" from the negated intermediate result, and store the results in "dst".
func MmFnmsubPs(v0 M128, v1 M128, v2 M128) M128 {
	return C._mm_fnmsub_ps(v0, v1, v2)
}

// Multiply packed double-precision (64-bit) floating-point elements in "a" and "b", subtract packed elements in "c" from the negated intermediate result, and store the results in "dst".
func MmFnmsubPd(v0 M128D, v1 M128D, v2 M128D) M128D {
	return C._mm_fnmsub_pd(v0, v1, v2)
}

// Multiply the lower single-precision (32-bit) floating-point elements in "a" and "b", and subtract the lower element in "c" from the negated intermediate result. Store the result in the lower element of "dst", and copy the upper 3 packed elements from "a" to the upper elements of "dst".
func MmFnmsubSs(v0 M128, v1 M128, v2 M128) M128 {
	return C._mm_fnmsub_ss(v0, v1, v2)
}

// Multiply the lower double-precision (64-bit) floating-point elements in "a" and "b", and subtract the lower element in "c" from the negated intermediate result. Store the result in the lower element of "dst", and copy the upper element from "a" to the upper element of "dst".
func MmFnmsubSd(v0 M128D, v1 M128D, v2 M128D) M128D {
	return C._mm_fnmsub_sd(v0, v1, v2)
}

// Multiply packed single-precision (32-bit) floating-point elements in "a" and "b", alternatively add and subtract packed elements in "c" to/from the intermediate result, and store the results in "dst".
func MmFmaddsubPs(v0 M128, v1 M128, v2 M128) M128 {
	return C._mm_fmaddsub_ps(v0, v1, v2)
}

// Multiply packed double-precision (64-bit) floating-point elements in "a" and "b", alternatively add and subtract packed elements in "c" to/from the intermediate result, and store the results in "dst".
func MmFmaddsubPd(v0 M128D, v1 M128D, v2 M128D) M128D {
	return C._mm_fmaddsub_pd(v0, v1, v2)
}

// Multiply packed single-precision (32-bit) floating-point elements in "a" and "b", alternatively subtract and add packed elements in "c" from/to the intermediate result, and store the results in "dst".
func MmFmsubaddPs(v0 M128, v1 M128, v2 M128) M128 {
	return C._mm_fmsubadd_ps(v0, v1, v2)
}

// Multiply packed double-precision (64-bit) floating-point elements in "a" and "b", alternatively subtract and add packed elements in "c" from/to the intermediate result, and store the results in "dst".
func MmFmsubaddPd(v0 M128D, v1 M128D, v2 M128D) M128D {
	return C._mm_fmsubadd_pd(v0, v1, v2)
}

// Multiply packed single-precision (32-bit) floating-point elements in "a" and "b", add the intermediate result to packed elements in "c", and store the results in "dst".
func Mm256FmaddPs(v0 M256, v1 M256, v2 M256) M256 {
	return C._mm256_fmadd_ps(v0, v1, v2)
}

// Multiply packed double-precision (64-bit) floating-point elements in "a" and "b", add the intermediate result to packed elements in "c", and store the results in "dst".
func Mm256FmaddPd(v0 M256D, v1 M256D, v2 M256D) M256D {
	return C._mm256_fmadd_pd(v0, v1, v2)
}

// Multiply packed single-precision (32-bit) floating-point elements in "a" and "b", subtract packed elements in "c" from the intermediate result, and store the results in "dst".
func Mm256FmsubPs(v0 M256, v1 M256, v2 M256) M256 {
	return C._mm256_fmsub_ps(v0, v1, v2)
}

// Multiply packed double-precision (64-bit) floating-point elements in "a" and "b", subtract packed elements in "c" from the intermediate result, and store the results in "dst".
func Mm256FmsubPd(v0 M256D, v1 M256D, v2 M256D) M256D {
	return C._mm256_fmsub_pd(v0, v1, v2)
}

// Multiply packed single-precision (32-bit) floating-point elements in "a" and "b", add the negated intermediate result to packed elements in "c", and store the results in "dst".
func Mm256FnmaddPs(v0 M256, v1 M256, v2 M256) M256 {
	return C._mm256_fnmadd_ps(v0, v1, v2)
}

// Multiply packed double-precision (64-bit) floating-point elements in "a" and "b", add the negated intermediate result to packed elements in "c", and store the results in "dst".
func Mm256FnmaddPd(v0 M256D, v1 M256D, v2 M256D) M256D {
	return C._mm256_fnmadd_pd(v0, v1, v2)
}

// Multiply packed single-precision (32-bit) floating-point elements in "a" and "b", subtract packed elements in "c" from the negated intermediate result, and store the results in "dst".
func Mm256FnmsubPs(v0 M256, v1 M256, v2 M256) M256 {
	return C._mm256_fnmsub_ps(v0, v1, v2)
}

// Multiply packed double-precision (64-bit) floating-point elements in "a" and "b", subtract packed elements in "c" from the negated intermediate result, and store the results in "dst".
func Mm256FnmsubPd(v0 M256D, v1 M256D, v2 M256D) M256D {
	return C._mm256_fnmsub_pd(v0, v1, v2)
}

// Multiply packed single-precision (32-bit) floating-point elements in "a" and "b", alternatively add and subtract packed elements in "c" to/from the intermediate result, and store the results in "dst".
func Mm256FmaddsubPs(v0 M256, v1 M256, v2 M256) M256 {
	return C._mm256_fmaddsub_ps(v0, v1, v2)
}

// Multiply packed double-precision (64-bit) floating-point elements in "a" and "b", alternatively add and subtract packed elements in "c" to/from the intermediate result, and store the results in "dst".
func Mm256FmaddsubPd(v0 M256D, v1 M256D, v2 M256D) M256D {
	return C._mm256_fmaddsub_pd(v0, v1, v2)
}

// Multiply packed single-precision (32-bit) floating-point elements in "a" and "b", alternatively subtract and add packed elements in "c" from/to the intermediate result, and store the results in "dst".
func Mm256FmsubaddPs(v0 M256, v1 M256, v2 M256) M256 {
	return C._mm256_fmsubadd_ps(v0, v1, v2)
}

// Multiply packed double-precision (64-bit) floating-point elements in "a" and "b", alternatively subtract and add packed elements in "c" from/to the intermediate result, and store the results in "dst".
func Mm256FmsubaddPd(v0 M256D, v1 M256D, v2 M256D) M256D {
	return C._mm256_fmsubadd_pd(v0, v1, v2)
}
