package easy

func rangeSumBST(root *TreeNode, L int, R int) int {
	ans := 0
	rangeSumBSTHepler(root, L, R, &ans)
	return ans
}
func rangeSumBSTHepler(root *TreeNode, L int, R int, ans *int) {
	if root == nil {
		return
	}
	if root.Val > R {
		rangeSumBSTHepler(root.Left, L, R, ans)
	} else if root.Val < L {
		rangeSumBSTHepler(root.Right, L, R, ans)
	} else {
		*ans += root.Val
		rangeSumBSTHepler(root.Left, L, R, ans)
		rangeSumBSTHepler(root.Right, L, R, ans)
	}

}
