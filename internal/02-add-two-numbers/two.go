package main

func atn(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	n1, n2, current, next := 0, 0, 0, head

	for l1 != nil || l2 != nil || current != 0 {
		n1, n2 = 0, 0
		if l1 != nil {
			n1, l1 = l1.Val, l1.Next
		}

		if l2 != nil {
			n2, l2 = l2.Val, l2.Next
		}

		node := &ListNode{Val: (n1 + n2 + current) % 10}
		current = (n1 + n2 + current) / 10
		next.Next = node
		next = next.Next
	}

	return head.Next
}
