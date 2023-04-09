func isPalindrome(head *ListNode) bool {
	var stack []int
	itr := head
	for itr != nil {
		stack = append(stack, itr.Val)
		itr = itr.Next
	}
	for head != nil {
		if stack[len(stack)-1] != head.Val {
			return false
		}
		stack = stack[:len(stack)-1]
		head = head.Next
	}
	return true
}
