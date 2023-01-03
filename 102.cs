
public class Solution
{
	public IList<IList<int>> LevelOrder(TreeNode root)
	{
		Queue<TreeNode> queue = new Queue<TreeNode>();
		if (root != null)
		{
			queue.Enqueue(root);
		}
		IList<IList<int>> result = new List<IList<int>>();
		while (queue.Any())
		{
			List<int> temp = new List<int>();
			int n = queue.Count;

			for (int i = 0; i < n; i++)
			{
				var node = queue.Dequeue();

				if (node.left != null)
				{
					queue.Enqueue(node.left);
				}

				if (node.right != null)
				{
					queue.Enqueue(node.right);
				}

				temp.Add(node.val);
			}
			result.Add(temp);
		}
		return result;
	}
}
