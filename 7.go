func reverse(x int) int {
	const MaxInt32 = 1<<31 - 1
	const MinInt32 = -1 << 31
	var result int
	for x != 0 {
		digit := x % 10
		x /= 10
		if result > MaxInt32/10 || (result == MaxInt32/10 && digit > 7) {
			return 0
		}
		if result < MinInt32/10 || (result == MinInt32/10 && digit < -8) {
			return 0
		}
		result = result*10 + digit
	}
	return result
}
