package list_tree

func merge2Lists(l1 *ListNode, l2 *ListNode) *ListNode {
	res := &ListNode{}
	cur := res
	for l1 != nil || l2 != nil {
		if l1 == nil {
			cur.Next = l2
			l2 = nil
			continue
		}
		if l2 == nil {
			cur.Next = l1
			l1 = nil
			continue
		}
		if l1.Val > l2.Val {
			cur.Next = l2
			cur, l2 = cur.Next, l2.Next
		} else {
			cur.Next = l1
			cur, l1 = cur.Next, l1.Next
		}
	}
	return res.Next
}

//每次合并一对链表，对合并完的链表重复执行
func mergeKLists(lists []*ListNode) *ListNode {
	length := len(lists)
	i := 1
	for i < length {
		for j := 0; j < length-i; j += i * 2 {
			lists[j] = merge2Lists(lists[j], lists[j+i])
		}
		i *= 2
	}
	if length > 0 {
		return lists[0]
	}
	return nil
}
