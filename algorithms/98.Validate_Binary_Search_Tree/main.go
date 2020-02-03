package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 1,
		},
		Right: &TreeNode{
			Val: 8,
			Left: &TreeNode{
				Val: 7,
			},
			Right: &TreeNode{
				Val: 9,
			},
		},
	}
	fmt.Println(isValidBST(root))
}

func isValidBST(root *TreeNode) bool {
	inOrder := inOrder(root)
	for i := 0; i < len(inOrder)-1; i++ {
		if inOrder[i] >= inOrder[i+1] {
			return false
		}
	}
	return true
}

func inOrder(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var l []int
	l = append(l, inOrder(root.Left)...)
	l = append(l, root.Val)
	l = append(l, inOrder(root.Right)...)
	return l
}
