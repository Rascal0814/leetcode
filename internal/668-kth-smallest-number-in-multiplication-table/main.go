package _68_kth_smallest_number_in_multiplication_table

//https://leetcode.cn/problems/kth-smallest-number-in-multiplication-table/
func findKthNumber(m int, n int, k int) int {
	check := func(x int) bool {
		cnt := 0
		for i := range m {
			cnt += min(x/(i+1), n)
		}
		return cnt < k
	}

	l, r := 0, m*n
	for l+1 < r {
		if mid := (l + r) >> 1; check(mid) {
			l = mid
		} else {
			r = mid
		}
	}
	return r

}
