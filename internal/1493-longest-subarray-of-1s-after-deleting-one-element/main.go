package _493_longest_subarray_of_1s_after_deleting_one_element

// https://leetcode.cn/problems/longest-subarray-of-1s-after-deleting-one-element/
func longestSubarray(nums []int) int {
	l, n, zeroCount, mx := 0, len(nums), 0, 0
	for r := 0; r < n; r++ {
		if nums[r] == 0 {
			zeroCount++
		}
		for zeroCount > 1 {
			if nums[l] == 0 {
				zeroCount--
			}
			l++
		}

		mx = max(mx, r-l)
	}
	return mx
}
