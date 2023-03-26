func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	set := map[int]struct{}{}
	for _, x := range nums {
		set[x] = struct{}{}
	}
	var seqs []int
	for _, x := range nums {
		if _, ok := set[x-1]; !ok {
			count := 1
			for {
				x++
				if _, ok := set[x]; ok {
					count++
					continue
				}
				seqs = append(seqs, count)
				break
			}
		}
	}
	max := seqs[0]
	for _, x := range seqs {
		if max < x {
			max = x
		}
	}
	return max
}
