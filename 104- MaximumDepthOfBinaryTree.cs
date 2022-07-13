public class Solution
{
	public int MaxDepth(TreeNode root)
	{
		if (root is null) return 0;
		var l = MaxDepth(root.left);
		var r = MaxDepth(root.right);
		return Math.Max(l, r) + 1;
	}
}
