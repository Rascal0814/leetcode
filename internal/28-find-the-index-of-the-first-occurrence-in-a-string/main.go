package _8_find_the_index_of_the_first_occurrence_in_a_string

import (
	"strings"
)

const _ = `
28. 找出字符串中第一个匹配项的下标

给你两个字符串 haystack 和 needle ，请你在 haystack 字符串中找出 needle 字符串的第一个匹配项的下标（下标从 0 开始）。如果 needle 不是 haystack 的一部分，则返回  -1 。`

func strStr(haystack string, needle string) int {
	return strings.Index(haystack, needle)
}

func strStr2(haystack string, needle string) int {
	if len(needle) > len(haystack) {
		return -1
	}
	var indexList []int
	for i := 0; i < len(haystack); i++ {
		if haystack[i] == needle[0] {
			indexList = append(indexList, i)
		}
	}

	for i := 0; i < len(indexList); i++ {
		index := indexList[i]
		if len(haystack[index:]) < len(needle) {
			continue
		}
		if haystack[index:len(needle)+index] == needle {
			return indexList[i]
		}
	}
	return -1
}
