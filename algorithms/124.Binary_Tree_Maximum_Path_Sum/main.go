package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val: 11,
				Left: &TreeNode{
					Val: 7,
				},
				Right: &TreeNode{
					Val: 2,
				},
			},
		},
		Right: &TreeNode{
			Val: 8,
			Left: &TreeNode{
				Val: 13,
			},
			Right: &TreeNode{
				Val: 4,
				Right: &TreeNode{
					Val: 1,
				},
			},
		},
	}
	fmt.Println(maxPathSum(root))
}

func maxPathSum(root *TreeNode) int {
	maxSum := math.MinInt64
	_ = throughPathSum(root, &maxSum)
	return maxSum
}

func throughPathSum(root *TreeNode, maxSum *int) int {
	if root == nil {
		return 0
	}
	var leftSum, rightSum, curMax, throughSum int
	if root.Left != nil && root.Right != nil {
		leftSum = throughPathSum(root.Left, maxSum)
		rightSum = throughPathSum(root.Right, maxSum)
		curMax = max(leftSum, rightSum, leftSum+rightSum+root.Val, root.Val,
			leftSum+root.Val, rightSum+root.Val)
		throughSum = max(leftSum+root.Val, root.Val, rightSum+root.Val)
	} else if root.Left != nil {
		leftSum = throughPathSum(root.Left, maxSum)
		curMax = max(leftSum, leftSum+root.Val, root.Val)
		throughSum = max(leftSum+root.Val, root.Val)
	} else if root.Right != nil {
		rightSum = throughPathSum(root.Right, maxSum)
		curMax = max(rightSum, rightSum+root.Val, root.Val)
		throughSum = max(root.Val, rightSum+root.Val)
	} else {
		curMax = root.Val
		throughSum = root.Val
	}
	if curMax > *maxSum {
		*maxSum = curMax
	}
	return throughSum
}

func max(args ...int) int {
	curMax := math.MinInt64
	for _, v := range args {
		if v > curMax {
			curMax = v
		}
	}
	return curMax
}
