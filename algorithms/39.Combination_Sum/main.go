package main

import (
	"sort"
	"fmt"
)

func main() {
	candidates := []int{5, 10, 8, 4, 3, 12, 9}
	target := 27
	fmt.Println(combinationSum(candidates, target))
}

var allCombinations [][]int

func combinationSum(candidates []int, target int) [][]int {
	var combinations [][]int
	for _, candidate := range candidates {
		var sum []int
		sum = append(sum, candidate)
		combination(candidates, sum, target)
	}
	combinations = allCombinations
	allCombinations = [][]int{}
	return combinations
}

func combination(candidates []int, sum []int, target int) {
	curSum := getSum(sum)
	if curSum == target {
		addCombination(sum)
		return
	} else if curSum > target {
		return
	}
	for _, candidate := range candidates {
		sum = append(sum, candidate)
		combination(candidates, sum, target)
		sum = sum[:len(sum)-1]
	}
}

func getSum(items []int) int {
	sum := 0
	for _, item := range items {
		sum += item
	}
	return sum
}

func addCombination(combination []int) {
	copyCom := make([]int, len(combination))
	copy(copyCom, combination)
	sort.Ints(copyCom)
	isHaveSame := false
	for _, items := range allCombinations {
		if len(items) == len(copyCom) {
			check := true
			for i := range items {
				if copyCom[i] != items[i] {
					check = false
				}
			}
			if check {
				isHaveSame = true
			}
		}
	}
	if !isHaveSame {
		allCombinations = append(allCombinations, copyCom)
	}
}