package _528_maximize_the_minimum_powered_city

import (
	"slices"
)

//https://leetcode.cn/problems/maximize-the-minimum-powered-city/
func maxPower(stations []int, r int, k int) int64 {
	// 定长滑窗
	var power = make([]int, len(stations))
	for _, v := range stations[:r+1] {
		power[0] += v
	}

	sum := power[0]
	for i := 1; i < len(stations); i++ {
		if i+r < len(stations) {
			sum += stations[i+r]
		}
		if i-r-1 >= 0 {
			sum -= stations[i-r-1]
		}
		power[i] = sum
	}

	left, right := slices.Min(power), slices.Max(power)
	if left == right {
		return int64(left + k)
	}

	check := func(x int) bool {
		for i := 0; i < len(power); i++ {
			ll, rr := i-r, min(i+r, len(power)-1)
			if i-r < 0 {
				ll = 0
			}
			if slices.Min(power[ll:rr+1])+k < x {
				return false
			}
		}
		return true
	}
	for left+1 < right {
		if mid := (left + right) >> 1; check(mid) {
			left = mid
		} else {
			right = mid
		}

	}
	return int64(left)
}
