package neon

//go:generate go run ../../cmd/generator -package=neon -head arm_neon.h -types -o types.go -- -arch arm64
//go:generate go run ../../cmd/generator -package=neon -head arm_neon.h -funcs -o functions.go -- -arch arm64
