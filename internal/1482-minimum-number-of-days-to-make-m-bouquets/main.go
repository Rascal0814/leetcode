package _482_minimum_number_of_days_to_make_m_bouquets

import "slices"

// https://leetcode.cn/problems/minimum-number-of-days-to-make-m-bouquets/
func minDays(bloomDay []int, m int, k int) int {
	l, r := 1, slices.Max(bloomDay)
	day := func(t int) int {
		num, tmp := 0, 0
		for _, v := range bloomDay {
			if v <= t { // 开花了
				tmp++
			} else {
				tmp = 0
			}
			if tmp == k {
				num++
				tmp = 0
			}
		}
		return num
	}
	for l <= r {
		mid := (l + r) >> 1
		if day(mid) >= m {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	if l > slices.Max(bloomDay) {
		return -1
	}
	return l
}
