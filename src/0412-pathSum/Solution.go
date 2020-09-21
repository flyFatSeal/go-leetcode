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
func pathSum(root *until.TreeNode, sum int) int {
	if root == nil {
		return 0
	}
	var res int
	var dfs func(*until.TreeNode, int)
	dfs = func(root *until.TreeNode, sum int) {
		if root == nil {
			return
		}
		if sum == root.Val {
			res = res + 1
		}
		dfs(root.Left, sum-root.Val)
		dfs(root.Right, sum-root.Val)
	}

	stack := []*until.TreeNode{root}
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		dfs(node, sum)
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
		if node.Right != nil {
			stack = append(stack, node.Right)
		}

	}

	return res

}

func main() {
	root := until.Add([]interface{}{1, -2, -3, 1, 3, -2, nil, -1})
	greaterRoot := pathSum(root, -1)
	fmt.Print(greaterRoot)
}
