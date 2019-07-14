package array_sort

import "container/heap"

func findKthLargest(nums []int, k int) int {
	h := &IntHeap{}
	heap.Init(h)

	length := len(nums)
	for i := 0; i < length; i++ {
		if i < k {
			heap.Push(h, nums[i])
		}else {
			if nums[i] > (*h)[0]{
				heap.Pop(h)
				heap.Push(h, nums[i])
			}
		}
	}
	return (*h)[0]
}

type IntHeap []int

func (h IntHeap) Len() int  {
	return len(h)
}
func (h IntHeap) Less(i, j int) bool {
	// 如果h[i]<h[j]生成的就是小根堆，如果h[i]>h[j]生成的就是大根堆
	return h[i] < h[j]
}
func (h IntHeap) Swap(i, j int)  {
	h[i], h[j] = h[j], h[i]
}
func (h *IntHeap) Pop() interface{}  {
	// 绑定pop方法，从最后拿出一个元素并返回
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
func (h *IntHeap) Push(x interface{}) {
	// 绑定push方法，插入新元素
	*h = append(*h, x.(int))
}
