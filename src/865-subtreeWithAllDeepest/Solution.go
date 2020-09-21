package Solution

import (
	"math"
	"until"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 层序遍历

// 方法一 自顶而下 由根节点开始左右进行剪枝，时间复杂度0(N^2)
func subtreeWithAllDeepest(root *until.TreeNode) *until.TreeNode {
	if root == nil {
		return root
	}

	hL := height(root.Left)
	hR := height(root.Right)

	if hL == hR {
		return root
	}
	if hL > hR {
		return subtreeWithAllDeepest(root.Left)
	}
	return subtreeWithAllDeepest(root.Right)
}

func height(root *until.TreeNode) int {
	if root == nil {
		return 0
	}
	return int(math.Max(float64(height(root.Left)), float64(height(root.Right)))) + 1
}

// 方法二 自下而上 通过后续遍历 由下向上递归 返回root  时间复杂度0(N)
func subtreeWithAllDeepestTwo(root *until.TreeNode) *until.TreeNode {
	if root == nil {
		return nil
	}

	var result *until.TreeNode
	max := 0
	helper(root, 1, &max, &result)

	return result
}

func helper(root *until.TreeNode, level int, max *int, result **until.TreeNode) int {
	if root == nil {
		return level - 1
	}

	left := helper(root.Left, level+1, max, result)
	right := helper(root.Right, level+1, max, result)

	*max = maxNum(*max, maxNum(left, right))
	if *max == left && *max == right {
		*result = root
	}

	return maxNum(left, right)
}

func maxNum(a, b int) int {
	if a > b {
		return a
	}
	return b
}
