package LCP_08_ju_qing_hong_fa_shi_jian

// https://leetcode.cn/problems/ju-qing-hong-fa-shi-jian/
func getTriggerTime(increase [][]int, requirements [][]int) []int {
	var ans []int
	var property [][3]int
	property = append(property, [3]int{0, 0, 0})
	for i := 0; i < len(increase); i++ {
		property = append(property, [3]int{increase[i][0] + property[i][0], increase[i][1] + property[i][1], increase[i][2] + property[i][2]})
	}

	lowerBound := func(property [][3]int, req []int) int {
		l, r := 0, len(property)-1
		for l <= r {
			if mid := (r + l) >> 1; property[mid][0] < req[0] || property[mid][1] < req[1] || property[mid][2] < req[2] {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}

		return l
	}
	for _, requirement := range requirements {
		i := lowerBound(property, requirement)
		if i >= len(property) {
			i = -1
		}
		ans = append(ans, i)
	}
	return ans
}
