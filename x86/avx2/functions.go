package avx2

import (
	"github.com/alivanz/go-simd/x86"
)

/*
#cgo CFLAGS: -mavx2
#include <immintrin.h>
*/
import "C"

// Compute the absolute value of packed signed 8-bit integers in "a", and store the unsigned results in "dst".
//
//go:linkname Mm256AbsEpi8 Mm256AbsEpi8
//go:noescape
func Mm256AbsEpi8(r *x86.M256I, v0 *x86.M256I)

// Compute the absolute value of packed signed 16-bit integers in "a", and store the unsigned results in "dst".
//
//go:linkname Mm256AbsEpi16 Mm256AbsEpi16
//go:noescape
func Mm256AbsEpi16(r *x86.M256I, v0 *x86.M256I)

// Compute the absolute value of packed signed 32-bit integers in "a", and store the unsigned results in "dst".
//
//go:linkname Mm256AbsEpi32 Mm256AbsEpi32
//go:noescape
func Mm256AbsEpi32(r *x86.M256I, v0 *x86.M256I)

// Convert packed signed 16-bit integers from "a" and "b" to packed 8-bit integers using signed saturation, and store the results in "dst".
//
//go:linkname Mm256PacksEpi16 Mm256PacksEpi16
//go:noescape
func Mm256PacksEpi16(r *x86.M256I, v0 *x86.M256I, v1 *x86.M256I)

// Convert packed signed 32-bit integers from "a" and "b" to packed 16-bit integers using signed saturation, and store the results in "dst".
//
//go:linkname Mm256PacksEpi32 Mm256PacksEpi32
//go:noescape
func Mm256PacksEpi32(r *x86.M256I, v0 *x86.M256I, v1 *x86.M256I)

// Convert packed signed 16-bit integers from "a" and "b" to packed 8-bit integers using unsigned saturation, and store the results in "dst".
//
//go:linkname Mm256PackusEpi16 Mm256PackusEpi16
//go:noescape
func Mm256PackusEpi16(r *x86.M256I, v0 *x86.M256I, v1 *x86.M256I)

// Convert packed signed 32-bit integers from "a" and "b" to packed 16-bit integers using unsigned saturation, and store the results in "dst".
//
//go:linkname Mm256PackusEpi32 Mm256PackusEpi32
//go:noescape
func Mm256PackusEpi32(r *x86.M256I, v0 *x86.M256I, v1 *x86.M256I)

// Add packed 8-bit integers in "a" and "b", and store the results in "dst".
//
//go:linkname Mm256AddEpi8 Mm256AddEpi8
//go:noescape
func Mm256AddEpi8(r *x86.M256I, v0 *x86.M256I, v1 *x86.M256I)

// Add packed 16-bit integers in "a" and "b", and store the results in "dst".
//
//go:linkname Mm256AddEpi16 Mm256AddEpi16
//go:noescape
func Mm256AddEpi16(r *x86.M256I, v0 *x86.M256I, v1 *x86.M256I)

// Add packed 32-bit integers in "a" and "b", and store the results in "dst".
//
//go:linkname Mm256AddEpi32 Mm256AddEpi32
//go:noescape
func Mm256AddEpi32(r *x86.M256I, v0 *x86.M256I, v1 *x86.M256I)

// Add packed 64-bit integers in "a" and "b", and store the results in "dst".
//
//go:linkname Mm256AddEpi64 Mm256AddEpi64
//go:noescape
func Mm256AddEpi64(r *x86.M256I, v0 *x86.M256I, v1 *x86.M256I)

// Add packed 8-bit integers in "a" and "b" using saturation, and store the results in "dst".
//
//go:linkname Mm256AddsEpi8 Mm256AddsEpi8
//go:noescape
func Mm256AddsEpi8(r *x86.M256I, v0 *x86.M256I, v1 *x86.M256I)

// Add packed 16-bit integers in "a" and "b" using saturation, and store the results in "dst".
//
//go:linkname Mm256AddsEpi16 Mm256AddsEpi16
//go:noescape
func Mm256AddsEpi16(r *x86.M256I, v0 *x86.M256I, v1 *x86.M256I)

// Add packed unsigned 8-bit integers in "a" and "b" using saturation, and store the results in "dst".
//
//go:linkname Mm256AddsEpu8 Mm256AddsEpu8
//go:noescape
func Mm256AddsEpu8(r *x86.M256I, v0 *x86.M256I, v1 *x86.M256I)

// Add packed unsigned 16-bit integers in "a" and "b" using saturation, and store the results in "dst".
//
//go:linkname Mm256AddsEpu16 Mm256AddsEpu16
//go:noescape
func Mm256AddsEpu16(r *x86.M256I, v0 *x86.M256I, v1 *x86.M256I)

// Compute the bitwise AND of 256 bits (representing integer data) in "a" and "b", and store the result in "dst".
//
//go:linkname Mm256AndSi256 Mm256AndSi256
//go:noescape
func Mm256AndSi256(r *x86.M256I, v0 *x86.M256I, v1 *x86.M256I)

// Compute the bitwise NOT of 256 bits (representing integer data) in "a" and then AND with "b", and store the result in "dst".
//
//go:linkname Mm256AndnotSi256 Mm256AndnotSi256
//go:noescape
func Mm256AndnotSi256(r *x86.M256I, v0 *x86.M256I, v1 *x86.M256I)

// Average packed unsigned 8-bit integers in "a" and "b", and store the results in "dst".
//
//go:linkname Mm256AvgEpu8 Mm256AvgEpu8
//go:noescape
func Mm256AvgEpu8(r *x86.M256I, v0 *x86.M256I, v1 *x86.M256I)

// Average packed unsigned 16-bit integers in "a" and "b", and store the results in "dst".
//
//go:linkname Mm256AvgEpu16 Mm256AvgEpu16
//go:noescape
func Mm256AvgEpu16(r *x86.M256I, v0 *x86.M256I, v1 *x86.M256I)

// Blend packed 8-bit integers from "a" and "b" using "mask", and store the results in "dst".
//
//go:linkname Mm256BlendvEpi8 Mm256BlendvEpi8
//go:noescape
func Mm256BlendvEpi8(r *x86.M256I, v0 *x86.M256I, v1 *x86.M256I, v2 *x86.M256I)

// Compare packed 8-bit integers in "a" and "b" for equality, and store the results in "dst".
//
//go:linkname Mm256CmpeqEpi8 Mm256CmpeqEpi8
//go:noescape
func Mm256CmpeqEpi8(r *x86.M256I, v0 *x86.M256I, v1 *x86.M256I)

// Compare packed 16-bit integers in "a" and "b" for equality, and store the results in "dst".
//
//go:linkname Mm256CmpeqEpi16 Mm256CmpeqEpi16
//go:noescape
func Mm256CmpeqEpi16(r *x86.M256I, v0 *x86.M256I, v1 *x86.M256I)

// Compare packed 32-bit integers in "a" and "b" for equality, and store the results in "dst".
//
//go:linkname Mm256CmpeqEpi32 Mm256CmpeqEpi32
//go:noescape
func Mm256CmpeqEpi32(r *x86.M256I, v0 *x86.M256I, v1 *x86.M256I)

// Compare packed 64-bit integers in "a" and "b" for equality, and store the results in "dst".
//
//go:linkname Mm256CmpeqEpi64 Mm256CmpeqEpi64
//go:noescape
func Mm256CmpeqEpi64(r *x86.M256I, v0 *x86.M256I, v1 *x86.M256I)

// Compare packed signed 8-bit integers in "a" and "b" for greater-than, and store the results in "dst".
//
//go:linkname Mm256CmpgtEpi8 Mm256CmpgtEpi8
//go:noescape
func Mm256CmpgtEpi8(r *x86.M256I, v0 *x86.M256I, v1 *x86.M256I)

// Compare packed signed 16-bit integers in "a" and "b" for greater-than, and store the results in "dst".
//
//go:linkname Mm256CmpgtEpi16 Mm256CmpgtEpi16
//go:noescape
func Mm256CmpgtEpi16(r *x86.M256I, v0 *x86.M256I, v1 *x86.M256I)

// Compare packed signed 32-bit integers in "a" and "b" for greater-than, and store the results in "dst".
//
//go:linkname Mm256CmpgtEpi32 Mm256CmpgtEpi32
//go:noescape
func Mm256CmpgtEpi32(r *x86.M256I, v0 *x86.M256I, v1 *x86.M256I)

// Compare packed signed 64-bit integers in "a" and "b" for greater-than, and store the results in "dst".
//
//go:linkname Mm256CmpgtEpi64 Mm256CmpgtEpi64
//go:noescape
func Mm256CmpgtEpi64(r *x86.M256I, v0 *x86.M256I, v1 *x86.M256I)

// Horizontally add adjacent pairs of 16-bit integers in "a" and "b", and pack the signed 16-bit results in "dst".
//
//go:linkname Mm256HaddEpi16 Mm256HaddEpi16
//go:noescape
func Mm256HaddEpi16(r *x86.M256I, v0 *x86.M256I, v1 *x86.M256I)

// Horizontally add adjacent pairs of 32-bit integers in "a" and "b", and pack the signed 32-bit results in "dst".
//
//go:linkname Mm256HaddEpi32 Mm256HaddEpi32
//go:noescape
func Mm256HaddEpi32(r *x86.M256I, v0 *x86.M256I, v1 *x86.M256I)

// Horizontally add adjacent pairs of signed 16-bit integers in "a" and "b" using saturation, and pack the signed 16-bit results in "dst".
//
//go:linkname Mm256HaddsEpi16 Mm256HaddsEpi16
//go:noescape
func Mm256HaddsEpi16(r *x86.M256I, v0 *x86.M256I, v1 *x86.M256I)

// Horizontally subtract adjacent pairs of 16-bit integers in "a" and "b", and pack the signed 16-bit results in "dst".
//
//go:linkname Mm256HsubEpi16 Mm256HsubEpi16
//go:noescape
func Mm256HsubEpi16(r *x86.M256I, v0 *x86.M256I, v1 *x86.M256I)

// Horizontally subtract adjacent pairs of 32-bit integers in "a" and "b", and pack the signed 32-bit results in "dst".
//
//go:linkname Mm256HsubEpi32 Mm256HsubEpi32
//go:noescape
func Mm256HsubEpi32(r *x86.M256I, v0 *x86.M256I, v1 *x86.M256I)

// Horizontally subtract adjacent pairs of signed 16-bit integers in "a" and "b" using saturation, and pack the signed 16-bit results in "dst".
//
//go:linkname Mm256HsubsEpi16 Mm256HsubsEpi16
//go:noescape
func Mm256HsubsEpi16(r *x86.M256I, v0 *x86.M256I, v1 *x86.M256I)

// Vertically multiply each unsigned 8-bit integer from "a" with the corresponding signed 8-bit integer from "b", producing intermediate signed 16-bit integers. Horizontally add adjacent pairs of intermediate signed 16-bit integers, and pack the saturated results in "dst".
//
//go:linkname Mm256MaddubsEpi16 Mm256MaddubsEpi16
//go:noescape
func Mm256MaddubsEpi16(r *x86.M256I, v0 *x86.M256I, v1 *x86.M256I)

// Multiply packed signed 16-bit integers in "a" and "b", producing intermediate signed 32-bit integers. Horizontally add adjacent pairs of intermediate 32-bit integers, and pack the results in "dst".
//
//go:linkname Mm256MaddEpi16 Mm256MaddEpi16
//go:noescape
func Mm256MaddEpi16(r *x86.M256I, v0 *x86.M256I, v1 *x86.M256I)

// Compare packed signed 8-bit integers in "a" and "b", and store packed maximum values in "dst".
//
//go:linkname Mm256MaxEpi8 Mm256MaxEpi8
//go:noescape
func Mm256MaxEpi8(r *x86.M256I, v0 *x86.M256I, v1 *x86.M256I)

// Compare packed signed 16-bit integers in "a" and "b", and store packed maximum values in "dst".
//
//go:linkname Mm256MaxEpi16 Mm256MaxEpi16
//go:noescape
func Mm256MaxEpi16(r *x86.M256I, v0 *x86.M256I, v1 *x86.M256I)

// Compare packed signed 32-bit integers in "a" and "b", and store packed maximum values in "dst".
//
//go:linkname Mm256MaxEpi32 Mm256MaxEpi32
//go:noescape
func Mm256MaxEpi32(r *x86.M256I, v0 *x86.M256I, v1 *x86.M256I)

// Compare packed unsigned 8-bit integers in "a" and "b", and store packed maximum values in "dst".
//
//go:linkname Mm256MaxEpu8 Mm256MaxEpu8
//go:noescape
func Mm256MaxEpu8(r *x86.M256I, v0 *x86.M256I, v1 *x86.M256I)

// Compare packed unsigned 16-bit integers in "a" and "b", and store packed maximum values in "dst".
//
//go:linkname Mm256MaxEpu16 Mm256MaxEpu16
//go:noescape
func Mm256MaxEpu16(r *x86.M256I, v0 *x86.M256I, v1 *x86.M256I)

// Compare packed unsigned 32-bit integers in "a" and "b", and store packed maximum values in "dst".
//
//go:linkname Mm256MaxEpu32 Mm256MaxEpu32
//go:noescape
func Mm256MaxEpu32(r *x86.M256I, v0 *x86.M256I, v1 *x86.M256I)

// Compare packed signed 8-bit integers in "a" and "b", and store packed minimum values in "dst".
//
//go:linkname Mm256MinEpi8 Mm256MinEpi8
//go:noescape
func Mm256MinEpi8(r *x86.M256I, v0 *x86.M256I, v1 *x86.M256I)

// Compare packed signed 16-bit integers in "a" and "b", and store packed minimum values in "dst".
//
//go:linkname Mm256MinEpi16 Mm256MinEpi16
//go:noescape
func Mm256MinEpi16(r *x86.M256I, v0 *x86.M256I, v1 *x86.M256I)

// Compare packed signed 32-bit integers in "a" and "b", and store packed minimum values in "dst".
//
//go:linkname Mm256MinEpi32 Mm256MinEpi32
//go:noescape
func Mm256MinEpi32(r *x86.M256I, v0 *x86.M256I, v1 *x86.M256I)

// Compare packed unsigned 8-bit integers in "a" and "b", and store packed minimum values in "dst".
//
//go:linkname Mm256MinEpu8 Mm256MinEpu8
//go:noescape
func Mm256MinEpu8(r *x86.M256I, v0 *x86.M256I, v1 *x86.M256I)

// Compare packed unsigned 16-bit integers in "a" and "b", and store packed minimum values in "dst".
//
//go:linkname Mm256MinEpu16 Mm256MinEpu16
//go:noescape
func Mm256MinEpu16(r *x86.M256I, v0 *x86.M256I, v1 *x86.M256I)

// Compare packed unsigned 32-bit integers in "a" and "b", and store packed minimum values in "dst".
//
//go:linkname Mm256MinEpu32 Mm256MinEpu32
//go:noescape
func Mm256MinEpu32(r *x86.M256I, v0 *x86.M256I, v1 *x86.M256I)

// Create mask from the most significant bit of each 8-bit element in "a", and store the result in "dst".
//
//go:linkname Mm256MovemaskEpi8 Mm256MovemaskEpi8
//go:noescape
func Mm256MovemaskEpi8(r *x86.Int, v0 *x86.M256I)

// Sign extend packed 8-bit integers in "a" to packed 16-bit integers, and store the results in "dst".
//
//go:linkname Mm256Cvtepi8Epi16 Mm256Cvtepi8Epi16
//go:noescape
func Mm256Cvtepi8Epi16(r *x86.M256I, v0 *x86.M128I)

// Sign extend packed 8-bit integers in "a" to packed 32-bit integers, and store the results in "dst".
//
//go:linkname Mm256Cvtepi8Epi32 Mm256Cvtepi8Epi32
//go:noescape
func Mm256Cvtepi8Epi32(r *x86.M256I, v0 *x86.M128I)

// Sign extend packed 8-bit integers in the low 8 bytes of "a" to packed 64-bit integers, and store the results in "dst".
//
//go:linkname Mm256Cvtepi8Epi64 Mm256Cvtepi8Epi64
//go:noescape
func Mm256Cvtepi8Epi64(r *x86.M256I, v0 *x86.M128I)

// Sign extend packed 16-bit integers in "a" to packed 32-bit integers, and store the results in "dst".
//
//go:linkname Mm256Cvtepi16Epi32 Mm256Cvtepi16Epi32
//go:noescape
func Mm256Cvtepi16Epi32(r *x86.M256I, v0 *x86.M128I)

// Sign extend packed 16-bit integers in "a" to packed 64-bit integers, and store the results in "dst".
//
//go:linkname Mm256Cvtepi16Epi64 Mm256Cvtepi16Epi64
//go:noescape
func Mm256Cvtepi16Epi64(r *x86.M256I, v0 *x86.M128I)

// Sign extend packed 32-bit integers in "a" to packed 64-bit integers, and store the results in "dst".
//
//go:linkname Mm256Cvtepi32Epi64 Mm256Cvtepi32Epi64
//go:noescape
func Mm256Cvtepi32Epi64(r *x86.M256I, v0 *x86.M128I)

// Zero extend packed unsigned 8-bit integers in "a" to packed 16-bit integers, and store the results in "dst".
//
//go:linkname Mm256Cvtepu8Epi16 Mm256Cvtepu8Epi16
//go:noescape
func Mm256Cvtepu8Epi16(r *x86.M256I, v0 *x86.M128I)

// Zero extend packed unsigned 8-bit integers in "a" to packed 32-bit integers, and store the results in "dst".
//
//go:linkname Mm256Cvtepu8Epi32 Mm256Cvtepu8Epi32
//go:noescape
func Mm256Cvtepu8Epi32(r *x86.M256I, v0 *x86.M128I)

// Zero extend packed unsigned 8-bit integers in the low 8 byte sof "a" to packed 64-bit integers, and store the results in "dst".
//
//go:linkname Mm256Cvtepu8Epi64 Mm256Cvtepu8Epi64
//go:noescape
func Mm256Cvtepu8Epi64(r *x86.M256I, v0 *x86.M128I)

// Zero extend packed unsigned 16-bit integers in "a" to packed 32-bit integers, and store the results in "dst".
//
//go:linkname Mm256Cvtepu16Epi32 Mm256Cvtepu16Epi32
//go:noescape
func Mm256Cvtepu16Epi32(r *x86.M256I, v0 *x86.M128I)

// Zero extend packed unsigned 16-bit integers in "a" to packed 64-bit integers, and store the results in "dst".
//
//go:linkname Mm256Cvtepu16Epi64 Mm256Cvtepu16Epi64
//go:noescape
func Mm256Cvtepu16Epi64(r *x86.M256I, v0 *x86.M128I)

// Zero extend packed unsigned 32-bit integers in "a" to packed 64-bit integers, and store the results in "dst".
//
//go:linkname Mm256Cvtepu32Epi64 Mm256Cvtepu32Epi64
//go:noescape
func Mm256Cvtepu32Epi64(r *x86.M256I, v0 *x86.M128I)

// Multiply the low signed 32-bit integers from each packed 64-bit element in "a" and "b", and store the signed 64-bit results in "dst".
//
//go:linkname Mm256MulEpi32 Mm256MulEpi32
//go:noescape
func Mm256MulEpi32(r *x86.M256I, v0 *x86.M256I, v1 *x86.M256I)

// Multiply packed signed 16-bit integers in "a" and "b", producing intermediate signed 32-bit integers. Truncate each intermediate integer to the 18 most significant bits, round by adding 1, and store bits [16:1] to "dst".
//
//go:linkname Mm256MulhrsEpi16 Mm256MulhrsEpi16
//go:noescape
func Mm256MulhrsEpi16(r *x86.M256I, v0 *x86.M256I, v1 *x86.M256I)

// Multiply the packed unsigned 16-bit integers in "a" and "b", producing intermediate 32-bit integers, and store the high 16 bits of the intermediate integers in "dst".
//
//go:linkname Mm256MulhiEpu16 Mm256MulhiEpu16
//go:noescape
func Mm256MulhiEpu16(r *x86.M256I, v0 *x86.M256I, v1 *x86.M256I)

// Multiply the packed signed 16-bit integers in "a" and "b", producing intermediate 32-bit integers, and store the high 16 bits of the intermediate integers in "dst".
//
//go:linkname Mm256MulhiEpi16 Mm256MulhiEpi16
//go:noescape
func Mm256MulhiEpi16(r *x86.M256I, v0 *x86.M256I, v1 *x86.M256I)

// Multiply the packed signed 16-bit integers in "a" and "b", producing intermediate 32-bit integers, and store the low 16 bits of the intermediate integers in "dst".
//
//go:linkname Mm256MulloEpi16 Mm256MulloEpi16
//go:noescape
func Mm256MulloEpi16(r *x86.M256I, v0 *x86.M256I, v1 *x86.M256I)

// Multiply the packed signed 32-bit integers in "a" and "b", producing intermediate 64-bit integers, and store the low 32 bits of the intermediate integers in "dst".
//
//go:linkname Mm256MulloEpi32 Mm256MulloEpi32
//go:noescape
func Mm256MulloEpi32(r *x86.M256I, v0 *x86.M256I, v1 *x86.M256I)

// Multiply the low unsigned 32-bit integers from each packed 64-bit element in "a" and "b", and store the unsigned 64-bit results in "dst".
//
//go:linkname Mm256MulEpu32 Mm256MulEpu32
//go:noescape
func Mm256MulEpu32(r *x86.M256I, v0 *x86.M256I, v1 *x86.M256I)

// Compute the bitwise OR of 256 bits (representing integer data) in "a" and "b", and store the result in "dst".
//
//go:linkname Mm256OrSi256 Mm256OrSi256
//go:noescape
func Mm256OrSi256(r *x86.M256I, v0 *x86.M256I, v1 *x86.M256I)

// Compute the absolute differences of packed unsigned 8-bit integers in "a" and "b", then horizontally sum each consecutive 8 differences to produce four unsigned 16-bit integers, and pack these unsigned 16-bit integers in the low 16 bits of 64-bit elements in "dst".
//
//go:linkname Mm256SadEpu8 Mm256SadEpu8
//go:noescape
func Mm256SadEpu8(r *x86.M256I, v0 *x86.M256I, v1 *x86.M256I)

// Shuffle 8-bit integers in "a" within 128-bit lanes according to shuffle control mask in the corresponding 8-bit element of "b", and store the results in "dst".
//
//go:linkname Mm256ShuffleEpi8 Mm256ShuffleEpi8
//go:noescape
func Mm256ShuffleEpi8(r *x86.M256I, v0 *x86.M256I, v1 *x86.M256I)

// Negate packed signed 8-bit integers in "a" when the corresponding signed 8-bit integer in "b" is negative, and store the results in "dst". Element in "dst" are zeroed out when the corresponding element in "b" is zero.
//
//go:linkname Mm256SignEpi8 Mm256SignEpi8
//go:noescape
func Mm256SignEpi8(r *x86.M256I, v0 *x86.M256I, v1 *x86.M256I)

// Negate packed signed 16-bit integers in "a" when the corresponding signed 16-bit integer in "b" is negative, and store the results in "dst". Element in "dst" are zeroed out when the corresponding element in "b" is zero.
//
//go:linkname Mm256SignEpi16 Mm256SignEpi16
//go:noescape
func Mm256SignEpi16(r *x86.M256I, v0 *x86.M256I, v1 *x86.M256I)

// Negate packed signed 32-bit integers in "a" when the corresponding signed 32-bit integer in "b" is negative, and store the results in "dst". Element in "dst" are zeroed out when the corresponding element in "b" is zero.
//
//go:linkname Mm256SignEpi32 Mm256SignEpi32
//go:noescape
func Mm256SignEpi32(r *x86.M256I, v0 *x86.M256I, v1 *x86.M256I)

// Shift packed 16-bit integers in "a" left by "imm8" while shifting in zeros, and store the results in "dst".
//
//go:linkname Mm256SlliEpi16 Mm256SlliEpi16
//go:noescape
func Mm256SlliEpi16(r *x86.M256I, v0 *x86.M256I, v1 *x86.Int)

// Shift packed 16-bit integers in "a" left by "count" while shifting in zeros, and store the results in "dst".
//
//go:linkname Mm256SllEpi16 Mm256SllEpi16
//go:noescape
func Mm256SllEpi16(r *x86.M256I, v0 *x86.M256I, v1 *x86.M128I)

// Shift packed 32-bit integers in "a" left by "imm8" while shifting in zeros, and store the results in "dst".
//
//go:linkname Mm256SlliEpi32 Mm256SlliEpi32
//go:noescape
func Mm256SlliEpi32(r *x86.M256I, v0 *x86.M256I, v1 *x86.Int)

// Shift packed 32-bit integers in "a" left by "count" while shifting in zeros, and store the results in "dst".
//
//go:linkname Mm256SllEpi32 Mm256SllEpi32
//go:noescape
func Mm256SllEpi32(r *x86.M256I, v0 *x86.M256I, v1 *x86.M128I)

// Shift packed 64-bit integers in "a" left by "imm8" while shifting in zeros, and store the results in "dst".
//
//go:linkname Mm256SlliEpi64 Mm256SlliEpi64
//go:noescape
func Mm256SlliEpi64(r *x86.M256I, v0 *x86.M256I, v1 *x86.Int)

// Shift packed 64-bit integers in "a" left by "count" while shifting in zeros, and store the results in "dst".
//
//go:linkname Mm256SllEpi64 Mm256SllEpi64
//go:noescape
func Mm256SllEpi64(r *x86.M256I, v0 *x86.M256I, v1 *x86.M128I)

// Shift packed 16-bit integers in "a" right by "imm8" while shifting in sign bits, and store the results in "dst".
//
//go:linkname Mm256SraiEpi16 Mm256SraiEpi16
//go:noescape
func Mm256SraiEpi16(r *x86.M256I, v0 *x86.M256I, v1 *x86.Int)

// Shift packed 16-bit integers in "a" right by "count" while shifting in sign bits, and store the results in "dst".
//
//go:linkname Mm256SraEpi16 Mm256SraEpi16
//go:noescape
func Mm256SraEpi16(r *x86.M256I, v0 *x86.M256I, v1 *x86.M128I)

// Shift packed 32-bit integers in "a" right by "imm8" while shifting in sign bits, and store the results in "dst".
//
//go:linkname Mm256SraiEpi32 Mm256SraiEpi32
//go:noescape
func Mm256SraiEpi32(r *x86.M256I, v0 *x86.M256I, v1 *x86.Int)

// Shift packed 32-bit integers in "a" right by "count" while shifting in sign bits, and store the results in "dst".
//
//go:linkname Mm256SraEpi32 Mm256SraEpi32
//go:noescape
func Mm256SraEpi32(r *x86.M256I, v0 *x86.M256I, v1 *x86.M128I)

// Shift packed 16-bit integers in "a" right by "imm8" while shifting in zeros, and store the results in "dst".
//
//go:linkname Mm256SrliEpi16 Mm256SrliEpi16
//go:noescape
func Mm256SrliEpi16(r *x86.M256I, v0 *x86.M256I, v1 *x86.Int)

// Shift packed 16-bit integers in "a" right by "count" while shifting in zeros, and store the results in "dst".
//
//go:linkname Mm256SrlEpi16 Mm256SrlEpi16
//go:noescape
func Mm256SrlEpi16(r *x86.M256I, v0 *x86.M256I, v1 *x86.M128I)

// Shift packed 32-bit integers in "a" right by "imm8" while shifting in zeros, and store the results in "dst".
//
//go:linkname Mm256SrliEpi32 Mm256SrliEpi32
//go:noescape
func Mm256SrliEpi32(r *x86.M256I, v0 *x86.M256I, v1 *x86.Int)

// Shift packed 32-bit integers in "a" right by "count" while shifting in zeros, and store the results in "dst".
//
//go:linkname Mm256SrlEpi32 Mm256SrlEpi32
//go:noescape
func Mm256SrlEpi32(r *x86.M256I, v0 *x86.M256I, v1 *x86.M128I)

// Shift packed 64-bit integers in "a" right by "imm8" while shifting in zeros, and store the results in "dst".
//
//go:linkname Mm256SrliEpi64 Mm256SrliEpi64
//go:noescape
func Mm256SrliEpi64(r *x86.M256I, v0 *x86.M256I, v1 *x86.Int)

// Shift packed 64-bit integers in "a" right by "count" while shifting in zeros, and store the results in "dst".
//
//go:linkname Mm256SrlEpi64 Mm256SrlEpi64
//go:noescape
func Mm256SrlEpi64(r *x86.M256I, v0 *x86.M256I, v1 *x86.M128I)

// Subtract packed 8-bit integers in "b" from packed 8-bit integers in "a", and store the results in "dst".
//
//go:linkname Mm256SubEpi8 Mm256SubEpi8
//go:noescape
func Mm256SubEpi8(r *x86.M256I, v0 *x86.M256I, v1 *x86.M256I)

// Subtract packed 16-bit integers in "b" from packed 16-bit integers in "a", and store the results in "dst".
//
//go:linkname Mm256SubEpi16 Mm256SubEpi16
//go:noescape
func Mm256SubEpi16(r *x86.M256I, v0 *x86.M256I, v1 *x86.M256I)

// Subtract packed 32-bit integers in "b" from packed 32-bit integers in "a", and store the results in "dst".
//
//go:linkname Mm256SubEpi32 Mm256SubEpi32
//go:noescape
func Mm256SubEpi32(r *x86.M256I, v0 *x86.M256I, v1 *x86.M256I)

// Subtract packed 64-bit integers in "b" from packed 64-bit integers in "a", and store the results in "dst".
//
//go:linkname Mm256SubEpi64 Mm256SubEpi64
//go:noescape
func Mm256SubEpi64(r *x86.M256I, v0 *x86.M256I, v1 *x86.M256I)

// Subtract packed signed 8-bit integers in "b" from packed 8-bit integers in "a" using saturation, and store the results in "dst".
//
//go:linkname Mm256SubsEpi8 Mm256SubsEpi8
//go:noescape
func Mm256SubsEpi8(r *x86.M256I, v0 *x86.M256I, v1 *x86.M256I)

// Subtract packed signed 16-bit integers in "b" from packed 16-bit integers in "a" using saturation, and store the results in "dst".
//
//go:linkname Mm256SubsEpi16 Mm256SubsEpi16
//go:noescape
func Mm256SubsEpi16(r *x86.M256I, v0 *x86.M256I, v1 *x86.M256I)

// Subtract packed unsigned 8-bit integers in "b" from packed unsigned 8-bit integers in "a" using saturation, and store the results in "dst".
//
//go:linkname Mm256SubsEpu8 Mm256SubsEpu8
//go:noescape
func Mm256SubsEpu8(r *x86.M256I, v0 *x86.M256I, v1 *x86.M256I)

// Subtract packed unsigned 16-bit integers in "b" from packed unsigned 16-bit integers in "a" using saturation, and store the results in "dst".
//
//go:linkname Mm256SubsEpu16 Mm256SubsEpu16
//go:noescape
func Mm256SubsEpu16(r *x86.M256I, v0 *x86.M256I, v1 *x86.M256I)

// Unpack and interleave 8-bit integers from the high half of each 128-bit lane in "a" and "b", and store the results in "dst".
//
//go:linkname Mm256UnpackhiEpi8 Mm256UnpackhiEpi8
//go:noescape
func Mm256UnpackhiEpi8(r *x86.M256I, v0 *x86.M256I, v1 *x86.M256I)

// Unpack and interleave 16-bit integers from the high half of each 128-bit lane in "a" and "b", and store the results in "dst".
//
//go:linkname Mm256UnpackhiEpi16 Mm256UnpackhiEpi16
//go:noescape
func Mm256UnpackhiEpi16(r *x86.M256I, v0 *x86.M256I, v1 *x86.M256I)

// Unpack and interleave 32-bit integers from the high half of each 128-bit lane in "a" and "b", and store the results in "dst".
//
//go:linkname Mm256UnpackhiEpi32 Mm256UnpackhiEpi32
//go:noescape
func Mm256UnpackhiEpi32(r *x86.M256I, v0 *x86.M256I, v1 *x86.M256I)

// Unpack and interleave 64-bit integers from the high half of each 128-bit lane in "a" and "b", and store the results in "dst".
//
//go:linkname Mm256UnpackhiEpi64 Mm256UnpackhiEpi64
//go:noescape
func Mm256UnpackhiEpi64(r *x86.M256I, v0 *x86.M256I, v1 *x86.M256I)

// Unpack and interleave 8-bit integers from the low half of each 128-bit lane in "a" and "b", and store the results in "dst".
//
//go:linkname Mm256UnpackloEpi8 Mm256UnpackloEpi8
//go:noescape
func Mm256UnpackloEpi8(r *x86.M256I, v0 *x86.M256I, v1 *x86.M256I)

// Unpack and interleave 16-bit integers from the low half of each 128-bit lane in "a" and "b", and store the results in "dst".
//
//go:linkname Mm256UnpackloEpi16 Mm256UnpackloEpi16
//go:noescape
func Mm256UnpackloEpi16(r *x86.M256I, v0 *x86.M256I, v1 *x86.M256I)

// Unpack and interleave 32-bit integers from the low half of each 128-bit lane in "a" and "b", and store the results in "dst".
//
//go:linkname Mm256UnpackloEpi32 Mm256UnpackloEpi32
//go:noescape
func Mm256UnpackloEpi32(r *x86.M256I, v0 *x86.M256I, v1 *x86.M256I)

// Unpack and interleave 64-bit integers from the low half of each 128-bit lane in "a" and "b", and store the results in "dst".
//
//go:linkname Mm256UnpackloEpi64 Mm256UnpackloEpi64
//go:noescape
func Mm256UnpackloEpi64(r *x86.M256I, v0 *x86.M256I, v1 *x86.M256I)

// Compute the bitwise XOR of 256 bits (representing integer data) in "a" and "b", and store the result in "dst".
//
//go:linkname Mm256XorSi256 Mm256XorSi256
//go:noescape
func Mm256XorSi256(r *x86.M256I, v0 *x86.M256I, v1 *x86.M256I)

// Broadcast the low single-precision (32-bit) floating-point element from "a" to all elements of "dst".
//
//go:linkname MmBroadcastssPs MmBroadcastssPs
//go:noescape
func MmBroadcastssPs(r *x86.M128, v0 *x86.M128)

// Broadcast the low double-precision (64-bit) floating-point element from "a" to all elements of "dst".
//
//go:linkname MmBroadcastsdPd MmBroadcastsdPd
//go:noescape
func MmBroadcastsdPd(r *x86.M128D, v0 *x86.M128D)

// Broadcast the low single-precision (32-bit) floating-point element from "a" to all elements of "dst".
//
//go:linkname Mm256BroadcastssPs Mm256BroadcastssPs
//go:noescape
func Mm256BroadcastssPs(r *x86.M256, v0 *x86.M128)

// Broadcast the low double-precision (64-bit) floating-point element from "a" to all elements of "dst".
//
//go:linkname Mm256BroadcastsdPd Mm256BroadcastsdPd
//go:noescape
func Mm256BroadcastsdPd(r *x86.M256D, v0 *x86.M128D)

// Broadcast 128 bits of integer data from "a" to all 128-bit lanes in "dst".
//
//go:linkname Mm256Broadcastsi128Si256 Mm256Broadcastsi128Si256
//go:noescape
func Mm256Broadcastsi128Si256(r *x86.M256I, v0 *x86.M128I)

// Broadcast the low packed 8-bit integer from "a" to all elements of "dst".
//
//go:linkname Mm256BroadcastbEpi8 Mm256BroadcastbEpi8
//go:noescape
func Mm256BroadcastbEpi8(r *x86.M256I, v0 *x86.M128I)

// Broadcast the low packed 16-bit integer from "a" to all elements of "dst".
//
//go:linkname Mm256BroadcastwEpi16 Mm256BroadcastwEpi16
//go:noescape
func Mm256BroadcastwEpi16(r *x86.M256I, v0 *x86.M128I)

// Broadcast the low packed 32-bit integer from "a" to all elements of "dst".
//
//go:linkname Mm256BroadcastdEpi32 Mm256BroadcastdEpi32
//go:noescape
func Mm256BroadcastdEpi32(r *x86.M256I, v0 *x86.M128I)

// Broadcast the low packed 64-bit integer from "a" to all elements of "dst".
//
//go:linkname Mm256BroadcastqEpi64 Mm256BroadcastqEpi64
//go:noescape
func Mm256BroadcastqEpi64(r *x86.M256I, v0 *x86.M128I)

// Broadcast the low packed 8-bit integer from "a" to all elements of "dst".
//
//go:linkname MmBroadcastbEpi8 MmBroadcastbEpi8
//go:noescape
func MmBroadcastbEpi8(r *x86.M128I, v0 *x86.M128I)

// Broadcast the low packed 16-bit integer from "a" to all elements of "dst".
//
//go:linkname MmBroadcastwEpi16 MmBroadcastwEpi16
//go:noescape
func MmBroadcastwEpi16(r *x86.M128I, v0 *x86.M128I)

// Broadcast the low packed 32-bit integer from "a" to all elements of "dst".
//
//go:linkname MmBroadcastdEpi32 MmBroadcastdEpi32
//go:noescape
func MmBroadcastdEpi32(r *x86.M128I, v0 *x86.M128I)

// Broadcast the low packed 64-bit integer from "a" to all elements of "dst".
//
//go:linkname MmBroadcastqEpi64 MmBroadcastqEpi64
//go:noescape
func MmBroadcastqEpi64(r *x86.M128I, v0 *x86.M128I)

// Shuffle 32-bit integers in "a" across lanes using the corresponding index in "idx", and store the results in "dst".
//
//go:linkname Mm256Permutevar8X32Epi32 Mm256Permutevar8X32Epi32
//go:noescape
func Mm256Permutevar8X32Epi32(r *x86.M256I, v0 *x86.M256I, v1 *x86.M256I)

// Shuffle single-precision (32-bit) floating-point elements in "a" across lanes using the corresponding index in "idx".
//
//go:linkname Mm256Permutevar8X32Ps Mm256Permutevar8X32Ps
//go:noescape
func Mm256Permutevar8X32Ps(r *x86.M256, v0 *x86.M256, v1 *x86.M256I)

// Shift packed 32-bit integers in "a" left by the amount specified by the corresponding element in "count" while shifting in zeros, and store the results in "dst".
//
//go:linkname Mm256SllvEpi32 Mm256SllvEpi32
//go:noescape
func Mm256SllvEpi32(r *x86.M256I, v0 *x86.M256I, v1 *x86.M256I)

// Shift packed 32-bit integers in "a" left by the amount specified by the corresponding element in "count" while shifting in zeros, and store the results in "dst".
//
//go:linkname MmSllvEpi32 MmSllvEpi32
//go:noescape
func MmSllvEpi32(r *x86.M128I, v0 *x86.M128I, v1 *x86.M128I)

// Shift packed 64-bit integers in "a" left by the amount specified by the corresponding element in "count" while shifting in zeros, and store the results in "dst".
//
//go:linkname Mm256SllvEpi64 Mm256SllvEpi64
//go:noescape
func Mm256SllvEpi64(r *x86.M256I, v0 *x86.M256I, v1 *x86.M256I)

// Shift packed 64-bit integers in "a" left by the amount specified by the corresponding element in "count" while shifting in zeros, and store the results in "dst".
//
//go:linkname MmSllvEpi64 MmSllvEpi64
//go:noescape
func MmSllvEpi64(r *x86.M128I, v0 *x86.M128I, v1 *x86.M128I)

// Shift packed 32-bit integers in "a" right by the amount specified by the corresponding element in "count" while shifting in sign bits, and store the results in "dst".
//
//go:linkname Mm256SravEpi32 Mm256SravEpi32
//go:noescape
func Mm256SravEpi32(r *x86.M256I, v0 *x86.M256I, v1 *x86.M256I)

// Shift packed 32-bit integers in "a" right by the amount specified by the corresponding element in "count" while shifting in sign bits, and store the results in "dst".
//
//go:linkname MmSravEpi32 MmSravEpi32
//go:noescape
func MmSravEpi32(r *x86.M128I, v0 *x86.M128I, v1 *x86.M128I)

// Shift packed 32-bit integers in "a" right by the amount specified by the corresponding element in "count" while shifting in zeros, and store the results in "dst".
//
//go:linkname Mm256SrlvEpi32 Mm256SrlvEpi32
//go:noescape
func Mm256SrlvEpi32(r *x86.M256I, v0 *x86.M256I, v1 *x86.M256I)

// Shift packed 32-bit integers in "a" right by the amount specified by the corresponding element in "count" while shifting in zeros, and store the results in "dst".
//
//go:linkname MmSrlvEpi32 MmSrlvEpi32
//go:noescape
func MmSrlvEpi32(r *x86.M128I, v0 *x86.M128I, v1 *x86.M128I)

// Shift packed 64-bit integers in "a" right by the amount specified by the corresponding element in "count" while shifting in zeros, and store the results in "dst".
//
//go:linkname Mm256SrlvEpi64 Mm256SrlvEpi64
//go:noescape
func Mm256SrlvEpi64(r *x86.M256I, v0 *x86.M256I, v1 *x86.M256I)

// Shift packed 64-bit integers in "a" right by the amount specified by the corresponding element in "count" while shifting in zeros, and store the results in "dst".
//
//go:linkname MmSrlvEpi64 MmSrlvEpi64
//go:noescape
func MmSrlvEpi64(r *x86.M128I, v0 *x86.M128I, v1 *x86.M128I)
