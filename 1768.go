func mergeAlternately(word1 string, word2 string) string {
	var buf bytes.Buffer
	i := 0
	for {
		if i >= len(word1) && i >= len(word2) {
			break
		}
		if i < len(word1) {
			buf.WriteRune(rune(word1[i]))
		}
		if i < len(word2) {
			buf.WriteRune(rune(word2[i]))
		}
		i++
	}
	return buf.String()
}
