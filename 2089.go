func targetIndices(nums []int, target int) []int {
	var result []int
	sort.Ints(nums)
	for i, x := range nums {
		if x == target {
			result = append(result, i)
		}
	}
	return result
}
