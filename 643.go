func findMaxAverage(nums []int, k int) float64 {
	if len(nums) == 1 {
		return float64(nums[0])
	}
	var sum int
	for i := 0; i < k; i++ {
		sum += nums[i]
	}
	maxAvg := float64(sum) / float64(k)
	for i := k; i < len(nums); i++ {
		sum -= nums[i-k]
		sum += nums[i]
		avg := float64(sum) / float64(k)
		if maxAvg < avg {
			maxAvg = avg
		}
	}
	return maxAvg
}
