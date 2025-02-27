package _25_valid_palindrome

import "strings"

const _ = `125. 验证回文串

如果在将所有大写字符转换为小写字符、并移除所有非字母数字字符之后，短语正着读和反着读都一样。则可以认为该短语是一个 回文串 。

字母和数字都属于字母数字字符。

给你一个字符串 s，如果它是 回文串 ，返回 true ；否则，返回 false 。`

func isPalindrome(s string) bool {
	s = strings.ToLower(s)

	l, r := 0, len(s)-1
	for l <= r {

		if !(s[l] >= 'a' && s[l] <= 'z') && !(s[l] <= '9' && s[l] >= '0') {
			l++
			continue
		}
		if !(s[r] >= 'a' && s[r] <= 'z') && !(s[r] <= '9' && s[r] >= '0') {
			r--
			continue
		}
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}

	return true
}
