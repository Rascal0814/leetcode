package _007_maximum_number_that_sum_of_the_prices_is_less_than_or_equal_to_k

import "fmt"

//https://leetcode.cn/problems/maximum-number-that-sum-of-the-prices-is-less-than-or-equal-to-k/
func findMaximumNumber(k int64, x int) int64 {
	check := func(n int64) bool {
		var c int64
		for i := x - 1; n>>i > 0; i += x {
			nn := n >> i
			fmt.Printf("%b\n", nn)
			c += nn >> 1 << i // 有一半的奇数，经过几轮i就相当于计算多少遍

			if nn&1 > 0 { // 当为奇数时需补充计算
				c += n&(1<<i-1) + 1
			}
		}
		return c <= k
	}
	var l, r = int64(0), (k + 1) << x
	for l+1 < r {
		mid := (l + r) >> 1
		if check(mid) {
			l = mid
		} else {
			r = mid
		}
	}

	return l
}
