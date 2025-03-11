package main

import (
	"fmt"
)

func checkInclusion(s1 string, s2 string) bool {

	s1Len := len(s1)
	s1Map := make(map[rune]int, s1Len)
	diff := s1Len
	window := ""
	for _, s := range s1 {
		s1Map[s] += 1
	}

	for i, s := range s2 {

		window = window[0:] + string(s)
		if _, ok := s1Map[s]; ok {
			s1Map[s] -= 1
			diff--
		}
		if diff == 0 {
			res := true
			for _, v := range s1Map {
				if v != 0 {
					res = false
					break
				}
			}
			if res {
				return true
			}

		}
		if i < s1Len-1 {
			continue
		}
		if _, ok := s1Map[rune(s2[i-s1Len+1])]; ok {
			diff++
			s1Map[rune(s2[i-s1Len+1])] += 1
		}

		window = window[1:]

	}

	return false
}

func checkInclusion1(s1 string, s2 string) bool {

	s1Len := len(s1)
	s2Len := len(s2)

	if s1Len > s2Len {
		return false
	}

	// 初始化字符计数数组
	s1Count := [26]int{}
	s2Count := [26]int{}

	for i := 0; i < s1Len; i++ {
		s1Count[s1[i]-'a']++
		s2Count[s2[i]-'a']++
	}

	// 检查初始窗口是否匹配
	if s1Count == s2Count {
		return true
	}

	// 滑动窗口
	for i := s1Len; i < s2Len; i++ {
		s2Count[s2[i]-'a']++
		s2Count[s2[i-s1Len]-'a']--

		if s1Count == s2Count {
			return true
		}
	}

	return false
}

func main() {
	s2 := "bbbca"
	s1 := "abc"

	res := checkInclusion(s1, s2)

	fmt.Println(res)
}
