package main

import (
	"fmt"
)

func trap(height []int) int {
	res := 0

	pre_max := make([]int, len(height))
	sub_max := make([]int, len(height))

	pre_max[0] = height[0]

	for i, v := range height[1:] {
		pre_max[i+1] = max(pre_max[i], v)
	}

	sub_max[len(sub_max)-1] = height[len(height)-1]

	for i := len(height) - 2; i >= 0; i-- {
		sub_max[i] = max(height[i], sub_max[i+1])
	}

	for i, v := range height {
		res += min(pre_max[i], sub_max[i]) - v
	}

	return res
}

func trap2(height []int) int {
	res := 0

	left := 0
	right := len(height) - 1
	pre_max := 0
	sub_max := 0

	for left < right {
		pre_max = max(pre_max, height[left])
		sub_max = max(sub_max, height[right])

		if pre_max < sub_max {
			res += pre_max - height[left]
			left++
		} else {
			res += sub_max - height[right]
			right--
		}
	}

	return res
}

func main() {
	numbers := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	res := trap(numbers)
	fmt.Println(res)
	res2 := trap2(numbers)
	fmt.Println(res2)

}
