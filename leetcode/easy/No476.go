package easy

//找出一个比它大的全部都是1的数，然后异或
func findComplement(num int) int {
	i := 1
	for i < num {
		i <<= 1
	}
	i--
	return i ^ num
}
