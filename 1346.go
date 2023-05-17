func checkIfExist(arr []int) bool {
	m := map[int]int{}
	for i, x := range arr {
		x *= 2
		m[x]=i
	}
	for i, x := range arr {
		if v, ok := m[x]; ok && v != i {
			return true
		}
	}
	return false
}
