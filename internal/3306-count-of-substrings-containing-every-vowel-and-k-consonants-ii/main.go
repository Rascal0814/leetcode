package _306_count_of_substrings_containing_every_vowel_and_k_consonants_ii

// https://leetcode.cn/problems/count-of-substrings-containing-every-vowel-and-k-consonants-ii/
func countOfSubstrings(word string, k int) int64 {
	var l1, l2, cnt1, cnt2, ans int
	cnt1Map, cnt2Map := [26]int{}, [26]int{}
	for _, v := range word {
		cnt1Map[v-'a']++
		cnt2Map[v-'a']++
		if v != 'a' && v != 'e' && v != 'i' && v != 'o' && v != 'u' {
			cnt1++
			cnt2++
		}
		for cnt1 >= k && cnt1Map['a'-'a'] >= 1 && cnt1Map['e'-'a'] >= 1 && cnt1Map['i'-'a'] >= 1 && cnt1Map['o'-'a'] >= 1 && cnt1Map['u'-'a'] >= 1 {
			if word[l1] != 'a' && word[l1] != 'e' && word[l1] != 'i' && word[l1] != 'o' && word[l1] != 'u' {
				cnt1--
			}
			cnt1Map[word[l1]-'a']--
			l1++
		}
		for cnt2 >= k+1 && cnt2Map['a'-'a'] >= 1 && cnt2Map['e'-'a'] >= 1 && cnt2Map['i'-'a'] >= 1 && cnt2Map['o'-'a'] >= 1 && cnt2Map['u'-'a'] >= 1 {
			if word[l2] != 'a' && word[l2] != 'e' && word[l2] != 'i' && word[l2] != 'o' && word[l2] != 'u' {
				cnt2--
			}
			cnt2Map[word[l2]-'a']--
			l2++
		}
		ans += l1 - l2
	}
	return int64(ans)
}
