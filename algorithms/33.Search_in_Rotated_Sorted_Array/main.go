package main

import "fmt"

func main() {
	nums := []int{4,5,6,7,0,1,2}
	fmt.Println(search(nums, 1))
}

func search(nums []int, target int) int {
	left := 0
	right := len(nums)-1
	for left <= right {
		mid := (left + right)/2
		if target == nums[mid] {
			return mid
		} else if nums[mid] >= nums[left] {
			if target >= nums[left] && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			if target > nums[mid] && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return -1
}