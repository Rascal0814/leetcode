package _802_maximum_value_at_a_given_index_in_a_bounded_array

// https://leetcode.cn/problems/maximum-value-at-a-given-index-in-a-bounded-array/
func maxValue(n int, index int, maxSum int) int {
	l, r := 1, maxSum

	sum := func(a1 int, n int) int {
		if n == 0 {
			return 0
		}
		if n <= a1 { // 形成等差数列 x,x-1,x-2,x-3...1 求和公式：sn = (2*a1 + (n-1)*d) * n / 2
			return (2*a1 + 1 - n) * n / 2
		}
		// x,x-1...1,1,1,1
		return (a1+1)*a1/2 + (n - a1)
	}

	for l < r {
		mid := (l + r) >> 1
		if sum(mid, index)+sum(mid, n-index-1)+mid >= maxSum {
			r = mid
		} else {
			l = mid + 1
		}
	}

	return l

}
