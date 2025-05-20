package _658_minimum_operations_to_reduce_x_to_zero

// https://leetcode.cn/problems/minimum-operations-to-reduce-x-to-zero/
func minOperations(nums []int, x int) int {
	var total int
	for _, num := range nums {
		total += num
	}
	target := total - x
	if target == 0 {
		return len(nums)
	}
	l, val, ans := 0, 0, -1
	for r, v := range nums {
		val += v
		for l < r && val > target {
			val -= nums[l]
			l++
		}
		if val == target {
			ans = max(ans, r-l+1)
		}
	}
	if ans == -1 {
		return ans
	}
	return len(nums) - ans
}
