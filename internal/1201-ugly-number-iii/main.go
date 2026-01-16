package _201_ugly_number_iii

func nthUglyNumber(n int, a int, b int, c int) int {
	lcmAB := a * b / gcd(a, b)          // a,c最小公倍数
	lcmAC := a * c / gcd(a, c)          // a,c最小公倍数
	lcmBC := b * c / gcd(b, c)          // b,c最小公倍数
	lcmABC := lcmAB * c / gcd(lcmAB, c) // a,b,c最小公倍数

	l, r := 0, min(a, b, b)*n

	check := func(x int) bool {
		return x/a+x/b+x/c+x/lcmABC-x/lcmAB-x/lcmBC-x/lcmAC >= n
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

// gcd 最大公因数
func gcd(a, b int) int {
	if a%b != 0 {
		return gcd(b, a%b)
	}
	return b
}
