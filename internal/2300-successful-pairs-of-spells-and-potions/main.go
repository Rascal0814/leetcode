package _300_successful_pairs_of_spells_and_potions

import (
	"slices"
)

// https://leetcode.cn/problems/successful-pairs-of-spells-and-potions/
func successfulPairs(spells []int, potions []int, success int64) []int {
	slices.Sort(potions)
	lowerBound := func(nums []int, target, params int) int {
		l, r := 0, len(nums)-1
		for l <= r {
			if mid := (r + l) >> 1; nums[mid]*params < target {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
		return l
	}
	var ans []int
	for _, spell := range spells {
		ans = append(ans, len(potions)-lowerBound(potions, int(success), spell))
	}
	return ans
}
