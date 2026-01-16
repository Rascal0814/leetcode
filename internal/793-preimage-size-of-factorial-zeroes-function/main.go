package _93_preimage_size_of_factorial_zeroes_function

//https://leetcode.cn/problems/preimage-size-of-factorial-zeroes-function/
func preimageSizeFZF(k int) int {
	fzf := func(k int) int {
		l, r := -1, 5*k+1
		check := func(x int) bool {
			cnt := 0
			for x > 0 {
				x /= 5
				cnt += x
			}
			return cnt >= k
		}
		for l+1 < r {
			if mid := (l + r) >> 1; check(mid) {
				r = mid
			} else {
				l = mid
			}
		}
		return r
	}
	return fzf(k+1) - fzf(k)
}
