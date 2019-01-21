package main

import "fmt"

func main() {
	fmt.Println(uniquePaths(3, 2))
}

func uniquePaths(m int, n int) int {
	if m < 1 || n < 1 {
		return 0
	}
	mem := make([][]int, n)
	for i:=0; i<n; i++ {
		mem[i] = make([]int, m)
	}
	mem[0][0] = 1
	for i:=0; i<n; i++ {
		for j:=0; j<m; j++ {
			if i == 0 && j == 0 {
				continue
			}
			down := 0
			right := 0
			if i-1 >= 0 {
				down = mem[i-1][j]
			}
			if j-1 >= 0 {
				right = mem[i][j-1]
			}
			mem[i][j] = down + right
		}
	}
	return mem[n-1][m-1]
}