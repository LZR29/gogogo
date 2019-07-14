package array_sort

func find(parents []int, i int) int {
	if parents[i] == i {
		return i
	}
	parents[i] = find(parents, parents[i])
	return parents[i]
}
func findCircleNum(M [][]int) int {
	nums := len(M)
	parents := make([]int, nums)
	for i := range parents {
		parents[i] = i
	}
	length := nums
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			if M[i][j] == 1 {
				p1 := find(parents, i)
				p2 := find(parents, j)
				if p1 != p2 {
					parents[p1] = p2
					nums--
				}
			}
		}
	}
	return nums
}