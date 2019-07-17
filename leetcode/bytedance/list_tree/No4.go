package list_tree

import (
	"sort"
)

func sortList(head *ListNode) *ListNode {
	var nums []int
	length := 0
	for head != nil {
		nums = append(nums, head.Val)
		head = head.Next
		length++
	}
	sort.Ints(nums)
	//	fmt.Println(nums)
	var newHead ListNode
	cur := &newHead
	for i := 0; i < length; i++ {
		cur.Next = &ListNode{Val: nums[i]}
		cur = cur.Next
	}
	return newHead.Next
}

//以下为别人的归并排序
func sortList2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	slow, fast := head, head
	var slowPre *ListNode //慢指针的前一个

	for fast != nil && fast.Next != nil {
		slowPre, slow = slow, slow.Next
		fast = fast.Next.Next
	}

	slowPre.Next = nil

	left := sortList(head)
	right := sortList(slow)

	return merge(left, right)
}

//归并两条排序好的链表
func merge(l1, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	var root, cur *ListNode
	if l1.Val <= l2.Val {
		root, cur = l1, l1
		l1 = l1.Next
	} else {
		root, cur = l2, l2
		l2 = l2.Next
	}

	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			cur.Next = l1
			l1 = l1.Next
		} else {
			cur.Next = l2
			l2 = l2.Next
		}
		cur = cur.Next
	}

	if l1 != nil {
		cur.Next = l1
	}
	if l2 != nil {
		cur.Next = l2
	}

	return root
}
