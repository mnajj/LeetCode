func pairSum(head *ListNode) int {
	slc := make([]int, 0)
	for curr := head; curr != nil; curr = curr.Next {
		slc = append(slc, curr.Val)
	}
	max := slc[0]
	for i := 0; i < len(slc)/2+1; i++ {
		sum := slc[i] + slc[len(slc)-1-i]
		if sum > max {
			max = sum
		}
	}
	return max
}
