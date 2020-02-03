package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := buildTree([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7})
	fmt.Println(root)
	fmt.Println(root.Left)
	fmt.Println(root.Left.Left)
	fmt.Println(root.Left.Right)
	fmt.Println(root.Right)
	fmt.Println(root.Right.Left)
	fmt.Println(root.Right.Right)
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}
	root := &TreeNode{
		Val: preorder[0],
	}
	index := findIndex(inorder, preorder[0])
	if index == -1 {
		return root
	}
	if len(preorder) > 0 && len(inorder) > 0 {
		root.Left = buildTree(preorder[1:1+index], inorder[0:index])
		root.Right = buildTree(preorder[1+index:], inorder[index+1:])
	}
	return root
}

func findIndex(data []int, val int) int {
	for i, d := range data {
		if d == val {
			return i
		}
	}
	return -1
}
