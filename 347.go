func topKFrequent(nums []int, k int) []int {
	freq := map[int]int{}
	for _, x := range nums {
		if _, ok := freq[x]; ok {
			freq[x]++
			continue
		}
		freq[x] = 1
	}
	type elem struct {
		num, freq, idx int
	}
	flat := make([]elem, 0, len(freq))
	for k, v := range freq {
		flat = append(flat, elem{
			num:  k,
			freq: v,
		})
	}
	sort.Slice(flat, func(i, j int) bool {
		return flat[i].freq > flat[j].freq
	})
	ret := make([]int, 0, k)
	for i := 0; i < k; i++ {
		ret = append(ret, flat[i].num)
	}
	return ret
}
