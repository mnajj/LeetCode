func longestOnes(nums []int, k int) int {
	z, mx, l, r := 0, 0, 0, 0
	for r < len(nums) {
		if nums[r] == 1 {
			r++
		} else if z < k {
			z++
			r++
		} else {
			for nums[l] != 0 {
				l++
			}
			l++
			z--
		}
		mx = max(mx, r-l)
	}
	return mx
}
