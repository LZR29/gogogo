package dp

func maxProfit(prices []int) int {
	if len(prices) < 1 {
		return 0
	}
	max := 0
	min := prices[0]
	for i := 1; i < len(prices); i++ {
		if prices[i] < min {
			min = prices[i]
		} else if max < prices[i]-min {
			max = prices[i] - min
		}
	}
	return max
}
