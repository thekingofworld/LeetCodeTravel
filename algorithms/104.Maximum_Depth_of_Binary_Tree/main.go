package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
		Right: &TreeNode{
			Val:   3,
			Left:  &TreeNode{Val: 4},
			Right: &TreeNode{Val: 5},
		},
	}
	fmt.Println(maxDepth(root))
}

func maxDepth(root *TreeNode) int {
	return maxDep(root, 0)
}

func maxDep(root *TreeNode, curDep int) int {
	if root == nil {
		return curDep
	}
	if root.Left != nil && root.Right != nil {
		return max(maxDep(root.Left, curDep+1), maxDep(root.Right, curDep+1))
	}
	if root.Left != nil {
		return maxDep(root.Left, curDep+1)
	}
	if root.Right != nil {
		return maxDep(root.Right, curDep+1)
	}
	return curDep + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
