package main

import "fmt"

/**
Implement strStr().

Return the index of the first occurrence of needle in haystack, or -1 if needle is not part of haystack.

Example 1:

Input: haystack = "hello", needle = "ll"
Output: 2
Example 2:

Input: haystack = "aaaaa", needle = "bba"
Output: -1
Clarification:

What should we return when needle is an empty string? This is a great question to ask during an interview.

For the purpose of this problem, we will return 0 when needle is an empty string. This is consistent to C's strstr() and Java's indexOf().
 */

func main() {
	haystack := "mississippi"
	needle := "issip"
	fmt.Println(strStr(haystack, needle))
}

func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	if len(haystack) < len(needle) {
		return -1
	}
	start := -1
	for i,j:=0,0; i<len(haystack) && j<len(needle); i++ {
		if haystack[i] != needle[j] {
			start = -1
			j = 0
		} else {
			fmt.Println(i)
			fmt.Println(j)
			if start == -1 {
				start = i
			}
			if j == len(needle)-1 {
				break
			}
			j++
		}
	}
	return start
}