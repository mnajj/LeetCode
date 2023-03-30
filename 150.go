func evalRPN(tokens []string) int {
	var num []int
	for _, t := range tokens {
		if n, ok := isNumber(t); ok {
			num = append(num, n)
			continue
		}
		y, x := num[len(num)-1], num[len(num)-2]
		num = num[:len(num)-2]
		res := compute(t, x, y)
		num = append(num, res)
	}
	return num[0]
}

func isNumber(str string) (int, bool) {
	x, err := strconv.Atoi(str)
	return x, err == nil
}

func compute(op string, x, y int) int {
	switch op {
	case "+":
		return x + y
	case "-":
		return x - y
	case "*":
		return x * y
	case "/":
		return x / y
	default:
		panic("")
	}
}
