package main

import (
	"fmt"
	"until"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func increasingBST(root *until.TreeNode) *until.TreeNode {
	if root == nil {
		return nil
	}
	stack := make([]*until.TreeNode, 0)
	res := new(until.TreeNode)
	dummeyHead := new(until.TreeNode)
	dummeyHead.Right = res
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		popNode := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res.Val = popNode.Val
		if root != nil || len(stack) > 0 || popNode.Right != nil {
			res.Right = new(until.TreeNode)
		}
		res = res.Right
		root = popNode.Right
	}
	return dummeyHead.Right
}

func main() {
	testData := []interface{}{5}
	testTree := until.Add(testData)
	fmt.Print(increasingBST(testTree))
}
