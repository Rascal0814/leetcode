package _058_find_the_minimum_and_maximum_number_of_nodes_between_critical_points

import "math"

const _ = `
2058. 找出临界点之间的最小和最大距离
中等
相关标签
相关企业
提示
链表中的 临界点 定义为一个 局部极大值点 或 局部极小值点 。

如果当前节点的值 严格大于 前一个节点和后一个节点，那么这个节点就是一个  局部极大值点 。

如果当前节点的值 严格小于 前一个节点和后一个节点，那么这个节点就是一个  局部极小值点 。

注意：节点只有在同时存在前一个节点和后一个节点的情况下，才能成为一个 局部极大值点 / 极小值点 。

给你一个链表 head ，返回一个长度为 2 的数组 [minDistance, maxDistance] ，其中 minDistance 是任意两个不同临界点之间的最小距离，maxDistance 是任意两个不同临界点之间的最大距离。如果临界点少于两个，则返回 [-1，-1] 。

`

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func nodesBetweenCriticalPoints(head *ListNode) []int {
	var points = make([]int, 0)
	preVal := head.Val
	head = head.Next
	if head.Next == nil {
		return []int{-1, -1}
	}
	curVal := head.Val
	head = head.Next
	var i = 2
	for head != nil {
		if curVal > preVal && curVal > head.Val {
			points = append(points, i)
		}
		if curVal < preVal && curVal < head.Val {
			points = append(points, i)
		}
		i++
		preVal = curVal
		curVal = head.Val
		head = head.Next
	}
	if len(points) < 2 {
		return []int{-1, -1}
	}
	var minDistance = math.MaxInt
	for i := 1; i < len(points); i++ {
		if points[i]-points[i-1] < minDistance {
			minDistance = points[i] - points[i-1]
		}
	}
	return []int{minDistance, points[len(points)-1] - points[0]}

}

type ListNode struct {
	Val  int
	Next *ListNode
}
