package main

import (
	"fmt"
)

func main() {
	s := "abcabcbb"
	fmt.Println(lengthOfLongestSubstring(s))
}

//func lengthOfLongestSubstring(s string) int {
//	max := 0
//	i := 0
//	for i < len(s) {
//		flag := 0
//		m := make(map[uint8]int)
//		for j := i; j < len(s); j++ {
//			if find, ok := m[s[j]]; ok {
//				flag = 1
//				i = find + 1
//				break
//			}
//			m[s[j]] = j
//		}
//		if len(m) > max {
//			max = len(m)
//		}
//		if flag == 0 {
//			i++
//		}
//	}
//	return max
//}

//func lengthOfLongestSubstring(s string) int {
//	max := 0
//	i := 0
//	j := 0
//	m := make(map[uint8]int)
//	for i < len(s) && j < len(s) {
//		if _, ok := m[s[j]]; !ok {
//			m[s[j]] = j
//			length := j - i + 1
//			if length > max {
//				max = length
//			}
//			j++
//		} else {
//			delete(m, s[i])
//			i++
//		}
//	}
//	return max
//}

func lengthOfLongestSubstring(s string) int {
	var m [128]int
	for i:=0; i < 128; i++ {
		m[i] = -1
	}
	//m := make(map[uint8]int)
	i := 0
	j := 0
	max := 0
	for j < len(s) {
		if m[s[j]] != -1 {
			if m[s[j]] + 1 > i {
				i = m[s[j]] + 1
			}
		}
		//if find,ok := m[s[j]]; ok {
		//	if find + 1 > i {
		//		i = find + 1
		//	}
		//}
		m[s[j]] = j
		length := j - i + 1
		if length > max {
			max = length
		}
		j++
	}
	return max
}