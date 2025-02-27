package _69_majority_element

import "slices"

const _ = `
169. 多数元素

给定一个大小为 n 的数组 nums ，返回其中的多数元素。多数元素是指在数组中出现次数 大于 ⌊ n/2 ⌋ 的元素。

你可以假设数组是非空的，并且给定的数组总是存在多数元素。`

func majorityElement(nums []int) int {
	var exitCount = make(map[int]int)
	for _, v := range nums {
		exitCount[v]++
		c := exitCount[v]
		if c > len(nums)/2 {
			return v
		}
	}
	return 0
}

func majorityElement2(nums []int) int {
	slices.Sort(nums)
	return nums[len(nums)/2]
}
