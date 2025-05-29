package _070_most_beautiful_item_for_each_query

import (
	"cmp"
	"slices"
)

func maximumBeauty(items [][]int, queries []int) []int {
	slices.SortFunc(items, func(a, b []int) int {
		return cmp.Compare(a[0], b[0])
	})

	// 前缀合最大值
	for i := 0; i < len(items); i++ {
		if i-1 >= 0 {
			items[i][1] = max(items[i][1], items[i-1][1])
		}
	}
	lowerBound := func(n [][]int, target int) int {
		l, r := 0, len(n)-1
		for l <= r {
			if mid := (r + l) >> 1; n[mid][0] < target {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
		return l
	}

	var ans []int
	for _, v := range queries {
		bound := lowerBound(items, v+1)
		if bound == 0 && items[bound][0] != v {
			ans = append(ans, 0)
			continue
		}
		ans = append(ans, items[bound-1][1])
	}

	return ans
}
