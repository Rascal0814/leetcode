package main

import "sort"

//-1 -1 0 0 0 1 2
func three(nums []int) [][]int {
	sort.Ints(nums)
	res, current, leftPoint, rightPoint := make([][]int, 0, 0), 0, 0, 0
	for current = 1; current < len(nums)-1; current++ {
		leftPoint, rightPoint = 0, len(nums)-1
		if current > 1 && nums[current] == nums[current-1] {
			leftPoint = current - 1
		}
		for leftPoint < current && rightPoint > current {
			if leftPoint > 0 && nums[leftPoint] == nums[leftPoint-1] {
				leftPoint++
				continue
			}
			if rightPoint < len(nums)-1 && nums[rightPoint] == nums[rightPoint+1] {
				rightPoint--
				continue
			}
			temp := nums[leftPoint] + nums[current] + nums[rightPoint]
			if temp > 0 {
				rightPoint--
			} else if temp < 0 {
				leftPoint++
			} else if temp == 0 {
				res = append(res, []int{nums[leftPoint], nums[current], nums[rightPoint]})
				leftPoint++
				rightPoint--
			}
		}
	}
	return res
}
