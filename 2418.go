import "sort"

func sortPeople(names []string, heights []int) []string {
	m := map[int]string{}
	for i := 0; i < len(names); i++ {
		m[heights[i]] = names[i]
	}

	sort.Slice(heights, func(i, j int) bool {
		return heights[i] > heights[j]
	})

	output := make([]string, len(names))
	for i, h := range heights {
		output[i] = m[h]
	}

	return output
}
