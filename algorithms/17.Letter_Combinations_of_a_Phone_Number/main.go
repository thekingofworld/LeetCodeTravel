package main

import "fmt"

/**
Given a string containing digits from 2-9 inclusive, return all possible letter combinations that the number could represent.

A mapping of digit to letters (just like on the telephone buttons) is given below. Note that 1 does not map to any letters.

Example:

Input: "23"
Output: ["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"].

Note:
Although the above answer is in lexicographical order, your answer could be in any order you want.
 */

func main() {
	input := "23"
	fmt.Println(letterCombinations(input))
}

var num2letters = map[string]string{"2":"abc", "3":"def", "4":"ghi", "5":"jkl", "6":"mno", "7":"pqrs", "8":"tuv", "9": "wxyz"}

func letterCombinations(digits string) []string {
	length := len(digits)
	iterates := make([]string, 0)
	for i:=0; i<length; i++ {
		val, ok := num2letters[string(digits[i])]
		if !ok {
			continue
		}
		if len(iterates) == 0 {
			for j:=0; j<len(val); j++ {
				iterates = append(iterates, string(val[j]))
			}
			continue
		}
		newIterates := make([]string, 0)
		for j:=0; j<len(iterates); j++ {
			for k:=0; k<len(val); k++ {
				newIterates = append(newIterates, string(iterates[j]) + string(val[k]))
			}
		}
		iterates = newIterates
	}
	return iterates
}