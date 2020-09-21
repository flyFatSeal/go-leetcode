package main

import "fmt"

// 假设 G(n) 为 从1...n为节点组成的二叉搜索树的总树 则设f(i) 为以i为节点 组成二叉搜索树的总数
// 则 f(i)  左边的树值区间为 [1,2....,i-1] 右边的值区间为 [i+1,i+2,...,n] 也就是说 f(i) 的值为左边区间G(left) * G(right)
// f(i) = G(i-1)*G(n-i)
// 而G(n) = f(1) + f(2) + ... f(n)
// 则 G(n) = G(0)*G(n-1) + G(1)*G(n-2) + G(2)*G(n-3) + ... + G(n-1)*G(0)
// G(0) = 0 G(1) = 1 G(2) = 2  G(3) = G(0)*G(2)+G(1)*G(1) +G(2)*G(0)
func numTrees(n int) int {
	// 给dp赋初值
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1
	for i := 2; i <= n; i++ {
		sum := 0
		for k := 1; k <= i; k++ {
			sum += dp[k-1] * dp[i-k]
		}
		dp[i] = sum
	}
	return dp[len(dp)-1]
}

func main() {
	fmt.Print(numTrees(3))
}
