func maximum69Number(num int) int {
	s := []rune(strconv.Itoa(num))
	for i, r := range s {
		if r == '6' {
			s[i] = '9'
			break
		}
	}
	n, err := strconv.Atoi(string(s))
	if err != nil {
		panic(err)
	}
	return n
}
