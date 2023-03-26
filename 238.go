func productExceptSelf(nums []int) []int {
	n := len(nums)
	ret := make([]int, len(nums))
	pre := 1
	for i := 0; i < n; i++ {
		ret[i] = pre
		pre *= nums[i]
	}
	post := 1
	for i := n - 1; i >= 0; i-- {
		ret[i] *= post
		post *= nums[i]
	}
	return ret
}
