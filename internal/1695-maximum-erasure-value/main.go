package _695_maximum_erasure_value

// https://leetcode.cn/problems/maximum-erasure-value/
func maximumUniqueSubarray(nums []int) int {
	l, val, ans, cnt := 0, 0, 0, make(map[int]bool)

	for _, v := range nums {
		for cnt[v] {
			val -= nums[l]
			cnt[nums[l]] = false
			l++
		}

		val += v
		cnt[v] = true
		ans = max(ans, val)
	}
	return ans
}
