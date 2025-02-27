package main

import (
	"fmt"
	"slices"
	"sort"
)

//Given an array nums of n integers, are there elements a, b, c in nums such that a + b + c = 0? Find all unique
//triplets in the array which gives the sum of zero.

//Note:
//The solution set must not contain duplicate triplets.

//Example:
//	Given array nums = [-1, 0, 1, 2, -1, -4],

//	A solution set is:
//	[
//		[-1, 0, 1],
//		[-1, -1, 2]
//	]

func main() {
	nums := []int{1, 2, -2, -1}
	//fmt.Println(three(nums))
	a := threeSum2(nums)
	fmt.Println(a)
	//fmt.Println(violence(nums))
}

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	result, start, end, index, addNum, length := make([][]int, 0), 0, 0, 0, 0, len(nums)
	for index = 1; index < length-1; index++ {
		start, end = 0, length-1
		if index > 1 && nums[index] == nums[index-1] {
			start = index - 1
		}
		for start < index && end > index {
			if start > 0 && nums[start] == nums[start-1] {
				start++
				continue
			}
			if end < length-1 && nums[end] == nums[end+1] {
				end--
				continue
			}
			addNum = nums[start] + nums[end] + nums[index]
			if addNum == 0 {
				result = append(result, []int{nums[start], nums[index], nums[end]})
				start++
				end--
			} else if addNum > 0 {
				end--
			} else {
				start++
			}
		}
	}
	return result

}

// 暴力破解
func violence(nums []int) [][]int {
	res := make([][]int, 0)
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			for k := j + 1; k < len(nums); k++ {
				if nums[i]+nums[j]+nums[k] == 0 {
					res = append(res, []int{nums[i], nums[j], nums[k]})
				}
			}
		}

	}
	return res
}

func threeSum2(nums []int) [][]int {
	var res [][]int
	slices.Sort(nums)
	for l := 0; l < len(nums); l++ {
		// 前一个数和当前相等
		if l > 0 && nums[l] == nums[l-1] {
			continue
		}
		r := len(nums) - 1
		target := -nums[l]
		for c := l + 1; c < r; c++ {
			if c > l+1 && nums[c] == nums[c-1] {
				// 中间的不相等
				continue
			}
			for c < r && nums[c]+nums[r] > target {
				r--
			}

			if c == r {
				break
			}
			if nums[c]+nums[r] == target {
				res = append(res, []int{nums[l], nums[c], nums[r]})
			}
		}
	}

	return res

}
