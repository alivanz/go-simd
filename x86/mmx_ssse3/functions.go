package mmx_ssse3

import (
	"github.com/alivanz/go-simd/x86"
)

/*
#cgo CFLAGS: -mmmx -mssse3
#include <immintrin.h>
*/
import "C"

// Compute the absolute value of packed signed 8-bit integers in "a", and store the unsigned results in "dst".
//
//go:linkname MmAbsPi8 MmAbsPi8
//go:noescape
func MmAbsPi8(r *x86.M64, v0 *x86.M64)

// Compute the absolute value of packed signed 16-bit integers in "a", and store the unsigned results in "dst".
//
//go:linkname MmAbsPi16 MmAbsPi16
//go:noescape
func MmAbsPi16(r *x86.M64, v0 *x86.M64)

// Compute the absolute value of packed signed 32-bit integers in "a", and store the unsigned results in "dst".
//
//go:linkname MmAbsPi32 MmAbsPi32
//go:noescape
func MmAbsPi32(r *x86.M64, v0 *x86.M64)

// Horizontally add adjacent pairs of 16-bit integers in "a" and "b", and pack the signed 16-bit results in "dst".
//
//go:linkname MmHaddPi16 MmHaddPi16
//go:noescape
func MmHaddPi16(r *x86.M64, v0 *x86.M64, v1 *x86.M64)

// Horizontally add adjacent pairs of 32-bit integers in "a" and "b", and pack the signed 32-bit results in "dst".
//
//go:linkname MmHaddPi32 MmHaddPi32
//go:noescape
func MmHaddPi32(r *x86.M64, v0 *x86.M64, v1 *x86.M64)

// Horizontally add adjacent pairs of signed 16-bit integers in "a" and "b" using saturation, and pack the signed 16-bit results in "dst".
//
//go:linkname MmHaddsPi16 MmHaddsPi16
//go:noescape
func MmHaddsPi16(r *x86.M64, v0 *x86.M64, v1 *x86.M64)

// Horizontally subtract adjacent pairs of 16-bit integers in "a" and "b", and pack the signed 16-bit results in "dst".
//
//go:linkname MmHsubPi16 MmHsubPi16
//go:noescape
func MmHsubPi16(r *x86.M64, v0 *x86.M64, v1 *x86.M64)

// Horizontally subtract adjacent pairs of 32-bit integers in "a" and "b", and pack the signed 32-bit results in "dst".
//
//go:linkname MmHsubPi32 MmHsubPi32
//go:noescape
func MmHsubPi32(r *x86.M64, v0 *x86.M64, v1 *x86.M64)

// Horizontally subtract adjacent pairs of signed 16-bit integers in "a" and "b" using saturation, and pack the signed 16-bit results in "dst".
//
//go:linkname MmHsubsPi16 MmHsubsPi16
//go:noescape
func MmHsubsPi16(r *x86.M64, v0 *x86.M64, v1 *x86.M64)

// Vertically multiply each unsigned 8-bit integer from "a" with the corresponding signed 8-bit integer from "b", producing intermediate signed 16-bit integers. Horizontally add adjacent pairs of intermediate signed 16-bit integers, and pack the saturated results in "dst".
//
//go:linkname MmMaddubsPi16 MmMaddubsPi16
//go:noescape
func MmMaddubsPi16(r *x86.M64, v0 *x86.M64, v1 *x86.M64)

// Multiply packed signed 16-bit integers in "a" and "b", producing intermediate signed 32-bit integers. Truncate each intermediate integer to the 18 most significant bits, round by adding 1, and store bits [16:1] to "dst".
//
//go:linkname MmMulhrsPi16 MmMulhrsPi16
//go:noescape
func MmMulhrsPi16(r *x86.M64, v0 *x86.M64, v1 *x86.M64)

// Shuffle packed 8-bit integers in "a" according to shuffle control mask in the corresponding 8-bit element of "b", and store the results in "dst".
//
//go:linkname MmShufflePi8 MmShufflePi8
//go:noescape
func MmShufflePi8(r *x86.M64, v0 *x86.M64, v1 *x86.M64)

// Negate packed 8-bit integers in "a" when the corresponding signed 8-bit integer in "b" is negative, and store the results in "dst". Element in "dst" are zeroed out when the corresponding element in "b" is zero.
//
//go:linkname MmSignPi8 MmSignPi8
//go:noescape
func MmSignPi8(r *x86.M64, v0 *x86.M64, v1 *x86.M64)

// Negate packed 16-bit integers in "a" when the corresponding signed 16-bit integer in "b" is negative, and store the results in "dst". Element in "dst" are zeroed out when the corresponding element in "b" is zero.
//
//go:linkname MmSignPi16 MmSignPi16
//go:noescape
func MmSignPi16(r *x86.M64, v0 *x86.M64, v1 *x86.M64)

// Negate packed 32-bit integers in "a" when the corresponding signed 32-bit integer in "b" is negative, and store the results in "dst". Element in "dst" are zeroed out when the corresponding element in "b" is zero.
//
//go:linkname MmSignPi32 MmSignPi32
//go:noescape
func MmSignPi32(r *x86.M64, v0 *x86.M64, v1 *x86.M64)
