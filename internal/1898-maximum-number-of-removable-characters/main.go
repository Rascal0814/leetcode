package _898_maximum_number_of_removable_characters

// https://leetcode.cn/problems/maximum-number-of-removable-characters/
func maximumRemovals(s string, p string, removable []int) int {

	l, r := 0, len(removable)

	check := func(x int) bool {
		var tmp = make(map[int]struct{})
		for _, v := range removable[:x+1] {
			tmp[v] = struct{}{}
		}
		var i = len(p) - 1
		for ii, v := range s {
			if _, ok := tmp[ii]; ok {
				continue
			}
			if v == int32(p[len(p)-1-i]) {
				i--
			}
			if i < 0 {
				return true
			}
		}

		return false
	}

	//for l <= r {
	//	if mid := (l + r) >> 1; check(mid) {
	//		l = mid + 1
	//	} else {
	//		r = mid - 1
	//	}
	//}

	for l+1 < r {
		if mid := (l + r) >> 1; check(mid) {
			l = mid
		} else {
			r = mid
		}
	}

	return l + 1
}
