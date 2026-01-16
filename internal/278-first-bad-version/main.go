package _78_first_bad_version

func firstBadVersion(n int) int {
	l, r := 0, n
	for l+1 < r {
		mid := (l + r) >> 1
		if isBadVersion(mid) {
			r = mid
		} else {
			l = mid
		}
	}
	return r
}
