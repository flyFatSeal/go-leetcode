package main

import (
	"fmt"
	"until"
)

/**
 * Definition for a binary tree node.
 * type until.TreeNode struct {
 *     Val int
 *     Left *until.TreeNode
 *     Right *until.TreeNode
 * }
 */

// 层序遍历 核心就是队列先进先出

func levelOrder(root *until.TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	queue := []*until.TreeNode{root}
	var res [][]int
	for len(queue) > 0 {
		// 出队 - 一次性把队列清空
		var level []int
		var tempNodeQueue []*until.TreeNode
		for _, v := range queue {
			node := v
			level = append(level, node.Val)
			if node.Left != nil {
				tempNodeQueue = append(tempNodeQueue, node.Left)
			}
			if node.Right != nil {
				tempNodeQueue = append(tempNodeQueue, node.Right)
			}
		}
		res = append(res, level)
		queue = tempNodeQueue
	}
	return res
}

func main() {
	root := until.Add([]interface{}{3, 9, 20, nil, nil, 15, 7})
	greaterRoot := levelOrder(root)
	fmt.Print(greaterRoot)
}
