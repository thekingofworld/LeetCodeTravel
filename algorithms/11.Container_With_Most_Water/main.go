package main

import "fmt"

func main() {
	x := []int{1,8,6,2,5,4,8,3,7}
	fmt.Println(maxArea(x))
}

func maxArea(height []int) int {
	l := 0
	r := len(height) - 1
	maxArea := 0
	for l < r {
		area := (r - l) * min(height[l], height[r])
		if area > maxArea {
			maxArea = area
		}
		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}
	return maxArea
}

//func maxArea(height []int) int {
//	maxArea := 0
//	for i := range height {
//		for j := i+1; j < len(height); j++ {
//			area := (j - i) * min(height[i], height[j])
//			if area > maxArea {
//				maxArea = area
//			}
//		}
//	}
//	return maxArea
//}
//
func min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}