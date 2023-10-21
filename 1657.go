func closeStrings(w1 string, w2 string) bool {
	if len(w1) != len(w2) {
		return false
	}
	m1, m2 := map[byte]int{}, map[byte]int{}
	for i := 0; i < len(w1); i++ {
		m1[w1[i]]++
		m2[w2[i]]++
	}
	c1, c2 := []int{}, []int{}
	for k1, v1 := range m1 {
		v2, ok := m2[k1]
		if !ok {
			return false
		}
		c1 = append(c1, v1)
		c2 = append(c2, v2)
	}
	sort.Ints(c1)
	sort.Ints(c2)
	return reflect.DeepEqual(c1, c2)
}
