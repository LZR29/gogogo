package array_sort

func search(nums []int, target int) int {
	return helperSearch(nums, 0, len(nums) - 1, target)
}

func helperSearch(nums []int, start, end, target int)  int {
	if start > end {
		return -1
	}
	mid  := start / 2 + end / 2
	if nums[mid] == target {
		return mid
	}
	if nums[mid] < nums[end]{
		if target <= nums[end] && target > nums[mid] {
			return helperSearch(nums,mid+1,end,target)
		}else {
			return helperSearch(nums,start,mid-1,target)
		}
	}else {
		if nums[mid] > target && target >= nums[start]{
			return helperSearch(nums,start,mid-1,target)
		}else {
			return helperSearch(nums,mid+1,end,target)
		}
	}
}