package main

import (
	"bytes"
	"strings"
)

func simplifyPath(path string) string {
	words := strings.Split(path, "/")
	var ans []string
	length := len(words)
	for i := 0; i < length; i++ {
		if strings.Compare(".", words[i]) == 0 || strings.Compare(words[i], "") == 0{
			continue
		}else if strings.Compare(words[i], "..") == 0{
			if len(ans) != 0{
				ans = ans[:len(ans) - 1]
			}
		}else {
			ans = append(ans, words[i])
		}
	}
	if len(ans) == 0{
		return "/"
	}
	var buffer bytes.Buffer
	for _, v := range ans{
		buffer.WriteString("/")
		buffer.WriteString(v)
	}
	return buffer.String()
}
