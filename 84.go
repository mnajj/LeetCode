func largestRectangleArea(heights []int) int {
	var st stack
	var max int
	for i, h := range heights {
		start := i
		for !st.isEmpty() && st.top().h > h {
			elem := st.top()
			st = st.pop()
			if ar := elem.h * (i - elem.i); max < ar {
				max = ar
			}
			start = elem.i
		}
		st = st.push(element{h, start})
	}
	for _, elem := range st {
		if ar := elem.h * (len(heights) - elem.i); ar > max {
			max = ar
		}
	}
	return max
}

type element struct {
	h, i int
}
type stack []element

func (s stack) top() element {
	return s[len(s)-1]
}

func (s stack) pop() stack {
	return s[:len(s)-1]
}

func (s stack) push(e element) stack {
	return append(s, e)
}

func (s stack) isEmpty() bool {
	return len(s) == 0
}
