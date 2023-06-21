package scanner

import (
	"reflect"
	"testing"
)

func TestTypes(t *testing.T) {
	result, err := Scan([]byte(`
		typedef char int8_t;
		typedef struct int32x4x3_t {
			int32x4_t val[3];
		} int32x4x3_t;
		int func(int a, int b, int c) { return a+b+c; }
	`))
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", result)
	if !reflect.DeepEqual(result, &ScanResult{
		Types: []Type{
			{
				Name: "int8_t",
				Full: "typedef char int8_t;",
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
	}) {
		t.Fail()
	}
}
