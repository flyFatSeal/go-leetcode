package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	res := make([][]int, 0)
	x, y, z := 0, 0, 0
	for i, _ := range nums[0 : len(nums)-2] {
		x = nums[i]
		if i > 0 && x == nums[i-1] {
			continue
		}
		if x+nums[i+1]+nums[i+2] > 0 {
			break
		}
		if x+nums[len(nums)-1]+nums[len(nums)-2] < 0 {
			continue
		}
		y = i + 1
		z = len(nums) - 1
		for y < z {
			s := x + nums[y] + nums[z]
			if s > 0 {
				z--
			} else if s < 0 {
				y++
			} else {
				res = append(res, []int{x, nums[y], nums[z]})
				y++
				for y < z && nums[y] == nums[y-1] {
					y++
				}
				z--
				for z > y && nums[z] == nums[z+1] {
					z--
				}
			}

		}
	}

	return res
}

func main() {
	numbers := []int{-1, 0, 1, 2, -1, -4}
	res := threeSum(numbers)

	fmt.Println(res)
}
