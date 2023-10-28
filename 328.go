func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	fEv := head.Next
	odd, even := head, head.Next
	for even != nil && even.Next != nil {
		odd.Next = even.Next
		odd = odd.Next
		even.Next = odd.Next
		even = even.Next
	}
	odd.Next = fEv
	return head
}
