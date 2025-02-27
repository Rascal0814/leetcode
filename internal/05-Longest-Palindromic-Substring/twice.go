package main

import "fmt"

// 动态规划
// 想要是回文子串必须头尾相等 也就是 s[i] = s[j] 长度为3时两边相等肯定是回文字串
func longestPalindromeTwice(s string) string {
	res, dp := "", make([][]bool, len(s))
	for i, _ := range s {
		dp[i] = make([]bool, len(s))
		dp[i][i] = true
	}
	// 由于需要dp[i+1] 所以必须从后向前遍历
	for i := len(s) - 1; i >= 0; i-- {
		for j := i; j < len(s); j++ {
			fmt.Println(s[i : j+1])
			// 将每个回文子串找出来 dp[i+1][j-1] 肯定在这轮之前
			if s[i] == s[j] && ((j-i < 3) || dp[i+1][j-1]) {
				dp[i][j] = true
			}
			// 找出最长的回文字串
			if dp[i][j] && j-i > len(res)-1 {
				res = s[i : j+1]
			}
		}
	}
	return res
}
