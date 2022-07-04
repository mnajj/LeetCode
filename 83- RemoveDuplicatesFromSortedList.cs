public class Solution
{
	public ListNode DeleteDuplicates(ListNode head)
	{
		var curr = head;
		while (curr is not null && curr.next is not null)
		{
			if (curr.val == curr.next.val)
			{
				curr.next = curr.next.next;
			}
			else
			{
				curr = curr.next;
			}
		}

		return head;
	}
}
