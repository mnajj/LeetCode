func minimumRecolors(blocks string, k int) int {
	var min int
	for i := 0; i < k; i++ {
		if blocks[i] == 'W' {
			min++
		}
	}
	curr := min
	for l, r := 1, k; r < len(blocks); l, r = l+1, r+1 {
		if blocks[l-1] == 'W' {
			curr--
		}
		if blocks[r] == 'W' {
			curr++
		}
		if curr < min {
			min = curr
		}
	}
	return min
}
