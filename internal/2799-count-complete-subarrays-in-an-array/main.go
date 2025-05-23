package _799_count_complete_subarrays_in_an_array

// https://leetcode.cn/problems/count-complete-subarrays-in-an-array/
func countCompleteSubarrays(nums []int) int {

	l, cnt, ans, allType := 0, make(map[int]int), 0, make(map[int]struct{})

	for _, v := range nums {
		allType[v] = struct{}{}
	}

	for _, v := range nums {
		cnt[v]++
		for len(cnt) == len(allType) {
			if cnt[nums[l]]--; cnt[nums[l]] == 0 {
				delete(cnt, nums[l])
			}
			l++
		}
		ans += l

	}
	return ans
}
