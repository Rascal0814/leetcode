package _841_maximum_sum_of_almost_unique_subarray

// https://leetcode.cn/problems/maximum-sum-of-almost-unique-subarray/
func maxSum(nums []int, m int, k int) int64 {
	var exitMap = make(map[int]int)
	var sum, res int
	for i := 0; i < len(nums); i++ {
		exitMap[nums[i]]++
		sum += nums[i]

		if i >= k {
			exitMap[nums[i-k]]--
			if exitMap[nums[i-k]] == 0 {
				delete(exitMap, nums[i-k])
			}
			sum -= nums[i-k]
		}

		if len(exitMap) >= m {
			res = max(res, sum)
		}
	}
	return int64(res)
}
