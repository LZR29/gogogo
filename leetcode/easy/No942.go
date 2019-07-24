package easy

func diStringMatch(S string) []int {
	n := len(S)
	ans := make([]int, n+1)
	low := 0
	high := n
	for i, v := range S {
		if v == 'I' {
			ans[i] = low
			low++
		} else if v == 'D' {
			ans[i] = high
			high--
		}
	}
	ans[n] = low
	return ans
}
