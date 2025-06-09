package _818_minimum_absolute_sum_difference

import (
	"slices"
	"sort"
)

// https://leetcode.cn/problems/minimum-absolute-sum-difference/
func minAbsoluteSumDiff(nums1 []int, nums2 []int) int {
	abs := func(a int) int {
		if a < 0 {
			return -a
		}
		return a
	}
	sortNum1 := make([]int, len(nums1))
	copy(sortNum1, nums1)
	slices.Sort(sortNum1)
	var sum, mxDiff int
	for i, v := range nums1 {
		sum += abs(v - nums2[i])
	}

	for i, v := range nums1 {
		left := sort.SearchInts(sortNum1, nums2[i])
		cutDiff := abs(v - nums2[i])
		if left < len(nums1) {
			mxDiff = max(mxDiff, cutDiff-abs(sortNum1[left]-nums2[i]))
		}
		if left != 0 {
			mxDiff = max(mxDiff, cutDiff-abs(sortNum1[left-1]-nums2[i]))
		}
	}

	return (sum - mxDiff) % (1e9 + 7)
}
