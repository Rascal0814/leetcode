package _208_get_equal_substrings_within_budget

// https://leetcode.cn/problems/get-equal-substrings-within-budget/
func equalSubstring(s string, t string, maxCost int) int {
	l, n, dif, cost, mx := 0, len(s), make([]int, len(s)), 0, 0
	for i := 0; i < n; i++ {
		if s[i] > t[i] {
			dif[i] = int(s[i] - t[i])
		} else {
			dif[i] = int(t[i] - s[i])

		}
	}
	for r := 0; r < n; r++ {
		cost += dif[r]
		for cost > maxCost {
			cost -= dif[l]
			l++
		}

		mx = max(mx, r-l+1)
	}
	return mx
}
