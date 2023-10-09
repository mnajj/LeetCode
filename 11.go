func maxArea(h []int) int {
	l, r, res := 0, len(h)-1, 0
	for l < r {
		a := int((r - l) * int(math.Min(float64(h[r]), float64(h[l]))))
		if res < a {
			res = a
		}
		if h[l] < h[r] {
			l++
		} else {
			r--
		}
	}
	return res
}
