func compress(chars []byte) int {
	oc, nc := 0, 0
	for oc < len(chars) {
		var i int
		for i = oc + 1; i < len(chars) && chars[i] == chars[oc]; i++ {
		}
		chars[nc] = chars[oc]
		nc++
		if cnt := i - oc; cnt > 1 {
			digits := []byte(strconv.FormatInt(int64(cnt), 10))
			for j := 0; j < len(digits); j++ {
				chars[nc+j] = digits[j]
			}
			nc += len(digits)
		}
		oc = i
	}
	chars = chars[:nc]
	return nc
}
