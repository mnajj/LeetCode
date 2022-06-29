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

    
# in-place
class Solution:
    def removeElements(self, head: ListNode, val: int) -> ListNode:
        prev = None
        i = 0
        curr = head
        while curr:
            if curr.val == val and i == 0:
                head = curr = curr.next
                continue
            else:
                i += 1
            if curr.val == val:
                prev.next = curr.next
                curr = prev.next
                continue
            prev = curr
            curr = curr.next
        return head
