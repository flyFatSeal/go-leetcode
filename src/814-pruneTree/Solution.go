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

// 后续遍历 为0的叶子节点去除
func pruneTree(root *until.TreeNode) *until.TreeNode {
	if root == nil {
		return nil
	}

	root.Left = pruneTree(root.Left)
	root.Right = pruneTree(root.Right)

	if root.Right == nil && root.Left == nil && root.Val == 0 {
		return nil
	}

	return root
}

func main() {
	testData := []interface{}{1, 0, 1, 0, 0, 0, 1}
	testTree := until.Add(testData)
	fmt.Print(pruneTree(testTree))
}
