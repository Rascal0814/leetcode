package _594_minimum_time_to_repair_cars

import (
	"math"
	"slices"
)

// https://leetcode.cn/problems/minimum-time-to-repair-cars/
func repairCars(ranks []int, cars int) int64 {
	l, r := 1, slices.Max(ranks)*cars*cars
	fixNum := func(t int) int {
		sum := 0
		for _, rank := range ranks {
			sum += int(math.Sqrt(float64(t) / float64(rank)))
		}
		return sum
	}
	for l <= r {
		mid := (l + r) >> 1
		if fixNum(mid) >= cars {
			r = mid - 1
		} else {
			l = mid + 1
		}

	}
	return int64(l)
}
