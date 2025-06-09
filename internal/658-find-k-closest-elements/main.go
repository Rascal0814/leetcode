package _58_find_k_closest_elements

import (
	"sort"
)

// https://leetcode.cn/problems/find-k-closest-elements/description/
func findClosestElements(arr []int, k int, x int) []int {
	right := sort.SearchInts(arr, x) // 找到x
	left := right - 1
	for ; k > 0; k-- {
		if left < 0 {
			right++
		} else if right >= len(arr) || x-arr[left] <= arr[right]-x { // arr[:left]肯定比x小 arr[right:]大于等于x
			left--
		} else {
			right++
		}
	}

	return arr[left+1 : right]
}
