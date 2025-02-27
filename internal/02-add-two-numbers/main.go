package main

import (
	"fmt"
)

/***

you are given two non-empty linked lists representing two non-negative integers. the digits are stored in reverse order and each of their nodes contain a single digit. add the two numbers and return it as a linked list.

you may assume the two numbers do not contain any leading zero, except the number 0 itself.

example:

	input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
	output: 7 -> 0 -> 8
	explanation: 342 + 465 = 807.

*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	a := &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9}}}}
	c := &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9}}}
	//[9,9,9,9,9,9,9]
	//	[9,9,9,9]
	//b := &ListNode{Val: 0}
	node := atn(a, c)

	for node != nil {
		fmt.Println(node.Val)
		node = node.Next
	}

}

func addtownumber(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}

	n1, n2, curry, next := 0, 0, 0, head

	for l1 != nil || l2 != nil || curry != 0 {
		n1, n2 = 0, 0
		if l1 != nil {
			n1, l1 = l1.Val, l1.Next
		}
		if l2 != nil {
			n2, l2 = l2.Val, l2.Next
		}

		node := &ListNode{Val: (n1 + n2 + curry) % 10}

		curry = (n1 + n2 + curry) / 10
		next.Next = node
		next = next.Next
	}

	return head.Next
}
