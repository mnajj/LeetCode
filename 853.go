type car struct {
	pos, speed int
}

func carFleet(target int, position []int, speed []int) int {
	cars := make([]car, 0, len(position))
	var fleets []car
	for i := 0; i < len(position); i++ {
		cars = append(cars, car{position[i], speed[i]})
	}
	sort.Slice(cars, func(i, j int) bool {
		return cars[i].pos < cars[j].pos
	})
	for i := len(cars) - 1; i >= 0; i-- {
		if len(fleets) == 0 || fleets[len(fleets)-1].pos == target {
			fleets = append(fleets, cars[i])
			continue
		}
		curr, next := arriveAfter(cars[i], target), arriveAfter(fleets[len(fleets)-1], target)
		if curr <= next {
			continue
		}
		fleets = append(fleets, cars[i])
	}
	return len(fleets)
}

func arriveAfter(c car, t int) float64 {
	return float64(t-c.pos) / float64(c.speed)
}
