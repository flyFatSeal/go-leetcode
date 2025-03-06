package main

import (
	"fmt"
	"math"
)

func minSwaps(nums []int) int {
	res := math.MaxInt
	swap := 0
	window := 0

	right := 0

	for _, v := range nums {
		if v == 1 {
			window++
		}
	}

	if window == 0 {
		return 0
	}

	for right < len(nums)+window-1 {
		if nums[right%len(nums)] == 0 {
			swap++
		}
		if right < window-1 {
			right++
			continue
		}
		res = min(res, swap)
		if nums[right-window+1] == 0 {
			swap--
		}
		right++

	}

	return res

}

func main() {
	numbers := []int{0, 0, 0}
	res := minSwaps(numbers)

	fmt.Println(res)
}
