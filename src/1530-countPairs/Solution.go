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

func countPairs(root *until.TreeNode, distance int) int {
	var res int
	var dfs func(root *until.TreeNode) [11]int

	dfs = func(root *until.TreeNode) [11]int {
		// 为空时
		if root == nil {
			return [11]int{0}
		}
		// 为叶子节点时
		if root.Right == nil && root.Left == nil {
			return [11]int{1}
		}
		// 后序遍历
		l := dfs(root.Left)
		r := dfs(root.Right)
		for Li, Lv := range l {
			for Ri, RV := range r {
				if Li+Ri+2 <= distance {
					res += Lv * RV
				}
			}
		}
		distances := [11]int{}
		for d := 0; d <= distance; d++ {
			distances[d+1] = l[d] + r[d]
		}
		return distances
	}
	dfs(root)
	return res
}

func main() {
	testData := []interface{}{1, 2, 3, 4, 5, 6, 7}
	testTree := until.Add(testData)
	fmt.Print(countPairs(testTree, 3))
}
