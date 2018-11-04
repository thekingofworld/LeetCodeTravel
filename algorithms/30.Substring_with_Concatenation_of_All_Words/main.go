package main

import (
	"fmt"
)

func main() {
	fmt.Println(findSubstring("wordgoodgoodgoodbestword", []string{"word","good","best","word"}))
}

func findSubstring(s string, words []string) []int {
	if len(s) == 0 {
		return []int{}
	}
	wordCount := len(words)
	if wordCount == 0 {
		return []int{}
	}
	wordLength := len(words[0])
	for _, val := range words {
		if len(val) != wordLength {
			return []int{}
		}
	}
	res := make([]int, 0)
	exist := make(map[string]int)
	for i:=0; i<len(words); i++ {
		if _, ok := exist[words[i]]; ok {
			exist[words[i]]++
		} else {
			exist[words[i]] = 1
		}
	}
	for i:=0; i<=len(s)-wordLength*wordCount; i++ {
		word := s[i:i+wordLength]
		if _, ok := exist[word]; ok {
			tmpMap := make(map[string]int)
			for k,v := range exist {
				tmpMap[k] = v
			}
			find := true
			for j:=0; j<wordCount; j++ {
				index := i+j*wordLength
				newWord := s[index:index+wordLength]
				if val, ok := tmpMap[newWord]; ok && val > 0 {
					tmpMap[newWord]--
				} else {
					find = false
					break
				}
			}
			if find {
				res = append(res, i)
			}
		}
	}
	return res
}