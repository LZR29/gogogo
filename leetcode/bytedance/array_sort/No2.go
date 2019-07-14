package array_sort

func maxAreaOfIsland(grid [][]int) int {
	if grid == nil || len(grid) == 0 {
		return 0
	}
	ans := 0
	row := len(grid)
	col := len(grid[0])
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if grid[i][j] == 1 {
				cur := helper(grid, i ,j)
				if cur > ans {
					ans = cur
				}
			}
		}
	}
	return ans
}
func helper(grid [][]int, i, j int) int  {
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) || grid[i][j] == 0 {
		return 0
	}
	grid[i][j] = 0;
	return 1 +
		helper(grid, i - 1, j) +
		helper(grid, i + 1, j) +
		helper(grid, i ,j - 1) +
		helper(grid, i, j + 1)
}
