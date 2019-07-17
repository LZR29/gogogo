package list_tree

import "math"

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	curA := headA
	curB := headB
	countA := 1
	countB := 1
	//比较下一个指针为空是为了统计个数后比较两个链表的最后一个节点是否相同来判断有没有相交
	for (*curA).Next != nil || (*curB).Next != nil {
		if (*curA).Next != nil {
			countA++
			curA = curA.Next
		}
		if (*curB).Next != nil {
			countB++
			curB = curB.Next
		}
	}
	if curA != curB {
		return nil
	}
	if countB > countA {
		headA, headB = headB, headA
	}
	dif := int(math.Abs(float64(countA - countB)))
	for i := 0; i < dif; i++ {
		headA = headA.Next
	}
	for headA != headB {
		headA = headA.Next
		headB = headB.Next
	}
	return headB
}
