package main

import "fmt"

/**
 * examples:
 * "*":匹配0个或多个字符
 * ".":匹配1个字符
 * s = "aa"  p = "a"
 * s = "aa"  p = "a*"
 * s = "ab"  p = ".*"
 * s = "aab" p = "c*a*b"
 * s = "mississippi" p = "mis*is*p*."
 */

func main() {
	s := "aaa"
	p := "aaaa"
	fmt.Println(s[3:])
	if isMatch(s, p) {
		fmt.Println("match")
	} else {
		fmt.Println("not match")
	}
}

//####solution 1####
//func isMatch(s string, p string) bool {
//	return match(s, len(s)-1, p, len(p)-1)
//}

//func match(s string, i int, p string, j int) bool {
//	if j == -1 {
//		if i == -1 {
//			return true
//		} else {
//			return false
//		}
//	}
//	if p[j] == '*' {
//		if i > -1 && (p[j-1] == s[i] || p[j-1] == '.') {
//			if match(s, i-1, p, j) {
//				return true
//			}
//		}
//		return match(s, i, p, j-2)
//	}
//	if i > -1 && j > -1 {
//		if s[i] == p[j] || p[j] == '.' {
//			return match(s, i-1, p, j-1)
//		}
//	}
//	return false
//}

//####solution 2####
func isMatch(s string, p string) bool {
	lengthS := len(s)
	lengthP := len(p)
	if lengthP == 0 {
		return lengthS == 0
	}

	var firstMatch = lengthS != 0 && (s[0] == p[0] || p[0] == '.')
	if lengthP >= 2 && p[1] == '*' {
		return (firstMatch && isMatch(s[1:], p)) || isMatch(s, p[2:])
	} else {
		return firstMatch && isMatch(s[1:], p[1:])
	}
}