func largestGoodInteger(num string) string {
	i, j := 0, 1
	m := map[int]string{}
	for j < len(num) {
		if num[i] != num[j] {
			i = j
		}
		j++
		if j-i == 3 {
			m[int(num[i])] = num[i:j]
		}
	}
	if len(m) == 0 {
		return ""
	}
	var maxKey int
	for k := range m {
		if k > maxKey {
			maxKey = k
		}
	}
	return m[maxKey]
}
