package main

import "strconv"

func multiply(num1 string, num2 string) string {
	len1 := len(num1)
	len2 := len(num2)
	var ans  = make([]int, len1 + len2)
	for i := len1 - 1; i >= 0; i-- {
		for j := len2 - 1; j >= 0; j-- {
			mul := int((num1[i] - '0') * (num2[j] - '0'))
			p1 := i + j
			p2 := i + j + 1
			mul += ans[p2]
			ans[p1] += mul / 10
			ans[p2] = mul % 10
		}
	}
	var res string
	for _, v := range ans {
		if !(len(res) == 0 && v == 0){
			res += strconv.Itoa(v)
		}
	}
	if len(res) == 0{
		return "0"
	}
	return res
}