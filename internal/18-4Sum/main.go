package main

import (
	"fmt"
	"sort"
)

const q = `Given an array nums of n integers and an integer target, are there elements a, b, c, and d in nums such that a + b + c + d = target? Find all unique quadruplets in the array which gives the sum of target.

Note:

The solution set must not contain duplicate quadruplets.

Example:


Given array nums = [1, 0, -1, 0, -2, 2], and target = 0.

A solution set is:
[
  [-1,  0, 0, 1],
  [-2, -1, 1, 2],
  [-2,  0, 0, 2]
]
[-2,-1,-1,1,1,2,2]
`

func main() {
	//fmt.Println(violence([]int{2, 2, 2, 2, 2}, 8))
	//fmt.Println(one([]int{2, 2, 2, 2, 2}, 8))
	fmt.Println(one([]int{-2, -1, -1, 1, 1, 2, 2}, 0))
}

// 暴力破解 O(n^4)
func violence(nums []int, target int) [][]int {
	res := make([][]int, 0)
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			for k := j + 1; k < len(nums); k++ {
				for f := k + 1; f < len(nums); f++ {
					if nums[i]+nums[j]+nums[k]+nums[f] == target {
						res = append(res, []int{nums[i], nums[j], nums[k], nums[f]})
					}
				}
			}
		}
	}
	return res
}

// 思路来源三数之和 排序➕双指针
func one(nums []int, target int) [][]int {
	sort.Ints(nums)
	l, c, r, res := 0, 0, 0, make([][]int, 0)
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		// 头尾指针
		for c = i + 2; c < len(nums); c++ {
			l, r = i+1, len(nums)-1
			if c > i+2 && nums[c] == nums[c-1] {
				l = c - 1
			}
			for l < c && r > c {
				if l > i+1 && nums[l] == nums[l-1] {
					l++
					continue
				}
				if r < len(nums)-1 && r-1 > c && nums[r] == nums[r-1] {
					r--
					continue
				}

				temp := nums[i] + nums[l] + nums[c] + nums[r]

				if temp == target {
					res = append(res, []int{nums[i], nums[l], nums[c], nums[r]})
					l++
					r--
				} else if temp > target {
					r--
				} else {
					l++
				}
			}
		}
	}
	return res
}

func fourSum(nums []int, target int) [][]int {

	n := len(nums)
	if n < 4 {
		return [][]int{}
	}

	sort.Ints(nums)

	ans := make([][]int, 0)
	for i := 0; i < n-3 && nums[i]+nums[i+1]+nums[i+2]+nums[i+3] <= target; i++ {

		if i > 0 && nums[i-1] == nums[i] || nums[i]+nums[n-3]+nums[n-2]+nums[n-1] < target {
			continue
		}
		for j := i + 1; j < n-2 && nums[i]+nums[j]+nums[j+1]+nums[j+2] <= target; j++ {

			if j > i+1 && nums[j-1] == nums[j] || nums[i]+nums[j]+nums[n-2]+nums[n-1] < target {
				continue
			}

			for left, right := j+1, n-1; left < right; {
				if sum := nums[i] + nums[j] + nums[left] + nums[right]; sum == target {
					ans = append(ans, []int{nums[i], nums[j], nums[left], nums[right]})

					for left++; left < right && nums[left] == nums[left-1]; left++ {
					}
					for right--; left < right && nums[right] == nums[right+1]; right-- {
					}
				} else if sum > target {
					right--
				} else {
					left++
				}
			}

		}
	}

	return ans

}
