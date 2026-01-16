package _74_guess_number_higher_or_lower

//https://leetcode.cn/problems/guess-number-higher-or-lower/
func guessNumber(n int) int {
	l, r := 0, n
	for l+1 < r {
		mid := (l + r) >> 1
		if guess(mid) <= 0 {
			r = mid
		} else {
			l = mid
		}
	}
	return r
}
