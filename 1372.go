func longestZigZag(root *TreeNode) int {
	return max(
		maxZigZag(root.Left, true, 0),
		maxZigZag(root.Right, false, 0))
}

func maxZigZag(root *TreeNode, left bool, len int) int {
	if root == nil {
		return len
	}
	if left {
		return max(
			maxZigZag(root.Left, true, 0),
			maxZigZag(root.Right, false, len+1))
	}
	return max(
		maxZigZag(root.Left, true, len+1),
		maxZigZag(root.Right, false, 0))
}
