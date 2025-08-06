package _560_house_robber_iv

import "slices"

//https://leetcode.cn/problems/house-robber-iv/
func minCapability(nums []int, k int) int {
	l, r := -1, slices.Max(nums)
	check := func(x int) bool {
		cnt := 0
		for i := 0; i < len(nums); i++ {
			if nums[i] <= x {
				cnt++
				i++
			}
		}
		if cnt >= k {
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
