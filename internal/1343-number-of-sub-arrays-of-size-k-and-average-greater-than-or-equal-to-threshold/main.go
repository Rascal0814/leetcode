package _343_number_of_sub_arrays_of_size_k_and_average_greater_than_or_equal_to_threshold

// https://leetcode.cn/problems/number-of-sub-arrays-of-size-k-and-average-greater-than-or-equal-to-threshold/
func numOfSubarrays(arr []int, k int, threshold int) int {
	var sum, res int
	for i := 0; i < len(arr); i++ {
		sum += arr[i]

		if i < k-1 {
			continue
		}
		if sum >= threshold*k {
			res++
		}
		sum -= arr[i-k+1]
	}

	return res
}
