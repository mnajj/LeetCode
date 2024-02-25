func minimumOperations(nums []int) int {
	var res int
	m := map[int]struct{}{}
	for _, x := range nums {
		if _, ok := m[x]; x != 0 && !ok {
			m[x] = struct{}{}
			res++
		}
	}
	return res
}
