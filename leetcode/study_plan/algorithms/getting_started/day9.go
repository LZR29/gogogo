package getting_started

import "fmt"

var dirs = [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}

func updateMatrix(mat [][]int) [][]int {
	m, n := len(mat), len(mat[0])
	dist := make([][]int, m)
	seen := make([][]bool, m)
	fmt.Println(dist, seen)
	queue := [][]int{}
	for i := 0; i < m; i++ {
		dist[i] = make([]int, n)
		seen[i] = make([]bool, n)
		for j := 0; j < n; j++ {
			if mat[i][j] == 0 {
				queue = append(queue, []int{i, j})
				seen[i][j] = true
			}
		}
	}
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		i, j := cur[0], cur[1]
		for d := 0; d < 4; d++ {
			ni := i + dirs[d][0]
			nj := j + dirs[d][1]
			if ni >= 0 && ni < m && nj >= 0 && nj < n && !seen[ni][nj] {
				dist[ni][nj] = dist[i][j] + 1
				seen[ni][nj] = true
				queue = append(queue, []int{ni, nj})
			}
		}
	}
	return dist
}

func orangesRotting(grid [][]int) int {
	queue := [][]int{}
	newQueue := [][]int{}
	m, n := len(grid), len(grid[0])
	ret := 0
	count := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 2 {
				queue = append(queue, []int{i, j})
			}
			if grid[i][j] == 1 {
				count++
			}
		}
	}
	for len(queue) > 0 && count > 0 {
		ret++
		for len(queue) > 0 {
			cur := queue[0]
			queue = queue[1:]
			for d := 0; d < 4; d++ {
				ni, nj := dirs[d][0]+cur[0], dirs[d][1]+cur[1]
				if ni >= 0 && ni < m && nj >= 0 && nj < n && grid[ni][nj] == 1 {
					grid[ni][nj] = 2
					count--
					newQueue = append(newQueue, []int{ni, nj})
				}
			}
		}
		queue = newQueue
		fmt.Println(queue)
		newQueue = [][]int{}
	}
	if count != 0 {
		return -1
	}
	return ret
}
