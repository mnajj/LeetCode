class Solution:
    def middleNode(self, head: ListNode) -> ListNode:
        len = 0
        h = head
        while h:
            len += 1
            h = h.next
        mid = len // 2
        i = 0
        h = head
        while i != mid:
            i += 1
            h = h.next
        return h


class Solution:
    def middleNode(self, head: ListNode) -> ListNode:
        totries = hare = head
        while hare.next:
            totries = totries.next
            hare = hare.next.next
        return totries
