package dp

//又是一道看题解的题
func maximalSquare(matrix [][]byte) int {
	if matrix == nil || len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	var max int = 0
	n := len(matrix)
	m := len(matrix[0])
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, m+1)
	}
	for i := range matrix {
		for j := range matrix[i] {
			if matrix[i][j] == '1' {
				min := dp[i][j]
				if min > dp[i+1][j] {
					min = dp[i+1][j]
				}
				if min > dp[i][j+1] {
					min = dp[i][j+1]
				}
				dp[i+1][j+1] = min + 1
				if max < min+1 {
					max = min + 1
				}
			}
		}
	}
	return max * max
}
