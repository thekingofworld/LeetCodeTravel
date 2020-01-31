package main

import "fmt"

func main() {
	in := []int{1, 2, 3, 4}
	fmt.Println(productExceptSelf(in))
}

func productExceptSelf(nums []int) []int {
	length := len(nums)
	if length == 0 {
		return nil
	}
	c := make([]int, length)
	for i := 0; i < length; i++ {
		if i == 0 {
			c[i] = 1
			continue
		}
		c[i] = c[i-1] * nums[i-1]
	}
	temp := 1
	for i := length - 2; i >= 0; i-- {
		temp = temp * nums[i+1]
		c[i] *= temp
	}
	return c
}
