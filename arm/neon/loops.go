package neon

import (
	"github.com/alivanz/go-simd/arm"
)

/*
#include <arm_neon.h>
*/
import "C"

// Signed Absolute Difference. This instruction subtracts the elements of the vector of the second source SIMD&FP register from the corresponding elements of the first source SIMD&FP register, places the the absolute values of the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VabdS8N VabdS8N
//go:noescape
func VabdS8N(r *arm.Int8, v0 *arm.Int8, v1 *arm.Int8, n int32)

// Signed Absolute Difference. This instruction subtracts the elements of the vector of the second source SIMD&FP register from the corresponding elements of the first source SIMD&FP register, places the the absolute values of the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VabdS16N VabdS16N
//go:noescape
func VabdS16N(r *arm.Int16, v0 *arm.Int16, v1 *arm.Int16, n int32)

// Signed Absolute Difference. This instruction subtracts the elements of the vector of the second source SIMD&FP register from the corresponding elements of the first source SIMD&FP register, places the the absolute values of the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VabdS32N VabdS32N
//go:noescape
func VabdS32N(r *arm.Int32, v0 *arm.Int32, v1 *arm.Int32, n int32)

// Unsigned Absolute Difference (vector). This instruction subtracts the elements of the vector of the second source SIMD&FP register from the corresponding elements of the first source SIMD&FP register, places the the absolute values of the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VabdU8N VabdU8N
//go:noescape
func VabdU8N(r *arm.Uint8, v0 *arm.Uint8, v1 *arm.Uint8, n int32)

// Unsigned Absolute Difference (vector). This instruction subtracts the elements of the vector of the second source SIMD&FP register from the corresponding elements of the first source SIMD&FP register, places the the absolute values of the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VabdU16N VabdU16N
//go:noescape
func VabdU16N(r *arm.Uint16, v0 *arm.Uint16, v1 *arm.Uint16, n int32)

// Unsigned Absolute Difference (vector). This instruction subtracts the elements of the vector of the second source SIMD&FP register from the corresponding elements of the first source SIMD&FP register, places the the absolute values of the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VabdU32N VabdU32N
//go:noescape
func VabdU32N(r *arm.Uint32, v0 *arm.Uint32, v1 *arm.Uint32, n int32)

// Floating-point Absolute Difference (vector). This instruction subtracts the floating-point values in the elements of the second source SIMD&FP register, from the corresponding floating-point values in the elements of the first source SIMD&FP register, places the absolute value of each result in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VabdF32N VabdF32N
//go:noescape
func VabdF32N(r *arm.Float32, v0 *arm.Float32, v1 *arm.Float32, n int32)

// Floating-point Absolute Difference (vector). This instruction subtracts the floating-point values in the elements of the second source SIMD&FP register, from the corresponding floating-point values in the elements of the first source SIMD&FP register, places the absolute value of each result in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VabdF64N VabdF64N
//go:noescape
func VabdF64N(r *arm.Float64, v0 *arm.Float64, v1 *arm.Float64, n int32)

// Floating-point Absolute Difference (vector). This instruction subtracts the floating-point values in the elements of the second source SIMD&FP register, from the corresponding floating-point values in the elements of the first source SIMD&FP register, places the absolute value of each result in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VabddF64N VabddF64N
//go:noescape
func VabddF64N(r *arm.Float64, v0 *arm.Float64, v1 *arm.Float64, n int32)

// Signed Absolute Difference. This instruction subtracts the elements of the vector of the second source SIMD&FP register from the corresponding elements of the first source SIMD&FP register, places the the absolute values of the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VabdqS8N VabdqS8N
//go:noescape
func VabdqS8N(r *arm.Int8, v0 *arm.Int8, v1 *arm.Int8, n int32)

// Signed Absolute Difference. This instruction subtracts the elements of the vector of the second source SIMD&FP register from the corresponding elements of the first source SIMD&FP register, places the the absolute values of the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VabdqS16N VabdqS16N
//go:noescape
func VabdqS16N(r *arm.Int16, v0 *arm.Int16, v1 *arm.Int16, n int32)

// Signed Absolute Difference. This instruction subtracts the elements of the vector of the second source SIMD&FP register from the corresponding elements of the first source SIMD&FP register, places the the absolute values of the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VabdqS32N VabdqS32N
//go:noescape
func VabdqS32N(r *arm.Int32, v0 *arm.Int32, v1 *arm.Int32, n int32)

// Unsigned Absolute Difference (vector). This instruction subtracts the elements of the vector of the second source SIMD&FP register from the corresponding elements of the first source SIMD&FP register, places the the absolute values of the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VabdqU8N VabdqU8N
//go:noescape
func VabdqU8N(r *arm.Uint8, v0 *arm.Uint8, v1 *arm.Uint8, n int32)

// Unsigned Absolute Difference (vector). This instruction subtracts the elements of the vector of the second source SIMD&FP register from the corresponding elements of the first source SIMD&FP register, places the the absolute values of the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VabdqU16N VabdqU16N
//go:noescape
func VabdqU16N(r *arm.Uint16, v0 *arm.Uint16, v1 *arm.Uint16, n int32)

// Unsigned Absolute Difference (vector). This instruction subtracts the elements of the vector of the second source SIMD&FP register from the corresponding elements of the first source SIMD&FP register, places the the absolute values of the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VabdqU32N VabdqU32N
//go:noescape
func VabdqU32N(r *arm.Uint32, v0 *arm.Uint32, v1 *arm.Uint32, n int32)

// Floating-point Absolute Difference (vector). This instruction subtracts the floating-point values in the elements of the second source SIMD&FP register, from the corresponding floating-point values in the elements of the first source SIMD&FP register, places the absolute value of each result in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VabdqF32N VabdqF32N
//go:noescape
func VabdqF32N(r *arm.Float32, v0 *arm.Float32, v1 *arm.Float32, n int32)

// Floating-point Absolute Difference (vector). This instruction subtracts the floating-point values in the elements of the second source SIMD&FP register, from the corresponding floating-point values in the elements of the first source SIMD&FP register, places the absolute value of each result in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VabdqF64N VabdqF64N
//go:noescape
func VabdqF64N(r *arm.Float64, v0 *arm.Float64, v1 *arm.Float64, n int32)

// Floating-point Absolute Difference (vector). This instruction subtracts the floating-point values in the elements of the second source SIMD&FP register, from the corresponding floating-point values in the elements of the first source SIMD&FP register, places the absolute value of each result in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VabdsF32N VabdsF32N
//go:noescape
func VabdsF32N(r *arm.Float32, v0 *arm.Float32, v1 *arm.Float32, n int32)

// Absolute value (vector). This instruction calculates the absolute value of each vector element in the source SIMD&FP register, puts the result into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VabsS8N VabsS8N
//go:noescape
func VabsS8N(r *arm.Int8, v0 *arm.Int8, n int32)

// Absolute value (vector). This instruction calculates the absolute value of each vector element in the source SIMD&FP register, puts the result into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VabsS16N VabsS16N
//go:noescape
func VabsS16N(r *arm.Int16, v0 *arm.Int16, n int32)

// Absolute value (vector). This instruction calculates the absolute value of each vector element in the source SIMD&FP register, puts the result into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VabsS32N VabsS32N
//go:noescape
func VabsS32N(r *arm.Int32, v0 *arm.Int32, n int32)

// Absolute value (vector). This instruction calculates the absolute value of each vector element in the source SIMD&FP register, puts the result into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VabsS64N VabsS64N
//go:noescape
func VabsS64N(r *arm.Int64, v0 *arm.Int64, n int32)

// Floating-point Absolute value (vector). This instruction calculates the absolute value of each vector element in the source SIMD&FP register, writes the result to a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VabsF32N VabsF32N
//go:noescape
func VabsF32N(r *arm.Float32, v0 *arm.Float32, n int32)

// Floating-point Absolute value (vector). This instruction calculates the absolute value of each vector element in the source SIMD&FP register, writes the result to a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VabsF64N VabsF64N
//go:noescape
func VabsF64N(r *arm.Float64, v0 *arm.Float64, n int32)

// Absolute value (vector). This instruction calculates the absolute value of each vector element in the source SIMD&FP register, puts the result into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VabsdS64N VabsdS64N
//go:noescape
func VabsdS64N(r *arm.Int64, v0 *arm.Int64, n int32)

// Absolute value (vector). This instruction calculates the absolute value of each vector element in the source SIMD&FP register, puts the result into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VabsqS8N VabsqS8N
//go:noescape
func VabsqS8N(r *arm.Int8, v0 *arm.Int8, n int32)

// Absolute value (vector). This instruction calculates the absolute value of each vector element in the source SIMD&FP register, puts the result into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VabsqS16N VabsqS16N
//go:noescape
func VabsqS16N(r *arm.Int16, v0 *arm.Int16, n int32)

// Absolute value (vector). This instruction calculates the absolute value of each vector element in the source SIMD&FP register, puts the result into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VabsqS32N VabsqS32N
//go:noescape
func VabsqS32N(r *arm.Int32, v0 *arm.Int32, n int32)

// Absolute value (vector). This instruction calculates the absolute value of each vector element in the source SIMD&FP register, puts the result into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VabsqS64N VabsqS64N
//go:noescape
func VabsqS64N(r *arm.Int64, v0 *arm.Int64, n int32)

// Floating-point Absolute value (vector). This instruction calculates the absolute value of each vector element in the source SIMD&FP register, writes the result to a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VabsqF32N VabsqF32N
//go:noescape
func VabsqF32N(r *arm.Float32, v0 *arm.Float32, n int32)

// Floating-point Absolute value (vector). This instruction calculates the absolute value of each vector element in the source SIMD&FP register, writes the result to a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VabsqF64N VabsqF64N
//go:noescape
func VabsqF64N(r *arm.Float64, v0 *arm.Float64, n int32)

// Add (vector). This instruction adds corresponding elements in the two source SIMD&FP registers, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VaddS8N VaddS8N
//go:noescape
func VaddS8N(r *arm.Int8, v0 *arm.Int8, v1 *arm.Int8, n int32)

// Add (vector). This instruction adds corresponding elements in the two source SIMD&FP registers, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VaddS16N VaddS16N
//go:noescape
func VaddS16N(r *arm.Int16, v0 *arm.Int16, v1 *arm.Int16, n int32)

// Add (vector). This instruction adds corresponding elements in the two source SIMD&FP registers, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VaddS32N VaddS32N
//go:noescape
func VaddS32N(r *arm.Int32, v0 *arm.Int32, v1 *arm.Int32, n int32)

// Add (vector). This instruction adds corresponding elements in the two source SIMD&FP registers, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VaddS64N VaddS64N
//go:noescape
func VaddS64N(r *arm.Int64, v0 *arm.Int64, v1 *arm.Int64, n int32)

// Add (vector). This instruction adds corresponding elements in the two source SIMD&FP registers, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VaddU8N VaddU8N
//go:noescape
func VaddU8N(r *arm.Uint8, v0 *arm.Uint8, v1 *arm.Uint8, n int32)

// Add (vector). This instruction adds corresponding elements in the two source SIMD&FP registers, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VaddU16N VaddU16N
//go:noescape
func VaddU16N(r *arm.Uint16, v0 *arm.Uint16, v1 *arm.Uint16, n int32)

// Add (vector). This instruction adds corresponding elements in the two source SIMD&FP registers, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VaddU32N VaddU32N
//go:noescape
func VaddU32N(r *arm.Uint32, v0 *arm.Uint32, v1 *arm.Uint32, n int32)

// Add (vector). This instruction adds corresponding elements in the two source SIMD&FP registers, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VaddU64N VaddU64N
//go:noescape
func VaddU64N(r *arm.Uint64, v0 *arm.Uint64, v1 *arm.Uint64, n int32)

// Floating-point Add (vector). This instruction adds corresponding vector elements in the two source SIMD&FP registers, writes the result into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are floating-point values.
//
//go:linkname VaddF32N VaddF32N
//go:noescape
func VaddF32N(r *arm.Float32, v0 *arm.Float32, v1 *arm.Float32, n int32)

// Floating-point Add (vector). This instruction adds corresponding vector elements in the two source SIMD&FP registers, writes the result into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are floating-point values.
//
//go:linkname VaddF64N VaddF64N
//go:noescape
func VaddF64N(r *arm.Float64, v0 *arm.Float64, v1 *arm.Float64, n int32)

// Add (vector). This instruction adds corresponding elements in the two source SIMD&FP registers, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VadddS64N VadddS64N
//go:noescape
func VadddS64N(r *arm.Int64, v0 *arm.Int64, v1 *arm.Int64, n int32)

// Add (vector). This instruction adds corresponding elements in the two source SIMD&FP registers, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VadddU64N VadddU64N
//go:noescape
func VadddU64N(r *arm.Uint64, v0 *arm.Uint64, v1 *arm.Uint64, n int32)

// Add (vector). This instruction adds corresponding elements in the two source SIMD&FP registers, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VaddqS8N VaddqS8N
//go:noescape
func VaddqS8N(r *arm.Int8, v0 *arm.Int8, v1 *arm.Int8, n int32)

// Add (vector). This instruction adds corresponding elements in the two source SIMD&FP registers, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VaddqS16N VaddqS16N
//go:noescape
func VaddqS16N(r *arm.Int16, v0 *arm.Int16, v1 *arm.Int16, n int32)

// Add (vector). This instruction adds corresponding elements in the two source SIMD&FP registers, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VaddqS32N VaddqS32N
//go:noescape
func VaddqS32N(r *arm.Int32, v0 *arm.Int32, v1 *arm.Int32, n int32)

// Add (vector). This instruction adds corresponding elements in the two source SIMD&FP registers, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VaddqS64N VaddqS64N
//go:noescape
func VaddqS64N(r *arm.Int64, v0 *arm.Int64, v1 *arm.Int64, n int32)

// Add (vector). This instruction adds corresponding elements in the two source SIMD&FP registers, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VaddqU8N VaddqU8N
//go:noescape
func VaddqU8N(r *arm.Uint8, v0 *arm.Uint8, v1 *arm.Uint8, n int32)

// Add (vector). This instruction adds corresponding elements in the two source SIMD&FP registers, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VaddqU16N VaddqU16N
//go:noescape
func VaddqU16N(r *arm.Uint16, v0 *arm.Uint16, v1 *arm.Uint16, n int32)

// Add (vector). This instruction adds corresponding elements in the two source SIMD&FP registers, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VaddqU32N VaddqU32N
//go:noescape
func VaddqU32N(r *arm.Uint32, v0 *arm.Uint32, v1 *arm.Uint32, n int32)

// Add (vector). This instruction adds corresponding elements in the two source SIMD&FP registers, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VaddqU64N VaddqU64N
//go:noescape
func VaddqU64N(r *arm.Uint64, v0 *arm.Uint64, v1 *arm.Uint64, n int32)

// Floating-point Add (vector). This instruction adds corresponding vector elements in the two source SIMD&FP registers, writes the result into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are floating-point values.
//
//go:linkname VaddqF32N VaddqF32N
//go:noescape
func VaddqF32N(r *arm.Float32, v0 *arm.Float32, v1 *arm.Float32, n int32)

// Floating-point Add (vector). This instruction adds corresponding vector elements in the two source SIMD&FP registers, writes the result into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are floating-point values.
//
//go:linkname VaddqF64N VaddqF64N
//go:noescape
func VaddqF64N(r *arm.Float64, v0 *arm.Float64, v1 *arm.Float64, n int32)

// Add across Vector. This instruction adds every vector element in the source SIMD&FP register together, and writes the scalar result to the destination SIMD&FP register.
//
//go:linkname VaddvS8N VaddvS8N
//go:noescape
func VaddvS8N(r *arm.Int8, v0 *arm.Int8, n int32)

// Add across Vector. This instruction adds every vector element in the source SIMD&FP register together, and writes the scalar result to the destination SIMD&FP register.
//
//go:linkname VaddvS16N VaddvS16N
//go:noescape
func VaddvS16N(r *arm.Int16, v0 *arm.Int16, n int32)

// Add across vector
//
//go:linkname VaddvS32N VaddvS32N
//go:noescape
func VaddvS32N(r *arm.Int32, v0 *arm.Int32, n int32)

// Add across Vector. This instruction adds every vector element in the source SIMD&FP register together, and writes the scalar result to the destination SIMD&FP register.
//
//go:linkname VaddvU8N VaddvU8N
//go:noescape
func VaddvU8N(r *arm.Uint8, v0 *arm.Uint8, n int32)

// Add across Vector. This instruction adds every vector element in the source SIMD&FP register together, and writes the scalar result to the destination SIMD&FP register.
//
//go:linkname VaddvU16N VaddvU16N
//go:noescape
func VaddvU16N(r *arm.Uint16, v0 *arm.Uint16, n int32)

// Add across vector
//
//go:linkname VaddvU32N VaddvU32N
//go:noescape
func VaddvU32N(r *arm.Uint32, v0 *arm.Uint32, n int32)

// Floating-point add across vector
//
//go:linkname VaddvF32N VaddvF32N
//go:noescape
func VaddvF32N(r *arm.Float32, v0 *arm.Float32, n int32)

// Add across Vector. This instruction adds every vector element in the source SIMD&FP register together, and writes the scalar result to the destination SIMD&FP register.
//
//go:linkname VaddvqS8N VaddvqS8N
//go:noescape
func VaddvqS8N(r *arm.Int8, v0 *arm.Int8, n int32)

// Add across Vector. This instruction adds every vector element in the source SIMD&FP register together, and writes the scalar result to the destination SIMD&FP register.
//
//go:linkname VaddvqS16N VaddvqS16N
//go:noescape
func VaddvqS16N(r *arm.Int16, v0 *arm.Int16, n int32)

// Add across Vector. This instruction adds every vector element in the source SIMD&FP register together, and writes the scalar result to the destination SIMD&FP register.
//
//go:linkname VaddvqS32N VaddvqS32N
//go:noescape
func VaddvqS32N(r *arm.Int32, v0 *arm.Int32, n int32)

// Add across vector
//
//go:linkname VaddvqS64N VaddvqS64N
//go:noescape
func VaddvqS64N(r *arm.Int64, v0 *arm.Int64, n int32)

// Add across Vector. This instruction adds every vector element in the source SIMD&FP register together, and writes the scalar result to the destination SIMD&FP register.
//
//go:linkname VaddvqU8N VaddvqU8N
//go:noescape
func VaddvqU8N(r *arm.Uint8, v0 *arm.Uint8, n int32)

// Add across Vector. This instruction adds every vector element in the source SIMD&FP register together, and writes the scalar result to the destination SIMD&FP register.
//
//go:linkname VaddvqU16N VaddvqU16N
//go:noescape
func VaddvqU16N(r *arm.Uint16, v0 *arm.Uint16, n int32)

// Add across Vector. This instruction adds every vector element in the source SIMD&FP register together, and writes the scalar result to the destination SIMD&FP register.
//
//go:linkname VaddvqU32N VaddvqU32N
//go:noescape
func VaddvqU32N(r *arm.Uint32, v0 *arm.Uint32, n int32)

// Add across vector
//
//go:linkname VaddvqU64N VaddvqU64N
//go:noescape
func VaddvqU64N(r *arm.Uint64, v0 *arm.Uint64, n int32)

// Floating-point add across vector
//
//go:linkname VaddvqF32N VaddvqF32N
//go:noescape
func VaddvqF32N(r *arm.Float32, v0 *arm.Float32, n int32)

// Floating-point add across vector
//
//go:linkname VaddvqF64N VaddvqF64N
//go:noescape
func VaddvqF64N(r *arm.Float64, v0 *arm.Float64, n int32)

// AES single round decryption.
//
//go:linkname VaesdqU8N VaesdqU8N
//go:noescape
func VaesdqU8N(r *arm.Uint8, v0 *arm.Uint8, v1 *arm.Uint8, n int32)

// AES single round encryption.
//
//go:linkname VaeseqU8N VaeseqU8N
//go:noescape
func VaeseqU8N(r *arm.Uint8, v0 *arm.Uint8, v1 *arm.Uint8, n int32)

// AES inverse mix columns.
//
//go:linkname VaesimcqU8N VaesimcqU8N
//go:noescape
func VaesimcqU8N(r *arm.Uint8, v0 *arm.Uint8, n int32)

// AES mix columns.
//
//go:linkname VaesmcqU8N VaesmcqU8N
//go:noescape
func VaesmcqU8N(r *arm.Uint8, v0 *arm.Uint8, n int32)

// Bitwise AND (vector). This instruction performs a bitwise AND between the two source SIMD&FP registers, and writes the result to the destination SIMD&FP register.
//
//go:linkname VandS8N VandS8N
//go:noescape
func VandS8N(r *arm.Int8, v0 *arm.Int8, v1 *arm.Int8, n int32)

// Bitwise AND (vector). This instruction performs a bitwise AND between the two source SIMD&FP registers, and writes the result to the destination SIMD&FP register.
//
//go:linkname VandS16N VandS16N
//go:noescape
func VandS16N(r *arm.Int16, v0 *arm.Int16, v1 *arm.Int16, n int32)

// Bitwise AND (vector). This instruction performs a bitwise AND between the two source SIMD&FP registers, and writes the result to the destination SIMD&FP register.
//
//go:linkname VandS32N VandS32N
//go:noescape
func VandS32N(r *arm.Int32, v0 *arm.Int32, v1 *arm.Int32, n int32)

// Bitwise AND (vector). This instruction performs a bitwise AND between the two source SIMD&FP registers, and writes the result to the destination SIMD&FP register.
//
//go:linkname VandS64N VandS64N
//go:noescape
func VandS64N(r *arm.Int64, v0 *arm.Int64, v1 *arm.Int64, n int32)

// Bitwise AND (vector). This instruction performs a bitwise AND between the two source SIMD&FP registers, and writes the result to the destination SIMD&FP register.
//
//go:linkname VandU8N VandU8N
//go:noescape
func VandU8N(r *arm.Uint8, v0 *arm.Uint8, v1 *arm.Uint8, n int32)

// Bitwise AND (vector). This instruction performs a bitwise AND between the two source SIMD&FP registers, and writes the result to the destination SIMD&FP register.
//
//go:linkname VandU16N VandU16N
//go:noescape
func VandU16N(r *arm.Uint16, v0 *arm.Uint16, v1 *arm.Uint16, n int32)

// Bitwise AND (vector). This instruction performs a bitwise AND between the two source SIMD&FP registers, and writes the result to the destination SIMD&FP register.
//
//go:linkname VandU32N VandU32N
//go:noescape
func VandU32N(r *arm.Uint32, v0 *arm.Uint32, v1 *arm.Uint32, n int32)

// Bitwise AND (vector). This instruction performs a bitwise AND between the two source SIMD&FP registers, and writes the result to the destination SIMD&FP register.
//
//go:linkname VandU64N VandU64N
//go:noescape
func VandU64N(r *arm.Uint64, v0 *arm.Uint64, v1 *arm.Uint64, n int32)

// Bitwise AND (vector). This instruction performs a bitwise AND between the two source SIMD&FP registers, and writes the result to the destination SIMD&FP register.
//
//go:linkname VandqS8N VandqS8N
//go:noescape
func VandqS8N(r *arm.Int8, v0 *arm.Int8, v1 *arm.Int8, n int32)

// Bitwise AND (vector). This instruction performs a bitwise AND between the two source SIMD&FP registers, and writes the result to the destination SIMD&FP register.
//
//go:linkname VandqS16N VandqS16N
//go:noescape
func VandqS16N(r *arm.Int16, v0 *arm.Int16, v1 *arm.Int16, n int32)

// Bitwise AND (vector). This instruction performs a bitwise AND between the two source SIMD&FP registers, and writes the result to the destination SIMD&FP register.
//
//go:linkname VandqS32N VandqS32N
//go:noescape
func VandqS32N(r *arm.Int32, v0 *arm.Int32, v1 *arm.Int32, n int32)

// Bitwise AND (vector). This instruction performs a bitwise AND between the two source SIMD&FP registers, and writes the result to the destination SIMD&FP register.
//
//go:linkname VandqS64N VandqS64N
//go:noescape
func VandqS64N(r *arm.Int64, v0 *arm.Int64, v1 *arm.Int64, n int32)

// Bitwise AND (vector). This instruction performs a bitwise AND between the two source SIMD&FP registers, and writes the result to the destination SIMD&FP register.
//
//go:linkname VandqU8N VandqU8N
//go:noescape
func VandqU8N(r *arm.Uint8, v0 *arm.Uint8, v1 *arm.Uint8, n int32)

// Bitwise AND (vector). This instruction performs a bitwise AND between the two source SIMD&FP registers, and writes the result to the destination SIMD&FP register.
//
//go:linkname VandqU16N VandqU16N
//go:noescape
func VandqU16N(r *arm.Uint16, v0 *arm.Uint16, v1 *arm.Uint16, n int32)

// Bitwise AND (vector). This instruction performs a bitwise AND between the two source SIMD&FP registers, and writes the result to the destination SIMD&FP register.
//
//go:linkname VandqU32N VandqU32N
//go:noescape
func VandqU32N(r *arm.Uint32, v0 *arm.Uint32, v1 *arm.Uint32, n int32)

// Bitwise AND (vector). This instruction performs a bitwise AND between the two source SIMD&FP registers, and writes the result to the destination SIMD&FP register.
//
//go:linkname VandqU64N VandqU64N
//go:noescape
func VandqU64N(r *arm.Uint64, v0 *arm.Uint64, v1 *arm.Uint64, n int32)

// Bitwise bit Clear (vector, register). This instruction performs a bitwise AND between the first source SIMD&FP register and the complement of the second source SIMD&FP register, and writes the result to the destination SIMD&FP register.
//
//go:linkname VbicS8N VbicS8N
//go:noescape
func VbicS8N(r *arm.Int8, v0 *arm.Int8, v1 *arm.Int8, n int32)

// Bitwise bit Clear (vector, register). This instruction performs a bitwise AND between the first source SIMD&FP register and the complement of the second source SIMD&FP register, and writes the result to the destination SIMD&FP register.
//
//go:linkname VbicS16N VbicS16N
//go:noescape
func VbicS16N(r *arm.Int16, v0 *arm.Int16, v1 *arm.Int16, n int32)

// Bitwise bit Clear (vector, register). This instruction performs a bitwise AND between the first source SIMD&FP register and the complement of the second source SIMD&FP register, and writes the result to the destination SIMD&FP register.
//
//go:linkname VbicS32N VbicS32N
//go:noescape
func VbicS32N(r *arm.Int32, v0 *arm.Int32, v1 *arm.Int32, n int32)

// Bitwise bit Clear (vector, register). This instruction performs a bitwise AND between the first source SIMD&FP register and the complement of the second source SIMD&FP register, and writes the result to the destination SIMD&FP register.
//
//go:linkname VbicS64N VbicS64N
//go:noescape
func VbicS64N(r *arm.Int64, v0 *arm.Int64, v1 *arm.Int64, n int32)

// Bitwise bit Clear (vector, register). This instruction performs a bitwise AND between the first source SIMD&FP register and the complement of the second source SIMD&FP register, and writes the result to the destination SIMD&FP register.
//
//go:linkname VbicU8N VbicU8N
//go:noescape
func VbicU8N(r *arm.Uint8, v0 *arm.Uint8, v1 *arm.Uint8, n int32)

// Bitwise bit Clear (vector, register). This instruction performs a bitwise AND between the first source SIMD&FP register and the complement of the second source SIMD&FP register, and writes the result to the destination SIMD&FP register.
//
//go:linkname VbicU16N VbicU16N
//go:noescape
func VbicU16N(r *arm.Uint16, v0 *arm.Uint16, v1 *arm.Uint16, n int32)

// Bitwise bit Clear (vector, register). This instruction performs a bitwise AND between the first source SIMD&FP register and the complement of the second source SIMD&FP register, and writes the result to the destination SIMD&FP register.
//
//go:linkname VbicU32N VbicU32N
//go:noescape
func VbicU32N(r *arm.Uint32, v0 *arm.Uint32, v1 *arm.Uint32, n int32)

// Bitwise bit Clear (vector, register). This instruction performs a bitwise AND between the first source SIMD&FP register and the complement of the second source SIMD&FP register, and writes the result to the destination SIMD&FP register.
//
//go:linkname VbicU64N VbicU64N
//go:noescape
func VbicU64N(r *arm.Uint64, v0 *arm.Uint64, v1 *arm.Uint64, n int32)

// Bitwise bit Clear (vector, register). This instruction performs a bitwise AND between the first source SIMD&FP register and the complement of the second source SIMD&FP register, and writes the result to the destination SIMD&FP register.
//
//go:linkname VbicqS8N VbicqS8N
//go:noescape
func VbicqS8N(r *arm.Int8, v0 *arm.Int8, v1 *arm.Int8, n int32)

// Bitwise bit Clear (vector, register). This instruction performs a bitwise AND between the first source SIMD&FP register and the complement of the second source SIMD&FP register, and writes the result to the destination SIMD&FP register.
//
//go:linkname VbicqS16N VbicqS16N
//go:noescape
func VbicqS16N(r *arm.Int16, v0 *arm.Int16, v1 *arm.Int16, n int32)

// Bitwise bit Clear (vector, register). This instruction performs a bitwise AND between the first source SIMD&FP register and the complement of the second source SIMD&FP register, and writes the result to the destination SIMD&FP register.
//
//go:linkname VbicqS32N VbicqS32N
//go:noescape
func VbicqS32N(r *arm.Int32, v0 *arm.Int32, v1 *arm.Int32, n int32)

// Bitwise bit Clear (vector, register). This instruction performs a bitwise AND between the first source SIMD&FP register and the complement of the second source SIMD&FP register, and writes the result to the destination SIMD&FP register.
//
//go:linkname VbicqS64N VbicqS64N
//go:noescape
func VbicqS64N(r *arm.Int64, v0 *arm.Int64, v1 *arm.Int64, n int32)

// Bitwise bit Clear (vector, register). This instruction performs a bitwise AND between the first source SIMD&FP register and the complement of the second source SIMD&FP register, and writes the result to the destination SIMD&FP register.
//
//go:linkname VbicqU8N VbicqU8N
//go:noescape
func VbicqU8N(r *arm.Uint8, v0 *arm.Uint8, v1 *arm.Uint8, n int32)

// Bitwise bit Clear (vector, register). This instruction performs a bitwise AND between the first source SIMD&FP register and the complement of the second source SIMD&FP register, and writes the result to the destination SIMD&FP register.
//
//go:linkname VbicqU16N VbicqU16N
//go:noescape
func VbicqU16N(r *arm.Uint16, v0 *arm.Uint16, v1 *arm.Uint16, n int32)

// Bitwise bit Clear (vector, register). This instruction performs a bitwise AND between the first source SIMD&FP register and the complement of the second source SIMD&FP register, and writes the result to the destination SIMD&FP register.
//
//go:linkname VbicqU32N VbicqU32N
//go:noescape
func VbicqU32N(r *arm.Uint32, v0 *arm.Uint32, v1 *arm.Uint32, n int32)

// Bitwise bit Clear (vector, register). This instruction performs a bitwise AND between the first source SIMD&FP register and the complement of the second source SIMD&FP register, and writes the result to the destination SIMD&FP register.
//
//go:linkname VbicqU64N VbicqU64N
//go:noescape
func VbicqU64N(r *arm.Uint64, v0 *arm.Uint64, v1 *arm.Uint64, n int32)

// Floating-point Complex Add.
//
//go:linkname VcaddRot270F32N VcaddRot270F32N
//go:noescape
func VcaddRot270F32N(r *arm.Float32, v0 *arm.Float32, v1 *arm.Float32, n int32)

// Floating-point Complex Add.
//
//go:linkname VcaddRot90F32N VcaddRot90F32N
//go:noescape
func VcaddRot90F32N(r *arm.Float32, v0 *arm.Float32, v1 *arm.Float32, n int32)

// Floating-point Complex Add.
//
//go:linkname VcaddqRot270F32N VcaddqRot270F32N
//go:noescape
func VcaddqRot270F32N(r *arm.Float32, v0 *arm.Float32, v1 *arm.Float32, n int32)

// Floating-point Complex Add.
//
//go:linkname VcaddqRot270F64N VcaddqRot270F64N
//go:noescape
func VcaddqRot270F64N(r *arm.Float64, v0 *arm.Float64, v1 *arm.Float64, n int32)

// Floating-point Complex Add.
//
//go:linkname VcaddqRot90F32N VcaddqRot90F32N
//go:noescape
func VcaddqRot90F32N(r *arm.Float32, v0 *arm.Float32, v1 *arm.Float32, n int32)

// Floating-point Complex Add.
//
//go:linkname VcaddqRot90F64N VcaddqRot90F64N
//go:noescape
func VcaddqRot90F64N(r *arm.Float64, v0 *arm.Float64, v1 *arm.Float64, n int32)

// Floating-point Absolute Compare Greater than or Equal (vector). This instruction compares the absolute value of each floating-point value in the first source SIMD&FP register with the absolute value of the corresponding floating-point value in the second source SIMD&FP register and if the first value is greater than or equal to the second value sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VcageF32N VcageF32N
//go:noescape
func VcageF32N(r *arm.Uint32, v0 *arm.Float32, v1 *arm.Float32, n int32)

// Floating-point Absolute Compare Greater than or Equal (vector). This instruction compares the absolute value of each floating-point value in the first source SIMD&FP register with the absolute value of the corresponding floating-point value in the second source SIMD&FP register and if the first value is greater than or equal to the second value sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VcageF64N VcageF64N
//go:noescape
func VcageF64N(r *arm.Uint64, v0 *arm.Float64, v1 *arm.Float64, n int32)

// Floating-point Absolute Compare Greater than or Equal (vector). This instruction compares the absolute value of each floating-point value in the first source SIMD&FP register with the absolute value of the corresponding floating-point value in the second source SIMD&FP register and if the first value is greater than or equal to the second value sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VcagedF64N VcagedF64N
//go:noescape
func VcagedF64N(r *arm.Uint64, v0 *arm.Float64, v1 *arm.Float64, n int32)

// Floating-point Absolute Compare Greater than or Equal (vector). This instruction compares the absolute value of each floating-point value in the first source SIMD&FP register with the absolute value of the corresponding floating-point value in the second source SIMD&FP register and if the first value is greater than or equal to the second value sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VcageqF32N VcageqF32N
//go:noescape
func VcageqF32N(r *arm.Uint32, v0 *arm.Float32, v1 *arm.Float32, n int32)

// Floating-point Absolute Compare Greater than or Equal (vector). This instruction compares the absolute value of each floating-point value in the first source SIMD&FP register with the absolute value of the corresponding floating-point value in the second source SIMD&FP register and if the first value is greater than or equal to the second value sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VcageqF64N VcageqF64N
//go:noescape
func VcageqF64N(r *arm.Uint64, v0 *arm.Float64, v1 *arm.Float64, n int32)

// Floating-point Absolute Compare Greater than or Equal (vector). This instruction compares the absolute value of each floating-point value in the first source SIMD&FP register with the absolute value of the corresponding floating-point value in the second source SIMD&FP register and if the first value is greater than or equal to the second value sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VcagesF32N VcagesF32N
//go:noescape
func VcagesF32N(r *arm.Uint32, v0 *arm.Float32, v1 *arm.Float32, n int32)

// Floating-point Absolute Compare Greater than (vector). This instruction compares the absolute value of each vector element in the first source SIMD&FP register with the absolute value of the corresponding vector element in the second source SIMD&FP register and if the first value is greater than the second value sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VcagtF32N VcagtF32N
//go:noescape
func VcagtF32N(r *arm.Uint32, v0 *arm.Float32, v1 *arm.Float32, n int32)

// Floating-point Absolute Compare Greater than (vector). This instruction compares the absolute value of each vector element in the first source SIMD&FP register with the absolute value of the corresponding vector element in the second source SIMD&FP register and if the first value is greater than the second value sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VcagtF64N VcagtF64N
//go:noescape
func VcagtF64N(r *arm.Uint64, v0 *arm.Float64, v1 *arm.Float64, n int32)

// Floating-point Absolute Compare Greater than (vector). This instruction compares the absolute value of each vector element in the first source SIMD&FP register with the absolute value of the corresponding vector element in the second source SIMD&FP register and if the first value is greater than the second value sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VcagtdF64N VcagtdF64N
//go:noescape
func VcagtdF64N(r *arm.Uint64, v0 *arm.Float64, v1 *arm.Float64, n int32)

// Floating-point Absolute Compare Greater than (vector). This instruction compares the absolute value of each vector element in the first source SIMD&FP register with the absolute value of the corresponding vector element in the second source SIMD&FP register and if the first value is greater than the second value sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VcagtqF32N VcagtqF32N
//go:noescape
func VcagtqF32N(r *arm.Uint32, v0 *arm.Float32, v1 *arm.Float32, n int32)

// Floating-point Absolute Compare Greater than (vector). This instruction compares the absolute value of each vector element in the first source SIMD&FP register with the absolute value of the corresponding vector element in the second source SIMD&FP register and if the first value is greater than the second value sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VcagtqF64N VcagtqF64N
//go:noescape
func VcagtqF64N(r *arm.Uint64, v0 *arm.Float64, v1 *arm.Float64, n int32)

// Floating-point Absolute Compare Greater than (vector). This instruction compares the absolute value of each vector element in the first source SIMD&FP register with the absolute value of the corresponding vector element in the second source SIMD&FP register and if the first value is greater than the second value sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VcagtsF32N VcagtsF32N
//go:noescape
func VcagtsF32N(r *arm.Uint32, v0 *arm.Float32, v1 *arm.Float32, n int32)

// Floating-point absolute compare less than or equal
//
//go:linkname VcaleF32N VcaleF32N
//go:noescape
func VcaleF32N(r *arm.Uint32, v0 *arm.Float32, v1 *arm.Float32, n int32)

// Floating-point absolute compare less than or equal
//
//go:linkname VcaleF64N VcaleF64N
//go:noescape
func VcaleF64N(r *arm.Uint64, v0 *arm.Float64, v1 *arm.Float64, n int32)

// Floating-point absolute compare less than or equal
//
//go:linkname VcaledF64N VcaledF64N
//go:noescape
func VcaledF64N(r *arm.Uint64, v0 *arm.Float64, v1 *arm.Float64, n int32)

// Floating-point absolute compare less than or equal
//
//go:linkname VcaleqF32N VcaleqF32N
//go:noescape
func VcaleqF32N(r *arm.Uint32, v0 *arm.Float32, v1 *arm.Float32, n int32)

// Floating-point absolute compare less than or equal
//
//go:linkname VcaleqF64N VcaleqF64N
//go:noescape
func VcaleqF64N(r *arm.Uint64, v0 *arm.Float64, v1 *arm.Float64, n int32)

// Floating-point absolute compare less than or equal
//
//go:linkname VcalesF32N VcalesF32N
//go:noescape
func VcalesF32N(r *arm.Uint32, v0 *arm.Float32, v1 *arm.Float32, n int32)

// Floating-point absolute compare less than
//
//go:linkname VcaltF32N VcaltF32N
//go:noescape
func VcaltF32N(r *arm.Uint32, v0 *arm.Float32, v1 *arm.Float32, n int32)

// Floating-point absolute compare less than
//
//go:linkname VcaltF64N VcaltF64N
//go:noescape
func VcaltF64N(r *arm.Uint64, v0 *arm.Float64, v1 *arm.Float64, n int32)

// Floating-point absolute compare less than
//
//go:linkname VcaltdF64N VcaltdF64N
//go:noescape
func VcaltdF64N(r *arm.Uint64, v0 *arm.Float64, v1 *arm.Float64, n int32)

// Floating-point absolute compare less than
//
//go:linkname VcaltqF32N VcaltqF32N
//go:noescape
func VcaltqF32N(r *arm.Uint32, v0 *arm.Float32, v1 *arm.Float32, n int32)

// Floating-point absolute compare less than
//
//go:linkname VcaltqF64N VcaltqF64N
//go:noescape
func VcaltqF64N(r *arm.Uint64, v0 *arm.Float64, v1 *arm.Float64, n int32)

// Floating-point absolute compare less than
//
//go:linkname VcaltsF32N VcaltsF32N
//go:noescape
func VcaltsF32N(r *arm.Uint32, v0 *arm.Float32, v1 *arm.Float32, n int32)

// Compare bitwise Equal (vector). This instruction compares each vector element from the first source SIMD&FP register with the corresponding vector element from the second source SIMD&FP register, and if the comparison is equal sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VceqS8N VceqS8N
//go:noescape
func VceqS8N(r *arm.Uint8, v0 *arm.Int8, v1 *arm.Int8, n int32)

// Compare bitwise Equal (vector). This instruction compares each vector element from the first source SIMD&FP register with the corresponding vector element from the second source SIMD&FP register, and if the comparison is equal sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VceqS16N VceqS16N
//go:noescape
func VceqS16N(r *arm.Uint16, v0 *arm.Int16, v1 *arm.Int16, n int32)

// Compare bitwise Equal (vector). This instruction compares each vector element from the first source SIMD&FP register with the corresponding vector element from the second source SIMD&FP register, and if the comparison is equal sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VceqS32N VceqS32N
//go:noescape
func VceqS32N(r *arm.Uint32, v0 *arm.Int32, v1 *arm.Int32, n int32)

// Compare bitwise Equal (vector). This instruction compares each vector element from the first source SIMD&FP register with the corresponding vector element from the second source SIMD&FP register, and if the comparison is equal sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VceqS64N VceqS64N
//go:noescape
func VceqS64N(r *arm.Uint64, v0 *arm.Int64, v1 *arm.Int64, n int32)

// Compare bitwise Equal (vector). This instruction compares each vector element from the first source SIMD&FP register with the corresponding vector element from the second source SIMD&FP register, and if the comparison is equal sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VceqU8N VceqU8N
//go:noescape
func VceqU8N(r *arm.Uint8, v0 *arm.Uint8, v1 *arm.Uint8, n int32)

// Compare bitwise Equal (vector). This instruction compares each vector element from the first source SIMD&FP register with the corresponding vector element from the second source SIMD&FP register, and if the comparison is equal sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VceqU16N VceqU16N
//go:noescape
func VceqU16N(r *arm.Uint16, v0 *arm.Uint16, v1 *arm.Uint16, n int32)

// Compare bitwise Equal (vector). This instruction compares each vector element from the first source SIMD&FP register with the corresponding vector element from the second source SIMD&FP register, and if the comparison is equal sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VceqU32N VceqU32N
//go:noescape
func VceqU32N(r *arm.Uint32, v0 *arm.Uint32, v1 *arm.Uint32, n int32)

// Compare bitwise Equal (vector). This instruction compares each vector element from the first source SIMD&FP register with the corresponding vector element from the second source SIMD&FP register, and if the comparison is equal sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VceqU64N VceqU64N
//go:noescape
func VceqU64N(r *arm.Uint64, v0 *arm.Uint64, v1 *arm.Uint64, n int32)

// Floating-point Compare Equal (vector). This instruction compares each floating-point value from the first source SIMD&FP register, with the corresponding floating-point value from the second source SIMD&FP register, and if the comparison is equal sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VceqF32N VceqF32N
//go:noescape
func VceqF32N(r *arm.Uint32, v0 *arm.Float32, v1 *arm.Float32, n int32)

// Floating-point Compare Equal (vector). This instruction compares each floating-point value from the first source SIMD&FP register, with the corresponding floating-point value from the second source SIMD&FP register, and if the comparison is equal sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VceqF64N VceqF64N
//go:noescape
func VceqF64N(r *arm.Uint64, v0 *arm.Float64, v1 *arm.Float64, n int32)

// Compare bitwise Equal (vector). This instruction compares each vector element from the first source SIMD&FP register with the corresponding vector element from the second source SIMD&FP register, and if the comparison is equal sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VceqdS64N VceqdS64N
//go:noescape
func VceqdS64N(r *arm.Uint64, v0 *arm.Int64, v1 *arm.Int64, n int32)

// Compare bitwise Equal (vector). This instruction compares each vector element from the first source SIMD&FP register with the corresponding vector element from the second source SIMD&FP register, and if the comparison is equal sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VceqdU64N VceqdU64N
//go:noescape
func VceqdU64N(r *arm.Uint64, v0 *arm.Uint64, v1 *arm.Uint64, n int32)

// Floating-point Compare Equal (vector). This instruction compares each floating-point value from the first source SIMD&FP register, with the corresponding floating-point value from the second source SIMD&FP register, and if the comparison is equal sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VceqdF64N VceqdF64N
//go:noescape
func VceqdF64N(r *arm.Uint64, v0 *arm.Float64, v1 *arm.Float64, n int32)

// Compare bitwise Equal (vector). This instruction compares each vector element from the first source SIMD&FP register with the corresponding vector element from the second source SIMD&FP register, and if the comparison is equal sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VceqqS8N VceqqS8N
//go:noescape
func VceqqS8N(r *arm.Uint8, v0 *arm.Int8, v1 *arm.Int8, n int32)

// Compare bitwise Equal (vector). This instruction compares each vector element from the first source SIMD&FP register with the corresponding vector element from the second source SIMD&FP register, and if the comparison is equal sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VceqqS16N VceqqS16N
//go:noescape
func VceqqS16N(r *arm.Uint16, v0 *arm.Int16, v1 *arm.Int16, n int32)

// Compare bitwise Equal (vector). This instruction compares each vector element from the first source SIMD&FP register with the corresponding vector element from the second source SIMD&FP register, and if the comparison is equal sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VceqqS32N VceqqS32N
//go:noescape
func VceqqS32N(r *arm.Uint32, v0 *arm.Int32, v1 *arm.Int32, n int32)

// Compare bitwise Equal (vector). This instruction compares each vector element from the first source SIMD&FP register with the corresponding vector element from the second source SIMD&FP register, and if the comparison is equal sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VceqqS64N VceqqS64N
//go:noescape
func VceqqS64N(r *arm.Uint64, v0 *arm.Int64, v1 *arm.Int64, n int32)

// Compare bitwise Equal (vector). This instruction compares each vector element from the first source SIMD&FP register with the corresponding vector element from the second source SIMD&FP register, and if the comparison is equal sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VceqqU8N VceqqU8N
//go:noescape
func VceqqU8N(r *arm.Uint8, v0 *arm.Uint8, v1 *arm.Uint8, n int32)

// Compare bitwise Equal (vector). This instruction compares each vector element from the first source SIMD&FP register with the corresponding vector element from the second source SIMD&FP register, and if the comparison is equal sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VceqqU16N VceqqU16N
//go:noescape
func VceqqU16N(r *arm.Uint16, v0 *arm.Uint16, v1 *arm.Uint16, n int32)

// Compare bitwise Equal (vector). This instruction compares each vector element from the first source SIMD&FP register with the corresponding vector element from the second source SIMD&FP register, and if the comparison is equal sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VceqqU32N VceqqU32N
//go:noescape
func VceqqU32N(r *arm.Uint32, v0 *arm.Uint32, v1 *arm.Uint32, n int32)

// Compare bitwise Equal (vector). This instruction compares each vector element from the first source SIMD&FP register with the corresponding vector element from the second source SIMD&FP register, and if the comparison is equal sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VceqqU64N VceqqU64N
//go:noescape
func VceqqU64N(r *arm.Uint64, v0 *arm.Uint64, v1 *arm.Uint64, n int32)

// Floating-point Compare Equal (vector). This instruction compares each floating-point value from the first source SIMD&FP register, with the corresponding floating-point value from the second source SIMD&FP register, and if the comparison is equal sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VceqqF32N VceqqF32N
//go:noescape
func VceqqF32N(r *arm.Uint32, v0 *arm.Float32, v1 *arm.Float32, n int32)

// Floating-point Compare Equal (vector). This instruction compares each floating-point value from the first source SIMD&FP register, with the corresponding floating-point value from the second source SIMD&FP register, and if the comparison is equal sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VceqqF64N VceqqF64N
//go:noescape
func VceqqF64N(r *arm.Uint64, v0 *arm.Float64, v1 *arm.Float64, n int32)

// Floating-point Compare Equal (vector). This instruction compares each floating-point value from the first source SIMD&FP register, with the corresponding floating-point value from the second source SIMD&FP register, and if the comparison is equal sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VceqsF32N VceqsF32N
//go:noescape
func VceqsF32N(r *arm.Uint32, v0 *arm.Float32, v1 *arm.Float32, n int32)

// Compare bitwise Equal to zero (vector). This instruction reads each vector element in the source SIMD&FP register and if the value is equal to zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VceqzS8N VceqzS8N
//go:noescape
func VceqzS8N(r *arm.Uint8, v0 *arm.Int8, n int32)

// Compare bitwise Equal to zero (vector). This instruction reads each vector element in the source SIMD&FP register and if the value is equal to zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VceqzS16N VceqzS16N
//go:noescape
func VceqzS16N(r *arm.Uint16, v0 *arm.Int16, n int32)

// Compare bitwise Equal to zero (vector). This instruction reads each vector element in the source SIMD&FP register and if the value is equal to zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VceqzS32N VceqzS32N
//go:noescape
func VceqzS32N(r *arm.Uint32, v0 *arm.Int32, n int32)

// Compare bitwise Equal to zero (vector). This instruction reads each vector element in the source SIMD&FP register and if the value is equal to zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VceqzS64N VceqzS64N
//go:noescape
func VceqzS64N(r *arm.Uint64, v0 *arm.Int64, n int32)

// Compare bitwise Equal to zero (vector). This instruction reads each vector element in the source SIMD&FP register and if the value is equal to zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VceqzU8N VceqzU8N
//go:noescape
func VceqzU8N(r *arm.Uint8, v0 *arm.Uint8, n int32)

// Compare bitwise Equal to zero (vector). This instruction reads each vector element in the source SIMD&FP register and if the value is equal to zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VceqzU16N VceqzU16N
//go:noescape
func VceqzU16N(r *arm.Uint16, v0 *arm.Uint16, n int32)

// Compare bitwise Equal to zero (vector). This instruction reads each vector element in the source SIMD&FP register and if the value is equal to zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VceqzU32N VceqzU32N
//go:noescape
func VceqzU32N(r *arm.Uint32, v0 *arm.Uint32, n int32)

// Compare bitwise Equal to zero (vector). This instruction reads each vector element in the source SIMD&FP register and if the value is equal to zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VceqzU64N VceqzU64N
//go:noescape
func VceqzU64N(r *arm.Uint64, v0 *arm.Uint64, n int32)

// Floating-point Compare Equal to zero (vector). This instruction reads each floating-point value in the source SIMD&FP register and if the value is equal to zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VceqzF32N VceqzF32N
//go:noescape
func VceqzF32N(r *arm.Uint32, v0 *arm.Float32, n int32)

// Floating-point Compare Equal to zero (vector). This instruction reads each floating-point value in the source SIMD&FP register and if the value is equal to zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VceqzF64N VceqzF64N
//go:noescape
func VceqzF64N(r *arm.Uint64, v0 *arm.Float64, n int32)

// Compare bitwise Equal to zero (vector). This instruction reads each vector element in the source SIMD&FP register and if the value is equal to zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VceqzdS64N VceqzdS64N
//go:noescape
func VceqzdS64N(r *arm.Uint64, v0 *arm.Int64, n int32)

// Compare bitwise Equal to zero (vector). This instruction reads each vector element in the source SIMD&FP register and if the value is equal to zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VceqzdU64N VceqzdU64N
//go:noescape
func VceqzdU64N(r *arm.Uint64, v0 *arm.Uint64, n int32)

// Floating-point Compare Equal to zero (vector). This instruction reads each floating-point value in the source SIMD&FP register and if the value is equal to zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VceqzdF64N VceqzdF64N
//go:noescape
func VceqzdF64N(r *arm.Uint64, v0 *arm.Float64, n int32)

// Compare bitwise Equal to zero (vector). This instruction reads each vector element in the source SIMD&FP register and if the value is equal to zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VceqzqS8N VceqzqS8N
//go:noescape
func VceqzqS8N(r *arm.Uint8, v0 *arm.Int8, n int32)

// Compare bitwise Equal to zero (vector). This instruction reads each vector element in the source SIMD&FP register and if the value is equal to zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VceqzqS16N VceqzqS16N
//go:noescape
func VceqzqS16N(r *arm.Uint16, v0 *arm.Int16, n int32)

// Compare bitwise Equal to zero (vector). This instruction reads each vector element in the source SIMD&FP register and if the value is equal to zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VceqzqS32N VceqzqS32N
//go:noescape
func VceqzqS32N(r *arm.Uint32, v0 *arm.Int32, n int32)

// Compare bitwise Equal to zero (vector). This instruction reads each vector element in the source SIMD&FP register and if the value is equal to zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VceqzqS64N VceqzqS64N
//go:noescape
func VceqzqS64N(r *arm.Uint64, v0 *arm.Int64, n int32)

// Compare bitwise Equal to zero (vector). This instruction reads each vector element in the source SIMD&FP register and if the value is equal to zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VceqzqU8N VceqzqU8N
//go:noescape
func VceqzqU8N(r *arm.Uint8, v0 *arm.Uint8, n int32)

// Compare bitwise Equal to zero (vector). This instruction reads each vector element in the source SIMD&FP register and if the value is equal to zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VceqzqU16N VceqzqU16N
//go:noescape
func VceqzqU16N(r *arm.Uint16, v0 *arm.Uint16, n int32)

// Compare bitwise Equal to zero (vector). This instruction reads each vector element in the source SIMD&FP register and if the value is equal to zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VceqzqU32N VceqzqU32N
//go:noescape
func VceqzqU32N(r *arm.Uint32, v0 *arm.Uint32, n int32)

// Compare bitwise Equal to zero (vector). This instruction reads each vector element in the source SIMD&FP register and if the value is equal to zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VceqzqU64N VceqzqU64N
//go:noescape
func VceqzqU64N(r *arm.Uint64, v0 *arm.Uint64, n int32)

// Floating-point Compare Equal to zero (vector). This instruction reads each floating-point value in the source SIMD&FP register and if the value is equal to zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VceqzqF32N VceqzqF32N
//go:noescape
func VceqzqF32N(r *arm.Uint32, v0 *arm.Float32, n int32)

// Floating-point Compare Equal to zero (vector). This instruction reads each floating-point value in the source SIMD&FP register and if the value is equal to zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VceqzqF64N VceqzqF64N
//go:noescape
func VceqzqF64N(r *arm.Uint64, v0 *arm.Float64, n int32)

// Floating-point Compare Equal to zero (vector). This instruction reads each floating-point value in the source SIMD&FP register and if the value is equal to zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VceqzsF32N VceqzsF32N
//go:noescape
func VceqzsF32N(r *arm.Uint32, v0 *arm.Float32, n int32)

// Compare signed Greater than or Equal (vector). This instruction compares each vector element in the first source SIMD&FP register with the corresponding vector element in the second source SIMD&FP register and if the first signed integer value is greater than or equal to the second signed integer value sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VcgeS8N VcgeS8N
//go:noescape
func VcgeS8N(r *arm.Uint8, v0 *arm.Int8, v1 *arm.Int8, n int32)

// Compare signed Greater than or Equal (vector). This instruction compares each vector element in the first source SIMD&FP register with the corresponding vector element in the second source SIMD&FP register and if the first signed integer value is greater than or equal to the second signed integer value sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VcgeS16N VcgeS16N
//go:noescape
func VcgeS16N(r *arm.Uint16, v0 *arm.Int16, v1 *arm.Int16, n int32)

// Compare signed Greater than or Equal (vector). This instruction compares each vector element in the first source SIMD&FP register with the corresponding vector element in the second source SIMD&FP register and if the first signed integer value is greater than or equal to the second signed integer value sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VcgeS32N VcgeS32N
//go:noescape
func VcgeS32N(r *arm.Uint32, v0 *arm.Int32, v1 *arm.Int32, n int32)

// Compare signed Greater than or Equal (vector). This instruction compares each vector element in the first source SIMD&FP register with the corresponding vector element in the second source SIMD&FP register and if the first signed integer value is greater than or equal to the second signed integer value sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VcgeS64N VcgeS64N
//go:noescape
func VcgeS64N(r *arm.Uint64, v0 *arm.Int64, v1 *arm.Int64, n int32)

// Compare unsigned Higher or Same (vector). This instruction compares each vector element in the first source SIMD&FP register with the corresponding vector element in the second source SIMD&FP register and if the first unsigned integer value is greater than or equal to the second unsigned integer value sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VcgeU8N VcgeU8N
//go:noescape
func VcgeU8N(r *arm.Uint8, v0 *arm.Uint8, v1 *arm.Uint8, n int32)

// Compare unsigned Higher or Same (vector). This instruction compares each vector element in the first source SIMD&FP register with the corresponding vector element in the second source SIMD&FP register and if the first unsigned integer value is greater than or equal to the second unsigned integer value sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VcgeU16N VcgeU16N
//go:noescape
func VcgeU16N(r *arm.Uint16, v0 *arm.Uint16, v1 *arm.Uint16, n int32)

// Compare unsigned Higher or Same (vector). This instruction compares each vector element in the first source SIMD&FP register with the corresponding vector element in the second source SIMD&FP register and if the first unsigned integer value is greater than or equal to the second unsigned integer value sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VcgeU32N VcgeU32N
//go:noescape
func VcgeU32N(r *arm.Uint32, v0 *arm.Uint32, v1 *arm.Uint32, n int32)

// Compare unsigned Higher or Same (vector). This instruction compares each vector element in the first source SIMD&FP register with the corresponding vector element in the second source SIMD&FP register and if the first unsigned integer value is greater than or equal to the second unsigned integer value sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VcgeU64N VcgeU64N
//go:noescape
func VcgeU64N(r *arm.Uint64, v0 *arm.Uint64, v1 *arm.Uint64, n int32)

// Floating-point Compare Greater than or Equal (vector). This instruction reads each floating-point value in the first source SIMD&FP register and if the value is greater than or equal to the corresponding floating-point value in the second source SIMD&FP register sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VcgeF32N VcgeF32N
//go:noescape
func VcgeF32N(r *arm.Uint32, v0 *arm.Float32, v1 *arm.Float32, n int32)

// Floating-point Compare Greater than or Equal (vector). This instruction reads each floating-point value in the first source SIMD&FP register and if the value is greater than or equal to the corresponding floating-point value in the second source SIMD&FP register sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VcgeF64N VcgeF64N
//go:noescape
func VcgeF64N(r *arm.Uint64, v0 *arm.Float64, v1 *arm.Float64, n int32)

// Compare signed Greater than or Equal (vector). This instruction compares each vector element in the first source SIMD&FP register with the corresponding vector element in the second source SIMD&FP register and if the first signed integer value is greater than or equal to the second signed integer value sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VcgedS64N VcgedS64N
//go:noescape
func VcgedS64N(r *arm.Uint64, v0 *arm.Int64, v1 *arm.Int64, n int32)

// Compare unsigned Higher or Same (vector). This instruction compares each vector element in the first source SIMD&FP register with the corresponding vector element in the second source SIMD&FP register and if the first unsigned integer value is greater than or equal to the second unsigned integer value sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VcgedU64N VcgedU64N
//go:noescape
func VcgedU64N(r *arm.Uint64, v0 *arm.Uint64, v1 *arm.Uint64, n int32)

// Floating-point Compare Greater than or Equal (vector). This instruction reads each floating-point value in the first source SIMD&FP register and if the value is greater than or equal to the corresponding floating-point value in the second source SIMD&FP register sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VcgedF64N VcgedF64N
//go:noescape
func VcgedF64N(r *arm.Uint64, v0 *arm.Float64, v1 *arm.Float64, n int32)

// Compare signed Greater than or Equal (vector). This instruction compares each vector element in the first source SIMD&FP register with the corresponding vector element in the second source SIMD&FP register and if the first signed integer value is greater than or equal to the second signed integer value sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VcgeqS8N VcgeqS8N
//go:noescape
func VcgeqS8N(r *arm.Uint8, v0 *arm.Int8, v1 *arm.Int8, n int32)

// Compare signed Greater than or Equal (vector). This instruction compares each vector element in the first source SIMD&FP register with the corresponding vector element in the second source SIMD&FP register and if the first signed integer value is greater than or equal to the second signed integer value sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VcgeqS16N VcgeqS16N
//go:noescape
func VcgeqS16N(r *arm.Uint16, v0 *arm.Int16, v1 *arm.Int16, n int32)

// Compare signed Greater than or Equal (vector). This instruction compares each vector element in the first source SIMD&FP register with the corresponding vector element in the second source SIMD&FP register and if the first signed integer value is greater than or equal to the second signed integer value sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VcgeqS32N VcgeqS32N
//go:noescape
func VcgeqS32N(r *arm.Uint32, v0 *arm.Int32, v1 *arm.Int32, n int32)

// Compare signed Greater than or Equal (vector). This instruction compares each vector element in the first source SIMD&FP register with the corresponding vector element in the second source SIMD&FP register and if the first signed integer value is greater than or equal to the second signed integer value sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VcgeqS64N VcgeqS64N
//go:noescape
func VcgeqS64N(r *arm.Uint64, v0 *arm.Int64, v1 *arm.Int64, n int32)

// Compare unsigned Higher or Same (vector). This instruction compares each vector element in the first source SIMD&FP register with the corresponding vector element in the second source SIMD&FP register and if the first unsigned integer value is greater than or equal to the second unsigned integer value sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VcgeqU8N VcgeqU8N
//go:noescape
func VcgeqU8N(r *arm.Uint8, v0 *arm.Uint8, v1 *arm.Uint8, n int32)

// Compare unsigned Higher or Same (vector). This instruction compares each vector element in the first source SIMD&FP register with the corresponding vector element in the second source SIMD&FP register and if the first unsigned integer value is greater than or equal to the second unsigned integer value sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VcgeqU16N VcgeqU16N
//go:noescape
func VcgeqU16N(r *arm.Uint16, v0 *arm.Uint16, v1 *arm.Uint16, n int32)

// Compare unsigned Higher or Same (vector). This instruction compares each vector element in the first source SIMD&FP register with the corresponding vector element in the second source SIMD&FP register and if the first unsigned integer value is greater than or equal to the second unsigned integer value sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VcgeqU32N VcgeqU32N
//go:noescape
func VcgeqU32N(r *arm.Uint32, v0 *arm.Uint32, v1 *arm.Uint32, n int32)

// Compare unsigned Higher or Same (vector). This instruction compares each vector element in the first source SIMD&FP register with the corresponding vector element in the second source SIMD&FP register and if the first unsigned integer value is greater than or equal to the second unsigned integer value sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VcgeqU64N VcgeqU64N
//go:noescape
func VcgeqU64N(r *arm.Uint64, v0 *arm.Uint64, v1 *arm.Uint64, n int32)

// Floating-point Compare Greater than or Equal (vector). This instruction reads each floating-point value in the first source SIMD&FP register and if the value is greater than or equal to the corresponding floating-point value in the second source SIMD&FP register sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VcgeqF32N VcgeqF32N
//go:noescape
func VcgeqF32N(r *arm.Uint32, v0 *arm.Float32, v1 *arm.Float32, n int32)

// Floating-point Compare Greater than or Equal (vector). This instruction reads each floating-point value in the first source SIMD&FP register and if the value is greater than or equal to the corresponding floating-point value in the second source SIMD&FP register sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VcgeqF64N VcgeqF64N
//go:noescape
func VcgeqF64N(r *arm.Uint64, v0 *arm.Float64, v1 *arm.Float64, n int32)

// Floating-point Compare Greater than or Equal (vector). This instruction reads each floating-point value in the first source SIMD&FP register and if the value is greater than or equal to the corresponding floating-point value in the second source SIMD&FP register sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VcgesF32N VcgesF32N
//go:noescape
func VcgesF32N(r *arm.Uint32, v0 *arm.Float32, v1 *arm.Float32, n int32)

// Compare signed Greater than or Equal to zero (vector). This instruction reads each vector element in the source SIMD&FP register and if the signed integer value is greater than or equal to zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VcgezS8N VcgezS8N
//go:noescape
func VcgezS8N(r *arm.Uint8, v0 *arm.Int8, n int32)

// Compare signed Greater than or Equal to zero (vector). This instruction reads each vector element in the source SIMD&FP register and if the signed integer value is greater than or equal to zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VcgezS16N VcgezS16N
//go:noescape
func VcgezS16N(r *arm.Uint16, v0 *arm.Int16, n int32)

// Compare signed Greater than or Equal to zero (vector). This instruction reads each vector element in the source SIMD&FP register and if the signed integer value is greater than or equal to zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VcgezS32N VcgezS32N
//go:noescape
func VcgezS32N(r *arm.Uint32, v0 *arm.Int32, n int32)

// Compare signed Greater than or Equal to zero (vector). This instruction reads each vector element in the source SIMD&FP register and if the signed integer value is greater than or equal to zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VcgezS64N VcgezS64N
//go:noescape
func VcgezS64N(r *arm.Uint64, v0 *arm.Int64, n int32)

// Floating-point Compare Greater than or Equal to zero (vector). This instruction reads each floating-point value in the source SIMD&FP register and if the value is greater than or equal to zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VcgezF32N VcgezF32N
//go:noescape
func VcgezF32N(r *arm.Uint32, v0 *arm.Float32, n int32)

// Floating-point Compare Greater than or Equal to zero (vector). This instruction reads each floating-point value in the source SIMD&FP register and if the value is greater than or equal to zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VcgezF64N VcgezF64N
//go:noescape
func VcgezF64N(r *arm.Uint64, v0 *arm.Float64, n int32)

// Compare signed Greater than or Equal to zero (vector). This instruction reads each vector element in the source SIMD&FP register and if the signed integer value is greater than or equal to zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VcgezdS64N VcgezdS64N
//go:noescape
func VcgezdS64N(r *arm.Uint64, v0 *arm.Int64, n int32)

// Floating-point Compare Greater than or Equal to zero (vector). This instruction reads each floating-point value in the source SIMD&FP register and if the value is greater than or equal to zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VcgezdF64N VcgezdF64N
//go:noescape
func VcgezdF64N(r *arm.Uint64, v0 *arm.Float64, n int32)

// Compare signed Greater than or Equal to zero (vector). This instruction reads each vector element in the source SIMD&FP register and if the signed integer value is greater than or equal to zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VcgezqS8N VcgezqS8N
//go:noescape
func VcgezqS8N(r *arm.Uint8, v0 *arm.Int8, n int32)

// Compare signed Greater than or Equal to zero (vector). This instruction reads each vector element in the source SIMD&FP register and if the signed integer value is greater than or equal to zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VcgezqS16N VcgezqS16N
//go:noescape
func VcgezqS16N(r *arm.Uint16, v0 *arm.Int16, n int32)

// Compare signed Greater than or Equal to zero (vector). This instruction reads each vector element in the source SIMD&FP register and if the signed integer value is greater than or equal to zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VcgezqS32N VcgezqS32N
//go:noescape
func VcgezqS32N(r *arm.Uint32, v0 *arm.Int32, n int32)

// Compare signed Greater than or Equal to zero (vector). This instruction reads each vector element in the source SIMD&FP register and if the signed integer value is greater than or equal to zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VcgezqS64N VcgezqS64N
//go:noescape
func VcgezqS64N(r *arm.Uint64, v0 *arm.Int64, n int32)

// Floating-point Compare Greater than or Equal to zero (vector). This instruction reads each floating-point value in the source SIMD&FP register and if the value is greater than or equal to zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VcgezqF32N VcgezqF32N
//go:noescape
func VcgezqF32N(r *arm.Uint32, v0 *arm.Float32, n int32)

// Floating-point Compare Greater than or Equal to zero (vector). This instruction reads each floating-point value in the source SIMD&FP register and if the value is greater than or equal to zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VcgezqF64N VcgezqF64N
//go:noescape
func VcgezqF64N(r *arm.Uint64, v0 *arm.Float64, n int32)

// Floating-point Compare Greater than or Equal to zero (vector). This instruction reads each floating-point value in the source SIMD&FP register and if the value is greater than or equal to zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VcgezsF32N VcgezsF32N
//go:noescape
func VcgezsF32N(r *arm.Uint32, v0 *arm.Float32, n int32)

// Compare signed Greater than (vector). This instruction compares each vector element in the first source SIMD&FP register with the corresponding vector element in the second source SIMD&FP register and if the first signed integer value is greater than the second signed integer value sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VcgtS8N VcgtS8N
//go:noescape
func VcgtS8N(r *arm.Uint8, v0 *arm.Int8, v1 *arm.Int8, n int32)

// Compare signed Greater than (vector). This instruction compares each vector element in the first source SIMD&FP register with the corresponding vector element in the second source SIMD&FP register and if the first signed integer value is greater than the second signed integer value sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VcgtS16N VcgtS16N
//go:noescape
func VcgtS16N(r *arm.Uint16, v0 *arm.Int16, v1 *arm.Int16, n int32)

// Compare signed Greater than (vector). This instruction compares each vector element in the first source SIMD&FP register with the corresponding vector element in the second source SIMD&FP register and if the first signed integer value is greater than the second signed integer value sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VcgtS32N VcgtS32N
//go:noescape
func VcgtS32N(r *arm.Uint32, v0 *arm.Int32, v1 *arm.Int32, n int32)

// Compare signed Greater than (vector). This instruction compares each vector element in the first source SIMD&FP register with the corresponding vector element in the second source SIMD&FP register and if the first signed integer value is greater than the second signed integer value sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VcgtS64N VcgtS64N
//go:noescape
func VcgtS64N(r *arm.Uint64, v0 *arm.Int64, v1 *arm.Int64, n int32)

// Compare unsigned Higher (vector). This instruction compares each vector element in the first source SIMD&FP register with the corresponding vector element in the second source SIMD&FP register and if the first unsigned integer value is greater than the second unsigned integer value sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VcgtU8N VcgtU8N
//go:noescape
func VcgtU8N(r *arm.Uint8, v0 *arm.Uint8, v1 *arm.Uint8, n int32)

// Compare unsigned Higher (vector). This instruction compares each vector element in the first source SIMD&FP register with the corresponding vector element in the second source SIMD&FP register and if the first unsigned integer value is greater than the second unsigned integer value sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VcgtU16N VcgtU16N
//go:noescape
func VcgtU16N(r *arm.Uint16, v0 *arm.Uint16, v1 *arm.Uint16, n int32)

// Compare unsigned Higher (vector). This instruction compares each vector element in the first source SIMD&FP register with the corresponding vector element in the second source SIMD&FP register and if the first unsigned integer value is greater than the second unsigned integer value sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VcgtU32N VcgtU32N
//go:noescape
func VcgtU32N(r *arm.Uint32, v0 *arm.Uint32, v1 *arm.Uint32, n int32)

// Compare unsigned Higher (vector). This instruction compares each vector element in the first source SIMD&FP register with the corresponding vector element in the second source SIMD&FP register and if the first unsigned integer value is greater than the second unsigned integer value sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VcgtU64N VcgtU64N
//go:noescape
func VcgtU64N(r *arm.Uint64, v0 *arm.Uint64, v1 *arm.Uint64, n int32)

// Floating-point Compare Greater than (vector). This instruction reads each floating-point value in the first source SIMD&FP register and if the value is greater than the corresponding floating-point value in the second source SIMD&FP register sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VcgtF32N VcgtF32N
//go:noescape
func VcgtF32N(r *arm.Uint32, v0 *arm.Float32, v1 *arm.Float32, n int32)

// Floating-point Compare Greater than (vector). This instruction reads each floating-point value in the first source SIMD&FP register and if the value is greater than the corresponding floating-point value in the second source SIMD&FP register sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VcgtF64N VcgtF64N
//go:noescape
func VcgtF64N(r *arm.Uint64, v0 *arm.Float64, v1 *arm.Float64, n int32)

// Compare signed Greater than (vector). This instruction compares each vector element in the first source SIMD&FP register with the corresponding vector element in the second source SIMD&FP register and if the first signed integer value is greater than the second signed integer value sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VcgtdS64N VcgtdS64N
//go:noescape
func VcgtdS64N(r *arm.Uint64, v0 *arm.Int64, v1 *arm.Int64, n int32)

// Compare unsigned Higher (vector). This instruction compares each vector element in the first source SIMD&FP register with the corresponding vector element in the second source SIMD&FP register and if the first unsigned integer value is greater than the second unsigned integer value sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VcgtdU64N VcgtdU64N
//go:noescape
func VcgtdU64N(r *arm.Uint64, v0 *arm.Uint64, v1 *arm.Uint64, n int32)

// Floating-point Compare Greater than (vector). This instruction reads each floating-point value in the first source SIMD&FP register and if the value is greater than the corresponding floating-point value in the second source SIMD&FP register sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VcgtdF64N VcgtdF64N
//go:noescape
func VcgtdF64N(r *arm.Uint64, v0 *arm.Float64, v1 *arm.Float64, n int32)

// Compare signed Greater than (vector). This instruction compares each vector element in the first source SIMD&FP register with the corresponding vector element in the second source SIMD&FP register and if the first signed integer value is greater than the second signed integer value sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VcgtqS8N VcgtqS8N
//go:noescape
func VcgtqS8N(r *arm.Uint8, v0 *arm.Int8, v1 *arm.Int8, n int32)

// Compare signed Greater than (vector). This instruction compares each vector element in the first source SIMD&FP register with the corresponding vector element in the second source SIMD&FP register and if the first signed integer value is greater than the second signed integer value sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VcgtqS16N VcgtqS16N
//go:noescape
func VcgtqS16N(r *arm.Uint16, v0 *arm.Int16, v1 *arm.Int16, n int32)

// Compare signed Greater than (vector). This instruction compares each vector element in the first source SIMD&FP register with the corresponding vector element in the second source SIMD&FP register and if the first signed integer value is greater than the second signed integer value sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VcgtqS32N VcgtqS32N
//go:noescape
func VcgtqS32N(r *arm.Uint32, v0 *arm.Int32, v1 *arm.Int32, n int32)

// Compare signed Greater than (vector). This instruction compares each vector element in the first source SIMD&FP register with the corresponding vector element in the second source SIMD&FP register and if the first signed integer value is greater than the second signed integer value sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VcgtqS64N VcgtqS64N
//go:noescape
func VcgtqS64N(r *arm.Uint64, v0 *arm.Int64, v1 *arm.Int64, n int32)

// Compare unsigned Higher (vector). This instruction compares each vector element in the first source SIMD&FP register with the corresponding vector element in the second source SIMD&FP register and if the first unsigned integer value is greater than the second unsigned integer value sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VcgtqU8N VcgtqU8N
//go:noescape
func VcgtqU8N(r *arm.Uint8, v0 *arm.Uint8, v1 *arm.Uint8, n int32)

// Compare unsigned Higher (vector). This instruction compares each vector element in the first source SIMD&FP register with the corresponding vector element in the second source SIMD&FP register and if the first unsigned integer value is greater than the second unsigned integer value sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VcgtqU16N VcgtqU16N
//go:noescape
func VcgtqU16N(r *arm.Uint16, v0 *arm.Uint16, v1 *arm.Uint16, n int32)

// Compare unsigned Higher (vector). This instruction compares each vector element in the first source SIMD&FP register with the corresponding vector element in the second source SIMD&FP register and if the first unsigned integer value is greater than the second unsigned integer value sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VcgtqU32N VcgtqU32N
//go:noescape
func VcgtqU32N(r *arm.Uint32, v0 *arm.Uint32, v1 *arm.Uint32, n int32)

// Compare unsigned Higher (vector). This instruction compares each vector element in the first source SIMD&FP register with the corresponding vector element in the second source SIMD&FP register and if the first unsigned integer value is greater than the second unsigned integer value sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VcgtqU64N VcgtqU64N
//go:noescape
func VcgtqU64N(r *arm.Uint64, v0 *arm.Uint64, v1 *arm.Uint64, n int32)

// Floating-point Compare Greater than (vector). This instruction reads each floating-point value in the first source SIMD&FP register and if the value is greater than the corresponding floating-point value in the second source SIMD&FP register sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VcgtqF32N VcgtqF32N
//go:noescape
func VcgtqF32N(r *arm.Uint32, v0 *arm.Float32, v1 *arm.Float32, n int32)

// Floating-point Compare Greater than (vector). This instruction reads each floating-point value in the first source SIMD&FP register and if the value is greater than the corresponding floating-point value in the second source SIMD&FP register sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VcgtqF64N VcgtqF64N
//go:noescape
func VcgtqF64N(r *arm.Uint64, v0 *arm.Float64, v1 *arm.Float64, n int32)

// Floating-point Compare Greater than (vector). This instruction reads each floating-point value in the first source SIMD&FP register and if the value is greater than the corresponding floating-point value in the second source SIMD&FP register sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VcgtsF32N VcgtsF32N
//go:noescape
func VcgtsF32N(r *arm.Uint32, v0 *arm.Float32, v1 *arm.Float32, n int32)

// Compare signed Greater than zero (vector). This instruction reads each vector element in the source SIMD&FP register and if the signed integer value is greater than zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VcgtzS8N VcgtzS8N
//go:noescape
func VcgtzS8N(r *arm.Uint8, v0 *arm.Int8, n int32)

// Compare signed Greater than zero (vector). This instruction reads each vector element in the source SIMD&FP register and if the signed integer value is greater than zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VcgtzS16N VcgtzS16N
//go:noescape
func VcgtzS16N(r *arm.Uint16, v0 *arm.Int16, n int32)

// Compare signed Greater than zero (vector). This instruction reads each vector element in the source SIMD&FP register and if the signed integer value is greater than zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VcgtzS32N VcgtzS32N
//go:noescape
func VcgtzS32N(r *arm.Uint32, v0 *arm.Int32, n int32)

// Compare signed Greater than zero (vector). This instruction reads each vector element in the source SIMD&FP register and if the signed integer value is greater than zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VcgtzS64N VcgtzS64N
//go:noescape
func VcgtzS64N(r *arm.Uint64, v0 *arm.Int64, n int32)

// Floating-point Compare Greater than zero (vector). This instruction reads each floating-point value in the source SIMD&FP register and if the value is greater than zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VcgtzF32N VcgtzF32N
//go:noescape
func VcgtzF32N(r *arm.Uint32, v0 *arm.Float32, n int32)

// Floating-point Compare Greater than zero (vector). This instruction reads each floating-point value in the source SIMD&FP register and if the value is greater than zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VcgtzF64N VcgtzF64N
//go:noescape
func VcgtzF64N(r *arm.Uint64, v0 *arm.Float64, n int32)

// Compare signed Greater than zero (vector). This instruction reads each vector element in the source SIMD&FP register and if the signed integer value is greater than zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VcgtzdS64N VcgtzdS64N
//go:noescape
func VcgtzdS64N(r *arm.Uint64, v0 *arm.Int64, n int32)

// Floating-point Compare Greater than zero (vector). This instruction reads each floating-point value in the source SIMD&FP register and if the value is greater than zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VcgtzdF64N VcgtzdF64N
//go:noescape
func VcgtzdF64N(r *arm.Uint64, v0 *arm.Float64, n int32)

// Compare signed Greater than zero (vector). This instruction reads each vector element in the source SIMD&FP register and if the signed integer value is greater than zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VcgtzqS8N VcgtzqS8N
//go:noescape
func VcgtzqS8N(r *arm.Uint8, v0 *arm.Int8, n int32)

// Compare signed Greater than zero (vector). This instruction reads each vector element in the source SIMD&FP register and if the signed integer value is greater than zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VcgtzqS16N VcgtzqS16N
//go:noescape
func VcgtzqS16N(r *arm.Uint16, v0 *arm.Int16, n int32)

// Compare signed Greater than zero (vector). This instruction reads each vector element in the source SIMD&FP register and if the signed integer value is greater than zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VcgtzqS32N VcgtzqS32N
//go:noescape
func VcgtzqS32N(r *arm.Uint32, v0 *arm.Int32, n int32)

// Compare signed Greater than zero (vector). This instruction reads each vector element in the source SIMD&FP register and if the signed integer value is greater than zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VcgtzqS64N VcgtzqS64N
//go:noescape
func VcgtzqS64N(r *arm.Uint64, v0 *arm.Int64, n int32)

// Floating-point Compare Greater than zero (vector). This instruction reads each floating-point value in the source SIMD&FP register and if the value is greater than zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VcgtzqF32N VcgtzqF32N
//go:noescape
func VcgtzqF32N(r *arm.Uint32, v0 *arm.Float32, n int32)

// Floating-point Compare Greater than zero (vector). This instruction reads each floating-point value in the source SIMD&FP register and if the value is greater than zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VcgtzqF64N VcgtzqF64N
//go:noescape
func VcgtzqF64N(r *arm.Uint64, v0 *arm.Float64, n int32)

// Floating-point Compare Greater than zero (vector). This instruction reads each floating-point value in the source SIMD&FP register and if the value is greater than zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VcgtzsF32N VcgtzsF32N
//go:noescape
func VcgtzsF32N(r *arm.Uint32, v0 *arm.Float32, n int32)

// Compare signed less than or equal
//
//go:linkname VcleS8N VcleS8N
//go:noescape
func VcleS8N(r *arm.Uint8, v0 *arm.Int8, v1 *arm.Int8, n int32)

// Compare signed less than or equal
//
//go:linkname VcleS16N VcleS16N
//go:noescape
func VcleS16N(r *arm.Uint16, v0 *arm.Int16, v1 *arm.Int16, n int32)

// Compare signed less than or equal
//
//go:linkname VcleS32N VcleS32N
//go:noescape
func VcleS32N(r *arm.Uint32, v0 *arm.Int32, v1 *arm.Int32, n int32)

// Compare signed less than or equal
//
//go:linkname VcleS64N VcleS64N
//go:noescape
func VcleS64N(r *arm.Uint64, v0 *arm.Int64, v1 *arm.Int64, n int32)

// Compare unsigned less than or equal
//
//go:linkname VcleU8N VcleU8N
//go:noescape
func VcleU8N(r *arm.Uint8, v0 *arm.Uint8, v1 *arm.Uint8, n int32)

// Compare unsigned less than or equal
//
//go:linkname VcleU16N VcleU16N
//go:noescape
func VcleU16N(r *arm.Uint16, v0 *arm.Uint16, v1 *arm.Uint16, n int32)

// Compare unsigned less than or equal
//
//go:linkname VcleU32N VcleU32N
//go:noescape
func VcleU32N(r *arm.Uint32, v0 *arm.Uint32, v1 *arm.Uint32, n int32)

// Compare unsigned less than or equal
//
//go:linkname VcleU64N VcleU64N
//go:noescape
func VcleU64N(r *arm.Uint64, v0 *arm.Uint64, v1 *arm.Uint64, n int32)

// Floating-point compare less than or equal
//
//go:linkname VcleF32N VcleF32N
//go:noescape
func VcleF32N(r *arm.Uint32, v0 *arm.Float32, v1 *arm.Float32, n int32)

// Floating-point compare less than or equal
//
//go:linkname VcleF64N VcleF64N
//go:noescape
func VcleF64N(r *arm.Uint64, v0 *arm.Float64, v1 *arm.Float64, n int32)

// Compare signed less than or equal
//
//go:linkname VcledS64N VcledS64N
//go:noescape
func VcledS64N(r *arm.Uint64, v0 *arm.Int64, v1 *arm.Int64, n int32)

// Compare unsigned less than or equal
//
//go:linkname VcledU64N VcledU64N
//go:noescape
func VcledU64N(r *arm.Uint64, v0 *arm.Uint64, v1 *arm.Uint64, n int32)

// Floating-point compare less than or equal
//
//go:linkname VcledF64N VcledF64N
//go:noescape
func VcledF64N(r *arm.Uint64, v0 *arm.Float64, v1 *arm.Float64, n int32)

// Compare signed less than or equal
//
//go:linkname VcleqS8N VcleqS8N
//go:noescape
func VcleqS8N(r *arm.Uint8, v0 *arm.Int8, v1 *arm.Int8, n int32)

// Compare signed less than or equal
//
//go:linkname VcleqS16N VcleqS16N
//go:noescape
func VcleqS16N(r *arm.Uint16, v0 *arm.Int16, v1 *arm.Int16, n int32)

// Compare signed less than or equal
//
//go:linkname VcleqS32N VcleqS32N
//go:noescape
func VcleqS32N(r *arm.Uint32, v0 *arm.Int32, v1 *arm.Int32, n int32)

// Compare signed less than or equal
//
//go:linkname VcleqS64N VcleqS64N
//go:noescape
func VcleqS64N(r *arm.Uint64, v0 *arm.Int64, v1 *arm.Int64, n int32)

// Compare unsigned less than or equal
//
//go:linkname VcleqU8N VcleqU8N
//go:noescape
func VcleqU8N(r *arm.Uint8, v0 *arm.Uint8, v1 *arm.Uint8, n int32)

// Compare unsigned less than or equal
//
//go:linkname VcleqU16N VcleqU16N
//go:noescape
func VcleqU16N(r *arm.Uint16, v0 *arm.Uint16, v1 *arm.Uint16, n int32)

// Compare unsigned less than or equal
//
//go:linkname VcleqU32N VcleqU32N
//go:noescape
func VcleqU32N(r *arm.Uint32, v0 *arm.Uint32, v1 *arm.Uint32, n int32)

// Compare unsigned less than or equal
//
//go:linkname VcleqU64N VcleqU64N
//go:noescape
func VcleqU64N(r *arm.Uint64, v0 *arm.Uint64, v1 *arm.Uint64, n int32)

// Floating-point compare less than or equal
//
//go:linkname VcleqF32N VcleqF32N
//go:noescape
func VcleqF32N(r *arm.Uint32, v0 *arm.Float32, v1 *arm.Float32, n int32)

// Floating-point compare less than or equal
//
//go:linkname VcleqF64N VcleqF64N
//go:noescape
func VcleqF64N(r *arm.Uint64, v0 *arm.Float64, v1 *arm.Float64, n int32)

// Floating-point compare less than or equal
//
//go:linkname VclesF32N VclesF32N
//go:noescape
func VclesF32N(r *arm.Uint32, v0 *arm.Float32, v1 *arm.Float32, n int32)

// Compare signed Less than or Equal to zero (vector). This instruction reads each vector element in the source SIMD&FP register and if the signed integer value is less than or equal to zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VclezS8N VclezS8N
//go:noescape
func VclezS8N(r *arm.Uint8, v0 *arm.Int8, n int32)

// Compare signed Less than or Equal to zero (vector). This instruction reads each vector element in the source SIMD&FP register and if the signed integer value is less than or equal to zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VclezS16N VclezS16N
//go:noescape
func VclezS16N(r *arm.Uint16, v0 *arm.Int16, n int32)

// Compare signed Less than or Equal to zero (vector). This instruction reads each vector element in the source SIMD&FP register and if the signed integer value is less than or equal to zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VclezS32N VclezS32N
//go:noescape
func VclezS32N(r *arm.Uint32, v0 *arm.Int32, n int32)

// Compare signed Less than or Equal to zero (vector). This instruction reads each vector element in the source SIMD&FP register and if the signed integer value is less than or equal to zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VclezS64N VclezS64N
//go:noescape
func VclezS64N(r *arm.Uint64, v0 *arm.Int64, n int32)

// Compare signed Less than or Equal to zero (vector). This instruction reads each vector element in the source SIMD&FP register and if the signed integer value is less than or equal to zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VclezF32N VclezF32N
//go:noescape
func VclezF32N(r *arm.Uint32, v0 *arm.Float32, n int32)

// Floating-point Compare Less than or Equal to zero (vector). This instruction reads each floating-point value in the source SIMD&FP register and if the value is less than or equal to zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VclezF64N VclezF64N
//go:noescape
func VclezF64N(r *arm.Uint64, v0 *arm.Float64, n int32)

// Compare signed Less than or Equal to zero (vector). This instruction reads each vector element in the source SIMD&FP register and if the signed integer value is less than or equal to zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VclezdS64N VclezdS64N
//go:noescape
func VclezdS64N(r *arm.Uint64, v0 *arm.Int64, n int32)

// Floating-point Compare Less than or Equal to zero (vector). This instruction reads each floating-point value in the source SIMD&FP register and if the value is less than or equal to zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VclezdF64N VclezdF64N
//go:noescape
func VclezdF64N(r *arm.Uint64, v0 *arm.Float64, n int32)

// Compare signed Less than or Equal to zero (vector). This instruction reads each vector element in the source SIMD&FP register and if the signed integer value is less than or equal to zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VclezqS8N VclezqS8N
//go:noescape
func VclezqS8N(r *arm.Uint8, v0 *arm.Int8, n int32)

// Compare signed Less than or Equal to zero (vector). This instruction reads each vector element in the source SIMD&FP register and if the signed integer value is less than or equal to zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VclezqS16N VclezqS16N
//go:noescape
func VclezqS16N(r *arm.Uint16, v0 *arm.Int16, n int32)

// Compare signed Less than or Equal to zero (vector). This instruction reads each vector element in the source SIMD&FP register and if the signed integer value is less than or equal to zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VclezqS32N VclezqS32N
//go:noescape
func VclezqS32N(r *arm.Uint32, v0 *arm.Int32, n int32)

// Compare signed Less than or Equal to zero (vector). This instruction reads each vector element in the source SIMD&FP register and if the signed integer value is less than or equal to zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VclezqS64N VclezqS64N
//go:noescape
func VclezqS64N(r *arm.Uint64, v0 *arm.Int64, n int32)

// Floating-point Compare Less than or Equal to zero (vector). This instruction reads each floating-point value in the source SIMD&FP register and if the value is less than or equal to zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VclezqF32N VclezqF32N
//go:noescape
func VclezqF32N(r *arm.Uint32, v0 *arm.Float32, n int32)

// Floating-point Compare Less than or Equal to zero (vector). This instruction reads each floating-point value in the source SIMD&FP register and if the value is less than or equal to zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VclezqF64N VclezqF64N
//go:noescape
func VclezqF64N(r *arm.Uint64, v0 *arm.Float64, n int32)

// Floating-point Compare Less than or Equal to zero (vector). This instruction reads each floating-point value in the source SIMD&FP register and if the value is less than or equal to zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VclezsF32N VclezsF32N
//go:noescape
func VclezsF32N(r *arm.Uint32, v0 *arm.Float32, n int32)

// Count Leading Sign bits (vector). This instruction counts the number of consecutive bits following the most significant bit that are the same as the most significant bit in each vector element in the source SIMD&FP register, places the result into a vector, and writes the vector to the destination SIMD&FP register. The count does not include the most significant bit itself.
//
//go:linkname VclsS8N VclsS8N
//go:noescape
func VclsS8N(r *arm.Int8, v0 *arm.Int8, n int32)

// Count Leading Sign bits (vector). This instruction counts the number of consecutive bits following the most significant bit that are the same as the most significant bit in each vector element in the source SIMD&FP register, places the result into a vector, and writes the vector to the destination SIMD&FP register. The count does not include the most significant bit itself.
//
//go:linkname VclsS16N VclsS16N
//go:noescape
func VclsS16N(r *arm.Int16, v0 *arm.Int16, n int32)

// Count Leading Sign bits (vector). This instruction counts the number of consecutive bits following the most significant bit that are the same as the most significant bit in each vector element in the source SIMD&FP register, places the result into a vector, and writes the vector to the destination SIMD&FP register. The count does not include the most significant bit itself.
//
//go:linkname VclsS32N VclsS32N
//go:noescape
func VclsS32N(r *arm.Int32, v0 *arm.Int32, n int32)

// Count Leading Sign bits (vector). This instruction counts the number of consecutive bits following the most significant bit that are the same as the most significant bit in each vector element in the source SIMD&FP register, places the result into a vector, and writes the vector to the destination SIMD&FP register. The count does not include the most significant bit itself.
//
//go:linkname VclsU8N VclsU8N
//go:noescape
func VclsU8N(r *arm.Int8, v0 *arm.Uint8, n int32)

// Count Leading Sign bits (vector). This instruction counts the number of consecutive bits following the most significant bit that are the same as the most significant bit in each vector element in the source SIMD&FP register, places the result into a vector, and writes the vector to the destination SIMD&FP register. The count does not include the most significant bit itself.
//
//go:linkname VclsU16N VclsU16N
//go:noescape
func VclsU16N(r *arm.Int16, v0 *arm.Uint16, n int32)

// Count Leading Sign bits (vector). This instruction counts the number of consecutive bits following the most significant bit that are the same as the most significant bit in each vector element in the source SIMD&FP register, places the result into a vector, and writes the vector to the destination SIMD&FP register. The count does not include the most significant bit itself.
//
//go:linkname VclsU32N VclsU32N
//go:noescape
func VclsU32N(r *arm.Int32, v0 *arm.Uint32, n int32)

// Count Leading Sign bits (vector). This instruction counts the number of consecutive bits following the most significant bit that are the same as the most significant bit in each vector element in the source SIMD&FP register, places the result into a vector, and writes the vector to the destination SIMD&FP register. The count does not include the most significant bit itself.
//
//go:linkname VclsqS8N VclsqS8N
//go:noescape
func VclsqS8N(r *arm.Int8, v0 *arm.Int8, n int32)

// Count Leading Sign bits (vector). This instruction counts the number of consecutive bits following the most significant bit that are the same as the most significant bit in each vector element in the source SIMD&FP register, places the result into a vector, and writes the vector to the destination SIMD&FP register. The count does not include the most significant bit itself.
//
//go:linkname VclsqS16N VclsqS16N
//go:noescape
func VclsqS16N(r *arm.Int16, v0 *arm.Int16, n int32)

// Count Leading Sign bits (vector). This instruction counts the number of consecutive bits following the most significant bit that are the same as the most significant bit in each vector element in the source SIMD&FP register, places the result into a vector, and writes the vector to the destination SIMD&FP register. The count does not include the most significant bit itself.
//
//go:linkname VclsqS32N VclsqS32N
//go:noescape
func VclsqS32N(r *arm.Int32, v0 *arm.Int32, n int32)

// Count Leading Sign bits (vector). This instruction counts the number of consecutive bits following the most significant bit that are the same as the most significant bit in each vector element in the source SIMD&FP register, places the result into a vector, and writes the vector to the destination SIMD&FP register. The count does not include the most significant bit itself.
//
//go:linkname VclsqU8N VclsqU8N
//go:noescape
func VclsqU8N(r *arm.Int8, v0 *arm.Uint8, n int32)

// Count Leading Sign bits (vector). This instruction counts the number of consecutive bits following the most significant bit that are the same as the most significant bit in each vector element in the source SIMD&FP register, places the result into a vector, and writes the vector to the destination SIMD&FP register. The count does not include the most significant bit itself.
//
//go:linkname VclsqU16N VclsqU16N
//go:noescape
func VclsqU16N(r *arm.Int16, v0 *arm.Uint16, n int32)

// Count Leading Sign bits (vector). This instruction counts the number of consecutive bits following the most significant bit that are the same as the most significant bit in each vector element in the source SIMD&FP register, places the result into a vector, and writes the vector to the destination SIMD&FP register. The count does not include the most significant bit itself.
//
//go:linkname VclsqU32N VclsqU32N
//go:noescape
func VclsqU32N(r *arm.Int32, v0 *arm.Uint32, n int32)

// Compare signed less than
//
//go:linkname VcltS8N VcltS8N
//go:noescape
func VcltS8N(r *arm.Uint8, v0 *arm.Int8, v1 *arm.Int8, n int32)

// Compare signed less than
//
//go:linkname VcltS16N VcltS16N
//go:noescape
func VcltS16N(r *arm.Uint16, v0 *arm.Int16, v1 *arm.Int16, n int32)

// Compare signed less than
//
//go:linkname VcltS32N VcltS32N
//go:noescape
func VcltS32N(r *arm.Uint32, v0 *arm.Int32, v1 *arm.Int32, n int32)

// Compare signed less than
//
//go:linkname VcltS64N VcltS64N
//go:noescape
func VcltS64N(r *arm.Uint64, v0 *arm.Int64, v1 *arm.Int64, n int32)

// Compare unsigned less than
//
//go:linkname VcltU8N VcltU8N
//go:noescape
func VcltU8N(r *arm.Uint8, v0 *arm.Uint8, v1 *arm.Uint8, n int32)

// Compare unsigned less than
//
//go:linkname VcltU16N VcltU16N
//go:noescape
func VcltU16N(r *arm.Uint16, v0 *arm.Uint16, v1 *arm.Uint16, n int32)

// Compare unsigned less than
//
//go:linkname VcltU32N VcltU32N
//go:noescape
func VcltU32N(r *arm.Uint32, v0 *arm.Uint32, v1 *arm.Uint32, n int32)

// Compare unsigned less than
//
//go:linkname VcltU64N VcltU64N
//go:noescape
func VcltU64N(r *arm.Uint64, v0 *arm.Uint64, v1 *arm.Uint64, n int32)

// Floating-point compare less than
//
//go:linkname VcltF32N VcltF32N
//go:noescape
func VcltF32N(r *arm.Uint32, v0 *arm.Float32, v1 *arm.Float32, n int32)

// Floating-point compare less than
//
//go:linkname VcltF64N VcltF64N
//go:noescape
func VcltF64N(r *arm.Uint64, v0 *arm.Float64, v1 *arm.Float64, n int32)

// Compare signed less than
//
//go:linkname VcltdS64N VcltdS64N
//go:noescape
func VcltdS64N(r *arm.Uint64, v0 *arm.Int64, v1 *arm.Int64, n int32)

// Compare unsigned less than
//
//go:linkname VcltdU64N VcltdU64N
//go:noescape
func VcltdU64N(r *arm.Uint64, v0 *arm.Uint64, v1 *arm.Uint64, n int32)

// Floating-point compare less than
//
//go:linkname VcltdF64N VcltdF64N
//go:noescape
func VcltdF64N(r *arm.Uint64, v0 *arm.Float64, v1 *arm.Float64, n int32)

// Compare signed less than
//
//go:linkname VcltqS8N VcltqS8N
//go:noescape
func VcltqS8N(r *arm.Uint8, v0 *arm.Int8, v1 *arm.Int8, n int32)

// Compare signed less than
//
//go:linkname VcltqS16N VcltqS16N
//go:noescape
func VcltqS16N(r *arm.Uint16, v0 *arm.Int16, v1 *arm.Int16, n int32)

// Compare signed less than
//
//go:linkname VcltqS32N VcltqS32N
//go:noescape
func VcltqS32N(r *arm.Uint32, v0 *arm.Int32, v1 *arm.Int32, n int32)

// Compare signed less than
//
//go:linkname VcltqS64N VcltqS64N
//go:noescape
func VcltqS64N(r *arm.Uint64, v0 *arm.Int64, v1 *arm.Int64, n int32)

// Compare unsigned less than
//
//go:linkname VcltqU8N VcltqU8N
//go:noescape
func VcltqU8N(r *arm.Uint8, v0 *arm.Uint8, v1 *arm.Uint8, n int32)

// Compare unsigned less than
//
//go:linkname VcltqU16N VcltqU16N
//go:noescape
func VcltqU16N(r *arm.Uint16, v0 *arm.Uint16, v1 *arm.Uint16, n int32)

// Compare unsigned less than
//
//go:linkname VcltqU32N VcltqU32N
//go:noescape
func VcltqU32N(r *arm.Uint32, v0 *arm.Uint32, v1 *arm.Uint32, n int32)

// Compare unsigned less than
//
//go:linkname VcltqU64N VcltqU64N
//go:noescape
func VcltqU64N(r *arm.Uint64, v0 *arm.Uint64, v1 *arm.Uint64, n int32)

// Floating-point compare less than
//
//go:linkname VcltqF32N VcltqF32N
//go:noescape
func VcltqF32N(r *arm.Uint32, v0 *arm.Float32, v1 *arm.Float32, n int32)

// Floating-point compare less than
//
//go:linkname VcltqF64N VcltqF64N
//go:noescape
func VcltqF64N(r *arm.Uint64, v0 *arm.Float64, v1 *arm.Float64, n int32)

// Floating-point compare less than
//
//go:linkname VcltsF32N VcltsF32N
//go:noescape
func VcltsF32N(r *arm.Uint32, v0 *arm.Float32, v1 *arm.Float32, n int32)

// Compare signed Less than zero (vector). This instruction reads each vector element in the source SIMD&FP register and if the signed integer value is less than zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VcltzS8N VcltzS8N
//go:noescape
func VcltzS8N(r *arm.Uint8, v0 *arm.Int8, n int32)

// Compare signed Less than zero (vector). This instruction reads each vector element in the source SIMD&FP register and if the signed integer value is less than zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VcltzS16N VcltzS16N
//go:noescape
func VcltzS16N(r *arm.Uint16, v0 *arm.Int16, n int32)

// Compare signed Less than zero (vector). This instruction reads each vector element in the source SIMD&FP register and if the signed integer value is less than zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VcltzS32N VcltzS32N
//go:noescape
func VcltzS32N(r *arm.Uint32, v0 *arm.Int32, n int32)

// Compare signed Less than zero (vector). This instruction reads each vector element in the source SIMD&FP register and if the signed integer value is less than zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VcltzS64N VcltzS64N
//go:noescape
func VcltzS64N(r *arm.Uint64, v0 *arm.Int64, n int32)

// Floating-point Compare Less than zero (vector). This instruction reads each floating-point value in the source SIMD&FP register and if the value is less than zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VcltzF32N VcltzF32N
//go:noescape
func VcltzF32N(r *arm.Uint32, v0 *arm.Float32, n int32)

// Floating-point Compare Less than zero (vector). This instruction reads each floating-point value in the source SIMD&FP register and if the value is less than zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VcltzF64N VcltzF64N
//go:noescape
func VcltzF64N(r *arm.Uint64, v0 *arm.Float64, n int32)

// Compare signed Less than zero (vector). This instruction reads each vector element in the source SIMD&FP register and if the signed integer value is less than zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VcltzdS64N VcltzdS64N
//go:noescape
func VcltzdS64N(r *arm.Uint64, v0 *arm.Int64, n int32)

// Floating-point Compare Less than zero (vector). This instruction reads each floating-point value in the source SIMD&FP register and if the value is less than zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VcltzdF64N VcltzdF64N
//go:noescape
func VcltzdF64N(r *arm.Uint64, v0 *arm.Float64, n int32)

// Compare signed Less than zero (vector). This instruction reads each vector element in the source SIMD&FP register and if the signed integer value is less than zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VcltzqS8N VcltzqS8N
//go:noescape
func VcltzqS8N(r *arm.Uint8, v0 *arm.Int8, n int32)

// Compare signed Less than zero (vector). This instruction reads each vector element in the source SIMD&FP register and if the signed integer value is less than zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VcltzqS16N VcltzqS16N
//go:noescape
func VcltzqS16N(r *arm.Uint16, v0 *arm.Int16, n int32)

// Compare signed Less than zero (vector). This instruction reads each vector element in the source SIMD&FP register and if the signed integer value is less than zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VcltzqS32N VcltzqS32N
//go:noescape
func VcltzqS32N(r *arm.Uint32, v0 *arm.Int32, n int32)

// Compare signed Less than zero (vector). This instruction reads each vector element in the source SIMD&FP register and if the signed integer value is less than zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VcltzqS64N VcltzqS64N
//go:noescape
func VcltzqS64N(r *arm.Uint64, v0 *arm.Int64, n int32)

// Floating-point Compare Less than zero (vector). This instruction reads each floating-point value in the source SIMD&FP register and if the value is less than zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VcltzqF32N VcltzqF32N
//go:noescape
func VcltzqF32N(r *arm.Uint32, v0 *arm.Float32, n int32)

// Floating-point Compare Less than zero (vector). This instruction reads each floating-point value in the source SIMD&FP register and if the value is less than zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VcltzqF64N VcltzqF64N
//go:noescape
func VcltzqF64N(r *arm.Uint64, v0 *arm.Float64, n int32)

// Floating-point Compare Less than zero (vector). This instruction reads each floating-point value in the source SIMD&FP register and if the value is less than zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VcltzsF32N VcltzsF32N
//go:noescape
func VcltzsF32N(r *arm.Uint32, v0 *arm.Float32, n int32)

// Count Leading Zero bits (vector). This instruction counts the number of consecutive zeros, starting from the most significant bit, in each vector element in the source SIMD&FP register, places the result into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VclzS8N VclzS8N
//go:noescape
func VclzS8N(r *arm.Int8, v0 *arm.Int8, n int32)

// Count Leading Zero bits (vector). This instruction counts the number of consecutive zeros, starting from the most significant bit, in each vector element in the source SIMD&FP register, places the result into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VclzS16N VclzS16N
//go:noescape
func VclzS16N(r *arm.Int16, v0 *arm.Int16, n int32)

// Count Leading Zero bits (vector). This instruction counts the number of consecutive zeros, starting from the most significant bit, in each vector element in the source SIMD&FP register, places the result into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VclzS32N VclzS32N
//go:noescape
func VclzS32N(r *arm.Int32, v0 *arm.Int32, n int32)

// Count Leading Zero bits (vector). This instruction counts the number of consecutive zeros, starting from the most significant bit, in each vector element in the source SIMD&FP register, places the result into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VclzU8N VclzU8N
//go:noescape
func VclzU8N(r *arm.Uint8, v0 *arm.Uint8, n int32)

// Count Leading Zero bits (vector). This instruction counts the number of consecutive zeros, starting from the most significant bit, in each vector element in the source SIMD&FP register, places the result into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VclzU16N VclzU16N
//go:noescape
func VclzU16N(r *arm.Uint16, v0 *arm.Uint16, n int32)

// Count Leading Zero bits (vector). This instruction counts the number of consecutive zeros, starting from the most significant bit, in each vector element in the source SIMD&FP register, places the result into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VclzU32N VclzU32N
//go:noescape
func VclzU32N(r *arm.Uint32, v0 *arm.Uint32, n int32)

// Count Leading Zero bits (vector). This instruction counts the number of consecutive zeros, starting from the most significant bit, in each vector element in the source SIMD&FP register, places the result into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VclzqS8N VclzqS8N
//go:noescape
func VclzqS8N(r *arm.Int8, v0 *arm.Int8, n int32)

// Count Leading Zero bits (vector). This instruction counts the number of consecutive zeros, starting from the most significant bit, in each vector element in the source SIMD&FP register, places the result into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VclzqS16N VclzqS16N
//go:noescape
func VclzqS16N(r *arm.Int16, v0 *arm.Int16, n int32)

// Count Leading Zero bits (vector). This instruction counts the number of consecutive zeros, starting from the most significant bit, in each vector element in the source SIMD&FP register, places the result into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VclzqS32N VclzqS32N
//go:noescape
func VclzqS32N(r *arm.Int32, v0 *arm.Int32, n int32)

// Count Leading Zero bits (vector). This instruction counts the number of consecutive zeros, starting from the most significant bit, in each vector element in the source SIMD&FP register, places the result into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VclzqU8N VclzqU8N
//go:noescape
func VclzqU8N(r *arm.Uint8, v0 *arm.Uint8, n int32)

// Count Leading Zero bits (vector). This instruction counts the number of consecutive zeros, starting from the most significant bit, in each vector element in the source SIMD&FP register, places the result into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VclzqU16N VclzqU16N
//go:noescape
func VclzqU16N(r *arm.Uint16, v0 *arm.Uint16, n int32)

// Count Leading Zero bits (vector). This instruction counts the number of consecutive zeros, starting from the most significant bit, in each vector element in the source SIMD&FP register, places the result into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VclzqU32N VclzqU32N
//go:noescape
func VclzqU32N(r *arm.Uint32, v0 *arm.Uint32, n int32)

// Population Count per byte. This instruction counts the number of bits that have a value of one in each vector element in the source SIMD&FP register, places the result into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VcntS8N VcntS8N
//go:noescape
func VcntS8N(r *arm.Int8, v0 *arm.Int8, n int32)

// Population Count per byte. This instruction counts the number of bits that have a value of one in each vector element in the source SIMD&FP register, places the result into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VcntU8N VcntU8N
//go:noescape
func VcntU8N(r *arm.Uint8, v0 *arm.Uint8, n int32)

// Population Count per byte. This instruction counts the number of bits that have a value of one in each vector element in the source SIMD&FP register, places the result into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VcntqS8N VcntqS8N
//go:noescape
func VcntqS8N(r *arm.Int8, v0 *arm.Int8, n int32)

// Population Count per byte. This instruction counts the number of bits that have a value of one in each vector element in the source SIMD&FP register, places the result into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VcntqU8N VcntqU8N
//go:noescape
func VcntqU8N(r *arm.Uint8, v0 *arm.Uint8, n int32)

// Join two smaller vectors into a single larger vector
//
//go:linkname VcombineS8N VcombineS8N
//go:noescape
func VcombineS8N(r *arm.Int8, v0 *arm.Int8, v1 *arm.Int8, n int32)

// Join two smaller vectors into a single larger vector
//
//go:linkname VcombineS16N VcombineS16N
//go:noescape
func VcombineS16N(r *arm.Int16, v0 *arm.Int16, v1 *arm.Int16, n int32)

// Join two smaller vectors into a single larger vector
//
//go:linkname VcombineS32N VcombineS32N
//go:noescape
func VcombineS32N(r *arm.Int32, v0 *arm.Int32, v1 *arm.Int32, n int32)

// Join two smaller vectors into a single larger vector
//
//go:linkname VcombineS64N VcombineS64N
//go:noescape
func VcombineS64N(r *arm.Int64, v0 *arm.Int64, v1 *arm.Int64, n int32)

// Join two smaller vectors into a single larger vector
//
//go:linkname VcombineU8N VcombineU8N
//go:noescape
func VcombineU8N(r *arm.Uint8, v0 *arm.Uint8, v1 *arm.Uint8, n int32)

// Join two smaller vectors into a single larger vector
//
//go:linkname VcombineU16N VcombineU16N
//go:noescape
func VcombineU16N(r *arm.Uint16, v0 *arm.Uint16, v1 *arm.Uint16, n int32)

// Join two smaller vectors into a single larger vector
//
//go:linkname VcombineU32N VcombineU32N
//go:noescape
func VcombineU32N(r *arm.Uint32, v0 *arm.Uint32, v1 *arm.Uint32, n int32)

// Join two smaller vectors into a single larger vector
//
//go:linkname VcombineU64N VcombineU64N
//go:noescape
func VcombineU64N(r *arm.Uint64, v0 *arm.Uint64, v1 *arm.Uint64, n int32)

// Join two smaller vectors into a single larger vector
//
//go:linkname VcombineF32N VcombineF32N
//go:noescape
func VcombineF32N(r *arm.Float32, v0 *arm.Float32, v1 *arm.Float32, n int32)

// Join two smaller vectors into a single larger vector
//
//go:linkname VcombineF64N VcombineF64N
//go:noescape
func VcombineF64N(r *arm.Float64, v0 *arm.Float64, v1 *arm.Float64, n int32)

// Signed fixed-point Convert to Floating-point (vector). This instruction converts each element in a vector from fixed-point to floating-point using the rounding mode that is specified by the FPCR, and writes the result to the SIMD&FP destination register.
//
//go:linkname VcvtF32S32N VcvtF32S32N
//go:noescape
func VcvtF32S32N(r *arm.Float32, v0 *arm.Int32, n int32)

// Unsigned fixed-point Convert to Floating-point (vector). This instruction converts each element in a vector from fixed-point to floating-point using the rounding mode that is specified by the FPCR, and writes the result to the SIMD&FP destination register.
//
//go:linkname VcvtF32U32N VcvtF32U32N
//go:noescape
func VcvtF32U32N(r *arm.Float32, v0 *arm.Uint32, n int32)

// Signed fixed-point Convert to Floating-point (vector). This instruction converts each element in a vector from fixed-point to floating-point using the rounding mode that is specified by the FPCR, and writes the result to the SIMD&FP destination register.
//
//go:linkname VcvtF64S64N VcvtF64S64N
//go:noescape
func VcvtF64S64N(r *arm.Float64, v0 *arm.Int64, n int32)

// Unsigned fixed-point Convert to Floating-point (vector). This instruction converts each element in a vector from fixed-point to floating-point using the rounding mode that is specified by the FPCR, and writes the result to the SIMD&FP destination register.
//
//go:linkname VcvtF64U64N VcvtF64U64N
//go:noescape
func VcvtF64U64N(r *arm.Float64, v0 *arm.Uint64, n int32)

// Floating-point Convert to Signed fixed-point, rounding toward Zero (vector). This instruction converts a scalar or each element in a vector from floating-point to fixed-point signed integer using the Round towards Zero rounding mode, and writes the result to the SIMD&FP destination register.
//
//go:linkname VcvtS32F32N VcvtS32F32N
//go:noescape
func VcvtS32F32N(r *arm.Int32, v0 *arm.Float32, n int32)

// Floating-point Convert to Signed fixed-point, rounding toward Zero (vector). This instruction converts a scalar or each element in a vector from floating-point to fixed-point signed integer using the Round towards Zero rounding mode, and writes the result to the SIMD&FP destination register.
//
//go:linkname VcvtS64F64N VcvtS64F64N
//go:noescape
func VcvtS64F64N(r *arm.Int64, v0 *arm.Float64, n int32)

// Floating-point Convert to Unsigned fixed-point, rounding toward Zero (vector). This instruction converts a scalar or each element in a vector from floating-point to fixed-point unsigned integer using the Round towards Zero rounding mode, and writes the result to the general-purpose destination register.
//
//go:linkname VcvtU32F32N VcvtU32F32N
//go:noescape
func VcvtU32F32N(r *arm.Uint32, v0 *arm.Float32, n int32)

// Floating-point Convert to Unsigned fixed-point, rounding toward Zero (vector). This instruction converts a scalar or each element in a vector from floating-point to fixed-point unsigned integer using the Round towards Zero rounding mode, and writes the result to the general-purpose destination register.
//
//go:linkname VcvtU64F64N VcvtU64F64N
//go:noescape
func VcvtU64F64N(r *arm.Uint64, v0 *arm.Float64, n int32)

// Floating-point Convert to Signed integer, rounding to nearest with ties to Away (vector). This instruction converts each element in a vector from a floating-point value to a signed integer value using the Round to Nearest with Ties to Away rounding mode and writes the result to the SIMD&FP destination register.
//
//go:linkname VcvtaS32F32N VcvtaS32F32N
//go:noescape
func VcvtaS32F32N(r *arm.Int32, v0 *arm.Float32, n int32)

// Floating-point Convert to Signed integer, rounding to nearest with ties to Away (vector). This instruction converts each element in a vector from a floating-point value to a signed integer value using the Round to Nearest with Ties to Away rounding mode and writes the result to the SIMD&FP destination register.
//
//go:linkname VcvtaS64F64N VcvtaS64F64N
//go:noescape
func VcvtaS64F64N(r *arm.Int64, v0 *arm.Float64, n int32)

// Floating-point Convert to Unsigned integer, rounding to nearest with ties to Away (vector). This instruction converts each element in a vector from a floating-point value to an unsigned integer value using the Round to Nearest with Ties to Away rounding mode and writes the result to the SIMD&FP destination register.
//
//go:linkname VcvtaU32F32N VcvtaU32F32N
//go:noescape
func VcvtaU32F32N(r *arm.Uint32, v0 *arm.Float32, n int32)

// Floating-point Convert to Unsigned integer, rounding to nearest with ties to Away (vector). This instruction converts each element in a vector from a floating-point value to an unsigned integer value using the Round to Nearest with Ties to Away rounding mode and writes the result to the SIMD&FP destination register.
//
//go:linkname VcvtaU64F64N VcvtaU64F64N
//go:noescape
func VcvtaU64F64N(r *arm.Uint64, v0 *arm.Float64, n int32)

// Floating-point Convert to Signed integer, rounding to nearest with ties to Away (vector). This instruction converts each element in a vector from a floating-point value to a signed integer value using the Round to Nearest with Ties to Away rounding mode and writes the result to the SIMD&FP destination register.
//
//go:linkname VcvtadS64F64N VcvtadS64F64N
//go:noescape
func VcvtadS64F64N(r *arm.Int64, v0 *arm.Float64, n int32)

// Floating-point Convert to Unsigned integer, rounding to nearest with ties to Away (vector). This instruction converts each element in a vector from a floating-point value to an unsigned integer value using the Round to Nearest with Ties to Away rounding mode and writes the result to the SIMD&FP destination register.
//
//go:linkname VcvtadU64F64N VcvtadU64F64N
//go:noescape
func VcvtadU64F64N(r *arm.Uint64, v0 *arm.Float64, n int32)

// Floating-point Convert to Signed integer, rounding to nearest with ties to Away (vector). This instruction converts each element in a vector from a floating-point value to a signed integer value using the Round to Nearest with Ties to Away rounding mode and writes the result to the SIMD&FP destination register.
//
//go:linkname VcvtaqS32F32N VcvtaqS32F32N
//go:noescape
func VcvtaqS32F32N(r *arm.Int32, v0 *arm.Float32, n int32)

// Floating-point Convert to Signed integer, rounding to nearest with ties to Away (vector). This instruction converts each element in a vector from a floating-point value to a signed integer value using the Round to Nearest with Ties to Away rounding mode and writes the result to the SIMD&FP destination register.
//
//go:linkname VcvtaqS64F64N VcvtaqS64F64N
//go:noescape
func VcvtaqS64F64N(r *arm.Int64, v0 *arm.Float64, n int32)

// Floating-point Convert to Unsigned integer, rounding to nearest with ties to Away (vector). This instruction converts each element in a vector from a floating-point value to an unsigned integer value using the Round to Nearest with Ties to Away rounding mode and writes the result to the SIMD&FP destination register.
//
//go:linkname VcvtaqU32F32N VcvtaqU32F32N
//go:noescape
func VcvtaqU32F32N(r *arm.Uint32, v0 *arm.Float32, n int32)

// Floating-point Convert to Unsigned integer, rounding to nearest with ties to Away (vector). This instruction converts each element in a vector from a floating-point value to an unsigned integer value using the Round to Nearest with Ties to Away rounding mode and writes the result to the SIMD&FP destination register.
//
//go:linkname VcvtaqU64F64N VcvtaqU64F64N
//go:noescape
func VcvtaqU64F64N(r *arm.Uint64, v0 *arm.Float64, n int32)

// Floating-point Convert to Signed integer, rounding to nearest with ties to Away (vector). This instruction converts each element in a vector from a floating-point value to a signed integer value using the Round to Nearest with Ties to Away rounding mode and writes the result to the SIMD&FP destination register.
//
//go:linkname VcvtasS32F32N VcvtasS32F32N
//go:noescape
func VcvtasS32F32N(r *arm.Int32, v0 *arm.Float32, n int32)

// Floating-point Convert to Unsigned integer, rounding to nearest with ties to Away (vector). This instruction converts each element in a vector from a floating-point value to an unsigned integer value using the Round to Nearest with Ties to Away rounding mode and writes the result to the SIMD&FP destination register.
//
//go:linkname VcvtasU32F32N VcvtasU32F32N
//go:noescape
func VcvtasU32F32N(r *arm.Uint32, v0 *arm.Float32, n int32)

// Signed fixed-point Convert to Floating-point (vector). This instruction converts each element in a vector from fixed-point to floating-point using the rounding mode that is specified by the FPCR, and writes the result to the SIMD&FP destination register.
//
//go:linkname VcvtdF64S64N VcvtdF64S64N
//go:noescape
func VcvtdF64S64N(r *arm.Float64, v0 *arm.Int64, n int32)

// Unsigned fixed-point Convert to Floating-point (vector). This instruction converts each element in a vector from fixed-point to floating-point using the rounding mode that is specified by the FPCR, and writes the result to the SIMD&FP destination register.
//
//go:linkname VcvtdF64U64N VcvtdF64U64N
//go:noescape
func VcvtdF64U64N(r *arm.Float64, v0 *arm.Uint64, n int32)

// Floating-point Convert to Signed fixed-point, rounding toward Zero (vector). This instruction converts a scalar or each element in a vector from floating-point to fixed-point signed integer using the Round towards Zero rounding mode, and writes the result to the SIMD&FP destination register.
//
//go:linkname VcvtdS64F64N VcvtdS64F64N
//go:noescape
func VcvtdS64F64N(r *arm.Int64, v0 *arm.Float64, n int32)

// Floating-point Convert to Unsigned fixed-point, rounding toward Zero (vector). This instruction converts a scalar or each element in a vector from floating-point to fixed-point unsigned integer using the Round towards Zero rounding mode, and writes the result to the general-purpose destination register.
//
//go:linkname VcvtdU64F64N VcvtdU64F64N
//go:noescape
func VcvtdU64F64N(r *arm.Uint64, v0 *arm.Float64, n int32)

// Floating-point Convert to Signed integer, rounding toward Minus infinity (vector). This instruction converts a scalar or each element in a vector from a floating-point value to a signed integer value using the Round towards Minus Infinity rounding mode, and writes the result to the SIMD&FP destination register.
//
//go:linkname VcvtmS32F32N VcvtmS32F32N
//go:noescape
func VcvtmS32F32N(r *arm.Int32, v0 *arm.Float32, n int32)

// Floating-point Convert to Signed integer, rounding toward Minus infinity (vector). This instruction converts a scalar or each element in a vector from a floating-point value to a signed integer value using the Round towards Minus Infinity rounding mode, and writes the result to the SIMD&FP destination register.
//
//go:linkname VcvtmS64F64N VcvtmS64F64N
//go:noescape
func VcvtmS64F64N(r *arm.Int64, v0 *arm.Float64, n int32)

// Floating-point Convert to Unsigned integer, rounding toward Minus infinity (vector). This instruction converts a scalar or each element in a vector from a floating-point value to an unsigned integer value using the Round towards Minus Infinity rounding mode, and writes the result to the SIMD&FP destination register.
//
//go:linkname VcvtmU32F32N VcvtmU32F32N
//go:noescape
func VcvtmU32F32N(r *arm.Uint32, v0 *arm.Float32, n int32)

// Floating-point Convert to Unsigned integer, rounding toward Minus infinity (vector). This instruction converts a scalar or each element in a vector from a floating-point value to an unsigned integer value using the Round towards Minus Infinity rounding mode, and writes the result to the SIMD&FP destination register.
//
//go:linkname VcvtmU64F64N VcvtmU64F64N
//go:noescape
func VcvtmU64F64N(r *arm.Uint64, v0 *arm.Float64, n int32)

// Floating-point Convert to Signed integer, rounding toward Minus infinity (vector). This instruction converts a scalar or each element in a vector from a floating-point value to a signed integer value using the Round towards Minus Infinity rounding mode, and writes the result to the SIMD&FP destination register.
//
//go:linkname VcvtmdS64F64N VcvtmdS64F64N
//go:noescape
func VcvtmdS64F64N(r *arm.Int64, v0 *arm.Float64, n int32)

// Floating-point Convert to Unsigned integer, rounding toward Minus infinity (vector). This instruction converts a scalar or each element in a vector from a floating-point value to an unsigned integer value using the Round towards Minus Infinity rounding mode, and writes the result to the SIMD&FP destination register.
//
//go:linkname VcvtmdU64F64N VcvtmdU64F64N
//go:noescape
func VcvtmdU64F64N(r *arm.Uint64, v0 *arm.Float64, n int32)

// Floating-point Convert to Signed integer, rounding toward Minus infinity (vector). This instruction converts a scalar or each element in a vector from a floating-point value to a signed integer value using the Round towards Minus Infinity rounding mode, and writes the result to the SIMD&FP destination register.
//
//go:linkname VcvtmqS32F32N VcvtmqS32F32N
//go:noescape
func VcvtmqS32F32N(r *arm.Int32, v0 *arm.Float32, n int32)

// Floating-point Convert to Signed integer, rounding toward Minus infinity (vector). This instruction converts a scalar or each element in a vector from a floating-point value to a signed integer value using the Round towards Minus Infinity rounding mode, and writes the result to the SIMD&FP destination register.
//
//go:linkname VcvtmqS64F64N VcvtmqS64F64N
//go:noescape
func VcvtmqS64F64N(r *arm.Int64, v0 *arm.Float64, n int32)

// Floating-point Convert to Unsigned integer, rounding toward Minus infinity (vector). This instruction converts a scalar or each element in a vector from a floating-point value to an unsigned integer value using the Round towards Minus Infinity rounding mode, and writes the result to the SIMD&FP destination register.
//
//go:linkname VcvtmqU32F32N VcvtmqU32F32N
//go:noescape
func VcvtmqU32F32N(r *arm.Uint32, v0 *arm.Float32, n int32)

// Floating-point Convert to Unsigned integer, rounding toward Minus infinity (vector). This instruction converts a scalar or each element in a vector from a floating-point value to an unsigned integer value using the Round towards Minus Infinity rounding mode, and writes the result to the SIMD&FP destination register.
//
//go:linkname VcvtmqU64F64N VcvtmqU64F64N
//go:noescape
func VcvtmqU64F64N(r *arm.Uint64, v0 *arm.Float64, n int32)

// Floating-point Convert to Signed integer, rounding toward Minus infinity (vector). This instruction converts a scalar or each element in a vector from a floating-point value to a signed integer value using the Round towards Minus Infinity rounding mode, and writes the result to the SIMD&FP destination register.
//
//go:linkname VcvtmsS32F32N VcvtmsS32F32N
//go:noescape
func VcvtmsS32F32N(r *arm.Int32, v0 *arm.Float32, n int32)

// Floating-point Convert to Unsigned integer, rounding toward Minus infinity (vector). This instruction converts a scalar or each element in a vector from a floating-point value to an unsigned integer value using the Round towards Minus Infinity rounding mode, and writes the result to the SIMD&FP destination register.
//
//go:linkname VcvtmsU32F32N VcvtmsU32F32N
//go:noescape
func VcvtmsU32F32N(r *arm.Uint32, v0 *arm.Float32, n int32)

// Floating-point Convert to Signed integer, rounding to nearest with ties to even (vector). This instruction converts a scalar or each element in a vector from a floating-point value to a signed integer value using the Round to Nearest rounding mode, and writes the result to the SIMD&FP destination register.
//
//go:linkname VcvtnS32F32N VcvtnS32F32N
//go:noescape
func VcvtnS32F32N(r *arm.Int32, v0 *arm.Float32, n int32)

// Floating-point Convert to Signed integer, rounding to nearest with ties to even (vector). This instruction converts a scalar or each element in a vector from a floating-point value to a signed integer value using the Round to Nearest rounding mode, and writes the result to the SIMD&FP destination register.
//
//go:linkname VcvtnS64F64N VcvtnS64F64N
//go:noescape
func VcvtnS64F64N(r *arm.Int64, v0 *arm.Float64, n int32)

// Floating-point Convert to Unsigned integer, rounding to nearest with ties to even (vector). This instruction converts a scalar or each element in a vector from a floating-point value to an unsigned integer value using the Round to Nearest rounding mode, and writes the result to the SIMD&FP destination register.
//
//go:linkname VcvtnU32F32N VcvtnU32F32N
//go:noescape
func VcvtnU32F32N(r *arm.Uint32, v0 *arm.Float32, n int32)

// Floating-point Convert to Unsigned integer, rounding to nearest with ties to even (vector). This instruction converts a scalar or each element in a vector from a floating-point value to an unsigned integer value using the Round to Nearest rounding mode, and writes the result to the SIMD&FP destination register.
//
//go:linkname VcvtnU64F64N VcvtnU64F64N
//go:noescape
func VcvtnU64F64N(r *arm.Uint64, v0 *arm.Float64, n int32)

// Floating-point Convert to Signed integer, rounding to nearest with ties to even (vector). This instruction converts a scalar or each element in a vector from a floating-point value to a signed integer value using the Round to Nearest rounding mode, and writes the result to the SIMD&FP destination register.
//
//go:linkname VcvtndS64F64N VcvtndS64F64N
//go:noescape
func VcvtndS64F64N(r *arm.Int64, v0 *arm.Float64, n int32)

// Floating-point Convert to Unsigned integer, rounding to nearest with ties to even (vector). This instruction converts a scalar or each element in a vector from a floating-point value to an unsigned integer value using the Round to Nearest rounding mode, and writes the result to the SIMD&FP destination register.
//
//go:linkname VcvtndU64F64N VcvtndU64F64N
//go:noescape
func VcvtndU64F64N(r *arm.Uint64, v0 *arm.Float64, n int32)

// Floating-point Convert to Signed integer, rounding to nearest with ties to even (vector). This instruction converts a scalar or each element in a vector from a floating-point value to a signed integer value using the Round to Nearest rounding mode, and writes the result to the SIMD&FP destination register.
//
//go:linkname VcvtnqS32F32N VcvtnqS32F32N
//go:noescape
func VcvtnqS32F32N(r *arm.Int32, v0 *arm.Float32, n int32)

// Floating-point Convert to Signed integer, rounding to nearest with ties to even (vector). This instruction converts a scalar or each element in a vector from a floating-point value to a signed integer value using the Round to Nearest rounding mode, and writes the result to the SIMD&FP destination register.
//
//go:linkname VcvtnqS64F64N VcvtnqS64F64N
//go:noescape
func VcvtnqS64F64N(r *arm.Int64, v0 *arm.Float64, n int32)

// Floating-point Convert to Unsigned integer, rounding to nearest with ties to even (vector). This instruction converts a scalar or each element in a vector from a floating-point value to an unsigned integer value using the Round to Nearest rounding mode, and writes the result to the SIMD&FP destination register.
//
//go:linkname VcvtnqU32F32N VcvtnqU32F32N
//go:noescape
func VcvtnqU32F32N(r *arm.Uint32, v0 *arm.Float32, n int32)

// Floating-point Convert to Unsigned integer, rounding to nearest with ties to even (vector). This instruction converts a scalar or each element in a vector from a floating-point value to an unsigned integer value using the Round to Nearest rounding mode, and writes the result to the SIMD&FP destination register.
//
//go:linkname VcvtnqU64F64N VcvtnqU64F64N
//go:noescape
func VcvtnqU64F64N(r *arm.Uint64, v0 *arm.Float64, n int32)

// Floating-point Convert to Signed integer, rounding to nearest with ties to even (vector). This instruction converts a scalar or each element in a vector from a floating-point value to a signed integer value using the Round to Nearest rounding mode, and writes the result to the SIMD&FP destination register.
//
//go:linkname VcvtnsS32F32N VcvtnsS32F32N
//go:noescape
func VcvtnsS32F32N(r *arm.Int32, v0 *arm.Float32, n int32)

// Floating-point Convert to Unsigned integer, rounding to nearest with ties to even (vector). This instruction converts a scalar or each element in a vector from a floating-point value to an unsigned integer value using the Round to Nearest rounding mode, and writes the result to the SIMD&FP destination register.
//
//go:linkname VcvtnsU32F32N VcvtnsU32F32N
//go:noescape
func VcvtnsU32F32N(r *arm.Uint32, v0 *arm.Float32, n int32)

// Floating-point Convert to Signed integer, rounding toward Plus infinity (vector). This instruction converts a scalar or each element in a vector from a floating-point value to a signed integer value using the Round towards Plus Infinity rounding mode, and writes the result to the SIMD&FP destination register.
//
//go:linkname VcvtpS32F32N VcvtpS32F32N
//go:noescape
func VcvtpS32F32N(r *arm.Int32, v0 *arm.Float32, n int32)

// Floating-point Convert to Signed integer, rounding toward Plus infinity (vector). This instruction converts a scalar or each element in a vector from a floating-point value to a signed integer value using the Round towards Plus Infinity rounding mode, and writes the result to the SIMD&FP destination register.
//
//go:linkname VcvtpS64F64N VcvtpS64F64N
//go:noescape
func VcvtpS64F64N(r *arm.Int64, v0 *arm.Float64, n int32)

// Floating-point Convert to Unsigned integer, rounding toward Plus infinity (vector). This instruction converts a scalar or each element in a vector from a floating-point value to an unsigned integer value using the Round towards Plus Infinity rounding mode, and writes the result to the SIMD&FP destination register.
//
//go:linkname VcvtpU32F32N VcvtpU32F32N
//go:noescape
func VcvtpU32F32N(r *arm.Uint32, v0 *arm.Float32, n int32)

// Floating-point Convert to Unsigned integer, rounding toward Plus infinity (vector). This instruction converts a scalar or each element in a vector from a floating-point value to an unsigned integer value using the Round towards Plus Infinity rounding mode, and writes the result to the SIMD&FP destination register.
//
//go:linkname VcvtpU64F64N VcvtpU64F64N
//go:noescape
func VcvtpU64F64N(r *arm.Uint64, v0 *arm.Float64, n int32)

// Floating-point Convert to Signed integer, rounding toward Plus infinity (vector). This instruction converts a scalar or each element in a vector from a floating-point value to a signed integer value using the Round towards Plus Infinity rounding mode, and writes the result to the SIMD&FP destination register.
//
//go:linkname VcvtpdS64F64N VcvtpdS64F64N
//go:noescape
func VcvtpdS64F64N(r *arm.Int64, v0 *arm.Float64, n int32)

// Floating-point Convert to Unsigned integer, rounding toward Plus infinity (vector). This instruction converts a scalar or each element in a vector from a floating-point value to an unsigned integer value using the Round towards Plus Infinity rounding mode, and writes the result to the SIMD&FP destination register.
//
//go:linkname VcvtpdU64F64N VcvtpdU64F64N
//go:noescape
func VcvtpdU64F64N(r *arm.Uint64, v0 *arm.Float64, n int32)

// Floating-point Convert to Signed integer, rounding toward Plus infinity (vector). This instruction converts a scalar or each element in a vector from a floating-point value to a signed integer value using the Round towards Plus Infinity rounding mode, and writes the result to the SIMD&FP destination register.
//
//go:linkname VcvtpqS32F32N VcvtpqS32F32N
//go:noescape
func VcvtpqS32F32N(r *arm.Int32, v0 *arm.Float32, n int32)

// Floating-point Convert to Signed integer, rounding toward Plus infinity (vector). This instruction converts a scalar or each element in a vector from a floating-point value to a signed integer value using the Round towards Plus Infinity rounding mode, and writes the result to the SIMD&FP destination register.
//
//go:linkname VcvtpqS64F64N VcvtpqS64F64N
//go:noescape
func VcvtpqS64F64N(r *arm.Int64, v0 *arm.Float64, n int32)

// Floating-point Convert to Unsigned integer, rounding toward Plus infinity (vector). This instruction converts a scalar or each element in a vector from a floating-point value to an unsigned integer value using the Round towards Plus Infinity rounding mode, and writes the result to the SIMD&FP destination register.
//
//go:linkname VcvtpqU32F32N VcvtpqU32F32N
//go:noescape
func VcvtpqU32F32N(r *arm.Uint32, v0 *arm.Float32, n int32)

// Floating-point Convert to Unsigned integer, rounding toward Plus infinity (vector). This instruction converts a scalar or each element in a vector from a floating-point value to an unsigned integer value using the Round towards Plus Infinity rounding mode, and writes the result to the SIMD&FP destination register.
//
//go:linkname VcvtpqU64F64N VcvtpqU64F64N
//go:noescape
func VcvtpqU64F64N(r *arm.Uint64, v0 *arm.Float64, n int32)

// Floating-point Convert to Signed integer, rounding toward Plus infinity (vector). This instruction converts a scalar or each element in a vector from a floating-point value to a signed integer value using the Round towards Plus Infinity rounding mode, and writes the result to the SIMD&FP destination register.
//
//go:linkname VcvtpsS32F32N VcvtpsS32F32N
//go:noescape
func VcvtpsS32F32N(r *arm.Int32, v0 *arm.Float32, n int32)

// Floating-point Convert to Unsigned integer, rounding toward Plus infinity (vector). This instruction converts a scalar or each element in a vector from a floating-point value to an unsigned integer value using the Round towards Plus Infinity rounding mode, and writes the result to the SIMD&FP destination register.
//
//go:linkname VcvtpsU32F32N VcvtpsU32F32N
//go:noescape
func VcvtpsU32F32N(r *arm.Uint32, v0 *arm.Float32, n int32)

// Signed fixed-point Convert to Floating-point (vector). This instruction converts each element in a vector from fixed-point to floating-point using the rounding mode that is specified by the FPCR, and writes the result to the SIMD&FP destination register.
//
//go:linkname VcvtqF32S32N VcvtqF32S32N
//go:noescape
func VcvtqF32S32N(r *arm.Float32, v0 *arm.Int32, n int32)

// Unsigned fixed-point Convert to Floating-point (vector). This instruction converts each element in a vector from fixed-point to floating-point using the rounding mode that is specified by the FPCR, and writes the result to the SIMD&FP destination register.
//
//go:linkname VcvtqF32U32N VcvtqF32U32N
//go:noescape
func VcvtqF32U32N(r *arm.Float32, v0 *arm.Uint32, n int32)

// Signed fixed-point Convert to Floating-point (vector). This instruction converts each element in a vector from fixed-point to floating-point using the rounding mode that is specified by the FPCR, and writes the result to the SIMD&FP destination register.
//
//go:linkname VcvtqF64S64N VcvtqF64S64N
//go:noescape
func VcvtqF64S64N(r *arm.Float64, v0 *arm.Int64, n int32)

// Unsigned fixed-point Convert to Floating-point (vector). This instruction converts each element in a vector from fixed-point to floating-point using the rounding mode that is specified by the FPCR, and writes the result to the SIMD&FP destination register.
//
//go:linkname VcvtqF64U64N VcvtqF64U64N
//go:noescape
func VcvtqF64U64N(r *arm.Float64, v0 *arm.Uint64, n int32)

// Floating-point Convert to Signed fixed-point, rounding toward Zero (vector). This instruction converts a scalar or each element in a vector from floating-point to fixed-point signed integer using the Round towards Zero rounding mode, and writes the result to the SIMD&FP destination register.
//
//go:linkname VcvtqS32F32N VcvtqS32F32N
//go:noescape
func VcvtqS32F32N(r *arm.Int32, v0 *arm.Float32, n int32)

// Floating-point Convert to Signed fixed-point, rounding toward Zero (vector). This instruction converts a scalar or each element in a vector from floating-point to fixed-point signed integer using the Round towards Zero rounding mode, and writes the result to the SIMD&FP destination register.
//
//go:linkname VcvtqS64F64N VcvtqS64F64N
//go:noescape
func VcvtqS64F64N(r *arm.Int64, v0 *arm.Float64, n int32)

// Floating-point Convert to Unsigned fixed-point, rounding toward Zero (vector). This instruction converts a scalar or each element in a vector from floating-point to fixed-point unsigned integer using the Round towards Zero rounding mode, and writes the result to the general-purpose destination register.
//
//go:linkname VcvtqU32F32N VcvtqU32F32N
//go:noescape
func VcvtqU32F32N(r *arm.Uint32, v0 *arm.Float32, n int32)

// Floating-point Convert to Unsigned fixed-point, rounding toward Zero (vector). This instruction converts a scalar or each element in a vector from floating-point to fixed-point unsigned integer using the Round towards Zero rounding mode, and writes the result to the general-purpose destination register.
//
//go:linkname VcvtqU64F64N VcvtqU64F64N
//go:noescape
func VcvtqU64F64N(r *arm.Uint64, v0 *arm.Float64, n int32)

// Signed fixed-point Convert to Floating-point (vector). This instruction converts each element in a vector from fixed-point to floating-point using the rounding mode that is specified by the FPCR, and writes the result to the SIMD&FP destination register.
//
//go:linkname VcvtsF32S32N VcvtsF32S32N
//go:noescape
func VcvtsF32S32N(r *arm.Float32, v0 *arm.Int32, n int32)

// Unsigned fixed-point Convert to Floating-point (vector). This instruction converts each element in a vector from fixed-point to floating-point using the rounding mode that is specified by the FPCR, and writes the result to the SIMD&FP destination register.
//
//go:linkname VcvtsF32U32N VcvtsF32U32N
//go:noescape
func VcvtsF32U32N(r *arm.Float32, v0 *arm.Uint32, n int32)

// Floating-point Convert to Signed fixed-point, rounding toward Zero (vector). This instruction converts a scalar or each element in a vector from floating-point to fixed-point signed integer using the Round towards Zero rounding mode, and writes the result to the SIMD&FP destination register.
//
//go:linkname VcvtsS32F32N VcvtsS32F32N
//go:noescape
func VcvtsS32F32N(r *arm.Int32, v0 *arm.Float32, n int32)

// Floating-point Convert to Unsigned fixed-point, rounding toward Zero (vector). This instruction converts a scalar or each element in a vector from floating-point to fixed-point unsigned integer using the Round towards Zero rounding mode, and writes the result to the general-purpose destination register.
//
//go:linkname VcvtsU32F32N VcvtsU32F32N
//go:noescape
func VcvtsU32F32N(r *arm.Uint32, v0 *arm.Float32, n int32)

// Floating-point Divide (vector). This instruction divides the floating-point values in the elements in the first source SIMD&FP register, by the floating-point values in the corresponding elements in the second source SIMD&FP register, places the results in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VdivF32N VdivF32N
//go:noescape
func VdivF32N(r *arm.Float32, v0 *arm.Float32, v1 *arm.Float32, n int32)

// Floating-point Divide (vector). This instruction divides the floating-point values in the elements in the first source SIMD&FP register, by the floating-point values in the corresponding elements in the second source SIMD&FP register, places the results in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VdivF64N VdivF64N
//go:noescape
func VdivF64N(r *arm.Float64, v0 *arm.Float64, v1 *arm.Float64, n int32)

// Floating-point Divide (vector). This instruction divides the floating-point values in the elements in the first source SIMD&FP register, by the floating-point values in the corresponding elements in the second source SIMD&FP register, places the results in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VdivqF32N VdivqF32N
//go:noescape
func VdivqF32N(r *arm.Float32, v0 *arm.Float32, v1 *arm.Float32, n int32)

// Floating-point Divide (vector). This instruction divides the floating-point values in the elements in the first source SIMD&FP register, by the floating-point values in the corresponding elements in the second source SIMD&FP register, places the results in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VdivqF64N VdivqF64N
//go:noescape
func VdivqF64N(r *arm.Float64, v0 *arm.Float64, v1 *arm.Float64, n int32)

// Duplicate vector element to vector or scalar. This instruction duplicates the vector element at the specified element index in the source SIMD&FP register into a scalar or each element in a vector, and writes the result to the destination SIMD&FP register.
//
//go:linkname VdupNS8N VdupNS8N
//go:noescape
func VdupNS8N(r *arm.Int8, v0 *arm.Int8, n int32)

// Duplicate vector element to vector or scalar. This instruction duplicates the vector element at the specified element index in the source SIMD&FP register into a scalar or each element in a vector, and writes the result to the destination SIMD&FP register.
//
//go:linkname VdupNS16N VdupNS16N
//go:noescape
func VdupNS16N(r *arm.Int16, v0 *arm.Int16, n int32)

// Duplicate vector element to vector or scalar. This instruction duplicates the vector element at the specified element index in the source SIMD&FP register into a scalar or each element in a vector, and writes the result to the destination SIMD&FP register.
//
//go:linkname VdupNS32N VdupNS32N
//go:noescape
func VdupNS32N(r *arm.Int32, v0 *arm.Int32, n int32)

// Insert vector element from another vector element. This instruction copies the vector element of the source SIMD&FP register to the specified vector element of the destination SIMD&FP register.
//
//go:linkname VdupNS64N VdupNS64N
//go:noescape
func VdupNS64N(r *arm.Int64, v0 *arm.Int64, n int32)

// Duplicate vector element to vector or scalar. This instruction duplicates the vector element at the specified element index in the source SIMD&FP register into a scalar or each element in a vector, and writes the result to the destination SIMD&FP register.
//
//go:linkname VdupNU8N VdupNU8N
//go:noescape
func VdupNU8N(r *arm.Uint8, v0 *arm.Uint8, n int32)

// Duplicate vector element to vector or scalar. This instruction duplicates the vector element at the specified element index in the source SIMD&FP register into a scalar or each element in a vector, and writes the result to the destination SIMD&FP register.
//
//go:linkname VdupNU16N VdupNU16N
//go:noescape
func VdupNU16N(r *arm.Uint16, v0 *arm.Uint16, n int32)

// Duplicate vector element to vector or scalar. This instruction duplicates the vector element at the specified element index in the source SIMD&FP register into a scalar or each element in a vector, and writes the result to the destination SIMD&FP register.
//
//go:linkname VdupNU32N VdupNU32N
//go:noescape
func VdupNU32N(r *arm.Uint32, v0 *arm.Uint32, n int32)

// Insert vector element from another vector element. This instruction copies the vector element of the source SIMD&FP register to the specified vector element of the destination SIMD&FP register.
//
//go:linkname VdupNU64N VdupNU64N
//go:noescape
func VdupNU64N(r *arm.Uint64, v0 *arm.Uint64, n int32)

// Duplicate vector element to vector or scalar. This instruction duplicates the vector element at the specified element index in the source SIMD&FP register into a scalar or each element in a vector, and writes the result to the destination SIMD&FP register.
//
//go:linkname VdupNF32N VdupNF32N
//go:noescape
func VdupNF32N(r *arm.Float32, v0 *arm.Float32, n int32)

// Insert vector element from another vector element. This instruction copies the vector element of the source SIMD&FP register to the specified vector element of the destination SIMD&FP register.
//
//go:linkname VdupNF64N VdupNF64N
//go:noescape
func VdupNF64N(r *arm.Float64, v0 *arm.Float64, n int32)

// Duplicate vector element to vector or scalar. This instruction duplicates the vector element at the specified element index in the source SIMD&FP register into a scalar or each element in a vector, and writes the result to the destination SIMD&FP register.
//
//go:linkname VdupqNS8N VdupqNS8N
//go:noescape
func VdupqNS8N(r *arm.Int8, v0 *arm.Int8, n int32)

// Duplicate vector element to vector or scalar. This instruction duplicates the vector element at the specified element index in the source SIMD&FP register into a scalar or each element in a vector, and writes the result to the destination SIMD&FP register.
//
//go:linkname VdupqNS16N VdupqNS16N
//go:noescape
func VdupqNS16N(r *arm.Int16, v0 *arm.Int16, n int32)

// Duplicate vector element to vector or scalar. This instruction duplicates the vector element at the specified element index in the source SIMD&FP register into a scalar or each element in a vector, and writes the result to the destination SIMD&FP register.
//
//go:linkname VdupqNS32N VdupqNS32N
//go:noescape
func VdupqNS32N(r *arm.Int32, v0 *arm.Int32, n int32)

// Duplicate vector element to vector or scalar. This instruction duplicates the vector element at the specified element index in the source SIMD&FP register into a scalar or each element in a vector, and writes the result to the destination SIMD&FP register.
//
//go:linkname VdupqNS64N VdupqNS64N
//go:noescape
func VdupqNS64N(r *arm.Int64, v0 *arm.Int64, n int32)

// Duplicate vector element to vector or scalar. This instruction duplicates the vector element at the specified element index in the source SIMD&FP register into a scalar or each element in a vector, and writes the result to the destination SIMD&FP register.
//
//go:linkname VdupqNU8N VdupqNU8N
//go:noescape
func VdupqNU8N(r *arm.Uint8, v0 *arm.Uint8, n int32)

// Duplicate vector element to vector or scalar. This instruction duplicates the vector element at the specified element index in the source SIMD&FP register into a scalar or each element in a vector, and writes the result to the destination SIMD&FP register.
//
//go:linkname VdupqNU16N VdupqNU16N
//go:noescape
func VdupqNU16N(r *arm.Uint16, v0 *arm.Uint16, n int32)

// Duplicate vector element to vector or scalar. This instruction duplicates the vector element at the specified element index in the source SIMD&FP register into a scalar or each element in a vector, and writes the result to the destination SIMD&FP register.
//
//go:linkname VdupqNU32N VdupqNU32N
//go:noescape
func VdupqNU32N(r *arm.Uint32, v0 *arm.Uint32, n int32)

// Duplicate vector element to vector or scalar. This instruction duplicates the vector element at the specified element index in the source SIMD&FP register into a scalar or each element in a vector, and writes the result to the destination SIMD&FP register.
//
//go:linkname VdupqNU64N VdupqNU64N
//go:noescape
func VdupqNU64N(r *arm.Uint64, v0 *arm.Uint64, n int32)

// Duplicate vector element to vector or scalar. This instruction duplicates the vector element at the specified element index in the source SIMD&FP register into a scalar or each element in a vector, and writes the result to the destination SIMD&FP register.
//
//go:linkname VdupqNF32N VdupqNF32N
//go:noescape
func VdupqNF32N(r *arm.Float32, v0 *arm.Float32, n int32)

// Duplicate vector element to vector or scalar. This instruction duplicates the vector element at the specified element index in the source SIMD&FP register into a scalar or each element in a vector, and writes the result to the destination SIMD&FP register.
//
//go:linkname VdupqNF64N VdupqNF64N
//go:noescape
func VdupqNF64N(r *arm.Float64, v0 *arm.Float64, n int32)

// Bitwise Exclusive OR (vector). This instruction performs a bitwise Exclusive OR operation between the two source SIMD&FP registers, and places the result in the destination SIMD&FP register.
//
//go:linkname VeorS8N VeorS8N
//go:noescape
func VeorS8N(r *arm.Int8, v0 *arm.Int8, v1 *arm.Int8, n int32)

// Bitwise Exclusive OR (vector). This instruction performs a bitwise Exclusive OR operation between the two source SIMD&FP registers, and places the result in the destination SIMD&FP register.
//
//go:linkname VeorS16N VeorS16N
//go:noescape
func VeorS16N(r *arm.Int16, v0 *arm.Int16, v1 *arm.Int16, n int32)

// Bitwise Exclusive OR (vector). This instruction performs a bitwise Exclusive OR operation between the two source SIMD&FP registers, and places the result in the destination SIMD&FP register.
//
//go:linkname VeorS32N VeorS32N
//go:noescape
func VeorS32N(r *arm.Int32, v0 *arm.Int32, v1 *arm.Int32, n int32)

// Bitwise Exclusive OR (vector). This instruction performs a bitwise Exclusive OR operation between the two source SIMD&FP registers, and places the result in the destination SIMD&FP register.
//
//go:linkname VeorS64N VeorS64N
//go:noescape
func VeorS64N(r *arm.Int64, v0 *arm.Int64, v1 *arm.Int64, n int32)

// Bitwise Exclusive OR (vector). This instruction performs a bitwise Exclusive OR operation between the two source SIMD&FP registers, and places the result in the destination SIMD&FP register.
//
//go:linkname VeorU8N VeorU8N
//go:noescape
func VeorU8N(r *arm.Uint8, v0 *arm.Uint8, v1 *arm.Uint8, n int32)

// Bitwise Exclusive OR (vector). This instruction performs a bitwise Exclusive OR operation between the two source SIMD&FP registers, and places the result in the destination SIMD&FP register.
//
//go:linkname VeorU16N VeorU16N
//go:noescape
func VeorU16N(r *arm.Uint16, v0 *arm.Uint16, v1 *arm.Uint16, n int32)

// Bitwise Exclusive OR (vector). This instruction performs a bitwise Exclusive OR operation between the two source SIMD&FP registers, and places the result in the destination SIMD&FP register.
//
//go:linkname VeorU32N VeorU32N
//go:noescape
func VeorU32N(r *arm.Uint32, v0 *arm.Uint32, v1 *arm.Uint32, n int32)

// Bitwise Exclusive OR (vector). This instruction performs a bitwise Exclusive OR operation between the two source SIMD&FP registers, and places the result in the destination SIMD&FP register.
//
//go:linkname VeorU64N VeorU64N
//go:noescape
func VeorU64N(r *arm.Uint64, v0 *arm.Uint64, v1 *arm.Uint64, n int32)

// Bitwise Exclusive OR (vector). This instruction performs a bitwise Exclusive OR operation between the two source SIMD&FP registers, and places the result in the destination SIMD&FP register.
//
//go:linkname VeorqS8N VeorqS8N
//go:noescape
func VeorqS8N(r *arm.Int8, v0 *arm.Int8, v1 *arm.Int8, n int32)

// Bitwise Exclusive OR (vector). This instruction performs a bitwise Exclusive OR operation between the two source SIMD&FP registers, and places the result in the destination SIMD&FP register.
//
//go:linkname VeorqS16N VeorqS16N
//go:noescape
func VeorqS16N(r *arm.Int16, v0 *arm.Int16, v1 *arm.Int16, n int32)

// Bitwise Exclusive OR (vector). This instruction performs a bitwise Exclusive OR operation between the two source SIMD&FP registers, and places the result in the destination SIMD&FP register.
//
//go:linkname VeorqS32N VeorqS32N
//go:noescape
func VeorqS32N(r *arm.Int32, v0 *arm.Int32, v1 *arm.Int32, n int32)

// Bitwise Exclusive OR (vector). This instruction performs a bitwise Exclusive OR operation between the two source SIMD&FP registers, and places the result in the destination SIMD&FP register.
//
//go:linkname VeorqS64N VeorqS64N
//go:noescape
func VeorqS64N(r *arm.Int64, v0 *arm.Int64, v1 *arm.Int64, n int32)

// Bitwise Exclusive OR (vector). This instruction performs a bitwise Exclusive OR operation between the two source SIMD&FP registers, and places the result in the destination SIMD&FP register.
//
//go:linkname VeorqU8N VeorqU8N
//go:noescape
func VeorqU8N(r *arm.Uint8, v0 *arm.Uint8, v1 *arm.Uint8, n int32)

// Bitwise Exclusive OR (vector). This instruction performs a bitwise Exclusive OR operation between the two source SIMD&FP registers, and places the result in the destination SIMD&FP register.
//
//go:linkname VeorqU16N VeorqU16N
//go:noescape
func VeorqU16N(r *arm.Uint16, v0 *arm.Uint16, v1 *arm.Uint16, n int32)

// Bitwise Exclusive OR (vector). This instruction performs a bitwise Exclusive OR operation between the two source SIMD&FP registers, and places the result in the destination SIMD&FP register.
//
//go:linkname VeorqU32N VeorqU32N
//go:noescape
func VeorqU32N(r *arm.Uint32, v0 *arm.Uint32, v1 *arm.Uint32, n int32)

// Bitwise Exclusive OR (vector). This instruction performs a bitwise Exclusive OR operation between the two source SIMD&FP registers, and places the result in the destination SIMD&FP register.
//
//go:linkname VeorqU64N VeorqU64N
//go:noescape
func VeorqU64N(r *arm.Uint64, v0 *arm.Uint64, v1 *arm.Uint64, n int32)

// Duplicate vector element to vector or scalar. This instruction duplicates the vector element at the specified element index in the source SIMD&FP register into a scalar or each element in a vector, and writes the result to the destination SIMD&FP register.
//
//go:linkname VgetHighS8N VgetHighS8N
//go:noescape
func VgetHighS8N(r *arm.Int8, v0 *arm.Int8, n int32)

// Duplicate vector element to vector or scalar. This instruction duplicates the vector element at the specified element index in the source SIMD&FP register into a scalar or each element in a vector, and writes the result to the destination SIMD&FP register.
//
//go:linkname VgetHighS16N VgetHighS16N
//go:noescape
func VgetHighS16N(r *arm.Int16, v0 *arm.Int16, n int32)

// Duplicate vector element to vector or scalar. This instruction duplicates the vector element at the specified element index in the source SIMD&FP register into a scalar or each element in a vector, and writes the result to the destination SIMD&FP register.
//
//go:linkname VgetHighS32N VgetHighS32N
//go:noescape
func VgetHighS32N(r *arm.Int32, v0 *arm.Int32, n int32)

// Duplicate vector element to vector or scalar. This instruction duplicates the vector element at the specified element index in the source SIMD&FP register into a scalar or each element in a vector, and writes the result to the destination SIMD&FP register.
//
//go:linkname VgetHighS64N VgetHighS64N
//go:noescape
func VgetHighS64N(r *arm.Int64, v0 *arm.Int64, n int32)

// Duplicate vector element to vector or scalar. This instruction duplicates the vector element at the specified element index in the source SIMD&FP register into a scalar or each element in a vector, and writes the result to the destination SIMD&FP register.
//
//go:linkname VgetHighU8N VgetHighU8N
//go:noescape
func VgetHighU8N(r *arm.Uint8, v0 *arm.Uint8, n int32)

// Duplicate vector element to vector or scalar. This instruction duplicates the vector element at the specified element index in the source SIMD&FP register into a scalar or each element in a vector, and writes the result to the destination SIMD&FP register.
//
//go:linkname VgetHighU16N VgetHighU16N
//go:noescape
func VgetHighU16N(r *arm.Uint16, v0 *arm.Uint16, n int32)

// Duplicate vector element to vector or scalar. This instruction duplicates the vector element at the specified element index in the source SIMD&FP register into a scalar or each element in a vector, and writes the result to the destination SIMD&FP register.
//
//go:linkname VgetHighU32N VgetHighU32N
//go:noescape
func VgetHighU32N(r *arm.Uint32, v0 *arm.Uint32, n int32)

// Duplicate vector element to vector or scalar. This instruction duplicates the vector element at the specified element index in the source SIMD&FP register into a scalar or each element in a vector, and writes the result to the destination SIMD&FP register.
//
//go:linkname VgetHighU64N VgetHighU64N
//go:noescape
func VgetHighU64N(r *arm.Uint64, v0 *arm.Uint64, n int32)

// Duplicate vector element to vector or scalar. This instruction duplicates the vector element at the specified element index in the source SIMD&FP register into a scalar or each element in a vector, and writes the result to the destination SIMD&FP register.
//
//go:linkname VgetHighF32N VgetHighF32N
//go:noescape
func VgetHighF32N(r *arm.Float32, v0 *arm.Float32, n int32)

// Duplicate vector element to vector or scalar. This instruction duplicates the vector element at the specified element index in the source SIMD&FP register into a scalar or each element in a vector, and writes the result to the destination SIMD&FP register.
//
//go:linkname VgetHighF64N VgetHighF64N
//go:noescape
func VgetHighF64N(r *arm.Float64, v0 *arm.Float64, n int32)

// Duplicate vector element to vector or scalar. This instruction duplicates the vector element at the specified element index in the source SIMD&FP register into a scalar or each element in a vector, and writes the result to the destination SIMD&FP register.
//
//go:linkname VgetLowS8N VgetLowS8N
//go:noescape
func VgetLowS8N(r *arm.Int8, v0 *arm.Int8, n int32)

// Duplicate vector element to vector or scalar. This instruction duplicates the vector element at the specified element index in the source SIMD&FP register into a scalar or each element in a vector, and writes the result to the destination SIMD&FP register.
//
//go:linkname VgetLowS16N VgetLowS16N
//go:noescape
func VgetLowS16N(r *arm.Int16, v0 *arm.Int16, n int32)

// Duplicate vector element to vector or scalar. This instruction duplicates the vector element at the specified element index in the source SIMD&FP register into a scalar or each element in a vector, and writes the result to the destination SIMD&FP register.
//
//go:linkname VgetLowS32N VgetLowS32N
//go:noescape
func VgetLowS32N(r *arm.Int32, v0 *arm.Int32, n int32)

// Duplicate vector element to vector or scalar. This instruction duplicates the vector element at the specified element index in the source SIMD&FP register into a scalar or each element in a vector, and writes the result to the destination SIMD&FP register.
//
//go:linkname VgetLowS64N VgetLowS64N
//go:noescape
func VgetLowS64N(r *arm.Int64, v0 *arm.Int64, n int32)

// Duplicate vector element to vector or scalar. This instruction duplicates the vector element at the specified element index in the source SIMD&FP register into a scalar or each element in a vector, and writes the result to the destination SIMD&FP register.
//
//go:linkname VgetLowU8N VgetLowU8N
//go:noescape
func VgetLowU8N(r *arm.Uint8, v0 *arm.Uint8, n int32)

// Duplicate vector element to vector or scalar. This instruction duplicates the vector element at the specified element index in the source SIMD&FP register into a scalar or each element in a vector, and writes the result to the destination SIMD&FP register.
//
//go:linkname VgetLowU16N VgetLowU16N
//go:noescape
func VgetLowU16N(r *arm.Uint16, v0 *arm.Uint16, n int32)

// Duplicate vector element to vector or scalar. This instruction duplicates the vector element at the specified element index in the source SIMD&FP register into a scalar or each element in a vector, and writes the result to the destination SIMD&FP register.
//
//go:linkname VgetLowU32N VgetLowU32N
//go:noescape
func VgetLowU32N(r *arm.Uint32, v0 *arm.Uint32, n int32)

// Duplicate vector element to vector or scalar. This instruction duplicates the vector element at the specified element index in the source SIMD&FP register into a scalar or each element in a vector, and writes the result to the destination SIMD&FP register.
//
//go:linkname VgetLowU64N VgetLowU64N
//go:noescape
func VgetLowU64N(r *arm.Uint64, v0 *arm.Uint64, n int32)

// Duplicate vector element to vector or scalar. This instruction duplicates the vector element at the specified element index in the source SIMD&FP register into a scalar or each element in a vector, and writes the result to the destination SIMD&FP register.
//
//go:linkname VgetLowF32N VgetLowF32N
//go:noescape
func VgetLowF32N(r *arm.Float32, v0 *arm.Float32, n int32)

// Duplicate vector element to vector or scalar. This instruction duplicates the vector element at the specified element index in the source SIMD&FP register into a scalar or each element in a vector, and writes the result to the destination SIMD&FP register.
//
//go:linkname VgetLowF64N VgetLowF64N
//go:noescape
func VgetLowF64N(r *arm.Float64, v0 *arm.Float64, n int32)

// Signed Halving Add. This instruction adds corresponding signed integer values from the two source SIMD&FP registers, shifts each result right one bit, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VhaddS8N VhaddS8N
//go:noescape
func VhaddS8N(r *arm.Int8, v0 *arm.Int8, v1 *arm.Int8, n int32)

// Signed Halving Add. This instruction adds corresponding signed integer values from the two source SIMD&FP registers, shifts each result right one bit, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VhaddS16N VhaddS16N
//go:noescape
func VhaddS16N(r *arm.Int16, v0 *arm.Int16, v1 *arm.Int16, n int32)

// Signed Halving Add. This instruction adds corresponding signed integer values from the two source SIMD&FP registers, shifts each result right one bit, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VhaddS32N VhaddS32N
//go:noescape
func VhaddS32N(r *arm.Int32, v0 *arm.Int32, v1 *arm.Int32, n int32)

// Unsigned Halving Add. This instruction adds corresponding unsigned integer values from the two source SIMD&FP registers, shifts each result right one bit, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VhaddU8N VhaddU8N
//go:noescape
func VhaddU8N(r *arm.Uint8, v0 *arm.Uint8, v1 *arm.Uint8, n int32)

// Unsigned Halving Add. This instruction adds corresponding unsigned integer values from the two source SIMD&FP registers, shifts each result right one bit, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VhaddU16N VhaddU16N
//go:noescape
func VhaddU16N(r *arm.Uint16, v0 *arm.Uint16, v1 *arm.Uint16, n int32)

// Unsigned Halving Add. This instruction adds corresponding unsigned integer values from the two source SIMD&FP registers, shifts each result right one bit, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VhaddU32N VhaddU32N
//go:noescape
func VhaddU32N(r *arm.Uint32, v0 *arm.Uint32, v1 *arm.Uint32, n int32)

// Signed Halving Add. This instruction adds corresponding signed integer values from the two source SIMD&FP registers, shifts each result right one bit, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VhaddqS8N VhaddqS8N
//go:noescape
func VhaddqS8N(r *arm.Int8, v0 *arm.Int8, v1 *arm.Int8, n int32)

// Signed Halving Add. This instruction adds corresponding signed integer values from the two source SIMD&FP registers, shifts each result right one bit, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VhaddqS16N VhaddqS16N
//go:noescape
func VhaddqS16N(r *arm.Int16, v0 *arm.Int16, v1 *arm.Int16, n int32)

// Signed Halving Add. This instruction adds corresponding signed integer values from the two source SIMD&FP registers, shifts each result right one bit, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VhaddqS32N VhaddqS32N
//go:noescape
func VhaddqS32N(r *arm.Int32, v0 *arm.Int32, v1 *arm.Int32, n int32)

// Unsigned Halving Add. This instruction adds corresponding unsigned integer values from the two source SIMD&FP registers, shifts each result right one bit, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VhaddqU8N VhaddqU8N
//go:noescape
func VhaddqU8N(r *arm.Uint8, v0 *arm.Uint8, v1 *arm.Uint8, n int32)

// Unsigned Halving Add. This instruction adds corresponding unsigned integer values from the two source SIMD&FP registers, shifts each result right one bit, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VhaddqU16N VhaddqU16N
//go:noescape
func VhaddqU16N(r *arm.Uint16, v0 *arm.Uint16, v1 *arm.Uint16, n int32)

// Unsigned Halving Add. This instruction adds corresponding unsigned integer values from the two source SIMD&FP registers, shifts each result right one bit, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VhaddqU32N VhaddqU32N
//go:noescape
func VhaddqU32N(r *arm.Uint32, v0 *arm.Uint32, v1 *arm.Uint32, n int32)

// Signed Halving Subtract. This instruction subtracts the elements in the vector in the second source SIMD&FP register from the corresponding elements in the vector in the first source SIMD&FP register, shifts each result right one bit, places each result into elements of a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VhsubS8N VhsubS8N
//go:noescape
func VhsubS8N(r *arm.Int8, v0 *arm.Int8, v1 *arm.Int8, n int32)

// Signed Halving Subtract. This instruction subtracts the elements in the vector in the second source SIMD&FP register from the corresponding elements in the vector in the first source SIMD&FP register, shifts each result right one bit, places each result into elements of a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VhsubS16N VhsubS16N
//go:noescape
func VhsubS16N(r *arm.Int16, v0 *arm.Int16, v1 *arm.Int16, n int32)

// Signed Halving Subtract. This instruction subtracts the elements in the vector in the second source SIMD&FP register from the corresponding elements in the vector in the first source SIMD&FP register, shifts each result right one bit, places each result into elements of a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VhsubS32N VhsubS32N
//go:noescape
func VhsubS32N(r *arm.Int32, v0 *arm.Int32, v1 *arm.Int32, n int32)

// Unsigned Halving Subtract. This instruction subtracts the vector elements in the second source SIMD&FP register from the corresponding vector elements in the first source SIMD&FP register, shifts each result right one bit, places each result into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VhsubU8N VhsubU8N
//go:noescape
func VhsubU8N(r *arm.Uint8, v0 *arm.Uint8, v1 *arm.Uint8, n int32)

// Unsigned Halving Subtract. This instruction subtracts the vector elements in the second source SIMD&FP register from the corresponding vector elements in the first source SIMD&FP register, shifts each result right one bit, places each result into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VhsubU16N VhsubU16N
//go:noescape
func VhsubU16N(r *arm.Uint16, v0 *arm.Uint16, v1 *arm.Uint16, n int32)

// Unsigned Halving Subtract. This instruction subtracts the vector elements in the second source SIMD&FP register from the corresponding vector elements in the first source SIMD&FP register, shifts each result right one bit, places each result into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VhsubU32N VhsubU32N
//go:noescape
func VhsubU32N(r *arm.Uint32, v0 *arm.Uint32, v1 *arm.Uint32, n int32)

// Signed Halving Subtract. This instruction subtracts the elements in the vector in the second source SIMD&FP register from the corresponding elements in the vector in the first source SIMD&FP register, shifts each result right one bit, places each result into elements of a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VhsubqS8N VhsubqS8N
//go:noescape
func VhsubqS8N(r *arm.Int8, v0 *arm.Int8, v1 *arm.Int8, n int32)

// Signed Halving Subtract. This instruction subtracts the elements in the vector in the second source SIMD&FP register from the corresponding elements in the vector in the first source SIMD&FP register, shifts each result right one bit, places each result into elements of a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VhsubqS16N VhsubqS16N
//go:noescape
func VhsubqS16N(r *arm.Int16, v0 *arm.Int16, v1 *arm.Int16, n int32)

// Signed Halving Subtract. This instruction subtracts the elements in the vector in the second source SIMD&FP register from the corresponding elements in the vector in the first source SIMD&FP register, shifts each result right one bit, places each result into elements of a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VhsubqS32N VhsubqS32N
//go:noescape
func VhsubqS32N(r *arm.Int32, v0 *arm.Int32, v1 *arm.Int32, n int32)

// Unsigned Halving Subtract. This instruction subtracts the vector elements in the second source SIMD&FP register from the corresponding vector elements in the first source SIMD&FP register, shifts each result right one bit, places each result into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VhsubqU8N VhsubqU8N
//go:noescape
func VhsubqU8N(r *arm.Uint8, v0 *arm.Uint8, v1 *arm.Uint8, n int32)

// Unsigned Halving Subtract. This instruction subtracts the vector elements in the second source SIMD&FP register from the corresponding vector elements in the first source SIMD&FP register, shifts each result right one bit, places each result into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VhsubqU16N VhsubqU16N
//go:noescape
func VhsubqU16N(r *arm.Uint16, v0 *arm.Uint16, v1 *arm.Uint16, n int32)

// Unsigned Halving Subtract. This instruction subtracts the vector elements in the second source SIMD&FP register from the corresponding vector elements in the first source SIMD&FP register, shifts each result right one bit, places each result into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VhsubqU32N VhsubqU32N
//go:noescape
func VhsubqU32N(r *arm.Uint32, v0 *arm.Uint32, v1 *arm.Uint32, n int32)

// Signed Maximum (vector). This instruction compares corresponding elements in the vectors in the two source SIMD&FP registers, places the larger of each pair of signed integer values into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VmaxS8N VmaxS8N
//go:noescape
func VmaxS8N(r *arm.Int8, v0 *arm.Int8, v1 *arm.Int8, n int32)

// Signed Maximum (vector). This instruction compares corresponding elements in the vectors in the two source SIMD&FP registers, places the larger of each pair of signed integer values into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VmaxS16N VmaxS16N
//go:noescape
func VmaxS16N(r *arm.Int16, v0 *arm.Int16, v1 *arm.Int16, n int32)

// Signed Maximum (vector). This instruction compares corresponding elements in the vectors in the two source SIMD&FP registers, places the larger of each pair of signed integer values into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VmaxS32N VmaxS32N
//go:noescape
func VmaxS32N(r *arm.Int32, v0 *arm.Int32, v1 *arm.Int32, n int32)

// Unsigned Maximum (vector). This instruction compares corresponding elements in the vectors in the two source SIMD&FP registers, places the larger of each pair of unsigned integer values into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VmaxU8N VmaxU8N
//go:noescape
func VmaxU8N(r *arm.Uint8, v0 *arm.Uint8, v1 *arm.Uint8, n int32)

// Unsigned Maximum (vector). This instruction compares corresponding elements in the vectors in the two source SIMD&FP registers, places the larger of each pair of unsigned integer values into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VmaxU16N VmaxU16N
//go:noescape
func VmaxU16N(r *arm.Uint16, v0 *arm.Uint16, v1 *arm.Uint16, n int32)

// Unsigned Maximum (vector). This instruction compares corresponding elements in the vectors in the two source SIMD&FP registers, places the larger of each pair of unsigned integer values into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VmaxU32N VmaxU32N
//go:noescape
func VmaxU32N(r *arm.Uint32, v0 *arm.Uint32, v1 *arm.Uint32, n int32)

// Floating-point Maximum (vector). This instruction compares corresponding vector elements in the two source SIMD&FP registers, places the larger of each of the two floating-point values into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VmaxF32N VmaxF32N
//go:noescape
func VmaxF32N(r *arm.Float32, v0 *arm.Float32, v1 *arm.Float32, n int32)

// Floating-point Maximum (vector). This instruction compares corresponding vector elements in the two source SIMD&FP registers, places the larger of each of the two floating-point values into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VmaxF64N VmaxF64N
//go:noescape
func VmaxF64N(r *arm.Float64, v0 *arm.Float64, v1 *arm.Float64, n int32)

// Floating-point Maximum Number (vector). This instruction compares corresponding vector elements in the two source SIMD&FP registers, writes the larger of the two floating-point values into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VmaxnmF32N VmaxnmF32N
//go:noescape
func VmaxnmF32N(r *arm.Float32, v0 *arm.Float32, v1 *arm.Float32, n int32)

// Floating-point Maximum Number (vector). This instruction compares corresponding vector elements in the two source SIMD&FP registers, writes the larger of the two floating-point values into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VmaxnmF64N VmaxnmF64N
//go:noescape
func VmaxnmF64N(r *arm.Float64, v0 *arm.Float64, v1 *arm.Float64, n int32)

// Floating-point Maximum Number (vector). This instruction compares corresponding vector elements in the two source SIMD&FP registers, writes the larger of the two floating-point values into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VmaxnmqF32N VmaxnmqF32N
//go:noescape
func VmaxnmqF32N(r *arm.Float32, v0 *arm.Float32, v1 *arm.Float32, n int32)

// Floating-point Maximum Number (vector). This instruction compares corresponding vector elements in the two source SIMD&FP registers, writes the larger of the two floating-point values into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VmaxnmqF64N VmaxnmqF64N
//go:noescape
func VmaxnmqF64N(r *arm.Float64, v0 *arm.Float64, v1 *arm.Float64, n int32)

// Floating-point Maximum Number Pairwise (vector). This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements in the two source SIMD&FP registers, writes the largest of each pair of values into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are floating-point values.
//
//go:linkname VmaxnmvF32N VmaxnmvF32N
//go:noescape
func VmaxnmvF32N(r *arm.Float32, v0 *arm.Float32, n int32)

// Floating-point Maximum Number across Vector. This instruction compares all the vector elements in the source SIMD&FP register, and writes the largest of the values as a scalar to the destination SIMD&FP register. All the values in this instruction are floating-point values.
//
//go:linkname VmaxnmvqF32N VmaxnmvqF32N
//go:noescape
func VmaxnmvqF32N(r *arm.Float32, v0 *arm.Float32, n int32)

// Floating-point Maximum Number Pairwise (vector). This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements in the two source SIMD&FP registers, writes the largest of each pair of values into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are floating-point values.
//
//go:linkname VmaxnmvqF64N VmaxnmvqF64N
//go:noescape
func VmaxnmvqF64N(r *arm.Float64, v0 *arm.Float64, n int32)

// Signed Maximum (vector). This instruction compares corresponding elements in the vectors in the two source SIMD&FP registers, places the larger of each pair of signed integer values into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VmaxqS8N VmaxqS8N
//go:noescape
func VmaxqS8N(r *arm.Int8, v0 *arm.Int8, v1 *arm.Int8, n int32)

// Signed Maximum (vector). This instruction compares corresponding elements in the vectors in the two source SIMD&FP registers, places the larger of each pair of signed integer values into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VmaxqS16N VmaxqS16N
//go:noescape
func VmaxqS16N(r *arm.Int16, v0 *arm.Int16, v1 *arm.Int16, n int32)

// Signed Maximum (vector). This instruction compares corresponding elements in the vectors in the two source SIMD&FP registers, places the larger of each pair of signed integer values into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VmaxqS32N VmaxqS32N
//go:noescape
func VmaxqS32N(r *arm.Int32, v0 *arm.Int32, v1 *arm.Int32, n int32)

// Unsigned Maximum (vector). This instruction compares corresponding elements in the vectors in the two source SIMD&FP registers, places the larger of each pair of unsigned integer values into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VmaxqU8N VmaxqU8N
//go:noescape
func VmaxqU8N(r *arm.Uint8, v0 *arm.Uint8, v1 *arm.Uint8, n int32)

// Unsigned Maximum (vector). This instruction compares corresponding elements in the vectors in the two source SIMD&FP registers, places the larger of each pair of unsigned integer values into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VmaxqU16N VmaxqU16N
//go:noescape
func VmaxqU16N(r *arm.Uint16, v0 *arm.Uint16, v1 *arm.Uint16, n int32)

// Unsigned Maximum (vector). This instruction compares corresponding elements in the vectors in the two source SIMD&FP registers, places the larger of each pair of unsigned integer values into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VmaxqU32N VmaxqU32N
//go:noescape
func VmaxqU32N(r *arm.Uint32, v0 *arm.Uint32, v1 *arm.Uint32, n int32)

// Floating-point Maximum (vector). This instruction compares corresponding vector elements in the two source SIMD&FP registers, places the larger of each of the two floating-point values into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VmaxqF32N VmaxqF32N
//go:noescape
func VmaxqF32N(r *arm.Float32, v0 *arm.Float32, v1 *arm.Float32, n int32)

// Floating-point Maximum (vector). This instruction compares corresponding vector elements in the two source SIMD&FP registers, places the larger of each of the two floating-point values into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VmaxqF64N VmaxqF64N
//go:noescape
func VmaxqF64N(r *arm.Float64, v0 *arm.Float64, v1 *arm.Float64, n int32)

// Signed Maximum across Vector. This instruction compares all the vector elements in the source SIMD&FP register, and writes the largest of the values as a scalar to the destination SIMD&FP register. All the values in this instruction are signed integer values.
//
//go:linkname VmaxvS8N VmaxvS8N
//go:noescape
func VmaxvS8N(r *arm.Int8, v0 *arm.Int8, n int32)

// Signed Maximum across Vector. This instruction compares all the vector elements in the source SIMD&FP register, and writes the largest of the values as a scalar to the destination SIMD&FP register. All the values in this instruction are signed integer values.
//
//go:linkname VmaxvS16N VmaxvS16N
//go:noescape
func VmaxvS16N(r *arm.Int16, v0 *arm.Int16, n int32)

// Signed Maximum Pairwise. This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements in the two source SIMD&FP registers, writes the largest of each pair of signed integer values into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VmaxvS32N VmaxvS32N
//go:noescape
func VmaxvS32N(r *arm.Int32, v0 *arm.Int32, n int32)

// Unsigned Maximum across Vector. This instruction compares all the vector elements in the source SIMD&FP register, and writes the largest of the values as a scalar to the destination SIMD&FP register. All the values in this instruction are unsigned integer values.
//
//go:linkname VmaxvU8N VmaxvU8N
//go:noescape
func VmaxvU8N(r *arm.Uint8, v0 *arm.Uint8, n int32)

// Unsigned Maximum across Vector. This instruction compares all the vector elements in the source SIMD&FP register, and writes the largest of the values as a scalar to the destination SIMD&FP register. All the values in this instruction are unsigned integer values.
//
//go:linkname VmaxvU16N VmaxvU16N
//go:noescape
func VmaxvU16N(r *arm.Uint16, v0 *arm.Uint16, n int32)

// Unsigned Maximum Pairwise. This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements in the two source SIMD&FP registers, writes the largest of each pair of unsigned integer values into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VmaxvU32N VmaxvU32N
//go:noescape
func VmaxvU32N(r *arm.Uint32, v0 *arm.Uint32, n int32)

// Floating-point Maximum Pairwise (vector). This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements from the concatenated vector, writes the larger of each pair of values into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are floating-point values.
//
//go:linkname VmaxvF32N VmaxvF32N
//go:noescape
func VmaxvF32N(r *arm.Float32, v0 *arm.Float32, n int32)

// Signed Maximum across Vector. This instruction compares all the vector elements in the source SIMD&FP register, and writes the largest of the values as a scalar to the destination SIMD&FP register. All the values in this instruction are signed integer values.
//
//go:linkname VmaxvqS8N VmaxvqS8N
//go:noescape
func VmaxvqS8N(r *arm.Int8, v0 *arm.Int8, n int32)

// Signed Maximum across Vector. This instruction compares all the vector elements in the source SIMD&FP register, and writes the largest of the values as a scalar to the destination SIMD&FP register. All the values in this instruction are signed integer values.
//
//go:linkname VmaxvqS16N VmaxvqS16N
//go:noescape
func VmaxvqS16N(r *arm.Int16, v0 *arm.Int16, n int32)

// Signed Maximum across Vector. This instruction compares all the vector elements in the source SIMD&FP register, and writes the largest of the values as a scalar to the destination SIMD&FP register. All the values in this instruction are signed integer values.
//
//go:linkname VmaxvqS32N VmaxvqS32N
//go:noescape
func VmaxvqS32N(r *arm.Int32, v0 *arm.Int32, n int32)

// Unsigned Maximum across Vector. This instruction compares all the vector elements in the source SIMD&FP register, and writes the largest of the values as a scalar to the destination SIMD&FP register. All the values in this instruction are unsigned integer values.
//
//go:linkname VmaxvqU8N VmaxvqU8N
//go:noescape
func VmaxvqU8N(r *arm.Uint8, v0 *arm.Uint8, n int32)

// Unsigned Maximum across Vector. This instruction compares all the vector elements in the source SIMD&FP register, and writes the largest of the values as a scalar to the destination SIMD&FP register. All the values in this instruction are unsigned integer values.
//
//go:linkname VmaxvqU16N VmaxvqU16N
//go:noescape
func VmaxvqU16N(r *arm.Uint16, v0 *arm.Uint16, n int32)

// Unsigned Maximum across Vector. This instruction compares all the vector elements in the source SIMD&FP register, and writes the largest of the values as a scalar to the destination SIMD&FP register. All the values in this instruction are unsigned integer values.
//
//go:linkname VmaxvqU32N VmaxvqU32N
//go:noescape
func VmaxvqU32N(r *arm.Uint32, v0 *arm.Uint32, n int32)

// Floating-point Maximum across Vector. This instruction compares all the vector elements in the source SIMD&FP register, and writes the largest of the values as a scalar to the destination SIMD&FP register. All the values in this instruction are floating-point values.
//
//go:linkname VmaxvqF32N VmaxvqF32N
//go:noescape
func VmaxvqF32N(r *arm.Float32, v0 *arm.Float32, n int32)

// Floating-point Maximum Pairwise (vector). This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements from the concatenated vector, writes the larger of each pair of values into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are floating-point values.
//
//go:linkname VmaxvqF64N VmaxvqF64N
//go:noescape
func VmaxvqF64N(r *arm.Float64, v0 *arm.Float64, n int32)

// Signed Minimum (vector). This instruction compares corresponding elements in the vectors in the two source SIMD&FP registers, places the smaller of each of the two signed integer values into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VminS8N VminS8N
//go:noescape
func VminS8N(r *arm.Int8, v0 *arm.Int8, v1 *arm.Int8, n int32)

// Signed Minimum (vector). This instruction compares corresponding elements in the vectors in the two source SIMD&FP registers, places the smaller of each of the two signed integer values into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VminS16N VminS16N
//go:noescape
func VminS16N(r *arm.Int16, v0 *arm.Int16, v1 *arm.Int16, n int32)

// Signed Minimum (vector). This instruction compares corresponding elements in the vectors in the two source SIMD&FP registers, places the smaller of each of the two signed integer values into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VminS32N VminS32N
//go:noescape
func VminS32N(r *arm.Int32, v0 *arm.Int32, v1 *arm.Int32, n int32)

// Unsigned Minimum (vector). This instruction compares corresponding vector elements in the two source SIMD&FP registers, places the smaller of each of the two unsigned integer values into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VminU8N VminU8N
//go:noescape
func VminU8N(r *arm.Uint8, v0 *arm.Uint8, v1 *arm.Uint8, n int32)

// Unsigned Minimum (vector). This instruction compares corresponding vector elements in the two source SIMD&FP registers, places the smaller of each of the two unsigned integer values into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VminU16N VminU16N
//go:noescape
func VminU16N(r *arm.Uint16, v0 *arm.Uint16, v1 *arm.Uint16, n int32)

// Unsigned Minimum (vector). This instruction compares corresponding vector elements in the two source SIMD&FP registers, places the smaller of each of the two unsigned integer values into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VminU32N VminU32N
//go:noescape
func VminU32N(r *arm.Uint32, v0 *arm.Uint32, v1 *arm.Uint32, n int32)

// Floating-point minimum (vector). This instruction compares corresponding elements in the vectors in the two source SIMD&FP registers, places the smaller of each of the two floating-point values into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VminF32N VminF32N
//go:noescape
func VminF32N(r *arm.Float32, v0 *arm.Float32, v1 *arm.Float32, n int32)

// Floating-point minimum (vector). This instruction compares corresponding elements in the vectors in the two source SIMD&FP registers, places the smaller of each of the two floating-point values into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VminF64N VminF64N
//go:noescape
func VminF64N(r *arm.Float64, v0 *arm.Float64, v1 *arm.Float64, n int32)

// Floating-point Minimum Number (vector). This instruction compares corresponding vector elements in the two source SIMD&FP registers, writes the smaller of the two floating-point values into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VminnmF32N VminnmF32N
//go:noescape
func VminnmF32N(r *arm.Float32, v0 *arm.Float32, v1 *arm.Float32, n int32)

// Floating-point Minimum Number (vector). This instruction compares corresponding vector elements in the two source SIMD&FP registers, writes the smaller of the two floating-point values into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VminnmF64N VminnmF64N
//go:noescape
func VminnmF64N(r *arm.Float64, v0 *arm.Float64, v1 *arm.Float64, n int32)

// Floating-point Minimum Number (vector). This instruction compares corresponding vector elements in the two source SIMD&FP registers, writes the smaller of the two floating-point values into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VminnmqF32N VminnmqF32N
//go:noescape
func VminnmqF32N(r *arm.Float32, v0 *arm.Float32, v1 *arm.Float32, n int32)

// Floating-point Minimum Number (vector). This instruction compares corresponding vector elements in the two source SIMD&FP registers, writes the smaller of the two floating-point values into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VminnmqF64N VminnmqF64N
//go:noescape
func VminnmqF64N(r *arm.Float64, v0 *arm.Float64, v1 *arm.Float64, n int32)

// Floating-point Minimum Number Pairwise (vector). This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements in the two source SIMD&FP registers, writes the smallest of each pair of floating-point values into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are floating-point values.
//
//go:linkname VminnmvF32N VminnmvF32N
//go:noescape
func VminnmvF32N(r *arm.Float32, v0 *arm.Float32, n int32)

// Floating-point Minimum Number across Vector. This instruction compares all the vector elements in the source SIMD&FP register, and writes the smallest of the values as a scalar to the destination SIMD&FP register. All the values in this instruction are floating-point values.
//
//go:linkname VminnmvqF32N VminnmvqF32N
//go:noescape
func VminnmvqF32N(r *arm.Float32, v0 *arm.Float32, n int32)

// Floating-point Minimum Number Pairwise (vector). This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements in the two source SIMD&FP registers, writes the smallest of each pair of floating-point values into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are floating-point values.
//
//go:linkname VminnmvqF64N VminnmvqF64N
//go:noescape
func VminnmvqF64N(r *arm.Float64, v0 *arm.Float64, n int32)

// Signed Minimum (vector). This instruction compares corresponding elements in the vectors in the two source SIMD&FP registers, places the smaller of each of the two signed integer values into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VminqS8N VminqS8N
//go:noescape
func VminqS8N(r *arm.Int8, v0 *arm.Int8, v1 *arm.Int8, n int32)

// Signed Minimum (vector). This instruction compares corresponding elements in the vectors in the two source SIMD&FP registers, places the smaller of each of the two signed integer values into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VminqS16N VminqS16N
//go:noescape
func VminqS16N(r *arm.Int16, v0 *arm.Int16, v1 *arm.Int16, n int32)

// Signed Minimum (vector). This instruction compares corresponding elements in the vectors in the two source SIMD&FP registers, places the smaller of each of the two signed integer values into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VminqS32N VminqS32N
//go:noescape
func VminqS32N(r *arm.Int32, v0 *arm.Int32, v1 *arm.Int32, n int32)

// Unsigned Minimum (vector). This instruction compares corresponding vector elements in the two source SIMD&FP registers, places the smaller of each of the two unsigned integer values into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VminqU8N VminqU8N
//go:noescape
func VminqU8N(r *arm.Uint8, v0 *arm.Uint8, v1 *arm.Uint8, n int32)

// Unsigned Minimum (vector). This instruction compares corresponding vector elements in the two source SIMD&FP registers, places the smaller of each of the two unsigned integer values into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VminqU16N VminqU16N
//go:noescape
func VminqU16N(r *arm.Uint16, v0 *arm.Uint16, v1 *arm.Uint16, n int32)

// Unsigned Minimum (vector). This instruction compares corresponding vector elements in the two source SIMD&FP registers, places the smaller of each of the two unsigned integer values into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VminqU32N VminqU32N
//go:noescape
func VminqU32N(r *arm.Uint32, v0 *arm.Uint32, v1 *arm.Uint32, n int32)

// Floating-point minimum (vector). This instruction compares corresponding elements in the vectors in the two source SIMD&FP registers, places the smaller of each of the two floating-point values into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VminqF32N VminqF32N
//go:noescape
func VminqF32N(r *arm.Float32, v0 *arm.Float32, v1 *arm.Float32, n int32)

// Floating-point minimum (vector). This instruction compares corresponding elements in the vectors in the two source SIMD&FP registers, places the smaller of each of the two floating-point values into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VminqF64N VminqF64N
//go:noescape
func VminqF64N(r *arm.Float64, v0 *arm.Float64, v1 *arm.Float64, n int32)

// Signed Minimum across Vector. This instruction compares all the vector elements in the source SIMD&FP register, and writes the smallest of the values as a scalar to the destination SIMD&FP register. All the values in this instruction are signed integer values.
//
//go:linkname VminvS8N VminvS8N
//go:noescape
func VminvS8N(r *arm.Int8, v0 *arm.Int8, n int32)

// Signed Minimum across Vector. This instruction compares all the vector elements in the source SIMD&FP register, and writes the smallest of the values as a scalar to the destination SIMD&FP register. All the values in this instruction are signed integer values.
//
//go:linkname VminvS16N VminvS16N
//go:noescape
func VminvS16N(r *arm.Int16, v0 *arm.Int16, n int32)

// Signed Minimum Pairwise. This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements in the two source SIMD&FP registers, writes the smallest of each pair of signed integer values into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VminvS32N VminvS32N
//go:noescape
func VminvS32N(r *arm.Int32, v0 *arm.Int32, n int32)

// Unsigned Minimum across Vector. This instruction compares all the vector elements in the source SIMD&FP register, and writes the smallest of the values as a scalar to the destination SIMD&FP register. All the values in this instruction are unsigned integer values.
//
//go:linkname VminvU8N VminvU8N
//go:noescape
func VminvU8N(r *arm.Uint8, v0 *arm.Uint8, n int32)

// Unsigned Minimum across Vector. This instruction compares all the vector elements in the source SIMD&FP register, and writes the smallest of the values as a scalar to the destination SIMD&FP register. All the values in this instruction are unsigned integer values.
//
//go:linkname VminvU16N VminvU16N
//go:noescape
func VminvU16N(r *arm.Uint16, v0 *arm.Uint16, n int32)

// Unsigned Minimum Pairwise. This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements in the two source SIMD&FP registers, writes the smallest of each pair of unsigned integer values into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VminvU32N VminvU32N
//go:noescape
func VminvU32N(r *arm.Uint32, v0 *arm.Uint32, n int32)

// Floating-point Minimum Pairwise (vector). This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements from the concatenated vector, writes the smaller of each pair of values into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are floating-point values.
//
//go:linkname VminvF32N VminvF32N
//go:noescape
func VminvF32N(r *arm.Float32, v0 *arm.Float32, n int32)

// Signed Minimum across Vector. This instruction compares all the vector elements in the source SIMD&FP register, and writes the smallest of the values as a scalar to the destination SIMD&FP register. All the values in this instruction are signed integer values.
//
//go:linkname VminvqS8N VminvqS8N
//go:noescape
func VminvqS8N(r *arm.Int8, v0 *arm.Int8, n int32)

// Signed Minimum across Vector. This instruction compares all the vector elements in the source SIMD&FP register, and writes the smallest of the values as a scalar to the destination SIMD&FP register. All the values in this instruction are signed integer values.
//
//go:linkname VminvqS16N VminvqS16N
//go:noescape
func VminvqS16N(r *arm.Int16, v0 *arm.Int16, n int32)

// Signed Minimum across Vector. This instruction compares all the vector elements in the source SIMD&FP register, and writes the smallest of the values as a scalar to the destination SIMD&FP register. All the values in this instruction are signed integer values.
//
//go:linkname VminvqS32N VminvqS32N
//go:noescape
func VminvqS32N(r *arm.Int32, v0 *arm.Int32, n int32)

// Unsigned Minimum across Vector. This instruction compares all the vector elements in the source SIMD&FP register, and writes the smallest of the values as a scalar to the destination SIMD&FP register. All the values in this instruction are unsigned integer values.
//
//go:linkname VminvqU8N VminvqU8N
//go:noescape
func VminvqU8N(r *arm.Uint8, v0 *arm.Uint8, n int32)

// Unsigned Minimum across Vector. This instruction compares all the vector elements in the source SIMD&FP register, and writes the smallest of the values as a scalar to the destination SIMD&FP register. All the values in this instruction are unsigned integer values.
//
//go:linkname VminvqU16N VminvqU16N
//go:noescape
func VminvqU16N(r *arm.Uint16, v0 *arm.Uint16, n int32)

// Unsigned Minimum across Vector. This instruction compares all the vector elements in the source SIMD&FP register, and writes the smallest of the values as a scalar to the destination SIMD&FP register. All the values in this instruction are unsigned integer values.
//
//go:linkname VminvqU32N VminvqU32N
//go:noescape
func VminvqU32N(r *arm.Uint32, v0 *arm.Uint32, n int32)

// Floating-point Minimum across Vector. This instruction compares all the vector elements in the source SIMD&FP register, and writes the smallest of the values as a scalar to the destination SIMD&FP register. All the values in this instruction are floating-point values.
//
//go:linkname VminvqF32N VminvqF32N
//go:noescape
func VminvqF32N(r *arm.Float32, v0 *arm.Float32, n int32)

// Floating-point Minimum Pairwise (vector). This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements from the concatenated vector, writes the smaller of each pair of values into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are floating-point values.
//
//go:linkname VminvqF64N VminvqF64N
//go:noescape
func VminvqF64N(r *arm.Float64, v0 *arm.Float64, n int32)

// Duplicate vector element to vector or scalar. This instruction duplicates the vector element at the specified element index in the source SIMD&FP register into a scalar or each element in a vector, and writes the result to the destination SIMD&FP register.
//
//go:linkname VmovNS8N VmovNS8N
//go:noescape
func VmovNS8N(r *arm.Int8, v0 *arm.Int8, n int32)

// Duplicate vector element to vector or scalar. This instruction duplicates the vector element at the specified element index in the source SIMD&FP register into a scalar or each element in a vector, and writes the result to the destination SIMD&FP register.
//
//go:linkname VmovNS16N VmovNS16N
//go:noescape
func VmovNS16N(r *arm.Int16, v0 *arm.Int16, n int32)

// Duplicate vector element to vector or scalar. This instruction duplicates the vector element at the specified element index in the source SIMD&FP register into a scalar or each element in a vector, and writes the result to the destination SIMD&FP register.
//
//go:linkname VmovNS32N VmovNS32N
//go:noescape
func VmovNS32N(r *arm.Int32, v0 *arm.Int32, n int32)

// Duplicate vector element to vector or scalar. This instruction duplicates the vector element at the specified element index in the source SIMD&FP register into a scalar or each element in a vector, and writes the result to the destination SIMD&FP register.
//
//go:linkname VmovNS64N VmovNS64N
//go:noescape
func VmovNS64N(r *arm.Int64, v0 *arm.Int64, n int32)

// Duplicate vector element to vector or scalar. This instruction duplicates the vector element at the specified element index in the source SIMD&FP register into a scalar or each element in a vector, and writes the result to the destination SIMD&FP register.
//
//go:linkname VmovNU8N VmovNU8N
//go:noescape
func VmovNU8N(r *arm.Uint8, v0 *arm.Uint8, n int32)

// Duplicate vector element to vector or scalar. This instruction duplicates the vector element at the specified element index in the source SIMD&FP register into a scalar or each element in a vector, and writes the result to the destination SIMD&FP register.
//
//go:linkname VmovNU16N VmovNU16N
//go:noescape
func VmovNU16N(r *arm.Uint16, v0 *arm.Uint16, n int32)

// Duplicate vector element to vector or scalar. This instruction duplicates the vector element at the specified element index in the source SIMD&FP register into a scalar or each element in a vector, and writes the result to the destination SIMD&FP register.
//
//go:linkname VmovNU32N VmovNU32N
//go:noescape
func VmovNU32N(r *arm.Uint32, v0 *arm.Uint32, n int32)

// Duplicate vector element to vector or scalar. This instruction duplicates the vector element at the specified element index in the source SIMD&FP register into a scalar or each element in a vector, and writes the result to the destination SIMD&FP register.
//
//go:linkname VmovNU64N VmovNU64N
//go:noescape
func VmovNU64N(r *arm.Uint64, v0 *arm.Uint64, n int32)

// Duplicate vector element to vector or scalar. This instruction duplicates the vector element at the specified element index in the source SIMD&FP register into a scalar or each element in a vector, and writes the result to the destination SIMD&FP register.
//
//go:linkname VmovNF32N VmovNF32N
//go:noescape
func VmovNF32N(r *arm.Float32, v0 *arm.Float32, n int32)

// Duplicate vector element to vector or scalar. This instruction duplicates the vector element at the specified element index in the source SIMD&FP register into a scalar or each element in a vector, and writes the result to the destination SIMD&FP register.
//
//go:linkname VmovNF64N VmovNF64N
//go:noescape
func VmovNF64N(r *arm.Float64, v0 *arm.Float64, n int32)

// Duplicate vector element to vector or scalar. This instruction duplicates the vector element at the specified element index in the source SIMD&FP register into a scalar or each element in a vector, and writes the result to the destination SIMD&FP register.
//
//go:linkname VmovqNS8N VmovqNS8N
//go:noescape
func VmovqNS8N(r *arm.Int8, v0 *arm.Int8, n int32)

// Duplicate vector element to vector or scalar. This instruction duplicates the vector element at the specified element index in the source SIMD&FP register into a scalar or each element in a vector, and writes the result to the destination SIMD&FP register.
//
//go:linkname VmovqNS16N VmovqNS16N
//go:noescape
func VmovqNS16N(r *arm.Int16, v0 *arm.Int16, n int32)

// Duplicate vector element to vector or scalar. This instruction duplicates the vector element at the specified element index in the source SIMD&FP register into a scalar or each element in a vector, and writes the result to the destination SIMD&FP register.
//
//go:linkname VmovqNS32N VmovqNS32N
//go:noescape
func VmovqNS32N(r *arm.Int32, v0 *arm.Int32, n int32)

// Duplicate vector element to vector or scalar. This instruction duplicates the vector element at the specified element index in the source SIMD&FP register into a scalar or each element in a vector, and writes the result to the destination SIMD&FP register.
//
//go:linkname VmovqNS64N VmovqNS64N
//go:noescape
func VmovqNS64N(r *arm.Int64, v0 *arm.Int64, n int32)

// Duplicate vector element to vector or scalar. This instruction duplicates the vector element at the specified element index in the source SIMD&FP register into a scalar or each element in a vector, and writes the result to the destination SIMD&FP register.
//
//go:linkname VmovqNU8N VmovqNU8N
//go:noescape
func VmovqNU8N(r *arm.Uint8, v0 *arm.Uint8, n int32)

// Duplicate vector element to vector or scalar. This instruction duplicates the vector element at the specified element index in the source SIMD&FP register into a scalar or each element in a vector, and writes the result to the destination SIMD&FP register.
//
//go:linkname VmovqNU16N VmovqNU16N
//go:noescape
func VmovqNU16N(r *arm.Uint16, v0 *arm.Uint16, n int32)

// Duplicate vector element to vector or scalar. This instruction duplicates the vector element at the specified element index in the source SIMD&FP register into a scalar or each element in a vector, and writes the result to the destination SIMD&FP register.
//
//go:linkname VmovqNU32N VmovqNU32N
//go:noescape
func VmovqNU32N(r *arm.Uint32, v0 *arm.Uint32, n int32)

// Duplicate vector element to vector or scalar. This instruction duplicates the vector element at the specified element index in the source SIMD&FP register into a scalar or each element in a vector, and writes the result to the destination SIMD&FP register.
//
//go:linkname VmovqNU64N VmovqNU64N
//go:noescape
func VmovqNU64N(r *arm.Uint64, v0 *arm.Uint64, n int32)

// Duplicate vector element to vector or scalar. This instruction duplicates the vector element at the specified element index in the source SIMD&FP register into a scalar or each element in a vector, and writes the result to the destination SIMD&FP register.
//
//go:linkname VmovqNF32N VmovqNF32N
//go:noescape
func VmovqNF32N(r *arm.Float32, v0 *arm.Float32, n int32)

// Duplicate vector element to vector or scalar. This instruction duplicates the vector element at the specified element index in the source SIMD&FP register into a scalar or each element in a vector, and writes the result to the destination SIMD&FP register.
//
//go:linkname VmovqNF64N VmovqNF64N
//go:noescape
func VmovqNF64N(r *arm.Float64, v0 *arm.Float64, n int32)

// Multiply (vector). This instruction multiplies corresponding elements in the vectors of the two source SIMD&FP registers, places the results in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VmulS8N VmulS8N
//go:noescape
func VmulS8N(r *arm.Int8, v0 *arm.Int8, v1 *arm.Int8, n int32)

// Multiply (vector). This instruction multiplies corresponding elements in the vectors of the two source SIMD&FP registers, places the results in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VmulS16N VmulS16N
//go:noescape
func VmulS16N(r *arm.Int16, v0 *arm.Int16, v1 *arm.Int16, n int32)

// Multiply (vector). This instruction multiplies corresponding elements in the vectors of the two source SIMD&FP registers, places the results in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VmulS32N VmulS32N
//go:noescape
func VmulS32N(r *arm.Int32, v0 *arm.Int32, v1 *arm.Int32, n int32)

// Multiply (vector). This instruction multiplies corresponding elements in the vectors of the two source SIMD&FP registers, places the results in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VmulU8N VmulU8N
//go:noescape
func VmulU8N(r *arm.Uint8, v0 *arm.Uint8, v1 *arm.Uint8, n int32)

// Multiply (vector). This instruction multiplies corresponding elements in the vectors of the two source SIMD&FP registers, places the results in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VmulU16N VmulU16N
//go:noescape
func VmulU16N(r *arm.Uint16, v0 *arm.Uint16, v1 *arm.Uint16, n int32)

// Multiply (vector). This instruction multiplies corresponding elements in the vectors of the two source SIMD&FP registers, places the results in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VmulU32N VmulU32N
//go:noescape
func VmulU32N(r *arm.Uint32, v0 *arm.Uint32, v1 *arm.Uint32, n int32)

// Floating-point Multiply (vector). This instruction multiplies corresponding floating-point values in the vectors in the two source SIMD&FP registers, places the result in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VmulF32N VmulF32N
//go:noescape
func VmulF32N(r *arm.Float32, v0 *arm.Float32, v1 *arm.Float32, n int32)

// Floating-point Multiply (vector). This instruction multiplies corresponding floating-point values in the vectors in the two source SIMD&FP registers, places the result in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VmulF64N VmulF64N
//go:noescape
func VmulF64N(r *arm.Float64, v0 *arm.Float64, v1 *arm.Float64, n int32)

// Multiply (vector). This instruction multiplies corresponding elements in the vectors of the two source SIMD&FP registers, places the results in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VmulqS8N VmulqS8N
//go:noescape
func VmulqS8N(r *arm.Int8, v0 *arm.Int8, v1 *arm.Int8, n int32)

// Multiply (vector). This instruction multiplies corresponding elements in the vectors of the two source SIMD&FP registers, places the results in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VmulqS16N VmulqS16N
//go:noescape
func VmulqS16N(r *arm.Int16, v0 *arm.Int16, v1 *arm.Int16, n int32)

// Multiply (vector). This instruction multiplies corresponding elements in the vectors of the two source SIMD&FP registers, places the results in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VmulqS32N VmulqS32N
//go:noescape
func VmulqS32N(r *arm.Int32, v0 *arm.Int32, v1 *arm.Int32, n int32)

// Multiply (vector). This instruction multiplies corresponding elements in the vectors of the two source SIMD&FP registers, places the results in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VmulqU8N VmulqU8N
//go:noescape
func VmulqU8N(r *arm.Uint8, v0 *arm.Uint8, v1 *arm.Uint8, n int32)

// Multiply (vector). This instruction multiplies corresponding elements in the vectors of the two source SIMD&FP registers, places the results in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VmulqU16N VmulqU16N
//go:noescape
func VmulqU16N(r *arm.Uint16, v0 *arm.Uint16, v1 *arm.Uint16, n int32)

// Multiply (vector). This instruction multiplies corresponding elements in the vectors of the two source SIMD&FP registers, places the results in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VmulqU32N VmulqU32N
//go:noescape
func VmulqU32N(r *arm.Uint32, v0 *arm.Uint32, v1 *arm.Uint32, n int32)

// Floating-point Multiply (vector). This instruction multiplies corresponding floating-point values in the vectors in the two source SIMD&FP registers, places the result in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VmulqF32N VmulqF32N
//go:noescape
func VmulqF32N(r *arm.Float32, v0 *arm.Float32, v1 *arm.Float32, n int32)

// Floating-point Multiply (vector). This instruction multiplies corresponding floating-point values in the vectors in the two source SIMD&FP registers, places the result in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VmulqF64N VmulqF64N
//go:noescape
func VmulqF64N(r *arm.Float64, v0 *arm.Float64, v1 *arm.Float64, n int32)

// Floating-point Multiply extended. This instruction multiplies corresponding floating-point values in the vectors of the two source SIMD&FP registers, places the resulting floating-point values in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VmulxF32N VmulxF32N
//go:noescape
func VmulxF32N(r *arm.Float32, v0 *arm.Float32, v1 *arm.Float32, n int32)

// Floating-point Multiply extended. This instruction multiplies corresponding floating-point values in the vectors of the two source SIMD&FP registers, places the resulting floating-point values in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VmulxF64N VmulxF64N
//go:noescape
func VmulxF64N(r *arm.Float64, v0 *arm.Float64, v1 *arm.Float64, n int32)

// Floating-point Multiply extended. This instruction multiplies corresponding floating-point values in the vectors of the two source SIMD&FP registers, places the resulting floating-point values in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VmulxdF64N VmulxdF64N
//go:noescape
func VmulxdF64N(r *arm.Float64, v0 *arm.Float64, v1 *arm.Float64, n int32)

// Floating-point Multiply extended. This instruction multiplies corresponding floating-point values in the vectors of the two source SIMD&FP registers, places the resulting floating-point values in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VmulxqF32N VmulxqF32N
//go:noescape
func VmulxqF32N(r *arm.Float32, v0 *arm.Float32, v1 *arm.Float32, n int32)

// Floating-point Multiply extended. This instruction multiplies corresponding floating-point values in the vectors of the two source SIMD&FP registers, places the resulting floating-point values in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VmulxqF64N VmulxqF64N
//go:noescape
func VmulxqF64N(r *arm.Float64, v0 *arm.Float64, v1 *arm.Float64, n int32)

// Floating-point Multiply extended. This instruction multiplies corresponding floating-point values in the vectors of the two source SIMD&FP registers, places the resulting floating-point values in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VmulxsF32N VmulxsF32N
//go:noescape
func VmulxsF32N(r *arm.Float32, v0 *arm.Float32, v1 *arm.Float32, n int32)

// Bitwise NOT (vector). This instruction reads each vector element from the source SIMD&FP register, places the inverse of each value into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VmvnS8N VmvnS8N
//go:noescape
func VmvnS8N(r *arm.Int8, v0 *arm.Int8, n int32)

// Bitwise NOT (vector). This instruction reads each vector element from the source SIMD&FP register, places the inverse of each value into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VmvnS16N VmvnS16N
//go:noescape
func VmvnS16N(r *arm.Int16, v0 *arm.Int16, n int32)

// Bitwise NOT (vector). This instruction reads each vector element from the source SIMD&FP register, places the inverse of each value into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VmvnS32N VmvnS32N
//go:noescape
func VmvnS32N(r *arm.Int32, v0 *arm.Int32, n int32)

// Bitwise NOT (vector). This instruction reads each vector element from the source SIMD&FP register, places the inverse of each value into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VmvnU8N VmvnU8N
//go:noescape
func VmvnU8N(r *arm.Uint8, v0 *arm.Uint8, n int32)

// Bitwise NOT (vector). This instruction reads each vector element from the source SIMD&FP register, places the inverse of each value into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VmvnU16N VmvnU16N
//go:noescape
func VmvnU16N(r *arm.Uint16, v0 *arm.Uint16, n int32)

// Bitwise NOT (vector). This instruction reads each vector element from the source SIMD&FP register, places the inverse of each value into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VmvnU32N VmvnU32N
//go:noescape
func VmvnU32N(r *arm.Uint32, v0 *arm.Uint32, n int32)

// Bitwise NOT (vector). This instruction reads each vector element from the source SIMD&FP register, places the inverse of each value into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VmvnqS8N VmvnqS8N
//go:noescape
func VmvnqS8N(r *arm.Int8, v0 *arm.Int8, n int32)

// Bitwise NOT (vector). This instruction reads each vector element from the source SIMD&FP register, places the inverse of each value into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VmvnqS16N VmvnqS16N
//go:noescape
func VmvnqS16N(r *arm.Int16, v0 *arm.Int16, n int32)

// Bitwise NOT (vector). This instruction reads each vector element from the source SIMD&FP register, places the inverse of each value into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VmvnqS32N VmvnqS32N
//go:noescape
func VmvnqS32N(r *arm.Int32, v0 *arm.Int32, n int32)

// Bitwise NOT (vector). This instruction reads each vector element from the source SIMD&FP register, places the inverse of each value into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VmvnqU8N VmvnqU8N
//go:noescape
func VmvnqU8N(r *arm.Uint8, v0 *arm.Uint8, n int32)

// Bitwise NOT (vector). This instruction reads each vector element from the source SIMD&FP register, places the inverse of each value into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VmvnqU16N VmvnqU16N
//go:noescape
func VmvnqU16N(r *arm.Uint16, v0 *arm.Uint16, n int32)

// Bitwise NOT (vector). This instruction reads each vector element from the source SIMD&FP register, places the inverse of each value into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VmvnqU32N VmvnqU32N
//go:noescape
func VmvnqU32N(r *arm.Uint32, v0 *arm.Uint32, n int32)

// Negate (vector). This instruction reads each vector element from the source SIMD&FP register, negates each value, puts the result into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VnegS8N VnegS8N
//go:noescape
func VnegS8N(r *arm.Int8, v0 *arm.Int8, n int32)

// Negate (vector). This instruction reads each vector element from the source SIMD&FP register, negates each value, puts the result into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VnegS16N VnegS16N
//go:noescape
func VnegS16N(r *arm.Int16, v0 *arm.Int16, n int32)

// Negate (vector). This instruction reads each vector element from the source SIMD&FP register, negates each value, puts the result into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VnegS32N VnegS32N
//go:noescape
func VnegS32N(r *arm.Int32, v0 *arm.Int32, n int32)

// Negate (vector). This instruction reads each vector element from the source SIMD&FP register, negates each value, puts the result into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VnegS64N VnegS64N
//go:noescape
func VnegS64N(r *arm.Int64, v0 *arm.Int64, n int32)

// Floating-point Negate (vector). This instruction negates the value of each vector element in the source SIMD&FP register, writes the result to a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VnegF32N VnegF32N
//go:noescape
func VnegF32N(r *arm.Float32, v0 *arm.Float32, n int32)

// Floating-point Negate (vector). This instruction negates the value of each vector element in the source SIMD&FP register, writes the result to a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VnegF64N VnegF64N
//go:noescape
func VnegF64N(r *arm.Float64, v0 *arm.Float64, n int32)

// Negate (vector). This instruction reads each vector element from the source SIMD&FP register, negates each value, puts the result into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VnegdS64N VnegdS64N
//go:noescape
func VnegdS64N(r *arm.Int64, v0 *arm.Int64, n int32)

// Negate (vector). This instruction reads each vector element from the source SIMD&FP register, negates each value, puts the result into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VnegqS8N VnegqS8N
//go:noescape
func VnegqS8N(r *arm.Int8, v0 *arm.Int8, n int32)

// Negate (vector). This instruction reads each vector element from the source SIMD&FP register, negates each value, puts the result into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VnegqS16N VnegqS16N
//go:noescape
func VnegqS16N(r *arm.Int16, v0 *arm.Int16, n int32)

// Negate (vector). This instruction reads each vector element from the source SIMD&FP register, negates each value, puts the result into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VnegqS32N VnegqS32N
//go:noescape
func VnegqS32N(r *arm.Int32, v0 *arm.Int32, n int32)

// Negate (vector). This instruction reads each vector element from the source SIMD&FP register, negates each value, puts the result into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VnegqS64N VnegqS64N
//go:noescape
func VnegqS64N(r *arm.Int64, v0 *arm.Int64, n int32)

// Floating-point Negate (vector). This instruction negates the value of each vector element in the source SIMD&FP register, writes the result to a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VnegqF32N VnegqF32N
//go:noescape
func VnegqF32N(r *arm.Float32, v0 *arm.Float32, n int32)

// Floating-point Negate (vector). This instruction negates the value of each vector element in the source SIMD&FP register, writes the result to a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VnegqF64N VnegqF64N
//go:noescape
func VnegqF64N(r *arm.Float64, v0 *arm.Float64, n int32)

// Bitwise inclusive OR NOT (vector). This instruction performs a bitwise OR NOT between the two source SIMD&FP registers, and writes the result to the destination SIMD&FP register.
//
//go:linkname VornS8N VornS8N
//go:noescape
func VornS8N(r *arm.Int8, v0 *arm.Int8, v1 *arm.Int8, n int32)

// Bitwise inclusive OR NOT (vector). This instruction performs a bitwise OR NOT between the two source SIMD&FP registers, and writes the result to the destination SIMD&FP register.
//
//go:linkname VornS16N VornS16N
//go:noescape
func VornS16N(r *arm.Int16, v0 *arm.Int16, v1 *arm.Int16, n int32)

// Bitwise inclusive OR NOT (vector). This instruction performs a bitwise OR NOT between the two source SIMD&FP registers, and writes the result to the destination SIMD&FP register.
//
//go:linkname VornS32N VornS32N
//go:noescape
func VornS32N(r *arm.Int32, v0 *arm.Int32, v1 *arm.Int32, n int32)

// Bitwise inclusive OR NOT (vector). This instruction performs a bitwise OR NOT between the two source SIMD&FP registers, and writes the result to the destination SIMD&FP register.
//
//go:linkname VornS64N VornS64N
//go:noescape
func VornS64N(r *arm.Int64, v0 *arm.Int64, v1 *arm.Int64, n int32)

// Bitwise inclusive OR NOT (vector). This instruction performs a bitwise OR NOT between the two source SIMD&FP registers, and writes the result to the destination SIMD&FP register.
//
//go:linkname VornU8N VornU8N
//go:noescape
func VornU8N(r *arm.Uint8, v0 *arm.Uint8, v1 *arm.Uint8, n int32)

// Bitwise inclusive OR NOT (vector). This instruction performs a bitwise OR NOT between the two source SIMD&FP registers, and writes the result to the destination SIMD&FP register.
//
//go:linkname VornU16N VornU16N
//go:noescape
func VornU16N(r *arm.Uint16, v0 *arm.Uint16, v1 *arm.Uint16, n int32)

// Bitwise inclusive OR NOT (vector). This instruction performs a bitwise OR NOT between the two source SIMD&FP registers, and writes the result to the destination SIMD&FP register.
//
//go:linkname VornU32N VornU32N
//go:noescape
func VornU32N(r *arm.Uint32, v0 *arm.Uint32, v1 *arm.Uint32, n int32)

// Bitwise inclusive OR NOT (vector). This instruction performs a bitwise OR NOT between the two source SIMD&FP registers, and writes the result to the destination SIMD&FP register.
//
//go:linkname VornU64N VornU64N
//go:noescape
func VornU64N(r *arm.Uint64, v0 *arm.Uint64, v1 *arm.Uint64, n int32)

// Bitwise inclusive OR NOT (vector). This instruction performs a bitwise OR NOT between the two source SIMD&FP registers, and writes the result to the destination SIMD&FP register.
//
//go:linkname VornqS8N VornqS8N
//go:noescape
func VornqS8N(r *arm.Int8, v0 *arm.Int8, v1 *arm.Int8, n int32)

// Bitwise inclusive OR NOT (vector). This instruction performs a bitwise OR NOT between the two source SIMD&FP registers, and writes the result to the destination SIMD&FP register.
//
//go:linkname VornqS16N VornqS16N
//go:noescape
func VornqS16N(r *arm.Int16, v0 *arm.Int16, v1 *arm.Int16, n int32)

// Bitwise inclusive OR NOT (vector). This instruction performs a bitwise OR NOT between the two source SIMD&FP registers, and writes the result to the destination SIMD&FP register.
//
//go:linkname VornqS32N VornqS32N
//go:noescape
func VornqS32N(r *arm.Int32, v0 *arm.Int32, v1 *arm.Int32, n int32)

// Bitwise inclusive OR NOT (vector). This instruction performs a bitwise OR NOT between the two source SIMD&FP registers, and writes the result to the destination SIMD&FP register.
//
//go:linkname VornqS64N VornqS64N
//go:noescape
func VornqS64N(r *arm.Int64, v0 *arm.Int64, v1 *arm.Int64, n int32)

// Bitwise inclusive OR NOT (vector). This instruction performs a bitwise OR NOT between the two source SIMD&FP registers, and writes the result to the destination SIMD&FP register.
//
//go:linkname VornqU8N VornqU8N
//go:noescape
func VornqU8N(r *arm.Uint8, v0 *arm.Uint8, v1 *arm.Uint8, n int32)

// Bitwise inclusive OR NOT (vector). This instruction performs a bitwise OR NOT between the two source SIMD&FP registers, and writes the result to the destination SIMD&FP register.
//
//go:linkname VornqU16N VornqU16N
//go:noescape
func VornqU16N(r *arm.Uint16, v0 *arm.Uint16, v1 *arm.Uint16, n int32)

// Bitwise inclusive OR NOT (vector). This instruction performs a bitwise OR NOT between the two source SIMD&FP registers, and writes the result to the destination SIMD&FP register.
//
//go:linkname VornqU32N VornqU32N
//go:noescape
func VornqU32N(r *arm.Uint32, v0 *arm.Uint32, v1 *arm.Uint32, n int32)

// Bitwise inclusive OR NOT (vector). This instruction performs a bitwise OR NOT between the two source SIMD&FP registers, and writes the result to the destination SIMD&FP register.
//
//go:linkname VornqU64N VornqU64N
//go:noescape
func VornqU64N(r *arm.Uint64, v0 *arm.Uint64, v1 *arm.Uint64, n int32)

// Bitwise inclusive OR (vector, register). This instruction performs a bitwise OR between the two source SIMD&FP registers, and writes the result to the destination SIMD&FP register.
//
//go:linkname VorrS8N VorrS8N
//go:noescape
func VorrS8N(r *arm.Int8, v0 *arm.Int8, v1 *arm.Int8, n int32)

// Bitwise inclusive OR (vector, register). This instruction performs a bitwise OR between the two source SIMD&FP registers, and writes the result to the destination SIMD&FP register.
//
//go:linkname VorrS16N VorrS16N
//go:noescape
func VorrS16N(r *arm.Int16, v0 *arm.Int16, v1 *arm.Int16, n int32)

// Bitwise inclusive OR (vector, register). This instruction performs a bitwise OR between the two source SIMD&FP registers, and writes the result to the destination SIMD&FP register.
//
//go:linkname VorrS32N VorrS32N
//go:noescape
func VorrS32N(r *arm.Int32, v0 *arm.Int32, v1 *arm.Int32, n int32)

// Bitwise inclusive OR (vector, register). This instruction performs a bitwise OR between the two source SIMD&FP registers, and writes the result to the destination SIMD&FP register.
//
//go:linkname VorrS64N VorrS64N
//go:noescape
func VorrS64N(r *arm.Int64, v0 *arm.Int64, v1 *arm.Int64, n int32)

// Bitwise inclusive OR (vector, register). This instruction performs a bitwise OR between the two source SIMD&FP registers, and writes the result to the destination SIMD&FP register.
//
//go:linkname VorrU8N VorrU8N
//go:noescape
func VorrU8N(r *arm.Uint8, v0 *arm.Uint8, v1 *arm.Uint8, n int32)

// Bitwise inclusive OR (vector, register). This instruction performs a bitwise OR between the two source SIMD&FP registers, and writes the result to the destination SIMD&FP register.
//
//go:linkname VorrU16N VorrU16N
//go:noescape
func VorrU16N(r *arm.Uint16, v0 *arm.Uint16, v1 *arm.Uint16, n int32)

// Bitwise inclusive OR (vector, register). This instruction performs a bitwise OR between the two source SIMD&FP registers, and writes the result to the destination SIMD&FP register.
//
//go:linkname VorrU32N VorrU32N
//go:noescape
func VorrU32N(r *arm.Uint32, v0 *arm.Uint32, v1 *arm.Uint32, n int32)

// Bitwise inclusive OR (vector, register). This instruction performs a bitwise OR between the two source SIMD&FP registers, and writes the result to the destination SIMD&FP register.
//
//go:linkname VorrU64N VorrU64N
//go:noescape
func VorrU64N(r *arm.Uint64, v0 *arm.Uint64, v1 *arm.Uint64, n int32)

// Bitwise inclusive OR (vector, register). This instruction performs a bitwise OR between the two source SIMD&FP registers, and writes the result to the destination SIMD&FP register.
//
//go:linkname VorrqS8N VorrqS8N
//go:noescape
func VorrqS8N(r *arm.Int8, v0 *arm.Int8, v1 *arm.Int8, n int32)

// Bitwise inclusive OR (vector, register). This instruction performs a bitwise OR between the two source SIMD&FP registers, and writes the result to the destination SIMD&FP register.
//
//go:linkname VorrqS16N VorrqS16N
//go:noescape
func VorrqS16N(r *arm.Int16, v0 *arm.Int16, v1 *arm.Int16, n int32)

// Bitwise inclusive OR (vector, register). This instruction performs a bitwise OR between the two source SIMD&FP registers, and writes the result to the destination SIMD&FP register.
//
//go:linkname VorrqS32N VorrqS32N
//go:noescape
func VorrqS32N(r *arm.Int32, v0 *arm.Int32, v1 *arm.Int32, n int32)

// Bitwise inclusive OR (vector, register). This instruction performs a bitwise OR between the two source SIMD&FP registers, and writes the result to the destination SIMD&FP register.
//
//go:linkname VorrqS64N VorrqS64N
//go:noescape
func VorrqS64N(r *arm.Int64, v0 *arm.Int64, v1 *arm.Int64, n int32)

// Bitwise inclusive OR (vector, register). This instruction performs a bitwise OR between the two source SIMD&FP registers, and writes the result to the destination SIMD&FP register.
//
//go:linkname VorrqU8N VorrqU8N
//go:noescape
func VorrqU8N(r *arm.Uint8, v0 *arm.Uint8, v1 *arm.Uint8, n int32)

// Bitwise inclusive OR (vector, register). This instruction performs a bitwise OR between the two source SIMD&FP registers, and writes the result to the destination SIMD&FP register.
//
//go:linkname VorrqU16N VorrqU16N
//go:noescape
func VorrqU16N(r *arm.Uint16, v0 *arm.Uint16, v1 *arm.Uint16, n int32)

// Bitwise inclusive OR (vector, register). This instruction performs a bitwise OR between the two source SIMD&FP registers, and writes the result to the destination SIMD&FP register.
//
//go:linkname VorrqU32N VorrqU32N
//go:noescape
func VorrqU32N(r *arm.Uint32, v0 *arm.Uint32, v1 *arm.Uint32, n int32)

// Bitwise inclusive OR (vector, register). This instruction performs a bitwise OR between the two source SIMD&FP registers, and writes the result to the destination SIMD&FP register.
//
//go:linkname VorrqU64N VorrqU64N
//go:noescape
func VorrqU64N(r *arm.Uint64, v0 *arm.Uint64, v1 *arm.Uint64, n int32)

// Add Pairwise (vector). This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements from the concatenated vector, adds each pair of values together, places the result into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VpaddS8N VpaddS8N
//go:noescape
func VpaddS8N(r *arm.Int8, v0 *arm.Int8, v1 *arm.Int8, n int32)

// Add Pairwise (vector). This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements from the concatenated vector, adds each pair of values together, places the result into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VpaddS16N VpaddS16N
//go:noescape
func VpaddS16N(r *arm.Int16, v0 *arm.Int16, v1 *arm.Int16, n int32)

// Add Pairwise (vector). This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements from the concatenated vector, adds each pair of values together, places the result into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VpaddS32N VpaddS32N
//go:noescape
func VpaddS32N(r *arm.Int32, v0 *arm.Int32, v1 *arm.Int32, n int32)

// Add Pairwise (vector). This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements from the concatenated vector, adds each pair of values together, places the result into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VpaddU8N VpaddU8N
//go:noescape
func VpaddU8N(r *arm.Uint8, v0 *arm.Uint8, v1 *arm.Uint8, n int32)

// Add Pairwise (vector). This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements from the concatenated vector, adds each pair of values together, places the result into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VpaddU16N VpaddU16N
//go:noescape
func VpaddU16N(r *arm.Uint16, v0 *arm.Uint16, v1 *arm.Uint16, n int32)

// Add Pairwise (vector). This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements from the concatenated vector, adds each pair of values together, places the result into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VpaddU32N VpaddU32N
//go:noescape
func VpaddU32N(r *arm.Uint32, v0 *arm.Uint32, v1 *arm.Uint32, n int32)

// Floating-point Add Pairwise (vector). This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements from the concatenated vector, adds each pair of values together, places the result into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are floating-point values.
//
//go:linkname VpaddF32N VpaddF32N
//go:noescape
func VpaddF32N(r *arm.Float32, v0 *arm.Float32, v1 *arm.Float32, n int32)

// Add Pairwise (vector). This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements from the concatenated vector, adds each pair of values together, places the result into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VpadddS64N VpadddS64N
//go:noescape
func VpadddS64N(r *arm.Int64, v0 *arm.Int64, n int32)

// Add Pairwise (vector). This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements from the concatenated vector, adds each pair of values together, places the result into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VpadddU64N VpadddU64N
//go:noescape
func VpadddU64N(r *arm.Uint64, v0 *arm.Uint64, n int32)

// Floating-point Add Pairwise (vector). This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements from the concatenated vector, adds each pair of values together, places the result into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are floating-point values.
//
//go:linkname VpadddF64N VpadddF64N
//go:noescape
func VpadddF64N(r *arm.Float64, v0 *arm.Float64, n int32)

// Add Pairwise (vector). This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements from the concatenated vector, adds each pair of values together, places the result into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VpaddqS8N VpaddqS8N
//go:noescape
func VpaddqS8N(r *arm.Int8, v0 *arm.Int8, v1 *arm.Int8, n int32)

// Add Pairwise (vector). This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements from the concatenated vector, adds each pair of values together, places the result into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VpaddqS16N VpaddqS16N
//go:noescape
func VpaddqS16N(r *arm.Int16, v0 *arm.Int16, v1 *arm.Int16, n int32)

// Add Pairwise (vector). This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements from the concatenated vector, adds each pair of values together, places the result into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VpaddqS32N VpaddqS32N
//go:noescape
func VpaddqS32N(r *arm.Int32, v0 *arm.Int32, v1 *arm.Int32, n int32)

// Add Pairwise (vector). This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements from the concatenated vector, adds each pair of values together, places the result into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VpaddqS64N VpaddqS64N
//go:noescape
func VpaddqS64N(r *arm.Int64, v0 *arm.Int64, v1 *arm.Int64, n int32)

// Add Pairwise (vector). This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements from the concatenated vector, adds each pair of values together, places the result into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VpaddqU8N VpaddqU8N
//go:noescape
func VpaddqU8N(r *arm.Uint8, v0 *arm.Uint8, v1 *arm.Uint8, n int32)

// Add Pairwise (vector). This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements from the concatenated vector, adds each pair of values together, places the result into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VpaddqU16N VpaddqU16N
//go:noescape
func VpaddqU16N(r *arm.Uint16, v0 *arm.Uint16, v1 *arm.Uint16, n int32)

// Add Pairwise (vector). This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements from the concatenated vector, adds each pair of values together, places the result into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VpaddqU32N VpaddqU32N
//go:noescape
func VpaddqU32N(r *arm.Uint32, v0 *arm.Uint32, v1 *arm.Uint32, n int32)

// Add Pairwise (vector). This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements from the concatenated vector, adds each pair of values together, places the result into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VpaddqU64N VpaddqU64N
//go:noescape
func VpaddqU64N(r *arm.Uint64, v0 *arm.Uint64, v1 *arm.Uint64, n int32)

// Floating-point Add Pairwise (vector). This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements from the concatenated vector, adds each pair of values together, places the result into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are floating-point values.
//
//go:linkname VpaddqF32N VpaddqF32N
//go:noescape
func VpaddqF32N(r *arm.Float32, v0 *arm.Float32, v1 *arm.Float32, n int32)

// Floating-point Add Pairwise (vector). This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements from the concatenated vector, adds each pair of values together, places the result into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are floating-point values.
//
//go:linkname VpaddqF64N VpaddqF64N
//go:noescape
func VpaddqF64N(r *arm.Float64, v0 *arm.Float64, v1 *arm.Float64, n int32)

// Floating-point Add Pairwise (vector). This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements from the concatenated vector, adds each pair of values together, places the result into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are floating-point values.
//
//go:linkname VpaddsF32N VpaddsF32N
//go:noescape
func VpaddsF32N(r *arm.Float32, v0 *arm.Float32, n int32)

// Signed Maximum Pairwise. This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements in the two source SIMD&FP registers, writes the largest of each pair of signed integer values into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VpmaxS8N VpmaxS8N
//go:noescape
func VpmaxS8N(r *arm.Int8, v0 *arm.Int8, v1 *arm.Int8, n int32)

// Signed Maximum Pairwise. This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements in the two source SIMD&FP registers, writes the largest of each pair of signed integer values into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VpmaxS16N VpmaxS16N
//go:noescape
func VpmaxS16N(r *arm.Int16, v0 *arm.Int16, v1 *arm.Int16, n int32)

// Signed Maximum Pairwise. This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements in the two source SIMD&FP registers, writes the largest of each pair of signed integer values into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VpmaxS32N VpmaxS32N
//go:noescape
func VpmaxS32N(r *arm.Int32, v0 *arm.Int32, v1 *arm.Int32, n int32)

// Unsigned Maximum Pairwise. This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements in the two source SIMD&FP registers, writes the largest of each pair of unsigned integer values into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VpmaxU8N VpmaxU8N
//go:noescape
func VpmaxU8N(r *arm.Uint8, v0 *arm.Uint8, v1 *arm.Uint8, n int32)

// Unsigned Maximum Pairwise. This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements in the two source SIMD&FP registers, writes the largest of each pair of unsigned integer values into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VpmaxU16N VpmaxU16N
//go:noescape
func VpmaxU16N(r *arm.Uint16, v0 *arm.Uint16, v1 *arm.Uint16, n int32)

// Unsigned Maximum Pairwise. This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements in the two source SIMD&FP registers, writes the largest of each pair of unsigned integer values into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VpmaxU32N VpmaxU32N
//go:noescape
func VpmaxU32N(r *arm.Uint32, v0 *arm.Uint32, v1 *arm.Uint32, n int32)

// Floating-point Maximum Pairwise (vector). This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements from the concatenated vector, writes the larger of each pair of values into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are floating-point values.
//
//go:linkname VpmaxF32N VpmaxF32N
//go:noescape
func VpmaxF32N(r *arm.Float32, v0 *arm.Float32, v1 *arm.Float32, n int32)

// Floating-point Maximum Number Pairwise (vector). This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements in the two source SIMD&FP registers, writes the largest of each pair of values into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are floating-point values.
//
//go:linkname VpmaxnmF32N VpmaxnmF32N
//go:noescape
func VpmaxnmF32N(r *arm.Float32, v0 *arm.Float32, v1 *arm.Float32, n int32)

// Floating-point Maximum Number Pairwise (vector). This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements in the two source SIMD&FP registers, writes the largest of each pair of values into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are floating-point values.
//
//go:linkname VpmaxnmqF32N VpmaxnmqF32N
//go:noescape
func VpmaxnmqF32N(r *arm.Float32, v0 *arm.Float32, v1 *arm.Float32, n int32)

// Floating-point Maximum Number Pairwise (vector). This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements in the two source SIMD&FP registers, writes the largest of each pair of values into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are floating-point values.
//
//go:linkname VpmaxnmqF64N VpmaxnmqF64N
//go:noescape
func VpmaxnmqF64N(r *arm.Float64, v0 *arm.Float64, v1 *arm.Float64, n int32)

// Floating-point Maximum Number Pairwise (vector). This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements in the two source SIMD&FP registers, writes the largest of each pair of values into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are floating-point values.
//
//go:linkname VpmaxnmqdF64N VpmaxnmqdF64N
//go:noescape
func VpmaxnmqdF64N(r *arm.Float64, v0 *arm.Float64, n int32)

// Floating-point Maximum Number Pairwise (vector). This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements in the two source SIMD&FP registers, writes the largest of each pair of values into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are floating-point values.
//
//go:linkname VpmaxnmsF32N VpmaxnmsF32N
//go:noescape
func VpmaxnmsF32N(r *arm.Float32, v0 *arm.Float32, n int32)

// Signed Maximum Pairwise. This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements in the two source SIMD&FP registers, writes the largest of each pair of signed integer values into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VpmaxqS8N VpmaxqS8N
//go:noescape
func VpmaxqS8N(r *arm.Int8, v0 *arm.Int8, v1 *arm.Int8, n int32)

// Signed Maximum Pairwise. This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements in the two source SIMD&FP registers, writes the largest of each pair of signed integer values into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VpmaxqS16N VpmaxqS16N
//go:noescape
func VpmaxqS16N(r *arm.Int16, v0 *arm.Int16, v1 *arm.Int16, n int32)

// Signed Maximum Pairwise. This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements in the two source SIMD&FP registers, writes the largest of each pair of signed integer values into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VpmaxqS32N VpmaxqS32N
//go:noescape
func VpmaxqS32N(r *arm.Int32, v0 *arm.Int32, v1 *arm.Int32, n int32)

// Unsigned Maximum Pairwise. This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements in the two source SIMD&FP registers, writes the largest of each pair of unsigned integer values into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VpmaxqU8N VpmaxqU8N
//go:noescape
func VpmaxqU8N(r *arm.Uint8, v0 *arm.Uint8, v1 *arm.Uint8, n int32)

// Unsigned Maximum Pairwise. This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements in the two source SIMD&FP registers, writes the largest of each pair of unsigned integer values into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VpmaxqU16N VpmaxqU16N
//go:noescape
func VpmaxqU16N(r *arm.Uint16, v0 *arm.Uint16, v1 *arm.Uint16, n int32)

// Unsigned Maximum Pairwise. This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements in the two source SIMD&FP registers, writes the largest of each pair of unsigned integer values into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VpmaxqU32N VpmaxqU32N
//go:noescape
func VpmaxqU32N(r *arm.Uint32, v0 *arm.Uint32, v1 *arm.Uint32, n int32)

// Floating-point Maximum Pairwise (vector). This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements from the concatenated vector, writes the larger of each pair of values into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are floating-point values.
//
//go:linkname VpmaxqF32N VpmaxqF32N
//go:noescape
func VpmaxqF32N(r *arm.Float32, v0 *arm.Float32, v1 *arm.Float32, n int32)

// Floating-point Maximum Pairwise (vector). This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements from the concatenated vector, writes the larger of each pair of values into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are floating-point values.
//
//go:linkname VpmaxqF64N VpmaxqF64N
//go:noescape
func VpmaxqF64N(r *arm.Float64, v0 *arm.Float64, v1 *arm.Float64, n int32)

// Floating-point Maximum Pairwise (vector). This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements from the concatenated vector, writes the larger of each pair of values into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are floating-point values.
//
//go:linkname VpmaxqdF64N VpmaxqdF64N
//go:noescape
func VpmaxqdF64N(r *arm.Float64, v0 *arm.Float64, n int32)

// Floating-point Maximum Pairwise (vector). This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements from the concatenated vector, writes the larger of each pair of values into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are floating-point values.
//
//go:linkname VpmaxsF32N VpmaxsF32N
//go:noescape
func VpmaxsF32N(r *arm.Float32, v0 *arm.Float32, n int32)

// Signed Minimum Pairwise. This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements in the two source SIMD&FP registers, writes the smallest of each pair of signed integer values into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VpminS8N VpminS8N
//go:noescape
func VpminS8N(r *arm.Int8, v0 *arm.Int8, v1 *arm.Int8, n int32)

// Signed Minimum Pairwise. This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements in the two source SIMD&FP registers, writes the smallest of each pair of signed integer values into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VpminS16N VpminS16N
//go:noescape
func VpminS16N(r *arm.Int16, v0 *arm.Int16, v1 *arm.Int16, n int32)

// Signed Minimum Pairwise. This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements in the two source SIMD&FP registers, writes the smallest of each pair of signed integer values into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VpminS32N VpminS32N
//go:noescape
func VpminS32N(r *arm.Int32, v0 *arm.Int32, v1 *arm.Int32, n int32)

// Unsigned Minimum Pairwise. This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements in the two source SIMD&FP registers, writes the smallest of each pair of unsigned integer values into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VpminU8N VpminU8N
//go:noescape
func VpminU8N(r *arm.Uint8, v0 *arm.Uint8, v1 *arm.Uint8, n int32)

// Unsigned Minimum Pairwise. This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements in the two source SIMD&FP registers, writes the smallest of each pair of unsigned integer values into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VpminU16N VpminU16N
//go:noescape
func VpminU16N(r *arm.Uint16, v0 *arm.Uint16, v1 *arm.Uint16, n int32)

// Unsigned Minimum Pairwise. This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements in the two source SIMD&FP registers, writes the smallest of each pair of unsigned integer values into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VpminU32N VpminU32N
//go:noescape
func VpminU32N(r *arm.Uint32, v0 *arm.Uint32, v1 *arm.Uint32, n int32)

// Floating-point Minimum Pairwise (vector). This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements from the concatenated vector, writes the smaller of each pair of values into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are floating-point values.
//
//go:linkname VpminF32N VpminF32N
//go:noescape
func VpminF32N(r *arm.Float32, v0 *arm.Float32, v1 *arm.Float32, n int32)

// Floating-point Minimum Number Pairwise (vector). This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements in the two source SIMD&FP registers, writes the smallest of each pair of floating-point values into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are floating-point values.
//
//go:linkname VpminnmF32N VpminnmF32N
//go:noescape
func VpminnmF32N(r *arm.Float32, v0 *arm.Float32, v1 *arm.Float32, n int32)

// Floating-point Minimum Number Pairwise (vector). This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements in the two source SIMD&FP registers, writes the smallest of each pair of floating-point values into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are floating-point values.
//
//go:linkname VpminnmqF32N VpminnmqF32N
//go:noescape
func VpminnmqF32N(r *arm.Float32, v0 *arm.Float32, v1 *arm.Float32, n int32)

// Floating-point Minimum Number Pairwise (vector). This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements in the two source SIMD&FP registers, writes the smallest of each pair of floating-point values into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are floating-point values.
//
//go:linkname VpminnmqF64N VpminnmqF64N
//go:noescape
func VpminnmqF64N(r *arm.Float64, v0 *arm.Float64, v1 *arm.Float64, n int32)

// Floating-point Minimum Number Pairwise (vector). This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements in the two source SIMD&FP registers, writes the smallest of each pair of floating-point values into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are floating-point values.
//
//go:linkname VpminnmqdF64N VpminnmqdF64N
//go:noescape
func VpminnmqdF64N(r *arm.Float64, v0 *arm.Float64, n int32)

// Floating-point Minimum Number Pairwise (vector). This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements in the two source SIMD&FP registers, writes the smallest of each pair of floating-point values into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are floating-point values.
//
//go:linkname VpminnmsF32N VpminnmsF32N
//go:noescape
func VpminnmsF32N(r *arm.Float32, v0 *arm.Float32, n int32)

// Signed Minimum Pairwise. This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements in the two source SIMD&FP registers, writes the smallest of each pair of signed integer values into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VpminqS8N VpminqS8N
//go:noescape
func VpminqS8N(r *arm.Int8, v0 *arm.Int8, v1 *arm.Int8, n int32)

// Signed Minimum Pairwise. This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements in the two source SIMD&FP registers, writes the smallest of each pair of signed integer values into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VpminqS16N VpminqS16N
//go:noescape
func VpminqS16N(r *arm.Int16, v0 *arm.Int16, v1 *arm.Int16, n int32)

// Signed Minimum Pairwise. This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements in the two source SIMD&FP registers, writes the smallest of each pair of signed integer values into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VpminqS32N VpminqS32N
//go:noescape
func VpminqS32N(r *arm.Int32, v0 *arm.Int32, v1 *arm.Int32, n int32)

// Unsigned Minimum Pairwise. This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements in the two source SIMD&FP registers, writes the smallest of each pair of unsigned integer values into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VpminqU8N VpminqU8N
//go:noescape
func VpminqU8N(r *arm.Uint8, v0 *arm.Uint8, v1 *arm.Uint8, n int32)

// Unsigned Minimum Pairwise. This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements in the two source SIMD&FP registers, writes the smallest of each pair of unsigned integer values into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VpminqU16N VpminqU16N
//go:noescape
func VpminqU16N(r *arm.Uint16, v0 *arm.Uint16, v1 *arm.Uint16, n int32)

// Unsigned Minimum Pairwise. This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements in the two source SIMD&FP registers, writes the smallest of each pair of unsigned integer values into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VpminqU32N VpminqU32N
//go:noescape
func VpminqU32N(r *arm.Uint32, v0 *arm.Uint32, v1 *arm.Uint32, n int32)

// Floating-point Minimum Pairwise (vector). This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements from the concatenated vector, writes the smaller of each pair of values into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are floating-point values.
//
//go:linkname VpminqF32N VpminqF32N
//go:noescape
func VpminqF32N(r *arm.Float32, v0 *arm.Float32, v1 *arm.Float32, n int32)

// Floating-point Minimum Pairwise (vector). This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements from the concatenated vector, writes the smaller of each pair of values into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are floating-point values.
//
//go:linkname VpminqF64N VpminqF64N
//go:noescape
func VpminqF64N(r *arm.Float64, v0 *arm.Float64, v1 *arm.Float64, n int32)

// Floating-point Minimum Pairwise (vector). This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements from the concatenated vector, writes the smaller of each pair of values into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are floating-point values.
//
//go:linkname VpminqdF64N VpminqdF64N
//go:noescape
func VpminqdF64N(r *arm.Float64, v0 *arm.Float64, n int32)

// Floating-point Minimum Pairwise (vector). This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements from the concatenated vector, writes the smaller of each pair of values into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are floating-point values.
//
//go:linkname VpminsF32N VpminsF32N
//go:noescape
func VpminsF32N(r *arm.Float32, v0 *arm.Float32, n int32)

// Signed saturating Absolute value. This instruction reads each vector element from the source SIMD&FP register, puts the absolute value of the result into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are signed integer values.
//
//go:linkname VqabsS8N VqabsS8N
//go:noescape
func VqabsS8N(r *arm.Int8, v0 *arm.Int8, n int32)

// Signed saturating Absolute value. This instruction reads each vector element from the source SIMD&FP register, puts the absolute value of the result into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are signed integer values.
//
//go:linkname VqabsS16N VqabsS16N
//go:noescape
func VqabsS16N(r *arm.Int16, v0 *arm.Int16, n int32)

// Signed saturating Absolute value. This instruction reads each vector element from the source SIMD&FP register, puts the absolute value of the result into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are signed integer values.
//
//go:linkname VqabsS32N VqabsS32N
//go:noescape
func VqabsS32N(r *arm.Int32, v0 *arm.Int32, n int32)

// Signed saturating Absolute value. This instruction reads each vector element from the source SIMD&FP register, puts the absolute value of the result into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are signed integer values.
//
//go:linkname VqabsS64N VqabsS64N
//go:noescape
func VqabsS64N(r *arm.Int64, v0 *arm.Int64, n int32)

// Signed saturating Absolute value. This instruction reads each vector element from the source SIMD&FP register, puts the absolute value of the result into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are signed integer values.
//
//go:linkname VqabsbS8N VqabsbS8N
//go:noescape
func VqabsbS8N(r *arm.Int8, v0 *arm.Int8, n int32)

// Signed saturating Absolute value. This instruction reads each vector element from the source SIMD&FP register, puts the absolute value of the result into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are signed integer values.
//
//go:linkname VqabsdS64N VqabsdS64N
//go:noescape
func VqabsdS64N(r *arm.Int64, v0 *arm.Int64, n int32)

// Signed saturating Absolute value. This instruction reads each vector element from the source SIMD&FP register, puts the absolute value of the result into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are signed integer values.
//
//go:linkname VqabshS16N VqabshS16N
//go:noescape
func VqabshS16N(r *arm.Int16, v0 *arm.Int16, n int32)

// Signed saturating Absolute value. This instruction reads each vector element from the source SIMD&FP register, puts the absolute value of the result into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are signed integer values.
//
//go:linkname VqabsqS8N VqabsqS8N
//go:noescape
func VqabsqS8N(r *arm.Int8, v0 *arm.Int8, n int32)

// Signed saturating Absolute value. This instruction reads each vector element from the source SIMD&FP register, puts the absolute value of the result into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are signed integer values.
//
//go:linkname VqabsqS16N VqabsqS16N
//go:noescape
func VqabsqS16N(r *arm.Int16, v0 *arm.Int16, n int32)

// Signed saturating Absolute value. This instruction reads each vector element from the source SIMD&FP register, puts the absolute value of the result into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are signed integer values.
//
//go:linkname VqabsqS32N VqabsqS32N
//go:noescape
func VqabsqS32N(r *arm.Int32, v0 *arm.Int32, n int32)

// Signed saturating Absolute value. This instruction reads each vector element from the source SIMD&FP register, puts the absolute value of the result into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are signed integer values.
//
//go:linkname VqabsqS64N VqabsqS64N
//go:noescape
func VqabsqS64N(r *arm.Int64, v0 *arm.Int64, n int32)

// Signed saturating Absolute value. This instruction reads each vector element from the source SIMD&FP register, puts the absolute value of the result into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are signed integer values.
//
//go:linkname VqabssS32N VqabssS32N
//go:noescape
func VqabssS32N(r *arm.Int32, v0 *arm.Int32, n int32)

// Signed saturating Add. This instruction adds the values of corresponding elements of the two source SIMD&FP registers, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VqaddS8N VqaddS8N
//go:noescape
func VqaddS8N(r *arm.Int8, v0 *arm.Int8, v1 *arm.Int8, n int32)

// Signed saturating Add. This instruction adds the values of corresponding elements of the two source SIMD&FP registers, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VqaddS16N VqaddS16N
//go:noescape
func VqaddS16N(r *arm.Int16, v0 *arm.Int16, v1 *arm.Int16, n int32)

// Signed saturating Add. This instruction adds the values of corresponding elements of the two source SIMD&FP registers, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VqaddS32N VqaddS32N
//go:noescape
func VqaddS32N(r *arm.Int32, v0 *arm.Int32, v1 *arm.Int32, n int32)

// Signed saturating Add. This instruction adds the values of corresponding elements of the two source SIMD&FP registers, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VqaddS64N VqaddS64N
//go:noescape
func VqaddS64N(r *arm.Int64, v0 *arm.Int64, v1 *arm.Int64, n int32)

// Unsigned saturating Add. This instruction adds the values of corresponding elements of the two source SIMD&FP registers, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VqaddU8N VqaddU8N
//go:noescape
func VqaddU8N(r *arm.Uint8, v0 *arm.Uint8, v1 *arm.Uint8, n int32)

// Unsigned saturating Add. This instruction adds the values of corresponding elements of the two source SIMD&FP registers, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VqaddU16N VqaddU16N
//go:noescape
func VqaddU16N(r *arm.Uint16, v0 *arm.Uint16, v1 *arm.Uint16, n int32)

// Unsigned saturating Add. This instruction adds the values of corresponding elements of the two source SIMD&FP registers, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VqaddU32N VqaddU32N
//go:noescape
func VqaddU32N(r *arm.Uint32, v0 *arm.Uint32, v1 *arm.Uint32, n int32)

// Unsigned saturating Add. This instruction adds the values of corresponding elements of the two source SIMD&FP registers, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VqaddU64N VqaddU64N
//go:noescape
func VqaddU64N(r *arm.Uint64, v0 *arm.Uint64, v1 *arm.Uint64, n int32)

// Signed saturating Add. This instruction adds the values of corresponding elements of the two source SIMD&FP registers, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VqaddbS8N VqaddbS8N
//go:noescape
func VqaddbS8N(r *arm.Int8, v0 *arm.Int8, v1 *arm.Int8, n int32)

// Unsigned saturating Add. This instruction adds the values of corresponding elements of the two source SIMD&FP registers, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VqaddbU8N VqaddbU8N
//go:noescape
func VqaddbU8N(r *arm.Uint8, v0 *arm.Uint8, v1 *arm.Uint8, n int32)

// Signed saturating Add. This instruction adds the values of corresponding elements of the two source SIMD&FP registers, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VqadddS64N VqadddS64N
//go:noescape
func VqadddS64N(r *arm.Int64, v0 *arm.Int64, v1 *arm.Int64, n int32)

// Unsigned saturating Add. This instruction adds the values of corresponding elements of the two source SIMD&FP registers, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VqadddU64N VqadddU64N
//go:noescape
func VqadddU64N(r *arm.Uint64, v0 *arm.Uint64, v1 *arm.Uint64, n int32)

// Signed saturating Add. This instruction adds the values of corresponding elements of the two source SIMD&FP registers, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VqaddhS16N VqaddhS16N
//go:noescape
func VqaddhS16N(r *arm.Int16, v0 *arm.Int16, v1 *arm.Int16, n int32)

// Unsigned saturating Add. This instruction adds the values of corresponding elements of the two source SIMD&FP registers, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VqaddhU16N VqaddhU16N
//go:noescape
func VqaddhU16N(r *arm.Uint16, v0 *arm.Uint16, v1 *arm.Uint16, n int32)

// Signed saturating Add. This instruction adds the values of corresponding elements of the two source SIMD&FP registers, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VqaddqS8N VqaddqS8N
//go:noescape
func VqaddqS8N(r *arm.Int8, v0 *arm.Int8, v1 *arm.Int8, n int32)

// Signed saturating Add. This instruction adds the values of corresponding elements of the two source SIMD&FP registers, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VqaddqS16N VqaddqS16N
//go:noescape
func VqaddqS16N(r *arm.Int16, v0 *arm.Int16, v1 *arm.Int16, n int32)

// Signed saturating Add. This instruction adds the values of corresponding elements of the two source SIMD&FP registers, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VqaddqS32N VqaddqS32N
//go:noescape
func VqaddqS32N(r *arm.Int32, v0 *arm.Int32, v1 *arm.Int32, n int32)

// Signed saturating Add. This instruction adds the values of corresponding elements of the two source SIMD&FP registers, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VqaddqS64N VqaddqS64N
//go:noescape
func VqaddqS64N(r *arm.Int64, v0 *arm.Int64, v1 *arm.Int64, n int32)

// Unsigned saturating Add. This instruction adds the values of corresponding elements of the two source SIMD&FP registers, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VqaddqU8N VqaddqU8N
//go:noescape
func VqaddqU8N(r *arm.Uint8, v0 *arm.Uint8, v1 *arm.Uint8, n int32)

// Unsigned saturating Add. This instruction adds the values of corresponding elements of the two source SIMD&FP registers, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VqaddqU16N VqaddqU16N
//go:noescape
func VqaddqU16N(r *arm.Uint16, v0 *arm.Uint16, v1 *arm.Uint16, n int32)

// Unsigned saturating Add. This instruction adds the values of corresponding elements of the two source SIMD&FP registers, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VqaddqU32N VqaddqU32N
//go:noescape
func VqaddqU32N(r *arm.Uint32, v0 *arm.Uint32, v1 *arm.Uint32, n int32)

// Unsigned saturating Add. This instruction adds the values of corresponding elements of the two source SIMD&FP registers, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VqaddqU64N VqaddqU64N
//go:noescape
func VqaddqU64N(r *arm.Uint64, v0 *arm.Uint64, v1 *arm.Uint64, n int32)

// Signed saturating Add. This instruction adds the values of corresponding elements of the two source SIMD&FP registers, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VqaddsS32N VqaddsS32N
//go:noescape
func VqaddsS32N(r *arm.Int32, v0 *arm.Int32, v1 *arm.Int32, n int32)

// Unsigned saturating Add. This instruction adds the values of corresponding elements of the two source SIMD&FP registers, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VqaddsU32N VqaddsU32N
//go:noescape
func VqaddsU32N(r *arm.Uint32, v0 *arm.Uint32, v1 *arm.Uint32, n int32)

// Signed saturating Doubling Multiply returning High half. This instruction multiplies the values of corresponding elements of the two source SIMD&FP registers, doubles the results, places the most significant half of the final results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VqdmulhS16N VqdmulhS16N
//go:noescape
func VqdmulhS16N(r *arm.Int16, v0 *arm.Int16, v1 *arm.Int16, n int32)

// Signed saturating Doubling Multiply returning High half. This instruction multiplies the values of corresponding elements of the two source SIMD&FP registers, doubles the results, places the most significant half of the final results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VqdmulhS32N VqdmulhS32N
//go:noescape
func VqdmulhS32N(r *arm.Int32, v0 *arm.Int32, v1 *arm.Int32, n int32)

// Signed saturating Doubling Multiply returning High half. This instruction multiplies the values of corresponding elements of the two source SIMD&FP registers, doubles the results, places the most significant half of the final results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VqdmulhhS16N VqdmulhhS16N
//go:noescape
func VqdmulhhS16N(r *arm.Int16, v0 *arm.Int16, v1 *arm.Int16, n int32)

// Signed saturating Doubling Multiply returning High half. This instruction multiplies the values of corresponding elements of the two source SIMD&FP registers, doubles the results, places the most significant half of the final results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VqdmulhqS16N VqdmulhqS16N
//go:noescape
func VqdmulhqS16N(r *arm.Int16, v0 *arm.Int16, v1 *arm.Int16, n int32)

// Signed saturating Doubling Multiply returning High half. This instruction multiplies the values of corresponding elements of the two source SIMD&FP registers, doubles the results, places the most significant half of the final results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VqdmulhqS32N VqdmulhqS32N
//go:noescape
func VqdmulhqS32N(r *arm.Int32, v0 *arm.Int32, v1 *arm.Int32, n int32)

// Signed saturating Doubling Multiply returning High half. This instruction multiplies the values of corresponding elements of the two source SIMD&FP registers, doubles the results, places the most significant half of the final results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VqdmulhsS32N VqdmulhsS32N
//go:noescape
func VqdmulhsS32N(r *arm.Int32, v0 *arm.Int32, v1 *arm.Int32, n int32)

// Signed saturating Negate. This instruction reads each vector element from the source SIMD&FP register, negates each value, places the result into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are signed integer values.
//
//go:linkname VqnegS8N VqnegS8N
//go:noescape
func VqnegS8N(r *arm.Int8, v0 *arm.Int8, n int32)

// Signed saturating Negate. This instruction reads each vector element from the source SIMD&FP register, negates each value, places the result into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are signed integer values.
//
//go:linkname VqnegS16N VqnegS16N
//go:noescape
func VqnegS16N(r *arm.Int16, v0 *arm.Int16, n int32)

// Signed saturating Negate. This instruction reads each vector element from the source SIMD&FP register, negates each value, places the result into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are signed integer values.
//
//go:linkname VqnegS32N VqnegS32N
//go:noescape
func VqnegS32N(r *arm.Int32, v0 *arm.Int32, n int32)

// Signed saturating Negate. This instruction reads each vector element from the source SIMD&FP register, negates each value, places the result into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are signed integer values.
//
//go:linkname VqnegS64N VqnegS64N
//go:noescape
func VqnegS64N(r *arm.Int64, v0 *arm.Int64, n int32)

// Signed saturating Negate. This instruction reads each vector element from the source SIMD&FP register, negates each value, places the result into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are signed integer values.
//
//go:linkname VqnegbS8N VqnegbS8N
//go:noescape
func VqnegbS8N(r *arm.Int8, v0 *arm.Int8, n int32)

// Signed saturating Negate. This instruction reads each vector element from the source SIMD&FP register, negates each value, places the result into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are signed integer values.
//
//go:linkname VqnegdS64N VqnegdS64N
//go:noescape
func VqnegdS64N(r *arm.Int64, v0 *arm.Int64, n int32)

// Signed saturating Negate. This instruction reads each vector element from the source SIMD&FP register, negates each value, places the result into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are signed integer values.
//
//go:linkname VqneghS16N VqneghS16N
//go:noescape
func VqneghS16N(r *arm.Int16, v0 *arm.Int16, n int32)

// Signed saturating Negate. This instruction reads each vector element from the source SIMD&FP register, negates each value, places the result into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are signed integer values.
//
//go:linkname VqnegqS8N VqnegqS8N
//go:noescape
func VqnegqS8N(r *arm.Int8, v0 *arm.Int8, n int32)

// Signed saturating Negate. This instruction reads each vector element from the source SIMD&FP register, negates each value, places the result into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are signed integer values.
//
//go:linkname VqnegqS16N VqnegqS16N
//go:noescape
func VqnegqS16N(r *arm.Int16, v0 *arm.Int16, n int32)

// Signed saturating Negate. This instruction reads each vector element from the source SIMD&FP register, negates each value, places the result into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are signed integer values.
//
//go:linkname VqnegqS32N VqnegqS32N
//go:noescape
func VqnegqS32N(r *arm.Int32, v0 *arm.Int32, n int32)

// Signed saturating Negate. This instruction reads each vector element from the source SIMD&FP register, negates each value, places the result into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are signed integer values.
//
//go:linkname VqnegqS64N VqnegqS64N
//go:noescape
func VqnegqS64N(r *arm.Int64, v0 *arm.Int64, n int32)

// Signed saturating Negate. This instruction reads each vector element from the source SIMD&FP register, negates each value, places the result into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are signed integer values.
//
//go:linkname VqnegsS32N VqnegsS32N
//go:noescape
func VqnegsS32N(r *arm.Int32, v0 *arm.Int32, n int32)

// Signed saturating Rounding Doubling Multiply returning High half. This instruction multiplies the values of corresponding elements of the two source SIMD&FP registers, doubles the results, places the most significant half of the final results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VqrdmulhS16N VqrdmulhS16N
//go:noescape
func VqrdmulhS16N(r *arm.Int16, v0 *arm.Int16, v1 *arm.Int16, n int32)

// Signed saturating Rounding Doubling Multiply returning High half. This instruction multiplies the values of corresponding elements of the two source SIMD&FP registers, doubles the results, places the most significant half of the final results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VqrdmulhS32N VqrdmulhS32N
//go:noescape
func VqrdmulhS32N(r *arm.Int32, v0 *arm.Int32, v1 *arm.Int32, n int32)

// Signed saturating Rounding Doubling Multiply returning High half. This instruction multiplies the values of corresponding elements of the two source SIMD&FP registers, doubles the results, places the most significant half of the final results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VqrdmulhhS16N VqrdmulhhS16N
//go:noescape
func VqrdmulhhS16N(r *arm.Int16, v0 *arm.Int16, v1 *arm.Int16, n int32)

// Signed saturating Rounding Doubling Multiply returning High half. This instruction multiplies the values of corresponding elements of the two source SIMD&FP registers, doubles the results, places the most significant half of the final results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VqrdmulhqS16N VqrdmulhqS16N
//go:noescape
func VqrdmulhqS16N(r *arm.Int16, v0 *arm.Int16, v1 *arm.Int16, n int32)

// Signed saturating Rounding Doubling Multiply returning High half. This instruction multiplies the values of corresponding elements of the two source SIMD&FP registers, doubles the results, places the most significant half of the final results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VqrdmulhqS32N VqrdmulhqS32N
//go:noescape
func VqrdmulhqS32N(r *arm.Int32, v0 *arm.Int32, v1 *arm.Int32, n int32)

// Signed saturating Rounding Doubling Multiply returning High half. This instruction multiplies the values of corresponding elements of the two source SIMD&FP registers, doubles the results, places the most significant half of the final results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VqrdmulhsS32N VqrdmulhsS32N
//go:noescape
func VqrdmulhsS32N(r *arm.Int32, v0 *arm.Int32, v1 *arm.Int32, n int32)

// Signed saturating Rounding Shift Left (register). This instruction takes each vector element in the first source SIMD&FP register, shifts it by a value from the least significant byte of the corresponding vector element of the second source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VqrshlS8N VqrshlS8N
//go:noescape
func VqrshlS8N(r *arm.Int8, v0 *arm.Int8, v1 *arm.Int8, n int32)

// Signed saturating Rounding Shift Left (register). This instruction takes each vector element in the first source SIMD&FP register, shifts it by a value from the least significant byte of the corresponding vector element of the second source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VqrshlS16N VqrshlS16N
//go:noescape
func VqrshlS16N(r *arm.Int16, v0 *arm.Int16, v1 *arm.Int16, n int32)

// Signed saturating Rounding Shift Left (register). This instruction takes each vector element in the first source SIMD&FP register, shifts it by a value from the least significant byte of the corresponding vector element of the second source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VqrshlS32N VqrshlS32N
//go:noescape
func VqrshlS32N(r *arm.Int32, v0 *arm.Int32, v1 *arm.Int32, n int32)

// Signed saturating Rounding Shift Left (register). This instruction takes each vector element in the first source SIMD&FP register, shifts it by a value from the least significant byte of the corresponding vector element of the second source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VqrshlS64N VqrshlS64N
//go:noescape
func VqrshlS64N(r *arm.Int64, v0 *arm.Int64, v1 *arm.Int64, n int32)

// Signed saturating Rounding Shift Left (register). This instruction takes each vector element in the first source SIMD&FP register, shifts it by a value from the least significant byte of the corresponding vector element of the second source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VqrshlbS8N VqrshlbS8N
//go:noescape
func VqrshlbS8N(r *arm.Int8, v0 *arm.Int8, v1 *arm.Int8, n int32)

// Signed saturating Rounding Shift Left (register). This instruction takes each vector element in the first source SIMD&FP register, shifts it by a value from the least significant byte of the corresponding vector element of the second source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VqrshldS64N VqrshldS64N
//go:noescape
func VqrshldS64N(r *arm.Int64, v0 *arm.Int64, v1 *arm.Int64, n int32)

// Signed saturating Rounding Shift Left (register). This instruction takes each vector element in the first source SIMD&FP register, shifts it by a value from the least significant byte of the corresponding vector element of the second source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VqrshlhS16N VqrshlhS16N
//go:noescape
func VqrshlhS16N(r *arm.Int16, v0 *arm.Int16, v1 *arm.Int16, n int32)

// Signed saturating Rounding Shift Left (register). This instruction takes each vector element in the first source SIMD&FP register, shifts it by a value from the least significant byte of the corresponding vector element of the second source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VqrshlqS8N VqrshlqS8N
//go:noescape
func VqrshlqS8N(r *arm.Int8, v0 *arm.Int8, v1 *arm.Int8, n int32)

// Signed saturating Rounding Shift Left (register). This instruction takes each vector element in the first source SIMD&FP register, shifts it by a value from the least significant byte of the corresponding vector element of the second source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VqrshlqS16N VqrshlqS16N
//go:noescape
func VqrshlqS16N(r *arm.Int16, v0 *arm.Int16, v1 *arm.Int16, n int32)

// Signed saturating Rounding Shift Left (register). This instruction takes each vector element in the first source SIMD&FP register, shifts it by a value from the least significant byte of the corresponding vector element of the second source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VqrshlqS32N VqrshlqS32N
//go:noescape
func VqrshlqS32N(r *arm.Int32, v0 *arm.Int32, v1 *arm.Int32, n int32)

// Signed saturating Rounding Shift Left (register). This instruction takes each vector element in the first source SIMD&FP register, shifts it by a value from the least significant byte of the corresponding vector element of the second source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VqrshlqS64N VqrshlqS64N
//go:noescape
func VqrshlqS64N(r *arm.Int64, v0 *arm.Int64, v1 *arm.Int64, n int32)

// Signed saturating Rounding Shift Left (register). This instruction takes each vector element in the first source SIMD&FP register, shifts it by a value from the least significant byte of the corresponding vector element of the second source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VqrshlsS32N VqrshlsS32N
//go:noescape
func VqrshlsS32N(r *arm.Int32, v0 *arm.Int32, v1 *arm.Int32, n int32)

// Signed saturating Shift Left (register). This instruction takes each element in the vector of the first source SIMD&FP register, shifts each element by a value from the least significant byte of the corresponding element of the second source SIMD&FP register, places the results in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VqshlS8N VqshlS8N
//go:noescape
func VqshlS8N(r *arm.Int8, v0 *arm.Int8, v1 *arm.Int8, n int32)

// Signed saturating Shift Left (register). This instruction takes each element in the vector of the first source SIMD&FP register, shifts each element by a value from the least significant byte of the corresponding element of the second source SIMD&FP register, places the results in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VqshlS16N VqshlS16N
//go:noescape
func VqshlS16N(r *arm.Int16, v0 *arm.Int16, v1 *arm.Int16, n int32)

// Signed saturating Shift Left (register). This instruction takes each element in the vector of the first source SIMD&FP register, shifts each element by a value from the least significant byte of the corresponding element of the second source SIMD&FP register, places the results in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VqshlS32N VqshlS32N
//go:noescape
func VqshlS32N(r *arm.Int32, v0 *arm.Int32, v1 *arm.Int32, n int32)

// Signed saturating Shift Left (register). This instruction takes each element in the vector of the first source SIMD&FP register, shifts each element by a value from the least significant byte of the corresponding element of the second source SIMD&FP register, places the results in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VqshlS64N VqshlS64N
//go:noescape
func VqshlS64N(r *arm.Int64, v0 *arm.Int64, v1 *arm.Int64, n int32)

// Signed saturating Shift Left (register). This instruction takes each element in the vector of the first source SIMD&FP register, shifts each element by a value from the least significant byte of the corresponding element of the second source SIMD&FP register, places the results in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VqshlbS8N VqshlbS8N
//go:noescape
func VqshlbS8N(r *arm.Int8, v0 *arm.Int8, v1 *arm.Int8, n int32)

// Signed saturating Shift Left (register). This instruction takes each element in the vector of the first source SIMD&FP register, shifts each element by a value from the least significant byte of the corresponding element of the second source SIMD&FP register, places the results in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VqshldS64N VqshldS64N
//go:noescape
func VqshldS64N(r *arm.Int64, v0 *arm.Int64, v1 *arm.Int64, n int32)

// Signed saturating Shift Left (register). This instruction takes each element in the vector of the first source SIMD&FP register, shifts each element by a value from the least significant byte of the corresponding element of the second source SIMD&FP register, places the results in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VqshlhS16N VqshlhS16N
//go:noescape
func VqshlhS16N(r *arm.Int16, v0 *arm.Int16, v1 *arm.Int16, n int32)

// Signed saturating Shift Left (register). This instruction takes each element in the vector of the first source SIMD&FP register, shifts each element by a value from the least significant byte of the corresponding element of the second source SIMD&FP register, places the results in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VqshlqS8N VqshlqS8N
//go:noescape
func VqshlqS8N(r *arm.Int8, v0 *arm.Int8, v1 *arm.Int8, n int32)

// Signed saturating Shift Left (register). This instruction takes each element in the vector of the first source SIMD&FP register, shifts each element by a value from the least significant byte of the corresponding element of the second source SIMD&FP register, places the results in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VqshlqS16N VqshlqS16N
//go:noescape
func VqshlqS16N(r *arm.Int16, v0 *arm.Int16, v1 *arm.Int16, n int32)

// Signed saturating Shift Left (register). This instruction takes each element in the vector of the first source SIMD&FP register, shifts each element by a value from the least significant byte of the corresponding element of the second source SIMD&FP register, places the results in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VqshlqS32N VqshlqS32N
//go:noescape
func VqshlqS32N(r *arm.Int32, v0 *arm.Int32, v1 *arm.Int32, n int32)

// Signed saturating Shift Left (register). This instruction takes each element in the vector of the first source SIMD&FP register, shifts each element by a value from the least significant byte of the corresponding element of the second source SIMD&FP register, places the results in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VqshlqS64N VqshlqS64N
//go:noescape
func VqshlqS64N(r *arm.Int64, v0 *arm.Int64, v1 *arm.Int64, n int32)

// Signed saturating Shift Left (register). This instruction takes each element in the vector of the first source SIMD&FP register, shifts each element by a value from the least significant byte of the corresponding element of the second source SIMD&FP register, places the results in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VqshlsS32N VqshlsS32N
//go:noescape
func VqshlsS32N(r *arm.Int32, v0 *arm.Int32, v1 *arm.Int32, n int32)

// Signed saturating Subtract. This instruction subtracts the element values of the second source SIMD&FP register from the corresponding element values of the first source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VqsubS8N VqsubS8N
//go:noescape
func VqsubS8N(r *arm.Int8, v0 *arm.Int8, v1 *arm.Int8, n int32)

// Signed saturating Subtract. This instruction subtracts the element values of the second source SIMD&FP register from the corresponding element values of the first source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VqsubS16N VqsubS16N
//go:noescape
func VqsubS16N(r *arm.Int16, v0 *arm.Int16, v1 *arm.Int16, n int32)

// Signed saturating Subtract. This instruction subtracts the element values of the second source SIMD&FP register from the corresponding element values of the first source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VqsubS32N VqsubS32N
//go:noescape
func VqsubS32N(r *arm.Int32, v0 *arm.Int32, v1 *arm.Int32, n int32)

// Signed saturating Subtract. This instruction subtracts the element values of the second source SIMD&FP register from the corresponding element values of the first source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VqsubS64N VqsubS64N
//go:noescape
func VqsubS64N(r *arm.Int64, v0 *arm.Int64, v1 *arm.Int64, n int32)

// Unsigned saturating Subtract. This instruction subtracts the element values of the second source SIMD&FP register from the corresponding element values of the first source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VqsubU8N VqsubU8N
//go:noescape
func VqsubU8N(r *arm.Uint8, v0 *arm.Uint8, v1 *arm.Uint8, n int32)

// Unsigned saturating Subtract. This instruction subtracts the element values of the second source SIMD&FP register from the corresponding element values of the first source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VqsubU16N VqsubU16N
//go:noescape
func VqsubU16N(r *arm.Uint16, v0 *arm.Uint16, v1 *arm.Uint16, n int32)

// Unsigned saturating Subtract. This instruction subtracts the element values of the second source SIMD&FP register from the corresponding element values of the first source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VqsubU32N VqsubU32N
//go:noescape
func VqsubU32N(r *arm.Uint32, v0 *arm.Uint32, v1 *arm.Uint32, n int32)

// Unsigned saturating Subtract. This instruction subtracts the element values of the second source SIMD&FP register from the corresponding element values of the first source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VqsubU64N VqsubU64N
//go:noescape
func VqsubU64N(r *arm.Uint64, v0 *arm.Uint64, v1 *arm.Uint64, n int32)

// Signed saturating Subtract. This instruction subtracts the element values of the second source SIMD&FP register from the corresponding element values of the first source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VqsubbS8N VqsubbS8N
//go:noescape
func VqsubbS8N(r *arm.Int8, v0 *arm.Int8, v1 *arm.Int8, n int32)

// Unsigned saturating Subtract. This instruction subtracts the element values of the second source SIMD&FP register from the corresponding element values of the first source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VqsubbU8N VqsubbU8N
//go:noescape
func VqsubbU8N(r *arm.Uint8, v0 *arm.Uint8, v1 *arm.Uint8, n int32)

// Signed saturating Subtract. This instruction subtracts the element values of the second source SIMD&FP register from the corresponding element values of the first source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VqsubdS64N VqsubdS64N
//go:noescape
func VqsubdS64N(r *arm.Int64, v0 *arm.Int64, v1 *arm.Int64, n int32)

// Unsigned saturating Subtract. This instruction subtracts the element values of the second source SIMD&FP register from the corresponding element values of the first source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VqsubdU64N VqsubdU64N
//go:noescape
func VqsubdU64N(r *arm.Uint64, v0 *arm.Uint64, v1 *arm.Uint64, n int32)

// Signed saturating Subtract. This instruction subtracts the element values of the second source SIMD&FP register from the corresponding element values of the first source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VqsubhS16N VqsubhS16N
//go:noescape
func VqsubhS16N(r *arm.Int16, v0 *arm.Int16, v1 *arm.Int16, n int32)

// Unsigned saturating Subtract. This instruction subtracts the element values of the second source SIMD&FP register from the corresponding element values of the first source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VqsubhU16N VqsubhU16N
//go:noescape
func VqsubhU16N(r *arm.Uint16, v0 *arm.Uint16, v1 *arm.Uint16, n int32)

// Signed saturating Subtract. This instruction subtracts the element values of the second source SIMD&FP register from the corresponding element values of the first source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VqsubqS8N VqsubqS8N
//go:noescape
func VqsubqS8N(r *arm.Int8, v0 *arm.Int8, v1 *arm.Int8, n int32)

// Signed saturating Subtract. This instruction subtracts the element values of the second source SIMD&FP register from the corresponding element values of the first source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VqsubqS16N VqsubqS16N
//go:noescape
func VqsubqS16N(r *arm.Int16, v0 *arm.Int16, v1 *arm.Int16, n int32)

// Signed saturating Subtract. This instruction subtracts the element values of the second source SIMD&FP register from the corresponding element values of the first source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VqsubqS32N VqsubqS32N
//go:noescape
func VqsubqS32N(r *arm.Int32, v0 *arm.Int32, v1 *arm.Int32, n int32)

// Signed saturating Subtract. This instruction subtracts the element values of the second source SIMD&FP register from the corresponding element values of the first source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VqsubqS64N VqsubqS64N
//go:noescape
func VqsubqS64N(r *arm.Int64, v0 *arm.Int64, v1 *arm.Int64, n int32)

// Unsigned saturating Subtract. This instruction subtracts the element values of the second source SIMD&FP register from the corresponding element values of the first source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VqsubqU8N VqsubqU8N
//go:noescape
func VqsubqU8N(r *arm.Uint8, v0 *arm.Uint8, v1 *arm.Uint8, n int32)

// Unsigned saturating Subtract. This instruction subtracts the element values of the second source SIMD&FP register from the corresponding element values of the first source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VqsubqU16N VqsubqU16N
//go:noescape
func VqsubqU16N(r *arm.Uint16, v0 *arm.Uint16, v1 *arm.Uint16, n int32)

// Unsigned saturating Subtract. This instruction subtracts the element values of the second source SIMD&FP register from the corresponding element values of the first source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VqsubqU32N VqsubqU32N
//go:noescape
func VqsubqU32N(r *arm.Uint32, v0 *arm.Uint32, v1 *arm.Uint32, n int32)

// Unsigned saturating Subtract. This instruction subtracts the element values of the second source SIMD&FP register from the corresponding element values of the first source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VqsubqU64N VqsubqU64N
//go:noescape
func VqsubqU64N(r *arm.Uint64, v0 *arm.Uint64, v1 *arm.Uint64, n int32)

// Signed saturating Subtract. This instruction subtracts the element values of the second source SIMD&FP register from the corresponding element values of the first source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VqsubsS32N VqsubsS32N
//go:noescape
func VqsubsS32N(r *arm.Int32, v0 *arm.Int32, v1 *arm.Int32, n int32)

// Unsigned saturating Subtract. This instruction subtracts the element values of the second source SIMD&FP register from the corresponding element values of the first source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VqsubsU32N VqsubsU32N
//go:noescape
func VqsubsU32N(r *arm.Uint32, v0 *arm.Uint32, v1 *arm.Uint32, n int32)

// Table vector Lookup. This instruction reads each value from the vector elements in the index source SIMD&FP register, uses each result as an index to perform a lookup in a table of bytes that is described by one to four source table SIMD&FP registers, places the lookup result in a vector, and writes the vector to the destination SIMD&FP register. If an index is out of range for the table, the result for that lookup is 0. If more than one source register is used to describe the table, the first source register describes the lowest bytes of the table.
//
//go:linkname Vqtbl1QU8N Vqtbl1QU8N
//go:noescape
func Vqtbl1QU8N(r *arm.Uint8, v0 *arm.Uint8, v1 *arm.Uint8, n int32)

// Rotate and Exclusive OR rotates each 64-bit element of the 128-bit vector in a source SIMD&FP register left by 1, performs a bitwise exclusive OR of the resulting 128-bit vector and the vector in another source SIMD&FP register, and writes the result to the destination SIMD&FP register.
//
//go:linkname Vrax1QU64N Vrax1QU64N
//go:noescape
func Vrax1QU64N(r *arm.Uint64, v0 *arm.Uint64, v1 *arm.Uint64, n int32)

// Reverse Bit order (vector). This instruction reads each vector element from the source SIMD&FP register, reverses the bits of the element, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VrbitS8N VrbitS8N
//go:noescape
func VrbitS8N(r *arm.Int8, v0 *arm.Int8, n int32)

// Reverse Bit order (vector). This instruction reads each vector element from the source SIMD&FP register, reverses the bits of the element, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VrbitU8N VrbitU8N
//go:noescape
func VrbitU8N(r *arm.Uint8, v0 *arm.Uint8, n int32)

// Reverse Bit order (vector). This instruction reads each vector element from the source SIMD&FP register, reverses the bits of the element, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VrbitqS8N VrbitqS8N
//go:noescape
func VrbitqS8N(r *arm.Int8, v0 *arm.Int8, n int32)

// Reverse Bit order (vector). This instruction reads each vector element from the source SIMD&FP register, reverses the bits of the element, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VrbitqU8N VrbitqU8N
//go:noescape
func VrbitqU8N(r *arm.Uint8, v0 *arm.Uint8, n int32)

// Unsigned Reciprocal Estimate. This instruction reads each vector element from the source SIMD&FP register, calculates an approximate inverse for the unsigned integer value, places the result into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VrecpeU32N VrecpeU32N
//go:noescape
func VrecpeU32N(r *arm.Uint32, v0 *arm.Uint32, n int32)

// Floating-point Reciprocal Estimate. This instruction finds an approximate reciprocal estimate for each vector element in the source SIMD&FP register, places the result in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VrecpeF32N VrecpeF32N
//go:noescape
func VrecpeF32N(r *arm.Float32, v0 *arm.Float32, n int32)

// Floating-point Reciprocal Estimate. This instruction finds an approximate reciprocal estimate for each vector element in the source SIMD&FP register, places the result in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VrecpeF64N VrecpeF64N
//go:noescape
func VrecpeF64N(r *arm.Float64, v0 *arm.Float64, n int32)

// Floating-point Reciprocal Estimate. This instruction finds an approximate reciprocal estimate for each vector element in the source SIMD&FP register, places the result in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VrecpedF64N VrecpedF64N
//go:noescape
func VrecpedF64N(r *arm.Float64, v0 *arm.Float64, n int32)

// Unsigned Reciprocal Estimate. This instruction reads each vector element from the source SIMD&FP register, calculates an approximate inverse for the unsigned integer value, places the result into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VrecpeqU32N VrecpeqU32N
//go:noescape
func VrecpeqU32N(r *arm.Uint32, v0 *arm.Uint32, n int32)

// Floating-point Reciprocal Estimate. This instruction finds an approximate reciprocal estimate for each vector element in the source SIMD&FP register, places the result in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VrecpeqF32N VrecpeqF32N
//go:noescape
func VrecpeqF32N(r *arm.Float32, v0 *arm.Float32, n int32)

// Floating-point Reciprocal Estimate. This instruction finds an approximate reciprocal estimate for each vector element in the source SIMD&FP register, places the result in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VrecpeqF64N VrecpeqF64N
//go:noescape
func VrecpeqF64N(r *arm.Float64, v0 *arm.Float64, n int32)

// Floating-point Reciprocal Estimate. This instruction finds an approximate reciprocal estimate for each vector element in the source SIMD&FP register, places the result in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VrecpesF32N VrecpesF32N
//go:noescape
func VrecpesF32N(r *arm.Float32, v0 *arm.Float32, n int32)

// Floating-point Reciprocal Step. This instruction multiplies the corresponding floating-point values in the vectors of the two source SIMD&FP registers, subtracts each of the products from 2.0, places the resulting floating-point values in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VrecpsF32N VrecpsF32N
//go:noescape
func VrecpsF32N(r *arm.Float32, v0 *arm.Float32, v1 *arm.Float32, n int32)

// Floating-point Reciprocal Step. This instruction multiplies the corresponding floating-point values in the vectors of the two source SIMD&FP registers, subtracts each of the products from 2.0, places the resulting floating-point values in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VrecpsF64N VrecpsF64N
//go:noescape
func VrecpsF64N(r *arm.Float64, v0 *arm.Float64, v1 *arm.Float64, n int32)

// Floating-point Reciprocal Step. This instruction multiplies the corresponding floating-point values in the vectors of the two source SIMD&FP registers, subtracts each of the products from 2.0, places the resulting floating-point values in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VrecpsdF64N VrecpsdF64N
//go:noescape
func VrecpsdF64N(r *arm.Float64, v0 *arm.Float64, v1 *arm.Float64, n int32)

// Floating-point Reciprocal Step. This instruction multiplies the corresponding floating-point values in the vectors of the two source SIMD&FP registers, subtracts each of the products from 2.0, places the resulting floating-point values in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VrecpsqF32N VrecpsqF32N
//go:noescape
func VrecpsqF32N(r *arm.Float32, v0 *arm.Float32, v1 *arm.Float32, n int32)

// Floating-point Reciprocal Step. This instruction multiplies the corresponding floating-point values in the vectors of the two source SIMD&FP registers, subtracts each of the products from 2.0, places the resulting floating-point values in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VrecpsqF64N VrecpsqF64N
//go:noescape
func VrecpsqF64N(r *arm.Float64, v0 *arm.Float64, v1 *arm.Float64, n int32)

// Floating-point Reciprocal Step. This instruction multiplies the corresponding floating-point values in the vectors of the two source SIMD&FP registers, subtracts each of the products from 2.0, places the resulting floating-point values in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VrecpssF32N VrecpssF32N
//go:noescape
func VrecpssF32N(r *arm.Float32, v0 *arm.Float32, v1 *arm.Float32, n int32)

// Floating-point Reciprocal exponent (scalar). This instruction finds an approximate reciprocal exponent for each vector element in the source SIMD&FP register, places the result in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VrecpxdF64N VrecpxdF64N
//go:noescape
func VrecpxdF64N(r *arm.Float64, v0 *arm.Float64, n int32)

// Floating-point Reciprocal exponent (scalar). This instruction finds an approximate reciprocal exponent for each vector element in the source SIMD&FP register, places the result in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VrecpxsF32N VrecpxsF32N
//go:noescape
func VrecpxsF32N(r *arm.Float32, v0 *arm.Float32, n int32)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretF32S32N VreinterpretF32S32N
//go:noescape
func VreinterpretF32S32N(r *arm.Float32, v0 *arm.Int32, n int32)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretF32U32N VreinterpretF32U32N
//go:noescape
func VreinterpretF32U32N(r *arm.Float32, v0 *arm.Uint32, n int32)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretF64S64N VreinterpretF64S64N
//go:noescape
func VreinterpretF64S64N(r *arm.Float64, v0 *arm.Int64, n int32)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretF64U64N VreinterpretF64U64N
//go:noescape
func VreinterpretF64U64N(r *arm.Float64, v0 *arm.Uint64, n int32)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretS16U16N VreinterpretS16U16N
//go:noescape
func VreinterpretS16U16N(r *arm.Int16, v0 *arm.Uint16, n int32)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretS32U32N VreinterpretS32U32N
//go:noescape
func VreinterpretS32U32N(r *arm.Int32, v0 *arm.Uint32, n int32)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretS32F32N VreinterpretS32F32N
//go:noescape
func VreinterpretS32F32N(r *arm.Int32, v0 *arm.Float32, n int32)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretS64U64N VreinterpretS64U64N
//go:noescape
func VreinterpretS64U64N(r *arm.Int64, v0 *arm.Uint64, n int32)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretS64F64N VreinterpretS64F64N
//go:noescape
func VreinterpretS64F64N(r *arm.Int64, v0 *arm.Float64, n int32)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretS8U8N VreinterpretS8U8N
//go:noescape
func VreinterpretS8U8N(r *arm.Int8, v0 *arm.Uint8, n int32)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretU16S16N VreinterpretU16S16N
//go:noescape
func VreinterpretU16S16N(r *arm.Uint16, v0 *arm.Int16, n int32)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretU32S32N VreinterpretU32S32N
//go:noescape
func VreinterpretU32S32N(r *arm.Uint32, v0 *arm.Int32, n int32)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretU32F32N VreinterpretU32F32N
//go:noescape
func VreinterpretU32F32N(r *arm.Uint32, v0 *arm.Float32, n int32)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretU64S64N VreinterpretU64S64N
//go:noescape
func VreinterpretU64S64N(r *arm.Uint64, v0 *arm.Int64, n int32)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretU64F64N VreinterpretU64F64N
//go:noescape
func VreinterpretU64F64N(r *arm.Uint64, v0 *arm.Float64, n int32)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretU8S8N VreinterpretU8S8N
//go:noescape
func VreinterpretU8S8N(r *arm.Uint8, v0 *arm.Int8, n int32)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretqF32S32N VreinterpretqF32S32N
//go:noescape
func VreinterpretqF32S32N(r *arm.Float32, v0 *arm.Int32, n int32)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretqF32U32N VreinterpretqF32U32N
//go:noescape
func VreinterpretqF32U32N(r *arm.Float32, v0 *arm.Uint32, n int32)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretqF64S64N VreinterpretqF64S64N
//go:noescape
func VreinterpretqF64S64N(r *arm.Float64, v0 *arm.Int64, n int32)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretqF64U64N VreinterpretqF64U64N
//go:noescape
func VreinterpretqF64U64N(r *arm.Float64, v0 *arm.Uint64, n int32)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretqS16U16N VreinterpretqS16U16N
//go:noescape
func VreinterpretqS16U16N(r *arm.Int16, v0 *arm.Uint16, n int32)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretqS32U32N VreinterpretqS32U32N
//go:noescape
func VreinterpretqS32U32N(r *arm.Int32, v0 *arm.Uint32, n int32)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretqS32F32N VreinterpretqS32F32N
//go:noescape
func VreinterpretqS32F32N(r *arm.Int32, v0 *arm.Float32, n int32)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretqS64U64N VreinterpretqS64U64N
//go:noescape
func VreinterpretqS64U64N(r *arm.Int64, v0 *arm.Uint64, n int32)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretqS64F64N VreinterpretqS64F64N
//go:noescape
func VreinterpretqS64F64N(r *arm.Int64, v0 *arm.Float64, n int32)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretqS8U8N VreinterpretqS8U8N
//go:noescape
func VreinterpretqS8U8N(r *arm.Int8, v0 *arm.Uint8, n int32)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretqU16S16N VreinterpretqU16S16N
//go:noescape
func VreinterpretqU16S16N(r *arm.Uint16, v0 *arm.Int16, n int32)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretqU32S32N VreinterpretqU32S32N
//go:noescape
func VreinterpretqU32S32N(r *arm.Uint32, v0 *arm.Int32, n int32)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretqU32F32N VreinterpretqU32F32N
//go:noescape
func VreinterpretqU32F32N(r *arm.Uint32, v0 *arm.Float32, n int32)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretqU64S64N VreinterpretqU64S64N
//go:noescape
func VreinterpretqU64S64N(r *arm.Uint64, v0 *arm.Int64, n int32)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretqU64F64N VreinterpretqU64F64N
//go:noescape
func VreinterpretqU64F64N(r *arm.Uint64, v0 *arm.Float64, n int32)

// Vector reinterpret cast operation
//
//go:linkname VreinterpretqU8S8N VreinterpretqU8S8N
//go:noescape
func VreinterpretqU8S8N(r *arm.Uint8, v0 *arm.Int8, n int32)

// Reverse elements in 16-bit halfwords (vector). This instruction reverses the order of 8-bit elements in each halfword of the vector in the source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname Vrev16S8N Vrev16S8N
//go:noescape
func Vrev16S8N(r *arm.Int8, v0 *arm.Int8, n int32)

// Reverse elements in 16-bit halfwords (vector). This instruction reverses the order of 8-bit elements in each halfword of the vector in the source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname Vrev16U8N Vrev16U8N
//go:noescape
func Vrev16U8N(r *arm.Uint8, v0 *arm.Uint8, n int32)

// Reverse elements in 16-bit halfwords (vector). This instruction reverses the order of 8-bit elements in each halfword of the vector in the source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname Vrev16QS8N Vrev16QS8N
//go:noescape
func Vrev16QS8N(r *arm.Int8, v0 *arm.Int8, n int32)

// Reverse elements in 16-bit halfwords (vector). This instruction reverses the order of 8-bit elements in each halfword of the vector in the source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname Vrev16QU8N Vrev16QU8N
//go:noescape
func Vrev16QU8N(r *arm.Uint8, v0 *arm.Uint8, n int32)

// Reverse elements in 32-bit words (vector). This instruction reverses the order of 8-bit or 16-bit elements in each word of the vector in the source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname Vrev32S8N Vrev32S8N
//go:noescape
func Vrev32S8N(r *arm.Int8, v0 *arm.Int8, n int32)

// Reverse elements in 32-bit words (vector). This instruction reverses the order of 8-bit or 16-bit elements in each word of the vector in the source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname Vrev32S16N Vrev32S16N
//go:noescape
func Vrev32S16N(r *arm.Int16, v0 *arm.Int16, n int32)

// Reverse elements in 32-bit words (vector). This instruction reverses the order of 8-bit or 16-bit elements in each word of the vector in the source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname Vrev32U8N Vrev32U8N
//go:noescape
func Vrev32U8N(r *arm.Uint8, v0 *arm.Uint8, n int32)

// Reverse elements in 32-bit words (vector). This instruction reverses the order of 8-bit or 16-bit elements in each word of the vector in the source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname Vrev32U16N Vrev32U16N
//go:noescape
func Vrev32U16N(r *arm.Uint16, v0 *arm.Uint16, n int32)

// Reverse elements in 32-bit words (vector). This instruction reverses the order of 8-bit or 16-bit elements in each word of the vector in the source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname Vrev32QS8N Vrev32QS8N
//go:noescape
func Vrev32QS8N(r *arm.Int8, v0 *arm.Int8, n int32)

// Reverse elements in 32-bit words (vector). This instruction reverses the order of 8-bit or 16-bit elements in each word of the vector in the source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname Vrev32QS16N Vrev32QS16N
//go:noescape
func Vrev32QS16N(r *arm.Int16, v0 *arm.Int16, n int32)

// Reverse elements in 32-bit words (vector). This instruction reverses the order of 8-bit or 16-bit elements in each word of the vector in the source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname Vrev32QU8N Vrev32QU8N
//go:noescape
func Vrev32QU8N(r *arm.Uint8, v0 *arm.Uint8, n int32)

// Reverse elements in 32-bit words (vector). This instruction reverses the order of 8-bit or 16-bit elements in each word of the vector in the source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname Vrev32QU16N Vrev32QU16N
//go:noescape
func Vrev32QU16N(r *arm.Uint16, v0 *arm.Uint16, n int32)

// Reverse elements in 64-bit doublewords (vector). This instruction reverses the order of 8-bit, 16-bit, or 32-bit elements in each doubleword of the vector in the source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname Vrev64S8N Vrev64S8N
//go:noescape
func Vrev64S8N(r *arm.Int8, v0 *arm.Int8, n int32)

// Reverse elements in 64-bit doublewords (vector). This instruction reverses the order of 8-bit, 16-bit, or 32-bit elements in each doubleword of the vector in the source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname Vrev64S16N Vrev64S16N
//go:noescape
func Vrev64S16N(r *arm.Int16, v0 *arm.Int16, n int32)

// Reverse elements in 64-bit doublewords (vector). This instruction reverses the order of 8-bit, 16-bit, or 32-bit elements in each doubleword of the vector in the source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname Vrev64S32N Vrev64S32N
//go:noescape
func Vrev64S32N(r *arm.Int32, v0 *arm.Int32, n int32)

// Reverse elements in 64-bit doublewords (vector). This instruction reverses the order of 8-bit, 16-bit, or 32-bit elements in each doubleword of the vector in the source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname Vrev64U8N Vrev64U8N
//go:noescape
func Vrev64U8N(r *arm.Uint8, v0 *arm.Uint8, n int32)

// Reverse elements in 64-bit doublewords (vector). This instruction reverses the order of 8-bit, 16-bit, or 32-bit elements in each doubleword of the vector in the source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname Vrev64U16N Vrev64U16N
//go:noescape
func Vrev64U16N(r *arm.Uint16, v0 *arm.Uint16, n int32)

// Reverse elements in 64-bit doublewords (vector). This instruction reverses the order of 8-bit, 16-bit, or 32-bit elements in each doubleword of the vector in the source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname Vrev64U32N Vrev64U32N
//go:noescape
func Vrev64U32N(r *arm.Uint32, v0 *arm.Uint32, n int32)

// Reverse elements in 64-bit doublewords (vector). This instruction reverses the order of 8-bit, 16-bit, or 32-bit elements in each doubleword of the vector in the source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname Vrev64F32N Vrev64F32N
//go:noescape
func Vrev64F32N(r *arm.Float32, v0 *arm.Float32, n int32)

// Reverse elements in 64-bit doublewords (vector). This instruction reverses the order of 8-bit, 16-bit, or 32-bit elements in each doubleword of the vector in the source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname Vrev64QS8N Vrev64QS8N
//go:noescape
func Vrev64QS8N(r *arm.Int8, v0 *arm.Int8, n int32)

// Reverse elements in 64-bit doublewords (vector). This instruction reverses the order of 8-bit, 16-bit, or 32-bit elements in each doubleword of the vector in the source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname Vrev64QS16N Vrev64QS16N
//go:noescape
func Vrev64QS16N(r *arm.Int16, v0 *arm.Int16, n int32)

// Reverse elements in 64-bit doublewords (vector). This instruction reverses the order of 8-bit, 16-bit, or 32-bit elements in each doubleword of the vector in the source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname Vrev64QS32N Vrev64QS32N
//go:noescape
func Vrev64QS32N(r *arm.Int32, v0 *arm.Int32, n int32)

// Reverse elements in 64-bit doublewords (vector). This instruction reverses the order of 8-bit, 16-bit, or 32-bit elements in each doubleword of the vector in the source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname Vrev64QU8N Vrev64QU8N
//go:noescape
func Vrev64QU8N(r *arm.Uint8, v0 *arm.Uint8, n int32)

// Reverse elements in 64-bit doublewords (vector). This instruction reverses the order of 8-bit, 16-bit, or 32-bit elements in each doubleword of the vector in the source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname Vrev64QU16N Vrev64QU16N
//go:noescape
func Vrev64QU16N(r *arm.Uint16, v0 *arm.Uint16, n int32)

// Reverse elements in 64-bit doublewords (vector). This instruction reverses the order of 8-bit, 16-bit, or 32-bit elements in each doubleword of the vector in the source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname Vrev64QU32N Vrev64QU32N
//go:noescape
func Vrev64QU32N(r *arm.Uint32, v0 *arm.Uint32, n int32)

// Reverse elements in 64-bit doublewords (vector). This instruction reverses the order of 8-bit, 16-bit, or 32-bit elements in each doubleword of the vector in the source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname Vrev64QF32N Vrev64QF32N
//go:noescape
func Vrev64QF32N(r *arm.Float32, v0 *arm.Float32, n int32)

// Signed Rounding Halving Add. This instruction adds corresponding signed integer values from the two source SIMD&FP registers, shifts each result right one bit, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VrhaddS8N VrhaddS8N
//go:noescape
func VrhaddS8N(r *arm.Int8, v0 *arm.Int8, v1 *arm.Int8, n int32)

// Signed Rounding Halving Add. This instruction adds corresponding signed integer values from the two source SIMD&FP registers, shifts each result right one bit, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VrhaddS16N VrhaddS16N
//go:noescape
func VrhaddS16N(r *arm.Int16, v0 *arm.Int16, v1 *arm.Int16, n int32)

// Signed Rounding Halving Add. This instruction adds corresponding signed integer values from the two source SIMD&FP registers, shifts each result right one bit, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VrhaddS32N VrhaddS32N
//go:noescape
func VrhaddS32N(r *arm.Int32, v0 *arm.Int32, v1 *arm.Int32, n int32)

// Unsigned Rounding Halving Add. This instruction adds corresponding unsigned integer values from the two source SIMD&FP registers, shifts each result right one bit, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VrhaddU8N VrhaddU8N
//go:noescape
func VrhaddU8N(r *arm.Uint8, v0 *arm.Uint8, v1 *arm.Uint8, n int32)

// Unsigned Rounding Halving Add. This instruction adds corresponding unsigned integer values from the two source SIMD&FP registers, shifts each result right one bit, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VrhaddU16N VrhaddU16N
//go:noescape
func VrhaddU16N(r *arm.Uint16, v0 *arm.Uint16, v1 *arm.Uint16, n int32)

// Unsigned Rounding Halving Add. This instruction adds corresponding unsigned integer values from the two source SIMD&FP registers, shifts each result right one bit, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VrhaddU32N VrhaddU32N
//go:noescape
func VrhaddU32N(r *arm.Uint32, v0 *arm.Uint32, v1 *arm.Uint32, n int32)

// Signed Rounding Halving Add. This instruction adds corresponding signed integer values from the two source SIMD&FP registers, shifts each result right one bit, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VrhaddqS8N VrhaddqS8N
//go:noescape
func VrhaddqS8N(r *arm.Int8, v0 *arm.Int8, v1 *arm.Int8, n int32)

// Signed Rounding Halving Add. This instruction adds corresponding signed integer values from the two source SIMD&FP registers, shifts each result right one bit, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VrhaddqS16N VrhaddqS16N
//go:noescape
func VrhaddqS16N(r *arm.Int16, v0 *arm.Int16, v1 *arm.Int16, n int32)

// Signed Rounding Halving Add. This instruction adds corresponding signed integer values from the two source SIMD&FP registers, shifts each result right one bit, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VrhaddqS32N VrhaddqS32N
//go:noescape
func VrhaddqS32N(r *arm.Int32, v0 *arm.Int32, v1 *arm.Int32, n int32)

// Unsigned Rounding Halving Add. This instruction adds corresponding unsigned integer values from the two source SIMD&FP registers, shifts each result right one bit, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VrhaddqU8N VrhaddqU8N
//go:noescape
func VrhaddqU8N(r *arm.Uint8, v0 *arm.Uint8, v1 *arm.Uint8, n int32)

// Unsigned Rounding Halving Add. This instruction adds corresponding unsigned integer values from the two source SIMD&FP registers, shifts each result right one bit, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VrhaddqU16N VrhaddqU16N
//go:noescape
func VrhaddqU16N(r *arm.Uint16, v0 *arm.Uint16, v1 *arm.Uint16, n int32)

// Unsigned Rounding Halving Add. This instruction adds corresponding unsigned integer values from the two source SIMD&FP registers, shifts each result right one bit, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VrhaddqU32N VrhaddqU32N
//go:noescape
func VrhaddqU32N(r *arm.Uint32, v0 *arm.Uint32, v1 *arm.Uint32, n int32)

// Floating-point Round to Integral, toward Zero (vector). This instruction rounds a vector of floating-point values in the SIMD&FP source register to integral floating-point values of the same size using the Round towards Zero rounding mode, and writes the result to the SIMD&FP destination register.
//
//go:linkname VrndF32N VrndF32N
//go:noescape
func VrndF32N(r *arm.Float32, v0 *arm.Float32, n int32)

// Floating-point Round to Integral, toward Zero (vector). This instruction rounds a vector of floating-point values in the SIMD&FP source register to integral floating-point values of the same size using the Round towards Zero rounding mode, and writes the result to the SIMD&FP destination register.
//
//go:linkname VrndF64N VrndF64N
//go:noescape
func VrndF64N(r *arm.Float64, v0 *arm.Float64, n int32)

// Floating-point Round to 32-bit Integer, using current rounding mode (vector). This instruction rounds a vector of floating-point values in the SIMD&FP source register to integral floating-point values that fit into a 32-bit integer size using the rounding mode that is determined by the FPCR, and writes the result to the SIMD&FP destination register.
//
//go:linkname Vrnd32XF32N Vrnd32XF32N
//go:noescape
func Vrnd32XF32N(r *arm.Float32, v0 *arm.Float32, n int32)

// Floating-point Round to 32-bit Integer, using current rounding mode (vector). This instruction rounds a vector of floating-point values in the SIMD&FP source register to integral floating-point values that fit into a 32-bit integer size using the rounding mode that is determined by the FPCR, and writes the result to the SIMD&FP destination register.
//
//go:linkname Vrnd32XF64N Vrnd32XF64N
//go:noescape
func Vrnd32XF64N(r *arm.Float64, v0 *arm.Float64, n int32)

// Floating-point Round to 32-bit Integer, using current rounding mode (vector). This instruction rounds a vector of floating-point values in the SIMD&FP source register to integral floating-point values that fit into a 32-bit integer size using the rounding mode that is determined by the FPCR, and writes the result to the SIMD&FP destination register.
//
//go:linkname Vrnd32XqF32N Vrnd32XqF32N
//go:noescape
func Vrnd32XqF32N(r *arm.Float32, v0 *arm.Float32, n int32)

// Floating-point Round to 32-bit Integer, using current rounding mode (vector). This instruction rounds a vector of floating-point values in the SIMD&FP source register to integral floating-point values that fit into a 32-bit integer size using the rounding mode that is determined by the FPCR, and writes the result to the SIMD&FP destination register.
//
//go:linkname Vrnd32XqF64N Vrnd32XqF64N
//go:noescape
func Vrnd32XqF64N(r *arm.Float64, v0 *arm.Float64, n int32)

// Floating-point Round to 32-bit Integer toward Zero (vector). This instruction rounds a vector of floating-point values in the SIMD&FP source register to integral floating-point values that fit into a 32-bit integer size using the Round towards Zero rounding mode, and writes the result to the SIMD&FP destination register.
//
//go:linkname Vrnd32ZF32N Vrnd32ZF32N
//go:noescape
func Vrnd32ZF32N(r *arm.Float32, v0 *arm.Float32, n int32)

// Floating-point Round to 32-bit Integer toward Zero (vector). This instruction rounds a vector of floating-point values in the SIMD&FP source register to integral floating-point values that fit into a 32-bit integer size using the Round towards Zero rounding mode, and writes the result to the SIMD&FP destination register.
//
//go:linkname Vrnd32ZF64N Vrnd32ZF64N
//go:noescape
func Vrnd32ZF64N(r *arm.Float64, v0 *arm.Float64, n int32)

// Floating-point Round to 32-bit Integer toward Zero (vector). This instruction rounds a vector of floating-point values in the SIMD&FP source register to integral floating-point values that fit into a 32-bit integer size using the Round towards Zero rounding mode, and writes the result to the SIMD&FP destination register.
//
//go:linkname Vrnd32ZqF32N Vrnd32ZqF32N
//go:noescape
func Vrnd32ZqF32N(r *arm.Float32, v0 *arm.Float32, n int32)

// Floating-point Round to 32-bit Integer toward Zero (vector). This instruction rounds a vector of floating-point values in the SIMD&FP source register to integral floating-point values that fit into a 32-bit integer size using the Round towards Zero rounding mode, and writes the result to the SIMD&FP destination register.
//
//go:linkname Vrnd32ZqF64N Vrnd32ZqF64N
//go:noescape
func Vrnd32ZqF64N(r *arm.Float64, v0 *arm.Float64, n int32)

// Floating-point Round to 64-bit Integer, using current rounding mode (vector). This instruction rounds a vector of floating-point values in the SIMD&FP source register to integral floating-point values that fit into a 64-bit integer size using the rounding mode that is determined by the FPCR, and writes the result to the SIMD&FP destination register.
//
//go:linkname Vrnd64XF32N Vrnd64XF32N
//go:noescape
func Vrnd64XF32N(r *arm.Float32, v0 *arm.Float32, n int32)

// Floating-point Round to 64-bit Integer, using current rounding mode (vector). This instruction rounds a vector of floating-point values in the SIMD&FP source register to integral floating-point values that fit into a 64-bit integer size using the rounding mode that is determined by the FPCR, and writes the result to the SIMD&FP destination register.
//
//go:linkname Vrnd64XF64N Vrnd64XF64N
//go:noescape
func Vrnd64XF64N(r *arm.Float64, v0 *arm.Float64, n int32)

// Floating-point Round to 64-bit Integer, using current rounding mode (vector). This instruction rounds a vector of floating-point values in the SIMD&FP source register to integral floating-point values that fit into a 64-bit integer size using the rounding mode that is determined by the FPCR, and writes the result to the SIMD&FP destination register.
//
//go:linkname Vrnd64XqF32N Vrnd64XqF32N
//go:noescape
func Vrnd64XqF32N(r *arm.Float32, v0 *arm.Float32, n int32)

// Floating-point Round to 64-bit Integer, using current rounding mode (vector). This instruction rounds a vector of floating-point values in the SIMD&FP source register to integral floating-point values that fit into a 64-bit integer size using the rounding mode that is determined by the FPCR, and writes the result to the SIMD&FP destination register.
//
//go:linkname Vrnd64XqF64N Vrnd64XqF64N
//go:noescape
func Vrnd64XqF64N(r *arm.Float64, v0 *arm.Float64, n int32)

// Floating-point Round to 64-bit Integer toward Zero (vector). This instruction rounds a vector of floating-point values in the SIMD&FP source register to integral floating-point values that fit into a 64-bit integer size using the Round towards Zero rounding mode, and writes the result to the SIMD&FP destination register.
//
//go:linkname Vrnd64ZF32N Vrnd64ZF32N
//go:noescape
func Vrnd64ZF32N(r *arm.Float32, v0 *arm.Float32, n int32)

// Floating-point Round to 64-bit Integer toward Zero (vector). This instruction rounds a vector of floating-point values in the SIMD&FP source register to integral floating-point values that fit into a 64-bit integer size using the Round towards Zero rounding mode, and writes the result to the SIMD&FP destination register.
//
//go:linkname Vrnd64ZF64N Vrnd64ZF64N
//go:noescape
func Vrnd64ZF64N(r *arm.Float64, v0 *arm.Float64, n int32)

// Floating-point Round to 64-bit Integer toward Zero (vector). This instruction rounds a vector of floating-point values in the SIMD&FP source register to integral floating-point values that fit into a 64-bit integer size using the Round towards Zero rounding mode, and writes the result to the SIMD&FP destination register.
//
//go:linkname Vrnd64ZqF32N Vrnd64ZqF32N
//go:noescape
func Vrnd64ZqF32N(r *arm.Float32, v0 *arm.Float32, n int32)

// Floating-point Round to 64-bit Integer toward Zero (vector). This instruction rounds a vector of floating-point values in the SIMD&FP source register to integral floating-point values that fit into a 64-bit integer size using the Round towards Zero rounding mode, and writes the result to the SIMD&FP destination register.
//
//go:linkname Vrnd64ZqF64N Vrnd64ZqF64N
//go:noescape
func Vrnd64ZqF64N(r *arm.Float64, v0 *arm.Float64, n int32)

// Floating-point Round to Integral, to nearest with ties to Away (vector). This instruction rounds a vector of floating-point values in the SIMD&FP source register to integral floating-point values of the same size using the Round to Nearest with Ties to Away rounding mode, and writes the result to the SIMD&FP destination register.
//
//go:linkname VrndaF32N VrndaF32N
//go:noescape
func VrndaF32N(r *arm.Float32, v0 *arm.Float32, n int32)

// Floating-point Round to Integral, to nearest with ties to Away (vector). This instruction rounds a vector of floating-point values in the SIMD&FP source register to integral floating-point values of the same size using the Round to Nearest with Ties to Away rounding mode, and writes the result to the SIMD&FP destination register.
//
//go:linkname VrndaF64N VrndaF64N
//go:noescape
func VrndaF64N(r *arm.Float64, v0 *arm.Float64, n int32)

// Floating-point Round to Integral, to nearest with ties to Away (vector). This instruction rounds a vector of floating-point values in the SIMD&FP source register to integral floating-point values of the same size using the Round to Nearest with Ties to Away rounding mode, and writes the result to the SIMD&FP destination register.
//
//go:linkname VrndaqF32N VrndaqF32N
//go:noescape
func VrndaqF32N(r *arm.Float32, v0 *arm.Float32, n int32)

// Floating-point Round to Integral, to nearest with ties to Away (vector). This instruction rounds a vector of floating-point values in the SIMD&FP source register to integral floating-point values of the same size using the Round to Nearest with Ties to Away rounding mode, and writes the result to the SIMD&FP destination register.
//
//go:linkname VrndaqF64N VrndaqF64N
//go:noescape
func VrndaqF64N(r *arm.Float64, v0 *arm.Float64, n int32)

// Floating-point Round to Integral, using current rounding mode (vector). This instruction rounds a vector of floating-point values in the SIMD&FP source register to integral floating-point values of the same size using the rounding mode that is determined by the FPCR, and writes the result to the SIMD&FP destination register.
//
//go:linkname VrndiF32N VrndiF32N
//go:noescape
func VrndiF32N(r *arm.Float32, v0 *arm.Float32, n int32)

// Floating-point Round to Integral, using current rounding mode (vector). This instruction rounds a vector of floating-point values in the SIMD&FP source register to integral floating-point values of the same size using the rounding mode that is determined by the FPCR, and writes the result to the SIMD&FP destination register.
//
//go:linkname VrndiF64N VrndiF64N
//go:noescape
func VrndiF64N(r *arm.Float64, v0 *arm.Float64, n int32)

// Floating-point Round to Integral, using current rounding mode (vector). This instruction rounds a vector of floating-point values in the SIMD&FP source register to integral floating-point values of the same size using the rounding mode that is determined by the FPCR, and writes the result to the SIMD&FP destination register.
//
//go:linkname VrndiqF32N VrndiqF32N
//go:noescape
func VrndiqF32N(r *arm.Float32, v0 *arm.Float32, n int32)

// Floating-point Round to Integral, using current rounding mode (vector). This instruction rounds a vector of floating-point values in the SIMD&FP source register to integral floating-point values of the same size using the rounding mode that is determined by the FPCR, and writes the result to the SIMD&FP destination register.
//
//go:linkname VrndiqF64N VrndiqF64N
//go:noescape
func VrndiqF64N(r *arm.Float64, v0 *arm.Float64, n int32)

// Floating-point Round to Integral, toward Minus infinity (vector). This instruction rounds a vector of floating-point values in the SIMD&FP source register to integral floating-point values of the same size using the Round towards Minus Infinity rounding mode, and writes the result to the SIMD&FP destination register.
//
//go:linkname VrndmF32N VrndmF32N
//go:noescape
func VrndmF32N(r *arm.Float32, v0 *arm.Float32, n int32)

// Floating-point Round to Integral, toward Minus infinity (vector). This instruction rounds a vector of floating-point values in the SIMD&FP source register to integral floating-point values of the same size using the Round towards Minus Infinity rounding mode, and writes the result to the SIMD&FP destination register.
//
//go:linkname VrndmF64N VrndmF64N
//go:noescape
func VrndmF64N(r *arm.Float64, v0 *arm.Float64, n int32)

// Floating-point Round to Integral, toward Minus infinity (vector). This instruction rounds a vector of floating-point values in the SIMD&FP source register to integral floating-point values of the same size using the Round towards Minus Infinity rounding mode, and writes the result to the SIMD&FP destination register.
//
//go:linkname VrndmqF32N VrndmqF32N
//go:noescape
func VrndmqF32N(r *arm.Float32, v0 *arm.Float32, n int32)

// Floating-point Round to Integral, toward Minus infinity (vector). This instruction rounds a vector of floating-point values in the SIMD&FP source register to integral floating-point values of the same size using the Round towards Minus Infinity rounding mode, and writes the result to the SIMD&FP destination register.
//
//go:linkname VrndmqF64N VrndmqF64N
//go:noescape
func VrndmqF64N(r *arm.Float64, v0 *arm.Float64, n int32)

// Floating-point Round to Integral, to nearest with ties to even (vector). This instruction rounds a vector of floating-point values in the SIMD&FP source register to integral floating-point values of the same size using the Round to Nearest rounding mode, and writes the result to the SIMD&FP destination register.
//
//go:linkname VrndnF32N VrndnF32N
//go:noescape
func VrndnF32N(r *arm.Float32, v0 *arm.Float32, n int32)

// Floating-point Round to Integral, to nearest with ties to even (vector). This instruction rounds a vector of floating-point values in the SIMD&FP source register to integral floating-point values of the same size using the Round to Nearest rounding mode, and writes the result to the SIMD&FP destination register.
//
//go:linkname VrndnF64N VrndnF64N
//go:noescape
func VrndnF64N(r *arm.Float64, v0 *arm.Float64, n int32)

// Floating-point Round to Integral, to nearest with ties to even (vector). This instruction rounds a vector of floating-point values in the SIMD&FP source register to integral floating-point values of the same size using the Round to Nearest rounding mode, and writes the result to the SIMD&FP destination register.
//
//go:linkname VrndnqF32N VrndnqF32N
//go:noescape
func VrndnqF32N(r *arm.Float32, v0 *arm.Float32, n int32)

// Floating-point Round to Integral, to nearest with ties to even (vector). This instruction rounds a vector of floating-point values in the SIMD&FP source register to integral floating-point values of the same size using the Round to Nearest rounding mode, and writes the result to the SIMD&FP destination register.
//
//go:linkname VrndnqF64N VrndnqF64N
//go:noescape
func VrndnqF64N(r *arm.Float64, v0 *arm.Float64, n int32)

// Floating-point Round to Integral, to nearest with ties to even (vector). This instruction rounds a vector of floating-point values in the SIMD&FP source register to integral floating-point values of the same size using the Round to Nearest rounding mode, and writes the result to the SIMD&FP destination register.
//
//go:linkname VrndnsF32N VrndnsF32N
//go:noescape
func VrndnsF32N(r *arm.Float32, v0 *arm.Float32, n int32)

// Floating-point Round to Integral, toward Plus infinity (vector). This instruction rounds a vector of floating-point values in the SIMD&FP source register to integral floating-point values of the same size using the Round towards Plus Infinity rounding mode, and writes the result to the SIMD&FP destination register.
//
//go:linkname VrndpF32N VrndpF32N
//go:noescape
func VrndpF32N(r *arm.Float32, v0 *arm.Float32, n int32)

// Floating-point Round to Integral, toward Plus infinity (vector). This instruction rounds a vector of floating-point values in the SIMD&FP source register to integral floating-point values of the same size using the Round towards Plus Infinity rounding mode, and writes the result to the SIMD&FP destination register.
//
//go:linkname VrndpF64N VrndpF64N
//go:noescape
func VrndpF64N(r *arm.Float64, v0 *arm.Float64, n int32)

// Floating-point Round to Integral, toward Plus infinity (vector). This instruction rounds a vector of floating-point values in the SIMD&FP source register to integral floating-point values of the same size using the Round towards Plus Infinity rounding mode, and writes the result to the SIMD&FP destination register.
//
//go:linkname VrndpqF32N VrndpqF32N
//go:noescape
func VrndpqF32N(r *arm.Float32, v0 *arm.Float32, n int32)

// Floating-point Round to Integral, toward Plus infinity (vector). This instruction rounds a vector of floating-point values in the SIMD&FP source register to integral floating-point values of the same size using the Round towards Plus Infinity rounding mode, and writes the result to the SIMD&FP destination register.
//
//go:linkname VrndpqF64N VrndpqF64N
//go:noescape
func VrndpqF64N(r *arm.Float64, v0 *arm.Float64, n int32)

// Floating-point Round to Integral, toward Zero (vector). This instruction rounds a vector of floating-point values in the SIMD&FP source register to integral floating-point values of the same size using the Round towards Zero rounding mode, and writes the result to the SIMD&FP destination register.
//
//go:linkname VrndqF32N VrndqF32N
//go:noescape
func VrndqF32N(r *arm.Float32, v0 *arm.Float32, n int32)

// Floating-point Round to Integral, toward Zero (vector). This instruction rounds a vector of floating-point values in the SIMD&FP source register to integral floating-point values of the same size using the Round towards Zero rounding mode, and writes the result to the SIMD&FP destination register.
//
//go:linkname VrndqF64N VrndqF64N
//go:noescape
func VrndqF64N(r *arm.Float64, v0 *arm.Float64, n int32)

// Floating-point Round to Integral exact, using current rounding mode (vector). This instruction rounds a vector of floating-point values in the SIMD&FP source register to integral floating-point values of the same size using the rounding mode that is determined by the FPCR, and writes the result to the SIMD&FP destination register.
//
//go:linkname VrndxF32N VrndxF32N
//go:noescape
func VrndxF32N(r *arm.Float32, v0 *arm.Float32, n int32)

// Floating-point Round to Integral exact, using current rounding mode (vector). This instruction rounds a vector of floating-point values in the SIMD&FP source register to integral floating-point values of the same size using the rounding mode that is determined by the FPCR, and writes the result to the SIMD&FP destination register.
//
//go:linkname VrndxF64N VrndxF64N
//go:noescape
func VrndxF64N(r *arm.Float64, v0 *arm.Float64, n int32)

// Floating-point Round to Integral exact, using current rounding mode (vector). This instruction rounds a vector of floating-point values in the SIMD&FP source register to integral floating-point values of the same size using the rounding mode that is determined by the FPCR, and writes the result to the SIMD&FP destination register.
//
//go:linkname VrndxqF32N VrndxqF32N
//go:noescape
func VrndxqF32N(r *arm.Float32, v0 *arm.Float32, n int32)

// Floating-point Round to Integral exact, using current rounding mode (vector). This instruction rounds a vector of floating-point values in the SIMD&FP source register to integral floating-point values of the same size using the rounding mode that is determined by the FPCR, and writes the result to the SIMD&FP destination register.
//
//go:linkname VrndxqF64N VrndxqF64N
//go:noescape
func VrndxqF64N(r *arm.Float64, v0 *arm.Float64, n int32)

// Signed Rounding Shift Left (register). This instruction takes each signed integer value in the vector of the first source SIMD&FP register, shifts it by a value from the least significant byte of the corresponding element of the second source SIMD&FP register, places the results in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VrshlS8N VrshlS8N
//go:noescape
func VrshlS8N(r *arm.Int8, v0 *arm.Int8, v1 *arm.Int8, n int32)

// Signed Rounding Shift Left (register). This instruction takes each signed integer value in the vector of the first source SIMD&FP register, shifts it by a value from the least significant byte of the corresponding element of the second source SIMD&FP register, places the results in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VrshlS16N VrshlS16N
//go:noescape
func VrshlS16N(r *arm.Int16, v0 *arm.Int16, v1 *arm.Int16, n int32)

// Signed Rounding Shift Left (register). This instruction takes each signed integer value in the vector of the first source SIMD&FP register, shifts it by a value from the least significant byte of the corresponding element of the second source SIMD&FP register, places the results in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VrshlS32N VrshlS32N
//go:noescape
func VrshlS32N(r *arm.Int32, v0 *arm.Int32, v1 *arm.Int32, n int32)

// Signed Rounding Shift Left (register). This instruction takes each signed integer value in the vector of the first source SIMD&FP register, shifts it by a value from the least significant byte of the corresponding element of the second source SIMD&FP register, places the results in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VrshlS64N VrshlS64N
//go:noescape
func VrshlS64N(r *arm.Int64, v0 *arm.Int64, v1 *arm.Int64, n int32)

// Signed Rounding Shift Left (register). This instruction takes each signed integer value in the vector of the first source SIMD&FP register, shifts it by a value from the least significant byte of the corresponding element of the second source SIMD&FP register, places the results in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VrshldS64N VrshldS64N
//go:noescape
func VrshldS64N(r *arm.Int64, v0 *arm.Int64, v1 *arm.Int64, n int32)

// Signed Rounding Shift Left (register). This instruction takes each signed integer value in the vector of the first source SIMD&FP register, shifts it by a value from the least significant byte of the corresponding element of the second source SIMD&FP register, places the results in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VrshlqS8N VrshlqS8N
//go:noescape
func VrshlqS8N(r *arm.Int8, v0 *arm.Int8, v1 *arm.Int8, n int32)

// Signed Rounding Shift Left (register). This instruction takes each signed integer value in the vector of the first source SIMD&FP register, shifts it by a value from the least significant byte of the corresponding element of the second source SIMD&FP register, places the results in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VrshlqS16N VrshlqS16N
//go:noescape
func VrshlqS16N(r *arm.Int16, v0 *arm.Int16, v1 *arm.Int16, n int32)

// Signed Rounding Shift Left (register). This instruction takes each signed integer value in the vector of the first source SIMD&FP register, shifts it by a value from the least significant byte of the corresponding element of the second source SIMD&FP register, places the results in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VrshlqS32N VrshlqS32N
//go:noescape
func VrshlqS32N(r *arm.Int32, v0 *arm.Int32, v1 *arm.Int32, n int32)

// Signed Rounding Shift Left (register). This instruction takes each signed integer value in the vector of the first source SIMD&FP register, shifts it by a value from the least significant byte of the corresponding element of the second source SIMD&FP register, places the results in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VrshlqS64N VrshlqS64N
//go:noescape
func VrshlqS64N(r *arm.Int64, v0 *arm.Int64, v1 *arm.Int64, n int32)

// Unsigned Reciprocal Square Root Estimate. This instruction reads each vector element from the source SIMD&FP register, calculates an approximate inverse square root for each value, places the result into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are unsigned integer values.
//
//go:linkname VrsqrteU32N VrsqrteU32N
//go:noescape
func VrsqrteU32N(r *arm.Uint32, v0 *arm.Uint32, n int32)

// Floating-point Reciprocal Square Root Estimate. This instruction calculates an approximate square root for each vector element in the source SIMD&FP register, places the result in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VrsqrteF32N VrsqrteF32N
//go:noescape
func VrsqrteF32N(r *arm.Float32, v0 *arm.Float32, n int32)

// Floating-point Reciprocal Square Root Estimate. This instruction calculates an approximate square root for each vector element in the source SIMD&FP register, places the result in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VrsqrteF64N VrsqrteF64N
//go:noescape
func VrsqrteF64N(r *arm.Float64, v0 *arm.Float64, n int32)

// Floating-point Reciprocal Square Root Estimate. This instruction calculates an approximate square root for each vector element in the source SIMD&FP register, places the result in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VrsqrtedF64N VrsqrtedF64N
//go:noescape
func VrsqrtedF64N(r *arm.Float64, v0 *arm.Float64, n int32)

// Unsigned Reciprocal Square Root Estimate. This instruction reads each vector element from the source SIMD&FP register, calculates an approximate inverse square root for each value, places the result into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are unsigned integer values.
//
//go:linkname VrsqrteqU32N VrsqrteqU32N
//go:noescape
func VrsqrteqU32N(r *arm.Uint32, v0 *arm.Uint32, n int32)

// Floating-point Reciprocal Square Root Estimate. This instruction calculates an approximate square root for each vector element in the source SIMD&FP register, places the result in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VrsqrteqF32N VrsqrteqF32N
//go:noescape
func VrsqrteqF32N(r *arm.Float32, v0 *arm.Float32, n int32)

// Floating-point Reciprocal Square Root Estimate. This instruction calculates an approximate square root for each vector element in the source SIMD&FP register, places the result in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VrsqrteqF64N VrsqrteqF64N
//go:noescape
func VrsqrteqF64N(r *arm.Float64, v0 *arm.Float64, n int32)

// Floating-point Reciprocal Square Root Estimate. This instruction calculates an approximate square root for each vector element in the source SIMD&FP register, places the result in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VrsqrtesF32N VrsqrtesF32N
//go:noescape
func VrsqrtesF32N(r *arm.Float32, v0 *arm.Float32, n int32)

// Floating-point Reciprocal Square Root Step. This instruction multiplies corresponding floating-point values in the vectors of the two source SIMD&FP registers, subtracts each of the products from 3.0, divides these results by 2.0, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VrsqrtsF32N VrsqrtsF32N
//go:noescape
func VrsqrtsF32N(r *arm.Float32, v0 *arm.Float32, v1 *arm.Float32, n int32)

// Floating-point Reciprocal Square Root Step. This instruction multiplies corresponding floating-point values in the vectors of the two source SIMD&FP registers, subtracts each of the products from 3.0, divides these results by 2.0, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VrsqrtsF64N VrsqrtsF64N
//go:noescape
func VrsqrtsF64N(r *arm.Float64, v0 *arm.Float64, v1 *arm.Float64, n int32)

// Floating-point Reciprocal Square Root Step. This instruction multiplies corresponding floating-point values in the vectors of the two source SIMD&FP registers, subtracts each of the products from 3.0, divides these results by 2.0, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VrsqrtsdF64N VrsqrtsdF64N
//go:noescape
func VrsqrtsdF64N(r *arm.Float64, v0 *arm.Float64, v1 *arm.Float64, n int32)

// Floating-point Reciprocal Square Root Step. This instruction multiplies corresponding floating-point values in the vectors of the two source SIMD&FP registers, subtracts each of the products from 3.0, divides these results by 2.0, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VrsqrtsqF32N VrsqrtsqF32N
//go:noescape
func VrsqrtsqF32N(r *arm.Float32, v0 *arm.Float32, v1 *arm.Float32, n int32)

// Floating-point Reciprocal Square Root Step. This instruction multiplies corresponding floating-point values in the vectors of the two source SIMD&FP registers, subtracts each of the products from 3.0, divides these results by 2.0, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VrsqrtsqF64N VrsqrtsqF64N
//go:noescape
func VrsqrtsqF64N(r *arm.Float64, v0 *arm.Float64, v1 *arm.Float64, n int32)

// Floating-point Reciprocal Square Root Step. This instruction multiplies corresponding floating-point values in the vectors of the two source SIMD&FP registers, subtracts each of the products from 3.0, divides these results by 2.0, places the results into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VrsqrtssF32N VrsqrtssF32N
//go:noescape
func VrsqrtssF32N(r *arm.Float32, v0 *arm.Float32, v1 *arm.Float32, n int32)

// SHA1 fixed rotate.
//
//go:linkname Vsha1HU32N Vsha1HU32N
//go:noescape
func Vsha1HU32N(r *arm.Uint32, v0 *arm.Uint32, n int32)

// SHA1 schedule update 1.
//
//go:linkname Vsha1Su1QU32N Vsha1Su1QU32N
//go:noescape
func Vsha1Su1QU32N(r *arm.Uint32, v0 *arm.Uint32, v1 *arm.Uint32, n int32)

// SHA256 schedule update 0.
//
//go:linkname Vsha256Su0QU32N Vsha256Su0QU32N
//go:noescape
func Vsha256Su0QU32N(r *arm.Uint32, v0 *arm.Uint32, v1 *arm.Uint32, n int32)

// SHA512 Schedule Update 0 takes the values from the two 128-bit source SIMD&FP registers and produces a 128-bit output value that combines the gamma0 functions of two iterations of the SHA512 schedule update that are performed after the first 16 iterations within a block. It returns this value to the destination SIMD&FP register.
//
//go:linkname Vsha512Su0QU64N Vsha512Su0QU64N
//go:noescape
func Vsha512Su0QU64N(r *arm.Uint64, v0 *arm.Uint64, v1 *arm.Uint64, n int32)

// Signed Shift Left (register). This instruction takes each signed integer value in the vector of the first source SIMD&FP register, shifts each value by a value from the least significant byte of the corresponding element of the second source SIMD&FP register, places the results in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VshlS8N VshlS8N
//go:noescape
func VshlS8N(r *arm.Int8, v0 *arm.Int8, v1 *arm.Int8, n int32)

// Signed Shift Left (register). This instruction takes each signed integer value in the vector of the first source SIMD&FP register, shifts each value by a value from the least significant byte of the corresponding element of the second source SIMD&FP register, places the results in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VshlS16N VshlS16N
//go:noescape
func VshlS16N(r *arm.Int16, v0 *arm.Int16, v1 *arm.Int16, n int32)

// Signed Shift Left (register). This instruction takes each signed integer value in the vector of the first source SIMD&FP register, shifts each value by a value from the least significant byte of the corresponding element of the second source SIMD&FP register, places the results in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VshlS32N VshlS32N
//go:noescape
func VshlS32N(r *arm.Int32, v0 *arm.Int32, v1 *arm.Int32, n int32)

// Signed Shift Left (register). This instruction takes each signed integer value in the vector of the first source SIMD&FP register, shifts each value by a value from the least significant byte of the corresponding element of the second source SIMD&FP register, places the results in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VshlS64N VshlS64N
//go:noescape
func VshlS64N(r *arm.Int64, v0 *arm.Int64, v1 *arm.Int64, n int32)

// Signed Shift Left (register). This instruction takes each signed integer value in the vector of the first source SIMD&FP register, shifts each value by a value from the least significant byte of the corresponding element of the second source SIMD&FP register, places the results in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VshldS64N VshldS64N
//go:noescape
func VshldS64N(r *arm.Int64, v0 *arm.Int64, v1 *arm.Int64, n int32)

// Signed Shift Left (register). This instruction takes each signed integer value in the vector of the first source SIMD&FP register, shifts each value by a value from the least significant byte of the corresponding element of the second source SIMD&FP register, places the results in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VshlqS8N VshlqS8N
//go:noescape
func VshlqS8N(r *arm.Int8, v0 *arm.Int8, v1 *arm.Int8, n int32)

// Signed Shift Left (register). This instruction takes each signed integer value in the vector of the first source SIMD&FP register, shifts each value by a value from the least significant byte of the corresponding element of the second source SIMD&FP register, places the results in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VshlqS16N VshlqS16N
//go:noescape
func VshlqS16N(r *arm.Int16, v0 *arm.Int16, v1 *arm.Int16, n int32)

// Signed Shift Left (register). This instruction takes each signed integer value in the vector of the first source SIMD&FP register, shifts each value by a value from the least significant byte of the corresponding element of the second source SIMD&FP register, places the results in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VshlqS32N VshlqS32N
//go:noescape
func VshlqS32N(r *arm.Int32, v0 *arm.Int32, v1 *arm.Int32, n int32)

// Signed Shift Left (register). This instruction takes each signed integer value in the vector of the first source SIMD&FP register, shifts each value by a value from the least significant byte of the corresponding element of the second source SIMD&FP register, places the results in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VshlqS64N VshlqS64N
//go:noescape
func VshlqS64N(r *arm.Int64, v0 *arm.Int64, v1 *arm.Int64, n int32)

// SM4 Key takes an input as a 128-bit vector from the first source SIMD&FP register and a 128-bit constant from the second SIMD&FP register. It derives four iterations of the output key, in accordance with the SM4 standard, returning the 128-bit result to the destination SIMD&FP register.
//
//go:linkname Vsm4EkeyqU32N Vsm4EkeyqU32N
//go:noescape
func Vsm4EkeyqU32N(r *arm.Uint32, v0 *arm.Uint32, v1 *arm.Uint32, n int32)

// SM4 Encode takes input data as a 128-bit vector from the first source SIMD&FP register, and four iterations of the round key held as the elements of the 128-bit vector in the second source SIMD&FP register. It encrypts the data by four rounds, in accordance with the SM4 standard, returning the 128-bit result to the destination SIMD&FP register.
//
//go:linkname Vsm4EqU32N Vsm4EqU32N
//go:noescape
func Vsm4EqU32N(r *arm.Uint32, v0 *arm.Uint32, v1 *arm.Uint32, n int32)

// Floating-point Square Root (vector). This instruction calculates the square root for each vector element in the source SIMD&FP register, places the result in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VsqrtF32N VsqrtF32N
//go:noescape
func VsqrtF32N(r *arm.Float32, v0 *arm.Float32, n int32)

// Floating-point Square Root (vector). This instruction calculates the square root for each vector element in the source SIMD&FP register, places the result in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VsqrtF64N VsqrtF64N
//go:noescape
func VsqrtF64N(r *arm.Float64, v0 *arm.Float64, n int32)

// Floating-point Square Root (vector). This instruction calculates the square root for each vector element in the source SIMD&FP register, places the result in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VsqrtqF32N VsqrtqF32N
//go:noescape
func VsqrtqF32N(r *arm.Float32, v0 *arm.Float32, n int32)

// Floating-point Square Root (vector). This instruction calculates the square root for each vector element in the source SIMD&FP register, places the result in a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VsqrtqF64N VsqrtqF64N
//go:noescape
func VsqrtqF64N(r *arm.Float64, v0 *arm.Float64, n int32)

// Subtract (vector). This instruction subtracts each vector element in the second source SIMD&FP register from the corresponding vector element in the first source SIMD&FP register, places the result into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VsubS8N VsubS8N
//go:noescape
func VsubS8N(r *arm.Int8, v0 *arm.Int8, v1 *arm.Int8, n int32)

// Subtract (vector). This instruction subtracts each vector element in the second source SIMD&FP register from the corresponding vector element in the first source SIMD&FP register, places the result into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VsubS16N VsubS16N
//go:noescape
func VsubS16N(r *arm.Int16, v0 *arm.Int16, v1 *arm.Int16, n int32)

// Subtract (vector). This instruction subtracts each vector element in the second source SIMD&FP register from the corresponding vector element in the first source SIMD&FP register, places the result into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VsubS32N VsubS32N
//go:noescape
func VsubS32N(r *arm.Int32, v0 *arm.Int32, v1 *arm.Int32, n int32)

// Subtract (vector). This instruction subtracts each vector element in the second source SIMD&FP register from the corresponding vector element in the first source SIMD&FP register, places the result into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VsubS64N VsubS64N
//go:noescape
func VsubS64N(r *arm.Int64, v0 *arm.Int64, v1 *arm.Int64, n int32)

// Subtract (vector). This instruction subtracts each vector element in the second source SIMD&FP register from the corresponding vector element in the first source SIMD&FP register, places the result into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VsubU8N VsubU8N
//go:noescape
func VsubU8N(r *arm.Uint8, v0 *arm.Uint8, v1 *arm.Uint8, n int32)

// Subtract (vector). This instruction subtracts each vector element in the second source SIMD&FP register from the corresponding vector element in the first source SIMD&FP register, places the result into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VsubU16N VsubU16N
//go:noescape
func VsubU16N(r *arm.Uint16, v0 *arm.Uint16, v1 *arm.Uint16, n int32)

// Subtract (vector). This instruction subtracts each vector element in the second source SIMD&FP register from the corresponding vector element in the first source SIMD&FP register, places the result into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VsubU32N VsubU32N
//go:noescape
func VsubU32N(r *arm.Uint32, v0 *arm.Uint32, v1 *arm.Uint32, n int32)

// Subtract (vector). This instruction subtracts each vector element in the second source SIMD&FP register from the corresponding vector element in the first source SIMD&FP register, places the result into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VsubU64N VsubU64N
//go:noescape
func VsubU64N(r *arm.Uint64, v0 *arm.Uint64, v1 *arm.Uint64, n int32)

// Floating-point Subtract (vector). This instruction subtracts the elements in the vector in the second source SIMD&FP register, from the corresponding elements in the vector in the first source SIMD&FP register, places each result into elements of a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VsubF32N VsubF32N
//go:noescape
func VsubF32N(r *arm.Float32, v0 *arm.Float32, v1 *arm.Float32, n int32)

// Floating-point Subtract (vector). This instruction subtracts the elements in the vector in the second source SIMD&FP register, from the corresponding elements in the vector in the first source SIMD&FP register, places each result into elements of a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VsubF64N VsubF64N
//go:noescape
func VsubF64N(r *arm.Float64, v0 *arm.Float64, v1 *arm.Float64, n int32)

// Subtract (vector). This instruction subtracts each vector element in the second source SIMD&FP register from the corresponding vector element in the first source SIMD&FP register, places the result into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VsubdS64N VsubdS64N
//go:noescape
func VsubdS64N(r *arm.Int64, v0 *arm.Int64, v1 *arm.Int64, n int32)

// Subtract (vector). This instruction subtracts each vector element in the second source SIMD&FP register from the corresponding vector element in the first source SIMD&FP register, places the result into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VsubdU64N VsubdU64N
//go:noescape
func VsubdU64N(r *arm.Uint64, v0 *arm.Uint64, v1 *arm.Uint64, n int32)

// Subtract (vector). This instruction subtracts each vector element in the second source SIMD&FP register from the corresponding vector element in the first source SIMD&FP register, places the result into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VsubqS8N VsubqS8N
//go:noescape
func VsubqS8N(r *arm.Int8, v0 *arm.Int8, v1 *arm.Int8, n int32)

// Subtract (vector). This instruction subtracts each vector element in the second source SIMD&FP register from the corresponding vector element in the first source SIMD&FP register, places the result into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VsubqS16N VsubqS16N
//go:noescape
func VsubqS16N(r *arm.Int16, v0 *arm.Int16, v1 *arm.Int16, n int32)

// Subtract (vector). This instruction subtracts each vector element in the second source SIMD&FP register from the corresponding vector element in the first source SIMD&FP register, places the result into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VsubqS32N VsubqS32N
//go:noescape
func VsubqS32N(r *arm.Int32, v0 *arm.Int32, v1 *arm.Int32, n int32)

// Subtract (vector). This instruction subtracts each vector element in the second source SIMD&FP register from the corresponding vector element in the first source SIMD&FP register, places the result into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VsubqS64N VsubqS64N
//go:noescape
func VsubqS64N(r *arm.Int64, v0 *arm.Int64, v1 *arm.Int64, n int32)

// Subtract (vector). This instruction subtracts each vector element in the second source SIMD&FP register from the corresponding vector element in the first source SIMD&FP register, places the result into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VsubqU8N VsubqU8N
//go:noescape
func VsubqU8N(r *arm.Uint8, v0 *arm.Uint8, v1 *arm.Uint8, n int32)

// Subtract (vector). This instruction subtracts each vector element in the second source SIMD&FP register from the corresponding vector element in the first source SIMD&FP register, places the result into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VsubqU16N VsubqU16N
//go:noescape
func VsubqU16N(r *arm.Uint16, v0 *arm.Uint16, v1 *arm.Uint16, n int32)

// Subtract (vector). This instruction subtracts each vector element in the second source SIMD&FP register from the corresponding vector element in the first source SIMD&FP register, places the result into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VsubqU32N VsubqU32N
//go:noescape
func VsubqU32N(r *arm.Uint32, v0 *arm.Uint32, v1 *arm.Uint32, n int32)

// Subtract (vector). This instruction subtracts each vector element in the second source SIMD&FP register from the corresponding vector element in the first source SIMD&FP register, places the result into a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VsubqU64N VsubqU64N
//go:noescape
func VsubqU64N(r *arm.Uint64, v0 *arm.Uint64, v1 *arm.Uint64, n int32)

// Floating-point Subtract (vector). This instruction subtracts the elements in the vector in the second source SIMD&FP register, from the corresponding elements in the vector in the first source SIMD&FP register, places each result into elements of a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VsubqF32N VsubqF32N
//go:noescape
func VsubqF32N(r *arm.Float32, v0 *arm.Float32, v1 *arm.Float32, n int32)

// Floating-point Subtract (vector). This instruction subtracts the elements in the vector in the second source SIMD&FP register, from the corresponding elements in the vector in the first source SIMD&FP register, places each result into elements of a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname VsubqF64N VsubqF64N
//go:noescape
func VsubqF64N(r *arm.Float64, v0 *arm.Float64, v1 *arm.Float64, n int32)

// Table vector Lookup. This instruction reads each value from the vector elements in the index source SIMD&FP register, uses each result as an index to perform a lookup in a table of bytes that is described by one to four source table SIMD&FP registers, places the lookup result in a vector, and writes the vector to the destination SIMD&FP register. If an index is out of range for the table, the result for that lookup is 0. If more than one source register is used to describe the table, the first source register describes the lowest bytes of the table.
//
//go:linkname Vtbl1S8N Vtbl1S8N
//go:noescape
func Vtbl1S8N(r *arm.Int8, v0 *arm.Int8, v1 *arm.Int8, n int32)

// Table vector Lookup. This instruction reads each value from the vector elements in the index source SIMD&FP register, uses each result as an index to perform a lookup in a table of bytes that is described by one to four source table SIMD&FP registers, places the lookup result in a vector, and writes the vector to the destination SIMD&FP register. If an index is out of range for the table, the result for that lookup is 0. If more than one source register is used to describe the table, the first source register describes the lowest bytes of the table.
//
//go:linkname Vtbl1U8N Vtbl1U8N
//go:noescape
func Vtbl1U8N(r *arm.Uint8, v0 *arm.Uint8, v1 *arm.Uint8, n int32)

// Transpose vectors (primary). This instruction reads corresponding even-numbered vector elements from the two source SIMD&FP registers, starting at zero, places each result into consecutive elements of a vector, and writes the vector to the destination SIMD&FP register. Vector elements from the first source register are placed into even-numbered elements of the destination vector, starting at zero, while vector elements from the second source register are placed into odd-numbered elements of the destination vector.
//
//go:linkname Vtrn1S8N Vtrn1S8N
//go:noescape
func Vtrn1S8N(r *arm.Int8, v0 *arm.Int8, v1 *arm.Int8, n int32)

// Transpose vectors (primary). This instruction reads corresponding even-numbered vector elements from the two source SIMD&FP registers, starting at zero, places each result into consecutive elements of a vector, and writes the vector to the destination SIMD&FP register. Vector elements from the first source register are placed into even-numbered elements of the destination vector, starting at zero, while vector elements from the second source register are placed into odd-numbered elements of the destination vector.
//
//go:linkname Vtrn1S16N Vtrn1S16N
//go:noescape
func Vtrn1S16N(r *arm.Int16, v0 *arm.Int16, v1 *arm.Int16, n int32)

// Transpose vectors (primary). This instruction reads corresponding even-numbered vector elements from the two source SIMD&FP registers, starting at zero, places each result into consecutive elements of a vector, and writes the vector to the destination SIMD&FP register. Vector elements from the first source register are placed into even-numbered elements of the destination vector, starting at zero, while vector elements from the second source register are placed into odd-numbered elements of the destination vector.
//
//go:linkname Vtrn1S32N Vtrn1S32N
//go:noescape
func Vtrn1S32N(r *arm.Int32, v0 *arm.Int32, v1 *arm.Int32, n int32)

// Transpose vectors (primary). This instruction reads corresponding even-numbered vector elements from the two source SIMD&FP registers, starting at zero, places each result into consecutive elements of a vector, and writes the vector to the destination SIMD&FP register. Vector elements from the first source register are placed into even-numbered elements of the destination vector, starting at zero, while vector elements from the second source register are placed into odd-numbered elements of the destination vector.
//
//go:linkname Vtrn1U8N Vtrn1U8N
//go:noescape
func Vtrn1U8N(r *arm.Uint8, v0 *arm.Uint8, v1 *arm.Uint8, n int32)

// Transpose vectors (primary). This instruction reads corresponding even-numbered vector elements from the two source SIMD&FP registers, starting at zero, places each result into consecutive elements of a vector, and writes the vector to the destination SIMD&FP register. Vector elements from the first source register are placed into even-numbered elements of the destination vector, starting at zero, while vector elements from the second source register are placed into odd-numbered elements of the destination vector.
//
//go:linkname Vtrn1U16N Vtrn1U16N
//go:noescape
func Vtrn1U16N(r *arm.Uint16, v0 *arm.Uint16, v1 *arm.Uint16, n int32)

// Transpose vectors (primary). This instruction reads corresponding even-numbered vector elements from the two source SIMD&FP registers, starting at zero, places each result into consecutive elements of a vector, and writes the vector to the destination SIMD&FP register. Vector elements from the first source register are placed into even-numbered elements of the destination vector, starting at zero, while vector elements from the second source register are placed into odd-numbered elements of the destination vector.
//
//go:linkname Vtrn1U32N Vtrn1U32N
//go:noescape
func Vtrn1U32N(r *arm.Uint32, v0 *arm.Uint32, v1 *arm.Uint32, n int32)

// Transpose vectors (primary). This instruction reads corresponding even-numbered vector elements from the two source SIMD&FP registers, starting at zero, places each result into consecutive elements of a vector, and writes the vector to the destination SIMD&FP register. Vector elements from the first source register are placed into even-numbered elements of the destination vector, starting at zero, while vector elements from the second source register are placed into odd-numbered elements of the destination vector.
//
//go:linkname Vtrn1F32N Vtrn1F32N
//go:noescape
func Vtrn1F32N(r *arm.Float32, v0 *arm.Float32, v1 *arm.Float32, n int32)

// Transpose vectors (primary). This instruction reads corresponding even-numbered vector elements from the two source SIMD&FP registers, starting at zero, places each result into consecutive elements of a vector, and writes the vector to the destination SIMD&FP register. Vector elements from the first source register are placed into even-numbered elements of the destination vector, starting at zero, while vector elements from the second source register are placed into odd-numbered elements of the destination vector.
//
//go:linkname Vtrn1QS8N Vtrn1QS8N
//go:noescape
func Vtrn1QS8N(r *arm.Int8, v0 *arm.Int8, v1 *arm.Int8, n int32)

// Transpose vectors (primary). This instruction reads corresponding even-numbered vector elements from the two source SIMD&FP registers, starting at zero, places each result into consecutive elements of a vector, and writes the vector to the destination SIMD&FP register. Vector elements from the first source register are placed into even-numbered elements of the destination vector, starting at zero, while vector elements from the second source register are placed into odd-numbered elements of the destination vector.
//
//go:linkname Vtrn1QS16N Vtrn1QS16N
//go:noescape
func Vtrn1QS16N(r *arm.Int16, v0 *arm.Int16, v1 *arm.Int16, n int32)

// Transpose vectors (primary). This instruction reads corresponding even-numbered vector elements from the two source SIMD&FP registers, starting at zero, places each result into consecutive elements of a vector, and writes the vector to the destination SIMD&FP register. Vector elements from the first source register are placed into even-numbered elements of the destination vector, starting at zero, while vector elements from the second source register are placed into odd-numbered elements of the destination vector.
//
//go:linkname Vtrn1QS32N Vtrn1QS32N
//go:noescape
func Vtrn1QS32N(r *arm.Int32, v0 *arm.Int32, v1 *arm.Int32, n int32)

// Transpose vectors (primary). This instruction reads corresponding even-numbered vector elements from the two source SIMD&FP registers, starting at zero, places each result into consecutive elements of a vector, and writes the vector to the destination SIMD&FP register. Vector elements from the first source register are placed into even-numbered elements of the destination vector, starting at zero, while vector elements from the second source register are placed into odd-numbered elements of the destination vector.
//
//go:linkname Vtrn1QS64N Vtrn1QS64N
//go:noescape
func Vtrn1QS64N(r *arm.Int64, v0 *arm.Int64, v1 *arm.Int64, n int32)

// Transpose vectors (primary). This instruction reads corresponding even-numbered vector elements from the two source SIMD&FP registers, starting at zero, places each result into consecutive elements of a vector, and writes the vector to the destination SIMD&FP register. Vector elements from the first source register are placed into even-numbered elements of the destination vector, starting at zero, while vector elements from the second source register are placed into odd-numbered elements of the destination vector.
//
//go:linkname Vtrn1QU8N Vtrn1QU8N
//go:noescape
func Vtrn1QU8N(r *arm.Uint8, v0 *arm.Uint8, v1 *arm.Uint8, n int32)

// Transpose vectors (primary). This instruction reads corresponding even-numbered vector elements from the two source SIMD&FP registers, starting at zero, places each result into consecutive elements of a vector, and writes the vector to the destination SIMD&FP register. Vector elements from the first source register are placed into even-numbered elements of the destination vector, starting at zero, while vector elements from the second source register are placed into odd-numbered elements of the destination vector.
//
//go:linkname Vtrn1QU16N Vtrn1QU16N
//go:noescape
func Vtrn1QU16N(r *arm.Uint16, v0 *arm.Uint16, v1 *arm.Uint16, n int32)

// Transpose vectors (primary). This instruction reads corresponding even-numbered vector elements from the two source SIMD&FP registers, starting at zero, places each result into consecutive elements of a vector, and writes the vector to the destination SIMD&FP register. Vector elements from the first source register are placed into even-numbered elements of the destination vector, starting at zero, while vector elements from the second source register are placed into odd-numbered elements of the destination vector.
//
//go:linkname Vtrn1QU32N Vtrn1QU32N
//go:noescape
func Vtrn1QU32N(r *arm.Uint32, v0 *arm.Uint32, v1 *arm.Uint32, n int32)

// Transpose vectors (primary). This instruction reads corresponding even-numbered vector elements from the two source SIMD&FP registers, starting at zero, places each result into consecutive elements of a vector, and writes the vector to the destination SIMD&FP register. Vector elements from the first source register are placed into even-numbered elements of the destination vector, starting at zero, while vector elements from the second source register are placed into odd-numbered elements of the destination vector.
//
//go:linkname Vtrn1QU64N Vtrn1QU64N
//go:noescape
func Vtrn1QU64N(r *arm.Uint64, v0 *arm.Uint64, v1 *arm.Uint64, n int32)

// Transpose vectors (primary). This instruction reads corresponding even-numbered vector elements from the two source SIMD&FP registers, starting at zero, places each result into consecutive elements of a vector, and writes the vector to the destination SIMD&FP register. Vector elements from the first source register are placed into even-numbered elements of the destination vector, starting at zero, while vector elements from the second source register are placed into odd-numbered elements of the destination vector.
//
//go:linkname Vtrn1QF32N Vtrn1QF32N
//go:noescape
func Vtrn1QF32N(r *arm.Float32, v0 *arm.Float32, v1 *arm.Float32, n int32)

// Transpose vectors (primary). This instruction reads corresponding even-numbered vector elements from the two source SIMD&FP registers, starting at zero, places each result into consecutive elements of a vector, and writes the vector to the destination SIMD&FP register. Vector elements from the first source register are placed into even-numbered elements of the destination vector, starting at zero, while vector elements from the second source register are placed into odd-numbered elements of the destination vector.
//
//go:linkname Vtrn1QF64N Vtrn1QF64N
//go:noescape
func Vtrn1QF64N(r *arm.Float64, v0 *arm.Float64, v1 *arm.Float64, n int32)

// Transpose vectors (secondary). This instruction reads corresponding odd-numbered vector elements from the two source SIMD&FP registers, places each result into consecutive elements of a vector, and writes the vector to the destination SIMD&FP register. Vector elements from the first source register are placed into even-numbered elements of the destination vector, starting at zero, while vector elements from the second source register are placed into odd-numbered elements of the destination vector.
//
//go:linkname Vtrn2S8N Vtrn2S8N
//go:noescape
func Vtrn2S8N(r *arm.Int8, v0 *arm.Int8, v1 *arm.Int8, n int32)

// Transpose vectors (secondary). This instruction reads corresponding odd-numbered vector elements from the two source SIMD&FP registers, places each result into consecutive elements of a vector, and writes the vector to the destination SIMD&FP register. Vector elements from the first source register are placed into even-numbered elements of the destination vector, starting at zero, while vector elements from the second source register are placed into odd-numbered elements of the destination vector.
//
//go:linkname Vtrn2S16N Vtrn2S16N
//go:noescape
func Vtrn2S16N(r *arm.Int16, v0 *arm.Int16, v1 *arm.Int16, n int32)

// Transpose vectors (secondary). This instruction reads corresponding odd-numbered vector elements from the two source SIMD&FP registers, places each result into consecutive elements of a vector, and writes the vector to the destination SIMD&FP register. Vector elements from the first source register are placed into even-numbered elements of the destination vector, starting at zero, while vector elements from the second source register are placed into odd-numbered elements of the destination vector.
//
//go:linkname Vtrn2S32N Vtrn2S32N
//go:noescape
func Vtrn2S32N(r *arm.Int32, v0 *arm.Int32, v1 *arm.Int32, n int32)

// Transpose vectors (secondary). This instruction reads corresponding odd-numbered vector elements from the two source SIMD&FP registers, places each result into consecutive elements of a vector, and writes the vector to the destination SIMD&FP register. Vector elements from the first source register are placed into even-numbered elements of the destination vector, starting at zero, while vector elements from the second source register are placed into odd-numbered elements of the destination vector.
//
//go:linkname Vtrn2U8N Vtrn2U8N
//go:noescape
func Vtrn2U8N(r *arm.Uint8, v0 *arm.Uint8, v1 *arm.Uint8, n int32)

// Transpose vectors (secondary). This instruction reads corresponding odd-numbered vector elements from the two source SIMD&FP registers, places each result into consecutive elements of a vector, and writes the vector to the destination SIMD&FP register. Vector elements from the first source register are placed into even-numbered elements of the destination vector, starting at zero, while vector elements from the second source register are placed into odd-numbered elements of the destination vector.
//
//go:linkname Vtrn2U16N Vtrn2U16N
//go:noescape
func Vtrn2U16N(r *arm.Uint16, v0 *arm.Uint16, v1 *arm.Uint16, n int32)

// Transpose vectors (secondary). This instruction reads corresponding odd-numbered vector elements from the two source SIMD&FP registers, places each result into consecutive elements of a vector, and writes the vector to the destination SIMD&FP register. Vector elements from the first source register are placed into even-numbered elements of the destination vector, starting at zero, while vector elements from the second source register are placed into odd-numbered elements of the destination vector.
//
//go:linkname Vtrn2U32N Vtrn2U32N
//go:noescape
func Vtrn2U32N(r *arm.Uint32, v0 *arm.Uint32, v1 *arm.Uint32, n int32)

// Transpose vectors (secondary). This instruction reads corresponding odd-numbered vector elements from the two source SIMD&FP registers, places each result into consecutive elements of a vector, and writes the vector to the destination SIMD&FP register. Vector elements from the first source register are placed into even-numbered elements of the destination vector, starting at zero, while vector elements from the second source register are placed into odd-numbered elements of the destination vector.
//
//go:linkname Vtrn2F32N Vtrn2F32N
//go:noescape
func Vtrn2F32N(r *arm.Float32, v0 *arm.Float32, v1 *arm.Float32, n int32)

// Transpose vectors (secondary). This instruction reads corresponding odd-numbered vector elements from the two source SIMD&FP registers, places each result into consecutive elements of a vector, and writes the vector to the destination SIMD&FP register. Vector elements from the first source register are placed into even-numbered elements of the destination vector, starting at zero, while vector elements from the second source register are placed into odd-numbered elements of the destination vector.
//
//go:linkname Vtrn2QS8N Vtrn2QS8N
//go:noescape
func Vtrn2QS8N(r *arm.Int8, v0 *arm.Int8, v1 *arm.Int8, n int32)

// Transpose vectors (secondary). This instruction reads corresponding odd-numbered vector elements from the two source SIMD&FP registers, places each result into consecutive elements of a vector, and writes the vector to the destination SIMD&FP register. Vector elements from the first source register are placed into even-numbered elements of the destination vector, starting at zero, while vector elements from the second source register are placed into odd-numbered elements of the destination vector.
//
//go:linkname Vtrn2QS16N Vtrn2QS16N
//go:noescape
func Vtrn2QS16N(r *arm.Int16, v0 *arm.Int16, v1 *arm.Int16, n int32)

// Transpose vectors (secondary). This instruction reads corresponding odd-numbered vector elements from the two source SIMD&FP registers, places each result into consecutive elements of a vector, and writes the vector to the destination SIMD&FP register. Vector elements from the first source register are placed into even-numbered elements of the destination vector, starting at zero, while vector elements from the second source register are placed into odd-numbered elements of the destination vector.
//
//go:linkname Vtrn2QS32N Vtrn2QS32N
//go:noescape
func Vtrn2QS32N(r *arm.Int32, v0 *arm.Int32, v1 *arm.Int32, n int32)

// Transpose vectors (secondary). This instruction reads corresponding odd-numbered vector elements from the two source SIMD&FP registers, places each result into consecutive elements of a vector, and writes the vector to the destination SIMD&FP register. Vector elements from the first source register are placed into even-numbered elements of the destination vector, starting at zero, while vector elements from the second source register are placed into odd-numbered elements of the destination vector.
//
//go:linkname Vtrn2QS64N Vtrn2QS64N
//go:noescape
func Vtrn2QS64N(r *arm.Int64, v0 *arm.Int64, v1 *arm.Int64, n int32)

// Transpose vectors (secondary). This instruction reads corresponding odd-numbered vector elements from the two source SIMD&FP registers, places each result into consecutive elements of a vector, and writes the vector to the destination SIMD&FP register. Vector elements from the first source register are placed into even-numbered elements of the destination vector, starting at zero, while vector elements from the second source register are placed into odd-numbered elements of the destination vector.
//
//go:linkname Vtrn2QU8N Vtrn2QU8N
//go:noescape
func Vtrn2QU8N(r *arm.Uint8, v0 *arm.Uint8, v1 *arm.Uint8, n int32)

// Transpose vectors (secondary). This instruction reads corresponding odd-numbered vector elements from the two source SIMD&FP registers, places each result into consecutive elements of a vector, and writes the vector to the destination SIMD&FP register. Vector elements from the first source register are placed into even-numbered elements of the destination vector, starting at zero, while vector elements from the second source register are placed into odd-numbered elements of the destination vector.
//
//go:linkname Vtrn2QU16N Vtrn2QU16N
//go:noescape
func Vtrn2QU16N(r *arm.Uint16, v0 *arm.Uint16, v1 *arm.Uint16, n int32)

// Transpose vectors (secondary). This instruction reads corresponding odd-numbered vector elements from the two source SIMD&FP registers, places each result into consecutive elements of a vector, and writes the vector to the destination SIMD&FP register. Vector elements from the first source register are placed into even-numbered elements of the destination vector, starting at zero, while vector elements from the second source register are placed into odd-numbered elements of the destination vector.
//
//go:linkname Vtrn2QU32N Vtrn2QU32N
//go:noescape
func Vtrn2QU32N(r *arm.Uint32, v0 *arm.Uint32, v1 *arm.Uint32, n int32)

// Transpose vectors (secondary). This instruction reads corresponding odd-numbered vector elements from the two source SIMD&FP registers, places each result into consecutive elements of a vector, and writes the vector to the destination SIMD&FP register. Vector elements from the first source register are placed into even-numbered elements of the destination vector, starting at zero, while vector elements from the second source register are placed into odd-numbered elements of the destination vector.
//
//go:linkname Vtrn2QU64N Vtrn2QU64N
//go:noescape
func Vtrn2QU64N(r *arm.Uint64, v0 *arm.Uint64, v1 *arm.Uint64, n int32)

// Transpose vectors (secondary). This instruction reads corresponding odd-numbered vector elements from the two source SIMD&FP registers, places each result into consecutive elements of a vector, and writes the vector to the destination SIMD&FP register. Vector elements from the first source register are placed into even-numbered elements of the destination vector, starting at zero, while vector elements from the second source register are placed into odd-numbered elements of the destination vector.
//
//go:linkname Vtrn2QF32N Vtrn2QF32N
//go:noescape
func Vtrn2QF32N(r *arm.Float32, v0 *arm.Float32, v1 *arm.Float32, n int32)

// Transpose vectors (secondary). This instruction reads corresponding odd-numbered vector elements from the two source SIMD&FP registers, places each result into consecutive elements of a vector, and writes the vector to the destination SIMD&FP register. Vector elements from the first source register are placed into even-numbered elements of the destination vector, starting at zero, while vector elements from the second source register are placed into odd-numbered elements of the destination vector.
//
//go:linkname Vtrn2QF64N Vtrn2QF64N
//go:noescape
func Vtrn2QF64N(r *arm.Float64, v0 *arm.Float64, v1 *arm.Float64, n int32)

// Compare bitwise Test bits nonzero (vector). This instruction reads each vector element in the first source SIMD&FP register, performs an AND with the corresponding vector element in the second source SIMD&FP register, and if the result is not zero, sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VtstS8N VtstS8N
//go:noescape
func VtstS8N(r *arm.Uint8, v0 *arm.Int8, v1 *arm.Int8, n int32)

// Compare bitwise Test bits nonzero (vector). This instruction reads each vector element in the first source SIMD&FP register, performs an AND with the corresponding vector element in the second source SIMD&FP register, and if the result is not zero, sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VtstS16N VtstS16N
//go:noescape
func VtstS16N(r *arm.Uint16, v0 *arm.Int16, v1 *arm.Int16, n int32)

// Compare bitwise Test bits nonzero (vector). This instruction reads each vector element in the first source SIMD&FP register, performs an AND with the corresponding vector element in the second source SIMD&FP register, and if the result is not zero, sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VtstS32N VtstS32N
//go:noescape
func VtstS32N(r *arm.Uint32, v0 *arm.Int32, v1 *arm.Int32, n int32)

// Compare bitwise Test bits nonzero (vector). This instruction reads each vector element in the first source SIMD&FP register, performs an AND with the corresponding vector element in the second source SIMD&FP register, and if the result is not zero, sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VtstS64N VtstS64N
//go:noescape
func VtstS64N(r *arm.Uint64, v0 *arm.Int64, v1 *arm.Int64, n int32)

// Compare bitwise Test bits nonzero (vector). This instruction reads each vector element in the first source SIMD&FP register, performs an AND with the corresponding vector element in the second source SIMD&FP register, and if the result is not zero, sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VtstU8N VtstU8N
//go:noescape
func VtstU8N(r *arm.Uint8, v0 *arm.Uint8, v1 *arm.Uint8, n int32)

// Compare bitwise Test bits nonzero (vector). This instruction reads each vector element in the first source SIMD&FP register, performs an AND with the corresponding vector element in the second source SIMD&FP register, and if the result is not zero, sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VtstU16N VtstU16N
//go:noescape
func VtstU16N(r *arm.Uint16, v0 *arm.Uint16, v1 *arm.Uint16, n int32)

// Compare bitwise Test bits nonzero (vector). This instruction reads each vector element in the first source SIMD&FP register, performs an AND with the corresponding vector element in the second source SIMD&FP register, and if the result is not zero, sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VtstU32N VtstU32N
//go:noescape
func VtstU32N(r *arm.Uint32, v0 *arm.Uint32, v1 *arm.Uint32, n int32)

// Compare bitwise Test bits nonzero (vector). This instruction reads each vector element in the first source SIMD&FP register, performs an AND with the corresponding vector element in the second source SIMD&FP register, and if the result is not zero, sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VtstU64N VtstU64N
//go:noescape
func VtstU64N(r *arm.Uint64, v0 *arm.Uint64, v1 *arm.Uint64, n int32)

// Compare bitwise Test bits nonzero (vector). This instruction reads each vector element in the first source SIMD&FP register, performs an AND with the corresponding vector element in the second source SIMD&FP register, and if the result is not zero, sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VtstdS64N VtstdS64N
//go:noescape
func VtstdS64N(r *arm.Uint64, v0 *arm.Int64, v1 *arm.Int64, n int32)

// Compare bitwise Test bits nonzero (vector). This instruction reads each vector element in the first source SIMD&FP register, performs an AND with the corresponding vector element in the second source SIMD&FP register, and if the result is not zero, sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VtstdU64N VtstdU64N
//go:noescape
func VtstdU64N(r *arm.Uint64, v0 *arm.Uint64, v1 *arm.Uint64, n int32)

// Compare bitwise Test bits nonzero (vector). This instruction reads each vector element in the first source SIMD&FP register, performs an AND with the corresponding vector element in the second source SIMD&FP register, and if the result is not zero, sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VtstqS8N VtstqS8N
//go:noescape
func VtstqS8N(r *arm.Uint8, v0 *arm.Int8, v1 *arm.Int8, n int32)

// Compare bitwise Test bits nonzero (vector). This instruction reads each vector element in the first source SIMD&FP register, performs an AND with the corresponding vector element in the second source SIMD&FP register, and if the result is not zero, sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VtstqS16N VtstqS16N
//go:noescape
func VtstqS16N(r *arm.Uint16, v0 *arm.Int16, v1 *arm.Int16, n int32)

// Compare bitwise Test bits nonzero (vector). This instruction reads each vector element in the first source SIMD&FP register, performs an AND with the corresponding vector element in the second source SIMD&FP register, and if the result is not zero, sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VtstqS32N VtstqS32N
//go:noescape
func VtstqS32N(r *arm.Uint32, v0 *arm.Int32, v1 *arm.Int32, n int32)

// Compare bitwise Test bits nonzero (vector). This instruction reads each vector element in the first source SIMD&FP register, performs an AND with the corresponding vector element in the second source SIMD&FP register, and if the result is not zero, sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VtstqS64N VtstqS64N
//go:noescape
func VtstqS64N(r *arm.Uint64, v0 *arm.Int64, v1 *arm.Int64, n int32)

// Compare bitwise Test bits nonzero (vector). This instruction reads each vector element in the first source SIMD&FP register, performs an AND with the corresponding vector element in the second source SIMD&FP register, and if the result is not zero, sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VtstqU8N VtstqU8N
//go:noescape
func VtstqU8N(r *arm.Uint8, v0 *arm.Uint8, v1 *arm.Uint8, n int32)

// Compare bitwise Test bits nonzero (vector). This instruction reads each vector element in the first source SIMD&FP register, performs an AND with the corresponding vector element in the second source SIMD&FP register, and if the result is not zero, sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VtstqU16N VtstqU16N
//go:noescape
func VtstqU16N(r *arm.Uint16, v0 *arm.Uint16, v1 *arm.Uint16, n int32)

// Compare bitwise Test bits nonzero (vector). This instruction reads each vector element in the first source SIMD&FP register, performs an AND with the corresponding vector element in the second source SIMD&FP register, and if the result is not zero, sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VtstqU32N VtstqU32N
//go:noescape
func VtstqU32N(r *arm.Uint32, v0 *arm.Uint32, v1 *arm.Uint32, n int32)

// Compare bitwise Test bits nonzero (vector). This instruction reads each vector element in the first source SIMD&FP register, performs an AND with the corresponding vector element in the second source SIMD&FP register, and if the result is not zero, sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
//
//go:linkname VtstqU64N VtstqU64N
//go:noescape
func VtstqU64N(r *arm.Uint64, v0 *arm.Uint64, v1 *arm.Uint64, n int32)

// Unzip vectors (primary). This instruction reads corresponding even-numbered vector elements from the two source SIMD&FP registers, starting at zero, places the result from the first source register into consecutive elements in the lower half of a vector, and the result from the second source register into consecutive elements in the upper half of a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname Vuzp1S8N Vuzp1S8N
//go:noescape
func Vuzp1S8N(r *arm.Int8, v0 *arm.Int8, v1 *arm.Int8, n int32)

// Unzip vectors (primary). This instruction reads corresponding even-numbered vector elements from the two source SIMD&FP registers, starting at zero, places the result from the first source register into consecutive elements in the lower half of a vector, and the result from the second source register into consecutive elements in the upper half of a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname Vuzp1S16N Vuzp1S16N
//go:noescape
func Vuzp1S16N(r *arm.Int16, v0 *arm.Int16, v1 *arm.Int16, n int32)

// Unzip vectors (primary). This instruction reads corresponding even-numbered vector elements from the two source SIMD&FP registers, starting at zero, places the result from the first source register into consecutive elements in the lower half of a vector, and the result from the second source register into consecutive elements in the upper half of a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname Vuzp1S32N Vuzp1S32N
//go:noescape
func Vuzp1S32N(r *arm.Int32, v0 *arm.Int32, v1 *arm.Int32, n int32)

// Unzip vectors (primary). This instruction reads corresponding even-numbered vector elements from the two source SIMD&FP registers, starting at zero, places the result from the first source register into consecutive elements in the lower half of a vector, and the result from the second source register into consecutive elements in the upper half of a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname Vuzp1U8N Vuzp1U8N
//go:noescape
func Vuzp1U8N(r *arm.Uint8, v0 *arm.Uint8, v1 *arm.Uint8, n int32)

// Unzip vectors (primary). This instruction reads corresponding even-numbered vector elements from the two source SIMD&FP registers, starting at zero, places the result from the first source register into consecutive elements in the lower half of a vector, and the result from the second source register into consecutive elements in the upper half of a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname Vuzp1U16N Vuzp1U16N
//go:noescape
func Vuzp1U16N(r *arm.Uint16, v0 *arm.Uint16, v1 *arm.Uint16, n int32)

// Unzip vectors (primary). This instruction reads corresponding even-numbered vector elements from the two source SIMD&FP registers, starting at zero, places the result from the first source register into consecutive elements in the lower half of a vector, and the result from the second source register into consecutive elements in the upper half of a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname Vuzp1U32N Vuzp1U32N
//go:noescape
func Vuzp1U32N(r *arm.Uint32, v0 *arm.Uint32, v1 *arm.Uint32, n int32)

// Unzip vectors (primary). This instruction reads corresponding even-numbered vector elements from the two source SIMD&FP registers, starting at zero, places the result from the first source register into consecutive elements in the lower half of a vector, and the result from the second source register into consecutive elements in the upper half of a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname Vuzp1F32N Vuzp1F32N
//go:noescape
func Vuzp1F32N(r *arm.Float32, v0 *arm.Float32, v1 *arm.Float32, n int32)

// Unzip vectors (primary). This instruction reads corresponding even-numbered vector elements from the two source SIMD&FP registers, starting at zero, places the result from the first source register into consecutive elements in the lower half of a vector, and the result from the second source register into consecutive elements in the upper half of a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname Vuzp1QS8N Vuzp1QS8N
//go:noescape
func Vuzp1QS8N(r *arm.Int8, v0 *arm.Int8, v1 *arm.Int8, n int32)

// Unzip vectors (primary). This instruction reads corresponding even-numbered vector elements from the two source SIMD&FP registers, starting at zero, places the result from the first source register into consecutive elements in the lower half of a vector, and the result from the second source register into consecutive elements in the upper half of a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname Vuzp1QS16N Vuzp1QS16N
//go:noescape
func Vuzp1QS16N(r *arm.Int16, v0 *arm.Int16, v1 *arm.Int16, n int32)

// Unzip vectors (primary). This instruction reads corresponding even-numbered vector elements from the two source SIMD&FP registers, starting at zero, places the result from the first source register into consecutive elements in the lower half of a vector, and the result from the second source register into consecutive elements in the upper half of a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname Vuzp1QS32N Vuzp1QS32N
//go:noescape
func Vuzp1QS32N(r *arm.Int32, v0 *arm.Int32, v1 *arm.Int32, n int32)

// Unzip vectors (primary). This instruction reads corresponding even-numbered vector elements from the two source SIMD&FP registers, starting at zero, places the result from the first source register into consecutive elements in the lower half of a vector, and the result from the second source register into consecutive elements in the upper half of a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname Vuzp1QS64N Vuzp1QS64N
//go:noescape
func Vuzp1QS64N(r *arm.Int64, v0 *arm.Int64, v1 *arm.Int64, n int32)

// Unzip vectors (primary). This instruction reads corresponding even-numbered vector elements from the two source SIMD&FP registers, starting at zero, places the result from the first source register into consecutive elements in the lower half of a vector, and the result from the second source register into consecutive elements in the upper half of a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname Vuzp1QU8N Vuzp1QU8N
//go:noescape
func Vuzp1QU8N(r *arm.Uint8, v0 *arm.Uint8, v1 *arm.Uint8, n int32)

// Unzip vectors (primary). This instruction reads corresponding even-numbered vector elements from the two source SIMD&FP registers, starting at zero, places the result from the first source register into consecutive elements in the lower half of a vector, and the result from the second source register into consecutive elements in the upper half of a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname Vuzp1QU16N Vuzp1QU16N
//go:noescape
func Vuzp1QU16N(r *arm.Uint16, v0 *arm.Uint16, v1 *arm.Uint16, n int32)

// Unzip vectors (primary). This instruction reads corresponding even-numbered vector elements from the two source SIMD&FP registers, starting at zero, places the result from the first source register into consecutive elements in the lower half of a vector, and the result from the second source register into consecutive elements in the upper half of a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname Vuzp1QU32N Vuzp1QU32N
//go:noescape
func Vuzp1QU32N(r *arm.Uint32, v0 *arm.Uint32, v1 *arm.Uint32, n int32)

// Unzip vectors (primary). This instruction reads corresponding even-numbered vector elements from the two source SIMD&FP registers, starting at zero, places the result from the first source register into consecutive elements in the lower half of a vector, and the result from the second source register into consecutive elements in the upper half of a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname Vuzp1QU64N Vuzp1QU64N
//go:noescape
func Vuzp1QU64N(r *arm.Uint64, v0 *arm.Uint64, v1 *arm.Uint64, n int32)

// Unzip vectors (primary). This instruction reads corresponding even-numbered vector elements from the two source SIMD&FP registers, starting at zero, places the result from the first source register into consecutive elements in the lower half of a vector, and the result from the second source register into consecutive elements in the upper half of a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname Vuzp1QF32N Vuzp1QF32N
//go:noescape
func Vuzp1QF32N(r *arm.Float32, v0 *arm.Float32, v1 *arm.Float32, n int32)

// Unzip vectors (primary). This instruction reads corresponding even-numbered vector elements from the two source SIMD&FP registers, starting at zero, places the result from the first source register into consecutive elements in the lower half of a vector, and the result from the second source register into consecutive elements in the upper half of a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname Vuzp1QF64N Vuzp1QF64N
//go:noescape
func Vuzp1QF64N(r *arm.Float64, v0 *arm.Float64, v1 *arm.Float64, n int32)

// Unzip vectors (secondary). This instruction reads corresponding odd-numbered vector elements from the two source SIMD&FP registers, places the result from the first source register into consecutive elements in the lower half of a vector, and the result from the second source register into consecutive elements in the upper half of a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname Vuzp2S8N Vuzp2S8N
//go:noescape
func Vuzp2S8N(r *arm.Int8, v0 *arm.Int8, v1 *arm.Int8, n int32)

// Unzip vectors (secondary). This instruction reads corresponding odd-numbered vector elements from the two source SIMD&FP registers, places the result from the first source register into consecutive elements in the lower half of a vector, and the result from the second source register into consecutive elements in the upper half of a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname Vuzp2S16N Vuzp2S16N
//go:noescape
func Vuzp2S16N(r *arm.Int16, v0 *arm.Int16, v1 *arm.Int16, n int32)

// Unzip vectors (secondary). This instruction reads corresponding odd-numbered vector elements from the two source SIMD&FP registers, places the result from the first source register into consecutive elements in the lower half of a vector, and the result from the second source register into consecutive elements in the upper half of a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname Vuzp2S32N Vuzp2S32N
//go:noescape
func Vuzp2S32N(r *arm.Int32, v0 *arm.Int32, v1 *arm.Int32, n int32)

// Unzip vectors (secondary). This instruction reads corresponding odd-numbered vector elements from the two source SIMD&FP registers, places the result from the first source register into consecutive elements in the lower half of a vector, and the result from the second source register into consecutive elements in the upper half of a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname Vuzp2U8N Vuzp2U8N
//go:noescape
func Vuzp2U8N(r *arm.Uint8, v0 *arm.Uint8, v1 *arm.Uint8, n int32)

// Unzip vectors (secondary). This instruction reads corresponding odd-numbered vector elements from the two source SIMD&FP registers, places the result from the first source register into consecutive elements in the lower half of a vector, and the result from the second source register into consecutive elements in the upper half of a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname Vuzp2U16N Vuzp2U16N
//go:noescape
func Vuzp2U16N(r *arm.Uint16, v0 *arm.Uint16, v1 *arm.Uint16, n int32)

// Unzip vectors (secondary). This instruction reads corresponding odd-numbered vector elements from the two source SIMD&FP registers, places the result from the first source register into consecutive elements in the lower half of a vector, and the result from the second source register into consecutive elements in the upper half of a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname Vuzp2U32N Vuzp2U32N
//go:noescape
func Vuzp2U32N(r *arm.Uint32, v0 *arm.Uint32, v1 *arm.Uint32, n int32)

// Unzip vectors (secondary). This instruction reads corresponding odd-numbered vector elements from the two source SIMD&FP registers, places the result from the first source register into consecutive elements in the lower half of a vector, and the result from the second source register into consecutive elements in the upper half of a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname Vuzp2F32N Vuzp2F32N
//go:noescape
func Vuzp2F32N(r *arm.Float32, v0 *arm.Float32, v1 *arm.Float32, n int32)

// Unzip vectors (secondary). This instruction reads corresponding odd-numbered vector elements from the two source SIMD&FP registers, places the result from the first source register into consecutive elements in the lower half of a vector, and the result from the second source register into consecutive elements in the upper half of a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname Vuzp2QS8N Vuzp2QS8N
//go:noescape
func Vuzp2QS8N(r *arm.Int8, v0 *arm.Int8, v1 *arm.Int8, n int32)

// Unzip vectors (secondary). This instruction reads corresponding odd-numbered vector elements from the two source SIMD&FP registers, places the result from the first source register into consecutive elements in the lower half of a vector, and the result from the second source register into consecutive elements in the upper half of a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname Vuzp2QS16N Vuzp2QS16N
//go:noescape
func Vuzp2QS16N(r *arm.Int16, v0 *arm.Int16, v1 *arm.Int16, n int32)

// Unzip vectors (secondary). This instruction reads corresponding odd-numbered vector elements from the two source SIMD&FP registers, places the result from the first source register into consecutive elements in the lower half of a vector, and the result from the second source register into consecutive elements in the upper half of a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname Vuzp2QS32N Vuzp2QS32N
//go:noescape
func Vuzp2QS32N(r *arm.Int32, v0 *arm.Int32, v1 *arm.Int32, n int32)

// Unzip vectors (secondary). This instruction reads corresponding odd-numbered vector elements from the two source SIMD&FP registers, places the result from the first source register into consecutive elements in the lower half of a vector, and the result from the second source register into consecutive elements in the upper half of a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname Vuzp2QS64N Vuzp2QS64N
//go:noescape
func Vuzp2QS64N(r *arm.Int64, v0 *arm.Int64, v1 *arm.Int64, n int32)

// Unzip vectors (secondary). This instruction reads corresponding odd-numbered vector elements from the two source SIMD&FP registers, places the result from the first source register into consecutive elements in the lower half of a vector, and the result from the second source register into consecutive elements in the upper half of a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname Vuzp2QU8N Vuzp2QU8N
//go:noescape
func Vuzp2QU8N(r *arm.Uint8, v0 *arm.Uint8, v1 *arm.Uint8, n int32)

// Unzip vectors (secondary). This instruction reads corresponding odd-numbered vector elements from the two source SIMD&FP registers, places the result from the first source register into consecutive elements in the lower half of a vector, and the result from the second source register into consecutive elements in the upper half of a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname Vuzp2QU16N Vuzp2QU16N
//go:noescape
func Vuzp2QU16N(r *arm.Uint16, v0 *arm.Uint16, v1 *arm.Uint16, n int32)

// Unzip vectors (secondary). This instruction reads corresponding odd-numbered vector elements from the two source SIMD&FP registers, places the result from the first source register into consecutive elements in the lower half of a vector, and the result from the second source register into consecutive elements in the upper half of a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname Vuzp2QU32N Vuzp2QU32N
//go:noescape
func Vuzp2QU32N(r *arm.Uint32, v0 *arm.Uint32, v1 *arm.Uint32, n int32)

// Unzip vectors (secondary). This instruction reads corresponding odd-numbered vector elements from the two source SIMD&FP registers, places the result from the first source register into consecutive elements in the lower half of a vector, and the result from the second source register into consecutive elements in the upper half of a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname Vuzp2QU64N Vuzp2QU64N
//go:noescape
func Vuzp2QU64N(r *arm.Uint64, v0 *arm.Uint64, v1 *arm.Uint64, n int32)

// Unzip vectors (secondary). This instruction reads corresponding odd-numbered vector elements from the two source SIMD&FP registers, places the result from the first source register into consecutive elements in the lower half of a vector, and the result from the second source register into consecutive elements in the upper half of a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname Vuzp2QF32N Vuzp2QF32N
//go:noescape
func Vuzp2QF32N(r *arm.Float32, v0 *arm.Float32, v1 *arm.Float32, n int32)

// Unzip vectors (secondary). This instruction reads corresponding odd-numbered vector elements from the two source SIMD&FP registers, places the result from the first source register into consecutive elements in the lower half of a vector, and the result from the second source register into consecutive elements in the upper half of a vector, and writes the vector to the destination SIMD&FP register.
//
//go:linkname Vuzp2QF64N Vuzp2QF64N
//go:noescape
func Vuzp2QF64N(r *arm.Float64, v0 *arm.Float64, v1 *arm.Float64, n int32)

// Zip vectors (primary). This instruction reads adjacent vector elements from the lower half of two source SIMD&FP registers as pairs, interleaves the pairs and places them into a vector, and writes the vector to the destination SIMD&FP register. The first pair from the first source register is placed into the two lowest vector elements, with subsequent pairs taken alternately from each source register.
//
//go:linkname Vzip1S8N Vzip1S8N
//go:noescape
func Vzip1S8N(r *arm.Int8, v0 *arm.Int8, v1 *arm.Int8, n int32)

// Zip vectors (primary). This instruction reads adjacent vector elements from the lower half of two source SIMD&FP registers as pairs, interleaves the pairs and places them into a vector, and writes the vector to the destination SIMD&FP register. The first pair from the first source register is placed into the two lowest vector elements, with subsequent pairs taken alternately from each source register.
//
//go:linkname Vzip1S16N Vzip1S16N
//go:noescape
func Vzip1S16N(r *arm.Int16, v0 *arm.Int16, v1 *arm.Int16, n int32)

// Zip vectors (primary). This instruction reads adjacent vector elements from the lower half of two source SIMD&FP registers as pairs, interleaves the pairs and places them into a vector, and writes the vector to the destination SIMD&FP register. The first pair from the first source register is placed into the two lowest vector elements, with subsequent pairs taken alternately from each source register.
//
//go:linkname Vzip1S32N Vzip1S32N
//go:noescape
func Vzip1S32N(r *arm.Int32, v0 *arm.Int32, v1 *arm.Int32, n int32)

// Zip vectors (primary). This instruction reads adjacent vector elements from the lower half of two source SIMD&FP registers as pairs, interleaves the pairs and places them into a vector, and writes the vector to the destination SIMD&FP register. The first pair from the first source register is placed into the two lowest vector elements, with subsequent pairs taken alternately from each source register.
//
//go:linkname Vzip1U8N Vzip1U8N
//go:noescape
func Vzip1U8N(r *arm.Uint8, v0 *arm.Uint8, v1 *arm.Uint8, n int32)

// Zip vectors (primary). This instruction reads adjacent vector elements from the lower half of two source SIMD&FP registers as pairs, interleaves the pairs and places them into a vector, and writes the vector to the destination SIMD&FP register. The first pair from the first source register is placed into the two lowest vector elements, with subsequent pairs taken alternately from each source register.
//
//go:linkname Vzip1U16N Vzip1U16N
//go:noescape
func Vzip1U16N(r *arm.Uint16, v0 *arm.Uint16, v1 *arm.Uint16, n int32)

// Zip vectors (primary). This instruction reads adjacent vector elements from the lower half of two source SIMD&FP registers as pairs, interleaves the pairs and places them into a vector, and writes the vector to the destination SIMD&FP register. The first pair from the first source register is placed into the two lowest vector elements, with subsequent pairs taken alternately from each source register.
//
//go:linkname Vzip1U32N Vzip1U32N
//go:noescape
func Vzip1U32N(r *arm.Uint32, v0 *arm.Uint32, v1 *arm.Uint32, n int32)

// Zip vectors (primary). This instruction reads adjacent vector elements from the lower half of two source SIMD&FP registers as pairs, interleaves the pairs and places them into a vector, and writes the vector to the destination SIMD&FP register. The first pair from the first source register is placed into the two lowest vector elements, with subsequent pairs taken alternately from each source register.
//
//go:linkname Vzip1F32N Vzip1F32N
//go:noescape
func Vzip1F32N(r *arm.Float32, v0 *arm.Float32, v1 *arm.Float32, n int32)

// Zip vectors (primary). This instruction reads adjacent vector elements from the lower half of two source SIMD&FP registers as pairs, interleaves the pairs and places them into a vector, and writes the vector to the destination SIMD&FP register. The first pair from the first source register is placed into the two lowest vector elements, with subsequent pairs taken alternately from each source register.
//
//go:linkname Vzip1QS8N Vzip1QS8N
//go:noescape
func Vzip1QS8N(r *arm.Int8, v0 *arm.Int8, v1 *arm.Int8, n int32)

// Zip vectors (primary). This instruction reads adjacent vector elements from the lower half of two source SIMD&FP registers as pairs, interleaves the pairs and places them into a vector, and writes the vector to the destination SIMD&FP register. The first pair from the first source register is placed into the two lowest vector elements, with subsequent pairs taken alternately from each source register.
//
//go:linkname Vzip1QS16N Vzip1QS16N
//go:noescape
func Vzip1QS16N(r *arm.Int16, v0 *arm.Int16, v1 *arm.Int16, n int32)

// Zip vectors (primary). This instruction reads adjacent vector elements from the lower half of two source SIMD&FP registers as pairs, interleaves the pairs and places them into a vector, and writes the vector to the destination SIMD&FP register. The first pair from the first source register is placed into the two lowest vector elements, with subsequent pairs taken alternately from each source register.
//
//go:linkname Vzip1QS32N Vzip1QS32N
//go:noescape
func Vzip1QS32N(r *arm.Int32, v0 *arm.Int32, v1 *arm.Int32, n int32)

// Zip vectors (primary). This instruction reads adjacent vector elements from the lower half of two source SIMD&FP registers as pairs, interleaves the pairs and places them into a vector, and writes the vector to the destination SIMD&FP register. The first pair from the first source register is placed into the two lowest vector elements, with subsequent pairs taken alternately from each source register.
//
//go:linkname Vzip1QS64N Vzip1QS64N
//go:noescape
func Vzip1QS64N(r *arm.Int64, v0 *arm.Int64, v1 *arm.Int64, n int32)

// Zip vectors (primary). This instruction reads adjacent vector elements from the lower half of two source SIMD&FP registers as pairs, interleaves the pairs and places them into a vector, and writes the vector to the destination SIMD&FP register. The first pair from the first source register is placed into the two lowest vector elements, with subsequent pairs taken alternately from each source register.
//
//go:linkname Vzip1QU8N Vzip1QU8N
//go:noescape
func Vzip1QU8N(r *arm.Uint8, v0 *arm.Uint8, v1 *arm.Uint8, n int32)

// Zip vectors (primary). This instruction reads adjacent vector elements from the lower half of two source SIMD&FP registers as pairs, interleaves the pairs and places them into a vector, and writes the vector to the destination SIMD&FP register. The first pair from the first source register is placed into the two lowest vector elements, with subsequent pairs taken alternately from each source register.
//
//go:linkname Vzip1QU16N Vzip1QU16N
//go:noescape
func Vzip1QU16N(r *arm.Uint16, v0 *arm.Uint16, v1 *arm.Uint16, n int32)

// Zip vectors (primary). This instruction reads adjacent vector elements from the lower half of two source SIMD&FP registers as pairs, interleaves the pairs and places them into a vector, and writes the vector to the destination SIMD&FP register. The first pair from the first source register is placed into the two lowest vector elements, with subsequent pairs taken alternately from each source register.
//
//go:linkname Vzip1QU32N Vzip1QU32N
//go:noescape
func Vzip1QU32N(r *arm.Uint32, v0 *arm.Uint32, v1 *arm.Uint32, n int32)

// Zip vectors (primary). This instruction reads adjacent vector elements from the lower half of two source SIMD&FP registers as pairs, interleaves the pairs and places them into a vector, and writes the vector to the destination SIMD&FP register. The first pair from the first source register is placed into the two lowest vector elements, with subsequent pairs taken alternately from each source register.
//
//go:linkname Vzip1QU64N Vzip1QU64N
//go:noescape
func Vzip1QU64N(r *arm.Uint64, v0 *arm.Uint64, v1 *arm.Uint64, n int32)

// Zip vectors (primary). This instruction reads adjacent vector elements from the lower half of two source SIMD&FP registers as pairs, interleaves the pairs and places them into a vector, and writes the vector to the destination SIMD&FP register. The first pair from the first source register is placed into the two lowest vector elements, with subsequent pairs taken alternately from each source register.
//
//go:linkname Vzip1QF32N Vzip1QF32N
//go:noescape
func Vzip1QF32N(r *arm.Float32, v0 *arm.Float32, v1 *arm.Float32, n int32)

// Zip vectors (primary). This instruction reads adjacent vector elements from the lower half of two source SIMD&FP registers as pairs, interleaves the pairs and places them into a vector, and writes the vector to the destination SIMD&FP register. The first pair from the first source register is placed into the two lowest vector elements, with subsequent pairs taken alternately from each source register.
//
//go:linkname Vzip1QF64N Vzip1QF64N
//go:noescape
func Vzip1QF64N(r *arm.Float64, v0 *arm.Float64, v1 *arm.Float64, n int32)

// Zip vectors (secondary). This instruction reads adjacent vector elements from the upper half of two source SIMD&FP registers as pairs, interleaves the pairs and places them into a vector, and writes the vector to the destination SIMD&FP register. The first pair from the first source register is placed into the two lowest vector elements, with subsequent pairs taken alternately from each source register.
//
//go:linkname Vzip2S8N Vzip2S8N
//go:noescape
func Vzip2S8N(r *arm.Int8, v0 *arm.Int8, v1 *arm.Int8, n int32)

// Zip vectors (secondary). This instruction reads adjacent vector elements from the upper half of two source SIMD&FP registers as pairs, interleaves the pairs and places them into a vector, and writes the vector to the destination SIMD&FP register. The first pair from the first source register is placed into the two lowest vector elements, with subsequent pairs taken alternately from each source register.
//
//go:linkname Vzip2S16N Vzip2S16N
//go:noescape
func Vzip2S16N(r *arm.Int16, v0 *arm.Int16, v1 *arm.Int16, n int32)

// Zip vectors (secondary). This instruction reads adjacent vector elements from the upper half of two source SIMD&FP registers as pairs, interleaves the pairs and places them into a vector, and writes the vector to the destination SIMD&FP register. The first pair from the first source register is placed into the two lowest vector elements, with subsequent pairs taken alternately from each source register.
//
//go:linkname Vzip2S32N Vzip2S32N
//go:noescape
func Vzip2S32N(r *arm.Int32, v0 *arm.Int32, v1 *arm.Int32, n int32)

// Zip vectors (secondary). This instruction reads adjacent vector elements from the upper half of two source SIMD&FP registers as pairs, interleaves the pairs and places them into a vector, and writes the vector to the destination SIMD&FP register. The first pair from the first source register is placed into the two lowest vector elements, with subsequent pairs taken alternately from each source register.
//
//go:linkname Vzip2U8N Vzip2U8N
//go:noescape
func Vzip2U8N(r *arm.Uint8, v0 *arm.Uint8, v1 *arm.Uint8, n int32)

// Zip vectors (secondary). This instruction reads adjacent vector elements from the upper half of two source SIMD&FP registers as pairs, interleaves the pairs and places them into a vector, and writes the vector to the destination SIMD&FP register. The first pair from the first source register is placed into the two lowest vector elements, with subsequent pairs taken alternately from each source register.
//
//go:linkname Vzip2U16N Vzip2U16N
//go:noescape
func Vzip2U16N(r *arm.Uint16, v0 *arm.Uint16, v1 *arm.Uint16, n int32)

// Zip vectors (secondary). This instruction reads adjacent vector elements from the upper half of two source SIMD&FP registers as pairs, interleaves the pairs and places them into a vector, and writes the vector to the destination SIMD&FP register. The first pair from the first source register is placed into the two lowest vector elements, with subsequent pairs taken alternately from each source register.
//
//go:linkname Vzip2U32N Vzip2U32N
//go:noescape
func Vzip2U32N(r *arm.Uint32, v0 *arm.Uint32, v1 *arm.Uint32, n int32)

// Zip vectors (secondary). This instruction reads adjacent vector elements from the upper half of two source SIMD&FP registers as pairs, interleaves the pairs and places them into a vector, and writes the vector to the destination SIMD&FP register. The first pair from the first source register is placed into the two lowest vector elements, with subsequent pairs taken alternately from each source register.
//
//go:linkname Vzip2F32N Vzip2F32N
//go:noescape
func Vzip2F32N(r *arm.Float32, v0 *arm.Float32, v1 *arm.Float32, n int32)

// Zip vectors (secondary). This instruction reads adjacent vector elements from the upper half of two source SIMD&FP registers as pairs, interleaves the pairs and places them into a vector, and writes the vector to the destination SIMD&FP register. The first pair from the first source register is placed into the two lowest vector elements, with subsequent pairs taken alternately from each source register.
//
//go:linkname Vzip2QS8N Vzip2QS8N
//go:noescape
func Vzip2QS8N(r *arm.Int8, v0 *arm.Int8, v1 *arm.Int8, n int32)

// Zip vectors (secondary). This instruction reads adjacent vector elements from the upper half of two source SIMD&FP registers as pairs, interleaves the pairs and places them into a vector, and writes the vector to the destination SIMD&FP register. The first pair from the first source register is placed into the two lowest vector elements, with subsequent pairs taken alternately from each source register.
//
//go:linkname Vzip2QS16N Vzip2QS16N
//go:noescape
func Vzip2QS16N(r *arm.Int16, v0 *arm.Int16, v1 *arm.Int16, n int32)

// Zip vectors (secondary). This instruction reads adjacent vector elements from the upper half of two source SIMD&FP registers as pairs, interleaves the pairs and places them into a vector, and writes the vector to the destination SIMD&FP register. The first pair from the first source register is placed into the two lowest vector elements, with subsequent pairs taken alternately from each source register.
//
//go:linkname Vzip2QS32N Vzip2QS32N
//go:noescape
func Vzip2QS32N(r *arm.Int32, v0 *arm.Int32, v1 *arm.Int32, n int32)

// Zip vectors (secondary). This instruction reads adjacent vector elements from the upper half of two source SIMD&FP registers as pairs, interleaves the pairs and places them into a vector, and writes the vector to the destination SIMD&FP register. The first pair from the first source register is placed into the two lowest vector elements, with subsequent pairs taken alternately from each source register.
//
//go:linkname Vzip2QS64N Vzip2QS64N
//go:noescape
func Vzip2QS64N(r *arm.Int64, v0 *arm.Int64, v1 *arm.Int64, n int32)

// Zip vectors (secondary). This instruction reads adjacent vector elements from the upper half of two source SIMD&FP registers as pairs, interleaves the pairs and places them into a vector, and writes the vector to the destination SIMD&FP register. The first pair from the first source register is placed into the two lowest vector elements, with subsequent pairs taken alternately from each source register.
//
//go:linkname Vzip2QU8N Vzip2QU8N
//go:noescape
func Vzip2QU8N(r *arm.Uint8, v0 *arm.Uint8, v1 *arm.Uint8, n int32)

// Zip vectors (secondary). This instruction reads adjacent vector elements from the upper half of two source SIMD&FP registers as pairs, interleaves the pairs and places them into a vector, and writes the vector to the destination SIMD&FP register. The first pair from the first source register is placed into the two lowest vector elements, with subsequent pairs taken alternately from each source register.
//
//go:linkname Vzip2QU16N Vzip2QU16N
//go:noescape
func Vzip2QU16N(r *arm.Uint16, v0 *arm.Uint16, v1 *arm.Uint16, n int32)

// Zip vectors (secondary). This instruction reads adjacent vector elements from the upper half of two source SIMD&FP registers as pairs, interleaves the pairs and places them into a vector, and writes the vector to the destination SIMD&FP register. The first pair from the first source register is placed into the two lowest vector elements, with subsequent pairs taken alternately from each source register.
//
//go:linkname Vzip2QU32N Vzip2QU32N
//go:noescape
func Vzip2QU32N(r *arm.Uint32, v0 *arm.Uint32, v1 *arm.Uint32, n int32)

// Zip vectors (secondary). This instruction reads adjacent vector elements from the upper half of two source SIMD&FP registers as pairs, interleaves the pairs and places them into a vector, and writes the vector to the destination SIMD&FP register. The first pair from the first source register is placed into the two lowest vector elements, with subsequent pairs taken alternately from each source register.
//
//go:linkname Vzip2QU64N Vzip2QU64N
//go:noescape
func Vzip2QU64N(r *arm.Uint64, v0 *arm.Uint64, v1 *arm.Uint64, n int32)

// Zip vectors (secondary). This instruction reads adjacent vector elements from the upper half of two source SIMD&FP registers as pairs, interleaves the pairs and places them into a vector, and writes the vector to the destination SIMD&FP register. The first pair from the first source register is placed into the two lowest vector elements, with subsequent pairs taken alternately from each source register.
//
//go:linkname Vzip2QF32N Vzip2QF32N
//go:noescape
func Vzip2QF32N(r *arm.Float32, v0 *arm.Float32, v1 *arm.Float32, n int32)

// Zip vectors (secondary). This instruction reads adjacent vector elements from the upper half of two source SIMD&FP registers as pairs, interleaves the pairs and places them into a vector, and writes the vector to the destination SIMD&FP register. The first pair from the first source register is placed into the two lowest vector elements, with subsequent pairs taken alternately from each source register.
//
//go:linkname Vzip2QF64N Vzip2QF64N
//go:noescape
func Vzip2QF64N(r *arm.Float64, v0 *arm.Float64, v1 *arm.Float64, n int32)
