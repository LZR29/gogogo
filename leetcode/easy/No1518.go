package easy

func numWaterBottles(numBottles int, numExchange int) int {
	res := numBottles
	remainder := 0 //余数
	total := 0     //空瓶子
	for numBottles+remainder >= numExchange {
		total = numBottles + remainder
		numBottles = total / numExchange
		res += numBottles
		remainder = total % numExchange
	}
	return res
}
