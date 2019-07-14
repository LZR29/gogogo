package main

import (
	"bytes"
	"fmt"
	"strconv"
)

func restoreIpAddresses(s string) []string {
	var ans []string
	var buffer bytes.Buffer
	getAns(s, 0, buffer, &ans, 0)
	return ans
}

func getAns(s string, start int, buffer bytes.Buffer, ans *[]string, count int)  {
	length := len(s)
	if length - start > 3 * (4 - count) {
		return
	}

	if start == length {
		if count == 4 {
			cur := buffer.String()
			cur = cur[:len(cur) - 1]
			fmt.Println(cur)
			*ans = append(*ans, cur)
		}
		return
	}

	if start > length || count == 4 {
		return
	}
	pre := buffer

	//加入一位数
	buffer.WriteByte(s[start])
	buffer.WriteByte('.')
	getAns(s, start + 1, buffer, ans, count + 1)

	if s[start] == '0' {
		return
	}

	//加入二位数
	if start + 1 < length {
		buffer = pre
		buffer.WriteString(s[start:start+2])
		buffer.WriteString(".")
		getAns(s, start + 2, buffer, ans, count + 1)
	}

	//加入三位数
	if start + 2 < length {
		buffer = pre
		sum , _ := strconv.ParseInt(s[start:start+3], 10, 32)
		if sum >= 0 && sum <= 255{
			buffer.WriteString(s[start:start+3])
			buffer.WriteString(".")
			getAns(s, start + 3, buffer, ans, count + 1)
		}
	}
}