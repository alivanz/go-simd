package scanner

func transform[A, B any](arr []A, fn func(i int, e A) B) []B {
	out := make([]B, len(arr))
	for i, e := range arr {
		out[i] = fn(i, e)
	}
	return out
}
