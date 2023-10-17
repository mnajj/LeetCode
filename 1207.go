func uniqueOccurrences(arr []int) bool {
	im := map[int]int{}
	for _, x := range arr {
		if v, ok := im[x]; ok {
			im[x] = v + 1
		} else {
			im[x] = 1
		}
	}
	ocr := map[int]struct{}{}
	for _, v := range im {
		if _, ok := ocr[v]; ok {
			return false
		} else {
			ocr[v] = struct{}{}
		}
	}
	return true
}
