package Solution

import "until"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 自顶而下递归解决问题  依次找到当前的根节点 然后接着进行下一层的寻找
// preorder[0] 前序遍历的第一个值肯定是根节点 根据这个值找到 处于中序遍历的位置 中序遍历 根节点的左边肯定是左子树 右边就是右子树
// 就这样一直递归 知道数组为空即可。
func buildTree(preorder []int, inorder []int) *until.TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}

	// 中顺序列找根结点
	var root int
	for k, v := range inorder {
		if v == preorder[0] {
			root = k
			break
		}
	}

	// 左右子树归类
	// pre_left, pre_right := preorder[1: root+1], preorder[root+1:]
	// in_left, in_right := inorder[0: root], inorder[root+1:]

	// 左右子树递归
	return &until.TreeNode{
		Val:   preorder[0],
		Left:  buildTree(preorder[1:root+1], inorder[0:root]),
		Right: buildTree(preorder[root+1:], inorder[root+1:]),
	}
}
