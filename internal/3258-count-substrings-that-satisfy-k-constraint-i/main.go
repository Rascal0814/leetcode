package _258_count_substrings_that_satisfy_k_constraint_i

// https://leetcode.cn/problems/count-substrings-that-satisfy-k-constraint-i/
func countKConstraintSubstrings(s string, k int) int {
	l, cnt, ans := 0, [2]int{}, 0
	for r, v := range s {
		cnt[v-'0']++
		for cnt[0] > k && cnt[1] > k {
			cnt[s[l]-'0']--
			l++
		}
		ans += r - l + 1
	}
	return ans
}
