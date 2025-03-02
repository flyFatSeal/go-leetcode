package main

import (
	"fmt"
)

func numSubarrayProductLessThanK(nums []int, k int) int {
	res := 0
	prod := 1
	left := 0
	if k <= 1 {
		return res
	}

	for right, v := range nums {
		prod *= v
		for prod >= k {
			prod = prod / nums[left]
			left++
		}
		res += right - left + 1
	}

	return res

}

func main() {
	numbers := []int{10, 5, 2, 6}
	res := numSubarrayProductLessThanK(numbers, 100)

	fmt.Println(res)
}
