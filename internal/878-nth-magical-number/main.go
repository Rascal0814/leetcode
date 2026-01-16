package _78_nth_magical_number

//https://leetcode.cn/problems/nth-magical-number/description/
func nthMagicalNumber(n int, a int, b int) int {
	mod := 1_0000_0000_7
	i := lcm(a, b)

	check := func(x int) bool {
		return x/a+x/b-x/i >= n
	}
	l, r := -1, min(a, b)*n
	for l+1 < r {
		if mid := (l + r) >> 1; check(mid) {
			r = mid
		} else {
			l = mid
		}
	}
	return r % mod
}

func gcd(a, b int) int {
	if a%b != 0 {
		return gcd(b, a%b)
	}
	return b
}

func lcm(a, b int) int {
	return a * b / gcd(a, b)
}
