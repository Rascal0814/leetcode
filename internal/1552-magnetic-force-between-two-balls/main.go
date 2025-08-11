package _552_magnetic_force_between_two_balls

import "slices"

//https://leetcode.cn/problems/magnetic-force-between-two-balls/
func maxDistance(position []int, m int) int {
	slices.Sort(position)
	l, r := 0, slices.Max(position)+1
	check := func(x int) bool {
		cnt, last := 1, position[0]
		for i := 1; i < len(position); i++ {
			if position[i]-last >= x {
				last = position[i]
				cnt++
			}
		}
		if cnt >= m {
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
