package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(countAndSay(5))
}

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}
	say := "1"
	for i:=1; i<n; i++ {
		tmpCount := 1
		tmpSay := ""
		for j:=1; j<len(say); j++ {
			if say[j] != say[j-1] {
				tmpSay = tmpSay + strconv.Itoa(tmpCount) + string(say[j-1])
				tmpCount = 1
			} else {
				tmpCount++
			}
		}
		tmpSay = tmpSay + strconv.Itoa(tmpCount) + string(say[len(say)-1])
		say = tmpSay
	}
	return say
}