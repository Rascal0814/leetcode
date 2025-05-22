package _09_minimum_size_subarray_sum

// https://leetcode.cn/problems/minimum-size-subarray-sum/description/
func minSubArrayLen(target int, nums []int) int {
	l, sum, ans := 0, 0, len(nums)+1
	for r, v := range nums {
		sum += v
		for sum-nums[l] >= target { // 竟可能的压缩子字符串的长度
			sum -= nums[l]
			l++
		}
		if sum >= target {
			ans = min(ans, r-l+1)
		}
	}
	if ans > len(nums) {
		return 0
	}
	return ans
}
