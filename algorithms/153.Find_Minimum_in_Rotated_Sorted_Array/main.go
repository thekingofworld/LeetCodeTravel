package main

import "fmt"

func main() {
	in := []int{3, 4, 5, 6, 1, 2}
	fmt.Println(findMin(in))
}

func findMin(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	low := 0
	high := len(nums) - 1
	mid := 0
	for low < high {
		mid = (low + high) / 2
		if nums[mid] > nums[high] {
			low = mid + 1
		} else {
			high = mid
		}
	}
	return nums[low]
}
