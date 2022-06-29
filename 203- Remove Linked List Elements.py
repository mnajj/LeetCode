class Solution:
    def removeElements(self, ls: ListNode, val: int) -> ListNode:
        head = ListNode()
        t = head
        while ls:
            if ls.val != val:
                t.next = ListNode(ls.val)
                t = t.next
            ls = ls.next
        return head.next
