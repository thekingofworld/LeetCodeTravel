package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{1, 2, 0}
	fmt.Println(firstMissingPositive(nums))
}

func firstMissingPositive(nums []int) int {
	sort.Ints(nums)
	curPositive := 1
	for _, num := range nums {
		if num == curPositive {
			curPositive++
		}
	}
	return curPositive
}