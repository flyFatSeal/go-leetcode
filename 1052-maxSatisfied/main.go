package main

import (
	"fmt"
	"math"
)

func maxSatisfied(customers []int, grumpy []int, minutes int) int {
	total := 0
	satisfied := -math.MaxInt
	satisfied_total := 0

	unsatisfied_total := 0

	for i, v := range grumpy {
		total += customers[i]
		if v == 1 {
			satisfied_total += customers[i]
			unsatisfied_total += customers[i]
		}
		if i < minutes-1 {
			continue
		}
		satisfied = max(satisfied, satisfied_total)

		if grumpy[i-minutes+1] == 1 {
			satisfied_total -= customers[i-minutes+1]
		}

	}

	return total - unsatisfied_total + satisfied
}

func main() {
	numbers := []int{7, 8, 8, 6}
	grumpy := []int{0, 1, 0, 1}
	res := maxSatisfied(numbers, grumpy, 3)

	fmt.Println(res)
}
