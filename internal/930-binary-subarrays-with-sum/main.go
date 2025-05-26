package _30_binary_subarrays_with_sum

func numSubarraysWithSum(nums []int, goal int) int {
	var l1, l2, cnt1, cnt2, ans int
	for r, v := range nums {
		cnt1 += v
		cnt2 += v
		for l1 <= r && cnt1 >= goal {
			cnt1 -= nums[l1]
			l1++
		}

		for l2 <= r && cnt2 >= goal+1 {
			cnt2 -= nums[l2]
			l2++
		}
		ans += l1 - l2
	}
	return ans
}
