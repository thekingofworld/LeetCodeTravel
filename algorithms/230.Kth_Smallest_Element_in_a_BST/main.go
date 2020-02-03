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
			Val: 3,
			Left: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 1,
				},
			},
			Right: &TreeNode{
				Val: 4,
			},
		},
		Right: &TreeNode{
			Val: 6,
		},
	}
	fmt.Println(kthSmallest(root, 3))
}

func kthSmallest(root *TreeNode, k int) int {
	inOrder := inOrder(root)
	return inOrder[k-1]
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
