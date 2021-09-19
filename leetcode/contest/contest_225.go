package contest

//5661. 替换隐藏数字得到的最晚时间
//i == 0时，如果i==1是?或在0-3的范围，则为2，否则为1
//i == 1时，如果i==0时2或?,则为3，否则为9
func maximumTime(time string) string {
	data := []byte(time)
	ret := []byte{}
	for i, v := range data {
		if v == '?' {
			if i == 0 {
				if data[1] == '?' || data[1] < '4' {
					ret = append(ret, '2')
				} else {
					ret = append(ret, '1')
				}
				continue
			}
			if i == 1 {
				if data[0] == '2' || data[0] == '?' {
					ret = append(ret, '3')
				} else {
					ret = append(ret, '9')
				}
				continue
			}
			if i == 3 {
				ret = append(ret, '5')
				continue
			}
			if i == 4 {
				ret = append(ret, '9')
				continue
			}

		}
		ret = append(ret, v)
	}
	return string(ret)
}
