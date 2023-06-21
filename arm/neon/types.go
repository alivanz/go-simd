package neon

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

// typedef long long int64_t;
type Int64 = C.int64_t

// typedef unsigned char uint8_t;
type Uint8 = C.uint8_t

// typedef unsigned short uint16_t;
type Uint16 = C.uint16_t

// typedef unsigned int uint32_t;
type Uint32 = C.uint32_t

// typedef unsigned long long uint64_t;
type Uint64 = C.uint64_t

// typedef int8_t int_least8_t;
type IntLeast8 = C.int_least8_t

// typedef int16_t int_least16_t;
type IntLeast16 = C.int_least16_t

// typedef int32_t int_least32_t;
type IntLeast32 = C.int_least32_t

// typedef int64_t int_least64_t;
type IntLeast64 = C.int_least64_t

// typedef uint8_t uint_least8_t;
type UintLeast8 = C.uint_least8_t

// typedef uint16_t uint_least16_t;
type UintLeast16 = C.uint_least16_t

// typedef uint32_t uint_least32_t;
type UintLeast32 = C.uint_least32_t

// typedef uint64_t uint_least64_t;
type UintLeast64 = C.uint_least64_t

// typedef int8_t int_fast8_t;
type IntFast8 = C.int_fast8_t

// typedef int16_t int_fast16_t;
type IntFast16 = C.int_fast16_t

// typedef int32_t int_fast32_t;
type IntFast32 = C.int_fast32_t

// typedef int64_t int_fast64_t;
type IntFast64 = C.int_fast64_t

// typedef uint8_t uint_fast8_t;
type UintFast8 = C.uint_fast8_t

// typedef uint16_t uint_fast16_t;
type UintFast16 = C.uint_fast16_t

// typedef uint32_t uint_fast32_t;
type UintFast32 = C.uint_fast32_t

// typedef uint64_t uint_fast64_t;
type UintFast64 = C.uint_fast64_t

// typedef unsigned char u_int8_t;
type UInt8 = C.u_int8_t

// typedef unsigned short u_int16_t;
type UInt16 = C.u_int16_t

// typedef unsigned int u_int32_t;
type UInt32 = C.u_int32_t

// typedef unsigned long long u_int64_t;
type UInt64 = C.u_int64_t

// typedef int64_t register_t;
type Register = C.register_t

// typedef unsigned long uintptr_t;
type Uintptr = C.uintptr_t

// typedef u_int64_t user_addr_t;
type UserAddr = C.user_addr_t

// typedef u_int64_t user_size_t;
type UserSize = C.user_size_t

// typedef int64_t user_ssize_t;
type UserSsize = C.user_ssize_t

// typedef int64_t user_long_t;
type UserLong = C.user_long_t

// typedef u_int64_t user_ulong_t;
type UserUlong = C.user_ulong_t

// typedef int64_t user_time_t;
type UserTime = C.user_time_t

// typedef int64_t user_off_t;
type UserOff = C.user_off_t

// typedef u_int64_t syscall_arg_t;
type SyscallArg = C.syscall_arg_t

// typedef __darwin_intptr_t intptr_t;
type Intptr = C.intptr_t

// typedef long int intmax_t;
type Intmax = C.intmax_t

// typedef long unsigned int uintmax_t;
type Uintmax = C.uintmax_t

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

// typedef struct int8x8x2_t {  int8x8_t val[2];} int8x8x2_t;
type Int8X8X2 = C.int8x8x2_t

// typedef struct int8x16x2_t {  int8x16_t val[2];} int8x16x2_t;
type Int8X16X2 = C.int8x16x2_t

// typedef struct int16x4x2_t {  int16x4_t val[2];} int16x4x2_t;
type Int16X4X2 = C.int16x4x2_t

// typedef struct int16x8x2_t {  int16x8_t val[2];} int16x8x2_t;
type Int16X8X2 = C.int16x8x2_t

// typedef struct int32x2x2_t {  int32x2_t val[2];} int32x2x2_t;
type Int32X2X2 = C.int32x2x2_t

// typedef struct int32x4x2_t {  int32x4_t val[2];} int32x4x2_t;
type Int32X4X2 = C.int32x4x2_t

// typedef struct int64x1x2_t {  int64x1_t val[2];} int64x1x2_t;
type Int64X1X2 = C.int64x1x2_t

// typedef struct int64x2x2_t {  int64x2_t val[2];} int64x2x2_t;
type Int64X2X2 = C.int64x2x2_t

// typedef struct uint8x8x2_t {  uint8x8_t val[2];} uint8x8x2_t;
type Uint8X8X2 = C.uint8x8x2_t

// typedef struct uint8x16x2_t {  uint8x16_t val[2];} uint8x16x2_t;
type Uint8X16X2 = C.uint8x16x2_t

// typedef struct uint16x4x2_t {  uint16x4_t val[2];} uint16x4x2_t;
type Uint16X4X2 = C.uint16x4x2_t

// typedef struct uint16x8x2_t {  uint16x8_t val[2];} uint16x8x2_t;
type Uint16X8X2 = C.uint16x8x2_t

// typedef struct uint32x2x2_t {  uint32x2_t val[2];} uint32x2x2_t;
type Uint32X2X2 = C.uint32x2x2_t

// typedef struct uint32x4x2_t {  uint32x4_t val[2];} uint32x4x2_t;
type Uint32X4X2 = C.uint32x4x2_t

// typedef struct uint64x1x2_t {  uint64x1_t val[2];} uint64x1x2_t;
type Uint64X1X2 = C.uint64x1x2_t

// typedef struct uint64x2x2_t {  uint64x2_t val[2];} uint64x2x2_t;
type Uint64X2X2 = C.uint64x2x2_t

// typedef struct float32x2x2_t {  float32x2_t val[2];} float32x2x2_t;
type Float32X2X2 = C.float32x2x2_t

// typedef struct float32x4x2_t {  float32x4_t val[2];} float32x4x2_t;
type Float32X4X2 = C.float32x4x2_t

// typedef struct float64x1x2_t {  float64x1_t val[2];} float64x1x2_t;
type Float64X1X2 = C.float64x1x2_t

// typedef struct float64x2x2_t {  float64x2_t val[2];} float64x2x2_t;
type Float64X2X2 = C.float64x2x2_t

// typedef struct poly8x8x2_t {  poly8x8_t val[2];} poly8x8x2_t;
type Poly8X8X2 = C.poly8x8x2_t

// typedef struct poly8x16x2_t {  poly8x16_t val[2];} poly8x16x2_t;
type Poly8X16X2 = C.poly8x16x2_t

// typedef struct poly16x4x2_t {  poly16x4_t val[2];} poly16x4x2_t;
type Poly16X4X2 = C.poly16x4x2_t

// typedef struct poly16x8x2_t {  poly16x8_t val[2];} poly16x8x2_t;
type Poly16X8X2 = C.poly16x8x2_t

// typedef struct poly64x1x2_t {  poly64x1_t val[2];} poly64x1x2_t;
type Poly64X1X2 = C.poly64x1x2_t

// typedef struct poly64x2x2_t {  poly64x2_t val[2];} poly64x2x2_t;
type Poly64X2X2 = C.poly64x2x2_t

// typedef struct int8x8x3_t {  int8x8_t val[3];} int8x8x3_t;
type Int8X8X3 = C.int8x8x3_t

// typedef struct int8x16x3_t {  int8x16_t val[3];} int8x16x3_t;
type Int8X16X3 = C.int8x16x3_t

// typedef struct int16x4x3_t {  int16x4_t val[3];} int16x4x3_t;
type Int16X4X3 = C.int16x4x3_t

// typedef struct int16x8x3_t {  int16x8_t val[3];} int16x8x3_t;
type Int16X8X3 = C.int16x8x3_t

// typedef struct int32x2x3_t {  int32x2_t val[3];} int32x2x3_t;
type Int32X2X3 = C.int32x2x3_t

// typedef struct int32x4x3_t {  int32x4_t val[3];} int32x4x3_t;
type Int32X4X3 = C.int32x4x3_t

// typedef struct int64x1x3_t {  int64x1_t val[3];} int64x1x3_t;
type Int64X1X3 = C.int64x1x3_t

// typedef struct int64x2x3_t {  int64x2_t val[3];} int64x2x3_t;
type Int64X2X3 = C.int64x2x3_t

// typedef struct uint8x8x3_t {  uint8x8_t val[3];} uint8x8x3_t;
type Uint8X8X3 = C.uint8x8x3_t

// typedef struct uint8x16x3_t {  uint8x16_t val[3];} uint8x16x3_t;
type Uint8X16X3 = C.uint8x16x3_t

// typedef struct uint16x4x3_t {  uint16x4_t val[3];} uint16x4x3_t;
type Uint16X4X3 = C.uint16x4x3_t

// typedef struct uint16x8x3_t {  uint16x8_t val[3];} uint16x8x3_t;
type Uint16X8X3 = C.uint16x8x3_t

// typedef struct uint32x2x3_t {  uint32x2_t val[3];} uint32x2x3_t;
type Uint32X2X3 = C.uint32x2x3_t

// typedef struct uint32x4x3_t {  uint32x4_t val[3];} uint32x4x3_t;
type Uint32X4X3 = C.uint32x4x3_t

// typedef struct uint64x1x3_t {  uint64x1_t val[3];} uint64x1x3_t;
type Uint64X1X3 = C.uint64x1x3_t

// typedef struct uint64x2x3_t {  uint64x2_t val[3];} uint64x2x3_t;
type Uint64X2X3 = C.uint64x2x3_t

// typedef struct float32x2x3_t {  float32x2_t val[3];} float32x2x3_t;
type Float32X2X3 = C.float32x2x3_t

// typedef struct float32x4x3_t {  float32x4_t val[3];} float32x4x3_t;
type Float32X4X3 = C.float32x4x3_t

// typedef struct float64x1x3_t {  float64x1_t val[3];} float64x1x3_t;
type Float64X1X3 = C.float64x1x3_t

// typedef struct float64x2x3_t {  float64x2_t val[3];} float64x2x3_t;
type Float64X2X3 = C.float64x2x3_t

// typedef struct poly8x8x3_t {  poly8x8_t val[3];} poly8x8x3_t;
type Poly8X8X3 = C.poly8x8x3_t

// typedef struct poly8x16x3_t {  poly8x16_t val[3];} poly8x16x3_t;
type Poly8X16X3 = C.poly8x16x3_t

// typedef struct poly16x4x3_t {  poly16x4_t val[3];} poly16x4x3_t;
type Poly16X4X3 = C.poly16x4x3_t

// typedef struct poly16x8x3_t {  poly16x8_t val[3];} poly16x8x3_t;
type Poly16X8X3 = C.poly16x8x3_t

// typedef struct poly64x1x3_t {  poly64x1_t val[3];} poly64x1x3_t;
type Poly64X1X3 = C.poly64x1x3_t

// typedef struct poly64x2x3_t {  poly64x2_t val[3];} poly64x2x3_t;
type Poly64X2X3 = C.poly64x2x3_t

// typedef struct int8x8x4_t {  int8x8_t val[4];} int8x8x4_t;
type Int8X8X4 = C.int8x8x4_t

// typedef struct int8x16x4_t {  int8x16_t val[4];} int8x16x4_t;
type Int8X16X4 = C.int8x16x4_t

// typedef struct int16x4x4_t {  int16x4_t val[4];} int16x4x4_t;
type Int16X4X4 = C.int16x4x4_t

// typedef struct int16x8x4_t {  int16x8_t val[4];} int16x8x4_t;
type Int16X8X4 = C.int16x8x4_t

// typedef struct int32x2x4_t {  int32x2_t val[4];} int32x2x4_t;
type Int32X2X4 = C.int32x2x4_t

// typedef struct int32x4x4_t {  int32x4_t val[4];} int32x4x4_t;
type Int32X4X4 = C.int32x4x4_t

// typedef struct int64x1x4_t {  int64x1_t val[4];} int64x1x4_t;
type Int64X1X4 = C.int64x1x4_t

// typedef struct int64x2x4_t {  int64x2_t val[4];} int64x2x4_t;
type Int64X2X4 = C.int64x2x4_t

// typedef struct uint8x8x4_t {  uint8x8_t val[4];} uint8x8x4_t;
type Uint8X8X4 = C.uint8x8x4_t

// typedef struct uint8x16x4_t {  uint8x16_t val[4];} uint8x16x4_t;
type Uint8X16X4 = C.uint8x16x4_t

// typedef struct uint16x4x4_t {  uint16x4_t val[4];} uint16x4x4_t;
type Uint16X4X4 = C.uint16x4x4_t

// typedef struct uint16x8x4_t {  uint16x8_t val[4];} uint16x8x4_t;
type Uint16X8X4 = C.uint16x8x4_t

// typedef struct uint32x2x4_t {  uint32x2_t val[4];} uint32x2x4_t;
type Uint32X2X4 = C.uint32x2x4_t

// typedef struct uint32x4x4_t {  uint32x4_t val[4];} uint32x4x4_t;
type Uint32X4X4 = C.uint32x4x4_t

// typedef struct uint64x1x4_t {  uint64x1_t val[4];} uint64x1x4_t;
type Uint64X1X4 = C.uint64x1x4_t

// typedef struct uint64x2x4_t {  uint64x2_t val[4];} uint64x2x4_t;
type Uint64X2X4 = C.uint64x2x4_t

// typedef struct float32x2x4_t {  float32x2_t val[4];} float32x2x4_t;
type Float32X2X4 = C.float32x2x4_t

// typedef struct float32x4x4_t {  float32x4_t val[4];} float32x4x4_t;
type Float32X4X4 = C.float32x4x4_t

// typedef struct float64x1x4_t {  float64x1_t val[4];} float64x1x4_t;
type Float64X1X4 = C.float64x1x4_t

// typedef struct float64x2x4_t {  float64x2_t val[4];} float64x2x4_t;
type Float64X2X4 = C.float64x2x4_t

// typedef struct poly8x8x4_t {  poly8x8_t val[4];} poly8x8x4_t;
type Poly8X8X4 = C.poly8x8x4_t

// typedef struct poly8x16x4_t {  poly8x16_t val[4];} poly8x16x4_t;
type Poly8X16X4 = C.poly8x16x4_t

// typedef struct poly16x4x4_t {  poly16x4_t val[4];} poly16x4x4_t;
type Poly16X4X4 = C.poly16x4x4_t

// typedef struct poly16x8x4_t {  poly16x8_t val[4];} poly16x8x4_t;
type Poly16X8X4 = C.poly16x8x4_t

// typedef struct poly64x1x4_t {  poly64x1_t val[4];} poly64x1x4_t;
type Poly64X1X4 = C.poly64x1x4_t

// typedef struct poly64x2x4_t {  poly64x2_t val[4];} poly64x2x4_t;
type Poly64X2X4 = C.poly64x2x4_t
