# straightforward
class Solution(object):
    def addTwoNumbers(self, l1, l2):
        flag = True
        total_list = head = ListNode()
        remainder = 0
        sum = 0
        while flag:
            if l1 is not None:
                sum += l1.val
                l1 = l1.next
            if l2 is not None:
                sum += l2.val
                l2 = l2.next
            sum += remainder
            remainder = 0
            if sum > 9:
                frac, whole = math.modf(sum/10)
                remainder = int(whole)
                total_list.val = int(round(frac, 1) * 10)
            else:
                total_list.val = sum
            sum = 0
            if l1 is None and l2 is None:
                if remainder != 0:
                    total_list.next = ListNode()
                    total_list = total_list.next
                    total_list.val = remainder
                    remainder = 0
                flag = False
                total_list.next = None
            else:
                total_list.next = ListNode()
                total_list = total_list.next
        return head
