func reverseWords(s string) string {
	res := make([]rune, 0)
	rs := []rune(strings.TrimRight(s, " "))
	rs = append([]rune{' '}, rs...)
	for len(rs) > 0 {
		rs = []rune(strings.TrimRight(string(rs), " "))
		for i := len(rs) - 1; i >= 0; i-- {
			if rs[i] == ' ' {
				res = append(res, rs[i+1:]...)
				res = append(res, ' ')
				rs = rs[:i]
				break
			}
		}
	}
	return strings.TrimRight(string(res), " ")
}
