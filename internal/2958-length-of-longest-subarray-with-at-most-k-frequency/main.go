package _958_length_of_longest_subarray_with_at_most_k_frequency

// https://leetcode.cn/problems/length-of-longest-subarray-with-at-most-k-frequency/
func maxSubarrayLength(nums []int, k int) int {
	l, ans, cnt := 0, 0, make(map[int]int)
	for r, v := range nums {
		cnt[v]++
		for cnt[v] > k {
			cnt[nums[l]]--
			l++
		}
		ans = max(ans, r-l+1)
	}

	return ans
}
