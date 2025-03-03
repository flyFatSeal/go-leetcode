package main

import (
	"fmt"
)

func numOfSubarrays(arr []int, k int, threshold int) int {
	res := 0
	sum := 0

	for i, v := range arr {
		sum += v
		if i < k-1 {
			continue
		}
		if sum/k >= threshold {
			res++
		}
		sum -= arr[i-k+1]
	}

	return res

}

func main() {
	numbers := []int{2, 2, 2, 2, 5, 5, 5, 8}
	res := numOfSubarrays(numbers, 3, 4)

	fmt.Println(res)
}
