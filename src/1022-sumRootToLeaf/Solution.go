package Solution

import "until"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumRootToLeaf(root *until.TreeNode) int {
	if root == nil {
		return 0
	}
	return _sumRootToLeaf(root, 0)
}

func _sumRootToLeaf(root *until.TreeNode, num int) int {
	var sum int
	num = (num << 1) + root.Val
	if root.Left == nil && root.Right == nil {
		return num
	}
	if root.Left != nil {
		sum += _sumRootToLeaf(root.Left, num)
	}
	if root.Right != nil {
		sum += _sumRootToLeaf(root.Right, num)
	}
	return sum
}
