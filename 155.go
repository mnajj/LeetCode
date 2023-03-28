type MinStack struct {
	topIdx     int
	stack, min []int
}

func Constructor() MinStack {
	return MinStack{topIdx: -1, stack: []int{}, min: []int{}}
}

func (s *MinStack) Push(val int) {
	s.stack = append(s.stack, val)
	if len(s.min) == 0 {
		s.min = append(s.min, val)
		return
	}
	minSoFar := s.min[len(s.min)-1]
	if minSoFar > val {
		minSoFar = val
	}
	s.min = append(s.min, minSoFar)
}

func (s *MinStack) Pop() {
	if len(s.stack) == 0 {
		return
	}
	s.stack = s.stack[:len(s.stack)-1]
	s.min = s.min[:len(s.min)-1]
}

func (s *MinStack) Top() int {
	if len(s.stack) == 0 {
		panic("")
	}
	return s.stack[len(s.stack)-1]
}

func (s *MinStack) GetMin() int {
	if len(s.min) == 0 {
		panic("")
	}
	return s.min[len(s.min)-1]
}
