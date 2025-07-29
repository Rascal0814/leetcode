package _982_find_longest_special_substring_that_occurs_thrice_ii

import "slices"

// https://leetcode.cn/problems/find-longest-special-substring-that-occurs-thrice-ii/description/
// 有哪些取出三个特殊子串的方法呢？
//  1. 从最长的特殊子串（a[0]）中取三个长度均为 a[0]−2 的特殊子串。例如示例 1 的 aaaa 可以取三个 aa。
//  2. 或者，从最长和次长的特殊子串（a[0],a[1]）中取三个长度一样的特殊子串：
//     - 如果 a[0]=a[1]，那么可以取三个长度均为 a[0]−1 的特殊子串。
//     - 如果 a[0]>a[1]，那么可以取三个长度均为 a[1] 的特殊子串：从最长中取两个，从次长中取一个。
//     - 这两种情况可以合并成 min(a[0]−1,a[1])，如果 a[0]−1<a[1]，这只能是第一种情况，因为 a[0]≥a[1]，我们取二者较小值 a[0]−1；如果 a[0]−1≥a[1]，即 a[0]>a[1]，这是第二种情况，我们也取的是二者较小值 a[1]。
//     3.又或者，从最长、次长、第三长的的特殊子串（a[0],a[1],a[2]）中各取一个长为 a[2] 的特殊子串。
//
// 这三种情况取最大值，即:		max(v[0]-2, min(v[0]-1, v[1]), v[2])
func maximumLength(s string) int {
	x := [26][]int{}

	var cnt int
	for i := 0; i < len(s); i++ {
		cnt++
		if i >= len(s)-1 || s[i] != s[i+1] {
			x[s[i]-'a'] = append(x[s[i]-'a'], cnt)
			cnt = 0
		}
	}

	var ans int
	for _, v := range x {
		if len(v) == 0 {
			continue
		}
		slices.SortFunc(v, func(a, b int) int { return b - a })
		v = append(v, 0, 0) // At least three elements are required to make judgments

		ans = max(ans, v[0]-2, min(v[0]-1, v[1]), v[2])

	}
	if ans == 0 {
		return -1
	}

	return ans
}
