package main

import "fmt"

/**
Write a function to find the longest common prefix string amongst an array of strings.

If there is no common prefix, return an empty string "".

Example 1:

Input: ["flower","flow","flight"]
Output: "fl"
Example 2:

Input: ["dog","racecar","car"]
Output: ""
Explanation: There is no common prefix among the input strings.
Note:

All given inputs are in lowercase letters a-z.
 */

func main() {
	strs := []string{"aa","a"}
	fmt.Println(longestCommonPrefix(strs))
}

func longestCommonPrefix(strs []string) string {
	length := len(strs)
	if length == 0 {
		return ""
	}
	if length == 1 {
		return strs[0]
	}
	prefix := make([]uint8, 0)
	first := strs[0]
	for i:=0; i<len(first); i++ {
		isSame := false
		for j:=1; j<length; j++ {
			if len(strs[j])-1 >= i && strs[j][i] == first[i] {
				isSame = true
			} else {
				isSame = false
				break
			}
		}
		if isSame {
			prefix = append(prefix, first[i])
		} else {
			break
		}
	}
	return string(prefix)
}