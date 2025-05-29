package _389_longest_subsequence_with_limited_sum

import "slices"

// https://leetcode.cn/problems/longest-subsequence-with-limited-sum/
func answerQueries(nums []int, queries []int) []int {
	lowerBound := func(n []int, target int) int {
		l, r := 0, len(n)-1
		for l <= r {
			if mid := (r + l) >> 1; n[mid] < target {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
		return l
	}
	//window := func(n []int, target int) int {
	//	var l, cnt, ans int
	//	for r, v := range n {
	//		cnt += v
	//		for cnt > target {
	//			cnt -= n[l]
	//			l++
	//		}
	//
	//		ans = max(r-l+1, ans)
	//	}
	//	return ans
	//}
	slices.Sort(nums)
	minf := func(n []int, target int) int {
		var cnt int
		for r, v := range n {
			if cnt += v; cnt > target {
				return r
			}
		}
		return len(n)
	}

	var ans []int
	for _, v := range queries {
		ans = append(ans, minf(nums, v))
	}
	return ans
}
