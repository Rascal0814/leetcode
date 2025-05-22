package _32_smallest_range_covering_elements_from_k_lists

import (
	"cmp"
	"math"
	"slices"
)

// https://leetcode.cn/problems/smallest-range-covering-elements-from-k-lists/
func smallestRange(nums [][]int) []int {
	bigNums := make([][2]int, 0)
	for index, num := range nums {
		for _, v := range num {
			bigNums = append(bigNums, [2]int{v, index})
		}
	}
	slices.SortFunc(bigNums, func(a, b [2]int) int {
		return cmp.Compare(a[0], b[0])
	})

	l, cnt, tCount, ansL, ansR := 0, make(map[int]int), 0, math.MaxInt, math.MaxInt
	for r, v := range bigNums {
		if cnt[v[1]] == 0 {
			tCount++
		}
		cnt[v[1]]++
		for tCount == len(nums) {
			if ansL == math.MaxInt || ansR == math.MaxInt || ansR-ansL > bigNums[r][0]-bigNums[l][0] {
				ansL, ansR = bigNums[l][0], bigNums[r][0]
			}
			if cnt[bigNums[l][1]]--; cnt[bigNums[l][1]] == 0 {
				tCount--
			}
			l++
		}

	}
	return []int{ansL, ansR}

}
