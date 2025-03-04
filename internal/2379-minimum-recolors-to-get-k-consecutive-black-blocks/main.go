package _379_minimum_recolors_to_get_k_consecutive_black_blocks

import "math"

// https://leetcode.cn/problems/minimum-recolors-to-get-k-consecutive-black-blocks/description/
func minimumRecolors(blocks string, k int) int {
	var count int
	res := math.MaxInt
	for i := 0; i < len(blocks); i++ {
		if blocks[i] == 'W' {
			count++
		}

		if i < k {
			res = count
			continue
		}
		if blocks[i-k] == 'W' {
			count--
		}

		res = min(count, res)
	}
	return res
}
