package list_tree

func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	help := helper(head)
	if help == nil {
		return nil
	}
	p1 := help
	p2 := head
	for p1 != p2 {
		p1 = p1.Next
		p2 = p2.Next
	}
	return p1
}
func helper(head *ListNode) *ListNode {
	fast := head
	slow := head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if slow == fast {
			return slow
		}
	}
	return nil
}
