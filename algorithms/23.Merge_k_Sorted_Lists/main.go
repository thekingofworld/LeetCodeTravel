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
	node2.Val = 4
	node3 := new(ListNode)
	node3.Val = 5
	node1.Next = node2
	node2.Next = node3

	node4 := new(ListNode)
	node4.Val = 1
	node5 := new(ListNode)
	node5.Val = 3
	node6 := new(ListNode)
	node6.Val = 4
	node4.Next = node5
	node5.Next = node6

	node7 := new(ListNode)
	node7.Val = 2
	node8 := new(ListNode)
	node8.Val = 6
	node7.Next = node8

	lists := []*ListNode{node1, node4, node7}
	l := mergeKLists(lists)
	for l != nil {
		fmt.Println(l.Val)
		l = l.Next
	}
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
//func mergeKLists(lists []*ListNode) *ListNode {
//	var res *ListNode
//	for i:=0; i<len(lists); i++ {
//		if i == 0 {
//			res = lists[i]
//			continue
//		}
//		res = mergeTwoLists(res, lists[i])
//	}
//	return res
//}

func mergeKLists(lists []*ListNode) *ListNode {
	amount := len(lists)
	if amount == 0 {
		return nil
	}
	interval := 1
	for interval < amount {
		for i:=0; i<amount && i+interval<amount; {
			lists[i] = mergeTwoLists(lists[i], lists[i+interval])
			i = i + 2*interval
		}
		interval *= 2
	}
	return lists[0]
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	var newHead *ListNode
	var split *ListNode
	i,j := l1,l2
	for i != nil && j != nil {
		if i.Val <= j.Val {
			if newHead == nil {
				newHead = i
			}
			if split == nil {
				split = i
			} else {
				split.Next = i
				split = i
			}
			i = i.Next
		} else {
			if newHead == nil {
				newHead = j
			}
			if split == nil {
				split = j
			} else {
				split.Next = j
				split = j
			}
			j = j.Next
		}
	}
	if i != nil {
		split.Next = i
	}
	if j != nil {
		split.Next = j
	}
	return newHead
}