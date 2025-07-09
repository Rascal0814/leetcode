package _75_heaters

import (
	"slices"
)

// https://leetcode.cn/problems/heaters/
func findRadius(houses []int, heaters []int) int {
	slices.Sort(heaters)
	var res int
	lowerBound := func(x int) int {
		l, r := 0, len(heaters)-1
		for l <= r {
			if mid := (l + r) >> 1; heaters[mid] <= x {
				l = mid + 1
			} else {
				r = mid - 1
			}

		}
		return l
	}
	for _, house := range houses {
		r := lowerBound(house) // 距离house 最近的热水器
		l := r - 1
		temp := 0
		if r >= len(heaters) {
			temp = house - heaters[l]
		} else if l < 0 {
			temp = heaters[r] - house
		} else {
			temp = min(house-heaters[l], heaters[r]-house)
		}
		res = max(res, temp)
	}
	return res
}
