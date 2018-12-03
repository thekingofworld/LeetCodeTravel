package main

import (
	"fmt"
)

func main() {
	nums := []int{-1,4,2,1,9,10}
	fmt.Println(firstMissingPositive(nums))
}

func firstMissingPositive(nums []int) int {
	n := len(nums)
	for i:=0; i<n; i++ {
		if nums[i]-1 >= 0 && nums[i]-1 < n && nums[i] != nums[nums[i]-1] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
			i--
		}
	}

	nextPositive := 1
	for i:=0; i<n; i++ {
		if nums[i] == nextPositive {
			nextPositive++
		}
	}

	return nextPositive
}