package easy

func buildArray(target []int, n int) []string {
	push := "Push"
	pop := "Pop"
	res := make([]string, 0)
	len := len(target)
	j := 0
	for i := 1; i <= n; i++ {
		if target[j] == i {
			res = append(res, push)
			j++
			if j >= len {
				break
			}
		} else {
			res = append(res, push, pop)
		}
	}
	return res
}
