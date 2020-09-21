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

// 用两个栈，一个栈是被删节点集合，一个是用于遍历的栈
func delNodes(root *TreeNode, toDelete []int) []*TreeNode {
	if root == nil {
		return nil
	}
	res, deleteMap := []*TreeNode{}, map[int]bool{}
	for _, v := range toDelete {
		deleteMap[v] = true
	}
	dfsDelNodes(root, deleteMap, true, &res)
	return res
}

func dfsDelNodes(root *TreeNode, toDel map[int]bool, isRoot bool, res *[]*TreeNode) bool {
	if root == nil {
		return false
	}
	if isRoot && !toDel[root.Val] {
		*res = append(*res, root)
	}
	isRoot = false
	if toDel[root.Val] {
		isRoot = true
	}
	if dfsDelNodes(root.Left, toDel, isRoot, res) {
		root.Left = nil
	}
	if dfsDelNodes(root.Right, toDel, isRoot, res) {
		root.Right = nil
	}
	return isRoot
}
func main() {
	root := until.Add([]interface{}{1, 2, 3, 4, 5, 6, 7})
	greaterRoot := delNodes(root, []int{3, 5})
	fmt.Print(greaterRoot)
}
