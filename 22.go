func generateParenthesis(n int) []string {
	var stack, res []string
	var backTrack func(int, int)
	backTrack = func(open, close int) {
		if open == close && open == n {
			res = append(res, strings.Join(stack, ""))
			return
		}
		if open < n {
			stack = append(stack, "(")
			backTrack(open+1, close)
			stack = stack[:len(stack)-1]
		}
		if close < open {
			stack = append(stack, ")")
			backTrack(open, close+1)
			stack = stack[:len(stack)-1]
		}
	}
	backTrack(0, 0)
	return res
}
