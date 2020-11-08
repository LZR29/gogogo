package easy

func moveZeroes(nums []int) {
	n := len(nums)
	j := 0
	for i := 0; i < n; i++ {
		if nums[i] != 0 {
			nums[j] = nums[i]
			if j != i {
				nums[i] = 0
			}
			j++
		}
	}

}
