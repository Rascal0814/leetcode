package _090_maximum_length_substring_with_two_occurrences

// https://leetcode.cn/problems/maximum-length-substring-with-two-occurrences/
func maximumLengthSubstring(s string) int {
	l, n, countMap, mx := 0, len(s), make(map[uint8]int), 0

	for r := 0; r < n; r++ {
		countMap[s[r]] += 1
		for countMap[s[r]] > 2 {
			countMap[s[l]] -= 1
			l++
		}

		mx = max(mx, r-l+1)
	}

	return mx

}
