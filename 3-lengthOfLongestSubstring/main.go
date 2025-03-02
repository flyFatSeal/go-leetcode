package main

import (
	"fmt"
)

func lengthOfLongestSubstring(s string) int {
	res := 0
	left := 0
	count := make(map[rune]int, 128)

	for right, v := range s {
		count[v] += 1
		for count[v] > 1 {
			count[rune(s[left])] -= 1
			left++
		}
		res = max(res, right-left+1)
	}
	return res
}

func main() {
	numbers := "abcabcbb"
	res := lengthOfLongestSubstring(numbers)

	fmt.Println(res)
}
