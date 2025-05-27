package _5_Search_Insert_Position

// https://leetcode.cn/problems/search-insert-position/
func searchInsert(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r { // 闭区间区间不为空 [left, right]
		if mid := (r + l) >> 1; nums[mid] < target {
			l = mid + 1 // [mid+1, right]
		} else {
			r = mid - 1 // [left, mid-1]
		}
	}
	return l
}
