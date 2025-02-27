package main

import "fmt"

const q = `
给你一个按照非递减顺序排列的整数数组 nums，和一个目标值 target。请你找出给定目标值在数组中的开始位置和结束位置。

如果数组中不存在目标值 target，返回[-1, -1]。

你必须设计并实现时间复杂度为O(log n)的算法解决此问题。

示例 1：

输入：nums = [5,7,7,8,8,10], target = 8
输出：[3,4]
示例2：

输入：nums = [5,7,7,8,8,10], target = 6
输出：[-1,-1]
示例 3：

输入：nums = [], target = 0
输出：[-1,-1]

`

func main() {
	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 8))
	fmt.Println(end([]int{8, 8, 8, 8, 8, 8}, 8))
}

// 二分查找
func searchRange(nums []int, target int) []int {
	return []int{searchFirstEqualElement(nums, target), searchLastEqualElement(nums, target)}
}

// 二分查找第一个与 target 相等的元素，时间复杂度 O(logn)
func searchFirstEqualElement(nums []int, target int) int {
	low, high := 0, len(nums)-1
	for low <= high {
		mid := low + ((high - low) >> 1)
		if nums[mid] > target {
			high = mid - 1
		} else if nums[mid] < target {
			low = mid + 1
		} else {
			// 相等的时候找第一个
			if (mid == 0) || (nums[mid-1] != target) {
				return mid
			}
			high = mid - 1
		}
	}
	return -1
}

func searchLastEqualElement(nums []int, target int) int {
	low, high := 0, len(nums)-1
	for low <= high {
		mid := low + ((high - low) >> 1)
		if nums[mid] > target {
			high = mid - 1
		} else if nums[mid] < target {
			low = mid + 1
		} else {
			// 相等的时候找最后一个
			if (mid == len(nums)-1) || (nums[mid+1] != target) {
				return mid
			}
			low = mid + 1
		}
	}
	return -1
}

func end(nums []int, target int) int {
	start, end := 0, len(nums)-1
	for start <= end {
		mid := start + ((end - start) >> 1)
		if nums[mid] > target {
			end = mid - 1
			continue
		} else if nums[mid] < target {
			start = mid + 1
			continue
		} else {
			if mid == len(nums)-1 || nums[mid+1] != target {
				return mid
			}
			start = mid + 1
		}
	}
	return -1
}
