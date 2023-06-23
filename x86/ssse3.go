package x86

/*
#cgo CFLAGS: -march=native
#include <immintrin.h>
*/
import "C"

// _mm_abs_epi8
// __always_inline__, __nodebug__, __target__("ssse3"), __min_vector_width__(64)
func MmAbsEpi8(v0 M128I) M128I {
	return C._mm_abs_epi8(v0)
}

// _mm_abs_epi16
// __always_inline__, __nodebug__, __target__("ssse3"), __min_vector_width__(64)
func MmAbsEpi16(v0 M128I) M128I {
	return C._mm_abs_epi16(v0)
}

// _mm_abs_epi32
// __always_inline__, __nodebug__, __target__("ssse3"), __min_vector_width__(64)
func MmAbsEpi32(v0 M128I) M128I {
	return C._mm_abs_epi32(v0)
}

// _mm_hadd_epi16
// __always_inline__, __nodebug__, __target__("ssse3"), __min_vector_width__(64)
func MmHaddEpi16(v0 M128I, v1 M128I) M128I {
	return C._mm_hadd_epi16(v0, v1)
}

// _mm_hadd_epi32
// __always_inline__, __nodebug__, __target__("ssse3"), __min_vector_width__(64)
func MmHaddEpi32(v0 M128I, v1 M128I) M128I {
	return C._mm_hadd_epi32(v0, v1)
}

// _mm_hadds_epi16
// __always_inline__, __nodebug__, __target__("ssse3"), __min_vector_width__(64)
func MmHaddsEpi16(v0 M128I, v1 M128I) M128I {
	return C._mm_hadds_epi16(v0, v1)
}

// _mm_hsub_epi16
// __always_inline__, __nodebug__, __target__("ssse3"), __min_vector_width__(64)
func MmHsubEpi16(v0 M128I, v1 M128I) M128I {
	return C._mm_hsub_epi16(v0, v1)
}

// _mm_hsub_epi32
// __always_inline__, __nodebug__, __target__("ssse3"), __min_vector_width__(64)
func MmHsubEpi32(v0 M128I, v1 M128I) M128I {
	return C._mm_hsub_epi32(v0, v1)
}

// _mm_hsubs_epi16
// __always_inline__, __nodebug__, __target__("ssse3"), __min_vector_width__(64)
func MmHsubsEpi16(v0 M128I, v1 M128I) M128I {
	return C._mm_hsubs_epi16(v0, v1)
}

// _mm_maddubs_epi16
// __always_inline__, __nodebug__, __target__("ssse3"), __min_vector_width__(64)
func MmMaddubsEpi16(v0 M128I, v1 M128I) M128I {
	return C._mm_maddubs_epi16(v0, v1)
}

// _mm_mulhrs_epi16
// __always_inline__, __nodebug__, __target__("ssse3"), __min_vector_width__(64)
func MmMulhrsEpi16(v0 M128I, v1 M128I) M128I {
	return C._mm_mulhrs_epi16(v0, v1)
}

// _mm_shuffle_epi8
// __always_inline__, __nodebug__, __target__("ssse3"), __min_vector_width__(64)
func MmShuffleEpi8(v0 M128I, v1 M128I) M128I {
	return C._mm_shuffle_epi8(v0, v1)
}

// _mm_sign_epi8
// __always_inline__, __nodebug__, __target__("ssse3"), __min_vector_width__(64)
func MmSignEpi8(v0 M128I, v1 M128I) M128I {
	return C._mm_sign_epi8(v0, v1)
}

// _mm_sign_epi16
// __always_inline__, __nodebug__, __target__("ssse3"), __min_vector_width__(64)
func MmSignEpi16(v0 M128I, v1 M128I) M128I {
	return C._mm_sign_epi16(v0, v1)
}

// _mm_sign_epi32
// __always_inline__, __nodebug__, __target__("ssse3"), __min_vector_width__(64)
func MmSignEpi32(v0 M128I, v1 M128I) M128I {
	return C._mm_sign_epi32(v0, v1)
}
