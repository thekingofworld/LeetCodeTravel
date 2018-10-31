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
	haystack := "a"
	needle := "a"
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
	j := 0
	for i:=0; i<len(haystack) && j<len(needle); {
		if haystack[i] != needle[j] {
			if start != -1 {
				i = start + 1
			} else {
				i++
			}
			start = -1
			j = 0
		} else {
			if start == -1 {
				start = i
			}
			if j == len(needle)-1 {
				break
			}
			i++
			j++
		}
	}
	if start + len(needle) -1 <= len(haystack)-1 && j == len(needle)-1 {
		return start
	}
	return -1
}