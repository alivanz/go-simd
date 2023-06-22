package scanner

func transform[A, B any](arr []A, fn func(i int, e A) B) []B {
	out := make([]B, len(arr))
	for i, e := range arr {
		out[i] = fn(i, e)
	}
	return out
}

func joinAll[T any](lists ...[]T) []T {
	var out []T
	for _, l := range lists {
		out = append(out, l...)
	}
	return out
}

func commaSplit(ss ...string) []string {
	switch len(ss) {
	case 0:
		return nil
	case 1:
		s := regWhitespace.ReplaceAllString(ss[0], " ")
		if len(s) == 0 {
			return nil
		}
		return regComma.Split(s, -1)
	default:
		return append(commaSplit(ss[0]), commaSplit(ss[1:]...)...)
	}
}
