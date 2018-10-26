package main

import (
	"sort"
	"fmt"
)

/**
Given an array nums of n integers and an integer target, find three integers in nums such that the sum is closest to target.
Return the sum of the three integers. You may assume that each input would have exactly one solution.

Example:

Given array nums = [-1, 2, 1, -4], and target = 1.

The sum that is closest to the target is 2. (-1 + 2 + 1 = 2).
 */

func main() {
	nums := []int{-1, 2, 1, -4}
	fmt.Println(threeSumClosest(nums, 1))
}

func threeSumClosest(nums []int, target int) int {
	closest := 0
	var distance int
	first := true
	sort.Ints(nums)
	length := len(nums)
	for i:=0; i<length-2; i++ {
		for j,k := i+1,length-1; j<k; {
			var curDistance int
			s := nums[i] + nums[j] + nums[k]
			if s == target {
				return s
			} else if s < target {
				curDistance = target - s
				j++
			} else {
				curDistance = s - target
				k--
			}
			if first {
				first = false
				distance = curDistance
				closest = s
			}
			if curDistance < distance {
				distance = curDistance
				closest = s
			}
		}
	}
	return closest
}