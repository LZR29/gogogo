package getting_started

func lengthOfLongestSubstring(s string) int {
	n := len(s)
	ret, start, end := 0, 0, 0
	bytesMap := make(map[byte]int) //下标map
	for ; end < n; end++ {
		index, ok := bytesMap[s[end]]
		if ok && start < index {
			start = index
		}
		cur := end - start + 1
		if cur > ret {
			ret = cur
		}
		bytesMap[s[end]] = end + 1
	}
	return ret
}

func checkInclusion(s1, s2 string) bool {
	n, m := len(s1), len(s2)
	if n > m {
		return false
	}
	cnt := [26]int{}
	for i, ch := range s1 {
		cnt[ch-'a']--
		cnt[s2[i]-'a']++
	}
	diff := 0
	for _, c := range cnt[:] {
		if c != 0 {
			diff++
		}
	}
	if diff == 0 {
		return true
	}
	for i := n; i < m; i++ {
		x, y := s2[i]-'a', s2[i-n]-'a'
		if x == y {
			continue
		}
		if cnt[x] == 0 {
			diff++
		}
		cnt[x]++
		if cnt[x] == 0 {
			diff--
		}
		if cnt[y] == 0 {
			diff++
		}
		cnt[y]--
		if cnt[y] == 0 {
			diff--
		}
		if diff == 0 {
			return true
		}
	}
	return false
}
