package _439_minimize_maximum_of_array

import "slices"

// https://leetcode.cn/problems/minimize-maximum-of-array/
func minimizeArrayValue(nums []int) int {
	l, r := 0, slices.Max(nums)

	check := func(x int) bool {
		step := 0
		for i := len(nums) - 1; i >= 0; i-- {
			if nums[i] < x {
				step += nums[i] - x
			} else {
				step += nums[i] - x
			}
			step = max(step, 0)
		}
		if step > 0 {
			return false
		}
		return true
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
