package easy

//异或运算+与运算判断+位移即可
func hammingDistance(x int, y int) int {
	//把xy异或得到某个数
	res := x ^ y
	count := 0
	for res != 0 {
		//res此时为奇数，即最低位为1，异或结果为1，则表明该位上x和y是不同的，满足要求
		if res&1 == 1 {
			count++
		}
		//右移下一位
		res >>= 1
	}
	return count
}
