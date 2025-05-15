package main

import "strings"

var (
	suffixOrder = []string{
		"_s8",
		"_s16",
		"_s32",
		"_s64",
		"_u8",
		"_u16",
		"_u32",
		"_u64",
		"_f32",
		"_f64",
	}
)

func sortGroup(name string) (string, int) {
	var (
		group = name
		index = -1
	)
	for i, suffix := range suffixOrder {
		if strings.HasSuffix(name, suffix) {
			group = strings.TrimSuffix(name, suffix)
			index = i
		}
	}
	return group, index
}
