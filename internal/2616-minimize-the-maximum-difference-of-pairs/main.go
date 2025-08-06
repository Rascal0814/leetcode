package _616_minimize_the_maximum_difference_of_pairs

import (
	"slices"
)

//https://leetcode.cn/problems/minimize-the-maximum-difference-of-pairs/
func minimizeMax(nums []int, p int) int {
	l, r := -1, 1_0000_0000_0
	abs := func(a int) int {
		if a < 0 {
			return -a
		}
		return a
	}

	slices.Sort(nums)
	check := func(x int) bool {
		cnt := 0
		for i := 0; i < len(nums)-1; i++ {
			if abs(nums[i+1]-nums[i]) <= x {
				cnt++
				i++
			}
		}
		if cnt >= p {
			return true
		}
		return false
	}
	for l+1 < r {
		if mid := (l + r) >> 1; check(mid) {
			r = mid
		} else {
			l = mid
		}
	}
	return r
}
