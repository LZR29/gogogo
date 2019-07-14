package array_sort

import "sort"

//排序解法
func longestConsecutive(nums []int) int {
	sort.Ints(nums)
	ans := 0
	length := len(nums)
	start := 0
	same := 0
	for i := 1; i < length; i++ {
		if nums[i] == nums[i-1] + 1 {
			continue
		}else if nums[i] == nums[i-1] {
			same++
		}else {
			cur := i - start - same
			if ans < cur {
				ans = cur
			}
			start = i
			same = 0
		}
	}
	cur := length - start - same
	if ans < cur {
		ans = cur
	}
	return ans
}
//o（n）解法，使用哈希
func longestConsecutive2(nums []int) int {
	ans := 0
	nums_map := map[int]int{}
	for _, v := range nums {
		nums_map[v] = 1
	}
	for v, _ := range nums_map {
		_, ok := nums_map[v-1]
		if !ok {
			cur := 1
			for true {
				_, ok = nums_map[v+1]
				if ok {
					v++
					cur++
					continue
				}else {
					if ans < cur {
						ans = cur
					}
					break
				}
			}
		}
	}
	return ans
}