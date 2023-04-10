func strStr(haystack string, needle string) int {
	haystack += "#"
	var curr int
	for i := 0; i < len(haystack); i++ {
		curr = 0
		for j := i; j < len(haystack); j++ {
			if curr == len(needle) {
				return i
			}
			if haystack[j] != needle[curr] {
				break
			}
			curr++
		}
	}
	return -1
}
