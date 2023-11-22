func pathSum(root *TreeNode, targetSum int) int {
	return dfs(root, targetSum, []int{})
}

func dfs(root *TreeNode, target int, currPath []int) int {
	if root == nil {
		return 0
	}
	newPath := make([]int, len(currPath))
	copy(newPath, currPath)
	newPath = append(newPath, root.Val)
	return count(newPath, target) +
		dfs(root.Left, target, newPath) +
		dfs(root.Right, target, newPath)
}

func count(path []int, target int) int {
	var cnt, sum int
	for i := len(path) - 1; i >= 0; i-- {
		sum += path[i]
		if sum == target {
			cnt++
		}
	}
	return cnt
}
