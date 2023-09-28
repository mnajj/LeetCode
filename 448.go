func findDisappearedNumbers(nums []int) []int {
	m := map[int]struct{}{}
	for _, v := range nums {
		if _, ok := m[v]; !ok {
			m[v] = struct{}{}
		}
	}
	res := make([]int, 0)
	for i := 1; i <= len(nums); i++ {
		if _, ok := m[i]; !ok {
			res = append(res, i)
		}
	}
	return res
}
