package LCP_68_1GxJYY

// https://leetcode.cn/problems/1GxJYY/
func beautifulBouquet(flowers []int, cnt int) int {
	l, ans, flowerMap := 0, 0, [100001]int{}
	for r, v := range flowers {
		flowerMap[v]++
		for flowerMap[v] > cnt {
			flowerMap[flowers[l]]--
			l++
		}
		ans += r - l + 1
	}
	return ans
}
