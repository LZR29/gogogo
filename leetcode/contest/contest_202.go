package contest

//5185. 存在连续三个奇数的数组
func threeConsecutiveOdds(arr []int) bool {
	n := len(arr)
	for i := 0; i < n-2; i++ {
		if arr[i]&arr[i+1]&arr[i+2]&1 == 1 {
			return true
		}
	}
	return false
}

//5488. 使数组中所有元素相等的最小操作数
func minOperations(n int) int {
	if n%2 == 0 {
		return n * n / 4
	} else {
		return (n + 1) * (n - 1) / 4
	}
}
