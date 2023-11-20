
public class Solution
{
	public List<int> Result { get; set; } = new List<int>();

	public IList<int> PostorderTraversal(TreeNode root)
	{
		if (root is not null)
		{
			PostorderTraversal(root.left);
			PostorderTraversal(root.right);
			Result.Add(root.val);
		}
		return Result;
	}
}
