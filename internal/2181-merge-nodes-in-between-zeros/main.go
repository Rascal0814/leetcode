package _181_merge_nodes_in_between_zeros

const _ = `2181. 合并零之间的节点
中等
相关标签
相关企业
提示
给你一个链表的头节点 head ，该链表包含由 0 分隔开的一连串整数。链表的 开端 和 末尾 的节点都满足 Node.val == 0 。

对于每两个相邻的 0 ，请你将它们之间的所有节点合并成一个节点，其值是所有已合并节点的值之和。然后将所有 0 移除，修改后的链表不应该含有任何 0 。

 返回修改后链表的头节点 head 。

`

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeNodes(head *ListNode) *ListNode {
	var list = make([]int, 0)
	head = head.Next
	total := 0
	for head != nil {
		if head.Val == 0 && total != 0 {
			list = append(list, total)
			total = 0
		} else {
			total += head.Val
		}
		head = head.Next
	}
	var res = &ListNode{Val: list[len(list)-1]}
	for i := len(list) - 2; i >= 0; i-- {
		res = &ListNode{Val: list[i], Next: res}
	}
	return res
}

type ListNode struct {
	Val  int
	Next *ListNode
}
