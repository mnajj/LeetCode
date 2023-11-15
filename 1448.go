func goodNodes(r *TreeNode) int {
	if r == nil {
		return 0
	}
	return trav(r.Left, r.Val) + trav(r.Right, r.Val) + 1
}

func trav(r *TreeNode, mx int) int {
	var cnt int
	if r == nil {
		return cnt
	}
	mx = max(mx, r.Val)
	if r.Val >= mx {
		cnt++
	}
	return trav(r.Left, mx) + trav(r.Right, mx) + cnt
}
