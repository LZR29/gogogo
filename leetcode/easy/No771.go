package easy

func numJewelsInStones(J string, S string) int {
	ans := 0
	m := map[rune]int{}
	for _, v := range S {
		m[v] = m[v] + 1
	}
	for _, v := range J {
		ans += m[v]
	}
	return ans
}
