package _170_compare_strings_by_frequency_of_the_smallest_character

import (
	"slices"
)

// https://leetcode.cn/problems/compare-strings-by-frequency-of-the-smallest-character/
func numSmallerByFrequency(queries []string, words []string) []int {
	f := func(w string) int {
		cnt := [26]int{}
		for _, v := range w {
			cnt[v-'a']++
		}
		for _, v := range cnt {
			if v > 0 {
				return v
			}
		}
		return 0
	}

	var wordsMin []int
	for _, word := range words {
		wordsMin = append(wordsMin, f(word))
	}

	lowerBound := func(n []int, target int) int {
		l, r := 0, len(n)-1
		for l <= r {
			if mid := (r + l) >> 1; n[mid] < target {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
		return l
	}

	var ans []int
	slices.Sort(wordsMin)
	for _, query := range queries {
		l := lowerBound(wordsMin, f(query)+1)
		ans = append(ans, len(wordsMin)-l)
	}
	return ans
}
