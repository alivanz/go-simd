package neon

import (
	"math/rand"
	"reflect"
	"testing"
	"unsafe"
)

func TestMult(t *testing.T) {
	var (
		a      = Int8X8{0, 1, 2, 3, 4, 5, 6, 7}
		b      = Int8X8{7, 6, 5, 4, 3, 2, 1, 0}
		r      = Int8X8{0, 6, 10, 12, 12, 10, 6, 0}
		result Int8X8
	)
	VmulS8(&result, &a, &b)
	if !reflect.DeepEqual(result, r) {
		t.Fatal(result)
	}
}

func TestMultFull(t *testing.T) {
	const N = 1024 * 16
	var (
		a      [N]int8
		b      [N]int8
		ref    [N]int8
		result [N]int8
	)
	for i := 0; i < N; i++ {
		a[i] = int8(rand.Int())
		b[i] = int8(rand.Int())
		ref[i] = a[i] * b[i]
	}
	vmulS8_full(&result[0], &a[0], &b[0], N)
	if !reflect.DeepEqual(result, ref) {
		t.Fail()
	}
}

func BenchmarkMultRef(t *testing.B) {
	const N = 1024 * 16
	var (
		a      [N]int8
		b      [N]int8
		result [N]int8
	)
	for i := 0; i < t.N; i++ {
		for j := 0; j < N; j++ {
			result[j] = a[j] * b[j]
		}
	}
}

func BenchmarkMultSimd(t *testing.B) {
	const N = 1024 * 16
	var (
		a      [N]int8
		b      [N]int8
		result [N]int8
	)
	for i := 0; i < t.N; i++ {
		for j := 0; j < N; j += 8 {
			VmulS8(
				(*Int8X8)(unsafe.Pointer(&result[j])),
				(*Int8X8)(unsafe.Pointer(&a[j])),
				(*Int8X8)(unsafe.Pointer(&b[j])),
			)
		}
	}
}

func BenchmarkMultSimdBypass(t *testing.B) {
	const N = 1024 * 16
	var (
		a      [N]int8
		b      [N]int8
		result [N]int8
	)
	for i := 0; i < t.N; i++ {
		for j := 0; j < N; j += 8 {
			vmulS8_bypass(
				(*Int8X8)(unsafe.Pointer(&result[j])),
				(*Int8X8)(unsafe.Pointer(&a[j])),
				(*Int8X8)(unsafe.Pointer(&b[j])),
			)
		}
	}
}

func BenchmarkMultSimdFull(t *testing.B) {
	const N = 1024 * 16
	var (
		a      [N]int8
		b      [N]int8
		result [N]int8
	)
	for i := 0; i < t.N; i++ {
		vmulS8_full(
			&result[0],
			&a[0],
			&b[0],
			N,
		)
	}
}

func BenchmarkMultSimdCgo(t *testing.B) {
	const N = 1024 * 16
	var (
		a      [N]int8
		b      [N]int8
		result [N]int8
	)
	for i := 0; i < t.N; i++ {
		for j := 0; j < N; j += 8 {
			vmulS8_cgo(
				(*Int8X8)(unsafe.Pointer(&result[j])),
				(*Int8X8)(unsafe.Pointer(&a[j])),
				(*Int8X8)(unsafe.Pointer(&b[j])),
			)
		}
	}
}
