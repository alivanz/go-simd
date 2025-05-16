package neon

import (
	"math/rand"
	"reflect"
	"testing"
	"unsafe"

	"github.com/alivanz/go-simd/arm"
)

func TestVabsS32N(t *testing.T) {
	const N = 1024 * 16
	var (
		r   = make([]arm.Int32, N)
		v   = make([]arm.Int32, N)
		ref = make([]arm.Int32, N)
	)
	for i := 0; i < N; i++ {
		r[i] = arm.Int32(int32(rand.Int()))
		v[i] = arm.Int32(int32(rand.Int()))
		if v[i] < 0 {
			ref[i] = -v[i]
		} else {
			ref[i] = v[i]
		}
	}
	VabsS32N(&r[0], &v[0], N)
	if !reflect.DeepEqual(r, ref) {
		t.Fatal(r)
	}
}

func TestVmulqF32N(t *testing.T) {
	const N = 1024 * 16
	var (
		r   = make([]arm.Float32, N)
		v1  = make([]arm.Float32, N)
		v2  = make([]arm.Float32, N)
		ref = make([]arm.Float32, N)
	)
	for i := 0; i < N; i++ {
		v1[i] = arm.Float32(rand.Float32())
		v2[i] = arm.Float32(rand.Float32())
		ref[i] = v1[i] * v2[i]
	}
	VmulqF32N(&r[0], &v1[0], &v2[0], N)
	if !reflect.DeepEqual(r, ref) {
		t.Fatal(r)
	}
}

// this benchmark is fully run on C code
func BenchmarkVmulqF32N(b *testing.B) {
	const N = 1024 * 1024
	var (
		r  = make([]arm.Float32, N)
		v1 = make([]arm.Float32, N)
		v2 = make([]arm.Float32, N)
	)
	b.SetBytes(N * 4)
	for i := int32(0); i < N; i++ {
		v1[i] = arm.Float32(rand.Float32())
		v2[i] = arm.Float32(rand.Float32())
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		VmulqF32N(&r[0], &v1[0], &v2[0], N)
	}
}

// this benchmark is calling the C code multiple times
func BenchmarkVmulqF32C(b *testing.B) {
	const N = 1024 * 1024
	var (
		r  = make([]arm.Float32, N)
		v1 = make([]arm.Float32, N)
		v2 = make([]arm.Float32, N)
	)
	b.SetBytes(N * 4)
	for i := int32(0); i < N; i++ {
		v1[i] = arm.Float32(rand.Float32())
		v2[i] = arm.Float32(rand.Float32())
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		for j := int32(0); j < N; j += 4 {
			VmulqF32(
				(*arm.Float32X4)(unsafe.Pointer(&r[j])),
				(*arm.Float32X4)(unsafe.Pointer(&v1[j])),
				(*arm.Float32X4)(unsafe.Pointer(&v2[j])),
			)
		}
	}
}

// this benchmark is Go runtime implementation
func BenchmarkVmulqF32Ref(b *testing.B) {
	const N = 1024 * 1024
	var (
		r  = make([]arm.Float32, N)
		v1 = make([]arm.Float32, N)
		v2 = make([]arm.Float32, N)
	)
	b.SetBytes(N * 4)
	for i := int32(0); i < N; i++ {
		v1[i] = arm.Float32(rand.Float32())
		v2[i] = arm.Float32(rand.Float32())
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		for j := int32(0); j < N; j++ {
			r[j] = v1[j] * v2[j]
		}
	}
}
