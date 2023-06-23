package x86

/*
#cgo CFLAGS: 
#include <immintrin.h>
*/
import "C"

// _mm_cvtps_pi32
// __always_inline__, __nodebug__, __target__("mmx,sse"), __min_vector_width__(64)
func MmCvtpsPi32(v0 M128) M64 {
	return C._mm_cvtps_pi32(v0)
}

// _mm_cvt_ps2pi
// __always_inline__, __nodebug__, __target__("mmx,sse"), __min_vector_width__(64)
func MmCvtPs2Pi(v0 M128) M64 {
	return C._mm_cvt_ps2pi(v0)
}

// _mm_cvttps_pi32
// __always_inline__, __nodebug__, __target__("mmx,sse"), __min_vector_width__(64)
func MmCvttpsPi32(v0 M128) M64 {
	return C._mm_cvttps_pi32(v0)
}

// _mm_cvtt_ps2pi
// __always_inline__, __nodebug__, __target__("mmx,sse"), __min_vector_width__(64)
func MmCvttPs2Pi(v0 M128) M64 {
	return C._mm_cvtt_ps2pi(v0)
}

// _mm_cvtpi32_ps
// __always_inline__, __nodebug__, __target__("mmx,sse"), __min_vector_width__(64)
func MmCvtpi32Ps(v0 M128, v1 M64) M128 {
	return C._mm_cvtpi32_ps(v0, v1)
}

// _mm_cvt_pi2ps
// __always_inline__, __nodebug__, __target__("mmx,sse"), __min_vector_width__(64)
func MmCvtPi2Ps(v0 M128, v1 M64) M128 {
	return C._mm_cvt_pi2ps(v0, v1)
}

// _mm_max_pi16
// __always_inline__, __nodebug__, __target__("mmx,sse"), __min_vector_width__(64)
func MmMaxPi16(v0 M64, v1 M64) M64 {
	return C._mm_max_pi16(v0, v1)
}

// _mm_max_pu8
// __always_inline__, __nodebug__, __target__("mmx,sse"), __min_vector_width__(64)
func MmMaxPu8(v0 M64, v1 M64) M64 {
	return C._mm_max_pu8(v0, v1)
}

// _mm_min_pi16
// __always_inline__, __nodebug__, __target__("mmx,sse"), __min_vector_width__(64)
func MmMinPi16(v0 M64, v1 M64) M64 {
	return C._mm_min_pi16(v0, v1)
}

// _mm_min_pu8
// __always_inline__, __nodebug__, __target__("mmx,sse"), __min_vector_width__(64)
func MmMinPu8(v0 M64, v1 M64) M64 {
	return C._mm_min_pu8(v0, v1)
}

// _mm_movemask_pi8
// __always_inline__, __nodebug__, __target__("mmx,sse"), __min_vector_width__(64)
func MmMovemaskPi8(v0 M64) Int {
	return C._mm_movemask_pi8(v0)
}

// _mm_mulhi_pu16
// __always_inline__, __nodebug__, __target__("mmx,sse"), __min_vector_width__(64)
func MmMulhiPu16(v0 M64, v1 M64) M64 {
	return C._mm_mulhi_pu16(v0, v1)
}

// _mm_avg_pu8
// __always_inline__, __nodebug__, __target__("mmx,sse"), __min_vector_width__(64)
func MmAvgPu8(v0 M64, v1 M64) M64 {
	return C._mm_avg_pu8(v0, v1)
}

// _mm_avg_pu16
// __always_inline__, __nodebug__, __target__("mmx,sse"), __min_vector_width__(64)
func MmAvgPu16(v0 M64, v1 M64) M64 {
	return C._mm_avg_pu16(v0, v1)
}

// _mm_sad_pu8
// __always_inline__, __nodebug__, __target__("mmx,sse"), __min_vector_width__(64)
func MmSadPu8(v0 M64, v1 M64) M64 {
	return C._mm_sad_pu8(v0, v1)
}

// _mm_cvtpi16_ps
// __always_inline__, __nodebug__, __target__("mmx,sse"), __min_vector_width__(64)
func MmCvtpi16Ps(v0 M64) M128 {
	return C._mm_cvtpi16_ps(v0)
}

// _mm_cvtpu16_ps
// __always_inline__, __nodebug__, __target__("mmx,sse"), __min_vector_width__(64)
func MmCvtpu16Ps(v0 M64) M128 {
	return C._mm_cvtpu16_ps(v0)
}

// _mm_cvtpi8_ps
// __always_inline__, __nodebug__, __target__("mmx,sse"), __min_vector_width__(64)
func MmCvtpi8Ps(v0 M64) M128 {
	return C._mm_cvtpi8_ps(v0)
}

// _mm_cvtpu8_ps
// __always_inline__, __nodebug__, __target__("mmx,sse"), __min_vector_width__(64)
func MmCvtpu8Ps(v0 M64) M128 {
	return C._mm_cvtpu8_ps(v0)
}

// _mm_cvtpi32x2_ps
// __always_inline__, __nodebug__, __target__("mmx,sse"), __min_vector_width__(64)
func MmCvtpi32X2Ps(v0 M64, v1 M64) M128 {
	return C._mm_cvtpi32x2_ps(v0, v1)
}

// _mm_cvtps_pi16
// __always_inline__, __nodebug__, __target__("mmx,sse"), __min_vector_width__(64)
func MmCvtpsPi16(v0 M128) M64 {
	return C._mm_cvtps_pi16(v0)
}

// _mm_cvtps_pi8
// __always_inline__, __nodebug__, __target__("mmx,sse"), __min_vector_width__(64)
func MmCvtpsPi8(v0 M128) M64 {
	return C._mm_cvtps_pi8(v0)
}
