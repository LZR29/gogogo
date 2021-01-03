package contest

import (
	"sort"
)

//5641. 卡车上的最大单元数
func maximumUnits(boxTypes [][]int, truckSize int) int {
	ret := 0
	//排序
	sort.Slice(boxTypes, func(i, j int) bool {
		return boxTypes[i][1] > boxTypes[j][1]
	})
	for _, box := range boxTypes {
		cnt := truckSize
		if box[0] < cnt {
			cnt = box[0]
		}
		ret += cnt * box[1]
		truckSize -= cnt
		if truckSize <= 0 {
			break
		}
	}
	return ret
}

//5642. 大餐计数
func countPairs(deliciousness []int) int {
	sum := []int{1}
	ret := 0
	mod := int(1e9 + 7)
	numMap := make(map[int]int)
	for i := 0; i < 22; i++ {
		sum = append(sum, 2<<i)
	}
	for _, v := range deliciousness {
		for _, temp := range sum {
			dif := temp - v
			if cnt, ok := numMap[dif]; ok {
				ret += cnt
			}
		}
		numMap[v]++
	}
	ret = ret % mod
	return ret
}
