package main

import "sort"

func four(nums []int) [][]int {
	sort.Ints(nums)
	l, c, r, res := 0, 0, 0, make([][]int, 0)
	for c = 1; c < len(nums)-1; c++ {
		l, r = 0, len(nums)-1
		if c > 1 && nums[c] == nums[c-1] {
			l = c - 1
		}

		for l < c && r > c {
			if l > 0 && nums[l] == nums[l-1] {
				l++
				continue
			}
			if r < len(nums)-1 && nums[r] == nums[r+1] {
				r--
				continue
			}
			sum := nums[l] + nums[c] + nums[r]
			if sum == 0 {
				res = append(res, []int{nums[l], nums[c], nums[r]})
				r--
				l++
			} else if sum > 0 {
				r--
			} else {
				l++
			}
		}
	}
	return res
}
