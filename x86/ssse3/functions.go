package ssse3

import (
	"github.com/alivanz/go-simd/x86"
)

/*
#cgo CFLAGS: -mssse3
#include <immintrin.h>
*/
import "C"

// Compute the absolute value of packed signed 8-bit integers in "a", and store the unsigned results in "dst".
//
//go:linkname MmAbsEpi8 MmAbsEpi8
//go:noescape
func MmAbsEpi8(r *x86.M128I, v0 *x86.M128I)

// Compute the absolute value of packed signed 16-bit integers in "a", and store the unsigned results in "dst".
//
//go:linkname MmAbsEpi16 MmAbsEpi16
//go:noescape
func MmAbsEpi16(r *x86.M128I, v0 *x86.M128I)

// Compute the absolute value of packed signed 32-bit integers in "a", and store the unsigned results in "dst".
//
//go:linkname MmAbsEpi32 MmAbsEpi32
//go:noescape
func MmAbsEpi32(r *x86.M128I, v0 *x86.M128I)

// Horizontally add adjacent pairs of 16-bit integers in "a" and "b", and pack the signed 16-bit results in "dst".
//
//go:linkname MmHaddEpi16 MmHaddEpi16
//go:noescape
func MmHaddEpi16(r *x86.M128I, v0 *x86.M128I, v1 *x86.M128I)

// Horizontally add adjacent pairs of 32-bit integers in "a" and "b", and pack the signed 32-bit results in "dst".
//
//go:linkname MmHaddEpi32 MmHaddEpi32
//go:noescape
func MmHaddEpi32(r *x86.M128I, v0 *x86.M128I, v1 *x86.M128I)

// Horizontally add adjacent pairs of signed 16-bit integers in "a" and "b" using saturation, and pack the signed 16-bit results in "dst".
//
//go:linkname MmHaddsEpi16 MmHaddsEpi16
//go:noescape
func MmHaddsEpi16(r *x86.M128I, v0 *x86.M128I, v1 *x86.M128I)

// Horizontally subtract adjacent pairs of 16-bit integers in "a" and "b", and pack the signed 16-bit results in "dst".
//
//go:linkname MmHsubEpi16 MmHsubEpi16
//go:noescape
func MmHsubEpi16(r *x86.M128I, v0 *x86.M128I, v1 *x86.M128I)

// Horizontally subtract adjacent pairs of 32-bit integers in "a" and "b", and pack the signed 32-bit results in "dst".
//
//go:linkname MmHsubEpi32 MmHsubEpi32
//go:noescape
func MmHsubEpi32(r *x86.M128I, v0 *x86.M128I, v1 *x86.M128I)

// Horizontally subtract adjacent pairs of signed 16-bit integers in "a" and "b" using saturation, and pack the signed 16-bit results in "dst".
//
//go:linkname MmHsubsEpi16 MmHsubsEpi16
//go:noescape
func MmHsubsEpi16(r *x86.M128I, v0 *x86.M128I, v1 *x86.M128I)

// Vertically multiply each unsigned 8-bit integer from "a" with the corresponding signed 8-bit integer from "b", producing intermediate signed 16-bit integers. Horizontally add adjacent pairs of intermediate signed 16-bit integers, and pack the saturated results in "dst".
//
//go:linkname MmMaddubsEpi16 MmMaddubsEpi16
//go:noescape
func MmMaddubsEpi16(r *x86.M128I, v0 *x86.M128I, v1 *x86.M128I)

// Multiply packed signed 16-bit integers in "a" and "b", producing intermediate signed 32-bit integers. Truncate each intermediate integer to the 18 most significant bits, round by adding 1, and store bits [16:1] to "dst".
//
//go:linkname MmMulhrsEpi16 MmMulhrsEpi16
//go:noescape
func MmMulhrsEpi16(r *x86.M128I, v0 *x86.M128I, v1 *x86.M128I)

// Shuffle packed 8-bit integers in "a" according to shuffle control mask in the corresponding 8-bit element of "b", and store the results in "dst".
//
//go:linkname MmShuffleEpi8 MmShuffleEpi8
//go:noescape
func MmShuffleEpi8(r *x86.M128I, v0 *x86.M128I, v1 *x86.M128I)

// Negate packed 8-bit integers in "a" when the corresponding signed 8-bit integer in "b" is negative, and store the results in "dst". Element in "dst" are zeroed out when the corresponding element in "b" is zero.
//
//go:linkname MmSignEpi8 MmSignEpi8
//go:noescape
func MmSignEpi8(r *x86.M128I, v0 *x86.M128I, v1 *x86.M128I)

// Negate packed 16-bit integers in "a" when the corresponding signed 16-bit integer in "b" is negative, and store the results in "dst". Element in "dst" are zeroed out when the corresponding element in "b" is zero.
//
//go:linkname MmSignEpi16 MmSignEpi16
//go:noescape
func MmSignEpi16(r *x86.M128I, v0 *x86.M128I, v1 *x86.M128I)

// Negate packed 32-bit integers in "a" when the corresponding signed 32-bit integer in "b" is negative, and store the results in "dst". Element in "dst" are zeroed out when the corresponding element in "b" is zero.
//
//go:linkname MmSignEpi32 MmSignEpi32
//go:noescape
func MmSignEpi32(r *x86.M128I, v0 *x86.M128I, v1 *x86.M128I)
