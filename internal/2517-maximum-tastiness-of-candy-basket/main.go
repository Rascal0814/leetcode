package _517_maximum_tastiness_of_candy_basket

import "slices"

//https://leetcode.cn/problems/maximum-tastiness-of-candy-basket/
func maximumTastiness(price []int, k int) int {
	slices.Sort(price)
	l, r := 0, price[len(price)-1]+1
	check := func(x int) bool {
		cnt, last := 1, price[0]
		for i := 1; i < len(price); i++ {
			t := price[i] - last
			if t >= x {
				last = price[i]
				cnt++
			}
		}
		if cnt >= k {
			return true
		}
		return false
	}
	for l+1 < r {
		if mid := (l + r) >> 1; check(mid) {
			l = mid
		} else {
			r = mid
		}
	}
	return l
}
