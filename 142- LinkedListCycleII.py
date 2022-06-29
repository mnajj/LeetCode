class Solution:
    def detectCycle(self, head: Optional[ListNode]) -> Optional[ListNode]:
        if not head or not head.next:
            return None
        tortise = hare = head
        while hare and hare.next:
            tortise = tortise.next
            hare = hare.next.next
            if hare == tortise:
                break
        if hare != tortise:
            return None
        hare = head
        while hare != tortise:
            hare = hare.next
            tortise = tortise.next
        return hare
