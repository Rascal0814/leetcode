package _576_find_the_maximum_number_of_marked_indices

import (
	"slices"
)

// https://leetcode.cn/problems/find-the-maximum-number-of-marked-indices/
func maxNumOfMarkedIndices(nums []int) int {
	slices.Sort(nums)

	var l, r = 0, len(nums) / 2

	check := func(cup int) bool {
		for i := 0; i < cup; i++ {
			if 2*nums[i] > nums[len(nums)-cup+i] {
				return false
			}
		}
		return true
	}

	for l <= r {
		if mid := (l + r) >> 1; check(mid) {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	return r * 2
}
