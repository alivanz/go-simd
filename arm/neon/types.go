package neon

/*
#include <arm_neon.h>
*/
import "C"

type Int8X8X2 C.int8x8x2_t

type Int8X16X2 C.int8x16x2_t

type Int16X4X2 C.int16x4x2_t

type Int16X8X2 C.int16x8x2_t

type Int32X2X2 C.int32x2x2_t

type Int32X4X2 C.int32x4x2_t

type Int64X1X2 C.int64x1x2_t

type Int64X2X2 C.int64x2x2_t

type Uint8X8X2 C.uint8x8x2_t

type Uint8X16X2 C.uint8x16x2_t

type Uint16X4X2 C.uint16x4x2_t

type Uint16X8X2 C.uint16x8x2_t

type Uint32X2X2 C.uint32x2x2_t

type Uint32X4X2 C.uint32x4x2_t

type Uint64X1X2 C.uint64x1x2_t

type Uint64X2X2 C.uint64x2x2_t

type Float32X2X2 C.float32x2x2_t

type Float32X4X2 C.float32x4x2_t

type Float64X1X2 C.float64x1x2_t

type Float64X2X2 C.float64x2x2_t

type Poly8X8X2 C.poly8x8x2_t

type Poly8X16X2 C.poly8x16x2_t

type Poly16X4X2 C.poly16x4x2_t

type Poly16X8X2 C.poly16x8x2_t

type Poly64X1X2 C.poly64x1x2_t

type Poly64X2X2 C.poly64x2x2_t

type Int8X8X3 C.int8x8x3_t

type Int8X16X3 C.int8x16x3_t

type Int16X4X3 C.int16x4x3_t

type Int16X8X3 C.int16x8x3_t

type Int32X2X3 C.int32x2x3_t

type Int32X4X3 C.int32x4x3_t

type Int64X1X3 C.int64x1x3_t

type Int64X2X3 C.int64x2x3_t

type Uint8X8X3 C.uint8x8x3_t

type Uint8X16X3 C.uint8x16x3_t

type Uint16X4X3 C.uint16x4x3_t

type Uint16X8X3 C.uint16x8x3_t

type Uint32X2X3 C.uint32x2x3_t

type Uint32X4X3 C.uint32x4x3_t

type Uint64X1X3 C.uint64x1x3_t

type Uint64X2X3 C.uint64x2x3_t

type Float32X2X3 C.float32x2x3_t

type Float32X4X3 C.float32x4x3_t

type Float64X1X3 C.float64x1x3_t

type Float64X2X3 C.float64x2x3_t

type Poly8X8X3 C.poly8x8x3_t

type Poly8X16X3 C.poly8x16x3_t

type Poly16X4X3 C.poly16x4x3_t

type Poly16X8X3 C.poly16x8x3_t

type Poly64X1X3 C.poly64x1x3_t

type Poly64X2X3 C.poly64x2x3_t

type Int8X8X4 C.int8x8x4_t

type Int8X16X4 C.int8x16x4_t

type Int16X4X4 C.int16x4x4_t

type Int16X8X4 C.int16x8x4_t

type Int32X2X4 C.int32x2x4_t

type Int32X4X4 C.int32x4x4_t

type Int64X1X4 C.int64x1x4_t

type Int64X2X4 C.int64x2x4_t

type Uint8X8X4 C.uint8x8x4_t

type Uint8X16X4 C.uint8x16x4_t

type Uint16X4X4 C.uint16x4x4_t

type Uint16X8X4 C.uint16x8x4_t

type Uint32X2X4 C.uint32x2x4_t

type Uint32X4X4 C.uint32x4x4_t

type Uint64X1X4 C.uint64x1x4_t

type Uint64X2X4 C.uint64x2x4_t

type Float32X2X4 C.float32x2x4_t

type Float32X4X4 C.float32x4x4_t

type Float64X1X4 C.float64x1x4_t

type Float64X2X4 C.float64x2x4_t

type Poly8X8X4 C.poly8x8x4_t

type Poly8X16X4 C.poly8x16x4_t

type Poly16X4X4 C.poly16x4x4_t

type Poly16X8X4 C.poly16x8x4_t

type Poly64X1X4 C.poly64x1x4_t

type Poly64X2X4 C.poly64x2x4_t

type Int8X8 C.int8x8_t

type Int8X16 C.int8x16_t

type Int16X4 C.int16x4_t

type Int16X8 C.int16x8_t

type Int32X2 C.int32x2_t

type Int32X4 C.int32x4_t

type Int64X1 C.int64x1_t

type Int64X2 C.int64x2_t

type Uint8X8 C.uint8x8_t

type Uint8X16 C.uint8x16_t

type Uint16X4 C.uint16x4_t

type Uint16X8 C.uint16x8_t

type Uint32X2 C.uint32x2_t

type Uint32X4 C.uint32x4_t

type Uint64X1 C.uint64x1_t

type Uint64X2 C.uint64x2_t

type Float32X2 C.float32x2_t

type Float32X4 C.float32x4_t

type Float64X1 C.float64x1_t

type Float64X2 C.float64x2_t

type Poly8X8 C.poly8x8_t

type Poly8X16 C.poly8x16_t

type Poly16X4 C.poly16x4_t

type Poly16X8 C.poly16x8_t

type Poly64X1 C.poly64x1_t

type Poly64X2 C.poly64x2_t
