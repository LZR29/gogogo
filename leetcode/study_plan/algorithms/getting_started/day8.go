package getting_started

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	if t1 == nil {
		return t2
	}
	if t2 == nil {
		return t1
	}
	node := TreeNode{Val:t1.Val + t2.Val}
	node.Left = mergeTrees(t1.Left, t2.Left)
	node.Right = mergeTrees(t1.Right, t2.Right)
	return &node
}
*/
func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil && root2 == nil {
		return nil
	}
	rootVal := 0
	if root1 != nil {
		rootVal += root1.Val
	}
	if root2 != nil {
		rootVal += root2.Val
	}
	newRoot := &TreeNode{Val: rootVal}
	merge(root1, root2, newRoot)
	return newRoot
}

func merge(node1 *TreeNode, node2 *TreeNode, root *TreeNode) {
	leftVal, rightVal := 0, 0
	leftTag, rightTag := false, false
	if node1 != nil && node1.Left != nil {
		leftTag = true
		leftVal += node1.Left.Val
	}
	if node1 != nil && node1.Right != nil {
		rightTag = true
		rightVal += node1.Right.Val
	}
	if node2 != nil && node2.Left != nil {
		leftTag = true
		leftVal += node2.Left.Val
	}
	if node2 != nil && node2.Right != nil {
		rightTag = true
		rightVal += node2.Right.Val
	}
	if leftTag {
		root.Left = &TreeNode{Val: leftVal}
		if node1 != nil && node2 != nil {
			merge(node1.Left, node2.Left, root.Left)
		}
		if node1 == nil && node2 != nil {
			merge(nil, node2.Left, root.Left)
		}
		if node1 != nil && node2 == nil {
			merge(node1.Left, nil, root.Left)
		}
	}
	if rightTag {
		root.Right = &TreeNode{Val: rightVal}
		if node1 != nil && node2 != nil {
			merge(node1.Right, node2.Right, root.Right)
		}
		if node1 == nil && node2 != nil {
			merge(nil, node2.Right, root.Right)
		}
		if node1 != nil && node2 == nil {
			merge(node1.Right, nil, root.Right)
		}
	}
}

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

//层序遍历
func connect(root *Node) *Node {
	if root == nil {
		return root
	}
	queue := []*Node{root}
	for len(queue) > 0 {
		temp := queue
		queue = nil
		for i, val := range temp {
			if i+1 < len(temp) {
				val.Next = temp[i+1]
			}
			if val.Left != nil {
				queue = append(queue, val.Left)
			}
			if val.Right != nil {
				queue = append(queue, val.Right)
			}
		}
	}
	return root
}
