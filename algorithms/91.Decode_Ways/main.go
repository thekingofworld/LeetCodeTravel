package main

import (
	"fmt"
)

func main() {
	fmt.Println(numDecodings("17"))
}

func numDecodings(s string) int {
	n := len(s)
	mem := make([]int, n+1)
	mem[0] = 1
	for i:=1; i<=n; i++ {
		oneLetter := 0
		doubleLetter := 0
		if s[i-1] != '0' {
			oneLetter = mem[i-1]
		}
		if i-2>=0 && ((s[i-2] == '1' ) || (s[i-2] == '2' && s[i-1] <= '6')) {
			doubleLetter = mem[i-2]
		}
		mem[i] = oneLetter + doubleLetter
	}
	return mem[n]
}