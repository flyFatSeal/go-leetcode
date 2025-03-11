package main

import (
	"fmt"
)

func findAnagrams(s string, p string) []int {
	plen := len(p)
	slen := len(s)
	res := []int{}
	s1 := [26]int{}
	s2 := [26]int{}

	if plen > slen {
		return res
	}

	for i, v := range p {
		s1[v-'a']++
		s2[s[i]-'a']++
	}
	if s1 == s2 {
		res = append(res, 0)
	}

	for i := plen; i < slen; i++ {
		s2[s[i-plen]-'a']--
		s2[s[i]-'a']++
		if s1 == s2 {
			res = append(res, i-plen+1)
		}
	}
	return res

}

func main() {
	s2 := "cbaebabacd"
	s1 := "abc"

	res := findAnagrams(s2, s1)

	fmt.Println(res)
}
