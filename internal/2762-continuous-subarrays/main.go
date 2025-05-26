package _762_continuous_subarrays

// https://leetcode.cn/problems/continuous-subarrays/
func continuousSubarrays(nums []int) int64 {
	l, cnt, ans := 0, make(map[int]int), 0
	for r, v := range nums {
		cnt[v]++
		for {
			mx, mn := v, v
			for k := range cnt {
				mx = max(mx, k)
				mn = min(mn, k)
			}

			if mx-mn <= 2 {
				break
			}

			if cnt[nums[l]]--; cnt[nums[l]] == 0 {
				delete(cnt, nums[l])
			}
			l++
		}
		ans += r - l + 1
	}
	return int64(ans)
}
