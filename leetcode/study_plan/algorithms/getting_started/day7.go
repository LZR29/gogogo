package getting_started

var (
	dx = []int{1, 0, 0, -1}
	dy = []int{0, 1, -1, 0}
)

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	cur := image[sr][sc]
	if cur != newColor {
		dfsFloodFill(image, sr, sc, cur, newColor)
	}
	return image
}

func dfsFloodFill(image [][]int, sr int, sc, curColor, newColor int) {
	if image[sr][sc] == curColor {
		image[sr][sc] = newColor
		for i := 0; i < 4; i++ {
			x, y := sr+dx[i], sc+dy[i]
			if x >= 0 && x < len(image) && y >= 0 && y < len(image[0]) {
				dfsFloodFill(image, x, y, curColor, newColor)
			}
		}
	}
}

func maxAreaOfIsland(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}
	x := len(grid)
	y := len(grid[0])
	ret := 0
	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			if grid[i][j] == 1 {
				cur := dfsMaxAreaOfIsland(grid, i, j)
				if cur > ret {
					ret = cur
				}
			}
		}
	}
	return ret
}

func dfsMaxAreaOfIsland(grid [][]int, x, y int) int {
	if x < 0 || y < 0 || x >= len(grid) || y >= len(grid[0]) || grid[x][y] == 0 {
		return 0
	}
	grid[x][y] = 0
	return 1 +
		dfsMaxAreaOfIsland(grid, x+1, y) +
		dfsMaxAreaOfIsland(grid, x-1, y) +
		dfsMaxAreaOfIsland(grid, x, y-1) +
		dfsMaxAreaOfIsland(grid, x, y+1)
}
