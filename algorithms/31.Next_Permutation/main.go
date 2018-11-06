package main

import (
	"fmt"
	"sort"
)

/**
Implement next permutation, which rearranges numbers into the lexicographically next greater permutation of numbers.

If such arrangement is not possible, it must rearrange it as the lowest possible order (ie, sorted in ascending order).

The replacement must be in-place and use only constant extra memory.

Here are some examples. Inputs are in the left-hand column and its corresponding outputs are in the right-hand column.

1,2,3 → 1,3,2
3,2,1 → 1,2,3
1,1,5 → 1,5,1
 */

func main() {
	nums := []int{4,2,0,2,3,2,0}
	nextPermutation(nums)
	fmt.Println(nums)
}

func nextPermutation(nums []int)  {
	start := -1
	end := len(nums)
	for i:=len(nums)-1; i>0; i-- {
		for j:=i-1; j>=0; j-- {
			if nums[i] > nums[j] && j > start {
				start = j
				end = i
				break
			}
		}
	}
	if start != -1 {
		nums[start],nums[end] = nums[end],nums[start]
		sort.Ints(nums[start+1:])
	} else {
		for i:=0; i<len(nums)/2; i++ {
			nums[i],nums[len(nums)-1-i] = nums[len(nums)-1-i], nums[i]
		}
	}
}