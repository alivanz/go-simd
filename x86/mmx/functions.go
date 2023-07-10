package mmx

import (
	"github.com/alivanz/go-simd/x86"
)

/*
#cgo CFLAGS: -mmmx
#include <immintrin.h>
*/
import "C"

// Empty the MMX state, which marks the x87 FPU registers as available for use by x87 instructions. This instruction must be used at the end of all MMX technology procedures.
//
//go:linkname MmEmpty MmEmpty
//go:noescape
func MmEmpty()

// Copy 32-bit integer "a" to the lower elements of "dst", and zero the upper element of "dst".
//
//go:linkname MmCvtsi32Si64 MmCvtsi32Si64
//go:noescape
func MmCvtsi32Si64(r *x86.M64, v0 *x86.Int)

// Copy the lower 32-bit integer in "a" to "dst".
//
//go:linkname MmCvtsi64Si32 MmCvtsi64Si32
//go:noescape
func MmCvtsi64Si32(r *x86.Int, v0 *x86.M64)

// Copy 64-bit integer "a" to "dst".
//
//go:linkname MmCvtsi64M64 MmCvtsi64M64
//go:noescape
func MmCvtsi64M64(r *x86.M64, v0 *x86.Longlong)

// Copy 64-bit integer "a" to "dst".
//
//go:linkname MmCvtm64Si64 MmCvtm64Si64
//go:noescape
func MmCvtm64Si64(r *x86.Longlong, v0 *x86.M64)

// Convert packed signed 16-bit integers from "a" and "b" to packed 8-bit integers using signed saturation, and store the results in "dst".
//
//go:linkname MmPacksPi16 MmPacksPi16
//go:noescape
func MmPacksPi16(r *x86.M64, v0 *x86.M64, v1 *x86.M64)

// Convert packed signed 32-bit integers from "a" and "b" to packed 16-bit integers using signed saturation, and store the results in "dst".
//
//go:linkname MmPacksPi32 MmPacksPi32
//go:noescape
func MmPacksPi32(r *x86.M64, v0 *x86.M64, v1 *x86.M64)

// Convert packed signed 16-bit integers from "a" and "b" to packed 8-bit integers using unsigned saturation, and store the results in "dst".
//
//go:linkname MmPacksPu16 MmPacksPu16
//go:noescape
func MmPacksPu16(r *x86.M64, v0 *x86.M64, v1 *x86.M64)

// Unpack and interleave 8-bit integers from the high half of "a" and "b", and store the results in "dst".
//
//go:linkname MmUnpackhiPi8 MmUnpackhiPi8
//go:noescape
func MmUnpackhiPi8(r *x86.M64, v0 *x86.M64, v1 *x86.M64)

// Unpack and interleave 16-bit integers from the high half of "a" and "b", and store the results in "dst".
//
//go:linkname MmUnpackhiPi16 MmUnpackhiPi16
//go:noescape
func MmUnpackhiPi16(r *x86.M64, v0 *x86.M64, v1 *x86.M64)

// Unpack and interleave 32-bit integers from the high half of "a" and "b", and store the results in "dst".
//
//go:linkname MmUnpackhiPi32 MmUnpackhiPi32
//go:noescape
func MmUnpackhiPi32(r *x86.M64, v0 *x86.M64, v1 *x86.M64)

// Unpack and interleave 8-bit integers from the low half of "a" and "b", and store the results in "dst".
//
//go:linkname MmUnpackloPi8 MmUnpackloPi8
//go:noescape
func MmUnpackloPi8(r *x86.M64, v0 *x86.M64, v1 *x86.M64)

// Unpack and interleave 16-bit integers from the low half of "a" and "b", and store the results in "dst".
//
//go:linkname MmUnpackloPi16 MmUnpackloPi16
//go:noescape
func MmUnpackloPi16(r *x86.M64, v0 *x86.M64, v1 *x86.M64)

// Unpack and interleave 32-bit integers from the low half of "a" and "b", and store the results in "dst".
//
//go:linkname MmUnpackloPi32 MmUnpackloPi32
//go:noescape
func MmUnpackloPi32(r *x86.M64, v0 *x86.M64, v1 *x86.M64)

// Add packed 8-bit integers in "a" and "b", and store the results in "dst".
//
//go:linkname MmAddPi8 MmAddPi8
//go:noescape
func MmAddPi8(r *x86.M64, v0 *x86.M64, v1 *x86.M64)

// Add packed 16-bit integers in "a" and "b", and store the results in "dst".
//
//go:linkname MmAddPi16 MmAddPi16
//go:noescape
func MmAddPi16(r *x86.M64, v0 *x86.M64, v1 *x86.M64)

// Add packed 32-bit integers in "a" and "b", and store the results in "dst".
//
//go:linkname MmAddPi32 MmAddPi32
//go:noescape
func MmAddPi32(r *x86.M64, v0 *x86.M64, v1 *x86.M64)

// Add packed signed 8-bit integers in "a" and "b" using saturation, and store the results in "dst".
//
//go:linkname MmAddsPi8 MmAddsPi8
//go:noescape
func MmAddsPi8(r *x86.M64, v0 *x86.M64, v1 *x86.M64)

// Add packed signed 16-bit integers in "a" and "b" using saturation, and store the results in "dst".
//
//go:linkname MmAddsPi16 MmAddsPi16
//go:noescape
func MmAddsPi16(r *x86.M64, v0 *x86.M64, v1 *x86.M64)

// Add packed unsigned 8-bit integers in "a" and "b" using saturation, and store the results in "dst".
//
//go:linkname MmAddsPu8 MmAddsPu8
//go:noescape
func MmAddsPu8(r *x86.M64, v0 *x86.M64, v1 *x86.M64)

// Add packed unsigned 16-bit integers in "a" and "b" using saturation, and store the results in "dst".
//
//go:linkname MmAddsPu16 MmAddsPu16
//go:noescape
func MmAddsPu16(r *x86.M64, v0 *x86.M64, v1 *x86.M64)

// Subtract packed 8-bit integers in "b" from packed 8-bit integers in "a", and store the results in "dst".
//
//go:linkname MmSubPi8 MmSubPi8
//go:noescape
func MmSubPi8(r *x86.M64, v0 *x86.M64, v1 *x86.M64)

// Subtract packed 16-bit integers in "b" from packed 16-bit integers in "a", and store the results in "dst".
//
//go:linkname MmSubPi16 MmSubPi16
//go:noescape
func MmSubPi16(r *x86.M64, v0 *x86.M64, v1 *x86.M64)

// Subtract packed 32-bit integers in "b" from packed 32-bit integers in "a", and store the results in "dst".
//
//go:linkname MmSubPi32 MmSubPi32
//go:noescape
func MmSubPi32(r *x86.M64, v0 *x86.M64, v1 *x86.M64)

// Subtract packed signed 8-bit integers in "b" from packed 8-bit integers in "a" using saturation, and store the results in "dst".
//
//go:linkname MmSubsPi8 MmSubsPi8
//go:noescape
func MmSubsPi8(r *x86.M64, v0 *x86.M64, v1 *x86.M64)

// Subtract packed signed 16-bit integers in "b" from packed 16-bit integers in "a" using saturation, and store the results in "dst".
//
//go:linkname MmSubsPi16 MmSubsPi16
//go:noescape
func MmSubsPi16(r *x86.M64, v0 *x86.M64, v1 *x86.M64)

// Subtract packed unsigned 8-bit integers in "b" from packed unsigned 8-bit integers in "a" using saturation, and store the results in "dst".
//
//go:linkname MmSubsPu8 MmSubsPu8
//go:noescape
func MmSubsPu8(r *x86.M64, v0 *x86.M64, v1 *x86.M64)

// Subtract packed unsigned 16-bit integers in "b" from packed unsigned 16-bit integers in "a" using saturation, and store the results in "dst".
//
//go:linkname MmSubsPu16 MmSubsPu16
//go:noescape
func MmSubsPu16(r *x86.M64, v0 *x86.M64, v1 *x86.M64)

// Multiply packed signed 16-bit integers in "a" and "b", producing intermediate signed 32-bit integers. Horizontally add adjacent pairs of intermediate 32-bit integers, and pack the results in "dst".
//
//go:linkname MmMaddPi16 MmMaddPi16
//go:noescape
func MmMaddPi16(r *x86.M64, v0 *x86.M64, v1 *x86.M64)

// Multiply the packed signed 16-bit integers in "a" and "b", producing intermediate 32-bit integers, and store the high 16 bits of the intermediate integers in "dst".
//
//go:linkname MmMulhiPi16 MmMulhiPi16
//go:noescape
func MmMulhiPi16(r *x86.M64, v0 *x86.M64, v1 *x86.M64)

// Multiply the packed 16-bit integers in "a" and "b", producing intermediate 32-bit integers, and store the low 16 bits of the intermediate integers in "dst".
//
//go:linkname MmMulloPi16 MmMulloPi16
//go:noescape
func MmMulloPi16(r *x86.M64, v0 *x86.M64, v1 *x86.M64)

// Shift packed 16-bit integers in "a" left by "count" while shifting in zeros, and store the results in "dst".
//
//go:linkname MmSllPi16 MmSllPi16
//go:noescape
func MmSllPi16(r *x86.M64, v0 *x86.M64, v1 *x86.M64)

// Shift packed 16-bit integers in "a" left by "imm8" while shifting in zeros, and store the results in "dst".
//
//go:linkname MmSlliPi16 MmSlliPi16
//go:noescape
func MmSlliPi16(r *x86.M64, v0 *x86.M64, v1 *x86.Int)

// Shift packed 32-bit integers in "a" left by "count" while shifting in zeros, and store the results in "dst".
//
//go:linkname MmSllPi32 MmSllPi32
//go:noescape
func MmSllPi32(r *x86.M64, v0 *x86.M64, v1 *x86.M64)

// Shift packed 32-bit integers in "a" left by "imm8" while shifting in zeros, and store the results in "dst".
//
//go:linkname MmSlliPi32 MmSlliPi32
//go:noescape
func MmSlliPi32(r *x86.M64, v0 *x86.M64, v1 *x86.Int)

// Shift 64-bit integer "a" left by "count" while shifting in zeros, and store the result in "dst".
//
//go:linkname MmSllSi64 MmSllSi64
//go:noescape
func MmSllSi64(r *x86.M64, v0 *x86.M64, v1 *x86.M64)

// Shift 64-bit integer "a" left by "imm8" while shifting in zeros, and store the result in "dst".
//
//go:linkname MmSlliSi64 MmSlliSi64
//go:noescape
func MmSlliSi64(r *x86.M64, v0 *x86.M64, v1 *x86.Int)

// Shift packed 16-bit integers in "a" right by "count" while shifting in sign bits, and store the results in "dst".
//
//go:linkname MmSraPi16 MmSraPi16
//go:noescape
func MmSraPi16(r *x86.M64, v0 *x86.M64, v1 *x86.M64)

// Shift packed 16-bit integers in "a" right by "imm8" while shifting in sign bits, and store the results in "dst".
//
//go:linkname MmSraiPi16 MmSraiPi16
//go:noescape
func MmSraiPi16(r *x86.M64, v0 *x86.M64, v1 *x86.Int)

// Shift packed 32-bit integers in "a" right by "count" while shifting in sign bits, and store the results in "dst".
//
//go:linkname MmSraPi32 MmSraPi32
//go:noescape
func MmSraPi32(r *x86.M64, v0 *x86.M64, v1 *x86.M64)

// Shift packed 32-bit integers in "a" right by "imm8" while shifting in sign bits, and store the results in "dst".
//
//go:linkname MmSraiPi32 MmSraiPi32
//go:noescape
func MmSraiPi32(r *x86.M64, v0 *x86.M64, v1 *x86.Int)

// Shift packed 16-bit integers in "a" right by "count" while shifting in zeros, and store the results in "dst".
//
//go:linkname MmSrlPi16 MmSrlPi16
//go:noescape
func MmSrlPi16(r *x86.M64, v0 *x86.M64, v1 *x86.M64)

// Shift packed 16-bit integers in "a" right by "imm8" while shifting in zeros, and store the results in "dst".
//
//go:linkname MmSrliPi16 MmSrliPi16
//go:noescape
func MmSrliPi16(r *x86.M64, v0 *x86.M64, v1 *x86.Int)

// Shift packed 32-bit integers in "a" right by "count" while shifting in zeros, and store the results in "dst".
//
//go:linkname MmSrlPi32 MmSrlPi32
//go:noescape
func MmSrlPi32(r *x86.M64, v0 *x86.M64, v1 *x86.M64)

// Shift packed 32-bit integers in "a" right by "imm8" while shifting in zeros, and store the results in "dst".
//
//go:linkname MmSrliPi32 MmSrliPi32
//go:noescape
func MmSrliPi32(r *x86.M64, v0 *x86.M64, v1 *x86.Int)

// Shift 64-bit integer "a" right by "count" while shifting in zeros, and store the result in "dst".
//
//go:linkname MmSrlSi64 MmSrlSi64
//go:noescape
func MmSrlSi64(r *x86.M64, v0 *x86.M64, v1 *x86.M64)

// Shift 64-bit integer "a" right by "imm8" while shifting in zeros, and store the result in "dst".
//
//go:linkname MmSrliSi64 MmSrliSi64
//go:noescape
func MmSrliSi64(r *x86.M64, v0 *x86.M64, v1 *x86.Int)

// Compute the bitwise AND of 64 bits (representing integer data) in "a" and "b", and store the result in "dst".
//
//go:linkname MmAndSi64 MmAndSi64
//go:noescape
func MmAndSi64(r *x86.M64, v0 *x86.M64, v1 *x86.M64)

// Compute the bitwise NOT of 64 bits (representing integer data) in "a" and then AND with "b", and store the result in "dst".
//
//go:linkname MmAndnotSi64 MmAndnotSi64
//go:noescape
func MmAndnotSi64(r *x86.M64, v0 *x86.M64, v1 *x86.M64)

// Compute the bitwise OR of 64 bits (representing integer data) in "a" and "b", and store the result in "dst".
//
//go:linkname MmOrSi64 MmOrSi64
//go:noescape
func MmOrSi64(r *x86.M64, v0 *x86.M64, v1 *x86.M64)

// Compute the bitwise XOR of 64 bits (representing integer data) in "a" and "b", and store the result in "dst".
//
//go:linkname MmXorSi64 MmXorSi64
//go:noescape
func MmXorSi64(r *x86.M64, v0 *x86.M64, v1 *x86.M64)

// Compare packed 8-bit integers in "a" and "b" for equality, and store the results in "dst".
//
//go:linkname MmCmpeqPi8 MmCmpeqPi8
//go:noescape
func MmCmpeqPi8(r *x86.M64, v0 *x86.M64, v1 *x86.M64)

// Compare packed 16-bit integers in "a" and "b" for equality, and store the results in "dst".
//
//go:linkname MmCmpeqPi16 MmCmpeqPi16
//go:noescape
func MmCmpeqPi16(r *x86.M64, v0 *x86.M64, v1 *x86.M64)

// Compare packed 32-bit integers in "a" and "b" for equality, and store the results in "dst".
//
//go:linkname MmCmpeqPi32 MmCmpeqPi32
//go:noescape
func MmCmpeqPi32(r *x86.M64, v0 *x86.M64, v1 *x86.M64)

// Compare packed signed 8-bit integers in "a" and "b" for greater-than, and store the results in "dst".
//
//go:linkname MmCmpgtPi8 MmCmpgtPi8
//go:noescape
func MmCmpgtPi8(r *x86.M64, v0 *x86.M64, v1 *x86.M64)

// Compare packed signed 16-bit integers in "a" and "b" for greater-than, and store the results in "dst".
//
//go:linkname MmCmpgtPi16 MmCmpgtPi16
//go:noescape
func MmCmpgtPi16(r *x86.M64, v0 *x86.M64, v1 *x86.M64)

// Compare packed signed 32-bit integers in "a" and "b" for greater-than, and store the results in "dst".
//
//go:linkname MmCmpgtPi32 MmCmpgtPi32
//go:noescape
func MmCmpgtPi32(r *x86.M64, v0 *x86.M64, v1 *x86.M64)

// Return vector of type __m64 with all elements set to zero.
//
//go:linkname MmSetzeroSi64 MmSetzeroSi64
//go:noescape
func MmSetzeroSi64(r *x86.M64, )

// Set packed 32-bit integers in "dst" with the supplied values.
//
//go:linkname MmSetPi32 MmSetPi32
//go:noescape
func MmSetPi32(r *x86.M64, v0 *x86.Int, v1 *x86.Int)

// Set packed 16-bit integers in "dst" with the supplied values.
//
//go:linkname MmSetPi16 MmSetPi16
//go:noescape
func MmSetPi16(r *x86.M64, v0 *x86.Short, v1 *x86.Short, v2 *x86.Short, v3 *x86.Short)

// Set packed 8-bit integers in "dst" with the supplied values.
//
//go:linkname MmSetPi8 MmSetPi8
//go:noescape
func MmSetPi8(r *x86.M64, v0 *x86.Char, v1 *x86.Char, v2 *x86.Char, v3 *x86.Char, v4 *x86.Char, v5 *x86.Char, v6 *x86.Char, v7 *x86.Char)

// Broadcast 32-bit integer "a" to all elements of "dst".
//
//go:linkname MmSet1Pi32 MmSet1Pi32
//go:noescape
func MmSet1Pi32(r *x86.M64, v0 *x86.Int)

// Broadcast 16-bit integer "a" to all all elements of "dst".
//
//go:linkname MmSet1Pi16 MmSet1Pi16
//go:noescape
func MmSet1Pi16(r *x86.M64, v0 *x86.Short)

// Broadcast 8-bit integer "a" to all elements of "dst".
//
//go:linkname MmSet1Pi8 MmSet1Pi8
//go:noescape
func MmSet1Pi8(r *x86.M64, v0 *x86.Char)

// Set packed 32-bit integers in "dst" with the supplied values in reverse order.
//
//go:linkname MmSetrPi32 MmSetrPi32
//go:noescape
func MmSetrPi32(r *x86.M64, v0 *x86.Int, v1 *x86.Int)

// Set packed 16-bit integers in "dst" with the supplied values in reverse order.
//
//go:linkname MmSetrPi16 MmSetrPi16
//go:noescape
func MmSetrPi16(r *x86.M64, v0 *x86.Short, v1 *x86.Short, v2 *x86.Short, v3 *x86.Short)

// Set packed 8-bit integers in "dst" with the supplied values in reverse order.
//
//go:linkname MmSetrPi8 MmSetrPi8
//go:noescape
func MmSetrPi8(r *x86.M64, v0 *x86.Char, v1 *x86.Char, v2 *x86.Char, v3 *x86.Char, v4 *x86.Char, v5 *x86.Char, v6 *x86.Char, v7 *x86.Char)
