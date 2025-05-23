package _537_count_the_number_of_good_subarrays

// https://leetcode.cn/problems/count-the-number-of-good-subarrays/
func countGood(nums []int, k int) int64 {
	l, ans, cnt, pairs := 0, 0, make(map[int]int), 0
	for _, v := range nums {
		pairs += cnt[v]
		cnt[v]++
		for pairs >= k {
			cnt[nums[l]]--
			pairs -= cnt[nums[l]]
			l++
		}

		ans += l
	}
	return int64(ans)
}
