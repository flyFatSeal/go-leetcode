package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	res := []int{}
	traversal(root, &res)

	return res
}

func traversal(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}
	*res = append(*res, root.Val)
	traversal(root.Left, res)
	traversal(root.Right, res)
}
