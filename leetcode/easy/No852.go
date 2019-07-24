package easy

func peakIndexInMountainArray(A []int) int {
	n := len(A)
	for i := 1; i < n-1; i++ {
		if A[i] > A[i+1] && A[i] > A[i-1] {
			return i
		}
	}
	return 0
}
