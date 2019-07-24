package easy

func flipAndInvertImage(A [][]int) [][]int {
	length := len(A)
	for i := 0; i < length; i++ {
		s := 0
		e := length - 1
		for s < e {
			A[i][s], A[i][e] = A[i][e], A[i][s]
			if A[i][s] == 0 {
				A[i][s] = 1
			} else {
				A[i][s] = 0
			}
			if A[i][e] == 0 {
				A[i][e] = 1
			} else {
				A[i][e] = 0
			}
			s++
			e--
		}
		if s == e {
			if A[i][e] == 0 {
				A[i][e] = 1
			} else {
				A[i][e] = 0
			}
		}
	}
	return A
}
