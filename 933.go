type RecentCounter struct {
	q []int
}

func Constructor() RecentCounter {
	return RecentCounter{
		q: make([]int, 0),
	}
}

func (rc *RecentCounter) Ping(t int) int {
	rc.q = append(rc.q, t)
	for rc.q[0] < t-3000 {
		rc.q = rc.q[1:]
	}
	return len(rc.q)
}
