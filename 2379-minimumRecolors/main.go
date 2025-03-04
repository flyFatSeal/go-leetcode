package main

import (
	"fmt"
	"math"
)

func minimumRecolors(blocks string, k int) int {
	res := math.MaxInt
	white_sum := 0

	for i, v := range blocks {
		if v == 'W' {
			white_sum++
		}
		if i < k-1 {
			continue
		}
		res = min(res, white_sum)
		if blocks[i-k+1] == 'W' {
			white_sum--
		}
	}
	return res

}

func main() {
	numbers := "WBBWWBBWBW"
	res := minimumRecolors(numbers, 7)

	fmt.Println(res)
}
