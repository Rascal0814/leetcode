package _004_max_consecutive_ones_iii

// https://leetcode.cn/problems/max-consecutive-ones-iii/
func longestOnes(nums []int, k int) int {
	var l, ans, zeroCount int
	for r, v := range nums {
		if v == 0 {
			zeroCount++
		}

		for zeroCount > k {
			if nums[l] == 0 {
				zeroCount--
			}
			l++
		}

		ans = max(ans, r-l+1)
	}
	return ans
}
