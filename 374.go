func guessNumber(n int) int {
	l, r := 1, n
	for l <= r {
		m := (l+r)/2
		g := guess(m)
		if g == 0 {
			return m
		} else if g == 1 {
			l = m + 1
		} else {
			r = m - 1
		}
	}
	return n
}
