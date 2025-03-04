package _461_maximum_sum_of_distinct_subarrays_with_length_k

// https://leetcode.cn/problems/maximum-sum-of-distinct-subarrays-with-length-k/
func maximumSubarraySum(nums []int, k int) int64 {
	var exitMap = make(map[int]int)
	var sum, res int
	for i := 0; i < len(nums); i++ {
		exitMap[nums[i]]++
		sum += nums[i]

		if i >= k {
			exitMap[nums[i-k]]--
			if v, ok := exitMap[nums[i-k]]; v <= 0 || !ok {
				delete(exitMap, nums[i-k])
			}
			sum -= nums[i-k]
		}

		if len(exitMap) >= k {
			res = max(res, sum)
		}

	}
	return int64(res)
}
