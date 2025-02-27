package main

import "fmt"

func main() {
	fmt.Println(two(""))
}

// 	滑动窗口：left，right 左右窗口，若右边元素在bucket中存在的话就左边滑动1，若不存在右边滑动1
func lengthOfLongestSubstring(s string) int {

	if len(s) == 0 {
		return 0
	}

	left, right, maxString := 0, 0, 0

	var bucket [256]bool

	for left < len(s) && right < len(s) {

		if bucket[s[right]] {
			bucket[s[left]] = false
			left++
		} else {
			bucket[s[right]] = true
			right++
		}

		if maxString < right-left {
			maxString = right - left
		}

	}
	return maxString
}
