package _290_convert_binary_number_in_a_linked_list_to_integer

const _ = `
1290. 二进制链表转整数
提示
给你一个单链表的引用结点 head。链表中每个结点的值不是 0 就是 1。已知此链表是一个整数数字的二进制表示形式。

请你返回该链表所表示数字的 十进制值 。`

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getDecimalValue(head *ListNode) int {
	var next = head
	var vals = make([]int, 0)
	for next != nil {
		vals = append(vals, next.Val)
		next = next.Next
	}

	var total = 0
	for i := 0; i < len(vals); i++ {
		if vals[i] == 0 {
			continue
		}
		total += 1 << (len(vals) - i - 1)
	}
	return total
}

type ListNode struct {
	Val  int
	Next *ListNode
}
