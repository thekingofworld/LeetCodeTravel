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
			Val: 4,
			Left: &TreeNode{
				Val: 1,
			},
			Right: &TreeNode{
				Val: 2,
			},
		},
		Right: &TreeNode{
			Val: 5,
		},
	}
	root2 := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val: 1,
		},
		Right: &TreeNode{
			Val: 2,
		},
	}
	fmt.Println(isSubtree(root, root2))
}

func isSubtree(s *TreeNode, t *TreeNode) bool {
	return checkIsSubTree(s, t, false)
}

func checkIsSubTree(s *TreeNode, t *TreeNode, checkEqual bool) bool {
	if s == nil && t == nil {
		return true
	}
	if s == nil && t != nil {
		return false
	}
	if s != nil && t == nil {
		return false
	}
	if checkEqual && s.Val != t.Val {
		return false
	}
	if s.Val == t.Val {
		res := checkIsSubTree(s.Left, t.Left, true) &&
			checkIsSubTree(s.Right, t.Right, true)
		if res {
			return true
		}
	}
	return checkIsSubTree(s.Left, t, false) ||
		checkIsSubTree(s.Right, t, false)
}
