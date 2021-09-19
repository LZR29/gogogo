package getting_started

func combine(n int, k int) [][]int {
	var combineRet [][]int
	var combineDfs func(curVal int, n, k int, curList []int)
	combineDfs = func(curVal int, n, k int, curList []int) {
		if len(curList) == k {
			//fmt.Println(curList)
			combineRet = append(combineRet, curList)
			return
		}
		for i := curVal + 1; i <= n; i++ {
			nextList := make([]int, len(curList))
			copy(nextList, curList)
			nextList = append(nextList, i)
			combineDfs(i, n, k, nextList)
		}
		return
	}
	for i := 1; i <= n-k+1; i++ {
		cur := make([]int, 0, 1)
		cur = append(cur, i)
		combineDfs(i, n, k, cur)
	}
	return combineRet
}

func permute(nums []int) [][]int {
	var ret [][]int
	n := len(nums)
	seen := make([]bool, n)
	var helper func(i int, nums []int, curRet []int)
	helper = func(i int, nums []int, curRet []int) {
		if len(curRet) == n {
			cur := make([]int, n)
			copy(cur, curRet)
			ret = append(ret, cur)
			return
		}
		for j := 0; j < n; j++ {
			if !seen[j] {
				seen[j] = true
				curRet = append(curRet, nums[j])
				helper(j, nums, curRet)
				seen[j] = false
				curRet = curRet[:len(curRet)-1]
			}
		}
	}
	helper(0, nums, []int{})
	return ret
}

func letterCasePermutation(s string) []string {
	bytes := []byte(s)
	n := len(bytes)
	ret := make([]string, 0)
	var helper func(i int, data []byte)
	helper = func(i int, data []byte) {
		ret = append(ret, string(data))
		if i == n {
			return
		}
		for j := i; j < n; j++ {
			if data[j] >= 'a' && data[j] <= 'z' {
				data[j] -= 32
				helper(j+1, data)
				data[j] += 32
			} else if data[j] >= 'A' && data[j] <= 'Z' {
				data[j] += 32
				helper(j+1, data)
				data[j] -= 32
			}
		}
	}
	helper(0, bytes)
	return ret
}
