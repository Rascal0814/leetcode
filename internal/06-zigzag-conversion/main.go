package main

import (
	"strings"
)

const q = `
将一个给定字符串 s 根据给定的行数 numRows ，以从上往下、从左到右进行 Z 字形排列。

比如输入字符串为 "PAYPALISHIRING" 行数为 3 时，排列如下：

P   A   H   N
A P L S I I G
Y   I   R
之后，你的输出需要从左往右逐行读取，产生出一个新的字符串，比如："PAHNAPLSIIGYIR"。

请你实现这个将字符串进行指定行数变换的函数：

string convert(string s, int numRows);
 

示例 1：
	输入：s = "PAYPALISHIRING", numRows = 3
	输出："PAHNAPLSIIGYIR"
	示例 2：
	输入：s = "PAYPALISHIRING", numRows = 4
	输出："PINALSIGYAHRPI"
	解释：
	P     I    N
	A   L S  I G
	Y A   H R
	P     I
示例 3：
	
	输入：s = "A", numRows = 1
	输出："A"
	 

提示：

1 <= s.length <= 1000
s 由英文字母（小写和大写）、',' 和 '.' 组成
1 <= numRows <= 1000
`

// 01210
func convert(s string, numRows int) string {
	if numRows < 2 {
		return s
	}
	var temp = make([][]string, numRows)
	i, flag := 0, -1
	for _, v := range s {
		temp[i] = append(temp[i], string(v))
		// 两边需要转向
		if i == 0 || i == numRows-1 {
			flag = -flag
		}
		i += flag
	}
	res := ""
	for _, re := range temp {
		res += strings.Join(re, "")
	}

	return res
}

func convert2(s string, numRows int) string {
	if numRows < 2 {
		return s
	}
	var lineList = make([][]string, numRows)
	var flag bool // 加还是减
	var lineNum = 0
	for i := 0; i < len(s); i++ {
		lineList[lineNum] = append(lineList[lineNum], s[i:i+1])
		if !flag {
			lineNum++
		} else {
			lineNum--
		}
		if lineNum == 0 || lineNum == numRows-1 {
			flag = !flag
		}

	}
	var ans string
	for _, v := range lineList {
		ans += strings.Join(v, "")
	}

	return ans
}
