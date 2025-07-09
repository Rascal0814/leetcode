package _187_minimum_time_to_complete_trips

import "slices"

// https://leetcode.cn/problems/minimum-time-to-complete-trips/submissions/637758973/
func minimumTime(time []int, totalTrips int) int64 {
	l := slices.Min(time)
	r := l * totalTrips

	for l <= r { // [left,right]
		mid := (l + r) >> 1
		count := 0
		for _, v := range time {
			count += mid / v
		}

		if count >= totalTrips {
			r = mid - 1 // [left,mid-1]
		} else {
			l = mid + 1 // [mid+1,right]
		}
	}
	return int64(l)
}
