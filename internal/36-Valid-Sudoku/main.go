package main

import "fmt"

const q = `
请你判断一个9 x 9 的数独是否有效。只需要 根据以下规则 ，验证已经填入的数字是否有效即可。

数字1-9在每一行只能出现一次。
数字1-9在每一列只能出现一次。
数字1-9在每一个以粗实线分隔的3x3宫内只能出现一次。（请参考示例图）


注意：

一个有效的数独（部分已被填充）不一定是可解的。
只需要根据以上规则，验证已经填入的数字是否有效即可。
空白格用'.'表示。


示例 1：


输入：board = 
[["5","3",".",".","7",".",".",".","."]
,["6",".",".","1","9","5",".",".","."]
,[".","9","8",".",".",".",".","6","."]
,["8",".",".",".","6",".",".",".","3"]
,["4",".",".","8",".","3",".",".","1"]
,["7",".",".",".","2",".",".",".","6"]
,[".","6",".",".",".",".","2","8","."]
,[".",".",".","4","1","9",".",".","5"]
,[".",".",".",".","8",".",".","7","9"]]
输出：true
示例 2：

输入：board = 
[["8","3",".",".","7",".",".",".","."]
,["6",".",".","1","9","5",".",".","."]
,[".","9","8",".",".",".",".","6","."]
,["8",".",".",".","6",".",".",".","3"]
,["4",".",".","8",".","3",".",".","1"]
,["7",".",".",".","2",".",".",".","6"]
,[".","6",".",".",".",".","2","8","."]
,[".",".",".","4","1","9",".",".","5"]
,[".",".",".",".","8",".",".","7","9"]]
输出：false
解释：除了第一行的第一个数字从 5 改为 8 以外，空格内其他数字均与 示例1 相同。 但由于位于左上角的 3x3 宫内有两个 8 存在, 因此这个数独是无效的。


提示：

board.length == 9
board[i].length == 9
board[i][j] 是一位数字（1-9）或者 '.'

`

func isValidSudoku(board [][]byte) bool {
	var rowMap = make([][]bool, 9)
	var cellMap = make([][]bool, 9)
	var boxMap = make([][]bool, 9)
	//初始化
	for i := 0; i < 9; i++ {
		rowMap[i] = make([]bool, 9)
		cellMap[i] = make([]bool, 9)
		boxMap[i] = make([]bool, 9)
	}
	for row := 0; row < 9; row++ {
		for cell := 0; cell < 9; cell++ {
			if board[row][cell] != '.' {
				// 每行的差值应该都不相同才会每个值都不一样 所以计算与1的差值 这里一定要减 1 因为数组最大长度为9 但是下标最大值为8
				value := board[row][cell] - '1'
				if rowMap[row][value] || cellMap[cell][value] || boxMap[row/3*3+cell/3][value] {
					return false
				}
				rowMap[row][value] = true            // 行
				cellMap[cell][value] = true          // 列
				boxMap[row/3*3+cell/3][value] = true // 3*3 可以想像成一排三个
			}
		}
	}

	return true
}

func main() {
	board := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}

	board1 := [][]byte{
		{'.', '.', '4', '.', '.', '.', '6', '3', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'5', '.', '.', '.', '.', '.', '.', '9', '.'},
		{'.', '.', '.', '5', '6', '.', '.', '.', '.'},
		{'4', '.', '3', '.', '.', '.', '.', '.', '1'},
		{'.', '.', '.', '7', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '5', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
	}
	fmt.Println("value========>>>>>>>>", isValidSudoku(board))
	fmt.Println("value2========>>>>>>>>", isValidSudoku(board1))
}
