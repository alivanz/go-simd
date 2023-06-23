package x86

/*
#cgo CFLAGS: -march=native
#include <immintrin.h>
*/
import "C"

// _mm_fmadd_ps
// __always_inline__, __nodebug__, __target__("fma"), __min_vector_width__(128)
func MmFmaddPs(v0 M128, v1 M128, v2 M128) M128 {
	return C._mm_fmadd_ps(v0, v1, v2)
}

// _mm_fmadd_pd
// __always_inline__, __nodebug__, __target__("fma"), __min_vector_width__(128)
func MmFmaddPd(v0 M128D, v1 M128D, v2 M128D) M128D {
	return C._mm_fmadd_pd(v0, v1, v2)
}

// _mm_fmadd_ss
// __always_inline__, __nodebug__, __target__("fma"), __min_vector_width__(128)
func MmFmaddSs(v0 M128, v1 M128, v2 M128) M128 {
	return C._mm_fmadd_ss(v0, v1, v2)
}

// _mm_fmadd_sd
// __always_inline__, __nodebug__, __target__("fma"), __min_vector_width__(128)
func MmFmaddSd(v0 M128D, v1 M128D, v2 M128D) M128D {
	return C._mm_fmadd_sd(v0, v1, v2)
}

// _mm_fmsub_ps
// __always_inline__, __nodebug__, __target__("fma"), __min_vector_width__(128)
func MmFmsubPs(v0 M128, v1 M128, v2 M128) M128 {
	return C._mm_fmsub_ps(v0, v1, v2)
}

// _mm_fmsub_pd
// __always_inline__, __nodebug__, __target__("fma"), __min_vector_width__(128)
func MmFmsubPd(v0 M128D, v1 M128D, v2 M128D) M128D {
	return C._mm_fmsub_pd(v0, v1, v2)
}

// _mm_fmsub_ss
// __always_inline__, __nodebug__, __target__("fma"), __min_vector_width__(128)
func MmFmsubSs(v0 M128, v1 M128, v2 M128) M128 {
	return C._mm_fmsub_ss(v0, v1, v2)
}

// _mm_fmsub_sd
// __always_inline__, __nodebug__, __target__("fma"), __min_vector_width__(128)
func MmFmsubSd(v0 M128D, v1 M128D, v2 M128D) M128D {
	return C._mm_fmsub_sd(v0, v1, v2)
}

// _mm_fnmadd_ps
// __always_inline__, __nodebug__, __target__("fma"), __min_vector_width__(128)
func MmFnmaddPs(v0 M128, v1 M128, v2 M128) M128 {
	return C._mm_fnmadd_ps(v0, v1, v2)
}

// _mm_fnmadd_pd
// __always_inline__, __nodebug__, __target__("fma"), __min_vector_width__(128)
func MmFnmaddPd(v0 M128D, v1 M128D, v2 M128D) M128D {
	return C._mm_fnmadd_pd(v0, v1, v2)
}

// _mm_fnmadd_ss
// __always_inline__, __nodebug__, __target__("fma"), __min_vector_width__(128)
func MmFnmaddSs(v0 M128, v1 M128, v2 M128) M128 {
	return C._mm_fnmadd_ss(v0, v1, v2)
}

// _mm_fnmadd_sd
// __always_inline__, __nodebug__, __target__("fma"), __min_vector_width__(128)
func MmFnmaddSd(v0 M128D, v1 M128D, v2 M128D) M128D {
	return C._mm_fnmadd_sd(v0, v1, v2)
}

// _mm_fnmsub_ps
// __always_inline__, __nodebug__, __target__("fma"), __min_vector_width__(128)
func MmFnmsubPs(v0 M128, v1 M128, v2 M128) M128 {
	return C._mm_fnmsub_ps(v0, v1, v2)
}

// _mm_fnmsub_pd
// __always_inline__, __nodebug__, __target__("fma"), __min_vector_width__(128)
func MmFnmsubPd(v0 M128D, v1 M128D, v2 M128D) M128D {
	return C._mm_fnmsub_pd(v0, v1, v2)
}

// _mm_fnmsub_ss
// __always_inline__, __nodebug__, __target__("fma"), __min_vector_width__(128)
func MmFnmsubSs(v0 M128, v1 M128, v2 M128) M128 {
	return C._mm_fnmsub_ss(v0, v1, v2)
}

// _mm_fnmsub_sd
// __always_inline__, __nodebug__, __target__("fma"), __min_vector_width__(128)
func MmFnmsubSd(v0 M128D, v1 M128D, v2 M128D) M128D {
	return C._mm_fnmsub_sd(v0, v1, v2)
}

// _mm_fmaddsub_ps
// __always_inline__, __nodebug__, __target__("fma"), __min_vector_width__(128)
func MmFmaddsubPs(v0 M128, v1 M128, v2 M128) M128 {
	return C._mm_fmaddsub_ps(v0, v1, v2)
}

// _mm_fmaddsub_pd
// __always_inline__, __nodebug__, __target__("fma"), __min_vector_width__(128)
func MmFmaddsubPd(v0 M128D, v1 M128D, v2 M128D) M128D {
	return C._mm_fmaddsub_pd(v0, v1, v2)
}

// _mm_fmsubadd_ps
// __always_inline__, __nodebug__, __target__("fma"), __min_vector_width__(128)
func MmFmsubaddPs(v0 M128, v1 M128, v2 M128) M128 {
	return C._mm_fmsubadd_ps(v0, v1, v2)
}

// _mm_fmsubadd_pd
// __always_inline__, __nodebug__, __target__("fma"), __min_vector_width__(128)
func MmFmsubaddPd(v0 M128D, v1 M128D, v2 M128D) M128D {
	return C._mm_fmsubadd_pd(v0, v1, v2)
}

// _mm256_fmadd_ps
// __always_inline__, __nodebug__, __target__("fma"), __min_vector_width__(256)
func Mm256FmaddPs(v0 M256, v1 M256, v2 M256) M256 {
	return C._mm256_fmadd_ps(v0, v1, v2)
}

// _mm256_fmadd_pd
// __always_inline__, __nodebug__, __target__("fma"), __min_vector_width__(256)
func Mm256FmaddPd(v0 M256D, v1 M256D, v2 M256D) M256D {
	return C._mm256_fmadd_pd(v0, v1, v2)
}

// _mm256_fmsub_ps
// __always_inline__, __nodebug__, __target__("fma"), __min_vector_width__(256)
func Mm256FmsubPs(v0 M256, v1 M256, v2 M256) M256 {
	return C._mm256_fmsub_ps(v0, v1, v2)
}

// _mm256_fmsub_pd
// __always_inline__, __nodebug__, __target__("fma"), __min_vector_width__(256)
func Mm256FmsubPd(v0 M256D, v1 M256D, v2 M256D) M256D {
	return C._mm256_fmsub_pd(v0, v1, v2)
}

// _mm256_fnmadd_ps
// __always_inline__, __nodebug__, __target__("fma"), __min_vector_width__(256)
func Mm256FnmaddPs(v0 M256, v1 M256, v2 M256) M256 {
	return C._mm256_fnmadd_ps(v0, v1, v2)
}

// _mm256_fnmadd_pd
// __always_inline__, __nodebug__, __target__("fma"), __min_vector_width__(256)
func Mm256FnmaddPd(v0 M256D, v1 M256D, v2 M256D) M256D {
	return C._mm256_fnmadd_pd(v0, v1, v2)
}

// _mm256_fnmsub_ps
// __always_inline__, __nodebug__, __target__("fma"), __min_vector_width__(256)
func Mm256FnmsubPs(v0 M256, v1 M256, v2 M256) M256 {
	return C._mm256_fnmsub_ps(v0, v1, v2)
}

// _mm256_fnmsub_pd
// __always_inline__, __nodebug__, __target__("fma"), __min_vector_width__(256)
func Mm256FnmsubPd(v0 M256D, v1 M256D, v2 M256D) M256D {
	return C._mm256_fnmsub_pd(v0, v1, v2)
}

// _mm256_fmaddsub_ps
// __always_inline__, __nodebug__, __target__("fma"), __min_vector_width__(256)
func Mm256FmaddsubPs(v0 M256, v1 M256, v2 M256) M256 {
	return C._mm256_fmaddsub_ps(v0, v1, v2)
}

// _mm256_fmaddsub_pd
// __always_inline__, __nodebug__, __target__("fma"), __min_vector_width__(256)
func Mm256FmaddsubPd(v0 M256D, v1 M256D, v2 M256D) M256D {
	return C._mm256_fmaddsub_pd(v0, v1, v2)
}

// _mm256_fmsubadd_ps
// __always_inline__, __nodebug__, __target__("fma"), __min_vector_width__(256)
func Mm256FmsubaddPs(v0 M256, v1 M256, v2 M256) M256 {
	return C._mm256_fmsubadd_ps(v0, v1, v2)
}

// _mm256_fmsubadd_pd
// __always_inline__, __nodebug__, __target__("fma"), __min_vector_width__(256)
func Mm256FmsubaddPd(v0 M256D, v1 M256D, v2 M256D) M256D {
	return C._mm256_fmsubadd_pd(v0, v1, v2)
}
