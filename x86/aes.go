package x86

/*
#cgo CFLAGS: -march=native
#include <immintrin.h>
*/
import "C"

// Perform one round of an AES encryption flow on data (state) in "a" using the round key in "RoundKey", and store the result in "dst"."
func MmAesencSi128(v0 M128I, v1 M128I) M128I {
	return C._mm_aesenc_si128(v0, v1)
}

// Perform the last round of an AES encryption flow on data (state) in "a" using the round key in "RoundKey", and store the result in "dst"."
func MmAesenclastSi128(v0 M128I, v1 M128I) M128I {
	return C._mm_aesenclast_si128(v0, v1)
}

// Perform one round of an AES decryption flow on data (state) in "a" using the round key in "RoundKey", and store the result in "dst".
func MmAesdecSi128(v0 M128I, v1 M128I) M128I {
	return C._mm_aesdec_si128(v0, v1)
}

// Perform the last round of an AES decryption flow on data (state) in "a" using the round key in "RoundKey", and store the result in "dst".
func MmAesdeclastSi128(v0 M128I, v1 M128I) M128I {
	return C._mm_aesdeclast_si128(v0, v1)
}

// Perform the InvMixColumns transformation on "a" and store the result in "dst".
func MmAesimcSi128(v0 M128I) M128I {
	return C._mm_aesimc_si128(v0)
}
