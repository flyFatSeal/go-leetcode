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
func zigzagLevelOrder(root *until.TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}
	reverse := true
	stack := []*until.TreeNode{root}
	for len(stack) > 0 {
		var storeOut []*until.TreeNode
		var childRes []int
		for i, node := range stack {
			if node.Left != nil {
				storeOut = append(storeOut, node.Left)
			}
			if node.Right != nil {
				storeOut = append(storeOut, node.Right)
			}

			if reverse {
				childRes = append(childRes, node.Val)
			} else {
				childRes = append(childRes, stack[len(stack)-1-i].Val)
			}
		}
		res = append(res, childRes)
		stack = storeOut
		reverse = !reverse
	}
	return res
}

func main() {
	testData := []interface{}{3, 9, 20, nil, nil, 15, 7}
	testTree := until.Add(testData)
	fmt.Print(zigzagLevelOrder(testTree))
}
