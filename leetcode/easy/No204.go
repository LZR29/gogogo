package easy

//计算质数，使用厄拉多塞筛法，即划掉那些质数的倍数，剩下的就是质数了
func countPrimes(n int) int {
	primes := make([]bool, n)
	for i := 2; i < n; i++ {
		primes[i] = true
	}

	for i := 2; i < n; i++ {
		//因为 i*2 到 i*i-1的那些数肯定不是质数了，所以从i*i
		if primes[i] {
			for j := i * i; j < n; j += i {
				primes[j] = false
			}
		}
	}

	count := 0
	for _, v := range primes {
		if v {
			count++
		}
	}
	return count
}
