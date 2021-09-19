package getting_started

func sortedSquares(nums []int) []int {
	index := 0
	n := len(nums)
	ret := make([]int, 0, n)
	for i := 1; i < n; i++ {
		if nums[i-1] >= 0 {
			index = i - 1
			break
		}
		if nums[i-1] < 0 && nums[i] > 0 {
			index = i
			break
		}
		index = i
	}
	pre := index - 1
	for pre >= 0 || index < n {
		preRet := 100000001
		indexRet := 100000001
		if pre >= 0 {
			preRet = nums[pre] * nums[pre]
		}
		if index < n {
			indexRet = nums[index] * nums[index]
		}
		if preRet < indexRet {
			ret = append(ret, preRet)
			pre--
		} else {
			ret = append(ret, indexRet)
			index++
		}
	}
	return ret
}

func rotate(nums []int, k int) {
	n := len(nums)
	k = k % n
	temp := make([]int, 0, k)
	for i := n - k; i < n; i++ {
		temp = append(temp, nums[i])
	}
	for i := n - 1 - k; i >= 0; i-- {
		nums[i+k] = nums[i]
	}
	for i := 0; i < k; i++ {
		nums[i] = temp[i]
	}
}
