package main

import "fmt"

func main() {
	nums := []int{3, 2, 4, 15}
	target := 6
	res := twoSum(nums, target)
	fmt.Println(res)
}

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, v := range nums {
		if j, ok := m[target - v]; ok {
			return []int{j, i}
		}
		m[v] = i
	}
	return nil
}