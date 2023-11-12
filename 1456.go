func maxVowels(s string, k int) int {
	lk := map[rune]struct{}{
		'a': {},
		'e': {},
		'i': {},
		'o': {},
		'u': {},
	}
	runes := []rune(s)
	var mx, curr int
	for i := 0; i < k; i++ {
		if _, ok := lk[runes[i]]; ok {
			curr++
		}
	}
	mx = curr

	for l, r := 0, k; r < len(runes); l, r = l+1, r+1 {
		if _, ok := lk[runes[l]]; ok {
			curr--
		}
		if _, ok := lk[runes[r]]; ok {
			curr++
		}
		mx = max(mx, curr)
	}
	return mx
}
