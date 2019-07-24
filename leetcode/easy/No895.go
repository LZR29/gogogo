package easy

func buddyStrings(A string, B string) bool {
	if len(A) != len(B) {
		return false
	}
	//如果两个字符串相等，那么只有存在某个字符的出现次数大于1时才能交换返回true，不然无法交换即false
	//如果两个字符串存在两个位置上不等，且两个位置上交换会使字符串相等才返回true否则是false
	if buddyStringsHelper(A, B) {
		m := make(map[byte]int)
		for i := 0; i < len(A); i++ {
			m[A[i]]++
		}
		for _, v := range m {
			if v >= 2 {
				return true
			}
		}
		return false
	}
	first := -1
	second := -1
	for i := 0; i < len(A); i++ {
		if A[i] != B[i] {
			if first == -1 {
				first = i
			} else if second == -1 {
				second = i
			} else {
				return false
			}
		}
	}
	return second != -1 && A[first] == B[second] && A[second] == B[first]
}

func buddyStringsHelper(A, B string) bool {
	for i := 0; i < len(A); i++ {
		if A[i] != B[i] {
			return false
		}
	}
	return true
}
