package _8_length_of_last_word

import "strings"

const _ = `
58. 最后一个单词的长度

给你一个字符串 s，由若干单词组成，单词前后用一些空格字符隔开。返回字符串中 最后一个 单词的长度。

单词 是指仅由字母组成、不包含任何空格字符的最大
子字符串。
`

func lengthOfLastWord(s string) int {
	split := strings.Split(strings.TrimSpace(s), " ")
	return len(split[len(split)-1])
}

func lengthOfLastWord2(s string) int {
	var count int
	var fistWord bool
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ' ' && !fistWord {
			continue
		}
		if s[i] == ' ' && fistWord {
			break
		}
		fistWord = true
		count++
	}
	return count
}
