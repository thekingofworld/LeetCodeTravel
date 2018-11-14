package main

import "fmt"

/**
Given a sorted array and a target value, return the index if the target is found. If not, return the index where it would be if it were inserted in order.

You may assume no duplicates in the array.

Example 1:

Input: [1,3,5,6], 5
Output: 2
Example 2:

Input: [1,3,5,6], 2
Output: 1
Example 3:

Input: [1,3,5,6], 7
Output: 4
Example 4:

Input: [1,3,5,6], 0
Output: 0
 */

func main() {
	nums := []int{1,3,5,6}
	fmt.Println(searchInsert(nums, 0))
}

func searchInsert(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}
	index := binarySearch(nums, target)
	if index == -1 {
		if target < nums[0] {
			return 0
		}
		for i:=len(nums)-1; i>=0; i-- {
			if target > nums[i] {
				index = i + 1
				break
			}
		}
	}
	return index
}

func binarySearch(nums []int, target int) int {
	left := 0
	right := len(nums)-1
	for left <= right {
		mid := (left+right)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}