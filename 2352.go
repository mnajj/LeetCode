const size = 200

func equalPairs(grid [][]int) int {
	n := len(grid)
	m := make(map[[size]int]int)
	arr := [size]int{}
	for i := 0; i < n; i++ {
		copy(arr[:], grid[i])
		m[arr]++
	}
	var res int
	for i := 0; i < n; i++ {
		arr := [size]int{}
		for j := 0; j < n; j++ {
			arr[j] = grid[j][i]
		}
		if v, ok := m[arr]; ok {
			res += v
		}
	}
	return res
}
