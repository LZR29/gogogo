package easy

func selfDividingNumbers(left int, right int) []int {
	ans := make([]int, 0)
	for left <= right {
		if selfDividingNumbersHelper(left) {
			ans = append(ans, left)
		}
		left++
	}
	return ans
}

func selfDividingNumbersHelper(num int) bool {
	temp := num
	for temp != 0 {
		k := temp % 10
		if k == 0 || num%k != 0 {
			return false
		}
		temp /= 10
	}
	return true
}
