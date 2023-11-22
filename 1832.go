func checkIfPangram(sentence string) bool {
	m := map[rune]struct{}{}
	for _, r := range sentence {
		m[r] = struct{}{}
	}
	return len(m) == 26
}
