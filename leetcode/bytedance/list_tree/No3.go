package list_tree

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	cur := head
	more := 0
	for l1 != nil || l2 != nil {
		if l1 != nil {
			more += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			more += l2.Val
			l2 = l2.Next
		}
		node := &ListNode{Val: more % 10}
		more /= 10
		cur.Next = node
		cur = cur.Next

	}
	if more != 0 {
		node := &ListNode{Val: more}
		cur.Next = node
		cur = cur.Next
	}
	return head.Next
}
