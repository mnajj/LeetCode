func maxCount(m int, n int, ops [][]int) int {
	if len(ops) == 0 {
		return m * n
	}
	minRow, minCol := math.MaxInt, math.MaxInt
	for _, op := range ops {
		if op[0] < minRow {
			minRow = op[0]
		}
		if op[1] < minCol {
			minCol = op[1]
		}
	}
	return minRow * minCol
}
