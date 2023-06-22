package neon

//go:generate go run ../../cmd/generator -package=neon -head arm_neon.h -raw raw.h -types types.go -funcs functions.go -- -arch arm64
