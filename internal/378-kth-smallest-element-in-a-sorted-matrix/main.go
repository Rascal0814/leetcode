package _78_kth_smallest_element_in_a_sorted_matrix

//https://leetcode.cn/problems/kth-smallest-element-in-a-sorted-matrix/
func kthSmallest(matrix [][]int, k int) int {
	n := len(matrix) - 1
	l, r := matrix[0][0], matrix[n][n]
	check := func(x int) bool {
		cnt, i, j := 0, n, 0
		for i >= 0 && j <= n {
			if matrix[i][j] <= x {
				cnt += i + 1
				j++
			} else {
				i--
			}
		}
		return cnt >= k
	}
	for l < r {
		if mid := (l + r) >> 1; check(mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return r
}
