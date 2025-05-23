package _325_count_substrings_with_k_frequency_characters_i

// https://leetcode.cn/problems/count-substrings-with-k-frequency-characters-i/
func numberOfSubstrings(s string, k int) int {
	l, ans, cnt := 0, 0, [26]int{}
	for _, v := range s {
		cnt[v-'a']++
		for cnt[v-'a'] >= k {
			cnt[s[l]-'a']--
			l++
		}
		ans += l

	}
	return ans
}
