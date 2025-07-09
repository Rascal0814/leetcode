package _283_find_the_smallest_divisor_given_a_threshold

import "slices"

func smallestDivisor(nums []int, threshold int) int {
	check := func(n int) bool {
		var sum int
		for _, v := range nums {
			sum += (n + v - 1) / n
			if sum > threshold {
				return false
			}
		}
		return true

	}
	l, r := 1, slices.Max(nums) // 除数只能从 1 开始
	for l <= r {
		mid := (l + r) >> 1
		if check(mid) {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return l
}
