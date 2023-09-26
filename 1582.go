func numSpecial(mat [][]int) int {
	pos := make([]struct {
		x, y int
	}, 0)
	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[i]); j++ {
			if mat[i][j] == 1 {
				pos = append(pos, struct {
					x, y int
				}{i, j})
			}
		}
	}

	for i := 0; i < len(pos); i++ {
		if pos[i].x == -1 {
			continue
		}
		isX, isY := false, false
		for j := 0; j < len(pos); j++ {
			if i == j {
				continue
			}
			if pos[i].x == pos[j].x {
				pos[j].x = -1
				isX = true
			}
			if pos[i].y == pos[j].y {
				pos[j].y = -1
				isY = true
			}
		}
		if isX {
			pos[i].x = -1
		}
		if isY {
			pos[i].y = -1
		}
	}
	var res int
	for _, p := range pos {
		if p.x != -1 && p.y != -1 {
			res++
		}
	}
	return res
}
