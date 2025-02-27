package _4_longest_common_prefix

import (
	"fmt"
	"sort"
)

const _ = `
14. 最长公共前缀

编写一个函数来查找字符串数组中的最长公共前缀。

如果不存在公共前缀，返回空字符串 ""。

`

func longestCommonPrefix(strs []string) string {
	fmt.Println(strs)
	var prefix = strs[0]
	for i := 1; i < len(strs); i++ {
		if len(prefix) == 0 {
			return ""
		}
		if len(prefix) > len(strs[i]) {
			prefix = prefix[:len(strs[i])]
		} else {
			strs[i] = strs[i][:len(prefix)]
		}
		for j := 0; j < len(prefix); j++ {
			if prefix[j] != strs[i][j] {
				prefix = prefix[:j]
			}
		}
	}
	return prefix
}
func longestCommonPrefix2(strs []string) string {
	sort.Strings(strs)
	last := len(strs) - 1
	prefix := ""
	for i := 0; i < min(len(strs[0]), len(strs[last])); i++ {
		if strs[0][i] != strs[last][i] {
			break
		}
		prefix += strs[0][i : i+1]
	}
	return prefix
}
