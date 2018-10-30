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
	node5 := new(ListNode)
	node5.Val = 5
	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5
	head := reverseKGroup(node1, 2)
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
func reverseKGroup(head *ListNode, k int) *ListNode {
	var newHead *ListNode
	first := head
	for first != nil {
		cur, next := reverse(first, k)
		if newHead == nil {
			newHead = cur
		}
		if cur == nil || next == nil {
			break
		}
		first.Next = next
		first = next
	}
	return newHead
}

func reverse(node *ListNode, k int) (*ListNode, *ListNode) {
	if k == 1 {
		return node, node.Next
	}
	if node.Next == nil {
		return nil, nil
	}
	k--
	cur, next := reverse(node.Next, k)
	if cur == nil || next == nil {
		return nil, nil
	}
	node.Next.Next = node
	return cur,next
}