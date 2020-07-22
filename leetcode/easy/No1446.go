package easy

func maxPower(s string) int {
	n := len(s)
	ans := 0
	temp := 1
	for i := 1; i < n; i++ {
		if s[i-1] == s[i] {
			temp++
		} else {
			if temp > ans {
				ans = temp
			}
			temp = 1
		}
	}
	if temp > ans {
		ans = temp
	}
	return ans
}
