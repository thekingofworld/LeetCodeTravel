package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

func main() {
	node1 := &ListNode{Val:1}
	node2 := &ListNode{Val:2}
	node3 := &ListNode{Val:3}
	node4 := &ListNode{Val:4}
	node5 := &ListNode{Val:5}
	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5
	head := removeNthFromEnd(node1, 2)
	var n *ListNode
	n = head
	for n != nil {
		fmt.Println(n.Val)
		n = n.Next
	}
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	node := new(ListNode)
	node.Val = 0
	node.Next = head
	remove(node, n)
	newHead := node.Next
	node = nil
	return newHead
}

func remove(node *ListNode, n int) int {
	if node.Next == nil {
		return 1
	}
	x := remove(node.Next, n)
	if x == n {
		node.Next = node.Next.Next
	}
	x++
	return x
}