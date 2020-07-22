package medium

func countTriplets(arr []int) int {
	ans := 0
	n := len(arr)
	if n < 2 {
		return 0
	}
	for i := 0; i < n-1; i++ {
		tem := arr[i]
		for k := i + 1; k < n; k++ {
			tem ^= arr[k]
			if tem == 0 {
				ans += k - i
			}
		}
	}
	return ans
}
