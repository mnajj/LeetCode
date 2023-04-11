func majorityElement(nums []int) int {
	ocr := make(map[int]int)
	h := len(nums) / 2
	for _, x := range nums {
		if _, ok := ocr[x]; !ok {
			ocr[x] = 0
		}
		ocr[x]++
		if ocr[x] > h {
			return x
		}
	}
  panic("")
}
