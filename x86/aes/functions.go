package aes

import (
	"github.com/alivanz/go-simd/x86"
)

/*
#cgo CFLAGS: -maes
#include <immintrin.h>
*/
import "C"

// Perform one round of an AES encryption flow on data (state) in "a" using the round key in "RoundKey", and store the result in "dst"."
//
//go:linkname MmAesencSi128 MmAesencSi128
//go:noescape
func MmAesencSi128(r *x86.M128I, v0 *x86.M128I, v1 *x86.M128I)

// Perform the last round of an AES encryption flow on data (state) in "a" using the round key in "RoundKey", and store the result in "dst"."
//
//go:linkname MmAesenclastSi128 MmAesenclastSi128
//go:noescape
func MmAesenclastSi128(r *x86.M128I, v0 *x86.M128I, v1 *x86.M128I)

// Perform one round of an AES decryption flow on data (state) in "a" using the round key in "RoundKey", and store the result in "dst".
//
//go:linkname MmAesdecSi128 MmAesdecSi128
//go:noescape
func MmAesdecSi128(r *x86.M128I, v0 *x86.M128I, v1 *x86.M128I)

// Perform the last round of an AES decryption flow on data (state) in "a" using the round key in "RoundKey", and store the result in "dst".
//
//go:linkname MmAesdeclastSi128 MmAesdeclastSi128
//go:noescape
func MmAesdeclastSi128(r *x86.M128I, v0 *x86.M128I, v1 *x86.M128I)

// Perform the InvMixColumns transformation on "a" and store the result in "dst".
//
//go:linkname MmAesimcSi128 MmAesimcSi128
//go:noescape
func MmAesimcSi128(r *x86.M128I, v0 *x86.M128I)
