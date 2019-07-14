package array_sort

import (
	"math"
)

func trap(height []int) int {
	max := math.MinInt64
	ans := 0
	length := len(height)
	for _, v := range height {
		if v > max {
			max = v
		}
	}
	for i := 1 ; i <= max; i++ {
		for j := 1; j < length-1; j++ {
			if height[j] == 0 && height[j-1] != 0{
				cur := 0
				for j < length {
					if height[j] == 0 {
						cur++
						j++
					}else {
						break
					}
				}
				if j != length{
					ans += cur
				}
			}
		}
		for j := 0; j < length; j++ {
			if height[j] > 0 {
				height[j]--
			}
		}
	}
	return ans
}
func main()  {

}