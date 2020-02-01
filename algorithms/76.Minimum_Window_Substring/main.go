package main

import (
	"fmt"
)

func main() {
	s := "a"
	fmt.Println(minWindow(s, "aa"))
}

func minWindow(s string, t string) string {
	mt := make(map[uint8]int)
	for i := 0; i < len(t); i++ {
		mt[t[i]]++
	}
	m := make([]int, 0)
	for i := 0; i < len(s); i++ {
		if _, ok := mt[s[i]]; ok {
			m = append(m, i)
		}
	}
	cnt := 0
	minLen := 0
	minLeft := -1
	i, j := 0, 0
	needCheckJ := true
	for i < len(m) && j < len(m) {
		if v, ok := mt[s[m[j]]]; ok && needCheckJ {
			if v >= 1 {
				cnt++
			}
			mt[s[m[j]]]--
		}
		needCheckJ = true
		if cnt == len(t) {
			tempLen := m[j] - m[i] + 1
			if minLeft == -1 || tempLen < minLen {
				minLen = tempLen
				minLeft = m[i]
			}
			mt[s[m[i]]]++
			if mt[s[m[i]]] > 0 {
				cnt--
			}
			needCheckJ = false
			i++
		} else {
			j++
		}
	}
	if minLeft == -1 {
		return ""
	}
	return s[minLeft : minLeft+minLen]
}
