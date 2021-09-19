package getting_started

//按位 与 运算
func isPowerOfTwo(n int) bool {
	return n > 0 && (n&(n-1)) == 0
}

func hammingWeight(num uint32) int {
	ret := 0
	k := uint32(1)
	for i := 0; i < 32; i++ {
		if num&k == 1 {
			ret++
		}
		num >>= 1
	}
	return ret
}
