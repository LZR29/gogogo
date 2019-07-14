package array_sort

import (
	"fmt"
	"sort"
)

func merge(intervals [][]int) [][]int {
	length := len(intervals)
	if length <= 1 {
		return intervals
	}
	sort.Sort(Interval(intervals))
	fmt.Println(intervals)
	var ans [][]int
	start, end := intervals[0][0], intervals[0][1]
	for i := 1; i < length; i++ {
		if end >= intervals[i][0] {
			if end < intervals[i][1] {
				end = intervals[i][1]
			}
		}else {
			ans = append(ans, []int{start, end})
			start = intervals[i][0]
			end = intervals[i][1]
		}
	}
	ans = append(ans, []int{start, end})
	return ans
}

type Interval [][]int

func (i Interval) Len() int {
	return len(i)
}
func (i Interval) Swap(j, k int) {
	i[j], i[k] = i[k], i[j]
}
func (i Interval) Less(j, k int) bool {
	return  i[j][0] < i[k][0]
}