// copilot: disable

package main

import (
	"fmt"
	"math"
)

func maxFreq(s string, maxLetters int, minSize int, maxSize int) int {
	freq := make(map[string]int)
	windowFreq := make(map[string]int, minSize)
	window := ""

	res := 0

	for i, s := range s {
		windowFreq[string(s)]++
		window = window + string(s)
		if i < minSize-1 {
			continue
		}
		if len(windowFreq) <= maxLetters {
			freq[window] += 1
		}
		windowFreq[string(window[0])]--
		if windowFreq[string(window[0])] == 0 {
			delete(windowFreq, string(window[0]))
		}
		window = window[1:]
	}

	for _, v := range freq {
		res = max(res, v)
	}

	return res
}

func maxFreq2(s string, maxLetters int, minSize int, maxSize int) int {
	if len(s) < minSize {
		return 0
	}

	diff := 0
	v := [26]int{}

	arr := []rune(s)

	m := make(map[string]int)

	l := 0
	for i, c := range s {
		if v[c-'a'] == 0 {
			diff++
		}
		v[c-'a']++

		if i-l+1 < minSize {
			continue
		}

		if diff <= maxLetters {
			m[s[l:i+1]]++
		}

		if v[arr[l]-'a'] == 1 {
			diff--
		}
		v[arr[l]-'a']--
		l++
	}

	res := math.MinInt64
	for _, v := range m {
		if v > res {
			res = v
		}
	}

	if res == math.MinInt64 {
		return 0
	}

	return res
}

func main() {
	numbers := "aaaa"
	res := maxFreq(numbers, 1, 3, 4)

	fmt.Println(res)
}
