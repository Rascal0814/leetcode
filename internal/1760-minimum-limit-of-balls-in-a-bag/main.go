package _760_minimum_limit_of_balls_in_a_bag

import "slices"

// https://leetcode.cn/problems/minimum-limit-of-balls-in-a-bag/
func minimumSize(nums []int, maxOperations int) int {
	l, r := 0, slices.Max(nums)
	check := func(x int) bool {
		cnt := 0
		for _, num := range nums {
			cnt += num/x - 1
			if num%x > 0 {
				cnt++
			}
		}
		return cnt > maxOperations
	}
	for l+1 < r {
		if mid := (l + r) >> 1; check(mid) {
			l = mid
		} else {
			r = mid
		}
	}
	return r
}
