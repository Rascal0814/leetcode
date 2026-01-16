package _476_closest_nodes_queries_in_a_binary_search_tree

import (
	"fmt"
	"slices"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func closestNodes(root *TreeNode, queries []int) [][]int {
	arr := bfs(root)
	slices.Sort(arr)
	fmt.Print(arr)

	var res [][]int
	for _, v := range queries {
		i, j := 0, len(arr)
		var mx, mn int

		for i+1 < j {
			mid := (i + j) >> 1
			if arr[mid] <= v {
				i = mid
			} else {
				j = mid
			}
		}
		if arr[i] <= v {
			mn = arr[i]
		} else {
			mn = -1
		}

		i, j = -1, len(arr)-1
		for i+1 < j {
			mid := (i + j) >> 1
			if arr[mid] >= v {
				j = mid
			} else {
				i = mid
			}
		}
		if arr[j] >= v {
			mx = arr[j]
		} else {
			mx = -1
		}

		res = append(res, []int{mn, mx})
	}
	return res
}

func bfs(root *TreeNode) []int {
	var res []int
	var queue []*TreeNode
	queue = append(queue, root)
	for len(queue) != 0 {
		tmp := queue[0]
		res = append(res, tmp.Val)
		if tmp.Left != nil {
			queue = append(queue, tmp.Left)
		}
		if tmp.Right != nil {
			queue = append(queue, tmp.Right)
		}
		queue = queue[1:]
	}
	return res
}
