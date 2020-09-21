package main

import (
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
func isSubStructure(A *until.TreeNode, B *until.TreeNode) bool {
	res := false
	if B == nil {
		return res
	}
	stack := make([]*until.TreeNode, 0)
	// 前序遍历
	for A != nil || len(stack) > 0 {
		for A != nil {
			stack = append(stack, A)
			A = A.Left
		}
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if node.Val == B.Val {
			res = compare(node, B)
			if res {
				return res
			}
		}
		A = node.Right
	}
	return res
}

func compare(A *until.TreeNode, B *until.TreeNode) bool {
	if B != nil && A == nil {
		return false
	}
	if B == nil && A == nil || B == nil && A != nil {
		return true
	}
	if B.Val != A.Val {
		return false
	}
	l := compare(A.Left, B.Left)
	r := compare(A.Right, B.Right)

	return l && r
}
