package _89_rotate_array

import (
	"slices"
)

const _ = `
189. 轮转数组

给定一个整数数组 nums，将数组中的元素向右轮转 k 个位置，其中 k 是非负数。
`

func rotate(nums []int, k int) {
	k %= len(nums) // k 可能小于 nums长度
	slices.Reverse(nums)
	slices.Reverse(nums[:k])
	slices.Reverse(nums[k:])
}

func rotate2(nums []int, k int) {
	var newNums = make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		newNums[(i+k)%len(nums)] = nums[i]
	}
	copy(nums, newNums)
}
