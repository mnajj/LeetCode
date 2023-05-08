// 1
func countCharacters(words []string, chars string) int {
	var sum int
	m := map[rune]int{}
	for _, r := range chars {
		if _, ok := m[r]; !ok {
			m[r] = 0
		}
		m[r]++
	}
	Main:
	for _, w := range words {
		c := mapCopy(m)
		for _, r := range w {
			if v, ok := c[r]; !ok || v == 0{
				continue Main
			}
			c[r]--
		}
		sum += len(w)
	}
	return sum
}

func mapCopy(m map[rune]int) map[rune]int {
	c := map[rune]int{}
	for k, v := range m {
		c[k] = v;
	}
	return c
}

// 2
func countCharacters(words []string, chars string) int {
	var sum int
	m := map[rune]int{}
	for _, r := range chars {
		if _, ok := m[r]; !ok {
			m[r] = 0
		}
		m[r]++
	}
	Main:
	for _, w := range words {
		o := map[rune]int{}
		for _, r := range w {
			if _, ok := o[r]; !ok {
				o[r] = 0
			}
			o[r]++
		}
		for k,v := range o {
			if count, ok := m[k]; !ok || v > count {
				continue Main
			}
		}
		sum += len(w)
	}
	return sum
}
