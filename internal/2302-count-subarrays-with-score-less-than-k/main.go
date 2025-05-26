package _302_count_subarrays_with_score_less_than_k

// https://leetcode.cn/problems/count-subarrays-with-score-less-than-k/
func countSubarrays(nums []int, k int64) int64 {
	l, cnt, ans := 0, 0, 0
	for r, v := range nums {
		cnt += v
		for cnt*(r-l+1) >= int(k) {
			cnt -= nums[l]
			l++
		}
		ans += r - l + 1
	}
	return int64(ans)
}
