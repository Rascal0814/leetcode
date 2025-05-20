package _3_Longest_Substring_Without_Repeating_Characters

// https://leetcode.cn/problems/longest-substring-without-repeating-characters/description/
func lengthOfLongestSubstring(s string) int {
	if len(s) <= 0 {
		return len(s)
	}
	l, r, n, exitMap, res := 0, 0, len(s), make(map[uint8]bool), 0
	for r < n {
		if exitMap[s[r]] {
			exitMap[s[l]] = false
			l++
		} else {
			exitMap[s[r]] = true
			r++
		}
		if r-l > res {
			res = r - l
		}
	}

	return res
}
