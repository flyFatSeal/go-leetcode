package main

import (
	"fmt"
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
// 直接中序遍历查看是否符合从小到大即可

func isValidBST(root *until.TreeNode) bool {
	if root == nil {
		return false
	}
	last := math.MinInt32
	stack := make([]*until.TreeNode, 0)

	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		node := stack[len(stack)-1]
		if node.Val <= last {
			return false
		}
		last = node.Val
		stack = stack[:len(stack)-1]
		root = node.Right
	}
	return true
}

func main() {
	root := until.Add([]interface{}{2, 1, 3})
	greaterRoot := isValidBST(root)
	fmt.Print(greaterRoot)
}
