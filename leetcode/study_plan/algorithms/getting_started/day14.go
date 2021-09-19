package getting_started

func reverseBits(n uint32) uint32 {
	ret := uint32(0)
	for i := 0; i < 32 && n > 0; i++ {
		ret += n & 1 << (31 - i)
		n >>= 1
	}
	return ret
}

func singleNumber(nums []int) int {
	ret := nums[0]
	for i := 1; i < len(nums); i++ {
		ret ^= nums[i]
	}
	return ret
}
