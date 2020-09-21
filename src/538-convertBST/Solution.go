package main

import (
	"fmt"
	"until"
)

// 累加就是由大到小  将中序遍历反过来 得到的就是由大到小
var sum int

func convertBST(root *until.TreeNode) *until.TreeNode {
	//先右子树，再求根节点的值，之后将根节点的值加到左子树上
	sum = 0
	helper(root)
	return root
}
func helper(root *until.TreeNode) {
	if root == nil {
		return
	}
	helper(root.Right)
	sum += root.Val
	root.Val = sum
	helper(root.Left)
	return
}

func main() {
	root := until.Add([]interface{}{2, 1, 3})
	greaterRoot := convertBST(root)
	fmt.Print(greaterRoot)
}
