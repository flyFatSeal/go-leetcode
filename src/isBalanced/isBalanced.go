/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

package isBalanced

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	return root == nil || checkChildIsSymmetric(root.Left, root.Right)
}

func checkChildIsSymmetric(Left *TreeNode, Right *TreeNode) bool {
	// 直到叶子节点时
	if Left == nil || Right == nil {
		return Left == Right
	}
	// 非叶子节点 先检查本身的值是否相等 再检查子树
	if Left.Val != Right.Val {
		return false
	}
	// 递归检查所有子节点
	return checkChildIsSymmetric(Left.Left, Right.Right) && checkChildIsSymmetric(Left.Right, Right.Left)

}
