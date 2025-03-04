package main

import (
	"fmt"
)

func maxScore(cardPoints []int, k int) int {
	n := len(cardPoints)
	m := n - k
	s := 0
	for _, x := range cardPoints[:m] {
		s += x
	}
	total := s
	minS := s
	for i := m; i < n; i++ {
		total += cardPoints[i]
		s += cardPoints[i] - cardPoints[i-m]
		minS = min(minS, s)
	}
	return total - minS
}

func main() {
	numbers := []int{9, 7, 7, 9, 7, 7, 9}
	res := maxScore(numbers, 7)

	fmt.Println(res)
}
