package _861_maximum_number_of_alloys

import (
	"slices"
)

//https://leetcode.cn/problems/maximum-number-of-alloys/
func maxNumberOfAlloys(n int, k int, budget int, composition [][]int, stock []int, cost []int) int {
	check := func(x int) bool {
		for _, comp := range composition {
			var need []int
			for _, v := range comp {
				need = append(need, v*x)
			}
			for i, v := range stock {
				need[i] -= v
				need[i] = max(need[i], 0)
			}

			var money int
			for i, v := range need {
				money += v * cost[i]
			}
			if money <= budget {
				return true
			}
		}
		return false
	}
	l, r := 0, slices.Min(stock)+budget+1
	for l < r-1 {
		mid := (l + r) >> 1
		if check(mid) {
			l = mid
		} else {
			r = mid
		}
	}
	return l
}
