package main

import (
	"fmt"
)

func hasAllCodes(s string, k int) bool {
	// 如果字符串长度小于k，无法形成长度为k的子串，返回false
	if len(s) < k {
		return false
	}

	// 创建一个布尔数组来记录所有可能的k位二进制数
	// 1<<k 表示2的k次方，即所有可能的k位二进制数的数量
	seen := make([]bool, 1<<k)
	total := 1 << k // 需要找到的所有可能的二进制串的数量
	count := 0      // 已找到的不同二进制串的数量

	// 创建一个k位的掩码，用于保持最后k位
	// 例如k=3时，mask=111(二进制)
	mask := total - 1

	// 初始化第一个窗口的值
	val := 0
	for i := 0; i < k; i++ {
		// 将每一位数字转换为二进制并拼接
		// s[i]-'0' 将字符转换为数字
		val = (val << 1) | int(s[i]-'0')
	}
	seen[val] = true // 标记第一个子串已经出现
	count++

	// 使用滑动窗口遍历剩余的子串
	for i := k; i < len(s); i++ {
		// 通过位运算更新窗口：
		// 1. val << 1 左移一位
		// 2. & mask 保持最后k位
		// 3. | int(s[i]-'0') 添加新的最低位
		val = ((val << 1) & mask) | int(s[i]-'0')
		if !seen[val] {
			// 如果这个子串之前没有出现过
			seen[val] = true
			count++
		}
	}

	// 检查是否找到了所有可能的k位二进制串
	return count == total
}

func main() {
	numbers := "00110110"
	res := hasAllCodes(numbers, 2)

	fmt.Println(res)
}
