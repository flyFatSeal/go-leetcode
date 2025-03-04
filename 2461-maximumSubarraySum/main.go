package main

import (
	"fmt"
)

func maximumSubarraySum(nums []int, k int) int64 {
	res := 0
	sum := 0
	elementCount := make(map[int]int)

	for i, v := range nums {
		sum += v
		elementCount[v]++
		if i < k-1 {
			continue
		}
		if len(elementCount) >= k {
			res = max(res, sum)
		}
		remove := nums[i-k+1]
		sum -= remove
		elementCount[remove]--
		if elementCount[remove] == 0 {
			delete(elementCount, remove)
		}
	}

	return int64(res)
}

func main() {
	numbers := []int{1, 5, 4, 2, 9, 9, 9}
	res := maximumSubarraySum(numbers, 3)

	fmt.Println(res)
}
