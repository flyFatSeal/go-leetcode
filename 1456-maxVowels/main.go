package main

import (
	"fmt"
)

func maxVowels(s string, k int) int {
	res := 0
	vowel := 0
	for i, v := range s {
		if v == 'a' || v == 'e' || v == 'i' || v == 'o' || v == 'u' {
			vowel++
		}
		res = max(res, vowel)
		if i < k-1 {
			continue
		}
		if s[i-k+1] == 'a' || s[i-k+1] == 'e' || s[i-k+1] == 'i' || s[i-k+1] == 'o' || s[i-k+1] == 'u' {
			vowel--
		}
	}

	return res
}

func main() {
	numbers := "abciiidef"
	res := maxVowels(numbers, 2)

	fmt.Println(res)
}
