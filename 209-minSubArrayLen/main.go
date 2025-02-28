package main

import (
	"fmt"
)

func minSubArrayLen(target int, nums []int) int {

	res := len(nums) + 1
	sum := 0
	left := 0

	for right, v := range nums {
		sum += v
		for sum >= target {
			res = min(res, right-left+1)
			sum -= nums[left]
			left++
		}
	}

	if res == len(nums)+1 {
		return 0
	}

	return res
}

func main() {
	numbers := []int{2, 3, 1, 2, 4, 3}
	res := minSubArrayLen(7, numbers)

	fmt.Println(res)
}
