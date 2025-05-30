package main

import "fmt"

func main() {
	s := "aaaa"
	fmt.Println(test(s))
	//fmt.Println(s[0:2])
}

func longestPalindrome(s string) string {
	for i := 0; i < len(s)-1; i++ {
		for j := i + 1; j < len(s); j++ {
			if j-i > 1 {
				fmt.Println(s[i:j])
			}
		}
	}

	return ""
}

func test(s string) string {
	res, dp := "", make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(s))
	}
	for i := len(s) - 1; i >= 0; i-- {
		for j := i; j < len(s); j++ {
			dp[i][j] = (s[i] == s[j]) && ((j-i < 3) || dp[i+1][j-1])
			if dp[i][j] && (res == "" || j-i+1 > len(res)) {
				res = s[i : j+1]
			}
		}
	}
	return res
}
