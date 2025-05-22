package _904_shortest_and_lexicographically_smallest_beautiful_string

import "strings"

// https://leetcode.cn/problems/shortest-and-lexicographically-smallest-beautiful-string/
func shortestBeautifulSubstring(s string, k int) string {
	if strings.Count(s, "1") < k {
		return ""
	}
	l, ans, cnt := 0, s, 0
	for r, v := range s {
		if v == '1' {
			cnt++
		}
		for cnt > k || s[l] == '0' {
			if s[l] == '1' {
				cnt--
			}
			l++
		}
		if cnt == k && (len(s[l:r+1]) < len(ans) || (len(s[l:r+1]) == len(ans) && s[l:r+1] < ans)) {
			ans = s[l : r+1]
		}
	}
	return ans
}
