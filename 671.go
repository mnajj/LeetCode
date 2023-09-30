func findSecondMinimumValue(root *TreeNode) int {
	return doWork(root, root.Val, -1)
}

func doWork(node *TreeNode, min, second int) int {
	if node.Left == nil {
		if node.Val != min && (min < second || second == -1) {
			return node.Val
		}
		return second
	}
	l, r := doWork(node.Left, min, second), doWork(node.Right, min, second)
	if l == -1 {
		return r
	}
	if r == -1 || l < r {
		return l
	}
	return r
}
