func deleteMiddle(head *ListNode) *ListNode {
	if head.Next == nil {
		return nil
	}
	s, f := head, head.Next
	for f.Next != nil && f.Next.Next != nil {
		f = f.Next.Next
		s = s.Next
	}
	s.Next = s.Next.Next
	return head
}
