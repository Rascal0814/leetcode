package _064_minimized_maximum_of_products_distributed_to_any_store

import "slices"

//https://leetcode.cn/problems/minimized-maximum-of-products-distributed-to-any-store/
func minimizedMaximum(n int, quantities []int) int {
	l, r := -1, slices.Max(quantities)
	check := func(x int) bool {
		if x == 0 {
			return true
		}
		cnt := 0
		for _, quantity := range quantities {
			t := quantity / x
			cnt += t
			if quantity%x > 0 {
				cnt++
			}
		}

		return cnt > n

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
