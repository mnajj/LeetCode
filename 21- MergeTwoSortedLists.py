class Solution:
    def mergeTwoLists(self, l1: ListNode, l2: ListNode) -> ListNode:
        head = ListNode()
        t = head
        while l1 and l2:
            if l1.val < l2.val:
                t.next = ListNode(l1.val)
                l1 = l1.next
            elif l1.val >= l2.val:
                t.next = ListNode(l2.val)
                l2 = l2.next
            t = t.next
        if l1 is not None:
            t.next = l1
        if l2 is not None:
            t.next = l2
        return head.next
