package main

import "sort"

//[-1, 0, 1, 2, -1, -4]
// [-4,-1,-1,0,1,2]
//		[-1, 0, 1],
//		[-1, -1, 2]
func two(nums []int) [][]int {
	sort.Ints(nums)

	res, sum, i, l, r := make([][]int, 0, 0), 0, 0, 0, 0
	for i = 1; i < len(nums)-1; i++ {
		l, r = 0, len(nums)-1
		if i > 1 && nums[i] == nums[i-1] {
			l = i - 1
		}
		for l < i && r > i {

			if l > 0 && nums[l] == nums[l-1] {
				l++
				continue
			}
			if r < len(nums)-1 && nums[r] == nums[r+1] {
				r--
				continue
			}
			if sum = nums[l] + nums[r] + nums[i]; sum == 0 {
				res = append(res, []int{nums[l], nums[r], nums[i]})
				l++
				r--
			} else if sum > 0 {
				r--
			} else {
				l++
			}
		}
	}

	return res
}
