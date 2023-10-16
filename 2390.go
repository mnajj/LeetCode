func removeStars(s string) string {
	stk := make([]rune, 0, len(s))
	for _, c := range s {
		if c == '*' {
			stk = stk[:len(stk)-1]
		} else {
			stk = append(stk, c)
		}
	}
	return string(stk)
}
