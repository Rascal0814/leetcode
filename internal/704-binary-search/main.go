package _04_binary_search

// https://leetcode.cn/problems/binary-search/
func search(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		if mid := (r + l) >> 1; nums[mid] < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	if l >= len(nums) || nums[l] != target {
		return -1
	}
	return l
}
