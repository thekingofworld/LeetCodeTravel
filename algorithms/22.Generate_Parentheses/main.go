package main

import "fmt"

/**
Given n pairs of parentheses, write a function to generate all combinations of well-formed parentheses.

For example, given n = 3, a solution set is:

[
  "((()))",
  "(()())",
  "(())()",
  "()(())",
  "()()()"
]
 */

func main() {
	fmt.Println(generateParenthesis(4))
}

func generateParenthesis(n int) []string {
	parenthesis := make(map[string]bool, 0)
	res := make([]string, 0)
	for i:=0; i<n; i++ {
		if i == 0 {
			res = append(res, "()")
			parenthesis["()"] = true
			continue
		}
		newRes := make([]string, 0)
		for j:=0; j<len(res); j++ {
			length := len(res[j])
			for k:=0; k<length; k++ {
				item := res[j][:k] + "()" + res[j][k:]
				if _, ok := parenthesis[item]; !ok {
					newRes = append(newRes, item)
					parenthesis[item] = true
				}
			}
			last := res[j] + "()"
			if _, ok := parenthesis[last]; !ok {
				newRes = append(newRes, last)
				parenthesis[last] = true
			}
		}
		res = newRes
	}
	return res
}