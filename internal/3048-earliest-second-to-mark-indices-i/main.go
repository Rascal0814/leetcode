package _048_earliest_second_to_mark_indices_i

import "slices"

// https://leetcode.cn/problems/earliest-second-to-mark-indices-i/
func earliestSecondToMarkIndices(nums []int, changeIndices []int) int {
	l, r := len(nums), len(changeIndices)
	for _, v := range nums {
		l += v
	}
	lastT := make([]int, len(nums))

	check := func(x int) bool {
		clear(lastT)
		for i, v := range changeIndices[:x] {
			lastT[v-1] = i + 1
		}
		// changeIndices[:x] 缺少 nums 中某个值
		if slices.Contains(lastT, 0) {
			return false
		}

		cnt := 0 // 可用的里程数
		for i, v := range changeIndices[:x] {
			if lastT[v-1] == i+1 { // 已经到要标记的位置了
				if nums[v-1] > cnt {
					return false
				}
				cnt -= nums[v-1] // 减去nums[v-1]的里程数
			} else {
				cnt++
			}
		}

		return true
	}
	for l <= r {
		mid := (l + r) >> 1
		if check(mid) {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	if l > len(changeIndices) {
		return -1
	}
	return l
}
