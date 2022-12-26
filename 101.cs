public class Solution
{
	public bool IsSymmetric(TreeNode root)
	{
		return Check(root.left, root.right);
	} 

	private bool Check(TreeNode left, TreeNode right)
	{
		if (left is null && right is null) return true;
		if(left == null || right == null || left.val != right.val) return false;
		return (Check(left.left,right.right) && Check(left.right,right.left));
	}
}
