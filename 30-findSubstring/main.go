package main

import (
	"fmt"
)

func findSubstring(s string, words []string) []int {
	if len(s) == 0 || len(words) == 0 {
		return []int{}
	}

	oneWord := len(words[0])
	wordNum := len(words)
	n := len(s)
	wordCount := make(map[string]int)
	for _, word := range words {
		wordCount[word]++
	}

	var res []int
	for i := 0; i < oneWord; i++ {
		curCount := 0
		left := i
		right := i
		curCounter := make(map[string]int)
		for right+oneWord <= n {
			w := s[right : right+oneWord]
			right += oneWord
			curCounter[w]++
			curCount++
			for curCounter[w] > wordCount[w] {
				leftW := s[left : left+oneWord]
				left += oneWord
				curCounter[leftW]--
				curCount--
			}
			if curCount == wordNum {
				res = append(res, left)
			}
		}
	}
	return res
}

func main() {
	s := "barfoothefoobarman"
	words := []string{"foo", "bar"}
	res := findSubstring(s, words)
	fmt.Println(res) // 输出: [0, 9]
}
