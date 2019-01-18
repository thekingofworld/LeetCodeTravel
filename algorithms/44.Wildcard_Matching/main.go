package main

import "fmt"

/**
Given an input string (s) and a pattern (p), implement wildcard pattern matching with support for '?' and '*'.

'?' Matches any single character.
'*' Matches any sequence of characters (including the empty sequence).
The matching should cover the entire input string (not partial).

Note:

s could be empty and contains only lowercase letters a-z.
p could be empty and contains only lowercase letters a-z, and characters like ? or *.
Example 1:

Input:
s = "aa"
p = "a"
Output: false
Explanation: "a" does not match the entire string "aa".
Example 2:

Input:
s = "aa"
p = "*"
Output: true
Explanation: '*' matches any sequence.
Example 3:

Input:
s = "cb"
p = "?a"
Output: false
Explanation: '?' matches 'c', but the second letter is 'a', which does not match 'b'.
Example 4:

Input:
s = "adceb"
p = "*a*b"
Output: true
Explanation: The first '*' matches the empty sequence, while the second '*' matches the substring "dce".
Example 5:

Input:
s = "acdcb"
p = "a*c?b"
Output: false
 */

func main() {
	if isMatch("aaaabaaaabbbbaabbbaabbaababbabbaaaababaaabbbbbbaabbbabababbaaabaabaaaaaabbaabbbbaababbababaabbbaababbbba",
		"*****b*aba***babaa*bbaba***a*aaba*b*aa**a*b**ba***a*a*") {
		fmt.Println("ok, is match")
	} else {
		fmt.Println("oh no, is not match")
	}
}

func isMatch(s string, p string) bool {
	return SubMatch(s, p, 0, 0)
}

func SubMatch(s string, p string, si int, pj int) bool {
	if si == len(s) && pj == len(p) {
		return true
	}
	if pj < len(p) && p[pj] == '*' {
		if si == len(s) {
			return SubMatch(s, p, si, pj+1)
		}
		return SubMatch(s, p, si+1, pj) || SubMatch(s, p, si+1, pj+1) || SubMatch(s, p, si, pj+1)
	}
	if si < len(s) && pj < len(p) && (p[pj] == '?' || s[si] == p[pj]) {
		return SubMatch(s, p, si+1, pj+1)
	}
	return false
}