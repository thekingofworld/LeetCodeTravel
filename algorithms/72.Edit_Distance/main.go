package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(minDistance("intention", "execution"))
}

func minDistance(word1 string, word2 string) int {
	m := len(word1)
	n := len(word2)
	mem := make([][]int, m+1)
	for i:=0; i<=m; i++ {
		mem[i] = make([]int, n+1)
	}
	mem[0][0] = 0
	for i:=0; i<=m; i++ {
		for j:=0; j<=n; j++ {
			if i == 0 && j == 0 {
				continue
			}
			var dist,D,I,S int
			if i-1 >= 0 {
				D = mem[i-1][j] + 1
			} else {
				D = math.MaxInt64
			}
			if j-1 >= 0 {
				I = mem[i][j-1] + 1
			} else {
				I = math.MaxInt64
			}
			if i-1 >= 0 && j-1 >= 0 {
				if word1[i-1] == word2[j-1] {
					dist = 0
				} else {
					dist = 1
				}
				S = mem[i-1][j-1] + dist
			} else {
				S = math.MaxInt64
			}
			mem[i][j] = minTree(D, I, S)
		}
	}
	return mem[m][n]
}

func minTree(a, b, c int) int {
	if a > b {
		if b > c {
			return c
		} else {
			return b
		}
	} else {
		if a > c {
			return c
		} else {
			return a
		}
	}
}