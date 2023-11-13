func leafSimilar(r1 *TreeNode, r2 *TreeNode) bool {
	l1, l2 := getLeafs(r1), getLeafs(r2)
	if len(l1) != len(l2) {
		return false
	}
	for i, x := range l1 {
		if x != l2[i] {
			return false
		}
	}
	return true
}

func getLeafs(r *TreeNode) []int {
	if r == nil {
		return []int{}
	}
	if r.Left == nil && r.Right == nil {
		return []int{r.Val}
	}
	return append(getLeafs(r.Left), getLeafs(r.Right)...)
}
