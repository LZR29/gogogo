package easy

//分割数组成正数和负数两部分，然后合并
func sortedSquares(A []int) []int {
	index := 0
	n := len(A)
	ans := make([]int, n)
	for i := 1; i < n; i++ {
		if A[i] > 0 && A[i-1] <= 0 {
			index = i
			break
		}
	}
	pos := index
	neg := index - 1
	i := 0
	for i = 0; i < n; i++ {
		if pos < n && neg >= 0 {
			num1 := A[pos] * A[pos]
			num2 := A[neg] * A[neg]
			if num1 > num2 {
				ans[i] = num2
				neg--
			} else {
				ans[i] = num1
				pos++
			}
		} else {
			break
		}
	}
	if pos < n {
		for ; i < n; i++ {
			ans[i] = A[pos] * A[pos]
			pos++
		}
	}
	if neg >= 0 {
		for ; i < n; i++ {
			ans[i] = A[neg] * A[neg]
			neg--
		}
	}
	return ans
}
