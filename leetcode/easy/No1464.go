package easy

func maxProduct(nums []int) int {
	res := 0
	n := len(nums)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			temp := (nums[i] - 1) * (nums[j] - 1)
			if res < temp {
				res = temp
			}
		}
	}
	return res
}
