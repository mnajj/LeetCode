class Solution:
    def removeNthFromEnd(self, head: ListNode, n: int) -> ListNode:
        slow, fast = head, head
        for i in range(n):
            fast = fast.next
        if not fast:
            return head.next
        while fast.next:
            fast = fast.next
            slow = slow.next
        slow.next = slow.next.next
        return head
