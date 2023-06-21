package scanner

import (
	"reflect"
	"testing"
)

func TestTypes(t *testing.T) {
	result, err := Scan([]byte(`
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
		Types: []Type{"int32x4x3_t"},
		Functions: []Function{
			{
				Name:   "func",
				Return: "int",
				Args:   []Type{"int", "int", "int"},
			},
		},
	}) {
		t.Fail()
	}
}
