
public class Solution
{
	public List<int> Result { get; set; } = new List<int>();

	public IList<int> InorderTraversal(TreeNode root)
	{
		if (root is not null)
		{
			InorderTraversal(root.left);
			Result.Add(root.val);
			InorderTraversal(root.right);
		}
		return Result;
	}
}
