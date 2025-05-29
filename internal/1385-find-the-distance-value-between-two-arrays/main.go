package _385_find_the_distance_value_between_two_arrays

import (
	"slices"
)

// https://leetcode.cn/problems/find-the-distance-value-between-two-arrays/
func findTheDistanceValue(arr1 []int, arr2 []int, d int) int {
	var ans int
	slices.Sort(arr2)
	lowerBound := func(nums []int, target int) int {
		l, r := 0, len(nums)-1
		for l <= r {
			if mid := (r + l) >> 1; nums[mid] < target {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
		return l
	}

	for _, v := range arr1 {
		t := lowerBound(arr2, v-d)
		if t >= len(arr2) || arr2[t] > v+d {
			ans++
		}
	}
	return ans
}
