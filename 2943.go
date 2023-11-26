func findWordsContaining(words []string, x byte) []int {
	var res []int
	r := rune(x)
	for i, w := range words {
		if strings.ContainsRune(w, r) {
			res = append(res, i)
		}
	}
	return res
}
