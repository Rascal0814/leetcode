package main

import (
	"fmt"
	"sort"
)

func two(nums []int) int {
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			// 将后面的元素向前移动一位
			copy(nums[i-1:], nums[i:])
			fmt.Println(nums)

			// 将数组长度减一
			nums = nums[:len(nums)-1]
			i--
		}
	}
	return len(nums)
}
