package contest

import "strings"

//5605. 检查两个字符串数组是否相等
func arrayStringsAreEqual(word1 []string, word2 []string) bool {
	str1 := strings.Join(word1, "")
	str2 := strings.Join(word2, "")
	return str1 == str2
}

//5606. 具有给定数值的最小字符串
func getSmallestString(n int, k int) string {
	bytes := make([]byte, n)
	for i := 0; i < n; i++ {
		bytes[i] = 97 //a=97
	}
	k = k - n //减去初始化为a的总和
	for i := n - 1; i >= 0; i-- {
		if k > 0 {
			if k <= 25 {
				bytes[i] += byte(k)
				break
			} else {
				k -= 25
				bytes[i] += 25
			}
		}
	}
	return string(bytes)
}
