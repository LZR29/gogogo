package contest

//5653. 可以形成最大正方形的矩形数目
func countGoodRectangles(rectangles [][]int) int {
	maxLen := 0
	min := 10 ^ 9
	sideMap := make(map[int]int)
	ret := 0
	for _, rectangle := range rectangles {
		min = rectangle[0]
		if min > rectangle[1] {
			min = rectangle[1]
		}
		if min > maxLen {
			maxLen = min
		}
		sideMap[min]++
	}
	for k, v := range sideMap {
		if k == maxLen {
			ret = v
			break
		}
	}
	return ret
}

//5243. 同积元组 一个组合就8种
func tupleSameProduct(nums []int) int {
	tempMap := make(map[int]int)
	ret := 0
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			tempMap[nums[i]*nums[j]]++
		}
	}
	for _, val := range tempMap {
		if val >= 2 {
			ret += (val - 1) * val / 2 * 8
		}
	}
	return ret
}
