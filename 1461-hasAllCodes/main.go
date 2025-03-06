package main

import (
	"fmt"
)

func hasAllCodes(s string, k int) bool {
	res := false
	count := 1 << k
	sMap := make(map[string]bool, count)
	window := string("")

	for i, _ := range s {
		window = window[:] + string(s[i])
		if i < k-1 {
			continue
		}
		sMap[window] = true
		window = window[1:]

	}
	if len(sMap) >= count {
		return true
	}
	return res
}

func main() {
	numbers := "00110110"
	res := hasAllCodes(numbers, 2)

	fmt.Println(res)
}
