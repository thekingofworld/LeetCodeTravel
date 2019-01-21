package main

import "fmt"

func main() {
	ob := [][]int{
		{0, 0, 0},
		{0, 1, 0},
		{0, 0, 0},
	}
	fmt.Println(uniquePathsWithObstacles(ob))
}

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	if len(obstacleGrid) > 0 && obstacleGrid[0][0] == 1 {
		return 0
	}
	n := len(obstacleGrid)
	m := len(obstacleGrid[n-1])
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
			if obstacleGrid[i][j] == 1 {
				mem[i][j] = 0
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