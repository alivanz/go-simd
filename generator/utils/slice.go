package utils

func Transform[A, B any](arr []A, fn func(i int, e A) B) []B {
	if arr == nil {
		return nil
	}
	out := make([]B, len(arr))
	for i, e := range arr {
		out[i] = fn(i, e)
	}
	return out
}

func Merge[T any](lists ...[]T) []T {
	var out []T
	for _, l := range lists {
		out = append(out, l...)
	}
	return out
}
