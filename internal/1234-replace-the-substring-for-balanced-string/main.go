package _234_replace_the_substring_for_balanced_string

// https://leetcode.cn/problems/replace-the-substring-for-balanced-string/
func balancedString(s string) int {
	m := len(s) / 4
	cnt := ['X']int{}
	for _, c := range s {
		cnt[c]++
	}
	if cnt['Q'] == m && cnt['W'] == m && cnt['E'] == m && cnt['R'] == m {
		return 0
	}
	l, ans := 0, len(s)
	for r, v := range s {
		cnt[v]--
		for cnt['Q'] <= m && cnt['W'] <= m && cnt['E'] <= m && cnt['R'] <= m { // 外部的需满足4个字符同时小于m s[l:r+1] 这里面的需要替换的字符串
			ans = min(ans, r-l+1)
			cnt[s[l]]++
			l++

		}
	}
	return ans

}
