package scanner

import (
	"reflect"
	"regexp"
	"testing"
)

func TestAttribute(t *testing.T) {
	reg := regexp.MustCompile(attr + ";")
	result := reg.FindAllString(`
		__attribute__((__vector_size__(32), __aligned__(32)));
		__attribute__((neon_vector_type(8)));
	`, -1)
	ref := []string{
		"__attribute__((__vector_size__(32), __aligned__(32)));",
		"__attribute__((neon_vector_type(8)));",
	}
	t.Log(result)
	t.Log(ref)
	if !reflect.DeepEqual(result, ref) {
		t.Fail()
	}
}

func TestScan(t *testing.T) {
	result, err := Scan([]byte(`
		typedef char int8_t;
		typedef __attribute__((neon_vector_type(8))) int8_t int8x8_t;
		typedef double __m256d __attribute__((__vector_size__(32), __aligned__(32)));
		typedef struct int32x4x3_t {
			int32x4_t val[3];
		} int32x4x3_t;
		int func(int a, int b, int c) { return a+b+c; }
		static __inline__ __m128 __attribute__((__always_inline__, __nodebug__, __target__("mmx, sse"), __min_vector_width__(128))) _mm_move_ss(__m128 __a, __m128 __b) { __a[0] = __b[0]; return __a; }
		static __inline__ long long __attribute__((__always_inline__, __nodebug__, __target__("mmx"), __min_vector_width__(64))) _mm_cvtm64_si64(__m64 __m) { return 1; }
		void lolo(int a, long long b) { }
		void vovo(void) { }
	`))
	if err != nil {
		t.Fatal(err)
	}
	ref := &ScanResult{
		Types: []Type{
			{
				Name: "int8_t",
				Full: "typedef char int8_t;",
			},
			{
				Name:       "int8x8_t",
				Full:       "typedef __attribute__((neon_vector_type(8))) int8_t int8x8_t;",
				Attributes: []string{"neon_vector_type(8)"},
			},
			{
				Name:       "__m256d",
				Full:       "typedef double __m256d __attribute__((__vector_size__(32), __aligned__(32)));",
				Attributes: []string{"__vector_size__(32)", "__aligned__(32)"},
			},
			{
				Name: "int32x4x3_t",
				Full: "typedef struct int32x4x3_t { int32x4_t val[3]; } int32x4x3_t;",
			},
		},
		Functions: []Function{
			{
				Name: "func",
				Return: &Type{
					Name: "int",
					Full: "int",
				},
				Args: []Type{
					{
						Name: "int",
						Full: "int a",
					}, {
						Name: "int",
						Full: "int b",
					}, {
						Name: "int",
						Full: "int c",
					},
				},
			},
			{
				Name:      "_mm_move_ss",
				Attribute: `__always_inline__, __nodebug__, __target__("mmx, sse"), __min_vector_width__(128)`,
				Return: &Type{
					Name: "__m128",
					Full: "__m128",
				},
				Args: []Type{
					{
						Name: "__m128",
						Full: "__m128 __a",
					},
					{
						Name: "__m128",
						Full: "__m128 __b",
					},
				},
			},
			{
				Name:      "_mm_cvtm64_si64",
				Attribute: `__always_inline__, __nodebug__, __target__("mmx"), __min_vector_width__(64)`,
				Return: &Type{
					Name: "longlong",
					Full: "longlong",
				},
				Args: []Type{
					{
						Name: "__m64",
						Full: "__m64 __m",
					},
				},
			},
			{
				Name: "lolo",
				Args: []Type{
					{
						Name: "int",
						Full: "int a",
					}, {
						Name: "longlong",
						Full: "longlong b",
					},
				},
			}, {
				Name: "vovo",
			},
		},
	}
	t.Logf("%+v", result.Functions[4].Return)
	if !reflect.DeepEqual(result.Types, ref.Types) {
		t.Logf("%+v", result.Types)
		t.Logf("%+v", ref.Types)
		t.Fatal()
	}
	if !reflect.DeepEqual(result, ref) {
		t.Logf("%+v", result.Functions)
		t.Logf("%+v", ref.Functions)
		t.Fatal()
	}
}
