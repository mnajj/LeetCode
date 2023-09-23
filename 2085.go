func countWords(words1 []string, words2 []string) int {
	m := map[string]struct{ one, two int }{}
	for _, x := range words1 {
		v, ok := m[x]
		if !ok {
			v.one = 1
			m[x] = v
			continue
		}
		v.one = v.one + 1
		m[x] = v
	}
	for _, x := range words2 {
		v, ok := m[x]
		if !ok {
			v.two = 1
			m[x] = v
			continue
		}
		v.two = v.two + 1
		m[x] = v
	}
	var res int
	for _, v := range m {
		if v.one == 1 && v.two == 1 {
			res++
		}
	}
	return res
}
