func canPlaceFlowers(flowerbed []int, n int) bool {
	if n == 0 {
		return true
	}
	flowerbed = append([]int{0}, flowerbed...)
	flowerbed = append(flowerbed, 0)
	for i := 1; i < len(flowerbed)-1; i++ {
		if flowerbed[i] == 0 && flowerbed[i-1] == 0 && flowerbed[i+1] == 0 {
			flowerbed[i] = 1
			n--
			if n == 0 {
				return true
			}
		}
	}
	return false
}
