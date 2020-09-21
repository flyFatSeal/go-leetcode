package main

import (
	"fmt"
	"until"
)

// 中序遍历
func rangeSumBST(root *until.TreeNode, L int, R int) int {
	res := 0
	stack := []*until.TreeNode{}

	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		popNode := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if popNode.Val <= R && popNode.Val >= L {
			res += popNode.Val
		}
		root = popNode.Right
	}
	return res
}

func main() {
	testData := []interface{}{10, 5, 15, 3, 7, nil, 18}
	testTree := until.Add(testData)
	fmt.Print(rangeSumBST(testTree, 7, 15))
}
