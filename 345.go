// 1

func reverseVowels(s string) string {
	rs := []rune(s)
	lk := map[rune]struct{}{
		'a': {},
		'A': {},
		'e': {},
		'E': {},
		'i': {},
		'I': {},
		'o': {},
		'O': {},
		'u': {},
		'U': {},
	}
	l, r := 0, len(rs)-1
	for l < r {
		_, okl := lk[rs[l]]
		_, okr := lk[rs[r]]
		if !okl {
			l++
		}
		if !okr {
			r--
		}
		if okl && okr {
			rs[l], rs[r] = rs[r], rs[l]
			l++
			r--
		}
	}
	return string(rs)
}

// 2
func reverseVowels(s string) string {
	rs := []rune(s)
	lk := map[rune]struct{}{
		'a': {},
		'A': {},
		'e': {},
		'E': {},
		'i': {},
		'I': {},
		'o': {},
		'O': {},
		'u': {},
		'U': {},
	}
	l, r := 0, len(rs)-1
	for l < r {
		for _, okl := lk[rs[l]]; !okl && l < r; _, okl = lk[rs[l]] {
			l++
		}
		for _, okr := lk[rs[r]]; !okr && l < r; _, okr = lk[rs[r]] {
			r--
		}
		rs[l], rs[r] = rs[r], rs[l]
		l++
		r--
	}
	return string(rs)
}
