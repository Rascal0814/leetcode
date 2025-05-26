package _248_count_number_of_nice_subarrays

// https://leetcode.cn/problems/count-number-of-nice-subarrays/
func numberOfSubarrays(nums []int, k int) int {
	var l1, l2, cnt1, cnt2, ans int
	for _, v := range nums {
		if v%2 != 0 {
			cnt1++
			cnt2++
		}

		for cnt1 >= k {
			if nums[l1]%2 != 0 {
				cnt1--
			}
			l1++
		}

		for cnt2 >= k+1 {
			if nums[l2]%2 != 0 {
				cnt2--
			}
			l2++
		}
		ans += l1 - l2
	}
	return ans
}
