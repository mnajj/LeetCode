public class Solution
{
	public ListNode ReverseList(ListNode head)
	{
		if (head is null)
		{
			return null;
		}
		var prev = head;
		var curr = head.next;
		prev.next = null;
		while (curr is not null)
		{
			var temp = curr.next;
			curr.next = prev;
			prev = curr;
			curr = temp;
		}

		return prev;
	}
}
