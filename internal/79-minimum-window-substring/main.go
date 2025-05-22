package _9_minimum_window_substring

// https://leetcode.cn/problems/minimum-window-substring/
func minWindow(s string, t string) string {
	tCnt := [123]int{}
	for _, v := range t {
		tCnt[v]++
	}
	cover := func(a [123]int, b [123]int) bool {
		for i, v := range b {
			if a[i] < v {
				return false
			}
		}
		return true
	}
	l, ans, SCnt := 0, s+"a", [123]int{}
	for r, v := range s {
		SCnt[v]++
		for cover(SCnt, tCnt) {
			if len(s[l:r+1]) < len(ans) {
				ans = s[l : r+1]
			}
			tCnt[s[l]]++
			l++
		}

	}

	if ans == s+"a" {
		return ""
	}

	return ans

}
