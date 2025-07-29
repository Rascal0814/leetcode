package _75_h_index_ii

// https://leetcode.cn/problems/h-index-ii/
func hIndex(citations []int) int {
	l, r, total := 0, len(citations)-1, len(citations)
	for l <= r {

		if mid := (l + r) >> 1; citations[mid] < total-mid {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	if l >= len(citations) {
		return 0
	}
	return total - l
}
