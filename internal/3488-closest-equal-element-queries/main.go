package _488_closest_equal_element_queries

import (
	"slices"
	"sort"
)

// https://leetcode.cn/problems/closest-equal-element-queries/
func solveQueries(nums []int, queries []int) []int {
	position := make(map[int][]int)
	for i, v := range nums {
		position[v] = append(position[v], i)
	}

	for i, v := range position {
		n := len(v)
		if n > 1 {
			p0 := v[0]
			v = slices.Insert(v, 0, v[n-1]-len(nums))
			position[i] = append(v, p0+len(nums))
		}
	}

	var ans []int
	for _, v := range queries {
		arr := position[nums[v]]
		if len(arr) == 1 {
			ans = append(ans, -1)
		} else {
			l := sort.SearchInts(arr, v) // 找到下标位置 左右都是最近的值
			ans = append(ans, min(arr[l+1]-v, v-arr[l-1]))
		}

	}
	return ans
}
