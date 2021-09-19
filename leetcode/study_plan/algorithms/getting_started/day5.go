package getting_started

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	newHead := &ListNode{-1, head}
	slow, fast := newHead, head
	for i := 0; i < n; i++ {
		fast = fast.Next
	}
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	//fmt.Println(slow)
	slow.Next = slow.Next.Next
	return newHead.Next
}
