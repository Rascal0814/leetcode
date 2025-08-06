package _10_split_array_largest_sum

import "slices"

//https://leetcode.cn/problems/split-array-largest-sum/
func splitArray(nums []int, k int) int {
	l, r := slices.Max(nums)-1, 0
	for _, v := range nums {
		r += v
	}
	check := func(x int) bool {
		sum, cnt := 0, 1
		for _, num := range nums {
			if sum+num > x {
				sum = num
				cnt++
				continue
			}
			sum += num
		}
		return cnt <= k
	}
	//右闭左开
	for l+1 < r {
		if mid := (l + r) >> 1; check(mid) {
			r = mid
		} else {
			l = mid
		}
	}

	return r
}

//[7,2,5,10,8]
