package main

import "strings"

func longestCommonPrefix(strs []string) string {
	if len(strs) < 1 {
		return ""
	}
	var ans  = strs[0]
	for i := 1; i < len(strs) ;  {
		if strings.Index(strs[i], ans) != 0 {
			ans = ans[: len(ans) - 1]
		}else {
			i++
		}
	}
	return ans
}

func main() {
	
}
