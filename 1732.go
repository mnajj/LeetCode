func largestAltitude(gain []int) int {
	curr, max := 0, 0
	for _, x := range gain {
		curr += x
		if max < curr {
			max = curr
		}
	}
	return max
}
