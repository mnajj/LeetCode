func pivotIndex(nums []int) int {
	n := len(nums)
	l, r := make([]int, 0, n), make([]int, 0, n)
	l = append(l, 0)
	r = append(r, 0)
	var sum int
	for i := 1; i < n; i++ {
		sum += nums[i-1]
		l = append(l, sum)
	}
	sum = 0
	for i := n - 2; i >= 0; i-- {
		sum += nums[i+1]
		r = append(r, sum)
	}
	for i := 0; i < n; i++ {
		if l[i] == r[n-1-i] {
			return i
		}
	}
	return -1
}
