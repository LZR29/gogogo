package contest

//5613. 最富有客户的资产总量
func maximumWealth(accounts [][]int) int {
	max := 0
	for _, cur := range accounts {
		tempMax := 0
		for _, v := range cur {
			tempMax += v
		}
		if tempMax > max {
			max = tempMax
		}
	}
	return max
}

//5614. 找出最具竞争力的子序列
//单调栈问题
func mostCompetitive(nums []int, k int) []int {
	n := len(nums)
	if n == k {
		return nums
	}
	stack := make([]int, 0, n)
	stackSize := 0
	for i := 0; i < n; i++ {
		//栈加剩余未遍历的总数要大于k
		for stackSize+n-i > k && stackSize > 0 && nums[i] < stack[stackSize-1] {
			stack = stack[:stackSize-1]
			stackSize--
		}
		stack = append(stack, nums[i])
		stackSize++
	}
	return stack[:k]
}
func main() {
	mostCompetitive([]int{3, 5, 2, 6}, 2)
}
