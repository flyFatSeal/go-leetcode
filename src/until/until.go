package until

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func Add(TreeArray []interface{}) *TreeNode {
	root := new(TreeNode)
	index := 0
	return CreatTree(root, TreeArray, index)
}

func CreatTree(tree *TreeNode, TreeArray []interface{}, index int) *TreeNode {
	if index > len(TreeArray)-1 || TreeArray[index] == nil {
		return tree
	}
	if tree == nil {
		tree = new(TreeNode)
	}
	tree.Val = TreeArray[index].(int)

	tree.Left = CreatTree(tree.Left, TreeArray, index*2+1)
	tree.Right = CreatTree(tree.Right, TreeArray, index*2+2)

	return tree

}

func helper(root *TreeNode, l *[]int) {
	if root == nil {
		return
	}
	helper(root.Left, l)
	*l = append(*l, root.Val)
	helper(root.Right, l)
}

func MinDiffInBST(root *TreeNode) []int {
	if root == nil {
		return []int{0}
	}
	var l []int
	helper(root, &l)
	return l
}
