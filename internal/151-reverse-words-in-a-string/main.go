package _51_reverse_words_in_a_string

import (
	"slices"
	"strings"
)

const _ = `
151. 反转字符串中的单词

给你一个字符串 s ，请你反转字符串中 单词 的顺序。

单词 是由非空格字符组成的字符串。s 中使用至少一个空格将字符串中的 单词 分隔开。

返回 单词 顺序颠倒且 单词 之间用单个空格连接的结果字符串。

注意：输入字符串 s中可能会存在前导空格、尾随空格或者单词间的多个空格。返回的结果字符串中，单词间应当仅用单个空格分隔，且不包含任何额外的空格。`

func reverseWords(s string) string {
	split := strings.Split(strings.TrimSpace(s), " ")
	slices.Reverse(split)
	var res []string
	for i := 0; i < len(split); i++ {
		if split[i] == "" {
			continue
		}
		res = append(res, split[i])
	}
	return strings.Join(res, " ")
}
func reverseWords2(s string) string {
	var res []string
	n := len(s)
	var words string
	for i := n - 1; i >= 0; i-- {
		if s[i] == ' ' {
			if words != "" {
				res = append(res, words)
				words = ""
			}
			continue
		}
		words = s[i:i+1] + words
	}
	if words != "" {
		res = append(res, words)
	}
	switch len(res) {
	case 0:
		return ""
	case 1:
		return res[0]
	}

	var ans = res[0]
	for _, re := range res[1:] {
		ans += " " + re
	}
	return ans
}
