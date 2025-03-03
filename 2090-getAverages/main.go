package main

import (
	"fmt"
)

func getAverages(nums []int, k int) []int {
	res := make([]int, len(nums))
	for i := range res {
		res[i] = -1
	}
	sum := 0

	for i, v := range nums {
		sum += v
		if i < k*2 {
			continue
		}
		res[i-k] = sum / (2*k + 1)

		sum -= nums[i-k*2]

	}

	return res
}

func main() {
	numbers := []int{100000}
	res := getAverages(numbers, 0)

	fmt.Println(res)
}
