package main

import "fmt"

func main() {
	grid := [][]int{
		{1,3,1},
		{1,5,1},
		{4,2,1},
	}
	fmt.Println(minPathSum(grid))
}

func minPathSum(grid [][]int) int {
	n := len(grid)
	m := len(grid[n-1])
	mem := make([][]int, n)
	for i:=0; i<n; i++ {
		mem[i] = make([]int, m)
	}
	for i:=0; i<n; i++ {
		for j:=0; j<m; j++ {
			var preMin int
			if i == 0 && j == 0 {
				preMin = 0
			} else if i-1 >= 0 && j-1 >= 0 {
				preMin = min(mem[i-1][j], mem[i][j-1])
			} else if i-1 >= 0 {
				preMin = mem[i-1][j]
			} else {
				preMin = mem[i][j-1]
			}
			mem[i][j] = grid[i][j] + preMin
		}
	}
	return mem[n-1][m-1]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}