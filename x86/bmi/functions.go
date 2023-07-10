package bmi

import (
	"github.com/alivanz/go-simd/x86"
)

/*
#cgo CFLAGS: -mbmi
#include <immintrin.h>
*/
import "C"

// __andn_u32
//
//go:linkname AndnU32 AndnU32
//go:noescape
func AndnU32(r *x86.Uint, v0 *x86.Uint, v1 *x86.Uint)

// __bextr_u32
//
//go:linkname BextrU32 BextrU32
//go:noescape
func BextrU32(r *x86.Uint, v0 *x86.Uint, v1 *x86.Uint)

// Extract contiguous bits from unsigned 32-bit integer "a", and store the result in "dst". Extract the number of bits specified by "len", starting at the bit specified by "start".
//
//go:linkname BextrU32 BextrU32
//go:noescape
func BextrU32(r *x86.Uint, v0 *x86.Uint, v1 *x86.Uint, v2 *x86.Uint)

// Extract contiguous bits from unsigned 32-bit integer "a", and store the result in "dst". Extract the number of bits specified by bits 15:8 of "control", starting at the bit specified by bits 0:7 of "control".
//
//go:linkname Bextr2U32 Bextr2U32
//go:noescape
func Bextr2U32(r *x86.Uint, v0 *x86.Uint, v1 *x86.Uint)

// __blsi_u32
//
//go:linkname BlsiU32 BlsiU32
//go:noescape
func BlsiU32(r *x86.Uint, v0 *x86.Uint)

// __blsmsk_u32
//
//go:linkname BlsmskU32 BlsmskU32
//go:noescape
func BlsmskU32(r *x86.Uint, v0 *x86.Uint)

// __blsr_u32
//
//go:linkname BlsrU32 BlsrU32
//go:noescape
func BlsrU32(r *x86.Uint, v0 *x86.Uint)

// __andn_u64
//
//go:linkname AndnU64 AndnU64
//go:noescape
func AndnU64(r *x86.Ulonglong, v0 *x86.Ulonglong, v1 *x86.Ulonglong)

// __bextr_u64
//
//go:linkname BextrU64 BextrU64
//go:noescape
func BextrU64(r *x86.Ulonglong, v0 *x86.Ulonglong, v1 *x86.Ulonglong)

// Extract contiguous bits from unsigned 64-bit integer "a", and store the result in "dst". Extract the number of bits specified by "len", starting at the bit specified by "start".
//
//go:linkname BextrU64 BextrU64
//go:noescape
func BextrU64(r *x86.Ulonglong, v0 *x86.Ulonglong, v1 *x86.Uint, v2 *x86.Uint)

// Extract contiguous bits from unsigned 64-bit integer "a", and store the result in "dst". Extract the number of bits specified by bits 15:8 of "control", starting at the bit specified by bits 0:7 of "control"..
//
//go:linkname Bextr2U64 Bextr2U64
//go:noescape
func Bextr2U64(r *x86.Ulonglong, v0 *x86.Ulonglong, v1 *x86.Ulonglong)

// __blsi_u64
//
//go:linkname BlsiU64 BlsiU64
//go:noescape
func BlsiU64(r *x86.Ulonglong, v0 *x86.Ulonglong)

// __blsmsk_u64
//
//go:linkname BlsmskU64 BlsmskU64
//go:noescape
func BlsmskU64(r *x86.Ulonglong, v0 *x86.Ulonglong)

// __blsr_u64
//
//go:linkname BlsrU64 BlsrU64
//go:noescape
func BlsrU64(r *x86.Ulonglong, v0 *x86.Ulonglong)
