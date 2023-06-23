package x86

/*
#cgo CFLAGS: -march=native
#include <immintrin.h>
*/
import "C"

// _mm_addsub_ps
// __always_inline__, __nodebug__, __target__("sse3"), __min_vector_width__(128)
func MmAddsubPs(v0 M128, v1 M128) M128 {
	return C._mm_addsub_ps(v0, v1)
}

// _mm_hadd_ps
// __always_inline__, __nodebug__, __target__("sse3"), __min_vector_width__(128)
func MmHaddPs(v0 M128, v1 M128) M128 {
	return C._mm_hadd_ps(v0, v1)
}

// _mm_hsub_ps
// __always_inline__, __nodebug__, __target__("sse3"), __min_vector_width__(128)
func MmHsubPs(v0 M128, v1 M128) M128 {
	return C._mm_hsub_ps(v0, v1)
}

// _mm_movehdup_ps
// __always_inline__, __nodebug__, __target__("sse3"), __min_vector_width__(128)
func MmMovehdupPs(v0 M128) M128 {
	return C._mm_movehdup_ps(v0)
}

// _mm_moveldup_ps
// __always_inline__, __nodebug__, __target__("sse3"), __min_vector_width__(128)
func MmMoveldupPs(v0 M128) M128 {
	return C._mm_moveldup_ps(v0)
}

// _mm_addsub_pd
// __always_inline__, __nodebug__, __target__("sse3"), __min_vector_width__(128)
func MmAddsubPd(v0 M128D, v1 M128D) M128D {
	return C._mm_addsub_pd(v0, v1)
}

// _mm_hadd_pd
// __always_inline__, __nodebug__, __target__("sse3"), __min_vector_width__(128)
func MmHaddPd(v0 M128D, v1 M128D) M128D {
	return C._mm_hadd_pd(v0, v1)
}

// _mm_hsub_pd
// __always_inline__, __nodebug__, __target__("sse3"), __min_vector_width__(128)
func MmHsubPd(v0 M128D, v1 M128D) M128D {
	return C._mm_hsub_pd(v0, v1)
}

// _mm_movedup_pd
// __always_inline__, __nodebug__, __target__("sse3"), __min_vector_width__(128)
func MmMovedupPd(v0 M128D) M128D {
	return C._mm_movedup_pd(v0)
}

// _mm_mwait
// __always_inline__, __nodebug__, __target__("sse3"), __min_vector_width__(128)
func MmMwait(v0 Unsigned, v1 Unsigned) {
	C._mm_mwait(v0, v1)
}
