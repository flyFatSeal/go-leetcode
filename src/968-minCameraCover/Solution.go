package Solution

import "math"

type TreeNode struct {
     Val int
    Left *TreeNode
     Right *TreeNode
 }

// 这是一道树结构的动态规划题，每个节点可以为三种状态 0 无摄像机 被子节点监控 1有摄像机 2 无摄像机被父节点监控
// 即用dp[root][0] dp[root][1] dp[root][2] 表示三个状态下各自的最优解  为什么不用dp[root] 直接表示 父节点的最优解 因为dp[root]
// 的最优解是会被子节点不同状态所影响的。
// 思路采取自下向上的解法 用后续遍历 依次由子节点返回给父节点
// 状态转移方程为
// dp[root][0] = dp[root.left][1]+dp[root.right][1]  -- 当前节点无摄像机 为了保证每个都被监控 所以各自的子节点必须要有摄像机 也就是dp[child][1]
// dp[root][1] = 1 + min(dp[root.left][0],[1],[2]) + min(dp[root.right][0],[1],[2])
// dp[root][2] = dp[root.left][1]+dp[root.right][1]
// 定义dp初始状态 dp[leaf][0] = int_max dp[leaf][1] = 1 dp[leaf][2] = 0
// dp[null] = 0

func min(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func _min(a int, b int, c int) int {
	return min(min(a, b), c)
}

func minCameraCover(root *TreeNode) int {
	res := _minCameraCover(root)

	return _min(res[0],res[1],res[2])

}

func _minCameraCover(root *TreeNode) []int{
	var dp = make([]int,3)
	if root == nil {

		dp[0] = 0
		dp[1] = 0
		dp[2] = 0
		return dp
	}
	if root.Left == nil && root.Right == nil {

		dp[0] = math.MaxInt32
		dp[1] = 1
		dp[2] = 0
		return dp
	}
	l := _minCameraCover(root.Left)
	r := _minCameraCover(root.Right)
	// 后序遍历
	dp[0] = min(r[0]+l[1],r[1]+l[0])
	dp[1] = 1 + _min(l[0],l[1],l[2]) + _min(r[0],r[1],r[2])
	dp[2] = min(l[0],l[1]) + min(r[0],r[1])

	return dp
}
