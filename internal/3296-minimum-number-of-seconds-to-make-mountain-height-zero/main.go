package _296_minimum_number_of_seconds_to_make_mountain_height_zero

// https://leetcode.cn/problems/minimum-number-of-seconds-to-make-mountain-height-zero/
func minNumberOfSeconds(mountainHeight int, workerTimes []int) int64 {
	var l, r int // 需要多少时间
	for _, v := range workerTimes {

		if r == 0 || r < v {
			r = v
		}
	}
	r = r * mountainHeight

	a := func(x int) int {

	}

	for l <= r {
		mid := (l + r) >> 1
		if a(mid) >= mountainHeight {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	return int64(l)
}
