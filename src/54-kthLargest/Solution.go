package main

import (
	"fmt"
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
// 反转中序遍历即可
func kthLargest(root *until.TreeNode, k int) int {
	if root == nil {
		return 0
	}
	stack := make([]*until.TreeNode, 0)
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Right
		}
		k = k - 1
		popNode := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if k == 0 {
			return popNode.Val
		}
		root = popNode.Left
	}
	return 0
}

func main() {
	testData := []interface{}{3, 1, 4, nil, 2}
	testTree := until.Add(testData)
	fmt.Print(kthLargest(testTree, 1))
}
