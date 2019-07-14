package array_sort

import (
	"sort"
)

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	var ans [][]int
	length := len(nums)
	for i := 0; i < length; i++ {
		if nums[i] > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left, right := i + 1, length - 1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum == 0 {
				cur := []int{nums[i], nums[left], nums[right]}
				ans = append(ans, cur)
				for left < right && nums[left] == nums[left+1]{
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				left++
				right--
			}else if sum > 0 {
				right--
			}else {
				left++
			}
		}
	}
	return ans
}

