package main

import "fmt"

type ListNode struct{
	Val int
	Next *ListNode
}

func main() {
	node2 := new(ListNode)
	node2.Val = 3
	node2.Next = nil
	node1 := new(ListNode)
	node1.Val = 4
	node1.Next = node2
	head := new(ListNode)
	head.Val = 2
	head.Next = node1

	node3 := new(ListNode)
	node3.Val = 4
	node3.Next = nil
	node4 := new(ListNode)
	node4.Val = 6
	node4.Next = node3
	head2 := new(ListNode)
	head2.Val = 5
	head2.Next = node4

	sum := addTwoNumbers(head, head2)
	for sum != nil {
		fmt.Println(sum.Val)
		sum = sum.Next
	}
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	carry := 0
	var head *ListNode
	var tmp *ListNode
	for l1 != nil || l2 != nil || carry != 0 {
		sum := 0
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		if carry != 0 {
			sum += carry
		}
		carry = sum / 10
		val := sum % 10
		node := new(ListNode)
		node.Val = val
		node.Next = nil

		if head == nil {
			head = node
		}
		if tmp == nil {
			tmp = node
		} else {
			tmp.Next = node
			tmp = node
		}
	}
	return head
}