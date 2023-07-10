package arm

/*
#include <arm_neon.h>
*/
import "C"

// typedef signed char int8_t;
type Int8 = C.int8_t

// typedef short int16_t;
type Int16 = C.int16_t

// typedef int int32_t;
type Int32 = C.int32_t

// typedef longlong int64_t;
type Int64 = C.int64_t

// typedef uchar uint8_t;
type Uint8 = C.uint8_t

// typedef ushort uint16_t;
type Uint16 = C.uint16_t

// typedef uint uint32_t;
type Uint32 = C.uint32_t

// typedef ulonglong uint64_t;
type Uint64 = C.uint64_t

// typedef float float32_t;
type Float32 = C.float32_t

// typedef double float64_t;
type Float64 = C.float64_t

// typedef uint8_t poly8_t;
type Poly8 = C.poly8_t

// typedef uint16_t poly16_t;
type Poly16 = C.poly16_t

// typedef uint64_t poly64_t;
type Poly64 = C.poly64_t

// typedef __uint128_t poly128_t;
type Poly128 = C.poly128_t

// typedef __attribute__((neon_vector_type(8))) int8_t int8x8_t;
type Int8X8 = C.int8x8_t

// typedef __attribute__((neon_vector_type(16))) int8_t int8x16_t;
type Int8X16 = C.int8x16_t

// typedef __attribute__((neon_vector_type(4))) int16_t int16x4_t;
type Int16X4 = C.int16x4_t

// typedef __attribute__((neon_vector_type(8))) int16_t int16x8_t;
type Int16X8 = C.int16x8_t

// typedef __attribute__((neon_vector_type(2))) int32_t int32x2_t;
type Int32X2 = C.int32x2_t

// typedef __attribute__((neon_vector_type(4))) int32_t int32x4_t;
type Int32X4 = C.int32x4_t

// typedef __attribute__((neon_vector_type(1))) int64_t int64x1_t;
type Int64X1 = C.int64x1_t

// typedef __attribute__((neon_vector_type(2))) int64_t int64x2_t;
type Int64X2 = C.int64x2_t

// typedef __attribute__((neon_vector_type(8))) uint8_t uint8x8_t;
type Uint8X8 = C.uint8x8_t

// typedef __attribute__((neon_vector_type(16))) uint8_t uint8x16_t;
type Uint8X16 = C.uint8x16_t

// typedef __attribute__((neon_vector_type(4))) uint16_t uint16x4_t;
type Uint16X4 = C.uint16x4_t

// typedef __attribute__((neon_vector_type(8))) uint16_t uint16x8_t;
type Uint16X8 = C.uint16x8_t

// typedef __attribute__((neon_vector_type(2))) uint32_t uint32x2_t;
type Uint32X2 = C.uint32x2_t

// typedef __attribute__((neon_vector_type(4))) uint32_t uint32x4_t;
type Uint32X4 = C.uint32x4_t

// typedef __attribute__((neon_vector_type(1))) uint64_t uint64x1_t;
type Uint64X1 = C.uint64x1_t

// typedef __attribute__((neon_vector_type(2))) uint64_t uint64x2_t;
type Uint64X2 = C.uint64x2_t

// typedef __attribute__((neon_vector_type(2))) float32_t float32x2_t;
type Float32X2 = C.float32x2_t

// typedef __attribute__((neon_vector_type(4))) float32_t float32x4_t;
type Float32X4 = C.float32x4_t

// typedef __attribute__((neon_vector_type(1))) float64_t float64x1_t;
type Float64X1 = C.float64x1_t

// typedef __attribute__((neon_vector_type(2))) float64_t float64x2_t;
type Float64X2 = C.float64x2_t

// typedef __attribute__((neon_polyvector_type(8))) poly8_t poly8x8_t;
type Poly8X8 = C.poly8x8_t

// typedef __attribute__((neon_polyvector_type(16))) poly8_t poly8x16_t;
type Poly8X16 = C.poly8x16_t

// typedef __attribute__((neon_polyvector_type(4))) poly16_t poly16x4_t;
type Poly16X4 = C.poly16x4_t

// typedef __attribute__((neon_polyvector_type(8))) poly16_t poly16x8_t;
type Poly16X8 = C.poly16x8_t

// typedef __attribute__((neon_polyvector_type(1))) poly64_t poly64x1_t;
type Poly64X1 = C.poly64x1_t

// typedef __attribute__((neon_polyvector_type(2))) poly64_t poly64x2_t;
type Poly64X2 = C.poly64x2_t

// typedef struct int8x8x2_t { int8x8_t val[2];} int8x8x2_t;
type Int8X8X2 = C.int8x8x2_t

// typedef struct int8x16x2_t { int8x16_t val[2];} int8x16x2_t;
type Int8X16X2 = C.int8x16x2_t

// typedef struct int16x4x2_t { int16x4_t val[2];} int16x4x2_t;
type Int16X4X2 = C.int16x4x2_t

// typedef struct int16x8x2_t { int16x8_t val[2];} int16x8x2_t;
type Int16X8X2 = C.int16x8x2_t

// typedef struct int32x2x2_t { int32x2_t val[2];} int32x2x2_t;
type Int32X2X2 = C.int32x2x2_t

// typedef struct int32x4x2_t { int32x4_t val[2];} int32x4x2_t;
type Int32X4X2 = C.int32x4x2_t

// typedef struct uint8x8x2_t { uint8x8_t val[2];} uint8x8x2_t;
type Uint8X8X2 = C.uint8x8x2_t

// typedef struct uint8x16x2_t { uint8x16_t val[2];} uint8x16x2_t;
type Uint8X16X2 = C.uint8x16x2_t

// typedef struct uint16x4x2_t { uint16x4_t val[2];} uint16x4x2_t;
type Uint16X4X2 = C.uint16x4x2_t

// typedef struct uint16x8x2_t { uint16x8_t val[2];} uint16x8x2_t;
type Uint16X8X2 = C.uint16x8x2_t

// typedef struct uint32x2x2_t { uint32x2_t val[2];} uint32x2x2_t;
type Uint32X2X2 = C.uint32x2x2_t

// typedef struct uint32x4x2_t { uint32x4_t val[2];} uint32x4x2_t;
type Uint32X4X2 = C.uint32x4x2_t

// typedef struct float32x2x2_t { float32x2_t val[2];} float32x2x2_t;
type Float32X2X2 = C.float32x2x2_t

// typedef struct float32x4x2_t { float32x4_t val[2];} float32x4x2_t;
type Float32X4X2 = C.float32x4x2_t

// typedef struct poly8x8x2_t { poly8x8_t val[2];} poly8x8x2_t;
type Poly8X8X2 = C.poly8x8x2_t

// typedef struct poly8x16x2_t { poly8x16_t val[2];} poly8x16x2_t;
type Poly8X16X2 = C.poly8x16x2_t

// typedef struct poly16x4x2_t { poly16x4_t val[2];} poly16x4x2_t;
type Poly16X4X2 = C.poly16x4x2_t

// typedef struct poly16x8x2_t { poly16x8_t val[2];} poly16x8x2_t;
type Poly16X8X2 = C.poly16x8x2_t

// typedef struct int8x8x3_t { int8x8_t val[3];} int8x8x3_t;
type Int8X8X3 = C.int8x8x3_t

// typedef struct int8x16x3_t { int8x16_t val[3];} int8x16x3_t;
type Int8X16X3 = C.int8x16x3_t

// typedef struct uint8x8x3_t { uint8x8_t val[3];} uint8x8x3_t;
type Uint8X8X3 = C.uint8x8x3_t

// typedef struct uint8x16x3_t { uint8x16_t val[3];} uint8x16x3_t;
type Uint8X16X3 = C.uint8x16x3_t

// typedef struct poly8x8x3_t { poly8x8_t val[3];} poly8x8x3_t;
type Poly8X8X3 = C.poly8x8x3_t

// typedef struct poly8x16x3_t { poly8x16_t val[3];} poly8x16x3_t;
type Poly8X16X3 = C.poly8x16x3_t

// typedef struct int8x8x4_t { int8x8_t val[4];} int8x8x4_t;
type Int8X8X4 = C.int8x8x4_t

// typedef struct int8x16x4_t { int8x16_t val[4];} int8x16x4_t;
type Int8X16X4 = C.int8x16x4_t

// typedef struct uint8x8x4_t { uint8x8_t val[4];} uint8x8x4_t;
type Uint8X8X4 = C.uint8x8x4_t

// typedef struct uint8x16x4_t { uint8x16_t val[4];} uint8x16x4_t;
type Uint8X16X4 = C.uint8x16x4_t

// typedef struct poly8x8x4_t { poly8x8_t val[4];} poly8x8x4_t;
type Poly8X8X4 = C.poly8x8x4_t

// typedef struct poly8x16x4_t { poly8x16_t val[4];} poly8x16x4_t;
type Poly8X16X4 = C.poly8x16x4_t
