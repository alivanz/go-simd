package x86

/*
#cgo CFLAGS: -march=native
#include <immintrin.h>
*/
import "C"

// _mm_abs_pi8
// __always_inline__, __nodebug__, __target__("mmx,ssse3"), __min_vector_width__(64)
func MmAbsPi8(v0 M64) M64 {
	return C._mm_abs_pi8(v0)
}

// _mm_abs_pi16
// __always_inline__, __nodebug__, __target__("mmx,ssse3"), __min_vector_width__(64)
func MmAbsPi16(v0 M64) M64 {
	return C._mm_abs_pi16(v0)
}

// _mm_abs_pi32
// __always_inline__, __nodebug__, __target__("mmx,ssse3"), __min_vector_width__(64)
func MmAbsPi32(v0 M64) M64 {
	return C._mm_abs_pi32(v0)
}

// _mm_hadd_pi16
// __always_inline__, __nodebug__, __target__("mmx,ssse3"), __min_vector_width__(64)
func MmHaddPi16(v0 M64, v1 M64) M64 {
	return C._mm_hadd_pi16(v0, v1)
}

// _mm_hadd_pi32
// __always_inline__, __nodebug__, __target__("mmx,ssse3"), __min_vector_width__(64)
func MmHaddPi32(v0 M64, v1 M64) M64 {
	return C._mm_hadd_pi32(v0, v1)
}

// _mm_hadds_pi16
// __always_inline__, __nodebug__, __target__("mmx,ssse3"), __min_vector_width__(64)
func MmHaddsPi16(v0 M64, v1 M64) M64 {
	return C._mm_hadds_pi16(v0, v1)
}

// _mm_hsub_pi16
// __always_inline__, __nodebug__, __target__("mmx,ssse3"), __min_vector_width__(64)
func MmHsubPi16(v0 M64, v1 M64) M64 {
	return C._mm_hsub_pi16(v0, v1)
}

// _mm_hsub_pi32
// __always_inline__, __nodebug__, __target__("mmx,ssse3"), __min_vector_width__(64)
func MmHsubPi32(v0 M64, v1 M64) M64 {
	return C._mm_hsub_pi32(v0, v1)
}

// _mm_hsubs_pi16
// __always_inline__, __nodebug__, __target__("mmx,ssse3"), __min_vector_width__(64)
func MmHsubsPi16(v0 M64, v1 M64) M64 {
	return C._mm_hsubs_pi16(v0, v1)
}

// _mm_maddubs_pi16
// __always_inline__, __nodebug__, __target__("mmx,ssse3"), __min_vector_width__(64)
func MmMaddubsPi16(v0 M64, v1 M64) M64 {
	return C._mm_maddubs_pi16(v0, v1)
}

// _mm_mulhrs_pi16
// __always_inline__, __nodebug__, __target__("mmx,ssse3"), __min_vector_width__(64)
func MmMulhrsPi16(v0 M64, v1 M64) M64 {
	return C._mm_mulhrs_pi16(v0, v1)
}

// _mm_shuffle_pi8
// __always_inline__, __nodebug__, __target__("mmx,ssse3"), __min_vector_width__(64)
func MmShufflePi8(v0 M64, v1 M64) M64 {
	return C._mm_shuffle_pi8(v0, v1)
}

// _mm_sign_pi8
// __always_inline__, __nodebug__, __target__("mmx,ssse3"), __min_vector_width__(64)
func MmSignPi8(v0 M64, v1 M64) M64 {
	return C._mm_sign_pi8(v0, v1)
}

// _mm_sign_pi16
// __always_inline__, __nodebug__, __target__("mmx,ssse3"), __min_vector_width__(64)
func MmSignPi16(v0 M64, v1 M64) M64 {
	return C._mm_sign_pi16(v0, v1)
}

// _mm_sign_pi32
// __always_inline__, __nodebug__, __target__("mmx,ssse3"), __min_vector_width__(64)
func MmSignPi32(v0 M64, v1 M64) M64 {
	return C._mm_sign_pi32(v0, v1)
}
