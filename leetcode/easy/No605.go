package easy

//看当前这块每种花的土地的前后是否都没种花
func canPlaceFlowers(flowerbed []int, n int) bool {
	length := len(flowerbed)
	pre := 0
	next := 0
	for i := 0; i < length; i++ {
		if flowerbed[i] != 1 {
			if i == length-1 {
				next = 0
			} else {
				next = flowerbed[i+1]
			}
			if i == 0 {
				pre = 0
			} else {
				pre = flowerbed[i-1]
			}
			if next == 0 && pre == 0 {
				n--
				flowerbed[i] = 1
			}
		}
	}
	return n <= 0
}
