package _090_k_radius_subarray_averages

// https://leetcode.cn/problems/k-radius-subarray-averages/
func getAverages(nums []int, k int) []int {
	var res []int
	if len(nums) < 2*k+1 {
		for range nums {
			res = append(res, -1)
		}
		return res
	}

	sum := 0
	// 计算半径为k的总和
	for i := 0; i < 2*k+1; i++ {
		sum += nums[i]
	}
	for i := 0; i < len(nums); i++ {
		if i < k || len(nums)-i-1 < k {
			res = append(res, -1)
			continue
		}
		if i > k {
			sum += nums[i+k] - nums[i-k-1]
		}
		res = append(res, sum/(2*k+1))
	}
	return res
}

func getAverages2(nums []int, k int) []int {
	var res = make([]int, len(nums))
	var sum int
	for i, num := range nums {
		sum += num
		if i < k || len(nums)-i-1 < k {
			res[i] = -1
		}
		if i < 2*k {
			continue
		}
		res[i-k] = sum / (2*k + 1)

		sum -= nums[i-2*k]

	}
	return res
}
