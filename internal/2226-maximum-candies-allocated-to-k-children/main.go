package _226_maximum_candies_allocated_to_k_children

import "slices"

// https://leetcode.cn/problems/maximum-candies-allocated-to-k-children/
func maximumCandies(candies []int, k int64) int {
	l, r := 1, slices.Max(candies)
	check := func(x int) int64 {
		sum := 0
		for _, v := range candies {
			sum += v / x
		}
		return int64(sum)
	}
	for l <= r {
		if mid := (l + r) >> 1; check(mid) < k {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return r
}
