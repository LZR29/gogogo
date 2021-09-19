package easy

func generateTheString(n int) string {
	data := make([]byte, 0, n)
	flag := false
	if n%2 == 0 {
		flag = true
		n--
	}
	for i := 0; i < n; i++ {
		data = append(data, 'a')
	}
	if flag {
		data = append(data, 'b')
	}
	return string(data)
}
