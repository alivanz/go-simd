package x86

/*
#cgo CFLAGS: -march=native
#include <immintrin.h>
*/
import "C"

// __andn_u32
func AndnU32(v0 Uint, v1 Uint) Uint {
	return C.__andn_u32(v0, v1)
}

// __bextr_u32
// func BextrU32(v0 Uint, v1 Uint) Uint {
// 	return C.__bextr_u32(v0, v1)
// }

// Extract contiguous bits from unsigned 32-bit integer "a", and store the result in "dst". Extract the number of bits specified by "len", starting at the bit specified by "start".
func BextrU32(v0 Uint, v1 Uint, v2 Uint) Uint {
	return C._bextr_u32(v0, v1, v2)
}

// Extract contiguous bits from unsigned 32-bit integer "a", and store the result in "dst". Extract the number of bits specified by bits 15:8 of "control", starting at the bit specified by bits 0:7 of "control".
// func Bextr2U32(v0 Uint, v1 Uint) Uint {
// 	return C._bextr2_u32(v0, v1)
// }

// __blsi_u32
func BlsiU32(v0 Uint) Uint {
	return C.__blsi_u32(v0)
}

// __blsmsk_u32
func BlsmskU32(v0 Uint) Uint {
	return C.__blsmsk_u32(v0)
}

// __blsr_u32
func BlsrU32(v0 Uint) Uint {
	return C.__blsr_u32(v0)
}

// __andn_u64
func AndnU64(v0 Ulonglong, v1 Ulonglong) Ulonglong {
	return C.__andn_u64(v0, v1)
}

// __bextr_u64
// func BextrU64(v0 Ulonglong, v1 Ulonglong) Ulonglong {
// 	return C.__bextr_u64(v0, v1)
// }

// Extract contiguous bits from unsigned 64-bit integer "a", and store the result in "dst". Extract the number of bits specified by "len", starting at the bit specified by "start".
func BextrU64(v0 Ulonglong, v1 Uint, v2 Uint) Ulonglong {
	return C._bextr_u64(v0, v1, v2)
}

// Extract contiguous bits from unsigned 64-bit integer "a", and store the result in "dst". Extract the number of bits specified by bits 15:8 of "control", starting at the bit specified by bits 0:7 of "control"..
// func Bextr2U64(v0 Ulonglong, v1 Ulonglong) Ulonglong {
// 	return C._bextr2_u64(v0, v1)
// }

// __blsi_u64
func BlsiU64(v0 Ulonglong) Ulonglong {
	return C.__blsi_u64(v0)
}

// __blsmsk_u64
func BlsmskU64(v0 Ulonglong) Ulonglong {
	return C.__blsmsk_u64(v0)
}

// __blsr_u64
func BlsrU64(v0 Ulonglong) Ulonglong {
	return C.__blsr_u64(v0)
}
