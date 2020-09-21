package Solution

import (
	"strconv"
	"strings"
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
//48ms
func tree2str(t *until.TreeNode) string {
	if t == nil {
		return ""
	}
	var res string
	var dfs func(root *until.TreeNode)
	dfs = func(root *until.TreeNode) {
		if root == nil {
			return
		}
		res += "(" + strconv.Itoa(root.Val)
		if root.Left == nil && root.Right != nil {
			res += "()"
		}
		dfs(root.Left)
		dfs(root.Right)

		res += ")"
	}
	dfs(t)

	return res[1 : len(res)-1]
}

// 0ms
func tree2str2(t *until.TreeNode) string {
	if t == nil {
		return ""
	}
	if t.Left == nil && t.Right == nil {
		return strconv.Itoa(t.Val)
	}

	var strBuilder strings.Builder

	strBuilder.WriteString(strconv.Itoa(t.Val))
	//左边无论如何必写()
	strBuilder.WriteString("(")
	strBuilder.WriteString(tree2str(t.Left))
	strBuilder.WriteString(")")

	if t.Right != nil {
		//右边不为空才写()
		strBuilder.WriteString("(")
		strBuilder.WriteString(tree2str(t.Right))
		strBuilder.WriteString(")")
	}

	return strBuilder.String()
}
