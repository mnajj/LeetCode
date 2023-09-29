func splitNum(num int) int {
	digits := make([]int, 0)
	for num != 0 {
		digits = append(digits, num%10)
		num /= 10
	}
	sort.Slice(digits, func(i, j int) bool { return digits[i] < digits[j] })
	var n1, n2 int
	for i, swap := 0, true; i < len(digits); i, swap = i+1, !swap {
		if swap {
			n1 = n1*10 + digits[i]
		} else {
			n2 = n2*10 + digits[i]
		}
	}
	return n1 + n2
}
