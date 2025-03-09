package main

import (
	"fmt"
)

func maxFreeTime(eventTime, k int, startTime, endTime []int) (ans int) {
	n := len(startTime)
	get := func(i int) int {
		if i == 0 {
			return startTime[0]
		}
		if i == n {
			return eventTime - endTime[n-1]
		}
		return startTime[i] - endTime[i-1]
	}

	s := 0
	for i := range n + 1 {
		s += get(i)
		if i < k {
			continue
		}
		ans = max(ans, s)
		s -= get(i - k)
	}
	return
}

func main() {
	start := []int{0, 2, 9}
	end := []int{1, 4, 10}
	res := maxFreeTime(10, 1, start, end)

	fmt.Println(res)
}
