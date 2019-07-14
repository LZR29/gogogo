package array_sort

func findLengthOfLCIS(nums []int) int {
	start := 0
	max := 0
	length := len(nums)
	for i := 1; i < length ; i++ {
		if nums[i] <= nums[i-1] {
			cur := i - start
			if cur > max {
				max = cur
			}
			start = i
		}
	}
	if max < length - start {
		max = length - start
	}
	return max
}
