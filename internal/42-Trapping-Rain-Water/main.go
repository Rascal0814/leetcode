package main

import (
	"fmt"
)

const q = `
42. 接雨水

给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。

示例 1：

	
	输入：height = [0,1,0,2,1,0,1,3,2,1,2,1]
	输出：6
	解释：上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）。 
	示例 2：
	
	输入：height = [4,2,0,3,2,5]
	输出：9
 

提示：

	n == height.length
	1 <= n <= 2 * 104
	0 <= height[i] <= 10^5
`

func trap(height []int) int {
	var sum int
	//计算左边高度的最大值
	var leftHeight = make([]int, len(height))
	leftHeight[0] = height[0]
	for i := 1; i < len(height); i++ {
		leftHeight[i] = max(leftHeight[i-1], height[i])
	}

	//计算右边高度的最大值
	var rightHeight = make([]int, len(height))
	rightHeight[len(height)-1] = height[len(height)-1]
	for i := len(height) - 2; i >= 0; i-- {
		rightHeight[i] = max(rightHeight[i+1], height[i])
	}

	//将两边最大值的 最小值减去当前高度的值就是储水量 这里可以去看看 leetcode 官网的题解图最直观
	for i := 0; i < len(height); i++ {
		sum += min(leftHeight[i], rightHeight[i]) - height[i]
	}
	return sum
}

func main() {
	fmt.Println(trap2([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
	fmt.Println(trap2([]int{4, 2, 0, 3, 2, 5}))
}

func trap2(height []int) int {
	var n = len(height) - 1
	var left = make([]int, len(height))
	left[0] = height[0]
	var right = make([]int, len(height))
	right[n] = height[n]
	for i := 1; i <= n; i++ {
		left[i] = max(left[i-1], height[i])
	}

	for i := n - 1; i >= 0; i-- {
		right[i] = max(right[i+1], height[i])
	}

	var total int
	for i := 0; i <= n; i++ {
		total += min(left[i], right[i]) - height[i]
	}
	return total
}
