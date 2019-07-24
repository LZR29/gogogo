package easy

import "strings"

var oneLine = map[byte]bool{'Q': true, 'W': true, 'E': true, 'R': true, 'T': true, 'Y': true, 'U': true, 'I': true, 'O': true, 'P': true}
var twoLine = map[byte]bool{'A': true, 'S': true, 'D': true, 'F': true, 'G': true, 'H': true, 'J': true, 'K': true, 'L': true}
var threeLine = map[byte]bool{'Z': true, 'X': true, 'C': true, 'V': true, 'B': true, 'N': true, 'M': true}

func findWords(words []string) []string {
	ans := make([]string, 0)
	for _, v := range words {
		up := strings.ToUpper(v)
		flag := true
		if _, ok := oneLine[up[0]]; ok {
			for i := 0; i < len(up); i++ {
				if !oneLine[up[i]] {
					flag = false
					break
				}
			}
		}
		if _, ok := twoLine[up[0]]; ok {
			for i := 0; i < len(up); i++ {
				if !twoLine[up[i]] {
					flag = false
					break
				}
			}
		}
		if _, ok := threeLine[up[0]]; ok {
			for i := 0; i < len(up); i++ {
				if !threeLine[up[i]] {
					flag = false
					break
				}
			}
		}
		if flag {
			ans = append(ans, v)
		}
	}
	return ans
}
