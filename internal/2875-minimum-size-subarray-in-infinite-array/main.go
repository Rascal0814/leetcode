package _875_minimum_size_subarray_in_infinite_array

import "math"

// https://leetcode.cn/problems/minimum-size-subarray-in-infinite-array/
func minSizeSubarray(nums []int, target int) int {
	l, val, sum := 0, 0, 0
	for _, v := range nums {
		sum += v
	}
	nums = append(nums, nums...)
	residue := target % sum // 余数肯定在这一个数组中
	ans := math.MaxInt
	for r, v := range nums {
		val += v
		for val > residue {
			val -= nums[l]
			l++
		}
		if val == residue {
			ans = min(ans, r-l+1)
		}
	}

	if ans == math.MaxInt {
		return -1
	}

	return ans + (target/sum)*len(nums)/2
}
