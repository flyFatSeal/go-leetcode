package main

import (
	"fmt"
)

func maxArea(height []int) int {
	left := 0
	right := len(height) - 1

	res := 0
	for left < right {
		res = max(res, (right-left)*min(height[left], height[right]))
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	return res
}

func main() {
	numbers := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	res := maxArea(numbers)

	fmt.Println(res)
}
