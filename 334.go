func increasingTriplet(nums []int) bool {
	f, s := math.MaxInt32, math.MaxInt32
	for _, n := range nums {
		if n <= f {
			f = n
		} else if n <= s {
			s = n
		} else {
			return true
		}
	}
	return false
}
