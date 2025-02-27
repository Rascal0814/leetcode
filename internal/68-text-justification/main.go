package _8_text_justification

import (
	"strings"
)

const _ = `
68.文本左右对齐

给定一个单词数组 words 和一个长度 maxWidth ，重新排版单词，使其成为每行恰好有 maxWidth 个字符，且左右两端对齐的文本。

你应该使用 “贪心算法” 来放置给定的单词；也就是说，尽可能多地往每行中放置单词。必要时可用空格 ' ' 填充，使得每行恰好有 maxWidth 个字符。

要求尽可能均匀分配单词间的空格数量。如果某一行单词间的空格不能均匀分配，则左侧放置的空格数要多于右侧的空格数。

文本的最后一行应为左对齐，且单词之间不插入额外的空格。

注意:

单词是指由非空格字符组成的字符序列。
每个单词的长度大于 0，小于等于 maxWidth。
输入单词数组 words 至少包含一个单词。
`

func fullJustify(words []string, maxWidth int) []string {
	var ans []string
	var wordsLen, startIndex int

	getSpace := func(num int) string {
		return strings.Repeat(" ", num)
	}
	for i := startIndex; startIndex < len(words) && i < len(words); {
		diff := wordsLen + len(words[i]) + len(words[startIndex:i]) - 1
		// 最后一行
		if i == len(words)-1 && diff < maxWidth {
			lineStr := strings.Join(words[startIndex:], " ")
			ans = append(ans, lineStr+getSpace(maxWidth-len(lineStr)))
			break
		}

		// 没满
		if diff < maxWidth {
			wordsLen += len(words[i])
			i++
			continue
		}

		// 已满
		if diff >= maxWidth {
			spaceNum := (maxWidth - wordsLen) / max(i-startIndex-1, 1)                           // 平均空格数
			extSpaceNum := (maxWidth - wordsLen) % max(i-startIndex-1, 1)                        // 额外空格数
			s1 := strings.Join(words[startIndex:startIndex+extSpaceNum+1], getSpace(spaceNum+1)) //前部分空格
			s2 := strings.Join(words[startIndex+extSpaceNum+1:i], getSpace(spaceNum))            //后部分空格
			ans = append(ans, s1+getSpace(spaceNum)+s2)
			wordsLen = 0
			startIndex = i
		}
	}
	return ans
}
