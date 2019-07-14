package main

import (
	"bytes"
	"strings"
)

func reverseWords(s string) string {
	words := strings.Split(s, " ")
	len := len(words)
	ans := ""
	for i := len - 1; i >= 0; i-- {
		if strings.Compare(words[i], "") != 0 {
			ans += words[i] + " "
		}
	}
	return strings.Trim(ans," ")
}
//参考使用了一些不知道的函数
func reverseWords2(s string) string {
	words := strings.Fields(s)
	var buffer bytes.Buffer
	for i := len(words) - 1; i>= 0; i--  {
		buffer.WriteString(words[i])
		if i != 0{
			buffer.WriteString(" ")
		}
	}
	return buffer.String()
}
