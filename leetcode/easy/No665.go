package easy

func checkPossibility(nums []int) bool {
	count1 := 0
	count2 := 0
	n := len(nums)
	judge := nums[0]
	for i := 1; i < n; i++ {
		if nums[i] >= judge {
			judge = nums[i]
		} else {
			count1++
			if count1 >= 2 {
				break
			}
		}
	}
	judge = nums[n-1]
	for i := n - 2; i >= 0; i-- {
		if judge >= nums[i] {
			judge = nums[i]
		} else {
			count2++
			if count2 >= 2 {
				break
			}
		}
	}
	return count1 < 2 || count2 < 2
}
