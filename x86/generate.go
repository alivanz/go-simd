package x86

//go:generate go run ../cmd/generator -package=x86 -head immintrin.h -raw raw.h -types types.go -funcs functions.go -- -arch x86_64
