package contest

//5649. 解码异或后的数组
// a^b^b=a , a^b=c => a^b^b=c^b => a=b^c;若a^b=c,则a=b^c
func decode(encoded []int, first int) []int {
	ret := make([]int, len(encoded)+1)
	ret[0] = first
	for i, v := range encoded {
		ret[i+1] = v ^ ret[i]
	}
	return ret
}

type ListNode struct {
	Val  int
	Next *ListNode
}

//5652. 交换链表中的节点
func swapNodes(head *ListNode, k int) *ListNode {
	fast := head
	slow := head
	positive := head
	for i := 1; i < k && fast.Next != nil; i++ {
		fast = fast.Next
	}
	positive = fast
	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}
	positive.Val, slow.Val = slow.Val, positive.Val
	return head
}
