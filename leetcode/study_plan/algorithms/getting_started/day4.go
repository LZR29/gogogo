package getting_started

func reverseString(s []byte) {
	start := 0
	end := len(s) - 1
	for start < end {
		s[start], s[end] = s[end], s[start]
		start++
		end--
	}
}

func reverseWords(s string) string {
	bytes := []byte(s)
	pre := 0
	for i := 0; i < len(bytes); i++ {
		if bytes[i] == ' ' {
			reverse(bytes, pre, i-1)
			pre = i + 1
		}
	}
	reverse(bytes, pre, len(bytes)-1)
	return string(bytes)
}

func reverse(s []byte, start, end int) {
	for start < end {
		s[start], s[end] = s[end], s[start]
		start++
		end--
	}
}
