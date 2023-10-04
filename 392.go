func isSubsequence(s string, t string) bool {
	p1 := 0
	for p2 := 0; p1 < len(s) && p2 < len(t); p2++ {
		if s[p1] == t[p2] {
			p1++
		}
	}
	return p1 == len(s)
}
