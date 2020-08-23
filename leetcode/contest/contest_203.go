package contest

import "sort"

//5495.圆形赛道上经过次数最多的扇区
//rounds[0] >= rounds[len - 1]，结果为[1, rounds[len - 1]]与[rounds[0], n]的并集
//rounds[0] < rounds[len - 1]，结果为[rounds[0],rounds[len - 1]]
func mostVisited(n int, rounds []int) []int {
	ret := make([]int, 0)
	length := len(rounds)
	start := rounds[0]
	end := rounds[length-1]
	if start > end {
		for i := 1; i <= end; i++ {
			ret = append(ret, i)
		}
		for i := start; i <= n; i++ {
			ret = append(ret, i)
		}
	} else {
		for i := start; i <= end; i++ {
			ret = append(ret, i)
		}
	}
	return ret
}

//5496你可以获得的最大硬币数目
func maxCoins(piles []int) int {
	sort.Ints(piles)
	n := len(piles)
	sum := 0
	k := n / 3
	for i := n - 1; i > 0; i = i - 2 {
		sum += piles[i-1]
		k--
		if k <= 0 {
			break
		}
	}
	return sum
}
