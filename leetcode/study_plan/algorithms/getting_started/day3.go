package getting_started

func moveZeroes(nums []int) {
	index := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[index] = nums[i]
			index++
		}
	}
	for i := index; i < len(nums); i++ {
		nums[i] = 0
	}
	return
}

func twoSum(nums []int, target int) []int {
	left, right := 0, len(nums)-1
	for left < right {
		cur := nums[left] + nums[right]
		if cur == target {
			return []int{left + 1, right + 1}
		}
		if cur > target {
			right--
		} else {
			left++
		}
	}
	return []int{}
}
