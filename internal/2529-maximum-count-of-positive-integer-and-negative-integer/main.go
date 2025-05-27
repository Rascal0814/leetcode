package _529_maximum_count_of_positive_integer_and_negative_integer

// https://leetcode.cn/problems/maximum-count-of-positive-integer-and-negative-integer/
func maximumCount(nums []int) int {
	lowerBound := func(list []int, target int) int {
		l, r := 0, len(list)-1
		for l <= r {
			if mid := (r + l) >> 1; list[mid] < target {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
		return l
	}

	neg := lowerBound(nums, 0)
	pos := lowerBound(nums, 1)
	if pos >= len(nums) {
		pos = 0
	} else {
		pos = len(nums) - pos
	}

	return max(neg, pos)
}
