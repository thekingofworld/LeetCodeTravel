package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

func main() {
	node1 := new(ListNode)
	node1.Val = 1
	node2 := new(ListNode)
	node2.Val = 2
	node3 := new(ListNode)
	node3.Val = 3
	node4 := new(ListNode)
	node4.Val = 4
	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	head := swapPairs(node1)
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}
	var newHead *ListNode
	n := 0
	interval := head
	last := head
	for interval != nil && interval.Next != nil {
		tmp := interval.Next
		interval.Next = interval.Next.Next
		tmp.Next = interval
		if n != 0 {
			last.Next = tmp
			last = interval
		}
		interval = interval.Next
		if newHead == nil {
			newHead = tmp
		}
		n++
	}
	return newHead
}