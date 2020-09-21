package main

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

type Node struct {
	Val      int
	Children []*Node
}

// 栈-前序遍历
func preorder(root *Node) []int {
	var res []int
	if root == nil {
		return res
	}
	stack := []*Node{root}
	var popNode *Node
	for len(stack) > 0 {
		// 出栈
		popNode = stack[len(stack)-1]
		res = append(res, popNode.Val)
		stack = stack[:len(stack)-1]
		root = popNode

		// 子节点压栈 - 压一次出一次
		if root != nil {
			for index := len(root.Children) - 1; index >= 0; index-- {
				stack = append(stack, root.Children[index])
			}
		}
	}
	return res
}

func main() {
	NTree := &Node{1, []*Node{}}
	next := &Node{3, []*Node{&Node{5, []*Node{}}, &Node{6, []*Node{}}}}
	NTree.Children = append(NTree.Children, next)
	NTree.Children = append(NTree.Children, &Node{2, []*Node{}})
	NTree.Children = append(NTree.Children, &Node{4, []*Node{}})

	preorder(NTree)
}
