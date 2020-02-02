package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 9,
		},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val: 15,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
	}
	fmt.Println(levelOrder(root))
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var res [][]int
	tmpLevel := make([]int, 0)
	curCnt := 0
	curLevelCnt := 1
	nextLevelCnt := 0
	unHandleQueue := []*TreeNode{root}
	for len(unHandleQueue) > 0 {
		tmpLevel = append(tmpLevel, unHandleQueue[0].Val)
		curCnt++
		if unHandleQueue[0].Left != nil {
			unHandleQueue = append(unHandleQueue, unHandleQueue[0].Left)
			nextLevelCnt++
		}
		if unHandleQueue[0].Right != nil {
			unHandleQueue = append(unHandleQueue, unHandleQueue[0].Right)
			nextLevelCnt++
		}
		if curCnt == curLevelCnt {
			res = append(res, tmpLevel)
			tmpLevel = make([]int, 0)
			curCnt = 0
			curLevelCnt = nextLevelCnt
			nextLevelCnt = 0
		}
		unHandleQueue = unHandleQueue[1:]
	}
	return res
}
