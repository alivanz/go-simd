package x86

/*
#cgo CFLAGS: -march=native
#include <immintrin.h>
*/
import "C"

// typedef longlong __m64 __attribute__((__vector_size__(8), __aligned__(8)));
type M64 = C.__m64

// typedef float __m128 __attribute__((__vector_size__(16), __aligned__(16)));
type M128 = C.__m128

// typedef double __m128d __attribute__((__vector_size__(16), __aligned__(16)));
type M128D = C.__m128d

// typedef longlong __m128i __attribute__((__vector_size__(16), __aligned__(16)));
type M128I = C.__m128i

// typedef double __m256d __attribute__((__vector_size__(32), __aligned__(32)));
type M256D = C.__m256d

// typedef longlong __m256i __attribute__((__vector_size__(32), __aligned__(32)));
type M256I = C.__m256i

// uint
type Uint = C.uint

// uchar __D
type Uchar = C.uchar

// ushort __D
type Ushort = C.ushort

// ulonglong
type Ulonglong = C.ulonglong

// int __i
type Int = C.int

// longlong __i
type Longlong = C.longlong

// short __s3
type Short = C.short

// char __b7
type Char = C.char

// float
type Float = C.float

// double
type Double = C.double

// unsigned __extensions
type Unsigned = C.unsigned

// __m256
type M256 = C.__m256
