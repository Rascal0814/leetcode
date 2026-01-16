package _9_sqrtx

import "math"

func mySqrt(x int) int {
	l, r := -1, min(int(math.Sqrt(math.MaxInt32)), x)+1
	for l+1 < r {
		mid := (l + r) >> 1
		if mid*mid > x {
			r = mid
		} else {
			l = mid
		}
	}
	return l
}
