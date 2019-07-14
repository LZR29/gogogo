package main

func lengthOfLongestSubstring(s string) int {
	if len(s) < 1 {
		return 0
	}
	max, start, end := 0, 0, 0
	m := make(map[byte]int)
	for start, end = 0, 0; end < len(s); end++ {
		v, ok := m[s[end]]
		if ok {
			if v > start{
				start = v
			}
		}
		cur := end - start + 1
		if cur > max {
			max = cur
		}
		m[s[end]] = end + 1
	}
	return max
}

func main() {
	
}
