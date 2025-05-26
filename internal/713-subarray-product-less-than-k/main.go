package _13_subarray_product_less_than_k

// https://leetcode.cn/problems/subarray-product-less-than-k/
func numSubarrayProductLessThanK(nums []int, k int) int {
	if k == 0 {
		return 0
	}
	l, cnt, ans := 0, 1, 0
	for r, v := range nums {
		cnt *= v
		for cnt >= k {
			cnt /= nums[l]
			l++
		}
		ans += r - l + 1
	}
	return ans
}
