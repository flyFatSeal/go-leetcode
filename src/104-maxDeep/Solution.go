package Solution

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	return getDepth(root, 0)
}

func getDepth(node *TreeNode, upDeep int) int {
	if node == nil {
		return upDeep
	}
	upDeep += 1

	leftChildDeep := getDepth(node.Left, upDeep)
	rightChildDeep := getDepth(node.Right, upDeep)
	if leftChildDeep > rightChildDeep {
		return leftChildDeep
	}
	return rightChildDeep
}
