package Solution

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	// root 处于两者之间肯定是最近公共节点
	if (root.Val >= p.Val && root.Val <= q.Val) || (root.Val <= p.Val && root.Val >= q.Val) {
		return root
	}

	if root.Val < q.Val && root.Val < p.Val {
		return lowestCommonAncestor(root.Right, p, q)
	}

	return lowestCommonAncestor(root.Left, p, q)
}
