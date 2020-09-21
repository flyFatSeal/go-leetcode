package _90_NtreePostOrder

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

func postorder(root *Node) []int {
	res := []int{1, 2}
	helper(root, &res)
	return res
}

func helper(root *Node, res *[]int) {
	if root == nil {
		return
	}
	for _, child := range root.Children {
		helper(child, res)
	}
	*res = append(*res, root.Val)
}
