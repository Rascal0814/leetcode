package main

import (
	"fmt"
	"strings"
)

// 回溯算法
const q = `
给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。答案可以按 任意顺序 返回。

给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。

示例 1：

输入：digits = "23"
输出：["ad","ae","af","bd","be","bf","cd","ce","cf"]
示例 2：

输入：digits = ""
输出：[]
示例 3：

输入：digits = "2"
输出：["a","b","c"]
 
`

func main() {
	letterCombinations("23")
}

func letterCombinations(digits string) []string {
	n := len(digits)
	if n == 0 {
		return []string{}
	}
	var phoneMap = []string{"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}
	ans := make([]string, 0, n)
	dfs := func(i int) string {
		if i == n {
			return strings.Join(ans, "")
		}
		for _, v := range phoneMap[i] {

		}
	}

	for i, v := range digits {
		fmt.Printf("i=%d;v=%d\n", i, v)
	}
	fmt.Println('1')
	return nil
}
