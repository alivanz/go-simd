package neon

/*
#cgo CFLAGS: -arch arm64
#include <arm_neon.h>
*/
import "C"

// Unsigned Absolute Difference (vector). This instruction subtracts the elements of the vector of the second source SIMD&FP register from the corresponding elements of the first source SIMD&FP register, places the the absolute values of the results into a vector, and writes the vector to the destination SIMD&FP register.
func VabdqU8(v0 Uint8X16, v1 Uint8X16) Uint8X16 {
	return C.vabdq_u8(v0, v1)
}

// Unsigned Absolute Difference (vector). This instruction subtracts the elements of the vector of the second source SIMD&FP register from the corresponding elements of the first source SIMD&FP register, places the the absolute values of the results into a vector, and writes the vector to the destination SIMD&FP register.
func VabdqU32(v0 Uint32X4, v1 Uint32X4) Uint32X4 {
	return C.vabdq_u32(v0, v1)
}

// Unsigned Absolute Difference (vector). This instruction subtracts the elements of the vector of the second source SIMD&FP register from the corresponding elements of the first source SIMD&FP register, places the the absolute values of the results into a vector, and writes the vector to the destination SIMD&FP register.
func VabdqU16(v0 Uint16X8, v1 Uint16X8) Uint16X8 {
	return C.vabdq_u16(v0, v1)
}

// Signed Absolute Difference. This instruction subtracts the elements of the vector of the second source SIMD&FP register from the corresponding elements of the first source SIMD&FP register, places the the absolute values of the results into a vector, and writes the vector to the destination SIMD&FP register.
func VabdqS8(v0 Int8X16, v1 Int8X16) Int8X16 {
	return C.vabdq_s8(v0, v1)
}

// Floating-point Absolute Difference (vector). This instruction subtracts the floating-point values in the elements of the second source SIMD&FP register, from the corresponding floating-point values in the elements of the first source SIMD&FP register, places the absolute value of each result in a vector, and writes the vector to the destination SIMD&FP register.
func VabdqF32(v0 Float32X4, v1 Float32X4) Float32X4 {
	return C.vabdq_f32(v0, v1)
}

// Signed Absolute Difference. This instruction subtracts the elements of the vector of the second source SIMD&FP register from the corresponding elements of the first source SIMD&FP register, places the the absolute values of the results into a vector, and writes the vector to the destination SIMD&FP register.
func VabdqS32(v0 Int32X4, v1 Int32X4) Int32X4 {
	return C.vabdq_s32(v0, v1)
}

// Signed Absolute Difference. This instruction subtracts the elements of the vector of the second source SIMD&FP register from the corresponding elements of the first source SIMD&FP register, places the the absolute values of the results into a vector, and writes the vector to the destination SIMD&FP register.
func VabdqS16(v0 Int16X8, v1 Int16X8) Int16X8 {
	return C.vabdq_s16(v0, v1)
}

// Unsigned Absolute Difference (vector). This instruction subtracts the elements of the vector of the second source SIMD&FP register from the corresponding elements of the first source SIMD&FP register, places the the absolute values of the results into a vector, and writes the vector to the destination SIMD&FP register.
func VabdU8(v0 Uint8X8, v1 Uint8X8) Uint8X8 {
	return C.vabd_u8(v0, v1)
}

// Unsigned Absolute Difference (vector). This instruction subtracts the elements of the vector of the second source SIMD&FP register from the corresponding elements of the first source SIMD&FP register, places the the absolute values of the results into a vector, and writes the vector to the destination SIMD&FP register.
func VabdU32(v0 Uint32X2, v1 Uint32X2) Uint32X2 {
	return C.vabd_u32(v0, v1)
}

// Unsigned Absolute Difference (vector). This instruction subtracts the elements of the vector of the second source SIMD&FP register from the corresponding elements of the first source SIMD&FP register, places the the absolute values of the results into a vector, and writes the vector to the destination SIMD&FP register.
func VabdU16(v0 Uint16X4, v1 Uint16X4) Uint16X4 {
	return C.vabd_u16(v0, v1)
}

// Signed Absolute Difference. This instruction subtracts the elements of the vector of the second source SIMD&FP register from the corresponding elements of the first source SIMD&FP register, places the the absolute values of the results into a vector, and writes the vector to the destination SIMD&FP register.
func VabdS8(v0 Int8X8, v1 Int8X8) Int8X8 {
	return C.vabd_s8(v0, v1)
}

// Floating-point Absolute Difference (vector). This instruction subtracts the floating-point values in the elements of the second source SIMD&FP register, from the corresponding floating-point values in the elements of the first source SIMD&FP register, places the absolute value of each result in a vector, and writes the vector to the destination SIMD&FP register.
func VabdF32(v0 Float32X2, v1 Float32X2) Float32X2 {
	return C.vabd_f32(v0, v1)
}

// Signed Absolute Difference. This instruction subtracts the elements of the vector of the second source SIMD&FP register from the corresponding elements of the first source SIMD&FP register, places the the absolute values of the results into a vector, and writes the vector to the destination SIMD&FP register.
func VabdS32(v0 Int32X2, v1 Int32X2) Int32X2 {
	return C.vabd_s32(v0, v1)
}

// Signed Absolute Difference. This instruction subtracts the elements of the vector of the second source SIMD&FP register from the corresponding elements of the first source SIMD&FP register, places the the absolute values of the results into a vector, and writes the vector to the destination SIMD&FP register.
func VabdS16(v0 Int16X4, v1 Int16X4) Int16X4 {
	return C.vabd_s16(v0, v1)
}

// Absolute value (vector). This instruction calculates the absolute value of each vector element in the source SIMD&FP register, puts the result into a vector, and writes the vector to the destination SIMD&FP register.
func VabsqS8(v0 Int8X16) Int8X16 {
	return C.vabsq_s8(v0)
}

// Floating-point Absolute value (vector). This instruction calculates the absolute value of each vector element in the source SIMD&FP register, writes the result to a vector, and writes the vector to the destination SIMD&FP register.
func VabsqF32(v0 Float32X4) Float32X4 {
	return C.vabsq_f32(v0)
}

// Absolute value (vector). This instruction calculates the absolute value of each vector element in the source SIMD&FP register, puts the result into a vector, and writes the vector to the destination SIMD&FP register.
func VabsqS32(v0 Int32X4) Int32X4 {
	return C.vabsq_s32(v0)
}

// Absolute value (vector). This instruction calculates the absolute value of each vector element in the source SIMD&FP register, puts the result into a vector, and writes the vector to the destination SIMD&FP register.
func VabsqS16(v0 Int16X8) Int16X8 {
	return C.vabsq_s16(v0)
}

// Absolute value (vector). This instruction calculates the absolute value of each vector element in the source SIMD&FP register, puts the result into a vector, and writes the vector to the destination SIMD&FP register.
func VabsS8(v0 Int8X8) Int8X8 {
	return C.vabs_s8(v0)
}

// Floating-point Absolute value (vector). This instruction calculates the absolute value of each vector element in the source SIMD&FP register, writes the result to a vector, and writes the vector to the destination SIMD&FP register.
func VabsF32(v0 Float32X2) Float32X2 {
	return C.vabs_f32(v0)
}

// Absolute value (vector). This instruction calculates the absolute value of each vector element in the source SIMD&FP register, puts the result into a vector, and writes the vector to the destination SIMD&FP register.
func VabsS32(v0 Int32X2) Int32X2 {
	return C.vabs_s32(v0)
}

// Absolute value (vector). This instruction calculates the absolute value of each vector element in the source SIMD&FP register, puts the result into a vector, and writes the vector to the destination SIMD&FP register.
func VabsS16(v0 Int16X4) Int16X4 {
	return C.vabs_s16(v0)
}

// Add (vector). This instruction adds corresponding elements in the two source SIMD&FP registers, places the results into a vector, and writes the vector to the destination SIMD&FP register.
func VaddqU8(v0 Uint8X16, v1 Uint8X16) Uint8X16 {
	return C.vaddq_u8(v0, v1)
}

// Add (vector). This instruction adds corresponding elements in the two source SIMD&FP registers, places the results into a vector, and writes the vector to the destination SIMD&FP register.
func VaddqU32(v0 Uint32X4, v1 Uint32X4) Uint32X4 {
	return C.vaddq_u32(v0, v1)
}

// Add (vector). This instruction adds corresponding elements in the two source SIMD&FP registers, places the results into a vector, and writes the vector to the destination SIMD&FP register.
func VaddqU64(v0 Uint64X2, v1 Uint64X2) Uint64X2 {
	return C.vaddq_u64(v0, v1)
}

// Add (vector). This instruction adds corresponding elements in the two source SIMD&FP registers, places the results into a vector, and writes the vector to the destination SIMD&FP register.
func VaddqU16(v0 Uint16X8, v1 Uint16X8) Uint16X8 {
	return C.vaddq_u16(v0, v1)
}

// Add (vector). This instruction adds corresponding elements in the two source SIMD&FP registers, places the results into a vector, and writes the vector to the destination SIMD&FP register.
func VaddqS8(v0 Int8X16, v1 Int8X16) Int8X16 {
	return C.vaddq_s8(v0, v1)
}

// Floating-point Add (vector). This instruction adds corresponding vector elements in the two source SIMD&FP registers, writes the result into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are floating-point values.
func VaddqF32(v0 Float32X4, v1 Float32X4) Float32X4 {
	return C.vaddq_f32(v0, v1)
}

// Add (vector). This instruction adds corresponding elements in the two source SIMD&FP registers, places the results into a vector, and writes the vector to the destination SIMD&FP register.
func VaddqS32(v0 Int32X4, v1 Int32X4) Int32X4 {
	return C.vaddq_s32(v0, v1)
}

// Add (vector). This instruction adds corresponding elements in the two source SIMD&FP registers, places the results into a vector, and writes the vector to the destination SIMD&FP register.
func VaddqS64(v0 Int64X2, v1 Int64X2) Int64X2 {
	return C.vaddq_s64(v0, v1)
}

// Add (vector). This instruction adds corresponding elements in the two source SIMD&FP registers, places the results into a vector, and writes the vector to the destination SIMD&FP register.
func VaddqS16(v0 Int16X8, v1 Int16X8) Int16X8 {
	return C.vaddq_s16(v0, v1)
}

// Add (vector). This instruction adds corresponding elements in the two source SIMD&FP registers, places the results into a vector, and writes the vector to the destination SIMD&FP register.
func VaddU8(v0 Uint8X8, v1 Uint8X8) Uint8X8 {
	return C.vadd_u8(v0, v1)
}

// Add (vector). This instruction adds corresponding elements in the two source SIMD&FP registers, places the results into a vector, and writes the vector to the destination SIMD&FP register.
func VaddU32(v0 Uint32X2, v1 Uint32X2) Uint32X2 {
	return C.vadd_u32(v0, v1)
}

// Add (vector). This instruction adds corresponding elements in the two source SIMD&FP registers, places the results into a vector, and writes the vector to the destination SIMD&FP register.
func VaddU64(v0 Uint64X1, v1 Uint64X1) Uint64X1 {
	return C.vadd_u64(v0, v1)
}

// Add (vector). This instruction adds corresponding elements in the two source SIMD&FP registers, places the results into a vector, and writes the vector to the destination SIMD&FP register.
func VaddU16(v0 Uint16X4, v1 Uint16X4) Uint16X4 {
	return C.vadd_u16(v0, v1)
}

// Add (vector). This instruction adds corresponding elements in the two source SIMD&FP registers, places the results into a vector, and writes the vector to the destination SIMD&FP register.
func VaddS8(v0 Int8X8, v1 Int8X8) Int8X8 {
	return C.vadd_s8(v0, v1)
}

// Floating-point Add (vector). This instruction adds corresponding vector elements in the two source SIMD&FP registers, writes the result into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are floating-point values.
func VaddF32(v0 Float32X2, v1 Float32X2) Float32X2 {
	return C.vadd_f32(v0, v1)
}

// Add (vector). This instruction adds corresponding elements in the two source SIMD&FP registers, places the results into a vector, and writes the vector to the destination SIMD&FP register.
func VaddS32(v0 Int32X2, v1 Int32X2) Int32X2 {
	return C.vadd_s32(v0, v1)
}

// Add (vector). This instruction adds corresponding elements in the two source SIMD&FP registers, places the results into a vector, and writes the vector to the destination SIMD&FP register.
func VaddS64(v0 Int64X1, v1 Int64X1) Int64X1 {
	return C.vadd_s64(v0, v1)
}

// Add (vector). This instruction adds corresponding elements in the two source SIMD&FP registers, places the results into a vector, and writes the vector to the destination SIMD&FP register.
func VaddS16(v0 Int16X4, v1 Int16X4) Int16X4 {
	return C.vadd_s16(v0, v1)
}

// Bitwise Exclusive OR (vector). This instruction performs a bitwise Exclusive OR operation between the two source SIMD&FP registers, and places the result in the destination SIMD&FP register.
func VaddP8(v0 Poly8X8, v1 Poly8X8) Poly8X8 {
	return C.vadd_p8(v0, v1)
}

// Bitwise Exclusive OR (vector). This instruction performs a bitwise Exclusive OR operation between the two source SIMD&FP registers, and places the result in the destination SIMD&FP register.
func VaddP64(v0 Poly64X1, v1 Poly64X1) Poly64X1 {
	return C.vadd_p64(v0, v1)
}

// Bitwise Exclusive OR (vector). This instruction performs a bitwise Exclusive OR operation between the two source SIMD&FP registers, and places the result in the destination SIMD&FP register.
func VaddP16(v0 Poly16X4, v1 Poly16X4) Poly16X4 {
	return C.vadd_p16(v0, v1)
}

// Bitwise Exclusive OR (vector). This instruction performs a bitwise Exclusive OR operation between the two source SIMD&FP registers, and places the result in the destination SIMD&FP register.
func VaddqP8(v0 Poly8X16, v1 Poly8X16) Poly8X16 {
	return C.vaddq_p8(v0, v1)
}

// Bitwise Exclusive OR (vector). This instruction performs a bitwise Exclusive OR operation between the two source SIMD&FP registers, and places the result in the destination SIMD&FP register.
func VaddqP64(v0 Poly64X2, v1 Poly64X2) Poly64X2 {
	return C.vaddq_p64(v0, v1)
}

// Bitwise Exclusive OR (vector). This instruction performs a bitwise Exclusive OR operation between the two source SIMD&FP registers, and places the result in the destination SIMD&FP register.
func VaddqP16(v0 Poly16X8, v1 Poly16X8) Poly16X8 {
	return C.vaddq_p16(v0, v1)
}

// Add returning High Narrow. This instruction adds each vector element in the first source SIMD&FP register to the corresponding vector element in the second source SIMD&FP register, places the most significant half of the result into a vector, and writes the vector to the lower or upper half of the destination SIMD&FP register.
func VaddhnU32(v0 Uint32X4, v1 Uint32X4) Uint16X4 {
	return C.vaddhn_u32(v0, v1)
}

// Add returning High Narrow. This instruction adds each vector element in the first source SIMD&FP register to the corresponding vector element in the second source SIMD&FP register, places the most significant half of the result into a vector, and writes the vector to the lower or upper half of the destination SIMD&FP register.
func VaddhnU64(v0 Uint64X2, v1 Uint64X2) Uint32X2 {
	return C.vaddhn_u64(v0, v1)
}

// Add returning High Narrow. This instruction adds each vector element in the first source SIMD&FP register to the corresponding vector element in the second source SIMD&FP register, places the most significant half of the result into a vector, and writes the vector to the lower or upper half of the destination SIMD&FP register.
func VaddhnU16(v0 Uint16X8, v1 Uint16X8) Uint8X8 {
	return C.vaddhn_u16(v0, v1)
}

// Add returning High Narrow. This instruction adds each vector element in the first source SIMD&FP register to the corresponding vector element in the second source SIMD&FP register, places the most significant half of the result into a vector, and writes the vector to the lower or upper half of the destination SIMD&FP register.
func VaddhnS32(v0 Int32X4, v1 Int32X4) Int16X4 {
	return C.vaddhn_s32(v0, v1)
}

// Add returning High Narrow. This instruction adds each vector element in the first source SIMD&FP register to the corresponding vector element in the second source SIMD&FP register, places the most significant half of the result into a vector, and writes the vector to the lower or upper half of the destination SIMD&FP register.
func VaddhnS64(v0 Int64X2, v1 Int64X2) Int32X2 {
	return C.vaddhn_s64(v0, v1)
}

// Add returning High Narrow. This instruction adds each vector element in the first source SIMD&FP register to the corresponding vector element in the second source SIMD&FP register, places the most significant half of the result into a vector, and writes the vector to the lower or upper half of the destination SIMD&FP register.
func VaddhnS16(v0 Int16X8, v1 Int16X8) Int8X8 {
	return C.vaddhn_s16(v0, v1)
}

// Bitwise AND (vector). This instruction performs a bitwise AND between the two source SIMD&FP registers, and writes the result to the destination SIMD&FP register.
func VandqU8(v0 Uint8X16, v1 Uint8X16) Uint8X16 {
	return C.vandq_u8(v0, v1)
}

// Bitwise AND (vector). This instruction performs a bitwise AND between the two source SIMD&FP registers, and writes the result to the destination SIMD&FP register.
func VandqU32(v0 Uint32X4, v1 Uint32X4) Uint32X4 {
	return C.vandq_u32(v0, v1)
}

// Bitwise AND (vector). This instruction performs a bitwise AND between the two source SIMD&FP registers, and writes the result to the destination SIMD&FP register.
func VandqU64(v0 Uint64X2, v1 Uint64X2) Uint64X2 {
	return C.vandq_u64(v0, v1)
}

// Bitwise AND (vector). This instruction performs a bitwise AND between the two source SIMD&FP registers, and writes the result to the destination SIMD&FP register.
func VandqU16(v0 Uint16X8, v1 Uint16X8) Uint16X8 {
	return C.vandq_u16(v0, v1)
}

// Bitwise AND (vector). This instruction performs a bitwise AND between the two source SIMD&FP registers, and writes the result to the destination SIMD&FP register.
func VandqS8(v0 Int8X16, v1 Int8X16) Int8X16 {
	return C.vandq_s8(v0, v1)
}

// Bitwise AND (vector). This instruction performs a bitwise AND between the two source SIMD&FP registers, and writes the result to the destination SIMD&FP register.
func VandqS32(v0 Int32X4, v1 Int32X4) Int32X4 {
	return C.vandq_s32(v0, v1)
}

// Bitwise AND (vector). This instruction performs a bitwise AND between the two source SIMD&FP registers, and writes the result to the destination SIMD&FP register.
func VandqS64(v0 Int64X2, v1 Int64X2) Int64X2 {
	return C.vandq_s64(v0, v1)
}

// Bitwise AND (vector). This instruction performs a bitwise AND between the two source SIMD&FP registers, and writes the result to the destination SIMD&FP register.
func VandqS16(v0 Int16X8, v1 Int16X8) Int16X8 {
	return C.vandq_s16(v0, v1)
}

// Bitwise AND (vector). This instruction performs a bitwise AND between the two source SIMD&FP registers, and writes the result to the destination SIMD&FP register.
func VandU8(v0 Uint8X8, v1 Uint8X8) Uint8X8 {
	return C.vand_u8(v0, v1)
}

// Bitwise AND (vector). This instruction performs a bitwise AND between the two source SIMD&FP registers, and writes the result to the destination SIMD&FP register.
func VandU32(v0 Uint32X2, v1 Uint32X2) Uint32X2 {
	return C.vand_u32(v0, v1)
}

// Bitwise AND (vector). This instruction performs a bitwise AND between the two source SIMD&FP registers, and writes the result to the destination SIMD&FP register.
func VandU64(v0 Uint64X1, v1 Uint64X1) Uint64X1 {
	return C.vand_u64(v0, v1)
}

// Bitwise AND (vector). This instruction performs a bitwise AND between the two source SIMD&FP registers, and writes the result to the destination SIMD&FP register.
func VandU16(v0 Uint16X4, v1 Uint16X4) Uint16X4 {
	return C.vand_u16(v0, v1)
}

// Bitwise AND (vector). This instruction performs a bitwise AND between the two source SIMD&FP registers, and writes the result to the destination SIMD&FP register.
func VandS8(v0 Int8X8, v1 Int8X8) Int8X8 {
	return C.vand_s8(v0, v1)
}

// Bitwise AND (vector). This instruction performs a bitwise AND between the two source SIMD&FP registers, and writes the result to the destination SIMD&FP register.
func VandS32(v0 Int32X2, v1 Int32X2) Int32X2 {
	return C.vand_s32(v0, v1)
}

// Bitwise AND (vector). This instruction performs a bitwise AND between the two source SIMD&FP registers, and writes the result to the destination SIMD&FP register.
func VandS64(v0 Int64X1, v1 Int64X1) Int64X1 {
	return C.vand_s64(v0, v1)
}

// Bitwise AND (vector). This instruction performs a bitwise AND between the two source SIMD&FP registers, and writes the result to the destination SIMD&FP register.
func VandS16(v0 Int16X4, v1 Int16X4) Int16X4 {
	return C.vand_s16(v0, v1)
}

// Bitwise bit Clear (vector, register). This instruction performs a bitwise AND between the first source SIMD&FP register and the complement of the second source SIMD&FP register, and writes the result to the destination SIMD&FP register.
func VbicqU8(v0 Uint8X16, v1 Uint8X16) Uint8X16 {
	return C.vbicq_u8(v0, v1)
}

// Bitwise bit Clear (vector, register). This instruction performs a bitwise AND between the first source SIMD&FP register and the complement of the second source SIMD&FP register, and writes the result to the destination SIMD&FP register.
func VbicqU32(v0 Uint32X4, v1 Uint32X4) Uint32X4 {
	return C.vbicq_u32(v0, v1)
}

// Bitwise bit Clear (vector, register). This instruction performs a bitwise AND between the first source SIMD&FP register and the complement of the second source SIMD&FP register, and writes the result to the destination SIMD&FP register.
func VbicqU64(v0 Uint64X2, v1 Uint64X2) Uint64X2 {
	return C.vbicq_u64(v0, v1)
}

// Bitwise bit Clear (vector, register). This instruction performs a bitwise AND between the first source SIMD&FP register and the complement of the second source SIMD&FP register, and writes the result to the destination SIMD&FP register.
func VbicqU16(v0 Uint16X8, v1 Uint16X8) Uint16X8 {
	return C.vbicq_u16(v0, v1)
}

// Bitwise bit Clear (vector, register). This instruction performs a bitwise AND between the first source SIMD&FP register and the complement of the second source SIMD&FP register, and writes the result to the destination SIMD&FP register.
func VbicqS8(v0 Int8X16, v1 Int8X16) Int8X16 {
	return C.vbicq_s8(v0, v1)
}

// Bitwise bit Clear (vector, register). This instruction performs a bitwise AND between the first source SIMD&FP register and the complement of the second source SIMD&FP register, and writes the result to the destination SIMD&FP register.
func VbicqS32(v0 Int32X4, v1 Int32X4) Int32X4 {
	return C.vbicq_s32(v0, v1)
}

// Bitwise bit Clear (vector, register). This instruction performs a bitwise AND between the first source SIMD&FP register and the complement of the second source SIMD&FP register, and writes the result to the destination SIMD&FP register.
func VbicqS64(v0 Int64X2, v1 Int64X2) Int64X2 {
	return C.vbicq_s64(v0, v1)
}

// Bitwise bit Clear (vector, register). This instruction performs a bitwise AND between the first source SIMD&FP register and the complement of the second source SIMD&FP register, and writes the result to the destination SIMD&FP register.
func VbicqS16(v0 Int16X8, v1 Int16X8) Int16X8 {
	return C.vbicq_s16(v0, v1)
}

// Bitwise bit Clear (vector, register). This instruction performs a bitwise AND between the first source SIMD&FP register and the complement of the second source SIMD&FP register, and writes the result to the destination SIMD&FP register.
func VbicU8(v0 Uint8X8, v1 Uint8X8) Uint8X8 {
	return C.vbic_u8(v0, v1)
}

// Bitwise bit Clear (vector, register). This instruction performs a bitwise AND between the first source SIMD&FP register and the complement of the second source SIMD&FP register, and writes the result to the destination SIMD&FP register.
func VbicU32(v0 Uint32X2, v1 Uint32X2) Uint32X2 {
	return C.vbic_u32(v0, v1)
}

// Bitwise bit Clear (vector, register). This instruction performs a bitwise AND between the first source SIMD&FP register and the complement of the second source SIMD&FP register, and writes the result to the destination SIMD&FP register.
func VbicU64(v0 Uint64X1, v1 Uint64X1) Uint64X1 {
	return C.vbic_u64(v0, v1)
}

// Bitwise bit Clear (vector, register). This instruction performs a bitwise AND between the first source SIMD&FP register and the complement of the second source SIMD&FP register, and writes the result to the destination SIMD&FP register.
func VbicU16(v0 Uint16X4, v1 Uint16X4) Uint16X4 {
	return C.vbic_u16(v0, v1)
}

// Bitwise bit Clear (vector, register). This instruction performs a bitwise AND between the first source SIMD&FP register and the complement of the second source SIMD&FP register, and writes the result to the destination SIMD&FP register.
func VbicS8(v0 Int8X8, v1 Int8X8) Int8X8 {
	return C.vbic_s8(v0, v1)
}

// Bitwise bit Clear (vector, register). This instruction performs a bitwise AND between the first source SIMD&FP register and the complement of the second source SIMD&FP register, and writes the result to the destination SIMD&FP register.
func VbicS32(v0 Int32X2, v1 Int32X2) Int32X2 {
	return C.vbic_s32(v0, v1)
}

// Bitwise bit Clear (vector, register). This instruction performs a bitwise AND between the first source SIMD&FP register and the complement of the second source SIMD&FP register, and writes the result to the destination SIMD&FP register.
func VbicS64(v0 Int64X1, v1 Int64X1) Int64X1 {
	return C.vbic_s64(v0, v1)
}

// Bitwise bit Clear (vector, register). This instruction performs a bitwise AND between the first source SIMD&FP register and the complement of the second source SIMD&FP register, and writes the result to the destination SIMD&FP register.
func VbicS16(v0 Int16X4, v1 Int16X4) Int16X4 {
	return C.vbic_s16(v0, v1)
}

// Bitwise Select. This instruction sets each bit in the destination SIMD&FP register to the corresponding bit from the first source SIMD&FP register when the original destination bit was 1, otherwise from the second source SIMD&FP register.
func VbslP8(v0 Uint8X8, v1 Poly8X8, v2 Poly8X8) Poly8X8 {
	return C.vbsl_p8(v0, v1, v2)
}

// Bitwise Select. This instruction sets each bit in the destination SIMD&FP register to the corresponding bit from the first source SIMD&FP register when the original destination bit was 1, otherwise from the second source SIMD&FP register.
func VbslP16(v0 Uint16X4, v1 Poly16X4, v2 Poly16X4) Poly16X4 {
	return C.vbsl_p16(v0, v1, v2)
}

// Bitwise Select. This instruction sets each bit in the destination SIMD&FP register to the corresponding bit from the first source SIMD&FP register when the original destination bit was 1, otherwise from the second source SIMD&FP register.
func VbslqP8(v0 Uint8X16, v1 Poly8X16, v2 Poly8X16) Poly8X16 {
	return C.vbslq_p8(v0, v1, v2)
}

// Bitwise Select. This instruction sets each bit in the destination SIMD&FP register to the corresponding bit from the first source SIMD&FP register when the original destination bit was 1, otherwise from the second source SIMD&FP register.
func VbslqP16(v0 Uint16X8, v1 Poly16X8, v2 Poly16X8) Poly16X8 {
	return C.vbslq_p16(v0, v1, v2)
}

// Bitwise Select. This instruction sets each bit in the destination SIMD&FP register to the corresponding bit from the first source SIMD&FP register when the original destination bit was 1, otherwise from the second source SIMD&FP register.
func VbslqU8(v0 Uint8X16, v1 Uint8X16, v2 Uint8X16) Uint8X16 {
	return C.vbslq_u8(v0, v1, v2)
}

// Bitwise Select. This instruction sets each bit in the destination SIMD&FP register to the corresponding bit from the first source SIMD&FP register when the original destination bit was 1, otherwise from the second source SIMD&FP register.
func VbslqU32(v0 Uint32X4, v1 Uint32X4, v2 Uint32X4) Uint32X4 {
	return C.vbslq_u32(v0, v1, v2)
}

// Bitwise Select. This instruction sets each bit in the destination SIMD&FP register to the corresponding bit from the first source SIMD&FP register when the original destination bit was 1, otherwise from the second source SIMD&FP register.
func VbslqU64(v0 Uint64X2, v1 Uint64X2, v2 Uint64X2) Uint64X2 {
	return C.vbslq_u64(v0, v1, v2)
}

// Bitwise Select. This instruction sets each bit in the destination SIMD&FP register to the corresponding bit from the first source SIMD&FP register when the original destination bit was 1, otherwise from the second source SIMD&FP register.
func VbslqU16(v0 Uint16X8, v1 Uint16X8, v2 Uint16X8) Uint16X8 {
	return C.vbslq_u16(v0, v1, v2)
}

// Bitwise Select. This instruction sets each bit in the destination SIMD&FP register to the corresponding bit from the first source SIMD&FP register when the original destination bit was 1, otherwise from the second source SIMD&FP register.
func VbslqS8(v0 Uint8X16, v1 Int8X16, v2 Int8X16) Int8X16 {
	return C.vbslq_s8(v0, v1, v2)
}

// Bitwise Select. This instruction sets each bit in the destination SIMD&FP register to the corresponding bit from the first source SIMD&FP register when the original destination bit was 1, otherwise from the second source SIMD&FP register.
func VbslqF32(v0 Uint32X4, v1 Float32X4, v2 Float32X4) Float32X4 {
	return C.vbslq_f32(v0, v1, v2)
}

// Bitwise Select. This instruction sets each bit in the destination SIMD&FP register to the corresponding bit from the first source SIMD&FP register when the original destination bit was 1, otherwise from the second source SIMD&FP register.
func VbslqS32(v0 Uint32X4, v1 Int32X4, v2 Int32X4) Int32X4 {
	return C.vbslq_s32(v0, v1, v2)
}

// Bitwise Select. This instruction sets each bit in the destination SIMD&FP register to the corresponding bit from the first source SIMD&FP register when the original destination bit was 1, otherwise from the second source SIMD&FP register.
func VbslqS64(v0 Uint64X2, v1 Int64X2, v2 Int64X2) Int64X2 {
	return C.vbslq_s64(v0, v1, v2)
}

// Bitwise Select. This instruction sets each bit in the destination SIMD&FP register to the corresponding bit from the first source SIMD&FP register when the original destination bit was 1, otherwise from the second source SIMD&FP register.
func VbslqS16(v0 Uint16X8, v1 Int16X8, v2 Int16X8) Int16X8 {
	return C.vbslq_s16(v0, v1, v2)
}

// Bitwise Select. This instruction sets each bit in the destination SIMD&FP register to the corresponding bit from the first source SIMD&FP register when the original destination bit was 1, otherwise from the second source SIMD&FP register.
func VbslU8(v0 Uint8X8, v1 Uint8X8, v2 Uint8X8) Uint8X8 {
	return C.vbsl_u8(v0, v1, v2)
}

// Bitwise Select. This instruction sets each bit in the destination SIMD&FP register to the corresponding bit from the first source SIMD&FP register when the original destination bit was 1, otherwise from the second source SIMD&FP register.
func VbslU32(v0 Uint32X2, v1 Uint32X2, v2 Uint32X2) Uint32X2 {
	return C.vbsl_u32(v0, v1, v2)
}

// Bitwise Select. This instruction sets each bit in the destination SIMD&FP register to the corresponding bit from the first source SIMD&FP register when the original destination bit was 1, otherwise from the second source SIMD&FP register.
func VbslU64(v0 Uint64X1, v1 Uint64X1, v2 Uint64X1) Uint64X1 {
	return C.vbsl_u64(v0, v1, v2)
}

// Bitwise Select. This instruction sets each bit in the destination SIMD&FP register to the corresponding bit from the first source SIMD&FP register when the original destination bit was 1, otherwise from the second source SIMD&FP register.
func VbslU16(v0 Uint16X4, v1 Uint16X4, v2 Uint16X4) Uint16X4 {
	return C.vbsl_u16(v0, v1, v2)
}

// Bitwise Select. This instruction sets each bit in the destination SIMD&FP register to the corresponding bit from the first source SIMD&FP register when the original destination bit was 1, otherwise from the second source SIMD&FP register.
func VbslS8(v0 Uint8X8, v1 Int8X8, v2 Int8X8) Int8X8 {
	return C.vbsl_s8(v0, v1, v2)
}

// Bitwise Select. This instruction sets each bit in the destination SIMD&FP register to the corresponding bit from the first source SIMD&FP register when the original destination bit was 1, otherwise from the second source SIMD&FP register.
func VbslF32(v0 Uint32X2, v1 Float32X2, v2 Float32X2) Float32X2 {
	return C.vbsl_f32(v0, v1, v2)
}

// Bitwise Select. This instruction sets each bit in the destination SIMD&FP register to the corresponding bit from the first source SIMD&FP register when the original destination bit was 1, otherwise from the second source SIMD&FP register.
func VbslS32(v0 Uint32X2, v1 Int32X2, v2 Int32X2) Int32X2 {
	return C.vbsl_s32(v0, v1, v2)
}

// Bitwise Select. This instruction sets each bit in the destination SIMD&FP register to the corresponding bit from the first source SIMD&FP register when the original destination bit was 1, otherwise from the second source SIMD&FP register.
func VbslS64(v0 Uint64X1, v1 Int64X1, v2 Int64X1) Int64X1 {
	return C.vbsl_s64(v0, v1, v2)
}

// Bitwise Select. This instruction sets each bit in the destination SIMD&FP register to the corresponding bit from the first source SIMD&FP register when the original destination bit was 1, otherwise from the second source SIMD&FP register.
func VbslS16(v0 Uint16X4, v1 Int16X4, v2 Int16X4) Int16X4 {
	return C.vbsl_s16(v0, v1, v2)
}

// Floating-point Absolute Compare Greater than or Equal (vector). This instruction compares the absolute value of each floating-point value in the first source SIMD&FP register with the absolute value of the corresponding floating-point value in the second source SIMD&FP register and if the first value is greater than or equal to the second value sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
func VcageqF32(v0 Float32X4, v1 Float32X4) Uint32X4 {
	return C.vcageq_f32(v0, v1)
}

// Floating-point Absolute Compare Greater than or Equal (vector). This instruction compares the absolute value of each floating-point value in the first source SIMD&FP register with the absolute value of the corresponding floating-point value in the second source SIMD&FP register and if the first value is greater than or equal to the second value sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
func VcageF32(v0 Float32X2, v1 Float32X2) Uint32X2 {
	return C.vcage_f32(v0, v1)
}

// Floating-point Absolute Compare Greater than (vector). This instruction compares the absolute value of each vector element in the first source SIMD&FP register with the absolute value of the corresponding vector element in the second source SIMD&FP register and if the first value is greater than the second value sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
func VcagtqF32(v0 Float32X4, v1 Float32X4) Uint32X4 {
	return C.vcagtq_f32(v0, v1)
}

// Floating-point Absolute Compare Greater than (vector). This instruction compares the absolute value of each vector element in the first source SIMD&FP register with the absolute value of the corresponding vector element in the second source SIMD&FP register and if the first value is greater than the second value sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
func VcagtF32(v0 Float32X2, v1 Float32X2) Uint32X2 {
	return C.vcagt_f32(v0, v1)
}

// Floating-point absolute compare less than or equal
func VcaleqF32(v0 Float32X4, v1 Float32X4) Uint32X4 {
	return C.vcaleq_f32(v0, v1)
}

// Floating-point absolute compare less than or equal
func VcaleF32(v0 Float32X2, v1 Float32X2) Uint32X2 {
	return C.vcale_f32(v0, v1)
}

// Floating-point absolute compare less than
func VcaltqF32(v0 Float32X4, v1 Float32X4) Uint32X4 {
	return C.vcaltq_f32(v0, v1)
}

// Floating-point absolute compare less than
func VcaltF32(v0 Float32X2, v1 Float32X2) Uint32X2 {
	return C.vcalt_f32(v0, v1)
}

// Compare bitwise Equal (vector). This instruction compares each vector element from the first source SIMD&FP register with the corresponding vector element from the second source SIMD&FP register, and if the comparison is equal sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
func VceqP8(v0 Poly8X8, v1 Poly8X8) Uint8X8 {
	return C.vceq_p8(v0, v1)
}

// Compare bitwise Equal (vector). This instruction compares each vector element from the first source SIMD&FP register with the corresponding vector element from the second source SIMD&FP register, and if the comparison is equal sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
func VceqqP8(v0 Poly8X16, v1 Poly8X16) Uint8X16 {
	return C.vceqq_p8(v0, v1)
}

// Compare bitwise Equal (vector). This instruction compares each vector element from the first source SIMD&FP register with the corresponding vector element from the second source SIMD&FP register, and if the comparison is equal sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
func VceqqU8(v0 Uint8X16, v1 Uint8X16) Uint8X16 {
	return C.vceqq_u8(v0, v1)
}

// Compare bitwise Equal (vector). This instruction compares each vector element from the first source SIMD&FP register with the corresponding vector element from the second source SIMD&FP register, and if the comparison is equal sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
func VceqqU32(v0 Uint32X4, v1 Uint32X4) Uint32X4 {
	return C.vceqq_u32(v0, v1)
}

// Compare bitwise Equal (vector). This instruction compares each vector element from the first source SIMD&FP register with the corresponding vector element from the second source SIMD&FP register, and if the comparison is equal sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
func VceqqU16(v0 Uint16X8, v1 Uint16X8) Uint16X8 {
	return C.vceqq_u16(v0, v1)
}

// Compare bitwise Equal (vector). This instruction compares each vector element from the first source SIMD&FP register with the corresponding vector element from the second source SIMD&FP register, and if the comparison is equal sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
func VceqqS8(v0 Int8X16, v1 Int8X16) Uint8X16 {
	return C.vceqq_s8(v0, v1)
}

// Floating-point Compare Equal (vector). This instruction compares each floating-point value from the first source SIMD&FP register, with the corresponding floating-point value from the second source SIMD&FP register, and if the comparison is equal sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
func VceqqF32(v0 Float32X4, v1 Float32X4) Uint32X4 {
	return C.vceqq_f32(v0, v1)
}

// Compare bitwise Equal (vector). This instruction compares each vector element from the first source SIMD&FP register with the corresponding vector element from the second source SIMD&FP register, and if the comparison is equal sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
func VceqqS32(v0 Int32X4, v1 Int32X4) Uint32X4 {
	return C.vceqq_s32(v0, v1)
}

// Compare bitwise Equal (vector). This instruction compares each vector element from the first source SIMD&FP register with the corresponding vector element from the second source SIMD&FP register, and if the comparison is equal sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
func VceqqS16(v0 Int16X8, v1 Int16X8) Uint16X8 {
	return C.vceqq_s16(v0, v1)
}

// Compare bitwise Equal (vector). This instruction compares each vector element from the first source SIMD&FP register with the corresponding vector element from the second source SIMD&FP register, and if the comparison is equal sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
func VceqU8(v0 Uint8X8, v1 Uint8X8) Uint8X8 {
	return C.vceq_u8(v0, v1)
}

// Compare bitwise Equal (vector). This instruction compares each vector element from the first source SIMD&FP register with the corresponding vector element from the second source SIMD&FP register, and if the comparison is equal sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
func VceqU32(v0 Uint32X2, v1 Uint32X2) Uint32X2 {
	return C.vceq_u32(v0, v1)
}

// Compare bitwise Equal (vector). This instruction compares each vector element from the first source SIMD&FP register with the corresponding vector element from the second source SIMD&FP register, and if the comparison is equal sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
func VceqU16(v0 Uint16X4, v1 Uint16X4) Uint16X4 {
	return C.vceq_u16(v0, v1)
}

// Compare bitwise Equal (vector). This instruction compares each vector element from the first source SIMD&FP register with the corresponding vector element from the second source SIMD&FP register, and if the comparison is equal sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
func VceqS8(v0 Int8X8, v1 Int8X8) Uint8X8 {
	return C.vceq_s8(v0, v1)
}

// Floating-point Compare Equal (vector). This instruction compares each floating-point value from the first source SIMD&FP register, with the corresponding floating-point value from the second source SIMD&FP register, and if the comparison is equal sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
func VceqF32(v0 Float32X2, v1 Float32X2) Uint32X2 {
	return C.vceq_f32(v0, v1)
}

// Compare bitwise Equal (vector). This instruction compares each vector element from the first source SIMD&FP register with the corresponding vector element from the second source SIMD&FP register, and if the comparison is equal sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
func VceqS32(v0 Int32X2, v1 Int32X2) Uint32X2 {
	return C.vceq_s32(v0, v1)
}

// Compare bitwise Equal (vector). This instruction compares each vector element from the first source SIMD&FP register with the corresponding vector element from the second source SIMD&FP register, and if the comparison is equal sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
func VceqS16(v0 Int16X4, v1 Int16X4) Uint16X4 {
	return C.vceq_s16(v0, v1)
}

// Compare unsigned Higher or Same (vector). This instruction compares each vector element in the first source SIMD&FP register with the corresponding vector element in the second source SIMD&FP register and if the first unsigned integer value is greater than or equal to the second unsigned integer value sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
func VcgeqU8(v0 Uint8X16, v1 Uint8X16) Uint8X16 {
	return C.vcgeq_u8(v0, v1)
}

// Compare unsigned Higher or Same (vector). This instruction compares each vector element in the first source SIMD&FP register with the corresponding vector element in the second source SIMD&FP register and if the first unsigned integer value is greater than or equal to the second unsigned integer value sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
func VcgeqU32(v0 Uint32X4, v1 Uint32X4) Uint32X4 {
	return C.vcgeq_u32(v0, v1)
}

// Compare unsigned Higher or Same (vector). This instruction compares each vector element in the first source SIMD&FP register with the corresponding vector element in the second source SIMD&FP register and if the first unsigned integer value is greater than or equal to the second unsigned integer value sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
func VcgeqU16(v0 Uint16X8, v1 Uint16X8) Uint16X8 {
	return C.vcgeq_u16(v0, v1)
}

// Compare signed Greater than or Equal (vector). This instruction compares each vector element in the first source SIMD&FP register with the corresponding vector element in the second source SIMD&FP register and if the first signed integer value is greater than or equal to the second signed integer value sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
func VcgeqS8(v0 Int8X16, v1 Int8X16) Uint8X16 {
	return C.vcgeq_s8(v0, v1)
}

// Floating-point Compare Greater than or Equal (vector). This instruction reads each floating-point value in the first source SIMD&FP register and if the value is greater than or equal to the corresponding floating-point value in the second source SIMD&FP register sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
func VcgeqF32(v0 Float32X4, v1 Float32X4) Uint32X4 {
	return C.vcgeq_f32(v0, v1)
}

// Compare signed Greater than or Equal (vector). This instruction compares each vector element in the first source SIMD&FP register with the corresponding vector element in the second source SIMD&FP register and if the first signed integer value is greater than or equal to the second signed integer value sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
func VcgeqS32(v0 Int32X4, v1 Int32X4) Uint32X4 {
	return C.vcgeq_s32(v0, v1)
}

// Compare signed Greater than or Equal (vector). This instruction compares each vector element in the first source SIMD&FP register with the corresponding vector element in the second source SIMD&FP register and if the first signed integer value is greater than or equal to the second signed integer value sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
func VcgeqS16(v0 Int16X8, v1 Int16X8) Uint16X8 {
	return C.vcgeq_s16(v0, v1)
}

// Compare unsigned Higher or Same (vector). This instruction compares each vector element in the first source SIMD&FP register with the corresponding vector element in the second source SIMD&FP register and if the first unsigned integer value is greater than or equal to the second unsigned integer value sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
func VcgeU8(v0 Uint8X8, v1 Uint8X8) Uint8X8 {
	return C.vcge_u8(v0, v1)
}

// Compare unsigned Higher or Same (vector). This instruction compares each vector element in the first source SIMD&FP register with the corresponding vector element in the second source SIMD&FP register and if the first unsigned integer value is greater than or equal to the second unsigned integer value sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
func VcgeU32(v0 Uint32X2, v1 Uint32X2) Uint32X2 {
	return C.vcge_u32(v0, v1)
}

// Compare unsigned Higher or Same (vector). This instruction compares each vector element in the first source SIMD&FP register with the corresponding vector element in the second source SIMD&FP register and if the first unsigned integer value is greater than or equal to the second unsigned integer value sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
func VcgeU16(v0 Uint16X4, v1 Uint16X4) Uint16X4 {
	return C.vcge_u16(v0, v1)
}

// Compare signed Greater than or Equal (vector). This instruction compares each vector element in the first source SIMD&FP register with the corresponding vector element in the second source SIMD&FP register and if the first signed integer value is greater than or equal to the second signed integer value sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
func VcgeS8(v0 Int8X8, v1 Int8X8) Uint8X8 {
	return C.vcge_s8(v0, v1)
}

// Floating-point Compare Greater than or Equal (vector). This instruction reads each floating-point value in the first source SIMD&FP register and if the value is greater than or equal to the corresponding floating-point value in the second source SIMD&FP register sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
func VcgeF32(v0 Float32X2, v1 Float32X2) Uint32X2 {
	return C.vcge_f32(v0, v1)
}

// Compare signed Greater than or Equal (vector). This instruction compares each vector element in the first source SIMD&FP register with the corresponding vector element in the second source SIMD&FP register and if the first signed integer value is greater than or equal to the second signed integer value sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
func VcgeS32(v0 Int32X2, v1 Int32X2) Uint32X2 {
	return C.vcge_s32(v0, v1)
}

// Compare signed Greater than or Equal (vector). This instruction compares each vector element in the first source SIMD&FP register with the corresponding vector element in the second source SIMD&FP register and if the first signed integer value is greater than or equal to the second signed integer value sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
func VcgeS16(v0 Int16X4, v1 Int16X4) Uint16X4 {
	return C.vcge_s16(v0, v1)
}

// Compare unsigned Higher (vector). This instruction compares each vector element in the first source SIMD&FP register with the corresponding vector element in the second source SIMD&FP register and if the first unsigned integer value is greater than the second unsigned integer value sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
func VcgtqU8(v0 Uint8X16, v1 Uint8X16) Uint8X16 {
	return C.vcgtq_u8(v0, v1)
}

// Compare unsigned Higher (vector). This instruction compares each vector element in the first source SIMD&FP register with the corresponding vector element in the second source SIMD&FP register and if the first unsigned integer value is greater than the second unsigned integer value sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
func VcgtqU32(v0 Uint32X4, v1 Uint32X4) Uint32X4 {
	return C.vcgtq_u32(v0, v1)
}

// Compare unsigned Higher (vector). This instruction compares each vector element in the first source SIMD&FP register with the corresponding vector element in the second source SIMD&FP register and if the first unsigned integer value is greater than the second unsigned integer value sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
func VcgtqU16(v0 Uint16X8, v1 Uint16X8) Uint16X8 {
	return C.vcgtq_u16(v0, v1)
}

// Compare signed Greater than (vector). This instruction compares each vector element in the first source SIMD&FP register with the corresponding vector element in the second source SIMD&FP register and if the first signed integer value is greater than the second signed integer value sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
func VcgtqS8(v0 Int8X16, v1 Int8X16) Uint8X16 {
	return C.vcgtq_s8(v0, v1)
}

// Floating-point Compare Greater than (vector). This instruction reads each floating-point value in the first source SIMD&FP register and if the value is greater than the corresponding floating-point value in the second source SIMD&FP register sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
func VcgtqF32(v0 Float32X4, v1 Float32X4) Uint32X4 {
	return C.vcgtq_f32(v0, v1)
}

// Compare signed Greater than (vector). This instruction compares each vector element in the first source SIMD&FP register with the corresponding vector element in the second source SIMD&FP register and if the first signed integer value is greater than the second signed integer value sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
func VcgtqS32(v0 Int32X4, v1 Int32X4) Uint32X4 {
	return C.vcgtq_s32(v0, v1)
}

// Compare signed Greater than (vector). This instruction compares each vector element in the first source SIMD&FP register with the corresponding vector element in the second source SIMD&FP register and if the first signed integer value is greater than the second signed integer value sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
func VcgtqS16(v0 Int16X8, v1 Int16X8) Uint16X8 {
	return C.vcgtq_s16(v0, v1)
}

// Compare unsigned Higher (vector). This instruction compares each vector element in the first source SIMD&FP register with the corresponding vector element in the second source SIMD&FP register and if the first unsigned integer value is greater than the second unsigned integer value sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
func VcgtU8(v0 Uint8X8, v1 Uint8X8) Uint8X8 {
	return C.vcgt_u8(v0, v1)
}

// Compare unsigned Higher (vector). This instruction compares each vector element in the first source SIMD&FP register with the corresponding vector element in the second source SIMD&FP register and if the first unsigned integer value is greater than the second unsigned integer value sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
func VcgtU32(v0 Uint32X2, v1 Uint32X2) Uint32X2 {
	return C.vcgt_u32(v0, v1)
}

// Compare unsigned Higher (vector). This instruction compares each vector element in the first source SIMD&FP register with the corresponding vector element in the second source SIMD&FP register and if the first unsigned integer value is greater than the second unsigned integer value sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
func VcgtU16(v0 Uint16X4, v1 Uint16X4) Uint16X4 {
	return C.vcgt_u16(v0, v1)
}

// Compare signed Greater than (vector). This instruction compares each vector element in the first source SIMD&FP register with the corresponding vector element in the second source SIMD&FP register and if the first signed integer value is greater than the second signed integer value sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
func VcgtS8(v0 Int8X8, v1 Int8X8) Uint8X8 {
	return C.vcgt_s8(v0, v1)
}

// Floating-point Compare Greater than (vector). This instruction reads each floating-point value in the first source SIMD&FP register and if the value is greater than the corresponding floating-point value in the second source SIMD&FP register sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
func VcgtF32(v0 Float32X2, v1 Float32X2) Uint32X2 {
	return C.vcgt_f32(v0, v1)
}

// Compare signed Greater than (vector). This instruction compares each vector element in the first source SIMD&FP register with the corresponding vector element in the second source SIMD&FP register and if the first signed integer value is greater than the second signed integer value sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
func VcgtS32(v0 Int32X2, v1 Int32X2) Uint32X2 {
	return C.vcgt_s32(v0, v1)
}

// Compare signed Greater than (vector). This instruction compares each vector element in the first source SIMD&FP register with the corresponding vector element in the second source SIMD&FP register and if the first signed integer value is greater than the second signed integer value sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
func VcgtS16(v0 Int16X4, v1 Int16X4) Uint16X4 {
	return C.vcgt_s16(v0, v1)
}

// Compare unsigned less than or equal
func VcleqU8(v0 Uint8X16, v1 Uint8X16) Uint8X16 {
	return C.vcleq_u8(v0, v1)
}

// Compare unsigned less than or equal
func VcleqU32(v0 Uint32X4, v1 Uint32X4) Uint32X4 {
	return C.vcleq_u32(v0, v1)
}

// Compare unsigned less than or equal
func VcleqU16(v0 Uint16X8, v1 Uint16X8) Uint16X8 {
	return C.vcleq_u16(v0, v1)
}

// Compare signed less than or equal
func VcleqS8(v0 Int8X16, v1 Int8X16) Uint8X16 {
	return C.vcleq_s8(v0, v1)
}

// Floating-point compare less than or equal
func VcleqF32(v0 Float32X4, v1 Float32X4) Uint32X4 {
	return C.vcleq_f32(v0, v1)
}

// Compare signed less than or equal
func VcleqS32(v0 Int32X4, v1 Int32X4) Uint32X4 {
	return C.vcleq_s32(v0, v1)
}

// Compare signed less than or equal
func VcleqS16(v0 Int16X8, v1 Int16X8) Uint16X8 {
	return C.vcleq_s16(v0, v1)
}

// Compare unsigned less than or equal
func VcleU8(v0 Uint8X8, v1 Uint8X8) Uint8X8 {
	return C.vcle_u8(v0, v1)
}

// Compare unsigned less than or equal
func VcleU32(v0 Uint32X2, v1 Uint32X2) Uint32X2 {
	return C.vcle_u32(v0, v1)
}

// Compare unsigned less than or equal
func VcleU16(v0 Uint16X4, v1 Uint16X4) Uint16X4 {
	return C.vcle_u16(v0, v1)
}

// Compare signed less than or equal
func VcleS8(v0 Int8X8, v1 Int8X8) Uint8X8 {
	return C.vcle_s8(v0, v1)
}

// Floating-point compare less than or equal
func VcleF32(v0 Float32X2, v1 Float32X2) Uint32X2 {
	return C.vcle_f32(v0, v1)
}

// Compare signed less than or equal
func VcleS32(v0 Int32X2, v1 Int32X2) Uint32X2 {
	return C.vcle_s32(v0, v1)
}

// Compare signed less than or equal
func VcleS16(v0 Int16X4, v1 Int16X4) Uint16X4 {
	return C.vcle_s16(v0, v1)
}

// Count Leading Sign bits (vector). This instruction counts the number of consecutive bits following the most significant bit that are the same as the most significant bit in each vector element in the source SIMD&FP register, places the result into a vector, and writes the vector to the destination SIMD&FP register. The count does not include the most significant bit itself.
func VclsqU8(v0 Uint8X16) Int8X16 {
	return C.vclsq_u8(v0)
}

// Count Leading Sign bits (vector). This instruction counts the number of consecutive bits following the most significant bit that are the same as the most significant bit in each vector element in the source SIMD&FP register, places the result into a vector, and writes the vector to the destination SIMD&FP register. The count does not include the most significant bit itself.
func VclsqU32(v0 Uint32X4) Int32X4 {
	return C.vclsq_u32(v0)
}

// Count Leading Sign bits (vector). This instruction counts the number of consecutive bits following the most significant bit that are the same as the most significant bit in each vector element in the source SIMD&FP register, places the result into a vector, and writes the vector to the destination SIMD&FP register. The count does not include the most significant bit itself.
func VclsqU16(v0 Uint16X8) Int16X8 {
	return C.vclsq_u16(v0)
}

// Count Leading Sign bits (vector). This instruction counts the number of consecutive bits following the most significant bit that are the same as the most significant bit in each vector element in the source SIMD&FP register, places the result into a vector, and writes the vector to the destination SIMD&FP register. The count does not include the most significant bit itself.
func VclsqS8(v0 Int8X16) Int8X16 {
	return C.vclsq_s8(v0)
}

// Count Leading Sign bits (vector). This instruction counts the number of consecutive bits following the most significant bit that are the same as the most significant bit in each vector element in the source SIMD&FP register, places the result into a vector, and writes the vector to the destination SIMD&FP register. The count does not include the most significant bit itself.
func VclsqS32(v0 Int32X4) Int32X4 {
	return C.vclsq_s32(v0)
}

// Count Leading Sign bits (vector). This instruction counts the number of consecutive bits following the most significant bit that are the same as the most significant bit in each vector element in the source SIMD&FP register, places the result into a vector, and writes the vector to the destination SIMD&FP register. The count does not include the most significant bit itself.
func VclsqS16(v0 Int16X8) Int16X8 {
	return C.vclsq_s16(v0)
}

// Count Leading Sign bits (vector). This instruction counts the number of consecutive bits following the most significant bit that are the same as the most significant bit in each vector element in the source SIMD&FP register, places the result into a vector, and writes the vector to the destination SIMD&FP register. The count does not include the most significant bit itself.
func VclsU8(v0 Uint8X8) Int8X8 {
	return C.vcls_u8(v0)
}

// Count Leading Sign bits (vector). This instruction counts the number of consecutive bits following the most significant bit that are the same as the most significant bit in each vector element in the source SIMD&FP register, places the result into a vector, and writes the vector to the destination SIMD&FP register. The count does not include the most significant bit itself.
func VclsU32(v0 Uint32X2) Int32X2 {
	return C.vcls_u32(v0)
}

// Count Leading Sign bits (vector). This instruction counts the number of consecutive bits following the most significant bit that are the same as the most significant bit in each vector element in the source SIMD&FP register, places the result into a vector, and writes the vector to the destination SIMD&FP register. The count does not include the most significant bit itself.
func VclsU16(v0 Uint16X4) Int16X4 {
	return C.vcls_u16(v0)
}

// Count Leading Sign bits (vector). This instruction counts the number of consecutive bits following the most significant bit that are the same as the most significant bit in each vector element in the source SIMD&FP register, places the result into a vector, and writes the vector to the destination SIMD&FP register. The count does not include the most significant bit itself.
func VclsS8(v0 Int8X8) Int8X8 {
	return C.vcls_s8(v0)
}

// Count Leading Sign bits (vector). This instruction counts the number of consecutive bits following the most significant bit that are the same as the most significant bit in each vector element in the source SIMD&FP register, places the result into a vector, and writes the vector to the destination SIMD&FP register. The count does not include the most significant bit itself.
func VclsS32(v0 Int32X2) Int32X2 {
	return C.vcls_s32(v0)
}

// Count Leading Sign bits (vector). This instruction counts the number of consecutive bits following the most significant bit that are the same as the most significant bit in each vector element in the source SIMD&FP register, places the result into a vector, and writes the vector to the destination SIMD&FP register. The count does not include the most significant bit itself.
func VclsS16(v0 Int16X4) Int16X4 {
	return C.vcls_s16(v0)
}

// Compare unsigned less than
func VcltqU8(v0 Uint8X16, v1 Uint8X16) Uint8X16 {
	return C.vcltq_u8(v0, v1)
}

// Compare unsigned less than
func VcltqU32(v0 Uint32X4, v1 Uint32X4) Uint32X4 {
	return C.vcltq_u32(v0, v1)
}

// Compare unsigned less than
func VcltqU16(v0 Uint16X8, v1 Uint16X8) Uint16X8 {
	return C.vcltq_u16(v0, v1)
}

// Compare signed less than
func VcltqS8(v0 Int8X16, v1 Int8X16) Uint8X16 {
	return C.vcltq_s8(v0, v1)
}

// Floating-point compare less than
func VcltqF32(v0 Float32X4, v1 Float32X4) Uint32X4 {
	return C.vcltq_f32(v0, v1)
}

// Compare signed less than
func VcltqS32(v0 Int32X4, v1 Int32X4) Uint32X4 {
	return C.vcltq_s32(v0, v1)
}

// Compare signed less than
func VcltqS16(v0 Int16X8, v1 Int16X8) Uint16X8 {
	return C.vcltq_s16(v0, v1)
}

// Compare unsigned less than
func VcltU8(v0 Uint8X8, v1 Uint8X8) Uint8X8 {
	return C.vclt_u8(v0, v1)
}

// Compare unsigned less than
func VcltU32(v0 Uint32X2, v1 Uint32X2) Uint32X2 {
	return C.vclt_u32(v0, v1)
}

// Compare unsigned less than
func VcltU16(v0 Uint16X4, v1 Uint16X4) Uint16X4 {
	return C.vclt_u16(v0, v1)
}

// Compare signed less than
func VcltS8(v0 Int8X8, v1 Int8X8) Uint8X8 {
	return C.vclt_s8(v0, v1)
}

// Floating-point compare less than
func VcltF32(v0 Float32X2, v1 Float32X2) Uint32X2 {
	return C.vclt_f32(v0, v1)
}

// Compare signed less than
func VcltS32(v0 Int32X2, v1 Int32X2) Uint32X2 {
	return C.vclt_s32(v0, v1)
}

// Compare signed less than
func VcltS16(v0 Int16X4, v1 Int16X4) Uint16X4 {
	return C.vclt_s16(v0, v1)
}

// Count Leading Zero bits (vector). This instruction counts the number of consecutive zeros, starting from the most significant bit, in each vector element in the source SIMD&FP register, places the result into a vector, and writes the vector to the destination SIMD&FP register.
func VclzqU8(v0 Uint8X16) Uint8X16 {
	return C.vclzq_u8(v0)
}

// Count Leading Zero bits (vector). This instruction counts the number of consecutive zeros, starting from the most significant bit, in each vector element in the source SIMD&FP register, places the result into a vector, and writes the vector to the destination SIMD&FP register.
func VclzqU32(v0 Uint32X4) Uint32X4 {
	return C.vclzq_u32(v0)
}

// Count Leading Zero bits (vector). This instruction counts the number of consecutive zeros, starting from the most significant bit, in each vector element in the source SIMD&FP register, places the result into a vector, and writes the vector to the destination SIMD&FP register.
func VclzqU16(v0 Uint16X8) Uint16X8 {
	return C.vclzq_u16(v0)
}

// Count Leading Zero bits (vector). This instruction counts the number of consecutive zeros, starting from the most significant bit, in each vector element in the source SIMD&FP register, places the result into a vector, and writes the vector to the destination SIMD&FP register.
func VclzqS8(v0 Int8X16) Int8X16 {
	return C.vclzq_s8(v0)
}

// Count Leading Zero bits (vector). This instruction counts the number of consecutive zeros, starting from the most significant bit, in each vector element in the source SIMD&FP register, places the result into a vector, and writes the vector to the destination SIMD&FP register.
func VclzqS32(v0 Int32X4) Int32X4 {
	return C.vclzq_s32(v0)
}

// Count Leading Zero bits (vector). This instruction counts the number of consecutive zeros, starting from the most significant bit, in each vector element in the source SIMD&FP register, places the result into a vector, and writes the vector to the destination SIMD&FP register.
func VclzqS16(v0 Int16X8) Int16X8 {
	return C.vclzq_s16(v0)
}

// Count Leading Zero bits (vector). This instruction counts the number of consecutive zeros, starting from the most significant bit, in each vector element in the source SIMD&FP register, places the result into a vector, and writes the vector to the destination SIMD&FP register.
func VclzU8(v0 Uint8X8) Uint8X8 {
	return C.vclz_u8(v0)
}

// Count Leading Zero bits (vector). This instruction counts the number of consecutive zeros, starting from the most significant bit, in each vector element in the source SIMD&FP register, places the result into a vector, and writes the vector to the destination SIMD&FP register.
func VclzU32(v0 Uint32X2) Uint32X2 {
	return C.vclz_u32(v0)
}

// Count Leading Zero bits (vector). This instruction counts the number of consecutive zeros, starting from the most significant bit, in each vector element in the source SIMD&FP register, places the result into a vector, and writes the vector to the destination SIMD&FP register.
func VclzU16(v0 Uint16X4) Uint16X4 {
	return C.vclz_u16(v0)
}

// Count Leading Zero bits (vector). This instruction counts the number of consecutive zeros, starting from the most significant bit, in each vector element in the source SIMD&FP register, places the result into a vector, and writes the vector to the destination SIMD&FP register.
func VclzS8(v0 Int8X8) Int8X8 {
	return C.vclz_s8(v0)
}

// Count Leading Zero bits (vector). This instruction counts the number of consecutive zeros, starting from the most significant bit, in each vector element in the source SIMD&FP register, places the result into a vector, and writes the vector to the destination SIMD&FP register.
func VclzS32(v0 Int32X2) Int32X2 {
	return C.vclz_s32(v0)
}

// Count Leading Zero bits (vector). This instruction counts the number of consecutive zeros, starting from the most significant bit, in each vector element in the source SIMD&FP register, places the result into a vector, and writes the vector to the destination SIMD&FP register.
func VclzS16(v0 Int16X4) Int16X4 {
	return C.vclz_s16(v0)
}

// Population Count per byte. This instruction counts the number of bits that have a value of one in each vector element in the source SIMD&FP register, places the result into a vector, and writes the vector to the destination SIMD&FP register.
func VcntP8(v0 Poly8X8) Poly8X8 {
	return C.vcnt_p8(v0)
}

// Population Count per byte. This instruction counts the number of bits that have a value of one in each vector element in the source SIMD&FP register, places the result into a vector, and writes the vector to the destination SIMD&FP register.
func VcntqP8(v0 Poly8X16) Poly8X16 {
	return C.vcntq_p8(v0)
}

// Population Count per byte. This instruction counts the number of bits that have a value of one in each vector element in the source SIMD&FP register, places the result into a vector, and writes the vector to the destination SIMD&FP register.
func VcntqU8(v0 Uint8X16) Uint8X16 {
	return C.vcntq_u8(v0)
}

// Population Count per byte. This instruction counts the number of bits that have a value of one in each vector element in the source SIMD&FP register, places the result into a vector, and writes the vector to the destination SIMD&FP register.
func VcntqS8(v0 Int8X16) Int8X16 {
	return C.vcntq_s8(v0)
}

// Population Count per byte. This instruction counts the number of bits that have a value of one in each vector element in the source SIMD&FP register, places the result into a vector, and writes the vector to the destination SIMD&FP register.
func VcntU8(v0 Uint8X8) Uint8X8 {
	return C.vcnt_u8(v0)
}

// Population Count per byte. This instruction counts the number of bits that have a value of one in each vector element in the source SIMD&FP register, places the result into a vector, and writes the vector to the destination SIMD&FP register.
func VcntS8(v0 Int8X8) Int8X8 {
	return C.vcnt_s8(v0)
}

// Join two smaller vectors into a single larger vector
func VcombineP8(v0 Poly8X8, v1 Poly8X8) Poly8X16 {
	return C.vcombine_p8(v0, v1)
}

// Join two smaller vectors into a single larger vector
func VcombineP16(v0 Poly16X4, v1 Poly16X4) Poly16X8 {
	return C.vcombine_p16(v0, v1)
}

// Join two smaller vectors into a single larger vector
func VcombineU8(v0 Uint8X8, v1 Uint8X8) Uint8X16 {
	return C.vcombine_u8(v0, v1)
}

// Join two smaller vectors into a single larger vector
func VcombineU32(v0 Uint32X2, v1 Uint32X2) Uint32X4 {
	return C.vcombine_u32(v0, v1)
}

// Join two smaller vectors into a single larger vector
func VcombineU64(v0 Uint64X1, v1 Uint64X1) Uint64X2 {
	return C.vcombine_u64(v0, v1)
}

// Join two smaller vectors into a single larger vector
func VcombineU16(v0 Uint16X4, v1 Uint16X4) Uint16X8 {
	return C.vcombine_u16(v0, v1)
}

// Join two smaller vectors into a single larger vector
func VcombineS8(v0 Int8X8, v1 Int8X8) Int8X16 {
	return C.vcombine_s8(v0, v1)
}

// Join two smaller vectors into a single larger vector
func VcombineF32(v0 Float32X2, v1 Float32X2) Float32X4 {
	return C.vcombine_f32(v0, v1)
}

// Join two smaller vectors into a single larger vector
func VcombineS32(v0 Int32X2, v1 Int32X2) Int32X4 {
	return C.vcombine_s32(v0, v1)
}

// Join two smaller vectors into a single larger vector
func VcombineS64(v0 Int64X1, v1 Int64X1) Int64X2 {
	return C.vcombine_s64(v0, v1)
}

// Join two smaller vectors into a single larger vector
func VcombineS16(v0 Int16X4, v1 Int16X4) Int16X8 {
	return C.vcombine_s16(v0, v1)
}

// Unsigned fixed-point Convert to Floating-point (vector). This instruction converts each element in a vector from fixed-point to floating-point using the rounding mode that is specified by the FPCR, and writes the result to the SIMD&FP destination register.
func VcvtqF32U32(v0 Uint32X4) Float32X4 {
	return C.vcvtq_f32_u32(v0)
}

// Signed fixed-point Convert to Floating-point (vector). This instruction converts each element in a vector from fixed-point to floating-point using the rounding mode that is specified by the FPCR, and writes the result to the SIMD&FP destination register.
func VcvtqF32S32(v0 Int32X4) Float32X4 {
	return C.vcvtq_f32_s32(v0)
}

// Unsigned fixed-point Convert to Floating-point (vector). This instruction converts each element in a vector from fixed-point to floating-point using the rounding mode that is specified by the FPCR, and writes the result to the SIMD&FP destination register.
func VcvtF32U32(v0 Uint32X2) Float32X2 {
	return C.vcvt_f32_u32(v0)
}

// Signed fixed-point Convert to Floating-point (vector). This instruction converts each element in a vector from fixed-point to floating-point using the rounding mode that is specified by the FPCR, and writes the result to the SIMD&FP destination register.
func VcvtF32S32(v0 Int32X2) Float32X2 {
	return C.vcvt_f32_s32(v0)
}

// Floating-point Convert to Signed fixed-point, rounding toward Zero (vector). This instruction converts a scalar or each element in a vector from floating-point to fixed-point signed integer using the Round towards Zero rounding mode, and writes the result to the SIMD&FP destination register.
func VcvtqS32F32(v0 Float32X4) Int32X4 {
	return C.vcvtq_s32_f32(v0)
}

// Floating-point Convert to Signed fixed-point, rounding toward Zero (vector). This instruction converts a scalar or each element in a vector from floating-point to fixed-point signed integer using the Round towards Zero rounding mode, and writes the result to the SIMD&FP destination register.
func VcvtS32F32(v0 Float32X2) Int32X2 {
	return C.vcvt_s32_f32(v0)
}

// Floating-point Convert to Unsigned fixed-point, rounding toward Zero (vector). This instruction converts a scalar or each element in a vector from floating-point to fixed-point unsigned integer using the Round towards Zero rounding mode, and writes the result to the general-purpose destination register.
func VcvtqU32F32(v0 Float32X4) Uint32X4 {
	return C.vcvtq_u32_f32(v0)
}

// Floating-point Convert to Unsigned fixed-point, rounding toward Zero (vector). This instruction converts a scalar or each element in a vector from floating-point to fixed-point unsigned integer using the Round towards Zero rounding mode, and writes the result to the general-purpose destination register.
func VcvtU32F32(v0 Float32X2) Uint32X2 {
	return C.vcvt_u32_f32(v0)
}

// Duplicate vector element to vector or scalar. This instruction duplicates the vector element at the specified element index in the source SIMD&FP register into a scalar or each element in a vector, and writes the result to the destination SIMD&FP register.
func VdupNP8(v0 Poly8) Poly8X8 {
	return C.vdup_n_p8(v0)
}

// Duplicate vector element to vector or scalar. This instruction duplicates the vector element at the specified element index in the source SIMD&FP register into a scalar or each element in a vector, and writes the result to the destination SIMD&FP register.
func VdupNP16(v0 Poly16) Poly16X4 {
	return C.vdup_n_p16(v0)
}

// Duplicate vector element to vector or scalar. This instruction duplicates the vector element at the specified element index in the source SIMD&FP register into a scalar or each element in a vector, and writes the result to the destination SIMD&FP register.
func VdupqNP8(v0 Poly8) Poly8X16 {
	return C.vdupq_n_p8(v0)
}

// Duplicate vector element to vector or scalar. This instruction duplicates the vector element at the specified element index in the source SIMD&FP register into a scalar or each element in a vector, and writes the result to the destination SIMD&FP register.
func VdupqNP16(v0 Poly16) Poly16X8 {
	return C.vdupq_n_p16(v0)
}

// Duplicate vector element to vector or scalar. This instruction duplicates the vector element at the specified element index in the source SIMD&FP register into a scalar or each element in a vector, and writes the result to the destination SIMD&FP register.
func VdupqNU8(v0 Uint8) Uint8X16 {
	return C.vdupq_n_u8(v0)
}

// Duplicate vector element to vector or scalar. This instruction duplicates the vector element at the specified element index in the source SIMD&FP register into a scalar or each element in a vector, and writes the result to the destination SIMD&FP register.
func VdupqNU32(v0 Uint32) Uint32X4 {
	return C.vdupq_n_u32(v0)
}

// Duplicate vector element to vector or scalar. This instruction duplicates the vector element at the specified element index in the source SIMD&FP register into a scalar or each element in a vector, and writes the result to the destination SIMD&FP register.
func VdupqNU64(v0 Uint64) Uint64X2 {
	return C.vdupq_n_u64(v0)
}

// Duplicate vector element to vector or scalar. This instruction duplicates the vector element at the specified element index in the source SIMD&FP register into a scalar or each element in a vector, and writes the result to the destination SIMD&FP register.
func VdupqNU16(v0 Uint16) Uint16X8 {
	return C.vdupq_n_u16(v0)
}

// Duplicate vector element to vector or scalar. This instruction duplicates the vector element at the specified element index in the source SIMD&FP register into a scalar or each element in a vector, and writes the result to the destination SIMD&FP register.
func VdupqNS8(v0 Int8) Int8X16 {
	return C.vdupq_n_s8(v0)
}

// Duplicate vector element to vector or scalar. This instruction duplicates the vector element at the specified element index in the source SIMD&FP register into a scalar or each element in a vector, and writes the result to the destination SIMD&FP register.
func VdupqNF32(v0 Float32) Float32X4 {
	return C.vdupq_n_f32(v0)
}

// Duplicate vector element to vector or scalar. This instruction duplicates the vector element at the specified element index in the source SIMD&FP register into a scalar or each element in a vector, and writes the result to the destination SIMD&FP register.
func VdupqNS32(v0 Int32) Int32X4 {
	return C.vdupq_n_s32(v0)
}

// Duplicate vector element to vector or scalar. This instruction duplicates the vector element at the specified element index in the source SIMD&FP register into a scalar or each element in a vector, and writes the result to the destination SIMD&FP register.
func VdupqNS64(v0 Int64) Int64X2 {
	return C.vdupq_n_s64(v0)
}

// Duplicate vector element to vector or scalar. This instruction duplicates the vector element at the specified element index in the source SIMD&FP register into a scalar or each element in a vector, and writes the result to the destination SIMD&FP register.
func VdupqNS16(v0 Int16) Int16X8 {
	return C.vdupq_n_s16(v0)
}

// Duplicate vector element to vector or scalar. This instruction duplicates the vector element at the specified element index in the source SIMD&FP register into a scalar or each element in a vector, and writes the result to the destination SIMD&FP register.
func VdupNU8(v0 Uint8) Uint8X8 {
	return C.vdup_n_u8(v0)
}

// Duplicate vector element to vector or scalar. This instruction duplicates the vector element at the specified element index in the source SIMD&FP register into a scalar or each element in a vector, and writes the result to the destination SIMD&FP register.
func VdupNU32(v0 Uint32) Uint32X2 {
	return C.vdup_n_u32(v0)
}

// Insert vector element from another vector element. This instruction copies the vector element of the source SIMD&FP register to the specified vector element of the destination SIMD&FP register.
func VdupNU64(v0 Uint64) Uint64X1 {
	return C.vdup_n_u64(v0)
}

// Duplicate vector element to vector or scalar. This instruction duplicates the vector element at the specified element index in the source SIMD&FP register into a scalar or each element in a vector, and writes the result to the destination SIMD&FP register.
func VdupNU16(v0 Uint16) Uint16X4 {
	return C.vdup_n_u16(v0)
}

// Duplicate vector element to vector or scalar. This instruction duplicates the vector element at the specified element index in the source SIMD&FP register into a scalar or each element in a vector, and writes the result to the destination SIMD&FP register.
func VdupNS8(v0 Int8) Int8X8 {
	return C.vdup_n_s8(v0)
}

// Duplicate vector element to vector or scalar. This instruction duplicates the vector element at the specified element index in the source SIMD&FP register into a scalar or each element in a vector, and writes the result to the destination SIMD&FP register.
func VdupNF32(v0 Float32) Float32X2 {
	return C.vdup_n_f32(v0)
}

// Duplicate vector element to vector or scalar. This instruction duplicates the vector element at the specified element index in the source SIMD&FP register into a scalar or each element in a vector, and writes the result to the destination SIMD&FP register.
func VdupNS32(v0 Int32) Int32X2 {
	return C.vdup_n_s32(v0)
}

// Insert vector element from another vector element. This instruction copies the vector element of the source SIMD&FP register to the specified vector element of the destination SIMD&FP register.
func VdupNS64(v0 Int64) Int64X1 {
	return C.vdup_n_s64(v0)
}

// Duplicate vector element to vector or scalar. This instruction duplicates the vector element at the specified element index in the source SIMD&FP register into a scalar or each element in a vector, and writes the result to the destination SIMD&FP register.
func VdupNS16(v0 Int16) Int16X4 {
	return C.vdup_n_s16(v0)
}

// Bitwise Exclusive OR (vector). This instruction performs a bitwise Exclusive OR operation between the two source SIMD&FP registers, and places the result in the destination SIMD&FP register.
func VeorqU8(v0 Uint8X16, v1 Uint8X16) Uint8X16 {
	return C.veorq_u8(v0, v1)
}

// Bitwise Exclusive OR (vector). This instruction performs a bitwise Exclusive OR operation between the two source SIMD&FP registers, and places the result in the destination SIMD&FP register.
func VeorqU32(v0 Uint32X4, v1 Uint32X4) Uint32X4 {
	return C.veorq_u32(v0, v1)
}

// Bitwise Exclusive OR (vector). This instruction performs a bitwise Exclusive OR operation between the two source SIMD&FP registers, and places the result in the destination SIMD&FP register.
func VeorqU64(v0 Uint64X2, v1 Uint64X2) Uint64X2 {
	return C.veorq_u64(v0, v1)
}

// Bitwise Exclusive OR (vector). This instruction performs a bitwise Exclusive OR operation between the two source SIMD&FP registers, and places the result in the destination SIMD&FP register.
func VeorqU16(v0 Uint16X8, v1 Uint16X8) Uint16X8 {
	return C.veorq_u16(v0, v1)
}

// Bitwise Exclusive OR (vector). This instruction performs a bitwise Exclusive OR operation between the two source SIMD&FP registers, and places the result in the destination SIMD&FP register.
func VeorqS8(v0 Int8X16, v1 Int8X16) Int8X16 {
	return C.veorq_s8(v0, v1)
}

// Bitwise Exclusive OR (vector). This instruction performs a bitwise Exclusive OR operation between the two source SIMD&FP registers, and places the result in the destination SIMD&FP register.
func VeorqS32(v0 Int32X4, v1 Int32X4) Int32X4 {
	return C.veorq_s32(v0, v1)
}

// Bitwise Exclusive OR (vector). This instruction performs a bitwise Exclusive OR operation between the two source SIMD&FP registers, and places the result in the destination SIMD&FP register.
func VeorqS64(v0 Int64X2, v1 Int64X2) Int64X2 {
	return C.veorq_s64(v0, v1)
}

// Bitwise Exclusive OR (vector). This instruction performs a bitwise Exclusive OR operation between the two source SIMD&FP registers, and places the result in the destination SIMD&FP register.
func VeorqS16(v0 Int16X8, v1 Int16X8) Int16X8 {
	return C.veorq_s16(v0, v1)
}

// Bitwise Exclusive OR (vector). This instruction performs a bitwise Exclusive OR operation between the two source SIMD&FP registers, and places the result in the destination SIMD&FP register.
func VeorU8(v0 Uint8X8, v1 Uint8X8) Uint8X8 {
	return C.veor_u8(v0, v1)
}

// Bitwise Exclusive OR (vector). This instruction performs a bitwise Exclusive OR operation between the two source SIMD&FP registers, and places the result in the destination SIMD&FP register.
func VeorU32(v0 Uint32X2, v1 Uint32X2) Uint32X2 {
	return C.veor_u32(v0, v1)
}

// Bitwise Exclusive OR (vector). This instruction performs a bitwise Exclusive OR operation between the two source SIMD&FP registers, and places the result in the destination SIMD&FP register.
func VeorU64(v0 Uint64X1, v1 Uint64X1) Uint64X1 {
	return C.veor_u64(v0, v1)
}

// Bitwise Exclusive OR (vector). This instruction performs a bitwise Exclusive OR operation between the two source SIMD&FP registers, and places the result in the destination SIMD&FP register.
func VeorU16(v0 Uint16X4, v1 Uint16X4) Uint16X4 {
	return C.veor_u16(v0, v1)
}

// Bitwise Exclusive OR (vector). This instruction performs a bitwise Exclusive OR operation between the two source SIMD&FP registers, and places the result in the destination SIMD&FP register.
func VeorS8(v0 Int8X8, v1 Int8X8) Int8X8 {
	return C.veor_s8(v0, v1)
}

// Bitwise Exclusive OR (vector). This instruction performs a bitwise Exclusive OR operation between the two source SIMD&FP registers, and places the result in the destination SIMD&FP register.
func VeorS32(v0 Int32X2, v1 Int32X2) Int32X2 {
	return C.veor_s32(v0, v1)
}

// Bitwise Exclusive OR (vector). This instruction performs a bitwise Exclusive OR operation between the two source SIMD&FP registers, and places the result in the destination SIMD&FP register.
func VeorS64(v0 Int64X1, v1 Int64X1) Int64X1 {
	return C.veor_s64(v0, v1)
}

// Bitwise Exclusive OR (vector). This instruction performs a bitwise Exclusive OR operation between the two source SIMD&FP registers, and places the result in the destination SIMD&FP register.
func VeorS16(v0 Int16X4, v1 Int16X4) Int16X4 {
	return C.veor_s16(v0, v1)
}

// Duplicate vector element to vector or scalar. This instruction duplicates the vector element at the specified element index in the source SIMD&FP register into a scalar or each element in a vector, and writes the result to the destination SIMD&FP register.
func VgetHighP8(v0 Poly8X16) Poly8X8 {
	return C.vget_high_p8(v0)
}

// Duplicate vector element to vector or scalar. This instruction duplicates the vector element at the specified element index in the source SIMD&FP register into a scalar or each element in a vector, and writes the result to the destination SIMD&FP register.
func VgetHighP16(v0 Poly16X8) Poly16X4 {
	return C.vget_high_p16(v0)
}

// Duplicate vector element to vector or scalar. This instruction duplicates the vector element at the specified element index in the source SIMD&FP register into a scalar or each element in a vector, and writes the result to the destination SIMD&FP register.
func VgetHighU8(v0 Uint8X16) Uint8X8 {
	return C.vget_high_u8(v0)
}

// Duplicate vector element to vector or scalar. This instruction duplicates the vector element at the specified element index in the source SIMD&FP register into a scalar or each element in a vector, and writes the result to the destination SIMD&FP register.
func VgetHighU32(v0 Uint32X4) Uint32X2 {
	return C.vget_high_u32(v0)
}

// Duplicate vector element to vector or scalar. This instruction duplicates the vector element at the specified element index in the source SIMD&FP register into a scalar or each element in a vector, and writes the result to the destination SIMD&FP register.
func VgetHighU64(v0 Uint64X2) Uint64X1 {
	return C.vget_high_u64(v0)
}

// Duplicate vector element to vector or scalar. This instruction duplicates the vector element at the specified element index in the source SIMD&FP register into a scalar or each element in a vector, and writes the result to the destination SIMD&FP register.
func VgetHighU16(v0 Uint16X8) Uint16X4 {
	return C.vget_high_u16(v0)
}

// Duplicate vector element to vector or scalar. This instruction duplicates the vector element at the specified element index in the source SIMD&FP register into a scalar or each element in a vector, and writes the result to the destination SIMD&FP register.
func VgetHighS8(v0 Int8X16) Int8X8 {
	return C.vget_high_s8(v0)
}

// Duplicate vector element to vector or scalar. This instruction duplicates the vector element at the specified element index in the source SIMD&FP register into a scalar or each element in a vector, and writes the result to the destination SIMD&FP register.
func VgetHighF32(v0 Float32X4) Float32X2 {
	return C.vget_high_f32(v0)
}

// Duplicate vector element to vector or scalar. This instruction duplicates the vector element at the specified element index in the source SIMD&FP register into a scalar or each element in a vector, and writes the result to the destination SIMD&FP register.
func VgetHighS32(v0 Int32X4) Int32X2 {
	return C.vget_high_s32(v0)
}

// Duplicate vector element to vector or scalar. This instruction duplicates the vector element at the specified element index in the source SIMD&FP register into a scalar or each element in a vector, and writes the result to the destination SIMD&FP register.
func VgetHighS64(v0 Int64X2) Int64X1 {
	return C.vget_high_s64(v0)
}

// Duplicate vector element to vector or scalar. This instruction duplicates the vector element at the specified element index in the source SIMD&FP register into a scalar or each element in a vector, and writes the result to the destination SIMD&FP register.
func VgetHighS16(v0 Int16X8) Int16X4 {
	return C.vget_high_s16(v0)
}

// Duplicate vector element to vector or scalar. This instruction duplicates the vector element at the specified element index in the source SIMD&FP register into a scalar or each element in a vector, and writes the result to the destination SIMD&FP register.
func VgetLowP8(v0 Poly8X16) Poly8X8 {
	return C.vget_low_p8(v0)
}

// Duplicate vector element to vector or scalar. This instruction duplicates the vector element at the specified element index in the source SIMD&FP register into a scalar or each element in a vector, and writes the result to the destination SIMD&FP register.
func VgetLowP16(v0 Poly16X8) Poly16X4 {
	return C.vget_low_p16(v0)
}

// Duplicate vector element to vector or scalar. This instruction duplicates the vector element at the specified element index in the source SIMD&FP register into a scalar or each element in a vector, and writes the result to the destination SIMD&FP register.
func VgetLowU8(v0 Uint8X16) Uint8X8 {
	return C.vget_low_u8(v0)
}

// Duplicate vector element to vector or scalar. This instruction duplicates the vector element at the specified element index in the source SIMD&FP register into a scalar or each element in a vector, and writes the result to the destination SIMD&FP register.
func VgetLowU32(v0 Uint32X4) Uint32X2 {
	return C.vget_low_u32(v0)
}

// Duplicate vector element to vector or scalar. This instruction duplicates the vector element at the specified element index in the source SIMD&FP register into a scalar or each element in a vector, and writes the result to the destination SIMD&FP register.
func VgetLowU64(v0 Uint64X2) Uint64X1 {
	return C.vget_low_u64(v0)
}

// Duplicate vector element to vector or scalar. This instruction duplicates the vector element at the specified element index in the source SIMD&FP register into a scalar or each element in a vector, and writes the result to the destination SIMD&FP register.
func VgetLowU16(v0 Uint16X8) Uint16X4 {
	return C.vget_low_u16(v0)
}

// Duplicate vector element to vector or scalar. This instruction duplicates the vector element at the specified element index in the source SIMD&FP register into a scalar or each element in a vector, and writes the result to the destination SIMD&FP register.
func VgetLowS8(v0 Int8X16) Int8X8 {
	return C.vget_low_s8(v0)
}

// Duplicate vector element to vector or scalar. This instruction duplicates the vector element at the specified element index in the source SIMD&FP register into a scalar or each element in a vector, and writes the result to the destination SIMD&FP register.
func VgetLowF32(v0 Float32X4) Float32X2 {
	return C.vget_low_f32(v0)
}

// Duplicate vector element to vector or scalar. This instruction duplicates the vector element at the specified element index in the source SIMD&FP register into a scalar or each element in a vector, and writes the result to the destination SIMD&FP register.
func VgetLowS32(v0 Int32X4) Int32X2 {
	return C.vget_low_s32(v0)
}

// Duplicate vector element to vector or scalar. This instruction duplicates the vector element at the specified element index in the source SIMD&FP register into a scalar or each element in a vector, and writes the result to the destination SIMD&FP register.
func VgetLowS64(v0 Int64X2) Int64X1 {
	return C.vget_low_s64(v0)
}

// Duplicate vector element to vector or scalar. This instruction duplicates the vector element at the specified element index in the source SIMD&FP register into a scalar or each element in a vector, and writes the result to the destination SIMD&FP register.
func VgetLowS16(v0 Int16X8) Int16X4 {
	return C.vget_low_s16(v0)
}

// Unsigned Halving Add. This instruction adds corresponding unsigned integer values from the two source SIMD&FP registers, shifts each result right one bit, places the results into a vector, and writes the vector to the destination SIMD&FP register.
func VhaddqU8(v0 Uint8X16, v1 Uint8X16) Uint8X16 {
	return C.vhaddq_u8(v0, v1)
}

// Unsigned Halving Add. This instruction adds corresponding unsigned integer values from the two source SIMD&FP registers, shifts each result right one bit, places the results into a vector, and writes the vector to the destination SIMD&FP register.
func VhaddqU32(v0 Uint32X4, v1 Uint32X4) Uint32X4 {
	return C.vhaddq_u32(v0, v1)
}

// Unsigned Halving Add. This instruction adds corresponding unsigned integer values from the two source SIMD&FP registers, shifts each result right one bit, places the results into a vector, and writes the vector to the destination SIMD&FP register.
func VhaddqU16(v0 Uint16X8, v1 Uint16X8) Uint16X8 {
	return C.vhaddq_u16(v0, v1)
}

// Signed Halving Add. This instruction adds corresponding signed integer values from the two source SIMD&FP registers, shifts each result right one bit, places the results into a vector, and writes the vector to the destination SIMD&FP register.
func VhaddqS8(v0 Int8X16, v1 Int8X16) Int8X16 {
	return C.vhaddq_s8(v0, v1)
}

// Signed Halving Add. This instruction adds corresponding signed integer values from the two source SIMD&FP registers, shifts each result right one bit, places the results into a vector, and writes the vector to the destination SIMD&FP register.
func VhaddqS32(v0 Int32X4, v1 Int32X4) Int32X4 {
	return C.vhaddq_s32(v0, v1)
}

// Signed Halving Add. This instruction adds corresponding signed integer values from the two source SIMD&FP registers, shifts each result right one bit, places the results into a vector, and writes the vector to the destination SIMD&FP register.
func VhaddqS16(v0 Int16X8, v1 Int16X8) Int16X8 {
	return C.vhaddq_s16(v0, v1)
}

// Unsigned Halving Add. This instruction adds corresponding unsigned integer values from the two source SIMD&FP registers, shifts each result right one bit, places the results into a vector, and writes the vector to the destination SIMD&FP register.
func VhaddU8(v0 Uint8X8, v1 Uint8X8) Uint8X8 {
	return C.vhadd_u8(v0, v1)
}

// Unsigned Halving Add. This instruction adds corresponding unsigned integer values from the two source SIMD&FP registers, shifts each result right one bit, places the results into a vector, and writes the vector to the destination SIMD&FP register.
func VhaddU32(v0 Uint32X2, v1 Uint32X2) Uint32X2 {
	return C.vhadd_u32(v0, v1)
}

// Unsigned Halving Add. This instruction adds corresponding unsigned integer values from the two source SIMD&FP registers, shifts each result right one bit, places the results into a vector, and writes the vector to the destination SIMD&FP register.
func VhaddU16(v0 Uint16X4, v1 Uint16X4) Uint16X4 {
	return C.vhadd_u16(v0, v1)
}

// Signed Halving Add. This instruction adds corresponding signed integer values from the two source SIMD&FP registers, shifts each result right one bit, places the results into a vector, and writes the vector to the destination SIMD&FP register.
func VhaddS8(v0 Int8X8, v1 Int8X8) Int8X8 {
	return C.vhadd_s8(v0, v1)
}

// Signed Halving Add. This instruction adds corresponding signed integer values from the two source SIMD&FP registers, shifts each result right one bit, places the results into a vector, and writes the vector to the destination SIMD&FP register.
func VhaddS32(v0 Int32X2, v1 Int32X2) Int32X2 {
	return C.vhadd_s32(v0, v1)
}

// Signed Halving Add. This instruction adds corresponding signed integer values from the two source SIMD&FP registers, shifts each result right one bit, places the results into a vector, and writes the vector to the destination SIMD&FP register.
func VhaddS16(v0 Int16X4, v1 Int16X4) Int16X4 {
	return C.vhadd_s16(v0, v1)
}

// Unsigned Halving Subtract. This instruction subtracts the vector elements in the second source SIMD&FP register from the corresponding vector elements in the first source SIMD&FP register, shifts each result right one bit, places each result into a vector, and writes the vector to the destination SIMD&FP register.
func VhsubqU8(v0 Uint8X16, v1 Uint8X16) Uint8X16 {
	return C.vhsubq_u8(v0, v1)
}

// Unsigned Halving Subtract. This instruction subtracts the vector elements in the second source SIMD&FP register from the corresponding vector elements in the first source SIMD&FP register, shifts each result right one bit, places each result into a vector, and writes the vector to the destination SIMD&FP register.
func VhsubqU32(v0 Uint32X4, v1 Uint32X4) Uint32X4 {
	return C.vhsubq_u32(v0, v1)
}

// Unsigned Halving Subtract. This instruction subtracts the vector elements in the second source SIMD&FP register from the corresponding vector elements in the first source SIMD&FP register, shifts each result right one bit, places each result into a vector, and writes the vector to the destination SIMD&FP register.
func VhsubqU16(v0 Uint16X8, v1 Uint16X8) Uint16X8 {
	return C.vhsubq_u16(v0, v1)
}

// Signed Halving Subtract. This instruction subtracts the elements in the vector in the second source SIMD&FP register from the corresponding elements in the vector in the first source SIMD&FP register, shifts each result right one bit, places each result into elements of a vector, and writes the vector to the destination SIMD&FP register.
func VhsubqS8(v0 Int8X16, v1 Int8X16) Int8X16 {
	return C.vhsubq_s8(v0, v1)
}

// Signed Halving Subtract. This instruction subtracts the elements in the vector in the second source SIMD&FP register from the corresponding elements in the vector in the first source SIMD&FP register, shifts each result right one bit, places each result into elements of a vector, and writes the vector to the destination SIMD&FP register.
func VhsubqS32(v0 Int32X4, v1 Int32X4) Int32X4 {
	return C.vhsubq_s32(v0, v1)
}

// Signed Halving Subtract. This instruction subtracts the elements in the vector in the second source SIMD&FP register from the corresponding elements in the vector in the first source SIMD&FP register, shifts each result right one bit, places each result into elements of a vector, and writes the vector to the destination SIMD&FP register.
func VhsubqS16(v0 Int16X8, v1 Int16X8) Int16X8 {
	return C.vhsubq_s16(v0, v1)
}

// Unsigned Halving Subtract. This instruction subtracts the vector elements in the second source SIMD&FP register from the corresponding vector elements in the first source SIMD&FP register, shifts each result right one bit, places each result into a vector, and writes the vector to the destination SIMD&FP register.
func VhsubU8(v0 Uint8X8, v1 Uint8X8) Uint8X8 {
	return C.vhsub_u8(v0, v1)
}

// Unsigned Halving Subtract. This instruction subtracts the vector elements in the second source SIMD&FP register from the corresponding vector elements in the first source SIMD&FP register, shifts each result right one bit, places each result into a vector, and writes the vector to the destination SIMD&FP register.
func VhsubU32(v0 Uint32X2, v1 Uint32X2) Uint32X2 {
	return C.vhsub_u32(v0, v1)
}

// Unsigned Halving Subtract. This instruction subtracts the vector elements in the second source SIMD&FP register from the corresponding vector elements in the first source SIMD&FP register, shifts each result right one bit, places each result into a vector, and writes the vector to the destination SIMD&FP register.
func VhsubU16(v0 Uint16X4, v1 Uint16X4) Uint16X4 {
	return C.vhsub_u16(v0, v1)
}

// Signed Halving Subtract. This instruction subtracts the elements in the vector in the second source SIMD&FP register from the corresponding elements in the vector in the first source SIMD&FP register, shifts each result right one bit, places each result into elements of a vector, and writes the vector to the destination SIMD&FP register.
func VhsubS8(v0 Int8X8, v1 Int8X8) Int8X8 {
	return C.vhsub_s8(v0, v1)
}

// Signed Halving Subtract. This instruction subtracts the elements in the vector in the second source SIMD&FP register from the corresponding elements in the vector in the first source SIMD&FP register, shifts each result right one bit, places each result into elements of a vector, and writes the vector to the destination SIMD&FP register.
func VhsubS32(v0 Int32X2, v1 Int32X2) Int32X2 {
	return C.vhsub_s32(v0, v1)
}

// Signed Halving Subtract. This instruction subtracts the elements in the vector in the second source SIMD&FP register from the corresponding elements in the vector in the first source SIMD&FP register, shifts each result right one bit, places each result into elements of a vector, and writes the vector to the destination SIMD&FP register.
func VhsubS16(v0 Int16X4, v1 Int16X4) Int16X4 {
	return C.vhsub_s16(v0, v1)
}

// Unsigned Maximum (vector). This instruction compares corresponding elements in the vectors in the two source SIMD&FP registers, places the larger of each pair of unsigned integer values into a vector, and writes the vector to the destination SIMD&FP register.
func VmaxqU8(v0 Uint8X16, v1 Uint8X16) Uint8X16 {
	return C.vmaxq_u8(v0, v1)
}

// Unsigned Maximum (vector). This instruction compares corresponding elements in the vectors in the two source SIMD&FP registers, places the larger of each pair of unsigned integer values into a vector, and writes the vector to the destination SIMD&FP register.
func VmaxqU32(v0 Uint32X4, v1 Uint32X4) Uint32X4 {
	return C.vmaxq_u32(v0, v1)
}

// Unsigned Maximum (vector). This instruction compares corresponding elements in the vectors in the two source SIMD&FP registers, places the larger of each pair of unsigned integer values into a vector, and writes the vector to the destination SIMD&FP register.
func VmaxqU16(v0 Uint16X8, v1 Uint16X8) Uint16X8 {
	return C.vmaxq_u16(v0, v1)
}

// Signed Maximum (vector). This instruction compares corresponding elements in the vectors in the two source SIMD&FP registers, places the larger of each pair of signed integer values into a vector, and writes the vector to the destination SIMD&FP register.
func VmaxqS8(v0 Int8X16, v1 Int8X16) Int8X16 {
	return C.vmaxq_s8(v0, v1)
}

// Floating-point Maximum (vector). This instruction compares corresponding vector elements in the two source SIMD&FP registers, places the larger of each of the two floating-point values into a vector, and writes the vector to the destination SIMD&FP register.
func VmaxqF32(v0 Float32X4, v1 Float32X4) Float32X4 {
	return C.vmaxq_f32(v0, v1)
}

// Signed Maximum (vector). This instruction compares corresponding elements in the vectors in the two source SIMD&FP registers, places the larger of each pair of signed integer values into a vector, and writes the vector to the destination SIMD&FP register.
func VmaxqS32(v0 Int32X4, v1 Int32X4) Int32X4 {
	return C.vmaxq_s32(v0, v1)
}

// Signed Maximum (vector). This instruction compares corresponding elements in the vectors in the two source SIMD&FP registers, places the larger of each pair of signed integer values into a vector, and writes the vector to the destination SIMD&FP register.
func VmaxqS16(v0 Int16X8, v1 Int16X8) Int16X8 {
	return C.vmaxq_s16(v0, v1)
}

// Unsigned Maximum (vector). This instruction compares corresponding elements in the vectors in the two source SIMD&FP registers, places the larger of each pair of unsigned integer values into a vector, and writes the vector to the destination SIMD&FP register.
func VmaxU8(v0 Uint8X8, v1 Uint8X8) Uint8X8 {
	return C.vmax_u8(v0, v1)
}

// Unsigned Maximum (vector). This instruction compares corresponding elements in the vectors in the two source SIMD&FP registers, places the larger of each pair of unsigned integer values into a vector, and writes the vector to the destination SIMD&FP register.
func VmaxU32(v0 Uint32X2, v1 Uint32X2) Uint32X2 {
	return C.vmax_u32(v0, v1)
}

// Unsigned Maximum (vector). This instruction compares corresponding elements in the vectors in the two source SIMD&FP registers, places the larger of each pair of unsigned integer values into a vector, and writes the vector to the destination SIMD&FP register.
func VmaxU16(v0 Uint16X4, v1 Uint16X4) Uint16X4 {
	return C.vmax_u16(v0, v1)
}

// Signed Maximum (vector). This instruction compares corresponding elements in the vectors in the two source SIMD&FP registers, places the larger of each pair of signed integer values into a vector, and writes the vector to the destination SIMD&FP register.
func VmaxS8(v0 Int8X8, v1 Int8X8) Int8X8 {
	return C.vmax_s8(v0, v1)
}

// Floating-point Maximum (vector). This instruction compares corresponding vector elements in the two source SIMD&FP registers, places the larger of each of the two floating-point values into a vector, and writes the vector to the destination SIMD&FP register.
func VmaxF32(v0 Float32X2, v1 Float32X2) Float32X2 {
	return C.vmax_f32(v0, v1)
}

// Signed Maximum (vector). This instruction compares corresponding elements in the vectors in the two source SIMD&FP registers, places the larger of each pair of signed integer values into a vector, and writes the vector to the destination SIMD&FP register.
func VmaxS32(v0 Int32X2, v1 Int32X2) Int32X2 {
	return C.vmax_s32(v0, v1)
}

// Signed Maximum (vector). This instruction compares corresponding elements in the vectors in the two source SIMD&FP registers, places the larger of each pair of signed integer values into a vector, and writes the vector to the destination SIMD&FP register.
func VmaxS16(v0 Int16X4, v1 Int16X4) Int16X4 {
	return C.vmax_s16(v0, v1)
}

// Unsigned Minimum (vector). This instruction compares corresponding vector elements in the two source SIMD&FP registers, places the smaller of each of the two unsigned integer values into a vector, and writes the vector to the destination SIMD&FP register.
func VminqU8(v0 Uint8X16, v1 Uint8X16) Uint8X16 {
	return C.vminq_u8(v0, v1)
}

// Unsigned Minimum (vector). This instruction compares corresponding vector elements in the two source SIMD&FP registers, places the smaller of each of the two unsigned integer values into a vector, and writes the vector to the destination SIMD&FP register.
func VminqU32(v0 Uint32X4, v1 Uint32X4) Uint32X4 {
	return C.vminq_u32(v0, v1)
}

// Unsigned Minimum (vector). This instruction compares corresponding vector elements in the two source SIMD&FP registers, places the smaller of each of the two unsigned integer values into a vector, and writes the vector to the destination SIMD&FP register.
func VminqU16(v0 Uint16X8, v1 Uint16X8) Uint16X8 {
	return C.vminq_u16(v0, v1)
}

// Signed Minimum (vector). This instruction compares corresponding elements in the vectors in the two source SIMD&FP registers, places the smaller of each of the two signed integer values into a vector, and writes the vector to the destination SIMD&FP register.
func VminqS8(v0 Int8X16, v1 Int8X16) Int8X16 {
	return C.vminq_s8(v0, v1)
}

// Floating-point minimum (vector). This instruction compares corresponding elements in the vectors in the two source SIMD&FP registers, places the smaller of each of the two floating-point values into a vector, and writes the vector to the destination SIMD&FP register.
func VminqF32(v0 Float32X4, v1 Float32X4) Float32X4 {
	return C.vminq_f32(v0, v1)
}

// Signed Minimum (vector). This instruction compares corresponding elements in the vectors in the two source SIMD&FP registers, places the smaller of each of the two signed integer values into a vector, and writes the vector to the destination SIMD&FP register.
func VminqS32(v0 Int32X4, v1 Int32X4) Int32X4 {
	return C.vminq_s32(v0, v1)
}

// Signed Minimum (vector). This instruction compares corresponding elements in the vectors in the two source SIMD&FP registers, places the smaller of each of the two signed integer values into a vector, and writes the vector to the destination SIMD&FP register.
func VminqS16(v0 Int16X8, v1 Int16X8) Int16X8 {
	return C.vminq_s16(v0, v1)
}

// Unsigned Minimum (vector). This instruction compares corresponding vector elements in the two source SIMD&FP registers, places the smaller of each of the two unsigned integer values into a vector, and writes the vector to the destination SIMD&FP register.
func VminU8(v0 Uint8X8, v1 Uint8X8) Uint8X8 {
	return C.vmin_u8(v0, v1)
}

// Unsigned Minimum (vector). This instruction compares corresponding vector elements in the two source SIMD&FP registers, places the smaller of each of the two unsigned integer values into a vector, and writes the vector to the destination SIMD&FP register.
func VminU32(v0 Uint32X2, v1 Uint32X2) Uint32X2 {
	return C.vmin_u32(v0, v1)
}

// Unsigned Minimum (vector). This instruction compares corresponding vector elements in the two source SIMD&FP registers, places the smaller of each of the two unsigned integer values into a vector, and writes the vector to the destination SIMD&FP register.
func VminU16(v0 Uint16X4, v1 Uint16X4) Uint16X4 {
	return C.vmin_u16(v0, v1)
}

// Signed Minimum (vector). This instruction compares corresponding elements in the vectors in the two source SIMD&FP registers, places the smaller of each of the two signed integer values into a vector, and writes the vector to the destination SIMD&FP register.
func VminS8(v0 Int8X8, v1 Int8X8) Int8X8 {
	return C.vmin_s8(v0, v1)
}

// Floating-point minimum (vector). This instruction compares corresponding elements in the vectors in the two source SIMD&FP registers, places the smaller of each of the two floating-point values into a vector, and writes the vector to the destination SIMD&FP register.
func VminF32(v0 Float32X2, v1 Float32X2) Float32X2 {
	return C.vmin_f32(v0, v1)
}

// Signed Minimum (vector). This instruction compares corresponding elements in the vectors in the two source SIMD&FP registers, places the smaller of each of the two signed integer values into a vector, and writes the vector to the destination SIMD&FP register.
func VminS32(v0 Int32X2, v1 Int32X2) Int32X2 {
	return C.vmin_s32(v0, v1)
}

// Signed Minimum (vector). This instruction compares corresponding elements in the vectors in the two source SIMD&FP registers, places the smaller of each of the two signed integer values into a vector, and writes the vector to the destination SIMD&FP register.
func VminS16(v0 Int16X4, v1 Int16X4) Int16X4 {
	return C.vmin_s16(v0, v1)
}

// Multiply-Add to accumulator (vector). This instruction multiplies corresponding elements in the vectors of the two source SIMD&FP registers, and accumulates the results with the vector elements of the destination SIMD&FP register.
func VmlaqU8(v0 Uint8X16, v1 Uint8X16, v2 Uint8X16) Uint8X16 {
	return C.vmlaq_u8(v0, v1, v2)
}

// Multiply-Add to accumulator (vector). This instruction multiplies corresponding elements in the vectors of the two source SIMD&FP registers, and accumulates the results with the vector elements of the destination SIMD&FP register.
func VmlaqU32(v0 Uint32X4, v1 Uint32X4, v2 Uint32X4) Uint32X4 {
	return C.vmlaq_u32(v0, v1, v2)
}

// Multiply-Add to accumulator (vector). This instruction multiplies corresponding elements in the vectors of the two source SIMD&FP registers, and accumulates the results with the vector elements of the destination SIMD&FP register.
func VmlaqU16(v0 Uint16X8, v1 Uint16X8, v2 Uint16X8) Uint16X8 {
	return C.vmlaq_u16(v0, v1, v2)
}

// Multiply-Add to accumulator (vector). This instruction multiplies corresponding elements in the vectors of the two source SIMD&FP registers, and accumulates the results with the vector elements of the destination SIMD&FP register.
func VmlaqS8(v0 Int8X16, v1 Int8X16, v2 Int8X16) Int8X16 {
	return C.vmlaq_s8(v0, v1, v2)
}

// Floating-point multiply-add to accumulator
func VmlaqF32(v0 Float32X4, v1 Float32X4, v2 Float32X4) Float32X4 {
	return C.vmlaq_f32(v0, v1, v2)
}

// Multiply-Add to accumulator (vector). This instruction multiplies corresponding elements in the vectors of the two source SIMD&FP registers, and accumulates the results with the vector elements of the destination SIMD&FP register.
func VmlaqS32(v0 Int32X4, v1 Int32X4, v2 Int32X4) Int32X4 {
	return C.vmlaq_s32(v0, v1, v2)
}

// Multiply-Add to accumulator (vector). This instruction multiplies corresponding elements in the vectors of the two source SIMD&FP registers, and accumulates the results with the vector elements of the destination SIMD&FP register.
func VmlaqS16(v0 Int16X8, v1 Int16X8, v2 Int16X8) Int16X8 {
	return C.vmlaq_s16(v0, v1, v2)
}

// Multiply-Add to accumulator (vector). This instruction multiplies corresponding elements in the vectors of the two source SIMD&FP registers, and accumulates the results with the vector elements of the destination SIMD&FP register.
func VmlaU8(v0 Uint8X8, v1 Uint8X8, v2 Uint8X8) Uint8X8 {
	return C.vmla_u8(v0, v1, v2)
}

// Multiply-Add to accumulator (vector). This instruction multiplies corresponding elements in the vectors of the two source SIMD&FP registers, and accumulates the results with the vector elements of the destination SIMD&FP register.
func VmlaU32(v0 Uint32X2, v1 Uint32X2, v2 Uint32X2) Uint32X2 {
	return C.vmla_u32(v0, v1, v2)
}

// Multiply-Add to accumulator (vector). This instruction multiplies corresponding elements in the vectors of the two source SIMD&FP registers, and accumulates the results with the vector elements of the destination SIMD&FP register.
func VmlaU16(v0 Uint16X4, v1 Uint16X4, v2 Uint16X4) Uint16X4 {
	return C.vmla_u16(v0, v1, v2)
}

// Multiply-Add to accumulator (vector). This instruction multiplies corresponding elements in the vectors of the two source SIMD&FP registers, and accumulates the results with the vector elements of the destination SIMD&FP register.
func VmlaS8(v0 Int8X8, v1 Int8X8, v2 Int8X8) Int8X8 {
	return C.vmla_s8(v0, v1, v2)
}

// Floating-point multiply-add to accumulator
func VmlaF32(v0 Float32X2, v1 Float32X2, v2 Float32X2) Float32X2 {
	return C.vmla_f32(v0, v1, v2)
}

// Multiply-Add to accumulator (vector). This instruction multiplies corresponding elements in the vectors of the two source SIMD&FP registers, and accumulates the results with the vector elements of the destination SIMD&FP register.
func VmlaS32(v0 Int32X2, v1 Int32X2, v2 Int32X2) Int32X2 {
	return C.vmla_s32(v0, v1, v2)
}

// Multiply-Add to accumulator (vector). This instruction multiplies corresponding elements in the vectors of the two source SIMD&FP registers, and accumulates the results with the vector elements of the destination SIMD&FP register.
func VmlaS16(v0 Int16X4, v1 Int16X4, v2 Int16X4) Int16X4 {
	return C.vmla_s16(v0, v1, v2)
}

// Vector multiply accumulate with scalar
func VmlaqNU32(v0 Uint32X4, v1 Uint32X4, v2 Uint32) Uint32X4 {
	return C.vmlaq_n_u32(v0, v1, v2)
}

// Vector multiply accumulate with scalar
func VmlaqNU16(v0 Uint16X8, v1 Uint16X8, v2 Uint16) Uint16X8 {
	return C.vmlaq_n_u16(v0, v1, v2)
}

// Vector multiply accumulate with scalar
func VmlaqNF32(v0 Float32X4, v1 Float32X4, v2 Float32) Float32X4 {
	return C.vmlaq_n_f32(v0, v1, v2)
}

// Vector multiply accumulate with scalar
func VmlaqNS32(v0 Int32X4, v1 Int32X4, v2 Int32) Int32X4 {
	return C.vmlaq_n_s32(v0, v1, v2)
}

// Vector multiply accumulate with scalar
func VmlaqNS16(v0 Int16X8, v1 Int16X8, v2 Int16) Int16X8 {
	return C.vmlaq_n_s16(v0, v1, v2)
}

// Vector multiply accumulate with scalar
func VmlaNU32(v0 Uint32X2, v1 Uint32X2, v2 Uint32) Uint32X2 {
	return C.vmla_n_u32(v0, v1, v2)
}

// Vector multiply accumulate with scalar
func VmlaNU16(v0 Uint16X4, v1 Uint16X4, v2 Uint16) Uint16X4 {
	return C.vmla_n_u16(v0, v1, v2)
}

// Vector multiply accumulate with scalar
func VmlaNF32(v0 Float32X2, v1 Float32X2, v2 Float32) Float32X2 {
	return C.vmla_n_f32(v0, v1, v2)
}

// Vector multiply accumulate with scalar
func VmlaNS32(v0 Int32X2, v1 Int32X2, v2 Int32) Int32X2 {
	return C.vmla_n_s32(v0, v1, v2)
}

// Vector multiply accumulate with scalar
func VmlaNS16(v0 Int16X4, v1 Int16X4, v2 Int16) Int16X4 {
	return C.vmla_n_s16(v0, v1, v2)
}

// Multiply-Subtract from accumulator (vector). This instruction multiplies corresponding elements in the vectors of the two source SIMD&FP registers, and subtracts the results from the vector elements of the destination SIMD&FP register.
func VmlsqU8(v0 Uint8X16, v1 Uint8X16, v2 Uint8X16) Uint8X16 {
	return C.vmlsq_u8(v0, v1, v2)
}

// Multiply-Subtract from accumulator (vector). This instruction multiplies corresponding elements in the vectors of the two source SIMD&FP registers, and subtracts the results from the vector elements of the destination SIMD&FP register.
func VmlsqU32(v0 Uint32X4, v1 Uint32X4, v2 Uint32X4) Uint32X4 {
	return C.vmlsq_u32(v0, v1, v2)
}

// Multiply-Subtract from accumulator (vector). This instruction multiplies corresponding elements in the vectors of the two source SIMD&FP registers, and subtracts the results from the vector elements of the destination SIMD&FP register.
func VmlsqU16(v0 Uint16X8, v1 Uint16X8, v2 Uint16X8) Uint16X8 {
	return C.vmlsq_u16(v0, v1, v2)
}

// Multiply-Subtract from accumulator (vector). This instruction multiplies corresponding elements in the vectors of the two source SIMD&FP registers, and subtracts the results from the vector elements of the destination SIMD&FP register.
func VmlsqS8(v0 Int8X16, v1 Int8X16, v2 Int8X16) Int8X16 {
	return C.vmlsq_s8(v0, v1, v2)
}

// Multiply-subtract from accumulator
func VmlsqF32(v0 Float32X4, v1 Float32X4, v2 Float32X4) Float32X4 {
	return C.vmlsq_f32(v0, v1, v2)
}

// Multiply-Subtract from accumulator (vector). This instruction multiplies corresponding elements in the vectors of the two source SIMD&FP registers, and subtracts the results from the vector elements of the destination SIMD&FP register.
func VmlsqS32(v0 Int32X4, v1 Int32X4, v2 Int32X4) Int32X4 {
	return C.vmlsq_s32(v0, v1, v2)
}

// Multiply-Subtract from accumulator (vector). This instruction multiplies corresponding elements in the vectors of the two source SIMD&FP registers, and subtracts the results from the vector elements of the destination SIMD&FP register.
func VmlsqS16(v0 Int16X8, v1 Int16X8, v2 Int16X8) Int16X8 {
	return C.vmlsq_s16(v0, v1, v2)
}

// Multiply-Subtract from accumulator (vector). This instruction multiplies corresponding elements in the vectors of the two source SIMD&FP registers, and subtracts the results from the vector elements of the destination SIMD&FP register.
func VmlsU8(v0 Uint8X8, v1 Uint8X8, v2 Uint8X8) Uint8X8 {
	return C.vmls_u8(v0, v1, v2)
}

// Multiply-Subtract from accumulator (vector). This instruction multiplies corresponding elements in the vectors of the two source SIMD&FP registers, and subtracts the results from the vector elements of the destination SIMD&FP register.
func VmlsU32(v0 Uint32X2, v1 Uint32X2, v2 Uint32X2) Uint32X2 {
	return C.vmls_u32(v0, v1, v2)
}

// Multiply-Subtract from accumulator (vector). This instruction multiplies corresponding elements in the vectors of the two source SIMD&FP registers, and subtracts the results from the vector elements of the destination SIMD&FP register.
func VmlsU16(v0 Uint16X4, v1 Uint16X4, v2 Uint16X4) Uint16X4 {
	return C.vmls_u16(v0, v1, v2)
}

// Multiply-Subtract from accumulator (vector). This instruction multiplies corresponding elements in the vectors of the two source SIMD&FP registers, and subtracts the results from the vector elements of the destination SIMD&FP register.
func VmlsS8(v0 Int8X8, v1 Int8X8, v2 Int8X8) Int8X8 {
	return C.vmls_s8(v0, v1, v2)
}

// Multiply-subtract from accumulator
func VmlsF32(v0 Float32X2, v1 Float32X2, v2 Float32X2) Float32X2 {
	return C.vmls_f32(v0, v1, v2)
}

// Multiply-Subtract from accumulator (vector). This instruction multiplies corresponding elements in the vectors of the two source SIMD&FP registers, and subtracts the results from the vector elements of the destination SIMD&FP register.
func VmlsS32(v0 Int32X2, v1 Int32X2, v2 Int32X2) Int32X2 {
	return C.vmls_s32(v0, v1, v2)
}

// Multiply-Subtract from accumulator (vector). This instruction multiplies corresponding elements in the vectors of the two source SIMD&FP registers, and subtracts the results from the vector elements of the destination SIMD&FP register.
func VmlsS16(v0 Int16X4, v1 Int16X4, v2 Int16X4) Int16X4 {
	return C.vmls_s16(v0, v1, v2)
}

// Vector multiply subtract with scalar
func VmlsqNU32(v0 Uint32X4, v1 Uint32X4, v2 Uint32) Uint32X4 {
	return C.vmlsq_n_u32(v0, v1, v2)
}

// Vector multiply subtract with scalar
func VmlsqNU16(v0 Uint16X8, v1 Uint16X8, v2 Uint16) Uint16X8 {
	return C.vmlsq_n_u16(v0, v1, v2)
}

// Vector multiply subtract with scalar
func VmlsqNF32(v0 Float32X4, v1 Float32X4, v2 Float32) Float32X4 {
	return C.vmlsq_n_f32(v0, v1, v2)
}

// Vector multiply subtract with scalar
func VmlsqNS32(v0 Int32X4, v1 Int32X4, v2 Int32) Int32X4 {
	return C.vmlsq_n_s32(v0, v1, v2)
}

// Vector multiply subtract with scalar
func VmlsqNS16(v0 Int16X8, v1 Int16X8, v2 Int16) Int16X8 {
	return C.vmlsq_n_s16(v0, v1, v2)
}

// Vector multiply subtract with scalar
func VmlsNU32(v0 Uint32X2, v1 Uint32X2, v2 Uint32) Uint32X2 {
	return C.vmls_n_u32(v0, v1, v2)
}

// Vector multiply subtract with scalar
func VmlsNU16(v0 Uint16X4, v1 Uint16X4, v2 Uint16) Uint16X4 {
	return C.vmls_n_u16(v0, v1, v2)
}

// Vector multiply subtract with scalar
func VmlsNF32(v0 Float32X2, v1 Float32X2, v2 Float32) Float32X2 {
	return C.vmls_n_f32(v0, v1, v2)
}

// Vector multiply subtract with scalar
func VmlsNS32(v0 Int32X2, v1 Int32X2, v2 Int32) Int32X2 {
	return C.vmls_n_s32(v0, v1, v2)
}

// Vector multiply subtract with scalar
func VmlsNS16(v0 Int16X4, v1 Int16X4, v2 Int16) Int16X4 {
	return C.vmls_n_s16(v0, v1, v2)
}

// Duplicate vector element to vector or scalar. This instruction duplicates the vector element at the specified element index in the source SIMD&FP register into a scalar or each element in a vector, and writes the result to the destination SIMD&FP register.
func VmovNP8(v0 Poly8) Poly8X8 {
	return C.vmov_n_p8(v0)
}

// Duplicate vector element to vector or scalar. This instruction duplicates the vector element at the specified element index in the source SIMD&FP register into a scalar or each element in a vector, and writes the result to the destination SIMD&FP register.
func VmovNP16(v0 Poly16) Poly16X4 {
	return C.vmov_n_p16(v0)
}

// Duplicate vector element to vector or scalar. This instruction duplicates the vector element at the specified element index in the source SIMD&FP register into a scalar or each element in a vector, and writes the result to the destination SIMD&FP register.
func VmovqNP8(v0 Poly8) Poly8X16 {
	return C.vmovq_n_p8(v0)
}

// Duplicate vector element to vector or scalar. This instruction duplicates the vector element at the specified element index in the source SIMD&FP register into a scalar or each element in a vector, and writes the result to the destination SIMD&FP register.
func VmovqNP16(v0 Poly16) Poly16X8 {
	return C.vmovq_n_p16(v0)
}

// Duplicate vector element to vector or scalar. This instruction duplicates the vector element at the specified element index in the source SIMD&FP register into a scalar or each element in a vector, and writes the result to the destination SIMD&FP register.
func VmovqNU8(v0 Uint8) Uint8X16 {
	return C.vmovq_n_u8(v0)
}

// Duplicate vector element to vector or scalar. This instruction duplicates the vector element at the specified element index in the source SIMD&FP register into a scalar or each element in a vector, and writes the result to the destination SIMD&FP register.
func VmovqNU32(v0 Uint32) Uint32X4 {
	return C.vmovq_n_u32(v0)
}

// Duplicate vector element to vector or scalar. This instruction duplicates the vector element at the specified element index in the source SIMD&FP register into a scalar or each element in a vector, and writes the result to the destination SIMD&FP register.
func VmovqNU64(v0 Uint64) Uint64X2 {
	return C.vmovq_n_u64(v0)
}

// Duplicate vector element to vector or scalar. This instruction duplicates the vector element at the specified element index in the source SIMD&FP register into a scalar or each element in a vector, and writes the result to the destination SIMD&FP register.
func VmovqNU16(v0 Uint16) Uint16X8 {
	return C.vmovq_n_u16(v0)
}

// Duplicate vector element to vector or scalar. This instruction duplicates the vector element at the specified element index in the source SIMD&FP register into a scalar or each element in a vector, and writes the result to the destination SIMD&FP register.
func VmovqNS8(v0 Int8) Int8X16 {
	return C.vmovq_n_s8(v0)
}

// Duplicate vector element to vector or scalar. This instruction duplicates the vector element at the specified element index in the source SIMD&FP register into a scalar or each element in a vector, and writes the result to the destination SIMD&FP register.
func VmovqNF32(v0 Float32) Float32X4 {
	return C.vmovq_n_f32(v0)
}

// Duplicate vector element to vector or scalar. This instruction duplicates the vector element at the specified element index in the source SIMD&FP register into a scalar or each element in a vector, and writes the result to the destination SIMD&FP register.
func VmovqNS32(v0 Int32) Int32X4 {
	return C.vmovq_n_s32(v0)
}

// Duplicate vector element to vector or scalar. This instruction duplicates the vector element at the specified element index in the source SIMD&FP register into a scalar or each element in a vector, and writes the result to the destination SIMD&FP register.
func VmovqNS64(v0 Int64) Int64X2 {
	return C.vmovq_n_s64(v0)
}

// Duplicate vector element to vector or scalar. This instruction duplicates the vector element at the specified element index in the source SIMD&FP register into a scalar or each element in a vector, and writes the result to the destination SIMD&FP register.
func VmovqNS16(v0 Int16) Int16X8 {
	return C.vmovq_n_s16(v0)
}

// Duplicate vector element to vector or scalar. This instruction duplicates the vector element at the specified element index in the source SIMD&FP register into a scalar or each element in a vector, and writes the result to the destination SIMD&FP register.
func VmovNU8(v0 Uint8) Uint8X8 {
	return C.vmov_n_u8(v0)
}

// Duplicate vector element to vector or scalar. This instruction duplicates the vector element at the specified element index in the source SIMD&FP register into a scalar or each element in a vector, and writes the result to the destination SIMD&FP register.
func VmovNU32(v0 Uint32) Uint32X2 {
	return C.vmov_n_u32(v0)
}

// Duplicate vector element to vector or scalar. This instruction duplicates the vector element at the specified element index in the source SIMD&FP register into a scalar or each element in a vector, and writes the result to the destination SIMD&FP register.
func VmovNU64(v0 Uint64) Uint64X1 {
	return C.vmov_n_u64(v0)
}

// Duplicate vector element to vector or scalar. This instruction duplicates the vector element at the specified element index in the source SIMD&FP register into a scalar or each element in a vector, and writes the result to the destination SIMD&FP register.
func VmovNU16(v0 Uint16) Uint16X4 {
	return C.vmov_n_u16(v0)
}

// Duplicate vector element to vector or scalar. This instruction duplicates the vector element at the specified element index in the source SIMD&FP register into a scalar or each element in a vector, and writes the result to the destination SIMD&FP register.
func VmovNS8(v0 Int8) Int8X8 {
	return C.vmov_n_s8(v0)
}

// Duplicate vector element to vector or scalar. This instruction duplicates the vector element at the specified element index in the source SIMD&FP register into a scalar or each element in a vector, and writes the result to the destination SIMD&FP register.
func VmovNF32(v0 Float32) Float32X2 {
	return C.vmov_n_f32(v0)
}

// Duplicate vector element to vector or scalar. This instruction duplicates the vector element at the specified element index in the source SIMD&FP register into a scalar or each element in a vector, and writes the result to the destination SIMD&FP register.
func VmovNS32(v0 Int32) Int32X2 {
	return C.vmov_n_s32(v0)
}

// Duplicate vector element to vector or scalar. This instruction duplicates the vector element at the specified element index in the source SIMD&FP register into a scalar or each element in a vector, and writes the result to the destination SIMD&FP register.
func VmovNS64(v0 Int64) Int64X1 {
	return C.vmov_n_s64(v0)
}

// Duplicate vector element to vector or scalar. This instruction duplicates the vector element at the specified element index in the source SIMD&FP register into a scalar or each element in a vector, and writes the result to the destination SIMD&FP register.
func VmovNS16(v0 Int16) Int16X4 {
	return C.vmov_n_s16(v0)
}

// Vector move
func VmovlU8(v0 Uint8X8) Uint16X8 {
	return C.vmovl_u8(v0)
}

// Vector move
func VmovlU32(v0 Uint32X2) Uint64X2 {
	return C.vmovl_u32(v0)
}

// Vector move
func VmovlU16(v0 Uint16X4) Uint32X4 {
	return C.vmovl_u16(v0)
}

// Vector move
func VmovlS8(v0 Int8X8) Int16X8 {
	return C.vmovl_s8(v0)
}

// Vector move
func VmovlS32(v0 Int32X2) Int64X2 {
	return C.vmovl_s32(v0)
}

// Vector move
func VmovlS16(v0 Int16X4) Int32X4 {
	return C.vmovl_s16(v0)
}

// Extract Narrow. This instruction reads each vector element from the source SIMD&FP register, narrows each value to half the original width, places the result into a vector, and writes the vector to the lower or upper half of the destination SIMD&FP register. The destination vector elements are half as long as the source vector elements.
func VmovnU32(v0 Uint32X4) Uint16X4 {
	return C.vmovn_u32(v0)
}

// Extract Narrow. This instruction reads each vector element from the source SIMD&FP register, narrows each value to half the original width, places the result into a vector, and writes the vector to the lower or upper half of the destination SIMD&FP register. The destination vector elements are half as long as the source vector elements.
func VmovnU64(v0 Uint64X2) Uint32X2 {
	return C.vmovn_u64(v0)
}

// Extract Narrow. This instruction reads each vector element from the source SIMD&FP register, narrows each value to half the original width, places the result into a vector, and writes the vector to the lower or upper half of the destination SIMD&FP register. The destination vector elements are half as long as the source vector elements.
func VmovnU16(v0 Uint16X8) Uint8X8 {
	return C.vmovn_u16(v0)
}

// Extract Narrow. This instruction reads each vector element from the source SIMD&FP register, narrows each value to half the original width, places the result into a vector, and writes the vector to the lower or upper half of the destination SIMD&FP register. The destination vector elements are half as long as the source vector elements.
func VmovnS32(v0 Int32X4) Int16X4 {
	return C.vmovn_s32(v0)
}

// Extract Narrow. This instruction reads each vector element from the source SIMD&FP register, narrows each value to half the original width, places the result into a vector, and writes the vector to the lower or upper half of the destination SIMD&FP register. The destination vector elements are half as long as the source vector elements.
func VmovnS64(v0 Int64X2) Int32X2 {
	return C.vmovn_s64(v0)
}

// Extract Narrow. This instruction reads each vector element from the source SIMD&FP register, narrows each value to half the original width, places the result into a vector, and writes the vector to the lower or upper half of the destination SIMD&FP register. The destination vector elements are half as long as the source vector elements.
func VmovnS16(v0 Int16X8) Int8X8 {
	return C.vmovn_s16(v0)
}

// Multiply (vector). This instruction multiplies corresponding elements in the vectors of the two source SIMD&FP registers, places the results in a vector, and writes the vector to the destination SIMD&FP register.
func VmulqU8(v0 Uint8X16, v1 Uint8X16) Uint8X16 {
	return C.vmulq_u8(v0, v1)
}

// Multiply (vector). This instruction multiplies corresponding elements in the vectors of the two source SIMD&FP registers, places the results in a vector, and writes the vector to the destination SIMD&FP register.
func VmulqU32(v0 Uint32X4, v1 Uint32X4) Uint32X4 {
	return C.vmulq_u32(v0, v1)
}

// Multiply (vector). This instruction multiplies corresponding elements in the vectors of the two source SIMD&FP registers, places the results in a vector, and writes the vector to the destination SIMD&FP register.
func VmulqU16(v0 Uint16X8, v1 Uint16X8) Uint16X8 {
	return C.vmulq_u16(v0, v1)
}

// Multiply (vector). This instruction multiplies corresponding elements in the vectors of the two source SIMD&FP registers, places the results in a vector, and writes the vector to the destination SIMD&FP register.
func VmulqS8(v0 Int8X16, v1 Int8X16) Int8X16 {
	return C.vmulq_s8(v0, v1)
}

// Floating-point Multiply (vector). This instruction multiplies corresponding floating-point values in the vectors in the two source SIMD&FP registers, places the result in a vector, and writes the vector to the destination SIMD&FP register.
func VmulqF32(v0 Float32X4, v1 Float32X4) Float32X4 {
	return C.vmulq_f32(v0, v1)
}

// Multiply (vector). This instruction multiplies corresponding elements in the vectors of the two source SIMD&FP registers, places the results in a vector, and writes the vector to the destination SIMD&FP register.
func VmulqS32(v0 Int32X4, v1 Int32X4) Int32X4 {
	return C.vmulq_s32(v0, v1)
}

// Multiply (vector). This instruction multiplies corresponding elements in the vectors of the two source SIMD&FP registers, places the results in a vector, and writes the vector to the destination SIMD&FP register.
func VmulqS16(v0 Int16X8, v1 Int16X8) Int16X8 {
	return C.vmulq_s16(v0, v1)
}

// Multiply (vector). This instruction multiplies corresponding elements in the vectors of the two source SIMD&FP registers, places the results in a vector, and writes the vector to the destination SIMD&FP register.
func VmulU8(v0 Uint8X8, v1 Uint8X8) Uint8X8 {
	return C.vmul_u8(v0, v1)
}

// Multiply (vector). This instruction multiplies corresponding elements in the vectors of the two source SIMD&FP registers, places the results in a vector, and writes the vector to the destination SIMD&FP register.
func VmulU32(v0 Uint32X2, v1 Uint32X2) Uint32X2 {
	return C.vmul_u32(v0, v1)
}

// Multiply (vector). This instruction multiplies corresponding elements in the vectors of the two source SIMD&FP registers, places the results in a vector, and writes the vector to the destination SIMD&FP register.
func VmulU16(v0 Uint16X4, v1 Uint16X4) Uint16X4 {
	return C.vmul_u16(v0, v1)
}

// Multiply (vector). This instruction multiplies corresponding elements in the vectors of the two source SIMD&FP registers, places the results in a vector, and writes the vector to the destination SIMD&FP register.
func VmulS8(v0 Int8X8, v1 Int8X8) Int8X8 {
	return C.vmul_s8(v0, v1)
}

// Floating-point Multiply (vector). This instruction multiplies corresponding floating-point values in the vectors in the two source SIMD&FP registers, places the result in a vector, and writes the vector to the destination SIMD&FP register.
func VmulF32(v0 Float32X2, v1 Float32X2) Float32X2 {
	return C.vmul_f32(v0, v1)
}

// Multiply (vector). This instruction multiplies corresponding elements in the vectors of the two source SIMD&FP registers, places the results in a vector, and writes the vector to the destination SIMD&FP register.
func VmulS32(v0 Int32X2, v1 Int32X2) Int32X2 {
	return C.vmul_s32(v0, v1)
}

// Multiply (vector). This instruction multiplies corresponding elements in the vectors of the two source SIMD&FP registers, places the results in a vector, and writes the vector to the destination SIMD&FP register.
func VmulS16(v0 Int16X4, v1 Int16X4) Int16X4 {
	return C.vmul_s16(v0, v1)
}

// Polynomial Multiply. This instruction multiplies corresponding elements in the vectors of the two source SIMD&FP registers, places the results in a vector, and writes the vector to the destination SIMD&FP register.
func VmulP8(v0 Poly8X8, v1 Poly8X8) Poly8X8 {
	return C.vmul_p8(v0, v1)
}

// Polynomial Multiply. This instruction multiplies corresponding elements in the vectors of the two source SIMD&FP registers, places the results in a vector, and writes the vector to the destination SIMD&FP register.
func VmulqP8(v0 Poly8X16, v1 Poly8X16) Poly8X16 {
	return C.vmulq_p8(v0, v1)
}

// Vector multiply by scalar
func VmulqNU32(v0 Uint32X4, v1 Uint32) Uint32X4 {
	return C.vmulq_n_u32(v0, v1)
}

// Vector multiply by scalar
func VmulqNU16(v0 Uint16X8, v1 Uint16) Uint16X8 {
	return C.vmulq_n_u16(v0, v1)
}

// Vector multiply by scalar
func VmulqNF32(v0 Float32X4, v1 Float32) Float32X4 {
	return C.vmulq_n_f32(v0, v1)
}

// Vector multiply by scalar
func VmulqNS32(v0 Int32X4, v1 Int32) Int32X4 {
	return C.vmulq_n_s32(v0, v1)
}

// Vector multiply by scalar
func VmulqNS16(v0 Int16X8, v1 Int16) Int16X8 {
	return C.vmulq_n_s16(v0, v1)
}

// Vector multiply by scalar
func VmulNU32(v0 Uint32X2, v1 Uint32) Uint32X2 {
	return C.vmul_n_u32(v0, v1)
}

// Vector multiply by scalar
func VmulNU16(v0 Uint16X4, v1 Uint16) Uint16X4 {
	return C.vmul_n_u16(v0, v1)
}

// Vector multiply by scalar
func VmulNF32(v0 Float32X2, v1 Float32) Float32X2 {
	return C.vmul_n_f32(v0, v1)
}

// Vector multiply by scalar
func VmulNS32(v0 Int32X2, v1 Int32) Int32X2 {
	return C.vmul_n_s32(v0, v1)
}

// Vector multiply by scalar
func VmulNS16(v0 Int16X4, v1 Int16) Int16X4 {
	return C.vmul_n_s16(v0, v1)
}

// Polynomial Multiply Long. This instruction multiplies corresponding elements in the lower or upper half of the vectors of the two source SIMD&FP registers, places the results in a vector, and writes the vector to the destination SIMD&FP register. The destination vector elements are twice as long as the elements that are multiplied.
func VmullP8(v0 Poly8X8, v1 Poly8X8) Poly16X8 {
	return C.vmull_p8(v0, v1)
}

// Unsigned Multiply long (vector). This instruction multiplies corresponding vector elements in the lower or upper half of the two source SIMD&FP registers, places the result in a vector, and writes the vector to the destination SIMD&FP register. The destination vector elements are twice as long as the elements that are multiplied. All the values in this instruction are unsigned integer values.
func VmullU8(v0 Uint8X8, v1 Uint8X8) Uint16X8 {
	return C.vmull_u8(v0, v1)
}

// Unsigned Multiply long (vector). This instruction multiplies corresponding vector elements in the lower or upper half of the two source SIMD&FP registers, places the result in a vector, and writes the vector to the destination SIMD&FP register. The destination vector elements are twice as long as the elements that are multiplied. All the values in this instruction are unsigned integer values.
func VmullU32(v0 Uint32X2, v1 Uint32X2) Uint64X2 {
	return C.vmull_u32(v0, v1)
}

// Unsigned Multiply long (vector). This instruction multiplies corresponding vector elements in the lower or upper half of the two source SIMD&FP registers, places the result in a vector, and writes the vector to the destination SIMD&FP register. The destination vector elements are twice as long as the elements that are multiplied. All the values in this instruction are unsigned integer values.
func VmullU16(v0 Uint16X4, v1 Uint16X4) Uint32X4 {
	return C.vmull_u16(v0, v1)
}

// Signed Multiply Long (vector). This instruction multiplies corresponding signed integer values in the lower or upper half of the vectors of the two source SIMD&FP registers, places the results in a vector, and writes the vector to the destination SIMD&FP register.
func VmullS8(v0 Int8X8, v1 Int8X8) Int16X8 {
	return C.vmull_s8(v0, v1)
}

// Signed Multiply Long (vector). This instruction multiplies corresponding signed integer values in the lower or upper half of the vectors of the two source SIMD&FP registers, places the results in a vector, and writes the vector to the destination SIMD&FP register.
func VmullS32(v0 Int32X2, v1 Int32X2) Int64X2 {
	return C.vmull_s32(v0, v1)
}

// Signed Multiply Long (vector). This instruction multiplies corresponding signed integer values in the lower or upper half of the vectors of the two source SIMD&FP registers, places the results in a vector, and writes the vector to the destination SIMD&FP register.
func VmullS16(v0 Int16X4, v1 Int16X4) Int32X4 {
	return C.vmull_s16(v0, v1)
}

// Vector long multiply with scalar
func VmullNU32(v0 Uint32X2, v1 Uint32) Uint64X2 {
	return C.vmull_n_u32(v0, v1)
}

// Vector long multiply with scalar
func VmullNU16(v0 Uint16X4, v1 Uint16) Uint32X4 {
	return C.vmull_n_u16(v0, v1)
}

// Vector long multiply with scalar
func VmullNS32(v0 Int32X2, v1 Int32) Int64X2 {
	return C.vmull_n_s32(v0, v1)
}

// Vector long multiply with scalar
func VmullNS16(v0 Int16X4, v1 Int16) Int32X4 {
	return C.vmull_n_s16(v0, v1)
}

// Bitwise NOT (vector). This instruction reads each vector element from the source SIMD&FP register, places the inverse of each value into a vector, and writes the vector to the destination SIMD&FP register.
func VmvnP8(v0 Poly8X8) Poly8X8 {
	return C.vmvn_p8(v0)
}

// Bitwise NOT (vector). This instruction reads each vector element from the source SIMD&FP register, places the inverse of each value into a vector, and writes the vector to the destination SIMD&FP register.
func VmvnqP8(v0 Poly8X16) Poly8X16 {
	return C.vmvnq_p8(v0)
}

// Bitwise NOT (vector). This instruction reads each vector element from the source SIMD&FP register, places the inverse of each value into a vector, and writes the vector to the destination SIMD&FP register.
func VmvnqU8(v0 Uint8X16) Uint8X16 {
	return C.vmvnq_u8(v0)
}

// Bitwise NOT (vector). This instruction reads each vector element from the source SIMD&FP register, places the inverse of each value into a vector, and writes the vector to the destination SIMD&FP register.
func VmvnqU32(v0 Uint32X4) Uint32X4 {
	return C.vmvnq_u32(v0)
}

// Bitwise NOT (vector). This instruction reads each vector element from the source SIMD&FP register, places the inverse of each value into a vector, and writes the vector to the destination SIMD&FP register.
func VmvnqU16(v0 Uint16X8) Uint16X8 {
	return C.vmvnq_u16(v0)
}

// Bitwise NOT (vector). This instruction reads each vector element from the source SIMD&FP register, places the inverse of each value into a vector, and writes the vector to the destination SIMD&FP register.
func VmvnqS8(v0 Int8X16) Int8X16 {
	return C.vmvnq_s8(v0)
}

// Bitwise NOT (vector). This instruction reads each vector element from the source SIMD&FP register, places the inverse of each value into a vector, and writes the vector to the destination SIMD&FP register.
func VmvnqS32(v0 Int32X4) Int32X4 {
	return C.vmvnq_s32(v0)
}

// Bitwise NOT (vector). This instruction reads each vector element from the source SIMD&FP register, places the inverse of each value into a vector, and writes the vector to the destination SIMD&FP register.
func VmvnqS16(v0 Int16X8) Int16X8 {
	return C.vmvnq_s16(v0)
}

// Bitwise NOT (vector). This instruction reads each vector element from the source SIMD&FP register, places the inverse of each value into a vector, and writes the vector to the destination SIMD&FP register.
func VmvnU8(v0 Uint8X8) Uint8X8 {
	return C.vmvn_u8(v0)
}

// Bitwise NOT (vector). This instruction reads each vector element from the source SIMD&FP register, places the inverse of each value into a vector, and writes the vector to the destination SIMD&FP register.
func VmvnU32(v0 Uint32X2) Uint32X2 {
	return C.vmvn_u32(v0)
}

// Bitwise NOT (vector). This instruction reads each vector element from the source SIMD&FP register, places the inverse of each value into a vector, and writes the vector to the destination SIMD&FP register.
func VmvnU16(v0 Uint16X4) Uint16X4 {
	return C.vmvn_u16(v0)
}

// Bitwise NOT (vector). This instruction reads each vector element from the source SIMD&FP register, places the inverse of each value into a vector, and writes the vector to the destination SIMD&FP register.
func VmvnS8(v0 Int8X8) Int8X8 {
	return C.vmvn_s8(v0)
}

// Bitwise NOT (vector). This instruction reads each vector element from the source SIMD&FP register, places the inverse of each value into a vector, and writes the vector to the destination SIMD&FP register.
func VmvnS32(v0 Int32X2) Int32X2 {
	return C.vmvn_s32(v0)
}

// Bitwise NOT (vector). This instruction reads each vector element from the source SIMD&FP register, places the inverse of each value into a vector, and writes the vector to the destination SIMD&FP register.
func VmvnS16(v0 Int16X4) Int16X4 {
	return C.vmvn_s16(v0)
}

// Negate (vector). This instruction reads each vector element from the source SIMD&FP register, negates each value, puts the result into a vector, and writes the vector to the destination SIMD&FP register.
func VnegqS8(v0 Int8X16) Int8X16 {
	return C.vnegq_s8(v0)
}

// Floating-point Negate (vector). This instruction negates the value of each vector element in the source SIMD&FP register, writes the result to a vector, and writes the vector to the destination SIMD&FP register.
func VnegqF32(v0 Float32X4) Float32X4 {
	return C.vnegq_f32(v0)
}

// Negate (vector). This instruction reads each vector element from the source SIMD&FP register, negates each value, puts the result into a vector, and writes the vector to the destination SIMD&FP register.
func VnegqS32(v0 Int32X4) Int32X4 {
	return C.vnegq_s32(v0)
}

// Negate (vector). This instruction reads each vector element from the source SIMD&FP register, negates each value, puts the result into a vector, and writes the vector to the destination SIMD&FP register.
func VnegqS16(v0 Int16X8) Int16X8 {
	return C.vnegq_s16(v0)
}

// Negate (vector). This instruction reads each vector element from the source SIMD&FP register, negates each value, puts the result into a vector, and writes the vector to the destination SIMD&FP register.
func VnegS8(v0 Int8X8) Int8X8 {
	return C.vneg_s8(v0)
}

// Floating-point Negate (vector). This instruction negates the value of each vector element in the source SIMD&FP register, writes the result to a vector, and writes the vector to the destination SIMD&FP register.
func VnegF32(v0 Float32X2) Float32X2 {
	return C.vneg_f32(v0)
}

// Negate (vector). This instruction reads each vector element from the source SIMD&FP register, negates each value, puts the result into a vector, and writes the vector to the destination SIMD&FP register.
func VnegS32(v0 Int32X2) Int32X2 {
	return C.vneg_s32(v0)
}

// Negate (vector). This instruction reads each vector element from the source SIMD&FP register, negates each value, puts the result into a vector, and writes the vector to the destination SIMD&FP register.
func VnegS16(v0 Int16X4) Int16X4 {
	return C.vneg_s16(v0)
}

// Bitwise inclusive OR NOT (vector). This instruction performs a bitwise OR NOT between the two source SIMD&FP registers, and writes the result to the destination SIMD&FP register.
func VornqU8(v0 Uint8X16, v1 Uint8X16) Uint8X16 {
	return C.vornq_u8(v0, v1)
}

// Bitwise inclusive OR NOT (vector). This instruction performs a bitwise OR NOT between the two source SIMD&FP registers, and writes the result to the destination SIMD&FP register.
func VornqU32(v0 Uint32X4, v1 Uint32X4) Uint32X4 {
	return C.vornq_u32(v0, v1)
}

// Bitwise inclusive OR NOT (vector). This instruction performs a bitwise OR NOT between the two source SIMD&FP registers, and writes the result to the destination SIMD&FP register.
func VornqU64(v0 Uint64X2, v1 Uint64X2) Uint64X2 {
	return C.vornq_u64(v0, v1)
}

// Bitwise inclusive OR NOT (vector). This instruction performs a bitwise OR NOT between the two source SIMD&FP registers, and writes the result to the destination SIMD&FP register.
func VornqU16(v0 Uint16X8, v1 Uint16X8) Uint16X8 {
	return C.vornq_u16(v0, v1)
}

// Bitwise inclusive OR NOT (vector). This instruction performs a bitwise OR NOT between the two source SIMD&FP registers, and writes the result to the destination SIMD&FP register.
func VornqS8(v0 Int8X16, v1 Int8X16) Int8X16 {
	return C.vornq_s8(v0, v1)
}

// Bitwise inclusive OR NOT (vector). This instruction performs a bitwise OR NOT between the two source SIMD&FP registers, and writes the result to the destination SIMD&FP register.
func VornqS32(v0 Int32X4, v1 Int32X4) Int32X4 {
	return C.vornq_s32(v0, v1)
}

// Bitwise inclusive OR NOT (vector). This instruction performs a bitwise OR NOT between the two source SIMD&FP registers, and writes the result to the destination SIMD&FP register.
func VornqS64(v0 Int64X2, v1 Int64X2) Int64X2 {
	return C.vornq_s64(v0, v1)
}

// Bitwise inclusive OR NOT (vector). This instruction performs a bitwise OR NOT between the two source SIMD&FP registers, and writes the result to the destination SIMD&FP register.
func VornqS16(v0 Int16X8, v1 Int16X8) Int16X8 {
	return C.vornq_s16(v0, v1)
}

// Bitwise inclusive OR NOT (vector). This instruction performs a bitwise OR NOT between the two source SIMD&FP registers, and writes the result to the destination SIMD&FP register.
func VornU8(v0 Uint8X8, v1 Uint8X8) Uint8X8 {
	return C.vorn_u8(v0, v1)
}

// Bitwise inclusive OR NOT (vector). This instruction performs a bitwise OR NOT between the two source SIMD&FP registers, and writes the result to the destination SIMD&FP register.
func VornU32(v0 Uint32X2, v1 Uint32X2) Uint32X2 {
	return C.vorn_u32(v0, v1)
}

// Bitwise inclusive OR NOT (vector). This instruction performs a bitwise OR NOT between the two source SIMD&FP registers, and writes the result to the destination SIMD&FP register.
func VornU64(v0 Uint64X1, v1 Uint64X1) Uint64X1 {
	return C.vorn_u64(v0, v1)
}

// Bitwise inclusive OR NOT (vector). This instruction performs a bitwise OR NOT between the two source SIMD&FP registers, and writes the result to the destination SIMD&FP register.
func VornU16(v0 Uint16X4, v1 Uint16X4) Uint16X4 {
	return C.vorn_u16(v0, v1)
}

// Bitwise inclusive OR NOT (vector). This instruction performs a bitwise OR NOT between the two source SIMD&FP registers, and writes the result to the destination SIMD&FP register.
func VornS8(v0 Int8X8, v1 Int8X8) Int8X8 {
	return C.vorn_s8(v0, v1)
}

// Bitwise inclusive OR NOT (vector). This instruction performs a bitwise OR NOT between the two source SIMD&FP registers, and writes the result to the destination SIMD&FP register.
func VornS32(v0 Int32X2, v1 Int32X2) Int32X2 {
	return C.vorn_s32(v0, v1)
}

// Bitwise inclusive OR NOT (vector). This instruction performs a bitwise OR NOT between the two source SIMD&FP registers, and writes the result to the destination SIMD&FP register.
func VornS64(v0 Int64X1, v1 Int64X1) Int64X1 {
	return C.vorn_s64(v0, v1)
}

// Bitwise inclusive OR NOT (vector). This instruction performs a bitwise OR NOT between the two source SIMD&FP registers, and writes the result to the destination SIMD&FP register.
func VornS16(v0 Int16X4, v1 Int16X4) Int16X4 {
	return C.vorn_s16(v0, v1)
}

// Bitwise inclusive OR (vector, register). This instruction performs a bitwise OR between the two source SIMD&FP registers, and writes the result to the destination SIMD&FP register.
func VorrqU8(v0 Uint8X16, v1 Uint8X16) Uint8X16 {
	return C.vorrq_u8(v0, v1)
}

// Bitwise inclusive OR (vector, register). This instruction performs a bitwise OR between the two source SIMD&FP registers, and writes the result to the destination SIMD&FP register.
func VorrqU32(v0 Uint32X4, v1 Uint32X4) Uint32X4 {
	return C.vorrq_u32(v0, v1)
}

// Bitwise inclusive OR (vector, register). This instruction performs a bitwise OR between the two source SIMD&FP registers, and writes the result to the destination SIMD&FP register.
func VorrqU64(v0 Uint64X2, v1 Uint64X2) Uint64X2 {
	return C.vorrq_u64(v0, v1)
}

// Bitwise inclusive OR (vector, register). This instruction performs a bitwise OR between the two source SIMD&FP registers, and writes the result to the destination SIMD&FP register.
func VorrqU16(v0 Uint16X8, v1 Uint16X8) Uint16X8 {
	return C.vorrq_u16(v0, v1)
}

// Bitwise inclusive OR (vector, register). This instruction performs a bitwise OR between the two source SIMD&FP registers, and writes the result to the destination SIMD&FP register.
func VorrqS8(v0 Int8X16, v1 Int8X16) Int8X16 {
	return C.vorrq_s8(v0, v1)
}

// Bitwise inclusive OR (vector, register). This instruction performs a bitwise OR between the two source SIMD&FP registers, and writes the result to the destination SIMD&FP register.
func VorrqS32(v0 Int32X4, v1 Int32X4) Int32X4 {
	return C.vorrq_s32(v0, v1)
}

// Bitwise inclusive OR (vector, register). This instruction performs a bitwise OR between the two source SIMD&FP registers, and writes the result to the destination SIMD&FP register.
func VorrqS64(v0 Int64X2, v1 Int64X2) Int64X2 {
	return C.vorrq_s64(v0, v1)
}

// Bitwise inclusive OR (vector, register). This instruction performs a bitwise OR between the two source SIMD&FP registers, and writes the result to the destination SIMD&FP register.
func VorrqS16(v0 Int16X8, v1 Int16X8) Int16X8 {
	return C.vorrq_s16(v0, v1)
}

// Bitwise inclusive OR (vector, register). This instruction performs a bitwise OR between the two source SIMD&FP registers, and writes the result to the destination SIMD&FP register.
func VorrU8(v0 Uint8X8, v1 Uint8X8) Uint8X8 {
	return C.vorr_u8(v0, v1)
}

// Bitwise inclusive OR (vector, register). This instruction performs a bitwise OR between the two source SIMD&FP registers, and writes the result to the destination SIMD&FP register.
func VorrU32(v0 Uint32X2, v1 Uint32X2) Uint32X2 {
	return C.vorr_u32(v0, v1)
}

// Bitwise inclusive OR (vector, register). This instruction performs a bitwise OR between the two source SIMD&FP registers, and writes the result to the destination SIMD&FP register.
func VorrU64(v0 Uint64X1, v1 Uint64X1) Uint64X1 {
	return C.vorr_u64(v0, v1)
}

// Bitwise inclusive OR (vector, register). This instruction performs a bitwise OR between the two source SIMD&FP registers, and writes the result to the destination SIMD&FP register.
func VorrU16(v0 Uint16X4, v1 Uint16X4) Uint16X4 {
	return C.vorr_u16(v0, v1)
}

// Bitwise inclusive OR (vector, register). This instruction performs a bitwise OR between the two source SIMD&FP registers, and writes the result to the destination SIMD&FP register.
func VorrS8(v0 Int8X8, v1 Int8X8) Int8X8 {
	return C.vorr_s8(v0, v1)
}

// Bitwise inclusive OR (vector, register). This instruction performs a bitwise OR between the two source SIMD&FP registers, and writes the result to the destination SIMD&FP register.
func VorrS32(v0 Int32X2, v1 Int32X2) Int32X2 {
	return C.vorr_s32(v0, v1)
}

// Bitwise inclusive OR (vector, register). This instruction performs a bitwise OR between the two source SIMD&FP registers, and writes the result to the destination SIMD&FP register.
func VorrS64(v0 Int64X1, v1 Int64X1) Int64X1 {
	return C.vorr_s64(v0, v1)
}

// Bitwise inclusive OR (vector, register). This instruction performs a bitwise OR between the two source SIMD&FP registers, and writes the result to the destination SIMD&FP register.
func VorrS16(v0 Int16X4, v1 Int16X4) Int16X4 {
	return C.vorr_s16(v0, v1)
}

// Unsigned Add and Accumulate Long Pairwise. This instruction adds pairs of adjacent unsigned integer values from the vector in the source SIMD&FP register and accumulates the results with the vector elements of the destination SIMD&FP register. The destination vector elements are twice as long as the source vector elements.
func VpadalqU8(v0 Uint16X8, v1 Uint8X16) Uint16X8 {
	return C.vpadalq_u8(v0, v1)
}

// Unsigned Add and Accumulate Long Pairwise. This instruction adds pairs of adjacent unsigned integer values from the vector in the source SIMD&FP register and accumulates the results with the vector elements of the destination SIMD&FP register. The destination vector elements are twice as long as the source vector elements.
func VpadalqU32(v0 Uint64X2, v1 Uint32X4) Uint64X2 {
	return C.vpadalq_u32(v0, v1)
}

// Unsigned Add and Accumulate Long Pairwise. This instruction adds pairs of adjacent unsigned integer values from the vector in the source SIMD&FP register and accumulates the results with the vector elements of the destination SIMD&FP register. The destination vector elements are twice as long as the source vector elements.
func VpadalqU16(v0 Uint32X4, v1 Uint16X8) Uint32X4 {
	return C.vpadalq_u16(v0, v1)
}

// Signed Add and Accumulate Long Pairwise. This instruction adds pairs of adjacent signed integer values from the vector in the source SIMD&FP register and accumulates the results into the vector elements of the destination SIMD&FP register. The destination vector elements are twice as long as the source vector elements.
func VpadalqS8(v0 Int16X8, v1 Int8X16) Int16X8 {
	return C.vpadalq_s8(v0, v1)
}

// Signed Add and Accumulate Long Pairwise. This instruction adds pairs of adjacent signed integer values from the vector in the source SIMD&FP register and accumulates the results into the vector elements of the destination SIMD&FP register. The destination vector elements are twice as long as the source vector elements.
func VpadalqS32(v0 Int64X2, v1 Int32X4) Int64X2 {
	return C.vpadalq_s32(v0, v1)
}

// Signed Add and Accumulate Long Pairwise. This instruction adds pairs of adjacent signed integer values from the vector in the source SIMD&FP register and accumulates the results into the vector elements of the destination SIMD&FP register. The destination vector elements are twice as long as the source vector elements.
func VpadalqS16(v0 Int32X4, v1 Int16X8) Int32X4 {
	return C.vpadalq_s16(v0, v1)
}

// Unsigned Add and Accumulate Long Pairwise. This instruction adds pairs of adjacent unsigned integer values from the vector in the source SIMD&FP register and accumulates the results with the vector elements of the destination SIMD&FP register. The destination vector elements are twice as long as the source vector elements.
func VpadalU8(v0 Uint16X4, v1 Uint8X8) Uint16X4 {
	return C.vpadal_u8(v0, v1)
}

// Unsigned Add and Accumulate Long Pairwise. This instruction adds pairs of adjacent unsigned integer values from the vector in the source SIMD&FP register and accumulates the results with the vector elements of the destination SIMD&FP register. The destination vector elements are twice as long as the source vector elements.
func VpadalU32(v0 Uint64X1, v1 Uint32X2) Uint64X1 {
	return C.vpadal_u32(v0, v1)
}

// Unsigned Add and Accumulate Long Pairwise. This instruction adds pairs of adjacent unsigned integer values from the vector in the source SIMD&FP register and accumulates the results with the vector elements of the destination SIMD&FP register. The destination vector elements are twice as long as the source vector elements.
func VpadalU16(v0 Uint32X2, v1 Uint16X4) Uint32X2 {
	return C.vpadal_u16(v0, v1)
}

// Signed Add and Accumulate Long Pairwise. This instruction adds pairs of adjacent signed integer values from the vector in the source SIMD&FP register and accumulates the results into the vector elements of the destination SIMD&FP register. The destination vector elements are twice as long as the source vector elements.
func VpadalS8(v0 Int16X4, v1 Int8X8) Int16X4 {
	return C.vpadal_s8(v0, v1)
}

// Signed Add and Accumulate Long Pairwise. This instruction adds pairs of adjacent signed integer values from the vector in the source SIMD&FP register and accumulates the results into the vector elements of the destination SIMD&FP register. The destination vector elements are twice as long as the source vector elements.
func VpadalS32(v0 Int64X1, v1 Int32X2) Int64X1 {
	return C.vpadal_s32(v0, v1)
}

// Signed Add and Accumulate Long Pairwise. This instruction adds pairs of adjacent signed integer values from the vector in the source SIMD&FP register and accumulates the results into the vector elements of the destination SIMD&FP register. The destination vector elements are twice as long as the source vector elements.
func VpadalS16(v0 Int32X2, v1 Int16X4) Int32X2 {
	return C.vpadal_s16(v0, v1)
}

// Add Pairwise (vector). This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements from the concatenated vector, adds each pair of values together, places the result into a vector, and writes the vector to the destination SIMD&FP register.
func VpaddU8(v0 Uint8X8, v1 Uint8X8) Uint8X8 {
	return C.vpadd_u8(v0, v1)
}

// Add Pairwise (vector). This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements from the concatenated vector, adds each pair of values together, places the result into a vector, and writes the vector to the destination SIMD&FP register.
func VpaddU32(v0 Uint32X2, v1 Uint32X2) Uint32X2 {
	return C.vpadd_u32(v0, v1)
}

// Add Pairwise (vector). This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements from the concatenated vector, adds each pair of values together, places the result into a vector, and writes the vector to the destination SIMD&FP register.
func VpaddU16(v0 Uint16X4, v1 Uint16X4) Uint16X4 {
	return C.vpadd_u16(v0, v1)
}

// Add Pairwise (vector). This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements from the concatenated vector, adds each pair of values together, places the result into a vector, and writes the vector to the destination SIMD&FP register.
func VpaddS8(v0 Int8X8, v1 Int8X8) Int8X8 {
	return C.vpadd_s8(v0, v1)
}

// Floating-point Add Pairwise (vector). This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements from the concatenated vector, adds each pair of values together, places the result into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are floating-point values.
func VpaddF32(v0 Float32X2, v1 Float32X2) Float32X2 {
	return C.vpadd_f32(v0, v1)
}

// Add Pairwise (vector). This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements from the concatenated vector, adds each pair of values together, places the result into a vector, and writes the vector to the destination SIMD&FP register.
func VpaddS32(v0 Int32X2, v1 Int32X2) Int32X2 {
	return C.vpadd_s32(v0, v1)
}

// Add Pairwise (vector). This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements from the concatenated vector, adds each pair of values together, places the result into a vector, and writes the vector to the destination SIMD&FP register.
func VpaddS16(v0 Int16X4, v1 Int16X4) Int16X4 {
	return C.vpadd_s16(v0, v1)
}

// Unsigned Add Long Pairwise. This instruction adds pairs of adjacent unsigned integer values from the vector in the source SIMD&FP register, places the result into a vector, and writes the vector to the destination SIMD&FP register. The destination vector elements are twice as long as the source vector elements.
func VpaddlqU8(v0 Uint8X16) Uint16X8 {
	return C.vpaddlq_u8(v0)
}

// Unsigned Add Long Pairwise. This instruction adds pairs of adjacent unsigned integer values from the vector in the source SIMD&FP register, places the result into a vector, and writes the vector to the destination SIMD&FP register. The destination vector elements are twice as long as the source vector elements.
func VpaddlqU32(v0 Uint32X4) Uint64X2 {
	return C.vpaddlq_u32(v0)
}

// Unsigned Add Long Pairwise. This instruction adds pairs of adjacent unsigned integer values from the vector in the source SIMD&FP register, places the result into a vector, and writes the vector to the destination SIMD&FP register. The destination vector elements are twice as long as the source vector elements.
func VpaddlqU16(v0 Uint16X8) Uint32X4 {
	return C.vpaddlq_u16(v0)
}

// Signed Add Long Pairwise. This instruction adds pairs of adjacent signed integer values from the vector in the source SIMD&FP register, places the result into a vector, and writes the vector to the destination SIMD&FP register. The destination vector elements are twice as long as the source vector elements.
func VpaddlqS8(v0 Int8X16) Int16X8 {
	return C.vpaddlq_s8(v0)
}

// Signed Add Long Pairwise. This instruction adds pairs of adjacent signed integer values from the vector in the source SIMD&FP register, places the result into a vector, and writes the vector to the destination SIMD&FP register. The destination vector elements are twice as long as the source vector elements.
func VpaddlqS32(v0 Int32X4) Int64X2 {
	return C.vpaddlq_s32(v0)
}

// Signed Add Long Pairwise. This instruction adds pairs of adjacent signed integer values from the vector in the source SIMD&FP register, places the result into a vector, and writes the vector to the destination SIMD&FP register. The destination vector elements are twice as long as the source vector elements.
func VpaddlqS16(v0 Int16X8) Int32X4 {
	return C.vpaddlq_s16(v0)
}

// Unsigned Add Long Pairwise. This instruction adds pairs of adjacent unsigned integer values from the vector in the source SIMD&FP register, places the result into a vector, and writes the vector to the destination SIMD&FP register. The destination vector elements are twice as long as the source vector elements.
func VpaddlU8(v0 Uint8X8) Uint16X4 {
	return C.vpaddl_u8(v0)
}

// Unsigned Add Long Pairwise. This instruction adds pairs of adjacent unsigned integer values from the vector in the source SIMD&FP register, places the result into a vector, and writes the vector to the destination SIMD&FP register. The destination vector elements are twice as long as the source vector elements.
func VpaddlU32(v0 Uint32X2) Uint64X1 {
	return C.vpaddl_u32(v0)
}

// Unsigned Add Long Pairwise. This instruction adds pairs of adjacent unsigned integer values from the vector in the source SIMD&FP register, places the result into a vector, and writes the vector to the destination SIMD&FP register. The destination vector elements are twice as long as the source vector elements.
func VpaddlU16(v0 Uint16X4) Uint32X2 {
	return C.vpaddl_u16(v0)
}

// Signed Add Long Pairwise. This instruction adds pairs of adjacent signed integer values from the vector in the source SIMD&FP register, places the result into a vector, and writes the vector to the destination SIMD&FP register. The destination vector elements are twice as long as the source vector elements.
func VpaddlS8(v0 Int8X8) Int16X4 {
	return C.vpaddl_s8(v0)
}

// Signed Add Long Pairwise. This instruction adds pairs of adjacent signed integer values from the vector in the source SIMD&FP register, places the result into a vector, and writes the vector to the destination SIMD&FP register. The destination vector elements are twice as long as the source vector elements.
func VpaddlS32(v0 Int32X2) Int64X1 {
	return C.vpaddl_s32(v0)
}

// Signed Add Long Pairwise. This instruction adds pairs of adjacent signed integer values from the vector in the source SIMD&FP register, places the result into a vector, and writes the vector to the destination SIMD&FP register. The destination vector elements are twice as long as the source vector elements.
func VpaddlS16(v0 Int16X4) Int32X2 {
	return C.vpaddl_s16(v0)
}

// Unsigned Maximum Pairwise. This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements in the two source SIMD&FP registers, writes the largest of each pair of unsigned integer values into a vector, and writes the vector to the destination SIMD&FP register.
func VpmaxU8(v0 Uint8X8, v1 Uint8X8) Uint8X8 {
	return C.vpmax_u8(v0, v1)
}

// Unsigned Maximum Pairwise. This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements in the two source SIMD&FP registers, writes the largest of each pair of unsigned integer values into a vector, and writes the vector to the destination SIMD&FP register.
func VpmaxU32(v0 Uint32X2, v1 Uint32X2) Uint32X2 {
	return C.vpmax_u32(v0, v1)
}

// Unsigned Maximum Pairwise. This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements in the two source SIMD&FP registers, writes the largest of each pair of unsigned integer values into a vector, and writes the vector to the destination SIMD&FP register.
func VpmaxU16(v0 Uint16X4, v1 Uint16X4) Uint16X4 {
	return C.vpmax_u16(v0, v1)
}

// Signed Maximum Pairwise. This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements in the two source SIMD&FP registers, writes the largest of each pair of signed integer values into a vector, and writes the vector to the destination SIMD&FP register.
func VpmaxS8(v0 Int8X8, v1 Int8X8) Int8X8 {
	return C.vpmax_s8(v0, v1)
}

// Floating-point Maximum Pairwise (vector). This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements from the concatenated vector, writes the larger of each pair of values into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are floating-point values.
func VpmaxF32(v0 Float32X2, v1 Float32X2) Float32X2 {
	return C.vpmax_f32(v0, v1)
}

// Signed Maximum Pairwise. This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements in the two source SIMD&FP registers, writes the largest of each pair of signed integer values into a vector, and writes the vector to the destination SIMD&FP register.
func VpmaxS32(v0 Int32X2, v1 Int32X2) Int32X2 {
	return C.vpmax_s32(v0, v1)
}

// Signed Maximum Pairwise. This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements in the two source SIMD&FP registers, writes the largest of each pair of signed integer values into a vector, and writes the vector to the destination SIMD&FP register.
func VpmaxS16(v0 Int16X4, v1 Int16X4) Int16X4 {
	return C.vpmax_s16(v0, v1)
}

// Unsigned Minimum Pairwise. This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements in the two source SIMD&FP registers, writes the smallest of each pair of unsigned integer values into a vector, and writes the vector to the destination SIMD&FP register.
func VpminU8(v0 Uint8X8, v1 Uint8X8) Uint8X8 {
	return C.vpmin_u8(v0, v1)
}

// Unsigned Minimum Pairwise. This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements in the two source SIMD&FP registers, writes the smallest of each pair of unsigned integer values into a vector, and writes the vector to the destination SIMD&FP register.
func VpminU32(v0 Uint32X2, v1 Uint32X2) Uint32X2 {
	return C.vpmin_u32(v0, v1)
}

// Unsigned Minimum Pairwise. This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements in the two source SIMD&FP registers, writes the smallest of each pair of unsigned integer values into a vector, and writes the vector to the destination SIMD&FP register.
func VpminU16(v0 Uint16X4, v1 Uint16X4) Uint16X4 {
	return C.vpmin_u16(v0, v1)
}

// Signed Minimum Pairwise. This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements in the two source SIMD&FP registers, writes the smallest of each pair of signed integer values into a vector, and writes the vector to the destination SIMD&FP register.
func VpminS8(v0 Int8X8, v1 Int8X8) Int8X8 {
	return C.vpmin_s8(v0, v1)
}

// Floating-point Minimum Pairwise (vector). This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements from the concatenated vector, writes the smaller of each pair of values into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are floating-point values.
func VpminF32(v0 Float32X2, v1 Float32X2) Float32X2 {
	return C.vpmin_f32(v0, v1)
}

// Signed Minimum Pairwise. This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements in the two source SIMD&FP registers, writes the smallest of each pair of signed integer values into a vector, and writes the vector to the destination SIMD&FP register.
func VpminS32(v0 Int32X2, v1 Int32X2) Int32X2 {
	return C.vpmin_s32(v0, v1)
}

// Signed Minimum Pairwise. This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements in the two source SIMD&FP registers, writes the smallest of each pair of signed integer values into a vector, and writes the vector to the destination SIMD&FP register.
func VpminS16(v0 Int16X4, v1 Int16X4) Int16X4 {
	return C.vpmin_s16(v0, v1)
}

// Signed saturating Absolute value. This instruction reads each vector element from the source SIMD&FP register, puts the absolute value of the result into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are signed integer values.
func VqabsqS8(v0 Int8X16) Int8X16 {
	return C.vqabsq_s8(v0)
}

// Signed saturating Absolute value. This instruction reads each vector element from the source SIMD&FP register, puts the absolute value of the result into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are signed integer values.
func VqabsqS32(v0 Int32X4) Int32X4 {
	return C.vqabsq_s32(v0)
}

// Signed saturating Absolute value. This instruction reads each vector element from the source SIMD&FP register, puts the absolute value of the result into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are signed integer values.
func VqabsqS16(v0 Int16X8) Int16X8 {
	return C.vqabsq_s16(v0)
}

// Signed saturating Absolute value. This instruction reads each vector element from the source SIMD&FP register, puts the absolute value of the result into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are signed integer values.
func VqabsS8(v0 Int8X8) Int8X8 {
	return C.vqabs_s8(v0)
}

// Signed saturating Absolute value. This instruction reads each vector element from the source SIMD&FP register, puts the absolute value of the result into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are signed integer values.
func VqabsS32(v0 Int32X2) Int32X2 {
	return C.vqabs_s32(v0)
}

// Signed saturating Absolute value. This instruction reads each vector element from the source SIMD&FP register, puts the absolute value of the result into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are signed integer values.
func VqabsS16(v0 Int16X4) Int16X4 {
	return C.vqabs_s16(v0)
}

// Unsigned saturating Add. This instruction adds the values of corresponding elements of the two source SIMD&FP registers, places the results into a vector, and writes the vector to the destination SIMD&FP register.
func VqaddqU8(v0 Uint8X16, v1 Uint8X16) Uint8X16 {
	return C.vqaddq_u8(v0, v1)
}

// Unsigned saturating Add. This instruction adds the values of corresponding elements of the two source SIMD&FP registers, places the results into a vector, and writes the vector to the destination SIMD&FP register.
func VqaddqU32(v0 Uint32X4, v1 Uint32X4) Uint32X4 {
	return C.vqaddq_u32(v0, v1)
}

// Unsigned saturating Add. This instruction adds the values of corresponding elements of the two source SIMD&FP registers, places the results into a vector, and writes the vector to the destination SIMD&FP register.
func VqaddqU64(v0 Uint64X2, v1 Uint64X2) Uint64X2 {
	return C.vqaddq_u64(v0, v1)
}

// Unsigned saturating Add. This instruction adds the values of corresponding elements of the two source SIMD&FP registers, places the results into a vector, and writes the vector to the destination SIMD&FP register.
func VqaddqU16(v0 Uint16X8, v1 Uint16X8) Uint16X8 {
	return C.vqaddq_u16(v0, v1)
}

// Signed saturating Add. This instruction adds the values of corresponding elements of the two source SIMD&FP registers, places the results into a vector, and writes the vector to the destination SIMD&FP register.
func VqaddqS8(v0 Int8X16, v1 Int8X16) Int8X16 {
	return C.vqaddq_s8(v0, v1)
}

// Signed saturating Add. This instruction adds the values of corresponding elements of the two source SIMD&FP registers, places the results into a vector, and writes the vector to the destination SIMD&FP register.
func VqaddqS32(v0 Int32X4, v1 Int32X4) Int32X4 {
	return C.vqaddq_s32(v0, v1)
}

// Signed saturating Add. This instruction adds the values of corresponding elements of the two source SIMD&FP registers, places the results into a vector, and writes the vector to the destination SIMD&FP register.
func VqaddqS64(v0 Int64X2, v1 Int64X2) Int64X2 {
	return C.vqaddq_s64(v0, v1)
}

// Signed saturating Add. This instruction adds the values of corresponding elements of the two source SIMD&FP registers, places the results into a vector, and writes the vector to the destination SIMD&FP register.
func VqaddqS16(v0 Int16X8, v1 Int16X8) Int16X8 {
	return C.vqaddq_s16(v0, v1)
}

// Unsigned saturating Add. This instruction adds the values of corresponding elements of the two source SIMD&FP registers, places the results into a vector, and writes the vector to the destination SIMD&FP register.
func VqaddU8(v0 Uint8X8, v1 Uint8X8) Uint8X8 {
	return C.vqadd_u8(v0, v1)
}

// Unsigned saturating Add. This instruction adds the values of corresponding elements of the two source SIMD&FP registers, places the results into a vector, and writes the vector to the destination SIMD&FP register.
func VqaddU32(v0 Uint32X2, v1 Uint32X2) Uint32X2 {
	return C.vqadd_u32(v0, v1)
}

// Unsigned saturating Add. This instruction adds the values of corresponding elements of the two source SIMD&FP registers, places the results into a vector, and writes the vector to the destination SIMD&FP register.
func VqaddU64(v0 Uint64X1, v1 Uint64X1) Uint64X1 {
	return C.vqadd_u64(v0, v1)
}

// Unsigned saturating Add. This instruction adds the values of corresponding elements of the two source SIMD&FP registers, places the results into a vector, and writes the vector to the destination SIMD&FP register.
func VqaddU16(v0 Uint16X4, v1 Uint16X4) Uint16X4 {
	return C.vqadd_u16(v0, v1)
}

// Signed saturating Add. This instruction adds the values of corresponding elements of the two source SIMD&FP registers, places the results into a vector, and writes the vector to the destination SIMD&FP register.
func VqaddS8(v0 Int8X8, v1 Int8X8) Int8X8 {
	return C.vqadd_s8(v0, v1)
}

// Signed saturating Add. This instruction adds the values of corresponding elements of the two source SIMD&FP registers, places the results into a vector, and writes the vector to the destination SIMD&FP register.
func VqaddS32(v0 Int32X2, v1 Int32X2) Int32X2 {
	return C.vqadd_s32(v0, v1)
}

// Signed saturating Add. This instruction adds the values of corresponding elements of the two source SIMD&FP registers, places the results into a vector, and writes the vector to the destination SIMD&FP register.
func VqaddS64(v0 Int64X1, v1 Int64X1) Int64X1 {
	return C.vqadd_s64(v0, v1)
}

// Signed saturating Add. This instruction adds the values of corresponding elements of the two source SIMD&FP registers, places the results into a vector, and writes the vector to the destination SIMD&FP register.
func VqaddS16(v0 Int16X4, v1 Int16X4) Int16X4 {
	return C.vqadd_s16(v0, v1)
}

// Signed saturating Doubling Multiply-Add Long. This instruction multiplies corresponding signed integer values in the lower or upper half of the vectors of the two source SIMD&FP registers, doubles the results, and accumulates the final results with the vector elements of the destination SIMD&FP register. The destination vector elements are twice as long as the elements that are multiplied.
func VqdmlalS32(v0 Int64X2, v1 Int32X2, v2 Int32X2) Int64X2 {
	return C.vqdmlal_s32(v0, v1, v2)
}

// Signed saturating Doubling Multiply-Add Long. This instruction multiplies corresponding signed integer values in the lower or upper half of the vectors of the two source SIMD&FP registers, doubles the results, and accumulates the final results with the vector elements of the destination SIMD&FP register. The destination vector elements are twice as long as the elements that are multiplied.
func VqdmlalS16(v0 Int32X4, v1 Int16X4, v2 Int16X4) Int32X4 {
	return C.vqdmlal_s16(v0, v1, v2)
}

// Vector widening saturating doubling multiply accumulate with scalar
func VqdmlalNS32(v0 Int64X2, v1 Int32X2, v2 Int32) Int64X2 {
	return C.vqdmlal_n_s32(v0, v1, v2)
}

// Vector widening saturating doubling multiply accumulate with scalar
func VqdmlalNS16(v0 Int32X4, v1 Int16X4, v2 Int16) Int32X4 {
	return C.vqdmlal_n_s16(v0, v1, v2)
}

// Signed saturating Doubling Multiply-Subtract Long. This instruction multiplies corresponding signed integer values in the lower or upper half of the vectors of the two source SIMD&FP registers, doubles the results, and subtracts the final results from the vector elements of the destination SIMD&FP register. The destination vector elements are twice as long as the elements that are multiplied.
func VqdmlslS32(v0 Int64X2, v1 Int32X2, v2 Int32X2) Int64X2 {
	return C.vqdmlsl_s32(v0, v1, v2)
}

// Signed saturating Doubling Multiply-Subtract Long. This instruction multiplies corresponding signed integer values in the lower or upper half of the vectors of the two source SIMD&FP registers, doubles the results, and subtracts the final results from the vector elements of the destination SIMD&FP register. The destination vector elements are twice as long as the elements that are multiplied.
func VqdmlslS16(v0 Int32X4, v1 Int16X4, v2 Int16X4) Int32X4 {
	return C.vqdmlsl_s16(v0, v1, v2)
}

// Vector widening saturating doubling multiply subtract with scalar
func VqdmlslNS32(v0 Int64X2, v1 Int32X2, v2 Int32) Int64X2 {
	return C.vqdmlsl_n_s32(v0, v1, v2)
}

// Vector widening saturating doubling multiply subtract with scalar
func VqdmlslNS16(v0 Int32X4, v1 Int16X4, v2 Int16) Int32X4 {
	return C.vqdmlsl_n_s16(v0, v1, v2)
}

// Signed saturating Doubling Multiply returning High half. This instruction multiplies the values of corresponding elements of the two source SIMD&FP registers, doubles the results, places the most significant half of the final results into a vector, and writes the vector to the destination SIMD&FP register.
func VqdmulhqS32(v0 Int32X4, v1 Int32X4) Int32X4 {
	return C.vqdmulhq_s32(v0, v1)
}

// Signed saturating Doubling Multiply returning High half. This instruction multiplies the values of corresponding elements of the two source SIMD&FP registers, doubles the results, places the most significant half of the final results into a vector, and writes the vector to the destination SIMD&FP register.
func VqdmulhqS16(v0 Int16X8, v1 Int16X8) Int16X8 {
	return C.vqdmulhq_s16(v0, v1)
}

// Signed saturating Doubling Multiply returning High half. This instruction multiplies the values of corresponding elements of the two source SIMD&FP registers, doubles the results, places the most significant half of the final results into a vector, and writes the vector to the destination SIMD&FP register.
func VqdmulhS32(v0 Int32X2, v1 Int32X2) Int32X2 {
	return C.vqdmulh_s32(v0, v1)
}

// Signed saturating Doubling Multiply returning High half. This instruction multiplies the values of corresponding elements of the two source SIMD&FP registers, doubles the results, places the most significant half of the final results into a vector, and writes the vector to the destination SIMD&FP register.
func VqdmulhS16(v0 Int16X4, v1 Int16X4) Int16X4 {
	return C.vqdmulh_s16(v0, v1)
}

// Vector saturating doubling multiply high with scalar
func VqdmulhqNS32(v0 Int32X4, v1 Int32) Int32X4 {
	return C.vqdmulhq_n_s32(v0, v1)
}

// Vector saturating doubling multiply high with scalar
func VqdmulhqNS16(v0 Int16X8, v1 Int16) Int16X8 {
	return C.vqdmulhq_n_s16(v0, v1)
}

// Vector saturating doubling multiply high with scalar
func VqdmulhNS32(v0 Int32X2, v1 Int32) Int32X2 {
	return C.vqdmulh_n_s32(v0, v1)
}

// Vector saturating doubling multiply high with scalar
func VqdmulhNS16(v0 Int16X4, v1 Int16) Int16X4 {
	return C.vqdmulh_n_s16(v0, v1)
}

// Signed saturating Doubling Multiply Long. This instruction multiplies corresponding vector elements in the lower or upper half of the two source SIMD&FP registers, doubles the results, places the final results in a vector, and writes the vector to the destination SIMD&FP register.
func VqdmullS32(v0 Int32X2, v1 Int32X2) Int64X2 {
	return C.vqdmull_s32(v0, v1)
}

// Signed saturating Doubling Multiply Long. This instruction multiplies corresponding vector elements in the lower or upper half of the two source SIMD&FP registers, doubles the results, places the final results in a vector, and writes the vector to the destination SIMD&FP register.
func VqdmullS16(v0 Int16X4, v1 Int16X4) Int32X4 {
	return C.vqdmull_s16(v0, v1)
}

// Vector saturating doubling long multiply with scalar
func VqdmullNS32(v0 Int32X2, v1 Int32) Int64X2 {
	return C.vqdmull_n_s32(v0, v1)
}

// Vector saturating doubling long multiply with scalar
func VqdmullNS16(v0 Int16X4, v1 Int16) Int32X4 {
	return C.vqdmull_n_s16(v0, v1)
}

// Unsigned saturating extract Narrow. This instruction reads each vector element from the source SIMD&FP register, saturates each value to half the original width, places the result into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are unsigned integer values.
func VqmovnU32(v0 Uint32X4) Uint16X4 {
	return C.vqmovn_u32(v0)
}

// Unsigned saturating extract Narrow. This instruction reads each vector element from the source SIMD&FP register, saturates each value to half the original width, places the result into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are unsigned integer values.
func VqmovnU64(v0 Uint64X2) Uint32X2 {
	return C.vqmovn_u64(v0)
}

// Unsigned saturating extract Narrow. This instruction reads each vector element from the source SIMD&FP register, saturates each value to half the original width, places the result into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are unsigned integer values.
func VqmovnU16(v0 Uint16X8) Uint8X8 {
	return C.vqmovn_u16(v0)
}

// Signed saturating extract Narrow. This instruction reads each vector element from the source SIMD&FP register, saturates the value to half the original width, places the result into a vector, and writes the vector to the lower or upper half of the destination SIMD&FP register. The destination vector elements are half as long as the source vector elements. All the values in this instruction are signed integer values.
func VqmovnS32(v0 Int32X4) Int16X4 {
	return C.vqmovn_s32(v0)
}

// Signed saturating extract Narrow. This instruction reads each vector element from the source SIMD&FP register, saturates the value to half the original width, places the result into a vector, and writes the vector to the lower or upper half of the destination SIMD&FP register. The destination vector elements are half as long as the source vector elements. All the values in this instruction are signed integer values.
func VqmovnS64(v0 Int64X2) Int32X2 {
	return C.vqmovn_s64(v0)
}

// Signed saturating extract Narrow. This instruction reads each vector element from the source SIMD&FP register, saturates the value to half the original width, places the result into a vector, and writes the vector to the lower or upper half of the destination SIMD&FP register. The destination vector elements are half as long as the source vector elements. All the values in this instruction are signed integer values.
func VqmovnS16(v0 Int16X8) Int8X8 {
	return C.vqmovn_s16(v0)
}

// Signed saturating extract Unsigned Narrow. This instruction reads each signed integer value in the vector of the source SIMD&FP register, saturates the value to an unsigned integer value that is half the original width, places the result into a vector, and writes the vector to the lower or upper half of the destination SIMD&FP register. The destination vector elements are half as long as the source vector elements.
func VqmovunS32(v0 Int32X4) Uint16X4 {
	return C.vqmovun_s32(v0)
}

// Signed saturating extract Unsigned Narrow. This instruction reads each signed integer value in the vector of the source SIMD&FP register, saturates the value to an unsigned integer value that is half the original width, places the result into a vector, and writes the vector to the lower or upper half of the destination SIMD&FP register. The destination vector elements are half as long as the source vector elements.
func VqmovunS64(v0 Int64X2) Uint32X2 {
	return C.vqmovun_s64(v0)
}

// Signed saturating extract Unsigned Narrow. This instruction reads each signed integer value in the vector of the source SIMD&FP register, saturates the value to an unsigned integer value that is half the original width, places the result into a vector, and writes the vector to the lower or upper half of the destination SIMD&FP register. The destination vector elements are half as long as the source vector elements.
func VqmovunS16(v0 Int16X8) Uint8X8 {
	return C.vqmovun_s16(v0)
}

// Signed saturating Negate. This instruction reads each vector element from the source SIMD&FP register, negates each value, places the result into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are signed integer values.
func VqnegqS8(v0 Int8X16) Int8X16 {
	return C.vqnegq_s8(v0)
}

// Signed saturating Negate. This instruction reads each vector element from the source SIMD&FP register, negates each value, places the result into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are signed integer values.
func VqnegqS32(v0 Int32X4) Int32X4 {
	return C.vqnegq_s32(v0)
}

// Signed saturating Negate. This instruction reads each vector element from the source SIMD&FP register, negates each value, places the result into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are signed integer values.
func VqnegqS16(v0 Int16X8) Int16X8 {
	return C.vqnegq_s16(v0)
}

// Signed saturating Negate. This instruction reads each vector element from the source SIMD&FP register, negates each value, places the result into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are signed integer values.
func VqnegS8(v0 Int8X8) Int8X8 {
	return C.vqneg_s8(v0)
}

// Signed saturating Negate. This instruction reads each vector element from the source SIMD&FP register, negates each value, places the result into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are signed integer values.
func VqnegS32(v0 Int32X2) Int32X2 {
	return C.vqneg_s32(v0)
}

// Signed saturating Negate. This instruction reads each vector element from the source SIMD&FP register, negates each value, places the result into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are signed integer values.
func VqnegS16(v0 Int16X4) Int16X4 {
	return C.vqneg_s16(v0)
}

// Signed saturating Rounding Doubling Multiply returning High half. This instruction multiplies the values of corresponding elements of the two source SIMD&FP registers, doubles the results, places the most significant half of the final results into a vector, and writes the vector to the destination SIMD&FP register.
func VqrdmulhqS32(v0 Int32X4, v1 Int32X4) Int32X4 {
	return C.vqrdmulhq_s32(v0, v1)
}

// Signed saturating Rounding Doubling Multiply returning High half. This instruction multiplies the values of corresponding elements of the two source SIMD&FP registers, doubles the results, places the most significant half of the final results into a vector, and writes the vector to the destination SIMD&FP register.
func VqrdmulhqS16(v0 Int16X8, v1 Int16X8) Int16X8 {
	return C.vqrdmulhq_s16(v0, v1)
}

// Signed saturating Rounding Doubling Multiply returning High half. This instruction multiplies the values of corresponding elements of the two source SIMD&FP registers, doubles the results, places the most significant half of the final results into a vector, and writes the vector to the destination SIMD&FP register.
func VqrdmulhS32(v0 Int32X2, v1 Int32X2) Int32X2 {
	return C.vqrdmulh_s32(v0, v1)
}

// Signed saturating Rounding Doubling Multiply returning High half. This instruction multiplies the values of corresponding elements of the two source SIMD&FP registers, doubles the results, places the most significant half of the final results into a vector, and writes the vector to the destination SIMD&FP register.
func VqrdmulhS16(v0 Int16X4, v1 Int16X4) Int16X4 {
	return C.vqrdmulh_s16(v0, v1)
}

// Vector saturating rounding doubling multiply high with scalar
func VqrdmulhqNS32(v0 Int32X4, v1 Int32) Int32X4 {
	return C.vqrdmulhq_n_s32(v0, v1)
}

// Vector saturating rounding doubling multiply high with scalar
func VqrdmulhqNS16(v0 Int16X8, v1 Int16) Int16X8 {
	return C.vqrdmulhq_n_s16(v0, v1)
}

// Vector saturating rounding doubling multiply high with scalar
func VqrdmulhNS32(v0 Int32X2, v1 Int32) Int32X2 {
	return C.vqrdmulh_n_s32(v0, v1)
}

// Vector saturating rounding doubling multiply high with scalar
func VqrdmulhNS16(v0 Int16X4, v1 Int16) Int16X4 {
	return C.vqrdmulh_n_s16(v0, v1)
}

// Unsigned saturating Rounding Shift Left (register). This instruction takes each vector element of the first source SIMD&FP register, shifts the vector element by a value from the least significant byte of the corresponding vector element of the second source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
func VqrshlqU8(v0 Uint8X16, v1 Int8X16) Uint8X16 {
	return C.vqrshlq_u8(v0, v1)
}

// Unsigned saturating Rounding Shift Left (register). This instruction takes each vector element of the first source SIMD&FP register, shifts the vector element by a value from the least significant byte of the corresponding vector element of the second source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
func VqrshlqU32(v0 Uint32X4, v1 Int32X4) Uint32X4 {
	return C.vqrshlq_u32(v0, v1)
}

// Unsigned saturating Rounding Shift Left (register). This instruction takes each vector element of the first source SIMD&FP register, shifts the vector element by a value from the least significant byte of the corresponding vector element of the second source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
func VqrshlqU64(v0 Uint64X2, v1 Int64X2) Uint64X2 {
	return C.vqrshlq_u64(v0, v1)
}

// Unsigned saturating Rounding Shift Left (register). This instruction takes each vector element of the first source SIMD&FP register, shifts the vector element by a value from the least significant byte of the corresponding vector element of the second source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
func VqrshlqU16(v0 Uint16X8, v1 Int16X8) Uint16X8 {
	return C.vqrshlq_u16(v0, v1)
}

// Signed saturating Rounding Shift Left (register). This instruction takes each vector element in the first source SIMD&FP register, shifts it by a value from the least significant byte of the corresponding vector element of the second source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
func VqrshlqS8(v0 Int8X16, v1 Int8X16) Int8X16 {
	return C.vqrshlq_s8(v0, v1)
}

// Signed saturating Rounding Shift Left (register). This instruction takes each vector element in the first source SIMD&FP register, shifts it by a value from the least significant byte of the corresponding vector element of the second source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
func VqrshlqS32(v0 Int32X4, v1 Int32X4) Int32X4 {
	return C.vqrshlq_s32(v0, v1)
}

// Signed saturating Rounding Shift Left (register). This instruction takes each vector element in the first source SIMD&FP register, shifts it by a value from the least significant byte of the corresponding vector element of the second source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
func VqrshlqS64(v0 Int64X2, v1 Int64X2) Int64X2 {
	return C.vqrshlq_s64(v0, v1)
}

// Signed saturating Rounding Shift Left (register). This instruction takes each vector element in the first source SIMD&FP register, shifts it by a value from the least significant byte of the corresponding vector element of the second source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
func VqrshlqS16(v0 Int16X8, v1 Int16X8) Int16X8 {
	return C.vqrshlq_s16(v0, v1)
}

// Unsigned saturating Rounding Shift Left (register). This instruction takes each vector element of the first source SIMD&FP register, shifts the vector element by a value from the least significant byte of the corresponding vector element of the second source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
func VqrshlU8(v0 Uint8X8, v1 Int8X8) Uint8X8 {
	return C.vqrshl_u8(v0, v1)
}

// Unsigned saturating Rounding Shift Left (register). This instruction takes each vector element of the first source SIMD&FP register, shifts the vector element by a value from the least significant byte of the corresponding vector element of the second source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
func VqrshlU32(v0 Uint32X2, v1 Int32X2) Uint32X2 {
	return C.vqrshl_u32(v0, v1)
}

// Unsigned saturating Rounding Shift Left (register). This instruction takes each vector element of the first source SIMD&FP register, shifts the vector element by a value from the least significant byte of the corresponding vector element of the second source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
func VqrshlU64(v0 Uint64X1, v1 Int64X1) Uint64X1 {
	return C.vqrshl_u64(v0, v1)
}

// Unsigned saturating Rounding Shift Left (register). This instruction takes each vector element of the first source SIMD&FP register, shifts the vector element by a value from the least significant byte of the corresponding vector element of the second source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
func VqrshlU16(v0 Uint16X4, v1 Int16X4) Uint16X4 {
	return C.vqrshl_u16(v0, v1)
}

// Signed saturating Rounding Shift Left (register). This instruction takes each vector element in the first source SIMD&FP register, shifts it by a value from the least significant byte of the corresponding vector element of the second source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
func VqrshlS8(v0 Int8X8, v1 Int8X8) Int8X8 {
	return C.vqrshl_s8(v0, v1)
}

// Signed saturating Rounding Shift Left (register). This instruction takes each vector element in the first source SIMD&FP register, shifts it by a value from the least significant byte of the corresponding vector element of the second source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
func VqrshlS32(v0 Int32X2, v1 Int32X2) Int32X2 {
	return C.vqrshl_s32(v0, v1)
}

// Signed saturating Rounding Shift Left (register). This instruction takes each vector element in the first source SIMD&FP register, shifts it by a value from the least significant byte of the corresponding vector element of the second source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
func VqrshlS64(v0 Int64X1, v1 Int64X1) Int64X1 {
	return C.vqrshl_s64(v0, v1)
}

// Signed saturating Rounding Shift Left (register). This instruction takes each vector element in the first source SIMD&FP register, shifts it by a value from the least significant byte of the corresponding vector element of the second source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
func VqrshlS16(v0 Int16X4, v1 Int16X4) Int16X4 {
	return C.vqrshl_s16(v0, v1)
}

// Unsigned saturating Shift Left (register). This instruction takes each element in the vector of the first source SIMD&FP register, shifts the element by a value from the least significant byte of the corresponding element of the second source SIMD&FP register, places the results in a vector, and writes the vector to the destination SIMD&FP register.
func VqshlqU8(v0 Uint8X16, v1 Int8X16) Uint8X16 {
	return C.vqshlq_u8(v0, v1)
}

// Unsigned saturating Shift Left (register). This instruction takes each element in the vector of the first source SIMD&FP register, shifts the element by a value from the least significant byte of the corresponding element of the second source SIMD&FP register, places the results in a vector, and writes the vector to the destination SIMD&FP register.
func VqshlqU32(v0 Uint32X4, v1 Int32X4) Uint32X4 {
	return C.vqshlq_u32(v0, v1)
}

// Unsigned saturating Shift Left (register). This instruction takes each element in the vector of the first source SIMD&FP register, shifts the element by a value from the least significant byte of the corresponding element of the second source SIMD&FP register, places the results in a vector, and writes the vector to the destination SIMD&FP register.
func VqshlqU64(v0 Uint64X2, v1 Int64X2) Uint64X2 {
	return C.vqshlq_u64(v0, v1)
}

// Unsigned saturating Shift Left (register). This instruction takes each element in the vector of the first source SIMD&FP register, shifts the element by a value from the least significant byte of the corresponding element of the second source SIMD&FP register, places the results in a vector, and writes the vector to the destination SIMD&FP register.
func VqshlqU16(v0 Uint16X8, v1 Int16X8) Uint16X8 {
	return C.vqshlq_u16(v0, v1)
}

// Signed saturating Shift Left (register). This instruction takes each element in the vector of the first source SIMD&FP register, shifts each element by a value from the least significant byte of the corresponding element of the second source SIMD&FP register, places the results in a vector, and writes the vector to the destination SIMD&FP register.
func VqshlqS8(v0 Int8X16, v1 Int8X16) Int8X16 {
	return C.vqshlq_s8(v0, v1)
}

// Signed saturating Shift Left (register). This instruction takes each element in the vector of the first source SIMD&FP register, shifts each element by a value from the least significant byte of the corresponding element of the second source SIMD&FP register, places the results in a vector, and writes the vector to the destination SIMD&FP register.
func VqshlqS32(v0 Int32X4, v1 Int32X4) Int32X4 {
	return C.vqshlq_s32(v0, v1)
}

// Signed saturating Shift Left (register). This instruction takes each element in the vector of the first source SIMD&FP register, shifts each element by a value from the least significant byte of the corresponding element of the second source SIMD&FP register, places the results in a vector, and writes the vector to the destination SIMD&FP register.
func VqshlqS64(v0 Int64X2, v1 Int64X2) Int64X2 {
	return C.vqshlq_s64(v0, v1)
}

// Signed saturating Shift Left (register). This instruction takes each element in the vector of the first source SIMD&FP register, shifts each element by a value from the least significant byte of the corresponding element of the second source SIMD&FP register, places the results in a vector, and writes the vector to the destination SIMD&FP register.
func VqshlqS16(v0 Int16X8, v1 Int16X8) Int16X8 {
	return C.vqshlq_s16(v0, v1)
}

// Unsigned saturating Shift Left (register). This instruction takes each element in the vector of the first source SIMD&FP register, shifts the element by a value from the least significant byte of the corresponding element of the second source SIMD&FP register, places the results in a vector, and writes the vector to the destination SIMD&FP register.
func VqshlU8(v0 Uint8X8, v1 Int8X8) Uint8X8 {
	return C.vqshl_u8(v0, v1)
}

// Unsigned saturating Shift Left (register). This instruction takes each element in the vector of the first source SIMD&FP register, shifts the element by a value from the least significant byte of the corresponding element of the second source SIMD&FP register, places the results in a vector, and writes the vector to the destination SIMD&FP register.
func VqshlU32(v0 Uint32X2, v1 Int32X2) Uint32X2 {
	return C.vqshl_u32(v0, v1)
}

// Unsigned saturating Shift Left (register). This instruction takes each element in the vector of the first source SIMD&FP register, shifts the element by a value from the least significant byte of the corresponding element of the second source SIMD&FP register, places the results in a vector, and writes the vector to the destination SIMD&FP register.
func VqshlU64(v0 Uint64X1, v1 Int64X1) Uint64X1 {
	return C.vqshl_u64(v0, v1)
}

// Unsigned saturating Shift Left (register). This instruction takes each element in the vector of the first source SIMD&FP register, shifts the element by a value from the least significant byte of the corresponding element of the second source SIMD&FP register, places the results in a vector, and writes the vector to the destination SIMD&FP register.
func VqshlU16(v0 Uint16X4, v1 Int16X4) Uint16X4 {
	return C.vqshl_u16(v0, v1)
}

// Signed saturating Shift Left (register). This instruction takes each element in the vector of the first source SIMD&FP register, shifts each element by a value from the least significant byte of the corresponding element of the second source SIMD&FP register, places the results in a vector, and writes the vector to the destination SIMD&FP register.
func VqshlS8(v0 Int8X8, v1 Int8X8) Int8X8 {
	return C.vqshl_s8(v0, v1)
}

// Signed saturating Shift Left (register). This instruction takes each element in the vector of the first source SIMD&FP register, shifts each element by a value from the least significant byte of the corresponding element of the second source SIMD&FP register, places the results in a vector, and writes the vector to the destination SIMD&FP register.
func VqshlS32(v0 Int32X2, v1 Int32X2) Int32X2 {
	return C.vqshl_s32(v0, v1)
}

// Signed saturating Shift Left (register). This instruction takes each element in the vector of the first source SIMD&FP register, shifts each element by a value from the least significant byte of the corresponding element of the second source SIMD&FP register, places the results in a vector, and writes the vector to the destination SIMD&FP register.
func VqshlS64(v0 Int64X1, v1 Int64X1) Int64X1 {
	return C.vqshl_s64(v0, v1)
}

// Signed saturating Shift Left (register). This instruction takes each element in the vector of the first source SIMD&FP register, shifts each element by a value from the least significant byte of the corresponding element of the second source SIMD&FP register, places the results in a vector, and writes the vector to the destination SIMD&FP register.
func VqshlS16(v0 Int16X4, v1 Int16X4) Int16X4 {
	return C.vqshl_s16(v0, v1)
}

// Unsigned saturating Subtract. This instruction subtracts the element values of the second source SIMD&FP register from the corresponding element values of the first source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
func VqsubqU8(v0 Uint8X16, v1 Uint8X16) Uint8X16 {
	return C.vqsubq_u8(v0, v1)
}

// Unsigned saturating Subtract. This instruction subtracts the element values of the second source SIMD&FP register from the corresponding element values of the first source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
func VqsubqU32(v0 Uint32X4, v1 Uint32X4) Uint32X4 {
	return C.vqsubq_u32(v0, v1)
}

// Unsigned saturating Subtract. This instruction subtracts the element values of the second source SIMD&FP register from the corresponding element values of the first source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
func VqsubqU64(v0 Uint64X2, v1 Uint64X2) Uint64X2 {
	return C.vqsubq_u64(v0, v1)
}

// Unsigned saturating Subtract. This instruction subtracts the element values of the second source SIMD&FP register from the corresponding element values of the first source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
func VqsubqU16(v0 Uint16X8, v1 Uint16X8) Uint16X8 {
	return C.vqsubq_u16(v0, v1)
}

// Signed saturating Subtract. This instruction subtracts the element values of the second source SIMD&FP register from the corresponding element values of the first source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
func VqsubqS8(v0 Int8X16, v1 Int8X16) Int8X16 {
	return C.vqsubq_s8(v0, v1)
}

// Signed saturating Subtract. This instruction subtracts the element values of the second source SIMD&FP register from the corresponding element values of the first source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
func VqsubqS32(v0 Int32X4, v1 Int32X4) Int32X4 {
	return C.vqsubq_s32(v0, v1)
}

// Signed saturating Subtract. This instruction subtracts the element values of the second source SIMD&FP register from the corresponding element values of the first source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
func VqsubqS64(v0 Int64X2, v1 Int64X2) Int64X2 {
	return C.vqsubq_s64(v0, v1)
}

// Signed saturating Subtract. This instruction subtracts the element values of the second source SIMD&FP register from the corresponding element values of the first source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
func VqsubqS16(v0 Int16X8, v1 Int16X8) Int16X8 {
	return C.vqsubq_s16(v0, v1)
}

// Unsigned saturating Subtract. This instruction subtracts the element values of the second source SIMD&FP register from the corresponding element values of the first source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
func VqsubU8(v0 Uint8X8, v1 Uint8X8) Uint8X8 {
	return C.vqsub_u8(v0, v1)
}

// Unsigned saturating Subtract. This instruction subtracts the element values of the second source SIMD&FP register from the corresponding element values of the first source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
func VqsubU32(v0 Uint32X2, v1 Uint32X2) Uint32X2 {
	return C.vqsub_u32(v0, v1)
}

// Unsigned saturating Subtract. This instruction subtracts the element values of the second source SIMD&FP register from the corresponding element values of the first source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
func VqsubU64(v0 Uint64X1, v1 Uint64X1) Uint64X1 {
	return C.vqsub_u64(v0, v1)
}

// Unsigned saturating Subtract. This instruction subtracts the element values of the second source SIMD&FP register from the corresponding element values of the first source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
func VqsubU16(v0 Uint16X4, v1 Uint16X4) Uint16X4 {
	return C.vqsub_u16(v0, v1)
}

// Signed saturating Subtract. This instruction subtracts the element values of the second source SIMD&FP register from the corresponding element values of the first source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
func VqsubS8(v0 Int8X8, v1 Int8X8) Int8X8 {
	return C.vqsub_s8(v0, v1)
}

// Signed saturating Subtract. This instruction subtracts the element values of the second source SIMD&FP register from the corresponding element values of the first source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
func VqsubS32(v0 Int32X2, v1 Int32X2) Int32X2 {
	return C.vqsub_s32(v0, v1)
}

// Signed saturating Subtract. This instruction subtracts the element values of the second source SIMD&FP register from the corresponding element values of the first source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
func VqsubS64(v0 Int64X1, v1 Int64X1) Int64X1 {
	return C.vqsub_s64(v0, v1)
}

// Signed saturating Subtract. This instruction subtracts the element values of the second source SIMD&FP register from the corresponding element values of the first source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
func VqsubS16(v0 Int16X4, v1 Int16X4) Int16X4 {
	return C.vqsub_s16(v0, v1)
}

// Rounding Add returning High Narrow. This instruction adds each vector element in the first source SIMD&FP register to the corresponding vector element in the second source SIMD&FP register, places the most significant half of the result into a vector, and writes the vector to the lower or upper half of the destination SIMD&FP register.
func VraddhnU32(v0 Uint32X4, v1 Uint32X4) Uint16X4 {
	return C.vraddhn_u32(v0, v1)
}

// Rounding Add returning High Narrow. This instruction adds each vector element in the first source SIMD&FP register to the corresponding vector element in the second source SIMD&FP register, places the most significant half of the result into a vector, and writes the vector to the lower or upper half of the destination SIMD&FP register.
func VraddhnU64(v0 Uint64X2, v1 Uint64X2) Uint32X2 {
	return C.vraddhn_u64(v0, v1)
}

// Rounding Add returning High Narrow. This instruction adds each vector element in the first source SIMD&FP register to the corresponding vector element in the second source SIMD&FP register, places the most significant half of the result into a vector, and writes the vector to the lower or upper half of the destination SIMD&FP register.
func VraddhnU16(v0 Uint16X8, v1 Uint16X8) Uint8X8 {
	return C.vraddhn_u16(v0, v1)
}

// Rounding Add returning High Narrow. This instruction adds each vector element in the first source SIMD&FP register to the corresponding vector element in the second source SIMD&FP register, places the most significant half of the result into a vector, and writes the vector to the lower or upper half of the destination SIMD&FP register.
func VraddhnS32(v0 Int32X4, v1 Int32X4) Int16X4 {
	return C.vraddhn_s32(v0, v1)
}

// Rounding Add returning High Narrow. This instruction adds each vector element in the first source SIMD&FP register to the corresponding vector element in the second source SIMD&FP register, places the most significant half of the result into a vector, and writes the vector to the lower or upper half of the destination SIMD&FP register.
func VraddhnS64(v0 Int64X2, v1 Int64X2) Int32X2 {
	return C.vraddhn_s64(v0, v1)
}

// Rounding Add returning High Narrow. This instruction adds each vector element in the first source SIMD&FP register to the corresponding vector element in the second source SIMD&FP register, places the most significant half of the result into a vector, and writes the vector to the lower or upper half of the destination SIMD&FP register.
func VraddhnS16(v0 Int16X8, v1 Int16X8) Int8X8 {
	return C.vraddhn_s16(v0, v1)
}

// Unsigned Reciprocal Estimate. This instruction reads each vector element from the source SIMD&FP register, calculates an approximate inverse for the unsigned integer value, places the result into a vector, and writes the vector to the destination SIMD&FP register.
func VrecpeqU32(v0 Uint32X4) Uint32X4 {
	return C.vrecpeq_u32(v0)
}

// Floating-point Reciprocal Estimate. This instruction finds an approximate reciprocal estimate for each vector element in the source SIMD&FP register, places the result in a vector, and writes the vector to the destination SIMD&FP register.
func VrecpeqF32(v0 Float32X4) Float32X4 {
	return C.vrecpeq_f32(v0)
}

// Unsigned Reciprocal Estimate. This instruction reads each vector element from the source SIMD&FP register, calculates an approximate inverse for the unsigned integer value, places the result into a vector, and writes the vector to the destination SIMD&FP register.
func VrecpeU32(v0 Uint32X2) Uint32X2 {
	return C.vrecpe_u32(v0)
}

// Floating-point Reciprocal Estimate. This instruction finds an approximate reciprocal estimate for each vector element in the source SIMD&FP register, places the result in a vector, and writes the vector to the destination SIMD&FP register.
func VrecpeF32(v0 Float32X2) Float32X2 {
	return C.vrecpe_f32(v0)
}

// Floating-point Reciprocal Step. This instruction multiplies the corresponding floating-point values in the vectors of the two source SIMD&FP registers, subtracts each of the products from 2.0, places the resulting floating-point values in a vector, and writes the vector to the destination SIMD&FP register.
func VrecpsqF32(v0 Float32X4, v1 Float32X4) Float32X4 {
	return C.vrecpsq_f32(v0, v1)
}

// Floating-point Reciprocal Step. This instruction multiplies the corresponding floating-point values in the vectors of the two source SIMD&FP registers, subtracts each of the products from 2.0, places the resulting floating-point values in a vector, and writes the vector to the destination SIMD&FP register.
func VrecpsF32(v0 Float32X2, v1 Float32X2) Float32X2 {
	return C.vrecps_f32(v0, v1)
}

// Reverse elements in 16-bit halfwords (vector). This instruction reverses the order of 8-bit elements in each halfword of the vector in the source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
func Vrev16P8(v0 Poly8X8) Poly8X8 {
	return C.vrev16_p8(v0)
}

// Reverse elements in 16-bit halfwords (vector). This instruction reverses the order of 8-bit elements in each halfword of the vector in the source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
func Vrev16QP8(v0 Poly8X16) Poly8X16 {
	return C.vrev16q_p8(v0)
}

// Reverse elements in 16-bit halfwords (vector). This instruction reverses the order of 8-bit elements in each halfword of the vector in the source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
func Vrev16QU8(v0 Uint8X16) Uint8X16 {
	return C.vrev16q_u8(v0)
}

// Reverse elements in 16-bit halfwords (vector). This instruction reverses the order of 8-bit elements in each halfword of the vector in the source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
func Vrev16QS8(v0 Int8X16) Int8X16 {
	return C.vrev16q_s8(v0)
}

// Reverse elements in 16-bit halfwords (vector). This instruction reverses the order of 8-bit elements in each halfword of the vector in the source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
func Vrev16U8(v0 Uint8X8) Uint8X8 {
	return C.vrev16_u8(v0)
}

// Reverse elements in 16-bit halfwords (vector). This instruction reverses the order of 8-bit elements in each halfword of the vector in the source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
func Vrev16S8(v0 Int8X8) Int8X8 {
	return C.vrev16_s8(v0)
}

// Reverse elements in 32-bit words (vector). This instruction reverses the order of 8-bit or 16-bit elements in each word of the vector in the source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
func Vrev32P8(v0 Poly8X8) Poly8X8 {
	return C.vrev32_p8(v0)
}

// Reverse elements in 32-bit words (vector). This instruction reverses the order of 8-bit or 16-bit elements in each word of the vector in the source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
func Vrev32P16(v0 Poly16X4) Poly16X4 {
	return C.vrev32_p16(v0)
}

// Reverse elements in 32-bit words (vector). This instruction reverses the order of 8-bit or 16-bit elements in each word of the vector in the source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
func Vrev32QP8(v0 Poly8X16) Poly8X16 {
	return C.vrev32q_p8(v0)
}

// Reverse elements in 32-bit words (vector). This instruction reverses the order of 8-bit or 16-bit elements in each word of the vector in the source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
func Vrev32QP16(v0 Poly16X8) Poly16X8 {
	return C.vrev32q_p16(v0)
}

// Reverse elements in 32-bit words (vector). This instruction reverses the order of 8-bit or 16-bit elements in each word of the vector in the source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
func Vrev32QU8(v0 Uint8X16) Uint8X16 {
	return C.vrev32q_u8(v0)
}

// Reverse elements in 32-bit words (vector). This instruction reverses the order of 8-bit or 16-bit elements in each word of the vector in the source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
func Vrev32QU16(v0 Uint16X8) Uint16X8 {
	return C.vrev32q_u16(v0)
}

// Reverse elements in 32-bit words (vector). This instruction reverses the order of 8-bit or 16-bit elements in each word of the vector in the source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
func Vrev32QS8(v0 Int8X16) Int8X16 {
	return C.vrev32q_s8(v0)
}

// Reverse elements in 32-bit words (vector). This instruction reverses the order of 8-bit or 16-bit elements in each word of the vector in the source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
func Vrev32QS16(v0 Int16X8) Int16X8 {
	return C.vrev32q_s16(v0)
}

// Reverse elements in 32-bit words (vector). This instruction reverses the order of 8-bit or 16-bit elements in each word of the vector in the source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
func Vrev32U8(v0 Uint8X8) Uint8X8 {
	return C.vrev32_u8(v0)
}

// Reverse elements in 32-bit words (vector). This instruction reverses the order of 8-bit or 16-bit elements in each word of the vector in the source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
func Vrev32U16(v0 Uint16X4) Uint16X4 {
	return C.vrev32_u16(v0)
}

// Reverse elements in 32-bit words (vector). This instruction reverses the order of 8-bit or 16-bit elements in each word of the vector in the source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
func Vrev32S8(v0 Int8X8) Int8X8 {
	return C.vrev32_s8(v0)
}

// Reverse elements in 32-bit words (vector). This instruction reverses the order of 8-bit or 16-bit elements in each word of the vector in the source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
func Vrev32S16(v0 Int16X4) Int16X4 {
	return C.vrev32_s16(v0)
}

// Reverse elements in 64-bit doublewords (vector). This instruction reverses the order of 8-bit, 16-bit, or 32-bit elements in each doubleword of the vector in the source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
func Vrev64P8(v0 Poly8X8) Poly8X8 {
	return C.vrev64_p8(v0)
}

// Reverse elements in 64-bit doublewords (vector). This instruction reverses the order of 8-bit, 16-bit, or 32-bit elements in each doubleword of the vector in the source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
func Vrev64P16(v0 Poly16X4) Poly16X4 {
	return C.vrev64_p16(v0)
}

// Reverse elements in 64-bit doublewords (vector). This instruction reverses the order of 8-bit, 16-bit, or 32-bit elements in each doubleword of the vector in the source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
func Vrev64QP8(v0 Poly8X16) Poly8X16 {
	return C.vrev64q_p8(v0)
}

// Reverse elements in 64-bit doublewords (vector). This instruction reverses the order of 8-bit, 16-bit, or 32-bit elements in each doubleword of the vector in the source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
func Vrev64QP16(v0 Poly16X8) Poly16X8 {
	return C.vrev64q_p16(v0)
}

// Reverse elements in 64-bit doublewords (vector). This instruction reverses the order of 8-bit, 16-bit, or 32-bit elements in each doubleword of the vector in the source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
func Vrev64QU8(v0 Uint8X16) Uint8X16 {
	return C.vrev64q_u8(v0)
}

// Reverse elements in 64-bit doublewords (vector). This instruction reverses the order of 8-bit, 16-bit, or 32-bit elements in each doubleword of the vector in the source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
func Vrev64QU32(v0 Uint32X4) Uint32X4 {
	return C.vrev64q_u32(v0)
}

// Reverse elements in 64-bit doublewords (vector). This instruction reverses the order of 8-bit, 16-bit, or 32-bit elements in each doubleword of the vector in the source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
func Vrev64QU16(v0 Uint16X8) Uint16X8 {
	return C.vrev64q_u16(v0)
}

// Reverse elements in 64-bit doublewords (vector). This instruction reverses the order of 8-bit, 16-bit, or 32-bit elements in each doubleword of the vector in the source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
func Vrev64QS8(v0 Int8X16) Int8X16 {
	return C.vrev64q_s8(v0)
}

// Reverse elements in 64-bit doublewords (vector). This instruction reverses the order of 8-bit, 16-bit, or 32-bit elements in each doubleword of the vector in the source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
func Vrev64QF32(v0 Float32X4) Float32X4 {
	return C.vrev64q_f32(v0)
}

// Reverse elements in 64-bit doublewords (vector). This instruction reverses the order of 8-bit, 16-bit, or 32-bit elements in each doubleword of the vector in the source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
func Vrev64QS32(v0 Int32X4) Int32X4 {
	return C.vrev64q_s32(v0)
}

// Reverse elements in 64-bit doublewords (vector). This instruction reverses the order of 8-bit, 16-bit, or 32-bit elements in each doubleword of the vector in the source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
func Vrev64QS16(v0 Int16X8) Int16X8 {
	return C.vrev64q_s16(v0)
}

// Reverse elements in 64-bit doublewords (vector). This instruction reverses the order of 8-bit, 16-bit, or 32-bit elements in each doubleword of the vector in the source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
func Vrev64U8(v0 Uint8X8) Uint8X8 {
	return C.vrev64_u8(v0)
}

// Reverse elements in 64-bit doublewords (vector). This instruction reverses the order of 8-bit, 16-bit, or 32-bit elements in each doubleword of the vector in the source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
func Vrev64U32(v0 Uint32X2) Uint32X2 {
	return C.vrev64_u32(v0)
}

// Reverse elements in 64-bit doublewords (vector). This instruction reverses the order of 8-bit, 16-bit, or 32-bit elements in each doubleword of the vector in the source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
func Vrev64U16(v0 Uint16X4) Uint16X4 {
	return C.vrev64_u16(v0)
}

// Reverse elements in 64-bit doublewords (vector). This instruction reverses the order of 8-bit, 16-bit, or 32-bit elements in each doubleword of the vector in the source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
func Vrev64S8(v0 Int8X8) Int8X8 {
	return C.vrev64_s8(v0)
}

// Reverse elements in 64-bit doublewords (vector). This instruction reverses the order of 8-bit, 16-bit, or 32-bit elements in each doubleword of the vector in the source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
func Vrev64F32(v0 Float32X2) Float32X2 {
	return C.vrev64_f32(v0)
}

// Reverse elements in 64-bit doublewords (vector). This instruction reverses the order of 8-bit, 16-bit, or 32-bit elements in each doubleword of the vector in the source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
func Vrev64S32(v0 Int32X2) Int32X2 {
	return C.vrev64_s32(v0)
}

// Reverse elements in 64-bit doublewords (vector). This instruction reverses the order of 8-bit, 16-bit, or 32-bit elements in each doubleword of the vector in the source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
func Vrev64S16(v0 Int16X4) Int16X4 {
	return C.vrev64_s16(v0)
}

// Unsigned Rounding Halving Add. This instruction adds corresponding unsigned integer values from the two source SIMD&FP registers, shifts each result right one bit, places the results into a vector, and writes the vector to the destination SIMD&FP register.
func VrhaddqU8(v0 Uint8X16, v1 Uint8X16) Uint8X16 {
	return C.vrhaddq_u8(v0, v1)
}

// Unsigned Rounding Halving Add. This instruction adds corresponding unsigned integer values from the two source SIMD&FP registers, shifts each result right one bit, places the results into a vector, and writes the vector to the destination SIMD&FP register.
func VrhaddqU32(v0 Uint32X4, v1 Uint32X4) Uint32X4 {
	return C.vrhaddq_u32(v0, v1)
}

// Unsigned Rounding Halving Add. This instruction adds corresponding unsigned integer values from the two source SIMD&FP registers, shifts each result right one bit, places the results into a vector, and writes the vector to the destination SIMD&FP register.
func VrhaddqU16(v0 Uint16X8, v1 Uint16X8) Uint16X8 {
	return C.vrhaddq_u16(v0, v1)
}

// Signed Rounding Halving Add. This instruction adds corresponding signed integer values from the two source SIMD&FP registers, shifts each result right one bit, places the results into a vector, and writes the vector to the destination SIMD&FP register.
func VrhaddqS8(v0 Int8X16, v1 Int8X16) Int8X16 {
	return C.vrhaddq_s8(v0, v1)
}

// Signed Rounding Halving Add. This instruction adds corresponding signed integer values from the two source SIMD&FP registers, shifts each result right one bit, places the results into a vector, and writes the vector to the destination SIMD&FP register.
func VrhaddqS32(v0 Int32X4, v1 Int32X4) Int32X4 {
	return C.vrhaddq_s32(v0, v1)
}

// Signed Rounding Halving Add. This instruction adds corresponding signed integer values from the two source SIMD&FP registers, shifts each result right one bit, places the results into a vector, and writes the vector to the destination SIMD&FP register.
func VrhaddqS16(v0 Int16X8, v1 Int16X8) Int16X8 {
	return C.vrhaddq_s16(v0, v1)
}

// Unsigned Rounding Halving Add. This instruction adds corresponding unsigned integer values from the two source SIMD&FP registers, shifts each result right one bit, places the results into a vector, and writes the vector to the destination SIMD&FP register.
func VrhaddU8(v0 Uint8X8, v1 Uint8X8) Uint8X8 {
	return C.vrhadd_u8(v0, v1)
}

// Unsigned Rounding Halving Add. This instruction adds corresponding unsigned integer values from the two source SIMD&FP registers, shifts each result right one bit, places the results into a vector, and writes the vector to the destination SIMD&FP register.
func VrhaddU32(v0 Uint32X2, v1 Uint32X2) Uint32X2 {
	return C.vrhadd_u32(v0, v1)
}

// Unsigned Rounding Halving Add. This instruction adds corresponding unsigned integer values from the two source SIMD&FP registers, shifts each result right one bit, places the results into a vector, and writes the vector to the destination SIMD&FP register.
func VrhaddU16(v0 Uint16X4, v1 Uint16X4) Uint16X4 {
	return C.vrhadd_u16(v0, v1)
}

// Signed Rounding Halving Add. This instruction adds corresponding signed integer values from the two source SIMD&FP registers, shifts each result right one bit, places the results into a vector, and writes the vector to the destination SIMD&FP register.
func VrhaddS8(v0 Int8X8, v1 Int8X8) Int8X8 {
	return C.vrhadd_s8(v0, v1)
}

// Signed Rounding Halving Add. This instruction adds corresponding signed integer values from the two source SIMD&FP registers, shifts each result right one bit, places the results into a vector, and writes the vector to the destination SIMD&FP register.
func VrhaddS32(v0 Int32X2, v1 Int32X2) Int32X2 {
	return C.vrhadd_s32(v0, v1)
}

// Signed Rounding Halving Add. This instruction adds corresponding signed integer values from the two source SIMD&FP registers, shifts each result right one bit, places the results into a vector, and writes the vector to the destination SIMD&FP register.
func VrhaddS16(v0 Int16X4, v1 Int16X4) Int16X4 {
	return C.vrhadd_s16(v0, v1)
}

// Unsigned Rounding Shift Left (register). This instruction takes each element in the vector of the first source SIMD&FP register, shifts the vector element by a value from the least significant byte of the corresponding element of the second source SIMD&FP register, places the results in a vector, and writes the vector to the destination SIMD&FP register.
func VrshlqU8(v0 Uint8X16, v1 Int8X16) Uint8X16 {
	return C.vrshlq_u8(v0, v1)
}

// Unsigned Rounding Shift Left (register). This instruction takes each element in the vector of the first source SIMD&FP register, shifts the vector element by a value from the least significant byte of the corresponding element of the second source SIMD&FP register, places the results in a vector, and writes the vector to the destination SIMD&FP register.
func VrshlqU32(v0 Uint32X4, v1 Int32X4) Uint32X4 {
	return C.vrshlq_u32(v0, v1)
}

// Unsigned Rounding Shift Left (register). This instruction takes each element in the vector of the first source SIMD&FP register, shifts the vector element by a value from the least significant byte of the corresponding element of the second source SIMD&FP register, places the results in a vector, and writes the vector to the destination SIMD&FP register.
func VrshlqU64(v0 Uint64X2, v1 Int64X2) Uint64X2 {
	return C.vrshlq_u64(v0, v1)
}

// Unsigned Rounding Shift Left (register). This instruction takes each element in the vector of the first source SIMD&FP register, shifts the vector element by a value from the least significant byte of the corresponding element of the second source SIMD&FP register, places the results in a vector, and writes the vector to the destination SIMD&FP register.
func VrshlqU16(v0 Uint16X8, v1 Int16X8) Uint16X8 {
	return C.vrshlq_u16(v0, v1)
}

// Signed Rounding Shift Left (register). This instruction takes each signed integer value in the vector of the first source SIMD&FP register, shifts it by a value from the least significant byte of the corresponding element of the second source SIMD&FP register, places the results in a vector, and writes the vector to the destination SIMD&FP register.
func VrshlqS8(v0 Int8X16, v1 Int8X16) Int8X16 {
	return C.vrshlq_s8(v0, v1)
}

// Signed Rounding Shift Left (register). This instruction takes each signed integer value in the vector of the first source SIMD&FP register, shifts it by a value from the least significant byte of the corresponding element of the second source SIMD&FP register, places the results in a vector, and writes the vector to the destination SIMD&FP register.
func VrshlqS32(v0 Int32X4, v1 Int32X4) Int32X4 {
	return C.vrshlq_s32(v0, v1)
}

// Signed Rounding Shift Left (register). This instruction takes each signed integer value in the vector of the first source SIMD&FP register, shifts it by a value from the least significant byte of the corresponding element of the second source SIMD&FP register, places the results in a vector, and writes the vector to the destination SIMD&FP register.
func VrshlqS64(v0 Int64X2, v1 Int64X2) Int64X2 {
	return C.vrshlq_s64(v0, v1)
}

// Signed Rounding Shift Left (register). This instruction takes each signed integer value in the vector of the first source SIMD&FP register, shifts it by a value from the least significant byte of the corresponding element of the second source SIMD&FP register, places the results in a vector, and writes the vector to the destination SIMD&FP register.
func VrshlqS16(v0 Int16X8, v1 Int16X8) Int16X8 {
	return C.vrshlq_s16(v0, v1)
}

// Unsigned Rounding Shift Left (register). This instruction takes each element in the vector of the first source SIMD&FP register, shifts the vector element by a value from the least significant byte of the corresponding element of the second source SIMD&FP register, places the results in a vector, and writes the vector to the destination SIMD&FP register.
func VrshlU8(v0 Uint8X8, v1 Int8X8) Uint8X8 {
	return C.vrshl_u8(v0, v1)
}

// Unsigned Rounding Shift Left (register). This instruction takes each element in the vector of the first source SIMD&FP register, shifts the vector element by a value from the least significant byte of the corresponding element of the second source SIMD&FP register, places the results in a vector, and writes the vector to the destination SIMD&FP register.
func VrshlU32(v0 Uint32X2, v1 Int32X2) Uint32X2 {
	return C.vrshl_u32(v0, v1)
}

// Unsigned Rounding Shift Left (register). This instruction takes each element in the vector of the first source SIMD&FP register, shifts the vector element by a value from the least significant byte of the corresponding element of the second source SIMD&FP register, places the results in a vector, and writes the vector to the destination SIMD&FP register.
func VrshlU64(v0 Uint64X1, v1 Int64X1) Uint64X1 {
	return C.vrshl_u64(v0, v1)
}

// Unsigned Rounding Shift Left (register). This instruction takes each element in the vector of the first source SIMD&FP register, shifts the vector element by a value from the least significant byte of the corresponding element of the second source SIMD&FP register, places the results in a vector, and writes the vector to the destination SIMD&FP register.
func VrshlU16(v0 Uint16X4, v1 Int16X4) Uint16X4 {
	return C.vrshl_u16(v0, v1)
}

// Signed Rounding Shift Left (register). This instruction takes each signed integer value in the vector of the first source SIMD&FP register, shifts it by a value from the least significant byte of the corresponding element of the second source SIMD&FP register, places the results in a vector, and writes the vector to the destination SIMD&FP register.
func VrshlS8(v0 Int8X8, v1 Int8X8) Int8X8 {
	return C.vrshl_s8(v0, v1)
}

// Signed Rounding Shift Left (register). This instruction takes each signed integer value in the vector of the first source SIMD&FP register, shifts it by a value from the least significant byte of the corresponding element of the second source SIMD&FP register, places the results in a vector, and writes the vector to the destination SIMD&FP register.
func VrshlS32(v0 Int32X2, v1 Int32X2) Int32X2 {
	return C.vrshl_s32(v0, v1)
}

// Signed Rounding Shift Left (register). This instruction takes each signed integer value in the vector of the first source SIMD&FP register, shifts it by a value from the least significant byte of the corresponding element of the second source SIMD&FP register, places the results in a vector, and writes the vector to the destination SIMD&FP register.
func VrshlS64(v0 Int64X1, v1 Int64X1) Int64X1 {
	return C.vrshl_s64(v0, v1)
}

// Signed Rounding Shift Left (register). This instruction takes each signed integer value in the vector of the first source SIMD&FP register, shifts it by a value from the least significant byte of the corresponding element of the second source SIMD&FP register, places the results in a vector, and writes the vector to the destination SIMD&FP register.
func VrshlS16(v0 Int16X4, v1 Int16X4) Int16X4 {
	return C.vrshl_s16(v0, v1)
}

// Unsigned Reciprocal Square Root Estimate. This instruction reads each vector element from the source SIMD&FP register, calculates an approximate inverse square root for each value, places the result into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are unsigned integer values.
func VrsqrteqU32(v0 Uint32X4) Uint32X4 {
	return C.vrsqrteq_u32(v0)
}

// Floating-point Reciprocal Square Root Estimate. This instruction calculates an approximate square root for each vector element in the source SIMD&FP register, places the result in a vector, and writes the vector to the destination SIMD&FP register.
func VrsqrteqF32(v0 Float32X4) Float32X4 {
	return C.vrsqrteq_f32(v0)
}

// Unsigned Reciprocal Square Root Estimate. This instruction reads each vector element from the source SIMD&FP register, calculates an approximate inverse square root for each value, places the result into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are unsigned integer values.
func VrsqrteU32(v0 Uint32X2) Uint32X2 {
	return C.vrsqrte_u32(v0)
}

// Floating-point Reciprocal Square Root Estimate. This instruction calculates an approximate square root for each vector element in the source SIMD&FP register, places the result in a vector, and writes the vector to the destination SIMD&FP register.
func VrsqrteF32(v0 Float32X2) Float32X2 {
	return C.vrsqrte_f32(v0)
}

// Floating-point Reciprocal Square Root Step. This instruction multiplies corresponding floating-point values in the vectors of the two source SIMD&FP registers, subtracts each of the products from 3.0, divides these results by 2.0, places the results into a vector, and writes the vector to the destination SIMD&FP register.
func VrsqrtsqF32(v0 Float32X4, v1 Float32X4) Float32X4 {
	return C.vrsqrtsq_f32(v0, v1)
}

// Floating-point Reciprocal Square Root Step. This instruction multiplies corresponding floating-point values in the vectors of the two source SIMD&FP registers, subtracts each of the products from 3.0, divides these results by 2.0, places the results into a vector, and writes the vector to the destination SIMD&FP register.
func VrsqrtsF32(v0 Float32X2, v1 Float32X2) Float32X2 {
	return C.vrsqrts_f32(v0, v1)
}

// Rounding Subtract returning High Narrow. This instruction subtracts each vector element of the second source SIMD&FP register from the corresponding vector element of the first source SIMD&FP register, places the most significant half of the result into a vector, and writes the vector to the lower or upper half of the destination SIMD&FP register.
func VrsubhnU32(v0 Uint32X4, v1 Uint32X4) Uint16X4 {
	return C.vrsubhn_u32(v0, v1)
}

// Rounding Subtract returning High Narrow. This instruction subtracts each vector element of the second source SIMD&FP register from the corresponding vector element of the first source SIMD&FP register, places the most significant half of the result into a vector, and writes the vector to the lower or upper half of the destination SIMD&FP register.
func VrsubhnU64(v0 Uint64X2, v1 Uint64X2) Uint32X2 {
	return C.vrsubhn_u64(v0, v1)
}

// Rounding Subtract returning High Narrow. This instruction subtracts each vector element of the second source SIMD&FP register from the corresponding vector element of the first source SIMD&FP register, places the most significant half of the result into a vector, and writes the vector to the lower or upper half of the destination SIMD&FP register.
func VrsubhnU16(v0 Uint16X8, v1 Uint16X8) Uint8X8 {
	return C.vrsubhn_u16(v0, v1)
}

// Rounding Subtract returning High Narrow. This instruction subtracts each vector element of the second source SIMD&FP register from the corresponding vector element of the first source SIMD&FP register, places the most significant half of the result into a vector, and writes the vector to the lower or upper half of the destination SIMD&FP register.
func VrsubhnS32(v0 Int32X4, v1 Int32X4) Int16X4 {
	return C.vrsubhn_s32(v0, v1)
}

// Rounding Subtract returning High Narrow. This instruction subtracts each vector element of the second source SIMD&FP register from the corresponding vector element of the first source SIMD&FP register, places the most significant half of the result into a vector, and writes the vector to the lower or upper half of the destination SIMD&FP register.
func VrsubhnS64(v0 Int64X2, v1 Int64X2) Int32X2 {
	return C.vrsubhn_s64(v0, v1)
}

// Rounding Subtract returning High Narrow. This instruction subtracts each vector element of the second source SIMD&FP register from the corresponding vector element of the first source SIMD&FP register, places the most significant half of the result into a vector, and writes the vector to the lower or upper half of the destination SIMD&FP register.
func VrsubhnS16(v0 Int16X8, v1 Int16X8) Int8X8 {
	return C.vrsubhn_s16(v0, v1)
}

// Unsigned Shift Left (register). This instruction takes each element in the vector of the first source SIMD&FP register, shifts each element by a value from the least significant byte of the corresponding element of the second source SIMD&FP register, places the results in a vector, and writes the vector to the destination SIMD&FP register.
func VshlqU8(v0 Uint8X16, v1 Int8X16) Uint8X16 {
	return C.vshlq_u8(v0, v1)
}

// Unsigned Shift Left (register). This instruction takes each element in the vector of the first source SIMD&FP register, shifts each element by a value from the least significant byte of the corresponding element of the second source SIMD&FP register, places the results in a vector, and writes the vector to the destination SIMD&FP register.
func VshlqU32(v0 Uint32X4, v1 Int32X4) Uint32X4 {
	return C.vshlq_u32(v0, v1)
}

// Unsigned Shift Left (register). This instruction takes each element in the vector of the first source SIMD&FP register, shifts each element by a value from the least significant byte of the corresponding element of the second source SIMD&FP register, places the results in a vector, and writes the vector to the destination SIMD&FP register.
func VshlqU64(v0 Uint64X2, v1 Int64X2) Uint64X2 {
	return C.vshlq_u64(v0, v1)
}

// Unsigned Shift Left (register). This instruction takes each element in the vector of the first source SIMD&FP register, shifts each element by a value from the least significant byte of the corresponding element of the second source SIMD&FP register, places the results in a vector, and writes the vector to the destination SIMD&FP register.
func VshlqU16(v0 Uint16X8, v1 Int16X8) Uint16X8 {
	return C.vshlq_u16(v0, v1)
}

// Signed Shift Left (register). This instruction takes each signed integer value in the vector of the first source SIMD&FP register, shifts each value by a value from the least significant byte of the corresponding element of the second source SIMD&FP register, places the results in a vector, and writes the vector to the destination SIMD&FP register.
func VshlqS8(v0 Int8X16, v1 Int8X16) Int8X16 {
	return C.vshlq_s8(v0, v1)
}

// Signed Shift Left (register). This instruction takes each signed integer value in the vector of the first source SIMD&FP register, shifts each value by a value from the least significant byte of the corresponding element of the second source SIMD&FP register, places the results in a vector, and writes the vector to the destination SIMD&FP register.
func VshlqS32(v0 Int32X4, v1 Int32X4) Int32X4 {
	return C.vshlq_s32(v0, v1)
}

// Signed Shift Left (register). This instruction takes each signed integer value in the vector of the first source SIMD&FP register, shifts each value by a value from the least significant byte of the corresponding element of the second source SIMD&FP register, places the results in a vector, and writes the vector to the destination SIMD&FP register.
func VshlqS64(v0 Int64X2, v1 Int64X2) Int64X2 {
	return C.vshlq_s64(v0, v1)
}

// Signed Shift Left (register). This instruction takes each signed integer value in the vector of the first source SIMD&FP register, shifts each value by a value from the least significant byte of the corresponding element of the second source SIMD&FP register, places the results in a vector, and writes the vector to the destination SIMD&FP register.
func VshlqS16(v0 Int16X8, v1 Int16X8) Int16X8 {
	return C.vshlq_s16(v0, v1)
}

// Unsigned Shift Left (register). This instruction takes each element in the vector of the first source SIMD&FP register, shifts each element by a value from the least significant byte of the corresponding element of the second source SIMD&FP register, places the results in a vector, and writes the vector to the destination SIMD&FP register.
func VshlU8(v0 Uint8X8, v1 Int8X8) Uint8X8 {
	return C.vshl_u8(v0, v1)
}

// Unsigned Shift Left (register). This instruction takes each element in the vector of the first source SIMD&FP register, shifts each element by a value from the least significant byte of the corresponding element of the second source SIMD&FP register, places the results in a vector, and writes the vector to the destination SIMD&FP register.
func VshlU32(v0 Uint32X2, v1 Int32X2) Uint32X2 {
	return C.vshl_u32(v0, v1)
}

// Unsigned Shift Left (register). This instruction takes each element in the vector of the first source SIMD&FP register, shifts each element by a value from the least significant byte of the corresponding element of the second source SIMD&FP register, places the results in a vector, and writes the vector to the destination SIMD&FP register.
func VshlU64(v0 Uint64X1, v1 Int64X1) Uint64X1 {
	return C.vshl_u64(v0, v1)
}

// Unsigned Shift Left (register). This instruction takes each element in the vector of the first source SIMD&FP register, shifts each element by a value from the least significant byte of the corresponding element of the second source SIMD&FP register, places the results in a vector, and writes the vector to the destination SIMD&FP register.
func VshlU16(v0 Uint16X4, v1 Int16X4) Uint16X4 {
	return C.vshl_u16(v0, v1)
}

// Signed Shift Left (register). This instruction takes each signed integer value in the vector of the first source SIMD&FP register, shifts each value by a value from the least significant byte of the corresponding element of the second source SIMD&FP register, places the results in a vector, and writes the vector to the destination SIMD&FP register.
func VshlS8(v0 Int8X8, v1 Int8X8) Int8X8 {
	return C.vshl_s8(v0, v1)
}

// Signed Shift Left (register). This instruction takes each signed integer value in the vector of the first source SIMD&FP register, shifts each value by a value from the least significant byte of the corresponding element of the second source SIMD&FP register, places the results in a vector, and writes the vector to the destination SIMD&FP register.
func VshlS32(v0 Int32X2, v1 Int32X2) Int32X2 {
	return C.vshl_s32(v0, v1)
}

// Signed Shift Left (register). This instruction takes each signed integer value in the vector of the first source SIMD&FP register, shifts each value by a value from the least significant byte of the corresponding element of the second source SIMD&FP register, places the results in a vector, and writes the vector to the destination SIMD&FP register.
func VshlS64(v0 Int64X1, v1 Int64X1) Int64X1 {
	return C.vshl_s64(v0, v1)
}

// Signed Shift Left (register). This instruction takes each signed integer value in the vector of the first source SIMD&FP register, shifts each value by a value from the least significant byte of the corresponding element of the second source SIMD&FP register, places the results in a vector, and writes the vector to the destination SIMD&FP register.
func VshlS16(v0 Int16X4, v1 Int16X4) Int16X4 {
	return C.vshl_s16(v0, v1)
}

// Subtract (vector). This instruction subtracts each vector element in the second source SIMD&FP register from the corresponding vector element in the first source SIMD&FP register, places the result into a vector, and writes the vector to the destination SIMD&FP register.
func VsubqU8(v0 Uint8X16, v1 Uint8X16) Uint8X16 {
	return C.vsubq_u8(v0, v1)
}

// Subtract (vector). This instruction subtracts each vector element in the second source SIMD&FP register from the corresponding vector element in the first source SIMD&FP register, places the result into a vector, and writes the vector to the destination SIMD&FP register.
func VsubqU32(v0 Uint32X4, v1 Uint32X4) Uint32X4 {
	return C.vsubq_u32(v0, v1)
}

// Subtract (vector). This instruction subtracts each vector element in the second source SIMD&FP register from the corresponding vector element in the first source SIMD&FP register, places the result into a vector, and writes the vector to the destination SIMD&FP register.
func VsubqU64(v0 Uint64X2, v1 Uint64X2) Uint64X2 {
	return C.vsubq_u64(v0, v1)
}

// Subtract (vector). This instruction subtracts each vector element in the second source SIMD&FP register from the corresponding vector element in the first source SIMD&FP register, places the result into a vector, and writes the vector to the destination SIMD&FP register.
func VsubqU16(v0 Uint16X8, v1 Uint16X8) Uint16X8 {
	return C.vsubq_u16(v0, v1)
}

// Subtract (vector). This instruction subtracts each vector element in the second source SIMD&FP register from the corresponding vector element in the first source SIMD&FP register, places the result into a vector, and writes the vector to the destination SIMD&FP register.
func VsubqS8(v0 Int8X16, v1 Int8X16) Int8X16 {
	return C.vsubq_s8(v0, v1)
}

// Floating-point Subtract (vector). This instruction subtracts the elements in the vector in the second source SIMD&FP register, from the corresponding elements in the vector in the first source SIMD&FP register, places each result into elements of a vector, and writes the vector to the destination SIMD&FP register.
func VsubqF32(v0 Float32X4, v1 Float32X4) Float32X4 {
	return C.vsubq_f32(v0, v1)
}

// Subtract (vector). This instruction subtracts each vector element in the second source SIMD&FP register from the corresponding vector element in the first source SIMD&FP register, places the result into a vector, and writes the vector to the destination SIMD&FP register.
func VsubqS32(v0 Int32X4, v1 Int32X4) Int32X4 {
	return C.vsubq_s32(v0, v1)
}

// Subtract (vector). This instruction subtracts each vector element in the second source SIMD&FP register from the corresponding vector element in the first source SIMD&FP register, places the result into a vector, and writes the vector to the destination SIMD&FP register.
func VsubqS64(v0 Int64X2, v1 Int64X2) Int64X2 {
	return C.vsubq_s64(v0, v1)
}

// Subtract (vector). This instruction subtracts each vector element in the second source SIMD&FP register from the corresponding vector element in the first source SIMD&FP register, places the result into a vector, and writes the vector to the destination SIMD&FP register.
func VsubqS16(v0 Int16X8, v1 Int16X8) Int16X8 {
	return C.vsubq_s16(v0, v1)
}

// Subtract (vector). This instruction subtracts each vector element in the second source SIMD&FP register from the corresponding vector element in the first source SIMD&FP register, places the result into a vector, and writes the vector to the destination SIMD&FP register.
func VsubU8(v0 Uint8X8, v1 Uint8X8) Uint8X8 {
	return C.vsub_u8(v0, v1)
}

// Subtract (vector). This instruction subtracts each vector element in the second source SIMD&FP register from the corresponding vector element in the first source SIMD&FP register, places the result into a vector, and writes the vector to the destination SIMD&FP register.
func VsubU32(v0 Uint32X2, v1 Uint32X2) Uint32X2 {
	return C.vsub_u32(v0, v1)
}

// Subtract (vector). This instruction subtracts each vector element in the second source SIMD&FP register from the corresponding vector element in the first source SIMD&FP register, places the result into a vector, and writes the vector to the destination SIMD&FP register.
func VsubU64(v0 Uint64X1, v1 Uint64X1) Uint64X1 {
	return C.vsub_u64(v0, v1)
}

// Subtract (vector). This instruction subtracts each vector element in the second source SIMD&FP register from the corresponding vector element in the first source SIMD&FP register, places the result into a vector, and writes the vector to the destination SIMD&FP register.
func VsubU16(v0 Uint16X4, v1 Uint16X4) Uint16X4 {
	return C.vsub_u16(v0, v1)
}

// Subtract (vector). This instruction subtracts each vector element in the second source SIMD&FP register from the corresponding vector element in the first source SIMD&FP register, places the result into a vector, and writes the vector to the destination SIMD&FP register.
func VsubS8(v0 Int8X8, v1 Int8X8) Int8X8 {
	return C.vsub_s8(v0, v1)
}

// Floating-point Subtract (vector). This instruction subtracts the elements in the vector in the second source SIMD&FP register, from the corresponding elements in the vector in the first source SIMD&FP register, places each result into elements of a vector, and writes the vector to the destination SIMD&FP register.
func VsubF32(v0 Float32X2, v1 Float32X2) Float32X2 {
	return C.vsub_f32(v0, v1)
}

// Subtract (vector). This instruction subtracts each vector element in the second source SIMD&FP register from the corresponding vector element in the first source SIMD&FP register, places the result into a vector, and writes the vector to the destination SIMD&FP register.
func VsubS32(v0 Int32X2, v1 Int32X2) Int32X2 {
	return C.vsub_s32(v0, v1)
}

// Subtract (vector). This instruction subtracts each vector element in the second source SIMD&FP register from the corresponding vector element in the first source SIMD&FP register, places the result into a vector, and writes the vector to the destination SIMD&FP register.
func VsubS64(v0 Int64X1, v1 Int64X1) Int64X1 {
	return C.vsub_s64(v0, v1)
}

// Subtract (vector). This instruction subtracts each vector element in the second source SIMD&FP register from the corresponding vector element in the first source SIMD&FP register, places the result into a vector, and writes the vector to the destination SIMD&FP register.
func VsubS16(v0 Int16X4, v1 Int16X4) Int16X4 {
	return C.vsub_s16(v0, v1)
}

// Subtract returning High Narrow. This instruction subtracts each vector element in the second source SIMD&FP register from the corresponding vector element in the first source SIMD&FP register, places the most significant half of the result into a vector, and writes the vector to the lower or upper half of the destination SIMD&FP register. All the values in this instruction are signed integer values.
func VsubhnU32(v0 Uint32X4, v1 Uint32X4) Uint16X4 {
	return C.vsubhn_u32(v0, v1)
}

// Subtract returning High Narrow. This instruction subtracts each vector element in the second source SIMD&FP register from the corresponding vector element in the first source SIMD&FP register, places the most significant half of the result into a vector, and writes the vector to the lower or upper half of the destination SIMD&FP register. All the values in this instruction are signed integer values.
func VsubhnU64(v0 Uint64X2, v1 Uint64X2) Uint32X2 {
	return C.vsubhn_u64(v0, v1)
}

// Subtract returning High Narrow. This instruction subtracts each vector element in the second source SIMD&FP register from the corresponding vector element in the first source SIMD&FP register, places the most significant half of the result into a vector, and writes the vector to the lower or upper half of the destination SIMD&FP register. All the values in this instruction are signed integer values.
func VsubhnU16(v0 Uint16X8, v1 Uint16X8) Uint8X8 {
	return C.vsubhn_u16(v0, v1)
}

// Subtract returning High Narrow. This instruction subtracts each vector element in the second source SIMD&FP register from the corresponding vector element in the first source SIMD&FP register, places the most significant half of the result into a vector, and writes the vector to the lower or upper half of the destination SIMD&FP register. All the values in this instruction are signed integer values.
func VsubhnS32(v0 Int32X4, v1 Int32X4) Int16X4 {
	return C.vsubhn_s32(v0, v1)
}

// Subtract returning High Narrow. This instruction subtracts each vector element in the second source SIMD&FP register from the corresponding vector element in the first source SIMD&FP register, places the most significant half of the result into a vector, and writes the vector to the lower or upper half of the destination SIMD&FP register. All the values in this instruction are signed integer values.
func VsubhnS64(v0 Int64X2, v1 Int64X2) Int32X2 {
	return C.vsubhn_s64(v0, v1)
}

// Subtract returning High Narrow. This instruction subtracts each vector element in the second source SIMD&FP register from the corresponding vector element in the first source SIMD&FP register, places the most significant half of the result into a vector, and writes the vector to the lower or upper half of the destination SIMD&FP register. All the values in this instruction are signed integer values.
func VsubhnS16(v0 Int16X8, v1 Int16X8) Int8X8 {
	return C.vsubhn_s16(v0, v1)
}

// Unsigned Subtract Long. This instruction subtracts each vector element in the lower or upper half of the second source SIMD&FP register from the corresponding vector element of the first source SIMD&FP register, places the result into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are unsigned integer values. The destination vector elements are twice as long as the source vector elements.
func VsublU8(v0 Uint8X8, v1 Uint8X8) Uint16X8 {
	return C.vsubl_u8(v0, v1)
}

// Unsigned Subtract Long. This instruction subtracts each vector element in the lower or upper half of the second source SIMD&FP register from the corresponding vector element of the first source SIMD&FP register, places the result into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are unsigned integer values. The destination vector elements are twice as long as the source vector elements.
func VsublU32(v0 Uint32X2, v1 Uint32X2) Uint64X2 {
	return C.vsubl_u32(v0, v1)
}

// Unsigned Subtract Long. This instruction subtracts each vector element in the lower or upper half of the second source SIMD&FP register from the corresponding vector element of the first source SIMD&FP register, places the result into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are unsigned integer values. The destination vector elements are twice as long as the source vector elements.
func VsublU16(v0 Uint16X4, v1 Uint16X4) Uint32X4 {
	return C.vsubl_u16(v0, v1)
}

// Signed Subtract Long. This instruction subtracts each vector element in the lower or upper half of the second source SIMD&FP register from the corresponding vector element of the first source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are signed integer values. The destination vector elements are twice as long as the source vector elements.
func VsublS8(v0 Int8X8, v1 Int8X8) Int16X8 {
	return C.vsubl_s8(v0, v1)
}

// Signed Subtract Long. This instruction subtracts each vector element in the lower or upper half of the second source SIMD&FP register from the corresponding vector element of the first source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are signed integer values. The destination vector elements are twice as long as the source vector elements.
func VsublS32(v0 Int32X2, v1 Int32X2) Int64X2 {
	return C.vsubl_s32(v0, v1)
}

// Signed Subtract Long. This instruction subtracts each vector element in the lower or upper half of the second source SIMD&FP register from the corresponding vector element of the first source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are signed integer values. The destination vector elements are twice as long as the source vector elements.
func VsublS16(v0 Int16X4, v1 Int16X4) Int32X4 {
	return C.vsubl_s16(v0, v1)
}

// Unsigned Subtract Wide. This instruction subtracts each vector element of the second source SIMD&FP register from the corresponding vector element in the lower or upper half of the first source SIMD&FP register, places the result in a vector, and writes the vector to the SIMD&FP destination register. All the values in this instruction are signed integer values.
func VsubwU8(v0 Uint16X8, v1 Uint8X8) Uint16X8 {
	return C.vsubw_u8(v0, v1)
}

// Unsigned Subtract Wide. This instruction subtracts each vector element of the second source SIMD&FP register from the corresponding vector element in the lower or upper half of the first source SIMD&FP register, places the result in a vector, and writes the vector to the SIMD&FP destination register. All the values in this instruction are signed integer values.
func VsubwU32(v0 Uint64X2, v1 Uint32X2) Uint64X2 {
	return C.vsubw_u32(v0, v1)
}

// Unsigned Subtract Wide. This instruction subtracts each vector element of the second source SIMD&FP register from the corresponding vector element in the lower or upper half of the first source SIMD&FP register, places the result in a vector, and writes the vector to the SIMD&FP destination register. All the values in this instruction are signed integer values.
func VsubwU16(v0 Uint32X4, v1 Uint16X4) Uint32X4 {
	return C.vsubw_u16(v0, v1)
}

// Signed Subtract Wide. This instruction subtracts each vector element in the lower or upper half of the second source SIMD&FP register from the corresponding vector element in the first source SIMD&FP register, places the result in a vector, and writes the vector to the SIMD&FP destination register. All the values in this instruction are signed integer values.
func VsubwS8(v0 Int16X8, v1 Int8X8) Int16X8 {
	return C.vsubw_s8(v0, v1)
}

// Signed Subtract Wide. This instruction subtracts each vector element in the lower or upper half of the second source SIMD&FP register from the corresponding vector element in the first source SIMD&FP register, places the result in a vector, and writes the vector to the SIMD&FP destination register. All the values in this instruction are signed integer values.
func VsubwS32(v0 Int64X2, v1 Int32X2) Int64X2 {
	return C.vsubw_s32(v0, v1)
}

// Signed Subtract Wide. This instruction subtracts each vector element in the lower or upper half of the second source SIMD&FP register from the corresponding vector element in the first source SIMD&FP register, places the result in a vector, and writes the vector to the SIMD&FP destination register. All the values in this instruction are signed integer values.
func VsubwS16(v0 Int32X4, v1 Int16X4) Int32X4 {
	return C.vsubw_s16(v0, v1)
}

// Table vector Lookup. This instruction reads each value from the vector elements in the index source SIMD&FP register, uses each result as an index to perform a lookup in a table of bytes that is described by one to four source table SIMD&FP registers, places the lookup result in a vector, and writes the vector to the destination SIMD&FP register. If an index is out of range for the table, the result for that lookup is 0. If more than one source register is used to describe the table, the first source register describes the lowest bytes of the table.
func Vtbl1P8(v0 Poly8X8, v1 Uint8X8) Poly8X8 {
	return C.vtbl1_p8(v0, v1)
}

// Table vector Lookup. This instruction reads each value from the vector elements in the index source SIMD&FP register, uses each result as an index to perform a lookup in a table of bytes that is described by one to four source table SIMD&FP registers, places the lookup result in a vector, and writes the vector to the destination SIMD&FP register. If an index is out of range for the table, the result for that lookup is 0. If more than one source register is used to describe the table, the first source register describes the lowest bytes of the table.
func Vtbl1U8(v0 Uint8X8, v1 Uint8X8) Uint8X8 {
	return C.vtbl1_u8(v0, v1)
}

// Table vector Lookup. This instruction reads each value from the vector elements in the index source SIMD&FP register, uses each result as an index to perform a lookup in a table of bytes that is described by one to four source table SIMD&FP registers, places the lookup result in a vector, and writes the vector to the destination SIMD&FP register. If an index is out of range for the table, the result for that lookup is 0. If more than one source register is used to describe the table, the first source register describes the lowest bytes of the table.
func Vtbl1S8(v0 Int8X8, v1 Int8X8) Int8X8 {
	return C.vtbl1_s8(v0, v1)
}

// Table vector Lookup. This instruction reads each value from the vector elements in the index source SIMD&FP register, uses each result as an index to perform a lookup in a table of bytes that is described by one to four source table SIMD&FP registers, places the lookup result in a vector, and writes the vector to the destination SIMD&FP register. If an index is out of range for the table, the result for that lookup is 0. If more than one source register is used to describe the table, the first source register describes the lowest bytes of the table.
func Vtbl2P8(v0 Poly8X8X2, v1 Uint8X8) Poly8X8 {
	return C.vtbl2_p8(v0, v1)
}

// Table vector Lookup. This instruction reads each value from the vector elements in the index source SIMD&FP register, uses each result as an index to perform a lookup in a table of bytes that is described by one to four source table SIMD&FP registers, places the lookup result in a vector, and writes the vector to the destination SIMD&FP register. If an index is out of range for the table, the result for that lookup is 0. If more than one source register is used to describe the table, the first source register describes the lowest bytes of the table.
func Vtbl2U8(v0 Uint8X8X2, v1 Uint8X8) Uint8X8 {
	return C.vtbl2_u8(v0, v1)
}

// Table vector Lookup. This instruction reads each value from the vector elements in the index source SIMD&FP register, uses each result as an index to perform a lookup in a table of bytes that is described by one to four source table SIMD&FP registers, places the lookup result in a vector, and writes the vector to the destination SIMD&FP register. If an index is out of range for the table, the result for that lookup is 0. If more than one source register is used to describe the table, the first source register describes the lowest bytes of the table.
func Vtbl2S8(v0 Int8X8X2, v1 Int8X8) Int8X8 {
	return C.vtbl2_s8(v0, v1)
}

// Table vector Lookup. This instruction reads each value from the vector elements in the index source SIMD&FP register, uses each result as an index to perform a lookup in a table of bytes that is described by one to four source table SIMD&FP registers, places the lookup result in a vector, and writes the vector to the destination SIMD&FP register. If an index is out of range for the table, the result for that lookup is 0. If more than one source register is used to describe the table, the first source register describes the lowest bytes of the table.
func Vtbl3P8(v0 Poly8X8X3, v1 Uint8X8) Poly8X8 {
	return C.vtbl3_p8(v0, v1)
}

// Table vector Lookup. This instruction reads each value from the vector elements in the index source SIMD&FP register, uses each result as an index to perform a lookup in a table of bytes that is described by one to four source table SIMD&FP registers, places the lookup result in a vector, and writes the vector to the destination SIMD&FP register. If an index is out of range for the table, the result for that lookup is 0. If more than one source register is used to describe the table, the first source register describes the lowest bytes of the table.
func Vtbl3U8(v0 Uint8X8X3, v1 Uint8X8) Uint8X8 {
	return C.vtbl3_u8(v0, v1)
}

// Table vector Lookup. This instruction reads each value from the vector elements in the index source SIMD&FP register, uses each result as an index to perform a lookup in a table of bytes that is described by one to four source table SIMD&FP registers, places the lookup result in a vector, and writes the vector to the destination SIMD&FP register. If an index is out of range for the table, the result for that lookup is 0. If more than one source register is used to describe the table, the first source register describes the lowest bytes of the table.
func Vtbl3S8(v0 Int8X8X3, v1 Int8X8) Int8X8 {
	return C.vtbl3_s8(v0, v1)
}

// Table vector Lookup. This instruction reads each value from the vector elements in the index source SIMD&FP register, uses each result as an index to perform a lookup in a table of bytes that is described by one to four source table SIMD&FP registers, places the lookup result in a vector, and writes the vector to the destination SIMD&FP register. If an index is out of range for the table, the result for that lookup is 0. If more than one source register is used to describe the table, the first source register describes the lowest bytes of the table.
func Vtbl4P8(v0 Poly8X8X4, v1 Uint8X8) Poly8X8 {
	return C.vtbl4_p8(v0, v1)
}

// Table vector Lookup. This instruction reads each value from the vector elements in the index source SIMD&FP register, uses each result as an index to perform a lookup in a table of bytes that is described by one to four source table SIMD&FP registers, places the lookup result in a vector, and writes the vector to the destination SIMD&FP register. If an index is out of range for the table, the result for that lookup is 0. If more than one source register is used to describe the table, the first source register describes the lowest bytes of the table.
func Vtbl4U8(v0 Uint8X8X4, v1 Uint8X8) Uint8X8 {
	return C.vtbl4_u8(v0, v1)
}

// Table vector Lookup. This instruction reads each value from the vector elements in the index source SIMD&FP register, uses each result as an index to perform a lookup in a table of bytes that is described by one to four source table SIMD&FP registers, places the lookup result in a vector, and writes the vector to the destination SIMD&FP register. If an index is out of range for the table, the result for that lookup is 0. If more than one source register is used to describe the table, the first source register describes the lowest bytes of the table.
func Vtbl4S8(v0 Int8X8X4, v1 Int8X8) Int8X8 {
	return C.vtbl4_s8(v0, v1)
}

// Table vector lookup extension
func Vtbx1P8(v0 Poly8X8, v1 Poly8X8, v2 Uint8X8) Poly8X8 {
	return C.vtbx1_p8(v0, v1, v2)
}

// Table vector lookup extension
func Vtbx1U8(v0 Uint8X8, v1 Uint8X8, v2 Uint8X8) Uint8X8 {
	return C.vtbx1_u8(v0, v1, v2)
}

// Table vector lookup extension
func Vtbx1S8(v0 Int8X8, v1 Int8X8, v2 Int8X8) Int8X8 {
	return C.vtbx1_s8(v0, v1, v2)
}

// Table vector lookup extension. This instruction reads each value from the vector elements in the index source SIMD&FP register, uses each result as an index to perform a lookup in a table of bytes that is described by one to four source table SIMD&FP registers, places the lookup result in a vector, and writes the vector to the destination SIMD&FP register. If an index is out of range for the table, the existing value in the vector element of the destination register is left unchanged. If more than one source register is used to describe the table, the first source register describes the lowest bytes of the table.
func Vtbx2P8(v0 Poly8X8, v1 Poly8X8X2, v2 Uint8X8) Poly8X8 {
	return C.vtbx2_p8(v0, v1, v2)
}

// Table vector lookup extension. This instruction reads each value from the vector elements in the index source SIMD&FP register, uses each result as an index to perform a lookup in a table of bytes that is described by one to four source table SIMD&FP registers, places the lookup result in a vector, and writes the vector to the destination SIMD&FP register. If an index is out of range for the table, the existing value in the vector element of the destination register is left unchanged. If more than one source register is used to describe the table, the first source register describes the lowest bytes of the table.
func Vtbx2U8(v0 Uint8X8, v1 Uint8X8X2, v2 Uint8X8) Uint8X8 {
	return C.vtbx2_u8(v0, v1, v2)
}

// Table vector lookup extension. This instruction reads each value from the vector elements in the index source SIMD&FP register, uses each result as an index to perform a lookup in a table of bytes that is described by one to four source table SIMD&FP registers, places the lookup result in a vector, and writes the vector to the destination SIMD&FP register. If an index is out of range for the table, the existing value in the vector element of the destination register is left unchanged. If more than one source register is used to describe the table, the first source register describes the lowest bytes of the table.
func Vtbx2S8(v0 Int8X8, v1 Int8X8X2, v2 Int8X8) Int8X8 {
	return C.vtbx2_s8(v0, v1, v2)
}

// Table vector lookup extension
func Vtbx3P8(v0 Poly8X8, v1 Poly8X8X3, v2 Uint8X8) Poly8X8 {
	return C.vtbx3_p8(v0, v1, v2)
}

// Table vector lookup extension
func Vtbx3U8(v0 Uint8X8, v1 Uint8X8X3, v2 Uint8X8) Uint8X8 {
	return C.vtbx3_u8(v0, v1, v2)
}

// Table vector lookup extension
func Vtbx3S8(v0 Int8X8, v1 Int8X8X3, v2 Int8X8) Int8X8 {
	return C.vtbx3_s8(v0, v1, v2)
}

// Table vector lookup extension. This instruction reads each value from the vector elements in the index source SIMD&FP register, uses each result as an index to perform a lookup in a table of bytes that is described by one to four source table SIMD&FP registers, places the lookup result in a vector, and writes the vector to the destination SIMD&FP register. If an index is out of range for the table, the existing value in the vector element of the destination register is left unchanged. If more than one source register is used to describe the table, the first source register describes the lowest bytes of the table.
func Vtbx4P8(v0 Poly8X8, v1 Poly8X8X4, v2 Uint8X8) Poly8X8 {
	return C.vtbx4_p8(v0, v1, v2)
}

// Table vector lookup extension. This instruction reads each value from the vector elements in the index source SIMD&FP register, uses each result as an index to perform a lookup in a table of bytes that is described by one to four source table SIMD&FP registers, places the lookup result in a vector, and writes the vector to the destination SIMD&FP register. If an index is out of range for the table, the existing value in the vector element of the destination register is left unchanged. If more than one source register is used to describe the table, the first source register describes the lowest bytes of the table.
func Vtbx4U8(v0 Uint8X8, v1 Uint8X8X4, v2 Uint8X8) Uint8X8 {
	return C.vtbx4_u8(v0, v1, v2)
}

// Table vector lookup extension. This instruction reads each value from the vector elements in the index source SIMD&FP register, uses each result as an index to perform a lookup in a table of bytes that is described by one to four source table SIMD&FP registers, places the lookup result in a vector, and writes the vector to the destination SIMD&FP register. If an index is out of range for the table, the existing value in the vector element of the destination register is left unchanged. If more than one source register is used to describe the table, the first source register describes the lowest bytes of the table.
func Vtbx4S8(v0 Int8X8, v1 Int8X8X4, v2 Int8X8) Int8X8 {
	return C.vtbx4_s8(v0, v1, v2)
}

// Transpose elements
func VtrnP8(v0 Poly8X8, v1 Poly8X8) Poly8X8X2 {
	return C.vtrn_p8(v0, v1)
}

// Transpose elements
func VtrnP16(v0 Poly16X4, v1 Poly16X4) Poly16X4X2 {
	return C.vtrn_p16(v0, v1)
}

// Transpose elements
func VtrnqP8(v0 Poly8X16, v1 Poly8X16) Poly8X16X2 {
	return C.vtrnq_p8(v0, v1)
}

// Transpose elements
func VtrnqP16(v0 Poly16X8, v1 Poly16X8) Poly16X8X2 {
	return C.vtrnq_p16(v0, v1)
}

// Transpose elements
func VtrnqU8(v0 Uint8X16, v1 Uint8X16) Uint8X16X2 {
	return C.vtrnq_u8(v0, v1)
}

// Transpose elements
func VtrnqU32(v0 Uint32X4, v1 Uint32X4) Uint32X4X2 {
	return C.vtrnq_u32(v0, v1)
}

// Transpose elements
func VtrnqU16(v0 Uint16X8, v1 Uint16X8) Uint16X8X2 {
	return C.vtrnq_u16(v0, v1)
}

// Transpose elements
func VtrnqS8(v0 Int8X16, v1 Int8X16) Int8X16X2 {
	return C.vtrnq_s8(v0, v1)
}

// Transpose elements
func VtrnqF32(v0 Float32X4, v1 Float32X4) Float32X4X2 {
	return C.vtrnq_f32(v0, v1)
}

// Transpose elements
func VtrnqS32(v0 Int32X4, v1 Int32X4) Int32X4X2 {
	return C.vtrnq_s32(v0, v1)
}

// Transpose elements
func VtrnqS16(v0 Int16X8, v1 Int16X8) Int16X8X2 {
	return C.vtrnq_s16(v0, v1)
}

// Transpose elements
func VtrnU8(v0 Uint8X8, v1 Uint8X8) Uint8X8X2 {
	return C.vtrn_u8(v0, v1)
}

// Transpose elements
func VtrnU32(v0 Uint32X2, v1 Uint32X2) Uint32X2X2 {
	return C.vtrn_u32(v0, v1)
}

// Transpose elements
func VtrnU16(v0 Uint16X4, v1 Uint16X4) Uint16X4X2 {
	return C.vtrn_u16(v0, v1)
}

// Transpose elements
func VtrnS8(v0 Int8X8, v1 Int8X8) Int8X8X2 {
	return C.vtrn_s8(v0, v1)
}

// Transpose elements
func VtrnF32(v0 Float32X2, v1 Float32X2) Float32X2X2 {
	return C.vtrn_f32(v0, v1)
}

// Transpose elements
func VtrnS32(v0 Int32X2, v1 Int32X2) Int32X2X2 {
	return C.vtrn_s32(v0, v1)
}

// Transpose elements
func VtrnS16(v0 Int16X4, v1 Int16X4) Int16X4X2 {
	return C.vtrn_s16(v0, v1)
}

// Compare bitwise Test bits nonzero (vector). This instruction reads each vector element in the first source SIMD&FP register, performs an AND with the corresponding vector element in the second source SIMD&FP register, and if the result is not zero, sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
func VtstP8(v0 Poly8X8, v1 Poly8X8) Uint8X8 {
	return C.vtst_p8(v0, v1)
}

// vtst_p16
func VtstP16(v0 Poly16X4, v1 Poly16X4) Uint16X4 {
	return C.vtst_p16(v0, v1)
}

// Compare bitwise Test bits nonzero (vector). This instruction reads each vector element in the first source SIMD&FP register, performs an AND with the corresponding vector element in the second source SIMD&FP register, and if the result is not zero, sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
func VtstqP8(v0 Poly8X16, v1 Poly8X16) Uint8X16 {
	return C.vtstq_p8(v0, v1)
}

// vtstq_p16
func VtstqP16(v0 Poly16X8, v1 Poly16X8) Uint16X8 {
	return C.vtstq_p16(v0, v1)
}

// Compare bitwise Test bits nonzero (vector). This instruction reads each vector element in the first source SIMD&FP register, performs an AND with the corresponding vector element in the second source SIMD&FP register, and if the result is not zero, sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
func VtstqU8(v0 Uint8X16, v1 Uint8X16) Uint8X16 {
	return C.vtstq_u8(v0, v1)
}

// Compare bitwise Test bits nonzero (vector). This instruction reads each vector element in the first source SIMD&FP register, performs an AND with the corresponding vector element in the second source SIMD&FP register, and if the result is not zero, sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
func VtstqU32(v0 Uint32X4, v1 Uint32X4) Uint32X4 {
	return C.vtstq_u32(v0, v1)
}

// Compare bitwise Test bits nonzero (vector). This instruction reads each vector element in the first source SIMD&FP register, performs an AND with the corresponding vector element in the second source SIMD&FP register, and if the result is not zero, sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
func VtstqU16(v0 Uint16X8, v1 Uint16X8) Uint16X8 {
	return C.vtstq_u16(v0, v1)
}

// Compare bitwise Test bits nonzero (vector). This instruction reads each vector element in the first source SIMD&FP register, performs an AND with the corresponding vector element in the second source SIMD&FP register, and if the result is not zero, sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
func VtstqS8(v0 Int8X16, v1 Int8X16) Uint8X16 {
	return C.vtstq_s8(v0, v1)
}

// Compare bitwise Test bits nonzero (vector). This instruction reads each vector element in the first source SIMD&FP register, performs an AND with the corresponding vector element in the second source SIMD&FP register, and if the result is not zero, sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
func VtstqS32(v0 Int32X4, v1 Int32X4) Uint32X4 {
	return C.vtstq_s32(v0, v1)
}

// Compare bitwise Test bits nonzero (vector). This instruction reads each vector element in the first source SIMD&FP register, performs an AND with the corresponding vector element in the second source SIMD&FP register, and if the result is not zero, sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
func VtstqS16(v0 Int16X8, v1 Int16X8) Uint16X8 {
	return C.vtstq_s16(v0, v1)
}

// Compare bitwise Test bits nonzero (vector). This instruction reads each vector element in the first source SIMD&FP register, performs an AND with the corresponding vector element in the second source SIMD&FP register, and if the result is not zero, sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
func VtstU8(v0 Uint8X8, v1 Uint8X8) Uint8X8 {
	return C.vtst_u8(v0, v1)
}

// Compare bitwise Test bits nonzero (vector). This instruction reads each vector element in the first source SIMD&FP register, performs an AND with the corresponding vector element in the second source SIMD&FP register, and if the result is not zero, sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
func VtstU32(v0 Uint32X2, v1 Uint32X2) Uint32X2 {
	return C.vtst_u32(v0, v1)
}

// Compare bitwise Test bits nonzero (vector). This instruction reads each vector element in the first source SIMD&FP register, performs an AND with the corresponding vector element in the second source SIMD&FP register, and if the result is not zero, sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
func VtstU16(v0 Uint16X4, v1 Uint16X4) Uint16X4 {
	return C.vtst_u16(v0, v1)
}

// Compare bitwise Test bits nonzero (vector). This instruction reads each vector element in the first source SIMD&FP register, performs an AND with the corresponding vector element in the second source SIMD&FP register, and if the result is not zero, sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
func VtstS8(v0 Int8X8, v1 Int8X8) Uint8X8 {
	return C.vtst_s8(v0, v1)
}

// Compare bitwise Test bits nonzero (vector). This instruction reads each vector element in the first source SIMD&FP register, performs an AND with the corresponding vector element in the second source SIMD&FP register, and if the result is not zero, sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
func VtstS32(v0 Int32X2, v1 Int32X2) Uint32X2 {
	return C.vtst_s32(v0, v1)
}

// Compare bitwise Test bits nonzero (vector). This instruction reads each vector element in the first source SIMD&FP register, performs an AND with the corresponding vector element in the second source SIMD&FP register, and if the result is not zero, sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
func VtstS16(v0 Int16X4, v1 Int16X4) Uint16X4 {
	return C.vtst_s16(v0, v1)
}

// Unzip vectors (secondary). This instruction reads corresponding odd-numbered vector elements from the two source SIMD&FP registers, places the result from the first source register into consecutive elements in the lower half of a vector, and the result from the second source register into consecutive elements in the upper half of a vector, and writes the vector to the destination SIMD&FP register.
func VuzpP8(v0 Poly8X8, v1 Poly8X8) Poly8X8X2 {
	return C.vuzp_p8(v0, v1)
}

// Unzip vectors (secondary). This instruction reads corresponding odd-numbered vector elements from the two source SIMD&FP registers, places the result from the first source register into consecutive elements in the lower half of a vector, and the result from the second source register into consecutive elements in the upper half of a vector, and writes the vector to the destination SIMD&FP register.
func VuzpP16(v0 Poly16X4, v1 Poly16X4) Poly16X4X2 {
	return C.vuzp_p16(v0, v1)
}

// Unzip vectors (secondary). This instruction reads corresponding odd-numbered vector elements from the two source SIMD&FP registers, places the result from the first source register into consecutive elements in the lower half of a vector, and the result from the second source register into consecutive elements in the upper half of a vector, and writes the vector to the destination SIMD&FP register.
func VuzpqP8(v0 Poly8X16, v1 Poly8X16) Poly8X16X2 {
	return C.vuzpq_p8(v0, v1)
}

// Unzip vectors (secondary). This instruction reads corresponding odd-numbered vector elements from the two source SIMD&FP registers, places the result from the first source register into consecutive elements in the lower half of a vector, and the result from the second source register into consecutive elements in the upper half of a vector, and writes the vector to the destination SIMD&FP register.
func VuzpqP16(v0 Poly16X8, v1 Poly16X8) Poly16X8X2 {
	return C.vuzpq_p16(v0, v1)
}

// Unzip vectors (secondary). This instruction reads corresponding odd-numbered vector elements from the two source SIMD&FP registers, places the result from the first source register into consecutive elements in the lower half of a vector, and the result from the second source register into consecutive elements in the upper half of a vector, and writes the vector to the destination SIMD&FP register.
func VuzpqU8(v0 Uint8X16, v1 Uint8X16) Uint8X16X2 {
	return C.vuzpq_u8(v0, v1)
}

// Unzip vectors (secondary). This instruction reads corresponding odd-numbered vector elements from the two source SIMD&FP registers, places the result from the first source register into consecutive elements in the lower half of a vector, and the result from the second source register into consecutive elements in the upper half of a vector, and writes the vector to the destination SIMD&FP register.
func VuzpqU32(v0 Uint32X4, v1 Uint32X4) Uint32X4X2 {
	return C.vuzpq_u32(v0, v1)
}

// Unzip vectors (secondary). This instruction reads corresponding odd-numbered vector elements from the two source SIMD&FP registers, places the result from the first source register into consecutive elements in the lower half of a vector, and the result from the second source register into consecutive elements in the upper half of a vector, and writes the vector to the destination SIMD&FP register.
func VuzpqU16(v0 Uint16X8, v1 Uint16X8) Uint16X8X2 {
	return C.vuzpq_u16(v0, v1)
}

// Unzip vectors (secondary). This instruction reads corresponding odd-numbered vector elements from the two source SIMD&FP registers, places the result from the first source register into consecutive elements in the lower half of a vector, and the result from the second source register into consecutive elements in the upper half of a vector, and writes the vector to the destination SIMD&FP register.
func VuzpqS8(v0 Int8X16, v1 Int8X16) Int8X16X2 {
	return C.vuzpq_s8(v0, v1)
}

// Unzip vectors (secondary). This instruction reads corresponding odd-numbered vector elements from the two source SIMD&FP registers, places the result from the first source register into consecutive elements in the lower half of a vector, and the result from the second source register into consecutive elements in the upper half of a vector, and writes the vector to the destination SIMD&FP register.
func VuzpqF32(v0 Float32X4, v1 Float32X4) Float32X4X2 {
	return C.vuzpq_f32(v0, v1)
}

// Unzip vectors (secondary). This instruction reads corresponding odd-numbered vector elements from the two source SIMD&FP registers, places the result from the first source register into consecutive elements in the lower half of a vector, and the result from the second source register into consecutive elements in the upper half of a vector, and writes the vector to the destination SIMD&FP register.
func VuzpqS32(v0 Int32X4, v1 Int32X4) Int32X4X2 {
	return C.vuzpq_s32(v0, v1)
}

// Unzip vectors (secondary). This instruction reads corresponding odd-numbered vector elements from the two source SIMD&FP registers, places the result from the first source register into consecutive elements in the lower half of a vector, and the result from the second source register into consecutive elements in the upper half of a vector, and writes the vector to the destination SIMD&FP register.
func VuzpqS16(v0 Int16X8, v1 Int16X8) Int16X8X2 {
	return C.vuzpq_s16(v0, v1)
}

// Unzip vectors (secondary). This instruction reads corresponding odd-numbered vector elements from the two source SIMD&FP registers, places the result from the first source register into consecutive elements in the lower half of a vector, and the result from the second source register into consecutive elements in the upper half of a vector, and writes the vector to the destination SIMD&FP register.
func VuzpU8(v0 Uint8X8, v1 Uint8X8) Uint8X8X2 {
	return C.vuzp_u8(v0, v1)
}

// Unzip vectors (secondary). This instruction reads corresponding odd-numbered vector elements from the two source SIMD&FP registers, places the result from the first source register into consecutive elements in the lower half of a vector, and the result from the second source register into consecutive elements in the upper half of a vector, and writes the vector to the destination SIMD&FP register.
func VuzpU32(v0 Uint32X2, v1 Uint32X2) Uint32X2X2 {
	return C.vuzp_u32(v0, v1)
}

// Unzip vectors (secondary). This instruction reads corresponding odd-numbered vector elements from the two source SIMD&FP registers, places the result from the first source register into consecutive elements in the lower half of a vector, and the result from the second source register into consecutive elements in the upper half of a vector, and writes the vector to the destination SIMD&FP register.
func VuzpU16(v0 Uint16X4, v1 Uint16X4) Uint16X4X2 {
	return C.vuzp_u16(v0, v1)
}

// Unzip vectors (secondary). This instruction reads corresponding odd-numbered vector elements from the two source SIMD&FP registers, places the result from the first source register into consecutive elements in the lower half of a vector, and the result from the second source register into consecutive elements in the upper half of a vector, and writes the vector to the destination SIMD&FP register.
func VuzpS8(v0 Int8X8, v1 Int8X8) Int8X8X2 {
	return C.vuzp_s8(v0, v1)
}

// Unzip vectors (secondary). This instruction reads corresponding odd-numbered vector elements from the two source SIMD&FP registers, places the result from the first source register into consecutive elements in the lower half of a vector, and the result from the second source register into consecutive elements in the upper half of a vector, and writes the vector to the destination SIMD&FP register.
func VuzpF32(v0 Float32X2, v1 Float32X2) Float32X2X2 {
	return C.vuzp_f32(v0, v1)
}

// Unzip vectors (secondary). This instruction reads corresponding odd-numbered vector elements from the two source SIMD&FP registers, places the result from the first source register into consecutive elements in the lower half of a vector, and the result from the second source register into consecutive elements in the upper half of a vector, and writes the vector to the destination SIMD&FP register.
func VuzpS32(v0 Int32X2, v1 Int32X2) Int32X2X2 {
	return C.vuzp_s32(v0, v1)
}

// Unzip vectors (secondary). This instruction reads corresponding odd-numbered vector elements from the two source SIMD&FP registers, places the result from the first source register into consecutive elements in the lower half of a vector, and the result from the second source register into consecutive elements in the upper half of a vector, and writes the vector to the destination SIMD&FP register.
func VuzpS16(v0 Int16X4, v1 Int16X4) Int16X4X2 {
	return C.vuzp_s16(v0, v1)
}

// Zip vectors (secondary). This instruction reads adjacent vector elements from the upper half of two source SIMD&FP registers as pairs, interleaves the pairs and places them into a vector, and writes the vector to the destination SIMD&FP register. The first pair from the first source register is placed into the two lowest vector elements, with subsequent pairs taken alternately from each source register.
func VzipP8(v0 Poly8X8, v1 Poly8X8) Poly8X8X2 {
	return C.vzip_p8(v0, v1)
}

// Zip vectors (secondary). This instruction reads adjacent vector elements from the upper half of two source SIMD&FP registers as pairs, interleaves the pairs and places them into a vector, and writes the vector to the destination SIMD&FP register. The first pair from the first source register is placed into the two lowest vector elements, with subsequent pairs taken alternately from each source register.
func VzipP16(v0 Poly16X4, v1 Poly16X4) Poly16X4X2 {
	return C.vzip_p16(v0, v1)
}

// Zip vectors (secondary). This instruction reads adjacent vector elements from the upper half of two source SIMD&FP registers as pairs, interleaves the pairs and places them into a vector, and writes the vector to the destination SIMD&FP register. The first pair from the first source register is placed into the two lowest vector elements, with subsequent pairs taken alternately from each source register.
func VzipqP8(v0 Poly8X16, v1 Poly8X16) Poly8X16X2 {
	return C.vzipq_p8(v0, v1)
}

// Zip vectors (secondary). This instruction reads adjacent vector elements from the upper half of two source SIMD&FP registers as pairs, interleaves the pairs and places them into a vector, and writes the vector to the destination SIMD&FP register. The first pair from the first source register is placed into the two lowest vector elements, with subsequent pairs taken alternately from each source register.
func VzipqP16(v0 Poly16X8, v1 Poly16X8) Poly16X8X2 {
	return C.vzipq_p16(v0, v1)
}

// Zip vectors (secondary). This instruction reads adjacent vector elements from the upper half of two source SIMD&FP registers as pairs, interleaves the pairs and places them into a vector, and writes the vector to the destination SIMD&FP register. The first pair from the first source register is placed into the two lowest vector elements, with subsequent pairs taken alternately from each source register.
func VzipqU8(v0 Uint8X16, v1 Uint8X16) Uint8X16X2 {
	return C.vzipq_u8(v0, v1)
}

// Zip vectors (secondary). This instruction reads adjacent vector elements from the upper half of two source SIMD&FP registers as pairs, interleaves the pairs and places them into a vector, and writes the vector to the destination SIMD&FP register. The first pair from the first source register is placed into the two lowest vector elements, with subsequent pairs taken alternately from each source register.
func VzipqU32(v0 Uint32X4, v1 Uint32X4) Uint32X4X2 {
	return C.vzipq_u32(v0, v1)
}

// Zip vectors (secondary). This instruction reads adjacent vector elements from the upper half of two source SIMD&FP registers as pairs, interleaves the pairs and places them into a vector, and writes the vector to the destination SIMD&FP register. The first pair from the first source register is placed into the two lowest vector elements, with subsequent pairs taken alternately from each source register.
func VzipqU16(v0 Uint16X8, v1 Uint16X8) Uint16X8X2 {
	return C.vzipq_u16(v0, v1)
}

// Zip vectors (secondary). This instruction reads adjacent vector elements from the upper half of two source SIMD&FP registers as pairs, interleaves the pairs and places them into a vector, and writes the vector to the destination SIMD&FP register. The first pair from the first source register is placed into the two lowest vector elements, with subsequent pairs taken alternately from each source register.
func VzipqS8(v0 Int8X16, v1 Int8X16) Int8X16X2 {
	return C.vzipq_s8(v0, v1)
}

// Zip vectors (secondary). This instruction reads adjacent vector elements from the upper half of two source SIMD&FP registers as pairs, interleaves the pairs and places them into a vector, and writes the vector to the destination SIMD&FP register. The first pair from the first source register is placed into the two lowest vector elements, with subsequent pairs taken alternately from each source register.
func VzipqF32(v0 Float32X4, v1 Float32X4) Float32X4X2 {
	return C.vzipq_f32(v0, v1)
}

// Zip vectors (secondary). This instruction reads adjacent vector elements from the upper half of two source SIMD&FP registers as pairs, interleaves the pairs and places them into a vector, and writes the vector to the destination SIMD&FP register. The first pair from the first source register is placed into the two lowest vector elements, with subsequent pairs taken alternately from each source register.
func VzipqS32(v0 Int32X4, v1 Int32X4) Int32X4X2 {
	return C.vzipq_s32(v0, v1)
}

// Zip vectors (secondary). This instruction reads adjacent vector elements from the upper half of two source SIMD&FP registers as pairs, interleaves the pairs and places them into a vector, and writes the vector to the destination SIMD&FP register. The first pair from the first source register is placed into the two lowest vector elements, with subsequent pairs taken alternately from each source register.
func VzipqS16(v0 Int16X8, v1 Int16X8) Int16X8X2 {
	return C.vzipq_s16(v0, v1)
}

// Zip vectors (secondary). This instruction reads adjacent vector elements from the upper half of two source SIMD&FP registers as pairs, interleaves the pairs and places them into a vector, and writes the vector to the destination SIMD&FP register. The first pair from the first source register is placed into the two lowest vector elements, with subsequent pairs taken alternately from each source register.
func VzipU8(v0 Uint8X8, v1 Uint8X8) Uint8X8X2 {
	return C.vzip_u8(v0, v1)
}

// Zip vectors (secondary). This instruction reads adjacent vector elements from the upper half of two source SIMD&FP registers as pairs, interleaves the pairs and places them into a vector, and writes the vector to the destination SIMD&FP register. The first pair from the first source register is placed into the two lowest vector elements, with subsequent pairs taken alternately from each source register.
func VzipU32(v0 Uint32X2, v1 Uint32X2) Uint32X2X2 {
	return C.vzip_u32(v0, v1)
}

// Zip vectors (secondary). This instruction reads adjacent vector elements from the upper half of two source SIMD&FP registers as pairs, interleaves the pairs and places them into a vector, and writes the vector to the destination SIMD&FP register. The first pair from the first source register is placed into the two lowest vector elements, with subsequent pairs taken alternately from each source register.
func VzipU16(v0 Uint16X4, v1 Uint16X4) Uint16X4X2 {
	return C.vzip_u16(v0, v1)
}

// Zip vectors (secondary). This instruction reads adjacent vector elements from the upper half of two source SIMD&FP registers as pairs, interleaves the pairs and places them into a vector, and writes the vector to the destination SIMD&FP register. The first pair from the first source register is placed into the two lowest vector elements, with subsequent pairs taken alternately from each source register.
func VzipS8(v0 Int8X8, v1 Int8X8) Int8X8X2 {
	return C.vzip_s8(v0, v1)
}

// Zip vectors (secondary). This instruction reads adjacent vector elements from the upper half of two source SIMD&FP registers as pairs, interleaves the pairs and places them into a vector, and writes the vector to the destination SIMD&FP register. The first pair from the first source register is placed into the two lowest vector elements, with subsequent pairs taken alternately from each source register.
func VzipF32(v0 Float32X2, v1 Float32X2) Float32X2X2 {
	return C.vzip_f32(v0, v1)
}

// Zip vectors (secondary). This instruction reads adjacent vector elements from the upper half of two source SIMD&FP registers as pairs, interleaves the pairs and places them into a vector, and writes the vector to the destination SIMD&FP register. The first pair from the first source register is placed into the two lowest vector elements, with subsequent pairs taken alternately from each source register.
func VzipS32(v0 Int32X2, v1 Int32X2) Int32X2X2 {
	return C.vzip_s32(v0, v1)
}

// Zip vectors (secondary). This instruction reads adjacent vector elements from the upper half of two source SIMD&FP registers as pairs, interleaves the pairs and places them into a vector, and writes the vector to the destination SIMD&FP register. The first pair from the first source register is placed into the two lowest vector elements, with subsequent pairs taken alternately from each source register.
func VzipS16(v0 Int16X4, v1 Int16X4) Int16X4X2 {
	return C.vzip_s16(v0, v1)
}

// Floating-point Convert to Signed integer, rounding to nearest with ties to Away (vector). This instruction converts each element in a vector from a floating-point value to a signed integer value using the Round to Nearest with Ties to Away rounding mode and writes the result to the SIMD&FP destination register.
func VcvtaqS32F32(v0 Float32X4) Int32X4 {
	return C.vcvtaq_s32_f32(v0)
}

// Floating-point Convert to Signed integer, rounding to nearest with ties to Away (vector). This instruction converts each element in a vector from a floating-point value to a signed integer value using the Round to Nearest with Ties to Away rounding mode and writes the result to the SIMD&FP destination register.
func VcvtaS32F32(v0 Float32X2) Int32X2 {
	return C.vcvta_s32_f32(v0)
}

// Floating-point Convert to Unsigned integer, rounding to nearest with ties to Away (vector). This instruction converts each element in a vector from a floating-point value to an unsigned integer value using the Round to Nearest with Ties to Away rounding mode and writes the result to the SIMD&FP destination register.
func VcvtaqU32F32(v0 Float32X4) Uint32X4 {
	return C.vcvtaq_u32_f32(v0)
}

// Floating-point Convert to Unsigned integer, rounding to nearest with ties to Away (vector). This instruction converts each element in a vector from a floating-point value to an unsigned integer value using the Round to Nearest with Ties to Away rounding mode and writes the result to the SIMD&FP destination register.
func VcvtaU32F32(v0 Float32X2) Uint32X2 {
	return C.vcvta_u32_f32(v0)
}

// Floating-point Convert to Signed integer, rounding toward Minus infinity (vector). This instruction converts a scalar or each element in a vector from a floating-point value to a signed integer value using the Round towards Minus Infinity rounding mode, and writes the result to the SIMD&FP destination register.
func VcvtmqS32F32(v0 Float32X4) Int32X4 {
	return C.vcvtmq_s32_f32(v0)
}

// Floating-point Convert to Signed integer, rounding toward Minus infinity (vector). This instruction converts a scalar or each element in a vector from a floating-point value to a signed integer value using the Round towards Minus Infinity rounding mode, and writes the result to the SIMD&FP destination register.
func VcvtmS32F32(v0 Float32X2) Int32X2 {
	return C.vcvtm_s32_f32(v0)
}

// Floating-point Convert to Unsigned integer, rounding toward Minus infinity (vector). This instruction converts a scalar or each element in a vector from a floating-point value to an unsigned integer value using the Round towards Minus Infinity rounding mode, and writes the result to the SIMD&FP destination register.
func VcvtmqU32F32(v0 Float32X4) Uint32X4 {
	return C.vcvtmq_u32_f32(v0)
}

// Floating-point Convert to Unsigned integer, rounding toward Minus infinity (vector). This instruction converts a scalar or each element in a vector from a floating-point value to an unsigned integer value using the Round towards Minus Infinity rounding mode, and writes the result to the SIMD&FP destination register.
func VcvtmU32F32(v0 Float32X2) Uint32X2 {
	return C.vcvtm_u32_f32(v0)
}

// Floating-point Convert to Signed integer, rounding to nearest with ties to even (vector). This instruction converts a scalar or each element in a vector from a floating-point value to a signed integer value using the Round to Nearest rounding mode, and writes the result to the SIMD&FP destination register.
func VcvtnqS32F32(v0 Float32X4) Int32X4 {
	return C.vcvtnq_s32_f32(v0)
}

// Floating-point Convert to Signed integer, rounding to nearest with ties to even (vector). This instruction converts a scalar or each element in a vector from a floating-point value to a signed integer value using the Round to Nearest rounding mode, and writes the result to the SIMD&FP destination register.
func VcvtnS32F32(v0 Float32X2) Int32X2 {
	return C.vcvtn_s32_f32(v0)
}

// Floating-point Convert to Unsigned integer, rounding to nearest with ties to even (vector). This instruction converts a scalar or each element in a vector from a floating-point value to an unsigned integer value using the Round to Nearest rounding mode, and writes the result to the SIMD&FP destination register.
func VcvtnqU32F32(v0 Float32X4) Uint32X4 {
	return C.vcvtnq_u32_f32(v0)
}

// Floating-point Convert to Unsigned integer, rounding to nearest with ties to even (vector). This instruction converts a scalar or each element in a vector from a floating-point value to an unsigned integer value using the Round to Nearest rounding mode, and writes the result to the SIMD&FP destination register.
func VcvtnU32F32(v0 Float32X2) Uint32X2 {
	return C.vcvtn_u32_f32(v0)
}

// Floating-point Convert to Signed integer, rounding toward Plus infinity (vector). This instruction converts a scalar or each element in a vector from a floating-point value to a signed integer value using the Round towards Plus Infinity rounding mode, and writes the result to the SIMD&FP destination register.
func VcvtpqS32F32(v0 Float32X4) Int32X4 {
	return C.vcvtpq_s32_f32(v0)
}

// Floating-point Convert to Signed integer, rounding toward Plus infinity (vector). This instruction converts a scalar or each element in a vector from a floating-point value to a signed integer value using the Round towards Plus Infinity rounding mode, and writes the result to the SIMD&FP destination register.
func VcvtpS32F32(v0 Float32X2) Int32X2 {
	return C.vcvtp_s32_f32(v0)
}

// Floating-point Convert to Unsigned integer, rounding toward Plus infinity (vector). This instruction converts a scalar or each element in a vector from a floating-point value to an unsigned integer value using the Round towards Plus Infinity rounding mode, and writes the result to the SIMD&FP destination register.
func VcvtpqU32F32(v0 Float32X4) Uint32X4 {
	return C.vcvtpq_u32_f32(v0)
}

// Floating-point Convert to Unsigned integer, rounding toward Plus infinity (vector). This instruction converts a scalar or each element in a vector from a floating-point value to an unsigned integer value using the Round towards Plus Infinity rounding mode, and writes the result to the SIMD&FP destination register.
func VcvtpU32F32(v0 Float32X2) Uint32X2 {
	return C.vcvtp_u32_f32(v0)
}

// AES single round decryption.
func VaesdqU8(v0 Uint8X16, v1 Uint8X16) Uint8X16 {
	return C.vaesdq_u8(v0, v1)
}

// AES single round encryption.
func VaeseqU8(v0 Uint8X16, v1 Uint8X16) Uint8X16 {
	return C.vaeseq_u8(v0, v1)
}

// AES inverse mix columns.
func VaesimcqU8(v0 Uint8X16) Uint8X16 {
	return C.vaesimcq_u8(v0)
}

// AES mix columns.
func VaesmcqU8(v0 Uint8X16) Uint8X16 {
	return C.vaesmcq_u8(v0)
}

// Floating-point Round to Integral, toward Zero (vector). This instruction rounds a vector of floating-point values in the SIMD&FP source register to integral floating-point values of the same size using the Round towards Zero rounding mode, and writes the result to the SIMD&FP destination register.
func VrndqF32(v0 Float32X4) Float32X4 {
	return C.vrndq_f32(v0)
}

// Floating-point Round to Integral, toward Zero (vector). This instruction rounds a vector of floating-point values in the SIMD&FP source register to integral floating-point values of the same size using the Round towards Zero rounding mode, and writes the result to the SIMD&FP destination register.
func VrndF32(v0 Float32X2) Float32X2 {
	return C.vrnd_f32(v0)
}

// Floating-point Round to Integral, to nearest with ties to Away (vector). This instruction rounds a vector of floating-point values in the SIMD&FP source register to integral floating-point values of the same size using the Round to Nearest with Ties to Away rounding mode, and writes the result to the SIMD&FP destination register.
func VrndaqF32(v0 Float32X4) Float32X4 {
	return C.vrndaq_f32(v0)
}

// Floating-point Round to Integral, to nearest with ties to Away (vector). This instruction rounds a vector of floating-point values in the SIMD&FP source register to integral floating-point values of the same size using the Round to Nearest with Ties to Away rounding mode, and writes the result to the SIMD&FP destination register.
func VrndaF32(v0 Float32X2) Float32X2 {
	return C.vrnda_f32(v0)
}

// Floating-point Round to Integral, using current rounding mode (vector). This instruction rounds a vector of floating-point values in the SIMD&FP source register to integral floating-point values of the same size using the rounding mode that is determined by the FPCR, and writes the result to the SIMD&FP destination register.
func VrndiqF32(v0 Float32X4) Float32X4 {
	return C.vrndiq_f32(v0)
}

// Floating-point Round to Integral, using current rounding mode (vector). This instruction rounds a vector of floating-point values in the SIMD&FP source register to integral floating-point values of the same size using the rounding mode that is determined by the FPCR, and writes the result to the SIMD&FP destination register.
func VrndiF32(v0 Float32X2) Float32X2 {
	return C.vrndi_f32(v0)
}

// Floating-point Round to Integral, toward Minus infinity (vector). This instruction rounds a vector of floating-point values in the SIMD&FP source register to integral floating-point values of the same size using the Round towards Minus Infinity rounding mode, and writes the result to the SIMD&FP destination register.
func VrndmqF32(v0 Float32X4) Float32X4 {
	return C.vrndmq_f32(v0)
}

// Floating-point Round to Integral, toward Minus infinity (vector). This instruction rounds a vector of floating-point values in the SIMD&FP source register to integral floating-point values of the same size using the Round towards Minus Infinity rounding mode, and writes the result to the SIMD&FP destination register.
func VrndmF32(v0 Float32X2) Float32X2 {
	return C.vrndm_f32(v0)
}

// Floating-point Round to Integral, to nearest with ties to even (vector). This instruction rounds a vector of floating-point values in the SIMD&FP source register to integral floating-point values of the same size using the Round to Nearest rounding mode, and writes the result to the SIMD&FP destination register.
func VrndnqF32(v0 Float32X4) Float32X4 {
	return C.vrndnq_f32(v0)
}

// Floating-point Round to Integral, to nearest with ties to even (vector). This instruction rounds a vector of floating-point values in the SIMD&FP source register to integral floating-point values of the same size using the Round to Nearest rounding mode, and writes the result to the SIMD&FP destination register.
func VrndnF32(v0 Float32X2) Float32X2 {
	return C.vrndn_f32(v0)
}

// Floating-point Round to Integral, to nearest with ties to even (vector). This instruction rounds a vector of floating-point values in the SIMD&FP source register to integral floating-point values of the same size using the Round to Nearest rounding mode, and writes the result to the SIMD&FP destination register.
func VrndnsF32(v0 Float32) Float32 {
	return C.vrndns_f32(v0)
}

// Floating-point Round to Integral, toward Plus infinity (vector). This instruction rounds a vector of floating-point values in the SIMD&FP source register to integral floating-point values of the same size using the Round towards Plus Infinity rounding mode, and writes the result to the SIMD&FP destination register.
func VrndpqF32(v0 Float32X4) Float32X4 {
	return C.vrndpq_f32(v0)
}

// Floating-point Round to Integral, toward Plus infinity (vector). This instruction rounds a vector of floating-point values in the SIMD&FP source register to integral floating-point values of the same size using the Round towards Plus Infinity rounding mode, and writes the result to the SIMD&FP destination register.
func VrndpF32(v0 Float32X2) Float32X2 {
	return C.vrndp_f32(v0)
}

// Floating-point Round to Integral exact, using current rounding mode (vector). This instruction rounds a vector of floating-point values in the SIMD&FP source register to integral floating-point values of the same size using the rounding mode that is determined by the FPCR, and writes the result to the SIMD&FP destination register.
func VrndxqF32(v0 Float32X4) Float32X4 {
	return C.vrndxq_f32(v0)
}

// Floating-point Round to Integral exact, using current rounding mode (vector). This instruction rounds a vector of floating-point values in the SIMD&FP source register to integral floating-point values of the same size using the rounding mode that is determined by the FPCR, and writes the result to the SIMD&FP destination register.
func VrndxF32(v0 Float32X2) Float32X2 {
	return C.vrndx_f32(v0)
}

// Floating-point Maximum Number (vector). This instruction compares corresponding vector elements in the two source SIMD&FP registers, writes the larger of the two floating-point values into a vector, and writes the vector to the destination SIMD&FP register.
func VmaxnmqF32(v0 Float32X4, v1 Float32X4) Float32X4 {
	return C.vmaxnmq_f32(v0, v1)
}

// Floating-point Maximum Number (vector). This instruction compares corresponding vector elements in the two source SIMD&FP registers, writes the larger of the two floating-point values into a vector, and writes the vector to the destination SIMD&FP register.
func VmaxnmF32(v0 Float32X2, v1 Float32X2) Float32X2 {
	return C.vmaxnm_f32(v0, v1)
}

// Floating-point Minimum Number (vector). This instruction compares corresponding vector elements in the two source SIMD&FP registers, writes the smaller of the two floating-point values into a vector, and writes the vector to the destination SIMD&FP register.
func VminnmqF32(v0 Float32X4, v1 Float32X4) Float32X4 {
	return C.vminnmq_f32(v0, v1)
}

// Floating-point Minimum Number (vector). This instruction compares corresponding vector elements in the two source SIMD&FP registers, writes the smaller of the two floating-point values into a vector, and writes the vector to the destination SIMD&FP register.
func VminnmF32(v0 Float32X2, v1 Float32X2) Float32X2 {
	return C.vminnm_f32(v0, v1)
}

// SHA1 hash update (choose).
func Vsha1CqU32(v0 Uint32X4, v1 Uint32, v2 Uint32X4) Uint32X4 {
	return C.vsha1cq_u32(v0, v1, v2)
}

// SHA1 fixed rotate.
func Vsha1HU32(v0 Uint32) Uint32 {
	return C.vsha1h_u32(v0)
}

// SHA1 hash update (majority).
func Vsha1MqU32(v0 Uint32X4, v1 Uint32, v2 Uint32X4) Uint32X4 {
	return C.vsha1mq_u32(v0, v1, v2)
}

// SHA1 hash update (parity).
func Vsha1PqU32(v0 Uint32X4, v1 Uint32, v2 Uint32X4) Uint32X4 {
	return C.vsha1pq_u32(v0, v1, v2)
}

// SHA1 schedule update 0.
func Vsha1Su0QU32(v0 Uint32X4, v1 Uint32X4, v2 Uint32X4) Uint32X4 {
	return C.vsha1su0q_u32(v0, v1, v2)
}

// SHA1 schedule update 1.
func Vsha1Su1QU32(v0 Uint32X4, v1 Uint32X4) Uint32X4 {
	return C.vsha1su1q_u32(v0, v1)
}

// SHA256 hash update (part 1).
func Vsha256HqU32(v0 Uint32X4, v1 Uint32X4, v2 Uint32X4) Uint32X4 {
	return C.vsha256hq_u32(v0, v1, v2)
}

// SHA256 hash update (part 2).
func Vsha256H2QU32(v0 Uint32X4, v1 Uint32X4, v2 Uint32X4) Uint32X4 {
	return C.vsha256h2q_u32(v0, v1, v2)
}

// SHA256 schedule update 0.
func Vsha256Su0QU32(v0 Uint32X4, v1 Uint32X4) Uint32X4 {
	return C.vsha256su0q_u32(v0, v1)
}

// SHA256 schedule update 1.
func Vsha256Su1QU32(v0 Uint32X4, v1 Uint32X4, v2 Uint32X4) Uint32X4 {
	return C.vsha256su1q_u32(v0, v1, v2)
}

// Bit Clear and Exclusive OR performs a bitwise AND of the 128-bit vector in a source SIMD&FP register and the complement of the vector in another source SIMD&FP register, then performs a bitwise exclusive OR of the resulting vector and the vector in a third source SIMD&FP register, and writes the result to the destination SIMD&FP register.
func VbcaxqU8(v0 Uint8X16, v1 Uint8X16, v2 Uint8X16) Uint8X16 {
	return C.vbcaxq_u8(v0, v1, v2)
}

// Bit Clear and Exclusive OR performs a bitwise AND of the 128-bit vector in a source SIMD&FP register and the complement of the vector in another source SIMD&FP register, then performs a bitwise exclusive OR of the resulting vector and the vector in a third source SIMD&FP register, and writes the result to the destination SIMD&FP register.
func VbcaxqU32(v0 Uint32X4, v1 Uint32X4, v2 Uint32X4) Uint32X4 {
	return C.vbcaxq_u32(v0, v1, v2)
}

// Bit Clear and Exclusive OR performs a bitwise AND of the 128-bit vector in a source SIMD&FP register and the complement of the vector in another source SIMD&FP register, then performs a bitwise exclusive OR of the resulting vector and the vector in a third source SIMD&FP register, and writes the result to the destination SIMD&FP register.
func VbcaxqU64(v0 Uint64X2, v1 Uint64X2, v2 Uint64X2) Uint64X2 {
	return C.vbcaxq_u64(v0, v1, v2)
}

// Bit Clear and Exclusive OR performs a bitwise AND of the 128-bit vector in a source SIMD&FP register and the complement of the vector in another source SIMD&FP register, then performs a bitwise exclusive OR of the resulting vector and the vector in a third source SIMD&FP register, and writes the result to the destination SIMD&FP register.
func VbcaxqU16(v0 Uint16X8, v1 Uint16X8, v2 Uint16X8) Uint16X8 {
	return C.vbcaxq_u16(v0, v1, v2)
}

// Bit Clear and Exclusive OR performs a bitwise AND of the 128-bit vector in a source SIMD&FP register and the complement of the vector in another source SIMD&FP register, then performs a bitwise exclusive OR of the resulting vector and the vector in a third source SIMD&FP register, and writes the result to the destination SIMD&FP register.
func VbcaxqS8(v0 Int8X16, v1 Int8X16, v2 Int8X16) Int8X16 {
	return C.vbcaxq_s8(v0, v1, v2)
}

// Bit Clear and Exclusive OR performs a bitwise AND of the 128-bit vector in a source SIMD&FP register and the complement of the vector in another source SIMD&FP register, then performs a bitwise exclusive OR of the resulting vector and the vector in a third source SIMD&FP register, and writes the result to the destination SIMD&FP register.
func VbcaxqS32(v0 Int32X4, v1 Int32X4, v2 Int32X4) Int32X4 {
	return C.vbcaxq_s32(v0, v1, v2)
}

// Bit Clear and Exclusive OR performs a bitwise AND of the 128-bit vector in a source SIMD&FP register and the complement of the vector in another source SIMD&FP register, then performs a bitwise exclusive OR of the resulting vector and the vector in a third source SIMD&FP register, and writes the result to the destination SIMD&FP register.
func VbcaxqS64(v0 Int64X2, v1 Int64X2, v2 Int64X2) Int64X2 {
	return C.vbcaxq_s64(v0, v1, v2)
}

// Bit Clear and Exclusive OR performs a bitwise AND of the 128-bit vector in a source SIMD&FP register and the complement of the vector in another source SIMD&FP register, then performs a bitwise exclusive OR of the resulting vector and the vector in a third source SIMD&FP register, and writes the result to the destination SIMD&FP register.
func VbcaxqS16(v0 Int16X8, v1 Int16X8, v2 Int16X8) Int16X8 {
	return C.vbcaxq_s16(v0, v1, v2)
}

// Three-way Exclusive OR performs a three-way exclusive OR of the values in the three source SIMD&FP registers, and writes the result to the destination SIMD&FP register.
func Veor3QU8(v0 Uint8X16, v1 Uint8X16, v2 Uint8X16) Uint8X16 {
	return C.veor3q_u8(v0, v1, v2)
}

// Three-way Exclusive OR performs a three-way exclusive OR of the values in the three source SIMD&FP registers, and writes the result to the destination SIMD&FP register.
func Veor3QU32(v0 Uint32X4, v1 Uint32X4, v2 Uint32X4) Uint32X4 {
	return C.veor3q_u32(v0, v1, v2)
}

// Three-way Exclusive OR performs a three-way exclusive OR of the values in the three source SIMD&FP registers, and writes the result to the destination SIMD&FP register.
func Veor3QU64(v0 Uint64X2, v1 Uint64X2, v2 Uint64X2) Uint64X2 {
	return C.veor3q_u64(v0, v1, v2)
}

// Three-way Exclusive OR performs a three-way exclusive OR of the values in the three source SIMD&FP registers, and writes the result to the destination SIMD&FP register.
func Veor3QU16(v0 Uint16X8, v1 Uint16X8, v2 Uint16X8) Uint16X8 {
	return C.veor3q_u16(v0, v1, v2)
}

// Three-way Exclusive OR performs a three-way exclusive OR of the values in the three source SIMD&FP registers, and writes the result to the destination SIMD&FP register.
func Veor3QS8(v0 Int8X16, v1 Int8X16, v2 Int8X16) Int8X16 {
	return C.veor3q_s8(v0, v1, v2)
}

// Three-way Exclusive OR performs a three-way exclusive OR of the values in the three source SIMD&FP registers, and writes the result to the destination SIMD&FP register.
func Veor3QS32(v0 Int32X4, v1 Int32X4, v2 Int32X4) Int32X4 {
	return C.veor3q_s32(v0, v1, v2)
}

// Three-way Exclusive OR performs a three-way exclusive OR of the values in the three source SIMD&FP registers, and writes the result to the destination SIMD&FP register.
func Veor3QS64(v0 Int64X2, v1 Int64X2, v2 Int64X2) Int64X2 {
	return C.veor3q_s64(v0, v1, v2)
}

// Three-way Exclusive OR performs a three-way exclusive OR of the values in the three source SIMD&FP registers, and writes the result to the destination SIMD&FP register.
func Veor3QS16(v0 Int16X8, v1 Int16X8, v2 Int16X8) Int16X8 {
	return C.veor3q_s16(v0, v1, v2)
}

// Rotate and Exclusive OR rotates each 64-bit element of the 128-bit vector in a source SIMD&FP register left by 1, performs a bitwise exclusive OR of the resulting 128-bit vector and the vector in another source SIMD&FP register, and writes the result to the destination SIMD&FP register.
func Vrax1QU64(v0 Uint64X2, v1 Uint64X2) Uint64X2 {
	return C.vrax1q_u64(v0, v1)
}

// SHA512 Hash update part 1 takes the values from the three 128-bit source SIMD&FP registers and produces a 128-bit output value that combines the sigma1 and chi functions of two iterations of the SHA512 computation. It returns this value to the destination SIMD&FP register.
func Vsha512HqU64(v0 Uint64X2, v1 Uint64X2, v2 Uint64X2) Uint64X2 {
	return C.vsha512hq_u64(v0, v1, v2)
}

// SHA512 Hash update part 2 takes the values from the three 128-bit source SIMD&FP registers and produces a 128-bit output value that combines the sigma0 and majority functions of two iterations of the SHA512 computation. It returns this value to the destination SIMD&FP register.
func Vsha512H2QU64(v0 Uint64X2, v1 Uint64X2, v2 Uint64X2) Uint64X2 {
	return C.vsha512h2q_u64(v0, v1, v2)
}

// SHA512 Schedule Update 0 takes the values from the two 128-bit source SIMD&FP registers and produces a 128-bit output value that combines the gamma0 functions of two iterations of the SHA512 schedule update that are performed after the first 16 iterations within a block. It returns this value to the destination SIMD&FP register.
func Vsha512Su0QU64(v0 Uint64X2, v1 Uint64X2) Uint64X2 {
	return C.vsha512su0q_u64(v0, v1)
}

// SHA512 Schedule Update 1 takes the values from the three source SIMD&FP registers and produces a 128-bit output value that combines the gamma1 functions of two iterations of the SHA512 schedule update that are performed after the first 16 iterations within a block. It returns this value to the destination SIMD&FP register.
func Vsha512Su1QU64(v0 Uint64X2, v1 Uint64X2, v2 Uint64X2) Uint64X2 {
	return C.vsha512su1q_u64(v0, v1, v2)
}

// SM3PARTW1 takes three 128-bit vectors from the three source SIMD&FP registers and returns a 128-bit result in the destination SIMD&FP register. The result is obtained by a three-way exclusive OR of the elements within the input vectors with some fixed rotations, see the Operation pseudocode for more information.
func Vsm3Partw1QU32(v0 Uint32X4, v1 Uint32X4, v2 Uint32X4) Uint32X4 {
	return C.vsm3partw1q_u32(v0, v1, v2)
}

// SM3PARTW2 takes three 128-bit vectors from three source SIMD&FP registers and returns a 128-bit result in the destination SIMD&FP register. The result is obtained by a three-way exclusive OR of the elements within the input vectors with some fixed rotations, see the Operation pseudocode for more information.
func Vsm3Partw2QU32(v0 Uint32X4, v1 Uint32X4, v2 Uint32X4) Uint32X4 {
	return C.vsm3partw2q_u32(v0, v1, v2)
}

// SM3SS1 rotates the top 32 bits of the 128-bit vector in the first source SIMD&FP register by 12, and adds that 32-bit value to the two other 32-bit values held in the top 32 bits of each of the 128-bit vectors in the second and third source SIMD&FP registers, rotating this result left by 7 and writing the final result into the top 32 bits of the vector in the destination SIMD&FP register, with the bottom 96 bits of the vector being written to 0.
func Vsm3Ss1QU32(v0 Uint32X4, v1 Uint32X4, v2 Uint32X4) Uint32X4 {
	return C.vsm3ss1q_u32(v0, v1, v2)
}

// SM4 Encode takes input data as a 128-bit vector from the first source SIMD&FP register, and four iterations of the round key held as the elements of the 128-bit vector in the second source SIMD&FP register. It encrypts the data by four rounds, in accordance with the SM4 standard, returning the 128-bit result to the destination SIMD&FP register.
func Vsm4EqU32(v0 Uint32X4, v1 Uint32X4) Uint32X4 {
	return C.vsm4eq_u32(v0, v1)
}

// SM4 Key takes an input as a 128-bit vector from the first source SIMD&FP register and a 128-bit constant from the second SIMD&FP register. It derives four iterations of the output key, in accordance with the SM4 standard, returning the 128-bit result to the destination SIMD&FP register.
func Vsm4EkeyqU32(v0 Uint32X4, v1 Uint32X4) Uint32X4 {
	return C.vsm4ekeyq_u32(v0, v1)
}

// Floating-point Convert to Signed integer, rounding to nearest with ties to Away (vector). This instruction converts each element in a vector from a floating-point value to a signed integer value using the Round to Nearest with Ties to Away rounding mode and writes the result to the SIMD&FP destination register.
func VcvtaqS64F64(v0 Float64X2) Int64X2 {
	return C.vcvtaq_s64_f64(v0)
}

// Floating-point Convert to Signed integer, rounding to nearest with ties to Away (vector). This instruction converts each element in a vector from a floating-point value to a signed integer value using the Round to Nearest with Ties to Away rounding mode and writes the result to the SIMD&FP destination register.
func VcvtaS64F64(v0 Float64X1) Int64X1 {
	return C.vcvta_s64_f64(v0)
}

// Floating-point Convert to Unsigned integer, rounding to nearest with ties to Away (vector). This instruction converts each element in a vector from a floating-point value to an unsigned integer value using the Round to Nearest with Ties to Away rounding mode and writes the result to the SIMD&FP destination register.
func VcvtaqU64F64(v0 Float64X2) Uint64X2 {
	return C.vcvtaq_u64_f64(v0)
}

// Floating-point Convert to Unsigned integer, rounding to nearest with ties to Away (vector). This instruction converts each element in a vector from a floating-point value to an unsigned integer value using the Round to Nearest with Ties to Away rounding mode and writes the result to the SIMD&FP destination register.
func VcvtaU64F64(v0 Float64X1) Uint64X1 {
	return C.vcvta_u64_f64(v0)
}

// Floating-point Convert to Signed integer, rounding toward Minus infinity (vector). This instruction converts a scalar or each element in a vector from a floating-point value to a signed integer value using the Round towards Minus Infinity rounding mode, and writes the result to the SIMD&FP destination register.
func VcvtmqS64F64(v0 Float64X2) Int64X2 {
	return C.vcvtmq_s64_f64(v0)
}

// Floating-point Convert to Signed integer, rounding toward Minus infinity (vector). This instruction converts a scalar or each element in a vector from a floating-point value to a signed integer value using the Round towards Minus Infinity rounding mode, and writes the result to the SIMD&FP destination register.
func VcvtmS64F64(v0 Float64X1) Int64X1 {
	return C.vcvtm_s64_f64(v0)
}

// Floating-point Convert to Unsigned integer, rounding toward Minus infinity (vector). This instruction converts a scalar or each element in a vector from a floating-point value to an unsigned integer value using the Round towards Minus Infinity rounding mode, and writes the result to the SIMD&FP destination register.
func VcvtmqU64F64(v0 Float64X2) Uint64X2 {
	return C.vcvtmq_u64_f64(v0)
}

// Floating-point Convert to Unsigned integer, rounding toward Minus infinity (vector). This instruction converts a scalar or each element in a vector from a floating-point value to an unsigned integer value using the Round towards Minus Infinity rounding mode, and writes the result to the SIMD&FP destination register.
func VcvtmU64F64(v0 Float64X1) Uint64X1 {
	return C.vcvtm_u64_f64(v0)
}

// Floating-point Convert to Signed integer, rounding to nearest with ties to even (vector). This instruction converts a scalar or each element in a vector from a floating-point value to a signed integer value using the Round to Nearest rounding mode, and writes the result to the SIMD&FP destination register.
func VcvtnqS64F64(v0 Float64X2) Int64X2 {
	return C.vcvtnq_s64_f64(v0)
}

// Floating-point Convert to Signed integer, rounding to nearest with ties to even (vector). This instruction converts a scalar or each element in a vector from a floating-point value to a signed integer value using the Round to Nearest rounding mode, and writes the result to the SIMD&FP destination register.
func VcvtnS64F64(v0 Float64X1) Int64X1 {
	return C.vcvtn_s64_f64(v0)
}

// Floating-point Convert to Unsigned integer, rounding to nearest with ties to even (vector). This instruction converts a scalar or each element in a vector from a floating-point value to an unsigned integer value using the Round to Nearest rounding mode, and writes the result to the SIMD&FP destination register.
func VcvtnqU64F64(v0 Float64X2) Uint64X2 {
	return C.vcvtnq_u64_f64(v0)
}

// Floating-point Convert to Unsigned integer, rounding to nearest with ties to even (vector). This instruction converts a scalar or each element in a vector from a floating-point value to an unsigned integer value using the Round to Nearest rounding mode, and writes the result to the SIMD&FP destination register.
func VcvtnU64F64(v0 Float64X1) Uint64X1 {
	return C.vcvtn_u64_f64(v0)
}

// Floating-point Convert to Signed integer, rounding toward Plus infinity (vector). This instruction converts a scalar or each element in a vector from a floating-point value to a signed integer value using the Round towards Plus Infinity rounding mode, and writes the result to the SIMD&FP destination register.
func VcvtpqS64F64(v0 Float64X2) Int64X2 {
	return C.vcvtpq_s64_f64(v0)
}

// Floating-point Convert to Signed integer, rounding toward Plus infinity (vector). This instruction converts a scalar or each element in a vector from a floating-point value to a signed integer value using the Round towards Plus Infinity rounding mode, and writes the result to the SIMD&FP destination register.
func VcvtpS64F64(v0 Float64X1) Int64X1 {
	return C.vcvtp_s64_f64(v0)
}

// Floating-point Convert to Unsigned integer, rounding toward Plus infinity (vector). This instruction converts a scalar or each element in a vector from a floating-point value to an unsigned integer value using the Round towards Plus Infinity rounding mode, and writes the result to the SIMD&FP destination register.
func VcvtpqU64F64(v0 Float64X2) Uint64X2 {
	return C.vcvtpq_u64_f64(v0)
}

// Floating-point Convert to Unsigned integer, rounding toward Plus infinity (vector). This instruction converts a scalar or each element in a vector from a floating-point value to an unsigned integer value using the Round towards Plus Infinity rounding mode, and writes the result to the SIMD&FP destination register.
func VcvtpU64F64(v0 Float64X1) Uint64X1 {
	return C.vcvtp_u64_f64(v0)
}

// Vector reinterpret cast operation
func VreinterpretP8P64(v0 Poly64X1) Poly8X8 {
	return C.vreinterpret_p8_p64(v0)
}

// Vector reinterpret cast operation
func VreinterpretP8P16(v0 Poly16X4) Poly8X8 {
	return C.vreinterpret_p8_p16(v0)
}

// Vector reinterpret cast operation
func VreinterpretP8U8(v0 Uint8X8) Poly8X8 {
	return C.vreinterpret_p8_u8(v0)
}

// Vector reinterpret cast operation
func VreinterpretP8U32(v0 Uint32X2) Poly8X8 {
	return C.vreinterpret_p8_u32(v0)
}

// Vector reinterpret cast operation
func VreinterpretP8U64(v0 Uint64X1) Poly8X8 {
	return C.vreinterpret_p8_u64(v0)
}

// Vector reinterpret cast operation
func VreinterpretP8U16(v0 Uint16X4) Poly8X8 {
	return C.vreinterpret_p8_u16(v0)
}

// Vector reinterpret cast operation
func VreinterpretP8S8(v0 Int8X8) Poly8X8 {
	return C.vreinterpret_p8_s8(v0)
}

// Vector reinterpret cast operation
func VreinterpretP8F64(v0 Float64X1) Poly8X8 {
	return C.vreinterpret_p8_f64(v0)
}

// Vector reinterpret cast operation
func VreinterpretP8F32(v0 Float32X2) Poly8X8 {
	return C.vreinterpret_p8_f32(v0)
}

// Vector reinterpret cast operation
func VreinterpretP8S32(v0 Int32X2) Poly8X8 {
	return C.vreinterpret_p8_s32(v0)
}

// Vector reinterpret cast operation
func VreinterpretP8S64(v0 Int64X1) Poly8X8 {
	return C.vreinterpret_p8_s64(v0)
}

// Vector reinterpret cast operation
func VreinterpretP8S16(v0 Int16X4) Poly8X8 {
	return C.vreinterpret_p8_s16(v0)
}

// Vector reinterpret cast operation
func VreinterpretP64P8(v0 Poly8X8) Poly64X1 {
	return C.vreinterpret_p64_p8(v0)
}

// Vector reinterpret cast operation
func VreinterpretP64P16(v0 Poly16X4) Poly64X1 {
	return C.vreinterpret_p64_p16(v0)
}

// Vector reinterpret cast operation
func VreinterpretP64U8(v0 Uint8X8) Poly64X1 {
	return C.vreinterpret_p64_u8(v0)
}

// Vector reinterpret cast operation
func VreinterpretP64U32(v0 Uint32X2) Poly64X1 {
	return C.vreinterpret_p64_u32(v0)
}

// Vector reinterpret cast operation
func VreinterpretP64U64(v0 Uint64X1) Poly64X1 {
	return C.vreinterpret_p64_u64(v0)
}

// Vector reinterpret cast operation
func VreinterpretP64U16(v0 Uint16X4) Poly64X1 {
	return C.vreinterpret_p64_u16(v0)
}

// Vector reinterpret cast operation
func VreinterpretP64S8(v0 Int8X8) Poly64X1 {
	return C.vreinterpret_p64_s8(v0)
}

// Vector reinterpret cast operation
func VreinterpretP64F64(v0 Float64X1) Poly64X1 {
	return C.vreinterpret_p64_f64(v0)
}

// Vector reinterpret cast operation
func VreinterpretP64F32(v0 Float32X2) Poly64X1 {
	return C.vreinterpret_p64_f32(v0)
}

// Vector reinterpret cast operation
func VreinterpretP64S32(v0 Int32X2) Poly64X1 {
	return C.vreinterpret_p64_s32(v0)
}

// vreinterpret_p64_s64
func VreinterpretP64S64(v0 Int64X1) Poly64X1 {
	return C.vreinterpret_p64_s64(v0)
}

// Vector reinterpret cast operation
func VreinterpretP64S16(v0 Int16X4) Poly64X1 {
	return C.vreinterpret_p64_s16(v0)
}

// Vector reinterpret cast operation
func VreinterpretP16P8(v0 Poly8X8) Poly16X4 {
	return C.vreinterpret_p16_p8(v0)
}

// Vector reinterpret cast operation
func VreinterpretP16P64(v0 Poly64X1) Poly16X4 {
	return C.vreinterpret_p16_p64(v0)
}

// Vector reinterpret cast operation
func VreinterpretP16U8(v0 Uint8X8) Poly16X4 {
	return C.vreinterpret_p16_u8(v0)
}

// Vector reinterpret cast operation
func VreinterpretP16U32(v0 Uint32X2) Poly16X4 {
	return C.vreinterpret_p16_u32(v0)
}

// Vector reinterpret cast operation
func VreinterpretP16U64(v0 Uint64X1) Poly16X4 {
	return C.vreinterpret_p16_u64(v0)
}

// Vector reinterpret cast operation
func VreinterpretP16U16(v0 Uint16X4) Poly16X4 {
	return C.vreinterpret_p16_u16(v0)
}

// Vector reinterpret cast operation
func VreinterpretP16S8(v0 Int8X8) Poly16X4 {
	return C.vreinterpret_p16_s8(v0)
}

// Vector reinterpret cast operation
func VreinterpretP16F64(v0 Float64X1) Poly16X4 {
	return C.vreinterpret_p16_f64(v0)
}

// Vector reinterpret cast operation
func VreinterpretP16F32(v0 Float32X2) Poly16X4 {
	return C.vreinterpret_p16_f32(v0)
}

// Vector reinterpret cast operation
func VreinterpretP16S32(v0 Int32X2) Poly16X4 {
	return C.vreinterpret_p16_s32(v0)
}

// Vector reinterpret cast operation
func VreinterpretP16S64(v0 Int64X1) Poly16X4 {
	return C.vreinterpret_p16_s64(v0)
}

// Vector reinterpret cast operation
func VreinterpretP16S16(v0 Int16X4) Poly16X4 {
	return C.vreinterpret_p16_s16(v0)
}

// Vector reinterpret cast operation
func VreinterpretqP8P128(v0 Poly128) Poly8X16 {
	return C.vreinterpretq_p8_p128(v0)
}

// Vector reinterpret cast operation
func VreinterpretqP8P64(v0 Poly64X2) Poly8X16 {
	return C.vreinterpretq_p8_p64(v0)
}

// Vector reinterpret cast operation
func VreinterpretqP8P16(v0 Poly16X8) Poly8X16 {
	return C.vreinterpretq_p8_p16(v0)
}

// Vector reinterpret cast operation
func VreinterpretqP8U8(v0 Uint8X16) Poly8X16 {
	return C.vreinterpretq_p8_u8(v0)
}

// Vector reinterpret cast operation
func VreinterpretqP8U32(v0 Uint32X4) Poly8X16 {
	return C.vreinterpretq_p8_u32(v0)
}

// Vector reinterpret cast operation
func VreinterpretqP8U64(v0 Uint64X2) Poly8X16 {
	return C.vreinterpretq_p8_u64(v0)
}

// Vector reinterpret cast operation
func VreinterpretqP8U16(v0 Uint16X8) Poly8X16 {
	return C.vreinterpretq_p8_u16(v0)
}

// Vector reinterpret cast operation
func VreinterpretqP8S8(v0 Int8X16) Poly8X16 {
	return C.vreinterpretq_p8_s8(v0)
}

// Vector reinterpret cast operation
func VreinterpretqP8F64(v0 Float64X2) Poly8X16 {
	return C.vreinterpretq_p8_f64(v0)
}

// Vector reinterpret cast operation
func VreinterpretqP8F32(v0 Float32X4) Poly8X16 {
	return C.vreinterpretq_p8_f32(v0)
}

// Vector reinterpret cast operation
func VreinterpretqP8S32(v0 Int32X4) Poly8X16 {
	return C.vreinterpretq_p8_s32(v0)
}

// Vector reinterpret cast operation
func VreinterpretqP8S64(v0 Int64X2) Poly8X16 {
	return C.vreinterpretq_p8_s64(v0)
}

// Vector reinterpret cast operation
func VreinterpretqP8S16(v0 Int16X8) Poly8X16 {
	return C.vreinterpretq_p8_s16(v0)
}

// Vector reinterpret cast operation
func VreinterpretqP128P8(v0 Poly8X16) Poly128 {
	return C.vreinterpretq_p128_p8(v0)
}

// vreinterpretq_p128_p64
func VreinterpretqP128P64(v0 Poly64X2) Poly128 {
	return C.vreinterpretq_p128_p64(v0)
}

// Vector reinterpret cast operation
func VreinterpretqP128P16(v0 Poly16X8) Poly128 {
	return C.vreinterpretq_p128_p16(v0)
}

// Vector reinterpret cast operation
func VreinterpretqP128U8(v0 Uint8X16) Poly128 {
	return C.vreinterpretq_p128_u8(v0)
}

// Vector reinterpret cast operation
func VreinterpretqP128U32(v0 Uint32X4) Poly128 {
	return C.vreinterpretq_p128_u32(v0)
}

// Vector reinterpret cast operation
func VreinterpretqP128U64(v0 Uint64X2) Poly128 {
	return C.vreinterpretq_p128_u64(v0)
}

// Vector reinterpret cast operation
func VreinterpretqP128U16(v0 Uint16X8) Poly128 {
	return C.vreinterpretq_p128_u16(v0)
}

// Vector reinterpret cast operation
func VreinterpretqP128S8(v0 Int8X16) Poly128 {
	return C.vreinterpretq_p128_s8(v0)
}

// Vector reinterpret cast operation
func VreinterpretqP128F64(v0 Float64X2) Poly128 {
	return C.vreinterpretq_p128_f64(v0)
}

// Vector reinterpret cast operation
func VreinterpretqP128F32(v0 Float32X4) Poly128 {
	return C.vreinterpretq_p128_f32(v0)
}

// Vector reinterpret cast operation
func VreinterpretqP128S32(v0 Int32X4) Poly128 {
	return C.vreinterpretq_p128_s32(v0)
}

// Vector reinterpret cast operation
func VreinterpretqP128S64(v0 Int64X2) Poly128 {
	return C.vreinterpretq_p128_s64(v0)
}

// Vector reinterpret cast operation
func VreinterpretqP128S16(v0 Int16X8) Poly128 {
	return C.vreinterpretq_p128_s16(v0)
}

// Vector reinterpret cast operation
func VreinterpretqP64P8(v0 Poly8X16) Poly64X2 {
	return C.vreinterpretq_p64_p8(v0)
}

// vreinterpretq_p64_p128
func VreinterpretqP64P128(v0 Poly128) Poly64X2 {
	return C.vreinterpretq_p64_p128(v0)
}

// Vector reinterpret cast operation
func VreinterpretqP64P16(v0 Poly16X8) Poly64X2 {
	return C.vreinterpretq_p64_p16(v0)
}

// Vector reinterpret cast operation
func VreinterpretqP64U8(v0 Uint8X16) Poly64X2 {
	return C.vreinterpretq_p64_u8(v0)
}

// Vector reinterpret cast operation
func VreinterpretqP64U32(v0 Uint32X4) Poly64X2 {
	return C.vreinterpretq_p64_u32(v0)
}

// Vector reinterpret cast operation
func VreinterpretqP64U64(v0 Uint64X2) Poly64X2 {
	return C.vreinterpretq_p64_u64(v0)
}

// Vector reinterpret cast operation
func VreinterpretqP64U16(v0 Uint16X8) Poly64X2 {
	return C.vreinterpretq_p64_u16(v0)
}

// Vector reinterpret cast operation
func VreinterpretqP64S8(v0 Int8X16) Poly64X2 {
	return C.vreinterpretq_p64_s8(v0)
}

// Vector reinterpret cast operation
func VreinterpretqP64F64(v0 Float64X2) Poly64X2 {
	return C.vreinterpretq_p64_f64(v0)
}

// Vector reinterpret cast operation
func VreinterpretqP64F32(v0 Float32X4) Poly64X2 {
	return C.vreinterpretq_p64_f32(v0)
}

// Vector reinterpret cast operation
func VreinterpretqP64S32(v0 Int32X4) Poly64X2 {
	return C.vreinterpretq_p64_s32(v0)
}

// Vector reinterpret cast operation
func VreinterpretqP64S64(v0 Int64X2) Poly64X2 {
	return C.vreinterpretq_p64_s64(v0)
}

// Vector reinterpret cast operation
func VreinterpretqP64S16(v0 Int16X8) Poly64X2 {
	return C.vreinterpretq_p64_s16(v0)
}

// Vector reinterpret cast operation
func VreinterpretqP16P8(v0 Poly8X16) Poly16X8 {
	return C.vreinterpretq_p16_p8(v0)
}

// Vector reinterpret cast operation
func VreinterpretqP16P128(v0 Poly128) Poly16X8 {
	return C.vreinterpretq_p16_p128(v0)
}

// Vector reinterpret cast operation
func VreinterpretqP16P64(v0 Poly64X2) Poly16X8 {
	return C.vreinterpretq_p16_p64(v0)
}

// Vector reinterpret cast operation
func VreinterpretqP16U8(v0 Uint8X16) Poly16X8 {
	return C.vreinterpretq_p16_u8(v0)
}

// Vector reinterpret cast operation
func VreinterpretqP16U32(v0 Uint32X4) Poly16X8 {
	return C.vreinterpretq_p16_u32(v0)
}

// Vector reinterpret cast operation
func VreinterpretqP16U64(v0 Uint64X2) Poly16X8 {
	return C.vreinterpretq_p16_u64(v0)
}

// Vector reinterpret cast operation
func VreinterpretqP16U16(v0 Uint16X8) Poly16X8 {
	return C.vreinterpretq_p16_u16(v0)
}

// Vector reinterpret cast operation
func VreinterpretqP16S8(v0 Int8X16) Poly16X8 {
	return C.vreinterpretq_p16_s8(v0)
}

// Vector reinterpret cast operation
func VreinterpretqP16F64(v0 Float64X2) Poly16X8 {
	return C.vreinterpretq_p16_f64(v0)
}

// Vector reinterpret cast operation
func VreinterpretqP16F32(v0 Float32X4) Poly16X8 {
	return C.vreinterpretq_p16_f32(v0)
}

// Vector reinterpret cast operation
func VreinterpretqP16S32(v0 Int32X4) Poly16X8 {
	return C.vreinterpretq_p16_s32(v0)
}

// Vector reinterpret cast operation
func VreinterpretqP16S64(v0 Int64X2) Poly16X8 {
	return C.vreinterpretq_p16_s64(v0)
}

// Vector reinterpret cast operation
func VreinterpretqP16S16(v0 Int16X8) Poly16X8 {
	return C.vreinterpretq_p16_s16(v0)
}

// Vector reinterpret cast operation
func VreinterpretqU8P8(v0 Poly8X16) Uint8X16 {
	return C.vreinterpretq_u8_p8(v0)
}

// Vector reinterpret cast operation
func VreinterpretqU8P128(v0 Poly128) Uint8X16 {
	return C.vreinterpretq_u8_p128(v0)
}

// Vector reinterpret cast operation
func VreinterpretqU8P64(v0 Poly64X2) Uint8X16 {
	return C.vreinterpretq_u8_p64(v0)
}

// Vector reinterpret cast operation
func VreinterpretqU8P16(v0 Poly16X8) Uint8X16 {
	return C.vreinterpretq_u8_p16(v0)
}

// Vector reinterpret cast operation
func VreinterpretqU8U32(v0 Uint32X4) Uint8X16 {
	return C.vreinterpretq_u8_u32(v0)
}

// Vector reinterpret cast operation
func VreinterpretqU8U64(v0 Uint64X2) Uint8X16 {
	return C.vreinterpretq_u8_u64(v0)
}

// Vector reinterpret cast operation
func VreinterpretqU8U16(v0 Uint16X8) Uint8X16 {
	return C.vreinterpretq_u8_u16(v0)
}

// Vector reinterpret cast operation
func VreinterpretqU8S8(v0 Int8X16) Uint8X16 {
	return C.vreinterpretq_u8_s8(v0)
}

// Vector reinterpret cast operation
func VreinterpretqU8F64(v0 Float64X2) Uint8X16 {
	return C.vreinterpretq_u8_f64(v0)
}

// Vector reinterpret cast operation
func VreinterpretqU8F32(v0 Float32X4) Uint8X16 {
	return C.vreinterpretq_u8_f32(v0)
}

// Vector reinterpret cast operation
func VreinterpretqU8S32(v0 Int32X4) Uint8X16 {
	return C.vreinterpretq_u8_s32(v0)
}

// Vector reinterpret cast operation
func VreinterpretqU8S64(v0 Int64X2) Uint8X16 {
	return C.vreinterpretq_u8_s64(v0)
}

// Vector reinterpret cast operation
func VreinterpretqU8S16(v0 Int16X8) Uint8X16 {
	return C.vreinterpretq_u8_s16(v0)
}

// Vector reinterpret cast operation
func VreinterpretqU32P8(v0 Poly8X16) Uint32X4 {
	return C.vreinterpretq_u32_p8(v0)
}

// Vector reinterpret cast operation
func VreinterpretqU32P128(v0 Poly128) Uint32X4 {
	return C.vreinterpretq_u32_p128(v0)
}

// Vector reinterpret cast operation
func VreinterpretqU32P64(v0 Poly64X2) Uint32X4 {
	return C.vreinterpretq_u32_p64(v0)
}

// Vector reinterpret cast operation
func VreinterpretqU32P16(v0 Poly16X8) Uint32X4 {
	return C.vreinterpretq_u32_p16(v0)
}

// Vector reinterpret cast operation
func VreinterpretqU32U8(v0 Uint8X16) Uint32X4 {
	return C.vreinterpretq_u32_u8(v0)
}

// Vector reinterpret cast operation
func VreinterpretqU32U64(v0 Uint64X2) Uint32X4 {
	return C.vreinterpretq_u32_u64(v0)
}

// Vector reinterpret cast operation
func VreinterpretqU32U16(v0 Uint16X8) Uint32X4 {
	return C.vreinterpretq_u32_u16(v0)
}

// Vector reinterpret cast operation
func VreinterpretqU32S8(v0 Int8X16) Uint32X4 {
	return C.vreinterpretq_u32_s8(v0)
}

// Vector reinterpret cast operation
func VreinterpretqU32F64(v0 Float64X2) Uint32X4 {
	return C.vreinterpretq_u32_f64(v0)
}

// Vector reinterpret cast operation
func VreinterpretqU32F32(v0 Float32X4) Uint32X4 {
	return C.vreinterpretq_u32_f32(v0)
}

// Vector reinterpret cast operation
func VreinterpretqU32S32(v0 Int32X4) Uint32X4 {
	return C.vreinterpretq_u32_s32(v0)
}

// Vector reinterpret cast operation
func VreinterpretqU32S64(v0 Int64X2) Uint32X4 {
	return C.vreinterpretq_u32_s64(v0)
}

// Vector reinterpret cast operation
func VreinterpretqU32S16(v0 Int16X8) Uint32X4 {
	return C.vreinterpretq_u32_s16(v0)
}

// Vector reinterpret cast operation
func VreinterpretqU64P8(v0 Poly8X16) Uint64X2 {
	return C.vreinterpretq_u64_p8(v0)
}

// Vector reinterpret cast operation
func VreinterpretqU64P128(v0 Poly128) Uint64X2 {
	return C.vreinterpretq_u64_p128(v0)
}

// Vector reinterpret cast operation
func VreinterpretqU64P64(v0 Poly64X2) Uint64X2 {
	return C.vreinterpretq_u64_p64(v0)
}

// Vector reinterpret cast operation
func VreinterpretqU64P16(v0 Poly16X8) Uint64X2 {
	return C.vreinterpretq_u64_p16(v0)
}

// Vector reinterpret cast operation
func VreinterpretqU64U8(v0 Uint8X16) Uint64X2 {
	return C.vreinterpretq_u64_u8(v0)
}

// Vector reinterpret cast operation
func VreinterpretqU64U32(v0 Uint32X4) Uint64X2 {
	return C.vreinterpretq_u64_u32(v0)
}

// Vector reinterpret cast operation
func VreinterpretqU64U16(v0 Uint16X8) Uint64X2 {
	return C.vreinterpretq_u64_u16(v0)
}

// Vector reinterpret cast operation
func VreinterpretqU64S8(v0 Int8X16) Uint64X2 {
	return C.vreinterpretq_u64_s8(v0)
}

// Vector reinterpret cast operation
func VreinterpretqU64F64(v0 Float64X2) Uint64X2 {
	return C.vreinterpretq_u64_f64(v0)
}

// Vector reinterpret cast operation
func VreinterpretqU64F32(v0 Float32X4) Uint64X2 {
	return C.vreinterpretq_u64_f32(v0)
}

// Vector reinterpret cast operation
func VreinterpretqU64S32(v0 Int32X4) Uint64X2 {
	return C.vreinterpretq_u64_s32(v0)
}

// Vector reinterpret cast operation
func VreinterpretqU64S64(v0 Int64X2) Uint64X2 {
	return C.vreinterpretq_u64_s64(v0)
}

// Vector reinterpret cast operation
func VreinterpretqU64S16(v0 Int16X8) Uint64X2 {
	return C.vreinterpretq_u64_s16(v0)
}

// Vector reinterpret cast operation
func VreinterpretqU16P8(v0 Poly8X16) Uint16X8 {
	return C.vreinterpretq_u16_p8(v0)
}

// Vector reinterpret cast operation
func VreinterpretqU16P128(v0 Poly128) Uint16X8 {
	return C.vreinterpretq_u16_p128(v0)
}

// Vector reinterpret cast operation
func VreinterpretqU16P64(v0 Poly64X2) Uint16X8 {
	return C.vreinterpretq_u16_p64(v0)
}

// Vector reinterpret cast operation
func VreinterpretqU16P16(v0 Poly16X8) Uint16X8 {
	return C.vreinterpretq_u16_p16(v0)
}

// Vector reinterpret cast operation
func VreinterpretqU16U8(v0 Uint8X16) Uint16X8 {
	return C.vreinterpretq_u16_u8(v0)
}

// Vector reinterpret cast operation
func VreinterpretqU16U32(v0 Uint32X4) Uint16X8 {
	return C.vreinterpretq_u16_u32(v0)
}

// Vector reinterpret cast operation
func VreinterpretqU16U64(v0 Uint64X2) Uint16X8 {
	return C.vreinterpretq_u16_u64(v0)
}

// Vector reinterpret cast operation
func VreinterpretqU16S8(v0 Int8X16) Uint16X8 {
	return C.vreinterpretq_u16_s8(v0)
}

// Vector reinterpret cast operation
func VreinterpretqU16F64(v0 Float64X2) Uint16X8 {
	return C.vreinterpretq_u16_f64(v0)
}

// Vector reinterpret cast operation
func VreinterpretqU16F32(v0 Float32X4) Uint16X8 {
	return C.vreinterpretq_u16_f32(v0)
}

// Vector reinterpret cast operation
func VreinterpretqU16S32(v0 Int32X4) Uint16X8 {
	return C.vreinterpretq_u16_s32(v0)
}

// Vector reinterpret cast operation
func VreinterpretqU16S64(v0 Int64X2) Uint16X8 {
	return C.vreinterpretq_u16_s64(v0)
}

// Vector reinterpret cast operation
func VreinterpretqU16S16(v0 Int16X8) Uint16X8 {
	return C.vreinterpretq_u16_s16(v0)
}

// Vector reinterpret cast operation
func VreinterpretqS8P8(v0 Poly8X16) Int8X16 {
	return C.vreinterpretq_s8_p8(v0)
}

// Vector reinterpret cast operation
func VreinterpretqS8P128(v0 Poly128) Int8X16 {
	return C.vreinterpretq_s8_p128(v0)
}

// Vector reinterpret cast operation
func VreinterpretqS8P64(v0 Poly64X2) Int8X16 {
	return C.vreinterpretq_s8_p64(v0)
}

// Vector reinterpret cast operation
func VreinterpretqS8P16(v0 Poly16X8) Int8X16 {
	return C.vreinterpretq_s8_p16(v0)
}

// Vector reinterpret cast operation
func VreinterpretqS8U8(v0 Uint8X16) Int8X16 {
	return C.vreinterpretq_s8_u8(v0)
}

// Vector reinterpret cast operation
func VreinterpretqS8U32(v0 Uint32X4) Int8X16 {
	return C.vreinterpretq_s8_u32(v0)
}

// Vector reinterpret cast operation
func VreinterpretqS8U64(v0 Uint64X2) Int8X16 {
	return C.vreinterpretq_s8_u64(v0)
}

// Vector reinterpret cast operation
func VreinterpretqS8U16(v0 Uint16X8) Int8X16 {
	return C.vreinterpretq_s8_u16(v0)
}

// Vector reinterpret cast operation
func VreinterpretqS8F64(v0 Float64X2) Int8X16 {
	return C.vreinterpretq_s8_f64(v0)
}

// Vector reinterpret cast operation
func VreinterpretqS8F32(v0 Float32X4) Int8X16 {
	return C.vreinterpretq_s8_f32(v0)
}

// Vector reinterpret cast operation
func VreinterpretqS8S32(v0 Int32X4) Int8X16 {
	return C.vreinterpretq_s8_s32(v0)
}

// Vector reinterpret cast operation
func VreinterpretqS8S64(v0 Int64X2) Int8X16 {
	return C.vreinterpretq_s8_s64(v0)
}

// Vector reinterpret cast operation
func VreinterpretqS8S16(v0 Int16X8) Int8X16 {
	return C.vreinterpretq_s8_s16(v0)
}

// Vector reinterpret cast operation
func VreinterpretqF64P8(v0 Poly8X16) Float64X2 {
	return C.vreinterpretq_f64_p8(v0)
}

// Vector reinterpret cast operation
func VreinterpretqF64P128(v0 Poly128) Float64X2 {
	return C.vreinterpretq_f64_p128(v0)
}

// Vector reinterpret cast operation
func VreinterpretqF64P64(v0 Poly64X2) Float64X2 {
	return C.vreinterpretq_f64_p64(v0)
}

// Vector reinterpret cast operation
func VreinterpretqF64P16(v0 Poly16X8) Float64X2 {
	return C.vreinterpretq_f64_p16(v0)
}

// Vector reinterpret cast operation
func VreinterpretqF64U8(v0 Uint8X16) Float64X2 {
	return C.vreinterpretq_f64_u8(v0)
}

// Vector reinterpret cast operation
func VreinterpretqF64U32(v0 Uint32X4) Float64X2 {
	return C.vreinterpretq_f64_u32(v0)
}

// Vector reinterpret cast operation
func VreinterpretqF64U64(v0 Uint64X2) Float64X2 {
	return C.vreinterpretq_f64_u64(v0)
}

// Vector reinterpret cast operation
func VreinterpretqF64U16(v0 Uint16X8) Float64X2 {
	return C.vreinterpretq_f64_u16(v0)
}

// Vector reinterpret cast operation
func VreinterpretqF64S8(v0 Int8X16) Float64X2 {
	return C.vreinterpretq_f64_s8(v0)
}

// Vector reinterpret cast operation
func VreinterpretqF64F32(v0 Float32X4) Float64X2 {
	return C.vreinterpretq_f64_f32(v0)
}

// Vector reinterpret cast operation
func VreinterpretqF64S32(v0 Int32X4) Float64X2 {
	return C.vreinterpretq_f64_s32(v0)
}

// Vector reinterpret cast operation
func VreinterpretqF64S64(v0 Int64X2) Float64X2 {
	return C.vreinterpretq_f64_s64(v0)
}

// Vector reinterpret cast operation
func VreinterpretqF64S16(v0 Int16X8) Float64X2 {
	return C.vreinterpretq_f64_s16(v0)
}

// Vector reinterpret cast operation
func VreinterpretqF32P8(v0 Poly8X16) Float32X4 {
	return C.vreinterpretq_f32_p8(v0)
}

// vreinterpretq_f32_p128
func VreinterpretqF32P128(v0 Poly128) Float32X4 {
	return C.vreinterpretq_f32_p128(v0)
}

// vreinterpretq_f32_p64
func VreinterpretqF32P64(v0 Poly64X2) Float32X4 {
	return C.vreinterpretq_f32_p64(v0)
}

// Vector reinterpret cast operation
func VreinterpretqF32P16(v0 Poly16X8) Float32X4 {
	return C.vreinterpretq_f32_p16(v0)
}

// Vector reinterpret cast operation
func VreinterpretqF32U8(v0 Uint8X16) Float32X4 {
	return C.vreinterpretq_f32_u8(v0)
}

// Vector reinterpret cast operation
func VreinterpretqF32U32(v0 Uint32X4) Float32X4 {
	return C.vreinterpretq_f32_u32(v0)
}

// Vector reinterpret cast operation
func VreinterpretqF32U64(v0 Uint64X2) Float32X4 {
	return C.vreinterpretq_f32_u64(v0)
}

// Vector reinterpret cast operation
func VreinterpretqF32U16(v0 Uint16X8) Float32X4 {
	return C.vreinterpretq_f32_u16(v0)
}

// Vector reinterpret cast operation
func VreinterpretqF32S8(v0 Int8X16) Float32X4 {
	return C.vreinterpretq_f32_s8(v0)
}

// Vector reinterpret cast operation
func VreinterpretqF32F64(v0 Float64X2) Float32X4 {
	return C.vreinterpretq_f32_f64(v0)
}

// Vector reinterpret cast operation
func VreinterpretqF32S32(v0 Int32X4) Float32X4 {
	return C.vreinterpretq_f32_s32(v0)
}

// Vector reinterpret cast operation
func VreinterpretqF32S64(v0 Int64X2) Float32X4 {
	return C.vreinterpretq_f32_s64(v0)
}

// Vector reinterpret cast operation
func VreinterpretqF32S16(v0 Int16X8) Float32X4 {
	return C.vreinterpretq_f32_s16(v0)
}

// Vector reinterpret cast operation
func VreinterpretqS32P8(v0 Poly8X16) Int32X4 {
	return C.vreinterpretq_s32_p8(v0)
}

// Vector reinterpret cast operation
func VreinterpretqS32P128(v0 Poly128) Int32X4 {
	return C.vreinterpretq_s32_p128(v0)
}

// Vector reinterpret cast operation
func VreinterpretqS32P64(v0 Poly64X2) Int32X4 {
	return C.vreinterpretq_s32_p64(v0)
}

// Vector reinterpret cast operation
func VreinterpretqS32P16(v0 Poly16X8) Int32X4 {
	return C.vreinterpretq_s32_p16(v0)
}

// Vector reinterpret cast operation
func VreinterpretqS32U8(v0 Uint8X16) Int32X4 {
	return C.vreinterpretq_s32_u8(v0)
}

// Vector reinterpret cast operation
func VreinterpretqS32U32(v0 Uint32X4) Int32X4 {
	return C.vreinterpretq_s32_u32(v0)
}

// Vector reinterpret cast operation
func VreinterpretqS32U64(v0 Uint64X2) Int32X4 {
	return C.vreinterpretq_s32_u64(v0)
}

// Vector reinterpret cast operation
func VreinterpretqS32U16(v0 Uint16X8) Int32X4 {
	return C.vreinterpretq_s32_u16(v0)
}

// Vector reinterpret cast operation
func VreinterpretqS32S8(v0 Int8X16) Int32X4 {
	return C.vreinterpretq_s32_s8(v0)
}

// Vector reinterpret cast operation
func VreinterpretqS32F64(v0 Float64X2) Int32X4 {
	return C.vreinterpretq_s32_f64(v0)
}

// Vector reinterpret cast operation
func VreinterpretqS32F32(v0 Float32X4) Int32X4 {
	return C.vreinterpretq_s32_f32(v0)
}

// Vector reinterpret cast operation
func VreinterpretqS32S64(v0 Int64X2) Int32X4 {
	return C.vreinterpretq_s32_s64(v0)
}

// Vector reinterpret cast operation
func VreinterpretqS32S16(v0 Int16X8) Int32X4 {
	return C.vreinterpretq_s32_s16(v0)
}

// Vector reinterpret cast operation
func VreinterpretqS64P8(v0 Poly8X16) Int64X2 {
	return C.vreinterpretq_s64_p8(v0)
}

// Vector reinterpret cast operation
func VreinterpretqS64P128(v0 Poly128) Int64X2 {
	return C.vreinterpretq_s64_p128(v0)
}

// Vector reinterpret cast operation
func VreinterpretqS64P64(v0 Poly64X2) Int64X2 {
	return C.vreinterpretq_s64_p64(v0)
}

// Vector reinterpret cast operation
func VreinterpretqS64P16(v0 Poly16X8) Int64X2 {
	return C.vreinterpretq_s64_p16(v0)
}

// Vector reinterpret cast operation
func VreinterpretqS64U8(v0 Uint8X16) Int64X2 {
	return C.vreinterpretq_s64_u8(v0)
}

// Vector reinterpret cast operation
func VreinterpretqS64U32(v0 Uint32X4) Int64X2 {
	return C.vreinterpretq_s64_u32(v0)
}

// Vector reinterpret cast operation
func VreinterpretqS64U64(v0 Uint64X2) Int64X2 {
	return C.vreinterpretq_s64_u64(v0)
}

// Vector reinterpret cast operation
func VreinterpretqS64U16(v0 Uint16X8) Int64X2 {
	return C.vreinterpretq_s64_u16(v0)
}

// Vector reinterpret cast operation
func VreinterpretqS64S8(v0 Int8X16) Int64X2 {
	return C.vreinterpretq_s64_s8(v0)
}

// Vector reinterpret cast operation
func VreinterpretqS64F64(v0 Float64X2) Int64X2 {
	return C.vreinterpretq_s64_f64(v0)
}

// Vector reinterpret cast operation
func VreinterpretqS64F32(v0 Float32X4) Int64X2 {
	return C.vreinterpretq_s64_f32(v0)
}

// Vector reinterpret cast operation
func VreinterpretqS64S32(v0 Int32X4) Int64X2 {
	return C.vreinterpretq_s64_s32(v0)
}

// Vector reinterpret cast operation
func VreinterpretqS64S16(v0 Int16X8) Int64X2 {
	return C.vreinterpretq_s64_s16(v0)
}

// Vector reinterpret cast operation
func VreinterpretqS16P8(v0 Poly8X16) Int16X8 {
	return C.vreinterpretq_s16_p8(v0)
}

// Vector reinterpret cast operation
func VreinterpretqS16P128(v0 Poly128) Int16X8 {
	return C.vreinterpretq_s16_p128(v0)
}

// Vector reinterpret cast operation
func VreinterpretqS16P64(v0 Poly64X2) Int16X8 {
	return C.vreinterpretq_s16_p64(v0)
}

// Vector reinterpret cast operation
func VreinterpretqS16P16(v0 Poly16X8) Int16X8 {
	return C.vreinterpretq_s16_p16(v0)
}

// Vector reinterpret cast operation
func VreinterpretqS16U8(v0 Uint8X16) Int16X8 {
	return C.vreinterpretq_s16_u8(v0)
}

// Vector reinterpret cast operation
func VreinterpretqS16U32(v0 Uint32X4) Int16X8 {
	return C.vreinterpretq_s16_u32(v0)
}

// Vector reinterpret cast operation
func VreinterpretqS16U64(v0 Uint64X2) Int16X8 {
	return C.vreinterpretq_s16_u64(v0)
}

// Vector reinterpret cast operation
func VreinterpretqS16U16(v0 Uint16X8) Int16X8 {
	return C.vreinterpretq_s16_u16(v0)
}

// Vector reinterpret cast operation
func VreinterpretqS16S8(v0 Int8X16) Int16X8 {
	return C.vreinterpretq_s16_s8(v0)
}

// Vector reinterpret cast operation
func VreinterpretqS16F64(v0 Float64X2) Int16X8 {
	return C.vreinterpretq_s16_f64(v0)
}

// Vector reinterpret cast operation
func VreinterpretqS16F32(v0 Float32X4) Int16X8 {
	return C.vreinterpretq_s16_f32(v0)
}

// Vector reinterpret cast operation
func VreinterpretqS16S32(v0 Int32X4) Int16X8 {
	return C.vreinterpretq_s16_s32(v0)
}

// Vector reinterpret cast operation
func VreinterpretqS16S64(v0 Int64X2) Int16X8 {
	return C.vreinterpretq_s16_s64(v0)
}

// Vector reinterpret cast operation
func VreinterpretU8P8(v0 Poly8X8) Uint8X8 {
	return C.vreinterpret_u8_p8(v0)
}

// Vector reinterpret cast operation
func VreinterpretU8P64(v0 Poly64X1) Uint8X8 {
	return C.vreinterpret_u8_p64(v0)
}

// Vector reinterpret cast operation
func VreinterpretU8P16(v0 Poly16X4) Uint8X8 {
	return C.vreinterpret_u8_p16(v0)
}

// Vector reinterpret cast operation
func VreinterpretU8U32(v0 Uint32X2) Uint8X8 {
	return C.vreinterpret_u8_u32(v0)
}

// Vector reinterpret cast operation
func VreinterpretU8U64(v0 Uint64X1) Uint8X8 {
	return C.vreinterpret_u8_u64(v0)
}

// Vector reinterpret cast operation
func VreinterpretU8U16(v0 Uint16X4) Uint8X8 {
	return C.vreinterpret_u8_u16(v0)
}

// Vector reinterpret cast operation
func VreinterpretU8S8(v0 Int8X8) Uint8X8 {
	return C.vreinterpret_u8_s8(v0)
}

// Vector reinterpret cast operation
func VreinterpretU8F64(v0 Float64X1) Uint8X8 {
	return C.vreinterpret_u8_f64(v0)
}

// Vector reinterpret cast operation
func VreinterpretU8F32(v0 Float32X2) Uint8X8 {
	return C.vreinterpret_u8_f32(v0)
}

// Vector reinterpret cast operation
func VreinterpretU8S32(v0 Int32X2) Uint8X8 {
	return C.vreinterpret_u8_s32(v0)
}

// Vector reinterpret cast operation
func VreinterpretU8S64(v0 Int64X1) Uint8X8 {
	return C.vreinterpret_u8_s64(v0)
}

// Vector reinterpret cast operation
func VreinterpretU8S16(v0 Int16X4) Uint8X8 {
	return C.vreinterpret_u8_s16(v0)
}

// Vector reinterpret cast operation
func VreinterpretU32P8(v0 Poly8X8) Uint32X2 {
	return C.vreinterpret_u32_p8(v0)
}

// Vector reinterpret cast operation
func VreinterpretU32P64(v0 Poly64X1) Uint32X2 {
	return C.vreinterpret_u32_p64(v0)
}

// Vector reinterpret cast operation
func VreinterpretU32P16(v0 Poly16X4) Uint32X2 {
	return C.vreinterpret_u32_p16(v0)
}

// Vector reinterpret cast operation
func VreinterpretU32U8(v0 Uint8X8) Uint32X2 {
	return C.vreinterpret_u32_u8(v0)
}

// Vector reinterpret cast operation
func VreinterpretU32U64(v0 Uint64X1) Uint32X2 {
	return C.vreinterpret_u32_u64(v0)
}

// Vector reinterpret cast operation
func VreinterpretU32U16(v0 Uint16X4) Uint32X2 {
	return C.vreinterpret_u32_u16(v0)
}

// Vector reinterpret cast operation
func VreinterpretU32S8(v0 Int8X8) Uint32X2 {
	return C.vreinterpret_u32_s8(v0)
}

// Vector reinterpret cast operation
func VreinterpretU32F64(v0 Float64X1) Uint32X2 {
	return C.vreinterpret_u32_f64(v0)
}

// Vector reinterpret cast operation
func VreinterpretU32F32(v0 Float32X2) Uint32X2 {
	return C.vreinterpret_u32_f32(v0)
}

// Vector reinterpret cast operation
func VreinterpretU32S32(v0 Int32X2) Uint32X2 {
	return C.vreinterpret_u32_s32(v0)
}

// Vector reinterpret cast operation
func VreinterpretU32S64(v0 Int64X1) Uint32X2 {
	return C.vreinterpret_u32_s64(v0)
}

// Vector reinterpret cast operation
func VreinterpretU32S16(v0 Int16X4) Uint32X2 {
	return C.vreinterpret_u32_s16(v0)
}

// Vector reinterpret cast operation
func VreinterpretU64P8(v0 Poly8X8) Uint64X1 {
	return C.vreinterpret_u64_p8(v0)
}

// Vector reinterpret cast operation
func VreinterpretU64P64(v0 Poly64X1) Uint64X1 {
	return C.vreinterpret_u64_p64(v0)
}

// Vector reinterpret cast operation
func VreinterpretU64P16(v0 Poly16X4) Uint64X1 {
	return C.vreinterpret_u64_p16(v0)
}

// Vector reinterpret cast operation
func VreinterpretU64U8(v0 Uint8X8) Uint64X1 {
	return C.vreinterpret_u64_u8(v0)
}

// Vector reinterpret cast operation
func VreinterpretU64U32(v0 Uint32X2) Uint64X1 {
	return C.vreinterpret_u64_u32(v0)
}

// Vector reinterpret cast operation
func VreinterpretU64U16(v0 Uint16X4) Uint64X1 {
	return C.vreinterpret_u64_u16(v0)
}

// Vector reinterpret cast operation
func VreinterpretU64S8(v0 Int8X8) Uint64X1 {
	return C.vreinterpret_u64_s8(v0)
}

// Vector reinterpret cast operation
func VreinterpretU64F64(v0 Float64X1) Uint64X1 {
	return C.vreinterpret_u64_f64(v0)
}

// Vector reinterpret cast operation
func VreinterpretU64F32(v0 Float32X2) Uint64X1 {
	return C.vreinterpret_u64_f32(v0)
}

// Vector reinterpret cast operation
func VreinterpretU64S32(v0 Int32X2) Uint64X1 {
	return C.vreinterpret_u64_s32(v0)
}

// Vector reinterpret cast operation
func VreinterpretU64S64(v0 Int64X1) Uint64X1 {
	return C.vreinterpret_u64_s64(v0)
}

// Vector reinterpret cast operation
func VreinterpretU64S16(v0 Int16X4) Uint64X1 {
	return C.vreinterpret_u64_s16(v0)
}

// Vector reinterpret cast operation
func VreinterpretU16P8(v0 Poly8X8) Uint16X4 {
	return C.vreinterpret_u16_p8(v0)
}

// Vector reinterpret cast operation
func VreinterpretU16P64(v0 Poly64X1) Uint16X4 {
	return C.vreinterpret_u16_p64(v0)
}

// Vector reinterpret cast operation
func VreinterpretU16P16(v0 Poly16X4) Uint16X4 {
	return C.vreinterpret_u16_p16(v0)
}

// Vector reinterpret cast operation
func VreinterpretU16U8(v0 Uint8X8) Uint16X4 {
	return C.vreinterpret_u16_u8(v0)
}

// Vector reinterpret cast operation
func VreinterpretU16U32(v0 Uint32X2) Uint16X4 {
	return C.vreinterpret_u16_u32(v0)
}

// Vector reinterpret cast operation
func VreinterpretU16U64(v0 Uint64X1) Uint16X4 {
	return C.vreinterpret_u16_u64(v0)
}

// Vector reinterpret cast operation
func VreinterpretU16S8(v0 Int8X8) Uint16X4 {
	return C.vreinterpret_u16_s8(v0)
}

// Vector reinterpret cast operation
func VreinterpretU16F64(v0 Float64X1) Uint16X4 {
	return C.vreinterpret_u16_f64(v0)
}

// Vector reinterpret cast operation
func VreinterpretU16F32(v0 Float32X2) Uint16X4 {
	return C.vreinterpret_u16_f32(v0)
}

// Vector reinterpret cast operation
func VreinterpretU16S32(v0 Int32X2) Uint16X4 {
	return C.vreinterpret_u16_s32(v0)
}

// Vector reinterpret cast operation
func VreinterpretU16S64(v0 Int64X1) Uint16X4 {
	return C.vreinterpret_u16_s64(v0)
}

// Vector reinterpret cast operation
func VreinterpretU16S16(v0 Int16X4) Uint16X4 {
	return C.vreinterpret_u16_s16(v0)
}

// Vector reinterpret cast operation
func VreinterpretS8P8(v0 Poly8X8) Int8X8 {
	return C.vreinterpret_s8_p8(v0)
}

// Vector reinterpret cast operation
func VreinterpretS8P64(v0 Poly64X1) Int8X8 {
	return C.vreinterpret_s8_p64(v0)
}

// Vector reinterpret cast operation
func VreinterpretS8P16(v0 Poly16X4) Int8X8 {
	return C.vreinterpret_s8_p16(v0)
}

// Vector reinterpret cast operation
func VreinterpretS8U8(v0 Uint8X8) Int8X8 {
	return C.vreinterpret_s8_u8(v0)
}

// Vector reinterpret cast operation
func VreinterpretS8U32(v0 Uint32X2) Int8X8 {
	return C.vreinterpret_s8_u32(v0)
}

// Vector reinterpret cast operation
func VreinterpretS8U64(v0 Uint64X1) Int8X8 {
	return C.vreinterpret_s8_u64(v0)
}

// Vector reinterpret cast operation
func VreinterpretS8U16(v0 Uint16X4) Int8X8 {
	return C.vreinterpret_s8_u16(v0)
}

// Vector reinterpret cast operation
func VreinterpretS8F64(v0 Float64X1) Int8X8 {
	return C.vreinterpret_s8_f64(v0)
}

// Vector reinterpret cast operation
func VreinterpretS8F32(v0 Float32X2) Int8X8 {
	return C.vreinterpret_s8_f32(v0)
}

// Vector reinterpret cast operation
func VreinterpretS8S32(v0 Int32X2) Int8X8 {
	return C.vreinterpret_s8_s32(v0)
}

// Vector reinterpret cast operation
func VreinterpretS8S64(v0 Int64X1) Int8X8 {
	return C.vreinterpret_s8_s64(v0)
}

// Vector reinterpret cast operation
func VreinterpretS8S16(v0 Int16X4) Int8X8 {
	return C.vreinterpret_s8_s16(v0)
}

// Vector reinterpret cast operation
func VreinterpretF64P8(v0 Poly8X8) Float64X1 {
	return C.vreinterpret_f64_p8(v0)
}

// Vector reinterpret cast operation
func VreinterpretF64P64(v0 Poly64X1) Float64X1 {
	return C.vreinterpret_f64_p64(v0)
}

// Vector reinterpret cast operation
func VreinterpretF64P16(v0 Poly16X4) Float64X1 {
	return C.vreinterpret_f64_p16(v0)
}

// Vector reinterpret cast operation
func VreinterpretF64U8(v0 Uint8X8) Float64X1 {
	return C.vreinterpret_f64_u8(v0)
}

// Vector reinterpret cast operation
func VreinterpretF64U32(v0 Uint32X2) Float64X1 {
	return C.vreinterpret_f64_u32(v0)
}

// Vector reinterpret cast operation
func VreinterpretF64U64(v0 Uint64X1) Float64X1 {
	return C.vreinterpret_f64_u64(v0)
}

// Vector reinterpret cast operation
func VreinterpretF64U16(v0 Uint16X4) Float64X1 {
	return C.vreinterpret_f64_u16(v0)
}

// Vector reinterpret cast operation
func VreinterpretF64S8(v0 Int8X8) Float64X1 {
	return C.vreinterpret_f64_s8(v0)
}

// Vector reinterpret cast operation
func VreinterpretF64F32(v0 Float32X2) Float64X1 {
	return C.vreinterpret_f64_f32(v0)
}

// Vector reinterpret cast operation
func VreinterpretF64S32(v0 Int32X2) Float64X1 {
	return C.vreinterpret_f64_s32(v0)
}

// Vector reinterpret cast operation
func VreinterpretF64S64(v0 Int64X1) Float64X1 {
	return C.vreinterpret_f64_s64(v0)
}

// Vector reinterpret cast operation
func VreinterpretF64S16(v0 Int16X4) Float64X1 {
	return C.vreinterpret_f64_s16(v0)
}

// Vector reinterpret cast operation
func VreinterpretF32P8(v0 Poly8X8) Float32X2 {
	return C.vreinterpret_f32_p8(v0)
}

// vreinterpret_f32_p64
func VreinterpretF32P64(v0 Poly64X1) Float32X2 {
	return C.vreinterpret_f32_p64(v0)
}

// Vector reinterpret cast operation
func VreinterpretF32P16(v0 Poly16X4) Float32X2 {
	return C.vreinterpret_f32_p16(v0)
}

// Vector reinterpret cast operation
func VreinterpretF32U8(v0 Uint8X8) Float32X2 {
	return C.vreinterpret_f32_u8(v0)
}

// Vector reinterpret cast operation
func VreinterpretF32U32(v0 Uint32X2) Float32X2 {
	return C.vreinterpret_f32_u32(v0)
}

// Vector reinterpret cast operation
func VreinterpretF32U64(v0 Uint64X1) Float32X2 {
	return C.vreinterpret_f32_u64(v0)
}

// Vector reinterpret cast operation
func VreinterpretF32U16(v0 Uint16X4) Float32X2 {
	return C.vreinterpret_f32_u16(v0)
}

// Vector reinterpret cast operation
func VreinterpretF32S8(v0 Int8X8) Float32X2 {
	return C.vreinterpret_f32_s8(v0)
}

// Vector reinterpret cast operation
func VreinterpretF32F64(v0 Float64X1) Float32X2 {
	return C.vreinterpret_f32_f64(v0)
}

// Vector reinterpret cast operation
func VreinterpretF32S32(v0 Int32X2) Float32X2 {
	return C.vreinterpret_f32_s32(v0)
}

// Vector reinterpret cast operation
func VreinterpretF32S64(v0 Int64X1) Float32X2 {
	return C.vreinterpret_f32_s64(v0)
}

// Vector reinterpret cast operation
func VreinterpretF32S16(v0 Int16X4) Float32X2 {
	return C.vreinterpret_f32_s16(v0)
}

// Vector reinterpret cast operation
func VreinterpretS32P8(v0 Poly8X8) Int32X2 {
	return C.vreinterpret_s32_p8(v0)
}

// Vector reinterpret cast operation
func VreinterpretS32P64(v0 Poly64X1) Int32X2 {
	return C.vreinterpret_s32_p64(v0)
}

// Vector reinterpret cast operation
func VreinterpretS32P16(v0 Poly16X4) Int32X2 {
	return C.vreinterpret_s32_p16(v0)
}

// Vector reinterpret cast operation
func VreinterpretS32U8(v0 Uint8X8) Int32X2 {
	return C.vreinterpret_s32_u8(v0)
}

// Vector reinterpret cast operation
func VreinterpretS32U32(v0 Uint32X2) Int32X2 {
	return C.vreinterpret_s32_u32(v0)
}

// Vector reinterpret cast operation
func VreinterpretS32U64(v0 Uint64X1) Int32X2 {
	return C.vreinterpret_s32_u64(v0)
}

// Vector reinterpret cast operation
func VreinterpretS32U16(v0 Uint16X4) Int32X2 {
	return C.vreinterpret_s32_u16(v0)
}

// Vector reinterpret cast operation
func VreinterpretS32S8(v0 Int8X8) Int32X2 {
	return C.vreinterpret_s32_s8(v0)
}

// Vector reinterpret cast operation
func VreinterpretS32F64(v0 Float64X1) Int32X2 {
	return C.vreinterpret_s32_f64(v0)
}

// Vector reinterpret cast operation
func VreinterpretS32F32(v0 Float32X2) Int32X2 {
	return C.vreinterpret_s32_f32(v0)
}

// Vector reinterpret cast operation
func VreinterpretS32S64(v0 Int64X1) Int32X2 {
	return C.vreinterpret_s32_s64(v0)
}

// Vector reinterpret cast operation
func VreinterpretS32S16(v0 Int16X4) Int32X2 {
	return C.vreinterpret_s32_s16(v0)
}

// Vector reinterpret cast operation
func VreinterpretS64P8(v0 Poly8X8) Int64X1 {
	return C.vreinterpret_s64_p8(v0)
}

// Vector reinterpret cast operation
func VreinterpretS64P64(v0 Poly64X1) Int64X1 {
	return C.vreinterpret_s64_p64(v0)
}

// Vector reinterpret cast operation
func VreinterpretS64P16(v0 Poly16X4) Int64X1 {
	return C.vreinterpret_s64_p16(v0)
}

// Vector reinterpret cast operation
func VreinterpretS64U8(v0 Uint8X8) Int64X1 {
	return C.vreinterpret_s64_u8(v0)
}

// Vector reinterpret cast operation
func VreinterpretS64U32(v0 Uint32X2) Int64X1 {
	return C.vreinterpret_s64_u32(v0)
}

// Vector reinterpret cast operation
func VreinterpretS64U64(v0 Uint64X1) Int64X1 {
	return C.vreinterpret_s64_u64(v0)
}

// Vector reinterpret cast operation
func VreinterpretS64U16(v0 Uint16X4) Int64X1 {
	return C.vreinterpret_s64_u16(v0)
}

// Vector reinterpret cast operation
func VreinterpretS64S8(v0 Int8X8) Int64X1 {
	return C.vreinterpret_s64_s8(v0)
}

// Vector reinterpret cast operation
func VreinterpretS64F64(v0 Float64X1) Int64X1 {
	return C.vreinterpret_s64_f64(v0)
}

// Vector reinterpret cast operation
func VreinterpretS64F32(v0 Float32X2) Int64X1 {
	return C.vreinterpret_s64_f32(v0)
}

// Vector reinterpret cast operation
func VreinterpretS64S32(v0 Int32X2) Int64X1 {
	return C.vreinterpret_s64_s32(v0)
}

// Vector reinterpret cast operation
func VreinterpretS64S16(v0 Int16X4) Int64X1 {
	return C.vreinterpret_s64_s16(v0)
}

// Vector reinterpret cast operation
func VreinterpretS16P8(v0 Poly8X8) Int16X4 {
	return C.vreinterpret_s16_p8(v0)
}

// Vector reinterpret cast operation
func VreinterpretS16P64(v0 Poly64X1) Int16X4 {
	return C.vreinterpret_s16_p64(v0)
}

// Vector reinterpret cast operation
func VreinterpretS16P16(v0 Poly16X4) Int16X4 {
	return C.vreinterpret_s16_p16(v0)
}

// Vector reinterpret cast operation
func VreinterpretS16U8(v0 Uint8X8) Int16X4 {
	return C.vreinterpret_s16_u8(v0)
}

// Vector reinterpret cast operation
func VreinterpretS16U32(v0 Uint32X2) Int16X4 {
	return C.vreinterpret_s16_u32(v0)
}

// Vector reinterpret cast operation
func VreinterpretS16U64(v0 Uint64X1) Int16X4 {
	return C.vreinterpret_s16_u64(v0)
}

// Vector reinterpret cast operation
func VreinterpretS16U16(v0 Uint16X4) Int16X4 {
	return C.vreinterpret_s16_u16(v0)
}

// Vector reinterpret cast operation
func VreinterpretS16S8(v0 Int8X8) Int16X4 {
	return C.vreinterpret_s16_s8(v0)
}

// Vector reinterpret cast operation
func VreinterpretS16F64(v0 Float64X1) Int16X4 {
	return C.vreinterpret_s16_f64(v0)
}

// Vector reinterpret cast operation
func VreinterpretS16F32(v0 Float32X2) Int16X4 {
	return C.vreinterpret_s16_f32(v0)
}

// Vector reinterpret cast operation
func VreinterpretS16S32(v0 Int32X2) Int16X4 {
	return C.vreinterpret_s16_s32(v0)
}

// Vector reinterpret cast operation
func VreinterpretS16S64(v0 Int64X1) Int16X4 {
	return C.vreinterpret_s16_s64(v0)
}

// Floating-point Round to Integral, toward Zero (vector). This instruction rounds a vector of floating-point values in the SIMD&FP source register to integral floating-point values of the same size using the Round towards Zero rounding mode, and writes the result to the SIMD&FP destination register.
func VrndqF64(v0 Float64X2) Float64X2 {
	return C.vrndq_f64(v0)
}

// Floating-point Round to Integral, toward Zero (vector). This instruction rounds a vector of floating-point values in the SIMD&FP source register to integral floating-point values of the same size using the Round towards Zero rounding mode, and writes the result to the SIMD&FP destination register.
func VrndF64(v0 Float64X1) Float64X1 {
	return C.vrnd_f64(v0)
}

// Floating-point Round to Integral, to nearest with ties to Away (vector). This instruction rounds a vector of floating-point values in the SIMD&FP source register to integral floating-point values of the same size using the Round to Nearest with Ties to Away rounding mode, and writes the result to the SIMD&FP destination register.
func VrndaqF64(v0 Float64X2) Float64X2 {
	return C.vrndaq_f64(v0)
}

// Floating-point Round to Integral, to nearest with ties to Away (vector). This instruction rounds a vector of floating-point values in the SIMD&FP source register to integral floating-point values of the same size using the Round to Nearest with Ties to Away rounding mode, and writes the result to the SIMD&FP destination register.
func VrndaF64(v0 Float64X1) Float64X1 {
	return C.vrnda_f64(v0)
}

// Floating-point Round to Integral, using current rounding mode (vector). This instruction rounds a vector of floating-point values in the SIMD&FP source register to integral floating-point values of the same size using the rounding mode that is determined by the FPCR, and writes the result to the SIMD&FP destination register.
func VrndiqF64(v0 Float64X2) Float64X2 {
	return C.vrndiq_f64(v0)
}

// Floating-point Round to Integral, using current rounding mode (vector). This instruction rounds a vector of floating-point values in the SIMD&FP source register to integral floating-point values of the same size using the rounding mode that is determined by the FPCR, and writes the result to the SIMD&FP destination register.
func VrndiF64(v0 Float64X1) Float64X1 {
	return C.vrndi_f64(v0)
}

// Floating-point Round to Integral, toward Minus infinity (vector). This instruction rounds a vector of floating-point values in the SIMD&FP source register to integral floating-point values of the same size using the Round towards Minus Infinity rounding mode, and writes the result to the SIMD&FP destination register.
func VrndmqF64(v0 Float64X2) Float64X2 {
	return C.vrndmq_f64(v0)
}

// Floating-point Round to Integral, toward Minus infinity (vector). This instruction rounds a vector of floating-point values in the SIMD&FP source register to integral floating-point values of the same size using the Round towards Minus Infinity rounding mode, and writes the result to the SIMD&FP destination register.
func VrndmF64(v0 Float64X1) Float64X1 {
	return C.vrndm_f64(v0)
}

// Floating-point Round to Integral, to nearest with ties to even (vector). This instruction rounds a vector of floating-point values in the SIMD&FP source register to integral floating-point values of the same size using the Round to Nearest rounding mode, and writes the result to the SIMD&FP destination register.
func VrndnqF64(v0 Float64X2) Float64X2 {
	return C.vrndnq_f64(v0)
}

// Floating-point Round to Integral, to nearest with ties to even (vector). This instruction rounds a vector of floating-point values in the SIMD&FP source register to integral floating-point values of the same size using the Round to Nearest rounding mode, and writes the result to the SIMD&FP destination register.
func VrndnF64(v0 Float64X1) Float64X1 {
	return C.vrndn_f64(v0)
}

// Floating-point Round to Integral, toward Plus infinity (vector). This instruction rounds a vector of floating-point values in the SIMD&FP source register to integral floating-point values of the same size using the Round towards Plus Infinity rounding mode, and writes the result to the SIMD&FP destination register.
func VrndpqF64(v0 Float64X2) Float64X2 {
	return C.vrndpq_f64(v0)
}

// Floating-point Round to Integral, toward Plus infinity (vector). This instruction rounds a vector of floating-point values in the SIMD&FP source register to integral floating-point values of the same size using the Round towards Plus Infinity rounding mode, and writes the result to the SIMD&FP destination register.
func VrndpF64(v0 Float64X1) Float64X1 {
	return C.vrndp_f64(v0)
}

// Floating-point Round to Integral exact, using current rounding mode (vector). This instruction rounds a vector of floating-point values in the SIMD&FP source register to integral floating-point values of the same size using the rounding mode that is determined by the FPCR, and writes the result to the SIMD&FP destination register.
func VrndxqF64(v0 Float64X2) Float64X2 {
	return C.vrndxq_f64(v0)
}

// Floating-point Round to Integral exact, using current rounding mode (vector). This instruction rounds a vector of floating-point values in the SIMD&FP source register to integral floating-point values of the same size using the rounding mode that is determined by the FPCR, and writes the result to the SIMD&FP destination register.
func VrndxF64(v0 Float64X1) Float64X1 {
	return C.vrndx_f64(v0)
}

// Floating-point Round to 32-bit Integer, using current rounding mode (vector). This instruction rounds a vector of floating-point values in the SIMD&FP source register to integral floating-point values that fit into a 32-bit integer size using the rounding mode that is determined by the FPCR, and writes the result to the SIMD&FP destination register.
func Vrnd32XqF32(v0 Float32X4) Float32X4 {
	return C.vrnd32xq_f32(v0)
}

// Floating-point Round to 32-bit Integer, using current rounding mode (vector). This instruction rounds a vector of floating-point values in the SIMD&FP source register to integral floating-point values that fit into a 32-bit integer size using the rounding mode that is determined by the FPCR, and writes the result to the SIMD&FP destination register.
func Vrnd32XF32(v0 Float32X2) Float32X2 {
	return C.vrnd32x_f32(v0)
}

// Floating-point Round to 32-bit Integer toward Zero (vector). This instruction rounds a vector of floating-point values in the SIMD&FP source register to integral floating-point values that fit into a 32-bit integer size using the Round towards Zero rounding mode, and writes the result to the SIMD&FP destination register.
func Vrnd32ZqF32(v0 Float32X4) Float32X4 {
	return C.vrnd32zq_f32(v0)
}

// Floating-point Round to 32-bit Integer toward Zero (vector). This instruction rounds a vector of floating-point values in the SIMD&FP source register to integral floating-point values that fit into a 32-bit integer size using the Round towards Zero rounding mode, and writes the result to the SIMD&FP destination register.
func Vrnd32ZF32(v0 Float32X2) Float32X2 {
	return C.vrnd32z_f32(v0)
}

// Floating-point Round to 64-bit Integer, using current rounding mode (vector). This instruction rounds a vector of floating-point values in the SIMD&FP source register to integral floating-point values that fit into a 64-bit integer size using the rounding mode that is determined by the FPCR, and writes the result to the SIMD&FP destination register.
func Vrnd64XqF32(v0 Float32X4) Float32X4 {
	return C.vrnd64xq_f32(v0)
}

// Floating-point Round to 64-bit Integer, using current rounding mode (vector). This instruction rounds a vector of floating-point values in the SIMD&FP source register to integral floating-point values that fit into a 64-bit integer size using the rounding mode that is determined by the FPCR, and writes the result to the SIMD&FP destination register.
func Vrnd64XF32(v0 Float32X2) Float32X2 {
	return C.vrnd64x_f32(v0)
}

// Floating-point Round to 64-bit Integer toward Zero (vector). This instruction rounds a vector of floating-point values in the SIMD&FP source register to integral floating-point values that fit into a 64-bit integer size using the Round towards Zero rounding mode, and writes the result to the SIMD&FP destination register.
func Vrnd64ZqF32(v0 Float32X4) Float32X4 {
	return C.vrnd64zq_f32(v0)
}

// Floating-point Round to 64-bit Integer toward Zero (vector). This instruction rounds a vector of floating-point values in the SIMD&FP source register to integral floating-point values that fit into a 64-bit integer size using the Round towards Zero rounding mode, and writes the result to the SIMD&FP destination register.
func Vrnd64ZF32(v0 Float32X2) Float32X2 {
	return C.vrnd64z_f32(v0)
}

// Floating-point Maximum Number (vector). This instruction compares corresponding vector elements in the two source SIMD&FP registers, writes the larger of the two floating-point values into a vector, and writes the vector to the destination SIMD&FP register.
func VmaxnmqF64(v0 Float64X2, v1 Float64X2) Float64X2 {
	return C.vmaxnmq_f64(v0, v1)
}

// Floating-point Maximum Number (vector). This instruction compares corresponding vector elements in the two source SIMD&FP registers, writes the larger of the two floating-point values into a vector, and writes the vector to the destination SIMD&FP register.
func VmaxnmF64(v0 Float64X1, v1 Float64X1) Float64X1 {
	return C.vmaxnm_f64(v0, v1)
}

// Floating-point Minimum Number (vector). This instruction compares corresponding vector elements in the two source SIMD&FP registers, writes the smaller of the two floating-point values into a vector, and writes the vector to the destination SIMD&FP register.
func VminnmqF64(v0 Float64X2, v1 Float64X2) Float64X2 {
	return C.vminnmq_f64(v0, v1)
}

// Floating-point Minimum Number (vector). This instruction compares corresponding vector elements in the two source SIMD&FP registers, writes the smaller of the two floating-point values into a vector, and writes the vector to the destination SIMD&FP register.
func VminnmF64(v0 Float64X1, v1 Float64X1) Float64X1 {
	return C.vminnm_f64(v0, v1)
}

// Floating-point Complex Add.
func VcaddRot270F32(v0 Float32X2, v1 Float32X2) Float32X2 {
	return C.vcadd_rot270_f32(v0, v1)
}

// Floating-point Complex Add.
func VcaddRot90F32(v0 Float32X2, v1 Float32X2) Float32X2 {
	return C.vcadd_rot90_f32(v0, v1)
}

// Floating-point Complex Add.
func VcaddqRot270F32(v0 Float32X4, v1 Float32X4) Float32X4 {
	return C.vcaddq_rot270_f32(v0, v1)
}

// Floating-point Complex Add.
func VcaddqRot90F32(v0 Float32X4, v1 Float32X4) Float32X4 {
	return C.vcaddq_rot90_f32(v0, v1)
}

// Floating-point Complex Add.
func VcaddqRot270F64(v0 Float64X2, v1 Float64X2) Float64X2 {
	return C.vcaddq_rot270_f64(v0, v1)
}

// Floating-point Complex Add.
func VcaddqRot90F64(v0 Float64X2, v1 Float64X2) Float64X2 {
	return C.vcaddq_rot90_f64(v0, v1)
}

// Dot Product unsigned arithmetic (vector). This instruction performs the dot product of the four unsigned 8-bit elements in each 32-bit element of the first source register with the four unsigned 8-bit elements of the corresponding 32-bit element in the second source register, accumulating the result into the corresponding 32-bit element of the destination register.
func VdotqU32(v0 Uint32X4, v1 Uint8X16, v2 Uint8X16) Uint32X4 {
	return C.vdotq_u32(v0, v1, v2)
}

// Dot Product signed arithmetic (vector). This instruction performs the dot product of the four signed 8-bit elements in each 32-bit element of the first source register with the four signed 8-bit elements of the corresponding 32-bit element in the second source register, accumulating the result into the corresponding 32-bit element of the destination register.
func VdotqS32(v0 Int32X4, v1 Int8X16, v2 Int8X16) Int32X4 {
	return C.vdotq_s32(v0, v1, v2)
}

// Dot Product unsigned arithmetic (vector). This instruction performs the dot product of the four unsigned 8-bit elements in each 32-bit element of the first source register with the four unsigned 8-bit elements of the corresponding 32-bit element in the second source register, accumulating the result into the corresponding 32-bit element of the destination register.
func VdotU32(v0 Uint32X2, v1 Uint8X8, v2 Uint8X8) Uint32X2 {
	return C.vdot_u32(v0, v1, v2)
}

// Dot Product signed arithmetic (vector). This instruction performs the dot product of the four signed 8-bit elements in each 32-bit element of the first source register with the four signed 8-bit elements of the corresponding 32-bit element in the second source register, accumulating the result into the corresponding 32-bit element of the destination register.
func VdotS32(v0 Int32X2, v1 Int8X8, v2 Int8X8) Int32X2 {
	return C.vdot_s32(v0, v1, v2)
}

// Floating-point fused Multiply-Add to accumulator (vector). This instruction multiplies corresponding floating-point values in the vectors in the two source SIMD&FP registers, adds the product to the corresponding vector element of the destination SIMD&FP register, and writes the result to the destination SIMD&FP register.
func VfmaqF32(v0 Float32X4, v1 Float32X4, v2 Float32X4) Float32X4 {
	return C.vfmaq_f32(v0, v1, v2)
}

// Floating-point fused Multiply-Add to accumulator (vector). This instruction multiplies corresponding floating-point values in the vectors in the two source SIMD&FP registers, adds the product to the corresponding vector element of the destination SIMD&FP register, and writes the result to the destination SIMD&FP register.
func VfmaF32(v0 Float32X2, v1 Float32X2, v2 Float32X2) Float32X2 {
	return C.vfma_f32(v0, v1, v2)
}

// Floating-point fused Multiply-Add to accumulator (vector). This instruction multiplies corresponding floating-point values in the vectors in the two source SIMD&FP registers, adds the product to the corresponding vector element of the destination SIMD&FP register, and writes the result to the destination SIMD&FP register.
func VfmaqNF32(v0 Float32X4, v1 Float32X4, v2 Float32) Float32X4 {
	return C.vfmaq_n_f32(v0, v1, v2)
}

// Floating-point fused Multiply-Add to accumulator (vector). This instruction multiplies corresponding floating-point values in the vectors in the two source SIMD&FP registers, adds the product to the corresponding vector element of the destination SIMD&FP register, and writes the result to the destination SIMD&FP register.
func VfmaNF32(v0 Float32X2, v1 Float32X2, v2 Float32) Float32X2 {
	return C.vfma_n_f32(v0, v1, v2)
}

// Floating-point fused Multiply-Subtract from accumulator (vector). This instruction multiplies corresponding floating-point values in the vectors in the two source SIMD&FP registers, negates the product, adds the result to the corresponding vector element of the destination SIMD&FP register, and writes the result to the destination SIMD&FP register.
func VfmsqF32(v0 Float32X4, v1 Float32X4, v2 Float32X4) Float32X4 {
	return C.vfmsq_f32(v0, v1, v2)
}

// Floating-point fused Multiply-Subtract from accumulator (vector). This instruction multiplies corresponding floating-point values in the vectors in the two source SIMD&FP registers, negates the product, adds the result to the corresponding vector element of the destination SIMD&FP register, and writes the result to the destination SIMD&FP register.
func VfmsF32(v0 Float32X2, v1 Float32X2, v2 Float32X2) Float32X2 {
	return C.vfms_f32(v0, v1, v2)
}

// Signed Saturating Rounding Doubling Multiply Accumulate returning High Half (vector). This instruction multiplies the vector elements of the first source SIMD&FP register with the corresponding vector elements of the second source SIMD&FP register without saturating the multiply results, doubles the results, and accumulates the most significant half of the final results with the vector elements of the destination SIMD&FP register. The results are rounded.
func VqrdmlahqS32(v0 Int32X4, v1 Int32X4, v2 Int32X4) Int32X4 {
	return C.vqrdmlahq_s32(v0, v1, v2)
}

// Signed Saturating Rounding Doubling Multiply Accumulate returning High Half (vector). This instruction multiplies the vector elements of the first source SIMD&FP register with the corresponding vector elements of the second source SIMD&FP register without saturating the multiply results, doubles the results, and accumulates the most significant half of the final results with the vector elements of the destination SIMD&FP register. The results are rounded.
func VqrdmlahqS16(v0 Int16X8, v1 Int16X8, v2 Int16X8) Int16X8 {
	return C.vqrdmlahq_s16(v0, v1, v2)
}

// Signed Saturating Rounding Doubling Multiply Accumulate returning High Half (vector). This instruction multiplies the vector elements of the first source SIMD&FP register with the corresponding vector elements of the second source SIMD&FP register without saturating the multiply results, doubles the results, and accumulates the most significant half of the final results with the vector elements of the destination SIMD&FP register. The results are rounded.
func VqrdmlahS32(v0 Int32X2, v1 Int32X2, v2 Int32X2) Int32X2 {
	return C.vqrdmlah_s32(v0, v1, v2)
}

// Signed Saturating Rounding Doubling Multiply Accumulate returning High Half (vector). This instruction multiplies the vector elements of the first source SIMD&FP register with the corresponding vector elements of the second source SIMD&FP register without saturating the multiply results, doubles the results, and accumulates the most significant half of the final results with the vector elements of the destination SIMD&FP register. The results are rounded.
func VqrdmlahS16(v0 Int16X4, v1 Int16X4, v2 Int16X4) Int16X4 {
	return C.vqrdmlah_s16(v0, v1, v2)
}

// Signed Saturating Rounding Doubling Multiply Subtract returning High Half (vector). This instruction multiplies the vector elements of the first source SIMD&FP register with the corresponding vector elements of the second source SIMD&FP register without saturating the multiply results, doubles the results, and subtracts the most significant half of the final results from the vector elements of the destination SIMD&FP register. The results are rounded.
func VqrdmlshqS32(v0 Int32X4, v1 Int32X4, v2 Int32X4) Int32X4 {
	return C.vqrdmlshq_s32(v0, v1, v2)
}

// Signed Saturating Rounding Doubling Multiply Subtract returning High Half (vector). This instruction multiplies the vector elements of the first source SIMD&FP register with the corresponding vector elements of the second source SIMD&FP register without saturating the multiply results, doubles the results, and subtracts the most significant half of the final results from the vector elements of the destination SIMD&FP register. The results are rounded.
func VqrdmlshqS16(v0 Int16X8, v1 Int16X8, v2 Int16X8) Int16X8 {
	return C.vqrdmlshq_s16(v0, v1, v2)
}

// Signed Saturating Rounding Doubling Multiply Subtract returning High Half (vector). This instruction multiplies the vector elements of the first source SIMD&FP register with the corresponding vector elements of the second source SIMD&FP register without saturating the multiply results, doubles the results, and subtracts the most significant half of the final results from the vector elements of the destination SIMD&FP register. The results are rounded.
func VqrdmlshS32(v0 Int32X2, v1 Int32X2, v2 Int32X2) Int32X2 {
	return C.vqrdmlsh_s32(v0, v1, v2)
}

// Signed Saturating Rounding Doubling Multiply Subtract returning High Half (vector). This instruction multiplies the vector elements of the first source SIMD&FP register with the corresponding vector elements of the second source SIMD&FP register without saturating the multiply results, doubles the results, and subtracts the most significant half of the final results from the vector elements of the destination SIMD&FP register. The results are rounded.
func VqrdmlshS16(v0 Int16X4, v1 Int16X4, v2 Int16X4) Int16X4 {
	return C.vqrdmlsh_s16(v0, v1, v2)
}

// Floating-point Absolute Difference (vector). This instruction subtracts the floating-point values in the elements of the second source SIMD&FP register, from the corresponding floating-point values in the elements of the first source SIMD&FP register, places the absolute value of each result in a vector, and writes the vector to the destination SIMD&FP register.
func VabdqF64(v0 Float64X2, v1 Float64X2) Float64X2 {
	return C.vabdq_f64(v0, v1)
}

// Floating-point Absolute Difference (vector). This instruction subtracts the floating-point values in the elements of the second source SIMD&FP register, from the corresponding floating-point values in the elements of the first source SIMD&FP register, places the absolute value of each result in a vector, and writes the vector to the destination SIMD&FP register.
func VabdF64(v0 Float64X1, v1 Float64X1) Float64X1 {
	return C.vabd_f64(v0, v1)
}

// Floating-point Absolute Difference (vector). This instruction subtracts the floating-point values in the elements of the second source SIMD&FP register, from the corresponding floating-point values in the elements of the first source SIMD&FP register, places the absolute value of each result in a vector, and writes the vector to the destination SIMD&FP register.
func VabddF64(v0 Float64, v1 Float64) Float64 {
	return C.vabdd_f64(v0, v1)
}

// Floating-point Absolute Difference (vector). This instruction subtracts the floating-point values in the elements of the second source SIMD&FP register, from the corresponding floating-point values in the elements of the first source SIMD&FP register, places the absolute value of each result in a vector, and writes the vector to the destination SIMD&FP register.
func VabdsF32(v0 Float32, v1 Float32) Float32 {
	return C.vabds_f32(v0, v1)
}

// Floating-point Absolute value (vector). This instruction calculates the absolute value of each vector element in the source SIMD&FP register, writes the result to a vector, and writes the vector to the destination SIMD&FP register.
func VabsqF64(v0 Float64X2) Float64X2 {
	return C.vabsq_f64(v0)
}

// Absolute value (vector). This instruction calculates the absolute value of each vector element in the source SIMD&FP register, puts the result into a vector, and writes the vector to the destination SIMD&FP register.
func VabsqS64(v0 Int64X2) Int64X2 {
	return C.vabsq_s64(v0)
}

// Floating-point Absolute value (vector). This instruction calculates the absolute value of each vector element in the source SIMD&FP register, writes the result to a vector, and writes the vector to the destination SIMD&FP register.
func VabsF64(v0 Float64X1) Float64X1 {
	return C.vabs_f64(v0)
}

// Absolute value (vector). This instruction calculates the absolute value of each vector element in the source SIMD&FP register, puts the result into a vector, and writes the vector to the destination SIMD&FP register.
func VabsS64(v0 Int64X1) Int64X1 {
	return C.vabs_s64(v0)
}

// Absolute value (vector). This instruction calculates the absolute value of each vector element in the source SIMD&FP register, puts the result into a vector, and writes the vector to the destination SIMD&FP register.
func VabsdS64(v0 Int64) Int64 {
	return C.vabsd_s64(v0)
}

// Floating-point Add (vector). This instruction adds corresponding vector elements in the two source SIMD&FP registers, writes the result into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are floating-point values.
func VaddqF64(v0 Float64X2, v1 Float64X2) Float64X2 {
	return C.vaddq_f64(v0, v1)
}

// Floating-point Add (vector). This instruction adds corresponding vector elements in the two source SIMD&FP registers, writes the result into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are floating-point values.
func VaddF64(v0 Float64X1, v1 Float64X1) Float64X1 {
	return C.vadd_f64(v0, v1)
}

// Add (vector). This instruction adds corresponding elements in the two source SIMD&FP registers, places the results into a vector, and writes the vector to the destination SIMD&FP register.
func VadddU64(v0 Uint64, v1 Uint64) Uint64 {
	return C.vaddd_u64(v0, v1)
}

// Add (vector). This instruction adds corresponding elements in the two source SIMD&FP registers, places the results into a vector, and writes the vector to the destination SIMD&FP register.
func VadddS64(v0 Int64, v1 Int64) Int64 {
	return C.vaddd_s64(v0, v1)
}

// Bitwise Exclusive OR (vector). This instruction performs a bitwise Exclusive OR operation between the two source SIMD&FP registers, and places the result in the destination SIMD&FP register.
func VaddqP128(v0 Poly128, v1 Poly128) Poly128 {
	return C.vaddq_p128(v0, v1)
}

// Add returning High Narrow. This instruction adds each vector element in the first source SIMD&FP register to the corresponding vector element in the second source SIMD&FP register, places the most significant half of the result into a vector, and writes the vector to the lower or upper half of the destination SIMD&FP register.
func VaddhnHighU32(v0 Uint16X4, v1 Uint32X4, v2 Uint32X4) Uint16X8 {
	return C.vaddhn_high_u32(v0, v1, v2)
}

// Add returning High Narrow. This instruction adds each vector element in the first source SIMD&FP register to the corresponding vector element in the second source SIMD&FP register, places the most significant half of the result into a vector, and writes the vector to the lower or upper half of the destination SIMD&FP register.
func VaddhnHighU64(v0 Uint32X2, v1 Uint64X2, v2 Uint64X2) Uint32X4 {
	return C.vaddhn_high_u64(v0, v1, v2)
}

// Add returning High Narrow. This instruction adds each vector element in the first source SIMD&FP register to the corresponding vector element in the second source SIMD&FP register, places the most significant half of the result into a vector, and writes the vector to the lower or upper half of the destination SIMD&FP register.
func VaddhnHighU16(v0 Uint8X8, v1 Uint16X8, v2 Uint16X8) Uint8X16 {
	return C.vaddhn_high_u16(v0, v1, v2)
}

// Add returning High Narrow. This instruction adds each vector element in the first source SIMD&FP register to the corresponding vector element in the second source SIMD&FP register, places the most significant half of the result into a vector, and writes the vector to the lower or upper half of the destination SIMD&FP register.
func VaddhnHighS32(v0 Int16X4, v1 Int32X4, v2 Int32X4) Int16X8 {
	return C.vaddhn_high_s32(v0, v1, v2)
}

// Add returning High Narrow. This instruction adds each vector element in the first source SIMD&FP register to the corresponding vector element in the second source SIMD&FP register, places the most significant half of the result into a vector, and writes the vector to the lower or upper half of the destination SIMD&FP register.
func VaddhnHighS64(v0 Int32X2, v1 Int64X2, v2 Int64X2) Int32X4 {
	return C.vaddhn_high_s64(v0, v1, v2)
}

// Add returning High Narrow. This instruction adds each vector element in the first source SIMD&FP register to the corresponding vector element in the second source SIMD&FP register, places the most significant half of the result into a vector, and writes the vector to the lower or upper half of the destination SIMD&FP register.
func VaddhnHighS16(v0 Int8X8, v1 Int16X8, v2 Int16X8) Int8X16 {
	return C.vaddhn_high_s16(v0, v1, v2)
}

// Unsigned sum Long across Vector. This instruction adds every vector element in the source SIMD&FP register together, and writes the scalar result to the destination SIMD&FP register. The destination scalar is twice as long as the source vector elements. All the values in this instruction are unsigned integer values.
func VaddlvqU8(v0 Uint8X16) Uint16 {
	return C.vaddlvq_u8(v0)
}

// Unsigned sum Long across Vector. This instruction adds every vector element in the source SIMD&FP register together, and writes the scalar result to the destination SIMD&FP register. The destination scalar is twice as long as the source vector elements. All the values in this instruction are unsigned integer values.
func VaddlvqU32(v0 Uint32X4) Uint64 {
	return C.vaddlvq_u32(v0)
}

// Unsigned sum Long across Vector. This instruction adds every vector element in the source SIMD&FP register together, and writes the scalar result to the destination SIMD&FP register. The destination scalar is twice as long as the source vector elements. All the values in this instruction are unsigned integer values.
func VaddlvqU16(v0 Uint16X8) Uint32 {
	return C.vaddlvq_u16(v0)
}

// Signed Add Long across Vector. This instruction adds every vector element in the source SIMD&FP register together, and writes the scalar result to the destination SIMD&FP register. The destination scalar is twice as long as the source vector elements. All the values in this instruction are signed integer values.
func VaddlvqS8(v0 Int8X16) Int16 {
	return C.vaddlvq_s8(v0)
}

// Signed Add Long across Vector. This instruction adds every vector element in the source SIMD&FP register together, and writes the scalar result to the destination SIMD&FP register. The destination scalar is twice as long as the source vector elements. All the values in this instruction are signed integer values.
func VaddlvqS32(v0 Int32X4) Int64 {
	return C.vaddlvq_s32(v0)
}

// Signed Add Long across Vector. This instruction adds every vector element in the source SIMD&FP register together, and writes the scalar result to the destination SIMD&FP register. The destination scalar is twice as long as the source vector elements. All the values in this instruction are signed integer values.
func VaddlvqS16(v0 Int16X8) Int32 {
	return C.vaddlvq_s16(v0)
}

// Unsigned sum Long across Vector. This instruction adds every vector element in the source SIMD&FP register together, and writes the scalar result to the destination SIMD&FP register. The destination scalar is twice as long as the source vector elements. All the values in this instruction are unsigned integer values.
func VaddlvU8(v0 Uint8X8) Uint16 {
	return C.vaddlv_u8(v0)
}

// Unsigned Add Long Pairwise. This instruction adds pairs of adjacent unsigned integer values from the vector in the source SIMD&FP register, places the result into a vector, and writes the vector to the destination SIMD&FP register. The destination vector elements are twice as long as the source vector elements.
func VaddlvU32(v0 Uint32X2) Uint64 {
	return C.vaddlv_u32(v0)
}

// Unsigned sum Long across Vector. This instruction adds every vector element in the source SIMD&FP register together, and writes the scalar result to the destination SIMD&FP register. The destination scalar is twice as long as the source vector elements. All the values in this instruction are unsigned integer values.
func VaddlvU16(v0 Uint16X4) Uint32 {
	return C.vaddlv_u16(v0)
}

// Signed Add Long across Vector. This instruction adds every vector element in the source SIMD&FP register together, and writes the scalar result to the destination SIMD&FP register. The destination scalar is twice as long as the source vector elements. All the values in this instruction are signed integer values.
func VaddlvS8(v0 Int8X8) Int16 {
	return C.vaddlv_s8(v0)
}

// Signed Add Long Pairwise. This instruction adds pairs of adjacent signed integer values from the vector in the source SIMD&FP register, places the result into a vector, and writes the vector to the destination SIMD&FP register. The destination vector elements are twice as long as the source vector elements.
func VaddlvS32(v0 Int32X2) Int64 {
	return C.vaddlv_s32(v0)
}

// Signed Add Long across Vector. This instruction adds every vector element in the source SIMD&FP register together, and writes the scalar result to the destination SIMD&FP register. The destination scalar is twice as long as the source vector elements. All the values in this instruction are signed integer values.
func VaddlvS16(v0 Int16X4) Int32 {
	return C.vaddlv_s16(v0)
}

// Add across Vector. This instruction adds every vector element in the source SIMD&FP register together, and writes the scalar result to the destination SIMD&FP register.
func VaddvqU8(v0 Uint8X16) Uint8 {
	return C.vaddvq_u8(v0)
}

// Add across Vector. This instruction adds every vector element in the source SIMD&FP register together, and writes the scalar result to the destination SIMD&FP register.
func VaddvqU32(v0 Uint32X4) Uint32 {
	return C.vaddvq_u32(v0)
}

// Add across vector
func VaddvqU64(v0 Uint64X2) Uint64 {
	return C.vaddvq_u64(v0)
}

// Add across Vector. This instruction adds every vector element in the source SIMD&FP register together, and writes the scalar result to the destination SIMD&FP register.
func VaddvqU16(v0 Uint16X8) Uint16 {
	return C.vaddvq_u16(v0)
}

// Add across Vector. This instruction adds every vector element in the source SIMD&FP register together, and writes the scalar result to the destination SIMD&FP register.
func VaddvqS8(v0 Int8X16) Int8 {
	return C.vaddvq_s8(v0)
}

// Floating-point add across vector
func VaddvqF64(v0 Float64X2) Float64 {
	return C.vaddvq_f64(v0)
}

// Floating-point add across vector
func VaddvqF32(v0 Float32X4) Float32 {
	return C.vaddvq_f32(v0)
}

// Add across Vector. This instruction adds every vector element in the source SIMD&FP register together, and writes the scalar result to the destination SIMD&FP register.
func VaddvqS32(v0 Int32X4) Int32 {
	return C.vaddvq_s32(v0)
}

// Add across vector
func VaddvqS64(v0 Int64X2) Int64 {
	return C.vaddvq_s64(v0)
}

// Add across Vector. This instruction adds every vector element in the source SIMD&FP register together, and writes the scalar result to the destination SIMD&FP register.
func VaddvqS16(v0 Int16X8) Int16 {
	return C.vaddvq_s16(v0)
}

// Add across Vector. This instruction adds every vector element in the source SIMD&FP register together, and writes the scalar result to the destination SIMD&FP register.
func VaddvU8(v0 Uint8X8) Uint8 {
	return C.vaddv_u8(v0)
}

// Add across vector
func VaddvU32(v0 Uint32X2) Uint32 {
	return C.vaddv_u32(v0)
}

// Add across Vector. This instruction adds every vector element in the source SIMD&FP register together, and writes the scalar result to the destination SIMD&FP register.
func VaddvU16(v0 Uint16X4) Uint16 {
	return C.vaddv_u16(v0)
}

// Add across Vector. This instruction adds every vector element in the source SIMD&FP register together, and writes the scalar result to the destination SIMD&FP register.
func VaddvS8(v0 Int8X8) Int8 {
	return C.vaddv_s8(v0)
}

// Floating-point add across vector
func VaddvF32(v0 Float32X2) Float32 {
	return C.vaddv_f32(v0)
}

// Add across vector
func VaddvS32(v0 Int32X2) Int32 {
	return C.vaddv_s32(v0)
}

// Add across Vector. This instruction adds every vector element in the source SIMD&FP register together, and writes the scalar result to the destination SIMD&FP register.
func VaddvS16(v0 Int16X4) Int16 {
	return C.vaddv_s16(v0)
}

// Bitwise Select. This instruction sets each bit in the destination SIMD&FP register to the corresponding bit from the first source SIMD&FP register when the original destination bit was 1, otherwise from the second source SIMD&FP register.
func VbslP64(v0 Uint64X1, v1 Poly64X1, v2 Poly64X1) Poly64X1 {
	return C.vbsl_p64(v0, v1, v2)
}

// Bitwise Select. This instruction sets each bit in the destination SIMD&FP register to the corresponding bit from the first source SIMD&FP register when the original destination bit was 1, otherwise from the second source SIMD&FP register.
func VbslqP64(v0 Uint64X2, v1 Poly64X2, v2 Poly64X2) Poly64X2 {
	return C.vbslq_p64(v0, v1, v2)
}

// Bitwise Select. This instruction sets each bit in the destination SIMD&FP register to the corresponding bit from the first source SIMD&FP register when the original destination bit was 1, otherwise from the second source SIMD&FP register.
func VbslqF64(v0 Uint64X2, v1 Float64X2, v2 Float64X2) Float64X2 {
	return C.vbslq_f64(v0, v1, v2)
}

// Bitwise Select. This instruction sets each bit in the destination SIMD&FP register to the corresponding bit from the first source SIMD&FP register when the original destination bit was 1, otherwise from the second source SIMD&FP register.
func VbslF64(v0 Uint64X1, v1 Float64X1, v2 Float64X1) Float64X1 {
	return C.vbsl_f64(v0, v1, v2)
}

// Floating-point Absolute Compare Greater than or Equal (vector). This instruction compares the absolute value of each floating-point value in the first source SIMD&FP register with the absolute value of the corresponding floating-point value in the second source SIMD&FP register and if the first value is greater than or equal to the second value sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
func VcageqF64(v0 Float64X2, v1 Float64X2) Uint64X2 {
	return C.vcageq_f64(v0, v1)
}

// Floating-point Absolute Compare Greater than or Equal (vector). This instruction compares the absolute value of each floating-point value in the first source SIMD&FP register with the absolute value of the corresponding floating-point value in the second source SIMD&FP register and if the first value is greater than or equal to the second value sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
func VcageF64(v0 Float64X1, v1 Float64X1) Uint64X1 {
	return C.vcage_f64(v0, v1)
}

// Floating-point Absolute Compare Greater than or Equal (vector). This instruction compares the absolute value of each floating-point value in the first source SIMD&FP register with the absolute value of the corresponding floating-point value in the second source SIMD&FP register and if the first value is greater than or equal to the second value sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
func VcagedF64(v0 Float64, v1 Float64) Uint64 {
	return C.vcaged_f64(v0, v1)
}

// Floating-point Absolute Compare Greater than or Equal (vector). This instruction compares the absolute value of each floating-point value in the first source SIMD&FP register with the absolute value of the corresponding floating-point value in the second source SIMD&FP register and if the first value is greater than or equal to the second value sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
func VcagesF32(v0 Float32, v1 Float32) Uint32 {
	return C.vcages_f32(v0, v1)
}

// Floating-point Absolute Compare Greater than (vector). This instruction compares the absolute value of each vector element in the first source SIMD&FP register with the absolute value of the corresponding vector element in the second source SIMD&FP register and if the first value is greater than the second value sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
func VcagtqF64(v0 Float64X2, v1 Float64X2) Uint64X2 {
	return C.vcagtq_f64(v0, v1)
}

// Floating-point Absolute Compare Greater than (vector). This instruction compares the absolute value of each vector element in the first source SIMD&FP register with the absolute value of the corresponding vector element in the second source SIMD&FP register and if the first value is greater than the second value sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
func VcagtF64(v0 Float64X1, v1 Float64X1) Uint64X1 {
	return C.vcagt_f64(v0, v1)
}

// Floating-point Absolute Compare Greater than (vector). This instruction compares the absolute value of each vector element in the first source SIMD&FP register with the absolute value of the corresponding vector element in the second source SIMD&FP register and if the first value is greater than the second value sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
func VcagtdF64(v0 Float64, v1 Float64) Uint64 {
	return C.vcagtd_f64(v0, v1)
}

// Floating-point Absolute Compare Greater than (vector). This instruction compares the absolute value of each vector element in the first source SIMD&FP register with the absolute value of the corresponding vector element in the second source SIMD&FP register and if the first value is greater than the second value sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
func VcagtsF32(v0 Float32, v1 Float32) Uint32 {
	return C.vcagts_f32(v0, v1)
}

// Floating-point absolute compare less than or equal
func VcaleqF64(v0 Float64X2, v1 Float64X2) Uint64X2 {
	return C.vcaleq_f64(v0, v1)
}

// Floating-point absolute compare less than or equal
func VcaleF64(v0 Float64X1, v1 Float64X1) Uint64X1 {
	return C.vcale_f64(v0, v1)
}

// Floating-point absolute compare less than or equal
func VcaledF64(v0 Float64, v1 Float64) Uint64 {
	return C.vcaled_f64(v0, v1)
}

// Floating-point absolute compare less than or equal
func VcalesF32(v0 Float32, v1 Float32) Uint32 {
	return C.vcales_f32(v0, v1)
}

// Floating-point absolute compare less than
func VcaltqF64(v0 Float64X2, v1 Float64X2) Uint64X2 {
	return C.vcaltq_f64(v0, v1)
}

// Floating-point absolute compare less than
func VcaltF64(v0 Float64X1, v1 Float64X1) Uint64X1 {
	return C.vcalt_f64(v0, v1)
}

// Floating-point absolute compare less than
func VcaltdF64(v0 Float64, v1 Float64) Uint64 {
	return C.vcaltd_f64(v0, v1)
}

// Floating-point absolute compare less than
func VcaltsF32(v0 Float32, v1 Float32) Uint32 {
	return C.vcalts_f32(v0, v1)
}

// Compare bitwise Equal (vector). This instruction compares each vector element from the first source SIMD&FP register with the corresponding vector element from the second source SIMD&FP register, and if the comparison is equal sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
func VceqP64(v0 Poly64X1, v1 Poly64X1) Uint64X1 {
	return C.vceq_p64(v0, v1)
}

// Compare bitwise Equal (vector). This instruction compares each vector element from the first source SIMD&FP register with the corresponding vector element from the second source SIMD&FP register, and if the comparison is equal sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
func VceqqP64(v0 Poly64X2, v1 Poly64X2) Uint64X2 {
	return C.vceqq_p64(v0, v1)
}

// Compare bitwise Equal (vector). This instruction compares each vector element from the first source SIMD&FP register with the corresponding vector element from the second source SIMD&FP register, and if the comparison is equal sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
func VceqqU64(v0 Uint64X2, v1 Uint64X2) Uint64X2 {
	return C.vceqq_u64(v0, v1)
}

// Floating-point Compare Equal (vector). This instruction compares each floating-point value from the first source SIMD&FP register, with the corresponding floating-point value from the second source SIMD&FP register, and if the comparison is equal sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
func VceqqF64(v0 Float64X2, v1 Float64X2) Uint64X2 {
	return C.vceqq_f64(v0, v1)
}

// Compare bitwise Equal (vector). This instruction compares each vector element from the first source SIMD&FP register with the corresponding vector element from the second source SIMD&FP register, and if the comparison is equal sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
func VceqqS64(v0 Int64X2, v1 Int64X2) Uint64X2 {
	return C.vceqq_s64(v0, v1)
}

// Compare bitwise Equal (vector). This instruction compares each vector element from the first source SIMD&FP register with the corresponding vector element from the second source SIMD&FP register, and if the comparison is equal sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
func VceqU64(v0 Uint64X1, v1 Uint64X1) Uint64X1 {
	return C.vceq_u64(v0, v1)
}

// Floating-point Compare Equal (vector). This instruction compares each floating-point value from the first source SIMD&FP register, with the corresponding floating-point value from the second source SIMD&FP register, and if the comparison is equal sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
func VceqF64(v0 Float64X1, v1 Float64X1) Uint64X1 {
	return C.vceq_f64(v0, v1)
}

// Compare bitwise Equal (vector). This instruction compares each vector element from the first source SIMD&FP register with the corresponding vector element from the second source SIMD&FP register, and if the comparison is equal sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
func VceqS64(v0 Int64X1, v1 Int64X1) Uint64X1 {
	return C.vceq_s64(v0, v1)
}

// Compare bitwise Equal (vector). This instruction compares each vector element from the first source SIMD&FP register with the corresponding vector element from the second source SIMD&FP register, and if the comparison is equal sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
func VceqdU64(v0 Uint64, v1 Uint64) Uint64 {
	return C.vceqd_u64(v0, v1)
}

// Compare bitwise Equal (vector). This instruction compares each vector element from the first source SIMD&FP register with the corresponding vector element from the second source SIMD&FP register, and if the comparison is equal sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
func VceqdS64(v0 Int64, v1 Int64) Uint64 {
	return C.vceqd_s64(v0, v1)
}

// Floating-point Compare Equal (vector). This instruction compares each floating-point value from the first source SIMD&FP register, with the corresponding floating-point value from the second source SIMD&FP register, and if the comparison is equal sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
func VceqdF64(v0 Float64, v1 Float64) Uint64 {
	return C.vceqd_f64(v0, v1)
}

// Floating-point Compare Equal (vector). This instruction compares each floating-point value from the first source SIMD&FP register, with the corresponding floating-point value from the second source SIMD&FP register, and if the comparison is equal sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
func VceqsF32(v0 Float32, v1 Float32) Uint32 {
	return C.vceqs_f32(v0, v1)
}

// Compare bitwise Equal to zero (vector). This instruction reads each vector element in the source SIMD&FP register and if the value is equal to zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
func VceqzP8(v0 Poly8X8) Uint8X8 {
	return C.vceqz_p8(v0)
}

// Compare bitwise Equal to zero (vector). This instruction reads each vector element in the source SIMD&FP register and if the value is equal to zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
func VceqzP64(v0 Poly64X1) Uint64X1 {
	return C.vceqz_p64(v0)
}

// Compare bitwise Equal to zero (vector). This instruction reads each vector element in the source SIMD&FP register and if the value is equal to zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
func VceqzqP8(v0 Poly8X16) Uint8X16 {
	return C.vceqzq_p8(v0)
}

// Compare bitwise Equal to zero (vector). This instruction reads each vector element in the source SIMD&FP register and if the value is equal to zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
func VceqzqP64(v0 Poly64X2) Uint64X2 {
	return C.vceqzq_p64(v0)
}

// Compare bitwise Equal to zero (vector). This instruction reads each vector element in the source SIMD&FP register and if the value is equal to zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
func VceqzqU8(v0 Uint8X16) Uint8X16 {
	return C.vceqzq_u8(v0)
}

// Compare bitwise Equal to zero (vector). This instruction reads each vector element in the source SIMD&FP register and if the value is equal to zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
func VceqzqU32(v0 Uint32X4) Uint32X4 {
	return C.vceqzq_u32(v0)
}

// Compare bitwise Equal to zero (vector). This instruction reads each vector element in the source SIMD&FP register and if the value is equal to zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
func VceqzqU64(v0 Uint64X2) Uint64X2 {
	return C.vceqzq_u64(v0)
}

// Compare bitwise Equal to zero (vector). This instruction reads each vector element in the source SIMD&FP register and if the value is equal to zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
func VceqzqU16(v0 Uint16X8) Uint16X8 {
	return C.vceqzq_u16(v0)
}

// Compare bitwise Equal to zero (vector). This instruction reads each vector element in the source SIMD&FP register and if the value is equal to zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
func VceqzqS8(v0 Int8X16) Uint8X16 {
	return C.vceqzq_s8(v0)
}

// Floating-point Compare Equal to zero (vector). This instruction reads each floating-point value in the source SIMD&FP register and if the value is equal to zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
func VceqzqF64(v0 Float64X2) Uint64X2 {
	return C.vceqzq_f64(v0)
}

// Floating-point Compare Equal to zero (vector). This instruction reads each floating-point value in the source SIMD&FP register and if the value is equal to zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
func VceqzqF32(v0 Float32X4) Uint32X4 {
	return C.vceqzq_f32(v0)
}

// Compare bitwise Equal to zero (vector). This instruction reads each vector element in the source SIMD&FP register and if the value is equal to zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
func VceqzqS32(v0 Int32X4) Uint32X4 {
	return C.vceqzq_s32(v0)
}

// Compare bitwise Equal to zero (vector). This instruction reads each vector element in the source SIMD&FP register and if the value is equal to zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
func VceqzqS64(v0 Int64X2) Uint64X2 {
	return C.vceqzq_s64(v0)
}

// Compare bitwise Equal to zero (vector). This instruction reads each vector element in the source SIMD&FP register and if the value is equal to zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
func VceqzqS16(v0 Int16X8) Uint16X8 {
	return C.vceqzq_s16(v0)
}

// Compare bitwise Equal to zero (vector). This instruction reads each vector element in the source SIMD&FP register and if the value is equal to zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
func VceqzU8(v0 Uint8X8) Uint8X8 {
	return C.vceqz_u8(v0)
}

// Compare bitwise Equal to zero (vector). This instruction reads each vector element in the source SIMD&FP register and if the value is equal to zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
func VceqzU32(v0 Uint32X2) Uint32X2 {
	return C.vceqz_u32(v0)
}

// Compare bitwise Equal to zero (vector). This instruction reads each vector element in the source SIMD&FP register and if the value is equal to zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
func VceqzU64(v0 Uint64X1) Uint64X1 {
	return C.vceqz_u64(v0)
}

// Compare bitwise Equal to zero (vector). This instruction reads each vector element in the source SIMD&FP register and if the value is equal to zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
func VceqzU16(v0 Uint16X4) Uint16X4 {
	return C.vceqz_u16(v0)
}

// Compare bitwise Equal to zero (vector). This instruction reads each vector element in the source SIMD&FP register and if the value is equal to zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
func VceqzS8(v0 Int8X8) Uint8X8 {
	return C.vceqz_s8(v0)
}

// Floating-point Compare Equal to zero (vector). This instruction reads each floating-point value in the source SIMD&FP register and if the value is equal to zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
func VceqzF64(v0 Float64X1) Uint64X1 {
	return C.vceqz_f64(v0)
}

// Floating-point Compare Equal to zero (vector). This instruction reads each floating-point value in the source SIMD&FP register and if the value is equal to zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
func VceqzF32(v0 Float32X2) Uint32X2 {
	return C.vceqz_f32(v0)
}

// Compare bitwise Equal to zero (vector). This instruction reads each vector element in the source SIMD&FP register and if the value is equal to zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
func VceqzS32(v0 Int32X2) Uint32X2 {
	return C.vceqz_s32(v0)
}

// Compare bitwise Equal to zero (vector). This instruction reads each vector element in the source SIMD&FP register and if the value is equal to zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
func VceqzS64(v0 Int64X1) Uint64X1 {
	return C.vceqz_s64(v0)
}

// Compare bitwise Equal to zero (vector). This instruction reads each vector element in the source SIMD&FP register and if the value is equal to zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
func VceqzS16(v0 Int16X4) Uint16X4 {
	return C.vceqz_s16(v0)
}

// Compare bitwise Equal to zero (vector). This instruction reads each vector element in the source SIMD&FP register and if the value is equal to zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
func VceqzdU64(v0 Uint64) Uint64 {
	return C.vceqzd_u64(v0)
}

// Compare bitwise Equal to zero (vector). This instruction reads each vector element in the source SIMD&FP register and if the value is equal to zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
func VceqzdS64(v0 Int64) Uint64 {
	return C.vceqzd_s64(v0)
}

// Floating-point Compare Equal to zero (vector). This instruction reads each floating-point value in the source SIMD&FP register and if the value is equal to zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
func VceqzdF64(v0 Float64) Uint64 {
	return C.vceqzd_f64(v0)
}

// Floating-point Compare Equal to zero (vector). This instruction reads each floating-point value in the source SIMD&FP register and if the value is equal to zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
func VceqzsF32(v0 Float32) Uint32 {
	return C.vceqzs_f32(v0)
}

// Compare unsigned Higher or Same (vector). This instruction compares each vector element in the first source SIMD&FP register with the corresponding vector element in the second source SIMD&FP register and if the first unsigned integer value is greater than or equal to the second unsigned integer value sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
func VcgeqU64(v0 Uint64X2, v1 Uint64X2) Uint64X2 {
	return C.vcgeq_u64(v0, v1)
}

// Floating-point Compare Greater than or Equal (vector). This instruction reads each floating-point value in the first source SIMD&FP register and if the value is greater than or equal to the corresponding floating-point value in the second source SIMD&FP register sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
func VcgeqF64(v0 Float64X2, v1 Float64X2) Uint64X2 {
	return C.vcgeq_f64(v0, v1)
}

// Compare signed Greater than or Equal (vector). This instruction compares each vector element in the first source SIMD&FP register with the corresponding vector element in the second source SIMD&FP register and if the first signed integer value is greater than or equal to the second signed integer value sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
func VcgeqS64(v0 Int64X2, v1 Int64X2) Uint64X2 {
	return C.vcgeq_s64(v0, v1)
}

// Compare unsigned Higher or Same (vector). This instruction compares each vector element in the first source SIMD&FP register with the corresponding vector element in the second source SIMD&FP register and if the first unsigned integer value is greater than or equal to the second unsigned integer value sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
func VcgeU64(v0 Uint64X1, v1 Uint64X1) Uint64X1 {
	return C.vcge_u64(v0, v1)
}

// Floating-point Compare Greater than or Equal (vector). This instruction reads each floating-point value in the first source SIMD&FP register and if the value is greater than or equal to the corresponding floating-point value in the second source SIMD&FP register sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
func VcgeF64(v0 Float64X1, v1 Float64X1) Uint64X1 {
	return C.vcge_f64(v0, v1)
}

// Compare signed Greater than or Equal (vector). This instruction compares each vector element in the first source SIMD&FP register with the corresponding vector element in the second source SIMD&FP register and if the first signed integer value is greater than or equal to the second signed integer value sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
func VcgeS64(v0 Int64X1, v1 Int64X1) Uint64X1 {
	return C.vcge_s64(v0, v1)
}

// Compare signed Greater than or Equal (vector). This instruction compares each vector element in the first source SIMD&FP register with the corresponding vector element in the second source SIMD&FP register and if the first signed integer value is greater than or equal to the second signed integer value sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
func VcgedS64(v0 Int64, v1 Int64) Uint64 {
	return C.vcged_s64(v0, v1)
}

// Compare unsigned Higher or Same (vector). This instruction compares each vector element in the first source SIMD&FP register with the corresponding vector element in the second source SIMD&FP register and if the first unsigned integer value is greater than or equal to the second unsigned integer value sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
func VcgedU64(v0 Uint64, v1 Uint64) Uint64 {
	return C.vcged_u64(v0, v1)
}

// Floating-point Compare Greater than or Equal (vector). This instruction reads each floating-point value in the first source SIMD&FP register and if the value is greater than or equal to the corresponding floating-point value in the second source SIMD&FP register sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
func VcgedF64(v0 Float64, v1 Float64) Uint64 {
	return C.vcged_f64(v0, v1)
}

// Floating-point Compare Greater than or Equal (vector). This instruction reads each floating-point value in the first source SIMD&FP register and if the value is greater than or equal to the corresponding floating-point value in the second source SIMD&FP register sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
func VcgesF32(v0 Float32, v1 Float32) Uint32 {
	return C.vcges_f32(v0, v1)
}

// Compare signed Greater than or Equal to zero (vector). This instruction reads each vector element in the source SIMD&FP register and if the signed integer value is greater than or equal to zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
func VcgezqS8(v0 Int8X16) Uint8X16 {
	return C.vcgezq_s8(v0)
}

// Floating-point Compare Greater than or Equal to zero (vector). This instruction reads each floating-point value in the source SIMD&FP register and if the value is greater than or equal to zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
func VcgezqF64(v0 Float64X2) Uint64X2 {
	return C.vcgezq_f64(v0)
}

// Floating-point Compare Greater than or Equal to zero (vector). This instruction reads each floating-point value in the source SIMD&FP register and if the value is greater than or equal to zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
func VcgezqF32(v0 Float32X4) Uint32X4 {
	return C.vcgezq_f32(v0)
}

// Compare signed Greater than or Equal to zero (vector). This instruction reads each vector element in the source SIMD&FP register and if the signed integer value is greater than or equal to zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
func VcgezqS32(v0 Int32X4) Uint32X4 {
	return C.vcgezq_s32(v0)
}

// Compare signed Greater than or Equal to zero (vector). This instruction reads each vector element in the source SIMD&FP register and if the signed integer value is greater than or equal to zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
func VcgezqS64(v0 Int64X2) Uint64X2 {
	return C.vcgezq_s64(v0)
}

// Compare signed Greater than or Equal to zero (vector). This instruction reads each vector element in the source SIMD&FP register and if the signed integer value is greater than or equal to zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
func VcgezqS16(v0 Int16X8) Uint16X8 {
	return C.vcgezq_s16(v0)
}

// Compare signed Greater than or Equal to zero (vector). This instruction reads each vector element in the source SIMD&FP register and if the signed integer value is greater than or equal to zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
func VcgezS8(v0 Int8X8) Uint8X8 {
	return C.vcgez_s8(v0)
}

// Floating-point Compare Greater than or Equal to zero (vector). This instruction reads each floating-point value in the source SIMD&FP register and if the value is greater than or equal to zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
func VcgezF64(v0 Float64X1) Uint64X1 {
	return C.vcgez_f64(v0)
}

// Floating-point Compare Greater than or Equal to zero (vector). This instruction reads each floating-point value in the source SIMD&FP register and if the value is greater than or equal to zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
func VcgezF32(v0 Float32X2) Uint32X2 {
	return C.vcgez_f32(v0)
}

// Compare signed Greater than or Equal to zero (vector). This instruction reads each vector element in the source SIMD&FP register and if the signed integer value is greater than or equal to zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
func VcgezS32(v0 Int32X2) Uint32X2 {
	return C.vcgez_s32(v0)
}

// Compare signed Greater than or Equal to zero (vector). This instruction reads each vector element in the source SIMD&FP register and if the signed integer value is greater than or equal to zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
func VcgezS64(v0 Int64X1) Uint64X1 {
	return C.vcgez_s64(v0)
}

// Compare signed Greater than or Equal to zero (vector). This instruction reads each vector element in the source SIMD&FP register and if the signed integer value is greater than or equal to zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
func VcgezS16(v0 Int16X4) Uint16X4 {
	return C.vcgez_s16(v0)
}

// Compare signed Greater than or Equal to zero (vector). This instruction reads each vector element in the source SIMD&FP register and if the signed integer value is greater than or equal to zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
func VcgezdS64(v0 Int64) Uint64 {
	return C.vcgezd_s64(v0)
}

// Floating-point Compare Greater than or Equal to zero (vector). This instruction reads each floating-point value in the source SIMD&FP register and if the value is greater than or equal to zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
func VcgezdF64(v0 Float64) Uint64 {
	return C.vcgezd_f64(v0)
}

// Floating-point Compare Greater than or Equal to zero (vector). This instruction reads each floating-point value in the source SIMD&FP register and if the value is greater than or equal to zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
func VcgezsF32(v0 Float32) Uint32 {
	return C.vcgezs_f32(v0)
}

// Compare unsigned Higher (vector). This instruction compares each vector element in the first source SIMD&FP register with the corresponding vector element in the second source SIMD&FP register and if the first unsigned integer value is greater than the second unsigned integer value sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
func VcgtqU64(v0 Uint64X2, v1 Uint64X2) Uint64X2 {
	return C.vcgtq_u64(v0, v1)
}

// Floating-point Compare Greater than (vector). This instruction reads each floating-point value in the first source SIMD&FP register and if the value is greater than the corresponding floating-point value in the second source SIMD&FP register sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
func VcgtqF64(v0 Float64X2, v1 Float64X2) Uint64X2 {
	return C.vcgtq_f64(v0, v1)
}

// Compare signed Greater than (vector). This instruction compares each vector element in the first source SIMD&FP register with the corresponding vector element in the second source SIMD&FP register and if the first signed integer value is greater than the second signed integer value sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
func VcgtqS64(v0 Int64X2, v1 Int64X2) Uint64X2 {
	return C.vcgtq_s64(v0, v1)
}

// Compare unsigned Higher (vector). This instruction compares each vector element in the first source SIMD&FP register with the corresponding vector element in the second source SIMD&FP register and if the first unsigned integer value is greater than the second unsigned integer value sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
func VcgtU64(v0 Uint64X1, v1 Uint64X1) Uint64X1 {
	return C.vcgt_u64(v0, v1)
}

// Floating-point Compare Greater than (vector). This instruction reads each floating-point value in the first source SIMD&FP register and if the value is greater than the corresponding floating-point value in the second source SIMD&FP register sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
func VcgtF64(v0 Float64X1, v1 Float64X1) Uint64X1 {
	return C.vcgt_f64(v0, v1)
}

// Compare signed Greater than (vector). This instruction compares each vector element in the first source SIMD&FP register with the corresponding vector element in the second source SIMD&FP register and if the first signed integer value is greater than the second signed integer value sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
func VcgtS64(v0 Int64X1, v1 Int64X1) Uint64X1 {
	return C.vcgt_s64(v0, v1)
}

// Compare signed Greater than (vector). This instruction compares each vector element in the first source SIMD&FP register with the corresponding vector element in the second source SIMD&FP register and if the first signed integer value is greater than the second signed integer value sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
func VcgtdS64(v0 Int64, v1 Int64) Uint64 {
	return C.vcgtd_s64(v0, v1)
}

// Compare unsigned Higher (vector). This instruction compares each vector element in the first source SIMD&FP register with the corresponding vector element in the second source SIMD&FP register and if the first unsigned integer value is greater than the second unsigned integer value sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
func VcgtdU64(v0 Uint64, v1 Uint64) Uint64 {
	return C.vcgtd_u64(v0, v1)
}

// Floating-point Compare Greater than (vector). This instruction reads each floating-point value in the first source SIMD&FP register and if the value is greater than the corresponding floating-point value in the second source SIMD&FP register sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
func VcgtdF64(v0 Float64, v1 Float64) Uint64 {
	return C.vcgtd_f64(v0, v1)
}

// Floating-point Compare Greater than (vector). This instruction reads each floating-point value in the first source SIMD&FP register and if the value is greater than the corresponding floating-point value in the second source SIMD&FP register sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
func VcgtsF32(v0 Float32, v1 Float32) Uint32 {
	return C.vcgts_f32(v0, v1)
}

// Compare signed Greater than zero (vector). This instruction reads each vector element in the source SIMD&FP register and if the signed integer value is greater than zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
func VcgtzqS8(v0 Int8X16) Uint8X16 {
	return C.vcgtzq_s8(v0)
}

// Floating-point Compare Greater than zero (vector). This instruction reads each floating-point value in the source SIMD&FP register and if the value is greater than zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
func VcgtzqF64(v0 Float64X2) Uint64X2 {
	return C.vcgtzq_f64(v0)
}

// Floating-point Compare Greater than zero (vector). This instruction reads each floating-point value in the source SIMD&FP register and if the value is greater than zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
func VcgtzqF32(v0 Float32X4) Uint32X4 {
	return C.vcgtzq_f32(v0)
}

// Compare signed Greater than zero (vector). This instruction reads each vector element in the source SIMD&FP register and if the signed integer value is greater than zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
func VcgtzqS32(v0 Int32X4) Uint32X4 {
	return C.vcgtzq_s32(v0)
}

// Compare signed Greater than zero (vector). This instruction reads each vector element in the source SIMD&FP register and if the signed integer value is greater than zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
func VcgtzqS64(v0 Int64X2) Uint64X2 {
	return C.vcgtzq_s64(v0)
}

// Compare signed Greater than zero (vector). This instruction reads each vector element in the source SIMD&FP register and if the signed integer value is greater than zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
func VcgtzqS16(v0 Int16X8) Uint16X8 {
	return C.vcgtzq_s16(v0)
}

// Compare signed Greater than zero (vector). This instruction reads each vector element in the source SIMD&FP register and if the signed integer value is greater than zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
func VcgtzS8(v0 Int8X8) Uint8X8 {
	return C.vcgtz_s8(v0)
}

// Floating-point Compare Greater than zero (vector). This instruction reads each floating-point value in the source SIMD&FP register and if the value is greater than zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
func VcgtzF64(v0 Float64X1) Uint64X1 {
	return C.vcgtz_f64(v0)
}

// Floating-point Compare Greater than zero (vector). This instruction reads each floating-point value in the source SIMD&FP register and if the value is greater than zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
func VcgtzF32(v0 Float32X2) Uint32X2 {
	return C.vcgtz_f32(v0)
}

// Compare signed Greater than zero (vector). This instruction reads each vector element in the source SIMD&FP register and if the signed integer value is greater than zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
func VcgtzS32(v0 Int32X2) Uint32X2 {
	return C.vcgtz_s32(v0)
}

// Compare signed Greater than zero (vector). This instruction reads each vector element in the source SIMD&FP register and if the signed integer value is greater than zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
func VcgtzS64(v0 Int64X1) Uint64X1 {
	return C.vcgtz_s64(v0)
}

// Compare signed Greater than zero (vector). This instruction reads each vector element in the source SIMD&FP register and if the signed integer value is greater than zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
func VcgtzS16(v0 Int16X4) Uint16X4 {
	return C.vcgtz_s16(v0)
}

// Compare signed Greater than zero (vector). This instruction reads each vector element in the source SIMD&FP register and if the signed integer value is greater than zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
func VcgtzdS64(v0 Int64) Uint64 {
	return C.vcgtzd_s64(v0)
}

// Floating-point Compare Greater than zero (vector). This instruction reads each floating-point value in the source SIMD&FP register and if the value is greater than zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
func VcgtzdF64(v0 Float64) Uint64 {
	return C.vcgtzd_f64(v0)
}

// Floating-point Compare Greater than zero (vector). This instruction reads each floating-point value in the source SIMD&FP register and if the value is greater than zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
func VcgtzsF32(v0 Float32) Uint32 {
	return C.vcgtzs_f32(v0)
}

// Compare unsigned less than or equal
func VcleqU64(v0 Uint64X2, v1 Uint64X2) Uint64X2 {
	return C.vcleq_u64(v0, v1)
}

// Floating-point compare less than or equal
func VcleqF64(v0 Float64X2, v1 Float64X2) Uint64X2 {
	return C.vcleq_f64(v0, v1)
}

// Compare signed less than or equal
func VcleqS64(v0 Int64X2, v1 Int64X2) Uint64X2 {
	return C.vcleq_s64(v0, v1)
}

// Compare unsigned less than or equal
func VcleU64(v0 Uint64X1, v1 Uint64X1) Uint64X1 {
	return C.vcle_u64(v0, v1)
}

// Floating-point compare less than or equal
func VcleF64(v0 Float64X1, v1 Float64X1) Uint64X1 {
	return C.vcle_f64(v0, v1)
}

// Compare signed less than or equal
func VcleS64(v0 Int64X1, v1 Int64X1) Uint64X1 {
	return C.vcle_s64(v0, v1)
}

// Compare unsigned less than or equal
func VcledU64(v0 Uint64, v1 Uint64) Uint64 {
	return C.vcled_u64(v0, v1)
}

// Compare signed less than or equal
func VcledS64(v0 Int64, v1 Int64) Uint64 {
	return C.vcled_s64(v0, v1)
}

// Floating-point compare less than or equal
func VcledF64(v0 Float64, v1 Float64) Uint64 {
	return C.vcled_f64(v0, v1)
}

// Floating-point compare less than or equal
func VclesF32(v0 Float32, v1 Float32) Uint32 {
	return C.vcles_f32(v0, v1)
}

// Compare signed Less than or Equal to zero (vector). This instruction reads each vector element in the source SIMD&FP register and if the signed integer value is less than or equal to zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
func VclezqS8(v0 Int8X16) Uint8X16 {
	return C.vclezq_s8(v0)
}

// Floating-point Compare Less than or Equal to zero (vector). This instruction reads each floating-point value in the source SIMD&FP register and if the value is less than or equal to zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
func VclezqF64(v0 Float64X2) Uint64X2 {
	return C.vclezq_f64(v0)
}

// Floating-point Compare Less than or Equal to zero (vector). This instruction reads each floating-point value in the source SIMD&FP register and if the value is less than or equal to zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
func VclezqF32(v0 Float32X4) Uint32X4 {
	return C.vclezq_f32(v0)
}

// Compare signed Less than or Equal to zero (vector). This instruction reads each vector element in the source SIMD&FP register and if the signed integer value is less than or equal to zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
func VclezqS32(v0 Int32X4) Uint32X4 {
	return C.vclezq_s32(v0)
}

// Compare signed Less than or Equal to zero (vector). This instruction reads each vector element in the source SIMD&FP register and if the signed integer value is less than or equal to zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
func VclezqS64(v0 Int64X2) Uint64X2 {
	return C.vclezq_s64(v0)
}

// Compare signed Less than or Equal to zero (vector). This instruction reads each vector element in the source SIMD&FP register and if the signed integer value is less than or equal to zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
func VclezqS16(v0 Int16X8) Uint16X8 {
	return C.vclezq_s16(v0)
}

// Compare signed Less than or Equal to zero (vector). This instruction reads each vector element in the source SIMD&FP register and if the signed integer value is less than or equal to zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
func VclezS8(v0 Int8X8) Uint8X8 {
	return C.vclez_s8(v0)
}

// Floating-point Compare Less than or Equal to zero (vector). This instruction reads each floating-point value in the source SIMD&FP register and if the value is less than or equal to zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
func VclezF64(v0 Float64X1) Uint64X1 {
	return C.vclez_f64(v0)
}

// Compare signed Less than or Equal to zero (vector). This instruction reads each vector element in the source SIMD&FP register and if the signed integer value is less than or equal to zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
func VclezF32(v0 Float32X2) Uint32X2 {
	return C.vclez_f32(v0)
}

// Compare signed Less than or Equal to zero (vector). This instruction reads each vector element in the source SIMD&FP register and if the signed integer value is less than or equal to zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
func VclezS32(v0 Int32X2) Uint32X2 {
	return C.vclez_s32(v0)
}

// Compare signed Less than or Equal to zero (vector). This instruction reads each vector element in the source SIMD&FP register and if the signed integer value is less than or equal to zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
func VclezS64(v0 Int64X1) Uint64X1 {
	return C.vclez_s64(v0)
}

// Compare signed Less than or Equal to zero (vector). This instruction reads each vector element in the source SIMD&FP register and if the signed integer value is less than or equal to zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
func VclezS16(v0 Int16X4) Uint16X4 {
	return C.vclez_s16(v0)
}

// Compare signed Less than or Equal to zero (vector). This instruction reads each vector element in the source SIMD&FP register and if the signed integer value is less than or equal to zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
func VclezdS64(v0 Int64) Uint64 {
	return C.vclezd_s64(v0)
}

// Floating-point Compare Less than or Equal to zero (vector). This instruction reads each floating-point value in the source SIMD&FP register and if the value is less than or equal to zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
func VclezdF64(v0 Float64) Uint64 {
	return C.vclezd_f64(v0)
}

// Floating-point Compare Less than or Equal to zero (vector). This instruction reads each floating-point value in the source SIMD&FP register and if the value is less than or equal to zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
func VclezsF32(v0 Float32) Uint32 {
	return C.vclezs_f32(v0)
}

// Compare unsigned less than
func VcltqU64(v0 Uint64X2, v1 Uint64X2) Uint64X2 {
	return C.vcltq_u64(v0, v1)
}

// Floating-point compare less than
func VcltqF64(v0 Float64X2, v1 Float64X2) Uint64X2 {
	return C.vcltq_f64(v0, v1)
}

// Compare signed less than
func VcltqS64(v0 Int64X2, v1 Int64X2) Uint64X2 {
	return C.vcltq_s64(v0, v1)
}

// Compare unsigned less than
func VcltU64(v0 Uint64X1, v1 Uint64X1) Uint64X1 {
	return C.vclt_u64(v0, v1)
}

// Floating-point compare less than
func VcltF64(v0 Float64X1, v1 Float64X1) Uint64X1 {
	return C.vclt_f64(v0, v1)
}

// Compare signed less than
func VcltS64(v0 Int64X1, v1 Int64X1) Uint64X1 {
	return C.vclt_s64(v0, v1)
}

// Compare unsigned less than
func VcltdU64(v0 Uint64, v1 Uint64) Uint64 {
	return C.vcltd_u64(v0, v1)
}

// Compare signed less than
func VcltdS64(v0 Int64, v1 Int64) Uint64 {
	return C.vcltd_s64(v0, v1)
}

// Floating-point compare less than
func VcltdF64(v0 Float64, v1 Float64) Uint64 {
	return C.vcltd_f64(v0, v1)
}

// Floating-point compare less than
func VcltsF32(v0 Float32, v1 Float32) Uint32 {
	return C.vclts_f32(v0, v1)
}

// Compare signed Less than zero (vector). This instruction reads each vector element in the source SIMD&FP register and if the signed integer value is less than zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
func VcltzqS8(v0 Int8X16) Uint8X16 {
	return C.vcltzq_s8(v0)
}

// Floating-point Compare Less than zero (vector). This instruction reads each floating-point value in the source SIMD&FP register and if the value is less than zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
func VcltzqF64(v0 Float64X2) Uint64X2 {
	return C.vcltzq_f64(v0)
}

// Floating-point Compare Less than zero (vector). This instruction reads each floating-point value in the source SIMD&FP register and if the value is less than zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
func VcltzqF32(v0 Float32X4) Uint32X4 {
	return C.vcltzq_f32(v0)
}

// Compare signed Less than zero (vector). This instruction reads each vector element in the source SIMD&FP register and if the signed integer value is less than zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
func VcltzqS32(v0 Int32X4) Uint32X4 {
	return C.vcltzq_s32(v0)
}

// Compare signed Less than zero (vector). This instruction reads each vector element in the source SIMD&FP register and if the signed integer value is less than zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
func VcltzqS64(v0 Int64X2) Uint64X2 {
	return C.vcltzq_s64(v0)
}

// Compare signed Less than zero (vector). This instruction reads each vector element in the source SIMD&FP register and if the signed integer value is less than zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
func VcltzqS16(v0 Int16X8) Uint16X8 {
	return C.vcltzq_s16(v0)
}

// Compare signed Less than zero (vector). This instruction reads each vector element in the source SIMD&FP register and if the signed integer value is less than zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
func VcltzS8(v0 Int8X8) Uint8X8 {
	return C.vcltz_s8(v0)
}

// Floating-point Compare Less than zero (vector). This instruction reads each floating-point value in the source SIMD&FP register and if the value is less than zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
func VcltzF64(v0 Float64X1) Uint64X1 {
	return C.vcltz_f64(v0)
}

// Floating-point Compare Less than zero (vector). This instruction reads each floating-point value in the source SIMD&FP register and if the value is less than zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
func VcltzF32(v0 Float32X2) Uint32X2 {
	return C.vcltz_f32(v0)
}

// Compare signed Less than zero (vector). This instruction reads each vector element in the source SIMD&FP register and if the signed integer value is less than zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
func VcltzS32(v0 Int32X2) Uint32X2 {
	return C.vcltz_s32(v0)
}

// Compare signed Less than zero (vector). This instruction reads each vector element in the source SIMD&FP register and if the signed integer value is less than zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
func VcltzS64(v0 Int64X1) Uint64X1 {
	return C.vcltz_s64(v0)
}

// Compare signed Less than zero (vector). This instruction reads each vector element in the source SIMD&FP register and if the signed integer value is less than zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
func VcltzS16(v0 Int16X4) Uint16X4 {
	return C.vcltz_s16(v0)
}

// Compare signed Less than zero (vector). This instruction reads each vector element in the source SIMD&FP register and if the signed integer value is less than zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
func VcltzdS64(v0 Int64) Uint64 {
	return C.vcltzd_s64(v0)
}

// Floating-point Compare Less than zero (vector). This instruction reads each floating-point value in the source SIMD&FP register and if the value is less than zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
func VcltzdF64(v0 Float64) Uint64 {
	return C.vcltzd_f64(v0)
}

// Floating-point Compare Less than zero (vector). This instruction reads each floating-point value in the source SIMD&FP register and if the value is less than zero sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
func VcltzsF32(v0 Float32) Uint32 {
	return C.vcltzs_f32(v0)
}

// Join two smaller vectors into a single larger vector
func VcombineP64(v0 Poly64X1, v1 Poly64X1) Poly64X2 {
	return C.vcombine_p64(v0, v1)
}

// Join two smaller vectors into a single larger vector
func VcombineF64(v0 Float64X1, v1 Float64X1) Float64X2 {
	return C.vcombine_f64(v0, v1)
}

// Signed fixed-point Convert to Floating-point (vector). This instruction converts each element in a vector from fixed-point to floating-point using the rounding mode that is specified by the FPCR, and writes the result to the SIMD&FP destination register.
func VcvtsF32S32(v0 Int32) Float32 {
	return C.vcvts_f32_s32(v0)
}

// Unsigned fixed-point Convert to Floating-point (vector). This instruction converts each element in a vector from fixed-point to floating-point using the rounding mode that is specified by the FPCR, and writes the result to the SIMD&FP destination register.
func VcvtsF32U32(v0 Uint32) Float32 {
	return C.vcvts_f32_u32(v0)
}

// Floating-point Convert to lower precision Narrow (vector). This instruction reads each vector element in the SIMD&FP source register, converts each result to half the precision of the source element, writes the final result to a vector, and writes the vector to the lower or upper half of the destination SIMD&FP register. The destination vector elements are half as long as the source vector elements. The rounding mode is determined by the FPCR.
func VcvtF32F64(v0 Float64X2) Float32X2 {
	return C.vcvt_f32_f64(v0)
}

// Signed fixed-point Convert to Floating-point (vector). This instruction converts each element in a vector from fixed-point to floating-point using the rounding mode that is specified by the FPCR, and writes the result to the SIMD&FP destination register.
func VcvtdF64S64(v0 Int64) Float64 {
	return C.vcvtd_f64_s64(v0)
}

// Unsigned fixed-point Convert to Floating-point (vector). This instruction converts each element in a vector from fixed-point to floating-point using the rounding mode that is specified by the FPCR, and writes the result to the SIMD&FP destination register.
func VcvtdF64U64(v0 Uint64) Float64 {
	return C.vcvtd_f64_u64(v0)
}

// Unsigned fixed-point Convert to Floating-point (vector). This instruction converts each element in a vector from fixed-point to floating-point using the rounding mode that is specified by the FPCR, and writes the result to the SIMD&FP destination register.
func VcvtqF64U64(v0 Uint64X2) Float64X2 {
	return C.vcvtq_f64_u64(v0)
}

// Signed fixed-point Convert to Floating-point (vector). This instruction converts each element in a vector from fixed-point to floating-point using the rounding mode that is specified by the FPCR, and writes the result to the SIMD&FP destination register.
func VcvtqF64S64(v0 Int64X2) Float64X2 {
	return C.vcvtq_f64_s64(v0)
}

// Unsigned fixed-point Convert to Floating-point (vector). This instruction converts each element in a vector from fixed-point to floating-point using the rounding mode that is specified by the FPCR, and writes the result to the SIMD&FP destination register.
func VcvtF64U64(v0 Uint64X1) Float64X1 {
	return C.vcvt_f64_u64(v0)
}

// Signed fixed-point Convert to Floating-point (vector). This instruction converts each element in a vector from fixed-point to floating-point using the rounding mode that is specified by the FPCR, and writes the result to the SIMD&FP destination register.
func VcvtF64S64(v0 Int64X1) Float64X1 {
	return C.vcvt_f64_s64(v0)
}

// Floating-point Convert to higher precision Long (vector). This instruction reads each element in a vector in the SIMD&FP source register, converts each value to double the precision of the source element using the rounding mode that is determined by the FPCR, and writes each result to the equivalent element of the vector in the SIMD&FP destination register.
func VcvtF64F32(v0 Float32X2) Float64X2 {
	return C.vcvt_f64_f32(v0)
}

// Floating-point Convert to lower precision Narrow (vector). This instruction reads each vector element in the SIMD&FP source register, converts each result to half the precision of the source element, writes the final result to a vector, and writes the vector to the lower or upper half of the destination SIMD&FP register. The destination vector elements are half as long as the source vector elements. The rounding mode is determined by the FPCR.
func VcvtHighF32F64(v0 Float32X2, v1 Float64X2) Float32X4 {
	return C.vcvt_high_f32_f64(v0, v1)
}

// Floating-point Convert to higher precision Long (vector). This instruction reads each element in a vector in the SIMD&FP source register, converts each value to double the precision of the source element using the rounding mode that is determined by the FPCR, and writes each result to the equivalent element of the vector in the SIMD&FP destination register.
func VcvtHighF64F32(v0 Float32X4) Float64X2 {
	return C.vcvt_high_f64_f32(v0)
}

// Floating-point Convert to Signed fixed-point, rounding toward Zero (vector). This instruction converts a scalar or each element in a vector from floating-point to fixed-point signed integer using the Round towards Zero rounding mode, and writes the result to the SIMD&FP destination register.
func VcvtsS32F32(v0 Float32) Int32 {
	return C.vcvts_s32_f32(v0)
}

// Floating-point Convert to Signed fixed-point, rounding toward Zero (vector). This instruction converts a scalar or each element in a vector from floating-point to fixed-point signed integer using the Round towards Zero rounding mode, and writes the result to the SIMD&FP destination register.
func VcvtdS64F64(v0 Float64) Int64 {
	return C.vcvtd_s64_f64(v0)
}

// Floating-point Convert to Signed fixed-point, rounding toward Zero (vector). This instruction converts a scalar or each element in a vector from floating-point to fixed-point signed integer using the Round towards Zero rounding mode, and writes the result to the SIMD&FP destination register.
func VcvtqS64F64(v0 Float64X2) Int64X2 {
	return C.vcvtq_s64_f64(v0)
}

// Floating-point Convert to Signed fixed-point, rounding toward Zero (vector). This instruction converts a scalar or each element in a vector from floating-point to fixed-point signed integer using the Round towards Zero rounding mode, and writes the result to the SIMD&FP destination register.
func VcvtS64F64(v0 Float64X1) Int64X1 {
	return C.vcvt_s64_f64(v0)
}

// Floating-point Convert to Unsigned fixed-point, rounding toward Zero (vector). This instruction converts a scalar or each element in a vector from floating-point to fixed-point unsigned integer using the Round towards Zero rounding mode, and writes the result to the general-purpose destination register.
func VcvtsU32F32(v0 Float32) Uint32 {
	return C.vcvts_u32_f32(v0)
}

// Floating-point Convert to Unsigned fixed-point, rounding toward Zero (vector). This instruction converts a scalar or each element in a vector from floating-point to fixed-point unsigned integer using the Round towards Zero rounding mode, and writes the result to the general-purpose destination register.
func VcvtdU64F64(v0 Float64) Uint64 {
	return C.vcvtd_u64_f64(v0)
}

// Floating-point Convert to Unsigned fixed-point, rounding toward Zero (vector). This instruction converts a scalar or each element in a vector from floating-point to fixed-point unsigned integer using the Round towards Zero rounding mode, and writes the result to the general-purpose destination register.
func VcvtqU64F64(v0 Float64X2) Uint64X2 {
	return C.vcvtq_u64_f64(v0)
}

// Floating-point Convert to Unsigned fixed-point, rounding toward Zero (vector). This instruction converts a scalar or each element in a vector from floating-point to fixed-point unsigned integer using the Round towards Zero rounding mode, and writes the result to the general-purpose destination register.
func VcvtU64F64(v0 Float64X1) Uint64X1 {
	return C.vcvt_u64_f64(v0)
}

// Floating-point Convert to Signed integer, rounding to nearest with ties to Away (vector). This instruction converts each element in a vector from a floating-point value to a signed integer value using the Round to Nearest with Ties to Away rounding mode and writes the result to the SIMD&FP destination register.
func VcvtasS32F32(v0 Float32) Int32 {
	return C.vcvtas_s32_f32(v0)
}

// Floating-point Convert to Signed integer, rounding to nearest with ties to Away (vector). This instruction converts each element in a vector from a floating-point value to a signed integer value using the Round to Nearest with Ties to Away rounding mode and writes the result to the SIMD&FP destination register.
func VcvtadS64F64(v0 Float64) Int64 {
	return C.vcvtad_s64_f64(v0)
}

// Floating-point Convert to Unsigned integer, rounding to nearest with ties to Away (vector). This instruction converts each element in a vector from a floating-point value to an unsigned integer value using the Round to Nearest with Ties to Away rounding mode and writes the result to the SIMD&FP destination register.
func VcvtasU32F32(v0 Float32) Uint32 {
	return C.vcvtas_u32_f32(v0)
}

// Floating-point Convert to Unsigned integer, rounding to nearest with ties to Away (vector). This instruction converts each element in a vector from a floating-point value to an unsigned integer value using the Round to Nearest with Ties to Away rounding mode and writes the result to the SIMD&FP destination register.
func VcvtadU64F64(v0 Float64) Uint64 {
	return C.vcvtad_u64_f64(v0)
}

// Floating-point Convert to Signed integer, rounding toward Minus infinity (vector). This instruction converts a scalar or each element in a vector from a floating-point value to a signed integer value using the Round towards Minus Infinity rounding mode, and writes the result to the SIMD&FP destination register.
func VcvtmsS32F32(v0 Float32) Int32 {
	return C.vcvtms_s32_f32(v0)
}

// Floating-point Convert to Signed integer, rounding toward Minus infinity (vector). This instruction converts a scalar or each element in a vector from a floating-point value to a signed integer value using the Round towards Minus Infinity rounding mode, and writes the result to the SIMD&FP destination register.
func VcvtmdS64F64(v0 Float64) Int64 {
	return C.vcvtmd_s64_f64(v0)
}

// Floating-point Convert to Unsigned integer, rounding toward Minus infinity (vector). This instruction converts a scalar or each element in a vector from a floating-point value to an unsigned integer value using the Round towards Minus Infinity rounding mode, and writes the result to the SIMD&FP destination register.
func VcvtmsU32F32(v0 Float32) Uint32 {
	return C.vcvtms_u32_f32(v0)
}

// Floating-point Convert to Unsigned integer, rounding toward Minus infinity (vector). This instruction converts a scalar or each element in a vector from a floating-point value to an unsigned integer value using the Round towards Minus Infinity rounding mode, and writes the result to the SIMD&FP destination register.
func VcvtmdU64F64(v0 Float64) Uint64 {
	return C.vcvtmd_u64_f64(v0)
}

// Floating-point Convert to Signed integer, rounding to nearest with ties to even (vector). This instruction converts a scalar or each element in a vector from a floating-point value to a signed integer value using the Round to Nearest rounding mode, and writes the result to the SIMD&FP destination register.
func VcvtnsS32F32(v0 Float32) Int32 {
	return C.vcvtns_s32_f32(v0)
}

// Floating-point Convert to Signed integer, rounding to nearest with ties to even (vector). This instruction converts a scalar or each element in a vector from a floating-point value to a signed integer value using the Round to Nearest rounding mode, and writes the result to the SIMD&FP destination register.
func VcvtndS64F64(v0 Float64) Int64 {
	return C.vcvtnd_s64_f64(v0)
}

// Floating-point Convert to Unsigned integer, rounding to nearest with ties to even (vector). This instruction converts a scalar or each element in a vector from a floating-point value to an unsigned integer value using the Round to Nearest rounding mode, and writes the result to the SIMD&FP destination register.
func VcvtnsU32F32(v0 Float32) Uint32 {
	return C.vcvtns_u32_f32(v0)
}

// Floating-point Convert to Unsigned integer, rounding to nearest with ties to even (vector). This instruction converts a scalar or each element in a vector from a floating-point value to an unsigned integer value using the Round to Nearest rounding mode, and writes the result to the SIMD&FP destination register.
func VcvtndU64F64(v0 Float64) Uint64 {
	return C.vcvtnd_u64_f64(v0)
}

// Floating-point Convert to Signed integer, rounding toward Plus infinity (vector). This instruction converts a scalar or each element in a vector from a floating-point value to a signed integer value using the Round towards Plus Infinity rounding mode, and writes the result to the SIMD&FP destination register.
func VcvtpsS32F32(v0 Float32) Int32 {
	return C.vcvtps_s32_f32(v0)
}

// Floating-point Convert to Signed integer, rounding toward Plus infinity (vector). This instruction converts a scalar or each element in a vector from a floating-point value to a signed integer value using the Round towards Plus Infinity rounding mode, and writes the result to the SIMD&FP destination register.
func VcvtpdS64F64(v0 Float64) Int64 {
	return C.vcvtpd_s64_f64(v0)
}

// Floating-point Convert to Unsigned integer, rounding toward Plus infinity (vector). This instruction converts a scalar or each element in a vector from a floating-point value to an unsigned integer value using the Round towards Plus Infinity rounding mode, and writes the result to the SIMD&FP destination register.
func VcvtpsU32F32(v0 Float32) Uint32 {
	return C.vcvtps_u32_f32(v0)
}

// Floating-point Convert to Unsigned integer, rounding toward Plus infinity (vector). This instruction converts a scalar or each element in a vector from a floating-point value to an unsigned integer value using the Round towards Plus Infinity rounding mode, and writes the result to the SIMD&FP destination register.
func VcvtpdU64F64(v0 Float64) Uint64 {
	return C.vcvtpd_u64_f64(v0)
}

// Floating-point Convert to lower precision Narrow, rounding to odd (vector). This instruction reads each vector element in the source SIMD&FP register, narrows each value to half the precision of the source element using the Round to Odd rounding mode, writes the result to a vector, and writes the vector to the destination SIMD&FP register.
func VcvtxdF32F64(v0 Float64) Float32 {
	return C.vcvtxd_f32_f64(v0)
}

// Floating-point Convert to lower precision Narrow, rounding to odd (vector). This instruction reads each vector element in the source SIMD&FP register, narrows each value to half the precision of the source element using the Round to Odd rounding mode, writes the result to a vector, and writes the vector to the destination SIMD&FP register.
func VcvtxF32F64(v0 Float64X2) Float32X2 {
	return C.vcvtx_f32_f64(v0)
}

// Floating-point Convert to lower precision Narrow, rounding to odd (vector). This instruction reads each vector element in the source SIMD&FP register, narrows each value to half the precision of the source element using the Round to Odd rounding mode, writes the result to a vector, and writes the vector to the destination SIMD&FP register.
func VcvtxHighF32F64(v0 Float32X2, v1 Float64X2) Float32X4 {
	return C.vcvtx_high_f32_f64(v0, v1)
}

// Floating-point Divide (vector). This instruction divides the floating-point values in the elements in the first source SIMD&FP register, by the floating-point values in the corresponding elements in the second source SIMD&FP register, places the results in a vector, and writes the vector to the destination SIMD&FP register.
func VdivqF64(v0 Float64X2, v1 Float64X2) Float64X2 {
	return C.vdivq_f64(v0, v1)
}

// Floating-point Divide (vector). This instruction divides the floating-point values in the elements in the first source SIMD&FP register, by the floating-point values in the corresponding elements in the second source SIMD&FP register, places the results in a vector, and writes the vector to the destination SIMD&FP register.
func VdivqF32(v0 Float32X4, v1 Float32X4) Float32X4 {
	return C.vdivq_f32(v0, v1)
}

// Floating-point Divide (vector). This instruction divides the floating-point values in the elements in the first source SIMD&FP register, by the floating-point values in the corresponding elements in the second source SIMD&FP register, places the results in a vector, and writes the vector to the destination SIMD&FP register.
func VdivF64(v0 Float64X1, v1 Float64X1) Float64X1 {
	return C.vdiv_f64(v0, v1)
}

// Floating-point Divide (vector). This instruction divides the floating-point values in the elements in the first source SIMD&FP register, by the floating-point values in the corresponding elements in the second source SIMD&FP register, places the results in a vector, and writes the vector to the destination SIMD&FP register.
func VdivF32(v0 Float32X2, v1 Float32X2) Float32X2 {
	return C.vdiv_f32(v0, v1)
}

// Insert vector element from another vector element. This instruction copies the vector element of the source SIMD&FP register to the specified vector element of the destination SIMD&FP register.
func VdupNP64(v0 Poly64) Poly64X1 {
	return C.vdup_n_p64(v0)
}

// Duplicate vector element to vector or scalar. This instruction duplicates the vector element at the specified element index in the source SIMD&FP register into a scalar or each element in a vector, and writes the result to the destination SIMD&FP register.
func VdupqNP64(v0 Poly64) Poly64X2 {
	return C.vdupq_n_p64(v0)
}

// Duplicate vector element to vector or scalar. This instruction duplicates the vector element at the specified element index in the source SIMD&FP register into a scalar or each element in a vector, and writes the result to the destination SIMD&FP register.
func VdupqNF64(v0 Float64) Float64X2 {
	return C.vdupq_n_f64(v0)
}

// Insert vector element from another vector element. This instruction copies the vector element of the source SIMD&FP register to the specified vector element of the destination SIMD&FP register.
func VdupNF64(v0 Float64) Float64X1 {
	return C.vdup_n_f64(v0)
}

// Floating-point fused Multiply-Add to accumulator (vector). This instruction multiplies corresponding floating-point values in the vectors in the two source SIMD&FP registers, adds the product to the corresponding vector element of the destination SIMD&FP register, and writes the result to the destination SIMD&FP register.
func VfmaqF64(v0 Float64X2, v1 Float64X2, v2 Float64X2) Float64X2 {
	return C.vfmaq_f64(v0, v1, v2)
}

// Floating-point fused Multiply-Add (scalar). This instruction multiplies the values of the first two SIMD&FP source registers, adds the product to the value of the third SIMD&FP source register, and writes the result to the SIMD&FP destination register.
func VfmaF64(v0 Float64X1, v1 Float64X1, v2 Float64X1) Float64X1 {
	return C.vfma_f64(v0, v1, v2)
}

// Floating-point fused Multiply-Add to accumulator (vector). This instruction multiplies corresponding floating-point values in the vectors in the two source SIMD&FP registers, adds the product to the corresponding vector element of the destination SIMD&FP register, and writes the result to the destination SIMD&FP register.
func VfmaqNF64(v0 Float64X2, v1 Float64X2, v2 Float64) Float64X2 {
	return C.vfmaq_n_f64(v0, v1, v2)
}

// Floating-point fused Multiply-Add (scalar). This instruction multiplies the values of the first two SIMD&FP source registers, adds the product to the value of the third SIMD&FP source register, and writes the result to the SIMD&FP destination register.
func VfmaNF64(v0 Float64X1, v1 Float64X1, v2 Float64) Float64X1 {
	return C.vfma_n_f64(v0, v1, v2)
}

// Floating-point fused Multiply-Subtract from accumulator (vector). This instruction multiplies corresponding floating-point values in the vectors in the two source SIMD&FP registers, negates the product, adds the result to the corresponding vector element of the destination SIMD&FP register, and writes the result to the destination SIMD&FP register.
func VfmsqF64(v0 Float64X2, v1 Float64X2, v2 Float64X2) Float64X2 {
	return C.vfmsq_f64(v0, v1, v2)
}

// Floating-point Fused Multiply-Subtract (scalar). This instruction multiplies the values of the first two SIMD&FP source registers, negates the product, adds that to the value of the third SIMD&FP source register, and writes the result to the SIMD&FP destination register.
func VfmsF64(v0 Float64X1, v1 Float64X1, v2 Float64X1) Float64X1 {
	return C.vfms_f64(v0, v1, v2)
}

// Floating-point fused Multiply-Subtract from accumulator (vector). This instruction multiplies corresponding floating-point values in the vectors in the two source SIMD&FP registers, negates the product, adds the result to the corresponding vector element of the destination SIMD&FP register, and writes the result to the destination SIMD&FP register.
func VfmsqNF64(v0 Float64X2, v1 Float64X2, v2 Float64) Float64X2 {
	return C.vfmsq_n_f64(v0, v1, v2)
}

// Floating-point fused Multiply-Subtract from accumulator (vector). This instruction multiplies corresponding floating-point values in the vectors in the two source SIMD&FP registers, negates the product, adds the result to the corresponding vector element of the destination SIMD&FP register, and writes the result to the destination SIMD&FP register.
func VfmsqNF32(v0 Float32X4, v1 Float32X4, v2 Float32) Float32X4 {
	return C.vfmsq_n_f32(v0, v1, v2)
}

// Floating-point Fused Multiply-Subtract (scalar). This instruction multiplies the values of the first two SIMD&FP source registers, negates the product, adds that to the value of the third SIMD&FP source register, and writes the result to the SIMD&FP destination register.
func VfmsNF64(v0 Float64X1, v1 Float64X1, v2 Float64) Float64X1 {
	return C.vfms_n_f64(v0, v1, v2)
}

// Floating-point fused Multiply-Subtract from accumulator (vector). This instruction multiplies corresponding floating-point values in the vectors in the two source SIMD&FP registers, negates the product, adds the result to the corresponding vector element of the destination SIMD&FP register, and writes the result to the destination SIMD&FP register.
func VfmsNF32(v0 Float32X2, v1 Float32X2, v2 Float32) Float32X2 {
	return C.vfms_n_f32(v0, v1, v2)
}

// Duplicate vector element to vector or scalar. This instruction duplicates the vector element at the specified element index in the source SIMD&FP register into a scalar or each element in a vector, and writes the result to the destination SIMD&FP register.
func VgetHighP64(v0 Poly64X2) Poly64X1 {
	return C.vget_high_p64(v0)
}

// Duplicate vector element to vector or scalar. This instruction duplicates the vector element at the specified element index in the source SIMD&FP register into a scalar or each element in a vector, and writes the result to the destination SIMD&FP register.
func VgetHighF64(v0 Float64X2) Float64X1 {
	return C.vget_high_f64(v0)
}

// Duplicate vector element to vector or scalar. This instruction duplicates the vector element at the specified element index in the source SIMD&FP register into a scalar or each element in a vector, and writes the result to the destination SIMD&FP register.
func VgetLowP64(v0 Poly64X2) Poly64X1 {
	return C.vget_low_p64(v0)
}

// Duplicate vector element to vector or scalar. This instruction duplicates the vector element at the specified element index in the source SIMD&FP register into a scalar or each element in a vector, and writes the result to the destination SIMD&FP register.
func VgetLowF64(v0 Float64X2) Float64X1 {
	return C.vget_low_f64(v0)
}

// Floating-point Maximum (vector). This instruction compares corresponding vector elements in the two source SIMD&FP registers, places the larger of each of the two floating-point values into a vector, and writes the vector to the destination SIMD&FP register.
func VmaxqF64(v0 Float64X2, v1 Float64X2) Float64X2 {
	return C.vmaxq_f64(v0, v1)
}

// Floating-point Maximum (vector). This instruction compares corresponding vector elements in the two source SIMD&FP registers, places the larger of each of the two floating-point values into a vector, and writes the vector to the destination SIMD&FP register.
func VmaxF64(v0 Float64X1, v1 Float64X1) Float64X1 {
	return C.vmax_f64(v0, v1)
}

// Floating-point Maximum Number Pairwise (vector). This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements in the two source SIMD&FP registers, writes the largest of each pair of values into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are floating-point values.
func VmaxnmvqF64(v0 Float64X2) Float64 {
	return C.vmaxnmvq_f64(v0)
}

// Floating-point Maximum Number across Vector. This instruction compares all the vector elements in the source SIMD&FP register, and writes the largest of the values as a scalar to the destination SIMD&FP register. All the values in this instruction are floating-point values.
func VmaxnmvqF32(v0 Float32X4) Float32 {
	return C.vmaxnmvq_f32(v0)
}

// Floating-point Maximum Number Pairwise (vector). This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements in the two source SIMD&FP registers, writes the largest of each pair of values into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are floating-point values.
func VmaxnmvF32(v0 Float32X2) Float32 {
	return C.vmaxnmv_f32(v0)
}

// Unsigned Maximum across Vector. This instruction compares all the vector elements in the source SIMD&FP register, and writes the largest of the values as a scalar to the destination SIMD&FP register. All the values in this instruction are unsigned integer values.
func VmaxvqU8(v0 Uint8X16) Uint8 {
	return C.vmaxvq_u8(v0)
}

// Unsigned Maximum across Vector. This instruction compares all the vector elements in the source SIMD&FP register, and writes the largest of the values as a scalar to the destination SIMD&FP register. All the values in this instruction are unsigned integer values.
func VmaxvqU32(v0 Uint32X4) Uint32 {
	return C.vmaxvq_u32(v0)
}

// Unsigned Maximum across Vector. This instruction compares all the vector elements in the source SIMD&FP register, and writes the largest of the values as a scalar to the destination SIMD&FP register. All the values in this instruction are unsigned integer values.
func VmaxvqU16(v0 Uint16X8) Uint16 {
	return C.vmaxvq_u16(v0)
}

// Signed Maximum across Vector. This instruction compares all the vector elements in the source SIMD&FP register, and writes the largest of the values as a scalar to the destination SIMD&FP register. All the values in this instruction are signed integer values.
func VmaxvqS8(v0 Int8X16) Int8 {
	return C.vmaxvq_s8(v0)
}

// Floating-point Maximum Pairwise (vector). This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements from the concatenated vector, writes the larger of each pair of values into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are floating-point values.
func VmaxvqF64(v0 Float64X2) Float64 {
	return C.vmaxvq_f64(v0)
}

// Floating-point Maximum across Vector. This instruction compares all the vector elements in the source SIMD&FP register, and writes the largest of the values as a scalar to the destination SIMD&FP register. All the values in this instruction are floating-point values.
func VmaxvqF32(v0 Float32X4) Float32 {
	return C.vmaxvq_f32(v0)
}

// Signed Maximum across Vector. This instruction compares all the vector elements in the source SIMD&FP register, and writes the largest of the values as a scalar to the destination SIMD&FP register. All the values in this instruction are signed integer values.
func VmaxvqS32(v0 Int32X4) Int32 {
	return C.vmaxvq_s32(v0)
}

// Signed Maximum across Vector. This instruction compares all the vector elements in the source SIMD&FP register, and writes the largest of the values as a scalar to the destination SIMD&FP register. All the values in this instruction are signed integer values.
func VmaxvqS16(v0 Int16X8) Int16 {
	return C.vmaxvq_s16(v0)
}

// Unsigned Maximum across Vector. This instruction compares all the vector elements in the source SIMD&FP register, and writes the largest of the values as a scalar to the destination SIMD&FP register. All the values in this instruction are unsigned integer values.
func VmaxvU8(v0 Uint8X8) Uint8 {
	return C.vmaxv_u8(v0)
}

// Unsigned Maximum Pairwise. This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements in the two source SIMD&FP registers, writes the largest of each pair of unsigned integer values into a vector, and writes the vector to the destination SIMD&FP register.
func VmaxvU32(v0 Uint32X2) Uint32 {
	return C.vmaxv_u32(v0)
}

// Unsigned Maximum across Vector. This instruction compares all the vector elements in the source SIMD&FP register, and writes the largest of the values as a scalar to the destination SIMD&FP register. All the values in this instruction are unsigned integer values.
func VmaxvU16(v0 Uint16X4) Uint16 {
	return C.vmaxv_u16(v0)
}

// Signed Maximum across Vector. This instruction compares all the vector elements in the source SIMD&FP register, and writes the largest of the values as a scalar to the destination SIMD&FP register. All the values in this instruction are signed integer values.
func VmaxvS8(v0 Int8X8) Int8 {
	return C.vmaxv_s8(v0)
}

// Floating-point Maximum Pairwise (vector). This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements from the concatenated vector, writes the larger of each pair of values into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are floating-point values.
func VmaxvF32(v0 Float32X2) Float32 {
	return C.vmaxv_f32(v0)
}

// Signed Maximum Pairwise. This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements in the two source SIMD&FP registers, writes the largest of each pair of signed integer values into a vector, and writes the vector to the destination SIMD&FP register.
func VmaxvS32(v0 Int32X2) Int32 {
	return C.vmaxv_s32(v0)
}

// Signed Maximum across Vector. This instruction compares all the vector elements in the source SIMD&FP register, and writes the largest of the values as a scalar to the destination SIMD&FP register. All the values in this instruction are signed integer values.
func VmaxvS16(v0 Int16X4) Int16 {
	return C.vmaxv_s16(v0)
}

// Floating-point minimum (vector). This instruction compares corresponding elements in the vectors in the two source SIMD&FP registers, places the smaller of each of the two floating-point values into a vector, and writes the vector to the destination SIMD&FP register.
func VminqF64(v0 Float64X2, v1 Float64X2) Float64X2 {
	return C.vminq_f64(v0, v1)
}

// Floating-point minimum (vector). This instruction compares corresponding elements in the vectors in the two source SIMD&FP registers, places the smaller of each of the two floating-point values into a vector, and writes the vector to the destination SIMD&FP register.
func VminF64(v0 Float64X1, v1 Float64X1) Float64X1 {
	return C.vmin_f64(v0, v1)
}

// Floating-point Minimum Number Pairwise (vector). This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements in the two source SIMD&FP registers, writes the smallest of each pair of floating-point values into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are floating-point values.
func VminnmvqF64(v0 Float64X2) Float64 {
	return C.vminnmvq_f64(v0)
}

// Floating-point Minimum Number across Vector. This instruction compares all the vector elements in the source SIMD&FP register, and writes the smallest of the values as a scalar to the destination SIMD&FP register. All the values in this instruction are floating-point values.
func VminnmvqF32(v0 Float32X4) Float32 {
	return C.vminnmvq_f32(v0)
}

// Floating-point Minimum Number Pairwise (vector). This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements in the two source SIMD&FP registers, writes the smallest of each pair of floating-point values into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are floating-point values.
func VminnmvF32(v0 Float32X2) Float32 {
	return C.vminnmv_f32(v0)
}

// Unsigned Minimum across Vector. This instruction compares all the vector elements in the source SIMD&FP register, and writes the smallest of the values as a scalar to the destination SIMD&FP register. All the values in this instruction are unsigned integer values.
func VminvqU8(v0 Uint8X16) Uint8 {
	return C.vminvq_u8(v0)
}

// Unsigned Minimum across Vector. This instruction compares all the vector elements in the source SIMD&FP register, and writes the smallest of the values as a scalar to the destination SIMD&FP register. All the values in this instruction are unsigned integer values.
func VminvqU32(v0 Uint32X4) Uint32 {
	return C.vminvq_u32(v0)
}

// Unsigned Minimum across Vector. This instruction compares all the vector elements in the source SIMD&FP register, and writes the smallest of the values as a scalar to the destination SIMD&FP register. All the values in this instruction are unsigned integer values.
func VminvqU16(v0 Uint16X8) Uint16 {
	return C.vminvq_u16(v0)
}

// Signed Minimum across Vector. This instruction compares all the vector elements in the source SIMD&FP register, and writes the smallest of the values as a scalar to the destination SIMD&FP register. All the values in this instruction are signed integer values.
func VminvqS8(v0 Int8X16) Int8 {
	return C.vminvq_s8(v0)
}

// Floating-point Minimum Pairwise (vector). This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements from the concatenated vector, writes the smaller of each pair of values into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are floating-point values.
func VminvqF64(v0 Float64X2) Float64 {
	return C.vminvq_f64(v0)
}

// Floating-point Minimum across Vector. This instruction compares all the vector elements in the source SIMD&FP register, and writes the smallest of the values as a scalar to the destination SIMD&FP register. All the values in this instruction are floating-point values.
func VminvqF32(v0 Float32X4) Float32 {
	return C.vminvq_f32(v0)
}

// Signed Minimum across Vector. This instruction compares all the vector elements in the source SIMD&FP register, and writes the smallest of the values as a scalar to the destination SIMD&FP register. All the values in this instruction are signed integer values.
func VminvqS32(v0 Int32X4) Int32 {
	return C.vminvq_s32(v0)
}

// Signed Minimum across Vector. This instruction compares all the vector elements in the source SIMD&FP register, and writes the smallest of the values as a scalar to the destination SIMD&FP register. All the values in this instruction are signed integer values.
func VminvqS16(v0 Int16X8) Int16 {
	return C.vminvq_s16(v0)
}

// Unsigned Minimum across Vector. This instruction compares all the vector elements in the source SIMD&FP register, and writes the smallest of the values as a scalar to the destination SIMD&FP register. All the values in this instruction are unsigned integer values.
func VminvU8(v0 Uint8X8) Uint8 {
	return C.vminv_u8(v0)
}

// Unsigned Minimum Pairwise. This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements in the two source SIMD&FP registers, writes the smallest of each pair of unsigned integer values into a vector, and writes the vector to the destination SIMD&FP register.
func VminvU32(v0 Uint32X2) Uint32 {
	return C.vminv_u32(v0)
}

// Unsigned Minimum across Vector. This instruction compares all the vector elements in the source SIMD&FP register, and writes the smallest of the values as a scalar to the destination SIMD&FP register. All the values in this instruction are unsigned integer values.
func VminvU16(v0 Uint16X4) Uint16 {
	return C.vminv_u16(v0)
}

// Signed Minimum across Vector. This instruction compares all the vector elements in the source SIMD&FP register, and writes the smallest of the values as a scalar to the destination SIMD&FP register. All the values in this instruction are signed integer values.
func VminvS8(v0 Int8X8) Int8 {
	return C.vminv_s8(v0)
}

// Floating-point Minimum Pairwise (vector). This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements from the concatenated vector, writes the smaller of each pair of values into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are floating-point values.
func VminvF32(v0 Float32X2) Float32 {
	return C.vminv_f32(v0)
}

// Signed Minimum Pairwise. This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements in the two source SIMD&FP registers, writes the smallest of each pair of signed integer values into a vector, and writes the vector to the destination SIMD&FP register.
func VminvS32(v0 Int32X2) Int32 {
	return C.vminv_s32(v0)
}

// Signed Minimum across Vector. This instruction compares all the vector elements in the source SIMD&FP register, and writes the smallest of the values as a scalar to the destination SIMD&FP register. All the values in this instruction are signed integer values.
func VminvS16(v0 Int16X4) Int16 {
	return C.vminv_s16(v0)
}

// Floating-point multiply-add to accumulator
func VmlaqF64(v0 Float64X2, v1 Float64X2, v2 Float64X2) Float64X2 {
	return C.vmlaq_f64(v0, v1, v2)
}

// Floating-point multiply-add to accumulator
func VmlaF64(v0 Float64X1, v1 Float64X1, v2 Float64X1) Float64X1 {
	return C.vmla_f64(v0, v1, v2)
}

// Multiply-subtract from accumulator
func VmlsqF64(v0 Float64X2, v1 Float64X2, v2 Float64X2) Float64X2 {
	return C.vmlsq_f64(v0, v1, v2)
}

// Multiply-subtract from accumulator
func VmlsF64(v0 Float64X1, v1 Float64X1, v2 Float64X1) Float64X1 {
	return C.vmls_f64(v0, v1, v2)
}

// vmov_n_p64
func VmovNP64(v0 Poly64) Poly64X1 {
	return C.vmov_n_p64(v0)
}

// vmovq_n_p64
func VmovqNP64(v0 Poly64) Poly64X2 {
	return C.vmovq_n_p64(v0)
}

// Duplicate vector element to vector or scalar. This instruction duplicates the vector element at the specified element index in the source SIMD&FP register into a scalar or each element in a vector, and writes the result to the destination SIMD&FP register.
func VmovqNF64(v0 Float64) Float64X2 {
	return C.vmovq_n_f64(v0)
}

// Duplicate vector element to vector or scalar. This instruction duplicates the vector element at the specified element index in the source SIMD&FP register into a scalar or each element in a vector, and writes the result to the destination SIMD&FP register.
func VmovNF64(v0 Float64) Float64X1 {
	return C.vmov_n_f64(v0)
}

// Vector move
func VmovlHighU8(v0 Uint8X16) Uint16X8 {
	return C.vmovl_high_u8(v0)
}

// Vector move
func VmovlHighU32(v0 Uint32X4) Uint64X2 {
	return C.vmovl_high_u32(v0)
}

// Vector move
func VmovlHighU16(v0 Uint16X8) Uint32X4 {
	return C.vmovl_high_u16(v0)
}

// Vector move
func VmovlHighS8(v0 Int8X16) Int16X8 {
	return C.vmovl_high_s8(v0)
}

// Vector move
func VmovlHighS32(v0 Int32X4) Int64X2 {
	return C.vmovl_high_s32(v0)
}

// Vector move
func VmovlHighS16(v0 Int16X8) Int32X4 {
	return C.vmovl_high_s16(v0)
}

// Extract Narrow. This instruction reads each vector element from the source SIMD&FP register, narrows each value to half the original width, places the result into a vector, and writes the vector to the lower or upper half of the destination SIMD&FP register. The destination vector elements are half as long as the source vector elements.
func VmovnHighU32(v0 Uint16X4, v1 Uint32X4) Uint16X8 {
	return C.vmovn_high_u32(v0, v1)
}

// Extract Narrow. This instruction reads each vector element from the source SIMD&FP register, narrows each value to half the original width, places the result into a vector, and writes the vector to the lower or upper half of the destination SIMD&FP register. The destination vector elements are half as long as the source vector elements.
func VmovnHighU64(v0 Uint32X2, v1 Uint64X2) Uint32X4 {
	return C.vmovn_high_u64(v0, v1)
}

// Extract Narrow. This instruction reads each vector element from the source SIMD&FP register, narrows each value to half the original width, places the result into a vector, and writes the vector to the lower or upper half of the destination SIMD&FP register. The destination vector elements are half as long as the source vector elements.
func VmovnHighU16(v0 Uint8X8, v1 Uint16X8) Uint8X16 {
	return C.vmovn_high_u16(v0, v1)
}

// Extract Narrow. This instruction reads each vector element from the source SIMD&FP register, narrows each value to half the original width, places the result into a vector, and writes the vector to the lower or upper half of the destination SIMD&FP register. The destination vector elements are half as long as the source vector elements.
func VmovnHighS32(v0 Int16X4, v1 Int32X4) Int16X8 {
	return C.vmovn_high_s32(v0, v1)
}

// Extract Narrow. This instruction reads each vector element from the source SIMD&FP register, narrows each value to half the original width, places the result into a vector, and writes the vector to the lower or upper half of the destination SIMD&FP register. The destination vector elements are half as long as the source vector elements.
func VmovnHighS64(v0 Int32X2, v1 Int64X2) Int32X4 {
	return C.vmovn_high_s64(v0, v1)
}

// Extract Narrow. This instruction reads each vector element from the source SIMD&FP register, narrows each value to half the original width, places the result into a vector, and writes the vector to the lower or upper half of the destination SIMD&FP register. The destination vector elements are half as long as the source vector elements.
func VmovnHighS16(v0 Int8X8, v1 Int16X8) Int8X16 {
	return C.vmovn_high_s16(v0, v1)
}

// Floating-point Multiply (vector). This instruction multiplies corresponding floating-point values in the vectors in the two source SIMD&FP registers, places the result in a vector, and writes the vector to the destination SIMD&FP register.
func VmulqF64(v0 Float64X2, v1 Float64X2) Float64X2 {
	return C.vmulq_f64(v0, v1)
}

// Floating-point Multiply (vector). This instruction multiplies corresponding floating-point values in the vectors in the two source SIMD&FP registers, places the result in a vector, and writes the vector to the destination SIMD&FP register.
func VmulF64(v0 Float64X1, v1 Float64X1) Float64X1 {
	return C.vmul_f64(v0, v1)
}

// Floating-point Multiply (vector). This instruction multiplies corresponding floating-point values in the vectors in the two source SIMD&FP registers, places the result in a vector, and writes the vector to the destination SIMD&FP register.
func VmulNF64(v0 Float64X1, v1 Float64) Float64X1 {
	return C.vmul_n_f64(v0, v1)
}

// Floating-point Multiply (vector). This instruction multiplies corresponding floating-point values in the vectors in the two source SIMD&FP registers, places the result in a vector, and writes the vector to the destination SIMD&FP register.
func VmulqNF64(v0 Float64X2, v1 Float64) Float64X2 {
	return C.vmulq_n_f64(v0, v1)
}

// Polynomial Multiply Long. This instruction multiplies corresponding elements in the lower or upper half of the vectors of the two source SIMD&FP registers, places the results in a vector, and writes the vector to the destination SIMD&FP register. The destination vector elements are twice as long as the elements that are multiplied.
func VmullP64(v0 Poly64, v1 Poly64) Poly128 {
	return C.vmull_p64(v0, v1)
}

// Polynomial Multiply Long. This instruction multiplies corresponding elements in the lower or upper half of the vectors of the two source SIMD&FP registers, places the results in a vector, and writes the vector to the destination SIMD&FP register. The destination vector elements are twice as long as the elements that are multiplied.
func VmullHighP8(v0 Poly8X16, v1 Poly8X16) Poly16X8 {
	return C.vmull_high_p8(v0, v1)
}

// Unsigned Multiply long (vector). This instruction multiplies corresponding vector elements in the lower or upper half of the two source SIMD&FP registers, places the result in a vector, and writes the vector to the destination SIMD&FP register. The destination vector elements are twice as long as the elements that are multiplied. All the values in this instruction are unsigned integer values.
func VmullHighU8(v0 Uint8X16, v1 Uint8X16) Uint16X8 {
	return C.vmull_high_u8(v0, v1)
}

// Unsigned Multiply long (vector). This instruction multiplies corresponding vector elements in the lower or upper half of the two source SIMD&FP registers, places the result in a vector, and writes the vector to the destination SIMD&FP register. The destination vector elements are twice as long as the elements that are multiplied. All the values in this instruction are unsigned integer values.
func VmullHighU32(v0 Uint32X4, v1 Uint32X4) Uint64X2 {
	return C.vmull_high_u32(v0, v1)
}

// Unsigned Multiply long (vector). This instruction multiplies corresponding vector elements in the lower or upper half of the two source SIMD&FP registers, places the result in a vector, and writes the vector to the destination SIMD&FP register. The destination vector elements are twice as long as the elements that are multiplied. All the values in this instruction are unsigned integer values.
func VmullHighU16(v0 Uint16X8, v1 Uint16X8) Uint32X4 {
	return C.vmull_high_u16(v0, v1)
}

// Signed Multiply Long (vector). This instruction multiplies corresponding signed integer values in the lower or upper half of the vectors of the two source SIMD&FP registers, places the results in a vector, and writes the vector to the destination SIMD&FP register.
func VmullHighS8(v0 Int8X16, v1 Int8X16) Int16X8 {
	return C.vmull_high_s8(v0, v1)
}

// Signed Multiply Long (vector). This instruction multiplies corresponding signed integer values in the lower or upper half of the vectors of the two source SIMD&FP registers, places the results in a vector, and writes the vector to the destination SIMD&FP register.
func VmullHighS32(v0 Int32X4, v1 Int32X4) Int64X2 {
	return C.vmull_high_s32(v0, v1)
}

// Signed Multiply Long (vector). This instruction multiplies corresponding signed integer values in the lower or upper half of the vectors of the two source SIMD&FP registers, places the results in a vector, and writes the vector to the destination SIMD&FP register.
func VmullHighS16(v0 Int16X8, v1 Int16X8) Int32X4 {
	return C.vmull_high_s16(v0, v1)
}

// Polynomial Multiply Long. This instruction multiplies corresponding elements in the lower or upper half of the vectors of the two source SIMD&FP registers, places the results in a vector, and writes the vector to the destination SIMD&FP register. The destination vector elements are twice as long as the elements that are multiplied.
func VmullHighP64(v0 Poly64X2, v1 Poly64X2) Poly128 {
	return C.vmull_high_p64(v0, v1)
}

// Unsigned Multiply long (vector). This instruction multiplies corresponding vector elements in the lower or upper half of the two source SIMD&FP registers, places the result in a vector, and writes the vector to the destination SIMD&FP register. The destination vector elements are twice as long as the elements that are multiplied. All the values in this instruction are unsigned integer values.
func VmullHighNU32(v0 Uint32X4, v1 Uint32) Uint64X2 {
	return C.vmull_high_n_u32(v0, v1)
}

// Unsigned Multiply long (vector). This instruction multiplies corresponding vector elements in the lower or upper half of the two source SIMD&FP registers, places the result in a vector, and writes the vector to the destination SIMD&FP register. The destination vector elements are twice as long as the elements that are multiplied. All the values in this instruction are unsigned integer values.
func VmullHighNU16(v0 Uint16X8, v1 Uint16) Uint32X4 {
	return C.vmull_high_n_u16(v0, v1)
}

// Signed Multiply Long (vector). This instruction multiplies corresponding signed integer values in the lower or upper half of the vectors of the two source SIMD&FP registers, places the results in a vector, and writes the vector to the destination SIMD&FP register.
func VmullHighNS32(v0 Int32X4, v1 Int32) Int64X2 {
	return C.vmull_high_n_s32(v0, v1)
}

// Signed Multiply Long (vector). This instruction multiplies corresponding signed integer values in the lower or upper half of the vectors of the two source SIMD&FP registers, places the results in a vector, and writes the vector to the destination SIMD&FP register.
func VmullHighNS16(v0 Int16X8, v1 Int16) Int32X4 {
	return C.vmull_high_n_s16(v0, v1)
}

// Floating-point Multiply extended. This instruction multiplies corresponding floating-point values in the vectors of the two source SIMD&FP registers, places the resulting floating-point values in a vector, and writes the vector to the destination SIMD&FP register.
func VmulxqF64(v0 Float64X2, v1 Float64X2) Float64X2 {
	return C.vmulxq_f64(v0, v1)
}

// Floating-point Multiply extended. This instruction multiplies corresponding floating-point values in the vectors of the two source SIMD&FP registers, places the resulting floating-point values in a vector, and writes the vector to the destination SIMD&FP register.
func VmulxqF32(v0 Float32X4, v1 Float32X4) Float32X4 {
	return C.vmulxq_f32(v0, v1)
}

// Floating-point Multiply extended. This instruction multiplies corresponding floating-point values in the vectors of the two source SIMD&FP registers, places the resulting floating-point values in a vector, and writes the vector to the destination SIMD&FP register.
func VmulxF64(v0 Float64X1, v1 Float64X1) Float64X1 {
	return C.vmulx_f64(v0, v1)
}

// Floating-point Multiply extended. This instruction multiplies corresponding floating-point values in the vectors of the two source SIMD&FP registers, places the resulting floating-point values in a vector, and writes the vector to the destination SIMD&FP register.
func VmulxF32(v0 Float32X2, v1 Float32X2) Float32X2 {
	return C.vmulx_f32(v0, v1)
}

// Floating-point Multiply extended. This instruction multiplies corresponding floating-point values in the vectors of the two source SIMD&FP registers, places the resulting floating-point values in a vector, and writes the vector to the destination SIMD&FP register.
func VmulxdF64(v0 Float64, v1 Float64) Float64 {
	return C.vmulxd_f64(v0, v1)
}

// Floating-point Multiply extended. This instruction multiplies corresponding floating-point values in the vectors of the two source SIMD&FP registers, places the resulting floating-point values in a vector, and writes the vector to the destination SIMD&FP register.
func VmulxsF32(v0 Float32, v1 Float32) Float32 {
	return C.vmulxs_f32(v0, v1)
}

// Floating-point Negate (vector). This instruction negates the value of each vector element in the source SIMD&FP register, writes the result to a vector, and writes the vector to the destination SIMD&FP register.
func VnegqF64(v0 Float64X2) Float64X2 {
	return C.vnegq_f64(v0)
}

// Negate (vector). This instruction reads each vector element from the source SIMD&FP register, negates each value, puts the result into a vector, and writes the vector to the destination SIMD&FP register.
func VnegqS64(v0 Int64X2) Int64X2 {
	return C.vnegq_s64(v0)
}

// Floating-point Negate (vector). This instruction negates the value of each vector element in the source SIMD&FP register, writes the result to a vector, and writes the vector to the destination SIMD&FP register.
func VnegF64(v0 Float64X1) Float64X1 {
	return C.vneg_f64(v0)
}

// Negate (vector). This instruction reads each vector element from the source SIMD&FP register, negates each value, puts the result into a vector, and writes the vector to the destination SIMD&FP register.
func VnegS64(v0 Int64X1) Int64X1 {
	return C.vneg_s64(v0)
}

// Negate (vector). This instruction reads each vector element from the source SIMD&FP register, negates each value, puts the result into a vector, and writes the vector to the destination SIMD&FP register.
func VnegdS64(v0 Int64) Int64 {
	return C.vnegd_s64(v0)
}

// Add Pairwise (vector). This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements from the concatenated vector, adds each pair of values together, places the result into a vector, and writes the vector to the destination SIMD&FP register.
func VpaddqU8(v0 Uint8X16, v1 Uint8X16) Uint8X16 {
	return C.vpaddq_u8(v0, v1)
}

// Add Pairwise (vector). This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements from the concatenated vector, adds each pair of values together, places the result into a vector, and writes the vector to the destination SIMD&FP register.
func VpaddqU32(v0 Uint32X4, v1 Uint32X4) Uint32X4 {
	return C.vpaddq_u32(v0, v1)
}

// Add Pairwise (vector). This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements from the concatenated vector, adds each pair of values together, places the result into a vector, and writes the vector to the destination SIMD&FP register.
func VpaddqU64(v0 Uint64X2, v1 Uint64X2) Uint64X2 {
	return C.vpaddq_u64(v0, v1)
}

// Add Pairwise (vector). This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements from the concatenated vector, adds each pair of values together, places the result into a vector, and writes the vector to the destination SIMD&FP register.
func VpaddqU16(v0 Uint16X8, v1 Uint16X8) Uint16X8 {
	return C.vpaddq_u16(v0, v1)
}

// Add Pairwise (vector). This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements from the concatenated vector, adds each pair of values together, places the result into a vector, and writes the vector to the destination SIMD&FP register.
func VpaddqS8(v0 Int8X16, v1 Int8X16) Int8X16 {
	return C.vpaddq_s8(v0, v1)
}

// Floating-point Add Pairwise (vector). This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements from the concatenated vector, adds each pair of values together, places the result into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are floating-point values.
func VpaddqF64(v0 Float64X2, v1 Float64X2) Float64X2 {
	return C.vpaddq_f64(v0, v1)
}

// Floating-point Add Pairwise (vector). This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements from the concatenated vector, adds each pair of values together, places the result into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are floating-point values.
func VpaddqF32(v0 Float32X4, v1 Float32X4) Float32X4 {
	return C.vpaddq_f32(v0, v1)
}

// Add Pairwise (vector). This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements from the concatenated vector, adds each pair of values together, places the result into a vector, and writes the vector to the destination SIMD&FP register.
func VpaddqS32(v0 Int32X4, v1 Int32X4) Int32X4 {
	return C.vpaddq_s32(v0, v1)
}

// Add Pairwise (vector). This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements from the concatenated vector, adds each pair of values together, places the result into a vector, and writes the vector to the destination SIMD&FP register.
func VpaddqS64(v0 Int64X2, v1 Int64X2) Int64X2 {
	return C.vpaddq_s64(v0, v1)
}

// Add Pairwise (vector). This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements from the concatenated vector, adds each pair of values together, places the result into a vector, and writes the vector to the destination SIMD&FP register.
func VpaddqS16(v0 Int16X8, v1 Int16X8) Int16X8 {
	return C.vpaddq_s16(v0, v1)
}

// Add Pairwise (vector). This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements from the concatenated vector, adds each pair of values together, places the result into a vector, and writes the vector to the destination SIMD&FP register.
func VpadddU64(v0 Uint64X2) Uint64 {
	return C.vpaddd_u64(v0)
}

// Floating-point Add Pairwise (vector). This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements from the concatenated vector, adds each pair of values together, places the result into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are floating-point values.
func VpadddF64(v0 Float64X2) Float64 {
	return C.vpaddd_f64(v0)
}

// Add Pairwise (vector). This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements from the concatenated vector, adds each pair of values together, places the result into a vector, and writes the vector to the destination SIMD&FP register.
func VpadddS64(v0 Int64X2) Int64 {
	return C.vpaddd_s64(v0)
}

// Floating-point Add Pairwise (vector). This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements from the concatenated vector, adds each pair of values together, places the result into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are floating-point values.
func VpaddsF32(v0 Float32X2) Float32 {
	return C.vpadds_f32(v0)
}

// Unsigned Maximum Pairwise. This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements in the two source SIMD&FP registers, writes the largest of each pair of unsigned integer values into a vector, and writes the vector to the destination SIMD&FP register.
func VpmaxqU8(v0 Uint8X16, v1 Uint8X16) Uint8X16 {
	return C.vpmaxq_u8(v0, v1)
}

// Unsigned Maximum Pairwise. This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements in the two source SIMD&FP registers, writes the largest of each pair of unsigned integer values into a vector, and writes the vector to the destination SIMD&FP register.
func VpmaxqU32(v0 Uint32X4, v1 Uint32X4) Uint32X4 {
	return C.vpmaxq_u32(v0, v1)
}

// Unsigned Maximum Pairwise. This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements in the two source SIMD&FP registers, writes the largest of each pair of unsigned integer values into a vector, and writes the vector to the destination SIMD&FP register.
func VpmaxqU16(v0 Uint16X8, v1 Uint16X8) Uint16X8 {
	return C.vpmaxq_u16(v0, v1)
}

// Signed Maximum Pairwise. This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements in the two source SIMD&FP registers, writes the largest of each pair of signed integer values into a vector, and writes the vector to the destination SIMD&FP register.
func VpmaxqS8(v0 Int8X16, v1 Int8X16) Int8X16 {
	return C.vpmaxq_s8(v0, v1)
}

// Floating-point Maximum Pairwise (vector). This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements from the concatenated vector, writes the larger of each pair of values into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are floating-point values.
func VpmaxqF64(v0 Float64X2, v1 Float64X2) Float64X2 {
	return C.vpmaxq_f64(v0, v1)
}

// Floating-point Maximum Pairwise (vector). This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements from the concatenated vector, writes the larger of each pair of values into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are floating-point values.
func VpmaxqF32(v0 Float32X4, v1 Float32X4) Float32X4 {
	return C.vpmaxq_f32(v0, v1)
}

// Signed Maximum Pairwise. This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements in the two source SIMD&FP registers, writes the largest of each pair of signed integer values into a vector, and writes the vector to the destination SIMD&FP register.
func VpmaxqS32(v0 Int32X4, v1 Int32X4) Int32X4 {
	return C.vpmaxq_s32(v0, v1)
}

// Signed Maximum Pairwise. This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements in the two source SIMD&FP registers, writes the largest of each pair of signed integer values into a vector, and writes the vector to the destination SIMD&FP register.
func VpmaxqS16(v0 Int16X8, v1 Int16X8) Int16X8 {
	return C.vpmaxq_s16(v0, v1)
}

// Floating-point Maximum Pairwise (vector). This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements from the concatenated vector, writes the larger of each pair of values into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are floating-point values.
func VpmaxqdF64(v0 Float64X2) Float64 {
	return C.vpmaxqd_f64(v0)
}

// Floating-point Maximum Pairwise (vector). This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements from the concatenated vector, writes the larger of each pair of values into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are floating-point values.
func VpmaxsF32(v0 Float32X2) Float32 {
	return C.vpmaxs_f32(v0)
}

// Floating-point Maximum Number Pairwise (vector). This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements in the two source SIMD&FP registers, writes the largest of each pair of values into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are floating-point values.
func VpmaxnmqF64(v0 Float64X2, v1 Float64X2) Float64X2 {
	return C.vpmaxnmq_f64(v0, v1)
}

// Floating-point Maximum Number Pairwise (vector). This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements in the two source SIMD&FP registers, writes the largest of each pair of values into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are floating-point values.
func VpmaxnmqF32(v0 Float32X4, v1 Float32X4) Float32X4 {
	return C.vpmaxnmq_f32(v0, v1)
}

// Floating-point Maximum Number Pairwise (vector). This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements in the two source SIMD&FP registers, writes the largest of each pair of values into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are floating-point values.
func VpmaxnmF32(v0 Float32X2, v1 Float32X2) Float32X2 {
	return C.vpmaxnm_f32(v0, v1)
}

// Floating-point Maximum Number Pairwise (vector). This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements in the two source SIMD&FP registers, writes the largest of each pair of values into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are floating-point values.
func VpmaxnmqdF64(v0 Float64X2) Float64 {
	return C.vpmaxnmqd_f64(v0)
}

// Floating-point Maximum Number Pairwise (vector). This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements in the two source SIMD&FP registers, writes the largest of each pair of values into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are floating-point values.
func VpmaxnmsF32(v0 Float32X2) Float32 {
	return C.vpmaxnms_f32(v0)
}

// Unsigned Minimum Pairwise. This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements in the two source SIMD&FP registers, writes the smallest of each pair of unsigned integer values into a vector, and writes the vector to the destination SIMD&FP register.
func VpminqU8(v0 Uint8X16, v1 Uint8X16) Uint8X16 {
	return C.vpminq_u8(v0, v1)
}

// Unsigned Minimum Pairwise. This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements in the two source SIMD&FP registers, writes the smallest of each pair of unsigned integer values into a vector, and writes the vector to the destination SIMD&FP register.
func VpminqU32(v0 Uint32X4, v1 Uint32X4) Uint32X4 {
	return C.vpminq_u32(v0, v1)
}

// Unsigned Minimum Pairwise. This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements in the two source SIMD&FP registers, writes the smallest of each pair of unsigned integer values into a vector, and writes the vector to the destination SIMD&FP register.
func VpminqU16(v0 Uint16X8, v1 Uint16X8) Uint16X8 {
	return C.vpminq_u16(v0, v1)
}

// Signed Minimum Pairwise. This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements in the two source SIMD&FP registers, writes the smallest of each pair of signed integer values into a vector, and writes the vector to the destination SIMD&FP register.
func VpminqS8(v0 Int8X16, v1 Int8X16) Int8X16 {
	return C.vpminq_s8(v0, v1)
}

// Floating-point Minimum Pairwise (vector). This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements from the concatenated vector, writes the smaller of each pair of values into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are floating-point values.
func VpminqF64(v0 Float64X2, v1 Float64X2) Float64X2 {
	return C.vpminq_f64(v0, v1)
}

// Floating-point Minimum Pairwise (vector). This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements from the concatenated vector, writes the smaller of each pair of values into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are floating-point values.
func VpminqF32(v0 Float32X4, v1 Float32X4) Float32X4 {
	return C.vpminq_f32(v0, v1)
}

// Signed Minimum Pairwise. This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements in the two source SIMD&FP registers, writes the smallest of each pair of signed integer values into a vector, and writes the vector to the destination SIMD&FP register.
func VpminqS32(v0 Int32X4, v1 Int32X4) Int32X4 {
	return C.vpminq_s32(v0, v1)
}

// Signed Minimum Pairwise. This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements in the two source SIMD&FP registers, writes the smallest of each pair of signed integer values into a vector, and writes the vector to the destination SIMD&FP register.
func VpminqS16(v0 Int16X8, v1 Int16X8) Int16X8 {
	return C.vpminq_s16(v0, v1)
}

// Floating-point Minimum Pairwise (vector). This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements from the concatenated vector, writes the smaller of each pair of values into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are floating-point values.
func VpminqdF64(v0 Float64X2) Float64 {
	return C.vpminqd_f64(v0)
}

// Floating-point Minimum Pairwise (vector). This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements from the concatenated vector, writes the smaller of each pair of values into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are floating-point values.
func VpminsF32(v0 Float32X2) Float32 {
	return C.vpmins_f32(v0)
}

// Floating-point Minimum Number Pairwise (vector). This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements in the two source SIMD&FP registers, writes the smallest of each pair of floating-point values into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are floating-point values.
func VpminnmqF64(v0 Float64X2, v1 Float64X2) Float64X2 {
	return C.vpminnmq_f64(v0, v1)
}

// Floating-point Minimum Number Pairwise (vector). This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements in the two source SIMD&FP registers, writes the smallest of each pair of floating-point values into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are floating-point values.
func VpminnmqF32(v0 Float32X4, v1 Float32X4) Float32X4 {
	return C.vpminnmq_f32(v0, v1)
}

// Floating-point Minimum Number Pairwise (vector). This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements in the two source SIMD&FP registers, writes the smallest of each pair of floating-point values into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are floating-point values.
func VpminnmF32(v0 Float32X2, v1 Float32X2) Float32X2 {
	return C.vpminnm_f32(v0, v1)
}

// Floating-point Minimum Number Pairwise (vector). This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements in the two source SIMD&FP registers, writes the smallest of each pair of floating-point values into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are floating-point values.
func VpminnmqdF64(v0 Float64X2) Float64 {
	return C.vpminnmqd_f64(v0)
}

// Floating-point Minimum Number Pairwise (vector). This instruction creates a vector by concatenating the vector elements of the first source SIMD&FP register after the vector elements of the second source SIMD&FP register, reads each pair of adjacent vector elements in the two source SIMD&FP registers, writes the smallest of each pair of floating-point values into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are floating-point values.
func VpminnmsF32(v0 Float32X2) Float32 {
	return C.vpminnms_f32(v0)
}

// Signed saturating Absolute value. This instruction reads each vector element from the source SIMD&FP register, puts the absolute value of the result into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are signed integer values.
func VqabsqS64(v0 Int64X2) Int64X2 {
	return C.vqabsq_s64(v0)
}

// Signed saturating Absolute value. This instruction reads each vector element from the source SIMD&FP register, puts the absolute value of the result into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are signed integer values.
func VqabsS64(v0 Int64X1) Int64X1 {
	return C.vqabs_s64(v0)
}

// Signed saturating Absolute value. This instruction reads each vector element from the source SIMD&FP register, puts the absolute value of the result into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are signed integer values.
func VqabsbS8(v0 Int8) Int8 {
	return C.vqabsb_s8(v0)
}

// Signed saturating Absolute value. This instruction reads each vector element from the source SIMD&FP register, puts the absolute value of the result into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are signed integer values.
func VqabssS32(v0 Int32) Int32 {
	return C.vqabss_s32(v0)
}

// Signed saturating Absolute value. This instruction reads each vector element from the source SIMD&FP register, puts the absolute value of the result into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are signed integer values.
func VqabsdS64(v0 Int64) Int64 {
	return C.vqabsd_s64(v0)
}

// Signed saturating Absolute value. This instruction reads each vector element from the source SIMD&FP register, puts the absolute value of the result into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are signed integer values.
func VqabshS16(v0 Int16) Int16 {
	return C.vqabsh_s16(v0)
}

// Unsigned saturating Add. This instruction adds the values of corresponding elements of the two source SIMD&FP registers, places the results into a vector, and writes the vector to the destination SIMD&FP register.
func VqaddbU8(v0 Uint8, v1 Uint8) Uint8 {
	return C.vqaddb_u8(v0, v1)
}

// Unsigned saturating Add. This instruction adds the values of corresponding elements of the two source SIMD&FP registers, places the results into a vector, and writes the vector to the destination SIMD&FP register.
func VqaddsU32(v0 Uint32, v1 Uint32) Uint32 {
	return C.vqadds_u32(v0, v1)
}

// Unsigned saturating Add. This instruction adds the values of corresponding elements of the two source SIMD&FP registers, places the results into a vector, and writes the vector to the destination SIMD&FP register.
func VqadddU64(v0 Uint64, v1 Uint64) Uint64 {
	return C.vqaddd_u64(v0, v1)
}

// Unsigned saturating Add. This instruction adds the values of corresponding elements of the two source SIMD&FP registers, places the results into a vector, and writes the vector to the destination SIMD&FP register.
func VqaddhU16(v0 Uint16, v1 Uint16) Uint16 {
	return C.vqaddh_u16(v0, v1)
}

// Signed saturating Add. This instruction adds the values of corresponding elements of the two source SIMD&FP registers, places the results into a vector, and writes the vector to the destination SIMD&FP register.
func VqaddbS8(v0 Int8, v1 Int8) Int8 {
	return C.vqaddb_s8(v0, v1)
}

// Signed saturating Add. This instruction adds the values of corresponding elements of the two source SIMD&FP registers, places the results into a vector, and writes the vector to the destination SIMD&FP register.
func VqaddsS32(v0 Int32, v1 Int32) Int32 {
	return C.vqadds_s32(v0, v1)
}

// Signed saturating Add. This instruction adds the values of corresponding elements of the two source SIMD&FP registers, places the results into a vector, and writes the vector to the destination SIMD&FP register.
func VqadddS64(v0 Int64, v1 Int64) Int64 {
	return C.vqaddd_s64(v0, v1)
}

// Signed saturating Add. This instruction adds the values of corresponding elements of the two source SIMD&FP registers, places the results into a vector, and writes the vector to the destination SIMD&FP register.
func VqaddhS16(v0 Int16, v1 Int16) Int16 {
	return C.vqaddh_s16(v0, v1)
}

// Signed saturating Doubling Multiply-Add Long. This instruction multiplies corresponding signed integer values in the lower or upper half of the vectors of the two source SIMD&FP registers, doubles the results, and accumulates the final results with the vector elements of the destination SIMD&FP register. The destination vector elements are twice as long as the elements that are multiplied.
func VqdmlalsS32(v0 Int64, v1 Int32, v2 Int32) Int64 {
	return C.vqdmlals_s32(v0, v1, v2)
}

// Signed saturating Doubling Multiply-Add Long. This instruction multiplies corresponding signed integer values in the lower or upper half of the vectors of the two source SIMD&FP registers, doubles the results, and accumulates the final results with the vector elements of the destination SIMD&FP register. The destination vector elements are twice as long as the elements that are multiplied.
func VqdmlalhS16(v0 Int32, v1 Int16, v2 Int16) Int32 {
	return C.vqdmlalh_s16(v0, v1, v2)
}

// Signed saturating Doubling Multiply-Add Long. This instruction multiplies corresponding signed integer values in the lower or upper half of the vectors of the two source SIMD&FP registers, doubles the results, and accumulates the final results with the vector elements of the destination SIMD&FP register. The destination vector elements are twice as long as the elements that are multiplied.
func VqdmlalHighS32(v0 Int64X2, v1 Int32X4, v2 Int32X4) Int64X2 {
	return C.vqdmlal_high_s32(v0, v1, v2)
}

// Signed saturating Doubling Multiply-Add Long. This instruction multiplies corresponding signed integer values in the lower or upper half of the vectors of the two source SIMD&FP registers, doubles the results, and accumulates the final results with the vector elements of the destination SIMD&FP register. The destination vector elements are twice as long as the elements that are multiplied.
func VqdmlalHighS16(v0 Int32X4, v1 Int16X8, v2 Int16X8) Int32X4 {
	return C.vqdmlal_high_s16(v0, v1, v2)
}

// Signed saturating Doubling Multiply-Add Long. This instruction multiplies corresponding signed integer values in the lower or upper half of the vectors of the two source SIMD&FP registers, doubles the results, and accumulates the final results with the vector elements of the destination SIMD&FP register. The destination vector elements are twice as long as the elements that are multiplied.
func VqdmlalHighNS32(v0 Int64X2, v1 Int32X4, v2 Int32) Int64X2 {
	return C.vqdmlal_high_n_s32(v0, v1, v2)
}

// Signed saturating Doubling Multiply-Add Long. This instruction multiplies corresponding signed integer values in the lower or upper half of the vectors of the two source SIMD&FP registers, doubles the results, and accumulates the final results with the vector elements of the destination SIMD&FP register. The destination vector elements are twice as long as the elements that are multiplied.
func VqdmlalHighNS16(v0 Int32X4, v1 Int16X8, v2 Int16) Int32X4 {
	return C.vqdmlal_high_n_s16(v0, v1, v2)
}

// Signed saturating Doubling Multiply-Subtract Long. This instruction multiplies corresponding signed integer values in the lower or upper half of the vectors of the two source SIMD&FP registers, doubles the results, and subtracts the final results from the vector elements of the destination SIMD&FP register. The destination vector elements are twice as long as the elements that are multiplied.
func VqdmlslsS32(v0 Int64, v1 Int32, v2 Int32) Int64 {
	return C.vqdmlsls_s32(v0, v1, v2)
}

// Signed saturating Doubling Multiply-Subtract Long. This instruction multiplies corresponding signed integer values in the lower or upper half of the vectors of the two source SIMD&FP registers, doubles the results, and subtracts the final results from the vector elements of the destination SIMD&FP register. The destination vector elements are twice as long as the elements that are multiplied.
func VqdmlslhS16(v0 Int32, v1 Int16, v2 Int16) Int32 {
	return C.vqdmlslh_s16(v0, v1, v2)
}

// Signed saturating Doubling Multiply-Subtract Long. This instruction multiplies corresponding signed integer values in the lower or upper half of the vectors of the two source SIMD&FP registers, doubles the results, and subtracts the final results from the vector elements of the destination SIMD&FP register. The destination vector elements are twice as long as the elements that are multiplied.
func VqdmlslHighS32(v0 Int64X2, v1 Int32X4, v2 Int32X4) Int64X2 {
	return C.vqdmlsl_high_s32(v0, v1, v2)
}

// Signed saturating Doubling Multiply-Subtract Long. This instruction multiplies corresponding signed integer values in the lower or upper half of the vectors of the two source SIMD&FP registers, doubles the results, and subtracts the final results from the vector elements of the destination SIMD&FP register. The destination vector elements are twice as long as the elements that are multiplied.
func VqdmlslHighS16(v0 Int32X4, v1 Int16X8, v2 Int16X8) Int32X4 {
	return C.vqdmlsl_high_s16(v0, v1, v2)
}

// Signed saturating Doubling Multiply-Subtract Long. This instruction multiplies corresponding signed integer values in the lower or upper half of the vectors of the two source SIMD&FP registers, doubles the results, and subtracts the final results from the vector elements of the destination SIMD&FP register. The destination vector elements are twice as long as the elements that are multiplied.
func VqdmlslHighNS32(v0 Int64X2, v1 Int32X4, v2 Int32) Int64X2 {
	return C.vqdmlsl_high_n_s32(v0, v1, v2)
}

// Signed saturating Doubling Multiply-Subtract Long. This instruction multiplies corresponding signed integer values in the lower or upper half of the vectors of the two source SIMD&FP registers, doubles the results, and subtracts the final results from the vector elements of the destination SIMD&FP register. The destination vector elements are twice as long as the elements that are multiplied.
func VqdmlslHighNS16(v0 Int32X4, v1 Int16X8, v2 Int16) Int32X4 {
	return C.vqdmlsl_high_n_s16(v0, v1, v2)
}

// Signed saturating Doubling Multiply returning High half. This instruction multiplies the values of corresponding elements of the two source SIMD&FP registers, doubles the results, places the most significant half of the final results into a vector, and writes the vector to the destination SIMD&FP register.
func VqdmulhsS32(v0 Int32, v1 Int32) Int32 {
	return C.vqdmulhs_s32(v0, v1)
}

// Signed saturating Doubling Multiply returning High half. This instruction multiplies the values of corresponding elements of the two source SIMD&FP registers, doubles the results, places the most significant half of the final results into a vector, and writes the vector to the destination SIMD&FP register.
func VqdmulhhS16(v0 Int16, v1 Int16) Int16 {
	return C.vqdmulhh_s16(v0, v1)
}

// Signed saturating Doubling Multiply Long. This instruction multiplies corresponding vector elements in the lower or upper half of the two source SIMD&FP registers, doubles the results, places the final results in a vector, and writes the vector to the destination SIMD&FP register.
func VqdmullsS32(v0 Int32, v1 Int32) Int64 {
	return C.vqdmulls_s32(v0, v1)
}

// Signed saturating Doubling Multiply Long. This instruction multiplies corresponding vector elements in the lower or upper half of the two source SIMD&FP registers, doubles the results, places the final results in a vector, and writes the vector to the destination SIMD&FP register.
func VqdmullhS16(v0 Int16, v1 Int16) Int32 {
	return C.vqdmullh_s16(v0, v1)
}

// Signed saturating Doubling Multiply Long. This instruction multiplies corresponding vector elements in the lower or upper half of the two source SIMD&FP registers, doubles the results, places the final results in a vector, and writes the vector to the destination SIMD&FP register.
func VqdmullHighS32(v0 Int32X4, v1 Int32X4) Int64X2 {
	return C.vqdmull_high_s32(v0, v1)
}

// Signed saturating Doubling Multiply Long. This instruction multiplies corresponding vector elements in the lower or upper half of the two source SIMD&FP registers, doubles the results, places the final results in a vector, and writes the vector to the destination SIMD&FP register.
func VqdmullHighS16(v0 Int16X8, v1 Int16X8) Int32X4 {
	return C.vqdmull_high_s16(v0, v1)
}

// Signed saturating Doubling Multiply Long. This instruction multiplies corresponding vector elements in the lower or upper half of the two source SIMD&FP registers, doubles the results, places the final results in a vector, and writes the vector to the destination SIMD&FP register.
func VqdmullHighNS32(v0 Int32X4, v1 Int32) Int64X2 {
	return C.vqdmull_high_n_s32(v0, v1)
}

// Signed saturating Doubling Multiply Long. This instruction multiplies corresponding vector elements in the lower or upper half of the two source SIMD&FP registers, doubles the results, places the final results in a vector, and writes the vector to the destination SIMD&FP register.
func VqdmullHighNS16(v0 Int16X8, v1 Int16) Int32X4 {
	return C.vqdmull_high_n_s16(v0, v1)
}

// Signed saturating extract Narrow. This instruction reads each vector element from the source SIMD&FP register, saturates the value to half the original width, places the result into a vector, and writes the vector to the lower or upper half of the destination SIMD&FP register. The destination vector elements are half as long as the source vector elements. All the values in this instruction are signed integer values.
func VqmovnsS32(v0 Int32) Int16 {
	return C.vqmovns_s32(v0)
}

// Signed saturating extract Narrow. This instruction reads each vector element from the source SIMD&FP register, saturates the value to half the original width, places the result into a vector, and writes the vector to the lower or upper half of the destination SIMD&FP register. The destination vector elements are half as long as the source vector elements. All the values in this instruction are signed integer values.
func VqmovndS64(v0 Int64) Int32 {
	return C.vqmovnd_s64(v0)
}

// Signed saturating extract Narrow. This instruction reads each vector element from the source SIMD&FP register, saturates the value to half the original width, places the result into a vector, and writes the vector to the lower or upper half of the destination SIMD&FP register. The destination vector elements are half as long as the source vector elements. All the values in this instruction are signed integer values.
func VqmovnhS16(v0 Int16) Int8 {
	return C.vqmovnh_s16(v0)
}

// Unsigned saturating extract Narrow. This instruction reads each vector element from the source SIMD&FP register, saturates each value to half the original width, places the result into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are unsigned integer values.
func VqmovnsU32(v0 Uint32) Uint16 {
	return C.vqmovns_u32(v0)
}

// Unsigned saturating extract Narrow. This instruction reads each vector element from the source SIMD&FP register, saturates each value to half the original width, places the result into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are unsigned integer values.
func VqmovndU64(v0 Uint64) Uint32 {
	return C.vqmovnd_u64(v0)
}

// Unsigned saturating extract Narrow. This instruction reads each vector element from the source SIMD&FP register, saturates each value to half the original width, places the result into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are unsigned integer values.
func VqmovnhU16(v0 Uint16) Uint8 {
	return C.vqmovnh_u16(v0)
}

// Unsigned saturating extract Narrow. This instruction reads each vector element from the source SIMD&FP register, saturates each value to half the original width, places the result into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are unsigned integer values.
func VqmovnHighU32(v0 Uint16X4, v1 Uint32X4) Uint16X8 {
	return C.vqmovn_high_u32(v0, v1)
}

// Unsigned saturating extract Narrow. This instruction reads each vector element from the source SIMD&FP register, saturates each value to half the original width, places the result into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are unsigned integer values.
func VqmovnHighU64(v0 Uint32X2, v1 Uint64X2) Uint32X4 {
	return C.vqmovn_high_u64(v0, v1)
}

// Unsigned saturating extract Narrow. This instruction reads each vector element from the source SIMD&FP register, saturates each value to half the original width, places the result into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are unsigned integer values.
func VqmovnHighU16(v0 Uint8X8, v1 Uint16X8) Uint8X16 {
	return C.vqmovn_high_u16(v0, v1)
}

// Signed saturating extract Narrow. This instruction reads each vector element from the source SIMD&FP register, saturates the value to half the original width, places the result into a vector, and writes the vector to the lower or upper half of the destination SIMD&FP register. The destination vector elements are half as long as the source vector elements. All the values in this instruction are signed integer values.
func VqmovnHighS32(v0 Int16X4, v1 Int32X4) Int16X8 {
	return C.vqmovn_high_s32(v0, v1)
}

// Signed saturating extract Narrow. This instruction reads each vector element from the source SIMD&FP register, saturates the value to half the original width, places the result into a vector, and writes the vector to the lower or upper half of the destination SIMD&FP register. The destination vector elements are half as long as the source vector elements. All the values in this instruction are signed integer values.
func VqmovnHighS64(v0 Int32X2, v1 Int64X2) Int32X4 {
	return C.vqmovn_high_s64(v0, v1)
}

// Signed saturating extract Narrow. This instruction reads each vector element from the source SIMD&FP register, saturates the value to half the original width, places the result into a vector, and writes the vector to the lower or upper half of the destination SIMD&FP register. The destination vector elements are half as long as the source vector elements. All the values in this instruction are signed integer values.
func VqmovnHighS16(v0 Int8X8, v1 Int16X8) Int8X16 {
	return C.vqmovn_high_s16(v0, v1)
}

// Signed saturating extract Unsigned Narrow. This instruction reads each signed integer value in the vector of the source SIMD&FP register, saturates the value to an unsigned integer value that is half the original width, places the result into a vector, and writes the vector to the lower or upper half of the destination SIMD&FP register. The destination vector elements are half as long as the source vector elements.
func VqmovunsS32(v0 Int32) Uint16 {
	return C.vqmovuns_s32(v0)
}

// Signed saturating extract Unsigned Narrow. This instruction reads each signed integer value in the vector of the source SIMD&FP register, saturates the value to an unsigned integer value that is half the original width, places the result into a vector, and writes the vector to the lower or upper half of the destination SIMD&FP register. The destination vector elements are half as long as the source vector elements.
func VqmovundS64(v0 Int64) Uint32 {
	return C.vqmovund_s64(v0)
}

// Signed saturating extract Unsigned Narrow. This instruction reads each signed integer value in the vector of the source SIMD&FP register, saturates the value to an unsigned integer value that is half the original width, places the result into a vector, and writes the vector to the lower or upper half of the destination SIMD&FP register. The destination vector elements are half as long as the source vector elements.
func VqmovunhS16(v0 Int16) Uint8 {
	return C.vqmovunh_s16(v0)
}

// Signed saturating extract Unsigned Narrow. This instruction reads each signed integer value in the vector of the source SIMD&FP register, saturates the value to an unsigned integer value that is half the original width, places the result into a vector, and writes the vector to the lower or upper half of the destination SIMD&FP register. The destination vector elements are half as long as the source vector elements.
func VqmovunHighS32(v0 Uint16X4, v1 Int32X4) Uint16X8 {
	return C.vqmovun_high_s32(v0, v1)
}

// Signed saturating extract Unsigned Narrow. This instruction reads each signed integer value in the vector of the source SIMD&FP register, saturates the value to an unsigned integer value that is half the original width, places the result into a vector, and writes the vector to the lower or upper half of the destination SIMD&FP register. The destination vector elements are half as long as the source vector elements.
func VqmovunHighS64(v0 Uint32X2, v1 Int64X2) Uint32X4 {
	return C.vqmovun_high_s64(v0, v1)
}

// Signed saturating extract Unsigned Narrow. This instruction reads each signed integer value in the vector of the source SIMD&FP register, saturates the value to an unsigned integer value that is half the original width, places the result into a vector, and writes the vector to the lower or upper half of the destination SIMD&FP register. The destination vector elements are half as long as the source vector elements.
func VqmovunHighS16(v0 Uint8X8, v1 Int16X8) Uint8X16 {
	return C.vqmovun_high_s16(v0, v1)
}

// Signed saturating Negate. This instruction reads each vector element from the source SIMD&FP register, negates each value, places the result into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are signed integer values.
func VqnegqS64(v0 Int64X2) Int64X2 {
	return C.vqnegq_s64(v0)
}

// Signed saturating Negate. This instruction reads each vector element from the source SIMD&FP register, negates each value, places the result into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are signed integer values.
func VqnegS64(v0 Int64X1) Int64X1 {
	return C.vqneg_s64(v0)
}

// Signed saturating Negate. This instruction reads each vector element from the source SIMD&FP register, negates each value, places the result into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are signed integer values.
func VqnegbS8(v0 Int8) Int8 {
	return C.vqnegb_s8(v0)
}

// Signed saturating Negate. This instruction reads each vector element from the source SIMD&FP register, negates each value, places the result into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are signed integer values.
func VqnegsS32(v0 Int32) Int32 {
	return C.vqnegs_s32(v0)
}

// Signed saturating Negate. This instruction reads each vector element from the source SIMD&FP register, negates each value, places the result into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are signed integer values.
func VqnegdS64(v0 Int64) Int64 {
	return C.vqnegd_s64(v0)
}

// Signed saturating Negate. This instruction reads each vector element from the source SIMD&FP register, negates each value, places the result into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are signed integer values.
func VqneghS16(v0 Int16) Int16 {
	return C.vqnegh_s16(v0)
}

// Signed saturating Rounding Doubling Multiply returning High half. This instruction multiplies the values of corresponding elements of the two source SIMD&FP registers, doubles the results, places the most significant half of the final results into a vector, and writes the vector to the destination SIMD&FP register.
func VqrdmulhsS32(v0 Int32, v1 Int32) Int32 {
	return C.vqrdmulhs_s32(v0, v1)
}

// Signed saturating Rounding Doubling Multiply returning High half. This instruction multiplies the values of corresponding elements of the two source SIMD&FP registers, doubles the results, places the most significant half of the final results into a vector, and writes the vector to the destination SIMD&FP register.
func VqrdmulhhS16(v0 Int16, v1 Int16) Int16 {
	return C.vqrdmulhh_s16(v0, v1)
}

// Unsigned saturating Rounding Shift Left (register). This instruction takes each vector element of the first source SIMD&FP register, shifts the vector element by a value from the least significant byte of the corresponding vector element of the second source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
func VqrshlbU8(v0 Uint8, v1 Int8) Uint8 {
	return C.vqrshlb_u8(v0, v1)
}

// Unsigned saturating Rounding Shift Left (register). This instruction takes each vector element of the first source SIMD&FP register, shifts the vector element by a value from the least significant byte of the corresponding vector element of the second source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
func VqrshlsU32(v0 Uint32, v1 Int32) Uint32 {
	return C.vqrshls_u32(v0, v1)
}

// Unsigned saturating Rounding Shift Left (register). This instruction takes each vector element of the first source SIMD&FP register, shifts the vector element by a value from the least significant byte of the corresponding vector element of the second source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
func VqrshldU64(v0 Uint64, v1 Int64) Uint64 {
	return C.vqrshld_u64(v0, v1)
}

// Unsigned saturating Rounding Shift Left (register). This instruction takes each vector element of the first source SIMD&FP register, shifts the vector element by a value from the least significant byte of the corresponding vector element of the second source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
func VqrshlhU16(v0 Uint16, v1 Int16) Uint16 {
	return C.vqrshlh_u16(v0, v1)
}

// Signed saturating Rounding Shift Left (register). This instruction takes each vector element in the first source SIMD&FP register, shifts it by a value from the least significant byte of the corresponding vector element of the second source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
func VqrshlbS8(v0 Int8, v1 Int8) Int8 {
	return C.vqrshlb_s8(v0, v1)
}

// Signed saturating Rounding Shift Left (register). This instruction takes each vector element in the first source SIMD&FP register, shifts it by a value from the least significant byte of the corresponding vector element of the second source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
func VqrshlsS32(v0 Int32, v1 Int32) Int32 {
	return C.vqrshls_s32(v0, v1)
}

// Signed saturating Rounding Shift Left (register). This instruction takes each vector element in the first source SIMD&FP register, shifts it by a value from the least significant byte of the corresponding vector element of the second source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
func VqrshldS64(v0 Int64, v1 Int64) Int64 {
	return C.vqrshld_s64(v0, v1)
}

// Signed saturating Rounding Shift Left (register). This instruction takes each vector element in the first source SIMD&FP register, shifts it by a value from the least significant byte of the corresponding vector element of the second source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
func VqrshlhS16(v0 Int16, v1 Int16) Int16 {
	return C.vqrshlh_s16(v0, v1)
}

// Unsigned saturating Shift Left (register). This instruction takes each element in the vector of the first source SIMD&FP register, shifts the element by a value from the least significant byte of the corresponding element of the second source SIMD&FP register, places the results in a vector, and writes the vector to the destination SIMD&FP register.
func VqshlbU8(v0 Uint8, v1 Int8) Uint8 {
	return C.vqshlb_u8(v0, v1)
}

// Unsigned saturating Shift Left (register). This instruction takes each element in the vector of the first source SIMD&FP register, shifts the element by a value from the least significant byte of the corresponding element of the second source SIMD&FP register, places the results in a vector, and writes the vector to the destination SIMD&FP register.
func VqshlsU32(v0 Uint32, v1 Int32) Uint32 {
	return C.vqshls_u32(v0, v1)
}

// Unsigned saturating Shift Left (register). This instruction takes each element in the vector of the first source SIMD&FP register, shifts the element by a value from the least significant byte of the corresponding element of the second source SIMD&FP register, places the results in a vector, and writes the vector to the destination SIMD&FP register.
func VqshldU64(v0 Uint64, v1 Int64) Uint64 {
	return C.vqshld_u64(v0, v1)
}

// Unsigned saturating Shift Left (register). This instruction takes each element in the vector of the first source SIMD&FP register, shifts the element by a value from the least significant byte of the corresponding element of the second source SIMD&FP register, places the results in a vector, and writes the vector to the destination SIMD&FP register.
func VqshlhU16(v0 Uint16, v1 Int16) Uint16 {
	return C.vqshlh_u16(v0, v1)
}

// Signed saturating Shift Left (register). This instruction takes each element in the vector of the first source SIMD&FP register, shifts each element by a value from the least significant byte of the corresponding element of the second source SIMD&FP register, places the results in a vector, and writes the vector to the destination SIMD&FP register.
func VqshlbS8(v0 Int8, v1 Int8) Int8 {
	return C.vqshlb_s8(v0, v1)
}

// Signed saturating Shift Left (register). This instruction takes each element in the vector of the first source SIMD&FP register, shifts each element by a value from the least significant byte of the corresponding element of the second source SIMD&FP register, places the results in a vector, and writes the vector to the destination SIMD&FP register.
func VqshlsS32(v0 Int32, v1 Int32) Int32 {
	return C.vqshls_s32(v0, v1)
}

// Signed saturating Shift Left (register). This instruction takes each element in the vector of the first source SIMD&FP register, shifts each element by a value from the least significant byte of the corresponding element of the second source SIMD&FP register, places the results in a vector, and writes the vector to the destination SIMD&FP register.
func VqshldS64(v0 Int64, v1 Int64) Int64 {
	return C.vqshld_s64(v0, v1)
}

// Signed saturating Shift Left (register). This instruction takes each element in the vector of the first source SIMD&FP register, shifts each element by a value from the least significant byte of the corresponding element of the second source SIMD&FP register, places the results in a vector, and writes the vector to the destination SIMD&FP register.
func VqshlhS16(v0 Int16, v1 Int16) Int16 {
	return C.vqshlh_s16(v0, v1)
}

// Unsigned saturating Subtract. This instruction subtracts the element values of the second source SIMD&FP register from the corresponding element values of the first source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
func VqsubbU8(v0 Uint8, v1 Uint8) Uint8 {
	return C.vqsubb_u8(v0, v1)
}

// Unsigned saturating Subtract. This instruction subtracts the element values of the second source SIMD&FP register from the corresponding element values of the first source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
func VqsubsU32(v0 Uint32, v1 Uint32) Uint32 {
	return C.vqsubs_u32(v0, v1)
}

// Unsigned saturating Subtract. This instruction subtracts the element values of the second source SIMD&FP register from the corresponding element values of the first source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
func VqsubdU64(v0 Uint64, v1 Uint64) Uint64 {
	return C.vqsubd_u64(v0, v1)
}

// Unsigned saturating Subtract. This instruction subtracts the element values of the second source SIMD&FP register from the corresponding element values of the first source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
func VqsubhU16(v0 Uint16, v1 Uint16) Uint16 {
	return C.vqsubh_u16(v0, v1)
}

// Signed saturating Subtract. This instruction subtracts the element values of the second source SIMD&FP register from the corresponding element values of the first source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
func VqsubbS8(v0 Int8, v1 Int8) Int8 {
	return C.vqsubb_s8(v0, v1)
}

// Signed saturating Subtract. This instruction subtracts the element values of the second source SIMD&FP register from the corresponding element values of the first source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
func VqsubsS32(v0 Int32, v1 Int32) Int32 {
	return C.vqsubs_s32(v0, v1)
}

// Signed saturating Subtract. This instruction subtracts the element values of the second source SIMD&FP register from the corresponding element values of the first source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
func VqsubdS64(v0 Int64, v1 Int64) Int64 {
	return C.vqsubd_s64(v0, v1)
}

// Signed saturating Subtract. This instruction subtracts the element values of the second source SIMD&FP register from the corresponding element values of the first source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register.
func VqsubhS16(v0 Int16, v1 Int16) Int16 {
	return C.vqsubh_s16(v0, v1)
}

// Table vector Lookup. This instruction reads each value from the vector elements in the index source SIMD&FP register, uses each result as an index to perform a lookup in a table of bytes that is described by one to four source table SIMD&FP registers, places the lookup result in a vector, and writes the vector to the destination SIMD&FP register. If an index is out of range for the table, the result for that lookup is 0. If more than one source register is used to describe the table, the first source register describes the lowest bytes of the table.
func Vqtbl1P8(v0 Poly8X16, v1 Uint8X8) Poly8X8 {
	return C.vqtbl1_p8(v0, v1)
}

// Table vector Lookup. This instruction reads each value from the vector elements in the index source SIMD&FP register, uses each result as an index to perform a lookup in a table of bytes that is described by one to four source table SIMD&FP registers, places the lookup result in a vector, and writes the vector to the destination SIMD&FP register. If an index is out of range for the table, the result for that lookup is 0. If more than one source register is used to describe the table, the first source register describes the lowest bytes of the table.
func Vqtbl1QP8(v0 Poly8X16, v1 Uint8X16) Poly8X16 {
	return C.vqtbl1q_p8(v0, v1)
}

// Table vector Lookup. This instruction reads each value from the vector elements in the index source SIMD&FP register, uses each result as an index to perform a lookup in a table of bytes that is described by one to four source table SIMD&FP registers, places the lookup result in a vector, and writes the vector to the destination SIMD&FP register. If an index is out of range for the table, the result for that lookup is 0. If more than one source register is used to describe the table, the first source register describes the lowest bytes of the table.
func Vqtbl1QU8(v0 Uint8X16, v1 Uint8X16) Uint8X16 {
	return C.vqtbl1q_u8(v0, v1)
}

// Table vector Lookup. This instruction reads each value from the vector elements in the index source SIMD&FP register, uses each result as an index to perform a lookup in a table of bytes that is described by one to four source table SIMD&FP registers, places the lookup result in a vector, and writes the vector to the destination SIMD&FP register. If an index is out of range for the table, the result for that lookup is 0. If more than one source register is used to describe the table, the first source register describes the lowest bytes of the table.
func Vqtbl1QS8(v0 Int8X16, v1 Uint8X16) Int8X16 {
	return C.vqtbl1q_s8(v0, v1)
}

// Table vector Lookup. This instruction reads each value from the vector elements in the index source SIMD&FP register, uses each result as an index to perform a lookup in a table of bytes that is described by one to four source table SIMD&FP registers, places the lookup result in a vector, and writes the vector to the destination SIMD&FP register. If an index is out of range for the table, the result for that lookup is 0. If more than one source register is used to describe the table, the first source register describes the lowest bytes of the table.
func Vqtbl1U8(v0 Uint8X16, v1 Uint8X8) Uint8X8 {
	return C.vqtbl1_u8(v0, v1)
}

// Table vector Lookup. This instruction reads each value from the vector elements in the index source SIMD&FP register, uses each result as an index to perform a lookup in a table of bytes that is described by one to four source table SIMD&FP registers, places the lookup result in a vector, and writes the vector to the destination SIMD&FP register. If an index is out of range for the table, the result for that lookup is 0. If more than one source register is used to describe the table, the first source register describes the lowest bytes of the table.
func Vqtbl1S8(v0 Int8X16, v1 Uint8X8) Int8X8 {
	return C.vqtbl1_s8(v0, v1)
}

// Table vector Lookup. This instruction reads each value from the vector elements in the index source SIMD&FP register, uses each result as an index to perform a lookup in a table of bytes that is described by one to four source table SIMD&FP registers, places the lookup result in a vector, and writes the vector to the destination SIMD&FP register. If an index is out of range for the table, the result for that lookup is 0. If more than one source register is used to describe the table, the first source register describes the lowest bytes of the table.
func Vqtbl2P8(v0 Poly8X16X2, v1 Uint8X8) Poly8X8 {
	return C.vqtbl2_p8(v0, v1)
}

// Table vector Lookup. This instruction reads each value from the vector elements in the index source SIMD&FP register, uses each result as an index to perform a lookup in a table of bytes that is described by one to four source table SIMD&FP registers, places the lookup result in a vector, and writes the vector to the destination SIMD&FP register. If an index is out of range for the table, the result for that lookup is 0. If more than one source register is used to describe the table, the first source register describes the lowest bytes of the table.
func Vqtbl2QP8(v0 Poly8X16X2, v1 Uint8X16) Poly8X16 {
	return C.vqtbl2q_p8(v0, v1)
}

// Table vector Lookup. This instruction reads each value from the vector elements in the index source SIMD&FP register, uses each result as an index to perform a lookup in a table of bytes that is described by one to four source table SIMD&FP registers, places the lookup result in a vector, and writes the vector to the destination SIMD&FP register. If an index is out of range for the table, the result for that lookup is 0. If more than one source register is used to describe the table, the first source register describes the lowest bytes of the table.
func Vqtbl2QU8(v0 Uint8X16X2, v1 Uint8X16) Uint8X16 {
	return C.vqtbl2q_u8(v0, v1)
}

// Table vector Lookup. This instruction reads each value from the vector elements in the index source SIMD&FP register, uses each result as an index to perform a lookup in a table of bytes that is described by one to four source table SIMD&FP registers, places the lookup result in a vector, and writes the vector to the destination SIMD&FP register. If an index is out of range for the table, the result for that lookup is 0. If more than one source register is used to describe the table, the first source register describes the lowest bytes of the table.
func Vqtbl2QS8(v0 Int8X16X2, v1 Uint8X16) Int8X16 {
	return C.vqtbl2q_s8(v0, v1)
}

// Table vector Lookup. This instruction reads each value from the vector elements in the index source SIMD&FP register, uses each result as an index to perform a lookup in a table of bytes that is described by one to four source table SIMD&FP registers, places the lookup result in a vector, and writes the vector to the destination SIMD&FP register. If an index is out of range for the table, the result for that lookup is 0. If more than one source register is used to describe the table, the first source register describes the lowest bytes of the table.
func Vqtbl2U8(v0 Uint8X16X2, v1 Uint8X8) Uint8X8 {
	return C.vqtbl2_u8(v0, v1)
}

// Table vector Lookup. This instruction reads each value from the vector elements in the index source SIMD&FP register, uses each result as an index to perform a lookup in a table of bytes that is described by one to four source table SIMD&FP registers, places the lookup result in a vector, and writes the vector to the destination SIMD&FP register. If an index is out of range for the table, the result for that lookup is 0. If more than one source register is used to describe the table, the first source register describes the lowest bytes of the table.
func Vqtbl2S8(v0 Int8X16X2, v1 Uint8X8) Int8X8 {
	return C.vqtbl2_s8(v0, v1)
}

// Table vector Lookup. This instruction reads each value from the vector elements in the index source SIMD&FP register, uses each result as an index to perform a lookup in a table of bytes that is described by one to four source table SIMD&FP registers, places the lookup result in a vector, and writes the vector to the destination SIMD&FP register. If an index is out of range for the table, the result for that lookup is 0. If more than one source register is used to describe the table, the first source register describes the lowest bytes of the table.
func Vqtbl3P8(v0 Poly8X16X3, v1 Uint8X8) Poly8X8 {
	return C.vqtbl3_p8(v0, v1)
}

// Table vector Lookup. This instruction reads each value from the vector elements in the index source SIMD&FP register, uses each result as an index to perform a lookup in a table of bytes that is described by one to four source table SIMD&FP registers, places the lookup result in a vector, and writes the vector to the destination SIMD&FP register. If an index is out of range for the table, the result for that lookup is 0. If more than one source register is used to describe the table, the first source register describes the lowest bytes of the table.
func Vqtbl3QP8(v0 Poly8X16X3, v1 Uint8X16) Poly8X16 {
	return C.vqtbl3q_p8(v0, v1)
}

// Table vector Lookup. This instruction reads each value from the vector elements in the index source SIMD&FP register, uses each result as an index to perform a lookup in a table of bytes that is described by one to four source table SIMD&FP registers, places the lookup result in a vector, and writes the vector to the destination SIMD&FP register. If an index is out of range for the table, the result for that lookup is 0. If more than one source register is used to describe the table, the first source register describes the lowest bytes of the table.
func Vqtbl3QU8(v0 Uint8X16X3, v1 Uint8X16) Uint8X16 {
	return C.vqtbl3q_u8(v0, v1)
}

// Table vector Lookup. This instruction reads each value from the vector elements in the index source SIMD&FP register, uses each result as an index to perform a lookup in a table of bytes that is described by one to four source table SIMD&FP registers, places the lookup result in a vector, and writes the vector to the destination SIMD&FP register. If an index is out of range for the table, the result for that lookup is 0. If more than one source register is used to describe the table, the first source register describes the lowest bytes of the table.
func Vqtbl3QS8(v0 Int8X16X3, v1 Uint8X16) Int8X16 {
	return C.vqtbl3q_s8(v0, v1)
}

// Table vector Lookup. This instruction reads each value from the vector elements in the index source SIMD&FP register, uses each result as an index to perform a lookup in a table of bytes that is described by one to four source table SIMD&FP registers, places the lookup result in a vector, and writes the vector to the destination SIMD&FP register. If an index is out of range for the table, the result for that lookup is 0. If more than one source register is used to describe the table, the first source register describes the lowest bytes of the table.
func Vqtbl3U8(v0 Uint8X16X3, v1 Uint8X8) Uint8X8 {
	return C.vqtbl3_u8(v0, v1)
}

// Table vector Lookup. This instruction reads each value from the vector elements in the index source SIMD&FP register, uses each result as an index to perform a lookup in a table of bytes that is described by one to four source table SIMD&FP registers, places the lookup result in a vector, and writes the vector to the destination SIMD&FP register. If an index is out of range for the table, the result for that lookup is 0. If more than one source register is used to describe the table, the first source register describes the lowest bytes of the table.
func Vqtbl3S8(v0 Int8X16X3, v1 Uint8X8) Int8X8 {
	return C.vqtbl3_s8(v0, v1)
}

// Table vector Lookup. This instruction reads each value from the vector elements in the index source SIMD&FP register, uses each result as an index to perform a lookup in a table of bytes that is described by one to four source table SIMD&FP registers, places the lookup result in a vector, and writes the vector to the destination SIMD&FP register. If an index is out of range for the table, the result for that lookup is 0. If more than one source register is used to describe the table, the first source register describes the lowest bytes of the table.
func Vqtbl4P8(v0 Poly8X16X4, v1 Uint8X8) Poly8X8 {
	return C.vqtbl4_p8(v0, v1)
}

// Table vector Lookup. This instruction reads each value from the vector elements in the index source SIMD&FP register, uses each result as an index to perform a lookup in a table of bytes that is described by one to four source table SIMD&FP registers, places the lookup result in a vector, and writes the vector to the destination SIMD&FP register. If an index is out of range for the table, the result for that lookup is 0. If more than one source register is used to describe the table, the first source register describes the lowest bytes of the table.
func Vqtbl4QP8(v0 Poly8X16X4, v1 Uint8X16) Poly8X16 {
	return C.vqtbl4q_p8(v0, v1)
}

// Table vector Lookup. This instruction reads each value from the vector elements in the index source SIMD&FP register, uses each result as an index to perform a lookup in a table of bytes that is described by one to four source table SIMD&FP registers, places the lookup result in a vector, and writes the vector to the destination SIMD&FP register. If an index is out of range for the table, the result for that lookup is 0. If more than one source register is used to describe the table, the first source register describes the lowest bytes of the table.
func Vqtbl4QU8(v0 Uint8X16X4, v1 Uint8X16) Uint8X16 {
	return C.vqtbl4q_u8(v0, v1)
}

// Table vector Lookup. This instruction reads each value from the vector elements in the index source SIMD&FP register, uses each result as an index to perform a lookup in a table of bytes that is described by one to four source table SIMD&FP registers, places the lookup result in a vector, and writes the vector to the destination SIMD&FP register. If an index is out of range for the table, the result for that lookup is 0. If more than one source register is used to describe the table, the first source register describes the lowest bytes of the table.
func Vqtbl4QS8(v0 Int8X16X4, v1 Uint8X16) Int8X16 {
	return C.vqtbl4q_s8(v0, v1)
}

// Table vector Lookup. This instruction reads each value from the vector elements in the index source SIMD&FP register, uses each result as an index to perform a lookup in a table of bytes that is described by one to four source table SIMD&FP registers, places the lookup result in a vector, and writes the vector to the destination SIMD&FP register. If an index is out of range for the table, the result for that lookup is 0. If more than one source register is used to describe the table, the first source register describes the lowest bytes of the table.
func Vqtbl4U8(v0 Uint8X16X4, v1 Uint8X8) Uint8X8 {
	return C.vqtbl4_u8(v0, v1)
}

// Table vector Lookup. This instruction reads each value from the vector elements in the index source SIMD&FP register, uses each result as an index to perform a lookup in a table of bytes that is described by one to four source table SIMD&FP registers, places the lookup result in a vector, and writes the vector to the destination SIMD&FP register. If an index is out of range for the table, the result for that lookup is 0. If more than one source register is used to describe the table, the first source register describes the lowest bytes of the table.
func Vqtbl4S8(v0 Int8X16X4, v1 Uint8X8) Int8X8 {
	return C.vqtbl4_s8(v0, v1)
}

// Table vector lookup extension. This instruction reads each value from the vector elements in the index source SIMD&FP register, uses each result as an index to perform a lookup in a table of bytes that is described by one to four source table SIMD&FP registers, places the lookup result in a vector, and writes the vector to the destination SIMD&FP register. If an index is out of range for the table, the existing value in the vector element of the destination register is left unchanged. If more than one source register is used to describe the table, the first source register describes the lowest bytes of the table.
func Vqtbx1P8(v0 Poly8X8, v1 Poly8X16, v2 Uint8X8) Poly8X8 {
	return C.vqtbx1_p8(v0, v1, v2)
}

// Table vector lookup extension. This instruction reads each value from the vector elements in the index source SIMD&FP register, uses each result as an index to perform a lookup in a table of bytes that is described by one to four source table SIMD&FP registers, places the lookup result in a vector, and writes the vector to the destination SIMD&FP register. If an index is out of range for the table, the existing value in the vector element of the destination register is left unchanged. If more than one source register is used to describe the table, the first source register describes the lowest bytes of the table.
func Vqtbx1QP8(v0 Poly8X16, v1 Poly8X16, v2 Uint8X16) Poly8X16 {
	return C.vqtbx1q_p8(v0, v1, v2)
}

// Table vector lookup extension. This instruction reads each value from the vector elements in the index source SIMD&FP register, uses each result as an index to perform a lookup in a table of bytes that is described by one to four source table SIMD&FP registers, places the lookup result in a vector, and writes the vector to the destination SIMD&FP register. If an index is out of range for the table, the existing value in the vector element of the destination register is left unchanged. If more than one source register is used to describe the table, the first source register describes the lowest bytes of the table.
func Vqtbx1QU8(v0 Uint8X16, v1 Uint8X16, v2 Uint8X16) Uint8X16 {
	return C.vqtbx1q_u8(v0, v1, v2)
}

// Table vector lookup extension. This instruction reads each value from the vector elements in the index source SIMD&FP register, uses each result as an index to perform a lookup in a table of bytes that is described by one to four source table SIMD&FP registers, places the lookup result in a vector, and writes the vector to the destination SIMD&FP register. If an index is out of range for the table, the existing value in the vector element of the destination register is left unchanged. If more than one source register is used to describe the table, the first source register describes the lowest bytes of the table.
func Vqtbx1QS8(v0 Int8X16, v1 Int8X16, v2 Uint8X16) Int8X16 {
	return C.vqtbx1q_s8(v0, v1, v2)
}

// Table vector lookup extension. This instruction reads each value from the vector elements in the index source SIMD&FP register, uses each result as an index to perform a lookup in a table of bytes that is described by one to four source table SIMD&FP registers, places the lookup result in a vector, and writes the vector to the destination SIMD&FP register. If an index is out of range for the table, the existing value in the vector element of the destination register is left unchanged. If more than one source register is used to describe the table, the first source register describes the lowest bytes of the table.
func Vqtbx1U8(v0 Uint8X8, v1 Uint8X16, v2 Uint8X8) Uint8X8 {
	return C.vqtbx1_u8(v0, v1, v2)
}

// Table vector lookup extension. This instruction reads each value from the vector elements in the index source SIMD&FP register, uses each result as an index to perform a lookup in a table of bytes that is described by one to four source table SIMD&FP registers, places the lookup result in a vector, and writes the vector to the destination SIMD&FP register. If an index is out of range for the table, the existing value in the vector element of the destination register is left unchanged. If more than one source register is used to describe the table, the first source register describes the lowest bytes of the table.
func Vqtbx1S8(v0 Int8X8, v1 Int8X16, v2 Uint8X8) Int8X8 {
	return C.vqtbx1_s8(v0, v1, v2)
}

// Table vector lookup extension. This instruction reads each value from the vector elements in the index source SIMD&FP register, uses each result as an index to perform a lookup in a table of bytes that is described by one to four source table SIMD&FP registers, places the lookup result in a vector, and writes the vector to the destination SIMD&FP register. If an index is out of range for the table, the existing value in the vector element of the destination register is left unchanged. If more than one source register is used to describe the table, the first source register describes the lowest bytes of the table.
func Vqtbx2P8(v0 Poly8X8, v1 Poly8X16X2, v2 Uint8X8) Poly8X8 {
	return C.vqtbx2_p8(v0, v1, v2)
}

// Table vector lookup extension. This instruction reads each value from the vector elements in the index source SIMD&FP register, uses each result as an index to perform a lookup in a table of bytes that is described by one to four source table SIMD&FP registers, places the lookup result in a vector, and writes the vector to the destination SIMD&FP register. If an index is out of range for the table, the existing value in the vector element of the destination register is left unchanged. If more than one source register is used to describe the table, the first source register describes the lowest bytes of the table.
func Vqtbx2QP8(v0 Poly8X16, v1 Poly8X16X2, v2 Uint8X16) Poly8X16 {
	return C.vqtbx2q_p8(v0, v1, v2)
}

// Table vector lookup extension. This instruction reads each value from the vector elements in the index source SIMD&FP register, uses each result as an index to perform a lookup in a table of bytes that is described by one to four source table SIMD&FP registers, places the lookup result in a vector, and writes the vector to the destination SIMD&FP register. If an index is out of range for the table, the existing value in the vector element of the destination register is left unchanged. If more than one source register is used to describe the table, the first source register describes the lowest bytes of the table.
func Vqtbx2QU8(v0 Uint8X16, v1 Uint8X16X2, v2 Uint8X16) Uint8X16 {
	return C.vqtbx2q_u8(v0, v1, v2)
}

// Table vector lookup extension. This instruction reads each value from the vector elements in the index source SIMD&FP register, uses each result as an index to perform a lookup in a table of bytes that is described by one to four source table SIMD&FP registers, places the lookup result in a vector, and writes the vector to the destination SIMD&FP register. If an index is out of range for the table, the existing value in the vector element of the destination register is left unchanged. If more than one source register is used to describe the table, the first source register describes the lowest bytes of the table.
func Vqtbx2QS8(v0 Int8X16, v1 Int8X16X2, v2 Uint8X16) Int8X16 {
	return C.vqtbx2q_s8(v0, v1, v2)
}

// Table vector lookup extension. This instruction reads each value from the vector elements in the index source SIMD&FP register, uses each result as an index to perform a lookup in a table of bytes that is described by one to four source table SIMD&FP registers, places the lookup result in a vector, and writes the vector to the destination SIMD&FP register. If an index is out of range for the table, the existing value in the vector element of the destination register is left unchanged. If more than one source register is used to describe the table, the first source register describes the lowest bytes of the table.
func Vqtbx2U8(v0 Uint8X8, v1 Uint8X16X2, v2 Uint8X8) Uint8X8 {
	return C.vqtbx2_u8(v0, v1, v2)
}

// Table vector lookup extension. This instruction reads each value from the vector elements in the index source SIMD&FP register, uses each result as an index to perform a lookup in a table of bytes that is described by one to four source table SIMD&FP registers, places the lookup result in a vector, and writes the vector to the destination SIMD&FP register. If an index is out of range for the table, the existing value in the vector element of the destination register is left unchanged. If more than one source register is used to describe the table, the first source register describes the lowest bytes of the table.
func Vqtbx2S8(v0 Int8X8, v1 Int8X16X2, v2 Uint8X8) Int8X8 {
	return C.vqtbx2_s8(v0, v1, v2)
}

// Table vector lookup extension. This instruction reads each value from the vector elements in the index source SIMD&FP register, uses each result as an index to perform a lookup in a table of bytes that is described by one to four source table SIMD&FP registers, places the lookup result in a vector, and writes the vector to the destination SIMD&FP register. If an index is out of range for the table, the existing value in the vector element of the destination register is left unchanged. If more than one source register is used to describe the table, the first source register describes the lowest bytes of the table.
func Vqtbx3P8(v0 Poly8X8, v1 Poly8X16X3, v2 Uint8X8) Poly8X8 {
	return C.vqtbx3_p8(v0, v1, v2)
}

// Table vector lookup extension. This instruction reads each value from the vector elements in the index source SIMD&FP register, uses each result as an index to perform a lookup in a table of bytes that is described by one to four source table SIMD&FP registers, places the lookup result in a vector, and writes the vector to the destination SIMD&FP register. If an index is out of range for the table, the existing value in the vector element of the destination register is left unchanged. If more than one source register is used to describe the table, the first source register describes the lowest bytes of the table.
func Vqtbx3QP8(v0 Poly8X16, v1 Poly8X16X3, v2 Uint8X16) Poly8X16 {
	return C.vqtbx3q_p8(v0, v1, v2)
}

// Table vector lookup extension. This instruction reads each value from the vector elements in the index source SIMD&FP register, uses each result as an index to perform a lookup in a table of bytes that is described by one to four source table SIMD&FP registers, places the lookup result in a vector, and writes the vector to the destination SIMD&FP register. If an index is out of range for the table, the existing value in the vector element of the destination register is left unchanged. If more than one source register is used to describe the table, the first source register describes the lowest bytes of the table.
func Vqtbx3QU8(v0 Uint8X16, v1 Uint8X16X3, v2 Uint8X16) Uint8X16 {
	return C.vqtbx3q_u8(v0, v1, v2)
}

// Table vector lookup extension. This instruction reads each value from the vector elements in the index source SIMD&FP register, uses each result as an index to perform a lookup in a table of bytes that is described by one to four source table SIMD&FP registers, places the lookup result in a vector, and writes the vector to the destination SIMD&FP register. If an index is out of range for the table, the existing value in the vector element of the destination register is left unchanged. If more than one source register is used to describe the table, the first source register describes the lowest bytes of the table.
func Vqtbx3QS8(v0 Int8X16, v1 Int8X16X3, v2 Uint8X16) Int8X16 {
	return C.vqtbx3q_s8(v0, v1, v2)
}

// Table vector lookup extension. This instruction reads each value from the vector elements in the index source SIMD&FP register, uses each result as an index to perform a lookup in a table of bytes that is described by one to four source table SIMD&FP registers, places the lookup result in a vector, and writes the vector to the destination SIMD&FP register. If an index is out of range for the table, the existing value in the vector element of the destination register is left unchanged. If more than one source register is used to describe the table, the first source register describes the lowest bytes of the table.
func Vqtbx3U8(v0 Uint8X8, v1 Uint8X16X3, v2 Uint8X8) Uint8X8 {
	return C.vqtbx3_u8(v0, v1, v2)
}

// Table vector lookup extension. This instruction reads each value from the vector elements in the index source SIMD&FP register, uses each result as an index to perform a lookup in a table of bytes that is described by one to four source table SIMD&FP registers, places the lookup result in a vector, and writes the vector to the destination SIMD&FP register. If an index is out of range for the table, the existing value in the vector element of the destination register is left unchanged. If more than one source register is used to describe the table, the first source register describes the lowest bytes of the table.
func Vqtbx3S8(v0 Int8X8, v1 Int8X16X3, v2 Uint8X8) Int8X8 {
	return C.vqtbx3_s8(v0, v1, v2)
}

// Table vector lookup extension. This instruction reads each value from the vector elements in the index source SIMD&FP register, uses each result as an index to perform a lookup in a table of bytes that is described by one to four source table SIMD&FP registers, places the lookup result in a vector, and writes the vector to the destination SIMD&FP register. If an index is out of range for the table, the existing value in the vector element of the destination register is left unchanged. If more than one source register is used to describe the table, the first source register describes the lowest bytes of the table.
func Vqtbx4P8(v0 Poly8X8, v1 Poly8X16X4, v2 Uint8X8) Poly8X8 {
	return C.vqtbx4_p8(v0, v1, v2)
}

// Table vector lookup extension. This instruction reads each value from the vector elements in the index source SIMD&FP register, uses each result as an index to perform a lookup in a table of bytes that is described by one to four source table SIMD&FP registers, places the lookup result in a vector, and writes the vector to the destination SIMD&FP register. If an index is out of range for the table, the existing value in the vector element of the destination register is left unchanged. If more than one source register is used to describe the table, the first source register describes the lowest bytes of the table.
func Vqtbx4QP8(v0 Poly8X16, v1 Poly8X16X4, v2 Uint8X16) Poly8X16 {
	return C.vqtbx4q_p8(v0, v1, v2)
}

// Table vector lookup extension. This instruction reads each value from the vector elements in the index source SIMD&FP register, uses each result as an index to perform a lookup in a table of bytes that is described by one to four source table SIMD&FP registers, places the lookup result in a vector, and writes the vector to the destination SIMD&FP register. If an index is out of range for the table, the existing value in the vector element of the destination register is left unchanged. If more than one source register is used to describe the table, the first source register describes the lowest bytes of the table.
func Vqtbx4QU8(v0 Uint8X16, v1 Uint8X16X4, v2 Uint8X16) Uint8X16 {
	return C.vqtbx4q_u8(v0, v1, v2)
}

// Table vector lookup extension. This instruction reads each value from the vector elements in the index source SIMD&FP register, uses each result as an index to perform a lookup in a table of bytes that is described by one to four source table SIMD&FP registers, places the lookup result in a vector, and writes the vector to the destination SIMD&FP register. If an index is out of range for the table, the existing value in the vector element of the destination register is left unchanged. If more than one source register is used to describe the table, the first source register describes the lowest bytes of the table.
func Vqtbx4QS8(v0 Int8X16, v1 Int8X16X4, v2 Uint8X16) Int8X16 {
	return C.vqtbx4q_s8(v0, v1, v2)
}

// Table vector lookup extension. This instruction reads each value from the vector elements in the index source SIMD&FP register, uses each result as an index to perform a lookup in a table of bytes that is described by one to four source table SIMD&FP registers, places the lookup result in a vector, and writes the vector to the destination SIMD&FP register. If an index is out of range for the table, the existing value in the vector element of the destination register is left unchanged. If more than one source register is used to describe the table, the first source register describes the lowest bytes of the table.
func Vqtbx4U8(v0 Uint8X8, v1 Uint8X16X4, v2 Uint8X8) Uint8X8 {
	return C.vqtbx4_u8(v0, v1, v2)
}

// Table vector lookup extension. This instruction reads each value from the vector elements in the index source SIMD&FP register, uses each result as an index to perform a lookup in a table of bytes that is described by one to four source table SIMD&FP registers, places the lookup result in a vector, and writes the vector to the destination SIMD&FP register. If an index is out of range for the table, the existing value in the vector element of the destination register is left unchanged. If more than one source register is used to describe the table, the first source register describes the lowest bytes of the table.
func Vqtbx4S8(v0 Int8X8, v1 Int8X16X4, v2 Uint8X8) Int8X8 {
	return C.vqtbx4_s8(v0, v1, v2)
}

// Rounding Add returning High Narrow. This instruction adds each vector element in the first source SIMD&FP register to the corresponding vector element in the second source SIMD&FP register, places the most significant half of the result into a vector, and writes the vector to the lower or upper half of the destination SIMD&FP register.
func VraddhnHighU32(v0 Uint16X4, v1 Uint32X4, v2 Uint32X4) Uint16X8 {
	return C.vraddhn_high_u32(v0, v1, v2)
}

// Rounding Add returning High Narrow. This instruction adds each vector element in the first source SIMD&FP register to the corresponding vector element in the second source SIMD&FP register, places the most significant half of the result into a vector, and writes the vector to the lower or upper half of the destination SIMD&FP register.
func VraddhnHighU64(v0 Uint32X2, v1 Uint64X2, v2 Uint64X2) Uint32X4 {
	return C.vraddhn_high_u64(v0, v1, v2)
}

// Rounding Add returning High Narrow. This instruction adds each vector element in the first source SIMD&FP register to the corresponding vector element in the second source SIMD&FP register, places the most significant half of the result into a vector, and writes the vector to the lower or upper half of the destination SIMD&FP register.
func VraddhnHighU16(v0 Uint8X8, v1 Uint16X8, v2 Uint16X8) Uint8X16 {
	return C.vraddhn_high_u16(v0, v1, v2)
}

// Rounding Add returning High Narrow. This instruction adds each vector element in the first source SIMD&FP register to the corresponding vector element in the second source SIMD&FP register, places the most significant half of the result into a vector, and writes the vector to the lower or upper half of the destination SIMD&FP register.
func VraddhnHighS32(v0 Int16X4, v1 Int32X4, v2 Int32X4) Int16X8 {
	return C.vraddhn_high_s32(v0, v1, v2)
}

// Rounding Add returning High Narrow. This instruction adds each vector element in the first source SIMD&FP register to the corresponding vector element in the second source SIMD&FP register, places the most significant half of the result into a vector, and writes the vector to the lower or upper half of the destination SIMD&FP register.
func VraddhnHighS64(v0 Int32X2, v1 Int64X2, v2 Int64X2) Int32X4 {
	return C.vraddhn_high_s64(v0, v1, v2)
}

// Rounding Add returning High Narrow. This instruction adds each vector element in the first source SIMD&FP register to the corresponding vector element in the second source SIMD&FP register, places the most significant half of the result into a vector, and writes the vector to the lower or upper half of the destination SIMD&FP register.
func VraddhnHighS16(v0 Int8X8, v1 Int16X8, v2 Int16X8) Int8X16 {
	return C.vraddhn_high_s16(v0, v1, v2)
}

// Reverse Bit order (vector). This instruction reads each vector element from the source SIMD&FP register, reverses the bits of the element, places the results into a vector, and writes the vector to the destination SIMD&FP register.
func VrbitP8(v0 Poly8X8) Poly8X8 {
	return C.vrbit_p8(v0)
}

// Reverse Bit order (vector). This instruction reads each vector element from the source SIMD&FP register, reverses the bits of the element, places the results into a vector, and writes the vector to the destination SIMD&FP register.
func VrbitqP8(v0 Poly8X16) Poly8X16 {
	return C.vrbitq_p8(v0)
}

// Reverse Bit order (vector). This instruction reads each vector element from the source SIMD&FP register, reverses the bits of the element, places the results into a vector, and writes the vector to the destination SIMD&FP register.
func VrbitqU8(v0 Uint8X16) Uint8X16 {
	return C.vrbitq_u8(v0)
}

// Reverse Bit order (vector). This instruction reads each vector element from the source SIMD&FP register, reverses the bits of the element, places the results into a vector, and writes the vector to the destination SIMD&FP register.
func VrbitqS8(v0 Int8X16) Int8X16 {
	return C.vrbitq_s8(v0)
}

// Reverse Bit order (vector). This instruction reads each vector element from the source SIMD&FP register, reverses the bits of the element, places the results into a vector, and writes the vector to the destination SIMD&FP register.
func VrbitU8(v0 Uint8X8) Uint8X8 {
	return C.vrbit_u8(v0)
}

// Reverse Bit order (vector). This instruction reads each vector element from the source SIMD&FP register, reverses the bits of the element, places the results into a vector, and writes the vector to the destination SIMD&FP register.
func VrbitS8(v0 Int8X8) Int8X8 {
	return C.vrbit_s8(v0)
}

// Floating-point Reciprocal Estimate. This instruction finds an approximate reciprocal estimate for each vector element in the source SIMD&FP register, places the result in a vector, and writes the vector to the destination SIMD&FP register.
func VrecpeqF64(v0 Float64X2) Float64X2 {
	return C.vrecpeq_f64(v0)
}

// Floating-point Reciprocal Estimate. This instruction finds an approximate reciprocal estimate for each vector element in the source SIMD&FP register, places the result in a vector, and writes the vector to the destination SIMD&FP register.
func VrecpeF64(v0 Float64X1) Float64X1 {
	return C.vrecpe_f64(v0)
}

// Floating-point Reciprocal Estimate. This instruction finds an approximate reciprocal estimate for each vector element in the source SIMD&FP register, places the result in a vector, and writes the vector to the destination SIMD&FP register.
func VrecpedF64(v0 Float64) Float64 {
	return C.vrecped_f64(v0)
}

// Floating-point Reciprocal Estimate. This instruction finds an approximate reciprocal estimate for each vector element in the source SIMD&FP register, places the result in a vector, and writes the vector to the destination SIMD&FP register.
func VrecpesF32(v0 Float32) Float32 {
	return C.vrecpes_f32(v0)
}

// Floating-point Reciprocal Step. This instruction multiplies the corresponding floating-point values in the vectors of the two source SIMD&FP registers, subtracts each of the products from 2.0, places the resulting floating-point values in a vector, and writes the vector to the destination SIMD&FP register.
func VrecpsqF64(v0 Float64X2, v1 Float64X2) Float64X2 {
	return C.vrecpsq_f64(v0, v1)
}

// Floating-point Reciprocal Step. This instruction multiplies the corresponding floating-point values in the vectors of the two source SIMD&FP registers, subtracts each of the products from 2.0, places the resulting floating-point values in a vector, and writes the vector to the destination SIMD&FP register.
func VrecpsF64(v0 Float64X1, v1 Float64X1) Float64X1 {
	return C.vrecps_f64(v0, v1)
}

// Floating-point Reciprocal Step. This instruction multiplies the corresponding floating-point values in the vectors of the two source SIMD&FP registers, subtracts each of the products from 2.0, places the resulting floating-point values in a vector, and writes the vector to the destination SIMD&FP register.
func VrecpsdF64(v0 Float64, v1 Float64) Float64 {
	return C.vrecpsd_f64(v0, v1)
}

// Floating-point Reciprocal Step. This instruction multiplies the corresponding floating-point values in the vectors of the two source SIMD&FP registers, subtracts each of the products from 2.0, places the resulting floating-point values in a vector, and writes the vector to the destination SIMD&FP register.
func VrecpssF32(v0 Float32, v1 Float32) Float32 {
	return C.vrecpss_f32(v0, v1)
}

// Floating-point Reciprocal exponent (scalar). This instruction finds an approximate reciprocal exponent for each vector element in the source SIMD&FP register, places the result in a vector, and writes the vector to the destination SIMD&FP register.
func VrecpxdF64(v0 Float64) Float64 {
	return C.vrecpxd_f64(v0)
}

// Floating-point Reciprocal exponent (scalar). This instruction finds an approximate reciprocal exponent for each vector element in the source SIMD&FP register, places the result in a vector, and writes the vector to the destination SIMD&FP register.
func VrecpxsF32(v0 Float32) Float32 {
	return C.vrecpxs_f32(v0)
}

// Unsigned Rounding Shift Left (register). This instruction takes each element in the vector of the first source SIMD&FP register, shifts the vector element by a value from the least significant byte of the corresponding element of the second source SIMD&FP register, places the results in a vector, and writes the vector to the destination SIMD&FP register.
func VrshldU64(v0 Uint64, v1 Int64) Uint64 {
	return C.vrshld_u64(v0, v1)
}

// Signed Rounding Shift Left (register). This instruction takes each signed integer value in the vector of the first source SIMD&FP register, shifts it by a value from the least significant byte of the corresponding element of the second source SIMD&FP register, places the results in a vector, and writes the vector to the destination SIMD&FP register.
func VrshldS64(v0 Int64, v1 Int64) Int64 {
	return C.vrshld_s64(v0, v1)
}

// Floating-point Reciprocal Square Root Estimate. This instruction calculates an approximate square root for each vector element in the source SIMD&FP register, places the result in a vector, and writes the vector to the destination SIMD&FP register.
func VrsqrteqF64(v0 Float64X2) Float64X2 {
	return C.vrsqrteq_f64(v0)
}

// Floating-point Reciprocal Square Root Estimate. This instruction calculates an approximate square root for each vector element in the source SIMD&FP register, places the result in a vector, and writes the vector to the destination SIMD&FP register.
func VrsqrteF64(v0 Float64X1) Float64X1 {
	return C.vrsqrte_f64(v0)
}

// Floating-point Reciprocal Square Root Estimate. This instruction calculates an approximate square root for each vector element in the source SIMD&FP register, places the result in a vector, and writes the vector to the destination SIMD&FP register.
func VrsqrtedF64(v0 Float64) Float64 {
	return C.vrsqrted_f64(v0)
}

// Floating-point Reciprocal Square Root Estimate. This instruction calculates an approximate square root for each vector element in the source SIMD&FP register, places the result in a vector, and writes the vector to the destination SIMD&FP register.
func VrsqrtesF32(v0 Float32) Float32 {
	return C.vrsqrtes_f32(v0)
}

// Floating-point Reciprocal Square Root Step. This instruction multiplies corresponding floating-point values in the vectors of the two source SIMD&FP registers, subtracts each of the products from 3.0, divides these results by 2.0, places the results into a vector, and writes the vector to the destination SIMD&FP register.
func VrsqrtsqF64(v0 Float64X2, v1 Float64X2) Float64X2 {
	return C.vrsqrtsq_f64(v0, v1)
}

// Floating-point Reciprocal Square Root Step. This instruction multiplies corresponding floating-point values in the vectors of the two source SIMD&FP registers, subtracts each of the products from 3.0, divides these results by 2.0, places the results into a vector, and writes the vector to the destination SIMD&FP register.
func VrsqrtsF64(v0 Float64X1, v1 Float64X1) Float64X1 {
	return C.vrsqrts_f64(v0, v1)
}

// Floating-point Reciprocal Square Root Step. This instruction multiplies corresponding floating-point values in the vectors of the two source SIMD&FP registers, subtracts each of the products from 3.0, divides these results by 2.0, places the results into a vector, and writes the vector to the destination SIMD&FP register.
func VrsqrtsdF64(v0 Float64, v1 Float64) Float64 {
	return C.vrsqrtsd_f64(v0, v1)
}

// Floating-point Reciprocal Square Root Step. This instruction multiplies corresponding floating-point values in the vectors of the two source SIMD&FP registers, subtracts each of the products from 3.0, divides these results by 2.0, places the results into a vector, and writes the vector to the destination SIMD&FP register.
func VrsqrtssF32(v0 Float32, v1 Float32) Float32 {
	return C.vrsqrtss_f32(v0, v1)
}

// Rounding Subtract returning High Narrow. This instruction subtracts each vector element of the second source SIMD&FP register from the corresponding vector element of the first source SIMD&FP register, places the most significant half of the result into a vector, and writes the vector to the lower or upper half of the destination SIMD&FP register.
func VrsubhnHighU32(v0 Uint16X4, v1 Uint32X4, v2 Uint32X4) Uint16X8 {
	return C.vrsubhn_high_u32(v0, v1, v2)
}

// Rounding Subtract returning High Narrow. This instruction subtracts each vector element of the second source SIMD&FP register from the corresponding vector element of the first source SIMD&FP register, places the most significant half of the result into a vector, and writes the vector to the lower or upper half of the destination SIMD&FP register.
func VrsubhnHighU64(v0 Uint32X2, v1 Uint64X2, v2 Uint64X2) Uint32X4 {
	return C.vrsubhn_high_u64(v0, v1, v2)
}

// Rounding Subtract returning High Narrow. This instruction subtracts each vector element of the second source SIMD&FP register from the corresponding vector element of the first source SIMD&FP register, places the most significant half of the result into a vector, and writes the vector to the lower or upper half of the destination SIMD&FP register.
func VrsubhnHighU16(v0 Uint8X8, v1 Uint16X8, v2 Uint16X8) Uint8X16 {
	return C.vrsubhn_high_u16(v0, v1, v2)
}

// Rounding Subtract returning High Narrow. This instruction subtracts each vector element of the second source SIMD&FP register from the corresponding vector element of the first source SIMD&FP register, places the most significant half of the result into a vector, and writes the vector to the lower or upper half of the destination SIMD&FP register.
func VrsubhnHighS32(v0 Int16X4, v1 Int32X4, v2 Int32X4) Int16X8 {
	return C.vrsubhn_high_s32(v0, v1, v2)
}

// Rounding Subtract returning High Narrow. This instruction subtracts each vector element of the second source SIMD&FP register from the corresponding vector element of the first source SIMD&FP register, places the most significant half of the result into a vector, and writes the vector to the lower or upper half of the destination SIMD&FP register.
func VrsubhnHighS64(v0 Int32X2, v1 Int64X2, v2 Int64X2) Int32X4 {
	return C.vrsubhn_high_s64(v0, v1, v2)
}

// Rounding Subtract returning High Narrow. This instruction subtracts each vector element of the second source SIMD&FP register from the corresponding vector element of the first source SIMD&FP register, places the most significant half of the result into a vector, and writes the vector to the lower or upper half of the destination SIMD&FP register.
func VrsubhnHighS16(v0 Int8X8, v1 Int16X8, v2 Int16X8) Int8X16 {
	return C.vrsubhn_high_s16(v0, v1, v2)
}

// Unsigned Shift Left (register). This instruction takes each element in the vector of the first source SIMD&FP register, shifts each element by a value from the least significant byte of the corresponding element of the second source SIMD&FP register, places the results in a vector, and writes the vector to the destination SIMD&FP register.
func VshldU64(v0 Uint64, v1 Int64) Uint64 {
	return C.vshld_u64(v0, v1)
}

// Signed Shift Left (register). This instruction takes each signed integer value in the vector of the first source SIMD&FP register, shifts each value by a value from the least significant byte of the corresponding element of the second source SIMD&FP register, places the results in a vector, and writes the vector to the destination SIMD&FP register.
func VshldS64(v0 Int64, v1 Int64) Int64 {
	return C.vshld_s64(v0, v1)
}

// Unsigned saturating Accumulate of Signed value. This instruction adds the signed integer values of the vector elements in the source SIMD&FP register to corresponding unsigned integer values of the vector elements in the destination SIMD&FP register, and accumulates the resulting unsigned integer values with the vector elements of the destination SIMD&FP register.
func VsqaddbU8(v0 Uint8, v1 Int8) Uint8 {
	return C.vsqaddb_u8(v0, v1)
}

// Unsigned saturating Accumulate of Signed value. This instruction adds the signed integer values of the vector elements in the source SIMD&FP register to corresponding unsigned integer values of the vector elements in the destination SIMD&FP register, and accumulates the resulting unsigned integer values with the vector elements of the destination SIMD&FP register.
func VsqaddsU32(v0 Uint32, v1 Int32) Uint32 {
	return C.vsqadds_u32(v0, v1)
}

// Unsigned saturating Accumulate of Signed value. This instruction adds the signed integer values of the vector elements in the source SIMD&FP register to corresponding unsigned integer values of the vector elements in the destination SIMD&FP register, and accumulates the resulting unsigned integer values with the vector elements of the destination SIMD&FP register.
func VsqadddU64(v0 Uint64, v1 Int64) Uint64 {
	return C.vsqaddd_u64(v0, v1)
}

// Unsigned saturating Accumulate of Signed value. This instruction adds the signed integer values of the vector elements in the source SIMD&FP register to corresponding unsigned integer values of the vector elements in the destination SIMD&FP register, and accumulates the resulting unsigned integer values with the vector elements of the destination SIMD&FP register.
func VsqaddhU16(v0 Uint16, v1 Int16) Uint16 {
	return C.vsqaddh_u16(v0, v1)
}

// Unsigned saturating Accumulate of Signed value. This instruction adds the signed integer values of the vector elements in the source SIMD&FP register to corresponding unsigned integer values of the vector elements in the destination SIMD&FP register, and accumulates the resulting unsigned integer values with the vector elements of the destination SIMD&FP register.
func VsqaddqU8(v0 Uint8X16, v1 Int8X16) Uint8X16 {
	return C.vsqaddq_u8(v0, v1)
}

// Unsigned saturating Accumulate of Signed value. This instruction adds the signed integer values of the vector elements in the source SIMD&FP register to corresponding unsigned integer values of the vector elements in the destination SIMD&FP register, and accumulates the resulting unsigned integer values with the vector elements of the destination SIMD&FP register.
func VsqaddqU32(v0 Uint32X4, v1 Int32X4) Uint32X4 {
	return C.vsqaddq_u32(v0, v1)
}

// Unsigned saturating Accumulate of Signed value. This instruction adds the signed integer values of the vector elements in the source SIMD&FP register to corresponding unsigned integer values of the vector elements in the destination SIMD&FP register, and accumulates the resulting unsigned integer values with the vector elements of the destination SIMD&FP register.
func VsqaddqU64(v0 Uint64X2, v1 Int64X2) Uint64X2 {
	return C.vsqaddq_u64(v0, v1)
}

// Unsigned saturating Accumulate of Signed value. This instruction adds the signed integer values of the vector elements in the source SIMD&FP register to corresponding unsigned integer values of the vector elements in the destination SIMD&FP register, and accumulates the resulting unsigned integer values with the vector elements of the destination SIMD&FP register.
func VsqaddqU16(v0 Uint16X8, v1 Int16X8) Uint16X8 {
	return C.vsqaddq_u16(v0, v1)
}

// Unsigned saturating Accumulate of Signed value. This instruction adds the signed integer values of the vector elements in the source SIMD&FP register to corresponding unsigned integer values of the vector elements in the destination SIMD&FP register, and accumulates the resulting unsigned integer values with the vector elements of the destination SIMD&FP register.
func VsqaddU8(v0 Uint8X8, v1 Int8X8) Uint8X8 {
	return C.vsqadd_u8(v0, v1)
}

// Unsigned saturating Accumulate of Signed value. This instruction adds the signed integer values of the vector elements in the source SIMD&FP register to corresponding unsigned integer values of the vector elements in the destination SIMD&FP register, and accumulates the resulting unsigned integer values with the vector elements of the destination SIMD&FP register.
func VsqaddU32(v0 Uint32X2, v1 Int32X2) Uint32X2 {
	return C.vsqadd_u32(v0, v1)
}

// Unsigned saturating Accumulate of Signed value. This instruction adds the signed integer values of the vector elements in the source SIMD&FP register to corresponding unsigned integer values of the vector elements in the destination SIMD&FP register, and accumulates the resulting unsigned integer values with the vector elements of the destination SIMD&FP register.
func VsqaddU64(v0 Uint64X1, v1 Int64X1) Uint64X1 {
	return C.vsqadd_u64(v0, v1)
}

// Unsigned saturating Accumulate of Signed value. This instruction adds the signed integer values of the vector elements in the source SIMD&FP register to corresponding unsigned integer values of the vector elements in the destination SIMD&FP register, and accumulates the resulting unsigned integer values with the vector elements of the destination SIMD&FP register.
func VsqaddU16(v0 Uint16X4, v1 Int16X4) Uint16X4 {
	return C.vsqadd_u16(v0, v1)
}

// Floating-point Square Root (vector). This instruction calculates the square root for each vector element in the source SIMD&FP register, places the result in a vector, and writes the vector to the destination SIMD&FP register.
func VsqrtqF64(v0 Float64X2) Float64X2 {
	return C.vsqrtq_f64(v0)
}

// Floating-point Square Root (vector). This instruction calculates the square root for each vector element in the source SIMD&FP register, places the result in a vector, and writes the vector to the destination SIMD&FP register.
func VsqrtqF32(v0 Float32X4) Float32X4 {
	return C.vsqrtq_f32(v0)
}

// Floating-point Square Root (vector). This instruction calculates the square root for each vector element in the source SIMD&FP register, places the result in a vector, and writes the vector to the destination SIMD&FP register.
func VsqrtF64(v0 Float64X1) Float64X1 {
	return C.vsqrt_f64(v0)
}

// Floating-point Square Root (vector). This instruction calculates the square root for each vector element in the source SIMD&FP register, places the result in a vector, and writes the vector to the destination SIMD&FP register.
func VsqrtF32(v0 Float32X2) Float32X2 {
	return C.vsqrt_f32(v0)
}

// Subtract (vector). This instruction subtracts each vector element in the second source SIMD&FP register from the corresponding vector element in the first source SIMD&FP register, places the result into a vector, and writes the vector to the destination SIMD&FP register.
func VsubdU64(v0 Uint64, v1 Uint64) Uint64 {
	return C.vsubd_u64(v0, v1)
}

// Subtract (vector). This instruction subtracts each vector element in the second source SIMD&FP register from the corresponding vector element in the first source SIMD&FP register, places the result into a vector, and writes the vector to the destination SIMD&FP register.
func VsubdS64(v0 Int64, v1 Int64) Int64 {
	return C.vsubd_s64(v0, v1)
}

// Floating-point Subtract (vector). This instruction subtracts the elements in the vector in the second source SIMD&FP register, from the corresponding elements in the vector in the first source SIMD&FP register, places each result into elements of a vector, and writes the vector to the destination SIMD&FP register.
func VsubqF64(v0 Float64X2, v1 Float64X2) Float64X2 {
	return C.vsubq_f64(v0, v1)
}

// Floating-point Subtract (vector). This instruction subtracts the elements in the vector in the second source SIMD&FP register, from the corresponding elements in the vector in the first source SIMD&FP register, places each result into elements of a vector, and writes the vector to the destination SIMD&FP register.
func VsubF64(v0 Float64X1, v1 Float64X1) Float64X1 {
	return C.vsub_f64(v0, v1)
}

// Subtract returning High Narrow. This instruction subtracts each vector element in the second source SIMD&FP register from the corresponding vector element in the first source SIMD&FP register, places the most significant half of the result into a vector, and writes the vector to the lower or upper half of the destination SIMD&FP register. All the values in this instruction are signed integer values.
func VsubhnHighU32(v0 Uint16X4, v1 Uint32X4, v2 Uint32X4) Uint16X8 {
	return C.vsubhn_high_u32(v0, v1, v2)
}

// Subtract returning High Narrow. This instruction subtracts each vector element in the second source SIMD&FP register from the corresponding vector element in the first source SIMD&FP register, places the most significant half of the result into a vector, and writes the vector to the lower or upper half of the destination SIMD&FP register. All the values in this instruction are signed integer values.
func VsubhnHighU64(v0 Uint32X2, v1 Uint64X2, v2 Uint64X2) Uint32X4 {
	return C.vsubhn_high_u64(v0, v1, v2)
}

// Subtract returning High Narrow. This instruction subtracts each vector element in the second source SIMD&FP register from the corresponding vector element in the first source SIMD&FP register, places the most significant half of the result into a vector, and writes the vector to the lower or upper half of the destination SIMD&FP register. All the values in this instruction are signed integer values.
func VsubhnHighU16(v0 Uint8X8, v1 Uint16X8, v2 Uint16X8) Uint8X16 {
	return C.vsubhn_high_u16(v0, v1, v2)
}

// Subtract returning High Narrow. This instruction subtracts each vector element in the second source SIMD&FP register from the corresponding vector element in the first source SIMD&FP register, places the most significant half of the result into a vector, and writes the vector to the lower or upper half of the destination SIMD&FP register. All the values in this instruction are signed integer values.
func VsubhnHighS32(v0 Int16X4, v1 Int32X4, v2 Int32X4) Int16X8 {
	return C.vsubhn_high_s32(v0, v1, v2)
}

// Subtract returning High Narrow. This instruction subtracts each vector element in the second source SIMD&FP register from the corresponding vector element in the first source SIMD&FP register, places the most significant half of the result into a vector, and writes the vector to the lower or upper half of the destination SIMD&FP register. All the values in this instruction are signed integer values.
func VsubhnHighS64(v0 Int32X2, v1 Int64X2, v2 Int64X2) Int32X4 {
	return C.vsubhn_high_s64(v0, v1, v2)
}

// Subtract returning High Narrow. This instruction subtracts each vector element in the second source SIMD&FP register from the corresponding vector element in the first source SIMD&FP register, places the most significant half of the result into a vector, and writes the vector to the lower or upper half of the destination SIMD&FP register. All the values in this instruction are signed integer values.
func VsubhnHighS16(v0 Int8X8, v1 Int16X8, v2 Int16X8) Int8X16 {
	return C.vsubhn_high_s16(v0, v1, v2)
}

// Unsigned Subtract Long. This instruction subtracts each vector element in the lower or upper half of the second source SIMD&FP register from the corresponding vector element of the first source SIMD&FP register, places the result into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are unsigned integer values. The destination vector elements are twice as long as the source vector elements.
func VsublHighU8(v0 Uint8X16, v1 Uint8X16) Uint16X8 {
	return C.vsubl_high_u8(v0, v1)
}

// Unsigned Subtract Long. This instruction subtracts each vector element in the lower or upper half of the second source SIMD&FP register from the corresponding vector element of the first source SIMD&FP register, places the result into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are unsigned integer values. The destination vector elements are twice as long as the source vector elements.
func VsublHighU32(v0 Uint32X4, v1 Uint32X4) Uint64X2 {
	return C.vsubl_high_u32(v0, v1)
}

// Unsigned Subtract Long. This instruction subtracts each vector element in the lower or upper half of the second source SIMD&FP register from the corresponding vector element of the first source SIMD&FP register, places the result into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are unsigned integer values. The destination vector elements are twice as long as the source vector elements.
func VsublHighU16(v0 Uint16X8, v1 Uint16X8) Uint32X4 {
	return C.vsubl_high_u16(v0, v1)
}

// Signed Subtract Long. This instruction subtracts each vector element in the lower or upper half of the second source SIMD&FP register from the corresponding vector element of the first source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are signed integer values. The destination vector elements are twice as long as the source vector elements.
func VsublHighS8(v0 Int8X16, v1 Int8X16) Int16X8 {
	return C.vsubl_high_s8(v0, v1)
}

// Signed Subtract Long. This instruction subtracts each vector element in the lower or upper half of the second source SIMD&FP register from the corresponding vector element of the first source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are signed integer values. The destination vector elements are twice as long as the source vector elements.
func VsublHighS32(v0 Int32X4, v1 Int32X4) Int64X2 {
	return C.vsubl_high_s32(v0, v1)
}

// Signed Subtract Long. This instruction subtracts each vector element in the lower or upper half of the second source SIMD&FP register from the corresponding vector element of the first source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register. All the values in this instruction are signed integer values. The destination vector elements are twice as long as the source vector elements.
func VsublHighS16(v0 Int16X8, v1 Int16X8) Int32X4 {
	return C.vsubl_high_s16(v0, v1)
}

// Unsigned Subtract Wide. This instruction subtracts each vector element of the second source SIMD&FP register from the corresponding vector element in the lower or upper half of the first source SIMD&FP register, places the result in a vector, and writes the vector to the SIMD&FP destination register. All the values in this instruction are signed integer values.
func VsubwHighU8(v0 Uint16X8, v1 Uint8X16) Uint16X8 {
	return C.vsubw_high_u8(v0, v1)
}

// Unsigned Subtract Wide. This instruction subtracts each vector element of the second source SIMD&FP register from the corresponding vector element in the lower or upper half of the first source SIMD&FP register, places the result in a vector, and writes the vector to the SIMD&FP destination register. All the values in this instruction are signed integer values.
func VsubwHighU32(v0 Uint64X2, v1 Uint32X4) Uint64X2 {
	return C.vsubw_high_u32(v0, v1)
}

// Unsigned Subtract Wide. This instruction subtracts each vector element of the second source SIMD&FP register from the corresponding vector element in the lower or upper half of the first source SIMD&FP register, places the result in a vector, and writes the vector to the SIMD&FP destination register. All the values in this instruction are signed integer values.
func VsubwHighU16(v0 Uint32X4, v1 Uint16X8) Uint32X4 {
	return C.vsubw_high_u16(v0, v1)
}

// Signed Subtract Wide. This instruction subtracts each vector element in the lower or upper half of the second source SIMD&FP register from the corresponding vector element in the first source SIMD&FP register, places the result in a vector, and writes the vector to the SIMD&FP destination register. All the values in this instruction are signed integer values.
func VsubwHighS8(v0 Int16X8, v1 Int8X16) Int16X8 {
	return C.vsubw_high_s8(v0, v1)
}

// Signed Subtract Wide. This instruction subtracts each vector element in the lower or upper half of the second source SIMD&FP register from the corresponding vector element in the first source SIMD&FP register, places the result in a vector, and writes the vector to the SIMD&FP destination register. All the values in this instruction are signed integer values.
func VsubwHighS32(v0 Int64X2, v1 Int32X4) Int64X2 {
	return C.vsubw_high_s32(v0, v1)
}

// Signed Subtract Wide. This instruction subtracts each vector element in the lower or upper half of the second source SIMD&FP register from the corresponding vector element in the first source SIMD&FP register, places the result in a vector, and writes the vector to the SIMD&FP destination register. All the values in this instruction are signed integer values.
func VsubwHighS16(v0 Int32X4, v1 Int16X8) Int32X4 {
	return C.vsubw_high_s16(v0, v1)
}

// Transpose vectors (primary). This instruction reads corresponding even-numbered vector elements from the two source SIMD&FP registers, starting at zero, places each result into consecutive elements of a vector, and writes the vector to the destination SIMD&FP register. Vector elements from the first source register are placed into even-numbered elements of the destination vector, starting at zero, while vector elements from the second source register are placed into odd-numbered elements of the destination vector.
func Vtrn1P8(v0 Poly8X8, v1 Poly8X8) Poly8X8 {
	return C.vtrn1_p8(v0, v1)
}

// Transpose vectors (primary). This instruction reads corresponding even-numbered vector elements from the two source SIMD&FP registers, starting at zero, places each result into consecutive elements of a vector, and writes the vector to the destination SIMD&FP register. Vector elements from the first source register are placed into even-numbered elements of the destination vector, starting at zero, while vector elements from the second source register are placed into odd-numbered elements of the destination vector.
func Vtrn1P16(v0 Poly16X4, v1 Poly16X4) Poly16X4 {
	return C.vtrn1_p16(v0, v1)
}

// Transpose vectors (primary). This instruction reads corresponding even-numbered vector elements from the two source SIMD&FP registers, starting at zero, places each result into consecutive elements of a vector, and writes the vector to the destination SIMD&FP register. Vector elements from the first source register are placed into even-numbered elements of the destination vector, starting at zero, while vector elements from the second source register are placed into odd-numbered elements of the destination vector.
func Vtrn1QP8(v0 Poly8X16, v1 Poly8X16) Poly8X16 {
	return C.vtrn1q_p8(v0, v1)
}

// Transpose vectors (primary). This instruction reads corresponding even-numbered vector elements from the two source SIMD&FP registers, starting at zero, places each result into consecutive elements of a vector, and writes the vector to the destination SIMD&FP register. Vector elements from the first source register are placed into even-numbered elements of the destination vector, starting at zero, while vector elements from the second source register are placed into odd-numbered elements of the destination vector.
func Vtrn1QP64(v0 Poly64X2, v1 Poly64X2) Poly64X2 {
	return C.vtrn1q_p64(v0, v1)
}

// Transpose vectors (primary). This instruction reads corresponding even-numbered vector elements from the two source SIMD&FP registers, starting at zero, places each result into consecutive elements of a vector, and writes the vector to the destination SIMD&FP register. Vector elements from the first source register are placed into even-numbered elements of the destination vector, starting at zero, while vector elements from the second source register are placed into odd-numbered elements of the destination vector.
func Vtrn1QP16(v0 Poly16X8, v1 Poly16X8) Poly16X8 {
	return C.vtrn1q_p16(v0, v1)
}

// Transpose vectors (primary). This instruction reads corresponding even-numbered vector elements from the two source SIMD&FP registers, starting at zero, places each result into consecutive elements of a vector, and writes the vector to the destination SIMD&FP register. Vector elements from the first source register are placed into even-numbered elements of the destination vector, starting at zero, while vector elements from the second source register are placed into odd-numbered elements of the destination vector.
func Vtrn1QU8(v0 Uint8X16, v1 Uint8X16) Uint8X16 {
	return C.vtrn1q_u8(v0, v1)
}

// Transpose vectors (primary). This instruction reads corresponding even-numbered vector elements from the two source SIMD&FP registers, starting at zero, places each result into consecutive elements of a vector, and writes the vector to the destination SIMD&FP register. Vector elements from the first source register are placed into even-numbered elements of the destination vector, starting at zero, while vector elements from the second source register are placed into odd-numbered elements of the destination vector.
func Vtrn1QU32(v0 Uint32X4, v1 Uint32X4) Uint32X4 {
	return C.vtrn1q_u32(v0, v1)
}

// Transpose vectors (primary). This instruction reads corresponding even-numbered vector elements from the two source SIMD&FP registers, starting at zero, places each result into consecutive elements of a vector, and writes the vector to the destination SIMD&FP register. Vector elements from the first source register are placed into even-numbered elements of the destination vector, starting at zero, while vector elements from the second source register are placed into odd-numbered elements of the destination vector.
func Vtrn1QU64(v0 Uint64X2, v1 Uint64X2) Uint64X2 {
	return C.vtrn1q_u64(v0, v1)
}

// Transpose vectors (primary). This instruction reads corresponding even-numbered vector elements from the two source SIMD&FP registers, starting at zero, places each result into consecutive elements of a vector, and writes the vector to the destination SIMD&FP register. Vector elements from the first source register are placed into even-numbered elements of the destination vector, starting at zero, while vector elements from the second source register are placed into odd-numbered elements of the destination vector.
func Vtrn1QU16(v0 Uint16X8, v1 Uint16X8) Uint16X8 {
	return C.vtrn1q_u16(v0, v1)
}

// Transpose vectors (primary). This instruction reads corresponding even-numbered vector elements from the two source SIMD&FP registers, starting at zero, places each result into consecutive elements of a vector, and writes the vector to the destination SIMD&FP register. Vector elements from the first source register are placed into even-numbered elements of the destination vector, starting at zero, while vector elements from the second source register are placed into odd-numbered elements of the destination vector.
func Vtrn1QS8(v0 Int8X16, v1 Int8X16) Int8X16 {
	return C.vtrn1q_s8(v0, v1)
}

// Transpose vectors (primary). This instruction reads corresponding even-numbered vector elements from the two source SIMD&FP registers, starting at zero, places each result into consecutive elements of a vector, and writes the vector to the destination SIMD&FP register. Vector elements from the first source register are placed into even-numbered elements of the destination vector, starting at zero, while vector elements from the second source register are placed into odd-numbered elements of the destination vector.
func Vtrn1QF64(v0 Float64X2, v1 Float64X2) Float64X2 {
	return C.vtrn1q_f64(v0, v1)
}

// Transpose vectors (primary). This instruction reads corresponding even-numbered vector elements from the two source SIMD&FP registers, starting at zero, places each result into consecutive elements of a vector, and writes the vector to the destination SIMD&FP register. Vector elements from the first source register are placed into even-numbered elements of the destination vector, starting at zero, while vector elements from the second source register are placed into odd-numbered elements of the destination vector.
func Vtrn1QF32(v0 Float32X4, v1 Float32X4) Float32X4 {
	return C.vtrn1q_f32(v0, v1)
}

// Transpose vectors (primary). This instruction reads corresponding even-numbered vector elements from the two source SIMD&FP registers, starting at zero, places each result into consecutive elements of a vector, and writes the vector to the destination SIMD&FP register. Vector elements from the first source register are placed into even-numbered elements of the destination vector, starting at zero, while vector elements from the second source register are placed into odd-numbered elements of the destination vector.
func Vtrn1QS32(v0 Int32X4, v1 Int32X4) Int32X4 {
	return C.vtrn1q_s32(v0, v1)
}

// Transpose vectors (primary). This instruction reads corresponding even-numbered vector elements from the two source SIMD&FP registers, starting at zero, places each result into consecutive elements of a vector, and writes the vector to the destination SIMD&FP register. Vector elements from the first source register are placed into even-numbered elements of the destination vector, starting at zero, while vector elements from the second source register are placed into odd-numbered elements of the destination vector.
func Vtrn1QS64(v0 Int64X2, v1 Int64X2) Int64X2 {
	return C.vtrn1q_s64(v0, v1)
}

// Transpose vectors (primary). This instruction reads corresponding even-numbered vector elements from the two source SIMD&FP registers, starting at zero, places each result into consecutive elements of a vector, and writes the vector to the destination SIMD&FP register. Vector elements from the first source register are placed into even-numbered elements of the destination vector, starting at zero, while vector elements from the second source register are placed into odd-numbered elements of the destination vector.
func Vtrn1QS16(v0 Int16X8, v1 Int16X8) Int16X8 {
	return C.vtrn1q_s16(v0, v1)
}

// Transpose vectors (primary). This instruction reads corresponding even-numbered vector elements from the two source SIMD&FP registers, starting at zero, places each result into consecutive elements of a vector, and writes the vector to the destination SIMD&FP register. Vector elements from the first source register are placed into even-numbered elements of the destination vector, starting at zero, while vector elements from the second source register are placed into odd-numbered elements of the destination vector.
func Vtrn1U8(v0 Uint8X8, v1 Uint8X8) Uint8X8 {
	return C.vtrn1_u8(v0, v1)
}

// Transpose vectors (primary). This instruction reads corresponding even-numbered vector elements from the two source SIMD&FP registers, starting at zero, places each result into consecutive elements of a vector, and writes the vector to the destination SIMD&FP register. Vector elements from the first source register are placed into even-numbered elements of the destination vector, starting at zero, while vector elements from the second source register are placed into odd-numbered elements of the destination vector.
func Vtrn1U32(v0 Uint32X2, v1 Uint32X2) Uint32X2 {
	return C.vtrn1_u32(v0, v1)
}

// Transpose vectors (primary). This instruction reads corresponding even-numbered vector elements from the two source SIMD&FP registers, starting at zero, places each result into consecutive elements of a vector, and writes the vector to the destination SIMD&FP register. Vector elements from the first source register are placed into even-numbered elements of the destination vector, starting at zero, while vector elements from the second source register are placed into odd-numbered elements of the destination vector.
func Vtrn1U16(v0 Uint16X4, v1 Uint16X4) Uint16X4 {
	return C.vtrn1_u16(v0, v1)
}

// Transpose vectors (primary). This instruction reads corresponding even-numbered vector elements from the two source SIMD&FP registers, starting at zero, places each result into consecutive elements of a vector, and writes the vector to the destination SIMD&FP register. Vector elements from the first source register are placed into even-numbered elements of the destination vector, starting at zero, while vector elements from the second source register are placed into odd-numbered elements of the destination vector.
func Vtrn1S8(v0 Int8X8, v1 Int8X8) Int8X8 {
	return C.vtrn1_s8(v0, v1)
}

// Transpose vectors (primary). This instruction reads corresponding even-numbered vector elements from the two source SIMD&FP registers, starting at zero, places each result into consecutive elements of a vector, and writes the vector to the destination SIMD&FP register. Vector elements from the first source register are placed into even-numbered elements of the destination vector, starting at zero, while vector elements from the second source register are placed into odd-numbered elements of the destination vector.
func Vtrn1F32(v0 Float32X2, v1 Float32X2) Float32X2 {
	return C.vtrn1_f32(v0, v1)
}

// Transpose vectors (primary). This instruction reads corresponding even-numbered vector elements from the two source SIMD&FP registers, starting at zero, places each result into consecutive elements of a vector, and writes the vector to the destination SIMD&FP register. Vector elements from the first source register are placed into even-numbered elements of the destination vector, starting at zero, while vector elements from the second source register are placed into odd-numbered elements of the destination vector.
func Vtrn1S32(v0 Int32X2, v1 Int32X2) Int32X2 {
	return C.vtrn1_s32(v0, v1)
}

// Transpose vectors (primary). This instruction reads corresponding even-numbered vector elements from the two source SIMD&FP registers, starting at zero, places each result into consecutive elements of a vector, and writes the vector to the destination SIMD&FP register. Vector elements from the first source register are placed into even-numbered elements of the destination vector, starting at zero, while vector elements from the second source register are placed into odd-numbered elements of the destination vector.
func Vtrn1S16(v0 Int16X4, v1 Int16X4) Int16X4 {
	return C.vtrn1_s16(v0, v1)
}

// Transpose vectors (secondary). This instruction reads corresponding odd-numbered vector elements from the two source SIMD&FP registers, places each result into consecutive elements of a vector, and writes the vector to the destination SIMD&FP register. Vector elements from the first source register are placed into even-numbered elements of the destination vector, starting at zero, while vector elements from the second source register are placed into odd-numbered elements of the destination vector.
func Vtrn2P8(v0 Poly8X8, v1 Poly8X8) Poly8X8 {
	return C.vtrn2_p8(v0, v1)
}

// Transpose vectors (secondary). This instruction reads corresponding odd-numbered vector elements from the two source SIMD&FP registers, places each result into consecutive elements of a vector, and writes the vector to the destination SIMD&FP register. Vector elements from the first source register are placed into even-numbered elements of the destination vector, starting at zero, while vector elements from the second source register are placed into odd-numbered elements of the destination vector.
func Vtrn2P16(v0 Poly16X4, v1 Poly16X4) Poly16X4 {
	return C.vtrn2_p16(v0, v1)
}

// Transpose vectors (secondary). This instruction reads corresponding odd-numbered vector elements from the two source SIMD&FP registers, places each result into consecutive elements of a vector, and writes the vector to the destination SIMD&FP register. Vector elements from the first source register are placed into even-numbered elements of the destination vector, starting at zero, while vector elements from the second source register are placed into odd-numbered elements of the destination vector.
func Vtrn2QP8(v0 Poly8X16, v1 Poly8X16) Poly8X16 {
	return C.vtrn2q_p8(v0, v1)
}

// Transpose vectors (secondary). This instruction reads corresponding odd-numbered vector elements from the two source SIMD&FP registers, places each result into consecutive elements of a vector, and writes the vector to the destination SIMD&FP register. Vector elements from the first source register are placed into even-numbered elements of the destination vector, starting at zero, while vector elements from the second source register are placed into odd-numbered elements of the destination vector.
func Vtrn2QP64(v0 Poly64X2, v1 Poly64X2) Poly64X2 {
	return C.vtrn2q_p64(v0, v1)
}

// Transpose vectors (secondary). This instruction reads corresponding odd-numbered vector elements from the two source SIMD&FP registers, places each result into consecutive elements of a vector, and writes the vector to the destination SIMD&FP register. Vector elements from the first source register are placed into even-numbered elements of the destination vector, starting at zero, while vector elements from the second source register are placed into odd-numbered elements of the destination vector.
func Vtrn2QP16(v0 Poly16X8, v1 Poly16X8) Poly16X8 {
	return C.vtrn2q_p16(v0, v1)
}

// Transpose vectors (secondary). This instruction reads corresponding odd-numbered vector elements from the two source SIMD&FP registers, places each result into consecutive elements of a vector, and writes the vector to the destination SIMD&FP register. Vector elements from the first source register are placed into even-numbered elements of the destination vector, starting at zero, while vector elements from the second source register are placed into odd-numbered elements of the destination vector.
func Vtrn2QU8(v0 Uint8X16, v1 Uint8X16) Uint8X16 {
	return C.vtrn2q_u8(v0, v1)
}

// Transpose vectors (secondary). This instruction reads corresponding odd-numbered vector elements from the two source SIMD&FP registers, places each result into consecutive elements of a vector, and writes the vector to the destination SIMD&FP register. Vector elements from the first source register are placed into even-numbered elements of the destination vector, starting at zero, while vector elements from the second source register are placed into odd-numbered elements of the destination vector.
func Vtrn2QU32(v0 Uint32X4, v1 Uint32X4) Uint32X4 {
	return C.vtrn2q_u32(v0, v1)
}

// Transpose vectors (secondary). This instruction reads corresponding odd-numbered vector elements from the two source SIMD&FP registers, places each result into consecutive elements of a vector, and writes the vector to the destination SIMD&FP register. Vector elements from the first source register are placed into even-numbered elements of the destination vector, starting at zero, while vector elements from the second source register are placed into odd-numbered elements of the destination vector.
func Vtrn2QU64(v0 Uint64X2, v1 Uint64X2) Uint64X2 {
	return C.vtrn2q_u64(v0, v1)
}

// Transpose vectors (secondary). This instruction reads corresponding odd-numbered vector elements from the two source SIMD&FP registers, places each result into consecutive elements of a vector, and writes the vector to the destination SIMD&FP register. Vector elements from the first source register are placed into even-numbered elements of the destination vector, starting at zero, while vector elements from the second source register are placed into odd-numbered elements of the destination vector.
func Vtrn2QU16(v0 Uint16X8, v1 Uint16X8) Uint16X8 {
	return C.vtrn2q_u16(v0, v1)
}

// Transpose vectors (secondary). This instruction reads corresponding odd-numbered vector elements from the two source SIMD&FP registers, places each result into consecutive elements of a vector, and writes the vector to the destination SIMD&FP register. Vector elements from the first source register are placed into even-numbered elements of the destination vector, starting at zero, while vector elements from the second source register are placed into odd-numbered elements of the destination vector.
func Vtrn2QS8(v0 Int8X16, v1 Int8X16) Int8X16 {
	return C.vtrn2q_s8(v0, v1)
}

// Transpose vectors (secondary). This instruction reads corresponding odd-numbered vector elements from the two source SIMD&FP registers, places each result into consecutive elements of a vector, and writes the vector to the destination SIMD&FP register. Vector elements from the first source register are placed into even-numbered elements of the destination vector, starting at zero, while vector elements from the second source register are placed into odd-numbered elements of the destination vector.
func Vtrn2QF64(v0 Float64X2, v1 Float64X2) Float64X2 {
	return C.vtrn2q_f64(v0, v1)
}

// Transpose vectors (secondary). This instruction reads corresponding odd-numbered vector elements from the two source SIMD&FP registers, places each result into consecutive elements of a vector, and writes the vector to the destination SIMD&FP register. Vector elements from the first source register are placed into even-numbered elements of the destination vector, starting at zero, while vector elements from the second source register are placed into odd-numbered elements of the destination vector.
func Vtrn2QF32(v0 Float32X4, v1 Float32X4) Float32X4 {
	return C.vtrn2q_f32(v0, v1)
}

// Transpose vectors (secondary). This instruction reads corresponding odd-numbered vector elements from the two source SIMD&FP registers, places each result into consecutive elements of a vector, and writes the vector to the destination SIMD&FP register. Vector elements from the first source register are placed into even-numbered elements of the destination vector, starting at zero, while vector elements from the second source register are placed into odd-numbered elements of the destination vector.
func Vtrn2QS32(v0 Int32X4, v1 Int32X4) Int32X4 {
	return C.vtrn2q_s32(v0, v1)
}

// Transpose vectors (secondary). This instruction reads corresponding odd-numbered vector elements from the two source SIMD&FP registers, places each result into consecutive elements of a vector, and writes the vector to the destination SIMD&FP register. Vector elements from the first source register are placed into even-numbered elements of the destination vector, starting at zero, while vector elements from the second source register are placed into odd-numbered elements of the destination vector.
func Vtrn2QS64(v0 Int64X2, v1 Int64X2) Int64X2 {
	return C.vtrn2q_s64(v0, v1)
}

// Transpose vectors (secondary). This instruction reads corresponding odd-numbered vector elements from the two source SIMD&FP registers, places each result into consecutive elements of a vector, and writes the vector to the destination SIMD&FP register. Vector elements from the first source register are placed into even-numbered elements of the destination vector, starting at zero, while vector elements from the second source register are placed into odd-numbered elements of the destination vector.
func Vtrn2QS16(v0 Int16X8, v1 Int16X8) Int16X8 {
	return C.vtrn2q_s16(v0, v1)
}

// Transpose vectors (secondary). This instruction reads corresponding odd-numbered vector elements from the two source SIMD&FP registers, places each result into consecutive elements of a vector, and writes the vector to the destination SIMD&FP register. Vector elements from the first source register are placed into even-numbered elements of the destination vector, starting at zero, while vector elements from the second source register are placed into odd-numbered elements of the destination vector.
func Vtrn2U8(v0 Uint8X8, v1 Uint8X8) Uint8X8 {
	return C.vtrn2_u8(v0, v1)
}

// Transpose vectors (secondary). This instruction reads corresponding odd-numbered vector elements from the two source SIMD&FP registers, places each result into consecutive elements of a vector, and writes the vector to the destination SIMD&FP register. Vector elements from the first source register are placed into even-numbered elements of the destination vector, starting at zero, while vector elements from the second source register are placed into odd-numbered elements of the destination vector.
func Vtrn2U32(v0 Uint32X2, v1 Uint32X2) Uint32X2 {
	return C.vtrn2_u32(v0, v1)
}

// Transpose vectors (secondary). This instruction reads corresponding odd-numbered vector elements from the two source SIMD&FP registers, places each result into consecutive elements of a vector, and writes the vector to the destination SIMD&FP register. Vector elements from the first source register are placed into even-numbered elements of the destination vector, starting at zero, while vector elements from the second source register are placed into odd-numbered elements of the destination vector.
func Vtrn2U16(v0 Uint16X4, v1 Uint16X4) Uint16X4 {
	return C.vtrn2_u16(v0, v1)
}

// Transpose vectors (secondary). This instruction reads corresponding odd-numbered vector elements from the two source SIMD&FP registers, places each result into consecutive elements of a vector, and writes the vector to the destination SIMD&FP register. Vector elements from the first source register are placed into even-numbered elements of the destination vector, starting at zero, while vector elements from the second source register are placed into odd-numbered elements of the destination vector.
func Vtrn2S8(v0 Int8X8, v1 Int8X8) Int8X8 {
	return C.vtrn2_s8(v0, v1)
}

// Transpose vectors (secondary). This instruction reads corresponding odd-numbered vector elements from the two source SIMD&FP registers, places each result into consecutive elements of a vector, and writes the vector to the destination SIMD&FP register. Vector elements from the first source register are placed into even-numbered elements of the destination vector, starting at zero, while vector elements from the second source register are placed into odd-numbered elements of the destination vector.
func Vtrn2F32(v0 Float32X2, v1 Float32X2) Float32X2 {
	return C.vtrn2_f32(v0, v1)
}

// Transpose vectors (secondary). This instruction reads corresponding odd-numbered vector elements from the two source SIMD&FP registers, places each result into consecutive elements of a vector, and writes the vector to the destination SIMD&FP register. Vector elements from the first source register are placed into even-numbered elements of the destination vector, starting at zero, while vector elements from the second source register are placed into odd-numbered elements of the destination vector.
func Vtrn2S32(v0 Int32X2, v1 Int32X2) Int32X2 {
	return C.vtrn2_s32(v0, v1)
}

// Transpose vectors (secondary). This instruction reads corresponding odd-numbered vector elements from the two source SIMD&FP registers, places each result into consecutive elements of a vector, and writes the vector to the destination SIMD&FP register. Vector elements from the first source register are placed into even-numbered elements of the destination vector, starting at zero, while vector elements from the second source register are placed into odd-numbered elements of the destination vector.
func Vtrn2S16(v0 Int16X4, v1 Int16X4) Int16X4 {
	return C.vtrn2_s16(v0, v1)
}

// Compare bitwise Test bits nonzero (vector). This instruction reads each vector element in the first source SIMD&FP register, performs an AND with the corresponding vector element in the second source SIMD&FP register, and if the result is not zero, sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
func VtstP64(v0 Poly64X1, v1 Poly64X1) Uint64X1 {
	return C.vtst_p64(v0, v1)
}

// Compare bitwise Test bits nonzero (vector). This instruction reads each vector element in the first source SIMD&FP register, performs an AND with the corresponding vector element in the second source SIMD&FP register, and if the result is not zero, sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
func VtstqP64(v0 Poly64X2, v1 Poly64X2) Uint64X2 {
	return C.vtstq_p64(v0, v1)
}

// Compare bitwise Test bits nonzero (vector). This instruction reads each vector element in the first source SIMD&FP register, performs an AND with the corresponding vector element in the second source SIMD&FP register, and if the result is not zero, sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
func VtstqU64(v0 Uint64X2, v1 Uint64X2) Uint64X2 {
	return C.vtstq_u64(v0, v1)
}

// Compare bitwise Test bits nonzero (vector). This instruction reads each vector element in the first source SIMD&FP register, performs an AND with the corresponding vector element in the second source SIMD&FP register, and if the result is not zero, sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
func VtstqS64(v0 Int64X2, v1 Int64X2) Uint64X2 {
	return C.vtstq_s64(v0, v1)
}

// Compare bitwise Test bits nonzero (vector). This instruction reads each vector element in the first source SIMD&FP register, performs an AND with the corresponding vector element in the second source SIMD&FP register, and if the result is not zero, sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
func VtstU64(v0 Uint64X1, v1 Uint64X1) Uint64X1 {
	return C.vtst_u64(v0, v1)
}

// Compare bitwise Test bits nonzero (vector). This instruction reads each vector element in the first source SIMD&FP register, performs an AND with the corresponding vector element in the second source SIMD&FP register, and if the result is not zero, sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
func VtstS64(v0 Int64X1, v1 Int64X1) Uint64X1 {
	return C.vtst_s64(v0, v1)
}

// Compare bitwise Test bits nonzero (vector). This instruction reads each vector element in the first source SIMD&FP register, performs an AND with the corresponding vector element in the second source SIMD&FP register, and if the result is not zero, sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
func VtstdU64(v0 Uint64, v1 Uint64) Uint64 {
	return C.vtstd_u64(v0, v1)
}

// Compare bitwise Test bits nonzero (vector). This instruction reads each vector element in the first source SIMD&FP register, performs an AND with the corresponding vector element in the second source SIMD&FP register, and if the result is not zero, sets every bit of the corresponding vector element in the destination SIMD&FP register to one, otherwise sets every bit of the corresponding vector element in the destination SIMD&FP register to zero.
func VtstdS64(v0 Int64, v1 Int64) Uint64 {
	return C.vtstd_s64(v0, v1)
}

// Signed saturating Accumulate of Unsigned value. This instruction adds the unsigned integer values of the vector elements in the source SIMD&FP register to corresponding signed integer values of the vector elements in the destination SIMD&FP register, and writes the resulting signed integer values to the destination SIMD&FP register.
func VuqaddbS8(v0 Int8, v1 Uint8) Int8 {
	return C.vuqaddb_s8(v0, v1)
}

// Signed saturating Accumulate of Unsigned value. This instruction adds the unsigned integer values of the vector elements in the source SIMD&FP register to corresponding signed integer values of the vector elements in the destination SIMD&FP register, and writes the resulting signed integer values to the destination SIMD&FP register.
func VuqaddsS32(v0 Int32, v1 Uint32) Int32 {
	return C.vuqadds_s32(v0, v1)
}

// Signed saturating Accumulate of Unsigned value. This instruction adds the unsigned integer values of the vector elements in the source SIMD&FP register to corresponding signed integer values of the vector elements in the destination SIMD&FP register, and writes the resulting signed integer values to the destination SIMD&FP register.
func VuqadddS64(v0 Int64, v1 Uint64) Int64 {
	return C.vuqaddd_s64(v0, v1)
}

// Signed saturating Accumulate of Unsigned value. This instruction adds the unsigned integer values of the vector elements in the source SIMD&FP register to corresponding signed integer values of the vector elements in the destination SIMD&FP register, and writes the resulting signed integer values to the destination SIMD&FP register.
func VuqaddhS16(v0 Int16, v1 Uint16) Int16 {
	return C.vuqaddh_s16(v0, v1)
}

// Signed saturating Accumulate of Unsigned value. This instruction adds the unsigned integer values of the vector elements in the source SIMD&FP register to corresponding signed integer values of the vector elements in the destination SIMD&FP register, and writes the resulting signed integer values to the destination SIMD&FP register.
func VuqaddqS8(v0 Int8X16, v1 Uint8X16) Int8X16 {
	return C.vuqaddq_s8(v0, v1)
}

// Signed saturating Accumulate of Unsigned value. This instruction adds the unsigned integer values of the vector elements in the source SIMD&FP register to corresponding signed integer values of the vector elements in the destination SIMD&FP register, and writes the resulting signed integer values to the destination SIMD&FP register.
func VuqaddqS32(v0 Int32X4, v1 Uint32X4) Int32X4 {
	return C.vuqaddq_s32(v0, v1)
}

// Signed saturating Accumulate of Unsigned value. This instruction adds the unsigned integer values of the vector elements in the source SIMD&FP register to corresponding signed integer values of the vector elements in the destination SIMD&FP register, and writes the resulting signed integer values to the destination SIMD&FP register.
func VuqaddqS64(v0 Int64X2, v1 Uint64X2) Int64X2 {
	return C.vuqaddq_s64(v0, v1)
}

// Signed saturating Accumulate of Unsigned value. This instruction adds the unsigned integer values of the vector elements in the source SIMD&FP register to corresponding signed integer values of the vector elements in the destination SIMD&FP register, and writes the resulting signed integer values to the destination SIMD&FP register.
func VuqaddqS16(v0 Int16X8, v1 Uint16X8) Int16X8 {
	return C.vuqaddq_s16(v0, v1)
}

// Signed saturating Accumulate of Unsigned value. This instruction adds the unsigned integer values of the vector elements in the source SIMD&FP register to corresponding signed integer values of the vector elements in the destination SIMD&FP register, and writes the resulting signed integer values to the destination SIMD&FP register.
func VuqaddS8(v0 Int8X8, v1 Uint8X8) Int8X8 {
	return C.vuqadd_s8(v0, v1)
}

// Signed saturating Accumulate of Unsigned value. This instruction adds the unsigned integer values of the vector elements in the source SIMD&FP register to corresponding signed integer values of the vector elements in the destination SIMD&FP register, and writes the resulting signed integer values to the destination SIMD&FP register.
func VuqaddS32(v0 Int32X2, v1 Uint32X2) Int32X2 {
	return C.vuqadd_s32(v0, v1)
}

// Signed saturating Accumulate of Unsigned value. This instruction adds the unsigned integer values of the vector elements in the source SIMD&FP register to corresponding signed integer values of the vector elements in the destination SIMD&FP register, and writes the resulting signed integer values to the destination SIMD&FP register.
func VuqaddS64(v0 Int64X1, v1 Uint64X1) Int64X1 {
	return C.vuqadd_s64(v0, v1)
}

// Signed saturating Accumulate of Unsigned value. This instruction adds the unsigned integer values of the vector elements in the source SIMD&FP register to corresponding signed integer values of the vector elements in the destination SIMD&FP register, and writes the resulting signed integer values to the destination SIMD&FP register.
func VuqaddS16(v0 Int16X4, v1 Uint16X4) Int16X4 {
	return C.vuqadd_s16(v0, v1)
}

// Unzip vectors (primary). This instruction reads corresponding even-numbered vector elements from the two source SIMD&FP registers, starting at zero, places the result from the first source register into consecutive elements in the lower half of a vector, and the result from the second source register into consecutive elements in the upper half of a vector, and writes the vector to the destination SIMD&FP register.
func Vuzp1P8(v0 Poly8X8, v1 Poly8X8) Poly8X8 {
	return C.vuzp1_p8(v0, v1)
}

// Unzip vectors (primary). This instruction reads corresponding even-numbered vector elements from the two source SIMD&FP registers, starting at zero, places the result from the first source register into consecutive elements in the lower half of a vector, and the result from the second source register into consecutive elements in the upper half of a vector, and writes the vector to the destination SIMD&FP register.
func Vuzp1P16(v0 Poly16X4, v1 Poly16X4) Poly16X4 {
	return C.vuzp1_p16(v0, v1)
}

// Unzip vectors (primary). This instruction reads corresponding even-numbered vector elements from the two source SIMD&FP registers, starting at zero, places the result from the first source register into consecutive elements in the lower half of a vector, and the result from the second source register into consecutive elements in the upper half of a vector, and writes the vector to the destination SIMD&FP register.
func Vuzp1QP8(v0 Poly8X16, v1 Poly8X16) Poly8X16 {
	return C.vuzp1q_p8(v0, v1)
}

// Unzip vectors (primary). This instruction reads corresponding even-numbered vector elements from the two source SIMD&FP registers, starting at zero, places the result from the first source register into consecutive elements in the lower half of a vector, and the result from the second source register into consecutive elements in the upper half of a vector, and writes the vector to the destination SIMD&FP register.
func Vuzp1QP64(v0 Poly64X2, v1 Poly64X2) Poly64X2 {
	return C.vuzp1q_p64(v0, v1)
}

// Unzip vectors (primary). This instruction reads corresponding even-numbered vector elements from the two source SIMD&FP registers, starting at zero, places the result from the first source register into consecutive elements in the lower half of a vector, and the result from the second source register into consecutive elements in the upper half of a vector, and writes the vector to the destination SIMD&FP register.
func Vuzp1QP16(v0 Poly16X8, v1 Poly16X8) Poly16X8 {
	return C.vuzp1q_p16(v0, v1)
}

// Unzip vectors (primary). This instruction reads corresponding even-numbered vector elements from the two source SIMD&FP registers, starting at zero, places the result from the first source register into consecutive elements in the lower half of a vector, and the result from the second source register into consecutive elements in the upper half of a vector, and writes the vector to the destination SIMD&FP register.
func Vuzp1QU8(v0 Uint8X16, v1 Uint8X16) Uint8X16 {
	return C.vuzp1q_u8(v0, v1)
}

// Unzip vectors (primary). This instruction reads corresponding even-numbered vector elements from the two source SIMD&FP registers, starting at zero, places the result from the first source register into consecutive elements in the lower half of a vector, and the result from the second source register into consecutive elements in the upper half of a vector, and writes the vector to the destination SIMD&FP register.
func Vuzp1QU32(v0 Uint32X4, v1 Uint32X4) Uint32X4 {
	return C.vuzp1q_u32(v0, v1)
}

// Unzip vectors (primary). This instruction reads corresponding even-numbered vector elements from the two source SIMD&FP registers, starting at zero, places the result from the first source register into consecutive elements in the lower half of a vector, and the result from the second source register into consecutive elements in the upper half of a vector, and writes the vector to the destination SIMD&FP register.
func Vuzp1QU64(v0 Uint64X2, v1 Uint64X2) Uint64X2 {
	return C.vuzp1q_u64(v0, v1)
}

// Unzip vectors (primary). This instruction reads corresponding even-numbered vector elements from the two source SIMD&FP registers, starting at zero, places the result from the first source register into consecutive elements in the lower half of a vector, and the result from the second source register into consecutive elements in the upper half of a vector, and writes the vector to the destination SIMD&FP register.
func Vuzp1QU16(v0 Uint16X8, v1 Uint16X8) Uint16X8 {
	return C.vuzp1q_u16(v0, v1)
}

// Unzip vectors (primary). This instruction reads corresponding even-numbered vector elements from the two source SIMD&FP registers, starting at zero, places the result from the first source register into consecutive elements in the lower half of a vector, and the result from the second source register into consecutive elements in the upper half of a vector, and writes the vector to the destination SIMD&FP register.
func Vuzp1QS8(v0 Int8X16, v1 Int8X16) Int8X16 {
	return C.vuzp1q_s8(v0, v1)
}

// Unzip vectors (primary). This instruction reads corresponding even-numbered vector elements from the two source SIMD&FP registers, starting at zero, places the result from the first source register into consecutive elements in the lower half of a vector, and the result from the second source register into consecutive elements in the upper half of a vector, and writes the vector to the destination SIMD&FP register.
func Vuzp1QF64(v0 Float64X2, v1 Float64X2) Float64X2 {
	return C.vuzp1q_f64(v0, v1)
}

// Unzip vectors (primary). This instruction reads corresponding even-numbered vector elements from the two source SIMD&FP registers, starting at zero, places the result from the first source register into consecutive elements in the lower half of a vector, and the result from the second source register into consecutive elements in the upper half of a vector, and writes the vector to the destination SIMD&FP register.
func Vuzp1QF32(v0 Float32X4, v1 Float32X4) Float32X4 {
	return C.vuzp1q_f32(v0, v1)
}

// Unzip vectors (primary). This instruction reads corresponding even-numbered vector elements from the two source SIMD&FP registers, starting at zero, places the result from the first source register into consecutive elements in the lower half of a vector, and the result from the second source register into consecutive elements in the upper half of a vector, and writes the vector to the destination SIMD&FP register.
func Vuzp1QS32(v0 Int32X4, v1 Int32X4) Int32X4 {
	return C.vuzp1q_s32(v0, v1)
}

// Unzip vectors (primary). This instruction reads corresponding even-numbered vector elements from the two source SIMD&FP registers, starting at zero, places the result from the first source register into consecutive elements in the lower half of a vector, and the result from the second source register into consecutive elements in the upper half of a vector, and writes the vector to the destination SIMD&FP register.
func Vuzp1QS64(v0 Int64X2, v1 Int64X2) Int64X2 {
	return C.vuzp1q_s64(v0, v1)
}

// Unzip vectors (primary). This instruction reads corresponding even-numbered vector elements from the two source SIMD&FP registers, starting at zero, places the result from the first source register into consecutive elements in the lower half of a vector, and the result from the second source register into consecutive elements in the upper half of a vector, and writes the vector to the destination SIMD&FP register.
func Vuzp1QS16(v0 Int16X8, v1 Int16X8) Int16X8 {
	return C.vuzp1q_s16(v0, v1)
}

// Unzip vectors (primary). This instruction reads corresponding even-numbered vector elements from the two source SIMD&FP registers, starting at zero, places the result from the first source register into consecutive elements in the lower half of a vector, and the result from the second source register into consecutive elements in the upper half of a vector, and writes the vector to the destination SIMD&FP register.
func Vuzp1U8(v0 Uint8X8, v1 Uint8X8) Uint8X8 {
	return C.vuzp1_u8(v0, v1)
}

// Unzip vectors (primary). This instruction reads corresponding even-numbered vector elements from the two source SIMD&FP registers, starting at zero, places the result from the first source register into consecutive elements in the lower half of a vector, and the result from the second source register into consecutive elements in the upper half of a vector, and writes the vector to the destination SIMD&FP register.
func Vuzp1U32(v0 Uint32X2, v1 Uint32X2) Uint32X2 {
	return C.vuzp1_u32(v0, v1)
}

// Unzip vectors (primary). This instruction reads corresponding even-numbered vector elements from the two source SIMD&FP registers, starting at zero, places the result from the first source register into consecutive elements in the lower half of a vector, and the result from the second source register into consecutive elements in the upper half of a vector, and writes the vector to the destination SIMD&FP register.
func Vuzp1U16(v0 Uint16X4, v1 Uint16X4) Uint16X4 {
	return C.vuzp1_u16(v0, v1)
}

// Unzip vectors (primary). This instruction reads corresponding even-numbered vector elements from the two source SIMD&FP registers, starting at zero, places the result from the first source register into consecutive elements in the lower half of a vector, and the result from the second source register into consecutive elements in the upper half of a vector, and writes the vector to the destination SIMD&FP register.
func Vuzp1S8(v0 Int8X8, v1 Int8X8) Int8X8 {
	return C.vuzp1_s8(v0, v1)
}

// Unzip vectors (primary). This instruction reads corresponding even-numbered vector elements from the two source SIMD&FP registers, starting at zero, places the result from the first source register into consecutive elements in the lower half of a vector, and the result from the second source register into consecutive elements in the upper half of a vector, and writes the vector to the destination SIMD&FP register.
func Vuzp1F32(v0 Float32X2, v1 Float32X2) Float32X2 {
	return C.vuzp1_f32(v0, v1)
}

// Unzip vectors (primary). This instruction reads corresponding even-numbered vector elements from the two source SIMD&FP registers, starting at zero, places the result from the first source register into consecutive elements in the lower half of a vector, and the result from the second source register into consecutive elements in the upper half of a vector, and writes the vector to the destination SIMD&FP register.
func Vuzp1S32(v0 Int32X2, v1 Int32X2) Int32X2 {
	return C.vuzp1_s32(v0, v1)
}

// Unzip vectors (primary). This instruction reads corresponding even-numbered vector elements from the two source SIMD&FP registers, starting at zero, places the result from the first source register into consecutive elements in the lower half of a vector, and the result from the second source register into consecutive elements in the upper half of a vector, and writes the vector to the destination SIMD&FP register.
func Vuzp1S16(v0 Int16X4, v1 Int16X4) Int16X4 {
	return C.vuzp1_s16(v0, v1)
}

// Unzip vectors (secondary). This instruction reads corresponding odd-numbered vector elements from the two source SIMD&FP registers, places the result from the first source register into consecutive elements in the lower half of a vector, and the result from the second source register into consecutive elements in the upper half of a vector, and writes the vector to the destination SIMD&FP register.
func Vuzp2P8(v0 Poly8X8, v1 Poly8X8) Poly8X8 {
	return C.vuzp2_p8(v0, v1)
}

// Unzip vectors (secondary). This instruction reads corresponding odd-numbered vector elements from the two source SIMD&FP registers, places the result from the first source register into consecutive elements in the lower half of a vector, and the result from the second source register into consecutive elements in the upper half of a vector, and writes the vector to the destination SIMD&FP register.
func Vuzp2P16(v0 Poly16X4, v1 Poly16X4) Poly16X4 {
	return C.vuzp2_p16(v0, v1)
}

// Unzip vectors (secondary). This instruction reads corresponding odd-numbered vector elements from the two source SIMD&FP registers, places the result from the first source register into consecutive elements in the lower half of a vector, and the result from the second source register into consecutive elements in the upper half of a vector, and writes the vector to the destination SIMD&FP register.
func Vuzp2QP8(v0 Poly8X16, v1 Poly8X16) Poly8X16 {
	return C.vuzp2q_p8(v0, v1)
}

// Unzip vectors (secondary). This instruction reads corresponding odd-numbered vector elements from the two source SIMD&FP registers, places the result from the first source register into consecutive elements in the lower half of a vector, and the result from the second source register into consecutive elements in the upper half of a vector, and writes the vector to the destination SIMD&FP register.
func Vuzp2QP64(v0 Poly64X2, v1 Poly64X2) Poly64X2 {
	return C.vuzp2q_p64(v0, v1)
}

// Unzip vectors (secondary). This instruction reads corresponding odd-numbered vector elements from the two source SIMD&FP registers, places the result from the first source register into consecutive elements in the lower half of a vector, and the result from the second source register into consecutive elements in the upper half of a vector, and writes the vector to the destination SIMD&FP register.
func Vuzp2QP16(v0 Poly16X8, v1 Poly16X8) Poly16X8 {
	return C.vuzp2q_p16(v0, v1)
}

// Unzip vectors (secondary). This instruction reads corresponding odd-numbered vector elements from the two source SIMD&FP registers, places the result from the first source register into consecutive elements in the lower half of a vector, and the result from the second source register into consecutive elements in the upper half of a vector, and writes the vector to the destination SIMD&FP register.
func Vuzp2QU8(v0 Uint8X16, v1 Uint8X16) Uint8X16 {
	return C.vuzp2q_u8(v0, v1)
}

// Unzip vectors (secondary). This instruction reads corresponding odd-numbered vector elements from the two source SIMD&FP registers, places the result from the first source register into consecutive elements in the lower half of a vector, and the result from the second source register into consecutive elements in the upper half of a vector, and writes the vector to the destination SIMD&FP register.
func Vuzp2QU32(v0 Uint32X4, v1 Uint32X4) Uint32X4 {
	return C.vuzp2q_u32(v0, v1)
}

// Unzip vectors (secondary). This instruction reads corresponding odd-numbered vector elements from the two source SIMD&FP registers, places the result from the first source register into consecutive elements in the lower half of a vector, and the result from the second source register into consecutive elements in the upper half of a vector, and writes the vector to the destination SIMD&FP register.
func Vuzp2QU64(v0 Uint64X2, v1 Uint64X2) Uint64X2 {
	return C.vuzp2q_u64(v0, v1)
}

// Unzip vectors (secondary). This instruction reads corresponding odd-numbered vector elements from the two source SIMD&FP registers, places the result from the first source register into consecutive elements in the lower half of a vector, and the result from the second source register into consecutive elements in the upper half of a vector, and writes the vector to the destination SIMD&FP register.
func Vuzp2QU16(v0 Uint16X8, v1 Uint16X8) Uint16X8 {
	return C.vuzp2q_u16(v0, v1)
}

// Unzip vectors (secondary). This instruction reads corresponding odd-numbered vector elements from the two source SIMD&FP registers, places the result from the first source register into consecutive elements in the lower half of a vector, and the result from the second source register into consecutive elements in the upper half of a vector, and writes the vector to the destination SIMD&FP register.
func Vuzp2QS8(v0 Int8X16, v1 Int8X16) Int8X16 {
	return C.vuzp2q_s8(v0, v1)
}

// Unzip vectors (secondary). This instruction reads corresponding odd-numbered vector elements from the two source SIMD&FP registers, places the result from the first source register into consecutive elements in the lower half of a vector, and the result from the second source register into consecutive elements in the upper half of a vector, and writes the vector to the destination SIMD&FP register.
func Vuzp2QF64(v0 Float64X2, v1 Float64X2) Float64X2 {
	return C.vuzp2q_f64(v0, v1)
}

// Unzip vectors (secondary). This instruction reads corresponding odd-numbered vector elements from the two source SIMD&FP registers, places the result from the first source register into consecutive elements in the lower half of a vector, and the result from the second source register into consecutive elements in the upper half of a vector, and writes the vector to the destination SIMD&FP register.
func Vuzp2QF32(v0 Float32X4, v1 Float32X4) Float32X4 {
	return C.vuzp2q_f32(v0, v1)
}

// Unzip vectors (secondary). This instruction reads corresponding odd-numbered vector elements from the two source SIMD&FP registers, places the result from the first source register into consecutive elements in the lower half of a vector, and the result from the second source register into consecutive elements in the upper half of a vector, and writes the vector to the destination SIMD&FP register.
func Vuzp2QS32(v0 Int32X4, v1 Int32X4) Int32X4 {
	return C.vuzp2q_s32(v0, v1)
}

// Unzip vectors (secondary). This instruction reads corresponding odd-numbered vector elements from the two source SIMD&FP registers, places the result from the first source register into consecutive elements in the lower half of a vector, and the result from the second source register into consecutive elements in the upper half of a vector, and writes the vector to the destination SIMD&FP register.
func Vuzp2QS64(v0 Int64X2, v1 Int64X2) Int64X2 {
	return C.vuzp2q_s64(v0, v1)
}

// Unzip vectors (secondary). This instruction reads corresponding odd-numbered vector elements from the two source SIMD&FP registers, places the result from the first source register into consecutive elements in the lower half of a vector, and the result from the second source register into consecutive elements in the upper half of a vector, and writes the vector to the destination SIMD&FP register.
func Vuzp2QS16(v0 Int16X8, v1 Int16X8) Int16X8 {
	return C.vuzp2q_s16(v0, v1)
}

// Unzip vectors (secondary). This instruction reads corresponding odd-numbered vector elements from the two source SIMD&FP registers, places the result from the first source register into consecutive elements in the lower half of a vector, and the result from the second source register into consecutive elements in the upper half of a vector, and writes the vector to the destination SIMD&FP register.
func Vuzp2U8(v0 Uint8X8, v1 Uint8X8) Uint8X8 {
	return C.vuzp2_u8(v0, v1)
}

// Unzip vectors (secondary). This instruction reads corresponding odd-numbered vector elements from the two source SIMD&FP registers, places the result from the first source register into consecutive elements in the lower half of a vector, and the result from the second source register into consecutive elements in the upper half of a vector, and writes the vector to the destination SIMD&FP register.
func Vuzp2U32(v0 Uint32X2, v1 Uint32X2) Uint32X2 {
	return C.vuzp2_u32(v0, v1)
}

// Unzip vectors (secondary). This instruction reads corresponding odd-numbered vector elements from the two source SIMD&FP registers, places the result from the first source register into consecutive elements in the lower half of a vector, and the result from the second source register into consecutive elements in the upper half of a vector, and writes the vector to the destination SIMD&FP register.
func Vuzp2U16(v0 Uint16X4, v1 Uint16X4) Uint16X4 {
	return C.vuzp2_u16(v0, v1)
}

// Unzip vectors (secondary). This instruction reads corresponding odd-numbered vector elements from the two source SIMD&FP registers, places the result from the first source register into consecutive elements in the lower half of a vector, and the result from the second source register into consecutive elements in the upper half of a vector, and writes the vector to the destination SIMD&FP register.
func Vuzp2S8(v0 Int8X8, v1 Int8X8) Int8X8 {
	return C.vuzp2_s8(v0, v1)
}

// Unzip vectors (secondary). This instruction reads corresponding odd-numbered vector elements from the two source SIMD&FP registers, places the result from the first source register into consecutive elements in the lower half of a vector, and the result from the second source register into consecutive elements in the upper half of a vector, and writes the vector to the destination SIMD&FP register.
func Vuzp2F32(v0 Float32X2, v1 Float32X2) Float32X2 {
	return C.vuzp2_f32(v0, v1)
}

// Unzip vectors (secondary). This instruction reads corresponding odd-numbered vector elements from the two source SIMD&FP registers, places the result from the first source register into consecutive elements in the lower half of a vector, and the result from the second source register into consecutive elements in the upper half of a vector, and writes the vector to the destination SIMD&FP register.
func Vuzp2S32(v0 Int32X2, v1 Int32X2) Int32X2 {
	return C.vuzp2_s32(v0, v1)
}

// Unzip vectors (secondary). This instruction reads corresponding odd-numbered vector elements from the two source SIMD&FP registers, places the result from the first source register into consecutive elements in the lower half of a vector, and the result from the second source register into consecutive elements in the upper half of a vector, and writes the vector to the destination SIMD&FP register.
func Vuzp2S16(v0 Int16X4, v1 Int16X4) Int16X4 {
	return C.vuzp2_s16(v0, v1)
}

// Zip vectors (primary). This instruction reads adjacent vector elements from the lower half of two source SIMD&FP registers as pairs, interleaves the pairs and places them into a vector, and writes the vector to the destination SIMD&FP register. The first pair from the first source register is placed into the two lowest vector elements, with subsequent pairs taken alternately from each source register.
func Vzip1P8(v0 Poly8X8, v1 Poly8X8) Poly8X8 {
	return C.vzip1_p8(v0, v1)
}

// Zip vectors (primary). This instruction reads adjacent vector elements from the lower half of two source SIMD&FP registers as pairs, interleaves the pairs and places them into a vector, and writes the vector to the destination SIMD&FP register. The first pair from the first source register is placed into the two lowest vector elements, with subsequent pairs taken alternately from each source register.
func Vzip1P16(v0 Poly16X4, v1 Poly16X4) Poly16X4 {
	return C.vzip1_p16(v0, v1)
}

// Zip vectors (primary). This instruction reads adjacent vector elements from the lower half of two source SIMD&FP registers as pairs, interleaves the pairs and places them into a vector, and writes the vector to the destination SIMD&FP register. The first pair from the first source register is placed into the two lowest vector elements, with subsequent pairs taken alternately from each source register.
func Vzip1QP8(v0 Poly8X16, v1 Poly8X16) Poly8X16 {
	return C.vzip1q_p8(v0, v1)
}

// Zip vectors (primary). This instruction reads adjacent vector elements from the lower half of two source SIMD&FP registers as pairs, interleaves the pairs and places them into a vector, and writes the vector to the destination SIMD&FP register. The first pair from the first source register is placed into the two lowest vector elements, with subsequent pairs taken alternately from each source register.
func Vzip1QP64(v0 Poly64X2, v1 Poly64X2) Poly64X2 {
	return C.vzip1q_p64(v0, v1)
}

// Zip vectors (primary). This instruction reads adjacent vector elements from the lower half of two source SIMD&FP registers as pairs, interleaves the pairs and places them into a vector, and writes the vector to the destination SIMD&FP register. The first pair from the first source register is placed into the two lowest vector elements, with subsequent pairs taken alternately from each source register.
func Vzip1QP16(v0 Poly16X8, v1 Poly16X8) Poly16X8 {
	return C.vzip1q_p16(v0, v1)
}

// Zip vectors (primary). This instruction reads adjacent vector elements from the lower half of two source SIMD&FP registers as pairs, interleaves the pairs and places them into a vector, and writes the vector to the destination SIMD&FP register. The first pair from the first source register is placed into the two lowest vector elements, with subsequent pairs taken alternately from each source register.
func Vzip1QU8(v0 Uint8X16, v1 Uint8X16) Uint8X16 {
	return C.vzip1q_u8(v0, v1)
}

// Zip vectors (primary). This instruction reads adjacent vector elements from the lower half of two source SIMD&FP registers as pairs, interleaves the pairs and places them into a vector, and writes the vector to the destination SIMD&FP register. The first pair from the first source register is placed into the two lowest vector elements, with subsequent pairs taken alternately from each source register.
func Vzip1QU32(v0 Uint32X4, v1 Uint32X4) Uint32X4 {
	return C.vzip1q_u32(v0, v1)
}

// Zip vectors (primary). This instruction reads adjacent vector elements from the lower half of two source SIMD&FP registers as pairs, interleaves the pairs and places them into a vector, and writes the vector to the destination SIMD&FP register. The first pair from the first source register is placed into the two lowest vector elements, with subsequent pairs taken alternately from each source register.
func Vzip1QU64(v0 Uint64X2, v1 Uint64X2) Uint64X2 {
	return C.vzip1q_u64(v0, v1)
}

// Zip vectors (primary). This instruction reads adjacent vector elements from the lower half of two source SIMD&FP registers as pairs, interleaves the pairs and places them into a vector, and writes the vector to the destination SIMD&FP register. The first pair from the first source register is placed into the two lowest vector elements, with subsequent pairs taken alternately from each source register.
func Vzip1QU16(v0 Uint16X8, v1 Uint16X8) Uint16X8 {
	return C.vzip1q_u16(v0, v1)
}

// Zip vectors (primary). This instruction reads adjacent vector elements from the lower half of two source SIMD&FP registers as pairs, interleaves the pairs and places them into a vector, and writes the vector to the destination SIMD&FP register. The first pair from the first source register is placed into the two lowest vector elements, with subsequent pairs taken alternately from each source register.
func Vzip1QS8(v0 Int8X16, v1 Int8X16) Int8X16 {
	return C.vzip1q_s8(v0, v1)
}

// Zip vectors (primary). This instruction reads adjacent vector elements from the lower half of two source SIMD&FP registers as pairs, interleaves the pairs and places them into a vector, and writes the vector to the destination SIMD&FP register. The first pair from the first source register is placed into the two lowest vector elements, with subsequent pairs taken alternately from each source register.
func Vzip1QF64(v0 Float64X2, v1 Float64X2) Float64X2 {
	return C.vzip1q_f64(v0, v1)
}

// Zip vectors (primary). This instruction reads adjacent vector elements from the lower half of two source SIMD&FP registers as pairs, interleaves the pairs and places them into a vector, and writes the vector to the destination SIMD&FP register. The first pair from the first source register is placed into the two lowest vector elements, with subsequent pairs taken alternately from each source register.
func Vzip1QF32(v0 Float32X4, v1 Float32X4) Float32X4 {
	return C.vzip1q_f32(v0, v1)
}

// Zip vectors (primary). This instruction reads adjacent vector elements from the lower half of two source SIMD&FP registers as pairs, interleaves the pairs and places them into a vector, and writes the vector to the destination SIMD&FP register. The first pair from the first source register is placed into the two lowest vector elements, with subsequent pairs taken alternately from each source register.
func Vzip1QS32(v0 Int32X4, v1 Int32X4) Int32X4 {
	return C.vzip1q_s32(v0, v1)
}

// Zip vectors (primary). This instruction reads adjacent vector elements from the lower half of two source SIMD&FP registers as pairs, interleaves the pairs and places them into a vector, and writes the vector to the destination SIMD&FP register. The first pair from the first source register is placed into the two lowest vector elements, with subsequent pairs taken alternately from each source register.
func Vzip1QS64(v0 Int64X2, v1 Int64X2) Int64X2 {
	return C.vzip1q_s64(v0, v1)
}

// Zip vectors (primary). This instruction reads adjacent vector elements from the lower half of two source SIMD&FP registers as pairs, interleaves the pairs and places them into a vector, and writes the vector to the destination SIMD&FP register. The first pair from the first source register is placed into the two lowest vector elements, with subsequent pairs taken alternately from each source register.
func Vzip1QS16(v0 Int16X8, v1 Int16X8) Int16X8 {
	return C.vzip1q_s16(v0, v1)
}

// Zip vectors (primary). This instruction reads adjacent vector elements from the lower half of two source SIMD&FP registers as pairs, interleaves the pairs and places them into a vector, and writes the vector to the destination SIMD&FP register. The first pair from the first source register is placed into the two lowest vector elements, with subsequent pairs taken alternately from each source register.
func Vzip1U8(v0 Uint8X8, v1 Uint8X8) Uint8X8 {
	return C.vzip1_u8(v0, v1)
}

// Zip vectors (primary). This instruction reads adjacent vector elements from the lower half of two source SIMD&FP registers as pairs, interleaves the pairs and places them into a vector, and writes the vector to the destination SIMD&FP register. The first pair from the first source register is placed into the two lowest vector elements, with subsequent pairs taken alternately from each source register.
func Vzip1U32(v0 Uint32X2, v1 Uint32X2) Uint32X2 {
	return C.vzip1_u32(v0, v1)
}

// Zip vectors (primary). This instruction reads adjacent vector elements from the lower half of two source SIMD&FP registers as pairs, interleaves the pairs and places them into a vector, and writes the vector to the destination SIMD&FP register. The first pair from the first source register is placed into the two lowest vector elements, with subsequent pairs taken alternately from each source register.
func Vzip1U16(v0 Uint16X4, v1 Uint16X4) Uint16X4 {
	return C.vzip1_u16(v0, v1)
}

// Zip vectors (primary). This instruction reads adjacent vector elements from the lower half of two source SIMD&FP registers as pairs, interleaves the pairs and places them into a vector, and writes the vector to the destination SIMD&FP register. The first pair from the first source register is placed into the two lowest vector elements, with subsequent pairs taken alternately from each source register.
func Vzip1S8(v0 Int8X8, v1 Int8X8) Int8X8 {
	return C.vzip1_s8(v0, v1)
}

// Zip vectors (primary). This instruction reads adjacent vector elements from the lower half of two source SIMD&FP registers as pairs, interleaves the pairs and places them into a vector, and writes the vector to the destination SIMD&FP register. The first pair from the first source register is placed into the two lowest vector elements, with subsequent pairs taken alternately from each source register.
func Vzip1F32(v0 Float32X2, v1 Float32X2) Float32X2 {
	return C.vzip1_f32(v0, v1)
}

// Zip vectors (primary). This instruction reads adjacent vector elements from the lower half of two source SIMD&FP registers as pairs, interleaves the pairs and places them into a vector, and writes the vector to the destination SIMD&FP register. The first pair from the first source register is placed into the two lowest vector elements, with subsequent pairs taken alternately from each source register.
func Vzip1S32(v0 Int32X2, v1 Int32X2) Int32X2 {
	return C.vzip1_s32(v0, v1)
}

// Zip vectors (primary). This instruction reads adjacent vector elements from the lower half of two source SIMD&FP registers as pairs, interleaves the pairs and places them into a vector, and writes the vector to the destination SIMD&FP register. The first pair from the first source register is placed into the two lowest vector elements, with subsequent pairs taken alternately from each source register.
func Vzip1S16(v0 Int16X4, v1 Int16X4) Int16X4 {
	return C.vzip1_s16(v0, v1)
}

// Zip vectors (secondary). This instruction reads adjacent vector elements from the upper half of two source SIMD&FP registers as pairs, interleaves the pairs and places them into a vector, and writes the vector to the destination SIMD&FP register. The first pair from the first source register is placed into the two lowest vector elements, with subsequent pairs taken alternately from each source register.
func Vzip2P8(v0 Poly8X8, v1 Poly8X8) Poly8X8 {
	return C.vzip2_p8(v0, v1)
}

// Zip vectors (secondary). This instruction reads adjacent vector elements from the upper half of two source SIMD&FP registers as pairs, interleaves the pairs and places them into a vector, and writes the vector to the destination SIMD&FP register. The first pair from the first source register is placed into the two lowest vector elements, with subsequent pairs taken alternately from each source register.
func Vzip2P16(v0 Poly16X4, v1 Poly16X4) Poly16X4 {
	return C.vzip2_p16(v0, v1)
}

// Zip vectors (secondary). This instruction reads adjacent vector elements from the upper half of two source SIMD&FP registers as pairs, interleaves the pairs and places them into a vector, and writes the vector to the destination SIMD&FP register. The first pair from the first source register is placed into the two lowest vector elements, with subsequent pairs taken alternately from each source register.
func Vzip2QP8(v0 Poly8X16, v1 Poly8X16) Poly8X16 {
	return C.vzip2q_p8(v0, v1)
}

// Zip vectors (secondary). This instruction reads adjacent vector elements from the upper half of two source SIMD&FP registers as pairs, interleaves the pairs and places them into a vector, and writes the vector to the destination SIMD&FP register. The first pair from the first source register is placed into the two lowest vector elements, with subsequent pairs taken alternately from each source register.
func Vzip2QP64(v0 Poly64X2, v1 Poly64X2) Poly64X2 {
	return C.vzip2q_p64(v0, v1)
}

// Zip vectors (secondary). This instruction reads adjacent vector elements from the upper half of two source SIMD&FP registers as pairs, interleaves the pairs and places them into a vector, and writes the vector to the destination SIMD&FP register. The first pair from the first source register is placed into the two lowest vector elements, with subsequent pairs taken alternately from each source register.
func Vzip2QP16(v0 Poly16X8, v1 Poly16X8) Poly16X8 {
	return C.vzip2q_p16(v0, v1)
}

// Zip vectors (secondary). This instruction reads adjacent vector elements from the upper half of two source SIMD&FP registers as pairs, interleaves the pairs and places them into a vector, and writes the vector to the destination SIMD&FP register. The first pair from the first source register is placed into the two lowest vector elements, with subsequent pairs taken alternately from each source register.
func Vzip2QU8(v0 Uint8X16, v1 Uint8X16) Uint8X16 {
	return C.vzip2q_u8(v0, v1)
}

// Zip vectors (secondary). This instruction reads adjacent vector elements from the upper half of two source SIMD&FP registers as pairs, interleaves the pairs and places them into a vector, and writes the vector to the destination SIMD&FP register. The first pair from the first source register is placed into the two lowest vector elements, with subsequent pairs taken alternately from each source register.
func Vzip2QU32(v0 Uint32X4, v1 Uint32X4) Uint32X4 {
	return C.vzip2q_u32(v0, v1)
}

// Zip vectors (secondary). This instruction reads adjacent vector elements from the upper half of two source SIMD&FP registers as pairs, interleaves the pairs and places them into a vector, and writes the vector to the destination SIMD&FP register. The first pair from the first source register is placed into the two lowest vector elements, with subsequent pairs taken alternately from each source register.
func Vzip2QU64(v0 Uint64X2, v1 Uint64X2) Uint64X2 {
	return C.vzip2q_u64(v0, v1)
}

// Zip vectors (secondary). This instruction reads adjacent vector elements from the upper half of two source SIMD&FP registers as pairs, interleaves the pairs and places them into a vector, and writes the vector to the destination SIMD&FP register. The first pair from the first source register is placed into the two lowest vector elements, with subsequent pairs taken alternately from each source register.
func Vzip2QU16(v0 Uint16X8, v1 Uint16X8) Uint16X8 {
	return C.vzip2q_u16(v0, v1)
}

// Zip vectors (secondary). This instruction reads adjacent vector elements from the upper half of two source SIMD&FP registers as pairs, interleaves the pairs and places them into a vector, and writes the vector to the destination SIMD&FP register. The first pair from the first source register is placed into the two lowest vector elements, with subsequent pairs taken alternately from each source register.
func Vzip2QS8(v0 Int8X16, v1 Int8X16) Int8X16 {
	return C.vzip2q_s8(v0, v1)
}

// Zip vectors (secondary). This instruction reads adjacent vector elements from the upper half of two source SIMD&FP registers as pairs, interleaves the pairs and places them into a vector, and writes the vector to the destination SIMD&FP register. The first pair from the first source register is placed into the two lowest vector elements, with subsequent pairs taken alternately from each source register.
func Vzip2QF64(v0 Float64X2, v1 Float64X2) Float64X2 {
	return C.vzip2q_f64(v0, v1)
}

// Zip vectors (secondary). This instruction reads adjacent vector elements from the upper half of two source SIMD&FP registers as pairs, interleaves the pairs and places them into a vector, and writes the vector to the destination SIMD&FP register. The first pair from the first source register is placed into the two lowest vector elements, with subsequent pairs taken alternately from each source register.
func Vzip2QF32(v0 Float32X4, v1 Float32X4) Float32X4 {
	return C.vzip2q_f32(v0, v1)
}

// Zip vectors (secondary). This instruction reads adjacent vector elements from the upper half of two source SIMD&FP registers as pairs, interleaves the pairs and places them into a vector, and writes the vector to the destination SIMD&FP register. The first pair from the first source register is placed into the two lowest vector elements, with subsequent pairs taken alternately from each source register.
func Vzip2QS32(v0 Int32X4, v1 Int32X4) Int32X4 {
	return C.vzip2q_s32(v0, v1)
}

// Zip vectors (secondary). This instruction reads adjacent vector elements from the upper half of two source SIMD&FP registers as pairs, interleaves the pairs and places them into a vector, and writes the vector to the destination SIMD&FP register. The first pair from the first source register is placed into the two lowest vector elements, with subsequent pairs taken alternately from each source register.
func Vzip2QS64(v0 Int64X2, v1 Int64X2) Int64X2 {
	return C.vzip2q_s64(v0, v1)
}

// Zip vectors (secondary). This instruction reads adjacent vector elements from the upper half of two source SIMD&FP registers as pairs, interleaves the pairs and places them into a vector, and writes the vector to the destination SIMD&FP register. The first pair from the first source register is placed into the two lowest vector elements, with subsequent pairs taken alternately from each source register.
func Vzip2QS16(v0 Int16X8, v1 Int16X8) Int16X8 {
	return C.vzip2q_s16(v0, v1)
}

// Zip vectors (secondary). This instruction reads adjacent vector elements from the upper half of two source SIMD&FP registers as pairs, interleaves the pairs and places them into a vector, and writes the vector to the destination SIMD&FP register. The first pair from the first source register is placed into the two lowest vector elements, with subsequent pairs taken alternately from each source register.
func Vzip2U8(v0 Uint8X8, v1 Uint8X8) Uint8X8 {
	return C.vzip2_u8(v0, v1)
}

// Zip vectors (secondary). This instruction reads adjacent vector elements from the upper half of two source SIMD&FP registers as pairs, interleaves the pairs and places them into a vector, and writes the vector to the destination SIMD&FP register. The first pair from the first source register is placed into the two lowest vector elements, with subsequent pairs taken alternately from each source register.
func Vzip2U32(v0 Uint32X2, v1 Uint32X2) Uint32X2 {
	return C.vzip2_u32(v0, v1)
}

// Zip vectors (secondary). This instruction reads adjacent vector elements from the upper half of two source SIMD&FP registers as pairs, interleaves the pairs and places them into a vector, and writes the vector to the destination SIMD&FP register. The first pair from the first source register is placed into the two lowest vector elements, with subsequent pairs taken alternately from each source register.
func Vzip2U16(v0 Uint16X4, v1 Uint16X4) Uint16X4 {
	return C.vzip2_u16(v0, v1)
}

// Zip vectors (secondary). This instruction reads adjacent vector elements from the upper half of two source SIMD&FP registers as pairs, interleaves the pairs and places them into a vector, and writes the vector to the destination SIMD&FP register. The first pair from the first source register is placed into the two lowest vector elements, with subsequent pairs taken alternately from each source register.
func Vzip2S8(v0 Int8X8, v1 Int8X8) Int8X8 {
	return C.vzip2_s8(v0, v1)
}

// Zip vectors (secondary). This instruction reads adjacent vector elements from the upper half of two source SIMD&FP registers as pairs, interleaves the pairs and places them into a vector, and writes the vector to the destination SIMD&FP register. The first pair from the first source register is placed into the two lowest vector elements, with subsequent pairs taken alternately from each source register.
func Vzip2F32(v0 Float32X2, v1 Float32X2) Float32X2 {
	return C.vzip2_f32(v0, v1)
}

// Zip vectors (secondary). This instruction reads adjacent vector elements from the upper half of two source SIMD&FP registers as pairs, interleaves the pairs and places them into a vector, and writes the vector to the destination SIMD&FP register. The first pair from the first source register is placed into the two lowest vector elements, with subsequent pairs taken alternately from each source register.
func Vzip2S32(v0 Int32X2, v1 Int32X2) Int32X2 {
	return C.vzip2_s32(v0, v1)
}

// Zip vectors (secondary). This instruction reads adjacent vector elements from the upper half of two source SIMD&FP registers as pairs, interleaves the pairs and places them into a vector, and writes the vector to the destination SIMD&FP register. The first pair from the first source register is placed into the two lowest vector elements, with subsequent pairs taken alternately from each source register.
func Vzip2S16(v0 Int16X4, v1 Int16X4) Int16X4 {
	return C.vzip2_s16(v0, v1)
}

// Unsigned Absolute difference and Accumulate. This instruction subtracts the elements of the vector of the second source SIMD&FP register from the corresponding elements of the first source SIMD&FP register, and accumulates the absolute values of the results into the elements of the vector of the destination SIMD&FP register.
func VabaqU8(v0 Uint8X16, v1 Uint8X16, v2 Uint8X16) Uint8X16 {
	return C.vabaq_u8(v0, v1, v2)
}

// Unsigned Absolute difference and Accumulate. This instruction subtracts the elements of the vector of the second source SIMD&FP register from the corresponding elements of the first source SIMD&FP register, and accumulates the absolute values of the results into the elements of the vector of the destination SIMD&FP register.
func VabaqU32(v0 Uint32X4, v1 Uint32X4, v2 Uint32X4) Uint32X4 {
	return C.vabaq_u32(v0, v1, v2)
}

// Unsigned Absolute difference and Accumulate. This instruction subtracts the elements of the vector of the second source SIMD&FP register from the corresponding elements of the first source SIMD&FP register, and accumulates the absolute values of the results into the elements of the vector of the destination SIMD&FP register.
func VabaqU16(v0 Uint16X8, v1 Uint16X8, v2 Uint16X8) Uint16X8 {
	return C.vabaq_u16(v0, v1, v2)
}

// Signed Absolute difference and Accumulate. This instruction subtracts the elements of the vector of the second source SIMD&FP register from the corresponding elements of the first source SIMD&FP register, and accumulates the absolute values of the results into the elements of the vector of the destination SIMD&FP register.
func VabaqS8(v0 Int8X16, v1 Int8X16, v2 Int8X16) Int8X16 {
	return C.vabaq_s8(v0, v1, v2)
}

// Signed Absolute difference and Accumulate. This instruction subtracts the elements of the vector of the second source SIMD&FP register from the corresponding elements of the first source SIMD&FP register, and accumulates the absolute values of the results into the elements of the vector of the destination SIMD&FP register.
func VabaqS32(v0 Int32X4, v1 Int32X4, v2 Int32X4) Int32X4 {
	return C.vabaq_s32(v0, v1, v2)
}

// Signed Absolute difference and Accumulate. This instruction subtracts the elements of the vector of the second source SIMD&FP register from the corresponding elements of the first source SIMD&FP register, and accumulates the absolute values of the results into the elements of the vector of the destination SIMD&FP register.
func VabaqS16(v0 Int16X8, v1 Int16X8, v2 Int16X8) Int16X8 {
	return C.vabaq_s16(v0, v1, v2)
}

// Unsigned Absolute difference and Accumulate. This instruction subtracts the elements of the vector of the second source SIMD&FP register from the corresponding elements of the first source SIMD&FP register, and accumulates the absolute values of the results into the elements of the vector of the destination SIMD&FP register.
func VabaU8(v0 Uint8X8, v1 Uint8X8, v2 Uint8X8) Uint8X8 {
	return C.vaba_u8(v0, v1, v2)
}

// Unsigned Absolute difference and Accumulate. This instruction subtracts the elements of the vector of the second source SIMD&FP register from the corresponding elements of the first source SIMD&FP register, and accumulates the absolute values of the results into the elements of the vector of the destination SIMD&FP register.
func VabaU32(v0 Uint32X2, v1 Uint32X2, v2 Uint32X2) Uint32X2 {
	return C.vaba_u32(v0, v1, v2)
}

// Unsigned Absolute difference and Accumulate. This instruction subtracts the elements of the vector of the second source SIMD&FP register from the corresponding elements of the first source SIMD&FP register, and accumulates the absolute values of the results into the elements of the vector of the destination SIMD&FP register.
func VabaU16(v0 Uint16X4, v1 Uint16X4, v2 Uint16X4) Uint16X4 {
	return C.vaba_u16(v0, v1, v2)
}

// Signed Absolute difference and Accumulate. This instruction subtracts the elements of the vector of the second source SIMD&FP register from the corresponding elements of the first source SIMD&FP register, and accumulates the absolute values of the results into the elements of the vector of the destination SIMD&FP register.
func VabaS8(v0 Int8X8, v1 Int8X8, v2 Int8X8) Int8X8 {
	return C.vaba_s8(v0, v1, v2)
}

// Signed Absolute difference and Accumulate. This instruction subtracts the elements of the vector of the second source SIMD&FP register from the corresponding elements of the first source SIMD&FP register, and accumulates the absolute values of the results into the elements of the vector of the destination SIMD&FP register.
func VabaS32(v0 Int32X2, v1 Int32X2, v2 Int32X2) Int32X2 {
	return C.vaba_s32(v0, v1, v2)
}

// Signed Absolute difference and Accumulate. This instruction subtracts the elements of the vector of the second source SIMD&FP register from the corresponding elements of the first source SIMD&FP register, and accumulates the absolute values of the results into the elements of the vector of the destination SIMD&FP register.
func VabaS16(v0 Int16X4, v1 Int16X4, v2 Int16X4) Int16X4 {
	return C.vaba_s16(v0, v1, v2)
}

// Unsigned Absolute Difference Long. This instruction subtracts the vector elements in the lower or upper half of the second source SIMD&FP register from the corresponding vector elements of the first source SIMD&FP register, places the absolute value of the result into a vector, and writes the vector to the destination SIMD&FP register. The destination vector elements are twice as long as the source vector elements. All the values in this instruction are unsigned integer values.
func VabdlU8(v0 Uint8X8, v1 Uint8X8) Uint16X8 {
	return C.vabdl_u8(v0, v1)
}

// Unsigned Absolute Difference Long. This instruction subtracts the vector elements in the lower or upper half of the second source SIMD&FP register from the corresponding vector elements of the first source SIMD&FP register, places the absolute value of the result into a vector, and writes the vector to the destination SIMD&FP register. The destination vector elements are twice as long as the source vector elements. All the values in this instruction are unsigned integer values.
func VabdlU32(v0 Uint32X2, v1 Uint32X2) Uint64X2 {
	return C.vabdl_u32(v0, v1)
}

// Unsigned Absolute Difference Long. This instruction subtracts the vector elements in the lower or upper half of the second source SIMD&FP register from the corresponding vector elements of the first source SIMD&FP register, places the absolute value of the result into a vector, and writes the vector to the destination SIMD&FP register. The destination vector elements are twice as long as the source vector elements. All the values in this instruction are unsigned integer values.
func VabdlU16(v0 Uint16X4, v1 Uint16X4) Uint32X4 {
	return C.vabdl_u16(v0, v1)
}

// Signed Absolute Difference Long. This instruction subtracts the vector elements of the second source SIMD&FP register from the corresponding vector elements of the first source SIMD&FP register, places the absolute value of the results into a vector, and writes the vector to the lower or upper half of the destination SIMD&FP register. The destination vector elements are twice as long as the source vector elements.
func VabdlS8(v0 Int8X8, v1 Int8X8) Int16X8 {
	return C.vabdl_s8(v0, v1)
}

// Signed Absolute Difference Long. This instruction subtracts the vector elements of the second source SIMD&FP register from the corresponding vector elements of the first source SIMD&FP register, places the absolute value of the results into a vector, and writes the vector to the lower or upper half of the destination SIMD&FP register. The destination vector elements are twice as long as the source vector elements.
func VabdlS32(v0 Int32X2, v1 Int32X2) Int64X2 {
	return C.vabdl_s32(v0, v1)
}

// Signed Absolute Difference Long. This instruction subtracts the vector elements of the second source SIMD&FP register from the corresponding vector elements of the first source SIMD&FP register, places the absolute value of the results into a vector, and writes the vector to the lower or upper half of the destination SIMD&FP register. The destination vector elements are twice as long as the source vector elements.
func VabdlS16(v0 Int16X4, v1 Int16X4) Int32X4 {
	return C.vabdl_s16(v0, v1)
}

// Unsigned Add Long (vector). This instruction adds each vector element in the lower or upper half of the first source SIMD&FP register to the corresponding vector element of the second source SIMD&FP register, places the result into a vector, and writes the vector to the destination SIMD&FP register. The destination vector elements are twice as long as the source vector elements. All the values in this instruction are unsigned integer values.
func VaddlU8(v0 Uint8X8, v1 Uint8X8) Uint16X8 {
	return C.vaddl_u8(v0, v1)
}

// Unsigned Add Long (vector). This instruction adds each vector element in the lower or upper half of the first source SIMD&FP register to the corresponding vector element of the second source SIMD&FP register, places the result into a vector, and writes the vector to the destination SIMD&FP register. The destination vector elements are twice as long as the source vector elements. All the values in this instruction are unsigned integer values.
func VaddlU32(v0 Uint32X2, v1 Uint32X2) Uint64X2 {
	return C.vaddl_u32(v0, v1)
}

// Unsigned Add Long (vector). This instruction adds each vector element in the lower or upper half of the first source SIMD&FP register to the corresponding vector element of the second source SIMD&FP register, places the result into a vector, and writes the vector to the destination SIMD&FP register. The destination vector elements are twice as long as the source vector elements. All the values in this instruction are unsigned integer values.
func VaddlU16(v0 Uint16X4, v1 Uint16X4) Uint32X4 {
	return C.vaddl_u16(v0, v1)
}

// Signed Add Long (vector). This instruction adds each vector element in the lower or upper half of the first source SIMD&FP register to the corresponding vector element of the second source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register. The destination vector elements are twice as long as the source vector elements. All the values in this instruction are signed integer values.
func VaddlS8(v0 Int8X8, v1 Int8X8) Int16X8 {
	return C.vaddl_s8(v0, v1)
}

// Signed Add Long (vector). This instruction adds each vector element in the lower or upper half of the first source SIMD&FP register to the corresponding vector element of the second source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register. The destination vector elements are twice as long as the source vector elements. All the values in this instruction are signed integer values.
func VaddlS32(v0 Int32X2, v1 Int32X2) Int64X2 {
	return C.vaddl_s32(v0, v1)
}

// Signed Add Long (vector). This instruction adds each vector element in the lower or upper half of the first source SIMD&FP register to the corresponding vector element of the second source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register. The destination vector elements are twice as long as the source vector elements. All the values in this instruction are signed integer values.
func VaddlS16(v0 Int16X4, v1 Int16X4) Int32X4 {
	return C.vaddl_s16(v0, v1)
}

// Unsigned Add Wide. This instruction adds the vector elements of the first source SIMD&FP register to the corresponding vector elements in the lower or upper half of the second source SIMD&FP register, places the result in a vector, and writes the vector to the SIMD&FP destination register. The vector elements of the destination register and the first source register are twice as long as the vector elements of the second source register. All the values in this instruction are unsigned integer values.
func VaddwU8(v0 Uint16X8, v1 Uint8X8) Uint16X8 {
	return C.vaddw_u8(v0, v1)
}

// Unsigned Add Wide. This instruction adds the vector elements of the first source SIMD&FP register to the corresponding vector elements in the lower or upper half of the second source SIMD&FP register, places the result in a vector, and writes the vector to the SIMD&FP destination register. The vector elements of the destination register and the first source register are twice as long as the vector elements of the second source register. All the values in this instruction are unsigned integer values.
func VaddwU32(v0 Uint64X2, v1 Uint32X2) Uint64X2 {
	return C.vaddw_u32(v0, v1)
}

// Unsigned Add Wide. This instruction adds the vector elements of the first source SIMD&FP register to the corresponding vector elements in the lower or upper half of the second source SIMD&FP register, places the result in a vector, and writes the vector to the SIMD&FP destination register. The vector elements of the destination register and the first source register are twice as long as the vector elements of the second source register. All the values in this instruction are unsigned integer values.
func VaddwU16(v0 Uint32X4, v1 Uint16X4) Uint32X4 {
	return C.vaddw_u16(v0, v1)
}

// Signed Add Wide. This instruction adds vector elements of the first source SIMD&FP register to the corresponding vector elements in the lower or upper half of the second source SIMD&FP register, places the results in a vector, and writes the vector to the SIMD&FP destination register.
func VaddwS8(v0 Int16X8, v1 Int8X8) Int16X8 {
	return C.vaddw_s8(v0, v1)
}

// Signed Add Wide. This instruction adds vector elements of the first source SIMD&FP register to the corresponding vector elements in the lower or upper half of the second source SIMD&FP register, places the results in a vector, and writes the vector to the SIMD&FP destination register.
func VaddwS32(v0 Int64X2, v1 Int32X2) Int64X2 {
	return C.vaddw_s32(v0, v1)
}

// Signed Add Wide. This instruction adds vector elements of the first source SIMD&FP register to the corresponding vector elements in the lower or upper half of the second source SIMD&FP register, places the results in a vector, and writes the vector to the SIMD&FP destination register.
func VaddwS16(v0 Int32X4, v1 Int16X4) Int32X4 {
	return C.vaddw_s16(v0, v1)
}

// Unsigned Multiply-Add Long (vector). This instruction multiplies the vector elements in the lower or upper half of the first source SIMD&FP register by the corresponding vector elements of the second source SIMD&FP register, and accumulates the results with the vector elements of the destination SIMD&FP register. The destination vector elements are twice as long as the elements that are multiplied.
func VmlalU8(v0 Uint16X8, v1 Uint8X8, v2 Uint8X8) Uint16X8 {
	return C.vmlal_u8(v0, v1, v2)
}

// Unsigned Multiply-Add Long (vector). This instruction multiplies the vector elements in the lower or upper half of the first source SIMD&FP register by the corresponding vector elements of the second source SIMD&FP register, and accumulates the results with the vector elements of the destination SIMD&FP register. The destination vector elements are twice as long as the elements that are multiplied.
func VmlalU32(v0 Uint64X2, v1 Uint32X2, v2 Uint32X2) Uint64X2 {
	return C.vmlal_u32(v0, v1, v2)
}

// Unsigned Multiply-Add Long (vector). This instruction multiplies the vector elements in the lower or upper half of the first source SIMD&FP register by the corresponding vector elements of the second source SIMD&FP register, and accumulates the results with the vector elements of the destination SIMD&FP register. The destination vector elements are twice as long as the elements that are multiplied.
func VmlalU16(v0 Uint32X4, v1 Uint16X4, v2 Uint16X4) Uint32X4 {
	return C.vmlal_u16(v0, v1, v2)
}

// Signed Multiply-Add Long (vector). This instruction multiplies corresponding signed integer values in the lower or upper half of the vectors of the two source SIMD&FP registers, and accumulates the results with the vector elements of the destination SIMD&FP register. The destination vector elements are twice as long as the elements that are multiplied.
func VmlalS8(v0 Int16X8, v1 Int8X8, v2 Int8X8) Int16X8 {
	return C.vmlal_s8(v0, v1, v2)
}

// Signed Multiply-Add Long (vector). This instruction multiplies corresponding signed integer values in the lower or upper half of the vectors of the two source SIMD&FP registers, and accumulates the results with the vector elements of the destination SIMD&FP register. The destination vector elements are twice as long as the elements that are multiplied.
func VmlalS32(v0 Int64X2, v1 Int32X2, v2 Int32X2) Int64X2 {
	return C.vmlal_s32(v0, v1, v2)
}

// Signed Multiply-Add Long (vector). This instruction multiplies corresponding signed integer values in the lower or upper half of the vectors of the two source SIMD&FP registers, and accumulates the results with the vector elements of the destination SIMD&FP register. The destination vector elements are twice as long as the elements that are multiplied.
func VmlalS16(v0 Int32X4, v1 Int16X4, v2 Int16X4) Int32X4 {
	return C.vmlal_s16(v0, v1, v2)
}

// Vector widening multiply accumulate with scalar
func VmlalNU32(v0 Uint64X2, v1 Uint32X2, v2 Uint32) Uint64X2 {
	return C.vmlal_n_u32(v0, v1, v2)
}

// Vector widening multiply accumulate with scalar
func VmlalNU16(v0 Uint32X4, v1 Uint16X4, v2 Uint16) Uint32X4 {
	return C.vmlal_n_u16(v0, v1, v2)
}

// Vector widening multiply accumulate with scalar
func VmlalNS32(v0 Int64X2, v1 Int32X2, v2 Int32) Int64X2 {
	return C.vmlal_n_s32(v0, v1, v2)
}

// Vector widening multiply accumulate with scalar
func VmlalNS16(v0 Int32X4, v1 Int16X4, v2 Int16) Int32X4 {
	return C.vmlal_n_s16(v0, v1, v2)
}

// Unsigned Multiply-Subtract Long (vector). This instruction multiplies corresponding vector elements in the lower or upper half of the two source SIMD&FP registers, and subtracts the results from the vector elements of the destination SIMD&FP register. The destination vector elements are twice as long as the elements that are multiplied. All the values in this instruction are unsigned integer values.
func VmlslU8(v0 Uint16X8, v1 Uint8X8, v2 Uint8X8) Uint16X8 {
	return C.vmlsl_u8(v0, v1, v2)
}

// Unsigned Multiply-Subtract Long (vector). This instruction multiplies corresponding vector elements in the lower or upper half of the two source SIMD&FP registers, and subtracts the results from the vector elements of the destination SIMD&FP register. The destination vector elements are twice as long as the elements that are multiplied. All the values in this instruction are unsigned integer values.
func VmlslU32(v0 Uint64X2, v1 Uint32X2, v2 Uint32X2) Uint64X2 {
	return C.vmlsl_u32(v0, v1, v2)
}

// Unsigned Multiply-Subtract Long (vector). This instruction multiplies corresponding vector elements in the lower or upper half of the two source SIMD&FP registers, and subtracts the results from the vector elements of the destination SIMD&FP register. The destination vector elements are twice as long as the elements that are multiplied. All the values in this instruction are unsigned integer values.
func VmlslU16(v0 Uint32X4, v1 Uint16X4, v2 Uint16X4) Uint32X4 {
	return C.vmlsl_u16(v0, v1, v2)
}

// Signed Multiply-Subtract Long (vector). This instruction multiplies corresponding signed integer values in the lower or upper half of the vectors of the two source SIMD&FP registers, and subtracts the results from the vector elements of the destination SIMD&FP register. The destination vector elements are twice as long as the elements that are multiplied.
func VmlslS8(v0 Int16X8, v1 Int8X8, v2 Int8X8) Int16X8 {
	return C.vmlsl_s8(v0, v1, v2)
}

// Signed Multiply-Subtract Long (vector). This instruction multiplies corresponding signed integer values in the lower or upper half of the vectors of the two source SIMD&FP registers, and subtracts the results from the vector elements of the destination SIMD&FP register. The destination vector elements are twice as long as the elements that are multiplied.
func VmlslS32(v0 Int64X2, v1 Int32X2, v2 Int32X2) Int64X2 {
	return C.vmlsl_s32(v0, v1, v2)
}

// Signed Multiply-Subtract Long (vector). This instruction multiplies corresponding signed integer values in the lower or upper half of the vectors of the two source SIMD&FP registers, and subtracts the results from the vector elements of the destination SIMD&FP register. The destination vector elements are twice as long as the elements that are multiplied.
func VmlslS16(v0 Int32X4, v1 Int16X4, v2 Int16X4) Int32X4 {
	return C.vmlsl_s16(v0, v1, v2)
}

// Vector widening multiply subtract with scalar
func VmlslNU32(v0 Uint64X2, v1 Uint32X2, v2 Uint32) Uint64X2 {
	return C.vmlsl_n_u32(v0, v1, v2)
}

// Vector widening multiply subtract with scalar
func VmlslNU16(v0 Uint32X4, v1 Uint16X4, v2 Uint16) Uint32X4 {
	return C.vmlsl_n_u16(v0, v1, v2)
}

// Vector widening multiply subtract with scalar
func VmlslNS32(v0 Int64X2, v1 Int32X2, v2 Int32) Int64X2 {
	return C.vmlsl_n_s32(v0, v1, v2)
}

// Vector widening multiply subtract with scalar
func VmlslNS16(v0 Int32X4, v1 Int16X4, v2 Int16) Int32X4 {
	return C.vmlsl_n_s16(v0, v1, v2)
}

// Signed Saturating Rounding Doubling Multiply Subtract returning High Half (vector). This instruction multiplies the vector elements of the first source SIMD&FP register with the corresponding vector elements of the second source SIMD&FP register without saturating the multiply results, doubles the results, and subtracts the most significant half of the final results from the vector elements of the destination SIMD&FP register. The results are rounded.
func VqrdmlahsS32(v0 Int32, v1 Int32, v2 Int32) Int32 {
	return C.vqrdmlahs_s32(v0, v1, v2)
}

// Signed Saturating Rounding Doubling Multiply Subtract returning High Half (vector). This instruction multiplies the vector elements of the first source SIMD&FP register with the corresponding vector elements of the second source SIMD&FP register without saturating the multiply results, doubles the results, and subtracts the most significant half of the final results from the vector elements of the destination SIMD&FP register. The results are rounded.
func VqrdmlahhS16(v0 Int16, v1 Int16, v2 Int16) Int16 {
	return C.vqrdmlahh_s16(v0, v1, v2)
}

// Signed Saturating Rounding Doubling Multiply Subtract returning High Half (vector). This instruction multiplies the vector elements of the first source SIMD&FP register with the corresponding vector elements of the second source SIMD&FP register without saturating the multiply results, doubles the results, and subtracts the most significant half of the final results from the vector elements of the destination SIMD&FP register. The results are rounded.
func VqrdmlshsS32(v0 Int32, v1 Int32, v2 Int32) Int32 {
	return C.vqrdmlshs_s32(v0, v1, v2)
}

// Signed Saturating Rounding Doubling Multiply Subtract returning High Half (vector). This instruction multiplies the vector elements of the first source SIMD&FP register with the corresponding vector elements of the second source SIMD&FP register without saturating the multiply results, doubles the results, and subtracts the most significant half of the final results from the vector elements of the destination SIMD&FP register. The results are rounded.
func VqrdmlshhS16(v0 Int16, v1 Int16, v2 Int16) Int16 {
	return C.vqrdmlshh_s16(v0, v1, v2)
}

// Unsigned Absolute Difference Long. This instruction subtracts the vector elements in the lower or upper half of the second source SIMD&FP register from the corresponding vector elements of the first source SIMD&FP register, places the absolute value of the result into a vector, and writes the vector to the destination SIMD&FP register. The destination vector elements are twice as long as the source vector elements. All the values in this instruction are unsigned integer values.
func VabdlHighU8(v0 Uint8X16, v1 Uint8X16) Uint16X8 {
	return C.vabdl_high_u8(v0, v1)
}

// Unsigned Absolute Difference Long. This instruction subtracts the vector elements in the lower or upper half of the second source SIMD&FP register from the corresponding vector elements of the first source SIMD&FP register, places the absolute value of the result into a vector, and writes the vector to the destination SIMD&FP register. The destination vector elements are twice as long as the source vector elements. All the values in this instruction are unsigned integer values.
func VabdlHighU32(v0 Uint32X4, v1 Uint32X4) Uint64X2 {
	return C.vabdl_high_u32(v0, v1)
}

// Unsigned Absolute Difference Long. This instruction subtracts the vector elements in the lower or upper half of the second source SIMD&FP register from the corresponding vector elements of the first source SIMD&FP register, places the absolute value of the result into a vector, and writes the vector to the destination SIMD&FP register. The destination vector elements are twice as long as the source vector elements. All the values in this instruction are unsigned integer values.
func VabdlHighU16(v0 Uint16X8, v1 Uint16X8) Uint32X4 {
	return C.vabdl_high_u16(v0, v1)
}

// Signed Absolute Difference Long. This instruction subtracts the vector elements of the second source SIMD&FP register from the corresponding vector elements of the first source SIMD&FP register, places the absolute value of the results into a vector, and writes the vector to the lower or upper half of the destination SIMD&FP register. The destination vector elements are twice as long as the source vector elements.
func VabdlHighS8(v0 Int8X16, v1 Int8X16) Int16X8 {
	return C.vabdl_high_s8(v0, v1)
}

// Signed Absolute Difference Long. This instruction subtracts the vector elements of the second source SIMD&FP register from the corresponding vector elements of the first source SIMD&FP register, places the absolute value of the results into a vector, and writes the vector to the lower or upper half of the destination SIMD&FP register. The destination vector elements are twice as long as the source vector elements.
func VabdlHighS32(v0 Int32X4, v1 Int32X4) Int64X2 {
	return C.vabdl_high_s32(v0, v1)
}

// Signed Absolute Difference Long. This instruction subtracts the vector elements of the second source SIMD&FP register from the corresponding vector elements of the first source SIMD&FP register, places the absolute value of the results into a vector, and writes the vector to the lower or upper half of the destination SIMD&FP register. The destination vector elements are twice as long as the source vector elements.
func VabdlHighS16(v0 Int16X8, v1 Int16X8) Int32X4 {
	return C.vabdl_high_s16(v0, v1)
}

// Unsigned Add Long (vector). This instruction adds each vector element in the lower or upper half of the first source SIMD&FP register to the corresponding vector element of the second source SIMD&FP register, places the result into a vector, and writes the vector to the destination SIMD&FP register. The destination vector elements are twice as long as the source vector elements. All the values in this instruction are unsigned integer values.
func VaddlHighU8(v0 Uint8X16, v1 Uint8X16) Uint16X8 {
	return C.vaddl_high_u8(v0, v1)
}

// Unsigned Add Long (vector). This instruction adds each vector element in the lower or upper half of the first source SIMD&FP register to the corresponding vector element of the second source SIMD&FP register, places the result into a vector, and writes the vector to the destination SIMD&FP register. The destination vector elements are twice as long as the source vector elements. All the values in this instruction are unsigned integer values.
func VaddlHighU32(v0 Uint32X4, v1 Uint32X4) Uint64X2 {
	return C.vaddl_high_u32(v0, v1)
}

// Unsigned Add Long (vector). This instruction adds each vector element in the lower or upper half of the first source SIMD&FP register to the corresponding vector element of the second source SIMD&FP register, places the result into a vector, and writes the vector to the destination SIMD&FP register. The destination vector elements are twice as long as the source vector elements. All the values in this instruction are unsigned integer values.
func VaddlHighU16(v0 Uint16X8, v1 Uint16X8) Uint32X4 {
	return C.vaddl_high_u16(v0, v1)
}

// Signed Add Long (vector). This instruction adds each vector element in the lower or upper half of the first source SIMD&FP register to the corresponding vector element of the second source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register. The destination vector elements are twice as long as the source vector elements. All the values in this instruction are signed integer values.
func VaddlHighS8(v0 Int8X16, v1 Int8X16) Int16X8 {
	return C.vaddl_high_s8(v0, v1)
}

// Signed Add Long (vector). This instruction adds each vector element in the lower or upper half of the first source SIMD&FP register to the corresponding vector element of the second source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register. The destination vector elements are twice as long as the source vector elements. All the values in this instruction are signed integer values.
func VaddlHighS32(v0 Int32X4, v1 Int32X4) Int64X2 {
	return C.vaddl_high_s32(v0, v1)
}

// Signed Add Long (vector). This instruction adds each vector element in the lower or upper half of the first source SIMD&FP register to the corresponding vector element of the second source SIMD&FP register, places the results into a vector, and writes the vector to the destination SIMD&FP register. The destination vector elements are twice as long as the source vector elements. All the values in this instruction are signed integer values.
func VaddlHighS16(v0 Int16X8, v1 Int16X8) Int32X4 {
	return C.vaddl_high_s16(v0, v1)
}

// Unsigned Add Wide. This instruction adds the vector elements of the first source SIMD&FP register to the corresponding vector elements in the lower or upper half of the second source SIMD&FP register, places the result in a vector, and writes the vector to the SIMD&FP destination register. The vector elements of the destination register and the first source register are twice as long as the vector elements of the second source register. All the values in this instruction are unsigned integer values.
func VaddwHighU8(v0 Uint16X8, v1 Uint8X16) Uint16X8 {
	return C.vaddw_high_u8(v0, v1)
}

// Unsigned Add Wide. This instruction adds the vector elements of the first source SIMD&FP register to the corresponding vector elements in the lower or upper half of the second source SIMD&FP register, places the result in a vector, and writes the vector to the SIMD&FP destination register. The vector elements of the destination register and the first source register are twice as long as the vector elements of the second source register. All the values in this instruction are unsigned integer values.
func VaddwHighU32(v0 Uint64X2, v1 Uint32X4) Uint64X2 {
	return C.vaddw_high_u32(v0, v1)
}

// Unsigned Add Wide. This instruction adds the vector elements of the first source SIMD&FP register to the corresponding vector elements in the lower or upper half of the second source SIMD&FP register, places the result in a vector, and writes the vector to the SIMD&FP destination register. The vector elements of the destination register and the first source register are twice as long as the vector elements of the second source register. All the values in this instruction are unsigned integer values.
func VaddwHighU16(v0 Uint32X4, v1 Uint16X8) Uint32X4 {
	return C.vaddw_high_u16(v0, v1)
}

// Signed Add Wide. This instruction adds vector elements of the first source SIMD&FP register to the corresponding vector elements in the lower or upper half of the second source SIMD&FP register, places the results in a vector, and writes the vector to the SIMD&FP destination register.
func VaddwHighS8(v0 Int16X8, v1 Int8X16) Int16X8 {
	return C.vaddw_high_s8(v0, v1)
}

// Signed Add Wide. This instruction adds vector elements of the first source SIMD&FP register to the corresponding vector elements in the lower or upper half of the second source SIMD&FP register, places the results in a vector, and writes the vector to the SIMD&FP destination register.
func VaddwHighS32(v0 Int64X2, v1 Int32X4) Int64X2 {
	return C.vaddw_high_s32(v0, v1)
}

// Signed Add Wide. This instruction adds vector elements of the first source SIMD&FP register to the corresponding vector elements in the lower or upper half of the second source SIMD&FP register, places the results in a vector, and writes the vector to the SIMD&FP destination register.
func VaddwHighS16(v0 Int32X4, v1 Int16X8) Int32X4 {
	return C.vaddw_high_s16(v0, v1)
}

// Unsigned Multiply-Add Long (vector). This instruction multiplies the vector elements in the lower or upper half of the first source SIMD&FP register by the corresponding vector elements of the second source SIMD&FP register, and accumulates the results with the vector elements of the destination SIMD&FP register. The destination vector elements are twice as long as the elements that are multiplied.
func VmlalHighU8(v0 Uint16X8, v1 Uint8X16, v2 Uint8X16) Uint16X8 {
	return C.vmlal_high_u8(v0, v1, v2)
}

// Unsigned Multiply-Add Long (vector). This instruction multiplies the vector elements in the lower or upper half of the first source SIMD&FP register by the corresponding vector elements of the second source SIMD&FP register, and accumulates the results with the vector elements of the destination SIMD&FP register. The destination vector elements are twice as long as the elements that are multiplied.
func VmlalHighU32(v0 Uint64X2, v1 Uint32X4, v2 Uint32X4) Uint64X2 {
	return C.vmlal_high_u32(v0, v1, v2)
}

// Unsigned Multiply-Add Long (vector). This instruction multiplies the vector elements in the lower or upper half of the first source SIMD&FP register by the corresponding vector elements of the second source SIMD&FP register, and accumulates the results with the vector elements of the destination SIMD&FP register. The destination vector elements are twice as long as the elements that are multiplied.
func VmlalHighU16(v0 Uint32X4, v1 Uint16X8, v2 Uint16X8) Uint32X4 {
	return C.vmlal_high_u16(v0, v1, v2)
}

// Signed Multiply-Add Long (vector). This instruction multiplies corresponding signed integer values in the lower or upper half of the vectors of the two source SIMD&FP registers, and accumulates the results with the vector elements of the destination SIMD&FP register. The destination vector elements are twice as long as the elements that are multiplied.
func VmlalHighS8(v0 Int16X8, v1 Int8X16, v2 Int8X16) Int16X8 {
	return C.vmlal_high_s8(v0, v1, v2)
}

// Signed Multiply-Add Long (vector). This instruction multiplies corresponding signed integer values in the lower or upper half of the vectors of the two source SIMD&FP registers, and accumulates the results with the vector elements of the destination SIMD&FP register. The destination vector elements are twice as long as the elements that are multiplied.
func VmlalHighS32(v0 Int64X2, v1 Int32X4, v2 Int32X4) Int64X2 {
	return C.vmlal_high_s32(v0, v1, v2)
}

// Signed Multiply-Add Long (vector). This instruction multiplies corresponding signed integer values in the lower or upper half of the vectors of the two source SIMD&FP registers, and accumulates the results with the vector elements of the destination SIMD&FP register. The destination vector elements are twice as long as the elements that are multiplied.
func VmlalHighS16(v0 Int32X4, v1 Int16X8, v2 Int16X8) Int32X4 {
	return C.vmlal_high_s16(v0, v1, v2)
}

// Unsigned Multiply-Add Long (vector). This instruction multiplies the vector elements in the lower or upper half of the first source SIMD&FP register by the corresponding vector elements of the second source SIMD&FP register, and accumulates the results with the vector elements of the destination SIMD&FP register. The destination vector elements are twice as long as the elements that are multiplied.
func VmlalHighNU32(v0 Uint64X2, v1 Uint32X4, v2 Uint32) Uint64X2 {
	return C.vmlal_high_n_u32(v0, v1, v2)
}

// Unsigned Multiply-Add Long (vector). This instruction multiplies the vector elements in the lower or upper half of the first source SIMD&FP register by the corresponding vector elements of the second source SIMD&FP register, and accumulates the results with the vector elements of the destination SIMD&FP register. The destination vector elements are twice as long as the elements that are multiplied.
func VmlalHighNU16(v0 Uint32X4, v1 Uint16X8, v2 Uint16) Uint32X4 {
	return C.vmlal_high_n_u16(v0, v1, v2)
}

// Signed Multiply-Add Long (vector). This instruction multiplies corresponding signed integer values in the lower or upper half of the vectors of the two source SIMD&FP registers, and accumulates the results with the vector elements of the destination SIMD&FP register. The destination vector elements are twice as long as the elements that are multiplied.
func VmlalHighNS32(v0 Int64X2, v1 Int32X4, v2 Int32) Int64X2 {
	return C.vmlal_high_n_s32(v0, v1, v2)
}

// Signed Multiply-Add Long (vector). This instruction multiplies corresponding signed integer values in the lower or upper half of the vectors of the two source SIMD&FP registers, and accumulates the results with the vector elements of the destination SIMD&FP register. The destination vector elements are twice as long as the elements that are multiplied.
func VmlalHighNS16(v0 Int32X4, v1 Int16X8, v2 Int16) Int32X4 {
	return C.vmlal_high_n_s16(v0, v1, v2)
}

// Unsigned Multiply-Subtract Long (vector). This instruction multiplies corresponding vector elements in the lower or upper half of the two source SIMD&FP registers, and subtracts the results from the vector elements of the destination SIMD&FP register. The destination vector elements are twice as long as the elements that are multiplied. All the values in this instruction are unsigned integer values.
func VmlslHighU8(v0 Uint16X8, v1 Uint8X16, v2 Uint8X16) Uint16X8 {
	return C.vmlsl_high_u8(v0, v1, v2)
}

// Unsigned Multiply-Subtract Long (vector). This instruction multiplies corresponding vector elements in the lower or upper half of the two source SIMD&FP registers, and subtracts the results from the vector elements of the destination SIMD&FP register. The destination vector elements are twice as long as the elements that are multiplied. All the values in this instruction are unsigned integer values.
func VmlslHighU32(v0 Uint64X2, v1 Uint32X4, v2 Uint32X4) Uint64X2 {
	return C.vmlsl_high_u32(v0, v1, v2)
}

// Unsigned Multiply-Subtract Long (vector). This instruction multiplies corresponding vector elements in the lower or upper half of the two source SIMD&FP registers, and subtracts the results from the vector elements of the destination SIMD&FP register. The destination vector elements are twice as long as the elements that are multiplied. All the values in this instruction are unsigned integer values.
func VmlslHighU16(v0 Uint32X4, v1 Uint16X8, v2 Uint16X8) Uint32X4 {
	return C.vmlsl_high_u16(v0, v1, v2)
}

// Signed Multiply-Subtract Long (vector). This instruction multiplies corresponding signed integer values in the lower or upper half of the vectors of the two source SIMD&FP registers, and subtracts the results from the vector elements of the destination SIMD&FP register. The destination vector elements are twice as long as the elements that are multiplied.
func VmlslHighS8(v0 Int16X8, v1 Int8X16, v2 Int8X16) Int16X8 {
	return C.vmlsl_high_s8(v0, v1, v2)
}

// Signed Multiply-Subtract Long (vector). This instruction multiplies corresponding signed integer values in the lower or upper half of the vectors of the two source SIMD&FP registers, and subtracts the results from the vector elements of the destination SIMD&FP register. The destination vector elements are twice as long as the elements that are multiplied.
func VmlslHighS32(v0 Int64X2, v1 Int32X4, v2 Int32X4) Int64X2 {
	return C.vmlsl_high_s32(v0, v1, v2)
}

// Signed Multiply-Subtract Long (vector). This instruction multiplies corresponding signed integer values in the lower or upper half of the vectors of the two source SIMD&FP registers, and subtracts the results from the vector elements of the destination SIMD&FP register. The destination vector elements are twice as long as the elements that are multiplied.
func VmlslHighS16(v0 Int32X4, v1 Int16X8, v2 Int16X8) Int32X4 {
	return C.vmlsl_high_s16(v0, v1, v2)
}

// Unsigned Multiply-Subtract Long (vector). This instruction multiplies corresponding vector elements in the lower or upper half of the two source SIMD&FP registers, and subtracts the results from the vector elements of the destination SIMD&FP register. The destination vector elements are twice as long as the elements that are multiplied. All the values in this instruction are unsigned integer values.
func VmlslHighNU32(v0 Uint64X2, v1 Uint32X4, v2 Uint32) Uint64X2 {
	return C.vmlsl_high_n_u32(v0, v1, v2)
}

// Unsigned Multiply-Subtract Long (vector). This instruction multiplies corresponding vector elements in the lower or upper half of the two source SIMD&FP registers, and subtracts the results from the vector elements of the destination SIMD&FP register. The destination vector elements are twice as long as the elements that are multiplied. All the values in this instruction are unsigned integer values.
func VmlslHighNU16(v0 Uint32X4, v1 Uint16X8, v2 Uint16) Uint32X4 {
	return C.vmlsl_high_n_u16(v0, v1, v2)
}

// Signed Multiply-Subtract Long (vector). This instruction multiplies corresponding signed integer values in the lower or upper half of the vectors of the two source SIMD&FP registers, and subtracts the results from the vector elements of the destination SIMD&FP register. The destination vector elements are twice as long as the elements that are multiplied.
func VmlslHighNS32(v0 Int64X2, v1 Int32X4, v2 Int32) Int64X2 {
	return C.vmlsl_high_n_s32(v0, v1, v2)
}

// Signed Multiply-Subtract Long (vector). This instruction multiplies corresponding signed integer values in the lower or upper half of the vectors of the two source SIMD&FP registers, and subtracts the results from the vector elements of the destination SIMD&FP register. The destination vector elements are twice as long as the elements that are multiplied.
func VmlslHighNS16(v0 Int32X4, v1 Int16X8, v2 Int16) Int32X4 {
	return C.vmlsl_high_n_s16(v0, v1, v2)
}

// Unsigned Absolute difference and Accumulate Long. This instruction subtracts the vector elements in the lower or upper half of the second source SIMD&FP register from the corresponding vector elements of the first source SIMD&FP register, and accumulates the absolute values of the results into the vector elements of the destination SIMD&FP register. The destination vector elements are twice as long as the source vector elements. All the values in this instruction are unsigned integer values.
func VabalU8(v0 Uint16X8, v1 Uint8X8, v2 Uint8X8) Uint16X8 {
	return C.vabal_u8(v0, v1, v2)
}

// Unsigned Absolute difference and Accumulate Long. This instruction subtracts the vector elements in the lower or upper half of the second source SIMD&FP register from the corresponding vector elements of the first source SIMD&FP register, and accumulates the absolute values of the results into the vector elements of the destination SIMD&FP register. The destination vector elements are twice as long as the source vector elements. All the values in this instruction are unsigned integer values.
func VabalU32(v0 Uint64X2, v1 Uint32X2, v2 Uint32X2) Uint64X2 {
	return C.vabal_u32(v0, v1, v2)
}

// Unsigned Absolute difference and Accumulate Long. This instruction subtracts the vector elements in the lower or upper half of the second source SIMD&FP register from the corresponding vector elements of the first source SIMD&FP register, and accumulates the absolute values of the results into the vector elements of the destination SIMD&FP register. The destination vector elements are twice as long as the source vector elements. All the values in this instruction are unsigned integer values.
func VabalU16(v0 Uint32X4, v1 Uint16X4, v2 Uint16X4) Uint32X4 {
	return C.vabal_u16(v0, v1, v2)
}

// Signed Absolute difference and Accumulate Long. This instruction subtracts the vector elements in the lower or upper half of the second source SIMD&FP register from the corresponding vector elements of the first source SIMD&FP register, and accumulates the absolute values of the results into the vector elements of the destination SIMD&FP register. The destination vector elements are twice as long as the source vector elements.
func VabalS8(v0 Int16X8, v1 Int8X8, v2 Int8X8) Int16X8 {
	return C.vabal_s8(v0, v1, v2)
}

// Signed Absolute difference and Accumulate Long. This instruction subtracts the vector elements in the lower or upper half of the second source SIMD&FP register from the corresponding vector elements of the first source SIMD&FP register, and accumulates the absolute values of the results into the vector elements of the destination SIMD&FP register. The destination vector elements are twice as long as the source vector elements.
func VabalS32(v0 Int64X2, v1 Int32X2, v2 Int32X2) Int64X2 {
	return C.vabal_s32(v0, v1, v2)
}

// Signed Absolute difference and Accumulate Long. This instruction subtracts the vector elements in the lower or upper half of the second source SIMD&FP register from the corresponding vector elements of the first source SIMD&FP register, and accumulates the absolute values of the results into the vector elements of the destination SIMD&FP register. The destination vector elements are twice as long as the source vector elements.
func VabalS16(v0 Int32X4, v1 Int16X4, v2 Int16X4) Int32X4 {
	return C.vabal_s16(v0, v1, v2)
}

// Unsigned Absolute difference and Accumulate Long. This instruction subtracts the vector elements in the lower or upper half of the second source SIMD&FP register from the corresponding vector elements of the first source SIMD&FP register, and accumulates the absolute values of the results into the vector elements of the destination SIMD&FP register. The destination vector elements are twice as long as the source vector elements. All the values in this instruction are unsigned integer values.
func VabalHighU8(v0 Uint16X8, v1 Uint8X16, v2 Uint8X16) Uint16X8 {
	return C.vabal_high_u8(v0, v1, v2)
}

// Unsigned Absolute difference and Accumulate Long. This instruction subtracts the vector elements in the lower or upper half of the second source SIMD&FP register from the corresponding vector elements of the first source SIMD&FP register, and accumulates the absolute values of the results into the vector elements of the destination SIMD&FP register. The destination vector elements are twice as long as the source vector elements. All the values in this instruction are unsigned integer values.
func VabalHighU32(v0 Uint64X2, v1 Uint32X4, v2 Uint32X4) Uint64X2 {
	return C.vabal_high_u32(v0, v1, v2)
}

// Unsigned Absolute difference and Accumulate Long. This instruction subtracts the vector elements in the lower or upper half of the second source SIMD&FP register from the corresponding vector elements of the first source SIMD&FP register, and accumulates the absolute values of the results into the vector elements of the destination SIMD&FP register. The destination vector elements are twice as long as the source vector elements. All the values in this instruction are unsigned integer values.
func VabalHighU16(v0 Uint32X4, v1 Uint16X8, v2 Uint16X8) Uint32X4 {
	return C.vabal_high_u16(v0, v1, v2)
}

// Signed Absolute difference and Accumulate Long. This instruction subtracts the vector elements in the lower or upper half of the second source SIMD&FP register from the corresponding vector elements of the first source SIMD&FP register, and accumulates the absolute values of the results into the vector elements of the destination SIMD&FP register. The destination vector elements are twice as long as the source vector elements.
func VabalHighS8(v0 Int16X8, v1 Int8X16, v2 Int8X16) Int16X8 {
	return C.vabal_high_s8(v0, v1, v2)
}

// Signed Absolute difference and Accumulate Long. This instruction subtracts the vector elements in the lower or upper half of the second source SIMD&FP register from the corresponding vector elements of the first source SIMD&FP register, and accumulates the absolute values of the results into the vector elements of the destination SIMD&FP register. The destination vector elements are twice as long as the source vector elements.
func VabalHighS32(v0 Int64X2, v1 Int32X4, v2 Int32X4) Int64X2 {
	return C.vabal_high_s32(v0, v1, v2)
}

// Signed Absolute difference and Accumulate Long. This instruction subtracts the vector elements in the lower or upper half of the second source SIMD&FP register from the corresponding vector elements of the first source SIMD&FP register, and accumulates the absolute values of the results into the vector elements of the destination SIMD&FP register. The destination vector elements are twice as long as the source vector elements.
func VabalHighS16(v0 Int32X4, v1 Int16X8, v2 Int16X8) Int32X4 {
	return C.vabal_high_s16(v0, v1, v2)
}
