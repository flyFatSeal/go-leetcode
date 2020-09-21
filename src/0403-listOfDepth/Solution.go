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
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// 用队列实现层序遍历
func listOfDepth(tree *until.TreeNode) []*until.ListNode {
	if tree == nil {
		return []*until.ListNode{}
	}
	queue := []*until.TreeNode{tree}
	var res []*until.ListNode
	for len(queue) > 0 {
		// 出队 - 一次性把队列清空
		var levelList *until.ListNode = &until.ListNode{-1, &until.ListNode{}}
		var dummeyList *until.ListNode = levelList.Next
		var tempNodeQueue []*until.TreeNode
		for i, node := range queue {
			levelList.Next.Val = node.Val
			if i != len(queue)-1 {
				levelList.Next.Next = &until.ListNode{}
			}
			if node.Left != nil {
				tempNodeQueue = append(tempNodeQueue, node.Left)
			}
			if node.Right != nil {
				tempNodeQueue = append(tempNodeQueue, node.Right)
			}
			levelList = levelList.Next
		}
		res = append(res, dummeyList)
		queue = tempNodeQueue
	}
	return res
}

func main() {
	root := until.Add([]interface{}{1, 2, 3, 4, 5, nil, 7, 8})
	greaterRoot := listOfDepth(root)
	fmt.Print(greaterRoot)
}
