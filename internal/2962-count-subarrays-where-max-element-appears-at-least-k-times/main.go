package _962_count_subarrays_where_max_element_appears_at_least_k_times

import (
	"slices"
)

// https://leetcode.cn/problems/count-subarrays-where-max-element-appears-at-least-k-times/
func countSubarrays(nums []int, k int) int64 {
	maxVal := slices.Max(nums)
	l, cnt, ans := 0, 0, 0
	for _, v := range nums {
		if v == maxVal {
			cnt++
		}
		for cnt >= k {
			if nums[l] == maxVal {
				cnt--
			}
			l++
		}
		ans += l
	}
	return int64(ans)
}
