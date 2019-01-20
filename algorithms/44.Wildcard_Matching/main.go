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
	mem := make([][]bool, len(s)+1)
	for i:=0; i<=len(s); i++ {
		mem[i] = make([]bool, len(p)+1)
	}
	mem[0][0] = true
	for i:=0; i<=len(s); i++ {
		for j:=1; j<=len(p); j++ {
			if p[j-1] == '*' {
				mem[i][j] = mem[i][j-1] || (i > 0 && mem[i-1][j])
			} else {
				mem[i][j] = i > 0 && mem[i-1][j-1] && (s[i-1] == p[j-1] || p[j-1] == '?')
			}
		}
	}
	return mem[len(s)][len(p)]
}