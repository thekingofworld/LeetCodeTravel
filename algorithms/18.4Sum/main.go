package main

import (
	"sort"
	"fmt"
)

/**
Given an array nums of n integers and an integer target, are there elements a, b, c, and d in nums such that a + b + c + d = target?
Find all unique quadruplets in the array which gives the sum of target.

Note:

The solution set must not contain duplicate quadruplets.

Example:

Given array nums = [1, 0, -1, 0, -2, 2], and target = 0.

A solution set is:
[
  [-1,  0, 0, 1],
  [-2, -1, 1, 2],
  [-2,  0, 0, 2]
]
 */

func main() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	fmt.Println(fourSum(nums, -1))
}

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	length := len(nums)
	res := make([][]int, 0)
	for i:=0; i<length-3; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j:=i+1; j<length-2; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			for k,l:=j+1,length-1; k < l; {
				s := nums[i] + nums[j] + nums[k] + nums[l]
				if s == target {
					res = append(res, []int{nums[i], nums[j], nums[k], nums[l]})
					for k < l && nums[k] == nums[k+1] {
						k++
					}
					for k < l && nums[l] == nums[l-1] {
						l--
					}
					k++
					l--
				} else if s < target {
					k++
				} else {
					l--
				}
			}
		}
	}
	return res
}