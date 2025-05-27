package _4_find_first_and_last_position_of_element_sorted_array

func searchRange(nums []int, target int) []int {
	lowerBound := func(n []int, t int) int {
		l, r := 0, len(n)-1
		for l <= r {
			if mid := (r + l) >> 1; n[mid] >= t {
				r = mid - 1
			} else {
				l = mid + 1
			}
		}
		return l
	}
	start := lowerBound(nums, target)
	if start == len(nums) || nums[start] != target {
		return []int{-1, -1}
	}
	return []int{start, lowerBound(nums, target+1) - 1}
}
