func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val < key {
		root.Right = deleteNode(root.Right, key)
	} else if root.Val > key {
		root.Left = deleteNode(root.Left, key)
	} else {
		if root.Left == nil && root.Right == nil {
			return nil
		} else if root.Left == nil {
			return root.Right
		} else if root.Right == nil {
			return root.Left
		} else {
			root.Val = findMin(root.Right)
			root.Right = deleteNode(root.Right, root.Val)
		}
	}
	return root
}

func findMin(root *TreeNode) int {
	min := root.Val
	for root != nil {
		if root.Val < min {
			min = root.Val
		}
		root = root.Left
	}
	return min
}
