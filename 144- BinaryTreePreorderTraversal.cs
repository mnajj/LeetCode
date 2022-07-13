
public class Solution
{
	public List<int> Result { get; set; } = new List<int>();

	public IList<int> PreorderTraversal(TreeNode root)
	{
		if (root is not null)
		{
			Result.Add(root.val);
			PreorderTraversal(root.left);
			PreorderTraversal(root.right);
		}
		return Result;
	}
}
