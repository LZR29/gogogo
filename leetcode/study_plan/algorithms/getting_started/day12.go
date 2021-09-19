package getting_started

func climbStairs(n int) int {
	if n == 1 || n == 2 {
		return n
	}
	dp := make([]int, n)
	dp[0] = 1
	dp[1] = 2
	for i := 2; i < n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n-1]
}

//用dp[i] 表示前 i 间房屋能偷窃到的最高总金额，那么就有如下的状态转移方程：
//dp[i] = max(dp[i-2]+nums[i],dp[i-1])
func rob(nums []int) int {
	n := len(nums)
	if n == 0 {
		return n
	}
	if n == 1 {
		return nums[0]
	}
	dp := make([]int, n)
	dp[0] = nums[0]
	dp[1] = max(dp[0], nums[1])
	for i := 2; i < n; i++ {
		dp[i] = max(dp[i-2]+nums[i], dp[i-1])
	}
	return dp[n-1]
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

//dp[i][j]表示顶点到第i层下标j的最小路径和
//dp[i][j] = min(dp[i-1][j],dp[i-1][j-1])+triangle[i][j]
func minimumTotal(triangle [][]int) int {
	n := len(triangle)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, len(triangle[i]))
	}
	dp[0][0] = triangle[0][0]
	for i := 1; i < n; i++ {
		curLen := len(triangle[i])
		for j := 0; j < curLen; j++ {
			if j == 0 {
				dp[i][j] = dp[i-1][j] + triangle[i][j]
			} else if j == curLen-1 {
				dp[i][j] = dp[i-1][j-1] + triangle[i][j]
			} else {
				dp[i][j] = min(dp[i-1][j], dp[i-1][j-1]) + triangle[i][j]
			}
		}
	}
	ret := 10001 //题目最大值
	for i := 0; i < len(dp[n-1]); i++ {
		ret = min(ret, dp[n-1][i])
	}
	return ret
}
