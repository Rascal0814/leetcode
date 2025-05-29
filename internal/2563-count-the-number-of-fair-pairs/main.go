package _563_count_the_number_of_fair_pairs

import (
	"slices"
	"sort"
)

// https://leetcode.cn/problems/count-the-number-of-fair-pairs/
func countFairPairs(nums []int, lower int, upper int) int64 {
	var ans int
	if len(nums) <= 1 {
		return 0
	}
	slices.Sort(nums)
	for i := 1; i < len(nums); i++ {
		lowerIndex := sort.SearchInts(nums[i:], lower-nums[i-1])
		upperIndex := sort.SearchInts(nums[i:], upper-nums[i-1]+1)
		if upperIndex >= len(nums[i:]) {
			upperIndex = len(nums[i:])
		}
		ans += upperIndex - lowerIndex
	}
	return int64(ans)
}
