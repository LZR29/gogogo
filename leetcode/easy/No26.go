package easy

func removeDuplicates(nums []int) int {
	j := 1
	n := len(nums)
	for i := 1; i < n; i++ {
		if nums[i] != nums[i-1] {
			nums[j] = nums[i]
			j++
		}
	}
	return j
}
