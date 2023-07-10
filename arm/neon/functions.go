package neon

import (
	"github.com/alivanz/go-simd/arm"
)

/*
#include <arm_neon.h>
*/
import "C"

// Unsigned Absolute Difference (vector). This instruction subtracts the elements of the vector of the second source SIMD&FP register from the corresponding elements of the first source SIMD&FP register, places the the absolute values of the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VabdqU8 VabdqU8
//go:noescape
func VabdqU8(r *arm.Uint8X16, v0 *arm.Uint8X16, v1 *arm.Uint8X16)

// Unsigned Absolute Difference (vector). This instruction subtracts the elements of the vector of the second source SIMD&FP register from the corresponding elements of the first source SIMD&FP register, places the the absolute values of the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VabdqU32 VabdqU32
//go:noescape
func VabdqU32(r *arm.Uint32X4, v0 *arm.Uint32X4, v1 *arm.Uint32X4)

// Unsigned Absolute Difference (vector). This instruction subtracts the elements of the vector of the second source SIMD&FP register from the corresponding elements of the first source SIMD&FP register, places the the absolute values of the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VabdqU16 VabdqU16
//go:noescape
func VabdqU16(r *arm.Uint16X8, v0 *arm.Uint16X8, v1 *arm.Uint16X8)

// Signed Absolute Difference. This instruction subtracts the elements of the vector of the second source SIMD&FP register from the corresponding elements of the first source SIMD&FP register, places the the absolute values of the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VabdqS8 VabdqS8
//go:noescape
func VabdqS8(r *arm.Int8X16, v0 *arm.Int8X16, v1 *arm.Int8X16)

// Floating-point Absolute Difference (vector). This instruction subtracts the floating-point values in the elements of the second source SIMD&FP register, from the corresponding floating-point values in the elements of the first source SIMD&FP register, places the absolute value of each result in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VabdqF32 VabdqF32
//go:noescape
func VabdqF32(r *arm.Float32X4, v0 *arm.Float32X4, v1 *arm.Float32X4)

// Signed Absolute Difference. This instruction subtracts the elements of the vector of the second source SIMD&FP register from the corresponding elements of the first source SIMD&FP register, places the the absolute values of the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VabdqS32 VabdqS32
//go:noescape
func VabdqS32(r *arm.Int32X4, v0 *arm.Int32X4, v1 *arm.Int32X4)

// Signed Absolute Difference. This instruction subtracts the elements of the vector of the second source SIMD&FP register from the corresponding elements of the first source SIMD&FP register, places the the absolute values of the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VabdqS16 VabdqS16
//go:noescape
func VabdqS16(r *arm.Int16X8, v0 *arm.Int16X8, v1 *arm.Int16X8)

// Unsigned Absolute Difference (vector). This instruction subtracts the elements of the vector of the second source SIMD&FP register from the corresponding elements of the first source SIMD&FP register, places the the absolute values of the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VabdU8 VabdU8
//go:noescape
func VabdU8(r *arm.Uint8X8, v0 *arm.Uint8X8, v1 *arm.Uint8X8)

// Unsigned Absolute Difference (vector). This instruction subtracts the elements of the vector of the second source SIMD&FP register from the corresponding elements of the first source SIMD&FP register, places the the absolute values of the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VabdU32 VabdU32
//go:noescape
func VabdU32(r *arm.Uint32X2, v0 *arm.Uint32X2, v1 *arm.Uint32X2)

// Unsigned Absolute Difference (vector). This instruction subtracts the elements of the vector of the second source SIMD&FP register from the corresponding elements of the first source SIMD&FP register, places the the absolute values of the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VabdU16 VabdU16
//go:noescape
func VabdU16(r *arm.Uint16X4, v0 *arm.Uint16X4, v1 *arm.Uint16X4)

// Signed Absolute Difference. This instruction subtracts the elements of the vector of the second source SIMD&FP register from the corresponding elements of the first source SIMD&FP register, places the the absolute values of the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VabdS8 VabdS8
//go:noescape
func VabdS8(r *arm.Int8X8, v0 *arm.Int8X8, v1 *arm.Int8X8)

// Floating-point Absolute Difference (vector). This instruction subtracts the floating-point values in the elements of the second source SIMD&FP register, from the corresponding floating-point values in the elements of the first source SIMD&FP register, places the absolute value of each result in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VabdF32 VabdF32
//go:noescape
func VabdF32(r *arm.Float32X2, v0 *arm.Float32X2, v1 *arm.Float32X2)

// Signed Absolute Difference. This instruction subtracts the elements of the vector of the second source SIMD&FP register from the corresponding elements of the first source SIMD&FP register, places the the absolute values of the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VabdS32 VabdS32
//go:noescape
func VabdS32(r *arm.Int32X2, v0 *arm.Int32X2, v1 *arm.Int32X2)

// Signed Absolute Difference. This instruction subtracts the elements of the vector of the second source SIMD&FP register from the corresponding elements of the first source SIMD&FP register, places the the absolute values of the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VabdS16 VabdS16
//go:noescape
func VabdS16(r *arm.Int16X4, v0 *arm.Int16X4, v1 *arm.Int16X4)

// Absolute value (vector). This instruction calculates the absolute value of each vector element in the source SIMD&FP register, puts the result into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VabsqS8 VabsqS8
//go:noescape
func VabsqS8(r *arm.Int8X16, v0 *arm.Int8X16)

// Floating-point Absolute value (vector). This instruction calculates the absolute value of each vector element in the source SIMD&FP register, writes the result to a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VabsqF32 VabsqF32
//go:noescape
func VabsqF32(r *arm.Float32X4, v0 *arm.Float32X4)

// Absolute value (vector). This instruction calculates the absolute value of each vector element in the source SIMD&FP register, puts the result into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VabsqS32 VabsqS32
//go:noescape
func VabsqS32(r *arm.Int32X4, v0 *arm.Int32X4)

// Absolute value (vector). This instruction calculates the absolute value of each vector element in the source SIMD&FP register, puts the result into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VabsqS16 VabsqS16
//go:noescape
func VabsqS16(r *arm.Int16X8, v0 *arm.Int16X8)

// Absolute value (vector). This instruction calculates the absolute value of each vector element in the source SIMD&FP register, puts the result into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VabsS8 VabsS8
//go:noescape
func VabsS8(r *arm.Int8X8, v0 *arm.Int8X8)

// Floating-point Absolute value (vector). This instruction calculates the absolute value of each vector element in the source SIMD&FP register, writes the result to a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VabsF32 VabsF32
//go:noescape
func VabsF32(r *arm.Float32X2, v0 *arm.Float32X2)

// Absolute value (vector). This instruction calculates the absolute value of each vector element in the source SIMD&FP register, puts the result into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VabsS32 VabsS32
//go:noescape
func VabsS32(r *arm.Int32X2, v0 *arm.Int32X2)

// Absolute value (vector). This instruction calculates the absolute value of each vector element in the source SIMD&FP register, puts the result into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VabsS16 VabsS16
//go:noescape
func VabsS16(r *arm.Int16X4, v0 *arm.Int16X4)

// Add (vector). This instruction adds corresponding elements in the two source SIMD&FP registers, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VaddqU8 VaddqU8
//go:noescape
func VaddqU8(r *arm.Uint8X16, v0 *arm.Uint8X16, v1 *arm.Uint8X16)

// Add (vector). This instruction adds corresponding elements in the two source SIMD&FP registers, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VaddqU32 VaddqU32
//go:noescape
func VaddqU32(r *arm.Uint32X4, v0 *arm.Uint32X4, v1 *arm.Uint32X4)

// Add (vector). This instruction adds corresponding elements in the two source SIMD&FP registers, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VaddqU64 VaddqU64
//go:noescape
func VaddqU64(r *arm.Uint64X2, v0 *arm.Uint64X2, v1 *arm.Uint64X2)

// Add (vector). This instruction adds corresponding elements in the two source SIMD&FP registers, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VaddqU16 VaddqU16
//go:noescape
func VaddqU16(r *arm.Uint16X8, v0 *arm.Uint16X8, v1 *arm.Uint16X8)

// Add (vector). This instruction adds corresponding elements in the two source SIMD&FP registers, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VaddqS8 VaddqS8
//go:noescape
func VaddqS8(r *arm.Int8X16, v0 *arm.Int8X16, v1 *arm.Int8X16)

// Floating-point Add (vector). This instruction adds corresponding vector elements in the two source SIMD&FP registers, writes the result into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are floating-point values.
//
//go:linkname VaddqF32 VaddqF32
//go:noescape
func VaddqF32(r *arm.Float32X4, v0 *arm.Float32X4, v1 *arm.Float32X4)

// Add (vector). This instruction adds corresponding elements in the two source SIMD&FP registers, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VaddqS32 VaddqS32
//go:noescape
func VaddqS32(r *arm.Int32X4, v0 *arm.Int32X4, v1 *arm.Int32X4)

// Add (vector). This instruction adds corresponding elements in the two source SIMD&FP registers, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VaddqS64 VaddqS64
//go:noescape
func VaddqS64(r *arm.Int64X2, v0 *arm.Int64X2, v1 *arm.Int64X2)

// Add (vector). This instruction adds corresponding elements in the two source SIMD&FP registers, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VaddqS16 VaddqS16
//go:noescape
func VaddqS16(r *arm.Int16X8, v0 *arm.Int16X8, v1 *arm.Int16X8)

// Add (vector). This instruction adds corresponding elements in the two source SIMD&FP registers, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VaddU8 VaddU8
//go:noescape
func VaddU8(r *arm.Uint8X8, v0 *arm.Uint8X8, v1 *arm.Uint8X8)

// Add (vector). This instruction adds corresponding elements in the two source SIMD&FP registers, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VaddU32 VaddU32
//go:noescape
func VaddU32(r *arm.Uint32X2, v0 *arm.Uint32X2, v1 *arm.Uint32X2)

// Add (vector). This instruction adds corresponding elements in the two source SIMD&FP registers, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VaddU64 VaddU64
//go:noescape
func VaddU64(r *arm.Uint64X1, v0 *arm.Uint64X1, v1 *arm.Uint64X1)

// Add (vector). This instruction adds corresponding elements in the two source SIMD&FP registers, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VaddU16 VaddU16
//go:noescape
func VaddU16(r *arm.Uint16X4, v0 *arm.Uint16X4, v1 *arm.Uint16X4)

// Add (vector). This instruction adds corresponding elements in the two source SIMD&FP registers, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VaddS8 VaddS8
//go:noescape
func VaddS8(r *arm.Int8X8, v0 *arm.Int8X8, v1 *arm.Int8X8)

// Floating-point Add (vector). This instruction adds corresponding vector elements in the two source SIMD&FP registers, writes the result into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are floating-point values.
//
//go:linkname VaddF32 VaddF32
//go:noescape
func VaddF32(r *arm.Float32X2, v0 *arm.Float32X2, v1 *arm.Float32X2)

// Add (vector). This instruction adds corresponding elements in the two source SIMD&FP registers, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VaddS32 VaddS32
//go:noescape
func VaddS32(r *arm.Int32X2, v0 *arm.Int32X2, v1 *arm.Int32X2)

// Add (vector). This instruction adds corresponding elements in the two source SIMD&FP registers, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VaddS64 VaddS64
//go:noescape
func VaddS64(r *arm.Int64X1, v0 *arm.Int64X1, v1 *arm.Int64X1)

// Add (vector). This instruction adds corresponding elements in the two source SIMD&FP registers, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VaddS16 VaddS16
//go:noescape
func VaddS16(r *arm.Int16X4, v0 *arm.Int16X4, v1 *arm.Int16X4)

// Bitwise Exclusive OR (vector). This instruction performs a bitwise Exclusive OR operation between the two source SIMD&FP registers, and places the result in the destination SIMD&FP register.
//
//go:linkname VaddP8 VaddP8
//go:noescape
func VaddP8(r *arm.Poly8X8, v0 *arm.Poly8X8, v1 *arm.Poly8X8)

// Bitwise Exclusive OR (vector). This instruction performs a bitwise Exclusive OR operation between the two source SIMD&FP registers, and places the result in the destination SIMD&FP register.
//
//go:linkname VaddP64 VaddP64
//go:noescape
func VaddP64(r *arm.Poly64X1, v0 *arm.Poly64X1, v1 *arm.Poly64X1)

// Bitwise Exclusive OR (vector). This instruction performs a bitwise Exclusive OR operation between the two source SIMD&FP registers, and places the result in the destination SIMD&FP register.
//
//go:linkname VaddP16 VaddP16
//go:noescape
func VaddP16(r *arm.Poly16X4, v0 *arm.Poly16X4, v1 *arm.Poly16X4)

// Bitwise Exclusive OR (vector). This instruction performs a bitwise Exclusive OR operation between the two source SIMD&FP registers, and places the result in the destination SIMD&FP register.
//
//go:linkname VaddqP8 VaddqP8
//go:noescape
func VaddqP8(r *arm.Poly8X16, v0 *arm.Poly8X16, v1 *arm.Poly8X16)

// Bitwise Exclusive OR (vector). This instruction performs a bitwise Exclusive OR operation between the two source SIMD&FP registers, and places the result in the destination SIMD&FP register.
//
//go:linkname VaddqP64 VaddqP64
//go:noescape
func VaddqP64(r *arm.Poly64X2, v0 *arm.Poly64X2, v1 *arm.Poly64X2)

// Bitwise Exclusive OR (vector). This instruction performs a bitwise Exclusive OR operation between the two source SIMD&FP registers, and places the result in the destination SIMD&FP register.
//
//go:linkname VaddqP16 VaddqP16
//go:noescape
func VaddqP16(r *arm.Poly16X8, v0 *arm.Poly16X8, v1 *arm.Poly16X8)

// Add returning High Narrow. This instruction adds each vector element in the first source SIMD&FP register to the corresponding vector element in the second source SIMD&FP register, places the most significant half of the result into a vector, and writes the vector to the lower or upper half of the destination SIMD&FP register.
//
//go:linkname VaddhnU32 VaddhnU32
//go:noescape
func VaddhnU32(r *arm.Uint16X4, v0 *arm.Uint32X4, v1 *arm.Uint32X4)

// Add returning High Narrow. This instruction adds each vector element in the first source SIMD&FP register to the corresponding vector element in the second source SIMD&FP register, places the most significant half of the result into a vector, and writes the vector to the lower or upper half of the destination SIMD&FP register.
//
//go:linkname VaddhnU64 VaddhnU64
//go:noescape
func VaddhnU64(r *arm.Uint32X2, v0 *arm.Uint64X2, v1 *arm.Uint64X2)

// Add returning High Narrow. This instruction adds each vector element in the first source SIMD&FP register to the corresponding vector element in the second source SIMD&FP register, places the most significant half of the result into a vector, and writes the vector to the lower or upper half of the destination SIMD&FP register.
//
//go:linkname VaddhnU16 VaddhnU16
//go:noescape
func VaddhnU16(r *arm.Uint8X8, v0 *arm.Uint16X8, v1 *arm.Uint16X8)

// Add returning High Narrow. This instruction adds each vector element in the first source SIMD&FP register to the corresponding vector element in the second source SIMD&FP register, places the most significant half of the result into a vector, and writes the vector to the lower or upper half of the destination SIMD&FP register.
//
//go:linkname VaddhnS32 VaddhnS32
//go:noescape
func VaddhnS32(r *arm.Int16X4, v0 *arm.Int32X4, v1 *arm.Int32X4)

// Add returning High Narrow. This instruction adds each vector element in the first source SIMD&FP register to the corresponding vector element in the second source SIMD&FP register, places the most significant half of the result into a vector, and writes the vector to the lower or upper half of the destination SIMD&FP register.
//
//go:linkname VaddhnS64 VaddhnS64
//go:noescape
func VaddhnS64(r *arm.Int32X2, v0 *arm.Int64X2, v1 *arm.Int64X2)

// Add returning High Narrow. This instruction adds each vector element in the first source SIMD&FP register to the corresponding vector element in the second source SIMD&FP register, places the most significant half of the result into a vector, and writes the vector to the lower or upper half of the destination SIMD&FP register.
//
//go:linkname VaddhnS16 VaddhnS16
//go:noescape
func VaddhnS16(r *arm.Int8X8, v0 *arm.Int16X8, v1 *arm.Int16X8)

// Bitwise AND (vector). This instruction performs a bitwise AND between the two source SIMD&FP registers, and writes the result to the destination SIMD&FP register.
//
//go:linkname VandqU8 VandqU8
//go:noescape
func VandqU8(r *arm.Uint8X16, v0 *arm.Uint8X16, v1 *arm.Uint8X16)

// Bitwise AND (vector). This instruction performs a bitwise AND between the two source SIMD&FP registers, and writes the result to the destination SIMD&FP register.
//
//go:linkname VandqU32 VandqU32
//go:noescape
func VandqU32(r *arm.Uint32X4, v0 *arm.Uint32X4, v1 *arm.Uint32X4)

// Bitwise AND (vector). This instruction performs a bitwise AND between the two source SIMD&FP registers, and writes the result to the destination SIMD&FP register.
//
//go:linkname VandqU64 VandqU64
//go:noescape
func VandqU64(r *arm.Uint64X2, v0 *arm.Uint64X2, v1 *arm.Uint64X2)

// Bitwise AND (vector). This instruction performs a bitwise AND between the two source SIMD&FP registers, and writes the result to the destination SIMD&FP register.
//
//go:linkname VandqU16 VandqU16
//go:noescape
func VandqU16(r *arm.Uint16X8, v0 *arm.Uint16X8, v1 *arm.Uint16X8)

// Bitwise AND (vector). This instruction performs a bitwise AND between the two source SIMD&FP registers, and writes the result to the destination SIMD&FP register.
//
//go:linkname VandqS8 VandqS8
//go:noescape
func VandqS8(r *arm.Int8X16, v0 *arm.Int8X16, v1 *arm.Int8X16)

// Bitwise AND (vector). This instruction performs a bitwise AND between the two source SIMD&FP registers, and writes the result to the destination SIMD&FP register.
//
//go:linkname VandqS32 VandqS32
//go:noescape
func VandqS32(r *arm.Int32X4, v0 *arm.Int32X4, v1 *arm.Int32X4)

// Bitwise AND (vector). This instruction performs a bitwise AND between the two source SIMD&FP registers, and writes the result to the destination SIMD&FP register.
//
//go:linkname VandqS64 VandqS64
//go:noescape
func VandqS64(r *arm.Int64X2, v0 *arm.Int64X2, v1 *arm.Int64X2)

// Bitwise AND (vector). This instruction performs a bitwise AND between the two source SIMD&FP registers, and writes the result to the destination SIMD&FP register.
//
//go:linkname VandqS16 VandqS16
//go:noescape
func VandqS16(r *arm.Int16X8, v0 *arm.Int16X8, v1 *arm.Int16X8)

// Bitwise AND (vector). This instruction performs a bitwise AND between the two source SIMD&FP registers, and writes the result to the destination SIMD&FP register.
//
//go:linkname VandU8 VandU8
//go:noescape
func VandU8(r *arm.Uint8X8, v0 *arm.Uint8X8, v1 *arm.Uint8X8)

// Bitwise AND (vector). This instruction performs a bitwise AND between the two source SIMD&FP registers, and writes the result to the destination SIMD&FP register.
//
//go:linkname VandU32 VandU32
//go:noescape
func VandU32(r *arm.Uint32X2, v0 *arm.Uint32X2, v1 *arm.Uint32X2)

// Bitwise AND (vector). This instruction performs a bitwise AND between the two source SIMD&FP registers, and writes the result to the destination SIMD&FP register.
//
//go:linkname VandU64 VandU64
//go:noescape
func VandU64(r *arm.Uint64X1, v0 *arm.Uint64X1, v1 *arm.Uint64X1)

// Bitwise AND (vector). This instruction performs a bitwise AND between the two source SIMD&FP registers, and writes the result to the destination SIMD&FP register.
//
//go:linkname VandU16 VandU16
//go:noescape
func VandU16(r *arm.Uint16X4, v0 *arm.Uint16X4, v1 *arm.Uint16X4)

// Bitwise AND (vector). This instruction performs a bitwise AND between the two source SIMD&FP registers, and writes the result to the destination SIMD&FP register.
//
//go:linkname VandS8 VandS8
//go:noescape
func VandS8(r *arm.Int8X8, v0 *arm.Int8X8, v1 *arm.Int8X8)

// Bitwise AND (vector). This instruction performs a bitwise AND between the two source SIMD&FP registers, and writes the result to the destination SIMD&FP register.
//
//go:linkname VandS32 VandS32
//go:noescape
func VandS32(r *arm.Int32X2, v0 *arm.Int32X2, v1 *arm.Int32X2)

// Bitwise AND (vector). This instruction performs a bitwise AND between the two source SIMD&FP registers, and writes the result to the destination SIMD&FP register.
//
//go:linkname VandS64 VandS64
//go:noescape
func VandS64(r *arm.Int64X1, v0 *arm.Int64X1, v1 *arm.Int64X1)

// Bitwise AND (vector). This instruction performs a bitwise AND between the two source SIMD&FP registers, and writes the result to the destination SIMD&FP register.
//
//go:linkname VandS16 VandS16
//go:noescape
func VandS16(r *arm.Int16X4, v0 *arm.Int16X4, v1 *arm.Int16X4)

// Bitwise bit Clear (vector, register). This instruction performs a bitwise AND between the first source SIMD&FP register and the complement of the second source SIMD&FP register, and writes the result to the destination SIMD&FP register.
//
//go:linkname VbicqU8 VbicqU8
//go:noescape
func VbicqU8(r *arm.Uint8X16, v0 *arm.Uint8X16, v1 *arm.Uint8X16)

// Bitwise bit Clear (vector, register). This instruction performs a bitwise AND between the first source SIMD&FP register and the complement of the second source SIMD&FP register, and writes the result to the destination SIMD&FP register.
//
//go:linkname VbicqU32 VbicqU32
//go:noescape
func VbicqU32(r *arm.Uint32X4, v0 *arm.Uint32X4, v1 *arm.Uint32X4)

// Bitwise bit Clear (vector, register). This instruction performs a bitwise AND between the first source SIMD&FP register and the complement of the second source SIMD&FP register, and writes the result to the destination SIMD&FP register.
//
//go:linkname VbicqU64 VbicqU64
//go:noescape
func VbicqU64(r *arm.Uint64X2, v0 *arm.Uint64X2, v1 *arm.Uint64X2)

// Bitwise bit Clear (vector, register). This instruction performs a bitwise AND between the first source SIMD&FP register and the complement of the second source SIMD&FP register, and writes the result to the destination SIMD&FP register.
//
//go:linkname VbicqU16 VbicqU16
//go:noescape
func VbicqU16(r *arm.Uint16X8, v0 *arm.Uint16X8, v1 *arm.Uint16X8)

// Bitwise bit Clear (vector, register). This instruction performs a bitwise AND between the first source SIMD&FP register and the complement of the second source SIMD&FP register, and writes the result to the destination SIMD&FP register.
//
//go:linkname VbicqS8 VbicqS8
//go:noescape
func VbicqS8(r *arm.Int8X16, v0 *arm.Int8X16, v1 *arm.Int8X16)

// Bitwise bit Clear (vector, register). This instruction performs a bitwise AND between the first source SIMD&FP register and the complement of the second source SIMD&FP register, and writes the result to the destination SIMD&FP register.
//
//go:linkname VbicqS32 VbicqS32
//go:noescape
func VbicqS32(r *arm.Int32X4, v0 *arm.Int32X4, v1 *arm.Int32X4)

// Bitwise bit Clear (vector, register). This instruction performs a bitwise AND between the first source SIMD&FP register and the complement of the second source SIMD&FP register, and writes the result to the destination SIMD&FP register.
//
//go:linkname VbicqS64 VbicqS64
//go:noescape
func VbicqS64(r *arm.Int64X2, v0 *arm.Int64X2, v1 *arm.Int64X2)

// Bitwise bit Clear (vector, register). This instruction performs a bitwise AND between the first source SIMD&FP register and the complement of the second source SIMD&FP register, and writes the result to the destination SIMD&FP register.
//
//go:linkname VbicqS16 VbicqS16
//go:noescape
func VbicqS16(r *arm.Int16X8, v0 *arm.Int16X8, v1 *arm.Int16X8)

// Bitwise bit Clear (vector, register). This instruction performs a bitwise AND between the first source SIMD&FP register and the complement of the second source SIMD&FP register, and writes the result to the destination SIMD&FP register.
//
//go:linkname VbicU8 VbicU8
//go:noescape
func VbicU8(r *arm.Uint8X8, v0 *arm.Uint8X8, v1 *arm.Uint8X8)

// Bitwise bit Clear (vector, register). This instruction performs a bitwise AND between the first source SIMD&FP register and the complement of the second source SIMD&FP register, and writes the result to the destination SIMD&FP register.
//
//go:linkname VbicU32 VbicU32
//go:noescape
func VbicU32(r *arm.Uint32X2, v0 *arm.Uint32X2, v1 *arm.Uint32X2)

// Bitwise bit Clear (vector, register). This instruction performs a bitwise AND between the first source SIMD&FP register and the complement of the second source SIMD&FP register, and writes the result to the destination SIMD&FP register.
//
//go:linkname VbicU64 VbicU64
//go:noescape
func VbicU64(r *arm.Uint64X1, v0 *arm.Uint64X1, v1 *arm.Uint64X1)

// Bitwise bit Clear (vector, register). This instruction performs a bitwise AND between the first source SIMD&FP register and the complement of the second source SIMD&FP register, and writes the result to the destination SIMD&FP register.
//
//go:linkname VbicU16 VbicU16
//go:noescape
func VbicU16(r *arm.Uint16X4, v0 *arm.Uint16X4, v1 *arm.Uint16X4)

// Bitwise bit Clear (vector, register). This instruction performs a bitwise AND between the first source SIMD&FP register and the complement of the second source SIMD&FP register, and writes the result to the destination SIMD&FP register.
//
//go:linkname VbicS8 VbicS8
//go:noescape
func VbicS8(r *arm.Int8X8, v0 *arm.Int8X8, v1 *arm.Int8X8)

// Bitwise bit Clear (vector, register). This instruction performs a bitwise AND between the first source SIMD&FP register and the complement of the second source SIMD&FP register, and writes the result to the destination SIMD&FP register.
//
//go:linkname VbicS32 VbicS32
//go:noescape
func VbicS32(r *arm.Int32X2, v0 *arm.Int32X2, v1 *arm.Int32X2)

// Bitwise bit Clear (vector, register). This instruction performs a bitwise AND between the first source SIMD&FP register and the complement of the second source SIMD&FP register, and writes the result to the destination SIMD&FP register.
//
//go:linkname VbicS64 VbicS64
//go:noescape
func VbicS64(r *arm.Int64X1, v0 *arm.Int64X1, v1 *arm.Int64X1)

// Bitwise bit Clear (vector, register). This instruction performs a bitwise AND between the first source SIMD&FP register and the complement of the second source SIMD&FP register, and writes the result to the destination SIMD&FP register.
//
//go:linkname VbicS16 VbicS16
//go:noescape
func VbicS16(r *arm.Int16X4, v0 *arm.Int16X4, v1 *arm.Int16X4)

// Bitwise Select. This instruction sets each bit in the destination SIMD&FP register to the corresponding bit from the first source SIMD&FP register when the original destination bit was 1, otherwise from the second source SIMD&FP register.
//
//go:linkname VbslP8 VbslP8
//go:noescape
func VbslP8(r *arm.Poly8X8, v0 *arm.Uint8X8, v1 *arm.Poly8X8, v2 *arm.Poly8X8)

// Bitwise Select. This instruction sets each bit in the destination SIMD&FP register to the corresponding bit from the first source SIMD&FP register when the original destination bit was 1, otherwise from the second source SIMD&FP register.
//
//go:linkname VbslP16 VbslP16
//go:noescape
func VbslP16(r *arm.Poly16X4, v0 *arm.Uint16X4, v1 *arm.Poly16X4, v2 *arm.Poly16X4)

// Bitwise Select. This instruction sets each bit in the destination SIMD&FP register to the corresponding bit from the first source SIMD&FP register when the original destination bit was 1, otherwise from the second source SIMD&FP register.
//
//go:linkname VbslqP8 VbslqP8
//go:noescape
func VbslqP8(r *arm.Poly8X16, v0 *arm.Uint8X16, v1 *arm.Poly8X16, v2 *arm.Poly8X16)

// Bitwise Select. This instruction sets each bit in the destination SIMD&FP register to the corresponding bit from the first source SIMD&FP register when the original destination bit was 1, otherwise from the second source SIMD&FP register.
//
//go:linkname VbslqP16 VbslqP16
//go:noescape
func VbslqP16(r *arm.Poly16X8, v0 *arm.Uint16X8, v1 *arm.Poly16X8, v2 *arm.Poly16X8)

// Bitwise Select. This instruction sets each bit in the destination SIMD&FP register to the corresponding bit from the first source SIMD&FP register when the original destination bit was 1, otherwise from the second source SIMD&FP register.
//
//go:linkname VbslqU8 VbslqU8
//go:noescape
func VbslqU8(r *arm.Uint8X16, v0 *arm.Uint8X16, v1 *arm.Uint8X16, v2 *arm.Uint8X16)

// Bitwise Select. This instruction sets each bit in the destination SIMD&FP register to the corresponding bit from the first source SIMD&FP register when the original destination bit was 1, otherwise from the second source SIMD&FP register.
//
//go:linkname VbslqU32 VbslqU32
//go:noescape
func VbslqU32(r *arm.Uint32X4, v0 *arm.Uint32X4, v1 *arm.Uint32X4, v2 *arm.Uint32X4)

// Bitwise Select. This instruction sets each bit in the destination SIMD&FP register to the corresponding bit from the first source SIMD&FP register when the original destination bit was 1, otherwise from the second source SIMD&FP register.
//
//go:linkname VbslqU64 VbslqU64
//go:noescape
func VbslqU64(r *arm.Uint64X2, v0 *arm.Uint64X2, v1 *arm.Uint64X2, v2 *arm.Uint64X2)

// Bitwise Select. This instruction sets each bit in the destination SIMD&FP register to the corresponding bit from the first source SIMD&FP register when the original destination bit was 1, otherwise from the second source SIMD&FP register.
//
//go:linkname VbslqU16 VbslqU16
//go:noescape
func VbslqU16(r *arm.Uint16X8, v0 *arm.Uint16X8, v1 *arm.Uint16X8, v2 *arm.Uint16X8)

// Bitwise Select. This instruction sets each bit in the destination SIMD&FP register to the corresponding bit from the first source SIMD&FP register when the original destination bit was 1, otherwise from the second source SIMD&FP register.
//
//go:linkname VbslqS8 VbslqS8
//go:noescape
func VbslqS8(r *arm.Int8X16, v0 *arm.Uint8X16, v1 *arm.Int8X16, v2 *arm.Int8X16)

// Bitwise Select. This instruction sets each bit in the destination SIMD&FP register to the corresponding bit from the first source SIMD&FP register when the original destination bit was 1, otherwise from the second source SIMD&FP register.
//
//go:linkname VbslqF32 VbslqF32
//go:noescape
func VbslqF32(r *arm.Float32X4, v0 *arm.Uint32X4, v1 *arm.Float32X4, v2 *arm.Float32X4)

// Bitwise Select. This instruction sets each bit in the destination SIMD&FP register to the corresponding bit from the first source SIMD&FP register when the original destination bit was 1, otherwise from the second source SIMD&FP register.
//
//go:linkname VbslqS32 VbslqS32
//go:noescape
func VbslqS32(r *arm.Int32X4, v0 *arm.Uint32X4, v1 *arm.Int32X4, v2 *arm.Int32X4)

// Bitwise Select. This instruction sets each bit in the destination SIMD&FP register to the corresponding bit from the first source SIMD&FP register when the original destination bit was 1, otherwise from the second source SIMD&FP register.
//
//go:linkname VbslqS64 VbslqS64
//go:noescape
func VbslqS64(r *arm.Int64X2, v0 *arm.Uint64X2, v1 *arm.Int64X2, v2 *arm.Int64X2)

// Bitwise Select. This instruction sets each bit in the destination SIMD&FP register to the corresponding bit from the first source SIMD&FP register when the original destination bit was 1, otherwise from the second source SIMD&FP register.
//
//go:linkname VbslqS16 VbslqS16
//go:noescape
func VbslqS16(r *arm.Int16X8, v0 *arm.Uint16X8, v1 *arm.Int16X8, v2 *arm.Int16X8)

// Bitwise Select. This instruction sets each bit in the destination SIMD&FP register to the corresponding bit from the first source SIMD&FP register when the original destination bit was 1, otherwise from the second source SIMD&FP register.
//
//go:linkname VbslU8 VbslU8
//go:noescape
func VbslU8(r *arm.Uint8X8, v0 *arm.Uint8X8, v1 *arm.Uint8X8, v2 *arm.Uint8X8)

// Bitwise Select. This instruction sets each bit in the destination SIMD&FP register to the corresponding bit from the first source SIMD&FP register when the original destination bit was 1, otherwise from the second source SIMD&FP register.
//
//go:linkname VbslU32 VbslU32
//go:noescape
func VbslU32(r *arm.Uint32X2, v0 *arm.Uint32X2, v1 *arm.Uint32X2, v2 *arm.Uint32X2)

// Bitwise Select. This instruction sets each bit in the destination SIMD&FP register to the corresponding bit from the first source SIMD&FP register when the original destination bit was 1, otherwise from the second source SIMD&FP register.
//
//go:linkname VbslU64 VbslU64
//go:noescape
func VbslU64(r *arm.Uint64X1, v0 *arm.Uint64X1, v1 *arm.Uint64X1, v2 *arm.Uint64X1)

// Bitwise Select. This instruction sets each bit in the destination SIMD&FP register to the corresponding bit from the first source SIMD&FP register when the original destination bit was 1, otherwise from the second source SIMD&FP register.
//
//go:linkname VbslU16 VbslU16
//go:noescape
func VbslU16(r *arm.Uint16X4, v0 *arm.Uint16X4, v1 *arm.Uint16X4, v2 *arm.Uint16X4)

// Bitwise Select. This instruction sets each bit in the destination SIMD&FP register to the corresponding bit from the first source SIMD&FP register when the original destination bit was 1, otherwise from the second source SIMD&FP register.
//
//go:linkname VbslS8 VbslS8
//go:noescape
func VbslS8(r *arm.Int8X8, v0 *arm.Uint8X8, v1 *arm.Int8X8, v2 *arm.Int8X8)

// Bitwise Select. This instruction sets each bit in the destination SIMD&FP register to the corresponding bit from the first source SIMD&FP register when the original destination bit was 1, otherwise from the second source SIMD&FP register.
//
//go:linkname VbslF32 VbslF32
//go:noescape
func VbslF32(r *arm.Float32X2, v0 *arm.Uint32X2, v1 *arm.Float32X2, v2 *arm.Float32X2)

// Bitwise Select. This instruction sets each bit in the destination SIMD&FP register to the corresponding bit from the first source SIMD&FP register when the original destination bit was 1, otherwise from the second source SIMD&FP register.
//
//go:linkname VbslS32 VbslS32
//go:noescape
func VbslS32(r *arm.Int32X2, v0 *arm.Uint32X2, v1 *arm.Int32X2, v2 *arm.Int32X2)

// Bitwise Select. This instruction sets each bit in the destination SIMD&FP register to the corresponding bit from the first source SIMD&FP register when the original destination bit was 1, otherwise from the second source SIMD&FP register.
//
//go:linkname VbslS64 VbslS64
//go:noescape
func VbslS64(r *arm.Int64X1, v0 *arm.Uint64X1, v1 *arm.Int64X1, v2 *arm.Int64X1)

// Bitwise Select. This instruction sets each bit in the destination SIMD&FP register to the corresponding bit from the first source SIMD&FP register when the original destination bit was 1, otherwise from the second source SIMD&FP register.
//
//go:linkname VbslS16 VbslS16
//go:noescape
func VbslS16(r *arm.Int16X4, v0 *arm.Uint16X4, v1 *arm.Int16X4, v2 *arm.Int16X4)

// Floating-point Absolute Compare Greater than or Equal (vector). This instruction compares the absolute value of each floating-point value in the first source SIMD&FP register with the absolute value of the corresponding floating-point value in the second source SIMD&FP register and if the first value is greater than or equal to the second value sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VcageqF32 VcageqF32
//go:noescape
func VcageqF32(r *arm.Uint32X4, v0 *arm.Float32X4, v1 *arm.Float32X4)

// Floating-point Absolute Compare Greater than or Equal (vector). This instruction compares the absolute value of each floating-point value in the first source SIMD&FP register with the absolute value of the corresponding floating-point value in the second source SIMD&FP register and if the first value is greater than or equal to the second value sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VcageF32 VcageF32
//go:noescape
func VcageF32(r *arm.Uint32X2, v0 *arm.Float32X2, v1 *arm.Float32X2)

// Floating-point Absolute Compare Greater than (vector). This instruction compares the absolute value of each vector element in the first source SIMD&FP register with the absolute value of the corresponding vector element in the second source SIMD&FP register and if the first value is greater than the second value sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VcagtqF32 VcagtqF32
//go:noescape
func VcagtqF32(r *arm.Uint32X4, v0 *arm.Float32X4, v1 *arm.Float32X4)

// Floating-point Absolute Compare Greater than (vector). This instruction compares the absolute value of each vector element in the first source SIMD&FP register with the absolute value of the corresponding vector element in the second source SIMD&FP register and if the first value is greater than the second value sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VcagtF32 VcagtF32
//go:noescape
func VcagtF32(r *arm.Uint32X2, v0 *arm.Float32X2, v1 *arm.Float32X2)

// Floating-point absolute compare less than or equal
//
//go:linkname VcaleqF32 VcaleqF32
//go:noescape
func VcaleqF32(r *arm.Uint32X4, v0 *arm.Float32X4, v1 *arm.Float32X4)

// Floating-point absolute compare less than or equal
//
//go:linkname VcaleF32 VcaleF32
//go:noescape
func VcaleF32(r *arm.Uint32X2, v0 *arm.Float32X2, v1 *arm.Float32X2)

// Floating-point absolute compare less than
//
//go:linkname VcaltqF32 VcaltqF32
//go:noescape
func VcaltqF32(r *arm.Uint32X4, v0 *arm.Float32X4, v1 *arm.Float32X4)

// Floating-point absolute compare less than
//
//go:linkname VcaltF32 VcaltF32
//go:noescape
func VcaltF32(r *arm.Uint32X2, v0 *arm.Float32X2, v1 *arm.Float32X2)

// Compare bitwise Equal (vector). This instruction compares each vector element from the first source SIMD&FP register with the corresponding vector element from the second source SIMD&FP register, and if the comparison is equal sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VceqP8 VceqP8
//go:noescape
func VceqP8(r *arm.Uint8X8, v0 *arm.Poly8X8, v1 *arm.Poly8X8)

// Compare bitwise Equal (vector). This instruction compares each vector element from the first source SIMD&FP register with the corresponding vector element from the second source SIMD&FP register, and if the comparison is equal sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VceqqP8 VceqqP8
//go:noescape
func VceqqP8(r *arm.Uint8X16, v0 *arm.Poly8X16, v1 *arm.Poly8X16)

// Compare bitwise Equal (vector). This instruction compares each vector element from the first source SIMD&FP register with the corresponding vector element from the second source SIMD&FP register, and if the comparison is equal sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VceqqU8 VceqqU8
//go:noescape
func VceqqU8(r *arm.Uint8X16, v0 *arm.Uint8X16, v1 *arm.Uint8X16)

// Compare bitwise Equal (vector). This instruction compares each vector element from the first source SIMD&FP register with the corresponding vector element from the second source SIMD&FP register, and if the comparison is equal sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VceqqU32 VceqqU32
//go:noescape
func VceqqU32(r *arm.Uint32X4, v0 *arm.Uint32X4, v1 *arm.Uint32X4)

// Compare bitwise Equal (vector). This instruction compares each vector element from the first source SIMD&FP register with the corresponding vector element from the second source SIMD&FP register, and if the comparison is equal sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VceqqU16 VceqqU16
//go:noescape
func VceqqU16(r *arm.Uint16X8, v0 *arm.Uint16X8, v1 *arm.Uint16X8)

// Compare bitwise Equal (vector). This instruction compares each vector element from the first source SIMD&FP register with the corresponding vector element from the second source SIMD&FP register, and if the comparison is equal sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VceqqS8 VceqqS8
//go:noescape
func VceqqS8(r *arm.Uint8X16, v0 *arm.Int8X16, v1 *arm.Int8X16)

// Floating-point Compare Equal (vector). This instruction compares each floating-point value from the first source SIMD&FP register, with the corresponding floating-point value from the second source SIMD&FP register, and if the comparison is equal sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VceqqF32 VceqqF32
//go:noescape
func VceqqF32(r *arm.Uint32X4, v0 *arm.Float32X4, v1 *arm.Float32X4)

// Compare bitwise Equal (vector). This instruction compares each vector element from the first source SIMD&FP register with the corresponding vector element from the second source SIMD&FP register, and if the comparison is equal sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VceqqS32 VceqqS32
//go:noescape
func VceqqS32(r *arm.Uint32X4, v0 *arm.Int32X4, v1 *arm.Int32X4)

// Compare bitwise Equal (vector). This instruction compares each vector element from the first source SIMD&FP register with the corresponding vector element from the second source SIMD&FP register, and if the comparison is equal sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VceqqS16 VceqqS16
//go:noescape
func VceqqS16(r *arm.Uint16X8, v0 *arm.Int16X8, v1 *arm.Int16X8)

// Compare bitwise Equal (vector). This instruction compares each vector element from the first source SIMD&FP register with the corresponding vector element from the second source SIMD&FP register, and if the comparison is equal sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VceqU8 VceqU8
//go:noescape
func VceqU8(r *arm.Uint8X8, v0 *arm.Uint8X8, v1 *arm.Uint8X8)

// Compare bitwise Equal (vector). This instruction compares each vector element from the first source SIMD&FP register with the corresponding vector element from the second source SIMD&FP register, and if the comparison is equal sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VceqU32 VceqU32
//go:noescape
func VceqU32(r *arm.Uint32X2, v0 *arm.Uint32X2, v1 *arm.Uint32X2)

// Compare bitwise Equal (vector). This instruction compares each vector element from the first source SIMD&FP register with the corresponding vector element from the second source SIMD&FP register, and if the comparison is equal sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VceqU16 VceqU16
//go:noescape
func VceqU16(r *arm.Uint16X4, v0 *arm.Uint16X4, v1 *arm.Uint16X4)

// Compare bitwise Equal (vector). This instruction compares each vector element from the first source SIMD&FP register with the corresponding vector element from the second source SIMD&FP register, and if the comparison is equal sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VceqS8 VceqS8
//go:noescape
func VceqS8(r *arm.Uint8X8, v0 *arm.Int8X8, v1 *arm.Int8X8)

// Floating-point Compare Equal (vector). This instruction compares each floating-point value from the first source SIMD&FP register, with the corresponding floating-point value from the second source SIMD&FP register, and if the comparison is equal sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VceqF32 VceqF32
//go:noescape
func VceqF32(r *arm.Uint32X2, v0 *arm.Float32X2, v1 *arm.Float32X2)

// Compare bitwise Equal (vector). This instruction compares each vector element from the first source SIMD&FP register with the corresponding vector element from the second source SIMD&FP register, and if the comparison is equal sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VceqS32 VceqS32
//go:noescape
func VceqS32(r *arm.Uint32X2, v0 *arm.Int32X2, v1 *arm.Int32X2)

// Compare bitwise Equal (vector). This instruction compares each vector element from the first source SIMD&FP register with the corresponding vector element from the second source SIMD&FP register, and if the comparison is equal sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VceqS16 VceqS16
//go:noescape
func VceqS16(r *arm.Uint16X4, v0 *arm.Int16X4, v1 *arm.Int16X4)

// Compare unsigned Higher or Same (vector). This instruction compares each vector element in the first source SIMD&FP register with the corresponding vector element in the second source SIMD&FP register and if the first unsigned integer value is greater than or equal to the second unsigned integer value sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VcgeqU8 VcgeqU8
//go:noescape
func VcgeqU8(r *arm.Uint8X16, v0 *arm.Uint8X16, v1 *arm.Uint8X16)

// Compare unsigned Higher or Same (vector). This instruction compares each vector element in the first source SIMD&FP register with the corresponding vector element in the second source SIMD&FP register and if the first unsigned integer value is greater than or equal to the second unsigned integer value sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VcgeqU32 VcgeqU32
//go:noescape
func VcgeqU32(r *arm.Uint32X4, v0 *arm.Uint32X4, v1 *arm.Uint32X4)

// Compare unsigned Higher or Same (vector). This instruction compares each vector element in the first source SIMD&FP register with the corresponding vector element in the second source SIMD&FP register and if the first unsigned integer value is greater than or equal to the second unsigned integer value sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VcgeqU16 VcgeqU16
//go:noescape
func VcgeqU16(r *arm.Uint16X8, v0 *arm.Uint16X8, v1 *arm.Uint16X8)

// Compare signed Greater than or Equal (vector). This instruction compares each vector element in the first source SIMD&FP register with the corresponding vector element in the second source SIMD&FP register and if the first signed integer value is greater than or equal to the second signed integer value sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VcgeqS8 VcgeqS8
//go:noescape
func VcgeqS8(r *arm.Uint8X16, v0 *arm.Int8X16, v1 *arm.Int8X16)

// Floating-point Compare Greater than or Equal (vector). This instruction reads each floating-point value in the first source SIMD&FP register and if the value is greater than or equal to the corresponding floating-point value in the second source SIMD&FP register sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VcgeqF32 VcgeqF32
//go:noescape
func VcgeqF32(r *arm.Uint32X4, v0 *arm.Float32X4, v1 *arm.Float32X4)

// Compare signed Greater than or Equal (vector). This instruction compares each vector element in the first source SIMD&FP register with the corresponding vector element in the second source SIMD&FP register and if the first signed integer value is greater than or equal to the second signed integer value sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VcgeqS32 VcgeqS32
//go:noescape
func VcgeqS32(r *arm.Uint32X4, v0 *arm.Int32X4, v1 *arm.Int32X4)

// Compare signed Greater than or Equal (vector). This instruction compares each vector element in the first source SIMD&FP register with the corresponding vector element in the second source SIMD&FP register and if the first signed integer value is greater than or equal to the second signed integer value sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VcgeqS16 VcgeqS16
//go:noescape
func VcgeqS16(r *arm.Uint16X8, v0 *arm.Int16X8, v1 *arm.Int16X8)

// Compare unsigned Higher or Same (vector). This instruction compares each vector element in the first source SIMD&FP register with the corresponding vector element in the second source SIMD&FP register and if the first unsigned integer value is greater than or equal to the second unsigned integer value sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VcgeU8 VcgeU8
//go:noescape
func VcgeU8(r *arm.Uint8X8, v0 *arm.Uint8X8, v1 *arm.Uint8X8)

// Compare unsigned Higher or Same (vector). This instruction compares each vector element in the first source SIMD&FP register with the corresponding vector element in the second source SIMD&FP register and if the first unsigned integer value is greater than or equal to the second unsigned integer value sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VcgeU32 VcgeU32
//go:noescape
func VcgeU32(r *arm.Uint32X2, v0 *arm.Uint32X2, v1 *arm.Uint32X2)

// Compare unsigned Higher or Same (vector). This instruction compares each vector element in the first source SIMD&FP register with the corresponding vector element in the second source SIMD&FP register and if the first unsigned integer value is greater than or equal to the second unsigned integer value sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VcgeU16 VcgeU16
//go:noescape
func VcgeU16(r *arm.Uint16X4, v0 *arm.Uint16X4, v1 *arm.Uint16X4)

// Compare signed Greater than or Equal (vector). This instruction compares each vector element in the first source SIMD&FP register with the corresponding vector element in the second source SIMD&FP register and if the first signed integer value is greater than or equal to the second signed integer value sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VcgeS8 VcgeS8
//go:noescape
func VcgeS8(r *arm.Uint8X8, v0 *arm.Int8X8, v1 *arm.Int8X8)

// Floating-point Compare Greater than or Equal (vector). This instruction reads each floating-point value in the first source SIMD&FP register and if the value is greater than or equal to the corresponding floating-point value in the second source SIMD&FP register sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VcgeF32 VcgeF32
//go:noescape
func VcgeF32(r *arm.Uint32X2, v0 *arm.Float32X2, v1 *arm.Float32X2)

// Compare signed Greater than or Equal (vector). This instruction compares each vector element in the first source SIMD&FP register with the corresponding vector element in the second source SIMD&FP register and if the first signed integer value is greater than or equal to the second signed integer value sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VcgeS32 VcgeS32
//go:noescape
func VcgeS32(r *arm.Uint32X2, v0 *arm.Int32X2, v1 *arm.Int32X2)

// Compare signed Greater than or Equal (vector). This instruction compares each vector element in the first source SIMD&FP register with the corresponding vector element in the second source SIMD&FP register and if the first signed integer value is greater than or equal to the second signed integer value sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VcgeS16 VcgeS16
//go:noescape
func VcgeS16(r *arm.Uint16X4, v0 *arm.Int16X4, v1 *arm.Int16X4)

// Compare unsigned Higher (vector). This instruction compares each vector element in the first source SIMD&FP register with the corresponding vector element in the second source SIMD&FP register and if the first unsigned integer value is greater than the second unsigned integer value sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VcgtqU8 VcgtqU8
//go:noescape
func VcgtqU8(r *arm.Uint8X16, v0 *arm.Uint8X16, v1 *arm.Uint8X16)

// Compare unsigned Higher (vector). This instruction compares each vector element in the first source SIMD&FP register with the corresponding vector element in the second source SIMD&FP register and if the first unsigned integer value is greater than the second unsigned integer value sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VcgtqU32 VcgtqU32
//go:noescape
func VcgtqU32(r *arm.Uint32X4, v0 *arm.Uint32X4, v1 *arm.Uint32X4)

// Compare unsigned Higher (vector). This instruction compares each vector element in the first source SIMD&FP register with the corresponding vector element in the second source SIMD&FP register and if the first unsigned integer value is greater than the second unsigned integer value sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VcgtqU16 VcgtqU16
//go:noescape
func VcgtqU16(r *arm.Uint16X8, v0 *arm.Uint16X8, v1 *arm.Uint16X8)

// Compare signed Greater than (vector). This instruction compares each vector element in the first source SIMD&FP register with the corresponding vector element in the second source SIMD&FP register and if the first signed integer value is greater than the second signed integer value sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VcgtqS8 VcgtqS8
//go:noescape
func VcgtqS8(r *arm.Uint8X16, v0 *arm.Int8X16, v1 *arm.Int8X16)

// Floating-point Compare Greater than (vector). This instruction reads each floating-point value in the first source SIMD&FP register and if the value is greater than the corresponding floating-point value in the second source SIMD&FP register sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VcgtqF32 VcgtqF32
//go:noescape
func VcgtqF32(r *arm.Uint32X4, v0 *arm.Float32X4, v1 *arm.Float32X4)

// Compare signed Greater than (vector). This instruction compares each vector element in the first source SIMD&FP register with the corresponding vector element in the second source SIMD&FP register and if the first signed integer value is greater than the second signed integer value sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VcgtqS32 VcgtqS32
//go:noescape
func VcgtqS32(r *arm.Uint32X4, v0 *arm.Int32X4, v1 *arm.Int32X4)

// Compare signed Greater than (vector). This instruction compares each vector element in the first source SIMD&FP register with the corresponding vector element in the second source SIMD&FP register and if the first signed integer value is greater than the second signed integer value sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VcgtqS16 VcgtqS16
//go:noescape
func VcgtqS16(r *arm.Uint16X8, v0 *arm.Int16X8, v1 *arm.Int16X8)

// Compare unsigned Higher (vector). This instruction compares each vector element in the first source SIMD&FP register with the corresponding vector element in the second source SIMD&FP register and if the first unsigned integer value is greater than the second unsigned integer value sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VcgtU8 VcgtU8
//go:noescape
func VcgtU8(r *arm.Uint8X8, v0 *arm.Uint8X8, v1 *arm.Uint8X8)

// Compare unsigned Higher (vector). This instruction compares each vector element in the first source SIMD&FP register with the corresponding vector element in the second source SIMD&FP register and if the first unsigned integer value is greater than the second unsigned integer value sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VcgtU32 VcgtU32
//go:noescape
func VcgtU32(r *arm.Uint32X2, v0 *arm.Uint32X2, v1 *arm.Uint32X2)

// Compare unsigned Higher (vector). This instruction compares each vector element in the first source SIMD&FP register with the corresponding vector element in the second source SIMD&FP register and if the first unsigned integer value is greater than the second unsigned integer value sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VcgtU16 VcgtU16
//go:noescape
func VcgtU16(r *arm.Uint16X4, v0 *arm.Uint16X4, v1 *arm.Uint16X4)

// Compare signed Greater than (vector). This instruction compares each vector element in the first source SIMD&FP register with the corresponding vector element in the second source SIMD&FP register and if the first signed integer value is greater than the second signed integer value sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VcgtS8 VcgtS8
//go:noescape
func VcgtS8(r *arm.Uint8X8, v0 *arm.Int8X8, v1 *arm.Int8X8)

// Floating-point Compare Greater than (vector). This instruction reads each floating-point value in the first source SIMD&FP register and if the value is greater than the corresponding floating-point value in the second source SIMD&FP register sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VcgtF32 VcgtF32
//go:noescape
func VcgtF32(r *arm.Uint32X2, v0 *arm.Float32X2, v1 *arm.Float32X2)

// Compare signed Greater than (vector). This instruction compares each vector element in the first source SIMD&FP register with the corresponding vector element in the second source SIMD&FP register and if the first signed integer value is greater than the second signed integer value sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VcgtS32 VcgtS32
//go:noescape
func VcgtS32(r *arm.Uint32X2, v0 *arm.Int32X2, v1 *arm.Int32X2)

// Compare signed Greater than (vector). This instruction compares each vector element in the first source SIMD&FP register with the corresponding vector element in the second source SIMD&FP register and if the first signed integer value is greater than the second signed integer value sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VcgtS16 VcgtS16
//go:noescape
func VcgtS16(r *arm.Uint16X4, v0 *arm.Int16X4, v1 *arm.Int16X4)

// Compare unsigned less than or equal
//
//go:linkname VcleqU8 VcleqU8
//go:noescape
func VcleqU8(r *arm.Uint8X16, v0 *arm.Uint8X16, v1 *arm.Uint8X16)

// Compare unsigned less than or equal
//
//go:linkname VcleqU32 VcleqU32
//go:noescape
func VcleqU32(r *arm.Uint32X4, v0 *arm.Uint32X4, v1 *arm.Uint32X4)

// Compare unsigned less than or equal
//
//go:linkname VcleqU16 VcleqU16
//go:noescape
func VcleqU16(r *arm.Uint16X8, v0 *arm.Uint16X8, v1 *arm.Uint16X8)

// Compare signed less than or equal
//
//go:linkname VcleqS8 VcleqS8
//go:noescape
func VcleqS8(r *arm.Uint8X16, v0 *arm.Int8X16, v1 *arm.Int8X16)

// Floating-point compare less than or equal
//
//go:linkname VcleqF32 VcleqF32
//go:noescape
func VcleqF32(r *arm.Uint32X4, v0 *arm.Float32X4, v1 *arm.Float32X4)

// Compare signed less than or equal
//
//go:linkname VcleqS32 VcleqS32
//go:noescape
func VcleqS32(r *arm.Uint32X4, v0 *arm.Int32X4, v1 *arm.Int32X4)

// Compare signed less than or equal
//
//go:linkname VcleqS16 VcleqS16
//go:noescape
func VcleqS16(r *arm.Uint16X8, v0 *arm.Int16X8, v1 *arm.Int16X8)

// Compare unsigned less than or equal
//
//go:linkname VcleU8 VcleU8
//go:noescape
func VcleU8(r *arm.Uint8X8, v0 *arm.Uint8X8, v1 *arm.Uint8X8)

// Compare unsigned less than or equal
//
//go:linkname VcleU32 VcleU32
//go:noescape
func VcleU32(r *arm.Uint32X2, v0 *arm.Uint32X2, v1 *arm.Uint32X2)

// Compare unsigned less than or equal
//
//go:linkname VcleU16 VcleU16
//go:noescape
func VcleU16(r *arm.Uint16X4, v0 *arm.Uint16X4, v1 *arm.Uint16X4)

// Compare signed less than or equal
//
//go:linkname VcleS8 VcleS8
//go:noescape
func VcleS8(r *arm.Uint8X8, v0 *arm.Int8X8, v1 *arm.Int8X8)

// Floating-point compare less than or equal
//
//go:linkname VcleF32 VcleF32
//go:noescape
func VcleF32(r *arm.Uint32X2, v0 *arm.Float32X2, v1 *arm.Float32X2)

// Compare signed less than or equal
//
//go:linkname VcleS32 VcleS32
//go:noescape
func VcleS32(r *arm.Uint32X2, v0 *arm.Int32X2, v1 *arm.Int32X2)

// Compare signed less than or equal
//
//go:linkname VcleS16 VcleS16
//go:noescape
func VcleS16(r *arm.Uint16X4, v0 *arm.Int16X4, v1 *arm.Int16X4)

// Count Leading Sign bits (vector). This instruction counts the number of consecutive bits following the most significant bit that are the same as the most significant bit in each vector element in the source SIMD&FP register, places the result into a vector, and writes the vector to the destination SIMD&FP register. The count does not include the most significant bit itself.
//
//go:linkname VclsqU8 VclsqU8
//go:noescape
func VclsqU8(r *arm.Int8X16, v0 *arm.Uint8X16)

// Count Leading Sign bits (vector). This instruction counts the number of consecutive bits following the most significant bit that are the same as the most significant bit in each vector element in the source SIMD&FP register, places the result into a vector, and writes the vector to the destination SIMD&FP register. The count does not include the most significant bit itself.
//
//go:linkname VclsqU32 VclsqU32
//go:noescape
func VclsqU32(r *arm.Int32X4, v0 *arm.Uint32X4)

// Count Leading Sign bits (vector). This instruction counts the number of consecutive bits following the most significant bit that are the same as the most significant bit in each vector element in the source SIMD&FP register, places the result into a vector, and writes the vector to the destination SIMD&FP register. The count does not include the most significant bit itself.
//
//go:linkname VclsqU16 VclsqU16
//go:noescape
func VclsqU16(r *arm.Int16X8, v0 *arm.Uint16X8)

// Count Leading Sign bits (vector). This instruction counts the number of consecutive bits following the most significant bit that are the same as the most significant bit in each vector element in the source SIMD&FP register, places the result into a vector, and writes the vector to the destination SIMD&FP register. The count does not include the most significant bit itself.
//
//go:linkname VclsqS8 VclsqS8
//go:noescape
func VclsqS8(r *arm.Int8X16, v0 *arm.Int8X16)

// Count Leading Sign bits (vector). This instruction counts the number of consecutive bits following the most significant bit that are the same as the most significant bit in each vector element in the source SIMD&FP register, places the result into a vector, and writes the vector to the destination SIMD&FP register. The count does not include the most significant bit itself.
//
//go:linkname VclsqS32 VclsqS32
//go:noescape
func VclsqS32(r *arm.Int32X4, v0 *arm.Int32X4)

// Count Leading Sign bits (vector). This instruction counts the number of consecutive bits following the most significant bit that are the same as the most significant bit in each vector element in the source SIMD&FP register, places the result into a vector, and writes the vector to the destination SIMD&FP register. The count does not include the most significant bit itself.
//
//go:linkname VclsqS16 VclsqS16
//go:noescape
func VclsqS16(r *arm.Int16X8, v0 *arm.Int16X8)

// Count Leading Sign bits (vector). This instruction counts the number of consecutive bits following the most significant bit that are the same as the most significant bit in each vector element in the source SIMD&FP register, places the result into a vector, and writes the vector to the destination SIMD&FP register. The count does not include the most significant bit itself.
//
//go:linkname VclsU8 VclsU8
//go:noescape
func VclsU8(r *arm.Int8X8, v0 *arm.Uint8X8)

// Count Leading Sign bits (vector). This instruction counts the number of consecutive bits following the most significant bit that are the same as the most significant bit in each vector element in the source SIMD&FP register, places the result into a vector, and writes the vector to the destination SIMD&FP register. The count does not include the most significant bit itself.
//
//go:linkname VclsU32 VclsU32
//go:noescape
func VclsU32(r *arm.Int32X2, v0 *arm.Uint32X2)

// Count Leading Sign bits (vector). This instruction counts the number of consecutive bits following the most significant bit that are the same as the most significant bit in each vector element in the source SIMD&FP register, places the result into a vector, and writes the vector to the destination SIMD&FP register. The count does not include the most significant bit itself.
//
//go:linkname VclsU16 VclsU16
//go:noescape
func VclsU16(r *arm.Int16X4, v0 *arm.Uint16X4)

// Count Leading Sign bits (vector). This instruction counts the number of consecutive bits following the most significant bit that are the same as the most significant bit in each vector element in the source SIMD&FP register, places the result into a vector, and writes the vector to the destination SIMD&FP register. The count does not include the most significant bit itself.
//
//go:linkname VclsS8 VclsS8
//go:noescape
func VclsS8(r *arm.Int8X8, v0 *arm.Int8X8)

// Count Leading Sign bits (vector). This instruction counts the number of consecutive bits following the most significant bit that are the same as the most significant bit in each vector element in the source SIMD&FP register, places the result into a vector, and writes the vector to the destination SIMD&FP register. The count does not include the most significant bit itself.
//
//go:linkname VclsS32 VclsS32
//go:noescape
func VclsS32(r *arm.Int32X2, v0 *arm.Int32X2)

// Count Leading Sign bits (vector). This instruction counts the number of consecutive bits following the most significant bit that are the same as the most significant bit in each vector element in the source SIMD&FP register, places the result into a vector, and writes the vector to the destination SIMD&FP register. The count does not include the most significant bit itself.
//
//go:linkname VclsS16 VclsS16
//go:noescape
func VclsS16(r *arm.Int16X4, v0 *arm.Int16X4)

// Compare unsigned less than
//
//go:linkname VcltqU8 VcltqU8
//go:noescape
func VcltqU8(r *arm.Uint8X16, v0 *arm.Uint8X16, v1 *arm.Uint8X16)

// Compare unsigned less than
//
//go:linkname VcltqU32 VcltqU32
//go:noescape
func VcltqU32(r *arm.Uint32X4, v0 *arm.Uint32X4, v1 *arm.Uint32X4)

// Compare unsigned less than
//
//go:linkname VcltqU16 VcltqU16
//go:noescape
func VcltqU16(r *arm.Uint16X8, v0 *arm.Uint16X8, v1 *arm.Uint16X8)

// Compare signed less than
//
//go:linkname VcltqS8 VcltqS8
//go:noescape
func VcltqS8(r *arm.Uint8X16, v0 *arm.Int8X16, v1 *arm.Int8X16)

// Floating-point compare less than
//
//go:linkname VcltqF32 VcltqF32
//go:noescape
func VcltqF32(r *arm.Uint32X4, v0 *arm.Float32X4, v1 *arm.Float32X4)

// Compare signed less than
//
//go:linkname VcltqS32 VcltqS32
//go:noescape
func VcltqS32(r *arm.Uint32X4, v0 *arm.Int32X4, v1 *arm.Int32X4)

// Compare signed less than
//
//go:linkname VcltqS16 VcltqS16
//go:noescape
func VcltqS16(r *arm.Uint16X8, v0 *arm.Int16X8, v1 *arm.Int16X8)

// Compare unsigned less than
//
//go:linkname VcltU8 VcltU8
//go:noescape
func VcltU8(r *arm.Uint8X8, v0 *arm.Uint8X8, v1 *arm.Uint8X8)

// Compare unsigned less than
//
//go:linkname VcltU32 VcltU32
//go:noescape
func VcltU32(r *arm.Uint32X2, v0 *arm.Uint32X2, v1 *arm.Uint32X2)

// Compare unsigned less than
//
//go:linkname VcltU16 VcltU16
//go:noescape
func VcltU16(r *arm.Uint16X4, v0 *arm.Uint16X4, v1 *arm.Uint16X4)

// Compare signed less than
//
//go:linkname VcltS8 VcltS8
//go:noescape
func VcltS8(r *arm.Uint8X8, v0 *arm.Int8X8, v1 *arm.Int8X8)

// Floating-point compare less than
//
//go:linkname VcltF32 VcltF32
//go:noescape
func VcltF32(r *arm.Uint32X2, v0 *arm.Float32X2, v1 *arm.Float32X2)

// Compare signed less than
//
//go:linkname VcltS32 VcltS32
//go:noescape
func VcltS32(r *arm.Uint32X2, v0 *arm.Int32X2, v1 *arm.Int32X2)

// Compare signed less than
//
//go:linkname VcltS16 VcltS16
//go:noescape
func VcltS16(r *arm.Uint16X4, v0 *arm.Int16X4, v1 *arm.Int16X4)

// Count Leading Zero bits (vector). This instruction counts the number of consecutive zeros, starting from the most significant bit, in each vector element in the source SIMD&FP register, places the result into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VclzqU8 VclzqU8
//go:noescape
func VclzqU8(r *arm.Uint8X16, v0 *arm.Uint8X16)

// Count Leading Zero bits (vector). This instruction counts the number of consecutive zeros, starting from the most significant bit, in each vector element in the source SIMD&FP register, places the result into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VclzqU32 VclzqU32
//go:noescape
func VclzqU32(r *arm.Uint32X4, v0 *arm.Uint32X4)

// Count Leading Zero bits (vector). This instruction counts the number of consecutive zeros, starting from the most significant bit, in each vector element in the source SIMD&FP register, places the result into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VclzqU16 VclzqU16
//go:noescape
func VclzqU16(r *arm.Uint16X8, v0 *arm.Uint16X8)

// Count Leading Zero bits (vector). This instruction counts the number of consecutive zeros, starting from the most significant bit, in each vector element in the source SIMD&FP register, places the result into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VclzqS8 VclzqS8
//go:noescape
func VclzqS8(r *arm.Int8X16, v0 *arm.Int8X16)

// Count Leading Zero bits (vector). This instruction counts the number of consecutive zeros, starting from the most significant bit, in each vector element in the source SIMD&FP register, places the result into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VclzqS32 VclzqS32
//go:noescape
func VclzqS32(r *arm.Int32X4, v0 *arm.Int32X4)

// Count Leading Zero bits (vector). This instruction counts the number of consecutive zeros, starting from the most significant bit, in each vector element in the source SIMD&FP register, places the result into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VclzqS16 VclzqS16
//go:noescape
func VclzqS16(r *arm.Int16X8, v0 *arm.Int16X8)

// Count Leading Zero bits (vector). This instruction counts the number of consecutive zeros, starting from the most significant bit, in each vector element in the source SIMD&FP register, places the result into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VclzU8 VclzU8
//go:noescape
func VclzU8(r *arm.Uint8X8, v0 *arm.Uint8X8)

// Count Leading Zero bits (vector). This instruction counts the number of consecutive zeros, starting from the most significant bit, in each vector element in the source SIMD&FP register, places the result into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VclzU32 VclzU32
//go:noescape
func VclzU32(r *arm.Uint32X2, v0 *arm.Uint32X2)

// Count Leading Zero bits (vector). This instruction counts the number of consecutive zeros, starting from the most significant bit, in each vector element in the source SIMD&FP register, places the result into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VclzU16 VclzU16
//go:noescape
func VclzU16(r *arm.Uint16X4, v0 *arm.Uint16X4)

// Count Leading Zero bits (vector). This instruction counts the number of consecutive zeros, starting from the most significant bit, in each vector element in the source SIMD&FP register, places the result into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VclzS8 VclzS8
//go:noescape
func VclzS8(r *arm.Int8X8, v0 *arm.Int8X8)

// Count Leading Zero bits (vector). This instruction counts the number of consecutive zeros, starting from the most significant bit, in each vector element in the source SIMD&FP register, places the result into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VclzS32 VclzS32
//go:noescape
func VclzS32(r *arm.Int32X2, v0 *arm.Int32X2)

// Count Leading Zero bits (vector). This instruction counts the number of consecutive zeros, starting from the most significant bit, in each vector element in the source SIMD&FP register, places the result into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VclzS16 VclzS16
//go:noescape
func VclzS16(r *arm.Int16X4, v0 *arm.Int16X4)

// Population Count per byte. This instruction counts the number of bits that have a value of one in each vector element in the source SIMD&FP register, places the result into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VcntP8 VcntP8
//go:noescape
func VcntP8(r *arm.Poly8X8, v0 *arm.Poly8X8)

// Population Count per byte. This instruction counts the number of bits that have a value of one in each vector element in the source SIMD&FP register, places the result into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VcntqP8 VcntqP8
//go:noescape
func VcntqP8(r *arm.Poly8X16, v0 *arm.Poly8X16)

// Population Count per byte. This instruction counts the number of bits that have a value of one in each vector element in the source SIMD&FP register, places the result into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VcntqU8 VcntqU8
//go:noescape
func VcntqU8(r *arm.Uint8X16, v0 *arm.Uint8X16)

// Population Count per byte. This instruction counts the number of bits that have a value of one in each vector element in the source SIMD&FP register, places the result into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VcntqS8 VcntqS8
//go:noescape
func VcntqS8(r *arm.Int8X16, v0 *arm.Int8X16)

// Population Count per byte. This instruction counts the number of bits that have a value of one in each vector element in the source SIMD&FP register, places the result into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VcntU8 VcntU8
//go:noescape
func VcntU8(r *arm.Uint8X8, v0 *arm.Uint8X8)

// Population Count per byte. This instruction counts the number of bits that have a value of one in each vector element in the source SIMD&FP register, places the result into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VcntS8 VcntS8
//go:noescape
func VcntS8(r *arm.Int8X8, v0 *arm.Int8X8)

// Join two smaller vectors into a single larger vector
//
//go:linkname VcombineP8 VcombineP8
//go:noescape
func VcombineP8(r *arm.Poly8X16, v0 *arm.Poly8X8, v1 *arm.Poly8X8)

// Join two smaller vectors into a single larger vector
//
//go:linkname VcombineP16 VcombineP16
//go:noescape
func VcombineP16(r *arm.Poly16X8, v0 *arm.Poly16X4, v1 *arm.Poly16X4)

// Join two smaller vectors into a single larger vector
//
//go:linkname VcombineU8 VcombineU8
//go:noescape
func VcombineU8(r *arm.Uint8X16, v0 *arm.Uint8X8, v1 *arm.Uint8X8)

// Join two smaller vectors into a single larger vector
//
//go:linkname VcombineU32 VcombineU32
//go:noescape
func VcombineU32(r *arm.Uint32X4, v0 *arm.Uint32X2, v1 *arm.Uint32X2)

// Join two smaller vectors into a single larger vector
//
//go:linkname VcombineU64 VcombineU64
//go:noescape
func VcombineU64(r *arm.Uint64X2, v0 *arm.Uint64X1, v1 *arm.Uint64X1)

// Join two smaller vectors into a single larger vector
//
//go:linkname VcombineU16 VcombineU16
//go:noescape
func VcombineU16(r *arm.Uint16X8, v0 *arm.Uint16X4, v1 *arm.Uint16X4)

// Join two smaller vectors into a single larger vector
//
//go:linkname VcombineS8 VcombineS8
//go:noescape
func VcombineS8(r *arm.Int8X16, v0 *arm.Int8X8, v1 *arm.Int8X8)

// Join two smaller vectors into a single larger vector
//
//go:linkname VcombineF32 VcombineF32
//go:noescape
func VcombineF32(r *arm.Float32X4, v0 *arm.Float32X2, v1 *arm.Float32X2)

// Join two smaller vectors into a single larger vector
//
//go:linkname VcombineS32 VcombineS32
//go:noescape
func VcombineS32(r *arm.Int32X4, v0 *arm.Int32X2, v1 *arm.Int32X2)

// Join two smaller vectors into a single larger vector
//
//go:linkname VcombineS64 VcombineS64
//go:noescape
func VcombineS64(r *arm.Int64X2, v0 *arm.Int64X1, v1 *arm.Int64X1)

// Join two smaller vectors into a single larger vector
//
//go:linkname VcombineS16 VcombineS16
//go:noescape
func VcombineS16(r *arm.Int16X8, v0 *arm.Int16X4, v1 *arm.Int16X4)

// Unsigned fixed-point Convert to Floating-point (vector). This instruction converts each element in a vector from fixed-point to floating-point using the rounding mode that is specified by the FPCR, and writes the result to the SIMD&FP destination register.
//
//go:linkname VcvtqF32U32 VcvtqF32U32
//go:noescape
func VcvtqF32U32(r *arm.Float32X4, v0 *arm.Uint32X4)

// Signed fixed-point Convert to Floating-point (vector). This instruction converts each element in a vector from fixed-point to floating-point using the rounding mode that is specified by the FPCR, and writes the result to the SIMD&FP destination register.
//
//go:linkname VcvtqF32S32 VcvtqF32S32
//go:noescape
func VcvtqF32S32(r *arm.Float32X4, v0 *arm.Int32X4)

// Unsigned fixed-point Convert to Floating-point (vector). This instruction converts each element in a vector from fixed-point to floating-point using the rounding mode that is specified by the FPCR, and writes the result to the SIMD&FP destination register.
//
//go:linkname VcvtF32U32 VcvtF32U32
//go:noescape
func VcvtF32U32(r *arm.Float32X2, v0 *arm.Uint32X2)

// Signed fixed-point Convert to Floating-point (vector). This instruction converts each element in a vector from fixed-point to floating-point using the rounding mode that is specified by the FPCR, and writes the result to the SIMD&FP destination register.
//
//go:linkname VcvtF32S32 VcvtF32S32
//go:noescape
func VcvtF32S32(r *arm.Float32X2, v0 *arm.Int32X2)

// Floating-point Convert to Signed fixed-point, rounding toward Zero (vector). This instruction converts a scalar or each element in a vector from floating-point to fixed-point signed integer using the Round towards Zero rounding mode, and writes the result to the SIMD&FP destination register.
//
//go:linkname VcvtqS32F32 VcvtqS32F32
//go:noescape
func VcvtqS32F32(r *arm.Int32X4, v0 *arm.Float32X4)

// Floating-point Convert to Signed fixed-point, rounding toward Zero (vector). This instruction converts a scalar or each element in a vector from floating-point to fixed-point signed integer using the Round towards Zero rounding mode, and writes the result to the SIMD&FP destination register.
//
//go:linkname VcvtS32F32 VcvtS32F32
//go:noescape
func VcvtS32F32(r *arm.Int32X2, v0 *arm.Float32X2)

// Floating-point Convert to Unsigned fixed-point, rounding toward Zero (vector). This instruction converts a scalar or each element in a vector from floating-point to fixed-point unsigned integer using the Round towards Zero rounding mode, and writes the result to the general-purpose destination register.
//
//go:linkname VcvtqU32F32 VcvtqU32F32
//go:noescape
func VcvtqU32F32(r *arm.Uint32X4, v0 *arm.Float32X4)

// Floating-point Convert to Unsigned fixed-point, rounding toward Zero (vector). This instruction converts a scalar or each element in a vector from floating-point to fixed-point unsigned integer using the Round towards Zero rounding mode, and writes the result to the general-purpose destination register.
//
//go:linkname VcvtU32F32 VcvtU32F32
//go:noescape
func VcvtU32F32(r *arm.Uint32X2, v0 *arm.Float32X2)

// Duplicate vector element to vector or scalar. This instruction duplicates the vector element at the specified element index in the source SIMD&FP register into a scalar or each element in a vector, and writes the result to the destination SIMD&FP register.
//
//go:linkname VdupNP8 VdupNP8
//go:noescape
func VdupNP8(r *arm.Poly8X8, v0 *arm.Poly8)

// Duplicate vector element to vector or scalar. This instruction duplicates the vector element at the specified element index in the source SIMD&FP register into a scalar or each element in a vector, and writes the result to the destination SIMD&FP register.
//
//go:linkname VdupNP16 VdupNP16
//go:noescape
func VdupNP16(r *arm.Poly16X4, v0 *arm.Poly16)

// Duplicate vector element to vector or scalar. This instruction duplicates the vector element at the specified element index in the source SIMD&FP register into a scalar or each element in a vector, and writes the result to the destination SIMD&FP register.
//
//go:linkname VdupqNP8 VdupqNP8
//go:noescape
func VdupqNP8(r *arm.Poly8X16, v0 *arm.Poly8)

// Duplicate vector element to vector or scalar. This instruction duplicates the vector element at the specified element index in the source SIMD&FP register into a scalar or each element in a vector, and writes the result to the destination SIMD&FP register.
//
//go:linkname VdupqNP16 VdupqNP16
//go:noescape
func VdupqNP16(r *arm.Poly16X8, v0 *arm.Poly16)

// Duplicate vector element to vector or scalar. This instruction duplicates the vector element at the specified element index in the source SIMD&FP register into a scalar or each element in a vector, and writes the result to the destination SIMD&FP register.
//
//go:linkname VdupqNU8 VdupqNU8
//go:noescape
func VdupqNU8(r *arm.Uint8X16, v0 *arm.Uint8)

// Duplicate vector element to vector or scalar. This instruction duplicates the vector element at the specified element index in the source SIMD&FP register into a scalar or each element in a vector, and writes the result to the destination SIMD&FP register.
//
//go:linkname VdupqNU32 VdupqNU32
//go:noescape
func VdupqNU32(r *arm.Uint32X4, v0 *arm.Uint32)

// Duplicate vector element to vector or scalar. This instruction duplicates the vector element at the specified element index in the source SIMD&FP register into a scalar or each element in a vector, and writes the result to the destination SIMD&FP register.
//
//go:linkname VdupqNU64 VdupqNU64
//go:noescape
func VdupqNU64(r *arm.Uint64X2, v0 *arm.Uint64)

// Duplicate vector element to vector or scalar. This instruction duplicates the vector element at the specified element index in the source SIMD&FP register into a scalar or each element in a vector, and writes the result to the destination SIMD&FP register.
//
//go:linkname VdupqNU16 VdupqNU16
//go:noescape
func VdupqNU16(r *arm.Uint16X8, v0 *arm.Uint16)

// Duplicate vector element to vector or scalar. This instruction duplicates the vector element at the specified element index in the source SIMD&FP register into a scalar or each element in a vector, and writes the result to the destination SIMD&FP register.
//
//go:linkname VdupqNS8 VdupqNS8
//go:noescape
func VdupqNS8(r *arm.Int8X16, v0 *arm.Int8)

// Duplicate vector element to vector or scalar. This instruction duplicates the vector element at the specified element index in the source SIMD&FP register into a scalar or each element in a vector, and writes the result to the destination SIMD&FP register.
//
//go:linkname VdupqNF32 VdupqNF32
//go:noescape
func VdupqNF32(r *arm.Float32X4, v0 *arm.Float32)

// Duplicate vector element to vector or scalar. This instruction duplicates the vector element at the specified element index in the source SIMD&FP register into a scalar or each element in a vector, and writes the result to the destination SIMD&FP register.
//
//go:linkname VdupqNS32 VdupqNS32
//go:noescape
func VdupqNS32(r *arm.Int32X4, v0 *arm.Int32)

// Duplicate vector element to vector or scalar. This instruction duplicates the vector element at the specified element index in the source SIMD&FP register into a scalar or each element in a vector, and writes the result to the destination SIMD&FP register.
//
//go:linkname VdupqNS64 VdupqNS64
//go:noescape
func VdupqNS64(r *arm.Int64X2, v0 *arm.Int64)

// Duplicate vector element to vector or scalar. This instruction duplicates the vector element at the specified element index in the source SIMD&FP register into a scalar or each element in a vector, and writes the result to the destination SIMD&FP register.
//
//go:linkname VdupqNS16 VdupqNS16
//go:noescape
func VdupqNS16(r *arm.Int16X8, v0 *arm.Int16)

// Duplicate vector element to vector or scalar. This instruction duplicates the vector element at the specified element index in the source SIMD&FP register into a scalar or each element in a vector, and writes the result to the destination SIMD&FP register.
//
//go:linkname VdupNU8 VdupNU8
//go:noescape
func VdupNU8(r *arm.Uint8X8, v0 *arm.Uint8)

// Duplicate vector element to vector or scalar. This instruction duplicates the vector element at the specified element index in the source SIMD&FP register into a scalar or each element in a vector, and writes the result to the destination SIMD&FP register.
//
//go:linkname VdupNU32 VdupNU32
//go:noescape
func VdupNU32(r *arm.Uint32X2, v0 *arm.Uint32)

// Insert vector element from another vector element. This instruction copies the vector element of the source SIMD&FP register to the specified vector element of the destination SIMD&FP register.
//
//go:linkname VdupNU64 VdupNU64
//go:noescape
func VdupNU64(r *arm.Uint64X1, v0 *arm.Uint64)

// Duplicate vector element to vector or scalar. This instruction duplicates the vector element at the specified element index in the source SIMD&FP register into a scalar or each element in a vector, and writes the result to the destination SIMD&FP register.
//
//go:linkname VdupNU16 VdupNU16
//go:noescape
func VdupNU16(r *arm.Uint16X4, v0 *arm.Uint16)

// Duplicate vector element to vector or scalar. This instruction duplicates the vector element at the specified element index in the source SIMD&FP register into a scalar or each element in a vector, and writes the result to the destination SIMD&FP register.
//
//go:linkname VdupNS8 VdupNS8
//go:noescape
func VdupNS8(r *arm.Int8X8, v0 *arm.Int8)

// Duplicate vector element to vector or scalar. This instruction duplicates the vector element at the specified element index in the source SIMD&FP register into a scalar or each element in a vector, and writes the result to the destination SIMD&FP register.
//
//go:linkname VdupNF32 VdupNF32
//go:noescape
func VdupNF32(r *arm.Float32X2, v0 *arm.Float32)

// Duplicate vector element to vector or scalar. This instruction duplicates the vector element at the specified element index in the source SIMD&FP register into a scalar or each element in a vector, and writes the result to the destination SIMD&FP register.
//
//go:linkname VdupNS32 VdupNS32
//go:noescape
func VdupNS32(r *arm.Int32X2, v0 *arm.Int32)

// Insert vector element from another vector element. This instruction copies the vector element of the source SIMD&FP register to the specified vector element of the destination SIMD&FP register.
//
//go:linkname VdupNS64 VdupNS64
//go:noescape
func VdupNS64(r *arm.Int64X1, v0 *arm.Int64)

// Duplicate vector element to vector or scalar. This instruction duplicates the vector element at the specified element index in the source SIMD&FP register into a scalar or each element in a vector, and writes the result to the destination SIMD&FP register.
//
//go:linkname VdupNS16 VdupNS16
//go:noescape
func VdupNS16(r *arm.Int16X4, v0 *arm.Int16)

// Bitwise Exclusive OR (vector). This instruction performs a bitwise Exclusive OR operation between the two source SIMD&FP registers, and places the result in the destination SIMD&FP register.
//
//go:linkname VeorqU8 VeorqU8
//go:noescape
func VeorqU8(r *arm.Uint8X16, v0 *arm.Uint8X16, v1 *arm.Uint8X16)

// Bitwise Exclusive OR (vector). This instruction performs a bitwise Exclusive OR operation between the two source SIMD&FP registers, and places the result in the destination SIMD&FP register.
//
//go:linkname VeorqU32 VeorqU32
//go:noescape
func VeorqU32(r *arm.Uint32X4, v0 *arm.Uint32X4, v1 *arm.Uint32X4)

// Bitwise Exclusive OR (vector). This instruction performs a bitwise Exclusive OR operation between the two source SIMD&FP registers, and places the result in the destination SIMD&FP register.
//
//go:linkname VeorqU64 VeorqU64
//go:noescape
func VeorqU64(r *arm.Uint64X2, v0 *arm.Uint64X2, v1 *arm.Uint64X2)

// Bitwise Exclusive OR (vector). This instruction performs a bitwise Exclusive OR operation between the two source SIMD&FP registers, and places the result in the destination SIMD&FP register.
//
//go:linkname VeorqU16 VeorqU16
//go:noescape
func VeorqU16(r *arm.Uint16X8, v0 *arm.Uint16X8, v1 *arm.Uint16X8)

// Bitwise Exclusive OR (vector). This instruction performs a bitwise Exclusive OR operation between the two source SIMD&FP registers, and places the result in the destination SIMD&FP register.
//
//go:linkname VeorqS8 VeorqS8
//go:noescape
func VeorqS8(r *arm.Int8X16, v0 *arm.Int8X16, v1 *arm.Int8X16)

// Bitwise Exclusive OR (vector). This instruction performs a bitwise Exclusive OR operation between the two source SIMD&FP registers, and places the result in the destination SIMD&FP register.
//
//go:linkname VeorqS32 VeorqS32
//go:noescape
func VeorqS32(r *arm.Int32X4, v0 *arm.Int32X4, v1 *arm.Int32X4)

// Bitwise Exclusive OR (vector). This instruction performs a bitwise Exclusive OR operation between the two source SIMD&FP registers, and places the result in the destination SIMD&FP register.
//
//go:linkname VeorqS64 VeorqS64
//go:noescape
func VeorqS64(r *arm.Int64X2, v0 *arm.Int64X2, v1 *arm.Int64X2)

// Bitwise Exclusive OR (vector). This instruction performs a bitwise Exclusive OR operation between the two source SIMD&FP registers, and places the result in the destination SIMD&FP register.
//
//go:linkname VeorqS16 VeorqS16
//go:noescape
func VeorqS16(r *arm.Int16X8, v0 *arm.Int16X8, v1 *arm.Int16X8)

// Bitwise Exclusive OR (vector). This instruction performs a bitwise Exclusive OR operation between the two source SIMD&FP registers, and places the result in the destination SIMD&FP register.
//
//go:linkname VeorU8 VeorU8
//go:noescape
func VeorU8(r *arm.Uint8X8, v0 *arm.Uint8X8, v1 *arm.Uint8X8)

// Bitwise Exclusive OR (vector). This instruction performs a bitwise Exclusive OR operation between the two source SIMD&FP registers, and places the result in the destination SIMD&FP register.
//
//go:linkname VeorU32 VeorU32
//go:noescape
func VeorU32(r *arm.Uint32X2, v0 *arm.Uint32X2, v1 *arm.Uint32X2)

// Bitwise Exclusive OR (vector). This instruction performs a bitwise Exclusive OR operation between the two source SIMD&FP registers, and places the result in the destination SIMD&FP register.
//
//go:linkname VeorU64 VeorU64
//go:noescape
func VeorU64(r *arm.Uint64X1, v0 *arm.Uint64X1, v1 *arm.Uint64X1)

// Bitwise Exclusive OR (vector). This instruction performs a bitwise Exclusive OR operation between the two source SIMD&FP registers, and places the result in the destination SIMD&FP register.
//
//go:linkname VeorU16 VeorU16
//go:noescape
func VeorU16(r *arm.Uint16X4, v0 *arm.Uint16X4, v1 *arm.Uint16X4)

// Bitwise Exclusive OR (vector). This instruction performs a bitwise Exclusive OR operation between the two source SIMD&FP registers, and places the result in the destination SIMD&FP register.
//
//go:linkname VeorS8 VeorS8
//go:noescape
func VeorS8(r *arm.Int8X8, v0 *arm.Int8X8, v1 *arm.Int8X8)

// Bitwise Exclusive OR (vector). This instruction performs a bitwise Exclusive OR operation between the two source SIMD&FP registers, and places the result in the destination SIMD&FP register.
//
//go:linkname VeorS32 VeorS32
//go:noescape
func VeorS32(r *arm.Int32X2, v0 *arm.Int32X2, v1 *arm.Int32X2)

// Bitwise Exclusive OR (vector). This instruction performs a bitwise Exclusive OR operation between the two source SIMD&FP registers, and places the result in the destination SIMD&FP register.
//
//go:linkname VeorS64 VeorS64
//go:noescape
func VeorS64(r *arm.Int64X1, v0 *arm.Int64X1, v1 *arm.Int64X1)

// Bitwise Exclusive OR (vector). This instruction performs a bitwise Exclusive OR operation between the two source SIMD&FP registers, and places the result in the destination SIMD&FP register.
//
//go:linkname VeorS16 VeorS16
//go:noescape
func VeorS16(r *arm.Int16X4, v0 *arm.Int16X4, v1 *arm.Int16X4)

// Duplicate vector element to vector or scalar. This instruction duplicates the vector element at the specified element index in the source SIMD&FP register into a scalar or each element in a vector, and writes the result to the destination SIMD&FP register.
//
//go:linkname VgetHighP8 VgetHighP8
//go:noescape
func VgetHighP8(r *arm.Poly8X8, v0 *arm.Poly8X16)

// Duplicate vector element to vector or scalar. This instruction duplicates the vector element at the specified element index in the source SIMD&FP register into a scalar or each element in a vector, and writes the result to the destination SIMD&FP register.
//
//go:linkname VgetHighP16 VgetHighP16
//go:noescape
func VgetHighP16(r *arm.Poly16X4, v0 *arm.Poly16X8)

// Duplicate vector element to vector or scalar. This instruction duplicates the vector element at the specified element index in the source SIMD&FP register into a scalar or each element in a vector, and writes the result to the destination SIMD&FP register.
//
//go:linkname VgetHighU8 VgetHighU8
//go:noescape
func VgetHighU8(r *arm.Uint8X8, v0 *arm.Uint8X16)

// Duplicate vector element to vector or scalar. This instruction duplicates the vector element at the specified element index in the source SIMD&FP register into a scalar or each element in a vector, and writes the result to the destination SIMD&FP register.
//
//go:linkname VgetHighU32 VgetHighU32
//go:noescape
func VgetHighU32(r *arm.Uint32X2, v0 *arm.Uint32X4)

// Duplicate vector element to vector or scalar. This instruction duplicates the vector element at the specified element index in the source SIMD&FP register into a scalar or each element in a vector, and writes the result to the destination SIMD&FP register.
//
//go:linkname VgetHighU64 VgetHighU64
//go:noescape
func VgetHighU64(r *arm.Uint64X1, v0 *arm.Uint64X2)

// Duplicate vector element to vector or scalar. This instruction duplicates the vector element at the specified element index in the source SIMD&FP register into a scalar or each element in a vector, and writes the result to the destination SIMD&FP register.
//
//go:linkname VgetHighU16 VgetHighU16
//go:noescape
func VgetHighU16(r *arm.Uint16X4, v0 *arm.Uint16X8)

// Duplicate vector element to vector or scalar. This instruction duplicates the vector element at the specified element index in the source SIMD&FP register into a scalar or each element in a vector, and writes the result to the destination SIMD&FP register.
//
//go:linkname VgetHighS8 VgetHighS8
//go:noescape
func VgetHighS8(r *arm.Int8X8, v0 *arm.Int8X16)

// Duplicate vector element to vector or scalar. This instruction duplicates the vector element at the specified element index in the source SIMD&FP register into a scalar or each element in a vector, and writes the result to the destination SIMD&FP register.
//
//go:linkname VgetHighF32 VgetHighF32
//go:noescape
func VgetHighF32(r *arm.Float32X2, v0 *arm.Float32X4)

// Duplicate vector element to vector or scalar. This instruction duplicates the vector element at the specified element index in the source SIMD&FP register into a scalar or each element in a vector, and writes the result to the destination SIMD&FP register.
//
//go:linkname VgetHighS32 VgetHighS32
//go:noescape
func VgetHighS32(r *arm.Int32X2, v0 *arm.Int32X4)

// Duplicate vector element to vector or scalar. This instruction duplicates the vector element at the specified element index in the source SIMD&FP register into a scalar or each element in a vector, and writes the result to the destination SIMD&FP register.
//
//go:linkname VgetHighS64 VgetHighS64
//go:noescape
func VgetHighS64(r *arm.Int64X1, v0 *arm.Int64X2)

// Duplicate vector element to vector or scalar. This instruction duplicates the vector element at the specified element index in the source SIMD&FP register into a scalar or each element in a vector, and writes the result to the destination SIMD&FP register.
//
//go:linkname VgetHighS16 VgetHighS16
//go:noescape
func VgetHighS16(r *arm.Int16X4, v0 *arm.Int16X8)

// Duplicate vector element to vector or scalar. This instruction duplicates the vector element at the specified element index in the source SIMD&FP register into a scalar or each element in a vector, and writes the result to the destination SIMD&FP register.
//
//go:linkname VgetLowP8 VgetLowP8
//go:noescape
func VgetLowP8(r *arm.Poly8X8, v0 *arm.Poly8X16)

// Duplicate vector element to vector or scalar. This instruction duplicates the vector element at the specified element index in the source SIMD&FP register into a scalar or each element in a vector, and writes the result to the destination SIMD&FP register.
//
//go:linkname VgetLowP16 VgetLowP16
//go:noescape
func VgetLowP16(r *arm.Poly16X4, v0 *arm.Poly16X8)

// Duplicate vector element to vector or scalar. This instruction duplicates the vector element at the specified element index in the source SIMD&FP register into a scalar or each element in a vector, and writes the result to the destination SIMD&FP register.
//
//go:linkname VgetLowU8 VgetLowU8
//go:noescape
func VgetLowU8(r *arm.Uint8X8, v0 *arm.Uint8X16)

// Duplicate vector element to vector or scalar. This instruction duplicates the vector element at the specified element index in the source SIMD&FP register into a scalar or each element in a vector, and writes the result to the destination SIMD&FP register.
//
//go:linkname VgetLowU32 VgetLowU32
//go:noescape
func VgetLowU32(r *arm.Uint32X2, v0 *arm.Uint32X4)

// Duplicate vector element to vector or scalar. This instruction duplicates the vector element at the specified element index in the source SIMD&FP register into a scalar or each element in a vector, and writes the result to the destination SIMD&FP register.
//
//go:linkname VgetLowU64 VgetLowU64
//go:noescape
func VgetLowU64(r *arm.Uint64X1, v0 *arm.Uint64X2)

// Duplicate vector element to vector or scalar. This instruction duplicates the vector element at the specified element index in the source SIMD&FP register into a scalar or each element in a vector, and writes the result to the destination SIMD&FP register.
//
//go:linkname VgetLowU16 VgetLowU16
//go:noescape
func VgetLowU16(r *arm.Uint16X4, v0 *arm.Uint16X8)

// Duplicate vector element to vector or scalar. This instruction duplicates the vector element at the specified element index in the source SIMD&FP register into a scalar or each element in a vector, and writes the result to the destination SIMD&FP register.
//
//go:linkname VgetLowS8 VgetLowS8
//go:noescape
func VgetLowS8(r *arm.Int8X8, v0 *arm.Int8X16)

// Duplicate vector element to vector or scalar. This instruction duplicates the vector element at the specified element index in the source SIMD&FP register into a scalar or each element in a vector, and writes the result to the destination SIMD&FP register.
//
//go:linkname VgetLowF32 VgetLowF32
//go:noescape
func VgetLowF32(r *arm.Float32X2, v0 *arm.Float32X4)

// Duplicate vector element to vector or scalar. This instruction duplicates the vector element at the specified element index in the source SIMD&FP register into a scalar or each element in a vector, and writes the result to the destination SIMD&FP register.
//
//go:linkname VgetLowS32 VgetLowS32
//go:noescape
func VgetLowS32(r *arm.Int32X2, v0 *arm.Int32X4)

// Duplicate vector element to vector or scalar. This instruction duplicates the vector element at the specified element index in the source SIMD&FP register into a scalar or each element in a vector, and writes the result to the destination SIMD&FP register.
//
//go:linkname VgetLowS64 VgetLowS64
//go:noescape
func VgetLowS64(r *arm.Int64X1, v0 *arm.Int64X2)

// Duplicate vector element to vector or scalar. This instruction duplicates the vector element at the specified element index in the source SIMD&FP register into a scalar or each element in a vector, and writes the result to the destination SIMD&FP register.
//
//go:linkname VgetLowS16 VgetLowS16
//go:noescape
func VgetLowS16(r *arm.Int16X4, v0 *arm.Int16X8)

// Unsigned Halving Add. This instruction adds corresponding unsigned integer values from the two source SIMD&FP registers, shifts each result right one bit, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VhaddqU8 VhaddqU8
//go:noescape
func VhaddqU8(r *arm.Uint8X16, v0 *arm.Uint8X16, v1 *arm.Uint8X16)

// Unsigned Halving Add. This instruction adds corresponding unsigned integer values from the two source SIMD&FP registers, shifts each result right one bit, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VhaddqU32 VhaddqU32
//go:noescape
func VhaddqU32(r *arm.Uint32X4, v0 *arm.Uint32X4, v1 *arm.Uint32X4)

// Unsigned Halving Add. This instruction adds corresponding unsigned integer values from the two source SIMD&FP registers, shifts each result right one bit, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VhaddqU16 VhaddqU16
//go:noescape
func VhaddqU16(r *arm.Uint16X8, v0 *arm.Uint16X8, v1 *arm.Uint16X8)

// Signed Halving Add. This instruction adds corresponding signed integer values from the two source SIMD&FP registers, shifts each result right one bit, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VhaddqS8 VhaddqS8
//go:noescape
func VhaddqS8(r *arm.Int8X16, v0 *arm.Int8X16, v1 *arm.Int8X16)

// Signed Halving Add. This instruction adds corresponding signed integer values from the two source SIMD&FP registers, shifts each result right one bit, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VhaddqS32 VhaddqS32
//go:noescape
func VhaddqS32(r *arm.Int32X4, v0 *arm.Int32X4, v1 *arm.Int32X4)

// Signed Halving Add. This instruction adds corresponding signed integer values from the two source SIMD&FP registers, shifts each result right one bit, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VhaddqS16 VhaddqS16
//go:noescape
func VhaddqS16(r *arm.Int16X8, v0 *arm.Int16X8, v1 *arm.Int16X8)

// Unsigned Halving Add. This instruction adds corresponding unsigned integer values from the two source SIMD&FP registers, shifts each result right one bit, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VhaddU8 VhaddU8
//go:noescape
func VhaddU8(r *arm.Uint8X8, v0 *arm.Uint8X8, v1 *arm.Uint8X8)

// Unsigned Halving Add. This instruction adds corresponding unsigned integer values from the two source SIMD&FP registers, shifts each result right one bit, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VhaddU32 VhaddU32
//go:noescape
func VhaddU32(r *arm.Uint32X2, v0 *arm.Uint32X2, v1 *arm.Uint32X2)

// Unsigned Halving Add. This instruction adds corresponding unsigned integer values from the two source SIMD&FP registers, shifts each result right one bit, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VhaddU16 VhaddU16
//go:noescape
func VhaddU16(r *arm.Uint16X4, v0 *arm.Uint16X4, v1 *arm.Uint16X4)

// Signed Halving Add. This instruction adds corresponding signed integer values from the two source SIMD&FP registers, shifts each result right one bit, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VhaddS8 VhaddS8
//go:noescape
func VhaddS8(r *arm.Int8X8, v0 *arm.Int8X8, v1 *arm.Int8X8)

// Signed Halving Add. This instruction adds corresponding signed integer values from the two source SIMD&FP registers, shifts each result right one bit, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VhaddS32 VhaddS32
//go:noescape
func VhaddS32(r *arm.Int32X2, v0 *arm.Int32X2, v1 *arm.Int32X2)

// Signed Halving Add. This instruction adds corresponding signed integer values from the two source SIMD&FP registers, shifts each result right one bit, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VhaddS16 VhaddS16
//go:noescape
func VhaddS16(r *arm.Int16X4, v0 *arm.Int16X4, v1 *arm.Int16X4)

// Unsigned Halving Subtract. This instruction subtracts the vector elements in the second source SIMD&FP register from the corresponding vector elements in the first source SIMD&FP register, shifts each result right one bit, places each result into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VhsubqU8 VhsubqU8
//go:noescape
func VhsubqU8(r *arm.Uint8X16, v0 *arm.Uint8X16, v1 *arm.Uint8X16)

// Unsigned Halving Subtract. This instruction subtracts the vector elements in the second source SIMD&FP register from the corresponding vector elements in the first source SIMD&FP register, shifts each result right one bit, places each result into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VhsubqU32 VhsubqU32
//go:noescape
func VhsubqU32(r *arm.Uint32X4, v0 *arm.Uint32X4, v1 *arm.Uint32X4)

// Unsigned Halving Subtract. This instruction subtracts the vector elements in the second source SIMD&FP register from the corresponding vector elements in the first source SIMD&FP register, shifts each result right one bit, places each result into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VhsubqU16 VhsubqU16
//go:noescape
func VhsubqU16(r *arm.Uint16X8, v0 *arm.Uint16X8, v1 *arm.Uint16X8)

// Signed Halving Subtract. This instruction subtracts the elements in the vector in the second source SIMD&FP register from the corresponding elements in the vector in the first source SIMD&FP register, shifts each result right one bit, places each result into elements of a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VhsubqS8 VhsubqS8
//go:noescape
func VhsubqS8(r *arm.Int8X16, v0 *arm.Int8X16, v1 *arm.Int8X16)

// Signed Halving Subtract. This instruction subtracts the elements in the vector in the second source SIMD&FP register from the corresponding elements in the vector in the first source SIMD&FP register, shifts each result right one bit, places each result into elements of a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VhsubqS32 VhsubqS32
//go:noescape
func VhsubqS32(r *arm.Int32X4, v0 *arm.Int32X4, v1 *arm.Int32X4)

// Signed Halving Subtract. This instruction subtracts the elements in the vector in the second source SIMD&FP register from the corresponding elements in the vector in the first source SIMD&FP register, shifts each result right one bit, places each result into elements of a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VhsubqS16 VhsubqS16
//go:noescape
func VhsubqS16(r *arm.Int16X8, v0 *arm.Int16X8, v1 *arm.Int16X8)

// Unsigned Halving Subtract. This instruction subtracts the vector elements in the second source SIMD&FP register from the corresponding vector elements in the first source SIMD&FP register, shifts each result right one bit, places each result into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VhsubU8 VhsubU8
//go:noescape
func VhsubU8(r *arm.Uint8X8, v0 *arm.Uint8X8, v1 *arm.Uint8X8)

// Unsigned Halving Subtract. This instruction subtracts the vector elements in the second source SIMD&FP register from the corresponding vector elements in the first source SIMD&FP register, shifts each result right one bit, places each result into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VhsubU32 VhsubU32
//go:noescape
func VhsubU32(r *arm.Uint32X2, v0 *arm.Uint32X2, v1 *arm.Uint32X2)

// Unsigned Halving Subtract. This instruction subtracts the vector elements in the second source SIMD&FP register from the corresponding vector elements in the first source SIMD&FP register, shifts each result right one bit, places each result into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VhsubU16 VhsubU16
//go:noescape
func VhsubU16(r *arm.Uint16X4, v0 *arm.Uint16X4, v1 *arm.Uint16X4)

// Signed Halving Subtract. This instruction subtracts the elements in the vector in the second source SIMD&FP register from the corresponding elements in the vector in the first source SIMD&FP register, shifts each result right one bit, places each result into elements of a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VhsubS8 VhsubS8
//go:noescape
func VhsubS8(r *arm.Int8X8, v0 *arm.Int8X8, v1 *arm.Int8X8)

// Signed Halving Subtract. This instruction subtracts the elements in the vector in the second source SIMD&FP register from the corresponding elements in the vector in the first source SIMD&FP register, shifts each result right one bit, places each result into elements of a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VhsubS32 VhsubS32
//go:noescape
func VhsubS32(r *arm.Int32X2, v0 *arm.Int32X2, v1 *arm.Int32X2)

// Signed Halving Subtract. This instruction subtracts the elements in the vector in the second source SIMD&FP register from the corresponding elements in the vector in the first source SIMD&FP register, shifts each result right one bit, places each result into elements of a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VhsubS16 VhsubS16
//go:noescape
func VhsubS16(r *arm.Int16X4, v0 *arm.Int16X4, v1 *arm.Int16X4)

// Unsigned Maximum (vector). This instruction compares corresponding elements in the vectors in the two source SIMD&FP registers, places the larger of each pair of unsigned integer values into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VmaxqU8 VmaxqU8
//go:noescape
func VmaxqU8(r *arm.Uint8X16, v0 *arm.Uint8X16, v1 *arm.Uint8X16)

// Unsigned Maximum (vector). This instruction compares corresponding elements in the vectors in the two source SIMD&FP registers, places the larger of each pair of unsigned integer values into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VmaxqU32 VmaxqU32
//go:noescape
func VmaxqU32(r *arm.Uint32X4, v0 *arm.Uint32X4, v1 *arm.Uint32X4)

// Unsigned Maximum (vector). This instruction compares corresponding elements in the vectors in the two source SIMD&FP registers, places the larger of each pair of unsigned integer values into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VmaxqU16 VmaxqU16
//go:noescape
func VmaxqU16(r *arm.Uint16X8, v0 *arm.Uint16X8, v1 *arm.Uint16X8)

// Signed Maximum (vector). This instruction compares corresponding elements in the vectors in the two source SIMD&FP registers, places the larger of each pair of signed integer values into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VmaxqS8 VmaxqS8
//go:noescape
func VmaxqS8(r *arm.Int8X16, v0 *arm.Int8X16, v1 *arm.Int8X16)

// Floating-point Maximum (vector). This instruction compares corresponding vector elements in the two source SIMD&FP registers, places the larger of each of the two floating-point values into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VmaxqF32 VmaxqF32
//go:noescape
func VmaxqF32(r *arm.Float32X4, v0 *arm.Float32X4, v1 *arm.Float32X4)

// Signed Maximum (vector). This instruction compares corresponding elements in the vectors in the two source SIMD&FP registers, places the larger of each pair of signed integer values into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VmaxqS32 VmaxqS32
//go:noescape
func VmaxqS32(r *arm.Int32X4, v0 *arm.Int32X4, v1 *arm.Int32X4)

// Signed Maximum (vector). This instruction compares corresponding elements in the vectors in the two source SIMD&FP registers, places the larger of each pair of signed integer values into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VmaxqS16 VmaxqS16
//go:noescape
func VmaxqS16(r *arm.Int16X8, v0 *arm.Int16X8, v1 *arm.Int16X8)

// Unsigned Maximum (vector). This instruction compares corresponding elements in the vectors in the two source SIMD&FP registers, places the larger of each pair of unsigned integer values into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VmaxU8 VmaxU8
//go:noescape
func VmaxU8(r *arm.Uint8X8, v0 *arm.Uint8X8, v1 *arm.Uint8X8)

// Unsigned Maximum (vector). This instruction compares corresponding elements in the vectors in the two source SIMD&FP registers, places the larger of each pair of unsigned integer values into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VmaxU32 VmaxU32
//go:noescape
func VmaxU32(r *arm.Uint32X2, v0 *arm.Uint32X2, v1 *arm.Uint32X2)

// Unsigned Maximum (vector). This instruction compares corresponding elements in the vectors in the two source SIMD&FP registers, places the larger of each pair of unsigned integer values into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VmaxU16 VmaxU16
//go:noescape
func VmaxU16(r *arm.Uint16X4, v0 *arm.Uint16X4, v1 *arm.Uint16X4)

// Signed Maximum (vector). This instruction compares corresponding elements in the vectors in the two source SIMD&FP registers, places the larger of each pair of signed integer values into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VmaxS8 VmaxS8
//go:noescape
func VmaxS8(r *arm.Int8X8, v0 *arm.Int8X8, v1 *arm.Int8X8)

// Floating-point Maximum (vector). This instruction compares corresponding vector elements in the two source SIMD&FP registers, places the larger of each of the two floating-point values into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VmaxF32 VmaxF32
//go:noescape
func VmaxF32(r *arm.Float32X2, v0 *arm.Float32X2, v1 *arm.Float32X2)

// Signed Maximum (vector). This instruction compares corresponding elements in the vectors in the two source SIMD&FP registers, places the larger of each pair of signed integer values into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VmaxS32 VmaxS32
//go:noescape
func VmaxS32(r *arm.Int32X2, v0 *arm.Int32X2, v1 *arm.Int32X2)

// Signed Maximum (vector). This instruction compares corresponding elements in the vectors in the two source SIMD&FP registers, places the larger of each pair of signed integer values into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VmaxS16 VmaxS16
//go:noescape
func VmaxS16(r *arm.Int16X4, v0 *arm.Int16X4, v1 *arm.Int16X4)

// Unsigned Minimum (vector). This instruction compares corresponding vector elements in the two source SIMD&FP registers, places the smaller of each of the two unsigned integer values into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VminqU8 VminqU8
//go:noescape
func VminqU8(r *arm.Uint8X16, v0 *arm.Uint8X16, v1 *arm.Uint8X16)

// Unsigned Minimum (vector). This instruction compares corresponding vector elements in the two source SIMD&FP registers, places the smaller of each of the two unsigned integer values into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VminqU32 VminqU32
//go:noescape
func VminqU32(r *arm.Uint32X4, v0 *arm.Uint32X4, v1 *arm.Uint32X4)

// Unsigned Minimum (vector). This instruction compares corresponding vector elements in the two source SIMD&FP registers, places the smaller of each of the two unsigned integer values into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VminqU16 VminqU16
//go:noescape
func VminqU16(r *arm.Uint16X8, v0 *arm.Uint16X8, v1 *arm.Uint16X8)

// Signed Minimum (vector). This instruction compares corresponding elements in the vectors in the two source SIMD&FP registers, places the smaller of each of the two signed integer values into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VminqS8 VminqS8
//go:noescape
func VminqS8(r *arm.Int8X16, v0 *arm.Int8X16, v1 *arm.Int8X16)

// Floating-point minimum (vector). This instruction compares corresponding elements in the vectors in the two source SIMD&FP registers, places the smaller of each of the two floating-point values into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VminqF32 VminqF32
//go:noescape
func VminqF32(r *arm.Float32X4, v0 *arm.Float32X4, v1 *arm.Float32X4)

// Signed Minimum (vector). This instruction compares corresponding elements in the vectors in the two source SIMD&FP registers, places the smaller of each of the two signed integer values into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VminqS32 VminqS32
//go:noescape
func VminqS32(r *arm.Int32X4, v0 *arm.Int32X4, v1 *arm.Int32X4)

// Signed Minimum (vector). This instruction compares corresponding elements in the vectors in the two source SIMD&FP registers, places the smaller of each of the two signed integer values into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VminqS16 VminqS16
//go:noescape
func VminqS16(r *arm.Int16X8, v0 *arm.Int16X8, v1 *arm.Int16X8)

// Unsigned Minimum (vector). This instruction compares corresponding vector elements in the two source SIMD&FP registers, places the smaller of each of the two unsigned integer values into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VminU8 VminU8
//go:noescape
func VminU8(r *arm.Uint8X8, v0 *arm.Uint8X8, v1 *arm.Uint8X8)

// Unsigned Minimum (vector). This instruction compares corresponding vector elements in the two source SIMD&FP registers, places the smaller of each of the two unsigned integer values into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VminU32 VminU32
//go:noescape
func VminU32(r *arm.Uint32X2, v0 *arm.Uint32X2, v1 *arm.Uint32X2)

// Unsigned Minimum (vector). This instruction compares corresponding vector elements in the two source SIMD&FP registers, places the smaller of each of the two unsigned integer values into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VminU16 VminU16
//go:noescape
func VminU16(r *arm.Uint16X4, v0 *arm.Uint16X4, v1 *arm.Uint16X4)

// Signed Minimum (vector). This instruction compares corresponding elements in the vectors in the two source SIMD&FP registers, places the smaller of each of the two signed integer values into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VminS8 VminS8
//go:noescape
func VminS8(r *arm.Int8X8, v0 *arm.Int8X8, v1 *arm.Int8X8)

// Floating-point minimum (vector). This instruction compares corresponding elements in the vectors in the two source SIMD&FP registers, places the smaller of each of the two floating-point values into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VminF32 VminF32
//go:noescape
func VminF32(r *arm.Float32X2, v0 *arm.Float32X2, v1 *arm.Float32X2)

// Signed Minimum (vector). This instruction compares corresponding elements in the vectors in the two source SIMD&FP registers, places the smaller of each of the two signed integer values into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VminS32 VminS32
//go:noescape
func VminS32(r *arm.Int32X2, v0 *arm.Int32X2, v1 *arm.Int32X2)

// Signed Minimum (vector). This instruction compares corresponding elements in the vectors in the two source SIMD&FP registers, places the smaller of each of the two signed integer values into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VminS16 VminS16
//go:noescape
func VminS16(r *arm.Int16X4, v0 *arm.Int16X4, v1 *arm.Int16X4)

// Multiply-Add to accumulator (vector). This instruction multiplies corresponding elements in the vectors of the two source SIMD&FP registers, and accumulates the results with the vector elements of the destination SIMD&FP register.
//
//go:linkname VmlaqU8 VmlaqU8
//go:noescape
func VmlaqU8(r *arm.Uint8X16, v0 *arm.Uint8X16, v1 *arm.Uint8X16, v2 *arm.Uint8X16)

// Multiply-Add to accumulator (vector). This instruction multiplies corresponding elements in the vectors of the two source SIMD&FP registers, and accumulates the results with the vector elements of the destination SIMD&FP register.
//
//go:linkname VmlaqU32 VmlaqU32
//go:noescape
func VmlaqU32(r *arm.Uint32X4, v0 *arm.Uint32X4, v1 *arm.Uint32X4, v2 *arm.Uint32X4)

// Multiply-Add to accumulator (vector). This instruction multiplies corresponding elements in the vectors of the two source SIMD&FP registers, and accumulates the results with the vector elements of the destination SIMD&FP register.
//
//go:linkname VmlaqU16 VmlaqU16
//go:noescape
func VmlaqU16(r *arm.Uint16X8, v0 *arm.Uint16X8, v1 *arm.Uint16X8, v2 *arm.Uint16X8)

// Multiply-Add to accumulator (vector). This instruction multiplies corresponding elements in the vectors of the two source SIMD&FP registers, and accumulates the results with the vector elements of the destination SIMD&FP register.
//
//go:linkname VmlaqS8 VmlaqS8
//go:noescape
func VmlaqS8(r *arm.Int8X16, v0 *arm.Int8X16, v1 *arm.Int8X16, v2 *arm.Int8X16)

// Floating-point multiply-add to accumulator
//
//go:linkname VmlaqF32 VmlaqF32
//go:noescape
func VmlaqF32(r *arm.Float32X4, v0 *arm.Float32X4, v1 *arm.Float32X4, v2 *arm.Float32X4)

// Multiply-Add to accumulator (vector). This instruction multiplies corresponding elements in the vectors of the two source SIMD&FP registers, and accumulates the results with the vector elements of the destination SIMD&FP register.
//
//go:linkname VmlaqS32 VmlaqS32
//go:noescape
func VmlaqS32(r *arm.Int32X4, v0 *arm.Int32X4, v1 *arm.Int32X4, v2 *arm.Int32X4)

// Multiply-Add to accumulator (vector). This instruction multiplies corresponding elements in the vectors of the two source SIMD&FP registers, and accumulates the results with the vector elements of the destination SIMD&FP register.
//
//go:linkname VmlaqS16 VmlaqS16
//go:noescape
func VmlaqS16(r *arm.Int16X8, v0 *arm.Int16X8, v1 *arm.Int16X8, v2 *arm.Int16X8)

// Multiply-Add to accumulator (vector). This instruction multiplies corresponding elements in the vectors of the two source SIMD&FP registers, and accumulates the results with the vector elements of the destination SIMD&FP register.
//
//go:linkname VmlaU8 VmlaU8
//go:noescape
func VmlaU8(r *arm.Uint8X8, v0 *arm.Uint8X8, v1 *arm.Uint8X8, v2 *arm.Uint8X8)

// Multiply-Add to accumulator (vector). This instruction multiplies corresponding elements in the vectors of the two source SIMD&FP registers, and accumulates the results with the vector elements of the destination SIMD&FP register.
//
//go:linkname VmlaU32 VmlaU32
//go:noescape
func VmlaU32(r *arm.Uint32X2, v0 *arm.Uint32X2, v1 *arm.Uint32X2, v2 *arm.Uint32X2)

// Multiply-Add to accumulator (vector). This instruction multiplies corresponding elements in the vectors of the two source SIMD&FP registers, and accumulates the results with the vector elements of the destination SIMD&FP register.
//
//go:linkname VmlaU16 VmlaU16
//go:noescape
func VmlaU16(r *arm.Uint16X4, v0 *arm.Uint16X4, v1 *arm.Uint16X4, v2 *arm.Uint16X4)

// Multiply-Add to accumulator (vector). This instruction multiplies corresponding elements in the vectors of the two source SIMD&FP registers, and accumulates the results with the vector elements of the destination SIMD&FP register.
//
//go:linkname VmlaS8 VmlaS8
//go:noescape
func VmlaS8(r *arm.Int8X8, v0 *arm.Int8X8, v1 *arm.Int8X8, v2 *arm.Int8X8)

// Floating-point multiply-add to accumulator
//
//go:linkname VmlaF32 VmlaF32
//go:noescape
func VmlaF32(r *arm.Float32X2, v0 *arm.Float32X2, v1 *arm.Float32X2, v2 *arm.Float32X2)

// Multiply-Add to accumulator (vector). This instruction multiplies corresponding elements in the vectors of the two source SIMD&FP registers, and accumulates the results with the vector elements of the destination SIMD&FP register.
//
//go:linkname VmlaS32 VmlaS32
//go:noescape
func VmlaS32(r *arm.Int32X2, v0 *arm.Int32X2, v1 *arm.Int32X2, v2 *arm.Int32X2)

// Multiply-Add to accumulator (vector). This instruction multiplies corresponding elements in the vectors of the two source SIMD&FP registers, and accumulates the results with the vector elements of the destination SIMD&FP register.
//
//go:linkname VmlaS16 VmlaS16
//go:noescape
func VmlaS16(r *arm.Int16X4, v0 *arm.Int16X4, v1 *arm.Int16X4, v2 *arm.Int16X4)

// Vector multiply accumulate with scalar
//
//go:linkname VmlaqNU32 VmlaqNU32
//go:noescape
func VmlaqNU32(r *arm.Uint32X4, v0 *arm.Uint32X4, v1 *arm.Uint32X4, v2 *arm.Uint32)

// Vector multiply accumulate with scalar
//
//go:linkname VmlaqNU16 VmlaqNU16
//go:noescape
func VmlaqNU16(r *arm.Uint16X8, v0 *arm.Uint16X8, v1 *arm.Uint16X8, v2 *arm.Uint16)

// Vector multiply accumulate with scalar
//
//go:linkname VmlaqNF32 VmlaqNF32
//go:noescape
func VmlaqNF32(r *arm.Float32X4, v0 *arm.Float32X4, v1 *arm.Float32X4, v2 *arm.Float32)

// Vector multiply accumulate with scalar
//
//go:linkname VmlaqNS32 VmlaqNS32
//go:noescape
func VmlaqNS32(r *arm.Int32X4, v0 *arm.Int32X4, v1 *arm.Int32X4, v2 *arm.Int32)

// Vector multiply accumulate with scalar
//
//go:linkname VmlaqNS16 VmlaqNS16
//go:noescape
func VmlaqNS16(r *arm.Int16X8, v0 *arm.Int16X8, v1 *arm.Int16X8, v2 *arm.Int16)

// Vector multiply accumulate with scalar
//
//go:linkname VmlaNU32 VmlaNU32
//go:noescape
func VmlaNU32(r *arm.Uint32X2, v0 *arm.Uint32X2, v1 *arm.Uint32X2, v2 *arm.Uint32)

// Vector multiply accumulate with scalar
//
//go:linkname VmlaNU16 VmlaNU16
//go:noescape
func VmlaNU16(r *arm.Uint16X4, v0 *arm.Uint16X4, v1 *arm.Uint16X4, v2 *arm.Uint16)

// Vector multiply accumulate with scalar
//
//go:linkname VmlaNF32 VmlaNF32
//go:noescape
func VmlaNF32(r *arm.Float32X2, v0 *arm.Float32X2, v1 *arm.Float32X2, v2 *arm.Float32)

// Vector multiply accumulate with scalar
//
//go:linkname VmlaNS32 VmlaNS32
//go:noescape
func VmlaNS32(r *arm.Int32X2, v0 *arm.Int32X2, v1 *arm.Int32X2, v2 *arm.Int32)

// Vector multiply accumulate with scalar
//
//go:linkname VmlaNS16 VmlaNS16
//go:noescape
func VmlaNS16(r *arm.Int16X4, v0 *arm.Int16X4, v1 *arm.Int16X4, v2 *arm.Int16)

// Multiply-Subtract from accumulator (vector). This instruction multiplies corresponding elements in the vectors of the two source SIMD&FP registers, and subtracts the results from the vector elements of the destination SIMD&FP register.
//
//go:linkname VmlsqU8 VmlsqU8
//go:noescape
func VmlsqU8(r *arm.Uint8X16, v0 *arm.Uint8X16, v1 *arm.Uint8X16, v2 *arm.Uint8X16)

// Multiply-Subtract from accumulator (vector). This instruction multiplies corresponding elements in the vectors of the two source SIMD&FP registers, and subtracts the results from the vector elements of the destination SIMD&FP register.
//
//go:linkname VmlsqU32 VmlsqU32
//go:noescape
func VmlsqU32(r *arm.Uint32X4, v0 *arm.Uint32X4, v1 *arm.Uint32X4, v2 *arm.Uint32X4)

// Multiply-Subtract from accumulator (vector). This instruction multiplies corresponding elements in the vectors of the two source SIMD&FP registers, and subtracts the results from the vector elements of the destination SIMD&FP register.
//
//go:linkname VmlsqU16 VmlsqU16
//go:noescape
func VmlsqU16(r *arm.Uint16X8, v0 *arm.Uint16X8, v1 *arm.Uint16X8, v2 *arm.Uint16X8)

// Multiply-Subtract from accumulator (vector). This instruction multiplies corresponding elements in the vectors of the two source SIMD&FP registers, and subtracts the results from the vector elements of the destination SIMD&FP register.
//
//go:linkname VmlsqS8 VmlsqS8
//go:noescape
func VmlsqS8(r *arm.Int8X16, v0 *arm.Int8X16, v1 *arm.Int8X16, v2 *arm.Int8X16)

// Multiply-subtract from accumulator
//
//go:linkname VmlsqF32 VmlsqF32
//go:noescape
func VmlsqF32(r *arm.Float32X4, v0 *arm.Float32X4, v1 *arm.Float32X4, v2 *arm.Float32X4)

// Multiply-Subtract from accumulator (vector). This instruction multiplies corresponding elements in the vectors of the two source SIMD&FP registers, and subtracts the results from the vector elements of the destination SIMD&FP register.
//
//go:linkname VmlsqS32 VmlsqS32
//go:noescape
func VmlsqS32(r *arm.Int32X4, v0 *arm.Int32X4, v1 *arm.Int32X4, v2 *arm.Int32X4)

// Multiply-Subtract from accumulator (vector). This instruction multiplies corresponding elements in the vectors of the two source SIMD&FP registers, and subtracts the results from the vector elements of the destination SIMD&FP register.
//
//go:linkname VmlsqS16 VmlsqS16
//go:noescape
func VmlsqS16(r *arm.Int16X8, v0 *arm.Int16X8, v1 *arm.Int16X8, v2 *arm.Int16X8)

// Multiply-Subtract from accumulator (vector). This instruction multiplies corresponding elements in the vectors of the two source SIMD&FP registers, and subtracts the results from the vector elements of the destination SIMD&FP register.
//
//go:linkname VmlsU8 VmlsU8
//go:noescape
func VmlsU8(r *arm.Uint8X8, v0 *arm.Uint8X8, v1 *arm.Uint8X8, v2 *arm.Uint8X8)

// Multiply-Subtract from accumulator (vector). This instruction multiplies corresponding elements in the vectors of the two source SIMD&FP registers, and subtracts the results from the vector elements of the destination SIMD&FP register.
//
//go:linkname VmlsU32 VmlsU32
//go:noescape
func VmlsU32(r *arm.Uint32X2, v0 *arm.Uint32X2, v1 *arm.Uint32X2, v2 *arm.Uint32X2)

// Multiply-Subtract from accumulator (vector). This instruction multiplies corresponding elements in the vectors of the two source SIMD&FP registers, and subtracts the results from the vector elements of the destination SIMD&FP register.
//
//go:linkname VmlsU16 VmlsU16
//go:noescape
func VmlsU16(r *arm.Uint16X4, v0 *arm.Uint16X4, v1 *arm.Uint16X4, v2 *arm.Uint16X4)

// Multiply-Subtract from accumulator (vector). This instruction multiplies corresponding elements in the vectors of the two source SIMD&FP registers, and subtracts the results from the vector elements of the destination SIMD&FP register.
//
//go:linkname VmlsS8 VmlsS8
//go:noescape
func VmlsS8(r *arm.Int8X8, v0 *arm.Int8X8, v1 *arm.Int8X8, v2 *arm.Int8X8)

// Multiply-subtract from accumulator
//
//go:linkname VmlsF32 VmlsF32
//go:noescape
func VmlsF32(r *arm.Float32X2, v0 *arm.Float32X2, v1 *arm.Float32X2, v2 *arm.Float32X2)

// Multiply-Subtract from accumulator (vector). This instruction multiplies corresponding elements in the vectors of the two source SIMD&FP registers, and subtracts the results from the vector elements of the destination SIMD&FP register.
//
//go:linkname VmlsS32 VmlsS32
//go:noescape
func VmlsS32(r *arm.Int32X2, v0 *arm.Int32X2, v1 *arm.Int32X2, v2 *arm.Int32X2)

// Multiply-Subtract from accumulator (vector). This instruction multiplies corresponding elements in the vectors of the two source SIMD&FP registers, and subtracts the results from the vector elements of the destination SIMD&FP register.
//
//go:linkname VmlsS16 VmlsS16
//go:noescape
func VmlsS16(r *arm.Int16X4, v0 *arm.Int16X4, v1 *arm.Int16X4, v2 *arm.Int16X4)

// Vector multiply subtract with scalar
//
//go:linkname VmlsqNU32 VmlsqNU32
//go:noescape
func VmlsqNU32(r *arm.Uint32X4, v0 *arm.Uint32X4, v1 *arm.Uint32X4, v2 *arm.Uint32)

// Vector multiply subtract with scalar
//
//go:linkname VmlsqNU16 VmlsqNU16
//go:noescape
func VmlsqNU16(r *arm.Uint16X8, v0 *arm.Uint16X8, v1 *arm.Uint16X8, v2 *arm.Uint16)

// Vector multiply subtract with scalar
//
//go:linkname VmlsqNF32 VmlsqNF32
//go:noescape
func VmlsqNF32(r *arm.Float32X4, v0 *arm.Float32X4, v1 *arm.Float32X4, v2 *arm.Float32)

// Vector multiply subtract with scalar
//
//go:linkname VmlsqNS32 VmlsqNS32
//go:noescape
func VmlsqNS32(r *arm.Int32X4, v0 *arm.Int32X4, v1 *arm.Int32X4, v2 *arm.Int32)

// Vector multiply subtract with scalar
//
//go:linkname VmlsqNS16 VmlsqNS16
//go:noescape
func VmlsqNS16(r *arm.Int16X8, v0 *arm.Int16X8, v1 *arm.Int16X8, v2 *arm.Int16)

// Vector multiply subtract with scalar
//
//go:linkname VmlsNU32 VmlsNU32
//go:noescape
func VmlsNU32(r *arm.Uint32X2, v0 *arm.Uint32X2, v1 *arm.Uint32X2, v2 *arm.Uint32)

// Vector multiply subtract with scalar
//
//go:linkname VmlsNU16 VmlsNU16
//go:noescape
func VmlsNU16(r *arm.Uint16X4, v0 *arm.Uint16X4, v1 *arm.Uint16X4, v2 *arm.Uint16)

// Vector multiply subtract with scalar
//
//go:linkname VmlsNF32 VmlsNF32
//go:noescape
func VmlsNF32(r *arm.Float32X2, v0 *arm.Float32X2, v1 *arm.Float32X2, v2 *arm.Float32)

// Vector multiply subtract with scalar
//
//go:linkname VmlsNS32 VmlsNS32
//go:noescape
func VmlsNS32(r *arm.Int32X2, v0 *arm.Int32X2, v1 *arm.Int32X2, v2 *arm.Int32)

// Vector multiply subtract with scalar
//
//go:linkname VmlsNS16 VmlsNS16
//go:noescape
func VmlsNS16(r *arm.Int16X4, v0 *arm.Int16X4, v1 *arm.Int16X4, v2 *arm.Int16)

// Duplicate vector element to vector or scalar. This instruction duplicates the vector element at the specified element index in the source SIMD&FP register into a scalar or each element in a vector, and writes the result to the destination SIMD&FP register.
//
//go:linkname VmovNP8 VmovNP8
//go:noescape
func VmovNP8(r *arm.Poly8X8, v0 *arm.Poly8)

// Duplicate vector element to vector or scalar. This instruction duplicates the vector element at the specified element index in the source SIMD&FP register into a scalar or each element in a vector, and writes the result to the destination SIMD&FP register.
//
//go:linkname VmovNP16 VmovNP16
//go:noescape
func VmovNP16(r *arm.Poly16X4, v0 *arm.Poly16)

// Duplicate vector element to vector or scalar. This instruction duplicates the vector element at the specified element index in the source SIMD&FP register into a scalar or each element in a vector, and writes the result to the destination SIMD&FP register.
//
//go:linkname VmovqNP8 VmovqNP8
//go:noescape
func VmovqNP8(r *arm.Poly8X16, v0 *arm.Poly8)

// Duplicate vector element to vector or scalar. This instruction duplicates the vector element at the specified element index in the source SIMD&FP register into a scalar or each element in a vector, and writes the result to the destination SIMD&FP register.
//
//go:linkname VmovqNP16 VmovqNP16
//go:noescape
func VmovqNP16(r *arm.Poly16X8, v0 *arm.Poly16)

// Duplicate vector element to vector or scalar. This instruction duplicates the vector element at the specified element index in the source SIMD&FP register into a scalar or each element in a vector, and writes the result to the destination SIMD&FP register.
//
//go:linkname VmovqNU8 VmovqNU8
//go:noescape
func VmovqNU8(r *arm.Uint8X16, v0 *arm.Uint8)

// Duplicate vector element to vector or scalar. This instruction duplicates the vector element at the specified element index in the source SIMD&FP register into a scalar or each element in a vector, and writes the result to the destination SIMD&FP register.
//
//go:linkname VmovqNU32 VmovqNU32
//go:noescape
func VmovqNU32(r *arm.Uint32X4, v0 *arm.Uint32)

// Duplicate vector element to vector or scalar. This instruction duplicates the vector element at the specified element index in the source SIMD&FP register into a scalar or each element in a vector, and writes the result to the destination SIMD&FP register.
//
//go:linkname VmovqNU64 VmovqNU64
//go:noescape
func VmovqNU64(r *arm.Uint64X2, v0 *arm.Uint64)

// Duplicate vector element to vector or scalar. This instruction duplicates the vector element at the specified element index in the source SIMD&FP register into a scalar or each element in a vector, and writes the result to the destination SIMD&FP register.
//
//go:linkname VmovqNU16 VmovqNU16
//go:noescape
func VmovqNU16(r *arm.Uint16X8, v0 *arm.Uint16)

// Duplicate vector element to vector or scalar. This instruction duplicates the vector element at the specified element index in the source SIMD&FP register into a scalar or each element in a vector, and writes the result to the destination SIMD&FP register.
//
//go:linkname VmovqNS8 VmovqNS8
//go:noescape
func VmovqNS8(r *arm.Int8X16, v0 *arm.Int8)

// Duplicate vector element to vector or scalar. This instruction duplicates the vector element at the specified element index in the source SIMD&FP register into a scalar or each element in a vector, and writes the result to the destination SIMD&FP register.
//
//go:linkname VmovqNF32 VmovqNF32
//go:noescape
func VmovqNF32(r *arm.Float32X4, v0 *arm.Float32)

// Duplicate vector element to vector or scalar. This instruction duplicates the vector element at the specified element index in the source SIMD&FP register into a scalar or each element in a vector, and writes the result to the destination SIMD&FP register.
//
//go:linkname VmovqNS32 VmovqNS32
//go:noescape
func VmovqNS32(r *arm.Int32X4, v0 *arm.Int32)

// Duplicate vector element to vector or scalar. This instruction duplicates the vector element at the specified element index in the source SIMD&FP register into a scalar or each element in a vector, and writes the result to the destination SIMD&FP register.
//
//go:linkname VmovqNS64 VmovqNS64
//go:noescape
func VmovqNS64(r *arm.Int64X2, v0 *arm.Int64)

// Duplicate vector element to vector or scalar. This instruction duplicates the vector element at the specified element index in the source SIMD&FP register into a scalar or each element in a vector, and writes the result to the destination SIMD&FP register.
//
//go:linkname VmovqNS16 VmovqNS16
//go:noescape
func VmovqNS16(r *arm.Int16X8, v0 *arm.Int16)

// Duplicate vector element to vector or scalar. This instruction duplicates the vector element at the specified element index in the source SIMD&FP register into a scalar or each element in a vector, and writes the result to the destination SIMD&FP register.
//
//go:linkname VmovNU8 VmovNU8
//go:noescape
func VmovNU8(r *arm.Uint8X8, v0 *arm.Uint8)

// Duplicate vector element to vector or scalar. This instruction duplicates the vector element at the specified element index in the source SIMD&FP register into a scalar or each element in a vector, and writes the result to the destination SIMD&FP register.
//
//go:linkname VmovNU32 VmovNU32
//go:noescape
func VmovNU32(r *arm.Uint32X2, v0 *arm.Uint32)

// Duplicate vector element to vector or scalar. This instruction duplicates the vector element at the specified element index in the source SIMD&FP register into a scalar or each element in a vector, and writes the result to the destination SIMD&FP register.
//
//go:linkname VmovNU64 VmovNU64
//go:noescape
func VmovNU64(r *arm.Uint64X1, v0 *arm.Uint64)

// Duplicate vector element to vector or scalar. This instruction duplicates the vector element at the specified element index in the source SIMD&FP register into a scalar or each element in a vector, and writes the result to the destination SIMD&FP register.
//
//go:linkname VmovNU16 VmovNU16
//go:noescape
func VmovNU16(r *arm.Uint16X4, v0 *arm.Uint16)

// Duplicate vector element to vector or scalar. This instruction duplicates the vector element at the specified element index in the source SIMD&FP register into a scalar or each element in a vector, and writes the result to the destination SIMD&FP register.
//
//go:linkname VmovNS8 VmovNS8
//go:noescape
func VmovNS8(r *arm.Int8X8, v0 *arm.Int8)

// Duplicate vector element to vector or scalar. This instruction duplicates the vector element at the specified element index in the source SIMD&FP register into a scalar or each element in a vector, and writes the result to the destination SIMD&FP register.
//
//go:linkname VmovNF32 VmovNF32
//go:noescape
func VmovNF32(r *arm.Float32X2, v0 *arm.Float32)

// Duplicate vector element to vector or scalar. This instruction duplicates the vector element at the specified element index in the source SIMD&FP register into a scalar or each element in a vector, and writes the result to the destination SIMD&FP register.
//
//go:linkname VmovNS32 VmovNS32
//go:noescape
func VmovNS32(r *arm.Int32X2, v0 *arm.Int32)

// Duplicate vector element to vector or scalar. This instruction duplicates the vector element at the specified element index in the source SIMD&FP register into a scalar or each element in a vector, and writes the result to the destination SIMD&FP register.
//
//go:linkname VmovNS64 VmovNS64
//go:noescape
func VmovNS64(r *arm.Int64X1, v0 *arm.Int64)

// Duplicate vector element to vector or scalar. This instruction duplicates the vector element at the specified element index in the source SIMD&FP register into a scalar or each element in a vector, and writes the result to the destination SIMD&FP register.
//
//go:linkname VmovNS16 VmovNS16
//go:noescape
func VmovNS16(r *arm.Int16X4, v0 *arm.Int16)

// Vector move
//
//go:linkname VmovlU8 VmovlU8
//go:noescape
func VmovlU8(r *arm.Uint16X8, v0 *arm.Uint8X8)

// Vector move
//
//go:linkname VmovlU32 VmovlU32
//go:noescape
func VmovlU32(r *arm.Uint64X2, v0 *arm.Uint32X2)

// Vector move
//
//go:linkname VmovlU16 VmovlU16
//go:noescape
func VmovlU16(r *arm.Uint32X4, v0 *arm.Uint16X4)

// Vector move
//
//go:linkname VmovlS8 VmovlS8
//go:noescape
func VmovlS8(r *arm.Int16X8, v0 *arm.Int8X8)

// Vector move
//
//go:linkname VmovlS32 VmovlS32
//go:noescape
func VmovlS32(r *arm.Int64X2, v0 *arm.Int32X2)

// Vector move
//
//go:linkname VmovlS16 VmovlS16
//go:noescape
func VmovlS16(r *arm.Int32X4, v0 *arm.Int16X4)

// Extract Narrow. This instruction reads each vector element from the source SIMD&FP register, narrows each value to half the original width, places the result into a vector, and writes the vector to the lower or upper half of the destination SIMD&FP register. The destination vector elements are half as long as the source vector elements.
//
//go:linkname VmovnU32 VmovnU32
//go:noescape
func VmovnU32(r *arm.Uint16X4, v0 *arm.Uint32X4)

// Extract Narrow. This instruction reads each vector element from the source SIMD&FP register, narrows each value to half the original width, places the result into a vector, and writes the vector to the lower or upper half of the destination SIMD&FP register. The destination vector elements are half as long as the source vector elements.
//
//go:linkname VmovnU64 VmovnU64
//go:noescape
func VmovnU64(r *arm.Uint32X2, v0 *arm.Uint64X2)

// Extract Narrow. This instruction reads each vector element from the source SIMD&FP register, narrows each value to half the original width, places the result into a vector, and writes the vector to the lower or upper half of the destination SIMD&FP register. The destination vector elements are half as long as the source vector elements.
//
//go:linkname VmovnU16 VmovnU16
//go:noescape
func VmovnU16(r *arm.Uint8X8, v0 *arm.Uint16X8)

// Extract Narrow. This instruction reads each vector element from the source SIMD&FP register, narrows each value to half the original width, places the result into a vector, and writes the vector to the lower or upper half of the destination SIMD&FP register. The destination vector elements are half as long as the source vector elements.
//
//go:linkname VmovnS32 VmovnS32
//go:noescape
func VmovnS32(r *arm.Int16X4, v0 *arm.Int32X4)

// Extract Narrow. This instruction reads each vector element from the source SIMD&FP register, narrows each value to half the original width, places the result into a vector, and writes the vector to the lower or upper half of the destination SIMD&FP register. The destination vector elements are half as long as the source vector elements.
//
//go:linkname VmovnS64 VmovnS64
//go:noescape
func VmovnS64(r *arm.Int32X2, v0 *arm.Int64X2)

// Extract Narrow. This instruction reads each vector element from the source SIMD&FP register, narrows each value to half the original width, places the result into a vector, and writes the vector to the lower or upper half of the destination SIMD&FP register. The destination vector elements are half as long as the source vector elements.
//
//go:linkname VmovnS16 VmovnS16
//go:noescape
func VmovnS16(r *arm.Int8X8, v0 *arm.Int16X8)

// Multiply (vector). This instruction multiplies corresponding elements in the vectors of the two source SIMD&FP registers, places the results in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VmulqU8 VmulqU8
//go:noescape
func VmulqU8(r *arm.Uint8X16, v0 *arm.Uint8X16, v1 *arm.Uint8X16)

// Multiply (vector). This instruction multiplies corresponding elements in the vectors of the two source SIMD&FP registers, places the results in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VmulqU32 VmulqU32
//go:noescape
func VmulqU32(r *arm.Uint32X4, v0 *arm.Uint32X4, v1 *arm.Uint32X4)

// Multiply (vector). This instruction multiplies corresponding elements in the vectors of the two source SIMD&FP registers, places the results in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VmulqU16 VmulqU16
//go:noescape
func VmulqU16(r *arm.Uint16X8, v0 *arm.Uint16X8, v1 *arm.Uint16X8)

// Multiply (vector). This instruction multiplies corresponding elements in the vectors of the two source SIMD&FP registers, places the results in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VmulqS8 VmulqS8
//go:noescape
func VmulqS8(r *arm.Int8X16, v0 *arm.Int8X16, v1 *arm.Int8X16)

// Floating-point Multiply (vector). This instruction multiplies corresponding floating-point values in the vectors in the two source SIMD&FP registers, places the result in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VmulqF32 VmulqF32
//go:noescape
func VmulqF32(r *arm.Float32X4, v0 *arm.Float32X4, v1 *arm.Float32X4)

// Multiply (vector). This instruction multiplies corresponding elements in the vectors of the two source SIMD&FP registers, places the results in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VmulqS32 VmulqS32
//go:noescape
func VmulqS32(r *arm.Int32X4, v0 *arm.Int32X4, v1 *arm.Int32X4)

// Multiply (vector). This instruction multiplies corresponding elements in the vectors of the two source SIMD&FP registers, places the results in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VmulqS16 VmulqS16
//go:noescape
func VmulqS16(r *arm.Int16X8, v0 *arm.Int16X8, v1 *arm.Int16X8)

// Multiply (vector). This instruction multiplies corresponding elements in the vectors of the two source SIMD&FP registers, places the results in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VmulU8 VmulU8
//go:noescape
func VmulU8(r *arm.Uint8X8, v0 *arm.Uint8X8, v1 *arm.Uint8X8)

// Multiply (vector). This instruction multiplies corresponding elements in the vectors of the two source SIMD&FP registers, places the results in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VmulU32 VmulU32
//go:noescape
func VmulU32(r *arm.Uint32X2, v0 *arm.Uint32X2, v1 *arm.Uint32X2)

// Multiply (vector). This instruction multiplies corresponding elements in the vectors of the two source SIMD&FP registers, places the results in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VmulU16 VmulU16
//go:noescape
func VmulU16(r *arm.Uint16X4, v0 *arm.Uint16X4, v1 *arm.Uint16X4)

// Multiply (vector). This instruction multiplies corresponding elements in the vectors of the two source SIMD&FP registers, places the results in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VmulS8 VmulS8
//go:noescape
func VmulS8(r *arm.Int8X8, v0 *arm.Int8X8, v1 *arm.Int8X8)

// Floating-point Multiply (vector). This instruction multiplies corresponding floating-point values in the vectors in the two source SIMD&FP registers, places the result in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VmulF32 VmulF32
//go:noescape
func VmulF32(r *arm.Float32X2, v0 *arm.Float32X2, v1 *arm.Float32X2)

// Multiply (vector). This instruction multiplies corresponding elements in the vectors of the two source SIMD&FP registers, places the results in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VmulS32 VmulS32
//go:noescape
func VmulS32(r *arm.Int32X2, v0 *arm.Int32X2, v1 *arm.Int32X2)

// Multiply (vector). This instruction multiplies corresponding elements in the vectors of the two source SIMD&FP registers, places the results in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VmulS16 VmulS16
//go:noescape
func VmulS16(r *arm.Int16X4, v0 *arm.Int16X4, v1 *arm.Int16X4)

// Polynomial Multiply. This instruction multiplies corresponding elements in the vectors of the two source SIMD&FP registers, places the results in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VmulP8 VmulP8
//go:noescape
func VmulP8(r *arm.Poly8X8, v0 *arm.Poly8X8, v1 *arm.Poly8X8)

// Polynomial Multiply. This instruction multiplies corresponding elements in the vectors of the two source SIMD&FP registers, places the results in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VmulqP8 VmulqP8
//go:noescape
func VmulqP8(r *arm.Poly8X16, v0 *arm.Poly8X16, v1 *arm.Poly8X16)

// Vector multiply by scalar
//
//go:linkname VmulqNU32 VmulqNU32
//go:noescape
func VmulqNU32(r *arm.Uint32X4, v0 *arm.Uint32X4, v1 *arm.Uint32)

// Vector multiply by scalar
//
//go:linkname VmulqNU16 VmulqNU16
//go:noescape
func VmulqNU16(r *arm.Uint16X8, v0 *arm.Uint16X8, v1 *arm.Uint16)

// Vector multiply by scalar
//
//go:linkname VmulqNF32 VmulqNF32
//go:noescape
func VmulqNF32(r *arm.Float32X4, v0 *arm.Float32X4, v1 *arm.Float32)

// Vector multiply by scalar
//
//go:linkname VmulqNS32 VmulqNS32
//go:noescape
func VmulqNS32(r *arm.Int32X4, v0 *arm.Int32X4, v1 *arm.Int32)

// Vector multiply by scalar
//
//go:linkname VmulqNS16 VmulqNS16
//go:noescape
func VmulqNS16(r *arm.Int16X8, v0 *arm.Int16X8, v1 *arm.Int16)

// Vector multiply by scalar
//
//go:linkname VmulNU32 VmulNU32
//go:noescape
func VmulNU32(r *arm.Uint32X2, v0 *arm.Uint32X2, v1 *arm.Uint32)

// Vector multiply by scalar
//
//go:linkname VmulNU16 VmulNU16
//go:noescape
func VmulNU16(r *arm.Uint16X4, v0 *arm.Uint16X4, v1 *arm.Uint16)

// Vector multiply by scalar
//
//go:linkname VmulNF32 VmulNF32
//go:noescape
func VmulNF32(r *arm.Float32X2, v0 *arm.Float32X2, v1 *arm.Float32)

// Vector multiply by scalar
//
//go:linkname VmulNS32 VmulNS32
//go:noescape
func VmulNS32(r *arm.Int32X2, v0 *arm.Int32X2, v1 *arm.Int32)

// Vector multiply by scalar
//
//go:linkname VmulNS16 VmulNS16
//go:noescape
func VmulNS16(r *arm.Int16X4, v0 *arm.Int16X4, v1 *arm.Int16)

// Polynomial Multiply Long. This instruction multiplies corresponding elements in the lower or upper half of the vectors of the two source SIMD&FP registers, places the results in a vector, and writes the vector to the destination SIMD&FP register. The destination vector elements are twice as long as the elements that are multiplied.
//
//go:linkname VmullP8 VmullP8
//go:noescape
func VmullP8(r *arm.Poly16X8, v0 *arm.Poly8X8, v1 *arm.Poly8X8)

// Unsigned Multiply long (vector). This instruction multiplies corresponding vector elements in the lower or upper half of the two source SIMD&FP registers, places the result in a vector, and writes the vector to the destination SIMD&FP register. The destination vector elements are twice as long as the elements that are multiplied. All the values in this instruction are unsigned integer values.
//
//go:linkname VmullU8 VmullU8
//go:noescape
func VmullU8(r *arm.Uint16X8, v0 *arm.Uint8X8, v1 *arm.Uint8X8)

// Unsigned Multiply long (vector). This instruction multiplies corresponding vector elements in the lower or upper half of the two source SIMD&FP registers, places the result in a vector, and writes the vector to the destination SIMD&FP register. The destination vector elements are twice as long as the elements that are multiplied. All the values in this instruction are unsigned integer values.
//
//go:linkname VmullU32 VmullU32
//go:noescape
func VmullU32(r *arm.Uint64X2, v0 *arm.Uint32X2, v1 *arm.Uint32X2)

// Unsigned Multiply long (vector). This instruction multiplies corresponding vector elements in the lower or upper half of the two source SIMD&FP registers, places the result in a vector, and writes the vector to the destination SIMD&FP register. The destination vector elements are twice as long as the elements that are multiplied. All the values in this instruction are unsigned integer values.
//
//go:linkname VmullU16 VmullU16
//go:noescape
func VmullU16(r *arm.Uint32X4, v0 *arm.Uint16X4, v1 *arm.Uint16X4)

// Signed Multiply Long (vector). This instruction multiplies corresponding signed integer values in the lower or upper half of the vectors of the two source SIMD&FP registers, places the results in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VmullS8 VmullS8
//go:noescape
func VmullS8(r *arm.Int16X8, v0 *arm.Int8X8, v1 *arm.Int8X8)

// Signed Multiply Long (vector). This instruction multiplies corresponding signed integer values in the lower or upper half of the vectors of the two source SIMD&FP registers, places the results in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VmullS32 VmullS32
//go:noescape
func VmullS32(r *arm.Int64X2, v0 *arm.Int32X2, v1 *arm.Int32X2)

// Signed Multiply Long (vector). This instruction multiplies corresponding signed integer values in the lower or upper half of the vectors of the two source SIMD&FP registers, places the results in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VmullS16 VmullS16
//go:noescape
func VmullS16(r *arm.Int32X4, v0 *arm.Int16X4, v1 *arm.Int16X4)

// Vector long multiply with scalar
//
//go:linkname VmullNU32 VmullNU32
//go:noescape
func VmullNU32(r *arm.Uint64X2, v0 *arm.Uint32X2, v1 *arm.Uint32)

// Vector long multiply with scalar
//
//go:linkname VmullNU16 VmullNU16
//go:noescape
func VmullNU16(r *arm.Uint32X4, v0 *arm.Uint16X4, v1 *arm.Uint16)

// Vector long multiply with scalar
//
//go:linkname VmullNS32 VmullNS32
//go:noescape
func VmullNS32(r *arm.Int64X2, v0 *arm.Int32X2, v1 *arm.Int32)

// Vector long multiply with scalar
//
//go:linkname VmullNS16 VmullNS16
//go:noescape
func VmullNS16(r *arm.Int32X4, v0 *arm.Int16X4, v1 *arm.Int16)

// Bitwise NOT (vector). This instruction reads each vector element from the source SIMD&FP register, places the inverse of each value into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VmvnP8 VmvnP8
//go:noescape
func VmvnP8(r *arm.Poly8X8, v0 *arm.Poly8X8)

// Bitwise NOT (vector). This instruction reads each vector element from the source SIMD&FP register, places the inverse of each value into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VmvnqP8 VmvnqP8
//go:noescape
func VmvnqP8(r *arm.Poly8X16, v0 *arm.Poly8X16)

// Bitwise NOT (vector). This instruction reads each vector element from the source SIMD&FP register, places the inverse of each value into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VmvnqU8 VmvnqU8
//go:noescape
func VmvnqU8(r *arm.Uint8X16, v0 *arm.Uint8X16)

// Bitwise NOT (vector). This instruction reads each vector element from the source SIMD&FP register, places the inverse of each value into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VmvnqU32 VmvnqU32
//go:noescape
func VmvnqU32(r *arm.Uint32X4, v0 *arm.Uint32X4)

// Bitwise NOT (vector). This instruction reads each vector element from the source SIMD&FP register, places the inverse of each value into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VmvnqU16 VmvnqU16
//go:noescape
func VmvnqU16(r *arm.Uint16X8, v0 *arm.Uint16X8)

// Bitwise NOT (vector). This instruction reads each vector element from the source SIMD&FP register, places the inverse of each value into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VmvnqS8 VmvnqS8
//go:noescape
func VmvnqS8(r *arm.Int8X16, v0 *arm.Int8X16)

// Bitwise NOT (vector). This instruction reads each vector element from the source SIMD&FP register, places the inverse of each value into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VmvnqS32 VmvnqS32
//go:noescape
func VmvnqS32(r *arm.Int32X4, v0 *arm.Int32X4)

// Bitwise NOT (vector). This instruction reads each vector element from the source SIMD&FP register, places the inverse of each value into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VmvnqS16 VmvnqS16
//go:noescape
func VmvnqS16(r *arm.Int16X8, v0 *arm.Int16X8)

// Bitwise NOT (vector). This instruction reads each vector element from the source SIMD&FP register, places the inverse of each value into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VmvnU8 VmvnU8
//go:noescape
func VmvnU8(r *arm.Uint8X8, v0 *arm.Uint8X8)

// Bitwise NOT (vector). This instruction reads each vector element from the source SIMD&FP register, places the inverse of each value into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VmvnU32 VmvnU32
//go:noescape
func VmvnU32(r *arm.Uint32X2, v0 *arm.Uint32X2)

// Bitwise NOT (vector). This instruction reads each vector element from the source SIMD&FP register, places the inverse of each value into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VmvnU16 VmvnU16
//go:noescape
func VmvnU16(r *arm.Uint16X4, v0 *arm.Uint16X4)

// Bitwise NOT (vector). This instruction reads each vector element from the source SIMD&FP register, places the inverse of each value into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VmvnS8 VmvnS8
//go:noescape
func VmvnS8(r *arm.Int8X8, v0 *arm.Int8X8)

// Bitwise NOT (vector). This instruction reads each vector element from the source SIMD&FP register, places the inverse of each value into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VmvnS32 VmvnS32
//go:noescape
func VmvnS32(r *arm.Int32X2, v0 *arm.Int32X2)

// Bitwise NOT (vector). This instruction reads each vector element from the source SIMD&FP register, places the inverse of each value into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VmvnS16 VmvnS16
//go:noescape
func VmvnS16(r *arm.Int16X4, v0 *arm.Int16X4)

// Negate (vector). This instruction reads each vector element from the source SIMD&FP register, negates each value, puts the result into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VnegqS8 VnegqS8
//go:noescape
func VnegqS8(r *arm.Int8X16, v0 *arm.Int8X16)

// Floating-point Negate (vector). This instruction negates the value of each vector element in the source SIMD&FP register, writes the result to a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VnegqF32 VnegqF32
//go:noescape
func VnegqF32(r *arm.Float32X4, v0 *arm.Float32X4)

// Negate (vector). This instruction reads each vector element from the source SIMD&FP register, negates each value, puts the result into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VnegqS32 VnegqS32
//go:noescape
func VnegqS32(r *arm.Int32X4, v0 *arm.Int32X4)

// Negate (vector). This instruction reads each vector element from the source SIMD&FP register, negates each value, puts the result into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VnegqS16 VnegqS16
//go:noescape
func VnegqS16(r *arm.Int16X8, v0 *arm.Int16X8)

// Negate (vector). This instruction reads each vector element from the source SIMD&FP register, negates each value, puts the result into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VnegS8 VnegS8
//go:noescape
func VnegS8(r *arm.Int8X8, v0 *arm.Int8X8)

// Floating-point Negate (vector). This instruction negates the value of each vector element in the source SIMD&FP register, writes the result to a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VnegF32 VnegF32
//go:noescape
func VnegF32(r *arm.Float32X2, v0 *arm.Float32X2)

// Negate (vector). This instruction reads each vector element from the source SIMD&FP register, negates each value, puts the result into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VnegS32 VnegS32
//go:noescape
func VnegS32(r *arm.Int32X2, v0 *arm.Int32X2)

// Negate (vector). This instruction reads each vector element from the source SIMD&FP register, negates each value, puts the result into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VnegS16 VnegS16
//go:noescape
func VnegS16(r *arm.Int16X4, v0 *arm.Int16X4)

// Bitwise inclusive OR NOT (vector). This instruction performs a bitwise OR NOT between the two source SIMD&FP registers, and writes the result to the destination SIMD&FP register.
//
//go:linkname VornqU8 VornqU8
//go:noescape
func VornqU8(r *arm.Uint8X16, v0 *arm.Uint8X16, v1 *arm.Uint8X16)

// Bitwise inclusive OR NOT (vector). This instruction performs a bitwise OR NOT between the two source SIMD&FP registers, and writes the result to the destination SIMD&FP register.
//
//go:linkname VornqU32 VornqU32
//go:noescape
func VornqU32(r *arm.Uint32X4, v0 *arm.Uint32X4, v1 *arm.Uint32X4)

// Bitwise inclusive OR NOT (vector). This instruction performs a bitwise OR NOT between the two source SIMD&FP registers, and writes the result to the destination SIMD&FP register.
//
//go:linkname VornqU64 VornqU64
//go:noescape
func VornqU64(r *arm.Uint64X2, v0 *arm.Uint64X2, v1 *arm.Uint64X2)

// Bitwise inclusive OR NOT (vector). This instruction performs a bitwise OR NOT between the two source SIMD&FP registers, and writes the result to the destination SIMD&FP register.
//
//go:linkname VornqU16 VornqU16
//go:noescape
func VornqU16(r *arm.Uint16X8, v0 *arm.Uint16X8, v1 *arm.Uint16X8)

// Bitwise inclusive OR NOT (vector). This instruction performs a bitwise OR NOT between the two source SIMD&FP registers, and writes the result to the destination SIMD&FP register.
//
//go:linkname VornqS8 VornqS8
//go:noescape
func VornqS8(r *arm.Int8X16, v0 *arm.Int8X16, v1 *arm.Int8X16)

// Bitwise inclusive OR NOT (vector). This instruction performs a bitwise OR NOT between the two source SIMD&FP registers, and writes the result to the destination SIMD&FP register.
//
//go:linkname VornqS32 VornqS32
//go:noescape
func VornqS32(r *arm.Int32X4, v0 *arm.Int32X4, v1 *arm.Int32X4)

// Bitwise inclusive OR NOT (vector). This instruction performs a bitwise OR NOT between the two source SIMD&FP registers, and writes the result to the destination SIMD&FP register.
//
//go:linkname VornqS64 VornqS64
//go:noescape
func VornqS64(r *arm.Int64X2, v0 *arm.Int64X2, v1 *arm.Int64X2)

// Bitwise inclusive OR NOT (vector). This instruction performs a bitwise OR NOT between the two source SIMD&FP registers, and writes the result to the destination SIMD&FP register.
//
//go:linkname VornqS16 VornqS16
//go:noescape
func VornqS16(r *arm.Int16X8, v0 *arm.Int16X8, v1 *arm.Int16X8)

// Bitwise inclusive OR NOT (vector). This instruction performs a bitwise OR NOT between the two source SIMD&FP registers, and writes the result to the destination SIMD&FP register.
//
//go:linkname VornU8 VornU8
//go:noescape
func VornU8(r *arm.Uint8X8, v0 *arm.Uint8X8, v1 *arm.Uint8X8)

// Bitwise inclusive OR NOT (vector). This instruction performs a bitwise OR NOT between the two source SIMD&FP registers, and writes the result to the destination SIMD&FP register.
//
//go:linkname VornU32 VornU32
//go:noescape
func VornU32(r *arm.Uint32X2, v0 *arm.Uint32X2, v1 *arm.Uint32X2)

// Bitwise inclusive OR NOT (vector). This instruction performs a bitwise OR NOT between the two source SIMD&FP registers, and writes the result to the destination SIMD&FP register.
//
//go:linkname VornU64 VornU64
//go:noescape
func VornU64(r *arm.Uint64X1, v0 *arm.Uint64X1, v1 *arm.Uint64X1)

// Bitwise inclusive OR NOT (vector). This instruction performs a bitwise OR NOT between the two source SIMD&FP registers, and writes the result to the destination SIMD&FP register.
//
//go:linkname VornU16 VornU16
//go:noescape
func VornU16(r *arm.Uint16X4, v0 *arm.Uint16X4, v1 *arm.Uint16X4)

// Bitwise inclusive OR NOT (vector). This instruction performs a bitwise OR NOT between the two source SIMD&FP registers, and writes the result to the destination SIMD&FP register.
//
//go:linkname VornS8 VornS8
//go:noescape
func VornS8(r *arm.Int8X8, v0 *arm.Int8X8, v1 *arm.Int8X8)

// Bitwise inclusive OR NOT (vector). This instruction performs a bitwise OR NOT between the two source SIMD&FP registers, and writes the result to the destination SIMD&FP register.
//
//go:linkname VornS32 VornS32
//go:noescape
func VornS32(r *arm.Int32X2, v0 *arm.Int32X2, v1 *arm.Int32X2)

// Bitwise inclusive OR NOT (vector). This instruction performs a bitwise OR NOT between the two source SIMD&FP registers, and writes the result to the destination SIMD&FP register.
//
//go:linkname VornS64 VornS64
//go:noescape
func VornS64(r *arm.Int64X1, v0 *arm.Int64X1, v1 *arm.Int64X1)

// Bitwise inclusive OR NOT (vector). This instruction performs a bitwise OR NOT between the two source SIMD&FP registers, and writes the result to the destination SIMD&FP register.
//
//go:linkname VornS16 VornS16
//go:noescape
func VornS16(r *arm.Int16X4, v0 *arm.Int16X4, v1 *arm.Int16X4)

// Bitwise inclusive OR (vector, register). This instruction performs a bitwise OR between the two source SIMD&FP registers, and writes the result to the destination SIMD&FP register.
//
//go:linkname VorrqU8 VorrqU8
//go:noescape
func VorrqU8(r *arm.Uint8X16, v0 *arm.Uint8X16, v1 *arm.Uint8X16)

// Bitwise inclusive OR (vector, register). This instruction performs a bitwise OR between the two source SIMD&FP registers, and writes the result to the destination SIMD&FP register.
//
//go:linkname VorrqU32 VorrqU32
//go:noescape
func VorrqU32(r *arm.Uint32X4, v0 *arm.Uint32X4, v1 *arm.Uint32X4)

// Bitwise inclusive OR (vector, register). This instruction performs a bitwise OR between the two source SIMD&FP registers, and writes the result to the destination SIMD&FP register.
//
//go:linkname VorrqU64 VorrqU64
//go:noescape
func VorrqU64(r *arm.Uint64X2, v0 *arm.Uint64X2, v1 *arm.Uint64X2)

// Bitwise inclusive OR (vector, register). This instruction performs a bitwise OR between the two source SIMD&FP registers, and writes the result to the destination SIMD&FP register.
//
//go:linkname VorrqU16 VorrqU16
//go:noescape
func VorrqU16(r *arm.Uint16X8, v0 *arm.Uint16X8, v1 *arm.Uint16X8)

// Bitwise inclusive OR (vector, register). This instruction performs a bitwise OR between the two source SIMD&FP registers, and writes the result to the destination SIMD&FP register.
//
//go:linkname VorrqS8 VorrqS8
//go:noescape
func VorrqS8(r *arm.Int8X16, v0 *arm.Int8X16, v1 *arm.Int8X16)

// Bitwise inclusive OR (vector, register). This instruction performs a bitwise OR between the two source SIMD&FP registers, and writes the result to the destination SIMD&FP register.
//
//go:linkname VorrqS32 VorrqS32
//go:noescape
func VorrqS32(r *arm.Int32X4, v0 *arm.Int32X4, v1 *arm.Int32X4)

// Bitwise inclusive OR (vector, register). This instruction performs a bitwise OR between the two source SIMD&FP registers, and writes the result to the destination SIMD&FP register.
//
//go:linkname VorrqS64 VorrqS64
//go:noescape
func VorrqS64(r *arm.Int64X2, v0 *arm.Int64X2, v1 *arm.Int64X2)

// Bitwise inclusive OR (vector, register). This instruction performs a bitwise OR between the two source SIMD&FP registers, and writes the result to the destination SIMD&FP register.
//
//go:linkname VorrqS16 VorrqS16
//go:noescape
func VorrqS16(r *arm.Int16X8, v0 *arm.Int16X8, v1 *arm.Int16X8)

// Bitwise inclusive OR (vector, register). This instruction performs a bitwise OR between the two source SIMD&FP registers, and writes the result to the destination SIMD&FP register.
//
//go:linkname VorrU8 VorrU8
//go:noescape
func VorrU8(r *arm.Uint8X8, v0 *arm.Uint8X8, v1 *arm.Uint8X8)

// Bitwise inclusive OR (vector, register). This instruction performs a bitwise OR between the two source SIMD&FP registers, and writes the result to the destination SIMD&FP register.
//
//go:linkname VorrU32 VorrU32
//go:noescape
func VorrU32(r *arm.Uint32X2, v0 *arm.Uint32X2, v1 *arm.Uint32X2)

// Bitwise inclusive OR (vector, register). This instruction performs a bitwise OR between the two source SIMD&FP registers, and writes the result to the destination SIMD&FP register.
//
//go:linkname VorrU64 VorrU64
//go:noescape
func VorrU64(r *arm.Uint64X1, v0 *arm.Uint64X1, v1 *arm.Uint64X1)

// Bitwise inclusive OR (vector, register). This instruction performs a bitwise OR between the two source SIMD&FP registers, and writes the result to the destination SIMD&FP register.
//
//go:linkname VorrU16 VorrU16
//go:noescape
func VorrU16(r *arm.Uint16X4, v0 *arm.Uint16X4, v1 *arm.Uint16X4)

// Bitwise inclusive OR (vector, register). This instruction performs a bitwise OR between the two source SIMD&FP registers, and writes the result to the destination SIMD&FP register.
//
//go:linkname VorrS8 VorrS8
//go:noescape
func VorrS8(r *arm.Int8X8, v0 *arm.Int8X8, v1 *arm.Int8X8)

// Bitwise inclusive OR (vector, register). This instruction performs a bitwise OR between the two source SIMD&FP registers, and writes the result to the destination SIMD&FP register.
//
//go:linkname VorrS32 VorrS32
//go:noescape
func VorrS32(r *arm.Int32X2, v0 *arm.Int32X2, v1 *arm.Int32X2)

// Bitwise inclusive OR (vector, register). This instruction performs a bitwise OR between the two source SIMD&FP registers, and writes the result to the destination SIMD&FP register.
//
//go:linkname VorrS64 VorrS64
//go:noescape
func VorrS64(r *arm.Int64X1, v0 *arm.Int64X1, v1 *arm.Int64X1)

// Bitwise inclusive OR (vector, register). This instruction performs a bitwise OR between the two source SIMD&FP registers, and writes the result to the destination SIMD&FP register.
//
//go:linkname VorrS16 VorrS16
//go:noescape
func VorrS16(r *arm.Int16X4, v0 *arm.Int16X4, v1 *arm.Int16X4)

// Unsigned Add and Accumulate Long Pairwise. This instruction adds pairs of adjacent unsigned integer values from the vector in the source SIMD&FP register and accumulates the results with the vector elements of the destination SIMD&FP register. The destination vector elements are twice as long as the source vector elements.
//
//go:linkname VpadalqU8 VpadalqU8
//go:noescape
func VpadalqU8(r *arm.Uint16X8, v0 *arm.Uint16X8, v1 *arm.Uint8X16)

// Unsigned Add and Accumulate Long Pairwise. This instruction adds pairs of adjacent unsigned integer values from the vector in the source SIMD&FP register and accumulates the results with the vector elements of the destination SIMD&FP register. The destination vector elements are twice as long as the source vector elements.
//
//go:linkname VpadalqU32 VpadalqU32
//go:noescape
func VpadalqU32(r *arm.Uint64X2, v0 *arm.Uint64X2, v1 *arm.Uint32X4)

// Unsigned Add and Accumulate Long Pairwise. This instruction adds pairs of adjacent unsigned integer values from the vector in the source SIMD&FP register and accumulates the results with the vector elements of the destination SIMD&FP register. The destination vector elements are twice as long as the source vector elements.
//
//go:linkname VpadalqU16 VpadalqU16
//go:noescape
func VpadalqU16(r *arm.Uint32X4, v0 *arm.Uint32X4, v1 *arm.Uint16X8)

// Signed Add and Accumulate Long Pairwise. This instruction adds pairs of adjacent signed integer values from the vector in the source SIMD&FP register and accumulates the results into the vector elements of the destination SIMD&FP register. The destination vector elements are twice as long as the source vector elements.
//
//go:linkname VpadalqS8 VpadalqS8
//go:noescape
func VpadalqS8(r *arm.Int16X8, v0 *arm.Int16X8, v1 *arm.Int8X16)

// Signed Add and Accumulate Long Pairwise. This instruction adds pairs of adjacent signed integer values from the vector in the source SIMD&FP register and accumulates the results into the vector elements of the destination SIMD&FP register. The destination vector elements are twice as long as the source vector elements.
//
//go:linkname VpadalqS32 VpadalqS32
//go:noescape
func VpadalqS32(r *arm.Int64X2, v0 *arm.Int64X2, v1 *arm.Int32X4)

// Signed Add and Accumulate Long Pairwise. This instruction adds pairs of adjacent signed integer values from the vector in the source SIMD&FP register and accumulates the results into the vector elements of the destination SIMD&FP register. The destination vector elements are twice as long as the source vector elements.
//
//go:linkname VpadalqS16 VpadalqS16
//go:noescape
func VpadalqS16(r *arm.Int32X4, v0 *arm.Int32X4, v1 *arm.Int16X8)

// Unsigned Add and Accumulate Long Pairwise. This instruction adds pairs of adjacent unsigned integer values from the vector in the source SIMD&FP register and accumulates the results with the vector elements of the destination SIMD&FP register. The destination vector elements are twice as long as the source vector elements.
//
//go:linkname VpadalU8 VpadalU8
//go:noescape
func VpadalU8(r *arm.Uint16X4, v0 *arm.Uint16X4, v1 *arm.Uint8X8)

// Unsigned Add and Accumulate Long Pairwise. This instruction adds pairs of adjacent unsigned integer values from the vector in the source SIMD&FP register and accumulates the results with the vector elements of the destination SIMD&FP register. The destination vector elements are twice as long as the source vector elements.
//
//go:linkname VpadalU32 VpadalU32
//go:noescape
func VpadalU32(r *arm.Uint64X1, v0 *arm.Uint64X1, v1 *arm.Uint32X2)

// Unsigned Add and Accumulate Long Pairwise. This instruction adds pairs of adjacent unsigned integer values from the vector in the source SIMD&FP register and accumulates the results with the vector elements of the destination SIMD&FP register. The destination vector elements are twice as long as the source vector elements.
//
//go:linkname VpadalU16 VpadalU16
//go:noescape
func VpadalU16(r *arm.Uint32X2, v0 *arm.Uint32X2, v1 *arm.Uint16X4)

// Signed Add and Accumulate Long Pairwise. This instruction adds pairs of adjacent signed integer values from the vector in the source SIMD&FP register and accumulates the results into the vector elements of the destination SIMD&FP register. The destination vector elements are twice as long as the source vector elements.
//
//go:linkname VpadalS8 VpadalS8
//go:noescape
func VpadalS8(r *arm.Int16X4, v0 *arm.Int16X4, v1 *arm.Int8X8)

// Signed Add and Accumulate Long Pairwise. This instruction adds pairs of adjacent signed integer values from the vector in the source SIMD&FP register and accumulates the results into the vector elements of the destination SIMD&FP register. The destination vector elements are twice as long as the source vector elements.
//
//go:linkname VpadalS32 VpadalS32
//go:noescape
func VpadalS32(r *arm.Int64X1, v0 *arm.Int64X1, v1 *arm.Int32X2)

// Signed Add and Accumulate Long Pairwise. This instruction adds pairs of adjacent signed integer values from the vector in the source SIMD&FP register and accumulates the results into the vector elements of the destination SIMD&FP register. The destination vector elements are twice as long as the source vector elements.
//
//go:linkname VpadalS16 VpadalS16
//go:noescape
func VpadalS16(r *arm.Int32X2, v0 *arm.Int32X2, v1 *arm.Int16X4)

// Add Pairwise (vector). This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements from the concatenated vector, adds each pair of values together, places the result into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VpaddU8 VpaddU8
//go:noescape
func VpaddU8(r *arm.Uint8X8, v0 *arm.Uint8X8, v1 *arm.Uint8X8)

// Add Pairwise (vector). This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements from the concatenated vector, adds each pair of values together, places the result into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VpaddU32 VpaddU32
//go:noescape
func VpaddU32(r *arm.Uint32X2, v0 *arm.Uint32X2, v1 *arm.Uint32X2)

// Add Pairwise (vector). This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements from the concatenated vector, adds each pair of values together, places the result into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VpaddU16 VpaddU16
//go:noescape
func VpaddU16(r *arm.Uint16X4, v0 *arm.Uint16X4, v1 *arm.Uint16X4)

// Add Pairwise (vector). This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements from the concatenated vector, adds each pair of values together, places the result into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VpaddS8 VpaddS8
//go:noescape
func VpaddS8(r *arm.Int8X8, v0 *arm.Int8X8, v1 *arm.Int8X8)

// Floating-point Add Pairwise (vector). This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements from the concatenated vector, adds each pair of values together, places the result into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are floating-point values.
//
//go:linkname VpaddF32 VpaddF32
//go:noescape
func VpaddF32(r *arm.Float32X2, v0 *arm.Float32X2, v1 *arm.Float32X2)

// Add Pairwise (vector). This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements from the concatenated vector, adds each pair of values together, places the result into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VpaddS32 VpaddS32
//go:noescape
func VpaddS32(r *arm.Int32X2, v0 *arm.Int32X2, v1 *arm.Int32X2)

// Add Pairwise (vector). This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements from the concatenated vector, adds each pair of values together, places the result into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VpaddS16 VpaddS16
//go:noescape
func VpaddS16(r *arm.Int16X4, v0 *arm.Int16X4, v1 *arm.Int16X4)

// Unsigned Add Long Pairwise. This instruction adds pairs of adjacent unsigned integer values from the vector in the source SIMD&FP register, places the result into a vector, and writes the vector to the destination SIMD&FP register. The destination vector elements are twice as long as the source vector elements.
//
//go:linkname VpaddlqU8 VpaddlqU8
//go:noescape
func VpaddlqU8(r *arm.Uint16X8, v0 *arm.Uint8X16)

// Unsigned Add Long Pairwise. This instruction adds pairs of adjacent unsigned integer values from the vector in the source SIMD&FP register, places the result into a vector, and writes the vector to the destination SIMD&FP register. The destination vector elements are twice as long as the source vector elements.
//
//go:linkname VpaddlqU32 VpaddlqU32
//go:noescape
func VpaddlqU32(r *arm.Uint64X2, v0 *arm.Uint32X4)

// Unsigned Add Long Pairwise. This instruction adds pairs of adjacent unsigned integer values from the vector in the source SIMD&FP register, places the result into a vector, and writes the vector to the destination SIMD&FP register. The destination vector elements are twice as long as the source vector elements.
//
//go:linkname VpaddlqU16 VpaddlqU16
//go:noescape
func VpaddlqU16(r *arm.Uint32X4, v0 *arm.Uint16X8)

// Signed Add Long Pairwise. This instruction adds pairs of adjacent signed integer values from the vector in the source SIMD&FP register, places the result into a vector, and writes the vector to the destination SIMD&FP register. The destination vector elements are twice as long as the source vector elements.
//
//go:linkname VpaddlqS8 VpaddlqS8
//go:noescape
func VpaddlqS8(r *arm.Int16X8, v0 *arm.Int8X16)

// Signed Add Long Pairwise. This instruction adds pairs of adjacent signed integer values from the vector in the source SIMD&FP register, places the result into a vector, and writes the vector to the destination SIMD&FP register. The destination vector elements are twice as long as the source vector elements.
//
//go:linkname VpaddlqS32 VpaddlqS32
//go:noescape
func VpaddlqS32(r *arm.Int64X2, v0 *arm.Int32X4)

// Signed Add Long Pairwise. This instruction adds pairs of adjacent signed integer values from the vector in the source SIMD&FP register, places the result into a vector, and writes the vector to the destination SIMD&FP register. The destination vector elements are twice as long as the source vector elements.
//
//go:linkname VpaddlqS16 VpaddlqS16
//go:noescape
func VpaddlqS16(r *arm.Int32X4, v0 *arm.Int16X8)

// Unsigned Add Long Pairwise. This instruction adds pairs of adjacent unsigned integer values from the vector in the source SIMD&FP register, places the result into a vector, and writes the vector to the destination SIMD&FP register. The destination vector elements are twice as long as the source vector elements.
//
//go:linkname VpaddlU8 VpaddlU8
//go:noescape
func VpaddlU8(r *arm.Uint16X4, v0 *arm.Uint8X8)

// Unsigned Add Long Pairwise. This instruction adds pairs of adjacent unsigned integer values from the vector in the source SIMD&FP register, places the result into a vector, and writes the vector to the destination SIMD&FP register. The destination vector elements are twice as long as the source vector elements.
//
//go:linkname VpaddlU32 VpaddlU32
//go:noescape
func VpaddlU32(r *arm.Uint64X1, v0 *arm.Uint32X2)

// Unsigned Add Long Pairwise. This instruction adds pairs of adjacent unsigned integer values from the vector in the source SIMD&FP register, places the result into a vector, and writes the vector to the destination SIMD&FP register. The destination vector elements are twice as long as the source vector elements.
//
//go:linkname VpaddlU16 VpaddlU16
//go:noescape
func VpaddlU16(r *arm.Uint32X2, v0 *arm.Uint16X4)

// Signed Add Long Pairwise. This instruction adds pairs of adjacent signed integer values from the vector in the source SIMD&FP register, places the result into a vector, and writes the vector to the destination SIMD&FP register. The destination vector elements are twice as long as the source vector elements.
//
//go:linkname VpaddlS8 VpaddlS8
//go:noescape
func VpaddlS8(r *arm.Int16X4, v0 *arm.Int8X8)

// Signed Add Long Pairwise. This instruction adds pairs of adjacent signed integer values from the vector in the source SIMD&FP register, places the result into a vector, and writes the vector to the destination SIMD&FP register. The destination vector elements are twice as long as the source vector elements.
//
//go:linkname VpaddlS32 VpaddlS32
//go:noescape
func VpaddlS32(r *arm.Int64X1, v0 *arm.Int32X2)

// Signed Add Long Pairwise. This instruction adds pairs of adjacent signed integer values from the vector in the source SIMD&FP register, places the result into a vector, and writes the vector to the destination SIMD&FP register. The destination vector elements are twice as long as the source vector elements.
//
//go:linkname VpaddlS16 VpaddlS16
//go:noescape
func VpaddlS16(r *arm.Int32X2, v0 *arm.Int16X4)

// Unsigned Maximum Pairwise. This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements in the two source SIMD&FP registers, writes the largest of each pair of unsigned integer values into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VpmaxU8 VpmaxU8
//go:noescape
func VpmaxU8(r *arm.Uint8X8, v0 *arm.Uint8X8, v1 *arm.Uint8X8)

// Unsigned Maximum Pairwise. This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements in the two source SIMD&FP registers, writes the largest of each pair of unsigned integer values into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VpmaxU32 VpmaxU32
//go:noescape
func VpmaxU32(r *arm.Uint32X2, v0 *arm.Uint32X2, v1 *arm.Uint32X2)

// Unsigned Maximum Pairwise. This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements in the two source SIMD&FP registers, writes the largest of each pair of unsigned integer values into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VpmaxU16 VpmaxU16
//go:noescape
func VpmaxU16(r *arm.Uint16X4, v0 *arm.Uint16X4, v1 *arm.Uint16X4)

// Signed Maximum Pairwise. This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements in the two source SIMD&FP registers, writes the largest of each pair of signed integer values into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VpmaxS8 VpmaxS8
//go:noescape
func VpmaxS8(r *arm.Int8X8, v0 *arm.Int8X8, v1 *arm.Int8X8)

// Floating-point Maximum Pairwise (vector). This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements from the concatenated vector, writes the larger of each pair of values into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are floating-point values.
//
//go:linkname VpmaxF32 VpmaxF32
//go:noescape
func VpmaxF32(r *arm.Float32X2, v0 *arm.Float32X2, v1 *arm.Float32X2)

// Signed Maximum Pairwise. This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements in the two source SIMD&FP registers, writes the largest of each pair of signed integer values into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VpmaxS32 VpmaxS32
//go:noescape
func VpmaxS32(r *arm.Int32X2, v0 *arm.Int32X2, v1 *arm.Int32X2)

// Signed Maximum Pairwise. This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements in the two source SIMD&FP registers, writes the largest of each pair of signed integer values into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VpmaxS16 VpmaxS16
//go:noescape
func VpmaxS16(r *arm.Int16X4, v0 *arm.Int16X4, v1 *arm.Int16X4)

// Unsigned Minimum Pairwise. This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements in the two source SIMD&FP registers, writes the smallest of each pair of unsigned integer values into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VpminU8 VpminU8
//go:noescape
func VpminU8(r *arm.Uint8X8, v0 *arm.Uint8X8, v1 *arm.Uint8X8)

// Unsigned Minimum Pairwise. This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements in the two source SIMD&FP registers, writes the smallest of each pair of unsigned integer values into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VpminU32 VpminU32
//go:noescape
func VpminU32(r *arm.Uint32X2, v0 *arm.Uint32X2, v1 *arm.Uint32X2)

// Unsigned Minimum Pairwise. This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements in the two source SIMD&FP registers, writes the smallest of each pair of unsigned integer values into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VpminU16 VpminU16
//go:noescape
func VpminU16(r *arm.Uint16X4, v0 *arm.Uint16X4, v1 *arm.Uint16X4)

// Signed Minimum Pairwise. This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements in the two source SIMD&FP registers, writes the smallest of each pair of signed integer values into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VpminS8 VpminS8
//go:noescape
func VpminS8(r *arm.Int8X8, v0 *arm.Int8X8, v1 *arm.Int8X8)

// Floating-point Minimum Pairwise (vector). This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements from the concatenated vector, writes the smaller of each pair of values into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are floating-point values.
//
//go:linkname VpminF32 VpminF32
//go:noescape
func VpminF32(r *arm.Float32X2, v0 *arm.Float32X2, v1 *arm.Float32X2)

// Signed Minimum Pairwise. This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements in the two source SIMD&FP registers, writes the smallest of each pair of signed integer values into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VpminS32 VpminS32
//go:noescape
func VpminS32(r *arm.Int32X2, v0 *arm.Int32X2, v1 *arm.Int32X2)

// Signed Minimum Pairwise. This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements in the two source SIMD&FP registers, writes the smallest of each pair of signed integer values into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VpminS16 VpminS16
//go:noescape
func VpminS16(r *arm.Int16X4, v0 *arm.Int16X4, v1 *arm.Int16X4)

// Signed saturating Absolute value. This instruction reads each vector element from the source SIMD&FP register, puts the absolute value of the result into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are signed integer values.
//
//go:linkname VqabsqS8 VqabsqS8
//go:noescape
func VqabsqS8(r *arm.Int8X16, v0 *arm.Int8X16)

// Signed saturating Absolute value. This instruction reads each vector element from the source SIMD&FP register, puts the absolute value of the result into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are signed integer values.
//
//go:linkname VqabsqS32 VqabsqS32
//go:noescape
func VqabsqS32(r *arm.Int32X4, v0 *arm.Int32X4)

// Signed saturating Absolute value. This instruction reads each vector element from the source SIMD&FP register, puts the absolute value of the result into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are signed integer values.
//
//go:linkname VqabsqS16 VqabsqS16
//go:noescape
func VqabsqS16(r *arm.Int16X8, v0 *arm.Int16X8)

// Signed saturating Absolute value. This instruction reads each vector element from the source SIMD&FP register, puts the absolute value of the result into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are signed integer values.
//
//go:linkname VqabsS8 VqabsS8
//go:noescape
func VqabsS8(r *arm.Int8X8, v0 *arm.Int8X8)

// Signed saturating Absolute value. This instruction reads each vector element from the source SIMD&FP register, puts the absolute value of the result into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are signed integer values.
//
//go:linkname VqabsS32 VqabsS32
//go:noescape
func VqabsS32(r *arm.Int32X2, v0 *arm.Int32X2)

// Signed saturating Absolute value. This instruction reads each vector element from the source SIMD&FP register, puts the absolute value of the result into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are signed integer values.
//
//go:linkname VqabsS16 VqabsS16
//go:noescape
func VqabsS16(r *arm.Int16X4, v0 *arm.Int16X4)

// Unsigned saturating Add. This instruction adds the values of corresponding elements of the two source SIMD&FP registers, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VqaddqU8 VqaddqU8
//go:noescape
func VqaddqU8(r *arm.Uint8X16, v0 *arm.Uint8X16, v1 *arm.Uint8X16)

// Unsigned saturating Add. This instruction adds the values of corresponding elements of the two source SIMD&FP registers, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VqaddqU32 VqaddqU32
//go:noescape
func VqaddqU32(r *arm.Uint32X4, v0 *arm.Uint32X4, v1 *arm.Uint32X4)

// Unsigned saturating Add. This instruction adds the values of corresponding elements of the two source SIMD&FP registers, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VqaddqU64 VqaddqU64
//go:noescape
func VqaddqU64(r *arm.Uint64X2, v0 *arm.Uint64X2, v1 *arm.Uint64X2)

// Unsigned saturating Add. This instruction adds the values of corresponding elements of the two source SIMD&FP registers, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VqaddqU16 VqaddqU16
//go:noescape
func VqaddqU16(r *arm.Uint16X8, v0 *arm.Uint16X8, v1 *arm.Uint16X8)

// Signed saturating Add. This instruction adds the values of corresponding elements of the two source SIMD&FP registers, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VqaddqS8 VqaddqS8
//go:noescape
func VqaddqS8(r *arm.Int8X16, v0 *arm.Int8X16, v1 *arm.Int8X16)

// Signed saturating Add. This instruction adds the values of corresponding elements of the two source SIMD&FP registers, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VqaddqS32 VqaddqS32
//go:noescape
func VqaddqS32(r *arm.Int32X4, v0 *arm.Int32X4, v1 *arm.Int32X4)

// Signed saturating Add. This instruction adds the values of corresponding elements of the two source SIMD&FP registers, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VqaddqS64 VqaddqS64
//go:noescape
func VqaddqS64(r *arm.Int64X2, v0 *arm.Int64X2, v1 *arm.Int64X2)

// Signed saturating Add. This instruction adds the values of corresponding elements of the two source SIMD&FP registers, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VqaddqS16 VqaddqS16
//go:noescape
func VqaddqS16(r *arm.Int16X8, v0 *arm.Int16X8, v1 *arm.Int16X8)

// Unsigned saturating Add. This instruction adds the values of corresponding elements of the two source SIMD&FP registers, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VqaddU8 VqaddU8
//go:noescape
func VqaddU8(r *arm.Uint8X8, v0 *arm.Uint8X8, v1 *arm.Uint8X8)

// Unsigned saturating Add. This instruction adds the values of corresponding elements of the two source SIMD&FP registers, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VqaddU32 VqaddU32
//go:noescape
func VqaddU32(r *arm.Uint32X2, v0 *arm.Uint32X2, v1 *arm.Uint32X2)

// Unsigned saturating Add. This instruction adds the values of corresponding elements of the two source SIMD&FP registers, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VqaddU64 VqaddU64
//go:noescape
func VqaddU64(r *arm.Uint64X1, v0 *arm.Uint64X1, v1 *arm.Uint64X1)

// Unsigned saturating Add. This instruction adds the values of corresponding elements of the two source SIMD&FP registers, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VqaddU16 VqaddU16
//go:noescape
func VqaddU16(r *arm.Uint16X4, v0 *arm.Uint16X4, v1 *arm.Uint16X4)

// Signed saturating Add. This instruction adds the values of corresponding elements of the two source SIMD&FP registers, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VqaddS8 VqaddS8
//go:noescape
func VqaddS8(r *arm.Int8X8, v0 *arm.Int8X8, v1 *arm.Int8X8)

// Signed saturating Add. This instruction adds the values of corresponding elements of the two source SIMD&FP registers, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VqaddS32 VqaddS32
//go:noescape
func VqaddS32(r *arm.Int32X2, v0 *arm.Int32X2, v1 *arm.Int32X2)

// Signed saturating Add. This instruction adds the values of corresponding elements of the two source SIMD&FP registers, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VqaddS64 VqaddS64
//go:noescape
func VqaddS64(r *arm.Int64X1, v0 *arm.Int64X1, v1 *arm.Int64X1)

// Signed saturating Add. This instruction adds the values of corresponding elements of the two source SIMD&FP registers, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VqaddS16 VqaddS16
//go:noescape
func VqaddS16(r *arm.Int16X4, v0 *arm.Int16X4, v1 *arm.Int16X4)

// Signed saturating Doubling Multiply-Add Long. This instruction multiplies corresponding signed integer values in the lower or upper half of the vectors of the two source SIMD&FP registers, doubles the results, and accumulates the final results with the vector elements of the destination SIMD&FP register. The destination vector elements are twice as long as the elements that are multiplied.
//
//go:linkname VqdmlalS32 VqdmlalS32
//go:noescape
func VqdmlalS32(r *arm.Int64X2, v0 *arm.Int64X2, v1 *arm.Int32X2, v2 *arm.Int32X2)

// Signed saturating Doubling Multiply-Add Long. This instruction multiplies corresponding signed integer values in the lower or upper half of the vectors of the two source SIMD&FP registers, doubles the results, and accumulates the final results with the vector elements of the destination SIMD&FP register. The destination vector elements are twice as long as the elements that are multiplied.
//
//go:linkname VqdmlalS16 VqdmlalS16
//go:noescape
func VqdmlalS16(r *arm.Int32X4, v0 *arm.Int32X4, v1 *arm.Int16X4, v2 *arm.Int16X4)

// Vector widening saturating doubling multiply accumulate with scalar
//
//go:linkname VqdmlalNS32 VqdmlalNS32
//go:noescape
func VqdmlalNS32(r *arm.Int64X2, v0 *arm.Int64X2, v1 *arm.Int32X2, v2 *arm.Int32)

// Vector widening saturating doubling multiply accumulate with scalar
//
//go:linkname VqdmlalNS16 VqdmlalNS16
//go:noescape
func VqdmlalNS16(r *arm.Int32X4, v0 *arm.Int32X4, v1 *arm.Int16X4, v2 *arm.Int16)

// Signed saturating Doubling Multiply-Subtract Long. This instruction multiplies corresponding signed integer values in the lower or upper half of the vectors of the two source SIMD&FP registers, doubles the results, and subtracts the final results from the vector elements of the destination SIMD&FP register. The destination vector elements are twice as long as the elements that are multiplied.
//
//go:linkname VqdmlslS32 VqdmlslS32
//go:noescape
func VqdmlslS32(r *arm.Int64X2, v0 *arm.Int64X2, v1 *arm.Int32X2, v2 *arm.Int32X2)

// Signed saturating Doubling Multiply-Subtract Long. This instruction multiplies corresponding signed integer values in the lower or upper half of the vectors of the two source SIMD&FP registers, doubles the results, and subtracts the final results from the vector elements of the destination SIMD&FP register. The destination vector elements are twice as long as the elements that are multiplied.
//
//go:linkname VqdmlslS16 VqdmlslS16
//go:noescape
func VqdmlslS16(r *arm.Int32X4, v0 *arm.Int32X4, v1 *arm.Int16X4, v2 *arm.Int16X4)

// Vector widening saturating doubling multiply subtract with scalar
//
//go:linkname VqdmlslNS32 VqdmlslNS32
//go:noescape
func VqdmlslNS32(r *arm.Int64X2, v0 *arm.Int64X2, v1 *arm.Int32X2, v2 *arm.Int32)

// Vector widening saturating doubling multiply subtract with scalar
//
//go:linkname VqdmlslNS16 VqdmlslNS16
//go:noescape
func VqdmlslNS16(r *arm.Int32X4, v0 *arm.Int32X4, v1 *arm.Int16X4, v2 *arm.Int16)

// Signed saturating Doubling Multiply returning High half. This instruction multiplies the values of corresponding elements of the two source SIMD&FP registers, doubles the results, places the most significant half of the final results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VqdmulhqS32 VqdmulhqS32
//go:noescape
func VqdmulhqS32(r *arm.Int32X4, v0 *arm.Int32X4, v1 *arm.Int32X4)

// Signed saturating Doubling Multiply returning High half. This instruction multiplies the values of corresponding elements of the two source SIMD&FP registers, doubles the results, places the most significant half of the final results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VqdmulhqS16 VqdmulhqS16
//go:noescape
func VqdmulhqS16(r *arm.Int16X8, v0 *arm.Int16X8, v1 *arm.Int16X8)

// Signed saturating Doubling Multiply returning High half. This instruction multiplies the values of corresponding elements of the two source SIMD&FP registers, doubles the results, places the most significant half of the final results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VqdmulhS32 VqdmulhS32
//go:noescape
func VqdmulhS32(r *arm.Int32X2, v0 *arm.Int32X2, v1 *arm.Int32X2)

// Signed saturating Doubling Multiply returning High half. This instruction multiplies the values of corresponding elements of the two source SIMD&FP registers, doubles the results, places the most significant half of the final results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VqdmulhS16 VqdmulhS16
//go:noescape
func VqdmulhS16(r *arm.Int16X4, v0 *arm.Int16X4, v1 *arm.Int16X4)

// Vector saturating doubling multiply high with scalar
//
//go:linkname VqdmulhqNS32 VqdmulhqNS32
//go:noescape
func VqdmulhqNS32(r *arm.Int32X4, v0 *arm.Int32X4, v1 *arm.Int32)

// Vector saturating doubling multiply high with scalar
//
//go:linkname VqdmulhqNS16 VqdmulhqNS16
//go:noescape
func VqdmulhqNS16(r *arm.Int16X8, v0 *arm.Int16X8, v1 *arm.Int16)

// Vector saturating doubling multiply high with scalar
//
//go:linkname VqdmulhNS32 VqdmulhNS32
//go:noescape
func VqdmulhNS32(r *arm.Int32X2, v0 *arm.Int32X2, v1 *arm.Int32)

// Vector saturating doubling multiply high with scalar
//
//go:linkname VqdmulhNS16 VqdmulhNS16
//go:noescape
func VqdmulhNS16(r *arm.Int16X4, v0 *arm.Int16X4, v1 *arm.Int16)

// Signed saturating Doubling Multiply Long. This instruction multiplies corresponding vector elements in the lower or upper half of the two source SIMD&FP registers, doubles the results, places the final results in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VqdmullS32 VqdmullS32
//go:noescape
func VqdmullS32(r *arm.Int64X2, v0 *arm.Int32X2, v1 *arm.Int32X2)

// Signed saturating Doubling Multiply Long. This instruction multiplies corresponding vector elements in the lower or upper half of the two source SIMD&FP registers, doubles the results, places the final results in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VqdmullS16 VqdmullS16
//go:noescape
func VqdmullS16(r *arm.Int32X4, v0 *arm.Int16X4, v1 *arm.Int16X4)

// Vector saturating doubling long multiply with scalar
//
//go:linkname VqdmullNS32 VqdmullNS32
//go:noescape
func VqdmullNS32(r *arm.Int64X2, v0 *arm.Int32X2, v1 *arm.Int32)

// Vector saturating doubling long multiply with scalar
//
//go:linkname VqdmullNS16 VqdmullNS16
//go:noescape
func VqdmullNS16(r *arm.Int32X4, v0 *arm.Int16X4, v1 *arm.Int16)

// Unsigned saturating extract Narrow. This instruction reads each vector element from the source SIMD&FP register, saturates each value to half the original width, places the result into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are unsigned integer values.
//
//go:linkname VqmovnU32 VqmovnU32
//go:noescape
func VqmovnU32(r *arm.Uint16X4, v0 *arm.Uint32X4)

// Unsigned saturating extract Narrow. This instruction reads each vector element from the source SIMD&FP register, saturates each value to half the original width, places the result into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are unsigned integer values.
//
//go:linkname VqmovnU64 VqmovnU64
//go:noescape
func VqmovnU64(r *arm.Uint32X2, v0 *arm.Uint64X2)

// Unsigned saturating extract Narrow. This instruction reads each vector element from the source SIMD&FP register, saturates each value to half the original width, places the result into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are unsigned integer values.
//
//go:linkname VqmovnU16 VqmovnU16
//go:noescape
func VqmovnU16(r *arm.Uint8X8, v0 *arm.Uint16X8)

// Signed saturating extract Narrow. This instruction reads each vector element from the source SIMD&FP register, saturates the value to half the original width, places the result into a vector, and writes the vector to the lower or upper half of the destination SIMD&FP register. The destination vector elements are half as long as the source vector elements. All the values in this instruction are signed integer values.
//
//go:linkname VqmovnS32 VqmovnS32
//go:noescape
func VqmovnS32(r *arm.Int16X4, v0 *arm.Int32X4)

// Signed saturating extract Narrow. This instruction reads each vector element from the source SIMD&FP register, saturates the value to half the original width, places the result into a vector, and writes the vector to the lower or upper half of the destination SIMD&FP register. The destination vector elements are half as long as the source vector elements. All the values in this instruction are signed integer values.
//
//go:linkname VqmovnS64 VqmovnS64
//go:noescape
func VqmovnS64(r *arm.Int32X2, v0 *arm.Int64X2)

// Signed saturating extract Narrow. This instruction reads each vector element from the source SIMD&FP register, saturates the value to half the original width, places the result into a vector, and writes the vector to the lower or upper half of the destination SIMD&FP register. The destination vector elements are half as long as the source vector elements. All the values in this instruction are signed integer values.
//
//go:linkname VqmovnS16 VqmovnS16
//go:noescape
func VqmovnS16(r *arm.Int8X8, v0 *arm.Int16X8)

// Signed saturating extract Unsigned Narrow. This instruction reads each signed integer value in the vector of the source SIMD&FP register, saturates the value to an unsigned integer value that is half the original width, places the result into a vector, and writes the vector to the lower or upper half of the destination SIMD&FP register. The destination vector elements are half as long as the source vector elements.
//
//go:linkname VqmovunS32 VqmovunS32
//go:noescape
func VqmovunS32(r *arm.Uint16X4, v0 *arm.Int32X4)

// Signed saturating extract Unsigned Narrow. This instruction reads each signed integer value in the vector of the source SIMD&FP register, saturates the value to an unsigned integer value that is half the original width, places the result into a vector, and writes the vector to the lower or upper half of the destination SIMD&FP register. The destination vector elements are half as long as the source vector elements.
//
//go:linkname VqmovunS64 VqmovunS64
//go:noescape
func VqmovunS64(r *arm.Uint32X2, v0 *arm.Int64X2)

// Signed saturating extract Unsigned Narrow. This instruction reads each signed integer value in the vector of the source SIMD&FP register, saturates the value to an unsigned integer value that is half the original width, places the result into a vector, and writes the vector to the lower or upper half of the destination SIMD&FP register. The destination vector elements are half as long as the source vector elements.
//
//go:linkname VqmovunS16 VqmovunS16
//go:noescape
func VqmovunS16(r *arm.Uint8X8, v0 *arm.Int16X8)

// Signed saturating Negate. This instruction reads each vector element from the source SIMD&FP register, negates each value, places the result into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are signed integer values.
//
//go:linkname VqnegqS8 VqnegqS8
//go:noescape
func VqnegqS8(r *arm.Int8X16, v0 *arm.Int8X16)

// Signed saturating Negate. This instruction reads each vector element from the source SIMD&FP register, negates each value, places the result into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are signed integer values.
//
//go:linkname VqnegqS32 VqnegqS32
//go:noescape
func VqnegqS32(r *arm.Int32X4, v0 *arm.Int32X4)

// Signed saturating Negate. This instruction reads each vector element from the source SIMD&FP register, negates each value, places the result into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are signed integer values.
//
//go:linkname VqnegqS16 VqnegqS16
//go:noescape
func VqnegqS16(r *arm.Int16X8, v0 *arm.Int16X8)

// Signed saturating Negate. This instruction reads each vector element from the source SIMD&FP register, negates each value, places the result into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are signed integer values.
//
//go:linkname VqnegS8 VqnegS8
//go:noescape
func VqnegS8(r *arm.Int8X8, v0 *arm.Int8X8)

// Signed saturating Negate. This instruction reads each vector element from the source SIMD&FP register, negates each value, places the result into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are signed integer values.
//
//go:linkname VqnegS32 VqnegS32
//go:noescape
func VqnegS32(r *arm.Int32X2, v0 *arm.Int32X2)

// Signed saturating Negate. This instruction reads each vector element from the source SIMD&FP register, negates each value, places the result into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are signed integer values.
//
//go:linkname VqnegS16 VqnegS16
//go:noescape
func VqnegS16(r *arm.Int16X4, v0 *arm.Int16X4)

// Signed saturating Rounding Doubling Multiply returning High half. This instruction multiplies the values of corresponding elements of the two source SIMD&FP registers, doubles the results, places the most significant half of the final results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VqrdmulhqS32 VqrdmulhqS32
//go:noescape
func VqrdmulhqS32(r *arm.Int32X4, v0 *arm.Int32X4, v1 *arm.Int32X4)

// Signed saturating Rounding Doubling Multiply returning High half. This instruction multiplies the values of corresponding elements of the two source SIMD&FP registers, doubles the results, places the most significant half of the final results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VqrdmulhqS16 VqrdmulhqS16
//go:noescape
func VqrdmulhqS16(r *arm.Int16X8, v0 *arm.Int16X8, v1 *arm.Int16X8)

// Signed saturating Rounding Doubling Multiply returning High half. This instruction multiplies the values of corresponding elements of the two source SIMD&FP registers, doubles the results, places the most significant half of the final results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VqrdmulhS32 VqrdmulhS32
//go:noescape
func VqrdmulhS32(r *arm.Int32X2, v0 *arm.Int32X2, v1 *arm.Int32X2)

// Signed saturating Rounding Doubling Multiply returning High half. This instruction multiplies the values of corresponding elements of the two source SIMD&FP registers, doubles the results, places the most significant half of the final results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VqrdmulhS16 VqrdmulhS16
//go:noescape
func VqrdmulhS16(r *arm.Int16X4, v0 *arm.Int16X4, v1 *arm.Int16X4)

// Vector saturating rounding doubling multiply high with scalar
//
//go:linkname VqrdmulhqNS32 VqrdmulhqNS32
//go:noescape
func VqrdmulhqNS32(r *arm.Int32X4, v0 *arm.Int32X4, v1 *arm.Int32)

// Vector saturating rounding doubling multiply high with scalar
//
//go:linkname VqrdmulhqNS16 VqrdmulhqNS16
//go:noescape
func VqrdmulhqNS16(r *arm.Int16X8, v0 *arm.Int16X8, v1 *arm.Int16)

// Vector saturating rounding doubling multiply high with scalar
//
//go:linkname VqrdmulhNS32 VqrdmulhNS32
//go:noescape
func VqrdmulhNS32(r *arm.Int32X2, v0 *arm.Int32X2, v1 *arm.Int32)

// Vector saturating rounding doubling multiply high with scalar
//
//go:linkname VqrdmulhNS16 VqrdmulhNS16
//go:noescape
func VqrdmulhNS16(r *arm.Int16X4, v0 *arm.Int16X4, v1 *arm.Int16)

// Unsigned saturating Rounding Shift Left (register). This instruction takes each vector element of the first source SIMD&FP register, shifts the vector element by a value from the least significant byte of the corresponding vector element of the second source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VqrshlqU8 VqrshlqU8
//go:noescape
func VqrshlqU8(r *arm.Uint8X16, v0 *arm.Uint8X16, v1 *arm.Int8X16)

// Unsigned saturating Rounding Shift Left (register). This instruction takes each vector element of the first source SIMD&FP register, shifts the vector element by a value from the least significant byte of the corresponding vector element of the second source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VqrshlqU32 VqrshlqU32
//go:noescape
func VqrshlqU32(r *arm.Uint32X4, v0 *arm.Uint32X4, v1 *arm.Int32X4)

// Unsigned saturating Rounding Shift Left (register). This instruction takes each vector element of the first source SIMD&FP register, shifts the vector element by a value from the least significant byte of the corresponding vector element of the second source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VqrshlqU64 VqrshlqU64
//go:noescape
func VqrshlqU64(r *arm.Uint64X2, v0 *arm.Uint64X2, v1 *arm.Int64X2)

// Unsigned saturating Rounding Shift Left (register). This instruction takes each vector element of the first source SIMD&FP register, shifts the vector element by a value from the least significant byte of the corresponding vector element of the second source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VqrshlqU16 VqrshlqU16
//go:noescape
func VqrshlqU16(r *arm.Uint16X8, v0 *arm.Uint16X8, v1 *arm.Int16X8)

// Signed saturating Rounding Shift Left (register). This instruction takes each vector element in the first source SIMD&FP register, shifts it by a value from the least significant byte of the corresponding vector element of the second source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VqrshlqS8 VqrshlqS8
//go:noescape
func VqrshlqS8(r *arm.Int8X16, v0 *arm.Int8X16, v1 *arm.Int8X16)

// Signed saturating Rounding Shift Left (register). This instruction takes each vector element in the first source SIMD&FP register, shifts it by a value from the least significant byte of the corresponding vector element of the second source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VqrshlqS32 VqrshlqS32
//go:noescape
func VqrshlqS32(r *arm.Int32X4, v0 *arm.Int32X4, v1 *arm.Int32X4)

// Signed saturating Rounding Shift Left (register). This instruction takes each vector element in the first source SIMD&FP register, shifts it by a value from the least significant byte of the corresponding vector element of the second source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VqrshlqS64 VqrshlqS64
//go:noescape
func VqrshlqS64(r *arm.Int64X2, v0 *arm.Int64X2, v1 *arm.Int64X2)

// Signed saturating Rounding Shift Left (register). This instruction takes each vector element in the first source SIMD&FP register, shifts it by a value from the least significant byte of the corresponding vector element of the second source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VqrshlqS16 VqrshlqS16
//go:noescape
func VqrshlqS16(r *arm.Int16X8, v0 *arm.Int16X8, v1 *arm.Int16X8)

// Unsigned saturating Rounding Shift Left (register). This instruction takes each vector element of the first source SIMD&FP register, shifts the vector element by a value from the least significant byte of the corresponding vector element of the second source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VqrshlU8 VqrshlU8
//go:noescape
func VqrshlU8(r *arm.Uint8X8, v0 *arm.Uint8X8, v1 *arm.Int8X8)

// Unsigned saturating Rounding Shift Left (register). This instruction takes each vector element of the first source SIMD&FP register, shifts the vector element by a value from the least significant byte of the corresponding vector element of the second source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VqrshlU32 VqrshlU32
//go:noescape
func VqrshlU32(r *arm.Uint32X2, v0 *arm.Uint32X2, v1 *arm.Int32X2)

// Unsigned saturating Rounding Shift Left (register). This instruction takes each vector element of the first source SIMD&FP register, shifts the vector element by a value from the least significant byte of the corresponding vector element of the second source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VqrshlU64 VqrshlU64
//go:noescape
func VqrshlU64(r *arm.Uint64X1, v0 *arm.Uint64X1, v1 *arm.Int64X1)

// Unsigned saturating Rounding Shift Left (register). This instruction takes each vector element of the first source SIMD&FP register, shifts the vector element by a value from the least significant byte of the corresponding vector element of the second source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VqrshlU16 VqrshlU16
//go:noescape
func VqrshlU16(r *arm.Uint16X4, v0 *arm.Uint16X4, v1 *arm.Int16X4)

// Signed saturating Rounding Shift Left (register). This instruction takes each vector element in the first source SIMD&FP register, shifts it by a value from the least significant byte of the corresponding vector element of the second source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VqrshlS8 VqrshlS8
//go:noescape
func VqrshlS8(r *arm.Int8X8, v0 *arm.Int8X8, v1 *arm.Int8X8)

// Signed saturating Rounding Shift Left (register). This instruction takes each vector element in the first source SIMD&FP register, shifts it by a value from the least significant byte of the corresponding vector element of the second source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VqrshlS32 VqrshlS32
//go:noescape
func VqrshlS32(r *arm.Int32X2, v0 *arm.Int32X2, v1 *arm.Int32X2)

// Signed saturating Rounding Shift Left (register). This instruction takes each vector element in the first source SIMD&FP register, shifts it by a value from the least significant byte of the corresponding vector element of the second source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VqrshlS64 VqrshlS64
//go:noescape
func VqrshlS64(r *arm.Int64X1, v0 *arm.Int64X1, v1 *arm.Int64X1)

// Signed saturating Rounding Shift Left (register). This instruction takes each vector element in the first source SIMD&FP register, shifts it by a value from the least significant byte of the corresponding vector element of the second source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VqrshlS16 VqrshlS16
//go:noescape
func VqrshlS16(r *arm.Int16X4, v0 *arm.Int16X4, v1 *arm.Int16X4)

// Unsigned saturating Shift Left (register). This instruction takes each element in the vector of the first source SIMD&FP register, shifts the element by a value from the least significant byte of the corresponding element of the second source SIMD&FP register, places the results in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VqshlqU8 VqshlqU8
//go:noescape
func VqshlqU8(r *arm.Uint8X16, v0 *arm.Uint8X16, v1 *arm.Int8X16)

// Unsigned saturating Shift Left (register). This instruction takes each element in the vector of the first source SIMD&FP register, shifts the element by a value from the least significant byte of the corresponding element of the second source SIMD&FP register, places the results in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VqshlqU32 VqshlqU32
//go:noescape
func VqshlqU32(r *arm.Uint32X4, v0 *arm.Uint32X4, v1 *arm.Int32X4)

// Unsigned saturating Shift Left (register). This instruction takes each element in the vector of the first source SIMD&FP register, shifts the element by a value from the least significant byte of the corresponding element of the second source SIMD&FP register, places the results in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VqshlqU64 VqshlqU64
//go:noescape
func VqshlqU64(r *arm.Uint64X2, v0 *arm.Uint64X2, v1 *arm.Int64X2)

// Unsigned saturating Shift Left (register). This instruction takes each element in the vector of the first source SIMD&FP register, shifts the element by a value from the least significant byte of the corresponding element of the second source SIMD&FP register, places the results in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VqshlqU16 VqshlqU16
//go:noescape
func VqshlqU16(r *arm.Uint16X8, v0 *arm.Uint16X8, v1 *arm.Int16X8)

// Signed saturating Shift Left (register). This instruction takes each element in the vector of the first source SIMD&FP register, shifts each element by a value from the least significant byte of the corresponding element of the second source SIMD&FP register, places the results in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VqshlqS8 VqshlqS8
//go:noescape
func VqshlqS8(r *arm.Int8X16, v0 *arm.Int8X16, v1 *arm.Int8X16)

// Signed saturating Shift Left (register). This instruction takes each element in the vector of the first source SIMD&FP register, shifts each element by a value from the least significant byte of the corresponding element of the second source SIMD&FP register, places the results in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VqshlqS32 VqshlqS32
//go:noescape
func VqshlqS32(r *arm.Int32X4, v0 *arm.Int32X4, v1 *arm.Int32X4)

// Signed saturating Shift Left (register). This instruction takes each element in the vector of the first source SIMD&FP register, shifts each element by a value from the least significant byte of the corresponding element of the second source SIMD&FP register, places the results in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VqshlqS64 VqshlqS64
//go:noescape
func VqshlqS64(r *arm.Int64X2, v0 *arm.Int64X2, v1 *arm.Int64X2)

// Signed saturating Shift Left (register). This instruction takes each element in the vector of the first source SIMD&FP register, shifts each element by a value from the least significant byte of the corresponding element of the second source SIMD&FP register, places the results in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VqshlqS16 VqshlqS16
//go:noescape
func VqshlqS16(r *arm.Int16X8, v0 *arm.Int16X8, v1 *arm.Int16X8)

// Unsigned saturating Shift Left (register). This instruction takes each element in the vector of the first source SIMD&FP register, shifts the element by a value from the least significant byte of the corresponding element of the second source SIMD&FP register, places the results in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VqshlU8 VqshlU8
//go:noescape
func VqshlU8(r *arm.Uint8X8, v0 *arm.Uint8X8, v1 *arm.Int8X8)

// Unsigned saturating Shift Left (register). This instruction takes each element in the vector of the first source SIMD&FP register, shifts the element by a value from the least significant byte of the corresponding element of the second source SIMD&FP register, places the results in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VqshlU32 VqshlU32
//go:noescape
func VqshlU32(r *arm.Uint32X2, v0 *arm.Uint32X2, v1 *arm.Int32X2)

// Unsigned saturating Shift Left (register). This instruction takes each element in the vector of the first source SIMD&FP register, shifts the element by a value from the least significant byte of the corresponding element of the second source SIMD&FP register, places the results in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VqshlU64 VqshlU64
//go:noescape
func VqshlU64(r *arm.Uint64X1, v0 *arm.Uint64X1, v1 *arm.Int64X1)

// Unsigned saturating Shift Left (register). This instruction takes each element in the vector of the first source SIMD&FP register, shifts the element by a value from the least significant byte of the corresponding element of the second source SIMD&FP register, places the results in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VqshlU16 VqshlU16
//go:noescape
func VqshlU16(r *arm.Uint16X4, v0 *arm.Uint16X4, v1 *arm.Int16X4)

// Signed saturating Shift Left (register). This instruction takes each element in the vector of the first source SIMD&FP register, shifts each element by a value from the least significant byte of the corresponding element of the second source SIMD&FP register, places the results in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VqshlS8 VqshlS8
//go:noescape
func VqshlS8(r *arm.Int8X8, v0 *arm.Int8X8, v1 *arm.Int8X8)

// Signed saturating Shift Left (register). This instruction takes each element in the vector of the first source SIMD&FP register, shifts each element by a value from the least significant byte of the corresponding element of the second source SIMD&FP register, places the results in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VqshlS32 VqshlS32
//go:noescape
func VqshlS32(r *arm.Int32X2, v0 *arm.Int32X2, v1 *arm.Int32X2)

// Signed saturating Shift Left (register). This instruction takes each element in the vector of the first source SIMD&FP register, shifts each element by a value from the least significant byte of the corresponding element of the second source SIMD&FP register, places the results in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VqshlS64 VqshlS64
//go:noescape
func VqshlS64(r *arm.Int64X1, v0 *arm.Int64X1, v1 *arm.Int64X1)

// Signed saturating Shift Left (register). This instruction takes each element in the vector of the first source SIMD&FP register, shifts each element by a value from the least significant byte of the corresponding element of the second source SIMD&FP register, places the results in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VqshlS16 VqshlS16
//go:noescape
func VqshlS16(r *arm.Int16X4, v0 *arm.Int16X4, v1 *arm.Int16X4)

// Unsigned saturating Subtract. This instruction subtracts the element values of the second source SIMD&FP register from the corresponding element values of the first source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VqsubqU8 VqsubqU8
//go:noescape
func VqsubqU8(r *arm.Uint8X16, v0 *arm.Uint8X16, v1 *arm.Uint8X16)

// Unsigned saturating Subtract. This instruction subtracts the element values of the second source SIMD&FP register from the corresponding element values of the first source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VqsubqU32 VqsubqU32
//go:noescape
func VqsubqU32(r *arm.Uint32X4, v0 *arm.Uint32X4, v1 *arm.Uint32X4)

// Unsigned saturating Subtract. This instruction subtracts the element values of the second source SIMD&FP register from the corresponding element values of the first source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VqsubqU64 VqsubqU64
//go:noescape
func VqsubqU64(r *arm.Uint64X2, v0 *arm.Uint64X2, v1 *arm.Uint64X2)

// Unsigned saturating Subtract. This instruction subtracts the element values of the second source SIMD&FP register from the corresponding element values of the first source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VqsubqU16 VqsubqU16
//go:noescape
func VqsubqU16(r *arm.Uint16X8, v0 *arm.Uint16X8, v1 *arm.Uint16X8)

// Signed saturating Subtract. This instruction subtracts the element values of the second source SIMD&FP register from the corresponding element values of the first source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VqsubqS8 VqsubqS8
//go:noescape
func VqsubqS8(r *arm.Int8X16, v0 *arm.Int8X16, v1 *arm.Int8X16)

// Signed saturating Subtract. This instruction subtracts the element values of the second source SIMD&FP register from the corresponding element values of the first source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VqsubqS32 VqsubqS32
//go:noescape
func VqsubqS32(r *arm.Int32X4, v0 *arm.Int32X4, v1 *arm.Int32X4)

// Signed saturating Subtract. This instruction subtracts the element values of the second source SIMD&FP register from the corresponding element values of the first source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VqsubqS64 VqsubqS64
//go:noescape
func VqsubqS64(r *arm.Int64X2, v0 *arm.Int64X2, v1 *arm.Int64X2)

// Signed saturating Subtract. This instruction subtracts the element values of the second source SIMD&FP register from the corresponding element values of the first source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VqsubqS16 VqsubqS16
//go:noescape
func VqsubqS16(r *arm.Int16X8, v0 *arm.Int16X8, v1 *arm.Int16X8)

// Unsigned saturating Subtract. This instruction subtracts the element values of the second source SIMD&FP register from the corresponding element values of the first source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VqsubU8 VqsubU8
//go:noescape
func VqsubU8(r *arm.Uint8X8, v0 *arm.Uint8X8, v1 *arm.Uint8X8)

// Unsigned saturating Subtract. This instruction subtracts the element values of the second source SIMD&FP register from the corresponding element values of the first source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VqsubU32 VqsubU32
//go:noescape
func VqsubU32(r *arm.Uint32X2, v0 *arm.Uint32X2, v1 *arm.Uint32X2)

// Unsigned saturating Subtract. This instruction subtracts the element values of the second source SIMD&FP register from the corresponding element values of the first source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VqsubU64 VqsubU64
//go:noescape
func VqsubU64(r *arm.Uint64X1, v0 *arm.Uint64X1, v1 *arm.Uint64X1)

// Unsigned saturating Subtract. This instruction subtracts the element values of the second source SIMD&FP register from the corresponding element values of the first source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VqsubU16 VqsubU16
//go:noescape
func VqsubU16(r *arm.Uint16X4, v0 *arm.Uint16X4, v1 *arm.Uint16X4)

// Signed saturating Subtract. This instruction subtracts the element values of the second source SIMD&FP register from the corresponding element values of the first source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VqsubS8 VqsubS8
//go:noescape
func VqsubS8(r *arm.Int8X8, v0 *arm.Int8X8, v1 *arm.Int8X8)

// Signed saturating Subtract. This instruction subtracts the element values of the second source SIMD&FP register from the corresponding element values of the first source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VqsubS32 VqsubS32
//go:noescape
func VqsubS32(r *arm.Int32X2, v0 *arm.Int32X2, v1 *arm.Int32X2)

// Signed saturating Subtract. This instruction subtracts the element values of the second source SIMD&FP register from the corresponding element values of the first source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VqsubS64 VqsubS64
//go:noescape
func VqsubS64(r *arm.Int64X1, v0 *arm.Int64X1, v1 *arm.Int64X1)

// Signed saturating Subtract. This instruction subtracts the element values of the second source SIMD&FP register from the corresponding element values of the first source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VqsubS16 VqsubS16
//go:noescape
func VqsubS16(r *arm.Int16X4, v0 *arm.Int16X4, v1 *arm.Int16X4)

// Rounding Add returning High Narrow. This instruction adds each vector element in the first source SIMD&FP register to the corresponding vector element in the second source SIMD&FP register, places the most significant half of the result into a vector, and writes the vector to the lower or upper half of the destination SIMD&FP register.
//
//go:linkname VraddhnU32 VraddhnU32
//go:noescape
func VraddhnU32(r *arm.Uint16X4, v0 *arm.Uint32X4, v1 *arm.Uint32X4)

// Rounding Add returning High Narrow. This instruction adds each vector element in the first source SIMD&FP register to the corresponding vector element in the second source SIMD&FP register, places the most significant half of the result into a vector, and writes the vector to the lower or upper half of the destination SIMD&FP register.
//
//go:linkname VraddhnU64 VraddhnU64
//go:noescape
func VraddhnU64(r *arm.Uint32X2, v0 *arm.Uint64X2, v1 *arm.Uint64X2)

// Rounding Add returning High Narrow. This instruction adds each vector element in the first source SIMD&FP register to the corresponding vector element in the second source SIMD&FP register, places the most significant half of the result into a vector, and writes the vector to the lower or upper half of the destination SIMD&FP register.
//
//go:linkname VraddhnU16 VraddhnU16
//go:noescape
func VraddhnU16(r *arm.Uint8X8, v0 *arm.Uint16X8, v1 *arm.Uint16X8)

// Rounding Add returning High Narrow. This instruction adds each vector element in the first source SIMD&FP register to the corresponding vector element in the second source SIMD&FP register, places the most significant half of the result into a vector, and writes the vector to the lower or upper half of the destination SIMD&FP register.
//
//go:linkname VraddhnS32 VraddhnS32
//go:noescape
func VraddhnS32(r *arm.Int16X4, v0 *arm.Int32X4, v1 *arm.Int32X4)

// Rounding Add returning High Narrow. This instruction adds each vector element in the first source SIMD&FP register to the corresponding vector element in the second source SIMD&FP register, places the most significant half of the result into a vector, and writes the vector to the lower or upper half of the destination SIMD&FP register.
//
//go:linkname VraddhnS64 VraddhnS64
//go:noescape
func VraddhnS64(r *arm.Int32X2, v0 *arm.Int64X2, v1 *arm.Int64X2)

// Rounding Add returning High Narrow. This instruction adds each vector element in the first source SIMD&FP register to the corresponding vector element in the second source SIMD&FP register, places the most significant half of the result into a vector, and writes the vector to the lower or upper half of the destination SIMD&FP register.
//
//go:linkname VraddhnS16 VraddhnS16
//go:noescape
func VraddhnS16(r *arm.Int8X8, v0 *arm.Int16X8, v1 *arm.Int16X8)

// Unsigned Reciprocal Estimate. This instruction reads each vector element from the source SIMD&FP register, calculates an approximate inverse for the unsigned integer value, places the result into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VrecpeqU32 VrecpeqU32
//go:noescape
func VrecpeqU32(r *arm.Uint32X4, v0 *arm.Uint32X4)

// Floating-point Reciprocal Estimate. This instruction finds an approximate reciprocal estimate for each vector element in the source SIMD&FP register, places the result in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VrecpeqF32 VrecpeqF32
//go:noescape
func VrecpeqF32(r *arm.Float32X4, v0 *arm.Float32X4)

// Unsigned Reciprocal Estimate. This instruction reads each vector element from the source SIMD&FP register, calculates an approximate inverse for the unsigned integer value, places the result into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VrecpeU32 VrecpeU32
//go:noescape
func VrecpeU32(r *arm.Uint32X2, v0 *arm.Uint32X2)

// Floating-point Reciprocal Estimate. This instruction finds an approximate reciprocal estimate for each vector element in the source SIMD&FP register, places the result in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VrecpeF32 VrecpeF32
//go:noescape
func VrecpeF32(r *arm.Float32X2, v0 *arm.Float32X2)

// Floating-point Reciprocal Step. This instruction multiplies the corresponding floating-point values in the vectors of the two source SIMD&FP registers, subtracts each of the products from 2.0, places the resulting floating-point values in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VrecpsqF32 VrecpsqF32
//go:noescape
func VrecpsqF32(r *arm.Float32X4, v0 *arm.Float32X4, v1 *arm.Float32X4)

// Floating-point Reciprocal Step. This instruction multiplies the corresponding floating-point values in the vectors of the two source SIMD&FP registers, subtracts each of the products from 2.0, places the resulting floating-point values in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VrecpsF32 VrecpsF32
//go:noescape
func VrecpsF32(r *arm.Float32X2, v0 *arm.Float32X2, v1 *arm.Float32X2)

// Reverse elements in 16-bit halfwords (vector). This instruction reverses the order of 8-bit elements in each halfword of the vector in the source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname Vrev16P8 Vrev16P8
//go:noescape
func Vrev16P8(r *arm.Poly8X8, v0 *arm.Poly8X8)

// Reverse elements in 16-bit halfwords (vector). This instruction reverses the order of 8-bit elements in each halfword of the vector in the source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname Vrev16QP8 Vrev16QP8
//go:noescape
func Vrev16QP8(r *arm.Poly8X16, v0 *arm.Poly8X16)

// Reverse elements in 16-bit halfwords (vector). This instruction reverses the order of 8-bit elements in each halfword of the vector in the source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname Vrev16QU8 Vrev16QU8
//go:noescape
func Vrev16QU8(r *arm.Uint8X16, v0 *arm.Uint8X16)

// Reverse elements in 16-bit halfwords (vector). This instruction reverses the order of 8-bit elements in each halfword of the vector in the source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname Vrev16QS8 Vrev16QS8
//go:noescape
func Vrev16QS8(r *arm.Int8X16, v0 *arm.Int8X16)

// Reverse elements in 16-bit halfwords (vector). This instruction reverses the order of 8-bit elements in each halfword of the vector in the source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname Vrev16U8 Vrev16U8
//go:noescape
func Vrev16U8(r *arm.Uint8X8, v0 *arm.Uint8X8)

// Reverse elements in 16-bit halfwords (vector). This instruction reverses the order of 8-bit elements in each halfword of the vector in the source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname Vrev16S8 Vrev16S8
//go:noescape
func Vrev16S8(r *arm.Int8X8, v0 *arm.Int8X8)

// Reverse elements in 32-bit words (vector). This instruction reverses the order of 8-bit or 16-bit elements in each word of the vector in the source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname Vrev32P8 Vrev32P8
//go:noescape
func Vrev32P8(r *arm.Poly8X8, v0 *arm.Poly8X8)

// Reverse elements in 32-bit words (vector). This instruction reverses the order of 8-bit or 16-bit elements in each word of the vector in the source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname Vrev32P16 Vrev32P16
//go:noescape
func Vrev32P16(r *arm.Poly16X4, v0 *arm.Poly16X4)

// Reverse elements in 32-bit words (vector). This instruction reverses the order of 8-bit or 16-bit elements in each word of the vector in the source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname Vrev32QP8 Vrev32QP8
//go:noescape
func Vrev32QP8(r *arm.Poly8X16, v0 *arm.Poly8X16)

// Reverse elements in 32-bit words (vector). This instruction reverses the order of 8-bit or 16-bit elements in each word of the vector in the source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname Vrev32QP16 Vrev32QP16
//go:noescape
func Vrev32QP16(r *arm.Poly16X8, v0 *arm.Poly16X8)

// Reverse elements in 32-bit words (vector). This instruction reverses the order of 8-bit or 16-bit elements in each word of the vector in the source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname Vrev32QU8 Vrev32QU8
//go:noescape
func Vrev32QU8(r *arm.Uint8X16, v0 *arm.Uint8X16)

// Reverse elements in 32-bit words (vector). This instruction reverses the order of 8-bit or 16-bit elements in each word of the vector in the source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname Vrev32QU16 Vrev32QU16
//go:noescape
func Vrev32QU16(r *arm.Uint16X8, v0 *arm.Uint16X8)

// Reverse elements in 32-bit words (vector). This instruction reverses the order of 8-bit or 16-bit elements in each word of the vector in the source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname Vrev32QS8 Vrev32QS8
//go:noescape
func Vrev32QS8(r *arm.Int8X16, v0 *arm.Int8X16)

// Reverse elements in 32-bit words (vector). This instruction reverses the order of 8-bit or 16-bit elements in each word of the vector in the source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname Vrev32QS16 Vrev32QS16
//go:noescape
func Vrev32QS16(r *arm.Int16X8, v0 *arm.Int16X8)

// Reverse elements in 32-bit words (vector). This instruction reverses the order of 8-bit or 16-bit elements in each word of the vector in the source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname Vrev32U8 Vrev32U8
//go:noescape
func Vrev32U8(r *arm.Uint8X8, v0 *arm.Uint8X8)

// Reverse elements in 32-bit words (vector). This instruction reverses the order of 8-bit or 16-bit elements in each word of the vector in the source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname Vrev32U16 Vrev32U16
//go:noescape
func Vrev32U16(r *arm.Uint16X4, v0 *arm.Uint16X4)

// Reverse elements in 32-bit words (vector). This instruction reverses the order of 8-bit or 16-bit elements in each word of the vector in the source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname Vrev32S8 Vrev32S8
//go:noescape
func Vrev32S8(r *arm.Int8X8, v0 *arm.Int8X8)

// Reverse elements in 32-bit words (vector). This instruction reverses the order of 8-bit or 16-bit elements in each word of the vector in the source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname Vrev32S16 Vrev32S16
//go:noescape
func Vrev32S16(r *arm.Int16X4, v0 *arm.Int16X4)

// Reverse elements in 64-bit doublewords (vector). This instruction reverses the order of 8-bit, 16-bit, or 32-bit elements in each doubleword of the vector in the source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname Vrev64P8 Vrev64P8
//go:noescape
func Vrev64P8(r *arm.Poly8X8, v0 *arm.Poly8X8)

// Reverse elements in 64-bit doublewords (vector). This instruction reverses the order of 8-bit, 16-bit, or 32-bit elements in each doubleword of the vector in the source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname Vrev64P16 Vrev64P16
//go:noescape
func Vrev64P16(r *arm.Poly16X4, v0 *arm.Poly16X4)

// Reverse elements in 64-bit doublewords (vector). This instruction reverses the order of 8-bit, 16-bit, or 32-bit elements in each doubleword of the vector in the source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname Vrev64QP8 Vrev64QP8
//go:noescape
func Vrev64QP8(r *arm.Poly8X16, v0 *arm.Poly8X16)

// Reverse elements in 64-bit doublewords (vector). This instruction reverses the order of 8-bit, 16-bit, or 32-bit elements in each doubleword of the vector in the source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname Vrev64QP16 Vrev64QP16
//go:noescape
func Vrev64QP16(r *arm.Poly16X8, v0 *arm.Poly16X8)

// Reverse elements in 64-bit doublewords (vector). This instruction reverses the order of 8-bit, 16-bit, or 32-bit elements in each doubleword of the vector in the source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname Vrev64QU8 Vrev64QU8
//go:noescape
func Vrev64QU8(r *arm.Uint8X16, v0 *arm.Uint8X16)

// Reverse elements in 64-bit doublewords (vector). This instruction reverses the order of 8-bit, 16-bit, or 32-bit elements in each doubleword of the vector in the source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname Vrev64QU32 Vrev64QU32
//go:noescape
func Vrev64QU32(r *arm.Uint32X4, v0 *arm.Uint32X4)

// Reverse elements in 64-bit doublewords (vector). This instruction reverses the order of 8-bit, 16-bit, or 32-bit elements in each doubleword of the vector in the source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname Vrev64QU16 Vrev64QU16
//go:noescape
func Vrev64QU16(r *arm.Uint16X8, v0 *arm.Uint16X8)

// Reverse elements in 64-bit doublewords (vector). This instruction reverses the order of 8-bit, 16-bit, or 32-bit elements in each doubleword of the vector in the source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname Vrev64QS8 Vrev64QS8
//go:noescape
func Vrev64QS8(r *arm.Int8X16, v0 *arm.Int8X16)

// Reverse elements in 64-bit doublewords (vector). This instruction reverses the order of 8-bit, 16-bit, or 32-bit elements in each doubleword of the vector in the source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname Vrev64QF32 Vrev64QF32
//go:noescape
func Vrev64QF32(r *arm.Float32X4, v0 *arm.Float32X4)

// Reverse elements in 64-bit doublewords (vector). This instruction reverses the order of 8-bit, 16-bit, or 32-bit elements in each doubleword of the vector in the source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname Vrev64QS32 Vrev64QS32
//go:noescape
func Vrev64QS32(r *arm.Int32X4, v0 *arm.Int32X4)

// Reverse elements in 64-bit doublewords (vector). This instruction reverses the order of 8-bit, 16-bit, or 32-bit elements in each doubleword of the vector in the source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname Vrev64QS16 Vrev64QS16
//go:noescape
func Vrev64QS16(r *arm.Int16X8, v0 *arm.Int16X8)

// Reverse elements in 64-bit doublewords (vector). This instruction reverses the order of 8-bit, 16-bit, or 32-bit elements in each doubleword of the vector in the source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname Vrev64U8 Vrev64U8
//go:noescape
func Vrev64U8(r *arm.Uint8X8, v0 *arm.Uint8X8)

// Reverse elements in 64-bit doublewords (vector). This instruction reverses the order of 8-bit, 16-bit, or 32-bit elements in each doubleword of the vector in the source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname Vrev64U32 Vrev64U32
//go:noescape
func Vrev64U32(r *arm.Uint32X2, v0 *arm.Uint32X2)

// Reverse elements in 64-bit doublewords (vector). This instruction reverses the order of 8-bit, 16-bit, or 32-bit elements in each doubleword of the vector in the source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname Vrev64U16 Vrev64U16
//go:noescape
func Vrev64U16(r *arm.Uint16X4, v0 *arm.Uint16X4)

// Reverse elements in 64-bit doublewords (vector). This instruction reverses the order of 8-bit, 16-bit, or 32-bit elements in each doubleword of the vector in the source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname Vrev64S8 Vrev64S8
//go:noescape
func Vrev64S8(r *arm.Int8X8, v0 *arm.Int8X8)

// Reverse elements in 64-bit doublewords (vector). This instruction reverses the order of 8-bit, 16-bit, or 32-bit elements in each doubleword of the vector in the source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname Vrev64F32 Vrev64F32
//go:noescape
func Vrev64F32(r *arm.Float32X2, v0 *arm.Float32X2)

// Reverse elements in 64-bit doublewords (vector). This instruction reverses the order of 8-bit, 16-bit, or 32-bit elements in each doubleword of the vector in the source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname Vrev64S32 Vrev64S32
//go:noescape
func Vrev64S32(r *arm.Int32X2, v0 *arm.Int32X2)

// Reverse elements in 64-bit doublewords (vector). This instruction reverses the order of 8-bit, 16-bit, or 32-bit elements in each doubleword of the vector in the source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname Vrev64S16 Vrev64S16
//go:noescape
func Vrev64S16(r *arm.Int16X4, v0 *arm.Int16X4)

// Unsigned Rounding Halving Add. This instruction adds corresponding unsigned integer values from the two source SIMD&FP registers, shifts each result right one bit, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VrhaddqU8 VrhaddqU8
//go:noescape
func VrhaddqU8(r *arm.Uint8X16, v0 *arm.Uint8X16, v1 *arm.Uint8X16)

// Unsigned Rounding Halving Add. This instruction adds corresponding unsigned integer values from the two source SIMD&FP registers, shifts each result right one bit, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VrhaddqU32 VrhaddqU32
//go:noescape
func VrhaddqU32(r *arm.Uint32X4, v0 *arm.Uint32X4, v1 *arm.Uint32X4)

// Unsigned Rounding Halving Add. This instruction adds corresponding unsigned integer values from the two source SIMD&FP registers, shifts each result right one bit, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VrhaddqU16 VrhaddqU16
//go:noescape
func VrhaddqU16(r *arm.Uint16X8, v0 *arm.Uint16X8, v1 *arm.Uint16X8)

// Signed Rounding Halving Add. This instruction adds corresponding signed integer values from the two source SIMD&FP registers, shifts each result right one bit, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VrhaddqS8 VrhaddqS8
//go:noescape
func VrhaddqS8(r *arm.Int8X16, v0 *arm.Int8X16, v1 *arm.Int8X16)

// Signed Rounding Halving Add. This instruction adds corresponding signed integer values from the two source SIMD&FP registers, shifts each result right one bit, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VrhaddqS32 VrhaddqS32
//go:noescape
func VrhaddqS32(r *arm.Int32X4, v0 *arm.Int32X4, v1 *arm.Int32X4)

// Signed Rounding Halving Add. This instruction adds corresponding signed integer values from the two source SIMD&FP registers, shifts each result right one bit, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VrhaddqS16 VrhaddqS16
//go:noescape
func VrhaddqS16(r *arm.Int16X8, v0 *arm.Int16X8, v1 *arm.Int16X8)

// Unsigned Rounding Halving Add. This instruction adds corresponding unsigned integer values from the two source SIMD&FP registers, shifts each result right one bit, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VrhaddU8 VrhaddU8
//go:noescape
func VrhaddU8(r *arm.Uint8X8, v0 *arm.Uint8X8, v1 *arm.Uint8X8)

// Unsigned Rounding Halving Add. This instruction adds corresponding unsigned integer values from the two source SIMD&FP registers, shifts each result right one bit, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VrhaddU32 VrhaddU32
//go:noescape
func VrhaddU32(r *arm.Uint32X2, v0 *arm.Uint32X2, v1 *arm.Uint32X2)

// Unsigned Rounding Halving Add. This instruction adds corresponding unsigned integer values from the two source SIMD&FP registers, shifts each result right one bit, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VrhaddU16 VrhaddU16
//go:noescape
func VrhaddU16(r *arm.Uint16X4, v0 *arm.Uint16X4, v1 *arm.Uint16X4)

// Signed Rounding Halving Add. This instruction adds corresponding signed integer values from the two source SIMD&FP registers, shifts each result right one bit, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VrhaddS8 VrhaddS8
//go:noescape
func VrhaddS8(r *arm.Int8X8, v0 *arm.Int8X8, v1 *arm.Int8X8)

// Signed Rounding Halving Add. This instruction adds corresponding signed integer values from the two source SIMD&FP registers, shifts each result right one bit, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VrhaddS32 VrhaddS32
//go:noescape
func VrhaddS32(r *arm.Int32X2, v0 *arm.Int32X2, v1 *arm.Int32X2)

// Signed Rounding Halving Add. This instruction adds corresponding signed integer values from the two source SIMD&FP registers, shifts each result right one bit, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VrhaddS16 VrhaddS16
//go:noescape
func VrhaddS16(r *arm.Int16X4, v0 *arm.Int16X4, v1 *arm.Int16X4)

// Unsigned Rounding Shift Left (register). This instruction takes each element in the vector of the first source SIMD&FP register, shifts the vector element by a value from the least significant byte of the corresponding element of the second source SIMD&FP register, places the results in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VrshlqU8 VrshlqU8
//go:noescape
func VrshlqU8(r *arm.Uint8X16, v0 *arm.Uint8X16, v1 *arm.Int8X16)

// Unsigned Rounding Shift Left (register). This instruction takes each element in the vector of the first source SIMD&FP register, shifts the vector element by a value from the least significant byte of the corresponding element of the second source SIMD&FP register, places the results in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VrshlqU32 VrshlqU32
//go:noescape
func VrshlqU32(r *arm.Uint32X4, v0 *arm.Uint32X4, v1 *arm.Int32X4)

// Unsigned Rounding Shift Left (register). This instruction takes each element in the vector of the first source SIMD&FP register, shifts the vector element by a value from the least significant byte of the corresponding element of the second source SIMD&FP register, places the results in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VrshlqU64 VrshlqU64
//go:noescape
func VrshlqU64(r *arm.Uint64X2, v0 *arm.Uint64X2, v1 *arm.Int64X2)

// Unsigned Rounding Shift Left (register). This instruction takes each element in the vector of the first source SIMD&FP register, shifts the vector element by a value from the least significant byte of the corresponding element of the second source SIMD&FP register, places the results in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VrshlqU16 VrshlqU16
//go:noescape
func VrshlqU16(r *arm.Uint16X8, v0 *arm.Uint16X8, v1 *arm.Int16X8)

// Signed Rounding Shift Left (register). This instruction takes each signed integer value in the vector of the first source SIMD&FP register, shifts it by a value from the least significant byte of the corresponding element of the second source SIMD&FP register, places the results in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VrshlqS8 VrshlqS8
//go:noescape
func VrshlqS8(r *arm.Int8X16, v0 *arm.Int8X16, v1 *arm.Int8X16)

// Signed Rounding Shift Left (register). This instruction takes each signed integer value in the vector of the first source SIMD&FP register, shifts it by a value from the least significant byte of the corresponding element of the second source SIMD&FP register, places the results in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VrshlqS32 VrshlqS32
//go:noescape
func VrshlqS32(r *arm.Int32X4, v0 *arm.Int32X4, v1 *arm.Int32X4)

// Signed Rounding Shift Left (register). This instruction takes each signed integer value in the vector of the first source SIMD&FP register, shifts it by a value from the least significant byte of the corresponding element of the second source SIMD&FP register, places the results in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VrshlqS64 VrshlqS64
//go:noescape
func VrshlqS64(r *arm.Int64X2, v0 *arm.Int64X2, v1 *arm.Int64X2)

// Signed Rounding Shift Left (register). This instruction takes each signed integer value in the vector of the first source SIMD&FP register, shifts it by a value from the least significant byte of the corresponding element of the second source SIMD&FP register, places the results in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VrshlqS16 VrshlqS16
//go:noescape
func VrshlqS16(r *arm.Int16X8, v0 *arm.Int16X8, v1 *arm.Int16X8)

// Unsigned Rounding Shift Left (register). This instruction takes each element in the vector of the first source SIMD&FP register, shifts the vector element by a value from the least significant byte of the corresponding element of the second source SIMD&FP register, places the results in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VrshlU8 VrshlU8
//go:noescape
func VrshlU8(r *arm.Uint8X8, v0 *arm.Uint8X8, v1 *arm.Int8X8)

// Unsigned Rounding Shift Left (register). This instruction takes each element in the vector of the first source SIMD&FP register, shifts the vector element by a value from the least significant byte of the corresponding element of the second source SIMD&FP register, places the results in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VrshlU32 VrshlU32
//go:noescape
func VrshlU32(r *arm.Uint32X2, v0 *arm.Uint32X2, v1 *arm.Int32X2)

// Unsigned Rounding Shift Left (register). This instruction takes each element in the vector of the first source SIMD&FP register, shifts the vector element by a value from the least significant byte of the corresponding element of the second source SIMD&FP register, places the results in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VrshlU64 VrshlU64
//go:noescape
func VrshlU64(r *arm.Uint64X1, v0 *arm.Uint64X1, v1 *arm.Int64X1)

// Unsigned Rounding Shift Left (register). This instruction takes each element in the vector of the first source SIMD&FP register, shifts the vector element by a value from the least significant byte of the corresponding element of the second source SIMD&FP register, places the results in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VrshlU16 VrshlU16
//go:noescape
func VrshlU16(r *arm.Uint16X4, v0 *arm.Uint16X4, v1 *arm.Int16X4)

// Signed Rounding Shift Left (register). This instruction takes each signed integer value in the vector of the first source SIMD&FP register, shifts it by a value from the least significant byte of the corresponding element of the second source SIMD&FP register, places the results in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VrshlS8 VrshlS8
//go:noescape
func VrshlS8(r *arm.Int8X8, v0 *arm.Int8X8, v1 *arm.Int8X8)

// Signed Rounding Shift Left (register). This instruction takes each signed integer value in the vector of the first source SIMD&FP register, shifts it by a value from the least significant byte of the corresponding element of the second source SIMD&FP register, places the results in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VrshlS32 VrshlS32
//go:noescape
func VrshlS32(r *arm.Int32X2, v0 *arm.Int32X2, v1 *arm.Int32X2)

// Signed Rounding Shift Left (register). This instruction takes each signed integer value in the vector of the first source SIMD&FP register, shifts it by a value from the least significant byte of the corresponding element of the second source SIMD&FP register, places the results in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VrshlS64 VrshlS64
//go:noescape
func VrshlS64(r *arm.Int64X1, v0 *arm.Int64X1, v1 *arm.Int64X1)

// Signed Rounding Shift Left (register). This instruction takes each signed integer value in the vector of the first source SIMD&FP register, shifts it by a value from the least significant byte of the corresponding element of the second source SIMD&FP register, places the results in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VrshlS16 VrshlS16
//go:noescape
func VrshlS16(r *arm.Int16X4, v0 *arm.Int16X4, v1 *arm.Int16X4)

// Unsigned Reciprocal Square Root Estimate. This instruction reads each vector element from the source SIMD&FP register, calculates an approximate inverse square root for each value, places the result into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are unsigned integer values.
//
//go:linkname VrsqrteqU32 VrsqrteqU32
//go:noescape
func VrsqrteqU32(r *arm.Uint32X4, v0 *arm.Uint32X4)

// Floating-point Reciprocal Square Root Estimate. This instruction calculates an approximate square root for each vector element in the source SIMD&FP register, places the result in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VrsqrteqF32 VrsqrteqF32
//go:noescape
func VrsqrteqF32(r *arm.Float32X4, v0 *arm.Float32X4)

// Unsigned Reciprocal Square Root Estimate. This instruction reads each vector element from the source SIMD&FP register, calculates an approximate inverse square root for each value, places the result into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are unsigned integer values.
//
//go:linkname VrsqrteU32 VrsqrteU32
//go:noescape
func VrsqrteU32(r *arm.Uint32X2, v0 *arm.Uint32X2)

// Floating-point Reciprocal Square Root Estimate. This instruction calculates an approximate square root for each vector element in the source SIMD&FP register, places the result in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VrsqrteF32 VrsqrteF32
//go:noescape
func VrsqrteF32(r *arm.Float32X2, v0 *arm.Float32X2)

// Floating-point Reciprocal Square Root Step. This instruction multiplies corresponding floating-point values in the vectors of the two source SIMD&FP registers, subtracts each of the products from 3.0, divides these results by 2.0, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VrsqrtsqF32 VrsqrtsqF32
//go:noescape
func VrsqrtsqF32(r *arm.Float32X4, v0 *arm.Float32X4, v1 *arm.Float32X4)

// Floating-point Reciprocal Square Root Step. This instruction multiplies corresponding floating-point values in the vectors of the two source SIMD&FP registers, subtracts each of the products from 3.0, divides these results by 2.0, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VrsqrtsF32 VrsqrtsF32
//go:noescape
func VrsqrtsF32(r *arm.Float32X2, v0 *arm.Float32X2, v1 *arm.Float32X2)

// Rounding Subtract returning High Narrow. This instruction subtracts each vector element of the second source SIMD&FP register from the corresponding vector element of the first source SIMD&FP register, places the most significant half of the result into a vector, and writes the vector to the lower or upper half of the destination SIMD&FP register.
//
//go:linkname VrsubhnU32 VrsubhnU32
//go:noescape
func VrsubhnU32(r *arm.Uint16X4, v0 *arm.Uint32X4, v1 *arm.Uint32X4)

// Rounding Subtract returning High Narrow. This instruction subtracts each vector element of the second source SIMD&FP register from the corresponding vector element of the first source SIMD&FP register, places the most significant half of the result into a vector, and writes the vector to the lower or upper half of the destination SIMD&FP register.
//
//go:linkname VrsubhnU64 VrsubhnU64
//go:noescape
func VrsubhnU64(r *arm.Uint32X2, v0 *arm.Uint64X2, v1 *arm.Uint64X2)

// Rounding Subtract returning High Narrow. This instruction subtracts each vector element of the second source SIMD&FP register from the corresponding vector element of the first source SIMD&FP register, places the most significant half of the result into a vector, and writes the vector to the lower or upper half of the destination SIMD&FP register.
//
//go:linkname VrsubhnU16 VrsubhnU16
//go:noescape
func VrsubhnU16(r *arm.Uint8X8, v0 *arm.Uint16X8, v1 *arm.Uint16X8)

// Rounding Subtract returning High Narrow. This instruction subtracts each vector element of the second source SIMD&FP register from the corresponding vector element of the first source SIMD&FP register, places the most significant half of the result into a vector, and writes the vector to the lower or upper half of the destination SIMD&FP register.
//
//go:linkname VrsubhnS32 VrsubhnS32
//go:noescape
func VrsubhnS32(r *arm.Int16X4, v0 *arm.Int32X4, v1 *arm.Int32X4)

// Rounding Subtract returning High Narrow. This instruction subtracts each vector element of the second source SIMD&FP register from the corresponding vector element of the first source SIMD&FP register, places the most significant half of the result into a vector, and writes the vector to the lower or upper half of the destination SIMD&FP register.
//
//go:linkname VrsubhnS64 VrsubhnS64
//go:noescape
func VrsubhnS64(r *arm.Int32X2, v0 *arm.Int64X2, v1 *arm.Int64X2)

// Rounding Subtract returning High Narrow. This instruction subtracts each vector element of the second source SIMD&FP register from the corresponding vector element of the first source SIMD&FP register, places the most significant half of the result into a vector, and writes the vector to the lower or upper half of the destination SIMD&FP register.
//
//go:linkname VrsubhnS16 VrsubhnS16
//go:noescape
func VrsubhnS16(r *arm.Int8X8, v0 *arm.Int16X8, v1 *arm.Int16X8)

// Unsigned Shift Left (register). This instruction takes each element in the vector of the first source SIMD&FP register, shifts each element by a value from the least significant byte of the corresponding element of the second source SIMD&FP register, places the results in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VshlqU8 VshlqU8
//go:noescape
func VshlqU8(r *arm.Uint8X16, v0 *arm.Uint8X16, v1 *arm.Int8X16)

// Unsigned Shift Left (register). This instruction takes each element in the vector of the first source SIMD&FP register, shifts each element by a value from the least significant byte of the corresponding element of the second source SIMD&FP register, places the results in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VshlqU32 VshlqU32
//go:noescape
func VshlqU32(r *arm.Uint32X4, v0 *arm.Uint32X4, v1 *arm.Int32X4)

// Unsigned Shift Left (register). This instruction takes each element in the vector of the first source SIMD&FP register, shifts each element by a value from the least significant byte of the corresponding element of the second source SIMD&FP register, places the results in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VshlqU64 VshlqU64
//go:noescape
func VshlqU64(r *arm.Uint64X2, v0 *arm.Uint64X2, v1 *arm.Int64X2)

// Unsigned Shift Left (register). This instruction takes each element in the vector of the first source SIMD&FP register, shifts each element by a value from the least significant byte of the corresponding element of the second source SIMD&FP register, places the results in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VshlqU16 VshlqU16
//go:noescape
func VshlqU16(r *arm.Uint16X8, v0 *arm.Uint16X8, v1 *arm.Int16X8)

// Signed Shift Left (register). This instruction takes each signed integer value in the vector of the first source SIMD&FP register, shifts each value by a value from the least significant byte of the corresponding element of the second source SIMD&FP register, places the results in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VshlqS8 VshlqS8
//go:noescape
func VshlqS8(r *arm.Int8X16, v0 *arm.Int8X16, v1 *arm.Int8X16)

// Signed Shift Left (register). This instruction takes each signed integer value in the vector of the first source SIMD&FP register, shifts each value by a value from the least significant byte of the corresponding element of the second source SIMD&FP register, places the results in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VshlqS32 VshlqS32
//go:noescape
func VshlqS32(r *arm.Int32X4, v0 *arm.Int32X4, v1 *arm.Int32X4)

// Signed Shift Left (register). This instruction takes each signed integer value in the vector of the first source SIMD&FP register, shifts each value by a value from the least significant byte of the corresponding element of the second source SIMD&FP register, places the results in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VshlqS64 VshlqS64
//go:noescape
func VshlqS64(r *arm.Int64X2, v0 *arm.Int64X2, v1 *arm.Int64X2)

// Signed Shift Left (register). This instruction takes each signed integer value in the vector of the first source SIMD&FP register, shifts each value by a value from the least significant byte of the corresponding element of the second source SIMD&FP register, places the results in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VshlqS16 VshlqS16
//go:noescape
func VshlqS16(r *arm.Int16X8, v0 *arm.Int16X8, v1 *arm.Int16X8)

// Unsigned Shift Left (register). This instruction takes each element in the vector of the first source SIMD&FP register, shifts each element by a value from the least significant byte of the corresponding element of the second source SIMD&FP register, places the results in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VshlU8 VshlU8
//go:noescape
func VshlU8(r *arm.Uint8X8, v0 *arm.Uint8X8, v1 *arm.Int8X8)

// Unsigned Shift Left (register). This instruction takes each element in the vector of the first source SIMD&FP register, shifts each element by a value from the least significant byte of the corresponding element of the second source SIMD&FP register, places the results in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VshlU32 VshlU32
//go:noescape
func VshlU32(r *arm.Uint32X2, v0 *arm.Uint32X2, v1 *arm.Int32X2)

// Unsigned Shift Left (register). This instruction takes each element in the vector of the first source SIMD&FP register, shifts each element by a value from the least significant byte of the corresponding element of the second source SIMD&FP register, places the results in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VshlU64 VshlU64
//go:noescape
func VshlU64(r *arm.Uint64X1, v0 *arm.Uint64X1, v1 *arm.Int64X1)

// Unsigned Shift Left (register). This instruction takes each element in the vector of the first source SIMD&FP register, shifts each element by a value from the least significant byte of the corresponding element of the second source SIMD&FP register, places the results in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VshlU16 VshlU16
//go:noescape
func VshlU16(r *arm.Uint16X4, v0 *arm.Uint16X4, v1 *arm.Int16X4)

// Signed Shift Left (register). This instruction takes each signed integer value in the vector of the first source SIMD&FP register, shifts each value by a value from the least significant byte of the corresponding element of the second source SIMD&FP register, places the results in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VshlS8 VshlS8
//go:noescape
func VshlS8(r *arm.Int8X8, v0 *arm.Int8X8, v1 *arm.Int8X8)

// Signed Shift Left (register). This instruction takes each signed integer value in the vector of the first source SIMD&FP register, shifts each value by a value from the least significant byte of the corresponding element of the second source SIMD&FP register, places the results in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VshlS32 VshlS32
//go:noescape
func VshlS32(r *arm.Int32X2, v0 *arm.Int32X2, v1 *arm.Int32X2)

// Signed Shift Left (register). This instruction takes each signed integer value in the vector of the first source SIMD&FP register, shifts each value by a value from the least significant byte of the corresponding element of the second source SIMD&FP register, places the results in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VshlS64 VshlS64
//go:noescape
func VshlS64(r *arm.Int64X1, v0 *arm.Int64X1, v1 *arm.Int64X1)

// Signed Shift Left (register). This instruction takes each signed integer value in the vector of the first source SIMD&FP register, shifts each value by a value from the least significant byte of the corresponding element of the second source SIMD&FP register, places the results in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VshlS16 VshlS16
//go:noescape
func VshlS16(r *arm.Int16X4, v0 *arm.Int16X4, v1 *arm.Int16X4)

// Subtract (vector). This instruction subtracts each vector element in the second source SIMD&FP register from the corresponding vector element in the first source SIMD&FP register, places the result into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VsubqU8 VsubqU8
//go:noescape
func VsubqU8(r *arm.Uint8X16, v0 *arm.Uint8X16, v1 *arm.Uint8X16)

// Subtract (vector). This instruction subtracts each vector element in the second source SIMD&FP register from the corresponding vector element in the first source SIMD&FP register, places the result into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VsubqU32 VsubqU32
//go:noescape
func VsubqU32(r *arm.Uint32X4, v0 *arm.Uint32X4, v1 *arm.Uint32X4)

// Subtract (vector). This instruction subtracts each vector element in the second source SIMD&FP register from the corresponding vector element in the first source SIMD&FP register, places the result into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VsubqU64 VsubqU64
//go:noescape
func VsubqU64(r *arm.Uint64X2, v0 *arm.Uint64X2, v1 *arm.Uint64X2)

// Subtract (vector). This instruction subtracts each vector element in the second source SIMD&FP register from the corresponding vector element in the first source SIMD&FP register, places the result into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VsubqU16 VsubqU16
//go:noescape
func VsubqU16(r *arm.Uint16X8, v0 *arm.Uint16X8, v1 *arm.Uint16X8)

// Subtract (vector). This instruction subtracts each vector element in the second source SIMD&FP register from the corresponding vector element in the first source SIMD&FP register, places the result into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VsubqS8 VsubqS8
//go:noescape
func VsubqS8(r *arm.Int8X16, v0 *arm.Int8X16, v1 *arm.Int8X16)

// Floating-point Subtract (vector). This instruction subtracts the elements in the vector in the second source SIMD&FP register, from the corresponding elements in the vector in the first source SIMD&FP register, places each result into elements of a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VsubqF32 VsubqF32
//go:noescape
func VsubqF32(r *arm.Float32X4, v0 *arm.Float32X4, v1 *arm.Float32X4)

// Subtract (vector). This instruction subtracts each vector element in the second source SIMD&FP register from the corresponding vector element in the first source SIMD&FP register, places the result into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VsubqS32 VsubqS32
//go:noescape
func VsubqS32(r *arm.Int32X4, v0 *arm.Int32X4, v1 *arm.Int32X4)

// Subtract (vector). This instruction subtracts each vector element in the second source SIMD&FP register from the corresponding vector element in the first source SIMD&FP register, places the result into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VsubqS64 VsubqS64
//go:noescape
func VsubqS64(r *arm.Int64X2, v0 *arm.Int64X2, v1 *arm.Int64X2)

// Subtract (vector). This instruction subtracts each vector element in the second source SIMD&FP register from the corresponding vector element in the first source SIMD&FP register, places the result into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VsubqS16 VsubqS16
//go:noescape
func VsubqS16(r *arm.Int16X8, v0 *arm.Int16X8, v1 *arm.Int16X8)

// Subtract (vector). This instruction subtracts each vector element in the second source SIMD&FP register from the corresponding vector element in the first source SIMD&FP register, places the result into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VsubU8 VsubU8
//go:noescape
func VsubU8(r *arm.Uint8X8, v0 *arm.Uint8X8, v1 *arm.Uint8X8)

// Subtract (vector). This instruction subtracts each vector element in the second source SIMD&FP register from the corresponding vector element in the first source SIMD&FP register, places the result into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VsubU32 VsubU32
//go:noescape
func VsubU32(r *arm.Uint32X2, v0 *arm.Uint32X2, v1 *arm.Uint32X2)

// Subtract (vector). This instruction subtracts each vector element in the second source SIMD&FP register from the corresponding vector element in the first source SIMD&FP register, places the result into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VsubU64 VsubU64
//go:noescape
func VsubU64(r *arm.Uint64X1, v0 *arm.Uint64X1, v1 *arm.Uint64X1)

// Subtract (vector). This instruction subtracts each vector element in the second source SIMD&FP register from the corresponding vector element in the first source SIMD&FP register, places the result into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VsubU16 VsubU16
//go:noescape
func VsubU16(r *arm.Uint16X4, v0 *arm.Uint16X4, v1 *arm.Uint16X4)

// Subtract (vector). This instruction subtracts each vector element in the second source SIMD&FP register from the corresponding vector element in the first source SIMD&FP register, places the result into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VsubS8 VsubS8
//go:noescape
func VsubS8(r *arm.Int8X8, v0 *arm.Int8X8, v1 *arm.Int8X8)

// Floating-point Subtract (vector). This instruction subtracts the elements in the vector in the second source SIMD&FP register, from the corresponding elements in the vector in the first source SIMD&FP register, places each result into elements of a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VsubF32 VsubF32
//go:noescape
func VsubF32(r *arm.Float32X2, v0 *arm.Float32X2, v1 *arm.Float32X2)

// Subtract (vector). This instruction subtracts each vector element in the second source SIMD&FP register from the corresponding vector element in the first source SIMD&FP register, places the result into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VsubS32 VsubS32
//go:noescape
func VsubS32(r *arm.Int32X2, v0 *arm.Int32X2, v1 *arm.Int32X2)

// Subtract (vector). This instruction subtracts each vector element in the second source SIMD&FP register from the corresponding vector element in the first source SIMD&FP register, places the result into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VsubS64 VsubS64
//go:noescape
func VsubS64(r *arm.Int64X1, v0 *arm.Int64X1, v1 *arm.Int64X1)

// Subtract (vector). This instruction subtracts each vector element in the second source SIMD&FP register from the corresponding vector element in the first source SIMD&FP register, places the result into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VsubS16 VsubS16
//go:noescape
func VsubS16(r *arm.Int16X4, v0 *arm.Int16X4, v1 *arm.Int16X4)

// Subtract returning High Narrow. This instruction subtracts each vector element in the second source SIMD&FP register from the corresponding vector element in the first source SIMD&FP register, places the most significant half of the result into a vector, and writes the vector to the lower or upper half of the destination SIMD&FP register. All the values in this instruction are signed integer values.
//
//go:linkname VsubhnU32 VsubhnU32
//go:noescape
func VsubhnU32(r *arm.Uint16X4, v0 *arm.Uint32X4, v1 *arm.Uint32X4)

// Subtract returning High Narrow. This instruction subtracts each vector element in the second source SIMD&FP register from the corresponding vector element in the first source SIMD&FP register, places the most significant half of the result into a vector, and writes the vector to the lower or upper half of the destination SIMD&FP register. All the values in this instruction are signed integer values.
//
//go:linkname VsubhnU64 VsubhnU64
//go:noescape
func VsubhnU64(r *arm.Uint32X2, v0 *arm.Uint64X2, v1 *arm.Uint64X2)

// Subtract returning High Narrow. This instruction subtracts each vector element in the second source SIMD&FP register from the corresponding vector element in the first source SIMD&FP register, places the most significant half of the result into a vector, and writes the vector to the lower or upper half of the destination SIMD&FP register. All the values in this instruction are signed integer values.
//
//go:linkname VsubhnU16 VsubhnU16
//go:noescape
func VsubhnU16(r *arm.Uint8X8, v0 *arm.Uint16X8, v1 *arm.Uint16X8)

// Subtract returning High Narrow. This instruction subtracts each vector element in the second source SIMD&FP register from the corresponding vector element in the first source SIMD&FP register, places the most significant half of the result into a vector, and writes the vector to the lower or upper half of the destination SIMD&FP register. All the values in this instruction are signed integer values.
//
//go:linkname VsubhnS32 VsubhnS32
//go:noescape
func VsubhnS32(r *arm.Int16X4, v0 *arm.Int32X4, v1 *arm.Int32X4)

// Subtract returning High Narrow. This instruction subtracts each vector element in the second source SIMD&FP register from the corresponding vector element in the first source SIMD&FP register, places the most significant half of the result into a vector, and writes the vector to the lower or upper half of the destination SIMD&FP register. All the values in this instruction are signed integer values.
//
//go:linkname VsubhnS64 VsubhnS64
//go:noescape
func VsubhnS64(r *arm.Int32X2, v0 *arm.Int64X2, v1 *arm.Int64X2)

// Subtract returning High Narrow. This instruction subtracts each vector element in the second source SIMD&FP register from the corresponding vector element in the first source SIMD&FP register, places the most significant half of the result into a vector, and writes the vector to the lower or upper half of the destination SIMD&FP register. All the values in this instruction are signed integer values.
//
//go:linkname VsubhnS16 VsubhnS16
//go:noescape
func VsubhnS16(r *arm.Int8X8, v0 *arm.Int16X8, v1 *arm.Int16X8)

// Unsigned Subtract Long. This instruction subtracts each vector element in the lower or upper half of the second source SIMD&FP register from the corresponding vector element of the first source SIMD&FP register, places the result into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are unsigned integer values. The destination vector elements are twice as long as the source vector elements.
//
//go:linkname VsublU8 VsublU8
//go:noescape
func VsublU8(r *arm.Uint16X8, v0 *arm.Uint8X8, v1 *arm.Uint8X8)

// Unsigned Subtract Long. This instruction subtracts each vector element in the lower or upper half of the second source SIMD&FP register from the corresponding vector element of the first source SIMD&FP register, places the result into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are unsigned integer values. The destination vector elements are twice as long as the source vector elements.
//
//go:linkname VsublU32 VsublU32
//go:noescape
func VsublU32(r *arm.Uint64X2, v0 *arm.Uint32X2, v1 *arm.Uint32X2)

// Unsigned Subtract Long. This instruction subtracts each vector element in the lower or upper half of the second source SIMD&FP register from the corresponding vector element of the first source SIMD&FP register, places the result into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are unsigned integer values. The destination vector elements are twice as long as the source vector elements.
//
//go:linkname VsublU16 VsublU16
//go:noescape
func VsublU16(r *arm.Uint32X4, v0 *arm.Uint16X4, v1 *arm.Uint16X4)

// Signed Subtract Long. This instruction subtracts each vector element in the lower or upper half of the second source SIMD&FP register from the corresponding vector element of the first source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are signed integer values. The destination vector elements are twice as long as the source vector elements.
//
//go:linkname VsublS8 VsublS8
//go:noescape
func VsublS8(r *arm.Int16X8, v0 *arm.Int8X8, v1 *arm.Int8X8)

// Signed Subtract Long. This instruction subtracts each vector element in the lower or upper half of the second source SIMD&FP register from the corresponding vector element of the first source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are signed integer values. The destination vector elements are twice as long as the source vector elements.
//
//go:linkname VsublS32 VsublS32
//go:noescape
func VsublS32(r *arm.Int64X2, v0 *arm.Int32X2, v1 *arm.Int32X2)

// Signed Subtract Long. This instruction subtracts each vector element in the lower or upper half of the second source SIMD&FP register from the corresponding vector element of the first source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are signed integer values. The destination vector elements are twice as long as the source vector elements.
//
//go:linkname VsublS16 VsublS16
//go:noescape
func VsublS16(r *arm.Int32X4, v0 *arm.Int16X4, v1 *arm.Int16X4)

// Unsigned Subtract Wide. This instruction subtracts each vector element of the second source SIMD&FP register from the corresponding vector element in the lower or upper half of the first source SIMD&FP register, places the result in a vector, and writes the vector to the SIMD&FP destination register. All the values in this instruction are signed integer values.
//
//go:linkname VsubwU8 VsubwU8
//go:noescape
func VsubwU8(r *arm.Uint16X8, v0 *arm.Uint16X8, v1 *arm.Uint8X8)

// Unsigned Subtract Wide. This instruction subtracts each vector element of the second source SIMD&FP register from the corresponding vector element in the lower or upper half of the first source SIMD&FP register, places the result in a vector, and writes the vector to the SIMD&FP destination register. All the values in this instruction are signed integer values.
//
//go:linkname VsubwU32 VsubwU32
//go:noescape
func VsubwU32(r *arm.Uint64X2, v0 *arm.Uint64X2, v1 *arm.Uint32X2)

// Unsigned Subtract Wide. This instruction subtracts each vector element of the second source SIMD&FP register from the corresponding vector element in the lower or upper half of the first source SIMD&FP register, places the result in a vector, and writes the vector to the SIMD&FP destination register. All the values in this instruction are signed integer values.
//
//go:linkname VsubwU16 VsubwU16
//go:noescape
func VsubwU16(r *arm.Uint32X4, v0 *arm.Uint32X4, v1 *arm.Uint16X4)

// Signed Subtract Wide. This instruction subtracts each vector element in the lower or upper half of the second source SIMD&FP register from the corresponding vector element in the first source SIMD&FP register, places the result in a vector, and writes the vector to the SIMD&FP destination register. All the values in this instruction are signed integer values.
//
//go:linkname VsubwS8 VsubwS8
//go:noescape
func VsubwS8(r *arm.Int16X8, v0 *arm.Int16X8, v1 *arm.Int8X8)

// Signed Subtract Wide. This instruction subtracts each vector element in the lower or upper half of the second source SIMD&FP register from the corresponding vector element in the first source SIMD&FP register, places the result in a vector, and writes the vector to the SIMD&FP destination register. All the values in this instruction are signed integer values.
//
//go:linkname VsubwS32 VsubwS32
//go:noescape
func VsubwS32(r *arm.Int64X2, v0 *arm.Int64X2, v1 *arm.Int32X2)

// Signed Subtract Wide. This instruction subtracts each vector element in the lower or upper half of the second source SIMD&FP register from the corresponding vector element in the first source SIMD&FP register, places the result in a vector, and writes the vector to the SIMD&FP destination register. All the values in this instruction are signed integer values.
//
//go:linkname VsubwS16 VsubwS16
//go:noescape
func VsubwS16(r *arm.Int32X4, v0 *arm.Int32X4, v1 *arm.Int16X4)

// Table vector Lookup. This instruction reads each value from the vector elements in the index source SIMD&FP register, uses each result as an index to perform a lookup in a table of bytes that is described by one to four source table SIMD&FP registers, places the lookup result in a vector, and writes the vector to the destination SIMD&FP register. If an index is out of range for the table, the result for that lookup is 0. If more than one source register is used to describe the table, the first source register describes the lowest bytes of the table.
//
//go:linkname Vtbl1P8 Vtbl1P8
//go:noescape
func Vtbl1P8(r *arm.Poly8X8, v0 *arm.Poly8X8, v1 *arm.Uint8X8)

// Table vector Lookup. This instruction reads each value from the vector elements in the index source SIMD&FP register, uses each result as an index to perform a lookup in a table of bytes that is described by one to four source table SIMD&FP registers, places the lookup result in a vector, and writes the vector to the destination SIMD&FP register. If an index is out of range for the table, the result for that lookup is 0. If more than one source register is used to describe the table, the first source register describes the lowest bytes of the table.
//
//go:linkname Vtbl1U8 Vtbl1U8
//go:noescape
func Vtbl1U8(r *arm.Uint8X8, v0 *arm.Uint8X8, v1 *arm.Uint8X8)

// Table vector Lookup. This instruction reads each value from the vector elements in the index source SIMD&FP register, uses each result as an index to perform a lookup in a table of bytes that is described by one to four source table SIMD&FP registers, places the lookup result in a vector, and writes the vector to the destination SIMD&FP register. If an index is out of range for the table, the result for that lookup is 0. If more than one source register is used to describe the table, the first source register describes the lowest bytes of the table.
//
//go:linkname Vtbl1S8 Vtbl1S8
//go:noescape
func Vtbl1S8(r *arm.Int8X8, v0 *arm.Int8X8, v1 *arm.Int8X8)

// Table vector Lookup. This instruction reads each value from the vector elements in the index source SIMD&FP register, uses each result as an index to perform a lookup in a table of bytes that is described by one to four source table SIMD&FP registers, places the lookup result in a vector, and writes the vector to the destination SIMD&FP register. If an index is out of range for the table, the result for that lookup is 0. If more than one source register is used to describe the table, the first source register describes the lowest bytes of the table.
//
//go:linkname Vtbl2P8 Vtbl2P8
//go:noescape
func Vtbl2P8(r *arm.Poly8X8, v0 *arm.Poly8X8X2, v1 *arm.Uint8X8)

// Table vector Lookup. This instruction reads each value from the vector elements in the index source SIMD&FP register, uses each result as an index to perform a lookup in a table of bytes that is described by one to four source table SIMD&FP registers, places the lookup result in a vector, and writes the vector to the destination SIMD&FP register. If an index is out of range for the table, the result for that lookup is 0. If more than one source register is used to describe the table, the first source register describes the lowest bytes of the table.
//
//go:linkname Vtbl2U8 Vtbl2U8
//go:noescape
func Vtbl2U8(r *arm.Uint8X8, v0 *arm.Uint8X8X2, v1 *arm.Uint8X8)

// Table vector Lookup. This instruction reads each value from the vector elements in the index source SIMD&FP register, uses each result as an index to perform a lookup in a table of bytes that is described by one to four source table SIMD&FP registers, places the lookup result in a vector, and writes the vector to the destination SIMD&FP register. If an index is out of range for the table, the result for that lookup is 0. If more than one source register is used to describe the table, the first source register describes the lowest bytes of the table.
//
//go:linkname Vtbl2S8 Vtbl2S8
//go:noescape
func Vtbl2S8(r *arm.Int8X8, v0 *arm.Int8X8X2, v1 *arm.Int8X8)

// Table vector Lookup. This instruction reads each value from the vector elements in the index source SIMD&FP register, uses each result as an index to perform a lookup in a table of bytes that is described by one to four source table SIMD&FP registers, places the lookup result in a vector, and writes the vector to the destination SIMD&FP register. If an index is out of range for the table, the result for that lookup is 0. If more than one source register is used to describe the table, the first source register describes the lowest bytes of the table.
//
//go:linkname Vtbl3P8 Vtbl3P8
//go:noescape
func Vtbl3P8(r *arm.Poly8X8, v0 *arm.Poly8X8X3, v1 *arm.Uint8X8)

// Table vector Lookup. This instruction reads each value from the vector elements in the index source SIMD&FP register, uses each result as an index to perform a lookup in a table of bytes that is described by one to four source table SIMD&FP registers, places the lookup result in a vector, and writes the vector to the destination SIMD&FP register. If an index is out of range for the table, the result for that lookup is 0. If more than one source register is used to describe the table, the first source register describes the lowest bytes of the table.
//
//go:linkname Vtbl3U8 Vtbl3U8
//go:noescape
func Vtbl3U8(r *arm.Uint8X8, v0 *arm.Uint8X8X3, v1 *arm.Uint8X8)

// Table vector Lookup. This instruction reads each value from the vector elements in the index source SIMD&FP register, uses each result as an index to perform a lookup in a table of bytes that is described by one to four source table SIMD&FP registers, places the lookup result in a vector, and writes the vector to the destination SIMD&FP register. If an index is out of range for the table, the result for that lookup is 0. If more than one source register is used to describe the table, the first source register describes the lowest bytes of the table.
//
//go:linkname Vtbl3S8 Vtbl3S8
//go:noescape
func Vtbl3S8(r *arm.Int8X8, v0 *arm.Int8X8X3, v1 *arm.Int8X8)

// Table vector Lookup. This instruction reads each value from the vector elements in the index source SIMD&FP register, uses each result as an index to perform a lookup in a table of bytes that is described by one to four source table SIMD&FP registers, places the lookup result in a vector, and writes the vector to the destination SIMD&FP register. If an index is out of range for the table, the result for that lookup is 0. If more than one source register is used to describe the table, the first source register describes the lowest bytes of the table.
//
//go:linkname Vtbl4P8 Vtbl4P8
//go:noescape
func Vtbl4P8(r *arm.Poly8X8, v0 *arm.Poly8X8X4, v1 *arm.Uint8X8)

// Table vector Lookup. This instruction reads each value from the vector elements in the index source SIMD&FP register, uses each result as an index to perform a lookup in a table of bytes that is described by one to four source table SIMD&FP registers, places the lookup result in a vector, and writes the vector to the destination SIMD&FP register. If an index is out of range for the table, the result for that lookup is 0. If more than one source register is used to describe the table, the first source register describes the lowest bytes of the table.
//
//go:linkname Vtbl4U8 Vtbl4U8
//go:noescape
func Vtbl4U8(r *arm.Uint8X8, v0 *arm.Uint8X8X4, v1 *arm.Uint8X8)

// Table vector Lookup. This instruction reads each value from the vector elements in the index source SIMD&FP register, uses each result as an index to perform a lookup in a table of bytes that is described by one to four source table SIMD&FP registers, places the lookup result in a vector, and writes the vector to the destination SIMD&FP register. If an index is out of range for the table, the result for that lookup is 0. If more than one source register is used to describe the table, the first source register describes the lowest bytes of the table.
//
//go:linkname Vtbl4S8 Vtbl4S8
//go:noescape
func Vtbl4S8(r *arm.Int8X8, v0 *arm.Int8X8X4, v1 *arm.Int8X8)

// Table vector lookup extension
//
//go:linkname Vtbx1P8 Vtbx1P8
//go:noescape
func Vtbx1P8(r *arm.Poly8X8, v0 *arm.Poly8X8, v1 *arm.Poly8X8, v2 *arm.Uint8X8)

// Table vector lookup extension
//
//go:linkname Vtbx1U8 Vtbx1U8
//go:noescape
func Vtbx1U8(r *arm.Uint8X8, v0 *arm.Uint8X8, v1 *arm.Uint8X8, v2 *arm.Uint8X8)

// Table vector lookup extension
//
//go:linkname Vtbx1S8 Vtbx1S8
//go:noescape
func Vtbx1S8(r *arm.Int8X8, v0 *arm.Int8X8, v1 *arm.Int8X8, v2 *arm.Int8X8)

// Table vector lookup extension. This instruction reads each value from the vector elements in the index source SIMD&FP register, uses each result as an index to perform a lookup in a table of bytes that is described by one to four source table SIMD&FP registers, places the lookup result in a vector, and writes the vector to the destination SIMD&FP register. If an index is out of range for the table, the existing value in the vector element of the destination register is left unchanged. If more than one source register is used to describe the table, the first source register describes the lowest bytes of the table.
//
//go:linkname Vtbx2P8 Vtbx2P8
//go:noescape
func Vtbx2P8(r *arm.Poly8X8, v0 *arm.Poly8X8, v1 *arm.Poly8X8X2, v2 *arm.Uint8X8)

// Table vector lookup extension. This instruction reads each value from the vector elements in the index source SIMD&FP register, uses each result as an index to perform a lookup in a table of bytes that is described by one to four source table SIMD&FP registers, places the lookup result in a vector, and writes the vector to the destination SIMD&FP register. If an index is out of range for the table, the existing value in the vector element of the destination register is left unchanged. If more than one source register is used to describe the table, the first source register describes the lowest bytes of the table.
//
//go:linkname Vtbx2U8 Vtbx2U8
//go:noescape
func Vtbx2U8(r *arm.Uint8X8, v0 *arm.Uint8X8, v1 *arm.Uint8X8X2, v2 *arm.Uint8X8)

// Table vector lookup extension. This instruction reads each value from the vector elements in the index source SIMD&FP register, uses each result as an index to perform a lookup in a table of bytes that is described by one to four source table SIMD&FP registers, places the lookup result in a vector, and writes the vector to the destination SIMD&FP register. If an index is out of range for the table, the existing value in the vector element of the destination register is left unchanged. If more than one source register is used to describe the table, the first source register describes the lowest bytes of the table.
//
//go:linkname Vtbx2S8 Vtbx2S8
//go:noescape
func Vtbx2S8(r *arm.Int8X8, v0 *arm.Int8X8, v1 *arm.Int8X8X2, v2 *arm.Int8X8)

// Table vector lookup extension
//
//go:linkname Vtbx3P8 Vtbx3P8
//go:noescape
func Vtbx3P8(r *arm.Poly8X8, v0 *arm.Poly8X8, v1 *arm.Poly8X8X3, v2 *arm.Uint8X8)

// Table vector lookup extension
//
//go:linkname Vtbx3U8 Vtbx3U8
//go:noescape
func Vtbx3U8(r *arm.Uint8X8, v0 *arm.Uint8X8, v1 *arm.Uint8X8X3, v2 *arm.Uint8X8)

// Table vector lookup extension
//
//go:linkname Vtbx3S8 Vtbx3S8
//go:noescape
func Vtbx3S8(r *arm.Int8X8, v0 *arm.Int8X8, v1 *arm.Int8X8X3, v2 *arm.Int8X8)

// Table vector lookup extension. This instruction reads each value from the vector elements in the index source SIMD&FP register, uses each result as an index to perform a lookup in a table of bytes that is described by one to four source table SIMD&FP registers, places the lookup result in a vector, and writes the vector to the destination SIMD&FP register. If an index is out of range for the table, the existing value in the vector element of the destination register is left unchanged. If more than one source register is used to describe the table, the first source register describes the lowest bytes of the table.
//
//go:linkname Vtbx4P8 Vtbx4P8
//go:noescape
func Vtbx4P8(r *arm.Poly8X8, v0 *arm.Poly8X8, v1 *arm.Poly8X8X4, v2 *arm.Uint8X8)

// Table vector lookup extension. This instruction reads each value from the vector elements in the index source SIMD&FP register, uses each result as an index to perform a lookup in a table of bytes that is described by one to four source table SIMD&FP registers, places the lookup result in a vector, and writes the vector to the destination SIMD&FP register. If an index is out of range for the table, the existing value in the vector element of the destination register is left unchanged. If more than one source register is used to describe the table, the first source register describes the lowest bytes of the table.
//
//go:linkname Vtbx4U8 Vtbx4U8
//go:noescape
func Vtbx4U8(r *arm.Uint8X8, v0 *arm.Uint8X8, v1 *arm.Uint8X8X4, v2 *arm.Uint8X8)

// Table vector lookup extension. This instruction reads each value from the vector elements in the index source SIMD&FP register, uses each result as an index to perform a lookup in a table of bytes that is described by one to four source table SIMD&FP registers, places the lookup result in a vector, and writes the vector to the destination SIMD&FP register. If an index is out of range for the table, the existing value in the vector element of the destination register is left unchanged. If more than one source register is used to describe the table, the first source register describes the lowest bytes of the table.
//
//go:linkname Vtbx4S8 Vtbx4S8
//go:noescape
func Vtbx4S8(r *arm.Int8X8, v0 *arm.Int8X8, v1 *arm.Int8X8X4, v2 *arm.Int8X8)

// Transpose elements
//
//go:linkname VtrnP8 VtrnP8
//go:noescape
func VtrnP8(r *arm.Poly8X8X2, v0 *arm.Poly8X8, v1 *arm.Poly8X8)

// Transpose elements
//
//go:linkname VtrnP16 VtrnP16
//go:noescape
func VtrnP16(r *arm.Poly16X4X2, v0 *arm.Poly16X4, v1 *arm.Poly16X4)

// Transpose elements
//
//go:linkname VtrnqP8 VtrnqP8
//go:noescape
func VtrnqP8(r *arm.Poly8X16X2, v0 *arm.Poly8X16, v1 *arm.Poly8X16)

// Transpose elements
//
//go:linkname VtrnqP16 VtrnqP16
//go:noescape
func VtrnqP16(r *arm.Poly16X8X2, v0 *arm.Poly16X8, v1 *arm.Poly16X8)

// Transpose elements
//
//go:linkname VtrnqU8 VtrnqU8
//go:noescape
func VtrnqU8(r *arm.Uint8X16X2, v0 *arm.Uint8X16, v1 *arm.Uint8X16)

// Transpose elements
//
//go:linkname VtrnqU32 VtrnqU32
//go:noescape
func VtrnqU32(r *arm.Uint32X4X2, v0 *arm.Uint32X4, v1 *arm.Uint32X4)

// Transpose elements
//
//go:linkname VtrnqU16 VtrnqU16
//go:noescape
func VtrnqU16(r *arm.Uint16X8X2, v0 *arm.Uint16X8, v1 *arm.Uint16X8)

// Transpose elements
//
//go:linkname VtrnqS8 VtrnqS8
//go:noescape
func VtrnqS8(r *arm.Int8X16X2, v0 *arm.Int8X16, v1 *arm.Int8X16)

// Transpose elements
//
//go:linkname VtrnqF32 VtrnqF32
//go:noescape
func VtrnqF32(r *arm.Float32X4X2, v0 *arm.Float32X4, v1 *arm.Float32X4)

// Transpose elements
//
//go:linkname VtrnqS32 VtrnqS32
//go:noescape
func VtrnqS32(r *arm.Int32X4X2, v0 *arm.Int32X4, v1 *arm.Int32X4)

// Transpose elements
//
//go:linkname VtrnqS16 VtrnqS16
//go:noescape
func VtrnqS16(r *arm.Int16X8X2, v0 *arm.Int16X8, v1 *arm.Int16X8)

// Transpose elements
//
//go:linkname VtrnU8 VtrnU8
//go:noescape
func VtrnU8(r *arm.Uint8X8X2, v0 *arm.Uint8X8, v1 *arm.Uint8X8)

// Transpose elements
//
//go:linkname VtrnU32 VtrnU32
//go:noescape
func VtrnU32(r *arm.Uint32X2X2, v0 *arm.Uint32X2, v1 *arm.Uint32X2)

// Transpose elements
//
//go:linkname VtrnU16 VtrnU16
//go:noescape
func VtrnU16(r *arm.Uint16X4X2, v0 *arm.Uint16X4, v1 *arm.Uint16X4)

// Transpose elements
//
//go:linkname VtrnS8 VtrnS8
//go:noescape
func VtrnS8(r *arm.Int8X8X2, v0 *arm.Int8X8, v1 *arm.Int8X8)

// Transpose elements
//
//go:linkname VtrnF32 VtrnF32
//go:noescape
func VtrnF32(r *arm.Float32X2X2, v0 *arm.Float32X2, v1 *arm.Float32X2)

// Transpose elements
//
//go:linkname VtrnS32 VtrnS32
//go:noescape
func VtrnS32(r *arm.Int32X2X2, v0 *arm.Int32X2, v1 *arm.Int32X2)

// Transpose elements
//
//go:linkname VtrnS16 VtrnS16
//go:noescape
func VtrnS16(r *arm.Int16X4X2, v0 *arm.Int16X4, v1 *arm.Int16X4)

// Compare bitwise Test bits nonzero (vector). This instruction reads each vector element in the first source SIMD&FP register, performs an AND with the corresponding vector element in the second source SIMD&FP register, and if the result is not zero, sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VtstP8 VtstP8
//go:noescape
func VtstP8(r *arm.Uint8X8, v0 *arm.Poly8X8, v1 *arm.Poly8X8)

// vtst_p16
//
//go:linkname VtstP16 VtstP16
//go:noescape
func VtstP16(r *arm.Uint16X4, v0 *arm.Poly16X4, v1 *arm.Poly16X4)

// Compare bitwise Test bits nonzero (vector). This instruction reads each vector element in the first source SIMD&FP register, performs an AND with the corresponding vector element in the second source SIMD&FP register, and if the result is not zero, sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VtstqP8 VtstqP8
//go:noescape
func VtstqP8(r *arm.Uint8X16, v0 *arm.Poly8X16, v1 *arm.Poly8X16)

// vtstq_p16
//
//go:linkname VtstqP16 VtstqP16
//go:noescape
func VtstqP16(r *arm.Uint16X8, v0 *arm.Poly16X8, v1 *arm.Poly16X8)

// Compare bitwise Test bits nonzero (vector). This instruction reads each vector element in the first source SIMD&FP register, performs an AND with the corresponding vector element in the second source SIMD&FP register, and if the result is not zero, sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VtstqU8 VtstqU8
//go:noescape
func VtstqU8(r *arm.Uint8X16, v0 *arm.Uint8X16, v1 *arm.Uint8X16)

// Compare bitwise Test bits nonzero (vector). This instruction reads each vector element in the first source SIMD&FP register, performs an AND with the corresponding vector element in the second source SIMD&FP register, and if the result is not zero, sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VtstqU32 VtstqU32
//go:noescape
func VtstqU32(r *arm.Uint32X4, v0 *arm.Uint32X4, v1 *arm.Uint32X4)

// Compare bitwise Test bits nonzero (vector). This instruction reads each vector element in the first source SIMD&FP register, performs an AND with the corresponding vector element in the second source SIMD&FP register, and if the result is not zero, sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VtstqU16 VtstqU16
//go:noescape
func VtstqU16(r *arm.Uint16X8, v0 *arm.Uint16X8, v1 *arm.Uint16X8)

// Compare bitwise Test bits nonzero (vector). This instruction reads each vector element in the first source SIMD&FP register, performs an AND with the corresponding vector element in the second source SIMD&FP register, and if the result is not zero, sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VtstqS8 VtstqS8
//go:noescape
func VtstqS8(r *arm.Uint8X16, v0 *arm.Int8X16, v1 *arm.Int8X16)

// Compare bitwise Test bits nonzero (vector). This instruction reads each vector element in the first source SIMD&FP register, performs an AND with the corresponding vector element in the second source SIMD&FP register, and if the result is not zero, sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VtstqS32 VtstqS32
//go:noescape
func VtstqS32(r *arm.Uint32X4, v0 *arm.Int32X4, v1 *arm.Int32X4)

// Compare bitwise Test bits nonzero (vector). This instruction reads each vector element in the first source SIMD&FP register, performs an AND with the corresponding vector element in the second source SIMD&FP register, and if the result is not zero, sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VtstqS16 VtstqS16
//go:noescape
func VtstqS16(r *arm.Uint16X8, v0 *arm.Int16X8, v1 *arm.Int16X8)

// Compare bitwise Test bits nonzero (vector). This instruction reads each vector element in the first source SIMD&FP register, performs an AND with the corresponding vector element in the second source SIMD&FP register, and if the result is not zero, sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VtstU8 VtstU8
//go:noescape
func VtstU8(r *arm.Uint8X8, v0 *arm.Uint8X8, v1 *arm.Uint8X8)

// Compare bitwise Test bits nonzero (vector). This instruction reads each vector element in the first source SIMD&FP register, performs an AND with the corresponding vector element in the second source SIMD&FP register, and if the result is not zero, sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VtstU32 VtstU32
//go:noescape
func VtstU32(r *arm.Uint32X2, v0 *arm.Uint32X2, v1 *arm.Uint32X2)

// Compare bitwise Test bits nonzero (vector). This instruction reads each vector element in the first source SIMD&FP register, performs an AND with the corresponding vector element in the second source SIMD&FP register, and if the result is not zero, sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VtstU16 VtstU16
//go:noescape
func VtstU16(r *arm.Uint16X4, v0 *arm.Uint16X4, v1 *arm.Uint16X4)

// Compare bitwise Test bits nonzero (vector). This instruction reads each vector element in the first source SIMD&FP register, performs an AND with the corresponding vector element in the second source SIMD&FP register, and if the result is not zero, sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VtstS8 VtstS8
//go:noescape
func VtstS8(r *arm.Uint8X8, v0 *arm.Int8X8, v1 *arm.Int8X8)

// Compare bitwise Test bits nonzero (vector). This instruction reads each vector element in the first source SIMD&FP register, performs an AND with the corresponding vector element in the second source SIMD&FP register, and if the result is not zero, sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VtstS32 VtstS32
//go:noescape
func VtstS32(r *arm.Uint32X2, v0 *arm.Int32X2, v1 *arm.Int32X2)

// Compare bitwise Test bits nonzero (vector). This instruction reads each vector element in the first source SIMD&FP register, performs an AND with the corresponding vector element in the second source SIMD&FP register, and if the result is not zero, sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VtstS16 VtstS16
//go:noescape
func VtstS16(r *arm.Uint16X4, v0 *arm.Int16X4, v1 *arm.Int16X4)

// Unzip vectors (secondary). This instruction reads corresponding odd-numbered vector elements from the two source SIMD&FP registers, places the result from the first source register into consecutive elements in the lower half of a vector, and the result from the second source register into consecutive elements in the upper half of a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VuzpP8 VuzpP8
//go:noescape
func VuzpP8(r *arm.Poly8X8X2, v0 *arm.Poly8X8, v1 *arm.Poly8X8)

// Unzip vectors (secondary). This instruction reads corresponding odd-numbered vector elements from the two source SIMD&FP registers, places the result from the first source register into consecutive elements in the lower half of a vector, and the result from the second source register into consecutive elements in the upper half of a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VuzpP16 VuzpP16
//go:noescape
func VuzpP16(r *arm.Poly16X4X2, v0 *arm.Poly16X4, v1 *arm.Poly16X4)

// Unzip vectors (secondary). This instruction reads corresponding odd-numbered vector elements from the two source SIMD&FP registers, places the result from the first source register into consecutive elements in the lower half of a vector, and the result from the second source register into consecutive elements in the upper half of a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VuzpqP8 VuzpqP8
//go:noescape
func VuzpqP8(r *arm.Poly8X16X2, v0 *arm.Poly8X16, v1 *arm.Poly8X16)

// Unzip vectors (secondary). This instruction reads corresponding odd-numbered vector elements from the two source SIMD&FP registers, places the result from the first source register into consecutive elements in the lower half of a vector, and the result from the second source register into consecutive elements in the upper half of a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VuzpqP16 VuzpqP16
//go:noescape
func VuzpqP16(r *arm.Poly16X8X2, v0 *arm.Poly16X8, v1 *arm.Poly16X8)

// Unzip vectors (secondary). This instruction reads corresponding odd-numbered vector elements from the two source SIMD&FP registers, places the result from the first source register into consecutive elements in the lower half of a vector, and the result from the second source register into consecutive elements in the upper half of a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VuzpqU8 VuzpqU8
//go:noescape
func VuzpqU8(r *arm.Uint8X16X2, v0 *arm.Uint8X16, v1 *arm.Uint8X16)

// Unzip vectors (secondary). This instruction reads corresponding odd-numbered vector elements from the two source SIMD&FP registers, places the result from the first source register into consecutive elements in the lower half of a vector, and the result from the second source register into consecutive elements in the upper half of a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VuzpqU32 VuzpqU32
//go:noescape
func VuzpqU32(r *arm.Uint32X4X2, v0 *arm.Uint32X4, v1 *arm.Uint32X4)

// Unzip vectors (secondary). This instruction reads corresponding odd-numbered vector elements from the two source SIMD&FP registers, places the result from the first source register into consecutive elements in the lower half of a vector, and the result from the second source register into consecutive elements in the upper half of a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VuzpqU16 VuzpqU16
//go:noescape
func VuzpqU16(r *arm.Uint16X8X2, v0 *arm.Uint16X8, v1 *arm.Uint16X8)

// Unzip vectors (secondary). This instruction reads corresponding odd-numbered vector elements from the two source SIMD&FP registers, places the result from the first source register into consecutive elements in the lower half of a vector, and the result from the second source register into consecutive elements in the upper half of a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VuzpqS8 VuzpqS8
//go:noescape
func VuzpqS8(r *arm.Int8X16X2, v0 *arm.Int8X16, v1 *arm.Int8X16)

// Unzip vectors (secondary). This instruction reads corresponding odd-numbered vector elements from the two source SIMD&FP registers, places the result from the first source register into consecutive elements in the lower half of a vector, and the result from the second source register into consecutive elements in the upper half of a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VuzpqF32 VuzpqF32
//go:noescape
func VuzpqF32(r *arm.Float32X4X2, v0 *arm.Float32X4, v1 *arm.Float32X4)

// Unzip vectors (secondary). This instruction reads corresponding odd-numbered vector elements from the two source SIMD&FP registers, places the result from the first source register into consecutive elements in the lower half of a vector, and the result from the second source register into consecutive elements in the upper half of a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VuzpqS32 VuzpqS32
//go:noescape
func VuzpqS32(r *arm.Int32X4X2, v0 *arm.Int32X4, v1 *arm.Int32X4)

// Unzip vectors (secondary). This instruction reads corresponding odd-numbered vector elements from the two source SIMD&FP registers, places the result from the first source register into consecutive elements in the lower half of a vector, and the result from the second source register into consecutive elements in the upper half of a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VuzpqS16 VuzpqS16
//go:noescape
func VuzpqS16(r *arm.Int16X8X2, v0 *arm.Int16X8, v1 *arm.Int16X8)

// Unzip vectors (secondary). This instruction reads corresponding odd-numbered vector elements from the two source SIMD&FP registers, places the result from the first source register into consecutive elements in the lower half of a vector, and the result from the second source register into consecutive elements in the upper half of a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VuzpU8 VuzpU8
//go:noescape
func VuzpU8(r *arm.Uint8X8X2, v0 *arm.Uint8X8, v1 *arm.Uint8X8)

// Unzip vectors (secondary). This instruction reads corresponding odd-numbered vector elements from the two source SIMD&FP registers, places the result from the first source register into consecutive elements in the lower half of a vector, and the result from the second source register into consecutive elements in the upper half of a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VuzpU32 VuzpU32
//go:noescape
func VuzpU32(r *arm.Uint32X2X2, v0 *arm.Uint32X2, v1 *arm.Uint32X2)

// Unzip vectors (secondary). This instruction reads corresponding odd-numbered vector elements from the two source SIMD&FP registers, places the result from the first source register into consecutive elements in the lower half of a vector, and the result from the second source register into consecutive elements in the upper half of a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VuzpU16 VuzpU16
//go:noescape
func VuzpU16(r *arm.Uint16X4X2, v0 *arm.Uint16X4, v1 *arm.Uint16X4)

// Unzip vectors (secondary). This instruction reads corresponding odd-numbered vector elements from the two source SIMD&FP registers, places the result from the first source register into consecutive elements in the lower half of a vector, and the result from the second source register into consecutive elements in the upper half of a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VuzpS8 VuzpS8
//go:noescape
func VuzpS8(r *arm.Int8X8X2, v0 *arm.Int8X8, v1 *arm.Int8X8)

// Unzip vectors (secondary). This instruction reads corresponding odd-numbered vector elements from the two source SIMD&FP registers, places the result from the first source register into consecutive elements in the lower half of a vector, and the result from the second source register into consecutive elements in the upper half of a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VuzpF32 VuzpF32
//go:noescape
func VuzpF32(r *arm.Float32X2X2, v0 *arm.Float32X2, v1 *arm.Float32X2)

// Unzip vectors (secondary). This instruction reads corresponding odd-numbered vector elements from the two source SIMD&FP registers, places the result from the first source register into consecutive elements in the lower half of a vector, and the result from the second source register into consecutive elements in the upper half of a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VuzpS32 VuzpS32
//go:noescape
func VuzpS32(r *arm.Int32X2X2, v0 *arm.Int32X2, v1 *arm.Int32X2)

// Unzip vectors (secondary). This instruction reads corresponding odd-numbered vector elements from the two source SIMD&FP registers, places the result from the first source register into consecutive elements in the lower half of a vector, and the result from the second source register into consecutive elements in the upper half of a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VuzpS16 VuzpS16
//go:noescape
func VuzpS16(r *arm.Int16X4X2, v0 *arm.Int16X4, v1 *arm.Int16X4)

// Zip vectors (secondary). This instruction reads adjacent vector elements from the upper half of two source SIMD&FP registers as pairs, interleaves the pairs and places them into a vector, and writes the vector to the destination SIMD&FP register. The first pair from the first source register is placed into the two lowest vector elements, with subsequent pairs taken alternately from each source register.
//
//go:linkname VzipP8 VzipP8
//go:noescape
func VzipP8(r *arm.Poly8X8X2, v0 *arm.Poly8X8, v1 *arm.Poly8X8)

// Zip vectors (secondary). This instruction reads adjacent vector elements from the upper half of two source SIMD&FP registers as pairs, interleaves the pairs and places them into a vector, and writes the vector to the destination SIMD&FP register. The first pair from the first source register is placed into the two lowest vector elements, with subsequent pairs taken alternately from each source register.
//
//go:linkname VzipP16 VzipP16
//go:noescape
func VzipP16(r *arm.Poly16X4X2, v0 *arm.Poly16X4, v1 *arm.Poly16X4)

// Zip vectors (secondary). This instruction reads adjacent vector elements from the upper half of two source SIMD&FP registers as pairs, interleaves the pairs and places them into a vector, and writes the vector to the destination SIMD&FP register. The first pair from the first source register is placed into the two lowest vector elements, with subsequent pairs taken alternately from each source register.
//
//go:linkname VzipqP8 VzipqP8
//go:noescape
func VzipqP8(r *arm.Poly8X16X2, v0 *arm.Poly8X16, v1 *arm.Poly8X16)

// Zip vectors (secondary). This instruction reads adjacent vector elements from the upper half of two source SIMD&FP registers as pairs, interleaves the pairs and places them into a vector, and writes the vector to the destination SIMD&FP register. The first pair from the first source register is placed into the two lowest vector elements, with subsequent pairs taken alternately from each source register.
//
//go:linkname VzipqP16 VzipqP16
//go:noescape
func VzipqP16(r *arm.Poly16X8X2, v0 *arm.Poly16X8, v1 *arm.Poly16X8)

// Zip vectors (secondary). This instruction reads adjacent vector elements from the upper half of two source SIMD&FP registers as pairs, interleaves the pairs and places them into a vector, and writes the vector to the destination SIMD&FP register. The first pair from the first source register is placed into the two lowest vector elements, with subsequent pairs taken alternately from each source register.
//
//go:linkname VzipqU8 VzipqU8
//go:noescape
func VzipqU8(r *arm.Uint8X16X2, v0 *arm.Uint8X16, v1 *arm.Uint8X16)

// Zip vectors (secondary). This instruction reads adjacent vector elements from the upper half of two source SIMD&FP registers as pairs, interleaves the pairs and places them into a vector, and writes the vector to the destination SIMD&FP register. The first pair from the first source register is placed into the two lowest vector elements, with subsequent pairs taken alternately from each source register.
//
//go:linkname VzipqU32 VzipqU32
//go:noescape
func VzipqU32(r *arm.Uint32X4X2, v0 *arm.Uint32X4, v1 *arm.Uint32X4)

// Zip vectors (secondary). This instruction reads adjacent vector elements from the upper half of two source SIMD&FP registers as pairs, interleaves the pairs and places them into a vector, and writes the vector to the destination SIMD&FP register. The first pair from the first source register is placed into the two lowest vector elements, with subsequent pairs taken alternately from each source register.
//
//go:linkname VzipqU16 VzipqU16
//go:noescape
func VzipqU16(r *arm.Uint16X8X2, v0 *arm.Uint16X8, v1 *arm.Uint16X8)

// Zip vectors (secondary). This instruction reads adjacent vector elements from the upper half of two source SIMD&FP registers as pairs, interleaves the pairs and places them into a vector, and writes the vector to the destination SIMD&FP register. The first pair from the first source register is placed into the two lowest vector elements, with subsequent pairs taken alternately from each source register.
//
//go:linkname VzipqS8 VzipqS8
//go:noescape
func VzipqS8(r *arm.Int8X16X2, v0 *arm.Int8X16, v1 *arm.Int8X16)

// Zip vectors (secondary). This instruction reads adjacent vector elements from the upper half of two source SIMD&FP registers as pairs, interleaves the pairs and places them into a vector, and writes the vector to the destination SIMD&FP register. The first pair from the first source register is placed into the two lowest vector elements, with subsequent pairs taken alternately from each source register.
//
//go:linkname VzipqF32 VzipqF32
//go:noescape
func VzipqF32(r *arm.Float32X4X2, v0 *arm.Float32X4, v1 *arm.Float32X4)

// Zip vectors (secondary). This instruction reads adjacent vector elements from the upper half of two source SIMD&FP registers as pairs, interleaves the pairs and places them into a vector, and writes the vector to the destination SIMD&FP register. The first pair from the first source register is placed into the two lowest vector elements, with subsequent pairs taken alternately from each source register.
//
//go:linkname VzipqS32 VzipqS32
//go:noescape
func VzipqS32(r *arm.Int32X4X2, v0 *arm.Int32X4, v1 *arm.Int32X4)

// Zip vectors (secondary). This instruction reads adjacent vector elements from the upper half of two source SIMD&FP registers as pairs, interleaves the pairs and places them into a vector, and writes the vector to the destination SIMD&FP register. The first pair from the first source register is placed into the two lowest vector elements, with subsequent pairs taken alternately from each source register.
//
//go:linkname VzipqS16 VzipqS16
//go:noescape
func VzipqS16(r *arm.Int16X8X2, v0 *arm.Int16X8, v1 *arm.Int16X8)

// Zip vectors (secondary). This instruction reads adjacent vector elements from the upper half of two source SIMD&FP registers as pairs, interleaves the pairs and places them into a vector, and writes the vector to the destination SIMD&FP register. The first pair from the first source register is placed into the two lowest vector elements, with subsequent pairs taken alternately from each source register.
//
//go:linkname VzipU8 VzipU8
//go:noescape
func VzipU8(r *arm.Uint8X8X2, v0 *arm.Uint8X8, v1 *arm.Uint8X8)

// Zip vectors (secondary). This instruction reads adjacent vector elements from the upper half of two source SIMD&FP registers as pairs, interleaves the pairs and places them into a vector, and writes the vector to the destination SIMD&FP register. The first pair from the first source register is placed into the two lowest vector elements, with subsequent pairs taken alternately from each source register.
//
//go:linkname VzipU32 VzipU32
//go:noescape
func VzipU32(r *arm.Uint32X2X2, v0 *arm.Uint32X2, v1 *arm.Uint32X2)

// Zip vectors (secondary). This instruction reads adjacent vector elements from the upper half of two source SIMD&FP registers as pairs, interleaves the pairs and places them into a vector, and writes the vector to the destination SIMD&FP register. The first pair from the first source register is placed into the two lowest vector elements, with subsequent pairs taken alternately from each source register.
//
//go:linkname VzipU16 VzipU16
//go:noescape
func VzipU16(r *arm.Uint16X4X2, v0 *arm.Uint16X4, v1 *arm.Uint16X4)

// Zip vectors (secondary). This instruction reads adjacent vector elements from the upper half of two source SIMD&FP registers as pairs, interleaves the pairs and places them into a vector, and writes the vector to the destination SIMD&FP register. The first pair from the first source register is placed into the two lowest vector elements, with subsequent pairs taken alternately from each source register.
//
//go:linkname VzipS8 VzipS8
//go:noescape
func VzipS8(r *arm.Int8X8X2, v0 *arm.Int8X8, v1 *arm.Int8X8)

// Zip vectors (secondary). This instruction reads adjacent vector elements from the upper half of two source SIMD&FP registers as pairs, interleaves the pairs and places them into a vector, and writes the vector to the destination SIMD&FP register. The first pair from the first source register is placed into the two lowest vector elements, with subsequent pairs taken alternately from each source register.
//
//go:linkname VzipF32 VzipF32
//go:noescape
func VzipF32(r *arm.Float32X2X2, v0 *arm.Float32X2, v1 *arm.Float32X2)

// Zip vectors (secondary). This instruction reads adjacent vector elements from the upper half of two source SIMD&FP registers as pairs, interleaves the pairs and places them into a vector, and writes the vector to the destination SIMD&FP register. The first pair from the first source register is placed into the two lowest vector elements, with subsequent pairs taken alternately from each source register.
//
//go:linkname VzipS32 VzipS32
//go:noescape
func VzipS32(r *arm.Int32X2X2, v0 *arm.Int32X2, v1 *arm.Int32X2)

// Zip vectors (secondary). This instruction reads adjacent vector elements from the upper half of two source SIMD&FP registers as pairs, interleaves the pairs and places them into a vector, and writes the vector to the destination SIMD&FP register. The first pair from the first source register is placed into the two lowest vector elements, with subsequent pairs taken alternately from each source register.
//
//go:linkname VzipS16 VzipS16
//go:noescape
func VzipS16(r *arm.Int16X4X2, v0 *arm.Int16X4, v1 *arm.Int16X4)

// Floating-point Convert to Signed integer, rounding to nearest with ties to Away (vector). This instruction converts each element in a vector from a floating-point value to a signed integer value using the Round to Nearest with Ties to Away rounding mode and writes the result to the SIMD&FP destination register.
//
//go:linkname VcvtaqS32F32 VcvtaqS32F32
//go:noescape
func VcvtaqS32F32(r *arm.Int32X4, v0 *arm.Float32X4)

// Floating-point Convert to Signed integer, rounding to nearest with ties to Away (vector). This instruction converts each element in a vector from a floating-point value to a signed integer value using the Round to Nearest with Ties to Away rounding mode and writes the result to the SIMD&FP destination register.
//
//go:linkname VcvtaS32F32 VcvtaS32F32
//go:noescape
func VcvtaS32F32(r *arm.Int32X2, v0 *arm.Float32X2)

// Floating-point Convert to Unsigned integer, rounding to nearest with ties to Away (vector). This instruction converts each element in a vector from a floating-point value to an unsigned integer value using the Round to Nearest with Ties to Away rounding mode and writes the result to the SIMD&FP destination register.
//
//go:linkname VcvtaqU32F32 VcvtaqU32F32
//go:noescape
func VcvtaqU32F32(r *arm.Uint32X4, v0 *arm.Float32X4)

// Floating-point Convert to Unsigned integer, rounding to nearest with ties to Away (vector). This instruction converts each element in a vector from a floating-point value to an unsigned integer value using the Round to Nearest with Ties to Away rounding mode and writes the result to the SIMD&FP destination register.
//
//go:linkname VcvtaU32F32 VcvtaU32F32
//go:noescape
func VcvtaU32F32(r *arm.Uint32X2, v0 *arm.Float32X2)

// Floating-point Convert to Signed integer, rounding toward Minus infinity (vector). This instruction converts a scalar or each element in a vector from a floating-point value to a signed integer value using the Round towards Minus Infinity rounding mode, and writes the result to the SIMD&FP destination register.
//
//go:linkname VcvtmqS32F32 VcvtmqS32F32
//go:noescape
func VcvtmqS32F32(r *arm.Int32X4, v0 *arm.Float32X4)

// Floating-point Convert to Signed integer, rounding toward Minus infinity (vector). This instruction converts a scalar or each element in a vector from a floating-point value to a signed integer value using the Round towards Minus Infinity rounding mode, and writes the result to the SIMD&FP destination register.
//
//go:linkname VcvtmS32F32 VcvtmS32F32
//go:noescape
func VcvtmS32F32(r *arm.Int32X2, v0 *arm.Float32X2)

// Floating-point Convert to Unsigned integer, rounding toward Minus infinity (vector). This instruction converts a scalar or each element in a vector from a floating-point value to an unsigned integer value using the Round towards Minus Infinity rounding mode, and writes the result to the SIMD&FP destination register.
//
//go:linkname VcvtmqU32F32 VcvtmqU32F32
//go:noescape
func VcvtmqU32F32(r *arm.Uint32X4, v0 *arm.Float32X4)

// Floating-point Convert to Unsigned integer, rounding toward Minus infinity (vector). This instruction converts a scalar or each element in a vector from a floating-point value to an unsigned integer value using the Round towards Minus Infinity rounding mode, and writes the result to the SIMD&FP destination register.
//
//go:linkname VcvtmU32F32 VcvtmU32F32
//go:noescape
func VcvtmU32F32(r *arm.Uint32X2, v0 *arm.Float32X2)

// Floating-point Convert to Signed integer, rounding to nearest with ties to even (vector). This instruction converts a scalar or each element in a vector from a floating-point value to a signed integer value using the Round to Nearest rounding mode, and writes the result to the SIMD&FP destination register.
//
//go:linkname VcvtnqS32F32 VcvtnqS32F32
//go:noescape
func VcvtnqS32F32(r *arm.Int32X4, v0 *arm.Float32X4)

// Floating-point Convert to Signed integer, rounding to nearest with ties to even (vector). This instruction converts a scalar or each element in a vector from a floating-point value to a signed integer value using the Round to Nearest rounding mode, and writes the result to the SIMD&FP destination register.
//
//go:linkname VcvtnS32F32 VcvtnS32F32
//go:noescape
func VcvtnS32F32(r *arm.Int32X2, v0 *arm.Float32X2)

// Floating-point Convert to Unsigned integer, rounding to nearest with ties to even (vector). This instruction converts a scalar or each element in a vector from a floating-point value to an unsigned integer value using the Round to Nearest rounding mode, and writes the result to the SIMD&FP destination register.
//
//go:linkname VcvtnqU32F32 VcvtnqU32F32
//go:noescape
func VcvtnqU32F32(r *arm.Uint32X4, v0 *arm.Float32X4)

// Floating-point Convert to Unsigned integer, rounding to nearest with ties to even (vector). This instruction converts a scalar or each element in a vector from a floating-point value to an unsigned integer value using the Round to Nearest rounding mode, and writes the result to the SIMD&FP destination register.
//
//go:linkname VcvtnU32F32 VcvtnU32F32
//go:noescape
func VcvtnU32F32(r *arm.Uint32X2, v0 *arm.Float32X2)

// Floating-point Convert to Signed integer, rounding toward Plus infinity (vector). This instruction converts a scalar or each element in a vector from a floating-point value to a signed integer value using the Round towards Plus Infinity rounding mode, and writes the result to the SIMD&FP destination register.
//
//go:linkname VcvtpqS32F32 VcvtpqS32F32
//go:noescape
func VcvtpqS32F32(r *arm.Int32X4, v0 *arm.Float32X4)

// Floating-point Convert to Signed integer, rounding toward Plus infinity (vector). This instruction converts a scalar or each element in a vector from a floating-point value to a signed integer value using the Round towards Plus Infinity rounding mode, and writes the result to the SIMD&FP destination register.
//
//go:linkname VcvtpS32F32 VcvtpS32F32
//go:noescape
func VcvtpS32F32(r *arm.Int32X2, v0 *arm.Float32X2)

// Floating-point Convert to Unsigned integer, rounding toward Plus infinity (vector). This instruction converts a scalar or each element in a vector from a floating-point value to an unsigned integer value using the Round towards Plus Infinity rounding mode, and writes the result to the SIMD&FP destination register.
//
//go:linkname VcvtpqU32F32 VcvtpqU32F32
//go:noescape
func VcvtpqU32F32(r *arm.Uint32X4, v0 *arm.Float32X4)

// Floating-point Convert to Unsigned integer, rounding toward Plus infinity (vector). This instruction converts a scalar or each element in a vector from a floating-point value to an unsigned integer value using the Round towards Plus Infinity rounding mode, and writes the result to the SIMD&FP destination register.
//
//go:linkname VcvtpU32F32 VcvtpU32F32
//go:noescape
func VcvtpU32F32(r *arm.Uint32X2, v0 *arm.Float32X2)

// AES single round decryption.
//
//go:linkname VaesdqU8 VaesdqU8
//go:noescape
func VaesdqU8(r *arm.Uint8X16, v0 *arm.Uint8X16, v1 *arm.Uint8X16)

// AES single round encryption.
//
//go:linkname VaeseqU8 VaeseqU8
//go:noescape
func VaeseqU8(r *arm.Uint8X16, v0 *arm.Uint8X16, v1 *arm.Uint8X16)

// AES inverse mix columns.
//
//go:linkname VaesimcqU8 VaesimcqU8
//go:noescape
func VaesimcqU8(r *arm.Uint8X16, v0 *arm.Uint8X16)

// AES mix columns.
//
//go:linkname VaesmcqU8 VaesmcqU8
//go:noescape
func VaesmcqU8(r *arm.Uint8X16, v0 *arm.Uint8X16)

// Floating-point Round to Integral, toward Zero (vector). This instruction rounds a vector of floating-point values in the SIMD&FP source register to integral floating-point values of the same size using the Round towards Zero rounding mode, and writes the result to the SIMD&FP destination register.
//
//go:linkname VrndqF32 VrndqF32
//go:noescape
func VrndqF32(r *arm.Float32X4, v0 *arm.Float32X4)

// Floating-point Round to Integral, toward Zero (vector). This instruction rounds a vector of floating-point values in the SIMD&FP source register to integral floating-point values of the same size using the Round towards Zero rounding mode, and writes the result to the SIMD&FP destination register.
//
//go:linkname VrndF32 VrndF32
//go:noescape
func VrndF32(r *arm.Float32X2, v0 *arm.Float32X2)

// Floating-point Round to Integral, to nearest with ties to Away (vector). This instruction rounds a vector of floating-point values in the SIMD&FP source register to integral floating-point values of the same size using the Round to Nearest with Ties to Away rounding mode, and writes the result to the SIMD&FP destination register.
//
//go:linkname VrndaqF32 VrndaqF32
//go:noescape
func VrndaqF32(r *arm.Float32X4, v0 *arm.Float32X4)

// Floating-point Round to Integral, to nearest with ties to Away (vector). This instruction rounds a vector of floating-point values in the SIMD&FP source register to integral floating-point values of the same size using the Round to Nearest with Ties to Away rounding mode, and writes the result to the SIMD&FP destination register.
//
//go:linkname VrndaF32 VrndaF32
//go:noescape
func VrndaF32(r *arm.Float32X2, v0 *arm.Float32X2)

// Floating-point Round to Integral, using current rounding mode (vector). This instruction rounds a vector of floating-point values in the SIMD&FP source register to integral floating-point values of the same size using the rounding mode that is determined by the FPCR, and writes the result to the SIMD&FP destination register.
//
//go:linkname VrndiqF32 VrndiqF32
//go:noescape
func VrndiqF32(r *arm.Float32X4, v0 *arm.Float32X4)

// Floating-point Round to Integral, using current rounding mode (vector). This instruction rounds a vector of floating-point values in the SIMD&FP source register to integral floating-point values of the same size using the rounding mode that is determined by the FPCR, and writes the result to the SIMD&FP destination register.
//
//go:linkname VrndiF32 VrndiF32
//go:noescape
func VrndiF32(r *arm.Float32X2, v0 *arm.Float32X2)

// Floating-point Round to Integral, toward Minus infinity (vector). This instruction rounds a vector of floating-point values in the SIMD&FP source register to integral floating-point values of the same size using the Round towards Minus Infinity rounding mode, and writes the result to the SIMD&FP destination register.
//
//go:linkname VrndmqF32 VrndmqF32
//go:noescape
func VrndmqF32(r *arm.Float32X4, v0 *arm.Float32X4)

// Floating-point Round to Integral, toward Minus infinity (vector). This instruction rounds a vector of floating-point values in the SIMD&FP source register to integral floating-point values of the same size using the Round towards Minus Infinity rounding mode, and writes the result to the SIMD&FP destination register.
//
//go:linkname VrndmF32 VrndmF32
//go:noescape
func VrndmF32(r *arm.Float32X2, v0 *arm.Float32X2)

// Floating-point Round to Integral, to nearest with ties to even (vector). This instruction rounds a vector of floating-point values in the SIMD&FP source register to integral floating-point values of the same size using the Round to Nearest rounding mode, and writes the result to the SIMD&FP destination register.
//
//go:linkname VrndnqF32 VrndnqF32
//go:noescape
func VrndnqF32(r *arm.Float32X4, v0 *arm.Float32X4)

// Floating-point Round to Integral, to nearest with ties to even (vector). This instruction rounds a vector of floating-point values in the SIMD&FP source register to integral floating-point values of the same size using the Round to Nearest rounding mode, and writes the result to the SIMD&FP destination register.
//
//go:linkname VrndnF32 VrndnF32
//go:noescape
func VrndnF32(r *arm.Float32X2, v0 *arm.Float32X2)

// Floating-point Round to Integral, to nearest with ties to even (vector). This instruction rounds a vector of floating-point values in the SIMD&FP source register to integral floating-point values of the same size using the Round to Nearest rounding mode, and writes the result to the SIMD&FP destination register.
//
//go:linkname VrndnsF32 VrndnsF32
//go:noescape
func VrndnsF32(r *arm.Float32, v0 *arm.Float32)

// Floating-point Round to Integral, toward Plus infinity (vector). This instruction rounds a vector of floating-point values in the SIMD&FP source register to integral floating-point values of the same size using the Round towards Plus Infinity rounding mode, and writes the result to the SIMD&FP destination register.
//
//go:linkname VrndpqF32 VrndpqF32
//go:noescape
func VrndpqF32(r *arm.Float32X4, v0 *arm.Float32X4)

// Floating-point Round to Integral, toward Plus infinity (vector). This instruction rounds a vector of floating-point values in the SIMD&FP source register to integral floating-point values of the same size using the Round towards Plus Infinity rounding mode, and writes the result to the SIMD&FP destination register.
//
//go:linkname VrndpF32 VrndpF32
//go:noescape
func VrndpF32(r *arm.Float32X2, v0 *arm.Float32X2)

// Floating-point Round to Integral exact, using current rounding mode (vector). This instruction rounds a vector of floating-point values in the SIMD&FP source register to integral floating-point values of the same size using the rounding mode that is determined by the FPCR, and writes the result to the SIMD&FP destination register.
//
//go:linkname VrndxqF32 VrndxqF32
//go:noescape
func VrndxqF32(r *arm.Float32X4, v0 *arm.Float32X4)

// Floating-point Round to Integral exact, using current rounding mode (vector). This instruction rounds a vector of floating-point values in the SIMD&FP source register to integral floating-point values of the same size using the rounding mode that is determined by the FPCR, and writes the result to the SIMD&FP destination register.
//
//go:linkname VrndxF32 VrndxF32
//go:noescape
func VrndxF32(r *arm.Float32X2, v0 *arm.Float32X2)

// Floating-point Maximum Number (vector). This instruction compares corresponding vector elements in the two source SIMD&FP registers, writes the larger of the two floating-point values into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VmaxnmqF32 VmaxnmqF32
//go:noescape
func VmaxnmqF32(r *arm.Float32X4, v0 *arm.Float32X4, v1 *arm.Float32X4)

// Floating-point Maximum Number (vector). This instruction compares corresponding vector elements in the two source SIMD&FP registers, writes the larger of the two floating-point values into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VmaxnmF32 VmaxnmF32
//go:noescape
func VmaxnmF32(r *arm.Float32X2, v0 *arm.Float32X2, v1 *arm.Float32X2)

// Floating-point Minimum Number (vector). This instruction compares corresponding vector elements in the two source SIMD&FP registers, writes the smaller of the two floating-point values into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VminnmqF32 VminnmqF32
//go:noescape
func VminnmqF32(r *arm.Float32X4, v0 *arm.Float32X4, v1 *arm.Float32X4)

// Floating-point Minimum Number (vector). This instruction compares corresponding vector elements in the two source SIMD&FP registers, writes the smaller of the two floating-point values into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VminnmF32 VminnmF32
//go:noescape
func VminnmF32(r *arm.Float32X2, v0 *arm.Float32X2, v1 *arm.Float32X2)

// SHA1 hash update (choose).
//
//go:linkname Vsha1CqU32 Vsha1CqU32
//go:noescape
func Vsha1CqU32(r *arm.Uint32X4, v0 *arm.Uint32X4, v1 *arm.Uint32, v2 *arm.Uint32X4)

// SHA1 fixed rotate.
//
//go:linkname Vsha1HU32 Vsha1HU32
//go:noescape
func Vsha1HU32(r *arm.Uint32, v0 *arm.Uint32)

// SHA1 hash update (majority).
//
//go:linkname Vsha1MqU32 Vsha1MqU32
//go:noescape
func Vsha1MqU32(r *arm.Uint32X4, v0 *arm.Uint32X4, v1 *arm.Uint32, v2 *arm.Uint32X4)

// SHA1 hash update (parity).
//
//go:linkname Vsha1PqU32 Vsha1PqU32
//go:noescape
func Vsha1PqU32(r *arm.Uint32X4, v0 *arm.Uint32X4, v1 *arm.Uint32, v2 *arm.Uint32X4)

// SHA1 schedule update 0.
//
//go:linkname Vsha1Su0QU32 Vsha1Su0QU32
//go:noescape
func Vsha1Su0QU32(r *arm.Uint32X4, v0 *arm.Uint32X4, v1 *arm.Uint32X4, v2 *arm.Uint32X4)

// SHA1 schedule update 1.
//
//go:linkname Vsha1Su1QU32 Vsha1Su1QU32
//go:noescape
func Vsha1Su1QU32(r *arm.Uint32X4, v0 *arm.Uint32X4, v1 *arm.Uint32X4)

// SHA256 hash update (part 1).
//
//go:linkname Vsha256HqU32 Vsha256HqU32
//go:noescape
func Vsha256HqU32(r *arm.Uint32X4, v0 *arm.Uint32X4, v1 *arm.Uint32X4, v2 *arm.Uint32X4)

// SHA256 hash update (part 2).
//
//go:linkname Vsha256H2QU32 Vsha256H2QU32
//go:noescape
func Vsha256H2QU32(r *arm.Uint32X4, v0 *arm.Uint32X4, v1 *arm.Uint32X4, v2 *arm.Uint32X4)

// SHA256 schedule update 0.
//
//go:linkname Vsha256Su0QU32 Vsha256Su0QU32
//go:noescape
func Vsha256Su0QU32(r *arm.Uint32X4, v0 *arm.Uint32X4, v1 *arm.Uint32X4)

// SHA256 schedule update 1.
//
//go:linkname Vsha256Su1QU32 Vsha256Su1QU32
//go:noescape
func Vsha256Su1QU32(r *arm.Uint32X4, v0 *arm.Uint32X4, v1 *arm.Uint32X4, v2 *arm.Uint32X4)

// Bit Clear and Exclusive OR performs a bitwise AND of the 128-bit vector in a source SIMD&FP register and the complement of the vector in another source SIMD&FP register, then performs a bitwise exclusive OR of the resulting vector and the vector in a third source SIMD&FP register, and writes the result to the destination SIMD&FP register.
//
//go:linkname VbcaxqU8 VbcaxqU8
//go:noescape
func VbcaxqU8(r *arm.Uint8X16, v0 *arm.Uint8X16, v1 *arm.Uint8X16, v2 *arm.Uint8X16)

// Bit Clear and Exclusive OR performs a bitwise AND of the 128-bit vector in a source SIMD&FP register and the complement of the vector in another source SIMD&FP register, then performs a bitwise exclusive OR of the resulting vector and the vector in a third source SIMD&FP register, and writes the result to the destination SIMD&FP register.
//
//go:linkname VbcaxqU32 VbcaxqU32
//go:noescape
func VbcaxqU32(r *arm.Uint32X4, v0 *arm.Uint32X4, v1 *arm.Uint32X4, v2 *arm.Uint32X4)

// Bit Clear and Exclusive OR performs a bitwise AND of the 128-bit vector in a source SIMD&FP register and the complement of the vector in another source SIMD&FP register, then performs a bitwise exclusive OR of the resulting vector and the vector in a third source SIMD&FP register, and writes the result to the destination SIMD&FP register.
//
//go:linkname VbcaxqU64 VbcaxqU64
//go:noescape
func VbcaxqU64(r *arm.Uint64X2, v0 *arm.Uint64X2, v1 *arm.Uint64X2, v2 *arm.Uint64X2)

// Bit Clear and Exclusive OR performs a bitwise AND of the 128-bit vector in a source SIMD&FP register and the complement of the vector in another source SIMD&FP register, then performs a bitwise exclusive OR of the resulting vector and the vector in a third source SIMD&FP register, and writes the result to the destination SIMD&FP register.
//
//go:linkname VbcaxqU16 VbcaxqU16
//go:noescape
func VbcaxqU16(r *arm.Uint16X8, v0 *arm.Uint16X8, v1 *arm.Uint16X8, v2 *arm.Uint16X8)

// Bit Clear and Exclusive OR performs a bitwise AND of the 128-bit vector in a source SIMD&FP register and the complement of the vector in another source SIMD&FP register, then performs a bitwise exclusive OR of the resulting vector and the vector in a third source SIMD&FP register, and writes the result to the destination SIMD&FP register.
//
//go:linkname VbcaxqS8 VbcaxqS8
//go:noescape
func VbcaxqS8(r *arm.Int8X16, v0 *arm.Int8X16, v1 *arm.Int8X16, v2 *arm.Int8X16)

// Bit Clear and Exclusive OR performs a bitwise AND of the 128-bit vector in a source SIMD&FP register and the complement of the vector in another source SIMD&FP register, then performs a bitwise exclusive OR of the resulting vector and the vector in a third source SIMD&FP register, and writes the result to the destination SIMD&FP register.
//
//go:linkname VbcaxqS32 VbcaxqS32
//go:noescape
func VbcaxqS32(r *arm.Int32X4, v0 *arm.Int32X4, v1 *arm.Int32X4, v2 *arm.Int32X4)

// Bit Clear and Exclusive OR performs a bitwise AND of the 128-bit vector in a source SIMD&FP register and the complement of the vector in another source SIMD&FP register, then performs a bitwise exclusive OR of the resulting vector and the vector in a third source SIMD&FP register, and writes the result to the destination SIMD&FP register.
//
//go:linkname VbcaxqS64 VbcaxqS64
//go:noescape
func VbcaxqS64(r *arm.Int64X2, v0 *arm.Int64X2, v1 *arm.Int64X2, v2 *arm.Int64X2)

// Bit Clear and Exclusive OR performs a bitwise AND of the 128-bit vector in a source SIMD&FP register and the complement of the vector in another source SIMD&FP register, then performs a bitwise exclusive OR of the resulting vector and the vector in a third source SIMD&FP register, and writes the result to the destination SIMD&FP register.
//
//go:linkname VbcaxqS16 VbcaxqS16
//go:noescape
func VbcaxqS16(r *arm.Int16X8, v0 *arm.Int16X8, v1 *arm.Int16X8, v2 *arm.Int16X8)

// Three-way Exclusive OR performs a three-way exclusive OR of the values in the three source SIMD&FP registers, and writes the result to the destination SIMD&FP register.
//
//go:linkname Veor3QU8 Veor3QU8
//go:noescape
func Veor3QU8(r *arm.Uint8X16, v0 *arm.Uint8X16, v1 *arm.Uint8X16, v2 *arm.Uint8X16)

// Three-way Exclusive OR performs a three-way exclusive OR of the values in the three source SIMD&FP registers, and writes the result to the destination SIMD&FP register.
//
//go:linkname Veor3QU32 Veor3QU32
//go:noescape
func Veor3QU32(r *arm.Uint32X4, v0 *arm.Uint32X4, v1 *arm.Uint32X4, v2 *arm.Uint32X4)

// Three-way Exclusive OR performs a three-way exclusive OR of the values in the three source SIMD&FP registers, and writes the result to the destination SIMD&FP register.
//
//go:linkname Veor3QU64 Veor3QU64
//go:noescape
func Veor3QU64(r *arm.Uint64X2, v0 *arm.Uint64X2, v1 *arm.Uint64X2, v2 *arm.Uint64X2)

// Three-way Exclusive OR performs a three-way exclusive OR of the values in the three source SIMD&FP registers, and writes the result to the destination SIMD&FP register.
//
//go:linkname Veor3QU16 Veor3QU16
//go:noescape
func Veor3QU16(r *arm.Uint16X8, v0 *arm.Uint16X8, v1 *arm.Uint16X8, v2 *arm.Uint16X8)

// Three-way Exclusive OR performs a three-way exclusive OR of the values in the three source SIMD&FP registers, and writes the result to the destination SIMD&FP register.
//
//go:linkname Veor3QS8 Veor3QS8
//go:noescape
func Veor3QS8(r *arm.Int8X16, v0 *arm.Int8X16, v1 *arm.Int8X16, v2 *arm.Int8X16)

// Three-way Exclusive OR performs a three-way exclusive OR of the values in the three source SIMD&FP registers, and writes the result to the destination SIMD&FP register.
//
//go:linkname Veor3QS32 Veor3QS32
//go:noescape
func Veor3QS32(r *arm.Int32X4, v0 *arm.Int32X4, v1 *arm.Int32X4, v2 *arm.Int32X4)

// Three-way Exclusive OR performs a three-way exclusive OR of the values in the three source SIMD&FP registers, and writes the result to the destination SIMD&FP register.
//
//go:linkname Veor3QS64 Veor3QS64
//go:noescape
func Veor3QS64(r *arm.Int64X2, v0 *arm.Int64X2, v1 *arm.Int64X2, v2 *arm.Int64X2)

// Three-way Exclusive OR performs a three-way exclusive OR of the values in the three source SIMD&FP registers, and writes the result to the destination SIMD&FP register.
//
//go:linkname Veor3QS16 Veor3QS16
//go:noescape
func Veor3QS16(r *arm.Int16X8, v0 *arm.Int16X8, v1 *arm.Int16X8, v2 *arm.Int16X8)

// Rotate and Exclusive OR rotates each 64-bit element of the 128-bit vector in a source SIMD&FP register left by 1, performs a bitwise exclusive OR of the resulting 128-bit vector and the vector in another source SIMD&FP register, and writes the result to the destination SIMD&FP register.
//
//go:linkname Vrax1QU64 Vrax1QU64
//go:noescape
func Vrax1QU64(r *arm.Uint64X2, v0 *arm.Uint64X2, v1 *arm.Uint64X2)

// SHA512 Hash update part 1 takes the values from the three 128-bit source SIMD&FP registers and produces a 128-bit output value that combines the sigma1 and chi functions of two iterations of the SHA512 computation. It returns this value to the destination SIMD&FP register.
//
//go:linkname Vsha512HqU64 Vsha512HqU64
//go:noescape
func Vsha512HqU64(r *arm.Uint64X2, v0 *arm.Uint64X2, v1 *arm.Uint64X2, v2 *arm.Uint64X2)

// SHA512 Hash update part 2 takes the values from the three 128-bit source SIMD&FP registers and produces a 128-bit output value that combines the sigma0 and majority functions of two iterations of the SHA512 computation. It returns this value to the destination SIMD&FP register.
//
//go:linkname Vsha512H2QU64 Vsha512H2QU64
//go:noescape
func Vsha512H2QU64(r *arm.Uint64X2, v0 *arm.Uint64X2, v1 *arm.Uint64X2, v2 *arm.Uint64X2)

// SHA512 Schedule Update 0 takes the values from the two 128-bit source SIMD&FP registers and produces a 128-bit output value that combines the gamma0 functions of two iterations of the SHA512 schedule update that are performed after the first 16 iterations within a block. It returns this value to the destination SIMD&FP register.
//
//go:linkname Vsha512Su0QU64 Vsha512Su0QU64
//go:noescape
func Vsha512Su0QU64(r *arm.Uint64X2, v0 *arm.Uint64X2, v1 *arm.Uint64X2)

// SHA512 Schedule Update 1 takes the values from the three source SIMD&FP registers and produces a 128-bit output value that combines the gamma1 functions of two iterations of the SHA512 schedule update that are performed after the first 16 iterations within a block. It returns this value to the destination SIMD&FP register.
//
//go:linkname Vsha512Su1QU64 Vsha512Su1QU64
//go:noescape
func Vsha512Su1QU64(r *arm.Uint64X2, v0 *arm.Uint64X2, v1 *arm.Uint64X2, v2 *arm.Uint64X2)

// SM3PARTW1 takes three 128-bit vectors from the three source SIMD&FP registers and returns a 128-bit result in the destination SIMD&FP register. The result is obtained by a three-way exclusive OR of the elements within the input vectors with some fixed rotations, see the Operation pseudocode for more information.
//
//go:linkname Vsm3Partw1QU32 Vsm3Partw1QU32
//go:noescape
func Vsm3Partw1QU32(r *arm.Uint32X4, v0 *arm.Uint32X4, v1 *arm.Uint32X4, v2 *arm.Uint32X4)

// SM3PARTW2 takes three 128-bit vectors from three source SIMD&FP registers and returns a 128-bit result in the destination SIMD&FP register. The result is obtained by a three-way exclusive OR of the elements within the input vectors with some fixed rotations, see the Operation pseudocode for more information.
//
//go:linkname Vsm3Partw2QU32 Vsm3Partw2QU32
//go:noescape
func Vsm3Partw2QU32(r *arm.Uint32X4, v0 *arm.Uint32X4, v1 *arm.Uint32X4, v2 *arm.Uint32X4)

// SM3SS1 rotates the top 32 bits of the 128-bit vector in the first source SIMD&FP register by 12, and adds that 32-bit value to the two other 32-bit values held in the top 32 bits of each of the 128-bit vectors in the second and third source SIMD&FP registers, rotating this result left by 7 and writing the final result into the top 32 bits of the vector in the destination SIMD&FP register, with the bottom 96 bits of the vector being written to 0.
//
//go:linkname Vsm3Ss1QU32 Vsm3Ss1QU32
//go:noescape
func Vsm3Ss1QU32(r *arm.Uint32X4, v0 *arm.Uint32X4, v1 *arm.Uint32X4, v2 *arm.Uint32X4)

// SM4 Encode takes input data as a 128-bit vector from the first source SIMD&FP register, and four iterations of the round key held as the elements of the 128-bit vector in the second source SIMD&FP register. It encrypts the data by four rounds, in accordance with the SM4 standard, returning the 128-bit result to the destination SIMD&FP register.
//
//go:linkname Vsm4EqU32 Vsm4EqU32
//go:noescape
func Vsm4EqU32(r *arm.Uint32X4, v0 *arm.Uint32X4, v1 *arm.Uint32X4)

// SM4 Key takes an input as a 128-bit vector from the first source SIMD&FP register and a 128-bit constant from the second SIMD&FP register. It derives four iterations of the output key, in accordance with the SM4 standard, returning the 128-bit result to the destination SIMD&FP register.
//
//go:linkname Vsm4EkeyqU32 Vsm4EkeyqU32
//go:noescape
func Vsm4EkeyqU32(r *arm.Uint32X4, v0 *arm.Uint32X4, v1 *arm.Uint32X4)

// Floating-point Convert to Signed integer, rounding to nearest with ties to Away (vector). This instruction converts each element in a vector from a floating-point value to a signed integer value using the Round to Nearest with Ties to Away rounding mode and writes the result to the SIMD&FP destination register.
//
//go:linkname VcvtaqS64F64 VcvtaqS64F64
//go:noescape
func VcvtaqS64F64(r *arm.Int64X2, v0 *arm.Float64X2)

// Floating-point Convert to Signed integer, rounding to nearest with ties to Away (vector). This instruction converts each element in a vector from a floating-point value to a signed integer value using the Round to Nearest with Ties to Away rounding mode and writes the result to the SIMD&FP destination register.
//
//go:linkname VcvtaS64F64 VcvtaS64F64
//go:noescape
func VcvtaS64F64(r *arm.Int64X1, v0 *arm.Float64X1)

// Floating-point Convert to Unsigned integer, rounding to nearest with ties to Away (vector). This instruction converts each element in a vector from a floating-point value to an unsigned integer value using the Round to Nearest with Ties to Away rounding mode and writes the result to the SIMD&FP destination register.
//
//go:linkname VcvtaqU64F64 VcvtaqU64F64
//go:noescape
func VcvtaqU64F64(r *arm.Uint64X2, v0 *arm.Float64X2)

// Floating-point Convert to Unsigned integer, rounding to nearest with ties to Away (vector). This instruction converts each element in a vector from a floating-point value to an unsigned integer value using the Round to Nearest with Ties to Away rounding mode and writes the result to the SIMD&FP destination register.
//
//go:linkname VcvtaU64F64 VcvtaU64F64
//go:noescape
func VcvtaU64F64(r *arm.Uint64X1, v0 *arm.Float64X1)

// Floating-point Convert to Signed integer, rounding toward Minus infinity (vector). This instruction converts a scalar or each element in a vector from a floating-point value to a signed integer value using the Round towards Minus Infinity rounding mode, and writes the result to the SIMD&FP destination register.
//
//go:linkname VcvtmqS64F64 VcvtmqS64F64
//go:noescape
func VcvtmqS64F64(r *arm.Int64X2, v0 *arm.Float64X2)

// Floating-point Convert to Signed integer, rounding toward Minus infinity (vector). This instruction converts a scalar or each element in a vector from a floating-point value to a signed integer value using the Round towards Minus Infinity rounding mode, and writes the result to the SIMD&FP destination register.
//
//go:linkname VcvtmS64F64 VcvtmS64F64
//go:noescape
func VcvtmS64F64(r *arm.Int64X1, v0 *arm.Float64X1)

// Floating-point Convert to Unsigned integer, rounding toward Minus infinity (vector). This instruction converts a scalar or each element in a vector from a floating-point value to an unsigned integer value using the Round towards Minus Infinity rounding mode, and writes the result to the SIMD&FP destination register.
//
//go:linkname VcvtmqU64F64 VcvtmqU64F64
//go:noescape
func VcvtmqU64F64(r *arm.Uint64X2, v0 *arm.Float64X2)

// Floating-point Convert to Unsigned integer, rounding toward Minus infinity (vector). This instruction converts a scalar or each element in a vector from a floating-point value to an unsigned integer value using the Round towards Minus Infinity rounding mode, and writes the result to the SIMD&FP destination register.
//
//go:linkname VcvtmU64F64 VcvtmU64F64
//go:noescape
func VcvtmU64F64(r *arm.Uint64X1, v0 *arm.Float64X1)

// Floating-point Convert to Signed integer, rounding to nearest with ties to even (vector). This instruction converts a scalar or each element in a vector from a floating-point value to a signed integer value using the Round to Nearest rounding mode, and writes the result to the SIMD&FP destination register.
//
//go:linkname VcvtnqS64F64 VcvtnqS64F64
//go:noescape
func VcvtnqS64F64(r *arm.Int64X2, v0 *arm.Float64X2)

// Floating-point Convert to Signed integer, rounding to nearest with ties to even (vector). This instruction converts a scalar or each element in a vector from a floating-point value to a signed integer value using the Round to Nearest rounding mode, and writes the result to the SIMD&FP destination register.
//
//go:linkname VcvtnS64F64 VcvtnS64F64
//go:noescape
func VcvtnS64F64(r *arm.Int64X1, v0 *arm.Float64X1)

// Floating-point Convert to Unsigned integer, rounding to nearest with ties to even (vector). This instruction converts a scalar or each element in a vector from a floating-point value to an unsigned integer value using the Round to Nearest rounding mode, and writes the result to the SIMD&FP destination register.
//
//go:linkname VcvtnqU64F64 VcvtnqU64F64
//go:noescape
func VcvtnqU64F64(r *arm.Uint64X2, v0 *arm.Float64X2)

// Floating-point Convert to Unsigned integer, rounding to nearest with ties to even (vector). This instruction converts a scalar or each element in a vector from a floating-point value to an unsigned integer value using the Round to Nearest rounding mode, and writes the result to the SIMD&FP destination register.
//
//go:linkname VcvtnU64F64 VcvtnU64F64
//go:noescape
func VcvtnU64F64(r *arm.Uint64X1, v0 *arm.Float64X1)

// Floating-point Convert to Signed integer, rounding toward Plus infinity (vector). This instruction converts a scalar or each element in a vector from a floating-point value to a signed integer value using the Round towards Plus Infinity rounding mode, and writes the result to the SIMD&FP destination register.
//
//go:linkname VcvtpqS64F64 VcvtpqS64F64
//go:noescape
func VcvtpqS64F64(r *arm.Int64X2, v0 *arm.Float64X2)

// Floating-point Convert to Signed integer, rounding toward Plus infinity (vector). This instruction converts a scalar or each element in a vector from a floating-point value to a signed integer value using the Round towards Plus Infinity rounding mode, and writes the result to the SIMD&FP destination register.
//
//go:linkname VcvtpS64F64 VcvtpS64F64
//go:noescape
func VcvtpS64F64(r *arm.Int64X1, v0 *arm.Float64X1)

// Floating-point Convert to Unsigned integer, rounding toward Plus infinity (vector). This instruction converts a scalar or each element in a vector from a floating-point value to an unsigned integer value using the Round towards Plus Infinity rounding mode, and writes the result to the SIMD&FP destination register.
//
//go:linkname VcvtpqU64F64 VcvtpqU64F64
//go:noescape
func VcvtpqU64F64(r *arm.Uint64X2, v0 *arm.Float64X2)

// Floating-point Convert to Unsigned integer, rounding toward Plus infinity (vector). This instruction converts a scalar or each element in a vector from a floating-point value to an unsigned integer value using the Round towards Plus Infinity rounding mode, and writes the result to the SIMD&FP destination register.
//
//go:linkname VcvtpU64F64 VcvtpU64F64
//go:noescape
func VcvtpU64F64(r *arm.Uint64X1, v0 *arm.Float64X1)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretP8P64 VreinterpretP8P64
//go:noescape
func VreinterpretP8P64(r *arm.Poly8X8, v0 *arm.Poly64X1)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretP8P16 VreinterpretP8P16
//go:noescape
func VreinterpretP8P16(r *arm.Poly8X8, v0 *arm.Poly16X4)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretP8U8 VreinterpretP8U8
//go:noescape
func VreinterpretP8U8(r *arm.Poly8X8, v0 *arm.Uint8X8)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretP8U32 VreinterpretP8U32
//go:noescape
func VreinterpretP8U32(r *arm.Poly8X8, v0 *arm.Uint32X2)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretP8U64 VreinterpretP8U64
//go:noescape
func VreinterpretP8U64(r *arm.Poly8X8, v0 *arm.Uint64X1)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretP8U16 VreinterpretP8U16
//go:noescape
func VreinterpretP8U16(r *arm.Poly8X8, v0 *arm.Uint16X4)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretP8S8 VreinterpretP8S8
//go:noescape
func VreinterpretP8S8(r *arm.Poly8X8, v0 *arm.Int8X8)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretP8F64 VreinterpretP8F64
//go:noescape
func VreinterpretP8F64(r *arm.Poly8X8, v0 *arm.Float64X1)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretP8F32 VreinterpretP8F32
//go:noescape
func VreinterpretP8F32(r *arm.Poly8X8, v0 *arm.Float32X2)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretP8S32 VreinterpretP8S32
//go:noescape
func VreinterpretP8S32(r *arm.Poly8X8, v0 *arm.Int32X2)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretP8S64 VreinterpretP8S64
//go:noescape
func VreinterpretP8S64(r *arm.Poly8X8, v0 *arm.Int64X1)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretP8S16 VreinterpretP8S16
//go:noescape
func VreinterpretP8S16(r *arm.Poly8X8, v0 *arm.Int16X4)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretP64P8 VreinterpretP64P8
//go:noescape
func VreinterpretP64P8(r *arm.Poly64X1, v0 *arm.Poly8X8)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretP64P16 VreinterpretP64P16
//go:noescape
func VreinterpretP64P16(r *arm.Poly64X1, v0 *arm.Poly16X4)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretP64U8 VreinterpretP64U8
//go:noescape
func VreinterpretP64U8(r *arm.Poly64X1, v0 *arm.Uint8X8)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretP64U32 VreinterpretP64U32
//go:noescape
func VreinterpretP64U32(r *arm.Poly64X1, v0 *arm.Uint32X2)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretP64U64 VreinterpretP64U64
//go:noescape
func VreinterpretP64U64(r *arm.Poly64X1, v0 *arm.Uint64X1)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretP64U16 VreinterpretP64U16
//go:noescape
func VreinterpretP64U16(r *arm.Poly64X1, v0 *arm.Uint16X4)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretP64S8 VreinterpretP64S8
//go:noescape
func VreinterpretP64S8(r *arm.Poly64X1, v0 *arm.Int8X8)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretP64F64 VreinterpretP64F64
//go:noescape
func VreinterpretP64F64(r *arm.Poly64X1, v0 *arm.Float64X1)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretP64F32 VreinterpretP64F32
//go:noescape
func VreinterpretP64F32(r *arm.Poly64X1, v0 *arm.Float32X2)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretP64S32 VreinterpretP64S32
//go:noescape
func VreinterpretP64S32(r *arm.Poly64X1, v0 *arm.Int32X2)

// vreinterpret_p64_s64
//
//go:linkname VreinterpretP64S64 VreinterpretP64S64
//go:noescape
func VreinterpretP64S64(r *arm.Poly64X1, v0 *arm.Int64X1)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretP64S16 VreinterpretP64S16
//go:noescape
func VreinterpretP64S16(r *arm.Poly64X1, v0 *arm.Int16X4)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretP16P8 VreinterpretP16P8
//go:noescape
func VreinterpretP16P8(r *arm.Poly16X4, v0 *arm.Poly8X8)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretP16P64 VreinterpretP16P64
//go:noescape
func VreinterpretP16P64(r *arm.Poly16X4, v0 *arm.Poly64X1)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretP16U8 VreinterpretP16U8
//go:noescape
func VreinterpretP16U8(r *arm.Poly16X4, v0 *arm.Uint8X8)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretP16U32 VreinterpretP16U32
//go:noescape
func VreinterpretP16U32(r *arm.Poly16X4, v0 *arm.Uint32X2)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretP16U64 VreinterpretP16U64
//go:noescape
func VreinterpretP16U64(r *arm.Poly16X4, v0 *arm.Uint64X1)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretP16U16 VreinterpretP16U16
//go:noescape
func VreinterpretP16U16(r *arm.Poly16X4, v0 *arm.Uint16X4)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretP16S8 VreinterpretP16S8
//go:noescape
func VreinterpretP16S8(r *arm.Poly16X4, v0 *arm.Int8X8)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretP16F64 VreinterpretP16F64
//go:noescape
func VreinterpretP16F64(r *arm.Poly16X4, v0 *arm.Float64X1)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretP16F32 VreinterpretP16F32
//go:noescape
func VreinterpretP16F32(r *arm.Poly16X4, v0 *arm.Float32X2)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretP16S32 VreinterpretP16S32
//go:noescape
func VreinterpretP16S32(r *arm.Poly16X4, v0 *arm.Int32X2)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretP16S64 VreinterpretP16S64
//go:noescape
func VreinterpretP16S64(r *arm.Poly16X4, v0 *arm.Int64X1)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretP16S16 VreinterpretP16S16
//go:noescape
func VreinterpretP16S16(r *arm.Poly16X4, v0 *arm.Int16X4)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretqP8P128 VreinterpretqP8P128
//go:noescape
func VreinterpretqP8P128(r *arm.Poly8X16, v0 *arm.Poly128)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretqP8P64 VreinterpretqP8P64
//go:noescape
func VreinterpretqP8P64(r *arm.Poly8X16, v0 *arm.Poly64X2)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretqP8P16 VreinterpretqP8P16
//go:noescape
func VreinterpretqP8P16(r *arm.Poly8X16, v0 *arm.Poly16X8)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretqP8U8 VreinterpretqP8U8
//go:noescape
func VreinterpretqP8U8(r *arm.Poly8X16, v0 *arm.Uint8X16)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretqP8U32 VreinterpretqP8U32
//go:noescape
func VreinterpretqP8U32(r *arm.Poly8X16, v0 *arm.Uint32X4)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretqP8U64 VreinterpretqP8U64
//go:noescape
func VreinterpretqP8U64(r *arm.Poly8X16, v0 *arm.Uint64X2)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretqP8U16 VreinterpretqP8U16
//go:noescape
func VreinterpretqP8U16(r *arm.Poly8X16, v0 *arm.Uint16X8)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretqP8S8 VreinterpretqP8S8
//go:noescape
func VreinterpretqP8S8(r *arm.Poly8X16, v0 *arm.Int8X16)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretqP8F64 VreinterpretqP8F64
//go:noescape
func VreinterpretqP8F64(r *arm.Poly8X16, v0 *arm.Float64X2)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretqP8F32 VreinterpretqP8F32
//go:noescape
func VreinterpretqP8F32(r *arm.Poly8X16, v0 *arm.Float32X4)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretqP8S32 VreinterpretqP8S32
//go:noescape
func VreinterpretqP8S32(r *arm.Poly8X16, v0 *arm.Int32X4)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretqP8S64 VreinterpretqP8S64
//go:noescape
func VreinterpretqP8S64(r *arm.Poly8X16, v0 *arm.Int64X2)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretqP8S16 VreinterpretqP8S16
//go:noescape
func VreinterpretqP8S16(r *arm.Poly8X16, v0 *arm.Int16X8)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretqP128P8 VreinterpretqP128P8
//go:noescape
func VreinterpretqP128P8(r *arm.Poly128, v0 *arm.Poly8X16)

// vreinterpretq_p128_p64
//
//go:linkname VreinterpretqP128P64 VreinterpretqP128P64
//go:noescape
func VreinterpretqP128P64(r *arm.Poly128, v0 *arm.Poly64X2)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretqP128P16 VreinterpretqP128P16
//go:noescape
func VreinterpretqP128P16(r *arm.Poly128, v0 *arm.Poly16X8)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretqP128U8 VreinterpretqP128U8
//go:noescape
func VreinterpretqP128U8(r *arm.Poly128, v0 *arm.Uint8X16)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretqP128U32 VreinterpretqP128U32
//go:noescape
func VreinterpretqP128U32(r *arm.Poly128, v0 *arm.Uint32X4)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretqP128U64 VreinterpretqP128U64
//go:noescape
func VreinterpretqP128U64(r *arm.Poly128, v0 *arm.Uint64X2)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretqP128U16 VreinterpretqP128U16
//go:noescape
func VreinterpretqP128U16(r *arm.Poly128, v0 *arm.Uint16X8)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretqP128S8 VreinterpretqP128S8
//go:noescape
func VreinterpretqP128S8(r *arm.Poly128, v0 *arm.Int8X16)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretqP128F64 VreinterpretqP128F64
//go:noescape
func VreinterpretqP128F64(r *arm.Poly128, v0 *arm.Float64X2)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretqP128F32 VreinterpretqP128F32
//go:noescape
func VreinterpretqP128F32(r *arm.Poly128, v0 *arm.Float32X4)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretqP128S32 VreinterpretqP128S32
//go:noescape
func VreinterpretqP128S32(r *arm.Poly128, v0 *arm.Int32X4)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretqP128S64 VreinterpretqP128S64
//go:noescape
func VreinterpretqP128S64(r *arm.Poly128, v0 *arm.Int64X2)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretqP128S16 VreinterpretqP128S16
//go:noescape
func VreinterpretqP128S16(r *arm.Poly128, v0 *arm.Int16X8)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretqP64P8 VreinterpretqP64P8
//go:noescape
func VreinterpretqP64P8(r *arm.Poly64X2, v0 *arm.Poly8X16)

// vreinterpretq_p64_p128
//
//go:linkname VreinterpretqP64P128 VreinterpretqP64P128
//go:noescape
func VreinterpretqP64P128(r *arm.Poly64X2, v0 *arm.Poly128)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretqP64P16 VreinterpretqP64P16
//go:noescape
func VreinterpretqP64P16(r *arm.Poly64X2, v0 *arm.Poly16X8)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretqP64U8 VreinterpretqP64U8
//go:noescape
func VreinterpretqP64U8(r *arm.Poly64X2, v0 *arm.Uint8X16)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretqP64U32 VreinterpretqP64U32
//go:noescape
func VreinterpretqP64U32(r *arm.Poly64X2, v0 *arm.Uint32X4)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretqP64U64 VreinterpretqP64U64
//go:noescape
func VreinterpretqP64U64(r *arm.Poly64X2, v0 *arm.Uint64X2)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretqP64U16 VreinterpretqP64U16
//go:noescape
func VreinterpretqP64U16(r *arm.Poly64X2, v0 *arm.Uint16X8)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretqP64S8 VreinterpretqP64S8
//go:noescape
func VreinterpretqP64S8(r *arm.Poly64X2, v0 *arm.Int8X16)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretqP64F64 VreinterpretqP64F64
//go:noescape
func VreinterpretqP64F64(r *arm.Poly64X2, v0 *arm.Float64X2)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretqP64F32 VreinterpretqP64F32
//go:noescape
func VreinterpretqP64F32(r *arm.Poly64X2, v0 *arm.Float32X4)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretqP64S32 VreinterpretqP64S32
//go:noescape
func VreinterpretqP64S32(r *arm.Poly64X2, v0 *arm.Int32X4)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretqP64S64 VreinterpretqP64S64
//go:noescape
func VreinterpretqP64S64(r *arm.Poly64X2, v0 *arm.Int64X2)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretqP64S16 VreinterpretqP64S16
//go:noescape
func VreinterpretqP64S16(r *arm.Poly64X2, v0 *arm.Int16X8)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretqP16P8 VreinterpretqP16P8
//go:noescape
func VreinterpretqP16P8(r *arm.Poly16X8, v0 *arm.Poly8X16)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretqP16P128 VreinterpretqP16P128
//go:noescape
func VreinterpretqP16P128(r *arm.Poly16X8, v0 *arm.Poly128)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretqP16P64 VreinterpretqP16P64
//go:noescape
func VreinterpretqP16P64(r *arm.Poly16X8, v0 *arm.Poly64X2)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretqP16U8 VreinterpretqP16U8
//go:noescape
func VreinterpretqP16U8(r *arm.Poly16X8, v0 *arm.Uint8X16)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretqP16U32 VreinterpretqP16U32
//go:noescape
func VreinterpretqP16U32(r *arm.Poly16X8, v0 *arm.Uint32X4)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretqP16U64 VreinterpretqP16U64
//go:noescape
func VreinterpretqP16U64(r *arm.Poly16X8, v0 *arm.Uint64X2)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretqP16U16 VreinterpretqP16U16
//go:noescape
func VreinterpretqP16U16(r *arm.Poly16X8, v0 *arm.Uint16X8)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretqP16S8 VreinterpretqP16S8
//go:noescape
func VreinterpretqP16S8(r *arm.Poly16X8, v0 *arm.Int8X16)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretqP16F64 VreinterpretqP16F64
//go:noescape
func VreinterpretqP16F64(r *arm.Poly16X8, v0 *arm.Float64X2)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretqP16F32 VreinterpretqP16F32
//go:noescape
func VreinterpretqP16F32(r *arm.Poly16X8, v0 *arm.Float32X4)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretqP16S32 VreinterpretqP16S32
//go:noescape
func VreinterpretqP16S32(r *arm.Poly16X8, v0 *arm.Int32X4)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretqP16S64 VreinterpretqP16S64
//go:noescape
func VreinterpretqP16S64(r *arm.Poly16X8, v0 *arm.Int64X2)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretqP16S16 VreinterpretqP16S16
//go:noescape
func VreinterpretqP16S16(r *arm.Poly16X8, v0 *arm.Int16X8)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretqU8P8 VreinterpretqU8P8
//go:noescape
func VreinterpretqU8P8(r *arm.Uint8X16, v0 *arm.Poly8X16)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretqU8P128 VreinterpretqU8P128
//go:noescape
func VreinterpretqU8P128(r *arm.Uint8X16, v0 *arm.Poly128)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretqU8P64 VreinterpretqU8P64
//go:noescape
func VreinterpretqU8P64(r *arm.Uint8X16, v0 *arm.Poly64X2)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretqU8P16 VreinterpretqU8P16
//go:noescape
func VreinterpretqU8P16(r *arm.Uint8X16, v0 *arm.Poly16X8)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretqU8U32 VreinterpretqU8U32
//go:noescape
func VreinterpretqU8U32(r *arm.Uint8X16, v0 *arm.Uint32X4)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretqU8U64 VreinterpretqU8U64
//go:noescape
func VreinterpretqU8U64(r *arm.Uint8X16, v0 *arm.Uint64X2)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretqU8U16 VreinterpretqU8U16
//go:noescape
func VreinterpretqU8U16(r *arm.Uint8X16, v0 *arm.Uint16X8)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretqU8S8 VreinterpretqU8S8
//go:noescape
func VreinterpretqU8S8(r *arm.Uint8X16, v0 *arm.Int8X16)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretqU8F64 VreinterpretqU8F64
//go:noescape
func VreinterpretqU8F64(r *arm.Uint8X16, v0 *arm.Float64X2)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretqU8F32 VreinterpretqU8F32
//go:noescape
func VreinterpretqU8F32(r *arm.Uint8X16, v0 *arm.Float32X4)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretqU8S32 VreinterpretqU8S32
//go:noescape
func VreinterpretqU8S32(r *arm.Uint8X16, v0 *arm.Int32X4)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretqU8S64 VreinterpretqU8S64
//go:noescape
func VreinterpretqU8S64(r *arm.Uint8X16, v0 *arm.Int64X2)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretqU8S16 VreinterpretqU8S16
//go:noescape
func VreinterpretqU8S16(r *arm.Uint8X16, v0 *arm.Int16X8)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretqU32P8 VreinterpretqU32P8
//go:noescape
func VreinterpretqU32P8(r *arm.Uint32X4, v0 *arm.Poly8X16)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretqU32P128 VreinterpretqU32P128
//go:noescape
func VreinterpretqU32P128(r *arm.Uint32X4, v0 *arm.Poly128)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretqU32P64 VreinterpretqU32P64
//go:noescape
func VreinterpretqU32P64(r *arm.Uint32X4, v0 *arm.Poly64X2)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretqU32P16 VreinterpretqU32P16
//go:noescape
func VreinterpretqU32P16(r *arm.Uint32X4, v0 *arm.Poly16X8)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretqU32U8 VreinterpretqU32U8
//go:noescape
func VreinterpretqU32U8(r *arm.Uint32X4, v0 *arm.Uint8X16)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretqU32U64 VreinterpretqU32U64
//go:noescape
func VreinterpretqU32U64(r *arm.Uint32X4, v0 *arm.Uint64X2)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretqU32U16 VreinterpretqU32U16
//go:noescape
func VreinterpretqU32U16(r *arm.Uint32X4, v0 *arm.Uint16X8)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretqU32S8 VreinterpretqU32S8
//go:noescape
func VreinterpretqU32S8(r *arm.Uint32X4, v0 *arm.Int8X16)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretqU32F64 VreinterpretqU32F64
//go:noescape
func VreinterpretqU32F64(r *arm.Uint32X4, v0 *arm.Float64X2)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretqU32F32 VreinterpretqU32F32
//go:noescape
func VreinterpretqU32F32(r *arm.Uint32X4, v0 *arm.Float32X4)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretqU32S32 VreinterpretqU32S32
//go:noescape
func VreinterpretqU32S32(r *arm.Uint32X4, v0 *arm.Int32X4)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretqU32S64 VreinterpretqU32S64
//go:noescape
func VreinterpretqU32S64(r *arm.Uint32X4, v0 *arm.Int64X2)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretqU32S16 VreinterpretqU32S16
//go:noescape
func VreinterpretqU32S16(r *arm.Uint32X4, v0 *arm.Int16X8)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretqU64P8 VreinterpretqU64P8
//go:noescape
func VreinterpretqU64P8(r *arm.Uint64X2, v0 *arm.Poly8X16)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretqU64P128 VreinterpretqU64P128
//go:noescape
func VreinterpretqU64P128(r *arm.Uint64X2, v0 *arm.Poly128)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretqU64P64 VreinterpretqU64P64
//go:noescape
func VreinterpretqU64P64(r *arm.Uint64X2, v0 *arm.Poly64X2)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretqU64P16 VreinterpretqU64P16
//go:noescape
func VreinterpretqU64P16(r *arm.Uint64X2, v0 *arm.Poly16X8)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretqU64U8 VreinterpretqU64U8
//go:noescape
func VreinterpretqU64U8(r *arm.Uint64X2, v0 *arm.Uint8X16)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretqU64U32 VreinterpretqU64U32
//go:noescape
func VreinterpretqU64U32(r *arm.Uint64X2, v0 *arm.Uint32X4)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretqU64U16 VreinterpretqU64U16
//go:noescape
func VreinterpretqU64U16(r *arm.Uint64X2, v0 *arm.Uint16X8)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretqU64S8 VreinterpretqU64S8
//go:noescape
func VreinterpretqU64S8(r *arm.Uint64X2, v0 *arm.Int8X16)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretqU64F64 VreinterpretqU64F64
//go:noescape
func VreinterpretqU64F64(r *arm.Uint64X2, v0 *arm.Float64X2)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretqU64F32 VreinterpretqU64F32
//go:noescape
func VreinterpretqU64F32(r *arm.Uint64X2, v0 *arm.Float32X4)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretqU64S32 VreinterpretqU64S32
//go:noescape
func VreinterpretqU64S32(r *arm.Uint64X2, v0 *arm.Int32X4)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretqU64S64 VreinterpretqU64S64
//go:noescape
func VreinterpretqU64S64(r *arm.Uint64X2, v0 *arm.Int64X2)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretqU64S16 VreinterpretqU64S16
//go:noescape
func VreinterpretqU64S16(r *arm.Uint64X2, v0 *arm.Int16X8)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretqU16P8 VreinterpretqU16P8
//go:noescape
func VreinterpretqU16P8(r *arm.Uint16X8, v0 *arm.Poly8X16)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretqU16P128 VreinterpretqU16P128
//go:noescape
func VreinterpretqU16P128(r *arm.Uint16X8, v0 *arm.Poly128)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretqU16P64 VreinterpretqU16P64
//go:noescape
func VreinterpretqU16P64(r *arm.Uint16X8, v0 *arm.Poly64X2)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretqU16P16 VreinterpretqU16P16
//go:noescape
func VreinterpretqU16P16(r *arm.Uint16X8, v0 *arm.Poly16X8)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretqU16U8 VreinterpretqU16U8
//go:noescape
func VreinterpretqU16U8(r *arm.Uint16X8, v0 *arm.Uint8X16)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretqU16U32 VreinterpretqU16U32
//go:noescape
func VreinterpretqU16U32(r *arm.Uint16X8, v0 *arm.Uint32X4)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretqU16U64 VreinterpretqU16U64
//go:noescape
func VreinterpretqU16U64(r *arm.Uint16X8, v0 *arm.Uint64X2)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretqU16S8 VreinterpretqU16S8
//go:noescape
func VreinterpretqU16S8(r *arm.Uint16X8, v0 *arm.Int8X16)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretqU16F64 VreinterpretqU16F64
//go:noescape
func VreinterpretqU16F64(r *arm.Uint16X8, v0 *arm.Float64X2)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretqU16F32 VreinterpretqU16F32
//go:noescape
func VreinterpretqU16F32(r *arm.Uint16X8, v0 *arm.Float32X4)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretqU16S32 VreinterpretqU16S32
//go:noescape
func VreinterpretqU16S32(r *arm.Uint16X8, v0 *arm.Int32X4)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretqU16S64 VreinterpretqU16S64
//go:noescape
func VreinterpretqU16S64(r *arm.Uint16X8, v0 *arm.Int64X2)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretqU16S16 VreinterpretqU16S16
//go:noescape
func VreinterpretqU16S16(r *arm.Uint16X8, v0 *arm.Int16X8)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretqS8P8 VreinterpretqS8P8
//go:noescape
func VreinterpretqS8P8(r *arm.Int8X16, v0 *arm.Poly8X16)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretqS8P128 VreinterpretqS8P128
//go:noescape
func VreinterpretqS8P128(r *arm.Int8X16, v0 *arm.Poly128)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretqS8P64 VreinterpretqS8P64
//go:noescape
func VreinterpretqS8P64(r *arm.Int8X16, v0 *arm.Poly64X2)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretqS8P16 VreinterpretqS8P16
//go:noescape
func VreinterpretqS8P16(r *arm.Int8X16, v0 *arm.Poly16X8)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretqS8U8 VreinterpretqS8U8
//go:noescape
func VreinterpretqS8U8(r *arm.Int8X16, v0 *arm.Uint8X16)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretqS8U32 VreinterpretqS8U32
//go:noescape
func VreinterpretqS8U32(r *arm.Int8X16, v0 *arm.Uint32X4)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretqS8U64 VreinterpretqS8U64
//go:noescape
func VreinterpretqS8U64(r *arm.Int8X16, v0 *arm.Uint64X2)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretqS8U16 VreinterpretqS8U16
//go:noescape
func VreinterpretqS8U16(r *arm.Int8X16, v0 *arm.Uint16X8)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretqS8F64 VreinterpretqS8F64
//go:noescape
func VreinterpretqS8F64(r *arm.Int8X16, v0 *arm.Float64X2)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretqS8F32 VreinterpretqS8F32
//go:noescape
func VreinterpretqS8F32(r *arm.Int8X16, v0 *arm.Float32X4)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretqS8S32 VreinterpretqS8S32
//go:noescape
func VreinterpretqS8S32(r *arm.Int8X16, v0 *arm.Int32X4)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretqS8S64 VreinterpretqS8S64
//go:noescape
func VreinterpretqS8S64(r *arm.Int8X16, v0 *arm.Int64X2)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretqS8S16 VreinterpretqS8S16
//go:noescape
func VreinterpretqS8S16(r *arm.Int8X16, v0 *arm.Int16X8)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretqF64P8 VreinterpretqF64P8
//go:noescape
func VreinterpretqF64P8(r *arm.Float64X2, v0 *arm.Poly8X16)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretqF64P128 VreinterpretqF64P128
//go:noescape
func VreinterpretqF64P128(r *arm.Float64X2, v0 *arm.Poly128)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretqF64P64 VreinterpretqF64P64
//go:noescape
func VreinterpretqF64P64(r *arm.Float64X2, v0 *arm.Poly64X2)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretqF64P16 VreinterpretqF64P16
//go:noescape
func VreinterpretqF64P16(r *arm.Float64X2, v0 *arm.Poly16X8)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretqF64U8 VreinterpretqF64U8
//go:noescape
func VreinterpretqF64U8(r *arm.Float64X2, v0 *arm.Uint8X16)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretqF64U32 VreinterpretqF64U32
//go:noescape
func VreinterpretqF64U32(r *arm.Float64X2, v0 *arm.Uint32X4)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretqF64U64 VreinterpretqF64U64
//go:noescape
func VreinterpretqF64U64(r *arm.Float64X2, v0 *arm.Uint64X2)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretqF64U16 VreinterpretqF64U16
//go:noescape
func VreinterpretqF64U16(r *arm.Float64X2, v0 *arm.Uint16X8)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretqF64S8 VreinterpretqF64S8
//go:noescape
func VreinterpretqF64S8(r *arm.Float64X2, v0 *arm.Int8X16)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretqF64F32 VreinterpretqF64F32
//go:noescape
func VreinterpretqF64F32(r *arm.Float64X2, v0 *arm.Float32X4)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretqF64S32 VreinterpretqF64S32
//go:noescape
func VreinterpretqF64S32(r *arm.Float64X2, v0 *arm.Int32X4)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretqF64S64 VreinterpretqF64S64
//go:noescape
func VreinterpretqF64S64(r *arm.Float64X2, v0 *arm.Int64X2)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretqF64S16 VreinterpretqF64S16
//go:noescape
func VreinterpretqF64S16(r *arm.Float64X2, v0 *arm.Int16X8)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretqF32P8 VreinterpretqF32P8
//go:noescape
func VreinterpretqF32P8(r *arm.Float32X4, v0 *arm.Poly8X16)

// vreinterpretq_f32_p128
//
//go:linkname VreinterpretqF32P128 VreinterpretqF32P128
//go:noescape
func VreinterpretqF32P128(r *arm.Float32X4, v0 *arm.Poly128)

// vreinterpretq_f32_p64
//
//go:linkname VreinterpretqF32P64 VreinterpretqF32P64
//go:noescape
func VreinterpretqF32P64(r *arm.Float32X4, v0 *arm.Poly64X2)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretqF32P16 VreinterpretqF32P16
//go:noescape
func VreinterpretqF32P16(r *arm.Float32X4, v0 *arm.Poly16X8)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretqF32U8 VreinterpretqF32U8
//go:noescape
func VreinterpretqF32U8(r *arm.Float32X4, v0 *arm.Uint8X16)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretqF32U32 VreinterpretqF32U32
//go:noescape
func VreinterpretqF32U32(r *arm.Float32X4, v0 *arm.Uint32X4)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretqF32U64 VreinterpretqF32U64
//go:noescape
func VreinterpretqF32U64(r *arm.Float32X4, v0 *arm.Uint64X2)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretqF32U16 VreinterpretqF32U16
//go:noescape
func VreinterpretqF32U16(r *arm.Float32X4, v0 *arm.Uint16X8)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretqF32S8 VreinterpretqF32S8
//go:noescape
func VreinterpretqF32S8(r *arm.Float32X4, v0 *arm.Int8X16)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretqF32F64 VreinterpretqF32F64
//go:noescape
func VreinterpretqF32F64(r *arm.Float32X4, v0 *arm.Float64X2)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretqF32S32 VreinterpretqF32S32
//go:noescape
func VreinterpretqF32S32(r *arm.Float32X4, v0 *arm.Int32X4)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretqF32S64 VreinterpretqF32S64
//go:noescape
func VreinterpretqF32S64(r *arm.Float32X4, v0 *arm.Int64X2)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretqF32S16 VreinterpretqF32S16
//go:noescape
func VreinterpretqF32S16(r *arm.Float32X4, v0 *arm.Int16X8)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretqS32P8 VreinterpretqS32P8
//go:noescape
func VreinterpretqS32P8(r *arm.Int32X4, v0 *arm.Poly8X16)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretqS32P128 VreinterpretqS32P128
//go:noescape
func VreinterpretqS32P128(r *arm.Int32X4, v0 *arm.Poly128)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretqS32P64 VreinterpretqS32P64
//go:noescape
func VreinterpretqS32P64(r *arm.Int32X4, v0 *arm.Poly64X2)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretqS32P16 VreinterpretqS32P16
//go:noescape
func VreinterpretqS32P16(r *arm.Int32X4, v0 *arm.Poly16X8)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretqS32U8 VreinterpretqS32U8
//go:noescape
func VreinterpretqS32U8(r *arm.Int32X4, v0 *arm.Uint8X16)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretqS32U32 VreinterpretqS32U32
//go:noescape
func VreinterpretqS32U32(r *arm.Int32X4, v0 *arm.Uint32X4)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretqS32U64 VreinterpretqS32U64
//go:noescape
func VreinterpretqS32U64(r *arm.Int32X4, v0 *arm.Uint64X2)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretqS32U16 VreinterpretqS32U16
//go:noescape
func VreinterpretqS32U16(r *arm.Int32X4, v0 *arm.Uint16X8)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretqS32S8 VreinterpretqS32S8
//go:noescape
func VreinterpretqS32S8(r *arm.Int32X4, v0 *arm.Int8X16)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretqS32F64 VreinterpretqS32F64
//go:noescape
func VreinterpretqS32F64(r *arm.Int32X4, v0 *arm.Float64X2)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretqS32F32 VreinterpretqS32F32
//go:noescape
func VreinterpretqS32F32(r *arm.Int32X4, v0 *arm.Float32X4)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretqS32S64 VreinterpretqS32S64
//go:noescape
func VreinterpretqS32S64(r *arm.Int32X4, v0 *arm.Int64X2)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretqS32S16 VreinterpretqS32S16
//go:noescape
func VreinterpretqS32S16(r *arm.Int32X4, v0 *arm.Int16X8)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretqS64P8 VreinterpretqS64P8
//go:noescape
func VreinterpretqS64P8(r *arm.Int64X2, v0 *arm.Poly8X16)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretqS64P128 VreinterpretqS64P128
//go:noescape
func VreinterpretqS64P128(r *arm.Int64X2, v0 *arm.Poly128)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretqS64P64 VreinterpretqS64P64
//go:noescape
func VreinterpretqS64P64(r *arm.Int64X2, v0 *arm.Poly64X2)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretqS64P16 VreinterpretqS64P16
//go:noescape
func VreinterpretqS64P16(r *arm.Int64X2, v0 *arm.Poly16X8)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretqS64U8 VreinterpretqS64U8
//go:noescape
func VreinterpretqS64U8(r *arm.Int64X2, v0 *arm.Uint8X16)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretqS64U32 VreinterpretqS64U32
//go:noescape
func VreinterpretqS64U32(r *arm.Int64X2, v0 *arm.Uint32X4)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretqS64U64 VreinterpretqS64U64
//go:noescape
func VreinterpretqS64U64(r *arm.Int64X2, v0 *arm.Uint64X2)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretqS64U16 VreinterpretqS64U16
//go:noescape
func VreinterpretqS64U16(r *arm.Int64X2, v0 *arm.Uint16X8)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretqS64S8 VreinterpretqS64S8
//go:noescape
func VreinterpretqS64S8(r *arm.Int64X2, v0 *arm.Int8X16)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretqS64F64 VreinterpretqS64F64
//go:noescape
func VreinterpretqS64F64(r *arm.Int64X2, v0 *arm.Float64X2)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretqS64F32 VreinterpretqS64F32
//go:noescape
func VreinterpretqS64F32(r *arm.Int64X2, v0 *arm.Float32X4)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretqS64S32 VreinterpretqS64S32
//go:noescape
func VreinterpretqS64S32(r *arm.Int64X2, v0 *arm.Int32X4)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretqS64S16 VreinterpretqS64S16
//go:noescape
func VreinterpretqS64S16(r *arm.Int64X2, v0 *arm.Int16X8)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretqS16P8 VreinterpretqS16P8
//go:noescape
func VreinterpretqS16P8(r *arm.Int16X8, v0 *arm.Poly8X16)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretqS16P128 VreinterpretqS16P128
//go:noescape
func VreinterpretqS16P128(r *arm.Int16X8, v0 *arm.Poly128)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretqS16P64 VreinterpretqS16P64
//go:noescape
func VreinterpretqS16P64(r *arm.Int16X8, v0 *arm.Poly64X2)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretqS16P16 VreinterpretqS16P16
//go:noescape
func VreinterpretqS16P16(r *arm.Int16X8, v0 *arm.Poly16X8)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretqS16U8 VreinterpretqS16U8
//go:noescape
func VreinterpretqS16U8(r *arm.Int16X8, v0 *arm.Uint8X16)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretqS16U32 VreinterpretqS16U32
//go:noescape
func VreinterpretqS16U32(r *arm.Int16X8, v0 *arm.Uint32X4)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretqS16U64 VreinterpretqS16U64
//go:noescape
func VreinterpretqS16U64(r *arm.Int16X8, v0 *arm.Uint64X2)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretqS16U16 VreinterpretqS16U16
//go:noescape
func VreinterpretqS16U16(r *arm.Int16X8, v0 *arm.Uint16X8)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretqS16S8 VreinterpretqS16S8
//go:noescape
func VreinterpretqS16S8(r *arm.Int16X8, v0 *arm.Int8X16)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretqS16F64 VreinterpretqS16F64
//go:noescape
func VreinterpretqS16F64(r *arm.Int16X8, v0 *arm.Float64X2)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretqS16F32 VreinterpretqS16F32
//go:noescape
func VreinterpretqS16F32(r *arm.Int16X8, v0 *arm.Float32X4)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretqS16S32 VreinterpretqS16S32
//go:noescape
func VreinterpretqS16S32(r *arm.Int16X8, v0 *arm.Int32X4)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretqS16S64 VreinterpretqS16S64
//go:noescape
func VreinterpretqS16S64(r *arm.Int16X8, v0 *arm.Int64X2)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretU8P8 VreinterpretU8P8
//go:noescape
func VreinterpretU8P8(r *arm.Uint8X8, v0 *arm.Poly8X8)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretU8P64 VreinterpretU8P64
//go:noescape
func VreinterpretU8P64(r *arm.Uint8X8, v0 *arm.Poly64X1)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretU8P16 VreinterpretU8P16
//go:noescape
func VreinterpretU8P16(r *arm.Uint8X8, v0 *arm.Poly16X4)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretU8U32 VreinterpretU8U32
//go:noescape
func VreinterpretU8U32(r *arm.Uint8X8, v0 *arm.Uint32X2)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretU8U64 VreinterpretU8U64
//go:noescape
func VreinterpretU8U64(r *arm.Uint8X8, v0 *arm.Uint64X1)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretU8U16 VreinterpretU8U16
//go:noescape
func VreinterpretU8U16(r *arm.Uint8X8, v0 *arm.Uint16X4)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretU8S8 VreinterpretU8S8
//go:noescape
func VreinterpretU8S8(r *arm.Uint8X8, v0 *arm.Int8X8)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretU8F64 VreinterpretU8F64
//go:noescape
func VreinterpretU8F64(r *arm.Uint8X8, v0 *arm.Float64X1)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretU8F32 VreinterpretU8F32
//go:noescape
func VreinterpretU8F32(r *arm.Uint8X8, v0 *arm.Float32X2)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretU8S32 VreinterpretU8S32
//go:noescape
func VreinterpretU8S32(r *arm.Uint8X8, v0 *arm.Int32X2)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretU8S64 VreinterpretU8S64
//go:noescape
func VreinterpretU8S64(r *arm.Uint8X8, v0 *arm.Int64X1)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretU8S16 VreinterpretU8S16
//go:noescape
func VreinterpretU8S16(r *arm.Uint8X8, v0 *arm.Int16X4)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretU32P8 VreinterpretU32P8
//go:noescape
func VreinterpretU32P8(r *arm.Uint32X2, v0 *arm.Poly8X8)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretU32P64 VreinterpretU32P64
//go:noescape
func VreinterpretU32P64(r *arm.Uint32X2, v0 *arm.Poly64X1)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretU32P16 VreinterpretU32P16
//go:noescape
func VreinterpretU32P16(r *arm.Uint32X2, v0 *arm.Poly16X4)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretU32U8 VreinterpretU32U8
//go:noescape
func VreinterpretU32U8(r *arm.Uint32X2, v0 *arm.Uint8X8)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretU32U64 VreinterpretU32U64
//go:noescape
func VreinterpretU32U64(r *arm.Uint32X2, v0 *arm.Uint64X1)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretU32U16 VreinterpretU32U16
//go:noescape
func VreinterpretU32U16(r *arm.Uint32X2, v0 *arm.Uint16X4)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretU32S8 VreinterpretU32S8
//go:noescape
func VreinterpretU32S8(r *arm.Uint32X2, v0 *arm.Int8X8)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretU32F64 VreinterpretU32F64
//go:noescape
func VreinterpretU32F64(r *arm.Uint32X2, v0 *arm.Float64X1)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretU32F32 VreinterpretU32F32
//go:noescape
func VreinterpretU32F32(r *arm.Uint32X2, v0 *arm.Float32X2)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretU32S32 VreinterpretU32S32
//go:noescape
func VreinterpretU32S32(r *arm.Uint32X2, v0 *arm.Int32X2)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretU32S64 VreinterpretU32S64
//go:noescape
func VreinterpretU32S64(r *arm.Uint32X2, v0 *arm.Int64X1)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretU32S16 VreinterpretU32S16
//go:noescape
func VreinterpretU32S16(r *arm.Uint32X2, v0 *arm.Int16X4)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretU64P8 VreinterpretU64P8
//go:noescape
func VreinterpretU64P8(r *arm.Uint64X1, v0 *arm.Poly8X8)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretU64P64 VreinterpretU64P64
//go:noescape
func VreinterpretU64P64(r *arm.Uint64X1, v0 *arm.Poly64X1)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretU64P16 VreinterpretU64P16
//go:noescape
func VreinterpretU64P16(r *arm.Uint64X1, v0 *arm.Poly16X4)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretU64U8 VreinterpretU64U8
//go:noescape
func VreinterpretU64U8(r *arm.Uint64X1, v0 *arm.Uint8X8)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretU64U32 VreinterpretU64U32
//go:noescape
func VreinterpretU64U32(r *arm.Uint64X1, v0 *arm.Uint32X2)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretU64U16 VreinterpretU64U16
//go:noescape
func VreinterpretU64U16(r *arm.Uint64X1, v0 *arm.Uint16X4)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretU64S8 VreinterpretU64S8
//go:noescape
func VreinterpretU64S8(r *arm.Uint64X1, v0 *arm.Int8X8)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretU64F64 VreinterpretU64F64
//go:noescape
func VreinterpretU64F64(r *arm.Uint64X1, v0 *arm.Float64X1)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretU64F32 VreinterpretU64F32
//go:noescape
func VreinterpretU64F32(r *arm.Uint64X1, v0 *arm.Float32X2)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretU64S32 VreinterpretU64S32
//go:noescape
func VreinterpretU64S32(r *arm.Uint64X1, v0 *arm.Int32X2)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretU64S64 VreinterpretU64S64
//go:noescape
func VreinterpretU64S64(r *arm.Uint64X1, v0 *arm.Int64X1)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretU64S16 VreinterpretU64S16
//go:noescape
func VreinterpretU64S16(r *arm.Uint64X1, v0 *arm.Int16X4)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretU16P8 VreinterpretU16P8
//go:noescape
func VreinterpretU16P8(r *arm.Uint16X4, v0 *arm.Poly8X8)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretU16P64 VreinterpretU16P64
//go:noescape
func VreinterpretU16P64(r *arm.Uint16X4, v0 *arm.Poly64X1)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretU16P16 VreinterpretU16P16
//go:noescape
func VreinterpretU16P16(r *arm.Uint16X4, v0 *arm.Poly16X4)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretU16U8 VreinterpretU16U8
//go:noescape
func VreinterpretU16U8(r *arm.Uint16X4, v0 *arm.Uint8X8)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretU16U32 VreinterpretU16U32
//go:noescape
func VreinterpretU16U32(r *arm.Uint16X4, v0 *arm.Uint32X2)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretU16U64 VreinterpretU16U64
//go:noescape
func VreinterpretU16U64(r *arm.Uint16X4, v0 *arm.Uint64X1)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretU16S8 VreinterpretU16S8
//go:noescape
func VreinterpretU16S8(r *arm.Uint16X4, v0 *arm.Int8X8)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretU16F64 VreinterpretU16F64
//go:noescape
func VreinterpretU16F64(r *arm.Uint16X4, v0 *arm.Float64X1)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretU16F32 VreinterpretU16F32
//go:noescape
func VreinterpretU16F32(r *arm.Uint16X4, v0 *arm.Float32X2)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretU16S32 VreinterpretU16S32
//go:noescape
func VreinterpretU16S32(r *arm.Uint16X4, v0 *arm.Int32X2)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretU16S64 VreinterpretU16S64
//go:noescape
func VreinterpretU16S64(r *arm.Uint16X4, v0 *arm.Int64X1)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretU16S16 VreinterpretU16S16
//go:noescape
func VreinterpretU16S16(r *arm.Uint16X4, v0 *arm.Int16X4)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretS8P8 VreinterpretS8P8
//go:noescape
func VreinterpretS8P8(r *arm.Int8X8, v0 *arm.Poly8X8)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretS8P64 VreinterpretS8P64
//go:noescape
func VreinterpretS8P64(r *arm.Int8X8, v0 *arm.Poly64X1)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretS8P16 VreinterpretS8P16
//go:noescape
func VreinterpretS8P16(r *arm.Int8X8, v0 *arm.Poly16X4)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretS8U8 VreinterpretS8U8
//go:noescape
func VreinterpretS8U8(r *arm.Int8X8, v0 *arm.Uint8X8)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretS8U32 VreinterpretS8U32
//go:noescape
func VreinterpretS8U32(r *arm.Int8X8, v0 *arm.Uint32X2)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretS8U64 VreinterpretS8U64
//go:noescape
func VreinterpretS8U64(r *arm.Int8X8, v0 *arm.Uint64X1)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretS8U16 VreinterpretS8U16
//go:noescape
func VreinterpretS8U16(r *arm.Int8X8, v0 *arm.Uint16X4)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretS8F64 VreinterpretS8F64
//go:noescape
func VreinterpretS8F64(r *arm.Int8X8, v0 *arm.Float64X1)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretS8F32 VreinterpretS8F32
//go:noescape
func VreinterpretS8F32(r *arm.Int8X8, v0 *arm.Float32X2)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretS8S32 VreinterpretS8S32
//go:noescape
func VreinterpretS8S32(r *arm.Int8X8, v0 *arm.Int32X2)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretS8S64 VreinterpretS8S64
//go:noescape
func VreinterpretS8S64(r *arm.Int8X8, v0 *arm.Int64X1)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretS8S16 VreinterpretS8S16
//go:noescape
func VreinterpretS8S16(r *arm.Int8X8, v0 *arm.Int16X4)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretF64P8 VreinterpretF64P8
//go:noescape
func VreinterpretF64P8(r *arm.Float64X1, v0 *arm.Poly8X8)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretF64P64 VreinterpretF64P64
//go:noescape
func VreinterpretF64P64(r *arm.Float64X1, v0 *arm.Poly64X1)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretF64P16 VreinterpretF64P16
//go:noescape
func VreinterpretF64P16(r *arm.Float64X1, v0 *arm.Poly16X4)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretF64U8 VreinterpretF64U8
//go:noescape
func VreinterpretF64U8(r *arm.Float64X1, v0 *arm.Uint8X8)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretF64U32 VreinterpretF64U32
//go:noescape
func VreinterpretF64U32(r *arm.Float64X1, v0 *arm.Uint32X2)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretF64U64 VreinterpretF64U64
//go:noescape
func VreinterpretF64U64(r *arm.Float64X1, v0 *arm.Uint64X1)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretF64U16 VreinterpretF64U16
//go:noescape
func VreinterpretF64U16(r *arm.Float64X1, v0 *arm.Uint16X4)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretF64S8 VreinterpretF64S8
//go:noescape
func VreinterpretF64S8(r *arm.Float64X1, v0 *arm.Int8X8)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretF64F32 VreinterpretF64F32
//go:noescape
func VreinterpretF64F32(r *arm.Float64X1, v0 *arm.Float32X2)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretF64S32 VreinterpretF64S32
//go:noescape
func VreinterpretF64S32(r *arm.Float64X1, v0 *arm.Int32X2)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretF64S64 VreinterpretF64S64
//go:noescape
func VreinterpretF64S64(r *arm.Float64X1, v0 *arm.Int64X1)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretF64S16 VreinterpretF64S16
//go:noescape
func VreinterpretF64S16(r *arm.Float64X1, v0 *arm.Int16X4)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretF32P8 VreinterpretF32P8
//go:noescape
func VreinterpretF32P8(r *arm.Float32X2, v0 *arm.Poly8X8)

// vreinterpret_f32_p64
//
//go:linkname VreinterpretF32P64 VreinterpretF32P64
//go:noescape
func VreinterpretF32P64(r *arm.Float32X2, v0 *arm.Poly64X1)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretF32P16 VreinterpretF32P16
//go:noescape
func VreinterpretF32P16(r *arm.Float32X2, v0 *arm.Poly16X4)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretF32U8 VreinterpretF32U8
//go:noescape
func VreinterpretF32U8(r *arm.Float32X2, v0 *arm.Uint8X8)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretF32U32 VreinterpretF32U32
//go:noescape
func VreinterpretF32U32(r *arm.Float32X2, v0 *arm.Uint32X2)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretF32U64 VreinterpretF32U64
//go:noescape
func VreinterpretF32U64(r *arm.Float32X2, v0 *arm.Uint64X1)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretF32U16 VreinterpretF32U16
//go:noescape
func VreinterpretF32U16(r *arm.Float32X2, v0 *arm.Uint16X4)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretF32S8 VreinterpretF32S8
//go:noescape
func VreinterpretF32S8(r *arm.Float32X2, v0 *arm.Int8X8)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretF32F64 VreinterpretF32F64
//go:noescape
func VreinterpretF32F64(r *arm.Float32X2, v0 *arm.Float64X1)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretF32S32 VreinterpretF32S32
//go:noescape
func VreinterpretF32S32(r *arm.Float32X2, v0 *arm.Int32X2)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretF32S64 VreinterpretF32S64
//go:noescape
func VreinterpretF32S64(r *arm.Float32X2, v0 *arm.Int64X1)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretF32S16 VreinterpretF32S16
//go:noescape
func VreinterpretF32S16(r *arm.Float32X2, v0 *arm.Int16X4)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretS32P8 VreinterpretS32P8
//go:noescape
func VreinterpretS32P8(r *arm.Int32X2, v0 *arm.Poly8X8)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretS32P64 VreinterpretS32P64
//go:noescape
func VreinterpretS32P64(r *arm.Int32X2, v0 *arm.Poly64X1)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretS32P16 VreinterpretS32P16
//go:noescape
func VreinterpretS32P16(r *arm.Int32X2, v0 *arm.Poly16X4)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretS32U8 VreinterpretS32U8
//go:noescape
func VreinterpretS32U8(r *arm.Int32X2, v0 *arm.Uint8X8)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretS32U32 VreinterpretS32U32
//go:noescape
func VreinterpretS32U32(r *arm.Int32X2, v0 *arm.Uint32X2)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretS32U64 VreinterpretS32U64
//go:noescape
func VreinterpretS32U64(r *arm.Int32X2, v0 *arm.Uint64X1)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretS32U16 VreinterpretS32U16
//go:noescape
func VreinterpretS32U16(r *arm.Int32X2, v0 *arm.Uint16X4)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretS32S8 VreinterpretS32S8
//go:noescape
func VreinterpretS32S8(r *arm.Int32X2, v0 *arm.Int8X8)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretS32F64 VreinterpretS32F64
//go:noescape
func VreinterpretS32F64(r *arm.Int32X2, v0 *arm.Float64X1)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretS32F32 VreinterpretS32F32
//go:noescape
func VreinterpretS32F32(r *arm.Int32X2, v0 *arm.Float32X2)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretS32S64 VreinterpretS32S64
//go:noescape
func VreinterpretS32S64(r *arm.Int32X2, v0 *arm.Int64X1)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretS32S16 VreinterpretS32S16
//go:noescape
func VreinterpretS32S16(r *arm.Int32X2, v0 *arm.Int16X4)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretS64P8 VreinterpretS64P8
//go:noescape
func VreinterpretS64P8(r *arm.Int64X1, v0 *arm.Poly8X8)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretS64P64 VreinterpretS64P64
//go:noescape
func VreinterpretS64P64(r *arm.Int64X1, v0 *arm.Poly64X1)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretS64P16 VreinterpretS64P16
//go:noescape
func VreinterpretS64P16(r *arm.Int64X1, v0 *arm.Poly16X4)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretS64U8 VreinterpretS64U8
//go:noescape
func VreinterpretS64U8(r *arm.Int64X1, v0 *arm.Uint8X8)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretS64U32 VreinterpretS64U32
//go:noescape
func VreinterpretS64U32(r *arm.Int64X1, v0 *arm.Uint32X2)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretS64U64 VreinterpretS64U64
//go:noescape
func VreinterpretS64U64(r *arm.Int64X1, v0 *arm.Uint64X1)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretS64U16 VreinterpretS64U16
//go:noescape
func VreinterpretS64U16(r *arm.Int64X1, v0 *arm.Uint16X4)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretS64S8 VreinterpretS64S8
//go:noescape
func VreinterpretS64S8(r *arm.Int64X1, v0 *arm.Int8X8)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretS64F64 VreinterpretS64F64
//go:noescape
func VreinterpretS64F64(r *arm.Int64X1, v0 *arm.Float64X1)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretS64F32 VreinterpretS64F32
//go:noescape
func VreinterpretS64F32(r *arm.Int64X1, v0 *arm.Float32X2)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretS64S32 VreinterpretS64S32
//go:noescape
func VreinterpretS64S32(r *arm.Int64X1, v0 *arm.Int32X2)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretS64S16 VreinterpretS64S16
//go:noescape
func VreinterpretS64S16(r *arm.Int64X1, v0 *arm.Int16X4)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretS16P8 VreinterpretS16P8
//go:noescape
func VreinterpretS16P8(r *arm.Int16X4, v0 *arm.Poly8X8)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretS16P64 VreinterpretS16P64
//go:noescape
func VreinterpretS16P64(r *arm.Int16X4, v0 *arm.Poly64X1)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretS16P16 VreinterpretS16P16
//go:noescape
func VreinterpretS16P16(r *arm.Int16X4, v0 *arm.Poly16X4)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretS16U8 VreinterpretS16U8
//go:noescape
func VreinterpretS16U8(r *arm.Int16X4, v0 *arm.Uint8X8)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretS16U32 VreinterpretS16U32
//go:noescape
func VreinterpretS16U32(r *arm.Int16X4, v0 *arm.Uint32X2)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretS16U64 VreinterpretS16U64
//go:noescape
func VreinterpretS16U64(r *arm.Int16X4, v0 *arm.Uint64X1)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretS16U16 VreinterpretS16U16
//go:noescape
func VreinterpretS16U16(r *arm.Int16X4, v0 *arm.Uint16X4)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretS16S8 VreinterpretS16S8
//go:noescape
func VreinterpretS16S8(r *arm.Int16X4, v0 *arm.Int8X8)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretS16F64 VreinterpretS16F64
//go:noescape
func VreinterpretS16F64(r *arm.Int16X4, v0 *arm.Float64X1)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretS16F32 VreinterpretS16F32
//go:noescape
func VreinterpretS16F32(r *arm.Int16X4, v0 *arm.Float32X2)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretS16S32 VreinterpretS16S32
//go:noescape
func VreinterpretS16S32(r *arm.Int16X4, v0 *arm.Int32X2)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretS16S64 VreinterpretS16S64
//go:noescape
func VreinterpretS16S64(r *arm.Int16X4, v0 *arm.Int64X1)

// Floating-point Round to Integral, toward Zero (vector). This instruction rounds a vector of floating-point values in the SIMD&FP source register to integral floating-point values of the same size using the Round towards Zero rounding mode, and writes the result to the SIMD&FP destination register.
//
//go:linkname VrndqF64 VrndqF64
//go:noescape
func VrndqF64(r *arm.Float64X2, v0 *arm.Float64X2)

// Floating-point Round to Integral, toward Zero (vector). This instruction rounds a vector of floating-point values in the SIMD&FP source register to integral floating-point values of the same size using the Round towards Zero rounding mode, and writes the result to the SIMD&FP destination register.
//
//go:linkname VrndF64 VrndF64
//go:noescape
func VrndF64(r *arm.Float64X1, v0 *arm.Float64X1)

// Floating-point Round to Integral, to nearest with ties to Away (vector). This instruction rounds a vector of floating-point values in the SIMD&FP source register to integral floating-point values of the same size using the Round to Nearest with Ties to Away rounding mode, and writes the result to the SIMD&FP destination register.
//
//go:linkname VrndaqF64 VrndaqF64
//go:noescape
func VrndaqF64(r *arm.Float64X2, v0 *arm.Float64X2)

// Floating-point Round to Integral, to nearest with ties to Away (vector). This instruction rounds a vector of floating-point values in the SIMD&FP source register to integral floating-point values of the same size using the Round to Nearest with Ties to Away rounding mode, and writes the result to the SIMD&FP destination register.
//
//go:linkname VrndaF64 VrndaF64
//go:noescape
func VrndaF64(r *arm.Float64X1, v0 *arm.Float64X1)

// Floating-point Round to Integral, using current rounding mode (vector). This instruction rounds a vector of floating-point values in the SIMD&FP source register to integral floating-point values of the same size using the rounding mode that is determined by the FPCR, and writes the result to the SIMD&FP destination register.
//
//go:linkname VrndiqF64 VrndiqF64
//go:noescape
func VrndiqF64(r *arm.Float64X2, v0 *arm.Float64X2)

// Floating-point Round to Integral, using current rounding mode (vector). This instruction rounds a vector of floating-point values in the SIMD&FP source register to integral floating-point values of the same size using the rounding mode that is determined by the FPCR, and writes the result to the SIMD&FP destination register.
//
//go:linkname VrndiF64 VrndiF64
//go:noescape
func VrndiF64(r *arm.Float64X1, v0 *arm.Float64X1)

// Floating-point Round to Integral, toward Minus infinity (vector). This instruction rounds a vector of floating-point values in the SIMD&FP source register to integral floating-point values of the same size using the Round towards Minus Infinity rounding mode, and writes the result to the SIMD&FP destination register.
//
//go:linkname VrndmqF64 VrndmqF64
//go:noescape
func VrndmqF64(r *arm.Float64X2, v0 *arm.Float64X2)

// Floating-point Round to Integral, toward Minus infinity (vector). This instruction rounds a vector of floating-point values in the SIMD&FP source register to integral floating-point values of the same size using the Round towards Minus Infinity rounding mode, and writes the result to the SIMD&FP destination register.
//
//go:linkname VrndmF64 VrndmF64
//go:noescape
func VrndmF64(r *arm.Float64X1, v0 *arm.Float64X1)

// Floating-point Round to Integral, to nearest with ties to even (vector). This instruction rounds a vector of floating-point values in the SIMD&FP source register to integral floating-point values of the same size using the Round to Nearest rounding mode, and writes the result to the SIMD&FP destination register.
//
//go:linkname VrndnqF64 VrndnqF64
//go:noescape
func VrndnqF64(r *arm.Float64X2, v0 *arm.Float64X2)

// Floating-point Round to Integral, to nearest with ties to even (vector). This instruction rounds a vector of floating-point values in the SIMD&FP source register to integral floating-point values of the same size using the Round to Nearest rounding mode, and writes the result to the SIMD&FP destination register.
//
//go:linkname VrndnF64 VrndnF64
//go:noescape
func VrndnF64(r *arm.Float64X1, v0 *arm.Float64X1)

// Floating-point Round to Integral, toward Plus infinity (vector). This instruction rounds a vector of floating-point values in the SIMD&FP source register to integral floating-point values of the same size using the Round towards Plus Infinity rounding mode, and writes the result to the SIMD&FP destination register.
//
//go:linkname VrndpqF64 VrndpqF64
//go:noescape
func VrndpqF64(r *arm.Float64X2, v0 *arm.Float64X2)

// Floating-point Round to Integral, toward Plus infinity (vector). This instruction rounds a vector of floating-point values in the SIMD&FP source register to integral floating-point values of the same size using the Round towards Plus Infinity rounding mode, and writes the result to the SIMD&FP destination register.
//
//go:linkname VrndpF64 VrndpF64
//go:noescape
func VrndpF64(r *arm.Float64X1, v0 *arm.Float64X1)

// Floating-point Round to Integral exact, using current rounding mode (vector). This instruction rounds a vector of floating-point values in the SIMD&FP source register to integral floating-point values of the same size using the rounding mode that is determined by the FPCR, and writes the result to the SIMD&FP destination register.
//
//go:linkname VrndxqF64 VrndxqF64
//go:noescape
func VrndxqF64(r *arm.Float64X2, v0 *arm.Float64X2)

// Floating-point Round to Integral exact, using current rounding mode (vector). This instruction rounds a vector of floating-point values in the SIMD&FP source register to integral floating-point values of the same size using the rounding mode that is determined by the FPCR, and writes the result to the SIMD&FP destination register.
//
//go:linkname VrndxF64 VrndxF64
//go:noescape
func VrndxF64(r *arm.Float64X1, v0 *arm.Float64X1)

// Floating-point Round to 32-bit Integer, using current rounding mode (vector). This instruction rounds a vector of floating-point values in the SIMD&FP source register to integral floating-point values that fit into a 32-bit integer size using the rounding mode that is determined by the FPCR, and writes the result to the SIMD&FP destination register.
//
//go:linkname Vrnd32XqF32 Vrnd32XqF32
//go:noescape
func Vrnd32XqF32(r *arm.Float32X4, v0 *arm.Float32X4)

// Floating-point Round to 32-bit Integer, using current rounding mode (vector). This instruction rounds a vector of floating-point values in the SIMD&FP source register to integral floating-point values that fit into a 32-bit integer size using the rounding mode that is determined by the FPCR, and writes the result to the SIMD&FP destination register.
//
//go:linkname Vrnd32XF32 Vrnd32XF32
//go:noescape
func Vrnd32XF32(r *arm.Float32X2, v0 *arm.Float32X2)

// Floating-point Round to 32-bit Integer toward Zero (vector). This instruction rounds a vector of floating-point values in the SIMD&FP source register to integral floating-point values that fit into a 32-bit integer size using the Round towards Zero rounding mode, and writes the result to the SIMD&FP destination register.
//
//go:linkname Vrnd32ZqF32 Vrnd32ZqF32
//go:noescape
func Vrnd32ZqF32(r *arm.Float32X4, v0 *arm.Float32X4)

// Floating-point Round to 32-bit Integer toward Zero (vector). This instruction rounds a vector of floating-point values in the SIMD&FP source register to integral floating-point values that fit into a 32-bit integer size using the Round towards Zero rounding mode, and writes the result to the SIMD&FP destination register.
//
//go:linkname Vrnd32ZF32 Vrnd32ZF32
//go:noescape
func Vrnd32ZF32(r *arm.Float32X2, v0 *arm.Float32X2)

// Floating-point Round to 64-bit Integer, using current rounding mode (vector). This instruction rounds a vector of floating-point values in the SIMD&FP source register to integral floating-point values that fit into a 64-bit integer size using the rounding mode that is determined by the FPCR, and writes the result to the SIMD&FP destination register.
//
//go:linkname Vrnd64XqF32 Vrnd64XqF32
//go:noescape
func Vrnd64XqF32(r *arm.Float32X4, v0 *arm.Float32X4)

// Floating-point Round to 64-bit Integer, using current rounding mode (vector). This instruction rounds a vector of floating-point values in the SIMD&FP source register to integral floating-point values that fit into a 64-bit integer size using the rounding mode that is determined by the FPCR, and writes the result to the SIMD&FP destination register.
//
//go:linkname Vrnd64XF32 Vrnd64XF32
//go:noescape
func Vrnd64XF32(r *arm.Float32X2, v0 *arm.Float32X2)

// Floating-point Round to 64-bit Integer toward Zero (vector). This instruction rounds a vector of floating-point values in the SIMD&FP source register to integral floating-point values that fit into a 64-bit integer size using the Round towards Zero rounding mode, and writes the result to the SIMD&FP destination register.
//
//go:linkname Vrnd64ZqF32 Vrnd64ZqF32
//go:noescape
func Vrnd64ZqF32(r *arm.Float32X4, v0 *arm.Float32X4)

// Floating-point Round to 64-bit Integer toward Zero (vector). This instruction rounds a vector of floating-point values in the SIMD&FP source register to integral floating-point values that fit into a 64-bit integer size using the Round towards Zero rounding mode, and writes the result to the SIMD&FP destination register.
//
//go:linkname Vrnd64ZF32 Vrnd64ZF32
//go:noescape
func Vrnd64ZF32(r *arm.Float32X2, v0 *arm.Float32X2)

// Floating-point Maximum Number (vector). This instruction compares corresponding vector elements in the two source SIMD&FP registers, writes the larger of the two floating-point values into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VmaxnmqF64 VmaxnmqF64
//go:noescape
func VmaxnmqF64(r *arm.Float64X2, v0 *arm.Float64X2, v1 *arm.Float64X2)

// Floating-point Maximum Number (vector). This instruction compares corresponding vector elements in the two source SIMD&FP registers, writes the larger of the two floating-point values into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VmaxnmF64 VmaxnmF64
//go:noescape
func VmaxnmF64(r *arm.Float64X1, v0 *arm.Float64X1, v1 *arm.Float64X1)

// Floating-point Minimum Number (vector). This instruction compares corresponding vector elements in the two source SIMD&FP registers, writes the smaller of the two floating-point values into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VminnmqF64 VminnmqF64
//go:noescape
func VminnmqF64(r *arm.Float64X2, v0 *arm.Float64X2, v1 *arm.Float64X2)

// Floating-point Minimum Number (vector). This instruction compares corresponding vector elements in the two source SIMD&FP registers, writes the smaller of the two floating-point values into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VminnmF64 VminnmF64
//go:noescape
func VminnmF64(r *arm.Float64X1, v0 *arm.Float64X1, v1 *arm.Float64X1)

// Floating-point Complex Add.
//
//go:linkname VcaddRot270F32 VcaddRot270F32
//go:noescape
func VcaddRot270F32(r *arm.Float32X2, v0 *arm.Float32X2, v1 *arm.Float32X2)

// Floating-point Complex Add.
//
//go:linkname VcaddRot90F32 VcaddRot90F32
//go:noescape
func VcaddRot90F32(r *arm.Float32X2, v0 *arm.Float32X2, v1 *arm.Float32X2)

// Floating-point Complex Add.
//
//go:linkname VcaddqRot270F32 VcaddqRot270F32
//go:noescape
func VcaddqRot270F32(r *arm.Float32X4, v0 *arm.Float32X4, v1 *arm.Float32X4)

// Floating-point Complex Add.
//
//go:linkname VcaddqRot90F32 VcaddqRot90F32
//go:noescape
func VcaddqRot90F32(r *arm.Float32X4, v0 *arm.Float32X4, v1 *arm.Float32X4)

// Floating-point Complex Add.
//
//go:linkname VcaddqRot270F64 VcaddqRot270F64
//go:noescape
func VcaddqRot270F64(r *arm.Float64X2, v0 *arm.Float64X2, v1 *arm.Float64X2)

// Floating-point Complex Add.
//
//go:linkname VcaddqRot90F64 VcaddqRot90F64
//go:noescape
func VcaddqRot90F64(r *arm.Float64X2, v0 *arm.Float64X2, v1 *arm.Float64X2)

// Dot Product unsigned arithmetic (vector). This instruction performs the dot product of the four unsigned 8-bit elements in each 32-bit element of the first source register with the four unsigned 8-bit elements of the corresponding 32-bit element in the second source register, accumulating the result into the corresponding 32-bit element of the destination register.
//
//go:linkname VdotqU32 VdotqU32
//go:noescape
func VdotqU32(r *arm.Uint32X4, v0 *arm.Uint32X4, v1 *arm.Uint8X16, v2 *arm.Uint8X16)

// Dot Product signed arithmetic (vector). This instruction performs the dot product of the four signed 8-bit elements in each 32-bit element of the first source register with the four signed 8-bit elements of the corresponding 32-bit element in the second source register, accumulating the result into the corresponding 32-bit element of the destination register.
//
//go:linkname VdotqS32 VdotqS32
//go:noescape
func VdotqS32(r *arm.Int32X4, v0 *arm.Int32X4, v1 *arm.Int8X16, v2 *arm.Int8X16)

// Dot Product unsigned arithmetic (vector). This instruction performs the dot product of the four unsigned 8-bit elements in each 32-bit element of the first source register with the four unsigned 8-bit elements of the corresponding 32-bit element in the second source register, accumulating the result into the corresponding 32-bit element of the destination register.
//
//go:linkname VdotU32 VdotU32
//go:noescape
func VdotU32(r *arm.Uint32X2, v0 *arm.Uint32X2, v1 *arm.Uint8X8, v2 *arm.Uint8X8)

// Dot Product signed arithmetic (vector). This instruction performs the dot product of the four signed 8-bit elements in each 32-bit element of the first source register with the four signed 8-bit elements of the corresponding 32-bit element in the second source register, accumulating the result into the corresponding 32-bit element of the destination register.
//
//go:linkname VdotS32 VdotS32
//go:noescape
func VdotS32(r *arm.Int32X2, v0 *arm.Int32X2, v1 *arm.Int8X8, v2 *arm.Int8X8)

// Floating-point fused Multiply-Add to accumulator (vector). This instruction multiplies corresponding floating-point values in the vectors in the two source SIMD&FP registers, adds the product to the corresponding vector element of the destination SIMD&FP register, and writes the result to the destination SIMD&FP register.
//
//go:linkname VfmaqF32 VfmaqF32
//go:noescape
func VfmaqF32(r *arm.Float32X4, v0 *arm.Float32X4, v1 *arm.Float32X4, v2 *arm.Float32X4)

// Floating-point fused Multiply-Add to accumulator (vector). This instruction multiplies corresponding floating-point values in the vectors in the two source SIMD&FP registers, adds the product to the corresponding vector element of the destination SIMD&FP register, and writes the result to the destination SIMD&FP register.
//
//go:linkname VfmaF32 VfmaF32
//go:noescape
func VfmaF32(r *arm.Float32X2, v0 *arm.Float32X2, v1 *arm.Float32X2, v2 *arm.Float32X2)

// Floating-point fused Multiply-Add to accumulator (vector). This instruction multiplies corresponding floating-point values in the vectors in the two source SIMD&FP registers, adds the product to the corresponding vector element of the destination SIMD&FP register, and writes the result to the destination SIMD&FP register.
//
//go:linkname VfmaqNF32 VfmaqNF32
//go:noescape
func VfmaqNF32(r *arm.Float32X4, v0 *arm.Float32X4, v1 *arm.Float32X4, v2 *arm.Float32)

// Floating-point fused Multiply-Add to accumulator (vector). This instruction multiplies corresponding floating-point values in the vectors in the two source SIMD&FP registers, adds the product to the corresponding vector element of the destination SIMD&FP register, and writes the result to the destination SIMD&FP register.
//
//go:linkname VfmaNF32 VfmaNF32
//go:noescape
func VfmaNF32(r *arm.Float32X2, v0 *arm.Float32X2, v1 *arm.Float32X2, v2 *arm.Float32)

// Floating-point fused Multiply-Subtract from accumulator (vector). This instruction multiplies corresponding floating-point values in the vectors in the two source SIMD&FP registers, negates the product, adds the result to the corresponding vector element of the destination SIMD&FP register, and writes the result to the destination SIMD&FP register.
//
//go:linkname VfmsqF32 VfmsqF32
//go:noescape
func VfmsqF32(r *arm.Float32X4, v0 *arm.Float32X4, v1 *arm.Float32X4, v2 *arm.Float32X4)

// Floating-point fused Multiply-Subtract from accumulator (vector). This instruction multiplies corresponding floating-point values in the vectors in the two source SIMD&FP registers, negates the product, adds the result to the corresponding vector element of the destination SIMD&FP register, and writes the result to the destination SIMD&FP register.
//
//go:linkname VfmsF32 VfmsF32
//go:noescape
func VfmsF32(r *arm.Float32X2, v0 *arm.Float32X2, v1 *arm.Float32X2, v2 *arm.Float32X2)

// Signed Saturating Rounding Doubling Multiply Accumulate returning High Half (vector). This instruction multiplies the vector elements of the first source SIMD&FP register with the corresponding vector elements of the second source SIMD&FP register without saturating the multiply results, doubles the results, and accumulates the most significant half of the final results with the vector elements of the destination SIMD&FP register. The results are rounded.
//
//go:linkname VqrdmlahqS32 VqrdmlahqS32
//go:noescape
func VqrdmlahqS32(r *arm.Int32X4, v0 *arm.Int32X4, v1 *arm.Int32X4, v2 *arm.Int32X4)

// Signed Saturating Rounding Doubling Multiply Accumulate returning High Half (vector). This instruction multiplies the vector elements of the first source SIMD&FP register with the corresponding vector elements of the second source SIMD&FP register without saturating the multiply results, doubles the results, and accumulates the most significant half of the final results with the vector elements of the destination SIMD&FP register. The results are rounded.
//
//go:linkname VqrdmlahqS16 VqrdmlahqS16
//go:noescape
func VqrdmlahqS16(r *arm.Int16X8, v0 *arm.Int16X8, v1 *arm.Int16X8, v2 *arm.Int16X8)

// Signed Saturating Rounding Doubling Multiply Accumulate returning High Half (vector). This instruction multiplies the vector elements of the first source SIMD&FP register with the corresponding vector elements of the second source SIMD&FP register without saturating the multiply results, doubles the results, and accumulates the most significant half of the final results with the vector elements of the destination SIMD&FP register. The results are rounded.
//
//go:linkname VqrdmlahS32 VqrdmlahS32
//go:noescape
func VqrdmlahS32(r *arm.Int32X2, v0 *arm.Int32X2, v1 *arm.Int32X2, v2 *arm.Int32X2)

// Signed Saturating Rounding Doubling Multiply Accumulate returning High Half (vector). This instruction multiplies the vector elements of the first source SIMD&FP register with the corresponding vector elements of the second source SIMD&FP register without saturating the multiply results, doubles the results, and accumulates the most significant half of the final results with the vector elements of the destination SIMD&FP register. The results are rounded.
//
//go:linkname VqrdmlahS16 VqrdmlahS16
//go:noescape
func VqrdmlahS16(r *arm.Int16X4, v0 *arm.Int16X4, v1 *arm.Int16X4, v2 *arm.Int16X4)

// Signed Saturating Rounding Doubling Multiply Subtract returning High Half (vector). This instruction multiplies the vector elements of the first source SIMD&FP register with the corresponding vector elements of the second source SIMD&FP register without saturating the multiply results, doubles the results, and subtracts the most significant half of the final results from the vector elements of the destination SIMD&FP register. The results are rounded.
//
//go:linkname VqrdmlshqS32 VqrdmlshqS32
//go:noescape
func VqrdmlshqS32(r *arm.Int32X4, v0 *arm.Int32X4, v1 *arm.Int32X4, v2 *arm.Int32X4)

// Signed Saturating Rounding Doubling Multiply Subtract returning High Half (vector). This instruction multiplies the vector elements of the first source SIMD&FP register with the corresponding vector elements of the second source SIMD&FP register without saturating the multiply results, doubles the results, and subtracts the most significant half of the final results from the vector elements of the destination SIMD&FP register. The results are rounded.
//
//go:linkname VqrdmlshqS16 VqrdmlshqS16
//go:noescape
func VqrdmlshqS16(r *arm.Int16X8, v0 *arm.Int16X8, v1 *arm.Int16X8, v2 *arm.Int16X8)

// Signed Saturating Rounding Doubling Multiply Subtract returning High Half (vector). This instruction multiplies the vector elements of the first source SIMD&FP register with the corresponding vector elements of the second source SIMD&FP register without saturating the multiply results, doubles the results, and subtracts the most significant half of the final results from the vector elements of the destination SIMD&FP register. The results are rounded.
//
//go:linkname VqrdmlshS32 VqrdmlshS32
//go:noescape
func VqrdmlshS32(r *arm.Int32X2, v0 *arm.Int32X2, v1 *arm.Int32X2, v2 *arm.Int32X2)

// Signed Saturating Rounding Doubling Multiply Subtract returning High Half (vector). This instruction multiplies the vector elements of the first source SIMD&FP register with the corresponding vector elements of the second source SIMD&FP register without saturating the multiply results, doubles the results, and subtracts the most significant half of the final results from the vector elements of the destination SIMD&FP register. The results are rounded.
//
//go:linkname VqrdmlshS16 VqrdmlshS16
//go:noescape
func VqrdmlshS16(r *arm.Int16X4, v0 *arm.Int16X4, v1 *arm.Int16X4, v2 *arm.Int16X4)

// Floating-point Absolute Difference (vector). This instruction subtracts the floating-point values in the elements of the second source SIMD&FP register, from the corresponding floating-point values in the elements of the first source SIMD&FP register, places the absolute value of each result in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VabdqF64 VabdqF64
//go:noescape
func VabdqF64(r *arm.Float64X2, v0 *arm.Float64X2, v1 *arm.Float64X2)

// Floating-point Absolute Difference (vector). This instruction subtracts the floating-point values in the elements of the second source SIMD&FP register, from the corresponding floating-point values in the elements of the first source SIMD&FP register, places the absolute value of each result in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VabdF64 VabdF64
//go:noescape
func VabdF64(r *arm.Float64X1, v0 *arm.Float64X1, v1 *arm.Float64X1)

// Floating-point Absolute Difference (vector). This instruction subtracts the floating-point values in the elements of the second source SIMD&FP register, from the corresponding floating-point values in the elements of the first source SIMD&FP register, places the absolute value of each result in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VabddF64 VabddF64
//go:noescape
func VabddF64(r *arm.Float64, v0 *arm.Float64, v1 *arm.Float64)

// Floating-point Absolute Difference (vector). This instruction subtracts the floating-point values in the elements of the second source SIMD&FP register, from the corresponding floating-point values in the elements of the first source SIMD&FP register, places the absolute value of each result in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VabdsF32 VabdsF32
//go:noescape
func VabdsF32(r *arm.Float32, v0 *arm.Float32, v1 *arm.Float32)

// Floating-point Absolute value (vector). This instruction calculates the absolute value of each vector element in the source SIMD&FP register, writes the result to a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VabsqF64 VabsqF64
//go:noescape
func VabsqF64(r *arm.Float64X2, v0 *arm.Float64X2)

// Absolute value (vector). This instruction calculates the absolute value of each vector element in the source SIMD&FP register, puts the result into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VabsqS64 VabsqS64
//go:noescape
func VabsqS64(r *arm.Int64X2, v0 *arm.Int64X2)

// Floating-point Absolute value (vector). This instruction calculates the absolute value of each vector element in the source SIMD&FP register, writes the result to a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VabsF64 VabsF64
//go:noescape
func VabsF64(r *arm.Float64X1, v0 *arm.Float64X1)

// Absolute value (vector). This instruction calculates the absolute value of each vector element in the source SIMD&FP register, puts the result into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VabsS64 VabsS64
//go:noescape
func VabsS64(r *arm.Int64X1, v0 *arm.Int64X1)

// Absolute value (vector). This instruction calculates the absolute value of each vector element in the source SIMD&FP register, puts the result into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VabsdS64 VabsdS64
//go:noescape
func VabsdS64(r *arm.Int64, v0 *arm.Int64)

// Floating-point Add (vector). This instruction adds corresponding vector elements in the two source SIMD&FP registers, writes the result into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are floating-point values.
//
//go:linkname VaddqF64 VaddqF64
//go:noescape
func VaddqF64(r *arm.Float64X2, v0 *arm.Float64X2, v1 *arm.Float64X2)

// Floating-point Add (vector). This instruction adds corresponding vector elements in the two source SIMD&FP registers, writes the result into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are floating-point values.
//
//go:linkname VaddF64 VaddF64
//go:noescape
func VaddF64(r *arm.Float64X1, v0 *arm.Float64X1, v1 *arm.Float64X1)

// Add (vector). This instruction adds corresponding elements in the two source SIMD&FP registers, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VadddU64 VadddU64
//go:noescape
func VadddU64(r *arm.Uint64, v0 *arm.Uint64, v1 *arm.Uint64)

// Add (vector). This instruction adds corresponding elements in the two source SIMD&FP registers, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VadddS64 VadddS64
//go:noescape
func VadddS64(r *arm.Int64, v0 *arm.Int64, v1 *arm.Int64)

// Bitwise Exclusive OR (vector). This instruction performs a bitwise Exclusive OR operation between the two source SIMD&FP registers, and places the result in the destination SIMD&FP register.
//
//go:linkname VaddqP128 VaddqP128
//go:noescape
func VaddqP128(r *arm.Poly128, v0 *arm.Poly128, v1 *arm.Poly128)

// Add returning High Narrow. This instruction adds each vector element in the first source SIMD&FP register to the corresponding vector element in the second source SIMD&FP register, places the most significant half of the result into a vector, and writes the vector to the lower or upper half of the destination SIMD&FP register.
//
//go:linkname VaddhnHighU32 VaddhnHighU32
//go:noescape
func VaddhnHighU32(r *arm.Uint16X8, v0 *arm.Uint16X4, v1 *arm.Uint32X4, v2 *arm.Uint32X4)

// Add returning High Narrow. This instruction adds each vector element in the first source SIMD&FP register to the corresponding vector element in the second source SIMD&FP register, places the most significant half of the result into a vector, and writes the vector to the lower or upper half of the destination SIMD&FP register.
//
//go:linkname VaddhnHighU64 VaddhnHighU64
//go:noescape
func VaddhnHighU64(r *arm.Uint32X4, v0 *arm.Uint32X2, v1 *arm.Uint64X2, v2 *arm.Uint64X2)

// Add returning High Narrow. This instruction adds each vector element in the first source SIMD&FP register to the corresponding vector element in the second source SIMD&FP register, places the most significant half of the result into a vector, and writes the vector to the lower or upper half of the destination SIMD&FP register.
//
//go:linkname VaddhnHighU16 VaddhnHighU16
//go:noescape
func VaddhnHighU16(r *arm.Uint8X16, v0 *arm.Uint8X8, v1 *arm.Uint16X8, v2 *arm.Uint16X8)

// Add returning High Narrow. This instruction adds each vector element in the first source SIMD&FP register to the corresponding vector element in the second source SIMD&FP register, places the most significant half of the result into a vector, and writes the vector to the lower or upper half of the destination SIMD&FP register.
//
//go:linkname VaddhnHighS32 VaddhnHighS32
//go:noescape
func VaddhnHighS32(r *arm.Int16X8, v0 *arm.Int16X4, v1 *arm.Int32X4, v2 *arm.Int32X4)

// Add returning High Narrow. This instruction adds each vector element in the first source SIMD&FP register to the corresponding vector element in the second source SIMD&FP register, places the most significant half of the result into a vector, and writes the vector to the lower or upper half of the destination SIMD&FP register.
//
//go:linkname VaddhnHighS64 VaddhnHighS64
//go:noescape
func VaddhnHighS64(r *arm.Int32X4, v0 *arm.Int32X2, v1 *arm.Int64X2, v2 *arm.Int64X2)

// Add returning High Narrow. This instruction adds each vector element in the first source SIMD&FP register to the corresponding vector element in the second source SIMD&FP register, places the most significant half of the result into a vector, and writes the vector to the lower or upper half of the destination SIMD&FP register.
//
//go:linkname VaddhnHighS16 VaddhnHighS16
//go:noescape
func VaddhnHighS16(r *arm.Int8X16, v0 *arm.Int8X8, v1 *arm.Int16X8, v2 *arm.Int16X8)

// Unsigned sum Long across Vector. This instruction adds every vector element in the source SIMD&FP register together, and writes the scalar result to the destination SIMD&FP register. The destination scalar is twice as long as the source vector elements. All the values in this instruction are unsigned integer values.
//
//go:linkname VaddlvqU8 VaddlvqU8
//go:noescape
func VaddlvqU8(r *arm.Uint16, v0 *arm.Uint8X16)

// Unsigned sum Long across Vector. This instruction adds every vector element in the source SIMD&FP register together, and writes the scalar result to the destination SIMD&FP register. The destination scalar is twice as long as the source vector elements. All the values in this instruction are unsigned integer values.
//
//go:linkname VaddlvqU32 VaddlvqU32
//go:noescape
func VaddlvqU32(r *arm.Uint64, v0 *arm.Uint32X4)

// Unsigned sum Long across Vector. This instruction adds every vector element in the source SIMD&FP register together, and writes the scalar result to the destination SIMD&FP register. The destination scalar is twice as long as the source vector elements. All the values in this instruction are unsigned integer values.
//
//go:linkname VaddlvqU16 VaddlvqU16
//go:noescape
func VaddlvqU16(r *arm.Uint32, v0 *arm.Uint16X8)

// Signed Add Long across Vector. This instruction adds every vector element in the source SIMD&FP register together, and writes the scalar result to the destination SIMD&FP register. The destination scalar is twice as long as the source vector elements. All the values in this instruction are signed integer values.
//
//go:linkname VaddlvqS8 VaddlvqS8
//go:noescape
func VaddlvqS8(r *arm.Int16, v0 *arm.Int8X16)

// Signed Add Long across Vector. This instruction adds every vector element in the source SIMD&FP register together, and writes the scalar result to the destination SIMD&FP register. The destination scalar is twice as long as the source vector elements. All the values in this instruction are signed integer values.
//
//go:linkname VaddlvqS32 VaddlvqS32
//go:noescape
func VaddlvqS32(r *arm.Int64, v0 *arm.Int32X4)

// Signed Add Long across Vector. This instruction adds every vector element in the source SIMD&FP register together, and writes the scalar result to the destination SIMD&FP register. The destination scalar is twice as long as the source vector elements. All the values in this instruction are signed integer values.
//
//go:linkname VaddlvqS16 VaddlvqS16
//go:noescape
func VaddlvqS16(r *arm.Int32, v0 *arm.Int16X8)

// Unsigned sum Long across Vector. This instruction adds every vector element in the source SIMD&FP register together, and writes the scalar result to the destination SIMD&FP register. The destination scalar is twice as long as the source vector elements. All the values in this instruction are unsigned integer values.
//
//go:linkname VaddlvU8 VaddlvU8
//go:noescape
func VaddlvU8(r *arm.Uint16, v0 *arm.Uint8X8)

// Unsigned Add Long Pairwise. This instruction adds pairs of adjacent unsigned integer values from the vector in the source SIMD&FP register, places the result into a vector, and writes the vector to the destination SIMD&FP register. The destination vector elements are twice as long as the source vector elements.
//
//go:linkname VaddlvU32 VaddlvU32
//go:noescape
func VaddlvU32(r *arm.Uint64, v0 *arm.Uint32X2)

// Unsigned sum Long across Vector. This instruction adds every vector element in the source SIMD&FP register together, and writes the scalar result to the destination SIMD&FP register. The destination scalar is twice as long as the source vector elements. All the values in this instruction are unsigned integer values.
//
//go:linkname VaddlvU16 VaddlvU16
//go:noescape
func VaddlvU16(r *arm.Uint32, v0 *arm.Uint16X4)

// Signed Add Long across Vector. This instruction adds every vector element in the source SIMD&FP register together, and writes the scalar result to the destination SIMD&FP register. The destination scalar is twice as long as the source vector elements. All the values in this instruction are signed integer values.
//
//go:linkname VaddlvS8 VaddlvS8
//go:noescape
func VaddlvS8(r *arm.Int16, v0 *arm.Int8X8)

// Signed Add Long Pairwise. This instruction adds pairs of adjacent signed integer values from the vector in the source SIMD&FP register, places the result into a vector, and writes the vector to the destination SIMD&FP register. The destination vector elements are twice as long as the source vector elements.
//
//go:linkname VaddlvS32 VaddlvS32
//go:noescape
func VaddlvS32(r *arm.Int64, v0 *arm.Int32X2)

// Signed Add Long across Vector. This instruction adds every vector element in the source SIMD&FP register together, and writes the scalar result to the destination SIMD&FP register. The destination scalar is twice as long as the source vector elements. All the values in this instruction are signed integer values.
//
//go:linkname VaddlvS16 VaddlvS16
//go:noescape
func VaddlvS16(r *arm.Int32, v0 *arm.Int16X4)

// Add across Vector. This instruction adds every vector element in the source SIMD&FP register together, and writes the scalar result to the destination SIMD&FP register.
//
//go:linkname VaddvqU8 VaddvqU8
//go:noescape
func VaddvqU8(r *arm.Uint8, v0 *arm.Uint8X16)

// Add across Vector. This instruction adds every vector element in the source SIMD&FP register together, and writes the scalar result to the destination SIMD&FP register.
//
//go:linkname VaddvqU32 VaddvqU32
//go:noescape
func VaddvqU32(r *arm.Uint32, v0 *arm.Uint32X4)

// Add across vector
//
//go:linkname VaddvqU64 VaddvqU64
//go:noescape
func VaddvqU64(r *arm.Uint64, v0 *arm.Uint64X2)

// Add across Vector. This instruction adds every vector element in the source SIMD&FP register together, and writes the scalar result to the destination SIMD&FP register.
//
//go:linkname VaddvqU16 VaddvqU16
//go:noescape
func VaddvqU16(r *arm.Uint16, v0 *arm.Uint16X8)

// Add across Vector. This instruction adds every vector element in the source SIMD&FP register together, and writes the scalar result to the destination SIMD&FP register.
//
//go:linkname VaddvqS8 VaddvqS8
//go:noescape
func VaddvqS8(r *arm.Int8, v0 *arm.Int8X16)

// Floating-point add across vector
//
//go:linkname VaddvqF64 VaddvqF64
//go:noescape
func VaddvqF64(r *arm.Float64, v0 *arm.Float64X2)

// Floating-point add across vector
//
//go:linkname VaddvqF32 VaddvqF32
//go:noescape
func VaddvqF32(r *arm.Float32, v0 *arm.Float32X4)

// Add across Vector. This instruction adds every vector element in the source SIMD&FP register together, and writes the scalar result to the destination SIMD&FP register.
//
//go:linkname VaddvqS32 VaddvqS32
//go:noescape
func VaddvqS32(r *arm.Int32, v0 *arm.Int32X4)

// Add across vector
//
//go:linkname VaddvqS64 VaddvqS64
//go:noescape
func VaddvqS64(r *arm.Int64, v0 *arm.Int64X2)

// Add across Vector. This instruction adds every vector element in the source SIMD&FP register together, and writes the scalar result to the destination SIMD&FP register.
//
//go:linkname VaddvqS16 VaddvqS16
//go:noescape
func VaddvqS16(r *arm.Int16, v0 *arm.Int16X8)

// Add across Vector. This instruction adds every vector element in the source SIMD&FP register together, and writes the scalar result to the destination SIMD&FP register.
//
//go:linkname VaddvU8 VaddvU8
//go:noescape
func VaddvU8(r *arm.Uint8, v0 *arm.Uint8X8)

// Add across vector
//
//go:linkname VaddvU32 VaddvU32
//go:noescape
func VaddvU32(r *arm.Uint32, v0 *arm.Uint32X2)

// Add across Vector. This instruction adds every vector element in the source SIMD&FP register together, and writes the scalar result to the destination SIMD&FP register.
//
//go:linkname VaddvU16 VaddvU16
//go:noescape
func VaddvU16(r *arm.Uint16, v0 *arm.Uint16X4)

// Add across Vector. This instruction adds every vector element in the source SIMD&FP register together, and writes the scalar result to the destination SIMD&FP register.
//
//go:linkname VaddvS8 VaddvS8
//go:noescape
func VaddvS8(r *arm.Int8, v0 *arm.Int8X8)

// Floating-point add across vector
//
//go:linkname VaddvF32 VaddvF32
//go:noescape
func VaddvF32(r *arm.Float32, v0 *arm.Float32X2)

// Add across vector
//
//go:linkname VaddvS32 VaddvS32
//go:noescape
func VaddvS32(r *arm.Int32, v0 *arm.Int32X2)

// Add across Vector. This instruction adds every vector element in the source SIMD&FP register together, and writes the scalar result to the destination SIMD&FP register.
//
//go:linkname VaddvS16 VaddvS16
//go:noescape
func VaddvS16(r *arm.Int16, v0 *arm.Int16X4)

// Bitwise Select. This instruction sets each bit in the destination SIMD&FP register to the corresponding bit from the first source SIMD&FP register when the original destination bit was 1, otherwise from the second source SIMD&FP register.
//
//go:linkname VbslP64 VbslP64
//go:noescape
func VbslP64(r *arm.Poly64X1, v0 *arm.Uint64X1, v1 *arm.Poly64X1, v2 *arm.Poly64X1)

// Bitwise Select. This instruction sets each bit in the destination SIMD&FP register to the corresponding bit from the first source SIMD&FP register when the original destination bit was 1, otherwise from the second source SIMD&FP register.
//
//go:linkname VbslqP64 VbslqP64
//go:noescape
func VbslqP64(r *arm.Poly64X2, v0 *arm.Uint64X2, v1 *arm.Poly64X2, v2 *arm.Poly64X2)

// Bitwise Select. This instruction sets each bit in the destination SIMD&FP register to the corresponding bit from the first source SIMD&FP register when the original destination bit was 1, otherwise from the second source SIMD&FP register.
//
//go:linkname VbslqF64 VbslqF64
//go:noescape
func VbslqF64(r *arm.Float64X2, v0 *arm.Uint64X2, v1 *arm.Float64X2, v2 *arm.Float64X2)

// Bitwise Select. This instruction sets each bit in the destination SIMD&FP register to the corresponding bit from the first source SIMD&FP register when the original destination bit was 1, otherwise from the second source SIMD&FP register.
//
//go:linkname VbslF64 VbslF64
//go:noescape
func VbslF64(r *arm.Float64X1, v0 *arm.Uint64X1, v1 *arm.Float64X1, v2 *arm.Float64X1)

// Floating-point Absolute Compare Greater than or Equal (vector). This instruction compares the absolute value of each floating-point value in the first source SIMD&FP register with the absolute value of the corresponding floating-point value in the second source SIMD&FP register and if the first value is greater than or equal to the second value sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VcageqF64 VcageqF64
//go:noescape
func VcageqF64(r *arm.Uint64X2, v0 *arm.Float64X2, v1 *arm.Float64X2)

// Floating-point Absolute Compare Greater than or Equal (vector). This instruction compares the absolute value of each floating-point value in the first source SIMD&FP register with the absolute value of the corresponding floating-point value in the second source SIMD&FP register and if the first value is greater than or equal to the second value sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VcageF64 VcageF64
//go:noescape
func VcageF64(r *arm.Uint64X1, v0 *arm.Float64X1, v1 *arm.Float64X1)

// Floating-point Absolute Compare Greater than or Equal (vector). This instruction compares the absolute value of each floating-point value in the first source SIMD&FP register with the absolute value of the corresponding floating-point value in the second source SIMD&FP register and if the first value is greater than or equal to the second value sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VcagedF64 VcagedF64
//go:noescape
func VcagedF64(r *arm.Uint64, v0 *arm.Float64, v1 *arm.Float64)

// Floating-point Absolute Compare Greater than or Equal (vector). This instruction compares the absolute value of each floating-point value in the first source SIMD&FP register with the absolute value of the corresponding floating-point value in the second source SIMD&FP register and if the first value is greater than or equal to the second value sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VcagesF32 VcagesF32
//go:noescape
func VcagesF32(r *arm.Uint32, v0 *arm.Float32, v1 *arm.Float32)

// Floating-point Absolute Compare Greater than (vector). This instruction compares the absolute value of each vector element in the first source SIMD&FP register with the absolute value of the corresponding vector element in the second source SIMD&FP register and if the first value is greater than the second value sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VcagtqF64 VcagtqF64
//go:noescape
func VcagtqF64(r *arm.Uint64X2, v0 *arm.Float64X2, v1 *arm.Float64X2)

// Floating-point Absolute Compare Greater than (vector). This instruction compares the absolute value of each vector element in the first source SIMD&FP register with the absolute value of the corresponding vector element in the second source SIMD&FP register and if the first value is greater than the second value sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VcagtF64 VcagtF64
//go:noescape
func VcagtF64(r *arm.Uint64X1, v0 *arm.Float64X1, v1 *arm.Float64X1)

// Floating-point Absolute Compare Greater than (vector). This instruction compares the absolute value of each vector element in the first source SIMD&FP register with the absolute value of the corresponding vector element in the second source SIMD&FP register and if the first value is greater than the second value sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VcagtdF64 VcagtdF64
//go:noescape
func VcagtdF64(r *arm.Uint64, v0 *arm.Float64, v1 *arm.Float64)

// Floating-point Absolute Compare Greater than (vector). This instruction compares the absolute value of each vector element in the first source SIMD&FP register with the absolute value of the corresponding vector element in the second source SIMD&FP register and if the first value is greater than the second value sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VcagtsF32 VcagtsF32
//go:noescape
func VcagtsF32(r *arm.Uint32, v0 *arm.Float32, v1 *arm.Float32)

// Floating-point absolute compare less than or equal
//
//go:linkname VcaleqF64 VcaleqF64
//go:noescape
func VcaleqF64(r *arm.Uint64X2, v0 *arm.Float64X2, v1 *arm.Float64X2)

// Floating-point absolute compare less than or equal
//
//go:linkname VcaleF64 VcaleF64
//go:noescape
func VcaleF64(r *arm.Uint64X1, v0 *arm.Float64X1, v1 *arm.Float64X1)

// Floating-point absolute compare less than or equal
//
//go:linkname VcaledF64 VcaledF64
//go:noescape
func VcaledF64(r *arm.Uint64, v0 *arm.Float64, v1 *arm.Float64)

// Floating-point absolute compare less than or equal
//
//go:linkname VcalesF32 VcalesF32
//go:noescape
func VcalesF32(r *arm.Uint32, v0 *arm.Float32, v1 *arm.Float32)

// Floating-point absolute compare less than
//
//go:linkname VcaltqF64 VcaltqF64
//go:noescape
func VcaltqF64(r *arm.Uint64X2, v0 *arm.Float64X2, v1 *arm.Float64X2)

// Floating-point absolute compare less than
//
//go:linkname VcaltF64 VcaltF64
//go:noescape
func VcaltF64(r *arm.Uint64X1, v0 *arm.Float64X1, v1 *arm.Float64X1)

// Floating-point absolute compare less than
//
//go:linkname VcaltdF64 VcaltdF64
//go:noescape
func VcaltdF64(r *arm.Uint64, v0 *arm.Float64, v1 *arm.Float64)

// Floating-point absolute compare less than
//
//go:linkname VcaltsF32 VcaltsF32
//go:noescape
func VcaltsF32(r *arm.Uint32, v0 *arm.Float32, v1 *arm.Float32)

// Compare bitwise Equal (vector). This instruction compares each vector element from the first source SIMD&FP register with the corresponding vector element from the second source SIMD&FP register, and if the comparison is equal sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VceqP64 VceqP64
//go:noescape
func VceqP64(r *arm.Uint64X1, v0 *arm.Poly64X1, v1 *arm.Poly64X1)

// Compare bitwise Equal (vector). This instruction compares each vector element from the first source SIMD&FP register with the corresponding vector element from the second source SIMD&FP register, and if the comparison is equal sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VceqqP64 VceqqP64
//go:noescape
func VceqqP64(r *arm.Uint64X2, v0 *arm.Poly64X2, v1 *arm.Poly64X2)

// Compare bitwise Equal (vector). This instruction compares each vector element from the first source SIMD&FP register with the corresponding vector element from the second source SIMD&FP register, and if the comparison is equal sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VceqqU64 VceqqU64
//go:noescape
func VceqqU64(r *arm.Uint64X2, v0 *arm.Uint64X2, v1 *arm.Uint64X2)

// Floating-point Compare Equal (vector). This instruction compares each floating-point value from the first source SIMD&FP register, with the corresponding floating-point value from the second source SIMD&FP register, and if the comparison is equal sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VceqqF64 VceqqF64
//go:noescape
func VceqqF64(r *arm.Uint64X2, v0 *arm.Float64X2, v1 *arm.Float64X2)

// Compare bitwise Equal (vector). This instruction compares each vector element from the first source SIMD&FP register with the corresponding vector element from the second source SIMD&FP register, and if the comparison is equal sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VceqqS64 VceqqS64
//go:noescape
func VceqqS64(r *arm.Uint64X2, v0 *arm.Int64X2, v1 *arm.Int64X2)

// Compare bitwise Equal (vector). This instruction compares each vector element from the first source SIMD&FP register with the corresponding vector element from the second source SIMD&FP register, and if the comparison is equal sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VceqU64 VceqU64
//go:noescape
func VceqU64(r *arm.Uint64X1, v0 *arm.Uint64X1, v1 *arm.Uint64X1)

// Floating-point Compare Equal (vector). This instruction compares each floating-point value from the first source SIMD&FP register, with the corresponding floating-point value from the second source SIMD&FP register, and if the comparison is equal sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VceqF64 VceqF64
//go:noescape
func VceqF64(r *arm.Uint64X1, v0 *arm.Float64X1, v1 *arm.Float64X1)

// Compare bitwise Equal (vector). This instruction compares each vector element from the first source SIMD&FP register with the corresponding vector element from the second source SIMD&FP register, and if the comparison is equal sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VceqS64 VceqS64
//go:noescape
func VceqS64(r *arm.Uint64X1, v0 *arm.Int64X1, v1 *arm.Int64X1)

// Compare bitwise Equal (vector). This instruction compares each vector element from the first source SIMD&FP register with the corresponding vector element from the second source SIMD&FP register, and if the comparison is equal sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VceqdU64 VceqdU64
//go:noescape
func VceqdU64(r *arm.Uint64, v0 *arm.Uint64, v1 *arm.Uint64)

// Compare bitwise Equal (vector). This instruction compares each vector element from the first source SIMD&FP register with the corresponding vector element from the second source SIMD&FP register, and if the comparison is equal sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VceqdS64 VceqdS64
//go:noescape
func VceqdS64(r *arm.Uint64, v0 *arm.Int64, v1 *arm.Int64)

// Floating-point Compare Equal (vector). This instruction compares each floating-point value from the first source SIMD&FP register, with the corresponding floating-point value from the second source SIMD&FP register, and if the comparison is equal sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VceqdF64 VceqdF64
//go:noescape
func VceqdF64(r *arm.Uint64, v0 *arm.Float64, v1 *arm.Float64)

// Floating-point Compare Equal (vector). This instruction compares each floating-point value from the first source SIMD&FP register, with the corresponding floating-point value from the second source SIMD&FP register, and if the comparison is equal sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VceqsF32 VceqsF32
//go:noescape
func VceqsF32(r *arm.Uint32, v0 *arm.Float32, v1 *arm.Float32)

// Compare bitwise Equal to zero (vector). This instruction reads each vector element in the source SIMD&FP register and if the value is equal to zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VceqzP8 VceqzP8
//go:noescape
func VceqzP8(r *arm.Uint8X8, v0 *arm.Poly8X8)

// Compare bitwise Equal to zero (vector). This instruction reads each vector element in the source SIMD&FP register and if the value is equal to zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VceqzP64 VceqzP64
//go:noescape
func VceqzP64(r *arm.Uint64X1, v0 *arm.Poly64X1)

// Compare bitwise Equal to zero (vector). This instruction reads each vector element in the source SIMD&FP register and if the value is equal to zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VceqzqP8 VceqzqP8
//go:noescape
func VceqzqP8(r *arm.Uint8X16, v0 *arm.Poly8X16)

// Compare bitwise Equal to zero (vector). This instruction reads each vector element in the source SIMD&FP register and if the value is equal to zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VceqzqP64 VceqzqP64
//go:noescape
func VceqzqP64(r *arm.Uint64X2, v0 *arm.Poly64X2)

// Compare bitwise Equal to zero (vector). This instruction reads each vector element in the source SIMD&FP register and if the value is equal to zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VceqzqU8 VceqzqU8
//go:noescape
func VceqzqU8(r *arm.Uint8X16, v0 *arm.Uint8X16)

// Compare bitwise Equal to zero (vector). This instruction reads each vector element in the source SIMD&FP register and if the value is equal to zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VceqzqU32 VceqzqU32
//go:noescape
func VceqzqU32(r *arm.Uint32X4, v0 *arm.Uint32X4)

// Compare bitwise Equal to zero (vector). This instruction reads each vector element in the source SIMD&FP register and if the value is equal to zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VceqzqU64 VceqzqU64
//go:noescape
func VceqzqU64(r *arm.Uint64X2, v0 *arm.Uint64X2)

// Compare bitwise Equal to zero (vector). This instruction reads each vector element in the source SIMD&FP register and if the value is equal to zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VceqzqU16 VceqzqU16
//go:noescape
func VceqzqU16(r *arm.Uint16X8, v0 *arm.Uint16X8)

// Compare bitwise Equal to zero (vector). This instruction reads each vector element in the source SIMD&FP register and if the value is equal to zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VceqzqS8 VceqzqS8
//go:noescape
func VceqzqS8(r *arm.Uint8X16, v0 *arm.Int8X16)

// Floating-point Compare Equal to zero (vector). This instruction reads each floating-point value in the source SIMD&FP register and if the value is equal to zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VceqzqF64 VceqzqF64
//go:noescape
func VceqzqF64(r *arm.Uint64X2, v0 *arm.Float64X2)

// Floating-point Compare Equal to zero (vector). This instruction reads each floating-point value in the source SIMD&FP register and if the value is equal to zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VceqzqF32 VceqzqF32
//go:noescape
func VceqzqF32(r *arm.Uint32X4, v0 *arm.Float32X4)

// Compare bitwise Equal to zero (vector). This instruction reads each vector element in the source SIMD&FP register and if the value is equal to zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VceqzqS32 VceqzqS32
//go:noescape
func VceqzqS32(r *arm.Uint32X4, v0 *arm.Int32X4)

// Compare bitwise Equal to zero (vector). This instruction reads each vector element in the source SIMD&FP register and if the value is equal to zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VceqzqS64 VceqzqS64
//go:noescape
func VceqzqS64(r *arm.Uint64X2, v0 *arm.Int64X2)

// Compare bitwise Equal to zero (vector). This instruction reads each vector element in the source SIMD&FP register and if the value is equal to zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VceqzqS16 VceqzqS16
//go:noescape
func VceqzqS16(r *arm.Uint16X8, v0 *arm.Int16X8)

// Compare bitwise Equal to zero (vector). This instruction reads each vector element in the source SIMD&FP register and if the value is equal to zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VceqzU8 VceqzU8
//go:noescape
func VceqzU8(r *arm.Uint8X8, v0 *arm.Uint8X8)

// Compare bitwise Equal to zero (vector). This instruction reads each vector element in the source SIMD&FP register and if the value is equal to zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VceqzU32 VceqzU32
//go:noescape
func VceqzU32(r *arm.Uint32X2, v0 *arm.Uint32X2)

// Compare bitwise Equal to zero (vector). This instruction reads each vector element in the source SIMD&FP register and if the value is equal to zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VceqzU64 VceqzU64
//go:noescape
func VceqzU64(r *arm.Uint64X1, v0 *arm.Uint64X1)

// Compare bitwise Equal to zero (vector). This instruction reads each vector element in the source SIMD&FP register and if the value is equal to zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VceqzU16 VceqzU16
//go:noescape
func VceqzU16(r *arm.Uint16X4, v0 *arm.Uint16X4)

// Compare bitwise Equal to zero (vector). This instruction reads each vector element in the source SIMD&FP register and if the value is equal to zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VceqzS8 VceqzS8
//go:noescape
func VceqzS8(r *arm.Uint8X8, v0 *arm.Int8X8)

// Floating-point Compare Equal to zero (vector). This instruction reads each floating-point value in the source SIMD&FP register and if the value is equal to zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VceqzF64 VceqzF64
//go:noescape
func VceqzF64(r *arm.Uint64X1, v0 *arm.Float64X1)

// Floating-point Compare Equal to zero (vector). This instruction reads each floating-point value in the source SIMD&FP register and if the value is equal to zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VceqzF32 VceqzF32
//go:noescape
func VceqzF32(r *arm.Uint32X2, v0 *arm.Float32X2)

// Compare bitwise Equal to zero (vector). This instruction reads each vector element in the source SIMD&FP register and if the value is equal to zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VceqzS32 VceqzS32
//go:noescape
func VceqzS32(r *arm.Uint32X2, v0 *arm.Int32X2)

// Compare bitwise Equal to zero (vector). This instruction reads each vector element in the source SIMD&FP register and if the value is equal to zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VceqzS64 VceqzS64
//go:noescape
func VceqzS64(r *arm.Uint64X1, v0 *arm.Int64X1)

// Compare bitwise Equal to zero (vector). This instruction reads each vector element in the source SIMD&FP register and if the value is equal to zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VceqzS16 VceqzS16
//go:noescape
func VceqzS16(r *arm.Uint16X4, v0 *arm.Int16X4)

// Compare bitwise Equal to zero (vector). This instruction reads each vector element in the source SIMD&FP register and if the value is equal to zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VceqzdU64 VceqzdU64
//go:noescape
func VceqzdU64(r *arm.Uint64, v0 *arm.Uint64)

// Compare bitwise Equal to zero (vector). This instruction reads each vector element in the source SIMD&FP register and if the value is equal to zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VceqzdS64 VceqzdS64
//go:noescape
func VceqzdS64(r *arm.Uint64, v0 *arm.Int64)

// Floating-point Compare Equal to zero (vector). This instruction reads each floating-point value in the source SIMD&FP register and if the value is equal to zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VceqzdF64 VceqzdF64
//go:noescape
func VceqzdF64(r *arm.Uint64, v0 *arm.Float64)

// Floating-point Compare Equal to zero (vector). This instruction reads each floating-point value in the source SIMD&FP register and if the value is equal to zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VceqzsF32 VceqzsF32
//go:noescape
func VceqzsF32(r *arm.Uint32, v0 *arm.Float32)

// Compare unsigned Higher or Same (vector). This instruction compares each vector element in the first source SIMD&FP register with the corresponding vector element in the second source SIMD&FP register and if the first unsigned integer value is greater than or equal to the second unsigned integer value sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VcgeqU64 VcgeqU64
//go:noescape
func VcgeqU64(r *arm.Uint64X2, v0 *arm.Uint64X2, v1 *arm.Uint64X2)

// Floating-point Compare Greater than or Equal (vector). This instruction reads each floating-point value in the first source SIMD&FP register and if the value is greater than or equal to the corresponding floating-point value in the second source SIMD&FP register sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VcgeqF64 VcgeqF64
//go:noescape
func VcgeqF64(r *arm.Uint64X2, v0 *arm.Float64X2, v1 *arm.Float64X2)

// Compare signed Greater than or Equal (vector). This instruction compares each vector element in the first source SIMD&FP register with the corresponding vector element in the second source SIMD&FP register and if the first signed integer value is greater than or equal to the second signed integer value sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VcgeqS64 VcgeqS64
//go:noescape
func VcgeqS64(r *arm.Uint64X2, v0 *arm.Int64X2, v1 *arm.Int64X2)

// Compare unsigned Higher or Same (vector). This instruction compares each vector element in the first source SIMD&FP register with the corresponding vector element in the second source SIMD&FP register and if the first unsigned integer value is greater than or equal to the second unsigned integer value sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VcgeU64 VcgeU64
//go:noescape
func VcgeU64(r *arm.Uint64X1, v0 *arm.Uint64X1, v1 *arm.Uint64X1)

// Floating-point Compare Greater than or Equal (vector). This instruction reads each floating-point value in the first source SIMD&FP register and if the value is greater than or equal to the corresponding floating-point value in the second source SIMD&FP register sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VcgeF64 VcgeF64
//go:noescape
func VcgeF64(r *arm.Uint64X1, v0 *arm.Float64X1, v1 *arm.Float64X1)

// Compare signed Greater than or Equal (vector). This instruction compares each vector element in the first source SIMD&FP register with the corresponding vector element in the second source SIMD&FP register and if the first signed integer value is greater than or equal to the second signed integer value sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VcgeS64 VcgeS64
//go:noescape
func VcgeS64(r *arm.Uint64X1, v0 *arm.Int64X1, v1 *arm.Int64X1)

// Compare signed Greater than or Equal (vector). This instruction compares each vector element in the first source SIMD&FP register with the corresponding vector element in the second source SIMD&FP register and if the first signed integer value is greater than or equal to the second signed integer value sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VcgedS64 VcgedS64
//go:noescape
func VcgedS64(r *arm.Uint64, v0 *arm.Int64, v1 *arm.Int64)

// Compare unsigned Higher or Same (vector). This instruction compares each vector element in the first source SIMD&FP register with the corresponding vector element in the second source SIMD&FP register and if the first unsigned integer value is greater than or equal to the second unsigned integer value sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VcgedU64 VcgedU64
//go:noescape
func VcgedU64(r *arm.Uint64, v0 *arm.Uint64, v1 *arm.Uint64)

// Floating-point Compare Greater than or Equal (vector). This instruction reads each floating-point value in the first source SIMD&FP register and if the value is greater than or equal to the corresponding floating-point value in the second source SIMD&FP register sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VcgedF64 VcgedF64
//go:noescape
func VcgedF64(r *arm.Uint64, v0 *arm.Float64, v1 *arm.Float64)

// Floating-point Compare Greater than or Equal (vector). This instruction reads each floating-point value in the first source SIMD&FP register and if the value is greater than or equal to the corresponding floating-point value in the second source SIMD&FP register sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VcgesF32 VcgesF32
//go:noescape
func VcgesF32(r *arm.Uint32, v0 *arm.Float32, v1 *arm.Float32)

// Compare signed Greater than or Equal to zero (vector). This instruction reads each vector element in the source SIMD&FP register and if the signed integer value is greater than or equal to zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VcgezqS8 VcgezqS8
//go:noescape
func VcgezqS8(r *arm.Uint8X16, v0 *arm.Int8X16)

// Floating-point Compare Greater than or Equal to zero (vector). This instruction reads each floating-point value in the source SIMD&FP register and if the value is greater than or equal to zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VcgezqF64 VcgezqF64
//go:noescape
func VcgezqF64(r *arm.Uint64X2, v0 *arm.Float64X2)

// Floating-point Compare Greater than or Equal to zero (vector). This instruction reads each floating-point value in the source SIMD&FP register and if the value is greater than or equal to zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VcgezqF32 VcgezqF32
//go:noescape
func VcgezqF32(r *arm.Uint32X4, v0 *arm.Float32X4)

// Compare signed Greater than or Equal to zero (vector). This instruction reads each vector element in the source SIMD&FP register and if the signed integer value is greater than or equal to zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VcgezqS32 VcgezqS32
//go:noescape
func VcgezqS32(r *arm.Uint32X4, v0 *arm.Int32X4)

// Compare signed Greater than or Equal to zero (vector). This instruction reads each vector element in the source SIMD&FP register and if the signed integer value is greater than or equal to zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VcgezqS64 VcgezqS64
//go:noescape
func VcgezqS64(r *arm.Uint64X2, v0 *arm.Int64X2)

// Compare signed Greater than or Equal to zero (vector). This instruction reads each vector element in the source SIMD&FP register and if the signed integer value is greater than or equal to zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VcgezqS16 VcgezqS16
//go:noescape
func VcgezqS16(r *arm.Uint16X8, v0 *arm.Int16X8)

// Compare signed Greater than or Equal to zero (vector). This instruction reads each vector element in the source SIMD&FP register and if the signed integer value is greater than or equal to zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VcgezS8 VcgezS8
//go:noescape
func VcgezS8(r *arm.Uint8X8, v0 *arm.Int8X8)

// Floating-point Compare Greater than or Equal to zero (vector). This instruction reads each floating-point value in the source SIMD&FP register and if the value is greater than or equal to zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VcgezF64 VcgezF64
//go:noescape
func VcgezF64(r *arm.Uint64X1, v0 *arm.Float64X1)

// Floating-point Compare Greater than or Equal to zero (vector). This instruction reads each floating-point value in the source SIMD&FP register and if the value is greater than or equal to zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VcgezF32 VcgezF32
//go:noescape
func VcgezF32(r *arm.Uint32X2, v0 *arm.Float32X2)

// Compare signed Greater than or Equal to zero (vector). This instruction reads each vector element in the source SIMD&FP register and if the signed integer value is greater than or equal to zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VcgezS32 VcgezS32
//go:noescape
func VcgezS32(r *arm.Uint32X2, v0 *arm.Int32X2)

// Compare signed Greater than or Equal to zero (vector). This instruction reads each vector element in the source SIMD&FP register and if the signed integer value is greater than or equal to zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VcgezS64 VcgezS64
//go:noescape
func VcgezS64(r *arm.Uint64X1, v0 *arm.Int64X1)

// Compare signed Greater than or Equal to zero (vector). This instruction reads each vector element in the source SIMD&FP register and if the signed integer value is greater than or equal to zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VcgezS16 VcgezS16
//go:noescape
func VcgezS16(r *arm.Uint16X4, v0 *arm.Int16X4)

// Compare signed Greater than or Equal to zero (vector). This instruction reads each vector element in the source SIMD&FP register and if the signed integer value is greater than or equal to zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VcgezdS64 VcgezdS64
//go:noescape
func VcgezdS64(r *arm.Uint64, v0 *arm.Int64)

// Floating-point Compare Greater than or Equal to zero (vector). This instruction reads each floating-point value in the source SIMD&FP register and if the value is greater than or equal to zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VcgezdF64 VcgezdF64
//go:noescape
func VcgezdF64(r *arm.Uint64, v0 *arm.Float64)

// Floating-point Compare Greater than or Equal to zero (vector). This instruction reads each floating-point value in the source SIMD&FP register and if the value is greater than or equal to zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VcgezsF32 VcgezsF32
//go:noescape
func VcgezsF32(r *arm.Uint32, v0 *arm.Float32)

// Compare unsigned Higher (vector). This instruction compares each vector element in the first source SIMD&FP register with the corresponding vector element in the second source SIMD&FP register and if the first unsigned integer value is greater than the second unsigned integer value sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VcgtqU64 VcgtqU64
//go:noescape
func VcgtqU64(r *arm.Uint64X2, v0 *arm.Uint64X2, v1 *arm.Uint64X2)

// Floating-point Compare Greater than (vector). This instruction reads each floating-point value in the first source SIMD&FP register and if the value is greater than the corresponding floating-point value in the second source SIMD&FP register sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VcgtqF64 VcgtqF64
//go:noescape
func VcgtqF64(r *arm.Uint64X2, v0 *arm.Float64X2, v1 *arm.Float64X2)

// Compare signed Greater than (vector). This instruction compares each vector element in the first source SIMD&FP register with the corresponding vector element in the second source SIMD&FP register and if the first signed integer value is greater than the second signed integer value sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VcgtqS64 VcgtqS64
//go:noescape
func VcgtqS64(r *arm.Uint64X2, v0 *arm.Int64X2, v1 *arm.Int64X2)

// Compare unsigned Higher (vector). This instruction compares each vector element in the first source SIMD&FP register with the corresponding vector element in the second source SIMD&FP register and if the first unsigned integer value is greater than the second unsigned integer value sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VcgtU64 VcgtU64
//go:noescape
func VcgtU64(r *arm.Uint64X1, v0 *arm.Uint64X1, v1 *arm.Uint64X1)

// Floating-point Compare Greater than (vector). This instruction reads each floating-point value in the first source SIMD&FP register and if the value is greater than the corresponding floating-point value in the second source SIMD&FP register sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VcgtF64 VcgtF64
//go:noescape
func VcgtF64(r *arm.Uint64X1, v0 *arm.Float64X1, v1 *arm.Float64X1)

// Compare signed Greater than (vector). This instruction compares each vector element in the first source SIMD&FP register with the corresponding vector element in the second source SIMD&FP register and if the first signed integer value is greater than the second signed integer value sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VcgtS64 VcgtS64
//go:noescape
func VcgtS64(r *arm.Uint64X1, v0 *arm.Int64X1, v1 *arm.Int64X1)

// Compare signed Greater than (vector). This instruction compares each vector element in the first source SIMD&FP register with the corresponding vector element in the second source SIMD&FP register and if the first signed integer value is greater than the second signed integer value sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VcgtdS64 VcgtdS64
//go:noescape
func VcgtdS64(r *arm.Uint64, v0 *arm.Int64, v1 *arm.Int64)

// Compare unsigned Higher (vector). This instruction compares each vector element in the first source SIMD&FP register with the corresponding vector element in the second source SIMD&FP register and if the first unsigned integer value is greater than the second unsigned integer value sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VcgtdU64 VcgtdU64
//go:noescape
func VcgtdU64(r *arm.Uint64, v0 *arm.Uint64, v1 *arm.Uint64)

// Floating-point Compare Greater than (vector). This instruction reads each floating-point value in the first source SIMD&FP register and if the value is greater than the corresponding floating-point value in the second source SIMD&FP register sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VcgtdF64 VcgtdF64
//go:noescape
func VcgtdF64(r *arm.Uint64, v0 *arm.Float64, v1 *arm.Float64)

// Floating-point Compare Greater than (vector). This instruction reads each floating-point value in the first source SIMD&FP register and if the value is greater than the corresponding floating-point value in the second source SIMD&FP register sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VcgtsF32 VcgtsF32
//go:noescape
func VcgtsF32(r *arm.Uint32, v0 *arm.Float32, v1 *arm.Float32)

// Compare signed Greater than zero (vector). This instruction reads each vector element in the source SIMD&FP register and if the signed integer value is greater than zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VcgtzqS8 VcgtzqS8
//go:noescape
func VcgtzqS8(r *arm.Uint8X16, v0 *arm.Int8X16)

// Floating-point Compare Greater than zero (vector). This instruction reads each floating-point value in the source SIMD&FP register and if the value is greater than zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VcgtzqF64 VcgtzqF64
//go:noescape
func VcgtzqF64(r *arm.Uint64X2, v0 *arm.Float64X2)

// Floating-point Compare Greater than zero (vector). This instruction reads each floating-point value in the source SIMD&FP register and if the value is greater than zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VcgtzqF32 VcgtzqF32
//go:noescape
func VcgtzqF32(r *arm.Uint32X4, v0 *arm.Float32X4)

// Compare signed Greater than zero (vector). This instruction reads each vector element in the source SIMD&FP register and if the signed integer value is greater than zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VcgtzqS32 VcgtzqS32
//go:noescape
func VcgtzqS32(r *arm.Uint32X4, v0 *arm.Int32X4)

// Compare signed Greater than zero (vector). This instruction reads each vector element in the source SIMD&FP register and if the signed integer value is greater than zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VcgtzqS64 VcgtzqS64
//go:noescape
func VcgtzqS64(r *arm.Uint64X2, v0 *arm.Int64X2)

// Compare signed Greater than zero (vector). This instruction reads each vector element in the source SIMD&FP register and if the signed integer value is greater than zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VcgtzqS16 VcgtzqS16
//go:noescape
func VcgtzqS16(r *arm.Uint16X8, v0 *arm.Int16X8)

// Compare signed Greater than zero (vector). This instruction reads each vector element in the source SIMD&FP register and if the signed integer value is greater than zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VcgtzS8 VcgtzS8
//go:noescape
func VcgtzS8(r *arm.Uint8X8, v0 *arm.Int8X8)

// Floating-point Compare Greater than zero (vector). This instruction reads each floating-point value in the source SIMD&FP register and if the value is greater than zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VcgtzF64 VcgtzF64
//go:noescape
func VcgtzF64(r *arm.Uint64X1, v0 *arm.Float64X1)

// Floating-point Compare Greater than zero (vector). This instruction reads each floating-point value in the source SIMD&FP register and if the value is greater than zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VcgtzF32 VcgtzF32
//go:noescape
func VcgtzF32(r *arm.Uint32X2, v0 *arm.Float32X2)

// Compare signed Greater than zero (vector). This instruction reads each vector element in the source SIMD&FP register and if the signed integer value is greater than zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VcgtzS32 VcgtzS32
//go:noescape
func VcgtzS32(r *arm.Uint32X2, v0 *arm.Int32X2)

// Compare signed Greater than zero (vector). This instruction reads each vector element in the source SIMD&FP register and if the signed integer value is greater than zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VcgtzS64 VcgtzS64
//go:noescape
func VcgtzS64(r *arm.Uint64X1, v0 *arm.Int64X1)

// Compare signed Greater than zero (vector). This instruction reads each vector element in the source SIMD&FP register and if the signed integer value is greater than zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VcgtzS16 VcgtzS16
//go:noescape
func VcgtzS16(r *arm.Uint16X4, v0 *arm.Int16X4)

// Compare signed Greater than zero (vector). This instruction reads each vector element in the source SIMD&FP register and if the signed integer value is greater than zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VcgtzdS64 VcgtzdS64
//go:noescape
func VcgtzdS64(r *arm.Uint64, v0 *arm.Int64)

// Floating-point Compare Greater than zero (vector). This instruction reads each floating-point value in the source SIMD&FP register and if the value is greater than zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VcgtzdF64 VcgtzdF64
//go:noescape
func VcgtzdF64(r *arm.Uint64, v0 *arm.Float64)

// Floating-point Compare Greater than zero (vector). This instruction reads each floating-point value in the source SIMD&FP register and if the value is greater than zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VcgtzsF32 VcgtzsF32
//go:noescape
func VcgtzsF32(r *arm.Uint32, v0 *arm.Float32)

// Compare unsigned less than or equal
//
//go:linkname VcleqU64 VcleqU64
//go:noescape
func VcleqU64(r *arm.Uint64X2, v0 *arm.Uint64X2, v1 *arm.Uint64X2)

// Floating-point compare less than or equal
//
//go:linkname VcleqF64 VcleqF64
//go:noescape
func VcleqF64(r *arm.Uint64X2, v0 *arm.Float64X2, v1 *arm.Float64X2)

// Compare signed less than or equal
//
//go:linkname VcleqS64 VcleqS64
//go:noescape
func VcleqS64(r *arm.Uint64X2, v0 *arm.Int64X2, v1 *arm.Int64X2)

// Compare unsigned less than or equal
//
//go:linkname VcleU64 VcleU64
//go:noescape
func VcleU64(r *arm.Uint64X1, v0 *arm.Uint64X1, v1 *arm.Uint64X1)

// Floating-point compare less than or equal
//
//go:linkname VcleF64 VcleF64
//go:noescape
func VcleF64(r *arm.Uint64X1, v0 *arm.Float64X1, v1 *arm.Float64X1)

// Compare signed less than or equal
//
//go:linkname VcleS64 VcleS64
//go:noescape
func VcleS64(r *arm.Uint64X1, v0 *arm.Int64X1, v1 *arm.Int64X1)

// Compare unsigned less than or equal
//
//go:linkname VcledU64 VcledU64
//go:noescape
func VcledU64(r *arm.Uint64, v0 *arm.Uint64, v1 *arm.Uint64)

// Compare signed less than or equal
//
//go:linkname VcledS64 VcledS64
//go:noescape
func VcledS64(r *arm.Uint64, v0 *arm.Int64, v1 *arm.Int64)

// Floating-point compare less than or equal
//
//go:linkname VcledF64 VcledF64
//go:noescape
func VcledF64(r *arm.Uint64, v0 *arm.Float64, v1 *arm.Float64)

// Floating-point compare less than or equal
//
//go:linkname VclesF32 VclesF32
//go:noescape
func VclesF32(r *arm.Uint32, v0 *arm.Float32, v1 *arm.Float32)

// Compare signed Less than or Equal to zero (vector). This instruction reads each vector element in the source SIMD&FP register and if the signed integer value is less than or equal to zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VclezqS8 VclezqS8
//go:noescape
func VclezqS8(r *arm.Uint8X16, v0 *arm.Int8X16)

// Floating-point Compare Less than or Equal to zero (vector). This instruction reads each floating-point value in the source SIMD&FP register and if the value is less than or equal to zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VclezqF64 VclezqF64
//go:noescape
func VclezqF64(r *arm.Uint64X2, v0 *arm.Float64X2)

// Floating-point Compare Less than or Equal to zero (vector). This instruction reads each floating-point value in the source SIMD&FP register and if the value is less than or equal to zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VclezqF32 VclezqF32
//go:noescape
func VclezqF32(r *arm.Uint32X4, v0 *arm.Float32X4)

// Compare signed Less than or Equal to zero (vector). This instruction reads each vector element in the source SIMD&FP register and if the signed integer value is less than or equal to zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VclezqS32 VclezqS32
//go:noescape
func VclezqS32(r *arm.Uint32X4, v0 *arm.Int32X4)

// Compare signed Less than or Equal to zero (vector). This instruction reads each vector element in the source SIMD&FP register and if the signed integer value is less than or equal to zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VclezqS64 VclezqS64
//go:noescape
func VclezqS64(r *arm.Uint64X2, v0 *arm.Int64X2)

// Compare signed Less than or Equal to zero (vector). This instruction reads each vector element in the source SIMD&FP register and if the signed integer value is less than or equal to zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VclezqS16 VclezqS16
//go:noescape
func VclezqS16(r *arm.Uint16X8, v0 *arm.Int16X8)

// Compare signed Less than or Equal to zero (vector). This instruction reads each vector element in the source SIMD&FP register and if the signed integer value is less than or equal to zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VclezS8 VclezS8
//go:noescape
func VclezS8(r *arm.Uint8X8, v0 *arm.Int8X8)

// Floating-point Compare Less than or Equal to zero (vector). This instruction reads each floating-point value in the source SIMD&FP register and if the value is less than or equal to zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VclezF64 VclezF64
//go:noescape
func VclezF64(r *arm.Uint64X1, v0 *arm.Float64X1)

// Compare signed Less than or Equal to zero (vector). This instruction reads each vector element in the source SIMD&FP register and if the signed integer value is less than or equal to zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VclezF32 VclezF32
//go:noescape
func VclezF32(r *arm.Uint32X2, v0 *arm.Float32X2)

// Compare signed Less than or Equal to zero (vector). This instruction reads each vector element in the source SIMD&FP register and if the signed integer value is less than or equal to zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VclezS32 VclezS32
//go:noescape
func VclezS32(r *arm.Uint32X2, v0 *arm.Int32X2)

// Compare signed Less than or Equal to zero (vector). This instruction reads each vector element in the source SIMD&FP register and if the signed integer value is less than or equal to zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VclezS64 VclezS64
//go:noescape
func VclezS64(r *arm.Uint64X1, v0 *arm.Int64X1)

// Compare signed Less than or Equal to zero (vector). This instruction reads each vector element in the source SIMD&FP register and if the signed integer value is less than or equal to zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VclezS16 VclezS16
//go:noescape
func VclezS16(r *arm.Uint16X4, v0 *arm.Int16X4)

// Compare signed Less than or Equal to zero (vector). This instruction reads each vector element in the source SIMD&FP register and if the signed integer value is less than or equal to zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VclezdS64 VclezdS64
//go:noescape
func VclezdS64(r *arm.Uint64, v0 *arm.Int64)

// Floating-point Compare Less than or Equal to zero (vector). This instruction reads each floating-point value in the source SIMD&FP register and if the value is less than or equal to zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VclezdF64 VclezdF64
//go:noescape
func VclezdF64(r *arm.Uint64, v0 *arm.Float64)

// Floating-point Compare Less than or Equal to zero (vector). This instruction reads each floating-point value in the source SIMD&FP register and if the value is less than or equal to zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VclezsF32 VclezsF32
//go:noescape
func VclezsF32(r *arm.Uint32, v0 *arm.Float32)

// Compare unsigned less than
//
//go:linkname VcltqU64 VcltqU64
//go:noescape
func VcltqU64(r *arm.Uint64X2, v0 *arm.Uint64X2, v1 *arm.Uint64X2)

// Floating-point compare less than
//
//go:linkname VcltqF64 VcltqF64
//go:noescape
func VcltqF64(r *arm.Uint64X2, v0 *arm.Float64X2, v1 *arm.Float64X2)

// Compare signed less than
//
//go:linkname VcltqS64 VcltqS64
//go:noescape
func VcltqS64(r *arm.Uint64X2, v0 *arm.Int64X2, v1 *arm.Int64X2)

// Compare unsigned less than
//
//go:linkname VcltU64 VcltU64
//go:noescape
func VcltU64(r *arm.Uint64X1, v0 *arm.Uint64X1, v1 *arm.Uint64X1)

// Floating-point compare less than
//
//go:linkname VcltF64 VcltF64
//go:noescape
func VcltF64(r *arm.Uint64X1, v0 *arm.Float64X1, v1 *arm.Float64X1)

// Compare signed less than
//
//go:linkname VcltS64 VcltS64
//go:noescape
func VcltS64(r *arm.Uint64X1, v0 *arm.Int64X1, v1 *arm.Int64X1)

// Compare unsigned less than
//
//go:linkname VcltdU64 VcltdU64
//go:noescape
func VcltdU64(r *arm.Uint64, v0 *arm.Uint64, v1 *arm.Uint64)

// Compare signed less than
//
//go:linkname VcltdS64 VcltdS64
//go:noescape
func VcltdS64(r *arm.Uint64, v0 *arm.Int64, v1 *arm.Int64)

// Floating-point compare less than
//
//go:linkname VcltdF64 VcltdF64
//go:noescape
func VcltdF64(r *arm.Uint64, v0 *arm.Float64, v1 *arm.Float64)

// Floating-point compare less than
//
//go:linkname VcltsF32 VcltsF32
//go:noescape
func VcltsF32(r *arm.Uint32, v0 *arm.Float32, v1 *arm.Float32)

// Compare signed Less than zero (vector). This instruction reads each vector element in the source SIMD&FP register and if the signed integer value is less than zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VcltzqS8 VcltzqS8
//go:noescape
func VcltzqS8(r *arm.Uint8X16, v0 *arm.Int8X16)

// Floating-point Compare Less than zero (vector). This instruction reads each floating-point value in the source SIMD&FP register and if the value is less than zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VcltzqF64 VcltzqF64
//go:noescape
func VcltzqF64(r *arm.Uint64X2, v0 *arm.Float64X2)

// Floating-point Compare Less than zero (vector). This instruction reads each floating-point value in the source SIMD&FP register and if the value is less than zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VcltzqF32 VcltzqF32
//go:noescape
func VcltzqF32(r *arm.Uint32X4, v0 *arm.Float32X4)

// Compare signed Less than zero (vector). This instruction reads each vector element in the source SIMD&FP register and if the signed integer value is less than zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VcltzqS32 VcltzqS32
//go:noescape
func VcltzqS32(r *arm.Uint32X4, v0 *arm.Int32X4)

// Compare signed Less than zero (vector). This instruction reads each vector element in the source SIMD&FP register and if the signed integer value is less than zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VcltzqS64 VcltzqS64
//go:noescape
func VcltzqS64(r *arm.Uint64X2, v0 *arm.Int64X2)

// Compare signed Less than zero (vector). This instruction reads each vector element in the source SIMD&FP register and if the signed integer value is less than zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VcltzqS16 VcltzqS16
//go:noescape
func VcltzqS16(r *arm.Uint16X8, v0 *arm.Int16X8)

// Compare signed Less than zero (vector). This instruction reads each vector element in the source SIMD&FP register and if the signed integer value is less than zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VcltzS8 VcltzS8
//go:noescape
func VcltzS8(r *arm.Uint8X8, v0 *arm.Int8X8)

// Floating-point Compare Less than zero (vector). This instruction reads each floating-point value in the source SIMD&FP register and if the value is less than zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VcltzF64 VcltzF64
//go:noescape
func VcltzF64(r *arm.Uint64X1, v0 *arm.Float64X1)

// Floating-point Compare Less than zero (vector). This instruction reads each floating-point value in the source SIMD&FP register and if the value is less than zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VcltzF32 VcltzF32
//go:noescape
func VcltzF32(r *arm.Uint32X2, v0 *arm.Float32X2)

// Compare signed Less than zero (vector). This instruction reads each vector element in the source SIMD&FP register and if the signed integer value is less than zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VcltzS32 VcltzS32
//go:noescape
func VcltzS32(r *arm.Uint32X2, v0 *arm.Int32X2)

// Compare signed Less than zero (vector). This instruction reads each vector element in the source SIMD&FP register and if the signed integer value is less than zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VcltzS64 VcltzS64
//go:noescape
func VcltzS64(r *arm.Uint64X1, v0 *arm.Int64X1)

// Compare signed Less than zero (vector). This instruction reads each vector element in the source SIMD&FP register and if the signed integer value is less than zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VcltzS16 VcltzS16
//go:noescape
func VcltzS16(r *arm.Uint16X4, v0 *arm.Int16X4)

// Compare signed Less than zero (vector). This instruction reads each vector element in the source SIMD&FP register and if the signed integer value is less than zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VcltzdS64 VcltzdS64
//go:noescape
func VcltzdS64(r *arm.Uint64, v0 *arm.Int64)

// Floating-point Compare Less than zero (vector). This instruction reads each floating-point value in the source SIMD&FP register and if the value is less than zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VcltzdF64 VcltzdF64
//go:noescape
func VcltzdF64(r *arm.Uint64, v0 *arm.Float64)

// Floating-point Compare Less than zero (vector). This instruction reads each floating-point value in the source SIMD&FP register and if the value is less than zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VcltzsF32 VcltzsF32
//go:noescape
func VcltzsF32(r *arm.Uint32, v0 *arm.Float32)

// Join two smaller vectors into a single larger vector
//
//go:linkname VcombineP64 VcombineP64
//go:noescape
func VcombineP64(r *arm.Poly64X2, v0 *arm.Poly64X1, v1 *arm.Poly64X1)

// Join two smaller vectors into a single larger vector
//
//go:linkname VcombineF64 VcombineF64
//go:noescape
func VcombineF64(r *arm.Float64X2, v0 *arm.Float64X1, v1 *arm.Float64X1)

// Signed fixed-point Convert to Floating-point (vector). This instruction converts each element in a vector from fixed-point to floating-point using the rounding mode that is specified by the FPCR, and writes the result to the SIMD&FP destination register.
//
//go:linkname VcvtsF32S32 VcvtsF32S32
//go:noescape
func VcvtsF32S32(r *arm.Float32, v0 *arm.Int32)

// Unsigned fixed-point Convert to Floating-point (vector). This instruction converts each element in a vector from fixed-point to floating-point using the rounding mode that is specified by the FPCR, and writes the result to the SIMD&FP destination register.
//
//go:linkname VcvtsF32U32 VcvtsF32U32
//go:noescape
func VcvtsF32U32(r *arm.Float32, v0 *arm.Uint32)

// Floating-point Convert to lower precision Narrow (vector). This instruction reads each vector element in the SIMD&FP source register, converts each result to half the precision of the source element, writes the final result to a vector, and writes the vector to the lower or upper half of the destination SIMD&FP register. The destination vector elements are half as long as the source vector elements. The rounding mode is determined by the FPCR.
//
//go:linkname VcvtF32F64 VcvtF32F64
//go:noescape
func VcvtF32F64(r *arm.Float32X2, v0 *arm.Float64X2)

// Signed fixed-point Convert to Floating-point (vector). This instruction converts each element in a vector from fixed-point to floating-point using the rounding mode that is specified by the FPCR, and writes the result to the SIMD&FP destination register.
//
//go:linkname VcvtdF64S64 VcvtdF64S64
//go:noescape
func VcvtdF64S64(r *arm.Float64, v0 *arm.Int64)

// Unsigned fixed-point Convert to Floating-point (vector). This instruction converts each element in a vector from fixed-point to floating-point using the rounding mode that is specified by the FPCR, and writes the result to the SIMD&FP destination register.
//
//go:linkname VcvtdF64U64 VcvtdF64U64
//go:noescape
func VcvtdF64U64(r *arm.Float64, v0 *arm.Uint64)

// Unsigned fixed-point Convert to Floating-point (vector). This instruction converts each element in a vector from fixed-point to floating-point using the rounding mode that is specified by the FPCR, and writes the result to the SIMD&FP destination register.
//
//go:linkname VcvtqF64U64 VcvtqF64U64
//go:noescape
func VcvtqF64U64(r *arm.Float64X2, v0 *arm.Uint64X2)

// Signed fixed-point Convert to Floating-point (vector). This instruction converts each element in a vector from fixed-point to floating-point using the rounding mode that is specified by the FPCR, and writes the result to the SIMD&FP destination register.
//
//go:linkname VcvtqF64S64 VcvtqF64S64
//go:noescape
func VcvtqF64S64(r *arm.Float64X2, v0 *arm.Int64X2)

// Unsigned fixed-point Convert to Floating-point (vector). This instruction converts each element in a vector from fixed-point to floating-point using the rounding mode that is specified by the FPCR, and writes the result to the SIMD&FP destination register.
//
//go:linkname VcvtF64U64 VcvtF64U64
//go:noescape
func VcvtF64U64(r *arm.Float64X1, v0 *arm.Uint64X1)

// Signed fixed-point Convert to Floating-point (vector). This instruction converts each element in a vector from fixed-point to floating-point using the rounding mode that is specified by the FPCR, and writes the result to the SIMD&FP destination register.
//
//go:linkname VcvtF64S64 VcvtF64S64
//go:noescape
func VcvtF64S64(r *arm.Float64X1, v0 *arm.Int64X1)

// Floating-point Convert to higher precision Long (vector). This instruction reads each element in a vector in the SIMD&FP source register, converts each value to double the precision of the source element using the rounding mode that is determined by the FPCR, and writes each result to the equivalent element of the vector in the SIMD&FP destination register.
//
//go:linkname VcvtF64F32 VcvtF64F32
//go:noescape
func VcvtF64F32(r *arm.Float64X2, v0 *arm.Float32X2)

// Floating-point Convert to lower precision Narrow (vector). This instruction reads each vector element in the SIMD&FP source register, converts each result to half the precision of the source element, writes the final result to a vector, and writes the vector to the lower or upper half of the destination SIMD&FP register. The destination vector elements are half as long as the source vector elements. The rounding mode is determined by the FPCR.
//
//go:linkname VcvtHighF32F64 VcvtHighF32F64
//go:noescape
func VcvtHighF32F64(r *arm.Float32X4, v0 *arm.Float32X2, v1 *arm.Float64X2)

// Floating-point Convert to higher precision Long (vector). This instruction reads each element in a vector in the SIMD&FP source register, converts each value to double the precision of the source element using the rounding mode that is determined by the FPCR, and writes each result to the equivalent element of the vector in the SIMD&FP destination register.
//
//go:linkname VcvtHighF64F32 VcvtHighF64F32
//go:noescape
func VcvtHighF64F32(r *arm.Float64X2, v0 *arm.Float32X4)

// Floating-point Convert to Signed fixed-point, rounding toward Zero (vector). This instruction converts a scalar or each element in a vector from floating-point to fixed-point signed integer using the Round towards Zero rounding mode, and writes the result to the SIMD&FP destination register.
//
//go:linkname VcvtsS32F32 VcvtsS32F32
//go:noescape
func VcvtsS32F32(r *arm.Int32, v0 *arm.Float32)

// Floating-point Convert to Signed fixed-point, rounding toward Zero (vector). This instruction converts a scalar or each element in a vector from floating-point to fixed-point signed integer using the Round towards Zero rounding mode, and writes the result to the SIMD&FP destination register.
//
//go:linkname VcvtdS64F64 VcvtdS64F64
//go:noescape
func VcvtdS64F64(r *arm.Int64, v0 *arm.Float64)

// Floating-point Convert to Signed fixed-point, rounding toward Zero (vector). This instruction converts a scalar or each element in a vector from floating-point to fixed-point signed integer using the Round towards Zero rounding mode, and writes the result to the SIMD&FP destination register.
//
//go:linkname VcvtqS64F64 VcvtqS64F64
//go:noescape
func VcvtqS64F64(r *arm.Int64X2, v0 *arm.Float64X2)

// Floating-point Convert to Signed fixed-point, rounding toward Zero (vector). This instruction converts a scalar or each element in a vector from floating-point to fixed-point signed integer using the Round towards Zero rounding mode, and writes the result to the SIMD&FP destination register.
//
//go:linkname VcvtS64F64 VcvtS64F64
//go:noescape
func VcvtS64F64(r *arm.Int64X1, v0 *arm.Float64X1)

// Floating-point Convert to Unsigned fixed-point, rounding toward Zero (vector). This instruction converts a scalar or each element in a vector from floating-point to fixed-point unsigned integer using the Round towards Zero rounding mode, and writes the result to the general-purpose destination register.
//
//go:linkname VcvtsU32F32 VcvtsU32F32
//go:noescape
func VcvtsU32F32(r *arm.Uint32, v0 *arm.Float32)

// Floating-point Convert to Unsigned fixed-point, rounding toward Zero (vector). This instruction converts a scalar or each element in a vector from floating-point to fixed-point unsigned integer using the Round towards Zero rounding mode, and writes the result to the general-purpose destination register.
//
//go:linkname VcvtdU64F64 VcvtdU64F64
//go:noescape
func VcvtdU64F64(r *arm.Uint64, v0 *arm.Float64)

// Floating-point Convert to Unsigned fixed-point, rounding toward Zero (vector). This instruction converts a scalar or each element in a vector from floating-point to fixed-point unsigned integer using the Round towards Zero rounding mode, and writes the result to the general-purpose destination register.
//
//go:linkname VcvtqU64F64 VcvtqU64F64
//go:noescape
func VcvtqU64F64(r *arm.Uint64X2, v0 *arm.Float64X2)

// Floating-point Convert to Unsigned fixed-point, rounding toward Zero (vector). This instruction converts a scalar or each element in a vector from floating-point to fixed-point unsigned integer using the Round towards Zero rounding mode, and writes the result to the general-purpose destination register.
//
//go:linkname VcvtU64F64 VcvtU64F64
//go:noescape
func VcvtU64F64(r *arm.Uint64X1, v0 *arm.Float64X1)

// Floating-point Convert to Signed integer, rounding to nearest with ties to Away (vector). This instruction converts each element in a vector from a floating-point value to a signed integer value using the Round to Nearest with Ties to Away rounding mode and writes the result to the SIMD&FP destination register.
//
//go:linkname VcvtasS32F32 VcvtasS32F32
//go:noescape
func VcvtasS32F32(r *arm.Int32, v0 *arm.Float32)

// Floating-point Convert to Signed integer, rounding to nearest with ties to Away (vector). This instruction converts each element in a vector from a floating-point value to a signed integer value using the Round to Nearest with Ties to Away rounding mode and writes the result to the SIMD&FP destination register.
//
//go:linkname VcvtadS64F64 VcvtadS64F64
//go:noescape
func VcvtadS64F64(r *arm.Int64, v0 *arm.Float64)

// Floating-point Convert to Unsigned integer, rounding to nearest with ties to Away (vector). This instruction converts each element in a vector from a floating-point value to an unsigned integer value using the Round to Nearest with Ties to Away rounding mode and writes the result to the SIMD&FP destination register.
//
//go:linkname VcvtasU32F32 VcvtasU32F32
//go:noescape
func VcvtasU32F32(r *arm.Uint32, v0 *arm.Float32)

// Floating-point Convert to Unsigned integer, rounding to nearest with ties to Away (vector). This instruction converts each element in a vector from a floating-point value to an unsigned integer value using the Round to Nearest with Ties to Away rounding mode and writes the result to the SIMD&FP destination register.
//
//go:linkname VcvtadU64F64 VcvtadU64F64
//go:noescape
func VcvtadU64F64(r *arm.Uint64, v0 *arm.Float64)

// Floating-point Convert to Signed integer, rounding toward Minus infinity (vector). This instruction converts a scalar or each element in a vector from a floating-point value to a signed integer value using the Round towards Minus Infinity rounding mode, and writes the result to the SIMD&FP destination register.
//
//go:linkname VcvtmsS32F32 VcvtmsS32F32
//go:noescape
func VcvtmsS32F32(r *arm.Int32, v0 *arm.Float32)

// Floating-point Convert to Signed integer, rounding toward Minus infinity (vector). This instruction converts a scalar or each element in a vector from a floating-point value to a signed integer value using the Round towards Minus Infinity rounding mode, and writes the result to the SIMD&FP destination register.
//
//go:linkname VcvtmdS64F64 VcvtmdS64F64
//go:noescape
func VcvtmdS64F64(r *arm.Int64, v0 *arm.Float64)

// Floating-point Convert to Unsigned integer, rounding toward Minus infinity (vector). This instruction converts a scalar or each element in a vector from a floating-point value to an unsigned integer value using the Round towards Minus Infinity rounding mode, and writes the result to the SIMD&FP destination register.
//
//go:linkname VcvtmsU32F32 VcvtmsU32F32
//go:noescape
func VcvtmsU32F32(r *arm.Uint32, v0 *arm.Float32)

// Floating-point Convert to Unsigned integer, rounding toward Minus infinity (vector). This instruction converts a scalar or each element in a vector from a floating-point value to an unsigned integer value using the Round towards Minus Infinity rounding mode, and writes the result to the SIMD&FP destination register.
//
//go:linkname VcvtmdU64F64 VcvtmdU64F64
//go:noescape
func VcvtmdU64F64(r *arm.Uint64, v0 *arm.Float64)

// Floating-point Convert to Signed integer, rounding to nearest with ties to even (vector). This instruction converts a scalar or each element in a vector from a floating-point value to a signed integer value using the Round to Nearest rounding mode, and writes the result to the SIMD&FP destination register.
//
//go:linkname VcvtnsS32F32 VcvtnsS32F32
//go:noescape
func VcvtnsS32F32(r *arm.Int32, v0 *arm.Float32)

// Floating-point Convert to Signed integer, rounding to nearest with ties to even (vector). This instruction converts a scalar or each element in a vector from a floating-point value to a signed integer value using the Round to Nearest rounding mode, and writes the result to the SIMD&FP destination register.
//
//go:linkname VcvtndS64F64 VcvtndS64F64
//go:noescape
func VcvtndS64F64(r *arm.Int64, v0 *arm.Float64)

// Floating-point Convert to Unsigned integer, rounding to nearest with ties to even (vector). This instruction converts a scalar or each element in a vector from a floating-point value to an unsigned integer value using the Round to Nearest rounding mode, and writes the result to the SIMD&FP destination register.
//
//go:linkname VcvtnsU32F32 VcvtnsU32F32
//go:noescape
func VcvtnsU32F32(r *arm.Uint32, v0 *arm.Float32)

// Floating-point Convert to Unsigned integer, rounding to nearest with ties to even (vector). This instruction converts a scalar or each element in a vector from a floating-point value to an unsigned integer value using the Round to Nearest rounding mode, and writes the result to the SIMD&FP destination register.
//
//go:linkname VcvtndU64F64 VcvtndU64F64
//go:noescape
func VcvtndU64F64(r *arm.Uint64, v0 *arm.Float64)

// Floating-point Convert to Signed integer, rounding toward Plus infinity (vector). This instruction converts a scalar or each element in a vector from a floating-point value to a signed integer value using the Round towards Plus Infinity rounding mode, and writes the result to the SIMD&FP destination register.
//
//go:linkname VcvtpsS32F32 VcvtpsS32F32
//go:noescape
func VcvtpsS32F32(r *arm.Int32, v0 *arm.Float32)

// Floating-point Convert to Signed integer, rounding toward Plus infinity (vector). This instruction converts a scalar or each element in a vector from a floating-point value to a signed integer value using the Round towards Plus Infinity rounding mode, and writes the result to the SIMD&FP destination register.
//
//go:linkname VcvtpdS64F64 VcvtpdS64F64
//go:noescape
func VcvtpdS64F64(r *arm.Int64, v0 *arm.Float64)

// Floating-point Convert to Unsigned integer, rounding toward Plus infinity (vector). This instruction converts a scalar or each element in a vector from a floating-point value to an unsigned integer value using the Round towards Plus Infinity rounding mode, and writes the result to the SIMD&FP destination register.
//
//go:linkname VcvtpsU32F32 VcvtpsU32F32
//go:noescape
func VcvtpsU32F32(r *arm.Uint32, v0 *arm.Float32)

// Floating-point Convert to Unsigned integer, rounding toward Plus infinity (vector). This instruction converts a scalar or each element in a vector from a floating-point value to an unsigned integer value using the Round towards Plus Infinity rounding mode, and writes the result to the SIMD&FP destination register.
//
//go:linkname VcvtpdU64F64 VcvtpdU64F64
//go:noescape
func VcvtpdU64F64(r *arm.Uint64, v0 *arm.Float64)

// Floating-point Convert to lower precision Narrow, rounding to odd (vector). This instruction reads each vector element in the source SIMD&FP register, narrows each value to half the precision of the source element using the Round to Odd rounding mode, writes the result to a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VcvtxdF32F64 VcvtxdF32F64
//go:noescape
func VcvtxdF32F64(r *arm.Float32, v0 *arm.Float64)

// Floating-point Convert to lower precision Narrow, rounding to odd (vector). This instruction reads each vector element in the source SIMD&FP register, narrows each value to half the precision of the source element using the Round to Odd rounding mode, writes the result to a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VcvtxF32F64 VcvtxF32F64
//go:noescape
func VcvtxF32F64(r *arm.Float32X2, v0 *arm.Float64X2)

// Floating-point Convert to lower precision Narrow, rounding to odd (vector). This instruction reads each vector element in the source SIMD&FP register, narrows each value to half the precision of the source element using the Round to Odd rounding mode, writes the result to a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VcvtxHighF32F64 VcvtxHighF32F64
//go:noescape
func VcvtxHighF32F64(r *arm.Float32X4, v0 *arm.Float32X2, v1 *arm.Float64X2)

// Floating-point Divide (vector). This instruction divides the floating-point values in the elements in the first source SIMD&FP register, by the floating-point values in the corresponding elements in the second source SIMD&FP register, places the results in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VdivqF64 VdivqF64
//go:noescape
func VdivqF64(r *arm.Float64X2, v0 *arm.Float64X2, v1 *arm.Float64X2)

// Floating-point Divide (vector). This instruction divides the floating-point values in the elements in the first source SIMD&FP register, by the floating-point values in the corresponding elements in the second source SIMD&FP register, places the results in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VdivqF32 VdivqF32
//go:noescape
func VdivqF32(r *arm.Float32X4, v0 *arm.Float32X4, v1 *arm.Float32X4)

// Floating-point Divide (vector). This instruction divides the floating-point values in the elements in the first source SIMD&FP register, by the floating-point values in the corresponding elements in the second source SIMD&FP register, places the results in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VdivF64 VdivF64
//go:noescape
func VdivF64(r *arm.Float64X1, v0 *arm.Float64X1, v1 *arm.Float64X1)

// Floating-point Divide (vector). This instruction divides the floating-point values in the elements in the first source SIMD&FP register, by the floating-point values in the corresponding elements in the second source SIMD&FP register, places the results in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VdivF32 VdivF32
//go:noescape
func VdivF32(r *arm.Float32X2, v0 *arm.Float32X2, v1 *arm.Float32X2)

// Insert vector element from another vector element. This instruction copies the vector element of the source SIMD&FP register to the specified vector element of the destination SIMD&FP register.
//
//go:linkname VdupNP64 VdupNP64
//go:noescape
func VdupNP64(r *arm.Poly64X1, v0 *arm.Poly64)

// Duplicate vector element to vector or scalar. This instruction duplicates the vector element at the specified element index in the source SIMD&FP register into a scalar or each element in a vector, and writes the result to the destination SIMD&FP register.
//
//go:linkname VdupqNP64 VdupqNP64
//go:noescape
func VdupqNP64(r *arm.Poly64X2, v0 *arm.Poly64)

// Duplicate vector element to vector or scalar. This instruction duplicates the vector element at the specified element index in the source SIMD&FP register into a scalar or each element in a vector, and writes the result to the destination SIMD&FP register.
//
//go:linkname VdupqNF64 VdupqNF64
//go:noescape
func VdupqNF64(r *arm.Float64X2, v0 *arm.Float64)

// Insert vector element from another vector element. This instruction copies the vector element of the source SIMD&FP register to the specified vector element of the destination SIMD&FP register.
//
//go:linkname VdupNF64 VdupNF64
//go:noescape
func VdupNF64(r *arm.Float64X1, v0 *arm.Float64)

// Floating-point fused Multiply-Add to accumulator (vector). This instruction multiplies corresponding floating-point values in the vectors in the two source SIMD&FP registers, adds the product to the corresponding vector element of the destination SIMD&FP register, and writes the result to the destination SIMD&FP register.
//
//go:linkname VfmaqF64 VfmaqF64
//go:noescape
func VfmaqF64(r *arm.Float64X2, v0 *arm.Float64X2, v1 *arm.Float64X2, v2 *arm.Float64X2)

// Floating-point fused Multiply-Add (scalar). This instruction multiplies the values of the first two SIMD&FP source registers, adds the product to the value of the third SIMD&FP source register, and writes the result to the SIMD&FP destination register.
//
//go:linkname VfmaF64 VfmaF64
//go:noescape
func VfmaF64(r *arm.Float64X1, v0 *arm.Float64X1, v1 *arm.Float64X1, v2 *arm.Float64X1)

// Floating-point fused Multiply-Add to accumulator (vector). This instruction multiplies corresponding floating-point values in the vectors in the two source SIMD&FP registers, adds the product to the corresponding vector element of the destination SIMD&FP register, and writes the result to the destination SIMD&FP register.
//
//go:linkname VfmaqNF64 VfmaqNF64
//go:noescape
func VfmaqNF64(r *arm.Float64X2, v0 *arm.Float64X2, v1 *arm.Float64X2, v2 *arm.Float64)

// Floating-point fused Multiply-Add (scalar). This instruction multiplies the values of the first two SIMD&FP source registers, adds the product to the value of the third SIMD&FP source register, and writes the result to the SIMD&FP destination register.
//
//go:linkname VfmaNF64 VfmaNF64
//go:noescape
func VfmaNF64(r *arm.Float64X1, v0 *arm.Float64X1, v1 *arm.Float64X1, v2 *arm.Float64)

// Floating-point fused Multiply-Subtract from accumulator (vector). This instruction multiplies corresponding floating-point values in the vectors in the two source SIMD&FP registers, negates the product, adds the result to the corresponding vector element of the destination SIMD&FP register, and writes the result to the destination SIMD&FP register.
//
//go:linkname VfmsqF64 VfmsqF64
//go:noescape
func VfmsqF64(r *arm.Float64X2, v0 *arm.Float64X2, v1 *arm.Float64X2, v2 *arm.Float64X2)

// Floating-point Fused Multiply-Subtract (scalar). This instruction multiplies the values of the first two SIMD&FP source registers, negates the product, adds that to the value of the third SIMD&FP source register, and writes the result to the SIMD&FP destination register.
//
//go:linkname VfmsF64 VfmsF64
//go:noescape
func VfmsF64(r *arm.Float64X1, v0 *arm.Float64X1, v1 *arm.Float64X1, v2 *arm.Float64X1)

// Floating-point fused Multiply-Subtract from accumulator (vector). This instruction multiplies corresponding floating-point values in the vectors in the two source SIMD&FP registers, negates the product, adds the result to the corresponding vector element of the destination SIMD&FP register, and writes the result to the destination SIMD&FP register.
//
//go:linkname VfmsqNF64 VfmsqNF64
//go:noescape
func VfmsqNF64(r *arm.Float64X2, v0 *arm.Float64X2, v1 *arm.Float64X2, v2 *arm.Float64)

// Floating-point fused Multiply-Subtract from accumulator (vector). This instruction multiplies corresponding floating-point values in the vectors in the two source SIMD&FP registers, negates the product, adds the result to the corresponding vector element of the destination SIMD&FP register, and writes the result to the destination SIMD&FP register.
//
//go:linkname VfmsqNF32 VfmsqNF32
//go:noescape
func VfmsqNF32(r *arm.Float32X4, v0 *arm.Float32X4, v1 *arm.Float32X4, v2 *arm.Float32)

// Floating-point Fused Multiply-Subtract (scalar). This instruction multiplies the values of the first two SIMD&FP source registers, negates the product, adds that to the value of the third SIMD&FP source register, and writes the result to the SIMD&FP destination register.
//
//go:linkname VfmsNF64 VfmsNF64
//go:noescape
func VfmsNF64(r *arm.Float64X1, v0 *arm.Float64X1, v1 *arm.Float64X1, v2 *arm.Float64)

// Floating-point fused Multiply-Subtract from accumulator (vector). This instruction multiplies corresponding floating-point values in the vectors in the two source SIMD&FP registers, negates the product, adds the result to the corresponding vector element of the destination SIMD&FP register, and writes the result to the destination SIMD&FP register.
//
//go:linkname VfmsNF32 VfmsNF32
//go:noescape
func VfmsNF32(r *arm.Float32X2, v0 *arm.Float32X2, v1 *arm.Float32X2, v2 *arm.Float32)

// Duplicate vector element to vector or scalar. This instruction duplicates the vector element at the specified element index in the source SIMD&FP register into a scalar or each element in a vector, and writes the result to the destination SIMD&FP register.
//
//go:linkname VgetHighP64 VgetHighP64
//go:noescape
func VgetHighP64(r *arm.Poly64X1, v0 *arm.Poly64X2)

// Duplicate vector element to vector or scalar. This instruction duplicates the vector element at the specified element index in the source SIMD&FP register into a scalar or each element in a vector, and writes the result to the destination SIMD&FP register.
//
//go:linkname VgetHighF64 VgetHighF64
//go:noescape
func VgetHighF64(r *arm.Float64X1, v0 *arm.Float64X2)

// Duplicate vector element to vector or scalar. This instruction duplicates the vector element at the specified element index in the source SIMD&FP register into a scalar or each element in a vector, and writes the result to the destination SIMD&FP register.
//
//go:linkname VgetLowP64 VgetLowP64
//go:noescape
func VgetLowP64(r *arm.Poly64X1, v0 *arm.Poly64X2)

// Duplicate vector element to vector or scalar. This instruction duplicates the vector element at the specified element index in the source SIMD&FP register into a scalar or each element in a vector, and writes the result to the destination SIMD&FP register.
//
//go:linkname VgetLowF64 VgetLowF64
//go:noescape
func VgetLowF64(r *arm.Float64X1, v0 *arm.Float64X2)

// Floating-point Maximum (vector). This instruction compares corresponding vector elements in the two source SIMD&FP registers, places the larger of each of the two floating-point values into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VmaxqF64 VmaxqF64
//go:noescape
func VmaxqF64(r *arm.Float64X2, v0 *arm.Float64X2, v1 *arm.Float64X2)

// Floating-point Maximum (vector). This instruction compares corresponding vector elements in the two source SIMD&FP registers, places the larger of each of the two floating-point values into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VmaxF64 VmaxF64
//go:noescape
func VmaxF64(r *arm.Float64X1, v0 *arm.Float64X1, v1 *arm.Float64X1)

// Floating-point Maximum Number Pairwise (vector). This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements in the two source SIMD&FP registers, writes the largest of each pair of values into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are floating-point values.
//
//go:linkname VmaxnmvqF64 VmaxnmvqF64
//go:noescape
func VmaxnmvqF64(r *arm.Float64, v0 *arm.Float64X2)

// Floating-point Maximum Number across Vector. This instruction compares all the vector elements in the source SIMD&FP register, and writes the largest of the values as a scalar to the destination SIMD&FP register. All the values in this instruction are floating-point values.
//
//go:linkname VmaxnmvqF32 VmaxnmvqF32
//go:noescape
func VmaxnmvqF32(r *arm.Float32, v0 *arm.Float32X4)

// Floating-point Maximum Number Pairwise (vector). This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements in the two source SIMD&FP registers, writes the largest of each pair of values into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are floating-point values.
//
//go:linkname VmaxnmvF32 VmaxnmvF32
//go:noescape
func VmaxnmvF32(r *arm.Float32, v0 *arm.Float32X2)

// Unsigned Maximum across Vector. This instruction compares all the vector elements in the source SIMD&FP register, and writes the largest of the values as a scalar to the destination SIMD&FP register. All the values in this instruction are unsigned integer values.
//
//go:linkname VmaxvqU8 VmaxvqU8
//go:noescape
func VmaxvqU8(r *arm.Uint8, v0 *arm.Uint8X16)

// Unsigned Maximum across Vector. This instruction compares all the vector elements in the source SIMD&FP register, and writes the largest of the values as a scalar to the destination SIMD&FP register. All the values in this instruction are unsigned integer values.
//
//go:linkname VmaxvqU32 VmaxvqU32
//go:noescape
func VmaxvqU32(r *arm.Uint32, v0 *arm.Uint32X4)

// Unsigned Maximum across Vector. This instruction compares all the vector elements in the source SIMD&FP register, and writes the largest of the values as a scalar to the destination SIMD&FP register. All the values in this instruction are unsigned integer values.
//
//go:linkname VmaxvqU16 VmaxvqU16
//go:noescape
func VmaxvqU16(r *arm.Uint16, v0 *arm.Uint16X8)

// Signed Maximum across Vector. This instruction compares all the vector elements in the source SIMD&FP register, and writes the largest of the values as a scalar to the destination SIMD&FP register. All the values in this instruction are signed integer values.
//
//go:linkname VmaxvqS8 VmaxvqS8
//go:noescape
func VmaxvqS8(r *arm.Int8, v0 *arm.Int8X16)

// Floating-point Maximum Pairwise (vector). This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements from the concatenated vector, writes the larger of each pair of values into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are floating-point values.
//
//go:linkname VmaxvqF64 VmaxvqF64
//go:noescape
func VmaxvqF64(r *arm.Float64, v0 *arm.Float64X2)

// Floating-point Maximum across Vector. This instruction compares all the vector elements in the source SIMD&FP register, and writes the largest of the values as a scalar to the destination SIMD&FP register. All the values in this instruction are floating-point values.
//
//go:linkname VmaxvqF32 VmaxvqF32
//go:noescape
func VmaxvqF32(r *arm.Float32, v0 *arm.Float32X4)

// Signed Maximum across Vector. This instruction compares all the vector elements in the source SIMD&FP register, and writes the largest of the values as a scalar to the destination SIMD&FP register. All the values in this instruction are signed integer values.
//
//go:linkname VmaxvqS32 VmaxvqS32
//go:noescape
func VmaxvqS32(r *arm.Int32, v0 *arm.Int32X4)

// Signed Maximum across Vector. This instruction compares all the vector elements in the source SIMD&FP register, and writes the largest of the values as a scalar to the destination SIMD&FP register. All the values in this instruction are signed integer values.
//
//go:linkname VmaxvqS16 VmaxvqS16
//go:noescape
func VmaxvqS16(r *arm.Int16, v0 *arm.Int16X8)

// Unsigned Maximum across Vector. This instruction compares all the vector elements in the source SIMD&FP register, and writes the largest of the values as a scalar to the destination SIMD&FP register. All the values in this instruction are unsigned integer values.
//
//go:linkname VmaxvU8 VmaxvU8
//go:noescape
func VmaxvU8(r *arm.Uint8, v0 *arm.Uint8X8)

// Unsigned Maximum Pairwise. This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements in the two source SIMD&FP registers, writes the largest of each pair of unsigned integer values into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VmaxvU32 VmaxvU32
//go:noescape
func VmaxvU32(r *arm.Uint32, v0 *arm.Uint32X2)

// Unsigned Maximum across Vector. This instruction compares all the vector elements in the source SIMD&FP register, and writes the largest of the values as a scalar to the destination SIMD&FP register. All the values in this instruction are unsigned integer values.
//
//go:linkname VmaxvU16 VmaxvU16
//go:noescape
func VmaxvU16(r *arm.Uint16, v0 *arm.Uint16X4)

// Signed Maximum across Vector. This instruction compares all the vector elements in the source SIMD&FP register, and writes the largest of the values as a scalar to the destination SIMD&FP register. All the values in this instruction are signed integer values.
//
//go:linkname VmaxvS8 VmaxvS8
//go:noescape
func VmaxvS8(r *arm.Int8, v0 *arm.Int8X8)

// Floating-point Maximum Pairwise (vector). This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements from the concatenated vector, writes the larger of each pair of values into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are floating-point values.
//
//go:linkname VmaxvF32 VmaxvF32
//go:noescape
func VmaxvF32(r *arm.Float32, v0 *arm.Float32X2)

// Signed Maximum Pairwise. This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements in the two source SIMD&FP registers, writes the largest of each pair of signed integer values into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VmaxvS32 VmaxvS32
//go:noescape
func VmaxvS32(r *arm.Int32, v0 *arm.Int32X2)

// Signed Maximum across Vector. This instruction compares all the vector elements in the source SIMD&FP register, and writes the largest of the values as a scalar to the destination SIMD&FP register. All the values in this instruction are signed integer values.
//
//go:linkname VmaxvS16 VmaxvS16
//go:noescape
func VmaxvS16(r *arm.Int16, v0 *arm.Int16X4)

// Floating-point minimum (vector). This instruction compares corresponding elements in the vectors in the two source SIMD&FP registers, places the smaller of each of the two floating-point values into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VminqF64 VminqF64
//go:noescape
func VminqF64(r *arm.Float64X2, v0 *arm.Float64X2, v1 *arm.Float64X2)

// Floating-point minimum (vector). This instruction compares corresponding elements in the vectors in the two source SIMD&FP registers, places the smaller of each of the two floating-point values into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VminF64 VminF64
//go:noescape
func VminF64(r *arm.Float64X1, v0 *arm.Float64X1, v1 *arm.Float64X1)

// Floating-point Minimum Number Pairwise (vector). This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements in the two source SIMD&FP registers, writes the smallest of each pair of floating-point values into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are floating-point values.
//
//go:linkname VminnmvqF64 VminnmvqF64
//go:noescape
func VminnmvqF64(r *arm.Float64, v0 *arm.Float64X2)

// Floating-point Minimum Number across Vector. This instruction compares all the vector elements in the source SIMD&FP register, and writes the smallest of the values as a scalar to the destination SIMD&FP register. All the values in this instruction are floating-point values.
//
//go:linkname VminnmvqF32 VminnmvqF32
//go:noescape
func VminnmvqF32(r *arm.Float32, v0 *arm.Float32X4)

// Floating-point Minimum Number Pairwise (vector). This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements in the two source SIMD&FP registers, writes the smallest of each pair of floating-point values into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are floating-point values.
//
//go:linkname VminnmvF32 VminnmvF32
//go:noescape
func VminnmvF32(r *arm.Float32, v0 *arm.Float32X2)

// Unsigned Minimum across Vector. This instruction compares all the vector elements in the source SIMD&FP register, and writes the smallest of the values as a scalar to the destination SIMD&FP register. All the values in this instruction are unsigned integer values.
//
//go:linkname VminvqU8 VminvqU8
//go:noescape
func VminvqU8(r *arm.Uint8, v0 *arm.Uint8X16)

// Unsigned Minimum across Vector. This instruction compares all the vector elements in the source SIMD&FP register, and writes the smallest of the values as a scalar to the destination SIMD&FP register. All the values in this instruction are unsigned integer values.
//
//go:linkname VminvqU32 VminvqU32
//go:noescape
func VminvqU32(r *arm.Uint32, v0 *arm.Uint32X4)

// Unsigned Minimum across Vector. This instruction compares all the vector elements in the source SIMD&FP register, and writes the smallest of the values as a scalar to the destination SIMD&FP register. All the values in this instruction are unsigned integer values.
//
//go:linkname VminvqU16 VminvqU16
//go:noescape
func VminvqU16(r *arm.Uint16, v0 *arm.Uint16X8)

// Signed Minimum across Vector. This instruction compares all the vector elements in the source SIMD&FP register, and writes the smallest of the values as a scalar to the destination SIMD&FP register. All the values in this instruction are signed integer values.
//
//go:linkname VminvqS8 VminvqS8
//go:noescape
func VminvqS8(r *arm.Int8, v0 *arm.Int8X16)

// Floating-point Minimum Pairwise (vector). This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements from the concatenated vector, writes the smaller of each pair of values into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are floating-point values.
//
//go:linkname VminvqF64 VminvqF64
//go:noescape
func VminvqF64(r *arm.Float64, v0 *arm.Float64X2)

// Floating-point Minimum across Vector. This instruction compares all the vector elements in the source SIMD&FP register, and writes the smallest of the values as a scalar to the destination SIMD&FP register. All the values in this instruction are floating-point values.
//
//go:linkname VminvqF32 VminvqF32
//go:noescape
func VminvqF32(r *arm.Float32, v0 *arm.Float32X4)

// Signed Minimum across Vector. This instruction compares all the vector elements in the source SIMD&FP register, and writes the smallest of the values as a scalar to the destination SIMD&FP register. All the values in this instruction are signed integer values.
//
//go:linkname VminvqS32 VminvqS32
//go:noescape
func VminvqS32(r *arm.Int32, v0 *arm.Int32X4)

// Signed Minimum across Vector. This instruction compares all the vector elements in the source SIMD&FP register, and writes the smallest of the values as a scalar to the destination SIMD&FP register. All the values in this instruction are signed integer values.
//
//go:linkname VminvqS16 VminvqS16
//go:noescape
func VminvqS16(r *arm.Int16, v0 *arm.Int16X8)

// Unsigned Minimum across Vector. This instruction compares all the vector elements in the source SIMD&FP register, and writes the smallest of the values as a scalar to the destination SIMD&FP register. All the values in this instruction are unsigned integer values.
//
//go:linkname VminvU8 VminvU8
//go:noescape
func VminvU8(r *arm.Uint8, v0 *arm.Uint8X8)

// Unsigned Minimum Pairwise. This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements in the two source SIMD&FP registers, writes the smallest of each pair of unsigned integer values into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VminvU32 VminvU32
//go:noescape
func VminvU32(r *arm.Uint32, v0 *arm.Uint32X2)

// Unsigned Minimum across Vector. This instruction compares all the vector elements in the source SIMD&FP register, and writes the smallest of the values as a scalar to the destination SIMD&FP register. All the values in this instruction are unsigned integer values.
//
//go:linkname VminvU16 VminvU16
//go:noescape
func VminvU16(r *arm.Uint16, v0 *arm.Uint16X4)

// Signed Minimum across Vector. This instruction compares all the vector elements in the source SIMD&FP register, and writes the smallest of the values as a scalar to the destination SIMD&FP register. All the values in this instruction are signed integer values.
//
//go:linkname VminvS8 VminvS8
//go:noescape
func VminvS8(r *arm.Int8, v0 *arm.Int8X8)

// Floating-point Minimum Pairwise (vector). This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements from the concatenated vector, writes the smaller of each pair of values into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are floating-point values.
//
//go:linkname VminvF32 VminvF32
//go:noescape
func VminvF32(r *arm.Float32, v0 *arm.Float32X2)

// Signed Minimum Pairwise. This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements in the two source SIMD&FP registers, writes the smallest of each pair of signed integer values into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VminvS32 VminvS32
//go:noescape
func VminvS32(r *arm.Int32, v0 *arm.Int32X2)

// Signed Minimum across Vector. This instruction compares all the vector elements in the source SIMD&FP register, and writes the smallest of the values as a scalar to the destination SIMD&FP register. All the values in this instruction are signed integer values.
//
//go:linkname VminvS16 VminvS16
//go:noescape
func VminvS16(r *arm.Int16, v0 *arm.Int16X4)

// Floating-point multiply-add to accumulator
//
//go:linkname VmlaqF64 VmlaqF64
//go:noescape
func VmlaqF64(r *arm.Float64X2, v0 *arm.Float64X2, v1 *arm.Float64X2, v2 *arm.Float64X2)

// Floating-point multiply-add to accumulator
//
//go:linkname VmlaF64 VmlaF64
//go:noescape
func VmlaF64(r *arm.Float64X1, v0 *arm.Float64X1, v1 *arm.Float64X1, v2 *arm.Float64X1)

// Multiply-subtract from accumulator
//
//go:linkname VmlsqF64 VmlsqF64
//go:noescape
func VmlsqF64(r *arm.Float64X2, v0 *arm.Float64X2, v1 *arm.Float64X2, v2 *arm.Float64X2)

// Multiply-subtract from accumulator
//
//go:linkname VmlsF64 VmlsF64
//go:noescape
func VmlsF64(r *arm.Float64X1, v0 *arm.Float64X1, v1 *arm.Float64X1, v2 *arm.Float64X1)

// vmov_n_p64
//
//go:linkname VmovNP64 VmovNP64
//go:noescape
func VmovNP64(r *arm.Poly64X1, v0 *arm.Poly64)

// vmovq_n_p64
//
//go:linkname VmovqNP64 VmovqNP64
//go:noescape
func VmovqNP64(r *arm.Poly64X2, v0 *arm.Poly64)

// Duplicate vector element to vector or scalar. This instruction duplicates the vector element at the specified element index in the source SIMD&FP register into a scalar or each element in a vector, and writes the result to the destination SIMD&FP register.
//
//go:linkname VmovqNF64 VmovqNF64
//go:noescape
func VmovqNF64(r *arm.Float64X2, v0 *arm.Float64)

// Duplicate vector element to vector or scalar. This instruction duplicates the vector element at the specified element index in the source SIMD&FP register into a scalar or each element in a vector, and writes the result to the destination SIMD&FP register.
//
//go:linkname VmovNF64 VmovNF64
//go:noescape
func VmovNF64(r *arm.Float64X1, v0 *arm.Float64)

// Vector move
//
//go:linkname VmovlHighU8 VmovlHighU8
//go:noescape
func VmovlHighU8(r *arm.Uint16X8, v0 *arm.Uint8X16)

// Vector move
//
//go:linkname VmovlHighU32 VmovlHighU32
//go:noescape
func VmovlHighU32(r *arm.Uint64X2, v0 *arm.Uint32X4)

// Vector move
//
//go:linkname VmovlHighU16 VmovlHighU16
//go:noescape
func VmovlHighU16(r *arm.Uint32X4, v0 *arm.Uint16X8)

// Vector move
//
//go:linkname VmovlHighS8 VmovlHighS8
//go:noescape
func VmovlHighS8(r *arm.Int16X8, v0 *arm.Int8X16)

// Vector move
//
//go:linkname VmovlHighS32 VmovlHighS32
//go:noescape
func VmovlHighS32(r *arm.Int64X2, v0 *arm.Int32X4)

// Vector move
//
//go:linkname VmovlHighS16 VmovlHighS16
//go:noescape
func VmovlHighS16(r *arm.Int32X4, v0 *arm.Int16X8)

// Extract Narrow. This instruction reads each vector element from the source SIMD&FP register, narrows each value to half the original width, places the result into a vector, and writes the vector to the lower or upper half of the destination SIMD&FP register. The destination vector elements are half as long as the source vector elements.
//
//go:linkname VmovnHighU32 VmovnHighU32
//go:noescape
func VmovnHighU32(r *arm.Uint16X8, v0 *arm.Uint16X4, v1 *arm.Uint32X4)

// Extract Narrow. This instruction reads each vector element from the source SIMD&FP register, narrows each value to half the original width, places the result into a vector, and writes the vector to the lower or upper half of the destination SIMD&FP register. The destination vector elements are half as long as the source vector elements.
//
//go:linkname VmovnHighU64 VmovnHighU64
//go:noescape
func VmovnHighU64(r *arm.Uint32X4, v0 *arm.Uint32X2, v1 *arm.Uint64X2)

// Extract Narrow. This instruction reads each vector element from the source SIMD&FP register, narrows each value to half the original width, places the result into a vector, and writes the vector to the lower or upper half of the destination SIMD&FP register. The destination vector elements are half as long as the source vector elements.
//
//go:linkname VmovnHighU16 VmovnHighU16
//go:noescape
func VmovnHighU16(r *arm.Uint8X16, v0 *arm.Uint8X8, v1 *arm.Uint16X8)

// Extract Narrow. This instruction reads each vector element from the source SIMD&FP register, narrows each value to half the original width, places the result into a vector, and writes the vector to the lower or upper half of the destination SIMD&FP register. The destination vector elements are half as long as the source vector elements.
//
//go:linkname VmovnHighS32 VmovnHighS32
//go:noescape
func VmovnHighS32(r *arm.Int16X8, v0 *arm.Int16X4, v1 *arm.Int32X4)

// Extract Narrow. This instruction reads each vector element from the source SIMD&FP register, narrows each value to half the original width, places the result into a vector, and writes the vector to the lower or upper half of the destination SIMD&FP register. The destination vector elements are half as long as the source vector elements.
//
//go:linkname VmovnHighS64 VmovnHighS64
//go:noescape
func VmovnHighS64(r *arm.Int32X4, v0 *arm.Int32X2, v1 *arm.Int64X2)

// Extract Narrow. This instruction reads each vector element from the source SIMD&FP register, narrows each value to half the original width, places the result into a vector, and writes the vector to the lower or upper half of the destination SIMD&FP register. The destination vector elements are half as long as the source vector elements.
//
//go:linkname VmovnHighS16 VmovnHighS16
//go:noescape
func VmovnHighS16(r *arm.Int8X16, v0 *arm.Int8X8, v1 *arm.Int16X8)

// Floating-point Multiply (vector). This instruction multiplies corresponding floating-point values in the vectors in the two source SIMD&FP registers, places the result in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VmulqF64 VmulqF64
//go:noescape
func VmulqF64(r *arm.Float64X2, v0 *arm.Float64X2, v1 *arm.Float64X2)

// Floating-point Multiply (vector). This instruction multiplies corresponding floating-point values in the vectors in the two source SIMD&FP registers, places the result in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VmulF64 VmulF64
//go:noescape
func VmulF64(r *arm.Float64X1, v0 *arm.Float64X1, v1 *arm.Float64X1)

// Floating-point Multiply (vector). This instruction multiplies corresponding floating-point values in the vectors in the two source SIMD&FP registers, places the result in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VmulNF64 VmulNF64
//go:noescape
func VmulNF64(r *arm.Float64X1, v0 *arm.Float64X1, v1 *arm.Float64)

// Floating-point Multiply (vector). This instruction multiplies corresponding floating-point values in the vectors in the two source SIMD&FP registers, places the result in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VmulqNF64 VmulqNF64
//go:noescape
func VmulqNF64(r *arm.Float64X2, v0 *arm.Float64X2, v1 *arm.Float64)

// Polynomial Multiply Long. This instruction multiplies corresponding elements in the lower or upper half of the vectors of the two source SIMD&FP registers, places the results in a vector, and writes the vector to the destination SIMD&FP register. The destination vector elements are twice as long as the elements that are multiplied.
//
//go:linkname VmullP64 VmullP64
//go:noescape
func VmullP64(r *arm.Poly128, v0 *arm.Poly64, v1 *arm.Poly64)

// Polynomial Multiply Long. This instruction multiplies corresponding elements in the lower or upper half of the vectors of the two source SIMD&FP registers, places the results in a vector, and writes the vector to the destination SIMD&FP register. The destination vector elements are twice as long as the elements that are multiplied.
//
//go:linkname VmullHighP8 VmullHighP8
//go:noescape
func VmullHighP8(r *arm.Poly16X8, v0 *arm.Poly8X16, v1 *arm.Poly8X16)

// Unsigned Multiply long (vector). This instruction multiplies corresponding vector elements in the lower or upper half of the two source SIMD&FP registers, places the result in a vector, and writes the vector to the destination SIMD&FP register. The destination vector elements are twice as long as the elements that are multiplied. All the values in this instruction are unsigned integer values.
//
//go:linkname VmullHighU8 VmullHighU8
//go:noescape
func VmullHighU8(r *arm.Uint16X8, v0 *arm.Uint8X16, v1 *arm.Uint8X16)

// Unsigned Multiply long (vector). This instruction multiplies corresponding vector elements in the lower or upper half of the two source SIMD&FP registers, places the result in a vector, and writes the vector to the destination SIMD&FP register. The destination vector elements are twice as long as the elements that are multiplied. All the values in this instruction are unsigned integer values.
//
//go:linkname VmullHighU32 VmullHighU32
//go:noescape
func VmullHighU32(r *arm.Uint64X2, v0 *arm.Uint32X4, v1 *arm.Uint32X4)

// Unsigned Multiply long (vector). This instruction multiplies corresponding vector elements in the lower or upper half of the two source SIMD&FP registers, places the result in a vector, and writes the vector to the destination SIMD&FP register. The destination vector elements are twice as long as the elements that are multiplied. All the values in this instruction are unsigned integer values.
//
//go:linkname VmullHighU16 VmullHighU16
//go:noescape
func VmullHighU16(r *arm.Uint32X4, v0 *arm.Uint16X8, v1 *arm.Uint16X8)

// Signed Multiply Long (vector). This instruction multiplies corresponding signed integer values in the lower or upper half of the vectors of the two source SIMD&FP registers, places the results in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VmullHighS8 VmullHighS8
//go:noescape
func VmullHighS8(r *arm.Int16X8, v0 *arm.Int8X16, v1 *arm.Int8X16)

// Signed Multiply Long (vector). This instruction multiplies corresponding signed integer values in the lower or upper half of the vectors of the two source SIMD&FP registers, places the results in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VmullHighS32 VmullHighS32
//go:noescape
func VmullHighS32(r *arm.Int64X2, v0 *arm.Int32X4, v1 *arm.Int32X4)

// Signed Multiply Long (vector). This instruction multiplies corresponding signed integer values in the lower or upper half of the vectors of the two source SIMD&FP registers, places the results in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VmullHighS16 VmullHighS16
//go:noescape
func VmullHighS16(r *arm.Int32X4, v0 *arm.Int16X8, v1 *arm.Int16X8)

// Polynomial Multiply Long. This instruction multiplies corresponding elements in the lower or upper half of the vectors of the two source SIMD&FP registers, places the results in a vector, and writes the vector to the destination SIMD&FP register. The destination vector elements are twice as long as the elements that are multiplied.
//
//go:linkname VmullHighP64 VmullHighP64
//go:noescape
func VmullHighP64(r *arm.Poly128, v0 *arm.Poly64X2, v1 *arm.Poly64X2)

// Unsigned Multiply long (vector). This instruction multiplies corresponding vector elements in the lower or upper half of the two source SIMD&FP registers, places the result in a vector, and writes the vector to the destination SIMD&FP register. The destination vector elements are twice as long as the elements that are multiplied. All the values in this instruction are unsigned integer values.
//
//go:linkname VmullHighNU32 VmullHighNU32
//go:noescape
func VmullHighNU32(r *arm.Uint64X2, v0 *arm.Uint32X4, v1 *arm.Uint32)

// Unsigned Multiply long (vector). This instruction multiplies corresponding vector elements in the lower or upper half of the two source SIMD&FP registers, places the result in a vector, and writes the vector to the destination SIMD&FP register. The destination vector elements are twice as long as the elements that are multiplied. All the values in this instruction are unsigned integer values.
//
//go:linkname VmullHighNU16 VmullHighNU16
//go:noescape
func VmullHighNU16(r *arm.Uint32X4, v0 *arm.Uint16X8, v1 *arm.Uint16)

// Signed Multiply Long (vector). This instruction multiplies corresponding signed integer values in the lower or upper half of the vectors of the two source SIMD&FP registers, places the results in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VmullHighNS32 VmullHighNS32
//go:noescape
func VmullHighNS32(r *arm.Int64X2, v0 *arm.Int32X4, v1 *arm.Int32)

// Signed Multiply Long (vector). This instruction multiplies corresponding signed integer values in the lower or upper half of the vectors of the two source SIMD&FP registers, places the results in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VmullHighNS16 VmullHighNS16
//go:noescape
func VmullHighNS16(r *arm.Int32X4, v0 *arm.Int16X8, v1 *arm.Int16)

// Floating-point Multiply extended. This instruction multiplies corresponding floating-point values in the vectors of the two source SIMD&FP registers, places the resulting floating-point values in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VmulxqF64 VmulxqF64
//go:noescape
func VmulxqF64(r *arm.Float64X2, v0 *arm.Float64X2, v1 *arm.Float64X2)

// Floating-point Multiply extended. This instruction multiplies corresponding floating-point values in the vectors of the two source SIMD&FP registers, places the resulting floating-point values in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VmulxqF32 VmulxqF32
//go:noescape
func VmulxqF32(r *arm.Float32X4, v0 *arm.Float32X4, v1 *arm.Float32X4)

// Floating-point Multiply extended. This instruction multiplies corresponding floating-point values in the vectors of the two source SIMD&FP registers, places the resulting floating-point values in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VmulxF64 VmulxF64
//go:noescape
func VmulxF64(r *arm.Float64X1, v0 *arm.Float64X1, v1 *arm.Float64X1)

// Floating-point Multiply extended. This instruction multiplies corresponding floating-point values in the vectors of the two source SIMD&FP registers, places the resulting floating-point values in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VmulxF32 VmulxF32
//go:noescape
func VmulxF32(r *arm.Float32X2, v0 *arm.Float32X2, v1 *arm.Float32X2)

// Floating-point Multiply extended. This instruction multiplies corresponding floating-point values in the vectors of the two source SIMD&FP registers, places the resulting floating-point values in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VmulxdF64 VmulxdF64
//go:noescape
func VmulxdF64(r *arm.Float64, v0 *arm.Float64, v1 *arm.Float64)

// Floating-point Multiply extended. This instruction multiplies corresponding floating-point values in the vectors of the two source SIMD&FP registers, places the resulting floating-point values in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VmulxsF32 VmulxsF32
//go:noescape
func VmulxsF32(r *arm.Float32, v0 *arm.Float32, v1 *arm.Float32)

// Floating-point Negate (vector). This instruction negates the value of each vector element in the source SIMD&FP register, writes the result to a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VnegqF64 VnegqF64
//go:noescape
func VnegqF64(r *arm.Float64X2, v0 *arm.Float64X2)

// Negate (vector). This instruction reads each vector element from the source SIMD&FP register, negates each value, puts the result into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VnegqS64 VnegqS64
//go:noescape
func VnegqS64(r *arm.Int64X2, v0 *arm.Int64X2)

// Floating-point Negate (vector). This instruction negates the value of each vector element in the source SIMD&FP register, writes the result to a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VnegF64 VnegF64
//go:noescape
func VnegF64(r *arm.Float64X1, v0 *arm.Float64X1)

// Negate (vector). This instruction reads each vector element from the source SIMD&FP register, negates each value, puts the result into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VnegS64 VnegS64
//go:noescape
func VnegS64(r *arm.Int64X1, v0 *arm.Int64X1)

// Negate (vector). This instruction reads each vector element from the source SIMD&FP register, negates each value, puts the result into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VnegdS64 VnegdS64
//go:noescape
func VnegdS64(r *arm.Int64, v0 *arm.Int64)

// Add Pairwise (vector). This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements from the concatenated vector, adds each pair of values together, places the result into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VpaddqU8 VpaddqU8
//go:noescape
func VpaddqU8(r *arm.Uint8X16, v0 *arm.Uint8X16, v1 *arm.Uint8X16)

// Add Pairwise (vector). This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements from the concatenated vector, adds each pair of values together, places the result into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VpaddqU32 VpaddqU32
//go:noescape
func VpaddqU32(r *arm.Uint32X4, v0 *arm.Uint32X4, v1 *arm.Uint32X4)

// Add Pairwise (vector). This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements from the concatenated vector, adds each pair of values together, places the result into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VpaddqU64 VpaddqU64
//go:noescape
func VpaddqU64(r *arm.Uint64X2, v0 *arm.Uint64X2, v1 *arm.Uint64X2)

// Add Pairwise (vector). This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements from the concatenated vector, adds each pair of values together, places the result into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VpaddqU16 VpaddqU16
//go:noescape
func VpaddqU16(r *arm.Uint16X8, v0 *arm.Uint16X8, v1 *arm.Uint16X8)

// Add Pairwise (vector). This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements from the concatenated vector, adds each pair of values together, places the result into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VpaddqS8 VpaddqS8
//go:noescape
func VpaddqS8(r *arm.Int8X16, v0 *arm.Int8X16, v1 *arm.Int8X16)

// Floating-point Add Pairwise (vector). This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements from the concatenated vector, adds each pair of values together, places the result into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are floating-point values.
//
//go:linkname VpaddqF64 VpaddqF64
//go:noescape
func VpaddqF64(r *arm.Float64X2, v0 *arm.Float64X2, v1 *arm.Float64X2)

// Floating-point Add Pairwise (vector). This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements from the concatenated vector, adds each pair of values together, places the result into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are floating-point values.
//
//go:linkname VpaddqF32 VpaddqF32
//go:noescape
func VpaddqF32(r *arm.Float32X4, v0 *arm.Float32X4, v1 *arm.Float32X4)

// Add Pairwise (vector). This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements from the concatenated vector, adds each pair of values together, places the result into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VpaddqS32 VpaddqS32
//go:noescape
func VpaddqS32(r *arm.Int32X4, v0 *arm.Int32X4, v1 *arm.Int32X4)

// Add Pairwise (vector). This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements from the concatenated vector, adds each pair of values together, places the result into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VpaddqS64 VpaddqS64
//go:noescape
func VpaddqS64(r *arm.Int64X2, v0 *arm.Int64X2, v1 *arm.Int64X2)

// Add Pairwise (vector). This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements from the concatenated vector, adds each pair of values together, places the result into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VpaddqS16 VpaddqS16
//go:noescape
func VpaddqS16(r *arm.Int16X8, v0 *arm.Int16X8, v1 *arm.Int16X8)

// Add Pairwise (vector). This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements from the concatenated vector, adds each pair of values together, places the result into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VpadddU64 VpadddU64
//go:noescape
func VpadddU64(r *arm.Uint64, v0 *arm.Uint64X2)

// Floating-point Add Pairwise (vector). This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements from the concatenated vector, adds each pair of values together, places the result into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are floating-point values.
//
//go:linkname VpadddF64 VpadddF64
//go:noescape
func VpadddF64(r *arm.Float64, v0 *arm.Float64X2)

// Add Pairwise (vector). This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements from the concatenated vector, adds each pair of values together, places the result into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VpadddS64 VpadddS64
//go:noescape
func VpadddS64(r *arm.Int64, v0 *arm.Int64X2)

// Floating-point Add Pairwise (vector). This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements from the concatenated vector, adds each pair of values together, places the result into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are floating-point values.
//
//go:linkname VpaddsF32 VpaddsF32
//go:noescape
func VpaddsF32(r *arm.Float32, v0 *arm.Float32X2)

// Unsigned Maximum Pairwise. This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements in the two source SIMD&FP registers, writes the largest of each pair of unsigned integer values into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VpmaxqU8 VpmaxqU8
//go:noescape
func VpmaxqU8(r *arm.Uint8X16, v0 *arm.Uint8X16, v1 *arm.Uint8X16)

// Unsigned Maximum Pairwise. This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements in the two source SIMD&FP registers, writes the largest of each pair of unsigned integer values into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VpmaxqU32 VpmaxqU32
//go:noescape
func VpmaxqU32(r *arm.Uint32X4, v0 *arm.Uint32X4, v1 *arm.Uint32X4)

// Unsigned Maximum Pairwise. This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements in the two source SIMD&FP registers, writes the largest of each pair of unsigned integer values into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VpmaxqU16 VpmaxqU16
//go:noescape
func VpmaxqU16(r *arm.Uint16X8, v0 *arm.Uint16X8, v1 *arm.Uint16X8)

// Signed Maximum Pairwise. This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements in the two source SIMD&FP registers, writes the largest of each pair of signed integer values into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VpmaxqS8 VpmaxqS8
//go:noescape
func VpmaxqS8(r *arm.Int8X16, v0 *arm.Int8X16, v1 *arm.Int8X16)

// Floating-point Maximum Pairwise (vector). This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements from the concatenated vector, writes the larger of each pair of values into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are floating-point values.
//
//go:linkname VpmaxqF64 VpmaxqF64
//go:noescape
func VpmaxqF64(r *arm.Float64X2, v0 *arm.Float64X2, v1 *arm.Float64X2)

// Floating-point Maximum Pairwise (vector). This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements from the concatenated vector, writes the larger of each pair of values into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are floating-point values.
//
//go:linkname VpmaxqF32 VpmaxqF32
//go:noescape
func VpmaxqF32(r *arm.Float32X4, v0 *arm.Float32X4, v1 *arm.Float32X4)

// Signed Maximum Pairwise. This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements in the two source SIMD&FP registers, writes the largest of each pair of signed integer values into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VpmaxqS32 VpmaxqS32
//go:noescape
func VpmaxqS32(r *arm.Int32X4, v0 *arm.Int32X4, v1 *arm.Int32X4)

// Signed Maximum Pairwise. This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements in the two source SIMD&FP registers, writes the largest of each pair of signed integer values into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VpmaxqS16 VpmaxqS16
//go:noescape
func VpmaxqS16(r *arm.Int16X8, v0 *arm.Int16X8, v1 *arm.Int16X8)

// Floating-point Maximum Pairwise (vector). This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements from the concatenated vector, writes the larger of each pair of values into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are floating-point values.
//
//go:linkname VpmaxqdF64 VpmaxqdF64
//go:noescape
func VpmaxqdF64(r *arm.Float64, v0 *arm.Float64X2)

// Floating-point Maximum Pairwise (vector). This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements from the concatenated vector, writes the larger of each pair of values into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are floating-point values.
//
//go:linkname VpmaxsF32 VpmaxsF32
//go:noescape
func VpmaxsF32(r *arm.Float32, v0 *arm.Float32X2)

// Floating-point Maximum Number Pairwise (vector). This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements in the two source SIMD&FP registers, writes the largest of each pair of values into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are floating-point values.
//
//go:linkname VpmaxnmqF64 VpmaxnmqF64
//go:noescape
func VpmaxnmqF64(r *arm.Float64X2, v0 *arm.Float64X2, v1 *arm.Float64X2)

// Floating-point Maximum Number Pairwise (vector). This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements in the two source SIMD&FP registers, writes the largest of each pair of values into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are floating-point values.
//
//go:linkname VpmaxnmqF32 VpmaxnmqF32
//go:noescape
func VpmaxnmqF32(r *arm.Float32X4, v0 *arm.Float32X4, v1 *arm.Float32X4)

// Floating-point Maximum Number Pairwise (vector). This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements in the two source SIMD&FP registers, writes the largest of each pair of values into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are floating-point values.
//
//go:linkname VpmaxnmF32 VpmaxnmF32
//go:noescape
func VpmaxnmF32(r *arm.Float32X2, v0 *arm.Float32X2, v1 *arm.Float32X2)

// Floating-point Maximum Number Pairwise (vector). This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements in the two source SIMD&FP registers, writes the largest of each pair of values into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are floating-point values.
//
//go:linkname VpmaxnmqdF64 VpmaxnmqdF64
//go:noescape
func VpmaxnmqdF64(r *arm.Float64, v0 *arm.Float64X2)

// Floating-point Maximum Number Pairwise (vector). This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements in the two source SIMD&FP registers, writes the largest of each pair of values into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are floating-point values.
//
//go:linkname VpmaxnmsF32 VpmaxnmsF32
//go:noescape
func VpmaxnmsF32(r *arm.Float32, v0 *arm.Float32X2)

// Unsigned Minimum Pairwise. This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements in the two source SIMD&FP registers, writes the smallest of each pair of unsigned integer values into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VpminqU8 VpminqU8
//go:noescape
func VpminqU8(r *arm.Uint8X16, v0 *arm.Uint8X16, v1 *arm.Uint8X16)

// Unsigned Minimum Pairwise. This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements in the two source SIMD&FP registers, writes the smallest of each pair of unsigned integer values into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VpminqU32 VpminqU32
//go:noescape
func VpminqU32(r *arm.Uint32X4, v0 *arm.Uint32X4, v1 *arm.Uint32X4)

// Unsigned Minimum Pairwise. This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements in the two source SIMD&FP registers, writes the smallest of each pair of unsigned integer values into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VpminqU16 VpminqU16
//go:noescape
func VpminqU16(r *arm.Uint16X8, v0 *arm.Uint16X8, v1 *arm.Uint16X8)

// Signed Minimum Pairwise. This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements in the two source SIMD&FP registers, writes the smallest of each pair of signed integer values into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VpminqS8 VpminqS8
//go:noescape
func VpminqS8(r *arm.Int8X16, v0 *arm.Int8X16, v1 *arm.Int8X16)

// Floating-point Minimum Pairwise (vector). This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements from the concatenated vector, writes the smaller of each pair of values into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are floating-point values.
//
//go:linkname VpminqF64 VpminqF64
//go:noescape
func VpminqF64(r *arm.Float64X2, v0 *arm.Float64X2, v1 *arm.Float64X2)

// Floating-point Minimum Pairwise (vector). This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements from the concatenated vector, writes the smaller of each pair of values into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are floating-point values.
//
//go:linkname VpminqF32 VpminqF32
//go:noescape
func VpminqF32(r *arm.Float32X4, v0 *arm.Float32X4, v1 *arm.Float32X4)

// Signed Minimum Pairwise. This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements in the two source SIMD&FP registers, writes the smallest of each pair of signed integer values into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VpminqS32 VpminqS32
//go:noescape
func VpminqS32(r *arm.Int32X4, v0 *arm.Int32X4, v1 *arm.Int32X4)

// Signed Minimum Pairwise. This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements in the two source SIMD&FP registers, writes the smallest of each pair of signed integer values into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VpminqS16 VpminqS16
//go:noescape
func VpminqS16(r *arm.Int16X8, v0 *arm.Int16X8, v1 *arm.Int16X8)

// Floating-point Minimum Pairwise (vector). This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements from the concatenated vector, writes the smaller of each pair of values into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are floating-point values.
//
//go:linkname VpminqdF64 VpminqdF64
//go:noescape
func VpminqdF64(r *arm.Float64, v0 *arm.Float64X2)

// Floating-point Minimum Pairwise (vector). This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements from the concatenated vector, writes the smaller of each pair of values into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are floating-point values.
//
//go:linkname VpminsF32 VpminsF32
//go:noescape
func VpminsF32(r *arm.Float32, v0 *arm.Float32X2)

// Floating-point Minimum Number Pairwise (vector). This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements in the two source SIMD&FP registers, writes the smallest of each pair of floating-point values into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are floating-point values.
//
//go:linkname VpminnmqF64 VpminnmqF64
//go:noescape
func VpminnmqF64(r *arm.Float64X2, v0 *arm.Float64X2, v1 *arm.Float64X2)

// Floating-point Minimum Number Pairwise (vector). This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements in the two source SIMD&FP registers, writes the smallest of each pair of floating-point values into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are floating-point values.
//
//go:linkname VpminnmqF32 VpminnmqF32
//go:noescape
func VpminnmqF32(r *arm.Float32X4, v0 *arm.Float32X4, v1 *arm.Float32X4)

// Floating-point Minimum Number Pairwise (vector). This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements in the two source SIMD&FP registers, writes the smallest of each pair of floating-point values into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are floating-point values.
//
//go:linkname VpminnmF32 VpminnmF32
//go:noescape
func VpminnmF32(r *arm.Float32X2, v0 *arm.Float32X2, v1 *arm.Float32X2)

// Floating-point Minimum Number Pairwise (vector). This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements in the two source SIMD&FP registers, writes the smallest of each pair of floating-point values into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are floating-point values.
//
//go:linkname VpminnmqdF64 VpminnmqdF64
//go:noescape
func VpminnmqdF64(r *arm.Float64, v0 *arm.Float64X2)

// Floating-point Minimum Number Pairwise (vector). This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements in the two source SIMD&FP registers, writes the smallest of each pair of floating-point values into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are floating-point values.
//
//go:linkname VpminnmsF32 VpminnmsF32
//go:noescape
func VpminnmsF32(r *arm.Float32, v0 *arm.Float32X2)

// Signed saturating Absolute value. This instruction reads each vector element from the source SIMD&FP register, puts the absolute value of the result into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are signed integer values.
//
//go:linkname VqabsqS64 VqabsqS64
//go:noescape
func VqabsqS64(r *arm.Int64X2, v0 *arm.Int64X2)

// Signed saturating Absolute value. This instruction reads each vector element from the source SIMD&FP register, puts the absolute value of the result into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are signed integer values.
//
//go:linkname VqabsS64 VqabsS64
//go:noescape
func VqabsS64(r *arm.Int64X1, v0 *arm.Int64X1)

// Signed saturating Absolute value. This instruction reads each vector element from the source SIMD&FP register, puts the absolute value of the result into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are signed integer values.
//
//go:linkname VqabsbS8 VqabsbS8
//go:noescape
func VqabsbS8(r *arm.Int8, v0 *arm.Int8)

// Signed saturating Absolute value. This instruction reads each vector element from the source SIMD&FP register, puts the absolute value of the result into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are signed integer values.
//
//go:linkname VqabssS32 VqabssS32
//go:noescape
func VqabssS32(r *arm.Int32, v0 *arm.Int32)

// Signed saturating Absolute value. This instruction reads each vector element from the source SIMD&FP register, puts the absolute value of the result into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are signed integer values.
//
//go:linkname VqabsdS64 VqabsdS64
//go:noescape
func VqabsdS64(r *arm.Int64, v0 *arm.Int64)

// Signed saturating Absolute value. This instruction reads each vector element from the source SIMD&FP register, puts the absolute value of the result into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are signed integer values.
//
//go:linkname VqabshS16 VqabshS16
//go:noescape
func VqabshS16(r *arm.Int16, v0 *arm.Int16)

// Unsigned saturating Add. This instruction adds the values of corresponding elements of the two source SIMD&FP registers, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VqaddbU8 VqaddbU8
//go:noescape
func VqaddbU8(r *arm.Uint8, v0 *arm.Uint8, v1 *arm.Uint8)

// Unsigned saturating Add. This instruction adds the values of corresponding elements of the two source SIMD&FP registers, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VqaddsU32 VqaddsU32
//go:noescape
func VqaddsU32(r *arm.Uint32, v0 *arm.Uint32, v1 *arm.Uint32)

// Unsigned saturating Add. This instruction adds the values of corresponding elements of the two source SIMD&FP registers, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VqadddU64 VqadddU64
//go:noescape
func VqadddU64(r *arm.Uint64, v0 *arm.Uint64, v1 *arm.Uint64)

// Unsigned saturating Add. This instruction adds the values of corresponding elements of the two source SIMD&FP registers, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VqaddhU16 VqaddhU16
//go:noescape
func VqaddhU16(r *arm.Uint16, v0 *arm.Uint16, v1 *arm.Uint16)

// Signed saturating Add. This instruction adds the values of corresponding elements of the two source SIMD&FP registers, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VqaddbS8 VqaddbS8
//go:noescape
func VqaddbS8(r *arm.Int8, v0 *arm.Int8, v1 *arm.Int8)

// Signed saturating Add. This instruction adds the values of corresponding elements of the two source SIMD&FP registers, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VqaddsS32 VqaddsS32
//go:noescape
func VqaddsS32(r *arm.Int32, v0 *arm.Int32, v1 *arm.Int32)

// Signed saturating Add. This instruction adds the values of corresponding elements of the two source SIMD&FP registers, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VqadddS64 VqadddS64
//go:noescape
func VqadddS64(r *arm.Int64, v0 *arm.Int64, v1 *arm.Int64)

// Signed saturating Add. This instruction adds the values of corresponding elements of the two source SIMD&FP registers, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VqaddhS16 VqaddhS16
//go:noescape
func VqaddhS16(r *arm.Int16, v0 *arm.Int16, v1 *arm.Int16)

// Signed saturating Doubling Multiply-Add Long. This instruction multiplies corresponding signed integer values in the lower or upper half of the vectors of the two source SIMD&FP registers, doubles the results, and accumulates the final results with the vector elements of the destination SIMD&FP register. The destination vector elements are twice as long as the elements that are multiplied.
//
//go:linkname VqdmlalsS32 VqdmlalsS32
//go:noescape
func VqdmlalsS32(r *arm.Int64, v0 *arm.Int64, v1 *arm.Int32, v2 *arm.Int32)

// Signed saturating Doubling Multiply-Add Long. This instruction multiplies corresponding signed integer values in the lower or upper half of the vectors of the two source SIMD&FP registers, doubles the results, and accumulates the final results with the vector elements of the destination SIMD&FP register. The destination vector elements are twice as long as the elements that are multiplied.
//
//go:linkname VqdmlalhS16 VqdmlalhS16
//go:noescape
func VqdmlalhS16(r *arm.Int32, v0 *arm.Int32, v1 *arm.Int16, v2 *arm.Int16)

// Signed saturating Doubling Multiply-Add Long. This instruction multiplies corresponding signed integer values in the lower or upper half of the vectors of the two source SIMD&FP registers, doubles the results, and accumulates the final results with the vector elements of the destination SIMD&FP register. The destination vector elements are twice as long as the elements that are multiplied.
//
//go:linkname VqdmlalHighS32 VqdmlalHighS32
//go:noescape
func VqdmlalHighS32(r *arm.Int64X2, v0 *arm.Int64X2, v1 *arm.Int32X4, v2 *arm.Int32X4)

// Signed saturating Doubling Multiply-Add Long. This instruction multiplies corresponding signed integer values in the lower or upper half of the vectors of the two source SIMD&FP registers, doubles the results, and accumulates the final results with the vector elements of the destination SIMD&FP register. The destination vector elements are twice as long as the elements that are multiplied.
//
//go:linkname VqdmlalHighS16 VqdmlalHighS16
//go:noescape
func VqdmlalHighS16(r *arm.Int32X4, v0 *arm.Int32X4, v1 *arm.Int16X8, v2 *arm.Int16X8)

// Signed saturating Doubling Multiply-Add Long. This instruction multiplies corresponding signed integer values in the lower or upper half of the vectors of the two source SIMD&FP registers, doubles the results, and accumulates the final results with the vector elements of the destination SIMD&FP register. The destination vector elements are twice as long as the elements that are multiplied.
//
//go:linkname VqdmlalHighNS32 VqdmlalHighNS32
//go:noescape
func VqdmlalHighNS32(r *arm.Int64X2, v0 *arm.Int64X2, v1 *arm.Int32X4, v2 *arm.Int32)

// Signed saturating Doubling Multiply-Add Long. This instruction multiplies corresponding signed integer values in the lower or upper half of the vectors of the two source SIMD&FP registers, doubles the results, and accumulates the final results with the vector elements of the destination SIMD&FP register. The destination vector elements are twice as long as the elements that are multiplied.
//
//go:linkname VqdmlalHighNS16 VqdmlalHighNS16
//go:noescape
func VqdmlalHighNS16(r *arm.Int32X4, v0 *arm.Int32X4, v1 *arm.Int16X8, v2 *arm.Int16)

// Signed saturating Doubling Multiply-Subtract Long. This instruction multiplies corresponding signed integer values in the lower or upper half of the vectors of the two source SIMD&FP registers, doubles the results, and subtracts the final results from the vector elements of the destination SIMD&FP register. The destination vector elements are twice as long as the elements that are multiplied.
//
//go:linkname VqdmlslsS32 VqdmlslsS32
//go:noescape
func VqdmlslsS32(r *arm.Int64, v0 *arm.Int64, v1 *arm.Int32, v2 *arm.Int32)

// Signed saturating Doubling Multiply-Subtract Long. This instruction multiplies corresponding signed integer values in the lower or upper half of the vectors of the two source SIMD&FP registers, doubles the results, and subtracts the final results from the vector elements of the destination SIMD&FP register. The destination vector elements are twice as long as the elements that are multiplied.
//
//go:linkname VqdmlslhS16 VqdmlslhS16
//go:noescape
func VqdmlslhS16(r *arm.Int32, v0 *arm.Int32, v1 *arm.Int16, v2 *arm.Int16)

// Signed saturating Doubling Multiply-Subtract Long. This instruction multiplies corresponding signed integer values in the lower or upper half of the vectors of the two source SIMD&FP registers, doubles the results, and subtracts the final results from the vector elements of the destination SIMD&FP register. The destination vector elements are twice as long as the elements that are multiplied.
//
//go:linkname VqdmlslHighS32 VqdmlslHighS32
//go:noescape
func VqdmlslHighS32(r *arm.Int64X2, v0 *arm.Int64X2, v1 *arm.Int32X4, v2 *arm.Int32X4)

// Signed saturating Doubling Multiply-Subtract Long. This instruction multiplies corresponding signed integer values in the lower or upper half of the vectors of the two source SIMD&FP registers, doubles the results, and subtracts the final results from the vector elements of the destination SIMD&FP register. The destination vector elements are twice as long as the elements that are multiplied.
//
//go:linkname VqdmlslHighS16 VqdmlslHighS16
//go:noescape
func VqdmlslHighS16(r *arm.Int32X4, v0 *arm.Int32X4, v1 *arm.Int16X8, v2 *arm.Int16X8)

// Signed saturating Doubling Multiply-Subtract Long. This instruction multiplies corresponding signed integer values in the lower or upper half of the vectors of the two source SIMD&FP registers, doubles the results, and subtracts the final results from the vector elements of the destination SIMD&FP register. The destination vector elements are twice as long as the elements that are multiplied.
//
//go:linkname VqdmlslHighNS32 VqdmlslHighNS32
//go:noescape
func VqdmlslHighNS32(r *arm.Int64X2, v0 *arm.Int64X2, v1 *arm.Int32X4, v2 *arm.Int32)

// Signed saturating Doubling Multiply-Subtract Long. This instruction multiplies corresponding signed integer values in the lower or upper half of the vectors of the two source SIMD&FP registers, doubles the results, and subtracts the final results from the vector elements of the destination SIMD&FP register. The destination vector elements are twice as long as the elements that are multiplied.
//
//go:linkname VqdmlslHighNS16 VqdmlslHighNS16
//go:noescape
func VqdmlslHighNS16(r *arm.Int32X4, v0 *arm.Int32X4, v1 *arm.Int16X8, v2 *arm.Int16)

// Signed saturating Doubling Multiply returning High half. This instruction multiplies the values of corresponding elements of the two source SIMD&FP registers, doubles the results, places the most significant half of the final results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VqdmulhsS32 VqdmulhsS32
//go:noescape
func VqdmulhsS32(r *arm.Int32, v0 *arm.Int32, v1 *arm.Int32)

// Signed saturating Doubling Multiply returning High half. This instruction multiplies the values of corresponding elements of the two source SIMD&FP registers, doubles the results, places the most significant half of the final results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VqdmulhhS16 VqdmulhhS16
//go:noescape
func VqdmulhhS16(r *arm.Int16, v0 *arm.Int16, v1 *arm.Int16)

// Signed saturating Doubling Multiply Long. This instruction multiplies corresponding vector elements in the lower or upper half of the two source SIMD&FP registers, doubles the results, places the final results in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VqdmullsS32 VqdmullsS32
//go:noescape
func VqdmullsS32(r *arm.Int64, v0 *arm.Int32, v1 *arm.Int32)

// Signed saturating Doubling Multiply Long. This instruction multiplies corresponding vector elements in the lower or upper half of the two source SIMD&FP registers, doubles the results, places the final results in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VqdmullhS16 VqdmullhS16
//go:noescape
func VqdmullhS16(r *arm.Int32, v0 *arm.Int16, v1 *arm.Int16)

// Signed saturating Doubling Multiply Long. This instruction multiplies corresponding vector elements in the lower or upper half of the two source SIMD&FP registers, doubles the results, places the final results in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VqdmullHighS32 VqdmullHighS32
//go:noescape
func VqdmullHighS32(r *arm.Int64X2, v0 *arm.Int32X4, v1 *arm.Int32X4)

// Signed saturating Doubling Multiply Long. This instruction multiplies corresponding vector elements in the lower or upper half of the two source SIMD&FP registers, doubles the results, places the final results in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VqdmullHighS16 VqdmullHighS16
//go:noescape
func VqdmullHighS16(r *arm.Int32X4, v0 *arm.Int16X8, v1 *arm.Int16X8)

// Signed saturating Doubling Multiply Long. This instruction multiplies corresponding vector elements in the lower or upper half of the two source SIMD&FP registers, doubles the results, places the final results in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VqdmullHighNS32 VqdmullHighNS32
//go:noescape
func VqdmullHighNS32(r *arm.Int64X2, v0 *arm.Int32X4, v1 *arm.Int32)

// Signed saturating Doubling Multiply Long. This instruction multiplies corresponding vector elements in the lower or upper half of the two source SIMD&FP registers, doubles the results, places the final results in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VqdmullHighNS16 VqdmullHighNS16
//go:noescape
func VqdmullHighNS16(r *arm.Int32X4, v0 *arm.Int16X8, v1 *arm.Int16)

// Signed saturating extract Narrow. This instruction reads each vector element from the source SIMD&FP register, saturates the value to half the original width, places the result into a vector, and writes the vector to the lower or upper half of the destination SIMD&FP register. The destination vector elements are half as long as the source vector elements. All the values in this instruction are signed integer values.
//
//go:linkname VqmovnsS32 VqmovnsS32
//go:noescape
func VqmovnsS32(r *arm.Int16, v0 *arm.Int32)

// Signed saturating extract Narrow. This instruction reads each vector element from the source SIMD&FP register, saturates the value to half the original width, places the result into a vector, and writes the vector to the lower or upper half of the destination SIMD&FP register. The destination vector elements are half as long as the source vector elements. All the values in this instruction are signed integer values.
//
//go:linkname VqmovndS64 VqmovndS64
//go:noescape
func VqmovndS64(r *arm.Int32, v0 *arm.Int64)

// Signed saturating extract Narrow. This instruction reads each vector element from the source SIMD&FP register, saturates the value to half the original width, places the result into a vector, and writes the vector to the lower or upper half of the destination SIMD&FP register. The destination vector elements are half as long as the source vector elements. All the values in this instruction are signed integer values.
//
//go:linkname VqmovnhS16 VqmovnhS16
//go:noescape
func VqmovnhS16(r *arm.Int8, v0 *arm.Int16)

// Unsigned saturating extract Narrow. This instruction reads each vector element from the source SIMD&FP register, saturates each value to half the original width, places the result into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are unsigned integer values.
//
//go:linkname VqmovnsU32 VqmovnsU32
//go:noescape
func VqmovnsU32(r *arm.Uint16, v0 *arm.Uint32)

// Unsigned saturating extract Narrow. This instruction reads each vector element from the source SIMD&FP register, saturates each value to half the original width, places the result into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are unsigned integer values.
//
//go:linkname VqmovndU64 VqmovndU64
//go:noescape
func VqmovndU64(r *arm.Uint32, v0 *arm.Uint64)

// Unsigned saturating extract Narrow. This instruction reads each vector element from the source SIMD&FP register, saturates each value to half the original width, places the result into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are unsigned integer values.
//
//go:linkname VqmovnhU16 VqmovnhU16
//go:noescape
func VqmovnhU16(r *arm.Uint8, v0 *arm.Uint16)

// Unsigned saturating extract Narrow. This instruction reads each vector element from the source SIMD&FP register, saturates each value to half the original width, places the result into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are unsigned integer values.
//
//go:linkname VqmovnHighU32 VqmovnHighU32
//go:noescape
func VqmovnHighU32(r *arm.Uint16X8, v0 *arm.Uint16X4, v1 *arm.Uint32X4)

// Unsigned saturating extract Narrow. This instruction reads each vector element from the source SIMD&FP register, saturates each value to half the original width, places the result into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are unsigned integer values.
//
//go:linkname VqmovnHighU64 VqmovnHighU64
//go:noescape
func VqmovnHighU64(r *arm.Uint32X4, v0 *arm.Uint32X2, v1 *arm.Uint64X2)

// Unsigned saturating extract Narrow. This instruction reads each vector element from the source SIMD&FP register, saturates each value to half the original width, places the result into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are unsigned integer values.
//
//go:linkname VqmovnHighU16 VqmovnHighU16
//go:noescape
func VqmovnHighU16(r *arm.Uint8X16, v0 *arm.Uint8X8, v1 *arm.Uint16X8)

// Signed saturating extract Narrow. This instruction reads each vector element from the source SIMD&FP register, saturates the value to half the original width, places the result into a vector, and writes the vector to the lower or upper half of the destination SIMD&FP register. The destination vector elements are half as long as the source vector elements. All the values in this instruction are signed integer values.
//
//go:linkname VqmovnHighS32 VqmovnHighS32
//go:noescape
func VqmovnHighS32(r *arm.Int16X8, v0 *arm.Int16X4, v1 *arm.Int32X4)

// Signed saturating extract Narrow. This instruction reads each vector element from the source SIMD&FP register, saturates the value to half the original width, places the result into a vector, and writes the vector to the lower or upper half of the destination SIMD&FP register. The destination vector elements are half as long as the source vector elements. All the values in this instruction are signed integer values.
//
//go:linkname VqmovnHighS64 VqmovnHighS64
//go:noescape
func VqmovnHighS64(r *arm.Int32X4, v0 *arm.Int32X2, v1 *arm.Int64X2)

// Signed saturating extract Narrow. This instruction reads each vector element from the source SIMD&FP register, saturates the value to half the original width, places the result into a vector, and writes the vector to the lower or upper half of the destination SIMD&FP register. The destination vector elements are half as long as the source vector elements. All the values in this instruction are signed integer values.
//
//go:linkname VqmovnHighS16 VqmovnHighS16
//go:noescape
func VqmovnHighS16(r *arm.Int8X16, v0 *arm.Int8X8, v1 *arm.Int16X8)

// Signed saturating extract Unsigned Narrow. This instruction reads each signed integer value in the vector of the source SIMD&FP register, saturates the value to an unsigned integer value that is half the original width, places the result into a vector, and writes the vector to the lower or upper half of the destination SIMD&FP register. The destination vector elements are half as long as the source vector elements.
//
//go:linkname VqmovunsS32 VqmovunsS32
//go:noescape
func VqmovunsS32(r *arm.Uint16, v0 *arm.Int32)

// Signed saturating extract Unsigned Narrow. This instruction reads each signed integer value in the vector of the source SIMD&FP register, saturates the value to an unsigned integer value that is half the original width, places the result into a vector, and writes the vector to the lower or upper half of the destination SIMD&FP register. The destination vector elements are half as long as the source vector elements.
//
//go:linkname VqmovundS64 VqmovundS64
//go:noescape
func VqmovundS64(r *arm.Uint32, v0 *arm.Int64)

// Signed saturating extract Unsigned Narrow. This instruction reads each signed integer value in the vector of the source SIMD&FP register, saturates the value to an unsigned integer value that is half the original width, places the result into a vector, and writes the vector to the lower or upper half of the destination SIMD&FP register. The destination vector elements are half as long as the source vector elements.
//
//go:linkname VqmovunhS16 VqmovunhS16
//go:noescape
func VqmovunhS16(r *arm.Uint8, v0 *arm.Int16)

// Signed saturating extract Unsigned Narrow. This instruction reads each signed integer value in the vector of the source SIMD&FP register, saturates the value to an unsigned integer value that is half the original width, places the result into a vector, and writes the vector to the lower or upper half of the destination SIMD&FP register. The destination vector elements are half as long as the source vector elements.
//
//go:linkname VqmovunHighS32 VqmovunHighS32
//go:noescape
func VqmovunHighS32(r *arm.Uint16X8, v0 *arm.Uint16X4, v1 *arm.Int32X4)

// Signed saturating extract Unsigned Narrow. This instruction reads each signed integer value in the vector of the source SIMD&FP register, saturates the value to an unsigned integer value that is half the original width, places the result into a vector, and writes the vector to the lower or upper half of the destination SIMD&FP register. The destination vector elements are half as long as the source vector elements.
//
//go:linkname VqmovunHighS64 VqmovunHighS64
//go:noescape
func VqmovunHighS64(r *arm.Uint32X4, v0 *arm.Uint32X2, v1 *arm.Int64X2)

// Signed saturating extract Unsigned Narrow. This instruction reads each signed integer value in the vector of the source SIMD&FP register, saturates the value to an unsigned integer value that is half the original width, places the result into a vector, and writes the vector to the lower or upper half of the destination SIMD&FP register. The destination vector elements are half as long as the source vector elements.
//
//go:linkname VqmovunHighS16 VqmovunHighS16
//go:noescape
func VqmovunHighS16(r *arm.Uint8X16, v0 *arm.Uint8X8, v1 *arm.Int16X8)

// Signed saturating Negate. This instruction reads each vector element from the source SIMD&FP register, negates each value, places the result into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are signed integer values.
//
//go:linkname VqnegqS64 VqnegqS64
//go:noescape
func VqnegqS64(r *arm.Int64X2, v0 *arm.Int64X2)

// Signed saturating Negate. This instruction reads each vector element from the source SIMD&FP register, negates each value, places the result into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are signed integer values.
//
//go:linkname VqnegS64 VqnegS64
//go:noescape
func VqnegS64(r *arm.Int64X1, v0 *arm.Int64X1)

// Signed saturating Negate. This instruction reads each vector element from the source SIMD&FP register, negates each value, places the result into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are signed integer values.
//
//go:linkname VqnegbS8 VqnegbS8
//go:noescape
func VqnegbS8(r *arm.Int8, v0 *arm.Int8)

// Signed saturating Negate. This instruction reads each vector element from the source SIMD&FP register, negates each value, places the result into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are signed integer values.
//
//go:linkname VqnegsS32 VqnegsS32
//go:noescape
func VqnegsS32(r *arm.Int32, v0 *arm.Int32)

// Signed saturating Negate. This instruction reads each vector element from the source SIMD&FP register, negates each value, places the result into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are signed integer values.
//
//go:linkname VqnegdS64 VqnegdS64
//go:noescape
func VqnegdS64(r *arm.Int64, v0 *arm.Int64)

// Signed saturating Negate. This instruction reads each vector element from the source SIMD&FP register, negates each value, places the result into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are signed integer values.
//
//go:linkname VqneghS16 VqneghS16
//go:noescape
func VqneghS16(r *arm.Int16, v0 *arm.Int16)

// Signed saturating Rounding Doubling Multiply returning High half. This instruction multiplies the values of corresponding elements of the two source SIMD&FP registers, doubles the results, places the most significant half of the final results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VqrdmulhsS32 VqrdmulhsS32
//go:noescape
func VqrdmulhsS32(r *arm.Int32, v0 *arm.Int32, v1 *arm.Int32)

// Signed saturating Rounding Doubling Multiply returning High half. This instruction multiplies the values of corresponding elements of the two source SIMD&FP registers, doubles the results, places the most significant half of the final results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VqrdmulhhS16 VqrdmulhhS16
//go:noescape
func VqrdmulhhS16(r *arm.Int16, v0 *arm.Int16, v1 *arm.Int16)

// Unsigned saturating Rounding Shift Left (register). This instruction takes each vector element of the first source SIMD&FP register, shifts the vector element by a value from the least significant byte of the corresponding vector element of the second source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VqrshlbU8 VqrshlbU8
//go:noescape
func VqrshlbU8(r *arm.Uint8, v0 *arm.Uint8, v1 *arm.Int8)

// Unsigned saturating Rounding Shift Left (register). This instruction takes each vector element of the first source SIMD&FP register, shifts the vector element by a value from the least significant byte of the corresponding vector element of the second source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VqrshlsU32 VqrshlsU32
//go:noescape
func VqrshlsU32(r *arm.Uint32, v0 *arm.Uint32, v1 *arm.Int32)

// Unsigned saturating Rounding Shift Left (register). This instruction takes each vector element of the first source SIMD&FP register, shifts the vector element by a value from the least significant byte of the corresponding vector element of the second source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VqrshldU64 VqrshldU64
//go:noescape
func VqrshldU64(r *arm.Uint64, v0 *arm.Uint64, v1 *arm.Int64)

// Unsigned saturating Rounding Shift Left (register). This instruction takes each vector element of the first source SIMD&FP register, shifts the vector element by a value from the least significant byte of the corresponding vector element of the second source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VqrshlhU16 VqrshlhU16
//go:noescape
func VqrshlhU16(r *arm.Uint16, v0 *arm.Uint16, v1 *arm.Int16)

// Signed saturating Rounding Shift Left (register). This instruction takes each vector element in the first source SIMD&FP register, shifts it by a value from the least significant byte of the corresponding vector element of the second source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VqrshlbS8 VqrshlbS8
//go:noescape
func VqrshlbS8(r *arm.Int8, v0 *arm.Int8, v1 *arm.Int8)

// Signed saturating Rounding Shift Left (register). This instruction takes each vector element in the first source SIMD&FP register, shifts it by a value from the least significant byte of the corresponding vector element of the second source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VqrshlsS32 VqrshlsS32
//go:noescape
func VqrshlsS32(r *arm.Int32, v0 *arm.Int32, v1 *arm.Int32)

// Signed saturating Rounding Shift Left (register). This instruction takes each vector element in the first source SIMD&FP register, shifts it by a value from the least significant byte of the corresponding vector element of the second source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VqrshldS64 VqrshldS64
//go:noescape
func VqrshldS64(r *arm.Int64, v0 *arm.Int64, v1 *arm.Int64)

// Signed saturating Rounding Shift Left (register). This instruction takes each vector element in the first source SIMD&FP register, shifts it by a value from the least significant byte of the corresponding vector element of the second source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VqrshlhS16 VqrshlhS16
//go:noescape
func VqrshlhS16(r *arm.Int16, v0 *arm.Int16, v1 *arm.Int16)

// Unsigned saturating Shift Left (register). This instruction takes each element in the vector of the first source SIMD&FP register, shifts the element by a value from the least significant byte of the corresponding element of the second source SIMD&FP register, places the results in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VqshlbU8 VqshlbU8
//go:noescape
func VqshlbU8(r *arm.Uint8, v0 *arm.Uint8, v1 *arm.Int8)

// Unsigned saturating Shift Left (register). This instruction takes each element in the vector of the first source SIMD&FP register, shifts the element by a value from the least significant byte of the corresponding element of the second source SIMD&FP register, places the results in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VqshlsU32 VqshlsU32
//go:noescape
func VqshlsU32(r *arm.Uint32, v0 *arm.Uint32, v1 *arm.Int32)

// Unsigned saturating Shift Left (register). This instruction takes each element in the vector of the first source SIMD&FP register, shifts the element by a value from the least significant byte of the corresponding element of the second source SIMD&FP register, places the results in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VqshldU64 VqshldU64
//go:noescape
func VqshldU64(r *arm.Uint64, v0 *arm.Uint64, v1 *arm.Int64)

// Unsigned saturating Shift Left (register). This instruction takes each element in the vector of the first source SIMD&FP register, shifts the element by a value from the least significant byte of the corresponding element of the second source SIMD&FP register, places the results in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VqshlhU16 VqshlhU16
//go:noescape
func VqshlhU16(r *arm.Uint16, v0 *arm.Uint16, v1 *arm.Int16)

// Signed saturating Shift Left (register). This instruction takes each element in the vector of the first source SIMD&FP register, shifts each element by a value from the least significant byte of the corresponding element of the second source SIMD&FP register, places the results in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VqshlbS8 VqshlbS8
//go:noescape
func VqshlbS8(r *arm.Int8, v0 *arm.Int8, v1 *arm.Int8)

// Signed saturating Shift Left (register). This instruction takes each element in the vector of the first source SIMD&FP register, shifts each element by a value from the least significant byte of the corresponding element of the second source SIMD&FP register, places the results in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VqshlsS32 VqshlsS32
//go:noescape
func VqshlsS32(r *arm.Int32, v0 *arm.Int32, v1 *arm.Int32)

// Signed saturating Shift Left (register). This instruction takes each element in the vector of the first source SIMD&FP register, shifts each element by a value from the least significant byte of the corresponding element of the second source SIMD&FP register, places the results in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VqshldS64 VqshldS64
//go:noescape
func VqshldS64(r *arm.Int64, v0 *arm.Int64, v1 *arm.Int64)

// Signed saturating Shift Left (register). This instruction takes each element in the vector of the first source SIMD&FP register, shifts each element by a value from the least significant byte of the corresponding element of the second source SIMD&FP register, places the results in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VqshlhS16 VqshlhS16
//go:noescape
func VqshlhS16(r *arm.Int16, v0 *arm.Int16, v1 *arm.Int16)

// Unsigned saturating Subtract. This instruction subtracts the element values of the second source SIMD&FP register from the corresponding element values of the first source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VqsubbU8 VqsubbU8
//go:noescape
func VqsubbU8(r *arm.Uint8, v0 *arm.Uint8, v1 *arm.Uint8)

// Unsigned saturating Subtract. This instruction subtracts the element values of the second source SIMD&FP register from the corresponding element values of the first source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VqsubsU32 VqsubsU32
//go:noescape
func VqsubsU32(r *arm.Uint32, v0 *arm.Uint32, v1 *arm.Uint32)

// Unsigned saturating Subtract. This instruction subtracts the element values of the second source SIMD&FP register from the corresponding element values of the first source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VqsubdU64 VqsubdU64
//go:noescape
func VqsubdU64(r *arm.Uint64, v0 *arm.Uint64, v1 *arm.Uint64)

// Unsigned saturating Subtract. This instruction subtracts the element values of the second source SIMD&FP register from the corresponding element values of the first source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VqsubhU16 VqsubhU16
//go:noescape
func VqsubhU16(r *arm.Uint16, v0 *arm.Uint16, v1 *arm.Uint16)

// Signed saturating Subtract. This instruction subtracts the element values of the second source SIMD&FP register from the corresponding element values of the first source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VqsubbS8 VqsubbS8
//go:noescape
func VqsubbS8(r *arm.Int8, v0 *arm.Int8, v1 *arm.Int8)

// Signed saturating Subtract. This instruction subtracts the element values of the second source SIMD&FP register from the corresponding element values of the first source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VqsubsS32 VqsubsS32
//go:noescape
func VqsubsS32(r *arm.Int32, v0 *arm.Int32, v1 *arm.Int32)

// Signed saturating Subtract. This instruction subtracts the element values of the second source SIMD&FP register from the corresponding element values of the first source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VqsubdS64 VqsubdS64
//go:noescape
func VqsubdS64(r *arm.Int64, v0 *arm.Int64, v1 *arm.Int64)

// Signed saturating Subtract. This instruction subtracts the element values of the second source SIMD&FP register from the corresponding element values of the first source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VqsubhS16 VqsubhS16
//go:noescape
func VqsubhS16(r *arm.Int16, v0 *arm.Int16, v1 *arm.Int16)

// Table vector Lookup. This instruction reads each value from the vector elements in the index source SIMD&FP register, uses each result as an index to perform a lookup in a table of bytes that is described by one to four source table SIMD&FP registers, places the lookup result in a vector, and writes the vector to the destination SIMD&FP register. If an index is out of range for the table, the result for that lookup is 0. If more than one source register is used to describe the table, the first source register describes the lowest bytes of the table.
//
//go:linkname Vqtbl1P8 Vqtbl1P8
//go:noescape
func Vqtbl1P8(r *arm.Poly8X8, v0 *arm.Poly8X16, v1 *arm.Uint8X8)

// Table vector Lookup. This instruction reads each value from the vector elements in the index source SIMD&FP register, uses each result as an index to perform a lookup in a table of bytes that is described by one to four source table SIMD&FP registers, places the lookup result in a vector, and writes the vector to the destination SIMD&FP register. If an index is out of range for the table, the result for that lookup is 0. If more than one source register is used to describe the table, the first source register describes the lowest bytes of the table.
//
//go:linkname Vqtbl1QP8 Vqtbl1QP8
//go:noescape
func Vqtbl1QP8(r *arm.Poly8X16, v0 *arm.Poly8X16, v1 *arm.Uint8X16)

// Table vector Lookup. This instruction reads each value from the vector elements in the index source SIMD&FP register, uses each result as an index to perform a lookup in a table of bytes that is described by one to four source table SIMD&FP registers, places the lookup result in a vector, and writes the vector to the destination SIMD&FP register. If an index is out of range for the table, the result for that lookup is 0. If more than one source register is used to describe the table, the first source register describes the lowest bytes of the table.
//
//go:linkname Vqtbl1QU8 Vqtbl1QU8
//go:noescape
func Vqtbl1QU8(r *arm.Uint8X16, v0 *arm.Uint8X16, v1 *arm.Uint8X16)

// Table vector Lookup. This instruction reads each value from the vector elements in the index source SIMD&FP register, uses each result as an index to perform a lookup in a table of bytes that is described by one to four source table SIMD&FP registers, places the lookup result in a vector, and writes the vector to the destination SIMD&FP register. If an index is out of range for the table, the result for that lookup is 0. If more than one source register is used to describe the table, the first source register describes the lowest bytes of the table.
//
//go:linkname Vqtbl1QS8 Vqtbl1QS8
//go:noescape
func Vqtbl1QS8(r *arm.Int8X16, v0 *arm.Int8X16, v1 *arm.Uint8X16)

// Table vector Lookup. This instruction reads each value from the vector elements in the index source SIMD&FP register, uses each result as an index to perform a lookup in a table of bytes that is described by one to four source table SIMD&FP registers, places the lookup result in a vector, and writes the vector to the destination SIMD&FP register. If an index is out of range for the table, the result for that lookup is 0. If more than one source register is used to describe the table, the first source register describes the lowest bytes of the table.
//
//go:linkname Vqtbl1U8 Vqtbl1U8
//go:noescape
func Vqtbl1U8(r *arm.Uint8X8, v0 *arm.Uint8X16, v1 *arm.Uint8X8)

// Table vector Lookup. This instruction reads each value from the vector elements in the index source SIMD&FP register, uses each result as an index to perform a lookup in a table of bytes that is described by one to four source table SIMD&FP registers, places the lookup result in a vector, and writes the vector to the destination SIMD&FP register. If an index is out of range for the table, the result for that lookup is 0. If more than one source register is used to describe the table, the first source register describes the lowest bytes of the table.
//
//go:linkname Vqtbl1S8 Vqtbl1S8
//go:noescape
func Vqtbl1S8(r *arm.Int8X8, v0 *arm.Int8X16, v1 *arm.Uint8X8)

// Table vector Lookup. This instruction reads each value from the vector elements in the index source SIMD&FP register, uses each result as an index to perform a lookup in a table of bytes that is described by one to four source table SIMD&FP registers, places the lookup result in a vector, and writes the vector to the destination SIMD&FP register. If an index is out of range for the table, the result for that lookup is 0. If more than one source register is used to describe the table, the first source register describes the lowest bytes of the table.
//
//go:linkname Vqtbl2P8 Vqtbl2P8
//go:noescape
func Vqtbl2P8(r *arm.Poly8X8, v0 *arm.Poly8X16X2, v1 *arm.Uint8X8)

// Table vector Lookup. This instruction reads each value from the vector elements in the index source SIMD&FP register, uses each result as an index to perform a lookup in a table of bytes that is described by one to four source table SIMD&FP registers, places the lookup result in a vector, and writes the vector to the destination SIMD&FP register. If an index is out of range for the table, the result for that lookup is 0. If more than one source register is used to describe the table, the first source register describes the lowest bytes of the table.
//
//go:linkname Vqtbl2QP8 Vqtbl2QP8
//go:noescape
func Vqtbl2QP8(r *arm.Poly8X16, v0 *arm.Poly8X16X2, v1 *arm.Uint8X16)

// Table vector Lookup. This instruction reads each value from the vector elements in the index source SIMD&FP register, uses each result as an index to perform a lookup in a table of bytes that is described by one to four source table SIMD&FP registers, places the lookup result in a vector, and writes the vector to the destination SIMD&FP register. If an index is out of range for the table, the result for that lookup is 0. If more than one source register is used to describe the table, the first source register describes the lowest bytes of the table.
//
//go:linkname Vqtbl2QU8 Vqtbl2QU8
//go:noescape
func Vqtbl2QU8(r *arm.Uint8X16, v0 *arm.Uint8X16X2, v1 *arm.Uint8X16)

// Table vector Lookup. This instruction reads each value from the vector elements in the index source SIMD&FP register, uses each result as an index to perform a lookup in a table of bytes that is described by one to four source table SIMD&FP registers, places the lookup result in a vector, and writes the vector to the destination SIMD&FP register. If an index is out of range for the table, the result for that lookup is 0. If more than one source register is used to describe the table, the first source register describes the lowest bytes of the table.
//
//go:linkname Vqtbl2QS8 Vqtbl2QS8
//go:noescape
func Vqtbl2QS8(r *arm.Int8X16, v0 *arm.Int8X16X2, v1 *arm.Uint8X16)

// Table vector Lookup. This instruction reads each value from the vector elements in the index source SIMD&FP register, uses each result as an index to perform a lookup in a table of bytes that is described by one to four source table SIMD&FP registers, places the lookup result in a vector, and writes the vector to the destination SIMD&FP register. If an index is out of range for the table, the result for that lookup is 0. If more than one source register is used to describe the table, the first source register describes the lowest bytes of the table.
//
//go:linkname Vqtbl2U8 Vqtbl2U8
//go:noescape
func Vqtbl2U8(r *arm.Uint8X8, v0 *arm.Uint8X16X2, v1 *arm.Uint8X8)

// Table vector Lookup. This instruction reads each value from the vector elements in the index source SIMD&FP register, uses each result as an index to perform a lookup in a table of bytes that is described by one to four source table SIMD&FP registers, places the lookup result in a vector, and writes the vector to the destination SIMD&FP register. If an index is out of range for the table, the result for that lookup is 0. If more than one source register is used to describe the table, the first source register describes the lowest bytes of the table.
//
//go:linkname Vqtbl2S8 Vqtbl2S8
//go:noescape
func Vqtbl2S8(r *arm.Int8X8, v0 *arm.Int8X16X2, v1 *arm.Uint8X8)

// Table vector Lookup. This instruction reads each value from the vector elements in the index source SIMD&FP register, uses each result as an index to perform a lookup in a table of bytes that is described by one to four source table SIMD&FP registers, places the lookup result in a vector, and writes the vector to the destination SIMD&FP register. If an index is out of range for the table, the result for that lookup is 0. If more than one source register is used to describe the table, the first source register describes the lowest bytes of the table.
//
//go:linkname Vqtbl3P8 Vqtbl3P8
//go:noescape
func Vqtbl3P8(r *arm.Poly8X8, v0 *arm.Poly8X16X3, v1 *arm.Uint8X8)

// Table vector Lookup. This instruction reads each value from the vector elements in the index source SIMD&FP register, uses each result as an index to perform a lookup in a table of bytes that is described by one to four source table SIMD&FP registers, places the lookup result in a vector, and writes the vector to the destination SIMD&FP register. If an index is out of range for the table, the result for that lookup is 0. If more than one source register is used to describe the table, the first source register describes the lowest bytes of the table.
//
//go:linkname Vqtbl3QP8 Vqtbl3QP8
//go:noescape
func Vqtbl3QP8(r *arm.Poly8X16, v0 *arm.Poly8X16X3, v1 *arm.Uint8X16)

// Table vector Lookup. This instruction reads each value from the vector elements in the index source SIMD&FP register, uses each result as an index to perform a lookup in a table of bytes that is described by one to four source table SIMD&FP registers, places the lookup result in a vector, and writes the vector to the destination SIMD&FP register. If an index is out of range for the table, the result for that lookup is 0. If more than one source register is used to describe the table, the first source register describes the lowest bytes of the table.
//
//go:linkname Vqtbl3QU8 Vqtbl3QU8
//go:noescape
func Vqtbl3QU8(r *arm.Uint8X16, v0 *arm.Uint8X16X3, v1 *arm.Uint8X16)

// Table vector Lookup. This instruction reads each value from the vector elements in the index source SIMD&FP register, uses each result as an index to perform a lookup in a table of bytes that is described by one to four source table SIMD&FP registers, places the lookup result in a vector, and writes the vector to the destination SIMD&FP register. If an index is out of range for the table, the result for that lookup is 0. If more than one source register is used to describe the table, the first source register describes the lowest bytes of the table.
//
//go:linkname Vqtbl3QS8 Vqtbl3QS8
//go:noescape
func Vqtbl3QS8(r *arm.Int8X16, v0 *arm.Int8X16X3, v1 *arm.Uint8X16)

// Table vector Lookup. This instruction reads each value from the vector elements in the index source SIMD&FP register, uses each result as an index to perform a lookup in a table of bytes that is described by one to four source table SIMD&FP registers, places the lookup result in a vector, and writes the vector to the destination SIMD&FP register. If an index is out of range for the table, the result for that lookup is 0. If more than one source register is used to describe the table, the first source register describes the lowest bytes of the table.
//
//go:linkname Vqtbl3U8 Vqtbl3U8
//go:noescape
func Vqtbl3U8(r *arm.Uint8X8, v0 *arm.Uint8X16X3, v1 *arm.Uint8X8)

// Table vector Lookup. This instruction reads each value from the vector elements in the index source SIMD&FP register, uses each result as an index to perform a lookup in a table of bytes that is described by one to four source table SIMD&FP registers, places the lookup result in a vector, and writes the vector to the destination SIMD&FP register. If an index is out of range for the table, the result for that lookup is 0. If more than one source register is used to describe the table, the first source register describes the lowest bytes of the table.
//
//go:linkname Vqtbl3S8 Vqtbl3S8
//go:noescape
func Vqtbl3S8(r *arm.Int8X8, v0 *arm.Int8X16X3, v1 *arm.Uint8X8)

// Table vector Lookup. This instruction reads each value from the vector elements in the index source SIMD&FP register, uses each result as an index to perform a lookup in a table of bytes that is described by one to four source table SIMD&FP registers, places the lookup result in a vector, and writes the vector to the destination SIMD&FP register. If an index is out of range for the table, the result for that lookup is 0. If more than one source register is used to describe the table, the first source register describes the lowest bytes of the table.
//
//go:linkname Vqtbl4P8 Vqtbl4P8
//go:noescape
func Vqtbl4P8(r *arm.Poly8X8, v0 *arm.Poly8X16X4, v1 *arm.Uint8X8)

// Table vector Lookup. This instruction reads each value from the vector elements in the index source SIMD&FP register, uses each result as an index to perform a lookup in a table of bytes that is described by one to four source table SIMD&FP registers, places the lookup result in a vector, and writes the vector to the destination SIMD&FP register. If an index is out of range for the table, the result for that lookup is 0. If more than one source register is used to describe the table, the first source register describes the lowest bytes of the table.
//
//go:linkname Vqtbl4QP8 Vqtbl4QP8
//go:noescape
func Vqtbl4QP8(r *arm.Poly8X16, v0 *arm.Poly8X16X4, v1 *arm.Uint8X16)

// Table vector Lookup. This instruction reads each value from the vector elements in the index source SIMD&FP register, uses each result as an index to perform a lookup in a table of bytes that is described by one to four source table SIMD&FP registers, places the lookup result in a vector, and writes the vector to the destination SIMD&FP register. If an index is out of range for the table, the result for that lookup is 0. If more than one source register is used to describe the table, the first source register describes the lowest bytes of the table.
//
//go:linkname Vqtbl4QU8 Vqtbl4QU8
//go:noescape
func Vqtbl4QU8(r *arm.Uint8X16, v0 *arm.Uint8X16X4, v1 *arm.Uint8X16)

// Table vector Lookup. This instruction reads each value from the vector elements in the index source SIMD&FP register, uses each result as an index to perform a lookup in a table of bytes that is described by one to four source table SIMD&FP registers, places the lookup result in a vector, and writes the vector to the destination SIMD&FP register. If an index is out of range for the table, the result for that lookup is 0. If more than one source register is used to describe the table, the first source register describes the lowest bytes of the table.
//
//go:linkname Vqtbl4QS8 Vqtbl4QS8
//go:noescape
func Vqtbl4QS8(r *arm.Int8X16, v0 *arm.Int8X16X4, v1 *arm.Uint8X16)

// Table vector Lookup. This instruction reads each value from the vector elements in the index source SIMD&FP register, uses each result as an index to perform a lookup in a table of bytes that is described by one to four source table SIMD&FP registers, places the lookup result in a vector, and writes the vector to the destination SIMD&FP register. If an index is out of range for the table, the result for that lookup is 0. If more than one source register is used to describe the table, the first source register describes the lowest bytes of the table.
//
//go:linkname Vqtbl4U8 Vqtbl4U8
//go:noescape
func Vqtbl4U8(r *arm.Uint8X8, v0 *arm.Uint8X16X4, v1 *arm.Uint8X8)

// Table vector Lookup. This instruction reads each value from the vector elements in the index source SIMD&FP register, uses each result as an index to perform a lookup in a table of bytes that is described by one to four source table SIMD&FP registers, places the lookup result in a vector, and writes the vector to the destination SIMD&FP register. If an index is out of range for the table, the result for that lookup is 0. If more than one source register is used to describe the table, the first source register describes the lowest bytes of the table.
//
//go:linkname Vqtbl4S8 Vqtbl4S8
//go:noescape
func Vqtbl4S8(r *arm.Int8X8, v0 *arm.Int8X16X4, v1 *arm.Uint8X8)

// Table vector lookup extension. This instruction reads each value from the vector elements in the index source SIMD&FP register, uses each result as an index to perform a lookup in a table of bytes that is described by one to four source table SIMD&FP registers, places the lookup result in a vector, and writes the vector to the destination SIMD&FP register. If an index is out of range for the table, the existing value in the vector element of the destination register is left unchanged. If more than one source register is used to describe the table, the first source register describes the lowest bytes of the table.
//
//go:linkname Vqtbx1P8 Vqtbx1P8
//go:noescape
func Vqtbx1P8(r *arm.Poly8X8, v0 *arm.Poly8X8, v1 *arm.Poly8X16, v2 *arm.Uint8X8)

// Table vector lookup extension. This instruction reads each value from the vector elements in the index source SIMD&FP register, uses each result as an index to perform a lookup in a table of bytes that is described by one to four source table SIMD&FP registers, places the lookup result in a vector, and writes the vector to the destination SIMD&FP register. If an index is out of range for the table, the existing value in the vector element of the destination register is left unchanged. If more than one source register is used to describe the table, the first source register describes the lowest bytes of the table.
//
//go:linkname Vqtbx1QP8 Vqtbx1QP8
//go:noescape
func Vqtbx1QP8(r *arm.Poly8X16, v0 *arm.Poly8X16, v1 *arm.Poly8X16, v2 *arm.Uint8X16)

// Table vector lookup extension. This instruction reads each value from the vector elements in the index source SIMD&FP register, uses each result as an index to perform a lookup in a table of bytes that is described by one to four source table SIMD&FP registers, places the lookup result in a vector, and writes the vector to the destination SIMD&FP register. If an index is out of range for the table, the existing value in the vector element of the destination register is left unchanged. If more than one source register is used to describe the table, the first source register describes the lowest bytes of the table.
//
//go:linkname Vqtbx1QU8 Vqtbx1QU8
//go:noescape
func Vqtbx1QU8(r *arm.Uint8X16, v0 *arm.Uint8X16, v1 *arm.Uint8X16, v2 *arm.Uint8X16)

// Table vector lookup extension. This instruction reads each value from the vector elements in the index source SIMD&FP register, uses each result as an index to perform a lookup in a table of bytes that is described by one to four source table SIMD&FP registers, places the lookup result in a vector, and writes the vector to the destination SIMD&FP register. If an index is out of range for the table, the existing value in the vector element of the destination register is left unchanged. If more than one source register is used to describe the table, the first source register describes the lowest bytes of the table.
//
//go:linkname Vqtbx1QS8 Vqtbx1QS8
//go:noescape
func Vqtbx1QS8(r *arm.Int8X16, v0 *arm.Int8X16, v1 *arm.Int8X16, v2 *arm.Uint8X16)

// Table vector lookup extension. This instruction reads each value from the vector elements in the index source SIMD&FP register, uses each result as an index to perform a lookup in a table of bytes that is described by one to four source table SIMD&FP registers, places the lookup result in a vector, and writes the vector to the destination SIMD&FP register. If an index is out of range for the table, the existing value in the vector element of the destination register is left unchanged. If more than one source register is used to describe the table, the first source register describes the lowest bytes of the table.
//
//go:linkname Vqtbx1U8 Vqtbx1U8
//go:noescape
func Vqtbx1U8(r *arm.Uint8X8, v0 *arm.Uint8X8, v1 *arm.Uint8X16, v2 *arm.Uint8X8)

// Table vector lookup extension. This instruction reads each value from the vector elements in the index source SIMD&FP register, uses each result as an index to perform a lookup in a table of bytes that is described by one to four source table SIMD&FP registers, places the lookup result in a vector, and writes the vector to the destination SIMD&FP register. If an index is out of range for the table, the existing value in the vector element of the destination register is left unchanged. If more than one source register is used to describe the table, the first source register describes the lowest bytes of the table.
//
//go:linkname Vqtbx1S8 Vqtbx1S8
//go:noescape
func Vqtbx1S8(r *arm.Int8X8, v0 *arm.Int8X8, v1 *arm.Int8X16, v2 *arm.Uint8X8)

// Table vector lookup extension. This instruction reads each value from the vector elements in the index source SIMD&FP register, uses each result as an index to perform a lookup in a table of bytes that is described by one to four source table SIMD&FP registers, places the lookup result in a vector, and writes the vector to the destination SIMD&FP register. If an index is out of range for the table, the existing value in the vector element of the destination register is left unchanged. If more than one source register is used to describe the table, the first source register describes the lowest bytes of the table.
//
//go:linkname Vqtbx2P8 Vqtbx2P8
//go:noescape
func Vqtbx2P8(r *arm.Poly8X8, v0 *arm.Poly8X8, v1 *arm.Poly8X16X2, v2 *arm.Uint8X8)

// Table vector lookup extension. This instruction reads each value from the vector elements in the index source SIMD&FP register, uses each result as an index to perform a lookup in a table of bytes that is described by one to four source table SIMD&FP registers, places the lookup result in a vector, and writes the vector to the destination SIMD&FP register. If an index is out of range for the table, the existing value in the vector element of the destination register is left unchanged. If more than one source register is used to describe the table, the first source register describes the lowest bytes of the table.
//
//go:linkname Vqtbx2QP8 Vqtbx2QP8
//go:noescape
func Vqtbx2QP8(r *arm.Poly8X16, v0 *arm.Poly8X16, v1 *arm.Poly8X16X2, v2 *arm.Uint8X16)

// Table vector lookup extension. This instruction reads each value from the vector elements in the index source SIMD&FP register, uses each result as an index to perform a lookup in a table of bytes that is described by one to four source table SIMD&FP registers, places the lookup result in a vector, and writes the vector to the destination SIMD&FP register. If an index is out of range for the table, the existing value in the vector element of the destination register is left unchanged. If more than one source register is used to describe the table, the first source register describes the lowest bytes of the table.
//
//go:linkname Vqtbx2QU8 Vqtbx2QU8
//go:noescape
func Vqtbx2QU8(r *arm.Uint8X16, v0 *arm.Uint8X16, v1 *arm.Uint8X16X2, v2 *arm.Uint8X16)

// Table vector lookup extension. This instruction reads each value from the vector elements in the index source SIMD&FP register, uses each result as an index to perform a lookup in a table of bytes that is described by one to four source table SIMD&FP registers, places the lookup result in a vector, and writes the vector to the destination SIMD&FP register. If an index is out of range for the table, the existing value in the vector element of the destination register is left unchanged. If more than one source register is used to describe the table, the first source register describes the lowest bytes of the table.
//
//go:linkname Vqtbx2QS8 Vqtbx2QS8
//go:noescape
func Vqtbx2QS8(r *arm.Int8X16, v0 *arm.Int8X16, v1 *arm.Int8X16X2, v2 *arm.Uint8X16)

// Table vector lookup extension. This instruction reads each value from the vector elements in the index source SIMD&FP register, uses each result as an index to perform a lookup in a table of bytes that is described by one to four source table SIMD&FP registers, places the lookup result in a vector, and writes the vector to the destination SIMD&FP register. If an index is out of range for the table, the existing value in the vector element of the destination register is left unchanged. If more than one source register is used to describe the table, the first source register describes the lowest bytes of the table.
//
//go:linkname Vqtbx2U8 Vqtbx2U8
//go:noescape
func Vqtbx2U8(r *arm.Uint8X8, v0 *arm.Uint8X8, v1 *arm.Uint8X16X2, v2 *arm.Uint8X8)

// Table vector lookup extension. This instruction reads each value from the vector elements in the index source SIMD&FP register, uses each result as an index to perform a lookup in a table of bytes that is described by one to four source table SIMD&FP registers, places the lookup result in a vector, and writes the vector to the destination SIMD&FP register. If an index is out of range for the table, the existing value in the vector element of the destination register is left unchanged. If more than one source register is used to describe the table, the first source register describes the lowest bytes of the table.
//
//go:linkname Vqtbx2S8 Vqtbx2S8
//go:noescape
func Vqtbx2S8(r *arm.Int8X8, v0 *arm.Int8X8, v1 *arm.Int8X16X2, v2 *arm.Uint8X8)

// Table vector lookup extension. This instruction reads each value from the vector elements in the index source SIMD&FP register, uses each result as an index to perform a lookup in a table of bytes that is described by one to four source table SIMD&FP registers, places the lookup result in a vector, and writes the vector to the destination SIMD&FP register. If an index is out of range for the table, the existing value in the vector element of the destination register is left unchanged. If more than one source register is used to describe the table, the first source register describes the lowest bytes of the table.
//
//go:linkname Vqtbx3P8 Vqtbx3P8
//go:noescape
func Vqtbx3P8(r *arm.Poly8X8, v0 *arm.Poly8X8, v1 *arm.Poly8X16X3, v2 *arm.Uint8X8)

// Table vector lookup extension. This instruction reads each value from the vector elements in the index source SIMD&FP register, uses each result as an index to perform a lookup in a table of bytes that is described by one to four source table SIMD&FP registers, places the lookup result in a vector, and writes the vector to the destination SIMD&FP register. If an index is out of range for the table, the existing value in the vector element of the destination register is left unchanged. If more than one source register is used to describe the table, the first source register describes the lowest bytes of the table.
//
//go:linkname Vqtbx3QP8 Vqtbx3QP8
//go:noescape
func Vqtbx3QP8(r *arm.Poly8X16, v0 *arm.Poly8X16, v1 *arm.Poly8X16X3, v2 *arm.Uint8X16)

// Table vector lookup extension. This instruction reads each value from the vector elements in the index source SIMD&FP register, uses each result as an index to perform a lookup in a table of bytes that is described by one to four source table SIMD&FP registers, places the lookup result in a vector, and writes the vector to the destination SIMD&FP register. If an index is out of range for the table, the existing value in the vector element of the destination register is left unchanged. If more than one source register is used to describe the table, the first source register describes the lowest bytes of the table.
//
//go:linkname Vqtbx3QU8 Vqtbx3QU8
//go:noescape
func Vqtbx3QU8(r *arm.Uint8X16, v0 *arm.Uint8X16, v1 *arm.Uint8X16X3, v2 *arm.Uint8X16)

// Table vector lookup extension. This instruction reads each value from the vector elements in the index source SIMD&FP register, uses each result as an index to perform a lookup in a table of bytes that is described by one to four source table SIMD&FP registers, places the lookup result in a vector, and writes the vector to the destination SIMD&FP register. If an index is out of range for the table, the existing value in the vector element of the destination register is left unchanged. If more than one source register is used to describe the table, the first source register describes the lowest bytes of the table.
//
//go:linkname Vqtbx3QS8 Vqtbx3QS8
//go:noescape
func Vqtbx3QS8(r *arm.Int8X16, v0 *arm.Int8X16, v1 *arm.Int8X16X3, v2 *arm.Uint8X16)

// Table vector lookup extension. This instruction reads each value from the vector elements in the index source SIMD&FP register, uses each result as an index to perform a lookup in a table of bytes that is described by one to four source table SIMD&FP registers, places the lookup result in a vector, and writes the vector to the destination SIMD&FP register. If an index is out of range for the table, the existing value in the vector element of the destination register is left unchanged. If more than one source register is used to describe the table, the first source register describes the lowest bytes of the table.
//
//go:linkname Vqtbx3U8 Vqtbx3U8
//go:noescape
func Vqtbx3U8(r *arm.Uint8X8, v0 *arm.Uint8X8, v1 *arm.Uint8X16X3, v2 *arm.Uint8X8)

// Table vector lookup extension. This instruction reads each value from the vector elements in the index source SIMD&FP register, uses each result as an index to perform a lookup in a table of bytes that is described by one to four source table SIMD&FP registers, places the lookup result in a vector, and writes the vector to the destination SIMD&FP register. If an index is out of range for the table, the existing value in the vector element of the destination register is left unchanged. If more than one source register is used to describe the table, the first source register describes the lowest bytes of the table.
//
//go:linkname Vqtbx3S8 Vqtbx3S8
//go:noescape
func Vqtbx3S8(r *arm.Int8X8, v0 *arm.Int8X8, v1 *arm.Int8X16X3, v2 *arm.Uint8X8)

// Table vector lookup extension. This instruction reads each value from the vector elements in the index source SIMD&FP register, uses each result as an index to perform a lookup in a table of bytes that is described by one to four source table SIMD&FP registers, places the lookup result in a vector, and writes the vector to the destination SIMD&FP register. If an index is out of range for the table, the existing value in the vector element of the destination register is left unchanged. If more than one source register is used to describe the table, the first source register describes the lowest bytes of the table.
//
//go:linkname Vqtbx4P8 Vqtbx4P8
//go:noescape
func Vqtbx4P8(r *arm.Poly8X8, v0 *arm.Poly8X8, v1 *arm.Poly8X16X4, v2 *arm.Uint8X8)

// Table vector lookup extension. This instruction reads each value from the vector elements in the index source SIMD&FP register, uses each result as an index to perform a lookup in a table of bytes that is described by one to four source table SIMD&FP registers, places the lookup result in a vector, and writes the vector to the destination SIMD&FP register. If an index is out of range for the table, the existing value in the vector element of the destination register is left unchanged. If more than one source register is used to describe the table, the first source register describes the lowest bytes of the table.
//
//go:linkname Vqtbx4QP8 Vqtbx4QP8
//go:noescape
func Vqtbx4QP8(r *arm.Poly8X16, v0 *arm.Poly8X16, v1 *arm.Poly8X16X4, v2 *arm.Uint8X16)

// Table vector lookup extension. This instruction reads each value from the vector elements in the index source SIMD&FP register, uses each result as an index to perform a lookup in a table of bytes that is described by one to four source table SIMD&FP registers, places the lookup result in a vector, and writes the vector to the destination SIMD&FP register. If an index is out of range for the table, the existing value in the vector element of the destination register is left unchanged. If more than one source register is used to describe the table, the first source register describes the lowest bytes of the table.
//
//go:linkname Vqtbx4QU8 Vqtbx4QU8
//go:noescape
func Vqtbx4QU8(r *arm.Uint8X16, v0 *arm.Uint8X16, v1 *arm.Uint8X16X4, v2 *arm.Uint8X16)

// Table vector lookup extension. This instruction reads each value from the vector elements in the index source SIMD&FP register, uses each result as an index to perform a lookup in a table of bytes that is described by one to four source table SIMD&FP registers, places the lookup result in a vector, and writes the vector to the destination SIMD&FP register. If an index is out of range for the table, the existing value in the vector element of the destination register is left unchanged. If more than one source register is used to describe the table, the first source register describes the lowest bytes of the table.
//
//go:linkname Vqtbx4QS8 Vqtbx4QS8
//go:noescape
func Vqtbx4QS8(r *arm.Int8X16, v0 *arm.Int8X16, v1 *arm.Int8X16X4, v2 *arm.Uint8X16)

// Table vector lookup extension. This instruction reads each value from the vector elements in the index source SIMD&FP register, uses each result as an index to perform a lookup in a table of bytes that is described by one to four source table SIMD&FP registers, places the lookup result in a vector, and writes the vector to the destination SIMD&FP register. If an index is out of range for the table, the existing value in the vector element of the destination register is left unchanged. If more than one source register is used to describe the table, the first source register describes the lowest bytes of the table.
//
//go:linkname Vqtbx4U8 Vqtbx4U8
//go:noescape
func Vqtbx4U8(r *arm.Uint8X8, v0 *arm.Uint8X8, v1 *arm.Uint8X16X4, v2 *arm.Uint8X8)

// Table vector lookup extension. This instruction reads each value from the vector elements in the index source SIMD&FP register, uses each result as an index to perform a lookup in a table of bytes that is described by one to four source table SIMD&FP registers, places the lookup result in a vector, and writes the vector to the destination SIMD&FP register. If an index is out of range for the table, the existing value in the vector element of the destination register is left unchanged. If more than one source register is used to describe the table, the first source register describes the lowest bytes of the table.
//
//go:linkname Vqtbx4S8 Vqtbx4S8
//go:noescape
func Vqtbx4S8(r *arm.Int8X8, v0 *arm.Int8X8, v1 *arm.Int8X16X4, v2 *arm.Uint8X8)

// Rounding Add returning High Narrow. This instruction adds each vector element in the first source SIMD&FP register to the corresponding vector element in the second source SIMD&FP register, places the most significant half of the result into a vector, and writes the vector to the lower or upper half of the destination SIMD&FP register.
//
//go:linkname VraddhnHighU32 VraddhnHighU32
//go:noescape
func VraddhnHighU32(r *arm.Uint16X8, v0 *arm.Uint16X4, v1 *arm.Uint32X4, v2 *arm.Uint32X4)

// Rounding Add returning High Narrow. This instruction adds each vector element in the first source SIMD&FP register to the corresponding vector element in the second source SIMD&FP register, places the most significant half of the result into a vector, and writes the vector to the lower or upper half of the destination SIMD&FP register.
//
//go:linkname VraddhnHighU64 VraddhnHighU64
//go:noescape
func VraddhnHighU64(r *arm.Uint32X4, v0 *arm.Uint32X2, v1 *arm.Uint64X2, v2 *arm.Uint64X2)

// Rounding Add returning High Narrow. This instruction adds each vector element in the first source SIMD&FP register to the corresponding vector element in the second source SIMD&FP register, places the most significant half of the result into a vector, and writes the vector to the lower or upper half of the destination SIMD&FP register.
//
//go:linkname VraddhnHighU16 VraddhnHighU16
//go:noescape
func VraddhnHighU16(r *arm.Uint8X16, v0 *arm.Uint8X8, v1 *arm.Uint16X8, v2 *arm.Uint16X8)

// Rounding Add returning High Narrow. This instruction adds each vector element in the first source SIMD&FP register to the corresponding vector element in the second source SIMD&FP register, places the most significant half of the result into a vector, and writes the vector to the lower or upper half of the destination SIMD&FP register.
//
//go:linkname VraddhnHighS32 VraddhnHighS32
//go:noescape
func VraddhnHighS32(r *arm.Int16X8, v0 *arm.Int16X4, v1 *arm.Int32X4, v2 *arm.Int32X4)

// Rounding Add returning High Narrow. This instruction adds each vector element in the first source SIMD&FP register to the corresponding vector element in the second source SIMD&FP register, places the most significant half of the result into a vector, and writes the vector to the lower or upper half of the destination SIMD&FP register.
//
//go:linkname VraddhnHighS64 VraddhnHighS64
//go:noescape
func VraddhnHighS64(r *arm.Int32X4, v0 *arm.Int32X2, v1 *arm.Int64X2, v2 *arm.Int64X2)

// Rounding Add returning High Narrow. This instruction adds each vector element in the first source SIMD&FP register to the corresponding vector element in the second source SIMD&FP register, places the most significant half of the result into a vector, and writes the vector to the lower or upper half of the destination SIMD&FP register.
//
//go:linkname VraddhnHighS16 VraddhnHighS16
//go:noescape
func VraddhnHighS16(r *arm.Int8X16, v0 *arm.Int8X8, v1 *arm.Int16X8, v2 *arm.Int16X8)

// Reverse Bit order (vector). This instruction reads each vector element from the source SIMD&FP register, reverses the bits of the element, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VrbitP8 VrbitP8
//go:noescape
func VrbitP8(r *arm.Poly8X8, v0 *arm.Poly8X8)

// Reverse Bit order (vector). This instruction reads each vector element from the source SIMD&FP register, reverses the bits of the element, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VrbitqP8 VrbitqP8
//go:noescape
func VrbitqP8(r *arm.Poly8X16, v0 *arm.Poly8X16)

// Reverse Bit order (vector). This instruction reads each vector element from the source SIMD&FP register, reverses the bits of the element, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VrbitqU8 VrbitqU8
//go:noescape
func VrbitqU8(r *arm.Uint8X16, v0 *arm.Uint8X16)

// Reverse Bit order (vector). This instruction reads each vector element from the source SIMD&FP register, reverses the bits of the element, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VrbitqS8 VrbitqS8
//go:noescape
func VrbitqS8(r *arm.Int8X16, v0 *arm.Int8X16)

// Reverse Bit order (vector). This instruction reads each vector element from the source SIMD&FP register, reverses the bits of the element, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VrbitU8 VrbitU8
//go:noescape
func VrbitU8(r *arm.Uint8X8, v0 *arm.Uint8X8)

// Reverse Bit order (vector). This instruction reads each vector element from the source SIMD&FP register, reverses the bits of the element, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VrbitS8 VrbitS8
//go:noescape
func VrbitS8(r *arm.Int8X8, v0 *arm.Int8X8)

// Floating-point Reciprocal Estimate. This instruction finds an approximate reciprocal estimate for each vector element in the source SIMD&FP register, places the result in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VrecpeqF64 VrecpeqF64
//go:noescape
func VrecpeqF64(r *arm.Float64X2, v0 *arm.Float64X2)

// Floating-point Reciprocal Estimate. This instruction finds an approximate reciprocal estimate for each vector element in the source SIMD&FP register, places the result in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VrecpeF64 VrecpeF64
//go:noescape
func VrecpeF64(r *arm.Float64X1, v0 *arm.Float64X1)

// Floating-point Reciprocal Estimate. This instruction finds an approximate reciprocal estimate for each vector element in the source SIMD&FP register, places the result in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VrecpedF64 VrecpedF64
//go:noescape
func VrecpedF64(r *arm.Float64, v0 *arm.Float64)

// Floating-point Reciprocal Estimate. This instruction finds an approximate reciprocal estimate for each vector element in the source SIMD&FP register, places the result in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VrecpesF32 VrecpesF32
//go:noescape
func VrecpesF32(r *arm.Float32, v0 *arm.Float32)

// Floating-point Reciprocal Step. This instruction multiplies the corresponding floating-point values in the vectors of the two source SIMD&FP registers, subtracts each of the products from 2.0, places the resulting floating-point values in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VrecpsqF64 VrecpsqF64
//go:noescape
func VrecpsqF64(r *arm.Float64X2, v0 *arm.Float64X2, v1 *arm.Float64X2)

// Floating-point Reciprocal Step. This instruction multiplies the corresponding floating-point values in the vectors of the two source SIMD&FP registers, subtracts each of the products from 2.0, places the resulting floating-point values in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VrecpsF64 VrecpsF64
//go:noescape
func VrecpsF64(r *arm.Float64X1, v0 *arm.Float64X1, v1 *arm.Float64X1)

// Floating-point Reciprocal Step. This instruction multiplies the corresponding floating-point values in the vectors of the two source SIMD&FP registers, subtracts each of the products from 2.0, places the resulting floating-point values in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VrecpsdF64 VrecpsdF64
//go:noescape
func VrecpsdF64(r *arm.Float64, v0 *arm.Float64, v1 *arm.Float64)

// Floating-point Reciprocal Step. This instruction multiplies the corresponding floating-point values in the vectors of the two source SIMD&FP registers, subtracts each of the products from 2.0, places the resulting floating-point values in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VrecpssF32 VrecpssF32
//go:noescape
func VrecpssF32(r *arm.Float32, v0 *arm.Float32, v1 *arm.Float32)

// Floating-point Reciprocal exponent (scalar). This instruction finds an approximate reciprocal exponent for each vector element in the source SIMD&FP register, places the result in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VrecpxdF64 VrecpxdF64
//go:noescape
func VrecpxdF64(r *arm.Float64, v0 *arm.Float64)

// Floating-point Reciprocal exponent (scalar). This instruction finds an approximate reciprocal exponent for each vector element in the source SIMD&FP register, places the result in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VrecpxsF32 VrecpxsF32
//go:noescape
func VrecpxsF32(r *arm.Float32, v0 *arm.Float32)

// Unsigned Rounding Shift Left (register). This instruction takes each element in the vector of the first source SIMD&FP register, shifts the vector element by a value from the least significant byte of the corresponding element of the second source SIMD&FP register, places the results in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VrshldU64 VrshldU64
//go:noescape
func VrshldU64(r *arm.Uint64, v0 *arm.Uint64, v1 *arm.Int64)

// Signed Rounding Shift Left (register). This instruction takes each signed integer value in the vector of the first source SIMD&FP register, shifts it by a value from the least significant byte of the corresponding element of the second source SIMD&FP register, places the results in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VrshldS64 VrshldS64
//go:noescape
func VrshldS64(r *arm.Int64, v0 *arm.Int64, v1 *arm.Int64)

// Floating-point Reciprocal Square Root Estimate. This instruction calculates an approximate square root for each vector element in the source SIMD&FP register, places the result in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VrsqrteqF64 VrsqrteqF64
//go:noescape
func VrsqrteqF64(r *arm.Float64X2, v0 *arm.Float64X2)

// Floating-point Reciprocal Square Root Estimate. This instruction calculates an approximate square root for each vector element in the source SIMD&FP register, places the result in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VrsqrteF64 VrsqrteF64
//go:noescape
func VrsqrteF64(r *arm.Float64X1, v0 *arm.Float64X1)

// Floating-point Reciprocal Square Root Estimate. This instruction calculates an approximate square root for each vector element in the source SIMD&FP register, places the result in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VrsqrtedF64 VrsqrtedF64
//go:noescape
func VrsqrtedF64(r *arm.Float64, v0 *arm.Float64)

// Floating-point Reciprocal Square Root Estimate. This instruction calculates an approximate square root for each vector element in the source SIMD&FP register, places the result in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VrsqrtesF32 VrsqrtesF32
//go:noescape
func VrsqrtesF32(r *arm.Float32, v0 *arm.Float32)

// Floating-point Reciprocal Square Root Step. This instruction multiplies corresponding floating-point values in the vectors of the two source SIMD&FP registers, subtracts each of the products from 3.0, divides these results by 2.0, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VrsqrtsqF64 VrsqrtsqF64
//go:noescape
func VrsqrtsqF64(r *arm.Float64X2, v0 *arm.Float64X2, v1 *arm.Float64X2)

// Floating-point Reciprocal Square Root Step. This instruction multiplies corresponding floating-point values in the vectors of the two source SIMD&FP registers, subtracts each of the products from 3.0, divides these results by 2.0, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VrsqrtsF64 VrsqrtsF64
//go:noescape
func VrsqrtsF64(r *arm.Float64X1, v0 *arm.Float64X1, v1 *arm.Float64X1)

// Floating-point Reciprocal Square Root Step. This instruction multiplies corresponding floating-point values in the vectors of the two source SIMD&FP registers, subtracts each of the products from 3.0, divides these results by 2.0, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VrsqrtsdF64 VrsqrtsdF64
//go:noescape
func VrsqrtsdF64(r *arm.Float64, v0 *arm.Float64, v1 *arm.Float64)

// Floating-point Reciprocal Square Root Step. This instruction multiplies corresponding floating-point values in the vectors of the two source SIMD&FP registers, subtracts each of the products from 3.0, divides these results by 2.0, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VrsqrtssF32 VrsqrtssF32
//go:noescape
func VrsqrtssF32(r *arm.Float32, v0 *arm.Float32, v1 *arm.Float32)

// Rounding Subtract returning High Narrow. This instruction subtracts each vector element of the second source SIMD&FP register from the corresponding vector element of the first source SIMD&FP register, places the most significant half of the result into a vector, and writes the vector to the lower or upper half of the destination SIMD&FP register.
//
//go:linkname VrsubhnHighU32 VrsubhnHighU32
//go:noescape
func VrsubhnHighU32(r *arm.Uint16X8, v0 *arm.Uint16X4, v1 *arm.Uint32X4, v2 *arm.Uint32X4)

// Rounding Subtract returning High Narrow. This instruction subtracts each vector element of the second source SIMD&FP register from the corresponding vector element of the first source SIMD&FP register, places the most significant half of the result into a vector, and writes the vector to the lower or upper half of the destination SIMD&FP register.
//
//go:linkname VrsubhnHighU64 VrsubhnHighU64
//go:noescape
func VrsubhnHighU64(r *arm.Uint32X4, v0 *arm.Uint32X2, v1 *arm.Uint64X2, v2 *arm.Uint64X2)

// Rounding Subtract returning High Narrow. This instruction subtracts each vector element of the second source SIMD&FP register from the corresponding vector element of the first source SIMD&FP register, places the most significant half of the result into a vector, and writes the vector to the lower or upper half of the destination SIMD&FP register.
//
//go:linkname VrsubhnHighU16 VrsubhnHighU16
//go:noescape
func VrsubhnHighU16(r *arm.Uint8X16, v0 *arm.Uint8X8, v1 *arm.Uint16X8, v2 *arm.Uint16X8)

// Rounding Subtract returning High Narrow. This instruction subtracts each vector element of the second source SIMD&FP register from the corresponding vector element of the first source SIMD&FP register, places the most significant half of the result into a vector, and writes the vector to the lower or upper half of the destination SIMD&FP register.
//
//go:linkname VrsubhnHighS32 VrsubhnHighS32
//go:noescape
func VrsubhnHighS32(r *arm.Int16X8, v0 *arm.Int16X4, v1 *arm.Int32X4, v2 *arm.Int32X4)

// Rounding Subtract returning High Narrow. This instruction subtracts each vector element of the second source SIMD&FP register from the corresponding vector element of the first source SIMD&FP register, places the most significant half of the result into a vector, and writes the vector to the lower or upper half of the destination SIMD&FP register.
//
//go:linkname VrsubhnHighS64 VrsubhnHighS64
//go:noescape
func VrsubhnHighS64(r *arm.Int32X4, v0 *arm.Int32X2, v1 *arm.Int64X2, v2 *arm.Int64X2)

// Rounding Subtract returning High Narrow. This instruction subtracts each vector element of the second source SIMD&FP register from the corresponding vector element of the first source SIMD&FP register, places the most significant half of the result into a vector, and writes the vector to the lower or upper half of the destination SIMD&FP register.
//
//go:linkname VrsubhnHighS16 VrsubhnHighS16
//go:noescape
func VrsubhnHighS16(r *arm.Int8X16, v0 *arm.Int8X8, v1 *arm.Int16X8, v2 *arm.Int16X8)

// Unsigned Shift Left (register). This instruction takes each element in the vector of the first source SIMD&FP register, shifts each element by a value from the least significant byte of the corresponding element of the second source SIMD&FP register, places the results in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VshldU64 VshldU64
//go:noescape
func VshldU64(r *arm.Uint64, v0 *arm.Uint64, v1 *arm.Int64)

// Signed Shift Left (register). This instruction takes each signed integer value in the vector of the first source SIMD&FP register, shifts each value by a value from the least significant byte of the corresponding element of the second source SIMD&FP register, places the results in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VshldS64 VshldS64
//go:noescape
func VshldS64(r *arm.Int64, v0 *arm.Int64, v1 *arm.Int64)

// Unsigned saturating Accumulate of Signed value. This instruction adds the signed integer values of the vector elements in the source SIMD&FP register to corresponding unsigned integer values of the vector elements in the destination SIMD&FP register, and accumulates the resulting unsigned integer values with the vector elements of the destination SIMD&FP register.
//
//go:linkname VsqaddbU8 VsqaddbU8
//go:noescape
func VsqaddbU8(r *arm.Uint8, v0 *arm.Uint8, v1 *arm.Int8)

// Unsigned saturating Accumulate of Signed value. This instruction adds the signed integer values of the vector elements in the source SIMD&FP register to corresponding unsigned integer values of the vector elements in the destination SIMD&FP register, and accumulates the resulting unsigned integer values with the vector elements of the destination SIMD&FP register.
//
//go:linkname VsqaddsU32 VsqaddsU32
//go:noescape
func VsqaddsU32(r *arm.Uint32, v0 *arm.Uint32, v1 *arm.Int32)

// Unsigned saturating Accumulate of Signed value. This instruction adds the signed integer values of the vector elements in the source SIMD&FP register to corresponding unsigned integer values of the vector elements in the destination SIMD&FP register, and accumulates the resulting unsigned integer values with the vector elements of the destination SIMD&FP register.
//
//go:linkname VsqadddU64 VsqadddU64
//go:noescape
func VsqadddU64(r *arm.Uint64, v0 *arm.Uint64, v1 *arm.Int64)

// Unsigned saturating Accumulate of Signed value. This instruction adds the signed integer values of the vector elements in the source SIMD&FP register to corresponding unsigned integer values of the vector elements in the destination SIMD&FP register, and accumulates the resulting unsigned integer values with the vector elements of the destination SIMD&FP register.
//
//go:linkname VsqaddhU16 VsqaddhU16
//go:noescape
func VsqaddhU16(r *arm.Uint16, v0 *arm.Uint16, v1 *arm.Int16)

// Unsigned saturating Accumulate of Signed value. This instruction adds the signed integer values of the vector elements in the source SIMD&FP register to corresponding unsigned integer values of the vector elements in the destination SIMD&FP register, and accumulates the resulting unsigned integer values with the vector elements of the destination SIMD&FP register.
//
//go:linkname VsqaddqU8 VsqaddqU8
//go:noescape
func VsqaddqU8(r *arm.Uint8X16, v0 *arm.Uint8X16, v1 *arm.Int8X16)

// Unsigned saturating Accumulate of Signed value. This instruction adds the signed integer values of the vector elements in the source SIMD&FP register to corresponding unsigned integer values of the vector elements in the destination SIMD&FP register, and accumulates the resulting unsigned integer values with the vector elements of the destination SIMD&FP register.
//
//go:linkname VsqaddqU32 VsqaddqU32
//go:noescape
func VsqaddqU32(r *arm.Uint32X4, v0 *arm.Uint32X4, v1 *arm.Int32X4)

// Unsigned saturating Accumulate of Signed value. This instruction adds the signed integer values of the vector elements in the source SIMD&FP register to corresponding unsigned integer values of the vector elements in the destination SIMD&FP register, and accumulates the resulting unsigned integer values with the vector elements of the destination SIMD&FP register.
//
//go:linkname VsqaddqU64 VsqaddqU64
//go:noescape
func VsqaddqU64(r *arm.Uint64X2, v0 *arm.Uint64X2, v1 *arm.Int64X2)

// Unsigned saturating Accumulate of Signed value. This instruction adds the signed integer values of the vector elements in the source SIMD&FP register to corresponding unsigned integer values of the vector elements in the destination SIMD&FP register, and accumulates the resulting unsigned integer values with the vector elements of the destination SIMD&FP register.
//
//go:linkname VsqaddqU16 VsqaddqU16
//go:noescape
func VsqaddqU16(r *arm.Uint16X8, v0 *arm.Uint16X8, v1 *arm.Int16X8)

// Unsigned saturating Accumulate of Signed value. This instruction adds the signed integer values of the vector elements in the source SIMD&FP register to corresponding unsigned integer values of the vector elements in the destination SIMD&FP register, and accumulates the resulting unsigned integer values with the vector elements of the destination SIMD&FP register.
//
//go:linkname VsqaddU8 VsqaddU8
//go:noescape
func VsqaddU8(r *arm.Uint8X8, v0 *arm.Uint8X8, v1 *arm.Int8X8)

// Unsigned saturating Accumulate of Signed value. This instruction adds the signed integer values of the vector elements in the source SIMD&FP register to corresponding unsigned integer values of the vector elements in the destination SIMD&FP register, and accumulates the resulting unsigned integer values with the vector elements of the destination SIMD&FP register.
//
//go:linkname VsqaddU32 VsqaddU32
//go:noescape
func VsqaddU32(r *arm.Uint32X2, v0 *arm.Uint32X2, v1 *arm.Int32X2)

// Unsigned saturating Accumulate of Signed value. This instruction adds the signed integer values of the vector elements in the source SIMD&FP register to corresponding unsigned integer values of the vector elements in the destination SIMD&FP register, and accumulates the resulting unsigned integer values with the vector elements of the destination SIMD&FP register.
//
//go:linkname VsqaddU64 VsqaddU64
//go:noescape
func VsqaddU64(r *arm.Uint64X1, v0 *arm.Uint64X1, v1 *arm.Int64X1)

// Unsigned saturating Accumulate of Signed value. This instruction adds the signed integer values of the vector elements in the source SIMD&FP register to corresponding unsigned integer values of the vector elements in the destination SIMD&FP register, and accumulates the resulting unsigned integer values with the vector elements of the destination SIMD&FP register.
//
//go:linkname VsqaddU16 VsqaddU16
//go:noescape
func VsqaddU16(r *arm.Uint16X4, v0 *arm.Uint16X4, v1 *arm.Int16X4)

// Floating-point Square Root (vector). This instruction calculates the square root for each vector element in the source SIMD&FP register, places the result in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VsqrtqF64 VsqrtqF64
//go:noescape
func VsqrtqF64(r *arm.Float64X2, v0 *arm.Float64X2)

// Floating-point Square Root (vector). This instruction calculates the square root for each vector element in the source SIMD&FP register, places the result in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VsqrtqF32 VsqrtqF32
//go:noescape
func VsqrtqF32(r *arm.Float32X4, v0 *arm.Float32X4)

// Floating-point Square Root (vector). This instruction calculates the square root for each vector element in the source SIMD&FP register, places the result in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VsqrtF64 VsqrtF64
//go:noescape
func VsqrtF64(r *arm.Float64X1, v0 *arm.Float64X1)

// Floating-point Square Root (vector). This instruction calculates the square root for each vector element in the source SIMD&FP register, places the result in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VsqrtF32 VsqrtF32
//go:noescape
func VsqrtF32(r *arm.Float32X2, v0 *arm.Float32X2)

// Subtract (vector). This instruction subtracts each vector element in the second source SIMD&FP register from the corresponding vector element in the first source SIMD&FP register, places the result into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VsubdU64 VsubdU64
//go:noescape
func VsubdU64(r *arm.Uint64, v0 *arm.Uint64, v1 *arm.Uint64)

// Subtract (vector). This instruction subtracts each vector element in the second source SIMD&FP register from the corresponding vector element in the first source SIMD&FP register, places the result into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VsubdS64 VsubdS64
//go:noescape
func VsubdS64(r *arm.Int64, v0 *arm.Int64, v1 *arm.Int64)

// Floating-point Subtract (vector). This instruction subtracts the elements in the vector in the second source SIMD&FP register, from the corresponding elements in the vector in the first source SIMD&FP register, places each result into elements of a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VsubqF64 VsubqF64
//go:noescape
func VsubqF64(r *arm.Float64X2, v0 *arm.Float64X2, v1 *arm.Float64X2)

// Floating-point Subtract (vector). This instruction subtracts the elements in the vector in the second source SIMD&FP register, from the corresponding elements in the vector in the first source SIMD&FP register, places each result into elements of a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VsubF64 VsubF64
//go:noescape
func VsubF64(r *arm.Float64X1, v0 *arm.Float64X1, v1 *arm.Float64X1)

// Subtract returning High Narrow. This instruction subtracts each vector element in the second source SIMD&FP register from the corresponding vector element in the first source SIMD&FP register, places the most significant half of the result into a vector, and writes the vector to the lower or upper half of the destination SIMD&FP register. All the values in this instruction are signed integer values.
//
//go:linkname VsubhnHighU32 VsubhnHighU32
//go:noescape
func VsubhnHighU32(r *arm.Uint16X8, v0 *arm.Uint16X4, v1 *arm.Uint32X4, v2 *arm.Uint32X4)

// Subtract returning High Narrow. This instruction subtracts each vector element in the second source SIMD&FP register from the corresponding vector element in the first source SIMD&FP register, places the most significant half of the result into a vector, and writes the vector to the lower or upper half of the destination SIMD&FP register. All the values in this instruction are signed integer values.
//
//go:linkname VsubhnHighU64 VsubhnHighU64
//go:noescape
func VsubhnHighU64(r *arm.Uint32X4, v0 *arm.Uint32X2, v1 *arm.Uint64X2, v2 *arm.Uint64X2)

// Subtract returning High Narrow. This instruction subtracts each vector element in the second source SIMD&FP register from the corresponding vector element in the first source SIMD&FP register, places the most significant half of the result into a vector, and writes the vector to the lower or upper half of the destination SIMD&FP register. All the values in this instruction are signed integer values.
//
//go:linkname VsubhnHighU16 VsubhnHighU16
//go:noescape
func VsubhnHighU16(r *arm.Uint8X16, v0 *arm.Uint8X8, v1 *arm.Uint16X8, v2 *arm.Uint16X8)

// Subtract returning High Narrow. This instruction subtracts each vector element in the second source SIMD&FP register from the corresponding vector element in the first source SIMD&FP register, places the most significant half of the result into a vector, and writes the vector to the lower or upper half of the destination SIMD&FP register. All the values in this instruction are signed integer values.
//
//go:linkname VsubhnHighS32 VsubhnHighS32
//go:noescape
func VsubhnHighS32(r *arm.Int16X8, v0 *arm.Int16X4, v1 *arm.Int32X4, v2 *arm.Int32X4)

// Subtract returning High Narrow. This instruction subtracts each vector element in the second source SIMD&FP register from the corresponding vector element in the first source SIMD&FP register, places the most significant half of the result into a vector, and writes the vector to the lower or upper half of the destination SIMD&FP register. All the values in this instruction are signed integer values.
//
//go:linkname VsubhnHighS64 VsubhnHighS64
//go:noescape
func VsubhnHighS64(r *arm.Int32X4, v0 *arm.Int32X2, v1 *arm.Int64X2, v2 *arm.Int64X2)

// Subtract returning High Narrow. This instruction subtracts each vector element in the second source SIMD&FP register from the corresponding vector element in the first source SIMD&FP register, places the most significant half of the result into a vector, and writes the vector to the lower or upper half of the destination SIMD&FP register. All the values in this instruction are signed integer values.
//
//go:linkname VsubhnHighS16 VsubhnHighS16
//go:noescape
func VsubhnHighS16(r *arm.Int8X16, v0 *arm.Int8X8, v1 *arm.Int16X8, v2 *arm.Int16X8)

// Unsigned Subtract Long. This instruction subtracts each vector element in the lower or upper half of the second source SIMD&FP register from the corresponding vector element of the first source SIMD&FP register, places the result into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are unsigned integer values. The destination vector elements are twice as long as the source vector elements.
//
//go:linkname VsublHighU8 VsublHighU8
//go:noescape
func VsublHighU8(r *arm.Uint16X8, v0 *arm.Uint8X16, v1 *arm.Uint8X16)

// Unsigned Subtract Long. This instruction subtracts each vector element in the lower or upper half of the second source SIMD&FP register from the corresponding vector element of the first source SIMD&FP register, places the result into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are unsigned integer values. The destination vector elements are twice as long as the source vector elements.
//
//go:linkname VsublHighU32 VsublHighU32
//go:noescape
func VsublHighU32(r *arm.Uint64X2, v0 *arm.Uint32X4, v1 *arm.Uint32X4)

// Unsigned Subtract Long. This instruction subtracts each vector element in the lower or upper half of the second source SIMD&FP register from the corresponding vector element of the first source SIMD&FP register, places the result into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are unsigned integer values. The destination vector elements are twice as long as the source vector elements.
//
//go:linkname VsublHighU16 VsublHighU16
//go:noescape
func VsublHighU16(r *arm.Uint32X4, v0 *arm.Uint16X8, v1 *arm.Uint16X8)

// Signed Subtract Long. This instruction subtracts each vector element in the lower or upper half of the second source SIMD&FP register from the corresponding vector element of the first source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are signed integer values. The destination vector elements are twice as long as the source vector elements.
//
//go:linkname VsublHighS8 VsublHighS8
//go:noescape
func VsublHighS8(r *arm.Int16X8, v0 *arm.Int8X16, v1 *arm.Int8X16)

// Signed Subtract Long. This instruction subtracts each vector element in the lower or upper half of the second source SIMD&FP register from the corresponding vector element of the first source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are signed integer values. The destination vector elements are twice as long as the source vector elements.
//
//go:linkname VsublHighS32 VsublHighS32
//go:noescape
func VsublHighS32(r *arm.Int64X2, v0 *arm.Int32X4, v1 *arm.Int32X4)

// Signed Subtract Long. This instruction subtracts each vector element in the lower or upper half of the second source SIMD&FP register from the corresponding vector element of the first source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are signed integer values. The destination vector elements are twice as long as the source vector elements.
//
//go:linkname VsublHighS16 VsublHighS16
//go:noescape
func VsublHighS16(r *arm.Int32X4, v0 *arm.Int16X8, v1 *arm.Int16X8)

// Unsigned Subtract Wide. This instruction subtracts each vector element of the second source SIMD&FP register from the corresponding vector element in the lower or upper half of the first source SIMD&FP register, places the result in a vector, and writes the vector to the SIMD&FP destination register. All the values in this instruction are signed integer values.
//
//go:linkname VsubwHighU8 VsubwHighU8
//go:noescape
func VsubwHighU8(r *arm.Uint16X8, v0 *arm.Uint16X8, v1 *arm.Uint8X16)

// Unsigned Subtract Wide. This instruction subtracts each vector element of the second source SIMD&FP register from the corresponding vector element in the lower or upper half of the first source SIMD&FP register, places the result in a vector, and writes the vector to the SIMD&FP destination register. All the values in this instruction are signed integer values.
//
//go:linkname VsubwHighU32 VsubwHighU32
//go:noescape
func VsubwHighU32(r *arm.Uint64X2, v0 *arm.Uint64X2, v1 *arm.Uint32X4)

// Unsigned Subtract Wide. This instruction subtracts each vector element of the second source SIMD&FP register from the corresponding vector element in the lower or upper half of the first source SIMD&FP register, places the result in a vector, and writes the vector to the SIMD&FP destination register. All the values in this instruction are signed integer values.
//
//go:linkname VsubwHighU16 VsubwHighU16
//go:noescape
func VsubwHighU16(r *arm.Uint32X4, v0 *arm.Uint32X4, v1 *arm.Uint16X8)

// Signed Subtract Wide. This instruction subtracts each vector element in the lower or upper half of the second source SIMD&FP register from the corresponding vector element in the first source SIMD&FP register, places the result in a vector, and writes the vector to the SIMD&FP destination register. All the values in this instruction are signed integer values.
//
//go:linkname VsubwHighS8 VsubwHighS8
//go:noescape
func VsubwHighS8(r *arm.Int16X8, v0 *arm.Int16X8, v1 *arm.Int8X16)

// Signed Subtract Wide. This instruction subtracts each vector element in the lower or upper half of the second source SIMD&FP register from the corresponding vector element in the first source SIMD&FP register, places the result in a vector, and writes the vector to the SIMD&FP destination register. All the values in this instruction are signed integer values.
//
//go:linkname VsubwHighS32 VsubwHighS32
//go:noescape
func VsubwHighS32(r *arm.Int64X2, v0 *arm.Int64X2, v1 *arm.Int32X4)

// Signed Subtract Wide. This instruction subtracts each vector element in the lower or upper half of the second source SIMD&FP register from the corresponding vector element in the first source SIMD&FP register, places the result in a vector, and writes the vector to the SIMD&FP destination register. All the values in this instruction are signed integer values.
//
//go:linkname VsubwHighS16 VsubwHighS16
//go:noescape
func VsubwHighS16(r *arm.Int32X4, v0 *arm.Int32X4, v1 *arm.Int16X8)

// Transpose vectors (primary). This instruction reads corresponding even-numbered vector elements from the two source SIMD&FP registers, starting at zero, places each result into consecutive elements of a vector, and writes the vector to the destination SIMD&FP register. Vector elements from the first source register are placed into even-numbered elements of the destination vector, starting at zero, while vector elements from the second source register are placed into odd-numbered elements of the destination vector.
//
//go:linkname Vtrn1P8 Vtrn1P8
//go:noescape
func Vtrn1P8(r *arm.Poly8X8, v0 *arm.Poly8X8, v1 *arm.Poly8X8)

// Transpose vectors (primary). This instruction reads corresponding even-numbered vector elements from the two source SIMD&FP registers, starting at zero, places each result into consecutive elements of a vector, and writes the vector to the destination SIMD&FP register. Vector elements from the first source register are placed into even-numbered elements of the destination vector, starting at zero, while vector elements from the second source register are placed into odd-numbered elements of the destination vector.
//
//go:linkname Vtrn1P16 Vtrn1P16
//go:noescape
func Vtrn1P16(r *arm.Poly16X4, v0 *arm.Poly16X4, v1 *arm.Poly16X4)

// Transpose vectors (primary). This instruction reads corresponding even-numbered vector elements from the two source SIMD&FP registers, starting at zero, places each result into consecutive elements of a vector, and writes the vector to the destination SIMD&FP register. Vector elements from the first source register are placed into even-numbered elements of the destination vector, starting at zero, while vector elements from the second source register are placed into odd-numbered elements of the destination vector.
//
//go:linkname Vtrn1QP8 Vtrn1QP8
//go:noescape
func Vtrn1QP8(r *arm.Poly8X16, v0 *arm.Poly8X16, v1 *arm.Poly8X16)

// Transpose vectors (primary). This instruction reads corresponding even-numbered vector elements from the two source SIMD&FP registers, starting at zero, places each result into consecutive elements of a vector, and writes the vector to the destination SIMD&FP register. Vector elements from the first source register are placed into even-numbered elements of the destination vector, starting at zero, while vector elements from the second source register are placed into odd-numbered elements of the destination vector.
//
//go:linkname Vtrn1QP64 Vtrn1QP64
//go:noescape
func Vtrn1QP64(r *arm.Poly64X2, v0 *arm.Poly64X2, v1 *arm.Poly64X2)

// Transpose vectors (primary). This instruction reads corresponding even-numbered vector elements from the two source SIMD&FP registers, starting at zero, places each result into consecutive elements of a vector, and writes the vector to the destination SIMD&FP register. Vector elements from the first source register are placed into even-numbered elements of the destination vector, starting at zero, while vector elements from the second source register are placed into odd-numbered elements of the destination vector.
//
//go:linkname Vtrn1QP16 Vtrn1QP16
//go:noescape
func Vtrn1QP16(r *arm.Poly16X8, v0 *arm.Poly16X8, v1 *arm.Poly16X8)

// Transpose vectors (primary). This instruction reads corresponding even-numbered vector elements from the two source SIMD&FP registers, starting at zero, places each result into consecutive elements of a vector, and writes the vector to the destination SIMD&FP register. Vector elements from the first source register are placed into even-numbered elements of the destination vector, starting at zero, while vector elements from the second source register are placed into odd-numbered elements of the destination vector.
//
//go:linkname Vtrn1QU8 Vtrn1QU8
//go:noescape
func Vtrn1QU8(r *arm.Uint8X16, v0 *arm.Uint8X16, v1 *arm.Uint8X16)

// Transpose vectors (primary). This instruction reads corresponding even-numbered vector elements from the two source SIMD&FP registers, starting at zero, places each result into consecutive elements of a vector, and writes the vector to the destination SIMD&FP register. Vector elements from the first source register are placed into even-numbered elements of the destination vector, starting at zero, while vector elements from the second source register are placed into odd-numbered elements of the destination vector.
//
//go:linkname Vtrn1QU32 Vtrn1QU32
//go:noescape
func Vtrn1QU32(r *arm.Uint32X4, v0 *arm.Uint32X4, v1 *arm.Uint32X4)

// Transpose vectors (primary). This instruction reads corresponding even-numbered vector elements from the two source SIMD&FP registers, starting at zero, places each result into consecutive elements of a vector, and writes the vector to the destination SIMD&FP register. Vector elements from the first source register are placed into even-numbered elements of the destination vector, starting at zero, while vector elements from the second source register are placed into odd-numbered elements of the destination vector.
//
//go:linkname Vtrn1QU64 Vtrn1QU64
//go:noescape
func Vtrn1QU64(r *arm.Uint64X2, v0 *arm.Uint64X2, v1 *arm.Uint64X2)

// Transpose vectors (primary). This instruction reads corresponding even-numbered vector elements from the two source SIMD&FP registers, starting at zero, places each result into consecutive elements of a vector, and writes the vector to the destination SIMD&FP register. Vector elements from the first source register are placed into even-numbered elements of the destination vector, starting at zero, while vector elements from the second source register are placed into odd-numbered elements of the destination vector.
//
//go:linkname Vtrn1QU16 Vtrn1QU16
//go:noescape
func Vtrn1QU16(r *arm.Uint16X8, v0 *arm.Uint16X8, v1 *arm.Uint16X8)

// Transpose vectors (primary). This instruction reads corresponding even-numbered vector elements from the two source SIMD&FP registers, starting at zero, places each result into consecutive elements of a vector, and writes the vector to the destination SIMD&FP register. Vector elements from the first source register are placed into even-numbered elements of the destination vector, starting at zero, while vector elements from the second source register are placed into odd-numbered elements of the destination vector.
//
//go:linkname Vtrn1QS8 Vtrn1QS8
//go:noescape
func Vtrn1QS8(r *arm.Int8X16, v0 *arm.Int8X16, v1 *arm.Int8X16)

// Transpose vectors (primary). This instruction reads corresponding even-numbered vector elements from the two source SIMD&FP registers, starting at zero, places each result into consecutive elements of a vector, and writes the vector to the destination SIMD&FP register. Vector elements from the first source register are placed into even-numbered elements of the destination vector, starting at zero, while vector elements from the second source register are placed into odd-numbered elements of the destination vector.
//
//go:linkname Vtrn1QF64 Vtrn1QF64
//go:noescape
func Vtrn1QF64(r *arm.Float64X2, v0 *arm.Float64X2, v1 *arm.Float64X2)

// Transpose vectors (primary). This instruction reads corresponding even-numbered vector elements from the two source SIMD&FP registers, starting at zero, places each result into consecutive elements of a vector, and writes the vector to the destination SIMD&FP register. Vector elements from the first source register are placed into even-numbered elements of the destination vector, starting at zero, while vector elements from the second source register are placed into odd-numbered elements of the destination vector.
//
//go:linkname Vtrn1QF32 Vtrn1QF32
//go:noescape
func Vtrn1QF32(r *arm.Float32X4, v0 *arm.Float32X4, v1 *arm.Float32X4)

// Transpose vectors (primary). This instruction reads corresponding even-numbered vector elements from the two source SIMD&FP registers, starting at zero, places each result into consecutive elements of a vector, and writes the vector to the destination SIMD&FP register. Vector elements from the first source register are placed into even-numbered elements of the destination vector, starting at zero, while vector elements from the second source register are placed into odd-numbered elements of the destination vector.
//
//go:linkname Vtrn1QS32 Vtrn1QS32
//go:noescape
func Vtrn1QS32(r *arm.Int32X4, v0 *arm.Int32X4, v1 *arm.Int32X4)

// Transpose vectors (primary). This instruction reads corresponding even-numbered vector elements from the two source SIMD&FP registers, starting at zero, places each result into consecutive elements of a vector, and writes the vector to the destination SIMD&FP register. Vector elements from the first source register are placed into even-numbered elements of the destination vector, starting at zero, while vector elements from the second source register are placed into odd-numbered elements of the destination vector.
//
//go:linkname Vtrn1QS64 Vtrn1QS64
//go:noescape
func Vtrn1QS64(r *arm.Int64X2, v0 *arm.Int64X2, v1 *arm.Int64X2)

// Transpose vectors (primary). This instruction reads corresponding even-numbered vector elements from the two source SIMD&FP registers, starting at zero, places each result into consecutive elements of a vector, and writes the vector to the destination SIMD&FP register. Vector elements from the first source register are placed into even-numbered elements of the destination vector, starting at zero, while vector elements from the second source register are placed into odd-numbered elements of the destination vector.
//
//go:linkname Vtrn1QS16 Vtrn1QS16
//go:noescape
func Vtrn1QS16(r *arm.Int16X8, v0 *arm.Int16X8, v1 *arm.Int16X8)

// Transpose vectors (primary). This instruction reads corresponding even-numbered vector elements from the two source SIMD&FP registers, starting at zero, places each result into consecutive elements of a vector, and writes the vector to the destination SIMD&FP register. Vector elements from the first source register are placed into even-numbered elements of the destination vector, starting at zero, while vector elements from the second source register are placed into odd-numbered elements of the destination vector.
//
//go:linkname Vtrn1U8 Vtrn1U8
//go:noescape
func Vtrn1U8(r *arm.Uint8X8, v0 *arm.Uint8X8, v1 *arm.Uint8X8)

// Transpose vectors (primary). This instruction reads corresponding even-numbered vector elements from the two source SIMD&FP registers, starting at zero, places each result into consecutive elements of a vector, and writes the vector to the destination SIMD&FP register. Vector elements from the first source register are placed into even-numbered elements of the destination vector, starting at zero, while vector elements from the second source register are placed into odd-numbered elements of the destination vector.
//
//go:linkname Vtrn1U32 Vtrn1U32
//go:noescape
func Vtrn1U32(r *arm.Uint32X2, v0 *arm.Uint32X2, v1 *arm.Uint32X2)

// Transpose vectors (primary). This instruction reads corresponding even-numbered vector elements from the two source SIMD&FP registers, starting at zero, places each result into consecutive elements of a vector, and writes the vector to the destination SIMD&FP register. Vector elements from the first source register are placed into even-numbered elements of the destination vector, starting at zero, while vector elements from the second source register are placed into odd-numbered elements of the destination vector.
//
//go:linkname Vtrn1U16 Vtrn1U16
//go:noescape
func Vtrn1U16(r *arm.Uint16X4, v0 *arm.Uint16X4, v1 *arm.Uint16X4)

// Transpose vectors (primary). This instruction reads corresponding even-numbered vector elements from the two source SIMD&FP registers, starting at zero, places each result into consecutive elements of a vector, and writes the vector to the destination SIMD&FP register. Vector elements from the first source register are placed into even-numbered elements of the destination vector, starting at zero, while vector elements from the second source register are placed into odd-numbered elements of the destination vector.
//
//go:linkname Vtrn1S8 Vtrn1S8
//go:noescape
func Vtrn1S8(r *arm.Int8X8, v0 *arm.Int8X8, v1 *arm.Int8X8)

// Transpose vectors (primary). This instruction reads corresponding even-numbered vector elements from the two source SIMD&FP registers, starting at zero, places each result into consecutive elements of a vector, and writes the vector to the destination SIMD&FP register. Vector elements from the first source register are placed into even-numbered elements of the destination vector, starting at zero, while vector elements from the second source register are placed into odd-numbered elements of the destination vector.
//
//go:linkname Vtrn1F32 Vtrn1F32
//go:noescape
func Vtrn1F32(r *arm.Float32X2, v0 *arm.Float32X2, v1 *arm.Float32X2)

// Transpose vectors (primary). This instruction reads corresponding even-numbered vector elements from the two source SIMD&FP registers, starting at zero, places each result into consecutive elements of a vector, and writes the vector to the destination SIMD&FP register. Vector elements from the first source register are placed into even-numbered elements of the destination vector, starting at zero, while vector elements from the second source register are placed into odd-numbered elements of the destination vector.
//
//go:linkname Vtrn1S32 Vtrn1S32
//go:noescape
func Vtrn1S32(r *arm.Int32X2, v0 *arm.Int32X2, v1 *arm.Int32X2)

// Transpose vectors (primary). This instruction reads corresponding even-numbered vector elements from the two source SIMD&FP registers, starting at zero, places each result into consecutive elements of a vector, and writes the vector to the destination SIMD&FP register. Vector elements from the first source register are placed into even-numbered elements of the destination vector, starting at zero, while vector elements from the second source register are placed into odd-numbered elements of the destination vector.
//
//go:linkname Vtrn1S16 Vtrn1S16
//go:noescape
func Vtrn1S16(r *arm.Int16X4, v0 *arm.Int16X4, v1 *arm.Int16X4)

// Transpose vectors (secondary). This instruction reads corresponding odd-numbered vector elements from the two source SIMD&FP registers, places each result into consecutive elements of a vector, and writes the vector to the destination SIMD&FP register. Vector elements from the first source register are placed into even-numbered elements of the destination vector, starting at zero, while vector elements from the second source register are placed into odd-numbered elements of the destination vector.
//
//go:linkname Vtrn2P8 Vtrn2P8
//go:noescape
func Vtrn2P8(r *arm.Poly8X8, v0 *arm.Poly8X8, v1 *arm.Poly8X8)

// Transpose vectors (secondary). This instruction reads corresponding odd-numbered vector elements from the two source SIMD&FP registers, places each result into consecutive elements of a vector, and writes the vector to the destination SIMD&FP register. Vector elements from the first source register are placed into even-numbered elements of the destination vector, starting at zero, while vector elements from the second source register are placed into odd-numbered elements of the destination vector.
//
//go:linkname Vtrn2P16 Vtrn2P16
//go:noescape
func Vtrn2P16(r *arm.Poly16X4, v0 *arm.Poly16X4, v1 *arm.Poly16X4)

// Transpose vectors (secondary). This instruction reads corresponding odd-numbered vector elements from the two source SIMD&FP registers, places each result into consecutive elements of a vector, and writes the vector to the destination SIMD&FP register. Vector elements from the first source register are placed into even-numbered elements of the destination vector, starting at zero, while vector elements from the second source register are placed into odd-numbered elements of the destination vector.
//
//go:linkname Vtrn2QP8 Vtrn2QP8
//go:noescape
func Vtrn2QP8(r *arm.Poly8X16, v0 *arm.Poly8X16, v1 *arm.Poly8X16)

// Transpose vectors (secondary). This instruction reads corresponding odd-numbered vector elements from the two source SIMD&FP registers, places each result into consecutive elements of a vector, and writes the vector to the destination SIMD&FP register. Vector elements from the first source register are placed into even-numbered elements of the destination vector, starting at zero, while vector elements from the second source register are placed into odd-numbered elements of the destination vector.
//
//go:linkname Vtrn2QP64 Vtrn2QP64
//go:noescape
func Vtrn2QP64(r *arm.Poly64X2, v0 *arm.Poly64X2, v1 *arm.Poly64X2)

// Transpose vectors (secondary). This instruction reads corresponding odd-numbered vector elements from the two source SIMD&FP registers, places each result into consecutive elements of a vector, and writes the vector to the destination SIMD&FP register. Vector elements from the first source register are placed into even-numbered elements of the destination vector, starting at zero, while vector elements from the second source register are placed into odd-numbered elements of the destination vector.
//
//go:linkname Vtrn2QP16 Vtrn2QP16
//go:noescape
func Vtrn2QP16(r *arm.Poly16X8, v0 *arm.Poly16X8, v1 *arm.Poly16X8)

// Transpose vectors (secondary). This instruction reads corresponding odd-numbered vector elements from the two source SIMD&FP registers, places each result into consecutive elements of a vector, and writes the vector to the destination SIMD&FP register. Vector elements from the first source register are placed into even-numbered elements of the destination vector, starting at zero, while vector elements from the second source register are placed into odd-numbered elements of the destination vector.
//
//go:linkname Vtrn2QU8 Vtrn2QU8
//go:noescape
func Vtrn2QU8(r *arm.Uint8X16, v0 *arm.Uint8X16, v1 *arm.Uint8X16)

// Transpose vectors (secondary). This instruction reads corresponding odd-numbered vector elements from the two source SIMD&FP registers, places each result into consecutive elements of a vector, and writes the vector to the destination SIMD&FP register. Vector elements from the first source register are placed into even-numbered elements of the destination vector, starting at zero, while vector elements from the second source register are placed into odd-numbered elements of the destination vector.
//
//go:linkname Vtrn2QU32 Vtrn2QU32
//go:noescape
func Vtrn2QU32(r *arm.Uint32X4, v0 *arm.Uint32X4, v1 *arm.Uint32X4)

// Transpose vectors (secondary). This instruction reads corresponding odd-numbered vector elements from the two source SIMD&FP registers, places each result into consecutive elements of a vector, and writes the vector to the destination SIMD&FP register. Vector elements from the first source register are placed into even-numbered elements of the destination vector, starting at zero, while vector elements from the second source register are placed into odd-numbered elements of the destination vector.
//
//go:linkname Vtrn2QU64 Vtrn2QU64
//go:noescape
func Vtrn2QU64(r *arm.Uint64X2, v0 *arm.Uint64X2, v1 *arm.Uint64X2)

// Transpose vectors (secondary). This instruction reads corresponding odd-numbered vector elements from the two source SIMD&FP registers, places each result into consecutive elements of a vector, and writes the vector to the destination SIMD&FP register. Vector elements from the first source register are placed into even-numbered elements of the destination vector, starting at zero, while vector elements from the second source register are placed into odd-numbered elements of the destination vector.
//
//go:linkname Vtrn2QU16 Vtrn2QU16
//go:noescape
func Vtrn2QU16(r *arm.Uint16X8, v0 *arm.Uint16X8, v1 *arm.Uint16X8)

// Transpose vectors (secondary). This instruction reads corresponding odd-numbered vector elements from the two source SIMD&FP registers, places each result into consecutive elements of a vector, and writes the vector to the destination SIMD&FP register. Vector elements from the first source register are placed into even-numbered elements of the destination vector, starting at zero, while vector elements from the second source register are placed into odd-numbered elements of the destination vector.
//
//go:linkname Vtrn2QS8 Vtrn2QS8
//go:noescape
func Vtrn2QS8(r *arm.Int8X16, v0 *arm.Int8X16, v1 *arm.Int8X16)

// Transpose vectors (secondary). This instruction reads corresponding odd-numbered vector elements from the two source SIMD&FP registers, places each result into consecutive elements of a vector, and writes the vector to the destination SIMD&FP register. Vector elements from the first source register are placed into even-numbered elements of the destination vector, starting at zero, while vector elements from the second source register are placed into odd-numbered elements of the destination vector.
//
//go:linkname Vtrn2QF64 Vtrn2QF64
//go:noescape
func Vtrn2QF64(r *arm.Float64X2, v0 *arm.Float64X2, v1 *arm.Float64X2)

// Transpose vectors (secondary). This instruction reads corresponding odd-numbered vector elements from the two source SIMD&FP registers, places each result into consecutive elements of a vector, and writes the vector to the destination SIMD&FP register. Vector elements from the first source register are placed into even-numbered elements of the destination vector, starting at zero, while vector elements from the second source register are placed into odd-numbered elements of the destination vector.
//
//go:linkname Vtrn2QF32 Vtrn2QF32
//go:noescape
func Vtrn2QF32(r *arm.Float32X4, v0 *arm.Float32X4, v1 *arm.Float32X4)

// Transpose vectors (secondary). This instruction reads corresponding odd-numbered vector elements from the two source SIMD&FP registers, places each result into consecutive elements of a vector, and writes the vector to the destination SIMD&FP register. Vector elements from the first source register are placed into even-numbered elements of the destination vector, starting at zero, while vector elements from the second source register are placed into odd-numbered elements of the destination vector.
//
//go:linkname Vtrn2QS32 Vtrn2QS32
//go:noescape
func Vtrn2QS32(r *arm.Int32X4, v0 *arm.Int32X4, v1 *arm.Int32X4)

// Transpose vectors (secondary). This instruction reads corresponding odd-numbered vector elements from the two source SIMD&FP registers, places each result into consecutive elements of a vector, and writes the vector to the destination SIMD&FP register. Vector elements from the first source register are placed into even-numbered elements of the destination vector, starting at zero, while vector elements from the second source register are placed into odd-numbered elements of the destination vector.
//
//go:linkname Vtrn2QS64 Vtrn2QS64
//go:noescape
func Vtrn2QS64(r *arm.Int64X2, v0 *arm.Int64X2, v1 *arm.Int64X2)

// Transpose vectors (secondary). This instruction reads corresponding odd-numbered vector elements from the two source SIMD&FP registers, places each result into consecutive elements of a vector, and writes the vector to the destination SIMD&FP register. Vector elements from the first source register are placed into even-numbered elements of the destination vector, starting at zero, while vector elements from the second source register are placed into odd-numbered elements of the destination vector.
//
//go:linkname Vtrn2QS16 Vtrn2QS16
//go:noescape
func Vtrn2QS16(r *arm.Int16X8, v0 *arm.Int16X8, v1 *arm.Int16X8)

// Transpose vectors (secondary). This instruction reads corresponding odd-numbered vector elements from the two source SIMD&FP registers, places each result into consecutive elements of a vector, and writes the vector to the destination SIMD&FP register. Vector elements from the first source register are placed into even-numbered elements of the destination vector, starting at zero, while vector elements from the second source register are placed into odd-numbered elements of the destination vector.
//
//go:linkname Vtrn2U8 Vtrn2U8
//go:noescape
func Vtrn2U8(r *arm.Uint8X8, v0 *arm.Uint8X8, v1 *arm.Uint8X8)

// Transpose vectors (secondary). This instruction reads corresponding odd-numbered vector elements from the two source SIMD&FP registers, places each result into consecutive elements of a vector, and writes the vector to the destination SIMD&FP register. Vector elements from the first source register are placed into even-numbered elements of the destination vector, starting at zero, while vector elements from the second source register are placed into odd-numbered elements of the destination vector.
//
//go:linkname Vtrn2U32 Vtrn2U32
//go:noescape
func Vtrn2U32(r *arm.Uint32X2, v0 *arm.Uint32X2, v1 *arm.Uint32X2)

// Transpose vectors (secondary). This instruction reads corresponding odd-numbered vector elements from the two source SIMD&FP registers, places each result into consecutive elements of a vector, and writes the vector to the destination SIMD&FP register. Vector elements from the first source register are placed into even-numbered elements of the destination vector, starting at zero, while vector elements from the second source register are placed into odd-numbered elements of the destination vector.
//
//go:linkname Vtrn2U16 Vtrn2U16
//go:noescape
func Vtrn2U16(r *arm.Uint16X4, v0 *arm.Uint16X4, v1 *arm.Uint16X4)

// Transpose vectors (secondary). This instruction reads corresponding odd-numbered vector elements from the two source SIMD&FP registers, places each result into consecutive elements of a vector, and writes the vector to the destination SIMD&FP register. Vector elements from the first source register are placed into even-numbered elements of the destination vector, starting at zero, while vector elements from the second source register are placed into odd-numbered elements of the destination vector.
//
//go:linkname Vtrn2S8 Vtrn2S8
//go:noescape
func Vtrn2S8(r *arm.Int8X8, v0 *arm.Int8X8, v1 *arm.Int8X8)

// Transpose vectors (secondary). This instruction reads corresponding odd-numbered vector elements from the two source SIMD&FP registers, places each result into consecutive elements of a vector, and writes the vector to the destination SIMD&FP register. Vector elements from the first source register are placed into even-numbered elements of the destination vector, starting at zero, while vector elements from the second source register are placed into odd-numbered elements of the destination vector.
//
//go:linkname Vtrn2F32 Vtrn2F32
//go:noescape
func Vtrn2F32(r *arm.Float32X2, v0 *arm.Float32X2, v1 *arm.Float32X2)

// Transpose vectors (secondary). This instruction reads corresponding odd-numbered vector elements from the two source SIMD&FP registers, places each result into consecutive elements of a vector, and writes the vector to the destination SIMD&FP register. Vector elements from the first source register are placed into even-numbered elements of the destination vector, starting at zero, while vector elements from the second source register are placed into odd-numbered elements of the destination vector.
//
//go:linkname Vtrn2S32 Vtrn2S32
//go:noescape
func Vtrn2S32(r *arm.Int32X2, v0 *arm.Int32X2, v1 *arm.Int32X2)

// Transpose vectors (secondary). This instruction reads corresponding odd-numbered vector elements from the two source SIMD&FP registers, places each result into consecutive elements of a vector, and writes the vector to the destination SIMD&FP register. Vector elements from the first source register are placed into even-numbered elements of the destination vector, starting at zero, while vector elements from the second source register are placed into odd-numbered elements of the destination vector.
//
//go:linkname Vtrn2S16 Vtrn2S16
//go:noescape
func Vtrn2S16(r *arm.Int16X4, v0 *arm.Int16X4, v1 *arm.Int16X4)

// Compare bitwise Test bits nonzero (vector). This instruction reads each vector element in the first source SIMD&FP register, performs an AND with the corresponding vector element in the second source SIMD&FP register, and if the result is not zero, sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VtstP64 VtstP64
//go:noescape
func VtstP64(r *arm.Uint64X1, v0 *arm.Poly64X1, v1 *arm.Poly64X1)

// Compare bitwise Test bits nonzero (vector). This instruction reads each vector element in the first source SIMD&FP register, performs an AND with the corresponding vector element in the second source SIMD&FP register, and if the result is not zero, sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VtstqP64 VtstqP64
//go:noescape
func VtstqP64(r *arm.Uint64X2, v0 *arm.Poly64X2, v1 *arm.Poly64X2)

// Compare bitwise Test bits nonzero (vector). This instruction reads each vector element in the first source SIMD&FP register, performs an AND with the corresponding vector element in the second source SIMD&FP register, and if the result is not zero, sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VtstqU64 VtstqU64
//go:noescape
func VtstqU64(r *arm.Uint64X2, v0 *arm.Uint64X2, v1 *arm.Uint64X2)

// Compare bitwise Test bits nonzero (vector). This instruction reads each vector element in the first source SIMD&FP register, performs an AND with the corresponding vector element in the second source SIMD&FP register, and if the result is not zero, sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VtstqS64 VtstqS64
//go:noescape
func VtstqS64(r *arm.Uint64X2, v0 *arm.Int64X2, v1 *arm.Int64X2)

// Compare bitwise Test bits nonzero (vector). This instruction reads each vector element in the first source SIMD&FP register, performs an AND with the corresponding vector element in the second source SIMD&FP register, and if the result is not zero, sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VtstU64 VtstU64
//go:noescape
func VtstU64(r *arm.Uint64X1, v0 *arm.Uint64X1, v1 *arm.Uint64X1)

// Compare bitwise Test bits nonzero (vector). This instruction reads each vector element in the first source SIMD&FP register, performs an AND with the corresponding vector element in the second source SIMD&FP register, and if the result is not zero, sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VtstS64 VtstS64
//go:noescape
func VtstS64(r *arm.Uint64X1, v0 *arm.Int64X1, v1 *arm.Int64X1)

// Compare bitwise Test bits nonzero (vector). This instruction reads each vector element in the first source SIMD&FP register, performs an AND with the corresponding vector element in the second source SIMD&FP register, and if the result is not zero, sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VtstdU64 VtstdU64
//go:noescape
func VtstdU64(r *arm.Uint64, v0 *arm.Uint64, v1 *arm.Uint64)

// Compare bitwise Test bits nonzero (vector). This instruction reads each vector element in the first source SIMD&FP register, performs an AND with the corresponding vector element in the second source SIMD&FP register, and if the result is not zero, sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VtstdS64 VtstdS64
//go:noescape
func VtstdS64(r *arm.Uint64, v0 *arm.Int64, v1 *arm.Int64)

// Signed saturating Accumulate of Unsigned value. This instruction adds the unsigned integer values of the vector elements in the source SIMD&FP register to corresponding signed integer values of the vector elements in the destination SIMD&FP register, and writes the resulting signed integer values to the destination SIMD&FP register.
//
//go:linkname VuqaddbS8 VuqaddbS8
//go:noescape
func VuqaddbS8(r *arm.Int8, v0 *arm.Int8, v1 *arm.Uint8)

// Signed saturating Accumulate of Unsigned value. This instruction adds the unsigned integer values of the vector elements in the source SIMD&FP register to corresponding signed integer values of the vector elements in the destination SIMD&FP register, and writes the resulting signed integer values to the destination SIMD&FP register.
//
//go:linkname VuqaddsS32 VuqaddsS32
//go:noescape
func VuqaddsS32(r *arm.Int32, v0 *arm.Int32, v1 *arm.Uint32)

// Signed saturating Accumulate of Unsigned value. This instruction adds the unsigned integer values of the vector elements in the source SIMD&FP register to corresponding signed integer values of the vector elements in the destination SIMD&FP register, and writes the resulting signed integer values to the destination SIMD&FP register.
//
//go:linkname VuqadddS64 VuqadddS64
//go:noescape
func VuqadddS64(r *arm.Int64, v0 *arm.Int64, v1 *arm.Uint64)

// Signed saturating Accumulate of Unsigned value. This instruction adds the unsigned integer values of the vector elements in the source SIMD&FP register to corresponding signed integer values of the vector elements in the destination SIMD&FP register, and writes the resulting signed integer values to the destination SIMD&FP register.
//
//go:linkname VuqaddhS16 VuqaddhS16
//go:noescape
func VuqaddhS16(r *arm.Int16, v0 *arm.Int16, v1 *arm.Uint16)

// Signed saturating Accumulate of Unsigned value. This instruction adds the unsigned integer values of the vector elements in the source SIMD&FP register to corresponding signed integer values of the vector elements in the destination SIMD&FP register, and writes the resulting signed integer values to the destination SIMD&FP register.
//
//go:linkname VuqaddqS8 VuqaddqS8
//go:noescape
func VuqaddqS8(r *arm.Int8X16, v0 *arm.Int8X16, v1 *arm.Uint8X16)

// Signed saturating Accumulate of Unsigned value. This instruction adds the unsigned integer values of the vector elements in the source SIMD&FP register to corresponding signed integer values of the vector elements in the destination SIMD&FP register, and writes the resulting signed integer values to the destination SIMD&FP register.
//
//go:linkname VuqaddqS32 VuqaddqS32
//go:noescape
func VuqaddqS32(r *arm.Int32X4, v0 *arm.Int32X4, v1 *arm.Uint32X4)

// Signed saturating Accumulate of Unsigned value. This instruction adds the unsigned integer values of the vector elements in the source SIMD&FP register to corresponding signed integer values of the vector elements in the destination SIMD&FP register, and writes the resulting signed integer values to the destination SIMD&FP register.
//
//go:linkname VuqaddqS64 VuqaddqS64
//go:noescape
func VuqaddqS64(r *arm.Int64X2, v0 *arm.Int64X2, v1 *arm.Uint64X2)

// Signed saturating Accumulate of Unsigned value. This instruction adds the unsigned integer values of the vector elements in the source SIMD&FP register to corresponding signed integer values of the vector elements in the destination SIMD&FP register, and writes the resulting signed integer values to the destination SIMD&FP register.
//
//go:linkname VuqaddqS16 VuqaddqS16
//go:noescape
func VuqaddqS16(r *arm.Int16X8, v0 *arm.Int16X8, v1 *arm.Uint16X8)

// Signed saturating Accumulate of Unsigned value. This instruction adds the unsigned integer values of the vector elements in the source SIMD&FP register to corresponding signed integer values of the vector elements in the destination SIMD&FP register, and writes the resulting signed integer values to the destination SIMD&FP register.
//
//go:linkname VuqaddS8 VuqaddS8
//go:noescape
func VuqaddS8(r *arm.Int8X8, v0 *arm.Int8X8, v1 *arm.Uint8X8)

// Signed saturating Accumulate of Unsigned value. This instruction adds the unsigned integer values of the vector elements in the source SIMD&FP register to corresponding signed integer values of the vector elements in the destination SIMD&FP register, and writes the resulting signed integer values to the destination SIMD&FP register.
//
//go:linkname VuqaddS32 VuqaddS32
//go:noescape
func VuqaddS32(r *arm.Int32X2, v0 *arm.Int32X2, v1 *arm.Uint32X2)

// Signed saturating Accumulate of Unsigned value. This instruction adds the unsigned integer values of the vector elements in the source SIMD&FP register to corresponding signed integer values of the vector elements in the destination SIMD&FP register, and writes the resulting signed integer values to the destination SIMD&FP register.
//
//go:linkname VuqaddS64 VuqaddS64
//go:noescape
func VuqaddS64(r *arm.Int64X1, v0 *arm.Int64X1, v1 *arm.Uint64X1)

// Signed saturating Accumulate of Unsigned value. This instruction adds the unsigned integer values of the vector elements in the source SIMD&FP register to corresponding signed integer values of the vector elements in the destination SIMD&FP register, and writes the resulting signed integer values to the destination SIMD&FP register.
//
//go:linkname VuqaddS16 VuqaddS16
//go:noescape
func VuqaddS16(r *arm.Int16X4, v0 *arm.Int16X4, v1 *arm.Uint16X4)

// Unzip vectors (primary). This instruction reads corresponding even-numbered vector elements from the two source SIMD&FP registers, starting at zero, places the result from the first source register into consecutive elements in the lower half of a vector, and the result from the second source register into consecutive elements in the upper half of a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname Vuzp1P8 Vuzp1P8
//go:noescape
func Vuzp1P8(r *arm.Poly8X8, v0 *arm.Poly8X8, v1 *arm.Poly8X8)

// Unzip vectors (primary). This instruction reads corresponding even-numbered vector elements from the two source SIMD&FP registers, starting at zero, places the result from the first source register into consecutive elements in the lower half of a vector, and the result from the second source register into consecutive elements in the upper half of a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname Vuzp1P16 Vuzp1P16
//go:noescape
func Vuzp1P16(r *arm.Poly16X4, v0 *arm.Poly16X4, v1 *arm.Poly16X4)

// Unzip vectors (primary). This instruction reads corresponding even-numbered vector elements from the two source SIMD&FP registers, starting at zero, places the result from the first source register into consecutive elements in the lower half of a vector, and the result from the second source register into consecutive elements in the upper half of a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname Vuzp1QP8 Vuzp1QP8
//go:noescape
func Vuzp1QP8(r *arm.Poly8X16, v0 *arm.Poly8X16, v1 *arm.Poly8X16)

// Unzip vectors (primary). This instruction reads corresponding even-numbered vector elements from the two source SIMD&FP registers, starting at zero, places the result from the first source register into consecutive elements in the lower half of a vector, and the result from the second source register into consecutive elements in the upper half of a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname Vuzp1QP64 Vuzp1QP64
//go:noescape
func Vuzp1QP64(r *arm.Poly64X2, v0 *arm.Poly64X2, v1 *arm.Poly64X2)

// Unzip vectors (primary). This instruction reads corresponding even-numbered vector elements from the two source SIMD&FP registers, starting at zero, places the result from the first source register into consecutive elements in the lower half of a vector, and the result from the second source register into consecutive elements in the upper half of a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname Vuzp1QP16 Vuzp1QP16
//go:noescape
func Vuzp1QP16(r *arm.Poly16X8, v0 *arm.Poly16X8, v1 *arm.Poly16X8)

// Unzip vectors (primary). This instruction reads corresponding even-numbered vector elements from the two source SIMD&FP registers, starting at zero, places the result from the first source register into consecutive elements in the lower half of a vector, and the result from the second source register into consecutive elements in the upper half of a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname Vuzp1QU8 Vuzp1QU8
//go:noescape
func Vuzp1QU8(r *arm.Uint8X16, v0 *arm.Uint8X16, v1 *arm.Uint8X16)

// Unzip vectors (primary). This instruction reads corresponding even-numbered vector elements from the two source SIMD&FP registers, starting at zero, places the result from the first source register into consecutive elements in the lower half of a vector, and the result from the second source register into consecutive elements in the upper half of a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname Vuzp1QU32 Vuzp1QU32
//go:noescape
func Vuzp1QU32(r *arm.Uint32X4, v0 *arm.Uint32X4, v1 *arm.Uint32X4)

// Unzip vectors (primary). This instruction reads corresponding even-numbered vector elements from the two source SIMD&FP registers, starting at zero, places the result from the first source register into consecutive elements in the lower half of a vector, and the result from the second source register into consecutive elements in the upper half of a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname Vuzp1QU64 Vuzp1QU64
//go:noescape
func Vuzp1QU64(r *arm.Uint64X2, v0 *arm.Uint64X2, v1 *arm.Uint64X2)

// Unzip vectors (primary). This instruction reads corresponding even-numbered vector elements from the two source SIMD&FP registers, starting at zero, places the result from the first source register into consecutive elements in the lower half of a vector, and the result from the second source register into consecutive elements in the upper half of a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname Vuzp1QU16 Vuzp1QU16
//go:noescape
func Vuzp1QU16(r *arm.Uint16X8, v0 *arm.Uint16X8, v1 *arm.Uint16X8)

// Unzip vectors (primary). This instruction reads corresponding even-numbered vector elements from the two source SIMD&FP registers, starting at zero, places the result from the first source register into consecutive elements in the lower half of a vector, and the result from the second source register into consecutive elements in the upper half of a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname Vuzp1QS8 Vuzp1QS8
//go:noescape
func Vuzp1QS8(r *arm.Int8X16, v0 *arm.Int8X16, v1 *arm.Int8X16)

// Unzip vectors (primary). This instruction reads corresponding even-numbered vector elements from the two source SIMD&FP registers, starting at zero, places the result from the first source register into consecutive elements in the lower half of a vector, and the result from the second source register into consecutive elements in the upper half of a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname Vuzp1QF64 Vuzp1QF64
//go:noescape
func Vuzp1QF64(r *arm.Float64X2, v0 *arm.Float64X2, v1 *arm.Float64X2)

// Unzip vectors (primary). This instruction reads corresponding even-numbered vector elements from the two source SIMD&FP registers, starting at zero, places the result from the first source register into consecutive elements in the lower half of a vector, and the result from the second source register into consecutive elements in the upper half of a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname Vuzp1QF32 Vuzp1QF32
//go:noescape
func Vuzp1QF32(r *arm.Float32X4, v0 *arm.Float32X4, v1 *arm.Float32X4)

// Unzip vectors (primary). This instruction reads corresponding even-numbered vector elements from the two source SIMD&FP registers, starting at zero, places the result from the first source register into consecutive elements in the lower half of a vector, and the result from the second source register into consecutive elements in the upper half of a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname Vuzp1QS32 Vuzp1QS32
//go:noescape
func Vuzp1QS32(r *arm.Int32X4, v0 *arm.Int32X4, v1 *arm.Int32X4)

// Unzip vectors (primary). This instruction reads corresponding even-numbered vector elements from the two source SIMD&FP registers, starting at zero, places the result from the first source register into consecutive elements in the lower half of a vector, and the result from the second source register into consecutive elements in the upper half of a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname Vuzp1QS64 Vuzp1QS64
//go:noescape
func Vuzp1QS64(r *arm.Int64X2, v0 *arm.Int64X2, v1 *arm.Int64X2)

// Unzip vectors (primary). This instruction reads corresponding even-numbered vector elements from the two source SIMD&FP registers, starting at zero, places the result from the first source register into consecutive elements in the lower half of a vector, and the result from the second source register into consecutive elements in the upper half of a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname Vuzp1QS16 Vuzp1QS16
//go:noescape
func Vuzp1QS16(r *arm.Int16X8, v0 *arm.Int16X8, v1 *arm.Int16X8)

// Unzip vectors (primary). This instruction reads corresponding even-numbered vector elements from the two source SIMD&FP registers, starting at zero, places the result from the first source register into consecutive elements in the lower half of a vector, and the result from the second source register into consecutive elements in the upper half of a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname Vuzp1U8 Vuzp1U8
//go:noescape
func Vuzp1U8(r *arm.Uint8X8, v0 *arm.Uint8X8, v1 *arm.Uint8X8)

// Unzip vectors (primary). This instruction reads corresponding even-numbered vector elements from the two source SIMD&FP registers, starting at zero, places the result from the first source register into consecutive elements in the lower half of a vector, and the result from the second source register into consecutive elements in the upper half of a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname Vuzp1U32 Vuzp1U32
//go:noescape
func Vuzp1U32(r *arm.Uint32X2, v0 *arm.Uint32X2, v1 *arm.Uint32X2)

// Unzip vectors (primary). This instruction reads corresponding even-numbered vector elements from the two source SIMD&FP registers, starting at zero, places the result from the first source register into consecutive elements in the lower half of a vector, and the result from the second source register into consecutive elements in the upper half of a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname Vuzp1U16 Vuzp1U16
//go:noescape
func Vuzp1U16(r *arm.Uint16X4, v0 *arm.Uint16X4, v1 *arm.Uint16X4)

// Unzip vectors (primary). This instruction reads corresponding even-numbered vector elements from the two source SIMD&FP registers, starting at zero, places the result from the first source register into consecutive elements in the lower half of a vector, and the result from the second source register into consecutive elements in the upper half of a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname Vuzp1S8 Vuzp1S8
//go:noescape
func Vuzp1S8(r *arm.Int8X8, v0 *arm.Int8X8, v1 *arm.Int8X8)

// Unzip vectors (primary). This instruction reads corresponding even-numbered vector elements from the two source SIMD&FP registers, starting at zero, places the result from the first source register into consecutive elements in the lower half of a vector, and the result from the second source register into consecutive elements in the upper half of a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname Vuzp1F32 Vuzp1F32
//go:noescape
func Vuzp1F32(r *arm.Float32X2, v0 *arm.Float32X2, v1 *arm.Float32X2)

// Unzip vectors (primary). This instruction reads corresponding even-numbered vector elements from the two source SIMD&FP registers, starting at zero, places the result from the first source register into consecutive elements in the lower half of a vector, and the result from the second source register into consecutive elements in the upper half of a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname Vuzp1S32 Vuzp1S32
//go:noescape
func Vuzp1S32(r *arm.Int32X2, v0 *arm.Int32X2, v1 *arm.Int32X2)

// Unzip vectors (primary). This instruction reads corresponding even-numbered vector elements from the two source SIMD&FP registers, starting at zero, places the result from the first source register into consecutive elements in the lower half of a vector, and the result from the second source register into consecutive elements in the upper half of a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname Vuzp1S16 Vuzp1S16
//go:noescape
func Vuzp1S16(r *arm.Int16X4, v0 *arm.Int16X4, v1 *arm.Int16X4)

// Unzip vectors (secondary). This instruction reads corresponding odd-numbered vector elements from the two source SIMD&FP registers, places the result from the first source register into consecutive elements in the lower half of a vector, and the result from the second source register into consecutive elements in the upper half of a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname Vuzp2P8 Vuzp2P8
//go:noescape
func Vuzp2P8(r *arm.Poly8X8, v0 *arm.Poly8X8, v1 *arm.Poly8X8)

// Unzip vectors (secondary). This instruction reads corresponding odd-numbered vector elements from the two source SIMD&FP registers, places the result from the first source register into consecutive elements in the lower half of a vector, and the result from the second source register into consecutive elements in the upper half of a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname Vuzp2P16 Vuzp2P16
//go:noescape
func Vuzp2P16(r *arm.Poly16X4, v0 *arm.Poly16X4, v1 *arm.Poly16X4)

// Unzip vectors (secondary). This instruction reads corresponding odd-numbered vector elements from the two source SIMD&FP registers, places the result from the first source register into consecutive elements in the lower half of a vector, and the result from the second source register into consecutive elements in the upper half of a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname Vuzp2QP8 Vuzp2QP8
//go:noescape
func Vuzp2QP8(r *arm.Poly8X16, v0 *arm.Poly8X16, v1 *arm.Poly8X16)

// Unzip vectors (secondary). This instruction reads corresponding odd-numbered vector elements from the two source SIMD&FP registers, places the result from the first source register into consecutive elements in the lower half of a vector, and the result from the second source register into consecutive elements in the upper half of a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname Vuzp2QP64 Vuzp2QP64
//go:noescape
func Vuzp2QP64(r *arm.Poly64X2, v0 *arm.Poly64X2, v1 *arm.Poly64X2)

// Unzip vectors (secondary). This instruction reads corresponding odd-numbered vector elements from the two source SIMD&FP registers, places the result from the first source register into consecutive elements in the lower half of a vector, and the result from the second source register into consecutive elements in the upper half of a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname Vuzp2QP16 Vuzp2QP16
//go:noescape
func Vuzp2QP16(r *arm.Poly16X8, v0 *arm.Poly16X8, v1 *arm.Poly16X8)

// Unzip vectors (secondary). This instruction reads corresponding odd-numbered vector elements from the two source SIMD&FP registers, places the result from the first source register into consecutive elements in the lower half of a vector, and the result from the second source register into consecutive elements in the upper half of a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname Vuzp2QU8 Vuzp2QU8
//go:noescape
func Vuzp2QU8(r *arm.Uint8X16, v0 *arm.Uint8X16, v1 *arm.Uint8X16)

// Unzip vectors (secondary). This instruction reads corresponding odd-numbered vector elements from the two source SIMD&FP registers, places the result from the first source register into consecutive elements in the lower half of a vector, and the result from the second source register into consecutive elements in the upper half of a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname Vuzp2QU32 Vuzp2QU32
//go:noescape
func Vuzp2QU32(r *arm.Uint32X4, v0 *arm.Uint32X4, v1 *arm.Uint32X4)

// Unzip vectors (secondary). This instruction reads corresponding odd-numbered vector elements from the two source SIMD&FP registers, places the result from the first source register into consecutive elements in the lower half of a vector, and the result from the second source register into consecutive elements in the upper half of a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname Vuzp2QU64 Vuzp2QU64
//go:noescape
func Vuzp2QU64(r *arm.Uint64X2, v0 *arm.Uint64X2, v1 *arm.Uint64X2)

// Unzip vectors (secondary). This instruction reads corresponding odd-numbered vector elements from the two source SIMD&FP registers, places the result from the first source register into consecutive elements in the lower half of a vector, and the result from the second source register into consecutive elements in the upper half of a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname Vuzp2QU16 Vuzp2QU16
//go:noescape
func Vuzp2QU16(r *arm.Uint16X8, v0 *arm.Uint16X8, v1 *arm.Uint16X8)

// Unzip vectors (secondary). This instruction reads corresponding odd-numbered vector elements from the two source SIMD&FP registers, places the result from the first source register into consecutive elements in the lower half of a vector, and the result from the second source register into consecutive elements in the upper half of a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname Vuzp2QS8 Vuzp2QS8
//go:noescape
func Vuzp2QS8(r *arm.Int8X16, v0 *arm.Int8X16, v1 *arm.Int8X16)

// Unzip vectors (secondary). This instruction reads corresponding odd-numbered vector elements from the two source SIMD&FP registers, places the result from the first source register into consecutive elements in the lower half of a vector, and the result from the second source register into consecutive elements in the upper half of a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname Vuzp2QF64 Vuzp2QF64
//go:noescape
func Vuzp2QF64(r *arm.Float64X2, v0 *arm.Float64X2, v1 *arm.Float64X2)

// Unzip vectors (secondary). This instruction reads corresponding odd-numbered vector elements from the two source SIMD&FP registers, places the result from the first source register into consecutive elements in the lower half of a vector, and the result from the second source register into consecutive elements in the upper half of a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname Vuzp2QF32 Vuzp2QF32
//go:noescape
func Vuzp2QF32(r *arm.Float32X4, v0 *arm.Float32X4, v1 *arm.Float32X4)

// Unzip vectors (secondary). This instruction reads corresponding odd-numbered vector elements from the two source SIMD&FP registers, places the result from the first source register into consecutive elements in the lower half of a vector, and the result from the second source register into consecutive elements in the upper half of a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname Vuzp2QS32 Vuzp2QS32
//go:noescape
func Vuzp2QS32(r *arm.Int32X4, v0 *arm.Int32X4, v1 *arm.Int32X4)

// Unzip vectors (secondary). This instruction reads corresponding odd-numbered vector elements from the two source SIMD&FP registers, places the result from the first source register into consecutive elements in the lower half of a vector, and the result from the second source register into consecutive elements in the upper half of a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname Vuzp2QS64 Vuzp2QS64
//go:noescape
func Vuzp2QS64(r *arm.Int64X2, v0 *arm.Int64X2, v1 *arm.Int64X2)

// Unzip vectors (secondary). This instruction reads corresponding odd-numbered vector elements from the two source SIMD&FP registers, places the result from the first source register into consecutive elements in the lower half of a vector, and the result from the second source register into consecutive elements in the upper half of a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname Vuzp2QS16 Vuzp2QS16
//go:noescape
func Vuzp2QS16(r *arm.Int16X8, v0 *arm.Int16X8, v1 *arm.Int16X8)

// Unzip vectors (secondary). This instruction reads corresponding odd-numbered vector elements from the two source SIMD&FP registers, places the result from the first source register into consecutive elements in the lower half of a vector, and the result from the second source register into consecutive elements in the upper half of a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname Vuzp2U8 Vuzp2U8
//go:noescape
func Vuzp2U8(r *arm.Uint8X8, v0 *arm.Uint8X8, v1 *arm.Uint8X8)

// Unzip vectors (secondary). This instruction reads corresponding odd-numbered vector elements from the two source SIMD&FP registers, places the result from the first source register into consecutive elements in the lower half of a vector, and the result from the second source register into consecutive elements in the upper half of a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname Vuzp2U32 Vuzp2U32
//go:noescape
func Vuzp2U32(r *arm.Uint32X2, v0 *arm.Uint32X2, v1 *arm.Uint32X2)

// Unzip vectors (secondary). This instruction reads corresponding odd-numbered vector elements from the two source SIMD&FP registers, places the result from the first source register into consecutive elements in the lower half of a vector, and the result from the second source register into consecutive elements in the upper half of a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname Vuzp2U16 Vuzp2U16
//go:noescape
func Vuzp2U16(r *arm.Uint16X4, v0 *arm.Uint16X4, v1 *arm.Uint16X4)

// Unzip vectors (secondary). This instruction reads corresponding odd-numbered vector elements from the two source SIMD&FP registers, places the result from the first source register into consecutive elements in the lower half of a vector, and the result from the second source register into consecutive elements in the upper half of a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname Vuzp2S8 Vuzp2S8
//go:noescape
func Vuzp2S8(r *arm.Int8X8, v0 *arm.Int8X8, v1 *arm.Int8X8)

// Unzip vectors (secondary). This instruction reads corresponding odd-numbered vector elements from the two source SIMD&FP registers, places the result from the first source register into consecutive elements in the lower half of a vector, and the result from the second source register into consecutive elements in the upper half of a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname Vuzp2F32 Vuzp2F32
//go:noescape
func Vuzp2F32(r *arm.Float32X2, v0 *arm.Float32X2, v1 *arm.Float32X2)

// Unzip vectors (secondary). This instruction reads corresponding odd-numbered vector elements from the two source SIMD&FP registers, places the result from the first source register into consecutive elements in the lower half of a vector, and the result from the second source register into consecutive elements in the upper half of a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname Vuzp2S32 Vuzp2S32
//go:noescape
func Vuzp2S32(r *arm.Int32X2, v0 *arm.Int32X2, v1 *arm.Int32X2)

// Unzip vectors (secondary). This instruction reads corresponding odd-numbered vector elements from the two source SIMD&FP registers, places the result from the first source register into consecutive elements in the lower half of a vector, and the result from the second source register into consecutive elements in the upper half of a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname Vuzp2S16 Vuzp2S16
//go:noescape
func Vuzp2S16(r *arm.Int16X4, v0 *arm.Int16X4, v1 *arm.Int16X4)

// Zip vectors (primary). This instruction reads adjacent vector elements from the lower half of two source SIMD&FP registers as pairs, interleaves the pairs and places them into a vector, and writes the vector to the destination SIMD&FP register. The first pair from the first source register is placed into the two lowest vector elements, with subsequent pairs taken alternately from each source register.
//
//go:linkname Vzip1P8 Vzip1P8
//go:noescape
func Vzip1P8(r *arm.Poly8X8, v0 *arm.Poly8X8, v1 *arm.Poly8X8)

// Zip vectors (primary). This instruction reads adjacent vector elements from the lower half of two source SIMD&FP registers as pairs, interleaves the pairs and places them into a vector, and writes the vector to the destination SIMD&FP register. The first pair from the first source register is placed into the two lowest vector elements, with subsequent pairs taken alternately from each source register.
//
//go:linkname Vzip1P16 Vzip1P16
//go:noescape
func Vzip1P16(r *arm.Poly16X4, v0 *arm.Poly16X4, v1 *arm.Poly16X4)

// Zip vectors (primary). This instruction reads adjacent vector elements from the lower half of two source SIMD&FP registers as pairs, interleaves the pairs and places them into a vector, and writes the vector to the destination SIMD&FP register. The first pair from the first source register is placed into the two lowest vector elements, with subsequent pairs taken alternately from each source register.
//
//go:linkname Vzip1QP8 Vzip1QP8
//go:noescape
func Vzip1QP8(r *arm.Poly8X16, v0 *arm.Poly8X16, v1 *arm.Poly8X16)

// Zip vectors (primary). This instruction reads adjacent vector elements from the lower half of two source SIMD&FP registers as pairs, interleaves the pairs and places them into a vector, and writes the vector to the destination SIMD&FP register. The first pair from the first source register is placed into the two lowest vector elements, with subsequent pairs taken alternately from each source register.
//
//go:linkname Vzip1QP64 Vzip1QP64
//go:noescape
func Vzip1QP64(r *arm.Poly64X2, v0 *arm.Poly64X2, v1 *arm.Poly64X2)

// Zip vectors (primary). This instruction reads adjacent vector elements from the lower half of two source SIMD&FP registers as pairs, interleaves the pairs and places them into a vector, and writes the vector to the destination SIMD&FP register. The first pair from the first source register is placed into the two lowest vector elements, with subsequent pairs taken alternately from each source register.
//
//go:linkname Vzip1QP16 Vzip1QP16
//go:noescape
func Vzip1QP16(r *arm.Poly16X8, v0 *arm.Poly16X8, v1 *arm.Poly16X8)

// Zip vectors (primary). This instruction reads adjacent vector elements from the lower half of two source SIMD&FP registers as pairs, interleaves the pairs and places them into a vector, and writes the vector to the destination SIMD&FP register. The first pair from the first source register is placed into the two lowest vector elements, with subsequent pairs taken alternately from each source register.
//
//go:linkname Vzip1QU8 Vzip1QU8
//go:noescape
func Vzip1QU8(r *arm.Uint8X16, v0 *arm.Uint8X16, v1 *arm.Uint8X16)

// Zip vectors (primary). This instruction reads adjacent vector elements from the lower half of two source SIMD&FP registers as pairs, interleaves the pairs and places them into a vector, and writes the vector to the destination SIMD&FP register. The first pair from the first source register is placed into the two lowest vector elements, with subsequent pairs taken alternately from each source register.
//
//go:linkname Vzip1QU32 Vzip1QU32
//go:noescape
func Vzip1QU32(r *arm.Uint32X4, v0 *arm.Uint32X4, v1 *arm.Uint32X4)

// Zip vectors (primary). This instruction reads adjacent vector elements from the lower half of two source SIMD&FP registers as pairs, interleaves the pairs and places them into a vector, and writes the vector to the destination SIMD&FP register. The first pair from the first source register is placed into the two lowest vector elements, with subsequent pairs taken alternately from each source register.
//
//go:linkname Vzip1QU64 Vzip1QU64
//go:noescape
func Vzip1QU64(r *arm.Uint64X2, v0 *arm.Uint64X2, v1 *arm.Uint64X2)

// Zip vectors (primary). This instruction reads adjacent vector elements from the lower half of two source SIMD&FP registers as pairs, interleaves the pairs and places them into a vector, and writes the vector to the destination SIMD&FP register. The first pair from the first source register is placed into the two lowest vector elements, with subsequent pairs taken alternately from each source register.
//
//go:linkname Vzip1QU16 Vzip1QU16
//go:noescape
func Vzip1QU16(r *arm.Uint16X8, v0 *arm.Uint16X8, v1 *arm.Uint16X8)

// Zip vectors (primary). This instruction reads adjacent vector elements from the lower half of two source SIMD&FP registers as pairs, interleaves the pairs and places them into a vector, and writes the vector to the destination SIMD&FP register. The first pair from the first source register is placed into the two lowest vector elements, with subsequent pairs taken alternately from each source register.
//
//go:linkname Vzip1QS8 Vzip1QS8
//go:noescape
func Vzip1QS8(r *arm.Int8X16, v0 *arm.Int8X16, v1 *arm.Int8X16)

// Zip vectors (primary). This instruction reads adjacent vector elements from the lower half of two source SIMD&FP registers as pairs, interleaves the pairs and places them into a vector, and writes the vector to the destination SIMD&FP register. The first pair from the first source register is placed into the two lowest vector elements, with subsequent pairs taken alternately from each source register.
//
//go:linkname Vzip1QF64 Vzip1QF64
//go:noescape
func Vzip1QF64(r *arm.Float64X2, v0 *arm.Float64X2, v1 *arm.Float64X2)

// Zip vectors (primary). This instruction reads adjacent vector elements from the lower half of two source SIMD&FP registers as pairs, interleaves the pairs and places them into a vector, and writes the vector to the destination SIMD&FP register. The first pair from the first source register is placed into the two lowest vector elements, with subsequent pairs taken alternately from each source register.
//
//go:linkname Vzip1QF32 Vzip1QF32
//go:noescape
func Vzip1QF32(r *arm.Float32X4, v0 *arm.Float32X4, v1 *arm.Float32X4)

// Zip vectors (primary). This instruction reads adjacent vector elements from the lower half of two source SIMD&FP registers as pairs, interleaves the pairs and places them into a vector, and writes the vector to the destination SIMD&FP register. The first pair from the first source register is placed into the two lowest vector elements, with subsequent pairs taken alternately from each source register.
//
//go:linkname Vzip1QS32 Vzip1QS32
//go:noescape
func Vzip1QS32(r *arm.Int32X4, v0 *arm.Int32X4, v1 *arm.Int32X4)

// Zip vectors (primary). This instruction reads adjacent vector elements from the lower half of two source SIMD&FP registers as pairs, interleaves the pairs and places them into a vector, and writes the vector to the destination SIMD&FP register. The first pair from the first source register is placed into the two lowest vector elements, with subsequent pairs taken alternately from each source register.
//
//go:linkname Vzip1QS64 Vzip1QS64
//go:noescape
func Vzip1QS64(r *arm.Int64X2, v0 *arm.Int64X2, v1 *arm.Int64X2)

// Zip vectors (primary). This instruction reads adjacent vector elements from the lower half of two source SIMD&FP registers as pairs, interleaves the pairs and places them into a vector, and writes the vector to the destination SIMD&FP register. The first pair from the first source register is placed into the two lowest vector elements, with subsequent pairs taken alternately from each source register.
//
//go:linkname Vzip1QS16 Vzip1QS16
//go:noescape
func Vzip1QS16(r *arm.Int16X8, v0 *arm.Int16X8, v1 *arm.Int16X8)

// Zip vectors (primary). This instruction reads adjacent vector elements from the lower half of two source SIMD&FP registers as pairs, interleaves the pairs and places them into a vector, and writes the vector to the destination SIMD&FP register. The first pair from the first source register is placed into the two lowest vector elements, with subsequent pairs taken alternately from each source register.
//
//go:linkname Vzip1U8 Vzip1U8
//go:noescape
func Vzip1U8(r *arm.Uint8X8, v0 *arm.Uint8X8, v1 *arm.Uint8X8)

// Zip vectors (primary). This instruction reads adjacent vector elements from the lower half of two source SIMD&FP registers as pairs, interleaves the pairs and places them into a vector, and writes the vector to the destination SIMD&FP register. The first pair from the first source register is placed into the two lowest vector elements, with subsequent pairs taken alternately from each source register.
//
//go:linkname Vzip1U32 Vzip1U32
//go:noescape
func Vzip1U32(r *arm.Uint32X2, v0 *arm.Uint32X2, v1 *arm.Uint32X2)

// Zip vectors (primary). This instruction reads adjacent vector elements from the lower half of two source SIMD&FP registers as pairs, interleaves the pairs and places them into a vector, and writes the vector to the destination SIMD&FP register. The first pair from the first source register is placed into the two lowest vector elements, with subsequent pairs taken alternately from each source register.
//
//go:linkname Vzip1U16 Vzip1U16
//go:noescape
func Vzip1U16(r *arm.Uint16X4, v0 *arm.Uint16X4, v1 *arm.Uint16X4)

// Zip vectors (primary). This instruction reads adjacent vector elements from the lower half of two source SIMD&FP registers as pairs, interleaves the pairs and places them into a vector, and writes the vector to the destination SIMD&FP register. The first pair from the first source register is placed into the two lowest vector elements, with subsequent pairs taken alternately from each source register.
//
//go:linkname Vzip1S8 Vzip1S8
//go:noescape
func Vzip1S8(r *arm.Int8X8, v0 *arm.Int8X8, v1 *arm.Int8X8)

// Zip vectors (primary). This instruction reads adjacent vector elements from the lower half of two source SIMD&FP registers as pairs, interleaves the pairs and places them into a vector, and writes the vector to the destination SIMD&FP register. The first pair from the first source register is placed into the two lowest vector elements, with subsequent pairs taken alternately from each source register.
//
//go:linkname Vzip1F32 Vzip1F32
//go:noescape
func Vzip1F32(r *arm.Float32X2, v0 *arm.Float32X2, v1 *arm.Float32X2)

// Zip vectors (primary). This instruction reads adjacent vector elements from the lower half of two source SIMD&FP registers as pairs, interleaves the pairs and places them into a vector, and writes the vector to the destination SIMD&FP register. The first pair from the first source register is placed into the two lowest vector elements, with subsequent pairs taken alternately from each source register.
//
//go:linkname Vzip1S32 Vzip1S32
//go:noescape
func Vzip1S32(r *arm.Int32X2, v0 *arm.Int32X2, v1 *arm.Int32X2)

// Zip vectors (primary). This instruction reads adjacent vector elements from the lower half of two source SIMD&FP registers as pairs, interleaves the pairs and places them into a vector, and writes the vector to the destination SIMD&FP register. The first pair from the first source register is placed into the two lowest vector elements, with subsequent pairs taken alternately from each source register.
//
//go:linkname Vzip1S16 Vzip1S16
//go:noescape
func Vzip1S16(r *arm.Int16X4, v0 *arm.Int16X4, v1 *arm.Int16X4)

// Zip vectors (secondary). This instruction reads adjacent vector elements from the upper half of two source SIMD&FP registers as pairs, interleaves the pairs and places them into a vector, and writes the vector to the destination SIMD&FP register. The first pair from the first source register is placed into the two lowest vector elements, with subsequent pairs taken alternately from each source register.
//
//go:linkname Vzip2P8 Vzip2P8
//go:noescape
func Vzip2P8(r *arm.Poly8X8, v0 *arm.Poly8X8, v1 *arm.Poly8X8)

// Zip vectors (secondary). This instruction reads adjacent vector elements from the upper half of two source SIMD&FP registers as pairs, interleaves the pairs and places them into a vector, and writes the vector to the destination SIMD&FP register. The first pair from the first source register is placed into the two lowest vector elements, with subsequent pairs taken alternately from each source register.
//
//go:linkname Vzip2P16 Vzip2P16
//go:noescape
func Vzip2P16(r *arm.Poly16X4, v0 *arm.Poly16X4, v1 *arm.Poly16X4)

// Zip vectors (secondary). This instruction reads adjacent vector elements from the upper half of two source SIMD&FP registers as pairs, interleaves the pairs and places them into a vector, and writes the vector to the destination SIMD&FP register. The first pair from the first source register is placed into the two lowest vector elements, with subsequent pairs taken alternately from each source register.
//
//go:linkname Vzip2QP8 Vzip2QP8
//go:noescape
func Vzip2QP8(r *arm.Poly8X16, v0 *arm.Poly8X16, v1 *arm.Poly8X16)

// Zip vectors (secondary). This instruction reads adjacent vector elements from the upper half of two source SIMD&FP registers as pairs, interleaves the pairs and places them into a vector, and writes the vector to the destination SIMD&FP register. The first pair from the first source register is placed into the two lowest vector elements, with subsequent pairs taken alternately from each source register.
//
//go:linkname Vzip2QP64 Vzip2QP64
//go:noescape
func Vzip2QP64(r *arm.Poly64X2, v0 *arm.Poly64X2, v1 *arm.Poly64X2)

// Zip vectors (secondary). This instruction reads adjacent vector elements from the upper half of two source SIMD&FP registers as pairs, interleaves the pairs and places them into a vector, and writes the vector to the destination SIMD&FP register. The first pair from the first source register is placed into the two lowest vector elements, with subsequent pairs taken alternately from each source register.
//
//go:linkname Vzip2QP16 Vzip2QP16
//go:noescape
func Vzip2QP16(r *arm.Poly16X8, v0 *arm.Poly16X8, v1 *arm.Poly16X8)

// Zip vectors (secondary). This instruction reads adjacent vector elements from the upper half of two source SIMD&FP registers as pairs, interleaves the pairs and places them into a vector, and writes the vector to the destination SIMD&FP register. The first pair from the first source register is placed into the two lowest vector elements, with subsequent pairs taken alternately from each source register.
//
//go:linkname Vzip2QU8 Vzip2QU8
//go:noescape
func Vzip2QU8(r *arm.Uint8X16, v0 *arm.Uint8X16, v1 *arm.Uint8X16)

// Zip vectors (secondary). This instruction reads adjacent vector elements from the upper half of two source SIMD&FP registers as pairs, interleaves the pairs and places them into a vector, and writes the vector to the destination SIMD&FP register. The first pair from the first source register is placed into the two lowest vector elements, with subsequent pairs taken alternately from each source register.
//
//go:linkname Vzip2QU32 Vzip2QU32
//go:noescape
func Vzip2QU32(r *arm.Uint32X4, v0 *arm.Uint32X4, v1 *arm.Uint32X4)

// Zip vectors (secondary). This instruction reads adjacent vector elements from the upper half of two source SIMD&FP registers as pairs, interleaves the pairs and places them into a vector, and writes the vector to the destination SIMD&FP register. The first pair from the first source register is placed into the two lowest vector elements, with subsequent pairs taken alternately from each source register.
//
//go:linkname Vzip2QU64 Vzip2QU64
//go:noescape
func Vzip2QU64(r *arm.Uint64X2, v0 *arm.Uint64X2, v1 *arm.Uint64X2)

// Zip vectors (secondary). This instruction reads adjacent vector elements from the upper half of two source SIMD&FP registers as pairs, interleaves the pairs and places them into a vector, and writes the vector to the destination SIMD&FP register. The first pair from the first source register is placed into the two lowest vector elements, with subsequent pairs taken alternately from each source register.
//
//go:linkname Vzip2QU16 Vzip2QU16
//go:noescape
func Vzip2QU16(r *arm.Uint16X8, v0 *arm.Uint16X8, v1 *arm.Uint16X8)

// Zip vectors (secondary). This instruction reads adjacent vector elements from the upper half of two source SIMD&FP registers as pairs, interleaves the pairs and places them into a vector, and writes the vector to the destination SIMD&FP register. The first pair from the first source register is placed into the two lowest vector elements, with subsequent pairs taken alternately from each source register.
//
//go:linkname Vzip2QS8 Vzip2QS8
//go:noescape
func Vzip2QS8(r *arm.Int8X16, v0 *arm.Int8X16, v1 *arm.Int8X16)

// Zip vectors (secondary). This instruction reads adjacent vector elements from the upper half of two source SIMD&FP registers as pairs, interleaves the pairs and places them into a vector, and writes the vector to the destination SIMD&FP register. The first pair from the first source register is placed into the two lowest vector elements, with subsequent pairs taken alternately from each source register.
//
//go:linkname Vzip2QF64 Vzip2QF64
//go:noescape
func Vzip2QF64(r *arm.Float64X2, v0 *arm.Float64X2, v1 *arm.Float64X2)

// Zip vectors (secondary). This instruction reads adjacent vector elements from the upper half of two source SIMD&FP registers as pairs, interleaves the pairs and places them into a vector, and writes the vector to the destination SIMD&FP register. The first pair from the first source register is placed into the two lowest vector elements, with subsequent pairs taken alternately from each source register.
//
//go:linkname Vzip2QF32 Vzip2QF32
//go:noescape
func Vzip2QF32(r *arm.Float32X4, v0 *arm.Float32X4, v1 *arm.Float32X4)

// Zip vectors (secondary). This instruction reads adjacent vector elements from the upper half of two source SIMD&FP registers as pairs, interleaves the pairs and places them into a vector, and writes the vector to the destination SIMD&FP register. The first pair from the first source register is placed into the two lowest vector elements, with subsequent pairs taken alternately from each source register.
//
//go:linkname Vzip2QS32 Vzip2QS32
//go:noescape
func Vzip2QS32(r *arm.Int32X4, v0 *arm.Int32X4, v1 *arm.Int32X4)

// Zip vectors (secondary). This instruction reads adjacent vector elements from the upper half of two source SIMD&FP registers as pairs, interleaves the pairs and places them into a vector, and writes the vector to the destination SIMD&FP register. The first pair from the first source register is placed into the two lowest vector elements, with subsequent pairs taken alternately from each source register.
//
//go:linkname Vzip2QS64 Vzip2QS64
//go:noescape
func Vzip2QS64(r *arm.Int64X2, v0 *arm.Int64X2, v1 *arm.Int64X2)

// Zip vectors (secondary). This instruction reads adjacent vector elements from the upper half of two source SIMD&FP registers as pairs, interleaves the pairs and places them into a vector, and writes the vector to the destination SIMD&FP register. The first pair from the first source register is placed into the two lowest vector elements, with subsequent pairs taken alternately from each source register.
//
//go:linkname Vzip2QS16 Vzip2QS16
//go:noescape
func Vzip2QS16(r *arm.Int16X8, v0 *arm.Int16X8, v1 *arm.Int16X8)

// Zip vectors (secondary). This instruction reads adjacent vector elements from the upper half of two source SIMD&FP registers as pairs, interleaves the pairs and places them into a vector, and writes the vector to the destination SIMD&FP register. The first pair from the first source register is placed into the two lowest vector elements, with subsequent pairs taken alternately from each source register.
//
//go:linkname Vzip2U8 Vzip2U8
//go:noescape
func Vzip2U8(r *arm.Uint8X8, v0 *arm.Uint8X8, v1 *arm.Uint8X8)

// Zip vectors (secondary). This instruction reads adjacent vector elements from the upper half of two source SIMD&FP registers as pairs, interleaves the pairs and places them into a vector, and writes the vector to the destination SIMD&FP register. The first pair from the first source register is placed into the two lowest vector elements, with subsequent pairs taken alternately from each source register.
//
//go:linkname Vzip2U32 Vzip2U32
//go:noescape
func Vzip2U32(r *arm.Uint32X2, v0 *arm.Uint32X2, v1 *arm.Uint32X2)

// Zip vectors (secondary). This instruction reads adjacent vector elements from the upper half of two source SIMD&FP registers as pairs, interleaves the pairs and places them into a vector, and writes the vector to the destination SIMD&FP register. The first pair from the first source register is placed into the two lowest vector elements, with subsequent pairs taken alternately from each source register.
//
//go:linkname Vzip2U16 Vzip2U16
//go:noescape
func Vzip2U16(r *arm.Uint16X4, v0 *arm.Uint16X4, v1 *arm.Uint16X4)

// Zip vectors (secondary). This instruction reads adjacent vector elements from the upper half of two source SIMD&FP registers as pairs, interleaves the pairs and places them into a vector, and writes the vector to the destination SIMD&FP register. The first pair from the first source register is placed into the two lowest vector elements, with subsequent pairs taken alternately from each source register.
//
//go:linkname Vzip2S8 Vzip2S8
//go:noescape
func Vzip2S8(r *arm.Int8X8, v0 *arm.Int8X8, v1 *arm.Int8X8)

// Zip vectors (secondary). This instruction reads adjacent vector elements from the upper half of two source SIMD&FP registers as pairs, interleaves the pairs and places them into a vector, and writes the vector to the destination SIMD&FP register. The first pair from the first source register is placed into the two lowest vector elements, with subsequent pairs taken alternately from each source register.
//
//go:linkname Vzip2F32 Vzip2F32
//go:noescape
func Vzip2F32(r *arm.Float32X2, v0 *arm.Float32X2, v1 *arm.Float32X2)

// Zip vectors (secondary). This instruction reads adjacent vector elements from the upper half of two source SIMD&FP registers as pairs, interleaves the pairs and places them into a vector, and writes the vector to the destination SIMD&FP register. The first pair from the first source register is placed into the two lowest vector elements, with subsequent pairs taken alternately from each source register.
//
//go:linkname Vzip2S32 Vzip2S32
//go:noescape
func Vzip2S32(r *arm.Int32X2, v0 *arm.Int32X2, v1 *arm.Int32X2)

// Zip vectors (secondary). This instruction reads adjacent vector elements from the upper half of two source SIMD&FP registers as pairs, interleaves the pairs and places them into a vector, and writes the vector to the destination SIMD&FP register. The first pair from the first source register is placed into the two lowest vector elements, with subsequent pairs taken alternately from each source register.
//
//go:linkname Vzip2S16 Vzip2S16
//go:noescape
func Vzip2S16(r *arm.Int16X4, v0 *arm.Int16X4, v1 *arm.Int16X4)

// Unsigned Absolute difference and Accumulate. This instruction subtracts the elements of the vector of the second source SIMD&FP register from the corresponding elements of the first source SIMD&FP register, and accumulates the absolute values of the results into the elements of the vector of the destination SIMD&FP register.
//
//go:linkname VabaqU8 VabaqU8
//go:noescape
func VabaqU8(r *arm.Uint8X16, v0 *arm.Uint8X16, v1 *arm.Uint8X16, v2 *arm.Uint8X16)

// Unsigned Absolute difference and Accumulate. This instruction subtracts the elements of the vector of the second source SIMD&FP register from the corresponding elements of the first source SIMD&FP register, and accumulates the absolute values of the results into the elements of the vector of the destination SIMD&FP register.
//
//go:linkname VabaqU32 VabaqU32
//go:noescape
func VabaqU32(r *arm.Uint32X4, v0 *arm.Uint32X4, v1 *arm.Uint32X4, v2 *arm.Uint32X4)

// Unsigned Absolute difference and Accumulate. This instruction subtracts the elements of the vector of the second source SIMD&FP register from the corresponding elements of the first source SIMD&FP register, and accumulates the absolute values of the results into the elements of the vector of the destination SIMD&FP register.
//
//go:linkname VabaqU16 VabaqU16
//go:noescape
func VabaqU16(r *arm.Uint16X8, v0 *arm.Uint16X8, v1 *arm.Uint16X8, v2 *arm.Uint16X8)

// Signed Absolute difference and Accumulate. This instruction subtracts the elements of the vector of the second source SIMD&FP register from the corresponding elements of the first source SIMD&FP register, and accumulates the absolute values of the results into the elements of the vector of the destination SIMD&FP register.
//
//go:linkname VabaqS8 VabaqS8
//go:noescape
func VabaqS8(r *arm.Int8X16, v0 *arm.Int8X16, v1 *arm.Int8X16, v2 *arm.Int8X16)

// Signed Absolute difference and Accumulate. This instruction subtracts the elements of the vector of the second source SIMD&FP register from the corresponding elements of the first source SIMD&FP register, and accumulates the absolute values of the results into the elements of the vector of the destination SIMD&FP register.
//
//go:linkname VabaqS32 VabaqS32
//go:noescape
func VabaqS32(r *arm.Int32X4, v0 *arm.Int32X4, v1 *arm.Int32X4, v2 *arm.Int32X4)

// Signed Absolute difference and Accumulate. This instruction subtracts the elements of the vector of the second source SIMD&FP register from the corresponding elements of the first source SIMD&FP register, and accumulates the absolute values of the results into the elements of the vector of the destination SIMD&FP register.
//
//go:linkname VabaqS16 VabaqS16
//go:noescape
func VabaqS16(r *arm.Int16X8, v0 *arm.Int16X8, v1 *arm.Int16X8, v2 *arm.Int16X8)

// Unsigned Absolute difference and Accumulate. This instruction subtracts the elements of the vector of the second source SIMD&FP register from the corresponding elements of the first source SIMD&FP register, and accumulates the absolute values of the results into the elements of the vector of the destination SIMD&FP register.
//
//go:linkname VabaU8 VabaU8
//go:noescape
func VabaU8(r *arm.Uint8X8, v0 *arm.Uint8X8, v1 *arm.Uint8X8, v2 *arm.Uint8X8)

// Unsigned Absolute difference and Accumulate. This instruction subtracts the elements of the vector of the second source SIMD&FP register from the corresponding elements of the first source SIMD&FP register, and accumulates the absolute values of the results into the elements of the vector of the destination SIMD&FP register.
//
//go:linkname VabaU32 VabaU32
//go:noescape
func VabaU32(r *arm.Uint32X2, v0 *arm.Uint32X2, v1 *arm.Uint32X2, v2 *arm.Uint32X2)

// Unsigned Absolute difference and Accumulate. This instruction subtracts the elements of the vector of the second source SIMD&FP register from the corresponding elements of the first source SIMD&FP register, and accumulates the absolute values of the results into the elements of the vector of the destination SIMD&FP register.
//
//go:linkname VabaU16 VabaU16
//go:noescape
func VabaU16(r *arm.Uint16X4, v0 *arm.Uint16X4, v1 *arm.Uint16X4, v2 *arm.Uint16X4)

// Signed Absolute difference and Accumulate. This instruction subtracts the elements of the vector of the second source SIMD&FP register from the corresponding elements of the first source SIMD&FP register, and accumulates the absolute values of the results into the elements of the vector of the destination SIMD&FP register.
//
//go:linkname VabaS8 VabaS8
//go:noescape
func VabaS8(r *arm.Int8X8, v0 *arm.Int8X8, v1 *arm.Int8X8, v2 *arm.Int8X8)

// Signed Absolute difference and Accumulate. This instruction subtracts the elements of the vector of the second source SIMD&FP register from the corresponding elements of the first source SIMD&FP register, and accumulates the absolute values of the results into the elements of the vector of the destination SIMD&FP register.
//
//go:linkname VabaS32 VabaS32
//go:noescape
func VabaS32(r *arm.Int32X2, v0 *arm.Int32X2, v1 *arm.Int32X2, v2 *arm.Int32X2)

// Signed Absolute difference and Accumulate. This instruction subtracts the elements of the vector of the second source SIMD&FP register from the corresponding elements of the first source SIMD&FP register, and accumulates the absolute values of the results into the elements of the vector of the destination SIMD&FP register.
//
//go:linkname VabaS16 VabaS16
//go:noescape
func VabaS16(r *arm.Int16X4, v0 *arm.Int16X4, v1 *arm.Int16X4, v2 *arm.Int16X4)

// Unsigned Absolute Difference Long. This instruction subtracts the vector elements in the lower or upper half of the second source SIMD&FP register from the corresponding vector elements of the first source SIMD&FP register, places the absolute value of the result into a vector, and writes the vector to the destination SIMD&FP register. The destination vector elements are twice as long as the source vector elements. All the values in this instruction are unsigned integer values.
//
//go:linkname VabdlU8 VabdlU8
//go:noescape
func VabdlU8(r *arm.Uint16X8, v0 *arm.Uint8X8, v1 *arm.Uint8X8)

// Unsigned Absolute Difference Long. This instruction subtracts the vector elements in the lower or upper half of the second source SIMD&FP register from the corresponding vector elements of the first source SIMD&FP register, places the absolute value of the result into a vector, and writes the vector to the destination SIMD&FP register. The destination vector elements are twice as long as the source vector elements. All the values in this instruction are unsigned integer values.
//
//go:linkname VabdlU32 VabdlU32
//go:noescape
func VabdlU32(r *arm.Uint64X2, v0 *arm.Uint32X2, v1 *arm.Uint32X2)

// Unsigned Absolute Difference Long. This instruction subtracts the vector elements in the lower or upper half of the second source SIMD&FP register from the corresponding vector elements of the first source SIMD&FP register, places the absolute value of the result into a vector, and writes the vector to the destination SIMD&FP register. The destination vector elements are twice as long as the source vector elements. All the values in this instruction are unsigned integer values.
//
//go:linkname VabdlU16 VabdlU16
//go:noescape
func VabdlU16(r *arm.Uint32X4, v0 *arm.Uint16X4, v1 *arm.Uint16X4)

// Signed Absolute Difference Long. This instruction subtracts the vector elements of the second source SIMD&FP register from the corresponding vector elements of the first source SIMD&FP register, places the absolute value of the results into a vector, and writes the vector to the lower or upper half of the destination SIMD&FP register. The destination vector elements are twice as long as the source vector elements.
//
//go:linkname VabdlS8 VabdlS8
//go:noescape
func VabdlS8(r *arm.Int16X8, v0 *arm.Int8X8, v1 *arm.Int8X8)

// Signed Absolute Difference Long. This instruction subtracts the vector elements of the second source SIMD&FP register from the corresponding vector elements of the first source SIMD&FP register, places the absolute value of the results into a vector, and writes the vector to the lower or upper half of the destination SIMD&FP register. The destination vector elements are twice as long as the source vector elements.
//
//go:linkname VabdlS32 VabdlS32
//go:noescape
func VabdlS32(r *arm.Int64X2, v0 *arm.Int32X2, v1 *arm.Int32X2)

// Signed Absolute Difference Long. This instruction subtracts the vector elements of the second source SIMD&FP register from the corresponding vector elements of the first source SIMD&FP register, places the absolute value of the results into a vector, and writes the vector to the lower or upper half of the destination SIMD&FP register. The destination vector elements are twice as long as the source vector elements.
//
//go:linkname VabdlS16 VabdlS16
//go:noescape
func VabdlS16(r *arm.Int32X4, v0 *arm.Int16X4, v1 *arm.Int16X4)

// Unsigned Add Long (vector). This instruction adds each vector element in the lower or upper half of the first source SIMD&FP register to the corresponding vector element of the second source SIMD&FP register, places the result into a vector, and writes the vector to the destination SIMD&FP register. The destination vector elements are twice as long as the source vector elements. All the values in this instruction are unsigned integer values.
//
//go:linkname VaddlU8 VaddlU8
//go:noescape
func VaddlU8(r *arm.Uint16X8, v0 *arm.Uint8X8, v1 *arm.Uint8X8)

// Unsigned Add Long (vector). This instruction adds each vector element in the lower or upper half of the first source SIMD&FP register to the corresponding vector element of the second source SIMD&FP register, places the result into a vector, and writes the vector to the destination SIMD&FP register. The destination vector elements are twice as long as the source vector elements. All the values in this instruction are unsigned integer values.
//
//go:linkname VaddlU32 VaddlU32
//go:noescape
func VaddlU32(r *arm.Uint64X2, v0 *arm.Uint32X2, v1 *arm.Uint32X2)

// Unsigned Add Long (vector). This instruction adds each vector element in the lower or upper half of the first source SIMD&FP register to the corresponding vector element of the second source SIMD&FP register, places the result into a vector, and writes the vector to the destination SIMD&FP register. The destination vector elements are twice as long as the source vector elements. All the values in this instruction are unsigned integer values.
//
//go:linkname VaddlU16 VaddlU16
//go:noescape
func VaddlU16(r *arm.Uint32X4, v0 *arm.Uint16X4, v1 *arm.Uint16X4)

// Signed Add Long (vector). This instruction adds each vector element in the lower or upper half of the first source SIMD&FP register to the corresponding vector element of the second source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register. The destination vector elements are twice as long as the source vector elements. All the values in this instruction are signed integer values.
//
//go:linkname VaddlS8 VaddlS8
//go:noescape
func VaddlS8(r *arm.Int16X8, v0 *arm.Int8X8, v1 *arm.Int8X8)

// Signed Add Long (vector). This instruction adds each vector element in the lower or upper half of the first source SIMD&FP register to the corresponding vector element of the second source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register. The destination vector elements are twice as long as the source vector elements. All the values in this instruction are signed integer values.
//
//go:linkname VaddlS32 VaddlS32
//go:noescape
func VaddlS32(r *arm.Int64X2, v0 *arm.Int32X2, v1 *arm.Int32X2)

// Signed Add Long (vector). This instruction adds each vector element in the lower or upper half of the first source SIMD&FP register to the corresponding vector element of the second source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register. The destination vector elements are twice as long as the source vector elements. All the values in this instruction are signed integer values.
//
//go:linkname VaddlS16 VaddlS16
//go:noescape
func VaddlS16(r *arm.Int32X4, v0 *arm.Int16X4, v1 *arm.Int16X4)

// Unsigned Add Wide. This instruction adds the vector elements of the first source SIMD&FP register to the corresponding vector elements in the lower or upper half of the second source SIMD&FP register, places the result in a vector, and writes the vector to the SIMD&FP destination register. The vector elements of the destination register and the first source register are twice as long as the vector elements of the second source register. All the values in this instruction are unsigned integer values.
//
//go:linkname VaddwU8 VaddwU8
//go:noescape
func VaddwU8(r *arm.Uint16X8, v0 *arm.Uint16X8, v1 *arm.Uint8X8)

// Unsigned Add Wide. This instruction adds the vector elements of the first source SIMD&FP register to the corresponding vector elements in the lower or upper half of the second source SIMD&FP register, places the result in a vector, and writes the vector to the SIMD&FP destination register. The vector elements of the destination register and the first source register are twice as long as the vector elements of the second source register. All the values in this instruction are unsigned integer values.
//
//go:linkname VaddwU32 VaddwU32
//go:noescape
func VaddwU32(r *arm.Uint64X2, v0 *arm.Uint64X2, v1 *arm.Uint32X2)

// Unsigned Add Wide. This instruction adds the vector elements of the first source SIMD&FP register to the corresponding vector elements in the lower or upper half of the second source SIMD&FP register, places the result in a vector, and writes the vector to the SIMD&FP destination register. The vector elements of the destination register and the first source register are twice as long as the vector elements of the second source register. All the values in this instruction are unsigned integer values.
//
//go:linkname VaddwU16 VaddwU16
//go:noescape
func VaddwU16(r *arm.Uint32X4, v0 *arm.Uint32X4, v1 *arm.Uint16X4)

// Signed Add Wide. This instruction adds vector elements of the first source SIMD&FP register to the corresponding vector elements in the lower or upper half of the second source SIMD&FP register, places the results in a vector, and writes the vector to the SIMD&FP destination register.
//
//go:linkname VaddwS8 VaddwS8
//go:noescape
func VaddwS8(r *arm.Int16X8, v0 *arm.Int16X8, v1 *arm.Int8X8)

// Signed Add Wide. This instruction adds vector elements of the first source SIMD&FP register to the corresponding vector elements in the lower or upper half of the second source SIMD&FP register, places the results in a vector, and writes the vector to the SIMD&FP destination register.
//
//go:linkname VaddwS32 VaddwS32
//go:noescape
func VaddwS32(r *arm.Int64X2, v0 *arm.Int64X2, v1 *arm.Int32X2)

// Signed Add Wide. This instruction adds vector elements of the first source SIMD&FP register to the corresponding vector elements in the lower or upper half of the second source SIMD&FP register, places the results in a vector, and writes the vector to the SIMD&FP destination register.
//
//go:linkname VaddwS16 VaddwS16
//go:noescape
func VaddwS16(r *arm.Int32X4, v0 *arm.Int32X4, v1 *arm.Int16X4)

// Unsigned Multiply-Add Long (vector). This instruction multiplies the vector elements in the lower or upper half of the first source SIMD&FP register by the corresponding vector elements of the second source SIMD&FP register, and accumulates the results with the vector elements of the destination SIMD&FP register. The destination vector elements are twice as long as the elements that are multiplied.
//
//go:linkname VmlalU8 VmlalU8
//go:noescape
func VmlalU8(r *arm.Uint16X8, v0 *arm.Uint16X8, v1 *arm.Uint8X8, v2 *arm.Uint8X8)

// Unsigned Multiply-Add Long (vector). This instruction multiplies the vector elements in the lower or upper half of the first source SIMD&FP register by the corresponding vector elements of the second source SIMD&FP register, and accumulates the results with the vector elements of the destination SIMD&FP register. The destination vector elements are twice as long as the elements that are multiplied.
//
//go:linkname VmlalU32 VmlalU32
//go:noescape
func VmlalU32(r *arm.Uint64X2, v0 *arm.Uint64X2, v1 *arm.Uint32X2, v2 *arm.Uint32X2)

// Unsigned Multiply-Add Long (vector). This instruction multiplies the vector elements in the lower or upper half of the first source SIMD&FP register by the corresponding vector elements of the second source SIMD&FP register, and accumulates the results with the vector elements of the destination SIMD&FP register. The destination vector elements are twice as long as the elements that are multiplied.
//
//go:linkname VmlalU16 VmlalU16
//go:noescape
func VmlalU16(r *arm.Uint32X4, v0 *arm.Uint32X4, v1 *arm.Uint16X4, v2 *arm.Uint16X4)

// Signed Multiply-Add Long (vector). This instruction multiplies corresponding signed integer values in the lower or upper half of the vectors of the two source SIMD&FP registers, and accumulates the results with the vector elements of the destination SIMD&FP register. The destination vector elements are twice as long as the elements that are multiplied.
//
//go:linkname VmlalS8 VmlalS8
//go:noescape
func VmlalS8(r *arm.Int16X8, v0 *arm.Int16X8, v1 *arm.Int8X8, v2 *arm.Int8X8)

// Signed Multiply-Add Long (vector). This instruction multiplies corresponding signed integer values in the lower or upper half of the vectors of the two source SIMD&FP registers, and accumulates the results with the vector elements of the destination SIMD&FP register. The destination vector elements are twice as long as the elements that are multiplied.
//
//go:linkname VmlalS32 VmlalS32
//go:noescape
func VmlalS32(r *arm.Int64X2, v0 *arm.Int64X2, v1 *arm.Int32X2, v2 *arm.Int32X2)

// Signed Multiply-Add Long (vector). This instruction multiplies corresponding signed integer values in the lower or upper half of the vectors of the two source SIMD&FP registers, and accumulates the results with the vector elements of the destination SIMD&FP register. The destination vector elements are twice as long as the elements that are multiplied.
//
//go:linkname VmlalS16 VmlalS16
//go:noescape
func VmlalS16(r *arm.Int32X4, v0 *arm.Int32X4, v1 *arm.Int16X4, v2 *arm.Int16X4)

// Vector widening multiply accumulate with scalar
//
//go:linkname VmlalNU32 VmlalNU32
//go:noescape
func VmlalNU32(r *arm.Uint64X2, v0 *arm.Uint64X2, v1 *arm.Uint32X2, v2 *arm.Uint32)

// Vector widening multiply accumulate with scalar
//
//go:linkname VmlalNU16 VmlalNU16
//go:noescape
func VmlalNU16(r *arm.Uint32X4, v0 *arm.Uint32X4, v1 *arm.Uint16X4, v2 *arm.Uint16)

// Vector widening multiply accumulate with scalar
//
//go:linkname VmlalNS32 VmlalNS32
//go:noescape
func VmlalNS32(r *arm.Int64X2, v0 *arm.Int64X2, v1 *arm.Int32X2, v2 *arm.Int32)

// Vector widening multiply accumulate with scalar
//
//go:linkname VmlalNS16 VmlalNS16
//go:noescape
func VmlalNS16(r *arm.Int32X4, v0 *arm.Int32X4, v1 *arm.Int16X4, v2 *arm.Int16)

// Unsigned Multiply-Subtract Long (vector). This instruction multiplies corresponding vector elements in the lower or upper half of the two source SIMD&FP registers, and subtracts the results from the vector elements of the destination SIMD&FP register. The destination vector elements are twice as long as the elements that are multiplied. All the values in this instruction are unsigned integer values.
//
//go:linkname VmlslU8 VmlslU8
//go:noescape
func VmlslU8(r *arm.Uint16X8, v0 *arm.Uint16X8, v1 *arm.Uint8X8, v2 *arm.Uint8X8)

// Unsigned Multiply-Subtract Long (vector). This instruction multiplies corresponding vector elements in the lower or upper half of the two source SIMD&FP registers, and subtracts the results from the vector elements of the destination SIMD&FP register. The destination vector elements are twice as long as the elements that are multiplied. All the values in this instruction are unsigned integer values.
//
//go:linkname VmlslU32 VmlslU32
//go:noescape
func VmlslU32(r *arm.Uint64X2, v0 *arm.Uint64X2, v1 *arm.Uint32X2, v2 *arm.Uint32X2)

// Unsigned Multiply-Subtract Long (vector). This instruction multiplies corresponding vector elements in the lower or upper half of the two source SIMD&FP registers, and subtracts the results from the vector elements of the destination SIMD&FP register. The destination vector elements are twice as long as the elements that are multiplied. All the values in this instruction are unsigned integer values.
//
//go:linkname VmlslU16 VmlslU16
//go:noescape
func VmlslU16(r *arm.Uint32X4, v0 *arm.Uint32X4, v1 *arm.Uint16X4, v2 *arm.Uint16X4)

// Signed Multiply-Subtract Long (vector). This instruction multiplies corresponding signed integer values in the lower or upper half of the vectors of the two source SIMD&FP registers, and subtracts the results from the vector elements of the destination SIMD&FP register. The destination vector elements are twice as long as the elements that are multiplied.
//
//go:linkname VmlslS8 VmlslS8
//go:noescape
func VmlslS8(r *arm.Int16X8, v0 *arm.Int16X8, v1 *arm.Int8X8, v2 *arm.Int8X8)

// Signed Multiply-Subtract Long (vector). This instruction multiplies corresponding signed integer values in the lower or upper half of the vectors of the two source SIMD&FP registers, and subtracts the results from the vector elements of the destination SIMD&FP register. The destination vector elements are twice as long as the elements that are multiplied.
//
//go:linkname VmlslS32 VmlslS32
//go:noescape
func VmlslS32(r *arm.Int64X2, v0 *arm.Int64X2, v1 *arm.Int32X2, v2 *arm.Int32X2)

// Signed Multiply-Subtract Long (vector). This instruction multiplies corresponding signed integer values in the lower or upper half of the vectors of the two source SIMD&FP registers, and subtracts the results from the vector elements of the destination SIMD&FP register. The destination vector elements are twice as long as the elements that are multiplied.
//
//go:linkname VmlslS16 VmlslS16
//go:noescape
func VmlslS16(r *arm.Int32X4, v0 *arm.Int32X4, v1 *arm.Int16X4, v2 *arm.Int16X4)

// Vector widening multiply subtract with scalar
//
//go:linkname VmlslNU32 VmlslNU32
//go:noescape
func VmlslNU32(r *arm.Uint64X2, v0 *arm.Uint64X2, v1 *arm.Uint32X2, v2 *arm.Uint32)

// Vector widening multiply subtract with scalar
//
//go:linkname VmlslNU16 VmlslNU16
//go:noescape
func VmlslNU16(r *arm.Uint32X4, v0 *arm.Uint32X4, v1 *arm.Uint16X4, v2 *arm.Uint16)

// Vector widening multiply subtract with scalar
//
//go:linkname VmlslNS32 VmlslNS32
//go:noescape
func VmlslNS32(r *arm.Int64X2, v0 *arm.Int64X2, v1 *arm.Int32X2, v2 *arm.Int32)

// Vector widening multiply subtract with scalar
//
//go:linkname VmlslNS16 VmlslNS16
//go:noescape
func VmlslNS16(r *arm.Int32X4, v0 *arm.Int32X4, v1 *arm.Int16X4, v2 *arm.Int16)

// Signed Saturating Rounding Doubling Multiply Subtract returning High Half (vector). This instruction multiplies the vector elements of the first source SIMD&FP register with the corresponding vector elements of the second source SIMD&FP register without saturating the multiply results, doubles the results, and subtracts the most significant half of the final results from the vector elements of the destination SIMD&FP register. The results are rounded.
//
//go:linkname VqrdmlahsS32 VqrdmlahsS32
//go:noescape
func VqrdmlahsS32(r *arm.Int32, v0 *arm.Int32, v1 *arm.Int32, v2 *arm.Int32)

// Signed Saturating Rounding Doubling Multiply Subtract returning High Half (vector). This instruction multiplies the vector elements of the first source SIMD&FP register with the corresponding vector elements of the second source SIMD&FP register without saturating the multiply results, doubles the results, and subtracts the most significant half of the final results from the vector elements of the destination SIMD&FP register. The results are rounded.
//
//go:linkname VqrdmlahhS16 VqrdmlahhS16
//go:noescape
func VqrdmlahhS16(r *arm.Int16, v0 *arm.Int16, v1 *arm.Int16, v2 *arm.Int16)

// Signed Saturating Rounding Doubling Multiply Subtract returning High Half (vector). This instruction multiplies the vector elements of the first source SIMD&FP register with the corresponding vector elements of the second source SIMD&FP register without saturating the multiply results, doubles the results, and subtracts the most significant half of the final results from the vector elements of the destination SIMD&FP register. The results are rounded.
//
//go:linkname VqrdmlshsS32 VqrdmlshsS32
//go:noescape
func VqrdmlshsS32(r *arm.Int32, v0 *arm.Int32, v1 *arm.Int32, v2 *arm.Int32)

// Signed Saturating Rounding Doubling Multiply Subtract returning High Half (vector). This instruction multiplies the vector elements of the first source SIMD&FP register with the corresponding vector elements of the second source SIMD&FP register without saturating the multiply results, doubles the results, and subtracts the most significant half of the final results from the vector elements of the destination SIMD&FP register. The results are rounded.
//
//go:linkname VqrdmlshhS16 VqrdmlshhS16
//go:noescape
func VqrdmlshhS16(r *arm.Int16, v0 *arm.Int16, v1 *arm.Int16, v2 *arm.Int16)

// Unsigned Absolute Difference Long. This instruction subtracts the vector elements in the lower or upper half of the second source SIMD&FP register from the corresponding vector elements of the first source SIMD&FP register, places the absolute value of the result into a vector, and writes the vector to the destination SIMD&FP register. The destination vector elements are twice as long as the source vector elements. All the values in this instruction are unsigned integer values.
//
//go:linkname VabdlHighU8 VabdlHighU8
//go:noescape
func VabdlHighU8(r *arm.Uint16X8, v0 *arm.Uint8X16, v1 *arm.Uint8X16)

// Unsigned Absolute Difference Long. This instruction subtracts the vector elements in the lower or upper half of the second source SIMD&FP register from the corresponding vector elements of the first source SIMD&FP register, places the absolute value of the result into a vector, and writes the vector to the destination SIMD&FP register. The destination vector elements are twice as long as the source vector elements. All the values in this instruction are unsigned integer values.
//
//go:linkname VabdlHighU32 VabdlHighU32
//go:noescape
func VabdlHighU32(r *arm.Uint64X2, v0 *arm.Uint32X4, v1 *arm.Uint32X4)

// Unsigned Absolute Difference Long. This instruction subtracts the vector elements in the lower or upper half of the second source SIMD&FP register from the corresponding vector elements of the first source SIMD&FP register, places the absolute value of the result into a vector, and writes the vector to the destination SIMD&FP register. The destination vector elements are twice as long as the source vector elements. All the values in this instruction are unsigned integer values.
//
//go:linkname VabdlHighU16 VabdlHighU16
//go:noescape
func VabdlHighU16(r *arm.Uint32X4, v0 *arm.Uint16X8, v1 *arm.Uint16X8)

// Signed Absolute Difference Long. This instruction subtracts the vector elements of the second source SIMD&FP register from the corresponding vector elements of the first source SIMD&FP register, places the absolute value of the results into a vector, and writes the vector to the lower or upper half of the destination SIMD&FP register. The destination vector elements are twice as long as the source vector elements.
//
//go:linkname VabdlHighS8 VabdlHighS8
//go:noescape
func VabdlHighS8(r *arm.Int16X8, v0 *arm.Int8X16, v1 *arm.Int8X16)

// Signed Absolute Difference Long. This instruction subtracts the vector elements of the second source SIMD&FP register from the corresponding vector elements of the first source SIMD&FP register, places the absolute value of the results into a vector, and writes the vector to the lower or upper half of the destination SIMD&FP register. The destination vector elements are twice as long as the source vector elements.
//
//go:linkname VabdlHighS32 VabdlHighS32
//go:noescape
func VabdlHighS32(r *arm.Int64X2, v0 *arm.Int32X4, v1 *arm.Int32X4)

// Signed Absolute Difference Long. This instruction subtracts the vector elements of the second source SIMD&FP register from the corresponding vector elements of the first source SIMD&FP register, places the absolute value of the results into a vector, and writes the vector to the lower or upper half of the destination SIMD&FP register. The destination vector elements are twice as long as the source vector elements.
//
//go:linkname VabdlHighS16 VabdlHighS16
//go:noescape
func VabdlHighS16(r *arm.Int32X4, v0 *arm.Int16X8, v1 *arm.Int16X8)

// Unsigned Add Long (vector). This instruction adds each vector element in the lower or upper half of the first source SIMD&FP register to the corresponding vector element of the second source SIMD&FP register, places the result into a vector, and writes the vector to the destination SIMD&FP register. The destination vector elements are twice as long as the source vector elements. All the values in this instruction are unsigned integer values.
//
//go:linkname VaddlHighU8 VaddlHighU8
//go:noescape
func VaddlHighU8(r *arm.Uint16X8, v0 *arm.Uint8X16, v1 *arm.Uint8X16)

// Unsigned Add Long (vector). This instruction adds each vector element in the lower or upper half of the first source SIMD&FP register to the corresponding vector element of the second source SIMD&FP register, places the result into a vector, and writes the vector to the destination SIMD&FP register. The destination vector elements are twice as long as the source vector elements. All the values in this instruction are unsigned integer values.
//
//go:linkname VaddlHighU32 VaddlHighU32
//go:noescape
func VaddlHighU32(r *arm.Uint64X2, v0 *arm.Uint32X4, v1 *arm.Uint32X4)

// Unsigned Add Long (vector). This instruction adds each vector element in the lower or upper half of the first source SIMD&FP register to the corresponding vector element of the second source SIMD&FP register, places the result into a vector, and writes the vector to the destination SIMD&FP register. The destination vector elements are twice as long as the source vector elements. All the values in this instruction are unsigned integer values.
//
//go:linkname VaddlHighU16 VaddlHighU16
//go:noescape
func VaddlHighU16(r *arm.Uint32X4, v0 *arm.Uint16X8, v1 *arm.Uint16X8)

// Signed Add Long (vector). This instruction adds each vector element in the lower or upper half of the first source SIMD&FP register to the corresponding vector element of the second source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register. The destination vector elements are twice as long as the source vector elements. All the values in this instruction are signed integer values.
//
//go:linkname VaddlHighS8 VaddlHighS8
//go:noescape
func VaddlHighS8(r *arm.Int16X8, v0 *arm.Int8X16, v1 *arm.Int8X16)

// Signed Add Long (vector). This instruction adds each vector element in the lower or upper half of the first source SIMD&FP register to the corresponding vector element of the second source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register. The destination vector elements are twice as long as the source vector elements. All the values in this instruction are signed integer values.
//
//go:linkname VaddlHighS32 VaddlHighS32
//go:noescape
func VaddlHighS32(r *arm.Int64X2, v0 *arm.Int32X4, v1 *arm.Int32X4)

// Signed Add Long (vector). This instruction adds each vector element in the lower or upper half of the first source SIMD&FP register to the corresponding vector element of the second source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register. The destination vector elements are twice as long as the source vector elements. All the values in this instruction are signed integer values.
//
//go:linkname VaddlHighS16 VaddlHighS16
//go:noescape
func VaddlHighS16(r *arm.Int32X4, v0 *arm.Int16X8, v1 *arm.Int16X8)

// Unsigned Add Wide. This instruction adds the vector elements of the first source SIMD&FP register to the corresponding vector elements in the lower or upper half of the second source SIMD&FP register, places the result in a vector, and writes the vector to the SIMD&FP destination register. The vector elements of the destination register and the first source register are twice as long as the vector elements of the second source register. All the values in this instruction are unsigned integer values.
//
//go:linkname VaddwHighU8 VaddwHighU8
//go:noescape
func VaddwHighU8(r *arm.Uint16X8, v0 *arm.Uint16X8, v1 *arm.Uint8X16)

// Unsigned Add Wide. This instruction adds the vector elements of the first source SIMD&FP register to the corresponding vector elements in the lower or upper half of the second source SIMD&FP register, places the result in a vector, and writes the vector to the SIMD&FP destination register. The vector elements of the destination register and the first source register are twice as long as the vector elements of the second source register. All the values in this instruction are unsigned integer values.
//
//go:linkname VaddwHighU32 VaddwHighU32
//go:noescape
func VaddwHighU32(r *arm.Uint64X2, v0 *arm.Uint64X2, v1 *arm.Uint32X4)

// Unsigned Add Wide. This instruction adds the vector elements of the first source SIMD&FP register to the corresponding vector elements in the lower or upper half of the second source SIMD&FP register, places the result in a vector, and writes the vector to the SIMD&FP destination register. The vector elements of the destination register and the first source register are twice as long as the vector elements of the second source register. All the values in this instruction are unsigned integer values.
//
//go:linkname VaddwHighU16 VaddwHighU16
//go:noescape
func VaddwHighU16(r *arm.Uint32X4, v0 *arm.Uint32X4, v1 *arm.Uint16X8)

// Signed Add Wide. This instruction adds vector elements of the first source SIMD&FP register to the corresponding vector elements in the lower or upper half of the second source SIMD&FP register, places the results in a vector, and writes the vector to the SIMD&FP destination register.
//
//go:linkname VaddwHighS8 VaddwHighS8
//go:noescape
func VaddwHighS8(r *arm.Int16X8, v0 *arm.Int16X8, v1 *arm.Int8X16)

// Signed Add Wide. This instruction adds vector elements of the first source SIMD&FP register to the corresponding vector elements in the lower or upper half of the second source SIMD&FP register, places the results in a vector, and writes the vector to the SIMD&FP destination register.
//
//go:linkname VaddwHighS32 VaddwHighS32
//go:noescape
func VaddwHighS32(r *arm.Int64X2, v0 *arm.Int64X2, v1 *arm.Int32X4)

// Signed Add Wide. This instruction adds vector elements of the first source SIMD&FP register to the corresponding vector elements in the lower or upper half of the second source SIMD&FP register, places the results in a vector, and writes the vector to the SIMD&FP destination register.
//
//go:linkname VaddwHighS16 VaddwHighS16
//go:noescape
func VaddwHighS16(r *arm.Int32X4, v0 *arm.Int32X4, v1 *arm.Int16X8)

// Unsigned Multiply-Add Long (vector). This instruction multiplies the vector elements in the lower or upper half of the first source SIMD&FP register by the corresponding vector elements of the second source SIMD&FP register, and accumulates the results with the vector elements of the destination SIMD&FP register. The destination vector elements are twice as long as the elements that are multiplied.
//
//go:linkname VmlalHighU8 VmlalHighU8
//go:noescape
func VmlalHighU8(r *arm.Uint16X8, v0 *arm.Uint16X8, v1 *arm.Uint8X16, v2 *arm.Uint8X16)

// Unsigned Multiply-Add Long (vector). This instruction multiplies the vector elements in the lower or upper half of the first source SIMD&FP register by the corresponding vector elements of the second source SIMD&FP register, and accumulates the results with the vector elements of the destination SIMD&FP register. The destination vector elements are twice as long as the elements that are multiplied.
//
//go:linkname VmlalHighU32 VmlalHighU32
//go:noescape
func VmlalHighU32(r *arm.Uint64X2, v0 *arm.Uint64X2, v1 *arm.Uint32X4, v2 *arm.Uint32X4)

// Unsigned Multiply-Add Long (vector). This instruction multiplies the vector elements in the lower or upper half of the first source SIMD&FP register by the corresponding vector elements of the second source SIMD&FP register, and accumulates the results with the vector elements of the destination SIMD&FP register. The destination vector elements are twice as long as the elements that are multiplied.
//
//go:linkname VmlalHighU16 VmlalHighU16
//go:noescape
func VmlalHighU16(r *arm.Uint32X4, v0 *arm.Uint32X4, v1 *arm.Uint16X8, v2 *arm.Uint16X8)

// Signed Multiply-Add Long (vector). This instruction multiplies corresponding signed integer values in the lower or upper half of the vectors of the two source SIMD&FP registers, and accumulates the results with the vector elements of the destination SIMD&FP register. The destination vector elements are twice as long as the elements that are multiplied.
//
//go:linkname VmlalHighS8 VmlalHighS8
//go:noescape
func VmlalHighS8(r *arm.Int16X8, v0 *arm.Int16X8, v1 *arm.Int8X16, v2 *arm.Int8X16)

// Signed Multiply-Add Long (vector). This instruction multiplies corresponding signed integer values in the lower or upper half of the vectors of the two source SIMD&FP registers, and accumulates the results with the vector elements of the destination SIMD&FP register. The destination vector elements are twice as long as the elements that are multiplied.
//
//go:linkname VmlalHighS32 VmlalHighS32
//go:noescape
func VmlalHighS32(r *arm.Int64X2, v0 *arm.Int64X2, v1 *arm.Int32X4, v2 *arm.Int32X4)

// Signed Multiply-Add Long (vector). This instruction multiplies corresponding signed integer values in the lower or upper half of the vectors of the two source SIMD&FP registers, and accumulates the results with the vector elements of the destination SIMD&FP register. The destination vector elements are twice as long as the elements that are multiplied.
//
//go:linkname VmlalHighS16 VmlalHighS16
//go:noescape
func VmlalHighS16(r *arm.Int32X4, v0 *arm.Int32X4, v1 *arm.Int16X8, v2 *arm.Int16X8)

// Unsigned Multiply-Add Long (vector). This instruction multiplies the vector elements in the lower or upper half of the first source SIMD&FP register by the corresponding vector elements of the second source SIMD&FP register, and accumulates the results with the vector elements of the destination SIMD&FP register. The destination vector elements are twice as long as the elements that are multiplied.
//
//go:linkname VmlalHighNU32 VmlalHighNU32
//go:noescape
func VmlalHighNU32(r *arm.Uint64X2, v0 *arm.Uint64X2, v1 *arm.Uint32X4, v2 *arm.Uint32)

// Unsigned Multiply-Add Long (vector). This instruction multiplies the vector elements in the lower or upper half of the first source SIMD&FP register by the corresponding vector elements of the second source SIMD&FP register, and accumulates the results with the vector elements of the destination SIMD&FP register. The destination vector elements are twice as long as the elements that are multiplied.
//
//go:linkname VmlalHighNU16 VmlalHighNU16
//go:noescape
func VmlalHighNU16(r *arm.Uint32X4, v0 *arm.Uint32X4, v1 *arm.Uint16X8, v2 *arm.Uint16)

// Signed Multiply-Add Long (vector). This instruction multiplies corresponding signed integer values in the lower or upper half of the vectors of the two source SIMD&FP registers, and accumulates the results with the vector elements of the destination SIMD&FP register. The destination vector elements are twice as long as the elements that are multiplied.
//
//go:linkname VmlalHighNS32 VmlalHighNS32
//go:noescape
func VmlalHighNS32(r *arm.Int64X2, v0 *arm.Int64X2, v1 *arm.Int32X4, v2 *arm.Int32)

// Signed Multiply-Add Long (vector). This instruction multiplies corresponding signed integer values in the lower or upper half of the vectors of the two source SIMD&FP registers, and accumulates the results with the vector elements of the destination SIMD&FP register. The destination vector elements are twice as long as the elements that are multiplied.
//
//go:linkname VmlalHighNS16 VmlalHighNS16
//go:noescape
func VmlalHighNS16(r *arm.Int32X4, v0 *arm.Int32X4, v1 *arm.Int16X8, v2 *arm.Int16)

// Unsigned Multiply-Subtract Long (vector). This instruction multiplies corresponding vector elements in the lower or upper half of the two source SIMD&FP registers, and subtracts the results from the vector elements of the destination SIMD&FP register. The destination vector elements are twice as long as the elements that are multiplied. All the values in this instruction are unsigned integer values.
//
//go:linkname VmlslHighU8 VmlslHighU8
//go:noescape
func VmlslHighU8(r *arm.Uint16X8, v0 *arm.Uint16X8, v1 *arm.Uint8X16, v2 *arm.Uint8X16)

// Unsigned Multiply-Subtract Long (vector). This instruction multiplies corresponding vector elements in the lower or upper half of the two source SIMD&FP registers, and subtracts the results from the vector elements of the destination SIMD&FP register. The destination vector elements are twice as long as the elements that are multiplied. All the values in this instruction are unsigned integer values.
//
//go:linkname VmlslHighU32 VmlslHighU32
//go:noescape
func VmlslHighU32(r *arm.Uint64X2, v0 *arm.Uint64X2, v1 *arm.Uint32X4, v2 *arm.Uint32X4)

// Unsigned Multiply-Subtract Long (vector). This instruction multiplies corresponding vector elements in the lower or upper half of the two source SIMD&FP registers, and subtracts the results from the vector elements of the destination SIMD&FP register. The destination vector elements are twice as long as the elements that are multiplied. All the values in this instruction are unsigned integer values.
//
//go:linkname VmlslHighU16 VmlslHighU16
//go:noescape
func VmlslHighU16(r *arm.Uint32X4, v0 *arm.Uint32X4, v1 *arm.Uint16X8, v2 *arm.Uint16X8)

// Signed Multiply-Subtract Long (vector). This instruction multiplies corresponding signed integer values in the lower or upper half of the vectors of the two source SIMD&FP registers, and subtracts the results from the vector elements of the destination SIMD&FP register. The destination vector elements are twice as long as the elements that are multiplied.
//
//go:linkname VmlslHighS8 VmlslHighS8
//go:noescape
func VmlslHighS8(r *arm.Int16X8, v0 *arm.Int16X8, v1 *arm.Int8X16, v2 *arm.Int8X16)

// Signed Multiply-Subtract Long (vector). This instruction multiplies corresponding signed integer values in the lower or upper half of the vectors of the two source SIMD&FP registers, and subtracts the results from the vector elements of the destination SIMD&FP register. The destination vector elements are twice as long as the elements that are multiplied.
//
//go:linkname VmlslHighS32 VmlslHighS32
//go:noescape
func VmlslHighS32(r *arm.Int64X2, v0 *arm.Int64X2, v1 *arm.Int32X4, v2 *arm.Int32X4)

// Signed Multiply-Subtract Long (vector). This instruction multiplies corresponding signed integer values in the lower or upper half of the vectors of the two source SIMD&FP registers, and subtracts the results from the vector elements of the destination SIMD&FP register. The destination vector elements are twice as long as the elements that are multiplied.
//
//go:linkname VmlslHighS16 VmlslHighS16
//go:noescape
func VmlslHighS16(r *arm.Int32X4, v0 *arm.Int32X4, v1 *arm.Int16X8, v2 *arm.Int16X8)

// Unsigned Multiply-Subtract Long (vector). This instruction multiplies corresponding vector elements in the lower or upper half of the two source SIMD&FP registers, and subtracts the results from the vector elements of the destination SIMD&FP register. The destination vector elements are twice as long as the elements that are multiplied. All the values in this instruction are unsigned integer values.
//
//go:linkname VmlslHighNU32 VmlslHighNU32
//go:noescape
func VmlslHighNU32(r *arm.Uint64X2, v0 *arm.Uint64X2, v1 *arm.Uint32X4, v2 *arm.Uint32)

// Unsigned Multiply-Subtract Long (vector). This instruction multiplies corresponding vector elements in the lower or upper half of the two source SIMD&FP registers, and subtracts the results from the vector elements of the destination SIMD&FP register. The destination vector elements are twice as long as the elements that are multiplied. All the values in this instruction are unsigned integer values.
//
//go:linkname VmlslHighNU16 VmlslHighNU16
//go:noescape
func VmlslHighNU16(r *arm.Uint32X4, v0 *arm.Uint32X4, v1 *arm.Uint16X8, v2 *arm.Uint16)

// Signed Multiply-Subtract Long (vector). This instruction multiplies corresponding signed integer values in the lower or upper half of the vectors of the two source SIMD&FP registers, and subtracts the results from the vector elements of the destination SIMD&FP register. The destination vector elements are twice as long as the elements that are multiplied.
//
//go:linkname VmlslHighNS32 VmlslHighNS32
//go:noescape
func VmlslHighNS32(r *arm.Int64X2, v0 *arm.Int64X2, v1 *arm.Int32X4, v2 *arm.Int32)

// Signed Multiply-Subtract Long (vector). This instruction multiplies corresponding signed integer values in the lower or upper half of the vectors of the two source SIMD&FP registers, and subtracts the results from the vector elements of the destination SIMD&FP register. The destination vector elements are twice as long as the elements that are multiplied.
//
//go:linkname VmlslHighNS16 VmlslHighNS16
//go:noescape
func VmlslHighNS16(r *arm.Int32X4, v0 *arm.Int32X4, v1 *arm.Int16X8, v2 *arm.Int16)

// Unsigned Absolute difference and Accumulate Long. This instruction subtracts the vector elements in the lower or upper half of the second source SIMD&FP register from the corresponding vector elements of the first source SIMD&FP register, and accumulates the absolute values of the results into the vector elements of the destination SIMD&FP register. The destination vector elements are twice as long as the source vector elements. All the values in this instruction are unsigned integer values.
//
//go:linkname VabalU8 VabalU8
//go:noescape
func VabalU8(r *arm.Uint16X8, v0 *arm.Uint16X8, v1 *arm.Uint8X8, v2 *arm.Uint8X8)

// Unsigned Absolute difference and Accumulate Long. This instruction subtracts the vector elements in the lower or upper half of the second source SIMD&FP register from the corresponding vector elements of the first source SIMD&FP register, and accumulates the absolute values of the results into the vector elements of the destination SIMD&FP register. The destination vector elements are twice as long as the source vector elements. All the values in this instruction are unsigned integer values.
//
//go:linkname VabalU32 VabalU32
//go:noescape
func VabalU32(r *arm.Uint64X2, v0 *arm.Uint64X2, v1 *arm.Uint32X2, v2 *arm.Uint32X2)

// Unsigned Absolute difference and Accumulate Long. This instruction subtracts the vector elements in the lower or upper half of the second source SIMD&FP register from the corresponding vector elements of the first source SIMD&FP register, and accumulates the absolute values of the results into the vector elements of the destination SIMD&FP register. The destination vector elements are twice as long as the source vector elements. All the values in this instruction are unsigned integer values.
//
//go:linkname VabalU16 VabalU16
//go:noescape
func VabalU16(r *arm.Uint32X4, v0 *arm.Uint32X4, v1 *arm.Uint16X4, v2 *arm.Uint16X4)

// Signed Absolute difference and Accumulate Long. This instruction subtracts the vector elements in the lower or upper half of the second source SIMD&FP register from the corresponding vector elements of the first source SIMD&FP register, and accumulates the absolute values of the results into the vector elements of the destination SIMD&FP register. The destination vector elements are twice as long as the source vector elements.
//
//go:linkname VabalS8 VabalS8
//go:noescape
func VabalS8(r *arm.Int16X8, v0 *arm.Int16X8, v1 *arm.Int8X8, v2 *arm.Int8X8)

// Signed Absolute difference and Accumulate Long. This instruction subtracts the vector elements in the lower or upper half of the second source SIMD&FP register from the corresponding vector elements of the first source SIMD&FP register, and accumulates the absolute values of the results into the vector elements of the destination SIMD&FP register. The destination vector elements are twice as long as the source vector elements.
//
//go:linkname VabalS32 VabalS32
//go:noescape
func VabalS32(r *arm.Int64X2, v0 *arm.Int64X2, v1 *arm.Int32X2, v2 *arm.Int32X2)

// Signed Absolute difference and Accumulate Long. This instruction subtracts the vector elements in the lower or upper half of the second source SIMD&FP register from the corresponding vector elements of the first source SIMD&FP register, and accumulates the absolute values of the results into the vector elements of the destination SIMD&FP register. The destination vector elements are twice as long as the source vector elements.
//
//go:linkname VabalS16 VabalS16
//go:noescape
func VabalS16(r *arm.Int32X4, v0 *arm.Int32X4, v1 *arm.Int16X4, v2 *arm.Int16X4)

// Unsigned Absolute difference and Accumulate Long. This instruction subtracts the vector elements in the lower or upper half of the second source SIMD&FP register from the corresponding vector elements of the first source SIMD&FP register, and accumulates the absolute values of the results into the vector elements of the destination SIMD&FP register. The destination vector elements are twice as long as the source vector elements. All the values in this instruction are unsigned integer values.
//
//go:linkname VabalHighU8 VabalHighU8
//go:noescape
func VabalHighU8(r *arm.Uint16X8, v0 *arm.Uint16X8, v1 *arm.Uint8X16, v2 *arm.Uint8X16)

// Unsigned Absolute difference and Accumulate Long. This instruction subtracts the vector elements in the lower or upper half of the second source SIMD&FP register from the corresponding vector elements of the first source SIMD&FP register, and accumulates the absolute values of the results into the vector elements of the destination SIMD&FP register. The destination vector elements are twice as long as the source vector elements. All the values in this instruction are unsigned integer values.
//
//go:linkname VabalHighU32 VabalHighU32
//go:noescape
func VabalHighU32(r *arm.Uint64X2, v0 *arm.Uint64X2, v1 *arm.Uint32X4, v2 *arm.Uint32X4)

// Unsigned Absolute difference and Accumulate Long. This instruction subtracts the vector elements in the lower or upper half of the second source SIMD&FP register from the corresponding vector elements of the first source SIMD&FP register, and accumulates the absolute values of the results into the vector elements of the destination SIMD&FP register. The destination vector elements are twice as long as the source vector elements. All the values in this instruction are unsigned integer values.
//
//go:linkname VabalHighU16 VabalHighU16
//go:noescape
func VabalHighU16(r *arm.Uint32X4, v0 *arm.Uint32X4, v1 *arm.Uint16X8, v2 *arm.Uint16X8)

// Signed Absolute difference and Accumulate Long. This instruction subtracts the vector elements in the lower or upper half of the second source SIMD&FP register from the corresponding vector elements of the first source SIMD&FP register, and accumulates the absolute values of the results into the vector elements of the destination SIMD&FP register. The destination vector elements are twice as long as the source vector elements.
//
//go:linkname VabalHighS8 VabalHighS8
//go:noescape
func VabalHighS8(r *arm.Int16X8, v0 *arm.Int16X8, v1 *arm.Int8X16, v2 *arm.Int8X16)

// Signed Absolute difference and Accumulate Long. This instruction subtracts the vector elements in the lower or upper half of the second source SIMD&FP register from the corresponding vector elements of the first source SIMD&FP register, and accumulates the absolute values of the results into the vector elements of the destination SIMD&FP register. The destination vector elements are twice as long as the source vector elements.
//
//go:linkname VabalHighS32 VabalHighS32
//go:noescape
func VabalHighS32(r *arm.Int64X2, v0 *arm.Int64X2, v1 *arm.Int32X4, v2 *arm.Int32X4)

// Signed Absolute difference and Accumulate Long. This instruction subtracts the vector elements in the lower or upper half of the second source SIMD&FP register from the corresponding vector elements of the first source SIMD&FP register, and accumulates the absolute values of the results into the vector elements of the destination SIMD&FP register. The destination vector elements are twice as long as the source vector elements.
//
//go:linkname VabalHighS16 VabalHighS16
//go:noescape
func VabalHighS16(r *arm.Int32X4, v0 *arm.Int32X4, v1 *arm.Int16X8, v2 *arm.Int16X8)
