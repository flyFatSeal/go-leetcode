package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func helper(root *TreeNode, l *[]int) {
	if root == nil {
		return
	}
	helper(root.Left, l)
	*l = append(*l, root.Val)
	helper(root.Right, l)
}

func minDiffInBST(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var l []int
	helper(root, &l)
	ret := math.MaxInt32
	for i := 1; i < len(l); i++ {
		if l[i]-l[i-1] < ret {
			ret = l[i] - l[i-1]
		}
	}
	return ret
}

func minDiffInBSTStack(root *TreeNode) int {
	var res []int
	stack := make([]*TreeNode, 0)
	min := 1<<63 - 1
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, node.Val)
		if len(res) > 1 {
			r := node.Val - res[len(res)-2]
			if min > r {
				min = r
			}
		}
		root = node.Right
	}
	if len(res) < 2 {
		return res[0]
	}
	return min
}
