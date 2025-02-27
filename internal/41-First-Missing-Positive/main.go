package main

import (
	"fmt"
)

const q = `
41. 缺失的第一个正数

给你一个未排序的整数数组 nums ，请你找出其中没有出现的最小的正整数。

请你实现时间复杂度为 O(n) 并且只使用常数级别额外空间的解决方案。
 

示例 1：

	输入：nums = [1,2,0]
	输出：3
	示例 2：

	输入：nums = [3,4,-1,1]
	输出：2
	示例 3：

	输入：nums = [7,8,9,11,12]
	输出：1
 

提示：

1 <= nums.length <= 5 * 105
-231 <= nums[i] <= 231 - 1
`

func firstMissingPositive(nums []int) int {

	n := len(nums)

	for i := 0; i < n; i++ {
		//nums[nums[i]-1] != nums[i] 两个如果重复的话就不需要变了
		if nums[i] > 0 && nums[i] < n && nums[nums[i]-1] != nums[i] {
			nums[nums[i]-1], nums[i] = nums[i], nums[nums[i]-1]
			// 如果换位了就需要判断换过来的值是否还需要换
			i--
		}
	}

	//如果位置上的值跟下标不想等就返回下标
	for i := 0; i < n; i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}

	// 如果都没找到的话就是全为比 数组长度大的值
	return n + 1
}

func main() {
	fmt.Println(firstMissingPositive([]int{1, 2, 0}))
	fmt.Println(firstMissingPositive([]int{3, 4, -1, 1}))
	fmt.Println(firstMissingPositive([]int{7, 8, 9, 11, 12}))
	fmt.Println(firstMissingPositive([]int{-1, 4, 2, 1, 9, 10}))
}
