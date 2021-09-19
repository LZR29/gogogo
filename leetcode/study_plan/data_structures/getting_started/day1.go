package getting_started

func containsDuplicate(nums []int) bool {
	m := make(map[int]bool)
	for i := 0; i < len(nums); i++ {
		if m[nums[i]] {
			return true
		}
		m[nums[i]] = true
	}
	return false
}

func maxSubArray(nums []int) int {
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i]+nums[i-1] > nums[i] {
			nums[i] += nums[i-1]
		}
		if nums[i] > max {
			max = nums[i]
		}
	}
	return max
}
