package main

import "fmt"

func main() {
	nums := []int{5,7,7,8,8,10}
	fmt.Println(searchRange(nums, 6))
}

func searchRange(nums []int, target int) []int {
	res := make([]int, 0)
	left := 0
	right := len(nums)-1
	for left <= right {
		mid := (left + right)/2
		if nums[mid] == target {
			res = append(res, []int{mid, mid}...)
			for i:=mid-1; i>=0; i-- {
				if nums[i] == target {
					res[0] = i
				} else {
					break
				}
			}
			for i:=mid+1; i<len(nums); i++ {
				if nums[i] == target {
					res[1] = i
				} else {
					break
				}
			}
			break
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	if len(res) == 0 {
		res = append(res, []int{-1, -1}...)
	}
	return res
}