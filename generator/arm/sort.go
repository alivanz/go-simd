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

func sortGroup(name string) (string, int, string) {
	var (
		group  = name
		index  = -1
		suffix = ""
	)
	for i, s := range suffixOrder {
		if strings.HasSuffix(name, s) {
			group = strings.TrimSuffix(name, s)
			index = i
			suffix = s
		}
	}
	return group, index, suffix
}
