package main

import "fmt"

/**
Given an integer array nums, find the contiguous subarray (containing at least one number) which has the largest sum and return its sum.

Example:

Input: [-2,1,-3,4,-1,2,1,-5,4],
Output: 6
Explanation: [4,-1,2,1] has the largest sum = 6.
Follow up:

If you have figured out the O(n) solution, try coding another solution using the divide and conquer approach, which is more subtle.
 */

func main() {
	nums := []int{-2,1,-3,4,-1,2,1,-5,4}
	fmt.Println(maxSubArray(nums))
}

func maxSubArray(nums []int) int {
	var max int
	mem := make([]int, len(nums))
	mem[0] = nums[0]
	max = nums[0]
	for i:=1; i<len(nums); i++ {
		if mem[i-1] > 0 {
			mem[i] = nums[i] + mem[i-1]
		} else {
			mem[i] = nums[i]
		}
		if mem[i] > max {
			max = mem[i]
		}
	}
	return max
}