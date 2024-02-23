func removeTrailingZeros(num string) string {
	l, r := 0, len(num)-1
	for l <= r {
		if num[l] != '0' && num[r] != '0' {
			break
		}
		if num[l] == '0' {
			l++
		}
		if num[r] == '0' {
			r--
		}
	}
	return num[l : r+1]
}
