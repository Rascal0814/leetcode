package _642_furthest_building_you_can_reach

import "slices"

//https://leetcode.cn/problems/furthest-building-you-can-reach/description/
func furthestBuilding(heights []int, bricks int, ladders int) int {
	var tmp []int
	for i := 0; i < len(heights)-1; i++ {
		v := heights[i+1] - heights[i]
		if v < 0 {
			v = 0
		}
		tmp = append(tmp, v)
	}
	l, r := 0, len(heights)
	check := func(x int) bool {
		var t = make([]int, x)
		copy(t, tmp[:x])
		slices.Sort(t)
		var s int
		if ladders > x {
			return true
		}
		for _, v := range t[:x-ladders] {
			if s += v; s > bricks {
				return false
			}
		}
		return true

	}
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
