package easy

//这个节点就是要删除的节点了
//这个直接赋值过去就行了
func deleteNode(node *ListNode) {
	next := node.Next
	node.Next = next.Next
	node.Val = next.Val
}
