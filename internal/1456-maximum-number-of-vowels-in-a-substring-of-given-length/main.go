package _456_maximum_number_of_vowels_in_a_substring_of_given_length

// https://leetcode.cn/problems/maximum-number-of-vowels-in-a-substring-of-given-length/
func maxVowels(s string, k int) int {
	isVowel := func(vowel uint8) int {
		if vowel == 'a' || vowel == 'e' || vowel == 'i' || vowel == 'o' || vowel == 'u' {
			return 1
		}
		return 0
	}
	maxVowel := 0
	for i := 0; i < k; i++ {
		maxVowel += isVowel(s[i])
	}
	temp := maxVowel
	for i := k; i < len(s); i++ {
		if maxVowel >= k {
			return maxVowel
		}
		temp += isVowel(s[i]) - isVowel(s[i-k])
		if temp > maxVowel {
			maxVowel = temp
		}
	}
	return maxVowel
}
