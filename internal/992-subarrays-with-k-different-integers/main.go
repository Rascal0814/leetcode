package _92_subarrays_with_k_different_integers

// https://leetcode.cn/problems/subarrays-with-k-different-integers/
func subarraysWithKDistinct(nums []int, k int) int {
	l1, l2, cnt1, cnt2, ans := 0, 0, make(map[int]int), make(map[int]int), 0
	for _, v := range nums {
		cnt1[v]++
		cnt2[v]++

		for len(cnt1) >= k {
			if cnt1[nums[l1]]--; cnt1[nums[l1]] == 0 {
				delete(cnt1, nums[l1])
			}
			l1++
		}
		for len(cnt2) >= k+1 {
			if cnt2[nums[l2]]--; cnt2[nums[l2]] == 0 {
				delete(cnt2, nums[l2])
			}
			l2++
		}

		ans += l1 - l2
	}
	return ans
}
