package main

import "fmt"

const q = `
给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。

示例 1：

	输入：nums = [1,2,3]
	输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
示例 2：

	输入：nums = [0,1]
	输出：[[0,1],[1,0]]
示例 3：

	输入：nums = [1]
	输出：[[1]]
 

提示：

	1 <= nums.length <= 6
	-10 <= nums[i] <= 10
	nums 中的所有整数 互不相同
`

func permute(nums []int) [][]int {
	var res = make([][]int, 0)

	var dfs func(nums []int, index int)
	dfs = func(nums []int, index int) {
		comb := make([]int, len(nums))
		for i := index; i < len(nums)-1; i++ {
			if index == len(nums) {
				temp
			}
			res = append(res, append(comb, nums[i]))
			dfs(nums, index+1)
		}
	}

	dfs(nums, 0)

	return res
}

func main() {
	fmt.Println(permute([]int{1, 2, 3}))
}
