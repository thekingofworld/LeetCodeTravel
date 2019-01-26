package main

import "fmt"

func main() {
	matrix := [][]byte{
		{'0','1','1','0','0','1','0','1','0','1'},
		{'0','0','1','0','1','0','1','0','1','0'},
		{'1','0','0','0','0','1','0','1','1','0'},
		{'0','1','1','1','1','1','1','0','1','0'},
		{'0','0','1','1','1','1','1','1','1','0'},
		{'1','1','0','1','0','1','1','1','1','0'},
		{'0','0','0','1','1','0','0','0','1','0'},
		{'1','1','0','1','1','0','0','1','1','1'},
		{'0','1','0','1','1','0','1','0','1','1'},
	}
	fmt.Println(maximalRectangle(matrix))
}

func maximalRectangle(matrix [][]byte) int {
	maxarea := 0
	m := len(matrix)
	if m == 0 {
		return 0
	}
	n := len(matrix[len(matrix)-1])
	if n == 0 {
		return 0
	}
	mem := make([][]int, m)
	for i:=0; i<m; i++ {
		mem[i] = make([]int, n)
	}
	for i:=0; i<m; i++ {
		for j:=0; j<n; j++ {
			if matrix[i][j] == '0' {
				mem[i][j] = 0
			} else {
				if i-1 < 0 {
					mem[i][j] = 1
				} else {
					mem[i][j] = mem[i-1][j] + 1
				}
			}
		}
	}
	for i:=0; i<m; i++ {
		left := make([]int, n)
		right := make([]int, n)
		for j:=0; j<n; j++ {
			if j-1<0 {
				left[j] = 0
				continue
			}
			if mem[i][j] == 0 {
				left[j] = 0
				continue
			}
			for k:=j-1; k >= 0 && mem[i][k] >= mem[i][j]; k-- {
				left[j] += 1
			}
		}
		for j:=n-1; j>=0; j-- {
			if j == (n-1) {
				right[j] = 0
				continue
			}
			if mem[i][j] == 0 {
				right[j] = 0
				continue
			}
			for k:=j+1; k<n && mem[i][k] >= mem[i][j]; k++ {
				right[j] += 1
			}
		}
		for j:=0; j<n; j++ {
			area := mem[i][j] * (left[j] + right[j] + 1)
			if area > maxarea {
				maxarea = area
			}
		}
	}
	return maxarea
}