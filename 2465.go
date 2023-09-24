func distinctAverages(nums []int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	m := map[float64]struct{}{}
	for i := 0; i < len(nums)/2; i++ {
		avg := float64(nums[i]+nums[len(nums)-i-1]) / float64(2)
		if _, ok := m[avg]; !ok {
			m[avg] = struct{}{}
		}
	}
	return len(m)
}
