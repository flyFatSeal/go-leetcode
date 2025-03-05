package main

import (
	"fmt"
)

func decrypt(code []int, k int) []int {
	len := len(code)
	res := make([]int, len)

	for i, _ := range code {
		left := i
		right := (i + k + len) % len
		decode := 0

		for left != right {
			index := left
			if k > 0 {
				index = (left + 1) % len
			} else {
				index = (left - 1 + len) % len
			}
			decode += code[index]
			left = index
		}
		res[i] = decode
	}

	return res

}

func main() {
	numbers := []int{5, 7, 1, 4}
	res := decrypt(numbers, 3)

	fmt.Println(res)
}
