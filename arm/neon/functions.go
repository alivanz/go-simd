package neon

/*
#include <arm_neon.h>
*/
import "C"

func VabdqU8(v0 Uint8X16, v1 Uint8X16) Uint8X16 {
	return Uint8X16(C.vabdq_u8(C.uint8x16_t(v0), C.uint8x16_t(v1)))
}

func VabdqU32(v0 Uint32X4, v1 Uint32X4) Uint32X4 {
	return Uint32X4(C.vabdq_u32(C.uint32x4_t(v0), C.uint32x4_t(v1)))
}

func VabdqU16(v0 Uint16X8, v1 Uint16X8) Uint16X8 {
	return Uint16X8(C.vabdq_u16(C.uint16x8_t(v0), C.uint16x8_t(v1)))
}

func VabdqS8(v0 Int8X16, v1 Int8X16) Int8X16 {
	return Int8X16(C.vabdq_s8(C.int8x16_t(v0), C.int8x16_t(v1)))
}

func VabdqF32(v0 Float32X4, v1 Float32X4) Float32X4 {
	return Float32X4(C.vabdq_f32(C.float32x4_t(v0), C.float32x4_t(v1)))
}

func VabdqS32(v0 Int32X4, v1 Int32X4) Int32X4 {
	return Int32X4(C.vabdq_s32(C.int32x4_t(v0), C.int32x4_t(v1)))
}

func VabdqS16(v0 Int16X8, v1 Int16X8) Int16X8 {
	return Int16X8(C.vabdq_s16(C.int16x8_t(v0), C.int16x8_t(v1)))
}

func VabdU8(v0 Uint8X8, v1 Uint8X8) Uint8X8 {
	return Uint8X8(C.vabd_u8(C.uint8x8_t(v0), C.uint8x8_t(v1)))
}

func VabdU32(v0 Uint32X2, v1 Uint32X2) Uint32X2 {
	return Uint32X2(C.vabd_u32(C.uint32x2_t(v0), C.uint32x2_t(v1)))
}

func VabdU16(v0 Uint16X4, v1 Uint16X4) Uint16X4 {
	return Uint16X4(C.vabd_u16(C.uint16x4_t(v0), C.uint16x4_t(v1)))
}

func VabdS8(v0 Int8X8, v1 Int8X8) Int8X8 {
	return Int8X8(C.vabd_s8(C.int8x8_t(v0), C.int8x8_t(v1)))
}

func VabdF32(v0 Float32X2, v1 Float32X2) Float32X2 {
	return Float32X2(C.vabd_f32(C.float32x2_t(v0), C.float32x2_t(v1)))
}

func VabdS32(v0 Int32X2, v1 Int32X2) Int32X2 {
	return Int32X2(C.vabd_s32(C.int32x2_t(v0), C.int32x2_t(v1)))
}

func VabdS16(v0 Int16X4, v1 Int16X4) Int16X4 {
	return Int16X4(C.vabd_s16(C.int16x4_t(v0), C.int16x4_t(v1)))
}

func VabsqS8(v0 Int8X16) Int8X16 {
	return Int8X16(C.vabsq_s8(C.int8x16_t(v0)))
}

func VabsqF32(v0 Float32X4) Float32X4 {
	return Float32X4(C.vabsq_f32(C.float32x4_t(v0)))
}

func VabsqS32(v0 Int32X4) Int32X4 {
	return Int32X4(C.vabsq_s32(C.int32x4_t(v0)))
}

func VabsqS16(v0 Int16X8) Int16X8 {
	return Int16X8(C.vabsq_s16(C.int16x8_t(v0)))
}

func VabsS8(v0 Int8X8) Int8X8 {
	return Int8X8(C.vabs_s8(C.int8x8_t(v0)))
}

func VabsF32(v0 Float32X2) Float32X2 {
	return Float32X2(C.vabs_f32(C.float32x2_t(v0)))
}

func VabsS32(v0 Int32X2) Int32X2 {
	return Int32X2(C.vabs_s32(C.int32x2_t(v0)))
}

func VabsS16(v0 Int16X4) Int16X4 {
	return Int16X4(C.vabs_s16(C.int16x4_t(v0)))
}

func VaddqU8(v0 Uint8X16, v1 Uint8X16) Uint8X16 {
	return Uint8X16(C.vaddq_u8(C.uint8x16_t(v0), C.uint8x16_t(v1)))
}

func VaddqU32(v0 Uint32X4, v1 Uint32X4) Uint32X4 {
	return Uint32X4(C.vaddq_u32(C.uint32x4_t(v0), C.uint32x4_t(v1)))
}

func VaddqU64(v0 Uint64X2, v1 Uint64X2) Uint64X2 {
	return Uint64X2(C.vaddq_u64(C.uint64x2_t(v0), C.uint64x2_t(v1)))
}

func VaddqU16(v0 Uint16X8, v1 Uint16X8) Uint16X8 {
	return Uint16X8(C.vaddq_u16(C.uint16x8_t(v0), C.uint16x8_t(v1)))
}

func VaddqS8(v0 Int8X16, v1 Int8X16) Int8X16 {
	return Int8X16(C.vaddq_s8(C.int8x16_t(v0), C.int8x16_t(v1)))
}

func VaddqF32(v0 Float32X4, v1 Float32X4) Float32X4 {
	return Float32X4(C.vaddq_f32(C.float32x4_t(v0), C.float32x4_t(v1)))
}

func VaddqS32(v0 Int32X4, v1 Int32X4) Int32X4 {
	return Int32X4(C.vaddq_s32(C.int32x4_t(v0), C.int32x4_t(v1)))
}

func VaddqS64(v0 Int64X2, v1 Int64X2) Int64X2 {
	return Int64X2(C.vaddq_s64(C.int64x2_t(v0), C.int64x2_t(v1)))
}

func VaddqS16(v0 Int16X8, v1 Int16X8) Int16X8 {
	return Int16X8(C.vaddq_s16(C.int16x8_t(v0), C.int16x8_t(v1)))
}

func VaddU8(v0 Uint8X8, v1 Uint8X8) Uint8X8 {
	return Uint8X8(C.vadd_u8(C.uint8x8_t(v0), C.uint8x8_t(v1)))
}

func VaddU32(v0 Uint32X2, v1 Uint32X2) Uint32X2 {
	return Uint32X2(C.vadd_u32(C.uint32x2_t(v0), C.uint32x2_t(v1)))
}

func VaddU64(v0 Uint64X1, v1 Uint64X1) Uint64X1 {
	return Uint64X1(C.vadd_u64(C.uint64x1_t(v0), C.uint64x1_t(v1)))
}

func VaddU16(v0 Uint16X4, v1 Uint16X4) Uint16X4 {
	return Uint16X4(C.vadd_u16(C.uint16x4_t(v0), C.uint16x4_t(v1)))
}

func VaddS8(v0 Int8X8, v1 Int8X8) Int8X8 {
	return Int8X8(C.vadd_s8(C.int8x8_t(v0), C.int8x8_t(v1)))
}

func VaddF32(v0 Float32X2, v1 Float32X2) Float32X2 {
	return Float32X2(C.vadd_f32(C.float32x2_t(v0), C.float32x2_t(v1)))
}

func VaddS32(v0 Int32X2, v1 Int32X2) Int32X2 {
	return Int32X2(C.vadd_s32(C.int32x2_t(v0), C.int32x2_t(v1)))
}

func VaddS64(v0 Int64X1, v1 Int64X1) Int64X1 {
	return Int64X1(C.vadd_s64(C.int64x1_t(v0), C.int64x1_t(v1)))
}

func VaddS16(v0 Int16X4, v1 Int16X4) Int16X4 {
	return Int16X4(C.vadd_s16(C.int16x4_t(v0), C.int16x4_t(v1)))
}

func VaddP8(v0 Poly8X8, v1 Poly8X8) Poly8X8 {
	return Poly8X8(C.vadd_p8(C.poly8x8_t(v0), C.poly8x8_t(v1)))
}

func VaddP64(v0 Poly64X1, v1 Poly64X1) Poly64X1 {
	return Poly64X1(C.vadd_p64(C.poly64x1_t(v0), C.poly64x1_t(v1)))
}

func VaddP16(v0 Poly16X4, v1 Poly16X4) Poly16X4 {
	return Poly16X4(C.vadd_p16(C.poly16x4_t(v0), C.poly16x4_t(v1)))
}

func VaddqP8(v0 Poly8X16, v1 Poly8X16) Poly8X16 {
	return Poly8X16(C.vaddq_p8(C.poly8x16_t(v0), C.poly8x16_t(v1)))
}

func VaddqP64(v0 Poly64X2, v1 Poly64X2) Poly64X2 {
	return Poly64X2(C.vaddq_p64(C.poly64x2_t(v0), C.poly64x2_t(v1)))
}

func VaddqP16(v0 Poly16X8, v1 Poly16X8) Poly16X8 {
	return Poly16X8(C.vaddq_p16(C.poly16x8_t(v0), C.poly16x8_t(v1)))
}

func VaddhnU32(v0 Uint32X4, v1 Uint32X4) Uint16X4 {
	return Uint16X4(C.vaddhn_u32(C.uint32x4_t(v0), C.uint32x4_t(v1)))
}

func VaddhnU64(v0 Uint64X2, v1 Uint64X2) Uint32X2 {
	return Uint32X2(C.vaddhn_u64(C.uint64x2_t(v0), C.uint64x2_t(v1)))
}

func VaddhnU16(v0 Uint16X8, v1 Uint16X8) Uint8X8 {
	return Uint8X8(C.vaddhn_u16(C.uint16x8_t(v0), C.uint16x8_t(v1)))
}

func VaddhnS32(v0 Int32X4, v1 Int32X4) Int16X4 {
	return Int16X4(C.vaddhn_s32(C.int32x4_t(v0), C.int32x4_t(v1)))
}

func VaddhnS64(v0 Int64X2, v1 Int64X2) Int32X2 {
	return Int32X2(C.vaddhn_s64(C.int64x2_t(v0), C.int64x2_t(v1)))
}

func VaddhnS16(v0 Int16X8, v1 Int16X8) Int8X8 {
	return Int8X8(C.vaddhn_s16(C.int16x8_t(v0), C.int16x8_t(v1)))
}

func VandqU8(v0 Uint8X16, v1 Uint8X16) Uint8X16 {
	return Uint8X16(C.vandq_u8(C.uint8x16_t(v0), C.uint8x16_t(v1)))
}

func VandqU32(v0 Uint32X4, v1 Uint32X4) Uint32X4 {
	return Uint32X4(C.vandq_u32(C.uint32x4_t(v0), C.uint32x4_t(v1)))
}

func VandqU64(v0 Uint64X2, v1 Uint64X2) Uint64X2 {
	return Uint64X2(C.vandq_u64(C.uint64x2_t(v0), C.uint64x2_t(v1)))
}

func VandqU16(v0 Uint16X8, v1 Uint16X8) Uint16X8 {
	return Uint16X8(C.vandq_u16(C.uint16x8_t(v0), C.uint16x8_t(v1)))
}

func VandqS8(v0 Int8X16, v1 Int8X16) Int8X16 {
	return Int8X16(C.vandq_s8(C.int8x16_t(v0), C.int8x16_t(v1)))
}

func VandqS32(v0 Int32X4, v1 Int32X4) Int32X4 {
	return Int32X4(C.vandq_s32(C.int32x4_t(v0), C.int32x4_t(v1)))
}

func VandqS64(v0 Int64X2, v1 Int64X2) Int64X2 {
	return Int64X2(C.vandq_s64(C.int64x2_t(v0), C.int64x2_t(v1)))
}

func VandqS16(v0 Int16X8, v1 Int16X8) Int16X8 {
	return Int16X8(C.vandq_s16(C.int16x8_t(v0), C.int16x8_t(v1)))
}

func VandU8(v0 Uint8X8, v1 Uint8X8) Uint8X8 {
	return Uint8X8(C.vand_u8(C.uint8x8_t(v0), C.uint8x8_t(v1)))
}

func VandU32(v0 Uint32X2, v1 Uint32X2) Uint32X2 {
	return Uint32X2(C.vand_u32(C.uint32x2_t(v0), C.uint32x2_t(v1)))
}

func VandU64(v0 Uint64X1, v1 Uint64X1) Uint64X1 {
	return Uint64X1(C.vand_u64(C.uint64x1_t(v0), C.uint64x1_t(v1)))
}

func VandU16(v0 Uint16X4, v1 Uint16X4) Uint16X4 {
	return Uint16X4(C.vand_u16(C.uint16x4_t(v0), C.uint16x4_t(v1)))
}

func VandS8(v0 Int8X8, v1 Int8X8) Int8X8 {
	return Int8X8(C.vand_s8(C.int8x8_t(v0), C.int8x8_t(v1)))
}

func VandS32(v0 Int32X2, v1 Int32X2) Int32X2 {
	return Int32X2(C.vand_s32(C.int32x2_t(v0), C.int32x2_t(v1)))
}

func VandS64(v0 Int64X1, v1 Int64X1) Int64X1 {
	return Int64X1(C.vand_s64(C.int64x1_t(v0), C.int64x1_t(v1)))
}

func VandS16(v0 Int16X4, v1 Int16X4) Int16X4 {
	return Int16X4(C.vand_s16(C.int16x4_t(v0), C.int16x4_t(v1)))
}

func VbicqU8(v0 Uint8X16, v1 Uint8X16) Uint8X16 {
	return Uint8X16(C.vbicq_u8(C.uint8x16_t(v0), C.uint8x16_t(v1)))
}

func VbicqU32(v0 Uint32X4, v1 Uint32X4) Uint32X4 {
	return Uint32X4(C.vbicq_u32(C.uint32x4_t(v0), C.uint32x4_t(v1)))
}

func VbicqU64(v0 Uint64X2, v1 Uint64X2) Uint64X2 {
	return Uint64X2(C.vbicq_u64(C.uint64x2_t(v0), C.uint64x2_t(v1)))
}

func VbicqU16(v0 Uint16X8, v1 Uint16X8) Uint16X8 {
	return Uint16X8(C.vbicq_u16(C.uint16x8_t(v0), C.uint16x8_t(v1)))
}

func VbicqS8(v0 Int8X16, v1 Int8X16) Int8X16 {
	return Int8X16(C.vbicq_s8(C.int8x16_t(v0), C.int8x16_t(v1)))
}

func VbicqS32(v0 Int32X4, v1 Int32X4) Int32X4 {
	return Int32X4(C.vbicq_s32(C.int32x4_t(v0), C.int32x4_t(v1)))
}

func VbicqS64(v0 Int64X2, v1 Int64X2) Int64X2 {
	return Int64X2(C.vbicq_s64(C.int64x2_t(v0), C.int64x2_t(v1)))
}

func VbicqS16(v0 Int16X8, v1 Int16X8) Int16X8 {
	return Int16X8(C.vbicq_s16(C.int16x8_t(v0), C.int16x8_t(v1)))
}

func VbicU8(v0 Uint8X8, v1 Uint8X8) Uint8X8 {
	return Uint8X8(C.vbic_u8(C.uint8x8_t(v0), C.uint8x8_t(v1)))
}

func VbicU32(v0 Uint32X2, v1 Uint32X2) Uint32X2 {
	return Uint32X2(C.vbic_u32(C.uint32x2_t(v0), C.uint32x2_t(v1)))
}

func VbicU64(v0 Uint64X1, v1 Uint64X1) Uint64X1 {
	return Uint64X1(C.vbic_u64(C.uint64x1_t(v0), C.uint64x1_t(v1)))
}

func VbicU16(v0 Uint16X4, v1 Uint16X4) Uint16X4 {
	return Uint16X4(C.vbic_u16(C.uint16x4_t(v0), C.uint16x4_t(v1)))
}

func VbicS8(v0 Int8X8, v1 Int8X8) Int8X8 {
	return Int8X8(C.vbic_s8(C.int8x8_t(v0), C.int8x8_t(v1)))
}

func VbicS32(v0 Int32X2, v1 Int32X2) Int32X2 {
	return Int32X2(C.vbic_s32(C.int32x2_t(v0), C.int32x2_t(v1)))
}

func VbicS64(v0 Int64X1, v1 Int64X1) Int64X1 {
	return Int64X1(C.vbic_s64(C.int64x1_t(v0), C.int64x1_t(v1)))
}

func VbicS16(v0 Int16X4, v1 Int16X4) Int16X4 {
	return Int16X4(C.vbic_s16(C.int16x4_t(v0), C.int16x4_t(v1)))
}

func VbslP8(v0 Uint8X8, v1 Poly8X8, v2 Poly8X8) Poly8X8 {
	return Poly8X8(C.vbsl_p8(C.uint8x8_t(v0), C.poly8x8_t(v1), C.poly8x8_t(v2)))
}

func VbslP16(v0 Uint16X4, v1 Poly16X4, v2 Poly16X4) Poly16X4 {
	return Poly16X4(C.vbsl_p16(C.uint16x4_t(v0), C.poly16x4_t(v1), C.poly16x4_t(v2)))
}

func VbslqP8(v0 Uint8X16, v1 Poly8X16, v2 Poly8X16) Poly8X16 {
	return Poly8X16(C.vbslq_p8(C.uint8x16_t(v0), C.poly8x16_t(v1), C.poly8x16_t(v2)))
}

func VbslqP16(v0 Uint16X8, v1 Poly16X8, v2 Poly16X8) Poly16X8 {
	return Poly16X8(C.vbslq_p16(C.uint16x8_t(v0), C.poly16x8_t(v1), C.poly16x8_t(v2)))
}

func VbslqU8(v0 Uint8X16, v1 Uint8X16, v2 Uint8X16) Uint8X16 {
	return Uint8X16(C.vbslq_u8(C.uint8x16_t(v0), C.uint8x16_t(v1), C.uint8x16_t(v2)))
}

func VbslqU32(v0 Uint32X4, v1 Uint32X4, v2 Uint32X4) Uint32X4 {
	return Uint32X4(C.vbslq_u32(C.uint32x4_t(v0), C.uint32x4_t(v1), C.uint32x4_t(v2)))
}

func VbslqU64(v0 Uint64X2, v1 Uint64X2, v2 Uint64X2) Uint64X2 {
	return Uint64X2(C.vbslq_u64(C.uint64x2_t(v0), C.uint64x2_t(v1), C.uint64x2_t(v2)))
}

func VbslqU16(v0 Uint16X8, v1 Uint16X8, v2 Uint16X8) Uint16X8 {
	return Uint16X8(C.vbslq_u16(C.uint16x8_t(v0), C.uint16x8_t(v1), C.uint16x8_t(v2)))
}

func VbslqS8(v0 Uint8X16, v1 Int8X16, v2 Int8X16) Int8X16 {
	return Int8X16(C.vbslq_s8(C.uint8x16_t(v0), C.int8x16_t(v1), C.int8x16_t(v2)))
}

func VbslqF32(v0 Uint32X4, v1 Float32X4, v2 Float32X4) Float32X4 {
	return Float32X4(C.vbslq_f32(C.uint32x4_t(v0), C.float32x4_t(v1), C.float32x4_t(v2)))
}

func VbslqS32(v0 Uint32X4, v1 Int32X4, v2 Int32X4) Int32X4 {
	return Int32X4(C.vbslq_s32(C.uint32x4_t(v0), C.int32x4_t(v1), C.int32x4_t(v2)))
}

func VbslqS64(v0 Uint64X2, v1 Int64X2, v2 Int64X2) Int64X2 {
	return Int64X2(C.vbslq_s64(C.uint64x2_t(v0), C.int64x2_t(v1), C.int64x2_t(v2)))
}

func VbslqS16(v0 Uint16X8, v1 Int16X8, v2 Int16X8) Int16X8 {
	return Int16X8(C.vbslq_s16(C.uint16x8_t(v0), C.int16x8_t(v1), C.int16x8_t(v2)))
}

func VbslU8(v0 Uint8X8, v1 Uint8X8, v2 Uint8X8) Uint8X8 {
	return Uint8X8(C.vbsl_u8(C.uint8x8_t(v0), C.uint8x8_t(v1), C.uint8x8_t(v2)))
}

func VbslU32(v0 Uint32X2, v1 Uint32X2, v2 Uint32X2) Uint32X2 {
	return Uint32X2(C.vbsl_u32(C.uint32x2_t(v0), C.uint32x2_t(v1), C.uint32x2_t(v2)))
}

func VbslU64(v0 Uint64X1, v1 Uint64X1, v2 Uint64X1) Uint64X1 {
	return Uint64X1(C.vbsl_u64(C.uint64x1_t(v0), C.uint64x1_t(v1), C.uint64x1_t(v2)))
}

func VbslU16(v0 Uint16X4, v1 Uint16X4, v2 Uint16X4) Uint16X4 {
	return Uint16X4(C.vbsl_u16(C.uint16x4_t(v0), C.uint16x4_t(v1), C.uint16x4_t(v2)))
}

func VbslS8(v0 Uint8X8, v1 Int8X8, v2 Int8X8) Int8X8 {
	return Int8X8(C.vbsl_s8(C.uint8x8_t(v0), C.int8x8_t(v1), C.int8x8_t(v2)))
}

func VbslF32(v0 Uint32X2, v1 Float32X2, v2 Float32X2) Float32X2 {
	return Float32X2(C.vbsl_f32(C.uint32x2_t(v0), C.float32x2_t(v1), C.float32x2_t(v2)))
}

func VbslS32(v0 Uint32X2, v1 Int32X2, v2 Int32X2) Int32X2 {
	return Int32X2(C.vbsl_s32(C.uint32x2_t(v0), C.int32x2_t(v1), C.int32x2_t(v2)))
}

func VbslS64(v0 Uint64X1, v1 Int64X1, v2 Int64X1) Int64X1 {
	return Int64X1(C.vbsl_s64(C.uint64x1_t(v0), C.int64x1_t(v1), C.int64x1_t(v2)))
}

func VbslS16(v0 Uint16X4, v1 Int16X4, v2 Int16X4) Int16X4 {
	return Int16X4(C.vbsl_s16(C.uint16x4_t(v0), C.int16x4_t(v1), C.int16x4_t(v2)))
}

func VcageqF32(v0 Float32X4, v1 Float32X4) Uint32X4 {
	return Uint32X4(C.vcageq_f32(C.float32x4_t(v0), C.float32x4_t(v1)))
}

func VcageF32(v0 Float32X2, v1 Float32X2) Uint32X2 {
	return Uint32X2(C.vcage_f32(C.float32x2_t(v0), C.float32x2_t(v1)))
}

func VcagtqF32(v0 Float32X4, v1 Float32X4) Uint32X4 {
	return Uint32X4(C.vcagtq_f32(C.float32x4_t(v0), C.float32x4_t(v1)))
}

func VcagtF32(v0 Float32X2, v1 Float32X2) Uint32X2 {
	return Uint32X2(C.vcagt_f32(C.float32x2_t(v0), C.float32x2_t(v1)))
}

func VcaleqF32(v0 Float32X4, v1 Float32X4) Uint32X4 {
	return Uint32X4(C.vcaleq_f32(C.float32x4_t(v0), C.float32x4_t(v1)))
}

func VcaleF32(v0 Float32X2, v1 Float32X2) Uint32X2 {
	return Uint32X2(C.vcale_f32(C.float32x2_t(v0), C.float32x2_t(v1)))
}

func VcaltqF32(v0 Float32X4, v1 Float32X4) Uint32X4 {
	return Uint32X4(C.vcaltq_f32(C.float32x4_t(v0), C.float32x4_t(v1)))
}

func VcaltF32(v0 Float32X2, v1 Float32X2) Uint32X2 {
	return Uint32X2(C.vcalt_f32(C.float32x2_t(v0), C.float32x2_t(v1)))
}

func VceqP8(v0 Poly8X8, v1 Poly8X8) Uint8X8 {
	return Uint8X8(C.vceq_p8(C.poly8x8_t(v0), C.poly8x8_t(v1)))
}

func VceqqP8(v0 Poly8X16, v1 Poly8X16) Uint8X16 {
	return Uint8X16(C.vceqq_p8(C.poly8x16_t(v0), C.poly8x16_t(v1)))
}

func VceqqU8(v0 Uint8X16, v1 Uint8X16) Uint8X16 {
	return Uint8X16(C.vceqq_u8(C.uint8x16_t(v0), C.uint8x16_t(v1)))
}

func VceqqU32(v0 Uint32X4, v1 Uint32X4) Uint32X4 {
	return Uint32X4(C.vceqq_u32(C.uint32x4_t(v0), C.uint32x4_t(v1)))
}

func VceqqU16(v0 Uint16X8, v1 Uint16X8) Uint16X8 {
	return Uint16X8(C.vceqq_u16(C.uint16x8_t(v0), C.uint16x8_t(v1)))
}

func VceqqS8(v0 Int8X16, v1 Int8X16) Uint8X16 {
	return Uint8X16(C.vceqq_s8(C.int8x16_t(v0), C.int8x16_t(v1)))
}

func VceqqF32(v0 Float32X4, v1 Float32X4) Uint32X4 {
	return Uint32X4(C.vceqq_f32(C.float32x4_t(v0), C.float32x4_t(v1)))
}

func VceqqS32(v0 Int32X4, v1 Int32X4) Uint32X4 {
	return Uint32X4(C.vceqq_s32(C.int32x4_t(v0), C.int32x4_t(v1)))
}

func VceqqS16(v0 Int16X8, v1 Int16X8) Uint16X8 {
	return Uint16X8(C.vceqq_s16(C.int16x8_t(v0), C.int16x8_t(v1)))
}

func VceqU8(v0 Uint8X8, v1 Uint8X8) Uint8X8 {
	return Uint8X8(C.vceq_u8(C.uint8x8_t(v0), C.uint8x8_t(v1)))
}

func VceqU32(v0 Uint32X2, v1 Uint32X2) Uint32X2 {
	return Uint32X2(C.vceq_u32(C.uint32x2_t(v0), C.uint32x2_t(v1)))
}

func VceqU16(v0 Uint16X4, v1 Uint16X4) Uint16X4 {
	return Uint16X4(C.vceq_u16(C.uint16x4_t(v0), C.uint16x4_t(v1)))
}

func VceqS8(v0 Int8X8, v1 Int8X8) Uint8X8 {
	return Uint8X8(C.vceq_s8(C.int8x8_t(v0), C.int8x8_t(v1)))
}

func VceqF32(v0 Float32X2, v1 Float32X2) Uint32X2 {
	return Uint32X2(C.vceq_f32(C.float32x2_t(v0), C.float32x2_t(v1)))
}

func VceqS32(v0 Int32X2, v1 Int32X2) Uint32X2 {
	return Uint32X2(C.vceq_s32(C.int32x2_t(v0), C.int32x2_t(v1)))
}

func VceqS16(v0 Int16X4, v1 Int16X4) Uint16X4 {
	return Uint16X4(C.vceq_s16(C.int16x4_t(v0), C.int16x4_t(v1)))
}

func VcgeqU8(v0 Uint8X16, v1 Uint8X16) Uint8X16 {
	return Uint8X16(C.vcgeq_u8(C.uint8x16_t(v0), C.uint8x16_t(v1)))
}

func VcgeqU32(v0 Uint32X4, v1 Uint32X4) Uint32X4 {
	return Uint32X4(C.vcgeq_u32(C.uint32x4_t(v0), C.uint32x4_t(v1)))
}

func VcgeqU16(v0 Uint16X8, v1 Uint16X8) Uint16X8 {
	return Uint16X8(C.vcgeq_u16(C.uint16x8_t(v0), C.uint16x8_t(v1)))
}

func VcgeqS8(v0 Int8X16, v1 Int8X16) Uint8X16 {
	return Uint8X16(C.vcgeq_s8(C.int8x16_t(v0), C.int8x16_t(v1)))
}

func VcgeqF32(v0 Float32X4, v1 Float32X4) Uint32X4 {
	return Uint32X4(C.vcgeq_f32(C.float32x4_t(v0), C.float32x4_t(v1)))
}

func VcgeqS32(v0 Int32X4, v1 Int32X4) Uint32X4 {
	return Uint32X4(C.vcgeq_s32(C.int32x4_t(v0), C.int32x4_t(v1)))
}

func VcgeqS16(v0 Int16X8, v1 Int16X8) Uint16X8 {
	return Uint16X8(C.vcgeq_s16(C.int16x8_t(v0), C.int16x8_t(v1)))
}

func VcgeU8(v0 Uint8X8, v1 Uint8X8) Uint8X8 {
	return Uint8X8(C.vcge_u8(C.uint8x8_t(v0), C.uint8x8_t(v1)))
}

func VcgeU32(v0 Uint32X2, v1 Uint32X2) Uint32X2 {
	return Uint32X2(C.vcge_u32(C.uint32x2_t(v0), C.uint32x2_t(v1)))
}

func VcgeU16(v0 Uint16X4, v1 Uint16X4) Uint16X4 {
	return Uint16X4(C.vcge_u16(C.uint16x4_t(v0), C.uint16x4_t(v1)))
}

func VcgeS8(v0 Int8X8, v1 Int8X8) Uint8X8 {
	return Uint8X8(C.vcge_s8(C.int8x8_t(v0), C.int8x8_t(v1)))
}

func VcgeF32(v0 Float32X2, v1 Float32X2) Uint32X2 {
	return Uint32X2(C.vcge_f32(C.float32x2_t(v0), C.float32x2_t(v1)))
}

func VcgeS32(v0 Int32X2, v1 Int32X2) Uint32X2 {
	return Uint32X2(C.vcge_s32(C.int32x2_t(v0), C.int32x2_t(v1)))
}

func VcgeS16(v0 Int16X4, v1 Int16X4) Uint16X4 {
	return Uint16X4(C.vcge_s16(C.int16x4_t(v0), C.int16x4_t(v1)))
}

func VcgtqU8(v0 Uint8X16, v1 Uint8X16) Uint8X16 {
	return Uint8X16(C.vcgtq_u8(C.uint8x16_t(v0), C.uint8x16_t(v1)))
}

func VcgtqU32(v0 Uint32X4, v1 Uint32X4) Uint32X4 {
	return Uint32X4(C.vcgtq_u32(C.uint32x4_t(v0), C.uint32x4_t(v1)))
}

func VcgtqU16(v0 Uint16X8, v1 Uint16X8) Uint16X8 {
	return Uint16X8(C.vcgtq_u16(C.uint16x8_t(v0), C.uint16x8_t(v1)))
}

func VcgtqS8(v0 Int8X16, v1 Int8X16) Uint8X16 {
	return Uint8X16(C.vcgtq_s8(C.int8x16_t(v0), C.int8x16_t(v1)))
}

func VcgtqF32(v0 Float32X4, v1 Float32X4) Uint32X4 {
	return Uint32X4(C.vcgtq_f32(C.float32x4_t(v0), C.float32x4_t(v1)))
}

func VcgtqS32(v0 Int32X4, v1 Int32X4) Uint32X4 {
	return Uint32X4(C.vcgtq_s32(C.int32x4_t(v0), C.int32x4_t(v1)))
}

func VcgtqS16(v0 Int16X8, v1 Int16X8) Uint16X8 {
	return Uint16X8(C.vcgtq_s16(C.int16x8_t(v0), C.int16x8_t(v1)))
}

func VcgtU8(v0 Uint8X8, v1 Uint8X8) Uint8X8 {
	return Uint8X8(C.vcgt_u8(C.uint8x8_t(v0), C.uint8x8_t(v1)))
}

func VcgtU32(v0 Uint32X2, v1 Uint32X2) Uint32X2 {
	return Uint32X2(C.vcgt_u32(C.uint32x2_t(v0), C.uint32x2_t(v1)))
}

func VcgtU16(v0 Uint16X4, v1 Uint16X4) Uint16X4 {
	return Uint16X4(C.vcgt_u16(C.uint16x4_t(v0), C.uint16x4_t(v1)))
}

func VcgtS8(v0 Int8X8, v1 Int8X8) Uint8X8 {
	return Uint8X8(C.vcgt_s8(C.int8x8_t(v0), C.int8x8_t(v1)))
}

func VcgtF32(v0 Float32X2, v1 Float32X2) Uint32X2 {
	return Uint32X2(C.vcgt_f32(C.float32x2_t(v0), C.float32x2_t(v1)))
}

func VcgtS32(v0 Int32X2, v1 Int32X2) Uint32X2 {
	return Uint32X2(C.vcgt_s32(C.int32x2_t(v0), C.int32x2_t(v1)))
}

func VcgtS16(v0 Int16X4, v1 Int16X4) Uint16X4 {
	return Uint16X4(C.vcgt_s16(C.int16x4_t(v0), C.int16x4_t(v1)))
}

func VcleqU8(v0 Uint8X16, v1 Uint8X16) Uint8X16 {
	return Uint8X16(C.vcleq_u8(C.uint8x16_t(v0), C.uint8x16_t(v1)))
}

func VcleqU32(v0 Uint32X4, v1 Uint32X4) Uint32X4 {
	return Uint32X4(C.vcleq_u32(C.uint32x4_t(v0), C.uint32x4_t(v1)))
}

func VcleqU16(v0 Uint16X8, v1 Uint16X8) Uint16X8 {
	return Uint16X8(C.vcleq_u16(C.uint16x8_t(v0), C.uint16x8_t(v1)))
}

func VcleqS8(v0 Int8X16, v1 Int8X16) Uint8X16 {
	return Uint8X16(C.vcleq_s8(C.int8x16_t(v0), C.int8x16_t(v1)))
}

func VcleqF32(v0 Float32X4, v1 Float32X4) Uint32X4 {
	return Uint32X4(C.vcleq_f32(C.float32x4_t(v0), C.float32x4_t(v1)))
}

func VcleqS32(v0 Int32X4, v1 Int32X4) Uint32X4 {
	return Uint32X4(C.vcleq_s32(C.int32x4_t(v0), C.int32x4_t(v1)))
}

func VcleqS16(v0 Int16X8, v1 Int16X8) Uint16X8 {
	return Uint16X8(C.vcleq_s16(C.int16x8_t(v0), C.int16x8_t(v1)))
}

func VcleU8(v0 Uint8X8, v1 Uint8X8) Uint8X8 {
	return Uint8X8(C.vcle_u8(C.uint8x8_t(v0), C.uint8x8_t(v1)))
}

func VcleU32(v0 Uint32X2, v1 Uint32X2) Uint32X2 {
	return Uint32X2(C.vcle_u32(C.uint32x2_t(v0), C.uint32x2_t(v1)))
}

func VcleU16(v0 Uint16X4, v1 Uint16X4) Uint16X4 {
	return Uint16X4(C.vcle_u16(C.uint16x4_t(v0), C.uint16x4_t(v1)))
}

func VcleS8(v0 Int8X8, v1 Int8X8) Uint8X8 {
	return Uint8X8(C.vcle_s8(C.int8x8_t(v0), C.int8x8_t(v1)))
}

func VcleF32(v0 Float32X2, v1 Float32X2) Uint32X2 {
	return Uint32X2(C.vcle_f32(C.float32x2_t(v0), C.float32x2_t(v1)))
}

func VcleS32(v0 Int32X2, v1 Int32X2) Uint32X2 {
	return Uint32X2(C.vcle_s32(C.int32x2_t(v0), C.int32x2_t(v1)))
}

func VcleS16(v0 Int16X4, v1 Int16X4) Uint16X4 {
	return Uint16X4(C.vcle_s16(C.int16x4_t(v0), C.int16x4_t(v1)))
}

func VclsqU8(v0 Uint8X16) Int8X16 {
	return Int8X16(C.vclsq_u8(C.uint8x16_t(v0)))
}

func VclsqU32(v0 Uint32X4) Int32X4 {
	return Int32X4(C.vclsq_u32(C.uint32x4_t(v0)))
}

func VclsqU16(v0 Uint16X8) Int16X8 {
	return Int16X8(C.vclsq_u16(C.uint16x8_t(v0)))
}

func VclsqS8(v0 Int8X16) Int8X16 {
	return Int8X16(C.vclsq_s8(C.int8x16_t(v0)))
}

func VclsqS32(v0 Int32X4) Int32X4 {
	return Int32X4(C.vclsq_s32(C.int32x4_t(v0)))
}

func VclsqS16(v0 Int16X8) Int16X8 {
	return Int16X8(C.vclsq_s16(C.int16x8_t(v0)))
}

func VclsU8(v0 Uint8X8) Int8X8 {
	return Int8X8(C.vcls_u8(C.uint8x8_t(v0)))
}

func VclsU32(v0 Uint32X2) Int32X2 {
	return Int32X2(C.vcls_u32(C.uint32x2_t(v0)))
}

func VclsU16(v0 Uint16X4) Int16X4 {
	return Int16X4(C.vcls_u16(C.uint16x4_t(v0)))
}

func VclsS8(v0 Int8X8) Int8X8 {
	return Int8X8(C.vcls_s8(C.int8x8_t(v0)))
}

func VclsS32(v0 Int32X2) Int32X2 {
	return Int32X2(C.vcls_s32(C.int32x2_t(v0)))
}

func VclsS16(v0 Int16X4) Int16X4 {
	return Int16X4(C.vcls_s16(C.int16x4_t(v0)))
}

func VcltqU8(v0 Uint8X16, v1 Uint8X16) Uint8X16 {
	return Uint8X16(C.vcltq_u8(C.uint8x16_t(v0), C.uint8x16_t(v1)))
}

func VcltqU32(v0 Uint32X4, v1 Uint32X4) Uint32X4 {
	return Uint32X4(C.vcltq_u32(C.uint32x4_t(v0), C.uint32x4_t(v1)))
}

func VcltqU16(v0 Uint16X8, v1 Uint16X8) Uint16X8 {
	return Uint16X8(C.vcltq_u16(C.uint16x8_t(v0), C.uint16x8_t(v1)))
}

func VcltqS8(v0 Int8X16, v1 Int8X16) Uint8X16 {
	return Uint8X16(C.vcltq_s8(C.int8x16_t(v0), C.int8x16_t(v1)))
}

func VcltqF32(v0 Float32X4, v1 Float32X4) Uint32X4 {
	return Uint32X4(C.vcltq_f32(C.float32x4_t(v0), C.float32x4_t(v1)))
}

func VcltqS32(v0 Int32X4, v1 Int32X4) Uint32X4 {
	return Uint32X4(C.vcltq_s32(C.int32x4_t(v0), C.int32x4_t(v1)))
}

func VcltqS16(v0 Int16X8, v1 Int16X8) Uint16X8 {
	return Uint16X8(C.vcltq_s16(C.int16x8_t(v0), C.int16x8_t(v1)))
}

func VcltU8(v0 Uint8X8, v1 Uint8X8) Uint8X8 {
	return Uint8X8(C.vclt_u8(C.uint8x8_t(v0), C.uint8x8_t(v1)))
}

func VcltU32(v0 Uint32X2, v1 Uint32X2) Uint32X2 {
	return Uint32X2(C.vclt_u32(C.uint32x2_t(v0), C.uint32x2_t(v1)))
}

func VcltU16(v0 Uint16X4, v1 Uint16X4) Uint16X4 {
	return Uint16X4(C.vclt_u16(C.uint16x4_t(v0), C.uint16x4_t(v1)))
}

func VcltS8(v0 Int8X8, v1 Int8X8) Uint8X8 {
	return Uint8X8(C.vclt_s8(C.int8x8_t(v0), C.int8x8_t(v1)))
}

func VcltF32(v0 Float32X2, v1 Float32X2) Uint32X2 {
	return Uint32X2(C.vclt_f32(C.float32x2_t(v0), C.float32x2_t(v1)))
}

func VcltS32(v0 Int32X2, v1 Int32X2) Uint32X2 {
	return Uint32X2(C.vclt_s32(C.int32x2_t(v0), C.int32x2_t(v1)))
}

func VcltS16(v0 Int16X4, v1 Int16X4) Uint16X4 {
	return Uint16X4(C.vclt_s16(C.int16x4_t(v0), C.int16x4_t(v1)))
}

func VclzqU8(v0 Uint8X16) Uint8X16 {
	return Uint8X16(C.vclzq_u8(C.uint8x16_t(v0)))
}

func VclzqU32(v0 Uint32X4) Uint32X4 {
	return Uint32X4(C.vclzq_u32(C.uint32x4_t(v0)))
}

func VclzqU16(v0 Uint16X8) Uint16X8 {
	return Uint16X8(C.vclzq_u16(C.uint16x8_t(v0)))
}

func VclzqS8(v0 Int8X16) Int8X16 {
	return Int8X16(C.vclzq_s8(C.int8x16_t(v0)))
}

func VclzqS32(v0 Int32X4) Int32X4 {
	return Int32X4(C.vclzq_s32(C.int32x4_t(v0)))
}

func VclzqS16(v0 Int16X8) Int16X8 {
	return Int16X8(C.vclzq_s16(C.int16x8_t(v0)))
}

func VclzU8(v0 Uint8X8) Uint8X8 {
	return Uint8X8(C.vclz_u8(C.uint8x8_t(v0)))
}

func VclzU32(v0 Uint32X2) Uint32X2 {
	return Uint32X2(C.vclz_u32(C.uint32x2_t(v0)))
}

func VclzU16(v0 Uint16X4) Uint16X4 {
	return Uint16X4(C.vclz_u16(C.uint16x4_t(v0)))
}

func VclzS8(v0 Int8X8) Int8X8 {
	return Int8X8(C.vclz_s8(C.int8x8_t(v0)))
}

func VclzS32(v0 Int32X2) Int32X2 {
	return Int32X2(C.vclz_s32(C.int32x2_t(v0)))
}

func VclzS16(v0 Int16X4) Int16X4 {
	return Int16X4(C.vclz_s16(C.int16x4_t(v0)))
}

func VcntP8(v0 Poly8X8) Poly8X8 {
	return Poly8X8(C.vcnt_p8(C.poly8x8_t(v0)))
}

func VcntqP8(v0 Poly8X16) Poly8X16 {
	return Poly8X16(C.vcntq_p8(C.poly8x16_t(v0)))
}

func VcntqU8(v0 Uint8X16) Uint8X16 {
	return Uint8X16(C.vcntq_u8(C.uint8x16_t(v0)))
}

func VcntqS8(v0 Int8X16) Int8X16 {
	return Int8X16(C.vcntq_s8(C.int8x16_t(v0)))
}

func VcntU8(v0 Uint8X8) Uint8X8 {
	return Uint8X8(C.vcnt_u8(C.uint8x8_t(v0)))
}

func VcntS8(v0 Int8X8) Int8X8 {
	return Int8X8(C.vcnt_s8(C.int8x8_t(v0)))
}

func VcombineP8(v0 Poly8X8, v1 Poly8X8) Poly8X16 {
	return Poly8X16(C.vcombine_p8(C.poly8x8_t(v0), C.poly8x8_t(v1)))
}

func VcombineP16(v0 Poly16X4, v1 Poly16X4) Poly16X8 {
	return Poly16X8(C.vcombine_p16(C.poly16x4_t(v0), C.poly16x4_t(v1)))
}

func VcombineU8(v0 Uint8X8, v1 Uint8X8) Uint8X16 {
	return Uint8X16(C.vcombine_u8(C.uint8x8_t(v0), C.uint8x8_t(v1)))
}

func VcombineU32(v0 Uint32X2, v1 Uint32X2) Uint32X4 {
	return Uint32X4(C.vcombine_u32(C.uint32x2_t(v0), C.uint32x2_t(v1)))
}

func VcombineU64(v0 Uint64X1, v1 Uint64X1) Uint64X2 {
	return Uint64X2(C.vcombine_u64(C.uint64x1_t(v0), C.uint64x1_t(v1)))
}

func VcombineU16(v0 Uint16X4, v1 Uint16X4) Uint16X8 {
	return Uint16X8(C.vcombine_u16(C.uint16x4_t(v0), C.uint16x4_t(v1)))
}

func VcombineS8(v0 Int8X8, v1 Int8X8) Int8X16 {
	return Int8X16(C.vcombine_s8(C.int8x8_t(v0), C.int8x8_t(v1)))
}

func VcombineF32(v0 Float32X2, v1 Float32X2) Float32X4 {
	return Float32X4(C.vcombine_f32(C.float32x2_t(v0), C.float32x2_t(v1)))
}

func VcombineS32(v0 Int32X2, v1 Int32X2) Int32X4 {
	return Int32X4(C.vcombine_s32(C.int32x2_t(v0), C.int32x2_t(v1)))
}

func VcombineS64(v0 Int64X1, v1 Int64X1) Int64X2 {
	return Int64X2(C.vcombine_s64(C.int64x1_t(v0), C.int64x1_t(v1)))
}

func VcombineS16(v0 Int16X4, v1 Int16X4) Int16X8 {
	return Int16X8(C.vcombine_s16(C.int16x4_t(v0), C.int16x4_t(v1)))
}

func VcvtqF32U32(v0 Uint32X4) Float32X4 {
	return Float32X4(C.vcvtq_f32_u32(C.uint32x4_t(v0)))
}

func VcvtqF32S32(v0 Int32X4) Float32X4 {
	return Float32X4(C.vcvtq_f32_s32(C.int32x4_t(v0)))
}

func VcvtF32U32(v0 Uint32X2) Float32X2 {
	return Float32X2(C.vcvt_f32_u32(C.uint32x2_t(v0)))
}

func VcvtF32S32(v0 Int32X2) Float32X2 {
	return Float32X2(C.vcvt_f32_s32(C.int32x2_t(v0)))
}

func VcvtqS32F32(v0 Float32X4) Int32X4 {
	return Int32X4(C.vcvtq_s32_f32(C.float32x4_t(v0)))
}

func VcvtS32F32(v0 Float32X2) Int32X2 {
	return Int32X2(C.vcvt_s32_f32(C.float32x2_t(v0)))
}

func VcvtqU32F32(v0 Float32X4) Uint32X4 {
	return Uint32X4(C.vcvtq_u32_f32(C.float32x4_t(v0)))
}

func VcvtU32F32(v0 Float32X2) Uint32X2 {
	return Uint32X2(C.vcvt_u32_f32(C.float32x2_t(v0)))
}

func VdupNP8(v0 Poly8) Poly8X8 {
	return Poly8X8(C.vdup_n_p8(C.poly8_t(v0)))
}

func VdupNP16(v0 Poly16) Poly16X4 {
	return Poly16X4(C.vdup_n_p16(C.poly16_t(v0)))
}

func VdupqNP8(v0 Poly8) Poly8X16 {
	return Poly8X16(C.vdupq_n_p8(C.poly8_t(v0)))
}

func VdupqNP16(v0 Poly16) Poly16X8 {
	return Poly16X8(C.vdupq_n_p16(C.poly16_t(v0)))
}

func VdupqNU8(v0 uint8) Uint8X16 {
	return Uint8X16(C.vdupq_n_u8(C.uint8_t(v0)))
}

func VdupqNU32(v0 uint32) Uint32X4 {
	return Uint32X4(C.vdupq_n_u32(C.uint32_t(v0)))
}

func VdupqNU64(v0 uint64) Uint64X2 {
	return Uint64X2(C.vdupq_n_u64(C.uint64_t(v0)))
}

func VdupqNU16(v0 uint16) Uint16X8 {
	return Uint16X8(C.vdupq_n_u16(C.uint16_t(v0)))
}

func VdupqNS8(v0 int8) Int8X16 {
	return Int8X16(C.vdupq_n_s8(C.int8_t(v0)))
}

func VdupqNF32(v0 float32) Float32X4 {
	return Float32X4(C.vdupq_n_f32(C.float32_t(v0)))
}

func VdupqNS32(v0 int32) Int32X4 {
	return Int32X4(C.vdupq_n_s32(C.int32_t(v0)))
}

func VdupqNS64(v0 int64) Int64X2 {
	return Int64X2(C.vdupq_n_s64(C.int64_t(v0)))
}

func VdupqNS16(v0 int16) Int16X8 {
	return Int16X8(C.vdupq_n_s16(C.int16_t(v0)))
}

func VdupNU8(v0 uint8) Uint8X8 {
	return Uint8X8(C.vdup_n_u8(C.uint8_t(v0)))
}

func VdupNU32(v0 uint32) Uint32X2 {
	return Uint32X2(C.vdup_n_u32(C.uint32_t(v0)))
}

func VdupNU64(v0 uint64) Uint64X1 {
	return Uint64X1(C.vdup_n_u64(C.uint64_t(v0)))
}

func VdupNU16(v0 uint16) Uint16X4 {
	return Uint16X4(C.vdup_n_u16(C.uint16_t(v0)))
}

func VdupNS8(v0 int8) Int8X8 {
	return Int8X8(C.vdup_n_s8(C.int8_t(v0)))
}

func VdupNF32(v0 float32) Float32X2 {
	return Float32X2(C.vdup_n_f32(C.float32_t(v0)))
}

func VdupNS32(v0 int32) Int32X2 {
	return Int32X2(C.vdup_n_s32(C.int32_t(v0)))
}

func VdupNS64(v0 int64) Int64X1 {
	return Int64X1(C.vdup_n_s64(C.int64_t(v0)))
}

func VdupNS16(v0 int16) Int16X4 {
	return Int16X4(C.vdup_n_s16(C.int16_t(v0)))
}

func VeorqU8(v0 Uint8X16, v1 Uint8X16) Uint8X16 {
	return Uint8X16(C.veorq_u8(C.uint8x16_t(v0), C.uint8x16_t(v1)))
}

func VeorqU32(v0 Uint32X4, v1 Uint32X4) Uint32X4 {
	return Uint32X4(C.veorq_u32(C.uint32x4_t(v0), C.uint32x4_t(v1)))
}

func VeorqU64(v0 Uint64X2, v1 Uint64X2) Uint64X2 {
	return Uint64X2(C.veorq_u64(C.uint64x2_t(v0), C.uint64x2_t(v1)))
}

func VeorqU16(v0 Uint16X8, v1 Uint16X8) Uint16X8 {
	return Uint16X8(C.veorq_u16(C.uint16x8_t(v0), C.uint16x8_t(v1)))
}

func VeorqS8(v0 Int8X16, v1 Int8X16) Int8X16 {
	return Int8X16(C.veorq_s8(C.int8x16_t(v0), C.int8x16_t(v1)))
}

func VeorqS32(v0 Int32X4, v1 Int32X4) Int32X4 {
	return Int32X4(C.veorq_s32(C.int32x4_t(v0), C.int32x4_t(v1)))
}

func VeorqS64(v0 Int64X2, v1 Int64X2) Int64X2 {
	return Int64X2(C.veorq_s64(C.int64x2_t(v0), C.int64x2_t(v1)))
}

func VeorqS16(v0 Int16X8, v1 Int16X8) Int16X8 {
	return Int16X8(C.veorq_s16(C.int16x8_t(v0), C.int16x8_t(v1)))
}

func VeorU8(v0 Uint8X8, v1 Uint8X8) Uint8X8 {
	return Uint8X8(C.veor_u8(C.uint8x8_t(v0), C.uint8x8_t(v1)))
}

func VeorU32(v0 Uint32X2, v1 Uint32X2) Uint32X2 {
	return Uint32X2(C.veor_u32(C.uint32x2_t(v0), C.uint32x2_t(v1)))
}

func VeorU64(v0 Uint64X1, v1 Uint64X1) Uint64X1 {
	return Uint64X1(C.veor_u64(C.uint64x1_t(v0), C.uint64x1_t(v1)))
}

func VeorU16(v0 Uint16X4, v1 Uint16X4) Uint16X4 {
	return Uint16X4(C.veor_u16(C.uint16x4_t(v0), C.uint16x4_t(v1)))
}

func VeorS8(v0 Int8X8, v1 Int8X8) Int8X8 {
	return Int8X8(C.veor_s8(C.int8x8_t(v0), C.int8x8_t(v1)))
}

func VeorS32(v0 Int32X2, v1 Int32X2) Int32X2 {
	return Int32X2(C.veor_s32(C.int32x2_t(v0), C.int32x2_t(v1)))
}

func VeorS64(v0 Int64X1, v1 Int64X1) Int64X1 {
	return Int64X1(C.veor_s64(C.int64x1_t(v0), C.int64x1_t(v1)))
}

func VeorS16(v0 Int16X4, v1 Int16X4) Int16X4 {
	return Int16X4(C.veor_s16(C.int16x4_t(v0), C.int16x4_t(v1)))
}

func VgetHighP8(v0 Poly8X16) Poly8X8 {
	return Poly8X8(C.vget_high_p8(C.poly8x16_t(v0)))
}

func VgetHighP16(v0 Poly16X8) Poly16X4 {
	return Poly16X4(C.vget_high_p16(C.poly16x8_t(v0)))
}

func VgetHighU8(v0 Uint8X16) Uint8X8 {
	return Uint8X8(C.vget_high_u8(C.uint8x16_t(v0)))
}

func VgetHighU32(v0 Uint32X4) Uint32X2 {
	return Uint32X2(C.vget_high_u32(C.uint32x4_t(v0)))
}

func VgetHighU64(v0 Uint64X2) Uint64X1 {
	return Uint64X1(C.vget_high_u64(C.uint64x2_t(v0)))
}

func VgetHighU16(v0 Uint16X8) Uint16X4 {
	return Uint16X4(C.vget_high_u16(C.uint16x8_t(v0)))
}

func VgetHighS8(v0 Int8X16) Int8X8 {
	return Int8X8(C.vget_high_s8(C.int8x16_t(v0)))
}

func VgetHighF32(v0 Float32X4) Float32X2 {
	return Float32X2(C.vget_high_f32(C.float32x4_t(v0)))
}

func VgetHighS32(v0 Int32X4) Int32X2 {
	return Int32X2(C.vget_high_s32(C.int32x4_t(v0)))
}

func VgetHighS64(v0 Int64X2) Int64X1 {
	return Int64X1(C.vget_high_s64(C.int64x2_t(v0)))
}

func VgetHighS16(v0 Int16X8) Int16X4 {
	return Int16X4(C.vget_high_s16(C.int16x8_t(v0)))
}

func VgetLowP8(v0 Poly8X16) Poly8X8 {
	return Poly8X8(C.vget_low_p8(C.poly8x16_t(v0)))
}

func VgetLowP16(v0 Poly16X8) Poly16X4 {
	return Poly16X4(C.vget_low_p16(C.poly16x8_t(v0)))
}

func VgetLowU8(v0 Uint8X16) Uint8X8 {
	return Uint8X8(C.vget_low_u8(C.uint8x16_t(v0)))
}

func VgetLowU32(v0 Uint32X4) Uint32X2 {
	return Uint32X2(C.vget_low_u32(C.uint32x4_t(v0)))
}

func VgetLowU64(v0 Uint64X2) Uint64X1 {
	return Uint64X1(C.vget_low_u64(C.uint64x2_t(v0)))
}

func VgetLowU16(v0 Uint16X8) Uint16X4 {
	return Uint16X4(C.vget_low_u16(C.uint16x8_t(v0)))
}

func VgetLowS8(v0 Int8X16) Int8X8 {
	return Int8X8(C.vget_low_s8(C.int8x16_t(v0)))
}

func VgetLowF32(v0 Float32X4) Float32X2 {
	return Float32X2(C.vget_low_f32(C.float32x4_t(v0)))
}

func VgetLowS32(v0 Int32X4) Int32X2 {
	return Int32X2(C.vget_low_s32(C.int32x4_t(v0)))
}

func VgetLowS64(v0 Int64X2) Int64X1 {
	return Int64X1(C.vget_low_s64(C.int64x2_t(v0)))
}

func VgetLowS16(v0 Int16X8) Int16X4 {
	return Int16X4(C.vget_low_s16(C.int16x8_t(v0)))
}

func VhaddqU8(v0 Uint8X16, v1 Uint8X16) Uint8X16 {
	return Uint8X16(C.vhaddq_u8(C.uint8x16_t(v0), C.uint8x16_t(v1)))
}

func VhaddqU32(v0 Uint32X4, v1 Uint32X4) Uint32X4 {
	return Uint32X4(C.vhaddq_u32(C.uint32x4_t(v0), C.uint32x4_t(v1)))
}

func VhaddqU16(v0 Uint16X8, v1 Uint16X8) Uint16X8 {
	return Uint16X8(C.vhaddq_u16(C.uint16x8_t(v0), C.uint16x8_t(v1)))
}

func VhaddqS8(v0 Int8X16, v1 Int8X16) Int8X16 {
	return Int8X16(C.vhaddq_s8(C.int8x16_t(v0), C.int8x16_t(v1)))
}

func VhaddqS32(v0 Int32X4, v1 Int32X4) Int32X4 {
	return Int32X4(C.vhaddq_s32(C.int32x4_t(v0), C.int32x4_t(v1)))
}

func VhaddqS16(v0 Int16X8, v1 Int16X8) Int16X8 {
	return Int16X8(C.vhaddq_s16(C.int16x8_t(v0), C.int16x8_t(v1)))
}

func VhaddU8(v0 Uint8X8, v1 Uint8X8) Uint8X8 {
	return Uint8X8(C.vhadd_u8(C.uint8x8_t(v0), C.uint8x8_t(v1)))
}

func VhaddU32(v0 Uint32X2, v1 Uint32X2) Uint32X2 {
	return Uint32X2(C.vhadd_u32(C.uint32x2_t(v0), C.uint32x2_t(v1)))
}

func VhaddU16(v0 Uint16X4, v1 Uint16X4) Uint16X4 {
	return Uint16X4(C.vhadd_u16(C.uint16x4_t(v0), C.uint16x4_t(v1)))
}

func VhaddS8(v0 Int8X8, v1 Int8X8) Int8X8 {
	return Int8X8(C.vhadd_s8(C.int8x8_t(v0), C.int8x8_t(v1)))
}

func VhaddS32(v0 Int32X2, v1 Int32X2) Int32X2 {
	return Int32X2(C.vhadd_s32(C.int32x2_t(v0), C.int32x2_t(v1)))
}

func VhaddS16(v0 Int16X4, v1 Int16X4) Int16X4 {
	return Int16X4(C.vhadd_s16(C.int16x4_t(v0), C.int16x4_t(v1)))
}

func VhsubqU8(v0 Uint8X16, v1 Uint8X16) Uint8X16 {
	return Uint8X16(C.vhsubq_u8(C.uint8x16_t(v0), C.uint8x16_t(v1)))
}

func VhsubqU32(v0 Uint32X4, v1 Uint32X4) Uint32X4 {
	return Uint32X4(C.vhsubq_u32(C.uint32x4_t(v0), C.uint32x4_t(v1)))
}

func VhsubqU16(v0 Uint16X8, v1 Uint16X8) Uint16X8 {
	return Uint16X8(C.vhsubq_u16(C.uint16x8_t(v0), C.uint16x8_t(v1)))
}

func VhsubqS8(v0 Int8X16, v1 Int8X16) Int8X16 {
	return Int8X16(C.vhsubq_s8(C.int8x16_t(v0), C.int8x16_t(v1)))
}

func VhsubqS32(v0 Int32X4, v1 Int32X4) Int32X4 {
	return Int32X4(C.vhsubq_s32(C.int32x4_t(v0), C.int32x4_t(v1)))
}

func VhsubqS16(v0 Int16X8, v1 Int16X8) Int16X8 {
	return Int16X8(C.vhsubq_s16(C.int16x8_t(v0), C.int16x8_t(v1)))
}

func VhsubU8(v0 Uint8X8, v1 Uint8X8) Uint8X8 {
	return Uint8X8(C.vhsub_u8(C.uint8x8_t(v0), C.uint8x8_t(v1)))
}

func VhsubU32(v0 Uint32X2, v1 Uint32X2) Uint32X2 {
	return Uint32X2(C.vhsub_u32(C.uint32x2_t(v0), C.uint32x2_t(v1)))
}

func VhsubU16(v0 Uint16X4, v1 Uint16X4) Uint16X4 {
	return Uint16X4(C.vhsub_u16(C.uint16x4_t(v0), C.uint16x4_t(v1)))
}

func VhsubS8(v0 Int8X8, v1 Int8X8) Int8X8 {
	return Int8X8(C.vhsub_s8(C.int8x8_t(v0), C.int8x8_t(v1)))
}

func VhsubS32(v0 Int32X2, v1 Int32X2) Int32X2 {
	return Int32X2(C.vhsub_s32(C.int32x2_t(v0), C.int32x2_t(v1)))
}

func VhsubS16(v0 Int16X4, v1 Int16X4) Int16X4 {
	return Int16X4(C.vhsub_s16(C.int16x4_t(v0), C.int16x4_t(v1)))
}

func VmaxqU8(v0 Uint8X16, v1 Uint8X16) Uint8X16 {
	return Uint8X16(C.vmaxq_u8(C.uint8x16_t(v0), C.uint8x16_t(v1)))
}

func VmaxqU32(v0 Uint32X4, v1 Uint32X4) Uint32X4 {
	return Uint32X4(C.vmaxq_u32(C.uint32x4_t(v0), C.uint32x4_t(v1)))
}

func VmaxqU16(v0 Uint16X8, v1 Uint16X8) Uint16X8 {
	return Uint16X8(C.vmaxq_u16(C.uint16x8_t(v0), C.uint16x8_t(v1)))
}

func VmaxqS8(v0 Int8X16, v1 Int8X16) Int8X16 {
	return Int8X16(C.vmaxq_s8(C.int8x16_t(v0), C.int8x16_t(v1)))
}

func VmaxqF32(v0 Float32X4, v1 Float32X4) Float32X4 {
	return Float32X4(C.vmaxq_f32(C.float32x4_t(v0), C.float32x4_t(v1)))
}

func VmaxqS32(v0 Int32X4, v1 Int32X4) Int32X4 {
	return Int32X4(C.vmaxq_s32(C.int32x4_t(v0), C.int32x4_t(v1)))
}

func VmaxqS16(v0 Int16X8, v1 Int16X8) Int16X8 {
	return Int16X8(C.vmaxq_s16(C.int16x8_t(v0), C.int16x8_t(v1)))
}

func VmaxU8(v0 Uint8X8, v1 Uint8X8) Uint8X8 {
	return Uint8X8(C.vmax_u8(C.uint8x8_t(v0), C.uint8x8_t(v1)))
}

func VmaxU32(v0 Uint32X2, v1 Uint32X2) Uint32X2 {
	return Uint32X2(C.vmax_u32(C.uint32x2_t(v0), C.uint32x2_t(v1)))
}

func VmaxU16(v0 Uint16X4, v1 Uint16X4) Uint16X4 {
	return Uint16X4(C.vmax_u16(C.uint16x4_t(v0), C.uint16x4_t(v1)))
}

func VmaxS8(v0 Int8X8, v1 Int8X8) Int8X8 {
	return Int8X8(C.vmax_s8(C.int8x8_t(v0), C.int8x8_t(v1)))
}

func VmaxF32(v0 Float32X2, v1 Float32X2) Float32X2 {
	return Float32X2(C.vmax_f32(C.float32x2_t(v0), C.float32x2_t(v1)))
}

func VmaxS32(v0 Int32X2, v1 Int32X2) Int32X2 {
	return Int32X2(C.vmax_s32(C.int32x2_t(v0), C.int32x2_t(v1)))
}

func VmaxS16(v0 Int16X4, v1 Int16X4) Int16X4 {
	return Int16X4(C.vmax_s16(C.int16x4_t(v0), C.int16x4_t(v1)))
}

func VminqU8(v0 Uint8X16, v1 Uint8X16) Uint8X16 {
	return Uint8X16(C.vminq_u8(C.uint8x16_t(v0), C.uint8x16_t(v1)))
}

func VminqU32(v0 Uint32X4, v1 Uint32X4) Uint32X4 {
	return Uint32X4(C.vminq_u32(C.uint32x4_t(v0), C.uint32x4_t(v1)))
}

func VminqU16(v0 Uint16X8, v1 Uint16X8) Uint16X8 {
	return Uint16X8(C.vminq_u16(C.uint16x8_t(v0), C.uint16x8_t(v1)))
}

func VminqS8(v0 Int8X16, v1 Int8X16) Int8X16 {
	return Int8X16(C.vminq_s8(C.int8x16_t(v0), C.int8x16_t(v1)))
}

func VminqF32(v0 Float32X4, v1 Float32X4) Float32X4 {
	return Float32X4(C.vminq_f32(C.float32x4_t(v0), C.float32x4_t(v1)))
}

func VminqS32(v0 Int32X4, v1 Int32X4) Int32X4 {
	return Int32X4(C.vminq_s32(C.int32x4_t(v0), C.int32x4_t(v1)))
}

func VminqS16(v0 Int16X8, v1 Int16X8) Int16X8 {
	return Int16X8(C.vminq_s16(C.int16x8_t(v0), C.int16x8_t(v1)))
}

func VminU8(v0 Uint8X8, v1 Uint8X8) Uint8X8 {
	return Uint8X8(C.vmin_u8(C.uint8x8_t(v0), C.uint8x8_t(v1)))
}

func VminU32(v0 Uint32X2, v1 Uint32X2) Uint32X2 {
	return Uint32X2(C.vmin_u32(C.uint32x2_t(v0), C.uint32x2_t(v1)))
}

func VminU16(v0 Uint16X4, v1 Uint16X4) Uint16X4 {
	return Uint16X4(C.vmin_u16(C.uint16x4_t(v0), C.uint16x4_t(v1)))
}

func VminS8(v0 Int8X8, v1 Int8X8) Int8X8 {
	return Int8X8(C.vmin_s8(C.int8x8_t(v0), C.int8x8_t(v1)))
}

func VminF32(v0 Float32X2, v1 Float32X2) Float32X2 {
	return Float32X2(C.vmin_f32(C.float32x2_t(v0), C.float32x2_t(v1)))
}

func VminS32(v0 Int32X2, v1 Int32X2) Int32X2 {
	return Int32X2(C.vmin_s32(C.int32x2_t(v0), C.int32x2_t(v1)))
}

func VminS16(v0 Int16X4, v1 Int16X4) Int16X4 {
	return Int16X4(C.vmin_s16(C.int16x4_t(v0), C.int16x4_t(v1)))
}

func VmlaqU8(v0 Uint8X16, v1 Uint8X16, v2 Uint8X16) Uint8X16 {
	return Uint8X16(C.vmlaq_u8(C.uint8x16_t(v0), C.uint8x16_t(v1), C.uint8x16_t(v2)))
}

func VmlaqU32(v0 Uint32X4, v1 Uint32X4, v2 Uint32X4) Uint32X4 {
	return Uint32X4(C.vmlaq_u32(C.uint32x4_t(v0), C.uint32x4_t(v1), C.uint32x4_t(v2)))
}

func VmlaqU16(v0 Uint16X8, v1 Uint16X8, v2 Uint16X8) Uint16X8 {
	return Uint16X8(C.vmlaq_u16(C.uint16x8_t(v0), C.uint16x8_t(v1), C.uint16x8_t(v2)))
}

func VmlaqS8(v0 Int8X16, v1 Int8X16, v2 Int8X16) Int8X16 {
	return Int8X16(C.vmlaq_s8(C.int8x16_t(v0), C.int8x16_t(v1), C.int8x16_t(v2)))
}

func VmlaqF32(v0 Float32X4, v1 Float32X4, v2 Float32X4) Float32X4 {
	return Float32X4(C.vmlaq_f32(C.float32x4_t(v0), C.float32x4_t(v1), C.float32x4_t(v2)))
}

func VmlaqS32(v0 Int32X4, v1 Int32X4, v2 Int32X4) Int32X4 {
	return Int32X4(C.vmlaq_s32(C.int32x4_t(v0), C.int32x4_t(v1), C.int32x4_t(v2)))
}

func VmlaqS16(v0 Int16X8, v1 Int16X8, v2 Int16X8) Int16X8 {
	return Int16X8(C.vmlaq_s16(C.int16x8_t(v0), C.int16x8_t(v1), C.int16x8_t(v2)))
}

func VmlaU8(v0 Uint8X8, v1 Uint8X8, v2 Uint8X8) Uint8X8 {
	return Uint8X8(C.vmla_u8(C.uint8x8_t(v0), C.uint8x8_t(v1), C.uint8x8_t(v2)))
}

func VmlaU32(v0 Uint32X2, v1 Uint32X2, v2 Uint32X2) Uint32X2 {
	return Uint32X2(C.vmla_u32(C.uint32x2_t(v0), C.uint32x2_t(v1), C.uint32x2_t(v2)))
}

func VmlaU16(v0 Uint16X4, v1 Uint16X4, v2 Uint16X4) Uint16X4 {
	return Uint16X4(C.vmla_u16(C.uint16x4_t(v0), C.uint16x4_t(v1), C.uint16x4_t(v2)))
}

func VmlaS8(v0 Int8X8, v1 Int8X8, v2 Int8X8) Int8X8 {
	return Int8X8(C.vmla_s8(C.int8x8_t(v0), C.int8x8_t(v1), C.int8x8_t(v2)))
}

func VmlaF32(v0 Float32X2, v1 Float32X2, v2 Float32X2) Float32X2 {
	return Float32X2(C.vmla_f32(C.float32x2_t(v0), C.float32x2_t(v1), C.float32x2_t(v2)))
}

func VmlaS32(v0 Int32X2, v1 Int32X2, v2 Int32X2) Int32X2 {
	return Int32X2(C.vmla_s32(C.int32x2_t(v0), C.int32x2_t(v1), C.int32x2_t(v2)))
}

func VmlaS16(v0 Int16X4, v1 Int16X4, v2 Int16X4) Int16X4 {
	return Int16X4(C.vmla_s16(C.int16x4_t(v0), C.int16x4_t(v1), C.int16x4_t(v2)))
}

func VmlaqNU32(v0 Uint32X4, v1 Uint32X4, v2 uint32) Uint32X4 {
	return Uint32X4(C.vmlaq_n_u32(C.uint32x4_t(v0), C.uint32x4_t(v1), C.uint32_t(v2)))
}

func VmlaqNU16(v0 Uint16X8, v1 Uint16X8, v2 uint16) Uint16X8 {
	return Uint16X8(C.vmlaq_n_u16(C.uint16x8_t(v0), C.uint16x8_t(v1), C.uint16_t(v2)))
}

func VmlaqNF32(v0 Float32X4, v1 Float32X4, v2 float32) Float32X4 {
	return Float32X4(C.vmlaq_n_f32(C.float32x4_t(v0), C.float32x4_t(v1), C.float32_t(v2)))
}

func VmlaqNS32(v0 Int32X4, v1 Int32X4, v2 int32) Int32X4 {
	return Int32X4(C.vmlaq_n_s32(C.int32x4_t(v0), C.int32x4_t(v1), C.int32_t(v2)))
}

func VmlaqNS16(v0 Int16X8, v1 Int16X8, v2 int16) Int16X8 {
	return Int16X8(C.vmlaq_n_s16(C.int16x8_t(v0), C.int16x8_t(v1), C.int16_t(v2)))
}

func VmlaNU32(v0 Uint32X2, v1 Uint32X2, v2 uint32) Uint32X2 {
	return Uint32X2(C.vmla_n_u32(C.uint32x2_t(v0), C.uint32x2_t(v1), C.uint32_t(v2)))
}

func VmlaNU16(v0 Uint16X4, v1 Uint16X4, v2 uint16) Uint16X4 {
	return Uint16X4(C.vmla_n_u16(C.uint16x4_t(v0), C.uint16x4_t(v1), C.uint16_t(v2)))
}

func VmlaNF32(v0 Float32X2, v1 Float32X2, v2 float32) Float32X2 {
	return Float32X2(C.vmla_n_f32(C.float32x2_t(v0), C.float32x2_t(v1), C.float32_t(v2)))
}

func VmlaNS32(v0 Int32X2, v1 Int32X2, v2 int32) Int32X2 {
	return Int32X2(C.vmla_n_s32(C.int32x2_t(v0), C.int32x2_t(v1), C.int32_t(v2)))
}

func VmlaNS16(v0 Int16X4, v1 Int16X4, v2 int16) Int16X4 {
	return Int16X4(C.vmla_n_s16(C.int16x4_t(v0), C.int16x4_t(v1), C.int16_t(v2)))
}

func VmlsqU8(v0 Uint8X16, v1 Uint8X16, v2 Uint8X16) Uint8X16 {
	return Uint8X16(C.vmlsq_u8(C.uint8x16_t(v0), C.uint8x16_t(v1), C.uint8x16_t(v2)))
}

func VmlsqU32(v0 Uint32X4, v1 Uint32X4, v2 Uint32X4) Uint32X4 {
	return Uint32X4(C.vmlsq_u32(C.uint32x4_t(v0), C.uint32x4_t(v1), C.uint32x4_t(v2)))
}

func VmlsqU16(v0 Uint16X8, v1 Uint16X8, v2 Uint16X8) Uint16X8 {
	return Uint16X8(C.vmlsq_u16(C.uint16x8_t(v0), C.uint16x8_t(v1), C.uint16x8_t(v2)))
}

func VmlsqS8(v0 Int8X16, v1 Int8X16, v2 Int8X16) Int8X16 {
	return Int8X16(C.vmlsq_s8(C.int8x16_t(v0), C.int8x16_t(v1), C.int8x16_t(v2)))
}

func VmlsqF32(v0 Float32X4, v1 Float32X4, v2 Float32X4) Float32X4 {
	return Float32X4(C.vmlsq_f32(C.float32x4_t(v0), C.float32x4_t(v1), C.float32x4_t(v2)))
}

func VmlsqS32(v0 Int32X4, v1 Int32X4, v2 Int32X4) Int32X4 {
	return Int32X4(C.vmlsq_s32(C.int32x4_t(v0), C.int32x4_t(v1), C.int32x4_t(v2)))
}

func VmlsqS16(v0 Int16X8, v1 Int16X8, v2 Int16X8) Int16X8 {
	return Int16X8(C.vmlsq_s16(C.int16x8_t(v0), C.int16x8_t(v1), C.int16x8_t(v2)))
}

func VmlsU8(v0 Uint8X8, v1 Uint8X8, v2 Uint8X8) Uint8X8 {
	return Uint8X8(C.vmls_u8(C.uint8x8_t(v0), C.uint8x8_t(v1), C.uint8x8_t(v2)))
}

func VmlsU32(v0 Uint32X2, v1 Uint32X2, v2 Uint32X2) Uint32X2 {
	return Uint32X2(C.vmls_u32(C.uint32x2_t(v0), C.uint32x2_t(v1), C.uint32x2_t(v2)))
}

func VmlsU16(v0 Uint16X4, v1 Uint16X4, v2 Uint16X4) Uint16X4 {
	return Uint16X4(C.vmls_u16(C.uint16x4_t(v0), C.uint16x4_t(v1), C.uint16x4_t(v2)))
}

func VmlsS8(v0 Int8X8, v1 Int8X8, v2 Int8X8) Int8X8 {
	return Int8X8(C.vmls_s8(C.int8x8_t(v0), C.int8x8_t(v1), C.int8x8_t(v2)))
}

func VmlsF32(v0 Float32X2, v1 Float32X2, v2 Float32X2) Float32X2 {
	return Float32X2(C.vmls_f32(C.float32x2_t(v0), C.float32x2_t(v1), C.float32x2_t(v2)))
}

func VmlsS32(v0 Int32X2, v1 Int32X2, v2 Int32X2) Int32X2 {
	return Int32X2(C.vmls_s32(C.int32x2_t(v0), C.int32x2_t(v1), C.int32x2_t(v2)))
}

func VmlsS16(v0 Int16X4, v1 Int16X4, v2 Int16X4) Int16X4 {
	return Int16X4(C.vmls_s16(C.int16x4_t(v0), C.int16x4_t(v1), C.int16x4_t(v2)))
}

func VmlsqNU32(v0 Uint32X4, v1 Uint32X4, v2 uint32) Uint32X4 {
	return Uint32X4(C.vmlsq_n_u32(C.uint32x4_t(v0), C.uint32x4_t(v1), C.uint32_t(v2)))
}

func VmlsqNU16(v0 Uint16X8, v1 Uint16X8, v2 uint16) Uint16X8 {
	return Uint16X8(C.vmlsq_n_u16(C.uint16x8_t(v0), C.uint16x8_t(v1), C.uint16_t(v2)))
}

func VmlsqNF32(v0 Float32X4, v1 Float32X4, v2 float32) Float32X4 {
	return Float32X4(C.vmlsq_n_f32(C.float32x4_t(v0), C.float32x4_t(v1), C.float32_t(v2)))
}

func VmlsqNS32(v0 Int32X4, v1 Int32X4, v2 int32) Int32X4 {
	return Int32X4(C.vmlsq_n_s32(C.int32x4_t(v0), C.int32x4_t(v1), C.int32_t(v2)))
}

func VmlsqNS16(v0 Int16X8, v1 Int16X8, v2 int16) Int16X8 {
	return Int16X8(C.vmlsq_n_s16(C.int16x8_t(v0), C.int16x8_t(v1), C.int16_t(v2)))
}

func VmlsNU32(v0 Uint32X2, v1 Uint32X2, v2 uint32) Uint32X2 {
	return Uint32X2(C.vmls_n_u32(C.uint32x2_t(v0), C.uint32x2_t(v1), C.uint32_t(v2)))
}

func VmlsNU16(v0 Uint16X4, v1 Uint16X4, v2 uint16) Uint16X4 {
	return Uint16X4(C.vmls_n_u16(C.uint16x4_t(v0), C.uint16x4_t(v1), C.uint16_t(v2)))
}

func VmlsNF32(v0 Float32X2, v1 Float32X2, v2 float32) Float32X2 {
	return Float32X2(C.vmls_n_f32(C.float32x2_t(v0), C.float32x2_t(v1), C.float32_t(v2)))
}

func VmlsNS32(v0 Int32X2, v1 Int32X2, v2 int32) Int32X2 {
	return Int32X2(C.vmls_n_s32(C.int32x2_t(v0), C.int32x2_t(v1), C.int32_t(v2)))
}

func VmlsNS16(v0 Int16X4, v1 Int16X4, v2 int16) Int16X4 {
	return Int16X4(C.vmls_n_s16(C.int16x4_t(v0), C.int16x4_t(v1), C.int16_t(v2)))
}

func VmovNP8(v0 Poly8) Poly8X8 {
	return Poly8X8(C.vmov_n_p8(C.poly8_t(v0)))
}

func VmovNP16(v0 Poly16) Poly16X4 {
	return Poly16X4(C.vmov_n_p16(C.poly16_t(v0)))
}

func VmovqNP8(v0 Poly8) Poly8X16 {
	return Poly8X16(C.vmovq_n_p8(C.poly8_t(v0)))
}

func VmovqNP16(v0 Poly16) Poly16X8 {
	return Poly16X8(C.vmovq_n_p16(C.poly16_t(v0)))
}

func VmovqNU8(v0 uint8) Uint8X16 {
	return Uint8X16(C.vmovq_n_u8(C.uint8_t(v0)))
}

func VmovqNU32(v0 uint32) Uint32X4 {
	return Uint32X4(C.vmovq_n_u32(C.uint32_t(v0)))
}

func VmovqNU64(v0 uint64) Uint64X2 {
	return Uint64X2(C.vmovq_n_u64(C.uint64_t(v0)))
}

func VmovqNU16(v0 uint16) Uint16X8 {
	return Uint16X8(C.vmovq_n_u16(C.uint16_t(v0)))
}

func VmovqNS8(v0 int8) Int8X16 {
	return Int8X16(C.vmovq_n_s8(C.int8_t(v0)))
}

func VmovqNF32(v0 float32) Float32X4 {
	return Float32X4(C.vmovq_n_f32(C.float32_t(v0)))
}

func VmovqNS32(v0 int32) Int32X4 {
	return Int32X4(C.vmovq_n_s32(C.int32_t(v0)))
}

func VmovqNS64(v0 int64) Int64X2 {
	return Int64X2(C.vmovq_n_s64(C.int64_t(v0)))
}

func VmovqNS16(v0 int16) Int16X8 {
	return Int16X8(C.vmovq_n_s16(C.int16_t(v0)))
}

func VmovNU8(v0 uint8) Uint8X8 {
	return Uint8X8(C.vmov_n_u8(C.uint8_t(v0)))
}

func VmovNU32(v0 uint32) Uint32X2 {
	return Uint32X2(C.vmov_n_u32(C.uint32_t(v0)))
}

func VmovNU64(v0 uint64) Uint64X1 {
	return Uint64X1(C.vmov_n_u64(C.uint64_t(v0)))
}

func VmovNU16(v0 uint16) Uint16X4 {
	return Uint16X4(C.vmov_n_u16(C.uint16_t(v0)))
}

func VmovNS8(v0 int8) Int8X8 {
	return Int8X8(C.vmov_n_s8(C.int8_t(v0)))
}

func VmovNF32(v0 float32) Float32X2 {
	return Float32X2(C.vmov_n_f32(C.float32_t(v0)))
}

func VmovNS32(v0 int32) Int32X2 {
	return Int32X2(C.vmov_n_s32(C.int32_t(v0)))
}

func VmovNS64(v0 int64) Int64X1 {
	return Int64X1(C.vmov_n_s64(C.int64_t(v0)))
}

func VmovNS16(v0 int16) Int16X4 {
	return Int16X4(C.vmov_n_s16(C.int16_t(v0)))
}

func VmovlU8(v0 Uint8X8) Uint16X8 {
	return Uint16X8(C.vmovl_u8(C.uint8x8_t(v0)))
}

func VmovlU32(v0 Uint32X2) Uint64X2 {
	return Uint64X2(C.vmovl_u32(C.uint32x2_t(v0)))
}

func VmovlU16(v0 Uint16X4) Uint32X4 {
	return Uint32X4(C.vmovl_u16(C.uint16x4_t(v0)))
}

func VmovlS8(v0 Int8X8) Int16X8 {
	return Int16X8(C.vmovl_s8(C.int8x8_t(v0)))
}

func VmovlS32(v0 Int32X2) Int64X2 {
	return Int64X2(C.vmovl_s32(C.int32x2_t(v0)))
}

func VmovlS16(v0 Int16X4) Int32X4 {
	return Int32X4(C.vmovl_s16(C.int16x4_t(v0)))
}

func VmovnU32(v0 Uint32X4) Uint16X4 {
	return Uint16X4(C.vmovn_u32(C.uint32x4_t(v0)))
}

func VmovnU64(v0 Uint64X2) Uint32X2 {
	return Uint32X2(C.vmovn_u64(C.uint64x2_t(v0)))
}

func VmovnU16(v0 Uint16X8) Uint8X8 {
	return Uint8X8(C.vmovn_u16(C.uint16x8_t(v0)))
}

func VmovnS32(v0 Int32X4) Int16X4 {
	return Int16X4(C.vmovn_s32(C.int32x4_t(v0)))
}

func VmovnS64(v0 Int64X2) Int32X2 {
	return Int32X2(C.vmovn_s64(C.int64x2_t(v0)))
}

func VmovnS16(v0 Int16X8) Int8X8 {
	return Int8X8(C.vmovn_s16(C.int16x8_t(v0)))
}

func VmulqU8(v0 Uint8X16, v1 Uint8X16) Uint8X16 {
	return Uint8X16(C.vmulq_u8(C.uint8x16_t(v0), C.uint8x16_t(v1)))
}

func VmulqU32(v0 Uint32X4, v1 Uint32X4) Uint32X4 {
	return Uint32X4(C.vmulq_u32(C.uint32x4_t(v0), C.uint32x4_t(v1)))
}

func VmulqU16(v0 Uint16X8, v1 Uint16X8) Uint16X8 {
	return Uint16X8(C.vmulq_u16(C.uint16x8_t(v0), C.uint16x8_t(v1)))
}

func VmulqS8(v0 Int8X16, v1 Int8X16) Int8X16 {
	return Int8X16(C.vmulq_s8(C.int8x16_t(v0), C.int8x16_t(v1)))
}

func VmulqF32(v0 Float32X4, v1 Float32X4) Float32X4 {
	return Float32X4(C.vmulq_f32(C.float32x4_t(v0), C.float32x4_t(v1)))
}

func VmulqS32(v0 Int32X4, v1 Int32X4) Int32X4 {
	return Int32X4(C.vmulq_s32(C.int32x4_t(v0), C.int32x4_t(v1)))
}

func VmulqS16(v0 Int16X8, v1 Int16X8) Int16X8 {
	return Int16X8(C.vmulq_s16(C.int16x8_t(v0), C.int16x8_t(v1)))
}

func VmulU8(v0 Uint8X8, v1 Uint8X8) Uint8X8 {
	return Uint8X8(C.vmul_u8(C.uint8x8_t(v0), C.uint8x8_t(v1)))
}

func VmulU32(v0 Uint32X2, v1 Uint32X2) Uint32X2 {
	return Uint32X2(C.vmul_u32(C.uint32x2_t(v0), C.uint32x2_t(v1)))
}

func VmulU16(v0 Uint16X4, v1 Uint16X4) Uint16X4 {
	return Uint16X4(C.vmul_u16(C.uint16x4_t(v0), C.uint16x4_t(v1)))
}

func VmulS8(v0 Int8X8, v1 Int8X8) Int8X8 {
	return Int8X8(C.vmul_s8(C.int8x8_t(v0), C.int8x8_t(v1)))
}

func VmulF32(v0 Float32X2, v1 Float32X2) Float32X2 {
	return Float32X2(C.vmul_f32(C.float32x2_t(v0), C.float32x2_t(v1)))
}

func VmulS32(v0 Int32X2, v1 Int32X2) Int32X2 {
	return Int32X2(C.vmul_s32(C.int32x2_t(v0), C.int32x2_t(v1)))
}

func VmulS16(v0 Int16X4, v1 Int16X4) Int16X4 {
	return Int16X4(C.vmul_s16(C.int16x4_t(v0), C.int16x4_t(v1)))
}

func VmulP8(v0 Poly8X8, v1 Poly8X8) Poly8X8 {
	return Poly8X8(C.vmul_p8(C.poly8x8_t(v0), C.poly8x8_t(v1)))
}

func VmulqP8(v0 Poly8X16, v1 Poly8X16) Poly8X16 {
	return Poly8X16(C.vmulq_p8(C.poly8x16_t(v0), C.poly8x16_t(v1)))
}

func VmulqNU32(v0 Uint32X4, v1 uint32) Uint32X4 {
	return Uint32X4(C.vmulq_n_u32(C.uint32x4_t(v0), C.uint32_t(v1)))
}

func VmulqNU16(v0 Uint16X8, v1 uint16) Uint16X8 {
	return Uint16X8(C.vmulq_n_u16(C.uint16x8_t(v0), C.uint16_t(v1)))
}

func VmulqNF32(v0 Float32X4, v1 float32) Float32X4 {
	return Float32X4(C.vmulq_n_f32(C.float32x4_t(v0), C.float32_t(v1)))
}

func VmulqNS32(v0 Int32X4, v1 int32) Int32X4 {
	return Int32X4(C.vmulq_n_s32(C.int32x4_t(v0), C.int32_t(v1)))
}

func VmulqNS16(v0 Int16X8, v1 int16) Int16X8 {
	return Int16X8(C.vmulq_n_s16(C.int16x8_t(v0), C.int16_t(v1)))
}

func VmulNU32(v0 Uint32X2, v1 uint32) Uint32X2 {
	return Uint32X2(C.vmul_n_u32(C.uint32x2_t(v0), C.uint32_t(v1)))
}

func VmulNU16(v0 Uint16X4, v1 uint16) Uint16X4 {
	return Uint16X4(C.vmul_n_u16(C.uint16x4_t(v0), C.uint16_t(v1)))
}

func VmulNF32(v0 Float32X2, v1 float32) Float32X2 {
	return Float32X2(C.vmul_n_f32(C.float32x2_t(v0), C.float32_t(v1)))
}

func VmulNS32(v0 Int32X2, v1 int32) Int32X2 {
	return Int32X2(C.vmul_n_s32(C.int32x2_t(v0), C.int32_t(v1)))
}

func VmulNS16(v0 Int16X4, v1 int16) Int16X4 {
	return Int16X4(C.vmul_n_s16(C.int16x4_t(v0), C.int16_t(v1)))
}

func VmullP8(v0 Poly8X8, v1 Poly8X8) Poly16X8 {
	return Poly16X8(C.vmull_p8(C.poly8x8_t(v0), C.poly8x8_t(v1)))
}

func VmullU8(v0 Uint8X8, v1 Uint8X8) Uint16X8 {
	return Uint16X8(C.vmull_u8(C.uint8x8_t(v0), C.uint8x8_t(v1)))
}

func VmullU32(v0 Uint32X2, v1 Uint32X2) Uint64X2 {
	return Uint64X2(C.vmull_u32(C.uint32x2_t(v0), C.uint32x2_t(v1)))
}

func VmullU16(v0 Uint16X4, v1 Uint16X4) Uint32X4 {
	return Uint32X4(C.vmull_u16(C.uint16x4_t(v0), C.uint16x4_t(v1)))
}

func VmullS8(v0 Int8X8, v1 Int8X8) Int16X8 {
	return Int16X8(C.vmull_s8(C.int8x8_t(v0), C.int8x8_t(v1)))
}

func VmullS32(v0 Int32X2, v1 Int32X2) Int64X2 {
	return Int64X2(C.vmull_s32(C.int32x2_t(v0), C.int32x2_t(v1)))
}

func VmullS16(v0 Int16X4, v1 Int16X4) Int32X4 {
	return Int32X4(C.vmull_s16(C.int16x4_t(v0), C.int16x4_t(v1)))
}

func VmullNU32(v0 Uint32X2, v1 uint32) Uint64X2 {
	return Uint64X2(C.vmull_n_u32(C.uint32x2_t(v0), C.uint32_t(v1)))
}

func VmullNU16(v0 Uint16X4, v1 uint16) Uint32X4 {
	return Uint32X4(C.vmull_n_u16(C.uint16x4_t(v0), C.uint16_t(v1)))
}

func VmullNS32(v0 Int32X2, v1 int32) Int64X2 {
	return Int64X2(C.vmull_n_s32(C.int32x2_t(v0), C.int32_t(v1)))
}

func VmullNS16(v0 Int16X4, v1 int16) Int32X4 {
	return Int32X4(C.vmull_n_s16(C.int16x4_t(v0), C.int16_t(v1)))
}

func VmvnP8(v0 Poly8X8) Poly8X8 {
	return Poly8X8(C.vmvn_p8(C.poly8x8_t(v0)))
}

func VmvnqP8(v0 Poly8X16) Poly8X16 {
	return Poly8X16(C.vmvnq_p8(C.poly8x16_t(v0)))
}

func VmvnqU8(v0 Uint8X16) Uint8X16 {
	return Uint8X16(C.vmvnq_u8(C.uint8x16_t(v0)))
}

func VmvnqU32(v0 Uint32X4) Uint32X4 {
	return Uint32X4(C.vmvnq_u32(C.uint32x4_t(v0)))
}

func VmvnqU16(v0 Uint16X8) Uint16X8 {
	return Uint16X8(C.vmvnq_u16(C.uint16x8_t(v0)))
}

func VmvnqS8(v0 Int8X16) Int8X16 {
	return Int8X16(C.vmvnq_s8(C.int8x16_t(v0)))
}

func VmvnqS32(v0 Int32X4) Int32X4 {
	return Int32X4(C.vmvnq_s32(C.int32x4_t(v0)))
}

func VmvnqS16(v0 Int16X8) Int16X8 {
	return Int16X8(C.vmvnq_s16(C.int16x8_t(v0)))
}

func VmvnU8(v0 Uint8X8) Uint8X8 {
	return Uint8X8(C.vmvn_u8(C.uint8x8_t(v0)))
}

func VmvnU32(v0 Uint32X2) Uint32X2 {
	return Uint32X2(C.vmvn_u32(C.uint32x2_t(v0)))
}

func VmvnU16(v0 Uint16X4) Uint16X4 {
	return Uint16X4(C.vmvn_u16(C.uint16x4_t(v0)))
}

func VmvnS8(v0 Int8X8) Int8X8 {
	return Int8X8(C.vmvn_s8(C.int8x8_t(v0)))
}

func VmvnS32(v0 Int32X2) Int32X2 {
	return Int32X2(C.vmvn_s32(C.int32x2_t(v0)))
}

func VmvnS16(v0 Int16X4) Int16X4 {
	return Int16X4(C.vmvn_s16(C.int16x4_t(v0)))
}

func VnegqS8(v0 Int8X16) Int8X16 {
	return Int8X16(C.vnegq_s8(C.int8x16_t(v0)))
}

func VnegqF32(v0 Float32X4) Float32X4 {
	return Float32X4(C.vnegq_f32(C.float32x4_t(v0)))
}

func VnegqS32(v0 Int32X4) Int32X4 {
	return Int32X4(C.vnegq_s32(C.int32x4_t(v0)))
}

func VnegqS16(v0 Int16X8) Int16X8 {
	return Int16X8(C.vnegq_s16(C.int16x8_t(v0)))
}

func VnegS8(v0 Int8X8) Int8X8 {
	return Int8X8(C.vneg_s8(C.int8x8_t(v0)))
}

func VnegF32(v0 Float32X2) Float32X2 {
	return Float32X2(C.vneg_f32(C.float32x2_t(v0)))
}

func VnegS32(v0 Int32X2) Int32X2 {
	return Int32X2(C.vneg_s32(C.int32x2_t(v0)))
}

func VnegS16(v0 Int16X4) Int16X4 {
	return Int16X4(C.vneg_s16(C.int16x4_t(v0)))
}

func VornqU8(v0 Uint8X16, v1 Uint8X16) Uint8X16 {
	return Uint8X16(C.vornq_u8(C.uint8x16_t(v0), C.uint8x16_t(v1)))
}

func VornqU32(v0 Uint32X4, v1 Uint32X4) Uint32X4 {
	return Uint32X4(C.vornq_u32(C.uint32x4_t(v0), C.uint32x4_t(v1)))
}

func VornqU64(v0 Uint64X2, v1 Uint64X2) Uint64X2 {
	return Uint64X2(C.vornq_u64(C.uint64x2_t(v0), C.uint64x2_t(v1)))
}

func VornqU16(v0 Uint16X8, v1 Uint16X8) Uint16X8 {
	return Uint16X8(C.vornq_u16(C.uint16x8_t(v0), C.uint16x8_t(v1)))
}

func VornqS8(v0 Int8X16, v1 Int8X16) Int8X16 {
	return Int8X16(C.vornq_s8(C.int8x16_t(v0), C.int8x16_t(v1)))
}

func VornqS32(v0 Int32X4, v1 Int32X4) Int32X4 {
	return Int32X4(C.vornq_s32(C.int32x4_t(v0), C.int32x4_t(v1)))
}

func VornqS64(v0 Int64X2, v1 Int64X2) Int64X2 {
	return Int64X2(C.vornq_s64(C.int64x2_t(v0), C.int64x2_t(v1)))
}

func VornqS16(v0 Int16X8, v1 Int16X8) Int16X8 {
	return Int16X8(C.vornq_s16(C.int16x8_t(v0), C.int16x8_t(v1)))
}

func VornU8(v0 Uint8X8, v1 Uint8X8) Uint8X8 {
	return Uint8X8(C.vorn_u8(C.uint8x8_t(v0), C.uint8x8_t(v1)))
}

func VornU32(v0 Uint32X2, v1 Uint32X2) Uint32X2 {
	return Uint32X2(C.vorn_u32(C.uint32x2_t(v0), C.uint32x2_t(v1)))
}

func VornU64(v0 Uint64X1, v1 Uint64X1) Uint64X1 {
	return Uint64X1(C.vorn_u64(C.uint64x1_t(v0), C.uint64x1_t(v1)))
}

func VornU16(v0 Uint16X4, v1 Uint16X4) Uint16X4 {
	return Uint16X4(C.vorn_u16(C.uint16x4_t(v0), C.uint16x4_t(v1)))
}

func VornS8(v0 Int8X8, v1 Int8X8) Int8X8 {
	return Int8X8(C.vorn_s8(C.int8x8_t(v0), C.int8x8_t(v1)))
}

func VornS32(v0 Int32X2, v1 Int32X2) Int32X2 {
	return Int32X2(C.vorn_s32(C.int32x2_t(v0), C.int32x2_t(v1)))
}

func VornS64(v0 Int64X1, v1 Int64X1) Int64X1 {
	return Int64X1(C.vorn_s64(C.int64x1_t(v0), C.int64x1_t(v1)))
}

func VornS16(v0 Int16X4, v1 Int16X4) Int16X4 {
	return Int16X4(C.vorn_s16(C.int16x4_t(v0), C.int16x4_t(v1)))
}

func VorrqU8(v0 Uint8X16, v1 Uint8X16) Uint8X16 {
	return Uint8X16(C.vorrq_u8(C.uint8x16_t(v0), C.uint8x16_t(v1)))
}

func VorrqU32(v0 Uint32X4, v1 Uint32X4) Uint32X4 {
	return Uint32X4(C.vorrq_u32(C.uint32x4_t(v0), C.uint32x4_t(v1)))
}

func VorrqU64(v0 Uint64X2, v1 Uint64X2) Uint64X2 {
	return Uint64X2(C.vorrq_u64(C.uint64x2_t(v0), C.uint64x2_t(v1)))
}

func VorrqU16(v0 Uint16X8, v1 Uint16X8) Uint16X8 {
	return Uint16X8(C.vorrq_u16(C.uint16x8_t(v0), C.uint16x8_t(v1)))
}

func VorrqS8(v0 Int8X16, v1 Int8X16) Int8X16 {
	return Int8X16(C.vorrq_s8(C.int8x16_t(v0), C.int8x16_t(v1)))
}

func VorrqS32(v0 Int32X4, v1 Int32X4) Int32X4 {
	return Int32X4(C.vorrq_s32(C.int32x4_t(v0), C.int32x4_t(v1)))
}

func VorrqS64(v0 Int64X2, v1 Int64X2) Int64X2 {
	return Int64X2(C.vorrq_s64(C.int64x2_t(v0), C.int64x2_t(v1)))
}

func VorrqS16(v0 Int16X8, v1 Int16X8) Int16X8 {
	return Int16X8(C.vorrq_s16(C.int16x8_t(v0), C.int16x8_t(v1)))
}

func VorrU8(v0 Uint8X8, v1 Uint8X8) Uint8X8 {
	return Uint8X8(C.vorr_u8(C.uint8x8_t(v0), C.uint8x8_t(v1)))
}

func VorrU32(v0 Uint32X2, v1 Uint32X2) Uint32X2 {
	return Uint32X2(C.vorr_u32(C.uint32x2_t(v0), C.uint32x2_t(v1)))
}

func VorrU64(v0 Uint64X1, v1 Uint64X1) Uint64X1 {
	return Uint64X1(C.vorr_u64(C.uint64x1_t(v0), C.uint64x1_t(v1)))
}

func VorrU16(v0 Uint16X4, v1 Uint16X4) Uint16X4 {
	return Uint16X4(C.vorr_u16(C.uint16x4_t(v0), C.uint16x4_t(v1)))
}

func VorrS8(v0 Int8X8, v1 Int8X8) Int8X8 {
	return Int8X8(C.vorr_s8(C.int8x8_t(v0), C.int8x8_t(v1)))
}

func VorrS32(v0 Int32X2, v1 Int32X2) Int32X2 {
	return Int32X2(C.vorr_s32(C.int32x2_t(v0), C.int32x2_t(v1)))
}

func VorrS64(v0 Int64X1, v1 Int64X1) Int64X1 {
	return Int64X1(C.vorr_s64(C.int64x1_t(v0), C.int64x1_t(v1)))
}

func VorrS16(v0 Int16X4, v1 Int16X4) Int16X4 {
	return Int16X4(C.vorr_s16(C.int16x4_t(v0), C.int16x4_t(v1)))
}

func VpadalqU8(v0 Uint16X8, v1 Uint8X16) Uint16X8 {
	return Uint16X8(C.vpadalq_u8(C.uint16x8_t(v0), C.uint8x16_t(v1)))
}

func VpadalqU32(v0 Uint64X2, v1 Uint32X4) Uint64X2 {
	return Uint64X2(C.vpadalq_u32(C.uint64x2_t(v0), C.uint32x4_t(v1)))
}

func VpadalqU16(v0 Uint32X4, v1 Uint16X8) Uint32X4 {
	return Uint32X4(C.vpadalq_u16(C.uint32x4_t(v0), C.uint16x8_t(v1)))
}

func VpadalqS8(v0 Int16X8, v1 Int8X16) Int16X8 {
	return Int16X8(C.vpadalq_s8(C.int16x8_t(v0), C.int8x16_t(v1)))
}

func VpadalqS32(v0 Int64X2, v1 Int32X4) Int64X2 {
	return Int64X2(C.vpadalq_s32(C.int64x2_t(v0), C.int32x4_t(v1)))
}

func VpadalqS16(v0 Int32X4, v1 Int16X8) Int32X4 {
	return Int32X4(C.vpadalq_s16(C.int32x4_t(v0), C.int16x8_t(v1)))
}

func VpadalU8(v0 Uint16X4, v1 Uint8X8) Uint16X4 {
	return Uint16X4(C.vpadal_u8(C.uint16x4_t(v0), C.uint8x8_t(v1)))
}

func VpadalU32(v0 Uint64X1, v1 Uint32X2) Uint64X1 {
	return Uint64X1(C.vpadal_u32(C.uint64x1_t(v0), C.uint32x2_t(v1)))
}

func VpadalU16(v0 Uint32X2, v1 Uint16X4) Uint32X2 {
	return Uint32X2(C.vpadal_u16(C.uint32x2_t(v0), C.uint16x4_t(v1)))
}

func VpadalS8(v0 Int16X4, v1 Int8X8) Int16X4 {
	return Int16X4(C.vpadal_s8(C.int16x4_t(v0), C.int8x8_t(v1)))
}

func VpadalS32(v0 Int64X1, v1 Int32X2) Int64X1 {
	return Int64X1(C.vpadal_s32(C.int64x1_t(v0), C.int32x2_t(v1)))
}

func VpadalS16(v0 Int32X2, v1 Int16X4) Int32X2 {
	return Int32X2(C.vpadal_s16(C.int32x2_t(v0), C.int16x4_t(v1)))
}

func VpaddU8(v0 Uint8X8, v1 Uint8X8) Uint8X8 {
	return Uint8X8(C.vpadd_u8(C.uint8x8_t(v0), C.uint8x8_t(v1)))
}

func VpaddU32(v0 Uint32X2, v1 Uint32X2) Uint32X2 {
	return Uint32X2(C.vpadd_u32(C.uint32x2_t(v0), C.uint32x2_t(v1)))
}

func VpaddU16(v0 Uint16X4, v1 Uint16X4) Uint16X4 {
	return Uint16X4(C.vpadd_u16(C.uint16x4_t(v0), C.uint16x4_t(v1)))
}

func VpaddS8(v0 Int8X8, v1 Int8X8) Int8X8 {
	return Int8X8(C.vpadd_s8(C.int8x8_t(v0), C.int8x8_t(v1)))
}

func VpaddF32(v0 Float32X2, v1 Float32X2) Float32X2 {
	return Float32X2(C.vpadd_f32(C.float32x2_t(v0), C.float32x2_t(v1)))
}

func VpaddS32(v0 Int32X2, v1 Int32X2) Int32X2 {
	return Int32X2(C.vpadd_s32(C.int32x2_t(v0), C.int32x2_t(v1)))
}

func VpaddS16(v0 Int16X4, v1 Int16X4) Int16X4 {
	return Int16X4(C.vpadd_s16(C.int16x4_t(v0), C.int16x4_t(v1)))
}

func VpaddlqU8(v0 Uint8X16) Uint16X8 {
	return Uint16X8(C.vpaddlq_u8(C.uint8x16_t(v0)))
}

func VpaddlqU32(v0 Uint32X4) Uint64X2 {
	return Uint64X2(C.vpaddlq_u32(C.uint32x4_t(v0)))
}

func VpaddlqU16(v0 Uint16X8) Uint32X4 {
	return Uint32X4(C.vpaddlq_u16(C.uint16x8_t(v0)))
}

func VpaddlqS8(v0 Int8X16) Int16X8 {
	return Int16X8(C.vpaddlq_s8(C.int8x16_t(v0)))
}

func VpaddlqS32(v0 Int32X4) Int64X2 {
	return Int64X2(C.vpaddlq_s32(C.int32x4_t(v0)))
}

func VpaddlqS16(v0 Int16X8) Int32X4 {
	return Int32X4(C.vpaddlq_s16(C.int16x8_t(v0)))
}

func VpaddlU8(v0 Uint8X8) Uint16X4 {
	return Uint16X4(C.vpaddl_u8(C.uint8x8_t(v0)))
}

func VpaddlU32(v0 Uint32X2) Uint64X1 {
	return Uint64X1(C.vpaddl_u32(C.uint32x2_t(v0)))
}

func VpaddlU16(v0 Uint16X4) Uint32X2 {
	return Uint32X2(C.vpaddl_u16(C.uint16x4_t(v0)))
}

func VpaddlS8(v0 Int8X8) Int16X4 {
	return Int16X4(C.vpaddl_s8(C.int8x8_t(v0)))
}

func VpaddlS32(v0 Int32X2) Int64X1 {
	return Int64X1(C.vpaddl_s32(C.int32x2_t(v0)))
}

func VpaddlS16(v0 Int16X4) Int32X2 {
	return Int32X2(C.vpaddl_s16(C.int16x4_t(v0)))
}

func VpmaxU8(v0 Uint8X8, v1 Uint8X8) Uint8X8 {
	return Uint8X8(C.vpmax_u8(C.uint8x8_t(v0), C.uint8x8_t(v1)))
}

func VpmaxU32(v0 Uint32X2, v1 Uint32X2) Uint32X2 {
	return Uint32X2(C.vpmax_u32(C.uint32x2_t(v0), C.uint32x2_t(v1)))
}

func VpmaxU16(v0 Uint16X4, v1 Uint16X4) Uint16X4 {
	return Uint16X4(C.vpmax_u16(C.uint16x4_t(v0), C.uint16x4_t(v1)))
}

func VpmaxS8(v0 Int8X8, v1 Int8X8) Int8X8 {
	return Int8X8(C.vpmax_s8(C.int8x8_t(v0), C.int8x8_t(v1)))
}

func VpmaxF32(v0 Float32X2, v1 Float32X2) Float32X2 {
	return Float32X2(C.vpmax_f32(C.float32x2_t(v0), C.float32x2_t(v1)))
}

func VpmaxS32(v0 Int32X2, v1 Int32X2) Int32X2 {
	return Int32X2(C.vpmax_s32(C.int32x2_t(v0), C.int32x2_t(v1)))
}

func VpmaxS16(v0 Int16X4, v1 Int16X4) Int16X4 {
	return Int16X4(C.vpmax_s16(C.int16x4_t(v0), C.int16x4_t(v1)))
}

func VpminU8(v0 Uint8X8, v1 Uint8X8) Uint8X8 {
	return Uint8X8(C.vpmin_u8(C.uint8x8_t(v0), C.uint8x8_t(v1)))
}

func VpminU32(v0 Uint32X2, v1 Uint32X2) Uint32X2 {
	return Uint32X2(C.vpmin_u32(C.uint32x2_t(v0), C.uint32x2_t(v1)))
}

func VpminU16(v0 Uint16X4, v1 Uint16X4) Uint16X4 {
	return Uint16X4(C.vpmin_u16(C.uint16x4_t(v0), C.uint16x4_t(v1)))
}

func VpminS8(v0 Int8X8, v1 Int8X8) Int8X8 {
	return Int8X8(C.vpmin_s8(C.int8x8_t(v0), C.int8x8_t(v1)))
}

func VpminF32(v0 Float32X2, v1 Float32X2) Float32X2 {
	return Float32X2(C.vpmin_f32(C.float32x2_t(v0), C.float32x2_t(v1)))
}

func VpminS32(v0 Int32X2, v1 Int32X2) Int32X2 {
	return Int32X2(C.vpmin_s32(C.int32x2_t(v0), C.int32x2_t(v1)))
}

func VpminS16(v0 Int16X4, v1 Int16X4) Int16X4 {
	return Int16X4(C.vpmin_s16(C.int16x4_t(v0), C.int16x4_t(v1)))
}

func VqabsqS8(v0 Int8X16) Int8X16 {
	return Int8X16(C.vqabsq_s8(C.int8x16_t(v0)))
}

func VqabsqS32(v0 Int32X4) Int32X4 {
	return Int32X4(C.vqabsq_s32(C.int32x4_t(v0)))
}

func VqabsqS16(v0 Int16X8) Int16X8 {
	return Int16X8(C.vqabsq_s16(C.int16x8_t(v0)))
}

func VqabsS8(v0 Int8X8) Int8X8 {
	return Int8X8(C.vqabs_s8(C.int8x8_t(v0)))
}

func VqabsS32(v0 Int32X2) Int32X2 {
	return Int32X2(C.vqabs_s32(C.int32x2_t(v0)))
}

func VqabsS16(v0 Int16X4) Int16X4 {
	return Int16X4(C.vqabs_s16(C.int16x4_t(v0)))
}

func VqaddqU8(v0 Uint8X16, v1 Uint8X16) Uint8X16 {
	return Uint8X16(C.vqaddq_u8(C.uint8x16_t(v0), C.uint8x16_t(v1)))
}

func VqaddqU32(v0 Uint32X4, v1 Uint32X4) Uint32X4 {
	return Uint32X4(C.vqaddq_u32(C.uint32x4_t(v0), C.uint32x4_t(v1)))
}

func VqaddqU64(v0 Uint64X2, v1 Uint64X2) Uint64X2 {
	return Uint64X2(C.vqaddq_u64(C.uint64x2_t(v0), C.uint64x2_t(v1)))
}

func VqaddqU16(v0 Uint16X8, v1 Uint16X8) Uint16X8 {
	return Uint16X8(C.vqaddq_u16(C.uint16x8_t(v0), C.uint16x8_t(v1)))
}

func VqaddqS8(v0 Int8X16, v1 Int8X16) Int8X16 {
	return Int8X16(C.vqaddq_s8(C.int8x16_t(v0), C.int8x16_t(v1)))
}

func VqaddqS32(v0 Int32X4, v1 Int32X4) Int32X4 {
	return Int32X4(C.vqaddq_s32(C.int32x4_t(v0), C.int32x4_t(v1)))
}

func VqaddqS64(v0 Int64X2, v1 Int64X2) Int64X2 {
	return Int64X2(C.vqaddq_s64(C.int64x2_t(v0), C.int64x2_t(v1)))
}

func VqaddqS16(v0 Int16X8, v1 Int16X8) Int16X8 {
	return Int16X8(C.vqaddq_s16(C.int16x8_t(v0), C.int16x8_t(v1)))
}

func VqaddU8(v0 Uint8X8, v1 Uint8X8) Uint8X8 {
	return Uint8X8(C.vqadd_u8(C.uint8x8_t(v0), C.uint8x8_t(v1)))
}

func VqaddU32(v0 Uint32X2, v1 Uint32X2) Uint32X2 {
	return Uint32X2(C.vqadd_u32(C.uint32x2_t(v0), C.uint32x2_t(v1)))
}

func VqaddU64(v0 Uint64X1, v1 Uint64X1) Uint64X1 {
	return Uint64X1(C.vqadd_u64(C.uint64x1_t(v0), C.uint64x1_t(v1)))
}

func VqaddU16(v0 Uint16X4, v1 Uint16X4) Uint16X4 {
	return Uint16X4(C.vqadd_u16(C.uint16x4_t(v0), C.uint16x4_t(v1)))
}

func VqaddS8(v0 Int8X8, v1 Int8X8) Int8X8 {
	return Int8X8(C.vqadd_s8(C.int8x8_t(v0), C.int8x8_t(v1)))
}

func VqaddS32(v0 Int32X2, v1 Int32X2) Int32X2 {
	return Int32X2(C.vqadd_s32(C.int32x2_t(v0), C.int32x2_t(v1)))
}

func VqaddS64(v0 Int64X1, v1 Int64X1) Int64X1 {
	return Int64X1(C.vqadd_s64(C.int64x1_t(v0), C.int64x1_t(v1)))
}

func VqaddS16(v0 Int16X4, v1 Int16X4) Int16X4 {
	return Int16X4(C.vqadd_s16(C.int16x4_t(v0), C.int16x4_t(v1)))
}

func VqdmlalS32(v0 Int64X2, v1 Int32X2, v2 Int32X2) Int64X2 {
	return Int64X2(C.vqdmlal_s32(C.int64x2_t(v0), C.int32x2_t(v1), C.int32x2_t(v2)))
}

func VqdmlalS16(v0 Int32X4, v1 Int16X4, v2 Int16X4) Int32X4 {
	return Int32X4(C.vqdmlal_s16(C.int32x4_t(v0), C.int16x4_t(v1), C.int16x4_t(v2)))
}

func VqdmlalNS32(v0 Int64X2, v1 Int32X2, v2 int32) Int64X2 {
	return Int64X2(C.vqdmlal_n_s32(C.int64x2_t(v0), C.int32x2_t(v1), C.int32_t(v2)))
}

func VqdmlalNS16(v0 Int32X4, v1 Int16X4, v2 int16) Int32X4 {
	return Int32X4(C.vqdmlal_n_s16(C.int32x4_t(v0), C.int16x4_t(v1), C.int16_t(v2)))
}

func VqdmlslS32(v0 Int64X2, v1 Int32X2, v2 Int32X2) Int64X2 {
	return Int64X2(C.vqdmlsl_s32(C.int64x2_t(v0), C.int32x2_t(v1), C.int32x2_t(v2)))
}

func VqdmlslS16(v0 Int32X4, v1 Int16X4, v2 Int16X4) Int32X4 {
	return Int32X4(C.vqdmlsl_s16(C.int32x4_t(v0), C.int16x4_t(v1), C.int16x4_t(v2)))
}

func VqdmlslNS32(v0 Int64X2, v1 Int32X2, v2 int32) Int64X2 {
	return Int64X2(C.vqdmlsl_n_s32(C.int64x2_t(v0), C.int32x2_t(v1), C.int32_t(v2)))
}

func VqdmlslNS16(v0 Int32X4, v1 Int16X4, v2 int16) Int32X4 {
	return Int32X4(C.vqdmlsl_n_s16(C.int32x4_t(v0), C.int16x4_t(v1), C.int16_t(v2)))
}

func VqdmulhqS32(v0 Int32X4, v1 Int32X4) Int32X4 {
	return Int32X4(C.vqdmulhq_s32(C.int32x4_t(v0), C.int32x4_t(v1)))
}

func VqdmulhqS16(v0 Int16X8, v1 Int16X8) Int16X8 {
	return Int16X8(C.vqdmulhq_s16(C.int16x8_t(v0), C.int16x8_t(v1)))
}

func VqdmulhS32(v0 Int32X2, v1 Int32X2) Int32X2 {
	return Int32X2(C.vqdmulh_s32(C.int32x2_t(v0), C.int32x2_t(v1)))
}

func VqdmulhS16(v0 Int16X4, v1 Int16X4) Int16X4 {
	return Int16X4(C.vqdmulh_s16(C.int16x4_t(v0), C.int16x4_t(v1)))
}

func VqdmulhqNS32(v0 Int32X4, v1 int32) Int32X4 {
	return Int32X4(C.vqdmulhq_n_s32(C.int32x4_t(v0), C.int32_t(v1)))
}

func VqdmulhqNS16(v0 Int16X8, v1 int16) Int16X8 {
	return Int16X8(C.vqdmulhq_n_s16(C.int16x8_t(v0), C.int16_t(v1)))
}

func VqdmulhNS32(v0 Int32X2, v1 int32) Int32X2 {
	return Int32X2(C.vqdmulh_n_s32(C.int32x2_t(v0), C.int32_t(v1)))
}

func VqdmulhNS16(v0 Int16X4, v1 int16) Int16X4 {
	return Int16X4(C.vqdmulh_n_s16(C.int16x4_t(v0), C.int16_t(v1)))
}

func VqdmullS32(v0 Int32X2, v1 Int32X2) Int64X2 {
	return Int64X2(C.vqdmull_s32(C.int32x2_t(v0), C.int32x2_t(v1)))
}

func VqdmullS16(v0 Int16X4, v1 Int16X4) Int32X4 {
	return Int32X4(C.vqdmull_s16(C.int16x4_t(v0), C.int16x4_t(v1)))
}

func VqdmullNS32(v0 Int32X2, v1 int32) Int64X2 {
	return Int64X2(C.vqdmull_n_s32(C.int32x2_t(v0), C.int32_t(v1)))
}

func VqdmullNS16(v0 Int16X4, v1 int16) Int32X4 {
	return Int32X4(C.vqdmull_n_s16(C.int16x4_t(v0), C.int16_t(v1)))
}

func VqmovnU32(v0 Uint32X4) Uint16X4 {
	return Uint16X4(C.vqmovn_u32(C.uint32x4_t(v0)))
}

func VqmovnU64(v0 Uint64X2) Uint32X2 {
	return Uint32X2(C.vqmovn_u64(C.uint64x2_t(v0)))
}

func VqmovnU16(v0 Uint16X8) Uint8X8 {
	return Uint8X8(C.vqmovn_u16(C.uint16x8_t(v0)))
}

func VqmovnS32(v0 Int32X4) Int16X4 {
	return Int16X4(C.vqmovn_s32(C.int32x4_t(v0)))
}

func VqmovnS64(v0 Int64X2) Int32X2 {
	return Int32X2(C.vqmovn_s64(C.int64x2_t(v0)))
}

func VqmovnS16(v0 Int16X8) Int8X8 {
	return Int8X8(C.vqmovn_s16(C.int16x8_t(v0)))
}

func VqmovunS32(v0 Int32X4) Uint16X4 {
	return Uint16X4(C.vqmovun_s32(C.int32x4_t(v0)))
}

func VqmovunS64(v0 Int64X2) Uint32X2 {
	return Uint32X2(C.vqmovun_s64(C.int64x2_t(v0)))
}

func VqmovunS16(v0 Int16X8) Uint8X8 {
	return Uint8X8(C.vqmovun_s16(C.int16x8_t(v0)))
}

func VqnegqS8(v0 Int8X16) Int8X16 {
	return Int8X16(C.vqnegq_s8(C.int8x16_t(v0)))
}

func VqnegqS32(v0 Int32X4) Int32X4 {
	return Int32X4(C.vqnegq_s32(C.int32x4_t(v0)))
}

func VqnegqS16(v0 Int16X8) Int16X8 {
	return Int16X8(C.vqnegq_s16(C.int16x8_t(v0)))
}

func VqnegS8(v0 Int8X8) Int8X8 {
	return Int8X8(C.vqneg_s8(C.int8x8_t(v0)))
}

func VqnegS32(v0 Int32X2) Int32X2 {
	return Int32X2(C.vqneg_s32(C.int32x2_t(v0)))
}

func VqnegS16(v0 Int16X4) Int16X4 {
	return Int16X4(C.vqneg_s16(C.int16x4_t(v0)))
}

func VqrdmulhqS32(v0 Int32X4, v1 Int32X4) Int32X4 {
	return Int32X4(C.vqrdmulhq_s32(C.int32x4_t(v0), C.int32x4_t(v1)))
}

func VqrdmulhqS16(v0 Int16X8, v1 Int16X8) Int16X8 {
	return Int16X8(C.vqrdmulhq_s16(C.int16x8_t(v0), C.int16x8_t(v1)))
}

func VqrdmulhS32(v0 Int32X2, v1 Int32X2) Int32X2 {
	return Int32X2(C.vqrdmulh_s32(C.int32x2_t(v0), C.int32x2_t(v1)))
}

func VqrdmulhS16(v0 Int16X4, v1 Int16X4) Int16X4 {
	return Int16X4(C.vqrdmulh_s16(C.int16x4_t(v0), C.int16x4_t(v1)))
}

func VqrdmulhqNS32(v0 Int32X4, v1 int32) Int32X4 {
	return Int32X4(C.vqrdmulhq_n_s32(C.int32x4_t(v0), C.int32_t(v1)))
}

func VqrdmulhqNS16(v0 Int16X8, v1 int16) Int16X8 {
	return Int16X8(C.vqrdmulhq_n_s16(C.int16x8_t(v0), C.int16_t(v1)))
}

func VqrdmulhNS32(v0 Int32X2, v1 int32) Int32X2 {
	return Int32X2(C.vqrdmulh_n_s32(C.int32x2_t(v0), C.int32_t(v1)))
}

func VqrdmulhNS16(v0 Int16X4, v1 int16) Int16X4 {
	return Int16X4(C.vqrdmulh_n_s16(C.int16x4_t(v0), C.int16_t(v1)))
}

func VqrshlqU8(v0 Uint8X16, v1 Int8X16) Uint8X16 {
	return Uint8X16(C.vqrshlq_u8(C.uint8x16_t(v0), C.int8x16_t(v1)))
}

func VqrshlqU32(v0 Uint32X4, v1 Int32X4) Uint32X4 {
	return Uint32X4(C.vqrshlq_u32(C.uint32x4_t(v0), C.int32x4_t(v1)))
}

func VqrshlqU64(v0 Uint64X2, v1 Int64X2) Uint64X2 {
	return Uint64X2(C.vqrshlq_u64(C.uint64x2_t(v0), C.int64x2_t(v1)))
}

func VqrshlqU16(v0 Uint16X8, v1 Int16X8) Uint16X8 {
	return Uint16X8(C.vqrshlq_u16(C.uint16x8_t(v0), C.int16x8_t(v1)))
}

func VqrshlqS8(v0 Int8X16, v1 Int8X16) Int8X16 {
	return Int8X16(C.vqrshlq_s8(C.int8x16_t(v0), C.int8x16_t(v1)))
}

func VqrshlqS32(v0 Int32X4, v1 Int32X4) Int32X4 {
	return Int32X4(C.vqrshlq_s32(C.int32x4_t(v0), C.int32x4_t(v1)))
}

func VqrshlqS64(v0 Int64X2, v1 Int64X2) Int64X2 {
	return Int64X2(C.vqrshlq_s64(C.int64x2_t(v0), C.int64x2_t(v1)))
}

func VqrshlqS16(v0 Int16X8, v1 Int16X8) Int16X8 {
	return Int16X8(C.vqrshlq_s16(C.int16x8_t(v0), C.int16x8_t(v1)))
}

func VqrshlU8(v0 Uint8X8, v1 Int8X8) Uint8X8 {
	return Uint8X8(C.vqrshl_u8(C.uint8x8_t(v0), C.int8x8_t(v1)))
}

func VqrshlU32(v0 Uint32X2, v1 Int32X2) Uint32X2 {
	return Uint32X2(C.vqrshl_u32(C.uint32x2_t(v0), C.int32x2_t(v1)))
}

func VqrshlU64(v0 Uint64X1, v1 Int64X1) Uint64X1 {
	return Uint64X1(C.vqrshl_u64(C.uint64x1_t(v0), C.int64x1_t(v1)))
}

func VqrshlU16(v0 Uint16X4, v1 Int16X4) Uint16X4 {
	return Uint16X4(C.vqrshl_u16(C.uint16x4_t(v0), C.int16x4_t(v1)))
}

func VqrshlS8(v0 Int8X8, v1 Int8X8) Int8X8 {
	return Int8X8(C.vqrshl_s8(C.int8x8_t(v0), C.int8x8_t(v1)))
}

func VqrshlS32(v0 Int32X2, v1 Int32X2) Int32X2 {
	return Int32X2(C.vqrshl_s32(C.int32x2_t(v0), C.int32x2_t(v1)))
}

func VqrshlS64(v0 Int64X1, v1 Int64X1) Int64X1 {
	return Int64X1(C.vqrshl_s64(C.int64x1_t(v0), C.int64x1_t(v1)))
}

func VqrshlS16(v0 Int16X4, v1 Int16X4) Int16X4 {
	return Int16X4(C.vqrshl_s16(C.int16x4_t(v0), C.int16x4_t(v1)))
}

func VqshlqU8(v0 Uint8X16, v1 Int8X16) Uint8X16 {
	return Uint8X16(C.vqshlq_u8(C.uint8x16_t(v0), C.int8x16_t(v1)))
}

func VqshlqU32(v0 Uint32X4, v1 Int32X4) Uint32X4 {
	return Uint32X4(C.vqshlq_u32(C.uint32x4_t(v0), C.int32x4_t(v1)))
}

func VqshlqU64(v0 Uint64X2, v1 Int64X2) Uint64X2 {
	return Uint64X2(C.vqshlq_u64(C.uint64x2_t(v0), C.int64x2_t(v1)))
}

func VqshlqU16(v0 Uint16X8, v1 Int16X8) Uint16X8 {
	return Uint16X8(C.vqshlq_u16(C.uint16x8_t(v0), C.int16x8_t(v1)))
}

func VqshlqS8(v0 Int8X16, v1 Int8X16) Int8X16 {
	return Int8X16(C.vqshlq_s8(C.int8x16_t(v0), C.int8x16_t(v1)))
}

func VqshlqS32(v0 Int32X4, v1 Int32X4) Int32X4 {
	return Int32X4(C.vqshlq_s32(C.int32x4_t(v0), C.int32x4_t(v1)))
}

func VqshlqS64(v0 Int64X2, v1 Int64X2) Int64X2 {
	return Int64X2(C.vqshlq_s64(C.int64x2_t(v0), C.int64x2_t(v1)))
}

func VqshlqS16(v0 Int16X8, v1 Int16X8) Int16X8 {
	return Int16X8(C.vqshlq_s16(C.int16x8_t(v0), C.int16x8_t(v1)))
}

func VqshlU8(v0 Uint8X8, v1 Int8X8) Uint8X8 {
	return Uint8X8(C.vqshl_u8(C.uint8x8_t(v0), C.int8x8_t(v1)))
}

func VqshlU32(v0 Uint32X2, v1 Int32X2) Uint32X2 {
	return Uint32X2(C.vqshl_u32(C.uint32x2_t(v0), C.int32x2_t(v1)))
}

func VqshlU64(v0 Uint64X1, v1 Int64X1) Uint64X1 {
	return Uint64X1(C.vqshl_u64(C.uint64x1_t(v0), C.int64x1_t(v1)))
}

func VqshlU16(v0 Uint16X4, v1 Int16X4) Uint16X4 {
	return Uint16X4(C.vqshl_u16(C.uint16x4_t(v0), C.int16x4_t(v1)))
}

func VqshlS8(v0 Int8X8, v1 Int8X8) Int8X8 {
	return Int8X8(C.vqshl_s8(C.int8x8_t(v0), C.int8x8_t(v1)))
}

func VqshlS32(v0 Int32X2, v1 Int32X2) Int32X2 {
	return Int32X2(C.vqshl_s32(C.int32x2_t(v0), C.int32x2_t(v1)))
}

func VqshlS64(v0 Int64X1, v1 Int64X1) Int64X1 {
	return Int64X1(C.vqshl_s64(C.int64x1_t(v0), C.int64x1_t(v1)))
}

func VqshlS16(v0 Int16X4, v1 Int16X4) Int16X4 {
	return Int16X4(C.vqshl_s16(C.int16x4_t(v0), C.int16x4_t(v1)))
}

func VqsubqU8(v0 Uint8X16, v1 Uint8X16) Uint8X16 {
	return Uint8X16(C.vqsubq_u8(C.uint8x16_t(v0), C.uint8x16_t(v1)))
}

func VqsubqU32(v0 Uint32X4, v1 Uint32X4) Uint32X4 {
	return Uint32X4(C.vqsubq_u32(C.uint32x4_t(v0), C.uint32x4_t(v1)))
}

func VqsubqU64(v0 Uint64X2, v1 Uint64X2) Uint64X2 {
	return Uint64X2(C.vqsubq_u64(C.uint64x2_t(v0), C.uint64x2_t(v1)))
}

func VqsubqU16(v0 Uint16X8, v1 Uint16X8) Uint16X8 {
	return Uint16X8(C.vqsubq_u16(C.uint16x8_t(v0), C.uint16x8_t(v1)))
}

func VqsubqS8(v0 Int8X16, v1 Int8X16) Int8X16 {
	return Int8X16(C.vqsubq_s8(C.int8x16_t(v0), C.int8x16_t(v1)))
}

func VqsubqS32(v0 Int32X4, v1 Int32X4) Int32X4 {
	return Int32X4(C.vqsubq_s32(C.int32x4_t(v0), C.int32x4_t(v1)))
}

func VqsubqS64(v0 Int64X2, v1 Int64X2) Int64X2 {
	return Int64X2(C.vqsubq_s64(C.int64x2_t(v0), C.int64x2_t(v1)))
}

func VqsubqS16(v0 Int16X8, v1 Int16X8) Int16X8 {
	return Int16X8(C.vqsubq_s16(C.int16x8_t(v0), C.int16x8_t(v1)))
}

func VqsubU8(v0 Uint8X8, v1 Uint8X8) Uint8X8 {
	return Uint8X8(C.vqsub_u8(C.uint8x8_t(v0), C.uint8x8_t(v1)))
}

func VqsubU32(v0 Uint32X2, v1 Uint32X2) Uint32X2 {
	return Uint32X2(C.vqsub_u32(C.uint32x2_t(v0), C.uint32x2_t(v1)))
}

func VqsubU64(v0 Uint64X1, v1 Uint64X1) Uint64X1 {
	return Uint64X1(C.vqsub_u64(C.uint64x1_t(v0), C.uint64x1_t(v1)))
}

func VqsubU16(v0 Uint16X4, v1 Uint16X4) Uint16X4 {
	return Uint16X4(C.vqsub_u16(C.uint16x4_t(v0), C.uint16x4_t(v1)))
}

func VqsubS8(v0 Int8X8, v1 Int8X8) Int8X8 {
	return Int8X8(C.vqsub_s8(C.int8x8_t(v0), C.int8x8_t(v1)))
}

func VqsubS32(v0 Int32X2, v1 Int32X2) Int32X2 {
	return Int32X2(C.vqsub_s32(C.int32x2_t(v0), C.int32x2_t(v1)))
}

func VqsubS64(v0 Int64X1, v1 Int64X1) Int64X1 {
	return Int64X1(C.vqsub_s64(C.int64x1_t(v0), C.int64x1_t(v1)))
}

func VqsubS16(v0 Int16X4, v1 Int16X4) Int16X4 {
	return Int16X4(C.vqsub_s16(C.int16x4_t(v0), C.int16x4_t(v1)))
}

func VraddhnU32(v0 Uint32X4, v1 Uint32X4) Uint16X4 {
	return Uint16X4(C.vraddhn_u32(C.uint32x4_t(v0), C.uint32x4_t(v1)))
}

func VraddhnU64(v0 Uint64X2, v1 Uint64X2) Uint32X2 {
	return Uint32X2(C.vraddhn_u64(C.uint64x2_t(v0), C.uint64x2_t(v1)))
}

func VraddhnU16(v0 Uint16X8, v1 Uint16X8) Uint8X8 {
	return Uint8X8(C.vraddhn_u16(C.uint16x8_t(v0), C.uint16x8_t(v1)))
}

func VraddhnS32(v0 Int32X4, v1 Int32X4) Int16X4 {
	return Int16X4(C.vraddhn_s32(C.int32x4_t(v0), C.int32x4_t(v1)))
}

func VraddhnS64(v0 Int64X2, v1 Int64X2) Int32X2 {
	return Int32X2(C.vraddhn_s64(C.int64x2_t(v0), C.int64x2_t(v1)))
}

func VraddhnS16(v0 Int16X8, v1 Int16X8) Int8X8 {
	return Int8X8(C.vraddhn_s16(C.int16x8_t(v0), C.int16x8_t(v1)))
}

func VrecpeqU32(v0 Uint32X4) Uint32X4 {
	return Uint32X4(C.vrecpeq_u32(C.uint32x4_t(v0)))
}

func VrecpeqF32(v0 Float32X4) Float32X4 {
	return Float32X4(C.vrecpeq_f32(C.float32x4_t(v0)))
}

func VrecpeU32(v0 Uint32X2) Uint32X2 {
	return Uint32X2(C.vrecpe_u32(C.uint32x2_t(v0)))
}

func VrecpeF32(v0 Float32X2) Float32X2 {
	return Float32X2(C.vrecpe_f32(C.float32x2_t(v0)))
}

func VrecpsqF32(v0 Float32X4, v1 Float32X4) Float32X4 {
	return Float32X4(C.vrecpsq_f32(C.float32x4_t(v0), C.float32x4_t(v1)))
}

func VrecpsF32(v0 Float32X2, v1 Float32X2) Float32X2 {
	return Float32X2(C.vrecps_f32(C.float32x2_t(v0), C.float32x2_t(v1)))
}

func Vrev16P8(v0 Poly8X8) Poly8X8 {
	return Poly8X8(C.vrev16_p8(C.poly8x8_t(v0)))
}

func Vrev16QP8(v0 Poly8X16) Poly8X16 {
	return Poly8X16(C.vrev16q_p8(C.poly8x16_t(v0)))
}

func Vrev16QU8(v0 Uint8X16) Uint8X16 {
	return Uint8X16(C.vrev16q_u8(C.uint8x16_t(v0)))
}

func Vrev16QS8(v0 Int8X16) Int8X16 {
	return Int8X16(C.vrev16q_s8(C.int8x16_t(v0)))
}

func Vrev16U8(v0 Uint8X8) Uint8X8 {
	return Uint8X8(C.vrev16_u8(C.uint8x8_t(v0)))
}

func Vrev16S8(v0 Int8X8) Int8X8 {
	return Int8X8(C.vrev16_s8(C.int8x8_t(v0)))
}

func Vrev32P8(v0 Poly8X8) Poly8X8 {
	return Poly8X8(C.vrev32_p8(C.poly8x8_t(v0)))
}

func Vrev32P16(v0 Poly16X4) Poly16X4 {
	return Poly16X4(C.vrev32_p16(C.poly16x4_t(v0)))
}

func Vrev32QP8(v0 Poly8X16) Poly8X16 {
	return Poly8X16(C.vrev32q_p8(C.poly8x16_t(v0)))
}

func Vrev32QP16(v0 Poly16X8) Poly16X8 {
	return Poly16X8(C.vrev32q_p16(C.poly16x8_t(v0)))
}

func Vrev32QU8(v0 Uint8X16) Uint8X16 {
	return Uint8X16(C.vrev32q_u8(C.uint8x16_t(v0)))
}

func Vrev32QU16(v0 Uint16X8) Uint16X8 {
	return Uint16X8(C.vrev32q_u16(C.uint16x8_t(v0)))
}

func Vrev32QS8(v0 Int8X16) Int8X16 {
	return Int8X16(C.vrev32q_s8(C.int8x16_t(v0)))
}

func Vrev32QS16(v0 Int16X8) Int16X8 {
	return Int16X8(C.vrev32q_s16(C.int16x8_t(v0)))
}

func Vrev32U8(v0 Uint8X8) Uint8X8 {
	return Uint8X8(C.vrev32_u8(C.uint8x8_t(v0)))
}

func Vrev32U16(v0 Uint16X4) Uint16X4 {
	return Uint16X4(C.vrev32_u16(C.uint16x4_t(v0)))
}

func Vrev32S8(v0 Int8X8) Int8X8 {
	return Int8X8(C.vrev32_s8(C.int8x8_t(v0)))
}

func Vrev32S16(v0 Int16X4) Int16X4 {
	return Int16X4(C.vrev32_s16(C.int16x4_t(v0)))
}

func Vrev64P8(v0 Poly8X8) Poly8X8 {
	return Poly8X8(C.vrev64_p8(C.poly8x8_t(v0)))
}

func Vrev64P16(v0 Poly16X4) Poly16X4 {
	return Poly16X4(C.vrev64_p16(C.poly16x4_t(v0)))
}

func Vrev64QP8(v0 Poly8X16) Poly8X16 {
	return Poly8X16(C.vrev64q_p8(C.poly8x16_t(v0)))
}

func Vrev64QP16(v0 Poly16X8) Poly16X8 {
	return Poly16X8(C.vrev64q_p16(C.poly16x8_t(v0)))
}

func Vrev64QU8(v0 Uint8X16) Uint8X16 {
	return Uint8X16(C.vrev64q_u8(C.uint8x16_t(v0)))
}

func Vrev64QU32(v0 Uint32X4) Uint32X4 {
	return Uint32X4(C.vrev64q_u32(C.uint32x4_t(v0)))
}

func Vrev64QU16(v0 Uint16X8) Uint16X8 {
	return Uint16X8(C.vrev64q_u16(C.uint16x8_t(v0)))
}

func Vrev64QS8(v0 Int8X16) Int8X16 {
	return Int8X16(C.vrev64q_s8(C.int8x16_t(v0)))
}

func Vrev64QF32(v0 Float32X4) Float32X4 {
	return Float32X4(C.vrev64q_f32(C.float32x4_t(v0)))
}

func Vrev64QS32(v0 Int32X4) Int32X4 {
	return Int32X4(C.vrev64q_s32(C.int32x4_t(v0)))
}

func Vrev64QS16(v0 Int16X8) Int16X8 {
	return Int16X8(C.vrev64q_s16(C.int16x8_t(v0)))
}

func Vrev64U8(v0 Uint8X8) Uint8X8 {
	return Uint8X8(C.vrev64_u8(C.uint8x8_t(v0)))
}

func Vrev64U32(v0 Uint32X2) Uint32X2 {
	return Uint32X2(C.vrev64_u32(C.uint32x2_t(v0)))
}

func Vrev64U16(v0 Uint16X4) Uint16X4 {
	return Uint16X4(C.vrev64_u16(C.uint16x4_t(v0)))
}

func Vrev64S8(v0 Int8X8) Int8X8 {
	return Int8X8(C.vrev64_s8(C.int8x8_t(v0)))
}

func Vrev64F32(v0 Float32X2) Float32X2 {
	return Float32X2(C.vrev64_f32(C.float32x2_t(v0)))
}

func Vrev64S32(v0 Int32X2) Int32X2 {
	return Int32X2(C.vrev64_s32(C.int32x2_t(v0)))
}

func Vrev64S16(v0 Int16X4) Int16X4 {
	return Int16X4(C.vrev64_s16(C.int16x4_t(v0)))
}

func VrhaddqU8(v0 Uint8X16, v1 Uint8X16) Uint8X16 {
	return Uint8X16(C.vrhaddq_u8(C.uint8x16_t(v0), C.uint8x16_t(v1)))
}

func VrhaddqU32(v0 Uint32X4, v1 Uint32X4) Uint32X4 {
	return Uint32X4(C.vrhaddq_u32(C.uint32x4_t(v0), C.uint32x4_t(v1)))
}

func VrhaddqU16(v0 Uint16X8, v1 Uint16X8) Uint16X8 {
	return Uint16X8(C.vrhaddq_u16(C.uint16x8_t(v0), C.uint16x8_t(v1)))
}

func VrhaddqS8(v0 Int8X16, v1 Int8X16) Int8X16 {
	return Int8X16(C.vrhaddq_s8(C.int8x16_t(v0), C.int8x16_t(v1)))
}

func VrhaddqS32(v0 Int32X4, v1 Int32X4) Int32X4 {
	return Int32X4(C.vrhaddq_s32(C.int32x4_t(v0), C.int32x4_t(v1)))
}

func VrhaddqS16(v0 Int16X8, v1 Int16X8) Int16X8 {
	return Int16X8(C.vrhaddq_s16(C.int16x8_t(v0), C.int16x8_t(v1)))
}

func VrhaddU8(v0 Uint8X8, v1 Uint8X8) Uint8X8 {
	return Uint8X8(C.vrhadd_u8(C.uint8x8_t(v0), C.uint8x8_t(v1)))
}

func VrhaddU32(v0 Uint32X2, v1 Uint32X2) Uint32X2 {
	return Uint32X2(C.vrhadd_u32(C.uint32x2_t(v0), C.uint32x2_t(v1)))
}

func VrhaddU16(v0 Uint16X4, v1 Uint16X4) Uint16X4 {
	return Uint16X4(C.vrhadd_u16(C.uint16x4_t(v0), C.uint16x4_t(v1)))
}

func VrhaddS8(v0 Int8X8, v1 Int8X8) Int8X8 {
	return Int8X8(C.vrhadd_s8(C.int8x8_t(v0), C.int8x8_t(v1)))
}

func VrhaddS32(v0 Int32X2, v1 Int32X2) Int32X2 {
	return Int32X2(C.vrhadd_s32(C.int32x2_t(v0), C.int32x2_t(v1)))
}

func VrhaddS16(v0 Int16X4, v1 Int16X4) Int16X4 {
	return Int16X4(C.vrhadd_s16(C.int16x4_t(v0), C.int16x4_t(v1)))
}

func VrshlqU8(v0 Uint8X16, v1 Int8X16) Uint8X16 {
	return Uint8X16(C.vrshlq_u8(C.uint8x16_t(v0), C.int8x16_t(v1)))
}

func VrshlqU32(v0 Uint32X4, v1 Int32X4) Uint32X4 {
	return Uint32X4(C.vrshlq_u32(C.uint32x4_t(v0), C.int32x4_t(v1)))
}

func VrshlqU64(v0 Uint64X2, v1 Int64X2) Uint64X2 {
	return Uint64X2(C.vrshlq_u64(C.uint64x2_t(v0), C.int64x2_t(v1)))
}

func VrshlqU16(v0 Uint16X8, v1 Int16X8) Uint16X8 {
	return Uint16X8(C.vrshlq_u16(C.uint16x8_t(v0), C.int16x8_t(v1)))
}

func VrshlqS8(v0 Int8X16, v1 Int8X16) Int8X16 {
	return Int8X16(C.vrshlq_s8(C.int8x16_t(v0), C.int8x16_t(v1)))
}

func VrshlqS32(v0 Int32X4, v1 Int32X4) Int32X4 {
	return Int32X4(C.vrshlq_s32(C.int32x4_t(v0), C.int32x4_t(v1)))
}

func VrshlqS64(v0 Int64X2, v1 Int64X2) Int64X2 {
	return Int64X2(C.vrshlq_s64(C.int64x2_t(v0), C.int64x2_t(v1)))
}

func VrshlqS16(v0 Int16X8, v1 Int16X8) Int16X8 {
	return Int16X8(C.vrshlq_s16(C.int16x8_t(v0), C.int16x8_t(v1)))
}

func VrshlU8(v0 Uint8X8, v1 Int8X8) Uint8X8 {
	return Uint8X8(C.vrshl_u8(C.uint8x8_t(v0), C.int8x8_t(v1)))
}

func VrshlU32(v0 Uint32X2, v1 Int32X2) Uint32X2 {
	return Uint32X2(C.vrshl_u32(C.uint32x2_t(v0), C.int32x2_t(v1)))
}

func VrshlU64(v0 Uint64X1, v1 Int64X1) Uint64X1 {
	return Uint64X1(C.vrshl_u64(C.uint64x1_t(v0), C.int64x1_t(v1)))
}

func VrshlU16(v0 Uint16X4, v1 Int16X4) Uint16X4 {
	return Uint16X4(C.vrshl_u16(C.uint16x4_t(v0), C.int16x4_t(v1)))
}

func VrshlS8(v0 Int8X8, v1 Int8X8) Int8X8 {
	return Int8X8(C.vrshl_s8(C.int8x8_t(v0), C.int8x8_t(v1)))
}

func VrshlS32(v0 Int32X2, v1 Int32X2) Int32X2 {
	return Int32X2(C.vrshl_s32(C.int32x2_t(v0), C.int32x2_t(v1)))
}

func VrshlS64(v0 Int64X1, v1 Int64X1) Int64X1 {
	return Int64X1(C.vrshl_s64(C.int64x1_t(v0), C.int64x1_t(v1)))
}

func VrshlS16(v0 Int16X4, v1 Int16X4) Int16X4 {
	return Int16X4(C.vrshl_s16(C.int16x4_t(v0), C.int16x4_t(v1)))
}

func VrsqrteqU32(v0 Uint32X4) Uint32X4 {
	return Uint32X4(C.vrsqrteq_u32(C.uint32x4_t(v0)))
}

func VrsqrteqF32(v0 Float32X4) Float32X4 {
	return Float32X4(C.vrsqrteq_f32(C.float32x4_t(v0)))
}

func VrsqrteU32(v0 Uint32X2) Uint32X2 {
	return Uint32X2(C.vrsqrte_u32(C.uint32x2_t(v0)))
}

func VrsqrteF32(v0 Float32X2) Float32X2 {
	return Float32X2(C.vrsqrte_f32(C.float32x2_t(v0)))
}

func VrsqrtsqF32(v0 Float32X4, v1 Float32X4) Float32X4 {
	return Float32X4(C.vrsqrtsq_f32(C.float32x4_t(v0), C.float32x4_t(v1)))
}

func VrsqrtsF32(v0 Float32X2, v1 Float32X2) Float32X2 {
	return Float32X2(C.vrsqrts_f32(C.float32x2_t(v0), C.float32x2_t(v1)))
}

func VrsubhnU32(v0 Uint32X4, v1 Uint32X4) Uint16X4 {
	return Uint16X4(C.vrsubhn_u32(C.uint32x4_t(v0), C.uint32x4_t(v1)))
}

func VrsubhnU64(v0 Uint64X2, v1 Uint64X2) Uint32X2 {
	return Uint32X2(C.vrsubhn_u64(C.uint64x2_t(v0), C.uint64x2_t(v1)))
}

func VrsubhnU16(v0 Uint16X8, v1 Uint16X8) Uint8X8 {
	return Uint8X8(C.vrsubhn_u16(C.uint16x8_t(v0), C.uint16x8_t(v1)))
}

func VrsubhnS32(v0 Int32X4, v1 Int32X4) Int16X4 {
	return Int16X4(C.vrsubhn_s32(C.int32x4_t(v0), C.int32x4_t(v1)))
}

func VrsubhnS64(v0 Int64X2, v1 Int64X2) Int32X2 {
	return Int32X2(C.vrsubhn_s64(C.int64x2_t(v0), C.int64x2_t(v1)))
}

func VrsubhnS16(v0 Int16X8, v1 Int16X8) Int8X8 {
	return Int8X8(C.vrsubhn_s16(C.int16x8_t(v0), C.int16x8_t(v1)))
}

func VshlqU8(v0 Uint8X16, v1 Int8X16) Uint8X16 {
	return Uint8X16(C.vshlq_u8(C.uint8x16_t(v0), C.int8x16_t(v1)))
}

func VshlqU32(v0 Uint32X4, v1 Int32X4) Uint32X4 {
	return Uint32X4(C.vshlq_u32(C.uint32x4_t(v0), C.int32x4_t(v1)))
}

func VshlqU64(v0 Uint64X2, v1 Int64X2) Uint64X2 {
	return Uint64X2(C.vshlq_u64(C.uint64x2_t(v0), C.int64x2_t(v1)))
}

func VshlqU16(v0 Uint16X8, v1 Int16X8) Uint16X8 {
	return Uint16X8(C.vshlq_u16(C.uint16x8_t(v0), C.int16x8_t(v1)))
}

func VshlqS8(v0 Int8X16, v1 Int8X16) Int8X16 {
	return Int8X16(C.vshlq_s8(C.int8x16_t(v0), C.int8x16_t(v1)))
}

func VshlqS32(v0 Int32X4, v1 Int32X4) Int32X4 {
	return Int32X4(C.vshlq_s32(C.int32x4_t(v0), C.int32x4_t(v1)))
}

func VshlqS64(v0 Int64X2, v1 Int64X2) Int64X2 {
	return Int64X2(C.vshlq_s64(C.int64x2_t(v0), C.int64x2_t(v1)))
}

func VshlqS16(v0 Int16X8, v1 Int16X8) Int16X8 {
	return Int16X8(C.vshlq_s16(C.int16x8_t(v0), C.int16x8_t(v1)))
}

func VshlU8(v0 Uint8X8, v1 Int8X8) Uint8X8 {
	return Uint8X8(C.vshl_u8(C.uint8x8_t(v0), C.int8x8_t(v1)))
}

func VshlU32(v0 Uint32X2, v1 Int32X2) Uint32X2 {
	return Uint32X2(C.vshl_u32(C.uint32x2_t(v0), C.int32x2_t(v1)))
}

func VshlU64(v0 Uint64X1, v1 Int64X1) Uint64X1 {
	return Uint64X1(C.vshl_u64(C.uint64x1_t(v0), C.int64x1_t(v1)))
}

func VshlU16(v0 Uint16X4, v1 Int16X4) Uint16X4 {
	return Uint16X4(C.vshl_u16(C.uint16x4_t(v0), C.int16x4_t(v1)))
}

func VshlS8(v0 Int8X8, v1 Int8X8) Int8X8 {
	return Int8X8(C.vshl_s8(C.int8x8_t(v0), C.int8x8_t(v1)))
}

func VshlS32(v0 Int32X2, v1 Int32X2) Int32X2 {
	return Int32X2(C.vshl_s32(C.int32x2_t(v0), C.int32x2_t(v1)))
}

func VshlS64(v0 Int64X1, v1 Int64X1) Int64X1 {
	return Int64X1(C.vshl_s64(C.int64x1_t(v0), C.int64x1_t(v1)))
}

func VshlS16(v0 Int16X4, v1 Int16X4) Int16X4 {
	return Int16X4(C.vshl_s16(C.int16x4_t(v0), C.int16x4_t(v1)))
}

func VsubqU8(v0 Uint8X16, v1 Uint8X16) Uint8X16 {
	return Uint8X16(C.vsubq_u8(C.uint8x16_t(v0), C.uint8x16_t(v1)))
}

func VsubqU32(v0 Uint32X4, v1 Uint32X4) Uint32X4 {
	return Uint32X4(C.vsubq_u32(C.uint32x4_t(v0), C.uint32x4_t(v1)))
}

func VsubqU64(v0 Uint64X2, v1 Uint64X2) Uint64X2 {
	return Uint64X2(C.vsubq_u64(C.uint64x2_t(v0), C.uint64x2_t(v1)))
}

func VsubqU16(v0 Uint16X8, v1 Uint16X8) Uint16X8 {
	return Uint16X8(C.vsubq_u16(C.uint16x8_t(v0), C.uint16x8_t(v1)))
}

func VsubqS8(v0 Int8X16, v1 Int8X16) Int8X16 {
	return Int8X16(C.vsubq_s8(C.int8x16_t(v0), C.int8x16_t(v1)))
}

func VsubqF32(v0 Float32X4, v1 Float32X4) Float32X4 {
	return Float32X4(C.vsubq_f32(C.float32x4_t(v0), C.float32x4_t(v1)))
}

func VsubqS32(v0 Int32X4, v1 Int32X4) Int32X4 {
	return Int32X4(C.vsubq_s32(C.int32x4_t(v0), C.int32x4_t(v1)))
}

func VsubqS64(v0 Int64X2, v1 Int64X2) Int64X2 {
	return Int64X2(C.vsubq_s64(C.int64x2_t(v0), C.int64x2_t(v1)))
}

func VsubqS16(v0 Int16X8, v1 Int16X8) Int16X8 {
	return Int16X8(C.vsubq_s16(C.int16x8_t(v0), C.int16x8_t(v1)))
}

func VsubU8(v0 Uint8X8, v1 Uint8X8) Uint8X8 {
	return Uint8X8(C.vsub_u8(C.uint8x8_t(v0), C.uint8x8_t(v1)))
}

func VsubU32(v0 Uint32X2, v1 Uint32X2) Uint32X2 {
	return Uint32X2(C.vsub_u32(C.uint32x2_t(v0), C.uint32x2_t(v1)))
}

func VsubU64(v0 Uint64X1, v1 Uint64X1) Uint64X1 {
	return Uint64X1(C.vsub_u64(C.uint64x1_t(v0), C.uint64x1_t(v1)))
}

func VsubU16(v0 Uint16X4, v1 Uint16X4) Uint16X4 {
	return Uint16X4(C.vsub_u16(C.uint16x4_t(v0), C.uint16x4_t(v1)))
}

func VsubS8(v0 Int8X8, v1 Int8X8) Int8X8 {
	return Int8X8(C.vsub_s8(C.int8x8_t(v0), C.int8x8_t(v1)))
}

func VsubF32(v0 Float32X2, v1 Float32X2) Float32X2 {
	return Float32X2(C.vsub_f32(C.float32x2_t(v0), C.float32x2_t(v1)))
}

func VsubS32(v0 Int32X2, v1 Int32X2) Int32X2 {
	return Int32X2(C.vsub_s32(C.int32x2_t(v0), C.int32x2_t(v1)))
}

func VsubS64(v0 Int64X1, v1 Int64X1) Int64X1 {
	return Int64X1(C.vsub_s64(C.int64x1_t(v0), C.int64x1_t(v1)))
}

func VsubS16(v0 Int16X4, v1 Int16X4) Int16X4 {
	return Int16X4(C.vsub_s16(C.int16x4_t(v0), C.int16x4_t(v1)))
}

func VsubhnU32(v0 Uint32X4, v1 Uint32X4) Uint16X4 {
	return Uint16X4(C.vsubhn_u32(C.uint32x4_t(v0), C.uint32x4_t(v1)))
}

func VsubhnU64(v0 Uint64X2, v1 Uint64X2) Uint32X2 {
	return Uint32X2(C.vsubhn_u64(C.uint64x2_t(v0), C.uint64x2_t(v1)))
}

func VsubhnU16(v0 Uint16X8, v1 Uint16X8) Uint8X8 {
	return Uint8X8(C.vsubhn_u16(C.uint16x8_t(v0), C.uint16x8_t(v1)))
}

func VsubhnS32(v0 Int32X4, v1 Int32X4) Int16X4 {
	return Int16X4(C.vsubhn_s32(C.int32x4_t(v0), C.int32x4_t(v1)))
}

func VsubhnS64(v0 Int64X2, v1 Int64X2) Int32X2 {
	return Int32X2(C.vsubhn_s64(C.int64x2_t(v0), C.int64x2_t(v1)))
}

func VsubhnS16(v0 Int16X8, v1 Int16X8) Int8X8 {
	return Int8X8(C.vsubhn_s16(C.int16x8_t(v0), C.int16x8_t(v1)))
}

func VsublU8(v0 Uint8X8, v1 Uint8X8) Uint16X8 {
	return Uint16X8(C.vsubl_u8(C.uint8x8_t(v0), C.uint8x8_t(v1)))
}

func VsublU32(v0 Uint32X2, v1 Uint32X2) Uint64X2 {
	return Uint64X2(C.vsubl_u32(C.uint32x2_t(v0), C.uint32x2_t(v1)))
}

func VsublU16(v0 Uint16X4, v1 Uint16X4) Uint32X4 {
	return Uint32X4(C.vsubl_u16(C.uint16x4_t(v0), C.uint16x4_t(v1)))
}

func VsublS8(v0 Int8X8, v1 Int8X8) Int16X8 {
	return Int16X8(C.vsubl_s8(C.int8x8_t(v0), C.int8x8_t(v1)))
}

func VsublS32(v0 Int32X2, v1 Int32X2) Int64X2 {
	return Int64X2(C.vsubl_s32(C.int32x2_t(v0), C.int32x2_t(v1)))
}

func VsublS16(v0 Int16X4, v1 Int16X4) Int32X4 {
	return Int32X4(C.vsubl_s16(C.int16x4_t(v0), C.int16x4_t(v1)))
}

func VsubwU8(v0 Uint16X8, v1 Uint8X8) Uint16X8 {
	return Uint16X8(C.vsubw_u8(C.uint16x8_t(v0), C.uint8x8_t(v1)))
}

func VsubwU32(v0 Uint64X2, v1 Uint32X2) Uint64X2 {
	return Uint64X2(C.vsubw_u32(C.uint64x2_t(v0), C.uint32x2_t(v1)))
}

func VsubwU16(v0 Uint32X4, v1 Uint16X4) Uint32X4 {
	return Uint32X4(C.vsubw_u16(C.uint32x4_t(v0), C.uint16x4_t(v1)))
}

func VsubwS8(v0 Int16X8, v1 Int8X8) Int16X8 {
	return Int16X8(C.vsubw_s8(C.int16x8_t(v0), C.int8x8_t(v1)))
}

func VsubwS32(v0 Int64X2, v1 Int32X2) Int64X2 {
	return Int64X2(C.vsubw_s32(C.int64x2_t(v0), C.int32x2_t(v1)))
}

func VsubwS16(v0 Int32X4, v1 Int16X4) Int32X4 {
	return Int32X4(C.vsubw_s16(C.int32x4_t(v0), C.int16x4_t(v1)))
}

func Vtbl1P8(v0 Poly8X8, v1 Uint8X8) Poly8X8 {
	return Poly8X8(C.vtbl1_p8(C.poly8x8_t(v0), C.uint8x8_t(v1)))
}

func Vtbl1U8(v0 Uint8X8, v1 Uint8X8) Uint8X8 {
	return Uint8X8(C.vtbl1_u8(C.uint8x8_t(v0), C.uint8x8_t(v1)))
}

func Vtbl1S8(v0 Int8X8, v1 Int8X8) Int8X8 {
	return Int8X8(C.vtbl1_s8(C.int8x8_t(v0), C.int8x8_t(v1)))
}

func Vtbl2P8(v0 Poly8X8X2, v1 Uint8X8) Poly8X8 {
	return Poly8X8(C.vtbl2_p8(C.poly8x8x2_t(v0), C.uint8x8_t(v1)))
}

func Vtbl2U8(v0 Uint8X8X2, v1 Uint8X8) Uint8X8 {
	return Uint8X8(C.vtbl2_u8(C.uint8x8x2_t(v0), C.uint8x8_t(v1)))
}

func Vtbl2S8(v0 Int8X8X2, v1 Int8X8) Int8X8 {
	return Int8X8(C.vtbl2_s8(C.int8x8x2_t(v0), C.int8x8_t(v1)))
}

func Vtbl3P8(v0 Poly8X8X3, v1 Uint8X8) Poly8X8 {
	return Poly8X8(C.vtbl3_p8(C.poly8x8x3_t(v0), C.uint8x8_t(v1)))
}

func Vtbl3U8(v0 Uint8X8X3, v1 Uint8X8) Uint8X8 {
	return Uint8X8(C.vtbl3_u8(C.uint8x8x3_t(v0), C.uint8x8_t(v1)))
}

func Vtbl3S8(v0 Int8X8X3, v1 Int8X8) Int8X8 {
	return Int8X8(C.vtbl3_s8(C.int8x8x3_t(v0), C.int8x8_t(v1)))
}

func Vtbl4P8(v0 Poly8X8X4, v1 Uint8X8) Poly8X8 {
	return Poly8X8(C.vtbl4_p8(C.poly8x8x4_t(v0), C.uint8x8_t(v1)))
}

func Vtbl4U8(v0 Uint8X8X4, v1 Uint8X8) Uint8X8 {
	return Uint8X8(C.vtbl4_u8(C.uint8x8x4_t(v0), C.uint8x8_t(v1)))
}

func Vtbl4S8(v0 Int8X8X4, v1 Int8X8) Int8X8 {
	return Int8X8(C.vtbl4_s8(C.int8x8x4_t(v0), C.int8x8_t(v1)))
}

func Vtbx1P8(v0 Poly8X8, v1 Poly8X8, v2 Uint8X8) Poly8X8 {
	return Poly8X8(C.vtbx1_p8(C.poly8x8_t(v0), C.poly8x8_t(v1), C.uint8x8_t(v2)))
}

func Vtbx1U8(v0 Uint8X8, v1 Uint8X8, v2 Uint8X8) Uint8X8 {
	return Uint8X8(C.vtbx1_u8(C.uint8x8_t(v0), C.uint8x8_t(v1), C.uint8x8_t(v2)))
}

func Vtbx1S8(v0 Int8X8, v1 Int8X8, v2 Int8X8) Int8X8 {
	return Int8X8(C.vtbx1_s8(C.int8x8_t(v0), C.int8x8_t(v1), C.int8x8_t(v2)))
}

func Vtbx2P8(v0 Poly8X8, v1 Poly8X8X2, v2 Uint8X8) Poly8X8 {
	return Poly8X8(C.vtbx2_p8(C.poly8x8_t(v0), C.poly8x8x2_t(v1), C.uint8x8_t(v2)))
}

func Vtbx2U8(v0 Uint8X8, v1 Uint8X8X2, v2 Uint8X8) Uint8X8 {
	return Uint8X8(C.vtbx2_u8(C.uint8x8_t(v0), C.uint8x8x2_t(v1), C.uint8x8_t(v2)))
}

func Vtbx2S8(v0 Int8X8, v1 Int8X8X2, v2 Int8X8) Int8X8 {
	return Int8X8(C.vtbx2_s8(C.int8x8_t(v0), C.int8x8x2_t(v1), C.int8x8_t(v2)))
}

func Vtbx3P8(v0 Poly8X8, v1 Poly8X8X3, v2 Uint8X8) Poly8X8 {
	return Poly8X8(C.vtbx3_p8(C.poly8x8_t(v0), C.poly8x8x3_t(v1), C.uint8x8_t(v2)))
}

func Vtbx3U8(v0 Uint8X8, v1 Uint8X8X3, v2 Uint8X8) Uint8X8 {
	return Uint8X8(C.vtbx3_u8(C.uint8x8_t(v0), C.uint8x8x3_t(v1), C.uint8x8_t(v2)))
}

func Vtbx3S8(v0 Int8X8, v1 Int8X8X3, v2 Int8X8) Int8X8 {
	return Int8X8(C.vtbx3_s8(C.int8x8_t(v0), C.int8x8x3_t(v1), C.int8x8_t(v2)))
}

func Vtbx4P8(v0 Poly8X8, v1 Poly8X8X4, v2 Uint8X8) Poly8X8 {
	return Poly8X8(C.vtbx4_p8(C.poly8x8_t(v0), C.poly8x8x4_t(v1), C.uint8x8_t(v2)))
}

func Vtbx4U8(v0 Uint8X8, v1 Uint8X8X4, v2 Uint8X8) Uint8X8 {
	return Uint8X8(C.vtbx4_u8(C.uint8x8_t(v0), C.uint8x8x4_t(v1), C.uint8x8_t(v2)))
}

func Vtbx4S8(v0 Int8X8, v1 Int8X8X4, v2 Int8X8) Int8X8 {
	return Int8X8(C.vtbx4_s8(C.int8x8_t(v0), C.int8x8x4_t(v1), C.int8x8_t(v2)))
}

func VtrnP8(v0 Poly8X8, v1 Poly8X8) Poly8X8X2 {
	return Poly8X8X2(C.vtrn_p8(C.poly8x8_t(v0), C.poly8x8_t(v1)))
}

func VtrnP16(v0 Poly16X4, v1 Poly16X4) Poly16X4X2 {
	return Poly16X4X2(C.vtrn_p16(C.poly16x4_t(v0), C.poly16x4_t(v1)))
}

func VtrnqP8(v0 Poly8X16, v1 Poly8X16) Poly8X16X2 {
	return Poly8X16X2(C.vtrnq_p8(C.poly8x16_t(v0), C.poly8x16_t(v1)))
}

func VtrnqP16(v0 Poly16X8, v1 Poly16X8) Poly16X8X2 {
	return Poly16X8X2(C.vtrnq_p16(C.poly16x8_t(v0), C.poly16x8_t(v1)))
}

func VtrnqU8(v0 Uint8X16, v1 Uint8X16) Uint8X16X2 {
	return Uint8X16X2(C.vtrnq_u8(C.uint8x16_t(v0), C.uint8x16_t(v1)))
}

func VtrnqU32(v0 Uint32X4, v1 Uint32X4) Uint32X4X2 {
	return Uint32X4X2(C.vtrnq_u32(C.uint32x4_t(v0), C.uint32x4_t(v1)))
}

func VtrnqU16(v0 Uint16X8, v1 Uint16X8) Uint16X8X2 {
	return Uint16X8X2(C.vtrnq_u16(C.uint16x8_t(v0), C.uint16x8_t(v1)))
}

func VtrnqS8(v0 Int8X16, v1 Int8X16) Int8X16X2 {
	return Int8X16X2(C.vtrnq_s8(C.int8x16_t(v0), C.int8x16_t(v1)))
}

func VtrnqF32(v0 Float32X4, v1 Float32X4) Float32X4X2 {
	return Float32X4X2(C.vtrnq_f32(C.float32x4_t(v0), C.float32x4_t(v1)))
}

func VtrnqS32(v0 Int32X4, v1 Int32X4) Int32X4X2 {
	return Int32X4X2(C.vtrnq_s32(C.int32x4_t(v0), C.int32x4_t(v1)))
}

func VtrnqS16(v0 Int16X8, v1 Int16X8) Int16X8X2 {
	return Int16X8X2(C.vtrnq_s16(C.int16x8_t(v0), C.int16x8_t(v1)))
}

func VtrnU8(v0 Uint8X8, v1 Uint8X8) Uint8X8X2 {
	return Uint8X8X2(C.vtrn_u8(C.uint8x8_t(v0), C.uint8x8_t(v1)))
}

func VtrnU32(v0 Uint32X2, v1 Uint32X2) Uint32X2X2 {
	return Uint32X2X2(C.vtrn_u32(C.uint32x2_t(v0), C.uint32x2_t(v1)))
}

func VtrnU16(v0 Uint16X4, v1 Uint16X4) Uint16X4X2 {
	return Uint16X4X2(C.vtrn_u16(C.uint16x4_t(v0), C.uint16x4_t(v1)))
}

func VtrnS8(v0 Int8X8, v1 Int8X8) Int8X8X2 {
	return Int8X8X2(C.vtrn_s8(C.int8x8_t(v0), C.int8x8_t(v1)))
}

func VtrnF32(v0 Float32X2, v1 Float32X2) Float32X2X2 {
	return Float32X2X2(C.vtrn_f32(C.float32x2_t(v0), C.float32x2_t(v1)))
}

func VtrnS32(v0 Int32X2, v1 Int32X2) Int32X2X2 {
	return Int32X2X2(C.vtrn_s32(C.int32x2_t(v0), C.int32x2_t(v1)))
}

func VtrnS16(v0 Int16X4, v1 Int16X4) Int16X4X2 {
	return Int16X4X2(C.vtrn_s16(C.int16x4_t(v0), C.int16x4_t(v1)))
}

func VtstP8(v0 Poly8X8, v1 Poly8X8) Uint8X8 {
	return Uint8X8(C.vtst_p8(C.poly8x8_t(v0), C.poly8x8_t(v1)))
}

func VtstP16(v0 Poly16X4, v1 Poly16X4) Uint16X4 {
	return Uint16X4(C.vtst_p16(C.poly16x4_t(v0), C.poly16x4_t(v1)))
}

func VtstqP8(v0 Poly8X16, v1 Poly8X16) Uint8X16 {
	return Uint8X16(C.vtstq_p8(C.poly8x16_t(v0), C.poly8x16_t(v1)))
}

func VtstqP16(v0 Poly16X8, v1 Poly16X8) Uint16X8 {
	return Uint16X8(C.vtstq_p16(C.poly16x8_t(v0), C.poly16x8_t(v1)))
}

func VtstqU8(v0 Uint8X16, v1 Uint8X16) Uint8X16 {
	return Uint8X16(C.vtstq_u8(C.uint8x16_t(v0), C.uint8x16_t(v1)))
}

func VtstqU32(v0 Uint32X4, v1 Uint32X4) Uint32X4 {
	return Uint32X4(C.vtstq_u32(C.uint32x4_t(v0), C.uint32x4_t(v1)))
}

func VtstqU16(v0 Uint16X8, v1 Uint16X8) Uint16X8 {
	return Uint16X8(C.vtstq_u16(C.uint16x8_t(v0), C.uint16x8_t(v1)))
}

func VtstqS8(v0 Int8X16, v1 Int8X16) Uint8X16 {
	return Uint8X16(C.vtstq_s8(C.int8x16_t(v0), C.int8x16_t(v1)))
}

func VtstqS32(v0 Int32X4, v1 Int32X4) Uint32X4 {
	return Uint32X4(C.vtstq_s32(C.int32x4_t(v0), C.int32x4_t(v1)))
}

func VtstqS16(v0 Int16X8, v1 Int16X8) Uint16X8 {
	return Uint16X8(C.vtstq_s16(C.int16x8_t(v0), C.int16x8_t(v1)))
}

func VtstU8(v0 Uint8X8, v1 Uint8X8) Uint8X8 {
	return Uint8X8(C.vtst_u8(C.uint8x8_t(v0), C.uint8x8_t(v1)))
}

func VtstU32(v0 Uint32X2, v1 Uint32X2) Uint32X2 {
	return Uint32X2(C.vtst_u32(C.uint32x2_t(v0), C.uint32x2_t(v1)))
}

func VtstU16(v0 Uint16X4, v1 Uint16X4) Uint16X4 {
	return Uint16X4(C.vtst_u16(C.uint16x4_t(v0), C.uint16x4_t(v1)))
}

func VtstS8(v0 Int8X8, v1 Int8X8) Uint8X8 {
	return Uint8X8(C.vtst_s8(C.int8x8_t(v0), C.int8x8_t(v1)))
}

func VtstS32(v0 Int32X2, v1 Int32X2) Uint32X2 {
	return Uint32X2(C.vtst_s32(C.int32x2_t(v0), C.int32x2_t(v1)))
}

func VtstS16(v0 Int16X4, v1 Int16X4) Uint16X4 {
	return Uint16X4(C.vtst_s16(C.int16x4_t(v0), C.int16x4_t(v1)))
}

func VuzpP8(v0 Poly8X8, v1 Poly8X8) Poly8X8X2 {
	return Poly8X8X2(C.vuzp_p8(C.poly8x8_t(v0), C.poly8x8_t(v1)))
}

func VuzpP16(v0 Poly16X4, v1 Poly16X4) Poly16X4X2 {
	return Poly16X4X2(C.vuzp_p16(C.poly16x4_t(v0), C.poly16x4_t(v1)))
}

func VuzpqP8(v0 Poly8X16, v1 Poly8X16) Poly8X16X2 {
	return Poly8X16X2(C.vuzpq_p8(C.poly8x16_t(v0), C.poly8x16_t(v1)))
}

func VuzpqP16(v0 Poly16X8, v1 Poly16X8) Poly16X8X2 {
	return Poly16X8X2(C.vuzpq_p16(C.poly16x8_t(v0), C.poly16x8_t(v1)))
}

func VuzpqU8(v0 Uint8X16, v1 Uint8X16) Uint8X16X2 {
	return Uint8X16X2(C.vuzpq_u8(C.uint8x16_t(v0), C.uint8x16_t(v1)))
}

func VuzpqU32(v0 Uint32X4, v1 Uint32X4) Uint32X4X2 {
	return Uint32X4X2(C.vuzpq_u32(C.uint32x4_t(v0), C.uint32x4_t(v1)))
}

func VuzpqU16(v0 Uint16X8, v1 Uint16X8) Uint16X8X2 {
	return Uint16X8X2(C.vuzpq_u16(C.uint16x8_t(v0), C.uint16x8_t(v1)))
}

func VuzpqS8(v0 Int8X16, v1 Int8X16) Int8X16X2 {
	return Int8X16X2(C.vuzpq_s8(C.int8x16_t(v0), C.int8x16_t(v1)))
}

func VuzpqF32(v0 Float32X4, v1 Float32X4) Float32X4X2 {
	return Float32X4X2(C.vuzpq_f32(C.float32x4_t(v0), C.float32x4_t(v1)))
}

func VuzpqS32(v0 Int32X4, v1 Int32X4) Int32X4X2 {
	return Int32X4X2(C.vuzpq_s32(C.int32x4_t(v0), C.int32x4_t(v1)))
}

func VuzpqS16(v0 Int16X8, v1 Int16X8) Int16X8X2 {
	return Int16X8X2(C.vuzpq_s16(C.int16x8_t(v0), C.int16x8_t(v1)))
}

func VuzpU8(v0 Uint8X8, v1 Uint8X8) Uint8X8X2 {
	return Uint8X8X2(C.vuzp_u8(C.uint8x8_t(v0), C.uint8x8_t(v1)))
}

func VuzpU32(v0 Uint32X2, v1 Uint32X2) Uint32X2X2 {
	return Uint32X2X2(C.vuzp_u32(C.uint32x2_t(v0), C.uint32x2_t(v1)))
}

func VuzpU16(v0 Uint16X4, v1 Uint16X4) Uint16X4X2 {
	return Uint16X4X2(C.vuzp_u16(C.uint16x4_t(v0), C.uint16x4_t(v1)))
}

func VuzpS8(v0 Int8X8, v1 Int8X8) Int8X8X2 {
	return Int8X8X2(C.vuzp_s8(C.int8x8_t(v0), C.int8x8_t(v1)))
}

func VuzpF32(v0 Float32X2, v1 Float32X2) Float32X2X2 {
	return Float32X2X2(C.vuzp_f32(C.float32x2_t(v0), C.float32x2_t(v1)))
}

func VuzpS32(v0 Int32X2, v1 Int32X2) Int32X2X2 {
	return Int32X2X2(C.vuzp_s32(C.int32x2_t(v0), C.int32x2_t(v1)))
}

func VuzpS16(v0 Int16X4, v1 Int16X4) Int16X4X2 {
	return Int16X4X2(C.vuzp_s16(C.int16x4_t(v0), C.int16x4_t(v1)))
}

func VzipP8(v0 Poly8X8, v1 Poly8X8) Poly8X8X2 {
	return Poly8X8X2(C.vzip_p8(C.poly8x8_t(v0), C.poly8x8_t(v1)))
}

func VzipP16(v0 Poly16X4, v1 Poly16X4) Poly16X4X2 {
	return Poly16X4X2(C.vzip_p16(C.poly16x4_t(v0), C.poly16x4_t(v1)))
}

func VzipqP8(v0 Poly8X16, v1 Poly8X16) Poly8X16X2 {
	return Poly8X16X2(C.vzipq_p8(C.poly8x16_t(v0), C.poly8x16_t(v1)))
}

func VzipqP16(v0 Poly16X8, v1 Poly16X8) Poly16X8X2 {
	return Poly16X8X2(C.vzipq_p16(C.poly16x8_t(v0), C.poly16x8_t(v1)))
}

func VzipqU8(v0 Uint8X16, v1 Uint8X16) Uint8X16X2 {
	return Uint8X16X2(C.vzipq_u8(C.uint8x16_t(v0), C.uint8x16_t(v1)))
}

func VzipqU32(v0 Uint32X4, v1 Uint32X4) Uint32X4X2 {
	return Uint32X4X2(C.vzipq_u32(C.uint32x4_t(v0), C.uint32x4_t(v1)))
}

func VzipqU16(v0 Uint16X8, v1 Uint16X8) Uint16X8X2 {
	return Uint16X8X2(C.vzipq_u16(C.uint16x8_t(v0), C.uint16x8_t(v1)))
}

func VzipqS8(v0 Int8X16, v1 Int8X16) Int8X16X2 {
	return Int8X16X2(C.vzipq_s8(C.int8x16_t(v0), C.int8x16_t(v1)))
}

func VzipqF32(v0 Float32X4, v1 Float32X4) Float32X4X2 {
	return Float32X4X2(C.vzipq_f32(C.float32x4_t(v0), C.float32x4_t(v1)))
}

func VzipqS32(v0 Int32X4, v1 Int32X4) Int32X4X2 {
	return Int32X4X2(C.vzipq_s32(C.int32x4_t(v0), C.int32x4_t(v1)))
}

func VzipqS16(v0 Int16X8, v1 Int16X8) Int16X8X2 {
	return Int16X8X2(C.vzipq_s16(C.int16x8_t(v0), C.int16x8_t(v1)))
}

func VzipU8(v0 Uint8X8, v1 Uint8X8) Uint8X8X2 {
	return Uint8X8X2(C.vzip_u8(C.uint8x8_t(v0), C.uint8x8_t(v1)))
}

func VzipU32(v0 Uint32X2, v1 Uint32X2) Uint32X2X2 {
	return Uint32X2X2(C.vzip_u32(C.uint32x2_t(v0), C.uint32x2_t(v1)))
}

func VzipU16(v0 Uint16X4, v1 Uint16X4) Uint16X4X2 {
	return Uint16X4X2(C.vzip_u16(C.uint16x4_t(v0), C.uint16x4_t(v1)))
}

func VzipS8(v0 Int8X8, v1 Int8X8) Int8X8X2 {
	return Int8X8X2(C.vzip_s8(C.int8x8_t(v0), C.int8x8_t(v1)))
}

func VzipF32(v0 Float32X2, v1 Float32X2) Float32X2X2 {
	return Float32X2X2(C.vzip_f32(C.float32x2_t(v0), C.float32x2_t(v1)))
}

func VzipS32(v0 Int32X2, v1 Int32X2) Int32X2X2 {
	return Int32X2X2(C.vzip_s32(C.int32x2_t(v0), C.int32x2_t(v1)))
}

func VzipS16(v0 Int16X4, v1 Int16X4) Int16X4X2 {
	return Int16X4X2(C.vzip_s16(C.int16x4_t(v0), C.int16x4_t(v1)))
}

func VcvtaqS32F32(v0 Float32X4) Int32X4 {
	return Int32X4(C.vcvtaq_s32_f32(C.float32x4_t(v0)))
}

func VcvtaS32F32(v0 Float32X2) Int32X2 {
	return Int32X2(C.vcvta_s32_f32(C.float32x2_t(v0)))
}

func VcvtaqU32F32(v0 Float32X4) Uint32X4 {
	return Uint32X4(C.vcvtaq_u32_f32(C.float32x4_t(v0)))
}

func VcvtaU32F32(v0 Float32X2) Uint32X2 {
	return Uint32X2(C.vcvta_u32_f32(C.float32x2_t(v0)))
}

func VcvtmqS32F32(v0 Float32X4) Int32X4 {
	return Int32X4(C.vcvtmq_s32_f32(C.float32x4_t(v0)))
}

func VcvtmS32F32(v0 Float32X2) Int32X2 {
	return Int32X2(C.vcvtm_s32_f32(C.float32x2_t(v0)))
}

func VcvtmqU32F32(v0 Float32X4) Uint32X4 {
	return Uint32X4(C.vcvtmq_u32_f32(C.float32x4_t(v0)))
}

func VcvtmU32F32(v0 Float32X2) Uint32X2 {
	return Uint32X2(C.vcvtm_u32_f32(C.float32x2_t(v0)))
}

func VcvtnqS32F32(v0 Float32X4) Int32X4 {
	return Int32X4(C.vcvtnq_s32_f32(C.float32x4_t(v0)))
}

func VcvtnS32F32(v0 Float32X2) Int32X2 {
	return Int32X2(C.vcvtn_s32_f32(C.float32x2_t(v0)))
}

func VcvtnqU32F32(v0 Float32X4) Uint32X4 {
	return Uint32X4(C.vcvtnq_u32_f32(C.float32x4_t(v0)))
}

func VcvtnU32F32(v0 Float32X2) Uint32X2 {
	return Uint32X2(C.vcvtn_u32_f32(C.float32x2_t(v0)))
}

func VcvtpqS32F32(v0 Float32X4) Int32X4 {
	return Int32X4(C.vcvtpq_s32_f32(C.float32x4_t(v0)))
}

func VcvtpS32F32(v0 Float32X2) Int32X2 {
	return Int32X2(C.vcvtp_s32_f32(C.float32x2_t(v0)))
}

func VcvtpqU32F32(v0 Float32X4) Uint32X4 {
	return Uint32X4(C.vcvtpq_u32_f32(C.float32x4_t(v0)))
}

func VcvtpU32F32(v0 Float32X2) Uint32X2 {
	return Uint32X2(C.vcvtp_u32_f32(C.float32x2_t(v0)))
}

func VaesdqU8(v0 Uint8X16, v1 Uint8X16) Uint8X16 {
	return Uint8X16(C.vaesdq_u8(C.uint8x16_t(v0), C.uint8x16_t(v1)))
}

func VaeseqU8(v0 Uint8X16, v1 Uint8X16) Uint8X16 {
	return Uint8X16(C.vaeseq_u8(C.uint8x16_t(v0), C.uint8x16_t(v1)))
}

func VaesimcqU8(v0 Uint8X16) Uint8X16 {
	return Uint8X16(C.vaesimcq_u8(C.uint8x16_t(v0)))
}

func VaesmcqU8(v0 Uint8X16) Uint8X16 {
	return Uint8X16(C.vaesmcq_u8(C.uint8x16_t(v0)))
}

func VrndqF32(v0 Float32X4) Float32X4 {
	return Float32X4(C.vrndq_f32(C.float32x4_t(v0)))
}

func VrndF32(v0 Float32X2) Float32X2 {
	return Float32X2(C.vrnd_f32(C.float32x2_t(v0)))
}

func VrndaqF32(v0 Float32X4) Float32X4 {
	return Float32X4(C.vrndaq_f32(C.float32x4_t(v0)))
}

func VrndaF32(v0 Float32X2) Float32X2 {
	return Float32X2(C.vrnda_f32(C.float32x2_t(v0)))
}

func VrndiqF32(v0 Float32X4) Float32X4 {
	return Float32X4(C.vrndiq_f32(C.float32x4_t(v0)))
}

func VrndiF32(v0 Float32X2) Float32X2 {
	return Float32X2(C.vrndi_f32(C.float32x2_t(v0)))
}

func VrndmqF32(v0 Float32X4) Float32X4 {
	return Float32X4(C.vrndmq_f32(C.float32x4_t(v0)))
}

func VrndmF32(v0 Float32X2) Float32X2 {
	return Float32X2(C.vrndm_f32(C.float32x2_t(v0)))
}

func VrndnqF32(v0 Float32X4) Float32X4 {
	return Float32X4(C.vrndnq_f32(C.float32x4_t(v0)))
}

func VrndnF32(v0 Float32X2) Float32X2 {
	return Float32X2(C.vrndn_f32(C.float32x2_t(v0)))
}

func VrndnsF32(v0 float32) float32 {
	return float32(C.vrndns_f32(C.float32_t(v0)))
}

func VrndpqF32(v0 Float32X4) Float32X4 {
	return Float32X4(C.vrndpq_f32(C.float32x4_t(v0)))
}

func VrndpF32(v0 Float32X2) Float32X2 {
	return Float32X2(C.vrndp_f32(C.float32x2_t(v0)))
}

func VrndxqF32(v0 Float32X4) Float32X4 {
	return Float32X4(C.vrndxq_f32(C.float32x4_t(v0)))
}

func VrndxF32(v0 Float32X2) Float32X2 {
	return Float32X2(C.vrndx_f32(C.float32x2_t(v0)))
}

func VmaxnmqF32(v0 Float32X4, v1 Float32X4) Float32X4 {
	return Float32X4(C.vmaxnmq_f32(C.float32x4_t(v0), C.float32x4_t(v1)))
}

func VmaxnmF32(v0 Float32X2, v1 Float32X2) Float32X2 {
	return Float32X2(C.vmaxnm_f32(C.float32x2_t(v0), C.float32x2_t(v1)))
}

func VminnmqF32(v0 Float32X4, v1 Float32X4) Float32X4 {
	return Float32X4(C.vminnmq_f32(C.float32x4_t(v0), C.float32x4_t(v1)))
}

func VminnmF32(v0 Float32X2, v1 Float32X2) Float32X2 {
	return Float32X2(C.vminnm_f32(C.float32x2_t(v0), C.float32x2_t(v1)))
}

func Vsha1CqU32(v0 Uint32X4, v1 uint32, v2 Uint32X4) Uint32X4 {
	return Uint32X4(C.vsha1cq_u32(C.uint32x4_t(v0), C.uint32_t(v1), C.uint32x4_t(v2)))
}

func Vsha1HU32(v0 uint32) uint32 {
	return uint32(C.vsha1h_u32(C.uint32_t(v0)))
}

func Vsha1MqU32(v0 Uint32X4, v1 uint32, v2 Uint32X4) Uint32X4 {
	return Uint32X4(C.vsha1mq_u32(C.uint32x4_t(v0), C.uint32_t(v1), C.uint32x4_t(v2)))
}

func Vsha1PqU32(v0 Uint32X4, v1 uint32, v2 Uint32X4) Uint32X4 {
	return Uint32X4(C.vsha1pq_u32(C.uint32x4_t(v0), C.uint32_t(v1), C.uint32x4_t(v2)))
}

func Vsha1Su0QU32(v0 Uint32X4, v1 Uint32X4, v2 Uint32X4) Uint32X4 {
	return Uint32X4(C.vsha1su0q_u32(C.uint32x4_t(v0), C.uint32x4_t(v1), C.uint32x4_t(v2)))
}

func Vsha1Su1QU32(v0 Uint32X4, v1 Uint32X4) Uint32X4 {
	return Uint32X4(C.vsha1su1q_u32(C.uint32x4_t(v0), C.uint32x4_t(v1)))
}

func Vsha256HqU32(v0 Uint32X4, v1 Uint32X4, v2 Uint32X4) Uint32X4 {
	return Uint32X4(C.vsha256hq_u32(C.uint32x4_t(v0), C.uint32x4_t(v1), C.uint32x4_t(v2)))
}

func Vsha256H2QU32(v0 Uint32X4, v1 Uint32X4, v2 Uint32X4) Uint32X4 {
	return Uint32X4(C.vsha256h2q_u32(C.uint32x4_t(v0), C.uint32x4_t(v1), C.uint32x4_t(v2)))
}

func Vsha256Su0QU32(v0 Uint32X4, v1 Uint32X4) Uint32X4 {
	return Uint32X4(C.vsha256su0q_u32(C.uint32x4_t(v0), C.uint32x4_t(v1)))
}

func Vsha256Su1QU32(v0 Uint32X4, v1 Uint32X4, v2 Uint32X4) Uint32X4 {
	return Uint32X4(C.vsha256su1q_u32(C.uint32x4_t(v0), C.uint32x4_t(v1), C.uint32x4_t(v2)))
}

func VbcaxqU8(v0 Uint8X16, v1 Uint8X16, v2 Uint8X16) Uint8X16 {
	return Uint8X16(C.vbcaxq_u8(C.uint8x16_t(v0), C.uint8x16_t(v1), C.uint8x16_t(v2)))
}

func VbcaxqU32(v0 Uint32X4, v1 Uint32X4, v2 Uint32X4) Uint32X4 {
	return Uint32X4(C.vbcaxq_u32(C.uint32x4_t(v0), C.uint32x4_t(v1), C.uint32x4_t(v2)))
}

func VbcaxqU64(v0 Uint64X2, v1 Uint64X2, v2 Uint64X2) Uint64X2 {
	return Uint64X2(C.vbcaxq_u64(C.uint64x2_t(v0), C.uint64x2_t(v1), C.uint64x2_t(v2)))
}

func VbcaxqU16(v0 Uint16X8, v1 Uint16X8, v2 Uint16X8) Uint16X8 {
	return Uint16X8(C.vbcaxq_u16(C.uint16x8_t(v0), C.uint16x8_t(v1), C.uint16x8_t(v2)))
}

func VbcaxqS8(v0 Int8X16, v1 Int8X16, v2 Int8X16) Int8X16 {
	return Int8X16(C.vbcaxq_s8(C.int8x16_t(v0), C.int8x16_t(v1), C.int8x16_t(v2)))
}

func VbcaxqS32(v0 Int32X4, v1 Int32X4, v2 Int32X4) Int32X4 {
	return Int32X4(C.vbcaxq_s32(C.int32x4_t(v0), C.int32x4_t(v1), C.int32x4_t(v2)))
}

func VbcaxqS64(v0 Int64X2, v1 Int64X2, v2 Int64X2) Int64X2 {
	return Int64X2(C.vbcaxq_s64(C.int64x2_t(v0), C.int64x2_t(v1), C.int64x2_t(v2)))
}

func VbcaxqS16(v0 Int16X8, v1 Int16X8, v2 Int16X8) Int16X8 {
	return Int16X8(C.vbcaxq_s16(C.int16x8_t(v0), C.int16x8_t(v1), C.int16x8_t(v2)))
}

func Veor3QU8(v0 Uint8X16, v1 Uint8X16, v2 Uint8X16) Uint8X16 {
	return Uint8X16(C.veor3q_u8(C.uint8x16_t(v0), C.uint8x16_t(v1), C.uint8x16_t(v2)))
}

func Veor3QU32(v0 Uint32X4, v1 Uint32X4, v2 Uint32X4) Uint32X4 {
	return Uint32X4(C.veor3q_u32(C.uint32x4_t(v0), C.uint32x4_t(v1), C.uint32x4_t(v2)))
}

func Veor3QU64(v0 Uint64X2, v1 Uint64X2, v2 Uint64X2) Uint64X2 {
	return Uint64X2(C.veor3q_u64(C.uint64x2_t(v0), C.uint64x2_t(v1), C.uint64x2_t(v2)))
}

func Veor3QU16(v0 Uint16X8, v1 Uint16X8, v2 Uint16X8) Uint16X8 {
	return Uint16X8(C.veor3q_u16(C.uint16x8_t(v0), C.uint16x8_t(v1), C.uint16x8_t(v2)))
}

func Veor3QS8(v0 Int8X16, v1 Int8X16, v2 Int8X16) Int8X16 {
	return Int8X16(C.veor3q_s8(C.int8x16_t(v0), C.int8x16_t(v1), C.int8x16_t(v2)))
}

func Veor3QS32(v0 Int32X4, v1 Int32X4, v2 Int32X4) Int32X4 {
	return Int32X4(C.veor3q_s32(C.int32x4_t(v0), C.int32x4_t(v1), C.int32x4_t(v2)))
}

func Veor3QS64(v0 Int64X2, v1 Int64X2, v2 Int64X2) Int64X2 {
	return Int64X2(C.veor3q_s64(C.int64x2_t(v0), C.int64x2_t(v1), C.int64x2_t(v2)))
}

func Veor3QS16(v0 Int16X8, v1 Int16X8, v2 Int16X8) Int16X8 {
	return Int16X8(C.veor3q_s16(C.int16x8_t(v0), C.int16x8_t(v1), C.int16x8_t(v2)))
}

func Vrax1QU64(v0 Uint64X2, v1 Uint64X2) Uint64X2 {
	return Uint64X2(C.vrax1q_u64(C.uint64x2_t(v0), C.uint64x2_t(v1)))
}

func Vsha512HqU64(v0 Uint64X2, v1 Uint64X2, v2 Uint64X2) Uint64X2 {
	return Uint64X2(C.vsha512hq_u64(C.uint64x2_t(v0), C.uint64x2_t(v1), C.uint64x2_t(v2)))
}

func Vsha512H2QU64(v0 Uint64X2, v1 Uint64X2, v2 Uint64X2) Uint64X2 {
	return Uint64X2(C.vsha512h2q_u64(C.uint64x2_t(v0), C.uint64x2_t(v1), C.uint64x2_t(v2)))
}

func Vsha512Su0QU64(v0 Uint64X2, v1 Uint64X2) Uint64X2 {
	return Uint64X2(C.vsha512su0q_u64(C.uint64x2_t(v0), C.uint64x2_t(v1)))
}

func Vsha512Su1QU64(v0 Uint64X2, v1 Uint64X2, v2 Uint64X2) Uint64X2 {
	return Uint64X2(C.vsha512su1q_u64(C.uint64x2_t(v0), C.uint64x2_t(v1), C.uint64x2_t(v2)))
}

func Vsm3Partw1QU32(v0 Uint32X4, v1 Uint32X4, v2 Uint32X4) Uint32X4 {
	return Uint32X4(C.vsm3partw1q_u32(C.uint32x4_t(v0), C.uint32x4_t(v1), C.uint32x4_t(v2)))
}

func Vsm3Partw2QU32(v0 Uint32X4, v1 Uint32X4, v2 Uint32X4) Uint32X4 {
	return Uint32X4(C.vsm3partw2q_u32(C.uint32x4_t(v0), C.uint32x4_t(v1), C.uint32x4_t(v2)))
}

func Vsm3Ss1QU32(v0 Uint32X4, v1 Uint32X4, v2 Uint32X4) Uint32X4 {
	return Uint32X4(C.vsm3ss1q_u32(C.uint32x4_t(v0), C.uint32x4_t(v1), C.uint32x4_t(v2)))
}

func Vsm4EqU32(v0 Uint32X4, v1 Uint32X4) Uint32X4 {
	return Uint32X4(C.vsm4eq_u32(C.uint32x4_t(v0), C.uint32x4_t(v1)))
}

func Vsm4EkeyqU32(v0 Uint32X4, v1 Uint32X4) Uint32X4 {
	return Uint32X4(C.vsm4ekeyq_u32(C.uint32x4_t(v0), C.uint32x4_t(v1)))
}

func VcvtaqS64F64(v0 Float64X2) Int64X2 {
	return Int64X2(C.vcvtaq_s64_f64(C.float64x2_t(v0)))
}

func VcvtaS64F64(v0 Float64X1) Int64X1 {
	return Int64X1(C.vcvta_s64_f64(C.float64x1_t(v0)))
}

func VcvtaqU64F64(v0 Float64X2) Uint64X2 {
	return Uint64X2(C.vcvtaq_u64_f64(C.float64x2_t(v0)))
}

func VcvtaU64F64(v0 Float64X1) Uint64X1 {
	return Uint64X1(C.vcvta_u64_f64(C.float64x1_t(v0)))
}

func VcvtmqS64F64(v0 Float64X2) Int64X2 {
	return Int64X2(C.vcvtmq_s64_f64(C.float64x2_t(v0)))
}

func VcvtmS64F64(v0 Float64X1) Int64X1 {
	return Int64X1(C.vcvtm_s64_f64(C.float64x1_t(v0)))
}

func VcvtmqU64F64(v0 Float64X2) Uint64X2 {
	return Uint64X2(C.vcvtmq_u64_f64(C.float64x2_t(v0)))
}

func VcvtmU64F64(v0 Float64X1) Uint64X1 {
	return Uint64X1(C.vcvtm_u64_f64(C.float64x1_t(v0)))
}

func VcvtnqS64F64(v0 Float64X2) Int64X2 {
	return Int64X2(C.vcvtnq_s64_f64(C.float64x2_t(v0)))
}

func VcvtnS64F64(v0 Float64X1) Int64X1 {
	return Int64X1(C.vcvtn_s64_f64(C.float64x1_t(v0)))
}

func VcvtnqU64F64(v0 Float64X2) Uint64X2 {
	return Uint64X2(C.vcvtnq_u64_f64(C.float64x2_t(v0)))
}

func VcvtnU64F64(v0 Float64X1) Uint64X1 {
	return Uint64X1(C.vcvtn_u64_f64(C.float64x1_t(v0)))
}

func VcvtpqS64F64(v0 Float64X2) Int64X2 {
	return Int64X2(C.vcvtpq_s64_f64(C.float64x2_t(v0)))
}

func VcvtpS64F64(v0 Float64X1) Int64X1 {
	return Int64X1(C.vcvtp_s64_f64(C.float64x1_t(v0)))
}

func VcvtpqU64F64(v0 Float64X2) Uint64X2 {
	return Uint64X2(C.vcvtpq_u64_f64(C.float64x2_t(v0)))
}

func VcvtpU64F64(v0 Float64X1) Uint64X1 {
	return Uint64X1(C.vcvtp_u64_f64(C.float64x1_t(v0)))
}

func VreinterpretP8P64(v0 Poly64X1) Poly8X8 {
	return Poly8X8(C.vreinterpret_p8_p64(C.poly64x1_t(v0)))
}

func VreinterpretP8P16(v0 Poly16X4) Poly8X8 {
	return Poly8X8(C.vreinterpret_p8_p16(C.poly16x4_t(v0)))
}

func VreinterpretP8U8(v0 Uint8X8) Poly8X8 {
	return Poly8X8(C.vreinterpret_p8_u8(C.uint8x8_t(v0)))
}

func VreinterpretP8U32(v0 Uint32X2) Poly8X8 {
	return Poly8X8(C.vreinterpret_p8_u32(C.uint32x2_t(v0)))
}

func VreinterpretP8U64(v0 Uint64X1) Poly8X8 {
	return Poly8X8(C.vreinterpret_p8_u64(C.uint64x1_t(v0)))
}

func VreinterpretP8U16(v0 Uint16X4) Poly8X8 {
	return Poly8X8(C.vreinterpret_p8_u16(C.uint16x4_t(v0)))
}

func VreinterpretP8S8(v0 Int8X8) Poly8X8 {
	return Poly8X8(C.vreinterpret_p8_s8(C.int8x8_t(v0)))
}

func VreinterpretP8F64(v0 Float64X1) Poly8X8 {
	return Poly8X8(C.vreinterpret_p8_f64(C.float64x1_t(v0)))
}

func VreinterpretP8F32(v0 Float32X2) Poly8X8 {
	return Poly8X8(C.vreinterpret_p8_f32(C.float32x2_t(v0)))
}

func VreinterpretP8S32(v0 Int32X2) Poly8X8 {
	return Poly8X8(C.vreinterpret_p8_s32(C.int32x2_t(v0)))
}

func VreinterpretP8S64(v0 Int64X1) Poly8X8 {
	return Poly8X8(C.vreinterpret_p8_s64(C.int64x1_t(v0)))
}

func VreinterpretP8S16(v0 Int16X4) Poly8X8 {
	return Poly8X8(C.vreinterpret_p8_s16(C.int16x4_t(v0)))
}

func VreinterpretP64P8(v0 Poly8X8) Poly64X1 {
	return Poly64X1(C.vreinterpret_p64_p8(C.poly8x8_t(v0)))
}

func VreinterpretP64P16(v0 Poly16X4) Poly64X1 {
	return Poly64X1(C.vreinterpret_p64_p16(C.poly16x4_t(v0)))
}

func VreinterpretP64U8(v0 Uint8X8) Poly64X1 {
	return Poly64X1(C.vreinterpret_p64_u8(C.uint8x8_t(v0)))
}

func VreinterpretP64U32(v0 Uint32X2) Poly64X1 {
	return Poly64X1(C.vreinterpret_p64_u32(C.uint32x2_t(v0)))
}

func VreinterpretP64U64(v0 Uint64X1) Poly64X1 {
	return Poly64X1(C.vreinterpret_p64_u64(C.uint64x1_t(v0)))
}

func VreinterpretP64U16(v0 Uint16X4) Poly64X1 {
	return Poly64X1(C.vreinterpret_p64_u16(C.uint16x4_t(v0)))
}

func VreinterpretP64S8(v0 Int8X8) Poly64X1 {
	return Poly64X1(C.vreinterpret_p64_s8(C.int8x8_t(v0)))
}

func VreinterpretP64F64(v0 Float64X1) Poly64X1 {
	return Poly64X1(C.vreinterpret_p64_f64(C.float64x1_t(v0)))
}

func VreinterpretP64F32(v0 Float32X2) Poly64X1 {
	return Poly64X1(C.vreinterpret_p64_f32(C.float32x2_t(v0)))
}

func VreinterpretP64S32(v0 Int32X2) Poly64X1 {
	return Poly64X1(C.vreinterpret_p64_s32(C.int32x2_t(v0)))
}

func VreinterpretP64S64(v0 Int64X1) Poly64X1 {
	return Poly64X1(C.vreinterpret_p64_s64(C.int64x1_t(v0)))
}

func VreinterpretP64S16(v0 Int16X4) Poly64X1 {
	return Poly64X1(C.vreinterpret_p64_s16(C.int16x4_t(v0)))
}

func VreinterpretP16P8(v0 Poly8X8) Poly16X4 {
	return Poly16X4(C.vreinterpret_p16_p8(C.poly8x8_t(v0)))
}

func VreinterpretP16P64(v0 Poly64X1) Poly16X4 {
	return Poly16X4(C.vreinterpret_p16_p64(C.poly64x1_t(v0)))
}

func VreinterpretP16U8(v0 Uint8X8) Poly16X4 {
	return Poly16X4(C.vreinterpret_p16_u8(C.uint8x8_t(v0)))
}

func VreinterpretP16U32(v0 Uint32X2) Poly16X4 {
	return Poly16X4(C.vreinterpret_p16_u32(C.uint32x2_t(v0)))
}

func VreinterpretP16U64(v0 Uint64X1) Poly16X4 {
	return Poly16X4(C.vreinterpret_p16_u64(C.uint64x1_t(v0)))
}

func VreinterpretP16U16(v0 Uint16X4) Poly16X4 {
	return Poly16X4(C.vreinterpret_p16_u16(C.uint16x4_t(v0)))
}

func VreinterpretP16S8(v0 Int8X8) Poly16X4 {
	return Poly16X4(C.vreinterpret_p16_s8(C.int8x8_t(v0)))
}

func VreinterpretP16F64(v0 Float64X1) Poly16X4 {
	return Poly16X4(C.vreinterpret_p16_f64(C.float64x1_t(v0)))
}

func VreinterpretP16F32(v0 Float32X2) Poly16X4 {
	return Poly16X4(C.vreinterpret_p16_f32(C.float32x2_t(v0)))
}

func VreinterpretP16S32(v0 Int32X2) Poly16X4 {
	return Poly16X4(C.vreinterpret_p16_s32(C.int32x2_t(v0)))
}

func VreinterpretP16S64(v0 Int64X1) Poly16X4 {
	return Poly16X4(C.vreinterpret_p16_s64(C.int64x1_t(v0)))
}

func VreinterpretP16S16(v0 Int16X4) Poly16X4 {
	return Poly16X4(C.vreinterpret_p16_s16(C.int16x4_t(v0)))
}

func VreinterpretqP8P128(v0 Poly128) Poly8X16 {
	return Poly8X16(C.vreinterpretq_p8_p128(C.poly128_t(v0)))
}

func VreinterpretqP8P64(v0 Poly64X2) Poly8X16 {
	return Poly8X16(C.vreinterpretq_p8_p64(C.poly64x2_t(v0)))
}

func VreinterpretqP8P16(v0 Poly16X8) Poly8X16 {
	return Poly8X16(C.vreinterpretq_p8_p16(C.poly16x8_t(v0)))
}

func VreinterpretqP8U8(v0 Uint8X16) Poly8X16 {
	return Poly8X16(C.vreinterpretq_p8_u8(C.uint8x16_t(v0)))
}

func VreinterpretqP8U32(v0 Uint32X4) Poly8X16 {
	return Poly8X16(C.vreinterpretq_p8_u32(C.uint32x4_t(v0)))
}

func VreinterpretqP8U64(v0 Uint64X2) Poly8X16 {
	return Poly8X16(C.vreinterpretq_p8_u64(C.uint64x2_t(v0)))
}

func VreinterpretqP8U16(v0 Uint16X8) Poly8X16 {
	return Poly8X16(C.vreinterpretq_p8_u16(C.uint16x8_t(v0)))
}

func VreinterpretqP8S8(v0 Int8X16) Poly8X16 {
	return Poly8X16(C.vreinterpretq_p8_s8(C.int8x16_t(v0)))
}

func VreinterpretqP8F64(v0 Float64X2) Poly8X16 {
	return Poly8X16(C.vreinterpretq_p8_f64(C.float64x2_t(v0)))
}

func VreinterpretqP8F32(v0 Float32X4) Poly8X16 {
	return Poly8X16(C.vreinterpretq_p8_f32(C.float32x4_t(v0)))
}

func VreinterpretqP8S32(v0 Int32X4) Poly8X16 {
	return Poly8X16(C.vreinterpretq_p8_s32(C.int32x4_t(v0)))
}

func VreinterpretqP8S64(v0 Int64X2) Poly8X16 {
	return Poly8X16(C.vreinterpretq_p8_s64(C.int64x2_t(v0)))
}

func VreinterpretqP8S16(v0 Int16X8) Poly8X16 {
	return Poly8X16(C.vreinterpretq_p8_s16(C.int16x8_t(v0)))
}

func VreinterpretqP128P8(v0 Poly8X16) Poly128 {
	return Poly128(C.vreinterpretq_p128_p8(C.poly8x16_t(v0)))
}

func VreinterpretqP128P64(v0 Poly64X2) Poly128 {
	return Poly128(C.vreinterpretq_p128_p64(C.poly64x2_t(v0)))
}

func VreinterpretqP128P16(v0 Poly16X8) Poly128 {
	return Poly128(C.vreinterpretq_p128_p16(C.poly16x8_t(v0)))
}

func VreinterpretqP128U8(v0 Uint8X16) Poly128 {
	return Poly128(C.vreinterpretq_p128_u8(C.uint8x16_t(v0)))
}

func VreinterpretqP128U32(v0 Uint32X4) Poly128 {
	return Poly128(C.vreinterpretq_p128_u32(C.uint32x4_t(v0)))
}

func VreinterpretqP128U64(v0 Uint64X2) Poly128 {
	return Poly128(C.vreinterpretq_p128_u64(C.uint64x2_t(v0)))
}

func VreinterpretqP128U16(v0 Uint16X8) Poly128 {
	return Poly128(C.vreinterpretq_p128_u16(C.uint16x8_t(v0)))
}

func VreinterpretqP128S8(v0 Int8X16) Poly128 {
	return Poly128(C.vreinterpretq_p128_s8(C.int8x16_t(v0)))
}

func VreinterpretqP128F64(v0 Float64X2) Poly128 {
	return Poly128(C.vreinterpretq_p128_f64(C.float64x2_t(v0)))
}

func VreinterpretqP128F32(v0 Float32X4) Poly128 {
	return Poly128(C.vreinterpretq_p128_f32(C.float32x4_t(v0)))
}

func VreinterpretqP128S32(v0 Int32X4) Poly128 {
	return Poly128(C.vreinterpretq_p128_s32(C.int32x4_t(v0)))
}

func VreinterpretqP128S64(v0 Int64X2) Poly128 {
	return Poly128(C.vreinterpretq_p128_s64(C.int64x2_t(v0)))
}

func VreinterpretqP128S16(v0 Int16X8) Poly128 {
	return Poly128(C.vreinterpretq_p128_s16(C.int16x8_t(v0)))
}

func VreinterpretqP64P8(v0 Poly8X16) Poly64X2 {
	return Poly64X2(C.vreinterpretq_p64_p8(C.poly8x16_t(v0)))
}

func VreinterpretqP64P128(v0 Poly128) Poly64X2 {
	return Poly64X2(C.vreinterpretq_p64_p128(C.poly128_t(v0)))
}

func VreinterpretqP64P16(v0 Poly16X8) Poly64X2 {
	return Poly64X2(C.vreinterpretq_p64_p16(C.poly16x8_t(v0)))
}

func VreinterpretqP64U8(v0 Uint8X16) Poly64X2 {
	return Poly64X2(C.vreinterpretq_p64_u8(C.uint8x16_t(v0)))
}

func VreinterpretqP64U32(v0 Uint32X4) Poly64X2 {
	return Poly64X2(C.vreinterpretq_p64_u32(C.uint32x4_t(v0)))
}

func VreinterpretqP64U64(v0 Uint64X2) Poly64X2 {
	return Poly64X2(C.vreinterpretq_p64_u64(C.uint64x2_t(v0)))
}

func VreinterpretqP64U16(v0 Uint16X8) Poly64X2 {
	return Poly64X2(C.vreinterpretq_p64_u16(C.uint16x8_t(v0)))
}

func VreinterpretqP64S8(v0 Int8X16) Poly64X2 {
	return Poly64X2(C.vreinterpretq_p64_s8(C.int8x16_t(v0)))
}

func VreinterpretqP64F64(v0 Float64X2) Poly64X2 {
	return Poly64X2(C.vreinterpretq_p64_f64(C.float64x2_t(v0)))
}

func VreinterpretqP64F32(v0 Float32X4) Poly64X2 {
	return Poly64X2(C.vreinterpretq_p64_f32(C.float32x4_t(v0)))
}

func VreinterpretqP64S32(v0 Int32X4) Poly64X2 {
	return Poly64X2(C.vreinterpretq_p64_s32(C.int32x4_t(v0)))
}

func VreinterpretqP64S64(v0 Int64X2) Poly64X2 {
	return Poly64X2(C.vreinterpretq_p64_s64(C.int64x2_t(v0)))
}

func VreinterpretqP64S16(v0 Int16X8) Poly64X2 {
	return Poly64X2(C.vreinterpretq_p64_s16(C.int16x8_t(v0)))
}

func VreinterpretqP16P8(v0 Poly8X16) Poly16X8 {
	return Poly16X8(C.vreinterpretq_p16_p8(C.poly8x16_t(v0)))
}

func VreinterpretqP16P128(v0 Poly128) Poly16X8 {
	return Poly16X8(C.vreinterpretq_p16_p128(C.poly128_t(v0)))
}

func VreinterpretqP16P64(v0 Poly64X2) Poly16X8 {
	return Poly16X8(C.vreinterpretq_p16_p64(C.poly64x2_t(v0)))
}

func VreinterpretqP16U8(v0 Uint8X16) Poly16X8 {
	return Poly16X8(C.vreinterpretq_p16_u8(C.uint8x16_t(v0)))
}

func VreinterpretqP16U32(v0 Uint32X4) Poly16X8 {
	return Poly16X8(C.vreinterpretq_p16_u32(C.uint32x4_t(v0)))
}

func VreinterpretqP16U64(v0 Uint64X2) Poly16X8 {
	return Poly16X8(C.vreinterpretq_p16_u64(C.uint64x2_t(v0)))
}

func VreinterpretqP16U16(v0 Uint16X8) Poly16X8 {
	return Poly16X8(C.vreinterpretq_p16_u16(C.uint16x8_t(v0)))
}

func VreinterpretqP16S8(v0 Int8X16) Poly16X8 {
	return Poly16X8(C.vreinterpretq_p16_s8(C.int8x16_t(v0)))
}

func VreinterpretqP16F64(v0 Float64X2) Poly16X8 {
	return Poly16X8(C.vreinterpretq_p16_f64(C.float64x2_t(v0)))
}

func VreinterpretqP16F32(v0 Float32X4) Poly16X8 {
	return Poly16X8(C.vreinterpretq_p16_f32(C.float32x4_t(v0)))
}

func VreinterpretqP16S32(v0 Int32X4) Poly16X8 {
	return Poly16X8(C.vreinterpretq_p16_s32(C.int32x4_t(v0)))
}

func VreinterpretqP16S64(v0 Int64X2) Poly16X8 {
	return Poly16X8(C.vreinterpretq_p16_s64(C.int64x2_t(v0)))
}

func VreinterpretqP16S16(v0 Int16X8) Poly16X8 {
	return Poly16X8(C.vreinterpretq_p16_s16(C.int16x8_t(v0)))
}

func VreinterpretqU8P8(v0 Poly8X16) Uint8X16 {
	return Uint8X16(C.vreinterpretq_u8_p8(C.poly8x16_t(v0)))
}

func VreinterpretqU8P128(v0 Poly128) Uint8X16 {
	return Uint8X16(C.vreinterpretq_u8_p128(C.poly128_t(v0)))
}

func VreinterpretqU8P64(v0 Poly64X2) Uint8X16 {
	return Uint8X16(C.vreinterpretq_u8_p64(C.poly64x2_t(v0)))
}

func VreinterpretqU8P16(v0 Poly16X8) Uint8X16 {
	return Uint8X16(C.vreinterpretq_u8_p16(C.poly16x8_t(v0)))
}

func VreinterpretqU8U32(v0 Uint32X4) Uint8X16 {
	return Uint8X16(C.vreinterpretq_u8_u32(C.uint32x4_t(v0)))
}

func VreinterpretqU8U64(v0 Uint64X2) Uint8X16 {
	return Uint8X16(C.vreinterpretq_u8_u64(C.uint64x2_t(v0)))
}

func VreinterpretqU8U16(v0 Uint16X8) Uint8X16 {
	return Uint8X16(C.vreinterpretq_u8_u16(C.uint16x8_t(v0)))
}

func VreinterpretqU8S8(v0 Int8X16) Uint8X16 {
	return Uint8X16(C.vreinterpretq_u8_s8(C.int8x16_t(v0)))
}

func VreinterpretqU8F64(v0 Float64X2) Uint8X16 {
	return Uint8X16(C.vreinterpretq_u8_f64(C.float64x2_t(v0)))
}

func VreinterpretqU8F32(v0 Float32X4) Uint8X16 {
	return Uint8X16(C.vreinterpretq_u8_f32(C.float32x4_t(v0)))
}

func VreinterpretqU8S32(v0 Int32X4) Uint8X16 {
	return Uint8X16(C.vreinterpretq_u8_s32(C.int32x4_t(v0)))
}

func VreinterpretqU8S64(v0 Int64X2) Uint8X16 {
	return Uint8X16(C.vreinterpretq_u8_s64(C.int64x2_t(v0)))
}

func VreinterpretqU8S16(v0 Int16X8) Uint8X16 {
	return Uint8X16(C.vreinterpretq_u8_s16(C.int16x8_t(v0)))
}

func VreinterpretqU32P8(v0 Poly8X16) Uint32X4 {
	return Uint32X4(C.vreinterpretq_u32_p8(C.poly8x16_t(v0)))
}

func VreinterpretqU32P128(v0 Poly128) Uint32X4 {
	return Uint32X4(C.vreinterpretq_u32_p128(C.poly128_t(v0)))
}

func VreinterpretqU32P64(v0 Poly64X2) Uint32X4 {
	return Uint32X4(C.vreinterpretq_u32_p64(C.poly64x2_t(v0)))
}

func VreinterpretqU32P16(v0 Poly16X8) Uint32X4 {
	return Uint32X4(C.vreinterpretq_u32_p16(C.poly16x8_t(v0)))
}

func VreinterpretqU32U8(v0 Uint8X16) Uint32X4 {
	return Uint32X4(C.vreinterpretq_u32_u8(C.uint8x16_t(v0)))
}

func VreinterpretqU32U64(v0 Uint64X2) Uint32X4 {
	return Uint32X4(C.vreinterpretq_u32_u64(C.uint64x2_t(v0)))
}

func VreinterpretqU32U16(v0 Uint16X8) Uint32X4 {
	return Uint32X4(C.vreinterpretq_u32_u16(C.uint16x8_t(v0)))
}

func VreinterpretqU32S8(v0 Int8X16) Uint32X4 {
	return Uint32X4(C.vreinterpretq_u32_s8(C.int8x16_t(v0)))
}

func VreinterpretqU32F64(v0 Float64X2) Uint32X4 {
	return Uint32X4(C.vreinterpretq_u32_f64(C.float64x2_t(v0)))
}

func VreinterpretqU32F32(v0 Float32X4) Uint32X4 {
	return Uint32X4(C.vreinterpretq_u32_f32(C.float32x4_t(v0)))
}

func VreinterpretqU32S32(v0 Int32X4) Uint32X4 {
	return Uint32X4(C.vreinterpretq_u32_s32(C.int32x4_t(v0)))
}

func VreinterpretqU32S64(v0 Int64X2) Uint32X4 {
	return Uint32X4(C.vreinterpretq_u32_s64(C.int64x2_t(v0)))
}

func VreinterpretqU32S16(v0 Int16X8) Uint32X4 {
	return Uint32X4(C.vreinterpretq_u32_s16(C.int16x8_t(v0)))
}

func VreinterpretqU64P8(v0 Poly8X16) Uint64X2 {
	return Uint64X2(C.vreinterpretq_u64_p8(C.poly8x16_t(v0)))
}

func VreinterpretqU64P128(v0 Poly128) Uint64X2 {
	return Uint64X2(C.vreinterpretq_u64_p128(C.poly128_t(v0)))
}

func VreinterpretqU64P64(v0 Poly64X2) Uint64X2 {
	return Uint64X2(C.vreinterpretq_u64_p64(C.poly64x2_t(v0)))
}

func VreinterpretqU64P16(v0 Poly16X8) Uint64X2 {
	return Uint64X2(C.vreinterpretq_u64_p16(C.poly16x8_t(v0)))
}

func VreinterpretqU64U8(v0 Uint8X16) Uint64X2 {
	return Uint64X2(C.vreinterpretq_u64_u8(C.uint8x16_t(v0)))
}

func VreinterpretqU64U32(v0 Uint32X4) Uint64X2 {
	return Uint64X2(C.vreinterpretq_u64_u32(C.uint32x4_t(v0)))
}

func VreinterpretqU64U16(v0 Uint16X8) Uint64X2 {
	return Uint64X2(C.vreinterpretq_u64_u16(C.uint16x8_t(v0)))
}

func VreinterpretqU64S8(v0 Int8X16) Uint64X2 {
	return Uint64X2(C.vreinterpretq_u64_s8(C.int8x16_t(v0)))
}

func VreinterpretqU64F64(v0 Float64X2) Uint64X2 {
	return Uint64X2(C.vreinterpretq_u64_f64(C.float64x2_t(v0)))
}

func VreinterpretqU64F32(v0 Float32X4) Uint64X2 {
	return Uint64X2(C.vreinterpretq_u64_f32(C.float32x4_t(v0)))
}

func VreinterpretqU64S32(v0 Int32X4) Uint64X2 {
	return Uint64X2(C.vreinterpretq_u64_s32(C.int32x4_t(v0)))
}

func VreinterpretqU64S64(v0 Int64X2) Uint64X2 {
	return Uint64X2(C.vreinterpretq_u64_s64(C.int64x2_t(v0)))
}

func VreinterpretqU64S16(v0 Int16X8) Uint64X2 {
	return Uint64X2(C.vreinterpretq_u64_s16(C.int16x8_t(v0)))
}

func VreinterpretqU16P8(v0 Poly8X16) Uint16X8 {
	return Uint16X8(C.vreinterpretq_u16_p8(C.poly8x16_t(v0)))
}

func VreinterpretqU16P128(v0 Poly128) Uint16X8 {
	return Uint16X8(C.vreinterpretq_u16_p128(C.poly128_t(v0)))
}

func VreinterpretqU16P64(v0 Poly64X2) Uint16X8 {
	return Uint16X8(C.vreinterpretq_u16_p64(C.poly64x2_t(v0)))
}

func VreinterpretqU16P16(v0 Poly16X8) Uint16X8 {
	return Uint16X8(C.vreinterpretq_u16_p16(C.poly16x8_t(v0)))
}

func VreinterpretqU16U8(v0 Uint8X16) Uint16X8 {
	return Uint16X8(C.vreinterpretq_u16_u8(C.uint8x16_t(v0)))
}

func VreinterpretqU16U32(v0 Uint32X4) Uint16X8 {
	return Uint16X8(C.vreinterpretq_u16_u32(C.uint32x4_t(v0)))
}

func VreinterpretqU16U64(v0 Uint64X2) Uint16X8 {
	return Uint16X8(C.vreinterpretq_u16_u64(C.uint64x2_t(v0)))
}

func VreinterpretqU16S8(v0 Int8X16) Uint16X8 {
	return Uint16X8(C.vreinterpretq_u16_s8(C.int8x16_t(v0)))
}

func VreinterpretqU16F64(v0 Float64X2) Uint16X8 {
	return Uint16X8(C.vreinterpretq_u16_f64(C.float64x2_t(v0)))
}

func VreinterpretqU16F32(v0 Float32X4) Uint16X8 {
	return Uint16X8(C.vreinterpretq_u16_f32(C.float32x4_t(v0)))
}

func VreinterpretqU16S32(v0 Int32X4) Uint16X8 {
	return Uint16X8(C.vreinterpretq_u16_s32(C.int32x4_t(v0)))
}

func VreinterpretqU16S64(v0 Int64X2) Uint16X8 {
	return Uint16X8(C.vreinterpretq_u16_s64(C.int64x2_t(v0)))
}

func VreinterpretqU16S16(v0 Int16X8) Uint16X8 {
	return Uint16X8(C.vreinterpretq_u16_s16(C.int16x8_t(v0)))
}

func VreinterpretqS8P8(v0 Poly8X16) Int8X16 {
	return Int8X16(C.vreinterpretq_s8_p8(C.poly8x16_t(v0)))
}

func VreinterpretqS8P128(v0 Poly128) Int8X16 {
	return Int8X16(C.vreinterpretq_s8_p128(C.poly128_t(v0)))
}

func VreinterpretqS8P64(v0 Poly64X2) Int8X16 {
	return Int8X16(C.vreinterpretq_s8_p64(C.poly64x2_t(v0)))
}

func VreinterpretqS8P16(v0 Poly16X8) Int8X16 {
	return Int8X16(C.vreinterpretq_s8_p16(C.poly16x8_t(v0)))
}

func VreinterpretqS8U8(v0 Uint8X16) Int8X16 {
	return Int8X16(C.vreinterpretq_s8_u8(C.uint8x16_t(v0)))
}

func VreinterpretqS8U32(v0 Uint32X4) Int8X16 {
	return Int8X16(C.vreinterpretq_s8_u32(C.uint32x4_t(v0)))
}

func VreinterpretqS8U64(v0 Uint64X2) Int8X16 {
	return Int8X16(C.vreinterpretq_s8_u64(C.uint64x2_t(v0)))
}

func VreinterpretqS8U16(v0 Uint16X8) Int8X16 {
	return Int8X16(C.vreinterpretq_s8_u16(C.uint16x8_t(v0)))
}

func VreinterpretqS8F64(v0 Float64X2) Int8X16 {
	return Int8X16(C.vreinterpretq_s8_f64(C.float64x2_t(v0)))
}

func VreinterpretqS8F32(v0 Float32X4) Int8X16 {
	return Int8X16(C.vreinterpretq_s8_f32(C.float32x4_t(v0)))
}

func VreinterpretqS8S32(v0 Int32X4) Int8X16 {
	return Int8X16(C.vreinterpretq_s8_s32(C.int32x4_t(v0)))
}

func VreinterpretqS8S64(v0 Int64X2) Int8X16 {
	return Int8X16(C.vreinterpretq_s8_s64(C.int64x2_t(v0)))
}

func VreinterpretqS8S16(v0 Int16X8) Int8X16 {
	return Int8X16(C.vreinterpretq_s8_s16(C.int16x8_t(v0)))
}

func VreinterpretqF64P8(v0 Poly8X16) Float64X2 {
	return Float64X2(C.vreinterpretq_f64_p8(C.poly8x16_t(v0)))
}

func VreinterpretqF64P128(v0 Poly128) Float64X2 {
	return Float64X2(C.vreinterpretq_f64_p128(C.poly128_t(v0)))
}

func VreinterpretqF64P64(v0 Poly64X2) Float64X2 {
	return Float64X2(C.vreinterpretq_f64_p64(C.poly64x2_t(v0)))
}

func VreinterpretqF64P16(v0 Poly16X8) Float64X2 {
	return Float64X2(C.vreinterpretq_f64_p16(C.poly16x8_t(v0)))
}

func VreinterpretqF64U8(v0 Uint8X16) Float64X2 {
	return Float64X2(C.vreinterpretq_f64_u8(C.uint8x16_t(v0)))
}

func VreinterpretqF64U32(v0 Uint32X4) Float64X2 {
	return Float64X2(C.vreinterpretq_f64_u32(C.uint32x4_t(v0)))
}

func VreinterpretqF64U64(v0 Uint64X2) Float64X2 {
	return Float64X2(C.vreinterpretq_f64_u64(C.uint64x2_t(v0)))
}

func VreinterpretqF64U16(v0 Uint16X8) Float64X2 {
	return Float64X2(C.vreinterpretq_f64_u16(C.uint16x8_t(v0)))
}

func VreinterpretqF64S8(v0 Int8X16) Float64X2 {
	return Float64X2(C.vreinterpretq_f64_s8(C.int8x16_t(v0)))
}

func VreinterpretqF64F32(v0 Float32X4) Float64X2 {
	return Float64X2(C.vreinterpretq_f64_f32(C.float32x4_t(v0)))
}

func VreinterpretqF64S32(v0 Int32X4) Float64X2 {
	return Float64X2(C.vreinterpretq_f64_s32(C.int32x4_t(v0)))
}

func VreinterpretqF64S64(v0 Int64X2) Float64X2 {
	return Float64X2(C.vreinterpretq_f64_s64(C.int64x2_t(v0)))
}

func VreinterpretqF64S16(v0 Int16X8) Float64X2 {
	return Float64X2(C.vreinterpretq_f64_s16(C.int16x8_t(v0)))
}

func VreinterpretqF32P8(v0 Poly8X16) Float32X4 {
	return Float32X4(C.vreinterpretq_f32_p8(C.poly8x16_t(v0)))
}

func VreinterpretqF32P128(v0 Poly128) Float32X4 {
	return Float32X4(C.vreinterpretq_f32_p128(C.poly128_t(v0)))
}

func VreinterpretqF32P64(v0 Poly64X2) Float32X4 {
	return Float32X4(C.vreinterpretq_f32_p64(C.poly64x2_t(v0)))
}

func VreinterpretqF32P16(v0 Poly16X8) Float32X4 {
	return Float32X4(C.vreinterpretq_f32_p16(C.poly16x8_t(v0)))
}

func VreinterpretqF32U8(v0 Uint8X16) Float32X4 {
	return Float32X4(C.vreinterpretq_f32_u8(C.uint8x16_t(v0)))
}

func VreinterpretqF32U32(v0 Uint32X4) Float32X4 {
	return Float32X4(C.vreinterpretq_f32_u32(C.uint32x4_t(v0)))
}

func VreinterpretqF32U64(v0 Uint64X2) Float32X4 {
	return Float32X4(C.vreinterpretq_f32_u64(C.uint64x2_t(v0)))
}

func VreinterpretqF32U16(v0 Uint16X8) Float32X4 {
	return Float32X4(C.vreinterpretq_f32_u16(C.uint16x8_t(v0)))
}

func VreinterpretqF32S8(v0 Int8X16) Float32X4 {
	return Float32X4(C.vreinterpretq_f32_s8(C.int8x16_t(v0)))
}

func VreinterpretqF32F64(v0 Float64X2) Float32X4 {
	return Float32X4(C.vreinterpretq_f32_f64(C.float64x2_t(v0)))
}

func VreinterpretqF32S32(v0 Int32X4) Float32X4 {
	return Float32X4(C.vreinterpretq_f32_s32(C.int32x4_t(v0)))
}

func VreinterpretqF32S64(v0 Int64X2) Float32X4 {
	return Float32X4(C.vreinterpretq_f32_s64(C.int64x2_t(v0)))
}

func VreinterpretqF32S16(v0 Int16X8) Float32X4 {
	return Float32X4(C.vreinterpretq_f32_s16(C.int16x8_t(v0)))
}

func VreinterpretqS32P8(v0 Poly8X16) Int32X4 {
	return Int32X4(C.vreinterpretq_s32_p8(C.poly8x16_t(v0)))
}

func VreinterpretqS32P128(v0 Poly128) Int32X4 {
	return Int32X4(C.vreinterpretq_s32_p128(C.poly128_t(v0)))
}

func VreinterpretqS32P64(v0 Poly64X2) Int32X4 {
	return Int32X4(C.vreinterpretq_s32_p64(C.poly64x2_t(v0)))
}

func VreinterpretqS32P16(v0 Poly16X8) Int32X4 {
	return Int32X4(C.vreinterpretq_s32_p16(C.poly16x8_t(v0)))
}

func VreinterpretqS32U8(v0 Uint8X16) Int32X4 {
	return Int32X4(C.vreinterpretq_s32_u8(C.uint8x16_t(v0)))
}

func VreinterpretqS32U32(v0 Uint32X4) Int32X4 {
	return Int32X4(C.vreinterpretq_s32_u32(C.uint32x4_t(v0)))
}

func VreinterpretqS32U64(v0 Uint64X2) Int32X4 {
	return Int32X4(C.vreinterpretq_s32_u64(C.uint64x2_t(v0)))
}

func VreinterpretqS32U16(v0 Uint16X8) Int32X4 {
	return Int32X4(C.vreinterpretq_s32_u16(C.uint16x8_t(v0)))
}

func VreinterpretqS32S8(v0 Int8X16) Int32X4 {
	return Int32X4(C.vreinterpretq_s32_s8(C.int8x16_t(v0)))
}

func VreinterpretqS32F64(v0 Float64X2) Int32X4 {
	return Int32X4(C.vreinterpretq_s32_f64(C.float64x2_t(v0)))
}

func VreinterpretqS32F32(v0 Float32X4) Int32X4 {
	return Int32X4(C.vreinterpretq_s32_f32(C.float32x4_t(v0)))
}

func VreinterpretqS32S64(v0 Int64X2) Int32X4 {
	return Int32X4(C.vreinterpretq_s32_s64(C.int64x2_t(v0)))
}

func VreinterpretqS32S16(v0 Int16X8) Int32X4 {
	return Int32X4(C.vreinterpretq_s32_s16(C.int16x8_t(v0)))
}

func VreinterpretqS64P8(v0 Poly8X16) Int64X2 {
	return Int64X2(C.vreinterpretq_s64_p8(C.poly8x16_t(v0)))
}

func VreinterpretqS64P128(v0 Poly128) Int64X2 {
	return Int64X2(C.vreinterpretq_s64_p128(C.poly128_t(v0)))
}

func VreinterpretqS64P64(v0 Poly64X2) Int64X2 {
	return Int64X2(C.vreinterpretq_s64_p64(C.poly64x2_t(v0)))
}

func VreinterpretqS64P16(v0 Poly16X8) Int64X2 {
	return Int64X2(C.vreinterpretq_s64_p16(C.poly16x8_t(v0)))
}

func VreinterpretqS64U8(v0 Uint8X16) Int64X2 {
	return Int64X2(C.vreinterpretq_s64_u8(C.uint8x16_t(v0)))
}

func VreinterpretqS64U32(v0 Uint32X4) Int64X2 {
	return Int64X2(C.vreinterpretq_s64_u32(C.uint32x4_t(v0)))
}

func VreinterpretqS64U64(v0 Uint64X2) Int64X2 {
	return Int64X2(C.vreinterpretq_s64_u64(C.uint64x2_t(v0)))
}

func VreinterpretqS64U16(v0 Uint16X8) Int64X2 {
	return Int64X2(C.vreinterpretq_s64_u16(C.uint16x8_t(v0)))
}

func VreinterpretqS64S8(v0 Int8X16) Int64X2 {
	return Int64X2(C.vreinterpretq_s64_s8(C.int8x16_t(v0)))
}

func VreinterpretqS64F64(v0 Float64X2) Int64X2 {
	return Int64X2(C.vreinterpretq_s64_f64(C.float64x2_t(v0)))
}

func VreinterpretqS64F32(v0 Float32X4) Int64X2 {
	return Int64X2(C.vreinterpretq_s64_f32(C.float32x4_t(v0)))
}

func VreinterpretqS64S32(v0 Int32X4) Int64X2 {
	return Int64X2(C.vreinterpretq_s64_s32(C.int32x4_t(v0)))
}

func VreinterpretqS64S16(v0 Int16X8) Int64X2 {
	return Int64X2(C.vreinterpretq_s64_s16(C.int16x8_t(v0)))
}

func VreinterpretqS16P8(v0 Poly8X16) Int16X8 {
	return Int16X8(C.vreinterpretq_s16_p8(C.poly8x16_t(v0)))
}

func VreinterpretqS16P128(v0 Poly128) Int16X8 {
	return Int16X8(C.vreinterpretq_s16_p128(C.poly128_t(v0)))
}

func VreinterpretqS16P64(v0 Poly64X2) Int16X8 {
	return Int16X8(C.vreinterpretq_s16_p64(C.poly64x2_t(v0)))
}

func VreinterpretqS16P16(v0 Poly16X8) Int16X8 {
	return Int16X8(C.vreinterpretq_s16_p16(C.poly16x8_t(v0)))
}

func VreinterpretqS16U8(v0 Uint8X16) Int16X8 {
	return Int16X8(C.vreinterpretq_s16_u8(C.uint8x16_t(v0)))
}

func VreinterpretqS16U32(v0 Uint32X4) Int16X8 {
	return Int16X8(C.vreinterpretq_s16_u32(C.uint32x4_t(v0)))
}

func VreinterpretqS16U64(v0 Uint64X2) Int16X8 {
	return Int16X8(C.vreinterpretq_s16_u64(C.uint64x2_t(v0)))
}

func VreinterpretqS16U16(v0 Uint16X8) Int16X8 {
	return Int16X8(C.vreinterpretq_s16_u16(C.uint16x8_t(v0)))
}

func VreinterpretqS16S8(v0 Int8X16) Int16X8 {
	return Int16X8(C.vreinterpretq_s16_s8(C.int8x16_t(v0)))
}

func VreinterpretqS16F64(v0 Float64X2) Int16X8 {
	return Int16X8(C.vreinterpretq_s16_f64(C.float64x2_t(v0)))
}

func VreinterpretqS16F32(v0 Float32X4) Int16X8 {
	return Int16X8(C.vreinterpretq_s16_f32(C.float32x4_t(v0)))
}

func VreinterpretqS16S32(v0 Int32X4) Int16X8 {
	return Int16X8(C.vreinterpretq_s16_s32(C.int32x4_t(v0)))
}

func VreinterpretqS16S64(v0 Int64X2) Int16X8 {
	return Int16X8(C.vreinterpretq_s16_s64(C.int64x2_t(v0)))
}

func VreinterpretU8P8(v0 Poly8X8) Uint8X8 {
	return Uint8X8(C.vreinterpret_u8_p8(C.poly8x8_t(v0)))
}

func VreinterpretU8P64(v0 Poly64X1) Uint8X8 {
	return Uint8X8(C.vreinterpret_u8_p64(C.poly64x1_t(v0)))
}

func VreinterpretU8P16(v0 Poly16X4) Uint8X8 {
	return Uint8X8(C.vreinterpret_u8_p16(C.poly16x4_t(v0)))
}

func VreinterpretU8U32(v0 Uint32X2) Uint8X8 {
	return Uint8X8(C.vreinterpret_u8_u32(C.uint32x2_t(v0)))
}

func VreinterpretU8U64(v0 Uint64X1) Uint8X8 {
	return Uint8X8(C.vreinterpret_u8_u64(C.uint64x1_t(v0)))
}

func VreinterpretU8U16(v0 Uint16X4) Uint8X8 {
	return Uint8X8(C.vreinterpret_u8_u16(C.uint16x4_t(v0)))
}

func VreinterpretU8S8(v0 Int8X8) Uint8X8 {
	return Uint8X8(C.vreinterpret_u8_s8(C.int8x8_t(v0)))
}

func VreinterpretU8F64(v0 Float64X1) Uint8X8 {
	return Uint8X8(C.vreinterpret_u8_f64(C.float64x1_t(v0)))
}

func VreinterpretU8F32(v0 Float32X2) Uint8X8 {
	return Uint8X8(C.vreinterpret_u8_f32(C.float32x2_t(v0)))
}

func VreinterpretU8S32(v0 Int32X2) Uint8X8 {
	return Uint8X8(C.vreinterpret_u8_s32(C.int32x2_t(v0)))
}

func VreinterpretU8S64(v0 Int64X1) Uint8X8 {
	return Uint8X8(C.vreinterpret_u8_s64(C.int64x1_t(v0)))
}

func VreinterpretU8S16(v0 Int16X4) Uint8X8 {
	return Uint8X8(C.vreinterpret_u8_s16(C.int16x4_t(v0)))
}

func VreinterpretU32P8(v0 Poly8X8) Uint32X2 {
	return Uint32X2(C.vreinterpret_u32_p8(C.poly8x8_t(v0)))
}

func VreinterpretU32P64(v0 Poly64X1) Uint32X2 {
	return Uint32X2(C.vreinterpret_u32_p64(C.poly64x1_t(v0)))
}

func VreinterpretU32P16(v0 Poly16X4) Uint32X2 {
	return Uint32X2(C.vreinterpret_u32_p16(C.poly16x4_t(v0)))
}

func VreinterpretU32U8(v0 Uint8X8) Uint32X2 {
	return Uint32X2(C.vreinterpret_u32_u8(C.uint8x8_t(v0)))
}

func VreinterpretU32U64(v0 Uint64X1) Uint32X2 {
	return Uint32X2(C.vreinterpret_u32_u64(C.uint64x1_t(v0)))
}

func VreinterpretU32U16(v0 Uint16X4) Uint32X2 {
	return Uint32X2(C.vreinterpret_u32_u16(C.uint16x4_t(v0)))
}

func VreinterpretU32S8(v0 Int8X8) Uint32X2 {
	return Uint32X2(C.vreinterpret_u32_s8(C.int8x8_t(v0)))
}

func VreinterpretU32F64(v0 Float64X1) Uint32X2 {
	return Uint32X2(C.vreinterpret_u32_f64(C.float64x1_t(v0)))
}

func VreinterpretU32F32(v0 Float32X2) Uint32X2 {
	return Uint32X2(C.vreinterpret_u32_f32(C.float32x2_t(v0)))
}

func VreinterpretU32S32(v0 Int32X2) Uint32X2 {
	return Uint32X2(C.vreinterpret_u32_s32(C.int32x2_t(v0)))
}

func VreinterpretU32S64(v0 Int64X1) Uint32X2 {
	return Uint32X2(C.vreinterpret_u32_s64(C.int64x1_t(v0)))
}

func VreinterpretU32S16(v0 Int16X4) Uint32X2 {
	return Uint32X2(C.vreinterpret_u32_s16(C.int16x4_t(v0)))
}

func VreinterpretU64P8(v0 Poly8X8) Uint64X1 {
	return Uint64X1(C.vreinterpret_u64_p8(C.poly8x8_t(v0)))
}

func VreinterpretU64P64(v0 Poly64X1) Uint64X1 {
	return Uint64X1(C.vreinterpret_u64_p64(C.poly64x1_t(v0)))
}

func VreinterpretU64P16(v0 Poly16X4) Uint64X1 {
	return Uint64X1(C.vreinterpret_u64_p16(C.poly16x4_t(v0)))
}

func VreinterpretU64U8(v0 Uint8X8) Uint64X1 {
	return Uint64X1(C.vreinterpret_u64_u8(C.uint8x8_t(v0)))
}

func VreinterpretU64U32(v0 Uint32X2) Uint64X1 {
	return Uint64X1(C.vreinterpret_u64_u32(C.uint32x2_t(v0)))
}

func VreinterpretU64U16(v0 Uint16X4) Uint64X1 {
	return Uint64X1(C.vreinterpret_u64_u16(C.uint16x4_t(v0)))
}

func VreinterpretU64S8(v0 Int8X8) Uint64X1 {
	return Uint64X1(C.vreinterpret_u64_s8(C.int8x8_t(v0)))
}

func VreinterpretU64F64(v0 Float64X1) Uint64X1 {
	return Uint64X1(C.vreinterpret_u64_f64(C.float64x1_t(v0)))
}

func VreinterpretU64F32(v0 Float32X2) Uint64X1 {
	return Uint64X1(C.vreinterpret_u64_f32(C.float32x2_t(v0)))
}

func VreinterpretU64S32(v0 Int32X2) Uint64X1 {
	return Uint64X1(C.vreinterpret_u64_s32(C.int32x2_t(v0)))
}

func VreinterpretU64S64(v0 Int64X1) Uint64X1 {
	return Uint64X1(C.vreinterpret_u64_s64(C.int64x1_t(v0)))
}

func VreinterpretU64S16(v0 Int16X4) Uint64X1 {
	return Uint64X1(C.vreinterpret_u64_s16(C.int16x4_t(v0)))
}

func VreinterpretU16P8(v0 Poly8X8) Uint16X4 {
	return Uint16X4(C.vreinterpret_u16_p8(C.poly8x8_t(v0)))
}

func VreinterpretU16P64(v0 Poly64X1) Uint16X4 {
	return Uint16X4(C.vreinterpret_u16_p64(C.poly64x1_t(v0)))
}

func VreinterpretU16P16(v0 Poly16X4) Uint16X4 {
	return Uint16X4(C.vreinterpret_u16_p16(C.poly16x4_t(v0)))
}

func VreinterpretU16U8(v0 Uint8X8) Uint16X4 {
	return Uint16X4(C.vreinterpret_u16_u8(C.uint8x8_t(v0)))
}

func VreinterpretU16U32(v0 Uint32X2) Uint16X4 {
	return Uint16X4(C.vreinterpret_u16_u32(C.uint32x2_t(v0)))
}

func VreinterpretU16U64(v0 Uint64X1) Uint16X4 {
	return Uint16X4(C.vreinterpret_u16_u64(C.uint64x1_t(v0)))
}

func VreinterpretU16S8(v0 Int8X8) Uint16X4 {
	return Uint16X4(C.vreinterpret_u16_s8(C.int8x8_t(v0)))
}

func VreinterpretU16F64(v0 Float64X1) Uint16X4 {
	return Uint16X4(C.vreinterpret_u16_f64(C.float64x1_t(v0)))
}

func VreinterpretU16F32(v0 Float32X2) Uint16X4 {
	return Uint16X4(C.vreinterpret_u16_f32(C.float32x2_t(v0)))
}

func VreinterpretU16S32(v0 Int32X2) Uint16X4 {
	return Uint16X4(C.vreinterpret_u16_s32(C.int32x2_t(v0)))
}

func VreinterpretU16S64(v0 Int64X1) Uint16X4 {
	return Uint16X4(C.vreinterpret_u16_s64(C.int64x1_t(v0)))
}

func VreinterpretU16S16(v0 Int16X4) Uint16X4 {
	return Uint16X4(C.vreinterpret_u16_s16(C.int16x4_t(v0)))
}

func VreinterpretS8P8(v0 Poly8X8) Int8X8 {
	return Int8X8(C.vreinterpret_s8_p8(C.poly8x8_t(v0)))
}

func VreinterpretS8P64(v0 Poly64X1) Int8X8 {
	return Int8X8(C.vreinterpret_s8_p64(C.poly64x1_t(v0)))
}

func VreinterpretS8P16(v0 Poly16X4) Int8X8 {
	return Int8X8(C.vreinterpret_s8_p16(C.poly16x4_t(v0)))
}

func VreinterpretS8U8(v0 Uint8X8) Int8X8 {
	return Int8X8(C.vreinterpret_s8_u8(C.uint8x8_t(v0)))
}

func VreinterpretS8U32(v0 Uint32X2) Int8X8 {
	return Int8X8(C.vreinterpret_s8_u32(C.uint32x2_t(v0)))
}

func VreinterpretS8U64(v0 Uint64X1) Int8X8 {
	return Int8X8(C.vreinterpret_s8_u64(C.uint64x1_t(v0)))
}

func VreinterpretS8U16(v0 Uint16X4) Int8X8 {
	return Int8X8(C.vreinterpret_s8_u16(C.uint16x4_t(v0)))
}

func VreinterpretS8F64(v0 Float64X1) Int8X8 {
	return Int8X8(C.vreinterpret_s8_f64(C.float64x1_t(v0)))
}

func VreinterpretS8F32(v0 Float32X2) Int8X8 {
	return Int8X8(C.vreinterpret_s8_f32(C.float32x2_t(v0)))
}

func VreinterpretS8S32(v0 Int32X2) Int8X8 {
	return Int8X8(C.vreinterpret_s8_s32(C.int32x2_t(v0)))
}

func VreinterpretS8S64(v0 Int64X1) Int8X8 {
	return Int8X8(C.vreinterpret_s8_s64(C.int64x1_t(v0)))
}

func VreinterpretS8S16(v0 Int16X4) Int8X8 {
	return Int8X8(C.vreinterpret_s8_s16(C.int16x4_t(v0)))
}

func VreinterpretF64P8(v0 Poly8X8) Float64X1 {
	return Float64X1(C.vreinterpret_f64_p8(C.poly8x8_t(v0)))
}

func VreinterpretF64P64(v0 Poly64X1) Float64X1 {
	return Float64X1(C.vreinterpret_f64_p64(C.poly64x1_t(v0)))
}

func VreinterpretF64P16(v0 Poly16X4) Float64X1 {
	return Float64X1(C.vreinterpret_f64_p16(C.poly16x4_t(v0)))
}

func VreinterpretF64U8(v0 Uint8X8) Float64X1 {
	return Float64X1(C.vreinterpret_f64_u8(C.uint8x8_t(v0)))
}

func VreinterpretF64U32(v0 Uint32X2) Float64X1 {
	return Float64X1(C.vreinterpret_f64_u32(C.uint32x2_t(v0)))
}

func VreinterpretF64U64(v0 Uint64X1) Float64X1 {
	return Float64X1(C.vreinterpret_f64_u64(C.uint64x1_t(v0)))
}

func VreinterpretF64U16(v0 Uint16X4) Float64X1 {
	return Float64X1(C.vreinterpret_f64_u16(C.uint16x4_t(v0)))
}

func VreinterpretF64S8(v0 Int8X8) Float64X1 {
	return Float64X1(C.vreinterpret_f64_s8(C.int8x8_t(v0)))
}

func VreinterpretF64F32(v0 Float32X2) Float64X1 {
	return Float64X1(C.vreinterpret_f64_f32(C.float32x2_t(v0)))
}

func VreinterpretF64S32(v0 Int32X2) Float64X1 {
	return Float64X1(C.vreinterpret_f64_s32(C.int32x2_t(v0)))
}

func VreinterpretF64S64(v0 Int64X1) Float64X1 {
	return Float64X1(C.vreinterpret_f64_s64(C.int64x1_t(v0)))
}

func VreinterpretF64S16(v0 Int16X4) Float64X1 {
	return Float64X1(C.vreinterpret_f64_s16(C.int16x4_t(v0)))
}

func VreinterpretF32P8(v0 Poly8X8) Float32X2 {
	return Float32X2(C.vreinterpret_f32_p8(C.poly8x8_t(v0)))
}

func VreinterpretF32P64(v0 Poly64X1) Float32X2 {
	return Float32X2(C.vreinterpret_f32_p64(C.poly64x1_t(v0)))
}

func VreinterpretF32P16(v0 Poly16X4) Float32X2 {
	return Float32X2(C.vreinterpret_f32_p16(C.poly16x4_t(v0)))
}

func VreinterpretF32U8(v0 Uint8X8) Float32X2 {
	return Float32X2(C.vreinterpret_f32_u8(C.uint8x8_t(v0)))
}

func VreinterpretF32U32(v0 Uint32X2) Float32X2 {
	return Float32X2(C.vreinterpret_f32_u32(C.uint32x2_t(v0)))
}

func VreinterpretF32U64(v0 Uint64X1) Float32X2 {
	return Float32X2(C.vreinterpret_f32_u64(C.uint64x1_t(v0)))
}

func VreinterpretF32U16(v0 Uint16X4) Float32X2 {
	return Float32X2(C.vreinterpret_f32_u16(C.uint16x4_t(v0)))
}

func VreinterpretF32S8(v0 Int8X8) Float32X2 {
	return Float32X2(C.vreinterpret_f32_s8(C.int8x8_t(v0)))
}

func VreinterpretF32F64(v0 Float64X1) Float32X2 {
	return Float32X2(C.vreinterpret_f32_f64(C.float64x1_t(v0)))
}

func VreinterpretF32S32(v0 Int32X2) Float32X2 {
	return Float32X2(C.vreinterpret_f32_s32(C.int32x2_t(v0)))
}

func VreinterpretF32S64(v0 Int64X1) Float32X2 {
	return Float32X2(C.vreinterpret_f32_s64(C.int64x1_t(v0)))
}

func VreinterpretF32S16(v0 Int16X4) Float32X2 {
	return Float32X2(C.vreinterpret_f32_s16(C.int16x4_t(v0)))
}

func VreinterpretS32P8(v0 Poly8X8) Int32X2 {
	return Int32X2(C.vreinterpret_s32_p8(C.poly8x8_t(v0)))
}

func VreinterpretS32P64(v0 Poly64X1) Int32X2 {
	return Int32X2(C.vreinterpret_s32_p64(C.poly64x1_t(v0)))
}

func VreinterpretS32P16(v0 Poly16X4) Int32X2 {
	return Int32X2(C.vreinterpret_s32_p16(C.poly16x4_t(v0)))
}

func VreinterpretS32U8(v0 Uint8X8) Int32X2 {
	return Int32X2(C.vreinterpret_s32_u8(C.uint8x8_t(v0)))
}

func VreinterpretS32U32(v0 Uint32X2) Int32X2 {
	return Int32X2(C.vreinterpret_s32_u32(C.uint32x2_t(v0)))
}

func VreinterpretS32U64(v0 Uint64X1) Int32X2 {
	return Int32X2(C.vreinterpret_s32_u64(C.uint64x1_t(v0)))
}

func VreinterpretS32U16(v0 Uint16X4) Int32X2 {
	return Int32X2(C.vreinterpret_s32_u16(C.uint16x4_t(v0)))
}

func VreinterpretS32S8(v0 Int8X8) Int32X2 {
	return Int32X2(C.vreinterpret_s32_s8(C.int8x8_t(v0)))
}

func VreinterpretS32F64(v0 Float64X1) Int32X2 {
	return Int32X2(C.vreinterpret_s32_f64(C.float64x1_t(v0)))
}

func VreinterpretS32F32(v0 Float32X2) Int32X2 {
	return Int32X2(C.vreinterpret_s32_f32(C.float32x2_t(v0)))
}

func VreinterpretS32S64(v0 Int64X1) Int32X2 {
	return Int32X2(C.vreinterpret_s32_s64(C.int64x1_t(v0)))
}

func VreinterpretS32S16(v0 Int16X4) Int32X2 {
	return Int32X2(C.vreinterpret_s32_s16(C.int16x4_t(v0)))
}

func VreinterpretS64P8(v0 Poly8X8) Int64X1 {
	return Int64X1(C.vreinterpret_s64_p8(C.poly8x8_t(v0)))
}

func VreinterpretS64P64(v0 Poly64X1) Int64X1 {
	return Int64X1(C.vreinterpret_s64_p64(C.poly64x1_t(v0)))
}

func VreinterpretS64P16(v0 Poly16X4) Int64X1 {
	return Int64X1(C.vreinterpret_s64_p16(C.poly16x4_t(v0)))
}

func VreinterpretS64U8(v0 Uint8X8) Int64X1 {
	return Int64X1(C.vreinterpret_s64_u8(C.uint8x8_t(v0)))
}

func VreinterpretS64U32(v0 Uint32X2) Int64X1 {
	return Int64X1(C.vreinterpret_s64_u32(C.uint32x2_t(v0)))
}

func VreinterpretS64U64(v0 Uint64X1) Int64X1 {
	return Int64X1(C.vreinterpret_s64_u64(C.uint64x1_t(v0)))
}

func VreinterpretS64U16(v0 Uint16X4) Int64X1 {
	return Int64X1(C.vreinterpret_s64_u16(C.uint16x4_t(v0)))
}

func VreinterpretS64S8(v0 Int8X8) Int64X1 {
	return Int64X1(C.vreinterpret_s64_s8(C.int8x8_t(v0)))
}

func VreinterpretS64F64(v0 Float64X1) Int64X1 {
	return Int64X1(C.vreinterpret_s64_f64(C.float64x1_t(v0)))
}

func VreinterpretS64F32(v0 Float32X2) Int64X1 {
	return Int64X1(C.vreinterpret_s64_f32(C.float32x2_t(v0)))
}

func VreinterpretS64S32(v0 Int32X2) Int64X1 {
	return Int64X1(C.vreinterpret_s64_s32(C.int32x2_t(v0)))
}

func VreinterpretS64S16(v0 Int16X4) Int64X1 {
	return Int64X1(C.vreinterpret_s64_s16(C.int16x4_t(v0)))
}

func VreinterpretS16P8(v0 Poly8X8) Int16X4 {
	return Int16X4(C.vreinterpret_s16_p8(C.poly8x8_t(v0)))
}

func VreinterpretS16P64(v0 Poly64X1) Int16X4 {
	return Int16X4(C.vreinterpret_s16_p64(C.poly64x1_t(v0)))
}

func VreinterpretS16P16(v0 Poly16X4) Int16X4 {
	return Int16X4(C.vreinterpret_s16_p16(C.poly16x4_t(v0)))
}

func VreinterpretS16U8(v0 Uint8X8) Int16X4 {
	return Int16X4(C.vreinterpret_s16_u8(C.uint8x8_t(v0)))
}

func VreinterpretS16U32(v0 Uint32X2) Int16X4 {
	return Int16X4(C.vreinterpret_s16_u32(C.uint32x2_t(v0)))
}

func VreinterpretS16U64(v0 Uint64X1) Int16X4 {
	return Int16X4(C.vreinterpret_s16_u64(C.uint64x1_t(v0)))
}

func VreinterpretS16U16(v0 Uint16X4) Int16X4 {
	return Int16X4(C.vreinterpret_s16_u16(C.uint16x4_t(v0)))
}

func VreinterpretS16S8(v0 Int8X8) Int16X4 {
	return Int16X4(C.vreinterpret_s16_s8(C.int8x8_t(v0)))
}

func VreinterpretS16F64(v0 Float64X1) Int16X4 {
	return Int16X4(C.vreinterpret_s16_f64(C.float64x1_t(v0)))
}

func VreinterpretS16F32(v0 Float32X2) Int16X4 {
	return Int16X4(C.vreinterpret_s16_f32(C.float32x2_t(v0)))
}

func VreinterpretS16S32(v0 Int32X2) Int16X4 {
	return Int16X4(C.vreinterpret_s16_s32(C.int32x2_t(v0)))
}

func VreinterpretS16S64(v0 Int64X1) Int16X4 {
	return Int16X4(C.vreinterpret_s16_s64(C.int64x1_t(v0)))
}

func VrndqF64(v0 Float64X2) Float64X2 {
	return Float64X2(C.vrndq_f64(C.float64x2_t(v0)))
}

func VrndF64(v0 Float64X1) Float64X1 {
	return Float64X1(C.vrnd_f64(C.float64x1_t(v0)))
}

func VrndaqF64(v0 Float64X2) Float64X2 {
	return Float64X2(C.vrndaq_f64(C.float64x2_t(v0)))
}

func VrndaF64(v0 Float64X1) Float64X1 {
	return Float64X1(C.vrnda_f64(C.float64x1_t(v0)))
}

func VrndiqF64(v0 Float64X2) Float64X2 {
	return Float64X2(C.vrndiq_f64(C.float64x2_t(v0)))
}

func VrndiF64(v0 Float64X1) Float64X1 {
	return Float64X1(C.vrndi_f64(C.float64x1_t(v0)))
}

func VrndmqF64(v0 Float64X2) Float64X2 {
	return Float64X2(C.vrndmq_f64(C.float64x2_t(v0)))
}

func VrndmF64(v0 Float64X1) Float64X1 {
	return Float64X1(C.vrndm_f64(C.float64x1_t(v0)))
}

func VrndnqF64(v0 Float64X2) Float64X2 {
	return Float64X2(C.vrndnq_f64(C.float64x2_t(v0)))
}

func VrndnF64(v0 Float64X1) Float64X1 {
	return Float64X1(C.vrndn_f64(C.float64x1_t(v0)))
}

func VrndpqF64(v0 Float64X2) Float64X2 {
	return Float64X2(C.vrndpq_f64(C.float64x2_t(v0)))
}

func VrndpF64(v0 Float64X1) Float64X1 {
	return Float64X1(C.vrndp_f64(C.float64x1_t(v0)))
}

func VrndxqF64(v0 Float64X2) Float64X2 {
	return Float64X2(C.vrndxq_f64(C.float64x2_t(v0)))
}

func VrndxF64(v0 Float64X1) Float64X1 {
	return Float64X1(C.vrndx_f64(C.float64x1_t(v0)))
}

func Vrnd32XqF32(v0 Float32X4) Float32X4 {
	return Float32X4(C.vrnd32xq_f32(C.float32x4_t(v0)))
}

func Vrnd32XF32(v0 Float32X2) Float32X2 {
	return Float32X2(C.vrnd32x_f32(C.float32x2_t(v0)))
}

func Vrnd32ZqF32(v0 Float32X4) Float32X4 {
	return Float32X4(C.vrnd32zq_f32(C.float32x4_t(v0)))
}

func Vrnd32ZF32(v0 Float32X2) Float32X2 {
	return Float32X2(C.vrnd32z_f32(C.float32x2_t(v0)))
}

func Vrnd64XqF32(v0 Float32X4) Float32X4 {
	return Float32X4(C.vrnd64xq_f32(C.float32x4_t(v0)))
}

func Vrnd64XF32(v0 Float32X2) Float32X2 {
	return Float32X2(C.vrnd64x_f32(C.float32x2_t(v0)))
}

func Vrnd64ZqF32(v0 Float32X4) Float32X4 {
	return Float32X4(C.vrnd64zq_f32(C.float32x4_t(v0)))
}

func Vrnd64ZF32(v0 Float32X2) Float32X2 {
	return Float32X2(C.vrnd64z_f32(C.float32x2_t(v0)))
}

func VmaxnmqF64(v0 Float64X2, v1 Float64X2) Float64X2 {
	return Float64X2(C.vmaxnmq_f64(C.float64x2_t(v0), C.float64x2_t(v1)))
}

func VmaxnmF64(v0 Float64X1, v1 Float64X1) Float64X1 {
	return Float64X1(C.vmaxnm_f64(C.float64x1_t(v0), C.float64x1_t(v1)))
}

func VminnmqF64(v0 Float64X2, v1 Float64X2) Float64X2 {
	return Float64X2(C.vminnmq_f64(C.float64x2_t(v0), C.float64x2_t(v1)))
}

func VminnmF64(v0 Float64X1, v1 Float64X1) Float64X1 {
	return Float64X1(C.vminnm_f64(C.float64x1_t(v0), C.float64x1_t(v1)))
}

func VcaddRot270F32(v0 Float32X2, v1 Float32X2) Float32X2 {
	return Float32X2(C.vcadd_rot270_f32(C.float32x2_t(v0), C.float32x2_t(v1)))
}

func VcaddRot90F32(v0 Float32X2, v1 Float32X2) Float32X2 {
	return Float32X2(C.vcadd_rot90_f32(C.float32x2_t(v0), C.float32x2_t(v1)))
}

func VcaddqRot270F32(v0 Float32X4, v1 Float32X4) Float32X4 {
	return Float32X4(C.vcaddq_rot270_f32(C.float32x4_t(v0), C.float32x4_t(v1)))
}

func VcaddqRot90F32(v0 Float32X4, v1 Float32X4) Float32X4 {
	return Float32X4(C.vcaddq_rot90_f32(C.float32x4_t(v0), C.float32x4_t(v1)))
}

func VcaddqRot270F64(v0 Float64X2, v1 Float64X2) Float64X2 {
	return Float64X2(C.vcaddq_rot270_f64(C.float64x2_t(v0), C.float64x2_t(v1)))
}

func VcaddqRot90F64(v0 Float64X2, v1 Float64X2) Float64X2 {
	return Float64X2(C.vcaddq_rot90_f64(C.float64x2_t(v0), C.float64x2_t(v1)))
}

func VdotqU32(v0 Uint32X4, v1 Uint8X16, v2 Uint8X16) Uint32X4 {
	return Uint32X4(C.vdotq_u32(C.uint32x4_t(v0), C.uint8x16_t(v1), C.uint8x16_t(v2)))
}

func VdotqS32(v0 Int32X4, v1 Int8X16, v2 Int8X16) Int32X4 {
	return Int32X4(C.vdotq_s32(C.int32x4_t(v0), C.int8x16_t(v1), C.int8x16_t(v2)))
}

func VdotU32(v0 Uint32X2, v1 Uint8X8, v2 Uint8X8) Uint32X2 {
	return Uint32X2(C.vdot_u32(C.uint32x2_t(v0), C.uint8x8_t(v1), C.uint8x8_t(v2)))
}

func VdotS32(v0 Int32X2, v1 Int8X8, v2 Int8X8) Int32X2 {
	return Int32X2(C.vdot_s32(C.int32x2_t(v0), C.int8x8_t(v1), C.int8x8_t(v2)))
}

func VfmaqF32(v0 Float32X4, v1 Float32X4, v2 Float32X4) Float32X4 {
	return Float32X4(C.vfmaq_f32(C.float32x4_t(v0), C.float32x4_t(v1), C.float32x4_t(v2)))
}

func VfmaF32(v0 Float32X2, v1 Float32X2, v2 Float32X2) Float32X2 {
	return Float32X2(C.vfma_f32(C.float32x2_t(v0), C.float32x2_t(v1), C.float32x2_t(v2)))
}

func VfmaqNF32(v0 Float32X4, v1 Float32X4, v2 float32) Float32X4 {
	return Float32X4(C.vfmaq_n_f32(C.float32x4_t(v0), C.float32x4_t(v1), C.float32_t(v2)))
}

func VfmaNF32(v0 Float32X2, v1 Float32X2, v2 float32) Float32X2 {
	return Float32X2(C.vfma_n_f32(C.float32x2_t(v0), C.float32x2_t(v1), C.float32_t(v2)))
}

func VfmsqF32(v0 Float32X4, v1 Float32X4, v2 Float32X4) Float32X4 {
	return Float32X4(C.vfmsq_f32(C.float32x4_t(v0), C.float32x4_t(v1), C.float32x4_t(v2)))
}

func VfmsF32(v0 Float32X2, v1 Float32X2, v2 Float32X2) Float32X2 {
	return Float32X2(C.vfms_f32(C.float32x2_t(v0), C.float32x2_t(v1), C.float32x2_t(v2)))
}

func VqrdmlahqS32(v0 Int32X4, v1 Int32X4, v2 Int32X4) Int32X4 {
	return Int32X4(C.vqrdmlahq_s32(C.int32x4_t(v0), C.int32x4_t(v1), C.int32x4_t(v2)))
}

func VqrdmlahqS16(v0 Int16X8, v1 Int16X8, v2 Int16X8) Int16X8 {
	return Int16X8(C.vqrdmlahq_s16(C.int16x8_t(v0), C.int16x8_t(v1), C.int16x8_t(v2)))
}

func VqrdmlahS32(v0 Int32X2, v1 Int32X2, v2 Int32X2) Int32X2 {
	return Int32X2(C.vqrdmlah_s32(C.int32x2_t(v0), C.int32x2_t(v1), C.int32x2_t(v2)))
}

func VqrdmlahS16(v0 Int16X4, v1 Int16X4, v2 Int16X4) Int16X4 {
	return Int16X4(C.vqrdmlah_s16(C.int16x4_t(v0), C.int16x4_t(v1), C.int16x4_t(v2)))
}

func VqrdmlshqS32(v0 Int32X4, v1 Int32X4, v2 Int32X4) Int32X4 {
	return Int32X4(C.vqrdmlshq_s32(C.int32x4_t(v0), C.int32x4_t(v1), C.int32x4_t(v2)))
}

func VqrdmlshqS16(v0 Int16X8, v1 Int16X8, v2 Int16X8) Int16X8 {
	return Int16X8(C.vqrdmlshq_s16(C.int16x8_t(v0), C.int16x8_t(v1), C.int16x8_t(v2)))
}

func VqrdmlshS32(v0 Int32X2, v1 Int32X2, v2 Int32X2) Int32X2 {
	return Int32X2(C.vqrdmlsh_s32(C.int32x2_t(v0), C.int32x2_t(v1), C.int32x2_t(v2)))
}

func VqrdmlshS16(v0 Int16X4, v1 Int16X4, v2 Int16X4) Int16X4 {
	return Int16X4(C.vqrdmlsh_s16(C.int16x4_t(v0), C.int16x4_t(v1), C.int16x4_t(v2)))
}

func VabdqF64(v0 Float64X2, v1 Float64X2) Float64X2 {
	return Float64X2(C.vabdq_f64(C.float64x2_t(v0), C.float64x2_t(v1)))
}

func VabdF64(v0 Float64X1, v1 Float64X1) Float64X1 {
	return Float64X1(C.vabd_f64(C.float64x1_t(v0), C.float64x1_t(v1)))
}

func VabddF64(v0 float64, v1 float64) float64 {
	return float64(C.vabdd_f64(C.float64_t(v0), C.float64_t(v1)))
}

func VabdsF32(v0 float32, v1 float32) float32 {
	return float32(C.vabds_f32(C.float32_t(v0), C.float32_t(v1)))
}

func VabsqF64(v0 Float64X2) Float64X2 {
	return Float64X2(C.vabsq_f64(C.float64x2_t(v0)))
}

func VabsqS64(v0 Int64X2) Int64X2 {
	return Int64X2(C.vabsq_s64(C.int64x2_t(v0)))
}

func VabsF64(v0 Float64X1) Float64X1 {
	return Float64X1(C.vabs_f64(C.float64x1_t(v0)))
}

func VabsS64(v0 Int64X1) Int64X1 {
	return Int64X1(C.vabs_s64(C.int64x1_t(v0)))
}

func VabsdS64(v0 int64) int64 {
	return int64(C.vabsd_s64(C.int64_t(v0)))
}

func VaddqF64(v0 Float64X2, v1 Float64X2) Float64X2 {
	return Float64X2(C.vaddq_f64(C.float64x2_t(v0), C.float64x2_t(v1)))
}

func VaddF64(v0 Float64X1, v1 Float64X1) Float64X1 {
	return Float64X1(C.vadd_f64(C.float64x1_t(v0), C.float64x1_t(v1)))
}

func VadddU64(v0 uint64, v1 uint64) uint64 {
	return uint64(C.vaddd_u64(C.uint64_t(v0), C.uint64_t(v1)))
}

func VadddS64(v0 int64, v1 int64) int64 {
	return int64(C.vaddd_s64(C.int64_t(v0), C.int64_t(v1)))
}

func VaddqP128(v0 Poly128, v1 Poly128) Poly128 {
	return Poly128(C.vaddq_p128(C.poly128_t(v0), C.poly128_t(v1)))
}

func VaddhnHighU32(v0 Uint16X4, v1 Uint32X4, v2 Uint32X4) Uint16X8 {
	return Uint16X8(C.vaddhn_high_u32(C.uint16x4_t(v0), C.uint32x4_t(v1), C.uint32x4_t(v2)))
}

func VaddhnHighU64(v0 Uint32X2, v1 Uint64X2, v2 Uint64X2) Uint32X4 {
	return Uint32X4(C.vaddhn_high_u64(C.uint32x2_t(v0), C.uint64x2_t(v1), C.uint64x2_t(v2)))
}

func VaddhnHighU16(v0 Uint8X8, v1 Uint16X8, v2 Uint16X8) Uint8X16 {
	return Uint8X16(C.vaddhn_high_u16(C.uint8x8_t(v0), C.uint16x8_t(v1), C.uint16x8_t(v2)))
}

func VaddhnHighS32(v0 Int16X4, v1 Int32X4, v2 Int32X4) Int16X8 {
	return Int16X8(C.vaddhn_high_s32(C.int16x4_t(v0), C.int32x4_t(v1), C.int32x4_t(v2)))
}

func VaddhnHighS64(v0 Int32X2, v1 Int64X2, v2 Int64X2) Int32X4 {
	return Int32X4(C.vaddhn_high_s64(C.int32x2_t(v0), C.int64x2_t(v1), C.int64x2_t(v2)))
}

func VaddhnHighS16(v0 Int8X8, v1 Int16X8, v2 Int16X8) Int8X16 {
	return Int8X16(C.vaddhn_high_s16(C.int8x8_t(v0), C.int16x8_t(v1), C.int16x8_t(v2)))
}

func VaddlvqU8(v0 Uint8X16) uint16 {
	return uint16(C.vaddlvq_u8(C.uint8x16_t(v0)))
}

func VaddlvqU32(v0 Uint32X4) uint64 {
	return uint64(C.vaddlvq_u32(C.uint32x4_t(v0)))
}

func VaddlvqU16(v0 Uint16X8) uint32 {
	return uint32(C.vaddlvq_u16(C.uint16x8_t(v0)))
}

func VaddlvqS8(v0 Int8X16) int16 {
	return int16(C.vaddlvq_s8(C.int8x16_t(v0)))
}

func VaddlvqS32(v0 Int32X4) int64 {
	return int64(C.vaddlvq_s32(C.int32x4_t(v0)))
}

func VaddlvqS16(v0 Int16X8) int32 {
	return int32(C.vaddlvq_s16(C.int16x8_t(v0)))
}

func VaddlvU8(v0 Uint8X8) uint16 {
	return uint16(C.vaddlv_u8(C.uint8x8_t(v0)))
}

func VaddlvU32(v0 Uint32X2) uint64 {
	return uint64(C.vaddlv_u32(C.uint32x2_t(v0)))
}

func VaddlvU16(v0 Uint16X4) uint32 {
	return uint32(C.vaddlv_u16(C.uint16x4_t(v0)))
}

func VaddlvS8(v0 Int8X8) int16 {
	return int16(C.vaddlv_s8(C.int8x8_t(v0)))
}

func VaddlvS32(v0 Int32X2) int64 {
	return int64(C.vaddlv_s32(C.int32x2_t(v0)))
}

func VaddlvS16(v0 Int16X4) int32 {
	return int32(C.vaddlv_s16(C.int16x4_t(v0)))
}

func VaddvqU8(v0 Uint8X16) uint8 {
	return uint8(C.vaddvq_u8(C.uint8x16_t(v0)))
}

func VaddvqU32(v0 Uint32X4) uint32 {
	return uint32(C.vaddvq_u32(C.uint32x4_t(v0)))
}

func VaddvqU64(v0 Uint64X2) uint64 {
	return uint64(C.vaddvq_u64(C.uint64x2_t(v0)))
}

func VaddvqU16(v0 Uint16X8) uint16 {
	return uint16(C.vaddvq_u16(C.uint16x8_t(v0)))
}

func VaddvqS8(v0 Int8X16) int8 {
	return int8(C.vaddvq_s8(C.int8x16_t(v0)))
}

func VaddvqF64(v0 Float64X2) float64 {
	return float64(C.vaddvq_f64(C.float64x2_t(v0)))
}

func VaddvqF32(v0 Float32X4) float32 {
	return float32(C.vaddvq_f32(C.float32x4_t(v0)))
}

func VaddvqS32(v0 Int32X4) int32 {
	return int32(C.vaddvq_s32(C.int32x4_t(v0)))
}

func VaddvqS64(v0 Int64X2) int64 {
	return int64(C.vaddvq_s64(C.int64x2_t(v0)))
}

func VaddvqS16(v0 Int16X8) int16 {
	return int16(C.vaddvq_s16(C.int16x8_t(v0)))
}

func VaddvU8(v0 Uint8X8) uint8 {
	return uint8(C.vaddv_u8(C.uint8x8_t(v0)))
}

func VaddvU32(v0 Uint32X2) uint32 {
	return uint32(C.vaddv_u32(C.uint32x2_t(v0)))
}

func VaddvU16(v0 Uint16X4) uint16 {
	return uint16(C.vaddv_u16(C.uint16x4_t(v0)))
}

func VaddvS8(v0 Int8X8) int8 {
	return int8(C.vaddv_s8(C.int8x8_t(v0)))
}

func VaddvF32(v0 Float32X2) float32 {
	return float32(C.vaddv_f32(C.float32x2_t(v0)))
}

func VaddvS32(v0 Int32X2) int32 {
	return int32(C.vaddv_s32(C.int32x2_t(v0)))
}

func VaddvS16(v0 Int16X4) int16 {
	return int16(C.vaddv_s16(C.int16x4_t(v0)))
}

func VbslP64(v0 Uint64X1, v1 Poly64X1, v2 Poly64X1) Poly64X1 {
	return Poly64X1(C.vbsl_p64(C.uint64x1_t(v0), C.poly64x1_t(v1), C.poly64x1_t(v2)))
}

func VbslqP64(v0 Uint64X2, v1 Poly64X2, v2 Poly64X2) Poly64X2 {
	return Poly64X2(C.vbslq_p64(C.uint64x2_t(v0), C.poly64x2_t(v1), C.poly64x2_t(v2)))
}

func VbslqF64(v0 Uint64X2, v1 Float64X2, v2 Float64X2) Float64X2 {
	return Float64X2(C.vbslq_f64(C.uint64x2_t(v0), C.float64x2_t(v1), C.float64x2_t(v2)))
}

func VbslF64(v0 Uint64X1, v1 Float64X1, v2 Float64X1) Float64X1 {
	return Float64X1(C.vbsl_f64(C.uint64x1_t(v0), C.float64x1_t(v1), C.float64x1_t(v2)))
}

func VcageqF64(v0 Float64X2, v1 Float64X2) Uint64X2 {
	return Uint64X2(C.vcageq_f64(C.float64x2_t(v0), C.float64x2_t(v1)))
}

func VcageF64(v0 Float64X1, v1 Float64X1) Uint64X1 {
	return Uint64X1(C.vcage_f64(C.float64x1_t(v0), C.float64x1_t(v1)))
}

func VcagedF64(v0 float64, v1 float64) uint64 {
	return uint64(C.vcaged_f64(C.float64_t(v0), C.float64_t(v1)))
}

func VcagesF32(v0 float32, v1 float32) uint32 {
	return uint32(C.vcages_f32(C.float32_t(v0), C.float32_t(v1)))
}

func VcagtqF64(v0 Float64X2, v1 Float64X2) Uint64X2 {
	return Uint64X2(C.vcagtq_f64(C.float64x2_t(v0), C.float64x2_t(v1)))
}

func VcagtF64(v0 Float64X1, v1 Float64X1) Uint64X1 {
	return Uint64X1(C.vcagt_f64(C.float64x1_t(v0), C.float64x1_t(v1)))
}

func VcagtdF64(v0 float64, v1 float64) uint64 {
	return uint64(C.vcagtd_f64(C.float64_t(v0), C.float64_t(v1)))
}

func VcagtsF32(v0 float32, v1 float32) uint32 {
	return uint32(C.vcagts_f32(C.float32_t(v0), C.float32_t(v1)))
}

func VcaleqF64(v0 Float64X2, v1 Float64X2) Uint64X2 {
	return Uint64X2(C.vcaleq_f64(C.float64x2_t(v0), C.float64x2_t(v1)))
}

func VcaleF64(v0 Float64X1, v1 Float64X1) Uint64X1 {
	return Uint64X1(C.vcale_f64(C.float64x1_t(v0), C.float64x1_t(v1)))
}

func VcaledF64(v0 float64, v1 float64) uint64 {
	return uint64(C.vcaled_f64(C.float64_t(v0), C.float64_t(v1)))
}

func VcalesF32(v0 float32, v1 float32) uint32 {
	return uint32(C.vcales_f32(C.float32_t(v0), C.float32_t(v1)))
}

func VcaltqF64(v0 Float64X2, v1 Float64X2) Uint64X2 {
	return Uint64X2(C.vcaltq_f64(C.float64x2_t(v0), C.float64x2_t(v1)))
}

func VcaltF64(v0 Float64X1, v1 Float64X1) Uint64X1 {
	return Uint64X1(C.vcalt_f64(C.float64x1_t(v0), C.float64x1_t(v1)))
}

func VcaltdF64(v0 float64, v1 float64) uint64 {
	return uint64(C.vcaltd_f64(C.float64_t(v0), C.float64_t(v1)))
}

func VcaltsF32(v0 float32, v1 float32) uint32 {
	return uint32(C.vcalts_f32(C.float32_t(v0), C.float32_t(v1)))
}

func VceqP64(v0 Poly64X1, v1 Poly64X1) Uint64X1 {
	return Uint64X1(C.vceq_p64(C.poly64x1_t(v0), C.poly64x1_t(v1)))
}

func VceqqP64(v0 Poly64X2, v1 Poly64X2) Uint64X2 {
	return Uint64X2(C.vceqq_p64(C.poly64x2_t(v0), C.poly64x2_t(v1)))
}

func VceqqU64(v0 Uint64X2, v1 Uint64X2) Uint64X2 {
	return Uint64X2(C.vceqq_u64(C.uint64x2_t(v0), C.uint64x2_t(v1)))
}

func VceqqF64(v0 Float64X2, v1 Float64X2) Uint64X2 {
	return Uint64X2(C.vceqq_f64(C.float64x2_t(v0), C.float64x2_t(v1)))
}

func VceqqS64(v0 Int64X2, v1 Int64X2) Uint64X2 {
	return Uint64X2(C.vceqq_s64(C.int64x2_t(v0), C.int64x2_t(v1)))
}

func VceqU64(v0 Uint64X1, v1 Uint64X1) Uint64X1 {
	return Uint64X1(C.vceq_u64(C.uint64x1_t(v0), C.uint64x1_t(v1)))
}

func VceqF64(v0 Float64X1, v1 Float64X1) Uint64X1 {
	return Uint64X1(C.vceq_f64(C.float64x1_t(v0), C.float64x1_t(v1)))
}

func VceqS64(v0 Int64X1, v1 Int64X1) Uint64X1 {
	return Uint64X1(C.vceq_s64(C.int64x1_t(v0), C.int64x1_t(v1)))
}

func VceqdU64(v0 uint64, v1 uint64) uint64 {
	return uint64(C.vceqd_u64(C.uint64_t(v0), C.uint64_t(v1)))
}

func VceqdS64(v0 int64, v1 int64) uint64 {
	return uint64(C.vceqd_s64(C.int64_t(v0), C.int64_t(v1)))
}

func VceqdF64(v0 float64, v1 float64) uint64 {
	return uint64(C.vceqd_f64(C.float64_t(v0), C.float64_t(v1)))
}

func VceqsF32(v0 float32, v1 float32) uint32 {
	return uint32(C.vceqs_f32(C.float32_t(v0), C.float32_t(v1)))
}

func VceqzP8(v0 Poly8X8) Uint8X8 {
	return Uint8X8(C.vceqz_p8(C.poly8x8_t(v0)))
}

func VceqzP64(v0 Poly64X1) Uint64X1 {
	return Uint64X1(C.vceqz_p64(C.poly64x1_t(v0)))
}

func VceqzqP8(v0 Poly8X16) Uint8X16 {
	return Uint8X16(C.vceqzq_p8(C.poly8x16_t(v0)))
}

func VceqzqP64(v0 Poly64X2) Uint64X2 {
	return Uint64X2(C.vceqzq_p64(C.poly64x2_t(v0)))
}

func VceqzqU8(v0 Uint8X16) Uint8X16 {
	return Uint8X16(C.vceqzq_u8(C.uint8x16_t(v0)))
}

func VceqzqU32(v0 Uint32X4) Uint32X4 {
	return Uint32X4(C.vceqzq_u32(C.uint32x4_t(v0)))
}

func VceqzqU64(v0 Uint64X2) Uint64X2 {
	return Uint64X2(C.vceqzq_u64(C.uint64x2_t(v0)))
}

func VceqzqU16(v0 Uint16X8) Uint16X8 {
	return Uint16X8(C.vceqzq_u16(C.uint16x8_t(v0)))
}

func VceqzqS8(v0 Int8X16) Uint8X16 {
	return Uint8X16(C.vceqzq_s8(C.int8x16_t(v0)))
}

func VceqzqF64(v0 Float64X2) Uint64X2 {
	return Uint64X2(C.vceqzq_f64(C.float64x2_t(v0)))
}

func VceqzqF32(v0 Float32X4) Uint32X4 {
	return Uint32X4(C.vceqzq_f32(C.float32x4_t(v0)))
}

func VceqzqS32(v0 Int32X4) Uint32X4 {
	return Uint32X4(C.vceqzq_s32(C.int32x4_t(v0)))
}

func VceqzqS64(v0 Int64X2) Uint64X2 {
	return Uint64X2(C.vceqzq_s64(C.int64x2_t(v0)))
}

func VceqzqS16(v0 Int16X8) Uint16X8 {
	return Uint16X8(C.vceqzq_s16(C.int16x8_t(v0)))
}

func VceqzU8(v0 Uint8X8) Uint8X8 {
	return Uint8X8(C.vceqz_u8(C.uint8x8_t(v0)))
}

func VceqzU32(v0 Uint32X2) Uint32X2 {
	return Uint32X2(C.vceqz_u32(C.uint32x2_t(v0)))
}

func VceqzU64(v0 Uint64X1) Uint64X1 {
	return Uint64X1(C.vceqz_u64(C.uint64x1_t(v0)))
}

func VceqzU16(v0 Uint16X4) Uint16X4 {
	return Uint16X4(C.vceqz_u16(C.uint16x4_t(v0)))
}

func VceqzS8(v0 Int8X8) Uint8X8 {
	return Uint8X8(C.vceqz_s8(C.int8x8_t(v0)))
}

func VceqzF64(v0 Float64X1) Uint64X1 {
	return Uint64X1(C.vceqz_f64(C.float64x1_t(v0)))
}

func VceqzF32(v0 Float32X2) Uint32X2 {
	return Uint32X2(C.vceqz_f32(C.float32x2_t(v0)))
}

func VceqzS32(v0 Int32X2) Uint32X2 {
	return Uint32X2(C.vceqz_s32(C.int32x2_t(v0)))
}

func VceqzS64(v0 Int64X1) Uint64X1 {
	return Uint64X1(C.vceqz_s64(C.int64x1_t(v0)))
}

func VceqzS16(v0 Int16X4) Uint16X4 {
	return Uint16X4(C.vceqz_s16(C.int16x4_t(v0)))
}

func VceqzdU64(v0 uint64) uint64 {
	return uint64(C.vceqzd_u64(C.uint64_t(v0)))
}

func VceqzdS64(v0 int64) uint64 {
	return uint64(C.vceqzd_s64(C.int64_t(v0)))
}

func VceqzdF64(v0 float64) uint64 {
	return uint64(C.vceqzd_f64(C.float64_t(v0)))
}

func VceqzsF32(v0 float32) uint32 {
	return uint32(C.vceqzs_f32(C.float32_t(v0)))
}

func VcgeqU64(v0 Uint64X2, v1 Uint64X2) Uint64X2 {
	return Uint64X2(C.vcgeq_u64(C.uint64x2_t(v0), C.uint64x2_t(v1)))
}

func VcgeqF64(v0 Float64X2, v1 Float64X2) Uint64X2 {
	return Uint64X2(C.vcgeq_f64(C.float64x2_t(v0), C.float64x2_t(v1)))
}

func VcgeqS64(v0 Int64X2, v1 Int64X2) Uint64X2 {
	return Uint64X2(C.vcgeq_s64(C.int64x2_t(v0), C.int64x2_t(v1)))
}

func VcgeU64(v0 Uint64X1, v1 Uint64X1) Uint64X1 {
	return Uint64X1(C.vcge_u64(C.uint64x1_t(v0), C.uint64x1_t(v1)))
}

func VcgeF64(v0 Float64X1, v1 Float64X1) Uint64X1 {
	return Uint64X1(C.vcge_f64(C.float64x1_t(v0), C.float64x1_t(v1)))
}

func VcgeS64(v0 Int64X1, v1 Int64X1) Uint64X1 {
	return Uint64X1(C.vcge_s64(C.int64x1_t(v0), C.int64x1_t(v1)))
}

func VcgedS64(v0 int64, v1 int64) uint64 {
	return uint64(C.vcged_s64(C.int64_t(v0), C.int64_t(v1)))
}

func VcgedU64(v0 uint64, v1 uint64) uint64 {
	return uint64(C.vcged_u64(C.uint64_t(v0), C.uint64_t(v1)))
}

func VcgedF64(v0 float64, v1 float64) uint64 {
	return uint64(C.vcged_f64(C.float64_t(v0), C.float64_t(v1)))
}

func VcgesF32(v0 float32, v1 float32) uint32 {
	return uint32(C.vcges_f32(C.float32_t(v0), C.float32_t(v1)))
}

func VcgezqS8(v0 Int8X16) Uint8X16 {
	return Uint8X16(C.vcgezq_s8(C.int8x16_t(v0)))
}

func VcgezqF64(v0 Float64X2) Uint64X2 {
	return Uint64X2(C.vcgezq_f64(C.float64x2_t(v0)))
}

func VcgezqF32(v0 Float32X4) Uint32X4 {
	return Uint32X4(C.vcgezq_f32(C.float32x4_t(v0)))
}

func VcgezqS32(v0 Int32X4) Uint32X4 {
	return Uint32X4(C.vcgezq_s32(C.int32x4_t(v0)))
}

func VcgezqS64(v0 Int64X2) Uint64X2 {
	return Uint64X2(C.vcgezq_s64(C.int64x2_t(v0)))
}

func VcgezqS16(v0 Int16X8) Uint16X8 {
	return Uint16X8(C.vcgezq_s16(C.int16x8_t(v0)))
}

func VcgezS8(v0 Int8X8) Uint8X8 {
	return Uint8X8(C.vcgez_s8(C.int8x8_t(v0)))
}

func VcgezF64(v0 Float64X1) Uint64X1 {
	return Uint64X1(C.vcgez_f64(C.float64x1_t(v0)))
}

func VcgezF32(v0 Float32X2) Uint32X2 {
	return Uint32X2(C.vcgez_f32(C.float32x2_t(v0)))
}

func VcgezS32(v0 Int32X2) Uint32X2 {
	return Uint32X2(C.vcgez_s32(C.int32x2_t(v0)))
}

func VcgezS64(v0 Int64X1) Uint64X1 {
	return Uint64X1(C.vcgez_s64(C.int64x1_t(v0)))
}

func VcgezS16(v0 Int16X4) Uint16X4 {
	return Uint16X4(C.vcgez_s16(C.int16x4_t(v0)))
}

func VcgezdS64(v0 int64) uint64 {
	return uint64(C.vcgezd_s64(C.int64_t(v0)))
}

func VcgezdF64(v0 float64) uint64 {
	return uint64(C.vcgezd_f64(C.float64_t(v0)))
}

func VcgezsF32(v0 float32) uint32 {
	return uint32(C.vcgezs_f32(C.float32_t(v0)))
}

func VcgtqU64(v0 Uint64X2, v1 Uint64X2) Uint64X2 {
	return Uint64X2(C.vcgtq_u64(C.uint64x2_t(v0), C.uint64x2_t(v1)))
}

func VcgtqF64(v0 Float64X2, v1 Float64X2) Uint64X2 {
	return Uint64X2(C.vcgtq_f64(C.float64x2_t(v0), C.float64x2_t(v1)))
}

func VcgtqS64(v0 Int64X2, v1 Int64X2) Uint64X2 {
	return Uint64X2(C.vcgtq_s64(C.int64x2_t(v0), C.int64x2_t(v1)))
}

func VcgtU64(v0 Uint64X1, v1 Uint64X1) Uint64X1 {
	return Uint64X1(C.vcgt_u64(C.uint64x1_t(v0), C.uint64x1_t(v1)))
}

func VcgtF64(v0 Float64X1, v1 Float64X1) Uint64X1 {
	return Uint64X1(C.vcgt_f64(C.float64x1_t(v0), C.float64x1_t(v1)))
}

func VcgtS64(v0 Int64X1, v1 Int64X1) Uint64X1 {
	return Uint64X1(C.vcgt_s64(C.int64x1_t(v0), C.int64x1_t(v1)))
}

func VcgtdS64(v0 int64, v1 int64) uint64 {
	return uint64(C.vcgtd_s64(C.int64_t(v0), C.int64_t(v1)))
}

func VcgtdU64(v0 uint64, v1 uint64) uint64 {
	return uint64(C.vcgtd_u64(C.uint64_t(v0), C.uint64_t(v1)))
}

func VcgtdF64(v0 float64, v1 float64) uint64 {
	return uint64(C.vcgtd_f64(C.float64_t(v0), C.float64_t(v1)))
}

func VcgtsF32(v0 float32, v1 float32) uint32 {
	return uint32(C.vcgts_f32(C.float32_t(v0), C.float32_t(v1)))
}

func VcgtzqS8(v0 Int8X16) Uint8X16 {
	return Uint8X16(C.vcgtzq_s8(C.int8x16_t(v0)))
}

func VcgtzqF64(v0 Float64X2) Uint64X2 {
	return Uint64X2(C.vcgtzq_f64(C.float64x2_t(v0)))
}

func VcgtzqF32(v0 Float32X4) Uint32X4 {
	return Uint32X4(C.vcgtzq_f32(C.float32x4_t(v0)))
}

func VcgtzqS32(v0 Int32X4) Uint32X4 {
	return Uint32X4(C.vcgtzq_s32(C.int32x4_t(v0)))
}

func VcgtzqS64(v0 Int64X2) Uint64X2 {
	return Uint64X2(C.vcgtzq_s64(C.int64x2_t(v0)))
}

func VcgtzqS16(v0 Int16X8) Uint16X8 {
	return Uint16X8(C.vcgtzq_s16(C.int16x8_t(v0)))
}

func VcgtzS8(v0 Int8X8) Uint8X8 {
	return Uint8X8(C.vcgtz_s8(C.int8x8_t(v0)))
}

func VcgtzF64(v0 Float64X1) Uint64X1 {
	return Uint64X1(C.vcgtz_f64(C.float64x1_t(v0)))
}

func VcgtzF32(v0 Float32X2) Uint32X2 {
	return Uint32X2(C.vcgtz_f32(C.float32x2_t(v0)))
}

func VcgtzS32(v0 Int32X2) Uint32X2 {
	return Uint32X2(C.vcgtz_s32(C.int32x2_t(v0)))
}

func VcgtzS64(v0 Int64X1) Uint64X1 {
	return Uint64X1(C.vcgtz_s64(C.int64x1_t(v0)))
}

func VcgtzS16(v0 Int16X4) Uint16X4 {
	return Uint16X4(C.vcgtz_s16(C.int16x4_t(v0)))
}

func VcgtzdS64(v0 int64) uint64 {
	return uint64(C.vcgtzd_s64(C.int64_t(v0)))
}

func VcgtzdF64(v0 float64) uint64 {
	return uint64(C.vcgtzd_f64(C.float64_t(v0)))
}

func VcgtzsF32(v0 float32) uint32 {
	return uint32(C.vcgtzs_f32(C.float32_t(v0)))
}

func VcleqU64(v0 Uint64X2, v1 Uint64X2) Uint64X2 {
	return Uint64X2(C.vcleq_u64(C.uint64x2_t(v0), C.uint64x2_t(v1)))
}

func VcleqF64(v0 Float64X2, v1 Float64X2) Uint64X2 {
	return Uint64X2(C.vcleq_f64(C.float64x2_t(v0), C.float64x2_t(v1)))
}

func VcleqS64(v0 Int64X2, v1 Int64X2) Uint64X2 {
	return Uint64X2(C.vcleq_s64(C.int64x2_t(v0), C.int64x2_t(v1)))
}

func VcleU64(v0 Uint64X1, v1 Uint64X1) Uint64X1 {
	return Uint64X1(C.vcle_u64(C.uint64x1_t(v0), C.uint64x1_t(v1)))
}

func VcleF64(v0 Float64X1, v1 Float64X1) Uint64X1 {
	return Uint64X1(C.vcle_f64(C.float64x1_t(v0), C.float64x1_t(v1)))
}

func VcleS64(v0 Int64X1, v1 Int64X1) Uint64X1 {
	return Uint64X1(C.vcle_s64(C.int64x1_t(v0), C.int64x1_t(v1)))
}

func VcledU64(v0 uint64, v1 uint64) uint64 {
	return uint64(C.vcled_u64(C.uint64_t(v0), C.uint64_t(v1)))
}

func VcledS64(v0 int64, v1 int64) uint64 {
	return uint64(C.vcled_s64(C.int64_t(v0), C.int64_t(v1)))
}

func VcledF64(v0 float64, v1 float64) uint64 {
	return uint64(C.vcled_f64(C.float64_t(v0), C.float64_t(v1)))
}

func VclesF32(v0 float32, v1 float32) uint32 {
	return uint32(C.vcles_f32(C.float32_t(v0), C.float32_t(v1)))
}

func VclezqS8(v0 Int8X16) Uint8X16 {
	return Uint8X16(C.vclezq_s8(C.int8x16_t(v0)))
}

func VclezqF64(v0 Float64X2) Uint64X2 {
	return Uint64X2(C.vclezq_f64(C.float64x2_t(v0)))
}

func VclezqF32(v0 Float32X4) Uint32X4 {
	return Uint32X4(C.vclezq_f32(C.float32x4_t(v0)))
}

func VclezqS32(v0 Int32X4) Uint32X4 {
	return Uint32X4(C.vclezq_s32(C.int32x4_t(v0)))
}

func VclezqS64(v0 Int64X2) Uint64X2 {
	return Uint64X2(C.vclezq_s64(C.int64x2_t(v0)))
}

func VclezqS16(v0 Int16X8) Uint16X8 {
	return Uint16X8(C.vclezq_s16(C.int16x8_t(v0)))
}

func VclezS8(v0 Int8X8) Uint8X8 {
	return Uint8X8(C.vclez_s8(C.int8x8_t(v0)))
}

func VclezF64(v0 Float64X1) Uint64X1 {
	return Uint64X1(C.vclez_f64(C.float64x1_t(v0)))
}

func VclezF32(v0 Float32X2) Uint32X2 {
	return Uint32X2(C.vclez_f32(C.float32x2_t(v0)))
}

func VclezS32(v0 Int32X2) Uint32X2 {
	return Uint32X2(C.vclez_s32(C.int32x2_t(v0)))
}

func VclezS64(v0 Int64X1) Uint64X1 {
	return Uint64X1(C.vclez_s64(C.int64x1_t(v0)))
}

func VclezS16(v0 Int16X4) Uint16X4 {
	return Uint16X4(C.vclez_s16(C.int16x4_t(v0)))
}

func VclezdS64(v0 int64) uint64 {
	return uint64(C.vclezd_s64(C.int64_t(v0)))
}

func VclezdF64(v0 float64) uint64 {
	return uint64(C.vclezd_f64(C.float64_t(v0)))
}

func VclezsF32(v0 float32) uint32 {
	return uint32(C.vclezs_f32(C.float32_t(v0)))
}

func VcltqU64(v0 Uint64X2, v1 Uint64X2) Uint64X2 {
	return Uint64X2(C.vcltq_u64(C.uint64x2_t(v0), C.uint64x2_t(v1)))
}

func VcltqF64(v0 Float64X2, v1 Float64X2) Uint64X2 {
	return Uint64X2(C.vcltq_f64(C.float64x2_t(v0), C.float64x2_t(v1)))
}

func VcltqS64(v0 Int64X2, v1 Int64X2) Uint64X2 {
	return Uint64X2(C.vcltq_s64(C.int64x2_t(v0), C.int64x2_t(v1)))
}

func VcltU64(v0 Uint64X1, v1 Uint64X1) Uint64X1 {
	return Uint64X1(C.vclt_u64(C.uint64x1_t(v0), C.uint64x1_t(v1)))
}

func VcltF64(v0 Float64X1, v1 Float64X1) Uint64X1 {
	return Uint64X1(C.vclt_f64(C.float64x1_t(v0), C.float64x1_t(v1)))
}

func VcltS64(v0 Int64X1, v1 Int64X1) Uint64X1 {
	return Uint64X1(C.vclt_s64(C.int64x1_t(v0), C.int64x1_t(v1)))
}

func VcltdU64(v0 uint64, v1 uint64) uint64 {
	return uint64(C.vcltd_u64(C.uint64_t(v0), C.uint64_t(v1)))
}

func VcltdS64(v0 int64, v1 int64) uint64 {
	return uint64(C.vcltd_s64(C.int64_t(v0), C.int64_t(v1)))
}

func VcltdF64(v0 float64, v1 float64) uint64 {
	return uint64(C.vcltd_f64(C.float64_t(v0), C.float64_t(v1)))
}

func VcltsF32(v0 float32, v1 float32) uint32 {
	return uint32(C.vclts_f32(C.float32_t(v0), C.float32_t(v1)))
}

func VcltzqS8(v0 Int8X16) Uint8X16 {
	return Uint8X16(C.vcltzq_s8(C.int8x16_t(v0)))
}

func VcltzqF64(v0 Float64X2) Uint64X2 {
	return Uint64X2(C.vcltzq_f64(C.float64x2_t(v0)))
}

func VcltzqF32(v0 Float32X4) Uint32X4 {
	return Uint32X4(C.vcltzq_f32(C.float32x4_t(v0)))
}

func VcltzqS32(v0 Int32X4) Uint32X4 {
	return Uint32X4(C.vcltzq_s32(C.int32x4_t(v0)))
}

func VcltzqS64(v0 Int64X2) Uint64X2 {
	return Uint64X2(C.vcltzq_s64(C.int64x2_t(v0)))
}

func VcltzqS16(v0 Int16X8) Uint16X8 {
	return Uint16X8(C.vcltzq_s16(C.int16x8_t(v0)))
}

func VcltzS8(v0 Int8X8) Uint8X8 {
	return Uint8X8(C.vcltz_s8(C.int8x8_t(v0)))
}

func VcltzF64(v0 Float64X1) Uint64X1 {
	return Uint64X1(C.vcltz_f64(C.float64x1_t(v0)))
}

func VcltzF32(v0 Float32X2) Uint32X2 {
	return Uint32X2(C.vcltz_f32(C.float32x2_t(v0)))
}

func VcltzS32(v0 Int32X2) Uint32X2 {
	return Uint32X2(C.vcltz_s32(C.int32x2_t(v0)))
}

func VcltzS64(v0 Int64X1) Uint64X1 {
	return Uint64X1(C.vcltz_s64(C.int64x1_t(v0)))
}

func VcltzS16(v0 Int16X4) Uint16X4 {
	return Uint16X4(C.vcltz_s16(C.int16x4_t(v0)))
}

func VcltzdS64(v0 int64) uint64 {
	return uint64(C.vcltzd_s64(C.int64_t(v0)))
}

func VcltzdF64(v0 float64) uint64 {
	return uint64(C.vcltzd_f64(C.float64_t(v0)))
}

func VcltzsF32(v0 float32) uint32 {
	return uint32(C.vcltzs_f32(C.float32_t(v0)))
}

func VcombineP64(v0 Poly64X1, v1 Poly64X1) Poly64X2 {
	return Poly64X2(C.vcombine_p64(C.poly64x1_t(v0), C.poly64x1_t(v1)))
}

func VcombineF64(v0 Float64X1, v1 Float64X1) Float64X2 {
	return Float64X2(C.vcombine_f64(C.float64x1_t(v0), C.float64x1_t(v1)))
}

func VcvtsF32S32(v0 int32) float32 {
	return float32(C.vcvts_f32_s32(C.int32_t(v0)))
}

func VcvtsF32U32(v0 uint32) float32 {
	return float32(C.vcvts_f32_u32(C.uint32_t(v0)))
}

func VcvtF32F64(v0 Float64X2) Float32X2 {
	return Float32X2(C.vcvt_f32_f64(C.float64x2_t(v0)))
}

func VcvtdF64S64(v0 int64) float64 {
	return float64(C.vcvtd_f64_s64(C.int64_t(v0)))
}

func VcvtdF64U64(v0 uint64) float64 {
	return float64(C.vcvtd_f64_u64(C.uint64_t(v0)))
}

func VcvtqF64U64(v0 Uint64X2) Float64X2 {
	return Float64X2(C.vcvtq_f64_u64(C.uint64x2_t(v0)))
}

func VcvtqF64S64(v0 Int64X2) Float64X2 {
	return Float64X2(C.vcvtq_f64_s64(C.int64x2_t(v0)))
}

func VcvtF64U64(v0 Uint64X1) Float64X1 {
	return Float64X1(C.vcvt_f64_u64(C.uint64x1_t(v0)))
}

func VcvtF64S64(v0 Int64X1) Float64X1 {
	return Float64X1(C.vcvt_f64_s64(C.int64x1_t(v0)))
}

func VcvtF64F32(v0 Float32X2) Float64X2 {
	return Float64X2(C.vcvt_f64_f32(C.float32x2_t(v0)))
}

func VcvtHighF32F64(v0 Float32X2, v1 Float64X2) Float32X4 {
	return Float32X4(C.vcvt_high_f32_f64(C.float32x2_t(v0), C.float64x2_t(v1)))
}

func VcvtHighF64F32(v0 Float32X4) Float64X2 {
	return Float64X2(C.vcvt_high_f64_f32(C.float32x4_t(v0)))
}

func VcvtsS32F32(v0 float32) int32 {
	return int32(C.vcvts_s32_f32(C.float32_t(v0)))
}

func VcvtdS64F64(v0 float64) int64 {
	return int64(C.vcvtd_s64_f64(C.float64_t(v0)))
}

func VcvtqS64F64(v0 Float64X2) Int64X2 {
	return Int64X2(C.vcvtq_s64_f64(C.float64x2_t(v0)))
}

func VcvtS64F64(v0 Float64X1) Int64X1 {
	return Int64X1(C.vcvt_s64_f64(C.float64x1_t(v0)))
}

func VcvtsU32F32(v0 float32) uint32 {
	return uint32(C.vcvts_u32_f32(C.float32_t(v0)))
}

func VcvtdU64F64(v0 float64) uint64 {
	return uint64(C.vcvtd_u64_f64(C.float64_t(v0)))
}

func VcvtqU64F64(v0 Float64X2) Uint64X2 {
	return Uint64X2(C.vcvtq_u64_f64(C.float64x2_t(v0)))
}

func VcvtU64F64(v0 Float64X1) Uint64X1 {
	return Uint64X1(C.vcvt_u64_f64(C.float64x1_t(v0)))
}

func VcvtasS32F32(v0 float32) int32 {
	return int32(C.vcvtas_s32_f32(C.float32_t(v0)))
}

func VcvtadS64F64(v0 float64) int64 {
	return int64(C.vcvtad_s64_f64(C.float64_t(v0)))
}

func VcvtasU32F32(v0 float32) uint32 {
	return uint32(C.vcvtas_u32_f32(C.float32_t(v0)))
}

func VcvtadU64F64(v0 float64) uint64 {
	return uint64(C.vcvtad_u64_f64(C.float64_t(v0)))
}

func VcvtmsS32F32(v0 float32) int32 {
	return int32(C.vcvtms_s32_f32(C.float32_t(v0)))
}

func VcvtmdS64F64(v0 float64) int64 {
	return int64(C.vcvtmd_s64_f64(C.float64_t(v0)))
}

func VcvtmsU32F32(v0 float32) uint32 {
	return uint32(C.vcvtms_u32_f32(C.float32_t(v0)))
}

func VcvtmdU64F64(v0 float64) uint64 {
	return uint64(C.vcvtmd_u64_f64(C.float64_t(v0)))
}

func VcvtnsS32F32(v0 float32) int32 {
	return int32(C.vcvtns_s32_f32(C.float32_t(v0)))
}

func VcvtndS64F64(v0 float64) int64 {
	return int64(C.vcvtnd_s64_f64(C.float64_t(v0)))
}

func VcvtnsU32F32(v0 float32) uint32 {
	return uint32(C.vcvtns_u32_f32(C.float32_t(v0)))
}

func VcvtndU64F64(v0 float64) uint64 {
	return uint64(C.vcvtnd_u64_f64(C.float64_t(v0)))
}

func VcvtpsS32F32(v0 float32) int32 {
	return int32(C.vcvtps_s32_f32(C.float32_t(v0)))
}

func VcvtpdS64F64(v0 float64) int64 {
	return int64(C.vcvtpd_s64_f64(C.float64_t(v0)))
}

func VcvtpsU32F32(v0 float32) uint32 {
	return uint32(C.vcvtps_u32_f32(C.float32_t(v0)))
}

func VcvtpdU64F64(v0 float64) uint64 {
	return uint64(C.vcvtpd_u64_f64(C.float64_t(v0)))
}

func VcvtxdF32F64(v0 float64) float32 {
	return float32(C.vcvtxd_f32_f64(C.float64_t(v0)))
}

func VcvtxF32F64(v0 Float64X2) Float32X2 {
	return Float32X2(C.vcvtx_f32_f64(C.float64x2_t(v0)))
}

func VcvtxHighF32F64(v0 Float32X2, v1 Float64X2) Float32X4 {
	return Float32X4(C.vcvtx_high_f32_f64(C.float32x2_t(v0), C.float64x2_t(v1)))
}

func VdivqF64(v0 Float64X2, v1 Float64X2) Float64X2 {
	return Float64X2(C.vdivq_f64(C.float64x2_t(v0), C.float64x2_t(v1)))
}

func VdivqF32(v0 Float32X4, v1 Float32X4) Float32X4 {
	return Float32X4(C.vdivq_f32(C.float32x4_t(v0), C.float32x4_t(v1)))
}

func VdivF64(v0 Float64X1, v1 Float64X1) Float64X1 {
	return Float64X1(C.vdiv_f64(C.float64x1_t(v0), C.float64x1_t(v1)))
}

func VdivF32(v0 Float32X2, v1 Float32X2) Float32X2 {
	return Float32X2(C.vdiv_f32(C.float32x2_t(v0), C.float32x2_t(v1)))
}

func VdupNP64(v0 Poly64) Poly64X1 {
	return Poly64X1(C.vdup_n_p64(C.poly64_t(v0)))
}

func VdupqNP64(v0 Poly64) Poly64X2 {
	return Poly64X2(C.vdupq_n_p64(C.poly64_t(v0)))
}

func VdupqNF64(v0 float64) Float64X2 {
	return Float64X2(C.vdupq_n_f64(C.float64_t(v0)))
}

func VdupNF64(v0 float64) Float64X1 {
	return Float64X1(C.vdup_n_f64(C.float64_t(v0)))
}

func VfmaqF64(v0 Float64X2, v1 Float64X2, v2 Float64X2) Float64X2 {
	return Float64X2(C.vfmaq_f64(C.float64x2_t(v0), C.float64x2_t(v1), C.float64x2_t(v2)))
}

func VfmaF64(v0 Float64X1, v1 Float64X1, v2 Float64X1) Float64X1 {
	return Float64X1(C.vfma_f64(C.float64x1_t(v0), C.float64x1_t(v1), C.float64x1_t(v2)))
}

func VfmaqNF64(v0 Float64X2, v1 Float64X2, v2 float64) Float64X2 {
	return Float64X2(C.vfmaq_n_f64(C.float64x2_t(v0), C.float64x2_t(v1), C.float64_t(v2)))
}

func VfmaNF64(v0 Float64X1, v1 Float64X1, v2 float64) Float64X1 {
	return Float64X1(C.vfma_n_f64(C.float64x1_t(v0), C.float64x1_t(v1), C.float64_t(v2)))
}

func VfmsqF64(v0 Float64X2, v1 Float64X2, v2 Float64X2) Float64X2 {
	return Float64X2(C.vfmsq_f64(C.float64x2_t(v0), C.float64x2_t(v1), C.float64x2_t(v2)))
}

func VfmsF64(v0 Float64X1, v1 Float64X1, v2 Float64X1) Float64X1 {
	return Float64X1(C.vfms_f64(C.float64x1_t(v0), C.float64x1_t(v1), C.float64x1_t(v2)))
}

func VfmsqNF64(v0 Float64X2, v1 Float64X2, v2 float64) Float64X2 {
	return Float64X2(C.vfmsq_n_f64(C.float64x2_t(v0), C.float64x2_t(v1), C.float64_t(v2)))
}

func VfmsqNF32(v0 Float32X4, v1 Float32X4, v2 float32) Float32X4 {
	return Float32X4(C.vfmsq_n_f32(C.float32x4_t(v0), C.float32x4_t(v1), C.float32_t(v2)))
}

func VfmsNF64(v0 Float64X1, v1 Float64X1, v2 float64) Float64X1 {
	return Float64X1(C.vfms_n_f64(C.float64x1_t(v0), C.float64x1_t(v1), C.float64_t(v2)))
}

func VfmsNF32(v0 Float32X2, v1 Float32X2, v2 float32) Float32X2 {
	return Float32X2(C.vfms_n_f32(C.float32x2_t(v0), C.float32x2_t(v1), C.float32_t(v2)))
}

func VgetHighP64(v0 Poly64X2) Poly64X1 {
	return Poly64X1(C.vget_high_p64(C.poly64x2_t(v0)))
}

func VgetHighF64(v0 Float64X2) Float64X1 {
	return Float64X1(C.vget_high_f64(C.float64x2_t(v0)))
}

func VgetLowP64(v0 Poly64X2) Poly64X1 {
	return Poly64X1(C.vget_low_p64(C.poly64x2_t(v0)))
}

func VgetLowF64(v0 Float64X2) Float64X1 {
	return Float64X1(C.vget_low_f64(C.float64x2_t(v0)))
}

func VmaxqF64(v0 Float64X2, v1 Float64X2) Float64X2 {
	return Float64X2(C.vmaxq_f64(C.float64x2_t(v0), C.float64x2_t(v1)))
}

func VmaxF64(v0 Float64X1, v1 Float64X1) Float64X1 {
	return Float64X1(C.vmax_f64(C.float64x1_t(v0), C.float64x1_t(v1)))
}

func VmaxnmvqF64(v0 Float64X2) float64 {
	return float64(C.vmaxnmvq_f64(C.float64x2_t(v0)))
}

func VmaxnmvqF32(v0 Float32X4) float32 {
	return float32(C.vmaxnmvq_f32(C.float32x4_t(v0)))
}

func VmaxnmvF32(v0 Float32X2) float32 {
	return float32(C.vmaxnmv_f32(C.float32x2_t(v0)))
}

func VmaxvqU8(v0 Uint8X16) uint8 {
	return uint8(C.vmaxvq_u8(C.uint8x16_t(v0)))
}

func VmaxvqU32(v0 Uint32X4) uint32 {
	return uint32(C.vmaxvq_u32(C.uint32x4_t(v0)))
}

func VmaxvqU16(v0 Uint16X8) uint16 {
	return uint16(C.vmaxvq_u16(C.uint16x8_t(v0)))
}

func VmaxvqS8(v0 Int8X16) int8 {
	return int8(C.vmaxvq_s8(C.int8x16_t(v0)))
}

func VmaxvqF64(v0 Float64X2) float64 {
	return float64(C.vmaxvq_f64(C.float64x2_t(v0)))
}

func VmaxvqF32(v0 Float32X4) float32 {
	return float32(C.vmaxvq_f32(C.float32x4_t(v0)))
}

func VmaxvqS32(v0 Int32X4) int32 {
	return int32(C.vmaxvq_s32(C.int32x4_t(v0)))
}

func VmaxvqS16(v0 Int16X8) int16 {
	return int16(C.vmaxvq_s16(C.int16x8_t(v0)))
}

func VmaxvU8(v0 Uint8X8) uint8 {
	return uint8(C.vmaxv_u8(C.uint8x8_t(v0)))
}

func VmaxvU32(v0 Uint32X2) uint32 {
	return uint32(C.vmaxv_u32(C.uint32x2_t(v0)))
}

func VmaxvU16(v0 Uint16X4) uint16 {
	return uint16(C.vmaxv_u16(C.uint16x4_t(v0)))
}

func VmaxvS8(v0 Int8X8) int8 {
	return int8(C.vmaxv_s8(C.int8x8_t(v0)))
}

func VmaxvF32(v0 Float32X2) float32 {
	return float32(C.vmaxv_f32(C.float32x2_t(v0)))
}

func VmaxvS32(v0 Int32X2) int32 {
	return int32(C.vmaxv_s32(C.int32x2_t(v0)))
}

func VmaxvS16(v0 Int16X4) int16 {
	return int16(C.vmaxv_s16(C.int16x4_t(v0)))
}

func VminqF64(v0 Float64X2, v1 Float64X2) Float64X2 {
	return Float64X2(C.vminq_f64(C.float64x2_t(v0), C.float64x2_t(v1)))
}

func VminF64(v0 Float64X1, v1 Float64X1) Float64X1 {
	return Float64X1(C.vmin_f64(C.float64x1_t(v0), C.float64x1_t(v1)))
}

func VminnmvqF64(v0 Float64X2) float64 {
	return float64(C.vminnmvq_f64(C.float64x2_t(v0)))
}

func VminnmvqF32(v0 Float32X4) float32 {
	return float32(C.vminnmvq_f32(C.float32x4_t(v0)))
}

func VminnmvF32(v0 Float32X2) float32 {
	return float32(C.vminnmv_f32(C.float32x2_t(v0)))
}

func VminvqU8(v0 Uint8X16) uint8 {
	return uint8(C.vminvq_u8(C.uint8x16_t(v0)))
}

func VminvqU32(v0 Uint32X4) uint32 {
	return uint32(C.vminvq_u32(C.uint32x4_t(v0)))
}

func VminvqU16(v0 Uint16X8) uint16 {
	return uint16(C.vminvq_u16(C.uint16x8_t(v0)))
}

func VminvqS8(v0 Int8X16) int8 {
	return int8(C.vminvq_s8(C.int8x16_t(v0)))
}

func VminvqF64(v0 Float64X2) float64 {
	return float64(C.vminvq_f64(C.float64x2_t(v0)))
}

func VminvqF32(v0 Float32X4) float32 {
	return float32(C.vminvq_f32(C.float32x4_t(v0)))
}

func VminvqS32(v0 Int32X4) int32 {
	return int32(C.vminvq_s32(C.int32x4_t(v0)))
}

func VminvqS16(v0 Int16X8) int16 {
	return int16(C.vminvq_s16(C.int16x8_t(v0)))
}

func VminvU8(v0 Uint8X8) uint8 {
	return uint8(C.vminv_u8(C.uint8x8_t(v0)))
}

func VminvU32(v0 Uint32X2) uint32 {
	return uint32(C.vminv_u32(C.uint32x2_t(v0)))
}

func VminvU16(v0 Uint16X4) uint16 {
	return uint16(C.vminv_u16(C.uint16x4_t(v0)))
}

func VminvS8(v0 Int8X8) int8 {
	return int8(C.vminv_s8(C.int8x8_t(v0)))
}

func VminvF32(v0 Float32X2) float32 {
	return float32(C.vminv_f32(C.float32x2_t(v0)))
}

func VminvS32(v0 Int32X2) int32 {
	return int32(C.vminv_s32(C.int32x2_t(v0)))
}

func VminvS16(v0 Int16X4) int16 {
	return int16(C.vminv_s16(C.int16x4_t(v0)))
}

func VmlaqF64(v0 Float64X2, v1 Float64X2, v2 Float64X2) Float64X2 {
	return Float64X2(C.vmlaq_f64(C.float64x2_t(v0), C.float64x2_t(v1), C.float64x2_t(v2)))
}

func VmlaF64(v0 Float64X1, v1 Float64X1, v2 Float64X1) Float64X1 {
	return Float64X1(C.vmla_f64(C.float64x1_t(v0), C.float64x1_t(v1), C.float64x1_t(v2)))
}

func VmlsqF64(v0 Float64X2, v1 Float64X2, v2 Float64X2) Float64X2 {
	return Float64X2(C.vmlsq_f64(C.float64x2_t(v0), C.float64x2_t(v1), C.float64x2_t(v2)))
}

func VmlsF64(v0 Float64X1, v1 Float64X1, v2 Float64X1) Float64X1 {
	return Float64X1(C.vmls_f64(C.float64x1_t(v0), C.float64x1_t(v1), C.float64x1_t(v2)))
}

func VmovNP64(v0 Poly64) Poly64X1 {
	return Poly64X1(C.vmov_n_p64(C.poly64_t(v0)))
}

func VmovqNP64(v0 Poly64) Poly64X2 {
	return Poly64X2(C.vmovq_n_p64(C.poly64_t(v0)))
}

func VmovqNF64(v0 float64) Float64X2 {
	return Float64X2(C.vmovq_n_f64(C.float64_t(v0)))
}

func VmovNF64(v0 float64) Float64X1 {
	return Float64X1(C.vmov_n_f64(C.float64_t(v0)))
}

func VmovlHighU8(v0 Uint8X16) Uint16X8 {
	return Uint16X8(C.vmovl_high_u8(C.uint8x16_t(v0)))
}

func VmovlHighU32(v0 Uint32X4) Uint64X2 {
	return Uint64X2(C.vmovl_high_u32(C.uint32x4_t(v0)))
}

func VmovlHighU16(v0 Uint16X8) Uint32X4 {
	return Uint32X4(C.vmovl_high_u16(C.uint16x8_t(v0)))
}

func VmovlHighS8(v0 Int8X16) Int16X8 {
	return Int16X8(C.vmovl_high_s8(C.int8x16_t(v0)))
}

func VmovlHighS32(v0 Int32X4) Int64X2 {
	return Int64X2(C.vmovl_high_s32(C.int32x4_t(v0)))
}

func VmovlHighS16(v0 Int16X8) Int32X4 {
	return Int32X4(C.vmovl_high_s16(C.int16x8_t(v0)))
}

func VmovnHighU32(v0 Uint16X4, v1 Uint32X4) Uint16X8 {
	return Uint16X8(C.vmovn_high_u32(C.uint16x4_t(v0), C.uint32x4_t(v1)))
}

func VmovnHighU64(v0 Uint32X2, v1 Uint64X2) Uint32X4 {
	return Uint32X4(C.vmovn_high_u64(C.uint32x2_t(v0), C.uint64x2_t(v1)))
}

func VmovnHighU16(v0 Uint8X8, v1 Uint16X8) Uint8X16 {
	return Uint8X16(C.vmovn_high_u16(C.uint8x8_t(v0), C.uint16x8_t(v1)))
}

func VmovnHighS32(v0 Int16X4, v1 Int32X4) Int16X8 {
	return Int16X8(C.vmovn_high_s32(C.int16x4_t(v0), C.int32x4_t(v1)))
}

func VmovnHighS64(v0 Int32X2, v1 Int64X2) Int32X4 {
	return Int32X4(C.vmovn_high_s64(C.int32x2_t(v0), C.int64x2_t(v1)))
}

func VmovnHighS16(v0 Int8X8, v1 Int16X8) Int8X16 {
	return Int8X16(C.vmovn_high_s16(C.int8x8_t(v0), C.int16x8_t(v1)))
}

func VmulqF64(v0 Float64X2, v1 Float64X2) Float64X2 {
	return Float64X2(C.vmulq_f64(C.float64x2_t(v0), C.float64x2_t(v1)))
}

func VmulF64(v0 Float64X1, v1 Float64X1) Float64X1 {
	return Float64X1(C.vmul_f64(C.float64x1_t(v0), C.float64x1_t(v1)))
}

func VmulNF64(v0 Float64X1, v1 float64) Float64X1 {
	return Float64X1(C.vmul_n_f64(C.float64x1_t(v0), C.float64_t(v1)))
}

func VmulqNF64(v0 Float64X2, v1 float64) Float64X2 {
	return Float64X2(C.vmulq_n_f64(C.float64x2_t(v0), C.float64_t(v1)))
}

func VmullP64(v0 Poly64, v1 Poly64) Poly128 {
	return Poly128(C.vmull_p64(C.poly64_t(v0), C.poly64_t(v1)))
}

func VmullHighP8(v0 Poly8X16, v1 Poly8X16) Poly16X8 {
	return Poly16X8(C.vmull_high_p8(C.poly8x16_t(v0), C.poly8x16_t(v1)))
}

func VmullHighU8(v0 Uint8X16, v1 Uint8X16) Uint16X8 {
	return Uint16X8(C.vmull_high_u8(C.uint8x16_t(v0), C.uint8x16_t(v1)))
}

func VmullHighU32(v0 Uint32X4, v1 Uint32X4) Uint64X2 {
	return Uint64X2(C.vmull_high_u32(C.uint32x4_t(v0), C.uint32x4_t(v1)))
}

func VmullHighU16(v0 Uint16X8, v1 Uint16X8) Uint32X4 {
	return Uint32X4(C.vmull_high_u16(C.uint16x8_t(v0), C.uint16x8_t(v1)))
}

func VmullHighS8(v0 Int8X16, v1 Int8X16) Int16X8 {
	return Int16X8(C.vmull_high_s8(C.int8x16_t(v0), C.int8x16_t(v1)))
}

func VmullHighS32(v0 Int32X4, v1 Int32X4) Int64X2 {
	return Int64X2(C.vmull_high_s32(C.int32x4_t(v0), C.int32x4_t(v1)))
}

func VmullHighS16(v0 Int16X8, v1 Int16X8) Int32X4 {
	return Int32X4(C.vmull_high_s16(C.int16x8_t(v0), C.int16x8_t(v1)))
}

func VmullHighP64(v0 Poly64X2, v1 Poly64X2) Poly128 {
	return Poly128(C.vmull_high_p64(C.poly64x2_t(v0), C.poly64x2_t(v1)))
}

func VmullHighNU32(v0 Uint32X4, v1 uint32) Uint64X2 {
	return Uint64X2(C.vmull_high_n_u32(C.uint32x4_t(v0), C.uint32_t(v1)))
}

func VmullHighNU16(v0 Uint16X8, v1 uint16) Uint32X4 {
	return Uint32X4(C.vmull_high_n_u16(C.uint16x8_t(v0), C.uint16_t(v1)))
}

func VmullHighNS32(v0 Int32X4, v1 int32) Int64X2 {
	return Int64X2(C.vmull_high_n_s32(C.int32x4_t(v0), C.int32_t(v1)))
}

func VmullHighNS16(v0 Int16X8, v1 int16) Int32X4 {
	return Int32X4(C.vmull_high_n_s16(C.int16x8_t(v0), C.int16_t(v1)))
}

func VmulxqF64(v0 Float64X2, v1 Float64X2) Float64X2 {
	return Float64X2(C.vmulxq_f64(C.float64x2_t(v0), C.float64x2_t(v1)))
}

func VmulxqF32(v0 Float32X4, v1 Float32X4) Float32X4 {
	return Float32X4(C.vmulxq_f32(C.float32x4_t(v0), C.float32x4_t(v1)))
}

func VmulxF64(v0 Float64X1, v1 Float64X1) Float64X1 {
	return Float64X1(C.vmulx_f64(C.float64x1_t(v0), C.float64x1_t(v1)))
}

func VmulxF32(v0 Float32X2, v1 Float32X2) Float32X2 {
	return Float32X2(C.vmulx_f32(C.float32x2_t(v0), C.float32x2_t(v1)))
}

func VmulxdF64(v0 float64, v1 float64) float64 {
	return float64(C.vmulxd_f64(C.float64_t(v0), C.float64_t(v1)))
}

func VmulxsF32(v0 float32, v1 float32) float32 {
	return float32(C.vmulxs_f32(C.float32_t(v0), C.float32_t(v1)))
}

func VnegqF64(v0 Float64X2) Float64X2 {
	return Float64X2(C.vnegq_f64(C.float64x2_t(v0)))
}

func VnegqS64(v0 Int64X2) Int64X2 {
	return Int64X2(C.vnegq_s64(C.int64x2_t(v0)))
}

func VnegF64(v0 Float64X1) Float64X1 {
	return Float64X1(C.vneg_f64(C.float64x1_t(v0)))
}

func VnegS64(v0 Int64X1) Int64X1 {
	return Int64X1(C.vneg_s64(C.int64x1_t(v0)))
}

func VnegdS64(v0 int64) int64 {
	return int64(C.vnegd_s64(C.int64_t(v0)))
}

func VpaddqU8(v0 Uint8X16, v1 Uint8X16) Uint8X16 {
	return Uint8X16(C.vpaddq_u8(C.uint8x16_t(v0), C.uint8x16_t(v1)))
}

func VpaddqU32(v0 Uint32X4, v1 Uint32X4) Uint32X4 {
	return Uint32X4(C.vpaddq_u32(C.uint32x4_t(v0), C.uint32x4_t(v1)))
}

func VpaddqU64(v0 Uint64X2, v1 Uint64X2) Uint64X2 {
	return Uint64X2(C.vpaddq_u64(C.uint64x2_t(v0), C.uint64x2_t(v1)))
}

func VpaddqU16(v0 Uint16X8, v1 Uint16X8) Uint16X8 {
	return Uint16X8(C.vpaddq_u16(C.uint16x8_t(v0), C.uint16x8_t(v1)))
}

func VpaddqS8(v0 Int8X16, v1 Int8X16) Int8X16 {
	return Int8X16(C.vpaddq_s8(C.int8x16_t(v0), C.int8x16_t(v1)))
}

func VpaddqF64(v0 Float64X2, v1 Float64X2) Float64X2 {
	return Float64X2(C.vpaddq_f64(C.float64x2_t(v0), C.float64x2_t(v1)))
}

func VpaddqF32(v0 Float32X4, v1 Float32X4) Float32X4 {
	return Float32X4(C.vpaddq_f32(C.float32x4_t(v0), C.float32x4_t(v1)))
}

func VpaddqS32(v0 Int32X4, v1 Int32X4) Int32X4 {
	return Int32X4(C.vpaddq_s32(C.int32x4_t(v0), C.int32x4_t(v1)))
}

func VpaddqS64(v0 Int64X2, v1 Int64X2) Int64X2 {
	return Int64X2(C.vpaddq_s64(C.int64x2_t(v0), C.int64x2_t(v1)))
}

func VpaddqS16(v0 Int16X8, v1 Int16X8) Int16X8 {
	return Int16X8(C.vpaddq_s16(C.int16x8_t(v0), C.int16x8_t(v1)))
}

func VpadddU64(v0 Uint64X2) uint64 {
	return uint64(C.vpaddd_u64(C.uint64x2_t(v0)))
}

func VpadddF64(v0 Float64X2) float64 {
	return float64(C.vpaddd_f64(C.float64x2_t(v0)))
}

func VpadddS64(v0 Int64X2) int64 {
	return int64(C.vpaddd_s64(C.int64x2_t(v0)))
}

func VpaddsF32(v0 Float32X2) float32 {
	return float32(C.vpadds_f32(C.float32x2_t(v0)))
}

func VpmaxqU8(v0 Uint8X16, v1 Uint8X16) Uint8X16 {
	return Uint8X16(C.vpmaxq_u8(C.uint8x16_t(v0), C.uint8x16_t(v1)))
}

func VpmaxqU32(v0 Uint32X4, v1 Uint32X4) Uint32X4 {
	return Uint32X4(C.vpmaxq_u32(C.uint32x4_t(v0), C.uint32x4_t(v1)))
}

func VpmaxqU16(v0 Uint16X8, v1 Uint16X8) Uint16X8 {
	return Uint16X8(C.vpmaxq_u16(C.uint16x8_t(v0), C.uint16x8_t(v1)))
}

func VpmaxqS8(v0 Int8X16, v1 Int8X16) Int8X16 {
	return Int8X16(C.vpmaxq_s8(C.int8x16_t(v0), C.int8x16_t(v1)))
}

func VpmaxqF64(v0 Float64X2, v1 Float64X2) Float64X2 {
	return Float64X2(C.vpmaxq_f64(C.float64x2_t(v0), C.float64x2_t(v1)))
}

func VpmaxqF32(v0 Float32X4, v1 Float32X4) Float32X4 {
	return Float32X4(C.vpmaxq_f32(C.float32x4_t(v0), C.float32x4_t(v1)))
}

func VpmaxqS32(v0 Int32X4, v1 Int32X4) Int32X4 {
	return Int32X4(C.vpmaxq_s32(C.int32x4_t(v0), C.int32x4_t(v1)))
}

func VpmaxqS16(v0 Int16X8, v1 Int16X8) Int16X8 {
	return Int16X8(C.vpmaxq_s16(C.int16x8_t(v0), C.int16x8_t(v1)))
}

func VpmaxqdF64(v0 Float64X2) float64 {
	return float64(C.vpmaxqd_f64(C.float64x2_t(v0)))
}

func VpmaxsF32(v0 Float32X2) float32 {
	return float32(C.vpmaxs_f32(C.float32x2_t(v0)))
}

func VpmaxnmqF64(v0 Float64X2, v1 Float64X2) Float64X2 {
	return Float64X2(C.vpmaxnmq_f64(C.float64x2_t(v0), C.float64x2_t(v1)))
}

func VpmaxnmqF32(v0 Float32X4, v1 Float32X4) Float32X4 {
	return Float32X4(C.vpmaxnmq_f32(C.float32x4_t(v0), C.float32x4_t(v1)))
}

func VpmaxnmF32(v0 Float32X2, v1 Float32X2) Float32X2 {
	return Float32X2(C.vpmaxnm_f32(C.float32x2_t(v0), C.float32x2_t(v1)))
}

func VpmaxnmqdF64(v0 Float64X2) float64 {
	return float64(C.vpmaxnmqd_f64(C.float64x2_t(v0)))
}

func VpmaxnmsF32(v0 Float32X2) float32 {
	return float32(C.vpmaxnms_f32(C.float32x2_t(v0)))
}

func VpminqU8(v0 Uint8X16, v1 Uint8X16) Uint8X16 {
	return Uint8X16(C.vpminq_u8(C.uint8x16_t(v0), C.uint8x16_t(v1)))
}

func VpminqU32(v0 Uint32X4, v1 Uint32X4) Uint32X4 {
	return Uint32X4(C.vpminq_u32(C.uint32x4_t(v0), C.uint32x4_t(v1)))
}

func VpminqU16(v0 Uint16X8, v1 Uint16X8) Uint16X8 {
	return Uint16X8(C.vpminq_u16(C.uint16x8_t(v0), C.uint16x8_t(v1)))
}

func VpminqS8(v0 Int8X16, v1 Int8X16) Int8X16 {
	return Int8X16(C.vpminq_s8(C.int8x16_t(v0), C.int8x16_t(v1)))
}

func VpminqF64(v0 Float64X2, v1 Float64X2) Float64X2 {
	return Float64X2(C.vpminq_f64(C.float64x2_t(v0), C.float64x2_t(v1)))
}

func VpminqF32(v0 Float32X4, v1 Float32X4) Float32X4 {
	return Float32X4(C.vpminq_f32(C.float32x4_t(v0), C.float32x4_t(v1)))
}

func VpminqS32(v0 Int32X4, v1 Int32X4) Int32X4 {
	return Int32X4(C.vpminq_s32(C.int32x4_t(v0), C.int32x4_t(v1)))
}

func VpminqS16(v0 Int16X8, v1 Int16X8) Int16X8 {
	return Int16X8(C.vpminq_s16(C.int16x8_t(v0), C.int16x8_t(v1)))
}

func VpminqdF64(v0 Float64X2) float64 {
	return float64(C.vpminqd_f64(C.float64x2_t(v0)))
}

func VpminsF32(v0 Float32X2) float32 {
	return float32(C.vpmins_f32(C.float32x2_t(v0)))
}

func VpminnmqF64(v0 Float64X2, v1 Float64X2) Float64X2 {
	return Float64X2(C.vpminnmq_f64(C.float64x2_t(v0), C.float64x2_t(v1)))
}

func VpminnmqF32(v0 Float32X4, v1 Float32X4) Float32X4 {
	return Float32X4(C.vpminnmq_f32(C.float32x4_t(v0), C.float32x4_t(v1)))
}

func VpminnmF32(v0 Float32X2, v1 Float32X2) Float32X2 {
	return Float32X2(C.vpminnm_f32(C.float32x2_t(v0), C.float32x2_t(v1)))
}

func VpminnmqdF64(v0 Float64X2) float64 {
	return float64(C.vpminnmqd_f64(C.float64x2_t(v0)))
}

func VpminnmsF32(v0 Float32X2) float32 {
	return float32(C.vpminnms_f32(C.float32x2_t(v0)))
}

func VqabsqS64(v0 Int64X2) Int64X2 {
	return Int64X2(C.vqabsq_s64(C.int64x2_t(v0)))
}

func VqabsS64(v0 Int64X1) Int64X1 {
	return Int64X1(C.vqabs_s64(C.int64x1_t(v0)))
}

func VqabsbS8(v0 int8) int8 {
	return int8(C.vqabsb_s8(C.int8_t(v0)))
}

func VqabssS32(v0 int32) int32 {
	return int32(C.vqabss_s32(C.int32_t(v0)))
}

func VqabsdS64(v0 int64) int64 {
	return int64(C.vqabsd_s64(C.int64_t(v0)))
}

func VqabshS16(v0 int16) int16 {
	return int16(C.vqabsh_s16(C.int16_t(v0)))
}

func VqaddbU8(v0 uint8, v1 uint8) uint8 {
	return uint8(C.vqaddb_u8(C.uint8_t(v0), C.uint8_t(v1)))
}

func VqaddsU32(v0 uint32, v1 uint32) uint32 {
	return uint32(C.vqadds_u32(C.uint32_t(v0), C.uint32_t(v1)))
}

func VqadddU64(v0 uint64, v1 uint64) uint64 {
	return uint64(C.vqaddd_u64(C.uint64_t(v0), C.uint64_t(v1)))
}

func VqaddhU16(v0 uint16, v1 uint16) uint16 {
	return uint16(C.vqaddh_u16(C.uint16_t(v0), C.uint16_t(v1)))
}

func VqaddbS8(v0 int8, v1 int8) int8 {
	return int8(C.vqaddb_s8(C.int8_t(v0), C.int8_t(v1)))
}

func VqaddsS32(v0 int32, v1 int32) int32 {
	return int32(C.vqadds_s32(C.int32_t(v0), C.int32_t(v1)))
}

func VqadddS64(v0 int64, v1 int64) int64 {
	return int64(C.vqaddd_s64(C.int64_t(v0), C.int64_t(v1)))
}

func VqaddhS16(v0 int16, v1 int16) int16 {
	return int16(C.vqaddh_s16(C.int16_t(v0), C.int16_t(v1)))
}

func VqdmlalsS32(v0 int64, v1 int32, v2 int32) int64 {
	return int64(C.vqdmlals_s32(C.int64_t(v0), C.int32_t(v1), C.int32_t(v2)))
}

func VqdmlalhS16(v0 int32, v1 int16, v2 int16) int32 {
	return int32(C.vqdmlalh_s16(C.int32_t(v0), C.int16_t(v1), C.int16_t(v2)))
}

func VqdmlalHighS32(v0 Int64X2, v1 Int32X4, v2 Int32X4) Int64X2 {
	return Int64X2(C.vqdmlal_high_s32(C.int64x2_t(v0), C.int32x4_t(v1), C.int32x4_t(v2)))
}

func VqdmlalHighS16(v0 Int32X4, v1 Int16X8, v2 Int16X8) Int32X4 {
	return Int32X4(C.vqdmlal_high_s16(C.int32x4_t(v0), C.int16x8_t(v1), C.int16x8_t(v2)))
}

func VqdmlalHighNS32(v0 Int64X2, v1 Int32X4, v2 int32) Int64X2 {
	return Int64X2(C.vqdmlal_high_n_s32(C.int64x2_t(v0), C.int32x4_t(v1), C.int32_t(v2)))
}

func VqdmlalHighNS16(v0 Int32X4, v1 Int16X8, v2 int16) Int32X4 {
	return Int32X4(C.vqdmlal_high_n_s16(C.int32x4_t(v0), C.int16x8_t(v1), C.int16_t(v2)))
}

func VqdmlslsS32(v0 int64, v1 int32, v2 int32) int64 {
	return int64(C.vqdmlsls_s32(C.int64_t(v0), C.int32_t(v1), C.int32_t(v2)))
}

func VqdmlslhS16(v0 int32, v1 int16, v2 int16) int32 {
	return int32(C.vqdmlslh_s16(C.int32_t(v0), C.int16_t(v1), C.int16_t(v2)))
}

func VqdmlslHighS32(v0 Int64X2, v1 Int32X4, v2 Int32X4) Int64X2 {
	return Int64X2(C.vqdmlsl_high_s32(C.int64x2_t(v0), C.int32x4_t(v1), C.int32x4_t(v2)))
}

func VqdmlslHighS16(v0 Int32X4, v1 Int16X8, v2 Int16X8) Int32X4 {
	return Int32X4(C.vqdmlsl_high_s16(C.int32x4_t(v0), C.int16x8_t(v1), C.int16x8_t(v2)))
}

func VqdmlslHighNS32(v0 Int64X2, v1 Int32X4, v2 int32) Int64X2 {
	return Int64X2(C.vqdmlsl_high_n_s32(C.int64x2_t(v0), C.int32x4_t(v1), C.int32_t(v2)))
}

func VqdmlslHighNS16(v0 Int32X4, v1 Int16X8, v2 int16) Int32X4 {
	return Int32X4(C.vqdmlsl_high_n_s16(C.int32x4_t(v0), C.int16x8_t(v1), C.int16_t(v2)))
}

func VqdmulhsS32(v0 int32, v1 int32) int32 {
	return int32(C.vqdmulhs_s32(C.int32_t(v0), C.int32_t(v1)))
}

func VqdmulhhS16(v0 int16, v1 int16) int16 {
	return int16(C.vqdmulhh_s16(C.int16_t(v0), C.int16_t(v1)))
}

func VqdmullsS32(v0 int32, v1 int32) int64 {
	return int64(C.vqdmulls_s32(C.int32_t(v0), C.int32_t(v1)))
}

func VqdmullhS16(v0 int16, v1 int16) int32 {
	return int32(C.vqdmullh_s16(C.int16_t(v0), C.int16_t(v1)))
}

func VqdmullHighS32(v0 Int32X4, v1 Int32X4) Int64X2 {
	return Int64X2(C.vqdmull_high_s32(C.int32x4_t(v0), C.int32x4_t(v1)))
}

func VqdmullHighS16(v0 Int16X8, v1 Int16X8) Int32X4 {
	return Int32X4(C.vqdmull_high_s16(C.int16x8_t(v0), C.int16x8_t(v1)))
}

func VqdmullHighNS32(v0 Int32X4, v1 int32) Int64X2 {
	return Int64X2(C.vqdmull_high_n_s32(C.int32x4_t(v0), C.int32_t(v1)))
}

func VqdmullHighNS16(v0 Int16X8, v1 int16) Int32X4 {
	return Int32X4(C.vqdmull_high_n_s16(C.int16x8_t(v0), C.int16_t(v1)))
}

func VqmovnsS32(v0 int32) int16 {
	return int16(C.vqmovns_s32(C.int32_t(v0)))
}

func VqmovndS64(v0 int64) int32 {
	return int32(C.vqmovnd_s64(C.int64_t(v0)))
}

func VqmovnhS16(v0 int16) int8 {
	return int8(C.vqmovnh_s16(C.int16_t(v0)))
}

func VqmovnsU32(v0 uint32) uint16 {
	return uint16(C.vqmovns_u32(C.uint32_t(v0)))
}

func VqmovndU64(v0 uint64) uint32 {
	return uint32(C.vqmovnd_u64(C.uint64_t(v0)))
}

func VqmovnhU16(v0 uint16) uint8 {
	return uint8(C.vqmovnh_u16(C.uint16_t(v0)))
}

func VqmovnHighU32(v0 Uint16X4, v1 Uint32X4) Uint16X8 {
	return Uint16X8(C.vqmovn_high_u32(C.uint16x4_t(v0), C.uint32x4_t(v1)))
}

func VqmovnHighU64(v0 Uint32X2, v1 Uint64X2) Uint32X4 {
	return Uint32X4(C.vqmovn_high_u64(C.uint32x2_t(v0), C.uint64x2_t(v1)))
}

func VqmovnHighU16(v0 Uint8X8, v1 Uint16X8) Uint8X16 {
	return Uint8X16(C.vqmovn_high_u16(C.uint8x8_t(v0), C.uint16x8_t(v1)))
}

func VqmovnHighS32(v0 Int16X4, v1 Int32X4) Int16X8 {
	return Int16X8(C.vqmovn_high_s32(C.int16x4_t(v0), C.int32x4_t(v1)))
}

func VqmovnHighS64(v0 Int32X2, v1 Int64X2) Int32X4 {
	return Int32X4(C.vqmovn_high_s64(C.int32x2_t(v0), C.int64x2_t(v1)))
}

func VqmovnHighS16(v0 Int8X8, v1 Int16X8) Int8X16 {
	return Int8X16(C.vqmovn_high_s16(C.int8x8_t(v0), C.int16x8_t(v1)))
}

func VqmovunsS32(v0 int32) uint16 {
	return uint16(C.vqmovuns_s32(C.int32_t(v0)))
}

func VqmovundS64(v0 int64) uint32 {
	return uint32(C.vqmovund_s64(C.int64_t(v0)))
}

func VqmovunhS16(v0 int16) uint8 {
	return uint8(C.vqmovunh_s16(C.int16_t(v0)))
}

func VqmovunHighS32(v0 Uint16X4, v1 Int32X4) Uint16X8 {
	return Uint16X8(C.vqmovun_high_s32(C.uint16x4_t(v0), C.int32x4_t(v1)))
}

func VqmovunHighS64(v0 Uint32X2, v1 Int64X2) Uint32X4 {
	return Uint32X4(C.vqmovun_high_s64(C.uint32x2_t(v0), C.int64x2_t(v1)))
}

func VqmovunHighS16(v0 Uint8X8, v1 Int16X8) Uint8X16 {
	return Uint8X16(C.vqmovun_high_s16(C.uint8x8_t(v0), C.int16x8_t(v1)))
}

func VqnegqS64(v0 Int64X2) Int64X2 {
	return Int64X2(C.vqnegq_s64(C.int64x2_t(v0)))
}

func VqnegS64(v0 Int64X1) Int64X1 {
	return Int64X1(C.vqneg_s64(C.int64x1_t(v0)))
}

func VqnegbS8(v0 int8) int8 {
	return int8(C.vqnegb_s8(C.int8_t(v0)))
}

func VqnegsS32(v0 int32) int32 {
	return int32(C.vqnegs_s32(C.int32_t(v0)))
}

func VqnegdS64(v0 int64) int64 {
	return int64(C.vqnegd_s64(C.int64_t(v0)))
}

func VqneghS16(v0 int16) int16 {
	return int16(C.vqnegh_s16(C.int16_t(v0)))
}

func VqrdmulhsS32(v0 int32, v1 int32) int32 {
	return int32(C.vqrdmulhs_s32(C.int32_t(v0), C.int32_t(v1)))
}

func VqrdmulhhS16(v0 int16, v1 int16) int16 {
	return int16(C.vqrdmulhh_s16(C.int16_t(v0), C.int16_t(v1)))
}

func VqrshlbU8(v0 uint8, v1 int8) uint8 {
	return uint8(C.vqrshlb_u8(C.uint8_t(v0), C.int8_t(v1)))
}

func VqrshlsU32(v0 uint32, v1 int32) uint32 {
	return uint32(C.vqrshls_u32(C.uint32_t(v0), C.int32_t(v1)))
}

func VqrshldU64(v0 uint64, v1 int64) uint64 {
	return uint64(C.vqrshld_u64(C.uint64_t(v0), C.int64_t(v1)))
}

func VqrshlhU16(v0 uint16, v1 int16) uint16 {
	return uint16(C.vqrshlh_u16(C.uint16_t(v0), C.int16_t(v1)))
}

func VqrshlbS8(v0 int8, v1 int8) int8 {
	return int8(C.vqrshlb_s8(C.int8_t(v0), C.int8_t(v1)))
}

func VqrshlsS32(v0 int32, v1 int32) int32 {
	return int32(C.vqrshls_s32(C.int32_t(v0), C.int32_t(v1)))
}

func VqrshldS64(v0 int64, v1 int64) int64 {
	return int64(C.vqrshld_s64(C.int64_t(v0), C.int64_t(v1)))
}

func VqrshlhS16(v0 int16, v1 int16) int16 {
	return int16(C.vqrshlh_s16(C.int16_t(v0), C.int16_t(v1)))
}

func VqshlbU8(v0 uint8, v1 int8) uint8 {
	return uint8(C.vqshlb_u8(C.uint8_t(v0), C.int8_t(v1)))
}

func VqshlsU32(v0 uint32, v1 int32) uint32 {
	return uint32(C.vqshls_u32(C.uint32_t(v0), C.int32_t(v1)))
}

func VqshldU64(v0 uint64, v1 int64) uint64 {
	return uint64(C.vqshld_u64(C.uint64_t(v0), C.int64_t(v1)))
}

func VqshlhU16(v0 uint16, v1 int16) uint16 {
	return uint16(C.vqshlh_u16(C.uint16_t(v0), C.int16_t(v1)))
}

func VqshlbS8(v0 int8, v1 int8) int8 {
	return int8(C.vqshlb_s8(C.int8_t(v0), C.int8_t(v1)))
}

func VqshlsS32(v0 int32, v1 int32) int32 {
	return int32(C.vqshls_s32(C.int32_t(v0), C.int32_t(v1)))
}

func VqshldS64(v0 int64, v1 int64) int64 {
	return int64(C.vqshld_s64(C.int64_t(v0), C.int64_t(v1)))
}

func VqshlhS16(v0 int16, v1 int16) int16 {
	return int16(C.vqshlh_s16(C.int16_t(v0), C.int16_t(v1)))
}

func VqsubbU8(v0 uint8, v1 uint8) uint8 {
	return uint8(C.vqsubb_u8(C.uint8_t(v0), C.uint8_t(v1)))
}

func VqsubsU32(v0 uint32, v1 uint32) uint32 {
	return uint32(C.vqsubs_u32(C.uint32_t(v0), C.uint32_t(v1)))
}

func VqsubdU64(v0 uint64, v1 uint64) uint64 {
	return uint64(C.vqsubd_u64(C.uint64_t(v0), C.uint64_t(v1)))
}

func VqsubhU16(v0 uint16, v1 uint16) uint16 {
	return uint16(C.vqsubh_u16(C.uint16_t(v0), C.uint16_t(v1)))
}

func VqsubbS8(v0 int8, v1 int8) int8 {
	return int8(C.vqsubb_s8(C.int8_t(v0), C.int8_t(v1)))
}

func VqsubsS32(v0 int32, v1 int32) int32 {
	return int32(C.vqsubs_s32(C.int32_t(v0), C.int32_t(v1)))
}

func VqsubdS64(v0 int64, v1 int64) int64 {
	return int64(C.vqsubd_s64(C.int64_t(v0), C.int64_t(v1)))
}

func VqsubhS16(v0 int16, v1 int16) int16 {
	return int16(C.vqsubh_s16(C.int16_t(v0), C.int16_t(v1)))
}

func Vqtbl1P8(v0 Poly8X16, v1 Uint8X8) Poly8X8 {
	return Poly8X8(C.vqtbl1_p8(C.poly8x16_t(v0), C.uint8x8_t(v1)))
}

func Vqtbl1QP8(v0 Poly8X16, v1 Uint8X16) Poly8X16 {
	return Poly8X16(C.vqtbl1q_p8(C.poly8x16_t(v0), C.uint8x16_t(v1)))
}

func Vqtbl1QU8(v0 Uint8X16, v1 Uint8X16) Uint8X16 {
	return Uint8X16(C.vqtbl1q_u8(C.uint8x16_t(v0), C.uint8x16_t(v1)))
}

func Vqtbl1QS8(v0 Int8X16, v1 Uint8X16) Int8X16 {
	return Int8X16(C.vqtbl1q_s8(C.int8x16_t(v0), C.uint8x16_t(v1)))
}

func Vqtbl1U8(v0 Uint8X16, v1 Uint8X8) Uint8X8 {
	return Uint8X8(C.vqtbl1_u8(C.uint8x16_t(v0), C.uint8x8_t(v1)))
}

func Vqtbl1S8(v0 Int8X16, v1 Uint8X8) Int8X8 {
	return Int8X8(C.vqtbl1_s8(C.int8x16_t(v0), C.uint8x8_t(v1)))
}

func Vqtbl2P8(v0 Poly8X16X2, v1 Uint8X8) Poly8X8 {
	return Poly8X8(C.vqtbl2_p8(C.poly8x16x2_t(v0), C.uint8x8_t(v1)))
}

func Vqtbl2QP8(v0 Poly8X16X2, v1 Uint8X16) Poly8X16 {
	return Poly8X16(C.vqtbl2q_p8(C.poly8x16x2_t(v0), C.uint8x16_t(v1)))
}

func Vqtbl2QU8(v0 Uint8X16X2, v1 Uint8X16) Uint8X16 {
	return Uint8X16(C.vqtbl2q_u8(C.uint8x16x2_t(v0), C.uint8x16_t(v1)))
}

func Vqtbl2QS8(v0 Int8X16X2, v1 Uint8X16) Int8X16 {
	return Int8X16(C.vqtbl2q_s8(C.int8x16x2_t(v0), C.uint8x16_t(v1)))
}

func Vqtbl2U8(v0 Uint8X16X2, v1 Uint8X8) Uint8X8 {
	return Uint8X8(C.vqtbl2_u8(C.uint8x16x2_t(v0), C.uint8x8_t(v1)))
}

func Vqtbl2S8(v0 Int8X16X2, v1 Uint8X8) Int8X8 {
	return Int8X8(C.vqtbl2_s8(C.int8x16x2_t(v0), C.uint8x8_t(v1)))
}

func Vqtbl3P8(v0 Poly8X16X3, v1 Uint8X8) Poly8X8 {
	return Poly8X8(C.vqtbl3_p8(C.poly8x16x3_t(v0), C.uint8x8_t(v1)))
}

func Vqtbl3QP8(v0 Poly8X16X3, v1 Uint8X16) Poly8X16 {
	return Poly8X16(C.vqtbl3q_p8(C.poly8x16x3_t(v0), C.uint8x16_t(v1)))
}

func Vqtbl3QU8(v0 Uint8X16X3, v1 Uint8X16) Uint8X16 {
	return Uint8X16(C.vqtbl3q_u8(C.uint8x16x3_t(v0), C.uint8x16_t(v1)))
}

func Vqtbl3QS8(v0 Int8X16X3, v1 Uint8X16) Int8X16 {
	return Int8X16(C.vqtbl3q_s8(C.int8x16x3_t(v0), C.uint8x16_t(v1)))
}

func Vqtbl3U8(v0 Uint8X16X3, v1 Uint8X8) Uint8X8 {
	return Uint8X8(C.vqtbl3_u8(C.uint8x16x3_t(v0), C.uint8x8_t(v1)))
}

func Vqtbl3S8(v0 Int8X16X3, v1 Uint8X8) Int8X8 {
	return Int8X8(C.vqtbl3_s8(C.int8x16x3_t(v0), C.uint8x8_t(v1)))
}

func Vqtbl4P8(v0 Poly8X16X4, v1 Uint8X8) Poly8X8 {
	return Poly8X8(C.vqtbl4_p8(C.poly8x16x4_t(v0), C.uint8x8_t(v1)))
}

func Vqtbl4QP8(v0 Poly8X16X4, v1 Uint8X16) Poly8X16 {
	return Poly8X16(C.vqtbl4q_p8(C.poly8x16x4_t(v0), C.uint8x16_t(v1)))
}

func Vqtbl4QU8(v0 Uint8X16X4, v1 Uint8X16) Uint8X16 {
	return Uint8X16(C.vqtbl4q_u8(C.uint8x16x4_t(v0), C.uint8x16_t(v1)))
}

func Vqtbl4QS8(v0 Int8X16X4, v1 Uint8X16) Int8X16 {
	return Int8X16(C.vqtbl4q_s8(C.int8x16x4_t(v0), C.uint8x16_t(v1)))
}

func Vqtbl4U8(v0 Uint8X16X4, v1 Uint8X8) Uint8X8 {
	return Uint8X8(C.vqtbl4_u8(C.uint8x16x4_t(v0), C.uint8x8_t(v1)))
}

func Vqtbl4S8(v0 Int8X16X4, v1 Uint8X8) Int8X8 {
	return Int8X8(C.vqtbl4_s8(C.int8x16x4_t(v0), C.uint8x8_t(v1)))
}

func Vqtbx1P8(v0 Poly8X8, v1 Poly8X16, v2 Uint8X8) Poly8X8 {
	return Poly8X8(C.vqtbx1_p8(C.poly8x8_t(v0), C.poly8x16_t(v1), C.uint8x8_t(v2)))
}

func Vqtbx1QP8(v0 Poly8X16, v1 Poly8X16, v2 Uint8X16) Poly8X16 {
	return Poly8X16(C.vqtbx1q_p8(C.poly8x16_t(v0), C.poly8x16_t(v1), C.uint8x16_t(v2)))
}

func Vqtbx1QU8(v0 Uint8X16, v1 Uint8X16, v2 Uint8X16) Uint8X16 {
	return Uint8X16(C.vqtbx1q_u8(C.uint8x16_t(v0), C.uint8x16_t(v1), C.uint8x16_t(v2)))
}

func Vqtbx1QS8(v0 Int8X16, v1 Int8X16, v2 Uint8X16) Int8X16 {
	return Int8X16(C.vqtbx1q_s8(C.int8x16_t(v0), C.int8x16_t(v1), C.uint8x16_t(v2)))
}

func Vqtbx1U8(v0 Uint8X8, v1 Uint8X16, v2 Uint8X8) Uint8X8 {
	return Uint8X8(C.vqtbx1_u8(C.uint8x8_t(v0), C.uint8x16_t(v1), C.uint8x8_t(v2)))
}

func Vqtbx1S8(v0 Int8X8, v1 Int8X16, v2 Uint8X8) Int8X8 {
	return Int8X8(C.vqtbx1_s8(C.int8x8_t(v0), C.int8x16_t(v1), C.uint8x8_t(v2)))
}

func Vqtbx2P8(v0 Poly8X8, v1 Poly8X16X2, v2 Uint8X8) Poly8X8 {
	return Poly8X8(C.vqtbx2_p8(C.poly8x8_t(v0), C.poly8x16x2_t(v1), C.uint8x8_t(v2)))
}

func Vqtbx2QP8(v0 Poly8X16, v1 Poly8X16X2, v2 Uint8X16) Poly8X16 {
	return Poly8X16(C.vqtbx2q_p8(C.poly8x16_t(v0), C.poly8x16x2_t(v1), C.uint8x16_t(v2)))
}

func Vqtbx2QU8(v0 Uint8X16, v1 Uint8X16X2, v2 Uint8X16) Uint8X16 {
	return Uint8X16(C.vqtbx2q_u8(C.uint8x16_t(v0), C.uint8x16x2_t(v1), C.uint8x16_t(v2)))
}

func Vqtbx2QS8(v0 Int8X16, v1 Int8X16X2, v2 Uint8X16) Int8X16 {
	return Int8X16(C.vqtbx2q_s8(C.int8x16_t(v0), C.int8x16x2_t(v1), C.uint8x16_t(v2)))
}

func Vqtbx2U8(v0 Uint8X8, v1 Uint8X16X2, v2 Uint8X8) Uint8X8 {
	return Uint8X8(C.vqtbx2_u8(C.uint8x8_t(v0), C.uint8x16x2_t(v1), C.uint8x8_t(v2)))
}

func Vqtbx2S8(v0 Int8X8, v1 Int8X16X2, v2 Uint8X8) Int8X8 {
	return Int8X8(C.vqtbx2_s8(C.int8x8_t(v0), C.int8x16x2_t(v1), C.uint8x8_t(v2)))
}

func Vqtbx3P8(v0 Poly8X8, v1 Poly8X16X3, v2 Uint8X8) Poly8X8 {
	return Poly8X8(C.vqtbx3_p8(C.poly8x8_t(v0), C.poly8x16x3_t(v1), C.uint8x8_t(v2)))
}

func Vqtbx3QP8(v0 Poly8X16, v1 Poly8X16X3, v2 Uint8X16) Poly8X16 {
	return Poly8X16(C.vqtbx3q_p8(C.poly8x16_t(v0), C.poly8x16x3_t(v1), C.uint8x16_t(v2)))
}

func Vqtbx3QU8(v0 Uint8X16, v1 Uint8X16X3, v2 Uint8X16) Uint8X16 {
	return Uint8X16(C.vqtbx3q_u8(C.uint8x16_t(v0), C.uint8x16x3_t(v1), C.uint8x16_t(v2)))
}

func Vqtbx3QS8(v0 Int8X16, v1 Int8X16X3, v2 Uint8X16) Int8X16 {
	return Int8X16(C.vqtbx3q_s8(C.int8x16_t(v0), C.int8x16x3_t(v1), C.uint8x16_t(v2)))
}

func Vqtbx3U8(v0 Uint8X8, v1 Uint8X16X3, v2 Uint8X8) Uint8X8 {
	return Uint8X8(C.vqtbx3_u8(C.uint8x8_t(v0), C.uint8x16x3_t(v1), C.uint8x8_t(v2)))
}

func Vqtbx3S8(v0 Int8X8, v1 Int8X16X3, v2 Uint8X8) Int8X8 {
	return Int8X8(C.vqtbx3_s8(C.int8x8_t(v0), C.int8x16x3_t(v1), C.uint8x8_t(v2)))
}

func Vqtbx4P8(v0 Poly8X8, v1 Poly8X16X4, v2 Uint8X8) Poly8X8 {
	return Poly8X8(C.vqtbx4_p8(C.poly8x8_t(v0), C.poly8x16x4_t(v1), C.uint8x8_t(v2)))
}

func Vqtbx4QP8(v0 Poly8X16, v1 Poly8X16X4, v2 Uint8X16) Poly8X16 {
	return Poly8X16(C.vqtbx4q_p8(C.poly8x16_t(v0), C.poly8x16x4_t(v1), C.uint8x16_t(v2)))
}

func Vqtbx4QU8(v0 Uint8X16, v1 Uint8X16X4, v2 Uint8X16) Uint8X16 {
	return Uint8X16(C.vqtbx4q_u8(C.uint8x16_t(v0), C.uint8x16x4_t(v1), C.uint8x16_t(v2)))
}

func Vqtbx4QS8(v0 Int8X16, v1 Int8X16X4, v2 Uint8X16) Int8X16 {
	return Int8X16(C.vqtbx4q_s8(C.int8x16_t(v0), C.int8x16x4_t(v1), C.uint8x16_t(v2)))
}

func Vqtbx4U8(v0 Uint8X8, v1 Uint8X16X4, v2 Uint8X8) Uint8X8 {
	return Uint8X8(C.vqtbx4_u8(C.uint8x8_t(v0), C.uint8x16x4_t(v1), C.uint8x8_t(v2)))
}

func Vqtbx4S8(v0 Int8X8, v1 Int8X16X4, v2 Uint8X8) Int8X8 {
	return Int8X8(C.vqtbx4_s8(C.int8x8_t(v0), C.int8x16x4_t(v1), C.uint8x8_t(v2)))
}

func VraddhnHighU32(v0 Uint16X4, v1 Uint32X4, v2 Uint32X4) Uint16X8 {
	return Uint16X8(C.vraddhn_high_u32(C.uint16x4_t(v0), C.uint32x4_t(v1), C.uint32x4_t(v2)))
}

func VraddhnHighU64(v0 Uint32X2, v1 Uint64X2, v2 Uint64X2) Uint32X4 {
	return Uint32X4(C.vraddhn_high_u64(C.uint32x2_t(v0), C.uint64x2_t(v1), C.uint64x2_t(v2)))
}

func VraddhnHighU16(v0 Uint8X8, v1 Uint16X8, v2 Uint16X8) Uint8X16 {
	return Uint8X16(C.vraddhn_high_u16(C.uint8x8_t(v0), C.uint16x8_t(v1), C.uint16x8_t(v2)))
}

func VraddhnHighS32(v0 Int16X4, v1 Int32X4, v2 Int32X4) Int16X8 {
	return Int16X8(C.vraddhn_high_s32(C.int16x4_t(v0), C.int32x4_t(v1), C.int32x4_t(v2)))
}

func VraddhnHighS64(v0 Int32X2, v1 Int64X2, v2 Int64X2) Int32X4 {
	return Int32X4(C.vraddhn_high_s64(C.int32x2_t(v0), C.int64x2_t(v1), C.int64x2_t(v2)))
}

func VraddhnHighS16(v0 Int8X8, v1 Int16X8, v2 Int16X8) Int8X16 {
	return Int8X16(C.vraddhn_high_s16(C.int8x8_t(v0), C.int16x8_t(v1), C.int16x8_t(v2)))
}

func VrbitP8(v0 Poly8X8) Poly8X8 {
	return Poly8X8(C.vrbit_p8(C.poly8x8_t(v0)))
}

func VrbitqP8(v0 Poly8X16) Poly8X16 {
	return Poly8X16(C.vrbitq_p8(C.poly8x16_t(v0)))
}

func VrbitqU8(v0 Uint8X16) Uint8X16 {
	return Uint8X16(C.vrbitq_u8(C.uint8x16_t(v0)))
}

func VrbitqS8(v0 Int8X16) Int8X16 {
	return Int8X16(C.vrbitq_s8(C.int8x16_t(v0)))
}

func VrbitU8(v0 Uint8X8) Uint8X8 {
	return Uint8X8(C.vrbit_u8(C.uint8x8_t(v0)))
}

func VrbitS8(v0 Int8X8) Int8X8 {
	return Int8X8(C.vrbit_s8(C.int8x8_t(v0)))
}

func VrecpeqF64(v0 Float64X2) Float64X2 {
	return Float64X2(C.vrecpeq_f64(C.float64x2_t(v0)))
}

func VrecpeF64(v0 Float64X1) Float64X1 {
	return Float64X1(C.vrecpe_f64(C.float64x1_t(v0)))
}

func VrecpedF64(v0 float64) float64 {
	return float64(C.vrecped_f64(C.float64_t(v0)))
}

func VrecpesF32(v0 float32) float32 {
	return float32(C.vrecpes_f32(C.float32_t(v0)))
}

func VrecpsqF64(v0 Float64X2, v1 Float64X2) Float64X2 {
	return Float64X2(C.vrecpsq_f64(C.float64x2_t(v0), C.float64x2_t(v1)))
}

func VrecpsF64(v0 Float64X1, v1 Float64X1) Float64X1 {
	return Float64X1(C.vrecps_f64(C.float64x1_t(v0), C.float64x1_t(v1)))
}

func VrecpsdF64(v0 float64, v1 float64) float64 {
	return float64(C.vrecpsd_f64(C.float64_t(v0), C.float64_t(v1)))
}

func VrecpssF32(v0 float32, v1 float32) float32 {
	return float32(C.vrecpss_f32(C.float32_t(v0), C.float32_t(v1)))
}

func VrecpxdF64(v0 float64) float64 {
	return float64(C.vrecpxd_f64(C.float64_t(v0)))
}

func VrecpxsF32(v0 float32) float32 {
	return float32(C.vrecpxs_f32(C.float32_t(v0)))
}

func VrshldU64(v0 uint64, v1 int64) uint64 {
	return uint64(C.vrshld_u64(C.uint64_t(v0), C.int64_t(v1)))
}

func VrshldS64(v0 int64, v1 int64) int64 {
	return int64(C.vrshld_s64(C.int64_t(v0), C.int64_t(v1)))
}

func VrsqrteqF64(v0 Float64X2) Float64X2 {
	return Float64X2(C.vrsqrteq_f64(C.float64x2_t(v0)))
}

func VrsqrteF64(v0 Float64X1) Float64X1 {
	return Float64X1(C.vrsqrte_f64(C.float64x1_t(v0)))
}

func VrsqrtedF64(v0 float64) float64 {
	return float64(C.vrsqrted_f64(C.float64_t(v0)))
}

func VrsqrtesF32(v0 float32) float32 {
	return float32(C.vrsqrtes_f32(C.float32_t(v0)))
}

func VrsqrtsqF64(v0 Float64X2, v1 Float64X2) Float64X2 {
	return Float64X2(C.vrsqrtsq_f64(C.float64x2_t(v0), C.float64x2_t(v1)))
}

func VrsqrtsF64(v0 Float64X1, v1 Float64X1) Float64X1 {
	return Float64X1(C.vrsqrts_f64(C.float64x1_t(v0), C.float64x1_t(v1)))
}

func VrsqrtsdF64(v0 float64, v1 float64) float64 {
	return float64(C.vrsqrtsd_f64(C.float64_t(v0), C.float64_t(v1)))
}

func VrsqrtssF32(v0 float32, v1 float32) float32 {
	return float32(C.vrsqrtss_f32(C.float32_t(v0), C.float32_t(v1)))
}

func VrsubhnHighU32(v0 Uint16X4, v1 Uint32X4, v2 Uint32X4) Uint16X8 {
	return Uint16X8(C.vrsubhn_high_u32(C.uint16x4_t(v0), C.uint32x4_t(v1), C.uint32x4_t(v2)))
}

func VrsubhnHighU64(v0 Uint32X2, v1 Uint64X2, v2 Uint64X2) Uint32X4 {
	return Uint32X4(C.vrsubhn_high_u64(C.uint32x2_t(v0), C.uint64x2_t(v1), C.uint64x2_t(v2)))
}

func VrsubhnHighU16(v0 Uint8X8, v1 Uint16X8, v2 Uint16X8) Uint8X16 {
	return Uint8X16(C.vrsubhn_high_u16(C.uint8x8_t(v0), C.uint16x8_t(v1), C.uint16x8_t(v2)))
}

func VrsubhnHighS32(v0 Int16X4, v1 Int32X4, v2 Int32X4) Int16X8 {
	return Int16X8(C.vrsubhn_high_s32(C.int16x4_t(v0), C.int32x4_t(v1), C.int32x4_t(v2)))
}

func VrsubhnHighS64(v0 Int32X2, v1 Int64X2, v2 Int64X2) Int32X4 {
	return Int32X4(C.vrsubhn_high_s64(C.int32x2_t(v0), C.int64x2_t(v1), C.int64x2_t(v2)))
}

func VrsubhnHighS16(v0 Int8X8, v1 Int16X8, v2 Int16X8) Int8X16 {
	return Int8X16(C.vrsubhn_high_s16(C.int8x8_t(v0), C.int16x8_t(v1), C.int16x8_t(v2)))
}

func VshldU64(v0 uint64, v1 int64) uint64 {
	return uint64(C.vshld_u64(C.uint64_t(v0), C.int64_t(v1)))
}

func VshldS64(v0 int64, v1 int64) int64 {
	return int64(C.vshld_s64(C.int64_t(v0), C.int64_t(v1)))
}

func VsqaddbU8(v0 uint8, v1 int8) uint8 {
	return uint8(C.vsqaddb_u8(C.uint8_t(v0), C.int8_t(v1)))
}

func VsqaddsU32(v0 uint32, v1 int32) uint32 {
	return uint32(C.vsqadds_u32(C.uint32_t(v0), C.int32_t(v1)))
}

func VsqadddU64(v0 uint64, v1 int64) uint64 {
	return uint64(C.vsqaddd_u64(C.uint64_t(v0), C.int64_t(v1)))
}

func VsqaddhU16(v0 uint16, v1 int16) uint16 {
	return uint16(C.vsqaddh_u16(C.uint16_t(v0), C.int16_t(v1)))
}

func VsqaddqU8(v0 Uint8X16, v1 Int8X16) Uint8X16 {
	return Uint8X16(C.vsqaddq_u8(C.uint8x16_t(v0), C.int8x16_t(v1)))
}

func VsqaddqU32(v0 Uint32X4, v1 Int32X4) Uint32X4 {
	return Uint32X4(C.vsqaddq_u32(C.uint32x4_t(v0), C.int32x4_t(v1)))
}

func VsqaddqU64(v0 Uint64X2, v1 Int64X2) Uint64X2 {
	return Uint64X2(C.vsqaddq_u64(C.uint64x2_t(v0), C.int64x2_t(v1)))
}

func VsqaddqU16(v0 Uint16X8, v1 Int16X8) Uint16X8 {
	return Uint16X8(C.vsqaddq_u16(C.uint16x8_t(v0), C.int16x8_t(v1)))
}

func VsqaddU8(v0 Uint8X8, v1 Int8X8) Uint8X8 {
	return Uint8X8(C.vsqadd_u8(C.uint8x8_t(v0), C.int8x8_t(v1)))
}

func VsqaddU32(v0 Uint32X2, v1 Int32X2) Uint32X2 {
	return Uint32X2(C.vsqadd_u32(C.uint32x2_t(v0), C.int32x2_t(v1)))
}

func VsqaddU64(v0 Uint64X1, v1 Int64X1) Uint64X1 {
	return Uint64X1(C.vsqadd_u64(C.uint64x1_t(v0), C.int64x1_t(v1)))
}

func VsqaddU16(v0 Uint16X4, v1 Int16X4) Uint16X4 {
	return Uint16X4(C.vsqadd_u16(C.uint16x4_t(v0), C.int16x4_t(v1)))
}

func VsqrtqF64(v0 Float64X2) Float64X2 {
	return Float64X2(C.vsqrtq_f64(C.float64x2_t(v0)))
}

func VsqrtqF32(v0 Float32X4) Float32X4 {
	return Float32X4(C.vsqrtq_f32(C.float32x4_t(v0)))
}

func VsqrtF64(v0 Float64X1) Float64X1 {
	return Float64X1(C.vsqrt_f64(C.float64x1_t(v0)))
}

func VsqrtF32(v0 Float32X2) Float32X2 {
	return Float32X2(C.vsqrt_f32(C.float32x2_t(v0)))
}

func VsubdU64(v0 uint64, v1 uint64) uint64 {
	return uint64(C.vsubd_u64(C.uint64_t(v0), C.uint64_t(v1)))
}

func VsubdS64(v0 int64, v1 int64) int64 {
	return int64(C.vsubd_s64(C.int64_t(v0), C.int64_t(v1)))
}

func VsubqF64(v0 Float64X2, v1 Float64X2) Float64X2 {
	return Float64X2(C.vsubq_f64(C.float64x2_t(v0), C.float64x2_t(v1)))
}

func VsubF64(v0 Float64X1, v1 Float64X1) Float64X1 {
	return Float64X1(C.vsub_f64(C.float64x1_t(v0), C.float64x1_t(v1)))
}

func VsubhnHighU32(v0 Uint16X4, v1 Uint32X4, v2 Uint32X4) Uint16X8 {
	return Uint16X8(C.vsubhn_high_u32(C.uint16x4_t(v0), C.uint32x4_t(v1), C.uint32x4_t(v2)))
}

func VsubhnHighU64(v0 Uint32X2, v1 Uint64X2, v2 Uint64X2) Uint32X4 {
	return Uint32X4(C.vsubhn_high_u64(C.uint32x2_t(v0), C.uint64x2_t(v1), C.uint64x2_t(v2)))
}

func VsubhnHighU16(v0 Uint8X8, v1 Uint16X8, v2 Uint16X8) Uint8X16 {
	return Uint8X16(C.vsubhn_high_u16(C.uint8x8_t(v0), C.uint16x8_t(v1), C.uint16x8_t(v2)))
}

func VsubhnHighS32(v0 Int16X4, v1 Int32X4, v2 Int32X4) Int16X8 {
	return Int16X8(C.vsubhn_high_s32(C.int16x4_t(v0), C.int32x4_t(v1), C.int32x4_t(v2)))
}

func VsubhnHighS64(v0 Int32X2, v1 Int64X2, v2 Int64X2) Int32X4 {
	return Int32X4(C.vsubhn_high_s64(C.int32x2_t(v0), C.int64x2_t(v1), C.int64x2_t(v2)))
}

func VsubhnHighS16(v0 Int8X8, v1 Int16X8, v2 Int16X8) Int8X16 {
	return Int8X16(C.vsubhn_high_s16(C.int8x8_t(v0), C.int16x8_t(v1), C.int16x8_t(v2)))
}

func VsublHighU8(v0 Uint8X16, v1 Uint8X16) Uint16X8 {
	return Uint16X8(C.vsubl_high_u8(C.uint8x16_t(v0), C.uint8x16_t(v1)))
}

func VsublHighU32(v0 Uint32X4, v1 Uint32X4) Uint64X2 {
	return Uint64X2(C.vsubl_high_u32(C.uint32x4_t(v0), C.uint32x4_t(v1)))
}

func VsublHighU16(v0 Uint16X8, v1 Uint16X8) Uint32X4 {
	return Uint32X4(C.vsubl_high_u16(C.uint16x8_t(v0), C.uint16x8_t(v1)))
}

func VsublHighS8(v0 Int8X16, v1 Int8X16) Int16X8 {
	return Int16X8(C.vsubl_high_s8(C.int8x16_t(v0), C.int8x16_t(v1)))
}

func VsublHighS32(v0 Int32X4, v1 Int32X4) Int64X2 {
	return Int64X2(C.vsubl_high_s32(C.int32x4_t(v0), C.int32x4_t(v1)))
}

func VsublHighS16(v0 Int16X8, v1 Int16X8) Int32X4 {
	return Int32X4(C.vsubl_high_s16(C.int16x8_t(v0), C.int16x8_t(v1)))
}

func VsubwHighU8(v0 Uint16X8, v1 Uint8X16) Uint16X8 {
	return Uint16X8(C.vsubw_high_u8(C.uint16x8_t(v0), C.uint8x16_t(v1)))
}

func VsubwHighU32(v0 Uint64X2, v1 Uint32X4) Uint64X2 {
	return Uint64X2(C.vsubw_high_u32(C.uint64x2_t(v0), C.uint32x4_t(v1)))
}

func VsubwHighU16(v0 Uint32X4, v1 Uint16X8) Uint32X4 {
	return Uint32X4(C.vsubw_high_u16(C.uint32x4_t(v0), C.uint16x8_t(v1)))
}

func VsubwHighS8(v0 Int16X8, v1 Int8X16) Int16X8 {
	return Int16X8(C.vsubw_high_s8(C.int16x8_t(v0), C.int8x16_t(v1)))
}

func VsubwHighS32(v0 Int64X2, v1 Int32X4) Int64X2 {
	return Int64X2(C.vsubw_high_s32(C.int64x2_t(v0), C.int32x4_t(v1)))
}

func VsubwHighS16(v0 Int32X4, v1 Int16X8) Int32X4 {
	return Int32X4(C.vsubw_high_s16(C.int32x4_t(v0), C.int16x8_t(v1)))
}

func Vtrn1P8(v0 Poly8X8, v1 Poly8X8) Poly8X8 {
	return Poly8X8(C.vtrn1_p8(C.poly8x8_t(v0), C.poly8x8_t(v1)))
}

func Vtrn1P16(v0 Poly16X4, v1 Poly16X4) Poly16X4 {
	return Poly16X4(C.vtrn1_p16(C.poly16x4_t(v0), C.poly16x4_t(v1)))
}

func Vtrn1QP8(v0 Poly8X16, v1 Poly8X16) Poly8X16 {
	return Poly8X16(C.vtrn1q_p8(C.poly8x16_t(v0), C.poly8x16_t(v1)))
}

func Vtrn1QP64(v0 Poly64X2, v1 Poly64X2) Poly64X2 {
	return Poly64X2(C.vtrn1q_p64(C.poly64x2_t(v0), C.poly64x2_t(v1)))
}

func Vtrn1QP16(v0 Poly16X8, v1 Poly16X8) Poly16X8 {
	return Poly16X8(C.vtrn1q_p16(C.poly16x8_t(v0), C.poly16x8_t(v1)))
}

func Vtrn1QU8(v0 Uint8X16, v1 Uint8X16) Uint8X16 {
	return Uint8X16(C.vtrn1q_u8(C.uint8x16_t(v0), C.uint8x16_t(v1)))
}

func Vtrn1QU32(v0 Uint32X4, v1 Uint32X4) Uint32X4 {
	return Uint32X4(C.vtrn1q_u32(C.uint32x4_t(v0), C.uint32x4_t(v1)))
}

func Vtrn1QU64(v0 Uint64X2, v1 Uint64X2) Uint64X2 {
	return Uint64X2(C.vtrn1q_u64(C.uint64x2_t(v0), C.uint64x2_t(v1)))
}

func Vtrn1QU16(v0 Uint16X8, v1 Uint16X8) Uint16X8 {
	return Uint16X8(C.vtrn1q_u16(C.uint16x8_t(v0), C.uint16x8_t(v1)))
}

func Vtrn1QS8(v0 Int8X16, v1 Int8X16) Int8X16 {
	return Int8X16(C.vtrn1q_s8(C.int8x16_t(v0), C.int8x16_t(v1)))
}

func Vtrn1QF64(v0 Float64X2, v1 Float64X2) Float64X2 {
	return Float64X2(C.vtrn1q_f64(C.float64x2_t(v0), C.float64x2_t(v1)))
}

func Vtrn1QF32(v0 Float32X4, v1 Float32X4) Float32X4 {
	return Float32X4(C.vtrn1q_f32(C.float32x4_t(v0), C.float32x4_t(v1)))
}

func Vtrn1QS32(v0 Int32X4, v1 Int32X4) Int32X4 {
	return Int32X4(C.vtrn1q_s32(C.int32x4_t(v0), C.int32x4_t(v1)))
}

func Vtrn1QS64(v0 Int64X2, v1 Int64X2) Int64X2 {
	return Int64X2(C.vtrn1q_s64(C.int64x2_t(v0), C.int64x2_t(v1)))
}

func Vtrn1QS16(v0 Int16X8, v1 Int16X8) Int16X8 {
	return Int16X8(C.vtrn1q_s16(C.int16x8_t(v0), C.int16x8_t(v1)))
}

func Vtrn1U8(v0 Uint8X8, v1 Uint8X8) Uint8X8 {
	return Uint8X8(C.vtrn1_u8(C.uint8x8_t(v0), C.uint8x8_t(v1)))
}

func Vtrn1U32(v0 Uint32X2, v1 Uint32X2) Uint32X2 {
	return Uint32X2(C.vtrn1_u32(C.uint32x2_t(v0), C.uint32x2_t(v1)))
}

func Vtrn1U16(v0 Uint16X4, v1 Uint16X4) Uint16X4 {
	return Uint16X4(C.vtrn1_u16(C.uint16x4_t(v0), C.uint16x4_t(v1)))
}

func Vtrn1S8(v0 Int8X8, v1 Int8X8) Int8X8 {
	return Int8X8(C.vtrn1_s8(C.int8x8_t(v0), C.int8x8_t(v1)))
}

func Vtrn1F32(v0 Float32X2, v1 Float32X2) Float32X2 {
	return Float32X2(C.vtrn1_f32(C.float32x2_t(v0), C.float32x2_t(v1)))
}

func Vtrn1S32(v0 Int32X2, v1 Int32X2) Int32X2 {
	return Int32X2(C.vtrn1_s32(C.int32x2_t(v0), C.int32x2_t(v1)))
}

func Vtrn1S16(v0 Int16X4, v1 Int16X4) Int16X4 {
	return Int16X4(C.vtrn1_s16(C.int16x4_t(v0), C.int16x4_t(v1)))
}

func Vtrn2P8(v0 Poly8X8, v1 Poly8X8) Poly8X8 {
	return Poly8X8(C.vtrn2_p8(C.poly8x8_t(v0), C.poly8x8_t(v1)))
}

func Vtrn2P16(v0 Poly16X4, v1 Poly16X4) Poly16X4 {
	return Poly16X4(C.vtrn2_p16(C.poly16x4_t(v0), C.poly16x4_t(v1)))
}

func Vtrn2QP8(v0 Poly8X16, v1 Poly8X16) Poly8X16 {
	return Poly8X16(C.vtrn2q_p8(C.poly8x16_t(v0), C.poly8x16_t(v1)))
}

func Vtrn2QP64(v0 Poly64X2, v1 Poly64X2) Poly64X2 {
	return Poly64X2(C.vtrn2q_p64(C.poly64x2_t(v0), C.poly64x2_t(v1)))
}

func Vtrn2QP16(v0 Poly16X8, v1 Poly16X8) Poly16X8 {
	return Poly16X8(C.vtrn2q_p16(C.poly16x8_t(v0), C.poly16x8_t(v1)))
}

func Vtrn2QU8(v0 Uint8X16, v1 Uint8X16) Uint8X16 {
	return Uint8X16(C.vtrn2q_u8(C.uint8x16_t(v0), C.uint8x16_t(v1)))
}

func Vtrn2QU32(v0 Uint32X4, v1 Uint32X4) Uint32X4 {
	return Uint32X4(C.vtrn2q_u32(C.uint32x4_t(v0), C.uint32x4_t(v1)))
}

func Vtrn2QU64(v0 Uint64X2, v1 Uint64X2) Uint64X2 {
	return Uint64X2(C.vtrn2q_u64(C.uint64x2_t(v0), C.uint64x2_t(v1)))
}

func Vtrn2QU16(v0 Uint16X8, v1 Uint16X8) Uint16X8 {
	return Uint16X8(C.vtrn2q_u16(C.uint16x8_t(v0), C.uint16x8_t(v1)))
}

func Vtrn2QS8(v0 Int8X16, v1 Int8X16) Int8X16 {
	return Int8X16(C.vtrn2q_s8(C.int8x16_t(v0), C.int8x16_t(v1)))
}

func Vtrn2QF64(v0 Float64X2, v1 Float64X2) Float64X2 {
	return Float64X2(C.vtrn2q_f64(C.float64x2_t(v0), C.float64x2_t(v1)))
}

func Vtrn2QF32(v0 Float32X4, v1 Float32X4) Float32X4 {
	return Float32X4(C.vtrn2q_f32(C.float32x4_t(v0), C.float32x4_t(v1)))
}

func Vtrn2QS32(v0 Int32X4, v1 Int32X4) Int32X4 {
	return Int32X4(C.vtrn2q_s32(C.int32x4_t(v0), C.int32x4_t(v1)))
}

func Vtrn2QS64(v0 Int64X2, v1 Int64X2) Int64X2 {
	return Int64X2(C.vtrn2q_s64(C.int64x2_t(v0), C.int64x2_t(v1)))
}

func Vtrn2QS16(v0 Int16X8, v1 Int16X8) Int16X8 {
	return Int16X8(C.vtrn2q_s16(C.int16x8_t(v0), C.int16x8_t(v1)))
}

func Vtrn2U8(v0 Uint8X8, v1 Uint8X8) Uint8X8 {
	return Uint8X8(C.vtrn2_u8(C.uint8x8_t(v0), C.uint8x8_t(v1)))
}

func Vtrn2U32(v0 Uint32X2, v1 Uint32X2) Uint32X2 {
	return Uint32X2(C.vtrn2_u32(C.uint32x2_t(v0), C.uint32x2_t(v1)))
}

func Vtrn2U16(v0 Uint16X4, v1 Uint16X4) Uint16X4 {
	return Uint16X4(C.vtrn2_u16(C.uint16x4_t(v0), C.uint16x4_t(v1)))
}

func Vtrn2S8(v0 Int8X8, v1 Int8X8) Int8X8 {
	return Int8X8(C.vtrn2_s8(C.int8x8_t(v0), C.int8x8_t(v1)))
}

func Vtrn2F32(v0 Float32X2, v1 Float32X2) Float32X2 {
	return Float32X2(C.vtrn2_f32(C.float32x2_t(v0), C.float32x2_t(v1)))
}

func Vtrn2S32(v0 Int32X2, v1 Int32X2) Int32X2 {
	return Int32X2(C.vtrn2_s32(C.int32x2_t(v0), C.int32x2_t(v1)))
}

func Vtrn2S16(v0 Int16X4, v1 Int16X4) Int16X4 {
	return Int16X4(C.vtrn2_s16(C.int16x4_t(v0), C.int16x4_t(v1)))
}

func VtstP64(v0 Poly64X1, v1 Poly64X1) Uint64X1 {
	return Uint64X1(C.vtst_p64(C.poly64x1_t(v0), C.poly64x1_t(v1)))
}

func VtstqP64(v0 Poly64X2, v1 Poly64X2) Uint64X2 {
	return Uint64X2(C.vtstq_p64(C.poly64x2_t(v0), C.poly64x2_t(v1)))
}

func VtstqU64(v0 Uint64X2, v1 Uint64X2) Uint64X2 {
	return Uint64X2(C.vtstq_u64(C.uint64x2_t(v0), C.uint64x2_t(v1)))
}

func VtstqS64(v0 Int64X2, v1 Int64X2) Uint64X2 {
	return Uint64X2(C.vtstq_s64(C.int64x2_t(v0), C.int64x2_t(v1)))
}

func VtstU64(v0 Uint64X1, v1 Uint64X1) Uint64X1 {
	return Uint64X1(C.vtst_u64(C.uint64x1_t(v0), C.uint64x1_t(v1)))
}

func VtstS64(v0 Int64X1, v1 Int64X1) Uint64X1 {
	return Uint64X1(C.vtst_s64(C.int64x1_t(v0), C.int64x1_t(v1)))
}

func VtstdU64(v0 uint64, v1 uint64) uint64 {
	return uint64(C.vtstd_u64(C.uint64_t(v0), C.uint64_t(v1)))
}

func VtstdS64(v0 int64, v1 int64) uint64 {
	return uint64(C.vtstd_s64(C.int64_t(v0), C.int64_t(v1)))
}

func VuqaddbS8(v0 int8, v1 uint8) int8 {
	return int8(C.vuqaddb_s8(C.int8_t(v0), C.uint8_t(v1)))
}

func VuqaddsS32(v0 int32, v1 uint32) int32 {
	return int32(C.vuqadds_s32(C.int32_t(v0), C.uint32_t(v1)))
}

func VuqadddS64(v0 int64, v1 uint64) int64 {
	return int64(C.vuqaddd_s64(C.int64_t(v0), C.uint64_t(v1)))
}

func VuqaddhS16(v0 int16, v1 uint16) int16 {
	return int16(C.vuqaddh_s16(C.int16_t(v0), C.uint16_t(v1)))
}

func VuqaddqS8(v0 Int8X16, v1 Uint8X16) Int8X16 {
	return Int8X16(C.vuqaddq_s8(C.int8x16_t(v0), C.uint8x16_t(v1)))
}

func VuqaddqS32(v0 Int32X4, v1 Uint32X4) Int32X4 {
	return Int32X4(C.vuqaddq_s32(C.int32x4_t(v0), C.uint32x4_t(v1)))
}

func VuqaddqS64(v0 Int64X2, v1 Uint64X2) Int64X2 {
	return Int64X2(C.vuqaddq_s64(C.int64x2_t(v0), C.uint64x2_t(v1)))
}

func VuqaddqS16(v0 Int16X8, v1 Uint16X8) Int16X8 {
	return Int16X8(C.vuqaddq_s16(C.int16x8_t(v0), C.uint16x8_t(v1)))
}

func VuqaddS8(v0 Int8X8, v1 Uint8X8) Int8X8 {
	return Int8X8(C.vuqadd_s8(C.int8x8_t(v0), C.uint8x8_t(v1)))
}

func VuqaddS32(v0 Int32X2, v1 Uint32X2) Int32X2 {
	return Int32X2(C.vuqadd_s32(C.int32x2_t(v0), C.uint32x2_t(v1)))
}

func VuqaddS64(v0 Int64X1, v1 Uint64X1) Int64X1 {
	return Int64X1(C.vuqadd_s64(C.int64x1_t(v0), C.uint64x1_t(v1)))
}

func VuqaddS16(v0 Int16X4, v1 Uint16X4) Int16X4 {
	return Int16X4(C.vuqadd_s16(C.int16x4_t(v0), C.uint16x4_t(v1)))
}

func Vuzp1P8(v0 Poly8X8, v1 Poly8X8) Poly8X8 {
	return Poly8X8(C.vuzp1_p8(C.poly8x8_t(v0), C.poly8x8_t(v1)))
}

func Vuzp1P16(v0 Poly16X4, v1 Poly16X4) Poly16X4 {
	return Poly16X4(C.vuzp1_p16(C.poly16x4_t(v0), C.poly16x4_t(v1)))
}

func Vuzp1QP8(v0 Poly8X16, v1 Poly8X16) Poly8X16 {
	return Poly8X16(C.vuzp1q_p8(C.poly8x16_t(v0), C.poly8x16_t(v1)))
}

func Vuzp1QP64(v0 Poly64X2, v1 Poly64X2) Poly64X2 {
	return Poly64X2(C.vuzp1q_p64(C.poly64x2_t(v0), C.poly64x2_t(v1)))
}

func Vuzp1QP16(v0 Poly16X8, v1 Poly16X8) Poly16X8 {
	return Poly16X8(C.vuzp1q_p16(C.poly16x8_t(v0), C.poly16x8_t(v1)))
}

func Vuzp1QU8(v0 Uint8X16, v1 Uint8X16) Uint8X16 {
	return Uint8X16(C.vuzp1q_u8(C.uint8x16_t(v0), C.uint8x16_t(v1)))
}

func Vuzp1QU32(v0 Uint32X4, v1 Uint32X4) Uint32X4 {
	return Uint32X4(C.vuzp1q_u32(C.uint32x4_t(v0), C.uint32x4_t(v1)))
}

func Vuzp1QU64(v0 Uint64X2, v1 Uint64X2) Uint64X2 {
	return Uint64X2(C.vuzp1q_u64(C.uint64x2_t(v0), C.uint64x2_t(v1)))
}

func Vuzp1QU16(v0 Uint16X8, v1 Uint16X8) Uint16X8 {
	return Uint16X8(C.vuzp1q_u16(C.uint16x8_t(v0), C.uint16x8_t(v1)))
}

func Vuzp1QS8(v0 Int8X16, v1 Int8X16) Int8X16 {
	return Int8X16(C.vuzp1q_s8(C.int8x16_t(v0), C.int8x16_t(v1)))
}

func Vuzp1QF64(v0 Float64X2, v1 Float64X2) Float64X2 {
	return Float64X2(C.vuzp1q_f64(C.float64x2_t(v0), C.float64x2_t(v1)))
}

func Vuzp1QF32(v0 Float32X4, v1 Float32X4) Float32X4 {
	return Float32X4(C.vuzp1q_f32(C.float32x4_t(v0), C.float32x4_t(v1)))
}

func Vuzp1QS32(v0 Int32X4, v1 Int32X4) Int32X4 {
	return Int32X4(C.vuzp1q_s32(C.int32x4_t(v0), C.int32x4_t(v1)))
}

func Vuzp1QS64(v0 Int64X2, v1 Int64X2) Int64X2 {
	return Int64X2(C.vuzp1q_s64(C.int64x2_t(v0), C.int64x2_t(v1)))
}

func Vuzp1QS16(v0 Int16X8, v1 Int16X8) Int16X8 {
	return Int16X8(C.vuzp1q_s16(C.int16x8_t(v0), C.int16x8_t(v1)))
}

func Vuzp1U8(v0 Uint8X8, v1 Uint8X8) Uint8X8 {
	return Uint8X8(C.vuzp1_u8(C.uint8x8_t(v0), C.uint8x8_t(v1)))
}

func Vuzp1U32(v0 Uint32X2, v1 Uint32X2) Uint32X2 {
	return Uint32X2(C.vuzp1_u32(C.uint32x2_t(v0), C.uint32x2_t(v1)))
}

func Vuzp1U16(v0 Uint16X4, v1 Uint16X4) Uint16X4 {
	return Uint16X4(C.vuzp1_u16(C.uint16x4_t(v0), C.uint16x4_t(v1)))
}

func Vuzp1S8(v0 Int8X8, v1 Int8X8) Int8X8 {
	return Int8X8(C.vuzp1_s8(C.int8x8_t(v0), C.int8x8_t(v1)))
}

func Vuzp1F32(v0 Float32X2, v1 Float32X2) Float32X2 {
	return Float32X2(C.vuzp1_f32(C.float32x2_t(v0), C.float32x2_t(v1)))
}

func Vuzp1S32(v0 Int32X2, v1 Int32X2) Int32X2 {
	return Int32X2(C.vuzp1_s32(C.int32x2_t(v0), C.int32x2_t(v1)))
}

func Vuzp1S16(v0 Int16X4, v1 Int16X4) Int16X4 {
	return Int16X4(C.vuzp1_s16(C.int16x4_t(v0), C.int16x4_t(v1)))
}

func Vuzp2P8(v0 Poly8X8, v1 Poly8X8) Poly8X8 {
	return Poly8X8(C.vuzp2_p8(C.poly8x8_t(v0), C.poly8x8_t(v1)))
}

func Vuzp2P16(v0 Poly16X4, v1 Poly16X4) Poly16X4 {
	return Poly16X4(C.vuzp2_p16(C.poly16x4_t(v0), C.poly16x4_t(v1)))
}

func Vuzp2QP8(v0 Poly8X16, v1 Poly8X16) Poly8X16 {
	return Poly8X16(C.vuzp2q_p8(C.poly8x16_t(v0), C.poly8x16_t(v1)))
}

func Vuzp2QP64(v0 Poly64X2, v1 Poly64X2) Poly64X2 {
	return Poly64X2(C.vuzp2q_p64(C.poly64x2_t(v0), C.poly64x2_t(v1)))
}

func Vuzp2QP16(v0 Poly16X8, v1 Poly16X8) Poly16X8 {
	return Poly16X8(C.vuzp2q_p16(C.poly16x8_t(v0), C.poly16x8_t(v1)))
}

func Vuzp2QU8(v0 Uint8X16, v1 Uint8X16) Uint8X16 {
	return Uint8X16(C.vuzp2q_u8(C.uint8x16_t(v0), C.uint8x16_t(v1)))
}

func Vuzp2QU32(v0 Uint32X4, v1 Uint32X4) Uint32X4 {
	return Uint32X4(C.vuzp2q_u32(C.uint32x4_t(v0), C.uint32x4_t(v1)))
}

func Vuzp2QU64(v0 Uint64X2, v1 Uint64X2) Uint64X2 {
	return Uint64X2(C.vuzp2q_u64(C.uint64x2_t(v0), C.uint64x2_t(v1)))
}

func Vuzp2QU16(v0 Uint16X8, v1 Uint16X8) Uint16X8 {
	return Uint16X8(C.vuzp2q_u16(C.uint16x8_t(v0), C.uint16x8_t(v1)))
}

func Vuzp2QS8(v0 Int8X16, v1 Int8X16) Int8X16 {
	return Int8X16(C.vuzp2q_s8(C.int8x16_t(v0), C.int8x16_t(v1)))
}

func Vuzp2QF64(v0 Float64X2, v1 Float64X2) Float64X2 {
	return Float64X2(C.vuzp2q_f64(C.float64x2_t(v0), C.float64x2_t(v1)))
}

func Vuzp2QF32(v0 Float32X4, v1 Float32X4) Float32X4 {
	return Float32X4(C.vuzp2q_f32(C.float32x4_t(v0), C.float32x4_t(v1)))
}

func Vuzp2QS32(v0 Int32X4, v1 Int32X4) Int32X4 {
	return Int32X4(C.vuzp2q_s32(C.int32x4_t(v0), C.int32x4_t(v1)))
}

func Vuzp2QS64(v0 Int64X2, v1 Int64X2) Int64X2 {
	return Int64X2(C.vuzp2q_s64(C.int64x2_t(v0), C.int64x2_t(v1)))
}

func Vuzp2QS16(v0 Int16X8, v1 Int16X8) Int16X8 {
	return Int16X8(C.vuzp2q_s16(C.int16x8_t(v0), C.int16x8_t(v1)))
}

func Vuzp2U8(v0 Uint8X8, v1 Uint8X8) Uint8X8 {
	return Uint8X8(C.vuzp2_u8(C.uint8x8_t(v0), C.uint8x8_t(v1)))
}

func Vuzp2U32(v0 Uint32X2, v1 Uint32X2) Uint32X2 {
	return Uint32X2(C.vuzp2_u32(C.uint32x2_t(v0), C.uint32x2_t(v1)))
}

func Vuzp2U16(v0 Uint16X4, v1 Uint16X4) Uint16X4 {
	return Uint16X4(C.vuzp2_u16(C.uint16x4_t(v0), C.uint16x4_t(v1)))
}

func Vuzp2S8(v0 Int8X8, v1 Int8X8) Int8X8 {
	return Int8X8(C.vuzp2_s8(C.int8x8_t(v0), C.int8x8_t(v1)))
}

func Vuzp2F32(v0 Float32X2, v1 Float32X2) Float32X2 {
	return Float32X2(C.vuzp2_f32(C.float32x2_t(v0), C.float32x2_t(v1)))
}

func Vuzp2S32(v0 Int32X2, v1 Int32X2) Int32X2 {
	return Int32X2(C.vuzp2_s32(C.int32x2_t(v0), C.int32x2_t(v1)))
}

func Vuzp2S16(v0 Int16X4, v1 Int16X4) Int16X4 {
	return Int16X4(C.vuzp2_s16(C.int16x4_t(v0), C.int16x4_t(v1)))
}

func Vzip1P8(v0 Poly8X8, v1 Poly8X8) Poly8X8 {
	return Poly8X8(C.vzip1_p8(C.poly8x8_t(v0), C.poly8x8_t(v1)))
}

func Vzip1P16(v0 Poly16X4, v1 Poly16X4) Poly16X4 {
	return Poly16X4(C.vzip1_p16(C.poly16x4_t(v0), C.poly16x4_t(v1)))
}

func Vzip1QP8(v0 Poly8X16, v1 Poly8X16) Poly8X16 {
	return Poly8X16(C.vzip1q_p8(C.poly8x16_t(v0), C.poly8x16_t(v1)))
}

func Vzip1QP64(v0 Poly64X2, v1 Poly64X2) Poly64X2 {
	return Poly64X2(C.vzip1q_p64(C.poly64x2_t(v0), C.poly64x2_t(v1)))
}

func Vzip1QP16(v0 Poly16X8, v1 Poly16X8) Poly16X8 {
	return Poly16X8(C.vzip1q_p16(C.poly16x8_t(v0), C.poly16x8_t(v1)))
}

func Vzip1QU8(v0 Uint8X16, v1 Uint8X16) Uint8X16 {
	return Uint8X16(C.vzip1q_u8(C.uint8x16_t(v0), C.uint8x16_t(v1)))
}

func Vzip1QU32(v0 Uint32X4, v1 Uint32X4) Uint32X4 {
	return Uint32X4(C.vzip1q_u32(C.uint32x4_t(v0), C.uint32x4_t(v1)))
}

func Vzip1QU64(v0 Uint64X2, v1 Uint64X2) Uint64X2 {
	return Uint64X2(C.vzip1q_u64(C.uint64x2_t(v0), C.uint64x2_t(v1)))
}

func Vzip1QU16(v0 Uint16X8, v1 Uint16X8) Uint16X8 {
	return Uint16X8(C.vzip1q_u16(C.uint16x8_t(v0), C.uint16x8_t(v1)))
}

func Vzip1QS8(v0 Int8X16, v1 Int8X16) Int8X16 {
	return Int8X16(C.vzip1q_s8(C.int8x16_t(v0), C.int8x16_t(v1)))
}

func Vzip1QF64(v0 Float64X2, v1 Float64X2) Float64X2 {
	return Float64X2(C.vzip1q_f64(C.float64x2_t(v0), C.float64x2_t(v1)))
}

func Vzip1QF32(v0 Float32X4, v1 Float32X4) Float32X4 {
	return Float32X4(C.vzip1q_f32(C.float32x4_t(v0), C.float32x4_t(v1)))
}

func Vzip1QS32(v0 Int32X4, v1 Int32X4) Int32X4 {
	return Int32X4(C.vzip1q_s32(C.int32x4_t(v0), C.int32x4_t(v1)))
}

func Vzip1QS64(v0 Int64X2, v1 Int64X2) Int64X2 {
	return Int64X2(C.vzip1q_s64(C.int64x2_t(v0), C.int64x2_t(v1)))
}

func Vzip1QS16(v0 Int16X8, v1 Int16X8) Int16X8 {
	return Int16X8(C.vzip1q_s16(C.int16x8_t(v0), C.int16x8_t(v1)))
}

func Vzip1U8(v0 Uint8X8, v1 Uint8X8) Uint8X8 {
	return Uint8X8(C.vzip1_u8(C.uint8x8_t(v0), C.uint8x8_t(v1)))
}

func Vzip1U32(v0 Uint32X2, v1 Uint32X2) Uint32X2 {
	return Uint32X2(C.vzip1_u32(C.uint32x2_t(v0), C.uint32x2_t(v1)))
}

func Vzip1U16(v0 Uint16X4, v1 Uint16X4) Uint16X4 {
	return Uint16X4(C.vzip1_u16(C.uint16x4_t(v0), C.uint16x4_t(v1)))
}

func Vzip1S8(v0 Int8X8, v1 Int8X8) Int8X8 {
	return Int8X8(C.vzip1_s8(C.int8x8_t(v0), C.int8x8_t(v1)))
}

func Vzip1F32(v0 Float32X2, v1 Float32X2) Float32X2 {
	return Float32X2(C.vzip1_f32(C.float32x2_t(v0), C.float32x2_t(v1)))
}

func Vzip1S32(v0 Int32X2, v1 Int32X2) Int32X2 {
	return Int32X2(C.vzip1_s32(C.int32x2_t(v0), C.int32x2_t(v1)))
}

func Vzip1S16(v0 Int16X4, v1 Int16X4) Int16X4 {
	return Int16X4(C.vzip1_s16(C.int16x4_t(v0), C.int16x4_t(v1)))
}

func Vzip2P8(v0 Poly8X8, v1 Poly8X8) Poly8X8 {
	return Poly8X8(C.vzip2_p8(C.poly8x8_t(v0), C.poly8x8_t(v1)))
}

func Vzip2P16(v0 Poly16X4, v1 Poly16X4) Poly16X4 {
	return Poly16X4(C.vzip2_p16(C.poly16x4_t(v0), C.poly16x4_t(v1)))
}

func Vzip2QP8(v0 Poly8X16, v1 Poly8X16) Poly8X16 {
	return Poly8X16(C.vzip2q_p8(C.poly8x16_t(v0), C.poly8x16_t(v1)))
}

func Vzip2QP64(v0 Poly64X2, v1 Poly64X2) Poly64X2 {
	return Poly64X2(C.vzip2q_p64(C.poly64x2_t(v0), C.poly64x2_t(v1)))
}

func Vzip2QP16(v0 Poly16X8, v1 Poly16X8) Poly16X8 {
	return Poly16X8(C.vzip2q_p16(C.poly16x8_t(v0), C.poly16x8_t(v1)))
}

func Vzip2QU8(v0 Uint8X16, v1 Uint8X16) Uint8X16 {
	return Uint8X16(C.vzip2q_u8(C.uint8x16_t(v0), C.uint8x16_t(v1)))
}

func Vzip2QU32(v0 Uint32X4, v1 Uint32X4) Uint32X4 {
	return Uint32X4(C.vzip2q_u32(C.uint32x4_t(v0), C.uint32x4_t(v1)))
}

func Vzip2QU64(v0 Uint64X2, v1 Uint64X2) Uint64X2 {
	return Uint64X2(C.vzip2q_u64(C.uint64x2_t(v0), C.uint64x2_t(v1)))
}

func Vzip2QU16(v0 Uint16X8, v1 Uint16X8) Uint16X8 {
	return Uint16X8(C.vzip2q_u16(C.uint16x8_t(v0), C.uint16x8_t(v1)))
}

func Vzip2QS8(v0 Int8X16, v1 Int8X16) Int8X16 {
	return Int8X16(C.vzip2q_s8(C.int8x16_t(v0), C.int8x16_t(v1)))
}

func Vzip2QF64(v0 Float64X2, v1 Float64X2) Float64X2 {
	return Float64X2(C.vzip2q_f64(C.float64x2_t(v0), C.float64x2_t(v1)))
}

func Vzip2QF32(v0 Float32X4, v1 Float32X4) Float32X4 {
	return Float32X4(C.vzip2q_f32(C.float32x4_t(v0), C.float32x4_t(v1)))
}

func Vzip2QS32(v0 Int32X4, v1 Int32X4) Int32X4 {
	return Int32X4(C.vzip2q_s32(C.int32x4_t(v0), C.int32x4_t(v1)))
}

func Vzip2QS64(v0 Int64X2, v1 Int64X2) Int64X2 {
	return Int64X2(C.vzip2q_s64(C.int64x2_t(v0), C.int64x2_t(v1)))
}

func Vzip2QS16(v0 Int16X8, v1 Int16X8) Int16X8 {
	return Int16X8(C.vzip2q_s16(C.int16x8_t(v0), C.int16x8_t(v1)))
}

func Vzip2U8(v0 Uint8X8, v1 Uint8X8) Uint8X8 {
	return Uint8X8(C.vzip2_u8(C.uint8x8_t(v0), C.uint8x8_t(v1)))
}

func Vzip2U32(v0 Uint32X2, v1 Uint32X2) Uint32X2 {
	return Uint32X2(C.vzip2_u32(C.uint32x2_t(v0), C.uint32x2_t(v1)))
}

func Vzip2U16(v0 Uint16X4, v1 Uint16X4) Uint16X4 {
	return Uint16X4(C.vzip2_u16(C.uint16x4_t(v0), C.uint16x4_t(v1)))
}

func Vzip2S8(v0 Int8X8, v1 Int8X8) Int8X8 {
	return Int8X8(C.vzip2_s8(C.int8x8_t(v0), C.int8x8_t(v1)))
}

func Vzip2F32(v0 Float32X2, v1 Float32X2) Float32X2 {
	return Float32X2(C.vzip2_f32(C.float32x2_t(v0), C.float32x2_t(v1)))
}

func Vzip2S32(v0 Int32X2, v1 Int32X2) Int32X2 {
	return Int32X2(C.vzip2_s32(C.int32x2_t(v0), C.int32x2_t(v1)))
}

func Vzip2S16(v0 Int16X4, v1 Int16X4) Int16X4 {
	return Int16X4(C.vzip2_s16(C.int16x4_t(v0), C.int16x4_t(v1)))
}

func VabaqU8(v0 Uint8X16, v1 Uint8X16, v2 Uint8X16) Uint8X16 {
	return Uint8X16(C.vabaq_u8(C.uint8x16_t(v0), C.uint8x16_t(v1), C.uint8x16_t(v2)))
}

func VabaqU32(v0 Uint32X4, v1 Uint32X4, v2 Uint32X4) Uint32X4 {
	return Uint32X4(C.vabaq_u32(C.uint32x4_t(v0), C.uint32x4_t(v1), C.uint32x4_t(v2)))
}

func VabaqU16(v0 Uint16X8, v1 Uint16X8, v2 Uint16X8) Uint16X8 {
	return Uint16X8(C.vabaq_u16(C.uint16x8_t(v0), C.uint16x8_t(v1), C.uint16x8_t(v2)))
}

func VabaqS8(v0 Int8X16, v1 Int8X16, v2 Int8X16) Int8X16 {
	return Int8X16(C.vabaq_s8(C.int8x16_t(v0), C.int8x16_t(v1), C.int8x16_t(v2)))
}

func VabaqS32(v0 Int32X4, v1 Int32X4, v2 Int32X4) Int32X4 {
	return Int32X4(C.vabaq_s32(C.int32x4_t(v0), C.int32x4_t(v1), C.int32x4_t(v2)))
}

func VabaqS16(v0 Int16X8, v1 Int16X8, v2 Int16X8) Int16X8 {
	return Int16X8(C.vabaq_s16(C.int16x8_t(v0), C.int16x8_t(v1), C.int16x8_t(v2)))
}

func VabaU8(v0 Uint8X8, v1 Uint8X8, v2 Uint8X8) Uint8X8 {
	return Uint8X8(C.vaba_u8(C.uint8x8_t(v0), C.uint8x8_t(v1), C.uint8x8_t(v2)))
}

func VabaU32(v0 Uint32X2, v1 Uint32X2, v2 Uint32X2) Uint32X2 {
	return Uint32X2(C.vaba_u32(C.uint32x2_t(v0), C.uint32x2_t(v1), C.uint32x2_t(v2)))
}

func VabaU16(v0 Uint16X4, v1 Uint16X4, v2 Uint16X4) Uint16X4 {
	return Uint16X4(C.vaba_u16(C.uint16x4_t(v0), C.uint16x4_t(v1), C.uint16x4_t(v2)))
}

func VabaS8(v0 Int8X8, v1 Int8X8, v2 Int8X8) Int8X8 {
	return Int8X8(C.vaba_s8(C.int8x8_t(v0), C.int8x8_t(v1), C.int8x8_t(v2)))
}

func VabaS32(v0 Int32X2, v1 Int32X2, v2 Int32X2) Int32X2 {
	return Int32X2(C.vaba_s32(C.int32x2_t(v0), C.int32x2_t(v1), C.int32x2_t(v2)))
}

func VabaS16(v0 Int16X4, v1 Int16X4, v2 Int16X4) Int16X4 {
	return Int16X4(C.vaba_s16(C.int16x4_t(v0), C.int16x4_t(v1), C.int16x4_t(v2)))
}

func VabdlU8(v0 Uint8X8, v1 Uint8X8) Uint16X8 {
	return Uint16X8(C.vabdl_u8(C.uint8x8_t(v0), C.uint8x8_t(v1)))
}

func VabdlU32(v0 Uint32X2, v1 Uint32X2) Uint64X2 {
	return Uint64X2(C.vabdl_u32(C.uint32x2_t(v0), C.uint32x2_t(v1)))
}

func VabdlU16(v0 Uint16X4, v1 Uint16X4) Uint32X4 {
	return Uint32X4(C.vabdl_u16(C.uint16x4_t(v0), C.uint16x4_t(v1)))
}

func VabdlS8(v0 Int8X8, v1 Int8X8) Int16X8 {
	return Int16X8(C.vabdl_s8(C.int8x8_t(v0), C.int8x8_t(v1)))
}

func VabdlS32(v0 Int32X2, v1 Int32X2) Int64X2 {
	return Int64X2(C.vabdl_s32(C.int32x2_t(v0), C.int32x2_t(v1)))
}

func VabdlS16(v0 Int16X4, v1 Int16X4) Int32X4 {
	return Int32X4(C.vabdl_s16(C.int16x4_t(v0), C.int16x4_t(v1)))
}

func VaddlU8(v0 Uint8X8, v1 Uint8X8) Uint16X8 {
	return Uint16X8(C.vaddl_u8(C.uint8x8_t(v0), C.uint8x8_t(v1)))
}

func VaddlU32(v0 Uint32X2, v1 Uint32X2) Uint64X2 {
	return Uint64X2(C.vaddl_u32(C.uint32x2_t(v0), C.uint32x2_t(v1)))
}

func VaddlU16(v0 Uint16X4, v1 Uint16X4) Uint32X4 {
	return Uint32X4(C.vaddl_u16(C.uint16x4_t(v0), C.uint16x4_t(v1)))
}

func VaddlS8(v0 Int8X8, v1 Int8X8) Int16X8 {
	return Int16X8(C.vaddl_s8(C.int8x8_t(v0), C.int8x8_t(v1)))
}

func VaddlS32(v0 Int32X2, v1 Int32X2) Int64X2 {
	return Int64X2(C.vaddl_s32(C.int32x2_t(v0), C.int32x2_t(v1)))
}

func VaddlS16(v0 Int16X4, v1 Int16X4) Int32X4 {
	return Int32X4(C.vaddl_s16(C.int16x4_t(v0), C.int16x4_t(v1)))
}

func VaddwU8(v0 Uint16X8, v1 Uint8X8) Uint16X8 {
	return Uint16X8(C.vaddw_u8(C.uint16x8_t(v0), C.uint8x8_t(v1)))
}

func VaddwU32(v0 Uint64X2, v1 Uint32X2) Uint64X2 {
	return Uint64X2(C.vaddw_u32(C.uint64x2_t(v0), C.uint32x2_t(v1)))
}

func VaddwU16(v0 Uint32X4, v1 Uint16X4) Uint32X4 {
	return Uint32X4(C.vaddw_u16(C.uint32x4_t(v0), C.uint16x4_t(v1)))
}

func VaddwS8(v0 Int16X8, v1 Int8X8) Int16X8 {
	return Int16X8(C.vaddw_s8(C.int16x8_t(v0), C.int8x8_t(v1)))
}

func VaddwS32(v0 Int64X2, v1 Int32X2) Int64X2 {
	return Int64X2(C.vaddw_s32(C.int64x2_t(v0), C.int32x2_t(v1)))
}

func VaddwS16(v0 Int32X4, v1 Int16X4) Int32X4 {
	return Int32X4(C.vaddw_s16(C.int32x4_t(v0), C.int16x4_t(v1)))
}

func VmlalU8(v0 Uint16X8, v1 Uint8X8, v2 Uint8X8) Uint16X8 {
	return Uint16X8(C.vmlal_u8(C.uint16x8_t(v0), C.uint8x8_t(v1), C.uint8x8_t(v2)))
}

func VmlalU32(v0 Uint64X2, v1 Uint32X2, v2 Uint32X2) Uint64X2 {
	return Uint64X2(C.vmlal_u32(C.uint64x2_t(v0), C.uint32x2_t(v1), C.uint32x2_t(v2)))
}

func VmlalU16(v0 Uint32X4, v1 Uint16X4, v2 Uint16X4) Uint32X4 {
	return Uint32X4(C.vmlal_u16(C.uint32x4_t(v0), C.uint16x4_t(v1), C.uint16x4_t(v2)))
}

func VmlalS8(v0 Int16X8, v1 Int8X8, v2 Int8X8) Int16X8 {
	return Int16X8(C.vmlal_s8(C.int16x8_t(v0), C.int8x8_t(v1), C.int8x8_t(v2)))
}

func VmlalS32(v0 Int64X2, v1 Int32X2, v2 Int32X2) Int64X2 {
	return Int64X2(C.vmlal_s32(C.int64x2_t(v0), C.int32x2_t(v1), C.int32x2_t(v2)))
}

func VmlalS16(v0 Int32X4, v1 Int16X4, v2 Int16X4) Int32X4 {
	return Int32X4(C.vmlal_s16(C.int32x4_t(v0), C.int16x4_t(v1), C.int16x4_t(v2)))
}

func VmlalNU32(v0 Uint64X2, v1 Uint32X2, v2 uint32) Uint64X2 {
	return Uint64X2(C.vmlal_n_u32(C.uint64x2_t(v0), C.uint32x2_t(v1), C.uint32_t(v2)))
}

func VmlalNU16(v0 Uint32X4, v1 Uint16X4, v2 uint16) Uint32X4 {
	return Uint32X4(C.vmlal_n_u16(C.uint32x4_t(v0), C.uint16x4_t(v1), C.uint16_t(v2)))
}

func VmlalNS32(v0 Int64X2, v1 Int32X2, v2 int32) Int64X2 {
	return Int64X2(C.vmlal_n_s32(C.int64x2_t(v0), C.int32x2_t(v1), C.int32_t(v2)))
}

func VmlalNS16(v0 Int32X4, v1 Int16X4, v2 int16) Int32X4 {
	return Int32X4(C.vmlal_n_s16(C.int32x4_t(v0), C.int16x4_t(v1), C.int16_t(v2)))
}

func VmlslU8(v0 Uint16X8, v1 Uint8X8, v2 Uint8X8) Uint16X8 {
	return Uint16X8(C.vmlsl_u8(C.uint16x8_t(v0), C.uint8x8_t(v1), C.uint8x8_t(v2)))
}

func VmlslU32(v0 Uint64X2, v1 Uint32X2, v2 Uint32X2) Uint64X2 {
	return Uint64X2(C.vmlsl_u32(C.uint64x2_t(v0), C.uint32x2_t(v1), C.uint32x2_t(v2)))
}

func VmlslU16(v0 Uint32X4, v1 Uint16X4, v2 Uint16X4) Uint32X4 {
	return Uint32X4(C.vmlsl_u16(C.uint32x4_t(v0), C.uint16x4_t(v1), C.uint16x4_t(v2)))
}

func VmlslS8(v0 Int16X8, v1 Int8X8, v2 Int8X8) Int16X8 {
	return Int16X8(C.vmlsl_s8(C.int16x8_t(v0), C.int8x8_t(v1), C.int8x8_t(v2)))
}

func VmlslS32(v0 Int64X2, v1 Int32X2, v2 Int32X2) Int64X2 {
	return Int64X2(C.vmlsl_s32(C.int64x2_t(v0), C.int32x2_t(v1), C.int32x2_t(v2)))
}

func VmlslS16(v0 Int32X4, v1 Int16X4, v2 Int16X4) Int32X4 {
	return Int32X4(C.vmlsl_s16(C.int32x4_t(v0), C.int16x4_t(v1), C.int16x4_t(v2)))
}

func VmlslNU32(v0 Uint64X2, v1 Uint32X2, v2 uint32) Uint64X2 {
	return Uint64X2(C.vmlsl_n_u32(C.uint64x2_t(v0), C.uint32x2_t(v1), C.uint32_t(v2)))
}

func VmlslNU16(v0 Uint32X4, v1 Uint16X4, v2 uint16) Uint32X4 {
	return Uint32X4(C.vmlsl_n_u16(C.uint32x4_t(v0), C.uint16x4_t(v1), C.uint16_t(v2)))
}

func VmlslNS32(v0 Int64X2, v1 Int32X2, v2 int32) Int64X2 {
	return Int64X2(C.vmlsl_n_s32(C.int64x2_t(v0), C.int32x2_t(v1), C.int32_t(v2)))
}

func VmlslNS16(v0 Int32X4, v1 Int16X4, v2 int16) Int32X4 {
	return Int32X4(C.vmlsl_n_s16(C.int32x4_t(v0), C.int16x4_t(v1), C.int16_t(v2)))
}

func VqrdmlahsS32(v0 int32, v1 int32, v2 int32) int32 {
	return int32(C.vqrdmlahs_s32(C.int32_t(v0), C.int32_t(v1), C.int32_t(v2)))
}

func VqrdmlahhS16(v0 int16, v1 int16, v2 int16) int16 {
	return int16(C.vqrdmlahh_s16(C.int16_t(v0), C.int16_t(v1), C.int16_t(v2)))
}

func VqrdmlshsS32(v0 int32, v1 int32, v2 int32) int32 {
	return int32(C.vqrdmlshs_s32(C.int32_t(v0), C.int32_t(v1), C.int32_t(v2)))
}

func VqrdmlshhS16(v0 int16, v1 int16, v2 int16) int16 {
	return int16(C.vqrdmlshh_s16(C.int16_t(v0), C.int16_t(v1), C.int16_t(v2)))
}

func VabdlHighU8(v0 Uint8X16, v1 Uint8X16) Uint16X8 {
	return Uint16X8(C.vabdl_high_u8(C.uint8x16_t(v0), C.uint8x16_t(v1)))
}

func VabdlHighU32(v0 Uint32X4, v1 Uint32X4) Uint64X2 {
	return Uint64X2(C.vabdl_high_u32(C.uint32x4_t(v0), C.uint32x4_t(v1)))
}

func VabdlHighU16(v0 Uint16X8, v1 Uint16X8) Uint32X4 {
	return Uint32X4(C.vabdl_high_u16(C.uint16x8_t(v0), C.uint16x8_t(v1)))
}

func VabdlHighS8(v0 Int8X16, v1 Int8X16) Int16X8 {
	return Int16X8(C.vabdl_high_s8(C.int8x16_t(v0), C.int8x16_t(v1)))
}

func VabdlHighS32(v0 Int32X4, v1 Int32X4) Int64X2 {
	return Int64X2(C.vabdl_high_s32(C.int32x4_t(v0), C.int32x4_t(v1)))
}

func VabdlHighS16(v0 Int16X8, v1 Int16X8) Int32X4 {
	return Int32X4(C.vabdl_high_s16(C.int16x8_t(v0), C.int16x8_t(v1)))
}

func VaddlHighU8(v0 Uint8X16, v1 Uint8X16) Uint16X8 {
	return Uint16X8(C.vaddl_high_u8(C.uint8x16_t(v0), C.uint8x16_t(v1)))
}

func VaddlHighU32(v0 Uint32X4, v1 Uint32X4) Uint64X2 {
	return Uint64X2(C.vaddl_high_u32(C.uint32x4_t(v0), C.uint32x4_t(v1)))
}

func VaddlHighU16(v0 Uint16X8, v1 Uint16X8) Uint32X4 {
	return Uint32X4(C.vaddl_high_u16(C.uint16x8_t(v0), C.uint16x8_t(v1)))
}

func VaddlHighS8(v0 Int8X16, v1 Int8X16) Int16X8 {
	return Int16X8(C.vaddl_high_s8(C.int8x16_t(v0), C.int8x16_t(v1)))
}

func VaddlHighS32(v0 Int32X4, v1 Int32X4) Int64X2 {
	return Int64X2(C.vaddl_high_s32(C.int32x4_t(v0), C.int32x4_t(v1)))
}

func VaddlHighS16(v0 Int16X8, v1 Int16X8) Int32X4 {
	return Int32X4(C.vaddl_high_s16(C.int16x8_t(v0), C.int16x8_t(v1)))
}

func VaddwHighU8(v0 Uint16X8, v1 Uint8X16) Uint16X8 {
	return Uint16X8(C.vaddw_high_u8(C.uint16x8_t(v0), C.uint8x16_t(v1)))
}

func VaddwHighU32(v0 Uint64X2, v1 Uint32X4) Uint64X2 {
	return Uint64X2(C.vaddw_high_u32(C.uint64x2_t(v0), C.uint32x4_t(v1)))
}

func VaddwHighU16(v0 Uint32X4, v1 Uint16X8) Uint32X4 {
	return Uint32X4(C.vaddw_high_u16(C.uint32x4_t(v0), C.uint16x8_t(v1)))
}

func VaddwHighS8(v0 Int16X8, v1 Int8X16) Int16X8 {
	return Int16X8(C.vaddw_high_s8(C.int16x8_t(v0), C.int8x16_t(v1)))
}

func VaddwHighS32(v0 Int64X2, v1 Int32X4) Int64X2 {
	return Int64X2(C.vaddw_high_s32(C.int64x2_t(v0), C.int32x4_t(v1)))
}

func VaddwHighS16(v0 Int32X4, v1 Int16X8) Int32X4 {
	return Int32X4(C.vaddw_high_s16(C.int32x4_t(v0), C.int16x8_t(v1)))
}

func VmlalHighU8(v0 Uint16X8, v1 Uint8X16, v2 Uint8X16) Uint16X8 {
	return Uint16X8(C.vmlal_high_u8(C.uint16x8_t(v0), C.uint8x16_t(v1), C.uint8x16_t(v2)))
}

func VmlalHighU32(v0 Uint64X2, v1 Uint32X4, v2 Uint32X4) Uint64X2 {
	return Uint64X2(C.vmlal_high_u32(C.uint64x2_t(v0), C.uint32x4_t(v1), C.uint32x4_t(v2)))
}

func VmlalHighU16(v0 Uint32X4, v1 Uint16X8, v2 Uint16X8) Uint32X4 {
	return Uint32X4(C.vmlal_high_u16(C.uint32x4_t(v0), C.uint16x8_t(v1), C.uint16x8_t(v2)))
}

func VmlalHighS8(v0 Int16X8, v1 Int8X16, v2 Int8X16) Int16X8 {
	return Int16X8(C.vmlal_high_s8(C.int16x8_t(v0), C.int8x16_t(v1), C.int8x16_t(v2)))
}

func VmlalHighS32(v0 Int64X2, v1 Int32X4, v2 Int32X4) Int64X2 {
	return Int64X2(C.vmlal_high_s32(C.int64x2_t(v0), C.int32x4_t(v1), C.int32x4_t(v2)))
}

func VmlalHighS16(v0 Int32X4, v1 Int16X8, v2 Int16X8) Int32X4 {
	return Int32X4(C.vmlal_high_s16(C.int32x4_t(v0), C.int16x8_t(v1), C.int16x8_t(v2)))
}

func VmlalHighNU32(v0 Uint64X2, v1 Uint32X4, v2 uint32) Uint64X2 {
	return Uint64X2(C.vmlal_high_n_u32(C.uint64x2_t(v0), C.uint32x4_t(v1), C.uint32_t(v2)))
}

func VmlalHighNU16(v0 Uint32X4, v1 Uint16X8, v2 uint16) Uint32X4 {
	return Uint32X4(C.vmlal_high_n_u16(C.uint32x4_t(v0), C.uint16x8_t(v1), C.uint16_t(v2)))
}

func VmlalHighNS32(v0 Int64X2, v1 Int32X4, v2 int32) Int64X2 {
	return Int64X2(C.vmlal_high_n_s32(C.int64x2_t(v0), C.int32x4_t(v1), C.int32_t(v2)))
}

func VmlalHighNS16(v0 Int32X4, v1 Int16X8, v2 int16) Int32X4 {
	return Int32X4(C.vmlal_high_n_s16(C.int32x4_t(v0), C.int16x8_t(v1), C.int16_t(v2)))
}

func VmlslHighU8(v0 Uint16X8, v1 Uint8X16, v2 Uint8X16) Uint16X8 {
	return Uint16X8(C.vmlsl_high_u8(C.uint16x8_t(v0), C.uint8x16_t(v1), C.uint8x16_t(v2)))
}

func VmlslHighU32(v0 Uint64X2, v1 Uint32X4, v2 Uint32X4) Uint64X2 {
	return Uint64X2(C.vmlsl_high_u32(C.uint64x2_t(v0), C.uint32x4_t(v1), C.uint32x4_t(v2)))
}

func VmlslHighU16(v0 Uint32X4, v1 Uint16X8, v2 Uint16X8) Uint32X4 {
	return Uint32X4(C.vmlsl_high_u16(C.uint32x4_t(v0), C.uint16x8_t(v1), C.uint16x8_t(v2)))
}

func VmlslHighS8(v0 Int16X8, v1 Int8X16, v2 Int8X16) Int16X8 {
	return Int16X8(C.vmlsl_high_s8(C.int16x8_t(v0), C.int8x16_t(v1), C.int8x16_t(v2)))
}

func VmlslHighS32(v0 Int64X2, v1 Int32X4, v2 Int32X4) Int64X2 {
	return Int64X2(C.vmlsl_high_s32(C.int64x2_t(v0), C.int32x4_t(v1), C.int32x4_t(v2)))
}

func VmlslHighS16(v0 Int32X4, v1 Int16X8, v2 Int16X8) Int32X4 {
	return Int32X4(C.vmlsl_high_s16(C.int32x4_t(v0), C.int16x8_t(v1), C.int16x8_t(v2)))
}

func VmlslHighNU32(v0 Uint64X2, v1 Uint32X4, v2 uint32) Uint64X2 {
	return Uint64X2(C.vmlsl_high_n_u32(C.uint64x2_t(v0), C.uint32x4_t(v1), C.uint32_t(v2)))
}

func VmlslHighNU16(v0 Uint32X4, v1 Uint16X8, v2 uint16) Uint32X4 {
	return Uint32X4(C.vmlsl_high_n_u16(C.uint32x4_t(v0), C.uint16x8_t(v1), C.uint16_t(v2)))
}

func VmlslHighNS32(v0 Int64X2, v1 Int32X4, v2 int32) Int64X2 {
	return Int64X2(C.vmlsl_high_n_s32(C.int64x2_t(v0), C.int32x4_t(v1), C.int32_t(v2)))
}

func VmlslHighNS16(v0 Int32X4, v1 Int16X8, v2 int16) Int32X4 {
	return Int32X4(C.vmlsl_high_n_s16(C.int32x4_t(v0), C.int16x8_t(v1), C.int16_t(v2)))
}

func VabalU8(v0 Uint16X8, v1 Uint8X8, v2 Uint8X8) Uint16X8 {
	return Uint16X8(C.vabal_u8(C.uint16x8_t(v0), C.uint8x8_t(v1), C.uint8x8_t(v2)))
}

func VabalU32(v0 Uint64X2, v1 Uint32X2, v2 Uint32X2) Uint64X2 {
	return Uint64X2(C.vabal_u32(C.uint64x2_t(v0), C.uint32x2_t(v1), C.uint32x2_t(v2)))
}

func VabalU16(v0 Uint32X4, v1 Uint16X4, v2 Uint16X4) Uint32X4 {
	return Uint32X4(C.vabal_u16(C.uint32x4_t(v0), C.uint16x4_t(v1), C.uint16x4_t(v2)))
}

func VabalS8(v0 Int16X8, v1 Int8X8, v2 Int8X8) Int16X8 {
	return Int16X8(C.vabal_s8(C.int16x8_t(v0), C.int8x8_t(v1), C.int8x8_t(v2)))
}

func VabalS32(v0 Int64X2, v1 Int32X2, v2 Int32X2) Int64X2 {
	return Int64X2(C.vabal_s32(C.int64x2_t(v0), C.int32x2_t(v1), C.int32x2_t(v2)))
}

func VabalS16(v0 Int32X4, v1 Int16X4, v2 Int16X4) Int32X4 {
	return Int32X4(C.vabal_s16(C.int32x4_t(v0), C.int16x4_t(v1), C.int16x4_t(v2)))
}

func VabalHighU8(v0 Uint16X8, v1 Uint8X16, v2 Uint8X16) Uint16X8 {
	return Uint16X8(C.vabal_high_u8(C.uint16x8_t(v0), C.uint8x16_t(v1), C.uint8x16_t(v2)))
}

func VabalHighU32(v0 Uint64X2, v1 Uint32X4, v2 Uint32X4) Uint64X2 {
	return Uint64X2(C.vabal_high_u32(C.uint64x2_t(v0), C.uint32x4_t(v1), C.uint32x4_t(v2)))
}

func VabalHighU16(v0 Uint32X4, v1 Uint16X8, v2 Uint16X8) Uint32X4 {
	return Uint32X4(C.vabal_high_u16(C.uint32x4_t(v0), C.uint16x8_t(v1), C.uint16x8_t(v2)))
}

func VabalHighS8(v0 Int16X8, v1 Int8X16, v2 Int8X16) Int16X8 {
	return Int16X8(C.vabal_high_s8(C.int16x8_t(v0), C.int8x16_t(v1), C.int8x16_t(v2)))
}

func VabalHighS32(v0 Int64X2, v1 Int32X4, v2 Int32X4) Int64X2 {
	return Int64X2(C.vabal_high_s32(C.int64x2_t(v0), C.int32x4_t(v1), C.int32x4_t(v2)))
}

func VabalHighS16(v0 Int32X4, v1 Int16X8, v2 Int16X8) Int32X4 {
	return Int32X4(C.vabal_high_s16(C.int32x4_t(v0), C.int16x8_t(v1), C.int16x8_t(v2)))
}

