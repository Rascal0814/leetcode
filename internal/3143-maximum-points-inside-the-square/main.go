package _143_maximum_points_inside_the_square

//https://leetcode.cn/problems/maximum-points-inside-the-square/
func maxPointsInsideSquare(points [][]int, s string) int {
	l, r := -1, 0

	for i := 0; i < len(points); i++ {
		if points[i][0] < 0 {
			points[i][0] *= -1
		}
		if points[i][1] < 0 {
			points[i][1] *= -1
		}
		r = max(points[i][0], points[i][1], r)
	}

	r += 1
	check := func(x int) int {
		var exit [26]int
		nums := 0

		for i, point := range points {
			if point[0] <= x && point[1] <= x {
				nums++
				if exit[s[i]-'a']++; exit[s[i]-'a'] > 1 {
					return -1
				}
			}
		}
		return nums
	}
	var res int
	for l+1 < r {
		mid := (l + r) >> 1
		tmp := check(mid)
		res = max(res, tmp)
		if tmp != -1 {
			l = mid
		} else {
			r = mid
		}
	}
	return res

}
