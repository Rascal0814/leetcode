package _43_maximum_average_subarray_i

// https://leetcode.cn/problems/maximum-average-subarray-i/
func findMaxAverage(nums []int, k int) float64 {
	var total int
	for i := 0; i < k; i++ {
		total += nums[i]
	}
	sum := total
	for i := k; i < len(nums); i++ {
		sum += nums[i] - nums[i-k]
		if sum > total {
			total = sum
		}
	}

	return float64(total) / float64(k)
}
