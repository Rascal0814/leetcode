package _298_count_substrings_that_can_be_rearranged_to_contain_a_string_ii

// https://leetcode.cn/problems/count-substrings-that-can-be-rearranged-to-contain-a-string-ii/
func validSubstringCount(word1 string, word2 string) int64 {
	l, cnt1, cnt2, ans := 0, [26]int{}, [26]int{}, 0

	for _, v := range word2 {
		cnt2[v-'a']++
	}
	cover := func(a [26]int, b [26]int) bool {
		for i, v := range b {
			if a[i] < v {
				return false
			}
		}
		return true
	}
	for _, v := range word1 {
		cnt1[v-'a']++

		for cover(cnt1, cnt2) {
			cnt1[word1[l]-'a']--
			l++
		}

		ans += l
	}
	return int64(ans)
}
