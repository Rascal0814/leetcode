package _648_sell_diminishing_valued_colored_balls

import "slices"

//https://leetcode.cn/problems/sell-diminishing-valued-colored-balls/
func maxProfit(inventory []int, orders int) int {
	l, r := 0, slices.Max(inventory)+1
	check := func(k int) bool {
		num := 0
		for _, v := range inventory {
			num += max(v-k, 0)
		}

		return num >= orders
	}
	for l+1 < r {
		if mid := (l + r) >> 1; check(mid) {
			l = mid
		} else {
			r = mid
		}
	}

	res := 0
	for _, v := range inventory {
		if v >= r {
			orders -= v - r + 1
			res += (v + r) * (v - r + 1) / 2
		}
	}
	if orders > 0 {
		res += orders * l
	} else {
		res += orders * r
	}

	return res % 1_0000_0000_7
}
