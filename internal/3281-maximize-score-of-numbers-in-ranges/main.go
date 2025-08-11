package _281_maximize_score_of_numbers_in_ranges

import "slices"

//https://leetcode.cn/problems/maximize-score-of-numbers-in-ranges/
func maxPossibleScore(start []int, d int) int {
	slices.Sort(start)
	l, r := 0, slices.Max(start)+d+1

	check := func(x int) bool {
		lastChoose := start[0]
		for i := 1; i < len(start); i++ {
			lastChoose += x
			if lastChoose < start[i] { // 若没在区间内就选择区间内的最小值
				lastChoose = start[i]
			}

			if lastChoose > start[i]+d { // 若没在区间外就说明这个值不符合
				return false
			}
		}

		return true
	}
	for l+1 < r {
		if mid := (l + r) >> 1; check(mid) {
			l = mid
		} else {
			r = mid
		}
	}
	return l

}
