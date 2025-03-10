package main

import (
	"fmt"
)

func minFlips(s string) int {
	n := len(s)
	ans := n
	// 枚举开头是 0 还是 1
	for head := byte('0'); head <= '1'; head++ {
		// 左边每个位置的不同字母个数
		leftDiff := make([]int, n)
		diff := 0
		for i := range s {
			if s[i] != head^byte(i&1) {
				diff++
			}
			leftDiff[i] = diff
		}

		// 右边每个位置的不同字母个数
		tail := head ^ 1
		diff = 0
		for i := n - 1; i >= 0; i-- {
			// 左边+右边即为整个字符串的不同字母个数，取最小值5
			ans = min(ans, leftDiff[i]+diff)
			if s[i] != tail^byte((n-1-i)&1) {
				diff++
			}
		}
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	numbers := "1011010"
	res := minFlips(numbers)

	fmt.Println(res)
}
