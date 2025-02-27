package main

import (
	"math"
	"sort"
)

//  Given an array nums of n integers and an integer target, find three integers in nums such that the sum is closest
//  to target. Return the sum of the three integers. You may assume that each input would have exactly one solution.
//
//  Example:
//  	Given array nums = [-1, 2, 1, -4], and target = 1.
//
//  	The sum that is closest to the target is 2. (-1 + 2 + 1 = 2).

func main() {
	//println(threeSumClosest_real([]int{-1, -1, 2, 2, 1, -4}, 1))
	println(threeSumClosest([]int{0, 1, 2}, 3))
}

// 暴力破解
func violence(nums []int, target int) int {
	res, min := 0, math.MaxInt
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			for k := j + 1; k < len(nums); k++ {
				if abs(nums[i]+nums[j]+nums[k]-target) < min {
					min = abs(nums[i] + nums[j] + nums[k] - target)
					res = nums[i] + nums[j] + nums[k]
				}
			}
		}

	}
	return res
}

// 排序+双指针
func threeSumClosest(nums []int, target int) int {
	// left前指针 current中间数 right后指针 res三数之和
	sort.Ints(nums)
	left, current, right, res := 0, 0, 0, math.MaxInt

	// 根据current循环
	for current = 1; current < len(nums); current++ {
		left, right = 0, len(nums)-1
		//如果左边的等于current的话就进位 不然就需要多判断几次
		if current > 1 && nums[current] == nums[current-1] {
			left = current - 1
		}

		for left < current && right > current {
			sum := nums[left] + nums[current] + nums[right]
			diff := target - sum
			if abs(diff) < abs(target-res) {
				res = sum
			}
			if sum > target {
				right--
			} else if sum < target {
				left++
			} else {
				return sum
			}
		}

	}
	return res
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
