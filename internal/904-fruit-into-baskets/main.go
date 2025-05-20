package _04_fruit_into_baskets

// https://leetcode.cn/problems/fruit-into-baskets/
func totalFruit(fruits []int) int {
	l, n, cnt, mx := 0, len(fruits), make(map[int]int), 0

	for r := 0; r < n; r++ {
		cnt[fruits[r]]++ // 水果类型计数
		for len(cnt) > 2 {
			cnt[fruits[l]]--
			if cnt[fruits[l]] == 0 {
				delete(cnt, fruits[l])
			}
			l++
		}
		mx = max(mx, r-l+1)
	}
	return mx
}
