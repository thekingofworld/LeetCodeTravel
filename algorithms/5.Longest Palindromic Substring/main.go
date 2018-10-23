package main

import (
	"fmt"
)

func main() {
	s := longestPalindrome("babad")
	fmt.Println(s)
}

func longestPalindrome(s string) string {
	var maxRight,pos,maxLen,maxPos int
	size, data := preProcess(s)
	var R = make([]int, size)
	var length = len(data)
	for i := 0; i < length; i++ {
		if i < maxRight {
			R[i] = min(R[2*pos-i], maxRight-i)
		} else {
			R[i] = 1
		}
		for (i - R[i] >= 0) && (i + R[i] < length) && (data[i-R[i]] == data[i+R[i]]) {
			R[i] = R[i] + 1
		}
		if i + R[i] - 1 > maxRight {
			maxRight = i + R[i] - 1
			pos = i
		}
		if R[i] > maxLen {
			maxLen = R[i]
			maxPos = i
		}
	}
	var res []rune
	for _,c := range data[(maxPos-(maxLen-1)):(maxPos-(maxLen-1))+(2*maxLen-1)] {
		if string(c) != "#" {
			res = append(res, c)
		}
	}
	return string(res)
}

func preProcess(s string) (int, []rune) {
	var runes = []rune(s)
	var size = len(runes)*2 + 1
	var data = make([]rune, 0, size)
	data = append(data, '#')
	for i := range runes {
		data = append(data, runes[i])
		data = append(data, '#')
	}
	return size, data
}

func min(a,b int) int {
	if a <= b {
		return a
	} else {
		return b
	}
}

//func oldlongestPalindrome(s string) string {
//	length := len(s)
//	if length == 0 {
//		return ""
//	}
//	if length == 1 || checkPalindrome(s) {
//		return s
//	}
//	for i := length; i > 0; i-- {
//		for j := 0; j <= length - i; j++ {
//			if checkPalindrome(s[j:j+i]) {
//				return s[j:j+i]
//			}
//		}
//	}
//	return ""
//}
//
//func checkPalindrome(s string) bool {
//	length := len(s)
//	find := 1
//	i := 0
//	for i < length/2 {
//		if s[i] != s[length-i-1] {
//			find = 0
//			break
//		}
//		i++
//	}
//	if find == 1 {
//		return true
//	}
//	return false
//}