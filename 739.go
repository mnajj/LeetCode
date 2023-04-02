func dailyTemperatures(temperatures []int) []int {
	type element struct {
		value, index int
	}
	var monoStack []element
	res := make([]int, len(temperatures))
	for i, x := range temperatures {
		if len(monoStack) == 0 {
			monoStack = append(monoStack, element{x, i})
			continue
		}
		for len(monoStack) > 0 {
			if last := monoStack[len(monoStack)-1]; x > last.value {
				res[last.index] = i - last.index
				monoStack = monoStack[:len(monoStack)-1]
				continue
			}
			break
		}
		monoStack = append(monoStack, element{x, i})
	}
	return res
}
