func summaryRanges(nums []int) []string {
	nums = append(nums, -1)
	var result []string
	var f, t = nums[0], nums[0]
	var buf bytes.Buffer
  
	for i := 0; i < len(nums)-1; i++ {
		if nums[i+1] != nums[i]+1 {
			t = nums[i]
			buf.WriteString(strconv.Itoa(f))
			if f != t {
				buf.WriteString(fmt.Sprintf("->%d", t))
			}
			result = append(result, buf.String())
			buf.Reset()
			f, t = nums[i+1], nums[i+1]
		}
	}
	return result
}
