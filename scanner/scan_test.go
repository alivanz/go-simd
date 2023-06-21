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
				Name: "int8x8_t",
				Full: "typedef __attribute__((neon_vector_type(8))) int8_t int8x8_t;",
			},
			{
				Name: "__m256d",
				Full: "typedef double __m256d __attribute__((__vector_size__(32), __aligned__(32)));",
			},
			{
				Name: "int32x4x3_t",
				Full: "typedef struct int32x4x3_t { int32x4_t val[3]; } int32x4x3_t;",
			},
		},
		Functions: []Function{
			{
				Name: "func",
				Return: Type{
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
		},
	}
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
