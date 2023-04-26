func arithmeticTriplets(nums []int, diff int) int {
	var sum int
	type triple struct {
		i, j, k int
	}
	touch := map[triple]struct{}{}
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			for k := j + 1; k < len(nums); k++ {
				f, s, t := nums[i], nums[j], nums[k]
				if s-f == diff && t-s == diff {
					if _, ok := touch[triple{f, s, t}]; !ok {
						sum++
					}
				}
			}
		}
	}
	return sum
}
