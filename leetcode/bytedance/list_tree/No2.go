package list_tree
func reverseList(head *ListNode) *ListNode {
	var newHead, node  *ListNode
	for head != nil {
		node = head.Next
		head.Next = newHead
		newHead = head
		head = node
	}
	return newHead
}
