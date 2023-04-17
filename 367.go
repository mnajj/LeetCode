func isPerfectSquare(num int) bool {
	l, r := 1, num
	for l <= r {
		mid := (l + r) / 2
		sqr := mid * mid
		if sqr == num {
			return true
		}
		if sqr < num {
			l = mid + 1
		}
		if sqr > num {
			r = mid - 1
		}
	}
	return false
}
