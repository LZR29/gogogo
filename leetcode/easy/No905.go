package easy

import "fmt"

func sortArrayByParity(A []int) []int {
	n := len(A)
	ans := make([]int, n)
	i := 0
	fmt.Println(A[0] ^ 1)
	for j := 0; j < n; j++ {
		if A[j]&1 == 1 {
			ans[i] = A[j]
			i++
		}
	}
	for j := 0; j < n; j++ {
		if A[j]&1 == 0 {
			ans[i] = A[j]
			i++
		}
	}
	return ans
}
