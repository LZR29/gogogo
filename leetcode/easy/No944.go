package easy

func minDeletionSize(A []string) int {
	row := len(A)
	col := len(A[0])
	count := 0
	for i := 0; i < col; i++ {
		for j := 0; j < row-1; j++ {
			if A[j][i] > A[j+1][i] {
				count++
				break
			}
		}
	}
	return count
}
