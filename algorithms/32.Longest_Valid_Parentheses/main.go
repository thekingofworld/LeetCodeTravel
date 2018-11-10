package main

import "fmt"

func main() {
	fmt.Println(longestValidParentheses("())"))
}

func longestValidParentheses(s string) int {
	longest := 0
	dp := make([]int, len(s))
	for i:=0; i<len(s); i++ {
		if s[i] == '(' {
			dp[i] = 0
		}
		if s[i] == ')' && i > 0 {
			if s[i-1] == '(' {
				if i >= 2 {
					dp[i] = dp[i-2] + 2
				} else {
					dp[i] = 2
				}
			} else if s[i-1] == ')' && i-dp[i-1]-1 >= 0 && s[i-dp[i-1]-1] == '(' {
				if i-dp[i-1]-2 >= 0 {
					dp[i] = dp[i-1] + dp[i-dp[i-1]-2] + 2
				} else {
					dp[i] = dp[i-1] + 2
				}
			}
		}
		if dp[i] > longest {
			longest = dp[i]
		}
	}
	return longest
}