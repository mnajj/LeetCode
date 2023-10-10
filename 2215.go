func findDifference(n1 []int, n2 []int) [][]int {
	m1, m2 := map[int]struct{}{}, map[int]struct{}{}
	for _, v := range n1 {
		if _, ok := m1[v]; !ok {
			m1[v] = struct{}{}
		}
	}
	for _, v := range n2 {
		if _, ok := m2[v]; !ok {
			m2[v] = struct{}{}
		}
	}
	res1, res2 := map[int]struct{}{}, map[int]struct{}{}
	for _, v := range n1 {
		if _, ok := m2[v]; !ok {
			res1[v] = struct{}{}
		}
	}
	for _, v := range n2 {
		if _, ok := m1[v]; !ok {
			res2[v] = struct{}{}
		}
	}
	return [][]int{mtoa(res1), mtoa(res2)}
}

func mtoa(m map[int]struct{}) []int {
	a := make([]int, 0, len(m))
	for k := range m {
		a = append(a, k)
	}
	return a
}
