package scanner

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
