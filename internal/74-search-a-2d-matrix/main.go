package _4_search_a_2d_matrix

//https://leetcode.cn/problems/search-a-2d-matrix/
func searchMatrix(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	var arr []int
	for i := m - 1; i >= 0; i-- {
		if matrix[i][0] <= target {
			arr = matrix[i]
			break
		}
	}
	if arr == nil {
		return false
	}

	l, r := 0, n
	for l+1 < r {
		mid := (l + r) >> 1
		if arr[mid] <= target {
			l = mid
		} else {
			r = mid
		}
	}
	return arr[l] == target
}
