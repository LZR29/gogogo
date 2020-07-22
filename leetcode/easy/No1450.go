package easy

func busyStudent(startTime []int, endTime []int, queryTime int) int {
	ans := 0
	n := len(startTime)
	for i := 0; i < n; i++ {
		if startTime[i] <= queryTime && endTime[i] >= queryTime {
			ans++
		}
	}
	return ans
}
