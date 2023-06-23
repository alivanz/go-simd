package x86

/*
#cgo CFLAGS: -march=native
#include <immintrin.h>
*/
import "C"

// __andn_u32
// __always_inline__, __nodebug__, __target__("bmi")
func AndnU32(v0 Uint, v1 Uint) Uint {
	return C.__andn_u32(v0, v1)
}

// __bextr_u32
// __always_inline__, __nodebug__, __target__("bmi")
// func BextrU32(v0 Uint, v1 Uint) Uint {
// 	return C.__bextr_u32(v0, v1)
// }

// _bextr_u32
// __always_inline__, __nodebug__, __target__("bmi")
func BextrU32(v0 Uint, v1 Uint, v2 Uint) Uint {
	return C._bextr_u32(v0, v1, v2)
}

// _bextr2_u32
// __always_inline__, __nodebug__, __target__("bmi")
// func Bextr2U32(v0 Uint, v1 Uint) Uint {
// 	return C._bextr2_u32(v0, v1)
// }

// __blsi_u32
// __always_inline__, __nodebug__, __target__("bmi")
func BlsiU32(v0 Uint) Uint {
	return C.__blsi_u32(v0)
}

// __blsmsk_u32
// __always_inline__, __nodebug__, __target__("bmi")
func BlsmskU32(v0 Uint) Uint {
	return C.__blsmsk_u32(v0)
}

// __blsr_u32
// __always_inline__, __nodebug__, __target__("bmi")
func BlsrU32(v0 Uint) Uint {
	return C.__blsr_u32(v0)
}

// __andn_u64
// __always_inline__, __nodebug__, __target__("bmi")
func AndnU64(v0 Ulonglong, v1 Ulonglong) Ulonglong {
	return C.__andn_u64(v0, v1)
}

// __bextr_u64
// __always_inline__, __nodebug__, __target__("bmi")
// func BextrU64(v0 Ulonglong, v1 Ulonglong) Ulonglong {
// 	return C.__bextr_u64(v0, v1)
// }

// _bextr_u64
// __always_inline__, __nodebug__, __target__("bmi")
func BextrU64(v0 Ulonglong, v1 Uint, v2 Uint) Ulonglong {
	return C._bextr_u64(v0, v1, v2)
}

// _bextr2_u64
// __always_inline__, __nodebug__, __target__("bmi")
// func Bextr2U64(v0 Ulonglong, v1 Ulonglong) Ulonglong {
// 	return C._bextr2_u64(v0, v1)
// }

// __blsi_u64
// __always_inline__, __nodebug__, __target__("bmi")
func BlsiU64(v0 Ulonglong) Ulonglong {
	return C.__blsi_u64(v0)
}

// __blsmsk_u64
// __always_inline__, __nodebug__, __target__("bmi")
func BlsmskU64(v0 Ulonglong) Ulonglong {
	return C.__blsmsk_u64(v0)
}

// __blsr_u64
// __always_inline__, __nodebug__, __target__("bmi")
func BlsrU64(v0 Ulonglong) Ulonglong {
	return C.__blsr_u64(v0)
}
