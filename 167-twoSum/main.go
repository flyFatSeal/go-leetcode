package main

import "fmt"

func twoSum(numbers []int, target int) []int {
	left := 0
	right := len(numbers) - 1

	res := make([]int, 2)
	for left < right {
		s := numbers[right] + numbers[left]
		if s > target {
			right--
		}
		if s < target {
			left++
		}
		if s == target {
			break
		}
	}

	res[0] = left + 1
	res[1] = right + 1

	return res
}

func main() {
	numbers := []int{2, 7, 11, 15}
	res := twoSum(numbers, 9)

	fmt.Println(res)
}
