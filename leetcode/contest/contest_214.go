package contest

//5561
//1.构造符合规则的数据后进行排序得到最大值
//2.构造过程进行保存最大值
func getMaximumGenerated(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	nums := make([]int, n+1)
	nums[0] = 0
	nums[1] = 1
	max := 1
	for i := 1; 2*i <= n; i++ {
		if 2*i <= n {
			nums[2*i] = nums[i]
			if max < nums[2*i] {
				max = nums[2*i]
			}
		}
		if 2*i+1 <= n {
			nums[2*i+1] = nums[i] + nums[i+1]
			if max < nums[2*i+1] {
				max = nums[2*i+1]
			}
		}
	}
	//sort.Ints(nums)
	return max
}

//统计字母的个数后构造一个set得到个数集合和一个重复个数的数组，然后对重复个数数组的每个数
//进行自减直到在set中不存在这个重复个数
func minDeletions(s string) int {
	ans := 0
	cnts := make(map[int32]int)
	for _, v := range s[:] {
		cnts[v]++
	}
	set := make(map[int]bool)
	repeatNums := []int{}
	for _, v := range cnts {
		if !set[v] {
			set[v] = true
		} else {
			repeatNums = append(repeatNums, v)
		}
	}
	for _, v := range repeatNums {
		for set[v] && v != 0 {
			v--
			ans++
		}
		set[v] = true
	}
	return ans
}
