package _75_koko_eating_bananas

import "slices"

// https://leetcode.cn/problems/koko-eating-bananas/
func minEatingSpeed(piles []int, h int) int {
	l, r := 1, slices.Max(piles)

	need := func(x int) int {
		var c int
		for _, pile := range piles {
			c += pile / x
			if pile%x != 0 {
				c++
			}
		}
		return c
	}

	for l <= r {
		mid := (l + r) >> 1
		if need(mid) <= h {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return l
}
