// 1
func maxOperations(nums []int, k int) int {
	m := map[int]int{}
	for _, x := range nums {
		v, ok := m[x]
		if !ok {
			m[x] = 1
		}
		m[x] = v + 1
	}
	var res int
	for _, x := range nums {
		sub := k - x
		if sub < 0 {
			continue
		}
		if v, ok := m[sub]; ok {
			if (x == sub && v > 1) || (x != sub && v > 0 && m[x] > 0) {
				m[sub] = v - 1
				m[x] = m[x] - 1
				res++
			}
		}
	}
	return res
}

// 2
func maxOperations(nums []int, k int) int {
	m := map[int]int{}
	for _, x := range nums {
		v, ok := m[x]
		if !ok {
			m[x] = 1
		}
		m[x] = v + 1
	}
	var res int
	for k1, v1 := range m {
		k2 := k - k1
		if v2, ok := m[k2]; ok {
			if k1 == k2 {
				res += v1 / 2
				m[k1] = 0
				continue
			}
			min := int(math.Min(float64(v1), float64(v2)))
			res += min
			m[k1] = v1 - min
			m[k2] = v2 - min
		}
	}
	return res
}
