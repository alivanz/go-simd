package utils

func Filter[T any](arr []T, fn func(e T) bool) []T {
	out := make([]T, 0, len(arr))
	for _, e := range arr {
		if !fn(e) {
			continue
		}
		out = append(out, e)
	}
	return out
}
