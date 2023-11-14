func isBalanced(r *TreeNode) bool {
	if r == nil {
		return true
	}
	diff := dep(r.Left) - dep(r.Right)
	if diff <= 1 && diff >= -1 {
		return isBalanced(r.Left) && isBalanced(r.Right)
	}
	return false
}

func dep(r *TreeNode) int {
	if r == nil {
		return 0
	}
	return max(dep(r.Left), dep(r.Right)) + 1
}
