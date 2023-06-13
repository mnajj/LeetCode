public class Solution
{
  public bool IsCompleteTree(TreeNode root)
  {
    var queue = new Queue<TreeNode>();
    queue.Enqueue(root);
    var isNoMoreAllowed = false;
    while (queue.Any())
    {
      var size = queue.Count;
      while (size > 0)
      {
        var curr = queue.Dequeue();
        if (curr.left is not null)
        {
          if (isNoMoreAllowed) return false;
          queue.Enqueue(curr.left);
        }
        else
        {
          isNoMoreAllowed = true;
        }
        if (curr.right is not null)
        {
          if (isNoMoreAllowed) return false;
          queue.Enqueue(curr.right);
        }
        else
        {
          isNoMoreAllowed = true;
        }
        size--;
      }
    }
    return true;
  }
}
