package _62_find_peak_element

//https://leetcode.cn/problems/find-peak-element/
func findPeakElement(nums []int) int {
	l, r := 0, len(nums)
	for l+1 < r {
		mid := (l + r) >> 1
		pre, next := mid-1, mid+1
		if pre < 0 {
			pre = len(nums) - 1
		}
		if next >= len(nums) {
			next = 0
		}
		if nums[mid] > nums[pre] && nums[mid] > nums[next] {
		}
	}
	return l
}
