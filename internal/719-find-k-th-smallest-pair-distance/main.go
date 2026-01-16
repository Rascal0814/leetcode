package _19_find_k_th_smallest_pair_distance

import (
	"fmt"
	"slices"
)

//https://leetcode.cn/problems/find-k-th-smallest-pair-distance/
func smallestDistancePair(nums []int, k int) int {
	slices.Sort(nums)
	check := func(x int) bool {
		num, j := 0, 0
		for i, v := range nums {
			for v-nums[j] > x {
				fmt.Printf("(%d,%d)\n", j, i)
				j++
			}
			num += i - j
		}
		return num >= k
	}
	l, r := 0, nums[len(nums)-1]
	for l < r {
		if mid := (l + r) >> 1; check(mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return r
}
