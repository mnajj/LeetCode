class Solution:
    def hasCycle(self, head: ListNode) -> bool:
        dict = {}
        while head:
            if head in dict:
                return True
            else:
                dict[head] = True
            head = head.next
        return False
      
      
class Solution:
    def hasCycle(self, head: ListNode) -> bool:
        if head is None:
            return False
        fst = head
        sec = head.next
        while fst != sec:
            if sec is None or sec.next is None:
                return False
            fst = fst.next
            sec = sec.next.next
        return True
